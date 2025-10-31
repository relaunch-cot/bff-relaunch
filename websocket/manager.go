package websocket

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID       string
	UserID   string
	Conn     *websocket.Conn
	Send     chan []byte
	Manager  *Manager
	ChatRoom string
	closed   bool
	mu       sync.Mutex
}

type Manager struct {
	clients map[string]*Client

	chatRooms map[string]map[string]*Client

	subscriptions map[string]map[string]bool

	register chan *Client

	unregister chan *Client

	broadcast chan *BroadcastMessage

	mu sync.RWMutex
}

type BroadcastMessage struct {
	UserID  string
	ChatID  string
	Message []byte
	Type    string
}

func NewManager() *Manager {
	return &Manager{
		clients:       make(map[string]*Client),
		chatRooms:     make(map[string]map[string]*Client),
		subscriptions: make(map[string]map[string]bool),
		register:      make(chan *Client),
		unregister:    make(chan *Client),
		broadcast:     make(chan *BroadcastMessage, 256),
	}
}

func (m *Manager) Run() {
	for {
		select {
		case client := <-m.register:
			m.registerClient(client)

		case client := <-m.unregister:
			m.unregisterClient(client)

		case message := <-m.broadcast:
			m.broadcastMessage(message)
		}
	}
}

func (m *Manager) registerClient(client *Client) {
	m.mu.Lock()

	if existingClient, exists := m.clients[client.UserID]; exists {
		existingClient.closeChannel()
		existingClient.Conn.Close()
		log.Printf("Disconnecting existing client: %s (UserID: %s)", existingClient.ID, existingClient.UserID)
	}

	m.clients[client.UserID] = client
	log.Printf("Client registered: %s (UserID: %s)", client.ID, client.UserID)

	shouldNotifyStatus := false
	var chatRoomToNotify string
	isPresenceConnection := client.ChatRoom == "" && client.Manager == PresenceManager

	if client.ChatRoom != "" {
		if _, exists := m.chatRooms[client.ChatRoom]; !exists {
			m.chatRooms[client.ChatRoom] = make(map[string]*Client)
		}
		m.chatRooms[client.ChatRoom][client.UserID] = client
		log.Printf("Client %s added to chat room %s", client.UserID, client.ChatRoom)

		shouldNotifyStatus = true
		chatRoomToNotify = client.ChatRoom
	}

	m.mu.Unlock()

	if shouldNotifyStatus {
		m.sendExistingParticipantsStatus(client, chatRoomToNotify)
		m.notifyUserStatus(chatRoomToNotify, client.UserID, true, client.UserID)
	}

	if isPresenceConnection {
		m.broadcastUserPresence(client.UserID, true, client.UserID)
	}
}

func (m *Manager) unregisterClient(client *Client) {
	m.mu.Lock()

	if _, exists := m.clients[client.UserID]; exists {
		delete(m.clients, client.UserID)
		client.closeChannel()
		log.Printf("Client unregistered: %s (UserID: %s)", client.ID, client.UserID)
	}

	isPresenceConnection := client.ChatRoom == "" && client.Manager == PresenceManager

	if client.ChatRoom != "" {
		if room, exists := m.chatRooms[client.ChatRoom]; exists {
			delete(room, client.UserID)
			if len(room) == 0 {
				delete(m.chatRooms, client.ChatRoom)
			}
		}

		chatRoom := client.ChatRoom
		userID := client.UserID
		m.mu.Unlock()
		m.notifyUserStatus(chatRoom, userID, false, userID)
		return
	}

	m.mu.Unlock()

	if isPresenceConnection {
		m.UnsubscribeFromAll(client.UserID)
		m.broadcastUserPresence(client.UserID, false, client.UserID)
	}
}

func (m *Manager) broadcastMessage(message *BroadcastMessage) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	switch message.Type {
	case "notification":
		if client, exists := m.clients[message.UserID]; exists {
			select {
			case client.Send <- message.Message:
			default:
				log.Printf("Failed to send message to client %s", message.UserID)
			}
		}
	case "chat":
		if room, exists := m.chatRooms[message.ChatID]; exists {
			for _, client := range room {
				select {
				case client.Send <- message.Message:
				default:
					log.Printf("Failed to send message to client %s in chat %s", client.UserID, message.ChatID)
				}
			}
		}
	}
}

func (m *Manager) SendToUser(userID string, message []byte) {
	m.broadcast <- &BroadcastMessage{
		UserID:  userID,
		Message: message,
		Type:    "notification",
	}
}

func (m *Manager) SendToChat(chatID string, message []byte) {
	m.broadcast <- &BroadcastMessage{
		ChatID:  chatID,
		Message: message,
		Type:    "chat",
	}
}

func (m *Manager) AddClientToChat(client *Client, chatID string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.chatRooms[chatID]; !exists {
		m.chatRooms[chatID] = make(map[string]*Client)
	}

	m.chatRooms[chatID][client.UserID] = client
	client.ChatRoom = chatID
	log.Printf("Client %s added to chat room %s", client.UserID, chatID)
}

func (m *Manager) IsUserOnline(userID string) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	_, exists := m.clients[userID]
	return exists
}

func (m *Manager) GetOnlineUsersCount() int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return len(m.clients)
}

func (c *Client) closeChannel() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.closed {
		close(c.Send)
		c.closed = true
	}
}

func (m *Manager) notifyUserStatus(chatID, userID string, isInChat bool, excludeUserID string) {
	m.mu.RLock()
	room, exists := m.chatRooms[chatID]
	if !exists {
		m.mu.RUnlock()
		return
	}

	statusMsg := map[string]interface{}{
		"type":     "USER_STATUS",
		"userId":   userID,
		"isInChat": isInChat,
		"chatId":   chatID,
	}

	data, err := json.Marshal(statusMsg)
	if err != nil {
		m.mu.RUnlock()
		log.Printf("Error marshaling user status: %v", err)
		return
	}

	for _, client := range room {
		if client.UserID != excludeUserID && client.ChatRoom == chatID {
			select {
			case client.Send <- data:
				log.Printf("User status sent: %s isInChat=%v in chat %s to user %s", userID, isInChat, chatID, client.UserID)
			default:
				log.Printf("Failed to send user status to client %s", client.UserID)
			}
		}
	}
	m.mu.RUnlock()
}

func (m *Manager) SendTypingIndicatorToChat(chatID, userID string, isTyping bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	room, exists := m.chatRooms[chatID]
	if !exists {
		log.Printf("Chat room '%s' not found when sending typing indicator", chatID)
		return
	}

	typingMsg := map[string]interface{}{
		"type":     "USER_TYPING",
		"userId":   userID,
		"isTyping": isTyping,
		"chatId":   chatID,
	}

	data, err := json.Marshal(typingMsg)
	if err != nil {
		log.Printf("Error marshaling typing indicator: %v", err)
		return
	}

	sentCount := 0
	for _, client := range room {
		if client.UserID != userID {
			select {
			case client.Send <- data:
				sentCount++
				log.Printf("Sent USER_TYPING to %s: user %s is typing: %v in chat %s", client.UserID, userID, isTyping, chatID)
			default:
				log.Printf("Failed to send typing indicator to client %s", client.UserID)
			}
		}
	}

	if sentCount == 0 {
		log.Printf("No other participants to send typing indicator (chat %s, total in room: %d)", chatID, len(room))
	}
}

func (m *Manager) GetChatParticipants(chatID string) []string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	room, exists := m.chatRooms[chatID]
	if !exists {
		log.Printf("Chat room '%s' not found. Total rooms: %d", chatID, len(m.chatRooms))
		return []string{}
	}

	participants := make([]string, 0, len(room))
	for userID := range room {
		participants = append(participants, userID)
	}

	log.Printf("Chat room '%s' has %d participants: %v", chatID, len(participants), participants)
	return participants
}

func (m *Manager) sendExistingParticipantsStatus(newClient *Client, chatID string) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	room, exists := m.chatRooms[chatID]
	if !exists {
		return
	}

	for userID, existingClient := range room {
		if userID != newClient.UserID && existingClient.ChatRoom == chatID {
			statusMsg := map[string]interface{}{
				"type":     "USER_STATUS",
				"userId":   userID,
				"isInChat": true,
				"chatId":   chatID,
			}

			data, err := json.Marshal(statusMsg)
			if err != nil {
				log.Printf("Error marshaling existing participant status: %v", err)
				continue
			}

			select {
			case newClient.Send <- data:
				log.Printf("Sent existing participant status: %s is in chat %s (to new client %s)", userID, chatID, newClient.UserID)
			default:
				log.Printf("Failed to send existing participant status to new client %s", newClient.UserID)
			}
		}
	}
}

func (m *Manager) SubscribeToPresence(observerUserID string, targetUserIDs []string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, targetUserID := range targetUserIDs {
		if m.subscriptions[targetUserID] == nil {
			m.subscriptions[targetUserID] = make(map[string]bool)
		}
		m.subscriptions[targetUserID][observerUserID] = true
	}

	log.Printf("User %s subscribed to presence of %d users", observerUserID, len(targetUserIDs))
}

func (m *Manager) UnsubscribeFromPresence(observerUserID string, targetUserIDs []string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, targetUserID := range targetUserIDs {
		if observers, exists := m.subscriptions[targetUserID]; exists {
			delete(observers, observerUserID)
			if len(observers) == 0 {
				delete(m.subscriptions, targetUserID)
			}
		}
	}

	log.Printf("User %s unsubscribed from presence of %d users", observerUserID, len(targetUserIDs))
}

func (m *Manager) UnsubscribeFromAll(observerUserID string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for targetUserID, observers := range m.subscriptions {
		delete(observers, observerUserID)
		if len(observers) == 0 {
			delete(m.subscriptions, targetUserID)
		}
	}

	log.Printf("User %s unsubscribed from all presence", observerUserID)
}

func (m *Manager) SendPresenceStatusToClient(client *Client, targetUserIDs []string) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, targetUserID := range targetUserIDs {
		isOnline := false
		if _, exists := m.clients[targetUserID]; exists {
			isOnline = true
		}

		statusMsg := map[string]interface{}{
			"type":     "USER_ONLINE",
			"userId":   targetUserID,
			"isOnline": isOnline,
		}

		data, err := json.Marshal(statusMsg)
		if err != nil {
			log.Printf("Error marshaling presence status: %v", err)
			continue
		}

		select {
		case client.Send <- data:
			log.Printf("Sent initial presence status: %s is %v to %s", targetUserID, isOnline, client.UserID)
		default:
			log.Printf("Failed to send initial presence status to client %s", client.UserID)
		}
	}
}

func (m *Manager) broadcastUserPresence(userID string, isOnline bool, excludeUserID string) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	observers, hasObservers := m.subscriptions[userID]
	if !hasObservers || len(observers) == 0 {
		return
	}

	statusMsg := map[string]interface{}{
		"type":     "USER_ONLINE",
		"userId":   userID,
		"isOnline": isOnline,
	}

	data, err := json.Marshal(statusMsg)
	if err != nil {
		log.Printf("Error marshaling user presence status: %v", err)
		return
	}

	for observerUserID := range observers {
		if observerUserID != excludeUserID {
			if client, exists := m.clients[observerUserID]; exists {
				select {
				case client.Send <- data:
					log.Printf("Sent presence update: %s is %v to observer %s", userID, isOnline, observerUserID)
				default:
					log.Printf("Failed to send presence update to observer %s", observerUserID)
				}
			}
		}
	}
}

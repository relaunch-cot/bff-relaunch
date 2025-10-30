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
		clients:    make(map[string]*Client),
		chatRooms:  make(map[string]map[string]*Client),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan *BroadcastMessage, 256),
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
}

func (m *Manager) unregisterClient(client *Client) {
	m.mu.Lock()

	if _, exists := m.clients[client.UserID]; exists {
		delete(m.clients, client.UserID)
		client.closeChannel()
		log.Printf("Client unregistered: %s (UserID: %s)", client.ID, client.UserID)
	}

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

func (m *Manager) notifyUserStatus(chatID, userID string, isOnline bool, excludeUserID string) {
	m.mu.RLock()
	room, exists := m.chatRooms[chatID]
	if !exists {
		m.mu.RUnlock()
		return
	}

	statusMsg := map[string]interface{}{
		"type":     "USER_STATUS",
		"userId":   userID,
		"isOnline": isOnline,
		"chatId":   chatID,
	}

	data, err := json.Marshal(statusMsg)
	if err != nil {
		m.mu.RUnlock()
		log.Printf("Error marshaling user status: %v", err)
		return
	}

	for _, client := range room {
		if client.UserID != excludeUserID {
			select {
			case client.Send <- data:
				log.Printf("User status sent: %s is %v in chat %s", userID, isOnline, chatID)
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

	for userID := range room {
		if userID != newClient.UserID {
			statusMsg := map[string]interface{}{
				"type":     "USER_STATUS",
				"userId":   userID,
				"isOnline": true,
				"chatId":   chatID,
			}

			data, err := json.Marshal(statusMsg)
			if err != nil {
				log.Printf("Error marshaling existing participant status: %v", err)
				continue
			}

			select {
			case newClient.Send <- data:
				log.Printf("Sent existing participant status: %s is online in chat %s (to new client %s)", userID, chatID, newClient.UserID)
			default:
				log.Printf("Failed to send existing participant status to new client %s", newClient.UserID)
			}
		}
	}
}

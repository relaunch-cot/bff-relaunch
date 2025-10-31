package websocket

import (
	"encoding/json"
	"log"
)

type NotificationMessage struct {
	Type         string                 `json:"type"`
	Notification map[string]interface{} `json:"notification,omitempty"`
}

type ChatMessage struct {
	Type    string                 `json:"type"`
	Message map[string]interface{} `json:"message,omitempty"`
}

func SendNewNotification(userID string, notification map[string]interface{}) {
	if NotificationManager == nil {
		log.Println("NotificationManager not initialized")
		return
	}

	msg := NotificationMessage{
		Type:         "NEW_NOTIFICATION",
		Notification: notification,
	}

	data, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Error marshaling notification: %v", err)
		return
	}

	NotificationManager.SendToUser(userID, data)
	log.Printf("Notification sent to user %s via WebSocket", userID)
}

func SendNotificationDeleted(userID string, notificationID string) {
	if NotificationManager == nil {
		log.Println("NotificationManager not initialized")
		return
	}

	msg := map[string]interface{}{
		"type":           "NOTIFICATION_DELETED",
		"notificationId": notificationID,
	}

	data, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Error marshaling notification deleted: %v", err)
		return
	}

	NotificationManager.SendToUser(userID, data)
	log.Printf("Notification deleted sent to user %s via WebSocket", userID)
}

func SendBadgeUpdate(userID string, count int) {
	if NotificationManager == nil {
		log.Println("NotificationManager not initialized")
		return
	}

	msg := map[string]interface{}{
		"type":  "BADGE_UPDATE",
		"count": count,
	}

	data, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Error marshaling badge update: %v", err)
		return
	}

	NotificationManager.SendToUser(userID, data)
	log.Printf("Badge update sent to user %s via WebSocket", userID)
}

func SendNewChatMessage(chatID string, message map[string]interface{}) {
	if ChatManager == nil {
		log.Println("ChatManager not initialized")
		return
	}

	msg := ChatMessage{
		Type:    "NEW_MESSAGE",
		Message: message,
	}

	data, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Error marshaling chat message: %v", err)
		return
	}

	ChatManager.SendToChat(chatID, data)
	log.Printf("Chat message sent to chat %s via WebSocket", chatID)
}

func SendTypingIndicator(chatID string, userID string, isTyping bool) {
	if ChatManager == nil {
		log.Println("ChatManager not initialized")
		return
	}

	msg := map[string]interface{}{
		"type":     "TYPING_INDICATOR",
		"chatId":   chatID,
		"userId":   userID,
		"isTyping": isTyping,
	}

	data, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Error marshaling typing indicator: %v", err)
		return
	}

	ChatManager.SendToChat(chatID, data)
}

func IsUserOnline(userID string) bool {
	if PresenceManager == nil {
		return false
	}
	return PresenceManager.IsUserOnline(userID)
}

func IsUserOnlineInChat(userID string) bool {
	if ChatManager == nil {
		return false
	}
	return ChatManager.IsUserOnline(userID)
}

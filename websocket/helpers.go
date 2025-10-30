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
	if WSManager == nil {
		log.Println("WebSocket manager not initialized")
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

	WSManager.SendToUser(userID, data)
	log.Printf("Notification sent to user %s via WebSocket", userID)
}

func SendNotificationDeleted(userID string, notificationID string) {
	if WSManager == nil {
		log.Println("WebSocket manager not initialized")
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

	WSManager.SendToUser(userID, data)
	log.Printf("Notification deleted sent to user %s via WebSocket", userID)
}

func SendBadgeUpdate(userID string, count int) {
	if WSManager == nil {
		log.Println("WebSocket manager not initialized")
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

	WSManager.SendToUser(userID, data)
	log.Printf("Badge update sent to user %s via WebSocket", userID)
}

func SendNewChatMessage(chatID string, message map[string]interface{}) {
	if WSManager == nil {
		log.Println("WebSocket manager not initialized")
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

	WSManager.SendToChat(chatID, data)
	log.Printf("Chat message sent to chat %s via WebSocket", chatID)
}

func SendTypingIndicator(chatID string, userID string, isTyping bool) {
	if WSManager == nil {
		log.Println("WebSocket manager not initialized")
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

	WSManager.SendToChat(chatID, data)
}

func IsUserOnline(userID string) bool {
	if WSManager == nil {
		return false
	}
	return WSManager.IsUserOnline(userID)
}

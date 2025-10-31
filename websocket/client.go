package websocket

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait = 10 * time.Second

	pongWait = 60 * time.Second

	pingPeriod = (pongWait * 9) / 10

	maxMessageSize = 512
)

type Message struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data,omitempty"`
}

func (c *Client) ReadPump() {
	defer func() {
		c.Manager.unregister <- c
		c.Conn.Close()
	}()

	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetPongHandler(func(string) error {
		c.Conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket error: %v", err)
			}
			break
		}

		c.handleMessage(message)
	}
}

func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			n := len(c.Send)
			for i := 0; i < n; i++ {
				w.Write([]byte{'\n'})
				w.Write(<-c.Send)
			}

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *Client) handleMessage(data []byte) {
	var msg Message
	if err := json.Unmarshal(data, &msg); err != nil {
		log.Printf("Error parsing message: %v", err)
		return
	}

	switch msg.Type {
	case "PING":
		response := Message{
			Type: "PONG",
		}
		c.sendMessage(response)

	case "JOIN_CHAT":
		if chatID, ok := msg.Data["chatId"].(string); ok {
			c.Manager.AddClientToChat(c, chatID)
		}

	case "TYPING":
		if c.ChatRoom != "" {
			isTyping := true
			if typing, ok := msg.Data["isTyping"].(bool); ok {
				isTyping = typing
			}

			c.Manager.SendTypingIndicatorToChat(c.ChatRoom, c.UserID, isTyping)
			log.Printf("Typing indicator from user %s in chat %s (typing: %v)", c.UserID, c.ChatRoom, isTyping)
		}

	case "SEND_MESSAGE":
		log.Printf("Received SEND_MESSAGE from user %s - messages should be sent via REST API", c.UserID)

		response := Message{
			Type: "ERROR",
			Data: map[string]interface{}{
				"message": "Use POST /v1/chat/send-message/:senderId to send messages",
				"code":    "USE_REST_API",
			},
		}
		c.sendMessage(response)

	case "SUBSCRIBE_PRESENCE":
		if userIds, ok := msg.Data["userIds"].([]interface{}); ok {
			targetUserIDs := make([]string, 0, len(userIds))
			for _, uid := range userIds {
				if userID, ok := uid.(string); ok {
					targetUserIDs = append(targetUserIDs, userID)
				}
			}

			if len(targetUserIDs) > 0 {
				c.Manager.SubscribeToPresence(c.UserID, targetUserIDs)
				c.Manager.SendPresenceStatusToClient(c, targetUserIDs)
			}
		}

	case "UNSUBSCRIBE_PRESENCE":
		if userIds, ok := msg.Data["userIds"].([]interface{}); ok {
			targetUserIDs := make([]string, 0, len(userIds))
			for _, uid := range userIds {
				if userID, ok := uid.(string); ok {
					targetUserIDs = append(targetUserIDs, userID)
				}
			}

			if len(targetUserIDs) > 0 {
				c.Manager.UnsubscribeFromPresence(c.UserID, targetUserIDs)
			}
		}

	default:
		log.Printf("Unknown message type: %s", msg.Type)
	}
}

func (c *Client) sendMessage(msg Message) {
	data, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Error marshaling message: %v", err)
		return
	}

	select {
	case c.Send <- data:
	default:
		log.Printf("Client %s send channel full, dropping message", c.UserID)
	}
}

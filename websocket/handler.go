package websocket

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
	"github.com/relaunch-cot/bff-relaunch/config"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func validateToken(tokenString string) (string, error) {
	secret := config.JWT_SECRET
	if strings.TrimSpace(secret) == "" {
		return "", jwt.ErrSignatureInvalid
	}

	if strings.HasPrefix(strings.ToLower(tokenString), "bearer ") {
		tokenString = strings.TrimSpace(tokenString[7:])
	}

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

	if exp, ok := claims["exp"].(float64); ok {
		_ = exp
	}

	var userId string
	if v, ok := claims["userId"].(string); ok {
		userId = v
	} else if v, ok := claims["user_id"].(string); ok {
		userId = v
	} else {
		return "", jwt.ErrTokenInvalidClaims
	}

	return userId, nil
}

func HandleWebSocketNotifications(manager *Manager) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "token is required"})
			return
		}

		userId, err := validateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid token", "details": err.Error()})
			return
		}

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Printf("Failed to upgrade connection: %v", err)
			return
		}

		client := &Client{
			ID:      generateClientID(),
			UserID:  userId,
			Conn:    conn,
			Send:    make(chan []byte, 256),
			Manager: manager,
		}

		manager.register <- client

		welcomeMsg := map[string]interface{}{
			"type":    "CONNECTED",
			"message": "Connected to notification service",
			"userId":  userId,
		}
		welcomeData, _ := json.Marshal(welcomeMsg)
		client.Send <- welcomeData

		go client.WritePump()
		go client.ReadPump()
	}
}

func HandleWebSocketChat(manager *Manager) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		chatId := c.Query("chatId")

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "token is required"})
			return
		}

		if chatId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "chatId is required"})
			return
		}

		userId, err := validateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid token", "details": err.Error()})
			return
		}

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Printf("Failed to upgrade connection: %v", err)
			return
		}

		client := &Client{
			ID:      generateClientID(),
			UserID:  userId,
			Conn:    conn,
			Send:    make(chan []byte, 256),
			Manager: manager,
		}

		manager.register <- client

		manager.AddClientToChat(client, chatId)

		welcomeMsg := map[string]interface{}{
			"type":    "CONNECTED",
			"message": "Connected to chat service",
			"userId":  userId,
			"chatId":  chatId,
		}
		welcomeData, _ := json.Marshal(welcomeMsg)
		client.Send <- welcomeData

		go client.WritePump()
		go client.ReadPump()
	}
}

func generateClientID() string {
	timestamp := time.Now().UnixNano()
	randomBytes := make([]byte, 8)
	if _, err := rand.Read(randomBytes); err != nil {
		return fmt.Sprintf("%d", timestamp)
	}

	return fmt.Sprintf("%d-%x", timestamp, randomBytes)
}

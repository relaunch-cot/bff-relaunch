package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/relaunch-cot/bff-relaunch/config"
	"google.golang.org/grpc/status"
)

func ValidateUserToken(c *gin.Context) {
	tokenHeader := c.GetHeader("Authorization")
	if strings.TrimSpace(tokenHeader) == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Header 'Authorization' is required"})
		return
	}

	tokenString := strings.TrimSpace(tokenHeader)
	if strings.HasPrefix(strings.ToLower(tokenString), "bearer ") {
		tokenString = strings.TrimSpace(tokenString[7:])
	}

	if tokenString == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Authorization token is empty"})
		return
	}

	secret := config.JWT_SECRET
	if strings.TrimSpace(secret) == "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "JWT secret not configured"})
		return
	}

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		// garantir método de assinatura esperado
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, status.Error(http.StatusInternalServerError, "unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token", "details": err.Error()})
		return
	}

	var userId interface{}
	if v, ok := claims["userId"]; ok {
		userId = v
	} else if v, ok := claims["user_id"]; ok {
		userId = v
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "token does not contain userId"})
		return
	}

	// coloca userId no contexto para handlers subsequentes
	c.Set("userId", userId)
	c.Next()
}

package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/relaunch-cot/bff-relaunch/config"
	"github.com/relaunch-cot/bff-relaunch/resource"
	"github.com/relaunch-cot/bff-relaunch/routes"
	"github.com/relaunch-cot/bff-relaunch/websocket"
)

func main() {
	resource.Inject()

	websocket.InitializeWebSocket()

	r := gin.Default()

	corsConfig := cors.Config{
		AllowOrigins:     []string{"https://relaunch-cot.netlify.app", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}

	r.Use(cors.New(corsConfig))

	routes.AddRoutes(r.Group(""))

	err := r.Run(":" + config.PORT)
	if err != nil {
		panic("[Error] failed to start Gin server due to : " + err.Error())
	}
}

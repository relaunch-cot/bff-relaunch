package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/relaunch-cot/bff/config"
	"github.com/relaunch-cot/bff/resource"
	"github.com/relaunch-cot/bff/routes"
)

func main() {
	resource.Inject()

	r := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true

	r.Use(cors.New(corsConfig))

	routes.AddRoutes(r.Group(""))

	err := r.Run(":" + config.PORT)
	if err != nil {
		panic("[Error] failed to start Gin server due to : " + err.Error())
	}
}

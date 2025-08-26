package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/relaunch-cot/bff-relaunch/resource"
)

func AddRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1")

	user := v1.Group("/user")
	user.POST("/register", resource.Servers.User.CreateUser)
	user.POST("/login", resource.Servers.User.LoginUser)
}

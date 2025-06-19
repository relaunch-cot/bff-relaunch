package routes

import (
	"bff.com/m/resource"
	"github.com/gin-gonic/gin"
)

func addRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1")

	user := v1.Group("/user")
	user.POST("", resource.Servers)
}

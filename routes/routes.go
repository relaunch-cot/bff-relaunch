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
	user.PUT("/:id", resource.Servers.User.UpdateUser)
	user.PATCH("", resource.Servers.User.UpdateUserPassword)
	user.DELETE("", resource.Servers.User.DeleteUser)
	user.POST("/send-email", resource.Servers.User.SendPasswordRecoveryEmail)

	reports := v1.Group("/reports")
	reports.POST("/generate-pdf", resource.Servers.User.GenerateReportPDF)

	chat := v1.Group("/chat")
	chat.POST("", resource.Servers.User.CreateNewChat)
	chat.POST("send-message/:senderId", resource.Servers.User.SendMessage)
	chat.GET("messages/:chatId", resource.Servers.User.GetAllMessagesFromChat)
}

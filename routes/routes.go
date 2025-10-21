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
	user.GET("/:userId", resource.Servers.User.GetUserProfile)
	user.GET("/userType/:userId", resource.Servers.User.GetUserType)

	reports := v1.Group("/reports")
	reports.POST("/generate-pdf", resource.Servers.User.GenerateReportPDF)

	chat := v1.Group("/chat")
	chat.POST("", resource.Servers.Chat.CreateNewChat)
	chat.POST("send-message/:senderId", resource.Servers.Chat.SendMessage)
	chat.GET("messages/:chatId", resource.Servers.Chat.GetAllMessagesFromChat)
	chat.GET("/:userId", resource.Servers.Chat.GetAllChatsFromUser)

	project := v1.Group("/project")
	project.POST("/:userId", resource.Servers.Project.CreateProject)
	project.GET("/:projectId", resource.Servers.Project.GetProject)
	project.GET("user/:userId", resource.Servers.Project.GetAllProjectsFromUser)
	project.PUT("/:projectId", resource.Servers.Project.UpdateProject)
}

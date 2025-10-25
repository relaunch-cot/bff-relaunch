package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/relaunch-cot/bff-relaunch/middleware"
	"github.com/relaunch-cot/bff-relaunch/resource"
)

func AddRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1")

	user := v1.Group("/user")
	user.POST("/register", resource.Servers.User.CreateUser)
	user.POST("/login", resource.Servers.User.LoginUser)
	user.PUT("/:id", middleware.ValidateUserToken, resource.Servers.User.UpdateUser)
	user.PATCH("", resource.Servers.User.UpdateUserPassword)
	user.DELETE("", middleware.ValidateUserToken, resource.Servers.User.DeleteUser)
	user.POST("/send-email", resource.Servers.User.SendPasswordRecoveryEmail)
	user.GET("/:userId", middleware.ValidateUserToken, resource.Servers.User.GetUserProfile)

	reports := v1.Group("/reports")
	reports.POST("/generate-pdf", middleware.ValidateUserToken, resource.Servers.User.GenerateReportPDF)

	chat := v1.Group("/chat")
	chat.POST("", middleware.ValidateUserToken, resource.Servers.Chat.CreateNewChat)
	chat.POST("/send-message/:senderId", middleware.ValidateUserToken, resource.Servers.Chat.SendMessage)
	chat.GET("/messages/:chatId", middleware.ValidateUserToken, resource.Servers.Chat.GetAllMessagesFromChat)
	chat.GET("/:userId", middleware.ValidateUserToken, resource.Servers.Chat.GetAllChatsFromUser)
	chat.GET("", middleware.ValidateUserToken, resource.Servers.Chat.GetChatFromUsers)

	project := v1.Group("/project")
	project.POST("/:userId", middleware.ValidateUserToken, resource.Servers.Project.CreateProject)
	project.GET("/:projectId", middleware.ValidateUserToken, resource.Servers.Project.GetProject)
	project.GET("/user/:userId", middleware.ValidateUserToken, resource.Servers.Project.GetAllProjectsFromUser)
	project.PUT("/:projectId", middleware.ValidateUserToken, resource.Servers.Project.UpdateProject)
	project.PATCH("/add-freelancer/:projectId", middleware.ValidateUserToken, resource.Servers.Project.AddFreelancerToProject)
	project.PATCH("/remove-freelancer/:projectId", middleware.ValidateUserToken, resource.Servers.Project.RemoveFreelancerFromProject)
	project.GET("", middleware.ValidateUserToken, resource.Servers.Project.GetAllProjects)

	notification := v1.Group("/notification")
	notification.POST("/:senderId", middleware.ValidateUserToken, resource.Servers.Notification.SendNotification)
	notification.GET("/:notificationId", middleware.ValidateUserToken, resource.Servers.Notification.GetNotification)
	notification.GET("/user/:userId", middleware.ValidateUserToken, resource.Servers.Notification.GetAllNotificationsFromUser)
	notification.DELETE("/:notificationId", middleware.ValidateUserToken, resource.Servers.Notification.DeleteNotification)
	notification.DELETE("/user/:userId", middleware.ValidateUserToken, resource.Servers.Notification.DeleteAllNotificationsFromUser)
}

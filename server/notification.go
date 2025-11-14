package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/relaunch-cot/bff-relaunch/handler"
	params "github.com/relaunch-cot/bff-relaunch/params/notification"
	"github.com/relaunch-cot/bff-relaunch/resource/transformer"
	"github.com/relaunch-cot/bff-relaunch/websocket"
	"github.com/relaunch-cot/lib-relaunch-cot/pkg/httpresponse"
	validation "github.com/relaunch-cot/lib-relaunch-cot/validate/notification"
)

type INotification interface {
	SendNotification(c *gin.Context)
	GetNotification(c *gin.Context)
	GetAllNotificationsFromUser(c *gin.Context)
	DeleteNotification(c *gin.Context)
	DeleteAllNotificationsFromUser(c *gin.Context)
}

func (r *resource) SendNotification(c *gin.Context) {
	senderId := c.Param("senderId")
	if senderId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "senderId is required"})
		return
	}

	in := new(params.SendNotificationPOST)
	if err := c.Bind(in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error getting body of the request"})
		return
	}

	sendNotificationRequest, err := transformer.SendNotificationToProto(senderId, in.ReceiverId, in.Title, in.Content, in.Type)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error transforming params to proto"})
		return
	}

	err = validation.ValidateSendNotificationRequest(sendNotificationRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error validating the body of the request. Details:" + err.Error()})
		return
	}

	ctx := c.Request.Context()

	err = r.handler.Notification.SendNotification(&ctx, sendNotificationRequest)
	if err != nil {
		c.JSON(httpresponse.TransformGrpcCodeToHttpStatus(err), gin.H{"message": err.Error()})
		return
	}

	go func() {
		notification := map[string]interface{}{
			"senderId":   senderId,
			"receiverId": in.ReceiverId,
			"title":      in.Title,
			"content":    in.Content,
			"type":       in.Type,
		}
		websocket.SendNewNotification(in.ReceiverId, notification)
	}()

	c.JSON(http.StatusOK, gin.H{"message": "notification sent successfully"})
}
func (r *resource) GetNotification(c *gin.Context) {
	notificationId := c.Param("notificationId")
	if notificationId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "userId is required"})
		return
	}

	getNotificationRequest, err := transformer.GetNotificationToProto(notificationId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error transforming params to proto"})
		return
	}

	err = validation.ValidateGetNotificationRequest(getNotificationRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error validating the body of the request. Details:" + err.Error()})
		return
	}

	ctx := c.Request.Context()

	getNotificationResponse, err := r.handler.Notification.GetNotification(&ctx, getNotificationRequest)
	if err != nil {
		c.JSON(httpresponse.TransformGrpcCodeToHttpStatus(err), gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, getNotificationResponse)
}

func (r *resource) GetAllNotificationsFromUser(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "userId is required"})
		return
	}

	getAllNotificationsFromUserRequest, err := transformer.GetAllNotificationsFromUserToProto(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error transforming params to proto"})
		return
	}

	err = validation.ValidateGetAllNotificationsFromUserRequest(getAllNotificationsFromUserRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error validating the body of the request. Details:" + err.Error()})
		return
	}

	ctx := c.Request.Context()

	getAllNotificationsFromUserResponse, err := r.handler.Notification.GetAllNotificationsFromUser(&ctx, getAllNotificationsFromUserRequest)
	if err != nil {
		c.JSON(httpresponse.TransformGrpcCodeToHttpStatus(err), gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, getAllNotificationsFromUserResponse)
}

func (r *resource) DeleteNotification(c *gin.Context) {
	notificationId := c.Param("notificationId")
	if notificationId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "notificationId is required"})
		return
	}

	getNotificationRequest, err := transformer.GetNotificationToProto(notificationId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error transforming params to proto"})
		return
	}

	err = validation.ValidateGetNotificationRequest(getNotificationRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error validating the body of the request. Details:" + err.Error()})
		return
	}

	ctx := c.Request.Context()

	getNotificationResponse, err := r.handler.Notification.GetNotification(&ctx, getNotificationRequest)
	var userId string
	if err == nil && getNotificationResponse != nil {
		userId = getNotificationResponse.Notification.ReceiverId
	}

	deleteNotificationRequest, err := transformer.DeleteNotificationToProto(notificationId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error transforming params to proto"})
		return
	}

	err = validation.ValidateDeleteNotificationRequest(deleteNotificationRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error validating the body of the request. Details:" + err.Error()})
		return
	}

	err = r.handler.Notification.DeleteNotification(&ctx, deleteNotificationRequest)
	if err != nil {
		c.JSON(httpresponse.TransformGrpcCodeToHttpStatus(err), gin.H{"message": err.Error()})
		return
	}

	if userId != "" {
		go func() {
			websocket.SendNotificationDeleted(userId, notificationId)
		}()
	}

	c.JSON(http.StatusOK, gin.H{"message": "notification deleted successfully"})
}

func (r *resource) DeleteAllNotificationsFromUser(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "userId is required"})
		return
	}

	deleteAllNotificationsFromUserRequest, err := transformer.DeleteAllNotificationsFromUserToProto(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error transforming params to proto"})
		return
	}

	err = validation.ValidateDeleteAllNotificationsFromUserRequest(deleteAllNotificationsFromUserRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error validating the body of the request. Details:" + err.Error()})
		return
	}

	ctx := c.Request.Context()

	err = r.handler.Notification.DeleteAllNotificationsFromUser(&ctx, deleteAllNotificationsFromUserRequest)
	if err != nil {
		c.JSON(httpresponse.TransformGrpcCodeToHttpStatus(err), gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "all notifications from user deleted successfully"})
}

func NewNotificationServer(handler *handler.Handlers) INotification {
	return &resource{
		handler: handler,
	}
}

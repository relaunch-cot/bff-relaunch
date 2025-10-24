package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/relaunch-cot/bff-relaunch/handler"
	params "github.com/relaunch-cot/bff-relaunch/params/notification"
	"github.com/relaunch-cot/bff-relaunch/resource/transformer"
	"github.com/relaunch-cot/lib-relaunch-cot/pkg/httpresponse"
	validate "github.com/relaunch-cot/lib-relaunch-cot/validate/notification"
)

type INotification interface {
	SendNotification(c *gin.Context)
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

	isValidNotificationType := validate.IsValidNotificationType(in.Type)
	if !isValidNotificationType {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid notification type"})
		return
	}

	sendNotificationRequest, err := transformer.SendNotificationToProto(senderId, in.ReceiverId, in.Title, in.Content, in.Type)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error transforming params to proto"})
		return
	}

	ctx := c.Request.Context()

	err = r.handler.Notification.SendNotification(&ctx, sendNotificationRequest)
	if err != nil {
		c.JSON(httpresponse.TransformGrpcCodeToHttpStatus(err), gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "notification sent successfully"})
}

func NewNotificationServer(handler *handler.Handlers) INotification {
	return &resource{
		handler: handler,
	}
}

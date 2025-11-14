package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/relaunch-cot/bff-relaunch/handler"
	params "github.com/relaunch-cot/bff-relaunch/params/chat"
	"github.com/relaunch-cot/bff-relaunch/resource/transformer"
	"github.com/relaunch-cot/bff-relaunch/websocket"
	"github.com/relaunch-cot/lib-relaunch-cot/pkg/httpresponse"
	validation "github.com/relaunch-cot/lib-relaunch-cot/validate/chat"
)

type IChat interface {
	CreateNewChat(c *gin.Context)
	SendMessage(c *gin.Context)
	GetAllMessagesFromChat(c *gin.Context)
	GetAllChatsFromUser(c *gin.Context)
	GetChatFromUsers(c *gin.Context)
	GetChatById(c *gin.Context)
}

func (r *resource) CreateNewChat(c *gin.Context) {
	in := new(params.CreateNewChatPOST)
	err := c.Bind(in)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error getting query params"})
		return
	}

	if len(in.UserIds) < 2 || len(in.UserIds) > 2 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "must be two userIds"})
		return
	}

	createNewChatRequest, err := transformer.CreateNewChatToProto(in.UserIds, in.CreatedBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error transforming params to proto"})
		return
	}

	err = validation.ValidateCreateNewChatRequest(createNewChatRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error validating the body of the request. Details:" + err.Error()})
		return
	}

	ctx := c.Request.Context()
	err = r.handler.Chat.CreateNewChat(&ctx, createNewChatRequest)
	if err != nil {
		c.JSON(httpresponse.TransformGrpcCodeToHttpStatus(err), gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "chat created successfully"})
}

func (r *resource) SendMessage(c *gin.Context) {
	senderId := c.Param("senderId")
	if senderId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "the path param senderId is required"})
		return
	}

	in := new(params.SendMessagePOST)
	err := c.Bind(in)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error getting query params"})
		return
	}

	sendMessageRequest, err := transformer.SendMessageToProto(in.ChatId, senderId, in.MessageContent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error transforming params to proto"})
		return
	}

	err = validation.ValidateSendMessageRequest(sendMessageRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error validating the body of the request. Details:" + err.Error()})
		return
	}

	ctx := c.Request.Context()

	err = r.handler.Chat.SendMessage(&ctx, sendMessageRequest)
	if err != nil {
		c.JSON(httpresponse.TransformGrpcCodeToHttpStatus(err), gin.H{"message": err.Error()})
		return
	}

	go func() {
		message := map[string]interface{}{
			"chatId":         in.ChatId,
			"senderId":       senderId,
			"messageContent": in.MessageContent,
		}
		websocket.SendNewChatMessage(in.ChatId, message)
	}()

	c.JSON(http.StatusCreated, gin.H{"message": "message sent successfully"})
}

func (r *resource) GetAllMessagesFromChat(c *gin.Context) {
	chatId := c.Param("chatId")
	if chatId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "the path param chatId is required"})
		return
	}

	getAllMessagesFromChatRequest, err := transformer.GetAllMessagesFromChatToProto(chatId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error transforming params to proto"})
		return
	}

	err = validation.ValidateGetAllMessagesFromChatRequest(getAllMessagesFromChatRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error validating the body of the request. Details:" + err.Error()})
		return
	}

	ctx := c.Request.Context()

	getAllMessagesFromChatResponse, err := r.handler.Chat.GetAllMessagesFromChat(&ctx, getAllMessagesFromChatRequest)
	if err != nil {
		c.JSON(httpresponse.TransformGrpcCodeToHttpStatus(err), gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, getAllMessagesFromChatResponse)
}

func (r *resource) GetAllChatsFromUser(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "the path param userId is required"})
		return
	}

	getAllChatsFromUserRequest, err := transformer.GetAllChatsFromUserToProto(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error transforming params to proto"})
		return
	}

	err = validation.ValidateGetAllChatsFromUserRequest(getAllChatsFromUserRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error validating the body of the request. Details:" + err.Error()})
		return
	}

	ctx := c.Request.Context()

	getAllChatsFromUserResponse, err := r.handler.Chat.GetAllChatsFromUser(&ctx, getAllChatsFromUserRequest)
	if err != nil {
		c.JSON(httpresponse.TransformGrpcCodeToHttpStatus(err), gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, getAllChatsFromUserResponse)
}

func (r *resource) GetChatFromUsers(c *gin.Context) {
	in := new(params.GetChatFromUsersGET)
	err := c.BindQuery(in)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error getting query params"})
		return
	}

	var userIds []string
	userIds = append(userIds, in.User1Id, in.User2Id)

	if len(userIds) < 2 || len(userIds) > 2 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "must be two userIds"})
		return
	}

	getChatFromUsersRequest, err := transformer.GetChatFromUsersToProto(userIds)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error transforming params to proto"})
		return
	}

	err = validation.ValidateGetChatFromUsersRequest(getChatFromUsersRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error validating the body of the request. Details:" + err.Error()})
		return
	}

	ctx := c.Request.Context()

	getChatFromUsersResponse, err := r.handler.Chat.GetChatFromUsers(&ctx, getChatFromUsersRequest)
	if err != nil {
		c.JSON(httpresponse.TransformGrpcCodeToHttpStatus(err), gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, getChatFromUsersResponse)
}

func (r *resource) GetChatById(c *gin.Context) {
	chatId := c.Param("chatId")
	if chatId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "the path param chatId is required"})
		return
	}

	getChatByIdRequest, err := transformer.GetChatByIdToProto(chatId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error transforming params to proto"})
		return
	}

	err = validation.ValidateGetChatByIdRequest(getChatByIdRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error validating the body of the request. Details:" + err.Error()})
		return
	}

	ctx := c.Request.Context()

	getChatByIdResponse, err := r.handler.Chat.GetChatById(&ctx, getChatByIdRequest)
	if err != nil {
		c.JSON(httpresponse.TransformGrpcCodeToHttpStatus(err), gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, getChatByIdResponse)
}

func NewChatServer(handler *handler.Handlers) IChat {
	return &resource{
		handler: handler,
	}
}

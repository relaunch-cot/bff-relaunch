package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/relaunch-cot/bff-relaunch/handler"
	params "github.com/relaunch-cot/bff-relaunch/params/chat"
	"github.com/relaunch-cot/bff-relaunch/resource/transformer"
)

type IChat interface {
	CreateNewChat(c *gin.Context)
	SendMessage(c *gin.Context)
	GetAllMessagesFromChat(c *gin.Context)
	GetAllChatsFromUser(c *gin.Context)
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

	ctx := c.Request.Context()
	err = r.handler.Chat.CreateNewChat(&ctx, createNewChatRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
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

	senderIdInt, err := strconv.Atoi(senderId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "senderId must be a integer number"})
		return
	}

	in := new(params.SendMessagePOST)
	err = c.Bind(in)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error getting query params"})
		return
	}

	sendMessageRequest, err := transformer.SendMessageToProto(in.ChatId, int64(senderIdInt), in.MessageContent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error transforming params to proto"})
		return
	}

	ctx := c.Request.Context()

	err = r.handler.Chat.SendMessage(&ctx, sendMessageRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "message sent successfully"})
}

func (r *resource) GetAllMessagesFromChat(c *gin.Context) {
	chatId := c.Param("chatId")
	if chatId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "the path param chatId is required"})
		return
	}

	chatIdInt, err := strconv.Atoi(chatId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "chatId must be a integer number"})
		return
	}

	getAllMessagesFromChatRequest, err := transformer.GetAllMessagesFromChatToProto(int64(chatIdInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error transforming params to proto"})
		return
	}

	ctx := c.Request.Context()

	getAllMessagesFromChatResponse, err := r.handler.Chat.GetAllMessagesFromChat(&ctx, getAllMessagesFromChatRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
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

	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "userId must be a integer number"})
		return
	}

	getAllChatsFromUserRequest, err := transformer.GetAllChatsFromUserToProto(int64(userIdInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error transforming params to proto"})
		return
	}

	ctx := c.Request.Context()

	getAllChatsFromUserResponse, err := r.handler.Chat.GetAllChatsFromUser(&ctx, getAllChatsFromUserRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, getAllChatsFromUserResponse)
}

func NewChatServer(handler *handler.Handlers) IChat {
	return &resource{
		handler: handler,
	}
}

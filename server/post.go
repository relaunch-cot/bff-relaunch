package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/relaunch-cot/bff-relaunch/handler"
	params "github.com/relaunch-cot/bff-relaunch/params/post"
	"github.com/relaunch-cot/bff-relaunch/resource/transformer"
	"github.com/relaunch-cot/lib-relaunch-cot/pkg/httpresponse"
)

type IPost interface {
	CreatePost(c *gin.Context)
}

func (r *resource) CreatePost(c *gin.Context) {
	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting user id from token"})
		return
	}

	in := new(params.CreatePostPOST)
	err := c.Bind(in)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error getting body of the request"})
		return
	}

	createPostRequest, err := transformer.CreatePostToProto(userId.(string), in.Title, in.Content, in.Type, in.UrlImagePost)
	if err != nil {
		c.JSON(httpresponse.TransformGrpcCodeToHttpStatus(err), gin.H{"message": err.Error()})
		return
	}

	ctx := c.Request.Context()

	err = r.handler.Post.CreatePost(&ctx, createPostRequest)
	if err != nil {
		c.JSON(httpresponse.TransformGrpcCodeToHttpStatus(err), gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "post created successfully"})
}

func NewPostServer(handler *handler.Handlers) IPost {
	return &resource{
		handler: handler,
	}
}

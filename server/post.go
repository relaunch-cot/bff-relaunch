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
	GetPost(c *gin.Context)
	GetAllPosts(c *gin.Context)
	GetAllPostsFromUser(c *gin.Context)
	UpdatePost(c *gin.Context)
	DeletePost(c *gin.Context)
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
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
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

func (r *resource) GetPost(c *gin.Context) {
	postId := c.Param("postId")
	if postId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "postId is required"})
		return
	}

	getPostRequest, err := transformer.GetPostToProto(postId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx := c.Request.Context()

	response, err := r.handler.Post.GetPost(&ctx, getPostRequest)
	if err != nil {
		c.JSON(httpresponse.TransformGrpcCodeToHttpStatus(err), gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (r *resource) GetAllPosts(c *gin.Context) {
	ctx := c.Request.Context()

	response, err := r.handler.Post.GetAllPosts(&ctx)
	if err != nil {
		c.JSON(httpresponse.TransformGrpcCodeToHttpStatus(err), gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (r *resource) GetAllPostsFromUser(c *gin.Context) {
	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting user id from token"})
		return
	}

	getAllPostsFromUserRequest, err := transformer.GetAllPostsFromUserToProto(userId.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx := c.Request.Context()

	response, err := r.handler.Post.GetAllPostsFromUser(&ctx, getAllPostsFromUserRequest)
	if err != nil {
		c.JSON(httpresponse.TransformGrpcCodeToHttpStatus(err), gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (r *resource) UpdatePost(c *gin.Context) {
	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting user id from token"})
		return
	}

	postId := c.Param("postId")
	if postId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "postId is required"})
		return
	}

	in := new(params.UpdatePostPUT)
	err := c.Bind(in)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error getting body of the request"})
		return
	}

	updatePostRequest, err := transformer.UpdatePostToProto(userId.(string), postId, in.Title, in.Content, in.UrlImagePost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx := c.Request.Context()

	response, err := r.handler.Post.UpdatePost(&ctx, updatePostRequest)
	if err != nil {
		c.JSON(httpresponse.TransformGrpcCodeToHttpStatus(err), gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (r *resource) DeletePost(c *gin.Context) {
	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting user id from token"})
		return
	}

	postId := c.Param("postId")
	if postId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "postId is required"})
		return
	}

	deletePostRequest, err := transformer.DeletePostToProto(userId.(string), postId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx := c.Request.Context()

	err = r.handler.Post.DeletePost(&ctx, deletePostRequest)
	if err != nil {
		c.JSON(httpresponse.TransformGrpcCodeToHttpStatus(err), gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "post deleted successfully"})
}

func NewPostServer(handler *handler.Handlers) IPost {
	return &resource{
		handler: handler,
	}
}

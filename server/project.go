package server

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/relaunch-cot/bff-relaunch/handler"
	params "github.com/relaunch-cot/bff-relaunch/params/project"
	"github.com/relaunch-cot/bff-relaunch/resource/transformer"
	"github.com/relaunch-cot/lib-relaunch-cot/pkg/httpresponse"
)

type IProject interface {
	CreateProject(c *gin.Context)
	GetProject(c *gin.Context)
	GetAllProjectsFromUser(c *gin.Context)
}

func (r *resource) CreateProject(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "the Authorization header is required"})
		return
	}

	userId := c.Param("userId")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "userId is required"})
		return
	}

	in := new(params.CreateProjectPOST)
	if err := c.Bind(in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error getting body of the request"})
		return
	}

	createProjectRequest, err := transformer.CreateProjectToProto(userId, in.FreelancerId, in.Name, in.Description, in.Category, in.ProjectDeliveryDeadline.Format("2006-01-02 15:04:05"), in.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error transforming params to proto"})
		return
	}

	ctx := c.Request.Context()

	err = r.handler.Project.CreateProject(&ctx, createProjectRequest)
	if err != nil {
		c.JSON(httpresponse.TransformGrpcCodeToHttpStatus(err), gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "project created"})
}

func (r *resource) GetProject(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "the Authorization header is required"})
		return
	}

	projectId := c.Param("projectId")
	if projectId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "projectId is required"})
		return
	}

	getProjectRequest, err := transformer.GetProjectToProto(projectId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error transforming params to proto"})
		return
	}

	ctx := c.Request.Context()

	response, err := r.handler.Project.GetProject(&ctx, getProjectRequest)
	if err != nil {
		c.JSON(httpresponse.TransformGrpcCodeToHttpStatus(err), gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (r *resource) GetAllProjectsFromUser(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "the Authorization header is required"})
		return
	}

	userId := c.Param("userId")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "userId is required"})
		return
	}

	in := new(params.GetAllProjectsFromUserGET)
	err := c.BindQuery(in)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error getting query params"})
		return
	}
	in.UserType = strings.TrimSpace(in.UserType)
	if strings.ToLower(in.UserType) != "client" && strings.ToLower(in.UserType) != "freelancer" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid user type"})
		return
	}

	getAllProjectsFromUserRequest, err := transformer.GetAllProjectsFromUserToProto(userId, in.UserType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error transforming params to proto"})
		return
	}

	ctx := c.Request.Context()

	response, err := r.handler.Project.GetAllProjectsFromUser(&ctx, getAllProjectsFromUserRequest)
	if err != nil {
		c.JSON(httpresponse.TransformGrpcCodeToHttpStatus(err), gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func NewProjectServer(handler *handler.Handlers) IProject {
	return &resource{
		handler: handler,
	}
}

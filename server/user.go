package server

import (
	"github.com/relaunch-cot/bff-relaunch/handler"
	models "github.com/relaunch-cot/bff-relaunch/models/user"
	"github.com/relaunch-cot/bff-relaunch/resource/transformer"
	"net/http"

	params "github.com/relaunch-cot/bff-relaunch/params/user"

	"github.com/gin-gonic/gin"
)

type IUser interface {
	CreateUser(c *gin.Context)
	LoginUser(c *gin.Context)
	UpdateUserPassword(c *gin.Context)
}

func (r *resource) CreateUser(c *gin.Context) {
	in := new(params.CreateUserPOST)
	err := c.Bind(in)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error getting query params"})
	}

	user := &models.User{
		Name:     in.Name,
		Email:    in.Email,
		Password: in.Password,
	}

	createUserReq, err := transformer.CreateUserToProto(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error transforming params to proto"})
	}

	ctx := c.Request.Context()

	err = r.handler.User.CreateUser(&ctx, createUserReq)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "error calling user service"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user registered successfully"})
}

func (r *resource) LoginUser(c *gin.Context) {
	in := new(params.LoginUserGET)
	err := c.Bind(in)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error getting query params"})
	}

	user := &models.User{
		Email:    in.Email,
		Password: in.Password,
	}

	loginUserReq, err := transformer.LoginUserToProto(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error transforming params to proto"})
	}

	ctx := c.Request.Context()

	loginUserResponse, err := r.handler.User.LoginUser(&ctx, loginUserReq)
	if err != nil {
		c.JSON(404, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusOK, loginUserResponse)
}

func (r *resource) UpdateUserPassword(c *gin.Context) {
	in := new(params.UpdateUserPasswordPATCH)
	err := c.Bind(in)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error getting query params"})
		return
	}

	updateUserPasswordReq, err := transformer.UpdateUserPasswordToProto(in.Email, in.CurrentPassword, in.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error transforming params to proto"})
	}

	ctx := c.Request.Context()
	err = r.handler.User.UpdateUserPassword(&ctx, updateUserPasswordReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "password updated successfully"})
}

func NewUserServer(handler *handler.Handlers) IUser {
	return &resource{
		handler: handler,
	}
}

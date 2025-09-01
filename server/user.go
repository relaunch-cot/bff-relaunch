package server

import (
	"net/http"
	"strconv"

	"github.com/relaunch-cot/bff-relaunch/handler"
	models "github.com/relaunch-cot/bff-relaunch/models/user"
	"github.com/relaunch-cot/bff-relaunch/resource/transformer"

	params "github.com/relaunch-cot/bff-relaunch/params/user"

	"github.com/gin-gonic/gin"
)

type IUser interface {
	CreateUser(c *gin.Context)
	LoginUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	UpdateUserPassword(c *gin.Context)
	DeleteUser(c *gin.Context)
}

func (r *resource) CreateUser(c *gin.Context) {
	in := new(params.CreateUserPOST)
	err := c.Bind(in)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error getting query params"})
		return
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
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user registered successfully"})
}

func (r *resource) LoginUser(c *gin.Context) {
	in := new(params.LoginUserGET)
	err := c.Bind(in)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error getting query params"})
		return
	}

	user := &models.User{
		Email:    in.Email,
		Password: in.Password,
	}

	loginUserReq, err := transformer.LoginUserToProto(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error transforming params to proto"})
		return
	}

	ctx := c.Request.Context()

	loginUserResponse, err := r.handler.User.LoginUser(&ctx, loginUserReq)
	if err != nil {
		c.JSON(404, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, loginUserResponse)
}

func (r *resource) UpdateUser(c *gin.Context) {
	userIdStr := c.Param("id")
	if userIdStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user id is required"})
		return
	}

	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid user id"})
		return
	}

	in := new(params.UpdateUserPUT)
	err = c.Bind(in)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error getting query params"})
		return
	}

	user := &models.User{
		UserId:   userId,
		Name:     in.Name,
		Email:    in.Email,
		Password: in.Password,
	}

	updateUserReq, err := transformer.UpdateUserToProto(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error transforming params to proto"})
		return
	}

	ctx := c.Request.Context()
	err = r.handler.User.UpdateUser(&ctx, updateUserReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user updated successfully"})
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

func (r *resource) DeleteUser(c *gin.Context) {
	in := new(params.DeleteUserDELETE)
	err := c.Bind(in)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error getting query params"})
		return
	}

	deleteUserReq, err := transformer.DeleteUserToProto(in.Email, in.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error transforming params to proto"})
		return
	}

	ctx := c.Request.Context()
	err = r.handler.User.DeleteUser(&ctx, deleteUserReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}

func NewUserServer(handler *handler.Handlers) IUser {
	return &resource{
		handler: handler,
	}
}

package server

import (
	"github.com/relaunch-cot/bff/handler"
	model "github.com/relaunch-cot/bff/model/user"
	"github.com/relaunch-cot/bff/resource/transformer"

	"github.com/gin-gonic/gin"
	"strings"
)

type IUser interface {
	CreateUser(c *gin.Context)
}

func (r *resource) CreateUser(c *gin.Context) {
	name := c.Param("name")
	if strings.TrimSpace(name) == "" {
		c.JSON(404, gin.H{"message": "the param name can not be empty"})
		return
	}

	email := c.Param("email")
	if strings.TrimSpace(email) == "" {
		c.JSON(404, gin.H{"message": "the param email can not be empty"})
		return
	}

	password := c.Param("password")
	if strings.TrimSpace(password) == "" {
		c.JSON(404, gin.H{"message": "the param password can not be empty"})
		return
	}

	user := &model.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	in, err := transformer.CreateUserToProto(user)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
	}

	ctx := c.Request.Context()

	err = r.handler.User.CreateUser(&ctx, in)
	if err != nil {
		c.JSON(404, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "user registered successfully"})
}

func NewUserServer(handler *handler.Handlers) IUser {
	return &resource{
		handler: handler,
	}
}

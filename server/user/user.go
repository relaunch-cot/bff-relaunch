package server

import "github.com/gin-gonic/gin"

type IUser interface {
	createUser(c *gin.Context)
}

func createUser(c *gin.Context) {

}

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
	GenerateReportPDF(c *gin.Context)
	SendPasswordRecoveryEmail(c *gin.Context)
	CreateNewChat(c *gin.Context)
}

func (r *resource) CreateUser(c *gin.Context) {
	in := new(params.CreateUserPOST)
	err := c.BindQuery(in)
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

	c.JSON(http.StatusCreated, gin.H{"message": "user registered successfully"})
}

func (r *resource) LoginUser(c *gin.Context) {
	in := new(params.LoginUserGET)
	err := c.BindQuery(in)
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
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.Header("Authorization", loginUserResponse.Token)
	c.Header("Access-Control-Expose-Headers", "Authorization")
	c.JSON(http.StatusOK, gin.H{"message": "user logged in successfully"})
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
		c.JSON(http.StatusBadRequest, gin.H{"message": "error getting body of the request"})
		return
	}

	updateUserPasswordReq, err := transformer.UpdateUserPasswordToProto(in.UserId, in.NewPassword)
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
	err := c.BindQuery(in)
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

func (r *resource) GenerateReportPDF(c *gin.Context) {
	in := new(params.GenerateReportPOST)
	err := c.ShouldBindJSON(in)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error parsing JSON: " + err.Error()})
		return
	}

	if in.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "title is required"})
		return
	}

	if len(in.Headers) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "headers are required"})
		return
	}

	reportData := &models.ReportData{
		Title:    in.Title,
		Subtitle: in.Subtitle,
		Headers:  in.Headers,
		Rows:     in.Rows,
		Footer:   in.Footer,
	}

	generateReportReq, err := transformer.ReportDataToProto(reportData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error transforming params to proto"})
		return
	}

	ctx := c.Request.Context()
	response, err := r.handler.User.GenerateReportPDF(&ctx, generateReportReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename=relatorio.pdf")
	c.Header("Content-Length", strconv.Itoa(len(response.PdfData)))

	c.Data(http.StatusCreated, "application/pdf", response.PdfData)
}

func (r *resource) SendPasswordRecoveryEmail(c *gin.Context) {
	in := new(params.SendPasswordRecoveryEmailPOST)
	err := c.Bind(in)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error getting query params"})
		return
	}

	sendPasswordRecoveryEmailReq, err := transformer.SendPasswordRecoveryEmailToProto(in.Email, in.RecoveryLink)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error transforming params to proto"})
		return
	}

	ctx := c.Request.Context()
	err = r.handler.User.SendPasswordRecoveryEmail(&ctx, sendPasswordRecoveryEmailReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "password recovery email sent successfully"})
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
	err = r.handler.User.CreateNewChat(&ctx, createNewChatRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "chat created successfully"})
}

func NewUserServer(handler *handler.Handlers) IUser {
	return &resource{
		handler: handler,
	}
}

package controller

import (
	"microauth/src/domain"
	"microauth/src/service"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var req *domain.SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		badRequestResponse(c)
	}
	service.NewUserService().SignUp(req)
	successResponse(c)
}

func ConfirmSignUp(c *gin.Context) {
	var req *domain.ConfirmSignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		badRequestResponse(c)
	}
	service.NewUserService().ConfirmSignUp(req)
	successResponse(c)
}

func ForgotPassword(c *gin.Context) {
	var req *domain.ForgotPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		badRequestResponse(c)
	}
	service.NewUserService().ForgotPassword(req)
	successResponse(c)
}

func ConfirmForgotPassword(c *gin.Context) {
	var req *domain.ConfirmForgotPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		badRequestResponse(c)
	}
	service.NewUserService().ConfirmForgotPassword(req)
	successResponse(c)
}

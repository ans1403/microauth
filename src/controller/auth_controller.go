package controller

import (
	"microauth/src/domain"
	"microauth/src/service"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (ctrl *AuthController) SignUp(c *gin.Context) {
	var req *domain.SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		badRequestResponse(c)
	}
	service.NewAuthService().SignUp(req)
	successResponse(c)
}

func (ctrl *AuthController) ConfirmSignUp(c *gin.Context) {
	var req *domain.ConfirmSignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		badRequestResponse(c)
	}
	service.NewAuthService().ConfirmSignUp(req)
	successResponse(c)
}

func (ctrl *AuthController) ForgotPassword(c *gin.Context) {
	var req *domain.ForgotPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		badRequestResponse(c)
	}
	service.NewAuthService().ForgotPassword(req)
	successResponse(c)
}

func (ctrl *AuthController) ConfirmForgotPassword(c *gin.Context) {
	var req *domain.ConfirmForgotPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		badRequestResponse(c)
	}
	service.NewAuthService().ConfirmForgotPassword(req)
	successResponse(c)
}
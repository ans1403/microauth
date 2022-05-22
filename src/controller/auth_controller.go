package controller

import (
	"microauth/src/domain"
	"microauth/src/service"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{
		service.NewAuthService(),
	}
}

func (ctrl *AuthController) SignUp(c *gin.Context) {
	var req *domain.SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		badRequestResponse(c)
	}
	ctrl.authService.SignUp(req)
	successResponse(c)
}

func (ctrl *AuthController) ConfirmSignUp(c *gin.Context) {
	var req *domain.ConfirmSignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		badRequestResponse(c)
	}
	ctrl.authService.ConfirmSignUp(req)
	successResponse(c)
}

func (ctrl *AuthController) ForgotPassword(c *gin.Context) {
	var req *domain.ForgotPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		badRequestResponse(c)
	}
	ctrl.authService.ForgotPassword(req)
	successResponse(c)
}

func (ctrl *AuthController) ConfirmForgotPassword(c *gin.Context) {
	var req *domain.ConfirmForgotPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		badRequestResponse(c)
	}
	ctrl.authService.ConfirmForgotPassword(req)
	successResponse(c)
}

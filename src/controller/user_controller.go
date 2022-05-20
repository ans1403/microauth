package controller

import (
	"microauth/src/domain"
	"microauth/src/service"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var req domain.SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		badRequestResponse(c)
	}
	s := service.UserService{}
	s.SignUp(req.Username, req.Password, req.Email)
	successResponse(c)
}

func ComfirmSignUp(c *gin.Context) {
	var req domain.ConfirmSignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		badRequestResponse(c)
	}
	s := service.UserService{}
	s.ConfirmSignUp(req.Username, req.ConfirmationCode)
	successResponse(c)
}

package controller

import (
	"microauth/src/domain"
	"microauth/src/logging"
	"microauth/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	SignUp(c *gin.Context)
	ConfirmSignUp(c *gin.Context)
	ForgotPassword(c *gin.Context)
	ConfirmForgotPassword(c *gin.Context)
	SignIn(c *gin.Context)
}

func NewAuthController() AuthController {
	return &authController{
		authService: service.NewAuthService(),
		logger:      logging.NewLogger(),
	}
}

type authController struct {
	authService service.AuthService
	logger      logging.Logger
}

func (ctrl *authController) SignUp(c *gin.Context) {
	var req *domain.SignUpRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		ctrl.logger.Info(err.Error())
		responseWithMessage(c, http.StatusBadRequest)
	}

	if err := ctrl.authService.SignUp(req); err != nil {
		ctrl.logger.Error(err.Error())
		responseWithMessage(c, http.StatusInternalServerError)
	}

	responseWithMessage(c, http.StatusOK)
}

func (ctrl *authController) ConfirmSignUp(c *gin.Context) {
	req := &domain.ConfirmSignUpRequest{
		Username:         c.Query("username"),
		ConfirmationCode: c.Query("confirmationCode"),
	}

	if err := ctrl.authService.ConfirmSignUp(req); err != nil {
		ctrl.logger.Error(err.Error())
		responseWithMessage(c, http.StatusInternalServerError)
	}

	responseWithMessage(c, http.StatusOK)
}

func (ctrl *authController) ForgotPassword(c *gin.Context) {
	var req *domain.ForgotPasswordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		ctrl.logger.Info(err.Error())
		responseWithMessage(c, http.StatusBadRequest)
	}

	if err := ctrl.authService.ForgotPassword(req); err != nil {
		ctrl.logger.Error(err.Error())
		responseWithMessage(c, http.StatusInternalServerError)
	}

	responseWithMessage(c, http.StatusOK)
}

func (ctrl *authController) ConfirmForgotPassword(c *gin.Context) {
	var req *domain.ConfirmForgotPasswordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		ctrl.logger.Info(err.Error())
		responseWithMessage(c, http.StatusBadRequest)
	}

	if err := ctrl.authService.ConfirmForgotPassword(req); err != nil {
		ctrl.logger.Error(err.Error())
		responseWithMessage(c, http.StatusInternalServerError)
	}

	responseWithMessage(c, http.StatusOK)
}

func (ctrl *authController) SignIn(c *gin.Context) {
	var req *domain.SignInRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		ctrl.logger.Info(err.Error())
		responseWithMessage(c, http.StatusBadRequest)
	}

	res, err := ctrl.authService.SignIn(req)
	if err != nil {
		ctrl.logger.Error(err.Error())
		responseWithMessage(c, http.StatusInternalServerError)
	}

	session := getDefaultSession(c)
	session.Set("cognitoAccessToken", res.AuthenticationResult.AccessToken)
	session.Set("cognitoIdToken", res.AuthenticationResult.IdToken)
	session.Set("cognitoRefreshToken", res.AuthenticationResult.RefreshToken)
	if err := session.Save(); err != nil {
		ctrl.logger.Error(err.Error())
	}

	// 普通は値を返すことはないけど勉強目的で階層構造の値を返してみる。
	responseWithMessageAndResults(c, http.StatusOK, &domain.CognitoTokens{
		AccessToken:  *res.AuthenticationResult.AccessToken,
		IdToken:      *res.AuthenticationResult.IdToken,
		RefreshToken: *res.AuthenticationResult.RefreshToken,
	})
}

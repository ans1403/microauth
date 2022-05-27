package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FrontController interface {
	SignUp(c *gin.Context)
	CompleteSignUp(c *gin.Context)
	SignIn(c *gin.Context)
}

func NewFrontController() FrontController {
	return &frontController{}
}

type frontController struct{}

func (ctrl *frontController) SignUp(c *gin.Context) {
	c.HTML(http.StatusOK, "signUp.html", nil)
}

func (ctrl *frontController) CompleteSignUp(c *gin.Context) {
	c.HTML(http.StatusOK, "completeSignUp.html", gin.H{})
}

func (ctrl *frontController) SignIn(c *gin.Context) {
	c.HTML(http.StatusOK, "signIn.html", gin.H{})
}

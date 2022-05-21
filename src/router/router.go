package router

import (
	"microauth/src/controller"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.POST("/signUp", controller.SignUp)
		v1.POST("/confirmSignUp", controller.ConfirmSignUp)
		v1.POST("/forgotPassword", controller.ForgotPassword)
		v1.POST("/confirmForgotPassword", controller.ConfirmForgotPassword)
	}
	return router
}

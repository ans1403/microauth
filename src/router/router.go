package router

import (
	"microauth/src/controller"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	authController := controller.NewAuthController()

	v1 := router.Group("api/v1")
	{
		v1.POST("/signUp", authController.SignUp)
		v1.POST("/confirmSignUp", authController.ConfirmSignUp)
		v1.POST("/forgotPassword", authController.ForgotPassword)
		v1.POST("/confirmForgotPassword", authController.ConfirmForgotPassword)
	}

	return router
}

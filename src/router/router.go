package router

import (
	"microauth/src/constants"
	"microauth/src/controller"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	app := constants.NewApp()

	store, err := redis.NewStore(10, "tcp", app.RedisHost+":"+app.RedisPort, "", []byte(app.SecretKey))
	if err != nil {
		panic(err.Error())
	}

	router := gin.Default()
	router.LoadHTMLGlob("src/templates/*.html")
	router.Use(sessions.Sessions("SESSION", store))

	front := router.Group("/")
	{
		frontController := controller.NewFrontController()
		front.GET("/signUp", frontController.SignUp)
		front.GET("/completeSignUp", frontController.CompleteSignUp)
		front.GET("/signIn", frontController.SignIn)
	}

	v1 := router.Group("api/v1")
	{
		authController := controller.NewAuthController()
		v1.POST("/signUp", authController.SignUp)
		v1.GET("/confirmSignUp", authController.ConfirmSignUp)
		v1.POST("/forgotPassword", authController.ForgotPassword)
		v1.POST("/confirmForgotPassword", authController.ConfirmForgotPassword)
		v1.POST("/signIn", authController.SignIn)
	}

	return router
}

package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type jsonResponse struct {
	Message string `json:"message"`
}

func successResponse(c *gin.Context) {
	c.JSON(http.StatusOK, jsonResponse{
		Message: "success",
	})
}

func badRequestResponse(c *gin.Context) {
	c.JSON(http.StatusBadRequest, jsonResponse{
		Message: "Bad Request",
	})
}

func getDefaultSession(c *gin.Context) sessions.Session {
	session := sessions.Default(c)
	session.Options(sessions.Options{
		MaxAge:   604800,
		Secure:   true,
		HttpOnly: true,
	})
	return session
}

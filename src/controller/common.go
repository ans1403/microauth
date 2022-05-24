package controller

import (
	"microauth/src/domain"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func responseWithMessage(c *gin.Context, httpStatus int) {
	c.JSON(httpStatus, domain.ResponseWithMessage{
		Message: http.StatusText(httpStatus),
	})
}

func responseWithMessageAndResults(c *gin.Context, httpStatus int, results interface{}) {
	c.JSON(httpStatus, domain.ResponseWithMessageAndResults{
		Message: http.StatusText(httpStatus),
		Results: results,
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

package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Login(c *gin.Context)
	AuthorizeUser(c *gin.Context)
}

func SetupAuthRouter(authHandler AuthHandler) *gin.Engine {
	router := gin.New()

	router.POST("/login", authHandler.Login)

	router.GET("/user/:user_id", authHandler.AuthorizeUser)
	router.PUT("/user/:user_id", authHandler.AuthorizeUser)
	router.POST("/user", func(c *gin.Context) { c.Status(http.StatusOK) })

	return router
}

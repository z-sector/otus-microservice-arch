package server

import (
	"github.com/gin-gonic/gin"
)

type Middlewares struct{}

type HealthCheckHandler interface {
	Handle(c *gin.Context)
}

type UserHandler interface {
	Create(c *gin.Context)
	Read(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type Handlers struct {
	HealthCheck HealthCheckHandler
	User        UserHandler
}

func SetupRouter(mws *Middlewares, handlers *Handlers) *gin.Engine {
	router := gin.New()

	router.GET("/health", handlers.HealthCheck.Handle)

	// user
	router.POST("/user", handlers.User.Create)
	router.GET("/user/:user_id", handlers.User.Read)
	router.PUT("/user/:user_id", handlers.User.Update)
	router.DELETE("/user/:user_id", handlers.User.Delete)

	return router
}

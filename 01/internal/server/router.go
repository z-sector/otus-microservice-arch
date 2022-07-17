package server

import (
	"github.com/gin-gonic/gin"
)

type Middlewares struct{}

type HealthCheckHandler interface {
	Handle(c *gin.Context)
}

type Handlers struct {
	HealthCheck HealthCheckHandler
}

func SetupRouter(mws *Middlewares, handlers *Handlers) *gin.Engine {
	router := gin.New()

	router.GET("/health", handlers.HealthCheck.Handle)

	return router
}

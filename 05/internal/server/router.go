package server

import (
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
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

	ginMetrics := ginmetrics.GetMonitor()
	ginMetrics.UseWithoutExposingEndpoint(router)

	router.GET("/health", handlers.HealthCheck.Handle)

	// user
	router.POST("/user", handlers.User.Create)
	router.GET("/user/:user_id", handlers.User.Read)
	router.PUT("/user/:user_id", handlers.User.Update)

	return router
}

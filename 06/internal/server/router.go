package server

import (
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
)

type Middlewares struct{}

type HealthCheckHandler interface {
	Handle(c *gin.Context)
}

type OrderHandler interface {
	Create(c *gin.Context)
}

type Handlers struct {
	HealthCheck HealthCheckHandler
	Order       OrderHandler
}

func SetupRouter(mws *Middlewares, handlers *Handlers) *gin.Engine {
	router := gin.New()

	ginMetrics := ginmetrics.GetMonitor()
	ginMetrics.UseWithoutExposingEndpoint(router)

	router.GET("/health", handlers.HealthCheck.Handle)

	router.POST("/orders", handlers.Order.Create)

	return router
}

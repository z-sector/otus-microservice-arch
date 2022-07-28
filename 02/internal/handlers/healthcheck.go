package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheckHandler struct{}

func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

func (h *HealthCheckHandler) Handle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

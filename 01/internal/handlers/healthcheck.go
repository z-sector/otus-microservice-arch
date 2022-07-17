package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheck struct{}

func NewHealthCheck() *HealthCheck {
	return &HealthCheck{}
}

func (h *HealthCheck) Handle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

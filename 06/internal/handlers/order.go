package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/wuzyk/otus-microservice-arch/06/internal/domain"
)

type CreateOrderRequest struct {
	ProductID       int    `json:"product_id"`
	Quantity        int    `json:"quantity"`
	ShippingAddress string `json:"shipping_address"`
	RequestID       string `json:"request_id"`
}

type OrderResponse struct {
	ID              int    `json:"id"`
	ProductID       int    `json:"product_id"`
	Quantity        int    `json:"quantity"`
	ShippingAddress string `json:"shipping_address"`
}

type OrderRepo interface {
	New(ctx context.Context, user *domain.Order) error
}

type OrderHandler struct {
	orderRepo OrderRepo
}

func NewOrderHandler(orderRepo OrderRepo) *OrderHandler {
	return &OrderHandler{orderRepo}
}

func (h *OrderHandler) Create(c *gin.Context) {
	var request CreateOrderRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	order := domain.Order{
		ProductID:       request.ProductID,
		Quantity:        request.Quantity,
		ShippingAddress: request.ShippingAddress,
		RequestID:       request.RequestID,
	}

	err := h.orderRepo.New(c, &order)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &OrderResponse{
		ID:              order.ID,
		ProductID:       order.ProductID,
		Quantity:        order.Quantity,
		ShippingAddress: order.ShippingAddress,
	})
}

package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/wuzyk/otus-microservice-arch/02/internal/domain"
)

type CreateUserRequest struct {
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type UpdateUserRequest struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type UserResponse struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type UserRepo interface {
	New(ctx context.Context, user *domain.User) error
	Get(ctx context.Context, userID int) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) (*domain.User, error)
	Delete(ctx context.Context, userID int) error
}

type UserHandler struct {
	userRepo UserRepo
}

func NewUserHandler(userRepo UserRepo) *UserHandler {
	return &UserHandler{userRepo}
}

func (h *UserHandler) Create(c *gin.Context) {
	var request CreateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user := domain.User{
		Username:  request.Username,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Phone:     request.Phone,
	}

	err := h.userRepo.New(c, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
	})
}

func (h *UserHandler) Read(c *gin.Context) {
	userIDParam := c.Param("user_id")

	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, err := h.userRepo.Get(c, userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
	})
}

func (h *UserHandler) Update(c *gin.Context) {
	userIDParam := c.Param("user_id")

	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var request UpdateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user := domain.User{
		ID:        userID,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Phone:     request.Phone,
	}

	updatedUser, err := h.userRepo.Update(c, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &UserResponse{
		ID:        updatedUser.ID,
		Username:  updatedUser.Username,
		FirstName: updatedUser.FirstName,
		LastName:  updatedUser.LastName,
		Email:     updatedUser.Email,
		Phone:     updatedUser.Phone,
	})
}

func (h *UserHandler) Delete(c *gin.Context) {
	userIDParam := c.Param("user_id")

	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = h.userRepo.Delete(c, userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Status(http.StatusOK)
}

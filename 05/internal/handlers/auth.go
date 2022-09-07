package handlers

import (
	"crypto/rsa"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/wuzyk/otus-microservice-arch/05/internal/domain"
)

const (
	tokenExpireTime = time.Hour
	tokenCookieName = "token"
	tokenUserIDKey  = "user_id"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthHandler struct {
	userRepo domain.UserRepo
	key      *rsa.PrivateKey
}

func NewAuthHandler(userRepo domain.UserRepo, key *rsa.PrivateKey) *AuthHandler {
	return &AuthHandler{userRepo, key}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, err := h.userRepo.SearchByEmail(c, request.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	if user == nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := h.generateToken(user.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.SetCookie(tokenCookieName, token, int(tokenExpireTime/time.Second), "/", "", false, true)
	c.Status(http.StatusOK)
}

func (h *AuthHandler) AuthorizeUser(c *gin.Context) {
	token, err := c.Cookie(tokenCookieName)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	tokenUserID, err := h.getUserIDFromToken(token)
	if err != nil {
		log.Print(err)
		c.Status(http.StatusUnauthorized)
		return
	}

	userIDParam := c.Param("user_id")
	requestedUserID, err := strconv.Atoi(userIDParam)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	if tokenUserID != requestedUserID {
		c.Status(http.StatusUnauthorized)
		return
	}

	c.Status(http.StatusOK)
}

func (h *AuthHandler) generateToken(userID int) (string, error) {
	token := jwt.New()

	_ = token.Set(jwt.ExpirationKey, time.Now().Add(tokenExpireTime).UTC())
	_ = token.Set(tokenUserIDKey, userID)

	signedBytes, err := jwt.Sign(token, jwa.RS256, h.key)

	return string(signedBytes), err
}

func (h *AuthHandler) getUserIDFromToken(token string) (int, error) {
	parsed, err := jwt.Parse([]byte(token),
		jwt.WithValidate(true),
		jwt.WithVerify(jwa.RS256, &h.key.PublicKey),
	)
	if err != nil {
		return 0, err
	}

	userIDClaim, ok := parsed.Get(tokenUserIDKey)
	if !ok {
		return 0, errors.New("failed to get user_id from token")
	}

	userID, ok := userIDClaim.(float64)
	if !ok {
		return 0, errors.New("invalid user_id in token")
	}

	return int(userID), nil
}

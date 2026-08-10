package handlers

import (
	"net/http"
	"task-mgmt/services"
	"task-mgmt/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type AuthHandler struct {
	db          *gorm.DB
	authService services.AuthService
}

type AuthRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewAuthHandler(db *gorm.DB, authService services.AuthService) *AuthHandler {
	return &AuthHandler{db: db, authService: authService}
}

// Token handles user login and token generation
func (h *AuthHandler) Token(c *gin.Context) {
	var auth AuthRequest

	if err := c.ShouldBindJSON(&auth); err != nil {
		c.Error(err)
		return
	}

	user, err := h.authService.LoginUser(h.db, auth.Username, auth.Password)
	if err != nil {
		c.Error(err)
		return
	}

	accessToken, refreshToken, err := h.authService.GenerateToken(h.db, user.ID, uuid.Nil)
	if err != nil {
		c.Error(err)
		return
	}

	// Implement the login and token generation logic here
	c.JSON(http.StatusOK, utils.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    utils.GetEnvAsDuration("JWT_EXPIRATION", time.Hour),
	})
}

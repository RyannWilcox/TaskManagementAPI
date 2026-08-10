package handlers

import (
	"net/http"
	"time"

	"task-mgmt/services"
	"task-mgmt/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RefreshHandler struct {
	db          *gorm.DB
	authService services.AuthService
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func NewRefreshHandler(db *gorm.DB, authService services.AuthService) *RefreshHandler {
	return &RefreshHandler{db: db, authService: authService}
}

func (h *RefreshHandler) Refresh(c *gin.Context) {

	var req RefreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	// Validate the refresh token and get the associated user ID
	oldRefreshToken, err := h.authService.ValidateRefreshToken(h.db, req.RefreshToken)
	if err != nil {
		c.Error(err)
		return
	}

	// Generate new access and refresh tokens for the user
	accessToken, refreshToken, err := h.authService.GenerateToken(h.db, oldRefreshToken.UserId, oldRefreshToken.ID)
	if err != nil {
		c.Error(err)
		return
	}

	// Implement the refresh token logic here
	c.JSON(http.StatusOK, utils.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    utils.GetEnvAsDuration("JWT_EXPIRATION", time.Hour),
	})
}

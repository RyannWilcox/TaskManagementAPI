package handlers

import (
	"errors"
	"net/http"

	"task-mgmt/models"
	"task-mgmt/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var ErrDuplicateUsername = errors.New("username already exists")

type RegisterHandler struct {
	db              *gorm.DB
	registerService services.RegisterService
}

func NewRegisterHandler(db *gorm.DB, registerService services.RegisterService) *RegisterHandler {
	return &RegisterHandler{db: db, registerService: registerService}
}

func (h *RegisterHandler) Registration(c *gin.Context) {
	// Implement the registration logic here
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	err := h.registerService.RegisterUser(h.db, user)
	if err != nil {
		if errors.Is(err, ErrDuplicateUsername) {
			c.JSON(http.StatusConflict, gin.H{"error": "username already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to register user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}

package handlers

import (
	"net/http"

	"task-mgmt/models"
	"task-mgmt/services"
	"task-mgmt/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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
		c.Error(err)
		return
	}

	if err := h.registerService.RegisterUser(h.db, user); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, utils.MessageResponse{
		Message: "user created successfully",
	})
}

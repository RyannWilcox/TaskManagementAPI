package handlers

import (
	"errors"
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
		c.JSON(http.StatusBadRequest, utils.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "badly formed json provided.",
		})
		return
	}

	err := h.registerService.RegisterUser(h.db, user)
	if err != nil {
		if errors.Is(err, utils.ErrDuplicateUsername) {
			c.JSON(http.StatusConflict, utils.HTTPError{
				Code:    http.StatusConflict,
				Message: "username already exists",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, utils.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "failed to register user",
		})
		return
	}
	c.JSON(http.StatusCreated, utils.MessageResponse{
		Message: "user created successfully",
	})
}

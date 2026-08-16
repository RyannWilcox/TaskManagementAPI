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

// Registration godoc
// @Summary      Register a new user
// @Description  Creates a new user account and assigns the default "user" role
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      models.User  true  "New user's username, email, and password"
// @Success      201      {object}  utils.MessageResponse
// @Failure      400      {object}  utils.HTTPError  "invalid request payload"
// @Failure      409      {object}  utils.HTTPError  "username already exists"
// @Router       /auth/register [post]
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

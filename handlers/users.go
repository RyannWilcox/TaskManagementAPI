package handlers

import (
	"net/http"
	"task-mgmt/services"
	"task-mgmt/utils"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type UserHandler struct {
	db          *gorm.DB
	userService services.UserService
}

func NewUserHandler(db *gorm.DB, userService services.UserService) *UserHandler {
	return &UserHandler{db: db, userService: userService}
}

func (h *UserHandler) GetUserProfile(c *gin.Context) {
	userID := c.MustGet("user_id").(uuid.UUID)
	user, err := h.userService.GetUserProfile(h.db, userID)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetUserProfileByUserId(c *gin.Context) {
	paramID, err := uuid.FromString(c.Param("user_id"))
	if err != nil {
		c.Error(err)
		return
	}

	contextUserID := c.MustGet("userID").(uuid.UUID)
	if contextUserID != paramID {
		c.Error(utils.ErrCannotViewProfile)
		return
	}

	user, err := h.userService.GetUserProfile(h.db, paramID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.userService.GetUsers(h.db)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	userID, err := uuid.FromString(c.Param("user_id"))
	if err != nil {
		c.Error(err)
		return
	}

	if err := h.userService.DeleteUser(h.db, userID); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusNoContent, utils.MessageResponse{
		Message: "User succesfully deleted",
	})
}

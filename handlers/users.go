package handlers

import (
	"net/http"
	"slices"
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

// GetUserProfile godoc
// @Summary      Get the authenticated user's profile
// @Description  Returns the profile of the currently authenticated user
// @Tags         users
// @Security     BearerAuth
// @Produce      json
// @Success      200  {object}  models.User
// @Failure      401  {object}  utils.HTTPError  "invalid or missing token"
// @Router       /users/profile [get]
func (h *UserHandler) GetUserProfile(c *gin.Context) {
	userID := c.MustGet("userID").(uuid.UUID)
	user, err := h.userService.GetUserProfile(h.db, userID)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, user)
}

// GetUserProfileByUserId godoc
// @Summary      Get a user's profile by ID
// @Description  Returns the profile of the specified user. Requires the "user:view" permission.
// @Tags         users
// @Security     BearerAuth
// @Produce      json
// @Param        user_id  path      string  true  "User ID"
// @Success      200      {object}  models.User
// @Failure      400      {object}  utils.HTTPError  "invalid user id"
// @Failure      403      {object}  utils.HTTPError  "insufficient permissions"
// @Failure      404      {object}  utils.HTTPError  "record not found"
// @Router       /users/profile/{user_id} [get]
func (h *UserHandler) GetUserProfileByUserId(c *gin.Context) {
	paramID, err := uuid.FromString(c.Param("user_id"))
	if err != nil {
		c.Error(utils.ErrInvalidUUID)
		return
	}

	// get context user id and roles
	contextUserID := c.MustGet("userID").(uuid.UUID)
	contextUserRoles := c.MustGet("userRoles").([]string)

	// an admin can view any profile..
	isAdmin := slices.Contains(contextUserRoles, "admin")

	if contextUserID != paramID && !isAdmin {
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

// GetUsers godoc
// @Summary      List all users
// @Description  Returns every user in the system. Requires the "admin" role.
// @Tags         users
// @Security     BearerAuth
// @Produce      json
// @Success      200  {array}   models.User
// @Failure      403  {object}  utils.HTTPError  "insufficient permissions"
// @Router       /users [get]
func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.userService.GetUsers(h.db)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, users)
}

// DeleteUser godoc
// @Summary      Delete a user
// @Description  Deletes the specified user account. Requires the "admin" role.
// @Tags         users
// @Security     BearerAuth
// @Produce      json
// @Param        user_id  path      string  true  "User ID"
// @Success      204      {object}  utils.MessageResponse
// @Failure      400      {object}  utils.HTTPError  "invalid user id"
// @Failure      403      {object}  utils.HTTPError  "insufficient permissions"
// @Failure      404      {object}  utils.HTTPError  "record not found"
// @Router       /users/{user_id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	userID, err := uuid.FromString(c.Param("user_id"))
	if err != nil {
		c.Error(utils.ErrInvalidUUID)
		return
	}

	if err := h.userService.DeleteUser(h.db, userID); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

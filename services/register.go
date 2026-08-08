package services

import (
	"task-mgmt/models"
	"task-mgmt/utils"

	"gorm.io/gorm"
)

type RegisterService interface {
	RegisterUser(db *gorm.DB, user models.User) error
}

type RegisterServiceImpl struct{}

func NewRegisterService() *RegisterServiceImpl {
	return &RegisterServiceImpl{}
}

// RegisterUser registers a new user in the database
func (s *RegisterServiceImpl) RegisterUser(db *gorm.DB, user models.User) error {
	// Check if the username already exists
	var existingUser models.User
	if err := db.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		return utils.ErrDuplicateUsername
	}

	hash, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	user.Password = hash

	// Create the new user
	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

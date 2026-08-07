package services

import (
	"errors"
	"task-mgmt/models"

	"gorm.io/gorm"
)

type RegisterService interface {
	RegisterUser(db *gorm.DB, user models.User) error
}

type RegisterServiceImpl struct{}

func NewRegisterService() *RegisterServiceImpl {
	return &RegisterServiceImpl{}
}

// TODO: Need to implement password hashing and validation before storing the user in the database.

func (s *RegisterServiceImpl) RegisterUser(db *gorm.DB, user models.User) error {
	// Check if the username already exists
	var existingUser models.User
	if err := db.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		return errors.New("username already exists")
	}

	// Create the new user
	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

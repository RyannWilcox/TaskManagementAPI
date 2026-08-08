package services

import (
	"task-mgmt/models"
	"task-mgmt/utils"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type AuthService interface {
	LoginUser(db *gorm.DB, username, password string) (*models.User, error)
	GenerateToken(db *gorm.DB, userID uuid.UUID) (string, string, error)
}

type AuthServiceImpl struct {
}

func NewAuthService() *AuthServiceImpl {
	return &AuthServiceImpl{}
}

func (s *AuthServiceImpl) LoginUser(db *gorm.DB, username, password string) (*models.User, error) {
	var user models.User
	// Check if the user exists
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, utils.ErrInvalidCredentials
	}

	// verify the user password and the provided password match.
	if !utils.VerifyPassword(user.Password, password) {
		return nil, utils.ErrInvalidCredentials
	}
	return &user, nil
}

func (s *AuthServiceImpl) GenerateToken(db *gorm.DB, userID uuid.UUID) (string, string, error) {
	return "", "", nil
}

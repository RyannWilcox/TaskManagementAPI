package services

import (
	"task-mgmt/models"

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

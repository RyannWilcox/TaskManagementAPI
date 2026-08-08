package services

import (
	"task-mgmt/models"
	"task-mgmt/utils"
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type AuthService interface {
	LoginUser(db *gorm.DB, username, password string) (*models.User, error)
	GenerateToken(db *gorm.DB, userID uuid.UUID) (string, string, int64, error)
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

// GenerateToken generates an access token and a refresh token for the given user ID.
func (s *AuthServiceImpl) GenerateToken(db *gorm.DB, userID uuid.UUID) (string, string, int64, error) {
	accessToken, expirationTime, err := utils.GenerateAccessToken(userID)
	if err != nil {
		return "", "", 0, err
	}
	refreshToken, err := uuid.NewV4()
	if err != nil {
		return "", "", 0, err
	}

	newToken := models.Token{
		UserId:       userID,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(expirationTime),
	}

	// Since this is a newly generated token,
	// we should delete the users currently stored one
	if err := db.Where("user_id = ?", userID).Delete(&models.Token{}).Error; err != nil {
		return "", "", 0, err
	}

	// Create the new token record in the db.
	if err := db.Create(&newToken).Error; err != nil {
		return "", "", 0, err
	}

	return accessToken, refreshToken.String(), int64(expirationTime.Seconds()), nil
}

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
	GenerateToken(db *gorm.DB, userID uuid.UUID, oldTokenID uuid.UUID) (string, string, error)
	ValidateRefreshToken(db *gorm.DB, refreshToken string) (*models.Token, error)
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
func (s *AuthServiceImpl) GenerateToken(db *gorm.DB, userID uuid.UUID, oldTokenID uuid.UUID) (string, string, error) {
	roles, permissions, err := s.getRolesAndPermissions(db, userID)
	if err != nil {
		return "", "", err
	}

	accessToken, expirationTime, err := utils.GenerateAccessToken(userID, roles, permissions)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := uuid.NewV4()
	if err != nil {
		return "", "", err
	}

	newToken := models.Token{
		UserId:       userID,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(expirationTime),
	}

	if oldTokenID != uuid.Nil {
		// we should delete the users currently stored one
		if err := db.Where("id = ?", oldTokenID).Delete(&models.Token{}).Error; err != nil {
			return "", "", err
		}
	}

	// Create the new token record in the db.
	if err := db.Create(&newToken).Error; err != nil {
		return "", "", err
	}

	return accessToken, refreshToken.String(), nil
}

// validates the provided refresh token and returns the associated user ID if valid.
func (s *AuthServiceImpl) ValidateRefreshToken(db *gorm.DB, refreshToken string) (*models.Token, error) {
	var token models.Token
	if err := db.Where("refresh_token = ?", refreshToken).First(&token).Error; err != nil {
		return nil, err
	}

	if time.Now().After(token.ExpiresAt) {
		return nil, utils.ErrRefreshTokenExpired
	}

	return &token, nil
}

func (s *AuthServiceImpl) getRolesAndPermissions(db *gorm.DB, userID uuid.UUID) ([]string, []string, error) {
	var user models.User
	if err := db.Preload("Roles.Permissions").Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, nil, err
	}

	roles := make([]string, 0, len(user.Roles))
	permissionSet := make(map[string]struct{})
	for _, role := range user.Roles {
		roles = append(roles, role.Name)
		for _, permission := range role.Permissions {
			permissionSet[permission.Resource+":"+permission.Action] = struct{}{}
		}
	}

	permissions := make([]string, 0, len(permissionSet))
	for permission := range permissionSet {
		permissions = append(permissions, permission)
	}

	return roles, permissions, nil
}

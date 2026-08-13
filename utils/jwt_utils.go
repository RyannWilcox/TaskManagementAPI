package utils

import (
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v5"
)

type TokenClaims struct {
	UserId      uuid.UUID
	Roles       []string
	Permissions []string
	jwt.RegisteredClaims
}

func GetSecret() []byte {
	return []byte(GetEnv("JWT_SECRET", ""))
}

func GetExpiration() time.Duration {
	return GetEnvAsDuration("JWT_EXPIRATION", time.Hour)
}

func GenerateAccessToken(UserId uuid.UUID, Roles []string, Permissions []string) (string, time.Duration, error) {
	expirationTime := GetExpiration()
	currentTime := time.Now()
	claims := &TokenClaims{
		UserId:      UserId,
		Roles:       Roles,
		Permissions: Permissions,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(currentTime),
			ExpiresAt: jwt.NewNumericDate(currentTime.Add(expirationTime)),
			NotBefore: jwt.NewNumericDate(currentTime),
			Issuer:    "task-manager-API",
			Subject:   "user-auth",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(GetSecret())
	if err != nil {
		return "", 0, err
	}

	return tokenString, expirationTime, nil
}

// Extracts the bearer token from the Authorization header,
// validates its signature and expiry, and returns its claims.
func ParseAndValidateToken(authHeader string) (*TokenClaims, error) {
	if authHeader == "" {
		return nil, ErrInvalidToken
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		return nil, ErrInvalidToken
	}
	tokenString := parts[1]

	claims := &TokenClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, ErrInvalidSigningMethod
			}
			return GetSecret(), nil
		})
	if err != nil || !token.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}

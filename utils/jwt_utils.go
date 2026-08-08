package utils

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v5"
)

type TokenClaims struct {
	UserId uuid.UUID
	jwt.RegisteredClaims
}

func GetSecret() []byte {
	return []byte(GetEnv("JWT_SECRET", ""))
}

func GetExpiration() time.Duration {
	return GetEnvAsDuration("JWT_EXPIRATION", time.Hour)
}

func GenerateAccessToken(UserId uuid.UUID) (string, time.Duration, error) {
	expirationTime := GetExpiration()
	currentTime := time.Now()
	claims := &TokenClaims{
		UserId: UserId,
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

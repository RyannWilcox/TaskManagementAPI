// utils/response.go
package utils

import "time"

// MessageResponse defines the structure for success messages
type MessageResponse struct {
	Message string `json:"message" example:"Media deleted"`
}

// HTTPError defines the structure for error responses
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"Invalid input"`
}

type TokenResponse struct {
	AccessToken  string        `json:"access_token"`
	RefreshToken string        `json:"refresh_token"`
	ExpiresIn    time.Duration `json:"expires_in"`
}

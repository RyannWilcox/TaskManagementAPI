package middleware

import (
	"encoding/json"
	"errors"
	"net/http"
	"task-mgmt/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			var statusCode int
			var message string

			switch {
			case errors.Is(err, utils.ErrDuplicateUsername):
				statusCode = http.StatusConflict
				message = "username already exists"
			case errors.Is(err, utils.ErrInvalidCredentials):
				statusCode = http.StatusUnauthorized
				message = "invalid username or password"
			case errors.Is(err, utils.ErrInvalidToken):
				statusCode = http.StatusUnauthorized
				message = "invalid token"
			case errors.Is(err, utils.ErrInvalidSigningMethod):
				statusCode = http.StatusUnauthorized
				message = "invalid signing method"
			case errors.Is(err, utils.ErrRefreshTokenExpired):
				statusCode = http.StatusUnauthorized
				message = "refresh token has expired"
			case errors.Is(err, gorm.ErrRecordNotFound):
				statusCode = http.StatusNotFound
				message = "record not found"
			case errors.As(err, new(*json.SyntaxError)), errors.As(err, new(*json.UnmarshalTypeError)):
				statusCode = http.StatusBadRequest
				message = "badly formed json provided."
			case errors.As(err, new(validator.ValidationErrors)):
				statusCode = http.StatusBadRequest
				message = "invalid request payload"
			default:
				statusCode = http.StatusInternalServerError
				message = "internal server error"
			}
			c.JSON(statusCode, utils.HTTPError{
				Code:    statusCode,
				Message: message,
			})
		}
	}
}

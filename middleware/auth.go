package middleware

import (
	"net/http"
	"task-mgmt/utils"

	"github.com/gin-gonic/gin"
)

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := utils.ParseAndValidateToken(c.GetHeader("Authorization"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.HTTPError{
				Code:    http.StatusUnauthorized,
				Message: "invalid or missing token",
			})
			return
		}
		c.Set("userID", claims.UserId)
		c.Next()
	}
}

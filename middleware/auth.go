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
			logAccessDenied("Missing or invalid token", map[string]interface{}{
				"error":  err,
				"path":   c.Request.URL.Path,
				"method": c.Request.Method,
			})
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.HTTPError{
				Code:    http.StatusUnauthorized,
				Message: "invalid or missing token",
			})
			return
		}
		// Expose user id, roles, and permissions from the token
		c.Set("userID", claims.UserId)
		c.Set("userRoles", claims.Roles)
		c.Set("userPermissions", claims.Permissions)
		c.Next()
	}
}

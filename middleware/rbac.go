package middleware

import (
	"net/http"
	"slices"
	"task-mgmt/utils"

	"github.com/gin-gonic/gin"
)

func RequireRole(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRoles, exists := c.Get("userRoles")
		if !exists || !isValid(userRoles.([]string), roles) {
			logAccessDenied("Missing required role", map[string]interface{}{
				"method":         c.Request.Method,
				"path":           c.Request.URL.Path,
				"roles":          userRoles.([]string),
				"required_roles": roles,
			})
			c.AbortWithStatusJSON(http.StatusForbidden, utils.HTTPError{
				Code:    http.StatusForbidden,
				Message: "Missing required role",
			})
			return
		}
		c.Next()
	}
}

func RequirePermission(permissions ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userPermissions, exists := c.Get("userPermissions")
		if !exists || !isValid(userPermissions.([]string), permissions) {
			logAccessDenied("Missing required permissions", map[string]interface{}{
				"method":         c.Request.Method,
				"path":           c.Request.URL.Path,
				"permissions":    userPermissions.([]string),
				"required_perms": permissions,
			})
			c.AbortWithStatusJSON(http.StatusForbidden, utils.HTTPError{
				Code:    http.StatusForbidden,
				Message: "Missing required permissions",
			})
			return
		}
		c.Next()
	}
}

// Find if there are any matches between actual and expected roles/permissions.
func isValid(actual []string, expected []string) bool {
	for _, item := range expected {
		if slices.Contains(actual, item) {
			return true
		}
	}
	return false
}

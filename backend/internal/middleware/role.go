package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	RoleSystemAdmin = "SYSTEM_ADMIN"
	RoleSchoolAdmin = "SCHOOL_ADMIN"
	RoleTeacher     = "TEACHER"
)

func RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "Role not found in context",
			})
			c.Abort()
			return
		}

		userRole := role.(string)
		allowed := false
		for _, allowedRole := range allowedRoles {
			if userRole == allowedRole {
				allowed = true
				break
			}
		}

		if !allowed {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "Insufficient permissions",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

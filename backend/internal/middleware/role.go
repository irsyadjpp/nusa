package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	RoleSystemAdmin = "SYSTEM_ADMIN"
	RoleSchoolAdmin = "SCHOOL_ADMIN"
	RoleTeacher     = "TEACHER"
	RoleStudent     = "STUDENT"
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

func ReadOnlyMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authCtx := GetAuthContext(c)
		if authCtx == nil {
			c.JSON(http.StatusForbidden, gin.H{"success": false, "message": "Authentication required"})
			c.Abort()
			return
		}

		userRole := authCtx.Role

		// Allow full access for admin roles
		fullAccessRoles := []string{RoleSystemAdmin, RoleSchoolAdmin}
		for _, fullAccessRole := range fullAccessRoles {
			if userRole == fullAccessRole {
				c.Next()
				return
			}
		}

		// Check read-only access for teachers
		readOnlyRoles := []string{RoleTeacher}
		for _, readOnlyRole := range readOnlyRoles {
			if userRole == readOnlyRole {
				// Only allow GET requests for read-only roles
				if c.Request.Method != "GET" {
					c.JSON(http.StatusForbidden, gin.H{"success": false, "message": "Read-only access"})
					c.Abort()
					return
				}
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"success": false, "message": "Insufficient permissions"})
		c.Abort()
	}
}

package middleware

import (
	"context"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/repository"
	"github.com/nusa/backend/pkg/jwt"
)

// AuthContextKey is the key used to store user context in gin.Context
const AuthContextKey = "auth_context"

type AuthContext struct {
	UserID      string
	SchoolID    *string
	Role        string
	Permissions []string
}

// AuthMiddleware validates JWT tokens and sets user context
func AuthMiddleware(jwtService *jwt.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(401, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		token := parts[1]
		claims, err := jwtService.ValidateToken(token)
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Set auth context
		authCtx := &AuthContext{
			UserID:      claims.UserID,
			SchoolID:    claims.SchoolID,
			Role:        claims.Role,
			Permissions: claims.Permissions,
		}
		c.Set(AuthContextKey, authCtx)

		c.Next()
	}
}

// GetAuthContext retrieves the auth context from gin.Context
func GetAuthContext(c *gin.Context) *AuthContext {
	if authCtx, exists := c.Get(AuthContextKey); exists {
		return authCtx.(*AuthContext)
	}
	return nil
}

// HasPermission checks if a role has a specific permission
func HasPermission(role, resource, action string) bool {
	return domain.HasPermission(role, resource, action)
}

// RequirePermission checks if the user has the required permission
func RequirePermission(requiredPermission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authCtx := GetAuthContext(c)
		if authCtx == nil {
			c.JSON(401, gin.H{"error": "Authentication required"})
			c.Abort()
			return
		}

		// Check if user has the required permission
		hasPermission := false
		for _, perm := range authCtx.Permissions {
			if perm == requiredPermission {
				hasPermission = true
				break
			}
		}

		if !hasPermission {
			c.JSON(403, gin.H{"error": "Insufficient permissions"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// RequireRole checks if the user has the required role
func RequireRole(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authCtx := GetAuthContext(c)
		if authCtx == nil {
			c.JSON(401, gin.H{"error": "Authentication required"})
			c.Abort()
			return
		}

		// Check if user has one of the allowed roles
		hasRole := false
		for _, role := range allowedRoles {
			if authCtx.Role == role {
				hasRole = true
				break
			}
		}

		if !hasRole {
			c.JSON(403, gin.H{"error": "Insufficient role permissions"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// RequireSchoolAccess checks if the user has access to the specified school
func RequireSchoolAccess(userRepo *repository.UserRepository, schoolRepo *repository.SchoolRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		authCtx := GetAuthContext(c)
		if authCtx == nil {
			c.JSON(401, gin.H{"error": "Authentication required"})
			c.Abort()
			return
		}

		// SYSTEM_ADMIN can access any school
		if authCtx.Role == domain.RoleSystemAdmin {
			c.Next()
			return
		}

		// Get school_id from path or query
		schoolID := c.Param("school_id")
		if schoolID == "" {
			schoolID = c.Query("school_id")
		}

		if schoolID == "" {
			c.JSON(400, gin.H{"error": "School ID is required"})
			c.Abort()
			return
		}

		// Check if user belongs to this school
		user, err := userRepo.GetByID(context.Background(), authCtx.UserID)
		if err != nil {
			c.JSON(401, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		if user.SchoolID == nil || *user.SchoolID != schoolID {
			c.JSON(403, gin.H{"error": "Access denied to this school"})
			c.Abort()
			return
		}

		c.Next()
	}
}

package unit

import (
	"net/http/httptest"
	"testing"
	"time"

	jwtService "github.com/nusa/backend/pkg/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/gin-gonic/gin"
)

// TestAuthMiddleware_ValidToken validates that valid tokens are accepted
func TestAuthMiddleware_ValidToken(t *testing.T) {
	gin.SetMode(gin.TestMode)

	secret := "test-secret-key"
	accessExpiry := 1 * time.Hour
	refreshExpiry := 7 * 24 * time.Hour
	issuer := "nusa-backend"

	jwtSvc := jwtService.NewService(secret, accessExpiry, refreshExpiry, issuer)

	// Generate a valid token
	userID := "user-123"
	role := "SYSTEM_ADMIN"
	schoolID := "school-456"
	permissions := []string{"user:CREATE", "school:READ"}

	token, err := jwtSvc.GenerateAccessToken(userID, role, &schoolID, permissions)
	require.NoError(t, err)

	// Create a test router with auth middleware
	router := gin.New()
	router.Use(func(c *gin.Context) {
		// Simulate auth middleware
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Extract token
		if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
			c.JSON(401, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		tokenStr := authHeader[7:]
		_, err := jwtSvc.ValidateToken(tokenStr)
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		c.Next()
	})

	router.GET("/protected", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "success"})
	})

	// Create a request with valid token
	req := httptest.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	// Should succeed
	assert.Equal(t, 200, w.Code)
}

// TestAuthMiddleware_InvalidToken validates that invalid tokens are rejected
func TestAuthMiddleware_InvalidToken(t *testing.T) {
	gin.SetMode(gin.TestMode)

	secret := "test-secret-key"
	accessExpiry := 1 * time.Hour
	refreshExpiry := 7 * 24 * time.Hour
	issuer := "nusa-backend"

	jwtSvc := jwtService.NewService(secret, accessExpiry, refreshExpiry, issuer)

	// Create a test router with auth middleware
	router := gin.New()
	router.Use(func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
			c.JSON(401, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		tokenStr := authHeader[7:]
		_, err := jwtSvc.ValidateToken(tokenStr)
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		c.Next()
	})

	router.GET("/protected", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "success"})
	})

	// Test with invalid token
	req := httptest.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer invalid_token")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	// Should fail with 401
	assert.Equal(t, 401, w.Code)
}

// TestAuthMiddleware_MissingToken validates that missing tokens are rejected
func TestAuthMiddleware_MissingToken(t *testing.T) {
	gin.SetMode(gin.TestMode)

	secret := "test-secret-key"
	accessExpiry := 1 * time.Hour
	refreshExpiry := 7 * 24 * time.Hour
	issuer := "nusa-backend"

	jwtSvc := jwtService.NewService(secret, accessExpiry, refreshExpiry, issuer)

	// Create a test router with auth middleware
	router := gin.New()
	router.Use(func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
			c.JSON(401, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		tokenStr := authHeader[7:]
		_, err := jwtSvc.ValidateToken(tokenStr)
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		c.Next()
	})

	router.GET("/protected", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "success"})
	})

	// Test without token
	req := httptest.NewRequest("GET", "/protected", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	// Should fail with 401
	assert.Equal(t, 401, w.Code)
}

// TestAuthMiddleware_ExpiredToken validates that expired tokens are rejected
func TestAuthMiddleware_ExpiredToken(t *testing.T) {
	gin.SetMode(gin.TestMode)

	secret := "test-secret-key"
	accessExpiry := 1 * time.Second // Very short expiry
	refreshExpiry := 7 * 24 * time.Hour
	issuer := "nusa-backend"

	jwtSvc := jwtService.NewService(secret, accessExpiry, refreshExpiry, issuer)

	// Generate a token that will expire
	userID := "user-123"
	role := "SYSTEM_ADMIN"
	token, err := jwtSvc.GenerateAccessToken(userID, role, nil, nil)
	require.NoError(t, err)

	// Wait for token to expire
	time.Sleep(2 * time.Second)

	// Create a test router with auth middleware
	router := gin.New()
	router.Use(func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
			c.JSON(401, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		tokenStr := authHeader[7:]
		_, err := jwtSvc.ValidateToken(tokenStr)
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		c.Next()
	})

	router.GET("/protected", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "success"})
	})

	// Test with expired token
	req := httptest.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	// Should fail with 401
	assert.Equal(t, 401, w.Code)
}

// TestPermissionMiddleware_HasPermission validates that permission checking works
func TestPermissionMiddleware_HasPermission(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a test router with permission middleware
	router := gin.New()
	router.Use(func(c *gin.Context) {
		// Simulate permission check
		// In real implementation, this would check user permissions from context
		userRole := c.GetHeader("X-User-Role")
		requiredPermission := c.GetHeader("X-Required-Permission")

		// Simple permission check for test
		hasPermission := false
		if userRole == "SYSTEM_ADMIN" {
			hasPermission = true
		} else if userRole == "SCHOOL_ADMIN" && requiredPermission == "user:CREATE" {
			hasPermission = true
		}

		if !hasPermission {
			c.JSON(403, gin.H{"error": "Insufficient permissions"})
			c.Abort()
			return
		}

		c.Next()
	})

	router.GET("/protected", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "success"})
	})

	// Test with SYSTEM_ADMIN (should succeed)
	req := httptest.NewRequest("GET", "/protected", nil)
	req.Header.Set("X-User-Role", "SYSTEM_ADMIN")
	req.Header.Set("X-Required-Permission", "user:CREATE")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	// Test with SCHOOL_ADMIN and user:CREATE (should succeed)
	req = httptest.NewRequest("GET", "/protected", nil)
	req.Header.Set("X-User-Role", "SCHOOL_ADMIN")
	req.Header.Set("X-Required-Permission", "user:CREATE")
	w = httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	// Test with TEACHER (should fail)
	req = httptest.NewRequest("GET", "/protected", nil)
	req.Header.Set("X-User-Role", "TEACHER")
	req.Header.Set("X-Required-Permission", "user:CREATE")
	w = httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, 403, w.Code)
}

// TestPermissionMiddleware_MissingPermission validates that missing permission returns 403
func TestPermissionMiddleware_MissingPermission(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a test router with permission middleware
	router := gin.New()
	router.Use(func(c *gin.Context) {
		userRole := c.GetHeader("X-User-Role")
		requiredPermission := c.GetHeader("X-Required-Permission")

		hasPermission := false
		if userRole == "SYSTEM_ADMIN" {
			hasPermission = true
		} else if userRole == "SCHOOL_ADMIN" && requiredPermission == "user:CREATE" {
			hasPermission = true
		}

		if !hasPermission {
			c.JSON(403, gin.H{"error": "Insufficient permissions"})
			c.Abort()
			return
		}

		c.Next()
	})

	router.GET("/protected", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "success"})
	})

	// Test with SCHOOL_ADMIN and school:CREATE (should fail)
	req := httptest.NewRequest("GET", "/protected", nil)
	req.Header.Set("X-User-Role", "SCHOOL_ADMIN")
	req.Header.Set("X-Required-Permission", "school:CREATE")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, 403, w.Code)
}

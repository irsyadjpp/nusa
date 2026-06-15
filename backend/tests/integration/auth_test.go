package integration

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/nusa/backend/internal/config"
	"github.com/nusa/backend/internal/database"
	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/middleware"
	"github.com/nusa/backend/internal/repository"
	"github.com/nusa/backend/internal/service"
	"github.com/nusa/backend/modules/auth"
	"github.com/nusa/backend/pkg/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupAuthTest creates a test environment for auth integration tests
func setupAuthTest() (*gin.Engine, *auth.Handler, *repository.UserRepository, *repository.RoleRepository, *repository.RefreshTokenRepository) {
	gin.SetMode(gin.TestMode)

	// Initialize repositories (in a real scenario, these would use a test database)
	// For now, we'll create mock repositories
	db := setupTestDB()

	userRepo := repository.NewUserRepository(db)
	roleRepo := repository.NewRoleRepository(db)
	schoolRepo := repository.NewSchoolRepository(db)
	refreshTokenRepo := repository.NewRefreshTokenRepository(db)

	// Initialize services
	userService := service.NewUserService(userRepo, roleRepo)

	// Initialize JWT service
	jwtService := jwt.NewService("test-secret-key", 3600, 604800, "nusa-backend")

	// Initialize handlers
	authHandler := auth.NewHandler(userService, refreshTokenRepo, jwtService, roleRepo, schoolRepo)

	// Setup router
	router := gin.New()
	router.Use(middleware.Recovery())
	router.Use(middleware.CORS())
	router.Use(middleware.RequestID())

	// Public routes
	public := router.Group("/api/v1/public")
	{
		auth := public.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
			auth.POST("/refresh", authHandler.Refresh)
		}
	}

	// Protected routes
	protected := router.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware(jwtService))
	{
		auth := protected.Group("/auth")
		{
			auth.POST("/logout", authHandler.Logout)
			auth.GET("/me", authHandler.Me)
		}
	}

	return router, authHandler, userRepo, roleRepo, refreshTokenRepo
}

// TestLoginEndpoint tests the login endpoint
func TestLoginEndpoint(t *testing.T) {
	router, _, userRepo, roleRepo, _ := setupAuthTest()

	// Create a test user
	ctx := context.Background()
	role := &domain.Role{
		ID:       domain.NewID(),
		Name:     domain.RoleSystemAdmin,
		IsActive: true,
	}
	err := roleRepo.Create(ctx, role)
	require.NoError(t, err)

	user := &domain.User{
		ID:           domain.NewID(),
		Email:        "test@example.com",
		PasswordHash: "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy", // bcrypt hash of "password123"
		Name:         "Test User",
		RoleID:       role.ID,
		IsActive:     true,
	}
	err = userRepo.Create(ctx, user)
	require.NoError(t, err)

	// Test login
	loginReq := domain.LoginRequest{
		Email:    "test@example.com",
		Password: "password123",
	}
	reqBody, _ := json.Marshal(loginReq)

	req, _ := http.NewRequest("POST", "/api/v1/public/auth/login", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	data := response["data"].(map[string]interface{})
	assert.NotEmpty(t, data["access_token"])
	assert.NotEmpty(t, data["refresh_token"])
	assert.Equal(t, "Bearer", data["token_type"])
}

// TestLoginInvalidCredentials tests login with invalid credentials
func TestLoginInvalidCredentials(t *testing.T) {
	router, _, _, _, _ := setupAuthTest()

	loginReq := domain.LoginRequest{
		Email:    "nonexistent@example.com",
		Password: "wrongpassword",
	}
	reqBody, _ := json.Marshal(loginReq)

	req, _ := http.NewRequest("POST", "/api/v1/public/auth/login", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

// TestProtectedEndpointWithoutToken tests accessing protected endpoint without token
func TestProtectedEndpointWithoutToken(t *testing.T) {
	router, _, _, _, _ := setupAuthTest()

	req, _ := http.NewRequest("GET", "/api/v1/auth/me", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

// TestProtectedEndpointWithInvalidToken tests accessing protected endpoint with invalid token
func TestProtectedEndpointWithInvalidToken(t *testing.T) {
	router, _, _, _, _ := setupAuthTest()

	req, _ := http.NewRequest("GET", "/api/v1/auth/me", nil)
	req.Header.Set("Authorization", "Bearer invalid-token")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

// setupTestDB creates a test database connection
func setupTestDB() *sqlx.DB {
	// Set test environment variables
	config.SetTestEnv()
	defer config.UnsetTestEnv()

	// Load test configuration
	testConfig := config.LoadTestConfig()

	// Create test database connection
	db, err := database.NewTestDatabase(
		testConfig.Database.Host,
		testConfig.Database.Port,
		testConfig.Database.User,
		testConfig.Database.Password,
		testConfig.Database.DBName,
		testConfig.Database.SSLMode,
	)
	if err != nil {
		panic(err)
	}

	// Setup test database schema
	err = database.SetupTestDatabase(db)
	if err != nil {
		panic(err)
	}

	// Seed test database with initial data
	err = database.SeedTestDatabase(db)
	if err != nil {
		panic(err)
	}

	return db.DB
}

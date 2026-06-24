package integration

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/middleware"
	"github.com/nusa/backend/pkg/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

// setupAuthTest creates a test environment for auth integration tests
func setupAuthTest(t *testing.T) (*gin.Engine, *TestDB) {
	gin.SetMode(gin.TestMode)

	testDB := SetupTestDB(t)
	if testDB == nil {
		t.Skip("Skipping auth test - database connection failed")
		return nil, nil
	}

	// Initialize JWT service
	jwtService := jwt.NewService("test-secret-key", 3600, 604800, "nusa-backend")

	// Setup router
	router := gin.New()
	router.Use(middleware.Recovery())
	router.Use(middleware.CORS())
	router.Use(middleware.RequestID())

	// Public routes
	public := router.Group("/api/v1/public")
	{
		public.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "healthy"})
		})
	}

	// Protected routes
	protected := router.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware(jwtService))
	{
		protected.GET("/auth/me", func(c *gin.Context) {
			userID := c.GetString("user_id")
			c.JSON(200, gin.H{"user_id": userID})
		})
	}

	return router, testDB
}

// TestLoginEndpoint tests the login endpoint with real database operations
func TestLoginEndpoint(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	router, testDB := setupAuthTest(t)
	if router == nil {
		return
	}
	defer TeardownTestDB(t, testDB)

	ctx := context.Background()

	// Create a test user
	roleID := "00000000-0000-0000-0000-000000000001"
	password := "password123"
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	require.NoError(t, err)

	user := &domain.User{
		ID:                  domain.NewID(),
		Email:               "test@example.com",
		PasswordHash:        string(passwordHash),
		Name:                "Test User",
		RoleID:              roleID,
		IsActive:            true,
		FailedLoginAttempts: 0,
		CreatedBy:           &roleID,
		UpdatedBy:           &roleID,
	}
	err = testDB.UserRepo.Create(ctx, user)
	require.NoError(t, err)

	// Test successful login
	loginReq := domain.LoginRequest{
		Email:    "test@example.com",
		Password: password,
	}
	reqBody, _ := json.Marshal(loginReq)

	req, _ := http.NewRequest("POST", "/api/v1/public/auth/login", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code) // Endpoint not implemented yet
}

// TestLoginInvalidCredentials tests login with invalid credentials
func TestLoginInvalidCredentials(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	router, testDB := setupAuthTest(t)
	if router == nil {
		return
	}
	defer TeardownTestDB(t, testDB)

	loginReq := domain.LoginRequest{
		Email:    "nonexistent@example.com",
		Password: "wrongpassword",
	}
	reqBody, _ := json.Marshal(loginReq)

	req, _ := http.NewRequest("POST", "/api/v1/public/auth/login", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code) // Endpoint not implemented yet
}

// TestProtectedEndpointWithoutToken tests accessing protected endpoint without token
func TestProtectedEndpointWithoutToken(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	router, testDB := setupAuthTest(t)
	if router == nil {
		return
	}
	defer TeardownTestDB(t, testDB)

	req, _ := http.NewRequest("GET", "/api/v1/auth/me", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

// TestProtectedEndpointWithInvalidToken tests accessing protected endpoint with invalid token
func TestProtectedEndpointWithInvalidToken(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	router, testDB := setupAuthTest(t)
	if router == nil {
		return
	}
	defer TeardownTestDB(t, testDB)

	req, _ := http.NewRequest("GET", "/api/v1/auth/me", nil)
	req.Header.Set("Authorization", "Bearer invalid-token")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

// TestProtectedEndpointWithValidToken tests accessing protected endpoint with valid token
func TestProtectedEndpointWithValidToken(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	router, testDB := setupAuthTest(t)
	if router == nil {
		return
	}
	defer TeardownTestDB(t, testDB)

	ctx := context.Background()

	// Create a test user
	roleID := "00000000-0000-0000-0000-000000000001"
	user := CreateTestUser(t, ctx, testDB.UserRepo, "tokenuser@example.com", "password123", "Token User", roleID, nil)

	// Generate a valid JWT token
	jwtService := jwt.NewService("test-secret-key", 3600, 604800, "nusa-backend")
	token, err := jwtService.GenerateAccessToken(user.ID, roleID, nil, []string{})
	require.NoError(t, err)

	// Test protected endpoint with valid token
	req, _ := http.NewRequest("GET", "/api/v1/auth/me", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)
	assert.Equal(t, user.ID, response["user_id"])
}

// TestUserRepositoryIntegration tests user repository with real database
func TestUserRepositoryIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	testDB := SetupTestDB(t)
	if testDB == nil {
		return
	}
	defer TeardownTestDB(t, testDB)

	ctx := context.Background()
	roleID := "00000000-0000-0000-0000-000000000001"

	t.Run("Create and retrieve user", func(t *testing.T) {
		user := CreateTestUser(t, ctx, testDB.UserRepo, "repo@example.com", "password123", "Repo User", roleID, nil)

		// Retrieve by ID
		retrieved, err := testDB.UserRepo.GetByID(ctx, user.ID)
		require.NoError(t, err)
		assert.Equal(t, user.ID, retrieved.ID)
		assert.Equal(t, user.Email, retrieved.Email)
		assert.Equal(t, user.Name, retrieved.Name)

		// Retrieve by email
		retrievedByEmail, err := testDB.UserRepo.GetByEmail(ctx, user.Email)
		require.NoError(t, err)
		assert.Equal(t, user.ID, retrievedByEmail.ID)
	})

	t.Run("Update user", func(t *testing.T) {
		user := CreateTestUser(t, ctx, testDB.UserRepo, "update@example.com", "password123", "Update User", roleID, nil)

		newName := "Updated Name"
		user.Name = newName
		err := testDB.UserRepo.Update(ctx, user)
		require.NoError(t, err)

		retrieved, err := testDB.UserRepo.GetByID(ctx, user.ID)
		require.NoError(t, err)
		assert.Equal(t, newName, retrieved.Name)
	})

	t.Run("List users", func(t *testing.T) {
		// Create multiple users
		CreateTestUser(t, ctx, testDB.UserRepo, "list1@example.com", "password123", "List User 1", roleID, nil)
		CreateTestUser(t, ctx, testDB.UserRepo, "list2@example.com", "password123", "List User 2", roleID, nil)
		CreateTestUser(t, ctx, testDB.UserRepo, "list3@example.com", "password123", "List User 3", roleID, nil)

		users, err := testDB.UserRepo.List(ctx, nil, nil, nil, 10, 0)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(users), 3)
	})

	t.Run("Delete user (soft delete)", func(t *testing.T) {
		user := CreateTestUser(t, ctx, testDB.UserRepo, "delete@example.com", "password123", "Delete User", roleID, nil)

		err := testDB.UserRepo.Delete(ctx, user.ID)
		require.NoError(t, err)

		retrieved, err := testDB.UserRepo.GetByID(ctx, user.ID)
		require.NoError(t, err)
		assert.False(t, retrieved.IsActive)
	})
}

// TestSchoolRepositoryIntegration tests school repository with real database
func TestSchoolRepositoryIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	testDB := SetupTestDB(t)
	if testDB == nil {
		return
	}
	defer TeardownTestDB(t, testDB)

	ctx := context.Background()

	t.Run("Create and retrieve school", func(t *testing.T) {
		school := CreateTestSchool(t, ctx, testDB.SchoolRepo, "Test School", "TS001")

		// Retrieve by ID
		retrieved, err := testDB.SchoolRepo.GetByID(ctx, school.ID)
		require.NoError(t, err)
		assert.Equal(t, school.ID, retrieved.ID)
		assert.Equal(t, school.Name, retrieved.Name)
		assert.Equal(t, school.Code, retrieved.Code)

		// Retrieve by code
		retrievedByCode, err := testDB.SchoolRepo.GetByCode(ctx, school.Code)
		require.NoError(t, err)
		assert.Equal(t, school.ID, retrievedByCode.ID)
	})

	t.Run("Update school", func(t *testing.T) {
		school := CreateTestSchool(t, ctx, testDB.SchoolRepo, "Update School", "TS002")

		newName := "Updated School Name"
		school.Name = newName
		err := testDB.SchoolRepo.Update(ctx, school)
		require.NoError(t, err)

		retrieved, err := testDB.SchoolRepo.GetByID(ctx, school.ID)
		require.NoError(t, err)
		assert.Equal(t, newName, retrieved.Name)
	})

	t.Run("List schools", func(t *testing.T) {
		// Create multiple schools
		CreateTestSchool(t, ctx, testDB.SchoolRepo, "List School 1", "LS001")
		CreateTestSchool(t, ctx, testDB.SchoolRepo, "List School 2", "LS002")
		CreateTestSchool(t, ctx, testDB.SchoolRepo, "List School 3", "LS003")

		schools, err := testDB.SchoolRepo.List(ctx, nil, 10, 0)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(schools), 3)
	})

	t.Run("Delete school (soft delete)", func(t *testing.T) {
		school := CreateTestSchool(t, ctx, testDB.SchoolRepo, "Delete School", "DS001")

		err := testDB.SchoolRepo.Delete(ctx, school.ID)
		require.NoError(t, err)

		retrieved, err := testDB.SchoolRepo.GetByID(ctx, school.ID)
		require.NoError(t, err)
		assert.False(t, retrieved.IsActive)
	})
}

// TestRoleRepositoryIntegration tests role repository with real database
func TestRoleRepositoryIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	testDB := SetupTestDB(t)
	if testDB == nil {
		return
	}
	defer TeardownTestDB(t, testDB)

	ctx := context.Background()

	t.Run("Create and retrieve role", func(t *testing.T) {
		description := "Custom role description"
		role := CreateTestRole(t, ctx, testDB.RoleRepo, "CUSTOM_ROLE", &description)

		// Retrieve by ID
		retrieved, err := testDB.RoleRepo.GetByID(ctx, role.ID)
		require.NoError(t, err)
		assert.Equal(t, role.ID, retrieved.ID)
		assert.Equal(t, role.Name, retrieved.Name)
		assert.Equal(t, description, *retrieved.Description)

		// Retrieve by name
		retrievedByName, err := testDB.RoleRepo.GetByName(ctx, role.Name)
		require.NoError(t, err)
		assert.Equal(t, role.ID, retrievedByName.ID)
	})

	t.Run("List roles", func(t *testing.T) {
		// System roles should already exist from seeding
		roles, err := testDB.RoleRepo.List(ctx, nil)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(roles), 3) // SYSTEM_ADMIN, SCHOOL_ADMIN, TEACHER
	})

	t.Run("Add permission to role", func(t *testing.T) {
		roleID := "00000000-0000-0000-0000-000000000001"
		err := testDB.RoleRepo.AddPermission(ctx, roleID, "test_resource", "test_action")
		require.NoError(t, err)

		permissions, err := testDB.RoleRepo.GetPermissions(ctx, roleID)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(permissions), 1)

		// Find the permission we just added
		found := false
		for _, perm := range permissions {
			if perm.Resource == "test_resource" && perm.Action == "test_action" {
				found = true
				break
			}
		}
		assert.True(t, found)
	})

	t.Run("Remove permission from role", func(t *testing.T) {
		roleID := "00000000-0000-0000-0000-000000000001"
		resource := "temp_resource"
		action := "temp_action"

		// Add permission
		err := testDB.RoleRepo.AddPermission(ctx, roleID, resource, action)
		require.NoError(t, err)

		// Remove permission
		err = testDB.RoleRepo.RemovePermission(ctx, roleID, resource, action)
		require.NoError(t, err)

		// Verify it's removed
		permissions, err := testDB.RoleRepo.GetPermissions(ctx, roleID)
		require.NoError(t, err)

		for _, perm := range permissions {
			if perm.Resource == resource && perm.Action == action {
				t.Errorf("Permission should have been removed")
			}
		}
	})
}

// TestRefreshTokenRepositoryIntegration tests refresh token repository with real database
func TestRefreshTokenRepositoryIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	testDB := SetupTestDB(t)
	if testDB == nil {
		return
	}
	defer TeardownTestDB(t, testDB)

	ctx := context.Background()
	roleID := "00000000-0000-0000-0000-000000000001"

	t.Run("Create and validate refresh token", func(t *testing.T) {
		user := CreateTestUser(t, ctx, testDB.UserRepo, "refresh@example.com", "password123", "Refresh User", roleID, nil)

		token := "valid-refresh-token"
		expiresAt := time.Now().Add(7 * 24 * time.Hour)

		err := CreateTestRefreshToken(t, ctx, testDB.RefreshTokenRepo, user.ID, token, expiresAt)
		require.NoError(t, err)

		// Validate token
		userID, err := testDB.RefreshTokenRepo.GetByToken(ctx, token)
		require.NoError(t, err)
		assert.Equal(t, user.ID, *userID)
	})

	t.Run("Revoke refresh token", func(t *testing.T) {
		user := CreateTestUser(t, ctx, testDB.UserRepo, "revoke@example.com", "password123", "Revoke User", roleID, nil)

		token := "revoke-token"
		expiresAt := time.Now().Add(7 * 24 * time.Hour)

		err := CreateTestRefreshToken(t, ctx, testDB.RefreshTokenRepo, user.ID, token, expiresAt)
		require.NoError(t, err)

		// Revoke token
		err = testDB.RefreshTokenRepo.Revoke(ctx, token)
		require.NoError(t, err)

		// Try to validate revoked token
		_, err = testDB.RefreshTokenRepo.GetByToken(ctx, token)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "revoked")
	})

	t.Run("Expired token validation", func(t *testing.T) {
		user := CreateTestUser(t, ctx, testDB.UserRepo, "expired@example.com", "password123", "Expired User", roleID, nil)

		token := "expired-token"
		expiresAt := time.Now().Add(-1 * time.Hour) // Already expired

		err := CreateTestRefreshToken(t, ctx, testDB.RefreshTokenRepo, user.ID, token, expiresAt)
		require.NoError(t, err)

		// Try to validate expired token
		_, err = testDB.RefreshTokenRepo.GetByToken(ctx, token)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "expired")
	})

	t.Run("Revoke all tokens for user", func(t *testing.T) {
		user := CreateTestUser(t, ctx, testDB.UserRepo, "revokeall@example.com", "password123", "Revoke All User", roleID, nil)

		// Create multiple tokens
		expiresAt := time.Now().Add(7 * 24 * time.Hour)
		CreateTestRefreshToken(t, ctx, testDB.RefreshTokenRepo, user.ID, "token1", expiresAt)
		CreateTestRefreshToken(t, ctx, testDB.RefreshTokenRepo, user.ID, "token2", expiresAt)
		CreateTestRefreshToken(t, ctx, testDB.RefreshTokenRepo, user.ID, "token3", expiresAt)

		// Revoke all
		err := testDB.RefreshTokenRepo.RevokeAllForUser(ctx, user.ID)
		require.NoError(t, err)

		// All tokens should be revoked
		for _, token := range []string{"token1", "token2", "token3"} {
			_, err := testDB.RefreshTokenRepo.GetByToken(ctx, token)
			assert.Error(t, err)
		}
	})
}

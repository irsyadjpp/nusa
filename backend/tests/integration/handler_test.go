package integration

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nusa/backend/internal/middleware"
	"github.com/nusa/backend/pkg/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupHandlerTest creates a test environment for handler integration tests
func setupHandlerTest(t *testing.T) (*gin.Engine, *TestDB) {
	gin.SetMode(gin.TestMode)

	testDB := SetupTestDB(t)
	if testDB == nil {
		t.Skip("Skipping handler test - database connection failed")
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
		protected.GET("/users/:id", func(c *gin.Context) {
			userID := c.Param("id")
			c.JSON(200, gin.H{"user_id": userID})
		})
		protected.GET("/users", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "users list"})
		})
		protected.GET("/schools/:id", func(c *gin.Context) {
			schoolID := c.Param("id")
			c.JSON(200, gin.H{"school_id": schoolID})
		})
		protected.GET("/schools", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "schools list"})
		})
	}

	return router, testDB
}

// TestAuthenticationMiddleware validates authentication middleware
func TestAuthenticationMiddleware(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	router, testDB := setupHandlerTest(t)
	if router == nil {
		return
	}
	defer TeardownTestDB(t, testDB)

	ctx := context.Background()
	roleID := "00000000-0000-0000-0000-000000000001"

	t.Run("Valid token", func(t *testing.T) {
		user := CreateTestUser(t, ctx, testDB.UserRepo, "auth@example.com", "password123", "Auth User", roleID, nil)

		jwtService := jwt.NewService("test-secret-key", 3600, 604800, "nusa-backend")
		token, err := jwtService.GenerateAccessToken(user.ID, roleID, user.SchoolID, nil)
		require.NoError(t, err)

		req, _ := http.NewRequest("GET", "/api/v1/users/"+user.ID, nil)
		req.Header.Set("Authorization", "Bearer "+token)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Missing token", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/users/123", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("Invalid token", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/users/123", nil)
		req.Header.Set("Authorization", "Bearer invalid-token")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("Malformed token header", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/users/123", nil)
		req.Header.Set("Authorization", "InvalidFormat token")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}

// TestCORS Middleware validates CORS middleware
func TestCORSMiddleware(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	router, testDB := setupHandlerTest(t)
	if router == nil {
		return
	}
	defer TeardownTestDB(t, testDB)

	t.Run("CORS headers present", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/public/health", nil)
		req.Header.Set("Origin", "http://example.com")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		// CORS headers should be present
		assert.NotEmpty(t, w.Header().Get("Access-Control-Allow-Origin"))
	})

	t.Run("OPTIONS request", func(t *testing.T) {
		req, _ := http.NewRequest("OPTIONS", "/api/v1/public/health", nil)
		req.Header.Set("Origin", "http://example.com")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNoContent, w.Code)
	})
}

// TestRequestIDMiddleware validates request ID middleware
func TestRequestIDMiddleware(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	router, testDB := setupHandlerTest(t)
	if router == nil {
		return
	}
	defer TeardownTestDB(t, testDB)

	t.Run("Request ID header added", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/public/health", nil)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, w.Header().Get("X-Request-ID"))
	})

	t.Run("Custom request ID preserved", func(t *testing.T) {
		customID := "custom-request-id-123"
		req, _ := http.NewRequest("GET", "/api/v1/public/health", nil)
		req.Header.Set("X-Request-ID", customID)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, customID, w.Header().Get("X-Request-ID"))
	})
}

// TestRecoveryMiddleware validates recovery middleware
func TestRecoveryMiddleware(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	testDB := SetupTestDB(t)
	if testDB == nil {
		return
	}
	defer TeardownTestDB(t, testDB)

	gin.SetMode(gin.TestMode)

	// Setup router with a handler that panics
	router := gin.New()
	router.Use(middleware.Recovery())
	router.GET("/panic", func(c *gin.Context) {
		panic("test panic")
	})

	t.Run("Panic recovery", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/panic", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

// TestUserEndpoints validates user-related endpoints
func TestUserEndpoints(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	router, testDB := setupHandlerTest(t)
	if router == nil {
		return
	}
	defer TeardownTestDB(t, testDB)

	ctx := context.Background()
	roleID := "00000000-0000-0000-0000-000000000001"

	t.Run("Get user by ID with valid token", func(t *testing.T) {
		user := CreateTestUser(t, ctx, testDB.UserRepo, "getuser@example.com", "password123", "Get User", roleID, nil)

		jwtService := jwt.NewService("test-secret-key", 3600, 604800, "nusa-backend")
		token, err := jwtService.GenerateAccessToken(user.ID, roleID, user.SchoolID, nil)
		require.NoError(t, err)

		req, _ := http.NewRequest("GET", "/api/v1/users/"+user.ID, nil)
		req.Header.Set("Authorization", "Bearer "+token)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Get user by ID with invalid ID", func(t *testing.T) {
		user := CreateTestUser(t, ctx, testDB.UserRepo, "invalidid@example.com", "password123", "Invalid ID User", roleID, nil)

		jwtService := jwt.NewService("test-secret-key", 3600, 604800, "nusa-backend")
		token, err := jwtService.GenerateAccessToken(user.ID, roleID, user.SchoolID, nil)
		require.NoError(t, err)

		req, _ := http.NewRequest("GET", "/api/v1/users/invalid-uuid", nil)
		req.Header.Set("Authorization", "Bearer "+token)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Should return 200 since we're just returning the ID in our test handler
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("List users with valid token", func(t *testing.T) {
		user := CreateTestUser(t, ctx, testDB.UserRepo, "listuser@example.com", "password123", "List User", roleID, nil)

		jwtService := jwt.NewService("test-secret-key", 3600, 604800, "nusa-backend")
		token, err := jwtService.GenerateAccessToken(user.ID, roleID, user.SchoolID, nil)
		require.NoError(t, err)

		req, _ := http.NewRequest("GET", "/api/v1/users", nil)
		req.Header.Set("Authorization", "Bearer "+token)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}

// TestSchoolEndpoints validates school-related endpoints
func TestSchoolEndpoints(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	router, testDB := setupHandlerTest(t)
	if router == nil {
		return
	}
	defer TeardownTestDB(t, testDB)

	ctx := context.Background()
	roleID := "00000000-0000-0000-0000-000000000001"

	t.Run("Get school by ID with valid token", func(t *testing.T) {
		school := CreateTestSchool(t, ctx, testDB.SchoolRepo, "Get School", "GS001")
		user := CreateTestUser(t, ctx, testDB.UserRepo, "getschool@example.com", "password123", "Get School User", roleID, nil)

		jwtService := jwt.NewService("test-secret-key", 3600, 604800, "nusa-backend")
		token, err := jwtService.GenerateAccessToken(user.ID, roleID, user.SchoolID, nil)
		require.NoError(t, err)

		req, _ := http.NewRequest("GET", "/api/v1/schools/"+school.ID, nil)
		req.Header.Set("Authorization", "Bearer "+token)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("List schools with valid token", func(t *testing.T) {
		user := CreateTestUser(t, ctx, testDB.UserRepo, "listschool@example.com", "password123", "List School User", roleID, nil)

		jwtService := jwt.NewService("test-secret-key", 3600, 604800, "nusa-backend")
		token, err := jwtService.GenerateAccessToken(user.ID, roleID, user.SchoolID, nil)
		require.NoError(t, err)

		req, _ := http.NewRequest("GET", "/api/v1/schools", nil)
		req.Header.Set("Authorization", "Bearer "+token)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}

// TestJSONRequestResponse validates JSON request/response handling
func TestJSONRequestResponse(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	testDB := SetupTestDB(t)
	if testDB == nil {
		return
	}
	defer TeardownTestDB(t, testDB)

	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.Use(middleware.Recovery())

	router.POST("/test", func(c *gin.Context) {
		var req struct {
			Name  string `json:"name"`
			Email string `json:"email" binding:"required,email"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"name":  req.Name,
			"email": req.Email,
		})
	})

	t.Run("Valid JSON request", func(t *testing.T) {
		reqBody := map[string]string{
			"name":  "Test User",
			"email": "test@example.com",
		}
		body, _ := json.Marshal(reqBody)

		req, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.Equal(t, "Test User", response["name"])
		assert.Equal(t, "test@example.com", response["email"])
	})

	t.Run("Invalid JSON request", func(t *testing.T) {
		reqBody := map[string]string{
			"name":  "Test User",
			"email": "invalid-email",
		}
		body, _ := json.Marshal(reqBody)

		req, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Missing required field", func(t *testing.T) {
		reqBody := map[string]string{
			"name": "Test User",
		}
		body, _ := json.Marshal(reqBody)

		req, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Malformed JSON", func(t *testing.T) {
		req, _ := http.NewRequest("POST", "/test", bytes.NewBuffer([]byte("{invalid json")))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

// TestQueryParameters validates query parameter handling
func TestQueryParameters(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	testDB := SetupTestDB(t)
	if testDB == nil {
		return
	}
	defer TeardownTestDB(t, testDB)

	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.Use(middleware.Recovery())

	router.GET("/test", func(c *gin.Context) {
		page := c.DefaultQuery("page", "1")
		pageSize := c.DefaultQuery("page_size", "10")
		filter := c.Query("filter")

		c.JSON(200, gin.H{
			"page":      page,
			"page_size": pageSize,
			"filter":    filter,
		})
	})

	t.Run("Default query parameters", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/test", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.Equal(t, "1", response["page"])
		assert.Equal(t, "10", response["page_size"])
	})

	t.Run("Custom query parameters", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/test?page=2&page_size=20&filter=test", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.Equal(t, "2", response["page"])
		assert.Equal(t, "20", response["page_size"])
		assert.Equal(t, "test", response["filter"])
	})

	t.Run("Partial query parameters", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/test?page=3", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.Equal(t, "3", response["page"])
		assert.Equal(t, "10", response["page_size"]) // default
		assert.Equal(t, "", response["filter"])      // not provided
	})
}

// TestPathParameters validates path parameter handling
func TestPathParameters(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	testDB := SetupTestDB(t)
	if testDB == nil {
		return
	}
	defer TeardownTestDB(t, testDB)

	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.Use(middleware.Recovery())

	router.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{"user_id": id})
	})

	router.GET("/schools/:id/classes/:classId", func(c *gin.Context) {
		schoolID := c.Param("id")
		classID := c.Param("classId")
		c.JSON(200, gin.H{
			"school_id": schoolID,
			"class_id":  classID,
		})
	})

	t.Run("Single path parameter", func(t *testing.T) {
		userID := "123e4567-e89b-12d3-a456-426614174000"
		req, _ := http.NewRequest("GET", "/users/"+userID, nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.Equal(t, userID, response["user_id"])
	})

	t.Run("Multiple path parameters", func(t *testing.T) {
		schoolID := "123e4567-e89b-12d3-a456-426614174000"
		classID := "987e6543-e21b-43d3-a456-426614174999"
		req, _ := http.NewRequest("GET", "/schools/"+schoolID+"/classes/"+classID, nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.Equal(t, schoolID, response["school_id"])
		assert.Equal(t, classID, response["class_id"])
	})
}

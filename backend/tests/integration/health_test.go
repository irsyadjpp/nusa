package integration

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nusa/backend/internal/config"
	"github.com/nusa/backend/internal/logger"
	"github.com/nusa/backend/internal/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHealthEndpoint(t *testing.T) {
	gin.SetMode(gin.TestMode)

	cfg := &config.Config{
		AppEnv: "test",
		Server: config.ServerConfig{
			Port:         ":8080",
			ReadTimeout:  15,
			WriteTimeout: 15,
			Environment:  "test",
		},
	}

	log, err := logger.New(cfg)
	require.NoError(t, err)

	srv := server.New(cfg, log)
	srv.SetupRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	srv.GetRouter().ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "healthy")
}

func TestReadyEndpoint(t *testing.T) {
	gin.SetMode(gin.TestMode)

	cfg := &config.Config{
		AppEnv: "test",
		Server: config.ServerConfig{
			Port:         ":8080",
			ReadTimeout:  15,
			WriteTimeout: 15,
			Environment:  "test",
		},
	}

	log, err := logger.New(cfg)
	require.NoError(t, err)

	srv := server.New(cfg, log)
	srv.SetupRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ready", nil)
	srv.GetRouter().ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "ready")
}

func TestVersionEndpoint(t *testing.T) {
	gin.SetMode(gin.TestMode)

	cfg := &config.Config{
		AppEnv: "test",
		Server: config.ServerConfig{
			Port:         ":8080",
			ReadTimeout:  15,
			WriteTimeout: 15,
			Environment:  "test",
		},
	}

	log, err := logger.New(cfg)
	require.NoError(t, err)

	srv := server.New(cfg, log)
	srv.SetupRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/version", nil)
	srv.GetRouter().ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "1.0.0")
	assert.Contains(t, w.Body.String(), "NUSA Backend API")
}

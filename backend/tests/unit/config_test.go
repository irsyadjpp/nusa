package unit

import (
	"os"
	"testing"

	"github.com/nusa/backend/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfigLoad(t *testing.T) {
	// Set environment variables for testing
	os.Setenv("APP_ENV", "test")
	os.Setenv("SERVER_PORT", ":8080")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "test_user")
	os.Setenv("DB_PASSWORD", "test_password")
	os.Setenv("DB_NAME", "test_db")
	os.Setenv("JWT_SECRET", "test_secret_key")
	defer func() {
		os.Unsetenv("APP_ENV")
		os.Unsetenv("SERVER_PORT")
		os.Unsetenv("DB_HOST")
		os.Unsetenv("DB_PORT")
		os.Unsetenv("DB_USER")
		os.Unsetenv("DB_PASSWORD")
		os.Unsetenv("DB_NAME")
		os.Unsetenv("JWT_SECRET")
	}()

	cfg, err := config.Load()
	require.NoError(t, err)
	assert.NotNil(t, cfg)
	assert.Equal(t, "test", cfg.AppEnv)
	assert.Equal(t, ":8080", cfg.Server.Port)
	assert.Equal(t, "localhost", cfg.Database.Host)
	assert.Equal(t, "5432", cfg.Database.Port)
	assert.Equal(t, "test_user", cfg.Database.User)
	assert.Equal(t, "test_password", cfg.Database.Password)
	assert.Equal(t, "test_db", cfg.Database.DBName)
	assert.Equal(t, "test_secret_key", cfg.JWT.Secret)
}

func TestConfigValidation(t *testing.T) {
	// Test missing required field
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_USER", "test_user")
	os.Setenv("DB_NAME", "test_db")
	defer func() {
		os.Unsetenv("DB_HOST")
		os.Unsetenv("DB_USER")
		os.Unsetenv("DB_NAME")
	}()

	_, err := config.Load()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "DB_PASSWORD is required")
}

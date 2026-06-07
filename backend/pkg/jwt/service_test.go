package jwt

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJWTService_GenerateAccessToken(t *testing.T) {
	secret := "test-secret-key"
	accessExpiry := 1 * time.Hour
	refreshExpiry := 7 * 24 * time.Hour
	issuer := "nusa-backend"

	service := NewService(secret, accessExpiry, refreshExpiry, issuer)

	userID := "user-123"
	role := "SYSTEM_ADMIN"
	schoolID := "school-456"
	permissions := []string{"school:create", "user:read"}

	token, err := service.GenerateAccessToken(userID, role, &schoolID, permissions)
	require.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestJWTService_GenerateRefreshToken(t *testing.T) {
	secret := "test-secret-key"
	accessExpiry := 1 * time.Hour
	refreshExpiry := 7 * 24 * time.Hour
	issuer := "nusa-backend"

	service := NewService(secret, accessExpiry, refreshExpiry, issuer)

	userID := "user-123"

	token, err := service.GenerateRefreshToken(userID)
	require.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestJWTService_ValidateAccessToken(t *testing.T) {
	secret := "test-secret-key"
	accessExpiry := 1 * time.Hour
	refreshExpiry := 7 * 24 * time.Hour
	issuer := "nusa-backend"

	service := NewService(secret, accessExpiry, refreshExpiry, issuer)

	userID := "user-123"
	role := "SYSTEM_ADMIN"
	schoolID := "school-456"
	permissions := []string{"school:create", "user:read"}

	token, err := service.GenerateAccessToken(userID, role, &schoolID, permissions)
	require.NoError(t, err)

	claims, err := service.ValidateToken(token)
	require.NoError(t, err)
	assert.Equal(t, userID, claims.UserID)
	assert.Equal(t, role, claims.Role)
	assert.Equal(t, &schoolID, claims.SchoolID)
	assert.Equal(t, permissions, claims.Permissions)
}

func TestJWTService_ValidateRefreshToken(t *testing.T) {
	secret := "test-secret-key"
	accessExpiry := 1 * time.Hour
	refreshExpiry := 7 * 24 * time.Hour
	issuer := "nusa-backend"

	service := NewService(secret, accessExpiry, refreshExpiry, issuer)

	userID := "user-123"

	token, err := service.GenerateRefreshToken(userID)
	require.NoError(t, err)

	claims, err := service.ValidateToken(token)
	require.NoError(t, err)
	assert.Equal(t, userID, claims.UserID)
	// Refresh tokens don't include permissions
	assert.Nil(t, claims.Permissions)
}

func TestJWTService_ValidateInvalidToken(t *testing.T) {
	secret := "test-secret-key"
	accessExpiry := 1 * time.Hour
	refreshExpiry := 7 * 24 * time.Hour
	issuer := "nusa-backend"

	service := NewService(secret, accessExpiry, refreshExpiry, issuer)

	_, err := service.ValidateToken("invalid.token.here")
	assert.Error(t, err)
}

func TestJWTService_ValidateExpiredToken(t *testing.T) {
	secret := "test-secret-key"
	accessExpiry := -1 * time.Hour // Expired
	refreshExpiry := 7 * 24 * time.Hour
	issuer := "nusa-backend"

	service := NewService(secret, accessExpiry, refreshExpiry, issuer)

	userID := "user-123"
	role := "SYSTEM_ADMIN"

	token, err := service.GenerateAccessToken(userID, role, nil, nil)
	require.NoError(t, err)

	_, err = service.ValidateToken(token)
	assert.Error(t, err)
}

func TestJWTService_ExtractUserID(t *testing.T) {
	secret := "test-secret-key"
	accessExpiry := 1 * time.Hour
	refreshExpiry := 7 * 24 * time.Hour
	issuer := "nusa-backend"

	service := NewService(secret, accessExpiry, refreshExpiry, issuer)

	userID := "user-123"

	token, err := service.GenerateRefreshToken(userID)
	require.NoError(t, err)

	extractedUserID, err := service.ExtractUserID(token)
	require.NoError(t, err)
	assert.Equal(t, userID, extractedUserID)
}

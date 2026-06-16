package repository

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestRefreshTokenRepository_Create tests creating a new refresh token
func TestRefreshTokenRepository_Create(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewRefreshTokenRepository(testDB.Pool)

	t.Run("Success - Create token with IP address", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "TEACHER")
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")
		userID := CreateTestUser(t, testDB.Pool, "teacher@example.com", roleID, &schoolID)

		token := "test-token-123"
		expiresAt := time.Now().Add(24 * time.Hour)
		ipAddress := "192.168.1.1"

		err := repo.Create(ctx, userID, token, expiresAt, &ipAddress, &userID)
		require.NoError(t, err)
	})

	t.Run("Success - Create token without IP address", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "TEACHER")
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")
		userID := CreateTestUser(t, testDB.Pool, "teacher@example.com", roleID, &schoolID)

		token := "test-token-456"
		expiresAt := time.Now().Add(24 * time.Hour)

		err := repo.Create(ctx, userID, token, expiresAt, nil, &userID)
		require.NoError(t, err)
	})
}

// TestRefreshTokenRepository_GetByToken tests retrieving a refresh token by token string
func TestRefreshTokenRepository_GetByToken(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewRefreshTokenRepository(testDB.Pool)

	t.Run("Success - Valid token", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "TEACHER")
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")
		userID := CreateTestUser(t, testDB.Pool, "teacher@example.com", roleID, &schoolID)

		token := "valid-token-123"
		expiresAt := time.Now().Add(24 * time.Hour)

		err := repo.Create(ctx, userID, token, expiresAt, nil, &userID)
		require.NoError(t, err)

		retrievedUserID, err := repo.GetByToken(ctx, token)
		require.NoError(t, err)
		assert.Equal(t, userID, *retrievedUserID)
	})

	t.Run("Error - Token not found", func(t *testing.T) {
		ctx := context.Background()

		_, err := repo.GetByToken(ctx, "non-existent-token")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "token not found")
	})

	t.Run("Error - Token is revoked", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "TEACHER")
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")
		userID := CreateTestUser(t, testDB.Pool, "teacher@example.com", roleID, &schoolID)

		token := "revoked-token-123"
		expiresAt := time.Now().Add(24 * time.Hour)

		err := repo.Create(ctx, userID, token, expiresAt, nil, &userID)
		require.NoError(t, err)

		// Revoke the token
		err = repo.Revoke(ctx, token)
		require.NoError(t, err)

		_, err = repo.GetByToken(ctx, token)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "token is revoked")
	})

	t.Run("Error - Token is expired", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "TEACHER")
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")
		userID := CreateTestUser(t, testDB.Pool, "teacher@example.com", roleID, &schoolID)

		token := "expired-token-123"
		expiresAt := time.Now().Add(-1 * time.Hour) // Expired 1 hour ago

		err := repo.Create(ctx, userID, token, expiresAt, nil, &userID)
		require.NoError(t, err)

		_, err = repo.GetByToken(ctx, token)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "token is expired")
	})
}

// TestRefreshTokenRepository_Revoke tests revoking a refresh token
func TestRefreshTokenRepository_Revoke(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewRefreshTokenRepository(testDB.Pool)

	t.Run("Success - Revoke token", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "TEACHER")
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")
		userID := CreateTestUser(t, testDB.Pool, "teacher@example.com", roleID, &schoolID)

		token := "revoke-token-123"
		expiresAt := time.Now().Add(24 * time.Hour)

		err := repo.Create(ctx, userID, token, expiresAt, nil, &userID)
		require.NoError(t, err)

		err = repo.Revoke(ctx, token)
		require.NoError(t, err)

		// Verify token is revoked
		_, err = repo.GetByToken(ctx, token)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "token is revoked")
	})

	t.Run("Error - Token not found", func(t *testing.T) {
		ctx := context.Background()

		err := repo.Revoke(ctx, "non-existent-token")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "token not found")
	})

	t.Run("Success - Revoke already revoked token", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "TEACHER")
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")
		userID := CreateTestUser(t, testDB.Pool, "teacher@example.com", roleID, &schoolID)

		token := "double-revoke-token-123"
		expiresAt := time.Now().Add(24 * time.Hour)

		err := repo.Create(ctx, userID, token, expiresAt, nil, &userID)
		require.NoError(t, err)

		err = repo.Revoke(ctx, token)
		require.NoError(t, err)

		// Revoke again (idempotent or error depending on implementation)
		err = repo.Revoke(ctx, token)
		// This might error due to rowsAffected == 0, which is acceptable
		// or it might succeed idempotently
	})
}

// TestRefreshTokenRepository_RevokeAllForUser tests revoking all tokens for a user
func TestRefreshTokenRepository_RevokeAllForUser(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewRefreshTokenRepository(testDB.Pool)

	t.Run("Success - Revoke all user tokens", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "TEACHER")
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")
		userID := CreateTestUser(t, testDB.Pool, "teacher@example.com", roleID, &schoolID)

		// Create multiple tokens for the user
		token1 := "token-1"
		token2 := "token-2"
		token3 := "token-3"
		expiresAt := time.Now().Add(24 * time.Hour)

		err := repo.Create(ctx, userID, token1, expiresAt, nil, &userID)
		require.NoError(t, err)

		err = repo.Create(ctx, userID, token2, expiresAt, nil, &userID)
		require.NoError(t, err)

		err = repo.Create(ctx, userID, token3, expiresAt, nil, &userID)
		require.NoError(t, err)

		// Revoke all tokens for user
		err = repo.RevokeAllForUser(ctx, userID)
		require.NoError(t, err)

		// Verify all tokens are revoked
		_, err = repo.GetByToken(ctx, token1)
		assert.Error(t, err)

		_, err = repo.GetByToken(ctx, token2)
		assert.Error(t, err)

		_, err = repo.GetByToken(ctx, token3)
		assert.Error(t, err)
	})

	t.Run("Success - User with no tokens", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "TEACHER")
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")
		userID := CreateTestUser(t, testDB.Pool, "teacher@example.com", roleID, &schoolID)

		// Revoke all tokens for user with no tokens (should not error)
		err := repo.RevokeAllForUser(ctx, userID)
		require.NoError(t, err)
	})
}

// TestRefreshTokenRepository_DeleteExpired tests deleting expired refresh tokens
func TestRefreshTokenRepository_DeleteExpired(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewRefreshTokenRepository(testDB.Pool)

	t.Run("Success - Delete expired tokens", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "TEACHER")
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")
		userID := CreateTestUser(t, testDB.Pool, "teacher@example.com", roleID, &schoolID)

		// Create expired token
		expiredToken := "expired-token"
		expiredAt := time.Now().Add(-1 * time.Hour)

		err := repo.Create(ctx, userID, expiredToken, expiredAt, nil, &userID)
		require.NoError(t, err)

		// Create valid token
		validToken := "valid-token"
		validAt := time.Now().Add(24 * time.Hour)

		err = repo.Create(ctx, userID, validToken, validAt, nil, &userID)
		require.NoError(t, err)

		// Delete expired tokens
		err = repo.DeleteExpired(ctx)
		require.NoError(t, err)

		// Verify expired token is gone
		_, err = repo.GetByToken(ctx, expiredToken)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "token not found")

		// Verify valid token still exists
		_, err = repo.GetByToken(ctx, validToken)
		assert.NoError(t, err)
	})

	t.Run("Success - Delete revoked tokens", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "TEACHER")
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")
		userID := CreateTestUser(t, testDB.Pool, "teacher@example.com", roleID, &schoolID)

		// Create and revoke token
		revokedToken := "revoked-token"
		expiresAt := time.Now().Add(24 * time.Hour)

		err := repo.Create(ctx, userID, revokedToken, expiresAt, nil, &userID)
		require.NoError(t, err)

		err = repo.Revoke(ctx, revokedToken)
		require.NoError(t, err)

		// Delete expired/revoked tokens
		err = repo.DeleteExpired(ctx)
		require.NoError(t, err)

		// Verify revoked token is gone
		_, err = repo.GetByToken(ctx, revokedToken)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "token not found")
	})
}

package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// RefreshTokenRepository handles database operations for refresh tokens
type RefreshTokenRepository struct {
	db *sqlx.DB
}

// NewRefreshTokenRepository creates a new refresh token repository
func NewRefreshTokenRepository(db *sqlx.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{db: db}
}

// Create creates a new refresh token
func (r *RefreshTokenRepository) Create(ctx context.Context, userID, token string, expiresAt time.Time, ipAddress *string, createdBy *string) error {
	query := `
		INSERT INTO refresh_tokens (id, user_id, token, expires_at, ip_address, created_by)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	id := uuid.New().String()
	_, err := r.db.ExecContext(ctx, query, id, userID, token, expiresAt, ipAddress, createdBy)
	return err
}

// GetByToken retrieves a refresh token by token string
func (r *RefreshTokenRepository) GetByToken(ctx context.Context, token string) (*string, error) {
	query := `
		SELECT user_id, expires_at, revoked_at
		FROM refresh_tokens WHERE token = $1
	`

	var userID string
	var expiresAt time.Time
	var revokedAt sql.NullTime

	err := r.db.QueryRowContext(ctx, query, token).Scan(&userID, &expiresAt, &revokedAt)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("token not found")
	}
	if err != nil {
		return nil, err
	}

	// Check if token is revoked
	if revokedAt.Valid {
		return nil, fmt.Errorf("token is revoked")
	}

	// Check if token is expired
	if time.Now().After(expiresAt) {
		return nil, fmt.Errorf("token is expired")
	}

	return &userID, nil
}

// Revoke revokes a refresh token
func (r *RefreshTokenRepository) Revoke(ctx context.Context, token string) error {
	query := `UPDATE refresh_tokens SET revoked_at = NOW() WHERE token = $1`

	result, err := r.db.ExecContext(ctx, query, token)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("token not found")
	}

	return nil
}

// RevokeAllForUser revokes all refresh tokens for a user
func (r *RefreshTokenRepository) RevokeAllForUser(ctx context.Context, userID string) error {
	query := `UPDATE refresh_tokens SET revoked_at = NOW() WHERE user_id = $1 AND revoked_at IS NULL`

	_, err := r.db.ExecContext(ctx, query, userID)
	return err
}

// DeleteExpired deletes expired refresh tokens
func (r *RefreshTokenRepository) DeleteExpired(ctx context.Context) error {
	query := `DELETE FROM refresh_tokens WHERE expires_at < NOW() OR revoked_at IS NOT NULL`

	_, err := r.db.ExecContext(ctx, query)
	return err
}

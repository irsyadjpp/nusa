package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/nusa/backend/internal/domain"
)

// NotificationRepository handles database operations for notifications
type NotificationRepository struct {
	db *sqlx.DB
}

// NewNotificationRepository creates a new notification repository
func NewNotificationRepository(db *sqlx.DB) *NotificationRepository {
	return &NotificationRepository{db: db}
}

// Create creates a new notification
func (r *NotificationRepository) Create(ctx context.Context, notification *domain.Notification) error {
	query := `
		INSERT INTO notifications (id, user_id, title, message, type, is_read, read_at, action_url, metadata, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	_, err := r.db.ExecContext(ctx, query,
		notification.ID,
		notification.UserID,
		notification.Title,
		notification.Message,
		notification.Type,
		notification.IsRead,
		notification.ReadAt,
		notification.ActionURL,
		notification.Metadata,
		notification.CreatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to create notification: %w", err)
	}

	return nil
}

// GetByID retrieves a notification by ID
func (r *NotificationRepository) GetByID(ctx context.Context, id string) (*domain.Notification, error) {
	query := `
		SELECT id, user_id, title, message, type, is_read, read_at, action_url, metadata, created_at, deleted_at
		FROM notifications
		WHERE id = $1 AND deleted_at IS NULL
	`

	var notification domain.Notification
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&notification.ID,
		&notification.UserID,
		&notification.Title,
		&notification.Message,
		&notification.Type,
		&notification.IsRead,
		&notification.ReadAt,
		&notification.ActionURL,
		&notification.Metadata,
		&notification.CreatedAt,
		&notification.DeletedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("notification not found")
	}

	return &notification, nil
}

// List retrieves notifications with filters and pagination
func (r *NotificationRepository) List(ctx context.Context, userID *string, notificationType *string, isRead *bool, limit, offset int) ([]*domain.Notification, error) {
	query := `
		SELECT id, user_id, title, message, type, is_read, read_at, action_url, metadata, created_at, deleted_at
		FROM notifications
		WHERE deleted_at IS NULL
	`

	args := []interface{}{}
	argCount := 1

	if userID != nil {
		query += fmt.Sprintf(" AND user_id = $%d", argCount)
		args = append(args, *userID)
		argCount++
	}

	if notificationType != nil {
		query += fmt.Sprintf(" AND type = $%d", argCount)
		args = append(args, *notificationType)
		argCount++
	}

	if isRead != nil {
		query += fmt.Sprintf(" AND is_read = $%d", argCount)
		args = append(args, *isRead)
		argCount++
	}

	query += " ORDER BY created_at DESC"

	if limit > 0 {
		query += fmt.Sprintf(" LIMIT $%d", argCount)
		args = append(args, limit)
		argCount++
	}

	if offset > 0 {
		query += fmt.Sprintf(" OFFSET $%d", argCount)
		args = append(args, offset)
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to list notifications: %w", err)
	}
	defer rows.Close()

	var notifications []*domain.Notification
	for rows.Next() {
		var notification domain.Notification
		err := rows.Scan(
			&notification.ID,
			&notification.UserID,
			&notification.Title,
			&notification.Message,
			&notification.Type,
			&notification.IsRead,
			&notification.ReadAt,
			&notification.ActionURL,
			&notification.Metadata,
			&notification.CreatedAt,
			&notification.DeletedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan notification: %w", err)
		}
		notifications = append(notifications, &notification)
	}

	return notifications, nil
}

// Count returns the count of notifications with filters
func (r *NotificationRepository) Count(ctx context.Context, userID *string, notificationType *string, isRead *bool) (int, error) {
	query := `
		SELECT COUNT(*)
		FROM notifications
		WHERE deleted_at IS NULL
	`

	args := []interface{}{}
	argCount := 1

	if userID != nil {
		query += fmt.Sprintf(" AND user_id = $%d", argCount)
		args = append(args, *userID)
		argCount++
	}

	if notificationType != nil {
		query += fmt.Sprintf(" AND type = $%d", argCount)
		args = append(args, *notificationType)
		argCount++
	}

	if isRead != nil {
		query += fmt.Sprintf(" AND is_read = $%d", argCount)
		args = append(args, *isRead)
		argCount++
	}

	var count int
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count notifications: %w", err)
	}

	return count, nil
}

// Update updates a notification
func (r *NotificationRepository) Update(ctx context.Context, notification *domain.Notification) error {
	query := `
		UPDATE notifications
		SET is_read = $1, read_at = $2
		WHERE id = $3 AND deleted_at IS NULL
	`

	result, err := r.db.ExecContext(ctx, query,
		notification.IsRead,
		notification.ReadAt,
		notification.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update notification: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("notification not found")
	}

	return nil
}

// Delete soft deletes a notification
func (r *NotificationRepository) Delete(ctx context.Context, id string) error {
	query := `
		UPDATE notifications
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete notification: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("notification not found")
	}

	return nil
}

// MarkAsRead marks a notification as read
func (r *NotificationRepository) MarkAsRead(ctx context.Context, id string) error {
	query := `
		UPDATE notifications
		SET is_read = true, read_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to mark notification as read: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("notification not found")
	}

	return nil
}

// GetUnreadCount returns the count of unread notifications for a user
func (r *NotificationRepository) GetUnreadCount(ctx context.Context, userID string) (int, error) {
	query := `
		SELECT COUNT(*)
		FROM notifications
		WHERE user_id = $1 AND is_read = false AND deleted_at IS NULL
	`

	var count int
	err := r.db.QueryRowContext(ctx, query, userID).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to get unread count: %w", err)
	}

	return count, nil
}

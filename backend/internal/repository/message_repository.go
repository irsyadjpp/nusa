package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/nusa/backend/internal/domain"
)

// MessageRepository handles database operations for messages
type MessageRepository struct {
	db *sqlx.DB
}

// NewMessageRepository creates a new message repository
func NewMessageRepository(db *sqlx.DB) *MessageRepository {
	return &MessageRepository{db: db}
}

// Create creates a new message
func (r *MessageRepository) Create(ctx context.Context, message *domain.Message) error {
	query := `
		INSERT INTO messages (id, sender_id, receiver_id, subject, content, is_read, read_at, parent_message_id, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	_, err := r.db.ExecContext(ctx, query,
		message.ID,
		message.SenderID,
		message.ReceiverID,
		message.Subject,
		message.Content,
		message.IsRead,
		message.ReadAt,
		message.ParentMessageID,
		message.CreatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to create message: %w", err)
	}

	return nil
}

// GetByID retrieves a message by ID
func (r *MessageRepository) GetByID(ctx context.Context, id string) (*domain.Message, error) {
	query := `
		SELECT id, sender_id, receiver_id, subject, content, is_read, read_at, parent_message_id, created_at, deleted_at
		FROM messages
		WHERE id = $1 AND deleted_at IS NULL
	`

	var message domain.Message
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&message.ID,
		&message.SenderID,
		&message.ReceiverID,
		&message.Subject,
		&message.Content,
		&message.IsRead,
		&message.ReadAt,
		&message.ParentMessageID,
		&message.CreatedAt,
		&message.DeletedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("message not found")
	}

	return &message, nil
}

// List retrieves messages with filters and pagination
func (r *MessageRepository) List(ctx context.Context, senderID, receiverID *string, isRead *bool, limit, offset int) ([]*domain.Message, error) {
	query := `
		SELECT id, sender_id, receiver_id, subject, content, is_read, read_at, parent_message_id, created_at, deleted_at
		FROM messages
		WHERE deleted_at IS NULL
	`

	args := []interface{}{}
	argCount := 1

	if senderID != nil {
		query += fmt.Sprintf(" AND sender_id = $%d", argCount)
		args = append(args, *senderID)
		argCount++
	}

	if receiverID != nil {
		query += fmt.Sprintf(" AND receiver_id = $%d", argCount)
		args = append(args, *receiverID)
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
		return nil, fmt.Errorf("failed to list messages: %w", err)
	}
	defer rows.Close()

	var messages []*domain.Message
	for rows.Next() {
		var message domain.Message
		err := rows.Scan(
			&message.ID,
			&message.SenderID,
			&message.ReceiverID,
			&message.Subject,
			&message.Content,
			&message.IsRead,
			&message.ReadAt,
			&message.ParentMessageID,
			&message.CreatedAt,
			&message.DeletedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan message: %w", err)
		}
		messages = append(messages, &message)
	}

	return messages, nil
}

// Count returns the count of messages with filters
func (r *MessageRepository) Count(ctx context.Context, senderID, receiverID *string, isRead *bool) (int, error) {
	query := `
		SELECT COUNT(*)
		FROM messages
		WHERE deleted_at IS NULL
	`

	args := []interface{}{}
	argCount := 1

	if senderID != nil {
		query += fmt.Sprintf(" AND sender_id = $%d", argCount)
		args = append(args, *senderID)
		argCount++
	}

	if receiverID != nil {
		query += fmt.Sprintf(" AND receiver_id = $%d", argCount)
		args = append(args, *receiverID)
		argCount++
	}

	if isRead != nil {
		query += fmt.Sprintf(" AND is_read = $%d", argCount)
		args = append(args, *isRead)
	}

	var count int
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count messages: %w", err)
	}

	return count, nil
}

// Update updates a message
func (r *MessageRepository) Update(ctx context.Context, message *domain.Message) error {
	query := `
		UPDATE messages
		SET is_read = $1, read_at = $2
		WHERE id = $3 AND deleted_at IS NULL
	`

	result, err := r.db.ExecContext(ctx, query,
		message.IsRead,
		message.ReadAt,
		message.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update message: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("message not found")
	}

	return nil
}

// Delete soft deletes a message
func (r *MessageRepository) Delete(ctx context.Context, id string) error {
	query := `
		UPDATE messages
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete message: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("message not found")
	}

	return nil
}

// MarkAsRead marks a message as read
func (r *MessageRepository) MarkAsRead(ctx context.Context, id string) error {
	query := `
		UPDATE messages
		SET is_read = true, read_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to mark message as read: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("message not found")
	}

	return nil
}

// GetConversation retrieves messages between two users
func (r *MessageRepository) GetConversation(ctx context.Context, userID1, userID2 string, limit, offset int) ([]*domain.Message, error) {
	query := `
		SELECT id, sender_id, receiver_id, subject, content, is_read, read_at, parent_message_id, created_at, deleted_at
		FROM messages
		WHERE deleted_at IS NULL
		AND ((sender_id = $1 AND receiver_id = $2) OR (sender_id = $2 AND receiver_id = $1))
		ORDER BY created_at DESC
	`

	args := []interface{}{userID1, userID2}
	argCount := 3

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
		return nil, fmt.Errorf("failed to get conversation: %w", err)
	}
	defer rows.Close()

	var messages []*domain.Message
	for rows.Next() {
		var message domain.Message
		err := rows.Scan(
			&message.ID,
			&message.SenderID,
			&message.ReceiverID,
			&message.Subject,
			&message.Content,
			&message.IsRead,
			&message.ReadAt,
			&message.ParentMessageID,
			&message.CreatedAt,
			&message.DeletedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan message: %w", err)
		}
		messages = append(messages, &message)
	}

	return messages, nil
}

// GetUnreadCount returns the count of unread messages for a user
func (r *MessageRepository) GetUnreadCount(ctx context.Context, userID string) (int, error) {
	query := `
		SELECT COUNT(*)
		FROM messages
		WHERE receiver_id = $1 AND is_read = false AND deleted_at IS NULL
	`

	var count int
	err := r.db.QueryRowContext(ctx, query, userID).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to get unread count: %w", err)
	}

	return count, nil
}

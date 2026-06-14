package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/nusa/backend/internal/domain"
)

// AnnouncementRepository handles database operations for announcements
type AnnouncementRepository struct {
	db *sqlx.DB
}

// NewAnnouncementRepository creates a new announcement repository
func NewAnnouncementRepository(db *sqlx.DB) *AnnouncementRepository {
	return &AnnouncementRepository{db: db}
}

// Create creates a new announcement
func (r *AnnouncementRepository) Create(ctx context.Context, announcement *domain.Announcement) error {
	query := `
		INSERT INTO announcements (id, school_id, title, content, priority, target_audience, published_by, published_at, expires_at, is_active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`

	_, err := r.db.ExecContext(ctx, query,
		announcement.ID,
		announcement.SchoolID,
		announcement.Title,
		announcement.Content,
		announcement.Priority,
		announcement.TargetAudience,
		announcement.PublishedBy,
		announcement.PublishedAt,
		announcement.ExpiresAt,
		announcement.IsActive,
		announcement.CreatedAt,
		announcement.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to create announcement: %w", err)
	}

	return nil
}

// GetByID retrieves an announcement by ID
func (r *AnnouncementRepository) GetByID(ctx context.Context, id string) (*domain.Announcement, error) {
	query := `
		SELECT id, school_id, title, content, priority, target_audience, published_by, published_at, expires_at, is_active, created_at, updated_at, deleted_at
		FROM announcements
		WHERE id = $1 AND deleted_at IS NULL
	`

	var announcement domain.Announcement
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&announcement.ID,
		&announcement.SchoolID,
		&announcement.Title,
		&announcement.Content,
		&announcement.Priority,
		&announcement.TargetAudience,
		&announcement.PublishedBy,
		&announcement.PublishedAt,
		&announcement.ExpiresAt,
		&announcement.IsActive,
		&announcement.CreatedAt,
		&announcement.UpdatedAt,
		&announcement.DeletedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("announcement not found")
	}

	return &announcement, nil
}

// List retrieves announcements with filters and pagination
func (r *AnnouncementRepository) List(ctx context.Context, schoolID *string, priority *string, targetAudience *string, isActive *bool, limit, offset int) ([]*domain.Announcement, error) {
	query := `
		SELECT id, school_id, title, content, priority, target_audience, published_by, published_at, expires_at, is_active, created_at, updated_at, deleted_at
		FROM announcements
		WHERE deleted_at IS NULL
	`

	args := []interface{}{}
	argCount := 1

	if schoolID != nil {
		query += fmt.Sprintf(" AND school_id = $%d", argCount)
		args = append(args, *schoolID)
		argCount++
	}

	if priority != nil {
		query += fmt.Sprintf(" AND priority = $%d", argCount)
		args = append(args, *priority)
		argCount++
	}

	if targetAudience != nil {
		query += fmt.Sprintf(" AND target_audience = $%d", argCount)
		args = append(args, *targetAudience)
		argCount++
	}

	if isActive != nil {
		query += fmt.Sprintf(" AND is_active = $%d", argCount)
		args = append(args, *isActive)
		argCount++
	}

	query += " ORDER BY published_at DESC"

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
		return nil, fmt.Errorf("failed to list announcements: %w", err)
	}
	defer rows.Close()

	var announcements []*domain.Announcement
	for rows.Next() {
		var announcement domain.Announcement
		err := rows.Scan(
			&announcement.ID,
			&announcement.SchoolID,
			&announcement.Title,
			&announcement.Content,
			&announcement.Priority,
			&announcement.TargetAudience,
			&announcement.PublishedBy,
			&announcement.PublishedAt,
			&announcement.ExpiresAt,
			&announcement.IsActive,
			&announcement.CreatedAt,
			&announcement.UpdatedAt,
			&announcement.DeletedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan announcement: %w", err)
		}
		announcements = append(announcements, &announcement)
	}

	return announcements, nil
}

// Count returns the count of announcements with filters
func (r *AnnouncementRepository) Count(ctx context.Context, schoolID *string, priority *string, targetAudience *string, isActive *bool) (int, error) {
	query := `
		SELECT COUNT(*)
		FROM announcements
		WHERE deleted_at IS NULL
	`

	args := []interface{}{}
	argCount := 1

	if schoolID != nil {
		query += fmt.Sprintf(" AND school_id = $%d", argCount)
		args = append(args, *schoolID)
		argCount++
	}

	if priority != nil {
		query += fmt.Sprintf(" AND priority = $%d", argCount)
		args = append(args, *priority)
		argCount++
	}

	if targetAudience != nil {
		query += fmt.Sprintf(" AND target_audience = $%d", argCount)
		args = append(args, *targetAudience)
		argCount++
	}

	if isActive != nil {
		query += fmt.Sprintf(" AND is_active = $%d", argCount)
		args = append(args, *isActive)
		argCount++
	}

	var count int
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count announcements: %w", err)
	}

	return count, nil
}

// Update updates an announcement
func (r *AnnouncementRepository) Update(ctx context.Context, announcement *domain.Announcement) error {
	query := `
		UPDATE announcements
		SET title = $1, content = $2, priority = $3, target_audience = $4, expires_at = $5, is_active = $6, updated_at = $7
		WHERE id = $8 AND deleted_at IS NULL
	`

	announcement.UpdatedAt = time.Now()

	result, err := r.db.ExecContext(ctx, query,
		announcement.Title,
		announcement.Content,
		announcement.Priority,
		announcement.TargetAudience,
		announcement.ExpiresAt,
		announcement.IsActive,
		announcement.UpdatedAt,
		announcement.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update announcement: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("announcement not found")
	}

	return nil
}

// Delete soft deletes an announcement
func (r *AnnouncementRepository) Delete(ctx context.Context, id string) error {
	query := `
		UPDATE announcements
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete announcement: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("announcement not found")
	}

	return nil
}

// GetBySchoolID retrieves all announcements for a school
func (r *AnnouncementRepository) GetBySchoolID(ctx context.Context, schoolID string) ([]*domain.Announcement, error) {
	query := `
		SELECT id, school_id, title, content, priority, target_audience, published_by, published_at, expires_at, is_active, created_at, updated_at, deleted_at
		FROM announcements
		WHERE school_id = $1 AND deleted_at IS NULL
		ORDER BY published_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, schoolID)
	if err != nil {
		return nil, fmt.Errorf("failed to get announcements by school: %w", err)
	}
	defer rows.Close()

	var announcements []*domain.Announcement
	for rows.Next() {
		var announcement domain.Announcement
		err := rows.Scan(
			&announcement.ID,
			&announcement.SchoolID,
			&announcement.Title,
			&announcement.Content,
			&announcement.Priority,
			&announcement.TargetAudience,
			&announcement.PublishedBy,
			&announcement.PublishedAt,
			&announcement.ExpiresAt,
			&announcement.IsActive,
			&announcement.CreatedAt,
			&announcement.UpdatedAt,
			&announcement.DeletedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan announcement: %w", err)
		}
		announcements = append(announcements, &announcement)
	}

	return announcements, nil
}

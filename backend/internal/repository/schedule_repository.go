package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/nusa/backend/internal/domain"
)

// ScheduleRepository handles database operations for schedules
type ScheduleRepository struct {
	db *sqlx.DB
}

// NewScheduleRepository creates a new schedule repository
func NewScheduleRepository(db *sqlx.DB) *ScheduleRepository {
	return &ScheduleRepository{db: db}
}

// Create creates a new schedule
func (r *ScheduleRepository) Create(ctx context.Context, schedule *domain.Schedule) error {
	query := `
		INSERT INTO schedules (id, class_id, day_of_week, start_time, end_time, room, is_active, created_at, updated_at, created_by, updated_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`

	_, err := r.db.ExecContext(ctx, query,
		schedule.ID,
		schedule.ClassID,
		schedule.DayOfWeek,
		schedule.StartTime,
		schedule.EndTime,
		schedule.Room,
		schedule.IsActive,
		schedule.CreatedAt,
		schedule.UpdatedAt,
		schedule.CreatedBy,
		schedule.UpdatedBy,
	)

	if err != nil {
		return fmt.Errorf("failed to create schedule: %w", err)
	}

	return nil
}

// GetByID retrieves a schedule by ID
func (r *ScheduleRepository) GetByID(ctx context.Context, id string) (*domain.Schedule, error) {
	query := `
		SELECT id, class_id, day_of_week, start_time, end_time, room, is_active, created_at, updated_at, created_by, updated_by, deleted_at
		FROM schedules
		WHERE id = $1 AND deleted_at IS NULL
	`

	var schedule domain.Schedule
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&schedule.ID,
		&schedule.ClassID,
		&schedule.DayOfWeek,
		&schedule.StartTime,
		&schedule.EndTime,
		&schedule.Room,
		&schedule.IsActive,
		&schedule.CreatedAt,
		&schedule.UpdatedAt,
		&schedule.CreatedBy,
		&schedule.UpdatedBy,
		&schedule.DeletedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("schedule not found")
	}

	return &schedule, nil
}

// List retrieves schedules with filters and pagination
func (r *ScheduleRepository) List(ctx context.Context, classID *string, dayOfWeek *int, isActive *bool, limit, offset int) ([]*domain.Schedule, error) {
	query := `
		SELECT id, class_id, day_of_week, start_time, end_time, room, is_active, created_at, updated_at, created_by, updated_by, deleted_at
		FROM schedules
		WHERE deleted_at IS NULL
	`

	args := []interface{}{}
	argCount := 1

	if classID != nil {
		query += fmt.Sprintf(" AND class_id = $%d", argCount)
		args = append(args, *classID)
		argCount++
	}

	if dayOfWeek != nil {
		query += fmt.Sprintf(" AND day_of_week = $%d", argCount)
		args = append(args, *dayOfWeek)
		argCount++
	}

	if isActive != nil {
		query += fmt.Sprintf(" AND is_active = $%d", argCount)
		args = append(args, *isActive)
		argCount++
	}

	query += " ORDER BY day_of_week, start_time"

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
		return nil, fmt.Errorf("failed to list schedules: %w", err)
	}
	defer rows.Close()

	var schedules []*domain.Schedule
	for rows.Next() {
		var schedule domain.Schedule
		err := rows.Scan(
			&schedule.ID,
			&schedule.ClassID,
			&schedule.DayOfWeek,
			&schedule.StartTime,
			&schedule.EndTime,
			&schedule.Room,
			&schedule.IsActive,
			&schedule.CreatedAt,
			&schedule.UpdatedAt,
			&schedule.CreatedBy,
			&schedule.UpdatedBy,
			&schedule.DeletedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan schedule: %w", err)
		}
		schedules = append(schedules, &schedule)
	}

	return schedules, nil
}

// Count returns the count of schedules with filters
func (r *ScheduleRepository) Count(ctx context.Context, classID *string, dayOfWeek *int, isActive *bool) (int, error) {
	query := `
		SELECT COUNT(*)
		FROM schedules
		WHERE deleted_at IS NULL
	`

	args := []interface{}{}
	argCount := 1

	if classID != nil {
		query += fmt.Sprintf(" AND class_id = $%d", argCount)
		args = append(args, *classID)
		argCount++
	}

	if dayOfWeek != nil {
		query += fmt.Sprintf(" AND day_of_week = $%d", argCount)
		args = append(args, *dayOfWeek)
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
		return 0, fmt.Errorf("failed to count schedules: %w", err)
	}

	return count, nil
}

// Update updates a schedule
func (r *ScheduleRepository) Update(ctx context.Context, schedule *domain.Schedule) error {
	query := `
		UPDATE schedules
		SET day_of_week = $1, start_time = $2, end_time = $3, room = $4, is_active = $5, updated_at = $6, updated_by = $7
		WHERE id = $8 AND deleted_at IS NULL
	`

	schedule.UpdatedAt = time.Now()

	result, err := r.db.ExecContext(ctx, query,
		schedule.DayOfWeek,
		schedule.StartTime,
		schedule.EndTime,
		schedule.Room,
		schedule.IsActive,
		schedule.UpdatedAt,
		schedule.UpdatedBy,
		schedule.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update schedule: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("schedule not found")
	}

	return nil
}

// Delete soft deletes a schedule
func (r *ScheduleRepository) Delete(ctx context.Context, id string) error {
	query := `
		UPDATE schedules
		SET deleted_at = $1
		WHERE id = $2 AND deleted_at IS NULL
	`

	now := time.Now()
	result, err := r.db.ExecContext(ctx, query, now, id)
	if err != nil {
		return fmt.Errorf("failed to delete schedule: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("schedule not found")
	}

	return nil
}

// GetByClassID retrieves all schedules for a class
func (r *ScheduleRepository) GetByClassID(ctx context.Context, classID string) ([]*domain.Schedule, error) {
	query := `
		SELECT id, class_id, day_of_week, start_time, end_time, room, is_active, created_at, updated_at, created_by, updated_by, deleted_at
		FROM schedules
		WHERE class_id = $1 AND deleted_at IS NULL
		ORDER BY day_of_week, start_time
	`

	rows, err := r.db.QueryContext(ctx, query, classID)
	if err != nil {
		return nil, fmt.Errorf("failed to get schedules by class: %w", err)
	}
	defer rows.Close()

	var schedules []*domain.Schedule
	for rows.Next() {
		var schedule domain.Schedule
		err := rows.Scan(
			&schedule.ID,
			&schedule.ClassID,
			&schedule.DayOfWeek,
			&schedule.StartTime,
			&schedule.EndTime,
			&schedule.Room,
			&schedule.IsActive,
			&schedule.CreatedAt,
			&schedule.UpdatedAt,
			&schedule.CreatedBy,
			&schedule.UpdatedBy,
			&schedule.DeletedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan schedule: %w", err)
		}
		schedules = append(schedules, &schedule)
	}

	return schedules, nil
}

package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/nusa/backend/internal/domain"
)

// ExamRepository handles database operations for exams
type ExamRepository struct {
	db *sqlx.DB
}

// NewExamRepository creates a new exam repository
func NewExamRepository(db *sqlx.DB) *ExamRepository {
	return &ExamRepository{db: db}
}

// Create creates a new exam
func (r *ExamRepository) Create(ctx context.Context, exam *domain.Exam) error {
	query := `
		INSERT INTO exams (id, class_id, assessment_id, exam_date, start_time, duration_minutes, room, status, created_at, updated_at, created_by, updated_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`

	_, err := r.db.ExecContext(ctx, query,
		exam.ID,
		exam.ClassID,
		exam.AssessmentID,
		exam.ExamDate,
		exam.StartTime,
		exam.DurationMinutes,
		exam.Room,
		exam.Status,
		exam.CreatedAt,
		exam.UpdatedAt,
		exam.CreatedBy,
		exam.UpdatedBy,
	)

	if err != nil {
		return fmt.Errorf("failed to create exam: %w", err)
	}

	return nil
}

// GetByID retrieves an exam by ID
func (r *ExamRepository) GetByID(ctx context.Context, id string) (*domain.Exam, error) {
	query := `
		SELECT id, class_id, assessment_id, exam_date, start_time, duration_minutes, room, status, created_at, updated_at, created_by, updated_by, deleted_at
		FROM exams
		WHERE id = $1 AND deleted_at IS NULL
	`

	var exam domain.Exam
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&exam.ID,
		&exam.ClassID,
		&exam.AssessmentID,
		&exam.ExamDate,
		&exam.StartTime,
		&exam.DurationMinutes,
		&exam.Room,
		&exam.Status,
		&exam.CreatedAt,
		&exam.UpdatedAt,
		&exam.CreatedBy,
		&exam.UpdatedBy,
		&exam.DeletedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("exam not found")
	}

	return &exam, nil
}

// List retrieves exams with filters and pagination
func (r *ExamRepository) List(ctx context.Context, classID, assessmentID, status *string, limit, offset int) ([]*domain.Exam, error) {
	query := `
		SELECT id, class_id, assessment_id, exam_date, start_time, duration_minutes, room, status, created_at, updated_at, created_by, updated_by, deleted_at
		FROM exams
		WHERE deleted_at IS NULL
	`

	args := []interface{}{}
	argCount := 1

	if classID != nil {
		query += fmt.Sprintf(" AND class_id = $%d", argCount)
		args = append(args, *classID)
		argCount++
	}

	if assessmentID != nil {
		query += fmt.Sprintf(" AND assessment_id = $%d", argCount)
		args = append(args, *assessmentID)
		argCount++
	}

	if status != nil {
		query += fmt.Sprintf(" AND status = $%d", argCount)
		args = append(args, *status)
		argCount++
	}

	query += " ORDER BY exam_date DESC"

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
		return nil, fmt.Errorf("failed to list exams: %w", err)
	}
	defer rows.Close()

	var exams []*domain.Exam
	for rows.Next() {
		var exam domain.Exam
		err := rows.Scan(
			&exam.ID,
			&exam.ClassID,
			&exam.AssessmentID,
			&exam.ExamDate,
			&exam.StartTime,
			&exam.DurationMinutes,
			&exam.Room,
			&exam.Status,
			&exam.CreatedAt,
			&exam.UpdatedAt,
			&exam.CreatedBy,
			&exam.UpdatedBy,
			&exam.DeletedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan exam: %w", err)
		}
		exams = append(exams, &exam)
	}

	return exams, nil
}

// Count returns the count of exams with filters
func (r *ExamRepository) Count(ctx context.Context, classID, assessmentID, status *string) (int, error) {
	query := `
		SELECT COUNT(*)
		FROM exams
		WHERE deleted_at IS NULL
	`

	args := []interface{}{}
	argCount := 1

	if classID != nil {
		query += fmt.Sprintf(" AND class_id = $%d", argCount)
		args = append(args, *classID)
		argCount++
	}

	if assessmentID != nil {
		query += fmt.Sprintf(" AND assessment_id = $%d", argCount)
		args = append(args, *assessmentID)
		argCount++
	}

	if status != nil {
		query += fmt.Sprintf(" AND status = $%d", argCount)
		args = append(args, *status)
		argCount++
	}

	var count int
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count exams: %w", err)
	}

	return count, nil
}

// Update updates an exam
func (r *ExamRepository) Update(ctx context.Context, exam *domain.Exam) error {
	query := `
		UPDATE exams
		SET exam_date = $1, start_time = $2, duration_minutes = $3, room = $4, status = $5, updated_at = $6, updated_by = $7
		WHERE id = $8 AND deleted_at IS NULL
	`

	exam.UpdatedAt = time.Now()

	result, err := r.db.ExecContext(ctx, query,
		exam.ExamDate,
		exam.StartTime,
		exam.DurationMinutes,
		exam.Room,
		exam.Status,
		exam.UpdatedAt,
		exam.UpdatedBy,
		exam.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update exam: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("exam not found")
	}

	return nil
}

// Delete soft deletes an exam
func (r *ExamRepository) Delete(ctx context.Context, id string) error {
	query := `
		UPDATE exams
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete exam: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("exam not found")
	}

	return nil
}

// GetByClassID retrieves all exams for a class
func (r *ExamRepository) GetByClassID(ctx context.Context, classID string) ([]*domain.Exam, error) {
	query := `
		SELECT id, class_id, assessment_id, exam_date, start_time, duration_minutes, room, status, created_at, updated_at, created_by, updated_by, deleted_at
		FROM exams
		WHERE class_id = $1 AND deleted_at IS NULL
		ORDER BY exam_date DESC
	`

	rows, err := r.db.QueryContext(ctx, query, classID)
	if err != nil {
		return nil, fmt.Errorf("failed to get exams by class: %w", err)
	}
	defer rows.Close()

	var exams []*domain.Exam
	for rows.Next() {
		var exam domain.Exam
		err := rows.Scan(
			&exam.ID,
			&exam.ClassID,
			&exam.AssessmentID,
			&exam.ExamDate,
			&exam.StartTime,
			&exam.DurationMinutes,
			&exam.Room,
			&exam.Status,
			&exam.CreatedAt,
			&exam.UpdatedAt,
			&exam.CreatedBy,
			&exam.UpdatedBy,
			&exam.DeletedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan exam: %w", err)
		}
		exams = append(exams, &exam)
	}

	return exams, nil
}

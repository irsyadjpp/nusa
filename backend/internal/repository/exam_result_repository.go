package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/nusa/backend/internal/domain"
)

// ExamResultRepository handles database operations for exam results
type ExamResultRepository struct {
	db *sqlx.DB
}

// NewExamResultRepository creates a new exam result repository
func NewExamResultRepository(db *sqlx.DB) *ExamResultRepository {
	return &ExamResultRepository{db: db}
}

// Create creates a new exam result
func (r *ExamResultRepository) Create(ctx context.Context, examResult *domain.ExamResult) error {
	query := `
		INSERT INTO exam_results (id, exam_id, student_id, score, grade, remarks, graded_at, graded_by, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		ON CONFLICT (exam_id, student_id) DO UPDATE
		SET score = EXCLUDED.score, grade = EXCLUDED.grade, remarks = EXCLUDED.remarks, graded_at = EXCLUDED.graded_at, graded_by = EXCLUDED.graded_by, updated_at = EXCLUDED.updated_at
	`

	_, err := r.db.ExecContext(ctx, query,
		examResult.ID,
		examResult.ExamID,
		examResult.StudentID,
		examResult.Score,
		examResult.Grade,
		examResult.Remarks,
		examResult.GradedAt,
		examResult.GradedBy,
		examResult.CreatedAt,
		examResult.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to create exam result: %w", err)
	}

	return nil
}

// GetByID retrieves an exam result by ID
func (r *ExamResultRepository) GetByID(ctx context.Context, id string) (*domain.ExamResult, error) {
	query := `
		SELECT id, exam_id, student_id, score, grade, remarks, graded_at, graded_by, created_at, updated_at, deleted_at
		FROM exam_results
		WHERE id = $1 AND deleted_at IS NULL
	`

	var examResult domain.ExamResult
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&examResult.ID,
		&examResult.ExamID,
		&examResult.StudentID,
		&examResult.Score,
		&examResult.Grade,
		&examResult.Remarks,
		&examResult.GradedAt,
		&examResult.GradedBy,
		&examResult.CreatedAt,
		&examResult.UpdatedAt,
		&examResult.DeletedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("exam result not found")
	}

	return &examResult, nil
}

// GetByExamAndStudent retrieves an exam result by exam ID and student ID
func (r *ExamResultRepository) GetByExamAndStudent(ctx context.Context, examID, studentID string) (*domain.ExamResult, error) {
	query := `
		SELECT id, exam_id, student_id, score, grade, remarks, graded_at, graded_by, created_at, updated_at, deleted_at
		FROM exam_results
		WHERE exam_id = $1 AND student_id = $2 AND deleted_at IS NULL
	`

	var examResult domain.ExamResult
	err := r.db.QueryRowContext(ctx, query, examID, studentID).Scan(
		&examResult.ID,
		&examResult.ExamID,
		&examResult.StudentID,
		&examResult.Score,
		&examResult.Grade,
		&examResult.Remarks,
		&examResult.GradedAt,
		&examResult.GradedBy,
		&examResult.CreatedAt,
		&examResult.UpdatedAt,
		&examResult.DeletedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("exam result not found")
	}

	return &examResult, nil
}

// List retrieves exam results with filters and pagination
func (r *ExamResultRepository) List(ctx context.Context, examID, studentID, grade *string, limit, offset int) ([]*domain.ExamResult, error) {
	query := `
		SELECT id, exam_id, student_id, score, grade, remarks, graded_at, graded_by, created_at, updated_at, deleted_at
		FROM exam_results
		WHERE deleted_at IS NULL
	`

	args := []interface{}{}
	argCount := 1

	if examID != nil {
		query += fmt.Sprintf(" AND exam_id = $%d", argCount)
		args = append(args, *examID)
		argCount++
	}

	if studentID != nil {
		query += fmt.Sprintf(" AND student_id = $%d", argCount)
		args = append(args, *studentID)
		argCount++
	}

	if grade != nil {
		query += fmt.Sprintf(" AND grade = $%d", argCount)
		args = append(args, *grade)
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
		return nil, fmt.Errorf("failed to list exam results: %w", err)
	}
	defer rows.Close()

	var examResults []*domain.ExamResult
	for rows.Next() {
		var examResult domain.ExamResult
		err := rows.Scan(
			&examResult.ID,
			&examResult.ExamID,
			&examResult.StudentID,
			&examResult.Score,
			&examResult.Grade,
			&examResult.Remarks,
			&examResult.GradedAt,
			&examResult.GradedBy,
			&examResult.CreatedAt,
			&examResult.UpdatedAt,
			&examResult.DeletedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan exam result: %w", err)
		}
		examResults = append(examResults, &examResult)
	}

	return examResults, nil
}

// Count returns the count of exam results with filters
func (r *ExamResultRepository) Count(ctx context.Context, examID, studentID, grade *string) (int, error) {
	query := `
		SELECT COUNT(*)
		FROM exam_results
		WHERE deleted_at IS NULL
	`

	args := []interface{}{}
	argCount := 1

	if examID != nil {
		query += fmt.Sprintf(" AND exam_id = $%d", argCount)
		args = append(args, *examID)
		argCount++
	}

	if studentID != nil {
		query += fmt.Sprintf(" AND student_id = $%d", argCount)
		args = append(args, *studentID)
		argCount++
	}

	if grade != nil {
		query += fmt.Sprintf(" AND grade = $%d", argCount)
		args = append(args, *grade)
		argCount++
	}

	var count int
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count exam results: %w", err)
	}

	return count, nil
}

// Update updates an exam result
func (r *ExamResultRepository) Update(ctx context.Context, examResult *domain.ExamResult) error {
	query := `
		UPDATE exam_results
		SET score = $1, grade = $2, remarks = $3, graded_at = $4, graded_by = $5, updated_at = $6
		WHERE id = $7 AND deleted_at IS NULL
	`

	examResult.UpdatedAt = time.Now()

	result, err := r.db.ExecContext(ctx, query,
		examResult.Score,
		examResult.Grade,
		examResult.Remarks,
		examResult.GradedAt,
		examResult.GradedBy,
		examResult.UpdatedAt,
		examResult.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update exam result: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("exam result not found")
	}

	return nil
}

// Delete soft deletes an exam result
func (r *ExamResultRepository) Delete(ctx context.Context, id string) error {
	query := `
		UPDATE exam_results
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete exam result: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("exam result not found")
	}

	return nil
}

// GetByExamID retrieves all exam results for an exam
func (r *ExamResultRepository) GetByExamID(ctx context.Context, examID string) ([]*domain.ExamResult, error) {
	query := `
		SELECT id, exam_id, student_id, score, grade, remarks, graded_at, graded_by, created_at, updated_at, deleted_at
		FROM exam_results
		WHERE exam_id = $1 AND deleted_at IS NULL
		ORDER BY score DESC
	`

	rows, err := r.db.QueryContext(ctx, query, examID)
	if err != nil {
		return nil, fmt.Errorf("failed to get exam results by exam: %w", err)
	}
	defer rows.Close()

	var examResults []*domain.ExamResult
	for rows.Next() {
		var examResult domain.ExamResult
		err := rows.Scan(
			&examResult.ID,
			&examResult.ExamID,
			&examResult.StudentID,
			&examResult.Score,
			&examResult.Grade,
			&examResult.Remarks,
			&examResult.GradedAt,
			&examResult.GradedBy,
			&examResult.CreatedAt,
			&examResult.UpdatedAt,
			&examResult.DeletedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan exam result: %w", err)
		}
		examResults = append(examResults, &examResult)
	}

	return examResults, nil
}

// GetByStudentID retrieves all exam results for a student
func (r *ExamResultRepository) GetByStudentID(ctx context.Context, studentID string) ([]*domain.ExamResult, error) {
	query := `
		SELECT id, exam_id, student_id, score, grade, remarks, graded_at, graded_by, created_at, updated_at, deleted_at
		FROM exam_results
		WHERE student_id = $1 AND deleted_at IS NULL
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, studentID)
	if err != nil {
		return nil, fmt.Errorf("failed to get exam results by student: %w", err)
	}
	defer rows.Close()

	var examResults []*domain.ExamResult
	for rows.Next() {
		var examResult domain.ExamResult
		err := rows.Scan(
			&examResult.ID,
			&examResult.ExamID,
			&examResult.StudentID,
			&examResult.Score,
			&examResult.Grade,
			&examResult.Remarks,
			&examResult.GradedAt,
			&examResult.GradedBy,
			&examResult.CreatedAt,
			&examResult.UpdatedAt,
			&examResult.DeletedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan exam result: %w", err)
		}
		examResults = append(examResults, &examResult)
	}

	return examResults, nil
}

package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/nusa/backend/internal/domain"
)

// AssignmentRepository handles database operations for assignments
type AssignmentRepository struct {
	db *sqlx.DB
}

// NewAssignmentRepository creates a new assignment repository
func NewAssignmentRepository(db *sqlx.DB) *AssignmentRepository {
	return &AssignmentRepository{db: db}
}

// Create creates a new assignment
func (r *AssignmentRepository) Create(ctx context.Context, assignment *domain.Assignment) error {
	query := `
		INSERT INTO assignments (id, class_id, assessment_id, title, description, due_date, max_score, status, created_at, updated_at, created_by, updated_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`

	_, err := r.db.ExecContext(ctx, query,
		assignment.ID,
		assignment.ClassID,
		assignment.AssessmentID,
		assignment.Title,
		assignment.Description,
		assignment.DueDate,
		assignment.MaxScore,
		assignment.Status,
		assignment.CreatedAt,
		assignment.UpdatedAt,
		assignment.CreatedBy,
		assignment.UpdatedBy,
	)

	if err != nil {
		return fmt.Errorf("failed to create assignment: %w", err)
	}

	return nil
}

// GetByID retrieves an assignment by ID
func (r *AssignmentRepository) GetByID(ctx context.Context, id string) (*domain.Assignment, error) {
	query := `
		SELECT id, class_id, assessment_id, title, description, due_date, max_score, status, created_at, updated_at, created_by, updated_by, deleted_at
		FROM assignments
		WHERE id = $1 AND deleted_at IS NULL
	`

	var assignment domain.Assignment
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&assignment.ID,
		&assignment.ClassID,
		&assignment.AssessmentID,
		&assignment.Title,
		&assignment.Description,
		&assignment.DueDate,
		&assignment.MaxScore,
		&assignment.Status,
		&assignment.CreatedAt,
		&assignment.UpdatedAt,
		&assignment.CreatedBy,
		&assignment.UpdatedBy,
		&assignment.DeletedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("assignment not found")
	}

	return &assignment, nil
}

// List retrieves assignments with filters and pagination
func (r *AssignmentRepository) List(ctx context.Context, classID, assessmentID, status *string, limit, offset int) ([]*domain.Assignment, error) {
	query := `
		SELECT id, class_id, assessment_id, title, description, due_date, max_score, status, created_at, updated_at, created_by, updated_by, deleted_at
		FROM assignments
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

	query += " ORDER BY due_date DESC"

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
		return nil, fmt.Errorf("failed to list assignments: %w", err)
	}
	defer rows.Close()

	var assignments []*domain.Assignment
	for rows.Next() {
		var assignment domain.Assignment
		err := rows.Scan(
			&assignment.ID,
			&assignment.ClassID,
			&assignment.AssessmentID,
			&assignment.Title,
			&assignment.Description,
			&assignment.DueDate,
			&assignment.MaxScore,
			&assignment.Status,
			&assignment.CreatedAt,
			&assignment.UpdatedAt,
			&assignment.CreatedBy,
			&assignment.UpdatedBy,
			&assignment.DeletedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan assignment: %w", err)
		}
		assignments = append(assignments, &assignment)
	}

	return assignments, nil
}

// Count returns the count of assignments with filters
func (r *AssignmentRepository) Count(ctx context.Context, classID, assessmentID, status *string) (int, error) {
	query := `
		SELECT COUNT(*)
		FROM assignments
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
	}

	var count int
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count assignments: %w", err)
	}

	return count, nil
}

// Update updates an assignment
func (r *AssignmentRepository) Update(ctx context.Context, assignment *domain.Assignment) error {
	query := `
		UPDATE assignments
		SET title = $1, description = $2, due_date = $3, max_score = $4, status = $5, updated_at = $6, updated_by = $7
		WHERE id = $8 AND deleted_at IS NULL
	`

	assignment.UpdatedAt = time.Now()

	result, err := r.db.ExecContext(ctx, query,
		assignment.Title,
		assignment.Description,
		assignment.DueDate,
		assignment.MaxScore,
		assignment.Status,
		assignment.UpdatedAt,
		assignment.UpdatedBy,
		assignment.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update assignment: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("assignment not found")
	}

	return nil
}

// Delete soft deletes an assignment
func (r *AssignmentRepository) Delete(ctx context.Context, id string) error {
	query := `
		UPDATE assignments
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete assignment: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("assignment not found")
	}

	return nil
}

// GetByClassID retrieves all assignments for a class
func (r *AssignmentRepository) GetByClassID(ctx context.Context, classID string) ([]*domain.Assignment, error) {
	query := `
		SELECT id, class_id, assessment_id, title, description, due_date, max_score, status, created_at, updated_at, created_by, updated_by, deleted_at
		FROM assignments
		WHERE class_id = $1 AND deleted_at IS NULL
		ORDER BY due_date DESC
	`

	rows, err := r.db.QueryContext(ctx, query, classID)
	if err != nil {
		return nil, fmt.Errorf("failed to get assignments by class: %w", err)
	}
	defer rows.Close()

	var assignments []*domain.Assignment
	for rows.Next() {
		var assignment domain.Assignment
		err := rows.Scan(
			&assignment.ID,
			&assignment.ClassID,
			&assignment.AssessmentID,
			&assignment.Title,
			&assignment.Description,
			&assignment.DueDate,
			&assignment.MaxScore,
			&assignment.Status,
			&assignment.CreatedAt,
			&assignment.UpdatedAt,
			&assignment.CreatedBy,
			&assignment.UpdatedBy,
			&assignment.DeletedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan assignment: %w", err)
		}
		assignments = append(assignments, &assignment)
	}

	return assignments, nil
}

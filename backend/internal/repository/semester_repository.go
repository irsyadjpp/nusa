package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/nusa/backend/internal/domain"
)

// SemesterRepository handles database operations for semesters
type SemesterRepository struct {
	db *sqlx.DB
}

// NewSemesterRepository creates a new semester repository
func NewSemesterRepository(db *sqlx.DB) *SemesterRepository {
	return &SemesterRepository{db: db}
}

// CreateSemester creates a new semester
func (r *SemesterRepository) CreateSemester(ctx context.Context, sem *domain.Semester) error {
	query := `
		INSERT INTO semesters (id, academic_year_id, type, name, start_date, end_date, status, sequence_number, created_by, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`

	_, err := r.db.ExecContext(ctx, query,
		sem.ID, sem.AcademicYearID, sem.Type, sem.Name, sem.StartDate, sem.EndDate,
		sem.Status, sem.SequenceNumber, sem.CreatedBy, sem.CreatedAt, sem.UpdatedAt)
	return err
}

// GetSemesterByID retrieves a semester by ID
func (r *SemesterRepository) GetSemesterByID(ctx context.Context, id string) (*domain.Semester, error) {
	query := `
		SELECT id, academic_year_id, type, name, start_date, end_date, status, sequence_number, created_by, created_at, updated_at
		FROM semesters WHERE id = $1
	`

	var sem domain.Semester
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&sem.ID, &sem.AcademicYearID, &sem.Type, &sem.Name, &sem.StartDate, &sem.EndDate,
		&sem.Status, &sem.SequenceNumber, &sem.CreatedBy, &sem.CreatedAt, &sem.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &sem, nil
}

// GetSemestersByAcademicYearID retrieves all semesters for an academic year
func (r *SemesterRepository) GetSemestersByAcademicYearID(ctx context.Context, academicYearID string) ([]*domain.Semester, error) {
	query := `
		SELECT id, academic_year_id, type, name, start_date, end_date, status, sequence_number, created_by, created_at, updated_at
		FROM semesters WHERE academic_year_id = $1 ORDER BY sequence_number ASC
	`

	rows, err := r.db.QueryContext(ctx, query, academicYearID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var semesters []*domain.Semester
	for rows.Next() {
		var sem domain.Semester
		err := rows.Scan(
			&sem.ID, &sem.AcademicYearID, &sem.Type, &sem.Name, &sem.StartDate, &sem.EndDate,
			&sem.Status, &sem.SequenceNumber, &sem.CreatedBy, &sem.CreatedAt, &sem.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		semesters = append(semesters, &sem)
	}
	return semesters, nil
}

// UpdateSemester updates a semester
func (r *SemesterRepository) UpdateSemester(ctx context.Context, sem *domain.Semester) error {
	query := `
		UPDATE semesters
		SET name = $1, start_date = $2, end_date = $3, status = $4, updated_at = $5
		WHERE id = $6
	`

	result, err := r.db.ExecContext(ctx, query,
		sem.Name, sem.StartDate, sem.EndDate, sem.Status, sem.UpdatedAt, sem.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("semester not found")
	}
	return nil
}

// DeleteSemester deletes a semester
func (r *SemesterRepository) DeleteSemester(ctx context.Context, id string) error {
	query := `DELETE FROM semesters WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("semester not found")
	}
	return nil
}

// CheckSemesterOverlap checks if a semester overlaps with existing ones in the same academic year
func (r *SemesterRepository) CheckSemesterOverlap(ctx context.Context, academicYearID string, startDate, endDate interface{}, excludeID string) (bool, error) {
	query := `
		SELECT COUNT(*) FROM semesters
		WHERE academic_year_id = $1
		AND (
			(start_date <= $2 AND end_date >= $3)
			OR (start_date <= $3 AND end_date >= $3)
			OR (start_date >= $2 AND end_date <= $3)
		)
	`
	args := []interface{}{academicYearID, startDate, endDate}

	if excludeID != "" {
		query += ` AND id != $4`
		args = append(args, excludeID)
	}

	var count int
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// CountSemestersByAcademicYearID counts the number of semesters for an academic year
func (r *SemesterRepository) CountSemestersByAcademicYearID(ctx context.Context, academicYearID string) (int, error) {
	query := `SELECT COUNT(*) FROM semesters WHERE academic_year_id = $1`
	var count int
	err := r.db.QueryRowContext(ctx, query, academicYearID).Scan(&count)
	return count, err
}

// CheckSequenceNumberExists checks if a sequence number already exists for an academic year
func (r *SemesterRepository) CheckSequenceNumberExists(ctx context.Context, academicYearID string, sequenceNumber int, excludeID string) (bool, error) {
	query := `
		SELECT COUNT(*) FROM semesters
		WHERE academic_year_id = $1 AND sequence_number = $2
	`
	args := []interface{}{academicYearID, sequenceNumber}

	if excludeID != "" {
		query += ` AND id != $3`
		args = append(args, excludeID)
	}

	var count int
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/nusa/backend/internal/domain"
)

// AcademicYearRepository handles database operations for academic years
type AcademicYearRepository struct {
	db *sqlx.DB
}

// NewAcademicYearRepository creates a new academic year repository
func NewAcademicYearRepository(db *sqlx.DB) *AcademicYearRepository {
	return &AcademicYearRepository{db: db}
}

// CreateAcademicYear creates a new academic year
func (r *AcademicYearRepository) CreateAcademicYear(ctx context.Context, ay *domain.AcademicYear) error {
	query := `
		INSERT INTO academic_years (id, school_id, name, start_date, end_date, status, created_by, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	_, err := r.db.ExecContext(ctx, query,
		ay.ID, ay.SchoolID, ay.Name, ay.StartDate, ay.EndDate, ay.Status,
		ay.CreatedBy, ay.CreatedAt, ay.UpdatedAt)
	return err
}

// GetAcademicYearByID retrieves an academic year by ID
func (r *AcademicYearRepository) GetAcademicYearByID(ctx context.Context, id string) (*domain.AcademicYear, error) {
	query := `
		SELECT id, school_id, name, start_date, end_date, status, created_by, created_at, updated_at
		FROM academic_years WHERE id = $1
	`

	var ay domain.AcademicYear
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&ay.ID, &ay.SchoolID, &ay.Name, &ay.StartDate, &ay.EndDate, &ay.Status,
		&ay.CreatedBy, &ay.CreatedAt, &ay.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &ay, nil
}

// GetAcademicYearsBySchoolID retrieves all academic years for a school
func (r *AcademicYearRepository) GetAcademicYearsBySchoolID(ctx context.Context, schoolID string) ([]*domain.AcademicYear, error) {
	query := `
		SELECT id, school_id, name, start_date, end_date, status, created_by, created_at, updated_at
		FROM academic_years WHERE school_id = $1 ORDER BY start_date DESC
	`

	rows, err := r.db.QueryContext(ctx, query, schoolID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var years []*domain.AcademicYear
	for rows.Next() {
		var ay domain.AcademicYear
		err := rows.Scan(
			&ay.ID, &ay.SchoolID, &ay.Name, &ay.StartDate, &ay.EndDate, &ay.Status,
			&ay.CreatedBy, &ay.CreatedAt, &ay.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		years = append(years, &ay)
	}
	return years, nil
}

// GetActiveAcademicYearBySchoolID retrieves the active academic year for a school
func (r *AcademicYearRepository) GetActiveAcademicYearBySchoolID(ctx context.Context, schoolID string) (*domain.AcademicYear, error) {
	query := `
		SELECT id, school_id, name, start_date, end_date, status, created_by, created_at, updated_at
		FROM academic_years WHERE school_id = $1 AND status = $2
	`

	var ay domain.AcademicYear
	err := r.db.QueryRowContext(ctx, query, schoolID, domain.AcademicYearStatusActive).Scan(
		&ay.ID, &ay.SchoolID, &ay.Name, &ay.StartDate, &ay.EndDate, &ay.Status,
		&ay.CreatedBy, &ay.CreatedAt, &ay.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &ay, nil
}

// UpdateAcademicYear updates an academic year
func (r *AcademicYearRepository) UpdateAcademicYear(ctx context.Context, ay *domain.AcademicYear) error {
	query := `
		UPDATE academic_years
		SET name = $1, start_date = $2, end_date = $3, status = $4, updated_at = $5
		WHERE id = $6
	`

	result, err := r.db.ExecContext(ctx, query,
		ay.Name, ay.StartDate, ay.EndDate, ay.Status, ay.UpdatedAt, ay.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("academic year not found")
	}
	return nil
}

// DeleteAcademicYear deletes an academic year
func (r *AcademicYearRepository) DeleteAcademicYear(ctx context.Context, id string) error {
	query := `DELETE FROM academic_years WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("academic year not found")
	}
	return nil
}

// CheckAcademicYearOverlap checks if an academic year overlaps with existing ones
// Used for enforcing BR-002: Non-overlapping academic years
func (r *AcademicYearRepository) CheckAcademicYearOverlap(ctx context.Context, schoolID string, startDate, endDate time.Time, excludeID string) (bool, error) {
	query := `
		SELECT COUNT(*) FROM academic_years
		WHERE school_id = $1
		AND status IN ($2, $3)
		AND (
			(start_date <= $4 AND end_date >= $5)
			OR (start_date <= $5 AND end_date >= $5)
			OR (start_date >= $4 AND end_date <= $5)
		)
	`
	args := []interface{}{schoolID, domain.AcademicYearStatusDraft, domain.AcademicYearStatusActive, startDate, endDate}

	if excludeID != "" {
		query += ` AND id != $6`
		args = append(args, excludeID)
	}

	var count int
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// ActivateAcademicYear activates an academic year and deactivates any previously active one
// Used for enforcing BR-001: Only one active academic year at a time
func (r *AcademicYearRepository) ActivateAcademicYear(ctx context.Context, id string) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Deactivate all active academic years in the same school
	deactivateQuery := `
		UPDATE academic_years
		SET status = $1, updated_at = $2
		WHERE id IN (
			SELECT id FROM academic_years WHERE school_id = (
				SELECT school_id FROM academic_years WHERE id = $3
			) AND status = $4
		)
	`
	_, err = tx.ExecContext(ctx, deactivateQuery, domain.AcademicYearStatusDraft, time.Now(), id, domain.AcademicYearStatusActive)
	if err != nil {
		return err
	}

	// Activate the target academic year
	activateQuery := `
		UPDATE academic_years
		SET status = $1, updated_at = $2
		WHERE id = $3
	`
	_, err = tx.ExecContext(ctx, activateQuery, domain.AcademicYearStatusActive, time.Now(), id)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// GetSemestersByAcademicYearID retrieves all semesters for an academic year
func (r *AcademicYearRepository) GetSemestersByAcademicYearID(ctx context.Context, academicYearID string) ([]*domain.Semester, error) {
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

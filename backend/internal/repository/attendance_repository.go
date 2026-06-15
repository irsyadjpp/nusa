package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/nusa/backend/internal/domain"
)

// AttendanceRepository handles database operations for attendance records
type AttendanceRepository struct {
	db *sqlx.DB
}

// NewAttendanceRepository creates a new attendance repository
func NewAttendanceRepository(db *sqlx.DB) *AttendanceRepository {
	return &AttendanceRepository{db: db}
}

// Create creates a new attendance record
func (r *AttendanceRepository) Create(ctx context.Context, attendance *domain.AttendanceRecord) error {
	query := `
		INSERT INTO attendance_records (id, class_id, student_id, date, status, notes, recorded_by, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		ON CONFLICT (class_id, student_id, date) DO UPDATE SET
			status = EXCLUDED.status,
			notes = EXCLUDED.notes,
			recorded_by = EXCLUDED.recorded_by,
			updated_at = EXCLUDED.updated_at
	`

	_, err := r.db.ExecContext(ctx, query,
		attendance.ID,
		attendance.ClassID,
		attendance.StudentID,
		attendance.Date,
		attendance.Status,
		attendance.Notes,
		attendance.RecordedBy,
		attendance.CreatedAt,
		attendance.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to create attendance record: %w", err)
	}

	return nil
}

// GetByID retrieves an attendance record by ID
func (r *AttendanceRepository) GetByID(ctx context.Context, id string) (*domain.AttendanceRecord, error) {
	query := `
		SELECT id, class_id, student_id, date, status, notes, recorded_by, created_at, updated_at, deleted_at
		FROM attendance_records
		WHERE id = $1 AND deleted_at IS NULL
	`

	var attendance domain.AttendanceRecord
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&attendance.ID,
		&attendance.ClassID,
		&attendance.StudentID,
		&attendance.Date,
		&attendance.Status,
		&attendance.Notes,
		&attendance.RecordedBy,
		&attendance.CreatedAt,
		&attendance.UpdatedAt,
		&attendance.DeletedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("attendance record not found")
	}

	return &attendance, nil
}

// GetByClassStudentDate retrieves an attendance record by class, student, and date
func (r *AttendanceRepository) GetByClassStudentDate(ctx context.Context, classID, studentID string, date time.Time) (*domain.AttendanceRecord, error) {
	query := `
		SELECT id, class_id, student_id, date, status, notes, recorded_by, created_at, updated_at, deleted_at
		FROM attendance_records
		WHERE class_id = $1 AND student_id = $2 AND date = $3 AND deleted_at IS NULL
	`

	var attendance domain.AttendanceRecord
	err := r.db.QueryRowContext(ctx, query, classID, studentID, date).Scan(
		&attendance.ID,
		&attendance.ClassID,
		&attendance.StudentID,
		&attendance.Date,
		&attendance.Status,
		&attendance.Notes,
		&attendance.RecordedBy,
		&attendance.CreatedAt,
		&attendance.UpdatedAt,
		&attendance.DeletedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("attendance record not found")
	}

	return &attendance, nil
}

// List retrieves attendance records with filters and pagination
func (r *AttendanceRepository) List(ctx context.Context, classID, studentID *string, status *string, startDate, endDate *time.Time, limit, offset int) ([]*domain.AttendanceRecord, error) {
	query := `
		SELECT id, class_id, student_id, date, status, notes, recorded_by, created_at, updated_at, deleted_at
		FROM attendance_records
		WHERE deleted_at IS NULL
	`

	args := []interface{}{}
	argCount := 1

	if classID != nil {
		query += fmt.Sprintf(" AND class_id = $%d", argCount)
		args = append(args, *classID)
		argCount++
	}

	if studentID != nil {
		query += fmt.Sprintf(" AND student_id = $%d", argCount)
		args = append(args, *studentID)
		argCount++
	}

	if status != nil {
		query += fmt.Sprintf(" AND status = $%d", argCount)
		args = append(args, *status)
		argCount++
	}

	if startDate != nil {
		query += fmt.Sprintf(" AND date >= $%d", argCount)
		args = append(args, *startDate)
		argCount++
	}

	if endDate != nil {
		query += fmt.Sprintf(" AND date <= $%d", argCount)
		args = append(args, *endDate)
		argCount++
	}

	query += " ORDER BY date DESC, created_at DESC"

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
		return nil, fmt.Errorf("failed to list attendance records: %w", err)
	}
	defer rows.Close()

	var attendances []*domain.AttendanceRecord
	for rows.Next() {
		var attendance domain.AttendanceRecord
		err := rows.Scan(
			&attendance.ID,
			&attendance.ClassID,
			&attendance.StudentID,
			&attendance.Date,
			&attendance.Status,
			&attendance.Notes,
			&attendance.RecordedBy,
			&attendance.CreatedAt,
			&attendance.UpdatedAt,
			&attendance.DeletedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan attendance record: %w", err)
		}
		attendances = append(attendances, &attendance)
	}

	return attendances, nil
}

// Count returns the count of attendance records with filters
func (r *AttendanceRepository) Count(ctx context.Context, classID, studentID *string, status *string, startDate, endDate *time.Time) (int, error) {
	query := `
		SELECT COUNT(*)
		FROM attendance_records
		WHERE deleted_at IS NULL
	`

	args := []interface{}{}
	argCount := 1

	if classID != nil {
		query += fmt.Sprintf(" AND class_id = $%d", argCount)
		args = append(args, *classID)
		argCount++
	}

	if studentID != nil {
		query += fmt.Sprintf(" AND student_id = $%d", argCount)
		args = append(args, *studentID)
		argCount++
	}

	if status != nil {
		query += fmt.Sprintf(" AND status = $%d", argCount)
		args = append(args, *status)
		argCount++
	}

	if startDate != nil {
		query += fmt.Sprintf(" AND date >= $%d", argCount)
		args = append(args, *startDate)
		argCount++
	}

	if endDate != nil {
		query += fmt.Sprintf(" AND date <= $%d", argCount)
		args = append(args, *endDate)
	}

	var count int
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count attendance records: %w", err)
	}

	return count, nil
}

// Update updates an attendance record
func (r *AttendanceRepository) Update(ctx context.Context, attendance *domain.AttendanceRecord) error {
	query := `
		UPDATE attendance_records
		SET status = $1, notes = $2, recorded_by = $3, updated_at = $4
		WHERE id = $5 AND deleted_at IS NULL
	`

	attendance.UpdatedAt = time.Now()

	result, err := r.db.ExecContext(ctx, query,
		attendance.Status,
		attendance.Notes,
		attendance.RecordedBy,
		attendance.UpdatedAt,
		attendance.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update attendance record: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("attendance record not found")
	}

	return nil
}

// Delete soft deletes an attendance record
func (r *AttendanceRepository) Delete(ctx context.Context, id string) error {
	query := `
		UPDATE attendance_records
		SET deleted_at = $1
		WHERE id = $2 AND deleted_at IS NULL
	`

	now := time.Now()
	result, err := r.db.ExecContext(ctx, query, now, id)
	if err != nil {
		return fmt.Errorf("failed to delete attendance record: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("attendance record not found")
	}

	return nil
}

// GetAttendanceStats returns attendance statistics for a class within a date range
func (r *AttendanceRepository) GetAttendanceStats(ctx context.Context, classID string, startDate, endDate time.Time) (map[string]int, error) {
	query := `
		SELECT status, COUNT(*) as count
		FROM attendance_records
		WHERE class_id = $1 AND date >= $2 AND date <= $3 AND deleted_at IS NULL
		GROUP BY status
	`

	rows, err := r.db.QueryContext(ctx, query, classID, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to get attendance stats: %w", err)
	}
	defer rows.Close()

	stats := make(map[string]int)
	for rows.Next() {
		var status string
		var count int
		err := rows.Scan(&status, &count)
		if err != nil {
			return nil, fmt.Errorf("failed to scan attendance stats: %w", err)
		}
		stats[status] = count
	}

	return stats, nil
}

// GetStudentAttendanceStats returns attendance statistics for a student within a date range
func (r *AttendanceRepository) GetStudentAttendanceStats(ctx context.Context, studentID string, startDate, endDate time.Time) (map[string]int, error) {
	query := `
		SELECT status, COUNT(*) as count
		FROM attendance_records
		WHERE student_id = $1 AND date >= $2 AND date <= $3 AND deleted_at IS NULL
		GROUP BY status
	`

	rows, err := r.db.QueryContext(ctx, query, studentID, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to get student attendance stats: %w", err)
	}
	defer rows.Close()

	stats := make(map[string]int)
	for rows.Next() {
		var status string
		var count int
		err := rows.Scan(&status, &count)
		if err != nil {
			return nil, fmt.Errorf("failed to scan student attendance stats: %w", err)
		}
		stats[status] = count
	}

	return stats, nil
}

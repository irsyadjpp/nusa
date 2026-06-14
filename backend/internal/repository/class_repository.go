package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/nusa/backend/internal/domain"
)

// ClassRepository handles database operations for classes
type ClassRepository struct {
	db *sqlx.DB
}

// NewClassRepository creates a new class repository
func NewClassRepository(db *sqlx.DB) *ClassRepository {
	return &ClassRepository{db: db}
}

// Create creates a new class
func (r *ClassRepository) Create(ctx context.Context, class *domain.Class) error {
	query := `
		INSERT INTO classes (id, school_id, academic_year_id, semester_id, subject_id, teacher_id, name, grade_level, room, max_students, is_active, created_by, updated_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	`

	_, err := r.db.ExecContext(ctx, query,
		class.ID, class.SchoolID, class.AcademicYearID, class.SemesterID, class.SubjectID, class.TeacherID,
		class.Name, class.GradeLevel, class.Room, class.MaxStudents, class.IsActive, class.CreatedBy, class.UpdatedBy)
	return err
}

// GetByID retrieves a class by ID
func (r *ClassRepository) GetByID(ctx context.Context, id string) (*domain.Class, error) {
	query := `
		SELECT id, school_id, academic_year_id, semester_id, subject_id, teacher_id, name, grade_level, room, max_students, is_active,
		       created_at, updated_at, created_by, updated_by, deleted_at
		FROM classes WHERE id = $1
	`

	var class domain.Class
	var room sql.NullString
	var createdBy sql.NullString
	var updatedBy sql.NullString
	var deletedAt sql.NullTime

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&class.ID, &class.SchoolID, &class.AcademicYearID, &class.SemesterID, &class.SubjectID, &class.TeacherID,
		&class.Name, &class.GradeLevel, &room, &class.MaxStudents, &class.IsActive,
		&class.CreatedAt, &class.UpdatedAt, &createdBy, &updatedBy, &deletedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("class not found")
	}
	if err != nil {
		return nil, err
	}

	if room.Valid {
		class.Room = &room.String
	}
	if createdBy.Valid {
		class.CreatedBy = &createdBy.String
	}
	if updatedBy.Valid {
		class.UpdatedBy = &updatedBy.String
	}
	if deletedAt.Valid {
		class.DeletedAt = &deletedAt.Time
	}

	return &class, nil
}

// List retrieves classes with pagination and filters
func (r *ClassRepository) List(ctx context.Context, schoolID, academicYearID, semesterID, subjectID, teacherID *string, isActive *bool, limit, offset int) ([]*domain.Class, error) {
	query := `
		SELECT id, school_id, academic_year_id, semester_id, subject_id, teacher_id, name, grade_level, room, max_students, is_active,
		       created_at, updated_at, created_by, updated_by, deleted_at
		FROM classes WHERE deleted_at IS NULL
	`

	args := []interface{}{}
	argCount := 1

	if schoolID != nil {
		query += fmt.Sprintf(" AND school_id = $%d", argCount)
		args = append(args, *schoolID)
		argCount++
	}
	if academicYearID != nil {
		query += fmt.Sprintf(" AND academic_year_id = $%d", argCount)
		args = append(args, *academicYearID)
		argCount++
	}
	if semesterID != nil {
		query += fmt.Sprintf(" AND semester_id = $%d", argCount)
		args = append(args, *semesterID)
		argCount++
	}
	if subjectID != nil {
		query += fmt.Sprintf(" AND subject_id = $%d", argCount)
		args = append(args, *subjectID)
		argCount++
	}
	if teacherID != nil {
		query += fmt.Sprintf(" AND teacher_id = $%d", argCount)
		args = append(args, *teacherID)
		argCount++
	}
	if isActive != nil {
		query += fmt.Sprintf(" AND is_active = $%d", argCount)
		args = append(args, *isActive)
		argCount++
	}

	query += fmt.Sprintf(" ORDER BY created_at DESC LIMIT $%d OFFSET $%d", argCount, argCount+1)
	args = append(args, limit, offset)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var classes []*domain.Class
	for rows.Next() {
		var class domain.Class
		var room sql.NullString
		var createdBy sql.NullString
		var updatedBy sql.NullString
		var deletedAt sql.NullTime

		err := rows.Scan(
			&class.ID, &class.SchoolID, &class.AcademicYearID, &class.SemesterID, &class.SubjectID, &class.TeacherID,
			&class.Name, &class.GradeLevel, &room, &class.MaxStudents, &class.IsActive,
			&class.CreatedAt, &class.UpdatedAt, &createdBy, &updatedBy, &deletedAt,
		)
		if err != nil {
			return nil, err
		}

		if room.Valid {
			class.Room = &room.String
		}
		if createdBy.Valid {
			class.CreatedBy = &createdBy.String
		}
		if updatedBy.Valid {
			class.UpdatedBy = &updatedBy.String
		}
		if deletedAt.Valid {
			class.DeletedAt = &deletedAt.Time
		}

		classes = append(classes, &class)
	}

	return classes, nil
}

// Count returns the total count of classes with filters
func (r *ClassRepository) Count(ctx context.Context, schoolID, academicYearID, semesterID, subjectID, teacherID *string, isActive *bool) (int, error) {
	query := `SELECT COUNT(*) FROM classes WHERE deleted_at IS NULL`

	args := []interface{}{}
	argCount := 1

	if schoolID != nil {
		query += fmt.Sprintf(" AND school_id = $%d", argCount)
		args = append(args, *schoolID)
		argCount++
	}
	if academicYearID != nil {
		query += fmt.Sprintf(" AND academic_year_id = $%d", argCount)
		args = append(args, *academicYearID)
		argCount++
	}
	if semesterID != nil {
		query += fmt.Sprintf(" AND semester_id = $%d", argCount)
		args = append(args, *semesterID)
		argCount++
	}
	if subjectID != nil {
		query += fmt.Sprintf(" AND subject_id = $%d", argCount)
		args = append(args, *subjectID)
		argCount++
	}
	if teacherID != nil {
		query += fmt.Sprintf(" AND teacher_id = $%d", argCount)
		args = append(args, *teacherID)
		argCount++
	}
	if isActive != nil {
		query += fmt.Sprintf(" AND is_active = $%d", argCount)
		args = append(args, *isActive)
		argCount++
	}

	var count int
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&count)
	return count, err
}

// Update updates a class
func (r *ClassRepository) Update(ctx context.Context, class *domain.Class) error {
	query := `
		UPDATE classes
		SET name = $2, grade_level = $3, room = $4, max_students = $5, is_active = $6, updated_by = $7, updated_at = $8
		WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, query,
		class.ID, class.Name, class.GradeLevel, class.Room, class.MaxStudents, class.IsActive, class.UpdatedBy, class.UpdatedAt)
	return err
}

// Delete soft deletes a class
func (r *ClassRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE classes SET deleted_at = NOW() WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

// GetStudentCount returns the number of students enrolled in a class
func (r *ClassRepository) GetStudentCount(ctx context.Context, classID string) (int, error) {
	query := `SELECT COUNT(*) FROM class_enrollments WHERE class_id = $1 AND status = 'ACTIVE' AND deleted_at IS NULL`

	var count int
	err := r.db.QueryRowContext(ctx, query, classID).Scan(&count)
	return count, err
}

// ClassEnrollmentRepository handles database operations for class enrollments
type ClassEnrollmentRepository struct {
	db *sqlx.DB
}

// NewClassEnrollmentRepository creates a new class enrollment repository
func NewClassEnrollmentRepository(db *sqlx.DB) *ClassEnrollmentRepository {
	return &ClassEnrollmentRepository{db: db}
}

// Create creates a new class enrollment
func (r *ClassEnrollmentRepository) Create(ctx context.Context, enrollment *domain.ClassEnrollment) error {
	query := `
		INSERT INTO class_enrollments (id, class_id, student_id, enrollment_date, status, notes)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.ExecContext(ctx, query,
		enrollment.ID, enrollment.ClassID, enrollment.StudentID, enrollment.EnrollmentDate, enrollment.Status, enrollment.Notes)
	return err
}

// GetByID retrieves a class enrollment by ID
func (r *ClassEnrollmentRepository) GetByID(ctx context.Context, id string) (*domain.ClassEnrollment, error) {
	query := `
		SELECT id, class_id, student_id, enrollment_date, status, notes, created_at, updated_at, deleted_at
		FROM class_enrollments WHERE id = $1
	`

	var enrollment domain.ClassEnrollment
	var notes sql.NullString
	var deletedAt sql.NullTime

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&enrollment.ID, &enrollment.ClassID, &enrollment.StudentID, &enrollment.EnrollmentDate,
		&enrollment.Status, &notes, &enrollment.CreatedAt, &enrollment.UpdatedAt, &deletedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("enrollment not found")
	}
	if err != nil {
		return nil, err
	}

	if notes.Valid {
		enrollment.Notes = &notes.String
	}
	if deletedAt.Valid {
		enrollment.DeletedAt = &deletedAt.Time
	}

	return &enrollment, nil
}

// List retrieves class enrollments with pagination and filters
func (r *ClassEnrollmentRepository) List(ctx context.Context, classID, studentID *string, status *string, limit, offset int) ([]*domain.ClassEnrollment, error) {
	query := `
		SELECT id, class_id, student_id, enrollment_date, status, notes, created_at, updated_at, deleted_at
		FROM class_enrollments WHERE deleted_at IS NULL
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

	query += fmt.Sprintf(" ORDER BY enrollment_date DESC LIMIT $%d OFFSET $%d", argCount, argCount+1)
	args = append(args, limit, offset)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var enrollments []*domain.ClassEnrollment
	for rows.Next() {
		var enrollment domain.ClassEnrollment
		var notes sql.NullString
		var deletedAt sql.NullTime

		err := rows.Scan(
			&enrollment.ID, &enrollment.ClassID, &enrollment.StudentID, &enrollment.EnrollmentDate,
			&enrollment.Status, &notes, &enrollment.CreatedAt, &enrollment.UpdatedAt, &deletedAt,
		)
		if err != nil {
			return nil, err
		}

		if notes.Valid {
			enrollment.Notes = &notes.String
		}
		if deletedAt.Valid {
			enrollment.DeletedAt = &deletedAt.Time
		}

		enrollments = append(enrollments, &enrollment)
	}

	return enrollments, nil
}

// Update updates a class enrollment
func (r *ClassEnrollmentRepository) Update(ctx context.Context, enrollment *domain.ClassEnrollment) error {
	query := `
		UPDATE class_enrollments
		SET status = $2, notes = $3, updated_at = $4
		WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, query,
		enrollment.ID, enrollment.Status, enrollment.Notes, enrollment.UpdatedAt)
	return err
}

// Delete soft deletes a class enrollment
func (r *ClassEnrollmentRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE class_enrollments SET deleted_at = NOW() WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

// CheckEnrollment checks if a student is enrolled in a class
func (r *ClassEnrollmentRepository) CheckEnrollment(ctx context.Context, classID, studentID string) (bool, error) {
	query := `SELECT id FROM class_enrollments WHERE class_id = $1 AND student_id = $2 AND status = 'ACTIVE' AND deleted_at IS NULL`

	var id string
	err := r.db.QueryRowContext(ctx, query, classID, studentID).Scan(&id)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/nusa/backend/internal/domain"
)

// SchoolRepository handles database operations for schools
type SchoolRepository struct {
	db *sqlx.DB
}

// NewSchoolRepository creates a new school repository
func NewSchoolRepository(db *sqlx.DB) *SchoolRepository {
	return &SchoolRepository{db: db}
}

// Create creates a new school
func (r *SchoolRepository) Create(ctx context.Context, school *domain.School) error {
	query := `
		INSERT INTO schools (id, name, code, address, phone, email, is_active, created_by, updated_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	_, err := r.db.ExecContext(ctx, query,
		school.ID, school.Name, school.Code, school.Address, school.Phone, school.Email,
		school.IsActive, school.CreatedBy, school.UpdatedBy)
	return err
}

// GetByID retrieves a school by ID
func (r *SchoolRepository) GetByID(ctx context.Context, id string) (*domain.School, error) {
	query := `
		SELECT id, name, code, address, phone, email, is_active, created_at, updated_at, created_by, updated_by
		FROM schools WHERE id = $1
	`

	var school domain.School
	var address, phone, email, createdBy, updatedBy sql.NullString

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&school.ID, &school.Name, &school.Code, &address, &phone, &email,
		&school.IsActive, &school.CreatedAt, &school.UpdatedAt, &createdBy, &updatedBy,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("school not found")
	}
	if err != nil {
		return nil, err
	}

	if address.Valid {
		school.Address = &address.String
	}
	if phone.Valid {
		school.Phone = &phone.String
	}
	if email.Valid {
		school.Email = &email.String
	}
	if createdBy.Valid {
		school.CreatedBy = &createdBy.String
	}
	if updatedBy.Valid {
		school.UpdatedBy = &updatedBy.String
	}

	return &school, nil
}

// GetByCode retrieves a school by code
func (r *SchoolRepository) GetByCode(ctx context.Context, code string) (*domain.School, error) {
	query := `
		SELECT id, name, code, address, phone, email, is_active, created_at, updated_at, created_by, updated_by
		FROM schools WHERE code = $1
	`

	var school domain.School
	var address, phone, email, createdBy, updatedBy sql.NullString

	err := r.db.QueryRowContext(ctx, query, code).Scan(
		&school.ID, &school.Name, &school.Code, &address, &phone, &email,
		&school.IsActive, &school.CreatedAt, &school.UpdatedAt, &createdBy, &updatedBy,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("school not found")
	}
	if err != nil {
		return nil, err
	}

	if address.Valid {
		school.Address = &address.String
	}
	if phone.Valid {
		school.Phone = &phone.String
	}
	if email.Valid {
		school.Email = &email.String
	}
	if createdBy.Valid {
		school.CreatedBy = &createdBy.String
	}
	if updatedBy.Valid {
		school.UpdatedBy = &updatedBy.String
	}

	return &school, nil
}

// Update updates a school
func (r *SchoolRepository) Update(ctx context.Context, school *domain.School) error {
	query := `
		UPDATE schools 
		SET name = $2, code = $3, address = $4, phone = $5, email = $6, 
		    updated_at = NOW(), updated_by = $7
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query,
		school.ID, school.Name, school.Code, school.Address, school.Phone, school.Email, school.UpdatedBy)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("school not found")
	}

	return nil
}

// UpdateStatus updates school status
func (r *SchoolRepository) UpdateStatus(ctx context.Context, id string, isActive bool) error {
	query := `UPDATE schools SET is_active = $2, updated_at = NOW() WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id, isActive)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("school not found")
	}

	return nil
}

// List retrieves schools with optional filters
func (r *SchoolRepository) List(ctx context.Context, isActive *bool, limit, offset int) ([]*domain.School, error) {
	query := `
		SELECT id, name, code, address, phone, email, is_active, created_at, updated_at, created_by, updated_by
		FROM schools
		WHERE 1=1
	`

	args := []interface{}{}
	argIndex := 1

	if isActive != nil {
		query += fmt.Sprintf(" AND is_active = $%d", argIndex)
		args = append(args, *isActive)
		argIndex++
	}

	query += " ORDER BY created_at DESC"

	if limit > 0 {
		query += fmt.Sprintf(" LIMIT $%d", argIndex)
		args = append(args, limit)
		argIndex++
	}

	if offset > 0 {
		query += fmt.Sprintf(" OFFSET $%d", argIndex)
		args = append(args, offset)
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schools []*domain.School
	for rows.Next() {
		var school domain.School
		var address, phone, email, createdBy, updatedBy sql.NullString

		err := rows.Scan(
			&school.ID, &school.Name, &school.Code, &address, &phone, &email,
			&school.IsActive, &school.CreatedAt, &school.UpdatedAt, &createdBy, &updatedBy,
		)
		if err != nil {
			return nil, err
		}

		if address.Valid {
			school.Address = &address.String
		}
		if phone.Valid {
			school.Phone = &phone.String
		}
		if email.Valid {
			school.Email = &email.String
		}
		if createdBy.Valid {
			school.CreatedBy = &createdBy.String
		}
		if updatedBy.Valid {
			school.UpdatedBy = &updatedBy.String
		}

		schools = append(schools, &school)
	}

	return schools, nil
}

// Count returns the total count of schools with optional filters
func (r *SchoolRepository) Count(ctx context.Context, isActive *bool) (int, error) {
	query := `SELECT COUNT(*) FROM schools WHERE 1=1`
	args := []interface{}{}

	if isActive != nil {
		query += " AND is_active = $1"
		args = append(args, *isActive)
	}

	var count int
	err := r.db.GetContext(ctx, &count, query, args...)
	return count, err
}

// Delete soft deletes a school (sets is_active to false)
func (r *SchoolRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE schools SET is_active = false, updated_at = NOW() WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("school not found")
	}

	return nil
}

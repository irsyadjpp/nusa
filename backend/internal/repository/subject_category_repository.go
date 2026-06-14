package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/nusa/backend/internal/domain"
)

// SubjectCategoryRepository handles database operations for subject categories
type SubjectCategoryRepository struct {
	db *sqlx.DB
}

// NewSubjectCategoryRepository creates a new subject category repository
func NewSubjectCategoryRepository(db *sqlx.DB) *SubjectCategoryRepository {
	return &SubjectCategoryRepository{db: db}
}

// CreateSubjectCategory creates a new subject category
func (r *SubjectCategoryRepository) CreateSubjectCategory(ctx context.Context, sc *domain.SubjectCategory) error {
	query := `
		INSERT INTO subject_categories (id, code, name, description, is_mandatory, is_active, created_by, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	_, err := r.db.ExecContext(ctx, query,
		sc.ID, sc.Code, sc.Name, sc.Description, sc.IsMandatory,
		sc.IsActive, sc.CreatedBy, sc.CreatedAt, sc.UpdatedAt)
	return err
}

// GetSubjectCategoryByID retrieves a subject category by ID
func (r *SubjectCategoryRepository) GetSubjectCategoryByID(ctx context.Context, id string) (*domain.SubjectCategory, error) {
	query := `
		SELECT id, code, name, description, is_mandatory, is_active, created_by, created_at, updated_at, updated_by
		FROM subject_categories WHERE id = $1
	`

	var sc domain.SubjectCategory
	var description, updatedBy sql.NullString

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&sc.ID, &sc.Code, &sc.Name, &description, &sc.IsMandatory,
		&sc.IsActive, &sc.CreatedBy, &sc.CreatedAt, &sc.UpdatedAt, &updatedBy,
	)
	if err != nil {
		return nil, err
	}

	if description.Valid {
		sc.Description = &description.String
	}
	if updatedBy.Valid {
		sc.UpdatedBy = &updatedBy.String
	}

	return &sc, nil
}

// GetAllSubjectCategories retrieves all subject categories
func (r *SubjectCategoryRepository) GetAllSubjectCategories(ctx context.Context) ([]*domain.SubjectCategory, error) {
	query := `
		SELECT id, code, name, description, is_mandatory, is_active, created_by, created_at, updated_at, updated_by
		FROM subject_categories ORDER BY code ASC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*domain.SubjectCategory
	for rows.Next() {
		var sc domain.SubjectCategory
		var description, updatedBy sql.NullString

		err := rows.Scan(
			&sc.ID, &sc.Code, &sc.Name, &description, &sc.IsMandatory,
			&sc.IsActive, &sc.CreatedBy, &sc.CreatedAt, &sc.UpdatedAt, &updatedBy,
		)
		if err != nil {
			return nil, err
		}

		if description.Valid {
			sc.Description = &description.String
		}
		if updatedBy.Valid {
			sc.UpdatedBy = &updatedBy.String
		}

		categories = append(categories, &sc)
	}
	return categories, nil
}

// GetActiveSubjectCategories retrieves all active subject categories
func (r *SubjectCategoryRepository) GetActiveSubjectCategories(ctx context.Context) ([]*domain.SubjectCategory, error) {
	query := `
		SELECT id, code, name, description, is_mandatory, is_active, created_by, created_at, updated_at, updated_by
		FROM subject_categories WHERE is_active = true ORDER BY code ASC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*domain.SubjectCategory
	for rows.Next() {
		var sc domain.SubjectCategory
		var description, updatedBy sql.NullString

		err := rows.Scan(
			&sc.ID, &sc.Code, &sc.Name, &description, &sc.IsMandatory,
			&sc.IsActive, &sc.CreatedBy, &sc.CreatedAt, &sc.UpdatedAt, &updatedBy,
		)
		if err != nil {
			return nil, err
		}

		if description.Valid {
			sc.Description = &description.String
		}
		if updatedBy.Valid {
			sc.UpdatedBy = &updatedBy.String
		}

		categories = append(categories, &sc)
	}
	return categories, nil
}

// UpdateSubjectCategory updates a subject category
func (r *SubjectCategoryRepository) UpdateSubjectCategory(ctx context.Context, sc *domain.SubjectCategory) error {
	query := `
		UPDATE subject_categories
		SET name = $1, description = $2, is_mandatory = $3, is_active = $4, updated_at = $5, updated_by = $6
		WHERE id = $7
	`

	result, err := r.db.ExecContext(ctx, query,
		sc.Name, sc.Description, sc.IsMandatory, sc.IsActive, sc.UpdatedAt, sc.UpdatedBy, sc.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("subject category not found")
	}
	return nil
}

// DeleteSubjectCategory deletes a subject category
func (r *SubjectCategoryRepository) DeleteSubjectCategory(ctx context.Context, id string) error {
	query := `DELETE FROM subject_categories WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("subject category not found")
	}
	return nil
}

// CheckCodeExists checks if a subject category code already exists
func (r *SubjectCategoryRepository) CheckCodeExists(ctx context.Context, code string, excludeID string) (bool, error) {
	query := `SELECT COUNT(*) FROM subject_categories WHERE code = $1`
	args := []interface{}{code}

	if excludeID != "" {
		query += ` AND id != $2`
		args = append(args, excludeID)
	}

	var count int
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

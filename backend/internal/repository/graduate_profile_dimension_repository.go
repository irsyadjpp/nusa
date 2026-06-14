package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/nusa/backend/internal/domain"
)

// GraduateProfileDimensionRepository handles database operations for graduate profile dimensions
type GraduateProfileDimensionRepository struct {
	db *sqlx.DB
}

// NewGraduateProfileDimensionRepository creates a new graduate profile dimension repository
func NewGraduateProfileDimensionRepository(db *sqlx.DB) *GraduateProfileDimensionRepository {
	return &GraduateProfileDimensionRepository{db: db}
}

// CreateGraduateProfileDimension creates a new graduate profile dimension
func (r *GraduateProfileDimensionRepository) CreateGraduateProfileDimension(ctx context.Context, gpd *domain.GraduateProfileDimension) error {
	query := `
		INSERT INTO graduate_profile_dimensions (id, code, name, description, sequence_number, is_active, created_by, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	_, err := r.db.ExecContext(ctx, query,
		gpd.ID, gpd.Code, gpd.Name, gpd.Description,
		gpd.SequenceNumber, gpd.IsActive, gpd.CreatedBy, gpd.CreatedAt, gpd.UpdatedAt)
	return err
}

// GetGraduateProfileDimensionByID retrieves a graduate profile dimension by ID
func (r *GraduateProfileDimensionRepository) GetGraduateProfileDimensionByID(ctx context.Context, id string) (*domain.GraduateProfileDimension, error) {
	query := `
		SELECT id, code, name, description, sequence_number, is_active, created_by, created_at, updated_at, updated_by
		FROM graduate_profile_dimensions WHERE id = $1
	`

	var gpd domain.GraduateProfileDimension
	var description, updatedBy sql.NullString

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&gpd.ID, &gpd.Code, &gpd.Name, &description,
		&gpd.SequenceNumber, &gpd.IsActive, &gpd.CreatedBy, &gpd.CreatedAt, &gpd.UpdatedAt, &updatedBy,
	)
	if err != nil {
		return nil, err
	}

	if description.Valid {
		gpd.Description = &description.String
	}
	if updatedBy.Valid {
		gpd.UpdatedBy = &updatedBy.String
	}

	return &gpd, nil
}

// GetAllGraduateProfileDimensions retrieves all graduate profile dimensions
func (r *GraduateProfileDimensionRepository) GetAllGraduateProfileDimensions(ctx context.Context) ([]*domain.GraduateProfileDimension, error) {
	query := `
		SELECT id, code, name, description, sequence_number, is_active, created_by, created_at, updated_at, updated_by
		FROM graduate_profile_dimensions ORDER BY sequence_number ASC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dimensions []*domain.GraduateProfileDimension
	for rows.Next() {
		var gpd domain.GraduateProfileDimension
		var description, updatedBy sql.NullString

		err := rows.Scan(
			&gpd.ID, &gpd.Code, &gpd.Name, &description,
			&gpd.SequenceNumber, &gpd.IsActive, &gpd.CreatedBy, &gpd.CreatedAt, &gpd.UpdatedAt, &updatedBy,
		)
		if err != nil {
			return nil, err
		}

		if description.Valid {
			gpd.Description = &description.String
		}
		if updatedBy.Valid {
			gpd.UpdatedBy = &updatedBy.String
		}

		dimensions = append(dimensions, &gpd)
	}
	return dimensions, nil
}

// GetActiveGraduateProfileDimensions retrieves all active graduate profile dimensions
func (r *GraduateProfileDimensionRepository) GetActiveGraduateProfileDimensions(ctx context.Context) ([]*domain.GraduateProfileDimension, error) {
	query := `
		SELECT id, code, name, description, sequence_number, is_active, created_by, created_at, updated_at, updated_by
		FROM graduate_profile_dimensions WHERE is_active = true ORDER BY sequence_number ASC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dimensions []*domain.GraduateProfileDimension
	for rows.Next() {
		var gpd domain.GraduateProfileDimension
		var description, updatedBy sql.NullString

		err := rows.Scan(
			&gpd.ID, &gpd.Code, &gpd.Name, &description,
			&gpd.SequenceNumber, &gpd.IsActive, &gpd.CreatedBy, &gpd.CreatedAt, &gpd.UpdatedAt, &updatedBy,
		)
		if err != nil {
			return nil, err
		}

		if description.Valid {
			gpd.Description = &description.String
		}
		if updatedBy.Valid {
			gpd.UpdatedBy = &updatedBy.String
		}

		dimensions = append(dimensions, &gpd)
	}
	return dimensions, nil
}

// UpdateGraduateProfileDimension updates a graduate profile dimension
func (r *GraduateProfileDimensionRepository) UpdateGraduateProfileDimension(ctx context.Context, gpd *domain.GraduateProfileDimension) error {
	query := `
		UPDATE graduate_profile_dimensions
		SET name = $1, description = $2, sequence_number = $3, is_active = $4, updated_at = $5, updated_by = $6
		WHERE id = $7
	`

	result, err := r.db.ExecContext(ctx, query,
		gpd.Name, gpd.Description, gpd.SequenceNumber,
		gpd.IsActive, gpd.UpdatedAt, gpd.UpdatedBy, gpd.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("graduate profile dimension not found")
	}
	return nil
}

// DeleteGraduateProfileDimension deletes a graduate profile dimension
func (r *GraduateProfileDimensionRepository) DeleteGraduateProfileDimension(ctx context.Context, id string) error {
	query := `DELETE FROM graduate_profile_dimensions WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("graduate profile dimension not found")
	}
	return nil
}

// CheckCodeExists checks if a graduate profile dimension code already exists
func (r *GraduateProfileDimensionRepository) CheckCodeExists(ctx context.Context, code string, excludeID string) (bool, error) {
	query := `SELECT COUNT(*) FROM graduate_profile_dimensions WHERE code = $1`
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

// CheckSequenceNumberExists checks if a sequence number already exists
func (r *GraduateProfileDimensionRepository) CheckSequenceNumberExists(ctx context.Context, sequenceNumber int, excludeID string) (bool, error) {
	query := `SELECT COUNT(*) FROM graduate_profile_dimensions WHERE sequence_number = $1`
	args := []interface{}{sequenceNumber}

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

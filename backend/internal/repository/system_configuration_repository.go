package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/nusa/backend/internal/domain"
)

// SystemConfigurationRepository handles database operations for system configurations
type SystemConfigurationRepository struct {
	db *sqlx.DB
}

// NewSystemConfigurationRepository creates a new system configuration repository
func NewSystemConfigurationRepository(db *sqlx.DB) *SystemConfigurationRepository {
	return &SystemConfigurationRepository{db: db}
}

// CreateSystemConfiguration creates a new system configuration
func (r *SystemConfigurationRepository) CreateSystemConfiguration(ctx context.Context, sc *domain.SystemConfiguration) error {
	query := `
		INSERT INTO system_configurations (id, key, value, value_type, description, category, is_system, is_active, created_by, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`

	_, err := r.db.ExecContext(ctx, query,
		sc.ID, sc.Key, sc.Value, sc.ValueType, sc.Description, sc.Category,
		sc.IsSystem, sc.IsActive, sc.CreatedBy, sc.CreatedAt, sc.UpdatedAt)
	return err
}

// GetSystemConfigurationByID retrieves a system configuration by ID
func (r *SystemConfigurationRepository) GetSystemConfigurationByID(ctx context.Context, id string) (*domain.SystemConfiguration, error) {
	query := `
		SELECT id, key, value, value_type, description, category, is_system, is_active, created_by, created_at, updated_at, updated_by
		FROM system_configurations WHERE id = $1
	`

	var sc domain.SystemConfiguration
	var description, updatedBy sql.NullString

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&sc.ID, &sc.Key, &sc.Value, &sc.ValueType, &description, &sc.Category,
		&sc.IsSystem, &sc.IsActive, &sc.CreatedBy, &sc.CreatedAt, &sc.UpdatedAt, &updatedBy,
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

// GetSystemConfigurationByKey retrieves a system configuration by key
func (r *SystemConfigurationRepository) GetSystemConfigurationByKey(ctx context.Context, key string) (*domain.SystemConfiguration, error) {
	query := `
		SELECT id, key, value, value_type, description, category, is_system, is_active, created_by, created_at, updated_at, updated_by
		FROM system_configurations WHERE key = $1 AND is_active = true
	`

	var sc domain.SystemConfiguration
	var description, updatedBy sql.NullString

	err := r.db.QueryRowContext(ctx, query, key).Scan(
		&sc.ID, &sc.Key, &sc.Value, &sc.ValueType, &description, &sc.Category,
		&sc.IsSystem, &sc.IsActive, &sc.CreatedBy, &sc.CreatedAt, &sc.UpdatedAt, &updatedBy,
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

// GetAllSystemConfigurations retrieves all system configurations
func (r *SystemConfigurationRepository) GetAllSystemConfigurations(ctx context.Context) ([]*domain.SystemConfiguration, error) {
	query := `
		SELECT id, key, value, value_type, description, category, is_system, is_active, created_by, created_at, updated_at, updated_by
		FROM system_configurations ORDER BY category ASC, key ASC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var configurations []*domain.SystemConfiguration
	for rows.Next() {
		var sc domain.SystemConfiguration
		var description, updatedBy sql.NullString

		err := rows.Scan(
			&sc.ID, &sc.Key, &sc.Value, &sc.ValueType, &description, &sc.Category,
			&sc.IsSystem, &sc.IsActive, &sc.CreatedBy, &sc.CreatedAt, &sc.UpdatedAt, &updatedBy,
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

		configurations = append(configurations, &sc)
	}
	return configurations, nil
}

// GetSystemConfigurationsByCategory retrieves all system configurations for a category
func (r *SystemConfigurationRepository) GetSystemConfigurationsByCategory(ctx context.Context, category string) ([]*domain.SystemConfiguration, error) {
	query := `
		SELECT id, key, value, value_type, description, category, is_system, is_active, created_by, created_at, updated_at, updated_by
		FROM system_configurations WHERE category = $1 ORDER BY key ASC
	`

	rows, err := r.db.QueryContext(ctx, query, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var configurations []*domain.SystemConfiguration
	for rows.Next() {
		var sc domain.SystemConfiguration
		var description, updatedBy sql.NullString

		err := rows.Scan(
			&sc.ID, &sc.Key, &sc.Value, &sc.ValueType, &description, &sc.Category,
			&sc.IsSystem, &sc.IsActive, &sc.CreatedBy, &sc.CreatedAt, &sc.UpdatedAt, &updatedBy,
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

		configurations = append(configurations, &sc)
	}
	return configurations, nil
}

// GetActiveSystemConfigurations retrieves all active system configurations
func (r *SystemConfigurationRepository) GetActiveSystemConfigurations(ctx context.Context) ([]*domain.SystemConfiguration, error) {
	query := `
		SELECT id, key, value, value_type, description, category, is_system, is_active, created_by, created_at, updated_at, updated_by
		FROM system_configurations WHERE is_active = true ORDER BY category ASC, key ASC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var configurations []*domain.SystemConfiguration
	for rows.Next() {
		var sc domain.SystemConfiguration
		var description, updatedBy sql.NullString

		err := rows.Scan(
			&sc.ID, &sc.Key, &sc.Value, &sc.ValueType, &description, &sc.Category,
			&sc.IsSystem, &sc.IsActive, &sc.CreatedBy, &sc.CreatedAt, &sc.UpdatedAt, &updatedBy,
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

		configurations = append(configurations, &sc)
	}
	return configurations, nil
}

// UpdateSystemConfiguration updates a system configuration
func (r *SystemConfigurationRepository) UpdateSystemConfiguration(ctx context.Context, sc *domain.SystemConfiguration) error {
	query := `
		UPDATE system_configurations
		SET value = $1, value_type = $2, description = $3, category = $4, is_active = $5, updated_at = $6, updated_by = $7
		WHERE id = $8
	`

	result, err := r.db.ExecContext(ctx, query,
		sc.Value, sc.ValueType, sc.Description, sc.Category, sc.IsActive,
		sc.UpdatedAt, sc.UpdatedBy, sc.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("system configuration not found")
	}
	return nil
}

// DeleteSystemConfiguration deletes a system configuration
func (r *SystemConfigurationRepository) DeleteSystemConfiguration(ctx context.Context, id string) error {
	query := `DELETE FROM system_configurations WHERE id = $1 AND is_system = false`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("system configuration not found or is a system configuration that cannot be deleted")
	}
	return nil
}

// CheckKeyExists checks if a configuration key already exists
func (r *SystemConfigurationRepository) CheckKeyExists(ctx context.Context, key string, excludeID string) (bool, error) {
	query := `SELECT COUNT(*) FROM system_configurations WHERE key = $1`
	args := []interface{}{key}

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

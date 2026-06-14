package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/nusa/backend/internal/domain"
)

// RoleRepository handles database operations for roles
type RoleRepository struct {
	db *sqlx.DB
}

// NewRoleRepository creates a new role repository
func NewRoleRepository(db *sqlx.DB) *RoleRepository {
	return &RoleRepository{db: db}
}

// Create creates a new role
func (r *RoleRepository) Create(ctx context.Context, role *domain.Role) error {
	query := `
		INSERT INTO roles (id, name, description, is_active)
		VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.ExecContext(ctx, query, role.ID, role.Name, role.Description, role.IsActive)
	return err
}

// GetByID retrieves a role by ID
func (r *RoleRepository) GetByID(ctx context.Context, id string) (*domain.Role, error) {
	query := `
		SELECT id, name, description, is_active, created_at, updated_at
		FROM roles WHERE id = $1
	`

	var role domain.Role
	var description sql.NullString

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&role.ID, &role.Name, &description, &role.IsActive, &role.CreatedAt, &role.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("role not found")
	}
	if err != nil {
		return nil, err
	}

	if description.Valid {
		role.Description = &description.String
	}

	return &role, nil
}

// GetByName retrieves a role by name
func (r *RoleRepository) GetByName(ctx context.Context, name string) (*domain.Role, error) {
	query := `
		SELECT id, name, description, is_active, created_at, updated_at
		FROM roles WHERE name = $1
	`

	var role domain.Role
	var description sql.NullString

	err := r.db.QueryRowContext(ctx, query, name).Scan(
		&role.ID, &role.Name, &description, &role.IsActive, &role.CreatedAt, &role.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("role not found")
	}
	if err != nil {
		return nil, err
	}

	if description.Valid {
		role.Description = &description.String
	}

	return &role, nil
}

// Update updates a role
func (r *RoleRepository) Update(ctx context.Context, role *domain.Role) error {
	query := `
		UPDATE roles 
		SET name = $2, description = $3, is_active = $4, updated_at = NOW()
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query, role.ID, role.Name, role.Description, role.IsActive)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("role not found")
	}

	return nil
}

// List retrieves all roles
func (r *RoleRepository) List(ctx context.Context, isActive *bool) ([]*domain.Role, error) {
	query := `
		SELECT id, name, description, is_active, created_at, updated_at
		FROM roles
		WHERE 1=1
	`

	args := []interface{}{}
	argIndex := 1

	if isActive != nil {
		query += fmt.Sprintf(" AND is_active = $%d", argIndex)
		args = append(args, *isActive)
		argIndex++
	}

	query += " ORDER BY name"
	_ = argIndex // Mark as used to satisfy linter

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []*domain.Role
	for rows.Next() {
		var role domain.Role
		var description sql.NullString

		err := rows.Scan(&role.ID, &role.Name, &description, &role.IsActive, &role.CreatedAt, &role.UpdatedAt)
		if err != nil {
			return nil, err
		}

		if description.Valid {
			role.Description = &description.String
		}

		roles = append(roles, &role)
	}

	return roles, nil
}

// AddPermission adds a permission to a role
func (r *RoleRepository) AddPermission(ctx context.Context, roleID, resource, action string) error {
	query := `
		INSERT INTO permissions (id, role_id, resource, action)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (role_id, resource, action) DO NOTHING
	`

	_, err := r.db.ExecContext(ctx, query, domain.NewID(), roleID, resource, action)
	return err
}

// GetPermissions retrieves all permissions for a role
func (r *RoleRepository) GetPermissions(ctx context.Context, roleID string) ([]*domain.Permission, error) {
	query := `
		SELECT id, role_id, resource, action, created_at
		FROM permissions WHERE role_id = $1
		ORDER BY resource, action
	`

	rows, err := r.db.QueryContext(ctx, query, roleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []*domain.Permission
	for rows.Next() {
		var perm domain.Permission
		err := rows.Scan(&perm.ID, &perm.RoleID, &perm.Resource, &perm.Action, &perm.CreatedAt)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, &perm)
	}

	return permissions, nil
}

// RemovePermission removes a permission from a role
func (r *RoleRepository) RemovePermission(ctx context.Context, roleID, resource, action string) error {
	query := `
		DELETE FROM permissions 
		WHERE role_id = $1 AND resource = $2 AND action = $3
	`

	_, err := r.db.ExecContext(ctx, query, roleID, resource, action)
	return err
}

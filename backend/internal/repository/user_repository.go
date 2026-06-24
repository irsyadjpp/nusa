package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/nusa/backend/internal/domain"
)

// UserRepository handles database operations for users
type UserRepository struct {
	db *sqlx.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create creates a new user
func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {
	query := `
		INSERT INTO users (id, email, password_hash, name, role_id, school_id, is_active, failed_login_attempts, created_by, updated_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	_, err := r.db.ExecContext(ctx, query,
		user.ID, user.Email, user.PasswordHash, user.Name, user.RoleID, user.SchoolID,
		user.IsActive, user.FailedLoginAttempts, user.CreatedBy, user.UpdatedBy)
	return err
}

// GetByID retrieves a user by ID
func (r *UserRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	query := `
		SELECT id, email, password_hash, name, role_id, school_id, is_active, 
		       failed_login_attempts, locked_until, created_at, updated_at, created_by, updated_by
		FROM users WHERE id = $1
	`

	var user domain.User
	var schoolID sql.NullString
	var lockedUntil sql.NullTime
	var createdBy sql.NullString
	var updatedBy sql.NullString

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID, &user.Email, &user.PasswordHash, &user.Name, &user.RoleID, &schoolID,
		&user.IsActive, &user.FailedLoginAttempts, &lockedUntil,
		&user.CreatedAt, &user.UpdatedAt, &createdBy, &updatedBy,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}
	if err != nil {
		return nil, err
	}

	if schoolID.Valid {
		user.SchoolID = &schoolID.String
	}
	if lockedUntil.Valid {
		user.LockedUntil = &lockedUntil.Time
	}
	if createdBy.Valid {
		user.CreatedBy = &createdBy.String
	}
	if updatedBy.Valid {
		user.UpdatedBy = &updatedBy.String
	}

	return &user, nil
}

// GetByEmail retrieves a user by email
func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `
		SELECT id, email, password_hash, name, role_id, school_id, is_active, 
		       failed_login_attempts, locked_until, created_at, updated_at, created_by, updated_by
		FROM users WHERE email = $1
	`

	var user domain.User
	var schoolID sql.NullString
	var lockedUntil sql.NullTime
	var createdBy sql.NullString
	var updatedBy sql.NullString

	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID, &user.Email, &user.PasswordHash, &user.Name, &user.RoleID, &schoolID,
		&user.IsActive, &user.FailedLoginAttempts, &lockedUntil,
		&user.CreatedAt, &user.UpdatedAt, &createdBy, &updatedBy,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}
	if err != nil {
		return nil, err
	}

	if schoolID.Valid {
		user.SchoolID = &schoolID.String
	}
	if lockedUntil.Valid {
		user.LockedUntil = &lockedUntil.Time
	}
	if createdBy.Valid {
		user.CreatedBy = &createdBy.String
	}
	if updatedBy.Valid {
		user.UpdatedBy = &updatedBy.String
	}

	return &user, nil
}

// Update updates a user
func (r *UserRepository) Update(ctx context.Context, user *domain.User) error {
	query := `
		UPDATE users 
		SET name = $2, role_id = $3, school_id = $4, is_active = $5, 
		    failed_login_attempts = $6, locked_until = $7, updated_at = NOW(), updated_by = $8
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query,
		user.ID, user.Name, user.RoleID, user.SchoolID, user.IsActive,
		user.FailedLoginAttempts, user.LockedUntil, user.UpdatedBy)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

// UpdateStatus updates user status and lock state
func (r *UserRepository) UpdateStatus(ctx context.Context, id string, isActive bool, lockedUntil *string, failedAttempts int) error {
	query := `
		UPDATE users 
		SET is_active = $2, locked_until = $3, failed_login_attempts = $4, updated_at = NOW()
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query, id, isActive, lockedUntil, failedAttempts)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

// List retrieves users with optional filters
func (r *UserRepository) List(ctx context.Context, schoolID *string, roleID *string, isActive *bool, limit, offset int) ([]*domain.User, error) {
	query := `
		SELECT id, email, password_hash, name, role_id, school_id, is_active, 
		       failed_login_attempts, locked_until, created_at, updated_at, created_by, updated_by
		FROM users
		WHERE 1=1
	`

	args := []interface{}{}
	argIndex := 1

	if schoolID != nil {
		query += fmt.Sprintf(" AND school_id = $%d", argIndex)
		args = append(args, *schoolID)
		argIndex++
	}

	if roleID != nil {
		query += fmt.Sprintf(" AND role_id = $%d", argIndex)
		args = append(args, *roleID)
		argIndex++
	}

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

	var users []*domain.User
	for rows.Next() {
		var user domain.User
		var schoolID sql.NullString
		var lockedUntil sql.NullTime
		var createdBy sql.NullString
		var updatedBy sql.NullString

		err := rows.Scan(
			&user.ID, &user.Email, &user.PasswordHash, &user.Name, &user.RoleID, &schoolID,
			&user.IsActive, &user.FailedLoginAttempts, &lockedUntil,
			&user.CreatedAt, &user.UpdatedAt, &createdBy, &updatedBy,
		)
		if err != nil {
			return nil, err
		}

		if schoolID.Valid {
			user.SchoolID = &schoolID.String
		}
		if lockedUntil.Valid {
			user.LockedUntil = &lockedUntil.Time
		}
		if createdBy.Valid {
			user.CreatedBy = &createdBy.String
		}
		if updatedBy.Valid {
			user.UpdatedBy = &updatedBy.String
		}

		users = append(users, &user)
	}

	return users, nil
}

// Count returns the total count of users with optional filters
func (r *UserRepository) Count(ctx context.Context, schoolID *string, roleID *string, isActive *bool) (int, error) {
	query := `SELECT COUNT(*) FROM users WHERE 1=1`
	args := []interface{}{}
	argIndex := 1

	if schoolID != nil {
		query += fmt.Sprintf(" AND school_id = $%d", argIndex)
		args = append(args, *schoolID)
		argIndex++
	}

	if roleID != nil {
		query += fmt.Sprintf(" AND role_id = $%d", argIndex)
		args = append(args, *roleID)
		argIndex++
	}

	if isActive != nil {
		query += fmt.Sprintf(" AND is_active = $%d", argIndex)
		args = append(args, *isActive)
		argIndex++
	}

	var count int
	err := r.db.GetContext(ctx, &count, query, args...)
	_ = argIndex // Mark as used to satisfy linter
	return count, err
}

// Delete soft deletes a user (sets is_active to false)
func (r *UserRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE users SET is_active = false, updated_at = NOW() WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

// ==================== School Boundary Helper Methods ====================

// GetUserSchoolID retrieves the school_id for a given user
func (r *UserRepository) GetUserSchoolID(ctx context.Context, userID string) (*string, error) {
	query := `SELECT school_id FROM users WHERE id = $1`

	var schoolID sql.NullString
	err := r.db.GetContext(ctx, &schoolID, query, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	if schoolID.Valid {
		return &schoolID.String, nil
	}

	return nil, nil
}

// GetUsersBySchool retrieves all users belonging to a specific school
func (r *UserRepository) GetUsersBySchool(ctx context.Context, schoolID string) ([]*domain.User, error) {
	query := `
		SELECT id, email, password_hash, name, role_id, school_id, is_active, 
		       failed_login_attempts, locked_until, created_at, updated_at, created_by, updated_by
		FROM users WHERE school_id = $1
	`

	rows, err := r.db.QueryContext(ctx, query, schoolID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		var user domain.User
		var userSchoolID sql.NullString
		var lockedUntil sql.NullTime
		var createdBy sql.NullString
		var updatedBy sql.NullString

		err := rows.Scan(
			&user.ID, &user.Email, &user.PasswordHash, &user.Name, &user.RoleID, &userSchoolID,
			&user.IsActive, &user.FailedLoginAttempts, &lockedUntil,
			&user.CreatedAt, &user.UpdatedAt, &createdBy, &updatedBy,
		)
		if err != nil {
			return nil, err
		}

		if userSchoolID.Valid {
			user.SchoolID = &userSchoolID.String
		}
		if lockedUntil.Valid {
			user.LockedUntil = &lockedUntil.Time
		}
		if createdBy.Valid {
			user.CreatedBy = &createdBy.String
		}
		if updatedBy.Valid {
			user.UpdatedBy = &updatedBy.String
		}

		users = append(users, &user)
	}

	return users, nil
}

// ListUsersBySchool retrieves users belonging to a specific school with pagination
func (r *UserRepository) ListUsersBySchool(ctx context.Context, schoolID string, limit, offset int) ([]*domain.User, error) {
	query := `
		SELECT id, email, password_hash, name, role_id, school_id, is_active, 
		       failed_login_attempts, locked_until, created_at, updated_at, created_by, updated_by
		FROM users WHERE school_id = $1
		ORDER BY created_at DESC
	`

	args := []interface{}{schoolID}
	argIndex := 2

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

	var users []*domain.User
	for rows.Next() {
		var user domain.User
		var userSchoolID sql.NullString
		var lockedUntil sql.NullTime
		var createdBy sql.NullString
		var updatedBy sql.NullString

		err := rows.Scan(
			&user.ID, &user.Email, &user.PasswordHash, &user.Name, &user.RoleID, &userSchoolID,
			&user.IsActive, &user.FailedLoginAttempts, &lockedUntil,
			&user.CreatedAt, &user.UpdatedAt, &createdBy, &updatedBy,
		)
		if err != nil {
			return nil, err
		}

		if userSchoolID.Valid {
			user.SchoolID = &userSchoolID.String
		}
		if lockedUntil.Valid {
			user.LockedUntil = &lockedUntil.Time
		}
		if createdBy.Valid {
			user.CreatedBy = &createdBy.String
		}
		if updatedBy.Valid {
			user.UpdatedBy = &updatedBy.String
		}

		users = append(users, &user)
	}

	return users, nil
}

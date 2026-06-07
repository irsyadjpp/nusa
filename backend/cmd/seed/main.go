package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/nusa/backend/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://nusa_user:nusa_password@localhost:5432/nusa_db?sslmode=disable"
	}

	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	ctx := context.Background()

	// Seed roles
	fmt.Println("Seeding roles...")
	if err := seedRoles(ctx, db); err != nil {
		log.Fatalf("Failed to seed roles: %v", err)
	}

	// Seed permissions
	fmt.Println("Seeding permissions...")
	if err := seedPermissions(ctx, db); err != nil {
		log.Fatalf("Failed to seed permissions: %v", err)
	}

	// Seed default school
	fmt.Println("Seeding default school...")
	schoolID, err := seedDefaultSchool(ctx, db)
	if err != nil {
		log.Fatalf("Failed to seed default school: %v", err)
	}

	// Seed admin user
	fmt.Println("Seeding admin user...")
	if err := seedAdminUser(ctx, db, schoolID); err != nil {
		log.Fatalf("Failed to seed admin user: %v", err)
	}

	fmt.Println("Seeding completed successfully!")
}

func seedRoles(ctx context.Context, db *sqlx.DB) error {
	roles := []struct {
		name        string
		description *string
		isActive    bool
	}{
		{domain.RoleSystemAdmin, strPtr("Platform administrator with full system access"), true},
		{domain.RoleSchoolAdmin, strPtr("School administrator with school-level access"), true},
		{domain.RoleTeacher, strPtr("Teacher with classroom-level access"), true},
	}

	for _, role := range roles {
		query := `
			INSERT INTO roles (id, name, description, is_active, created_at, updated_at)
			VALUES ($1, $2, $3, $4, NOW(), NOW())
			ON CONFLICT (name) DO UPDATE SET
				description = EXCLUDED.description,
				is_active = EXCLUDED.is_active,
				updated_at = NOW()
			RETURNING id
		`
		_, err := db.ExecContext(ctx, query, domain.NewID(), role.name, role.description, role.isActive)
		if err != nil {
			return fmt.Errorf("failed to insert role %s: %w", role.name, err)
		}
	}

	return nil
}

func seedPermissions(ctx context.Context, db *sqlx.DB) error {
	rolePermissions := domain.GetRolePermissions()

	for roleName, permissions := range rolePermissions {
		// Get role ID
		var roleID string
		err := db.GetContext(ctx, &roleID, "SELECT id FROM roles WHERE name = $1", roleName)
		if err != nil {
			return fmt.Errorf("failed to get role ID for %s: %w", roleName, err)
		}

		// Add permissions
		for _, perm := range permissions {
			query := `
				INSERT INTO permissions (id, role_id, resource, action, created_at)
				VALUES ($1, $2, $3, $4, NOW())
				ON CONFLICT (role_id, resource, action) DO NOTHING
			`
			_, err := db.ExecContext(ctx, query, domain.NewID(), roleID, permResource(perm), permAction(perm))
			if err != nil {
				return fmt.Errorf("failed to insert permission %s for role %s: %w", perm, roleName, err)
			}
		}
	}

	return nil
}

func seedDefaultSchool(ctx context.Context, db *sqlx.DB) (string, error) {
	schoolID := domain.NewID()
	query := `
		INSERT INTO schools (id, name, code, is_active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW())
		ON CONFLICT (code) DO UPDATE SET
			name = EXCLUDED.name,
			is_active = EXCLUDED.is_active,
			updated_at = NOW()
	`

	_, err := db.ExecContext(ctx, query, schoolID, "Default School", "SCH-001", true)
	if err != nil {
		return "", fmt.Errorf("failed to insert default school: %w", err)
	}

	// Get the school ID (whether it was inserted or already existed)
	var id string
	err = db.GetContext(ctx, &id, "SELECT id FROM schools WHERE code = $1", "SCH-001")
	if err != nil {
		return "", fmt.Errorf("failed to get default school ID: %w", err)
	}

	return id, nil
}

func seedAdminUser(ctx context.Context, db *sqlx.DB, schoolID string) error {
	// Get SYSTEM_ADMIN role ID
	var roleID string
	err := db.GetContext(ctx, &roleID, "SELECT id FROM roles WHERE name = $1", domain.RoleSystemAdmin)
	if err != nil {
		return fmt.Errorf("failed to get SYSTEM_ADMIN role ID: %w", err)
	}

	// Hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// Insert admin user
	query := `
		INSERT INTO users (id, email, password_hash, name, role_id, school_id, is_active, failed_login_attempts, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW(), NOW())
		ON CONFLICT (email) DO UPDATE SET
			password_hash = EXCLUDED.password_hash,
			name = EXCLUDED.name,
			role_id = EXCLUDED.role_id,
			school_id = EXCLUDED.school_id,
			is_active = EXCLUDED.is_active,
			updated_at = NOW()
	`

	adminEmail := "admin@nusa.local"
	_, err = db.ExecContext(ctx, query, domain.NewID(), adminEmail, string(passwordHash), "System Administrator", roleID, &schoolID, true, 0)
	if err != nil {
		return fmt.Errorf("failed to insert admin user: %w", err)
	}

	fmt.Printf("Admin user created: %s / admin123\n", adminEmail)
	return nil
}

func strPtr(s string) *string {
	return &s
}

func permResource(perm string) string {
	// perm format is "resource:action"
	for i, char := range perm {
		if char == ':' {
			return perm[:i]
		}
	}
	return perm
}

func permAction(perm string) string {
	// perm format is "resource:action"
	for i, char := range perm {
		if char == ':' {
			return perm[i+1:]
		}
	}
	return perm
}

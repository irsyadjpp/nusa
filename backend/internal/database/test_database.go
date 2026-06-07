package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// NewTestDatabase creates a new test database connection
func NewTestDatabase(host, port, user, password, dbname, sslmode string) (*Database, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to test database: %w", err)
	}

	// Use smaller connection pool for tests
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(5 * time.Minute)

	return &Database{DB: db}, nil
}

// SetupTestDatabase creates and initializes the test database schema
func SetupTestDatabase(db *Database) error {
	ctx := context.Background()
	
	// Test database schema setup
	schema := `
	-- Enable UUID extension
	CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
	
	-- Users table
	CREATE TABLE IF NOT EXISTS users (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		email VARCHAR(255) UNIQUE NOT NULL,
		password_hash VARCHAR(255) NOT NULL,
		name VARCHAR(100) NOT NULL,
		role_id UUID NOT NULL,
		school_id UUID,
		phone VARCHAR(20),
		address VARCHAR(500),
		status VARCHAR(20) NOT NULL DEFAULT 'ACTIVE',
		created_by UUID NOT NULL,
		updated_by UUID NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP WITH TIME ZONE
	);
	
	-- Schools table
	CREATE TABLE IF NOT EXISTS schools (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		code VARCHAR(20) UNIQUE NOT NULL,
		name VARCHAR(100) NOT NULL,
		address VARCHAR(500),
		phone VARCHAR(20),
		email VARCHAR(255),
		status VARCHAR(20) NOT NULL DEFAULT 'ACTIVE',
		created_by UUID NOT NULL,
		updated_by UUID NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP WITH TIME ZONE
	);
	
	-- Roles table
	CREATE TABLE IF NOT EXISTS roles (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		name VARCHAR(50) UNIQUE NOT NULL,
		is_active BOOLEAN NOT NULL DEFAULT true,
		created_by UUID NOT NULL,
		updated_by UUID NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP WITH TIME ZONE
	);
	
	-- Refresh tokens table
	CREATE TABLE IF NOT EXISTS refresh_tokens (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
		token VARCHAR(500) UNIQUE NOT NULL,
		expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		revoked_at TIMESTAMP WITH TIME ZONE
	);
	
	-- Create indexes
	CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
	CREATE INDEX IF NOT EXISTS idx_users_role_id ON users(role_id);
	CREATE INDEX IF NOT EXISTS idx_users_school_id ON users(school_id);
	CREATE INDEX IF NOT EXISTS idx_users_status ON users(status);
	CREATE INDEX IF NOT EXISTS idx_schools_code ON schools(code);
	CREATE INDEX IF NOT EXISTS idx_schools_status ON schools(status);
	CREATE INDEX IF NOT EXISTS idx_roles_name ON roles(name);
	CREATE INDEX IF NOT EXISTS idx_refresh_tokens_user_id ON refresh_tokens(user_id);
	CREATE INDEX IF NOT EXISTS idx_refresh_tokens_token ON refresh_tokens(token);
	`

	_, err := db.DB.ExecContext(ctx, schema)
	if err != nil {
		return fmt.Errorf("failed to create test database schema: %w", err)
	}

	return nil
}

// CleanupTestDatabase drops all tables in the test database
func CleanupTestDatabase(db *Database) error {
	ctx := context.Background()
	
	// Drop tables in correct order (reverse of creation due to foreign keys)
	tables := []string{
		"refresh_tokens",
		"users",
		"schools",
		"roles",
	}

	for _, table := range tables {
		_, err := db.DB.ExecContext(ctx, fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE", table))
		if err != nil {
			return fmt.Errorf("failed to drop table %s: %w", table, err)
		}
	}

	return nil
}

// TruncateTestDatabase truncates all tables in the test database
func TruncateTestDatabase(db *Database) error {
	ctx := context.Background()
	
	// Truncate tables in correct order (reverse of creation due to foreign keys)
	tables := []string{
		"refresh_tokens",
		"users",
		"schools",
		"roles",
	}

	for _, table := range tables {
		_, err := db.DB.ExecContext(ctx, fmt.Sprintf("TRUNCATE TABLE %s CASCADE", table))
		if err != nil {
			return fmt.Errorf("failed to truncate table %s: %w", table, err)
		}
	}

	return nil
}

// SeedTestDatabase seeds the test database with initial data
func SeedTestDatabase(db *Database) error {
	ctx := context.Background()
	
	// Insert system roles
	roles := []struct {
		id        string
		name      string
		isActive  bool
		createdBy string
		updatedBy string
	}{
		{"00000000-0000-0000-0000-000000000001", "SYSTEM_ADMIN", true, "00000000-0000-0000-0000-000000000001", "00000000-0000-0000-0000-000000000001"},
		{"00000000-0000-0000-0000-000000000002", "SCHOOL_ADMIN", true, "00000000-0000-0000-0000-000000000001", "00000000-0000-0000-0000-000000000001"},
		{"00000000-0000-0000-0000-000000000003", "TEACHER", true, "00000000-0000-0000-0000-000000000001", "00000000-0000-0000-0000-000000000001"},
	}

	for _, role := range roles {
		_, err := db.DB.ExecContext(ctx, `
			INSERT INTO roles (id, name, is_active, created_by, updated_by)
			VALUES ($1, $2, $3, $4, $5)
			ON CONFLICT (name) DO NOTHING
		`, role.id, role.name, role.isActive, role.createdBy, role.updatedBy)
		if err != nil {
			return fmt.Errorf("failed to seed role %s: %w", role.name, err)
		}
	}

	return nil
}

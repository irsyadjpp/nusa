package repository

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/nusa/backend/internal/config"
	"github.com/nusa/backend/internal/db"
	"github.com/stretchr/testify/require"
)

// TestDB holds the test database connection and cleanup functions
type TestDB struct {
	Pool    *sqlx.DB
	Cleanup func()
}

// SetupTestDB creates a test database connection and returns a TestDB struct
// with a cleanup function to close the connection
func SetupTestDB(t *testing.T) *TestDB {
	t.Helper()

	// Load test configuration
	cfg := config.LoadTestConfig()

	// Override with environment variables if set
	if host := os.Getenv("TEST_DB_HOST"); host != "" {
		cfg.Database.Host = host
	}
	if port := os.Getenv("TEST_DB_PORT"); port != "" {
		cfg.Database.Port = port
	}
	if user := os.Getenv("TEST_DB_USER"); user != "" {
		cfg.Database.User = user
	}
	if password := os.Getenv("TEST_DB_PASSWORD"); password != "" {
		cfg.Database.Password = password
	}
	if dbname := os.Getenv("TEST_DB_NAME"); dbname != "" {
		cfg.Database.DBName = dbname
	}

	// Create PostgreSQL connection
	pg, err := db.NewPostgres(&cfg.Database)
	if err != nil {
		t.Skipf("Skipping test - database connection failed: %v", err)
		return nil
	}

	// Get sqlx.DB instance
	sqlxDB, err := pg.GetSQLXDB()
	require.NoError(t, err, "Failed to get sqlx.DB instance")

	// Verify connection works
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = sqlxDB.PingContext(ctx)
	require.NoError(t, err, "Failed to ping database")

	return &TestDB{
		Pool: sqlxDB,
		Cleanup: func() {
			pg.Close()
		},
	}
}

// CleanupTestData cleans up test data from all relevant tables
// This should be called after each test to ensure data isolation
func CleanupTestData(t *testing.T, db *sqlx.DB) {
	t.Helper()

	ctx := context.Background()

	// Delete from child tables first (respecting foreign key constraints)
	tables := []string{
		// Evidence and assessment related
		"evidence",
		"evaluation",
		"assessments",
		"exam_results",
		"exams",
		"assignment_submissions",
		"assignments",
		"attendance_records",

		// Learning planning related
		"tp",
		"tp_sets",
		"atp",
		"modul_ajar",

		// Academic structure
		"class_enrollments",
		"classes",
		"schedules",
		"subjects",
		"subject_categories",
		"semesters",
		"academic_years",

		// Curriculum and outcomes
		"cp_alignments",
		"graduate_profile_dimensions",
		"cp", // Learning outcomes
		"curricula",

		// System and configuration
		"system_configurations",
		"refresh_tokens",
		"notifications",
		"messages",
		"announcements",

		// Core entities
		"users",
		"roles",
		"schools",
	}

	for _, table := range tables {
		_, err := db.ExecContext(ctx, fmt.Sprintf("DELETE FROM %s WHERE true", table))
		if err != nil {
			// Table might not exist, which is okay for some tests
			// Log but don't fail
			t.Logf("Warning: failed to clean up table %s: %v", table, err)
		}
	}
}

// CreateTestRole creates a test role in the database
func CreateTestRole(t *testing.T, db *sqlx.DB, name string) string {
	t.Helper()

	id := fmt.Sprintf("role-%s-%d", name, time.Now().UnixNano())
	query := `
		INSERT INTO roles (id, name, description, created_by, updated_by)
		VALUES ($1, $2, $3, $1, $1)
	`
	_, err := db.Exec(query, id, name, fmt.Sprintf("Test role %s", name))
	require.NoError(t, err, "Failed to create test role")

	return id
}

// CreateTestSchool creates a test school in the database
func CreateTestSchool(t *testing.T, db *sqlx.DB, name string) string {
	t.Helper()

	id := fmt.Sprintf("school-%s-%d", name, time.Now().UnixNano())
	code := fmt.Sprintf("SCH-%d", time.Now().UnixNano())
	query := `
		INSERT INTO schools (id, code, name, address, created_by, updated_by)
		VALUES ($1, $2, $3, $4, $1, $1)
	`
	_, err := db.Exec(query, id, code, name, "Test Address")
	require.NoError(t, err, "Failed to create test school")

	return id
}

// CreateTestUser creates a test user in the database
func CreateTestUser(t *testing.T, db *sqlx.DB, email string, roleID string, schoolID *string) string {
	t.Helper()

	id := fmt.Sprintf("user-%d", time.Now().UnixNano())
	query := `
		INSERT INTO users (id, email, password_hash, name, role_id, school_id, is_active, created_by, updated_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $1, $1)
	`
	_, err := db.Exec(query, id, email, "hashed_password", "Test User", roleID, schoolID, true)
	require.NoError(t, err, "Failed to create test user")

	return id
}

// CreateTestAcademicYear creates a test academic year
func CreateTestAcademicYear(t *testing.T, db *sqlx.DB, schoolID, year string) string {
	t.Helper()

	id := fmt.Sprintf("ay-%s-%d", year, time.Now().UnixNano())
	query := `
		INSERT INTO academic_years (id, school_id, year, start_date, end_date, is_current, created_by, updated_by)
		VALUES ($1, $2, $3, $4, $5, $6, $1, $1)
	`
	startDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)
	_, err := db.Exec(query, id, schoolID, year, startDate, endDate, false)
	require.NoError(t, err, "Failed to create test academic year")

	return id
}

// CreateTestSemester creates a test semester
func CreateTestSemester(t *testing.T, db *sqlx.DB, academicYearID, name string) string {
	t.Helper()

	id := fmt.Sprintf("sem-%s-%d", name, time.Now().UnixNano())
	query := `
		INSERT INTO semesters (id, academic_year_id, name, start_date, end_date, is_active, created_by, updated_by)
		VALUES ($1, $2, $3, $4, $5, $6, $1, $1)
	`
	startDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2024, 6, 30, 0, 0, 0, 0, time.UTC)
	_, err := db.Exec(query, id, academicYearID, name, startDate, endDate, true)
	require.NoError(t, err, "Failed to create test semester")

	return id
}

// CreateTestCurriculum creates a test curriculum
func CreateTestCurriculum(t *testing.T, db *sqlx.DB, schoolID, name string) string {
	t.Helper()

	id := fmt.Sprintf("cur-%s-%d", name, time.Now().UnixNano())
	query := `
		INSERT INTO curricula (id, school_id, name, version, created_by, updated_by)
		VALUES ($1, $2, $3, $4, $1, $1)
	`
	_, err := db.Exec(query, id, schoolID, name, 1)
	require.NoError(t, err, "Failed to create test curriculum")

	return id
}

// StringPtr returns a pointer to a string
func StringPtr(s string) *string {
	return &s
}

// IntPtr returns a pointer to an int
func IntPtr(i int) *int {
	return &i
}

// BoolPtr returns a pointer to a bool
func BoolPtr(b bool) *bool {
	return &b
}

// TimePtr returns a pointer to a time.Time
func TimePtr(t time.Time) *time.Time {
	return &t
}

// NullString returns a sql.NullString
func NullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: true}
}

// NullInt returns a sql.NullInt64
func NullInt(i int64) sql.NullInt64 {
	return sql.NullInt64{Int64: i, Valid: true}
}

// NullTime returns a sql.NullTime
func NullTime(t time.Time) sql.NullTime {
	return sql.NullTime{Time: t, Valid: true}
}

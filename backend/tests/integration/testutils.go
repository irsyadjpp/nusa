package integration

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/nusa/backend/internal/config"
	"github.com/nusa/backend/internal/database"
	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/repository"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

// TestDB holds the test database connection and repositories
type TestDB struct {
	DB               *sqlx.DB
	UserRepo         *repository.UserRepository
	RoleRepo         *repository.RoleRepository
	SchoolRepo       *repository.SchoolRepository
	RefreshTokenRepo *repository.RefreshTokenRepository
}

// SetupTestDB creates and initializes a test database connection
func SetupTestDB(t require.TestingT) *TestDB {
	config.SetTestEnv()
	defer config.UnsetTestEnv()

	testConfig := config.LoadTestConfig()

	db, err := database.NewTestDatabase(
		testConfig.Database.Host,
		testConfig.Database.Port,
		testConfig.Database.User,
		testConfig.Database.Password,
		testConfig.Database.DBName,
		testConfig.Database.SSLMode,
	)
	if err != nil {
		require.NoError(t, err)
		return nil
	}

	// Setup test database schema
	err = database.SetupTestDatabase(db)
	if err != nil {
		require.NoError(t, err)
		return nil
	}

	// Seed test database with initial data
	err = database.SeedTestDatabase(db)
	if err != nil {
		require.NoError(t, err)
		return nil
	}

	return &TestDB{
		DB:               db.DB,
		UserRepo:         repository.NewUserRepository(db.DB),
		RoleRepo:         repository.NewRoleRepository(db.DB),
		SchoolRepo:       repository.NewSchoolRepository(db.DB),
		RefreshTokenRepo: repository.NewRefreshTokenRepository(db.DB),
	}
}

// TeardownTestDB truncates all tables and closes the connection
func TeardownTestDB(t require.TestingT, testDB *TestDB) {
	if testDB == nil || testDB.DB == nil {
		return
	}

	err := database.TruncateTestDatabase(&database.Database{DB: testDB.DB})
	require.NoError(t, err)

	testDB.DB.Close()
}

// CreateTestUser creates a test user with a hashed password
func CreateTestUser(t require.TestingT, ctx context.Context, userRepo *repository.UserRepository, email, password, name, roleID string, schoolID *string) *domain.User {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	require.NoError(t, err)

	user := &domain.User{
		ID:                  domain.NewID(),
		Email:               email,
		PasswordHash:        string(passwordHash),
		Name:                name,
		RoleID:              roleID,
		SchoolID:            schoolID,
		IsActive:            true,
		FailedLoginAttempts: 0,
		CreatedBy:           &roleID,
		UpdatedBy:           &roleID,
	}

	err = userRepo.Create(ctx, user)
	require.NoError(t, err)

	return user
}

// CreateTestRefreshToken creates a test refresh token
func CreateTestRefreshToken(t require.TestingT, ctx context.Context, tokenRepo *repository.RefreshTokenRepository, userID, token string, expiresAt time.Time) error {
	return tokenRepo.Create(ctx, userID, token, expiresAt, nil, &userID)
}

// CreateTestSchool creates a test school
func CreateTestSchool(t require.TestingT, ctx context.Context, schoolRepo *repository.SchoolRepository, name, code string) *domain.School {
	adminID := "00000000-0000-0000-0000-000000000001"

	school := &domain.School{
		ID:        domain.NewID(),
		Name:      name,
		Code:      code,
		IsActive:  true,
		CreatedBy: &adminID,
		UpdatedBy: &adminID,
	}

	err := schoolRepo.Create(ctx, school)
	require.NoError(t, err)

	return school
}

// CreateTestRole creates a test role
func CreateTestRole(t require.TestingT, ctx context.Context, roleRepo *repository.RoleRepository, name string, description *string) *domain.Role {
	role := &domain.Role{
		ID:          domain.NewID(),
		Name:        name,
		Description: description,
		IsActive:    true,
	}

	err := roleRepo.Create(ctx, role)
	require.NoError(t, err)

	return role
}

// WithTransaction runs a function within a database transaction and rolls it back
func WithTransaction(t require.TestingT, db *sqlx.DB, fn func(ctx context.Context, tx *sqlx.Tx) error) {
	ctx := context.Background()
	tx, err := db.BeginTxx(ctx, nil)
	require.NoError(t, err)

	defer func() {
		if err := recover(); err != nil {
			_ = tx.Rollback()
			panic(err)
		}
	}()

	err = fn(ctx, tx)
	if err != nil {
		_ = tx.Rollback()
		require.NoError(t, err)
		return
	}

	err = tx.Rollback()
	require.NoError(t, err)
}

// Wait is a helper for waiting in tests
func Wait(duration time.Duration) {
	time.Sleep(duration)
}

// AssertEqualTime compares two times with a small tolerance
func AssertEqualTime(t require.TestingT, expected, actual time.Time, tolerance time.Duration) {
	diff := expected.Sub(actual)
	if diff < 0 {
		diff = -diff
	}
	require.True(t, diff <= tolerance, "time difference %v exceeds tolerance %v", diff, tolerance)
}

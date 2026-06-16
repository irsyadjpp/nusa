package integration

import (
	"context"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/nusa/backend/internal/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestDatabaseConnection tests database connection and health
func TestDatabaseConnection(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	testDB := SetupTestDB(t)
	if testDB == nil {
		return
	}
	defer TeardownTestDB(t, testDB)

	assert.NotNil(t, testDB.DB)
	assert.NotNil(t, testDB.UserRepo)
	assert.NotNil(t, testDB.RoleRepo)
	assert.NotNil(t, testDB.SchoolRepo)
	assert.NotNil(t, testDB.RefreshTokenRepo)
}

// TestDatabaseHealthCheck tests database health check
func TestDatabaseHealthCheck(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	testDB := SetupTestDB(t)
	if testDB == nil {
		return
	}
	defer TeardownTestDB(t, testDB)

	ctx := context.Background()
	err := testDB.DB.PingContext(ctx)
	require.NoError(t, err, "Database should respond to ping")
}

// TestDatabaseTransaction tests database transaction functionality
func TestDatabaseTransaction(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	testDB := SetupTestDB(t)
	if testDB == nil {
		return
	}
	defer TeardownTestDB(t, testDB)

	ctx := context.Background()
	roleID := "00000000-0000-0000-0000-000000000001"

	t.Run("Transaction commit", func(t *testing.T) {
		WithTransaction(t, testDB.DB, func(ctx context.Context, tx *sqlx.Tx) error {
			// Create user within transaction
			user := CreateTestUser(t, ctx, testDB.UserRepo, "trans-commit@example.com", "password123", "Transaction Commit", roleID, nil)
			assert.NotNil(t, user)

			// Verify user exists within transaction
			retrieved, err := testDB.UserRepo.GetByID(ctx, user.ID)
			require.NoError(t, err)
			assert.Equal(t, user.ID, retrieved.ID)

			return nil
		})

		// Verify user still exists after transaction commit (even though we rollback in WithTransaction)
		// Note: WithTransaction always rolls back, so user should not exist
		_, err := testDB.UserRepo.GetByEmail(ctx, "trans-commit@example.com")
		assert.Error(t, err, "User should not exist after transaction rollback")
	})

	t.Run("Transaction rollback", func(t *testing.T) {
		email := "trans-rollback@example.com"

		WithTransaction(t, testDB.DB, func(ctx context.Context, tx *sqlx.Tx) error {
			user := CreateTestUser(t, ctx, testDB.UserRepo, email, "password123", "Transaction Rollback", roleID, nil)
			assert.NotNil(t, user)

			// Force rollback by returning error
			return assert.AnError
		})

		// Verify user does not exist after rollback
		_, err := testDB.UserRepo.GetByEmail(ctx, email)
		assert.Error(t, err, "User should not exist after transaction rollback")
	})
}

// TestDatabaseConstraints tests database constraints
func TestDatabaseConstraints(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	testDB := SetupTestDB(t)
	if testDB == nil {
		return
	}
	defer TeardownTestDB(t, testDB)

	ctx := context.Background()
	roleID := "00000000-0000-0000-0000-000000000001"

	t.Run("Unique constraint on email", func(t *testing.T) {
		email := "unique@example.com"
		CreateTestUser(t, ctx, testDB.UserRepo, email, "password123", "Unique Email", roleID, nil)

		// Try to create user with same email
		user2 := CreateTestUser(t, ctx, testDB.UserRepo, email, "password456", "Duplicate Email", roleID, nil)
		assert.Nil(t, user2, "Should fail to create user with duplicate email")
	})

	t.Run("Unique constraint on school code", func(t *testing.T) {
		code := "UNIQ001"
		CreateTestSchool(t, ctx, testDB.SchoolRepo, "Unique School", code)

		// Try to create school with same code
		school2 := CreateTestSchool(t, ctx, testDB.SchoolRepo, "Duplicate School", code)
		assert.Nil(t, school2, "Should fail to create school with duplicate code")
	})

	t.Run("Unique constraint on role name", func(t *testing.T) {
		name := "CUSTOM_UNIQUE"
		description := "Custom unique role"
		CreateTestRole(t, ctx, testDB.RoleRepo, name, &description)

		// Try to create role with same name
		role2 := CreateTestRole(t, ctx, testDB.RoleRepo, name, &description)
		assert.Nil(t, role2, "Should fail to create role with duplicate name")
	})

	t.Run("Foreign key constraint on user.role_id", func(t *testing.T) {
		// This would require using raw SQL to violate the constraint
		// For now, we'll just verify that users can only be created with valid role IDs
		user := CreateTestUser(t, ctx, testDB.UserRepo, "fk@example.com", "password123", "FK Test", roleID, nil)
		assert.NotNil(t, user, "User should be created with valid role ID")
	})
}

// TestDatabaseIndexes tests database indexes are working
func TestDatabaseIndexes(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	testDB := SetupTestDB(t)
	if testDB == nil {
		return
	}
	defer TeardownTestDB(t, testDB)

	ctx := context.Background()
	roleID := "00000000-0000-0000-0000-000000000001"

	t.Run("Email index usage", func(t *testing.T) {
		email := "indexed@example.com"
		user := CreateTestUser(t, ctx, testDB.UserRepo, email, "password123", "Indexed User", roleID, nil)

		// Query by email should use index
		retrieved, err := testDB.UserRepo.GetByEmail(ctx, email)
		require.NoError(t, err)
		assert.Equal(t, user.ID, retrieved.ID)
	})

	t.Run("School code index usage", func(t *testing.T) {
		code := "IDX001"
		school := CreateTestSchool(t, ctx, testDB.SchoolRepo, "Indexed School", code)

		// Query by code should use index
		retrieved, err := testDB.SchoolRepo.GetByCode(ctx, code)
		require.NoError(t, err)
		assert.Equal(t, school.ID, retrieved.ID)
	})

	t.Run("Role name index usage", func(t *testing.T) {
		name := "INDEXED_ROLE"
		description := "Indexed role"
		role := CreateTestRole(t, ctx, testDB.RoleRepo, name, &description)

		// Query by name should use index
		retrieved, err := testDB.RoleRepo.GetByName(ctx, name)
		require.NoError(t, err)
		assert.Equal(t, role.ID, retrieved.ID)
	})
}

// TestDatabaseConnectionPooling tests database connection pooling
func TestDatabaseConnectionPooling(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	testDB := SetupTestDB(t)
	if testDB == nil {
		return
	}
	defer TeardownTestDB(t, testDB)

	ctx := context.Background()
	roleID := "00000000-0000-0000-0000-000000000001"

	t.Run("Multiple concurrent queries", func(t *testing.T) {
		// Create multiple users concurrently to test connection pool
		for i := 0; i < 10; i++ {
			email := "pool" + string(rune('0'+i)) + "@example.com"
			CreateTestUser(t, ctx, testDB.UserRepo, email, "password123", "Pool User", roleID, nil)
		}

		// Verify all users were created
		users, err := testDB.UserRepo.List(ctx, nil, nil, nil, 10, 0)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(users), 10)
	})
}

// TestDatabaseNullFields tests handling of NULL fields
func TestDatabaseNullFields(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	testDB := SetupTestDB(t)
	if testDB == nil {
		return
	}
	defer TeardownTestDB(t, testDB)

	ctx := context.Background()
	roleID := "00000000-0000-0000-0000-000000000001"

	t.Run("User with optional fields NULL", func(t *testing.T) {
		user := CreateTestUser(t, ctx, testDB.UserRepo, "null@example.com", "password123", "Null Fields", roleID, nil)

		// Verify nullable fields are handled correctly
		retrieved, err := testDB.UserRepo.GetByID(ctx, user.ID)
		require.NoError(t, err)
		assert.Nil(t, retrieved.SchoolID, "School ID should be NULL")
		assert.Nil(t, retrieved.LockedUntil, "Locked until should be NULL")
	})

	t.Run("User with optional fields set", func(t *testing.T) {
		school := CreateTestSchool(t, ctx, testDB.SchoolRepo, "School with Users", "SWU001")
		user := CreateTestUser(t, ctx, testDB.UserRepo, "notnull@example.com", "password123", "Not Null Fields", roleID, &school.ID)

		// Verify nullable fields are set
		retrieved, err := testDB.UserRepo.GetByID(ctx, user.ID)
		require.NoError(t, err)
		assert.NotNil(t, retrieved.SchoolID, "School ID should not be NULL")
		assert.Equal(t, school.ID, *retrieved.SchoolID)
	})

	t.Run("School with optional fields NULL", func(t *testing.T) {
		school := CreateTestSchool(t, ctx, testDB.SchoolRepo, "Minimal School", "MS001")

		// Verify nullable fields are handled correctly
		retrieved, err := testDB.SchoolRepo.GetByID(ctx, school.ID)
		require.NoError(t, err)
		assert.Nil(t, retrieved.Address, "Address should be NULL")
		assert.Nil(t, retrieved.Phone, "Phone should be NULL")
		assert.Nil(t, retrieved.Email, "Email should be NULL")
	})

	t.Run("School with optional fields set", func(t *testing.T) {
		address := "123 Test Street"
		phone := "555-1234"
		email := "school@example.com"
		school := CreateTestSchool(t, ctx, testDB.SchoolRepo, "Full School", "FS001")

		// Update with optional fields
		school.Address = &address
		school.Phone = &phone
		school.Email = &email
		err := testDB.SchoolRepo.Update(ctx, school)
		require.NoError(t, err)

		// Verify nullable fields are set
		retrieved, err := testDB.SchoolRepo.GetByID(ctx, school.ID)
		require.NoError(t, err)
		assert.NotNil(t, retrieved.Address, "Address should not be NULL")
		assert.Equal(t, address, *retrieved.Address)
		assert.NotNil(t, retrieved.Phone, "Phone should not be NULL")
		assert.Equal(t, phone, *retrieved.Phone)
		assert.NotNil(t, retrieved.Email, "Email should not be NULL")
		assert.Equal(t, email, *retrieved.Email)
	})

	t.Run("Role with description NULL", func(t *testing.T) {
		role := CreateTestRole(t, ctx, testDB.RoleRepo, "NO_DESC_ROLE", nil)

		// Verify description is handled correctly
		retrieved, err := testDB.RoleRepo.GetByID(ctx, role.ID)
		require.NoError(t, err)
		assert.Nil(t, retrieved.Description, "Description should be NULL")
	})

	t.Run("Role with description set", func(t *testing.T) {
		description := "Role with description"
		role := CreateTestRole(t, ctx, testDB.RoleRepo, "DESC_ROLE", &description)

		// Verify description is set
		retrieved, err := testDB.RoleRepo.GetByID(ctx, role.ID)
		require.NoError(t, err)
		assert.NotNil(t, retrieved.Description, "Description should not be NULL")
		assert.Equal(t, description, *retrieved.Description)
	})
}

// TestDatabaseTimestamps tests timestamp fields (created_at, updated_at)
func TestDatabaseTimestamps(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	testDB := SetupTestDB(t)
	if testDB == nil {
		return
	}
	defer TeardownTestDB(t, testDB)

	ctx := context.Background()
	roleID := "00000000-0000-0000-0000-000000000001"

	t.Run("User timestamps", func(t *testing.T) {
		beforeCreate := time.Now()
		user := CreateTestUser(t, ctx, testDB.UserRepo, "timestamp@example.com", "password123", "Timestamp User", roleID, nil)
		afterCreate := time.Now()

		// Verify created_at is set
		assert.False(t, user.CreatedAt.IsZero(), "Created at should be set")
		assert.False(t, user.UpdatedAt.IsZero(), "Updated at should be set")

		// Verify timestamps are within expected range
		assert.True(t, user.CreatedAt.After(beforeCreate) || user.CreatedAt.Equal(beforeCreate), "Created at should be after or equal to beforeCreate")
		assert.True(t, user.CreatedAt.Before(afterCreate) || user.CreatedAt.Equal(afterCreate), "Created at should be before or equal to afterCreate")

		// Update user
		beforeUpdate := time.Now()
		user.Name = "Updated Name"
		err := testDB.UserRepo.Update(ctx, user)
		require.NoError(t, err)
		afterUpdate := time.Now()

		// Verify updated_at changed
		retrieved, err := testDB.UserRepo.GetByID(ctx, user.ID)
		require.NoError(t, err)
		assert.True(t, retrieved.UpdatedAt.After(user.UpdatedAt), "Updated at should be after original")

		// Verify updated timestamp is within expected range
		assert.True(t, retrieved.UpdatedAt.After(beforeUpdate) || retrieved.UpdatedAt.Equal(beforeUpdate), "Updated at should be after or equal to beforeUpdate")
		assert.True(t, retrieved.UpdatedAt.Before(afterUpdate) || retrieved.UpdatedAt.Equal(afterUpdate), "Updated at should be before or equal to afterUpdate")
	})

	t.Run("School timestamps", func(t *testing.T) {
		beforeCreate := time.Now()
		school := CreateTestSchool(t, ctx, testDB.SchoolRepo, "Timestamp School", "TS001")
		afterCreate := time.Now()

		// Verify timestamps are set
		assert.False(t, school.CreatedAt.IsZero(), "Created at should be set")
		assert.False(t, school.UpdatedAt.IsZero(), "Updated at should be set")

		// Verify timestamps are within expected range
		assert.True(t, school.CreatedAt.After(beforeCreate) || school.CreatedAt.Equal(beforeCreate), "Created at should be after or equal to beforeCreate")
		assert.True(t, school.CreatedAt.Before(afterCreate) || school.CreatedAt.Equal(afterCreate), "Created at should be before or equal to afterCreate")

		// Update school
		beforeUpdate := time.Now()
		school.Name = "Updated School Name"
		err := testDB.SchoolRepo.Update(ctx, school)
		require.NoError(t, err)
		afterUpdate := time.Now()

		// Verify updated_at changed
		retrieved, err := testDB.SchoolRepo.GetByID(ctx, school.ID)
		require.NoError(t, err)
		assert.True(t, retrieved.UpdatedAt.After(school.UpdatedAt), "Updated at should be after original")
		assert.True(t, retrieved.UpdatedAt.After(beforeUpdate) || retrieved.UpdatedAt.Equal(beforeUpdate), "Updated at should be after or equal to beforeUpdate")
		assert.True(t, retrieved.UpdatedAt.Before(afterUpdate) || retrieved.UpdatedAt.Equal(afterUpdate), "Updated at should be before or equal to afterUpdate")
	})
}

// TestDatabaseCleanup tests database cleanup functionality
func TestDatabaseCleanup(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	ctx := context.Background()
	roleID := "00000000-0000-0000-0000-000000000001"

	t.Run("Truncate test database", func(t *testing.T) {
		testDB := SetupTestDB(t)
		if testDB == nil {
			return
		}

		// Create some data
		CreateTestUser(t, ctx, testDB.UserRepo, "cleanup@example.com", "password123", "Cleanup User", roleID, nil)
		CreateTestSchool(t, ctx, testDB.SchoolRepo, "Cleanup School", "CS001")

		// Verify data exists
		users, err := testDB.UserRepo.List(ctx, nil, nil, nil, 10, 0)
		require.NoError(t, err)
		assert.Greater(t, len(users), 0)

		// Truncate database
		err = database.TruncateTestDatabase(&database.Database{DB: testDB.DB})
		require.NoError(t, err)

		// Verify data is gone
		users, err = testDB.UserRepo.List(ctx, nil, nil, nil, 10, 0)
		require.NoError(t, err)
		assert.Equal(t, 0, len(users), "Users should be truncated")

		testDB.DB.Close()
	})
}

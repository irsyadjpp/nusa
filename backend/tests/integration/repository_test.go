package integration

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestAdvancedUserRepositoryOperations tests advanced user repository operations
func TestAdvancedUserRepositoryOperations(t *testing.T) {
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

	t.Run("User count with filters", func(t *testing.T) {
		school := CreateTestSchool(t, ctx, testDB.SchoolRepo, "Count School", "CS001")

		// Create users in the school
		CreateTestUser(t, ctx, testDB.UserRepo, "count1@example.com", "password123", "Count User 1", roleID, &school.ID)
		CreateTestUser(t, ctx, testDB.UserRepo, "count2@example.com", "password123", "Count User 2", roleID, &school.ID)

		// Count all users
		totalCount, err := testDB.UserRepo.Count(ctx, nil, nil, nil)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, totalCount, 2)

		// Count users by school
		schoolCount, err := testDB.UserRepo.Count(ctx, &school.ID, nil, nil)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, schoolCount, 2)

		// Count active users
		isActive := true
		activeCount, err := testDB.UserRepo.Count(ctx, nil, nil, &isActive)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, activeCount, 2)
	})

	t.Run("User with null school ID", func(t *testing.T) {
		user := CreateTestUser(t, ctx, testDB.UserRepo, "nullschool@example.com", "password123", "Null School User", roleID, nil)

		retrieved, err := testDB.UserRepo.GetByID(ctx, user.ID)
		require.NoError(t, err)
		assert.Nil(t, retrieved.SchoolID)
	})

	t.Run("User with set school ID", func(t *testing.T) {
		school := CreateTestSchool(t, ctx, testDB.SchoolRepo, "Set School", "SS001")
		user := CreateTestUser(t, ctx, testDB.UserRepo, "setschool@example.com", "password123", "Set School User", roleID, &school.ID)

		retrieved, err := testDB.UserRepo.GetByID(ctx, user.ID)
		require.NoError(t, err)
		assert.NotNil(t, retrieved.SchoolID)
		assert.Equal(t, school.ID, *retrieved.SchoolID)
	})
}

// TestAdvancedSchoolRepositoryOperations tests advanced school repository operations
func TestAdvancedSchoolRepositoryOperations(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	testDB := SetupTestDB(t)
	if testDB == nil {
		return
	}
	defer TeardownTestDB(t, testDB)

	ctx := context.Background()

	t.Run("School count with filters", func(t *testing.T) {
		// Create some schools
		CreateTestSchool(t, ctx, testDB.SchoolRepo, "Count School 1", "CS001")
		CreateTestSchool(t, ctx, testDB.SchoolRepo, "Count School 2", "CS002")
		CreateTestSchool(t, ctx, testDB.SchoolRepo, "Count School 3", "CS003")

		// Count all schools
		totalCount, err := testDB.SchoolRepo.Count(ctx, nil)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, totalCount, 3)

		// Count active schools
		isActive := true
		activeCount, err := testDB.SchoolRepo.Count(ctx, &isActive)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, activeCount, 3)
	})

	t.Run("School with null optional fields", func(t *testing.T) {
		school := CreateTestSchool(t, ctx, testDB.SchoolRepo, "Null Fields School", "NFS001")

		retrieved, err := testDB.SchoolRepo.GetByID(ctx, school.ID)
		require.NoError(t, err)
		assert.Nil(t, retrieved.Address)
		assert.Nil(t, retrieved.Phone)
		assert.Nil(t, retrieved.Email)
	})

	t.Run("School with set optional fields", func(t *testing.T) {
		address := "123 Test Street"
		phone := "555-TEST"
		email := "test@school.com"

		school := CreateTestSchool(t, ctx, testDB.SchoolRepo, "Full Fields School", "FFS001")
		school.Address = &address
		school.Phone = &phone
		school.Email = &email

		err := testDB.SchoolRepo.Update(ctx, school)
		require.NoError(t, err)

		retrieved, err := testDB.SchoolRepo.GetByID(ctx, school.ID)
		require.NoError(t, err)
		assert.NotNil(t, retrieved.Address)
		assert.Equal(t, address, *retrieved.Address)
		assert.NotNil(t, retrieved.Phone)
		assert.Equal(t, phone, *retrieved.Phone)
		assert.NotNil(t, retrieved.Email)
		assert.Equal(t, email, *retrieved.Email)
	})
}

// TestAdvancedRoleRepositoryOperations tests advanced role repository operations
func TestAdvancedRoleRepositoryOperations(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	testDB := SetupTestDB(t)
	if testDB == nil {
		return
	}
	defer TeardownTestDB(t, testDB)

	ctx := context.Background()

	t.Run("Add duplicate permission", func(t *testing.T) {
		roleID := "00000000-0000-0000-0000-000000000001"
		resource := "duplicate_resource"
		action := "duplicate_action"

		// Add permission first time
		err := testDB.RoleRepo.AddPermission(ctx, roleID, resource, action)
		require.NoError(t, err)

		// Try to add same permission again (should not error due to ON CONFLICT DO NOTHING)
		err = testDB.RoleRepo.AddPermission(ctx, roleID, resource, action)
		require.NoError(t, err)

		// Verify only one permission exists
		permissions, err := testDB.RoleRepo.GetPermissions(ctx, roleID)
		require.NoError(t, err)

		count := 0
		for _, perm := range permissions {
			if perm.Resource == resource && perm.Action == action {
				count++
			}
		}
		assert.Equal(t, 1, count, "Duplicate permission should not be added")
	})

	t.Run("Remove non-existent permission", func(t *testing.T) {
		roleID := "00000000-0000-0000-0000-000000000001"

		// Try to remove non-existent permission (should not error)
		err := testDB.RoleRepo.RemovePermission(ctx, roleID, "nonexistent", "action")
		require.NoError(t, err)
	})
}

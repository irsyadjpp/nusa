package integration

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestSoftDeleteUser validates soft delete functionality for users
func TestSoftDeleteUser(t *testing.T) {
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

	t.Run("Soft delete user", func(t *testing.T) {
		user := CreateTestUser(t, ctx, testDB.UserRepo, "softdelete@example.com", "password123", "Soft Delete User", roleID, nil)

		// Verify user is active
		assert.True(t, user.IsActive)

		// Soft delete user
		err := testDB.UserRepo.Delete(ctx, user.ID)
		require.NoError(t, err)

		// Verify user is still in database but inactive
		retrieved, err := testDB.UserRepo.GetByID(ctx, user.ID)
		require.NoError(t, err)
		assert.False(t, retrieved.IsActive, "User should be inactive after soft delete")
		assert.Equal(t, user.ID, retrieved.ID, "User ID should remain the same")
		assert.Equal(t, user.Email, retrieved.Email, "User email should remain the same")
	})

	t.Run("Soft delete is idempotent", func(t *testing.T) {
		user := CreateTestUser(t, ctx, testDB.UserRepo, "idempotent@example.com", "password123", "Idempotent User", roleID, nil)

		// First delete
		err := testDB.UserRepo.Delete(ctx, user.ID)
		require.NoError(t, err)

		// Second delete (should not error)
		err = testDB.UserRepo.Delete(ctx, user.ID)
		assert.Error(t, err, "Second delete should fail with 'user not found'")
	})
}

// TestCannotRetrieveSoftDeletedUser validates soft deleted users cannot be retrieved
func TestCannotRetrieveSoftDeletedUser(t *testing.T) {
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

	t.Run("Soft deleted user filtered from list", func(t *testing.T) {
		// Create active and soft deleted users
		activeUser := CreateTestUser(t, ctx, testDB.UserRepo, "active@example.com", "password123", "Active User", roleID, nil)
		deletedUser := CreateTestUser(t, ctx, testDB.UserRepo, "tobedeleted@example.com", "password123", "To Be Deleted", roleID, nil)

		// Soft delete one user
		err := testDB.UserRepo.Delete(ctx, deletedUser.ID)
		require.NoError(t, err)

		// List active users only
		isActive := true
		users, err := testDB.UserRepo.List(ctx, nil, nil, &isActive, 10, 0)
		require.NoError(t, err)

		// Verify active user is in list
		foundActive := false
		foundDeleted := false
		for _, user := range users {
			if user.ID == activeUser.ID {
				foundActive = true
			}
			if user.ID == deletedUser.ID {
				foundDeleted = true
			}
		}
		assert.True(t, foundActive, "Active user should be in list")
		assert.False(t, foundDeleted, "Soft deleted user should not be in active list")
	})

	t.Run("Soft deleted user can still be retrieved by ID", func(t *testing.T) {
		user := CreateTestUser(t, ctx, testDB.UserRepo, "retrieve@example.com", "password123", "Retrieve User", roleID, nil)

		// Soft delete user
		err := testDB.UserRepo.Delete(ctx, user.ID)
		require.NoError(t, err)

		// Can still retrieve by ID (for audit purposes)
		retrieved, err := testDB.UserRepo.GetByID(ctx, user.ID)
		require.NoError(t, err)
		assert.Equal(t, user.ID, retrieved.ID)
		assert.False(t, retrieved.IsActive)
	})

	t.Run("Soft deleted user can be retrieved by email", func(t *testing.T) {
		user := CreateTestUser(t, ctx, testDB.UserRepo, "email@example.com", "password123", "Email User", roleID, nil)

		// Soft delete user
		err := testDB.UserRepo.Delete(ctx, user.ID)
		require.NoError(t, err)

		// Can still retrieve by email
		retrieved, err := testDB.UserRepo.GetByEmail(ctx, user.Email)
		require.NoError(t, err)
		assert.Equal(t, user.ID, retrieved.ID)
		assert.False(t, retrieved.IsActive)
	})
}

// TestRestoreSoftDeletedUser validates restore functionality for soft deleted users
func TestRestoreSoftDeletedUser(t *testing.T) {
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

	t.Run("Restore soft deleted user", func(t *testing.T) {
		user := CreateTestUser(t, ctx, testDB.UserRepo, "restore@example.com", "password123", "Restore User", roleID, nil)

		// Soft delete user
		err := testDB.UserRepo.Delete(ctx, user.ID)
		require.NoError(t, err)

		// Verify user is inactive
		retrieved, err := testDB.UserRepo.GetByID(ctx, user.ID)
		require.NoError(t, err)
		assert.False(t, retrieved.IsActive)

		// Restore user (reactivate)
		retrieved.IsActive = true
		err = testDB.UserRepo.Update(ctx, retrieved)
		require.NoError(t, err)

		// Verify user is active again
		restored, err := testDB.UserRepo.GetByID(ctx, user.ID)
		require.NoError(t, err)
		assert.True(t, restored.IsActive, "User should be active after restore")

		// Verify user appears in active list
		isActive := true
		users, err := testDB.UserRepo.List(ctx, nil, nil, &isActive, 10, 0)
		require.NoError(t, err)

		found := false
		for _, u := range users {
			if u.ID == user.ID {
				found = true
				break
			}
		}
		assert.True(t, found, "Restored user should appear in active list")
	})
}

// TestSoftDeleteSchool validates soft delete functionality for schools
func TestSoftDeleteSchool(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	testDB := SetupTestDB(t)
	if testDB == nil {
		return
	}
	defer TeardownTestDB(t, testDB)

	ctx := context.Background()

	t.Run("Soft delete school", func(t *testing.T) {
		school := CreateTestSchool(t, ctx, testDB.SchoolRepo, "Soft Delete School", "SDS001")

		// Verify school is active
		assert.True(t, school.IsActive)

		// Soft delete school
		err := testDB.SchoolRepo.Delete(ctx, school.ID)
		require.NoError(t, err)

		// Verify school is still in database but inactive
		retrieved, err := testDB.SchoolRepo.GetByID(ctx, school.ID)
		require.NoError(t, err)
		assert.False(t, retrieved.IsActive, "School should be inactive after soft delete")
		assert.Equal(t, school.ID, retrieved.ID, "School ID should remain the same")
		assert.Equal(t, school.Code, retrieved.Code, "School code should remain the same")
	})

	t.Run("Soft delete school with users", func(t *testing.T) {
		roleID := "00000000-0000-0000-0000-000000000001"

		school := CreateTestSchool(t, ctx, testDB.SchoolRepo, "School with Users", "SWU001")
		user := CreateTestUser(t, ctx, testDB.UserRepo, "schooluser@example.com", "password123", "School User", roleID, &school.ID)

		// Soft delete school
		err := testDB.SchoolRepo.Delete(ctx, school.ID)
		require.NoError(t, err)

		// Verify school is inactive
		retrievedSchool, err := testDB.SchoolRepo.GetByID(ctx, school.ID)
		require.NoError(t, err)
		assert.False(t, retrievedSchool.IsActive)

		// Verify user still exists (no cascade delete)
		retrievedUser, err := testDB.UserRepo.GetByID(ctx, user.ID)
		require.NoError(t, err)
		assert.Equal(t, user.ID, retrievedUser.ID)
		assert.NotNil(t, retrievedUser.SchoolID)
		assert.Equal(t, school.ID, *retrievedUser.SchoolID)

		// User should still be active (school deactivation doesn't cascade)
		assert.True(t, retrievedUser.IsActive)
	})

	t.Run("Soft deleted school filtered from list", func(t *testing.T) {
		// Create active and soft deleted schools
		activeSchool := CreateTestSchool(t, ctx, testDB.SchoolRepo, "Active School", "AS001")
		deletedSchool := CreateTestSchool(t, ctx, testDB.SchoolRepo, "To Be Deleted", "TBD001")

		// Soft delete one school
		err := testDB.SchoolRepo.Delete(ctx, deletedSchool.ID)
		require.NoError(t, err)

		// List active schools only
		isActive := true
		schools, err := testDB.SchoolRepo.List(ctx, &isActive, 10, 0)
		require.NoError(t, err)

		// Verify active school is in list
		foundActive := false
		foundDeleted := false
		for _, school := range schools {
			if school.ID == activeSchool.ID {
				foundActive = true
			}
			if school.ID == deletedSchool.ID {
				foundDeleted = true
			}
		}
		assert.True(t, foundActive, "Active school should be in list")
		assert.False(t, foundDeleted, "Soft deleted school should not be in active list")
	})

	t.Run("Restore soft deleted school", func(t *testing.T) {
		school := CreateTestSchool(t, ctx, testDB.SchoolRepo, "Restore School", "RS001")

		// Soft delete school
		err := testDB.SchoolRepo.Delete(ctx, school.ID)
		require.NoError(t, err)

		// Verify school is inactive
		retrieved, err := testDB.SchoolRepo.GetByID(ctx, school.ID)
		require.NoError(t, err)
		assert.False(t, retrieved.IsActive)

		// Restore school (reactivate)
		retrieved.IsActive = true
		err = testDB.SchoolRepo.Update(ctx, retrieved)
		require.NoError(t, err)

		// Verify school is active again
		restored, err := testDB.SchoolRepo.GetByID(ctx, school.ID)
		require.NoError(t, err)
		assert.True(t, restored.IsActive, "School should be active after restore")

		// Verify school appears in active list
		isActive := true
		schools, err := testDB.SchoolRepo.List(ctx, &isActive, 10, 0)
		require.NoError(t, err)

		found := false
		for _, s := range schools {
			if s.ID == school.ID {
				found = true
				break
			}
		}
		assert.True(t, found, "Restored school should appear in active list")
	})
}

// TestSoftDeleteRole validates soft delete functionality for roles
func TestSoftDeleteRole(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	testDB := SetupTestDB(t)
	if testDB == nil {
		return
	}
	defer TeardownTestDB(t, testDB)

	ctx := context.Background()

	t.Run("Soft delete custom role", func(t *testing.T) {
		name := "CUSTOM_SOFT_DELETE"
		role := CreateTestRole(t, ctx, testDB.RoleRepo, name, nil)

		// Verify role is active
		assert.True(t, role.IsActive)

		// Soft delete role (deactivate)
		role.IsActive = false
		err := testDB.RoleRepo.Update(ctx, role)
		require.NoError(t, err)

		// Verify role is inactive
		retrieved, err := testDB.RoleRepo.GetByID(ctx, role.ID)
		require.NoError(t, err)
		assert.False(t, retrieved.IsActive, "Role should be inactive after soft delete")
	})

	t.Run("Soft deleted role filtered from list", func(t *testing.T) {
		// Create active and inactive roles
		activeRole := CreateTestRole(t, ctx, testDB.RoleRepo, "ACTIVE_CUSTOM_ROLE", nil)
		inactiveRole := CreateTestRole(t, ctx, testDB.RoleRepo, "INACTIVE_CUSTOM_ROLE", nil)

		// Deactivate one role
		inactiveRole.IsActive = false
		err := testDB.RoleRepo.Update(ctx, inactiveRole)
		require.NoError(t, err)

		// List active roles only
		isActive := true
		roles, err := testDB.RoleRepo.List(ctx, &isActive)
		require.NoError(t, err)

		// Verify active role is in list
		foundActive := false
		foundInactive := false
		for _, role := range roles {
			if role.ID == activeRole.ID {
				foundActive = true
			}
			if role.ID == inactiveRole.ID {
				foundInactive = true
			}
		}
		assert.True(t, foundActive, "Active role should be in list")
		assert.False(t, foundInactive, "Inactive role should not be in active list")
	})
}

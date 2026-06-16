package repository

import (
	"context"
	"testing"

	"github.com/nusa/backend/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestRoleRepository_Create tests creating a new role
func TestRoleRepository_Create(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewRoleRepository(testDB.Pool)

	t.Run("Success - Create role with description", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		description := "Teacher role for classroom management"
		role := &domain.Role{
			ID:          "role-001",
			Name:        "TEACHER",
			Description: &description,
			IsActive:    true,
		}

		err := repo.Create(ctx, role)
		require.NoError(t, err)
		assert.NotEmpty(t, role.ID)
	})

	t.Run("Success - Create role without description", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		role := &domain.Role{
			ID:       "role-002",
			Name:     "ADMIN",
			IsActive: true,
		}

		err := repo.Create(ctx, role)
		require.NoError(t, err)
	})

	t.Run("Error - Duplicate name", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()

		role1 := &domain.Role{
			ID:       "role-003",
			Name:     "DUPLICATE",
			IsActive: true,
		}

		err := repo.Create(ctx, role1)
		require.NoError(t, err)

		role2 := &domain.Role{
			ID:       "role-004",
			Name:     "DUPLICATE", // Same name
			IsActive: true,
		}

		err = repo.Create(ctx, role2)
		assert.Error(t, err)
	})
}

// TestRoleRepository_GetByID tests retrieving a role by ID
func TestRoleRepository_GetByID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewRoleRepository(testDB.Pool)

	t.Run("Success - Role found", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "TEACHER")

		role, err := repo.GetByID(ctx, roleID)
		require.NoError(t, err)
		assert.NotNil(t, role)
		assert.Equal(t, roleID, role.ID)
		assert.Equal(t, "TEACHER", role.Name)
		assert.True(t, role.IsActive)
	})

	t.Run("Success - Role with description", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "ADMIN")

		role, err := repo.GetByID(ctx, roleID)
		require.NoError(t, err)
		assert.NotNil(t, role)
		assert.NotNil(t, role.Description)
	})

	t.Run("Error - Role not found", func(t *testing.T) {
		ctx := context.Background()

		role, err := repo.GetByID(ctx, "non-existent-id")
		assert.Error(t, err)
		assert.Nil(t, role)
		assert.Contains(t, err.Error(), "role not found")
	})
}

// TestRoleRepository_GetByName tests retrieving a role by name
func TestRoleRepository_GetByName(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewRoleRepository(testDB.Pool)

	t.Run("Success - Role found by name", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "TEACHER")

		role, err := repo.GetByName(ctx, "TEACHER")
		require.NoError(t, err)
		assert.NotNil(t, role)
		assert.Equal(t, roleID, role.ID)
		assert.Equal(t, "TEACHER", role.Name)
	})

	t.Run("Error - Role not found by name", func(t *testing.T) {
		ctx := context.Background()

		role, err := repo.GetByName(ctx, "NONEXISTENT")
		assert.Error(t, err)
		assert.Nil(t, role)
		assert.Contains(t, err.Error(), "role not found")
	})
}

// TestRoleRepository_Update tests updating a role
func TestRoleRepository_Update(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewRoleRepository(testDB.Pool)

	t.Run("Success - Update role name", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "OLD_NAME")

		role, err := repo.GetByID(ctx, roleID)
		require.NoError(t, err)

		role.Name = "NEW_NAME"

		err = repo.Update(ctx, role)
		require.NoError(t, err)

		updatedRole, err := repo.GetByID(ctx, roleID)
		require.NoError(t, err)
		assert.Equal(t, "NEW_NAME", updatedRole.Name)
	})

	t.Run("Success - Update role description", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "TEACHER")

		role, err := repo.GetByID(ctx, roleID)
		require.NoError(t, err)

		newDescription := "Updated teacher description"
		role.Description = &newDescription

		err = repo.Update(ctx, role)
		require.NoError(t, err)

		updatedRole, err := repo.GetByID(ctx, roleID)
		require.NoError(t, err)
		assert.NotNil(t, updatedRole.Description)
		assert.Equal(t, newDescription, *updatedRole.Description)
	})

	t.Run("Success - Deactivate role", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "TEACHER")

		role, err := repo.GetByID(ctx, roleID)
		require.NoError(t, err)

		role.IsActive = false

		err = repo.Update(ctx, role)
		require.NoError(t, err)

		updatedRole, err := repo.GetByID(ctx, roleID)
		require.NoError(t, err)
		assert.False(t, updatedRole.IsActive)
	})

	t.Run("Error - Role not found", func(t *testing.T) {
		ctx := context.Background()

		role := &domain.Role{
			ID:   "non-existent-id",
			Name: "Non-existent",
		}

		err := repo.Update(ctx, role)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "role not found")
	})
}

// TestRoleRepository_List tests listing roles with filters
func TestRoleRepository_List(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewRoleRepository(testDB.Pool)

	t.Run("Success - List all roles", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()

		CreateTestRole(t, testDB.Pool, "TEACHER")
		CreateTestRole(t, testDB.Pool, "ADMIN")
		CreateTestRole(t, testDB.Pool, "STUDENT")

		roles, err := repo.List(ctx, nil)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(roles), 3)
	})

	t.Run("Success - Filter active roles", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()

		activeRoleID := CreateTestRole(t, testDB.Pool, "ACTIVE_ROLE")
		inactiveRoleID := CreateTestRole(t, testDB.Pool, "INACTIVE_ROLE")

		// Deactivate one role
		inactiveRole, err := repo.GetByID(ctx, inactiveRoleID)
		require.NoError(t, err)
		inactiveRole.IsActive = false
		err = repo.Update(ctx, inactiveRole)
		require.NoError(t, err)

		isActive := true
		roles, err := repo.List(ctx, &isActive)
		require.NoError(t, err)

		roleIDs := make([]string, len(roles))
		for i, r := range roles {
			roleIDs[i] = r.ID
		}
		assert.Contains(t, roleIDs, activeRoleID)
		assert.NotContains(t, roleIDs, inactiveRoleID)
	})

	t.Run("Success - Filter inactive roles", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()

		activeRoleID := CreateTestRole(t, testDB.Pool, "ACTIVE_ROLE")
		inactiveRoleID := CreateTestRole(t, testDB.Pool, "INACTIVE_ROLE")

		// Deactivate one role
		inactiveRole, err := repo.GetByID(ctx, inactiveRoleID)
		require.NoError(t, err)
		inactiveRole.IsActive = false
		err = repo.Update(ctx, inactiveRole)
		require.NoError(t, err)

		isActive := false
		roles, err := repo.List(ctx, &isActive)
		require.NoError(t, err)

		roleIDs := make([]string, len(roles))
		for i, r := range roles {
			roleIDs[i] = r.ID
		}
		assert.NotContains(t, roleIDs, activeRoleID)
		assert.Contains(t, roleIDs, inactiveRoleID)
	})
}

// TestRoleRepository_AddPermission tests adding permissions to a role
func TestRoleRepository_AddPermission(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewRoleRepository(testDB.Pool)

	t.Run("Success - Add permission", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "TEACHER")

		err := repo.AddPermission(ctx, roleID, "class", "create")
		require.NoError(t, err)

		permissions, err := repo.GetPermissions(ctx, roleID)
		require.NoError(t, err)
		assert.Equal(t, 1, len(permissions))
		assert.Equal(t, "class", permissions[0].Resource)
		assert.Equal(t, "create", permissions[0].Action)
	})

	t.Run("Success - Add multiple permissions", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "TEACHER")

		repo.AddPermission(ctx, roleID, "class", "create")
		repo.AddPermission(ctx, roleID, "class", "read")
		repo.AddPermission(ctx, roleID, "student", "read")

		permissions, err := repo.GetPermissions(ctx, roleID)
		require.NoError(t, err)
		assert.Equal(t, 3, len(permissions))
	})

	t.Run("Success - Duplicate permission idempotent", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "TEACHER")

		// Add same permission twice
		err := repo.AddPermission(ctx, roleID, "class", "create")
		require.NoError(t, err)

		err = repo.AddPermission(ctx, roleID, "class", "create")
		require.NoError(t, err) // Should not error due to ON CONFLICT

		permissions, err := repo.GetPermissions(ctx, roleID)
		require.NoError(t, err)
		assert.Equal(t, 1, len(permissions))
	})
}

// TestRoleRepository_GetPermissions tests retrieving permissions for a role
func TestRoleRepository_GetPermissions(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewRoleRepository(testDB.Pool)

	t.Run("Success - Get permissions for role", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "TEACHER")

		repo.AddPermission(ctx, roleID, "class", "create")
		repo.AddPermission(ctx, roleID, "class", "read")
		repo.AddPermission(ctx, roleID, "student", "read")

		permissions, err := repo.GetPermissions(ctx, roleID)
		require.NoError(t, err)
		assert.Equal(t, 3, len(permissions))
	})

	t.Run("Success - No permissions for role", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "TEACHER")

		permissions, err := repo.GetPermissions(ctx, roleID)
		require.NoError(t, err)
		assert.Equal(t, 0, len(permissions))
	})

	t.Run("Success - Permissions sorted by resource and action", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "TEACHER")

		repo.AddPermission(ctx, roleID, "student", "read")
		repo.AddPermission(ctx, roleID, "class", "create")
		repo.AddPermission(ctx, roleID, "class", "read")

		permissions, err := repo.GetPermissions(ctx, roleID)
		require.NoError(t, err)
		assert.Equal(t, 3, len(permissions))

		// Should be sorted by resource, then action
		assert.Equal(t, "class", permissions[0].Resource)
		assert.Equal(t, "class", permissions[1].Resource)
		assert.Equal(t, "student", permissions[2].Resource)
	})
}

// TestRoleRepository_RemovePermission tests removing permissions from a role
func TestRoleRepository_RemovePermission(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewRoleRepository(testDB.Pool)

	t.Run("Success - Remove permission", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "TEACHER")

		repo.AddPermission(ctx, roleID, "class", "create")
		repo.AddPermission(ctx, roleID, "class", "read")

		err := repo.RemovePermission(ctx, roleID, "class", "create")
		require.NoError(t, err)

		permissions, err := repo.GetPermissions(ctx, roleID)
		require.NoError(t, err)
		assert.Equal(t, 1, len(permissions))
		assert.Equal(t, "read", permissions[0].Action)
	})

	t.Run("Success - Remove non-existent permission (no error)", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "TEACHER")

		// Remove permission that doesn't exist
		err := repo.RemovePermission(ctx, roleID, "class", "create")
		require.NoError(t, err)
	})
}

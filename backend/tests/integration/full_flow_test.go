package integration

import (
	"context"
	"testing"
	"time"

	"github.com/nusa/backend/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestFullUserLifecycle validates complete user CRUD lifecycle
func TestFullUserLifecycle(t *testing.T) {
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

	t.Run("Complete user lifecycle", func(t *testing.T) {
		// 1. Create user
		email := "lifecycle@example.com"
		password := "password123"
		name := "Lifecycle User"

		user := CreateTestUser(t, ctx, testDB.UserRepo, email, password, name, roleID, nil)
		assert.NotNil(t, user)
		assert.NotEmpty(t, user.ID)
		assert.Equal(t, email, user.Email)
		assert.Equal(t, name, user.Name)
		assert.True(t, user.IsActive)

		// 2. Retrieve user by ID
		retrieved, err := testDB.UserRepo.GetByID(ctx, user.ID)
		require.NoError(t, err)
		assert.Equal(t, user.ID, retrieved.ID)
		assert.Equal(t, email, retrieved.Email)
		assert.Equal(t, name, retrieved.Name)

		// 3. Retrieve user by email
		retrievedByEmail, err := testDB.UserRepo.GetByEmail(ctx, email)
		require.NoError(t, err)
		assert.Equal(t, user.ID, retrievedByEmail.ID)

		// 4. Update user information
		newName := "Updated Lifecycle User"
		retrieved.Name = newName
		err = testDB.UserRepo.Update(ctx, retrieved)
		require.NoError(t, err)

		updated, err := testDB.UserRepo.GetByID(ctx, user.ID)
		require.NoError(t, err)
		assert.Equal(t, newName, updated.Name)

		// 5. List users with pagination
		users, err := testDB.UserRepo.List(ctx, nil, nil, nil, 10, 0)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(users), 1)

		// 6. Deactivate user (soft delete)
		err = testDB.UserRepo.Delete(ctx, user.ID)
		require.NoError(t, err)

		deleted, err := testDB.UserRepo.GetByID(ctx, user.ID)
		require.NoError(t, err)
		assert.False(t, deleted.IsActive, "User should be deactivated")

		// 7. Reactivate user
		deleted.IsActive = true
		err = testDB.UserRepo.Update(ctx, deleted)
		require.NoError(t, err)

		reactivated, err := testDB.UserRepo.GetByID(ctx, user.ID)
		require.NoError(t, err)
		assert.True(t, reactivated.IsActive, "User should be reactivated")
	})
}

// TestFullSchoolLifecycle validates complete school CRUD lifecycle
func TestFullSchoolLifecycle(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	testDB := SetupTestDB(t)
	if testDB == nil {
		return
	}
	defer TeardownTestDB(t, testDB)

	ctx := context.Background()

	t.Run("Complete school lifecycle", func(t *testing.T) {
		// 1. Create school
		code := "LCS001"
		name := "Lifecycle School"

		school := CreateTestSchool(t, ctx, testDB.SchoolRepo, name, code)
		assert.NotNil(t, school)
		assert.NotEmpty(t, school.ID)
		assert.Equal(t, code, school.Code)
		assert.Equal(t, name, school.Name)
		assert.True(t, school.IsActive)

		// 2. Retrieve school by ID
		retrieved, err := testDB.SchoolRepo.GetByID(ctx, school.ID)
		require.NoError(t, err)
		assert.Equal(t, school.ID, retrieved.ID)
		assert.Equal(t, code, retrieved.Code)
		assert.Equal(t, name, retrieved.Name)

		// 3. Retrieve school by code
		retrievedByCode, err := testDB.SchoolRepo.GetByCode(ctx, code)
		require.NoError(t, err)
		assert.Equal(t, school.ID, retrievedByCode.ID)

		// 4. Update school information
		newName := "Updated Lifecycle School"
		newAddress := "123 Updated Street"
		school.Name = newName
		school.Address = &newAddress
		err = testDB.SchoolRepo.Update(ctx, school)
		require.NoError(t, err)

		updated, err := testDB.SchoolRepo.GetByID(ctx, school.ID)
		require.NoError(t, err)
		assert.Equal(t, newName, updated.Name)
		assert.Equal(t, newAddress, *updated.Address)

		// 5. List schools with pagination
		schools, err := testDB.SchoolRepo.List(ctx, nil, 10, 0)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(schools), 1)

		// 6. Deactivate school (soft delete)
		err = testDB.SchoolRepo.Delete(ctx, school.ID)
		require.NoError(t, err)

		deleted, err := testDB.SchoolRepo.GetByID(ctx, school.ID)
		require.NoError(t, err)
		assert.False(t, deleted.IsActive, "School should be deactivated")

		// 7. Reactivate school
		deleted.IsActive = true
		err = testDB.SchoolRepo.Update(ctx, deleted)
		require.NoError(t, err)

		reactivated, err := testDB.SchoolRepo.GetByID(ctx, school.ID)
		require.NoError(t, err)
		assert.True(t, reactivated.IsActive, "School should be reactivated")
	})
}

// TestFullRoleManagementFlow validates complete role management operations
func TestFullRoleManagementFlow(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	testDB := SetupTestDB(t)
	if testDB == nil {
		return
	}
	defer TeardownTestDB(t, testDB)

	ctx := context.Background()

	t.Run("Complete role lifecycle", func(t *testing.T) {
		// 1. Create custom role
		name := "CUSTOM_LIFECYCLE"
		description := "Custom role for lifecycle testing"

		role := CreateTestRole(t, ctx, testDB.RoleRepo, name, &description)
		assert.NotNil(t, role)
		assert.NotEmpty(t, role.ID)
		assert.Equal(t, name, role.Name)
		assert.Equal(t, description, *role.Description)
		assert.True(t, role.IsActive)

		// 2. Retrieve role by ID
		retrieved, err := testDB.RoleRepo.GetByID(ctx, role.ID)
		require.NoError(t, err)
		assert.Equal(t, role.ID, retrieved.ID)
		assert.Equal(t, name, retrieved.Name)

		// 3. Retrieve role by name
		retrievedByName, err := testDB.RoleRepo.GetByName(ctx, name)
		require.NoError(t, err)
		assert.Equal(t, role.ID, retrievedByName.ID)

		// 4. Add permissions to role
		err = testDB.RoleRepo.AddPermission(ctx, role.ID, "test_resource", "test_action")
		require.NoError(t, err)

		err = testDB.RoleRepo.AddPermission(ctx, role.ID, "another_resource", "read")
		require.NoError(t, err)

		// 5. Retrieve role with permissions
		permissions, err := testDB.RoleRepo.GetPermissions(ctx, role.ID)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(permissions), 2)

		// 6. Update role information
		newDescription := "Updated role description"
		role.Description = &newDescription
		err = testDB.RoleRepo.Update(ctx, role)
		require.NoError(t, err)

		updated, err := testDB.RoleRepo.GetByID(ctx, role.ID)
		require.NoError(t, err)
		assert.Equal(t, newDescription, *updated.Description)

		// 7. List roles
		roles, err := testDB.RoleRepo.List(ctx, nil)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(roles), 4) // 3 system + 1 custom

		// 8. Remove permission from role
		err = testDB.RoleRepo.RemovePermission(ctx, role.ID, "test_resource", "test_action")
		require.NoError(t, err)

		permissionsAfter, err := testDB.RoleRepo.GetPermissions(ctx, role.ID)
		require.NoError(t, err)
		assert.Equal(t, len(permissions)-1, len(permissionsAfter), "One permission should be removed")
	})
}

// TestUserWithSchoolAssignment validates user-school relationship
func TestUserWithSchoolAssignment(t *testing.T) {
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

	t.Run("User with school assignment flow", func(t *testing.T) {
		// 1. Create school
		school := CreateTestSchool(t, ctx, testDB.SchoolRepo, "User School", "US001")
		assert.NotNil(t, school)

		// 2. Create user with school assignment
		user := CreateTestUser(t, ctx, testDB.UserRepo, "schooluser@example.com", "password123", "School User", roleID, &school.ID)
		assert.NotNil(t, user)
		assert.NotNil(t, user.SchoolID)
		assert.Equal(t, school.ID, *user.SchoolID)

		// 3. Verify school association
		retrieved, err := testDB.UserRepo.GetByID(ctx, user.ID)
		require.NoError(t, err)
		assert.NotNil(t, retrieved.SchoolID)
		assert.Equal(t, school.ID, *retrieved.SchoolID)

		// 4. Update user school
		newSchool := CreateTestSchool(t, ctx, testDB.SchoolRepo, "New School", "NS002")
		retrieved.SchoolID = &newSchool.ID
		err = testDB.UserRepo.Update(ctx, retrieved)
		require.NoError(t, err)

		updated, err := testDB.UserRepo.GetByID(ctx, user.ID)
		require.NoError(t, err)
		assert.Equal(t, newSchool.ID, *updated.SchoolID)

		// 5. List users by school
		users, err := testDB.UserRepo.List(ctx, &newSchool.ID, nil, nil, 10, 0)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(users), 1)

		// 6. Verify user is in the list
		found := false
		for _, u := range users {
			if u.ID == user.ID {
				found = true
				break
			}
		}
		assert.True(t, found, "User should be in school user list")
	})
}

// TestRefreshTokenLifecycle validates complete refresh token lifecycle
func TestRefreshTokenLifecycle(t *testing.T) {
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

	t.Run("Complete refresh token lifecycle", func(t *testing.T) {
		// 1. Create user
		user := CreateTestUser(t, ctx, testDB.UserRepo, "tokenuser@example.com", "password123", "Token User", roleID, nil)

		// 2. Create refresh token
		token := "lifecycle-refresh-token"
		expiresAt := time.Now().Add(7 * 24 * time.Hour)

		err := CreateTestRefreshToken(t, ctx, testDB.RefreshTokenRepo, user.ID, token, expiresAt)
		require.NoError(t, err)

		// 3. Validate token is active
		userID, err := testDB.RefreshTokenRepo.GetByToken(ctx, token)
		require.NoError(t, err)
		assert.Equal(t, user.ID, *userID)

		// 4. Create another token for the same user
		token2 := "lifecycle-refresh-token-2"
		expiresAt2 := time.Now().Add(7 * 24 * time.Hour)

		err = CreateTestRefreshToken(t, ctx, testDB.RefreshTokenRepo, user.ID, token2, expiresAt2)
		require.NoError(t, err)

		// 5. Revoke first token
		err = testDB.RefreshTokenRepo.Revoke(ctx, token)
		require.NoError(t, err)

		// 6. Verify first token is revoked
		_, err = testDB.RefreshTokenRepo.GetByToken(ctx, token)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "revoked")

		// 7. Verify second token is still active
		userID2, err := testDB.RefreshTokenRepo.GetByToken(ctx, token2)
		require.NoError(t, err)
		assert.Equal(t, user.ID, *userID2)

		// 8. Revoke all tokens for user
		err = testDB.RefreshTokenRepo.RevokeAllForUser(ctx, user.ID)
		require.NoError(t, err)

		// 9. Verify all tokens are revoked
		_, err = testDB.RefreshTokenRepo.GetByToken(ctx, token2)
		assert.Error(t, err)
	})
}

// TestPermissionManagementFlow validates complete permission management
func TestPermissionManagementFlow(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	testDB := SetupTestDB(t)
	if testDB == nil {
		return
	}
	defer TeardownTestDB(t, testDB)

	ctx := context.Background()

	t.Run("Complete permission management flow", func(t *testing.T) {
		// 1. Get system admin role
		role, err := testDB.RoleRepo.GetByName(ctx, domain.RoleSystemAdmin)
		require.NoError(t, err)

		// 2. Add custom permission
		err = testDB.RoleRepo.AddPermission(ctx, role.ID, "custom_resource", "custom_action")
		require.NoError(t, err)

		// 3. Verify permission was added
		permissions, err := testDB.RoleRepo.GetPermissions(ctx, role.ID)
		require.NoError(t, err)

		found := false
		for _, perm := range permissions {
			if perm.Resource == "custom_resource" && perm.Action == "custom_action" {
				found = true
				break
			}
		}
		assert.True(t, found, "Custom permission should be added")

		// 4. Add another permission
		err = testDB.RoleRepo.AddPermission(ctx, role.ID, "another_resource", "another_action")
		require.NoError(t, err)

		// 5. Verify both permissions exist
		permissions, err = testDB.RoleRepo.GetPermissions(ctx, role.ID)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(permissions), 2)

		// 6. Remove one permission
		err = testDB.RoleRepo.RemovePermission(ctx, role.ID, "custom_resource", "custom_action")
		require.NoError(t, err)

		// 7. Verify permission was removed
		permissions, err = testDB.RoleRepo.GetPermissions(ctx, role.ID)
		require.NoError(t, err)

		found = false
		for _, perm := range permissions {
			if perm.Resource == "custom_resource" && perm.Action == "custom_action" {
				found = true
				break
			}
		}
		assert.False(t, found, "Custom permission should be removed")
	})
}

// TestComplexScenario validates complex multi-entity scenario
func TestComplexScenario(t *testing.T) {
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

	t.Run("School with multiple users and roles", func(t *testing.T) {
		// 1. Create school
		school := CreateTestSchool(t, ctx, testDB.SchoolRepo, "Complex School", "CS001")

		// 2. Create multiple users in the school
		adminUser := CreateTestUser(t, ctx, testDB.UserRepo, "admin@complex.com", "password123", "School Admin", roleID, &school.ID)

		teacherRoleID := "00000000-0000-0000-0000-000000000003"
		teacher1 := CreateTestUser(t, ctx, testDB.UserRepo, "teacher1@complex.com", "password123", "Teacher 1", teacherRoleID, &school.ID)
		teacher2 := CreateTestUser(t, ctx, testDB.UserRepo, "teacher2@complex.com", "password123", "Teacher 2", teacherRoleID, &school.ID)

		// 3. List users by school
		schoolUsers, err := testDB.UserRepo.List(ctx, &school.ID, nil, nil, 10, 0)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(schoolUsers), 3)

		// 4. Verify all users are in the list
		userIDs := make(map[string]bool)
		for _, u := range schoolUsers {
			userIDs[u.ID] = true
		}
		assert.True(t, userIDs[adminUser.ID], "Admin user should be in school list")
		assert.True(t, userIDs[teacher1.ID], "Teacher 1 should be in school list")
		assert.True(t, userIDs[teacher2.ID], "Teacher 2 should be in school list")

		// 5. List users by role
				teacherRoleID = "00000000-0000-0000-0000-000000000003"
				teachers, err := testDB.UserRepo.List(ctx, nil, &teacherRoleID, nil, 10, 0)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(teachers), 2)

		// 6. Update school status
		err = testDB.SchoolRepo.UpdateStatus(ctx, school.ID, false)
		require.NoError(t, err)

		updatedSchool, err := testDB.SchoolRepo.GetByID(ctx, school.ID)
		require.NoError(t, err)
		assert.False(t, updatedSchool.IsActive, "School should be deactivated")

		// 7. Verify users still exist (they should not be cascade deleted)
		_, err = testDB.UserRepo.GetByID(ctx, adminUser.ID)
		assert.NoError(t, err, "Admin user should still exist after school deactivation")
	})

	t.Run("User with refresh tokens and permissions", func(t *testing.T) {
		// 1. Create user
		user := CreateTestUser(t, ctx, testDB.UserRepo, "complex@user.com", "password123", "Complex User", roleID, nil)

		// 2. Create multiple refresh tokens
		token1 := "complex-token-1"
		token2 := "complex-token-2"
		token3 := "complex-token-3"
		expiresAt := time.Now().Add(7 * 24 * time.Hour)

		CreateTestRefreshToken(t, ctx, testDB.RefreshTokenRepo, user.ID, token1, expiresAt)
		CreateTestRefreshToken(t, ctx, testDB.RefreshTokenRepo, user.ID, token2, expiresAt)
		CreateTestRefreshToken(t, ctx, testDB.RefreshTokenRepo, user.ID, token3, expiresAt)

		// 3. Add permissions to user's role
		err := testDB.RoleRepo.AddPermission(ctx, roleID, "complex_resource", "read")
		require.NoError(t, err)
		err = testDB.RoleRepo.AddPermission(ctx, roleID, "complex_resource", "write")
		require.NoError(t, err)

		// 4. Verify all tokens are active
		for _, token := range []string{token1, token2, token3} {
			userID, err := testDB.RefreshTokenRepo.GetByToken(ctx, token)
			require.NoError(t, err)
			assert.Equal(t, user.ID, *userID)
		}

		// 5. Verify permissions exist
		permissions, err := testDB.RoleRepo.GetPermissions(ctx, roleID)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(permissions), 2)

		// 6. Deactivate user
		err = testDB.UserRepo.Delete(ctx, user.ID)
		require.NoError(t, err)

		// 7. Revoke all tokens
		err = testDB.RefreshTokenRepo.RevokeAllForUser(ctx, user.ID)
		require.NoError(t, err)

		// 8. Verify all tokens are revoked
		for _, token := range []string{token1, token2, token3} {
			_, err := testDB.RefreshTokenRepo.GetByToken(ctx, token)
			assert.Error(t, err)
		}
	})
}

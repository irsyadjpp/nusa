package repository

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/nusa/backend/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestUserRepository_Create tests creating a new user
func TestUserRepository_Create(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewUserRepository(testDB.Pool)

	t.Run("Success - Create user with school", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "TEACHER")
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")

		user := &domain.User{
			ID:                  "user-001",
			Email:               "teacher@example.com",
			Name:                "Test Teacher",
			PasswordHash:        "hashed_password",
			RoleID:              roleID,
			SchoolID:            &schoolID,
			IsActive:            true,
			FailedLoginAttempts: 0,
			CreatedBy:           &schoolID,
			UpdatedBy:           &schoolID,
		}

		err := repo.Create(ctx, user)
		require.NoError(t, err)
		assert.NotEmpty(t, user.ID)
	})

	t.Run("Success - Create user without school", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "ADMIN")

		user := &domain.User{
			ID:                  "user-002",
			Email:               "admin@example.com",
			Name:                "Test Admin",
			PasswordHash:        "hashed_password",
			RoleID:              roleID,
			SchoolID:            nil,
			IsActive:            true,
			FailedLoginAttempts: 0,
			CreatedBy:           StringPtr("system"),
			UpdatedBy:           StringPtr("system"),
		}

		err := repo.Create(ctx, user)
		require.NoError(t, err)
	})

	t.Run("Error - Duplicate email", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "TEACHER")
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")

		user1 := &domain.User{
			ID:           "user-003",
			Email:        "duplicate@example.com",
			Name:         "Test User 1",
			PasswordHash: "hashed_password",
			RoleID:       roleID,
			SchoolID:     &schoolID,
			IsActive:     true,
			CreatedBy:    &schoolID,
			UpdatedBy:    &schoolID,
		}

		err := repo.Create(ctx, user1)
		require.NoError(t, err)

		user2 := &domain.User{
			ID:           "user-004",
			Email:        "duplicate@example.com", // Same email
			Name:         "Test User 2",
			PasswordHash: "hashed_password",
			RoleID:       roleID,
			SchoolID:     &schoolID,
			IsActive:     true,
			CreatedBy:    &schoolID,
			UpdatedBy:    &schoolID,
		}

		err = repo.Create(ctx, user2)
		assert.Error(t, err)
	})
}

// TestUserRepository_GetByID tests retrieving a user by ID
func TestUserRepository_GetByID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewUserRepository(testDB.Pool)

	t.Run("Success - User found", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "TEACHER")
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")
		userID := CreateTestUser(t, testDB.Pool, "teacher@example.com", roleID, &schoolID)

		user, err := repo.GetByID(ctx, userID)
		require.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, userID, user.ID)
		assert.Equal(t, "teacher@example.com", user.Email)
		assert.Equal(t, "Test User", user.Name)
		assert.Equal(t, roleID, user.RoleID)
		assert.NotNil(t, user.SchoolID)
		assert.Equal(t, schoolID, *user.SchoolID)
		assert.True(t, user.IsActive)
	})

	t.Run("Success - User without school", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "ADMIN")
		userID := CreateTestUser(t, testDB.Pool, "admin@example.com", roleID, nil)

		user, err := repo.GetByID(ctx, userID)
		require.NoError(t, err)
		assert.NotNil(t, user)
		assert.Nil(t, user.SchoolID)
	})

	t.Run("Error - User not found", func(t *testing.T) {
		ctx := context.Background()

		user, err := repo.GetByID(ctx, "non-existent-id")
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Contains(t, err.Error(), "user not found")
	})
}

// TestUserRepository_GetByEmail tests retrieving a user by email
func TestUserRepository_GetByEmail(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewUserRepository(testDB.Pool)

	t.Run("Success - User found by email", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "TEACHER")
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")
		userID := CreateTestUser(t, testDB.Pool, "teacher@example.com", roleID, &schoolID)

		user, err := repo.GetByEmail(ctx, "teacher@example.com")
		require.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, userID, user.ID)
		assert.Equal(t, "teacher@example.com", user.Email)
	})

	t.Run("Error - User not found by email", func(t *testing.T) {
		ctx := context.Background()

		user, err := repo.GetByEmail(ctx, "nonexistent@example.com")
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Contains(t, err.Error(), "user not found")
	})
}

// TestUserRepository_Update tests updating a user
func TestUserRepository_Update(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewUserRepository(testDB.Pool)

	t.Run("Success - Update user", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "TEACHER")
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")
		userID := CreateTestUser(t, testDB.Pool, "teacher@example.com", roleID, &schoolID)

		user, err := repo.GetByID(ctx, userID)
		require.NoError(t, err)

		user.Name = "Updated Name"
		user.UpdatedBy = &schoolID

		err = repo.Update(ctx, user)
		require.NoError(t, err)

		updatedUser, err := repo.GetByID(ctx, userID)
		require.NoError(t, err)
		assert.Equal(t, "Updated Name", updatedUser.Name)
	})

	t.Run("Error - User not found", func(t *testing.T) {
		ctx := context.Background()

		user := &domain.User{
			ID:   "non-existent-id",
			Name: "Non-existent",
		}

		err := repo.Update(ctx, user)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "user not found")
	})
}

// TestUserRepository_UpdateStatus tests updating user status
func TestUserRepository_UpdateStatus(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewUserRepository(testDB.Pool)

	t.Run("Success - Deactivate user", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "TEACHER")
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")
		userID := CreateTestUser(t, testDB.Pool, "teacher@example.com", roleID, &schoolID)

		err := repo.UpdateStatus(ctx, userID, false, nil, 0)
		require.NoError(t, err)

		user, err := repo.GetByID(ctx, userID)
		require.NoError(t, err)
		assert.False(t, user.IsActive)
	})

	t.Run("Success - Update failed login attempts", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "TEACHER")
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")
		userID := CreateTestUser(t, testDB.Pool, "teacher@example.com", roleID, &schoolID)

		err := repo.UpdateStatus(ctx, userID, true, nil, 3)
		require.NoError(t, err)

		user, err := repo.GetByID(ctx, userID)
		require.NoError(t, err)
		assert.Equal(t, 3, user.FailedLoginAttempts)
	})

	t.Run("Success - Lock user", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "TEACHER")
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")
		userID := CreateTestUser(t, testDB.Pool, "teacher@example.com", roleID, &schoolID)

		lockedUntil := time.Now().Add(1 * time.Hour).Format(time.RFC3339)
		err := repo.UpdateStatus(ctx, userID, true, StringPtr(lockedUntil), 5)
		require.NoError(t, err)

		user, err := repo.GetByID(ctx, userID)
		require.NoError(t, err)
		assert.NotNil(t, user.LockedUntil)
	})

	t.Run("Error - User not found", func(t *testing.T) {
		ctx := context.Background()

		err := repo.UpdateStatus(ctx, "non-existent-id", false, nil, 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "user not found")
	})
}

// TestUserRepository_List tests listing users with filters
func TestUserRepository_List(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewUserRepository(testDB.Pool)

	t.Run("Success - List all users", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "TEACHER")
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")

		CreateTestUser(t, testDB.Pool, "teacher1@example.com", roleID, &schoolID)
		CreateTestUser(t, testDB.Pool, "teacher2@example.com", roleID, &schoolID)
		CreateTestUser(t, testDB.Pool, "teacher3@example.com", roleID, &schoolID)

		users, err := repo.List(ctx, nil, nil, nil, 0, 0)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(users), 3)
	})

	t.Run("Success - Filter by school", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "TEACHER")
		schoolID1 := CreateTestSchool(t, testDB.Pool, "School 1")
		schoolID2 := CreateTestSchool(t, testDB.Pool, "School 2")

		CreateTestUser(t, testDB.Pool, "teacher1@example.com", roleID, &schoolID1)
		CreateTestUser(t, testDB.Pool, "teacher2@example.com", roleID, &schoolID1)
		CreateTestUser(t, testDB.Pool, "teacher3@example.com", roleID, &schoolID2)

		users, err := repo.List(ctx, &schoolID1, nil, nil, 0, 0)
		require.NoError(t, err)
		assert.Equal(t, 2, len(users))
	})

	t.Run("Success - Filter by role", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		teacherRoleID := CreateTestRole(t, testDB.Pool, "TEACHER")
		adminRoleID := CreateTestRole(t, testDB.Pool, "ADMIN")
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")

		CreateTestUser(t, testDB.Pool, "teacher1@example.com", teacherRoleID, &schoolID)
		CreateTestUser(t, testDB.Pool, "teacher2@example.com", teacherRoleID, &schoolID)
		CreateTestUser(t, testDB.Pool, "admin@example.com", adminRoleID, &schoolID)

		users, err := repo.List(ctx, nil, &teacherRoleID, nil, 0, 0)
		require.NoError(t, err)
		assert.Equal(t, 2, len(users))
	})

	t.Run("Success - Filter by active status", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "TEACHER")
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")

		userID1 := CreateTestUser(t, testDB.Pool, "active@example.com", roleID, &schoolID)
		userID2 := CreateTestUser(t, testDB.Pool, "inactive@example.com", roleID, &schoolID)

		// Deactivate second user
		repo.UpdateStatus(ctx, userID2, false, nil, 0)

		isActive := true
		users, err := repo.List(ctx, nil, nil, &isActive, 0, 0)
		require.NoError(t, err)

		// Should only have active user
		userIDs := make([]string, len(users))
		for i, u := range users {
			userIDs[i] = u.ID
		}
		assert.Contains(t, userIDs, userID1)
		assert.NotContains(t, userIDs, userID2)
	})

	t.Run("Success - With pagination", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "TEACHER")
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")

		for i := 0; i < 5; i++ {
			CreateTestUser(t, testDB.Pool, fmt.Sprintf("teacher%d@example.com", i), roleID, &schoolID)
		}

		users, err := repo.List(ctx, nil, nil, nil, 2, 0)
		require.NoError(t, err)
		assert.Equal(t, 2, len(users))

		users2, err := repo.List(ctx, nil, nil, nil, 2, 2)
		require.NoError(t, err)
		assert.Equal(t, 2, len(users2))
	})
}

// TestUserRepository_Count tests counting users with filters
func TestUserRepository_Count(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewUserRepository(testDB.Pool)

	t.Run("Success - Count all users", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "TEACHER")
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")

		CreateTestUser(t, testDB.Pool, "teacher1@example.com", roleID, &schoolID)
		CreateTestUser(t, testDB.Pool, "teacher2@example.com", roleID, &schoolID)
		CreateTestUser(t, testDB.Pool, "teacher3@example.com", roleID, &schoolID)

		count, err := repo.Count(ctx, nil, nil, nil)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, count, 3)
	})

	t.Run("Success - Count by school", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "TEACHER")
		schoolID1 := CreateTestSchool(t, testDB.Pool, "School 1")
		schoolID2 := CreateTestSchool(t, testDB.Pool, "School 2")

		CreateTestUser(t, testDB.Pool, "teacher1@example.com", roleID, &schoolID1)
		CreateTestUser(t, testDB.Pool, "teacher2@example.com", roleID, &schoolID1)
		CreateTestUser(t, testDB.Pool, "teacher3@example.com", roleID, &schoolID2)

		count, err := repo.Count(ctx, &schoolID1, nil, nil)
		require.NoError(t, err)
		assert.Equal(t, 2, count)
	})
}

// TestUserRepository_Delete tests soft deleting a user
func TestUserRepository_Delete(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewUserRepository(testDB.Pool)

	t.Run("Success - Soft delete user", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "TEACHER")
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")
		userID := CreateTestUser(t, testDB.Pool, "teacher@example.com", roleID, &schoolID)

		err := repo.Delete(ctx, userID)
		require.NoError(t, err)

		user, err := repo.GetByID(ctx, userID)
		require.NoError(t, err)
		assert.False(t, user.IsActive)
	})

	t.Run("Error - User not found", func(t *testing.T) {
		ctx := context.Background()

		err := repo.Delete(ctx, "non-existent-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "user not found")
	})
}

// TestUserRepository_GetUserSchoolID tests getting user's school ID
func TestUserRepository_GetUserSchoolID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewUserRepository(testDB.Pool)

	t.Run("Success - User has school", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "TEACHER")
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")
		userID := CreateTestUser(t, testDB.Pool, "teacher@example.com", roleID, &schoolID)

		userSchoolID, err := repo.GetUserSchoolID(ctx, userID)
		require.NoError(t, err)
		assert.NotNil(t, userSchoolID)
		assert.Equal(t, schoolID, *userSchoolID)
	})

	t.Run("Success - User without school", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "ADMIN")
		userID := CreateTestUser(t, testDB.Pool, "admin@example.com", roleID, nil)

		userSchoolID, err := repo.GetUserSchoolID(ctx, userID)
		require.NoError(t, err)
		assert.Nil(t, userSchoolID)
	})

	t.Run("Error - User not found", func(t *testing.T) {
		ctx := context.Background()

		_, err := repo.GetUserSchoolID(ctx, "non-existent-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "user not found")
	})
}

// TestUserRepository_GetUsersBySchool tests getting all users by school
func TestUserRepository_GetUsersBySchool(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewUserRepository(testDB.Pool)

	t.Run("Success - Get users by school", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		roleID := CreateTestRole(t, testDB.Pool, "TEACHER")
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")

		CreateTestUser(t, testDB.Pool, "teacher1@example.com", roleID, &schoolID)
		CreateTestUser(t, testDB.Pool, "teacher2@example.com", roleID, &schoolID)

		users, err := repo.GetUsersBySchool(ctx, schoolID)
		require.NoError(t, err)
		assert.Equal(t, 2, len(users))
	})
}

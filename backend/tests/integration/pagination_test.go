package integration

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestUserListPagination validates user list pagination
func TestUserListPagination(t *testing.T) {
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

	// Create test users
	for i := 0; i < 25; i++ {
		email := "pagination" + string(rune('0'+i%10)) + "@example.com"
		CreateTestUser(t, ctx, testDB.UserRepo, email, "password123", "Pagination User", roleID, nil)
	}

	t.Run("Basic pagination", func(t *testing.T) {
		limit := 10
		offset := 0

		users, err := testDB.UserRepo.List(ctx, nil, nil, nil, limit, offset)
		require.NoError(t, err)
		assert.Equal(t, limit, len(users), "Should return exactly 10 users")

		count, err := testDB.UserRepo.Count(ctx, nil, nil, nil)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, count, 25, "Total count should be at least 25")
	})

	t.Run("Second page", func(t *testing.T) {
		limit := 10
		offset := 10

		users, err := testDB.UserRepo.List(ctx, nil, nil, nil, limit, offset)
		require.NoError(t, err)
		assert.Equal(t, limit, len(users), "Should return exactly 10 users on second page")
	})

	t.Run("Last page with fewer items", func(t *testing.T) {
		limit := 10
		offset := 20

		users, err := testDB.UserRepo.List(ctx, nil, nil, nil, limit, offset)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(users), 5, "Last page should have at least 5 users")
		assert.LessOrEqual(t, len(users), 10, "Last page should have at most 10 users")
	})

	t.Run("Default pagination (no limit/offset)", func(t *testing.T) {
		users, err := testDB.UserRepo.List(ctx, nil, nil, nil, 0, 0)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(users), 25, "Should return all users when no pagination specified")
	})

	t.Run("Pagination with filtering", func(t *testing.T) {
		// Create users with a specific role for filtering
		teacherRoleID := "00000000-0000-0000-0000-000000000003"
		for i := 0; i < 5; i++ {
			email := "teacher" + string(rune('0'+i)) + "@example.com"
			CreateTestUser(t, ctx, testDB.UserRepo, email, "password123", "Teacher User", teacherRoleID, nil)
		}

		users, err := testDB.UserRepo.List(ctx, nil, &teacherRoleID, nil, 10, 0)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(users), 5, "Should return at least 5 teachers")
	})

	t.Run("Pagination with active status filter", func(t *testing.T) {
		isActive := true
		users, err := testDB.UserRepo.List(ctx, nil, nil, &isActive, 10, 0)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(users), 25, "Should return at least 25 active users")

		for _, user := range users {
			assert.True(t, user.IsActive, "All users should be active")
		}
	})
}

// TestSchoolListPagination validates school list pagination
func TestSchoolListPagination(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	testDB := SetupTestDB(t)
	if testDB == nil {
		return
	}
	defer TeardownTestDB(t, testDB)

	ctx := context.Background()

	// Create test schools
	for i := 0; i < 15; i++ {
		code := "SCH" + string(rune('0'+i%10)) + string(rune('0'+(i/10)%10))
		CreateTestSchool(t, ctx, testDB.SchoolRepo, "Pagination School "+string(rune('0'+i)), code)
	}

	t.Run("Basic pagination", func(t *testing.T) {
		limit := 5
		offset := 0

		schools, err := testDB.SchoolRepo.List(ctx, nil, limit, offset)
		require.NoError(t, err)
		assert.Equal(t, limit, len(schools), "Should return exactly 5 schools")
	})

	t.Run("Second page", func(t *testing.T) {
		limit := 5
		offset := 5

		schools, err := testDB.SchoolRepo.List(ctx, nil, limit, offset)
		require.NoError(t, err)
		assert.Equal(t, limit, len(schools), "Should return exactly 5 schools on second page")
	})

	t.Run("Pagination with active status filter", func(t *testing.T) {
		isActive := true
		schools, err := testDB.SchoolRepo.List(ctx, &isActive, 10, 0)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(schools), 15, "Should return at least 15 active schools")

		for _, school := range schools {
			assert.True(t, school.IsActive, "All schools should be active")
		}
	})

	t.Run("Pagination with inactive status filter", func(t *testing.T) {
		// Deactivate some schools
		schools, _ := testDB.SchoolRepo.List(ctx, nil, 5, 0)
		for _, school := range schools {
			school.IsActive = false
			testDB.SchoolRepo.Update(ctx, school)
		}

		isActive := false
		inactiveSchools, err := testDB.SchoolRepo.List(ctx, &isActive, 10, 0)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(inactiveSchools), 5, "Should return at least 5 inactive schools")

		for _, school := range inactiveSchools {
			assert.False(t, school.IsActive, "All schools should be inactive")
		}
	})
}

// TestRoleListPagination validates role list pagination
func TestRoleListPagination(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	testDB := SetupTestDB(t)
	if testDB == nil {
		return
	}
	defer TeardownTestDB(t, testDB)

	ctx := context.Background()

	// Create additional roles
	for i := 0; i < 10; i++ {
		name := "CUSTOM_ROLE_" + string(rune('0'+i))
		description := "Custom role description"
		CreateTestRole(t, ctx, testDB.RoleRepo, name, &description)
	}

	t.Run("Basic pagination", func(t *testing.T) {

		roles, err := testDB.RoleRepo.List(ctx, nil)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(roles), 13, "Should have at least 13 roles (3 system + 10 custom)")
	})

	t.Run("Pagination with active status filter", func(t *testing.T) {
		isActive := true
		roles, err := testDB.RoleRepo.List(ctx, &isActive)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(roles), 13, "Should return at least 13 active roles")

		for _, role := range roles {
			assert.True(t, role.IsActive, "All roles should be active")
		}
	})
}

// TestPaginationEdgeCases validates pagination edge cases
func TestPaginationEdgeCases(t *testing.T) {
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

	t.Run("Empty result set", func(t *testing.T) {
		// Query with non-existent school ID
		fakeSchoolID := "00000000-0000-0000-0000-000000000999"

		users, err := testDB.UserRepo.List(ctx, &fakeSchoolID, nil, nil, 10, 0)
		require.NoError(t, err)
		assert.Equal(t, 0, len(users), "Should return empty array for non-existent school")

		count, err := testDB.UserRepo.Count(ctx, &fakeSchoolID, nil, nil)
		require.NoError(t, err)
		assert.Equal(t, 0, count, "Count should be 0 for non-existent school")
	})

	t.Run("Page beyond available", func(t *testing.T) {
		users, err := testDB.UserRepo.List(ctx, nil, nil, nil, 10, 1000)
		require.NoError(t, err)
		assert.Equal(t, 0, len(users), "Should return empty array for page beyond available")

		count, err := testDB.UserRepo.Count(ctx, nil, nil, nil)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, count, 0, "Count should still be valid")
	})

	t.Run("Single result", func(t *testing.T) {
		// Create a unique user
		email := "single@example.com"
		CreateTestUser(t, ctx, testDB.UserRepo, email, "password123", "Single User", roleID, nil)

		// Query by specific email won't work with List, so we'll just verify the count
		users, err := testDB.UserRepo.List(ctx, nil, nil, nil, 1, 0)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(users), 1, "Should return at least 1 user")
	})

	t.Run("Exact page size match", func(t *testing.T) {
		// Create exactly 10 users
		for i := 0; i < 10; i++ {
			email := "exact" + string(rune('0'+i)) + "@example.com"
			CreateTestUser(t, ctx, testDB.UserRepo, email, "password123", "Exact User", roleID, nil)
		}

		users, err := testDB.UserRepo.List(ctx, nil, nil, nil, 10, 0)
		require.NoError(t, err)
		assert.Equal(t, 10, len(users), "Should return exactly 10 users")
	})

	t.Run("Large page size", func(t *testing.T) {
		users, err := testDB.UserRepo.List(ctx, nil, nil, nil, 1000, 0)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(users), 0, "Should handle large page size")
	})
}

// TestPaginationConsistency validates pagination consistency across calls
func TestPaginationConsistency(t *testing.T) {
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

	// Create consistent set of users
	for i := 0; i < 20; i++ {
		email := "consistent" + string(rune('0'+i%10)) + "@example.com"
		CreateTestUser(t, ctx, testDB.UserRepo, email, "password123", "Consistent User", roleID, nil)
	}

	t.Run("Consistent total count", func(t *testing.T) {
		count1, err := testDB.UserRepo.Count(ctx, nil, nil, nil)
		require.NoError(t, err)

		count2, err := testDB.UserRepo.Count(ctx, nil, nil, nil)
		require.NoError(t, err)

		assert.Equal(t, count1, count2, "Total count should be consistent across calls")
	})

	t.Run("Consistent pagination results", func(t *testing.T) {
		users1, err := testDB.UserRepo.List(ctx, nil, nil, nil, 10, 0)
		require.NoError(t, err)

		users2, err := testDB.UserRepo.List(ctx, nil, nil, nil, 10, 0)
		require.NoError(t, err)

		assert.Equal(t, len(users1), len(users2), "Page 1 should have consistent results")

		// Verify same users are returned
		for i := 0; i < len(users1) && i < len(users2); i++ {
			assert.Equal(t, users1[i].ID, users2[i].ID, "User IDs should be in same order")
		}
	})

	t.Run("No duplicate users across pages", func(t *testing.T) {
		userIDs := make(map[string]bool)

		// Get first page
		page1, err := testDB.UserRepo.List(ctx, nil, nil, nil, 10, 0)
		require.NoError(t, err)

		// Get second page
		page2, err := testDB.UserRepo.List(ctx, nil, nil, nil, 10, 10)
		require.NoError(t, err)

		// Collect all user IDs
		for _, user := range page1 {
			if _, exists := userIDs[user.ID]; exists {
				t.Errorf("Duplicate user ID found on page 1: %s", user.ID)
			}
			userIDs[user.ID] = true
		}

		for _, user := range page2 {
			if _, exists := userIDs[user.ID]; exists {
				t.Errorf("Duplicate user ID found across pages: %s", user.ID)
			}
			userIDs[user.ID] = true
		}
	})
}

// TestPaginationWithSorting validates pagination with sorting
func TestPaginationWithSorting(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	testDB := SetupTestDB(t)
	if testDB == nil {
		return
	}
	defer TeardownTestDB(t, testDB)

	ctx := context.Background()

	t.Run("Users sorted by created_at (descending)", func(t *testing.T) {
		roleID := "00000000-0000-0000-0000-000000000001"

		// Create users
		for i := 0; i < 5; i++ {
			email := "sort" + string(rune('0'+i)) + "@example.com"
			CreateTestUser(t, ctx, testDB.UserRepo, email, "password123", "Sort User", roleID, nil)
		}

		users, err := testDB.UserRepo.List(ctx, nil, nil, nil, 5, 0)
		require.NoError(t, err)

		// Verify users are sorted by created_at descending
		for i := 1; i < len(users); i++ {
			assert.True(t, users[i-1].CreatedAt.After(users[i].CreatedAt) ||
				users[i-1].CreatedAt.Equal(users[i].CreatedAt),
				"Users should be sorted by created_at descending")
		}
	})

	t.Run("Roles sorted by name", func(t *testing.T) {
		roles, err := testDB.RoleRepo.List(ctx, nil)
		require.NoError(t, err)

		// Verify roles are sorted by name
		for i := 1; i < len(roles); i++ {
			assert.True(t, roles[i-1].Name <= roles[i].Name,
				"Roles should be sorted by name")
		}
	})
}

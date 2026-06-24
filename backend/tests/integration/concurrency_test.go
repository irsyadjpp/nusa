package integration

import (
	"context"
	"sync"
	"testing"

	"github.com/nusa/backend/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestConcurrentUserCreation validates concurrent user creation
func TestConcurrentUserCreation(t *testing.T) {
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

	t.Run("Concurrent user creation with same email", func(t *testing.T) {
		email := "concurrent-same@example.com"
		password := "password123"
		name := "Concurrent User"

		var wg sync.WaitGroup
		numGoroutines := 10
		successCount := 0
		var mu sync.Mutex

		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				user := &domain.User{
					ID:                  domain.NewID(),
					Email:               email,
					PasswordHash:        password, // Simplified for concurrency test
					Name:                name,
					RoleID:              roleID,
					IsActive:            true,
					FailedLoginAttempts: 0,
					CreatedBy:           &roleID,
					UpdatedBy:           &roleID,
				}

				err := testDB.UserRepo.Create(ctx, user)
				mu.Lock()
				if err == nil {
					successCount++
				}
				mu.Unlock()
			}()
		}

		wg.Wait()

		// Due to unique constraint on email, only one should succeed
		assert.Equal(t, 1, successCount, "Only one user creation should succeed with duplicate email")

		// Verify only one user exists
		user, err := testDB.UserRepo.GetByEmail(ctx, email)
		require.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, email, user.Email)
	})

	t.Run("Concurrent user creation with different emails", func(t *testing.T) {
		var wg sync.WaitGroup
		numGoroutines := 10
		successCount := 0
		var mu sync.Mutex

		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func(idx int) {
				defer wg.Done()
				email := "concurrent-different" + string(rune('0'+idx)) + "@example.com"
				user := CreateTestUser(t, ctx, testDB.UserRepo, email, "password123", "Concurrent User", roleID, nil)

				mu.Lock()
				if user != nil {
					successCount++
				}
				mu.Unlock()
			}(i)
		}

		wg.Wait()

		// All should succeed with different emails
		assert.Equal(t, numGoroutines, successCount, "All user creations should succeed with different emails")
	})
}

// TestConcurrentUpdate validates concurrent update operations
func TestConcurrentUpdate(t *testing.T) {
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

	t.Run("Concurrent user update", func(t *testing.T) {
		user := CreateTestUser(t, ctx, testDB.UserRepo, "update-concurrent@example.com", "password123", "Update Concurrent", roleID, nil)

		var wg sync.WaitGroup
		numGoroutines := 10
		successCount := 0
		var mu sync.Mutex

		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func(idx int) {
				defer wg.Done()
				user.Name = "Updated Name " + string(rune('0'+idx))
				err := testDB.UserRepo.Update(ctx, user)

				mu.Lock()
				if err == nil {
					successCount++
				}
				mu.Unlock()
			}(i)
		}

		wg.Wait()

		// All updates should succeed (last write wins)
		assert.Equal(t, numGoroutines, successCount, "All updates should succeed")

		// Verify final state
		retrieved, err := testDB.UserRepo.GetByID(ctx, user.ID)
		require.NoError(t, err)
		assert.Contains(t, retrieved.Name, "Updated Name")
	})

	t.Run("Concurrent school update", func(t *testing.T) {
		school := CreateTestSchool(t, ctx, testDB.SchoolRepo, "Update School Concurrent", "USC001")

		var wg sync.WaitGroup
		numGoroutines := 10
		successCount := 0
		var mu sync.Mutex

		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func(idx int) {
				defer wg.Done()
				school.Name = "Updated School Name " + string(rune('0'+idx))
				err := testDB.SchoolRepo.Update(ctx, school)

				mu.Lock()
				if err == nil {
					successCount++
				}
				mu.Unlock()
			}(i)
		}

		wg.Wait()

		// All updates should succeed (last write wins)
		assert.Equal(t, numGoroutines, successCount, "All updates should succeed")

		// Verify final state
		retrieved, err := testDB.SchoolRepo.GetByID(ctx, school.ID)
		require.NoError(t, err)
		assert.Contains(t, retrieved.Name, "Updated School Name")
	})
}

// TestConcurrentDelete validates concurrent delete operations
func TestConcurrentDelete(t *testing.T) {
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

	t.Run("Concurrent user delete", func(t *testing.T) {
		user := CreateTestUser(t, ctx, testDB.UserRepo, "delete-concurrent@example.com", "password123", "Delete Concurrent", roleID, nil)

		var wg sync.WaitGroup
		numGoroutines := 10
		successCount := 0
		var mu sync.Mutex

		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				err := testDB.UserRepo.Delete(ctx, user.ID)

				mu.Lock()
				if err == nil {
					successCount++
				}
				mu.Unlock()
			}()
		}

		wg.Wait()

		// Only first delete should succeed, others should get "user not found"
		assert.Equal(t, 1, successCount, "Only first delete should succeed")

		// Verify user is soft deleted
		retrieved, err := testDB.UserRepo.GetByID(ctx, user.ID)
		require.NoError(t, err)
		assert.False(t, retrieved.IsActive, "User should be soft deleted")

		// Additional deletes should fail (user not found or already deleted)
		err = testDB.UserRepo.Delete(ctx, user.ID)
		assert.Error(t, err)
	})
}

// TestConcurrentRead validates concurrent read operations
func TestConcurrentRead(t *testing.T) {
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

	t.Run("Concurrent user reads", func(t *testing.T) {
		user := CreateTestUser(t, ctx, testDB.UserRepo, "read-concurrent@example.com", "password123", "Read Concurrent", roleID, nil)

		var wg sync.WaitGroup
		numGoroutines := 10
		successCount := 0
		var mu sync.Mutex

		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				_, err := testDB.UserRepo.GetByID(ctx, user.ID)

				mu.Lock()
				if err == nil {
					successCount++
				}
				mu.Unlock()
			}()
		}

		wg.Wait()

		// All reads should succeed
		assert.Equal(t, numGoroutines, successCount, "All reads should succeed")
	})

	t.Run("Concurrent list operations", func(t *testing.T) {
		// Create some users
		for i := 0; i < 5; i++ {
			CreateTestUser(t, ctx, testDB.UserRepo, "list"+string(rune('0'+i))+"@example.com", "password123", "List User", roleID, nil)
		}

		var wg sync.WaitGroup
		numGoroutines := 10
		successCount := 0
		var mu sync.Mutex

		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				_, err := testDB.UserRepo.List(ctx, nil, nil, nil, 10, 0)

				mu.Lock()
				if err == nil {
					successCount++
				}
				mu.Unlock()
			}()
		}

		wg.Wait()

		// All list operations should succeed
		assert.Equal(t, numGoroutines, successCount, "All list operations should succeed")
	})
}

// TestConcurrentMixedOperations validates mixed concurrent operations
func TestConcurrentMixedOperations(t *testing.T) {
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

	t.Run("Mixed create, read, update, and delete", func(t *testing.T) {
		var wg sync.WaitGroup
		numGoroutines := 20
		var createSuccess, readSuccess, updateSuccess, deleteSuccess int
		var mu sync.Mutex

		// Create initial user for update/delete operations
		user := CreateTestUser(t, ctx, testDB.UserRepo, "mixed@example.com", "password123", "Mixed User", roleID, nil)

		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func(idx int) {
				defer wg.Done()

				switch idx % 4 {
				case 0: // Create
					email := "mixed" + string(rune('0'+idx)) + "@example.com"
					CreateTestUser(t, ctx, testDB.UserRepo, email, "password123", "Mixed Create", roleID, nil)
					mu.Lock()
					createSuccess++
					mu.Unlock()
				case 1: // Read
					_, err := testDB.UserRepo.GetByID(ctx, user.ID)
					if err == nil {
						mu.Lock()
						readSuccess++
						mu.Unlock()
					}
				case 2: // Update
					user.Name = "Updated Name " + string(rune('0'+idx))
					err := testDB.UserRepo.Update(ctx, user)
					if err == nil {
						mu.Lock()
						updateSuccess++
						mu.Unlock()
					}
				case 3: // Delete
					// Only attempt delete once
					if idx == 3 {
						err := testDB.UserRepo.Delete(ctx, user.ID)
						if err == nil {
							mu.Lock()
							deleteSuccess++
							mu.Unlock()
						}
					}
				}
			}(i)
		}

		wg.Wait()

		// Verify operations completed
		assert.Greater(t, createSuccess, 0, "Some creates should succeed")
		assert.Greater(t, readSuccess, 0, "Some reads should succeed")
		assert.Greater(t, updateSuccess, 0, "Some updates should succeed")
		assert.Equal(t, 1, deleteSuccess, "Delete should succeed once")
	})
}

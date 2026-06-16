package repository

import (
	"context"
	"fmt"
	"testing"

	"github.com/nusa/backend/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestSchoolRepository_Create tests creating a new school
func TestSchoolRepository_Create(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewSchoolRepository(testDB.Pool)

	t.Run("Success - Create school with all fields", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		address := "123 Test Street"
		phone := "123-456-7890"
		email := "school@example.com"
		school := &domain.School{
			ID:        "school-001",
			Name:      "Test School",
			Code:      "SCH001",
			Address:   &address,
			Phone:     &phone,
			Email:     &email,
			IsActive:  true,
			CreatedBy: StringPtr("system"),
			UpdatedBy: StringPtr("system"),
		}

		err := repo.Create(ctx, school)
		require.NoError(t, err)
		assert.NotEmpty(t, school.ID)
	})

	t.Run("Success - Create school with minimal fields", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		school := &domain.School{
			ID:        "school-002",
			Name:      "Minimal School",
			Code:      "SCH002",
			Address:   nil,
			Phone:     nil,
			Email:     nil,
			IsActive:  true,
			CreatedBy: StringPtr("system"),
			UpdatedBy: StringPtr("system"),
		}

		err := repo.Create(ctx, school)
		require.NoError(t, err)
	})

	t.Run("Error - Duplicate code", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()

		school1 := &domain.School{
			ID:        "school-003",
			Name:      "School 1",
			Code:      "DUPLICATE",
			IsActive:  true,
			CreatedBy: StringPtr("system"),
			UpdatedBy: StringPtr("system"),
		}

		err := repo.Create(ctx, school1)
		require.NoError(t, err)

		school2 := &domain.School{
			ID:        "school-004",
			Name:      "School 2",
			Code:      "DUPLICATE", // Same code
			IsActive:  true,
			CreatedBy: StringPtr("system"),
			UpdatedBy: StringPtr("system"),
		}

		err = repo.Create(ctx, school2)
		assert.Error(t, err)
	})
}

// TestSchoolRepository_GetByID tests retrieving a school by ID
func TestSchoolRepository_GetByID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewSchoolRepository(testDB.Pool)

	t.Run("Success - School found", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")

		school, err := repo.GetByID(ctx, schoolID)
		require.NoError(t, err)
		assert.NotNil(t, school)
		assert.Equal(t, schoolID, school.ID)
		assert.Equal(t, "Test School", school.Name)
		assert.True(t, school.IsActive)
	})

	t.Run("Success - School with all fields", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		address := "123 Test Street"
		phone := "123-456-7890"
		email := "school@example.com"

		schoolID := CreateTestSchool(t, testDB.Pool, "Full School")

		// Update with full details
		school, err := repo.GetByID(ctx, schoolID)
		require.NoError(t, err)
		school.Address = &address
		school.Phone = &phone
		school.Email = &email
		err = repo.Update(ctx, school)
		require.NoError(t, err)

		school, err = repo.GetByID(ctx, schoolID)
		require.NoError(t, err)
		assert.NotNil(t, school.Address)
		assert.NotNil(t, school.Phone)
		assert.NotNil(t, school.Email)
	})

	t.Run("Error - School not found", func(t *testing.T) {
		ctx := context.Background()

		school, err := repo.GetByID(ctx, "non-existent-id")
		assert.Error(t, err)
		assert.Nil(t, school)
		assert.Contains(t, err.Error(), "school not found")
	})
}

// TestSchoolRepository_GetByCode tests retrieving a school by code
func TestSchoolRepository_GetByCode(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewSchoolRepository(testDB.Pool)

	t.Run("Success - School found by code", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")

		school, err := repo.GetByID(ctx, schoolID)
		require.NoError(t, err)
		code := school.Code

		schoolByCode, err := repo.GetByCode(ctx, code)
		require.NoError(t, err)
		assert.NotNil(t, schoolByCode)
		assert.Equal(t, schoolID, schoolByCode.ID)
		assert.Equal(t, code, schoolByCode.Code)
	})

	t.Run("Error - School not found by code", func(t *testing.T) {
		ctx := context.Background()

		school, err := repo.GetByCode(ctx, "NONEXISTENT")
		assert.Error(t, err)
		assert.Nil(t, school)
		assert.Contains(t, err.Error(), "school not found")
	})
}

// TestSchoolRepository_Update tests updating a school
func TestSchoolRepository_Update(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewSchoolRepository(testDB.Pool)

	t.Run("Success - Update school name", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		schoolID := CreateTestSchool(t, testDB.Pool, "Old Name")

		school, err := repo.GetByID(ctx, schoolID)
		require.NoError(t, err)

		school.Name = "New Name"

		err = repo.Update(ctx, school)
		require.NoError(t, err)

		updatedSchool, err := repo.GetByID(ctx, schoolID)
		require.NoError(t, err)
		assert.Equal(t, "New Name", updatedSchool.Name)
	})

	t.Run("Success - Update school address", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")

		school, err := repo.GetByID(ctx, schoolID)
		require.NoError(t, err)

		newAddress := "456 New Address"
		school.Address = &newAddress

		err = repo.Update(ctx, school)
		require.NoError(t, err)

		updatedSchool, err := repo.GetByID(ctx, schoolID)
		require.NoError(t, err)
		assert.NotNil(t, updatedSchool.Address)
		assert.Equal(t, newAddress, *updatedSchool.Address)
	})

	t.Run("Success - Update school contact info", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")

		school, err := repo.GetByID(ctx, schoolID)
		require.NoError(t, err)

		newPhone := "987-654-3210"
		newEmail := "newemail@example.com"
		school.Phone = &newPhone
		school.Email = &newEmail

		err = repo.Update(ctx, school)
		require.NoError(t, err)

		updatedSchool, err := repo.GetByID(ctx, schoolID)
		require.NoError(t, err)
		assert.NotNil(t, updatedSchool.Phone)
		assert.NotNil(t, updatedSchool.Email)
		assert.Equal(t, newPhone, *updatedSchool.Phone)
		assert.Equal(t, newEmail, *updatedSchool.Email)
	})

	t.Run("Error - School not found", func(t *testing.T) {
		ctx := context.Background()

		school := &domain.School{
			ID:   "non-existent-id",
			Name: "Non-existent",
			Code: "NONEXIST",
		}

		err := repo.Update(ctx, school)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "school not found")
	})
}

// TestSchoolRepository_UpdateStatus tests updating school status
func TestSchoolRepository_UpdateStatus(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewSchoolRepository(testDB.Pool)

	t.Run("Success - Deactivate school", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")

		err := repo.UpdateStatus(ctx, schoolID, false)
		require.NoError(t, err)

		school, err := repo.GetByID(ctx, schoolID)
		require.NoError(t, err)
		assert.False(t, school.IsActive)
	})

	t.Run("Success - Reactivate school", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")

		// Deactivate
		repo.UpdateStatus(ctx, schoolID, false)

		// Reactivate
		err := repo.UpdateStatus(ctx, schoolID, true)
		require.NoError(t, err)

		school, err := repo.GetByID(ctx, schoolID)
		require.NoError(t, err)
		assert.True(t, school.IsActive)
	})

	t.Run("Error - School not found", func(t *testing.T) {
		ctx := context.Background()

		err := repo.UpdateStatus(ctx, "non-existent-id", false)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "school not found")
	})
}

// TestSchoolRepository_List tests listing schools with filters
func TestSchoolRepository_List(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewSchoolRepository(testDB.Pool)

	t.Run("Success - List all schools", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()

		CreateTestSchool(t, testDB.Pool, "School 1")
		CreateTestSchool(t, testDB.Pool, "School 2")
		CreateTestSchool(t, testDB.Pool, "School 3")

		schools, err := repo.List(ctx, nil, 0, 0)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(schools), 3)
	})

	t.Run("Success - Filter active schools", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()

		activeSchoolID := CreateTestSchool(t, testDB.Pool, "Active School")
		inactiveSchoolID := CreateTestSchool(t, testDB.Pool, "Inactive School")

		// Deactivate one school
		repo.UpdateStatus(ctx, inactiveSchoolID, false)

		isActive := true
		schools, err := repo.List(ctx, &isActive, 0, 0)
		require.NoError(t, err)

		schoolIDs := make([]string, len(schools))
		for i, s := range schools {
			schoolIDs[i] = s.ID
		}
		assert.Contains(t, schoolIDs, activeSchoolID)
		assert.NotContains(t, schoolIDs, inactiveSchoolID)
	})

	t.Run("Success - With pagination", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()

		for i := 0; i < 5; i++ {
			CreateTestSchool(t, testDB.Pool, fmt.Sprintf("School %d", i))
		}

		schools, err := repo.List(ctx, nil, 2, 0)
		require.NoError(t, err)
		assert.Equal(t, 2, len(schools))

		schools2, err := repo.List(ctx, nil, 2, 2)
		require.NoError(t, err)
		assert.Equal(t, 2, len(schools2))
	})
}

// TestSchoolRepository_Count tests counting schools with filters
func TestSchoolRepository_Count(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewSchoolRepository(testDB.Pool)

	t.Run("Success - Count all schools", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()

		CreateTestSchool(t, testDB.Pool, "School 1")
		CreateTestSchool(t, testDB.Pool, "School 2")
		CreateTestSchool(t, testDB.Pool, "School 3")

		count, err := repo.Count(ctx, nil)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, count, 3)
	})

	t.Run("Success - Count active schools", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()

		CreateTestSchool(t, testDB.Pool, "Active School 1")
		CreateTestSchool(t, testDB.Pool, "Active School 2")
		inactiveSchoolID := CreateTestSchool(t, testDB.Pool, "Inactive School")

		repo.UpdateStatus(ctx, inactiveSchoolID, false)

		isActive := true
		count, err := repo.Count(ctx, &isActive)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, count, 2)
	})
}

// TestSchoolRepository_Delete tests soft deleting a school
func TestSchoolRepository_Delete(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewSchoolRepository(testDB.Pool)

	t.Run("Success - Soft delete school", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")

		err := repo.Delete(ctx, schoolID)
		require.NoError(t, err)

		school, err := repo.GetByID(ctx, schoolID)
		require.NoError(t, err)
		assert.False(t, school.IsActive)
	})

	t.Run("Error - School not found", func(t *testing.T) {
		ctx := context.Background()

		err := repo.Delete(ctx, "non-existent-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "school not found")
	})
}

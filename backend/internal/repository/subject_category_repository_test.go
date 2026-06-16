package repository

import (
	"context"
	"testing"
	"time"

	"github.com/nusa/backend/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestSubjectCategoryRepository_CreateSubjectCategory tests creating a subject category
func TestSubjectCategoryRepository_CreateSubjectCategory(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewSubjectCategoryRepository(testDB.Pool)

	t.Run("Success - Create subject category", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		description := "Required subject category"
		createdBy := "system"
		now := time.Now()

		sc := &domain.SubjectCategory{
			ID:          "sc-001",
			Code:        "REQUIRED",
			Name:        "Required",
			Description: &description,
			IsMandatory: true,
			IsActive:    true,
			CreatedBy:   createdBy,
			CreatedAt:   now,
			UpdatedAt:   now,
		}

		err := repo.CreateSubjectCategory(ctx, sc)
		require.NoError(t, err)
	})
}

// TestSubjectCategoryRepository_GetSubjectCategoryByID tests retrieving by ID
func TestSubjectCategoryRepository_GetSubjectCategoryByID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewSubjectCategoryRepository(testDB.Pool)

	t.Run("Success - Category found", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		description := "Required subject category"
		now := time.Now()

		sc := &domain.SubjectCategory{
			ID:          "sc-001",
			Code:        "REQUIRED",
			Name:        "Required",
			Description: &description,
			IsMandatory: true,
			IsActive:    true,
			CreatedBy:   "system",
			CreatedAt:   now,
			UpdatedAt:   now,
		}

		err := repo.CreateSubjectCategory(ctx, sc)
		require.NoError(t, err)

		retrieved, err := repo.GetSubjectCategoryByID(ctx, sc.ID)
		require.NoError(t, err)
		assert.Equal(t, sc.ID, retrieved.ID)
		assert.Equal(t, sc.Name, retrieved.Name)
	})

	t.Run("Error - Category not found", func(t *testing.T) {
		ctx := context.Background()

		_, err := repo.GetSubjectCategoryByID(ctx, "non-existent-id")
		assert.Error(t, err)
	})
}

// TestSubjectCategoryRepository_GetAllSubjectCategories tests retrieving all categories
func TestSubjectCategoryRepository_GetAllSubjectCategories(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewSubjectCategoryRepository(testDB.Pool)

	t.Run("Success - Get all categories", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		now := time.Now()

		sc1 := &domain.SubjectCategory{
			ID:        "sc-001",
			Code:      "REQUIRED",
			Name:      "Required",
			IsMandatory: true,
			IsActive:  true,
			CreatedBy: "system",
			CreatedAt: now,
			UpdatedAt: now,
		}

		sc2 := &domain.SubjectCategory{
			ID:        "sc-002",
			Code:      "OPTIONAL",
			Name:      "Optional",
			IsMandatory: false,
			IsActive:  true,
			CreatedBy: "system",
			CreatedAt: now,
			UpdatedAt: now,
		}

		repo.CreateSubjectCategory(ctx, sc1)
		repo.CreateSubjectCategory(ctx, sc2)

		categories, err := repo.GetAllSubjectCategories(ctx)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(categories), 2)
	})
}

// TestSubjectCategoryRepository_GetActiveSubjectCategories tests retrieving active categories
func TestSubjectCategoryRepository_GetActiveSubjectCategories(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewSubjectCategoryRepository(testDB.Pool)

	t.Run("Success - Get active categories", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		now := time.Now()

		sc1 := &domain.SubjectCategory{
			ID:        "sc-001",
			Code:      "ACTIVE",
			Name:      "Active",
			IsMandatory: true,
			IsActive:  true,
			CreatedBy: "system",
			CreatedAt: now,
			UpdatedAt: now,
		}

		sc2 := &domain.SubjectCategory{
			ID:        "sc-002",
			Code:      "INACTIVE",
			Name:      "Inactive",
			IsMandatory: false,
			IsActive:  false,
			CreatedBy: "system",
			CreatedAt: now,
			UpdatedAt: now,
		}

		repo.CreateSubjectCategory(ctx, sc1)
		repo.CreateSubjectCategory(ctx, sc2)

		categories, err := repo.GetActiveSubjectCategories(ctx)
		require.NoError(t, err)

		codes := make([]string, len(categories))
		for i, c := range categories {
			codes[i] = c.Code
		}
		assert.Contains(t, codes, "ACTIVE")
		assert.NotContains(t, codes, "INACTIVE")
	})
}

// TestSubjectCategoryRepository_UpdateSubjectCategory tests updating a category
func TestSubjectCategoryRepository_UpdateSubjectCategory(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewSubjectCategoryRepository(testDB.Pool)

	t.Run("Success - Update category", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		now := time.Now()

		sc := &domain.SubjectCategory{
			ID:        "sc-001",
			Code:      "REQUIRED",
			Name:      "Required",
			IsMandatory: true,
			IsActive:  true,
			CreatedBy: "system",
			CreatedAt: now,
			UpdatedAt: now,
		}

		err := repo.CreateSubjectCategory(ctx, sc)
		require.NoError(t, err)

		sc.Name = "Updated Required"
		sc.UpdatedBy = StringPtr("admin")
		sc.UpdatedAt = time.Now()

		err = repo.UpdateSubjectCategory(ctx, sc)
		require.NoError(t, err)

		updated, err := repo.GetSubjectCategoryByID(ctx, sc.ID)
		require.NoError(t, err)
		assert.Equal(t, "Updated Required", updated.Name)
	})

	t.Run("Error - Category not found", func(t *testing.T) {
		ctx := context.Background()

		sc := &domain.SubjectCategory{
			ID:   "non-existent-id",
			Name: "Non-existent",
		}

		err := repo.UpdateSubjectCategory(ctx, sc)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "subject category not found")
	})
}

// TestSubjectCategoryRepository_DeleteSubjectCategory tests deleting a category
func TestSubjectCategoryRepository_DeleteSubjectCategory(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewSubjectCategoryRepository(testDB.Pool)

	t.Run("Success - Delete category", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		now := time.Now()

		sc := &domain.SubjectCategory{
			ID:        "sc-001",
			Code:      "DELETE",
			Name:      "Delete Me",
			IsMandatory: false,
			IsActive:  true,
			CreatedBy: "system",
			CreatedAt: now,
			UpdatedAt: now,
		}

		err := repo.CreateSubjectCategory(ctx, sc)
		require.NoError(t, err)

		err = repo.DeleteSubjectCategory(ctx, sc.ID)
		require.NoError(t, err)

		_, err = repo.GetSubjectCategoryByID(ctx, sc.ID)
		assert.Error(t, err)
	})

	t.Run("Error - Category not found", func(t *testing.T) {
		ctx := context.Background()

		err := repo.DeleteSubjectCategory(ctx, "non-existent-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "subject category not found")
	})
}

// TestSubjectCategoryRepository_CheckCodeExists tests checking code uniqueness
func TestSubjectCategoryRepository_CheckCodeExists(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewSubjectCategoryRepository(testDB.Pool)

	t.Run("Success - Code exists", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		now := time.Now()

		sc := &domain.SubjectCategory{
			ID:        "sc-001",
			Code:      "UNIQUE",
			Name:      "Unique",
			IsMandatory: true,
			IsActive:  true,
			CreatedBy: "system",
			CreatedAt: now,
			UpdatedAt: now,
		}

		repo.CreateSubjectCategory(ctx, sc)

		exists, err := repo.CheckCodeExists(ctx, "UNIQUE", "")
		require.NoError(t, err)
		assert.True(t, exists)
	})

	t.Run("Success - Code does not exist", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()

		exists, err := repo.CheckCodeExists(ctx, "NONEXISTENT", "")
		require.NoError(t, err)
		assert.False(t, exists)
	})
}

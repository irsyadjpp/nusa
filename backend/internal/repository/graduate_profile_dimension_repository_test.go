package repository

import (
	"context"
	"testing"
	"time"

	"github.com/nusa/backend/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestGraduateProfileDimensionRepository_CreateGraduateProfileDimension tests creating a dimension
func TestGraduateProfileDimensionRepository_CreateGraduateProfileDimension(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewGraduateProfileDimensionRepository(testDB.Pool)

	t.Run("Success - Create dimension", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		description := "Character development dimension"
		now := time.Now()

		gpd := &domain.GraduateProfileDimension{
			ID:             "gpd-001",
			Code:           "CHARACTER",
			Name:           "Character",
			Description:    &description,
			SequenceNumber: 1,
			IsActive:       true,
			CreatedBy:      "system",
			CreatedAt:      now,
			UpdatedAt:      now,
		}

		err := repo.CreateGraduateProfileDimension(ctx, gpd)
		require.NoError(t, err)
	})
}

// TestGraduateProfileDimensionRepository_GetGraduateProfileDimensionByID tests retrieving by ID
func TestGraduateProfileDimensionRepository_GetGraduateProfileDimensionByID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewGraduateProfileDimensionRepository(testDB.Pool)

	t.Run("Success - Dimension found", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		description := "Character development dimension"
		now := time.Now()

		gpd := &domain.GraduateProfileDimension{
			ID:             "gpd-001",
			Code:           "CHARACTER",
			Name:           "Character",
			Description:    &description,
			SequenceNumber: 1,
			IsActive:       true,
			CreatedBy:      "system",
			CreatedAt:      now,
			UpdatedAt:      now,
		}

		err := repo.CreateGraduateProfileDimension(ctx, gpd)
		require.NoError(t, err)

		retrieved, err := repo.GetGraduateProfileDimensionByID(ctx, gpd.ID)
		require.NoError(t, err)
		assert.Equal(t, gpd.ID, retrieved.ID)
		assert.Equal(t, gpd.Name, retrieved.Name)
	})

	t.Run("Error - Dimension not found", func(t *testing.T) {
		ctx := context.Background()

		_, err := repo.GetGraduateProfileDimensionByID(ctx, "non-existent-id")
		assert.Error(t, err)
	})
}

// TestGraduateProfileDimensionRepository_GetAllGraduateProfileDimensions tests retrieving all dimensions
func TestGraduateProfileDimensionRepository_GetAllGraduateProfileDimensions(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewGraduateProfileDimensionRepository(testDB.Pool)

	t.Run("Success - Get all dimensions", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		now := time.Now()

		gpd1 := &domain.GraduateProfileDimension{
			ID:             "gpd-001",
			Code:           "CHARACTER",
			Name:           "Character",
			SequenceNumber: 1,
			IsActive:       true,
			CreatedBy:      "system",
			CreatedAt:      now,
			UpdatedAt:      now,
		}

		gpd2 := &domain.GraduateProfileDimension{
			ID:             "gpd-002",
			Code:           "COMPETENCE",
			Name:           "Competence",
			SequenceNumber: 2,
			IsActive:       true,
			CreatedBy:      "system",
			CreatedAt:      now,
			UpdatedAt:      now,
		}

		repo.CreateGraduateProfileDimension(ctx, gpd1)
		repo.CreateGraduateProfileDimension(ctx, gpd2)

		dimensions, err := repo.GetAllGraduateProfileDimensions(ctx)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(dimensions), 2)
	})
}

// TestGraduateProfileDimensionRepository_GetActiveGraduateProfileDimensions tests retrieving active dimensions
func TestGraduateProfileDimensionRepository_GetActiveGraduateProfileDimensions(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewGraduateProfileDimensionRepository(testDB.Pool)

	t.Run("Success - Get active dimensions", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		now := time.Now()

		gpd1 := &domain.GraduateProfileDimension{
			ID:             "gpd-001",
			Code:           "ACTIVE",
			Name:           "Active",
			SequenceNumber: 1,
			IsActive:       true,
			CreatedBy:      "system",
			CreatedAt:      now,
			UpdatedAt:      now,
		}

		gpd2 := &domain.GraduateProfileDimension{
			ID:             "gpd-002",
			Code:           "INACTIVE",
			Name:           "Inactive",
			SequenceNumber: 2,
			IsActive:       false,
			CreatedBy:      "system",
			CreatedAt:      now,
			UpdatedAt:      now,
		}

		repo.CreateGraduateProfileDimension(ctx, gpd1)
		repo.CreateGraduateProfileDimension(ctx, gpd2)

		dimensions, err := repo.GetActiveGraduateProfileDimensions(ctx)
		require.NoError(t, err)

		codes := make([]string, len(dimensions))
		for i, d := range dimensions {
			codes[i] = d.Code
		}
		assert.Contains(t, codes, "ACTIVE")
		assert.NotContains(t, codes, "INACTIVE")
	})
}

// TestGraduateProfileDimensionRepository_UpdateGraduateProfileDimension tests updating a dimension
func TestGraduateProfileDimensionRepository_UpdateGraduateProfileDimension(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewGraduateProfileDimensionRepository(testDB.Pool)

	t.Run("Success - Update dimension", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		now := time.Now()

		gpd := &domain.GraduateProfileDimension{
			ID:             "gpd-001",
			Code:           "CHARACTER",
			Name:           "Character",
			SequenceNumber: 1,
			IsActive:       true,
			CreatedBy:      "system",
			CreatedAt:      now,
			UpdatedAt:      now,
		}

		err := repo.CreateGraduateProfileDimension(ctx, gpd)
		require.NoError(t, err)

		gpd.Name = "Updated Character"
		gpd.UpdatedBy = StringPtr("admin")
		gpd.UpdatedAt = time.Now()

		err = repo.UpdateGraduateProfileDimension(ctx, gpd)
		require.NoError(t, err)

		updated, err := repo.GetGraduateProfileDimensionByID(ctx, gpd.ID)
		require.NoError(t, err)
		assert.Equal(t, "Updated Character", updated.Name)
	})

	t.Run("Error - Dimension not found", func(t *testing.T) {
		ctx := context.Background()

		gpd := &domain.GraduateProfileDimension{
			ID:   "non-existent-id",
			Name: "Non-existent",
		}

		err := repo.UpdateGraduateProfileDimension(ctx, gpd)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "graduate profile dimension not found")
	})
}

// TestGraduateProfileDimensionRepository_DeleteGraduateProfileDimension tests deleting a dimension
func TestGraduateProfileDimensionRepository_DeleteGraduateProfileDimension(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewGraduateProfileDimensionRepository(testDB.Pool)

	t.Run("Success - Delete dimension", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		now := time.Now()

		gpd := &domain.GraduateProfileDimension{
			ID:             "gpd-001",
			Code:           "DELETE",
			Name:           "Delete Me",
			SequenceNumber: 1,
			IsActive:       true,
			CreatedBy:      "system",
			CreatedAt:      now,
			UpdatedAt:      now,
		}

		err := repo.CreateGraduateProfileDimension(ctx, gpd)
		require.NoError(t, err)

		err = repo.DeleteGraduateProfileDimension(ctx, gpd.ID)
		require.NoError(t, err)

		_, err = repo.GetGraduateProfileDimensionByID(ctx, gpd.ID)
		assert.Error(t, err)
	})

	t.Run("Error - Dimension not found", func(t *testing.T) {
		ctx := context.Background()

		err := repo.DeleteGraduateProfileDimension(ctx, "non-existent-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "graduate profile dimension not found")
	})
}

// TestGraduateProfileDimensionRepository_CheckCodeExists tests checking code uniqueness
func TestGraduateProfileDimensionRepository_CheckCodeExists(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewGraduateProfileDimensionRepository(testDB.Pool)

	t.Run("Success - Code exists", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		now := time.Now()

		gpd := &domain.GraduateProfileDimension{
			ID:             "gpd-001",
			Code:           "UNIQUE",
			Name:           "Unique",
			SequenceNumber: 1,
			IsActive:       true,
			CreatedBy:      "system",
			CreatedAt:      now,
			UpdatedAt:      now,
		}

		repo.CreateGraduateProfileDimension(ctx, gpd)

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

// TestGraduateProfileDimensionRepository_CheckSequenceNumberExists tests checking sequence number uniqueness
func TestGraduateProfileDimensionRepository_CheckSequenceNumberExists(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewGraduateProfileDimensionRepository(testDB.Pool)

	t.Run("Success - Sequence number exists", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		now := time.Now()

		gpd := &domain.GraduateProfileDimension{
			ID:             "gpd-001",
			Code:           "TEST",
			Name:           "Test",
			SequenceNumber: 1,
			IsActive:       true,
			CreatedBy:      "system",
			CreatedAt:      now,
			UpdatedAt:      now,
		}

		repo.CreateGraduateProfileDimension(ctx, gpd)

		exists, err := repo.CheckSequenceNumberExists(ctx, 1, "")
		require.NoError(t, err)
		assert.True(t, exists)
	})

	t.Run("Success - Sequence number does not exist", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()

		exists, err := repo.CheckSequenceNumberExists(ctx, 99, "")
		require.NoError(t, err)
		assert.False(t, exists)
	})
}

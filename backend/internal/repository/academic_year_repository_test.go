package repository

import (
	"context"
	"testing"
	"time"

	"github.com/nusa/backend/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestAcademicYearRepository_CreateAcademicYear tests creating a new academic year
func TestAcademicYearRepository_CreateAcademicYear(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewAcademicYearRepository(testDB.Pool)

	t.Run("Success - Create academic year", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")

		ay := &domain.AcademicYear{
			ID:        "ay-001",
			SchoolID:  schoolID,
			Name:      "2024-2025",
			StartDate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
			Status:    domain.AcademicYearStatusDraft,
			CreatedBy: schoolID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		err := repo.CreateAcademicYear(ctx, ay)
		require.NoError(t, err)
		assert.NotEmpty(t, ay.ID)
	})
}

// TestAcademicYearRepository_GetAcademicYearByID tests retrieving an academic year by ID
func TestAcademicYearRepository_GetAcademicYearByID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewAcademicYearRepository(testDB.Pool)

	t.Run("Success - Academic year found", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")

		ay := &domain.AcademicYear{
			ID:        "ay-001",
			SchoolID:  schoolID,
			Name:      "2024-2025",
			StartDate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
			Status:    domain.AcademicYearStatusDraft,
			CreatedBy: schoolID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		err := repo.CreateAcademicYear(ctx, ay)
		require.NoError(t, err)

		retrieved, err := repo.GetAcademicYearByID(ctx, ay.ID)
		require.NoError(t, err)
		assert.Equal(t, ay.ID, retrieved.ID)
		assert.Equal(t, ay.Name, retrieved.Name)
	})

	t.Run("Error - Academic year not found", func(t *testing.T) {
		ctx := context.Background()

		_, err := repo.GetAcademicYearByID(ctx, "non-existent-id")
		assert.Error(t, err)
	})
}

// TestAcademicYearRepository_GetAcademicYearsBySchoolID tests retrieving academic years by school
func TestAcademicYearRepository_GetAcademicYearsBySchoolID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewAcademicYearRepository(testDB.Pool)

	t.Run("Success - Get academic years by school", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")

		ay1 := &domain.AcademicYear{
			ID:        "ay-001",
			SchoolID:  schoolID,
			Name:      "2023-2024",
			StartDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC),
			Status:    domain.AcademicYearStatusDraft,
			CreatedBy: schoolID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		ay2 := &domain.AcademicYear{
			ID:        "ay-002",
			SchoolID:  schoolID,
			Name:      "2024-2025",
			StartDate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
			Status:    domain.AcademicYearStatusDraft,
			CreatedBy: schoolID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		repo.CreateAcademicYear(ctx, ay1)
		repo.CreateAcademicYear(ctx, ay2)

		years, err := repo.GetAcademicYearsBySchoolID(ctx, schoolID)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(years), 2)
	})
}

// TestAcademicYearRepository_GetActiveAcademicYearBySchoolID tests retrieving active academic year
func TestAcademicYearRepository_GetActiveAcademicYearBySchoolID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewAcademicYearRepository(testDB.Pool)

	t.Run("Success - Get active academic year", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")

		ay := &domain.AcademicYear{
			ID:        "ay-001",
			SchoolID:  schoolID,
			Name:      "2024-2025",
			StartDate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
			Status:    domain.AcademicYearStatusActive,
			CreatedBy: schoolID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		err := repo.CreateAcademicYear(ctx, ay)
		require.NoError(t, err)

		activeAY, err := repo.GetActiveAcademicYearBySchoolID(ctx, schoolID)
		require.NoError(t, err)
		assert.Equal(t, ay.ID, activeAY.ID)
		assert.Equal(t, domain.AcademicYearStatusActive, activeAY.Status)
	})

	t.Run("Error - No active academic year", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")

		_, err := repo.GetActiveAcademicYearBySchoolID(ctx, schoolID)
		assert.Error(t, err)
	})
}

// TestAcademicYearRepository_UpdateAcademicYear tests updating an academic year
func TestAcademicYearRepository_UpdateAcademicYear(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewAcademicYearRepository(testDB.Pool)

	t.Run("Success - Update academic year", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")

		ay := &domain.AcademicYear{
			ID:        "ay-001",
			SchoolID:  schoolID,
			Name:      "2024-2025",
			StartDate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
			Status:    domain.AcademicYearStatusDraft,
			CreatedBy: schoolID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		err := repo.CreateAcademicYear(ctx, ay)
		require.NoError(t, err)

		ay.Name = "2024-2025 Updated"
		ay.Status = domain.AcademicYearStatusActive
		ay.UpdatedAt = time.Now()

		err = repo.UpdateAcademicYear(ctx, ay)
		require.NoError(t, err)

		updated, err := repo.GetAcademicYearByID(ctx, ay.ID)
		require.NoError(t, err)
		assert.Equal(t, "2024-2025 Updated", updated.Name)
		assert.Equal(t, domain.AcademicYearStatusActive, updated.Status)
	})

	t.Run("Error - Academic year not found", func(t *testing.T) {
		ctx := context.Background()

		ay := &domain.AcademicYear{
			ID:   "non-existent-id",
			Name: "Non-existent",
		}

		err := repo.UpdateAcademicYear(ctx, ay)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "academic year not found")
	})
}

// TestAcademicYearRepository_DeleteAcademicYear tests deleting an academic year
func TestAcademicYearRepository_DeleteAcademicYear(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewAcademicYearRepository(testDB.Pool)

	t.Run("Success - Delete academic year", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")

		ay := &domain.AcademicYear{
			ID:        "ay-001",
			SchoolID:  schoolID,
			Name:      "2024-2025",
			StartDate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
			Status:    domain.AcademicYearStatusDraft,
			CreatedBy: schoolID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		err := repo.CreateAcademicYear(ctx, ay)
		require.NoError(t, err)

		err = repo.DeleteAcademicYear(ctx, ay.ID)
		require.NoError(t, err)

		_, err = repo.GetAcademicYearByID(ctx, ay.ID)
		assert.Error(t, err)
	})

	t.Run("Error - Academic year not found", func(t *testing.T) {
		ctx := context.Background()

		err := repo.DeleteAcademicYear(ctx, "non-existent-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "academic year not found")
	})
}

// TestAcademicYearRepository_CheckAcademicYearOverlap tests overlap checking
func TestAcademicYearRepository_CheckAcademicYearOverlap(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewAcademicYearRepository(testDB.Pool)

	t.Run("Success - No overlap", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")

		ay := &domain.AcademicYear{
			ID:        "ay-001",
			SchoolID:  schoolID,
			Name:      "2024-2025",
			StartDate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
			Status:    domain.AcademicYearStatusDraft,
			CreatedBy: schoolID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		repo.CreateAcademicYear(ctx, ay)

		// Check for overlap with non-overlapping dates
		startDate := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
		endDate := time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC)

		overlaps, err := repo.CheckAcademicYearOverlap(ctx, schoolID, startDate, endDate, "")
		require.NoError(t, err)
		assert.False(t, overlaps)
	})

	t.Run("Success - Overlap detected", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")

		ay := &domain.AcademicYear{
			ID:        "ay-001",
			SchoolID:  schoolID,
			Name:      "2024-2025",
			StartDate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
			Status:    domain.AcademicYearStatusActive,
			CreatedBy: schoolID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		repo.CreateAcademicYear(ctx, ay)

		// Check for overlap with overlapping dates
		startDate := time.Date(2024, 6, 1, 0, 0, 0, 0, time.UTC)
		endDate := time.Date(2025, 6, 1, 0, 0, 0, 0, time.UTC)

		overlaps, err := repo.CheckAcademicYearOverlap(ctx, schoolID, startDate, endDate, "")
		require.NoError(t, err)
		assert.True(t, overlaps)
	})
}

// TestAcademicYearRepository_ActivateAcademicYear tests activating an academic year
func TestAcademicYearRepository_ActivateAcademicYear(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewAcademicYearRepository(testDB.Pool)

	t.Run("Success - Activate academic year", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")

		ay1 := &domain.AcademicYear{
			ID:        "ay-001",
			SchoolID:  schoolID,
			Name:      "2023-2024",
			StartDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC),
			Status:    domain.AcademicYearStatusActive,
			CreatedBy: schoolID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		ay2 := &domain.AcademicYear{
			ID:        "ay-002",
			SchoolID:  schoolID,
			Name:      "2024-2025",
			StartDate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
			Status:    domain.AcademicYearStatusDraft,
			CreatedBy: schoolID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		repo.CreateAcademicYear(ctx, ay1)
		repo.CreateAcademicYear(ctx, ay2)

		// Activate ay2, should deactivate ay1
		err := repo.ActivateAcademicYear(ctx, ay2.ID)
		require.NoError(t, err)

		activeAY, err := repo.GetActiveAcademicYearBySchoolID(ctx, schoolID)
		require.NoError(t, err)
		assert.Equal(t, ay2.ID, activeAY.ID)
	})
}

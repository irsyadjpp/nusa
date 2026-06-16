package repository

import (
	"context"
	"testing"
	"time"

	"github.com/nusa/backend/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestSemesterRepository_CreateSemester tests creating a new semester
func TestSemesterRepository_CreateSemester(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewSemesterRepository(testDB.Pool)

	t.Run("Success - Create semester", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")
		ayID := CreateTestAcademicYear(t, testDB.Pool, schoolID, "2024-2025")

		sem := &domain.Semester{
			ID:             "sem-001",
			AcademicYearID: ayID,
			Type:           domain.SemesterTypeGanjil,
			Name:           "Semester 1",
			StartDate:      time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			EndDate:        time.Date(2024, 6, 30, 0, 0, 0, 0, time.UTC),
			Status:         domain.SemesterStatusActive,
			SequenceNumber: 1,
			CreatedBy:      schoolID,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		err := repo.CreateSemester(ctx, sem)
		require.NoError(t, err)
		assert.NotEmpty(t, sem.ID)
	})
}

// TestSemesterRepository_GetSemesterByID tests retrieving a semester by ID
func TestSemesterRepository_GetSemesterByID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewSemesterRepository(testDB.Pool)

	t.Run("Success - Semester found", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")
		ayID := CreateTestAcademicYear(t, testDB.Pool, schoolID, "2024-2025")

		sem := &domain.Semester{
			ID:             "sem-001",
			AcademicYearID: ayID,
			Type:           domain.SemesterTypeGanjil,
			Name:           "Semester 1",
			StartDate:      time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			EndDate:        time.Date(2024, 6, 30, 0, 0, 0, 0, time.UTC),
			Status:         domain.SemesterStatusActive,
			SequenceNumber: 1,
			CreatedBy:      schoolID,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		err := repo.CreateSemester(ctx, sem)
		require.NoError(t, err)

		retrieved, err := repo.GetSemesterByID(ctx, sem.ID)
		require.NoError(t, err)
		assert.Equal(t, sem.ID, retrieved.ID)
		assert.Equal(t, sem.Name, retrieved.Name)
	})

	t.Run("Error - Semester not found", func(t *testing.T) {
		ctx := context.Background()

		_, err := repo.GetSemesterByID(ctx, "non-existent-id")
		assert.Error(t, err)
	})
}

// TestSemesterRepository_GetSemestersByAcademicYearID tests retrieving semesters by academic year
func TestSemesterRepository_GetSemestersByAcademicYearID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewSemesterRepository(testDB.Pool)

	t.Run("Success - Get semesters by academic year", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")
		ayID := CreateTestAcademicYear(t, testDB.Pool, schoolID, "2024-2025")

		sem1 := &domain.Semester{
			ID:             "sem-001",
			AcademicYearID: ayID,
			Type:           domain.SemesterTypeGanjil,
			Name:           "Semester 1",
			StartDate:      time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			EndDate:        time.Date(2024, 6, 30, 0, 0, 0, 0, time.UTC),
			Status:         domain.SemesterStatusActive,
			SequenceNumber: 1,
			CreatedBy:      schoolID,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		sem2 := &domain.Semester{
			ID:             "sem-002",
			AcademicYearID: ayID,
			Type:           domain.SemesterTypeGanjil,
			Name:           "Semester 2",
			StartDate:      time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC),
			EndDate:        time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
			Status:         domain.SemesterStatusActive,
			SequenceNumber: 2,
			CreatedBy:      schoolID,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		repo.CreateSemester(ctx, sem1)
		repo.CreateSemester(ctx, sem2)

		semesters, err := repo.GetSemestersByAcademicYearID(ctx, ayID)
		require.NoError(t, err)
		assert.Equal(t, 2, len(semesters))
	})
}

// TestSemesterRepository_UpdateSemester tests updating a semester
func TestSemesterRepository_UpdateSemester(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewSemesterRepository(testDB.Pool)

	t.Run("Success - Update semester", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")
		ayID := CreateTestAcademicYear(t, testDB.Pool, schoolID, "2024-2025")

		sem := &domain.Semester{
			ID:             "sem-001",
			AcademicYearID: ayID,
			Type:           domain.SemesterTypeGanjil,
			Name:           "Semester 1",
			StartDate:      time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			EndDate:        time.Date(2024, 6, 30, 0, 0, 0, 0, time.UTC),
			Status:         domain.SemesterStatusActive,
			SequenceNumber: 1,
			CreatedBy:      schoolID,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		err := repo.CreateSemester(ctx, sem)
		require.NoError(t, err)

		sem.Name = "Semester 1 Updated"
		sem.Status = domain.SemesterStatusActive
		sem.UpdatedAt = time.Now()

		err = repo.UpdateSemester(ctx, sem)
		require.NoError(t, err)

		updated, err := repo.GetSemesterByID(ctx, sem.ID)
		require.NoError(t, err)
		assert.Equal(t, "Semester 1 Updated", updated.Name)
		assert.Equal(t, domain.SemesterStatusActive, updated.Status)
	})

	t.Run("Error - Semester not found", func(t *testing.T) {
		ctx := context.Background()

		sem := &domain.Semester{
			ID:   "non-existent-id",
			Name: "Non-existent",
		}

		err := repo.UpdateSemester(ctx, sem)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "semester not found")
	})
}

// TestSemesterRepository_DeleteSemester tests deleting a semester
func TestSemesterRepository_DeleteSemester(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewSemesterRepository(testDB.Pool)

	t.Run("Success - Delete semester", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")
		ayID := CreateTestAcademicYear(t, testDB.Pool, schoolID, "2024-2025")

		sem := &domain.Semester{
			ID:             "sem-001",
			AcademicYearID: ayID,
			Type:           domain.SemesterTypeGanjil,
			Name:           "Semester 1",
			StartDate:      time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			EndDate:        time.Date(2024, 6, 30, 0, 0, 0, 0, time.UTC),
			Status:         domain.SemesterStatusActive,
			SequenceNumber: 1,
			CreatedBy:      schoolID,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		err := repo.CreateSemester(ctx, sem)
		require.NoError(t, err)

		err = repo.DeleteSemester(ctx, sem.ID)
		require.NoError(t, err)

		_, err = repo.GetSemesterByID(ctx, sem.ID)
		assert.Error(t, err)
	})

	t.Run("Error - Semester not found", func(t *testing.T) {
		ctx := context.Background()

		err := repo.DeleteSemester(ctx, "non-existent-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "semester not found")
	})
}

// TestSemesterRepository_CheckSemesterOverlap tests overlap checking
func TestSemesterRepository_CheckSemesterOverlap(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewSemesterRepository(testDB.Pool)

	t.Run("Success - No overlap", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")
		ayID := CreateTestAcademicYear(t, testDB.Pool, schoolID, "2024-2025")

		sem := &domain.Semester{
			ID:             "sem-001",
			AcademicYearID: ayID,
			Type:           domain.SemesterTypeGanjil,
			Name:           "Semester 1",
			StartDate:      time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			EndDate:        time.Date(2024, 6, 30, 0, 0, 0, 0, time.UTC),
			Status:         domain.SemesterStatusActive,
			SequenceNumber: 1,
			CreatedBy:      schoolID,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		repo.CreateSemester(ctx, sem)

		startDate := time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC)
		endDate := time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)

		overlaps, err := repo.CheckSemesterOverlap(ctx, ayID, startDate, endDate, "")
		require.NoError(t, err)
		assert.False(t, overlaps)
	})

	t.Run("Success - Overlap detected", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")
		ayID := CreateTestAcademicYear(t, testDB.Pool, schoolID, "2024-2025")

		sem := &domain.Semester{
			ID:             "sem-001",
			AcademicYearID: ayID,
			Type:           domain.SemesterTypeGanjil,
			Name:           "Semester 1",
			StartDate:      time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			EndDate:        time.Date(2024, 6, 30, 0, 0, 0, 0, time.UTC),
			Status:         domain.SemesterStatusActive,
			SequenceNumber: 1,
			CreatedBy:      schoolID,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		repo.CreateSemester(ctx, sem)

		startDate := time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC)
		endDate := time.Date(2024, 9, 30, 0, 0, 0, 0, time.UTC)

		overlaps, err := repo.CheckSemesterOverlap(ctx, ayID, startDate, endDate, "")
		require.NoError(t, err)
		assert.True(t, overlaps)
	})
}

// TestSemesterRepository_CountSemestersByAcademicYearID tests counting semesters
func TestSemesterRepository_CountSemestersByAcademicYearID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewSemesterRepository(testDB.Pool)

	t.Run("Success - Count semesters", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")
		ayID := CreateTestAcademicYear(t, testDB.Pool, schoolID, "2024-2025")

		sem := &domain.Semester{
			ID:             "sem-001",
			AcademicYearID: ayID,
			Type:           domain.SemesterTypeGanjil,
			Name:           "Semester 1",
			StartDate:      time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			EndDate:        time.Date(2024, 6, 30, 0, 0, 0, 0, time.UTC),
			Status:         domain.SemesterStatusActive,
			SequenceNumber: 1,
			CreatedBy:      schoolID,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		repo.CreateSemester(ctx, sem)

		count, err := repo.CountSemestersByAcademicYearID(ctx, ayID)
		require.NoError(t, err)
		assert.Equal(t, 1, count)
	})
}

// TestSemesterRepository_CheckSequenceNumberExists tests checking sequence number uniqueness
func TestSemesterRepository_CheckSequenceNumberExists(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewSemesterRepository(testDB.Pool)

	t.Run("Success - Sequence number exists", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")
		ayID := CreateTestAcademicYear(t, testDB.Pool, schoolID, "2024-2025")

		sem := &domain.Semester{
			ID:             "sem-001",
			AcademicYearID: ayID,
			Type:           domain.SemesterTypeGanjil,
			Name:           "Semester 1",
			StartDate:      time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			EndDate:        time.Date(2024, 6, 30, 0, 0, 0, 0, time.UTC),
			Status:         domain.SemesterStatusActive,
			SequenceNumber: 1,
			CreatedBy:      schoolID,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		repo.CreateSemester(ctx, sem)

		exists, err := repo.CheckSequenceNumberExists(ctx, ayID, 1, "")
		require.NoError(t, err)
		assert.True(t, exists)
	})

	t.Run("Success - Sequence number does not exist", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		schoolID := CreateTestSchool(t, testDB.Pool, "Test School")
		ayID := CreateTestAcademicYear(t, testDB.Pool, schoolID, "2024-2025")

		exists, err := repo.CheckSequenceNumberExists(ctx, ayID, 2, "")
		require.NoError(t, err)
		assert.False(t, exists)
	})
}

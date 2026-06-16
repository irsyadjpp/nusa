package repository

import (
	"context"
	"testing"
	"time"

	"github.com/nusa/backend/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestCPAlignmentRepository_CreateCPAlignment tests creating a CP alignment
func TestCPAlignmentRepository_CreateCPAlignment(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewCPAlignmentRepository(testDB.Pool)

	t.Run("Success - Create CP alignment", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		now := time.Now()

		// Create prerequisite data
		subject := &domain.CurriculumSubject{
			ID:       "subj-001",
			Code:     "MATH",
			Name:     "Mathematics",
			IsActive: true,
		}
		curriculumRepo := NewCurriculumRepository(testDB.Pool)
		err := curriculumRepo.CreateCurriculumSubject(ctx, subject)
		require.NoError(t, err)

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
		gpdRepo := NewGraduateProfileDimensionRepository(testDB.Pool)
		err = gpdRepo.CreateGraduateProfileDimension(ctx, gpd)
		require.NoError(t, err)

		description := "Math develops character"
		cpa := &domain.CPAlignment{
			ID:                         "cpa-001",
			CurriculumSubjectID:         subject.ID,
			GraduateProfileDimensionID:  gpd.ID,
			AlignmentDescription:       &description,
			IsActive:                   true,
			CreatedBy:                  "system",
			CreatedAt:                  now,
			UpdatedAt:                  now,
		}

		err = repo.CreateCPAlignment(ctx, cpa)
		require.NoError(t, err)
	})
}

// TestCPAlignmentRepository_GetCPAlignmentByID tests retrieving by ID
func TestCPAlignmentRepository_GetCPAlignmentByID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewCPAlignmentRepository(testDB.Pool)

	t.Run("Success - Alignment found", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		now := time.Now()

		subject := &domain.CurriculumSubject{
			ID:       "subj-001",
			Code:     "MATH",
			Name:     "Mathematics",
			IsActive: true,
		}
		curriculumRepo := NewCurriculumRepository(testDB.Pool)
		curriculumRepo.CreateCurriculumSubject(ctx, subject)

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
		gpdRepo := NewGraduateProfileDimensionRepository(testDB.Pool)
		gpdRepo.CreateGraduateProfileDimension(ctx, gpd)

		description := "Math develops character"
		cpa := &domain.CPAlignment{
			ID:                         "cpa-001",
			CurriculumSubjectID:         subject.ID,
			GraduateProfileDimensionID:  gpd.ID,
			AlignmentDescription:       &description,
			IsActive:                   true,
			CreatedBy:                  "system",
			CreatedAt:                  now,
			UpdatedAt:                  now,
		}

		err := repo.CreateCPAlignment(ctx, cpa)
		require.NoError(t, err)

		retrieved, err := repo.GetCPAlignmentByID(ctx, cpa.ID)
		require.NoError(t, err)
		assert.Equal(t, cpa.ID, retrieved.ID)
	})

	t.Run("Error - Alignment not found", func(t *testing.T) {
		ctx := context.Background()

		_, err := repo.GetCPAlignmentByID(ctx, "non-existent-id")
		assert.Error(t, err)
	})
}

// TestCPAlignmentRepository_GetCPAlignmentsByCurriculumSubjectID tests retrieving by subject
func TestCPAlignmentRepository_GetCPAlignmentsByCurriculumSubjectID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewCPAlignmentRepository(testDB.Pool)

	t.Run("Success - Get alignments by subject", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		now := time.Now()

		subject := &domain.CurriculumSubject{
			ID:       "subj-001",
			Code:     "MATH",
			Name:     "Mathematics",
			IsActive: true,
		}
		curriculumRepo := NewCurriculumRepository(testDB.Pool)
		curriculumRepo.CreateCurriculumSubject(ctx, subject)

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
		gpdRepo := NewGraduateProfileDimensionRepository(testDB.Pool)
		gpdRepo.CreateGraduateProfileDimension(ctx, gpd1)
		gpdRepo.CreateGraduateProfileDimension(ctx, gpd2)

		cpa1 := &domain.CPAlignment{
			ID:                         "cpa-001",
			CurriculumSubjectID:         subject.ID,
			GraduateProfileDimensionID:  gpd1.ID,
			IsActive:                   true,
			CreatedBy:                  "system",
			CreatedAt:                  now,
			UpdatedAt:                  now,
		}

		cpa2 := &domain.CPAlignment{
			ID:                         "cpa-002",
			CurriculumSubjectID:         subject.ID,
			GraduateProfileDimensionID:  gpd2.ID,
			IsActive:                   true,
			CreatedBy:                  "system",
			CreatedAt:                  now,
			UpdatedAt:                  now,
		}

		repo.CreateCPAlignment(ctx, cpa1)
		repo.CreateCPAlignment(ctx, cpa2)

		alignments, err := repo.GetCPAlignmentsByCurriculumSubjectID(ctx, subject.ID)
		require.NoError(t, err)
		assert.Equal(t, 2, len(alignments))
	})
}

// TestCPAlignmentRepository_UpdateCPAlignment tests updating an alignment
func TestCPAlignmentRepository_UpdateCPAlignment(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewCPAlignmentRepository(testDB.Pool)

	t.Run("Success - Update alignment", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		now := time.Now()

		subject := &domain.CurriculumSubject{
			ID:       "subj-001",
			Code:     "MATH",
			Name:     "Mathematics",
			IsActive: true,
		}
		curriculumRepo := NewCurriculumRepository(testDB.Pool)
		curriculumRepo.CreateCurriculumSubject(ctx, subject)

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
		gpdRepo := NewGraduateProfileDimensionRepository(testDB.Pool)
		gpdRepo.CreateGraduateProfileDimension(ctx, gpd)

		cpa := &domain.CPAlignment{
			ID:                         "cpa-001",
			CurriculumSubjectID:         subject.ID,
			GraduateProfileDimensionID:  gpd.ID,
			IsActive:                   true,
			CreatedBy:                  "system",
			CreatedAt:                  now,
			UpdatedAt:                  now,
		}

		err := repo.CreateCPAlignment(ctx, cpa)
		require.NoError(t, err)

		newDescription := "Updated description"
		cpa.AlignmentDescription = &newDescription
		cpa.UpdatedBy = StringPtr("admin")
		cpa.UpdatedAt = time.Now()

		err = repo.UpdateCPAlignment(ctx, cpa)
		require.NoError(t, err)

		updated, err := repo.GetCPAlignmentByID(ctx, cpa.ID)
		require.NoError(t, err)
		assert.NotNil(t, updated.AlignmentDescription)
		assert.Equal(t, newDescription, *updated.AlignmentDescription)
	})

	t.Run("Error - Alignment not found", func(t *testing.T) {
		ctx := context.Background()

		cpa := &domain.CPAlignment{
			ID: "non-existent-id",
		}

		err := repo.UpdateCPAlignment(ctx, cpa)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "CP alignment not found")
	})
}

// TestCPAlignmentRepository_DeleteCPAlignment tests deleting an alignment
func TestCPAlignmentRepository_DeleteCPAlignment(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewCPAlignmentRepository(testDB.Pool)

	t.Run("Success - Delete alignment", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		now := time.Now()

		subject := &domain.CurriculumSubject{
			ID:       "subj-001",
			Code:     "MATH",
			Name:     "Mathematics",
			IsActive: true,
		}
		curriculumRepo := NewCurriculumRepository(testDB.Pool)
		curriculumRepo.CreateCurriculumSubject(ctx, subject)

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
		gpdRepo := NewGraduateProfileDimensionRepository(testDB.Pool)
		gpdRepo.CreateGraduateProfileDimension(ctx, gpd)

		cpa := &domain.CPAlignment{
			ID:                         "cpa-001",
			CurriculumSubjectID:         subject.ID,
			GraduateProfileDimensionID:  gpd.ID,
			IsActive:                   true,
			CreatedBy:                  "system",
			CreatedAt:                  now,
			UpdatedAt:                  now,
		}

		err := repo.CreateCPAlignment(ctx, cpa)
		require.NoError(t, err)

		err = repo.DeleteCPAlignment(ctx, cpa.ID)
		require.NoError(t, err)

		_, err = repo.GetCPAlignmentByID(ctx, cpa.ID)
		assert.Error(t, err)
	})

	t.Run("Error - Alignment not found", func(t *testing.T) {
		ctx := context.Background()

		err := repo.DeleteCPAlignment(ctx, "non-existent-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "CP alignment not found")
	})
}

// TestCPAlignmentRepository_CheckAlignmentExists tests checking alignment uniqueness
func TestCPAlignmentRepository_CheckAlignmentExists(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewCPAlignmentRepository(testDB.Pool)

	t.Run("Success - Alignment exists", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		now := time.Now()

		subject := &domain.CurriculumSubject{
			ID:       "subj-001",
			Code:     "MATH",
			Name:     "Mathematics",
			IsActive: true,
		}
		curriculumRepo := NewCurriculumRepository(testDB.Pool)
		curriculumRepo.CreateCurriculumSubject(ctx, subject)

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
		gpdRepo := NewGraduateProfileDimensionRepository(testDB.Pool)
		gpdRepo.CreateGraduateProfileDimension(ctx, gpd)

		cpa := &domain.CPAlignment{
			ID:                         "cpa-001",
			CurriculumSubjectID:         subject.ID,
			GraduateProfileDimensionID:  gpd.ID,
			IsActive:                   true,
			CreatedBy:                  "system",
			CreatedAt:                  now,
			UpdatedAt:                  now,
		}

		repo.CreateCPAlignment(ctx, cpa)

		exists, err := repo.CheckAlignmentExists(ctx, subject.ID, gpd.ID, "")
		require.NoError(t, err)
		assert.True(t, exists)
	})

	t.Run("Success - Alignment does not exist", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()

		exists, err := repo.CheckAlignmentExists(ctx, "non-existent-subject", "non-existent-gpd", "")
		require.NoError(t, err)
		assert.False(t, exists)
	})
}

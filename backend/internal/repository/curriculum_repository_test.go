package repository

import (
	"context"
	"testing"

	"github.com/nusa/backend/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestCurriculumRepository_CreateCurriculumSubject tests creating a curriculum subject
func TestCurriculumRepository_CreateCurriculumSubject(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewCurriculumRepository(testDB.Pool)

	t.Run("Success - Create subject", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		description := "Mathematics subject"
		subject := &domain.CurriculumSubject{
			ID:          "subj-001",
			Code:        "MATH",
			Name:        "Mathematics",
			Description: &description,
			IsActive:    true,
		}

		err := repo.CreateCurriculumSubject(ctx, subject)
		require.NoError(t, err)
	})

	t.Run("Success - Create subject without description", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		subject := &domain.CurriculumSubject{
			ID:       "subj-002",
			Code:     "SCI",
			Name:     "Science",
			IsActive: true,
		}

		err := repo.CreateCurriculumSubject(ctx, subject)
		require.NoError(t, err)
	})
}

// TestCurriculumRepository_GetCurriculumSubjectByID tests retrieving a subject by ID
func TestCurriculumRepository_GetCurriculumSubjectByID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewCurriculumRepository(testDB.Pool)

	t.Run("Success - Subject found", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		description := "Mathematics subject"
		subject := &domain.CurriculumSubject{
			ID:          "subj-001",
			Code:        "MATH",
			Name:        "Mathematics",
			Description: &description,
			IsActive:    true,
		}

		err := repo.CreateCurriculumSubject(ctx, subject)
		require.NoError(t, err)

		retrieved, err := repo.GetCurriculumSubjectByID(ctx, subject.ID)
		require.NoError(t, err)
		assert.Equal(t, subject.ID, retrieved.ID)
		assert.Equal(t, subject.Name, retrieved.Name)
	})

	t.Run("Error - Subject not found", func(t *testing.T) {
		ctx := context.Background()

		_, err := repo.GetCurriculumSubjectByID(ctx, "non-existent-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "curriculum subject not found")
	})
}

// TestCurriculumRepository_GetCurriculumSubjectByCode tests retrieving by code
func TestCurriculumRepository_GetCurriculumSubjectByCode(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewCurriculumRepository(testDB.Pool)

	t.Run("Success - Subject found by code", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		subject := &domain.CurriculumSubject{
			ID:       "subj-001",
			Code:     "MATH",
			Name:     "Mathematics",
			IsActive: true,
		}

		err := repo.CreateCurriculumSubject(ctx, subject)
		require.NoError(t, err)

		retrieved, err := repo.GetCurriculumSubjectByCode(ctx, "MATH")
		require.NoError(t, err)
		assert.Equal(t, subject.ID, retrieved.ID)
	})

	t.Run("Error - Subject not found by code", func(t *testing.T) {
		ctx := context.Background()

		_, err := repo.GetCurriculumSubjectByCode(ctx, "NONEXISTENT")
		assert.Error(t, err)
	})
}

// TestCurriculumRepository_ListCurriculumSubjects tests listing subjects
func TestCurriculumRepository_ListCurriculumSubjects(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewCurriculumRepository(testDB.Pool)

	t.Run("Success - List all subjects", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()

		subject1 := &domain.CurriculumSubject{ID: "subj-001", Code: "MATH", Name: "Mathematics", IsActive: true}
		subject2 := &domain.CurriculumSubject{ID: "subj-002", Code: "SCI", Name: "Science", IsActive: true}

		repo.CreateCurriculumSubject(ctx, subject1)
		repo.CreateCurriculumSubject(ctx, subject2)

		subjects, err := repo.ListCurriculumSubjects(ctx, nil, 0, 0)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(subjects), 2)
	})

	t.Run("Success - Filter active subjects", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()

		subject1 := &domain.CurriculumSubject{ID: "subj-001", Code: "MATH", Name: "Mathematics", IsActive: true}
		subject2 := &domain.CurriculumSubject{ID: "subj-002", Code: "SCI", Name: "Science", IsActive: false}

		repo.CreateCurriculumSubject(ctx, subject1)
		repo.CreateCurriculumSubject(ctx, subject2)

		isActive := true
		subjects, err := repo.ListCurriculumSubjects(ctx, &isActive, 0, 0)
		require.NoError(t, err)

		codes := make([]string, len(subjects))
		for i, s := range subjects {
			codes[i] = s.Code
		}
		assert.Contains(t, codes, "MATH")
	})
}

// TestCurriculumRepository_UpdateCurriculumSubject tests updating a subject
func TestCurriculumRepository_UpdateCurriculumSubject(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewCurriculumRepository(testDB.Pool)

	t.Run("Success - Update subject", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		subject := &domain.CurriculumSubject{
			ID:       "subj-001",
			Code:     "MATH",
			Name:     "Mathematics",
			IsActive: true,
		}

		err := repo.CreateCurriculumSubject(ctx, subject)
		require.NoError(t, err)

		subject.Name = "Advanced Mathematics"
		err = repo.UpdateCurriculumSubject(ctx, subject)
		require.NoError(t, err)

		updated, err := repo.GetCurriculumSubjectByID(ctx, subject.ID)
		require.NoError(t, err)
		assert.Equal(t, "Advanced Mathematics", updated.Name)
	})

	t.Run("Error - Subject not found", func(t *testing.T) {
		ctx := context.Background()

		subject := &domain.CurriculumSubject{
			ID:   "non-existent-id",
			Name: "Non-existent",
		}

		err := repo.UpdateCurriculumSubject(ctx, subject)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "curriculum subject not found")
	})
}

// TestCurriculumRepository_DeleteCurriculumSubject tests deleting a subject
func TestCurriculumRepository_DeleteCurriculumSubject(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewCurriculumRepository(testDB.Pool)

	t.Run("Success - Delete subject", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		subject := &domain.CurriculumSubject{
			ID:       "subj-001",
			Code:     "MATH",
			Name:     "Mathematics",
			IsActive: true,
		}

		err := repo.CreateCurriculumSubject(ctx, subject)
		require.NoError(t, err)

		err = repo.DeleteCurriculumSubject(ctx, subject.ID)
		require.NoError(t, err)

		_, err = repo.GetCurriculumSubjectByID(ctx, subject.ID)
		assert.Error(t, err)
	})

	t.Run("Error - Subject not found", func(t *testing.T) {
		ctx := context.Background()

		err := repo.DeleteCurriculumSubject(ctx, "non-existent-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "curriculum subject not found")
	})
}

// TestCurriculumRepository_CreateCurriculumPhase tests creating a curriculum phase
func TestCurriculumRepository_CreateCurriculumPhase(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewCurriculumRepository(testDB.Pool)

	t.Run("Success - Create phase", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		description := "Primary education phase"
		gradeStart := 1
		gradeEnd := 6
		phase := &domain.CurriculumPhase{
			ID:              "phase-001",
			Code:            "PRIMARY",
			Name:            "Primary",
			Description:     &description,
			GradeLevelStart: &gradeStart,
			GradeLevelEnd:   &gradeEnd,
			IsActive:        true,
		}

		err := repo.CreateCurriculumPhase(ctx, phase)
		require.NoError(t, err)
	})
}

// TestCurriculumRepository_GetCurriculumPhaseByID tests retrieving a phase
func TestCurriculumRepository_GetCurriculumPhaseByID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewCurriculumRepository(testDB.Pool)

	t.Run("Success - Phase found", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		gradeStart := 1
		gradeEnd := 6
		phase := &domain.CurriculumPhase{
			ID:              "phase-001",
			Code:            "PRIMARY",
			Name:            "Primary",
			GradeLevelStart: &gradeStart,
			GradeLevelEnd:   &gradeEnd,
			IsActive:        true,
		}

		err := repo.CreateCurriculumPhase(ctx, phase)
		require.NoError(t, err)

		retrieved, err := repo.GetCurriculumPhaseByID(ctx, phase.ID)
		require.NoError(t, err)
		assert.Equal(t, phase.ID, retrieved.ID)
		assert.NotNil(t, retrieved.GradeLevelStart)
		assert.NotNil(t, retrieved.GradeLevelEnd)
	})

	t.Run("Error - Phase not found", func(t *testing.T) {
		ctx := context.Background()

		_, err := repo.GetCurriculumPhaseByID(ctx, "non-existent-id")
		assert.Error(t, err)
	})
}

// TestCurriculumRepository_ListCurriculumPhases tests listing phases
func TestCurriculumRepository_ListCurriculumPhases(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewCurriculumRepository(testDB.Pool)

	t.Run("Success - List all phases", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()

		gradeStart1 := 1
		gradeEnd1 := 6
		phase1 := &domain.CurriculumPhase{
			ID:              "phase-001",
			Code:            "PRIMARY",
			Name:            "Primary",
			GradeLevelStart: &gradeStart1,
			GradeLevelEnd:   &gradeEnd1,
			IsActive:        true,
		}

		gradeStart2 := 7
		gradeEnd2 := 9
		phase2 := &domain.CurriculumPhase{
			ID:              "phase-002",
			Code:            "SECONDARY",
			Name:            "Secondary",
			GradeLevelStart: &gradeStart2,
			GradeLevelEnd:   &gradeEnd2,
			IsActive:        true,
		}

		repo.CreateCurriculumPhase(ctx, phase1)
		repo.CreateCurriculumPhase(ctx, phase2)

		phases, err := repo.ListCurriculumPhases(ctx, nil, 0, 0)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(phases), 2)
	})
}

package repository

import (
	"context"
	"testing"
	"time"

	"github.com/nusa/backend/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestSystemConfigurationRepository_CreateSystemConfiguration tests creating a system configuration
func TestSystemConfigurationRepository_CreateSystemConfiguration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewSystemConfigurationRepository(testDB.Pool)

	t.Run("Success - Create system configuration", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		description := "Maximum file upload size"
		now := time.Now()

		sc := &domain.SystemConfiguration{
			ID:          "sc-001",
			Key:         "max_file_size",
			Value:       "10485760",
			ValueType:   "integer",
			Description: &description,
			Category:    "upload",
			IsSystem:    true,
			IsActive:    true,
			CreatedBy:   "system",
			CreatedAt:   now,
			UpdatedAt:   now,
		}

		err := repo.CreateSystemConfiguration(ctx, sc)
		require.NoError(t, err)
	})
}

// TestSystemConfigurationRepository_GetSystemConfigurationByID tests retrieving by ID
func TestSystemConfigurationRepository_GetSystemConfigurationByID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewSystemConfigurationRepository(testDB.Pool)

	t.Run("Success - Configuration found", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		description := "Maximum file upload size"
		now := time.Now()

		sc := &domain.SystemConfiguration{
			ID:          "sc-001",
			Key:         "max_file_size",
			Value:       "10485760",
			ValueType:   "integer",
			Description: &description,
			Category:    "upload",
			IsSystem:    true,
			IsActive:    true,
			CreatedBy:   "system",
			CreatedAt:   now,
			UpdatedAt:   now,
		}

		err := repo.CreateSystemConfiguration(ctx, sc)
		require.NoError(t, err)

		retrieved, err := repo.GetSystemConfigurationByID(ctx, sc.ID)
		require.NoError(t, err)
		assert.Equal(t, sc.ID, retrieved.ID)
		assert.Equal(t, sc.Key, retrieved.Key)
	})

	t.Run("Error - Configuration not found", func(t *testing.T) {
		ctx := context.Background()

		_, err := repo.GetSystemConfigurationByID(ctx, "non-existent-id")
		assert.Error(t, err)
	})
}

// TestSystemConfigurationRepository_GetSystemConfigurationByKey tests retrieving by key
func TestSystemConfigurationRepository_GetSystemConfigurationByKey(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewSystemConfigurationRepository(testDB.Pool)

	t.Run("Success - Configuration found by key", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		now := time.Now()

		sc := &domain.SystemConfiguration{
			ID:        "sc-001",
			Key:       "max_file_size",
			Value:     "10485760",
			ValueType: "integer",
			Category:  "upload",
			IsSystem:  true,
			IsActive:  true,
			CreatedBy: "system",
			CreatedAt: now,
			UpdatedAt: now,
		}

		err := repo.CreateSystemConfiguration(ctx, sc)
		require.NoError(t, err)

		retrieved, err := repo.GetSystemConfigurationByKey(ctx, "max_file_size")
		require.NoError(t, err)
		assert.Equal(t, sc.ID, retrieved.ID)
		assert.Equal(t, sc.Key, retrieved.Key)
	})

	t.Run("Error - Configuration not found by key", func(t *testing.T) {
		ctx := context.Background()

		_, err := repo.GetSystemConfigurationByKey(ctx, "non-existent-key")
		assert.Error(t, err)
	})
}

// TestSystemConfigurationRepository_GetAllSystemConfigurations tests retrieving all configurations
func TestSystemConfigurationRepository_GetAllSystemConfigurations(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewSystemConfigurationRepository(testDB.Pool)

	t.Run("Success - Get all configurations", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		now := time.Now()

		sc1 := &domain.SystemConfiguration{
			ID:        "sc-001",
			Key:       "max_file_size",
			Value:     "10485760",
			ValueType: "integer",
			Category:  "upload",
			IsSystem:  true,
			IsActive:  true,
			CreatedBy: "system",
			CreatedAt: now,
			UpdatedAt: now,
		}

		sc2 := &domain.SystemConfiguration{
			ID:        "sc-002",
			Key:       "session_timeout",
			Value:     "3600",
			ValueType: "integer",
			Category:  "security",
			IsSystem:  true,
			IsActive:  true,
			CreatedBy: "system",
			CreatedAt: now,
			UpdatedAt: now,
		}

		repo.CreateSystemConfiguration(ctx, sc1)
		repo.CreateSystemConfiguration(ctx, sc2)

		configurations, err := repo.GetAllSystemConfigurations(ctx)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(configurations), 2)
	})
}

// TestSystemConfigurationRepository_GetSystemConfigurationsByCategory tests retrieving by category
func TestSystemConfigurationRepository_GetSystemConfigurationsByCategory(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewSystemConfigurationRepository(testDB.Pool)

	t.Run("Success - Get configurations by category", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		now := time.Now()

		sc1 := &domain.SystemConfiguration{
			ID:        "sc-001",
			Key:       "max_file_size",
			Value:     "10485760",
			ValueType: "integer",
			Category:  "upload",
			IsSystem:  true,
			IsActive:  true,
			CreatedBy: "system",
			CreatedAt: now,
			UpdatedAt: now,
		}

		sc2 := &domain.SystemConfiguration{
			ID:        "sc-002",
			Key:       "allowed_extensions",
			Value:     "pdf,doc,docx",
			ValueType: "string",
			Category:  "upload",
			IsSystem:  false,
			IsActive:  true,
			CreatedBy: "system",
			CreatedAt: now,
			UpdatedAt: now,
		}

		repo.CreateSystemConfiguration(ctx, sc1)
		repo.CreateSystemConfiguration(ctx, sc2)

		configurations, err := repo.GetSystemConfigurationsByCategory(ctx, "upload")
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(configurations), 2)
	})
}

// TestSystemConfigurationRepository_GetActiveSystemConfigurations tests retrieving active configurations
func TestSystemConfigurationRepository_GetActiveSystemConfigurations(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewSystemConfigurationRepository(testDB.Pool)

	t.Run("Success - Get active configurations", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		now := time.Now()

		sc1 := &domain.SystemConfiguration{
			ID:        "sc-001",
			Key:       "ACTIVE",
			Value:     "value1",
			ValueType: "string",
			Category:  "test",
			IsSystem:  false,
			IsActive:  true,
			CreatedBy: "system",
			CreatedAt: now,
			UpdatedAt: now,
		}

		sc2 := &domain.SystemConfiguration{
			ID:        "sc-002",
			Key:       "INACTIVE",
			Value:     "value2",
			ValueType: "string",
			Category:  "test",
			IsSystem:  false,
			IsActive:  false,
			CreatedBy: "system",
			CreatedAt: now,
			UpdatedAt: now,
		}

		repo.CreateSystemConfiguration(ctx, sc1)
		repo.CreateSystemConfiguration(ctx, sc2)

		configurations, err := repo.GetActiveSystemConfigurations(ctx)
		require.NoError(t, err)

		keys := make([]string, len(configurations))
		for i, c := range configurations {
			keys[i] = c.Key
		}
		assert.Contains(t, keys, "ACTIVE")
		assert.NotContains(t, keys, "INACTIVE")
	})
}

// TestSystemConfigurationRepository_UpdateSystemConfiguration tests updating a configuration
func TestSystemConfigurationRepository_UpdateSystemConfiguration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewSystemConfigurationRepository(testDB.Pool)

	t.Run("Success - Update configuration", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		now := time.Now()

		sc := &domain.SystemConfiguration{
			ID:        "sc-001",
			Key:       "max_file_size",
			Value:     "10485760",
			ValueType: "integer",
			Category:  "upload",
			IsSystem:  false,
			IsActive:  true,
			CreatedBy: "system",
			CreatedAt: now,
			UpdatedAt: now,
		}

		err := repo.CreateSystemConfiguration(ctx, sc)
		require.NoError(t, err)

		sc.Value = "20971520"
		sc.UpdatedBy = StringPtr("admin")
		sc.UpdatedAt = time.Now()

		err = repo.UpdateSystemConfiguration(ctx, sc)
		require.NoError(t, err)

		updated, err := repo.GetSystemConfigurationByID(ctx, sc.ID)
		require.NoError(t, err)
		assert.Equal(t, "20971520", updated.Value)
	})

	t.Run("Error - Configuration not found", func(t *testing.T) {
		ctx := context.Background()

		sc := &domain.SystemConfiguration{
			ID:   "non-existent-id",
			Key:  "non-existent",
			Value: "value",
		}

		err := repo.UpdateSystemConfiguration(ctx, sc)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "system configuration not found")
	})
}

// TestSystemConfigurationRepository_DeleteSystemConfiguration tests deleting a configuration
func TestSystemConfigurationRepository_DeleteSystemConfiguration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewSystemConfigurationRepository(testDB.Pool)

	t.Run("Success - Delete non-system configuration", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		now := time.Now()

		sc := &domain.SystemConfiguration{
			ID:        "sc-001",
			Key:       "deletable_config",
			Value:     "value",
			ValueType: "string",
			Category:  "test",
			IsSystem:  false,
			IsActive:  true,
			CreatedBy: "system",
			CreatedAt: now,
			UpdatedAt: now,
		}

		err := repo.CreateSystemConfiguration(ctx, sc)
		require.NoError(t, err)

		err = repo.DeleteSystemConfiguration(ctx, sc.ID)
		require.NoError(t, err)

		_, err = repo.GetSystemConfigurationByID(ctx, sc.ID)
		assert.Error(t, err)
	})

	t.Run("Error - Cannot delete system configuration", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		now := time.Now()

		sc := &domain.SystemConfiguration{
			ID:        "sc-001",
			Key:       "system_config",
			Value:     "value",
			ValueType: "string",
			Category:  "system",
			IsSystem:  true,
			IsActive:  true,
			CreatedBy: "system",
			CreatedAt: now,
			UpdatedAt: now,
		}

		err := repo.CreateSystemConfiguration(ctx, sc)
		require.NoError(t, err)

		err = repo.DeleteSystemConfiguration(ctx, sc.ID)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "system configuration not found or is a system configuration that cannot be deleted")
	})
}

// TestSystemConfigurationRepository_CheckKeyExists tests checking key uniqueness
func TestSystemConfigurationRepository_CheckKeyExists(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping repository test in short mode")
	}

	testDB := SetupTestDB(t)
	defer testDB.Cleanup()
	defer CleanupTestData(t, testDB.Pool)

	repo := NewSystemConfigurationRepository(testDB.Pool)

	t.Run("Success - Key exists", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()
		now := time.Now()

		sc := &domain.SystemConfiguration{
			ID:        "sc-001",
			Key:       "UNIQUE_KEY",
			Value:     "value",
			ValueType: "string",
			Category:  "test",
			IsSystem:  false,
			IsActive:  true,
			CreatedBy: "system",
			CreatedAt: now,
			UpdatedAt: now,
		}

		repo.CreateSystemConfiguration(ctx, sc)

		exists, err := repo.CheckKeyExists(ctx, "UNIQUE_KEY", "")
		require.NoError(t, err)
		assert.True(t, exists)
	})

	t.Run("Success - Key does not exist", func(t *testing.T) {
		CleanupTestData(t, testDB.Pool)

		ctx := context.Background()

		exists, err := repo.CheckKeyExists(ctx, "NONEXISTENT_KEY", "")
		require.NoError(t, err)
		assert.False(t, exists)
	})
}

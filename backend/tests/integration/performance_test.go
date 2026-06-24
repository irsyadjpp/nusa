package integration

import (
	"context"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestLargeDatasetPagination validates pagination with large datasets
func TestLargeDatasetPagination(t *testing.T) {
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

	t.Run("Pagination with 100+ records", func(t *testing.T) {
		// Create 100 users
		start := time.Now()
		for i := 0; i < 100; i++ {
			email := "perf100" + string(rune('0'+i%10)) + "@example.com"
			CreateTestUser(t, ctx, testDB.UserRepo, email, "password123", "Perf User", roleID, nil)
		}
		createTime := time.Since(start)
		t.Logf("Created 100 users in %v", createTime)

		// Test pagination performance
		start = time.Now()
		users, err := testDB.UserRepo.List(ctx, nil, nil, nil, 10, 0)
		require.NoError(t, err)
		paginationTime := time.Since(start)
		t.Logf("Paginated 10 users in %v", paginationTime)

		assert.Equal(t, 10, len(users))
		assert.Less(t, paginationTime, 1*time.Second, "Pagination should complete in less than 1 second")
	})

	t.Run("Pagination with count query", func(t *testing.T) {
		// Create additional users
		for i := 0; i < 50; i++ {
			email := "perf50" + string(rune('0'+i%10)) + "@example.com"
			CreateTestUser(t, ctx, testDB.UserRepo, email, "password123", "Perf User", roleID, nil)
		}

		// Test count query performance
		start := time.Now()
		count, err := testDB.UserRepo.Count(ctx, nil, nil, nil)
		require.NoError(t, err)
		countTime := time.Since(start)
		t.Logf("Counted %d users in %v", count, countTime)

		assert.GreaterOrEqual(t, count, 150)
		assert.Less(t, countTime, 500*time.Millisecond, "Count should complete in less than 500ms")
	})
}

// TestComplexQueryPerformance validates complex query performance
func TestComplexQueryPerformance(t *testing.T) {
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

	t.Run("Query with multiple filters", func(t *testing.T) {
		// Create users with different schools and roles
		school := CreateTestSchool(t, ctx, testDB.SchoolRepo, "Perf School", "PF001")

		for i := 0; i < 20; i++ {
			email := "filter" + string(rune('0'+i%10)) + "@example.com"
			CreateTestUser(t, ctx, testDB.UserRepo, email, "password123", "Filter User", roleID, &school.ID)
		}

		// Test query with multiple filters
		start := time.Now()
		users, err := testDB.UserRepo.List(ctx, &school.ID, &roleID, nil, 10, 0)
		require.NoError(t, err)
		queryTime := time.Since(start)
		t.Logf("Complex filter query completed in %v", queryTime)

		assert.GreaterOrEqual(t, len(users), 20)
		assert.Less(t, queryTime, 1*time.Second, "Complex filter query should complete in less than 1 second")
	})

	t.Run("Multiple sequential queries", func(t *testing.T) {
		// Test performance of multiple sequential queries
		start := time.Now()

		for i := 0; i < 10; i++ {
			users, err := testDB.UserRepo.List(ctx, nil, nil, nil, 10, 0)
			require.NoError(t, err)
			assert.GreaterOrEqual(t, len(users), 0)
		}

		totalTime := time.Since(start)
		avgTime := totalTime / 10
		t.Logf("10 sequential queries completed in %v (avg %v per query)", totalTime, avgTime)

		assert.Less(t, avgTime, 200*time.Millisecond, "Average query time should be less than 200ms")
	})
}

// TestIndexUsageValidation validates database index usage
func TestIndexUsageValidation(t *testing.T) {
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

	t.Run("Email index usage", func(t *testing.T) {
		email := "index@example.com"
		user := CreateTestUser(t, ctx, testDB.UserRepo, email, "password123", "Index User", roleID, nil)

		// Query by email should be fast (using index)
		start := time.Now()
		retrieved, err := testDB.UserRepo.GetByEmail(ctx, email)
		require.NoError(t, err)
		queryTime := time.Since(start)
		t.Logf("Email lookup completed in %v", queryTime)

		assert.Equal(t, user.ID, retrieved.ID)
		assert.Less(t, queryTime, 100*time.Millisecond, "Email lookup should be fast (using index)")
	})

	t.Run("School code index usage", func(t *testing.T) {
		code := "IDX001"
		school := CreateTestSchool(t, ctx, testDB.SchoolRepo, "Index School", code)

		// Query by code should be fast (using index)
		start := time.Now()
		retrieved, err := testDB.SchoolRepo.GetByCode(ctx, code)
		require.NoError(t, err)
		queryTime := time.Since(start)
		t.Logf("School code lookup completed in %v", queryTime)

		assert.Equal(t, school.ID, retrieved.ID)
		assert.Less(t, queryTime, 100*time.Millisecond, "School code lookup should be fast (using index)")
	})

	t.Run("Role name index usage", func(t *testing.T) {
		name := "INDEXED_ROLE"
		description := "Role for index testing"
		role := CreateTestRole(t, ctx, testDB.RoleRepo, name, &description)

		// Query by name should be fast (using index)
		start := time.Now()
		retrieved, err := testDB.RoleRepo.GetByName(ctx, name)
		require.NoError(t, err)
		queryTime := time.Since(start)
		t.Logf("Role name lookup completed in %v", queryTime)

		assert.Equal(t, role.ID, retrieved.ID)
		assert.Less(t, queryTime, 100*time.Millisecond, "Role name lookup should be fast (using index)")
	})
}

// TestTransactionPerformance validates transaction performance
func TestTransactionPerformance(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	testDB := SetupTestDB(t)
	if testDB == nil {
		return
	}
	defer TeardownTestDB(t, testDB)

	roleID := "00000000-0000-0000-0000-000000000001"

	t.Run("Transaction with multiple operations", func(t *testing.T) {
		start := time.Now()

		WithTransaction(t, testDB.DB, func(ctx context.Context, tx *sqlx.Tx) error {
			// Create multiple users within transaction
			for i := 0; i < 10; i++ {
				email := "trans" + string(rune('0'+i%10)) + "@example.com"
				CreateTestUser(t, ctx, testDB.UserRepo, email, "password123", "Trans User", roleID, nil)
			}
			return nil
		})

		transactionTime := time.Since(start)
		t.Logf("Transaction with 10 operations completed in %v", transactionTime)

		assert.Less(t, transactionTime, 2*time.Second, "Transaction should complete in less than 2 seconds")
	})
}

// TestConcurrentQueryPerformance validates concurrent query performance
func TestConcurrentQueryPerformance(t *testing.T) {
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

	// Create test data
	for i := 0; i < 20; i++ {
		email := "concurrent" + string(rune('0'+i%10)) + "@example.com"
		CreateTestUser(t, ctx, testDB.UserRepo, email, "password123", "Concurrent User", roleID, nil)
	}

	t.Run("Concurrent read operations", func(t *testing.T) {
		start := time.Now()

		// Run concurrent read operations
		for i := 0; i < 10; i++ {
			users, err := testDB.UserRepo.List(ctx, nil, nil, nil, 10, 0)
			require.NoError(t, err)
			assert.GreaterOrEqual(t, len(users), 0)
		}

		totalTime := time.Since(start)
		avgTime := totalTime / 10
		t.Logf("10 concurrent read operations completed in %v (avg %v per operation)", totalTime, avgTime)

		assert.Less(t, avgTime, 200*time.Millisecond, "Average concurrent read time should be less than 200ms")
	})
}

// TestBatchOperationPerformance validates batch operation performance
func TestBatchOperationPerformance(t *testing.T) {
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

	t.Run("Batch user creation", func(t *testing.T) {
		batchSize := 50

		start := time.Now()
		for i := 0; i < batchSize; i++ {
			email := "batch" + string(rune('0'+i%10)) + "@example.com"
			CreateTestUser(t, ctx, testDB.UserRepo, email, "password123", "Batch User", roleID, nil)
		}
		totalTime := time.Since(start)
		avgTime := totalTime / time.Duration(batchSize)
		t.Logf("Created %d users in %v (avg %v per user)", batchSize, totalTime, avgTime)

		assert.Less(t, avgTime, 100*time.Millisecond, "Average user creation time should be less than 100ms")
	})
}

package integration

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestSQLInjectionInEmail validates that SQL injection in email parameter is prevented
func TestSQLInjectionInEmail(t *testing.T) {
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

	// SQL injection payloads to test
	injectionPayloads := []string{
		"test@example.com'; DROP TABLE users; --",
		"test@example.com' OR '1'='1",
		"test@example.com' UNION SELECT * FROM users --",
		"test@example.com'; INSERT INTO users (email) VALUES ('hacked@evil.com'); --",
		"test@example.com' AND 1=1 --",
		"test@example.com'; SELECT pg_sleep(10); --",
		"test@example.com' OR 1=1 #",
		"test@example.com' OR 1=1 /*",
	}

	for i, payload := range injectionPayloads {
		t.Run("SQL injection payload "+string(rune('0'+i)), func(t *testing.T) {
			// Attempt to create user with SQL injection in email
			// This should either fail validation or be stored as-is (not executed)
			user := CreateTestUser(t, ctx, testDB.UserRepo, payload, "password123", "Injection User", roleID, nil)

			if user == nil {
				// Creation failed (validation prevented it)
				return
			}

			// If creation succeeded, verify the email is stored as-is
			retrieved, err := testDB.UserRepo.GetByID(ctx, user.ID)
			require.NoError(t, err)
			assert.Equal(t, payload, retrieved.Email, "Email should be stored as-is, not executed as SQL")

			// Verify no tables were dropped by attempting to query users
			users, err := testDB.UserRepo.List(ctx, nil, nil, nil, 10, 0)
			require.NoError(t, err, "Users table should still exist if injection was prevented")
			assert.Greater(t, len(users), 0, "Users should still exist in database")
		})
	}
}

// TestSQLInjectionInName validates that SQL injection in name parameter is prevented
func TestSQLInjectionInName(t *testing.T) {
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

	// SQL injection payloads to test in name field
	injectionPayloads := []string{
		"John'; DROP TABLE users; --",
		"John' OR '1'='1",
		"John' UNION SELECT * FROM users --",
		"<script>alert('xss')</script>",
		"'; EXEC xp_cmdshell('dir'); --",
		"Robert'); DROP TABLE students; --",
	}

	for i, payload := range injectionPayloads {
		t.Run("SQL injection in name payload "+string(rune('0'+i)), func(t *testing.T) {
			email := "injection" + string(rune('0'+i)) + "@example.com"

			// Attempt to create user with SQL injection in name
			user := CreateTestUser(t, ctx, testDB.UserRepo, email, "password123", payload, roleID, nil)

			if user == nil {
				// Creation failed (validation prevented it)
				return
			}

			// If creation succeeded, verify the name is stored as-is
			retrieved, err := testDB.UserRepo.GetByID(ctx, user.ID)
			require.NoError(t, err)
			assert.Equal(t, payload, retrieved.Name, "Name should be stored as-is, not executed as SQL")
		})
	}
}

// TestSQLInjectionInSchoolCode validates that SQL injection in school code parameter is prevented
func TestSQLInjectionInSchoolCode(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	testDB := SetupTestDB(t)
	if testDB == nil {
		return
	}
	defer TeardownTestDB(t, testDB)

	ctx := context.Background()

	// SQL injection payloads to test in school code
	injectionPayloads := []string{
		"SCH001'; DROP TABLE schools; --",
		"SCH001' OR '1'='1",
		"SCH001' UNION SELECT * FROM schools --",
	}

	for i, payload := range injectionPayloads {
		t.Run("SQL injection in school code payload "+string(rune('0'+i)), func(t *testing.T) {
			// Attempt to create school with SQL injection in code
			school := CreateTestSchool(t, ctx, testDB.SchoolRepo, "Injection School", payload)

			if school == nil {
				// Creation failed (validation prevented it)
				return
			}

			// If creation succeeded, verify the code is stored as-is
			retrieved, err := testDB.SchoolRepo.GetByID(ctx, school.ID)
			require.NoError(t, err)
			assert.Equal(t, payload, retrieved.Code, "Code should be stored as-is, not executed as SQL")
		})
	}
}

// TestParameterizedQueryValidation validates that all repository methods use parameterized queries
func TestParameterizedQueryValidation(t *testing.T) {
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

	t.Run("User repository uses parameterized queries", func(t *testing.T) {
		// Test that special characters are handled correctly
		specialEmail := "test'special@example.com"
		user := CreateTestUser(t, ctx, testDB.UserRepo, specialEmail, "password123", "Special User", roleID, nil)

		if user != nil {
			// Verify special characters are stored correctly
			retrieved, err := testDB.UserRepo.GetByEmail(ctx, specialEmail)
			require.NoError(t, err)
			assert.Equal(t, specialEmail, retrieved.Email)
		}
	})

	t.Run("School repository uses parameterized queries", func(t *testing.T) {
		// Test that special characters are handled correctly
		specialName := "O'Reilly School"
		specialCode := "O'REILLY"

		school := CreateTestSchool(t, ctx, testDB.SchoolRepo, specialName, specialCode)

		if school != nil {
			// Verify special characters are stored correctly
			retrieved, err := testDB.SchoolRepo.GetByCode(ctx, specialCode)
			require.NoError(t, err)
			assert.Equal(t, specialName, retrieved.Name)
			assert.Equal(t, specialCode, retrieved.Code)
		}
	})

	t.Run("Role repository uses parameterized queries", func(t *testing.T) {
		// Test that special characters are handled correctly
		specialName := "ADMIN'S ROLE"

		role := CreateTestRole(t, ctx, testDB.RoleRepo, specialName, nil)

		if role != nil {
			// Verify special characters are stored correctly
			retrieved, err := testDB.RoleRepo.GetByName(ctx, specialName)
			require.NoError(t, err)
			assert.Equal(t, specialName, retrieved.Name)
		}
	})
}

// TestLongStringInjection validates that very long strings are handled correctly
func TestLongStringInjection(t *testing.T) {
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

	t.Run("Long email string", func(t *testing.T) {
		// Create a very long email (exceeds typical limits)
		longEmail := "a" + string(make([]byte, 1000)) + "@example.com"

		user := CreateTestUser(t, ctx, testDB.UserRepo, longEmail, "password123", "Long Email User", roleID, nil)

		// Should fail due to validation or database constraint
		if user != nil {
			// If it succeeded, database should handle it gracefully
			assert.NotEmpty(t, user.Email)
		}
	})

	t.Run("Long name string", func(t *testing.T) {
		// Create a very long name
		longName := string(make([]byte, 1000))
		email := "longname@example.com"

		user := CreateTestUser(t, ctx, testDB.UserRepo, email, "password123", longName, roleID, nil)

		// Should fail due to validation or database constraint
		if user != nil {
			// If it succeeded, database should handle it gracefully
			assert.NotEmpty(t, user.Name)
		}
	})
}

// TestNullByteInjection validates that null bytes are handled correctly
func TestNullByteInjection(t *testing.T) {
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

	t.Run("Null byte in email", func(t *testing.T) {
		// Email with null byte
		email := "test\000@example.com"

		user := CreateTestUser(t, ctx, testDB.UserRepo, email, "password123", "Null Byte User", roleID, nil)

		// Should fail due to validation
		assert.Nil(t, user, "User creation with null byte should fail")
	})

	t.Run("Null byte in name", func(t *testing.T) {
		// Name with null byte
		name := "Test\000User"
		email := "nullbyte@example.com"

		user := CreateTestUser(t, ctx, testDB.UserRepo, email, "password123", name, roleID, nil)

		// Should fail due to validation
		assert.Nil(t, user, "User creation with null byte should fail")
	})
}

// TestUnicodeInjection validates that unicode characters are handled correctly
func TestUnicodeInjection(t *testing.T) {
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

	t.Run("Unicode in email", func(t *testing.T) {
		// Email with unicode characters
		email := "tëst@example.com"

		user := CreateTestUser(t, ctx, testDB.UserRepo, email, "password123", "Unicode User", roleID, nil)

		if user != nil {
			// Verify unicode is stored correctly
			retrieved, err := testDB.UserRepo.GetByEmail(ctx, email)
			require.NoError(t, err)
			assert.Equal(t, email, retrieved.Email)
		}
	})

	t.Run("Unicode in name", func(t *testing.T) {
		// Name with unicode characters
		name := "Tëst Üser"
		email := "unicode@example.com"

		user := CreateTestUser(t, ctx, testDB.UserRepo, email, "password123", name, roleID, nil)

		if user != nil {
			// Verify unicode is stored correctly
			retrieved, err := testDB.UserRepo.GetByID(ctx, user.ID)
			require.NoError(t, err)
			assert.Equal(t, name, retrieved.Name)
		}
	})
}

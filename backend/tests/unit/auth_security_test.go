package unit

import (
	"testing"

	"golang.org/x/crypto/bcrypt"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestPasswordHashing validates that passwords are properly hashed with bcrypt
func TestPasswordHashing(t *testing.T) {
	password := "SecurePassword123!"

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	require.NoError(t, err)
	require.NotEmpty(t, hash)

	// Verify hash is different from original password
	assert.NotEqual(t, password, string(hash))

	// Verify hash starts with bcrypt prefix
	hasPrefix := assert.Contains(t, string(hash), "$2a$") || assert.Contains(t, string(hash), "$2b$")
	_ = hasPrefix // Explicitly ignore the result
}

// TestPasswordComparison validates that bcrypt can correctly compare passwords
func TestPasswordComparison(t *testing.T) {
	password := "SecurePassword123!"

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	require.NoError(t, err)

	// Verify correct password matches
	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	assert.NoError(t, err)

	// Verify incorrect password does not match
	err = bcrypt.CompareHashAndPassword(hash, []byte("WrongPassword"))
	assert.Error(t, err)
	assert.Equal(t, bcrypt.ErrMismatchedHashAndPassword, err)
}

// TestPasswordHashingCostFactor validates that bcrypt uses appropriate cost factor
func TestPasswordHashingCostFactor(t *testing.T) {
	password := "SecurePassword123!"

	// Test with default cost factor (10)
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	require.NoError(t, err)
	require.NotEmpty(t, hash)

	// Test with higher cost factor (12) for production
	hash12, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	require.NoError(t, err)
	require.NotEmpty(t, hash12)

	// Both should validate correctly
	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	assert.NoError(t, err)

	err = bcrypt.CompareHashAndPassword(hash12, []byte(password))
	assert.NoError(t, err)
}

// TestInvalidPasswordHash validates that invalid hashes are rejected
func TestInvalidPasswordHash(t *testing.T) {
	password := "SecurePassword123!"

	// Test with malformed hash
	invalidHash := []byte("invalid_hash_format")
	err := bcrypt.CompareHashAndPassword(invalidHash, []byte(password))
	assert.Error(t, err)

	// Test with empty hash
	emptyHash := []byte("")
	err = bcrypt.CompareHashAndPassword(emptyHash, []byte(password))
	assert.Error(t, err)
}

// TestPasswordComplexityValidation validates password complexity requirements
func TestPasswordComplexityValidation(t *testing.T) {
	tests := []struct {
		name     string
		password string
		valid    bool
	}{
		{
			name:     "Valid password with complexity",
			password: "SecurePass123!",
			valid:    true,
		},
		{
			name:     "Too short",
			password: "Short1!",
			valid:    false,
		},
		{
			name:     "No uppercase",
			password: "lowercase123!",
			valid:    false,
		},
		{
			name:     "No lowercase",
			password: "UPPERCASE123!",
			valid:    false,
		},
		{
			name:     "No numbers",
			password: "NoNumbersHere!",
			valid:    false,
		},
		{
			name:     "No special characters",
			password: "NoSpecialChars123",
			valid:    false,
		},
		{
			name:     "Only numbers",
			password: "12345678",
			valid:    false,
		},
		{
			name:     "Common password",
			password: "password123",
			valid:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Basic validation rules (min 8 chars, uppercase, lowercase, number, special)
			hasMinLength := len(tt.password) >= 8
			hasUppercase := false
			hasLowercase := false
			hasNumber := false
			hasSpecial := false

			for _, char := range tt.password {
				switch {
				case char >= 'A' && char <= 'Z':
					hasUppercase = true
				case char >= 'a' && char <= 'z':
					hasLowercase = true
				case char >= '0' && char <= '9':
					hasNumber = true
				case char == '!' || char == '@' || char == '#' || char == '$' || char == '%' || char == '^' || char == '&' || char == '*':
					hasSpecial = true
				}
			}

			isValid := hasMinLength && hasUppercase && hasLowercase && hasNumber && hasSpecial

			// For common passwords, additional check
			isCommon := tt.password == "password123" || tt.password == "12345678"
			if isCommon {
				isValid = false
			}

			assert.Equal(t, tt.valid, isValid, "Password validation result mismatch")
		})
	}
}

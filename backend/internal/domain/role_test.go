package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasPermission(t *testing.T) {
	tests := []struct {
		name     string
		role     string
		resource string
		action   string
		expected bool
	}{
		{
			name:     "SYSTEM_ADMIN has user:CREATE permission",
			role:     RoleSystemAdmin,
			resource: "user",
			action:   "CREATE",
			expected: true,
		},
		{
			name:     "SYSTEM_ADMIN has school:READ permission",
			role:     RoleSystemAdmin,
			resource: "school",
			action:   "READ",
			expected: true,
		},
		{
			name:     "SCHOOL_ADMIN has user:CREATE permission",
			role:     RoleSchoolAdmin,
			resource: "user",
			action:   "CREATE",
			expected: true,
		},
		{
			name:     "SCHOOL_ADMIN does not have school:CREATE permission",
			role:     RoleSchoolAdmin,
			resource: "school",
			action:   "CREATE",
			expected: false,
		},
		{
			name:     "TEACHER has tp:READ permission",
			role:     RoleTeacher,
			resource: "tp",
			action:   "READ",
			expected: true,
		},
		{
			name:     "TEACHER does not have user:CREATE permission",
			role:     RoleTeacher,
			resource: "user",
			action:   "CREATE",
			expected: false,
		},
		{
			name:     "Unknown role has no permissions",
			role:     "UNKNOWN_ROLE",
			resource: "user",
			action:   "CREATE",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HasPermission(tt.role, tt.resource, tt.action)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGetRolePermissions(t *testing.T) {
	permissions := GetRolePermissions()

	// Check that all expected roles exist
	assert.Contains(t, permissions, RoleSystemAdmin)
	assert.Contains(t, permissions, RoleSchoolAdmin)
	assert.Contains(t, permissions, RoleTeacher)

	// Check that SYSTEM_ADMIN has all permissions
	systemAdminPerms := permissions[RoleSystemAdmin]
	assert.NotEmpty(t, systemAdminPerms)
	assert.Contains(t, systemAdminPerms, "user:CREATE")
	assert.Contains(t, systemAdminPerms, "school:CREATE")

	// Check that SCHOOL_ADMIN has expected permissions
	schoolAdminPerms := permissions[RoleSchoolAdmin]
	assert.NotEmpty(t, schoolAdminPerms)
	assert.Contains(t, schoolAdminPerms, "user:CREATE")
	assert.NotContains(t, schoolAdminPerms, "school:CREATE")

	// Check that TEACHER has expected permissions
	teacherPerms := permissions[RoleTeacher]
	assert.NotEmpty(t, teacherPerms)
	assert.Contains(t, teacherPerms, "tp:READ")
	assert.NotContains(t, teacherPerms, "user:CREATE")
}

func TestUserStatus(t *testing.T) {
	assert.Equal(t, UserStatus("ACTIVE"), UserStatusActive)
	assert.Equal(t, UserStatus("INACTIVE"), UserStatusInactive)
	assert.Equal(t, UserStatus("SUSPENDED"), UserStatusSuspended)
}

func TestSchoolStatus(t *testing.T) {
	assert.Equal(t, SchoolStatus("ACTIVE"), SchoolStatusActive)
	assert.Equal(t, SchoolStatus("INACTIVE"), SchoolStatusInactive)
}

func TestNewID(t *testing.T) {
	id := NewID()
	assert.NotEmpty(t, id)
	assert.NotEqual(t, NewID(), id) // Different IDs should be generated
}

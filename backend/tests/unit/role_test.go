package unit

import (
	"testing"

	"github.com/nusa/backend/internal/domain"
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
			role:     domain.RoleSystemAdmin,
			resource: "user",
			action:   "CREATE",
			expected: true,
		},
		{
			name:     "SYSTEM_ADMIN has school:READ permission",
			role:     domain.RoleSystemAdmin,
			resource: "school",
			action:   "READ",
			expected: true,
		},
		{
			name:     "SCHOOL_ADMIN has user:CREATE permission",
			role:     domain.RoleSchoolAdmin,
			resource: "user",
			action:   "CREATE",
			expected: true,
		},
		{
			name:     "SCHOOL_ADMIN does not have school:CREATE permission",
			role:     domain.RoleSchoolAdmin,
			resource: "school",
			action:   "CREATE",
			expected: false,
		},
		{
			name:     "TEACHER has tp:READ permission",
			role:     domain.RoleTeacher,
			resource: "tp",
			action:   "READ",
			expected: true,
		},
		{
			name:     "TEACHER does not have user:CREATE permission",
			role:     domain.RoleTeacher,
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
			result := domain.HasPermission(tt.role, tt.resource, tt.action)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGetRolePermissions(t *testing.T) {
	permissions := domain.GetRolePermissions()

	// Check that all expected roles exist
	assert.Contains(t, permissions, domain.RoleSystemAdmin)
	assert.Contains(t, permissions, domain.RoleSchoolAdmin)
	assert.Contains(t, permissions, domain.RoleTeacher)

	// Check that SYSTEM_ADMIN has all permissions
	systemAdminPerms := permissions[domain.RoleSystemAdmin]
	assert.NotEmpty(t, systemAdminPerms)
	assert.Contains(t, systemAdminPerms, "user:CREATE")
	assert.Contains(t, systemAdminPerms, "school:CREATE")

	// Check that SCHOOL_ADMIN has expected permissions
	schoolAdminPerms := permissions[domain.RoleSchoolAdmin]
	assert.NotEmpty(t, schoolAdminPerms)
	assert.Contains(t, schoolAdminPerms, "user:CREATE")
	assert.NotContains(t, schoolAdminPerms, "school:CREATE")

	// Check that TEACHER has expected permissions
	teacherPerms := permissions[domain.RoleTeacher]
	assert.NotEmpty(t, teacherPerms)
	assert.Contains(t, teacherPerms, "tp:READ")
	assert.NotContains(t, teacherPerms, "user:CREATE")
}

func TestUserStatus(t *testing.T) {
	assert.Equal(t, domain.UserStatus("ACTIVE"), domain.UserStatusActive)
	assert.Equal(t, domain.UserStatus("INACTIVE"), domain.UserStatusInactive)
	assert.Equal(t, domain.UserStatus("SUSPENDED"), domain.UserStatusSuspended)
}

func TestSchoolStatus(t *testing.T) {
	assert.Equal(t, domain.SchoolStatus("ACTIVE"), domain.SchoolStatusActive)
	assert.Equal(t, domain.SchoolStatus("INACTIVE"), domain.SchoolStatusInactive)
}

func TestNewID(t *testing.T) {
	id := domain.NewID()
	assert.NotEmpty(t, id)
	assert.NotEqual(t, domain.NewID(), id) // Different IDs should be generated
}

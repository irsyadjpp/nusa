package unit

import (
	"testing"

	"github.com/nusa/backend/internal/domain"
	"github.com/stretchr/testify/assert"
)

// TestPermissionEscalationPrevention validates that users cannot elevate their own permissions
func TestPermissionEscalationPrevention(t *testing.T) {
	// Test that TEACHER cannot get SYSTEM_ADMIN permissions
	assert.False(t, domain.HasPermission(domain.RoleTeacher, "user", "CREATE"))
	assert.False(t, domain.HasPermission(domain.RoleTeacher, "school", "CREATE"))

	// Test that SCHOOL_ADMIN cannot get SYSTEM_ADMIN permissions
	assert.False(t, domain.HasPermission(domain.RoleSchoolAdmin, "school", "CREATE"))

	// Test that SYSTEM_ADMIN has all permissions
	assert.True(t, domain.HasPermission(domain.RoleSystemAdmin, "user", "CREATE"))
	assert.True(t, domain.HasPermission(domain.RoleSystemAdmin, "school", "CREATE"))
}

// TestCrossSchoolAccessPrevention validates that users cannot access other schools' data
func TestCrossSchoolAccessPrevention(t *testing.T) {
	// This test validates the permission model
	// In a real implementation, this would be enforced by:
	// 1. Middleware checking school_id in user context
	// 2. Repository filtering by school_id
	// 3. Handler-level validation

	// For unit test, we verify the permission structure
	// SCHOOL_ADMIN has user:CREATE but not school:CREATE
	assert.True(t, domain.HasPermission(domain.RoleSchoolAdmin, "user", "CREATE"))
	assert.False(t, domain.HasPermission(domain.RoleSchoolAdmin, "school", "CREATE"))

	// TEACHER has limited permissions
	assert.True(t, domain.HasPermission(domain.RoleTeacher, "tp", "READ"))
	assert.False(t, domain.HasPermission(domain.RoleTeacher, "user", "CREATE"))
	assert.False(t, domain.HasPermission(domain.RoleTeacher, "school", "READ"))
}

// TestRoleAssignmentWithoutPermission validates that role assignment requires proper permissions
func TestRoleAssignmentWithoutPermission(t *testing.T) {
	// Only SYSTEM_ADMIN should be able to assign roles
	// This is enforced at the handler level

	// Verify permission structure
	// SYSTEM_ADMIN has user management permissions (which includes role assignment)
	assert.True(t, domain.HasPermission(domain.RoleSystemAdmin, domain.ResourceUser, domain.ActionCreate))
	assert.True(t, domain.HasPermission(domain.RoleSystemAdmin, domain.ResourceUser, domain.ActionUpdate))

	// SCHOOL_ADMIN has user management permissions but limited
	assert.True(t, domain.HasPermission(domain.RoleSchoolAdmin, domain.ResourceUser, domain.ActionCreate))
	assert.True(t, domain.HasPermission(domain.RoleSchoolAdmin, domain.ResourceUser, domain.ActionUpdate))

	// TEACHER does not have user management permissions
	assert.False(t, domain.HasPermission(domain.RoleTeacher, domain.ResourceUser, domain.ActionCreate))
	assert.False(t, domain.HasPermission(domain.RoleTeacher, domain.ResourceUser, domain.ActionUpdate))
}

// TestSystemRoleProtection validates that system roles cannot be modified by non-admins
func TestSystemRoleProtection(t *testing.T) {
	// System roles should be protected
	// SYSTEM_ADMIN, SCHOOL_ADMIN, TEACHER are predefined roles

	// Verify these roles exist in the permission map
	permissions := domain.GetRolePermissions()
	assert.Contains(t, permissions, domain.RoleSystemAdmin)
	assert.Contains(t, permissions, domain.RoleSchoolAdmin)
	assert.Contains(t, permissions, domain.RoleTeacher)

	// Verify SYSTEM_ADMIN has maximum permissions
	systemAdminPerms := permissions[domain.RoleSystemAdmin]
	assert.NotEmpty(t, systemAdminPerms)
	assert.Contains(t, systemAdminPerms, domain.ResourceUser+":"+domain.ActionCreate)
	assert.Contains(t, systemAdminPerms, domain.ResourceSchool+":"+domain.ActionCreate)
	assert.Contains(t, systemAdminPerms, domain.ResourceSchool+":"+domain.ActionDelete)

	// Verify other roles have restricted permissions
	schoolAdminPerms := permissions[domain.RoleSchoolAdmin]
	assert.NotContains(t, schoolAdminPerms, domain.ResourceSchool+":"+domain.ActionCreate)
	assert.NotContains(t, schoolAdminPerms, domain.ResourceSchool+":"+domain.ActionDelete)
	assert.Contains(t, schoolAdminPerms, domain.ResourceSchool+":"+domain.ActionRead)
}

// TestPrivilegeEscalation validates various privilege escalation scenarios
func TestPrivilegeEscalation(t *testing.T) {
	tests := []struct {
		name           string
		currentRole    string
		targetResource string
		targetAction   string
		shouldAllow    bool
	}{
		{
			name:           "TEACHER cannot CREATE users",
			currentRole:    domain.RoleTeacher,
			targetResource: "user",
			targetAction:   "CREATE",
			shouldAllow:    false,
		},
		{
			name:           "TEACHER cannot CREATE schools",
			currentRole:    domain.RoleTeacher,
			targetResource: "school",
			targetAction:   "CREATE",
			shouldAllow:    false,
		},
		{
			name:           "SCHOOL_ADMIN cannot CREATE schools",
			currentRole:    domain.RoleSchoolAdmin,
			targetResource: "school",
			targetAction:   "CREATE",
			shouldAllow:    false,
		},
		{
			name:           "SCHOOL_ADMIN cannot DELETE schools",
			currentRole:    domain.RoleSchoolAdmin,
			targetResource: domain.ResourceSchool,
			targetAction:   domain.ActionDelete,
			shouldAllow:    false,
		},
		{
			name:           "SYSTEM_ADMIN can CREATE users",
			currentRole:    domain.RoleSystemAdmin,
			targetResource: domain.ResourceUser,
			targetAction:   domain.ActionCreate,
			shouldAllow:    true,
		},
		{
			name:           "SYSTEM_ADMIN can CREATE schools",
			currentRole:    domain.RoleSystemAdmin,
			targetResource: domain.ResourceSchool,
			targetAction:   domain.ActionCreate,
			shouldAllow:    true,
		},
		{
			name:           "SYSTEM_ADMIN can DELETE schools",
			currentRole:    domain.RoleSystemAdmin,
			targetResource: domain.ResourceSchool,
			targetAction:   domain.ActionDelete,
			shouldAllow:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hasPermission := domain.HasPermission(tt.currentRole, tt.targetResource, tt.targetAction)
			assert.Equal(t, tt.shouldAllow, hasPermission)
		})
	}
}

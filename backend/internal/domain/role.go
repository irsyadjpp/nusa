package domain

import "time"

// Role represents a user role in the system
type Role struct {
	ID          string    `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description *string   `json:"description,omitempty" db:"description"`
	IsActive    bool      `json:"is_active" db:"is_active"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// Permission represents a permission for a role
type Permission struct {
	ID        string    `json:"id" db:"id"`
	RoleID    string    `json:"role_id" db:"role_id"`
	Resource  string    `json:"resource" db:"resource"`
	Action    string    `json:"action" db:"action"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// Role constants
const (
	RoleSystemAdmin = "SYSTEM_ADMIN"
	RoleSchoolAdmin = "SCHOOL_ADMIN"
	RoleTeacher     = "TEACHER"
)

// Permission action constants
const (
	ActionCreate  = "CREATE"
	ActionRead    = "READ"
	ActionUpdate  = "UPDATE"
	ActionDelete  = "DELETE"
	ActionApprove = "APPROVE"
)

// Resource constants
const (
	ResourceSchool     = "school"
	ResourceUser       = "user"
	ResourceCurriculum = "curriculum"
	ResourceTP         = "tp"
	ResourceAssessment = "assessment"
	ResourceReporting  = "reporting"
)

// GetRolePermissions returns all permissions for a role
func GetRolePermissions() map[string][]string {
	return map[string][]string{
		RoleSystemAdmin: {
			ResourceSchool + ":" + ActionCreate,
			ResourceSchool + ":" + ActionUpdate,
			ResourceSchool + ":" + ActionDelete,
			ResourceSchool + ":" + ActionRead,
			ResourceUser + ":" + ActionCreate,
			ResourceUser + ":" + ActionUpdate,
			ResourceUser + ":" + ActionDelete,
			ResourceUser + ":" + ActionRead,
			ResourceCurriculum + ":" + ActionRead,
			ResourceCurriculum + ":" + ActionUpdate,
			ResourceCurriculum + ":" + ActionApprove,
			ResourceTP + ":" + ActionRead,
			ResourceTP + ":" + ActionCreate,
			ResourceTP + ":" + ActionUpdate,
			ResourceTP + ":" + ActionApprove,
			ResourceAssessment + ":" + ActionRead,
			ResourceAssessment + ":" + ActionCreate,
			ResourceAssessment + ":" + ActionUpdate,
			ResourceAssessment + ":" + ActionDelete,
			ResourceReporting + ":" + ActionRead,
		},
		RoleSchoolAdmin: {
			ResourceSchool + ":" + ActionRead,
			ResourceUser + ":" + ActionCreate,
			ResourceUser + ":" + ActionUpdate,
			ResourceUser + ":" + ActionRead,
			ResourceUser + ":" + ActionDelete,
			ResourceCurriculum + ":" + ActionRead,
			ResourceCurriculum + ":" + ActionUpdate,
			ResourceTP + ":" + ActionRead,
			ResourceTP + ":" + ActionCreate,
			ResourceTP + ":" + ActionApprove,
			ResourceAssessment + ":" + ActionRead,
			ResourceAssessment + ":" + ActionCreate,
			ResourceAssessment + ":" + ActionUpdate,
			ResourceReporting + ":" + ActionRead,
		},
		RoleTeacher: {
			ResourceTP + ":" + ActionRead,
			ResourceTP + ":" + ActionCreate,
			ResourceAssessment + ":" + ActionRead,
			ResourceAssessment + ":" + ActionCreate,
			ResourceReporting + ":" + ActionRead,
		},
	}
}

// HasPermission checks if a role has a specific permission
func HasPermission(role, resource, action string) bool {
	permissions := GetRolePermissions()
	rolePerms, exists := permissions[role]
	if !exists {
		return false
	}

	for _, perm := range rolePerms {
		if perm == resource+":"+action {
			return true
		}
	}
	return false
}

// CreateRoleRequest represents the request to create a new role
type CreateRoleRequest struct {
	Name        string  `json:"name" binding:"required,min=2"`
	Description *string `json:"description,omitempty"`
}

// UpdateRoleRequest represents the request to update a role
type UpdateRoleRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	IsActive    *bool   `json:"is_active,omitempty"`
}

// CreatePermissionRequest represents the request to add a permission to a role
type CreatePermissionRequest struct {
	Resource string `json:"resource" binding:"required"`
	Action   string `json:"action" binding:"required"`
}

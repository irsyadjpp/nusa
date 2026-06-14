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
	RoleSystemAdmin     = "SYSTEM_ADMIN"
	RoleSchoolAdmin     = "SCHOOL_ADMIN"
	RoleTeacher         = "TEACHER"
	RoleCurriculumAdmin = "CURRICULUM_ADMIN" // Sprint 4: Academic Foundation
)

// Permission action constants
const (
	ActionCreate   = "CREATE"
	ActionRead     = "READ"
	ActionUpdate   = "UPDATE"
	ActionDelete   = "DELETE"
	ActionApprove  = "APPROVE"
	ActionActivate = "ACTIVATE" // Sprint 4: For academic year activation
	ActionArchive  = "ARCHIVE"  // Sprint 4: For academic year archival
)

// Resource constants
const (
	ResourceSchool          = "school"
	ResourceUser            = "user"
	ResourceCurriculum      = "curriculum"
	ResourceTP              = "tp"
	ResourceAssessment      = "assessment"
	ResourceReporting       = "reporting"
	ResourceAcademicYear    = "academic_year"              // Sprint 4: Academic Foundation
	ResourceSemester        = "semester"                   // Sprint 4: Academic Foundation
	ResourceSubjectCategory = "subject_category"           // Sprint 4: Academic Foundation
	ResourceGraduateProfile = "graduate_profile_dimension" // Sprint 4: Academic Foundation
	ResourceCPAlignment     = "cp_alignment"               // Sprint 4: Academic Foundation
	ResourceSystemConfig    = "system_config"              // Sprint 4: Academic Foundation
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
			// Sprint 4: System Config permissions
			ResourceSystemConfig + ":" + ActionRead,
			ResourceSystemConfig + ":" + ActionUpdate,
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
			// Sprint 4: Academic Year permissions
			ResourceAcademicYear + ":" + ActionRead,
			ResourceAcademicYear + ":" + ActionCreate,
			ResourceAcademicYear + ":" + ActionUpdate,
			ResourceAcademicYear + ":" + ActionActivate,
			ResourceAcademicYear + ":" + ActionArchive,
			// Sprint 4: Semester permissions
			ResourceSemester + ":" + ActionRead,
			ResourceSemester + ":" + ActionCreate,
			ResourceSemester + ":" + ActionUpdate,
			ResourceSemester + ":" + ActionDelete,
			// Sprint 4: Alignment Report read access
			ResourceCPAlignment + ":" + ActionRead,
		},
		RoleCurriculumAdmin: {
			// Sprint 4: Curriculum Governance permissions
			ResourceCurriculum + ":" + ActionRead,
			ResourceCurriculum + ":" + ActionUpdate,
			ResourceSubjectCategory + ":" + ActionCreate,
			ResourceSubjectCategory + ":" + ActionUpdate,
			ResourceSubjectCategory + ":" + ActionDelete,
			ResourceGraduateProfile + ":" + ActionCreate,
			ResourceGraduateProfile + ":" + ActionUpdate,
			ResourceGraduateProfile + ":" + ActionDelete,
			ResourceCPAlignment + ":" + ActionCreate,
			ResourceCPAlignment + ":" + ActionUpdate,
			ResourceCPAlignment + ":" + ActionDelete,
			ResourceSystemConfig + ":" + ActionRead,
			// Academic Year read-only access
			ResourceAcademicYear + ":" + ActionRead,
			// Semester read-only access
			ResourceSemester + ":" + ActionRead,
		},
		RoleTeacher: {
			ResourceTP + ":" + ActionRead,
			ResourceTP + ":" + ActionCreate,
			ResourceAssessment + ":" + ActionRead,
			ResourceAssessment + ":" + ActionCreate,
			ResourceReporting + ":" + ActionRead,
			// Sprint 4: Academic Year read-only access
			ResourceAcademicYear + ":" + ActionRead,
			// Sprint 4: Subject Category read-only access
			ResourceSubjectCategory + ":" + ActionRead,
			// Sprint 4: Graduate Profile read-only access
			ResourceGraduateProfile + ":" + ActionRead,
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

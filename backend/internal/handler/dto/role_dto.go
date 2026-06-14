package dto

// CreateRoleRequest represents the request to create a role
type CreateRoleRequest struct {
	Name        string   `json:"name" binding:"required,min=2,max=100"`
	Description string   `json:"description" binding:"omitempty,max=500"`
	Permissions []string `json:"permissions" binding:"required"`
}

// UpdateRoleRequest represents the request to update a role
type UpdateRoleRequest struct {
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
}

// AddPermissionRequest represents the request to add a permission to a role
type AddPermissionRequest struct {
	Permission string `json:"permission" binding:"required"`
}

// RemovePermissionRequest represents the request to remove a permission from a role
type RemovePermissionRequest struct {
	Permission string `json:"permission" binding:"required"`
}

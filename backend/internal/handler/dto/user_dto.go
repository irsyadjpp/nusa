package dto

// CreateUserRequest represents the request to create a user
type CreateUserRequest struct {
	Email    string  `json:"email" binding:"required,email"`
	Password string  `json:"password" binding:"required,min=8"`
	FullName string  `json:"full_name" binding:"required"`
	RoleID   string  `json:"role_id" binding:"required"`
	SchoolID *string `json:"school_id,omitempty"`
	IsActive bool    `json:"is_active"`
}

// UpdateUserRequest represents the request to update a user
type UpdateUserRequest struct {
	FullName *string `json:"full_name,omitempty"`
	Email    *string `json:"email,omitempty"`
	RoleID   *string `json:"role_id,omitempty"`
	SchoolID *string `json:"school_id,omitempty"`
	IsActive *bool   `json:"is_active,omitempty"`
}

// UpdateUserStatusRequest represents the request to update user status
type UpdateUserStatusRequest struct {
	IsActive bool `json:"is_active" binding:"required"`
}

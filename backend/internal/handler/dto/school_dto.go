package dto

// CreateSchoolRequest represents the request to create a school
type CreateSchoolRequest struct {
	Name        string `json:"name" binding:"required,min=2,max=200"`
	Code        string `json:"code" binding:"required,min=2,max=20"`
	Address     string `json:"address" binding:"required,max=500"`
	Phone       string `json:"phone" binding:"omitempty,max=20"`
	Email       string `json:"email" binding:"omitempty,email,max=100"`
	PrincipalID string `json:"principal_id,omitempty"`
}

// UpdateSchoolRequest represents the request to update a school
type UpdateSchoolRequest struct {
	Name        *string `json:"name,omitempty"`
	Code        *string `json:"code,omitempty"`
	Address     *string `json:"address,omitempty"`
	Phone       *string `json:"phone,omitempty"`
	Email       *string `json:"email,omitempty"`
	PrincipalID *string `json:"principal_id,omitempty"`
}

// UpdateSchoolStatusRequest represents the request to update school status
type UpdateSchoolStatusRequest struct {
	IsActive bool `json:"is_active" binding:"required"`
}

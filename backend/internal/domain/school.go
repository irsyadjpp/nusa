package domain

import "time"

// SchoolStatus represents the status of a school
type SchoolStatus string

const (
	SchoolStatusActive   SchoolStatus = "ACTIVE"
	SchoolStatusInactive SchoolStatus = "INACTIVE"
)

// School represents a school in the system
type School struct {
	ID        string       `json:"id" db:"id"`
	Name      string       `json:"name" db:"name"`
	Code      string       `json:"code" db:"code"`
	Address   *string      `json:"address,omitempty" db:"address"`
	Phone     *string      `json:"phone,omitempty" db:"phone"`
	Email     *string      `json:"email,omitempty" db:"email"`
	IsActive  bool         `json:"is_active" db:"is_active"`
	Status    SchoolStatus `json:"status" db:"-"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
	CreatedBy *string      `json:"created_by,omitempty" db:"created_by"`
	UpdatedBy *string      `json:"updated_by,omitempty" db:"updated_by"`
}

// CreateSchoolRequest represents the request to create a new school
type CreateSchoolRequest struct {
	Name    string `json:"name" binding:"required,min=2"`
	Code    string `json:"code" binding:"required,min=2"`
	Address string `json:"address,omitempty"`
	Phone   string `json:"phone,omitempty"`
	Email   string `json:"email,omitempty"`
}

// UpdateSchoolRequest represents the request to update a school
type UpdateSchoolRequest struct {
	Name    *string `json:"name,omitempty"`
	Address *string `json:"address,omitempty"`
	Phone   *string `json:"phone,omitempty"`
	Email   *string `json:"email,omitempty"`
}

// UpdateSchoolStatusRequest represents the request to update school status
type UpdateSchoolStatusRequest struct {
	Status SchoolStatus `json:"status" binding:"required,oneof=ACTIVE INACTIVE"`
}

// ToSchoolStatus converts bool to SchoolStatus
func (s *School) ToSchoolStatus() SchoolStatus {
	if s.IsActive {
		return SchoolStatusActive
	}
	return SchoolStatusInactive
}

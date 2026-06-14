package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// GraduateProfileDimension represents a dimension of the Profil Lulusan (Graduate Profile)
// Maps to the 6 dimensions of Profil Lulusan in Kurikulum Merdeka:
// 1. Beriman, bertakwa kepada Tuhan YME, dan berakhlak mulia
// 2. Berkebinekaan global
// 3. Gotong royong
// 4. Mandiri
// 5. Bernalar kritis
// 6. Kreatif
type GraduateProfileDimension struct {
	ID             string    `json:"id" db:"id"`
	Code           string    `json:"code" db:"code"`
	Name           string    `json:"name" db:"name"`
	Description    *string   `json:"description,omitempty" db:"description"`
	SequenceNumber int       `json:"sequence_number" db:"sequence_number"`
	IsActive       bool      `json:"is_active" db:"is_active"`
	CreatedBy      string    `json:"created_by" db:"created_by"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
	UpdatedBy      *string   `json:"updated_by,omitempty" db:"updated_by"`
}

// NewGraduateProfileDimension creates a new GraduateProfileDimension entity
func NewGraduateProfileDimension(code, name string, sequenceNumber int, createdBy string) (*GraduateProfileDimension, error) {
	if code == "" {
		return nil, errors.New("graduate profile dimension code is required")
	}
	if name == "" {
		return nil, errors.New("graduate profile dimension name is required")
	}
	if sequenceNumber < 1 {
		return nil, errors.New("sequence number must be positive")
	}

	return &GraduateProfileDimension{
		ID:             uuid.New().String(),
		Code:           code,
		Name:           name,
		SequenceNumber: sequenceNumber,
		IsActive:       true,
		CreatedBy:      createdBy,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}, nil
}

// Validate performs business rule validation on the GraduateProfileDimension
func (gpd *GraduateProfileDimension) Validate() error {
	if gpd.Code == "" {
		return errors.New("graduate profile dimension code is required")
	}
	if gpd.Name == "" {
		return errors.New("graduate profile dimension name is required")
	}
	if len(gpd.Code) > 20 {
		return errors.New("code must be less than 20 characters")
	}
	if len(gpd.Name) > 100 {
		return errors.New("name must be less than 100 characters")
	}
	if gpd.Description != nil && len(*gpd.Description) > 1000 {
		return errors.New("description must be less than 1000 characters")
	}
	if gpd.SequenceNumber < 1 || gpd.SequenceNumber > 6 {
		return errors.New("sequence number must be between 1 and 6")
	}
	return nil
}

// Activate marks the graduate profile dimension as active
func (gpd *GraduateProfileDimension) Activate() {
	gpd.IsActive = true
	gpd.UpdatedAt = time.Now()
}

// Deactivate marks the graduate profile dimension as inactive
func (gpd *GraduateProfileDimension) Deactivate() {
	gpd.IsActive = false
	gpd.UpdatedAt = time.Now()
}

// IsActiveStatus checks if the graduate profile dimension is active
func (gpd *GraduateProfileDimension) IsActiveStatus() bool {
	return gpd.IsActive
}

// CreateGraduateProfileDimensionRequest represents the request to create a graduate profile dimension
type CreateGraduateProfileDimensionRequest struct {
	Code           string  `json:"code" binding:"required,min=1,max=20"`
	Name           string  `json:"name" binding:"required,min=1,max=100"`
	Description    *string `json:"description,omitempty" binding:"omitempty,max=1000"`
	SequenceNumber int     `json:"sequence_number" binding:"required,min=1,max=6"`
}

// Validate performs validation on the request
func (r *CreateGraduateProfileDimensionRequest) Validate() error {
	if r.Code == "" {
		return errors.New("code is required")
	}
	if len(r.Code) > 20 {
		return errors.New("code must be less than 20 characters")
	}
	if r.Name == "" {
		return errors.New("name is required")
	}
	if len(r.Name) > 100 {
		return errors.New("name must be less than 100 characters")
	}
	if r.Description != nil && len(*r.Description) > 1000 {
		return errors.New("description must be less than 1000 characters")
	}
	if r.SequenceNumber < 1 || r.SequenceNumber > 6 {
		return errors.New("sequence_number must be between 1 and 6")
	}
	return nil
}

// UpdateGraduateProfileDimensionRequest represents the request to update a graduate profile dimension
type UpdateGraduateProfileDimensionRequest struct {
	Name           *string `json:"name,omitempty" binding:"omitempty,max=100"`
	Description    *string `json:"description,omitempty" binding:"omitempty,max=1000"`
	SequenceNumber *int    `json:"sequence_number,omitempty" binding:"omitempty,min=1,max=6"`
	IsActive       *bool   `json:"is_active,omitempty"`
}

// Validate performs validation on the update request
func (r *UpdateGraduateProfileDimensionRequest) Validate() error {
	if r.Name != nil {
		if *r.Name == "" {
			return errors.New("name cannot be empty")
		}
		if len(*r.Name) > 100 {
			return errors.New("name must be less than 100 characters")
		}
	}
	if r.Description != nil && len(*r.Description) > 1000 {
		return errors.New("description must be less than 1000 characters")
	}
	if r.SequenceNumber != nil {
		if *r.SequenceNumber < 1 || *r.SequenceNumber > 6 {
			return errors.New("sequence_number must be between 1 and 6")
		}
	}
	return nil
}

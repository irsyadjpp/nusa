package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// SubjectCategory represents a category for grouping curriculum subjects
// Maps to Kurikulum Merdeka subject groupings (Kelompok Mata Pelajaran)
type SubjectCategory struct {
	ID          string    `json:"id" db:"id"`
	Code        string    `json:"code" db:"code"`
	Name        string    `json:"name" db:"name"`
	Description *string   `json:"description,omitempty" db:"description"`
	IsMandatory bool      `json:"is_mandatory" db:"is_mandatory"`
	IsActive    bool      `json:"is_active" db:"is_active"`
	CreatedBy   string    `json:"created_by" db:"created_by"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	UpdatedBy   *string   `json:"updated_by,omitempty" db:"updated_by"`
}

// NewSubjectCategory creates a new SubjectCategory entity
func NewSubjectCategory(code, name string, isMandatory bool, createdBy string) (*SubjectCategory, error) {
	if code == "" {
		return nil, errors.New("subject category code is required")
	}
	if name == "" {
		return nil, errors.New("subject category name is required")
	}

	return &SubjectCategory{
		ID:          uuid.New().String(),
		Code:        code,
		Name:        name,
		IsMandatory: isMandatory,
		IsActive:    true,
		CreatedBy:   createdBy,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

// Validate performs business rule validation on the SubjectCategory
func (sc *SubjectCategory) Validate() error {
	if sc.Code == "" {
		return errors.New("subject category code is required")
	}
	if sc.Name == "" {
		return errors.New("subject category name is required")
	}
	if len(sc.Code) > 20 {
		return errors.New("code must be less than 20 characters")
	}
	if len(sc.Name) > 100 {
		return errors.New("name must be less than 100 characters")
	}
	if sc.Description != nil && len(*sc.Description) > 500 {
		return errors.New("description must be less than 500 characters")
	}
	return nil
}

// Activate marks the subject category as active
func (sc *SubjectCategory) Activate() {
	sc.IsActive = true
	sc.UpdatedAt = time.Now()
}

// Deactivate marks the subject category as inactive
func (sc *SubjectCategory) Deactivate() {
	sc.IsActive = false
	sc.UpdatedAt = time.Now()
}

// IsActive checks if the subject category is active
func (sc *SubjectCategory) IsActiveStatus() bool {
	return sc.IsActive
}

// CreateSubjectCategoryRequest represents the request to create a subject category
type CreateSubjectCategoryRequest struct {
	Code        string  `json:"code" binding:"required,min=1,max=20"`
	Name        string  `json:"name" binding:"required,min=1,max=100"`
	Description *string `json:"description,omitempty" binding:"omitempty,max=500"`
	IsMandatory bool    `json:"is_mandatory"`
}

// Validate performs validation on the request
func (r *CreateSubjectCategoryRequest) Validate() error {
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
	if r.Description != nil && len(*r.Description) > 500 {
		return errors.New("description must be less than 500 characters")
	}
	return nil
}

// UpdateSubjectCategoryRequest represents the request to update a subject category
type UpdateSubjectCategoryRequest struct {
	Name        *string `json:"name,omitempty" binding:"omitempty,max=100"`
	Description *string `json:"description,omitempty" binding:"omitempty,max=500"`
	IsMandatory *bool   `json:"is_mandatory,omitempty"`
	IsActive    *bool   `json:"is_active,omitempty"`
}

// Validate performs validation on the update request
func (r *UpdateSubjectCategoryRequest) Validate() error {
	if r.Name != nil {
		if *r.Name == "" {
			return errors.New("name cannot be empty")
		}
		if len(*r.Name) > 100 {
			return errors.New("name must be less than 100 characters")
		}
	}
	if r.Description != nil && len(*r.Description) > 500 {
		return errors.New("description must be less than 500 characters")
	}
	return nil
}

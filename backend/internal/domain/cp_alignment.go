package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// CPAlignment represents the alignment between a Curriculum Subject (CP) and a Graduate Profile Dimension
// This maps which CPs support which dimensions of the Profil Lulusan
type CPAlignment struct {
	ID                         string    `json:"id" db:"id"`
	CurriculumSubjectID        string    `json:"curriculum_subject_id" db:"curriculum_subject_id"`
	GraduateProfileDimensionID string    `json:"graduate_profile_dimension_id" db:"graduate_profile_dimension_id"`
	AlignmentDescription       *string   `json:"alignment_description,omitempty" db:"alignment_description"`
	IsActive                   bool      `json:"is_active" db:"is_active"`
	CreatedBy                  string    `json:"created_by" db:"created_by"`
	CreatedAt                  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt                  time.Time `json:"updated_at" db:"updated_at"`
	UpdatedBy                  *string   `json:"updated_by,omitempty" db:"updated_by"`
}

// NewCPAlignment creates a new CPAlignment entity
func NewCPAlignment(curriculumSubjectID, graduateProfileDimensionID string, createdBy string) (*CPAlignment, error) {
	if curriculumSubjectID == "" {
		return nil, errors.New("curriculum subject ID is required")
	}
	if graduateProfileDimensionID == "" {
		return nil, errors.New("graduate profile dimension ID is required")
	}

	return &CPAlignment{
		ID:                         uuid.New().String(),
		CurriculumSubjectID:        curriculumSubjectID,
		GraduateProfileDimensionID: graduateProfileDimensionID,
		IsActive:                   true,
		CreatedBy:                  createdBy,
		CreatedAt:                  time.Now(),
		UpdatedAt:                  time.Now(),
	}, nil
}

// Validate performs business rule validation on the CPAlignment
func (cpa *CPAlignment) Validate() error {
	if cpa.CurriculumSubjectID == "" {
		return errors.New("curriculum subject ID is required")
	}
	if cpa.GraduateProfileDimensionID == "" {
		return errors.New("graduate profile dimension ID is required")
	}
	if cpa.AlignmentDescription != nil && len(*cpa.AlignmentDescription) > 500 {
		return errors.New("alignment description must be less than 500 characters")
	}
	return nil
}

// Activate marks the CP alignment as active
func (cpa *CPAlignment) Activate() {
	cpa.IsActive = true
	cpa.UpdatedAt = time.Now()
}

// Deactivate marks the CP alignment as inactive
func (cpa *CPAlignment) Deactivate() {
	cpa.IsActive = false
	cpa.UpdatedAt = time.Now()
}

// IsActiveStatus checks if the CP alignment is active
func (cpa *CPAlignment) IsActiveStatus() bool {
	return cpa.IsActive
}

// CreateCPAlignmentRequest represents the request to create a CP alignment
type CreateCPAlignmentRequest struct {
	CurriculumSubjectID        string  `json:"curriculum_subject_id" binding:"required"`
	GraduateProfileDimensionID string  `json:"graduate_profile_dimension_id" binding:"required"`
	AlignmentDescription       *string `json:"alignment_description,omitempty" binding:"omitempty,max=500"`
}

// Validate performs validation on the request
func (r *CreateCPAlignmentRequest) Validate() error {
	if r.CurriculumSubjectID == "" {
		return errors.New("curriculum_subject_id is required")
	}
	if r.GraduateProfileDimensionID == "" {
		return errors.New("graduate_profile_dimension_id is required")
	}
	if r.AlignmentDescription != nil && len(*r.AlignmentDescription) > 500 {
		return errors.New("alignment_description must be less than 500 characters")
	}
	return nil
}

// UpdateCPAlignmentRequest represents the request to update a CP alignment
type UpdateCPAlignmentRequest struct {
	AlignmentDescription *string `json:"alignment_description,omitempty" binding:"omitempty,max=500"`
	IsActive             *bool   `json:"is_active,omitempty"`
}

// Validate performs validation on the update request
func (r *UpdateCPAlignmentRequest) Validate() error {
	if r.AlignmentDescription != nil && len(*r.AlignmentDescription) > 500 {
		return errors.New("alignment_description must be less than 500 characters")
	}
	return nil
}

// CPAlignmentReport represents a report showing CP coverage across all graduate profile dimensions
// This is used for the BR-004 business rule: Minimum CP coverage percentage
type CPAlignmentReport struct {
	GraduateProfileDimensionID   string  `json:"graduate_profile_dimension_id"`
	GraduateProfileDimensionName string  `json:"graduate_profile_dimension_name"`
	TotalCPCount                 int     `json:"total_cp_count"`
	AlignedCPCount               int     `json:"aligned_cp_count"`
	CoveragePercentage           float64 `json:"coverage_percentage"`
	MeetsThreshold               bool    `json:"meets_threshold"`
}

// CreateCPAlignmentBulkRequest represents the request to create multiple CP alignments
type CreateCPAlignmentBulkRequest struct {
	CurriculumSubjectID  string   `json:"curriculum_subject_id" binding:"required"`
	AlignmentIDs         []string `json:"alignment_ids" binding:"required"`
	AlignmentDescription *string  `json:"alignment_description,omitempty" binding:"omitempty,max=500"`
}

// Validate performs validation on the bulk request
func (r *CreateCPAlignmentBulkRequest) Validate() error {
	if r.CurriculumSubjectID == "" {
		return errors.New("curriculum_subject_id is required")
	}
	if len(r.AlignmentIDs) == 0 {
		return errors.New("at least one alignment_id is required")
	}
	if len(r.AlignmentIDs) > 6 {
		return errors.New("maximum 6 alignments allowed (one per dimension)")
	}
	if r.AlignmentDescription != nil && len(*r.AlignmentDescription) > 500 {
		return errors.New("alignment_description must be less than 500 characters")
	}
	return nil
}

package dto

import (
	"time"
)

// CreateCurriculumSubjectRequest represents the request to create a curriculum subject
type CreateCurriculumSubjectRequest struct {
	Code        string  `json:"code" binding:"required,min=2"`
	Name        string  `json:"name" binding:"required,min=2"`
	Description *string `json:"description,omitempty"`
}

// UpdateCurriculumSubjectRequest represents the request to update a curriculum subject
type UpdateCurriculumSubjectRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	IsActive    *bool   `json:"is_active,omitempty"`
}

// CreateCurriculumPhaseRequest represents the request to create a curriculum phase
type CreateCurriculumPhaseRequest struct {
	Code            string  `json:"code" binding:"required,min=2"`
	Name            string  `json:"name" binding:"required,min=2"`
	Description     *string `json:"description,omitempty"`
	GradeLevelStart *int    `json:"grade_level_start,omitempty"`
	GradeLevelEnd   *int    `json:"grade_level_end,omitempty"`
}

// UpdateCurriculumPhaseRequest represents the request to update a curriculum phase
type UpdateCurriculumPhaseRequest struct {
	Name            *string `json:"name,omitempty"`
	Description     *string `json:"description,omitempty"`
	GradeLevelStart *int    `json:"grade_level_start,omitempty"`
	GradeLevelEnd   *int    `json:"grade_level_end,omitempty"`
	IsActive        *bool   `json:"is_active,omitempty"`
}

// CreateCurriculumElementRequest represents the request to create a curriculum element
type CreateCurriculumElementRequest struct {
	SubjectID   string  `json:"subject_id" binding:"required"`
	PhaseID     string  `json:"phase_id" binding:"required"`
	Code        string  `json:"code" binding:"required,min=2"`
	Name        string  `json:"name" binding:"required,min=2"`
	Description *string `json:"description,omitempty"`
}

// UpdateCurriculumElementRequest represents the request to update a curriculum element
type UpdateCurriculumElementRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	IsActive    *bool   `json:"is_active,omitempty"`
}

// CreateCurriculumSubelementRequest represents the request to create a curriculum subelement
type CreateCurriculumSubelementRequest struct {
	ElementID   string  `json:"element_id" binding:"required"`
	Code        string  `json:"code" binding:"required,min=2"`
	Name        string  `json:"name" binding:"required,min=2"`
	Description *string `json:"description,omitempty"`
}

// UpdateCurriculumSubelementRequest represents the request to update a curriculum subelement
type UpdateCurriculumSubelementRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	IsActive    *bool   `json:"is_active,omitempty"`
}

// CreateCPRequest represents the request to create a CP (Capaian Pembelajaran)
type CreateCPRequest struct {
	SubjectID           string      `json:"subject_id" binding:"required"`
	PhaseID             string      `json:"phase_id" binding:"required"`
	ElementID           string      `json:"element_id" binding:"required"`
	SubelementID        string      `json:"subelement_id" binding:"required"`
	Code                string      `json:"code" binding:"required"`
	Description         string      `json:"description" binding:"required"`
	CompetencyCode      *string     `json:"competency_code,omitempty"`
	LearningObjectives  interface{} `json:"learning_objectives" binding:"required"`
	CompetencyStandards interface{} `json:"competency_standards" binding:"required"`
	TimeAllocationHours int         `json:"time_allocation_hours" binding:"required,min=1"`
	HoursPerWeek        int         `json:"hours_per_week" binding:"required,min=1,max=40"`
	Version             string      `json:"version" binding:"required"`
}

// UpdateCPRequest represents the request to update a CP
type UpdateCPRequest struct {
	Description         *string     `json:"description,omitempty"`
	CompetencyCode      *string     `json:"competency_code,omitempty"`
	LearningObjectives  interface{} `json:"learning_objectives,omitempty"`
	CompetencyStandards interface{} `json:"competency_standards,omitempty"`
	TimeAllocationHours *int        `json:"time_allocation_hours,omitempty"`
	HoursPerWeek        *int        `json:"hours_per_week,omitempty"`
	Version             *string     `json:"version,omitempty"`
	IsActive            *bool       `json:"is_active,omitempty"`
}

// ImportCPRequest represents the request to import CP data
type ImportCPRequest struct {
	CPs []CreateCPRequest `json:"cps" binding:"required,min=1"`
}

// CPResponse represents the response for CP with related data
type CPResponse struct {
	ID                  string      `json:"id"`
	SubjectID           string      `json:"subject_id"`
	SubjectCode         string      `json:"subject_code"`
	SubjectName         string      `json:"subject_name"`
	PhaseID             string      `json:"phase_id"`
	PhaseCode           string      `json:"phase_code"`
	PhaseName           string      `json:"phase_name"`
	ElementID           string      `json:"element_id"`
	ElementCode         string      `json:"element_code"`
	ElementName         string      `json:"element_name"`
	SubelementID        string      `json:"subelement_id"`
	SubelementCode      string      `json:"subelement_code"`
	SubelementName      string      `json:"subelement_name"`
	Code                string      `json:"code"`
	Description         string      `json:"description"`
	CompetencyCode      *string     `json:"competency_code,omitempty"`
	LearningObjectives  interface{} `json:"learning_objectives"`
	CompetencyStandards interface{} `json:"competency_standards"`
	TimeAllocationHours int         `json:"time_allocation_hours"`
	HoursPerWeek        int         `json:"hours_per_week"`
	Version             string      `json:"version"`
	IsActive            bool        `json:"is_active"`
	ImportedAt          time.Time   `json:"imported_at"`
	ImportedBy          *string     `json:"imported_by,omitempty"`
	ImportedByName      *string     `json:"imported_by_name,omitempty"`
}

package domain

import "time"

// CurriculumSubject represents a curriculum subject (Mata Pelajaran)
type CurriculumSubject struct {
	ID          string    `json:"id" db:"id"`
	Code        string    `json:"code" db:"code"`
	Name        string    `json:"name" db:"name"`
	Description *string   `json:"description,omitempty" db:"description"`
	IsActive    bool      `json:"is_active" db:"is_active"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// CurriculumPhase represents a curriculum phase (Fase)
type CurriculumPhase struct {
	ID              string    `json:"id" db:"id"`
	Code            string    `json:"code" db:"code"`
	Name            string    `json:"name" db:"name"`
	Description     *string   `json:"description,omitempty" db:"description"`
	GradeLevelStart *int      `json:"grade_level_start,omitempty" db:"grade_level_start"`
	GradeLevelEnd   *int      `json:"grade_level_end,omitempty" db:"grade_level_end"`
	IsActive        bool      `json:"is_active" db:"is_active"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

// CurriculumElement represents a curriculum element (Elemen)
type CurriculumElement struct {
	ID          string    `json:"id" db:"id"`
	SubjectID   string    `json:"subject_id" db:"subject_id"`
	PhaseID     string    `json:"phase_id" db:"phase_id"`
	Code        string    `json:"code" db:"code"`
	Name        string    `json:"name" db:"name"`
	Description *string   `json:"description,omitempty" db:"description"`
	IsActive    bool      `json:"is_active" db:"is_active"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// CurriculumSubelement represents a curriculum subelement (Subelemen)
type CurriculumSubelement struct {
	ID          string    `json:"id" db:"id"`
	ElementID   string    `json:"element_id" db:"element_id"`
	Code        string    `json:"code" db:"code"`
	Name        string    `json:"name" db:"name"`
	Description *string   `json:"description,omitempty" db:"description"`
	IsActive    bool      `json:"is_active" db:"is_active"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// WorkflowStatus represents the workflow status for artifacts
type WorkflowStatus string

const (
	WorkflowStatusDraft       WorkflowStatus = "DRAFT"
	WorkflowStatusUnderReview WorkflowStatus = "UNDER_REVIEW"
	WorkflowStatusApproved    WorkflowStatus = "APPROVED"
	WorkflowStatusRejected    WorkflowStatus = "REJECTED"
	WorkflowStatusArchived    WorkflowStatus = "ARCHIVED"
)

// GenerationSource represents the source of generation
type GenerationSource string

const (
	GenerationSourceAIGenerated GenerationSource = "AI_GENERATED"
	GenerationSourceManual      GenerationSource = "MANUAL"
)

// CP represents a Curriculum Plan (Capaian Pembelajaran)
type CP struct {
	ID                  string      `json:"id" db:"id"`
	SubjectID           string      `json:"subject_id" db:"subject_id"`
	PhaseID             string      `json:"phase_id" db:"phase_id"`
	ElementID           string      `json:"element_id" db:"element_id"`
	SubelementID        string      `json:"subelement_id" db:"subelement_id"`
	Code                string      `json:"code" db:"code"`
	Description         string      `json:"description" db:"description"`
	CompetencyCode      *string     `json:"competency_code,omitempty" db:"competency_code"`
	LearningObjectives  interface{} `json:"learning_objectives" db:"learning_objectives"`
	CompetencyStandards interface{} `json:"competency_standards" db:"competency_standards"`
	TimeAllocationHours int         `json:"time_allocation_hours" db:"time_allocation_hours"`
	HoursPerWeek        int         `json:"hours_per_week" db:"hours_per_week"`
	Version             string      `json:"version" db:"version"`
	IsActive            bool        `json:"is_active" db:"is_active"`
	ImportedAt          time.Time   `json:"imported_at" db:"imported_at"`
	ImportedBy          *string     `json:"imported_by,omitempty" db:"imported_by"`
}

package domain

import "time"

// CurriculumSubject represents a curriculum subject (Mata Pelajaran)
type CurriculumSubject struct {
	ID          string    `json:"id" db:"id"`
	Code        string    `json:"code" db:"code"`
	Name        string    `json:"name" db:"name"`
	NameEN      *string   `json:"name_en,omitempty" db:"name_en"`
	Description *string   `json:"description,omitempty" db:"description"`
	IsActive    bool      `json:"is_active" db:"is_active"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// CreateCurriculumSubjectRequest represents the request to create a curriculum subject
type CreateCurriculumSubjectRequest struct {
	Code        string  `json:"code" binding:"required,min=2"`
	Name        string  `json:"name" binding:"required,min=2"`
	NameEN      *string `json:"name_en,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateCurriculumSubjectRequest represents the request to update a curriculum subject
type UpdateCurriculumSubjectRequest struct {
	Name        *string `json:"name,omitempty"`
	NameEN      *string `json:"name_en,omitempty"`
	Description *string `json:"description,omitempty"`
	IsActive    *bool   `json:"is_active,omitempty"`
}

// CurriculumPhase represents a curriculum phase (Fase)
type CurriculumPhase struct {
	ID              string    `json:"id" db:"id"`
	Code            string    `json:"code" db:"code"`
	Name            string    `json:"name" db:"name"`
	NameEN          *string   `json:"name_en,omitempty" db:"name_en"`
	Description     *string   `json:"description,omitempty" db:"description"`
	GradeLevelStart *int      `json:"grade_level_start,omitempty" db:"grade_level_start"`
	GradeLevelEnd   *int      `json:"grade_level_end,omitempty" db:"grade_level_end"`
	IsActive        bool      `json:"is_active" db:"is_active"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

// CreateCurriculumPhaseRequest represents the request to create a curriculum phase
type CreateCurriculumPhaseRequest struct {
	Code            string  `json:"code" binding:"required,min=2"`
	Name            string  `json:"name" binding:"required,min=2"`
	NameEN          *string `json:"name_en,omitempty"`
	Description     *string `json:"description,omitempty"`
	GradeLevelStart *int    `json:"grade_level_start,omitempty"`
	GradeLevelEnd   *int    `json:"grade_level_end,omitempty"`
}

// UpdateCurriculumPhaseRequest represents the request to update a curriculum phase
type UpdateCurriculumPhaseRequest struct {
	Name            *string `json:"name,omitempty"`
	NameEN          *string `json:"name_en,omitempty"`
	Description     *string `json:"description,omitempty"`
	GradeLevelStart *int    `json:"grade_level_start,omitempty"`
	GradeLevelEnd   *int    `json:"grade_level_end,omitempty"`
	IsActive        *bool   `json:"is_active,omitempty"`
}

// CurriculumElement represents a curriculum element (Elemen)
type CurriculumElement struct {
	ID          string    `json:"id" db:"id"`
	SubjectID   string    `json:"subject_id" db:"subject_id"`
	PhaseID     string    `json:"phase_id" db:"phase_id"`
	Code        string    `json:"code" db:"code"`
	Name        string    `json:"name" db:"name"`
	NameEN      *string   `json:"name_en,omitempty" db:"name_en"`
	Description *string   `json:"description,omitempty" db:"description"`
	IsActive    bool      `json:"is_active" db:"is_active"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// CreateCurriculumElementRequest represents the request to create a curriculum element
type CreateCurriculumElementRequest struct {
	SubjectID   string  `json:"subject_id" binding:"required"`
	PhaseID     string  `json:"phase_id" binding:"required"`
	Code        string  `json:"code" binding:"required,min=2"`
	Name        string  `json:"name" binding:"required,min=2"`
	NameEN      *string `json:"name_en,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateCurriculumElementRequest represents the request to update a curriculum element
type UpdateCurriculumElementRequest struct {
	Name        *string `json:"name,omitempty"`
	NameEN      *string `json:"name_en,omitempty"`
	Description *string `json:"description,omitempty"`
	IsActive    *bool   `json:"is_active,omitempty"`
}

// CurriculumSubelement represents a curriculum subelement (Subelemen)
type CurriculumSubelement struct {
	ID          string    `json:"id" db:"id"`
	ElementID   string    `json:"element_id" db:"element_id"`
	Code        string    `json:"code" db:"code"`
	Name        string    `json:"name" db:"name"`
	NameEN      *string   `json:"name_en,omitempty" db:"name_en"`
	Description *string   `json:"description,omitempty" db:"description"`
	IsActive    bool      `json:"is_active" db:"is_active"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// CreateCurriculumSubelementRequest represents the request to create a curriculum subelement
type CreateCurriculumSubelementRequest struct {
	ElementID   string  `json:"element_id" binding:"required"`
	Code        string  `json:"code" binding:"required,min=2"`
	Name        string  `json:"name" binding:"required,min=2"`
	NameEN      *string `json:"name_en,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateCurriculumSubelementRequest represents the request to update a curriculum subelement
type UpdateCurriculumSubelementRequest struct {
	Name        *string `json:"name,omitempty"`
	NameEN      *string `json:"name_en,omitempty"`
	Description *string `json:"description,omitempty"`
	IsActive    *bool   `json:"is_active,omitempty"`
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
	ID                   string       `json:"id" db:"id"`
	SubjectID            string       `json:"subject_id" db:"subject_id"`
	PhaseID              string       `json:"phase_id" db:"phase_id"`
	ElementID            string       `json:"element_id" db:"element_id"`
	SubelementID         string       `json:"subelement_id" db:"subelement_id"`
	Code                 string       `json:"code" db:"code"`
	Description          string       `json:"description" db:"description"`
	CompetencyCode       *string      `json:"competency_code,omitempty" db:"competency_code"`
	LearningObjectives   interface{}  `json:"learning_objectives" db:"learning_objectives"`
	CompetencyStandards  interface{}  `json:"competency_standards" db:"competency_standards"`
	TimeAllocationHours  int          `json:"time_allocation_hours" db:"time_allocation_hours"`
	HoursPerWeek         int          `json:"hours_per_week" db:"hours_per_week"`
	Version              string       `json:"version" db:"version"`
	IsActive             bool         `json:"is_active" db:"is_active"`
	ImportedAt           time.Time    `json:"imported_at" db:"imported_at"`
	ImportedBy           *string      `json:"imported_by,omitempty" db:"imported_by"`
}

// CPResponse represents the response for CP with related data
type CPResponse struct {
	ID                  string       `json:"id"`
	SubjectID           string       `json:"subject_id"`
	SubjectCode         string       `json:"subject_code"`
	SubjectName         string       `json:"subject_name"`
	PhaseID             string       `json:"phase_id"`
	PhaseCode           string       `json:"phase_code"`
	PhaseName           string       `json:"phase_name"`
	ElementID           string       `json:"element_id"`
	ElementCode         string       `json:"element_code"`
	ElementName         string       `json:"element_name"`
	SubelementID        string       `json:"subelement_id"`
	SubelementCode      string       `json:"subelement_code"`
	SubelementName      string       `json:"subelement_name"`
	Code                string       `json:"code"`
	Description         string       `json:"description"`
	CompetencyCode      *string      `json:"competency_code,omitempty"`
	LearningObjectives  interface{}  `json:"learning_objectives"`
	CompetencyStandards interface{}  `json:"competency_standards"`
	TimeAllocationHours int          `json:"time_allocation_hours"`
	HoursPerWeek        int          `json:"hours_per_week"`
	Version             string       `json:"version"`
	IsActive            bool         `json:"is_active"`
	ImportedAt          time.Time    `json:"imported_at"`
	ImportedBy          *string      `json:"imported_by,omitempty"`
	ImportedByName      *string      `json:"imported_by_name,omitempty"`
}

// CreateCPRequest represents the request to create a CP
type CreateCPRequest struct {
	SubjectID           string       `json:"subject_id" binding:"required"`
	PhaseID             string       `json:"phase_id" binding:"required"`
	ElementID           string       `json:"element_id" binding:"required"`
	SubelementID        string       `json:"subelement_id" binding:"required"`
	Code                string       `json:"code" binding:"required"`
	Description         string       `json:"description" binding:"required"`
	CompetencyCode      *string      `json:"competency_code,omitempty"`
	LearningObjectives   interface{} `json:"learning_objectives" binding:"required"`
	CompetencyStandards interface{} `json:"competency_standards" binding:"required"`
	TimeAllocationHours int         `json:"time_allocation_hours" binding:"required,min=1"`
	HoursPerWeek        int         `json:"hours_per_week" binding:"required,min=1,max=40"`
	Version             string      `json:"version" binding:"required"`
}

// UpdateCPRequest represents the request to update a CP
type UpdateCPRequest struct {
	Description         *string     `json:"description,omitempty"`
	CompetencyCode      *string     `json:"competency_code,omitempty"`
	LearningObjectives   interface{} `json:"learning_objectives,omitempty"`
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

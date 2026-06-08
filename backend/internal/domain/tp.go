package domain

import (
	"encoding/json"
	"errors"
	"time"
)

// KKTPCriteria represents the embedded Success Criteria (KKTP) Value Object
// This is NOT an aggregate - it's embedded inside TP
type KKTPCriteria struct {
	MasteryThresholds     MasteryThresholds     `json:"mastery_thresholds"`
	PerformanceIndicators PerformanceIndicators `json:"performance_indicators"`
	MinimumRequirements   MinimumRequirements   `json:"minimum_requirements"`
}

// MasteryThresholds defines the mastery levels for assessment
type MasteryThresholds struct {
	ExcellentThreshold  float64 `json:"excellent_threshold"`
	ProficientThreshold float64 `json:"proficient_threshold"`
	DevelopingThreshold float64 `json:"developing_threshold"`
	BeginningThreshold  float64 `json:"beginning_threshold"`
}

// PerformanceIndicators defines observable behaviors
type PerformanceIndicators struct {
	Cognitive   []string `json:"cognitive"`
	Psychomotor []string `json:"psychomotor"`
	Affective   []string `json:"affective"`
}

// MinimumRequirements defines minimum acceptable standards
type MinimumRequirements struct {
	CoreCompetencies []string `json:"core_competencies"`
	EssentialSkills  []string `json:"essential_skills"`
	RequiredEvidence []string `json:"required_evidence"`
}

// Validate validates the KKTPCriteria
func (k *KKTPCriteria) Validate() error {
	if k.MasteryThresholds.ExcellentThreshold <= k.MasteryThresholds.ProficientThreshold {
		return errors.New("excellent threshold must be greater than proficient threshold")
	}
	if k.MasteryThresholds.ProficientThreshold <= k.MasteryThresholds.DevelopingThreshold {
		return errors.New("proficient threshold must be greater than developing threshold")
	}
	if k.MasteryThresholds.DevelopingThreshold <= k.MasteryThresholds.BeginningThreshold {
		return errors.New("developing threshold must be greater than beginning threshold")
	}
	if k.MasteryThresholds.BeginningThreshold < 0 {
		return errors.New("beginning threshold must be non-negative")
	}
	if len(k.MinimumRequirements.CoreCompetencies) == 0 {
		return errors.New("at least one core competency is required")
	}
	return nil
}

// ToJSON converts KKTPCriteria to JSON
func (k *KKTPCriteria) ToJSON() (interface{}, error) {
	data, err := json.Marshal(k)
	if err != nil {
		return nil, err
	}
	var result interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// FromJSON creates KKTPCriteria from JSON
func FromJSONToKKTPCriteria(data interface{}) (*KKTPCriteria, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	var criteria KKTPCriteria
	if err := json.Unmarshal(jsonData, &criteria); err != nil {
		return nil, err
	}
	if err := criteria.Validate(); err != nil {
		return nil, err
	}
	return &criteria, nil
}

// TPSet represents a Teaching Plan Set
type TPSet struct {
	ID               string           `json:"id" db:"id"`
	CPID             string           `json:"cp_id" db:"cp_id"`
	VersionNo        int              `json:"version_no" db:"version_no"`
	Status           WorkflowStatus   `json:"status" db:"status"`
	GenerationSource GenerationSource `json:"generation_source" db:"generation_source"`
	GenerationReason *string          `json:"generation_reason,omitempty" db:"generation_reason"`
	GeneratedBy      string           `json:"generated_by" db:"generated_by"`
	AIGenerationID   *string          `json:"ai_generation_id,omitempty" db:"ai_generation_id"`
	ApprovedBy       *string          `json:"approved_by,omitempty" db:"approved_by"`
	ApprovedAt       *time.Time       `json:"approved_at,omitempty" db:"approved_at"`
	CreatedAt        time.Time        `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time        `json:"updated_at" db:"updated_at"`
}

// TPSetResponse represents the response for TP Set with related data
type TPSetResponse struct {
	ID               string           `json:"id"`
	CPID             string           `json:"cp_id"`
	CPCode           string           `json:"cp_code"`
	CPText           string           `json:"cp_text"`
	VersionNo        int              `json:"version_no"`
	Status           WorkflowStatus   `json:"status"`
	GenerationSource GenerationSource `json:"generation_source"`
	GenerationReason *string          `json:"generation_reason,omitempty"`
	GeneratedBy      string           `json:"generated_by"`
	GeneratedByName  string           `json:"generated_by_name"`
	AIGenerationID   *string          `json:"ai_generation_id,omitempty"`
	ApprovedBy       *string          `json:"approved_by,omitempty"`
	ApprovedByName   *string          `json:"approved_by_name,omitempty"`
	ApprovedAt       *time.Time       `json:"approved_at,omitempty"`
	CreatedAt        time.Time        `json:"created_at"`
	UpdatedAt        time.Time        `json:"updated_at"`
}

// CreateTPSetRequest represents the request to create a TP Set
type CreateTPSetRequest struct {
	CPID             string           `json:"cp_id" binding:"required"`
	VersionNo        int              `json:"version_no" binding:"required,min=1"`
	GenerationSource GenerationSource `json:"generation_source" binding:"required,oneof=AI_GENERATED MANUAL"`
	GenerationReason *string          `json:"generation_reason,omitempty"`
}

// UpdateTPSetRequest represents the request to update a TP Set
type UpdateTPSetRequest struct {
	Status           *WorkflowStatus `json:"status,omitempty"`
	GenerationReason *string         `json:"generation_reason,omitempty"`
}

// TP represents a Teaching Plan Item
type TP struct {
	ID                 string         `json:"id" db:"id"`
	TPSetID            string         `json:"tp_set_id" db:"tp_set_id"`
	SequenceNumber     int            `json:"sequence_number" db:"sequence_number"`
	CPID               string         `json:"cp_id" db:"cp_id"`
	SubjectID          string         `json:"subject_id" db:"subject_id"`
	PhaseID            string         `json:"phase_id" db:"phase_id"`
	ElementID          string         `json:"element_id" db:"element_id"`
	SubelementID       string         `json:"subelement_id" db:"subelement_id"`
	UserID             string         `json:"user_id" db:"user_id"`
	Status             WorkflowStatus `json:"status" db:"status"`
	Title              *string        `json:"title,omitempty" db:"title"`
	LearningObjectives interface{}    `json:"learning_objectives" db:"learning_objectives"`
	TimeAllocation     interface{}    `json:"time_allocation" db:"time_allocation"`
	Prerequisites      interface{}    `json:"prerequisites,omitempty" db:"prerequisites"`
	EstimatedWeeks     *int           `json:"estimated_weeks,omitempty" db:"estimated_weeks"`
	SuccessCriteria    interface{}    `json:"success_criteria" db:"success_criteria"` // Embedded KKTPCriteria as JSONB
	// Version Tracking
	VersionNo        int       `json:"version_no" db:"version_no"`
	IsCurrentVersion bool      `json:"is_current_version" db:"is_current_version"`
	ParentVersionID  *string   `json:"parent_version_id,omitempty" db:"parent_version_id"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
}

// TPResponse represents the response for TP with related data
type TPResponse struct {
	ID                 string         `json:"id"`
	TPSetID            string         `json:"tp_set_id"`
	SequenceNumber     int            `json:"sequence_number"`
	CPID               string         `json:"cp_id"`
	CPCode             string         `json:"cp_code"`
	CPText             string         `json:"cp_text"`
	SubjectID          string         `json:"subject_id"`
	SubjectCode        string         `json:"subject_code"`
	SubjectName        string         `json:"subject_name"`
	PhaseID            string         `json:"phase_id"`
	PhaseCode          string         `json:"phase_code"`
	PhaseName          string         `json:"phase_name"`
	ElementID          string         `json:"element_id"`
	ElementCode        string         `json:"element_code"`
	ElementName        string         `json:"element_name"`
	SubelementID       string         `json:"subelement_id"`
	SubelementCode     string         `json:"subelement_code"`
	SubelementName     string         `json:"subelement_name"`
	UserID             string         `json:"user_id"`
	UserName           string         `json:"user_name"`
	Status             WorkflowStatus `json:"status"`
	Title              *string        `json:"title,omitempty"`
	LearningObjectives interface{}    `json:"learning_objectives"`
	TimeAllocation     interface{}    `json:"time_allocation"`
	Prerequisites      interface{}    `json:"prerequisites,omitempty"`
	EstimatedWeeks     *int           `json:"estimated_weeks,omitempty"`
	SuccessCriteria    interface{}    `json:"success_criteria"`
	// Version Tracking
	VersionNo        int       `json:"version_no"`
	IsCurrentVersion bool      `json:"is_current_version"`
	ParentVersionID  *string   `json:"parent_version_id,omitempty"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// CreateTPRequest represents the request to create a TP
type CreateTPRequest struct {
	TPSetID            string      `json:"tp_set_id" binding:"required"`
	SequenceNumber     int         `json:"sequence_number" binding:"required,min=1"`
	CPID               string      `json:"cp_id" binding:"required"`
	SubjectID          string      `json:"subject_id" binding:"required"`
	PhaseID            string      `json:"phase_id" binding:"required"`
	ElementID          string      `json:"element_id" binding:"required"`
	SubelementID       string      `json:"subelement_id" binding:"required"`
	UserID             string      `json:"user_id" binding:"required"`
	Title              *string     `json:"title,omitempty"`
	LearningObjectives interface{} `json:"learning_objectives" binding:"required"`
	TimeAllocation     interface{} `json:"time_allocation" binding:"required"`
	Prerequisites      interface{} `json:"prerequisites,omitempty"`
	EstimatedWeeks     *int        `json:"estimated_weeks,omitempty"`
	SuccessCriteria    interface{} `json:"success_criteria"` // Embedded KKTPCriteria
}

// UpdateTPRequest represents the request to update a TP
type UpdateTPRequest struct {
	Title              *string         `json:"title,omitempty"`
	LearningObjectives interface{}     `json:"learning_objectives,omitempty"`
	TimeAllocation     interface{}     `json:"time_allocation,omitempty"`
	Prerequisites      interface{}     `json:"prerequisites,omitempty"`
	EstimatedWeeks     *int            `json:"estimated_weeks,omitempty"`
	SuccessCriteria    interface{}     `json:"success_criteria,omitempty"`
	Status             *WorkflowStatus `json:"status,omitempty"`
}

// GenerateTPRequest represents the request to generate TPs from CP
type GenerateTPRequest struct {
	CPID             string  `json:"cp_id" binding:"required"`
	GenerationReason *string `json:"generation_reason,omitempty"`
}

// ApproveTPSetRequest represents the request to approve a TP Set
type ApproveTPSetRequest struct {
	Reason string `json:"reason" binding:"required,min=5"`
}

// RejectTPSetRequest represents the request to reject a TP Set
type RejectTPSetRequest struct {
	Reason string `json:"reason" binding:"required,min=5"`
}

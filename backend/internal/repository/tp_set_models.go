package repository

import (
	"time"
)

// TPSetDBModel represents the database model for tp_sets table
// This is used for mapping between database rows and domain objects
type TPSetDBModel struct {
	ID               string     `db:"id"`
	CPID             string     `db:"cp_id"`
	VersionNo        int        `db:"version_no"`
	Status           string     `db:"status"`
	GenerationSource string     `db:"generation_source"`
	GenerationReason *string    `db:"generation_reason"`
	GeneratedBy      string     `db:"generated_by"`
	AIGenerationID   *string    `db:"ai_generation_id"`
	ApprovedBy       *string    `db:"approved_by"`
	ApprovedAt       *time.Time `db:"approved_at"`
	CreatedAt        time.Time  `db:"created_at"`
	UpdatedAt        time.Time  `db:"updated_at"`
}

// TPDBModel represents the database model for tp table
type TPDBModel struct {
	ID                 string      `db:"id"`
	TPSetID            string      `db:"tp_set_id"`
	SequenceNumber     int         `db:"sequence_number"`
	CPID               string      `db:"cp_id"`
	SubjectID          string      `db:"subject_id"`
	PhaseID            string      `db:"phase_id"`
	ElementID          string      `db:"element_id"`
	SubelementID       string      `db:"subelement_id"`
	UserID             string      `db:"user_id"`
	Status             string      `db:"status"`
	Title              *string     `db:"title"`
	LearningObjectives interface{} `db:"learning_objectives"`
	TimeAllocation     interface{} `db:"time_allocation"`
	Prerequisites      interface{} `db:"prerequisites"`
	EstimatedWeeks     *int        `db:"estimated_weeks"`
	SuccessCriteria    interface{} `db:"success_criteria"`
	VersionNo          int         `db:"version_no"`
	IsCurrentVersion   bool        `db:"is_current_version"`
	ParentVersionID    *string     `db:"parent_version_id"`
	CreatedAt          time.Time   `db:"created_at"`
	UpdatedAt          time.Time   `db:"updated_at"`
}

package domain

import "time"

// ATPSet represents an Annual Teaching Plan Set
type ATPSet struct {
	ID               string             `json:"id" db:"id"`
	TPSetID          string             `json:"tp_set_id" db:"tp_set_id"`
	VersionNo        int                `json:"version_no" db:"version_no"`
	Status           WorkflowStatus     `json:"status" db:"status"`
	GenerationSource GenerationSource  `json:"generation_source" db:"generation_source"`
	GenerationReason *string            `json:"generation_reason,omitempty" db:"generation_reason"`
	GeneratedBy      string             `json:"generated_by" db:"generated_by"`
	AIGenerationID   *string            `json:"ai_generation_id,omitempty" db:"ai_generation_id"`
	ApprovedBy       *string            `json:"approved_by,omitempty" db:"approved_by"`
	ApprovedAt       *time.Time         `json:"approved_at,omitempty" db:"approved_at"`
	CreatedAt        time.Time          `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at" db:"updated_at"`
}

// ATPSetResponse represents the response for ATP Set with related data
type ATPSetResponse struct {
	ID               string             `json:"id"`
	TPSetID          string             `json:"tp_set_id"`
	TPSetVersionNo   int                `json:"tp_set_version_no"`
	VersionNo        int                `json:"version_no"`
	Status           WorkflowStatus     `json:"status"`
	GenerationSource GenerationSource  `json:"generation_source"`
	GenerationReason *string            `json:"generation_reason,omitempty"`
	GeneratedBy      string             `json:"generated_by"`
	GeneratedByName  string             `json:"generated_by_name"`
	AIGenerationID   *string            `json:"ai_generation_id,omitempty"`
	ApprovedBy       *string            `json:"approved_by,omitempty"`
	ApprovedByName   *string            `json:"approved_by_name,omitempty"`
	ApprovedAt       *time.Time         `json:"approved_at,omitempty"`
	CreatedAt        time.Time          `json:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at"`
}

// CreateATPSetRequest represents the request to create an ATP Set
type CreateATPSetRequest struct {
	TPSetID          string            `json:"tp_set_id" binding:"required"`
	VersionNo        int               `json:"version_no" binding:"required,min=1"`
	GenerationSource GenerationSource `json:"generation_source" binding:"required,oneof=AI_GENERATED MANUAL"`
	GenerationReason *string           `json:"generation_reason,omitempty"`
}

// ATP represents an Annual Teaching Plan Item
type ATP struct {
	ID               string       `json:"id" db:"id"`
	ATPSetID         string       `json:"atp_set_id" db:"atp_set_id"`
	TPID             string       `json:"tp_id" db:"tp_id"`
	UserID           string       `json:"user_id" db:"user_id"`
	Status           WorkflowStatus `json:"status" db:"status"`
	AcademicCalendar interface{} `json:"academic_calendar" db:"academic_calendar"`
	ClassSchedule    interface{} `json:"class_schedule" db:"class_schedule"`
	WeeklySequence   interface{} `json:"weekly_sequence" db:"weekly_sequence"`
	AssessmentSchedule interface{} `json:"assessment_schedule,omitempty" db:"assessment_schedule"`
	CreatedAt        time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time    `json:"updated_at" db:"updated_at"`
}

// ATPResponse represents the response for ATP with related data
type ATPResponse struct {
	ID               string       `json:"id"`
	ATPSetID         string       `json:"atp_set_id"`
	TPID             string       `json:"tp_id"`
	TPTitle          string       `json:"tp_title"`
	UserID           string       `json:"user_id"`
	UserName         string       `json:"user_name"`
	Status           WorkflowStatus `json:"status"`
	AcademicCalendar interface{} `json:"academic_calendar"`
	ClassSchedule    interface{} `json:"class_schedule"`
	WeeklySequence   interface{} `json:"weekly_sequence"`
	AssessmentSchedule interface{} `json:"assessment_schedule,omitempty"`
	CreatedAt        time.Time    `json:"created_at"`
	UpdatedAt        time.Time    `json:"updated_at"`
}

// CreateATPRequest represents the request to create an ATP
type CreateATPRequest struct {
	ATPSetID         string       `json:"atp_set_id" binding:"required"`
	TPID             string       `json:"tp_id" binding:"required"`
	UserID           string       `json:"user_id" binding:"required"`
	AcademicCalendar interface{} `json:"academic_calendar" binding:"required"`
	ClassSchedule    interface{} `json:"class_schedule" binding:"required"`
	WeeklySequence   interface{} `json:"weekly_sequence" binding:"required"`
	AssessmentSchedule interface{} `json:"assessment_schedule,omitempty"`
}

// UpdateATPRequest represents the request to update an ATP
type UpdateATPRequest struct {
	AcademicCalendar interface{} `json:"academic_calendar,omitempty"`
	ClassSchedule    interface{} `json:"class_schedule,omitempty"`
	WeeklySequence   interface{} `json:"weekly_sequence,omitempty"`
	AssessmentSchedule interface{} `json:"assessment_schedule,omitempty"`
	Status           *WorkflowStatus `json:"status,omitempty"`
}

// ModulAjarSet represents a Modul Ajar Set
type ModulAjarSet struct {
	ID               string             `json:"id" db:"id"`
	ATPSetID         string             `json:"atp_set_id" db:"atp_set_id"`
	VersionNo        int                `json:"version_no" db:"version_no"`
	Status           WorkflowStatus     `json:"status" db:"status"`
	GenerationSource GenerationSource  `json:"generation_source" db:"generation_source"`
	GenerationReason *string            `json:"generation_reason,omitempty" db:"generation_reason"`
	GeneratedBy      string             `json:"generated_by" db:"generated_by"`
	AIGenerationID   *string            `json:"ai_generation_id,omitempty" db:"ai_generation_id"`
	ApprovedBy       *string            `json:"approved_by,omitempty" db:"approved_by"`
	ApprovedAt       *time.Time         `json:"approved_at,omitempty" db:"approved_at"`
	CreatedAt        time.Time          `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at" db:"updated_at"`
}

// ModulAjarSetResponse represents the response for Modul Ajar Set with related data
type ModulAjarSetResponse struct {
	ID               string             `json:"id"`
	ATPSetID         string             `json:"atp_set_id"`
	ATPSetVersionNo  int                `json:"atp_set_version_no"`
	VersionNo        int                `json:"version_no"`
	Status           WorkflowStatus     `json:"status"`
	GenerationSource GenerationSource  `json:"generation_source"`
	GenerationReason *string            `json:"generation_reason,omitempty"`
	GeneratedBy      string             `json:"generated_by"`
	GeneratedByName  string             `json:"generated_by_name"`
	AIGenerationID   *string            `json:"ai_generation_id,omitempty"`
	ApprovedBy       *string            `json:"approved_by,omitempty"`
	ApprovedByName   *string            `json:"approved_by_name,omitempty"`
	ApprovedAt       *time.Time         `json:"approved_at,omitempty"`
	CreatedAt        time.Time          `json:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at"`
}

// CreateModulAjarSetRequest represents the request to create a Modul Ajar Set
type CreateModulAjarSetRequest struct {
	ATPSetID         string            `json:"atp_set_id" binding:"required"`
	VersionNo        int               `json:"version_no" binding:"required,min=1"`
	GenerationSource GenerationSource `json:"generation_source" binding:"required,oneof=AI_GENERATED MANUAL"`
	GenerationReason *string           `json:"generation_reason,omitempty"`
}

// ModulAjar represents a Modul Ajar Item
type ModulAjar struct {
	ID                  string       `json:"id" db:"id"`
	ModulAjarSetID      string       `json:"modul_ajar_set_id" db:"modul_ajar_set_id"`
	ATPID               string       `json:"atp_id" db:"atp_id"`
	Week                int          `json:"week" db:"week"`
	Topic               interface{} `json:"topic" db:"topic"`
	Resources           interface{} `json:"resources,omitempty" db:"resources"`
	ClassCharacteristics interface{} `json:"class_characteristics,omitempty" db:"class_characteristics"`
	LearningActivities  interface{} `json:"learning_activities" db:"learning_activities"`
	ResourceRequirements interface{} `json:"resource_requirements,omitempty" db:"resource_requirements"`
	AssessmentMethods   interface{} `json:"assessment_methods,omitempty" db:"assessment_methods"`
	Status              WorkflowStatus `json:"status" db:"status"`
	CreatedAt           time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt           time.Time    `json:"updated_at" db:"updated_at"`
}

// ModulAjarResponse represents the response for Modul Ajar with related data
type ModulAjarResponse struct {
	ID                  string       `json:"id"`
	ModulAjarSetID      string       `json:"modul_ajar_set_id"`
	ATPID               string       `json:"atp_id"`
	ATPWeek             int          `json:"atp_week"`
	Week                int          `json:"week"`
	Topic               interface{} `json:"topic"`
	Resources           interface{} `json:"resources,omitempty"`
	ClassCharacteristics interface{} `json:"class_characteristics,omitempty"`
	LearningActivities  interface{} `json:"learning_activities"`
	ResourceRequirements interface{} `json:"resource_requirements,omitempty"`
	AssessmentMethods   interface{} `json:"assessment_methods,omitempty"`
	Status              WorkflowStatus `json:"status"`
	CreatedAt           time.Time    `json:"created_at"`
	UpdatedAt           time.Time    `json:"updated_at"`
}

// CreateModulAjarRequest represents the request to create a Modul Ajar
type CreateModulAjarRequest struct {
	ModulAjarSetID      string       `json:"modul_ajar_set_id" binding:"required"`
	ATPID               string       `json:"atp_id" binding:"required"`
	Week                int          `json:"week" binding:"required,min=1"`
	Topic               interface{} `json:"topic" binding:"required"`
	Resources           interface{} `json:"resources,omitempty"`
	ClassCharacteristics interface{} `json:"class_characteristics,omitempty"`
	LearningActivities  interface{} `json:"learning_activities" binding:"required"`
	ResourceRequirements interface{} `json:"resource_requirements,omitempty"`
	AssessmentMethods   interface{} `json:"assessment_methods,omitempty"`
}

// UpdateModulAjarRequest represents the request to update a Modul Ajar
type UpdateModulAjarRequest struct {
	Topic               interface{} `json:"topic,omitempty"`
	Resources           interface{} `json:"resources,omitempty"`
	ClassCharacteristics interface{} `json:"class_characteristics,omitempty"`
	LearningActivities  interface{} `json:"learning_activities,omitempty"`
	ResourceRequirements interface{} `json:"resource_requirements,omitempty"`
	AssessmentMethods   interface{} `json:"assessment_methods,omitempty"`
	Status              *WorkflowStatus `json:"status,omitempty"`
}

// GenerateATPRequest represents the request to generate ATP from TP Set
type GenerateATPRequest struct {
	TPSetID          string  `json:"tp_set_id" binding:"required"`
	GenerationReason *string `json:"generation_reason,omitempty"`
}

// GenerateModulAjarRequest represents the request to generate Modul Ajar from ATP Set
type GenerateModulAjarRequest struct {
	ATPSetID         string  `json:"atp_set_id" binding:"required"`
	GenerationReason *string `json:"generation_reason,omitempty"`
}

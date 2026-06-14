package dto

// ==================== ATP Set DTOs ====================

// CreateATPSetRequest represents the request body for creating an ATP Set
type CreateATPSetRequest struct {
	TPSetID          string `json:"tp_set_id" binding:"required"`
	VersionNo        int    `json:"version_no" binding:"required,min=1"`
	GenerationSource string `json:"generation_source" binding:"required,oneof=AI_GENERATED MANUAL"`
	GenerationReason string `json:"generation_reason"`
}

// ATPSetResponse represents the response body for ATP Set operations
type ATPSetResponse struct {
	ID               string  `json:"id"`
	TPSetID          string  `json:"tp_set_id"`
	TPSetVersionNo   int     `json:"tp_set_version_no"`
	VersionNo        int     `json:"version_no"`
	Status           string  `json:"status"`
	GenerationSource string  `json:"generation_source"`
	GenerationReason *string `json:"generation_reason"`
	GeneratedBy      string  `json:"generated_by"`
	GeneratedByName  string  `json:"generated_by_name"`
	AIGenerationID   *string `json:"ai_generation_id"`
	ApprovedBy       *string `json:"approved_by"`
	ApprovedByName   *string `json:"approved_by_name,omitempty"`
	ApprovedAt       *string `json:"approved_at"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
}

// UpdateATPSetRequest represents the request body for updating an ATP Set
type UpdateATPSetRequest struct {
	VersionNo *int    `json:"version_no,omitempty"`
	Status    *string `json:"status,omitempty"`
}

// ==================== ATP DTOs ====================

// CreateATPRequest represents the request body for creating an ATP
type CreateATPRequest struct {
	ATPSetID           string      `json:"atp_set_id" binding:"required"`
	TPID               string      `json:"tp_id" binding:"required"`
	UserID             string      `json:"user_id" binding:"required"`
	AcademicCalendar   interface{} `json:"academic_calendar" binding:"required"`
	ClassSchedule      interface{} `json:"class_schedule" binding:"required"`
	WeeklySequence     interface{} `json:"weekly_sequence" binding:"required"`
	AssessmentSchedule interface{} `json:"assessment_schedule,omitempty"`
}

// UpdateATPRequest represents the request body for updating an ATP
type UpdateATPRequest struct {
	TPID               *string     `json:"tp_id,omitempty"`
	AcademicCalendar   interface{} `json:"academic_calendar,omitempty"`
	ClassSchedule      interface{} `json:"class_schedule,omitempty"`
	WeeklySequence     interface{} `json:"weekly_sequence,omitempty"`
	AssessmentSchedule interface{} `json:"assessment_schedule,omitempty"`
	Status             *string     `json:"status,omitempty"`
}

// ATPResponse represents the response body for ATP operations
type ATPResponse struct {
	ID                 string      `json:"id"`
	ATPSetID           string      `json:"atp_set_id"`
	TPID               string      `json:"tp_id"`
	TPTitle            string      `json:"tp_title"`
	UserID             string      `json:"user_id"`
	UserName           string      `json:"user_name"`
	Status             string      `json:"status"`
	AcademicCalendar   interface{} `json:"academic_calendar"`
	ClassSchedule      interface{} `json:"class_schedule"`
	WeeklySequence     interface{} `json:"weekly_sequence"`
	AssessmentSchedule interface{} `json:"assessment_schedule,omitempty"`
	CreatedAt          string      `json:"created_at"`
	UpdatedAt          string      `json:"updated_at"`
}

// ==================== Modul Ajar Set DTOs ====================

// CreateModulAjarSetRequest represents the request body for creating a Modul Ajar Set
type CreateModulAjarSetRequest struct {
	ATPSetID         string `json:"atp_set_id" binding:"required"`
	VersionNo        int    `json:"version_no" binding:"required,min=1"`
	GenerationSource string `json:"generation_source" binding:"required,oneof=AI_GENERATED MANUAL"`
	GenerationReason string `json:"generation_reason"`
}

// ModulAjarSetResponse represents the response body for Modul Ajar Set operations
type ModulAjarSetResponse struct {
	ID               string  `json:"id"`
	ATPSetID         string  `json:"atp_set_id"`
	ATPSetVersionNo  int     `json:"atp_set_version_no"`
	VersionNo        int     `json:"version_no"`
	Status           string  `json:"status"`
	GenerationSource string  `json:"generation_source"`
	GenerationReason *string `json:"generation_reason"`
	GeneratedBy      string  `json:"generated_by"`
	GeneratedByName  string  `json:"generated_by_name"`
	AIGenerationID   *string `json:"ai_generation_id"`
	ApprovedBy       *string `json:"approved_by"`
	ApprovedByName   *string `json:"approved_by_name,omitempty"`
	ApprovedAt       *string `json:"approved_at"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
}

// UpdateModulAjarSetRequest represents the request body for updating a Modul Ajar Set
type UpdateModulAjarSetRequest struct {
	VersionNo *int    `json:"version_no,omitempty"`
	Status    *string `json:"status,omitempty"`
}

// ==================== Modul Ajar DTOs ====================

// CreateModulAjarRequest represents the request body for creating a Modul Ajar
type CreateModulAjarRequest struct {
	ModulAjarSetID       string      `json:"modul_ajar_set_id" binding:"required"`
	ATPID                string      `json:"atp_id" binding:"required"`
	Week                 int         `json:"week" binding:"required,min=1,max=52"`
	Topic                interface{} `json:"topic" binding:"required"`
	Resources            interface{} `json:"resources,omitempty"`
	ClassCharacteristics interface{} `json:"class_characteristics,omitempty"`
	TeachingMethods      interface{} `json:"teaching_methods,omitempty"`
	LearningActivities   interface{} `json:"learning_activities" binding:"required"`
	AssessmentMethods    interface{} `json:"assessment_methods" binding:"required"`
	Reflection           interface{} `json:"reflection,omitempty"`
	Status               string      `json:"status" binding:"required,oneof=DRAFT PENDING APPROVED REJECTED"`
}

// UpdateModulAjarRequest represents the request body for updating a Modul Ajar
type UpdateModulAjarRequest struct {
	ATPID                *string     `json:"atp_id,omitempty"`
	Week                 *int        `json:"week,omitempty"`
	Topic                interface{} `json:"topic,omitempty"`
	Resources            interface{} `json:"resources,omitempty"`
	ClassCharacteristics interface{} `json:"class_characteristics,omitempty"`
	TeachingMethods      interface{} `json:"teaching_methods,omitempty"`
	LearningActivities   interface{} `json:"learning_activities,omitempty"`
	AssessmentMethods    interface{} `json:"assessment_methods,omitempty"`
	Reflection           interface{} `json:"reflection,omitempty"`
	Status               *string     `json:"status,omitempty"`
}

// ModulAjarResponse represents the response body for Modul Ajar operations
type ModulAjarResponse struct {
	ID                   string      `json:"id"`
	ModulAjarSetID       string      `json:"modul_ajar_set_id"`
	ATPID                string      `json:"atp_id"`
	ATPTitle             string      `json:"atp_title"`
	Week                 int         `json:"week"`
	Topic                interface{} `json:"topic"`
	Resources            interface{} `json:"resources,omitempty"`
	ClassCharacteristics interface{} `json:"class_characteristics,omitempty"`
	TeachingMethods      interface{} `json:"teaching_methods,omitempty"`
	LearningActivities   interface{} `json:"learning_activities"`
	AssessmentMethods    interface{} `json:"assessment_methods"`
	Reflection           interface{} `json:"reflection,omitempty"`
	Status               string      `json:"status"`
	CreatedAt            string      `json:"created_at"`
	UpdatedAt            string      `json:"updated_at"`
}

// ListATPSetsResponse represents the response body for listing ATP Sets
type ListATPSetsResponse struct {
	ATPSets []ATPSetResponse `json:"atp_sets"`
	Total   int              `json:"total"`
}

// ListATPsResponse represents the response body for listing ATPs
type ListATPsResponse struct {
	ATPs     []ATPResponse `json:"atps"`
	Total    int           `json:"total"`
	Page     int           `json:"page"`
	PageSize int           `json:"page_size"`
}

// ListModulAjarSetsResponse represents the response body for listing Modul Ajar Sets
type ListModulAjarSetsResponse struct {
	ModulAjarSets []ModulAjarSetResponse `json:"modul_ajar_sets"`
	Total         int                    `json:"total"`
}

// ListModulAjarsResponse represents the response body for listing Modul Ajars
type ListModulAjarsResponse struct {
	ModulAjars []ModulAjarResponse `json:"modul_ajars"`
	Total      int                 `json:"total"`
	Page       int                 `json:"page"`
	PageSize   int                 `json:"page_size"`
}

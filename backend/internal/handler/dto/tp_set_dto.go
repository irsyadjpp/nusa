package dto

// CreateTPSetRequest represents the request body for creating a TP Set
// Matches OpenAPI schema: CreateTPSetRequest
type CreateTPSetRequest struct {
	CPID             string `json:"cp_id" binding:"required"`
	VersionNo        int    `json:"version_no" binding:"required,min=1"`
	GenerationSource string `json:"generation_source" binding:"required,oneof=AI_GENERATED MANUAL"`
	GenerationReason string `json:"generation_reason"`
}

// TPSetResponse represents the response body for TP Set operations
// Matches OpenAPI schema: TPSetResponse
type TPSetResponse struct {
	ID               string  `json:"id"`
	CPID             string  `json:"cp_id"`
	CPCode           string  `json:"cp_code"`
	CPText           string  `json:"cp_text"`
	VersionNo        int     `json:"version_no"`
	Status           string  `json:"status"`
	GenerationSource string  `json:"generation_source"`
	GenerationReason *string `json:"generation_reason"`
	GeneratedBy      string  `json:"generated_by"`
	GeneratedByName  string  `json:"generated_by_name"`
	AIGenerationID   *string `json:"ai_generation_id"`
	ApprovedBy       *string `json:"approved_by"`
	ApprovedByName   *string `json:"approved_by_name"`
	ApprovedAt       *string `json:"approved_at"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
}

// ListTPSetsResponse represents the response body for listing TP Sets
// Matches OpenAPI schema for GET /tp-sets
type ListTPSetsResponse struct {
	TPSets   []TPSetResponse `json:"tp_sets"`
	Total    int             `json:"total"`
	Page     int             `json:"page"`
	PageSize int             `json:"page_size"`
}

// ApproveTPSetRequest represents the request body for approving a TP Set
// Matches OpenAPI schema: ApproveTPSetRequest
type ApproveTPSetRequest struct {
	Reason string `json:"reason" binding:"required,min=5"`
}

// ApproveTPSetResponse represents the response body for approving a TP Set
// Matches OpenAPI schema for POST /tp-sets/{id}/approve
type ApproveTPSetResponse struct {
	Message string `json:"message"`
}

// CreateTPRequest represents the request body for creating a TP
// Matches OpenAPI schema: CreateTPRequest
type CreateTPRequest struct {
	TPSetID            string      `json:"tp_set_id" binding:"required"`
	SequenceNumber     int         `json:"sequence_number" binding:"required"`
	CPID               string      `json:"cp_id" binding:"required"`
	SubjectID          string      `json:"subject_id" binding:"required"`
	PhaseID            string      `json:"phase_id" binding:"required"`
	ElementID          string      `json:"element_id" binding:"required"`
	SubelementID       string      `json:"subelement_id" binding:"required"`
	Status             string      `json:"status" binding:"required,oneof=DRAFT PENDING APPROVED REJECTED"`
	Title              *string     `json:"title"`
	LearningObjectives interface{} `json:"learning_objectives" binding:"required"`
	TimeAllocation     interface{} `json:"time_allocation" binding:"required"`
	Prerequisites      interface{} `json:"prerequisites"`
	EstimatedWeeks     *int        `json:"estimated_weeks"`
	SuccessCriteria    interface{} `json:"success_criteria"`
}

// TPResponse represents the response body for TP operations
// Matches OpenAPI schema: TPResponse
type TPResponse struct {
	ID                 string      `json:"id"`
	TPSetID            string      `json:"tp_set_id"`
	SequenceNumber     int         `json:"sequence_number"`
	CPID               string      `json:"cp_id"`
	CPCode             string      `json:"cp_code"`
	CPText             string      `json:"cp_text"`
	SubjectID          string      `json:"subject_id"`
	SubjectCode        string      `json:"subject_code"`
	SubjectName        string      `json:"subject_name"`
	PhaseID            string      `json:"phase_id"`
	PhaseCode          string      `json:"phase_code"`
	PhaseName          string      `json:"phase_name"`
	ElementID          string      `json:"element_id"`
	ElementCode        string      `json:"element_code"`
	ElementName        string      `json:"element_name"`
	SubelementID       string      `json:"subelement_id"`
	SubelementCode     string      `json:"subelement_code"`
	SubelementName     string      `json:"subelement_name"`
	UserID             string      `json:"user_id"`
	UserName           string      `json:"user_name"`
	Status             string      `json:"status"`
	Title              *string     `json:"title"`
	LearningObjectives interface{} `json:"learning_objectives"`
	TimeAllocation     interface{} `json:"time_allocation"`
	Prerequisites      interface{} `json:"prerequisites"`
	EstimatedWeeks     *int        `json:"estimated_weeks"`
	SuccessCriteria    interface{} `json:"success_criteria"`
	CreatedAt          string      `json:"created_at"`
	UpdatedAt          string      `json:"updated_at"`
}

// ListTPsResponse represents the response body for listing TPs
// Matches OpenAPI schema for GET /tps
type ListTPsResponse struct {
	TPs      []TPResponse `json:"tps"`
	Total    int          `json:"total"`
	Page     int          `json:"page"`
	PageSize int          `json:"page_size"`
}

// ErrorResponse represents the standard error response
// Matches OpenAPI schema: Error
type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

// ErrorDetail represents error details
type ErrorDetail struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}

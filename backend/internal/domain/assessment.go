package domain

import "time"

// AssessmentType represents the type of assessment
type AssessmentType string

const (
	AssessmentTypeFormative AssessmentType = "FORMATIVE"
	AssessmentTypeSummative AssessmentType = "SUMMATIVE"
)

// Assessment represents an assessment
type Assessment struct {
	ID                      string      `json:"id" db:"id"`
	TPID                    string      `json:"tp_id" db:"tp_id"`                                         // References TP instead of ModulAjar
	TPVersionNo             int         `json:"tp_version_no" db:"tp_version_no"`                         // Snapshot of TP version
	SuccessCriteriaSnapshot interface{} `json:"success_criteria_snapshot" db:"success_criteria_snapshot"` // Snapshot of TP's SuccessCriteria
	// Expanded TP Snapshot
	TPTitleSnapshot              *string        `json:"tp_title_snapshot,omitempty" db:"tp_title_snapshot"`
	TPLearningObjectivesSnapshot interface{}    `json:"tp_learning_objectives_snapshot,omitempty" db:"tp_learning_objectives_snapshot"`
	TPTimeAllocationSnapshot     interface{}    `json:"tp_time_allocation_snapshot,omitempty" db:"tp_time_allocation_snapshot"`
	UserID                       string         `json:"user_id" db:"user_id"`
	AssessmentType               AssessmentType `json:"assessment_type" db:"assessment_type"`
	Status                       WorkflowStatus `json:"status" db:"status"`
	AssessmentItems              interface{}    `json:"assessment_items" db:"assessment_items"`
	AnswerKey                    interface{}    `json:"answer_key" db:"answer_key"`
	ScoringGuidelines            interface{}    `json:"scoring_guidelines" db:"scoring_guidelines"`
	AiConfidenceScore            *float64       `json:"ai_confidence_score,omitempty" db:"ai_confidence_score"`
	AiGeneratedAt                *time.Time     `json:"ai_generated_at,omitempty" db:"ai_generated_at"`
	AiAgentVersion               *string        `json:"ai_agent_version,omitempty" db:"ai_agent_version"`
	VersionNo                    int            `json:"version_no" db:"version_no"`
	IsCurrentVersion             bool           `json:"is_current_version" db:"is_current_version"`
	ParentVersionID              *string        `json:"parent_version_id,omitempty" db:"parent_version_id"`
	CreatedAt                    time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt                    time.Time      `json:"updated_at" db:"updated_at"`
	ApprovedAt                   *time.Time     `json:"approved_at,omitempty" db:"approved_at"`
	ApprovedBy                   *string        `json:"approved_by,omitempty" db:"approved_by"`
}

// AssessmentResponse represents the response for assessment with related data
type AssessmentResponse struct {
	ID                      string      `json:"id"`
	TPID                    string      `json:"tp_id"`
	TPTitle                 string      `json:"tp_title"`
	TPVersionNo             int         `json:"tp_version_no"`
	SuccessCriteriaSnapshot interface{} `json:"success_criteria_snapshot"`
	// Expanded TP Snapshot
	TPLearningObjectivesSnapshot interface{}    `json:"tp_learning_objectives_snapshot,omitempty"`
	TPTimeAllocationSnapshot     interface{}    `json:"tp_time_allocation_snapshot,omitempty"`
	UserID                       string         `json:"user_id"`
	UserName                     string         `json:"user_name"`
	AssessmentType               AssessmentType `json:"assessment_type"`
	Status                       WorkflowStatus `json:"status"`
	AssessmentItems              interface{}    `json:"assessment_items"`
	AnswerKey                    interface{}    `json:"answer_key"`
	ScoringGuidelines            interface{}    `json:"scoring_guidelines"`
	AiConfidenceScore            *float64       `json:"ai_confidence_score,omitempty"`
	AiGeneratedAt                *time.Time     `json:"ai_generated_at,omitempty"`
	AiAgentVersion               *string        `json:"ai_agent_version,omitempty"`
	VersionNo                    int            `json:"version_no"`
	IsCurrentVersion             bool           `json:"is_current_version"`
	ParentVersionID              *string        `json:"parent_version_id,omitempty"`
	CreatedAt                    time.Time      `json:"created_at"`
	UpdatedAt                    time.Time      `json:"updated_at"`
	ApprovedAt                   *time.Time     `json:"approved_at,omitempty"`
	ApprovedBy                   *string        `json:"approved_by,omitempty"`
	ApprovedByName               *string        `json:"approved_by_name,omitempty"`
}

// CreateAssessmentRequest represents the request to create an assessment
type CreateAssessmentRequest struct {
	TPID                         string         `json:"tp_id" binding:"required"`
	TPVersionNo                  int            `json:"tp_version_no" binding:"required,min=1"`
	SuccessCriteriaSnapshot      interface{}    `json:"success_criteria_snapshot" binding:"required"`
	TPTitleSnapshot              *string        `json:"tp_title_snapshot,omitempty"`
	TPLearningObjectivesSnapshot interface{}    `json:"tp_learning_objectives_snapshot,omitempty"`
	TPTimeAllocationSnapshot     interface{}    `json:"tp_time_allocation_snapshot,omitempty"`
	AssessmentType               AssessmentType `json:"assessment_type" binding:"required,oneof=FORMATIVE SUMMATIVE"`
	AssessmentItems              interface{}    `json:"assessment_items" binding:"required"`
	AnswerKey                    interface{}    `json:"answer_key" binding:"required"`
	ScoringGuidelines            interface{}    `json:"scoring_guidelines" binding:"required"`
}

// UpdateAssessmentRequest represents the request to update an assessment
type UpdateAssessmentRequest struct {
	AssessmentItems   interface{}     `json:"assessment_items,omitempty"`
	AnswerKey         interface{}     `json:"answer_key,omitempty"`
	ScoringGuidelines interface{}     `json:"scoring_guidelines,omitempty"`
	Status            *WorkflowStatus `json:"status,omitempty"`
}

// RubricType represents the type of rubric
type RubricType string

const (
	RubricTypeAnalytic RubricType = "ANALYTIC"
	RubricTypeHolistic RubricType = "HOLISTIC"
)

// Rubric represents a rubric
type Rubric struct {
	ID                  string         `json:"id" db:"id"`
	AssessmentID        string         `json:"assessment_id" db:"assessment_id"`
	UserID              string         `json:"user_id" db:"user_id"`
	RubricType          RubricType     `json:"rubric_type" db:"rubric_type"`
	Status              WorkflowStatus `json:"status" db:"status"`
	PerformanceCriteria interface{}    `json:"performance_criteria" db:"performance_criteria"`
	PerformanceLevels   interface{}    `json:"performance_levels" db:"performance_levels"`
	ScoringGuidelines   interface{}    `json:"scoring_guidelines" db:"scoring_guidelines"`
	AiConfidenceScore   *float64       `json:"ai_confidence_score,omitempty" db:"ai_confidence_score"`
	AiGeneratedAt       *time.Time     `json:"ai_generated_at,omitempty" db:"ai_generated_at"`
	AiAgentVersion      *string        `json:"ai_agent_version,omitempty" db:"ai_agent_version"`
	VersionNo           int            `json:"version_no" db:"version_no"`
	IsCurrentVersion    bool           `json:"is_current_version" db:"is_current_version"`
	ParentVersionID     *string        `json:"parent_version_id,omitempty" db:"parent_version_id"`
	CreatedAt           time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at" db:"updated_at"`
	ApprovedAt          *time.Time     `json:"approved_at,omitempty" db:"approved_at"`
	ApprovedBy          *string        `json:"approved_by,omitempty" db:"approved_by"`
}

// RubricResponse represents the response for rubric with related data
type RubricResponse struct {
	ID                  string         `json:"id"`
	AssessmentID        string         `json:"assessment_id"`
	AssessmentType      string         `json:"assessment_type"`
	UserID              string         `json:"user_id"`
	UserName            string         `json:"user_name"`
	RubricType          RubricType     `json:"rubric_type"`
	Status              WorkflowStatus `json:"status"`
	PerformanceCriteria interface{}    `json:"performance_criteria"`
	PerformanceLevels   interface{}    `json:"performance_levels"`
	ScoringGuidelines   interface{}    `json:"scoring_guidelines"`
	AiConfidenceScore   *float64       `json:"ai_confidence_score,omitempty"`
	AiGeneratedAt       *time.Time     `json:"ai_generated_at,omitempty"`
	AiAgentVersion      *string        `json:"ai_agent_version,omitempty"`
	VersionNo           int            `json:"version_no"`
	IsCurrentVersion    bool           `json:"is_current_version"`
	ParentVersionID     *string        `json:"parent_version_id,omitempty"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	ApprovedAt          *time.Time     `json:"approved_at,omitempty"`
	ApprovedBy          *string        `json:"approved_by,omitempty"`
	ApprovedByName      *string        `json:"approved_by_name,omitempty"`
}

// CreateRubricRequest represents the request to create a rubric
type CreateRubricRequest struct {
	AssessmentID        string      `json:"assessment_id" binding:"required"`
	RubricType          RubricType  `json:"rubric_type" binding:"required,oneof=ANALYTIC HOLISTIC"`
	PerformanceCriteria interface{} `json:"performance_criteria" binding:"required"`
	PerformanceLevels   interface{} `json:"performance_levels" binding:"required"`
	ScoringGuidelines   interface{} `json:"scoring_guidelines" binding:"required"`
}

// UpdateRubricRequest represents the request to update a rubric
type UpdateRubricRequest struct {
	PerformanceCriteria interface{}     `json:"performance_criteria,omitempty"`
	PerformanceLevels   interface{}     `json:"performance_levels,omitempty"`
	ScoringGuidelines   interface{}     `json:"scoring_guidelines,omitempty"`
	Status              *WorkflowStatus `json:"status,omitempty"`
}

// EvidenceType represents the type of evidence
type EvidenceType string

const (
	EvidenceTypeStudentWork      EvidenceType = "STUDENT_WORK"
	EvidenceTypeAssessmentResult EvidenceType = "ASSESSMENT_RESULT"
	EvidenceTypeObservation      EvidenceType = "OBSERVATION"
)

// EvidenceStatus represents the status of evidence
type EvidenceStatus string

const (
	EvidenceStatusCollected EvidenceStatus = "COLLECTED"
	EvidenceStatusLinked    EvidenceStatus = "LINKED"
	EvidenceStatusEvaluated EvidenceStatus = "EVALUATED"
)

// Evidence represents student learning evidence (Aggregate Root)
type Evidence struct {
	ID              string         `json:"id" db:"id"`
	StudentID       string         `json:"student_id" db:"student_id"`
	AssessmentID    string         `json:"assessment_id" db:"assessment_id"`
	UserID          string         `json:"user_id" db:"user_id"`
	EvidenceType    EvidenceType   `json:"evidence_type" db:"evidence_type"`
	Status          EvidenceStatus `json:"status" db:"status"`
	EvidenceData    interface{}    `json:"evidence_data" db:"evidence_data"`
	TeacherNotes    *string        `json:"teacher_notes,omitempty" db:"teacher_notes"`
	RubricID        *string        `json:"rubric_id,omitempty" db:"rubric_id"`
	LinkedCriteria  interface{}    `json:"linked_criteria,omitempty" db:"linked_criteria"`
	EvaluationNotes *string        `json:"evaluation_notes,omitempty" db:"evaluation_notes"`
	Evaluations     []Evaluation   `json:"evaluations"` // Child entities
	CreatedAt       time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at" db:"updated_at"`
}

// EvidenceResponse represents the response for evidence with related data
type EvidenceResponse struct {
	ID              string         `json:"id"`
	StudentID       string         `json:"student_id"`
	StudentName     string         `json:"student_name"`
	AssessmentID    string         `json:"assessment_id"`
	AssessmentType  string         `json:"assessment_type"`
	UserID          string         `json:"user_id"`
	UserName        string         `json:"user_name"`
	EvidenceType    EvidenceType   `json:"evidence_type"`
	Status          EvidenceStatus `json:"status"`
	EvidenceData    interface{}    `json:"evidence_data"`
	TeacherNotes    *string        `json:"teacher_notes,omitempty"`
	RubricID        *string        `json:"rubric_id,omitempty"`
	LinkedCriteria  interface{}    `json:"linked_criteria,omitempty"`
	EvaluationNotes *string        `json:"evaluation_notes,omitempty"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
}

// CreateEvidenceRequest represents the request to create evidence
type CreateEvidenceRequest struct {
	StudentID      string       `json:"student_id" binding:"required"`
	AssessmentID   string       `json:"assessment_id" binding:"required"`
	EvidenceType   EvidenceType `json:"evidence_type" binding:"required,oneof=STUDENT_WORK ASSESSMENT_RESULT OBSERVATION"`
	EvidenceData   interface{}  `json:"evidence_data" binding:"required"`
	TeacherNotes   *string      `json:"teacher_notes,omitempty"`
	RubricID       *string      `json:"rubric_id,omitempty"`
	LinkedCriteria interface{}  `json:"linked_criteria,omitempty"`
}

// UpdateEvidenceRequest represents the request to update evidence
type UpdateEvidenceRequest struct {
	EvidenceData    interface{}     `json:"evidence_data,omitempty"`
	TeacherNotes    *string         `json:"teacher_notes,omitempty"`
	RubricID        *string         `json:"rubric_id,omitempty"`
	LinkedCriteria  interface{}     `json:"linked_criteria,omitempty"`
	EvaluationNotes *string         `json:"evaluation_notes,omitempty"`
	Status          *EvidenceStatus `json:"status,omitempty"`
}

// PerformanceLevel represents the performance level
type PerformanceLevel string

const (
	PerformanceLevelExcellent  PerformanceLevel = "EXCELLENT"
	PerformanceLevelProficient PerformanceLevel = "PROFICIENT"
	PerformanceLevelDeveloping PerformanceLevel = "DEVELOPING"
	PerformanceLevelBeginning  PerformanceLevel = "BEGINNING"
)

// Evaluation represents student performance evaluation (Child Entity of Evidence)
type Evaluation struct {
	ID                string           `json:"id" db:"id"`
	StudentID         string           `json:"student_id" db:"student_id"`
	RubricID          string           `json:"rubric_id" db:"rubric_id"`
	EvidenceID        string           `json:"evidence_id" db:"evidence_id"`
	UserID            string           `json:"user_id" db:"user_id"`
	PerformanceScores interface{}      `json:"performance_scores" db:"performance_scores"`
	TotalScore        int              `json:"total_score" db:"total_score"`
	MaxScore          int              `json:"max_score" db:"max_score"`
	PerformanceLevel  PerformanceLevel `json:"performance_level" db:"performance_level"`
	TeacherFeedback   *string          `json:"teacher_feedback,omitempty" db:"teacher_feedback"`
	RevisionNo        int              `json:"revision_no" db:"revision_no"`
	IsCurrentVersion  bool             `json:"is_current_version" db:"is_current_version"`
	ParentRevisionID  *string          `json:"parent_revision_id,omitempty" db:"parent_revision_id"`
	EvaluatedAt       time.Time        `json:"evaluated_at" db:"evaluated_at"`
	CreatedAt         time.Time        `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time        `json:"updated_at" db:"updated_at"`
}

// EvaluationResponse represents the response for evaluation with related data
type EvaluationResponse struct {
	ID                string           `json:"id"`
	StudentID         string           `json:"student_id"`
	StudentName       string           `json:"student_name"`
	RubricID          string           `json:"rubric_id"`
	RubricType        string           `json:"rubric_type"`
	EvidenceID        string           `json:"evidence_id"`
	EvidenceType      string           `json:"evidence_type"`
	UserID            string           `json:"user_id"`
	UserName          string           `json:"user_name"`
	PerformanceScores interface{}      `json:"performance_scores"`
	TotalScore        int              `json:"total_score"`
	MaxScore          int              `json:"max_score"`
	PerformanceLevel  PerformanceLevel `json:"performance_level"`
	TeacherFeedback   *string          `json:"teacher_feedback,omitempty"`
	RevisionNo        int              `json:"revision_no"`
	EvaluatedAt       time.Time        `json:"evaluated_at"`
	CreatedAt         time.Time        `json:"created_at"`
	UpdatedAt         time.Time        `json:"updated_at"`
}

// CreateEvaluationRequest represents the request to create an evaluation
type CreateEvaluationRequest struct {
	StudentID         string           `json:"student_id" binding:"required"`
	RubricID          string           `json:"rubric_id" binding:"required"`
	EvidenceID        string           `json:"evidence_id" binding:"required"`
	PerformanceScores interface{}      `json:"performance_scores" binding:"required"`
	TotalScore        int              `json:"total_score" binding:"required,min=0"`
	MaxScore          int              `json:"max_score" binding:"required,min=1"`
	PerformanceLevel  PerformanceLevel `json:"performance_level" binding:"required,oneof=EXCELLENT PROFICIENT DEVELOPING BEGINNING"`
	TeacherFeedback   *string          `json:"teacher_feedback,omitempty"`
}

// UpdateEvaluationRequest represents the request to update an evaluation
type UpdateEvaluationRequest struct {
	PerformanceScores interface{}       `json:"performance_scores,omitempty"`
	TotalScore        *int              `json:"total_score,omitempty"`
	MaxScore          *int              `json:"max_score,omitempty"`
	PerformanceLevel  *PerformanceLevel `json:"performance_level,omitempty"`
	TeacherFeedback   *string           `json:"teacher_feedback,omitempty"`
}

// EvaluationFeedbackHistory represents the history of teacher feedback changes for an evaluation
type EvaluationFeedbackHistory struct {
	ID              string    `json:"id" db:"id"`
	EvaluationID    string    `json:"evaluation_id" db:"evaluation_id"`
	TeacherFeedback string    `json:"teacher_feedback" db:"teacher_feedback"`
	ChangedBy       string    `json:"changed_by" db:"changed_by"`
	ChangedAt       time.Time `json:"changed_at" db:"changed_at"`
}

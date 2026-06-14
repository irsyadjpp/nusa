package dto

import (
	"time"
)

// ==================== Assessment DTOs ====================

// AssessmentType represents the type of assessment
type AssessmentType string

const (
	AssessmentTypeFormative AssessmentType = "FORMATIVE"
	AssessmentTypeSummative AssessmentType = "SUMMATIVE"
)

// SuccessCriteriaSnapshot represents KKTP (Kriteria Ketuntasan Tujuan Pembelajaran)
// This is a Value Object embedded in TP but also snapshotted in Assessment
type SuccessCriteriaSnapshot struct {
	MasteryLevel         string  `json:"mastery_level"`
	PerformanceIndicator string  `json:"performance_indicator"`
	MinimumRequirement   string  `json:"minimum_requirement"`
	Weight               float64 `json:"weight"`
	Description          string  `json:"description"`
}

// AssessmentItem represents a single question or task in an assessment
type AssessmentItem struct {
	ID             string                 `json:"id"`
	SequenceNumber int                    `json:"sequence_number"`
	Question       string                 `json:"question"`
	QuestionType   string                 `json:"question_type"` // multiple_choice, essay, true_false, etc.
	Options        []AssessmentItemOption `json:"options,omitempty"`
	Points         float64                `json:"points"`
	IsMandatory    bool                   `json:"is_mandatory"`
	Metadata       interface{}            `json:"metadata,omitempty"`
}

// AssessmentItemOption represents a choice for multiple choice questions
type AssessmentItemOption struct {
	ID         string `json:"id"`
	OptionText string `json:"option_text"`
	IsCorrect  bool   `json:"is_correct"`
}

// AnswerKey maps assessment item IDs to correct answers
type AnswerKey struct {
	ItemID      string  `json:"item_id"`
	Answer      string  `json:"answer"` // Can be option ID or text answer
	Explanation string  `json:"explanation,omitempty"`
	Points      float64 `json:"points"`
}

// ScoringGuideline represents scoring rules for an assessment
type ScoringGuideline struct {
	Criterion   string  `json:"criterion"`
	Description string  `json:"description"`
	MaxPoints   float64 `json:"max_points"`
	Weight      float64 `json:"weight"`
	RubricLevel string  `json:"rubric_level,omitempty"` // For analytic rubrics
}

// CreateAssessmentRequest represents the request to create an assessment
type CreateAssessmentRequest struct {
	TPID                         string                    `json:"tp_id" binding:"required"`
	TPVersionNo                  int                       `json:"tp_version_no" binding:"required,min=1"`
	SuccessCriteriaSnapshot      []SuccessCriteriaSnapshot `json:"success_criteria_snapshot" binding:"required"`
	TPTitleSnapshot              *string                   `json:"tp_title_snapshot,omitempty"`
	TPLearningObjectivesSnapshot []string                  `json:"tp_learning_objectives_snapshot,omitempty"`
	TPTimeAllocationSnapshot     map[string]int            `json:"tp_time_allocation_snapshot,omitempty"`
	AssessmentType               AssessmentType            `json:"assessment_type" binding:"required,oneof=FORMATIVE SUMMATIVE"`
	AssessmentItems              []AssessmentItem          `json:"assessment_items" binding:"required"`
	AnswerKey                    []AnswerKey               `json:"answer_key" binding:"required"`
	ScoringGuidelines            []ScoringGuideline        `json:"scoring_guidelines" binding:"required"`
}

// UpdateAssessmentRequest represents the request to update an assessment
type UpdateAssessmentRequest struct {
	AssessmentItems   []AssessmentItem   `json:"assessment_items,omitempty"`
	AnswerKey         []AnswerKey        `json:"answer_key,omitempty"`
	ScoringGuidelines []ScoringGuideline `json:"scoring_guidelines,omitempty"`
	Status            *string            `json:"status,omitempty"`
}

// AssessmentResponse represents the response for assessment with related data
type AssessmentResponse struct {
	ID                      string                    `json:"id"`
	TPID                    string                    `json:"tp_id"`
	TPTitle                 string                    `json:"tp_title"`
	TPVersionNo             int                       `json:"tp_version_no"`
	SuccessCriteriaSnapshot []SuccessCriteriaSnapshot `json:"success_criteria_snapshot"`
	// Expanded TP Snapshot
	TPLearningObjectivesSnapshot []string           `json:"tp_learning_objectives_snapshot,omitempty"`
	TPTimeAllocationSnapshot     map[string]int     `json:"tp_time_allocation_snapshot,omitempty"`
	UserID                       string             `json:"user_id"`
	UserName                     string             `json:"user_name"`
	AssessmentType               AssessmentType     `json:"assessment_type"`
	Status                       string             `json:"status"`
	AssessmentItems              []AssessmentItem   `json:"assessment_items"`
	AnswerKey                    []AnswerKey        `json:"answer_key"`
	ScoringGuidelines            []ScoringGuideline `json:"scoring_guidelines"`
	AiConfidenceScore            *float64           `json:"ai_confidence_score,omitempty"`
	AiGeneratedAt                *time.Time         `json:"ai_generated_at,omitempty"`
	AiAgentVersion               *string            `json:"ai_agent_version,omitempty"`
	VersionNo                    int                `json:"version_no"`
	IsCurrentVersion             bool               `json:"is_current_version"`
	ParentVersionID              *string            `json:"parent_version_id,omitempty"`
	CreatedAt                    time.Time          `json:"created_at"`
	UpdatedAt                    time.Time          `json:"updated_at"`
	ApprovedAt                   *time.Time         `json:"approved_at,omitempty"`
	ApprovedBy                   *string            `json:"approved_by,omitempty"`
	ApprovedByName               *string            `json:"approved_by_name,omitempty"`
}

// ==================== Rubric DTOs ====================

// RubricType represents the type of rubric
type RubricType string

const (
	RubricTypeAnalytic RubricType = "ANALYTIC"
	RubricTypeHolistic RubricType = "HOLISTIC"
)

// PerformanceCriterion represents a single criterion in a rubric
type PerformanceCriterion struct {
	ID          string  `json:"id"`
	Description string  `json:"description"`
	Weight      float64 `json:"weight"`
	Required    bool    `json:"required"`
	MaxPoints   float64 `json:"max_points"`
}

// PerformanceLevel represents achievement levels (e.g., Exemplary, Proficient, Developing)
type PerformanceLevel struct {
	Level       string `json:"level"`
	Description string `json:"description"`
	ScoreRange  string `json:"score_range"`
	ColorCode   string `json:"color_code"`
}

// CreateRubricRequest represents the request to create a rubric
type CreateRubricRequest struct {
	AssessmentID        string                 `json:"assessment_id" binding:"required"`
	RubricType          RubricType             `json:"rubric_type" binding:"required,oneof=ANALYTIC HOLISTIC"`
	PerformanceCriteria []PerformanceCriterion `json:"performance_criteria" binding:"required"`
	PerformanceLevels   []PerformanceLevel     `json:"performance_levels" binding:"required"`
	ScoringGuidelines   []ScoringGuideline     `json:"scoring_guidelines" binding:"required"`
}

// UpdateRubricRequest represents the request to update a rubric
type UpdateRubricRequest struct {
	PerformanceCriteria []PerformanceCriterion `json:"performance_criteria,omitempty"`
	PerformanceLevels   []PerformanceLevel     `json:"performance_levels,omitempty"`
	ScoringGuidelines   []ScoringGuideline     `json:"scoring_guidelines,omitempty"`
	Status              *string                `json:"status,omitempty"`
}

// RubricResponse represents the response for rubric with related data
type RubricResponse struct {
	ID                  string                 `json:"id"`
	AssessmentID        string                 `json:"assessment_id"`
	AssessmentType      string                 `json:"assessment_type"`
	UserID              string                 `json:"user_id"`
	UserName            string                 `json:"user_name"`
	RubricType          RubricType             `json:"rubric_type"`
	Status              string                 `json:"status"`
	PerformanceCriteria []PerformanceCriterion `json:"performance_criteria"`
	PerformanceLevels   []PerformanceLevel     `json:"performance_levels"`
	ScoringGuidelines   []ScoringGuideline     `json:"scoring_guidelines"`
	AiConfidenceScore   *float64               `json:"ai_confidence_score,omitempty"`
	AiGeneratedAt       *time.Time             `json:"ai_generated_at,omitempty"`
	AiAgentVersion      *string                `json:"ai_agent_version,omitempty"`
	VersionNo           int                    `json:"version_no"`
	IsCurrentVersion    bool                   `json:"is_current_version"`
	ParentVersionID     *string                `json:"parent_version_id,omitempty"`
	CreatedAt           time.Time              `json:"created_at"`
	UpdatedAt           time.Time              `json:"updated_at"`
	ApprovedAt          *time.Time             `json:"approved_at,omitempty"`
	ApprovedBy          *string                `json:"approved_by,omitempty"`
	ApprovedByName      *string                `json:"approved_by_name,omitempty"`
}

// ==================== Evidence DTOs ====================

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

// ==================== Evaluation DTOs ====================

// EvaluationPerformanceLevel represents the performance level
type EvaluationPerformanceLevel string

const (
	EvaluationPerformanceLevelExcellent  EvaluationPerformanceLevel = "EXCELLENT"
	EvaluationPerformanceLevelProficient EvaluationPerformanceLevel = "PROFICIENT"
	EvaluationPerformanceLevelDeveloping EvaluationPerformanceLevel = "DEVELOPING"
	EvaluationPerformanceLevelBeginning  EvaluationPerformanceLevel = "BEGINNING"
)

// CreateEvaluationRequest represents the request to create an evaluation
type CreateEvaluationRequest struct {
	StudentID         string                     `json:"student_id" binding:"required"`
	RubricID          string                     `json:"rubric_id" binding:"required"`
	EvidenceID        string                     `json:"evidence_id" binding:"required"`
	PerformanceScores interface{}                `json:"performance_scores" binding:"required"`
	TotalScore        int                        `json:"total_score" binding:"required,min=0"`
	MaxScore          int                        `json:"max_score" binding:"required,min=1"`
	PerformanceLevel  EvaluationPerformanceLevel `json:"performance_level" binding:"required,oneof=EXCELLENT PROFICIENT DEVELOPING BEGINNING"`
	TeacherFeedback   *string                    `json:"teacher_feedback,omitempty"`
}

// UpdateEvaluationRequest represents the request to update an evaluation
type UpdateEvaluationRequest struct {
	PerformanceScores interface{}                 `json:"performance_scores,omitempty"`
	TotalScore        *int                        `json:"total_score,omitempty"`
	MaxScore          *int                        `json:"max_score,omitempty"`
	PerformanceLevel  *EvaluationPerformanceLevel `json:"performance_level,omitempty"`
	TeacherFeedback   *string                     `json:"teacher_feedback,omitempty"`
}

// EvaluationResponse represents the response for evaluation with related data
type EvaluationResponse struct {
	ID                string                     `json:"id"`
	StudentID         string                     `json:"student_id"`
	StudentName       string                     `json:"student_name"`
	RubricID          string                     `json:"rubric_id"`
	RubricType        string                     `json:"rubric_type"`
	EvidenceID        string                     `json:"evidence_id"`
	EvidenceType      string                     `json:"evidence_type"`
	UserID            string                     `json:"user_id"`
	UserName          string                     `json:"user_name"`
	PerformanceScores interface{}                `json:"performance_scores"`
	TotalScore        int                        `json:"total_score"`
	MaxScore          int                        `json:"max_score"`
	PerformanceLevel  EvaluationPerformanceLevel `json:"performance_level"`
	TeacherFeedback   *string                    `json:"teacher_feedback,omitempty"`
	RevisionNo        int                        `json:"revision_no"`
	EvaluatedAt       time.Time                  `json:"evaluated_at"`
	CreatedAt         time.Time                  `json:"created_at"`
	UpdatedAt         time.Time                  `json:"updated_at"`
}

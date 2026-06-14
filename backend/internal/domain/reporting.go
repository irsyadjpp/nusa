package domain

import "time"

// ReportLanguage represents the language of the report
type ReportLanguage string

const (
	ReportLanguageIndonesian ReportLanguage = "INDONESIAN"
)

// NarrativeReport represents a narrative report
type NarrativeReport struct {
	ID           string         `json:"id" db:"id"`
	StudentID    string         `json:"student_id" db:"student_id"`
	ClassID      string         `json:"class_id" db:"class_id"`
	UserID       string         `json:"user_id" db:"user_id"`
	Status       WorkflowStatus `json:"status" db:"status"`
	ReportPeriod interface{}    `json:"report_period" db:"report_period"`
	Language     ReportLanguage `json:"language" db:"language"`
	Content      interface{}    `json:"content" db:"content"`

	// Achievement Integration
	AchievementSummaryID        *string     `json:"achievement_summary_id,omitempty" db:"achievement_summary_id"`
	AchievementData             interface{} `json:"achievement_data,omitempty" db:"achievement_data"`
	LastAchievementCalculatedAt *time.Time  `json:"last_achievement_calculated_at,omitempty" db:"last_achievement_calculated_at"`

	// Metadata
	AiConfidenceScore *float64   `json:"ai_confidence_score,omitempty" db:"ai_confidence_score"`
	AiGeneratedAt     *time.Time `json:"ai_generated_at,omitempty" db:"ai_generated_at"`
	AiAgentVersion    *string    `json:"ai_agent_version,omitempty" db:"ai_agent_version"`
	VersionNo         int        `json:"version_no" db:"version_no"`
	IsCurrentVersion  bool       `json:"is_current_version" db:"is_current_version"`
	ParentVersionID   *string    `json:"parent_version_id,omitempty" db:"parent_version_id"`
	CreatedAt         time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at" db:"updated_at"`
	ApprovedAt        *time.Time `json:"approved_at,omitempty" db:"approved_at"`
	ApprovedBy        *string    `json:"approved_by,omitempty" db:"approved_by"`
}

// NarrativeReportResponse represents the response for narrative report with related data
type NarrativeReportResponse struct {
	ID           string         `json:"id"`
	StudentID    string         `json:"student_id"`
	StudentName  string         `json:"student_name"`
	ClassID      string         `json:"class_id"`
	ClassName    string         `json:"class_name"`
	UserID       string         `json:"user_id"`
	UserName     string         `json:"user_name"`
	Status       WorkflowStatus `json:"status"`
	ReportPeriod interface{}    `json:"report_period"`
	Language     ReportLanguage `json:"language"`
	Content      interface{}    `json:"content"`

	// Achievement Integration
	AchievementSummaryID        *string     `json:"achievement_summary_id,omitempty"`
	AchievementData             interface{} `json:"achievement_data,omitempty"`
	LastAchievementCalculatedAt *time.Time  `json:"last_achievement_calculated_at,omitempty"`

	// Metadata
	AiConfidenceScore *float64   `json:"ai_confidence_score,omitempty"`
	AiGeneratedAt     *time.Time `json:"ai_generated_at,omitempty"`
	AiAgentVersion    *string    `json:"ai_agent_version,omitempty"`
	VersionNo         int        `json:"version_no"`
	IsCurrentVersion  bool       `json:"is_current_version"`
	ParentVersionID   *string    `json:"parent_version_id,omitempty"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	ApprovedAt        *time.Time `json:"approved_at,omitempty"`
	ApprovedBy        *string    `json:"approved_by,omitempty"`
	ApprovedByName    *string    `json:"approved_by_name,omitempty"`
}

// CreateNarrativeReportRequest represents the request to create a narrative report
type CreateNarrativeReportRequest struct {
	StudentID    string         `json:"student_id" binding:"required"`
	ClassID      string         `json:"class_id" binding:"required"`
	ReportPeriod interface{}    `json:"report_period" binding:"required"`
	Language     ReportLanguage `json:"language" binding:"required,oneof=INDONESIAN"`
	Content      interface{}    `json:"content" binding:"required"`
}

// UpdateNarrativeReportRequest represents the request to update a narrative report
type UpdateNarrativeReportRequest struct {
	ReportPeriod interface{}     `json:"report_period,omitempty"`
	Content      interface{}     `json:"content,omitempty"`
	Status       *WorkflowStatus `json:"status,omitempty"`
}

// GenerateNarrativeReportRequest represents the request to generate a narrative report
type GenerateNarrativeReportRequest struct {
	StudentID    string         `json:"student_id" binding:"required"`
	ReportPeriod interface{}    `json:"report_period" binding:"required"`
	Language     ReportLanguage `json:"language" binding:"required,oneof=INDONESIAN"`
}

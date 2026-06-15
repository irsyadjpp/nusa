package dto

import "time"

// CreateNarrativeReportRequest represents the request to create a narrative report
type CreateNarrativeReportRequest struct {
	StudentID      string `json:"student_id" binding:"required"`
	ClassID        string `json:"class_id" binding:"required"`
	AcademicYearID string `json:"academic_year_id" binding:"required"`
	SemesterID     string `json:"semester_id" binding:"required"`
	SubjectID      string `json:"subject_id" binding:"required"`
	ReportPeriod   string `json:"report_period" binding:"required"`
	TeacherID      string `json:"teacher_id" binding:"required"`
}

// UpdateNarrativeReportRequest represents the request to update a narrative report
type UpdateNarrativeReportRequest struct {
	Content     *string    `json:"content,omitempty"`
	TeacherID   *string    `json:"teacher_id,omitempty"`
	Status      *string    `json:"status,omitempty"`
	SubmittedAt *time.Time `json:"submitted_at,omitempty"`
}

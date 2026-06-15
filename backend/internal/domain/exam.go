package domain

import (
	"fmt"
	"time"
)

// ExamStatus represents the status of an exam
type ExamStatus string

const (
	ExamStatusScheduled  ExamStatus = "SCHEDULED"
	ExamStatusInProgress ExamStatus = "IN_PROGRESS"
	ExamStatusCompleted  ExamStatus = "COMPLETED"
	ExamStatusCancelled  ExamStatus = "CANCELLED"
)

// Exam represents an exam
type Exam struct {
	ID              string     `json:"id" db:"id"`
	ClassID         string     `json:"class_id" db:"class_id"`
	AssessmentID    string     `json:"assessment_id" db:"assessment_id"`
	ExamDate        time.Time  `json:"exam_date" db:"exam_date"`
	StartTime       string     `json:"start_time" db:"start_time"`
	DurationMinutes int        `json:"duration_minutes" db:"duration_minutes"`
	Room            *string    `json:"room,omitempty" db:"room"`
	Status          string     `json:"status" db:"status"`
	CreatedAt       time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at" db:"updated_at"`
	CreatedBy       *string    `json:"created_by,omitempty" db:"created_by"`
	UpdatedBy       *string    `json:"updated_by,omitempty" db:"updated_by"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

// CreateExamRequest represents the request to create an exam
type CreateExamRequest struct {
	ClassID         string    `json:"class_id" binding:"required"`
	AssessmentID    string    `json:"assessment_id" binding:"required"`
	ExamDate        time.Time `json:"exam_date" binding:"required"`
	StartTime       string    `json:"start_time" binding:"required"`
	DurationMinutes int       `json:"duration_minutes" binding:"required,min=1"`
	Room            *string   `json:"room,omitempty" binding:"omitempty,max=100"`
}

// UpdateExamRequest represents the request to update an exam
type UpdateExamRequest struct {
	ExamDate        *time.Time  `json:"exam_date,omitempty"`
	StartTime       *string     `json:"start_time,omitempty" binding:"omitempty"`
	DurationMinutes *int        `json:"duration_minutes,omitempty" binding:"omitempty,min=1"`
	Room            *string     `json:"room,omitempty" binding:"omitempty,max=100"`
	Status          *ExamStatus `json:"status,omitempty" binding:"omitempty,oneof=SCHEDULED IN_PROGRESS COMPLETED CANCELLED"`
}

// ExamResponse represents the exam data returned to clients
type ExamResponse struct {
	ID              string     `json:"id"`
	ClassID         string     `json:"class_id"`
	ClassName       *string    `json:"class_name,omitempty"`
	AssessmentID    string     `json:"assessment_id"`
	AssessmentType  *string    `json:"assessment_type,omitempty"`
	ExamDate        string     `json:"exam_date"`
	StartTime       string     `json:"start_time"`
	DurationMinutes int        `json:"duration_minutes"`
	Room            *string    `json:"room,omitempty"`
	Status          ExamStatus `json:"status"`
	CreatedAt       string     `json:"created_at"`
	UpdatedAt       string     `json:"updated_at"`
	CreatedBy       *string    `json:"created_by,omitempty"`
	CreatedByName   *string    `json:"created_by_name,omitempty"`
	UpdatedBy       *string    `json:"updated_by,omitempty"`
	UpdatedByName   *string    `json:"updated_by_name,omitempty"`
}

// ToExamResponse converts Exam to ExamResponse
func (e *Exam) ToExamResponse(className, assessmentType, createdByName, updatedByName string) *ExamResponse {
	var classNamePtr, assessmentTypePtr, roomPtr, createdByPtr, createdByNamePtr, updatedByPtr, updatedByNamePtr *string

	if className != "" {
		classNamePtr = &className
	}
	if assessmentType != "" {
		assessmentTypePtr = &assessmentType
	}
	if e.Room != nil {
		roomPtr = e.Room
	}
	if e.CreatedBy != nil {
		createdByPtr = e.CreatedBy
	}
	if createdByName != "" {
		createdByNamePtr = &createdByName
	}
	if e.UpdatedBy != nil {
		updatedByPtr = e.UpdatedBy
	}
	if updatedByName != "" {
		updatedByNamePtr = &updatedByName
	}

	return &ExamResponse{
		ID:              e.ID,
		ClassID:         e.ClassID,
		ClassName:       classNamePtr,
		AssessmentID:    e.AssessmentID,
		AssessmentType:  assessmentTypePtr,
		ExamDate:        e.ExamDate.Format(time.RFC3339),
		StartTime:       e.StartTime,
		DurationMinutes: e.DurationMinutes,
		Room:            roomPtr,
		Status:          ExamStatus(e.Status),
		CreatedAt:       e.CreatedAt.Format(time.RFC3339),
		UpdatedAt:       e.UpdatedAt.Format(time.RFC3339),
		CreatedBy:       createdByPtr,
		CreatedByName:   createdByNamePtr,
		UpdatedBy:       updatedByPtr,
		UpdatedByName:   updatedByNamePtr,
	}
}

// Validate validates the exam entity
func (e *Exam) Validate() error {
	if e.ID == "" {
		return fmt.Errorf("id is required")
	}
	if e.ClassID == "" {
		return fmt.Errorf("class_id is required")
	}
	if e.AssessmentID == "" {
		return fmt.Errorf("assessment_id is required")
	}
	if e.ExamDate.IsZero() {
		return fmt.Errorf("exam_date is required")
	}
	if e.StartTime == "" {
		return fmt.Errorf("start_time is required")
	}
	if e.DurationMinutes <= 0 {
		return fmt.Errorf("duration_minutes must be greater than 0")
	}
	if e.Status == "" {
		return fmt.Errorf("status is required")
	}
	validStatuses := map[string]bool{
		string(ExamStatusScheduled):  true,
		string(ExamStatusInProgress): true,
		string(ExamStatusCompleted):  true,
		string(ExamStatusCancelled):  true,
	}
	if !validStatuses[e.Status] {
		return fmt.Errorf("invalid status: %s", e.Status)
	}
	return nil
}

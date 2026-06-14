package domain

import (
	"fmt"
	"time"
)

// AssignmentStatus represents the status of an assignment
type AssignmentStatus string

const (
	AssignmentStatusAssigned   AssignmentStatus = "ASSIGNED"
	AssignmentStatusInProgress AssignmentStatus = "IN_PROGRESS"
	AssignmentStatusSubmitted  AssignmentStatus = "SUBMITTED"
	AssignmentStatusGraded     AssignmentStatus = "GRADED"
	AssignmentStatusCancelled  AssignmentStatus = "CANCELLED"
)

// Assignment represents an assignment
type Assignment struct {
	ID           string     `json:"id" db:"id"`
	ClassID      string     `json:"class_id" db:"class_id"`
	AssessmentID string     `json:"assessment_id" db:"assessment_id"`
	Title        string     `json:"title" db:"title"`
	Description  *string    `json:"description,omitempty" db:"description"`
	DueDate      time.Time  `json:"due_date" db:"due_date"`
	MaxScore     int        `json:"max_score" db:"max_score"`
	Status       string     `json:"status" db:"status"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
	CreatedBy    *string    `json:"created_by,omitempty" db:"created_by"`
	UpdatedBy    *string    `json:"updated_by,omitempty" db:"updated_by"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

// CreateAssignmentRequest represents the request to create an assignment
type CreateAssignmentRequest struct {
	ClassID      string     `json:"class_id" binding:"required"`
	AssessmentID string     `json:"assessment_id" binding:"required"`
	Title        string     `json:"title" binding:"required,max=255"`
	Description  *string    `json:"description,omitempty"`
	DueDate      time.Time  `json:"due_date" binding:"required"`
	MaxScore     int        `json:"max_score" binding:"required,min=1"`
}

// UpdateAssignmentRequest represents the request to update an assignment
type UpdateAssignmentRequest struct {
	Title       *string          `json:"title,omitempty" binding:"omitempty,max=255"`
	Description *string          `json:"description,omitempty"`
	DueDate     *time.Time       `json:"due_date,omitempty"`
	MaxScore    *int             `json:"max_score,omitempty" binding:"omitempty,min=1"`
	Status      *AssignmentStatus `json:"status,omitempty" binding:"omitempty,oneof=ASSIGNED IN_PROGRESS SUBMITTED GRADED CANCELLED"`
}

// AssignmentResponse represents the assignment data returned to clients
type AssignmentResponse struct {
	ID             string           `json:"id"`
	ClassID        string           `json:"class_id"`
	ClassName      *string          `json:"class_name,omitempty"`
	AssessmentID   string           `json:"assessment_id"`
	AssessmentType *string          `json:"assessment_type,omitempty"`
	Title          string           `json:"title"`
	Description    *string          `json:"description,omitempty"`
	DueDate        string           `json:"due_date"`
	MaxScore       int              `json:"max_score"`
	Status         AssignmentStatus `json:"status"`
	CreatedAt      string           `json:"created_at"`
	UpdatedAt      string           `json:"updated_at"`
	CreatedBy      *string          `json:"created_by,omitempty"`
	CreatedByName  *string          `json:"created_by_name,omitempty"`
	UpdatedBy      *string          `json:"updated_by,omitempty"`
	UpdatedByName  *string          `json:"updated_by_name,omitempty"`
}

// ToAssignmentResponse converts Assignment to AssignmentResponse
func (a *Assignment) ToAssignmentResponse(className, assessmentType, createdByName, updatedByName string) *AssignmentResponse {
	var classNamePtr, assessmentTypePtr, descriptionPtr, createdByPtr, createdByNamePtr, updatedByPtr, updatedByNamePtr *string

	if className != "" {
		classNamePtr = &className
	}
	if assessmentType != "" {
		assessmentTypePtr = &assessmentType
	}
	if a.Description != nil {
		descriptionPtr = a.Description
	}
	if a.CreatedBy != nil {
		createdByPtr = a.CreatedBy
	}
	if createdByName != "" {
		createdByNamePtr = &createdByName
	}
	if a.UpdatedBy != nil {
		updatedByPtr = a.UpdatedBy
	}
	if updatedByName != "" {
		updatedByNamePtr = &updatedByName
	}

	return &AssignmentResponse{
		ID:             a.ID,
		ClassID:        a.ClassID,
		ClassName:      classNamePtr,
		AssessmentID:   a.AssessmentID,
		AssessmentType: assessmentTypePtr,
		Title:          a.Title,
		Description:    descriptionPtr,
		DueDate:        a.DueDate.Format(time.RFC3339),
		MaxScore:       a.MaxScore,
		Status:         AssignmentStatus(a.Status),
		CreatedAt:      a.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      a.UpdatedAt.Format(time.RFC3339),
		CreatedBy:      createdByPtr,
		CreatedByName:  createdByNamePtr,
		UpdatedBy:      updatedByPtr,
		UpdatedByName:  updatedByNamePtr,
	}
}

// IsOverdue checks if the assignment is overdue
func (a *Assignment) IsOverdue() bool {
	return time.Now().After(a.DueDate) && a.Status != string(AssignmentStatusGraded) && a.Status != string(AssignmentStatusCancelled)
}

// Validate validates the assignment entity
func (a *Assignment) Validate() error {
	if a.ID == "" {
		return fmt.Errorf("id is required")
	}
	if a.ClassID == "" {
		return fmt.Errorf("class_id is required")
	}
	if a.AssessmentID == "" {
		return fmt.Errorf("assessment_id is required")
	}
	if a.Title == "" {
		return fmt.Errorf("title is required")
	}
	if a.DueDate.IsZero() {
		return fmt.Errorf("due_date is required")
	}
	if a.MaxScore <= 0 {
		return fmt.Errorf("max_score must be greater than 0")
	}
	if a.Status == "" {
		return fmt.Errorf("status is required")
	}
	validStatuses := map[string]bool{
		string(AssignmentStatusAssigned):   true,
		string(AssignmentStatusInProgress): true,
		string(AssignmentStatusSubmitted):  true,
		string(AssignmentStatusGraded):     true,
		string(AssignmentStatusCancelled):  true,
	}
	if !validStatuses[a.Status] {
		return fmt.Errorf("invalid status: %s", a.Status)
	}
	return nil
}

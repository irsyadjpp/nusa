package domain

import (
	"fmt"
	"time"
)

// ExamResult represents an exam result
type ExamResult struct {
	ID        string     `json:"id" db:"id"`
	ExamID    string     `json:"exam_id" db:"exam_id"`
	StudentID string     `json:"student_id" db:"student_id"`
	Score     *float64   `json:"score,omitempty" db:"score"`
	Grade     *string    `json:"grade,omitempty" db:"grade"`
	Remarks   *string    `json:"remarks,omitempty" db:"remarks"`
	GradedAt  *time.Time `json:"graded_at,omitempty" db:"graded_at"`
	GradedBy  *string    `json:"graded_by,omitempty" db:"graded_by"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

// CreateExamResultRequest represents the request to create an exam result
type CreateExamResultRequest struct {
	ExamID    string   `json:"exam_id" binding:"required"`
	StudentID string   `json:"student_id" binding:"required"`
	Score     *float64 `json:"score,omitempty" binding:"omitempty,gte=0,lte=100"`
	Grade     *string  `json:"grade,omitempty" binding:"omitempty,max=10"`
	Remarks   *string  `json:"remarks,omitempty"`
}

// UpdateExamResultRequest represents the request to update an exam result
type UpdateExamResultRequest struct {
	Score    *float64 `json:"score,omitempty" binding:"omitempty,gte=0,lte=100"`
	Grade    *string  `json:"grade,omitempty" binding:"omitempty,max=10"`
	Remarks  *string  `json:"remarks,omitempty"`
	GradedBy *string  `json:"graded_by,omitempty"`
}

// ExamResultResponse represents the exam result data returned to clients
type ExamResultResponse struct {
	ID           string   `json:"id"`
	ExamID       string   `json:"exam_id"`
	ExamDate     *string  `json:"exam_date,omitempty"`
	ExamTitle    *string  `json:"exam_title,omitempty"`
	StudentID    string   `json:"student_id"`
	StudentName  *string  `json:"student_name,omitempty"`
	Score        *float64 `json:"score,omitempty"`
	Grade        *string  `json:"grade,omitempty"`
	Remarks      *string  `json:"remarks,omitempty"`
	GradedAt     *string  `json:"graded_at,omitempty"`
	GradedBy     *string  `json:"graded_by,omitempty"`
	GradedByName *string  `json:"graded_by_name,omitempty"`
	CreatedAt    string   `json:"created_at"`
	UpdatedAt    string   `json:"updated_at"`
}

// ToExamResultResponse converts ExamResult to ExamResultResponse
func (er *ExamResult) ToExamResultResponse(examDate, examTitle, studentName, gradedByName string) *ExamResultResponse {
	var examDatePtr, examTitlePtr, studentNamePtr, gradePtr, remarksPtr, gradedAtPtr, gradedByPtr, gradedByNamePtr *string

	if examDate != "" {
		examDatePtr = &examDate
	}
	if examTitle != "" {
		examTitlePtr = &examTitle
	}
	if studentName != "" {
		studentNamePtr = &studentName
	}
	if er.Grade != nil {
		gradePtr = er.Grade
	}
	if er.Remarks != nil {
		remarksPtr = er.Remarks
	}
	if er.GradedAt != nil {
		gradedAtStr := er.GradedAt.Format(time.RFC3339)
		gradedAtPtr = &gradedAtStr
	}
	if er.GradedBy != nil {
		gradedByPtr = er.GradedBy
	}
	if gradedByName != "" {
		gradedByNamePtr = &gradedByName
	}

	return &ExamResultResponse{
		ID:           er.ID,
		ExamID:       er.ExamID,
		ExamDate:     examDatePtr,
		ExamTitle:    examTitlePtr,
		StudentID:    er.StudentID,
		StudentName:  studentNamePtr,
		Score:        er.Score,
		Grade:        gradePtr,
		Remarks:      remarksPtr,
		GradedAt:     gradedAtPtr,
		GradedBy:     gradedByPtr,
		GradedByName: gradedByNamePtr,
		CreatedAt:    er.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    er.UpdatedAt.Format(time.RFC3339),
	}
}

// IsGraded checks if the exam result has been graded
func (er *ExamResult) IsGraded() bool {
	return er.GradedAt != nil && er.GradedBy != nil
}

// MarkAsGraded marks the exam result as graded
func (er *ExamResult) MarkAsGraded(gradedBy string) {
	now := time.Now()
	er.GradedAt = &now
	er.GradedBy = &gradedBy
	er.UpdatedAt = now
}

// Validate validates the exam result entity
func (er *ExamResult) Validate() error {
	if er.ID == "" {
		return fmt.Errorf("id is required")
	}
	if er.ExamID == "" {
		return fmt.Errorf("exam_id is required")
	}
	if er.StudentID == "" {
		return fmt.Errorf("student_id is required")
	}
	if er.Score != nil {
		if *er.Score < 0 || *er.Score > 100 {
			return fmt.Errorf("score must be between 0 and 100")
		}
	}
	return nil
}

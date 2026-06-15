package domain

import (
	"fmt"
	"time"
)

// AttendanceRecord represents a student attendance record
type AttendanceRecord struct {
	ID         string     `json:"id" db:"id"`
	ClassID    string     `json:"class_id" db:"class_id"`
	StudentID  string     `json:"student_id" db:"student_id"`
	Date       time.Time  `json:"date" db:"date"`
	Status     string     `json:"status" db:"status"`
	Notes      *string    `json:"notes,omitempty" db:"notes"`
	RecordedBy string     `json:"recorded_by" db:"recorded_by"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

// AttendanceStatus represents the status of attendance
type AttendanceStatus string

const (
	AttendanceStatusPresent AttendanceStatus = "PRESENT"
	AttendanceStatusAbsent  AttendanceStatus = "ABSENT"
	AttendanceStatusLate    AttendanceStatus = "LATE"
	AttendanceStatusExcused AttendanceStatus = "EXCUSED"
)

// CreateAttendanceRequest represents the request to create an attendance record
type CreateAttendanceRequest struct {
	ClassID   string           `json:"class_id" binding:"required"`
	StudentID string           `json:"student_id" binding:"required"`
	Date      time.Time        `json:"date" binding:"required"`
	Status    AttendanceStatus `json:"status" binding:"required,oneof=PRESENT ABSENT LATE EXCUSED"`
	Notes     *string          `json:"notes,omitempty"`
}

// UpdateAttendanceRequest represents the request to update an attendance record
type UpdateAttendanceRequest struct {
	Status AttendanceStatus `json:"status" binding:"required,oneof=PRESENT ABSENT LATE EXCUSED"`
	Notes  *string          `json:"notes,omitempty"`
}

// AttendanceResponse represents the attendance data returned to clients
type AttendanceResponse struct {
	ID             string           `json:"id"`
	ClassID        string           `json:"class_id"`
	ClassName      *string          `json:"class_name,omitempty"`
	StudentID      string           `json:"student_id"`
	StudentName    *string          `json:"student_name,omitempty"`
	Date           string           `json:"date"`
	Status         AttendanceStatus `json:"status"`
	Notes          *string          `json:"notes,omitempty"`
	RecordedBy     string           `json:"recorded_by"`
	RecordedByName *string          `json:"recorded_by_name,omitempty"`
	CreatedAt      string           `json:"created_at"`
	UpdatedAt      string           `json:"updated_at"`
}

// ToAttendanceResponse converts AttendanceRecord to AttendanceResponse
func (a *AttendanceRecord) ToAttendanceResponse(className, studentName, recordedByName string) *AttendanceResponse {
	var classNamePtr, studentNamePtr, recordedByNamePtr, notesPtr *string

	if className != "" {
		classNamePtr = &className
	}
	if studentName != "" {
		studentNamePtr = &studentName
	}
	if recordedByName != "" {
		recordedByNamePtr = &recordedByName
	}
	if a.Notes != nil {
		notesPtr = a.Notes
	}

	return &AttendanceResponse{
		ID:             a.ID,
		ClassID:        a.ClassID,
		ClassName:      classNamePtr,
		StudentID:      a.StudentID,
		StudentName:    studentNamePtr,
		Date:           a.Date.Format("2006-01-02"),
		Status:         AttendanceStatus(a.Status),
		Notes:          notesPtr,
		RecordedBy:     a.RecordedBy,
		RecordedByName: recordedByNamePtr,
		CreatedAt:      a.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      a.UpdatedAt.Format(time.RFC3339),
	}
}

// Validate validates the attendance record entity
func (a *AttendanceRecord) Validate() error {
	if a.ID == "" {
		return fmt.Errorf("id is required")
	}
	if a.ClassID == "" {
		return fmt.Errorf("class_id is required")
	}
	if a.StudentID == "" {
		return fmt.Errorf("student_id is required")
	}
	if a.Date.IsZero() {
		return fmt.Errorf("date is required")
	}
	if a.Status == "" {
		return fmt.Errorf("status is required")
	}
	validStatuses := map[string]bool{
		string(AttendanceStatusPresent): true,
		string(AttendanceStatusAbsent):  true,
		string(AttendanceStatusLate):    true,
		string(AttendanceStatusExcused): true,
	}
	if !validStatuses[a.Status] {
		return fmt.Errorf("invalid status: %s", a.Status)
	}
	if a.RecordedBy == "" {
		return fmt.Errorf("recorded_by is required")
	}
	return nil
}

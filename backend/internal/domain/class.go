package domain

import (
	"fmt"
	"time"
)

// Class represents a class in the system
type Class struct {
	ID             string     `json:"id" db:"id"`
	SchoolID       string     `json:"school_id" db:"school_id"`
	AcademicYearID string     `json:"academic_year_id" db:"academic_year_id"`
	SemesterID     string     `json:"semester_id" db:"semester_id"`
	SubjectID      string     `json:"subject_id" db:"subject_id"`
	TeacherID      string     `json:"teacher_id" db:"teacher_id"`
	Name           string     `json:"name" db:"name"`
	GradeLevel     string     `json:"grade_level" db:"grade_level"`
	Room           *string    `json:"room,omitempty" db:"room"`
	MaxStudents    int        `json:"max_students" db:"max_students"`
	IsActive       bool       `json:"is_active" db:"is_active"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at" db:"updated_at"`
	CreatedBy      *string    `json:"created_by,omitempty" db:"created_by"`
	UpdatedBy      *string    `json:"updated_by,omitempty" db:"updated_by"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

// CreateClassRequest represents the request to create a new class
type CreateClassRequest struct {
	SchoolID       string  `json:"school_id" binding:"required"`
	AcademicYearID string  `json:"academic_year_id" binding:"required"`
	SemesterID     string  `json:"semester_id" binding:"required"`
	SubjectID      string  `json:"subject_id" binding:"required"`
	TeacherID      string  `json:"teacher_id" binding:"required"`
	Name           string  `json:"name" binding:"required,min=2,max=255"`
	GradeLevel     string  `json:"grade_level" binding:"required,max=50"`
	Room           *string `json:"room,omitempty" binding:"omitempty,max=100"`
	MaxStudents    int     `json:"max_students" binding:"required,min=1,max=100"`
}

// UpdateClassRequest represents the request to update a class
type UpdateClassRequest struct {
	Name        *string `json:"name,omitempty" binding:"omitempty,min=2,max=255"`
	GradeLevel  *string `json:"grade_level,omitempty" binding:"omitempty,max=50"`
	Room        *string `json:"room,omitempty" binding:"omitempty,max=100"`
	MaxStudents *int    `json:"max_students,omitempty" binding:"omitempty,min=1,max=100"`
	IsActive    *bool   `json:"is_active,omitempty"`
}

// ClassResponse represents the class data returned to clients
type ClassResponse struct {
	ID              string    `json:"id"`
	SchoolID        string    `json:"school_id"`
	SchoolName      *string   `json:"school_name,omitempty"`
	AcademicYearID  string    `json:"academic_year_id"`
	AcademicYear    *string   `json:"academic_year,omitempty"`
	SemesterID      string    `json:"semester_id"`
	SemesterName    *string   `json:"semester_name,omitempty"`
	SubjectID       string    `json:"subject_id"`
	SubjectName     *string   `json:"subject_name,omitempty"`
	TeacherID       string    `json:"teacher_id"`
	TeacherName     *string   `json:"teacher_name,omitempty"`
	Name            string    `json:"name"`
	GradeLevel      string    `json:"grade_level"`
	Room            *string   `json:"room,omitempty"`
	MaxStudents     int       `json:"max_students"`
	CurrentStudents int       `json:"current_students"`
	IsActive        bool      `json:"is_active"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// ToClassResponse converts Class to ClassResponse
func (c *Class) ToClassResponse(schoolName, academicYear, semesterName, subjectName, teacherName string, currentStudents int) *ClassResponse {
	var schoolNamePtr, academicYearPtr, semesterNamePtr, subjectNamePtr, teacherNamePtr, roomPtr *string

	if schoolName != "" {
		schoolNamePtr = &schoolName
	}
	if academicYear != "" {
		academicYearPtr = &academicYear
	}
	if semesterName != "" {
		semesterNamePtr = &semesterName
	}
	if subjectName != "" {
		subjectNamePtr = &subjectName
	}
	if teacherName != "" {
		teacherNamePtr = &teacherName
	}
	if c.Room != nil {
		roomPtr = c.Room
	}

	return &ClassResponse{
		ID:              c.ID,
		SchoolID:        c.SchoolID,
		SchoolName:      schoolNamePtr,
		AcademicYearID:  c.AcademicYearID,
		AcademicYear:    academicYearPtr,
		SemesterID:      c.SemesterID,
		SemesterName:    semesterNamePtr,
		SubjectID:       c.SubjectID,
		SubjectName:     subjectNamePtr,
		TeacherID:       c.TeacherID,
		TeacherName:     teacherNamePtr,
		Name:            c.Name,
		GradeLevel:      c.GradeLevel,
		Room:            roomPtr,
		MaxStudents:     c.MaxStudents,
		CurrentStudents: currentStudents,
		IsActive:        c.IsActive,
		CreatedAt:       c.CreatedAt,
		UpdatedAt:       c.UpdatedAt,
	}
}

// Validate validates the class entity
func (c *Class) Validate() error {
	if c.ID == "" {
		return fmt.Errorf("id is required")
	}
	if c.SchoolID == "" {
		return fmt.Errorf("school_id is required")
	}
	if c.AcademicYearID == "" {
		return fmt.Errorf("academic_year_id is required")
	}
	if c.SemesterID == "" {
		return fmt.Errorf("semester_id is required")
	}
	if c.SubjectID == "" {
		return fmt.Errorf("subject_id is required")
	}
	if c.TeacherID == "" {
		return fmt.Errorf("teacher_id is required")
	}
	if c.Name == "" {
		return fmt.Errorf("name is required")
	}
	if c.GradeLevel == "" {
		return fmt.Errorf("grade_level is required")
	}
	if c.MaxStudents < 1 || c.MaxStudents > 100 {
		return fmt.Errorf("max_students must be between 1 and 100")
	}
	return nil
}

// ClassEnrollment represents a student enrollment in a class
type ClassEnrollment struct {
	ID             string     `json:"id" db:"id"`
	ClassID        string     `json:"class_id" db:"class_id"`
	StudentID      string     `json:"student_id" db:"student_id"`
	EnrollmentDate time.Time  `json:"enrollment_date" db:"enrollment_date"`
	Status         string     `json:"status" db:"status"`
	Notes          *string    `json:"notes,omitempty" db:"notes"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

// EnrollmentStatus represents the status of an enrollment
type EnrollmentStatus string

const (
	EnrollmentStatusActive    EnrollmentStatus = "ACTIVE"
	EnrollmentStatusInactive  EnrollmentStatus = "INACTIVE"
	EnrollmentStatusWithdrawn EnrollmentStatus = "WITHDRAWN"
	EnrollmentStatusCompleted EnrollmentStatus = "COMPLETED"
)

// CreateClassEnrollmentRequest represents the request to enroll a student in a class
type CreateClassEnrollmentRequest struct {
	ClassID   string  `json:"class_id" binding:"required"`
	StudentID string  `json:"student_id" binding:"required"`
	Notes     *string `json:"notes,omitempty"`
}

// UpdateClassEnrollmentRequest represents the request to update an enrollment
type UpdateClassEnrollmentRequest struct {
	Status EnrollmentStatus `json:"status" binding:"required,oneof=ACTIVE INACTIVE WITHDRAWN COMPLETED"`
	Notes  *string          `json:"notes,omitempty"`
}

// ClassEnrollmentResponse represents the enrollment data returned to clients
type ClassEnrollmentResponse struct {
	ID             string           `json:"id"`
	ClassID        string           `json:"class_id"`
	ClassName      *string          `json:"class_name,omitempty"`
	StudentID      string           `json:"student_id"`
	StudentName    *string          `json:"student_name,omitempty"`
	EnrollmentDate time.Time        `json:"enrollment_date"`
	Status         EnrollmentStatus `json:"status"`
	Notes          *string          `json:"notes,omitempty"`
	CreatedAt      time.Time        `json:"created_at"`
	UpdatedAt      time.Time        `json:"updated_at"`
}

// ToClassEnrollmentResponse converts ClassEnrollment to ClassEnrollmentResponse
func (ce *ClassEnrollment) ToClassEnrollmentResponse(className, studentName string) *ClassEnrollmentResponse {
	var classNamePtr, studentNamePtr, notesPtr *string

	if className != "" {
		classNamePtr = &className
	}
	if studentName != "" {
		studentNamePtr = &studentName
	}
	if ce.Notes != nil {
		notesPtr = ce.Notes
	}

	return &ClassEnrollmentResponse{
		ID:             ce.ID,
		ClassID:        ce.ClassID,
		ClassName:      classNamePtr,
		StudentID:      ce.StudentID,
		StudentName:    studentNamePtr,
		EnrollmentDate: ce.EnrollmentDate,
		Status:         EnrollmentStatus(ce.Status),
		Notes:          notesPtr,
		CreatedAt:      ce.CreatedAt,
		UpdatedAt:      ce.UpdatedAt,
	}
}

// Validate validates the class enrollment entity
func (ce *ClassEnrollment) Validate() error {
	if ce.ID == "" {
		return fmt.Errorf("id is required")
	}
	if ce.ClassID == "" {
		return fmt.Errorf("class_id is required")
	}
	if ce.StudentID == "" {
		return fmt.Errorf("student_id is required")
	}
	if ce.Status == "" {
		return fmt.Errorf("status is required")
	}
	return nil
}

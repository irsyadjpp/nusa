package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// SemesterType represents the type of semester
type SemesterType string

const (
	SemesterTypeGanjil SemesterType = "GANJIL" // Semester 1 (Odd semester)
	SemesterTypeGenap  SemesterType = "GENAP"  // Semester 2 (Even semester)
)

// SemesterStatus represents the status of a semester
type SemesterStatus string

const (
	SemesterStatusActive   SemesterStatus = "ACTIVE"
	SemesterStatusInactive SemesterStatus = "INACTIVE"
)

// Semester represents a semester within an academic year
type Semester struct {
	ID              string          `json:"id" db:"id"`
	AcademicYearID  string          `json:"academic_year_id" db:"academic_year_id"`
	Type            SemesterType    `json:"type" db:"type"`
	Name            string          `json:"name" db:"name"`
	StartDate       time.Time       `json:"start_date" db:"start_date"`
	EndDate         time.Time       `json:"end_date" db:"end_date"`
	Status          SemesterStatus  `json:"status" db:"status"`
	SequenceNumber  int             `json:"sequence_number" db:"sequence_number"`
	CreatedBy       string          `json:"created_by" db:"created_by"`
	CreatedAt       time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at" db:"updated_at"`
}

// NewSemester creates a new Semester entity
func NewSemester(academicYearID, name string, semType SemesterType, startDate, endDate time.Time, sequenceNumber int, createdBy string) (*Semester, error) {
	if academicYearID == "" {
		return nil, errors.New("academic year ID is required")
	}
	if name == "" {
		return nil, errors.New("semester name is required")
	}
	if startDate.After(endDate) {
		return nil, errors.New("start date must be before end date")
	}
	
	// Validate semester type
	if semType != SemesterTypeGanjil && semType != SemesterTypeGenap {
		return nil, errors.New("invalid semester type, must be GANJIL or GENAP")
	}

	return &Semester{
		ID:             uuid.New().String(),
		AcademicYearID: academicYearID,
		Type:           semType,
		Name:           name,
		StartDate:      startDate,
		EndDate:        endDate,
		Status:         SemesterStatusActive,
		SequenceNumber: sequenceNumber,
		CreatedBy:      createdBy,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}, nil
}

// Validate performs business rule validation on the Semester
func (s *Semester) Validate() error {
	if s.AcademicYearID == "" {
		return errors.New("academic year ID is required")
	}
	if s.Name == "" {
		return errors.New("semester name is required")
	}
	if s.StartDate.After(s.EndDate) {
		return errors.New("start date must be before end date")
	}
	if s.Type != SemesterTypeGanjil && s.Type != SemesterTypeGenap {
		return errors.New("invalid semester type")
	}
	if s.Status != SemesterStatusActive && s.Status != SemesterStatusInactive {
		return errors.New("invalid semester status")
	}
	if s.SequenceNumber < 1 || s.SequenceNumber > 2 {
		return errors.New("sequence number must be 1 or 2")
	}
	return nil
}

// Activate sets the semester to ACTIVE status
func (s *Semester) Activate() {
	s.Status = SemesterStatusActive
	s.UpdatedAt = time.Now()
}

// Deactivate sets the semester to INACTIVE status
func (s *Semester) Deactivate() {
	s.Status = SemesterStatusInactive
	s.UpdatedAt = time.Now()
}

// IsActive checks if the semester is currently active
func (s *Semester) IsActive() bool {
	return s.Status == SemesterStatusActive
}

// ContainsDate checks if a given date falls within the semester range
func (s *Semester) ContainsDate(date time.Time) bool {
	return (date.Equal(s.StartDate) || date.After(s.StartDate)) && date.Before(s.EndDate)
}

// OverlapsWith checks if this semester overlaps with another
func (s *Semester) OverlapsWith(other *Semester) bool {
	if s.AcademicYearID != other.AcademicYearID {
		return false
	}
	
	// Check date overlap
	overlapStart := s.StartDate
	if other.StartDate.After(s.StartDate) {
		overlapStart = other.StartDate
	}
	
	overlapEnd := s.EndDate
	if other.EndDate.Before(s.EndDate) {
		overlapEnd = other.EndDate
	}
	
	return overlapStart.Before(overlapEnd)
}

// IsGanjil checks if this is a Ganjil (odd) semester
func (s *Semester) IsGanjil() bool {
	return s.Type == SemesterTypeGanjil
}

// IsGenap checks if this is a Genap (even) semester
func (s *Semester) IsGenap() bool {
	return s.Type == SemesterTypeGenap
}

// CreateSemesterRequest represents the request to create a semester
type CreateSemesterRequest struct {
	AcademicYearID string       `json:"academic_year_id" binding:"required"`
	Type           SemesterType `json:"type" binding:"required"`
	Name           string       `json:"name" binding:"required,min=1,max=100"`
	StartDate      time.Time    `json:"start_date" binding:"required"`
	EndDate        time.Time    `json:"end_date" binding:"required"`
	SequenceNumber int          `json:"sequence_number" binding:"required,min=1,max=2"`
}

// Validate performs validation on the request
func (r *CreateSemesterRequest) Validate() error {
	if r.AcademicYearID == "" {
		return errors.New("academic_year_id is required")
	}
	if r.Type != SemesterTypeGanjil && r.Type != SemesterTypeGenap {
		return errors.New("type must be GANJIL or GENAP")
	}
	if r.Name == "" {
		return errors.New("name is required")
	}
	if len(r.Name) > 100 {
		return errors.New("name must be less than 100 characters")
	}
	if r.StartDate.IsZero() {
		return errors.New("start_date is required")
	}
	if r.EndDate.IsZero() {
		return errors.New("end_date is required")
	}
	if r.StartDate.After(r.EndDate) {
		return errors.New("start_date must be before end_date")
	}
	if r.SequenceNumber < 1 || r.SequenceNumber > 2 {
		return errors.New("sequence_number must be 1 or 2")
	}
	return nil
}

// UpdateSemesterRequest represents the request to update a semester
type UpdateSemesterRequest struct {
	Name      *string    `json:"name,omitempty"`
	StartDate *time.Time `json:"start_date,omitempty"`
	EndDate   *time.Time `json:"end_date,omitempty"`
	Status    *SemesterStatus `json:"status,omitempty"`
}

// Validate performs validation on the update request
func (r *UpdateSemesterRequest) Validate() error {
	if r.Name != nil {
		if *r.Name == "" {
			return errors.New("name cannot be empty")
		}
		if len(*r.Name) > 100 {
			return errors.New("name must be less than 100 characters")
		}
	}
	if r.StartDate != nil && r.EndDate != nil {
		if r.StartDate.After(*r.EndDate) {
			return errors.New("start_date must be before end date")
		}
	}
	if r.Status != nil {
		if *r.Status != SemesterStatusActive && *r.Status != SemesterStatusInactive {
			return errors.New("status must be ACTIVE or INACTIVE")
		}
	}
	return nil
}
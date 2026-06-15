package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// AcademicYearStatus represents the status of an academic year
type AcademicYearStatus string

const (
	AcademicYearStatusDraft    AcademicYearStatus = "DRAFT"
	AcademicYearStatusActive   AcademicYearStatus = "ACTIVE"
	AcademicYearStatusArchived AcademicYearStatus = "ARCHIVED"
)

// AcademicYear represents an academic year in the school calendar
// Workflow: DRAFT → ACTIVE → ARCHIVED (simplified, no approval workflow)
type AcademicYear struct {
	ID        string             `json:"id" db:"id"`
	SchoolID  string             `json:"school_id" db:"school_id"`
	Name      string             `json:"name" db:"name"`
	StartDate time.Time          `json:"start_date" db:"start_date"`
	EndDate   time.Time          `json:"end_date" db:"end_date"`
	Status    AcademicYearStatus `json:"status" db:"status"`
	CreatedBy string             `json:"created_by" db:"created_by"`
	CreatedAt time.Time          `json:"created_at" db:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" db:"updated_at"`
}

// NewAcademicYear creates a new AcademicYear entity
func NewAcademicYear(schoolID, name string, startDate, endDate time.Time, createdBy string) (*AcademicYear, error) {
	if name == "" {
		return nil, errors.New("academic year name is required")
	}
	if startDate.After(endDate) {
		return nil, errors.New("start date must be before end date")
	}
	if startDate.Before(time.Now().AddDate(0, 0, 30)) {
		return nil, errors.New("academic year must be created at least 30 days in advance")
	}

	return &AcademicYear{
		ID:        uuid.New().String(),
		SchoolID:  schoolID,
		Name:      name,
		StartDate: startDate,
		EndDate:   endDate,
		Status:    AcademicYearStatusDraft,
		CreatedBy: createdBy,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

// Validate performs business rule validation on the AcademicYear
func (ay *AcademicYear) Validate() error {
	// BR-002: Academic Year Non-Overlap (checked by service, not domain)
	// BR-003: Academic Year Lead Time (validated in NewAcademicYear)

	if ay.Name == "" {
		return errors.New("academic year name is required")
	}
	if ay.StartDate.After(ay.EndDate) {
		return errors.New("start date must be before end date")
	}
	if ay.StartDate.Before(time.Now().AddDate(0, 0, 30)) {
		return errors.New("academic year must be created at least 30 days in advance")
	}

	// Validate status
	switch ay.Status {
	case AcademicYearStatusDraft, AcademicYearStatusActive, AcademicYearStatusArchived:
		// Valid statuses
	default:
		return errors.New("invalid academic year status")
	}

	return nil
}

// Activate transitions the academic year to ACTIVE status
func (ay *AcademicYear) Activate() error {
	if ay.Status != AcademicYearStatusDraft {
		return errors.New("only DRAFT academic years can be activated")
	}

	// BR-001: Only one academic year can be active at a time (checked by service)

	ay.Status = AcademicYearStatusActive
	ay.UpdatedAt = time.Now()
	return nil
}

// Archive transitions the academic year to ARCHIVED status
func (ay *AcademicYear) Archive() error {
	if ay.Status != AcademicYearStatusActive {
		return errors.New("only ACTIVE academic years can be archived")
	}

	ay.Status = AcademicYearStatusArchived
	ay.UpdatedAt = time.Now()
	return nil
}

// IsActive checks if the academic year is currently active
func (ay *AcademicYear) IsActive() bool {
	return ay.Status == AcademicYearStatusActive
}

// IsDraft checks if the academic year is in DRAFT status
func (ay *AcademicYear) IsDraft() bool {
	return ay.Status == AcademicYearStatusDraft
}

// CanBeModified checks if the academic year can be modified
func (ay *AcademicYear) CanBeModified() bool {
	return ay.Status == AcademicYearStatusDraft
}

// ContainsDate checks if a given date falls within the academic year range
func (ay *AcademicYear) ContainsDate(date time.Time) bool {
	return (date.Equal(ay.StartDate) || date.After(ay.StartDate)) && date.Before(ay.EndDate)
}

// OverlapsWith checks if this academic year overlaps with another
func (ay *AcademicYear) OverlapsWith(other *AcademicYear) bool {
	if ay.SchoolID != other.SchoolID {
		return false
	}

	// Check date overlap
	overlapStart := ay.StartDate
	if other.StartDate.After(ay.StartDate) {
		overlapStart = other.StartDate
	}

	overlapEnd := ay.EndDate
	if other.EndDate.Before(ay.EndDate) {
		overlapEnd = other.EndDate
	}

	return overlapStart.Before(overlapEnd)
}

// CreateAcademicYearRequest represents the request to create an academic year
type CreateAcademicYearRequest struct {
	SchoolID    string    `json:"school_id" binding:"required"`
	Name        string    `json:"name" binding:"required,min=1,max=100"`
	StartDate   time.Time `json:"start_date" binding:"required"`
	EndDate     time.Time `json:"end_date" binding:"required"`
	Description string    `json:"description,omitempty" binding:"max=500"`
}

// Validate performs validation on the request
func (r *CreateAcademicYearRequest) Validate() error {
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
	if r.StartDate.Before(time.Now().AddDate(0, 0, 30)) {
		return errors.New("start_date must be at least 30 days in the future")
	}
	if len(r.Description) > 500 {
		return errors.New("description must be less than 500 characters")
	}
	return nil
}

// UpdateAcademicYearRequest represents the request to update an academic year
type UpdateAcademicYearRequest struct {
	Name        *string    `json:"name,omitempty"`
	StartDate   *time.Time `json:"start_date,omitempty"`
	EndDate     *time.Time `json:"end_date,omitempty"`
	Description *string    `json:"description,omitempty"`
}

// Validate performs validation on the update request
func (r *UpdateAcademicYearRequest) Validate() error {
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
	if r.StartDate != nil {
		if r.StartDate.Before(time.Now().AddDate(0, 0, 30)) {
			return errors.New("start_date must be at least 30 days in the future")
		}
	}
	if r.Description != nil && len(*r.Description) > 500 {
		return errors.New("description must be less than 500 characters")
	}
	return nil
}

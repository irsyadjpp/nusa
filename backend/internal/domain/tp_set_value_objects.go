package domain

import (
	"errors"
)

// CPCode represents the CP Code value object
type CPCode struct {
	value string
}

// NewCPCode creates a new CPCode value object
func NewCPCode(code string) (*CPCode, error) {
	if code == "" {
		return nil, errors.New("CP code cannot be empty")
	}
	return &CPCode{value: code}, nil
}

// String returns the string representation
func (c *CPCode) String() string {
	return c.value
}

// Equals checks equality
func (c *CPCode) Equals(other *CPCode) bool {
	if c == nil || other == nil {
		return false
	}
	return c.value == other.value
}

// CPText represents the CP Text value object
type CPText struct {
	value string
}

// NewCPText creates a new CPText value object
func NewCPText(text string) (*CPText, error) {
	if text == "" {
		return nil, errors.New("CP text cannot be empty")
	}
	return &CPText{value: text}, nil
}

// String returns the string representation
func (c *CPText) String() string {
	return c.value
}

// Equals checks equality
func (c *CPText) Equals(other *CPText) bool {
	if c == nil || other == nil {
		return false
	}
	return c.value == other.value
}

// LearningObjective represents a learning objective value object
type LearningObjective struct {
	text string
}

// NewLearningObjective creates a new LearningObjective value object
func NewLearningObjective(text string) (*LearningObjective, error) {
	if text == "" {
		return nil, errors.New("learning objective cannot be empty")
	}
	return &LearningObjective{text: text}, nil
}

// String returns the string representation
func (l *LearningObjective) String() string {
	return l.text
}

// Equals checks equality
func (l *LearningObjective) Equals(other *LearningObjective) bool {
	if l == nil || other == nil {
		return false
	}
	return l.text == other.text
}

// TimeAllocation represents time allocation value object
type TimeAllocation struct {
	weeks   int
	hours   int
	minutes int
}

// NewTimeAllocation creates a new TimeAllocation value object
func NewTimeAllocation(weeks, hours, minutes int) (*TimeAllocation, error) {
	if weeks < 0 {
		return nil, errors.New("weeks cannot be negative")
	}
	if hours < 0 {
		return nil, errors.New("hours cannot be negative")
	}
	if minutes < 0 {
		return nil, errors.New("minutes cannot be negative")
	}
	if minutes >= 60 {
		return nil, errors.New("minutes must be less than 60")
	}
	return &TimeAllocation{
		weeks:   weeks,
		hours:   hours,
		minutes: minutes,
	}, nil
}

// GetWeeks returns the weeks
func (t *TimeAllocation) GetWeeks() int {
	return t.weeks
}

// GetHours returns the hours
func (t *TimeAllocation) GetHours() int {
	return t.hours
}

// GetMinutes returns the minutes
func (t *TimeAllocation) GetMinutes() int {
	return t.minutes
}

// Equals checks equality
func (t *TimeAllocation) Equals(other *TimeAllocation) bool {
	if t == nil || other == nil {
		return false
	}
	return t.weeks == other.weeks &&
		t.hours == other.hours &&
		t.minutes == other.minutes
}

// SuccessCriteria represents success criteria value object
type SuccessCriteria struct {
	criteria interface{}
}

// NewSuccessCriteria creates a new SuccessCriteria value object
func NewSuccessCriteria(criteria interface{}) (*SuccessCriteria, error) {
	if criteria == nil {
		return nil, errors.New("success criteria cannot be nil")
	}
	return &SuccessCriteria{criteria: criteria}, nil
}

// GetCriteria returns the criteria
func (s *SuccessCriteria) GetCriteria() interface{} {
	return s.criteria
}

// Equals checks equality
func (s *SuccessCriteria) Equals(other *SuccessCriteria) bool {
	if s == nil || other == nil {
		return false
	}
	// For interface{}, we can only do shallow comparison
	return s.criteria == other.criteria
}

// GenerationReason represents the generation reason value object
type GenerationReason struct {
	reason string
}

// NewGenerationReason creates a new GenerationReason value object
func NewGenerationReason(reason string) (*GenerationReason, error) {
	if reason == "" {
		return nil, errors.New("generation reason cannot be empty")
	}
	if len(reason) > 500 {
		return nil, errors.New("generation reason cannot exceed 500 characters")
	}
	return &GenerationReason{reason: reason}, nil
}

// String returns the string representation
func (g *GenerationReason) String() string {
	return g.reason
}

// Equals checks equality
func (g *GenerationReason) Equals(other *GenerationReason) bool {
	if g == nil || other == nil {
		return false
	}
	return g.reason == other.reason
}

// TPSetID represents the TPSet ID value object
type TPSetID struct {
	value string
}

// NewTPSetID creates a new TPSetID value object
func NewTPSetID(id string) (*TPSetID, error) {
	if id == "" {
		return nil, errors.New("TPSet ID cannot be empty")
	}
	return &TPSetID{value: id}, nil
}

// String returns the string representation
func (t *TPSetID) String() string {
	return t.value
}

// Equals checks equality
func (t *TPSetID) Equals(other *TPSetID) bool {
	if t == nil || other == nil {
		return false
	}
	return t.value == other.value
}

// SchoolID represents the School ID value object
type SchoolID struct {
	value string
}

// NewSchoolID creates a new SchoolID value object
func NewSchoolID(id string) (*SchoolID, error) {
	if id == "" {
		return nil, errors.New("School ID cannot be empty")
	}
	return &SchoolID{value: id}, nil
}

// String returns the string representation
func (s *SchoolID) String() string {
	return s.value
}

// Equals checks equality
func (s *SchoolID) Equals(other *SchoolID) bool {
	if s == nil || other == nil {
		return false
	}
	return s.value == other.value
}

// UserID represents the User ID value object
type UserID struct {
	value string
}

// NewUserID creates a new UserID value object
func NewUserID(id string) (*UserID, error) {
	if id == "" {
		return nil, errors.New("User ID cannot be empty")
	}
	return &UserID{value: id}, nil
}

// String returns the string representation
func (u *UserID) String() string {
	return u.value
}

// Equals checks equality
func (u *UserID) Equals(other *UserID) bool {
	if u == nil || other == nil {
		return false
	}
	return u.value == other.value
}

// WorkflowStatus is already defined in curriculum.go
// Using the existing definition to avoid redeclaration

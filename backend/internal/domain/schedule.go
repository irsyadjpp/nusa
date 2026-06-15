package domain

import (
	"fmt"
	"time"
)

// Schedule represents a class schedule
type Schedule struct {
	ID        string     `json:"id" db:"id"`
	ClassID   string     `json:"class_id" db:"class_id"`
	DayOfWeek int        `json:"day_of_week" db:"day_of_week"`
	StartTime string     `json:"start_time" db:"start_time"`
	EndTime   string     `json:"end_time" db:"end_time"`
	Room      *string    `json:"room,omitempty" db:"room"`
	IsActive  bool       `json:"is_active" db:"is_active"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	CreatedBy *string    `json:"created_by,omitempty" db:"created_by"`
	UpdatedBy *string    `json:"updated_by,omitempty" db:"updated_by"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

// CreateScheduleRequest represents the request to create a schedule
type CreateScheduleRequest struct {
	ClassID   string  `json:"class_id" binding:"required"`
	DayOfWeek int     `json:"day_of_week" binding:"required,min=1,max=7"`
	StartTime string  `json:"start_time" binding:"required"`
	EndTime   string  `json:"end_time" binding:"required"`
	Room      *string `json:"room,omitempty" binding:"omitempty,max=100"`
}

// UpdateScheduleRequest represents the request to update a schedule
type UpdateScheduleRequest struct {
	DayOfWeek *int    `json:"day_of_week,omitempty" binding:"omitempty,min=1,max=7"`
	StartTime *string `json:"start_time,omitempty" binding:"omitempty"`
	EndTime   *string `json:"end_time,omitempty" binding:"omitempty"`
	Room      *string `json:"room,omitempty" binding:"omitempty,max=100"`
	IsActive  *bool   `json:"is_active,omitempty"`
}

// ScheduleResponse represents the schedule data returned to clients
type ScheduleResponse struct {
	ID        string  `json:"id"`
	ClassID   string  `json:"class_id"`
	ClassName *string `json:"class_name,omitempty"`
	DayOfWeek int     `json:"day_of_week"`
	DayName   string  `json:"day_name"`
	StartTime string  `json:"start_time"`
	EndTime   string  `json:"end_time"`
	Room      *string `json:"room,omitempty"`
	IsActive  bool    `json:"is_active"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

// ToScheduleResponse converts Schedule to ScheduleResponse
func (s *Schedule) ToScheduleResponse(className string) *ScheduleResponse {
	var classNamePtr, roomPtr *string

	if className != "" {
		classNamePtr = &className
	}
	if s.Room != nil {
		roomPtr = s.Room
	}

	return &ScheduleResponse{
		ID:        s.ID,
		ClassID:   s.ClassID,
		ClassName: classNamePtr,
		DayOfWeek: s.DayOfWeek,
		DayName:   getDayName(s.DayOfWeek),
		StartTime: s.StartTime,
		EndTime:   s.EndTime,
		Room:      roomPtr,
		IsActive:  s.IsActive,
		CreatedAt: s.CreatedAt.Format(time.RFC3339),
		UpdatedAt: s.UpdatedAt.Format(time.RFC3339),
	}
}

// getDayName returns the day name for a given day of week (1-7)
func getDayName(dayOfWeek int) string {
	days := map[int]string{
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thursday",
		5: "Friday",
		6: "Saturday",
		7: "Sunday",
	}
	return days[dayOfWeek]
}

// Validate validates the schedule entity
func (s *Schedule) Validate() error {
	if s.ID == "" {
		return fmt.Errorf("id is required")
	}
	if s.ClassID == "" {
		return fmt.Errorf("class_id is required")
	}
	if s.DayOfWeek < 1 || s.DayOfWeek > 7 {
		return fmt.Errorf("day_of_week must be between 1 and 7")
	}
	if s.StartTime == "" {
		return fmt.Errorf("start_time is required")
	}
	if s.EndTime == "" {
		return fmt.Errorf("end_time is required")
	}
	return nil
}

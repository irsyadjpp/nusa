package domain

import (
	"testing"
	"time"
)

// TestSchedule_ToScheduleResponse tests conversion to response format
func TestSchedule_ToScheduleResponse(t *testing.T) {
	now := time.Now()
	room := "Room 101"

	tests := []struct {
		name            string
		schedule       *Schedule
		className       string
		expectedClassName *string
		expectedRoom   *string
		expectedDayName string
	}{
		{
			name: "Full schedule with all related data",
			schedule: &Schedule{
				ID:        "schedule1",
				ClassID:   "class1",
				DayOfWeek: 1,
				StartTime: "08:00",
				EndTime:   "09:00",
				Room:      &room,
				IsActive:  true,
				CreatedAt: now,
				UpdatedAt: now,
			},
			className:       "Math Class",
			expectedClassName: stringPtr("Math Class"),
			expectedRoom:   &room,
			expectedDayName: "Monday",
		},
		{
			name: "Schedule without class name",
			schedule: &Schedule{
				ID:        "schedule2",
				ClassID:   "class1",
				DayOfWeek: 2,
				StartTime: "10:00",
				EndTime:   "11:00",
				Room:      nil,
				IsActive:  true,
				CreatedAt: now,
				UpdatedAt: now,
			},
			className:       "",
			expectedClassName: nil,
			expectedRoom:   nil,
			expectedDayName: "Tuesday",
		},
		{
			name: "Schedule without room",
			schedule: &Schedule{
				ID:        "schedule3",
				ClassID:   "class1",
				DayOfWeek: 3,
				StartTime: "13:00",
				EndTime:   "14:00",
				Room:      nil,
				IsActive:  false,
				CreatedAt: now,
				UpdatedAt: now,
			},
			className:       "Science Class",
			expectedClassName: stringPtr("Science Class"),
			expectedRoom:   nil,
			expectedDayName: "Wednesday",
		},
		{
			name: "Schedule for Friday",
			schedule: &Schedule{
				ID:        "schedule4",
				ClassID:   "class1",
				DayOfWeek: 5,
				StartTime: "09:00",
				EndTime:   "10:00",
				Room:      &room,
				IsActive:  true,
				CreatedAt: now,
				UpdatedAt: now,
			},
			className:       "History Class",
			expectedClassName: stringPtr("History Class"),
			expectedRoom:   &room,
			expectedDayName: "Friday",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			resp := tc.schedule.ToScheduleResponse(tc.className)

			if resp.ID != tc.schedule.ID {
				t.Errorf("expected ID %s, got %s", tc.schedule.ID, resp.ID)
			}
			if resp.ClassID != tc.schedule.ClassID {
				t.Errorf("expected ClassID %s, got %s", tc.schedule.ClassID, resp.ClassID)
			}
			if resp.DayOfWeek != tc.schedule.DayOfWeek {
				t.Errorf("expected DayOfWeek %d, got %d", tc.schedule.DayOfWeek, resp.DayOfWeek)
			}
			if resp.DayName != tc.expectedDayName {
				t.Errorf("expected DayName %s, got %s", tc.expectedDayName, resp.DayName)
			}
			if resp.StartTime != tc.schedule.StartTime {
				t.Errorf("expected StartTime %s, got %s", tc.schedule.StartTime, resp.StartTime)
			}
			if resp.EndTime != tc.schedule.EndTime {
				t.Errorf("expected EndTime %s, got %s", tc.schedule.EndTime, resp.EndTime)
			}
			if resp.IsActive != tc.schedule.IsActive {
				t.Errorf("expected IsActive %v, got %v", tc.schedule.IsActive, resp.IsActive)
			}

			// Check optional fields
			if (resp.ClassName == nil) != (tc.expectedClassName == nil) {
				t.Errorf("ClassName pointer mismatch")
			} else if resp.ClassName != nil && tc.expectedClassName != nil {
				if *resp.ClassName != *tc.expectedClassName {
					t.Errorf("expected ClassName %s, got %s", *tc.expectedClassName, *resp.ClassName)
				}
			}

			if (resp.Room == nil) != (tc.expectedRoom == nil) {
				t.Errorf("Room pointer mismatch")
			} else if resp.Room != nil && tc.expectedRoom != nil {
				if *resp.Room != *tc.expectedRoom {
					t.Errorf("expected Room %s, got %s", *tc.expectedRoom, *resp.Room)
				}
			}
		})
	}
}

// TestGetDayName tests the getDayName function
func TestGetDayName(t *testing.T) {
	tests := []struct {
		name          string
		dayOfWeek     int
		expectedDayName string
	}{
		{
			name:          "Monday",
			dayOfWeek:     1,
			expectedDayName: "Monday",
		},
		{
			name:          "Tuesday",
			dayOfWeek:     2,
			expectedDayName: "Tuesday",
		},
		{
			name:          "Wednesday",
			dayOfWeek:     3,
			expectedDayName: "Wednesday",
		},
		{
			name:          "Thursday",
			dayOfWeek:     4,
			expectedDayName: "Thursday",
		},
		{
			name:          "Friday",
			dayOfWeek:     5,
			expectedDayName: "Friday",
		},
		{
			name:          "Saturday",
			dayOfWeek:     6,
			expectedDayName: "Saturday",
		},
		{
			name:          "Sunday",
			dayOfWeek:     7,
			expectedDayName: "Sunday",
		},
		{
			name:          "Invalid day - 0",
			dayOfWeek:     0,
			expectedDayName: "",
		},
		{
			name:          "Invalid day - 8",
			dayOfWeek:     8,
			expectedDayName: "",
		},
		{
			name:          "Invalid day - negative",
			dayOfWeek:     -1,
			expectedDayName: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dayName := getDayName(tc.dayOfWeek)
			if dayName != tc.expectedDayName {
				t.Errorf("expected day name %s, got %s", tc.expectedDayName, dayName)
			}
		})
	}
}

// TestSchedule_Validate tests schedule validation
func TestSchedule_Validate(t *testing.T) {
	now := time.Now()
	room := "Room 101"

	tests := []struct {
		name        string
		schedule    *Schedule
		expectedErr string
	}{
		{
			name: "Valid schedule",
			schedule: &Schedule{
				ID:        "schedule1",
				ClassID:   "class1",
				DayOfWeek: 1,
				StartTime: "08:00",
				EndTime:   "09:00",
				Room:      &room,
				IsActive:  true,
				CreatedAt: now,
				UpdatedAt: now,
			},
			expectedErr: "",
		},
		{
			name: "Valid schedule without room",
			schedule: &Schedule{
				ID:        "schedule2",
				ClassID:   "class1",
				DayOfWeek: 5,
				StartTime: "10:00",
				EndTime:   "11:00",
				Room:      nil,
				IsActive:  true,
				CreatedAt: now,
				UpdatedAt: now,
			},
			expectedErr: "",
		},
		{
			name: "Valid schedule for all days - Monday",
			schedule: &Schedule{
				ID:        "schedule3",
				ClassID:   "class1",
				DayOfWeek: 1,
				StartTime: "08:00",
				EndTime:   "09:00",
				CreatedAt: now,
				UpdatedAt: now,
			},
			expectedErr: "",
		},
		{
			name: "Valid schedule for all days - Sunday",
			schedule: &Schedule{
				ID:        "schedule4",
				ClassID:   "class1",
				DayOfWeek: 7,
				StartTime: "08:00",
				EndTime:   "09:00",
				CreatedAt: now,
				UpdatedAt: now,
			},
			expectedErr: "",
		},
		{
			name: "Invalid - missing ID",
			schedule: &Schedule{
				ID:        "",
				ClassID:   "class1",
				DayOfWeek: 1,
				StartTime: "08:00",
				EndTime:   "09:00",
				CreatedAt: now,
				UpdatedAt: now,
			},
			expectedErr: "id is required",
		},
		{
			name: "Invalid - missing ClassID",
			schedule: &Schedule{
				ID:        "schedule1",
				ClassID:   "",
				DayOfWeek: 1,
				StartTime: "08:00",
				EndTime:   "09:00",
				CreatedAt: now,
				UpdatedAt: now,
			},
			expectedErr: "class_id is required",
		},
		{
			name: "Invalid - DayOfWeek below minimum (0)",
			schedule: &Schedule{
				ID:        "schedule1",
				ClassID:   "class1",
				DayOfWeek: 0,
				StartTime: "08:00",
				EndTime:   "09:00",
				CreatedAt: now,
				UpdatedAt: now,
			},
			expectedErr: "day_of_week must be between 1 and 7",
		},
		{
			name: "Invalid - DayOfWeek above maximum (8)",
			schedule: &Schedule{
				ID:        "schedule1",
				ClassID:   "class1",
				DayOfWeek: 8,
				StartTime: "08:00",
				EndTime:   "09:00",
				CreatedAt: now,
				UpdatedAt: now,
			},
			expectedErr: "day_of_week must be between 1 and 7",
		},
		{
			name: "Invalid - negative DayOfWeek",
			schedule: &Schedule{
				ID:        "schedule1",
				ClassID:   "class1",
				DayOfWeek: -1,
				StartTime: "08:00",
				EndTime:   "09:00",
				CreatedAt: now,
				UpdatedAt: now,
			},
			expectedErr: "day_of_week must be between 1 and 7",
		},
		{
			name: "Invalid - missing StartTime",
			schedule: &Schedule{
				ID:        "schedule1",
				ClassID:   "class1",
				DayOfWeek: 1,
				StartTime: "",
				EndTime:   "09:00",
				CreatedAt: now,
				UpdatedAt: now,
			},
			expectedErr: "start_time is required",
		},
		{
			name: "Invalid - missing EndTime",
			schedule: &Schedule{
				ID:        "schedule1",
				ClassID:   "class1",
				DayOfWeek: 1,
				StartTime: "08:00",
				EndTime:   "",
				CreatedAt: now,
				UpdatedAt: now,
			},
			expectedErr: "end_time is required",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.schedule.Validate()
			if tc.expectedErr == "" && err != nil {
				t.Errorf("expected no error, got: %v", err)
			}
			if tc.expectedErr != "" && err == nil {
				t.Errorf("expected error: %s, got nil", tc.expectedErr)
			}
			if tc.expectedErr != "" && err != nil && err.Error() != tc.expectedErr {
				t.Errorf("expected error: %s, got: %s", tc.expectedErr, err.Error())
			}
		})
	}
}

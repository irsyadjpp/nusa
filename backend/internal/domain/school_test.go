package domain

import (
	"testing"
	"time"
)

// TestSchool_ToSchoolStatus tests conversion of IsActive to SchoolStatus
func TestSchool_ToSchoolStatus(t *testing.T) {
	now := time.Now()
	address := "123 Main St"
	phone := "555-1234"
	email := "school@example.com"

	tests := []struct {
		name            string
		school          *School
		expectedStatus SchoolStatus
	}{
		{
			name: "Active school returns ACTIVE status",
			school: &School{
				ID:       "school1",
				Name:     "Sekolah Indonesia",
				Code:     "SCH001",
				Address:  &address,
				Phone:    &phone,
				Email:    &email,
				IsActive: true,
				CreatedAt: now,
				UpdatedAt: now,
			},
			expectedStatus: SchoolStatusActive,
		},
		{
			name: "Inactive school returns INACTIVE status",
			school: &School{
				ID:       "school2",
				Name:     "Sekolah Indonesia 2",
				Code:     "SCH002",
				Address:  nil,
				Phone:    nil,
				Email:    nil,
				IsActive: false,
				CreatedAt: now,
				UpdatedAt: now,
			},
			expectedStatus: SchoolStatusInactive,
		},
		{
			name: "School with optional fields nil but active",
			school: &School{
				ID:       "school3",
				Name:     "Sekolah Indonesia 3",
				Code:     "SCH003",
				Address:  nil,
				Phone:    nil,
				Email:    nil,
				IsActive: true,
				CreatedAt: now,
				UpdatedAt: now,
			},
			expectedStatus: SchoolStatusActive,
		},
		{
			name: "School with optional fields set but inactive",
			school: &School{
				ID:       "school4",
				Name:     "Sekolah Indonesia 4",
				Code:     "SCH004",
				Address:  &address,
				Phone:    &phone,
				Email:    &email,
				IsActive: false,
				CreatedAt: now,
				UpdatedAt: now,
			},
			expectedStatus: SchoolStatusInactive,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			status := tc.school.ToSchoolStatus()
			if status != tc.expectedStatus {
				t.Errorf("expected status %s, got %s", tc.expectedStatus, status)
			}
		})
	}
}

package domain

import (
	"testing"
	"time"
)

// TestNewAcademicYear tests the AcademicYear constructor
func TestNewAcademicYear(t *testing.T) {
	tests := []struct {
		name        string
		schoolID    string
		yearName    string
		startDate   time.Time
		endDate     time.Time
		createdBy   string
		wantError   bool
		expectedErr string
	}{
		{
			name:      "Valid academic year creation",
			schoolID:  "school-1",
			yearName:  "2026-2027",
			startDate: time.Now().AddDate(0, 0, 40),
			endDate:   time.Now().AddDate(1, 0, 40),
			createdBy: "user-1",
			wantError: false,
		},
		{
			name:      "Valid - minimum lead time (30 days)",
			schoolID:  "school-1",
			yearName:  "2026-2027",
			startDate: time.Now().AddDate(0, 0, 31),
			endDate:   time.Now().AddDate(1, 0, 31),
			createdBy: "user-1",
			wantError: false,
		},
		{
			name:        "Invalid - empty name",
			schoolID:    "school-1",
			yearName:    "",
			startDate:   time.Now().AddDate(0, 0, 40),
			endDate:     time.Now().AddDate(1, 0, 40),
			createdBy:   "user-1",
			wantError:   true,
			expectedErr: "academic year name is required",
		},
		{
			name:        "Invalid - start date after end date",
			schoolID:    "school-1",
			yearName:    "2026-2027",
			startDate:   time.Now().AddDate(1, 0, 0),
			endDate:     time.Now().AddDate(0, 0, 40),
			createdBy:   "user-1",
			wantError:   true,
			expectedErr: "start date must be before end date",
		},
		{
			name:        "Invalid - less than 30 days lead time",
			schoolID:    "school-1",
			yearName:    "2026-2027",
			startDate:   time.Now().AddDate(0, 0, 29),
			endDate:     time.Now().AddDate(1, 0, 29),
			createdBy:   "user-1",
			wantError:   true,
			expectedErr: "academic year must be created at least 30 days in advance",
		},
		{
			name:        "Invalid - exactly 29 days",
			schoolID:    "school-1",
			yearName:    "2026-2027",
			startDate:   time.Now().AddDate(0, 0, 29),
			endDate:     time.Now().AddDate(1, 0, 29),
			createdBy:   "user-1",
			wantError:   true,
			expectedErr: "academic year must be created at least 30 days in advance",
		},
		{
			name:        "Invalid - past date",
			schoolID:    "school-1",
			yearName:    "2026-2027",
			startDate:   time.Now().AddDate(-1, 0, 0),
			endDate:     time.Now().AddDate(0, 0, 40),
			createdBy:   "user-1",
			wantError:   true,
			expectedErr: "academic year must be created at least 30 days in advance",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ay, err := NewAcademicYear(tc.schoolID, tc.yearName, tc.startDate, tc.endDate, tc.createdBy)

			if tc.wantError && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tc.wantError && err != nil {
				t.Errorf("expected no error, got: %v", err)
			}

			if !tc.wantError {
				if ay == nil {
					t.Fatal("expected academic year, got nil")
				}
				if ay.ID == "" {
					t.Error("expected ID to be generated")
				}
				if ay.SchoolID != tc.schoolID {
					t.Errorf("expected school ID %s, got %s", tc.schoolID, ay.SchoolID)
				}
				if ay.Name != tc.yearName {
					t.Errorf("expected name %s, got %s", tc.yearName, ay.Name)
				}
				if !ay.StartDate.Equal(tc.startDate) {
					t.Error("start date mismatch")
				}
				if !ay.EndDate.Equal(tc.endDate) {
					t.Error("end date mismatch")
				}
				if ay.Status != AcademicYearStatusDraft {
					t.Errorf("expected status DRAFT, got %s", ay.Status)
				}
				if ay.CreatedBy != tc.createdBy {
					t.Errorf("expected created_by %s, got %s", tc.createdBy, ay.CreatedBy)
				}
				if ay.CreatedAt.IsZero() {
					t.Error("CreatedAt should not be zero")
				}
				if ay.UpdatedAt.IsZero() {
					t.Error("UpdatedAt should not be zero")
				}
			}

			if tc.wantError && err != nil {
				if tc.expectedErr != "" && !contains(err.Error(), tc.expectedErr) {
					t.Errorf("expected error to contain %s, got %s", tc.expectedErr, err.Error())
				}
			}
		})
	}
}

// TestAcademicYear_Validate tests domain validation
func TestAcademicYear_Validate(t *testing.T) {
	tests := []struct {
		name         string
		academicYear AcademicYear
		wantError    bool
		expectedErr  string
	}{
		{
			name: "Valid DRAFT academic year",
			academicYear: AcademicYear{
				ID:        "ay-1",
				SchoolID:  "school-1",
				Name:      "2026-2027",
				StartDate: time.Now().AddDate(0, 0, 40),
				EndDate:   time.Now().AddDate(1, 0, 40),
				Status:    AcademicYearStatusDraft,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantError: false,
		},
		{
			name: "Valid ACTIVE academic year",
			academicYear: AcademicYear{
				ID:        "ay-1",
				SchoolID:  "school-1",
				Name:      "2026-2027",
				StartDate: time.Now().AddDate(0, 0, 40),
				EndDate:   time.Now().AddDate(1, 0, 40),
				Status:    AcademicYearStatusActive,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantError: false,
		},
		{
			name: "Valid ARCHIVED academic year",
			academicYear: AcademicYear{
				ID:        "ay-1",
				SchoolID:  "school-1",
				Name:      "2026-2027",
				StartDate: time.Now().AddDate(0, 0, 40),
				EndDate:   time.Now().AddDate(1, 0, 40),
				Status:    AcademicYearStatusArchived,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantError: false,
		},
		{
			name: "Invalid - empty name",
			academicYear: AcademicYear{
				ID:        "ay-1",
				SchoolID:  "school-1",
				Name:      "",
				StartDate: time.Now().AddDate(0, 0, 40),
				EndDate:   time.Now().AddDate(1, 0, 40),
				Status:    AcademicYearStatusDraft,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantError:   true,
			expectedErr: "academic year name is required",
		},
		{
			name: "Invalid - start date after end date",
			academicYear: AcademicYear{
				ID:        "ay-1",
				SchoolID:  "school-1",
				Name:      "2026-2027",
				StartDate: time.Now().AddDate(1, 0, 0),
				EndDate:   time.Now().AddDate(0, 0, 40),
				Status:    AcademicYearStatusDraft,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantError:   true,
			expectedErr: "start date must be before end date",
		},
		{
			name: "Invalid - less than 30 days lead time",
			academicYear: AcademicYear{
				ID:        "ay-1",
				SchoolID:  "school-1",
				Name:      "2026-2027",
				StartDate: time.Now().AddDate(0, 0, 29),
				EndDate:   time.Now().AddDate(1, 0, 29),
				Status:    AcademicYearStatusDraft,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantError:   true,
			expectedErr: "academic year must be created at least 30 days in advance",
		},
		{
			name: "Invalid - unknown status",
			academicYear: AcademicYear{
				ID:        "ay-1",
				SchoolID:  "school-1",
				Name:      "2026-2027",
				StartDate: time.Now().AddDate(0, 0, 40),
				EndDate:   time.Now().AddDate(1, 0, 40),
				Status:    "UNKNOWN",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantError:   true,
			expectedErr: "invalid academic year status",
		},
		{
			name: "Invalid - empty status",
			academicYear: AcademicYear{
				ID:        "ay-1",
				SchoolID:  "school-1",
				Name:      "2026-2027",
				StartDate: time.Now().AddDate(0, 0, 40),
				EndDate:   time.Now().AddDate(1, 0, 40),
				Status:    "",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantError:   true,
			expectedErr: "invalid academic year status",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.academicYear.Validate()

			if tc.wantError && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tc.wantError && err != nil {
				t.Errorf("expected no error, got: %v", err)
			}

			if tc.wantError && err != nil {
				if tc.expectedErr != "" && !contains(err.Error(), tc.expectedErr) {
					t.Errorf("expected error to contain %s, got %s", tc.expectedErr, err.Error())
				}
			}
		})
	}
}

// TestAcademicYear_Activate tests activation transition
func TestAcademicYear_Activate(t *testing.T) {
	tests := []struct {
		name           string
		academicYear   AcademicYear
		wantError      bool
		expectedErr    string
		expectedStatus AcademicYearStatus
	}{
		{
			name: "Valid activation from DRAFT",
			academicYear: AcademicYear{
				ID:        "ay-1",
				SchoolID:  "school-1",
				Name:      "2026-2027",
				StartDate: time.Now().AddDate(0, 0, 40),
				EndDate:   time.Now().AddDate(1, 0, 40),
				Status:    AcademicYearStatusDraft,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantError:      false,
			expectedStatus: AcademicYearStatusActive,
		},
		{
			name: "Invalid - already ACTIVE",
			academicYear: AcademicYear{
				ID:        "ay-1",
				SchoolID:  "school-1",
				Name:      "2026-2027",
				StartDate: time.Now().AddDate(0, 0, 40),
				EndDate:   time.Now().AddDate(1, 0, 40),
				Status:    AcademicYearStatusActive,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantError:      true,
			expectedErr:    "only DRAFT academic years can be activated",
			expectedStatus: AcademicYearStatusActive,
		},
		{
			name: "Invalid - already ARCHIVED",
			academicYear: AcademicYear{
				ID:        "ay-1",
				SchoolID:  "school-1",
				Name:      "2026-2027",
				StartDate: time.Now().AddDate(0, 0, 40),
				EndDate:   time.Now().AddDate(1, 0, 40),
				Status:    AcademicYearStatusArchived,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantError:      true,
			expectedErr:    "only DRAFT academic years can be activated",
			expectedStatus: AcademicYearStatusArchived,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			originalUpdatedAt := tc.academicYear.UpdatedAt
			err := tc.academicYear.Activate()

			if tc.wantError && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tc.wantError && err != nil {
				t.Errorf("expected no error, got: %v", err)
			}

			if !tc.wantError {
				if tc.academicYear.Status != tc.expectedStatus {
					t.Errorf("expected status %s, got %s", tc.expectedStatus, tc.academicYear.Status)
				}
				if !tc.academicYear.UpdatedAt.After(originalUpdatedAt) {
					t.Error("UpdatedAt should be updated")
				}
			}

			if tc.wantError && err != nil {
				if tc.expectedErr != "" && !contains(err.Error(), tc.expectedErr) {
					t.Errorf("expected error to contain %s, got %s", tc.expectedErr, err.Error())
				}
				if tc.academicYear.Status != tc.expectedStatus {
					t.Errorf("status should not change on error, expected %s, got %s", tc.expectedStatus, tc.academicYear.Status)
				}
			}
		})
	}
}

// TestAcademicYear_Archive tests archiving transition
func TestAcademicYear_Archive(t *testing.T) {
	tests := []struct {
		name           string
		academicYear   AcademicYear
		wantError      bool
		expectedErr    string
		expectedStatus AcademicYearStatus
	}{
		{
			name: "Valid archiving from ACTIVE",
			academicYear: AcademicYear{
				ID:        "ay-1",
				SchoolID:  "school-1",
				Name:      "2026-2027",
				StartDate: time.Now().AddDate(0, 0, 40),
				EndDate:   time.Now().AddDate(1, 0, 40),
				Status:    AcademicYearStatusActive,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantError:      false,
			expectedStatus: AcademicYearStatusArchived,
		},
		{
			name: "Invalid - still DRAFT",
			academicYear: AcademicYear{
				ID:        "ay-1",
				SchoolID:  "school-1",
				Name:      "2026-2027",
				StartDate: time.Now().AddDate(0, 0, 40),
				EndDate:   time.Now().AddDate(1, 0, 40),
				Status:    AcademicYearStatusDraft,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantError:      true,
			expectedErr:    "only ACTIVE academic years can be archived",
			expectedStatus: AcademicYearStatusDraft,
		},
		{
			name: "Invalid - already ARCHIVED",
			academicYear: AcademicYear{
				ID:        "ay-1",
				SchoolID:  "school-1",
				Name:      "2026-2027",
				StartDate: time.Now().AddDate(0, 0, 40),
				EndDate:   time.Now().AddDate(1, 0, 40),
				Status:    AcademicYearStatusArchived,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantError:      true,
			expectedErr:    "only ACTIVE academic years can be archived",
			expectedStatus: AcademicYearStatusArchived,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			originalUpdatedAt := tc.academicYear.UpdatedAt
			err := tc.academicYear.Archive()

			if tc.wantError && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tc.wantError && err != nil {
				t.Errorf("expected no error, got: %v", err)
			}

			if !tc.wantError {
				if tc.academicYear.Status != tc.expectedStatus {
					t.Errorf("expected status %s, got %s", tc.expectedStatus, tc.academicYear.Status)
				}
				if !tc.academicYear.UpdatedAt.After(originalUpdatedAt) {
					t.Error("UpdatedAt should be updated")
				}
			}

			if tc.wantError && err != nil {
				if tc.expectedErr != "" && !contains(err.Error(), tc.expectedErr) {
					t.Errorf("expected error to contain %s, got %s", tc.expectedErr, err.Error())
				}
				if tc.academicYear.Status != tc.expectedStatus {
					t.Errorf("status should not change on error, expected %s, got %s", tc.expectedStatus, tc.academicYear.Status)
				}
			}
		})
	}
}

// TestAcademicYear_IsActive tests status check methods
func TestAcademicYear_StatusChecks(t *testing.T) {
	tests := []struct {
		name               string
		status             AcademicYearStatus
		expectedActive     bool
		expectedDraft      bool
		expectedModifiable bool
	}{
		{
			name:               "ACTIVE status",
			status:             AcademicYearStatusActive,
			expectedActive:     true,
			expectedDraft:      false,
			expectedModifiable: false,
		},
		{
			name:               "DRAFT status",
			status:             AcademicYearStatusDraft,
			expectedActive:     false,
			expectedDraft:      true,
			expectedModifiable: true,
		},
		{
			name:               "ARCHIVED status",
			status:             AcademicYearStatusArchived,
			expectedActive:     false,
			expectedDraft:      false,
			expectedModifiable: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ay := &AcademicYear{
				ID:        "ay-1",
				SchoolID:  "school-1",
				Name:      "2026-2027",
				StartDate: time.Now().AddDate(0, 0, 40),
				EndDate:   time.Now().AddDate(1, 0, 40),
				Status:    tc.status,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}

			if ay.IsActive() != tc.expectedActive {
				t.Errorf("IsActive: expected %v, got %v", tc.expectedActive, ay.IsActive())
			}
			if ay.IsDraft() != tc.expectedDraft {
				t.Errorf("IsDraft: expected %v, got %v", tc.expectedDraft, ay.IsDraft())
			}
			if ay.CanBeModified() != tc.expectedModifiable {
				t.Errorf("CanBeModified: expected %v, got %v", tc.expectedModifiable, ay.CanBeModified())
			}
		})
	}
}

// TestAcademicYear_ContainsDate tests date range checking
func TestAcademicYear_ContainsDate(t *testing.T) {
	startDate := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC)

	tests := []struct {
		name     string
		ay       AcademicYear
		testDate time.Time
		expected bool
	}{
		{
			name: "Date at start - should be contained",
			ay: AcademicYear{
				StartDate: startDate,
				EndDate:   endDate,
			},
			testDate: startDate,
			expected: true,
		},
		{
			name: "Date in middle - should be contained",
			ay: AcademicYear{
				StartDate: startDate,
				EndDate:   endDate,
			},
			testDate: time.Date(2026, 6, 15, 12, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name: "Date at end - should not be contained (exclusive)",
			ay: AcademicYear{
				StartDate: startDate,
				EndDate:   endDate,
			},
			testDate: endDate,
			expected: false,
		},
		{
			name: "Date before start - should not be contained",
			ay: AcademicYear{
				StartDate: startDate,
				EndDate:   endDate,
			},
			testDate: time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC),
			expected: false,
		},
		{
			name: "Date after end - should not be contained",
			ay: AcademicYear{
				StartDate: startDate,
				EndDate:   endDate,
			},
			testDate: time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name: "Date just after start - should be contained",
			ay: AcademicYear{
				StartDate: startDate,
				EndDate:   endDate,
			},
			testDate: startDate.Add(1 * time.Second),
			expected: true,
		},
		{
			name: "Date just before end - should be contained",
			ay: AcademicYear{
				StartDate: startDate,
				EndDate:   endDate,
			},
			testDate: endDate.Add(-1 * time.Second),
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.ay.ContainsDate(tc.testDate)
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

// TestAcademicYear_OverlapsWith tests overlap detection
func TestAcademicYear_OverlapsWith(t *testing.T) {
	tests := []struct {
		name     string
		ay       AcademicYear
		other    AcademicYear
		expected bool
	}{
		{
			name: "Complete overlap - same dates",
			ay: AcademicYear{
				SchoolID:  "school-1",
				StartDate: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC),
			},
			other: AcademicYear{
				SchoolID:  "school-1",
				StartDate: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC),
			},
			expected: true,
		},
		{
			name: "Partial overlap - other starts during ay",
			ay: AcademicYear{
				SchoolID:  "school-1",
				StartDate: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC),
			},
			other: AcademicYear{
				SchoolID:  "school-1",
				StartDate: time.Date(2026, 6, 1, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2027, 6, 30, 23, 59, 59, 0, time.UTC),
			},
			expected: true,
		},
		{
			name: "Partial overlap - other ends during ay",
			ay: AcademicYear{
				SchoolID:  "school-1",
				StartDate: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC),
			},
			other: AcademicYear{
				SchoolID:  "school-1",
				StartDate: time.Date(2025, 6, 1, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2026, 6, 30, 23, 59, 59, 0, time.UTC),
			},
			expected: true,
		},
		{
			name: "No overlap - other completely before",
			ay: AcademicYear{
				SchoolID:  "school-1",
				StartDate: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC),
			},
			other: AcademicYear{
				SchoolID:  "school-1",
				StartDate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC),
			},
			expected: false,
		},
		{
			name: "No overlap - other completely after",
			ay: AcademicYear{
				SchoolID:  "school-1",
				StartDate: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC),
			},
			other: AcademicYear{
				SchoolID:  "school-1",
				StartDate: time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2028, 12, 31, 23, 59, 59, 0, time.UTC),
			},
			expected: false,
		},
		{
			name: "No overlap - different school",
			ay: AcademicYear{
				SchoolID:  "school-1",
				StartDate: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC),
			},
			other: AcademicYear{
				SchoolID:  "school-2",
				StartDate: time.Date(2026, 6, 1, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2027, 6, 30, 23, 59, 59, 0, time.UTC),
			},
			expected: false,
		},
		{
			name: "Boundary case - other ends exactly when ay starts",
			ay: AcademicYear{
				SchoolID:  "school-1",
				StartDate: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC),
			},
			other: AcademicYear{
				SchoolID:  "school-1",
				StartDate: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			expected: false,
		},
		{
			name: "Boundary case - other starts exactly when ay ends",
			ay: AcademicYear{
				SchoolID:  "school-1",
				StartDate: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC),
			},
			other: AcademicYear{
				SchoolID:  "school-1",
				StartDate: time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC),
				EndDate:   time.Date(2027, 12, 31, 23, 59, 59, 0, time.UTC),
			},
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.ay.OverlapsWith(&tc.other)
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

// TestCreateAcademicYearRequest_Validate tests request validation
func TestCreateAcademicYearRequest_Validate(t *testing.T) {
	futureDate := time.Now().AddDate(0, 0, 40)
	futureEndDate := time.Now().AddDate(1, 0, 40)

	tests := []struct {
		name        string
		request     CreateAcademicYearRequest
		wantError   bool
		expectedErr string
	}{
		{
			name: "Valid request",
			request: CreateAcademicYearRequest{
				SchoolID:    "school-1",
				Name:        "2026-2027",
				StartDate:   futureDate,
				EndDate:     futureEndDate,
				Description: "Academic year for 2026-2027",
			},
			wantError: false,
		},
		{
			name: "Valid - minimal required fields",
			request: CreateAcademicYearRequest{
				SchoolID:  "school-1",
				Name:      "2026-2027",
				StartDate: futureDate,
				EndDate:   futureEndDate,
			},
			wantError: false,
		},
		{
			name: "Invalid - empty name",
			request: CreateAcademicYearRequest{
				SchoolID:  "school-1",
				Name:      "",
				StartDate: futureDate,
				EndDate:   futureEndDate,
			},
			wantError:   true,
			expectedErr: "name is required",
		},
		{
			name: "Invalid - name too long (>100 chars)",
			request: CreateAcademicYearRequest{
				SchoolID:  "school-1",
				Name:      string(make([]byte, 101)),
				StartDate: futureDate,
				EndDate:   futureEndDate,
			},
			wantError:   true,
			expectedErr: "name must be less than 100 characters",
		},
		{
			name: "Invalid - zero start date",
			request: CreateAcademicYearRequest{
				SchoolID:  "school-1",
				Name:      "2026-2027",
				StartDate: time.Time{},
				EndDate:   futureEndDate,
			},
			wantError:   true,
			expectedErr: "start_date is required",
		},
		{
			name: "Invalid - zero end date",
			request: CreateAcademicYearRequest{
				SchoolID:  "school-1",
				Name:      "2026-2027",
				StartDate: futureDate,
				EndDate:   time.Time{},
			},
			wantError:   true,
			expectedErr: "end_date is required",
		},
		{
			name: "Invalid - start date after end date",
			request: CreateAcademicYearRequest{
				SchoolID:  "school-1",
				Name:      "2026-2027",
				StartDate: futureEndDate,
				EndDate:   futureDate,
			},
			wantError: true,
		},
		{
			name: "Invalid - start date less than 30 days",
			request: CreateAcademicYearRequest{
				SchoolID:  "school-1",
				Name:      "2026-2027",
				StartDate: time.Now().AddDate(0, 0, 29),
				EndDate:   futureEndDate,
			},
			wantError:   true,
			expectedErr: "start_date must be at least 30 days in the future",
		},
		{
			name: "Invalid - description too long (>500 chars)",
			request: CreateAcademicYearRequest{
				SchoolID:    "school-1",
				Name:        "2026-2027",
				StartDate:   futureDate,
				EndDate:     futureEndDate,
				Description: string(make([]byte, 501)),
			},
			wantError:   true,
			expectedErr: "description must be less than 500 characters",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.request.Validate()

			if tc.wantError && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tc.wantError && err != nil {
				t.Errorf("expected no error, got: %v", err)
			}

			if tc.wantError && err != nil {
				if tc.expectedErr != "" && !contains(err.Error(), tc.expectedErr) {
					t.Errorf("expected error to contain %s, got %s", tc.expectedErr, err.Error())
				}
			}
		})
	}
}

// TestUpdateAcademicYearRequest_Validate tests update request validation
func TestUpdateAcademicYearRequest_Validate(t *testing.T) {
	futureDate := time.Now().AddDate(0, 0, 40)
	futureEndDate := time.Now().AddDate(1, 0, 40)

	tests := []struct {
		name        string
		request     UpdateAcademicYearRequest
		wantError   bool
		expectedErr string
	}{
		{
			name: "Valid request - all fields",
			request: UpdateAcademicYearRequest{
				Name:        makeStringPtr("2026-2027 Updated"),
				StartDate:   &futureDate,
				EndDate:     &futureEndDate,
				Description: makeStringPtr("Updated description"),
			},
			wantError: false,
		},
		{
			name: "Valid request - only name",
			request: UpdateAcademicYearRequest{
				Name:        makeStringPtr("2026-2027"),
				StartDate:   nil,
				EndDate:     nil,
				Description: nil,
			},
			wantError: false,
		},
		{
			name: "Valid request - no fields (all nil)",
			request: UpdateAcademicYearRequest{
				Name:        nil,
				StartDate:   nil,
				EndDate:     nil,
				Description: nil,
			},
			wantError: false,
		},
		{
			name: "Invalid - empty name",
			request: UpdateAcademicYearRequest{
				Name:        makeStringPtr(""),
				StartDate:   nil,
				EndDate:     nil,
				Description: nil,
			},
			wantError:   true,
			expectedErr: "name cannot be empty",
		},
		{
			name: "Invalid - name too long (>100 chars)",
			request: UpdateAcademicYearRequest{
				Name:        makeStringPtr(string(make([]byte, 101))),
				StartDate:   nil,
				EndDate:     nil,
				Description: nil,
			},
			wantError:   true,
			expectedErr: "name must be less than 100 characters",
		},
		{
			name: "Invalid - start date after end date",
			request: UpdateAcademicYearRequest{
				Name:        nil,
				StartDate:   &futureEndDate,
				EndDate:     &futureDate,
				Description: nil,
			},
			wantError:   true,
			expectedErr: "start_date must be before end date",
		},
		{
			name: "Invalid - start date less than 30 days",
			request: UpdateAcademicYearRequest{
				Name:        nil,
				StartDate:   func() *time.Time { t := time.Now().AddDate(0, 0, 29); return &t }(),
				EndDate:     nil,
				Description: nil,
			},
			wantError:   true,
			expectedErr: "start_date must be at least 30 days in the future",
		},
		{
			name: "Invalid - description too long (>500 chars)",
			request: UpdateAcademicYearRequest{
				Name:        nil,
				StartDate:   nil,
				EndDate:     nil,
				Description: makeStringPtr(string(make([]byte, 501))),
			},
			wantError:   true,
			expectedErr: "description must be less than 500 characters",
		},
		{
			name: "Valid - start date without end date",
			request: UpdateAcademicYearRequest{
				Name:        nil,
				StartDate:   &futureDate,
				EndDate:     nil,
				Description: nil,
			},
			wantError: false,
		},
		{
			name: "Valid - end date without start date",
			request: UpdateAcademicYearRequest{
				Name:        nil,
				StartDate:   nil,
				EndDate:     &futureEndDate,
				Description: nil,
			},
			wantError: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.request.Validate()

			if tc.wantError && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tc.wantError && err != nil {
				t.Errorf("expected no error, got: %v", err)
			}

			if tc.wantError && err != nil {
				if tc.expectedErr != "" && !contains(err.Error(), tc.expectedErr) {
					t.Errorf("expected error to contain %s, got %s", tc.expectedErr, err.Error())
				}
			}
		})
	}
}

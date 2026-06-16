package domain

import (
	"testing"
	"time"
)

// TestNewSemester tests the Semester constructor
func TestNewSemester(t *testing.T) {
	tests := []struct {
		name           string
		academicYearID string
		semName        string
		semType        SemesterType
		startDate      time.Time
		endDate        time.Time
		sequenceNumber int
		createdBy      string
		wantError      bool
		expectedErr    string
	}{
		{
			name:           "Valid GANJIL semester creation",
			academicYearID: "ay-1",
			semName:        "Semester Ganjil 2026-2027",
			semType:        SemesterTypeGanjil,
			startDate:      time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC),
			endDate:        time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
			sequenceNumber: 1,
			createdBy:      "user-1",
			wantError:      false,
		},
		{
			name:           "Valid GENAP semester creation",
			academicYearID: "ay-1",
			semName:        "Semester Genap 2026-2027",
			semType:        SemesterTypeGenap,
			startDate:      time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC),
			endDate:        time.Date(2027, 6, 30, 0, 0, 0, 0, time.UTC),
			sequenceNumber: 2,
			createdBy:      "user-1",
			wantError:      false,
		},
		{
			name:           "Valid - sequence number 1",
			academicYearID: "ay-1",
			semName:        "Semester 1",
			semType:        SemesterTypeGanjil,
			startDate:      time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC),
			endDate:        time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
			sequenceNumber: 1,
			createdBy:      "user-1",
			wantError:      false,
		},
		{
			name:           "Valid - sequence number 2",
			academicYearID: "ay-1",
			semName:        "Semester 2",
			semType:        SemesterTypeGenap,
			startDate:      time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC),
			endDate:        time.Date(2027, 6, 30, 0, 0, 0, 0, time.UTC),
			sequenceNumber: 2,
			createdBy:      "user-1",
			wantError:      false,
		},
		{
			name:           "Invalid - empty academic year ID",
			academicYearID: "",
			semName:        "Semester Ganjil",
			semType:        SemesterTypeGanjil,
			startDate:      time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC),
			endDate:        time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
			sequenceNumber: 1,
			createdBy:      "user-1",
			wantError:      true,
			expectedErr:    "academic year ID is required",
		},
		{
			name:           "Invalid - empty name",
			academicYearID: "ay-1",
			semName:        "",
			semType:        SemesterTypeGanjil,
			startDate:      time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC),
			endDate:        time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
			sequenceNumber: 1,
			createdBy:      "user-1",
			wantError:      true,
			expectedErr:    "semester name is required",
		},
		{
			name:           "Invalid - start date after end date",
			academicYearID: "ay-1",
			semName:        "Semester Ganjil",
			semType:        SemesterTypeGanjil,
			startDate:      time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC),
			endDate:        time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
			sequenceNumber: 1,
			createdBy:      "user-1",
			wantError:      true,
			expectedErr:    "start date must be before end date",
		},
		{
			name:           "Invalid - invalid semester type",
			academicYearID: "ay-1",
			semName:        "Semester",
			semType:        "INVALID",
			startDate:      time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC),
			endDate:        time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
			sequenceNumber: 1,
			createdBy:      "user-1",
			wantError:      true,
			expectedErr:    "invalid semester type",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			semester, err := NewSemester(tc.academicYearID, tc.semName, tc.semType, tc.startDate, tc.endDate, tc.sequenceNumber, tc.createdBy)

			if tc.wantError && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tc.wantError && err != nil {
				t.Errorf("expected no error, got: %v", err)
			}

			if !tc.wantError {
				if semester == nil {
					t.Fatal("expected semester, got nil")
				}
				if semester.ID == "" {
					t.Error("expected ID to be generated")
				}
				if semester.AcademicYearID != tc.academicYearID {
					t.Errorf("expected academic year ID %s, got %s", tc.academicYearID, semester.AcademicYearID)
				}
				if semester.Name != tc.semName {
					t.Errorf("expected name %s, got %s", tc.semName, semester.Name)
				}
				if semester.Type != tc.semType {
					t.Errorf("expected type %s, got %s", tc.semType, semester.Type)
				}
				if !semester.StartDate.Equal(tc.startDate) {
					t.Error("start date mismatch")
				}
				if !semester.EndDate.Equal(tc.endDate) {
					t.Error("end date mismatch")
				}
				if semester.Status != SemesterStatusActive {
					t.Errorf("expected status ACTIVE, got %s", semester.Status)
				}
				if semester.SequenceNumber != tc.sequenceNumber {
					t.Errorf("expected sequence number %d, got %d", tc.sequenceNumber, semester.SequenceNumber)
				}
				if semester.CreatedBy != tc.createdBy {
					t.Errorf("expected created_by %s, got %s", tc.createdBy, semester.CreatedBy)
				}
				if semester.CreatedAt.IsZero() {
					t.Error("CreatedAt should not be zero")
				}
				if semester.UpdatedAt.IsZero() {
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

// TestSemester_Validate tests domain validation
func TestSemester_Validate(t *testing.T) {
	tests := []struct {
		name        string
		semester    Semester
		wantError   bool
		expectedErr string
	}{
		{
			name: "Valid ACTIVE GANJIL semester",
			semester: Semester{
				ID:             "sem-1",
				AcademicYearID: "ay-1",
				Type:           SemesterTypeGanjil,
				Name:           "Semester Ganjil",
				StartDate:      time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC),
				EndDate:        time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
				Status:         SemesterStatusActive,
				SequenceNumber: 1,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError: false,
		},
		{
			name: "Valid INACTIVE GENAP semester",
			semester: Semester{
				ID:             "sem-1",
				AcademicYearID: "ay-1",
				Type:           SemesterTypeGenap,
				Name:           "Semester Genap",
				StartDate:      time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC),
				EndDate:        time.Date(2027, 6, 30, 0, 0, 0, 0, time.UTC),
				Status:         SemesterStatusInactive,
				SequenceNumber: 2,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError: false,
		},
		{
			name: "Invalid - empty academic year ID",
			semester: Semester{
				ID:             "sem-1",
				AcademicYearID: "",
				Type:           SemesterTypeGanjil,
				Name:           "Semester Ganjil",
				StartDate:      time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC),
				EndDate:        time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
				Status:         SemesterStatusActive,
				SequenceNumber: 1,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError:   true,
			expectedErr: "academic year ID is required",
		},
		{
			name: "Invalid - empty name",
			semester: Semester{
				ID:             "sem-1",
				AcademicYearID: "ay-1",
				Type:           SemesterTypeGanjil,
				Name:           "",
				StartDate:      time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC),
				EndDate:        time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
				Status:         SemesterStatusActive,
				SequenceNumber: 1,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError:   true,
			expectedErr: "semester name is required",
		},
		{
			name: "Invalid - start date after end date",
			semester: Semester{
				ID:             "sem-1",
				AcademicYearID: "ay-1",
				Type:           SemesterTypeGanjil,
				Name:           "Semester Ganjil",
				StartDate:      time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC),
				EndDate:        time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
				Status:         SemesterStatusActive,
				SequenceNumber: 1,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError:   true,
			expectedErr: "start date must be before end date",
		},
		{
			name: "Invalid - invalid semester type",
			semester: Semester{
				ID:             "sem-1",
				AcademicYearID: "ay-1",
				Type:           "INVALID",
				Name:           "Semester",
				StartDate:      time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC),
				EndDate:        time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
				Status:         SemesterStatusActive,
				SequenceNumber: 1,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError:   true,
			expectedErr: "invalid semester type",
		},
		{
			name: "Invalid - invalid semester status",
			semester: Semester{
				ID:             "sem-1",
				AcademicYearID: "ay-1",
				Type:           SemesterTypeGanjil,
				Name:           "Semester Ganjil",
				StartDate:      time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC),
				EndDate:        time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
				Status:         "INVALID",
				SequenceNumber: 1,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError:   true,
			expectedErr: "invalid semester status",
		},
		{
			name: "Invalid - sequence number 0",
			semester: Semester{
				ID:             "sem-1",
				AcademicYearID: "ay-1",
				Type:           SemesterTypeGanjil,
				Name:           "Semester Ganjil",
				StartDate:      time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC),
				EndDate:        time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
				Status:         SemesterStatusActive,
				SequenceNumber: 0,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError:   true,
			expectedErr: "sequence number must be 1 or 2",
		},
		{
			name: "Invalid - sequence number 3",
			semester: Semester{
				ID:             "sem-1",
				AcademicYearID: "ay-1",
				Type:           SemesterTypeGanjil,
				Name:           "Semester Ganjil",
				StartDate:      time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC),
				EndDate:        time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
				Status:         SemesterStatusActive,
				SequenceNumber: 3,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError:   true,
			expectedErr: "sequence number must be 1 or 2",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.semester.Validate()

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

// TestSemester_Activate tests the Activate method
func TestSemester_Activate(t *testing.T) {
	semester := &Semester{
		ID:             "sem-1",
		AcademicYearID: "ay-1",
		Type:           SemesterTypeGanjil,
		Name:           "Semester Ganjil",
		StartDate:      time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC),
		EndDate:        time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
		Status:         SemesterStatusInactive,
		SequenceNumber: 1,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	oldUpdatedAt := semester.UpdatedAt
	time.Sleep(1 * time.Millisecond) // Ensure time difference

	semester.Activate()

	if semester.Status != SemesterStatusActive {
		t.Errorf("expected status ACTIVE, got %s", semester.Status)
	}
	if !semester.UpdatedAt.After(oldUpdatedAt) {
		t.Error("UpdatedAt should be updated after Activate")
	}
}

// TestSemester_Deactivate tests the Deactivate method
func TestSemester_Deactivate(t *testing.T) {
	semester := &Semester{
		ID:             "sem-1",
		AcademicYearID: "ay-1",
		Type:           SemesterTypeGanjil,
		Name:           "Semester Ganjil",
		StartDate:      time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC),
		EndDate:        time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
		Status:         SemesterStatusActive,
		SequenceNumber: 1,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	oldUpdatedAt := semester.UpdatedAt
	time.Sleep(1 * time.Millisecond) // Ensure time difference

	semester.Deactivate()

	if semester.Status != SemesterStatusInactive {
		t.Errorf("expected status INACTIVE, got %s", semester.Status)
	}
	if !semester.UpdatedAt.After(oldUpdatedAt) {
		t.Error("UpdatedAt should be updated after Deactivate")
	}
}

// TestSemester_IsActive tests the IsActive method
func TestSemester_IsActive(t *testing.T) {
	tests := []struct {
		name     string
		status   SemesterStatus
		expected bool
	}{
		{
			name:     "Active semester returns true",
			status:   SemesterStatusActive,
			expected: true,
		},
		{
			name:     "Inactive semester returns false",
			status:   SemesterStatusInactive,
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			semester := &Semester{
				ID:             "sem-1",
				AcademicYearID: "ay-1",
				Type:           SemesterTypeGanjil,
				Name:           "Semester Ganjil",
				StartDate:      time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC),
				EndDate:        time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
				Status:         tc.status,
				SequenceNumber: 1,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			}

			result := semester.IsActive()
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

// TestSemester_ContainsDate tests the ContainsDate method
func TestSemester_ContainsDate(t *testing.T) {
	tests := []struct {
		name     string
		start    time.Time
		end      time.Time
		date     time.Time
		expected bool
	}{
		{
			name:     "Date within semester range",
			start:    time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
			date:     time.Date(2026, 9, 1, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "Date on start date",
			start:    time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
			date:     time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "Date before start date",
			start:    time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
			date:     time.Date(2026, 6, 1, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "Date on end date",
			start:    time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
			date:     time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "Date after end date",
			start:    time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
			date:     time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			semester := &Semester{
				ID:             "sem-1",
				AcademicYearID: "ay-1",
				Type:           SemesterTypeGanjil,
				Name:           "Semester Ganjil",
				StartDate:      tc.start,
				EndDate:        tc.end,
				Status:         SemesterStatusActive,
				SequenceNumber: 1,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			}

			result := semester.ContainsDate(tc.date)
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

// TestSemester_OverlapsWith tests the OverlapsWith method
func TestSemester_OverlapsWith(t *testing.T) {
	tests := []struct {
		name     string
		sem1     Semester
		sem2     Semester
		expected bool
	}{
		{
			name: "Overlapping semesters in same academic year",
			sem1: Semester{
				ID:             "sem-1",
				AcademicYearID: "ay-1",
				Type:           SemesterTypeGanjil,
				Name:           "Semester Ganjil",
				StartDate:      time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC),
				EndDate:        time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
				Status:         SemesterStatusActive,
				SequenceNumber: 1,
			},
			sem2: Semester{
				ID:             "sem-2",
				AcademicYearID: "ay-1",
				Type:           SemesterTypeGenap,
				Name:           "Semester Genap",
				StartDate:      time.Date(2026, 11, 1, 0, 0, 0, 0, time.UTC),
				EndDate:        time.Date(2027, 6, 30, 0, 0, 0, 0, time.UTC),
				Status:         SemesterStatusActive,
				SequenceNumber: 2,
			},
			expected: true,
		},
		{
			name: "Non-overlapping semesters in same academic year",
			sem1: Semester{
				ID:             "sem-1",
				AcademicYearID: "ay-1",
				Type:           SemesterTypeGanjil,
				Name:           "Semester Ganjil",
				StartDate:      time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC),
				EndDate:        time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
				Status:         SemesterStatusActive,
				SequenceNumber: 1,
			},
			sem2: Semester{
				ID:             "sem-2",
				AcademicYearID: "ay-1",
				Type:           SemesterTypeGenap,
				Name:           "Semester Genap",
				StartDate:      time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC),
				EndDate:        time.Date(2027, 6, 30, 0, 0, 0, 0, time.UTC),
				Status:         SemesterStatusActive,
				SequenceNumber: 2,
			},
			expected: false,
		},
		{
			name: "Semesters in different academic years do not overlap",
			sem1: Semester{
				ID:             "sem-1",
				AcademicYearID: "ay-1",
				Type:           SemesterTypeGanjil,
				Name:           "Semester Ganjil",
				StartDate:      time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC),
				EndDate:        time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
				Status:         SemesterStatusActive,
				SequenceNumber: 1,
			},
			sem2: Semester{
				ID:             "sem-2",
				AcademicYearID: "ay-2",
				Type:           SemesterTypeGanjil,
				Name:           "Semester Ganjil",
				StartDate:      time.Date(2027, 7, 15, 0, 0, 0, 0, time.UTC),
				EndDate:        time.Date(2027, 12, 31, 0, 0, 0, 0, time.UTC),
				Status:         SemesterStatusActive,
				SequenceNumber: 1,
			},
			expected: false,
		},
		{
			name: "Semesters with same start date",
			sem1: Semester{
				ID:             "sem-1",
				AcademicYearID: "ay-1",
				Type:           SemesterTypeGanjil,
				Name:           "Semester Ganjil",
				StartDate:      time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC),
				EndDate:        time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
				Status:         SemesterStatusActive,
				SequenceNumber: 1,
			},
			sem2: Semester{
				ID:             "sem-2",
				AcademicYearID: "ay-1",
				Type:           SemesterTypeGenap,
				Name:           "Semester Genap",
				StartDate:      time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC),
				EndDate:        time.Date(2026, 11, 30, 0, 0, 0, 0, time.UTC),
				Status:         SemesterStatusActive,
				SequenceNumber: 2,
			},
			expected: true,
		},
		{
			name: "Semesters with same end date",
			sem1: Semester{
				ID:             "sem-1",
				AcademicYearID: "ay-1",
				Type:           SemesterTypeGanjil,
				Name:           "Semester Ganjil",
				StartDate:      time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC),
				EndDate:        time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
				Status:         SemesterStatusActive,
				SequenceNumber: 1,
			},
			sem2: Semester{
				ID:             "sem-2",
				AcademicYearID: "ay-1",
				Type:           SemesterTypeGenap,
				Name:           "Semester Genap",
				StartDate:      time.Date(2026, 10, 1, 0, 0, 0, 0, time.UTC),
				EndDate:        time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
				Status:         SemesterStatusActive,
				SequenceNumber: 2,
			},
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.sem1.OverlapsWith(&tc.sem2)
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

// TestSemester_IsGanjil tests the IsGanjil method
func TestSemester_IsGanjil(t *testing.T) {
	tests := []struct {
		name     string
		semType  SemesterType
		expected bool
	}{
		{
			name:     "GANJIL semester returns true",
			semType:  SemesterTypeGanjil,
			expected: true,
		},
		{
			name:     "GENAP semester returns false",
			semType:  SemesterTypeGenap,
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			semester := &Semester{
				ID:             "sem-1",
				AcademicYearID: "ay-1",
				Type:           tc.semType,
				Name:           "Semester",
				StartDate:      time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC),
				EndDate:        time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
				Status:         SemesterStatusActive,
				SequenceNumber: 1,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			}

			result := semester.IsGanjil()
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

// TestSemester_IsGenap tests the IsGenap method
func TestSemester_IsGenap(t *testing.T) {
	tests := []struct {
		name     string
		semType  SemesterType
		expected bool
	}{
		{
			name:     "GENAP semester returns true",
			semType:  SemesterTypeGenap,
			expected: true,
		},
		{
			name:     "GANJIL semester returns false",
			semType:  SemesterTypeGanjil,
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			semester := &Semester{
				ID:             "sem-1",
				AcademicYearID: "ay-1",
				Type:           tc.semType,
				Name:           "Semester",
				StartDate:      time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC),
				EndDate:        time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
				Status:         SemesterStatusActive,
				SequenceNumber: 1,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			}

			result := semester.IsGenap()
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

// TestCreateSemesterRequest_Validate tests the CreateSemesterRequest Validate method
func TestCreateSemesterRequest_Validate(t *testing.T) {
	tests := []struct {
		name        string
		request     CreateSemesterRequest
		wantError   bool
		expectedErr string
	}{
		{
			name: "Valid GANJIL semester request",
			request: CreateSemesterRequest{
				AcademicYearID: "ay-1",
				Type:           SemesterTypeGanjil,
				Name:           "Semester Ganjil",
				StartDate:      time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC),
				EndDate:        time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
				SequenceNumber: 1,
			},
			wantError: false,
		},
		{
			name: "Valid GENAP semester request",
			request: CreateSemesterRequest{
				AcademicYearID: "ay-1",
				Type:           SemesterTypeGenap,
				Name:           "Semester Genap",
				StartDate:      time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC),
				EndDate:        time.Date(2027, 6, 30, 0, 0, 0, 0, time.UTC),
				SequenceNumber: 2,
			},
			wantError: false,
		},
		{
			name: "Invalid - empty academic year ID",
			request: CreateSemesterRequest{
				AcademicYearID: "",
				Type:           SemesterTypeGanjil,
				Name:           "Semester Ganjil",
				StartDate:      time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC),
				EndDate:        time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
				SequenceNumber: 1,
			},
			wantError:   true,
			expectedErr: "academic_year_id is required",
		},
		{
			name: "Invalid - invalid semester type",
			request: CreateSemesterRequest{
				AcademicYearID: "ay-1",
				Type:           "INVALID",
				Name:           "Semester",
				StartDate:      time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC),
				EndDate:        time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
				SequenceNumber: 1,
			},
			wantError:   true,
			expectedErr: "type must be GANJIL or GENAP",
		},
		{
			name: "Invalid - empty name",
			request: CreateSemesterRequest{
				AcademicYearID: "ay-1",
				Type:           SemesterTypeGanjil,
				Name:           "",
				StartDate:      time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC),
				EndDate:        time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
				SequenceNumber: 1,
			},
			wantError:   true,
			expectedErr: "name is required",
		},
		{
			name: "Invalid - name too long",
			request: CreateSemesterRequest{
				AcademicYearID: "ay-1",
				Type:           SemesterTypeGanjil,
				Name:           string(make([]byte, 101)), // 101 characters
				StartDate:      time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC),
				EndDate:        time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
				SequenceNumber: 1,
			},
			wantError:   true,
			expectedErr: "name must be less than 100 characters",
		},
		{
			name: "Invalid - zero start date",
			request: CreateSemesterRequest{
				AcademicYearID: "ay-1",
				Type:           SemesterTypeGanjil,
				Name:           "Semester Ganjil",
				StartDate:      time.Time{},
				EndDate:        time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
				SequenceNumber: 1,
			},
			wantError:   true,
			expectedErr: "start_date is required",
		},
		{
			name: "Invalid - zero end date",
			request: CreateSemesterRequest{
				AcademicYearID: "ay-1",
				Type:           SemesterTypeGanjil,
				Name:           "Semester Ganjil",
				StartDate:      time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC),
				EndDate:        time.Time{},
				SequenceNumber: 1,
			},
			wantError:   true,
			expectedErr: "end_date is required",
		},
		{
			name: "Invalid - start date after end date",
			request: CreateSemesterRequest{
				AcademicYearID: "ay-1",
				Type:           SemesterTypeGanjil,
				Name:           "Semester Ganjil",
				StartDate:      time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC),
				EndDate:        time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
				SequenceNumber: 1,
			},
			wantError:   true,
			expectedErr: "start_date must be before end_date",
		},
		{
			name: "Invalid - sequence number 0",
			request: CreateSemesterRequest{
				AcademicYearID: "ay-1",
				Type:           SemesterTypeGanjil,
				Name:           "Semester Ganjil",
				StartDate:      time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC),
				EndDate:        time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
				SequenceNumber: 0,
			},
			wantError:   true,
			expectedErr: "sequence_number must be 1 or 2",
		},
		{
			name: "Invalid - sequence number 3",
			request: CreateSemesterRequest{
				AcademicYearID: "ay-1",
				Type:           SemesterTypeGanjil,
				Name:           "Semester Ganjil",
				StartDate:      time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC),
				EndDate:        time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
				SequenceNumber: 3,
			},
			wantError:   true,
			expectedErr: "sequence_number must be 1 or 2",
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

// TestUpdateSemesterRequest_Validate tests the UpdateSemesterRequest Validate method
func TestUpdateSemesterRequest_Validate(t *testing.T) {
	tests := []struct {
		name        string
		request     UpdateSemesterRequest
		wantError   bool
		expectedErr string
	}{
		{
			name: "Valid update with name only",
			request: UpdateSemesterRequest{
				Name: &[]string{"Updated Semester Name"}[0],
			},
			wantError: false,
		},
		{
			name: "Valid update with dates",
			request: UpdateSemesterRequest{
				StartDate: &[]time.Time{time.Date(2026, 8, 1, 0, 0, 0, 0, time.UTC)}[0],
				EndDate:   &[]time.Time{time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC)}[0],
			},
			wantError: false,
		},
		{
			name: "Valid update with status",
			request: UpdateSemesterRequest{
				Status: &[]SemesterStatus{SemesterStatusInactive}[0],
			},
			wantError: false,
		},
		{
			name: "Valid - all nil (no updates)",
			request: UpdateSemesterRequest{
				Name:      nil,
				StartDate: nil,
				EndDate:   nil,
				Status:    nil,
			},
			wantError: false,
		},
		{
			name: "Invalid - empty name",
			request: UpdateSemesterRequest{
				Name: &[]string{""}[0],
			},
			wantError:   true,
			expectedErr: "name cannot be empty",
		},
		{
			name: "Invalid - name too long",
			request: UpdateSemesterRequest{
				Name: &[]string{string(make([]byte, 101))}[0], // 101 characters
			},
			wantError:   true,
			expectedErr: "name must be less than 100 characters",
		},
		{
			name: "Invalid - start date after end date",
			request: UpdateSemesterRequest{
				StartDate: &[]time.Time{time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC)}[0],
				EndDate:   &[]time.Time{time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC)}[0],
			},
			wantError:   true,
			expectedErr: "start_date must be before end date",
		},
		{
			name: "Invalid - invalid status",
			request: UpdateSemesterRequest{
				Status: &[]SemesterStatus{"INVALID"}[0],
			},
			wantError:   true,
			expectedErr: "status must be ACTIVE or INACTIVE",
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

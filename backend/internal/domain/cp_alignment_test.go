package domain

import (
	"testing"
	"time"
)

// TestNewCPAlignment tests the CPAlignment constructor
func TestNewCPAlignment(t *testing.T) {
	tests := []struct {
		name                        string
		curriculumSubjectID         string
		graduateProfileDimensionID  string
		createdBy                   string
		wantError                   bool
		expectedErr                 string
	}{
		{
			name:                       "Valid CP alignment",
			curriculumSubjectID:        "cp-1",
			graduateProfileDimensionID: "gpd-1",
			createdBy:                  "user-1",
			wantError:                  false,
		},
		{
			name:                       "Invalid - empty curriculum subject ID",
			curriculumSubjectID:        "",
			graduateProfileDimensionID: "gpd-1",
			createdBy:                  "user-1",
			wantError:                  true,
			expectedErr:                "curriculum subject ID is required",
		},
		{
			name:                       "Invalid - empty graduate profile dimension ID",
			curriculumSubjectID:        "cp-1",
			graduateProfileDimensionID: "",
			createdBy:                  "user-1",
			wantError:                  true,
			expectedErr:                "graduate profile dimension ID is required",
		},
		{
			name:                       "Invalid - both IDs empty",
			curriculumSubjectID:        "",
			graduateProfileDimensionID: "",
			createdBy:                  "user-1",
			wantError:                  true,
			expectedErr:                "curriculum subject ID is required",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			alignment, err := NewCPAlignment(tc.curriculumSubjectID, tc.graduateProfileDimensionID, tc.createdBy)

			if tc.wantError && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tc.wantError && err != nil {
				t.Errorf("expected no error, got: %v", err)
			}

			if !tc.wantError {
				if alignment == nil {
					t.Fatal("expected CP alignment, got nil")
				}
				if alignment.ID == "" {
					t.Error("expected ID to be generated")
				}
				if alignment.CurriculumSubjectID != tc.curriculumSubjectID {
					t.Errorf("expected curriculum subject ID %s, got %s", tc.curriculumSubjectID, alignment.CurriculumSubjectID)
				}
				if alignment.GraduateProfileDimensionID != tc.graduateProfileDimensionID {
					t.Errorf("expected graduate profile dimension ID %s, got %s", tc.graduateProfileDimensionID, alignment.GraduateProfileDimensionID)
				}
				if !alignment.IsActive {
					t.Error("expected IsActive to be true by default")
				}
				if alignment.CreatedBy != tc.createdBy {
					t.Errorf("expected created_by %s, got %s", tc.createdBy, alignment.CreatedBy)
				}
				if alignment.CreatedAt.IsZero() {
					t.Error("CreatedAt should not be zero")
				}
				if alignment.UpdatedAt.IsZero() {
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

// TestCPAlignment_Validate tests domain validation
func TestCPAlignment_Validate(t *testing.T) {
	tests := []struct {
		name      string
		alignment CPAlignment
		wantError bool
		expectedErr string
	}{
		{
			name: "Valid active CP alignment",
			alignment: CPAlignment{
				ID:                         "cpa-1",
				CurriculumSubjectID:        "cp-1",
				GraduateProfileDimensionID: "gpd-1",
				AlignmentDescription:       &[]string{"This CP supports dimension 1"}[0],
				IsActive:                   true,
				CreatedAt:                  time.Now(),
				UpdatedAt:                  time.Now(),
			},
			wantError: false,
		},
		{
			name: "Valid inactive CP alignment",
			alignment: CPAlignment{
				ID:                         "cpa-1",
				CurriculumSubjectID:        "cp-1",
				GraduateProfileDimensionID: "gpd-1",
				AlignmentDescription:       &[]string{"This CP supports dimension 1"}[0],
				IsActive:                   false,
				CreatedAt:                  time.Now(),
				UpdatedAt:                  time.Now(),
			},
			wantError: false,
		},
		{
			name: "Valid - no description",
			alignment: CPAlignment{
				ID:                         "cpa-1",
				CurriculumSubjectID:        "cp-1",
				GraduateProfileDimensionID: "gpd-1",
				AlignmentDescription:       nil,
				IsActive:                   true,
				CreatedAt:                  time.Now(),
				UpdatedAt:                  time.Now(),
			},
			wantError: false,
		},
		{
			name: "Invalid - empty curriculum subject ID",
			alignment: CPAlignment{
				ID:                         "cpa-1",
				CurriculumSubjectID:        "",
				GraduateProfileDimensionID: "gpd-1",
				IsActive:                   true,
				CreatedAt:                  time.Now(),
				UpdatedAt:                  time.Now(),
			},
			wantError:   true,
			expectedErr: "curriculum subject ID is required",
		},
		{
			name: "Invalid - empty graduate profile dimension ID",
			alignment: CPAlignment{
				ID:                         "cpa-1",
				CurriculumSubjectID:        "cp-1",
				GraduateProfileDimensionID: "",
				IsActive:                   true,
				CreatedAt:                  time.Now(),
				UpdatedAt:                  time.Now(),
			},
			wantError:   true,
			expectedErr: "graduate profile dimension ID is required",
		},
		{
			name: "Invalid - description too long (501 characters)",
			alignment: CPAlignment{
				ID:                         "cpa-1",
				CurriculumSubjectID:        "cp-1",
				GraduateProfileDimensionID: "gpd-1",
				AlignmentDescription:       &[]string{string(make([]byte, 501))}[0],
				IsActive:                   true,
				CreatedAt:                  time.Now(),
				UpdatedAt:                  time.Now(),
			},
			wantError:   true,
			expectedErr: "alignment description must be less than 500 characters",
		},
		{
			name: "Valid - description at boundary (500 characters)",
			alignment: CPAlignment{
				ID:                         "cpa-1",
				CurriculumSubjectID:        "cp-1",
				GraduateProfileDimensionID: "gpd-1",
				AlignmentDescription:       &[]string{string(make([]byte, 500))}[0],
				IsActive:                   true,
				CreatedAt:                  time.Now(),
				UpdatedAt:                  time.Now(),
			},
			wantError: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.alignment.Validate()

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

// TestCPAlignment_Activate tests the Activate method
func TestCPAlignment_Activate(t *testing.T) {
	alignment := &CPAlignment{
		ID:                         "cpa-1",
		CurriculumSubjectID:        "cp-1",
		GraduateProfileDimensionID: "gpd-1",
		IsActive:                   false,
		CreatedAt:                  time.Now(),
		UpdatedAt:                  time.Now(),
	}

	oldUpdatedAt := alignment.UpdatedAt
	time.Sleep(1 * time.Millisecond) // Ensure time difference

	alignment.Activate()

	if !alignment.IsActive {
		t.Error("expected IsActive to be true after Activate")
	}
	if !alignment.UpdatedAt.After(oldUpdatedAt) {
		t.Error("UpdatedAt should be updated after Activate")
	}
}

// TestCPAlignment_Deactivate tests the Deactivate method
func TestCPAlignment_Deactivate(t *testing.T) {
	alignment := &CPAlignment{
		ID:                         "cpa-1",
		CurriculumSubjectID:        "cp-1",
		GraduateProfileDimensionID: "gpd-1",
		IsActive:                   true,
		CreatedAt:                  time.Now(),
		UpdatedAt:                  time.Now(),
	}

	oldUpdatedAt := alignment.UpdatedAt
	time.Sleep(1 * time.Millisecond) // Ensure time difference

	alignment.Deactivate()

	if alignment.IsActive {
		t.Error("expected IsActive to be false after Deactivate")
	}
	if !alignment.UpdatedAt.After(oldUpdatedAt) {
		t.Error("UpdatedAt should be updated after Deactivate")
	}
}

// TestCPAlignment_IsActiveStatus tests the IsActiveStatus method
func TestCPAlignment_IsActiveStatus(t *testing.T) {
	tests := []struct {
		name     string
		isActive bool
		expected bool
	}{
		{
			name:     "Active alignment returns true",
			isActive: true,
			expected: true,
		},
		{
			name:     "Inactive alignment returns false",
			isActive: false,
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			alignment := &CPAlignment{
				ID:                         "cpa-1",
				CurriculumSubjectID:        "cp-1",
				GraduateProfileDimensionID: "gpd-1",
				IsActive:                   tc.isActive,
				CreatedAt:                  time.Now(),
				UpdatedAt:                  time.Now(),
			}

			result := alignment.IsActiveStatus()
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

// TestCreateCPAlignmentRequest_Validate tests the CreateCPAlignmentRequest Validate method
func TestCreateCPAlignmentRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		request CreateCPAlignmentRequest
		wantError bool
		expectedErr string
	}{
		{
			name: "Valid request with description",
			request: CreateCPAlignmentRequest{
				CurriculumSubjectID:        "cp-1",
				GraduateProfileDimensionID: "gpd-1",
				AlignmentDescription:       &[]string{"This CP supports dimension 1"}[0],
			},
			wantError: false,
		},
		{
			name: "Valid request without description",
			request: CreateCPAlignmentRequest{
				CurriculumSubjectID:        "cp-1",
				GraduateProfileDimensionID: "gpd-1",
			},
			wantError: false,
		},
		{
			name: "Invalid - empty curriculum subject ID",
			request: CreateCPAlignmentRequest{
				CurriculumSubjectID:        "",
				GraduateProfileDimensionID: "gpd-1",
			},
			wantError:   true,
			expectedErr: "curriculum_subject_id is required",
		},
		{
			name: "Invalid - empty graduate profile dimension ID",
			request: CreateCPAlignmentRequest{
				CurriculumSubjectID:        "cp-1",
				GraduateProfileDimensionID: "",
			},
			wantError:   true,
			expectedErr: "graduate_profile_dimension_id is required",
		},
		{
			name: "Invalid - both IDs empty",
			request: CreateCPAlignmentRequest{
				CurriculumSubjectID:        "",
				GraduateProfileDimensionID: "",
			},
			wantError:   true,
			expectedErr: "curriculum_subject_id is required",
		},
		{
			name: "Invalid - description too long (501 characters)",
			request: CreateCPAlignmentRequest{
				CurriculumSubjectID:        "cp-1",
				GraduateProfileDimensionID: "gpd-1",
				AlignmentDescription:       &[]string{string(make([]byte, 501))}[0],
			},
			wantError:   true,
			expectedErr: "alignment_description must be less than 500 characters",
		},
		{
			name: "Valid - description at boundary (500 characters)",
			request: CreateCPAlignmentRequest{
				CurriculumSubjectID:        "cp-1",
				GraduateProfileDimensionID: "gpd-1",
				AlignmentDescription:       &[]string{string(make([]byte, 500))}[0],
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

// TestUpdateCPAlignmentRequest_Validate tests the UpdateCPAlignmentRequest Validate method
func TestUpdateCPAlignmentRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		request UpdateCPAlignmentRequest
		wantError bool
		expectedErr string
	}{
		{
			name: "Valid update with description",
			request: UpdateCPAlignmentRequest{
				AlignmentDescription: &[]string{"Updated description"}[0],
			},
			wantError: false,
		},
		{
			name: "Valid update with is_active",
			request: UpdateCPAlignmentRequest{
				IsActive: &[]bool{false}[0],
			},
			wantError: false,
		},
		{
			name: "Valid - all nil (no updates)",
			request: UpdateCPAlignmentRequest{
				AlignmentDescription: nil,
				IsActive:             nil,
			},
			wantError: false,
		},
		{
			name: "Invalid - description too long (501 characters)",
			request: UpdateCPAlignmentRequest{
				AlignmentDescription: &[]string{string(make([]byte, 501))}[0],
			},
			wantError:   true,
			expectedErr: "alignment_description must be less than 500 characters",
		},
		{
			name: "Valid - description at boundary (500 characters)",
			request: UpdateCPAlignmentRequest{
				AlignmentDescription: &[]string{string(make([]byte, 500))}[0],
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

// TestCreateCPAlignmentBulkRequest_Validate tests the CreateCPAlignmentBulkRequest Validate method
func TestCreateCPAlignmentBulkRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		request CreateCPAlignmentBulkRequest
		wantError bool
		expectedErr string
	}{
		{
			name: "Valid bulk request with 1 alignment",
			request: CreateCPAlignmentBulkRequest{
				CurriculumSubjectID:  "cp-1",
				AlignmentIDs:         []string{"gpd-1"},
				AlignmentDescription: &[]string{"Description"}[0],
			},
			wantError: false,
		},
		{
			name: "Valid bulk request with 6 alignments (max)",
			request: CreateCPAlignmentBulkRequest{
				CurriculumSubjectID:  "cp-1",
				AlignmentIDs:         []string{"gpd-1", "gpd-2", "gpd-3", "gpd-4", "gpd-5", "gpd-6"},
				AlignmentDescription: &[]string{"Description"}[0],
			},
			wantError: false,
		},
		{
			name: "Valid - no description",
			request: CreateCPAlignmentBulkRequest{
				CurriculumSubjectID: "cp-1",
				AlignmentIDs:        []string{"gpd-1"},
			},
			wantError: false,
		},
		{
			name: "Invalid - empty curriculum subject ID",
			request: CreateCPAlignmentBulkRequest{
				CurriculumSubjectID:  "",
				AlignmentIDs:         []string{"gpd-1"},
			},
			wantError:   true,
			expectedErr: "curriculum_subject_id is required",
		},
		{
			name: "Invalid - empty alignment IDs",
			request: CreateCPAlignmentBulkRequest{
				CurriculumSubjectID: "cp-1",
				AlignmentIDs:        []string{},
			},
			wantError:   true,
			expectedErr: "at least one alignment_id is required",
		},
		{
			name: "Invalid - nil alignment IDs",
			request: CreateCPAlignmentBulkRequest{
				CurriculumSubjectID: "cp-1",
				AlignmentIDs:        nil,
			},
			wantError:   true,
			expectedErr: "at least one alignment_id is required",
		},
		{
			name: "Invalid - 7 alignments (exceeds max of 6)",
			request: CreateCPAlignmentBulkRequest{
				CurriculumSubjectID: "cp-1",
				AlignmentIDs:        []string{"gpd-1", "gpd-2", "gpd-3", "gpd-4", "gpd-5", "gpd-6", "gpd-7"},
			},
			wantError:   true,
			expectedErr: "maximum 6 alignments allowed",
		},
		{
			name: "Invalid - description too long (501 characters)",
			request: CreateCPAlignmentBulkRequest{
				CurriculumSubjectID:  "cp-1",
				AlignmentIDs:         []string{"gpd-1"},
				AlignmentDescription: &[]string{string(make([]byte, 501))}[0],
			},
			wantError:   true,
			expectedErr: "alignment_description must be less than 500 characters",
		},
		{
			name: "Valid - description at boundary (500 characters)",
			request: CreateCPAlignmentBulkRequest{
				CurriculumSubjectID:  "cp-1",
				AlignmentIDs:         []string{"gpd-1"},
				AlignmentDescription: &[]string{string(make([]byte, 500))}[0],
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

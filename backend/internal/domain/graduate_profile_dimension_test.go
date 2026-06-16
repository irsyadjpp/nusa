package domain

import (
	"testing"
	"time"
)

// TestNewGraduateProfileDimension tests the GraduateProfileDimension constructor
func TestNewGraduateProfileDimension(t *testing.T) {
	tests := []struct {
		name            string
		code            string
		dimName         string
		sequenceNumber  int
		createdBy       string
		wantError       bool
		expectedErr     string
	}{
		{
			name:           "Valid graduate profile dimension",
			code:           "DIM1",
			dimName:        "Beriman, bertakwa kepada Tuhan YME, dan berakhlak mulia",
			sequenceNumber: 1,
			createdBy:      "user-1",
			wantError:      false,
		},
		{
			name:           "Valid - sequence number 6 (max)",
			code:           "DIM6",
			dimName:        "Kreatif",
			sequenceNumber: 6,
			createdBy:      "user-1",
			wantError:      false,
		},
		{
			name:           "Invalid - empty code",
			code:           "",
			dimName:        "Beriman",
			sequenceNumber: 1,
			createdBy:      "user-1",
			wantError:      true,
			expectedErr:    "graduate profile dimension code is required",
		},
		{
			name:           "Invalid - empty name",
			code:           "DIM1",
			dimName:        "",
			sequenceNumber: 1,
			createdBy:      "user-1",
			wantError:      true,
			expectedErr:    "graduate profile dimension name is required",
		},
		{
			name:           "Invalid - sequence number 0",
			code:           "DIM1",
			dimName:        "Beriman",
			sequenceNumber: 0,
			createdBy:      "user-1",
			wantError:      true,
			expectedErr:    "sequence number must be positive",
		},
		{
			name:           "Invalid - negative sequence number",
			code:           "DIM1",
			dimName:        "Beriman",
			sequenceNumber: -1,
			createdBy:      "user-1",
			wantError:      true,
			expectedErr:    "sequence number must be positive",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dimension, err := NewGraduateProfileDimension(tc.code, tc.dimName, tc.sequenceNumber, tc.createdBy)

			if tc.wantError && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tc.wantError && err != nil {
				t.Errorf("expected no error, got: %v", err)
			}

			if !tc.wantError {
				if dimension == nil {
					t.Fatal("expected graduate profile dimension, got nil")
				}
				if dimension.ID == "" {
					t.Error("expected ID to be generated")
				}
				if dimension.Code != tc.code {
					t.Errorf("expected code %s, got %s", tc.code, dimension.Code)
				}
				if dimension.Name != tc.dimName {
					t.Errorf("expected name %s, got %s", tc.dimName, dimension.Name)
				}
				if dimension.SequenceNumber != tc.sequenceNumber {
					t.Errorf("expected sequence number %d, got %d", tc.sequenceNumber, dimension.SequenceNumber)
				}
				if !dimension.IsActive {
					t.Error("expected IsActive to be true by default")
				}
				if dimension.CreatedBy != tc.createdBy {
					t.Errorf("expected created_by %s, got %s", tc.createdBy, dimension.CreatedBy)
				}
				if dimension.CreatedAt.IsZero() {
					t.Error("CreatedAt should not be zero")
				}
				if dimension.UpdatedAt.IsZero() {
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

// TestGraduateProfileDimension_Validate tests domain validation
func TestGraduateProfileDimension_Validate(t *testing.T) {
	tests := []struct {
		name      string
		dimension GraduateProfileDimension
		wantError bool
		expectedErr string
	}{
		{
			name: "Valid active dimension",
			dimension: GraduateProfileDimension{
				ID:             "gpd-1",
				Code:           "DIM1",
				Name:           "Beriman, bertakwa kepada Tuhan YME, dan berakhlak mulia",
				Description:    &[]string{"First dimension of graduate profile"}[0],
				SequenceNumber: 1,
				IsActive:       true,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError: false,
		},
		{
			name: "Valid inactive dimension",
			dimension: GraduateProfileDimension{
				ID:             "gpd-1",
				Code:           "DIM6",
				Name:           "Kreatif",
				Description:    &[]string{"Sixth dimension of graduate profile"}[0],
				SequenceNumber: 6,
				IsActive:       false,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError: false,
		},
		{
			name: "Valid - no description",
			dimension: GraduateProfileDimension{
				ID:             "gpd-1",
				Code:           "DIM1",
				Name:           "Beriman",
				Description:    nil,
				SequenceNumber: 1,
				IsActive:       true,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError: false,
		},
		{
			name: "Valid - sequence number at minimum (1)",
			dimension: GraduateProfileDimension{
				ID:             "gpd-1",
				Code:           "DIM1",
				Name:           "Beriman",
				SequenceNumber: 1,
				IsActive:       true,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError: false,
		},
		{
			name: "Valid - sequence number at maximum (6)",
			dimension: GraduateProfileDimension{
				ID:             "gpd-1",
				Code:           "DIM6",
				Name:           "Kreatif",
				SequenceNumber: 6,
				IsActive:       true,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError: false,
		},
		{
			name: "Invalid - empty code",
			dimension: GraduateProfileDimension{
				ID:             "gpd-1",
				Code:           "",
				Name:           "Beriman",
				SequenceNumber: 1,
				IsActive:       true,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError:   true,
			expectedErr: "graduate profile dimension code is required",
		},
		{
			name: "Invalid - empty name",
			dimension: GraduateProfileDimension{
				ID:             "gpd-1",
				Code:           "DIM1",
				Name:           "",
				SequenceNumber: 1,
				IsActive:       true,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError:   true,
			expectedErr: "graduate profile dimension name is required",
		},
		{
			name: "Invalid - code too long (21 characters)",
			dimension: GraduateProfileDimension{
				ID:             "gpd-1",
				Code:           string(make([]byte, 21)),
				Name:           "Beriman",
				SequenceNumber: 1,
				IsActive:       true,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError:   true,
			expectedErr: "code must be less than 20 characters",
		},
		{
			name: "Invalid - name too long (101 characters)",
			dimension: GraduateProfileDimension{
				ID:             "gpd-1",
				Code:           "DIM1",
				Name:           string(make([]byte, 101)),
				SequenceNumber: 1,
				IsActive:       true,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError:   true,
			expectedErr: "name must be less than 100 characters",
		},
		{
			name: "Invalid - description too long (1001 characters)",
			dimension: GraduateProfileDimension{
				ID:             "gpd-1",
				Code:           "DIM1",
				Name:           "Beriman",
				Description:    &[]string{string(make([]byte, 1001))}[0],
				SequenceNumber: 1,
				IsActive:       true,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError:   true,
			expectedErr: "description must be less than 1000 characters",
		},
		{
			name: "Invalid - sequence number 0",
			dimension: GraduateProfileDimension{
				ID:             "gpd-1",
				Code:           "DIM1",
				Name:           "Beriman",
				SequenceNumber: 0,
				IsActive:       true,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError:   true,
			expectedErr: "sequence number must be between 1 and 6",
		},
		{
			name: "Invalid - sequence number 7",
			dimension: GraduateProfileDimension{
				ID:             "gpd-1",
				Code:           "DIM1",
				Name:           "Beriman",
				SequenceNumber: 7,
				IsActive:       true,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError:   true,
			expectedErr: "sequence number must be between 1 and 6",
		},
		{
			name: "Valid - code at boundary (20 characters)",
			dimension: GraduateProfileDimension{
				ID:             "gpd-1",
				Code:           string(make([]byte, 20)),
				Name:           "Beriman",
				SequenceNumber: 1,
				IsActive:       true,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError: false,
		},
		{
			name: "Valid - name at boundary (100 characters)",
			dimension: GraduateProfileDimension{
				ID:             "gpd-1",
				Code:           "DIM1",
				Name:           string(make([]byte, 100)),
				SequenceNumber: 1,
				IsActive:       true,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError: false,
		},
		{
			name: "Valid - description at boundary (1000 characters)",
			dimension: GraduateProfileDimension{
				ID:             "gpd-1",
				Code:           "DIM1",
				Name:           "Beriman",
				Description:    &[]string{string(make([]byte, 1000))}[0],
				SequenceNumber: 1,
				IsActive:       true,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.dimension.Validate()

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

// TestGraduateProfileDimension_Activate tests the Activate method
func TestGraduateProfileDimension_Activate(t *testing.T) {
	dimension := &GraduateProfileDimension{
		ID:             "gpd-1",
		Code:           "DIM1",
		Name:           "Beriman",
		SequenceNumber: 1,
		IsActive:       false,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	oldUpdatedAt := dimension.UpdatedAt
	time.Sleep(1 * time.Millisecond) // Ensure time difference

	dimension.Activate()

	if !dimension.IsActive {
		t.Error("expected IsActive to be true after Activate")
	}
	if !dimension.UpdatedAt.After(oldUpdatedAt) {
		t.Error("UpdatedAt should be updated after Activate")
	}
}

// TestGraduateProfileDimension_Deactivate tests the Deactivate method
func TestGraduateProfileDimension_Deactivate(t *testing.T) {
	dimension := &GraduateProfileDimension{
		ID:             "gpd-1",
		Code:           "DIM1",
		Name:           "Beriman",
		SequenceNumber: 1,
		IsActive:       true,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	oldUpdatedAt := dimension.UpdatedAt
	time.Sleep(1 * time.Millisecond) // Ensure time difference

	dimension.Deactivate()

	if dimension.IsActive {
		t.Error("expected IsActive to be false after Deactivate")
	}
	if !dimension.UpdatedAt.After(oldUpdatedAt) {
		t.Error("UpdatedAt should be updated after Deactivate")
	}
}

// TestGraduateProfileDimension_IsActiveStatus tests the IsActiveStatus method
func TestGraduateProfileDimension_IsActiveStatus(t *testing.T) {
	tests := []struct {
		name     string
		isActive bool
		expected bool
	}{
		{
			name:     "Active dimension returns true",
			isActive: true,
			expected: true,
		},
		{
			name:     "Inactive dimension returns false",
			isActive: false,
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dimension := &GraduateProfileDimension{
				ID:             "gpd-1",
				Code:           "DIM1",
				Name:           "Beriman",
				SequenceNumber: 1,
				IsActive:       tc.isActive,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			}

			result := dimension.IsActiveStatus()
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

// TestCreateGraduateProfileDimensionRequest_Validate tests the CreateGraduateProfileDimensionRequest Validate method
func TestCreateGraduateProfileDimensionRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		request CreateGraduateProfileDimensionRequest
		wantError bool
		expectedErr string
	}{
		{
			name: "Valid request with description",
			request: CreateGraduateProfileDimensionRequest{
				Code:           "DIM1",
				Name:           "Beriman, bertakwa kepada Tuhan YME, dan berakhlak mulia",
				Description:    &[]string{"First dimension"}[0],
				SequenceNumber: 1,
			},
			wantError: false,
		},
		{
			name: "Valid request without description",
			request: CreateGraduateProfileDimensionRequest{
				Code:           "DIM6",
				Name:           "Kreatif",
				SequenceNumber: 6,
			},
			wantError: false,
		},
		{
			name: "Valid - sequence number at minimum (1)",
			request: CreateGraduateProfileDimensionRequest{
				Code:           "DIM1",
				Name:           "Beriman",
				SequenceNumber: 1,
			},
			wantError: false,
		},
		{
			name: "Valid - sequence number at maximum (6)",
			request: CreateGraduateProfileDimensionRequest{
				Code:           "DIM6",
				Name:           "Kreatif",
				SequenceNumber: 6,
			},
			wantError: false,
		},
		{
			name: "Invalid - empty code",
			request: CreateGraduateProfileDimensionRequest{
				Code:           "",
				Name:           "Beriman",
				SequenceNumber: 1,
			},
			wantError:   true,
			expectedErr: "code is required",
		},
		{
			name: "Invalid - code too long (21 characters)",
			request: CreateGraduateProfileDimensionRequest{
				Code:           string(make([]byte, 21)),
				Name:           "Beriman",
				SequenceNumber: 1,
			},
			wantError:   true,
			expectedErr: "code must be less than 20 characters",
		},
		{
			name: "Invalid - empty name",
			request: CreateGraduateProfileDimensionRequest{
				Code:           "DIM1",
				Name:           "",
				SequenceNumber: 1,
			},
			wantError:   true,
			expectedErr: "name is required",
		},
		{
			name: "Invalid - name too long (101 characters)",
			request: CreateGraduateProfileDimensionRequest{
				Code:           "DIM1",
				Name:           string(make([]byte, 101)),
				SequenceNumber: 1,
			},
			wantError:   true,
			expectedErr: "name must be less than 100 characters",
		},
		{
			name: "Invalid - description too long (1001 characters)",
			request: CreateGraduateProfileDimensionRequest{
				Code:           "DIM1",
				Name:           "Beriman",
				Description:    &[]string{string(make([]byte, 1001))}[0],
				SequenceNumber: 1,
			},
			wantError:   true,
			expectedErr: "description must be less than 1000 characters",
		},
		{
			name: "Invalid - sequence number 0",
			request: CreateGraduateProfileDimensionRequest{
				Code:           "DIM1",
				Name:           "Beriman",
				SequenceNumber: 0,
			},
			wantError:   true,
			expectedErr: "sequence_number must be between 1 and 6",
		},
		{
			name: "Invalid - sequence number 7",
			request: CreateGraduateProfileDimensionRequest{
				Code:           "DIM1",
				Name:           "Beriman",
				SequenceNumber: 7,
			},
			wantError:   true,
			expectedErr: "sequence_number must be between 1 and 6",
		},
		{
			name: "Valid - code at boundary (20 characters)",
			request: CreateGraduateProfileDimensionRequest{
				Code:           string(make([]byte, 20)),
				Name:           "Beriman",
				SequenceNumber: 1,
			},
			wantError: false,
		},
		{
			name: "Valid - name at boundary (100 characters)",
			request: CreateGraduateProfileDimensionRequest{
				Code:           "DIM1",
				Name:           string(make([]byte, 100)),
				SequenceNumber: 1,
			},
			wantError: false,
		},
		{
			name: "Valid - description at boundary (1000 characters)",
			request: CreateGraduateProfileDimensionRequest{
				Code:           "DIM1",
				Name:           "Beriman",
				Description:    &[]string{string(make([]byte, 1000))}[0],
				SequenceNumber: 1,
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

// TestUpdateGraduateProfileDimensionRequest_Validate tests the UpdateGraduateProfileDimensionRequest Validate method
func TestUpdateGraduateProfileDimensionRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		request UpdateGraduateProfileDimensionRequest
		wantError bool
		expectedErr string
	}{
		{
			name: "Valid update with name only",
			request: UpdateGraduateProfileDimensionRequest{
				Name: &[]string{"Updated Name"}[0],
			},
			wantError: false,
		},
		{
			name: "Valid update with description",
			request: UpdateGraduateProfileDimensionRequest{
				Description: &[]string{"Updated description"}[0],
			},
			wantError: false,
		},
		{
			name: "Valid update with sequence number",
			request: UpdateGraduateProfileDimensionRequest{
				SequenceNumber: &[]int{2}[0],
			},
			wantError: false,
		},
		{
			name: "Valid update with is_active",
			request: UpdateGraduateProfileDimensionRequest{
				IsActive: &[]bool{false}[0],
			},
			wantError: false,
		},
		{
			name: "Valid - all nil (no updates)",
			request: UpdateGraduateProfileDimensionRequest{
				Name:           nil,
				Description:    nil,
				SequenceNumber: nil,
				IsActive:       nil,
			},
			wantError: false,
		},
		{
			name: "Invalid - empty name",
			request: UpdateGraduateProfileDimensionRequest{
				Name: &[]string{""}[0],
			},
			wantError:   true,
			expectedErr: "name cannot be empty",
		},
		{
			name: "Invalid - name too long (101 characters)",
			request: UpdateGraduateProfileDimensionRequest{
				Name: &[]string{string(make([]byte, 101))}[0],
			},
			wantError:   true,
			expectedErr: "name must be less than 100 characters",
		},
		{
			name: "Invalid - description too long (1001 characters)",
			request: UpdateGraduateProfileDimensionRequest{
				Description: &[]string{string(make([]byte, 1001))}[0],
			},
			wantError:   true,
			expectedErr: "description must be less than 1000 characters",
		},
		{
			name: "Invalid - sequence number 0",
			request: UpdateGraduateProfileDimensionRequest{
				SequenceNumber: &[]int{0}[0],
			},
			wantError:   true,
			expectedErr: "sequence_number must be between 1 and 6",
		},
		{
			name: "Invalid - sequence number 7",
			request: UpdateGraduateProfileDimensionRequest{
				SequenceNumber: &[]int{7}[0],
			},
			wantError:   true,
			expectedErr: "sequence_number must be between 1 and 6",
		},
		{
			name: "Valid - name at boundary (100 characters)",
			request: UpdateGraduateProfileDimensionRequest{
				Name: &[]string{string(make([]byte, 100))}[0],
			},
			wantError: false,
		},
		{
			name: "Valid - description at boundary (1000 characters)",
			request: UpdateGraduateProfileDimensionRequest{
				Description: &[]string{string(make([]byte, 1000))}[0],
			},
			wantError: false,
		},
		{
			name: "Valid - sequence number at minimum (1)",
			request: UpdateGraduateProfileDimensionRequest{
				SequenceNumber: &[]int{1}[0],
			},
			wantError: false,
		},
		{
			name: "Valid - sequence number at maximum (6)",
			request: UpdateGraduateProfileDimensionRequest{
				SequenceNumber: &[]int{6}[0],
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

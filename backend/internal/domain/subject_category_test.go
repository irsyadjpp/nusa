package domain

import (
	"testing"
	"time"
)

// TestNewSubjectCategory tests the SubjectCategory constructor
func TestNewSubjectCategory(t *testing.T) {
	tests := []struct {
		name        string
		code        string
		catName     string
		isMandatory bool
		createdBy   string
		wantError   bool
		expectedErr string
	}{
		{
			name:        "Valid mandatory subject category",
			code:        "MATH",
			catName:     "Mathematics",
			isMandatory: true,
			createdBy:   "user-1",
			wantError:   false,
		},
		{
			name:        "Valid optional subject category",
			code:        "ART",
			catName:     "Arts",
			isMandatory: false,
			createdBy:   "user-1",
			wantError:   false,
		},
		{
			name:        "Invalid - empty code",
			code:        "",
			catName:     "Mathematics",
			isMandatory: true,
			createdBy:   "user-1",
			wantError:   true,
			expectedErr: "subject category code is required",
		},
		{
			name:        "Invalid - empty name",
			code:        "MATH",
			catName:     "",
			isMandatory: true,
			createdBy:   "user-1",
			wantError:   true,
			expectedErr: "subject category name is required",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			category, err := NewSubjectCategory(tc.code, tc.catName, tc.isMandatory, tc.createdBy)

			if tc.wantError && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tc.wantError && err != nil {
				t.Errorf("expected no error, got: %v", err)
			}

			if !tc.wantError {
				if category == nil {
					t.Fatal("expected subject category, got nil")
				}
				if category.ID == "" {
					t.Error("expected ID to be generated")
				}
				if category.Code != tc.code {
					t.Errorf("expected code %s, got %s", tc.code, category.Code)
				}
				if category.Name != tc.catName {
					t.Errorf("expected name %s, got %s", tc.catName, category.Name)
				}
				if category.IsMandatory != tc.isMandatory {
					t.Errorf("expected is_mandatory %v, got %v", tc.isMandatory, category.IsMandatory)
				}
				if !category.IsActive {
					t.Error("expected IsActive to be true by default")
				}
				if category.CreatedBy != tc.createdBy {
					t.Errorf("expected created_by %s, got %s", tc.createdBy, category.CreatedBy)
				}
				if category.CreatedAt.IsZero() {
					t.Error("CreatedAt should not be zero")
				}
				if category.UpdatedAt.IsZero() {
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

// TestSubjectCategory_Validate tests domain validation
func TestSubjectCategory_Validate(t *testing.T) {
	tests := []struct {
		name     string
		category SubjectCategory
		wantError bool
		expectedErr string
	}{
		{
			name: "Valid active subject category",
			category: SubjectCategory{
				ID:          "sc-1",
				Code:        "MATH",
				Name:        "Mathematics",
				Description: &[]string{"Mathematics subject group"}[0],
				IsMandatory: true,
				IsActive:    true,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			wantError: false,
		},
		{
			name: "Valid inactive subject category",
			category: SubjectCategory{
				ID:          "sc-1",
				Code:        "ART",
				Name:        "Arts",
				Description: &[]string{"Arts subject group"}[0],
				IsMandatory: false,
				IsActive:    false,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			wantError: false,
		},
		{
			name: "Valid - no description",
			category: SubjectCategory{
				ID:          "sc-1",
				Code:        "MATH",
				Name:        "Mathematics",
				Description: nil,
				IsMandatory: true,
				IsActive:    true,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			wantError: false,
		},
		{
			name: "Invalid - empty code",
			category: SubjectCategory{
				ID:          "sc-1",
				Code:        "",
				Name:        "Mathematics",
				IsMandatory: true,
				IsActive:    true,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			wantError:   true,
			expectedErr: "subject category code is required",
		},
		{
			name: "Invalid - empty name",
			category: SubjectCategory{
				ID:          "sc-1",
				Code:        "MATH",
				Name:        "",
				IsMandatory: true,
				IsActive:    true,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			wantError:   true,
			expectedErr: "subject category name is required",
		},
		{
			name: "Invalid - code too long (21 characters)",
			category: SubjectCategory{
				ID:          "sc-1",
				Code:        string(make([]byte, 21)),
				Name:        "Mathematics",
				IsMandatory: true,
				IsActive:    true,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			wantError:   true,
			expectedErr: "code must be less than 20 characters",
		},
		{
			name: "Invalid - name too long (101 characters)",
			category: SubjectCategory{
				ID:          "sc-1",
				Code:        "MATH",
				Name:        string(make([]byte, 101)),
				IsMandatory: true,
				IsActive:    true,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			wantError:   true,
			expectedErr: "name must be less than 100 characters",
		},
		{
			name: "Invalid - description too long (501 characters)",
			category: SubjectCategory{
				ID:          "sc-1",
				Code:        "MATH",
				Name:        "Mathematics",
				Description: &[]string{string(make([]byte, 501))}[0],
				IsMandatory: true,
				IsActive:    true,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			wantError:   true,
			expectedErr: "description must be less than 500 characters",
		},
		{
			name: "Valid - code at boundary (20 characters)",
			category: SubjectCategory{
				ID:          "sc-1",
				Code:        string(make([]byte, 20)),
				Name:        "Mathematics",
				IsMandatory: true,
				IsActive:    true,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			wantError: false,
		},
		{
			name: "Valid - name at boundary (100 characters)",
			category: SubjectCategory{
				ID:          "sc-1",
				Code:        "MATH",
				Name:        string(make([]byte, 100)),
				IsMandatory: true,
				IsActive:    true,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			wantError: false,
		},
		{
			name: "Valid - description at boundary (500 characters)",
			category: SubjectCategory{
				ID:          "sc-1",
				Code:        "MATH",
				Name:        "Mathematics",
				Description: &[]string{string(make([]byte, 500))}[0],
				IsMandatory: true,
				IsActive:    true,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			wantError: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.category.Validate()

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

// TestSubjectCategory_Activate tests the Activate method
func TestSubjectCategory_Activate(t *testing.T) {
	category := &SubjectCategory{
		ID:          "sc-1",
		Code:        "MATH",
		Name:        "Mathematics",
		IsMandatory: true,
		IsActive:    false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	oldUpdatedAt := category.UpdatedAt
	time.Sleep(1 * time.Millisecond) // Ensure time difference

	category.Activate()

	if !category.IsActive {
		t.Error("expected IsActive to be true after Activate")
	}
	if !category.UpdatedAt.After(oldUpdatedAt) {
		t.Error("UpdatedAt should be updated after Activate")
	}
}

// TestSubjectCategory_Deactivate tests the Deactivate method
func TestSubjectCategory_Deactivate(t *testing.T) {
	category := &SubjectCategory{
		ID:          "sc-1",
		Code:        "MATH",
		Name:        "Mathematics",
		IsMandatory: true,
		IsActive:    true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	oldUpdatedAt := category.UpdatedAt
	time.Sleep(1 * time.Millisecond) // Ensure time difference

	category.Deactivate()

	if category.IsActive {
		t.Error("expected IsActive to be false after Deactivate")
	}
	if !category.UpdatedAt.After(oldUpdatedAt) {
		t.Error("UpdatedAt should be updated after Deactivate")
	}
}

// TestSubjectCategory_IsActiveStatus tests the IsActiveStatus method
func TestSubjectCategory_IsActiveStatus(t *testing.T) {
	tests := []struct {
		name     string
		isActive bool
		expected bool
	}{
		{
			name:     "Active category returns true",
			isActive: true,
			expected: true,
		},
		{
			name:     "Inactive category returns false",
			isActive: false,
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			category := &SubjectCategory{
				ID:          "sc-1",
				Code:        "MATH",
				Name:        "Mathematics",
				IsMandatory: true,
				IsActive:    tc.isActive,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}

			result := category.IsActiveStatus()
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

// TestCreateSubjectCategoryRequest_Validate tests the CreateSubjectCategoryRequest Validate method
func TestCreateSubjectCategoryRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		request CreateSubjectCategoryRequest
		wantError bool
		expectedErr string
	}{
		{
			name: "Valid mandatory subject category request",
			request: CreateSubjectCategoryRequest{
				Code:        "MATH",
				Name:        "Mathematics",
				Description: &[]string{"Mathematics subject group"}[0],
				IsMandatory: true,
			},
			wantError: false,
		},
		{
			name: "Valid optional subject category request",
			request: CreateSubjectCategoryRequest{
				Code:        "ART",
				Name:        "Arts",
				IsMandatory: false,
			},
			wantError: false,
		},
		{
			name: "Valid - no description",
			request: CreateSubjectCategoryRequest{
				Code:        "MATH",
				Name:        "Mathematics",
				Description: nil,
				IsMandatory: true,
			},
			wantError: false,
		},
		{
			name: "Invalid - empty code",
			request: CreateSubjectCategoryRequest{
				Code:        "",
				Name:        "Mathematics",
				IsMandatory: true,
			},
			wantError:   true,
			expectedErr: "code is required",
		},
		{
			name: "Invalid - code too long (21 characters)",
			request: CreateSubjectCategoryRequest{
				Code:        string(make([]byte, 21)),
				Name:        "Mathematics",
				IsMandatory: true,
			},
			wantError:   true,
			expectedErr: "code must be less than 20 characters",
		},
		{
			name: "Invalid - empty name",
			request: CreateSubjectCategoryRequest{
				Code:        "MATH",
				Name:        "",
				IsMandatory: true,
			},
			wantError:   true,
			expectedErr: "name is required",
		},
		{
			name: "Invalid - name too long (101 characters)",
			request: CreateSubjectCategoryRequest{
				Code:        "MATH",
				Name:        string(make([]byte, 101)),
				IsMandatory: true,
			},
			wantError:   true,
			expectedErr: "name must be less than 100 characters",
		},
		{
			name: "Invalid - description too long (501 characters)",
			request: CreateSubjectCategoryRequest{
				Code:        "MATH",
				Name:        "Mathematics",
				Description: &[]string{string(make([]byte, 501))}[0],
				IsMandatory: true,
			},
			wantError:   true,
			expectedErr: "description must be less than 500 characters",
		},
		{
			name: "Valid - code at boundary (20 characters)",
			request: CreateSubjectCategoryRequest{
				Code:        string(make([]byte, 20)),
				Name:        "Mathematics",
				IsMandatory: true,
			},
			wantError: false,
		},
		{
			name: "Valid - name at boundary (100 characters)",
			request: CreateSubjectCategoryRequest{
				Code:        "MATH",
				Name:        string(make([]byte, 100)),
				IsMandatory: true,
			},
			wantError: false,
		},
		{
			name: "Valid - description at boundary (500 characters)",
			request: CreateSubjectCategoryRequest{
				Code:        "MATH",
				Name:        "Mathematics",
				Description: &[]string{string(make([]byte, 500))}[0],
				IsMandatory: true,
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

// TestUpdateSubjectCategoryRequest_Validate tests the UpdateSubjectCategoryRequest Validate method
func TestUpdateSubjectCategoryRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		request UpdateSubjectCategoryRequest
		wantError bool
		expectedErr string
	}{
		{
			name: "Valid update with name only",
			request: UpdateSubjectCategoryRequest{
				Name: &[]string{"Updated Mathematics"}[0],
			},
			wantError: false,
		},
		{
			name: "Valid update with description",
			request: UpdateSubjectCategoryRequest{
				Description: &[]string{"Updated description"}[0],
			},
			wantError: false,
		},
		{
			name: "Valid update with is_mandatory",
			request: UpdateSubjectCategoryRequest{
				IsMandatory: &[]bool{false}[0],
			},
			wantError: false,
		},
		{
			name: "Valid update with is_active",
			request: UpdateSubjectCategoryRequest{
				IsActive: &[]bool{false}[0],
			},
			wantError: false,
		},
		{
			name: "Valid - all nil (no updates)",
			request: UpdateSubjectCategoryRequest{
				Name:        nil,
				Description: nil,
				IsMandatory: nil,
				IsActive:    nil,
			},
			wantError: false,
		},
		{
			name: "Invalid - empty name",
			request: UpdateSubjectCategoryRequest{
				Name: &[]string{""}[0],
			},
			wantError:   true,
			expectedErr: "name cannot be empty",
		},
		{
			name: "Invalid - name too long (101 characters)",
			request: UpdateSubjectCategoryRequest{
				Name: &[]string{string(make([]byte, 101))}[0],
			},
			wantError:   true,
			expectedErr: "name must be less than 100 characters",
		},
		{
			name: "Invalid - description too long (501 characters)",
			request: UpdateSubjectCategoryRequest{
				Description: &[]string{string(make([]byte, 501))}[0],
			},
			wantError:   true,
			expectedErr: "description must be less than 500 characters",
		},
		{
			name: "Valid - name at boundary (100 characters)",
			request: UpdateSubjectCategoryRequest{
				Name: &[]string{string(make([]byte, 100))}[0],
			},
			wantError: false,
		},
		{
			name: "Valid - description at boundary (500 characters)",
			request: UpdateSubjectCategoryRequest{
				Description: &[]string{string(make([]byte, 500))}[0],
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

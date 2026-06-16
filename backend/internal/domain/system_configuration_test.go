package domain

import (
	"testing"
	"time"
)

// TestNewSystemConfiguration tests the SystemConfiguration constructor
func TestNewSystemConfiguration(t *testing.T) {
	tests := []struct {
		name        string
		key         string
		value       string
		valueType   string
		category    string
		isSystem    bool
		createdBy   string
		wantError   bool
		expectedErr string
	}{
		{
			name:      "Valid system configuration",
			key:       "CP_ALIGNMENT_THRESHOLD",
			value:     "60",
			valueType: "number",
			category:  "CURRICULUM",
			isSystem:  true,
			createdBy: "user-1",
			wantError: false,
		},
		{
			name:      "Valid user configuration",
			key:       "SCHOOL_NAME",
			value:     "NUSA School",
			valueType: "string",
			category:  "SCHOOL",
			isSystem:  false,
			createdBy: "user-1",
			wantError: false,
		},
		{
			name:      "Valid - boolean value type",
			key:       "ENABLE_FEATURE",
			value:     "true",
			valueType: "boolean",
			category:  "FEATURE",
			isSystem:  false,
			createdBy: "user-1",
			wantError: false,
		},
		{
			name:      "Valid - json value type",
			key:       "CONFIG_JSON",
			value:     `{"key": "value"}`,
			valueType: "json",
			category:  "GENERAL",
			isSystem:  false,
			createdBy: "user-1",
			wantError: false,
		},
		{
			name:        "Invalid - empty key",
			key:         "",
			value:       "60",
			valueType:   "number",
			category:    "CURRICULUM",
			isSystem:    true,
			createdBy:   "user-1",
			wantError:   true,
			expectedErr: "configuration key is required",
		},
		{
			name:        "Invalid - empty value",
			key:         "CP_ALIGNMENT_THRESHOLD",
			value:       "",
			valueType:   "number",
			category:    "CURRICULUM",
			isSystem:    true,
			createdBy:   "user-1",
			wantError:   true,
			expectedErr: "configuration value is required",
		},
		{
			name:        "Invalid - empty value type",
			key:         "CP_ALIGNMENT_THRESHOLD",
			value:       "60",
			valueType:   "",
			category:    "CURRICULUM",
			isSystem:    true,
			createdBy:   "user-1",
			wantError:   true,
			expectedErr: "configuration value type is required",
		},
		{
			name:        "Invalid - empty category",
			key:         "CP_ALIGNMENT_THRESHOLD",
			value:       "60",
			valueType:   "number",
			category:    "",
			isSystem:    true,
			createdBy:   "user-1",
			wantError:   true,
			expectedErr: "configuration category is required",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			config, err := NewSystemConfiguration(tc.key, tc.value, tc.valueType, tc.category, tc.isSystem, tc.createdBy)

			if tc.wantError && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tc.wantError && err != nil {
				t.Errorf("expected no error, got: %v", err)
			}

			if !tc.wantError {
				if config == nil {
					t.Fatal("expected system configuration, got nil")
				}
				if config.ID == "" {
					t.Error("expected ID to be generated")
				}
				if config.Key != tc.key {
					t.Errorf("expected key %s, got %s", tc.key, config.Key)
				}
				if config.Value != tc.value {
					t.Errorf("expected value %s, got %s", tc.value, config.Value)
				}
				if config.ValueType != tc.valueType {
					t.Errorf("expected value type %s, got %s", tc.valueType, config.ValueType)
				}
				if config.Category != tc.category {
					t.Errorf("expected category %s, got %s", tc.category, config.Category)
				}
				if config.IsSystem != tc.isSystem {
					t.Errorf("expected is_system %v, got %v", tc.isSystem, config.IsSystem)
				}
				if !config.IsActive {
					t.Error("expected IsActive to be true by default")
				}
				if config.CreatedBy != tc.createdBy {
					t.Errorf("expected created_by %s, got %s", tc.createdBy, config.CreatedBy)
				}
				if config.CreatedAt.IsZero() {
					t.Error("CreatedAt should not be zero")
				}
				if config.UpdatedAt.IsZero() {
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

// TestSystemConfiguration_Validate tests domain validation
func TestSystemConfiguration_Validate(t *testing.T) {
	tests := []struct {
		name     string
		config   SystemConfiguration
		wantError bool
		expectedErr string
	}{
		{
			name: "Valid system configuration",
			config: SystemConfiguration{
				ID:        "sc-1",
				Key:       "CP_ALIGNMENT_THRESHOLD",
				Value:     "60",
				ValueType: "number",
				Category:  "CURRICULUM",
				IsSystem:  true,
				IsActive:  true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantError: false,
		},
		{
			name: "Valid - string value type",
			config: SystemConfiguration{
				ID:        "sc-1",
				Key:       "SCHOOL_NAME",
				Value:     "NUSA School",
				ValueType: "string",
				Category:  "SCHOOL",
				IsSystem:  false,
				IsActive:  true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantError: false,
		},
		{
			name: "Valid - boolean value type",
			config: SystemConfiguration{
				ID:        "sc-1",
				Key:       "ENABLE_FEATURE",
				Value:     "true",
				ValueType: "boolean",
				Category:  "FEATURE",
				IsSystem:  false,
				IsActive:  true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantError: false,
		},
		{
			name: "Valid - json value type",
			config: SystemConfiguration{
				ID:        "sc-1",
				Key:       "CONFIG_JSON",
				Value:     `{"key": "value"}`,
				ValueType: "json",
				Category:  "GENERAL",
				IsSystem:  false,
				IsActive:  true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantError: false,
		},
		{
			name: "Valid - no description",
			config: SystemConfiguration{
				ID:        "sc-1",
				Key:       "CP_ALIGNMENT_THRESHOLD",
				Value:     "60",
				ValueType: "number",
				Category:  "CURRICULUM",
				Description: nil,
				IsSystem:  true,
				IsActive:  true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantError: false,
		},
		{
			name: "Invalid - empty key",
			config: SystemConfiguration{
				ID:        "sc-1",
				Key:       "",
				Value:     "60",
				ValueType: "number",
				Category:  "CURRICULUM",
				IsActive:  true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantError:   true,
			expectedErr: "configuration key is required",
		},
		{
			name: "Invalid - empty value",
			config: SystemConfiguration{
				ID:        "sc-1",
				Key:       "CP_ALIGNMENT_THRESHOLD",
				Value:     "",
				ValueType: "number",
				Category:  "CURRICULUM",
				IsActive:  true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantError:   true,
			expectedErr: "configuration value is required",
		},
		{
			name: "Invalid - empty value type",
			config: SystemConfiguration{
				ID:        "sc-1",
				Key:       "CP_ALIGNMENT_THRESHOLD",
				Value:     "60",
				ValueType: "",
				Category:  "CURRICULUM",
				IsActive:  true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantError:   true,
			expectedErr: "configuration value type is required",
		},
		{
			name: "Invalid - empty category",
			config: SystemConfiguration{
				ID:        "sc-1",
				Key:       "CP_ALIGNMENT_THRESHOLD",
				Value:     "60",
				ValueType: "number",
				Category:  "",
				IsActive:  true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantError:   true,
			expectedErr: "configuration category is required",
		},
		{
			name: "Invalid - invalid value type",
			config: SystemConfiguration{
				ID:        "sc-1",
				Key:       "CP_ALIGNMENT_THRESHOLD",
				Value:     "60",
				ValueType: "invalid",
				Category:  "CURRICULUM",
				IsActive:  true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantError:   true,
			expectedErr: "invalid value type",
		},
		{
			name: "Invalid - key too long (101 characters)",
			config: SystemConfiguration{
				ID:        "sc-1",
				Key:       string(make([]byte, 101)),
				Value:     "60",
				ValueType: "number",
				Category:  "CURRICULUM",
				IsActive:  true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantError:   true,
			expectedErr: "key must be less than 100 characters",
		},
		{
			name: "Invalid - category too long (51 characters)",
			config: SystemConfiguration{
				ID:        "sc-1",
				Key:       "CP_ALIGNMENT_THRESHOLD",
				Value:     "60",
				ValueType: "number",
				Category:  string(make([]byte, 51)),
				IsActive:  true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantError:   true,
			expectedErr: "category must be less than 50 characters",
		},
		{
			name: "Invalid - description too long (501 characters)",
			config: SystemConfiguration{
				ID:        "sc-1",
				Key:       "CP_ALIGNMENT_THRESHOLD",
				Value:     "60",
				ValueType: "number",
				Category:  "CURRICULUM",
				Description: &[]string{string(make([]byte, 501))}[0],
				IsActive:  true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantError:   true,
			expectedErr: "description must be less than 500 characters",
		},
		{
			name: "Valid - key at boundary (100 characters)",
			config: SystemConfiguration{
				ID:        "sc-1",
				Key:       string(make([]byte, 100)),
				Value:     "60",
				ValueType: "number",
				Category:  "CURRICULUM",
				IsActive:  true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantError: false,
		},
		{
			name: "Valid - category at boundary (50 characters)",
			config: SystemConfiguration{
				ID:        "sc-1",
				Key:       "CP_ALIGNMENT_THRESHOLD",
				Value:     "60",
				ValueType: "number",
				Category:  string(make([]byte, 50)),
				IsActive:  true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantError: false,
		},
		{
			name: "Valid - description at boundary (500 characters)",
			config: SystemConfiguration{
				ID:        "sc-1",
				Key:       "CP_ALIGNMENT_THRESHOLD",
				Value:     "60",
				ValueType: "number",
				Category:  "CURRICULUM",
				Description: &[]string{string(make([]byte, 500))}[0],
				IsActive:  true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantError: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.config.Validate()

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

// TestSystemConfiguration_Activate tests the Activate method
func TestSystemConfiguration_Activate(t *testing.T) {
	config := &SystemConfiguration{
		ID:        "sc-1",
		Key:       "CP_ALIGNMENT_THRESHOLD",
		Value:     "60",
		ValueType: "number",
		Category:  "CURRICULUM",
		IsSystem:  true,
		IsActive:  false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	oldUpdatedAt := config.UpdatedAt
	time.Sleep(1 * time.Millisecond) // Ensure time difference

	config.Activate()

	if !config.IsActive {
		t.Error("expected IsActive to be true after Activate")
	}
	if !config.UpdatedAt.After(oldUpdatedAt) {
		t.Error("UpdatedAt should be updated after Activate")
	}
}

// TestSystemConfiguration_Deactivate tests the Deactivate method
func TestSystemConfiguration_Deactivate(t *testing.T) {
	config := &SystemConfiguration{
		ID:        "sc-1",
		Key:       "CP_ALIGNMENT_THRESHOLD",
		Value:     "60",
		ValueType: "number",
		Category:  "CURRICULUM",
		IsSystem:  true,
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	oldUpdatedAt := config.UpdatedAt
	time.Sleep(1 * time.Millisecond) // Ensure time difference

	config.Deactivate()

	if config.IsActive {
		t.Error("expected IsActive to be false after Deactivate")
	}
	if !config.UpdatedAt.After(oldUpdatedAt) {
		t.Error("UpdatedAt should be updated after Deactivate")
	}
}

// TestSystemConfiguration_IsSystemConfig tests the IsSystemConfig method
func TestSystemConfiguration_IsSystemConfig(t *testing.T) {
	tests := []struct {
		name     string
		isSystem bool
		expected bool
	}{
		{
			name:     "System configuration returns true",
			isSystem: true,
			expected: true,
		},
		{
			name:     "User configuration returns false",
			isSystem: false,
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			config := &SystemConfiguration{
				ID:        "sc-1",
				Key:       "CP_ALIGNMENT_THRESHOLD",
				Value:     "60",
				ValueType: "number",
				Category:  "CURRICULUM",
				IsSystem:  tc.isSystem,
				IsActive:  true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}

			result := config.IsSystemConfig()
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

// TestSystemConfiguration_CanBeDeleted tests the CanBeDeleted method
func TestSystemConfiguration_CanBeDeleted(t *testing.T) {
	tests := []struct {
		name     string
		isSystem bool
		expected bool
	}{
		{
			name:     "System configuration cannot be deleted",
			isSystem: true,
			expected: false,
		},
		{
			name:     "User configuration can be deleted",
			isSystem: false,
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			config := &SystemConfiguration{
				ID:        "sc-1",
				Key:       "CP_ALIGNMENT_THRESHOLD",
				Value:     "60",
				ValueType: "number",
				Category:  "CURRICULUM",
				IsSystem:  tc.isSystem,
				IsActive:  true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}

			result := config.CanBeDeleted()
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

// TestCreateSystemConfigurationRequest_Validate tests the CreateSystemConfigurationRequest Validate method
func TestCreateSystemConfigurationRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		request CreateSystemConfigurationRequest
		wantError bool
		expectedErr string
	}{
		{
			name: "Valid system configuration request",
			request: CreateSystemConfigurationRequest{
				Key:         "CP_ALIGNMENT_THRESHOLD",
				Value:       "60",
				ValueType:   "number",
				Description: &[]string{"CP alignment threshold percentage"}[0],
				Category:    "CURRICULUM",
				IsSystem:    true,
			},
			wantError: false,
		},
		{
			name: "Valid - string value type",
			request: CreateSystemConfigurationRequest{
				Key:       "SCHOOL_NAME",
				Value:     "NUSA School",
				ValueType: "string",
				Category:  "SCHOOL",
				IsSystem:  false,
			},
			wantError: false,
		},
		{
			name: "Valid - boolean value type",
			request: CreateSystemConfigurationRequest{
				Key:       "ENABLE_FEATURE",
				Value:     "true",
				ValueType: "boolean",
				Category:  "FEATURE",
				IsSystem:  false,
			},
			wantError: false,
		},
		{
			name: "Valid - json value type",
			request: CreateSystemConfigurationRequest{
				Key:       "CONFIG_JSON",
				Value:     `{"key": "value"}`,
				ValueType: "json",
				Category:  "GENERAL",
				IsSystem:  false,
			},
			wantError: false,
		},
		{
			name: "Valid - no description",
			request: CreateSystemConfigurationRequest{
				Key:         "CP_ALIGNMENT_THRESHOLD",
				Value:       "60",
				ValueType:   "number",
				Description: nil,
				Category:    "CURRICULUM",
				IsSystem:    true,
			},
			wantError: false,
		},
		{
			name: "Invalid - empty key",
			request: CreateSystemConfigurationRequest{
				Key:       "",
				Value:     "60",
				ValueType: "number",
				Category:  "CURRICULUM",
				IsSystem:  true,
			},
			wantError:   true,
			expectedErr: "key is required",
		},
		{
			name: "Invalid - key too long (101 characters)",
			request: CreateSystemConfigurationRequest{
				Key:       string(make([]byte, 101)),
				Value:     "60",
				ValueType: "number",
				Category:  "CURRICULUM",
				IsSystem:  true,
			},
			wantError:   true,
			expectedErr: "key must be less than 100 characters",
		},
		{
			name: "Invalid - empty value",
			request: CreateSystemConfigurationRequest{
				Key:       "CP_ALIGNMENT_THRESHOLD",
				Value:     "",
				ValueType: "number",
				Category:  "CURRICULUM",
				IsSystem:  true,
			},
			wantError:   true,
			expectedErr: "value is required",
		},
		{
			name: "Invalid - empty value type",
			request: CreateSystemConfigurationRequest{
				Key:       "CP_ALIGNMENT_THRESHOLD",
				Value:     "60",
				ValueType: "",
				Category:  "CURRICULUM",
				IsSystem:  true,
			},
			wantError:   true,
			expectedErr: "value_type is required",
		},
		{
			name: "Invalid - invalid value type",
			request: CreateSystemConfigurationRequest{
				Key:       "CP_ALIGNMENT_THRESHOLD",
				Value:     "60",
				ValueType: "invalid",
				Category:  "CURRICULUM",
				IsSystem:  true,
			},
			wantError:   true,
			expectedErr: "value_type must be string, number, boolean, or json",
		},
		{
			name: "Invalid - empty category",
			request: CreateSystemConfigurationRequest{
				Key:       "CP_ALIGNMENT_THRESHOLD",
				Value:     "60",
				ValueType: "number",
				Category:  "",
				IsSystem:  true,
			},
			wantError:   true,
			expectedErr: "category is required",
		},
		{
			name: "Invalid - category too long (51 characters)",
			request: CreateSystemConfigurationRequest{
				Key:       "CP_ALIGNMENT_THRESHOLD",
				Value:     "60",
				ValueType: "number",
				Category:  string(make([]byte, 51)),
				IsSystem:  true,
			},
			wantError:   true,
			expectedErr: "category must be less than 50 characters",
		},
		{
			name: "Invalid - description too long (501 characters)",
			request: CreateSystemConfigurationRequest{
				Key:         "CP_ALIGNMENT_THRESHOLD",
				Value:       "60",
				ValueType:   "number",
				Description: &[]string{string(make([]byte, 501))}[0],
				Category:    "CURRICULUM",
				IsSystem:    true,
			},
			wantError:   true,
			expectedErr: "description must be less than 500 characters",
		},
		{
			name: "Valid - key at boundary (100 characters)",
			request: CreateSystemConfigurationRequest{
				Key:       string(make([]byte, 100)),
				Value:     "60",
				ValueType: "number",
				Category:  "CURRICULUM",
				IsSystem:  true,
			},
			wantError: false,
		},
		{
			name: "Valid - category at boundary (50 characters)",
			request: CreateSystemConfigurationRequest{
				Key:       "CP_ALIGNMENT_THRESHOLD",
				Value:     "60",
				ValueType: "number",
				Category:  string(make([]byte, 50)),
				IsSystem:  true,
			},
			wantError: false,
		},
		{
			name: "Valid - description at boundary (500 characters)",
			request: CreateSystemConfigurationRequest{
				Key:         "CP_ALIGNMENT_THRESHOLD",
				Value:       "60",
				ValueType:   "number",
				Description: &[]string{string(make([]byte, 500))}[0],
				Category:    "CURRICULUM",
				IsSystem:    true,
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

// TestUpdateSystemConfigurationRequest_Validate tests the UpdateSystemConfigurationRequest Validate method
func TestUpdateSystemConfigurationRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		request UpdateSystemConfigurationRequest
		wantError bool
		expectedErr string
	}{
		{
			name: "Valid update with value",
			request: UpdateSystemConfigurationRequest{
				Value: &[]string{"new value"}[0],
			},
			wantError: false,
		},
		{
			name: "Valid update with value type",
			request: UpdateSystemConfigurationRequest{
				ValueType: &[]string{"number"}[0],
			},
			wantError: false,
		},
		{
			name: "Valid update with description",
			request: UpdateSystemConfigurationRequest{
				Description: &[]string{"Updated description"}[0],
			},
			wantError: false,
		},
		{
			name: "Valid update with category",
			request: UpdateSystemConfigurationRequest{
				Category: &[]string{"NEW_CATEGORY"}[0],
			},
			wantError: false,
		},
		{
			name: "Valid update with is_active",
			request: UpdateSystemConfigurationRequest{
				IsActive: &[]bool{false}[0],
			},
			wantError: false,
		},
		{
			name: "Valid - all nil (no updates)",
			request: UpdateSystemConfigurationRequest{
				Value:       nil,
				ValueType:   nil,
				Description: nil,
				Category:    nil,
				IsActive:    nil,
			},
			wantError: false,
		},
		{
			name: "Invalid - invalid value type",
			request: UpdateSystemConfigurationRequest{
				ValueType: &[]string{"invalid"}[0],
			},
			wantError:   true,
			expectedErr: "value_type must be string, number, boolean, or json",
		},
		{
			name: "Invalid - description too long (501 characters)",
			request: UpdateSystemConfigurationRequest{
				Description: &[]string{string(make([]byte, 501))}[0],
			},
			wantError:   true,
			expectedErr: "description must be less than 500 characters",
		},
		{
			name: "Invalid - category too long (51 characters)",
			request: UpdateSystemConfigurationRequest{
				Category: &[]string{string(make([]byte, 51))}[0],
			},
			wantError:   true,
			expectedErr: "category must be less than 50 characters",
		},
		{
			name: "Valid - description at boundary (500 characters)",
			request: UpdateSystemConfigurationRequest{
				Description: &[]string{string(make([]byte, 500))}[0],
			},
			wantError: false,
		},
		{
			name: "Valid - category at boundary (50 characters)",
			request: UpdateSystemConfigurationRequest{
				Category: &[]string{string(make([]byte, 50))}[0],
			},
			wantError: false,
		},
		{
			name: "Valid - valid value types",
			request: UpdateSystemConfigurationRequest{
				ValueType: &[]string{"string"}[0],
			},
			wantError: false,
		},
		{
			name: "Valid - valid value type boolean",
			request: UpdateSystemConfigurationRequest{
				ValueType: &[]string{"boolean"}[0],
			},
			wantError: false,
		},
		{
			name: "Valid - valid value type json",
			request: UpdateSystemConfigurationRequest{
				ValueType: &[]string{"json"}[0],
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

// TestGetDefaultCPAlignmentThreshold tests the GetDefaultCPAlignmentThreshold function
func TestGetDefaultCPAlignmentThreshold(t *testing.T) {
	threshold := GetDefaultCPAlignmentThreshold()
	if threshold != 60.0 {
		t.Errorf("expected default threshold 60.0, got %f", threshold)
	}
}

// TestGetCPAlignmentThresholdCategory tests the GetCPAlignmentThresholdCategory function
func TestGetCPAlignmentThresholdCategory(t *testing.T) {
	category := GetCPAlignmentThresholdCategory()
	if category != "CURRICULUM" {
		t.Errorf("expected category CURRICULUM, got %s", category)
	}
}

package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// SystemConfiguration represents system-wide configuration settings
// Used for configurable business rules like CP alignment threshold
type SystemConfiguration struct {
	ID            string    `json:"id" db:"id"`
	Key           string    `json:"key" db:"key"`
	Value         string    `json:"value" db:"value"`
	ValueType     string    `json:"value_type" db:"value_type"` // "string", "number", "boolean", "json"
	Description   *string   `json:"description,omitempty" db:"description"`
	Category      string    `json:"category" db:"category"`
	IsSystem      bool      `json:"is_system" db:"is_system"` // System configs cannot be deleted by users
	IsActive      bool      `json:"is_active" db:"is_active"`
	CreatedBy     string    `json:"created_by" db:"created_by"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
	UpdatedBy     *string   `json:"updated_by,omitempty" db:"updated_by"`
}

// NewSystemConfiguration creates a new SystemConfiguration entity
func NewSystemConfiguration(key, value, valueType, category string, isSystem bool, createdBy string) (*SystemConfiguration, error) {
	if key == "" {
		return nil, errors.New("configuration key is required")
	}
	if value == "" {
		return nil, errors.New("configuration value is required")
	}
	if valueType == "" {
		return nil, errors.New("configuration value type is required")
	}
	if category == "" {
		return nil, errors.New("configuration category is required")
	}

	return &SystemConfiguration{
		ID:        uuid.New().String(),
		Key:       key,
		Value:     value,
		ValueType: valueType,
		Category:  category,
		IsSystem:  isSystem,
		IsActive:  true,
		CreatedBy: createdBy,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

// Validate performs business rule validation on the SystemConfiguration
func (sc *SystemConfiguration) Validate() error {
	if sc.Key == "" {
		return errors.New("configuration key is required")
	}
	if sc.Value == "" {
		return errors.New("configuration value is required")
	}
	if sc.ValueType == "" {
		return errors.New("configuration value type is required")
	}
	if sc.Category == "" {
		return errors.New("configuration category is required")
	}
	
	// Validate value type
	validTypes := map[string]bool{
		"string":  true,
		"number":  true,
		"boolean": true,
		"json":    true,
	}
	if !validTypes[sc.ValueType] {
		return errors.New("invalid value type, must be string, number, boolean, or json")
	}
	
	if len(sc.Key) > 100 {
		return errors.New("key must be less than 100 characters")
	}
	if len(sc.Category) > 50 {
		return errors.New("category must be less than 50 characters")
	}
	if sc.Description != nil && len(*sc.Description) > 500 {
		return errors.New("description must be less than 500 characters")
	}
	return nil
}

// Activate marks the system configuration as active
func (sc *SystemConfiguration) Activate() {
	sc.IsActive = true
	sc.UpdatedAt = time.Now()
}

// Deactivate marks the system configuration as inactive
func (sc *SystemConfiguration) Deactivate() {
	sc.IsActive = false
	sc.UpdatedAt = time.Now()
}

// IsSystemConfig checks if this is a system configuration (cannot be deleted by users)
func (sc *SystemConfiguration) IsSystemConfig() bool {
	return sc.IsSystem
}

// CanBeDeleted checks if the configuration can be deleted
func (sc *SystemConfiguration) CanBeDeleted() bool {
	return !sc.IsSystem
}

// CreateSystemConfigurationRequest represents the request to create a system configuration
type CreateSystemConfigurationRequest struct {
	Key         string  `json:"key" binding:"required,min=1,max=100"`
	Value       string  `json:"value" binding:"required"`
	ValueType   string  `json:"value_type" binding:"required,oneof=string number boolean json"`
	Description *string `json:"description,omitempty" binding:"omitempty,max=500"`
	Category    string  `json:"category" binding:"required,min=1,max=50"`
	IsSystem    bool    `json:"is_system"`
}

// Validate performs validation on the request
func (r *CreateSystemConfigurationRequest) Validate() error {
	if r.Key == "" {
		return errors.New("key is required")
	}
	if len(r.Key) > 100 {
		return errors.New("key must be less than 100 characters")
	}
	if r.Value == "" {
		return errors.New("value is required")
	}
	if r.ValueType == "" {
		return errors.New("value_type is required")
	}
	if r.ValueType != "string" && r.ValueType != "number" && r.ValueType != "boolean" && r.ValueType != "json" {
		return errors.New("value_type must be string, number, boolean, or json")
	}
	if r.Category == "" {
		return errors.New("category is required")
	}
	if len(r.Category) > 50 {
		return errors.New("category must be less than 50 characters")
	}
	if r.Description != nil && len(*r.Description) > 500 {
		return errors.New("description must be less than 500 characters")
	}
	return nil
}

// UpdateSystemConfigurationRequest represents the request to update a system configuration
type UpdateSystemConfigurationRequest struct {
	Value       *string `json:"value,omitempty"`
	ValueType   *string `json:"value_type,omitempty" binding:"omitempty,oneof=string number boolean json"`
	Description *string `json:"description,omitempty" binding:"omitempty,max=500"`
	Category    *string `json:"category,omitempty" binding:"omitempty,max=50"`
	IsActive    *bool   `json:"is_active,omitempty"`
}

// Validate performs validation on the update request
func (r *UpdateSystemConfigurationRequest) Validate() error {
	if r.ValueType != nil {
		if *r.ValueType != "string" && *r.ValueType != "number" && *r.ValueType != "boolean" && *r.ValueType != "json" {
			return errors.New("value_type must be string, number, boolean, or json")
		}
	}
	if r.Description != nil && len(*r.Description) > 500 {
		return errors.New("description must be less than 500 characters")
	}
	if r.Category != nil && len(*r.Category) > 50 {
		return errors.New("category must be less than 50 characters")
	}
	return nil
}

// Known configuration keys
const (
	// ConfigCPAlignmentThreshold is the key for CP alignment minimum percentage threshold
	// Default: 60 (as per Sprint 4 requirement BR-004)
	ConfigCPAlignmentThreshold = "CP_ALIGNMENT_THRESHOLD"
)

// GetDefaultCPAlignmentThreshold returns the default CP alignment threshold (60%)
func GetDefaultCPAlignmentThreshold() float64 {
	return 60.0
}

// GetCPAlignmentThresholdCategory returns the category for CP alignment threshold configuration
func GetCPAlignmentThresholdCategory() string {
	return "CURRICULUM"
}
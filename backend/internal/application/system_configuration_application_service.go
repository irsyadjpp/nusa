package application

import (
	"context"
	"fmt"
	"time"

	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/repository"
)

// SystemConfigurationApplicationService orchestrates system configuration use cases
type SystemConfigurationApplicationService struct {
	systemConfigRepo *repository.SystemConfigurationRepository
	userRepo         *repository.UserRepository
}

// NewSystemConfigurationApplicationService creates a new system configuration application service
func NewSystemConfigurationApplicationService(
	systemConfigRepo *repository.SystemConfigurationRepository,
	userRepo *repository.UserRepository,
) *SystemConfigurationApplicationService {
	return &SystemConfigurationApplicationService{
		systemConfigRepo: systemConfigRepo,
		userRepo:         userRepo,
	}
}

// CreateSystemConfigurationCommand represents the command to create a system configuration
type CreateSystemConfigurationCommand struct {
	Key         string
	Value       string
	ValueType   string
	Description *string
	Category    string
	IsSystem    bool
	UserID      string
}

// CreateSystemConfigurationResponse represents the response for creating a system configuration
type CreateSystemConfigurationResponse struct {
	SystemConfigurationID string
}

// CreateSystemConfiguration creates a new system configuration
func (s *SystemConfigurationApplicationService) CreateSystemConfiguration(ctx context.Context, cmd *CreateSystemConfigurationCommand) (*CreateSystemConfigurationResponse, error) {
	// 1. Authorization: Get user
	user, err := s.userRepo.GetByID(ctx, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Only System Admin can create system configurations
	if user.RoleID != domain.RoleSystemAdmin {
		return nil, fmt.Errorf("user does not have permission to create system configurations")
	}

	// 2. Check key uniqueness
	keyExists, err := s.systemConfigRepo.CheckKeyExists(ctx, cmd.Key, "")
	if err != nil {
		return nil, fmt.Errorf("failed to check key uniqueness: %w", err)
	}
	if keyExists {
		return nil, fmt.Errorf("configuration key already exists")
	}

	// 3. Create domain entity
	systemConfig, err := domain.NewSystemConfiguration(cmd.Key, cmd.Value, cmd.ValueType, cmd.Category, cmd.IsSystem, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to create system configuration: %w", err)
	}

	// 4. Apply optional fields
	if cmd.Description != nil {
		systemConfig.Description = cmd.Description
	}

	// 5. Business rule validation
	if err := systemConfig.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	// 6. Persist
	if err := s.systemConfigRepo.CreateSystemConfiguration(ctx, systemConfig); err != nil {
		return nil, fmt.Errorf("failed to create system configuration: %w", err)
	}

	return &CreateSystemConfigurationResponse{
		SystemConfigurationID: systemConfig.ID,
	}, nil
}

// UpdateSystemConfigurationCommand represents the command to update a system configuration
type UpdateSystemConfigurationCommand struct {
	SystemConfigurationID string
	Value                 *string
	ValueType             *string
	Description           *string
	Category              *string
	IsActive              *bool
	UserID                string
}

// UpdateSystemConfigurationResponse represents the response for updating a system configuration
type UpdateSystemConfigurationResponse struct {
	Success bool
}

// UpdateSystemConfiguration updates a system configuration
func (s *SystemConfigurationApplicationService) UpdateSystemConfiguration(ctx context.Context, cmd *UpdateSystemConfigurationCommand) (*UpdateSystemConfigurationResponse, error) {
	// 1. Authorization: Get user
	user, err := s.userRepo.GetByID(ctx, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// System Admin can update all configurations
	// Curriculum Admin can only update non-system configurations
	config, err := s.systemConfigRepo.GetSystemConfigurationByID(ctx, cmd.SystemConfigurationID)
	if err != nil {
		return nil, fmt.Errorf("system configuration not found: %w", err)
	}

	if user.RoleID != domain.RoleSystemAdmin {
		if config.IsSystem {
			return nil, fmt.Errorf("user does not have permission to update system configurations")
		}
		if user.RoleID != domain.RoleCurriculumAdmin {
			return nil, fmt.Errorf("user does not have permission to update system configurations")
		}
	}

	// 2. Apply updates
	updatedBy := cmd.UserID
	if cmd.Value != nil {
		config.Value = *cmd.Value
	}
	if cmd.ValueType != nil {
		config.ValueType = *cmd.ValueType
	}
	if cmd.Description != nil {
		config.Description = cmd.Description
	}
	if cmd.Category != nil {
		config.Category = *cmd.Category
	}
	if cmd.IsActive != nil {
		config.IsActive = *cmd.IsActive
	}
	config.UpdatedBy = &updatedBy
	config.UpdatedAt = time.Now()

	// 3. Business rule validation
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	// 4. Persist
	if err := s.systemConfigRepo.UpdateSystemConfiguration(ctx, config); err != nil {
		return nil, fmt.Errorf("failed to update system configuration: %w", err)
	}

	return &UpdateSystemConfigurationResponse{
		Success: true,
	}, nil
}

// DeleteSystemConfigurationCommand represents the command to delete a system configuration
type DeleteSystemConfigurationCommand struct {
	SystemConfigurationID string
	UserID                 string
}

// DeleteSystemConfigurationResponse represents the response for deleting a system configuration
type DeleteSystemConfigurationResponse struct {
	Success bool
}

// DeleteSystemConfiguration deletes a system configuration
func (s *SystemConfigurationApplicationService) DeleteSystemConfiguration(ctx context.Context, cmd *DeleteSystemConfigurationCommand) (*DeleteSystemConfigurationResponse, error) {
	// 1. Authorization: Get user
	user, err := s.userRepo.GetByID(ctx, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Only System Admin can delete system configurations
	if user.RoleID != domain.RoleSystemAdmin {
		return nil, fmt.Errorf("user does not have permission to delete system configurations")
	}

	// 2. Persist
	if err := s.systemConfigRepo.DeleteSystemConfiguration(ctx, cmd.SystemConfigurationID); err != nil {
		return nil, fmt.Errorf("failed to delete system configuration: %w", err)
	}

	return &DeleteSystemConfigurationResponse{
		Success: true,
	}, nil
}

// GetSystemConfigurationCommand represents the command to get a system configuration
type GetSystemConfigurationCommand struct {
	SystemConfigurationID string
	UserID                string
}

// GetSystemConfigurationResponse represents the response for getting a system configuration
type GetSystemConfigurationResponse struct {
	*domain.SystemConfiguration
}

// GetSystemConfiguration retrieves a system configuration
func (s *SystemConfigurationApplicationService) GetSystemConfiguration(ctx context.Context, cmd *GetSystemConfigurationCommand) (*GetSystemConfigurationResponse, error) {
	// 1. Authorization: Get user
	user, err := s.userRepo.GetByID(ctx, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// System Admin and Curriculum Admin can read
	if user.RoleID != domain.RoleSystemAdmin && user.RoleID != domain.RoleCurriculumAdmin {
		return nil, fmt.Errorf("user does not have permission to view system configurations")
	}

	// 2. Load system configuration
	config, err := s.systemConfigRepo.GetSystemConfigurationByID(ctx, cmd.SystemConfigurationID)
	if err != nil {
		return nil, fmt.Errorf("system configuration not found: %w", err)
	}

	return &GetSystemConfigurationResponse{
		SystemConfiguration: config,
	}, nil
}

// GetSystemConfigurationByKeyCommand represents the command to get a system configuration by key
type GetSystemConfigurationByKeyCommand struct {
	Key    string
	UserID string
}

// GetSystemConfigurationByKeyResponse represents the response for getting a system configuration by key
type GetSystemConfigurationByKeyResponse struct {
	*domain.SystemConfiguration
}

// GetSystemConfigurationByKey retrieves a system configuration by key
func (s *SystemConfigurationApplicationService) GetSystemConfigurationByKey(ctx context.Context, cmd *GetSystemConfigurationByKeyCommand) (*GetSystemConfigurationByKeyResponse, error) {
	// 1. Authorization: Get user
	user, err := s.userRepo.GetByID(ctx, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// System Admin and Curriculum Admin can read
	if user.RoleID != domain.RoleSystemAdmin && user.RoleID != domain.RoleCurriculumAdmin {
		return nil, fmt.Errorf("user does not have permission to view system configurations")
	}

	// 2. Load system configuration
	config, err := s.systemConfigRepo.GetSystemConfigurationByKey(ctx, cmd.Key)
	if err != nil {
		return nil, fmt.Errorf("system configuration not found: %w", err)
	}

	return &GetSystemConfigurationByKeyResponse{
		SystemConfiguration: config,
	}, nil
}

// ListSystemConfigurationsCommand represents the command to list system configurations
type ListSystemConfigurationsCommand struct {
	Category   *string
	ActiveOnly bool
	UserID     string
}

// ListSystemConfigurationsResponse represents the response for listing system configurations
type ListSystemConfigurationsResponse struct {
	SystemConfigurations []*domain.SystemConfiguration
}

// ListSystemConfigurations retrieves system configurations
func (s *SystemConfigurationApplicationService) ListSystemConfigurations(ctx context.Context, cmd *ListSystemConfigurationsCommand) (*ListSystemConfigurationsResponse, error) {
	// 1. Authorization: Get user
	user, err := s.userRepo.GetByID(ctx, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// System Admin and Curriculum Admin can read
	if user.RoleID != domain.RoleSystemAdmin && user.RoleID != domain.RoleCurriculumAdmin {
		return nil, fmt.Errorf("user does not have permission to view system configurations")
	}

	// 2. Load system configurations
	var configs []*domain.SystemConfiguration
	if cmd.Category != nil {
		configs, err = s.systemConfigRepo.GetSystemConfigurationsByCategory(ctx, *cmd.Category)
	} else if cmd.ActiveOnly {
		configs, err = s.systemConfigRepo.GetActiveSystemConfigurations(ctx)
	} else {
		configs, err = s.systemConfigRepo.GetAllSystemConfigurations(ctx)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to load system configurations: %w", err)
	}

	return &ListSystemConfigurationsResponse{
		SystemConfigurations: configs,
	}, nil
}
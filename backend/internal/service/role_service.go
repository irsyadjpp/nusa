package service

import (
	"context"
	"fmt"

	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/repository"
)

// RoleService handles business logic for role operations
type RoleService struct {
	roleRepo repository.RoleRepositoryInterface
}

// NewRoleService creates a new role service
func NewRoleService(roleRepo repository.RoleRepositoryInterface) *RoleService {
	return &RoleService{
		roleRepo: roleRepo,
	}
}

// CreateRole creates a new role
func (s *RoleService) CreateRole(ctx context.Context, req *domain.CreateRoleRequest) (*domain.Role, error) {
	// Check if role name already exists
	_, err := s.roleRepo.GetByName(ctx, req.Name)
	if err == nil {
		return nil, fmt.Errorf("role with name %s already exists", req.Name)
	}

	role := &domain.Role{
		ID:          domain.NewID(),
		Name:        req.Name,
		Description: req.Description,
		IsActive:    true,
	}

	if err := s.roleRepo.Create(ctx, role); err != nil {
		return nil, fmt.Errorf("failed to create role: %w", err)
	}

	return role, nil
}

// GetRole retrieves a role by ID
func (s *RoleService) GetRole(ctx context.Context, id string) (*domain.Role, error) {
	return s.roleRepo.GetByID(ctx, id)
}

// GetRoleByName retrieves a role by name
func (s *RoleService) GetRoleByName(ctx context.Context, name string) (*domain.Role, error) {
	return s.roleRepo.GetByName(ctx, name)
}

// ListRoles retrieves all roles with optional filters
func (s *RoleService) ListRoles(ctx context.Context, isActive *bool) ([]*domain.Role, error) {
	return s.roleRepo.List(ctx, isActive)
}

// UpdateRole updates role information
func (s *RoleService) UpdateRole(ctx context.Context, id string, req *domain.UpdateRoleRequest) (*domain.Role, error) {
	role, err := s.roleRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("role not found")
	}

	if req.Name != nil {
		// Check if new name already exists (excluding current role)
		existingRole, err := s.roleRepo.GetByName(ctx, *req.Name)
		if err == nil && existingRole.ID != id {
			return nil, fmt.Errorf("role with name %s already exists", *req.Name)
		}
		role.Name = *req.Name
	}

	if req.Description != nil {
		role.Description = req.Description
	}

	if req.IsActive != nil {
		role.IsActive = *req.IsActive
	}

	if err := s.roleRepo.Update(ctx, role); err != nil {
		return nil, fmt.Errorf("failed to update role: %w", err)
	}

	return role, nil
}

// DeleteRole soft deletes a role (sets is_active to false)
func (s *RoleService) DeleteRole(ctx context.Context, id string) error {
	role, err := s.roleRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("role not found")
	}

	// Prevent deletion of system roles
	if role.Name == domain.RoleSystemAdmin || role.Name == domain.RoleSchoolAdmin || role.Name == domain.RoleTeacher {
		return fmt.Errorf("cannot delete system role: %s", role.Name)
	}

	role.IsActive = false
	if err := s.roleRepo.Update(ctx, role); err != nil {
		return fmt.Errorf("failed to delete role: %w", err)
	}

	return nil
}

// AddPermission adds a permission to a role
func (s *RoleService) AddPermission(ctx context.Context, roleID, resource, action string) error {
	// Verify role exists
	role, err := s.roleRepo.GetByID(ctx, roleID)
	if err != nil {
		return fmt.Errorf("role not found")
	}

	if !role.IsActive {
		return fmt.Errorf("role is not active")
	}

	return s.roleRepo.AddPermission(ctx, roleID, resource, action)
}

// RemovePermission removes a permission from a role
func (s *RoleService) RemovePermission(ctx context.Context, roleID, resource, action string) error {
	return s.roleRepo.RemovePermission(ctx, roleID, resource, action)
}

// GetPermissions retrieves all permissions for a role
func (s *RoleService) GetPermissions(ctx context.Context, roleID string) ([]*domain.Permission, error) {
	return s.roleRepo.GetPermissions(ctx, roleID)
}

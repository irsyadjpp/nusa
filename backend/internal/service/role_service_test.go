package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/nusa/backend/internal/domain"
)

func TestRoleService_CreateRole_Success(t *testing.T) {
	mockRoleRepo := new(MockRoleRepository)
	service := NewRoleService(mockRoleRepo)

	description := "A new role"
	req := &domain.CreateRoleRequest{
		Name:        "New Role",
		Description: &description,
	}

	mockRoleRepo.On("GetByName", mock.Anything, "New Role").Return(nil, errors.New("not found"))
	mockRoleRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.Role")).Return(nil)

	result, err := service.CreateRole(context.Background(), req)

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "New Role", result.Name)
	assert.Equal(t, "A new role", *result.Description)
	assert.True(t, result.IsActive)

	mockRoleRepo.AssertExpectations(t)
}

func TestRoleService_CreateRole_NameExists(t *testing.T) {
	mockRoleRepo := new(MockRoleRepository)
	service := NewRoleService(mockRoleRepo)

	description := "An existing role"
	req := &domain.CreateRoleRequest{
		Name:        "Existing Role",
		Description: &description,
	}

	existingRole := &domain.Role{
		ID:   "role-1",
		Name: "Existing Role",
	}

	mockRoleRepo.On("GetByName", mock.Anything, "Existing Role").Return(existingRole, nil)

	result, err := service.CreateRole(context.Background(), req)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "already exists")

	mockRoleRepo.AssertExpectations(t)
}

func TestRoleService_CreateRole_CreateError(t *testing.T) {
	mockRoleRepo := new(MockRoleRepository)
	service := NewRoleService(mockRoleRepo)

	description := "A new role"
	req := &domain.CreateRoleRequest{
		Name:        "New Role",
		Description: &description,
	}

	mockRoleRepo.On("GetByName", mock.Anything, "New Role").Return(nil, errors.New("not found"))
	mockRoleRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.Role")).Return(errors.New("database error"))

	result, err := service.CreateRole(context.Background(), req)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to create role")

	mockRoleRepo.AssertExpectations(t)
}

func TestRoleService_GetRole_Success(t *testing.T) {
	mockRoleRepo := new(MockRoleRepository)
	service := NewRoleService(mockRoleRepo)

	role := &domain.Role{
		ID:   "role-1",
		Name: "Teacher",
	}

	mockRoleRepo.On("GetByID", mock.Anything, "role-1").Return(role, nil)

	result, err := service.GetRole(context.Background(), "role-1")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, role.ID, result.ID)

	mockRoleRepo.AssertExpectations(t)
}

func TestRoleService_GetRole_NotFound(t *testing.T) {
	mockRoleRepo := new(MockRoleRepository)
	service := NewRoleService(mockRoleRepo)

	mockRoleRepo.On("GetByID", mock.Anything, "role-1").Return(nil, errors.New("not found"))

	result, err := service.GetRole(context.Background(), "role-1")

	require.Error(t, err)
	assert.Nil(t, result)

	mockRoleRepo.AssertExpectations(t)
}

func TestRoleService_GetRoleByName_Success(t *testing.T) {
	mockRoleRepo := new(MockRoleRepository)
	service := NewRoleService(mockRoleRepo)

	role := &domain.Role{
		ID:   "role-1",
		Name: "Teacher",
	}

	mockRoleRepo.On("GetByName", mock.Anything, "Teacher").Return(role, nil)

	result, err := service.GetRoleByName(context.Background(), "Teacher")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, role.ID, result.ID)

	mockRoleRepo.AssertExpectations(t)
}

func TestRoleService_ListRoles_Success(t *testing.T) {
	mockRoleRepo := new(MockRoleRepository)
	service := NewRoleService(mockRoleRepo)

	roles := []*domain.Role{
		{ID: "role-1", Name: "Teacher"},
		{ID: "role-2", Name: "Admin"},
	}

	mockRoleRepo.On("List", mock.Anything, (*bool)(nil)).Return(roles, nil)

	result, err := service.ListRoles(context.Background(), nil)

	require.NoError(t, err)
	assert.Len(t, result, 2)

	mockRoleRepo.AssertExpectations(t)
}

func TestRoleService_ListRoles_WithFilter(t *testing.T) {
	mockRoleRepo := new(MockRoleRepository)
	service := NewRoleService(mockRoleRepo)

	isActive := true

	roles := []*domain.Role{
		{ID: "role-1", Name: "Teacher", IsActive: true},
	}

	mockRoleRepo.On("List", mock.Anything, &isActive).Return(roles, nil)

	result, err := service.ListRoles(context.Background(), &isActive)

	require.NoError(t, err)
	assert.Len(t, result, 1)

	mockRoleRepo.AssertExpectations(t)
}

func TestRoleService_UpdateRole_Success(t *testing.T) {
	mockRoleRepo := new(MockRoleRepository)
	service := NewRoleService(mockRoleRepo)

	description := "Old description"
	role := &domain.Role{
		ID:          "role-1",
		Name:        "Old Name",
		Description: &description,
	}

	newName := "New Name"
	newDescription := "New description"

	req := &domain.UpdateRoleRequest{
		Name:        &newName,
		Description: &newDescription,
	}

	mockRoleRepo.On("GetByID", mock.Anything, "role-1").Return(role, nil)
	mockRoleRepo.On("GetByName", mock.Anything, "New Name").Return(nil, errors.New("not found"))
	mockRoleRepo.On("Update", mock.Anything, mock.AnythingOfType("*domain.Role")).Return(nil)

	result, err := service.UpdateRole(context.Background(), "role-1", req)

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, newName, result.Name)
	assert.Equal(t, newDescription, *result.Description)

	mockRoleRepo.AssertExpectations(t)
}

func TestRoleService_UpdateRole_NameExists(t *testing.T) {
	mockRoleRepo := new(MockRoleRepository)
	service := NewRoleService(mockRoleRepo)

	description := "Old description"
	role := &domain.Role{
		ID:          "role-1",
		Name:        "Old Name",
		Description: &description,
	}

	existingRole := &domain.Role{
		ID:   "role-2",
		Name: "Existing Name",
	}

	newName := "Existing Name"

	req := &domain.UpdateRoleRequest{
		Name: &newName,
	}

	mockRoleRepo.On("GetByID", mock.Anything, "role-1").Return(role, nil)
	mockRoleRepo.On("GetByName", mock.Anything, "Existing Name").Return(existingRole, nil)

	result, err := service.UpdateRole(context.Background(), "role-1", req)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "already exists")

	mockRoleRepo.AssertExpectations(t)
}

func TestRoleService_UpdateRole_SameName(t *testing.T) {
	mockRoleRepo := new(MockRoleRepository)
	service := NewRoleService(mockRoleRepo)

	description := "Old description"
	role := &domain.Role{
		ID:          "role-1",
		Name:        "Same Name",
		Description: &description,
	}

	newName := "Same Name"

	req := &domain.UpdateRoleRequest{
		Name: &newName,
	}

	mockRoleRepo.On("GetByID", mock.Anything, "role-1").Return(role, nil)
	mockRoleRepo.On("GetByName", mock.Anything, "Same Name").Return(role, nil)
	mockRoleRepo.On("Update", mock.Anything, mock.AnythingOfType("*domain.Role")).Return(nil)

	result, err := service.UpdateRole(context.Background(), "role-1", req)

	require.NoError(t, err)
	assert.NotNil(t, result)

	mockRoleRepo.AssertExpectations(t)
}

func TestRoleService_UpdateRole_NotFound(t *testing.T) {
	mockRoleRepo := new(MockRoleRepository)
	service := NewRoleService(mockRoleRepo)

	req := &domain.UpdateRoleRequest{}

	mockRoleRepo.On("GetByID", mock.Anything, "role-1").Return(nil, errors.New("not found"))

	result, err := service.UpdateRole(context.Background(), "role-1", req)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "role not found")

	mockRoleRepo.AssertExpectations(t)
}

func TestRoleService_DeleteRole_Success(t *testing.T) {
	mockRoleRepo := new(MockRoleRepository)
	service := NewRoleService(mockRoleRepo)

	role := &domain.Role{
		ID:   "role-1",
		Name: "Custom Role",
	}

	mockRoleRepo.On("GetByID", mock.Anything, "role-1").Return(role, nil)
	mockRoleRepo.On("Update", mock.Anything, mock.AnythingOfType("*domain.Role")).Return(nil)

	err := service.DeleteRole(context.Background(), "role-1")

	require.NoError(t, err)

	mockRoleRepo.AssertExpectations(t)
}

func TestRoleService_DeleteRole_SystemRole(t *testing.T) {
	mockRoleRepo := new(MockRoleRepository)
	service := NewRoleService(mockRoleRepo)

	role := &domain.Role{
		ID:   "role-1",
		Name: domain.RoleSystemAdmin,
	}

	mockRoleRepo.On("GetByID", mock.Anything, "role-1").Return(role, nil)

	err := service.DeleteRole(context.Background(), "role-1")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "cannot delete system role")

	mockRoleRepo.AssertExpectations(t)
}

func TestRoleService_DeleteRole_NotFound(t *testing.T) {
	mockRoleRepo := new(MockRoleRepository)
	service := NewRoleService(mockRoleRepo)

	mockRoleRepo.On("GetByID", mock.Anything, "role-1").Return(nil, errors.New("not found"))

	err := service.DeleteRole(context.Background(), "role-1")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "role not found")

	mockRoleRepo.AssertExpectations(t)
}

func TestRoleService_AddPermission_Success(t *testing.T) {
	mockRoleRepo := new(MockRoleRepository)
	service := NewRoleService(mockRoleRepo)

	role := &domain.Role{
		ID:       "role-1",
		Name:     "Teacher",
		IsActive: true,
	}

	mockRoleRepo.On("GetByID", mock.Anything, "role-1").Return(role, nil)
	mockRoleRepo.On("AddPermission", mock.Anything, "role-1", "assessment", "create").Return(nil)

	err := service.AddPermission(context.Background(), "role-1", "assessment", "create")

	require.NoError(t, err)

	mockRoleRepo.AssertExpectations(t)
}

func TestRoleService_AddPermission_RoleNotFound(t *testing.T) {
	mockRoleRepo := new(MockRoleRepository)
	service := NewRoleService(mockRoleRepo)

	mockRoleRepo.On("GetByID", mock.Anything, "role-1").Return(nil, errors.New("not found"))

	err := service.AddPermission(context.Background(), "role-1", "assessment", "create")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "role not found")

	mockRoleRepo.AssertExpectations(t)
}

func TestRoleService_AddPermission_RoleInactive(t *testing.T) {
	mockRoleRepo := new(MockRoleRepository)
	service := NewRoleService(mockRoleRepo)

	role := &domain.Role{
		ID:       "role-1",
		Name:     "Teacher",
		IsActive: false,
	}

	mockRoleRepo.On("GetByID", mock.Anything, "role-1").Return(role, nil)

	err := service.AddPermission(context.Background(), "role-1", "assessment", "create")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "role is not active")

	mockRoleRepo.AssertExpectations(t)
}

func TestRoleService_RemovePermission_Success(t *testing.T) {
	mockRoleRepo := new(MockRoleRepository)
	service := NewRoleService(mockRoleRepo)

	mockRoleRepo.On("RemovePermission", mock.Anything, "role-1", "assessment", "create").Return(nil)

	err := service.RemovePermission(context.Background(), "role-1", "assessment", "create")

	require.NoError(t, err)

	mockRoleRepo.AssertExpectations(t)
}

func TestRoleService_GetPermissions_Success(t *testing.T) {
	mockRoleRepo := new(MockRoleRepository)
	service := NewRoleService(mockRoleRepo)

	permissions := []*domain.Permission{
		{Resource: "assessment", Action: "create"},
		{Resource: "assessment", Action: "read"},
	}

	mockRoleRepo.On("GetPermissions", mock.Anything, "role-1").Return(permissions, nil)

	result, err := service.GetPermissions(context.Background(), "role-1")

	require.NoError(t, err)
	assert.Len(t, result, 2)

	mockRoleRepo.AssertExpectations(t)
}

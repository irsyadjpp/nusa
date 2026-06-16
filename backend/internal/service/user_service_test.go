package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"

	"github.com/nusa/backend/internal/domain"
)

// MockUserRepository is a mock implementation of UserRepositoryInterface
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(ctx context.Context, user *domain.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) Update(ctx context.Context, user *domain.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) UpdateStatus(ctx context.Context, id string, isActive bool, lockedUntil *string, failedAttempts int) error {
	args := m.Called(ctx, id, isActive, lockedUntil, failedAttempts)
	return args.Error(0)
}

func (m *MockUserRepository) List(ctx context.Context, schoolID, roleID *string, isActive *bool, limit, offset int) ([]*domain.User, error) {
	args := m.Called(ctx, schoolID, roleID, isActive, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.User), args.Error(1)
}

func (m *MockUserRepository) Count(ctx context.Context, schoolID, roleID *string, isActive *bool) (int, error) {
	args := m.Called(ctx, schoolID, roleID, isActive)
	return args.Int(0), args.Error(1)
}

func (m *MockUserRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockUserRepository) GetUserSchoolID(ctx context.Context, userID string) (*string, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*string), args.Error(1)
}

func (m *MockUserRepository) ListUsersBySchool(ctx context.Context, schoolID string, limit, offset int) ([]*domain.User, error) {
	args := m.Called(ctx, schoolID, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.User), args.Error(1)
}

// MockRoleRepository is a mock implementation of RoleRepositoryInterface
type MockRoleRepository struct {
	mock.Mock
}

func (m *MockRoleRepository) Create(ctx context.Context, role *domain.Role) error {
	args := m.Called(ctx, role)
	return args.Error(0)
}

func (m *MockRoleRepository) GetByID(ctx context.Context, id string) (*domain.Role, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Role), args.Error(1)
}

func (m *MockRoleRepository) GetByName(ctx context.Context, name string) (*domain.Role, error) {
	args := m.Called(ctx, name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Role), args.Error(1)
}

func (m *MockRoleRepository) Update(ctx context.Context, role *domain.Role) error {
	args := m.Called(ctx, role)
	return args.Error(0)
}

func (m *MockRoleRepository) List(ctx context.Context, isActive *bool) ([]*domain.Role, error) {
	args := m.Called(ctx, isActive)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Role), args.Error(1)
}

func (m *MockRoleRepository) AddPermission(ctx context.Context, roleID, resource, action string) error {
	args := m.Called(ctx, roleID, resource, action)
	return args.Error(0)
}

func (m *MockRoleRepository) GetPermissions(ctx context.Context, roleID string) ([]*domain.Permission, error) {
	args := m.Called(ctx, roleID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Permission), args.Error(1)
}

func (m *MockRoleRepository) RemovePermission(ctx context.Context, roleID, resource, action string) error {
	args := m.Called(ctx, roleID, resource, action)
	return args.Error(0)
}

func TestUserService_Register_Success(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	mockRoleRepo := new(MockRoleRepository)
	service := NewUserService(mockUserRepo, mockRoleRepo)

	schoolID := "school-1"
	req := &domain.CreateUserRequest{
		Email:    "test@example.com",
		Password: "password123",
		RoleID:   "role-1",
		Name:     "Test User",
		SchoolID: &schoolID,
	}

	role := &domain.Role{
		ID:       "role-1",
		Name:     "Teacher",
		IsActive: true,
	}

	mockRoleRepo.On("GetByID", mock.Anything, "role-1").Return(role, nil)
	mockUserRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil)

	result, err := service.Register(context.Background(), req, "creator-1")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, req.Email, result.Email)
	assert.Equal(t, req.Name, result.Name)
	assert.Equal(t, req.RoleID, result.RoleID)
	assert.Equal(t, req.SchoolID, result.SchoolID)
	assert.True(t, result.IsActive)
	assert.NotEmpty(t, result.PasswordHash)

	mockRoleRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestUserService_Register_RoleNotFound(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	mockRoleRepo := new(MockRoleRepository)
	service := NewUserService(mockUserRepo, mockRoleRepo)

	req := &domain.CreateUserRequest{
		Email:    "test@example.com",
		Password: "password123",
		RoleID:   "role-1",
		Name:     "Test User",
	}

	mockRoleRepo.On("GetByID", mock.Anything, "role-1").Return(nil, errors.New("role not found"))

	result, err := service.Register(context.Background(), req, "creator-1")

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "invalid role")

	mockRoleRepo.AssertExpectations(t)
}

func TestUserService_Register_RoleInactive(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	mockRoleRepo := new(MockRoleRepository)
	service := NewUserService(mockUserRepo, mockRoleRepo)

	req := &domain.CreateUserRequest{
		Email:    "test@example.com",
		Password: "password123",
		RoleID:   "role-1",
		Name:     "Test User",
	}

	role := &domain.Role{
		ID:       "role-1",
		Name:     "Teacher",
		IsActive: false,
	}

	mockRoleRepo.On("GetByID", mock.Anything, "role-1").Return(role, nil)

	result, err := service.Register(context.Background(), req, "creator-1")

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "role is not active")

	mockRoleRepo.AssertExpectations(t)
}

func TestUserService_Register_CreateError(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	mockRoleRepo := new(MockRoleRepository)
	service := NewUserService(mockUserRepo, mockRoleRepo)

	req := &domain.CreateUserRequest{
		Email:    "test@example.com",
		Password: "password123",
		RoleID:   "role-1",
		Name:     "Test User",
	}

	role := &domain.Role{
		ID:       "role-1",
		Name:     "Teacher",
		IsActive: true,
	}

	mockRoleRepo.On("GetByID", mock.Anything, "role-1").Return(role, nil)
	mockUserRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.User")).Return(errors.New("database error"))

	result, err := service.Register(context.Background(), req, "creator-1")

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to create user")

	mockRoleRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestUserService_ValidateCredentials_Success(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	mockRoleRepo := new(MockRoleRepository)
	service := NewUserService(mockUserRepo, mockRoleRepo)

	// Generate a real bcrypt hash for "password"
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	user := &domain.User{
		ID:                  "user-1",
		Email:               "test@example.com",
		PasswordHash:        string(passwordHash),
		IsActive:            true,
		FailedLoginAttempts: 0,
	}

	mockUserRepo.On("GetByEmail", mock.Anything, "test@example.com").Return(user, nil)
	mockUserRepo.On("UpdateStatus", mock.Anything, "user-1", true, mock.Anything, mock.Anything).Return(nil).Maybe()

	result, err := service.ValidateCredentials(context.Background(), "test@example.com", "password")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, user.ID, result.ID)

	mockUserRepo.AssertExpectations(t)
}

func TestUserService_ValidateCredentials_UserNotFound(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	mockRoleRepo := new(MockRoleRepository)
	service := NewUserService(mockUserRepo, mockRoleRepo)

	mockUserRepo.On("GetByEmail", mock.Anything, "test@example.com").Return(nil, errors.New("user not found"))

	result, err := service.ValidateCredentials(context.Background(), "test@example.com", "password")

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "invalid credentials")

	mockUserRepo.AssertExpectations(t)
}

func TestUserService_ValidateCredentials_AccountLocked(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	mockRoleRepo := new(MockRoleRepository)
	service := NewUserService(mockUserRepo, mockRoleRepo)

	lockTime := time.Now().Add(1 * time.Hour)
	user := &domain.User{
		ID:           "user-1",
		Email:        "test@example.com",
		PasswordHash: "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy",
		IsActive:     true,
		LockedUntil:  &lockTime,
	}

	mockUserRepo.On("GetByEmail", mock.Anything, "test@example.com").Return(user, nil)

	result, err := service.ValidateCredentials(context.Background(), "test@example.com", "password")

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "account is locked")

	mockUserRepo.AssertExpectations(t)
}

func TestUserService_ValidateCredentials_AccountInactive(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	mockRoleRepo := new(MockRoleRepository)
	service := NewUserService(mockUserRepo, mockRoleRepo)

	user := &domain.User{
		ID:           "user-1",
		Email:        "test@example.com",
		PasswordHash: "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy",
		IsActive:     false,
	}

	mockUserRepo.On("GetByEmail", mock.Anything, "test@example.com").Return(user, nil)

	result, err := service.ValidateCredentials(context.Background(), "test@example.com", "password")

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "account is not active")

	mockUserRepo.AssertExpectations(t)
}

func TestUserService_ValidateCredentials_WrongPassword(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	mockRoleRepo := new(MockRoleRepository)
	service := NewUserService(mockUserRepo, mockRoleRepo)

	user := &domain.User{
		ID:                  "user-1",
		Email:               "test@example.com",
		PasswordHash:        "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy",
		IsActive:            true,
		FailedLoginAttempts: 0,
	}

	mockUserRepo.On("GetByEmail", mock.Anything, "test@example.com").Return(user, nil)
	mockUserRepo.On("GetByID", mock.Anything, "user-1").Return(user, nil)
	mockUserRepo.On("UpdateStatus", mock.Anything, "user-1", true, mock.Anything, mock.Anything).Return(nil)

	result, err := service.ValidateCredentials(context.Background(), "test@example.com", "wrongpassword")

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "invalid credentials")

	mockUserRepo.AssertExpectations(t)
}

func TestUserService_GetUser_Success(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	mockRoleRepo := new(MockRoleRepository)
	service := NewUserService(mockUserRepo, mockRoleRepo)

	user := &domain.User{
		ID:    "user-1",
		Email: "test@example.com",
	}

	mockUserRepo.On("GetByID", mock.Anything, "user-1").Return(user, nil)

	result, err := service.GetUser(context.Background(), "user-1")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, user.ID, result.ID)

	mockUserRepo.AssertExpectations(t)
}

func TestUserService_GetUser_NotFound(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	mockRoleRepo := new(MockRoleRepository)
	service := NewUserService(mockUserRepo, mockRoleRepo)

	mockUserRepo.On("GetByID", mock.Anything, "user-1").Return(nil, errors.New("user not found"))

	result, err := service.GetUser(context.Background(), "user-1")

	require.Error(t, err)
	assert.Nil(t, result)

	mockUserRepo.AssertExpectations(t)
}

func TestUserService_ListUsers_Success(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	mockRoleRepo := new(MockRoleRepository)
	service := NewUserService(mockUserRepo, mockRoleRepo)

	users := []*domain.User{
		{ID: "user-1", Email: "user1@example.com"},
		{ID: "user-2", Email: "user2@example.com"},
	}

	mockUserRepo.On("List", mock.Anything, (*string)(nil), (*string)(nil), (*bool)(nil), 10, 0).Return(users, nil)
	mockUserRepo.On("Count", mock.Anything, (*string)(nil), (*string)(nil), (*bool)(nil)).Return(2, nil)

	result, total, err := service.ListUsers(context.Background(), nil, nil, nil, 1, 10)

	require.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, 2, total)

	mockUserRepo.AssertExpectations(t)
}

func TestUserService_ListUsers_WithFilters(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	mockRoleRepo := new(MockRoleRepository)
	service := NewUserService(mockUserRepo, mockRoleRepo)

	schoolID := "school-1"
	roleID := "role-1"
	isActive := true

	users := []*domain.User{
		{ID: "user-1", Email: "user1@example.com"},
	}

	mockUserRepo.On("List", mock.Anything, &schoolID, &roleID, &isActive, 10, 0).Return(users, nil)
	mockUserRepo.On("Count", mock.Anything, &schoolID, &roleID, &isActive).Return(1, nil)

	result, total, err := service.ListUsers(context.Background(), &schoolID, &roleID, &isActive, 1, 10)

	require.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, 1, total)

	mockUserRepo.AssertExpectations(t)
}

func TestUserService_ListUsers_Error(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	mockRoleRepo := new(MockRoleRepository)
	service := NewUserService(mockUserRepo, mockRoleRepo)

	mockUserRepo.On("List", mock.Anything, (*string)(nil), (*string)(nil), (*bool)(nil), 10, 0).Return(nil, errors.New("database error"))

	result, total, err := service.ListUsers(context.Background(), nil, nil, nil, 1, 10)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, 0, total)

	mockUserRepo.AssertExpectations(t)
}

func TestUserService_UpdateUser_Success(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	mockRoleRepo := new(MockRoleRepository)
	service := NewUserService(mockUserRepo, mockRoleRepo)

	user := &domain.User{
		ID:     "user-1",
		Name:   "Old Name",
		RoleID: "role-1",
	}

	role := &domain.Role{
		ID:       "role-2",
		Name:     "Admin",
		IsActive: true,
	}

	newName := "New Name"
	newRoleID := "role-2"

	req := &domain.UpdateUserRequest{
		Name:   &newName,
		RoleID: &newRoleID,
	}

	mockUserRepo.On("GetByID", mock.Anything, "user-1").Return(user, nil)
	mockRoleRepo.On("GetByID", mock.Anything, "role-2").Return(role, nil)
	mockUserRepo.On("Update", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil)

	result, err := service.UpdateUser(context.Background(), "user-1", req, "updater-1")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, newName, result.Name)
	assert.Equal(t, newRoleID, result.RoleID)

	mockUserRepo.AssertExpectations(t)
	mockRoleRepo.AssertExpectations(t)
}

func TestUserService_UpdateUser_UserNotFound(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	mockRoleRepo := new(MockRoleRepository)
	service := NewUserService(mockUserRepo, mockRoleRepo)

	req := &domain.UpdateUserRequest{}

	mockUserRepo.On("GetByID", mock.Anything, "user-1").Return(nil, errors.New("user not found"))

	result, err := service.UpdateUser(context.Background(), "user-1", req, "updater-1")

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "user not found")

	mockUserRepo.AssertExpectations(t)
}

func TestUserService_UpdateUser_RoleInactive(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	mockRoleRepo := new(MockRoleRepository)
	service := NewUserService(mockUserRepo, mockRoleRepo)

	user := &domain.User{
		ID:     "user-1",
		Name:   "Old Name",
		RoleID: "role-1",
	}

	newRoleID := "role-2"

	req := &domain.UpdateUserRequest{
		RoleID: &newRoleID,
	}

	role := &domain.Role{
		ID:       "role-2",
		Name:     "Admin",
		IsActive: false,
	}

	mockUserRepo.On("GetByID", mock.Anything, "user-1").Return(user, nil)
	mockRoleRepo.On("GetByID", mock.Anything, "role-2").Return(role, nil)

	result, err := service.UpdateUser(context.Background(), "user-1", req, "updater-1")

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "role is not active")

	mockUserRepo.AssertExpectations(t)
	mockRoleRepo.AssertExpectations(t)
}

func TestUserService_UpdateUserStatus_Success(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	mockRoleRepo := new(MockRoleRepository)
	service := NewUserService(mockUserRepo, mockRoleRepo)

	user := &domain.User{
		ID:    "user-1",
		Email: "test@example.com",
	}

	mockUserRepo.On("GetByID", mock.Anything, "user-1").Return(user, nil)
	mockUserRepo.On("UpdateStatus", mock.Anything, "user-1", true, (*string)(nil), 0).Return(nil)

	err := service.UpdateUserStatus(context.Background(), "user-1", domain.UserStatusActive)

	require.NoError(t, err)

	mockUserRepo.AssertExpectations(t)
}

func TestUserService_UpdateUserStatus_Suspend(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	mockRoleRepo := new(MockRoleRepository)
	service := NewUserService(mockUserRepo, mockRoleRepo)

	user := &domain.User{
		ID:    "user-1",
		Email: "test@example.com",
	}

	mockUserRepo.On("GetByID", mock.Anything, "user-1").Return(user, nil)
	mockUserRepo.On("UpdateStatus", mock.Anything, "user-1", false, mock.AnythingOfType("*string"), 0).Return(nil)

	err := service.UpdateUserStatus(context.Background(), "user-1", domain.UserStatusSuspended)

	require.NoError(t, err)

	mockUserRepo.AssertExpectations(t)
}

func TestUserService_UpdateUserStatus_UserNotFound(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	mockRoleRepo := new(MockRoleRepository)
	service := NewUserService(mockUserRepo, mockRoleRepo)

	mockUserRepo.On("GetByID", mock.Anything, "user-1").Return(nil, errors.New("user not found"))

	err := service.UpdateUserStatus(context.Background(), "user-1", domain.UserStatusActive)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user not found")

	mockUserRepo.AssertExpectations(t)
}

func TestUserService_DeleteUser_Success(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	mockRoleRepo := new(MockRoleRepository)
	service := NewUserService(mockUserRepo, mockRoleRepo)

	mockUserRepo.On("Delete", mock.Anything, "user-1").Return(nil)

	err := service.DeleteUser(context.Background(), "user-1")

	require.NoError(t, err)

	mockUserRepo.AssertExpectations(t)
}

func TestUserService_DeleteUser_Error(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	mockRoleRepo := new(MockRoleRepository)
	service := NewUserService(mockUserRepo, mockRoleRepo)

	mockUserRepo.On("Delete", mock.Anything, "user-1").Return(errors.New("database error"))

	err := service.DeleteUser(context.Background(), "user-1")

	require.Error(t, err)

	mockUserRepo.AssertExpectations(t)
}

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

// MockSchoolRepository is a mock implementation of SchoolRepositoryInterface
type MockSchoolRepository struct {
	mock.Mock
}

func (m *MockSchoolRepository) Create(ctx context.Context, school *domain.School) error {
	args := m.Called(ctx, school)
	return args.Error(0)
}

func (m *MockSchoolRepository) GetByID(ctx context.Context, id string) (*domain.School, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.School), args.Error(1)
}

func (m *MockSchoolRepository) GetByCode(ctx context.Context, code string) (*domain.School, error) {
	args := m.Called(ctx, code)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.School), args.Error(1)
}

func (m *MockSchoolRepository) Update(ctx context.Context, school *domain.School) error {
	args := m.Called(ctx, school)
	return args.Error(0)
}

func (m *MockSchoolRepository) UpdateStatus(ctx context.Context, id string, isActive bool) error {
	args := m.Called(ctx, id, isActive)
	return args.Error(0)
}

func (m *MockSchoolRepository) List(ctx context.Context, isActive *bool, limit, offset int) ([]*domain.School, error) {
	args := m.Called(ctx, isActive, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.School), args.Error(1)
}

func (m *MockSchoolRepository) Count(ctx context.Context, isActive *bool) (int, error) {
	args := m.Called(ctx, isActive)
	return args.Int(0), args.Error(1)
}

func (m *MockSchoolRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestSchoolService_CreateSchool_Success(t *testing.T) {
	mockSchoolRepo := new(MockSchoolRepository)
	service := NewSchoolService(mockSchoolRepo)

	req := &domain.CreateSchoolRequest{
		Name:    "Test School",
		Code:    "TS001",
		Address: "123 Test St",
		Phone:   "123-456-7890",
		Email:   "test@school.com",
	}

	mockSchoolRepo.On("GetByCode", mock.Anything, "TS001").Return(nil, errors.New("not found"))
	mockSchoolRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.School")).Return(nil)

	result, err := service.CreateSchool(context.Background(), req, "creator-1")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Test School", result.Name)
	assert.Equal(t, "TS001", result.Code)
	assert.NotNil(t, result.Address)
	assert.NotNil(t, result.Phone)
	assert.NotNil(t, result.Email)
	assert.True(t, result.IsActive)

	mockSchoolRepo.AssertExpectations(t)
}

func TestSchoolService_CreateSchool_CodeExists(t *testing.T) {
	mockSchoolRepo := new(MockSchoolRepository)
	service := NewSchoolService(mockSchoolRepo)

	req := &domain.CreateSchoolRequest{
		Name:    "Test School",
		Code:    "TS001",
		Address: "123 Test St",
	}

	existingSchool := &domain.School{
		ID:   "school-1",
		Code: "TS001",
	}

	mockSchoolRepo.On("GetByCode", mock.Anything, "TS001").Return(existingSchool, nil)

	result, err := service.CreateSchool(context.Background(), req, "creator-1")

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "already exists")

	mockSchoolRepo.AssertExpectations(t)
}

func TestSchoolService_CreateSchool_CreateError(t *testing.T) {
	mockSchoolRepo := new(MockSchoolRepository)
	service := NewSchoolService(mockSchoolRepo)

	req := &domain.CreateSchoolRequest{
		Name: "Test School",
		Code: "TS001",
	}

	mockSchoolRepo.On("GetByCode", mock.Anything, "TS001").Return(nil, errors.New("not found"))
	mockSchoolRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.School")).Return(errors.New("database error"))

	result, err := service.CreateSchool(context.Background(), req, "creator-1")

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to create school")

	mockSchoolRepo.AssertExpectations(t)
}

func TestSchoolService_GetSchool_Success(t *testing.T) {
	mockSchoolRepo := new(MockSchoolRepository)
	service := NewSchoolService(mockSchoolRepo)

	school := &domain.School{
		ID:   "school-1",
		Name: "Test School",
	}

	mockSchoolRepo.On("GetByID", mock.Anything, "school-1").Return(school, nil)

	result, err := service.GetSchool(context.Background(), "school-1")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, school.ID, result.ID)

	mockSchoolRepo.AssertExpectations(t)
}

func TestSchoolService_GetSchool_NotFound(t *testing.T) {
	mockSchoolRepo := new(MockSchoolRepository)
	service := NewSchoolService(mockSchoolRepo)

	mockSchoolRepo.On("GetByID", mock.Anything, "school-1").Return(nil, errors.New("not found"))

	result, err := service.GetSchool(context.Background(), "school-1")

	require.Error(t, err)
	assert.Nil(t, result)

	mockSchoolRepo.AssertExpectations(t)
}

func TestSchoolService_GetSchoolByCode_Success(t *testing.T) {
	mockSchoolRepo := new(MockSchoolRepository)
	service := NewSchoolService(mockSchoolRepo)

	school := &domain.School{
		ID:   "school-1",
		Code: "TS001",
	}

	mockSchoolRepo.On("GetByCode", mock.Anything, "TS001").Return(school, nil)

	result, err := service.GetSchoolByCode(context.Background(), "TS001")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, school.ID, result.ID)

	mockSchoolRepo.AssertExpectations(t)
}

func TestSchoolService_ListSchools_Success(t *testing.T) {
	mockSchoolRepo := new(MockSchoolRepository)
	service := NewSchoolService(mockSchoolRepo)

	schools := []*domain.School{
		{ID: "school-1", Name: "School 1"},
		{ID: "school-2", Name: "School 2"},
	}

	mockSchoolRepo.On("List", mock.Anything, (*bool)(nil), 10, 0).Return(schools, nil)
	mockSchoolRepo.On("Count", mock.Anything, (*bool)(nil)).Return(2, nil)

	result, total, err := service.ListSchools(context.Background(), nil, 1, 10)

	require.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, 2, total)

	mockSchoolRepo.AssertExpectations(t)
}

func TestSchoolService_ListSchools_WithFilter(t *testing.T) {
	mockSchoolRepo := new(MockSchoolRepository)
	service := NewSchoolService(mockSchoolRepo)

	isActive := true

	schools := []*domain.School{
		{ID: "school-1", Name: "School 1", IsActive: true},
	}

	mockSchoolRepo.On("List", mock.Anything, &isActive, 10, 0).Return(schools, nil)
	mockSchoolRepo.On("Count", mock.Anything, &isActive).Return(1, nil)

	result, total, err := service.ListSchools(context.Background(), &isActive, 1, 10)

	require.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, 1, total)

	mockSchoolRepo.AssertExpectations(t)
}

func TestSchoolService_ListSchools_Error(t *testing.T) {
	mockSchoolRepo := new(MockSchoolRepository)
	service := NewSchoolService(mockSchoolRepo)

	mockSchoolRepo.On("List", mock.Anything, (*bool)(nil), 10, 0).Return(nil, errors.New("database error"))

	result, total, err := service.ListSchools(context.Background(), nil, 1, 10)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, 0, total)

	mockSchoolRepo.AssertExpectations(t)
}

func TestSchoolService_UpdateSchool_Success(t *testing.T) {
	mockSchoolRepo := new(MockSchoolRepository)
	service := NewSchoolService(mockSchoolRepo)

	school := &domain.School{
		ID:   "school-1",
		Name: "Old Name",
	}

	newName := "New Name"

	req := &domain.UpdateSchoolRequest{
		Name: &newName,
	}

	mockSchoolRepo.On("GetByID", mock.Anything, "school-1").Return(school, nil)
	mockSchoolRepo.On("Update", mock.Anything, mock.AnythingOfType("*domain.School")).Return(nil)

	result, err := service.UpdateSchool(context.Background(), "school-1", req, "updater-1")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, newName, result.Name)

	mockSchoolRepo.AssertExpectations(t)
}

func TestSchoolService_UpdateSchool_NotFound(t *testing.T) {
	mockSchoolRepo := new(MockSchoolRepository)
	service := NewSchoolService(mockSchoolRepo)

	req := &domain.UpdateSchoolRequest{}

	mockSchoolRepo.On("GetByID", mock.Anything, "school-1").Return(nil, errors.New("not found"))

	result, err := service.UpdateSchool(context.Background(), "school-1", req, "updater-1")

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "school not found")

	mockSchoolRepo.AssertExpectations(t)
}

func TestSchoolService_UpdateSchoolStatus_Success(t *testing.T) {
	mockSchoolRepo := new(MockSchoolRepository)
	service := NewSchoolService(mockSchoolRepo)

	mockSchoolRepo.On("UpdateStatus", mock.Anything, "school-1", true).Return(nil)

	err := service.UpdateSchoolStatus(context.Background(), "school-1", domain.SchoolStatusActive)

	require.NoError(t, err)

	mockSchoolRepo.AssertExpectations(t)
}

func TestSchoolService_UpdateSchoolStatus_Inactive(t *testing.T) {
	mockSchoolRepo := new(MockSchoolRepository)
	service := NewSchoolService(mockSchoolRepo)

	mockSchoolRepo.On("UpdateStatus", mock.Anything, "school-1", false).Return(nil)

	err := service.UpdateSchoolStatus(context.Background(), "school-1", domain.SchoolStatusInactive)

	require.NoError(t, err)

	mockSchoolRepo.AssertExpectations(t)
}

func TestSchoolService_DeleteSchool_Success(t *testing.T) {
	mockSchoolRepo := new(MockSchoolRepository)
	service := NewSchoolService(mockSchoolRepo)

	mockSchoolRepo.On("Delete", mock.Anything, "school-1").Return(nil)

	err := service.DeleteSchool(context.Background(), "school-1")

	require.NoError(t, err)

	mockSchoolRepo.AssertExpectations(t)
}

func TestSchoolService_DeleteSchool_Error(t *testing.T) {
	mockSchoolRepo := new(MockSchoolRepository)
	service := NewSchoolService(mockSchoolRepo)

	mockSchoolRepo.On("Delete", mock.Anything, "school-1").Return(errors.New("database error"))

	err := service.DeleteSchool(context.Background(), "school-1")

	require.Error(t, err)

	mockSchoolRepo.AssertExpectations(t)
}

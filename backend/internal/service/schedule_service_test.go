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

// MockScheduleRepository is a mock implementation of ScheduleRepositoryInterface
type MockScheduleRepository struct {
	mock.Mock
}

func (m *MockScheduleRepository) Create(ctx context.Context, schedule *domain.Schedule) error {
	args := m.Called(ctx, schedule)
	return args.Error(0)
}

func (m *MockScheduleRepository) GetByID(ctx context.Context, id string) (*domain.Schedule, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Schedule), args.Error(1)
}

func (m *MockScheduleRepository) List(ctx context.Context, classID *string, dayOfWeek *int, isActive *bool, limit, offset int) ([]*domain.Schedule, error) {
	args := m.Called(ctx, classID, dayOfWeek, isActive, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Schedule), args.Error(1)
}

func (m *MockScheduleRepository) Count(ctx context.Context, classID *string, dayOfWeek *int, isActive *bool) (int, error) {
	args := m.Called(ctx, classID, dayOfWeek, isActive)
	return args.Int(0), args.Error(1)
}

func (m *MockScheduleRepository) Update(ctx context.Context, schedule *domain.Schedule) error {
	args := m.Called(ctx, schedule)
	return args.Error(0)
}

func (m *MockScheduleRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockScheduleRepository) GetByClassID(ctx context.Context, classID string) ([]*domain.Schedule, error) {
	args := m.Called(ctx, classID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Schedule), args.Error(1)
}

func TestScheduleService_CreateSchedule_Success(t *testing.T) {
	mockScheduleRepo := new(MockScheduleRepository)
	mockClassRepo := new(MockClassRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewScheduleService(mockScheduleRepo, mockClassRepo, mockUserRepo)

	room := "Room 101"
	req := &domain.CreateScheduleRequest{
		ClassID:   "class-1",
		DayOfWeek: 1,
		StartTime: "08:00",
		EndTime:   "09:00",
		Room:      &room,
	}

	class := &domain.Class{
		ID:       "class-1",
		Name:     "Test Class",
		IsActive: true,
	}

	mockClassRepo.On("GetByID", mock.Anything, "class-1").Return(class, nil)
	mockScheduleRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.Schedule")).Return(nil)

	result, err := service.CreateSchedule(context.Background(), req, "creator-1")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "class-1", result.ClassID)

	mockScheduleRepo.AssertExpectations(t)
	mockClassRepo.AssertExpectations(t)
}

func TestScheduleService_CreateSchedule_ClassNotFound(t *testing.T) {
	mockScheduleRepo := new(MockScheduleRepository)
	mockClassRepo := new(MockClassRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewScheduleService(mockScheduleRepo, mockClassRepo, mockUserRepo)

	req := &domain.CreateScheduleRequest{
		ClassID:   "class-1",
		DayOfWeek: 1,
		StartTime: "08:00",
		EndTime:   "09:00",
	}

	mockClassRepo.On("GetByID", mock.Anything, "class-1").Return(nil, errors.New("not found"))

	result, err := service.CreateSchedule(context.Background(), req, "creator-1")

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "class not found")

	mockClassRepo.AssertExpectations(t)
}

func TestScheduleService_GetSchedule_Success(t *testing.T) {
	mockScheduleRepo := new(MockScheduleRepository)
	mockClassRepo := new(MockClassRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewScheduleService(mockScheduleRepo, mockClassRepo, mockUserRepo)

	schedule := &domain.Schedule{
		ID:      "schedule-1",
		ClassID: "class-1",
	}

	mockScheduleRepo.On("GetByID", mock.Anything, "schedule-1").Return(schedule, nil)

	result, err := service.GetSchedule(context.Background(), "schedule-1")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, schedule.ID, result.ID)

	mockScheduleRepo.AssertExpectations(t)
}

func TestScheduleService_ListSchedules_Success(t *testing.T) {
	mockScheduleRepo := new(MockScheduleRepository)
	mockClassRepo := new(MockClassRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewScheduleService(mockScheduleRepo, mockClassRepo, mockUserRepo)

	schedules := []*domain.Schedule{
		{ID: "schedule-1", ClassID: "class-1"},
		{ID: "schedule-2", ClassID: "class-2"},
	}

	mockScheduleRepo.On("List", mock.Anything, (*string)(nil), (*int)(nil), (*bool)(nil), 10, 0).Return(schedules, nil)
	mockScheduleRepo.On("Count", mock.Anything, (*string)(nil), (*int)(nil), (*bool)(nil)).Return(2, nil)

	result, total, err := service.ListSchedules(context.Background(), nil, nil, nil, 1, 10)

	require.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, 2, total)

	mockScheduleRepo.AssertExpectations(t)
}

func TestScheduleService_UpdateSchedule_Success(t *testing.T) {
	mockScheduleRepo := new(MockScheduleRepository)
	mockClassRepo := new(MockClassRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewScheduleService(mockScheduleRepo, mockClassRepo, mockUserRepo)

	schedule := &domain.Schedule{
		ID:        "schedule-1",
		ClassID:   "class-1",
		DayOfWeek: 1,
		StartTime: "08:00",
	}

	newDayOfWeek := 2

	req := &domain.UpdateScheduleRequest{
		DayOfWeek: &newDayOfWeek,
	}

	mockScheduleRepo.On("GetByID", mock.Anything, "schedule-1").Return(schedule, nil)
	mockScheduleRepo.On("Update", mock.Anything, mock.AnythingOfType("*domain.Schedule")).Return(nil)

	result, err := service.UpdateSchedule(context.Background(), "schedule-1", req, "updater-1")

	require.NoError(t, err)
	assert.NotNil(t, result)

	mockScheduleRepo.AssertExpectations(t)
}

func TestScheduleService_DeleteSchedule_Success(t *testing.T) {
	mockScheduleRepo := new(MockScheduleRepository)
	mockClassRepo := new(MockClassRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewScheduleService(mockScheduleRepo, mockClassRepo, mockUserRepo)

	mockScheduleRepo.On("Delete", mock.Anything, "schedule-1").Return(nil)

	err := service.DeleteSchedule(context.Background(), "schedule-1")

	require.NoError(t, err)

	mockScheduleRepo.AssertExpectations(t)
}

func TestScheduleService_GetClassSchedules_Success(t *testing.T) {
	mockScheduleRepo := new(MockScheduleRepository)
	mockClassRepo := new(MockClassRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewScheduleService(mockScheduleRepo, mockClassRepo, mockUserRepo)

	class := &domain.Class{
		ID:   "class-1",
		Name: "Test Class",
	}

	schedules := []*domain.Schedule{
		{ID: "schedule-1", ClassID: "class-1"},
	}

	mockClassRepo.On("GetByID", mock.Anything, "class-1").Return(class, nil)
	mockScheduleRepo.On("GetByClassID", mock.Anything, "class-1").Return(schedules, nil)

	result, err := service.GetClassSchedules(context.Background(), "class-1")

	require.NoError(t, err)
	assert.Len(t, result, 1)

	mockScheduleRepo.AssertExpectations(t)
	mockClassRepo.AssertExpectations(t)
}

func TestScheduleService_GetClassSchedules_ClassNotFound(t *testing.T) {
	mockScheduleRepo := new(MockScheduleRepository)
	mockClassRepo := new(MockClassRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewScheduleService(mockScheduleRepo, mockClassRepo, mockUserRepo)

	mockClassRepo.On("GetByID", mock.Anything, "class-1").Return(nil, errors.New("not found"))

	result, err := service.GetClassSchedules(context.Background(), "class-1")

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "class not found")

	mockClassRepo.AssertExpectations(t)
}

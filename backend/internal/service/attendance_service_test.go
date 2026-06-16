package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/nusa/backend/internal/domain"
)

// MockAttendanceRepository is a mock implementation of AttendanceRepositoryInterface
type MockAttendanceRepository struct {
	mock.Mock
}

func (m *MockAttendanceRepository) Create(ctx context.Context, attendance *domain.AttendanceRecord) error {
	args := m.Called(ctx, attendance)
	return args.Error(0)
}

func (m *MockAttendanceRepository) GetByID(ctx context.Context, id string) (*domain.AttendanceRecord, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.AttendanceRecord), args.Error(1)
}

func (m *MockAttendanceRepository) List(ctx context.Context, classID, studentID *string, status *string, startDate, endDate *time.Time, limit, offset int) ([]*domain.AttendanceRecord, error) {
	args := m.Called(ctx, classID, studentID, status, startDate, endDate, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.AttendanceRecord), args.Error(1)
}

func (m *MockAttendanceRepository) Count(ctx context.Context, classID, studentID *string, status *string, startDate, endDate *time.Time) (int, error) {
	args := m.Called(ctx, classID, studentID, status, startDate, endDate)
	return args.Int(0), args.Error(1)
}

func (m *MockAttendanceRepository) Update(ctx context.Context, attendance *domain.AttendanceRecord) error {
	args := m.Called(ctx, attendance)
	return args.Error(0)
}

func (m *MockAttendanceRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockAttendanceRepository) GetAttendanceStats(ctx context.Context, classID string, startDate, endDate time.Time) (map[string]int, error) {
	args := m.Called(ctx, classID, startDate, endDate)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(map[string]int), args.Error(1)
}

func (m *MockAttendanceRepository) GetStudentAttendanceStats(ctx context.Context, studentID string, startDate, endDate time.Time) (map[string]int, error) {
	args := m.Called(ctx, studentID, startDate, endDate)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(map[string]int), args.Error(1)
}

// MockClassRepository is a mock implementation of ClassRepositoryInterface
type MockClassRepository struct {
	mock.Mock
}

func (m *MockClassRepository) Create(ctx context.Context, class *domain.Class) error {
	args := m.Called(ctx, class)
	return args.Error(0)
}

func (m *MockClassRepository) GetByID(ctx context.Context, id string) (*domain.Class, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Class), args.Error(1)
}

func (m *MockClassRepository) List(ctx context.Context, schoolID, academicYearID, semesterID, subjectID, teacherID *string, isActive *bool, limit, offset int) ([]*domain.Class, error) {
	args := m.Called(ctx, schoolID, academicYearID, semesterID, subjectID, teacherID, isActive, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Class), args.Error(1)
}

func (m *MockClassRepository) Count(ctx context.Context, schoolID, academicYearID, semesterID, subjectID, teacherID *string, isActive *bool) (int, error) {
	args := m.Called(ctx, schoolID, academicYearID, semesterID, subjectID, teacherID, isActive)
	return args.Int(0), args.Error(1)
}

func (m *MockClassRepository) Update(ctx context.Context, class *domain.Class) error {
	args := m.Called(ctx, class)
	return args.Error(0)
}

func (m *MockClassRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockClassRepository) GetStudentCount(ctx context.Context, classID string) (int, error) {
	args := m.Called(ctx, classID)
	return args.Int(0), args.Error(1)
}

func TestAttendanceService_RecordAttendance_Success(t *testing.T) {
	mockAttendanceRepo := new(MockAttendanceRepository)
	mockClassRepo := new(MockClassRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewAttendanceService(mockAttendanceRepo, mockClassRepo, mockUserRepo)

	req := &domain.CreateAttendanceRequest{
		ClassID:   "class-1",
		StudentID: "student-1",
		Date:      time.Now(),
		Status:    domain.AttendanceStatusPresent,
	}

	class := &domain.Class{
		ID:       "class-1",
		Name:     "Test Class",
		IsActive: true,
	}

	student := &domain.User{
		ID:       "student-1",
		Name:     "Student",
		IsActive: true,
	}

	recorder := &domain.User{
		ID:       "recorder-1",
		Name:     "Recorder",
		IsActive: true,
	}

	mockClassRepo.On("GetByID", mock.Anything, "class-1").Return(class, nil)
	mockUserRepo.On("GetByID", mock.Anything, "student-1").Return(student, nil)
	mockUserRepo.On("GetByID", mock.Anything, "recorder-1").Return(recorder, nil)
	mockAttendanceRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.AttendanceRecord")).Return(nil)

	result, err := service.RecordAttendance(context.Background(), req, "recorder-1")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "class-1", result.ClassID)
	assert.Equal(t, "student-1", result.StudentID)

	mockAttendanceRepo.AssertExpectations(t)
	mockClassRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestAttendanceService_RecordAttendance_ClassNotFound(t *testing.T) {
	mockAttendanceRepo := new(MockAttendanceRepository)
	mockClassRepo := new(MockClassRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewAttendanceService(mockAttendanceRepo, mockClassRepo, mockUserRepo)

	req := &domain.CreateAttendanceRequest{
		ClassID:   "class-1",
		StudentID: "student-1",
		Date:      time.Now(),
	}

	mockClassRepo.On("GetByID", mock.Anything, "class-1").Return(nil, errors.New("not found"))

	result, err := service.RecordAttendance(context.Background(), req, "recorder-1")

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "class not found")

	mockClassRepo.AssertExpectations(t)
}

func TestAttendanceService_GetAttendance_Success(t *testing.T) {
	mockAttendanceRepo := new(MockAttendanceRepository)
	mockClassRepo := new(MockClassRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewAttendanceService(mockAttendanceRepo, mockClassRepo, mockUserRepo)

	attendance := &domain.AttendanceRecord{
		ID:     "attendance-1",
		Status: string(domain.AttendanceStatusPresent),
	}

	mockAttendanceRepo.On("GetByID", mock.Anything, "attendance-1").Return(attendance, nil)

	result, err := service.GetAttendance(context.Background(), "attendance-1")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, attendance.ID, result.ID)

	mockAttendanceRepo.AssertExpectations(t)
}

func TestAttendanceService_ListAttendances_Success(t *testing.T) {
	mockAttendanceRepo := new(MockAttendanceRepository)
	mockClassRepo := new(MockClassRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewAttendanceService(mockAttendanceRepo, mockClassRepo, mockUserRepo)

	attendances := []*domain.AttendanceRecord{
		{ID: "attendance-1", Status: string(domain.AttendanceStatusPresent)},
		{ID: "attendance-2", Status: string(domain.AttendanceStatusAbsent)},
	}

	mockAttendanceRepo.On("List", mock.Anything, (*string)(nil), (*string)(nil), (*string)(nil), (*time.Time)(nil), (*time.Time)(nil), 10, 0).Return(attendances, nil)
	mockAttendanceRepo.On("Count", mock.Anything, (*string)(nil), (*string)(nil), (*string)(nil), (*time.Time)(nil), (*time.Time)(nil)).Return(2, nil)

	result, total, err := service.ListAttendances(context.Background(), nil, nil, nil, nil, nil, 1, 10)

	require.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, 2, total)

	mockAttendanceRepo.AssertExpectations(t)
}

func TestAttendanceService_UpdateAttendance_Success(t *testing.T) {
	mockAttendanceRepo := new(MockAttendanceRepository)
	mockClassRepo := new(MockClassRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewAttendanceService(mockAttendanceRepo, mockClassRepo, mockUserRepo)

	attendance := &domain.AttendanceRecord{
		ID:     "attendance-1",
		Status: string(domain.AttendanceStatusPresent),
	}

	req := &domain.UpdateAttendanceRequest{
		Status: domain.AttendanceStatusAbsent,
	}

	recorder := &domain.User{
		ID:       "recorder-1",
		Name:     "Recorder",
		IsActive: true,
	}

	mockAttendanceRepo.On("GetByID", mock.Anything, "attendance-1").Return(attendance, nil)
	mockUserRepo.On("GetByID", mock.Anything, "recorder-1").Return(recorder, nil)
	mockAttendanceRepo.On("Update", mock.Anything, mock.AnythingOfType("*domain.AttendanceRecord")).Return(nil)

	result, err := service.UpdateAttendance(context.Background(), "attendance-1", req, "recorder-1")

	require.NoError(t, err)
	assert.NotNil(t, result)

	mockAttendanceRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestAttendanceService_DeleteAttendance_Success(t *testing.T) {
	mockAttendanceRepo := new(MockAttendanceRepository)
	mockClassRepo := new(MockClassRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewAttendanceService(mockAttendanceRepo, mockClassRepo, mockUserRepo)

	mockAttendanceRepo.On("Delete", mock.Anything, "attendance-1").Return(nil)

	err := service.DeleteAttendance(context.Background(), "attendance-1")

	require.NoError(t, err)

	mockAttendanceRepo.AssertExpectations(t)
}

func TestAttendanceService_GetClassAttendanceStats_Success(t *testing.T) {
	mockAttendanceRepo := new(MockAttendanceRepository)
	mockClassRepo := new(MockClassRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewAttendanceService(mockAttendanceRepo, mockClassRepo, mockUserRepo)

	class := &domain.Class{
		ID:   "class-1",
		Name: "Test Class",
	}

	stats := map[string]int{
		"present": 20,
		"absent":  2,
		"late":    3,
	}

	mockClassRepo.On("GetByID", mock.Anything, "class-1").Return(class, nil)
	mockAttendanceRepo.On("GetAttendanceStats", mock.Anything, "class-1", mock.Anything, mock.Anything).Return(stats, nil)

	result, err := service.GetClassAttendanceStats(context.Background(), "class-1", time.Now(), time.Now())

	require.NoError(t, err)
	assert.Equal(t, 20, result["present"])

	mockAttendanceRepo.AssertExpectations(t)
	mockClassRepo.AssertExpectations(t)
}

func TestAttendanceService_GetStudentAttendanceStats_Success(t *testing.T) {
	mockAttendanceRepo := new(MockAttendanceRepository)
	mockClassRepo := new(MockClassRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewAttendanceService(mockAttendanceRepo, mockClassRepo, mockUserRepo)

	student := &domain.User{
		ID:   "student-1",
		Name: "Student",
	}

	stats := map[string]int{
		"present": 18,
		"absent":  1,
		"late":    2,
	}

	mockUserRepo.On("GetByID", mock.Anything, "student-1").Return(student, nil)
	mockAttendanceRepo.On("GetStudentAttendanceStats", mock.Anything, "student-1", mock.Anything, mock.Anything).Return(stats, nil)

	result, err := service.GetStudentAttendanceStats(context.Background(), "student-1", time.Now(), time.Now())

	require.NoError(t, err)
	assert.Equal(t, 18, result["present"])

	mockAttendanceRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

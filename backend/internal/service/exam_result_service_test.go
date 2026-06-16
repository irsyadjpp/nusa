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

// MockExamResultRepository is a mock implementation of ExamResultRepositoryInterface
type MockExamResultRepository struct {
	mock.Mock
}

func (m *MockExamResultRepository) Create(ctx context.Context, result *domain.ExamResult) error {
	args := m.Called(ctx, result)
	return args.Error(0)
}

func (m *MockExamResultRepository) GetByID(ctx context.Context, id string) (*domain.ExamResult, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.ExamResult), args.Error(1)
}

func (m *MockExamResultRepository) GetByExamAndStudent(ctx context.Context, examID, studentID string) (*domain.ExamResult, error) {
	args := m.Called(ctx, examID, studentID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.ExamResult), args.Error(1)
}

func (m *MockExamResultRepository) List(ctx context.Context, examID, studentID, grade *string, limit, offset int) ([]*domain.ExamResult, error) {
	args := m.Called(ctx, examID, studentID, grade, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.ExamResult), args.Error(1)
}

func (m *MockExamResultRepository) Count(ctx context.Context, examID, studentID, grade *string) (int, error) {
	args := m.Called(ctx, examID, studentID, grade)
	return args.Int(0), args.Error(1)
}

func (m *MockExamResultRepository) Update(ctx context.Context, result *domain.ExamResult) error {
	args := m.Called(ctx, result)
	return args.Error(0)
}

func (m *MockExamResultRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockExamResultRepository) GetByExamID(ctx context.Context, examID string) ([]*domain.ExamResult, error) {
	args := m.Called(ctx, examID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.ExamResult), args.Error(1)
}

func (m *MockExamResultRepository) GetByStudentID(ctx context.Context, studentID string) ([]*domain.ExamResult, error) {
	args := m.Called(ctx, studentID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.ExamResult), args.Error(1)
}

func TestExamResultService_CreateExamResult_Success(t *testing.T) {
	mockExamResultRepo := new(MockExamResultRepository)
	mockExamRepo := new(MockExamRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewExamResultService(mockExamResultRepo, mockExamRepo, mockUserRepo)

	score := 85.0
	grade := "A"

	req := &domain.CreateExamResultRequest{
		ExamID:    "exam-1",
		StudentID: "student-1",
		Score:     &score,
		Grade:     &grade,
	}

	exam := &domain.Exam{
		ID:      "exam-1",
		ClassID: "class-1",
	}

	student := &domain.User{
		ID:       "student-1",
		Name:     "Student",
		IsActive: true,
	}

	grader := &domain.User{
		ID:       "grader-1",
		Name:     "Grader",
		IsActive: true,
	}

	mockExamRepo.On("GetByID", mock.Anything, "exam-1").Return(exam, nil)
	mockUserRepo.On("GetByID", mock.Anything, "student-1").Return(student, nil)
	mockUserRepo.On("GetByID", mock.Anything, "grader-1").Return(grader, nil)
	mockExamResultRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.ExamResult")).Return(nil)

	result, err := service.CreateExamResult(context.Background(), req, "grader-1")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "exam-1", result.ExamID)
	assert.Equal(t, "student-1", result.StudentID)

	mockExamResultRepo.AssertExpectations(t)
	mockExamRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestExamResultService_CreateExamResult_ExamNotFound(t *testing.T) {
	mockExamResultRepo := new(MockExamResultRepository)
	mockExamRepo := new(MockExamRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewExamResultService(mockExamResultRepo, mockExamRepo, mockUserRepo)

	req := &domain.CreateExamResultRequest{
		ExamID:    "exam-1",
		StudentID: "student-1",
	}

	mockExamRepo.On("GetByID", mock.Anything, "exam-1").Return(nil, errors.New("not found"))

	result, err := service.CreateExamResult(context.Background(), req, "grader-1")

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "exam not found")

	mockExamRepo.AssertExpectations(t)
}

func TestExamResultService_GetExamResult_Success(t *testing.T) {
	mockExamResultRepo := new(MockExamResultRepository)
	mockExamRepo := new(MockExamRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewExamResultService(mockExamResultRepo, mockExamRepo, mockUserRepo)

	examResult := &domain.ExamResult{
		ID:        "result-1",
		ExamID:    "exam-1",
		StudentID: "student-1",
	}

	mockExamResultRepo.On("GetByID", mock.Anything, "result-1").Return(examResult, nil)

	result, err := service.GetExamResult(context.Background(), "result-1")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, examResult.ID, result.ID)

	mockExamResultRepo.AssertExpectations(t)
}

func TestExamResultService_GetExamResultByExamAndStudent_Success(t *testing.T) {
	mockExamResultRepo := new(MockExamResultRepository)
	mockExamRepo := new(MockExamRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewExamResultService(mockExamResultRepo, mockExamRepo, mockUserRepo)

	examResult := &domain.ExamResult{
		ID:        "result-1",
		ExamID:    "exam-1",
		StudentID: "student-1",
	}

	mockExamResultRepo.On("GetByExamAndStudent", mock.Anything, "exam-1", "student-1").Return(examResult, nil)

	result, err := service.GetExamResultByExamAndStudent(context.Background(), "exam-1", "student-1")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, examResult.ID, result.ID)

	mockExamResultRepo.AssertExpectations(t)
}

func TestExamResultService_ListExamResults_Success(t *testing.T) {
	mockExamResultRepo := new(MockExamResultRepository)
	mockExamRepo := new(MockExamRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewExamResultService(mockExamResultRepo, mockExamRepo, mockUserRepo)

	examResults := []*domain.ExamResult{
		{ID: "result-1", ExamID: "exam-1"},
		{ID: "result-2", ExamID: "exam-2"},
	}

	mockExamResultRepo.On("List", mock.Anything, (*string)(nil), (*string)(nil), (*string)(nil), 10, 0).Return(examResults, nil)
	mockExamResultRepo.On("Count", mock.Anything, (*string)(nil), (*string)(nil), (*string)(nil)).Return(2, nil)

	result, total, err := service.ListExamResults(context.Background(), nil, nil, nil, 1, 10)

	require.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, 2, total)

	mockExamResultRepo.AssertExpectations(t)
}

func TestExamResultService_UpdateExamResult_Success(t *testing.T) {
	mockExamResultRepo := new(MockExamResultRepository)
	mockExamRepo := new(MockExamRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewExamResultService(mockExamResultRepo, mockExamRepo, mockUserRepo)

	examResult := &domain.ExamResult{
		ID:     "result-1",
		ExamID: "exam-1",
		Score:  nil,
	}

	newScore := 90.0

	req := &domain.UpdateExamResultRequest{
		Score: &newScore,
	}

	mockExamResultRepo.On("GetByID", mock.Anything, "result-1").Return(examResult, nil)
	mockExamResultRepo.On("Update", mock.Anything, mock.AnythingOfType("*domain.ExamResult")).Return(nil)

	result, err := service.UpdateExamResult(context.Background(), "result-1", req, "grader-1")

	require.NoError(t, err)
	assert.NotNil(t, result)

	mockExamResultRepo.AssertExpectations(t)
}

func TestExamResultService_DeleteExamResult_Success(t *testing.T) {
	mockExamResultRepo := new(MockExamResultRepository)
	mockExamRepo := new(MockExamRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewExamResultService(mockExamResultRepo, mockExamRepo, mockUserRepo)

	mockExamResultRepo.On("Delete", mock.Anything, "result-1").Return(nil)

	err := service.DeleteExamResult(context.Background(), "result-1")

	require.NoError(t, err)

	mockExamResultRepo.AssertExpectations(t)
}

func TestExamResultService_GetExamResultsByExam_Success(t *testing.T) {
	mockExamResultRepo := new(MockExamResultRepository)
	mockExamRepo := new(MockExamRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewExamResultService(mockExamResultRepo, mockExamRepo, mockUserRepo)

	exam := &domain.Exam{
		ID:      "exam-1",
		ClassID: "class-1",
	}

	examResults := []*domain.ExamResult{
		{ID: "result-1", ExamID: "exam-1"},
	}

	mockExamRepo.On("GetByID", mock.Anything, "exam-1").Return(exam, nil)
	mockExamResultRepo.On("GetByExamID", mock.Anything, "exam-1").Return(examResults, nil)

	result, err := service.GetExamResultsByExam(context.Background(), "exam-1")

	require.NoError(t, err)
	assert.Len(t, result, 1)

	mockExamResultRepo.AssertExpectations(t)
	mockExamRepo.AssertExpectations(t)
}

func TestExamResultService_GetExamResultsByExam_ExamNotFound(t *testing.T) {
	mockExamResultRepo := new(MockExamResultRepository)
	mockExamRepo := new(MockExamRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewExamResultService(mockExamResultRepo, mockExamRepo, mockUserRepo)

	mockExamRepo.On("GetByID", mock.Anything, "exam-1").Return(nil, errors.New("not found"))

	result, err := service.GetExamResultsByExam(context.Background(), "exam-1")

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "exam not found")

	mockExamRepo.AssertExpectations(t)
}

func TestExamResultService_GetExamResultsByStudent_Success(t *testing.T) {
	mockExamResultRepo := new(MockExamResultRepository)
	mockExamRepo := new(MockExamRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewExamResultService(mockExamResultRepo, mockExamRepo, mockUserRepo)

	student := &domain.User{
		ID:       "student-1",
		Name:     "Student",
		IsActive: true,
	}

	examResults := []*domain.ExamResult{
		{ID: "result-1", StudentID: "student-1"},
	}

	mockUserRepo.On("GetByID", mock.Anything, "student-1").Return(student, nil)
	mockExamResultRepo.On("GetByStudentID", mock.Anything, "student-1").Return(examResults, nil)

	result, err := service.GetExamResultsByStudent(context.Background(), "student-1")

	require.NoError(t, err)
	assert.Len(t, result, 1)

	mockExamResultRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestExamResultService_GetExamResultsByStudent_StudentNotFound(t *testing.T) {
	mockExamResultRepo := new(MockExamResultRepository)
	mockExamRepo := new(MockExamRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewExamResultService(mockExamResultRepo, mockExamRepo, mockUserRepo)

	mockUserRepo.On("GetByID", mock.Anything, "student-1").Return(nil, errors.New("not found"))

	result, err := service.GetExamResultsByStudent(context.Background(), "student-1")

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "student not found")

	mockUserRepo.AssertExpectations(t)
}

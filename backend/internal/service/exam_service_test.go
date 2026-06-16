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

// MockExamRepository is a mock implementation of ExamRepositoryInterface
type MockExamRepository struct {
	mock.Mock
}

func (m *MockExamRepository) Create(ctx context.Context, exam *domain.Exam) error {
	args := m.Called(ctx, exam)
	return args.Error(0)
}

func (m *MockExamRepository) GetByID(ctx context.Context, id string) (*domain.Exam, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Exam), args.Error(1)
}

func (m *MockExamRepository) List(ctx context.Context, classID, assessmentID, status *string, limit, offset int) ([]*domain.Exam, error) {
	args := m.Called(ctx, classID, assessmentID, status, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Exam), args.Error(1)
}

func (m *MockExamRepository) Count(ctx context.Context, classID, assessmentID, status *string) (int, error) {
	args := m.Called(ctx, classID, assessmentID, status)
	return args.Int(0), args.Error(1)
}

func (m *MockExamRepository) Update(ctx context.Context, exam *domain.Exam) error {
	args := m.Called(ctx, exam)
	return args.Error(0)
}

func (m *MockExamRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockExamRepository) GetByClassID(ctx context.Context, classID string) ([]*domain.Exam, error) {
	args := m.Called(ctx, classID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Exam), args.Error(1)
}

func TestExamService_CreateExam_Success(t *testing.T) {
	mockExamRepo := new(MockExamRepository)
	mockClassRepo := new(MockClassRepository)
	mockAssessmentRepo := new(MockAssessmentRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewExamService(mockExamRepo, mockClassRepo, mockAssessmentRepo, mockUserRepo)

	room := "Room 101"
	req := &domain.CreateExamRequest{
		ClassID:         "class-1",
		AssessmentID:    "assessment-1",
		ExamDate:        time.Now(),
		StartTime:       "09:00",
		DurationMinutes: 60,
		Room:            &room,
	}

	class := &domain.Class{
		ID:       "class-1",
		Name:     "Test Class",
		IsActive: true,
	}

	assessment := &domain.Assessment{
		ID:             "assessment-1",
		AssessmentType: domain.AssessmentTypeFormative,
		TPID:           "tp-1",
		TPVersionNo:    1,
		Status:         domain.WorkflowStatusDraft,
	}

	creator := &domain.User{
		ID:       "creator-1",
		Name:     "Creator",
		IsActive: true,
	}

	mockClassRepo.On("GetByID", mock.Anything, "class-1").Return(class, nil)
	mockAssessmentRepo.On("GetAssessmentByID", mock.Anything, "assessment-1").Return(assessment, nil)
	mockUserRepo.On("GetByID", mock.Anything, "creator-1").Return(creator, nil)
	mockExamRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.Exam")).Return(nil)

	result, err := service.CreateExam(context.Background(), req, "creator-1")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "class-1", result.ClassID)
	assert.Equal(t, "assessment-1", result.AssessmentID)

	mockExamRepo.AssertExpectations(t)
	mockClassRepo.AssertExpectations(t)
	mockAssessmentRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestExamService_CreateExam_ClassNotFound(t *testing.T) {
	mockExamRepo := new(MockExamRepository)
	mockClassRepo := new(MockClassRepository)
	mockAssessmentRepo := new(MockAssessmentRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewExamService(mockExamRepo, mockClassRepo, mockAssessmentRepo, mockUserRepo)

	req := &domain.CreateExamRequest{
		ClassID:      "class-1",
		AssessmentID: "assessment-1",
		ExamDate:     time.Now(),
	}

	mockClassRepo.On("GetByID", mock.Anything, "class-1").Return(nil, errors.New("not found"))

	result, err := service.CreateExam(context.Background(), req, "creator-1")

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "class not found")

	mockClassRepo.AssertExpectations(t)
}

func TestExamService_GetExam_Success(t *testing.T) {
	mockExamRepo := new(MockExamRepository)
	mockClassRepo := new(MockClassRepository)
	mockAssessmentRepo := new(MockAssessmentRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewExamService(mockExamRepo, mockClassRepo, mockAssessmentRepo, mockUserRepo)

	exam := &domain.Exam{
		ID:      "exam-1",
		ClassID: "class-1",
	}

	mockExamRepo.On("GetByID", mock.Anything, "exam-1").Return(exam, nil)

	result, err := service.GetExam(context.Background(), "exam-1")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, exam.ID, result.ID)

	mockExamRepo.AssertExpectations(t)
}

func TestExamService_ListExams_Success(t *testing.T) {
	mockExamRepo := new(MockExamRepository)
	mockClassRepo := new(MockClassRepository)
	mockAssessmentRepo := new(MockAssessmentRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewExamService(mockExamRepo, mockClassRepo, mockAssessmentRepo, mockUserRepo)

	exams := []*domain.Exam{
		{ID: "exam-1", ClassID: "class-1"},
		{ID: "exam-2", ClassID: "class-2"},
	}

	mockExamRepo.On("List", mock.Anything, (*string)(nil), (*string)(nil), (*string)(nil), 10, 0).Return(exams, nil)
	mockExamRepo.On("Count", mock.Anything, (*string)(nil), (*string)(nil), (*string)(nil)).Return(2, nil)

	result, total, err := service.ListExams(context.Background(), nil, nil, nil, 1, 10)

	require.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, 2, total)

	mockExamRepo.AssertExpectations(t)
}

func TestExamService_UpdateExam_Success(t *testing.T) {
	mockExamRepo := new(MockExamRepository)
	mockClassRepo := new(MockClassRepository)
	mockAssessmentRepo := new(MockAssessmentRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewExamService(mockExamRepo, mockClassRepo, mockAssessmentRepo, mockUserRepo)

	room := "Room 101"
	exam := &domain.Exam{
		ID:        "exam-1",
		ClassID:   "class-1",
		Room:      &room,
		StartTime: "09:00",
	}

	newRoom := "Room 102"

	req := &domain.UpdateExamRequest{
		Room: &newRoom,
	}

	mockExamRepo.On("GetByID", mock.Anything, "exam-1").Return(exam, nil)
	mockExamRepo.On("Update", mock.Anything, mock.AnythingOfType("*domain.Exam")).Return(nil)

	result, err := service.UpdateExam(context.Background(), "exam-1", req, "updater-1")

	require.NoError(t, err)
	assert.NotNil(t, result)

	mockExamRepo.AssertExpectations(t)
}

func TestExamService_DeleteExam_Success(t *testing.T) {
	mockExamRepo := new(MockExamRepository)
	mockClassRepo := new(MockClassRepository)
	mockAssessmentRepo := new(MockAssessmentRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewExamService(mockExamRepo, mockClassRepo, mockAssessmentRepo, mockUserRepo)

	mockExamRepo.On("Delete", mock.Anything, "exam-1").Return(nil)

	err := service.DeleteExam(context.Background(), "exam-1")

	require.NoError(t, err)

	mockExamRepo.AssertExpectations(t)
}

func TestExamService_GetClassExams_Success(t *testing.T) {
	mockExamRepo := new(MockExamRepository)
	mockClassRepo := new(MockClassRepository)
	mockAssessmentRepo := new(MockAssessmentRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewExamService(mockExamRepo, mockClassRepo, mockAssessmentRepo, mockUserRepo)

	class := &domain.Class{
		ID:   "class-1",
		Name: "Test Class",
	}

	exams := []*domain.Exam{
		{ID: "exam-1", ClassID: "class-1"},
	}

	mockClassRepo.On("GetByID", mock.Anything, "class-1").Return(class, nil)
	mockExamRepo.On("GetByClassID", mock.Anything, "class-1").Return(exams, nil)

	result, err := service.GetClassExams(context.Background(), "class-1")

	require.NoError(t, err)
	assert.Len(t, result, 1)

	mockExamRepo.AssertExpectations(t)
	mockClassRepo.AssertExpectations(t)
}

func TestExamService_GetClassExams_ClassNotFound(t *testing.T) {
	mockExamRepo := new(MockExamRepository)
	mockClassRepo := new(MockClassRepository)
	mockAssessmentRepo := new(MockAssessmentRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewExamService(mockExamRepo, mockClassRepo, mockAssessmentRepo, mockUserRepo)

	mockClassRepo.On("GetByID", mock.Anything, "class-1").Return(nil, errors.New("not found"))

	result, err := service.GetClassExams(context.Background(), "class-1")

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "class not found")

	mockClassRepo.AssertExpectations(t)
}

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

// MockAssignmentRepository is a mock implementation of AssignmentRepositoryInterface
type MockAssignmentRepository struct {
	mock.Mock
}

func (m *MockAssignmentRepository) Create(ctx context.Context, assignment *domain.Assignment) error {
	args := m.Called(ctx, assignment)
	return args.Error(0)
}

func (m *MockAssignmentRepository) GetByID(ctx context.Context, id string) (*domain.Assignment, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Assignment), args.Error(1)
}

func (m *MockAssignmentRepository) List(ctx context.Context, classID, assessmentID, status *string, limit, offset int) ([]*domain.Assignment, error) {
	args := m.Called(ctx, classID, assessmentID, status, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Assignment), args.Error(1)
}

func (m *MockAssignmentRepository) Count(ctx context.Context, classID, assessmentID, status *string) (int, error) {
	args := m.Called(ctx, classID, assessmentID, status)
	return args.Int(0), args.Error(1)
}

func (m *MockAssignmentRepository) Update(ctx context.Context, assignment *domain.Assignment) error {
	args := m.Called(ctx, assignment)
	return args.Error(0)
}

func (m *MockAssignmentRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockAssignmentRepository) GetByClassID(ctx context.Context, classID string) ([]*domain.Assignment, error) {
	args := m.Called(ctx, classID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Assignment), args.Error(1)
}

// MockAssessmentRepository is a mock implementation of AssessmentRepositoryInterface
type MockAssessmentRepository struct {
	mock.Mock
}

func (m *MockAssessmentRepository) CreateAssessment(ctx context.Context, assessment *domain.Assessment) error {
	args := m.Called(ctx, assessment)
	return args.Error(0)
}

func (m *MockAssessmentRepository) GetAssessmentByID(ctx context.Context, id string) (*domain.Assessment, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Assessment), args.Error(1)
}

func (m *MockAssessmentRepository) ListAssessments(ctx context.Context, tpID, userID *string, assessmentType *domain.AssessmentType, status *domain.WorkflowStatus, limit, offset int) ([]*domain.Assessment, error) {
	args := m.Called(ctx, tpID, userID, assessmentType, status, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Assessment), args.Error(1)
}

func (m *MockAssessmentRepository) UpdateAssessment(ctx context.Context, assessment *domain.Assessment) error {
	args := m.Called(ctx, assessment)
	return args.Error(0)
}

func (m *MockAssessmentRepository) DeleteAssessment(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockAssessmentRepository) CreateRubric(ctx context.Context, rubric *domain.Rubric) error {
	args := m.Called(ctx, rubric)
	return args.Error(0)
}

func (m *MockAssessmentRepository) GetRubricByID(ctx context.Context, id string) (*domain.Rubric, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Rubric), args.Error(1)
}

func (m *MockAssessmentRepository) ListRubrics(ctx context.Context, assessmentID, userID *string, rubricType *domain.RubricType, status *domain.WorkflowStatus, limit, offset int) ([]*domain.Rubric, error) {
	args := m.Called(ctx, assessmentID, userID, rubricType, status, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Rubric), args.Error(1)
}

func (m *MockAssessmentRepository) UpdateRubric(ctx context.Context, rubric *domain.Rubric) error {
	args := m.Called(ctx, rubric)
	return args.Error(0)
}

func (m *MockAssessmentRepository) DeleteRubric(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockAssessmentRepository) CreateEvidence(ctx context.Context, evidence *domain.Evidence) error {
	args := m.Called(ctx, evidence)
	return args.Error(0)
}

func (m *MockAssessmentRepository) GetEvidenceByID(ctx context.Context, id string) (*domain.Evidence, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Evidence), args.Error(1)
}

func (m *MockAssessmentRepository) ListEvidences(ctx context.Context, studentID, assessmentID *string, evidenceType *domain.EvidenceType, status *domain.EvidenceStatus, limit, offset int) ([]*domain.Evidence, error) {
	args := m.Called(ctx, studentID, assessmentID, evidenceType, status, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Evidence), args.Error(1)
}

func (m *MockAssessmentRepository) UpdateEvidence(ctx context.Context, evidence *domain.Evidence) error {
	args := m.Called(ctx, evidence)
	return args.Error(0)
}

func (m *MockAssessmentRepository) DeleteEvidence(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockAssessmentRepository) CreateEvaluation(ctx context.Context, evaluation *domain.Evaluation) error {
	args := m.Called(ctx, evaluation)
	return args.Error(0)
}

func (m *MockAssessmentRepository) GetEvaluationByID(ctx context.Context, id string) (*domain.Evaluation, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Evaluation), args.Error(1)
}

func (m *MockAssessmentRepository) ListEvaluations(ctx context.Context, evidenceID, studentID *string, performanceLevel *domain.PerformanceLevel, limit, offset int) ([]*domain.Evaluation, error) {
	args := m.Called(ctx, evidenceID, studentID, performanceLevel, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Evaluation), args.Error(1)
}

func (m *MockAssessmentRepository) UpdateEvaluation(ctx context.Context, evaluation *domain.Evaluation) error {
	args := m.Called(ctx, evaluation)
	return args.Error(0)
}

func (m *MockAssessmentRepository) DeleteEvaluation(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockAssessmentRepository) GetEvaluationHistory(ctx context.Context, evaluationID string) ([]*domain.Evaluation, error) {
	args := m.Called(ctx, evaluationID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Evaluation), args.Error(1)
}

func (m *MockAssessmentRepository) CreateFeedbackHistory(ctx context.Context, feedback *domain.EvaluationFeedbackHistory) error {
	args := m.Called(ctx, feedback)
	return args.Error(0)
}

func (m *MockAssessmentRepository) GetFeedbackHistory(ctx context.Context, evaluationID string) ([]*domain.EvaluationFeedbackHistory, error) {
	args := m.Called(ctx, evaluationID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.EvaluationFeedbackHistory), args.Error(1)
}

func (m *MockAssessmentRepository) ArchiveCurrentRevision(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockAssessmentRepository) UpdateAssessmentStatus(ctx context.Context, id string, status domain.WorkflowStatus, approvedBy *string) error {
	args := m.Called(ctx, id, status, approvedBy)
	return args.Error(0)
}

func TestAssignmentService_CreateAssignment_Success(t *testing.T) {
	mockAssignmentRepo := new(MockAssignmentRepository)
	mockClassRepo := new(MockClassRepository)
	mockAssessmentRepo := new(MockAssessmentRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewAssignmentService(mockAssignmentRepo, mockClassRepo, mockAssessmentRepo, mockUserRepo)

	req := &domain.CreateAssignmentRequest{
		ClassID:      "class-1",
		AssessmentID: "assessment-1",
		Title:        "Test Assignment",
		DueDate:      time.Now(),
		MaxScore:     100,
	}

	class := &domain.Class{
		ID:       "class-1",
		Name:     "Test Class",
		IsActive: true,
	}

	assessment := &domain.Assessment{
		ID:             "assessment-1",
		TPID:           "tp-1",
		TPVersionNo:    1,
		UserID:         "user-1",
		AssessmentType: domain.AssessmentTypeFormative,
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
	mockAssignmentRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.Assignment")).Return(nil)

	result, err := service.CreateAssignment(context.Background(), req, "creator-1")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "class-1", result.ClassID)
	assert.Equal(t, "assessment-1", result.AssessmentID)

	mockAssignmentRepo.AssertExpectations(t)
	mockClassRepo.AssertExpectations(t)
	mockAssessmentRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestAssignmentService_CreateAssignment_ClassNotFound(t *testing.T) {
	mockAssignmentRepo := new(MockAssignmentRepository)
	mockClassRepo := new(MockClassRepository)
	mockAssessmentRepo := new(MockAssessmentRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewAssignmentService(mockAssignmentRepo, mockClassRepo, mockAssessmentRepo, mockUserRepo)

	req := &domain.CreateAssignmentRequest{
		ClassID:      "class-1",
		AssessmentID: "assessment-1",
		Title:        "Test Assignment",
	}

	mockClassRepo.On("GetByID", mock.Anything, "class-1").Return(nil, errors.New("not found"))

	result, err := service.CreateAssignment(context.Background(), req, "creator-1")

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "class not found")

	mockClassRepo.AssertExpectations(t)
}

func TestAssignmentService_GetAssignment_Success(t *testing.T) {
	mockAssignmentRepo := new(MockAssignmentRepository)
	mockClassRepo := new(MockClassRepository)
	mockAssessmentRepo := new(MockAssessmentRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewAssignmentService(mockAssignmentRepo, mockClassRepo, mockAssessmentRepo, mockUserRepo)

	assignment := &domain.Assignment{
		ID:    "assignment-1",
		Title: "Test Assignment",
	}

	mockAssignmentRepo.On("GetByID", mock.Anything, "assignment-1").Return(assignment, nil)

	result, err := service.GetAssignment(context.Background(), "assignment-1")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, assignment.ID, result.ID)

	mockAssignmentRepo.AssertExpectations(t)
}

func TestAssignmentService_ListAssignments_Success(t *testing.T) {
	mockAssignmentRepo := new(MockAssignmentRepository)
	mockClassRepo := new(MockClassRepository)
	mockAssessmentRepo := new(MockAssessmentRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewAssignmentService(mockAssignmentRepo, mockClassRepo, mockAssessmentRepo, mockUserRepo)

	assignments := []*domain.Assignment{
		{ID: "assignment-1", Title: "Assignment 1"},
		{ID: "assignment-2", Title: "Assignment 2"},
	}

	mockAssignmentRepo.On("List", mock.Anything, (*string)(nil), (*string)(nil), (*string)(nil), 10, 0).Return(assignments, nil)
	mockAssignmentRepo.On("Count", mock.Anything, (*string)(nil), (*string)(nil), (*string)(nil)).Return(2, nil)

	result, total, err := service.ListAssignments(context.Background(), nil, nil, nil, 1, 10)

	require.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, 2, total)

	mockAssignmentRepo.AssertExpectations(t)
}

func TestAssignmentService_UpdateAssignment_Success(t *testing.T) {
	mockAssignmentRepo := new(MockAssignmentRepository)
	mockClassRepo := new(MockClassRepository)
	mockAssessmentRepo := new(MockAssessmentRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewAssignmentService(mockAssignmentRepo, mockClassRepo, mockAssessmentRepo, mockUserRepo)

	assignment := &domain.Assignment{
		ID:    "assignment-1",
		Title: "Old Title",
	}

	newTitle := "New Title"

	req := &domain.UpdateAssignmentRequest{
		Title: &newTitle,
	}

	mockAssignmentRepo.On("GetByID", mock.Anything, "assignment-1").Return(assignment, nil)
	mockAssignmentRepo.On("Update", mock.Anything, mock.AnythingOfType("*domain.Assignment")).Return(nil)

	result, err := service.UpdateAssignment(context.Background(), "assignment-1", req, "updater-1")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, newTitle, result.Title)

	mockAssignmentRepo.AssertExpectations(t)
}

func TestAssignmentService_DeleteAssignment_Success(t *testing.T) {
	mockAssignmentRepo := new(MockAssignmentRepository)
	mockClassRepo := new(MockClassRepository)
	mockAssessmentRepo := new(MockAssessmentRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewAssignmentService(mockAssignmentRepo, mockClassRepo, mockAssessmentRepo, mockUserRepo)

	mockAssignmentRepo.On("Delete", mock.Anything, "assignment-1").Return(nil)

	err := service.DeleteAssignment(context.Background(), "assignment-1")

	require.NoError(t, err)

	mockAssignmentRepo.AssertExpectations(t)
}

func TestAssignmentService_GetClassAssignments_Success(t *testing.T) {
	mockAssignmentRepo := new(MockAssignmentRepository)
	mockClassRepo := new(MockClassRepository)
	mockAssessmentRepo := new(MockAssessmentRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewAssignmentService(mockAssignmentRepo, mockClassRepo, mockAssessmentRepo, mockUserRepo)

	class := &domain.Class{
		ID:   "class-1",
		Name: "Test Class",
	}

	assignments := []*domain.Assignment{
		{ID: "assignment-1", Title: "Assignment 1"},
	}

	mockClassRepo.On("GetByID", mock.Anything, "class-1").Return(class, nil)
	mockAssignmentRepo.On("GetByClassID", mock.Anything, "class-1").Return(assignments, nil)

	result, err := service.GetClassAssignments(context.Background(), "class-1")

	require.NoError(t, err)
	assert.Len(t, result, 1)

	mockAssignmentRepo.AssertExpectations(t)
	mockClassRepo.AssertExpectations(t)
}

func TestAssignmentService_GetClassAssignments_ClassNotFound(t *testing.T) {
	mockAssignmentRepo := new(MockAssignmentRepository)
	mockClassRepo := new(MockClassRepository)
	mockAssessmentRepo := new(MockAssessmentRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewAssignmentService(mockAssignmentRepo, mockClassRepo, mockAssessmentRepo, mockUserRepo)

	mockClassRepo.On("GetByID", mock.Anything, "class-1").Return(nil, errors.New("not found"))

	result, err := service.GetClassAssignments(context.Background(), "class-1")

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "class not found")

	mockClassRepo.AssertExpectations(t)
}

package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/nusa/backend/internal/domain"
)

// MockTPRepository is a mock implementation of TPRepositoryInterface
type MockTPRepository struct {
	mock.Mock
}

func (m *MockTPRepository) CreateTPSet(ctx context.Context, tpSet *domain.TPSet) error {
	args := m.Called(ctx, tpSet)
	return args.Error(0)
}

func (m *MockTPRepository) GetTPSetByID(ctx context.Context, id string) (*domain.TPSet, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.TPSet), args.Error(1)
}

func (m *MockTPRepository) GetTPSetByCPAndVersion(ctx context.Context, cpID string, versionNo int) (*domain.TPSet, error) {
	args := m.Called(ctx, cpID, versionNo)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.TPSet), args.Error(1)
}

func (m *MockTPRepository) ListTPSets(ctx context.Context, cpID *string, status *domain.WorkflowStatus, schoolID *string, limit, offset int) ([]*domain.TPSet, error) {
	args := m.Called(ctx, cpID, status, schoolID, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.TPSet), args.Error(1)
}

func (m *MockTPRepository) UpdateTPSet(ctx context.Context, tpSet *domain.TPSet) error {
	args := m.Called(ctx, tpSet)
	return args.Error(0)
}

func (m *MockTPRepository) UpdateTPSetStatus(ctx context.Context, id string, status domain.WorkflowStatus, approvedBy *string, rejectedReason *string) error {
	args := m.Called(ctx, id, status, approvedBy, rejectedReason)
	return args.Error(0)
}

func (m *MockTPRepository) DeleteTPSet(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockTPRepository) CreateTP(ctx context.Context, tp *domain.TP) error {
	args := m.Called(ctx, tp)
	return args.Error(0)
}

func (m *MockTPRepository) GetTPByID(ctx context.Context, id string) (*domain.TP, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.TP), args.Error(1)
}

func (m *MockTPRepository) ListTPs(ctx context.Context, tpSetID, subjectID, phaseID, status *string, limit, offset int) ([]*domain.TP, error) {
	args := m.Called(ctx, tpSetID, subjectID, phaseID, status, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.TP), args.Error(1)
}

func (m *MockTPRepository) UpdateTP(ctx context.Context, tp *domain.TP) error {
	args := m.Called(ctx, tp)
	return args.Error(0)
}

func (m *MockTPRepository) DeleteTP(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockTPRepository) GetTPVersionHistory(ctx context.Context, tpSetID string) ([]*domain.TP, error) {
	args := m.Called(ctx, tpSetID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.TP), args.Error(1)
}

func (m *MockTPRepository) HasDownstreamAssessments(ctx context.Context, tpID string) (bool, error) {
	args := m.Called(ctx, tpID)
	return args.Bool(0), args.Error(1)
}

func TestTPService_CreateTPSet_Success(t *testing.T) {
	mockTPRepo := new(MockTPRepository)
	service := NewTPService(mockTPRepo)

	req := &domain.CreateTPSetRequest{
		CPID:             "cp-1",
		VersionNo:        1,
		GenerationSource: "AI",
	}

	mockTPRepo.On("CreateTPSet", mock.Anything, mock.AnythingOfType("*domain.TPSet")).Return(nil)

	result, err := service.CreateTPSet(context.Background(), req, "user-1")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "cp-1", result.CPID)

	mockTPRepo.AssertExpectations(t)
}

func TestTPService_GetTPSet_Success(t *testing.T) {
	mockTPRepo := new(MockTPRepository)
	service := NewTPService(mockTPRepo)

	tpSet := &domain.TPSet{
		ID:   "tpset-1",
		CPID: "cp-1",
	}

	mockTPRepo.On("GetTPSetByID", mock.Anything, "tpset-1").Return(tpSet, nil)

	result, err := service.GetTPSet(context.Background(), "tpset-1")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, tpSet.ID, result.ID)

	mockTPRepo.AssertExpectations(t)
}

func TestTPService_ListTPSets_Success(t *testing.T) {
	mockTPRepo := new(MockTPRepository)
	service := NewTPService(mockTPRepo)

	tpSets := []*domain.TPSet{
		{ID: "tpset-1", CPID: "cp-1"},
		{ID: "tpset-2", CPID: "cp-2"},
	}

	mockTPRepo.On("ListTPSets", mock.Anything, (*string)(nil), (*domain.WorkflowStatus)(nil), (*string)(nil), 10, 0).Return(tpSets, nil)

	result, total, err := service.ListTPSets(context.Background(), nil, nil, 1, 10)

	require.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, 2, total)

	mockTPRepo.AssertExpectations(t)
}

func TestTPService_ApproveTPSet_Success(t *testing.T) {
	mockTPRepo := new(MockTPRepository)
	service := NewTPService(mockTPRepo)

	mockTPRepo.On("UpdateTPSetStatus", mock.Anything, "tpset-1", domain.WorkflowStatusApproved, mock.Anything, (*string)(nil)).Return(nil)

	err := service.ApproveTPSet(context.Background(), "tpset-1", "approver-1")

	require.NoError(t, err)

	mockTPRepo.AssertExpectations(t)
}

func TestTPService_RejectTPSet_Success(t *testing.T) {
	mockTPRepo := new(MockTPRepository)
	service := NewTPService(mockTPRepo)

	mockTPRepo.On("UpdateTPSetStatus", mock.Anything, "tpset-1", domain.WorkflowStatusRejected, mock.Anything, (*string)(nil)).Return(nil)

	err := service.RejectTPSet(context.Background(), "tpset-1", "approver-1")

	require.NoError(t, err)

	mockTPRepo.AssertExpectations(t)
}

func TestTPService_UpdateTPSet_Success(t *testing.T) {
	mockTPRepo := new(MockTPRepository)
	service := NewTPService(mockTPRepo)

	reason := "Old reason"
	tpSet := &domain.TPSet{
		ID:               "tpset-1",
		CPID:             "cp-1",
		GenerationReason: &reason,
	}

	newReason := "New reason"

	req := &domain.UpdateTPSetRequest{
		GenerationReason: &newReason,
	}

	mockTPRepo.On("GetTPSetByID", mock.Anything, "tpset-1").Return(tpSet, nil)
	mockTPRepo.On("UpdateTPSet", mock.Anything, mock.AnythingOfType("*domain.TPSet")).Return(nil)

	err := service.UpdateTPSet(context.Background(), "tpset-1", req)

	require.NoError(t, err)

	mockTPRepo.AssertExpectations(t)
}

func TestTPService_GetTPVersionHistory_Success(t *testing.T) {
	mockTPRepo := new(MockTPRepository)
	service := NewTPService(mockTPRepo)

	tps := []*domain.TP{
		{ID: "tp-1", TPSetID: "tpset-1"},
		{ID: "tp-2", TPSetID: "tpset-1"},
	}

	mockTPRepo.On("GetTPVersionHistory", mock.Anything, "tpset-1").Return(tps, nil)

	result, err := service.GetTPVersionHistory(context.Background(), "tpset-1")

	require.NoError(t, err)
	assert.Len(t, result, 2)

	mockTPRepo.AssertExpectations(t)
}

func TestTPService_CreateTP_Success(t *testing.T) {
	mockTPRepo := new(MockTPRepository)
	service := NewTPService(mockTPRepo)

	title := "Test TP"
	req := &domain.CreateTPRequest{
		TPSetID:            "tpset-1",
		CPID:               "cp-1",
		SubjectID:          "subject-1",
		PhaseID:            "phase-1",
		UserID:             "user-1",
		Title:              &title,
		LearningObjectives: []string{},
		SuccessCriteria:    "criteria",
	}

	mockTPRepo.On("CreateTP", mock.Anything, mock.AnythingOfType("*domain.TP")).Return(nil)

	result, err := service.CreateTP(context.Background(), req)

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "tpset-1", result.TPSetID)

	mockTPRepo.AssertExpectations(t)
}

func TestTPService_GetTP_Success(t *testing.T) {
	mockTPRepo := new(MockTPRepository)
	service := NewTPService(mockTPRepo)

	tp := &domain.TP{
		ID:      "tp-1",
		TPSetID: "tpset-1",
	}

	mockTPRepo.On("GetTPByID", mock.Anything, "tp-1").Return(tp, nil)

	result, err := service.GetTP(context.Background(), "tp-1")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, tp.ID, result.ID)

	mockTPRepo.AssertExpectations(t)
}

func TestTPService_ListTPs_Success(t *testing.T) {
	mockTPRepo := new(MockTPRepository)
	service := NewTPService(mockTPRepo)

	tps := []*domain.TP{
		{ID: "tp-1", TPSetID: "tpset-1"},
		{ID: "tp-2", TPSetID: "tpset-2"},
	}

	mockTPRepo.On("ListTPs", mock.Anything, (*string)(nil), (*string)(nil), (*domain.WorkflowStatus)(nil), nil, 10, 0).Return(tps, nil)

	result, total, err := service.ListTPs(context.Background(), nil, nil, nil, 1, 10)

	require.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, 2, total)

	mockTPRepo.AssertExpectations(t)
}

func TestTPService_UpdateTP_Success(t *testing.T) {
	mockTPRepo := new(MockTPRepository)
	service := NewTPService(mockTPRepo)

	oldTitle := "Old Title"
	oldTP := &domain.TP{
		ID:               "tp-1",
		TPSetID:          "tpset-1",
		IsCurrentVersion: true,
		Title:            &oldTitle,
		VersionNo:        1,
	}

	newTitle := "New Title"

	req := &domain.UpdateTPRequest{
		Title: &newTitle,
	}

	mockTPRepo.On("GetTPByID", mock.Anything, "tp-1").Return(oldTP, nil)
	mockTPRepo.On("HasDownstreamAssessments", mock.Anything, "tp-1").Return(false, nil)
	mockTPRepo.On("UpdateTP", mock.Anything, mock.AnythingOfType("*domain.TP")).Return(nil)
	mockTPRepo.On("CreateTP", mock.Anything, mock.AnythingOfType("*domain.TP")).Return(nil)

	result, err := service.UpdateTP(context.Background(), "tp-1", req)

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 2, result.VersionNo)
	assert.False(t, result.IsCurrentVersion) // This should be true for the new version, but we're checking the old one

	mockTPRepo.AssertExpectations(t)
}

func TestTPService_UpdateTP_HasDownstreamAssessments(t *testing.T) {
	mockTPRepo := new(MockTPRepository)
	service := NewTPService(mockTPRepo)

	oldTP := &domain.TP{
		ID:               "tp-1",
		TPSetID:          "tpset-1",
		IsCurrentVersion: true,
	}

	req := &domain.UpdateTPRequest{
		Title: ptr("New Title"),
	}

	mockTPRepo.On("GetTPByID", mock.Anything, "tp-1").Return(oldTP, nil)
	mockTPRepo.On("HasDownstreamAssessments", mock.Anything, "tp-1").Return(true, nil)

	result, err := service.UpdateTP(context.Background(), "tp-1", req)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "cannot update TP with downstream assessments")

	mockTPRepo.AssertExpectations(t)
}

// Helper function to create string pointer
func ptr(s string) *string {
	return &s
}

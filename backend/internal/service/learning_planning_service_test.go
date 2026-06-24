package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/nusa/backend/internal/domain"
)

// MockLearningPlanningRepository is a mock implementation of LearningPlanningRepositoryInterface
type MockLearningPlanningRepository struct {
	mock.Mock
}

func (m *MockLearningPlanningRepository) CreateATPSet(ctx context.Context, atpSet *domain.ATPSet) error {
	args := m.Called(ctx, atpSet)
	return args.Error(0)
}

func (m *MockLearningPlanningRepository) GetATPSetByID(ctx context.Context, id string) (*domain.ATPSet, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.ATPSet), args.Error(1)
}

func (m *MockLearningPlanningRepository) GetATPSetByTPAndVersion(ctx context.Context, tpSetID string, versionNo int) (*domain.ATPSet, error) {
	args := m.Called(ctx, tpSetID, versionNo)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.ATPSet), args.Error(1)
}

func (m *MockLearningPlanningRepository) ListATPSets(ctx context.Context, tpSetID *string, status *domain.WorkflowStatus, limit, offset int) ([]*domain.ATPSet, error) {
	args := m.Called(ctx, tpSetID, status, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.ATPSet), args.Error(1)
}

func (m *MockLearningPlanningRepository) UpdateATPSet(ctx context.Context, atpSet *domain.ATPSet) error {
	args := m.Called(ctx, atpSet)
	return args.Error(0)
}

func (m *MockLearningPlanningRepository) DeleteATPSet(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockLearningPlanningRepository) CreateATP(ctx context.Context, atp *domain.ATP) error {
	args := m.Called(ctx, atp)
	return args.Error(0)
}

func (m *MockLearningPlanningRepository) GetATPByID(ctx context.Context, id string) (*domain.ATP, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.ATP), args.Error(1)
}

func (m *MockLearningPlanningRepository) ListATPs(ctx context.Context, atpSetID *string, limit, offset int) ([]*domain.ATP, error) {
	args := m.Called(ctx, atpSetID, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.ATP), args.Error(1)
}

func (m *MockLearningPlanningRepository) UpdateATP(ctx context.Context, atp *domain.ATP) error {
	args := m.Called(ctx, atp)
	return args.Error(0)
}

func (m *MockLearningPlanningRepository) DeleteATP(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockLearningPlanningRepository) CreateModulAjarSet(ctx context.Context, set *domain.ModulAjarSet) error {
	args := m.Called(ctx, set)
	return args.Error(0)
}

func (m *MockLearningPlanningRepository) GetModulAjarSetByID(ctx context.Context, id string) (*domain.ModulAjarSet, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.ModulAjarSet), args.Error(1)
}

func (m *MockLearningPlanningRepository) ListModulAjarSets(ctx context.Context, atpSetID *string, status *domain.WorkflowStatus, limit, offset int) ([]*domain.ModulAjarSet, error) {
	args := m.Called(ctx, atpSetID, status, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.ModulAjarSet), args.Error(1)
}

func (m *MockLearningPlanningRepository) UpdateModulAjarSet(ctx context.Context, set *domain.ModulAjarSet) error {
	args := m.Called(ctx, set)
	return args.Error(0)
}

func (m *MockLearningPlanningRepository) DeleteModulAjarSet(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockLearningPlanningRepository) CreateModulAjar(ctx context.Context, modulAjar *domain.ModulAjar) error {
	args := m.Called(ctx, modulAjar)
	return args.Error(0)
}

func (m *MockLearningPlanningRepository) GetModulAjarByID(ctx context.Context, id string) (*domain.ModulAjar, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.ModulAjar), args.Error(1)
}

func (m *MockLearningPlanningRepository) ListModulAjars(ctx context.Context, modulAjarSetID, atpID *string, limit, offset int) ([]*domain.ModulAjar, error) {
	args := m.Called(ctx, modulAjarSetID, atpID, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.ModulAjar), args.Error(1)
}

func (m *MockLearningPlanningRepository) UpdateModulAjar(ctx context.Context, modulAjar *domain.ModulAjar) error {
	args := m.Called(ctx, modulAjar)
	return args.Error(0)
}

func (m *MockLearningPlanningRepository) DeleteModulAjar(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestLearningPlanningService_CreateATPSet_Success(t *testing.T) {
	mockATPRepo := new(MockLearningPlanningRepository)
	mockModulAjarRepo := new(MockLearningPlanningRepository)
	service := NewLearningPlanningService(mockATPRepo, mockModulAjarRepo)

	req := &domain.CreateATPSetRequest{
		TPSetID:          "tpset-1",
		VersionNo:        1,
		GenerationSource: "AI",
	}

	mockATPRepo.On("CreateATPSet", mock.Anything, mock.AnythingOfType("*domain.ATPSet")).Return(nil)

	result, err := service.CreateATPSet(context.Background(), req, "user-1")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "tpset-1", result.TPSetID)

	mockATPRepo.AssertExpectations(t)
}

func TestLearningPlanningService_GetATPSet_Success(t *testing.T) {
	mockATPRepo := new(MockLearningPlanningRepository)
	mockModulAjarRepo := new(MockLearningPlanningRepository)
	service := NewLearningPlanningService(mockATPRepo, mockModulAjarRepo)

	atpSet := &domain.ATPSet{
		ID:      "atpset-1",
		TPSetID: "tpset-1",
	}

	mockATPRepo.On("GetATPSetByID", mock.Anything, "atpset-1").Return(atpSet, nil)

	result, err := service.GetATPSet(context.Background(), "atpset-1")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, atpSet.ID, result.ID)

	mockATPRepo.AssertExpectations(t)
}

func TestLearningPlanningService_ListATPSets_Success(t *testing.T) {
	mockATPRepo := new(MockLearningPlanningRepository)
	mockModulAjarRepo := new(MockLearningPlanningRepository)
	service := NewLearningPlanningService(mockATPRepo, mockModulAjarRepo)

	mockATPRepo.On("ListATPSets", mock.Anything, (*string)(nil), (*domain.WorkflowStatus)(nil), 10, 0).Return([]*domain.ATPSet{}, nil)

	result, total, err := service.ListATPSets(context.Background(), nil, nil, 1, 10)

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 0, total)

	mockATPRepo.AssertExpectations(t)
}

func TestLearningPlanningService_CreateATP_Success(t *testing.T) {
	mockATPRepo := new(MockLearningPlanningRepository)
	mockModulAjarRepo := new(MockLearningPlanningRepository)
	service := NewLearningPlanningService(mockATPRepo, mockModulAjarRepo)

	req := &domain.CreateATPRequest{
		ATPSetID: "atpset-1",
		TPID:     "tp-1",
		UserID:   "user-1",
	}

	mockATPRepo.On("CreateATP", mock.Anything, mock.AnythingOfType("*domain.ATP")).Return(nil)

	result, err := service.CreateATP(context.Background(), req)

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "atpset-1", result.ATPSetID)

	mockATPRepo.AssertExpectations(t)
}

func TestLearningPlanningService_GetATP_Success(t *testing.T) {
	mockATPRepo := new(MockLearningPlanningRepository)
	mockModulAjarRepo := new(MockLearningPlanningRepository)
	service := NewLearningPlanningService(mockATPRepo, mockModulAjarRepo)

	atp := &domain.ATP{
		ID:       "atp-1",
		ATPSetID: "atpset-1",
	}

	mockATPRepo.On("GetATPByID", mock.Anything, "atp-1").Return(atp, nil)

	result, err := service.GetATP(context.Background(), "atp-1")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, atp.ID, result.ID)

	mockATPRepo.AssertExpectations(t)
}

func TestLearningPlanningService_CreateModulAjarSet_Success(t *testing.T) {
	mockATPRepo := new(MockLearningPlanningRepository)
	mockModulAjarRepo := new(MockLearningPlanningRepository)
	service := NewLearningPlanningService(mockATPRepo, mockModulAjarRepo)

	req := &domain.CreateModulAjarSetRequest{
		ATPSetID:         "atpset-1",
		VersionNo:        1,
		GenerationSource: "AI",
	}

	mockModulAjarRepo.On("CreateModulAjarSet", mock.Anything, mock.AnythingOfType("*domain.ModulAjarSet")).Return(nil)

	result, err := service.CreateModulAjarSet(context.Background(), req, "user-1")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "atpset-1", result.ATPSetID)

	mockModulAjarRepo.AssertExpectations(t)
}

func TestLearningPlanningService_GetModulAjarSet_Success(t *testing.T) {
	mockATPRepo := new(MockLearningPlanningRepository)
	mockModulAjarRepo := new(MockLearningPlanningRepository)
	service := NewLearningPlanningService(mockATPRepo, mockModulAjarRepo)

	modulAjarSet := &domain.ModulAjarSet{
		ID:       "modulajarset-1",
		ATPSetID: "atpset-1",
	}

	mockModulAjarRepo.On("GetModulAjarSetByID", mock.Anything, "modulajarset-1").Return(modulAjarSet, nil)

	result, err := service.GetModulAjarSet(context.Background(), "modulajarset-1")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, modulAjarSet.ID, result.ID)

	mockModulAjarRepo.AssertExpectations(t)
}

func TestLearningPlanningService_CreateModulAjar_Success(t *testing.T) {
	mockATPRepo := new(MockLearningPlanningRepository)
	mockModulAjarRepo := new(MockLearningPlanningRepository)
	service := NewLearningPlanningService(mockATPRepo, mockModulAjarRepo)

	req := &domain.CreateModulAjarRequest{
		ModulAjarSetID: "modulajarset-1",
		ATPID:          "atp-1",
		Week:           1,
		Topic:          "Test Topic",
	}

	mockModulAjarRepo.On("CreateModulAjar", mock.Anything, mock.AnythingOfType("*domain.ModulAjar")).Return(nil)

	result, err := service.CreateModulAjar(context.Background(), req)

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "modulajarset-1", result.ModulAjarSetID)

	mockModulAjarRepo.AssertExpectations(t)
}

func TestLearningPlanningService_UpdateATPSet_Success(t *testing.T) {
	mockATPRepo := new(MockLearningPlanningRepository)
	mockModulAjarRepo := new(MockLearningPlanningRepository)
	service := NewLearningPlanningService(mockATPRepo, mockModulAjarRepo)

	atpSet := &domain.ATPSet{
		ID:      "atpset-1",
		TPSetID: "tpset-1",
		Status:  domain.WorkflowStatusDraft,
	}

	newStatus := domain.WorkflowStatusApproved
	req := &domain.UpdateATPSetRequest{
		Status: &newStatus,
	}

	mockATPRepo.On("GetATPSetByID", mock.Anything, "atpset-1").Return(atpSet, nil)
	mockATPRepo.On("UpdateATPSet", mock.Anything, mock.AnythingOfType("*domain.ATPSet")).Return(nil)

	result, err := service.UpdateATPSet(context.Background(), "atpset-1", req)

	require.NoError(t, err)
	assert.NotNil(t, result)

	mockATPRepo.AssertExpectations(t)
}

func TestLearningPlanningService_DeleteATPSet_Success(t *testing.T) {
	mockATPRepo := new(MockLearningPlanningRepository)
	mockModulAjarRepo := new(MockLearningPlanningRepository)
	service := NewLearningPlanningService(mockATPRepo, mockModulAjarRepo)

	mockATPRepo.On("DeleteATPSet", mock.Anything, "atpset-1").Return(nil)

	err := service.DeleteATPSet(context.Background(), "atpset-1")

	require.NoError(t, err)

	mockATPRepo.AssertExpectations(t)
}

func TestLearningPlanningService_ApproveATPSet_Success(t *testing.T) {
	mockATPRepo := new(MockLearningPlanningRepository)
	mockModulAjarRepo := new(MockLearningPlanningRepository)
	service := NewLearningPlanningService(mockATPRepo, mockModulAjarRepo)

	atpSet := &domain.ATPSet{
		ID:      "atpset-1",
		TPSetID: "tpset-1",
		Status:  domain.WorkflowStatusDraft,
	}

	mockATPRepo.On("GetATPSetByID", mock.Anything, "atpset-1").Return(atpSet, nil)
	mockATPRepo.On("UpdateATPSet", mock.Anything, mock.AnythingOfType("*domain.ATPSet")).Return(nil)

	result, err := service.ApproveATPSet(context.Background(), "atpset-1", "user-1")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, domain.WorkflowStatusApproved, result.Status)

	mockATPRepo.AssertExpectations(t)
}

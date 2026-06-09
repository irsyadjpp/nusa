package application

import (
	"context"
	"testing"

	"github.com/nusa/backend/internal/domain"
)

// Mock implementations for testing

type mockUserRepository struct {
	users map[string]*domain.User
}

func (m *mockUserRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	if user, ok := m.users[id]; ok {
		return user, nil
	}
	return nil, &domain.UserNotFoundError{}
}

type mockTPRepository struct {
	tpSets map[string]*domain.TPSet
}

func (m *mockTPRepository) CreateTPSet(ctx context.Context, tpSet *domain.TPSet) error {
	m.tpSets[tpSet.ID] = tpSet
	return nil
}

func (m *mockTPRepository) GetTPSetByID(ctx context.Context, id string) (*domain.TPSet, error) {
	if tpSet, ok := m.tpSets[id]; ok {
		return tpSet, nil
	}
	return nil, &domain.TPSetNotFoundError{}
}

func (m *mockTPRepository) UpdateTPSet(ctx context.Context, tpSet *domain.TPSet) error {
	m.tpSets[tpSet.ID] = tpSet
	return nil
}

func (m *mockTPRepository) ListTPSets(ctx context.Context, cpID *string, status *domain.WorkflowStatus, limit, offset int) ([]*domain.TPSet, error) {
	var result []*domain.TPSet
	for _, tpSet := range m.tpSets {
		result = append(result, tpSet)
	}
	return result, nil
}

type mockSchoolRepository struct{}

// Test CreateTPSet

func TestTPSetApplicationService_CreateTPSet_Success(t *testing.T) {
	// Arrange
	userRepo := &mockUserRepository{
		users: map[string]*domain.User{
			"user-1": {
				ID:       "user-1",
				SchoolID: stringPtr("school-1"),
			},
		},
	}
	tpRepo := &mockTPRepository{
		tpSets: make(map[string]*domain.TPSet),
	}
	schoolRepo := &mockSchoolRepository{}

	service := NewTPSetApplicationService(tpRepo, userRepo, schoolRepo)

	cmd := &CreateTPSetCommand{
		CPID:             "cp-123",
		VersionNo:        1,
		GenerationSource: domain.GenerationSourceManual,
		GenerationReason: stringPtr("Initial version"),
		UserID:           "user-1",
	}

	// Act
	response, err := service.CreateTPSet(context.Background(), cmd)

	// Assert
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if response == nil {
		t.Fatal("Expected response, got nil")
	}
	if response.TPSetID == "" {
		t.Error("Expected TPSetID, got empty string")
	}
	if response.Status != domain.WorkflowStatusDraft {
		t.Errorf("Expected status DRAFT, got %s", response.Status)
	}
}

func TestTPSetApplicationService_CreateTPSet_NoSchoolID(t *testing.T) {
	// Arrange
	userRepo := &mockUserRepository{
		users: map[string]*domain.User{
			"user-1": {
				ID:       "user-1",
				SchoolID: nil,
			},
		},
	}
	tpRepo := &mockTPRepository{
		tpSets: make(map[string]*domain.TPSet),
	}
	schoolRepo := &mockSchoolRepository{}

	service := NewTPSetApplicationService(tpRepo, userRepo, schoolRepo)

	cmd := &CreateTPSetCommand{
		CPID:             "cp-123",
		VersionNo:        1,
		GenerationSource: domain.GenerationSourceManual,
		UserID:           "user-1",
	}

	// Act
	_, err := service.CreateTPSet(context.Background(), cmd)

	// Assert
	if err == nil {
		t.Error("Expected error for user without school ID, got nil")
	}
}

func TestTPSetApplicationService_CreateTPSet_UserNotFound(t *testing.T) {
	// Arrange
	userRepo := &mockUserRepository{
		users: map[string]*domain.User{},
	}
	tpRepo := &mockTPRepository{
		tpSets: make(map[string]*domain.TPSet),
	}
	schoolRepo := &mockSchoolRepository{}

	service := NewTPSetApplicationService(tpRepo, userRepo, schoolRepo)

	cmd := &CreateTPSetCommand{
		CPID:             "cp-123",
		VersionNo:        1,
		GenerationSource: domain.GenerationSourceManual,
		UserID:           "user-1",
	}

	// Act
	_, err := service.CreateTPSet(context.Background(), cmd)

	// Assert
	if err == nil {
		t.Error("Expected error for user not found, got nil")
	}
}

// Test UpdateTPSet

func TestTPSetApplicationService_UpdateTPSet_Success(t *testing.T) {
	// Arrange
	userRepo := &mockUserRepository{
		users: map[string]*domain.User{
			"user-1": {
				ID:       "user-1",
				SchoolID: stringPtr("school-1"),
				RoleID:   "TEACHER",
			},
		},
	}
	tpRepo := &mockTPRepository{
		tpSets: map[string]*domain.TPSet{
			"tp-set-1": {
				ID:          "tp-set-1",
				CPID:        "cp-123",
				VersionNo:   1,
				Status:      domain.WorkflowStatusDraft,
				GeneratedBy: "user-1",
			},
		},
	}
	schoolRepo := &mockSchoolRepository{}

	service := NewTPSetApplicationService(tpRepo, userRepo, schoolRepo)

	status := domain.WorkflowStatusUnderReview
	cmd := &UpdateTPSetCommand{
		TPSetID:          "tp-set-1",
		Status:           &status,
		GenerationReason: stringPtr("Updated reason"),
		UserID:           "user-1",
	}

	// Act
	response, err := service.UpdateTPSet(context.Background(), cmd)

	// Assert
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if response == nil {
		t.Fatal("Expected response, got nil")
	}
	if response.TPSetID != "tp-set-1" {
		t.Errorf("Expected TPSetID tp-set-1, got %s", response.TPSetID)
	}
}

func TestTPSetApplicationService_UpdateTPSet_NotOwner(t *testing.T) {
	// Arrange
	userRepo := &mockUserRepository{
		users: map[string]*domain.User{
			"user-1": {
				ID:       "user-1",
				SchoolID: stringPtr("school-1"),
				RoleID:   "TEACHER",
			},
			"user-2": {
				ID:       "user-2",
				SchoolID: stringPtr("school-1"),
				RoleID:   "TEACHER",
			},
		},
	}
	tpRepo := &mockTPRepository{
		tpSets: map[string]*domain.TPSet{
			"tp-set-1": {
				ID:          "tp-set-1",
				CPID:        "cp-123",
				VersionNo:   1,
				Status:      domain.WorkflowStatusDraft,
				GeneratedBy: "user-2", // Different user
			},
		},
	}
	schoolRepo := &mockSchoolRepository{}

	service := NewTPSetApplicationService(tpRepo, userRepo, schoolRepo)

	cmd := &UpdateTPSetCommand{
		TPSetID: "tp-set-1",
		UserID:  "user-1",
	}

	// Act
	_, err := service.UpdateTPSet(context.Background(), cmd)

	// Assert
	if err == nil {
		t.Error("Expected error for non-owner update, got nil")
	}
}

// Test ApproveTPSet

func TestTPSetApplicationService_ApproveTPSet_Success(t *testing.T) {
	// Arrange
	userRepo := &mockUserRepository{
		users: map[string]*domain.User{
			"admin-1": {
				ID:       "admin-1",
				SchoolID: stringPtr("school-1"),
				RoleID:   "SCHOOL_ADMIN",
			},
			"user-1": {
				ID:       "user-1",
				SchoolID: stringPtr("school-1"),
				RoleID:   "TEACHER",
			},
		},
	}
	tpRepo := &mockTPRepository{
		tpSets: map[string]*domain.TPSet{
			"tp-set-1": {
				ID:          "tp-set-1",
				CPID:        "cp-123",
				VersionNo:   1,
				Status:      domain.WorkflowStatusUnderReview,
				GeneratedBy: "user-1",
			},
		},
	}
	schoolRepo := &mockSchoolRepository{}

	service := NewTPSetApplicationService(tpRepo, userRepo, schoolRepo)

	cmd := &ApproveTPSetCommand{
		TPSetID:    "tp-set-1",
		ApproverID: "admin-1",
		Reason:     "Approved for use",
	}

	// Act
	response, err := service.ApproveTPSet(context.Background(), cmd)

	// Assert
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if response == nil {
		t.Fatal("Expected response, got nil")
	}
	if response.Status != domain.WorkflowStatusApproved {
		t.Errorf("Expected status APPROVED, got %s", response.Status)
	}
}

func TestTPSetApplicationService_ApproveTPSet_NotAdmin(t *testing.T) {
	// Arrange
	userRepo := &mockUserRepository{
		users: map[string]*domain.User{
			"user-1": {
				ID:       "user-1",
				SchoolID: stringPtr("school-1"),
				RoleID:   "TEACHER",
			},
		},
	}
	tpRepo := &mockTPRepository{
		tpSets: map[string]*domain.TPSet{
			"tp-set-1": {
				ID:          "tp-set-1",
				CPID:        "cp-123",
				VersionNo:   1,
				Status:      domain.WorkflowStatusUnderReview,
				GeneratedBy: "user-1",
			},
		},
	}
	schoolRepo := &mockSchoolRepository{}

	service := NewTPSetApplicationService(tpRepo, userRepo, schoolRepo)

	cmd := &ApproveTPSetCommand{
		TPSetID:    "tp-set-1",
		ApproverID: "user-1",
		Reason:     "Should not be allowed",
	}

	// Act
	_, err := service.ApproveTPSet(context.Background(), cmd)

	// Assert
	if err == nil {
		t.Error("Expected error for non-admin approval, got nil")
	}
}

// Test GetTPSet

func TestTPSetApplicationService_GetTPSet_Success(t *testing.T) {
	// Arrange
	userRepo := &mockUserRepository{
		users: map[string]*domain.User{
			"user-1": {
				ID:       "user-1",
				SchoolID: stringPtr("school-1"),
				RoleID:   "TEACHER",
			},
		},
	}
	tpRepo := &mockTPRepository{
		tpSets: map[string]*domain.TPSet{
			"tp-set-1": {
				ID:          "tp-set-1",
				CPID:        "cp-123",
				VersionNo:   1,
				Status:      domain.WorkflowStatusDraft,
				GeneratedBy: "user-1",
			},
		},
	}
	schoolRepo := &mockSchoolRepository{}

	service := NewTPSetApplicationService(tpRepo, userRepo, schoolRepo)

	query := &GetTPSetQuery{
		TPSetID: "tp-set-1",
		UserID:  "user-1",
	}

	// Act
	response, err := service.GetTPSet(context.Background(), query)

	// Assert
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if response == nil {
		t.Fatal("Expected response, got nil")
	}
	if response.TPSet == nil {
		t.Error("Expected TPSet, got nil")
	}
}

func TestTPSetApplicationService_GetTPSet_CrossSchoolAccess(t *testing.T) {
	// Arrange
	userRepo := &mockUserRepository{
		users: map[string]*domain.User{
			"user-1": {
				ID:       "user-1",
				SchoolID: stringPtr("school-1"),
				RoleID:   "TEACHER",
			},
			"user-2": {
				ID:       "user-2",
				SchoolID: stringPtr("school-2"),
				RoleID:   "TEACHER",
			},
		},
	}
	tpRepo := &mockTPRepository{
		tpSets: map[string]*domain.TPSet{
			"tp-set-1": {
				ID:          "tp-set-1",
				CPID:        "cp-123",
				VersionNo:   1,
				Status:      domain.WorkflowStatusDraft,
				GeneratedBy: "user-2", // Different school
			},
		},
	}
	schoolRepo := &mockSchoolRepository{}

	service := NewTPSetApplicationService(tpRepo, userRepo, schoolRepo)

	query := &GetTPSetQuery{
		TPSetID: "tp-set-1",
		UserID:  "user-1",
	}

	// Act
	_, err := service.GetTPSet(context.Background(), query)

	// Assert
	if err == nil {
		t.Error("Expected error for cross-school access, got nil")
	}
}

// Test ListTPSets

func TestTPSetApplicationService_ListTPSets_Success(t *testing.T) {
	// Arrange
	userRepo := &mockUserRepository{
		users: map[string]*domain.User{
			"user-1": {
				ID:       "user-1",
				SchoolID: stringPtr("school-1"),
				RoleID:   "TEACHER",
			},
		},
	}
	tpRepo := &mockTPRepository{
		tpSets: map[string]*domain.TPSet{
			"tp-set-1": {
				ID:          "tp-set-1",
				CPID:        "cp-123",
				VersionNo:   1,
				Status:      domain.WorkflowStatusDraft,
				GeneratedBy: "user-1",
			},
		},
	}
	schoolRepo := &mockSchoolRepository{}

	service := NewTPSetApplicationService(tpRepo, userRepo, schoolRepo)

	query := &ListTPSetsQuery{
		UserID:  "user-1",
		Page:    1,
		PageSize: 10,
	}

	// Act
	response, err := service.ListTPSets(context.Background(), query)

	// Assert
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if response == nil {
		t.Fatal("Expected response, got nil")
	}
	if response.TPSets == nil {
		t.Error("Expected TPSets, got nil")
	}
}

// Helper functions

func stringPtr(s string) *string {
	return &s
}

// Custom error types for testing

type UserNotFoundError struct{}

func (e *UserNotFoundError) Error() string {
	return "user not found"
}

type TPSetNotFoundError struct{}

func (e *TPSetNotFoundError) Error() string {
	return "tp set not found"
}

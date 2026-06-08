package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/repository"
)

// TPService handles business logic for TP operations
type TPService struct {
	tpRepo *repository.TPRepository
}

// NewTPService creates a new TP service
func NewTPService(tpRepo *repository.TPRepository) *TPService {
	return &TPService{tpRepo: tpRepo}
}

// CreateTPSet creates a new TP Set
func (s *TPService) CreateTPSet(ctx context.Context, req *domain.CreateTPSetRequest, userID string) (*domain.TPSet, error) {
	tpSet := &domain.TPSet{
		ID:               uuid.New().String(),
		CPID:             req.CPID,
		VersionNo:        req.VersionNo,
		Status:           domain.WorkflowStatusDraft,
		GenerationSource: req.GenerationSource,
		GenerationReason: req.GenerationReason,
		GeneratedBy:      userID,
	}

	if err := s.tpRepo.CreateTPSet(ctx, tpSet); err != nil {
		return nil, fmt.Errorf("failed to create TP set: %w", err)
	}

	return tpSet, nil
}

// GetTPSet retrieves a TP Set by ID
func (s *TPService) GetTPSet(ctx context.Context, id string) (*domain.TPSet, error) {
	return s.tpRepo.GetTPSetByID(ctx, id)
}

// ListTPSets retrieves TP Sets with optional filters
func (s *TPService) ListTPSets(ctx context.Context, cpID *string, status *domain.WorkflowStatus, page, pageSize int) ([]*domain.TPSet, int, error) {
	limit := pageSize
	offset := (page - 1) * pageSize

	sets, err := s.tpRepo.ListTPSets(ctx, cpID, status, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return sets, len(sets), nil
}

// ApproveTPSet approves a TP Set
func (s *TPService) ApproveTPSet(ctx context.Context, id string, approverID string) error {
	err := s.tpRepo.UpdateTPSetStatus(ctx, id, domain.WorkflowStatusApproved, &approverID, nil)
	if err != nil {
		return fmt.Errorf("failed to approve TP set: %w", err)
	}
	return nil
}

// RejectTPSet rejects a TP Set
func (s *TPService) RejectTPSet(ctx context.Context, id string, approverID string) error {
	err := s.tpRepo.UpdateTPSetStatus(ctx, id, domain.WorkflowStatusRejected, &approverID, nil)
	if err != nil {
		return fmt.Errorf("failed to reject TP set: %w", err)
	}
	return nil
}

// CreateTP creates a new TP
func (s *TPService) CreateTP(ctx context.Context, req *domain.CreateTPRequest) (*domain.TP, error) {
	tp := &domain.TP{
		ID:                 uuid.New().String(),
		TPSetID:            req.TPSetID,
		SequenceNumber:     req.SequenceNumber,
		CPID:               req.CPID,
		SubjectID:          req.SubjectID,
		PhaseID:            req.PhaseID,
		ElementID:          req.ElementID,
		SubelementID:       req.SubelementID,
		UserID:             req.UserID,
		Status:             domain.WorkflowStatusDraft,
		Title:              req.Title,
		LearningObjectives: req.LearningObjectives,
		TimeAllocation:     req.TimeAllocation,
		Prerequisites:      req.Prerequisites,
		EstimatedWeeks:     req.EstimatedWeeks,
		SuccessCriteria:    req.SuccessCriteria,
		VersionNo:          1,
		IsCurrentVersion:   true,
		ParentVersionID:    nil,
	}

	if err := s.tpRepo.CreateTP(ctx, tp); err != nil {
		return nil, fmt.Errorf("failed to create TP: %w", err)
	}

	return tp, nil
}

// GetTP retrieves a TP by ID
func (s *TPService) GetTP(ctx context.Context, id string) (*domain.TP, error) {
	return s.tpRepo.GetTPByID(ctx, id)
}

// ListTPs retrieves TPs with optional filters
func (s *TPService) ListTPs(ctx context.Context, tpSetID, cpID *string, status *domain.WorkflowStatus, page, pageSize int) ([]*domain.TP, int, error) {
	limit := pageSize
	offset := (page - 1) * pageSize

	tps, err := s.tpRepo.ListTPs(ctx, tpSetID, cpID, status, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return tps, len(tps), nil
}

// UpdateTP updates a TP (creates new version instead of in-place update)
func (s *TPService) UpdateTP(ctx context.Context, id string, req *domain.UpdateTPRequest) (*domain.TP, error) {
	oldTP, err := s.tpRepo.GetTPByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("tp not found")
	}

	// Check if TP has downstream assessments before allowing update
	hasDownstream, err := s.tpRepo.HasDownstreamAssessments(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to check downstream assessments: %w", err)
	}
	if hasDownstream {
		return nil, fmt.Errorf("cannot update TP with downstream assessments")
	}

	// Mark old version as not current
	oldTP.IsCurrentVersion = false
	if err := s.tpRepo.UpdateTP(ctx, oldTP); err != nil {
		return nil, fmt.Errorf("failed to mark old version: %w", err)
	}

	// Prepare new version values
	title := oldTP.Title
	learningObjectives := oldTP.LearningObjectives
	timeAllocation := oldTP.TimeAllocation
	prerequisites := oldTP.Prerequisites
	estimatedWeeks := oldTP.EstimatedWeeks
	successCriteria := oldTP.SuccessCriteria
	status := oldTP.Status

	if req.Title != nil {
		title = req.Title
	}
	if req.LearningObjectives != nil {
		learningObjectives = req.LearningObjectives
	}
	if req.TimeAllocation != nil {
		timeAllocation = req.TimeAllocation
	}
	if req.Prerequisites != nil {
		prerequisites = req.Prerequisites
	}
	if req.EstimatedWeeks != nil {
		estimatedWeeks = req.EstimatedWeeks
	}
	if req.SuccessCriteria != nil {
		successCriteria = req.SuccessCriteria
	}
	if req.Status != nil {
		status = *req.Status
	}

	// Create new version
	newTP := &domain.TP{
		ID:                 uuid.New().String(),
		TPSetID:            oldTP.TPSetID,
		SequenceNumber:     oldTP.SequenceNumber,
		CPID:               oldTP.CPID,
		SubjectID:          oldTP.SubjectID,
		PhaseID:            oldTP.PhaseID,
		ElementID:          oldTP.ElementID,
		SubelementID:       oldTP.SubelementID,
		UserID:             oldTP.UserID,
		Status:             status,
		Title:              title,
		LearningObjectives: learningObjectives,
		TimeAllocation:     timeAllocation,
		Prerequisites:      prerequisites,
		EstimatedWeeks:     estimatedWeeks,
		SuccessCriteria:    successCriteria,
		VersionNo:          oldTP.VersionNo + 1,
		IsCurrentVersion:   true,
		ParentVersionID:    &oldTP.ID,
	}

	if err := s.tpRepo.CreateTP(ctx, newTP); err != nil {
		return nil, fmt.Errorf("failed to create new version: %w", err)
	}

	return newTP, nil
}

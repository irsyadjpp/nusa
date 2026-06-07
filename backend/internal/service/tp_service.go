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
		LearningObjectives:  req.LearningObjectives,
		TimeAllocation:     req.TimeAllocation,
		Prerequisites:      req.Prerequisites,
		EstimatedWeeks:     req.EstimatedWeeks,
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

// UpdateTP updates a TP
func (s *TPService) UpdateTP(ctx context.Context, id string, req *domain.UpdateTPRequest) (*domain.TP, error) {
	tp, err := s.tpRepo.GetTPByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("tp not found")
	}

	if req.Title != nil {
		tp.Title = req.Title
	}
	if req.LearningObjectives != nil {
		tp.LearningObjectives = req.LearningObjectives
	}
	if req.TimeAllocation != nil {
		tp.TimeAllocation = req.TimeAllocation
	}
	if req.Prerequisites != nil {
		tp.Prerequisites = req.Prerequisites
	}
	if req.EstimatedWeeks != nil {
		tp.EstimatedWeeks = req.EstimatedWeeks
	}
	if req.Status != nil {
		tp.Status = *req.Status
	}

	if err := s.tpRepo.UpdateTP(ctx, tp); err != nil {
		return nil, fmt.Errorf("failed to update TP: %w", err)
	}

	return tp, nil
}

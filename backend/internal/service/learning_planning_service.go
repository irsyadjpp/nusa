package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/repository"
)

// LearningPlanningService handles business logic for ATP and Modul Ajar operations
type LearningPlanningService struct {
	atpRepo         *repository.LearningPlanningRepository
	modulAjarRepo   *repository.LearningPlanningRepository
}

// NewLearningPlanningService creates a new learning planning service
func NewLearningPlanningService(atpRepo, modulAjarRepo *repository.LearningPlanningRepository) *LearningPlanningService {
	return &LearningPlanningService{
		atpRepo:       atpRepo,
		modulAjarRepo: modulAjarRepo,
	}
}

// CreateATPSet creates a new ATP Set
func (s *LearningPlanningService) CreateATPSet(ctx context.Context, req *domain.CreateATPSetRequest, userID string) (*domain.ATPSet, error) {
	atpSet := &domain.ATPSet{
		ID:               uuid.New().String(),
		TPSetID:          req.TPSetID,
		VersionNo:        req.VersionNo,
		Status:           domain.WorkflowStatusDraft,
		GenerationSource: req.GenerationSource,
		GenerationReason: req.GenerationReason,
		GeneratedBy:      userID,
	}

	if err := s.atpRepo.CreateATPSet(ctx, atpSet); err != nil {
		return nil, fmt.Errorf("failed to create ATP set: %w", err)
	}

	return atpSet, nil
}

// GetATPSet retrieves an ATP Set by ID
func (s *LearningPlanningService) GetATPSet(ctx context.Context, id string) (*domain.ATPSet, error) {
	return s.atpRepo.GetATPSetByID(ctx, id)
}

// ListATPSets retrieves ATP Sets with optional filters
func (s *LearningPlanningService) ListATPSets(ctx context.Context, tpSetID *string, status *domain.WorkflowStatus, page, pageSize int) ([]*domain.ATPSet, int, error) {
	limit := pageSize
	offset := (page - 1) * pageSize
	sets, err := s.atpRepo.ListATPSets(ctx, tpSetID, status, limit, offset)
	return sets, len(sets), err
}

// CreateATP creates a new ATP
func (s *LearningPlanningService) CreateATP(ctx context.Context, req *domain.CreateATPRequest) (*domain.ATP, error) {
	atp := &domain.ATP{
		ID:               uuid.New().String(),
		ATPSetID:         req.ATPSetID,
		TPID:             req.TPID,
		UserID:           req.UserID,
		Status:           domain.WorkflowStatusDraft,
		AcademicCalendar: req.AcademicCalendar,
		ClassSchedule:    req.ClassSchedule,
		WeeklySequence:   req.WeeklySequence,
		AssessmentSchedule: req.AssessmentSchedule,
	}
	if err := s.atpRepo.CreateATP(ctx, atp); err != nil {
		return nil, fmt.Errorf("failed to create ATP: %w", err)
	}
	return atp, nil
}

// GetATP retrieves an ATP by ID
func (s *LearningPlanningService) GetATP(ctx context.Context, id string) (*domain.ATP, error) {
	return s.atpRepo.GetATPByID(ctx, id)
}

// CreateModulAjarSet creates a new Modul Ajar Set
func (s *LearningPlanningService) CreateModulAjarSet(ctx context.Context, req *domain.CreateModulAjarSetRequest, userID string) (*domain.ModulAjarSet, error) {
	set := &domain.ModulAjarSet{
		ID:               uuid.New().String(),
		ATPSetID:         req.ATPSetID,
		VersionNo:        req.VersionNo,
		Status:           domain.WorkflowStatusDraft,
		GenerationSource: req.GenerationSource,
		GenerationReason: req.GenerationReason,
		GeneratedBy:      userID,
	}
	if err := s.modulAjarRepo.CreateModulAjarSet(ctx, set); err != nil {
		return nil, fmt.Errorf("failed to create modul ajar set: %w", err)
	}
	return set, nil
}

// GetModulAjarSet retrieves a Modul Ajar Set by ID
func (s *LearningPlanningService) GetModulAjarSet(ctx context.Context, id string) (*domain.ModulAjarSet, error) {
	return s.modulAjarRepo.GetModulAjarSetByID(ctx, id)
}

// ListModulAjarSets retrieves Modul Ajar Sets with optional filters
func (s *LearningPlanningService) ListModulAjarSets(ctx context.Context, atpSetID *string, status *domain.WorkflowStatus, page, pageSize int) ([]*domain.ModulAjarSet, int, error) {
	limit := pageSize
	offset := (page - 1) * pageSize
	sets, err := s.modulAjarRepo.ListModulAjarSets(ctx, atpSetID, status, limit, offset)
	return sets, len(sets), err
}

// CreateModulAjar creates a new Modul Ajar
func (s *LearningPlanningService) CreateModulAjar(ctx context.Context, req *domain.CreateModulAjarRequest) (*domain.ModulAjar, error) {
	modulAjar := &domain.ModulAjar{
		ID:                  uuid.New().String(),
		ModulAjarSetID:      req.ModulAjarSetID,
		ATPID:               req.ATPID,
		Week:                req.Week,
		Topic:               req.Topic,
		Resources:           req.Resources,
		ClassCharacteristics: req.ClassCharacteristics,
		LearningActivities:  req.LearningActivities,
		ResourceRequirements: req.ResourceRequirements,
		AssessmentMethods:   req.AssessmentMethods,
		Status:              domain.WorkflowStatusDraft,
	}
	if err := s.modulAjarRepo.CreateModulAjar(ctx, modulAjar); err != nil {
		return nil, fmt.Errorf("failed to create modul ajar: %w", err)
	}
	return modulAjar, nil
}

// GetModulAjar retrieves a Modul Ajar by ID
func (s *LearningPlanningService) GetModulAjar(ctx context.Context, id string) (*domain.ModulAjar, error) {
	return s.modulAjarRepo.GetModulAjarByID(ctx, id)
}

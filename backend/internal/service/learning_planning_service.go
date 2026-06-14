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
	atpRepo       *repository.LearningPlanningRepository
	modulAjarRepo *repository.LearningPlanningRepository
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
	return sets, len(sets), fmt.Errorf("failed to list ATP sets: %w", err)
}

// CreateATP creates a new ATP
func (s *LearningPlanningService) CreateATP(ctx context.Context, req *domain.CreateATPRequest) (*domain.ATP, error) {
	atp := &domain.ATP{
		ID:                 uuid.New().String(),
		ATPSetID:           req.ATPSetID,
		TPID:               req.TPID,
		UserID:             req.UserID,
		Status:             domain.WorkflowStatusDraft,
		AcademicCalendar:   req.AcademicCalendar,
		ClassSchedule:      req.ClassSchedule,
		WeeklySequence:     req.WeeklySequence,
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
	return sets, len(sets), fmt.Errorf("failed to list Modul Ajar sets: %w", err)
}

// CreateModulAjar creates a new Modul Ajar
func (s *LearningPlanningService) CreateModulAjar(ctx context.Context, req *domain.CreateModulAjarRequest) (*domain.ModulAjar, error) {
	modulAjar := &domain.ModulAjar{
		ID:                   uuid.New().String(),
		ModulAjarSetID:       req.ModulAjarSetID,
		ATPID:                req.ATPID,
		Week:                 req.Week,
		Topic:                req.Topic,
		Resources:            req.Resources,
		ClassCharacteristics: req.ClassCharacteristics,
		LearningActivities:   req.LearningActivities,
		ResourceRequirements: req.ResourceRequirements,
		AssessmentMethods:    req.AssessmentMethods,
		Status:               domain.WorkflowStatusDraft,
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

// GetATPSetDetail retrieves an ATP Set with detailed information
func (s *LearningPlanningService) GetATPSetDetail(ctx context.Context, id string) (*domain.ATPSet, error) {
	return s.atpRepo.GetATPSetByID(ctx, id)
}

// GetModulAjarSetDetail retrieves a Modul Ajar Set with detailed information
func (s *LearningPlanningService) GetModulAjarSetDetail(ctx context.Context, id string) (*domain.ModulAjarSet, error) {
	return s.modulAjarRepo.GetModulAjarSetByID(ctx, id)
}

// UpdateATPSet updates an existing ATP Set
func (s *LearningPlanningService) UpdateATPSet(ctx context.Context, id string, req *domain.UpdateATPSetRequest) (*domain.ATPSet, error) {
	atpSet, err := s.atpRepo.GetATPSetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get ATP set: %w", err)
	}

	if req.Status != nil {
		atpSet.Status = *req.Status
	}
	if req.VersionNo != nil {
		atpSet.VersionNo = *req.VersionNo
	}

	if err := s.atpRepo.UpdateATPSet(ctx, atpSet); err != nil {
		return nil, fmt.Errorf("failed to update ATP set: %w", err)
	}

	return atpSet, nil
}

// DeleteATPSet deletes an ATP Set
func (s *LearningPlanningService) DeleteATPSet(ctx context.Context, id string) error {
	return s.atpRepo.DeleteATPSet(ctx, id)
}

// ApproveATPSet approves an ATP Set
func (s *LearningPlanningService) ApproveATPSet(ctx context.Context, id string, userID string) (*domain.ATPSet, error) {
	atpSet, err := s.atpRepo.GetATPSetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get ATP set: %w", err)
	}

	atpSet.Status = domain.WorkflowStatusApproved
	if err := s.atpRepo.UpdateATPSet(ctx, atpSet); err != nil {
		return nil, fmt.Errorf("failed to approve ATP set: %w", err)
	}

	return atpSet, nil
}

// UpdateATP updates an existing ATP
func (s *LearningPlanningService) UpdateATP(ctx context.Context, id string, req *domain.UpdateATPRequest) (*domain.ATP, error) {
	atp, err := s.atpRepo.GetATPByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get ATP: %w", err)
	}

	if req.TPID != nil {
		atp.TPID = *req.TPID
	}
	if req.AcademicCalendar != nil {
		atp.AcademicCalendar = req.AcademicCalendar
	}
	if req.ClassSchedule != nil {
		atp.ClassSchedule = req.ClassSchedule
	}
	if req.WeeklySequence != nil {
		atp.WeeklySequence = req.WeeklySequence
	}
	if req.AssessmentSchedule != nil {
		atp.AssessmentSchedule = req.AssessmentSchedule
	}
	if req.Status != nil {
		atp.Status = *req.Status
	}

	if err := s.atpRepo.UpdateATP(ctx, atp); err != nil {
		return nil, fmt.Errorf("failed to update ATP: %w", err)
	}

	return atp, nil
}

// DeleteATP deletes an ATP
func (s *LearningPlanningService) DeleteATP(ctx context.Context, id string) error {
	return s.atpRepo.DeleteATP(ctx, id)
}

// UpdateModulAjarSet updates an existing Modul Ajar Set
func (s *LearningPlanningService) UpdateModulAjarSet(ctx context.Context, id string, req *domain.UpdateModulAjarSetRequest) (*domain.ModulAjarSet, error) {
	set, err := s.modulAjarRepo.GetModulAjarSetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get modul ajar set: %w", err)
	}

	if req.Status != nil {
		set.Status = *req.Status
	}
	if req.VersionNo != nil {
		set.VersionNo = *req.VersionNo
	}

	if err := s.modulAjarRepo.UpdateModulAjarSet(ctx, set); err != nil {
		return nil, fmt.Errorf("failed to update modul ajar set: %w", err)
	}

	return set, nil
}

// DeleteModulAjarSet deletes a Modul Ajar Set
func (s *LearningPlanningService) DeleteModulAjarSet(ctx context.Context, id string) error {
	return s.modulAjarRepo.DeleteModulAjarSet(ctx, id)
}

// ApproveModulAjarSet approves a Modul Ajar Set
func (s *LearningPlanningService) ApproveModulAjarSet(ctx context.Context, id string, userID string) (*domain.ModulAjarSet, error) {
	set, err := s.modulAjarRepo.GetModulAjarSetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get modul ajar set: %w", err)
	}

	set.Status = domain.WorkflowStatusApproved
	if err := s.modulAjarRepo.UpdateModulAjarSet(ctx, set); err != nil {
		return nil, fmt.Errorf("failed to approve modul ajar set: %w", err)
	}

	return set, nil
}

// UpdateModulAjar updates an existing Modul Ajar
func (s *LearningPlanningService) UpdateModulAjar(ctx context.Context, id string, req *domain.UpdateModulAjarRequest) (*domain.ModulAjar, error) {
	modulAjar, err := s.modulAjarRepo.GetModulAjarByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get modul ajar: %w", err)
	}

	if req.Topic != nil {
		modulAjar.Topic = req.Topic
	}
	if req.Resources != nil {
		modulAjar.Resources = req.Resources
	}
	if req.ClassCharacteristics != nil {
		modulAjar.ClassCharacteristics = req.ClassCharacteristics
	}
	if req.LearningActivities != nil {
		modulAjar.LearningActivities = req.LearningActivities
	}
	if req.ResourceRequirements != nil {
		modulAjar.ResourceRequirements = req.ResourceRequirements
	}
	if req.AssessmentMethods != nil {
		modulAjar.AssessmentMethods = req.AssessmentMethods
	}
	if req.Status != nil {
		modulAjar.Status = *req.Status
	}

	if err := s.modulAjarRepo.UpdateModulAjar(ctx, modulAjar); err != nil {
		return nil, fmt.Errorf("failed to update modul ajar: %w", err)
	}

	return modulAjar, nil
}

// DeleteModulAjar deletes a Modul Ajar
func (s *LearningPlanningService) DeleteModulAjar(ctx context.Context, id string) error {
	return s.modulAjarRepo.DeleteModulAjar(ctx, id)
}

package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/repository"
)

// CurriculumService handles business logic for curriculum operations
type CurriculumService struct {
	curriculumRepo *repository.CurriculumRepository
}

// NewCurriculumService creates a new curriculum service
func NewCurriculumService(curriculumRepo *repository.CurriculumRepository) *CurriculumService {
	return &CurriculumService{curriculumRepo: curriculumRepo}
}

// CreateCurriculumSubject creates a new curriculum subject
func (s *CurriculumService) CreateCurriculumSubject(ctx context.Context, req *domain.CreateCurriculumSubjectRequest) (*domain.CurriculumSubject, error) {
	subject := &domain.CurriculumSubject{
		ID:          uuid.New().String(),
		Code:        req.Code,
		Name:        req.Name,
		NameEN:      req.NameEN,
		Description: req.Description,
		IsActive:    true,
	}

	if err := s.curriculumRepo.CreateCurriculumSubject(ctx, subject); err != nil {
		return nil, fmt.Errorf("failed to create curriculum subject: %w", err)
	}

	return subject, nil
}

// GetCurriculumSubject retrieves a curriculum subject by ID
func (s *CurriculumService) GetCurriculumSubject(ctx context.Context, id string) (*domain.CurriculumSubject, error) {
	return s.curriculumRepo.GetCurriculumSubjectByID(ctx, id)
}

// ListCurriculumSubjects retrieves curriculum subjects with optional filters
func (s *CurriculumService) ListCurriculumSubjects(ctx context.Context, isActive *bool, page, pageSize int) ([]*domain.CurriculumSubject, int, error) {
	limit := pageSize
	offset := (page - 1) * pageSize

	subjects, err := s.curriculumRepo.ListCurriculumSubjects(ctx, isActive, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return subjects, len(subjects), nil
}

// UpdateCurriculumSubject updates a curriculum subject
func (s *CurriculumService) UpdateCurriculumSubject(ctx context.Context, id string, req *domain.UpdateCurriculumSubjectRequest) (*domain.CurriculumSubject, error) {
	subject, err := s.curriculumRepo.GetCurriculumSubjectByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("curriculum subject not found")
	}

	if req.Name != nil {
		subject.Name = *req.Name
	}
	if req.NameEN != nil {
		subject.NameEN = req.NameEN
	}
	if req.Description != nil {
		subject.Description = req.Description
	}
	if req.IsActive != nil {
		subject.IsActive = *req.IsActive
	}

	if err := s.curriculumRepo.UpdateCurriculumSubject(ctx, subject); err != nil {
		return nil, fmt.Errorf("failed to update curriculum subject: %w", err)
	}

	return subject, nil
}

// CreateCurriculumPhase creates a new curriculum phase
func (s *CurriculumService) CreateCurriculumPhase(ctx context.Context, req *domain.CreateCurriculumPhaseRequest) (*domain.CurriculumPhase, error) {
	phase := &domain.CurriculumPhase{
		ID:              uuid.New().String(),
		Code:            req.Code,
		Name:            req.Name,
		NameEN:          req.NameEN,
		Description:     req.Description,
		GradeLevelStart: req.GradeLevelStart,
		GradeLevelEnd:   req.GradeLevelEnd,
		IsActive:        true,
	}

	if err := s.curriculumRepo.CreateCurriculumPhase(ctx, phase); err != nil {
		return nil, fmt.Errorf("failed to create curriculum phase: %w", err)
	}

	return phase, nil
}

// GetCurriculumPhase retrieves a curriculum phase by ID
func (s *CurriculumService) GetCurriculumPhase(ctx context.Context, id string) (*domain.CurriculumPhase, error) {
	return s.curriculumRepo.GetCurriculumPhaseByID(ctx, id)
}

// ListCurriculumPhases retrieves curriculum phases with optional filters
func (s *CurriculumService) ListCurriculumPhases(ctx context.Context, isActive *bool, page, pageSize int) ([]*domain.CurriculumPhase, int, error) {
	limit := pageSize
	offset := (page - 1) * pageSize

	phases, err := s.curriculumRepo.ListCurriculumPhases(ctx, isActive, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return phases, len(phases), nil
}

// UpdateCurriculumPhase updates a curriculum phase
func (s *CurriculumService) UpdateCurriculumPhase(ctx context.Context, id string, req *domain.UpdateCurriculumPhaseRequest) (*domain.CurriculumPhase, error) {
	phase, err := s.curriculumRepo.GetCurriculumPhaseByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("curriculum phase not found")
	}

	if req.Name != nil {
		phase.Name = *req.Name
	}
	if req.NameEN != nil {
		phase.NameEN = req.NameEN
	}
	if req.Description != nil {
		phase.Description = req.Description
	}
	if req.GradeLevelStart != nil {
		phase.GradeLevelStart = req.GradeLevelStart
	}
	if req.GradeLevelEnd != nil {
		phase.GradeLevelEnd = req.GradeLevelEnd
	}
	if req.IsActive != nil {
		phase.IsActive = *req.IsActive
	}

	if err := s.curriculumRepo.UpdateCurriculumPhase(ctx, phase); err != nil {
		return nil, fmt.Errorf("failed to update curriculum phase: %w", err)
	}

	return phase, nil
}

// CreateCurriculumElement creates a new curriculum element
func (s *CurriculumService) CreateCurriculumElement(ctx context.Context, req *domain.CreateCurriculumElementRequest) (*domain.CurriculumElement, error) {
	subject, err := s.curriculumRepo.GetCurriculumSubjectByID(ctx, req.SubjectID)
	if err != nil {
		return nil, fmt.Errorf("invalid subject: %w", err)
	}
	if !subject.IsActive {
		return nil, fmt.Errorf("subject is not active")
	}

	phase, err := s.curriculumRepo.GetCurriculumPhaseByID(ctx, req.PhaseID)
	if err != nil {
		return nil, fmt.Errorf("invalid phase: %w", err)
	}
	if !phase.IsActive {
		return nil, fmt.Errorf("phase is not active")
	}

	element := &domain.CurriculumElement{
		ID:        uuid.New().String(),
		SubjectID:  req.SubjectID,
		PhaseID:    req.PhaseID,
		Code:       req.Code,
		Name:       req.Name,
		NameEN:     req.NameEN,
		Description: req.Description,
		IsActive:   true,
	}

	if err := s.curriculumRepo.CreateCurriculumElement(ctx, element); err != nil {
		return nil, fmt.Errorf("failed to create curriculum element: %w", err)
	}

	return element, nil
}

// GetCurriculumElement retrieves a curriculum element by ID
func (s *CurriculumService) GetCurriculumElement(ctx context.Context, id string) (*domain.CurriculumElement, error) {
	return s.curriculumRepo.GetCurriculumElementByID(ctx, id)
}

// ListCurriculumElements retrieves curriculum elements with optional filters
func (s *CurriculumService) ListCurriculumElements(ctx context.Context, subjectID, phaseID *string, isActive *bool, page, pageSize int) ([]*domain.CurriculumElement, int, error) {
	limit := pageSize
	offset := (page - 1) * pageSize
	elements, err := s.curriculumRepo.ListCurriculumElements(ctx, subjectID, phaseID, isActive, limit, offset)
	return elements, len(elements), err
}

// UpdateCurriculumElement updates a curriculum element
func (s *CurriculumService) UpdateCurriculumElement(ctx context.Context, id string, req *domain.UpdateCurriculumElementRequest) (*domain.CurriculumElement, error) {
	element, err := s.curriculumRepo.GetCurriculumElementByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("curriculum element not found")
	}

	if req.Name != nil {
		element.Name = *req.Name
	}
	if req.NameEN != nil {
		element.NameEN = req.NameEN
	}
	if req.Description != nil {
		element.Description = req.Description
	}
	if req.IsActive != nil {
		element.IsActive = *req.IsActive
	}

	if err := s.curriculumRepo.UpdateCurriculumElement(ctx, element); err != nil {
		return nil, fmt.Errorf("failed to update curriculum element: %w", err)
	}

	return element, nil
}

// CreateCurriculumSubelement creates a new curriculum subelement
func (s *CurriculumService) CreateCurriculumSubelement(ctx context.Context, req *domain.CreateCurriculumSubelementRequest) (*domain.CurriculumSubelement, error) {
	element, err := s.curriculumRepo.GetCurriculumElementByID(ctx, req.ElementID)
	if err != nil {
		return nil, fmt.Errorf("invalid element: %w", err)
	}
	if !element.IsActive {
		return nil, fmt.Errorf("element is not active")
	}

	subelement := &domain.CurriculumSubelement{
		ID:        uuid.New().String(),
		ElementID:  req.ElementID,
		Code:       req.Code,
		Name:       req.Name,
		NameEN:     req.NameEN,
		Description: req.Description,
		IsActive:   true,
	}

	if err := s.curriculumRepo.CreateCurriculumSubelement(ctx, subelement); err != nil {
		return nil, fmt.Errorf("failed to create curriculum subelement: %w", err)
	}

	return subelement, nil
}

// GetCurriculumSubelement retrieves a curriculum subelement by ID
func (s *CurriculumService) GetCurriculumSubelement(ctx context.Context, id string) (*domain.CurriculumSubelement, error) {
	return s.curriculumRepo.GetCurriculumSubelementByID(ctx, id)
}

// ListCurriculumSubelements retrieves curriculum subelements with optional filters
func (s *CurriculumService) ListCurriculumSubelements(ctx context.Context, elementID *string, isActive *bool, page, pageSize int) ([]*domain.CurriculumSubelement, int, error) {
	limit := pageSize
	offset := (page - 1) * pageSize
	subelements, err := s.curriculumRepo.ListCurriculumSubelements(ctx, elementID, isActive, limit, offset)
	return subelements, len(subelements), err
}

// ImportCP imports CP data in bulk
func (s *CurriculumService) ImportCP(ctx context.Context, req *domain.ImportCPRequest, importerID string) ([]*domain.CP, error) {
	var cps []*domain.CP

	for _, cpReq := range req.CPs {
		if _, err := s.curriculumRepo.GetCurriculumSubjectByID(ctx, cpReq.SubjectID); err != nil {
			return nil, fmt.Errorf("invalid subject for CP %s: %w", cpReq.Code, err)
		}
		if _, err := s.curriculumRepo.GetCurriculumPhaseByID(ctx, cpReq.PhaseID); err != nil {
			return nil, fmt.Errorf("invalid phase for CP %s: %w", cpReq.Code, err)
		}
		if _, err := s.curriculumRepo.GetCurriculumElementByID(ctx, cpReq.ElementID); err != nil {
			return nil, fmt.Errorf("invalid element for CP %s: %w", cpReq.Code, err)
		}
		if _, err := s.curriculumRepo.GetCurriculumSubelementByID(ctx, cpReq.SubelementID); err != nil {
			return nil, fmt.Errorf("invalid subelement for CP %s: %w", cpReq.Code, err)
		}

		cp := &domain.CP{
			ID:                  uuid.New().String(),
			SubjectID:           cpReq.SubjectID,
			PhaseID:             cpReq.PhaseID,
			ElementID:           cpReq.ElementID,
			SubelementID:        cpReq.SubelementID,
			Code:                cpReq.Code,
			Description:         cpReq.Description,
			CompetencyCode:      cpReq.CompetencyCode,
			LearningObjectives:   cpReq.LearningObjectives,
			CompetencyStandards: cpReq.CompetencyStandards,
			TimeAllocationHours:  cpReq.TimeAllocationHours,
			HoursPerWeek:         cpReq.HoursPerWeek,
			Version:             cpReq.Version,
			IsActive:            true,
			ImportedBy:          &importerID,
		}

		if err := s.curriculumRepo.CreateCP(ctx, cp); err != nil {
			return nil, fmt.Errorf("failed to create CP %s: %w", cpReq.Code, err)
		}

		cps = append(cps, cp)
	}

	return cps, nil
}

// GetCP retrieves a CP by ID
func (s *CurriculumService) GetCP(ctx context.Context, id string) (*domain.CP, error) {
	return s.curriculumRepo.GetCPByID(ctx, id)
}

// ListCPs retrieves CPs with optional filters
func (s *CurriculumService) ListCPs(ctx context.Context, subjectID, phaseID, elementID *string, version *string, isActive *bool, page, pageSize int) ([]*domain.CP, int, error) {
	limit := pageSize
	offset := (page - 1) * pageSize
	cps, err := s.curriculumRepo.ListCPs(ctx, subjectID, phaseID, elementID, version, isActive, limit, offset)
	return cps, len(cps), err
}

// UpdateCP updates a CP
func (s *CurriculumService) UpdateCP(ctx context.Context, id string, req *domain.UpdateCPRequest) (*domain.CP, error) {
	cp, err := s.curriculumRepo.GetCPByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("cp not found")
	}

	if req.Description != nil {
		cp.Description = *req.Description
	}
	if req.CompetencyCode != nil {
		cp.CompetencyCode = req.CompetencyCode
	}
	if req.LearningObjectives != nil {
		cp.LearningObjectives = req.LearningObjectives
	}
	if req.CompetencyStandards != nil {
		cp.CompetencyStandards = req.CompetencyStandards
	}
	if req.TimeAllocationHours != nil {
		cp.TimeAllocationHours = *req.TimeAllocationHours
	}
	if req.HoursPerWeek != nil {
		cp.HoursPerWeek = *req.HoursPerWeek
	}
	if req.Version != nil {
		cp.Version = *req.Version
	}
	if req.IsActive != nil {
		cp.IsActive = *req.IsActive
	}

	if err := s.curriculumRepo.UpdateCP(ctx, cp); err != nil {
		return nil, fmt.Errorf("failed to update CP: %w", err)
	}

	return cp, nil
}

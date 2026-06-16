package service

import (
	"context"
	"fmt"
	"time"

	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/repository"
)

// AssignmentService handles business logic for assignment operations
type AssignmentService struct {
	assignmentRepo repository.AssignmentRepositoryInterface
	classRepo      repository.ClassRepositoryInterface
	assessmentRepo repository.AssessmentRepositoryInterface
	userRepo       repository.UserRepositoryInterface
}

// NewAssignmentService creates a new assignment service
func NewAssignmentService(
	assignmentRepo repository.AssignmentRepositoryInterface,
	classRepo repository.ClassRepositoryInterface,
	assessmentRepo repository.AssessmentRepositoryInterface,
	userRepo repository.UserRepositoryInterface,
) *AssignmentService {
	return &AssignmentService{
		assignmentRepo: assignmentRepo,
		classRepo:      classRepo,
		assessmentRepo: assessmentRepo,
		userRepo:       userRepo,
	}
}

// CreateAssignment creates a new assignment
func (s *AssignmentService) CreateAssignment(ctx context.Context, req *domain.CreateAssignmentRequest, creatorID string) (*domain.Assignment, error) {
	// Verify class exists and is active
	class, err := s.classRepo.GetByID(ctx, req.ClassID)
	if err != nil {
		return nil, fmt.Errorf("class not found")
	}
	if !class.IsActive {
		return nil, fmt.Errorf("class is not active")
	}

	// Verify assessment exists
	_, err = s.assessmentRepo.GetAssessmentByID(ctx, req.AssessmentID)
	if err != nil {
		return nil, fmt.Errorf("assessment not found")
	}

	// Verify creator exists and is active
	creator, err := s.userRepo.GetByID(ctx, creatorID)
	if err != nil {
		return nil, fmt.Errorf("creator not found")
	}
	if !creator.IsActive {
		return nil, fmt.Errorf("creator is not active")
	}

	assignment := &domain.Assignment{
		ID:           domain.NewID(),
		ClassID:      req.ClassID,
		AssessmentID: req.AssessmentID,
		Title:        req.Title,
		Description:  req.Description,
		DueDate:      req.DueDate,
		MaxScore:     req.MaxScore,
		Status:       string(domain.AssignmentStatusAssigned),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		CreatedBy:    &creatorID,
		UpdatedBy:    &creatorID,
	}

	if err := s.assignmentRepo.Create(ctx, assignment); err != nil {
		return nil, fmt.Errorf("failed to create assignment: %w", err)
	}

	return assignment, nil
}

// GetAssignment retrieves an assignment by ID
func (s *AssignmentService) GetAssignment(ctx context.Context, id string) (*domain.Assignment, error) {
	assignment, err := s.assignmentRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("assignment not found")
	}
	return assignment, nil
}

// ListAssignments retrieves assignments with filters and pagination
func (s *AssignmentService) ListAssignments(ctx context.Context, classID, assessmentID, status *string, page, pageSize int) ([]*domain.Assignment, int, error) {
	limit := pageSize
	offset := (page - 1) * pageSize

	assignments, err := s.assignmentRepo.List(ctx, classID, assessmentID, status, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list assignments: %w", err)
	}

	total, err := s.assignmentRepo.Count(ctx, classID, assessmentID, status)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count assignments: %w", err)
	}

	return assignments, total, nil
}

// UpdateAssignment updates assignment information
func (s *AssignmentService) UpdateAssignment(ctx context.Context, id string, req *domain.UpdateAssignmentRequest, updaterID string) (*domain.Assignment, error) {
	assignment, err := s.assignmentRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("assignment not found")
	}

	if req.Title != nil {
		assignment.Title = *req.Title
	}
	if req.Description != nil {
		assignment.Description = req.Description
	}
	if req.DueDate != nil {
		assignment.DueDate = *req.DueDate
	}
	if req.MaxScore != nil {
		assignment.MaxScore = *req.MaxScore
	}
	if req.Status != nil {
		assignment.Status = string(*req.Status)
	}

	assignment.UpdatedBy = &updaterID

	if err := s.assignmentRepo.Update(ctx, assignment); err != nil {
		return nil, fmt.Errorf("failed to update assignment: %w", err)
	}

	return assignment, nil
}

// DeleteAssignment soft deletes an assignment
func (s *AssignmentService) DeleteAssignment(ctx context.Context, id string) error {
	return s.assignmentRepo.Delete(ctx, id)
}

// GetClassAssignments retrieves all assignments for a class
func (s *AssignmentService) GetClassAssignments(ctx context.Context, classID string) ([]*domain.Assignment, error) {
	// Verify class exists
	_, err := s.classRepo.GetByID(ctx, classID)
	if err != nil {
		return nil, fmt.Errorf("class not found")
	}

	assignments, err := s.assignmentRepo.GetByClassID(ctx, classID)
	if err != nil {
		return nil, fmt.Errorf("failed to get class assignments: %w", err)
	}

	return assignments, nil
}

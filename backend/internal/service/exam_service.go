package service

import (
	"context"
	"fmt"
	"time"

	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/repository"
)

// ExamService handles business logic for exam operations
type ExamService struct {
	examRepo       *repository.ExamRepository
	classRepo      *repository.ClassRepository
	assessmentRepo *repository.AssessmentRepository
	userRepo       *repository.UserRepository
}

// NewExamService creates a new exam service
func NewExamService(
	examRepo *repository.ExamRepository,
	classRepo *repository.ClassRepository,
	assessmentRepo *repository.AssessmentRepository,
	userRepo *repository.UserRepository,
) *ExamService {
	return &ExamService{
		examRepo:       examRepo,
		classRepo:      classRepo,
		assessmentRepo: assessmentRepo,
		userRepo:       userRepo,
	}
}

// CreateExam creates a new exam
func (s *ExamService) CreateExam(ctx context.Context, req *domain.CreateExamRequest, creatorID string) (*domain.Exam, error) {
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

	exam := &domain.Exam{
		ID:              domain.NewID(),
		ClassID:         req.ClassID,
		AssessmentID:    req.AssessmentID,
		ExamDate:        req.ExamDate,
		StartTime:       req.StartTime,
		DurationMinutes: req.DurationMinutes,
		Room:            req.Room,
		Status:          string(domain.ExamStatusScheduled),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		CreatedBy:       &creatorID,
		UpdatedBy:       &creatorID,
	}

	if err := s.examRepo.Create(ctx, exam); err != nil {
		return nil, fmt.Errorf("failed to create exam: %w", err)
	}

	return exam, nil
}

// GetExam retrieves an exam by ID
func (s *ExamService) GetExam(ctx context.Context, id string) (*domain.Exam, error) {
	exam, err := s.examRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("exam not found")
	}
	return exam, nil
}

// ListExams retrieves exams with filters and pagination
func (s *ExamService) ListExams(ctx context.Context, classID, assessmentID, status *string, page, pageSize int) ([]*domain.Exam, int, error) {
	limit := pageSize
	offset := (page - 1) * pageSize

	exams, err := s.examRepo.List(ctx, classID, assessmentID, status, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list exams: %w", err)
	}

	total, err := s.examRepo.Count(ctx, classID, assessmentID, status)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count exams: %w", err)
	}

	return exams, total, nil
}

// UpdateExam updates exam information
func (s *ExamService) UpdateExam(ctx context.Context, id string, req *domain.UpdateExamRequest, updaterID string) (*domain.Exam, error) {
	exam, err := s.examRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("exam not found")
	}

	if req.ExamDate != nil {
		exam.ExamDate = *req.ExamDate
	}
	if req.StartTime != nil {
		exam.StartTime = *req.StartTime
	}
	if req.DurationMinutes != nil {
		exam.DurationMinutes = *req.DurationMinutes
	}
	if req.Room != nil {
		exam.Room = req.Room
	}
	if req.Status != nil {
		exam.Status = string(*req.Status)
	}

	exam.UpdatedBy = &updaterID

	if err := s.examRepo.Update(ctx, exam); err != nil {
		return nil, fmt.Errorf("failed to update exam: %w", err)
	}

	return exam, nil
}

// DeleteExam soft deletes an exam
func (s *ExamService) DeleteExam(ctx context.Context, id string) error {
	return s.examRepo.Delete(ctx, id)
}

// GetClassExams retrieves all exams for a class
func (s *ExamService) GetClassExams(ctx context.Context, classID string) ([]*domain.Exam, error) {
	// Verify class exists
	_, err := s.classRepo.GetByID(ctx, classID)
	if err != nil {
		return nil, fmt.Errorf("class not found")
	}

	exams, err := s.examRepo.GetByClassID(ctx, classID)
	if err != nil {
		return nil, fmt.Errorf("failed to get class exams: %w", err)
	}

	return exams, nil
}

package service

import (
	"context"
	"fmt"
	"time"

	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/repository"
)

// ExamResultService handles business logic for exam result operations
type ExamResultService struct {
	examResultRepo *repository.ExamResultRepository
	examRepo       *repository.ExamRepository
	userRepo       *repository.UserRepository
}

// NewExamResultService creates a new exam result service
func NewExamResultService(
	examResultRepo *repository.ExamResultRepository,
	examRepo *repository.ExamRepository,
	userRepo *repository.UserRepository,
) *ExamResultService {
	return &ExamResultService{
		examResultRepo: examResultRepo,
		examRepo:       examRepo,
		userRepo:       userRepo,
	}
}

// CreateExamResult creates a new exam result
func (s *ExamResultService) CreateExamResult(ctx context.Context, req *domain.CreateExamResultRequest, graderID string) (*domain.ExamResult, error) {
	// Verify exam exists
	_, err := s.examRepo.GetByID(ctx, req.ExamID)
	if err != nil {
		return nil, fmt.Errorf("exam not found")
	}

	// Verify student exists and is active
	student, err := s.userRepo.GetByID(ctx, req.StudentID)
	if err != nil {
		return nil, fmt.Errorf("student not found")
	}
	if !student.IsActive {
		return nil, fmt.Errorf("student is not active")
	}

	// Verify grader exists and is active (if grading is being done)
	if req.Score != nil {
		grader, err := s.userRepo.GetByID(ctx, graderID)
		if err != nil {
			return nil, fmt.Errorf("grader not found")
		}
		if !grader.IsActive {
			return nil, fmt.Errorf("grader is not active")
		}
	}

	examResult := &domain.ExamResult{
		ID:        domain.NewID(),
		ExamID:    req.ExamID,
		StudentID: req.StudentID,
		Score:     req.Score,
		Grade:     req.Grade,
		Remarks:   req.Remarks,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// If score is provided, mark as graded
	if req.Score != nil {
		examResult.MarkAsGraded(graderID)
	}

	if err := s.examResultRepo.Create(ctx, examResult); err != nil {
		return nil, fmt.Errorf("failed to create exam result: %w", err)
	}

	return examResult, nil
}

// GetExamResult retrieves an exam result by ID
func (s *ExamResultService) GetExamResult(ctx context.Context, id string) (*domain.ExamResult, error) {
	examResult, err := s.examResultRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("exam result not found")
	}
	return examResult, nil
}

// GetExamResultByExamAndStudent retrieves an exam result by exam ID and student ID
func (s *ExamResultService) GetExamResultByExamAndStudent(ctx context.Context, examID, studentID string) (*domain.ExamResult, error) {
	examResult, err := s.examResultRepo.GetByExamAndStudent(ctx, examID, studentID)
	if err != nil {
		return nil, fmt.Errorf("exam result not found")
	}
	return examResult, nil
}

// ListExamResults retrieves exam results with filters and pagination
func (s *ExamResultService) ListExamResults(ctx context.Context, examID, studentID, grade *string, page, pageSize int) ([]*domain.ExamResult, int, error) {
	limit := pageSize
	offset := (page - 1) * pageSize

	examResults, err := s.examResultRepo.List(ctx, examID, studentID, grade, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list exam results: %w", err)
	}

	total, err := s.examResultRepo.Count(ctx, examID, studentID, grade)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count exam results: %w", err)
	}

	return examResults, total, nil
}

// UpdateExamResult updates exam result information
func (s *ExamResultService) UpdateExamResult(ctx context.Context, id string, req *domain.UpdateExamResultRequest, graderID string) (*domain.ExamResult, error) {
	examResult, err := s.examResultRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("exam result not found")
	}

	if req.Score != nil {
		examResult.Score = req.Score
	}
	if req.Grade != nil {
		examResult.Grade = req.Grade
	}
	if req.Remarks != nil {
		examResult.Remarks = req.Remarks
	}
	if req.GradedBy != nil {
		examResult.GradedBy = req.GradedBy
	}

	// If score is being updated, mark as graded
	if req.Score != nil {
		examResult.MarkAsGraded(graderID)
	}

	if err := s.examResultRepo.Update(ctx, examResult); err != nil {
		return nil, fmt.Errorf("failed to update exam result: %w", err)
	}

	return examResult, nil
}

// DeleteExamResult soft deletes an exam result
func (s *ExamResultService) DeleteExamResult(ctx context.Context, id string) error {
	return s.examResultRepo.Delete(ctx, id)
}

// GetExamResultsByExam retrieves all exam results for an exam
func (s *ExamResultService) GetExamResultsByExam(ctx context.Context, examID string) ([]*domain.ExamResult, error) {
	// Verify exam exists
	_, err := s.examRepo.GetByID(ctx, examID)
	if err != nil {
		return nil, fmt.Errorf("exam not found")
	}

	examResults, err := s.examResultRepo.GetByExamID(ctx, examID)
	if err != nil {
		return nil, fmt.Errorf("failed to get exam results by exam: %w", err)
	}

	return examResults, nil
}

// GetExamResultsByStudent retrieves all exam results for a student
func (s *ExamResultService) GetExamResultsByStudent(ctx context.Context, studentID string) ([]*domain.ExamResult, error) {
	// Verify student exists and is active
	student, err := s.userRepo.GetByID(ctx, studentID)
	if err != nil {
		return nil, fmt.Errorf("student not found")
	}
	if !student.IsActive {
		return nil, fmt.Errorf("student is not active")
	}

	examResults, err := s.examResultRepo.GetByStudentID(ctx, studentID)
	if err != nil {
		return nil, fmt.Errorf("failed to get exam results by student: %w", err)
	}

	return examResults, nil
}

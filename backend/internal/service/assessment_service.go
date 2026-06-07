package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/repository"
)

// AssessmentService handles business logic for assessment operations
type AssessmentService struct {
	assessmentRepo  *repository.AssessmentRepository
}

// NewAssessmentService creates a new assessment service
func NewAssessmentService(assessmentRepo *repository.AssessmentRepository) *AssessmentService {
	return &AssessmentService{assessmentRepo: assessmentRepo}
}

// CreateAssessment creates a new assessment
func (s *AssessmentService) CreateAssessment(ctx context.Context, req *domain.CreateAssessmentRequest, userID string) (*domain.Assessment, error) {
	assessment := &domain.Assessment{
		ID:                uuid.New().String(),
		ModulAjarID:       req.ModulAjarID,
		UserID:            userID,
		AssessmentType:    req.AssessmentType,
		Status:            domain.WorkflowStatusDraft,
		AssessmentItems:   req.AssessmentItems,
		AnswerKey:         req.AnswerKey,
		ScoringGuidelines: req.ScoringGuidelines,
		VersionNo:         1,
		IsCurrentVersion:  true,
	}

	if err := s.assessmentRepo.CreateAssessment(ctx, assessment); err != nil {
		return nil, fmt.Errorf("failed to create assessment: %w", err)
	}

	return assessment, nil
}

// GetAssessment retrieves an assessment by ID
func (s *AssessmentService) GetAssessment(ctx context.Context, id string) (*domain.Assessment, error) {
	return s.assessmentRepo.GetAssessmentByID(ctx, id)
}

// ListAssessments retrieves assessments with optional filters
func (s *AssessmentService) ListAssessments(ctx context.Context, modulAjarID, userID *string, assessmentType *domain.AssessmentType, status *domain.WorkflowStatus, page, pageSize int) ([]*domain.Assessment, int, error) {
	limit := pageSize
	offset := (page - 1) * pageSize
	assessments, err := s.assessmentRepo.ListAssessments(ctx, modulAjarID, userID, assessmentType, status, limit, offset)
	return assessments, len(assessments), err
}

// UpdateAssessment updates an assessment
func (s *AssessmentService) UpdateAssessment(ctx context.Context, id string, req *domain.UpdateAssessmentRequest) (*domain.Assessment, error) {
	assessment, err := s.assessmentRepo.GetAssessmentByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("assessment not found")
	}

	if req.AssessmentItems != nil {
		assessment.AssessmentItems = req.AssessmentItems
	}
	if req.AnswerKey != nil {
		assessment.AnswerKey = req.AnswerKey
	}
	if req.ScoringGuidelines != nil {
		assessment.ScoringGuidelines = req.ScoringGuidelines
	}
	if req.Status != nil {
		assessment.Status = *req.Status
	}

	if err := s.assessmentRepo.UpdateAssessment(ctx, assessment); err != nil {
		return nil, fmt.Errorf("failed to update assessment: %w", err)
	}

	return assessment, nil
}

// CreateRubric creates a new rubric
func (s *AssessmentService) CreateRubric(ctx context.Context, req *domain.CreateRubricRequest, userID string) (*domain.Rubric, error) {
	rubric := &domain.Rubric{
		ID:                  uuid.New().String(),
		AssessmentID:         req.AssessmentID,
		UserID:              userID,
		RubricType:          req.RubricType,
		Status:              domain.WorkflowStatusDraft,
		PerformanceCriteria: req.PerformanceCriteria,
		PerformanceLevels:   req.PerformanceLevels,
		ScoringGuidelines:   req.ScoringGuidelines,
		VersionNo:           1,
		IsCurrentVersion:    true,
	}

	if err := s.assessmentRepo.CreateRubric(ctx, rubric); err != nil {
		return nil, fmt.Errorf("failed to create rubric: %w", err)
	}

	return rubric, nil
}

// GetRubric retrieves a rubric by ID
func (s *AssessmentService) GetRubric(ctx context.Context, id string) (*domain.Rubric, error) {
	return s.assessmentRepo.GetRubricByID(ctx, id)
}

// ListRubrics retrieves rubrics with optional filters
func (s *AssessmentService) ListRubrics(ctx context.Context, assessmentID, userID *string, rubricType *domain.RubricType, status *domain.WorkflowStatus, page, pageSize int) ([]*domain.Rubric, int, error) {
	limit := pageSize
	offset := (page - 1) * pageSize
	rubrics, err := s.assessmentRepo.ListRubrics(ctx, assessmentID, userID, rubricType, status, limit, offset)
	return rubrics, len(rubrics), err
}

// UpdateRubric updates a rubric
func (s *AssessmentService) UpdateRubric(ctx context.Context, id string, req *domain.UpdateRubricRequest) (*domain.Rubric, error) {
	rubric, err := s.assessmentRepo.GetRubricByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("rubric not found")
	}

	if req.PerformanceCriteria != nil {
		rubric.PerformanceCriteria = req.PerformanceCriteria
	}
	if req.PerformanceLevels != nil {
		rubric.PerformanceLevels = req.PerformanceLevels
	}
	if req.ScoringGuidelines != nil {
		rubric.ScoringGuidelines = req.ScoringGuidelines
	}
	if req.Status != nil {
		rubric.Status = *req.Status
	}

	if err := s.assessmentRepo.UpdateRubric(ctx, rubric); err != nil {
		return nil, fmt.Errorf("failed to update rubric: %w", err)
	}

	return rubric, nil
}

// CreateEvidence creates a new evidence
func (s *AssessmentService) CreateEvidence(ctx context.Context, req *domain.CreateEvidenceRequest, userID string) (*domain.Evidence, error) {
	evidence := &domain.Evidence{
		ID:           uuid.New().String(),
		StudentID:    req.StudentID,
		AssessmentID: req.AssessmentID,
		UserID:       userID,
		EvidenceType: req.EvidenceType,
		Status:       domain.EvidenceStatusCollected,
		EvidenceData: req.EvidenceData,
		TeacherNotes: req.TeacherNotes,
		RubricID:     req.RubricID,
		LinkedCriteria: req.LinkedCriteria,
	}

	if err := s.assessmentRepo.CreateEvidence(ctx, evidence); err != nil {
		return nil, fmt.Errorf("failed to create evidence: %w", err)
	}

	return evidence, nil
}

// GetEvidence retrieves an evidence by ID
func (s *AssessmentService) GetEvidence(ctx context.Context, id string) (*domain.Evidence, error) {
	return s.assessmentRepo.GetEvidenceByID(ctx, id)
}

// ListEvidences retrieves evidences with optional filters
func (s *AssessmentService) ListEvidences(ctx context.Context, studentID, assessmentID *string, evidenceType *domain.EvidenceType, status *domain.EvidenceStatus, page, pageSize int) ([]*domain.Evidence, int, error) {
	limit := pageSize
	offset := (page - 1) * pageSize
	evidences, err := s.assessmentRepo.ListEvidences(ctx, studentID, assessmentID, evidenceType, status, limit, offset)
	return evidences, len(evidences), err
}

// UpdateEvidence updates an evidence
func (s *AssessmentService) UpdateEvidence(ctx context.Context, id string, req *domain.UpdateEvidenceRequest) (*domain.Evidence, error) {
	evidence, err := s.assessmentRepo.GetEvidenceByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("evidence not found")
	}

	if req.EvidenceData != nil {
		evidence.EvidenceData = req.EvidenceData
	}
	if req.TeacherNotes != nil {
		evidence.TeacherNotes = req.TeacherNotes
	}
	if req.RubricID != nil {
		evidence.RubricID = req.RubricID
	}
	if req.LinkedCriteria != nil {
		evidence.LinkedCriteria = req.LinkedCriteria
	}
	if req.EvaluationNotes != nil {
		evidence.EvaluationNotes = req.EvaluationNotes
	}
	if req.Status != nil {
		evidence.Status = *req.Status
	}

	if err := s.assessmentRepo.UpdateEvidence(ctx, evidence); err != nil {
		return nil, fmt.Errorf("failed to update evidence: %w", err)
	}

	return evidence, nil
}

// CreateEvaluation creates a new evaluation
func (s *AssessmentService) CreateEvaluation(ctx context.Context, req *domain.CreateEvaluationRequest, userID string) (*domain.Evaluation, error) {
	evaluation := &domain.Evaluation{
		ID:               uuid.New().String(),
		StudentID:        req.StudentID,
		RubricID:         req.RubricID,
		EvidenceID:       req.EvidenceID,
		UserID:           userID,
		PerformanceScores: req.PerformanceScores,
		TotalScore:       req.TotalScore,
		MaxScore:         req.MaxScore,
		PerformanceLevel: req.PerformanceLevel,
	}

	if err := s.assessmentRepo.CreateEvaluation(ctx, evaluation); err != nil {
		return nil, fmt.Errorf("failed to create evaluation: %w", err)
	}

	return evaluation, nil
}

// GetEvaluation retrieves an evaluation by ID
func (s *AssessmentService) GetEvaluation(ctx context.Context, id string) (*domain.Evaluation, error) {
	return s.assessmentRepo.GetEvaluationByID(ctx, id)
}

// ListEvaluations retrieves evaluations with optional filters
func (s *AssessmentService) ListEvaluations(ctx context.Context, studentID, rubricID, evidenceID *string, performanceLevel *domain.PerformanceLevel, page, pageSize int) ([]*domain.Evaluation, int, error) {
	limit := pageSize
	offset := (page - 1) * pageSize
	evaluations, err := s.assessmentRepo.ListEvaluations(ctx, studentID, rubricID, evidenceID, performanceLevel, limit, offset)
	return evaluations, len(evaluations), err
}

// UpdateEvaluation updates an evaluation
func (s *AssessmentService) UpdateEvaluation(ctx context.Context, id string, req *domain.UpdateEvaluationRequest) (*domain.Evaluation, error) {
	evaluation, err := s.assessmentRepo.GetEvaluationByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("evaluation not found")
	}

	if req.PerformanceScores != nil {
		evaluation.PerformanceScores = req.PerformanceScores
	}
	if req.TotalScore != nil {
		evaluation.TotalScore = *req.TotalScore
	}
	if req.MaxScore != nil {
		evaluation.MaxScore = *req.MaxScore
	}
	if req.PerformanceLevel != nil {
		evaluation.PerformanceLevel = *req.PerformanceLevel
	}

	if err := s.assessmentRepo.UpdateEvaluation(ctx, evaluation); err != nil {
		return nil, fmt.Errorf("failed to update evaluation: %w", err)
	}

	return evaluation, nil
}

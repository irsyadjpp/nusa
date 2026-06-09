package service

import (
	"context"
	"fmt"

	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/repository"
)

// AchievementService handles business logic for achievement calculations
type AchievementService struct {
	achievementDomainService *domain.AchievementService
	assessmentRepo           *repository.AssessmentRepository
	tpRepo                   *repository.TPRepository
}

// NewAchievementService creates a new achievement service
func NewAchievementService(
	assessmentRepo *repository.AssessmentRepository,
	tpRepo *repository.TPRepository,
) *AchievementService {
	return &AchievementService{
		achievementDomainService: domain.NewAchievementService(),
		assessmentRepo:           assessmentRepo,
		tpRepo:                   tpRepo,
	}
}

// CalculateStudentAchievement calculates achievement for a specific student and TP
func (s *AchievementService) CalculateStudentAchievement(
	ctx context.Context,
	studentID string,
	studentName string,
	tpID string,
) (*domain.Achievement, error) {
	// Get TP to retrieve success criteria
	tp, err := s.tpRepo.GetTPByID(ctx, tpID)
	if err != nil {
		return nil, fmt.Errorf("failed to get TP: %w", err)
	}

	// Get evaluations for this student and TP
	// In a real implementation, we would query evaluations by student_id and evidence_id (linked to assessments by tp_id)
	// For now, we'll use a simplified approach
	evaluations := []domain.Evaluation{}

	// Calculate achievement using domain service
	achievement, err := s.achievementDomainService.CalculateStudentAchievement(
		studentID,
		tpID,
		evaluations,
		tp.SuccessCriteria,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate achievement: %w", err)
	}

	// Set student name
	achievement.StudentName = studentName
	achievement.TPTitle = *tp.Title

	return achievement, nil
}

// CalculateCompetencyProgress calculates progress across competencies for a student
func (s *AchievementService) CalculateCompetencyProgress(
	ctx context.Context,
	studentID string,
	studentName string,
	subjectID string,
	subjectName string,
	phaseID string,
	phaseName string,
) (*domain.CompetencyProgress, error) {
	// Get TPs for this subject and phase
	tpsPtr, err := s.tpRepo.ListTPs(ctx, nil, &subjectID, nil, nil, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get TPs: %w", err)
	}

	// Convert to slice of domain.TP
	tps := make([]domain.TP, len(tpsPtr))
	for i, tp := range tpsPtr {
		tps[i] = *tp
	}

	// Get evaluations for this student
	// In a real implementation, we would query evaluations by student_id
	evaluations := []domain.Evaluation{}

	// Calculate progress using domain service
	progress, err := s.achievementDomainService.CalculateCompetencyProgress(
		studentID,
		studentName,
		subjectID,
		subjectName,
		phaseID,
		phaseName,
		tps,
		evaluations,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate competency progress: %w", err)
	}

	return progress, nil
}

// GenerateAchievementSummary generates a comprehensive achievement summary for a student
func (s *AchievementService) GenerateAchievementSummary(
	ctx context.Context,
	studentID string,
	studentName string,
	classID string,
	className string,
	reportPeriod interface{},
) (*domain.AchievementSummary, error) {
	// Get all achievements for this student
	// In a real implementation, we would query all TPs and calculate achievements
	achievements := []domain.Achievement{}

	// Generate summary using domain service
	summary, err := s.achievementDomainService.GenerateAchievementSummary(
		studentID,
		studentName,
		classID,
		className,
		reportPeriod,
		achievements,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to generate achievement summary: %w", err)
	}

	return summary, nil
}

// GenerateClassAchievement generates achievement summary for an entire class
func (s *AchievementService) GenerateClassAchievement(
	ctx context.Context,
	classID string,
	className string,
	subjectID string,
	subjectName string,
) (*domain.ClassAchievement, error) {
	// Get all student achievements for this class and subject
	// In a real implementation, we would query all students in the class and calculate their achievements
	studentAchievements := []domain.Achievement{}

	// Generate class achievement using domain service
	classAchievement, err := s.achievementDomainService.GenerateClassAchievement(
		classID,
		className,
		subjectID,
		subjectName,
		studentAchievements,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to generate class achievement: %w", err)
	}

	return classAchievement, nil
}

package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/repository"
)

// ReportingService handles business logic for reporting operations
type ReportingService struct {
	reportingRepo      repository.ReportingRepositoryInterface
	achievementService *AchievementService
}

// NewReportingService creates a new reporting service
func NewReportingService(
	reportingRepo repository.ReportingRepositoryInterface,
	achievementService *AchievementService,
) *ReportingService {
	return &ReportingService{
		reportingRepo:      reportingRepo,
		achievementService: achievementService,
	}
}

// CreateNarrativeReport creates a new narrative report
func (s *ReportingService) CreateNarrativeReport(ctx context.Context, req *domain.CreateNarrativeReportRequest, userID string) (*domain.NarrativeReport, error) {
	report := &domain.NarrativeReport{
		ID:               uuid.New().String(),
		StudentID:        req.StudentID,
		ClassID:          req.ClassID,
		UserID:           userID,
		Status:           domain.WorkflowStatusDraft,
		ReportPeriod:     req.ReportPeriod,
		Language:         req.Language,
		Content:          req.Content,
		VersionNo:        1,
		IsCurrentVersion: true,
	}

	if err := s.reportingRepo.CreateNarrativeReport(ctx, report); err != nil {
		return nil, fmt.Errorf("failed to create narrative report: %w", err)
	}

	// Calculate initial achievement data
	studentName := "" // In a real implementation, fetch from user service
	className := ""   // In a real implementation, fetch from class service

	achievementData, err := s.achievementService.GenerateAchievementSummary(
		ctx, report.StudentID, studentName, report.ClassID, className, report.ReportPeriod,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate initial achievement: %w", err)
	}
	report.AchievementData = achievementData

	now := time.Now()
	report.LastAchievementCalculatedAt = &now

	// Update report with achievement data
	if err := s.reportingRepo.UpdateNarrativeReport(ctx, report); err != nil {
		return nil, fmt.Errorf("failed to update report with achievement data: %w", err)
	}

	return report, nil
}

// GetNarrativeReport retrieves a narrative report by ID
func (s *ReportingService) GetNarrativeReport(ctx context.Context, id string) (*domain.NarrativeReport, error) {
	return s.reportingRepo.GetNarrativeReportByID(ctx, id)
}

// ListNarrativeReports retrieves narrative reports with optional filters
func (s *ReportingService) ListNarrativeReports(ctx context.Context, studentID, userID *string, language *domain.ReportLanguage, status *domain.WorkflowStatus, page, pageSize int) ([]*domain.NarrativeReport, int, error) {
	limit := pageSize
	offset := (page - 1) * pageSize
	reports, err := s.reportingRepo.ListNarrativeReports(ctx, studentID, userID, language, status, limit, offset)
	return reports, len(reports), fmt.Errorf("failed to list narrative reports: %w", err)
}

// UpdateNarrativeReport updates a narrative report
func (s *ReportingService) UpdateNarrativeReport(ctx context.Context, id string, req *domain.UpdateNarrativeReportRequest) (*domain.NarrativeReport, error) {
	report, err := s.reportingRepo.GetNarrativeReportByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("narrative report not found: %w", err)
	}

	if req.ReportPeriod != nil {
		report.ReportPeriod = req.ReportPeriod
	}
	if req.Content != nil {
		report.Content = req.Content
	}
	if req.Status != nil {
		report.Status = *req.Status
	}

	if err := s.reportingRepo.UpdateNarrativeReport(ctx, report); err != nil {
		return nil, fmt.Errorf("failed to update narrative report: %w", err)
	}

	return report, nil
}

// RefreshAchievementData refreshes achievement data for a narrative report
func (s *ReportingService) RefreshAchievementData(ctx context.Context, reportID string) (*domain.NarrativeReport, error) {
	report, err := s.reportingRepo.GetNarrativeReportByID(ctx, reportID)
	if err != nil {
		return nil, fmt.Errorf("narrative report not found: %w", err)
	}

	// Calculate achievement data using achievement service
	studentName := "" // In a real implementation, fetch from user service
	className := ""   // In a real implementation, fetch from class service

	achievementData, err := s.achievementService.GenerateAchievementSummary(
		ctx, report.StudentID, studentName, report.ClassID, className, report.ReportPeriod,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate achievement: %w", err)
	}
	report.AchievementData = achievementData

	now := time.Now()
	report.LastAchievementCalculatedAt = &now

	if err := s.reportingRepo.UpdateNarrativeReport(ctx, report); err != nil {
		return nil, fmt.Errorf("failed to update report with achievement data: %w", err)
	}

	return report, nil
}

// DeleteNarrativeReport deletes a narrative report
func (s *ReportingService) DeleteNarrativeReport(ctx context.Context, id string) error {
	return s.reportingRepo.DeleteNarrativeReport(ctx, id)
}

package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/repository"
)

// ReportingService handles business logic for reporting operations
type ReportingService struct {
	reportingRepo *repository.ReportingRepository
}

// NewReportingService creates a new reporting service
func NewReportingService(reportingRepo *repository.ReportingRepository) *ReportingService {
	return &ReportingService{reportingRepo: reportingRepo}
}

// CreateNarrativeReport creates a new narrative report
func (s *ReportingService) CreateNarrativeReport(ctx context.Context, req *domain.CreateNarrativeReportRequest, userID string) (*domain.NarrativeReport, error) {
	report := &domain.NarrativeReport{
		ID:           uuid.New().String(),
		StudentID:    req.StudentID,
		UserID:       userID,
		Status:       domain.WorkflowStatusDraft,
		ReportPeriod: req.ReportPeriod,
		Language:     req.Language,
		Content:      req.Content,
		VersionNo:    1,
		IsCurrentVersion: true,
	}

	if err := s.reportingRepo.CreateNarrativeReport(ctx, report); err != nil {
		return nil, fmt.Errorf("failed to create narrative report: %w", err)
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
	return reports, len(reports), err
}

// UpdateNarrativeReport updates a narrative report
func (s *ReportingService) UpdateNarrativeReport(ctx context.Context, id string, req *domain.UpdateNarrativeReportRequest) (*domain.NarrativeReport, error) {
	report, err := s.reportingRepo.GetNarrativeReportByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("narrative report not found")
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

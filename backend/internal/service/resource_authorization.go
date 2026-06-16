package service

import (
	"context"
	"fmt"

	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/repository"
)

// ResourceAuthorizationService handles resource-level authorization checks
type ResourceAuthorizationService struct {
	userRepo             repository.UserRepositoryInterface
	tpRepo               repository.TPRepositoryInterface
	learningPlanningRepo repository.LearningPlanningRepositoryInterface
	assessmentRepo       repository.AssessmentRepositoryInterface
	reportRepo           repository.ReportingRepositoryInterface
}

// NewResourceAuthorizationService creates a new resource authorization service
func NewResourceAuthorizationService(
	userRepo repository.UserRepositoryInterface,
	tpRepo repository.TPRepositoryInterface,
	learningPlanningRepo repository.LearningPlanningRepositoryInterface,
	assessmentRepo repository.AssessmentRepositoryInterface,
	reportRepo repository.ReportingRepositoryInterface,
) *ResourceAuthorizationService {
	return &ResourceAuthorizationService{
		userRepo:             userRepo,
		tpRepo:               tpRepo,
		learningPlanningRepo: learningPlanningRepo,
		assessmentRepo:       assessmentRepo,
		reportRepo:           reportRepo,
	}
}

// AuthorizeSchoolAccess checks if user has access to a specific school
func (s *ResourceAuthorizationService) AuthorizeSchoolAccess(ctx context.Context, userID, schoolID string) error {
	// Get user to check school_id
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	// Check if user belongs to the school
	if user.SchoolID == nil || *user.SchoolID != schoolID {
		return fmt.Errorf("user does not have access to school %s", schoolID)
	}

	return nil
}

// AuthorizeOwnership checks if user owns a specific resource (based on user_id field)
func (s *ResourceAuthorizationService) AuthorizeOwnership(ctx context.Context, userID, resourceUserID string, roleID string) error {
	// Get user to check role
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	// SYSTEM_ADMIN can access any resource
	if user.RoleID == getSystemAdminRoleID() {
		return nil
	}

	// Check ownership
	if resourceUserID != userID {
		return fmt.Errorf("user does not own this resource")
	}

	return nil
}

// GetUserSchoolID returns the school ID of a user
func (s *ResourceAuthorizationService) GetUserSchoolID(ctx context.Context, userID string) (*string, error) {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return user.SchoolID, nil
}

// CheckResourcePermission checks if user's role has permission for a resource action
func (s *ResourceAuthorizationService) CheckResourcePermission(ctx context.Context, userID string, resource, action string) error {
	// Get user to check role
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	// Map role_id to role name (simplified - in production would query role table)
	roleName := mapRoleIDToName(user.RoleID)

	// Check if role has permission
	if !domain.HasPermission(roleName, resource, action) {
		return fmt.Errorf("user does not have %s permission on resource %s", action, resource)
	}

	return nil
}

// Helper function to get SYSTEM_ADMIN role ID (would be from config or DB in production)
func getSystemAdminRoleID() string {
	return "1" // Simplified for now
}

// Helper function to map role_id to role name (would be from DB in production)
func mapRoleIDToName(roleID string) string {
	// Simplified mapping - in production would query roles table
	roleMap := map[string]string{
		"1": domain.RoleSystemAdmin,
		"2": domain.RoleSchoolAdmin,
		"3": domain.RoleTeacher,
	}
	if roleName, exists := roleMap[roleID]; exists {
		return roleName
	}
	return domain.RoleTeacher // Default to teacher for unknown roles
}

// AuthorizeTPOwnership checks if user owns the specified TP
func (s *ResourceAuthorizationService) AuthorizeTPOwnership(ctx context.Context, userID, tpID string) error {
	// Check if user is system admin
	if s.isSystemAdmin(ctx, userID) {
		return nil
	}

	tp, err := s.tpRepo.GetTPByID(ctx, tpID)
	if err != nil {
		return fmt.Errorf("failed to get TP: %w", err)
	}

	if tp.UserID != userID {
		return fmt.Errorf("user does not own this TP")
	}

	return nil
}

// AuthorizeATPOwnership checks if user owns the specified ATP (via ATPSet)
func (s *ResourceAuthorizationService) AuthorizeATPOwnership(ctx context.Context, userID, atpID string) error {
	// Check if user is system admin
	if s.isSystemAdmin(ctx, userID) {
		return nil
	}

	atp, err := s.learningPlanningRepo.GetATPByID(ctx, atpID)
	if err != nil {
		return fmt.Errorf("failed to get ATP: %w", err)
	}

	// For ATP, ownership is derived from the ATPSet
	// Get the ATPSet to check generated_by
	atpSet, err := s.learningPlanningRepo.GetATPSetByID(ctx, atp.ATPSetID)
	if err != nil {
		return fmt.Errorf("failed to get ATPSet: %w", err)
	}

	if atpSet.GeneratedBy != userID {
		return fmt.Errorf("user does not own this ATP")
	}

	return nil
}

// AuthorizeAssessmentOwnership checks if user owns the specified Assessment
func (s *ResourceAuthorizationService) AuthorizeAssessmentOwnership(ctx context.Context, userID, assessmentID string) error {
	// Check if user is system admin
	if s.isSystemAdmin(ctx, userID) {
		return nil
	}

	assessment, err := s.assessmentRepo.GetAssessmentByID(ctx, assessmentID)
	if err != nil {
		return fmt.Errorf("failed to get Assessment: %w", err)
	}

	if assessment.UserID != userID {
		return fmt.Errorf("user does not own this Assessment")
	}

	return nil
}

// AuthorizeEvidenceOwnership checks if user owns the specified Evidence
func (s *ResourceAuthorizationService) AuthorizeEvidenceOwnership(ctx context.Context, userID, evidenceID string) error {
	// Check if user is system admin
	if s.isSystemAdmin(ctx, userID) {
		return nil
	}

	evidence, err := s.assessmentRepo.GetEvidenceByID(ctx, evidenceID)
	if err != nil {
		return fmt.Errorf("failed to get Evidence: %w", err)
	}

	if evidence.UserID != userID {
		return fmt.Errorf("user does not own this Evidence")
	}

	return nil
}

// AuthorizeEvaluationOwnership checks if user owns the specified Evaluation
func (s *ResourceAuthorizationService) AuthorizeEvaluationOwnership(ctx context.Context, userID, evaluationID string) error {
	// Check if user is system admin
	if s.isSystemAdmin(ctx, userID) {
		return nil
	}

	evaluation, err := s.assessmentRepo.GetEvaluationByID(ctx, evaluationID)
	if err != nil {
		return fmt.Errorf("failed to get Evaluation: %w", err)
	}

	if evaluation.UserID != userID {
		return fmt.Errorf("user does not own this Evaluation")
	}

	return nil
}

// AuthorizeNarrativeReportOwnership checks if user owns the specified Narrative Report
func (s *ResourceAuthorizationService) AuthorizeNarrativeReportOwnership(ctx context.Context, userID, reportID string) error {
	// Check if user is system admin
	if s.isSystemAdmin(ctx, userID) {
		return nil
	}

	report, err := s.reportRepo.GetNarrativeReportByID(ctx, reportID)
	if err != nil {
		return fmt.Errorf("failed to get Narrative Report: %w", err)
	}

	if report.UserID != userID {
		return fmt.Errorf("user does not own this Narrative Report")
	}

	return nil
}

// isSystemAdmin checks if user is a system admin
func (s *ResourceAuthorizationService) isSystemAdmin(ctx context.Context, userID string) bool {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return false
	}
	return user.RoleID == getSystemAdminRoleID()
}

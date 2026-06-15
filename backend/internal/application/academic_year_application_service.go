package application

import (
	"context"
	"fmt"
	"time"

	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/repository"
)

// AcademicYearApplicationService orchestrates Academic Year use cases
type AcademicYearApplicationService struct {
	academicYearRepo *repository.AcademicYearRepository
	userRepo         *repository.UserRepository
	semesterRepo     *repository.SemesterRepository
}

// NewAcademicYearApplicationService creates a new academic year application service
func NewAcademicYearApplicationService(
	academicYearRepo *repository.AcademicYearRepository,
	userRepo *repository.UserRepository,
	semesterRepo *repository.SemesterRepository,
) *AcademicYearApplicationService {
	return &AcademicYearApplicationService{
		academicYearRepo: academicYearRepo,
		userRepo:         userRepo,
		semesterRepo:     semesterRepo,
	}
}

// CreateAcademicYearCommand represents the command to create an academic year
type CreateAcademicYearCommand struct {
	SchoolID    string
	Name        string
	StartDate   time.Time
	EndDate     time.Time
	Description string
	UserID      string
}

// CreateAcademicYearResponse represents the response for creating an academic year
type CreateAcademicYearResponse struct {
	AcademicYearID string
	Status         domain.AcademicYearStatus
}

// CreateAcademicYear creates a new academic year
// Enforces: BR-001 (single active year), BR-002 (non-overlap), BR-003 (lead time)
func (s *AcademicYearApplicationService) CreateAcademicYear(ctx context.Context, cmd *CreateAcademicYearCommand) (*CreateAcademicYearResponse, error) {
	// 1. Authorization: Get user and validate role
	user, err := s.userRepo.GetByID(ctx, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Only School Admin and System Admin can create academic years
	if user.RoleID != domain.RoleSchoolAdmin && user.RoleID != domain.RoleSystemAdmin {
		return nil, fmt.Errorf("user does not have permission to create academic years")
	}

	// 2. School scope: Validate user belongs to the school
	if user.SchoolID == nil || *user.SchoolID != cmd.SchoolID {
		if user.RoleID != domain.RoleSystemAdmin {
			return nil, fmt.Errorf("user does not belong to the specified school")
		}
	}

	// 3. Create domain entity (enforces BR-003: lead time)
	academicYear, err := domain.NewAcademicYear(cmd.SchoolID, cmd.Name, cmd.StartDate, cmd.EndDate, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to create academic year: %w", err)
	}

	// 4. Business rule validation
	if err := academicYear.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	// 5. Check BR-002: Non-overlapping academic years
	overlap, err := s.academicYearRepo.CheckAcademicYearOverlap(ctx, cmd.SchoolID, cmd.StartDate, cmd.EndDate, "")
	if err != nil {
		return nil, fmt.Errorf("failed to check overlap: %w", err)
	}
	if overlap {
		return nil, fmt.Errorf("academic year dates overlap with existing academic year")
	}

	// 6. Persist
	if err := s.academicYearRepo.CreateAcademicYear(ctx, academicYear); err != nil {
		return nil, fmt.Errorf("failed to create academic year: %w", err)
	}

	return &CreateAcademicYearResponse{
		AcademicYearID: academicYear.ID,
		Status:         academicYear.Status,
	}, nil
}

// UpdateAcademicYearCommand represents the command to update an academic year
type UpdateAcademicYearCommand struct {
	AcademicYearID string
	Name           *string
	StartDate      *time.Time
	EndDate        *time.Time
	UserID         string
}

// UpdateAcademicYearResponse represents the response for updating an academic year
type UpdateAcademicYearResponse struct {
	AcademicYearID string
	Status         domain.AcademicYearStatus
}

// UpdateAcademicYear updates an academic year
// Only DRAFT academic years can be modified
func (s *AcademicYearApplicationService) UpdateAcademicYear(ctx context.Context, cmd *UpdateAcademicYearCommand) (*UpdateAcademicYearResponse, error) {
	// 1. Authorization: Get user
	user, err := s.userRepo.GetByID(ctx, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Only School Admin and System Admin can update academic years
	if user.RoleID != domain.RoleSchoolAdmin && user.RoleID != domain.RoleSystemAdmin {
		return nil, fmt.Errorf("user does not have permission to update academic years")
	}

	// 2. Load existing academic year
	academicYear, err := s.academicYearRepo.GetAcademicYearByID(ctx, cmd.AcademicYearID)
	if err != nil {
		return nil, fmt.Errorf("academic year not found: %w", err)
	}

	// 3. School scope: Validate user belongs to same school
	if user.SchoolID != nil && *user.SchoolID != "" && academicYear.SchoolID != *user.SchoolID {
		if user.RoleID != domain.RoleSystemAdmin {
			return nil, fmt.Errorf("cross-school access not allowed")
		}
	}

	// 4. Business rule: Only DRAFT years can be modified
	if !academicYear.CanBeModified() {
		return nil, fmt.Errorf("only DRAFT academic years can be modified")
	}

	// 5. Apply updates
	if cmd.Name != nil {
		academicYear.Name = *cmd.Name
	}
	if cmd.StartDate != nil {
		academicYear.StartDate = *cmd.StartDate
	}
	if cmd.EndDate != nil {
		academicYear.EndDate = *cmd.EndDate
	}
	academicYear.UpdatedAt = time.Now()

	// 6. Business rule validation
	if err := academicYear.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	// 7. Check BR-002: Non-overlapping academic years
	overlap, err := s.academicYearRepo.CheckAcademicYearOverlap(ctx, academicYear.SchoolID, academicYear.StartDate, academicYear.EndDate, academicYear.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to check overlap: %w", err)
	}
	if overlap {
		return nil, fmt.Errorf("academic year dates overlap with existing academic year")
	}

	// 8. Persist
	if err := s.academicYearRepo.UpdateAcademicYear(ctx, academicYear); err != nil {
		return nil, fmt.Errorf("failed to update academic year: %w", err)
	}

	return &UpdateAcademicYearResponse{
		AcademicYearID: academicYear.ID,
		Status:         academicYear.Status,
	}, nil
}

// ActivateAcademicYearCommand represents the command to activate an academic year
type ActivateAcademicYearCommand struct {
	AcademicYearID string
	UserID         string
}

// ActivateAcademicYearResponse represents the response for activating an academic year
type ActivateAcademicYearResponse struct {
	AcademicYearID string
	Status         domain.AcademicYearStatus
}

// ActivateAcademicYear activates an academic year
// Enforces BR-001: Only one active academic year at a time
func (s *AcademicYearApplicationService) ActivateAcademicYear(ctx context.Context, cmd *ActivateAcademicYearCommand) (*ActivateAcademicYearResponse, error) {
	// 1. Authorization: Get user
	user, err := s.userRepo.GetByID(ctx, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Only School Admin and System Admin can activate academic years
	if user.RoleID != domain.RoleSchoolAdmin && user.RoleID != domain.RoleSystemAdmin {
		return nil, fmt.Errorf("user does not have permission to activate academic years")
	}

	// 2. Load existing academic year
	academicYear, err := s.academicYearRepo.GetAcademicYearByID(ctx, cmd.AcademicYearID)
	if err != nil {
		return nil, fmt.Errorf("academic year not found: %w", err)
	}

	// 3. School scope: Validate user belongs to same school
	if user.SchoolID != nil && *user.SchoolID != "" && academicYear.SchoolID != *user.SchoolID {
		if user.RoleID != domain.RoleSystemAdmin {
			return nil, fmt.Errorf("cross-school access not allowed")
		}
	}

	// 4. Business rule: Only DRAFT years can be activated
	if !academicYear.IsDraft() {
		return nil, fmt.Errorf("only DRAFT academic years can be activated")
	}

	// 5. Check if academic year has exactly 2 semesters configured
	// This is a business rule for activation
	semesterCount, err := s.semesterRepo.CountSemestersByAcademicYearID(ctx, academicYear.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to check semesters: %w", err)
	}
	if semesterCount != 2 {
		return nil, fmt.Errorf("academic year must have exactly 2 semesters to be activated")
	}

	// 6. Activate through domain
	if err := academicYear.Activate(); err != nil {
		return nil, fmt.Errorf("failed to activate academic year: %w", err)
	}

	// 7. Persist (this also deactivates any previously active year)
	if err := s.academicYearRepo.ActivateAcademicYear(ctx, academicYear.ID); err != nil {
		return nil, fmt.Errorf("failed to activate academic year: %w", err)
	}

	return &ActivateAcademicYearResponse{
		AcademicYearID: academicYear.ID,
		Status:         academicYear.Status,
	}, nil
}

// ArchiveAcademicYearCommand represents the command to archive an academic year
type ArchiveAcademicYearCommand struct {
	AcademicYearID string
	UserID         string
}

// ArchiveAcademicYearResponse represents the response for archiving an academic year
type ArchiveAcademicYearResponse struct {
	AcademicYearID string
	Status         domain.AcademicYearStatus
}

// ArchiveAcademicYear archives an academic year
func (s *AcademicYearApplicationService) ArchiveAcademicYear(ctx context.Context, cmd *ArchiveAcademicYearCommand) (*ArchiveAcademicYearResponse, error) {
	// 1. Authorization: Get user
	user, err := s.userRepo.GetByID(ctx, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Only School Admin and System Admin can archive academic years
	if user.RoleID != domain.RoleSchoolAdmin && user.RoleID != domain.RoleSystemAdmin {
		return nil, fmt.Errorf("user does not have permission to archive academic years")
	}

	// 2. Load existing academic year
	academicYear, err := s.academicYearRepo.GetAcademicYearByID(ctx, cmd.AcademicYearID)
	if err != nil {
		return nil, fmt.Errorf("academic year not found: %w", err)
	}

	// 3. School scope: Validate user belongs to same school
	if user.SchoolID != nil && *user.SchoolID != "" && academicYear.SchoolID != *user.SchoolID {
		if user.RoleID != domain.RoleSystemAdmin {
			return nil, fmt.Errorf("cross-school access not allowed")
		}
	}

	// 4. Business rule: Only ACTIVE years can be archived
	if !academicYear.IsActive() {
		return nil, fmt.Errorf("only ACTIVE academic years can be archived")
	}

	// 5. Archive through domain
	if err := academicYear.Archive(); err != nil {
		return nil, fmt.Errorf("failed to archive academic year: %w", err)
	}

	// 6. Persist
	if err := s.academicYearRepo.UpdateAcademicYear(ctx, academicYear); err != nil {
		return nil, fmt.Errorf("failed to archive academic year: %w", err)
	}

	return &ArchiveAcademicYearResponse{
		AcademicYearID: academicYear.ID,
		Status:         academicYear.Status,
	}, nil
}

// GetAcademicYearCommand represents the command to get an academic year
type GetAcademicYearCommand struct {
	AcademicYearID string
	UserID         string
}

// GetAcademicYearResponse represents the response for getting an academic year
type GetAcademicYearResponse struct {
	*domain.AcademicYear
	Semesters []*domain.Semester
}

// GetAcademicYear retrieves an academic year with its semesters
func (s *AcademicYearApplicationService) GetAcademicYear(ctx context.Context, cmd *GetAcademicYearCommand) (*GetAcademicYearResponse, error) {
	// 1. Authorization: Get user
	user, err := s.userRepo.GetByID(ctx, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// 2. Load academic year
	academicYear, err := s.academicYearRepo.GetAcademicYearByID(ctx, cmd.AcademicYearID)
	if err != nil {
		return nil, fmt.Errorf("academic year not found: %w", err)
	}

	// 3. School scope: Validate user belongs to same school
	if user.SchoolID != nil && *user.SchoolID != "" && academicYear.SchoolID != *user.SchoolID {
		if user.RoleID != domain.RoleSystemAdmin {
			return nil, fmt.Errorf("cross-school access not allowed")
		}
	}

	// 4. Load semesters
	semesters, err := s.academicYearRepo.GetSemestersByAcademicYearID(ctx, academicYear.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to load semesters: %w", err)
	}

	return &GetAcademicYearResponse{
		AcademicYear: academicYear,
		Semesters:    semesters,
	}, nil
}

// ListAcademicYearsCommand represents the command to list academic years
type ListAcademicYearsCommand struct {
	SchoolID string
	UserID   string
}

// ListAcademicYearsResponse represents the response for listing academic years
type ListAcademicYearsResponse struct {
	AcademicYears []*domain.AcademicYear
}

// ListAcademicYears retrieves all academic years for a school
func (s *AcademicYearApplicationService) ListAcademicYears(ctx context.Context, cmd *ListAcademicYearsCommand) (*ListAcademicYearsResponse, error) {
	// 1. Authorization: Get user
	user, err := s.userRepo.GetByID(ctx, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// 2. School scope: Validate user belongs to school
	if user.SchoolID == nil || *user.SchoolID != cmd.SchoolID {
		if user.RoleID != domain.RoleSystemAdmin {
			return nil, fmt.Errorf("user does not belong to the specified school")
		}
	}

	// 3. Load academic years
	academicYears, err := s.academicYearRepo.GetAcademicYearsBySchoolID(ctx, cmd.SchoolID)
	if err != nil {
		return nil, fmt.Errorf("failed to load academic years: %w", err)
	}

	return &ListAcademicYearsResponse{
		AcademicYears: academicYears,
	}, nil
}

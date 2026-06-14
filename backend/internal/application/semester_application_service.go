package application

import (
	"context"
	"fmt"
	"time"

	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/repository"
)

// SemesterApplicationService orchestrates Semester use cases
type SemesterApplicationService struct {
	semesterRepo     *repository.SemesterRepository
	academicYearRepo *repository.AcademicYearRepository
	userRepo         *repository.UserRepository
}

// NewSemesterApplicationService creates a new semester application service
func NewSemesterApplicationService(
	semesterRepo *repository.SemesterRepository,
	academicYearRepo *repository.AcademicYearRepository,
	userRepo *repository.UserRepository,
) *SemesterApplicationService {
	return &SemesterApplicationService{
		semesterRepo:     semesterRepo,
		academicYearRepo: academicYearRepo,
		userRepo:         userRepo,
	}
}

// CreateSemesterCommand represents the command to create a semester
type CreateSemesterCommand struct {
	AcademicYearID string
	Type           domain.SemesterType
	Name           string
	StartDate      time.Time
	EndDate        time.Time
	SequenceNumber int
	UserID         string
}

// CreateSemesterResponse represents the response for creating a semester
type CreateSemesterResponse struct {
	SemesterID string
	Status     domain.SemesterStatus
}

// CreateSemester creates a new semester
func (s *SemesterApplicationService) CreateSemester(ctx context.Context, cmd *CreateSemesterCommand) (*CreateSemesterResponse, error) {
	// 1. Authorization: Get user
	user, err := s.userRepo.GetByID(ctx, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// School Admin, Curriculum Admin, and System Admin can create semesters
	if user.RoleID != domain.RoleSchoolAdmin && user.RoleID != domain.RoleCurriculumAdmin && user.RoleID != domain.RoleSystemAdmin {
		return nil, fmt.Errorf("user does not have permission to create semesters")
	}

	// 2. Load academic year to validate
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

	// 4. Business rule: Only DRAFT academic years can have semesters added
	if !academicYear.IsDraft() {
		return nil, fmt.Errorf("semesters can only be added to DRAFT academic years")
	}

	// 5. Create domain entity
	semester, err := domain.NewSemester(cmd.AcademicYearID, cmd.Name, cmd.Type, cmd.StartDate, cmd.EndDate, cmd.SequenceNumber, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to create semester: %w", err)
	}

	// 6. Business rule validation
	if err := semester.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	// 7. Check for overlap with existing semesters
	overlap, err := s.semesterRepo.CheckSemesterOverlap(ctx, cmd.AcademicYearID, cmd.StartDate, cmd.EndDate, "")
	if err != nil {
		return nil, fmt.Errorf("failed to check overlap: %w", err)
	}
	if overlap {
		return nil, fmt.Errorf("semester dates overlap with existing semester")
	}

	// 8. Check sequence number uniqueness
	seqExists, err := s.semesterRepo.CheckSequenceNumberExists(ctx, cmd.AcademicYearID, cmd.SequenceNumber, "")
	if err != nil {
		return nil, fmt.Errorf("failed to check sequence number: %w", err)
	}
	if seqExists {
		return nil, fmt.Errorf("sequence number %d already exists for this academic year", cmd.SequenceNumber)
	}

	// 9. Persist
	if err := s.semesterRepo.CreateSemester(ctx, semester); err != nil {
		return nil, fmt.Errorf("failed to create semester: %w", err)
	}

	return &CreateSemesterResponse{
		SemesterID: semester.ID,
		Status:     semester.Status,
	}, nil
}

// UpdateSemesterCommand represents the command to update a semester
type UpdateSemesterCommand struct {
	SemesterID string
	Name       *string
	StartDate  *time.Time
	EndDate    *time.Time
	Status     *domain.SemesterStatus
	UserID     string
}

// UpdateSemesterResponse represents the response for updating a semester
type UpdateSemesterResponse struct {
	SemesterID string
	Status     domain.SemesterStatus
}

// UpdateSemester updates a semester
func (s *SemesterApplicationService) UpdateSemester(ctx context.Context, cmd *UpdateSemesterCommand) (*UpdateSemesterResponse, error) {
	// 1. Authorization: Get user
	user, err := s.userRepo.GetByID(ctx, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// School Admin, Curriculum Admin, and System Admin can update semesters
	if user.RoleID != domain.RoleSchoolAdmin && user.RoleID != domain.RoleCurriculumAdmin && user.RoleID != domain.RoleSystemAdmin {
		return nil, fmt.Errorf("user does not have permission to update semesters")
	}

	// 2. Load existing semester
	semester, err := s.semesterRepo.GetSemesterByID(ctx, cmd.SemesterID)
	if err != nil {
		return nil, fmt.Errorf("semester not found: %w", err)
	}

	// 3. Load academic year
	academicYear, err := s.academicYearRepo.GetAcademicYearByID(ctx, semester.AcademicYearID)
	if err != nil {
		return nil, fmt.Errorf("academic year not found: %w", err)
	}

	// 4. School scope: Validate user belongs to same school
	if user.SchoolID != nil && *user.SchoolID != "" && academicYear.SchoolID != *user.SchoolID {
		if user.RoleID != domain.RoleSystemAdmin {
			return nil, fmt.Errorf("cross-school access not allowed")
		}
	}

	// 5. Business rule: Only DRAFT academic years can have semesters modified
	if !academicYear.IsDraft() {
		return nil, fmt.Errorf("semesters can only be modified in DRAFT academic years")
	}

	// 6. Apply updates
	if cmd.Name != nil {
		semester.Name = *cmd.Name
	}
	if cmd.StartDate != nil {
		semester.StartDate = *cmd.StartDate
	}
	if cmd.EndDate != nil {
		semester.EndDate = *cmd.EndDate
	}
	if cmd.Status != nil {
		semester.Status = *cmd.Status
	}
	semester.UpdatedAt = time.Now()

	// 7. Business rule validation
	if err := semester.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	// 8. Check for overlap with existing semesters (if dates changed)
	if cmd.StartDate != nil || cmd.EndDate != nil {
		checkStartDate := semester.StartDate
		checkEndDate := semester.EndDate
		overlap, err := s.semesterRepo.CheckSemesterOverlap(ctx, semester.AcademicYearID, checkStartDate, checkEndDate, semester.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to check overlap: %w", err)
		}
		if overlap {
			return nil, fmt.Errorf("semester dates overlap with existing semester")
		}
	}

	// 9. Persist
	if err := s.semesterRepo.UpdateSemester(ctx, semester); err != nil {
		return nil, fmt.Errorf("failed to update semester: %w", err)
	}

	return &UpdateSemesterResponse{
		SemesterID: semester.ID,
		Status:     semester.Status,
	}, nil
}

// DeleteSemesterCommand represents the command to delete a semester
type DeleteSemesterCommand struct {
	SemesterID string
	UserID     string
}

// DeleteSemesterResponse represents the response for deleting a semester
type DeleteSemesterResponse struct {
	Success bool
}

// DeleteSemester deletes a semester
func (s *SemesterApplicationService) DeleteSemester(ctx context.Context, cmd *DeleteSemesterCommand) (*DeleteSemesterResponse, error) {
	// 1. Authorization: Get user
	user, err := s.userRepo.GetByID(ctx, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// School Admin, Curriculum Admin, and System Admin can delete semesters
	if user.RoleID != domain.RoleSchoolAdmin && user.RoleID != domain.RoleCurriculumAdmin && user.RoleID != domain.RoleSystemAdmin {
		return nil, fmt.Errorf("user does not have permission to delete semesters")
	}

	// 2. Load existing semester
	semester, err := s.semesterRepo.GetSemesterByID(ctx, cmd.SemesterID)
	if err != nil {
		return nil, fmt.Errorf("semester not found: %w", err)
	}

	// 3. Load academic year
	academicYear, err := s.academicYearRepo.GetAcademicYearByID(ctx, semester.AcademicYearID)
	if err != nil {
		return nil, fmt.Errorf("academic year not found: %w", err)
	}

	// 4. School scope: Validate user belongs to same school
	if user.SchoolID != nil && *user.SchoolID != "" && academicYear.SchoolID != *user.SchoolID {
		if user.RoleID != domain.RoleSystemAdmin {
			return nil, fmt.Errorf("cross-school access not allowed")
		}
	}

	// 5. Business rule: Only DRAFT academic years can have semesters deleted
	if !academicYear.IsDraft() {
		return nil, fmt.Errorf("semesters can only be deleted from DRAFT academic years")
	}

	// 6. Persist
	if err := s.semesterRepo.DeleteSemester(ctx, cmd.SemesterID); err != nil {
		return nil, fmt.Errorf("failed to delete semester: %w", err)
	}

	return &DeleteSemesterResponse{
		Success: true,
	}, nil
}

// GetSemesterCommand represents the command to get a semester
type GetSemesterCommand struct {
	SemesterID string
	UserID     string
}

// GetSemesterResponse represents the response for getting a semester
type GetSemesterResponse struct {
	*domain.Semester
}

// GetSemester retrieves a semester
func (s *SemesterApplicationService) GetSemester(ctx context.Context, cmd *GetSemesterCommand) (*GetSemesterResponse, error) {
	// 1. Authorization: Get user
	user, err := s.userRepo.GetByID(ctx, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// 2. Load semester
	semester, err := s.semesterRepo.GetSemesterByID(ctx, cmd.SemesterID)
	if err != nil {
		return nil, fmt.Errorf("semester not found: %w", err)
	}

	// 3. Load academic year
	academicYear, err := s.academicYearRepo.GetAcademicYearByID(ctx, semester.AcademicYearID)
	if err != nil {
		return nil, fmt.Errorf("academic year not found: %w", err)
	}

	// 4. School scope: Validate user belongs to same school
	if user.SchoolID != nil && *user.SchoolID != "" && academicYear.SchoolID != *user.SchoolID {
		if user.RoleID != domain.RoleSystemAdmin {
			return nil, fmt.Errorf("cross-school access not allowed")
		}
	}

	return &GetSemesterResponse{
		Semester: semester,
	}, nil
}

// ListSemestersCommand represents the command to list semesters
type ListSemestersCommand struct {
	AcademicYearID string
	UserID         string
}

// ListSemestersResponse represents the response for listing semesters
type ListSemestersResponse struct {
	Semesters []*domain.Semester
}

// ListSemesters retrieves all semesters for an academic year
func (s *SemesterApplicationService) ListSemesters(ctx context.Context, cmd *ListSemestersCommand) (*ListSemestersResponse, error) {
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
	semesters, err := s.semesterRepo.GetSemestersByAcademicYearID(ctx, cmd.AcademicYearID)
	if err != nil {
		return nil, fmt.Errorf("failed to load semesters: %w", err)
	}

	return &ListSemestersResponse{
		Semesters: semesters,
	}, nil
}

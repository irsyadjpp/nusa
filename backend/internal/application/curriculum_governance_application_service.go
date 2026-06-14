package application

import (
	"context"
	"fmt"
	"time"

	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/repository"
)

// CurriculumGovernanceApplicationService orchestrates curriculum governance use cases
// This includes Subject Category, Graduate Profile Dimension, and CP Alignment operations
type CurriculumGovernanceApplicationService struct {
	subjectCategoryRepo *repository.SubjectCategoryRepository
	graduateProfileRepo *repository.GraduateProfileDimensionRepository
	cpAlignmentRepo     *repository.CPAlignmentRepository
	curriculumRepo      *repository.CurriculumRepository
	systemConfigRepo    *repository.SystemConfigurationRepository
	userRepo            *repository.UserRepository
}

// NewCurriculumGovernanceApplicationService creates a new curriculum governance application service
func NewCurriculumGovernanceApplicationService(
	subjectCategoryRepo *repository.SubjectCategoryRepository,
	graduateProfileRepo *repository.GraduateProfileDimensionRepository,
	cpAlignmentRepo *repository.CPAlignmentRepository,
	curriculumRepo *repository.CurriculumRepository,
	systemConfigRepo *repository.SystemConfigurationRepository,
	userRepo *repository.UserRepository,
) *CurriculumGovernanceApplicationService {
	return &CurriculumGovernanceApplicationService{
		subjectCategoryRepo: subjectCategoryRepo,
		graduateProfileRepo: graduateProfileRepo,
		cpAlignmentRepo:     cpAlignmentRepo,
		curriculumRepo:      curriculumRepo,
		systemConfigRepo:    systemConfigRepo,
		userRepo:            userRepo,
	}
}

// ==================== Subject Category Operations ====================

// CreateSubjectCategoryCommand represents the command to create a subject category
type CreateSubjectCategoryCommand struct {
	Code        string
	Name        string
	Description *string
	IsMandatory bool
	UserID      string
}

// CreateSubjectCategoryResponse represents the response for creating a subject category
type CreateSubjectCategoryResponse struct {
	SubjectCategoryID string
}

// CreateSubjectCategory creates a new subject category
func (s *CurriculumGovernanceApplicationService) CreateSubjectCategory(ctx context.Context, cmd *CreateSubjectCategoryCommand) (*CreateSubjectCategoryResponse, error) {
	// 1. Authorization: Get user
	user, err := s.userRepo.GetByID(ctx, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Only Curriculum Admin and System Admin can create subject categories
	if user.RoleID != domain.RoleCurriculumAdmin && user.RoleID != domain.RoleSystemAdmin {
		return nil, fmt.Errorf("user does not have permission to create subject categories")
	}

	// 2. Check code uniqueness
	codeExists, err := s.subjectCategoryRepo.CheckCodeExists(ctx, cmd.Code, "")
	if err != nil {
		return nil, fmt.Errorf("failed to check code uniqueness: %w", err)
	}
	if codeExists {
		return nil, fmt.Errorf("subject category code already exists")
	}

	// 3. Create domain entity
	subjectCategory, err := domain.NewSubjectCategory(cmd.Code, cmd.Name, cmd.IsMandatory, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to create subject category: %w", err)
	}

	// 4. Apply optional fields
	if cmd.Description != nil {
		subjectCategory.Description = cmd.Description
	}

	// 5. Business rule validation
	if err := subjectCategory.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	// 6. Persist
	if err := s.subjectCategoryRepo.CreateSubjectCategory(ctx, subjectCategory); err != nil {
		return nil, fmt.Errorf("failed to create subject category: %w", err)
	}

	return &CreateSubjectCategoryResponse{
		SubjectCategoryID: subjectCategory.ID,
	}, nil
}

// UpdateSubjectCategoryCommand represents the command to update a subject category
type UpdateSubjectCategoryCommand struct {
	SubjectCategoryID string
	Name              *string
	Description       *string
	IsMandatory       *bool
	IsActive          *bool
	UserID            string
}

// UpdateSubjectCategoryResponse represents the response for updating a subject category
type UpdateSubjectCategoryResponse struct {
	Success bool
}

// UpdateSubjectCategory updates a subject category
func (s *CurriculumGovernanceApplicationService) UpdateSubjectCategory(ctx context.Context, cmd *UpdateSubjectCategoryCommand) (*UpdateSubjectCategoryResponse, error) {
	// 1. Authorization: Get user
	user, err := s.userRepo.GetByID(ctx, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Only Curriculum Admin and System Admin can update subject categories
	if user.RoleID != domain.RoleCurriculumAdmin && user.RoleID != domain.RoleSystemAdmin {
		return nil, fmt.Errorf("user does not have permission to update subject categories")
	}

	// 2. Load existing subject category
	subjectCategory, err := s.subjectCategoryRepo.GetSubjectCategoryByID(ctx, cmd.SubjectCategoryID)
	if err != nil {
		return nil, fmt.Errorf("subject category not found: %w", err)
	}

	// 3. Apply updates
	updatedBy := cmd.UserID
	if cmd.Name != nil {
		subjectCategory.Name = *cmd.Name
	}
	if cmd.Description != nil {
		subjectCategory.Description = cmd.Description
	}
	if cmd.IsMandatory != nil {
		subjectCategory.IsMandatory = *cmd.IsMandatory
	}
	if cmd.IsActive != nil {
		subjectCategory.IsActive = *cmd.IsActive
	}
	subjectCategory.UpdatedBy = &updatedBy
	subjectCategory.UpdatedAt = time.Now()

	// 4. Business rule validation
	if err := subjectCategory.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	// 5. Persist
	if err := s.subjectCategoryRepo.UpdateSubjectCategory(ctx, subjectCategory); err != nil {
		return nil, fmt.Errorf("failed to update subject category: %w", err)
	}

	return &UpdateSubjectCategoryResponse{
		Success: true,
	}, nil
}

// DeleteSubjectCategoryCommand represents the command to delete a subject category
type DeleteSubjectCategoryCommand struct {
	SubjectCategoryID string
	UserID            string
}

// DeleteSubjectCategoryResponse represents the response for deleting a subject category
type DeleteSubjectCategoryResponse struct {
	Success bool
}

// DeleteSubjectCategory deletes a subject category
func (s *CurriculumGovernanceApplicationService) DeleteSubjectCategory(ctx context.Context, cmd *DeleteSubjectCategoryCommand) (*DeleteSubjectCategoryResponse, error) {
	// 1. Authorization: Get user
	user, err := s.userRepo.GetByID(ctx, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Only Curriculum Admin and System Admin can delete subject categories
	if user.RoleID != domain.RoleCurriculumAdmin && user.RoleID != domain.RoleSystemAdmin {
		return nil, fmt.Errorf("user does not have permission to delete subject categories")
	}

	// 2. Persist
	if err := s.subjectCategoryRepo.DeleteSubjectCategory(ctx, cmd.SubjectCategoryID); err != nil {
		return nil, fmt.Errorf("failed to delete subject category: %w", err)
	}

	return &DeleteSubjectCategoryResponse{
		Success: true,
	}, nil
}

// ListSubjectCategoriesCommand represents the command to list subject categories
type ListSubjectCategoriesCommand struct {
	ActiveOnly bool
	UserID     string
}

// ListSubjectCategoriesResponse represents the response for listing subject categories
type ListSubjectCategoriesResponse struct {
	SubjectCategories []*domain.SubjectCategory
}

// ListSubjectCategories retrieves all subject categories
func (s *CurriculumGovernanceApplicationService) ListSubjectCategories(ctx context.Context, cmd *ListSubjectCategoriesCommand) (*ListSubjectCategoriesResponse, error) {
	// 1. Authorization: Get user
	user, err := s.userRepo.GetByID(ctx, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Teachers can read, Curriculum Admin and System Admin can read
	if user.RoleID != domain.RoleTeacher && user.RoleID != domain.RoleCurriculumAdmin && user.RoleID != domain.RoleSystemAdmin {
		return nil, fmt.Errorf("user does not have permission to view subject categories")
	}

	// 2. Load subject categories
	var categories []*domain.SubjectCategory
	if cmd.ActiveOnly {
		categories, err = s.subjectCategoryRepo.GetActiveSubjectCategories(ctx)
	} else {
		categories, err = s.subjectCategoryRepo.GetAllSubjectCategories(ctx)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to load subject categories: %w", err)
	}

	return &ListSubjectCategoriesResponse{
		SubjectCategories: categories,
	}, nil
}

// ==================== Graduate Profile Dimension Operations ====================

// CreateGraduateProfileDimensionCommand represents the command to create a graduate profile dimension
type CreateGraduateProfileDimensionCommand struct {
	Code           string
	Name           string
	Description    *string
	SequenceNumber int
	UserID         string
}

// CreateGraduateProfileDimensionResponse represents the response for creating a graduate profile dimension
type CreateGraduateProfileDimensionResponse struct {
	GraduateProfileDimensionID string
}

// CreateGraduateProfileDimension creates a new graduate profile dimension
func (s *CurriculumGovernanceApplicationService) CreateGraduateProfileDimension(ctx context.Context, cmd *CreateGraduateProfileDimensionCommand) (*CreateGraduateProfileDimensionResponse, error) {
	// 1. Authorization: Get user
	user, err := s.userRepo.GetByID(ctx, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Only Curriculum Admin and System Admin can create graduate profile dimensions
	if user.RoleID != domain.RoleCurriculumAdmin && user.RoleID != domain.RoleSystemAdmin {
		return nil, fmt.Errorf("user does not have permission to create graduate profile dimensions")
	}

	// 2. Check code uniqueness
	codeExists, err := s.graduateProfileRepo.CheckCodeExists(ctx, cmd.Code, "")
	if err != nil {
		return nil, fmt.Errorf("failed to check code uniqueness: %w", err)
	}
	if codeExists {
		return nil, fmt.Errorf("graduate profile dimension code already exists")
	}

	// 3. Check sequence number uniqueness
	seqExists, err := s.graduateProfileRepo.CheckSequenceNumberExists(ctx, cmd.SequenceNumber, "")
	if err != nil {
		return nil, fmt.Errorf("failed to check sequence number: %w", err)
	}
	if seqExists {
		return nil, fmt.Errorf("sequence number %d already exists", cmd.SequenceNumber)
	}

	// 4. Create domain entity
	graduateProfile, err := domain.NewGraduateProfileDimension(cmd.Code, cmd.Name, cmd.SequenceNumber, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to create graduate profile dimension: %w", err)
	}

	// 5. Apply optional fields
	if cmd.Description != nil {
		graduateProfile.Description = cmd.Description
	}

	// 6. Business rule validation
	if err := graduateProfile.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	// 7. Persist
	if err := s.graduateProfileRepo.CreateGraduateProfileDimension(ctx, graduateProfile); err != nil {
		return nil, fmt.Errorf("failed to create graduate profile dimension: %w", err)
	}

	return &CreateGraduateProfileDimensionResponse{
		GraduateProfileDimensionID: graduateProfile.ID,
	}, nil
}

// UpdateGraduateProfileDimensionCommand represents the command to update a graduate profile dimension
type UpdateGraduateProfileDimensionCommand struct {
	GraduateProfileDimensionID string
	Name                       *string
	Description                *string
	SequenceNumber             *int
	IsActive                   *bool
	UserID                     string
}

// UpdateGraduateProfileDimensionResponse represents the response for updating a graduate profile dimension
type UpdateGraduateProfileDimensionResponse struct {
	Success bool
}

// UpdateGraduateProfileDimension updates a graduate profile dimension
func (s *CurriculumGovernanceApplicationService) UpdateGraduateProfileDimension(ctx context.Context, cmd *UpdateGraduateProfileDimensionCommand) (*UpdateGraduateProfileDimensionResponse, error) {
	// 1. Authorization: Get user
	user, err := s.userRepo.GetByID(ctx, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Only Curriculum Admin and System Admin can update graduate profile dimensions
	if user.RoleID != domain.RoleCurriculumAdmin && user.RoleID != domain.RoleSystemAdmin {
		return nil, fmt.Errorf("user does not have permission to update graduate profile dimensions")
	}

	// 2. Load existing graduate profile dimension
	graduateProfile, err := s.graduateProfileRepo.GetGraduateProfileDimensionByID(ctx, cmd.GraduateProfileDimensionID)
	if err != nil {
		return nil, fmt.Errorf("graduate profile dimension not found: %w", err)
	}

	// 3. Check sequence number uniqueness (if changing)
	if cmd.SequenceNumber != nil && *cmd.SequenceNumber != graduateProfile.SequenceNumber {
		seqExists, err := s.graduateProfileRepo.CheckSequenceNumberExists(ctx, *cmd.SequenceNumber, graduateProfile.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to check sequence number: %w", err)
		}
		if seqExists {
			return nil, fmt.Errorf("sequence number %d already exists", *cmd.SequenceNumber)
		}
	}

	// 4. Apply updates
	updatedBy := cmd.UserID
	if cmd.Name != nil {
		graduateProfile.Name = *cmd.Name
	}
	if cmd.Description != nil {
		graduateProfile.Description = cmd.Description
	}
	if cmd.SequenceNumber != nil {
		graduateProfile.SequenceNumber = *cmd.SequenceNumber
	}
	if cmd.IsActive != nil {
		graduateProfile.IsActive = *cmd.IsActive
	}
	graduateProfile.UpdatedBy = &updatedBy
	graduateProfile.UpdatedAt = time.Now()

	// 5. Business rule validation
	if err := graduateProfile.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	// 6. Persist
	if err := s.graduateProfileRepo.UpdateGraduateProfileDimension(ctx, graduateProfile); err != nil {
		return nil, fmt.Errorf("failed to update graduate profile dimension: %w", err)
	}

	return &UpdateGraduateProfileDimensionResponse{
		Success: true,
	}, nil
}

// DeleteGraduateProfileDimensionCommand represents the command to delete a graduate profile dimension
type DeleteGraduateProfileDimensionCommand struct {
	GraduateProfileDimensionID string
	UserID                     string
}

// DeleteGraduateProfileDimensionResponse represents the response for deleting a graduate profile dimension
type DeleteGraduateProfileDimensionResponse struct {
	Success bool
}

// DeleteGraduateProfileDimension deletes a graduate profile dimension
func (s *CurriculumGovernanceApplicationService) DeleteGraduateProfileDimension(ctx context.Context, cmd *DeleteGraduateProfileDimensionCommand) (*DeleteGraduateProfileDimensionResponse, error) {
	// 1. Authorization: Get user
	user, err := s.userRepo.GetByID(ctx, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Only Curriculum Admin and System Admin can delete graduate profile dimensions
	if user.RoleID != domain.RoleCurriculumAdmin && user.RoleID != domain.RoleSystemAdmin {
		return nil, fmt.Errorf("user does not have permission to delete graduate profile dimensions")
	}

	// 2. Persist
	if err := s.graduateProfileRepo.DeleteGraduateProfileDimension(ctx, cmd.GraduateProfileDimensionID); err != nil {
		return nil, fmt.Errorf("failed to delete graduate profile dimension: %w", err)
	}

	return &DeleteGraduateProfileDimensionResponse{
		Success: true,
	}, nil
}

// ListGraduateProfileDimensionsCommand represents the command to list graduate profile dimensions
type ListGraduateProfileDimensionsCommand struct {
	ActiveOnly bool
	UserID     string
}

// ListGraduateProfileDimensionsResponse represents the response for listing graduate profile dimensions
type ListGraduateProfileDimensionsResponse struct {
	GraduateProfileDimensions []*domain.GraduateProfileDimension
}

// ListGraduateProfileDimensions retrieves all graduate profile dimensions
func (s *CurriculumGovernanceApplicationService) ListGraduateProfileDimensions(ctx context.Context, cmd *ListGraduateProfileDimensionsCommand) (*ListGraduateProfileDimensionsResponse, error) {
	// 1. Authorization: Get user
	user, err := s.userRepo.GetByID(ctx, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Teachers can read, Curriculum Admin and System Admin can read
	if user.RoleID != domain.RoleTeacher && user.RoleID != domain.RoleCurriculumAdmin && user.RoleID != domain.RoleSystemAdmin {
		return nil, fmt.Errorf("user does not have permission to view graduate profile dimensions")
	}

	// 2. Load graduate profile dimensions
	var dimensions []*domain.GraduateProfileDimension
	if cmd.ActiveOnly {
		dimensions, err = s.graduateProfileRepo.GetActiveGraduateProfileDimensions(ctx)
	} else {
		dimensions, err = s.graduateProfileRepo.GetAllGraduateProfileDimensions(ctx)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to load graduate profile dimensions: %w", err)
	}

	return &ListGraduateProfileDimensionsResponse{
		GraduateProfileDimensions: dimensions,
	}, nil
}

// ==================== CP Alignment Operations ====================

// CreateCPAlignmentCommand represents the command to create a CP alignment
type CreateCPAlignmentCommand struct {
	CurriculumSubjectID        string
	GraduateProfileDimensionID string
	AlignmentDescription       *string
	UserID                     string
}

// CreateCPAlignmentResponse represents the response for creating a CP alignment
type CreateCPAlignmentResponse struct {
	CPAlignmentID string
}

// CreateCPAlignment creates a new CP alignment
func (s *CurriculumGovernanceApplicationService) CreateCPAlignment(ctx context.Context, cmd *CreateCPAlignmentCommand) (*CreateCPAlignmentResponse, error) {
	// 1. Authorization: Get user
	user, err := s.userRepo.GetByID(ctx, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Only Curriculum Admin and System Admin can create CP alignments
	if user.RoleID != domain.RoleCurriculumAdmin && user.RoleID != domain.RoleSystemAdmin {
		return nil, fmt.Errorf("user does not have permission to create CP alignments")
	}

	// 2. Check if alignment already exists
	alignmentExists, err := s.cpAlignmentRepo.CheckAlignmentExists(ctx, cmd.CurriculumSubjectID, cmd.GraduateProfileDimensionID, "")
	if err != nil {
		return nil, fmt.Errorf("failed to check alignment existence: %w", err)
	}
	if alignmentExists {
		return nil, fmt.Errorf("CP alignment already exists for this combination")
	}

	// 3. Create domain entity
	cpAlignment, err := domain.NewCPAlignment(cmd.CurriculumSubjectID, cmd.GraduateProfileDimensionID, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to create CP alignment: %w", err)
	}

	// 4. Apply optional fields
	if cmd.AlignmentDescription != nil {
		cpAlignment.AlignmentDescription = cmd.AlignmentDescription
	}

	// 5. Business rule validation
	if err := cpAlignment.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	// 6. Persist
	if err := s.cpAlignmentRepo.CreateCPAlignment(ctx, cpAlignment); err != nil {
		return nil, fmt.Errorf("failed to create CP alignment: %w", err)
	}

	return &CreateCPAlignmentResponse{
		CPAlignmentID: cpAlignment.ID,
	}, nil
}

// CreateCPAlignmentBulkCommand represents the command to create multiple CP alignments
type CreateCPAlignmentBulkCommand struct {
	CurriculumSubjectID  string
	AlignmentIDs         []string
	AlignmentDescription *string
	UserID               string
}

// CreateCPAlignmentBulkResponse represents the response for creating multiple CP alignments
type CreateCPAlignmentBulkResponse struct {
	CPAlignmentIDs []string
}

// CreateCPAlignmentBulk creates multiple CP alignments for a curriculum subject
func (s *CurriculumGovernanceApplicationService) CreateCPAlignmentBulk(ctx context.Context, cmd *CreateCPAlignmentBulkCommand) (*CreateCPAlignmentBulkResponse, error) {
	// 1. Authorization: Get user
	user, err := s.userRepo.GetByID(ctx, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Only Curriculum Admin and System Admin can create CP alignments
	if user.RoleID != domain.RoleCurriculumAdmin && user.RoleID != domain.RoleSystemAdmin {
		return nil, fmt.Errorf("user does not have permission to create CP alignments")
	}

	// 2. Delete existing alignments for this curriculum subject
	if err := s.cpAlignmentRepo.DeleteAlignmentsByCurriculumSubjectID(ctx, cmd.CurriculumSubjectID); err != nil {
		return nil, fmt.Errorf("failed to delete existing alignments: %w", err)
	}

	// 3. Create new alignments
	var alignments []*domain.CPAlignment
	var alignmentIDs []string

	for _, dimID := range cmd.AlignmentIDs {
		cpAlignment, err := domain.NewCPAlignment(cmd.CurriculumSubjectID, dimID, cmd.UserID)
		if err != nil {
			return nil, fmt.Errorf("failed to create CP alignment: %w", err)
		}

		if cmd.AlignmentDescription != nil {
			cpAlignment.AlignmentDescription = cmd.AlignmentDescription
		}

		alignments = append(alignments, cpAlignment)
		alignmentIDs = append(alignmentIDs, cpAlignment.ID)
	}

	// 4. Bulk persist
	if err := s.cpAlignmentRepo.BulkCreateCPAlignments(ctx, alignments); err != nil {
		return nil, fmt.Errorf("failed to create CP alignments: %w", err)
	}

	return &CreateCPAlignmentBulkResponse{
		CPAlignmentIDs: alignmentIDs,
	}, nil
}

// UpdateCPAlignmentCommand represents the command to update a CP alignment
type UpdateCPAlignmentCommand struct {
	CPAlignmentID        string
	AlignmentDescription *string
	IsActive             *bool
	UserID               string
}

// UpdateCPAlignmentResponse represents the response for updating a CP alignment
type UpdateCPAlignmentResponse struct {
	Success bool
}

// UpdateCPAlignment updates a CP alignment
func (s *CurriculumGovernanceApplicationService) UpdateCPAlignment(ctx context.Context, cmd *UpdateCPAlignmentCommand) (*UpdateCPAlignmentResponse, error) {
	// 1. Authorization: Get user
	user, err := s.userRepo.GetByID(ctx, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Only Curriculum Admin and System Admin can update CP alignments
	if user.RoleID != domain.RoleCurriculumAdmin && user.RoleID != domain.RoleSystemAdmin {
		return nil, fmt.Errorf("user does not have permission to update CP alignments")
	}

	// 2. Load existing CP alignment
	cpAlignment, err := s.cpAlignmentRepo.GetCPAlignmentByID(ctx, cmd.CPAlignmentID)
	if err != nil {
		return nil, fmt.Errorf("CP alignment not found: %w", err)
	}

	// 3. Apply updates
	updatedBy := cmd.UserID
	if cmd.AlignmentDescription != nil {
		cpAlignment.AlignmentDescription = cmd.AlignmentDescription
	}
	if cmd.IsActive != nil {
		cpAlignment.IsActive = *cmd.IsActive
	}
	cpAlignment.UpdatedBy = &updatedBy
	cpAlignment.UpdatedAt = time.Now()

	// 4. Business rule validation
	if err := cpAlignment.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	// 5. Persist
	if err := s.cpAlignmentRepo.UpdateCPAlignment(ctx, cpAlignment); err != nil {
		return nil, fmt.Errorf("failed to update CP alignment: %w", err)
	}

	return &UpdateCPAlignmentResponse{
		Success: true,
	}, nil
}

// DeleteCPAlignmentCommand represents the command to delete a CP alignment
type DeleteCPAlignmentCommand struct {
	CPAlignmentID string
	UserID        string
}

// DeleteCPAlignmentResponse represents the response for deleting a CP alignment
type DeleteCPAlignmentResponse struct {
	Success bool
}

// DeleteCPAlignment deletes a CP alignment
func (s *CurriculumGovernanceApplicationService) DeleteCPAlignment(ctx context.Context, cmd *DeleteCPAlignmentCommand) (*DeleteCPAlignmentResponse, error) {
	// 1. Authorization: Get user
	user, err := s.userRepo.GetByID(ctx, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Only Curriculum Admin and System Admin can delete CP alignments
	if user.RoleID != domain.RoleCurriculumAdmin && user.RoleID != domain.RoleSystemAdmin {
		return nil, fmt.Errorf("user does not have permission to delete CP alignments")
	}

	// 2. Persist
	if err := s.cpAlignmentRepo.DeleteCPAlignment(ctx, cmd.CPAlignmentID); err != nil {
		return nil, fmt.Errorf("failed to delete CP alignment: %w", err)
	}

	return &DeleteCPAlignmentResponse{
		Success: true,
	}, nil
}

// GetCPAlignmentsByCurriculumSubjectCommand represents the command to get CP alignments by curriculum subject
type GetCPAlignmentsByCurriculumSubjectCommand struct {
	CurriculumSubjectID string
	UserID              string
}

// GetCPAlignmentsByCurriculumSubjectResponse represents the response for getting CP alignments by curriculum subject
type GetCPAlignmentsByCurriculumSubjectResponse struct {
	CPAlignments []*domain.CPAlignment
}

// GetCPAlignmentsByCurriculumSubject retrieves CP alignments for a curriculum subject
func (s *CurriculumGovernanceApplicationService) GetCPAlignmentsByCurriculumSubject(ctx context.Context, cmd *GetCPAlignmentsByCurriculumSubjectCommand) (*GetCPAlignmentsByCurriculumSubjectResponse, error) {
	// 1. Authorization: Get user
	user, err := s.userRepo.GetByID(ctx, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Curriculum Admin, School Admin, Teacher, and System Admin can read
	if user.RoleID != domain.RoleCurriculumAdmin && user.RoleID != domain.RoleSchoolAdmin && user.RoleID != domain.RoleTeacher && user.RoleID != domain.RoleSystemAdmin {
		return nil, fmt.Errorf("user does not have permission to view CP alignments")
	}

	// 2. Load CP alignments
	alignments, err := s.cpAlignmentRepo.GetCPAlignmentsByCurriculumSubjectID(ctx, cmd.CurriculumSubjectID)
	if err != nil {
		return nil, fmt.Errorf("failed to load CP alignments: %w", err)
	}

	return &GetCPAlignmentsByCurriculumSubjectResponse{
		CPAlignments: alignments,
	}, nil
}

// GenerateCPAlignmentReportCommand represents the command to generate CP alignment report
type GenerateCPAlignmentReportCommand struct {
	UserID string
}

// GenerateCPAlignmentReportResponse represents the response for generating CP alignment report
type GenerateCPAlignmentReportResponse struct {
	Reports []*domain.CPAlignmentReport
}

// GenerateCPAlignmentReport generates a report showing CP coverage across all graduate profile dimensions
// This implements BR-004: Minimum CP coverage percentage
func (s *CurriculumGovernanceApplicationService) GenerateCPAlignmentReport(ctx context.Context, cmd *GenerateCPAlignmentReportCommand) (*GenerateCPAlignmentReportResponse, error) {
	// 1. Authorization: Get user
	user, err := s.userRepo.GetByID(ctx, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// School Admin and System Admin can generate reports
	if user.RoleID != domain.RoleSchoolAdmin && user.RoleID != domain.RoleSystemAdmin {
		return nil, fmt.Errorf("user does not have permission to generate CP alignment reports")
	}

	// 2. Get CP alignment threshold from system configuration
	threshold := domain.GetDefaultCPAlignmentThreshold()
	config, err := s.systemConfigRepo.GetSystemConfigurationByKey(ctx, domain.ConfigCPAlignmentThreshold)
	if err == nil && config != nil {
		// Parse the value (assuming it's stored as a string representation of a number)
		// In a real implementation, you'd parse based on value_type
		// For now, use the default
	}

	// 3. Generate report
	reports, err := s.cpAlignmentRepo.GenerateCPAlignmentReport(ctx, threshold)
	if err != nil {
		return nil, fmt.Errorf("failed to generate CP alignment report: %w", err)
	}

	return &GenerateCPAlignmentReportResponse{
		Reports: reports,
	}, nil
}

// ListCPAlignmentsCommand represents the command to list all CP alignments
type ListCPAlignmentsCommand struct {
	UserID string
}

// ListCPAlignmentsResponse represents the response for listing CP alignments
type ListCPAlignmentsResponse struct {
	CPAlignments []*domain.CPAlignment
}

// ListCPAlignments retrieves all CP alignments
func (s *CurriculumGovernanceApplicationService) ListCPAlignments(ctx context.Context, cmd *ListCPAlignmentsCommand) (*ListCPAlignmentsResponse, error) {
	// 1. Authorization: Get user
	user, err := s.userRepo.GetByID(ctx, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Curriculum Admin, School Admin, Teacher, and System Admin can read
	if user.RoleID != domain.RoleCurriculumAdmin && user.RoleID != domain.RoleSchoolAdmin && user.RoleID != domain.RoleTeacher && user.RoleID != domain.RoleSystemAdmin {
		return nil, fmt.Errorf("user does not have permission to view CP alignments")
	}

	// 2. Load CP alignments
	alignments, err := s.cpAlignmentRepo.GetAllCPAlignments(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load CP alignments: %w", err)
	}

	return &ListCPAlignmentsResponse{
		CPAlignments: alignments,
	}, nil
}

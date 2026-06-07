package service

import (
	"context"
	"fmt"

	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/repository"
)

// SchoolService handles business logic for school operations
type SchoolService struct {
	schoolRepo *repository.SchoolRepository
}

// NewSchoolService creates a new school service
func NewSchoolService(schoolRepo *repository.SchoolRepository) *SchoolService {
	return &SchoolService{
		schoolRepo: schoolRepo,
	}
}

// CreateSchool creates a new school
func (s *SchoolService) CreateSchool(ctx context.Context, req *domain.CreateSchoolRequest, creatorID string) (*domain.School, error) {
	// Check if school code already exists
	_, err := s.schoolRepo.GetByCode(ctx, req.Code)
	if err == nil {
		return nil, fmt.Errorf("school with code %s already exists", req.Code)
	}

	var address, phone, email *string
	if req.Address != "" {
		address = &req.Address
	}
	if req.Phone != "" {
		phone = &req.Phone
	}
	if req.Email != "" {
		email = &req.Email
	}

	school := &domain.School{
		ID:        domain.NewID(),
		Name:      req.Name,
		Code:      req.Code,
		Address:   address,
		Phone:     phone,
		Email:     email,
		IsActive:  true,
		CreatedBy: &creatorID,
		UpdatedBy: &creatorID,
	}

	if err := s.schoolRepo.Create(ctx, school); err != nil {
		return nil, fmt.Errorf("failed to create school: %w", err)
	}

	return school, nil
}

// GetSchool retrieves a school by ID
func (s *SchoolService) GetSchool(ctx context.Context, id string) (*domain.School, error) {
	return s.schoolRepo.GetByID(ctx, id)
}

// GetSchoolByCode retrieves a school by code
func (s *SchoolService) GetSchoolByCode(ctx context.Context, code string) (*domain.School, error) {
	return s.schoolRepo.GetByCode(ctx, code)
}

// ListSchools retrieves schools with pagination and filters
func (s *SchoolService) ListSchools(ctx context.Context, isActive *bool, page, pageSize int) ([]*domain.School, int, error) {
	limit := pageSize
	offset := (page - 1) * pageSize

	schools, err := s.schoolRepo.List(ctx, isActive, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.schoolRepo.Count(ctx, isActive)
	if err != nil {
		return nil, 0, err
	}

	return schools, total, nil
}

// UpdateSchool updates school information
func (s *SchoolService) UpdateSchool(ctx context.Context, id string, req *domain.UpdateSchoolRequest, updaterID string) (*domain.School, error) {
	school, err := s.schoolRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("school not found")
	}

	if req.Name != nil {
		school.Name = *req.Name
	}
	if req.Address != nil {
		school.Address = req.Address
	}
	if req.Phone != nil {
		school.Phone = req.Phone
	}
	if req.Email != nil {
		school.Email = req.Email
	}

	school.UpdatedBy = &updaterID

	if err := s.schoolRepo.Update(ctx, school); err != nil {
		return nil, fmt.Errorf("failed to update school: %w", err)
	}

	return school, nil
}

// UpdateSchoolStatus updates school status
func (s *SchoolService) UpdateSchoolStatus(ctx context.Context, id string, status domain.SchoolStatus) error {
	isActive := status == domain.SchoolStatusActive
	return s.schoolRepo.UpdateStatus(ctx, id, isActive)
}

// DeleteSchool soft deletes a school
func (s *SchoolService) DeleteSchool(ctx context.Context, id string) error {
	return s.schoolRepo.Delete(ctx, id)
}

package service

import (
	"context"
	"fmt"
	"time"

	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/repository"
)

// AnnouncementService handles business logic for announcement operations
type AnnouncementService struct {
	announcementRepo *repository.AnnouncementRepository
	schoolRepo       *repository.SchoolRepository
	userRepo         *repository.UserRepository
}

// NewAnnouncementService creates a new announcement service
func NewAnnouncementService(
	announcementRepo *repository.AnnouncementRepository,
	schoolRepo *repository.SchoolRepository,
	userRepo *repository.UserRepository,
) *AnnouncementService {
	return &AnnouncementService{
		announcementRepo: announcementRepo,
		schoolRepo:       schoolRepo,
		userRepo:         userRepo,
	}
}

// CreateAnnouncement creates a new announcement
func (s *AnnouncementService) CreateAnnouncement(ctx context.Context, req *domain.CreateAnnouncementRequest, publisherID string) (*domain.Announcement, error) {
	// Verify school exists
	_, err := s.schoolRepo.GetByID(ctx, req.SchoolID)
	if err != nil {
		return nil, fmt.Errorf("school not found")
	}

	// Verify publisher exists and is active
	publisher, err := s.userRepo.GetByID(ctx, publisherID)
	if err != nil {
		return nil, fmt.Errorf("publisher not found")
	}
	if !publisher.IsActive {
		return nil, fmt.Errorf("publisher is not active")
	}

	announcement := &domain.Announcement{
		ID:             domain.NewID(),
		SchoolID:       req.SchoolID,
		Title:          req.Title,
		Content:        req.Content,
		Priority:       string(req.Priority),
		TargetAudience: string(req.TargetAudience),
		PublishedBy:    publisherID,
		PublishedAt:    time.Now(),
		ExpiresAt:      req.ExpiresAt,
		IsActive:       true,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := s.announcementRepo.Create(ctx, announcement); err != nil {
		return nil, fmt.Errorf("failed to create announcement: %w", err)
	}

	return announcement, nil
}

// GetAnnouncement retrieves an announcement by ID
func (s *AnnouncementService) GetAnnouncement(ctx context.Context, id string) (*domain.Announcement, error) {
	announcement, err := s.announcementRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("announcement not found")
	}
	return announcement, nil
}

// ListAnnouncements retrieves announcements with filters and pagination
func (s *AnnouncementService) ListAnnouncements(ctx context.Context, schoolID *string, priority *string, targetAudience *string, isActive *bool, page, pageSize int) ([]*domain.Announcement, int, error) {
	limit := pageSize
	offset := (page - 1) * pageSize

	announcements, err := s.announcementRepo.List(ctx, schoolID, priority, targetAudience, isActive, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list announcements: %w", err)
	}

	total, err := s.announcementRepo.Count(ctx, schoolID, priority, targetAudience, isActive)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count announcements: %w", err)
	}

	return announcements, total, nil
}

// UpdateAnnouncement updates announcement information
func (s *AnnouncementService) UpdateAnnouncement(ctx context.Context, id string, req *domain.UpdateAnnouncementRequest) (*domain.Announcement, error) {
	announcement, err := s.announcementRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("announcement not found")
	}

	if req.Title != nil {
		announcement.Title = *req.Title
	}
	if req.Content != nil {
		announcement.Content = *req.Content
	}
	if req.Priority != nil {
		announcement.Priority = string(*req.Priority)
	}
	if req.TargetAudience != nil {
		announcement.TargetAudience = string(*req.TargetAudience)
	}
	if req.ExpiresAt != nil {
		announcement.ExpiresAt = req.ExpiresAt
	}
	if req.IsActive != nil {
		announcement.IsActive = *req.IsActive
	}

	if err := s.announcementRepo.Update(ctx, announcement); err != nil {
		return nil, fmt.Errorf("failed to update announcement: %w", err)
	}

	return announcement, nil
}

// DeleteAnnouncement soft deletes an announcement
func (s *AnnouncementService) DeleteAnnouncement(ctx context.Context, id string) error {
	return s.announcementRepo.Delete(ctx, id)
}

// GetSchoolAnnouncements retrieves all announcements for a school
func (s *AnnouncementService) GetSchoolAnnouncements(ctx context.Context, schoolID string) ([]*domain.Announcement, error) {
	// Verify school exists
	_, err := s.schoolRepo.GetByID(ctx, schoolID)
	if err != nil {
		return nil, fmt.Errorf("school not found")
	}

	announcements, err := s.announcementRepo.GetBySchoolID(ctx, schoolID)
	if err != nil {
		return nil, fmt.Errorf("failed to get school announcements: %w", err)
	}

	return announcements, nil
}

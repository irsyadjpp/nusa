package service

import (
	"context"
	"fmt"
	"time"

	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/repository"
)

// ScheduleService handles business logic for schedule operations
type ScheduleService struct {
	scheduleRepo repository.ScheduleRepositoryInterface
	classRepo    repository.ClassRepositoryInterface
	userRepo     repository.UserRepositoryInterface
}

// NewScheduleService creates a new schedule service
func NewScheduleService(
	scheduleRepo repository.ScheduleRepositoryInterface,
	classRepo repository.ClassRepositoryInterface,
	userRepo repository.UserRepositoryInterface,
) *ScheduleService {
	return &ScheduleService{
		scheduleRepo: scheduleRepo,
		classRepo:    classRepo,
		userRepo:     userRepo,
	}
}

// CreateSchedule creates a new schedule
func (s *ScheduleService) CreateSchedule(ctx context.Context, req *domain.CreateScheduleRequest, creatorID string) (*domain.Schedule, error) {
	// Verify class exists and is active
	class, err := s.classRepo.GetByID(ctx, req.ClassID)
	if err != nil {
		return nil, fmt.Errorf("class not found")
	}
	if !class.IsActive {
		return nil, fmt.Errorf("class is not active")
	}

	schedule := &domain.Schedule{
		ID:        domain.NewID(),
		ClassID:   req.ClassID,
		DayOfWeek: req.DayOfWeek,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Room:      req.Room,
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: &creatorID,
		UpdatedBy: &creatorID,
	}

	if err := s.scheduleRepo.Create(ctx, schedule); err != nil {
		return nil, fmt.Errorf("failed to create schedule: %w", err)
	}

	return schedule, nil
}

// GetSchedule retrieves a schedule by ID
func (s *ScheduleService) GetSchedule(ctx context.Context, id string) (*domain.Schedule, error) {
	schedule, err := s.scheduleRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("schedule not found")
	}
	return schedule, nil
}

// ListSchedules retrieves schedules with filters and pagination
func (s *ScheduleService) ListSchedules(ctx context.Context, classID *string, dayOfWeek *int, isActive *bool, page, pageSize int) ([]*domain.Schedule, int, error) {
	limit := pageSize
	offset := (page - 1) * pageSize

	schedules, err := s.scheduleRepo.List(ctx, classID, dayOfWeek, isActive, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list schedules: %w", err)
	}

	total, err := s.scheduleRepo.Count(ctx, classID, dayOfWeek, isActive)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count schedules: %w", err)
	}

	return schedules, total, nil
}

// UpdateSchedule updates schedule information
func (s *ScheduleService) UpdateSchedule(ctx context.Context, id string, req *domain.UpdateScheduleRequest, updaterID string) (*domain.Schedule, error) {
	schedule, err := s.scheduleRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("schedule not found")
	}

	if req.DayOfWeek != nil {
		schedule.DayOfWeek = *req.DayOfWeek
	}
	if req.StartTime != nil {
		schedule.StartTime = *req.StartTime
	}
	if req.EndTime != nil {
		schedule.EndTime = *req.EndTime
	}
	if req.Room != nil {
		schedule.Room = req.Room
	}
	if req.IsActive != nil {
		schedule.IsActive = *req.IsActive
	}

	schedule.UpdatedBy = &updaterID
	schedule.UpdatedAt = time.Now()

	if err := s.scheduleRepo.Update(ctx, schedule); err != nil {
		return nil, fmt.Errorf("failed to update schedule: %w", err)
	}

	return schedule, nil
}

// DeleteSchedule soft deletes a schedule
func (s *ScheduleService) DeleteSchedule(ctx context.Context, id string) error {
	return s.scheduleRepo.Delete(ctx, id)
}

// GetClassSchedules retrieves all schedules for a class
func (s *ScheduleService) GetClassSchedules(ctx context.Context, classID string) ([]*domain.Schedule, error) {
	// Verify class exists
	_, err := s.classRepo.GetByID(ctx, classID)
	if err != nil {
		return nil, fmt.Errorf("class not found")
	}

	schedules, err := s.scheduleRepo.GetByClassID(ctx, classID)
	if err != nil {
		return nil, fmt.Errorf("failed to get class schedules: %w", err)
	}

	return schedules, nil
}

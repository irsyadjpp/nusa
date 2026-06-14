package service

import (
	"context"
	"fmt"
	"time"

	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/repository"
)

// AttendanceService handles business logic for attendance operations
type AttendanceService struct {
	attendanceRepo *repository.AttendanceRepository
	classRepo      *repository.ClassRepository
	userRepo       *repository.UserRepository
}

// NewAttendanceService creates a new attendance service
func NewAttendanceService(
	attendanceRepo *repository.AttendanceRepository,
	classRepo *repository.ClassRepository,
	userRepo *repository.UserRepository,
) *AttendanceService {
	return &AttendanceService{
		attendanceRepo: attendanceRepo,
		classRepo:      classRepo,
		userRepo:       userRepo,
	}
}

// RecordAttendance records attendance for a student
func (s *AttendanceService) RecordAttendance(ctx context.Context, req *domain.CreateAttendanceRequest, recorderID string) (*domain.AttendanceRecord, error) {
	// Verify class exists and is active
	class, err := s.classRepo.GetByID(ctx, req.ClassID)
	if err != nil {
		return nil, fmt.Errorf("class not found")
	}
	if !class.IsActive {
		return nil, fmt.Errorf("class is not active")
	}

	// Verify student exists and is active
	student, err := s.userRepo.GetByID(ctx, req.StudentID)
	if err != nil {
		return nil, fmt.Errorf("student not found")
	}
	if !student.IsActive {
		return nil, fmt.Errorf("student is not active")
	}

	// Verify recorder exists and is active
	recorder, err := s.userRepo.GetByID(ctx, recorderID)
	if err != nil {
		return nil, fmt.Errorf("recorder not found")
	}
	if !recorder.IsActive {
		return nil, fmt.Errorf("recorder is not active")
	}

	attendance := &domain.AttendanceRecord{
		ID:         domain.NewID(),
		ClassID:    req.ClassID,
		StudentID:  req.StudentID,
		Date:       req.Date,
		Status:     string(req.Status),
		Notes:      req.Notes,
		RecordedBy: recorderID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := s.attendanceRepo.Create(ctx, attendance); err != nil {
		return nil, fmt.Errorf("failed to record attendance: %w", err)
	}

	return attendance, nil
}

// GetAttendance retrieves an attendance record by ID
func (s *AttendanceService) GetAttendance(ctx context.Context, id string) (*domain.AttendanceRecord, error) {
	attendance, err := s.attendanceRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("attendance record not found")
	}
	return attendance, nil
}

// ListAttendances retrieves attendance records with filters and pagination
func (s *AttendanceService) ListAttendances(ctx context.Context, classID, studentID *string, status *string, startDate, endDate *time.Time, page, pageSize int) ([]*domain.AttendanceRecord, int, error) {
	limit := pageSize
	offset := (page - 1) * pageSize

	attendances, err := s.attendanceRepo.List(ctx, classID, studentID, status, startDate, endDate, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list attendances: %w", err)
	}

	total, err := s.attendanceRepo.Count(ctx, classID, studentID, status, startDate, endDate)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count attendances: %w", err)
	}

	return attendances, total, nil
}

// UpdateAttendance updates an attendance record
func (s *AttendanceService) UpdateAttendance(ctx context.Context, id string, req *domain.UpdateAttendanceRequest, recorderID string) (*domain.AttendanceRecord, error) {
	attendance, err := s.attendanceRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("attendance record not found")
	}

	// Verify recorder exists and is active
	recorder, err := s.userRepo.GetByID(ctx, recorderID)
	if err != nil {
		return nil, fmt.Errorf("recorder not found")
	}
	if !recorder.IsActive {
		return nil, fmt.Errorf("recorder is not active")
	}

	attendance.Status = string(req.Status)
	attendance.Notes = req.Notes
	attendance.RecordedBy = recorderID
	attendance.UpdatedAt = time.Now()

	if err := s.attendanceRepo.Update(ctx, attendance); err != nil {
		return nil, fmt.Errorf("failed to update attendance: %w", err)
	}

	return attendance, nil
}

// DeleteAttendance soft deletes an attendance record
func (s *AttendanceService) DeleteAttendance(ctx context.Context, id string) error {
	return s.attendanceRepo.Delete(ctx, id)
}

// GetClassAttendanceStats returns attendance statistics for a class
func (s *AttendanceService) GetClassAttendanceStats(ctx context.Context, classID string, startDate, endDate time.Time) (map[string]int, error) {
	// Verify class exists
	_, err := s.classRepo.GetByID(ctx, classID)
	if err != nil {
		return nil, fmt.Errorf("class not found")
	}

	stats, err := s.attendanceRepo.GetAttendanceStats(ctx, classID, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to get class attendance stats: %w", err)
	}

	return stats, nil
}

// GetStudentAttendanceStats returns attendance statistics for a student
func (s *AttendanceService) GetStudentAttendanceStats(ctx context.Context, studentID string, startDate, endDate time.Time) (map[string]int, error) {
	// Verify student exists
	_, err := s.userRepo.GetByID(ctx, studentID)
	if err != nil {
		return nil, fmt.Errorf("student not found")
	}

	stats, err := s.attendanceRepo.GetStudentAttendanceStats(ctx, studentID, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to get student attendance stats: %w", err)
	}

	return stats, nil
}

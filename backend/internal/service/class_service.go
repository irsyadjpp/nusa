package service

import (
	"context"
	"fmt"
	"time"

	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/repository"
)

// ClassService handles business logic for class operations
type ClassService struct {
	classRepo        *repository.ClassRepository
	enrollmentRepo   *repository.ClassEnrollmentRepository
	userRepo         *repository.UserRepository
	academicYearRepo *repository.AcademicYearRepository
	semesterRepo     *repository.SemesterRepository
}

// NewClassService creates a new class service
func NewClassService(
	classRepo *repository.ClassRepository,
	enrollmentRepo *repository.ClassEnrollmentRepository,
	userRepo *repository.UserRepository,
	academicYearRepo *repository.AcademicYearRepository,
	semesterRepo *repository.SemesterRepository,
) *ClassService {
	return &ClassService{
		classRepo:        classRepo,
		enrollmentRepo:   enrollmentRepo,
		userRepo:         userRepo,
		academicYearRepo: academicYearRepo,
		semesterRepo:     semesterRepo,
	}
}

// CreateClass creates a new class
func (s *ClassService) CreateClass(ctx context.Context, req *domain.CreateClassRequest, creatorID string) (*domain.Class, error) {
	// Verify school exists
	_, err := s.userRepo.GetByID(ctx, req.SchoolID)
	if err != nil {
		return nil, fmt.Errorf("invalid school: %w", err)
	}

	// Verify academic year exists
	_, err = s.academicYearRepo.GetAcademicYearByID(ctx, req.AcademicYearID)
	if err != nil {
		return nil, fmt.Errorf("invalid academic year: %w", err)
	}

	// Verify semester exists
	_, err = s.semesterRepo.GetSemesterByID(ctx, req.SemesterID)
	if err != nil {
		return nil, fmt.Errorf("invalid semester: %w", err)
	}

	// Verify teacher exists and is active
	teacher, err := s.userRepo.GetByID(ctx, req.TeacherID)
	if err != nil {
		return nil, fmt.Errorf("invalid teacher: %w", err)
	}
	if !teacher.IsActive {
		return nil, fmt.Errorf("teacher is not active")
	}

	class := &domain.Class{
		ID:             domain.NewID(),
		SchoolID:       req.SchoolID,
		AcademicYearID: req.AcademicYearID,
		SemesterID:     req.SemesterID,
		SubjectID:      req.SubjectID,
		TeacherID:      req.TeacherID,
		Name:           req.Name,
		GradeLevel:     req.GradeLevel,
		Room:           req.Room,
		MaxStudents:    req.MaxStudents,
		IsActive:       true,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		CreatedBy:      &creatorID,
		UpdatedBy:      &creatorID,
	}

	if err := s.classRepo.Create(ctx, class); err != nil {
		return nil, fmt.Errorf("failed to create class: %w", err)
	}

	return class, nil
}

// GetClass retrieves a class by ID
func (s *ClassService) GetClass(ctx context.Context, id string) (*domain.Class, error) {
	class, err := s.classRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("class not found")
	}
	return class, nil
}

// ListClasses retrieves classes with pagination and filters
func (s *ClassService) ListClasses(ctx context.Context, schoolID, academicYearID, semesterID, subjectID, teacherID *string, isActive *bool, page, pageSize int) ([]*domain.Class, int, error) {
	limit := pageSize
	offset := (page - 1) * pageSize

	classes, err := s.classRepo.List(ctx, schoolID, academicYearID, semesterID, subjectID, teacherID, isActive, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list classes: %w", err)
	}

	total, err := s.classRepo.Count(ctx, schoolID, academicYearID, semesterID, subjectID, teacherID, isActive)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count classes: %w", err)
	}

	return classes, total, nil
}

// UpdateClass updates class information
func (s *ClassService) UpdateClass(ctx context.Context, id string, req *domain.UpdateClassRequest, updaterID string) (*domain.Class, error) {
	class, err := s.classRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("class not found")
	}

	if req.Name != nil {
		class.Name = *req.Name
	}
	if req.GradeLevel != nil {
		class.GradeLevel = *req.GradeLevel
	}
	if req.Room != nil {
		class.Room = req.Room
	}
	if req.MaxStudents != nil {
		// Check if current students exceed new max
		currentCount, err := s.classRepo.GetStudentCount(ctx, id)
		if err != nil {
			return nil, fmt.Errorf("failed to check student count: %w", err)
		}
		if currentCount > *req.MaxStudents {
			return nil, fmt.Errorf("cannot reduce max students below current enrollment count")
		}
		class.MaxStudents = *req.MaxStudents
	}
	if req.IsActive != nil {
		class.IsActive = *req.IsActive
	}

	class.UpdatedBy = &updaterID
	class.UpdatedAt = time.Now()

	if err := s.classRepo.Update(ctx, class); err != nil {
		return nil, fmt.Errorf("failed to update class: %w", err)
	}

	return class, nil
}

// DeleteClass soft deletes a class
func (s *ClassService) DeleteClass(ctx context.Context, id string) error {
	// Check if class has active enrollments
	status := string(domain.EnrollmentStatusActive)
	enrollments, err := s.enrollmentRepo.List(ctx, &id, nil, &status, 1, 0)
	if err != nil {
		return fmt.Errorf("failed to check enrollments: %w", err)
	}
	if len(enrollments) > 0 {
		return fmt.Errorf("cannot delete class with active enrollments")
	}

	return s.classRepo.Delete(ctx, id)
}

// EnrollStudent enrolls a student in a class
func (s *ClassService) EnrollStudent(ctx context.Context, req *domain.CreateClassEnrollmentRequest, creatorID string) (*domain.ClassEnrollment, error) {
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

	// Check if student is already enrolled
	alreadyEnrolled, err := s.enrollmentRepo.CheckEnrollment(ctx, req.ClassID, req.StudentID)
	if err != nil {
		return nil, fmt.Errorf("failed to check enrollment: %w", err)
	}
	if alreadyEnrolled {
		return nil, fmt.Errorf("student is already enrolled in this class")
	}

	// Check if class is at capacity
	currentCount, err := s.classRepo.GetStudentCount(ctx, req.ClassID)
	if err != nil {
		return nil, fmt.Errorf("failed to check student count: %w", err)
	}
	if currentCount >= class.MaxStudents {
		return nil, fmt.Errorf("class is at maximum capacity")
	}

	enrollment := &domain.ClassEnrollment{
		ID:             domain.NewID(),
		ClassID:        req.ClassID,
		StudentID:      req.StudentID,
		EnrollmentDate: time.Now(),
		Status:         string(domain.EnrollmentStatusActive),
		Notes:          req.Notes,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := s.enrollmentRepo.Create(ctx, enrollment); err != nil {
		return nil, fmt.Errorf("failed to create enrollment: %w", err)
	}

	return enrollment, nil
}

// WithdrawStudent withdraws a student from a class
func (s *ClassService) WithdrawStudent(ctx context.Context, classID, studentID string) error {
	status := string(domain.EnrollmentStatusActive)
	enrollments, err := s.enrollmentRepo.List(ctx, &classID, &studentID, &status, 1, 0)
	if err != nil {
		return fmt.Errorf("failed to find enrollment: %w", err)
	}
	if len(enrollments) == 0 {
		return fmt.Errorf("enrollment not found")
	}

	enrollment := enrollments[0]
	enrollment.Status = string(domain.EnrollmentStatusWithdrawn)
	enrollment.UpdatedAt = time.Now()

	if err := s.enrollmentRepo.Update(ctx, enrollment); err != nil {
		return fmt.Errorf("failed to withdraw student: %w", err)
	}

	return nil
}

// ListEnrollments retrieves enrollments with pagination and filters
func (s *ClassService) ListEnrollments(ctx context.Context, classID, studentID *string, status *string, page, pageSize int) ([]*domain.ClassEnrollment, error) {
	limit := pageSize
	offset := (page - 1) * pageSize

	enrollments, err := s.enrollmentRepo.List(ctx, classID, studentID, status, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list enrollments: %w", err)
	}

	return enrollments, nil
}

// UpdateEnrollment updates enrollment status
func (s *ClassService) UpdateEnrollment(ctx context.Context, id string, req *domain.UpdateClassEnrollmentRequest) error {
	enrollment, err := s.enrollmentRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("enrollment not found")
	}

	enrollment.Status = string(req.Status)
	enrollment.Notes = req.Notes
	enrollment.UpdatedAt = time.Now()

	if err := s.enrollmentRepo.Update(ctx, enrollment); err != nil {
		return fmt.Errorf("failed to update enrollment: %w", err)
	}

	return nil
}

// GetStudentCount returns the number of students enrolled in a class
func (s *ClassService) GetStudentCount(ctx context.Context, classID string) (int, error) {
	return s.classRepo.GetStudentCount(ctx, classID)
}

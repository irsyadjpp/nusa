package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/nusa/backend/internal/domain"
)

// MockClassEnrollmentRepository is a mock implementation of ClassEnrollmentRepositoryInterface
type MockClassEnrollmentRepository struct {
	mock.Mock
}

func (m *MockClassEnrollmentRepository) Create(ctx context.Context, enrollment *domain.ClassEnrollment) error {
	args := m.Called(ctx, enrollment)
	return args.Error(0)
}

func (m *MockClassEnrollmentRepository) GetByID(ctx context.Context, id string) (*domain.ClassEnrollment, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.ClassEnrollment), args.Error(1)
}

func (m *MockClassEnrollmentRepository) List(ctx context.Context, classID, studentID *string, status *string, limit, offset int) ([]*domain.ClassEnrollment, error) {
	args := m.Called(ctx, classID, studentID, status, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.ClassEnrollment), args.Error(1)
}

func (m *MockClassEnrollmentRepository) Update(ctx context.Context, enrollment *domain.ClassEnrollment) error {
	args := m.Called(ctx, enrollment)
	return args.Error(0)
}

func (m *MockClassEnrollmentRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockClassEnrollmentRepository) CheckEnrollment(ctx context.Context, classID, studentID string) (bool, error) {
	args := m.Called(ctx, classID, studentID)
	return args.Bool(0), args.Error(1)
}

// MockAcademicYearRepository is a mock implementation of AcademicYearRepositoryInterface
type MockAcademicYearRepository struct {
	mock.Mock
}

func (m *MockAcademicYearRepository) Create(ctx context.Context, academicYear *domain.AcademicYear) error {
	args := m.Called(ctx, academicYear)
	return args.Error(0)
}

func (m *MockAcademicYearRepository) GetByID(ctx context.Context, id string) (*domain.AcademicYear, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.AcademicYear), args.Error(1)
}

func (m *MockAcademicYearRepository) GetAcademicYearByID(ctx context.Context, id string) (*domain.AcademicYear, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.AcademicYear), args.Error(1)
}

func (m *MockAcademicYearRepository) GetByYear(ctx context.Context, year int) (*domain.AcademicYear, error) {
	args := m.Called(ctx, year)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.AcademicYear), args.Error(1)
}

func (m *MockAcademicYearRepository) List(ctx context.Context, isActive *bool, limit, offset int) ([]*domain.AcademicYear, error) {
	args := m.Called(ctx, isActive, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.AcademicYear), args.Error(1)
}

func (m *MockAcademicYearRepository) Update(ctx context.Context, academicYear *domain.AcademicYear) error {
	args := m.Called(ctx, academicYear)
	return args.Error(0)
}

func (m *MockAcademicYearRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockAcademicYearRepository) GetCurrent(ctx context.Context) (*domain.AcademicYear, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.AcademicYear), args.Error(1)
}

// MockSemesterRepository is a mock implementation of SemesterRepositoryInterface
type MockSemesterRepository struct {
	mock.Mock
}

func (m *MockSemesterRepository) Create(ctx context.Context, semester *domain.Semester) error {
	args := m.Called(ctx, semester)
	return args.Error(0)
}

func (m *MockSemesterRepository) GetByID(ctx context.Context, id string) (*domain.Semester, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Semester), args.Error(1)
}

func (m *MockSemesterRepository) GetSemesterByID(ctx context.Context, id string) (*domain.Semester, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Semester), args.Error(1)
}

func (m *MockSemesterRepository) GetByCode(ctx context.Context, code string) (*domain.Semester, error) {
	args := m.Called(ctx, code)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Semester), args.Error(1)
}

func (m *MockSemesterRepository) List(ctx context.Context, academicYearID *string, isActive *bool, limit, offset int) ([]*domain.Semester, error) {
	args := m.Called(ctx, academicYearID, isActive, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Semester), args.Error(1)
}

func (m *MockSemesterRepository) Update(ctx context.Context, semester *domain.Semester) error {
	args := m.Called(ctx, semester)
	return args.Error(0)
}

func (m *MockSemesterRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockSemesterRepository) GetCurrent(ctx context.Context, academicYearID string) (*domain.Semester, error) {
	args := m.Called(ctx, academicYearID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Semester), args.Error(1)
}

func TestClassService_CreateClass_Success(t *testing.T) {
	mockClassRepo := new(MockClassRepository)
	mockEnrollmentRepo := new(MockClassEnrollmentRepository)
	mockUserRepo := new(MockUserRepository)
	mockAcademicYearRepo := new(MockAcademicYearRepository)
	mockSemesterRepo := new(MockSemesterRepository)
	service := NewClassService(mockClassRepo, mockEnrollmentRepo, mockUserRepo, mockAcademicYearRepo, mockSemesterRepo)

	req := &domain.CreateClassRequest{
		SchoolID:       "school-1",
		AcademicYearID: "year-1",
		SemesterID:     "semester-1",
		TeacherID:      "teacher-1",
		Name:           "Test Class",
		GradeLevel:     "10",
		Room:           stringPtr("Room 101"),
		MaxStudents:    30,
	}

	school := &domain.User{
		ID:   "school-1",
		Name: "School",
	}

	academicYear := &domain.AcademicYear{
		ID:   "year-1",
		Name: "2024-2025",
	}

	semester := &domain.Semester{
		ID:   "semester-1",
		Name: "Semester 1",
	}

	teacher := &domain.User{
		ID:       "teacher-1",
		Name:     "Teacher",
		IsActive: true,
	}

	mockUserRepo.On("GetByID", mock.Anything, "school-1").Return(school, nil)
	mockAcademicYearRepo.On("GetAcademicYearByID", mock.Anything, "year-1").Return(academicYear, nil)
	mockSemesterRepo.On("GetSemesterByID", mock.Anything, "semester-1").Return(semester, nil)
	mockUserRepo.On("GetByID", mock.Anything, "teacher-1").Return(teacher, nil)
	mockClassRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.Class")).Return(nil)

	result, err := service.CreateClass(context.Background(), req, "creator-1")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Test Class", result.Name)

	mockClassRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
	mockAcademicYearRepo.AssertExpectations(t)
	mockSemesterRepo.AssertExpectations(t)
}

func TestClassService_GetClass_Success(t *testing.T) {
	mockClassRepo := new(MockClassRepository)
	mockEnrollmentRepo := new(MockClassEnrollmentRepository)
	mockUserRepo := new(MockUserRepository)
	mockAcademicYearRepo := new(MockAcademicYearRepository)
	mockSemesterRepo := new(MockSemesterRepository)
	service := NewClassService(mockClassRepo, mockEnrollmentRepo, mockUserRepo, mockAcademicYearRepo, mockSemesterRepo)

	class := &domain.Class{
		ID:   "class-1",
		Name: "Test Class",
	}

	mockClassRepo.On("GetByID", mock.Anything, "class-1").Return(class, nil)

	result, err := service.GetClass(context.Background(), "class-1")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, class.ID, result.ID)

	mockClassRepo.AssertExpectations(t)
}

func TestClassService_ListClasses_Success(t *testing.T) {
	mockClassRepo := new(MockClassRepository)
	mockEnrollmentRepo := new(MockClassEnrollmentRepository)
	mockUserRepo := new(MockUserRepository)
	mockAcademicYearRepo := new(MockAcademicYearRepository)
	mockSemesterRepo := new(MockSemesterRepository)
	service := NewClassService(mockClassRepo, mockEnrollmentRepo, mockUserRepo, mockAcademicYearRepo, mockSemesterRepo)

	classes := []*domain.Class{
		{ID: "class-1", Name: "Class 1"},
		{ID: "class-2", Name: "Class 2"},
	}

	mockClassRepo.On("List", mock.Anything, (*string)(nil), (*string)(nil), (*string)(nil), (*string)(nil), (*string)(nil), (*bool)(nil), 10, 0).Return(classes, nil)
	mockClassRepo.On("Count", mock.Anything, (*string)(nil), (*string)(nil), (*string)(nil), (*string)(nil), (*string)(nil), (*bool)(nil)).Return(2, nil)

	result, total, err := service.ListClasses(context.Background(), nil, nil, nil, nil, nil, nil, 1, 10)

	require.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, 2, total)

	mockClassRepo.AssertExpectations(t)
}

func TestClassService_UpdateClass_Success(t *testing.T) {
	mockClassRepo := new(MockClassRepository)
	mockEnrollmentRepo := new(MockClassEnrollmentRepository)
	mockUserRepo := new(MockUserRepository)
	mockAcademicYearRepo := new(MockAcademicYearRepository)
	mockSemesterRepo := new(MockSemesterRepository)
	service := NewClassService(mockClassRepo, mockEnrollmentRepo, mockUserRepo, mockAcademicYearRepo, mockSemesterRepo)

	class := &domain.Class{
		ID:   "class-1",
		Name: "Old Name",
	}

	newName := "New Name"

	req := &domain.UpdateClassRequest{
		Name: &newName,
	}

	mockClassRepo.On("GetByID", mock.Anything, "class-1").Return(class, nil)
	mockClassRepo.On("Update", mock.Anything, mock.AnythingOfType("*domain.Class")).Return(nil)

	result, err := service.UpdateClass(context.Background(), "class-1", req, "updater-1")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, newName, result.Name)

	mockClassRepo.AssertExpectations(t)
}

func TestClassService_DeleteClass_Success(t *testing.T) {
	mockClassRepo := new(MockClassRepository)
	mockEnrollmentRepo := new(MockClassEnrollmentRepository)
	mockUserRepo := new(MockUserRepository)
	mockAcademicYearRepo := new(MockAcademicYearRepository)
	mockSemesterRepo := new(MockSemesterRepository)
	service := NewClassService(mockClassRepo, mockEnrollmentRepo, mockUserRepo, mockAcademicYearRepo, mockSemesterRepo)

	status := string(domain.EnrollmentStatusActive)
	mockEnrollmentRepo.On("List", mock.Anything, &[]string{"class-1"}[0], (*string)(nil), &status, 1, 0).Return([]*domain.ClassEnrollment{}, nil)
	mockClassRepo.On("Delete", mock.Anything, "class-1").Return(nil)

	err := service.DeleteClass(context.Background(), "class-1")

	require.NoError(t, err)

	mockClassRepo.AssertExpectations(t)
	mockEnrollmentRepo.AssertExpectations(t)
}

func TestClassService_DeleteClass_HasActiveEnrollments(t *testing.T) {
	mockClassRepo := new(MockClassRepository)
	mockEnrollmentRepo := new(MockClassEnrollmentRepository)
	mockUserRepo := new(MockUserRepository)
	mockAcademicYearRepo := new(MockAcademicYearRepository)
	mockSemesterRepo := new(MockSemesterRepository)
	service := NewClassService(mockClassRepo, mockEnrollmentRepo, mockUserRepo, mockAcademicYearRepo, mockSemesterRepo)

	status := string(domain.EnrollmentStatusActive)
	enrollments := []*domain.ClassEnrollment{
		{ID: "enrollment-1", Status: status},
	}

	mockEnrollmentRepo.On("List", mock.Anything, &[]string{"class-1"}[0], (*string)(nil), &status, 1, 0).Return(enrollments, nil)

	err := service.DeleteClass(context.Background(), "class-1")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "cannot delete class with active enrollments")

	mockEnrollmentRepo.AssertExpectations(t)
}

func TestClassService_EnrollStudent_Success(t *testing.T) {
	mockClassRepo := new(MockClassRepository)
	mockEnrollmentRepo := new(MockClassEnrollmentRepository)
	mockUserRepo := new(MockUserRepository)
	mockAcademicYearRepo := new(MockAcademicYearRepository)
	mockSemesterRepo := new(MockSemesterRepository)
	service := NewClassService(mockClassRepo, mockEnrollmentRepo, mockUserRepo, mockAcademicYearRepo, mockSemesterRepo)

	req := &domain.CreateClassEnrollmentRequest{
		ClassID:   "class-1",
		StudentID: "student-1",
	}

	class := &domain.Class{
		ID:          "class-1",
		Name:        "Test Class",
		IsActive:    true,
		MaxStudents: 30,
	}

	student := &domain.User{
		ID:       "student-1",
		Name:     "Student",
		IsActive: true,
	}

	mockClassRepo.On("GetByID", mock.Anything, "class-1").Return(class, nil)
	mockUserRepo.On("GetByID", mock.Anything, "student-1").Return(student, nil)
	mockEnrollmentRepo.On("CheckEnrollment", mock.Anything, "class-1", "student-1").Return(false, nil)
	mockClassRepo.On("GetStudentCount", mock.Anything, "class-1").Return(10, nil)
	mockEnrollmentRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.ClassEnrollment")).Return(nil)

	result, err := service.EnrollStudent(context.Background(), req, "creator-1")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "class-1", result.ClassID)

	mockClassRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
	mockEnrollmentRepo.AssertExpectations(t)
}

func TestClassService_WithdrawStudent_Success(t *testing.T) {
	mockClassRepo := new(MockClassRepository)
	mockEnrollmentRepo := new(MockClassEnrollmentRepository)
	mockUserRepo := new(MockUserRepository)
	mockAcademicYearRepo := new(MockAcademicYearRepository)
	mockSemesterRepo := new(MockSemesterRepository)
	service := NewClassService(mockClassRepo, mockEnrollmentRepo, mockUserRepo, mockAcademicYearRepo, mockSemesterRepo)

	status := string(domain.EnrollmentStatusActive)
	enrollment := &domain.ClassEnrollment{
		ID:     "enrollment-1",
		Status: status,
	}

	mockEnrollmentRepo.On("List", mock.Anything, &[]string{"class-1"}[0], &[]string{"student-1"}[0], &status, 1, 0).Return([]*domain.ClassEnrollment{enrollment}, nil)
	mockEnrollmentRepo.On("Update", mock.Anything, mock.AnythingOfType("*domain.ClassEnrollment")).Return(nil)

	err := service.WithdrawStudent(context.Background(), "class-1", "student-1")

	require.NoError(t, err)

	mockEnrollmentRepo.AssertExpectations(t)
}

func TestClassService_ListEnrollments_Success(t *testing.T) {
	mockClassRepo := new(MockClassRepository)
	mockEnrollmentRepo := new(MockClassEnrollmentRepository)
	mockUserRepo := new(MockUserRepository)
	mockAcademicYearRepo := new(MockAcademicYearRepository)
	mockSemesterRepo := new(MockSemesterRepository)
	service := NewClassService(mockClassRepo, mockEnrollmentRepo, mockUserRepo, mockAcademicYearRepo, mockSemesterRepo)

	enrollments := []*domain.ClassEnrollment{
		{ID: "enrollment-1", ClassID: "class-1"},
	}

	mockEnrollmentRepo.On("List", mock.Anything, (*string)(nil), (*string)(nil), (*string)(nil), 10, 0).Return(enrollments, nil)

	result, err := service.ListEnrollments(context.Background(), nil, nil, nil, 1, 10)

	require.NoError(t, err)
	assert.Len(t, result, 1)

	mockEnrollmentRepo.AssertExpectations(t)
}

func stringPtr(s string) *string {
	return &s
}

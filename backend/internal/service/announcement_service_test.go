package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/nusa/backend/internal/domain"
)

// MockAnnouncementRepository is a mock implementation of AnnouncementRepositoryInterface
type MockAnnouncementRepository struct {
	mock.Mock
}

func (m *MockAnnouncementRepository) Create(ctx context.Context, announcement *domain.Announcement) error {
	args := m.Called(ctx, announcement)
	return args.Error(0)
}

func (m *MockAnnouncementRepository) GetByID(ctx context.Context, id string) (*domain.Announcement, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Announcement), args.Error(1)
}

func (m *MockAnnouncementRepository) List(ctx context.Context, schoolID, priority, targetAudience *string, isActive *bool, limit, offset int) ([]*domain.Announcement, error) {
	args := m.Called(ctx, schoolID, priority, targetAudience, isActive, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Announcement), args.Error(1)
}

func (m *MockAnnouncementRepository) Count(ctx context.Context, schoolID, priority, targetAudience *string, isActive *bool) (int, error) {
	args := m.Called(ctx, schoolID, priority, targetAudience, isActive)
	return args.Int(0), args.Error(1)
}

func (m *MockAnnouncementRepository) Update(ctx context.Context, announcement *domain.Announcement) error {
	args := m.Called(ctx, announcement)
	return args.Error(0)
}

func (m *MockAnnouncementRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockAnnouncementRepository) GetBySchoolID(ctx context.Context, schoolID string) ([]*domain.Announcement, error) {
	args := m.Called(ctx, schoolID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Announcement), args.Error(1)
}

func TestAnnouncementService_CreateAnnouncement_Success(t *testing.T) {
	mockAnnouncementRepo := new(MockAnnouncementRepository)
	mockSchoolRepo := new(MockSchoolRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewAnnouncementService(mockAnnouncementRepo, mockSchoolRepo, mockUserRepo)

	req := &domain.CreateAnnouncementRequest{
		SchoolID:       "school-1",
		Title:          "Test Announcement",
		Content:        "Test content",
		Priority:       "high",
		TargetAudience: "all",
		ExpiresAt:      nil,
	}

	school := &domain.School{
		ID:   "school-1",
		Name: "Test School",
	}

	publisher := &domain.User{
		ID:       "user-1",
		Name:     "Publisher",
		IsActive: true,
	}

	mockSchoolRepo.On("GetByID", mock.Anything, "school-1").Return(school, nil)
	mockUserRepo.On("GetByID", mock.Anything, "user-1").Return(publisher, nil)
	mockAnnouncementRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.Announcement")).Return(nil)

	result, err := service.CreateAnnouncement(context.Background(), req, "user-1")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Test Announcement", result.Title)
	assert.Equal(t, "school-1", result.SchoolID)
	assert.True(t, result.IsActive)

	mockAnnouncementRepo.AssertExpectations(t)
	mockSchoolRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestAnnouncementService_CreateAnnouncement_SchoolNotFound(t *testing.T) {
	mockAnnouncementRepo := new(MockAnnouncementRepository)
	mockSchoolRepo := new(MockSchoolRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewAnnouncementService(mockAnnouncementRepo, mockSchoolRepo, mockUserRepo)

	req := &domain.CreateAnnouncementRequest{
		SchoolID: "school-1",
		Title:    "Test Announcement",
	}

	mockSchoolRepo.On("GetByID", mock.Anything, "school-1").Return(nil, errors.New("not found"))

	result, err := service.CreateAnnouncement(context.Background(), req, "user-1")

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "school not found")

	mockSchoolRepo.AssertExpectations(t)
}

func TestAnnouncementService_CreateAnnouncement_PublisherNotFound(t *testing.T) {
	mockAnnouncementRepo := new(MockAnnouncementRepository)
	mockSchoolRepo := new(MockSchoolRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewAnnouncementService(mockAnnouncementRepo, mockSchoolRepo, mockUserRepo)

	req := &domain.CreateAnnouncementRequest{
		SchoolID: "school-1",
		Title:    "Test Announcement",
	}

	school := &domain.School{
		ID:   "school-1",
		Name: "Test School",
	}

	mockSchoolRepo.On("GetByID", mock.Anything, "school-1").Return(school, nil)
	mockUserRepo.On("GetByID", mock.Anything, "user-1").Return(nil, errors.New("not found"))

	result, err := service.CreateAnnouncement(context.Background(), req, "user-1")

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "publisher not found")

	mockSchoolRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestAnnouncementService_CreateAnnouncement_PublisherInactive(t *testing.T) {
	mockAnnouncementRepo := new(MockAnnouncementRepository)
	mockSchoolRepo := new(MockSchoolRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewAnnouncementService(mockAnnouncementRepo, mockSchoolRepo, mockUserRepo)

	req := &domain.CreateAnnouncementRequest{
		SchoolID: "school-1",
		Title:    "Test Announcement",
	}

	school := &domain.School{
		ID:   "school-1",
		Name: "Test School",
	}

	publisher := &domain.User{
		ID:       "user-1",
		Name:     "Publisher",
		IsActive: false,
	}

	mockSchoolRepo.On("GetByID", mock.Anything, "school-1").Return(school, nil)
	mockUserRepo.On("GetByID", mock.Anything, "user-1").Return(publisher, nil)

	result, err := service.CreateAnnouncement(context.Background(), req, "user-1")

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "publisher is not active")

	mockSchoolRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestAnnouncementService_GetAnnouncement_Success(t *testing.T) {
	mockAnnouncementRepo := new(MockAnnouncementRepository)
	mockSchoolRepo := new(MockSchoolRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewAnnouncementService(mockAnnouncementRepo, mockSchoolRepo, mockUserRepo)

	announcement := &domain.Announcement{
		ID:    "announcement-1",
		Title: "Test Announcement",
	}

	mockAnnouncementRepo.On("GetByID", mock.Anything, "announcement-1").Return(announcement, nil)

	result, err := service.GetAnnouncement(context.Background(), "announcement-1")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, announcement.ID, result.ID)

	mockAnnouncementRepo.AssertExpectations(t)
}

func TestAnnouncementService_GetAnnouncement_NotFound(t *testing.T) {
	mockAnnouncementRepo := new(MockAnnouncementRepository)
	mockSchoolRepo := new(MockSchoolRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewAnnouncementService(mockAnnouncementRepo, mockSchoolRepo, mockUserRepo)

	mockAnnouncementRepo.On("GetByID", mock.Anything, "announcement-1").Return(nil, errors.New("not found"))

	result, err := service.GetAnnouncement(context.Background(), "announcement-1")

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "announcement not found")

	mockAnnouncementRepo.AssertExpectations(t)
}

func TestAnnouncementService_ListAnnouncements_Success(t *testing.T) {
	mockAnnouncementRepo := new(MockAnnouncementRepository)
	mockSchoolRepo := new(MockSchoolRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewAnnouncementService(mockAnnouncementRepo, mockSchoolRepo, mockUserRepo)

	announcements := []*domain.Announcement{
		{ID: "announcement-1", Title: "Announcement 1"},
		{ID: "announcement-2", Title: "Announcement 2"},
	}

	mockAnnouncementRepo.On("List", mock.Anything, (*string)(nil), (*string)(nil), (*string)(nil), (*bool)(nil), 10, 0).Return(announcements, nil)
	mockAnnouncementRepo.On("Count", mock.Anything, (*string)(nil), (*string)(nil), (*string)(nil), (*bool)(nil)).Return(2, nil)

	result, total, err := service.ListAnnouncements(context.Background(), nil, nil, nil, nil, 1, 10)

	require.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, 2, total)

	mockAnnouncementRepo.AssertExpectations(t)
}

func TestAnnouncementService_UpdateAnnouncement_Success(t *testing.T) {
	mockAnnouncementRepo := new(MockAnnouncementRepository)
	mockSchoolRepo := new(MockSchoolRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewAnnouncementService(mockAnnouncementRepo, mockSchoolRepo, mockUserRepo)

	announcement := &domain.Announcement{
		ID:    "announcement-1",
		Title: "Old Title",
	}

	newTitle := "New Title"

	req := &domain.UpdateAnnouncementRequest{
		Title: &newTitle,
	}

	mockAnnouncementRepo.On("GetByID", mock.Anything, "announcement-1").Return(announcement, nil)
	mockAnnouncementRepo.On("Update", mock.Anything, mock.AnythingOfType("*domain.Announcement")).Return(nil)

	result, err := service.UpdateAnnouncement(context.Background(), "announcement-1", req)

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, newTitle, result.Title)

	mockAnnouncementRepo.AssertExpectations(t)
}

func TestAnnouncementService_UpdateAnnouncement_NotFound(t *testing.T) {
	mockAnnouncementRepo := new(MockAnnouncementRepository)
	mockSchoolRepo := new(MockSchoolRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewAnnouncementService(mockAnnouncementRepo, mockSchoolRepo, mockUserRepo)

	req := &domain.UpdateAnnouncementRequest{}

	mockAnnouncementRepo.On("GetByID", mock.Anything, "announcement-1").Return(nil, errors.New("not found"))

	result, err := service.UpdateAnnouncement(context.Background(), "announcement-1", req)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "announcement not found")

	mockAnnouncementRepo.AssertExpectations(t)
}

func TestAnnouncementService_DeleteAnnouncement_Success(t *testing.T) {
	mockAnnouncementRepo := new(MockAnnouncementRepository)
	mockSchoolRepo := new(MockSchoolRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewAnnouncementService(mockAnnouncementRepo, mockSchoolRepo, mockUserRepo)

	mockAnnouncementRepo.On("Delete", mock.Anything, "announcement-1").Return(nil)

	err := service.DeleteAnnouncement(context.Background(), "announcement-1")

	require.NoError(t, err)

	mockAnnouncementRepo.AssertExpectations(t)
}

func TestAnnouncementService_GetSchoolAnnouncements_Success(t *testing.T) {
	mockAnnouncementRepo := new(MockAnnouncementRepository)
	mockSchoolRepo := new(MockSchoolRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewAnnouncementService(mockAnnouncementRepo, mockSchoolRepo, mockUserRepo)

	school := &domain.School{
		ID:   "school-1",
		Name: "Test School",
	}

	announcements := []*domain.Announcement{
		{ID: "announcement-1", Title: "Announcement 1"},
	}

	mockSchoolRepo.On("GetByID", mock.Anything, "school-1").Return(school, nil)
	mockAnnouncementRepo.On("GetBySchoolID", mock.Anything, "school-1").Return(announcements, nil)

	result, err := service.GetSchoolAnnouncements(context.Background(), "school-1")

	require.NoError(t, err)
	assert.Len(t, result, 1)

	mockAnnouncementRepo.AssertExpectations(t)
	mockSchoolRepo.AssertExpectations(t)
}

func TestAnnouncementService_GetSchoolAnnouncements_SchoolNotFound(t *testing.T) {
	mockAnnouncementRepo := new(MockAnnouncementRepository)
	mockSchoolRepo := new(MockSchoolRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewAnnouncementService(mockAnnouncementRepo, mockSchoolRepo, mockUserRepo)

	mockSchoolRepo.On("GetByID", mock.Anything, "school-1").Return(nil, errors.New("not found"))

	result, err := service.GetSchoolAnnouncements(context.Background(), "school-1")

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "school not found")

	mockSchoolRepo.AssertExpectations(t)
}

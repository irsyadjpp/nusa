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

// MockNotificationRepository is a mock implementation of NotificationRepositoryInterface
type MockNotificationRepository struct {
	mock.Mock
}

func (m *MockNotificationRepository) Create(ctx context.Context, notification *domain.Notification) error {
	args := m.Called(ctx, notification)
	return args.Error(0)
}

func (m *MockNotificationRepository) GetByID(ctx context.Context, id string) (*domain.Notification, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Notification), args.Error(1)
}

func (m *MockNotificationRepository) List(ctx context.Context, userID, notificationType *string, isRead *bool, limit, offset int) ([]*domain.Notification, error) {
	args := m.Called(ctx, userID, notificationType, isRead, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Notification), args.Error(1)
}

func (m *MockNotificationRepository) Count(ctx context.Context, userID, notificationType *string, isRead *bool) (int, error) {
	args := m.Called(ctx, userID, notificationType, isRead)
	return args.Int(0), args.Error(1)
}

func (m *MockNotificationRepository) Update(ctx context.Context, notification *domain.Notification) error {
	args := m.Called(ctx, notification)
	return args.Error(0)
}

func (m *MockNotificationRepository) MarkAsRead(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockNotificationRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockNotificationRepository) GetUnreadCount(ctx context.Context, userID string) (int, error) {
	args := m.Called(ctx, userID)
	return args.Int(0), args.Error(1)
}

func TestNotificationService_CreateNotification_Success(t *testing.T) {
	mockNotificationRepo := new(MockNotificationRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewNotificationService(mockNotificationRepo, mockUserRepo)

	req := &domain.CreateNotificationRequest{
		UserID:  "user-1",
		Title:   "Test Notification",
		Message: "Test message",
		Type:    domain.NotificationTypeInfo,
	}

	user := &domain.User{
		ID:       "user-1",
		Name:     "Test User",
		IsActive: true,
	}

	mockUserRepo.On("GetByID", mock.Anything, "user-1").Return(user, nil)
	mockNotificationRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.Notification")).Return(nil)

	result, err := service.CreateNotification(context.Background(), req)

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Test Notification", result.Title)
	assert.Equal(t, "user-1", result.UserID)
	assert.False(t, result.IsRead)

	mockNotificationRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestNotificationService_CreateNotification_UserNotFound(t *testing.T) {
	mockNotificationRepo := new(MockNotificationRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewNotificationService(mockNotificationRepo, mockUserRepo)

	req := &domain.CreateNotificationRequest{
		UserID: "user-1",
		Title:  "Test Notification",
	}

	mockUserRepo.On("GetByID", mock.Anything, "user-1").Return(nil, errors.New("not found"))

	result, err := service.CreateNotification(context.Background(), req)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "user not found")

	mockUserRepo.AssertExpectations(t)
}

func TestNotificationService_CreateNotification_UserInactive(t *testing.T) {
	mockNotificationRepo := new(MockNotificationRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewNotificationService(mockNotificationRepo, mockUserRepo)

	req := &domain.CreateNotificationRequest{
		UserID: "user-1",
		Title:  "Test Notification",
	}

	user := &domain.User{
		ID:       "user-1",
		Name:     "Test User",
		IsActive: false,
	}

	mockUserRepo.On("GetByID", mock.Anything, "user-1").Return(user, nil)

	result, err := service.CreateNotification(context.Background(), req)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "user is not active")

	mockUserRepo.AssertExpectations(t)
}

func TestNotificationService_GetNotification_Success(t *testing.T) {
	mockNotificationRepo := new(MockNotificationRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewNotificationService(mockNotificationRepo, mockUserRepo)

	notification := &domain.Notification{
		ID:    "notification-1",
		Title: "Test Notification",
	}

	mockNotificationRepo.On("GetByID", mock.Anything, "notification-1").Return(notification, nil)

	result, err := service.GetNotification(context.Background(), "notification-1")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, notification.ID, result.ID)

	mockNotificationRepo.AssertExpectations(t)
}

func TestNotificationService_GetNotification_NotFound(t *testing.T) {
	mockNotificationRepo := new(MockNotificationRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewNotificationService(mockNotificationRepo, mockUserRepo)

	mockNotificationRepo.On("GetByID", mock.Anything, "notification-1").Return(nil, errors.New("not found"))

	result, err := service.GetNotification(context.Background(), "notification-1")

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "notification not found")

	mockNotificationRepo.AssertExpectations(t)
}

func TestNotificationService_ListNotifications_Success(t *testing.T) {
	mockNotificationRepo := new(MockNotificationRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewNotificationService(mockNotificationRepo, mockUserRepo)

	notifications := []*domain.Notification{
		{ID: "notification-1", Title: "Notification 1"},
		{ID: "notification-2", Title: "Notification 2"},
	}

	mockNotificationRepo.On("List", mock.Anything, (*string)(nil), (*string)(nil), (*bool)(nil), 10, 0).Return(notifications, nil)
	mockNotificationRepo.On("Count", mock.Anything, (*string)(nil), (*string)(nil), (*bool)(nil)).Return(2, nil)

	result, total, err := service.ListNotifications(context.Background(), nil, nil, nil, 1, 10)

	require.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, 2, total)

	mockNotificationRepo.AssertExpectations(t)
}

func TestNotificationService_MarkAsRead_Success(t *testing.T) {
	mockNotificationRepo := new(MockNotificationRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewNotificationService(mockNotificationRepo, mockUserRepo)

	notification := &domain.Notification{
		ID:     "notification-1",
		Title:  "Test Notification",
		IsRead: false,
	}

	mockNotificationRepo.On("GetByID", mock.Anything, "notification-1").Return(notification, nil)
	mockNotificationRepo.On("Update", mock.Anything, mock.AnythingOfType("*domain.Notification")).Return(nil)

	err := service.MarkAsRead(context.Background(), "notification-1")

	require.NoError(t, err)

	mockNotificationRepo.AssertExpectations(t)
}

func TestNotificationService_MarkAsRead_NotFound(t *testing.T) {
	mockNotificationRepo := new(MockNotificationRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewNotificationService(mockNotificationRepo, mockUserRepo)

	mockNotificationRepo.On("GetByID", mock.Anything, "notification-1").Return(nil, errors.New("not found"))

	err := service.MarkAsRead(context.Background(), "notification-1")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "notification not found")

	mockNotificationRepo.AssertExpectations(t)
}

func TestNotificationService_DeleteNotification_Success(t *testing.T) {
	mockNotificationRepo := new(MockNotificationRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewNotificationService(mockNotificationRepo, mockUserRepo)

	mockNotificationRepo.On("Delete", mock.Anything, "notification-1").Return(nil)

	err := service.DeleteNotification(context.Background(), "notification-1")

	require.NoError(t, err)

	mockNotificationRepo.AssertExpectations(t)
}

func TestNotificationService_GetUnreadCount_Success(t *testing.T) {
	mockNotificationRepo := new(MockNotificationRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewNotificationService(mockNotificationRepo, mockUserRepo)

	user := &domain.User{
		ID:   "user-1",
		Name: "Test User",
	}

	mockUserRepo.On("GetByID", mock.Anything, "user-1").Return(user, nil)
	mockNotificationRepo.On("GetUnreadCount", mock.Anything, "user-1").Return(5, nil)

	count, err := service.GetUnreadCount(context.Background(), "user-1")

	require.NoError(t, err)
	assert.Equal(t, 5, count)

	mockNotificationRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestNotificationService_GetUnreadCount_UserNotFound(t *testing.T) {
	mockNotificationRepo := new(MockNotificationRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewNotificationService(mockNotificationRepo, mockUserRepo)

	mockUserRepo.On("GetByID", mock.Anything, "user-1").Return(nil, errors.New("not found"))

	count, err := service.GetUnreadCount(context.Background(), "user-1")

	require.Error(t, err)
	assert.Equal(t, 0, count)
	assert.Contains(t, err.Error(), "user not found")

	mockUserRepo.AssertExpectations(t)
}

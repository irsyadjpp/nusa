package service

import (
	"context"
	"fmt"
	"time"

	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/repository"
)

// NotificationService handles business logic for notification operations
type NotificationService struct {
	notificationRepo repository.NotificationRepositoryInterface
	userRepo         repository.UserRepositoryInterface
}

// NewNotificationService creates a new notification service
func NewNotificationService(
	notificationRepo repository.NotificationRepositoryInterface,
	userRepo repository.UserRepositoryInterface,
) *NotificationService {
	return &NotificationService{
		notificationRepo: notificationRepo,
		userRepo:         userRepo,
	}
}

// CreateNotification creates a new notification
func (s *NotificationService) CreateNotification(ctx context.Context, req *domain.CreateNotificationRequest) (*domain.Notification, error) {
	// Verify user exists and is active
	user, err := s.userRepo.GetByID(ctx, req.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}
	if !user.IsActive {
		return nil, fmt.Errorf("user is not active")
	}

	notification := &domain.Notification{
		ID:        domain.NewID(),
		UserID:    req.UserID,
		Title:     req.Title,
		Message:   req.Message,
		Type:      string(req.Type),
		IsRead:    false,
		ActionURL: req.ActionURL,
		Metadata:  req.Metadata,
		CreatedAt: time.Now(),
	}

	if err := s.notificationRepo.Create(ctx, notification); err != nil {
		return nil, fmt.Errorf("failed to create notification: %v", err)
	}

	return notification, nil
}

// GetNotification retrieves a notification by ID
func (s *NotificationService) GetNotification(ctx context.Context, id string) (*domain.Notification, error) {
	notification, err := s.notificationRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("notification not found")
	}
	return notification, nil
}

// ListNotifications retrieves notifications with filters and pagination
func (s *NotificationService) ListNotifications(ctx context.Context, userID *string, notificationType *string, isRead *bool, page, pageSize int) ([]*domain.Notification, int, error) {
	limit := pageSize
	offset := (page - 1) * pageSize

	notifications, err := s.notificationRepo.List(ctx, userID, notificationType, isRead, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list notifications: %v", err)
	}

	total, err := s.notificationRepo.Count(ctx, userID, notificationType, isRead)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count notifications: %v", err)
	}

	return notifications, total, nil
}

// MarkAsRead marks a notification as read
func (s *NotificationService) MarkAsRead(ctx context.Context, id string) error {
	notification, err := s.notificationRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("notification not found")
	}

	notification.MarkAsRead()

	if err := s.notificationRepo.Update(ctx, notification); err != nil {
		return fmt.Errorf("failed to mark notification as read: %v", err)
	}

	return nil
}

// DeleteNotification soft deletes a notification
func (s *NotificationService) DeleteNotification(ctx context.Context, id string) error {
	return s.notificationRepo.Delete(ctx, id)
}

// GetUnreadCount returns the count of unread notifications for a user
func (s *NotificationService) GetUnreadCount(ctx context.Context, userID string) (int, error) {
	// Verify user exists
	_, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return 0, fmt.Errorf("user not found")
	}

	count, err := s.notificationRepo.GetUnreadCount(ctx, userID)
	if err != nil {
		return 0, fmt.Errorf("failed to get unread count: %v", err)
	}

	return count, nil
}

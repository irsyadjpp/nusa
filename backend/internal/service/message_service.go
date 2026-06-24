package service

import (
	"context"
	"fmt"
	"time"

	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/repository"
)

// MessageService handles business logic for message operations
type MessageService struct {
	messageRepo repository.MessageRepositoryInterface
	userRepo    repository.UserRepositoryInterface
}

// NewMessageService creates a new message service
func NewMessageService(
	messageRepo repository.MessageRepositoryInterface,
	userRepo repository.UserRepositoryInterface,
) *MessageService {
	return &MessageService{
		messageRepo: messageRepo,
		userRepo:    userRepo,
	}
}

// CreateMessage creates a new message
func (s *MessageService) CreateMessage(ctx context.Context, req *domain.CreateMessageRequest) (*domain.Message, error) {
	// Verify sender exists and is active
	sender, err := s.userRepo.GetByID(ctx, req.SenderID)
	if err != nil {
		return nil, fmt.Errorf("sender not found")
	}
	if !sender.IsActive {
		return nil, fmt.Errorf("sender is not active")
	}

	// Verify receiver exists and is active
	receiver, err := s.userRepo.GetByID(ctx, req.ReceiverID)
	if err != nil {
		return nil, fmt.Errorf("receiver not found")
	}
	if !receiver.IsActive {
		return nil, fmt.Errorf("receiver is not active")
	}

	// Verify sender and receiver are different
	if req.SenderID == req.ReceiverID {
		return nil, fmt.Errorf("sender and receiver cannot be the same")
	}

	message := &domain.Message{
		ID:              domain.NewID(),
		SenderID:        req.SenderID,
		ReceiverID:      req.ReceiverID,
		Subject:         req.Subject,
		Content:         req.Content,
		IsRead:          false,
		ParentMessageID: req.ParentMessageID,
		CreatedAt:       time.Now(),
	}

	if err := s.messageRepo.Create(ctx, message); err != nil {
		return nil, fmt.Errorf("failed to create message: %v", err)
	}

	return message, nil
}

// GetMessage retrieves a message by ID
func (s *MessageService) GetMessage(ctx context.Context, id string) (*domain.Message, error) {
	message, err := s.messageRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("message not found")
	}
	return message, nil
}

// ListMessages retrieves messages with filters and pagination
func (s *MessageService) ListMessages(ctx context.Context, senderID, receiverID *string, isRead *bool, page, pageSize int) ([]*domain.Message, int, error) {
	limit := pageSize
	offset := (page - 1) * pageSize

	messages, err := s.messageRepo.List(ctx, senderID, receiverID, isRead, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list messages: %v", err)
	}

	total, err := s.messageRepo.Count(ctx, senderID, receiverID, isRead)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count messages: %v", err)
	}

	return messages, total, nil
}

// MarkAsRead marks a message as read
func (s *MessageService) MarkAsRead(ctx context.Context, id string) error {
	message, err := s.messageRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("message not found")
	}

	message.MarkAsRead()

	if err := s.messageRepo.Update(ctx, message); err != nil {
		return fmt.Errorf("failed to mark message as read: %v", err)
	}

	return nil
}

// DeleteMessage soft deletes a message
func (s *MessageService) DeleteMessage(ctx context.Context, id string) error {
	return s.messageRepo.Delete(ctx, id)
}

// GetConversation retrieves messages between two users
func (s *MessageService) GetConversation(ctx context.Context, userID1, userID2 string, page, pageSize int) ([]*domain.Message, error) {
	// Verify both users exist
	_, err := s.userRepo.GetByID(ctx, userID1)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}
	_, err = s.userRepo.GetByID(ctx, userID2)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	limit := pageSize
	offset := (page - 1) * pageSize

	messages, err := s.messageRepo.GetConversation(ctx, userID1, userID2, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get conversation: %v", err)
	}

	return messages, nil
}

// GetUnreadCount returns the count of unread messages for a user
func (s *MessageService) GetUnreadCount(ctx context.Context, userID string) (int, error) {
	// Verify user exists
	_, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return 0, fmt.Errorf("user not found")
	}

	count, err := s.messageRepo.GetUnreadCount(ctx, userID)
	if err != nil {
		return 0, fmt.Errorf("failed to get unread count: %v", err)
	}

	return count, nil
}

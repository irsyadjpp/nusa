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

// MockMessageRepository is a mock implementation of MessageRepositoryInterface
type MockMessageRepository struct {
	mock.Mock
}

func (m *MockMessageRepository) Create(ctx context.Context, message *domain.Message) error {
	args := m.Called(ctx, message)
	return args.Error(0)
}

func (m *MockMessageRepository) GetByID(ctx context.Context, id string) (*domain.Message, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Message), args.Error(1)
}

func (m *MockMessageRepository) List(ctx context.Context, senderID, receiverID *string, isRead *bool, limit, offset int) ([]*domain.Message, error) {
	args := m.Called(ctx, senderID, receiverID, isRead, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Message), args.Error(1)
}

func (m *MockMessageRepository) Count(ctx context.Context, senderID, receiverID *string, isRead *bool) (int, error) {
	args := m.Called(ctx, senderID, receiverID, isRead)
	return args.Int(0), args.Error(1)
}

func (m *MockMessageRepository) Update(ctx context.Context, message *domain.Message) error {
	args := m.Called(ctx, message)
	return args.Error(0)
}

func (m *MockMessageRepository) MarkAsRead(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockMessageRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockMessageRepository) GetConversation(ctx context.Context, userID1, userID2 string, limit, offset int) ([]*domain.Message, error) {
	args := m.Called(ctx, userID1, userID2, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Message), args.Error(1)
}

func (m *MockMessageRepository) GetUnreadCount(ctx context.Context, userID string) (int, error) {
	args := m.Called(ctx, userID)
	return args.Int(0), args.Error(1)
}

func TestMessageService_CreateMessage_Success(t *testing.T) {
	mockMessageRepo := new(MockMessageRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewMessageService(mockMessageRepo, mockUserRepo)

	subject := "Test Subject"
	req := &domain.CreateMessageRequest{
		SenderID:   "user-1",
		ReceiverID: "user-2",
		Subject:    &subject,
		Content:    "Test content",
	}

	sender := &domain.User{
		ID:       "user-1",
		Name:     "Sender",
		IsActive: true,
	}

	receiver := &domain.User{
		ID:       "user-2",
		Name:     "Receiver",
		IsActive: true,
	}

	mockUserRepo.On("GetByID", mock.Anything, "user-1").Return(sender, nil)
	mockUserRepo.On("GetByID", mock.Anything, "user-2").Return(receiver, nil)
	mockMessageRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.Message")).Return(nil)

	result, err := service.CreateMessage(context.Background(), req)

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Test Subject", *result.Subject)
	assert.Equal(t, "user-1", result.SenderID)
	assert.Equal(t, "user-2", result.ReceiverID)

	mockMessageRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestMessageService_CreateMessage_SenderNotFound(t *testing.T) {
	mockMessageRepo := new(MockMessageRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewMessageService(mockMessageRepo, mockUserRepo)

	subject := "Test Subject"
	req := &domain.CreateMessageRequest{
		SenderID:   "user-1",
		ReceiverID: "user-2",
		Subject:    &subject,
	}

	mockUserRepo.On("GetByID", mock.Anything, "user-1").Return(nil, errors.New("not found"))

	result, err := service.CreateMessage(context.Background(), req)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "sender not found")

	mockUserRepo.AssertExpectations(t)
}

func TestMessageService_CreateMessage_ReceiverNotFound(t *testing.T) {
	mockMessageRepo := new(MockMessageRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewMessageService(mockMessageRepo, mockUserRepo)

	subject := "Test Subject"
	req := &domain.CreateMessageRequest{
		SenderID:   "user-1",
		ReceiverID: "user-2",
		Subject:    &subject,
	}

	sender := &domain.User{
		ID:       "user-1",
		Name:     "Sender",
		IsActive: true,
	}

	mockUserRepo.On("GetByID", mock.Anything, "user-1").Return(sender, nil)
	mockUserRepo.On("GetByID", mock.Anything, "user-2").Return(nil, errors.New("not found"))

	result, err := service.CreateMessage(context.Background(), req)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "receiver not found")

	mockUserRepo.AssertExpectations(t)
}

func TestMessageService_CreateMessage_SenderInactive(t *testing.T) {
	mockMessageRepo := new(MockMessageRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewMessageService(mockMessageRepo, mockUserRepo)

	subject := "Test Subject"
	req := &domain.CreateMessageRequest{
		SenderID:   "user-1",
		ReceiverID: "user-2",
		Subject:    &subject,
	}

	sender := &domain.User{
		ID:       "user-1",
		Name:     "Sender",
		IsActive: false,
	}

	mockUserRepo.On("GetByID", mock.Anything, "user-1").Return(sender, nil)

	result, err := service.CreateMessage(context.Background(), req)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "sender is not active")

	mockUserRepo.AssertExpectations(t)
}

func TestMessageService_CreateMessage_ReceiverInactive(t *testing.T) {
	mockMessageRepo := new(MockMessageRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewMessageService(mockMessageRepo, mockUserRepo)

	subject := "Test Subject"
	req := &domain.CreateMessageRequest{
		SenderID:   "user-1",
		ReceiverID: "user-2",
		Subject:    &subject,
	}

	sender := &domain.User{
		ID:       "user-1",
		Name:     "Sender",
		IsActive: true,
	}

	receiver := &domain.User{
		ID:       "user-2",
		Name:     "Receiver",
		IsActive: false,
	}

	mockUserRepo.On("GetByID", mock.Anything, "user-1").Return(sender, nil)
	mockUserRepo.On("GetByID", mock.Anything, "user-2").Return(receiver, nil)

	result, err := service.CreateMessage(context.Background(), req)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "receiver is not active")

	mockUserRepo.AssertExpectations(t)
}

func TestMessageService_CreateMessage_SameUser(t *testing.T) {
	mockMessageRepo := new(MockMessageRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewMessageService(mockMessageRepo, mockUserRepo)

	subject := "Test Subject"
	req := &domain.CreateMessageRequest{
		SenderID:   "user-1",
		ReceiverID: "user-1",
		Subject:    &subject,
	}

	mockUserRepo.On("GetByID", mock.Anything, "user-1").Return(&domain.User{ID: "user-1", IsActive: true}, nil)

	result, err := service.CreateMessage(context.Background(), req)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "sender and receiver cannot be the same")
}

func TestMessageService_GetMessage_Success(t *testing.T) {
	mockMessageRepo := new(MockMessageRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewMessageService(mockMessageRepo, mockUserRepo)

	subject := "Test Subject"
	message := &domain.Message{
		ID:      "message-1",
		Subject: &subject,
	}

	mockMessageRepo.On("GetByID", mock.Anything, "message-1").Return(message, nil)

	result, err := service.GetMessage(context.Background(), "message-1")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, message.ID, result.ID)

	mockMessageRepo.AssertExpectations(t)
}

func TestMessageService_GetMessage_NotFound(t *testing.T) {
	mockMessageRepo := new(MockMessageRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewMessageService(mockMessageRepo, mockUserRepo)

	mockMessageRepo.On("GetByID", mock.Anything, "message-1").Return(nil, errors.New("not found"))

	result, err := service.GetMessage(context.Background(), "message-1")

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "message not found")

	mockMessageRepo.AssertExpectations(t)
}

func TestMessageService_ListMessages_Success(t *testing.T) {
	mockMessageRepo := new(MockMessageRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewMessageService(mockMessageRepo, mockUserRepo)

	subject1 := "Message 1"
	subject2 := "Message 2"
	messages := []*domain.Message{
		{ID: "message-1", Subject: &subject1},
		{ID: "message-2", Subject: &subject2},
	}

	mockMessageRepo.On("List", mock.Anything, (*string)(nil), (*string)(nil), (*bool)(nil), 10, 0).Return(messages, nil)
	mockMessageRepo.On("Count", mock.Anything, (*string)(nil), (*string)(nil), (*bool)(nil)).Return(2, nil)

	result, total, err := service.ListMessages(context.Background(), nil, nil, nil, 1, 10)

	require.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, 2, total)

	mockMessageRepo.AssertExpectations(t)
}

func TestMessageService_MarkAsRead_Success(t *testing.T) {
	mockMessageRepo := new(MockMessageRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewMessageService(mockMessageRepo, mockUserRepo)

	subject := "Test Subject"
	message := &domain.Message{
		ID:      "message-1",
		Subject: &subject,
		IsRead:  false,
	}

	mockMessageRepo.On("GetByID", mock.Anything, "message-1").Return(message, nil)
	mockMessageRepo.On("Update", mock.Anything, mock.AnythingOfType("*domain.Message")).Return(nil)

	err := service.MarkAsRead(context.Background(), "message-1")

	require.NoError(t, err)

	mockMessageRepo.AssertExpectations(t)
}

func TestMessageService_MarkAsRead_NotFound(t *testing.T) {
	mockMessageRepo := new(MockMessageRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewMessageService(mockMessageRepo, mockUserRepo)

	mockMessageRepo.On("GetByID", mock.Anything, "message-1").Return(nil, errors.New("not found"))

	err := service.MarkAsRead(context.Background(), "message-1")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "message not found")

	mockMessageRepo.AssertExpectations(t)
}

func TestMessageService_DeleteMessage_Success(t *testing.T) {
	mockMessageRepo := new(MockMessageRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewMessageService(mockMessageRepo, mockUserRepo)

	mockMessageRepo.On("Delete", mock.Anything, "message-1").Return(nil)

	err := service.DeleteMessage(context.Background(), "message-1")

	require.NoError(t, err)

	mockMessageRepo.AssertExpectations(t)
}

func TestMessageService_GetConversation_Success(t *testing.T) {
	mockMessageRepo := new(MockMessageRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewMessageService(mockMessageRepo, mockUserRepo)

	user1 := &domain.User{ID: "user-1", Name: "User 1"}
	user2 := &domain.User{ID: "user-2", Name: "User 2"}

	subject := "Message 1"
	messages := []*domain.Message{
		{ID: "message-1", Subject: &subject},
	}

	mockUserRepo.On("GetByID", mock.Anything, "user-1").Return(user1, nil)
	mockUserRepo.On("GetByID", mock.Anything, "user-2").Return(user2, nil)
	mockMessageRepo.On("GetConversation", mock.Anything, "user-1", "user-2", 10, 0).Return(messages, nil)

	result, err := service.GetConversation(context.Background(), "user-1", "user-2", 1, 10)

	require.NoError(t, err)
	assert.Len(t, result, 1)

	mockMessageRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestMessageService_GetConversation_UserNotFound(t *testing.T) {
	mockMessageRepo := new(MockMessageRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewMessageService(mockMessageRepo, mockUserRepo)

	mockUserRepo.On("GetByID", mock.Anything, "user-1").Return(nil, errors.New("not found"))

	result, err := service.GetConversation(context.Background(), "user-1", "user-2", 1, 10)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "user not found")

	mockUserRepo.AssertExpectations(t)
}

func TestMessageService_GetUnreadCount_Success(t *testing.T) {
	mockMessageRepo := new(MockMessageRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewMessageService(mockMessageRepo, mockUserRepo)

	user := &domain.User{ID: "user-1", Name: "User 1"}

	mockUserRepo.On("GetByID", mock.Anything, "user-1").Return(user, nil)
	mockMessageRepo.On("GetUnreadCount", mock.Anything, "user-1").Return(5, nil)

	count, err := service.GetUnreadCount(context.Background(), "user-1")

	require.NoError(t, err)
	assert.Equal(t, 5, count)

	mockMessageRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestMessageService_GetUnreadCount_UserNotFound(t *testing.T) {
	mockMessageRepo := new(MockMessageRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewMessageService(mockMessageRepo, mockUserRepo)

	mockUserRepo.On("GetByID", mock.Anything, "user-1").Return(nil, errors.New("not found"))

	count, err := service.GetUnreadCount(context.Background(), "user-1")

	require.Error(t, err)
	assert.Equal(t, 0, count)
	assert.Contains(t, err.Error(), "user not found")

	mockUserRepo.AssertExpectations(t)
}

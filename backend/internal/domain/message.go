package domain

import (
	"fmt"
	"time"
)

// Message represents a user message
type Message struct {
	ID              string     `json:"id" db:"id"`
	SenderID        string     `json:"sender_id" db:"sender_id"`
	ReceiverID      string     `json:"receiver_id" db:"receiver_id"`
	Subject         *string    `json:"subject,omitempty" db:"subject"`
	Content         string     `json:"content" db:"content"`
	IsRead          bool       `json:"is_read" db:"is_read"`
	ReadAt          *time.Time `json:"read_at,omitempty" db:"read_at"`
	ParentMessageID *string    `json:"parent_message_id,omitempty" db:"parent_message_id"`
	CreatedAt       time.Time  `json:"created_at" db:"created_at"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

// CreateMessageRequest represents the request to create a message
type CreateMessageRequest struct {
	SenderID        string  `json:"sender_id" binding:"required"`
	ReceiverID      string  `json:"receiver_id" binding:"required"`
	Subject         *string `json:"subject,omitempty" binding:"omitempty,max=255"`
	Content         string  `json:"content" binding:"required"`
	ParentMessageID *string `json:"parent_message_id,omitempty"`
}

// UpdateMessageRequest represents the request to update a message
type UpdateMessageRequest struct {
	IsRead *bool `json:"is_read,omitempty"`
}

// MessageResponse represents the message data returned to clients
type MessageResponse struct {
	ID              string  `json:"id"`
	SenderID        string  `json:"sender_id"`
	SenderName      *string `json:"sender_name,omitempty"`
	ReceiverID      string  `json:"receiver_id"`
	ReceiverName    *string `json:"receiver_name,omitempty"`
	Subject         *string `json:"subject,omitempty"`
	Content         string  `json:"content"`
	IsRead          bool    `json:"is_read"`
	ReadAt          *string `json:"read_at,omitempty"`
	ParentMessageID *string `json:"parent_message_id,omitempty"`
	CreatedAt       string  `json:"created_at"`
}

// ToMessageResponse converts Message to MessageResponse
func (m *Message) ToMessageResponse(senderName, receiverName string) *MessageResponse {
	var senderNamePtr, receiverNamePtr, subjectPtr, readAtPtr, parentMessageIDPtr *string

	if senderName != "" {
		senderNamePtr = &senderName
	}
	if receiverName != "" {
		receiverNamePtr = &receiverName
	}
	if m.Subject != nil {
		subjectPtr = m.Subject
	}
	if m.ReadAt != nil {
		readAtStr := m.ReadAt.Format(time.RFC3339)
		readAtPtr = &readAtStr
	}
	if m.ParentMessageID != nil {
		parentMessageIDPtr = m.ParentMessageID
	}

	return &MessageResponse{
		ID:              m.ID,
		SenderID:        m.SenderID,
		SenderName:      senderNamePtr,
		ReceiverID:      m.ReceiverID,
		ReceiverName:    receiverNamePtr,
		Subject:         subjectPtr,
		Content:         m.Content,
		IsRead:          m.IsRead,
		ReadAt:          readAtPtr,
		ParentMessageID: parentMessageIDPtr,
		CreatedAt:       m.CreatedAt.Format(time.RFC3339),
	}
}

// MarkAsRead marks the message as read
func (m *Message) MarkAsRead() {
	now := time.Now()
	m.IsRead = true
	m.ReadAt = &now
}

// Validate validates the message entity
func (m *Message) Validate() error {
	if m.ID == "" {
		return fmt.Errorf("id is required")
	}
	if m.SenderID == "" {
		return fmt.Errorf("sender_id is required")
	}
	if m.ReceiverID == "" {
		return fmt.Errorf("receiver_id is required")
	}
	if m.SenderID == m.ReceiverID {
		return fmt.Errorf("sender_id and receiver_id cannot be the same")
	}
	if m.Content == "" {
		return fmt.Errorf("content is required")
	}
	return nil
}

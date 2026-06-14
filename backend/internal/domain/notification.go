package domain

import (
	"encoding/json"
	"fmt"
	"time"
)

// Notification represents a user notification
type Notification struct {
	ID        string          `json:"id" db:"id"`
	UserID    string          `json:"user_id" db:"user_id"`
	Title     string          `json:"title" db:"title"`
	Message   string          `json:"message" db:"message"`
	Type      string          `json:"type" db:"type"`
	IsRead    bool            `json:"is_read" db:"is_read"`
	ReadAt    *time.Time      `json:"read_at,omitempty" db:"read_at"`
	ActionURL *string         `json:"action_url,omitempty" db:"action_url"`
	Metadata  json.RawMessage `json:"metadata,omitempty" db:"metadata"`
	CreatedAt time.Time       `json:"created_at" db:"created_at"`
	DeletedAt *time.Time      `json:"deleted_at,omitempty" db:"deleted_at"`
}

// NotificationType represents the type of notification
type NotificationType string

const (
	NotificationTypeInfo    NotificationType = "INFO"
	NotificationTypeWarning NotificationType = "WARNING"
	NotificationTypeError   NotificationType = "ERROR"
	NotificationTypeSuccess NotificationType = "SUCCESS"
)

// CreateNotificationRequest represents the request to create a notification
type CreateNotificationRequest struct {
	UserID    string          `json:"user_id" binding:"required"`
	Title     string          `json:"title" binding:"required,max=255"`
	Message   string          `json:"message" binding:"required"`
	Type      NotificationType `json:"type" binding:"required,oneof=INFO WARNING ERROR SUCCESS"`
	ActionURL *string         `json:"action_url,omitempty" binding:"omitempty,max=500"`
	Metadata  json.RawMessage `json:"metadata,omitempty"`
}

// UpdateNotificationRequest represents the request to update a notification
type UpdateNotificationRequest struct {
	IsRead *bool `json:"is_read,omitempty"`
}

// NotificationResponse represents the notification data returned to clients
type NotificationResponse struct {
	ID        string          `json:"id"`
	UserID    string          `json:"user_id"`
	UserName  *string         `json:"user_name,omitempty"`
	Title     string          `json:"title"`
	Message   string          `json:"message"`
	Type      NotificationType `json:"type"`
	IsRead    bool            `json:"is_read"`
	ReadAt    *string         `json:"read_at,omitempty"`
	ActionURL *string         `json:"action_url,omitempty"`
	Metadata  json.RawMessage `json:"metadata,omitempty"`
	CreatedAt string          `json:"created_at"`
}

// ToNotificationResponse converts Notification to NotificationResponse
func (n *Notification) ToNotificationResponse(userName string) *NotificationResponse {
	var userNamePtr, readAtPtr, actionURLPtr *string

	if userName != "" {
		userNamePtr = &userName
	}
	if n.ReadAt != nil {
		readAtStr := n.ReadAt.Format(time.RFC3339)
		readAtPtr = &readAtStr
	}
	if n.ActionURL != nil {
		actionURLPtr = n.ActionURL
	}

	return &NotificationResponse{
		ID:        n.ID,
		UserID:    n.UserID,
		UserName:  userNamePtr,
		Title:     n.Title,
		Message:   n.Message,
		Type:      NotificationType(n.Type),
		IsRead:    n.IsRead,
		ReadAt:    readAtPtr,
		ActionURL: actionURLPtr,
		Metadata:  n.Metadata,
		CreatedAt: n.CreatedAt.Format(time.RFC3339),
	}
}

// MarkAsRead marks the notification as read
func (n *Notification) MarkAsRead() {
	now := time.Now()
	n.IsRead = true
	n.ReadAt = &now
}

// Validate validates the notification entity
func (n *Notification) Validate() error {
	if n.ID == "" {
		return fmt.Errorf("id is required")
	}
	if n.UserID == "" {
		return fmt.Errorf("user_id is required")
	}
	if n.Title == "" {
		return fmt.Errorf("title is required")
	}
	if n.Message == "" {
		return fmt.Errorf("message is required")
	}
	if n.Type == "" {
		return fmt.Errorf("type is required")
	}
	validTypes := map[string]bool{
		string(NotificationTypeInfo):    true,
		string(NotificationTypeWarning): true,
		string(NotificationTypeError):   true,
		string(NotificationTypeSuccess): true,
	}
	if !validTypes[n.Type] {
		return fmt.Errorf("invalid type: %s", n.Type)
	}
	return nil
}

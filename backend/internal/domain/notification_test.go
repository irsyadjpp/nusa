package domain

import (
	"encoding/json"
	"testing"
	"time"
)

// TestNotification_ToNotificationResponse tests conversion to response format
func TestNotification_ToNotificationResponse(t *testing.T) {
	now := time.Now()
	actionURL := "https://example.com/action"
	readAt := now.Add(1 * time.Hour)
	metadata := json.RawMessage(`{"key": "value"}`)

	tests := []struct {
		name            string
		notification    *Notification
		userName        string
		expectedUserName *string
		expectedReadAt  *string
		expectedActionURL *string
	}{
		{
			name: "Full notification with all related data",
			notification: &Notification{
				ID:        "notif1",
				UserID:    "user1",
				Title:     "New Assignment",
				Message:   "You have a new assignment",
				Type:      "INFO",
				IsRead:    true,
				ReadAt:    &readAt,
				ActionURL: &actionURL,
				Metadata:  metadata,
				CreatedAt: now,
			},
			userName:        "John Doe",
			expectedUserName: stringPtr("John Doe"),
			expectedReadAt:  stringPtr(readAt.Format(time.RFC3339)),
			expectedActionURL: &actionURL,
		},
		{
			name: "Notification without user name",
			notification: &Notification{
				ID:        "notif2",
				UserID:    "user1",
				Title:     "Test",
				Message:   "Test message",
				Type:      "WARNING",
				IsRead:    false,
				ReadAt:    nil,
				ActionURL: nil,
				Metadata:  nil,
				CreatedAt: now,
			},
			userName:        "",
			expectedUserName: nil,
			expectedReadAt:  nil,
			expectedActionURL: nil,
		},
		{
			name: "Unread notification without read_at",
			notification: &Notification{
				ID:        "notif3",
				UserID:    "user1",
				Title:     "Test",
				Message:   "Test message",
				Type:      "ERROR",
				IsRead:    false,
				ReadAt:    nil,
				ActionURL: &actionURL,
				Metadata:  metadata,
				CreatedAt: now,
			},
			userName:        "John Doe",
			expectedUserName: stringPtr("John Doe"),
			expectedReadAt:  nil,
			expectedActionURL: &actionURL,
		},
		{
			name: "Notification without action URL",
			notification: &Notification{
				ID:        "notif4",
				UserID:    "user1",
				Title:     "Test",
				Message:   "Test message",
				Type:      "SUCCESS",
				IsRead:    true,
				ReadAt:    &readAt,
				ActionURL: nil,
				Metadata:  nil,
				CreatedAt: now,
			},
			userName:        "John Doe",
			expectedUserName: stringPtr("John Doe"),
			expectedReadAt:  stringPtr(readAt.Format(time.RFC3339)),
			expectedActionURL: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			resp := tc.notification.ToNotificationResponse(tc.userName)

			if resp.ID != tc.notification.ID {
				t.Errorf("expected ID %s, got %s", tc.notification.ID, resp.ID)
			}
			if resp.UserID != tc.notification.UserID {
				t.Errorf("expected UserID %s, got %s", tc.notification.UserID, resp.UserID)
			}
			if resp.Title != tc.notification.Title {
				t.Errorf("expected Title %s, got %s", tc.notification.Title, resp.Title)
			}
			if resp.Message != tc.notification.Message {
				t.Errorf("expected Message %s, got %s", tc.notification.Message, resp.Message)
			}
			if resp.Type != NotificationType(tc.notification.Type) {
				t.Errorf("expected Type %s, got %s", tc.notification.Type, resp.Type)
			}
			if resp.IsRead != tc.notification.IsRead {
				t.Errorf("expected IsRead %v, got %v", tc.notification.IsRead, resp.IsRead)
			}

			// Check optional fields
			if (resp.UserName == nil) != (tc.expectedUserName == nil) {
				t.Errorf("UserName pointer mismatch")
			} else if resp.UserName != nil && tc.expectedUserName != nil {
				if *resp.UserName != *tc.expectedUserName {
					t.Errorf("expected UserName %s, got %s", *tc.expectedUserName, *resp.UserName)
				}
			}

			if (resp.ReadAt == nil) != (tc.expectedReadAt == nil) {
				t.Errorf("ReadAt pointer mismatch")
			} else if resp.ReadAt != nil && tc.expectedReadAt != nil {
				if *resp.ReadAt != *tc.expectedReadAt {
					t.Errorf("expected ReadAt %s, got %s", *tc.expectedReadAt, *resp.ReadAt)
				}
			}

			if (resp.ActionURL == nil) != (tc.expectedActionURL == nil) {
				t.Errorf("ActionURL pointer mismatch")
			} else if resp.ActionURL != nil && tc.expectedActionURL != nil {
				if *resp.ActionURL != *tc.expectedActionURL {
					t.Errorf("expected ActionURL %s, got %s", *tc.expectedActionURL, *resp.ActionURL)
				}
			}
		})
	}
}

// TestNotification_MarkAsRead tests marking notification as read
func TestNotification_MarkAsRead(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name       string
		notification *Notification
		expectedRead bool
	}{
		{
			name: "Mark unread notification as read",
			notification: &Notification{
				ID:        "notif1",
				UserID:    "user1",
				Title:     "Test",
				Message:   "Test message",
				Type:      "INFO",
				IsRead:    false,
				ReadAt:    nil,
				CreatedAt: now,
			},
			expectedRead: true,
		},
		{
			name: "Mark already read notification",
			notification: &Notification{
				ID:        "notif2",
				UserID:    "user1",
				Title:     "Test",
				Message:   "Test message",
				Type:      "WARNING",
				IsRead:    true,
				ReadAt:    timePtr(now.Add(-1 * time.Hour)),
				CreatedAt: now,
			},
			expectedRead: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.notification.MarkAsRead()

			if !tc.notification.IsRead {
				t.Errorf("expected IsRead to be true")
			}
			if tc.notification.ReadAt == nil {
				t.Errorf("expected ReadAt to be set")
			}
			if tc.notification.ReadAt.Before(tc.notification.CreatedAt) {
				t.Errorf("expected ReadAt to be after CreatedAt")
			}
		})
	}
}

// TestNotification_Validate tests notification validation
func TestNotification_Validate(t *testing.T) {
	now := time.Now()
	actionURL := "https://example.com/action"
	metadata := json.RawMessage(`{"key": "value"}`)

	tests := []struct {
		name        string
		notification *Notification
		expectedErr string
	}{
		{
			name: "Valid notification - INFO",
			notification: &Notification{
				ID:        "notif1",
				UserID:    "user1",
				Title:     "New Assignment",
				Message:   "You have a new assignment",
				Type:      "INFO",
				ActionURL: &actionURL,
				Metadata:  metadata,
				CreatedAt: now,
			},
			expectedErr: "",
		},
		{
			name: "Valid notification - WARNING",
			notification: &Notification{
				ID:        "notif2",
				UserID:    "user1",
				Title:     "Test",
				Message:   "Test message",
				Type:      "WARNING",
				CreatedAt: now,
			},
			expectedErr: "",
		},
		{
			name: "Valid notification - ERROR",
			notification: &Notification{
				ID:        "notif3",
				UserID:    "user1",
				Title:     "Test",
				Message:   "Test message",
				Type:      "ERROR",
				CreatedAt: now,
			},
			expectedErr: "",
		},
		{
			name: "Valid notification - SUCCESS",
			notification: &Notification{
				ID:        "notif4",
				UserID:    "user1",
				Title:     "Test",
				Message:   "Test message",
				Type:      "SUCCESS",
				CreatedAt: now,
			},
			expectedErr: "",
		},
		{
			name: "Valid notification without action URL and metadata",
			notification: &Notification{
				ID:        "notif5",
				UserID:    "user1",
				Title:     "Test",
				Message:   "Test message",
				Type:      "INFO",
				ActionURL: nil,
				Metadata:  nil,
				CreatedAt: now,
			},
			expectedErr: "",
		},
		{
			name: "Invalid - missing ID",
			notification: &Notification{
				ID:        "",
				UserID:    "user1",
				Title:     "Test",
				Message:   "Test message",
				Type:      "INFO",
				CreatedAt: now,
			},
			expectedErr: "id is required",
		},
		{
			name: "Invalid - missing UserID",
			notification: &Notification{
				ID:        "notif1",
				UserID:    "",
				Title:     "Test",
				Message:   "Test message",
				Type:      "INFO",
				CreatedAt: now,
			},
			expectedErr: "user_id is required",
		},
		{
			name: "Invalid - missing Title",
			notification: &Notification{
				ID:        "notif1",
				UserID:    "user1",
				Title:     "",
				Message:   "Test message",
				Type:      "INFO",
				CreatedAt: now,
			},
			expectedErr: "title is required",
		},
		{
			name: "Invalid - missing Message",
			notification: &Notification{
				ID:        "notif1",
				UserID:    "user1",
				Title:     "Test",
				Message:   "",
				Type:      "INFO",
				CreatedAt: now,
			},
			expectedErr: "message is required",
		},
		{
			name: "Invalid - missing Type",
			notification: &Notification{
				ID:        "notif1",
				UserID:    "user1",
				Title:     "Test",
				Message:   "Test message",
				Type:      "",
				CreatedAt: now,
			},
			expectedErr: "type is required",
		},
		{
			name: "Invalid - invalid Type",
			notification: &Notification{
				ID:        "notif1",
				UserID:    "user1",
				Title:     "Test",
				Message:   "Test message",
				Type:      "INVALID",
				CreatedAt: now,
			},
			expectedErr: "invalid type: INVALID",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.notification.Validate()
			if tc.expectedErr == "" && err != nil {
				t.Errorf("expected no error, got: %v", err)
			}
			if tc.expectedErr != "" && err == nil {
				t.Errorf("expected error: %s, got nil", tc.expectedErr)
			}
			if tc.expectedErr != "" && err != nil && err.Error() != tc.expectedErr {
				t.Errorf("expected error: %s, got: %s", tc.expectedErr, err.Error())
			}
		})
	}
}

package domain

import (
	"testing"
	"time"
)

// TestMessage_ToMessageResponse tests conversion to response format
func TestMessage_ToMessageResponse(t *testing.T) {
	now := time.Now()
	subject := "Test Subject"
	parentMessageID := "parent123"
	readAt := now.Add(1 * time.Hour)

	tests := []struct {
		name            string
		message         *Message
		senderName      string
		receiverName    string
		expectedSubject *string
		expectedReadAt  *string
		expectedParentMessageID *string
	}{
		{
			name: "Full message with all related data",
			message: &Message{
				ID:              "msg1",
				SenderID:        "user1",
				ReceiverID:      "user2",
				Subject:         &subject,
				Content:         "Hello, how are you?",
				IsRead:          true,
				ReadAt:          &readAt,
				ParentMessageID: &parentMessageID,
				CreatedAt:       now,
			},
			senderName:      "John Doe",
			receiverName:    "Jane Smith",
			expectedSubject: &subject,
			expectedReadAt:  stringPtr(readAt.Format(time.RFC3339)),
			expectedParentMessageID: &parentMessageID,
		},
		{
			name: "Message without subject",
			message: &Message{
				ID:              "msg2",
				SenderID:        "user1",
				ReceiverID:      "user2",
				Subject:         nil,
				Content:         "Hello",
				IsRead:          false,
				ReadAt:          nil,
				ParentMessageID: nil,
				CreatedAt:       now,
			},
			senderName:      "John Doe",
			receiverName:    "Jane Smith",
			expectedSubject: nil,
			expectedReadAt:  nil,
			expectedParentMessageID: nil,
		},
		{
			name: "Message without sender name",
			message: &Message{
				ID:              "msg3",
				SenderID:        "user1",
				ReceiverID:      "user2",
				Subject:         &subject,
				Content:         "Hello",
				IsRead:          true,
				ReadAt:          &readAt,
				ParentMessageID: nil,
				CreatedAt:       now,
			},
			senderName:      "",
			receiverName:    "Jane Smith",
			expectedSubject: &subject,
			expectedReadAt:  stringPtr(readAt.Format(time.RFC3339)),
			expectedParentMessageID: nil,
		},
		{
			name: "Message without receiver name",
			message: &Message{
				ID:              "msg4",
				SenderID:        "user1",
				ReceiverID:      "user2",
				Subject:         &subject,
				Content:         "Hello",
				IsRead:          false,
				ReadAt:          nil,
				ParentMessageID: &parentMessageID,
				CreatedAt:       now,
			},
			senderName:      "John Doe",
			receiverName:    "",
			expectedSubject: &subject,
			expectedReadAt:  nil,
			expectedParentMessageID: &parentMessageID,
		},
		{
			name: "Unread message without read_at",
			message: &Message{
				ID:              "msg5",
				SenderID:        "user1",
				ReceiverID:      "user2",
				Subject:         &subject,
				Content:         "Hello",
				IsRead:          false,
				ReadAt:          nil,
				ParentMessageID: nil,
				CreatedAt:       now,
			},
			senderName:      "John Doe",
			receiverName:    "Jane Smith",
			expectedSubject: &subject,
			expectedReadAt:  nil,
			expectedParentMessageID: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			resp := tc.message.ToMessageResponse(tc.senderName, tc.receiverName)

			if resp.ID != tc.message.ID {
				t.Errorf("expected ID %s, got %s", tc.message.ID, resp.ID)
			}
			if resp.SenderID != tc.message.SenderID {
				t.Errorf("expected SenderID %s, got %s", tc.message.SenderID, resp.SenderID)
			}
			if resp.ReceiverID != tc.message.ReceiverID {
				t.Errorf("expected ReceiverID %s, got %s", tc.message.ReceiverID, resp.ReceiverID)
			}
			if resp.Content != tc.message.Content {
				t.Errorf("expected Content %s, got %s", tc.message.Content, resp.Content)
			}
			if resp.IsRead != tc.message.IsRead {
				t.Errorf("expected IsRead %v, got %v", tc.message.IsRead, resp.IsRead)
			}

			// Check optional fields
			if (resp.SenderName == nil) != (tc.senderName == "") {
				t.Errorf("SenderName pointer mismatch")
			} else if resp.SenderName != nil && tc.senderName != "" {
				if *resp.SenderName != tc.senderName {
					t.Errorf("expected SenderName %s, got %s", tc.senderName, *resp.SenderName)
				}
			}

			if (resp.ReceiverName == nil) != (tc.receiverName == "") {
				t.Errorf("ReceiverName pointer mismatch")
			} else if resp.ReceiverName != nil && tc.receiverName != "" {
				if *resp.ReceiverName != tc.receiverName {
					t.Errorf("expected ReceiverName %s, got %s", tc.receiverName, *resp.ReceiverName)
				}
			}

			if (resp.Subject == nil) != (tc.expectedSubject == nil) {
				t.Errorf("Subject pointer mismatch")
			} else if resp.Subject != nil && tc.expectedSubject != nil {
				if *resp.Subject != *tc.expectedSubject {
					t.Errorf("expected Subject %s, got %s", *tc.expectedSubject, *resp.Subject)
				}
			}

			if (resp.ReadAt == nil) != (tc.expectedReadAt == nil) {
				t.Errorf("ReadAt pointer mismatch")
			} else if resp.ReadAt != nil && tc.expectedReadAt != nil {
				if *resp.ReadAt != *tc.expectedReadAt {
					t.Errorf("expected ReadAt %s, got %s", *tc.expectedReadAt, *resp.ReadAt)
				}
			}

			if (resp.ParentMessageID == nil) != (tc.expectedParentMessageID == nil) {
				t.Errorf("ParentMessageID pointer mismatch")
			} else if resp.ParentMessageID != nil && tc.expectedParentMessageID != nil {
				if *resp.ParentMessageID != *tc.expectedParentMessageID {
					t.Errorf("expected ParentMessageID %s, got %s", *tc.expectedParentMessageID, *resp.ParentMessageID)
				}
			}
		})
	}
}

// TestMessage_MarkAsRead tests marking message as read
func TestMessage_MarkAsRead(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name       string
		message    *Message
		expectedRead bool
	}{
		{
			name: "Mark unread message as read",
			message: &Message{
				ID:         "msg1",
				SenderID:   "user1",
				ReceiverID: "user2",
				Content:    "Hello",
				IsRead:     false,
				ReadAt:     nil,
				CreatedAt:  now,
			},
			expectedRead: true,
		},
		{
			name: "Mark already read message",
			message: &Message{
				ID:         "msg2",
				SenderID:   "user1",
				ReceiverID: "user2",
				Content:    "Hello",
				IsRead:     true,
				ReadAt:     timePtr(now.Add(-1 * time.Hour)),
				CreatedAt:  now,
			},
			expectedRead: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.message.MarkAsRead()

			if !tc.message.IsRead {
				t.Errorf("expected IsRead to be true")
			}
			if tc.message.ReadAt == nil {
				t.Errorf("expected ReadAt to be set")
			}
			if tc.message.ReadAt.Before(tc.message.CreatedAt) {
				t.Errorf("expected ReadAt to be after CreatedAt")
			}
		})
	}
}

// TestMessage_Validate tests message validation
func TestMessage_Validate(t *testing.T) {
	now := time.Now()
	subject := "Test Subject"

	tests := []struct {
		name        string
		message     *Message
		expectedErr string
	}{
		{
			name: "Valid message",
			message: &Message{
				ID:         "msg1",
				SenderID:   "user1",
				ReceiverID: "user2",
				Subject:    &subject,
				Content:    "Hello, how are you?",
				IsRead:     false,
				CreatedAt:  now,
			},
			expectedErr: "",
		},
		{
			name: "Valid message without subject",
			message: &Message{
				ID:         "msg2",
				SenderID:   "user1",
				ReceiverID: "user2",
				Subject:    nil,
				Content:    "Hello",
				IsRead:     false,
				CreatedAt:  now,
			},
			expectedErr: "",
		},
		{
			name: "Invalid - missing ID",
			message: &Message{
				ID:         "",
				SenderID:   "user1",
				ReceiverID: "user2",
				Content:    "Hello",
				CreatedAt:  now,
			},
			expectedErr: "id is required",
		},
		{
			name: "Invalid - missing SenderID",
			message: &Message{
				ID:         "msg1",
				SenderID:   "",
				ReceiverID: "user2",
				Content:    "Hello",
				CreatedAt:  now,
			},
			expectedErr: "sender_id is required",
		},
		{
			name: "Invalid - missing ReceiverID",
			message: &Message{
				ID:         "msg1",
				SenderID:   "user1",
				ReceiverID: "",
				Content:    "Hello",
				CreatedAt:  now,
			},
			expectedErr: "receiver_id is required",
		},
		{
			name: "Invalid - sender and receiver are the same",
			message: &Message{
				ID:         "msg1",
				SenderID:   "user1",
				ReceiverID: "user1",
				Content:    "Hello",
				CreatedAt:  now,
			},
			expectedErr: "sender_id and receiver_id cannot be the same",
		},
		{
			name: "Invalid - missing Content",
			message: &Message{
				ID:         "msg1",
				SenderID:   "user1",
				ReceiverID: "user2",
				Content:    "",
				CreatedAt:  now,
			},
			expectedErr: "content is required",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.message.Validate()
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

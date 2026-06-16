package domain

import (
	"testing"
	"time"
)

// TestAnnouncement_ToAnnouncementResponse tests conversion to response format
func TestAnnouncement_ToAnnouncementResponse(t *testing.T) {
	now := time.Now()
	future := now.Add(24 * time.Hour)

	tests := []struct {
		name                    string
		announcement            *Announcement
		schoolName              string
		publishedByName         string
		expectedSchoolName      *string
		expectedPublishedByName *string
		expectedExpiresAt       *string
	}{
		{
			name: "Full announcement with all related data",
			announcement: &Announcement{
				ID:             "ann1",
				SchoolID:       "school1",
				Title:          "School Holiday",
				Content:        "School will be closed",
				Priority:       "HIGH",
				TargetAudience: "ALL",
				PublishedBy:    "user1",
				PublishedAt:    now,
				ExpiresAt:      &future,
				IsActive:       true,
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			schoolName:              "Sekolah Indonesia",
			publishedByName:         "Admin User",
			expectedSchoolName:      stringPtr("Sekolah Indonesia"),
			expectedPublishedByName: stringPtr("Admin User"),
			expectedExpiresAt:       stringPtr(future.Format(time.RFC3339)),
		},
		{
			name: "Announcement without expires_at",
			announcement: &Announcement{
				ID:             "ann2",
				SchoolID:       "school1",
				Title:          "Meeting",
				Content:        "Staff meeting at 10am",
				Priority:       "NORMAL",
				TargetAudience: "TEACHERS",
				PublishedBy:    "user1",
				PublishedAt:    now,
				ExpiresAt:      nil,
				IsActive:       true,
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			schoolName:              "Sekolah Indonesia",
			publishedByName:         "Admin User",
			expectedSchoolName:      stringPtr("Sekolah Indonesia"),
			expectedPublishedByName: stringPtr("Admin User"),
			expectedExpiresAt:       nil,
		},
		{
			name: "Announcement without school name",
			announcement: &Announcement{
				ID:             "ann3",
				SchoolID:       "school1",
				Title:          "Test",
				Content:        "Test content",
				Priority:       "LOW",
				TargetAudience: "ADMIN",
				PublishedBy:    "user1",
				PublishedAt:    now,
				ExpiresAt:      nil,
				IsActive:       true,
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			schoolName:              "",
			publishedByName:         "Admin User",
			expectedSchoolName:      nil,
			expectedPublishedByName: stringPtr("Admin User"),
			expectedExpiresAt:       nil,
		},
		{
			name: "Announcement without published by name",
			announcement: &Announcement{
				ID:             "ann4",
				SchoolID:       "school1",
				Title:          "Test",
				Content:        "Test content",
				Priority:       "URGENT",
				TargetAudience: "STUDENTS",
				PublishedBy:    "user1",
				PublishedAt:    now,
				ExpiresAt:      &future,
				IsActive:       true,
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			schoolName:              "Sekolah Indonesia",
			publishedByName:         "",
			expectedSchoolName:      stringPtr("Sekolah Indonesia"),
			expectedPublishedByName: nil,
			expectedExpiresAt:       stringPtr(future.Format(time.RFC3339)),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			resp := tc.announcement.ToAnnouncementResponse(tc.schoolName, tc.publishedByName)

			if resp.ID != tc.announcement.ID {
				t.Errorf("expected ID %s, got %s", tc.announcement.ID, resp.ID)
			}
			if resp.SchoolID != tc.announcement.SchoolID {
				t.Errorf("expected SchoolID %s, got %s", tc.announcement.SchoolID, resp.SchoolID)
			}
			if resp.Title != tc.announcement.Title {
				t.Errorf("expected Title %s, got %s", tc.announcement.Title, resp.Title)
			}
			if resp.Content != tc.announcement.Content {
				t.Errorf("expected Content %s, got %s", tc.announcement.Content, resp.Content)
			}
			if resp.Priority != AnnouncementPriority(tc.announcement.Priority) {
				t.Errorf("expected Priority %s, got %s", tc.announcement.Priority, resp.Priority)
			}
			if resp.TargetAudience != TargetAudience(tc.announcement.TargetAudience) {
				t.Errorf("expected TargetAudience %s, got %s", tc.announcement.TargetAudience, resp.TargetAudience)
			}
			if resp.PublishedBy != tc.announcement.PublishedBy {
				t.Errorf("expected PublishedBy %s, got %s", tc.announcement.PublishedBy, resp.PublishedBy)
			}
			if resp.IsActive != tc.announcement.IsActive {
				t.Errorf("expected IsActive %v, got %v", tc.announcement.IsActive, resp.IsActive)
			}

			// Check optional fields
			if (resp.SchoolName == nil) != (tc.expectedSchoolName == nil) {
				t.Errorf("SchoolName pointer mismatch")
			} else if resp.SchoolName != nil && tc.expectedSchoolName != nil {
				if *resp.SchoolName != *tc.expectedSchoolName {
					t.Errorf("expected SchoolName %s, got %s", *tc.expectedSchoolName, *resp.SchoolName)
				}
			}

			if (resp.PublishedByName == nil) != (tc.expectedPublishedByName == nil) {
				t.Errorf("PublishedByName pointer mismatch")
			} else if resp.PublishedByName != nil && tc.expectedPublishedByName != nil {
				if *resp.PublishedByName != *tc.expectedPublishedByName {
					t.Errorf("expected PublishedByName %s, got %s", *tc.expectedPublishedByName, *resp.PublishedByName)
				}
			}

			if (resp.ExpiresAt == nil) != (tc.expectedExpiresAt == nil) {
				t.Errorf("ExpiresAt pointer mismatch")
			} else if resp.ExpiresAt != nil && tc.expectedExpiresAt != nil {
				if *resp.ExpiresAt != *tc.expectedExpiresAt {
					t.Errorf("expected ExpiresAt %s, got %s", *tc.expectedExpiresAt, *resp.ExpiresAt)
				}
			}
		})
	}
}

// TestAnnouncement_IsExpired tests expiration check
func TestAnnouncement_IsExpired(t *testing.T) {
	now := time.Now()
	past := now.Add(-1 * time.Hour)
	future := now.Add(24 * time.Hour)

	tests := []struct {
		name            string
		expiresAt       *time.Time
		expectedExpired bool
	}{
		{
			name:            "Announcement without expiration",
			expiresAt:       nil,
			expectedExpired: false,
		},
		{
			name:            "Announcement with future expiration",
			expiresAt:       &future,
			expectedExpired: false,
		},
		{
			name:            "Announcement with past expiration",
			expiresAt:       &past,
			expectedExpired: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			announcement := &Announcement{
				ID:             "ann1",
				SchoolID:       "school1",
				Title:          "Test",
				Content:        "Test content",
				Priority:       "NORMAL",
				TargetAudience: "ALL",
				PublishedBy:    "user1",
				PublishedAt:    now,
				ExpiresAt:      tc.expiresAt,
				IsActive:       true,
				CreatedAt:      now,
				UpdatedAt:      now,
			}

			expired := announcement.IsExpired()
			if expired != tc.expectedExpired {
				t.Errorf("expected expired %v, got %v", tc.expectedExpired, expired)
			}
		})
	}
}

// TestAnnouncement_Validate tests announcement validation
func TestAnnouncement_Validate(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name         string
		announcement *Announcement
		expectedErr  string
	}{
		{
			name: "Valid announcement",
			announcement: &Announcement{
				ID:             "ann1",
				SchoolID:       "school1",
				Title:          "School Holiday",
				Content:        "School will be closed for holidays",
				Priority:       "HIGH",
				TargetAudience: "ALL",
				PublishedBy:    "user1",
				PublishedAt:    now,
				IsActive:       true,
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			expectedErr: "",
		},
		{
			name: "Valid announcement with all priority types - LOW",
			announcement: &Announcement{
				ID:             "ann1",
				SchoolID:       "school1",
				Title:          "Test",
				Content:        "Test content",
				Priority:       "LOW",
				TargetAudience: "ALL",
				PublishedBy:    "user1",
				PublishedAt:    now,
				IsActive:       true,
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			expectedErr: "",
		},
		{
			name: "Valid announcement with all priority types - NORMAL",
			announcement: &Announcement{
				ID:             "ann1",
				SchoolID:       "school1",
				Title:          "Test",
				Content:        "Test content",
				Priority:       "NORMAL",
				TargetAudience: "ALL",
				PublishedBy:    "user1",
				PublishedAt:    now,
				IsActive:       true,
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			expectedErr: "",
		},
		{
			name: "Valid announcement with all priority types - URGENT",
			announcement: &Announcement{
				ID:             "ann1",
				SchoolID:       "school1",
				Title:          "Test",
				Content:        "Test content",
				Priority:       "URGENT",
				TargetAudience: "ALL",
				PublishedBy:    "user1",
				PublishedAt:    now,
				IsActive:       true,
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			expectedErr: "",
		},
		{
			name: "Valid announcement with all audience types - TEACHERS",
			announcement: &Announcement{
				ID:             "ann1",
				SchoolID:       "school1",
				Title:          "Test",
				Content:        "Test content",
				Priority:       "NORMAL",
				TargetAudience: "TEACHERS",
				PublishedBy:    "user1",
				PublishedAt:    now,
				IsActive:       true,
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			expectedErr: "",
		},
		{
			name: "Valid announcement with all audience types - STUDENTS",
			announcement: &Announcement{
				ID:             "ann1",
				SchoolID:       "school1",
				Title:          "Test",
				Content:        "Test content",
				Priority:       "NORMAL",
				TargetAudience: "STUDENTS",
				PublishedBy:    "user1",
				PublishedAt:    now,
				IsActive:       true,
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			expectedErr: "",
		},
		{
			name: "Valid announcement with all audience types - PARENTS",
			announcement: &Announcement{
				ID:             "ann1",
				SchoolID:       "school1",
				Title:          "Test",
				Content:        "Test content",
				Priority:       "NORMAL",
				TargetAudience: "PARENTS",
				PublishedBy:    "user1",
				PublishedAt:    now,
				IsActive:       true,
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			expectedErr: "",
		},
		{
			name: "Valid announcement with all audience types - ADMIN",
			announcement: &Announcement{
				ID:             "ann1",
				SchoolID:       "school1",
				Title:          "Test",
				Content:        "Test content",
				Priority:       "NORMAL",
				TargetAudience: "ADMIN",
				PublishedBy:    "user1",
				PublishedAt:    now,
				IsActive:       true,
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			expectedErr: "",
		},
		{
			name: "Invalid - missing ID",
			announcement: &Announcement{
				ID:             "",
				SchoolID:       "school1",
				Title:          "Test",
				Content:        "Test content",
				Priority:       "NORMAL",
				TargetAudience: "ALL",
				PublishedBy:    "user1",
				PublishedAt:    now,
				IsActive:       true,
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			expectedErr: "id is required",
		},
		{
			name: "Invalid - missing SchoolID",
			announcement: &Announcement{
				ID:             "ann1",
				SchoolID:       "",
				Title:          "Test",
				Content:        "Test content",
				Priority:       "NORMAL",
				TargetAudience: "ALL",
				PublishedBy:    "user1",
				PublishedAt:    now,
				IsActive:       true,
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			expectedErr: "school_id is required",
		},
		{
			name: "Invalid - missing Title",
			announcement: &Announcement{
				ID:             "ann1",
				SchoolID:       "school1",
				Title:          "",
				Content:        "Test content",
				Priority:       "NORMAL",
				TargetAudience: "ALL",
				PublishedBy:    "user1",
				PublishedAt:    now,
				IsActive:       true,
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			expectedErr: "title is required",
		},
		{
			name: "Invalid - missing Content",
			announcement: &Announcement{
				ID:             "ann1",
				SchoolID:       "school1",
				Title:          "Test",
				Content:        "",
				Priority:       "NORMAL",
				TargetAudience: "ALL",
				PublishedBy:    "user1",
				PublishedAt:    now,
				IsActive:       true,
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			expectedErr: "content is required",
		},
		{
			name: "Invalid - missing Priority",
			announcement: &Announcement{
				ID:             "ann1",
				SchoolID:       "school1",
				Title:          "Test",
				Content:        "Test content",
				Priority:       "",
				TargetAudience: "ALL",
				PublishedBy:    "user1",
				PublishedAt:    now,
				IsActive:       true,
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			expectedErr: "priority is required",
		},
		{
			name: "Invalid - invalid Priority",
			announcement: &Announcement{
				ID:             "ann1",
				SchoolID:       "school1",
				Title:          "Test",
				Content:        "Test content",
				Priority:       "INVALID",
				TargetAudience: "ALL",
				PublishedBy:    "user1",
				PublishedAt:    now,
				IsActive:       true,
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			expectedErr: "invalid priority: INVALID",
		},
		{
			name: "Invalid - missing TargetAudience",
			announcement: &Announcement{
				ID:             "ann1",
				SchoolID:       "school1",
				Title:          "Test",
				Content:        "Test content",
				Priority:       "NORMAL",
				TargetAudience: "",
				PublishedBy:    "user1",
				PublishedAt:    now,
				IsActive:       true,
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			expectedErr: "target_audience is required",
		},
		{
			name: "Invalid - invalid TargetAudience",
			announcement: &Announcement{
				ID:             "ann1",
				SchoolID:       "school1",
				Title:          "Test",
				Content:        "Test content",
				Priority:       "NORMAL",
				TargetAudience: "INVALID",
				PublishedBy:    "user1",
				PublishedAt:    now,
				IsActive:       true,
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			expectedErr: "invalid target_audience: INVALID",
		},
		{
			name: "Invalid - missing PublishedBy",
			announcement: &Announcement{
				ID:             "ann1",
				SchoolID:       "school1",
				Title:          "Test",
				Content:        "Test content",
				Priority:       "NORMAL",
				TargetAudience: "ALL",
				PublishedBy:    "",
				PublishedAt:    now,
				IsActive:       true,
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			expectedErr: "published_by is required",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.announcement.Validate()
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

package domain

import (
	"testing"
	"time"
)

// TestUser_ToUserResponse tests conversion to response format
func TestUser_ToUserResponse(t *testing.T) {
	now := time.Now()
	future := now.Add(24 * time.Hour)
	schoolID := "school123"

	tests := []struct {
		name            string
		user            *User
		roleName        string
		schoolName      string
		expectedStatus  UserStatus
		expectedSchoolID *string
		expectedSchoolName *string
	}{
		{
			name: "Active user with school",
			user: &User{
				ID:                  "user1",
				Email:               "user@example.com",
				PasswordHash:        "hash",
				Name:                "John Doe",
				RoleID:              "role1",
				SchoolID:            &schoolID,
				IsActive:            true,
				FailedLoginAttempts: 0,
				LockedUntil:         nil,
				CreatedAt:           now,
				UpdatedAt:           now,
			},
			roleName:        "Teacher",
			schoolName:      "Sekolah Indonesia",
			expectedStatus:  UserStatusActive,
			expectedSchoolID: &schoolID,
			expectedSchoolName: stringPtr("Sekolah Indonesia"),
		},
		{
			name: "Inactive user",
			user: &User{
				ID:                  "user2",
				Email:               "user2@example.com",
				PasswordHash:        "hash",
				Name:                "Jane Smith",
				RoleID:              "role2",
				SchoolID:            nil,
				IsActive:            false,
				FailedLoginAttempts: 0,
				LockedUntil:         nil,
				CreatedAt:           now,
				UpdatedAt:           now,
			},
			roleName:        "Admin",
			schoolName:      "",
			expectedStatus:  UserStatusInactive,
			expectedSchoolID: nil,
			expectedSchoolName: nil,
		},
		{
			name: "Suspended user (locked)",
			user: &User{
				ID:                  "user3",
				Email:               "user3@example.com",
				PasswordHash:        "hash",
				Name:                "Bob Johnson",
				RoleID:              "role3",
				SchoolID:            &schoolID,
				IsActive:            true,
				FailedLoginAttempts: 5,
				LockedUntil:         &future,
				CreatedAt:           now,
				UpdatedAt:           now,
			},
			roleName:        "Student",
			schoolName:      "Sekolah Indonesia",
			expectedStatus:  UserStatusSuspended,
			expectedSchoolID: &schoolID,
			expectedSchoolName: stringPtr("Sekolah Indonesia"),
		},
		{
			name: "Active user without school",
			user: &User{
				ID:                  "user4",
				Email:               "user4@example.com",
				PasswordHash:        "hash",
				Name:                "Alice Williams",
				RoleID:              "role4",
				SchoolID:            nil,
				IsActive:            true,
				FailedLoginAttempts: 0,
				LockedUntil:         nil,
				CreatedAt:           now,
				UpdatedAt:           now,
			},
			roleName:        "Super Admin",
			schoolName:      "",
			expectedStatus:  UserStatusActive,
			expectedSchoolID: nil,
			expectedSchoolName: nil,
		},
		{
			name: "User with school but no school name in response",
			user: &User{
				ID:                  "user5",
				Email:               "user5@example.com",
				PasswordHash:        "hash",
				Name:                "Charlie Brown",
				RoleID:              "role5",
				SchoolID:            &schoolID,
				IsActive:            true,
				FailedLoginAttempts: 0,
				LockedUntil:         nil,
				CreatedAt:           now,
				UpdatedAt:           now,
			},
			roleName:        "Teacher",
			schoolName:      "",
			expectedStatus:  UserStatusActive,
			expectedSchoolID: &schoolID,
			expectedSchoolName: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			resp := tc.user.ToUserResponse(tc.roleName, tc.schoolName)

			if resp.ID != tc.user.ID {
				t.Errorf("expected ID %s, got %s", tc.user.ID, resp.ID)
			}
			if resp.Email != tc.user.Email {
				t.Errorf("expected Email %s, got %s", tc.user.Email, resp.Email)
			}
			if resp.Name != tc.user.Name {
				t.Errorf("expected Name %s, got %s", tc.user.Name, resp.Name)
			}
			if resp.RoleID != tc.user.RoleID {
				t.Errorf("expected RoleID %s, got %s", tc.user.RoleID, resp.RoleID)
			}
			if resp.RoleName != tc.roleName {
				t.Errorf("expected RoleName %s, got %s", tc.roleName, resp.RoleName)
			}
			if resp.IsActive != tc.user.IsActive {
				t.Errorf("expected IsActive %v, got %v", tc.user.IsActive, resp.IsActive)
			}
			if resp.Status != tc.expectedStatus {
				t.Errorf("expected Status %s, got %s", tc.expectedStatus, resp.Status)
			}

			// Check optional fields
			if (resp.SchoolID == nil) != (tc.expectedSchoolID == nil) {
				t.Errorf("SchoolID pointer mismatch")
			} else if resp.SchoolID != nil && tc.expectedSchoolID != nil {
				if *resp.SchoolID != *tc.expectedSchoolID {
					t.Errorf("expected SchoolID %s, got %s", *tc.expectedSchoolID, *resp.SchoolID)
				}
			}

			if (resp.SchoolName == nil) != (tc.expectedSchoolName == nil) {
				t.Errorf("SchoolName pointer mismatch")
			} else if resp.SchoolName != nil && tc.expectedSchoolName != nil {
				if *resp.SchoolName != *tc.expectedSchoolName {
					t.Errorf("expected SchoolName %s, got %s", *tc.expectedSchoolName, *resp.SchoolName)
				}
			}
		})
	}
}

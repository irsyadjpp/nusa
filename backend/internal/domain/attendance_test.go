package domain

import (
	"testing"
	"time"
)

// TestAttendanceRecord_Validate tests validation of attendance records
func TestAttendanceRecord_Validate(t *testing.T) {
	tests := []struct {
		name                  string
		attendance            AttendanceRecord
		wantError             bool
		expectedErrorContains string
	}{
		{
			name: "Valid attendance record",
			attendance: AttendanceRecord{
				ID:         "att-1",
				ClassID:    "class-1",
				StudentID:  "student-1",
				Date:       time.Now(),
				Status:     string(AttendanceStatusPresent),
				RecordedBy: "user-1",
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now(),
			},
			wantError: false,
		},
		{
			name: "Valid - absent status",
			attendance: AttendanceRecord{
				ID:         "att-1",
				ClassID:    "class-1",
				StudentID:  "student-1",
				Date:       time.Now(),
				Status:     string(AttendanceStatusAbsent),
				RecordedBy: "user-1",
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now(),
			},
			wantError: false,
		},
		{
			name: "Valid - late status",
			attendance: AttendanceRecord{
				ID:         "att-1",
				ClassID:    "class-1",
				StudentID:  "student-1",
				Date:       time.Now(),
				Status:     string(AttendanceStatusLate),
				RecordedBy: "user-1",
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now(),
			},
			wantError: false,
		},
		{
			name: "Valid - excused status",
			attendance: AttendanceRecord{
				ID:         "att-1",
				ClassID:    "class-1",
				StudentID:  "student-1",
				Date:       time.Now(),
				Status:     string(AttendanceStatusExcused),
				RecordedBy: "user-1",
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now(),
			},
			wantError: false,
		},
		{
			name: "Valid - with notes",
			attendance: AttendanceRecord{
				ID:         "att-1",
				ClassID:    "class-1",
				StudentID:  "student-1",
				Date:       time.Now(),
				Status:     string(AttendanceStatusPresent),
				Notes:      makeStringPtr("Student was present"),
				RecordedBy: "user-1",
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now(),
			},
			wantError: false,
		},
		{
			name: "Valid - with deleted_at",
			attendance: AttendanceRecord{
				ID:         "att-1",
				ClassID:    "class-1",
				StudentID:  "student-1",
				Date:       time.Now(),
				Status:     string(AttendanceStatusPresent),
				RecordedBy: "user-1",
				DeletedAt:  func() *time.Time { t := time.Now(); return &t }(),
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now(),
			},
			wantError: false,
		},
		{
			name: "Invalid - empty ID",
			attendance: AttendanceRecord{
				ID:         "",
				ClassID:    "class-1",
				StudentID:  "student-1",
				Date:       time.Now(),
				Status:     string(AttendanceStatusPresent),
				RecordedBy: "user-1",
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now(),
			},
			wantError:             true,
			expectedErrorContains: "id is required",
		},
		{
			name: "Invalid - empty class ID",
			attendance: AttendanceRecord{
				ID:         "att-1",
				ClassID:    "",
				StudentID:  "student-1",
				Date:       time.Now(),
				Status:     string(AttendanceStatusPresent),
				RecordedBy: "user-1",
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now(),
			},
			wantError:             true,
			expectedErrorContains: "class_id is required",
		},
		{
			name: "Invalid - empty student ID",
			attendance: AttendanceRecord{
				ID:         "att-1",
				ClassID:    "class-1",
				StudentID:  "",
				Date:       time.Now(),
				Status:     string(AttendanceStatusPresent),
				RecordedBy: "user-1",
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now(),
			},
			wantError:             true,
			expectedErrorContains: "student_id is required",
		},
		{
			name: "Invalid - zero date",
			attendance: AttendanceRecord{
				ID:         "att-1",
				ClassID:    "class-1",
				StudentID:  "student-1",
				Date:       time.Time{},
				Status:     string(AttendanceStatusPresent),
				RecordedBy: "user-1",
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now(),
			},
			wantError:             true,
			expectedErrorContains: "date is required",
		},
		{
			name: "Invalid - empty status",
			attendance: AttendanceRecord{
				ID:         "att-1",
				ClassID:    "class-1",
				StudentID:  "student-1",
				Date:       time.Now(),
				Status:     "",
				RecordedBy: "user-1",
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now(),
			},
			wantError:             true,
			expectedErrorContains: "status is required",
		},
		{
			name: "Invalid - invalid status",
			attendance: AttendanceRecord{
				ID:         "att-1",
				ClassID:    "class-1",
				StudentID:  "student-1",
				Date:       time.Now(),
				Status:     "INVALID_STATUS",
				RecordedBy: "user-1",
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now(),
			},
			wantError:             true,
			expectedErrorContains: "invalid status: INVALID_STATUS",
		},
		{
			name: "Invalid - empty recorded_by",
			attendance: AttendanceRecord{
				ID:         "att-1",
				ClassID:    "class-1",
				StudentID:  "student-1",
				Date:       time.Now(),
				Status:     string(AttendanceStatusPresent),
				RecordedBy: "",
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now(),
			},
			wantError:             true,
			expectedErrorContains: "recorded_by is required",
		},
		{
			name: "Invalid - all required fields missing",
			attendance: AttendanceRecord{
				ID:         "",
				ClassID:    "",
				StudentID:  "",
				Date:       time.Time{},
				Status:     "",
				RecordedBy: "",
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now(),
			},
			wantError:             true,
			expectedErrorContains: "id is required", // Should return first error
		},
		{
			name: "Edge case - whitespace only ID",
			attendance: AttendanceRecord{
				ID:         "   ",
				ClassID:    "class-1",
				StudentID:  "student-1",
				Date:       time.Now(),
				Status:     string(AttendanceStatusPresent),
				RecordedBy: "user-1",
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now(),
			},
			wantError: false, // Validation doesn't trim whitespace
		},
		{
			name: "Edge case - mixed case status",
			attendance: AttendanceRecord{
				ID:         "att-1",
				ClassID:    "class-1",
				StudentID:  "student-1",
				Date:       time.Now(),
				Status:     "present", // lowercase
				RecordedBy: "user-1",
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now(),
			},
			wantError:             true,
			expectedErrorContains: "invalid status: present",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.attendance.Validate()

			if tc.wantError && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tc.wantError && err != nil {
				t.Errorf("expected no error, got: %v", err)
			}

			if tc.wantError && err != nil {
				if tc.expectedErrorContains != "" && !contains(err.Error(), tc.expectedErrorContains) {
					t.Errorf("expected error to contain %s, got %s", tc.expectedErrorContains, err.Error())
				}
			}
		})
	}
}

// TestAttendanceRecord_ToAttendanceResponse tests DTO transformation
func TestAttendanceRecord_ToAttendanceResponse(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name                    string
		attendance              AttendanceRecord
		className               string
		studentName             string
		recordedByName          string
		expectedID              string
		expectedStatus          AttendanceStatus
		expectedHasNotes        bool
		expectedHasClassName    bool
		expectedHasStudentName  bool
		expectedHasRecordedName bool
	}{
		{
			name: "Full transformation with all data",
			attendance: AttendanceRecord{
				ID:         "att-1",
				ClassID:    "class-1",
				StudentID:  "student-1",
				Date:       now,
				Status:     string(AttendanceStatusPresent),
				Notes:      makeStringPtr("Student was present"),
				RecordedBy: "user-1",
				CreatedAt:  now,
				UpdatedAt:  now,
			},
			className:               "Class A",
			studentName:             "John Doe",
			recordedByName:          "Jane Smith",
			expectedID:              "att-1",
			expectedStatus:          AttendanceStatusPresent,
			expectedHasNotes:        true,
			expectedHasClassName:    true,
			expectedHasStudentName:  true,
			expectedHasRecordedName: true,
		},
		{
			name: "Transformation with empty optional fields",
			attendance: AttendanceRecord{
				ID:         "att-1",
				ClassID:    "class-1",
				StudentID:  "student-1",
				Date:       now,
				Status:     string(AttendanceStatusAbsent),
				Notes:      nil,
				RecordedBy: "user-1",
				CreatedAt:  now,
				UpdatedAt:  now,
			},
			className:               "",
			studentName:             "",
			recordedByName:          "",
			expectedID:              "att-1",
			expectedStatus:          AttendanceStatusAbsent,
			expectedHasNotes:        false,
			expectedHasClassName:    false,
			expectedHasStudentName:  false,
			expectedHasRecordedName: false,
		},
		{
			name: "Transformation with only class name",
			attendance: AttendanceRecord{
				ID:         "att-1",
				ClassID:    "class-1",
				StudentID:  "student-1",
				Date:       now,
				Status:     string(AttendanceStatusLate),
				RecordedBy: "user-1",
				CreatedAt:  now,
				UpdatedAt:  now,
			},
			className:               "Class A",
			studentName:             "",
			recordedByName:          "",
			expectedID:              "att-1",
			expectedStatus:          AttendanceStatusLate,
			expectedHasNotes:        false,
			expectedHasClassName:    true,
			expectedHasStudentName:  false,
			expectedHasRecordedName: false,
		},
		{
			name: "Transformation with only student name",
			attendance: AttendanceRecord{
				ID:         "att-1",
				ClassID:    "class-1",
				StudentID:  "student-1",
				Date:       now,
				Status:     string(AttendanceStatusExcused),
				RecordedBy: "user-1",
				CreatedAt:  now,
				UpdatedAt:  now,
			},
			className:               "",
			studentName:             "John Doe",
			recordedByName:          "",
			expectedID:              "att-1",
			expectedStatus:          AttendanceStatusExcused,
			expectedHasNotes:        false,
			expectedHasClassName:    false,
			expectedHasStudentName:  true,
			expectedHasRecordedName: false,
		},
		{
			name: "Transformation with only recorded name",
			attendance: AttendanceRecord{
				ID:         "att-1",
				ClassID:    "class-1",
				StudentID:  "student-1",
				Date:       now,
				Status:     string(AttendanceStatusPresent),
				RecordedBy: "user-1",
				CreatedAt:  now,
				UpdatedAt:  now,
			},
			className:               "",
			studentName:             "",
			recordedByName:          "Jane Smith",
			expectedID:              "att-1",
			expectedStatus:          AttendanceStatusPresent,
			expectedHasNotes:        false,
			expectedHasClassName:    false,
			expectedHasStudentName:  false,
			expectedHasRecordedName: true,
		},
		{
			name: "Transformation with empty notes string",
			attendance: AttendanceRecord{
				ID:         "att-1",
				ClassID:    "class-1",
				StudentID:  "student-1",
				Date:       now,
				Status:     string(AttendanceStatusPresent),
				Notes:      makeStringPtr(""),
				RecordedBy: "user-1",
				CreatedAt:  now,
				UpdatedAt:  now,
			},
			className:               "Class A",
			studentName:             "John Doe",
			recordedByName:          "Jane Smith",
			expectedID:              "att-1",
			expectedStatus:          AttendanceStatusPresent,
			expectedHasNotes:        true, // Empty string is still a string
			expectedHasClassName:    true,
			expectedHasStudentName:  true,
			expectedHasRecordedName: true,
		},
		{
			name: "Transformation - date formatting",
			attendance: AttendanceRecord{
				ID:         "att-1",
				ClassID:    "class-1",
				StudentID:  "student-1",
				Date:       now,
				Status:     string(AttendanceStatusPresent),
				RecordedBy: "user-1",
				CreatedAt:  now,
				UpdatedAt:  now,
			},
			className:               "",
			studentName:             "",
			recordedByName:          "",
			expectedID:              "att-1",
			expectedStatus:          AttendanceStatusPresent,
			expectedHasNotes:        false,
			expectedHasClassName:    false,
			expectedHasStudentName:  false,
			expectedHasRecordedName: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			response := tc.attendance.ToAttendanceResponse(tc.className, tc.studentName, tc.recordedByName)

			if response == nil {
				t.Fatal("expected response, got nil")
			}

			if response.ID != tc.expectedID {
				t.Errorf("expected ID %s, got %s", tc.expectedID, response.ID)
			}
			if response.ClassID != tc.attendance.ClassID {
				t.Errorf("expected class ID %s, got %s", tc.attendance.ClassID, response.ClassID)
			}
			if response.StudentID != tc.attendance.StudentID {
				t.Errorf("expected student ID %s, got %s", tc.attendance.StudentID, response.StudentID)
			}
			if response.Status != tc.expectedStatus {
				t.Errorf("expected status %s, got %s", tc.expectedStatus, response.Status)
			}
			if response.RecordedBy != tc.attendance.RecordedBy {
				t.Errorf("expected recorded_by %s, got %s", tc.attendance.RecordedBy, response.RecordedBy)
			}

			// Check optional fields
			if tc.expectedHasClassName && response.ClassName == nil {
				t.Error("expected ClassName to be set")
			}
			if !tc.expectedHasClassName && response.ClassName != nil {
				t.Error("expected ClassName to be nil")
			}
			if tc.expectedHasStudentName && response.StudentName == nil {
				t.Error("expected StudentName to be set")
			}
			if !tc.expectedHasStudentName && response.StudentName != nil {
				t.Error("expected StudentName to be nil")
			}
			if tc.expectedHasRecordedName && response.RecordedByName == nil {
				t.Error("expected RecordedByName to be set")
			}
			if !tc.expectedHasRecordedName && response.RecordedByName != nil {
				t.Error("expected RecordedByName to be nil")
			}
			if tc.expectedHasNotes && response.Notes == nil {
				t.Error("expected Notes to be set")
			}
			if !tc.expectedHasNotes && response.Notes != nil {
				t.Error("expected Notes to be nil")
			}

			// Check date formatting
			expectedDateFormat := now.Format("2006-01-02")
			if response.Date != expectedDateFormat {
				t.Errorf("expected date %s, got %s", expectedDateFormat, response.Date)
			}

			// Check timestamp formatting
			expectedTimestampFormat := now.Format(time.RFC3339)
			if response.CreatedAt != expectedTimestampFormat {
				t.Errorf("expected created_at %s, got %s", expectedTimestampFormat, response.CreatedAt)
			}
			if response.UpdatedAt != expectedTimestampFormat {
				t.Errorf("expected updated_at %s, got %s", expectedTimestampFormat, response.UpdatedAt)
			}
		})
	}
}

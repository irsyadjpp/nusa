package domain

import (
	"testing"
	"time"
)

func TestExamResult_Validate(t *testing.T) {
	tests := []struct {
		name       string
		examResult *ExamResult
		wantErr    bool
	}{
		{
			name: "valid exam result",
			examResult: &ExamResult{
				ID:        "exam-result-123",
				ExamID:    "exam-123",
				StudentID: "student-123",
				Score:     float64Ptr(85.5),
				Grade:     stringPtr("A"),
			},
			wantErr: false,
		},
		{
			name: "valid exam result without score",
			examResult: &ExamResult{
				ID:        "exam-result-123",
				ExamID:    "exam-123",
				StudentID: "student-123",
			},
			wantErr: false,
		},
		{
			name: "missing id",
			examResult: &ExamResult{
				ExamID:    "exam-123",
				StudentID: "student-123",
				Score:     float64Ptr(85.5),
			},
			wantErr: true,
		},
		{
			name: "missing exam_id",
			examResult: &ExamResult{
				ID:        "exam-result-123",
				StudentID: "student-123",
				Score:     float64Ptr(85.5),
			},
			wantErr: true,
		},
		{
			name: "missing student_id",
			examResult: &ExamResult{
				ID:     "exam-result-123",
				ExamID: "exam-123",
				Score:  float64Ptr(85.5),
			},
			wantErr: true,
		},
		{
			name: "invalid score (negative)",
			examResult: &ExamResult{
				ID:        "exam-result-123",
				ExamID:    "exam-123",
				StudentID: "student-123",
				Score:     float64Ptr(-1.0),
			},
			wantErr: true,
		},
		{
			name: "invalid score (greater than 100)",
			examResult: &ExamResult{
				ID:        "exam-result-123",
				ExamID:    "exam-123",
				StudentID: "student-123",
				Score:     float64Ptr(101.0),
			},
			wantErr: true,
		},
		{
			name: "valid score (exactly 0)",
			examResult: &ExamResult{
				ID:        "exam-result-123",
				ExamID:    "exam-123",
				StudentID: "student-123",
				Score:     float64Ptr(0.0),
			},
			wantErr: false,
		},
		{
			name: "valid score (exactly 100)",
			examResult: &ExamResult{
				ID:        "exam-result-123",
				ExamID:    "exam-123",
				StudentID: "student-123",
				Score:     float64Ptr(100.0),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.examResult.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ExamResult.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestExamResult_IsGraded(t *testing.T) {
	tests := []struct {
		name       string
		examResult *ExamResult
		want       bool
	}{
		{
			name: "exam result is graded",
			examResult: &ExamResult{
				GradedAt: timePtr(time.Now()),
				GradedBy: stringPtr("teacher-123"),
			},
			want: true,
		},
		{
			name: "exam result is not graded (no graded_at)",
			examResult: &ExamResult{
				GradedBy: stringPtr("teacher-123"),
			},
			want: false,
		},
		{
			name: "exam result is not graded (no graded_by)",
			examResult: &ExamResult{
				GradedAt: timePtr(time.Now()),
			},
			want: false,
		},
		{
			name: "exam result is not graded (both missing)",
			examResult: &ExamResult{},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.examResult.IsGraded(); got != tt.want {
				t.Errorf("ExamResult.IsGraded() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExamResult_MarkAsGraded(t *testing.T) {
	examResult := &ExamResult{
		ID:        "exam-result-123",
		ExamID:    "exam-123",
		StudentID: "student-123",
	}

	graderID := "teacher-123"
	examResult.MarkAsGraded(graderID)

	if examResult.GradedAt == nil {
		t.Error("MarkAsGraded() did not set GradedAt")
	}
	if examResult.GradedBy == nil {
		t.Error("MarkAsGraded() did not set GradedBy")
	}
	if *examResult.GradedBy != graderID {
		t.Errorf("MarkAsGraded() GradedBy = %v, want %v", *examResult.GradedBy, graderID)
	}
}

func TestExamResult_ToExamResultResponse(t *testing.T) {
	gradedAt := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
	examResult := &ExamResult{
		ID:        "exam-result-123",
		ExamID:    "exam-123",
		StudentID: "student-123",
		Score:     float64Ptr(85.5),
		Grade:     stringPtr("A"),
		Remarks:   stringPtr("Good performance"),
		GradedAt:  &gradedAt,
		GradedBy:  stringPtr("teacher-123"),
		CreatedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	response := examResult.ToExamResultResponse("2024-01-15", "Math Exam", "John Doe", "Jane Smith")

	if response.ID != examResult.ID {
		t.Errorf("ToExamResultResponse() ID = %v, want %v", response.ID, examResult.ID)
	}
	if response.ExamID != examResult.ExamID {
		t.Errorf("ToExamResultResponse() ExamID = %v, want %v", response.ExamID, examResult.ExamID)
	}
	if *response.ExamDate != "2024-01-15" {
		t.Errorf("ToExamResultResponse() ExamDate = %v, want 2024-01-15", *response.ExamDate)
	}
	if *response.StudentName != "John Doe" {
		t.Errorf("ToExamResultResponse() StudentName = %v, want John Doe", *response.StudentName)
	}
	if *response.Score != 85.5 {
		t.Errorf("ToExamResultResponse() Score = %v, want 85.5", *response.Score)
	}
	if *response.GradedByName != "Jane Smith" {
		t.Errorf("ToExamResultResponse() GradedByName = %v, want Jane Smith", *response.GradedByName)
	}
}

func float64Ptr(f float64) *float64 {
	return &f
}

func timePtr(t time.Time) *time.Time {
	return &t
}

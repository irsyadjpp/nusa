package domain

import (
	"testing"
	"time"
)

func TestExam_Validate(t *testing.T) {
	tests := []struct {
		name    string
		exam    *Exam
		wantErr bool
	}{
		{
			name: "valid exam",
			exam: &Exam{
				ID:              "exam-123",
				ClassID:         "class-123",
				AssessmentID:    "assessment-123",
				ExamDate:        time.Now().Add(24 * time.Hour),
				StartTime:       "09:00",
				DurationMinutes: 60,
				Status:          string(ExamStatusScheduled),
			},
			wantErr: false,
		},
		{
			name: "missing id",
			exam: &Exam{
				ClassID:         "class-123",
				AssessmentID:    "assessment-123",
				ExamDate:        time.Now().Add(24 * time.Hour),
				StartTime:       "09:00",
				DurationMinutes: 60,
				Status:          string(ExamStatusScheduled),
			},
			wantErr: true,
		},
		{
			name: "missing class_id",
			exam: &Exam{
				ID:              "exam-123",
				AssessmentID:    "assessment-123",
				ExamDate:        time.Now().Add(24 * time.Hour),
				StartTime:       "09:00",
				DurationMinutes: 60,
				Status:          string(ExamStatusScheduled),
			},
			wantErr: true,
		},
		{
			name: "missing assessment_id",
			exam: &Exam{
				ID:              "exam-123",
				ClassID:         "class-123",
				ExamDate:        time.Now().Add(24 * time.Hour),
				StartTime:       "09:00",
				DurationMinutes: 60,
				Status:          string(ExamStatusScheduled),
			},
			wantErr: true,
		},
		{
			name: "missing exam_date",
			exam: &Exam{
				ID:              "exam-123",
				ClassID:         "class-123",
				AssessmentID:    "assessment-123",
				StartTime:       "09:00",
				DurationMinutes: 60,
				Status:          string(ExamStatusScheduled),
			},
			wantErr: true,
		},
		{
			name: "missing start_time",
			exam: &Exam{
				ID:              "exam-123",
				ClassID:         "class-123",
				AssessmentID:    "assessment-123",
				ExamDate:        time.Now().Add(24 * time.Hour),
				DurationMinutes: 60,
				Status:          string(ExamStatusScheduled),
			},
			wantErr: true,
		},
		{
			name: "invalid duration_minutes",
			exam: &Exam{
				ID:              "exam-123",
				ClassID:         "class-123",
				AssessmentID:    "assessment-123",
				ExamDate:        time.Now().Add(24 * time.Hour),
				StartTime:       "09:00",
				DurationMinutes: 0,
				Status:          string(ExamStatusScheduled),
			},
			wantErr: true,
		},
		{
			name: "missing status",
			exam: &Exam{
				ID:              "exam-123",
				ClassID:         "class-123",
				AssessmentID:    "assessment-123",
				ExamDate:        time.Now().Add(24 * time.Hour),
				StartTime:       "09:00",
				DurationMinutes: 60,
			},
			wantErr: true,
		},
		{
			name: "invalid status",
			exam: &Exam{
				ID:              "exam-123",
				ClassID:         "class-123",
				AssessmentID:    "assessment-123",
				ExamDate:        time.Now().Add(24 * time.Hour),
				StartTime:       "09:00",
				DurationMinutes: 60,
				Status:          "INVALID",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.exam.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Exam.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestExam_ToExamResponse(t *testing.T) {
	exam := &Exam{
		ID:              "exam-123",
		ClassID:         "class-123",
		AssessmentID:    "assessment-123",
		ExamDate:        time.Date(2024, 1, 15, 9, 0, 0, 0, time.UTC),
		StartTime:       "09:00",
		DurationMinutes: 60,
		Room:            stringPtr("Room 101"),
		Status:          string(ExamStatusScheduled),
		CreatedAt:       time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:       time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		CreatedBy:       stringPtr("user-123"),
		UpdatedBy:       stringPtr("user-123"),
	}

	response := exam.ToExamResponse("Class A", "Formative", "John Doe", "John Doe")

	if response.ID != exam.ID {
		t.Errorf("ToExamResponse() ID = %v, want %v", response.ID, exam.ID)
	}
	if response.ClassID != exam.ClassID {
		t.Errorf("ToExamResponse() ClassID = %v, want %v", response.ClassID, exam.ClassID)
	}
	if *response.ClassName != "Class A" {
		t.Errorf("ToExamResponse() ClassName = %v, want Class A", *response.ClassName)
	}
	if response.ExamDate != exam.ExamDate.Format(time.RFC3339) {
		t.Errorf("ToExamResponse() ExamDate = %v, want %v", response.ExamDate, exam.ExamDate.Format(time.RFC3339))
	}
	if response.Status != ExamStatus(exam.Status) {
		t.Errorf("ToExamResponse() Status = %v, want %v", response.Status, exam.Status)
	}
}

func stringPtr(s string) *string {
	return &s
}

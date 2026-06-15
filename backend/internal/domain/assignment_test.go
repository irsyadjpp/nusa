package domain

import (
	"testing"
	"time"
)

func TestAssignment_Validate(t *testing.T) {
	tests := []struct {
		name       string
		assignment *Assignment
		wantErr    bool
	}{
		{
			name: "valid assignment",
			assignment: &Assignment{
				ID:           "assignment-123",
				ClassID:      "class-123",
				AssessmentID: "assessment-123",
				Title:        "Math Homework",
				DueDate:      time.Now().Add(7 * 24 * time.Hour),
				MaxScore:     100,
				Status:       string(AssignmentStatusAssigned),
			},
			wantErr: false,
		},
		{
			name: "missing id",
			assignment: &Assignment{
				ClassID:      "class-123",
				AssessmentID: "assessment-123",
				Title:        "Math Homework",
				DueDate:      time.Now().Add(7 * 24 * time.Hour),
				MaxScore:     100,
				Status:       string(AssignmentStatusAssigned),
			},
			wantErr: true,
		},
		{
			name: "missing class_id",
			assignment: &Assignment{
				ID:           "assignment-123",
				AssessmentID: "assessment-123",
				Title:        "Math Homework",
				DueDate:      time.Now().Add(7 * 24 * time.Hour),
				MaxScore:     100,
				Status:       string(AssignmentStatusAssigned),
			},
			wantErr: true,
		},
		{
			name: "missing assessment_id",
			assignment: &Assignment{
				ID:       "assignment-123",
				ClassID:  "class-123",
				Title:    "Math Homework",
				DueDate:  time.Now().Add(7 * 24 * time.Hour),
				MaxScore: 100,
				Status:   string(AssignmentStatusAssigned),
			},
			wantErr: true,
		},
		{
			name: "missing title",
			assignment: &Assignment{
				ID:           "assignment-123",
				ClassID:      "class-123",
				AssessmentID: "assessment-123",
				DueDate:      time.Now().Add(7 * 24 * time.Hour),
				MaxScore:     100,
				Status:       string(AssignmentStatusAssigned),
			},
			wantErr: true,
		},
		{
			name: "missing due_date",
			assignment: &Assignment{
				ID:           "assignment-123",
				ClassID:      "class-123",
				AssessmentID: "assessment-123",
				Title:        "Math Homework",
				MaxScore:     100,
				Status:       string(AssignmentStatusAssigned),
			},
			wantErr: true,
		},
		{
			name: "invalid max_score",
			assignment: &Assignment{
				ID:           "assignment-123",
				ClassID:      "class-123",
				AssessmentID: "assessment-123",
				Title:        "Math Homework",
				DueDate:      time.Now().Add(7 * 24 * time.Hour),
				MaxScore:     0,
				Status:       string(AssignmentStatusAssigned),
			},
			wantErr: true,
		},
		{
			name: "missing status",
			assignment: &Assignment{
				ID:           "assignment-123",
				ClassID:      "class-123",
				AssessmentID: "assessment-123",
				Title:        "Math Homework",
				DueDate:      time.Now().Add(7 * 24 * time.Hour),
				MaxScore:     100,
			},
			wantErr: true,
		},
		{
			name: "invalid status",
			assignment: &Assignment{
				ID:           "assignment-123",
				ClassID:      "class-123",
				AssessmentID: "assessment-123",
				Title:        "Math Homework",
				DueDate:      time.Now().Add(7 * 24 * time.Hour),
				MaxScore:     100,
				Status:       "INVALID",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.assignment.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Assignment.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAssignment_IsOverdue(t *testing.T) {
	tests := []struct {
		name       string
		assignment *Assignment
		want       bool
	}{
		{
			name: "assignment overdue",
			assignment: &Assignment{
				DueDate: time.Now().Add(-24 * time.Hour),
				Status:  string(AssignmentStatusAssigned),
			},
			want: true,
		},
		{
			name: "assignment not overdue",
			assignment: &Assignment{
				DueDate: time.Now().Add(24 * time.Hour),
				Status:  string(AssignmentStatusAssigned),
			},
			want: false,
		},
		{
			name: "assignment graded (not overdue even if past due date)",
			assignment: &Assignment{
				DueDate: time.Now().Add(-24 * time.Hour),
				Status:  string(AssignmentStatusGraded),
			},
			want: false,
		},
		{
			name: "assignment cancelled (not overdue even if past due date)",
			assignment: &Assignment{
				DueDate: time.Now().Add(-24 * time.Hour),
				Status:  string(AssignmentStatusCancelled),
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.assignment.IsOverdue(); got != tt.want {
				t.Errorf("Assignment.IsOverdue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAssignment_ToAssignmentResponse(t *testing.T) {
	assignment := &Assignment{
		ID:           "assignment-123",
		ClassID:      "class-123",
		AssessmentID: "assessment-123",
		Title:        "Math Homework",
		Description:  stringPtr("Complete exercises 1-10"),
		DueDate:      time.Date(2024, 1, 15, 23, 59, 0, 0, time.UTC),
		MaxScore:     100,
		Status:       string(AssignmentStatusAssigned),
		CreatedAt:    time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:    time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		CreatedBy:    stringPtr("user-123"),
		UpdatedBy:    stringPtr("user-123"),
	}

	response := assignment.ToAssignmentResponse("Class A", "Formative", "John Doe", "John Doe")

	if response.ID != assignment.ID {
		t.Errorf("ToAssignmentResponse() ID = %v, want %v", response.ID, assignment.ID)
	}
	if response.ClassID != assignment.ClassID {
		t.Errorf("ToAssignmentResponse() ClassID = %v, want %v", response.ClassID, assignment.ClassID)
	}
	if *response.ClassName != "Class A" {
		t.Errorf("ToAssignmentResponse() ClassName = %v, want Class A", *response.ClassName)
	}
	if response.DueDate != assignment.DueDate.Format(time.RFC3339) {
		t.Errorf("ToAssignmentResponse() DueDate = %v, want %v", response.DueDate, assignment.DueDate.Format(time.RFC3339))
	}
	if response.Status != AssignmentStatus(assignment.Status) {
		t.Errorf("ToAssignmentResponse() Status = %v, want %v", response.Status, assignment.Status)
	}
}

package domain

import (
	"testing"
	"time"
)

// TestEvaluation_CreateRevision tests creating new evaluation revisions
func TestEvaluation_CreateRevision(t *testing.T) {
	tests := []struct {
		name                  string
		evaluation            Evaluation
		newPerformanceScores  interface{}
		newTotalScore         int
		newMaxScore           int
		newPerformanceLevel   EvaluationPerformanceLevel
		newTeacherFeedback    *string
		userID                string
		wantError             bool
		expectedErrorContains string
		expectedRevisionNo    int
	}{
		{
			name: "Valid revision creation",
			evaluation: Evaluation{
				ID:                "eval-1",
				StudentID:         "student-1",
				RubricID:          "rubric-1",
				EvidenceID:        "evidence-1",
				UserID:            "user-1",
				PerformanceScores: map[string]interface{}{"criterion1": 85},
				TotalScore:        85,
				MaxScore:          100,
				PerformanceLevel:  EvaluationPerformanceLevelProficient,
				TeacherFeedback:   makeStringPtr("Good work"),
				RevisionNo:        1,
				IsCurrentVersion:  true,
				EvaluatedAt:       time.Now(),
				CreatedAt:         time.Now(),
				UpdatedAt:         time.Now(),
			},
			newPerformanceScores: map[string]interface{}{"criterion1": 90},
			newTotalScore:        90,
			newMaxScore:          100,
			newPerformanceLevel:  EvaluationPerformanceLevelExcellent,
			newTeacherFeedback:   makeStringPtr("Excellent work"),
			userID:               "user-2",
			wantError:            false,
			expectedRevisionNo:   2,
		},
		{
			name: "First revision (revision 1)",
			evaluation: Evaluation{
				ID:                "eval-1",
				StudentID:         "student-1",
				RubricID:          "rubric-1",
				EvidenceID:        "evidence-1",
				UserID:            "user-1",
				PerformanceScores: map[string]interface{}{"criterion1": 75},
				TotalScore:        75,
				MaxScore:          100,
				PerformanceLevel:  EvaluationPerformanceLevelProficient,
				RevisionNo:        1,
				IsCurrentVersion:  true,
				EvaluatedAt:       time.Now(),
				CreatedAt:         time.Now(),
				UpdatedAt:         time.Now(),
			},
			newPerformanceScores: map[string]interface{}{"criterion1": 80},
			newTotalScore:        80,
			newMaxScore:          100,
			newPerformanceLevel:  EvaluationPerformanceLevelProficient,
			newTeacherFeedback:   nil,
			userID:               "user-1",
			wantError:            false,
			expectedRevisionNo:   2,
		},
		{
			name: "Higher revision number (revision 5)",
			evaluation: Evaluation{
				ID:                "eval-1",
				StudentID:         "student-1",
				RubricID:          "rubric-1",
				EvidenceID:        "evidence-1",
				UserID:            "user-1",
				PerformanceScores: map[string]interface{}{"criterion1": 75},
				TotalScore:        75,
				MaxScore:          100,
				PerformanceLevel:  EvaluationPerformanceLevelProficient,
				RevisionNo:        5,
				IsCurrentVersion:  true,
				EvaluatedAt:       time.Now(),
				CreatedAt:         time.Now(),
				UpdatedAt:         time.Now(),
			},
			newPerformanceScores: map[string]interface{}{"criterion1": 85},
			newTotalScore:        85,
			newMaxScore:          100,
			newPerformanceLevel:  EvaluationPerformanceLevelProficient,
			newTeacherFeedback:   makeStringPtr("Improved"),
			userID:               "user-1",
			wantError:            false,
			expectedRevisionNo:   6,
		},
		{
			name: "Invalid - empty evaluation ID",
			evaluation: Evaluation{
				ID:                "",
				StudentID:         "student-1",
				RubricID:          "rubric-1",
				EvidenceID:        "evidence-1",
				UserID:            "user-1",
				PerformanceScores: map[string]interface{}{"criterion1": 75},
				TotalScore:        75,
				MaxScore:          100,
				PerformanceLevel:  EvaluationPerformanceLevelProficient,
				RevisionNo:        1,
				IsCurrentVersion:  true,
				EvaluatedAt:       time.Now(),
				CreatedAt:         time.Now(),
				UpdatedAt:         time.Now(),
			},
			newPerformanceScores:  map[string]interface{}{"criterion1": 80},
			newTotalScore:         80,
			newMaxScore:           100,
			newPerformanceLevel:   EvaluationPerformanceLevelProficient,
			newTeacherFeedback:    nil,
			userID:                "user-1",
			wantError:             true,
			expectedErrorContains: "evaluation ID is required",
		},
		{
			name: "Invalid - empty evidence ID",
			evaluation: Evaluation{
				ID:                "eval-1",
				StudentID:         "student-1",
				RubricID:          "rubric-1",
				EvidenceID:        "",
				UserID:            "user-1",
				PerformanceScores: map[string]interface{}{"criterion1": 75},
				TotalScore:        75,
				MaxScore:          100,
				PerformanceLevel:  EvaluationPerformanceLevelProficient,
				RevisionNo:        1,
				IsCurrentVersion:  true,
				EvaluatedAt:       time.Now(),
				CreatedAt:         time.Now(),
				UpdatedAt:         time.Now(),
			},
			newPerformanceScores:  map[string]interface{}{"criterion1": 80},
			newTotalScore:         80,
			newMaxScore:           100,
			newPerformanceLevel:   EvaluationPerformanceLevelProficient,
			newTeacherFeedback:    nil,
			userID:                "user-1",
			wantError:             true,
			expectedErrorContains: "evidence ID is required",
		},
		{
			name: "Use default values - nil scores, zero values",
			evaluation: Evaluation{
				ID:                "eval-1",
				StudentID:         "student-1",
				RubricID:          "rubric-1",
				EvidenceID:        "evidence-1",
				UserID:            "user-1",
				PerformanceScores: map[string]interface{}{"criterion1": 75},
				TotalScore:        75,
				MaxScore:          100,
				PerformanceLevel:  EvaluationPerformanceLevelProficient,
				TeacherFeedback:   makeStringPtr("Original feedback"),
				RevisionNo:        1,
				IsCurrentVersion:  true,
				EvaluatedAt:       time.Now(),
				CreatedAt:         time.Now(),
				UpdatedAt:         time.Now(),
			},
			newPerformanceScores: nil, // Should use existing
			newTotalScore:        0,   // Should use existing
			newMaxScore:          0,   // Should use existing
			newPerformanceLevel:  "",  // Should use existing
			newTeacherFeedback:   nil, // Should use existing
			userID:               "user-2",
			wantError:            false,
			expectedRevisionNo:   2,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			newRevision, err := tc.evaluation.CreateRevision(
				tc.newPerformanceScores,
				tc.newTotalScore,
				tc.newMaxScore,
				tc.newPerformanceLevel,
				tc.newTeacherFeedback,
				tc.userID,
			)

			if tc.wantError && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tc.wantError && err != nil {
				t.Errorf("expected no error, got: %v", err)
			}

			if !tc.wantError {
				if newRevision == nil {
					t.Fatal("expected new revision, got nil")
				}

				// ID is intentionally empty (set by repository)
				if newRevision.ID != "" {
					t.Errorf("expected empty ID (set by repository), got %s", newRevision.ID)
				}
				if newRevision.StudentID != tc.evaluation.StudentID {
					t.Errorf("expected student ID %s, got %s", tc.evaluation.StudentID, newRevision.StudentID)
				}
				if newRevision.RubricID != tc.evaluation.RubricID {
					t.Errorf("expected rubric ID %s, got %s", tc.evaluation.RubricID, newRevision.RubricID)
				}
				if newRevision.EvidenceID != tc.evaluation.EvidenceID {
					t.Errorf("expected evidence ID %s, got %s", tc.evaluation.EvidenceID, newRevision.EvidenceID)
				}
				if newRevision.RevisionNo != tc.expectedRevisionNo {
					t.Errorf("expected revision no %d, got %d", tc.expectedRevisionNo, newRevision.RevisionNo)
				}
				if !newRevision.IsCurrentVersion {
					t.Error("new revision should be current version")
				}
				if newRevision.ParentRevisionID == nil {
					t.Error("new revision should have parent revision ID")
				}
				if *newRevision.ParentRevisionID != tc.evaluation.ID {
					t.Errorf("expected parent revision ID %s, got %s", tc.evaluation.ID, *newRevision.ParentRevisionID)
				}
				if newRevision.UserID != tc.userID {
					t.Errorf("expected user ID %s, got %s", tc.userID, newRevision.UserID)
				}

				// For the default values test case, verify defaults are used
				if tc.name == "Use default values - nil scores, zero values" {
					if newRevision.TotalScore != tc.evaluation.TotalScore {
						t.Errorf("expected default total score %d, got %d", tc.evaluation.TotalScore, newRevision.TotalScore)
					}
					if newRevision.MaxScore != tc.evaluation.MaxScore {
						t.Errorf("expected default max score %d, got %d", tc.evaluation.MaxScore, newRevision.MaxScore)
					}
					if newRevision.PerformanceLevel != tc.evaluation.PerformanceLevel {
						t.Errorf("expected default performance level %s, got %s", tc.evaluation.PerformanceLevel, newRevision.PerformanceLevel)
					}
					if newRevision.TeacherFeedback == nil || *newRevision.TeacherFeedback != *tc.evaluation.TeacherFeedback {
						t.Errorf("expected default teacher feedback, got %v", newRevision.TeacherFeedback)
					}
				} else {
					if newRevision.TotalScore != tc.newTotalScore {
						t.Errorf("expected total score %d, got %d", tc.newTotalScore, newRevision.TotalScore)
					}
					if newRevision.MaxScore != tc.newMaxScore {
						t.Errorf("expected max score %d, got %d", tc.newMaxScore, newRevision.MaxScore)
					}
					if newRevision.PerformanceLevel != tc.newPerformanceLevel {
						t.Errorf("expected performance level %s, got %s", tc.newPerformanceLevel, newRevision.PerformanceLevel)
					}
				}
				if newRevision.CreatedAt.IsZero() {
					t.Error("CreatedAt should not be zero")
				}
				if newRevision.UpdatedAt.IsZero() {
					t.Error("UpdatedAt should not be zero")
				}
			}

			if tc.wantError && err != nil {
				if tc.expectedErrorContains != "" && !contains(err.Error(), tc.expectedErrorContains) {
					t.Errorf("expected error to contain %s, got %s", tc.expectedErrorContains, err.Error())
				}
			}
		})
	}
}

// TestEvaluation_IsValidRevision tests validation of evaluation for revision
func TestEvaluation_IsValidRevision(t *testing.T) {
	tests := []struct {
		name                  string
		evaluation            Evaluation
		wantError             bool
		expectedErrorContains string
	}{
		{
			name: "Valid evaluation",
			evaluation: Evaluation{
				ID:                "eval-1",
				StudentID:         "student-1",
				RubricID:          "rubric-1",
				EvidenceID:        "evidence-1",
				UserID:            "user-1",
				PerformanceScores: map[string]interface{}{"criterion1": 85},
				TotalScore:        85,
				MaxScore:          100,
				PerformanceLevel:  EvaluationPerformanceLevelProficient,
				RevisionNo:        1,
				IsCurrentVersion:  true,
				EvaluatedAt:       time.Now(),
				CreatedAt:         time.Now(),
				UpdatedAt:         time.Now(),
			},
			wantError: false,
		},
		{
			name: "Valid - zero total score is allowed",
			evaluation: Evaluation{
				ID:                "eval-1",
				StudentID:         "student-1",
				RubricID:          "rubric-1",
				EvidenceID:        "evidence-1",
				UserID:            "user-1",
				PerformanceScores: map[string]interface{}{"criterion1": 0},
				TotalScore:        0,
				MaxScore:          100,
				PerformanceLevel:  EvaluationPerformanceLevelBeginning,
				RevisionNo:        1,
				IsCurrentVersion:  true,
				EvaluatedAt:       time.Now(),
				CreatedAt:         time.Now(),
				UpdatedAt:         time.Now(),
			},
			wantError: false,
		},
		{
			name: "Invalid - empty ID",
			evaluation: Evaluation{
				ID:                "",
				StudentID:         "student-1",
				RubricID:          "rubric-1",
				EvidenceID:        "evidence-1",
				UserID:            "user-1",
				PerformanceScores: map[string]interface{}{"criterion1": 85},
				TotalScore:        85,
				MaxScore:          100,
				PerformanceLevel:  EvaluationPerformanceLevelProficient,
				RevisionNo:        1,
				IsCurrentVersion:  true,
				EvaluatedAt:       time.Now(),
				CreatedAt:         time.Now(),
				UpdatedAt:         time.Now(),
			},
			wantError:             true,
			expectedErrorContains: "evaluation ID is required",
		},
		{
			name: "Invalid - empty evidence ID",
			evaluation: Evaluation{
				ID:                "eval-1",
				StudentID:         "student-1",
				RubricID:          "rubric-1",
				EvidenceID:        "",
				UserID:            "user-1",
				PerformanceScores: map[string]interface{}{"criterion1": 85},
				TotalScore:        85,
				MaxScore:          100,
				PerformanceLevel:  EvaluationPerformanceLevelProficient,
				RevisionNo:        1,
				IsCurrentVersion:  true,
				EvaluatedAt:       time.Now(),
				CreatedAt:         time.Now(),
				UpdatedAt:         time.Now(),
			},
			wantError:             true,
			expectedErrorContains: "evidence ID is required",
		},
		{
			name: "Invalid - zero revision number",
			evaluation: Evaluation{
				ID:                "eval-1",
				StudentID:         "student-1",
				RubricID:          "rubric-1",
				EvidenceID:        "evidence-1",
				UserID:            "user-1",
				PerformanceScores: map[string]interface{}{"criterion1": 85},
				TotalScore:        85,
				MaxScore:          100,
				PerformanceLevel:  EvaluationPerformanceLevelProficient,
				RevisionNo:        0,
				IsCurrentVersion:  true,
				EvaluatedAt:       time.Now(),
				CreatedAt:         time.Now(),
				UpdatedAt:         time.Now(),
			},
			wantError:             true,
			expectedErrorContains: "revision number must be at least 1",
		},
		{
			name: "Invalid - negative total score",
			evaluation: Evaluation{
				ID:                "eval-1",
				StudentID:         "student-1",
				RubricID:          "rubric-1",
				EvidenceID:        "evidence-1",
				UserID:            "user-1",
				PerformanceScores: map[string]interface{}{"criterion1": 85},
				TotalScore:        -10,
				MaxScore:          100,
				PerformanceLevel:  EvaluationPerformanceLevelProficient,
				RevisionNo:        1,
				IsCurrentVersion:  true,
				EvaluatedAt:       time.Now(),
				CreatedAt:         time.Now(),
				UpdatedAt:         time.Now(),
			},
			wantError:             true,
			expectedErrorContains: "total score cannot be negative",
		},
		{
			name: "Invalid - zero max score",
			evaluation: Evaluation{
				ID:                "eval-1",
				StudentID:         "student-1",
				RubricID:          "rubric-1",
				EvidenceID:        "evidence-1",
				UserID:            "user-1",
				PerformanceScores: map[string]interface{}{"criterion1": 85},
				TotalScore:        85,
				MaxScore:          0,
				PerformanceLevel:  EvaluationPerformanceLevelProficient,
				RevisionNo:        1,
				IsCurrentVersion:  true,
				EvaluatedAt:       time.Now(),
				CreatedAt:         time.Now(),
				UpdatedAt:         time.Now(),
			},
			wantError:             true,
			expectedErrorContains: "max score must be at least 1",
		},
		{
			name: "Invalid - negative max score",
			evaluation: Evaluation{
				ID:                "eval-1",
				StudentID:         "student-1",
				RubricID:          "rubric-1",
				EvidenceID:        "evidence-1",
				UserID:            "user-1",
				PerformanceScores: map[string]interface{}{"criterion1": 85},
				TotalScore:        85,
				MaxScore:          -100,
				PerformanceLevel:  EvaluationPerformanceLevelProficient,
				RevisionNo:        1,
				IsCurrentVersion:  true,
				EvaluatedAt:       time.Now(),
				CreatedAt:         time.Now(),
				UpdatedAt:         time.Now(),
			},
			wantError:             true,
			expectedErrorContains: "max score must be at least 1",
		},
		{
			name: "Invalid - total score greater than max score",
			evaluation: Evaluation{
				ID:                "eval-1",
				StudentID:         "student-1",
				RubricID:          "rubric-1",
				EvidenceID:        "evidence-1",
				UserID:            "user-1",
				PerformanceScores: map[string]interface{}{"criterion1": 85},
				TotalScore:        150,
				MaxScore:          100,
				PerformanceLevel:  EvaluationPerformanceLevelProficient,
				RevisionNo:        1,
				IsCurrentVersion:  true,
				EvaluatedAt:       time.Now(),
				CreatedAt:         time.Now(),
				UpdatedAt:         time.Now(),
			},
			wantError:             true,
			expectedErrorContains: "total score cannot exceed max score",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.evaluation.IsValidRevision()

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

// TestEvaluation_GetCurrentRevision tests if evaluation is current revision
func TestEvaluation_GetCurrentRevision(t *testing.T) {
	tests := []struct {
		name       string
		evaluation Evaluation
		expected   bool
	}{
		{
			name: "Is current version",
			evaluation: Evaluation{
				ID:               "eval-1",
				IsCurrentVersion: true,
			},
			expected: true,
		},
		{
			name: "Is not current version",
			evaluation: Evaluation{
				ID:               "eval-1",
				IsCurrentVersion: false,
			},
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.evaluation.GetCurrentRevision()
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

// TestEvaluation_IsFirstRevision tests if evaluation is first revision
func TestEvaluation_IsFirstRevision(t *testing.T) {
	tests := []struct {
		name       string
		evaluation Evaluation
		expected   bool
	}{
		{
			name: "Is first revision",
			evaluation: Evaluation{
				ID:               "eval-1",
				RevisionNo:       1,
				ParentRevisionID: nil,
			},
			expected: true,
		},
		{
			name: "Not first revision - higher revision number",
			evaluation: Evaluation{
				ID:               "eval-1",
				RevisionNo:       2,
				ParentRevisionID: makeStringPtr("eval-1"),
			},
			expected: false,
		},
		{
			name: "Not first revision - has parent ID",
			evaluation: Evaluation{
				ID:               "eval-1",
				RevisionNo:       1,
				ParentRevisionID: makeStringPtr("eval-0"),
			},
			expected: false,
		},
		{
			name: "Not first revision - both conditions",
			evaluation: Evaluation{
				ID:               "eval-1",
				RevisionNo:       3,
				ParentRevisionID: makeStringPtr("eval-2"),
			},
			expected: false,
		},
		{
			name: "Edge case - revision 1 but has parent",
			evaluation: Evaluation{
				ID:               "eval-1",
				RevisionNo:       1,
				ParentRevisionID: makeStringPtr("parent-eval"),
			},
			expected: false,
		},
		{
			name: "Edge case - no parent but revision 2",
			evaluation: Evaluation{
				ID:               "eval-1",
				RevisionNo:       2,
				ParentRevisionID: nil,
			},
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.evaluation.IsFirstRevision()
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

// TestEvaluation_CanBeRevised tests if evaluation can be revised
func TestEvaluation_CanBeRevised(t *testing.T) {
	tests := []struct {
		name       string
		evaluation Evaluation
		expected   bool
	}{
		{
			name: "Can be revised - current version",
			evaluation: Evaluation{
				ID:               "eval-1",
				IsCurrentVersion: true,
			},
			expected: true,
		},
		{
			name: "Cannot be revised - not current version",
			evaluation: Evaluation{
				ID:               "eval-1",
				IsCurrentVersion: false,
			},
			expected: false,
		},
		{
			name: "Cannot be revised - archived",
			evaluation: Evaluation{
				ID:               "eval-1",
				IsCurrentVersion: false,
				ParentRevisionID: makeStringPtr("parent-eval"),
			},
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.evaluation.CanBeRevised()
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

// TestEvaluation_Archive tests archiving evaluations
func TestEvaluation_Archive(t *testing.T) {
	tests := []struct {
		name                  string
		evaluation            Evaluation
		wantError             bool
		expectedErrorContains string
		expectedIsCurrent     bool
	}{
		{
			name: "Valid archive",
			evaluation: Evaluation{
				ID:               "eval-1",
				StudentID:        "student-1",
				RubricID:         "rubric-1",
				EvidenceID:       "evidence-1",
				UserID:           "user-1",
				TotalScore:       85,
				MaxScore:         100,
				PerformanceLevel: EvaluationPerformanceLevelProficient,
				RevisionNo:       1,
				IsCurrentVersion: true,
				EvaluatedAt:      time.Now(),
				CreatedAt:        time.Now(),
				UpdatedAt:        time.Now(),
			},
			wantError:         false,
			expectedIsCurrent: false,
		},
		{
			name: "Invalid - already archived",
			evaluation: Evaluation{
				ID:               "eval-1",
				StudentID:        "student-1",
				RubricID:         "rubric-1",
				EvidenceID:       "evidence-1",
				UserID:           "user-1",
				TotalScore:       85,
				MaxScore:         100,
				PerformanceLevel: EvaluationPerformanceLevelProficient,
				RevisionNo:       2,
				IsCurrentVersion: false,
				EvaluatedAt:      time.Now(),
				CreatedAt:        time.Now(),
				UpdatedAt:        time.Now(),
			},
			wantError:             true,
			expectedErrorContains: "evaluation is already archived",
			expectedIsCurrent:     false,
		},
		{
			name: "Valid archive - update timestamp",
			evaluation: Evaluation{
				ID:               "eval-1",
				StudentID:        "student-1",
				RubricID:         "rubric-1",
				EvidenceID:       "evidence-1",
				UserID:           "user-1",
				TotalScore:       85,
				MaxScore:         100,
				PerformanceLevel: EvaluationPerformanceLevelProficient,
				RevisionNo:       1,
				IsCurrentVersion: true,
				EvaluatedAt:      time.Now(),
				CreatedAt:        time.Now(),
				UpdatedAt:        time.Now().Add(-1 * time.Hour),
			},
			wantError:         false,
			expectedIsCurrent: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			originalUpdatedAt := tc.evaluation.UpdatedAt
			err := tc.evaluation.Archive()

			if tc.wantError && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tc.wantError && err != nil {
				t.Errorf("expected no error, got: %v", err)
			}

			if tc.evaluation.IsCurrentVersion != tc.expectedIsCurrent {
				t.Errorf("expected IsCurrentVersion %v, got %v", tc.expectedIsCurrent, tc.evaluation.IsCurrentVersion)
			}

			if !tc.wantError {
				if tc.evaluation.UpdatedAt.Before(originalUpdatedAt) || tc.evaluation.UpdatedAt.Equal(originalUpdatedAt) {
					t.Error("UpdatedAt should be updated after archive")
				}
			}

			if tc.wantError && err != nil {
				if tc.expectedErrorContains != "" && !contains(err.Error(), tc.expectedErrorContains) {
					t.Errorf("expected error to contain %s, got %s", tc.expectedErrorContains, err.Error())
				}
			}
		})
	}
}

// TestEvaluation_HasFeedbackChanged tests feedback change detection
func TestEvaluation_HasFeedbackChanged(t *testing.T) {
	tests := []struct {
		name        string
		oldFeedback *string
		newFeedback *string
		expected    bool
	}{
		{
			name:        "Both nil - no change",
			oldFeedback: nil,
			newFeedback: nil,
			expected:    false,
		},
		{
			name:        "Old nil, new not nil - change",
			oldFeedback: nil,
			newFeedback: makeStringPtr("New feedback"),
			expected:    true,
		},
		{
			name:        "Old not nil, new nil - change",
			oldFeedback: makeStringPtr("Old feedback"),
			newFeedback: nil,
			expected:    true,
		},
		{
			name:        "Both not nil, same value - no change",
			oldFeedback: makeStringPtr("Same feedback"),
			newFeedback: makeStringPtr("Same feedback"),
			expected:    false,
		},
		{
			name:        "Both not nil, different value - change",
			oldFeedback: makeStringPtr("Old feedback"),
			newFeedback: makeStringPtr("New feedback"),
			expected:    true,
		},
		{
			name:        "Both not nil, case sensitive - change",
			oldFeedback: makeStringPtr("Feedback"),
			newFeedback: makeStringPtr("feedback"),
			expected:    true,
		},
		{
			name:        "Both not nil, whitespace difference - change",
			oldFeedback: makeStringPtr("Feedback"),
			newFeedback: makeStringPtr("Feedback "),
			expected:    true,
		},
		{
			name:        "Both empty strings - no change",
			oldFeedback: makeStringPtr(""),
			newFeedback: makeStringPtr(""),
			expected:    false,
		},
		{
			name:        "Old empty, new not empty - change",
			oldFeedback: makeStringPtr(""),
			newFeedback: makeStringPtr("Feedback"),
			expected:    true,
		},
		{
			name:        "Old not empty, new empty - change",
			oldFeedback: makeStringPtr("Feedback"),
			newFeedback: makeStringPtr(""),
			expected:    true,
		},
		{
			name:        "Both whitespace - no change",
			oldFeedback: makeStringPtr("   "),
			newFeedback: makeStringPtr("   "),
			expected:    false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			evaluation := Evaluation{}
			result := evaluation.HasFeedbackChanged(tc.oldFeedback, tc.newFeedback)
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

// Helper functions
func makeStringPtr(s string) *string {
	return &s
}

func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if i+len(substr) > len(s) {
			continue
		}
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

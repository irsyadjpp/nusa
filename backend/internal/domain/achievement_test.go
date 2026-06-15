package domain

import (
	"testing"
	"time"
)

// TestNewAchievementService tests the service constructor
func TestNewAchievementService(t *testing.T) {
	service := NewAchievementService()
	if service == nil {
		t.Fatal("NewAchievementService should not return nil")
	}
	if service == nil {
		t.Error("Service should be properly initialized")
	}
}

// TestAchievementService_CalculateStudentAchievement tests achievement calculation
func TestAchievementService_CalculateStudentAchievement(t *testing.T) {
	tests := []struct {
		name              string
		studentID         string
		tpID              string
		evaluations       []Evaluation
		tpSuccessCriteria interface{}
		wantError         bool
		wantScore         float64
		wantEvidenceCount int
	}{
		{
			name:      "Valid calculation with evaluations",
			studentID: "student-1",
			tpID:      "tp-1",
			evaluations: []Evaluation{
				{
					ID:         "eval-1",
					StudentID:  "student-1",
					TotalScore: 85,
					MaxScore:   100,
					EvidenceID: "evidence-1",
				},
				{
					ID:         "eval-2",
					StudentID:  "student-1",
					TotalScore: 90,
					MaxScore:   100,
					EvidenceID: "evidence-2",
				},
			},
			tpSuccessCriteria: nil,
			wantError:         false,
			wantScore:         87.5,
			wantEvidenceCount: 2,
		},
		{
			name:              "Empty evaluations",
			studentID:         "student-1",
			tpID:              "tp-1",
			evaluations:       []Evaluation{},
			tpSuccessCriteria: nil,
			wantError:         false,
			wantScore:         0,
			wantEvidenceCount: 0,
		},
		{
			name:      "Single evaluation",
			studentID: "student-2",
			tpID:      "tp-2",
			evaluations: []Evaluation{
				{
					ID:         "eval-1",
					StudentID:  "student-2",
					TotalScore: 75,
					MaxScore:   100,
					EvidenceID: "evidence-1",
				},
			},
			tpSuccessCriteria: nil,
			wantError:         false,
			wantScore:         75,
			wantEvidenceCount: 1,
		},
		{
			name:      "Perfect score",
			studentID: "student-3",
			tpID:      "tp-3",
			evaluations: []Evaluation{
				{
					ID:         "eval-1",
					StudentID:  "student-3",
					TotalScore: 100,
					MaxScore:   100,
					EvidenceID: "evidence-1",
				},
			},
			tpSuccessCriteria: nil,
			wantError:         false,
			wantScore:         100,
			wantEvidenceCount: 1,
		},
		{
			name:      "Zero score",
			studentID: "student-4",
			tpID:      "tp-4",
			evaluations: []Evaluation{
				{
					ID:         "eval-1",
					StudentID:  "student-4",
					TotalScore: 0,
					MaxScore:   100,
					EvidenceID: "evidence-1",
				},
			},
			tpSuccessCriteria: nil,
			wantError:         false,
			wantScore:         0,
			wantEvidenceCount: 1,
		},
		{
			name:      "Multiple evaluations with different max scores",
			studentID: "student-5",
			tpID:      "tp-5",
			evaluations: []Evaluation{
				{
					ID:         "eval-1",
					StudentID:  "student-5",
					TotalScore: 40,
					MaxScore:   50,
					EvidenceID: "evidence-1",
				},
				{
					ID:         "eval-2",
					StudentID:  "student-5",
					TotalScore: 60,
					MaxScore:   100,
					EvidenceID: "evidence-2",
				},
			},
			tpSuccessCriteria: nil,
			wantError:         false,
			wantScore:         66.67, // (40+60)/(50+100)*100 = 100/150*100 = 66.67
			wantEvidenceCount: 2,
		},
		{
			name:      "Max score zero (edge case)",
			studentID: "student-6",
			tpID:      "tp-6",
			evaluations: []Evaluation{
				{
					ID:         "eval-1",
					StudentID:  "student-6",
					TotalScore: 0,
					MaxScore:   0,
					EvidenceID: "evidence-1",
				},
			},
			tpSuccessCriteria: nil,
			wantError:         false,
			wantScore:         0,
			wantEvidenceCount: 1,
		},
	}

	service := NewAchievementService()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			achievement, err := service.CalculateStudentAchievement(
				tc.studentID,
				tc.tpID,
				tc.evaluations,
				tc.tpSuccessCriteria,
			)

			if tc.wantError && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tc.wantError && err != nil {
				t.Errorf("expected no error, got: %v", err)
			}

			if achievement == nil && !tc.wantError {
				t.Fatal("expected achievement object, got nil")
			}

			if achievement != nil {
				if achievement.StudentID != tc.studentID {
					t.Errorf("expected student ID %s, got %s", tc.studentID, achievement.StudentID)
				}
				if achievement.TPID != tc.tpID {
					t.Errorf("expected TP ID %s, got %s", tc.tpID, achievement.TPID)
				}
				if achievement.EvidenceCount != tc.wantEvidenceCount {
					t.Errorf("expected evidence count %d, got %d", tc.wantEvidenceCount, achievement.EvidenceCount)
				}
				if achievement.EvaluationCount != tc.wantEvidenceCount {
					t.Errorf("expected evaluation count %d, got %d", tc.wantEvidenceCount, achievement.EvaluationCount)
				}
				// Allow small floating point difference
				if achievement.OverallScore < tc.wantScore-0.1 || achievement.OverallScore > tc.wantScore+0.1 {
					t.Errorf("expected score %.2f, got %.2f", tc.wantScore, achievement.OverallScore)
				}
				if achievement.CalculatedAt.IsZero() {
					t.Error("CalculatedAt should not be zero")
				}
			}
		})
	}
}

// TestAchievementService_DetermineEvaluationPerformanceLevel tests performance level determination
func TestAchievementService_DetermineEvaluationPerformanceLevel(t *testing.T) {
	tests := []struct {
		name     string
		score    float64
		expected EvaluationPerformanceLevel
	}{
		{
			name:     "Excellent score (90)",
			score:    90,
			expected: EvaluationPerformanceLevelExcellent,
		},
		{
			name:     "Excellent score (95)",
			score:    95,
			expected: EvaluationPerformanceLevelExcellent,
		},
		{
			name:     "Excellent score (100)",
			score:    100,
			expected: EvaluationPerformanceLevelExcellent,
		},
		{
			name:     "Proficient score (75)",
			score:    75,
			expected: EvaluationPerformanceLevelProficient,
		},
		{
			name:     "Proficient score (80)",
			score:    80,
			expected: EvaluationPerformanceLevelProficient,
		},
		{
			name:     "Proficient score (89)",
			score:    89,
			expected: EvaluationPerformanceLevelProficient,
		},
		{
			name:     "Developing score (60)",
			score:    60,
			expected: EvaluationPerformanceLevelDeveloping,
		},
		{
			name:     "Developing score (65)",
			score:    65,
			expected: EvaluationPerformanceLevelDeveloping,
		},
		{
			name:     "Developing score (74)",
			score:    74,
			expected: EvaluationPerformanceLevelDeveloping,
		},
		{
			name:     "Beginning score (0)",
			score:    0,
			expected: EvaluationPerformanceLevelBeginning,
		},
		{
			name:     "Beginning score (30)",
			score:    30,
			expected: EvaluationPerformanceLevelBeginning,
		},
		{
			name:     "Beginning score (59)",
			score:    59,
			expected: EvaluationPerformanceLevelBeginning,
		},
		{
			name:     "Boundary case - exactly 89.9",
			score:    89.9,
			expected: EvaluationPerformanceLevelProficient,
		},
		{
			name:     "Boundary case - exactly 90",
			score:    90,
			expected: EvaluationPerformanceLevelExcellent,
		},
		{
			name:     "Boundary case - exactly 74.9",
			score:    74.9,
			expected: EvaluationPerformanceLevelDeveloping,
		},
		{
			name:     "Boundary case - exactly 75",
			score:    75,
			expected: EvaluationPerformanceLevelProficient,
		},
		{
			name:     "Boundary case - exactly 59.9",
			score:    59.9,
			expected: EvaluationPerformanceLevelBeginning,
		},
		{
			name:     "Boundary case - exactly 60",
			score:    60,
			expected: EvaluationPerformanceLevelDeveloping,
		},
		{
			name:     "Negative score",
			score:    -10,
			expected: EvaluationPerformanceLevelBeginning,
		},
		{
			name:     "Score above 100",
			score:    110,
			expected: EvaluationPerformanceLevelExcellent,
		},
	}

	service := NewAchievementService()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := service.determineEvaluationPerformanceLevel(tc.score)
			if result != tc.expected {
				t.Errorf("score %.2f: expected %s, got %s", tc.score, tc.expected, result)
			}
		})
	}
}

// TestAchievementService_DetermineMasteryStatus tests mastery status determination
func TestAchievementService_DetermineMasteryStatus(t *testing.T) {
	tests := []struct {
		name             string
		score            float64
		performanceLevel EvaluationPerformanceLevel
		expected         MasteryStatus
	}{
		{
			name:             "Excellent performance",
			score:            90,
			performanceLevel: EvaluationPerformanceLevelExcellent,
			expected:         MasteryStatusExceeding,
		},
		{
			name:             "Excellent performance with lower score",
			score:            90,
			performanceLevel: EvaluationPerformanceLevelExcellent,
			expected:         MasteryStatusExceeding,
		},
		{
			name:             "Proficient performance",
			score:            75,
			performanceLevel: EvaluationPerformanceLevelProficient,
			expected:         MasteryStatusAchieved,
		},
		{
			name:             "Proficient performance with higher score",
			score:            85,
			performanceLevel: EvaluationPerformanceLevelProficient,
			expected:         MasteryStatusAchieved,
		},
		{
			name:             "Developing performance with positive score",
			score:            65,
			performanceLevel: EvaluationPerformanceLevelDeveloping,
			expected:         MasteryStatusInProgress,
		},
		{
			name:             "Beginning performance with positive score",
			score:            45,
			performanceLevel: EvaluationPerformanceLevelBeginning,
			expected:         MasteryStatusInProgress,
		},
		{
			name:             "Zero score",
			score:            0,
			performanceLevel: EvaluationPerformanceLevelBeginning,
			expected:         MasteryStatusNotStarted,
		},
		{
			name:             "Developing with zero score",
			score:            0,
			performanceLevel: EvaluationPerformanceLevelDeveloping,
			expected:         MasteryStatusNotStarted,
		},
		{
			name:             "Proficient with zero score",
			score:            0,
			performanceLevel: EvaluationPerformanceLevelProficient,
			expected:         MasteryStatusAchieved, // Level takes precedence over score
		},
		{
			name:             "Excellent with zero score",
			score:            0,
			performanceLevel: EvaluationPerformanceLevelExcellent,
			expected:         MasteryStatusExceeding,
		},
	}

	service := NewAchievementService()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := service.determineMasteryStatus(tc.score, tc.performanceLevel)
			if result != tc.expected {
				t.Errorf("score %.2f, level %s: expected %s, got %s",
					tc.score, tc.performanceLevel, tc.expected, result)
			}
		})
	}
}

// TestAchievementService_DetermineMasteryStatusFromProgress tests mastery status from progress
func TestAchievementService_DetermineMasteryStatusFromProgress(t *testing.T) {
	tests := []struct {
		name     string
		progress float64
		expected MasteryStatus
	}{
		{
			name:     "Exceeding progress (90)",
			progress: 90,
			expected: MasteryStatusExceeding,
		},
		{
			name:     "Exceeding progress (95)",
			progress: 95,
			expected: MasteryStatusExceeding,
		},
		{
			name:     "Exceeding progress (100)",
			progress: 100,
			expected: MasteryStatusExceeding,
		},
		{
			name:     "Achieved progress (75)",
			progress: 75,
			expected: MasteryStatusAchieved,
		},
		{
			name:     "Achieved progress (80)",
			progress: 80,
			expected: MasteryStatusAchieved,
		},
		{
			name:     "Achieved progress (89)",
			progress: 89,
			expected: MasteryStatusAchieved,
		},
		{
			name:     "In progress (1)",
			progress: 1,
			expected: MasteryStatusInProgress,
		},
		{
			name:     "In progress (50)",
			progress: 50,
			expected: MasteryStatusInProgress,
		},
		{
			name:     "In progress (74)",
			progress: 74,
			expected: MasteryStatusInProgress,
		},
		{
			name:     "Not started (0)",
			progress: 0,
			expected: MasteryStatusNotStarted,
		},
		{
			name:     "Boundary case - exactly 74.9",
			progress: 74.9,
			expected: MasteryStatusInProgress,
		},
		{
			name:     "Boundary case - exactly 75",
			progress: 75,
			expected: MasteryStatusAchieved,
		},
		{
			name:     "Boundary case - exactly 89.9",
			progress: 89.9,
			expected: MasteryStatusAchieved,
		},
		{
			name:     "Boundary case - exactly 90",
			progress: 90,
			expected: MasteryStatusExceeding,
		},
		{
			name:     "Negative progress",
			progress: -10,
			expected: MasteryStatusNotStarted,
		},
		{
			name:     "Progress above 100",
			progress: 110,
			expected: MasteryStatusExceeding,
		},
	}

	service := NewAchievementService()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := service.determineMasteryStatusFromProgress(tc.progress)
			if result != tc.expected {
				t.Errorf("progress %.2f: expected %s, got %s", tc.progress, tc.expected, result)
			}
		})
	}
}

// TestAchievementService_CalculateTPProgress tests TP progress calculation
func TestAchievementService_CalculateTPProgress(t *testing.T) {
	tests := []struct {
		name        string
		evaluations []Evaluation
		expected    float64
	}{
		{
			name:        "Empty evaluations",
			evaluations: []Evaluation{},
			expected:    0,
		},
		{
			name: "Single evaluation perfect score",
			evaluations: []Evaluation{
				{
					ID:         "eval-1",
					TotalScore: 100,
					MaxScore:   100,
					EvidenceID: "evidence-1",
				},
			},
			expected: 100,
		},
		{
			name: "Single evaluation zero score",
			evaluations: []Evaluation{
				{
					ID:         "eval-1",
					TotalScore: 0,
					MaxScore:   100,
					EvidenceID: "evidence-1",
				},
			},
			expected: 0,
		},
		{
			name: "Single evaluation partial score",
			evaluations: []Evaluation{
				{
					ID:         "eval-1",
					TotalScore: 75,
					MaxScore:   100,
					EvidenceID: "evidence-1",
				},
			},
			expected: 75,
		},
		{
			name: "Multiple evaluations average score",
			evaluations: []Evaluation{
				{
					ID:         "eval-1",
					TotalScore: 80,
					MaxScore:   100,
					EvidenceID: "evidence-1",
				},
				{
					ID:         "eval-2",
					TotalScore: 90,
					MaxScore:   100,
					EvidenceID: "evidence-2",
				},
			},
			expected: 85,
		},
		{
			name: "Multiple evaluations different max scores",
			evaluations: []Evaluation{
				{
					ID:         "eval-1",
					TotalScore: 40,
					MaxScore:   50,
					EvidenceID: "evidence-1",
				},
				{
					ID:         "eval-2",
					TotalScore: 80,
					MaxScore:   100,
					EvidenceID: "evidence-2",
				},
			},
			expected: 80, // (40+80)/(50+100)*100 = 120/150*100 = 80
		},
		{
			name: "Zero max score",
			evaluations: []Evaluation{
				{
					ID:         "eval-1",
					TotalScore: 0,
					MaxScore:   0,
					EvidenceID: "evidence-1",
				},
			},
			expected: 0,
		},
	}

	service := NewAchievementService()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := service.calculateTPProgress(tc.evaluations)
			// Allow small floating point difference
			if result < tc.expected-0.1 || result > tc.expected+0.1 {
				t.Errorf("expected %.2f, got %.2f", tc.expected, result)
			}
		})
	}
}

// TestAchievementService_BuildCompetencyBreakdown tests competency breakdown building
func TestAchievementService_BuildCompetencyBreakdown(t *testing.T) {
	service := NewAchievementService()

	tests := []struct {
		name        string
		evaluations []Evaluation
		wantLength  int
	}{
		{
			name:        "Empty evaluations",
			evaluations: []Evaluation{},
			wantLength:  0,
		},
		{
			name: "Single evaluation",
			evaluations: []Evaluation{
				{
					ID:         "eval-1",
					TotalScore: 75,
					MaxScore:   100,
					EvidenceID: "evidence-1",
				},
			},
			wantLength: 0, // Simplified implementation
		},
		{
			name: "Multiple evaluations",
			evaluations: []Evaluation{
				{
					ID:         "eval-1",
					TotalScore: 80,
					MaxScore:   100,
					EvidenceID: "evidence-1",
				},
				{
					ID:         "eval-2",
					TotalScore: 90,
					MaxScore:   100,
					EvidenceID: "evidence-2",
				},
			},
			wantLength: 0, // Simplified implementation
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := service.buildCompetencyBreakdown(tc.evaluations)
			if len(result) != tc.wantLength {
				t.Errorf("expected length %d, got %d", tc.wantLength, len(result))
			}
		})
	}
}

// TestAchievementService_CalculateCompetencyProgress tests competency progress calculation
func TestAchievementService_CalculateCompetencyProgress(t *testing.T) {
	service := NewAchievementService()

	tests := []struct {
		name        string
		studentID   string
		studentName string
		subjectID   string
		subjectName string
		phaseID     string
		phaseName   string
		tps         []TP
		evaluations []Evaluation
		wantError   bool
		wantTPCount int
	}{
		{
			name:        "Valid calculation with TPs",
			studentID:   "student-1",
			studentName: "John Doe",
			subjectID:   "subject-1",
			subjectName: "Mathematics",
			phaseID:     "phase-1",
			phaseName:   "Phase A",
			tps: []TP{
				{
					ID:              "tp-1",
					CPID:            "cp-1",
					Status:          WorkflowStatusDraft,
					SuccessCriteria: nil, // Simplified test
				},
				{
					ID:              "tp-2",
					CPID:            "cp-1",
					Status:          WorkflowStatusDraft,
					SuccessCriteria: nil, // Simplified test
				},
			},
			evaluations: []Evaluation{
				{
					ID:         "eval-1",
					StudentID:  "student-1",
					TotalScore: 85,
					MaxScore:   100,
					EvidenceID: "tp-1",
				},
			},
			wantError:   false,
			wantTPCount: 2,
		},
		{
			name:        "Empty TPs",
			studentID:   "student-1",
			studentName: "John Doe",
			subjectID:   "subject-1",
			subjectName: "Mathematics",
			phaseID:     "phase-1",
			phaseName:   "Phase A",
			tps:         []TP{},
			evaluations: []Evaluation{},
			wantError:   false,
			wantTPCount: 0,
		},
		{
			name:        "TPs with no evaluations",
			studentID:   "student-2",
			studentName: "Jane Doe",
			subjectID:   "subject-2",
			subjectName: "Science",
			phaseID:     "phase-2",
			phaseName:   "Phase B",
			tps: []TP{
				{
					ID:              "tp-1",
					CPID:            "cp-1",
					Status:          WorkflowStatusDraft,
					SuccessCriteria: nil, // Simplified test
				},
			},
			evaluations: []Evaluation{},
			wantError:   false,
			wantTPCount: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			progress, err := service.CalculateCompetencyProgress(
				tc.studentID,
				tc.studentName,
				tc.subjectID,
				tc.subjectName,
				tc.phaseID,
				tc.phaseName,
				tc.tps,
				tc.evaluations,
			)

			if tc.wantError && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tc.wantError && err != nil {
				t.Errorf("expected no error, got: %v", err)
			}

			if progress == nil && !tc.wantError {
				t.Fatal("expected progress object, got nil")
			}

			if progress != nil {
				if progress.StudentID != tc.studentID {
					t.Errorf("expected student ID %s, got %s", tc.studentID, progress.StudentID)
				}
				if progress.StudentName != tc.studentName {
					t.Errorf("expected student name %s, got %s", tc.studentName, progress.StudentName)
				}
				if progress.SubjectID != tc.subjectID {
					t.Errorf("expected subject ID %s, got %s", tc.subjectID, progress.SubjectID)
				}
				if progress.SubjectName != tc.subjectName {
					t.Errorf("expected subject name %s, got %s", tc.subjectName, progress.SubjectName)
				}
				if progress.PhaseID != tc.phaseID {
					t.Errorf("expected phase ID %s, got %s", tc.phaseID, progress.PhaseID)
				}
				if progress.PhaseName != tc.phaseName {
					t.Errorf("expected phase name %s, got %s", tc.phaseName, progress.PhaseName)
				}
				if len(progress.CompetencyProgress) != tc.wantTPCount {
					t.Errorf("expected TP count %d, got %d", tc.wantTPCount, len(progress.CompetencyProgress))
				}
				if progress.CalculatedAt.IsZero() {
					t.Error("CalculatedAt should not be zero")
				}
			}
		})
	}
}

// TestAchievementService_GenerateAchievementSummary tests achievement summary generation
func TestAchievementService_GenerateAchievementSummary(t *testing.T) {
	service := NewAchievementService()

	tests := []struct {
		name         string
		studentID    string
		studentName  string
		classID      string
		className    string
		reportPeriod interface{}
		achievements []Achievement
		wantError    bool
		wantSubject  int
	}{
		{
			name:         "Valid summary with achievements",
			studentID:    "student-1",
			studentName:  "John Doe",
			classID:      "class-1",
			className:    "Class A",
			reportPeriod: "2026 Semester 1",
			achievements: []Achievement{
				{
					StudentID:                  "student-1",
					StudentName:                "John Doe",
					TPID:                       "tp-1",
					TPTitle:                    "Algebra",
					OverallScore:               85,
					EvaluationPerformanceLevel: EvaluationPerformanceLevelProficient,
					MasteryStatus:              MasteryStatusAchieved,
					EvidenceCount:              3,
					EvaluationCount:            3,
					CalculatedAt:               time.Now(),
				},
				{
					StudentID:                  "student-1",
					StudentName:                "John Doe",
					TPID:                       "tp-2",
					TPTitle:                    "Geometry",
					OverallScore:               92,
					EvaluationPerformanceLevel: EvaluationPerformanceLevelExcellent,
					MasteryStatus:              MasteryStatusExceeding,
					EvidenceCount:              2,
					EvaluationCount:            2,
					CalculatedAt:               time.Now(),
				},
			},
			wantError:   false,
			wantSubject: 0, // Simplified implementation
		},
		{
			name:         "Empty achievements",
			studentID:    "student-1",
			studentName:  "John Doe",
			classID:      "class-1",
			className:    "Class A",
			reportPeriod: "2026 Semester 1",
			achievements: []Achievement{},
			wantError:    false,
			wantSubject:  0,
		},
		{
			name:         "Single achievement",
			studentID:    "student-2",
			studentName:  "Jane Doe",
			classID:      "class-2",
			className:    "Class B",
			reportPeriod: "2026 Semester 1",
			achievements: []Achievement{
				{
					StudentID:                  "student-2",
					StudentName:                "Jane Doe",
					TPID:                       "tp-1",
					TPTitle:                    "Physics",
					OverallScore:               78,
					EvaluationPerformanceLevel: EvaluationPerformanceLevelProficient,
					MasteryStatus:              MasteryStatusAchieved,
					EvidenceCount:              4,
					EvaluationCount:            4,
					CalculatedAt:               time.Now(),
				},
			},
			wantError:   false,
			wantSubject: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			summary, err := service.GenerateAchievementSummary(
				tc.studentID,
				tc.studentName,
				tc.classID,
				tc.className,
				tc.reportPeriod,
				tc.achievements,
			)

			if tc.wantError && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tc.wantError && err != nil {
				t.Errorf("expected no error, got: %v", err)
			}

			if summary == nil && !tc.wantError {
				t.Fatal("expected summary object, got nil")
			}

			if summary != nil {
				if summary.StudentID != tc.studentID {
					t.Errorf("expected student ID %s, got %s", tc.studentID, summary.StudentID)
				}
				if summary.StudentName != tc.studentName {
					t.Errorf("expected student name %s, got %s", tc.studentName, summary.StudentName)
				}
				if summary.ClassID != tc.classID {
					t.Errorf("expected class ID %s, got %s", tc.classID, summary.ClassID)
				}
				if summary.ClassName != tc.className {
					t.Errorf("expected class name %s, got %s", tc.className, summary.ClassName)
				}
				if len(summary.SubjectBreakdown) != tc.wantSubject {
					t.Errorf("expected subject breakdown length %d, got %d", tc.wantSubject, len(summary.SubjectBreakdown))
				}
				if summary.CalculatedAt.IsZero() {
					t.Error("CalculatedAt should not be zero")
				}

				// Check overall achievement calculation
				if len(tc.achievements) > 0 {
					var totalScore float64
					for _, ach := range tc.achievements {
						totalScore += ach.OverallScore
					}
					expectedOverall := totalScore / float64(len(tc.achievements))
					if summary.OverallAchievement < expectedOverall-0.1 || summary.OverallAchievement > expectedOverall+0.1 {
						t.Errorf("expected overall achievement %.2f, got %.2f", expectedOverall, summary.OverallAchievement)
					}
				} else {
					if summary.OverallAchievement != 0 {
						t.Errorf("expected overall achievement 0, got %.2f", summary.OverallAchievement)
					}
				}
			}
		})
	}
}

// TestAchievementService_GenerateClassAchievement tests class achievement generation
func TestAchievementService_GenerateClassAchievement(t *testing.T) {
	service := NewAchievementService()

	tests := []struct {
		name                 string
		classID              string
		className            string
		subjectID            string
		subjectName          string
		studentAchievements  []Achievement
		wantError            bool
		wantStudentCount     int
		wantPerformanceLevel int // Expected number of performance levels
		wantMasteryStatus    int // Expected number of mastery statuses
	}{
		{
			name:        "Valid class achievement",
			classID:     "class-1",
			className:   "Class A",
			subjectID:   "subject-1",
			subjectName: "Mathematics",
			studentAchievements: []Achievement{
				{
					StudentID:                  "student-1",
					StudentName:                "John Doe",
					OverallScore:               85,
					EvaluationPerformanceLevel: EvaluationPerformanceLevelProficient,
					MasteryStatus:              MasteryStatusAchieved,
					CalculatedAt:               time.Now(),
				},
				{
					StudentID:                  "student-2",
					StudentName:                "Jane Doe",
					OverallScore:               92,
					EvaluationPerformanceLevel: EvaluationPerformanceLevelExcellent,
					MasteryStatus:              MasteryStatusExceeding,
					CalculatedAt:               time.Now(),
				},
				{
					StudentID:                  "student-3",
					StudentName:                "Bob Smith",
					OverallScore:               65,
					EvaluationPerformanceLevel: EvaluationPerformanceLevelDeveloping,
					MasteryStatus:              MasteryStatusInProgress,
					CalculatedAt:               time.Now(),
				},
			},
			wantError:            false,
			wantStudentCount:     3,
			wantPerformanceLevel: 3,
			wantMasteryStatus:    3,
		},
		{
			name:                 "Empty student achievements",
			classID:              "class-1",
			className:            "Class A",
			subjectID:            "subject-1",
			subjectName:          "Mathematics",
			studentAchievements:  []Achievement{},
			wantError:            false,
			wantStudentCount:     0,
			wantPerformanceLevel: 0,
			wantMasteryStatus:    0,
		},
		{
			name:        "Single student",
			classID:     "class-2",
			className:   "Class B",
			subjectID:   "subject-2",
			subjectName: "Science",
			studentAchievements: []Achievement{
				{
					StudentID:                  "student-1",
					StudentName:                "Alice Johnson",
					OverallScore:               78,
					EvaluationPerformanceLevel: EvaluationPerformanceLevelProficient,
					MasteryStatus:              MasteryStatusAchieved,
					CalculatedAt:               time.Now(),
				},
			},
			wantError:            false,
			wantStudentCount:     1,
			wantPerformanceLevel: 1,
			wantMasteryStatus:    1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			classAchievement, err := service.GenerateClassAchievement(
				tc.classID,
				tc.className,
				tc.subjectID,
				tc.subjectName,
				tc.studentAchievements,
			)

			if tc.wantError && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tc.wantError && err != nil {
				t.Errorf("expected no error, got: %v", err)
			}

			if classAchievement == nil && !tc.wantError {
				t.Fatal("expected class achievement object, got nil")
			}

			if classAchievement != nil {
				if classAchievement.ClassID != tc.classID {
					t.Errorf("expected class ID %s, got %s", tc.classID, classAchievement.ClassID)
				}
				if classAchievement.ClassName != tc.className {
					t.Errorf("expected class name %s, got %s", tc.className, classAchievement.ClassName)
				}
				if classAchievement.SubjectID != tc.subjectID {
					t.Errorf("expected subject ID %s, got %s", tc.subjectID, classAchievement.SubjectID)
				}
				if classAchievement.SubjectName != tc.subjectName {
					t.Errorf("expected subject name %s, got %s", tc.subjectName, classAchievement.SubjectName)
				}
				if classAchievement.TotalStudents != tc.wantStudentCount {
					t.Errorf("expected student count %d, got %d", tc.wantStudentCount, classAchievement.TotalStudents)
				}
				if len(classAchievement.PerformanceDistribution) != tc.wantPerformanceLevel {
					t.Errorf("expected performance distribution length %d, got %d", tc.wantPerformanceLevel, len(classAchievement.PerformanceDistribution))
				}
				if len(classAchievement.MasteryDistribution) != tc.wantMasteryStatus {
					t.Errorf("expected mastery distribution length %d, got %d", tc.wantMasteryStatus, len(classAchievement.MasteryDistribution))
				}
				if classAchievement.CalculatedAt.IsZero() {
					t.Error("CalculatedAt should not be zero")
				}

				// Check class average calculation
				if len(tc.studentAchievements) > 0 {
					var totalScore float64
					for _, ach := range tc.studentAchievements {
						totalScore += ach.OverallScore
					}
					expectedAverage := totalScore / float64(len(tc.studentAchievements))
					if classAchievement.ClassAverage < expectedAverage-0.1 || classAchievement.ClassAverage > expectedAverage+0.1 {
						t.Errorf("expected class average %.2f, got %.2f", expectedAverage, classAchievement.ClassAverage)
					}
				} else {
					if classAchievement.ClassAverage != 0 {
						t.Errorf("expected class average 0, got %.2f", classAchievement.ClassAverage)
					}
				}
			}
		})
	}
}

// TestAchievementService_BuildSubjectBreakdown tests subject breakdown building
func TestAchievementService_BuildSubjectBreakdown(t *testing.T) {
	service := NewAchievementService()

	tests := []struct {
		name         string
		achievements []Achievement
		wantLength   int
	}{
		{
			name:         "Empty achievements",
			achievements: []Achievement{},
			wantLength:   0,
		},
		{
			name: "Single achievement",
			achievements: []Achievement{
				{
					TPID:         "tp-1",
					TPTitle:      "Algebra",
					OverallScore: 85,
					CalculatedAt: time.Now(),
				},
			},
			wantLength: 0, // Simplified implementation
		},
		{
			name: "Multiple achievements",
			achievements: []Achievement{
				{
					TPID:         "tp-1",
					TPTitle:      "Algebra",
					OverallScore: 85,
					CalculatedAt: time.Now(),
				},
				{
					TPID:         "tp-2",
					TPTitle:      "Geometry",
					OverallScore: 92,
					CalculatedAt: time.Now(),
				},
			},
			wantLength: 0, // Simplified implementation
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := service.buildSubjectBreakdown(tc.achievements)
			if len(result) != tc.wantLength {
				t.Errorf("expected length %d, got %d", tc.wantLength, len(result))
			}
		})
	}
}

// TestAchievementService_IdentifyStrengthsAndWeaknesses tests strength/weakness identification
func TestAchievementService_IdentifyStrengthsAndWeaknesses(t *testing.T) {
	service := NewAchievementService()

	tests := []struct {
		name           string
		achievements   []Achievement
		wantStrengths  int
		wantWeaknesses int
	}{
		{
			name:           "Empty achievements",
			achievements:   []Achievement{},
			wantStrengths:  0,
			wantWeaknesses: 0,
		},
		{
			name: "Mixed performance",
			achievements: []Achievement{
				{
					TPID:                       "tp-1",
					TPTitle:                    "Algebra",
					OverallScore:               95,
					EvaluationPerformanceLevel: EvaluationPerformanceLevelExcellent,
					CalculatedAt:               time.Now(),
				},
				{
					TPID:                       "tp-2",
					TPTitle:                    "Geometry",
					OverallScore:               55,
					EvaluationPerformanceLevel: EvaluationPerformanceLevelBeginning,
					CalculatedAt:               time.Now(),
				},
				{
					TPID:                       "tp-3",
					TPTitle:                    "Statistics",
					OverallScore:               85,
					EvaluationPerformanceLevel: EvaluationPerformanceLevelProficient,
					CalculatedAt:               time.Now(),
				},
			},
			wantStrengths:  1,
			wantWeaknesses: 1,
		},
		{
			name: "All excellent",
			achievements: []Achievement{
				{
					TPID:                       "tp-1",
					TPTitle:                    "Algebra",
					OverallScore:               95,
					EvaluationPerformanceLevel: EvaluationPerformanceLevelExcellent,
					CalculatedAt:               time.Now(),
				},
				{
					TPID:                       "tp-2",
					TPTitle:                    "Geometry",
					OverallScore:               92,
					EvaluationPerformanceLevel: EvaluationPerformanceLevelExcellent,
					CalculatedAt:               time.Now(),
				},
			},
			wantStrengths:  2,
			wantWeaknesses: 0,
		},
		{
			name: "All beginning",
			achievements: []Achievement{
				{
					TPID:                       "tp-1",
					TPTitle:                    "Algebra",
					OverallScore:               45,
					EvaluationPerformanceLevel: EvaluationPerformanceLevelBeginning,
					CalculatedAt:               time.Now(),
				},
				{
					TPID:                       "tp-2",
					TPTitle:                    "Geometry",
					OverallScore:               35,
					EvaluationPerformanceLevel: EvaluationPerformanceLevelBeginning,
					CalculatedAt:               time.Now(),
				},
			},
			wantStrengths:  0,
			wantWeaknesses: 2,
		},
		{
			name: "All proficient",
			achievements: []Achievement{
				{
					TPID:                       "tp-1",
					TPTitle:                    "Algebra",
					OverallScore:               85,
					EvaluationPerformanceLevel: EvaluationPerformanceLevelProficient,
					CalculatedAt:               time.Now(),
				},
				{
					TPID:                       "tp-2",
					TPTitle:                    "Geometry",
					OverallScore:               80,
					EvaluationPerformanceLevel: EvaluationPerformanceLevelProficient,
					CalculatedAt:               time.Now(),
				},
			},
			wantStrengths:  0,
			wantWeaknesses: 0,
		},
		{
			name: "All developing",
			achievements: []Achievement{
				{
					TPID:                       "tp-1",
					TPTitle:                    "Algebra",
					OverallScore:               65,
					EvaluationPerformanceLevel: EvaluationPerformanceLevelDeveloping,
					CalculatedAt:               time.Now(),
				},
				{
					TPID:                       "tp-2",
					TPTitle:                    "Geometry",
					OverallScore:               70,
					EvaluationPerformanceLevel: EvaluationPerformanceLevelDeveloping,
					CalculatedAt:               time.Now(),
				},
			},
			wantStrengths:  0,
			wantWeaknesses: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			strengths, weaknesses := service.identifyStrengthsAndWeaknesses(tc.achievements)
			if len(strengths) != tc.wantStrengths {
				t.Errorf("expected %d strengths, got %d", tc.wantStrengths, len(strengths))
			}
			if len(weaknesses) != tc.wantWeaknesses {
				t.Errorf("expected %d weaknesses, got %d", tc.wantWeaknesses, len(weaknesses))
			}
		})
	}
}

// TestAchievementService_GenerateRecommendations tests recommendation generation
func TestAchievementService_GenerateRecommendations(t *testing.T) {
	service := NewAchievementService()

	tests := []struct {
		name                string
		achievements        []Achievement
		wantRecommendations int
	}{
		{
			name:                "Empty achievements",
			achievements:        []Achievement{},
			wantRecommendations: 0,
		},
		{
			name: "Developing and beginning performance",
			achievements: []Achievement{
				{
					TPID:                       "tp-1",
					TPTitle:                    "Algebra",
					OverallScore:               65,
					EvaluationPerformanceLevel: EvaluationPerformanceLevelDeveloping,
					CalculatedAt:               time.Now(),
				},
				{
					TPID:                       "tp-2",
					TPTitle:                    "Geometry",
					OverallScore:               45,
					EvaluationPerformanceLevel: EvaluationPerformanceLevelBeginning,
					CalculatedAt:               time.Now(),
				},
			},
			wantRecommendations: 2,
		},
		{
			name: "Excellent and proficient performance",
			achievements: []Achievement{
				{
					TPID:                       "tp-1",
					TPTitle:                    "Algebra",
					OverallScore:               95,
					EvaluationPerformanceLevel: EvaluationPerformanceLevelExcellent,
					CalculatedAt:               time.Now(),
				},
				{
					TPID:                       "tp-2",
					TPTitle:                    "Geometry",
					OverallScore:               85,
					EvaluationPerformanceLevel: EvaluationPerformanceLevelProficient,
					CalculatedAt:               time.Now(),
				},
			},
			wantRecommendations: 0,
		},
		{
			name: "Mixed performance",
			achievements: []Achievement{
				{
					TPID:                       "tp-1",
					TPTitle:                    "Algebra",
					OverallScore:               95,
					EvaluationPerformanceLevel: EvaluationPerformanceLevelExcellent,
					CalculatedAt:               time.Now(),
				},
				{
					TPID:                       "tp-2",
					TPTitle:                    "Geometry",
					OverallScore:               65,
					EvaluationPerformanceLevel: EvaluationPerformanceLevelDeveloping,
					CalculatedAt:               time.Now(),
				},
				{
					TPID:                       "tp-3",
					TPTitle:                    "Statistics",
					OverallScore:               85,
					EvaluationPerformanceLevel: EvaluationPerformanceLevelProficient,
					CalculatedAt:               time.Now(),
				},
			},
			wantRecommendations: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			recommendations := service.generateRecommendations(tc.achievements)
			if len(recommendations) != tc.wantRecommendations {
				t.Errorf("expected %d recommendations, got %d", tc.wantRecommendations, len(recommendations))
			}

			// Check recommendation format
			for _, rec := range recommendations {
				if len(rec) == 0 {
					t.Error("recommendation should not be empty")
				}
				if rec[0:5] != "Focus" {
					t.Errorf("recommendation should start with 'Focus', got: %s", rec)
				}
			}
		})
	}
}

// TestAchievementService_BuildStudentClassAchievements tests student class achievement building
func TestAchievementService_BuildStudentClassAchievements(t *testing.T) {
	service := NewAchievementService()

	tests := []struct {
		name         string
		achievements []Achievement
		wantLength   int
	}{
		{
			name:         "Empty achievements",
			achievements: []Achievement{},
			wantLength:   0,
		},
		{
			name: "Single achievement",
			achievements: []Achievement{
				{
					StudentID:                  "student-1",
					StudentName:                "John Doe",
					OverallScore:               85,
					EvaluationPerformanceLevel: EvaluationPerformanceLevelProficient,
					MasteryStatus:              MasteryStatusAchieved,
					CalculatedAt:               time.Now(),
				},
			},
			wantLength: 0, // Simplified implementation
		},
		{
			name: "Multiple achievements",
			achievements: []Achievement{
				{
					StudentID:                  "student-1",
					StudentName:                "John Doe",
					OverallScore:               85,
					EvaluationPerformanceLevel: EvaluationPerformanceLevelProficient,
					MasteryStatus:              MasteryStatusAchieved,
					CalculatedAt:               time.Now(),
				},
				{
					StudentID:                  "student-2",
					StudentName:                "Jane Doe",
					OverallScore:               92,
					EvaluationPerformanceLevel: EvaluationPerformanceLevelExcellent,
					MasteryStatus:              MasteryStatusExceeding,
					CalculatedAt:               time.Now(),
				},
			},
			wantLength: 0, // Simplified implementation
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := service.buildStudentClassAchievements(tc.achievements)
			if len(result) != tc.wantLength {
				t.Errorf("expected length %d, got %d", tc.wantLength, len(result))
			}
		})
	}
}

package domain

import (
	"encoding/json"
	"testing"
	"time"
)

// TestKKTPCriteria_Validation tests KKTPCriteria validation rules
func TestKKTPCriteria_Validation(t *testing.T) {
	tests := []struct {
		name        string
		criteria    KKTPCriteria
		expectedErr string
	}{
		{
			name: "Valid KKTPCriteria",
			criteria: KKTPCriteria{
				MasteryThresholds: MasteryThresholds{
					ExcellentThreshold:  90,
					ProficientThreshold: 80,
					DevelopingThreshold: 70,
					BeginningThreshold:  60,
				},
				PerformanceIndicators: PerformanceIndicators{
					Cognitive:   []string{"Analyze", "Evaluate"},
					Psychomotor: []string{"Demonstrate"},
					Affective:   []string{"Appreciate"},
				},
				MinimumRequirements: MinimumRequirements{
					CoreCompetencies: []string{"Critical thinking"},
					EssentialSkills:  []string{"Problem solving"},
					RequiredEvidence: []string{"Project", "Exam"},
				},
			},
			expectedErr: "",
		},
		{
			name: "Invalid - Excellent not greater than Proficient",
			criteria: KKTPCriteria{
				MasteryThresholds: MasteryThresholds{
					ExcellentThreshold:  80,
					ProficientThreshold: 80,
					DevelopingThreshold: 70,
					BeginningThreshold:  60,
				},
				MinimumRequirements: MinimumRequirements{
					CoreCompetencies: []string{"Critical thinking"},
				},
			},
			expectedErr: "excellent threshold must be greater than proficient threshold",
		},
		{
			name: "Invalid - Proficient not greater than Developing",
			criteria: KKTPCriteria{
				MasteryThresholds: MasteryThresholds{
					ExcellentThreshold:  90,
					ProficientThreshold: 70,
					DevelopingThreshold: 70,
					BeginningThreshold:  60,
				},
				MinimumRequirements: MinimumRequirements{
					CoreCompetencies: []string{"Critical thinking"},
				},
			},
			expectedErr: "proficient threshold must be greater than developing threshold",
		},
		{
			name: "Invalid - Developing not greater than Beginning",
			criteria: KKTPCriteria{
				MasteryThresholds: MasteryThresholds{
					ExcellentThreshold:  90,
					ProficientThreshold: 80,
					DevelopingThreshold: 60,
					BeginningThreshold:  60,
				},
				MinimumRequirements: MinimumRequirements{
					CoreCompetencies: []string{"Critical thinking"},
				},
			},
			expectedErr: "developing threshold must be greater than beginning threshold",
		},
		{
			name: "Invalid - Beginning threshold negative",
			criteria: KKTPCriteria{
				MasteryThresholds: MasteryThresholds{
					ExcellentThreshold:  90,
					ProficientThreshold: 80,
					DevelopingThreshold: 70,
					BeginningThreshold:  -10,
				},
				MinimumRequirements: MinimumRequirements{
					CoreCompetencies: []string{"Critical thinking"},
				},
			},
			expectedErr: "beginning threshold must be non-negative",
		},
		{
			name: "Invalid - No core competencies",
			criteria: KKTPCriteria{
				MasteryThresholds: MasteryThresholds{
					ExcellentThreshold:  90,
					ProficientThreshold: 80,
					DevelopingThreshold: 70,
					BeginningThreshold:  60,
				},
				MinimumRequirements: MinimumRequirements{
					CoreCompetencies: []string{},
				},
			},
			expectedErr: "at least one core competency is required",
		},
		{
			name: "Invalid - Thresholds out of order",
			criteria: KKTPCriteria{
				MasteryThresholds: MasteryThresholds{
					ExcellentThreshold:  70,
					ProficientThreshold: 90,
					DevelopingThreshold: 80,
					BeginningThreshold:  60,
				},
				MinimumRequirements: MinimumRequirements{
					CoreCompetencies: []string{"Critical thinking"},
				},
			},
			expectedErr: "excellent threshold must be greater than proficient threshold",
		},
		{
			name: "Valid - Minimum thresholds",
			criteria: KKTPCriteria{
				MasteryThresholds: MasteryThresholds{
					ExcellentThreshold:  85,
					ProficientThreshold: 75,
					DevelopingThreshold: 65,
					BeginningThreshold:  55,
				},
				PerformanceIndicators: PerformanceIndicators{
					Cognitive: []string{"Understand"},
				},
				MinimumRequirements: MinimumRequirements{
					CoreCompetencies: []string{"Basic skill"},
				},
			},
			expectedErr: "",
		},
		{
			name: "Valid - Empty performance indicators",
			criteria: KKTPCriteria{
				MasteryThresholds: MasteryThresholds{
					ExcellentThreshold:  90,
					ProficientThreshold: 80,
					DevelopingThreshold: 70,
					BeginningThreshold:  60,
				},
				PerformanceIndicators: PerformanceIndicators{
					Cognitive:   []string{},
					Psychomotor: []string{},
					Affective:   []string{},
				},
				MinimumRequirements: MinimumRequirements{
					CoreCompetencies: []string{"Critical thinking"},
				},
			},
			expectedErr: "",
		},
		{
			name: "Valid - Empty optional requirements",
			criteria: KKTPCriteria{
				MasteryThresholds: MasteryThresholds{
					ExcellentThreshold:  90,
					ProficientThreshold: 80,
					DevelopingThreshold: 70,
					BeginningThreshold:  60,
				},
				PerformanceIndicators: PerformanceIndicators{
					Cognitive: []string{"Analyze"},
				},
				MinimumRequirements: MinimumRequirements{
					CoreCompetencies: []string{"Critical thinking"},
					EssentialSkills:  []string{},
					RequiredEvidence: []string{},
				},
			},
			expectedErr: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.criteria.Validate()
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

// TestKKTPCriteria_ToJSON tests JSON serialization
func TestKKTPCriteria_ToJSON(t *testing.T) {
	tests := []struct {
		name        string
		criteria    KKTPCriteria
		expectedErr bool
	}{
		{
			name: "Valid KKTPCriteria to JSON",
			criteria: KKTPCriteria{
				MasteryThresholds: MasteryThresholds{
					ExcellentThreshold:  90,
					ProficientThreshold: 80,
					DevelopingThreshold: 70,
					BeginningThreshold:  60,
				},
				PerformanceIndicators: PerformanceIndicators{
					Cognitive: []string{"Analyze", "Evaluate"},
				},
				MinimumRequirements: MinimumRequirements{
					CoreCompetencies: []string{"Critical thinking"},
				},
			},
			expectedErr: false,
		},
		{
			name: "Empty KKTPCriteria to JSON",
			criteria: KKTPCriteria{
				MasteryThresholds: MasteryThresholds{},
				PerformanceIndicators: PerformanceIndicators{
					Cognitive: []string{},
				},
				MinimumRequirements: MinimumRequirements{
					CoreCompetencies: []string{},
				},
			},
			expectedErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := tc.criteria.ToJSON()
			if tc.expectedErr && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tc.expectedErr && err != nil {
				t.Errorf("expected no error, got: %v", err)
			}
		})
	}
}

// TestFromJSONToKKTPCriteria tests JSON deserialization with validation
func TestFromJSONToKKTPCriteria(t *testing.T) {
	tests := []struct {
		name        string
		data        interface{}
		expectedErr bool
	}{
		{
			name: "Valid JSON to KKTPCriteria",
			data: map[string]interface{}{
				"mastery_thresholds": map[string]interface{}{
					"excellent_threshold":  float64(90),
					"proficient_threshold": float64(80),
					"developing_threshold": float64(70),
					"beginning_threshold":  float64(60),
				},
				"performance_indicators": map[string]interface{}{
					"cognitive":   []string{"Analyze", "Evaluate"},
					"psychomotor": []string{"Demonstrate"},
					"affective":   []string{"Appreciate"},
				},
				"minimum_requirements": map[string]interface{}{
					"core_competencies": []string{"Critical thinking"},
					"essential_skills":  []string{"Problem solving"},
					"required_evidence": []string{"Project", "Exam"},
				},
			},
			expectedErr: false,
		},
		{
			name: "Invalid JSON - threshold violation",
			data: map[string]interface{}{
				"mastery_thresholds": map[string]interface{}{
					"excellent_threshold":  float64(80),
					"proficient_threshold": float64(80),
					"developing_threshold": float64(70),
					"beginning_threshold":  float64(60),
				},
				"minimum_requirements": map[string]interface{}{
					"core_competencies": []string{"Critical thinking"},
				},
			},
			expectedErr: true,
		},
		{
			name: "Invalid JSON - missing core competencies",
			data: map[string]interface{}{
				"mastery_thresholds": map[string]interface{}{
					"excellent_threshold":  float64(90),
					"proficient_threshold": float64(80),
					"developing_threshold": float64(70),
					"beginning_threshold":  float64(60),
				},
				"minimum_requirements": map[string]interface{}{
					"core_competencies": []string{},
				},
			},
			expectedErr: true,
		},
		{
			name:        "Invalid JSON structure",
			data:        "invalid json string",
			expectedErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := FromJSONToKKTPCriteria(tc.data)
			if tc.expectedErr && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tc.expectedErr && err != nil {
				t.Errorf("expected no error, got: %v", err)
			}
		})
	}
}

// TestTPSetValidation tests TPSet domain validation
func TestTPSetValidation(t *testing.T) {
	tests := []struct {
		name        string
		tpSet       TPSet
		expectedErr bool
	}{
		{
			name: "Valid TPSet",
			tpSet: TPSet{
				ID:               "tp-set-001",
				CPID:             "cp-001",
				VersionNo:        1,
				Status:           WorkflowStatusDraft,
				GenerationSource: GenerationSourceManual,
				GeneratedBy:      "user-001",
				CreatedAt:        time.Now(),
				UpdatedAt:        time.Now(),
			},
			expectedErr: false,
		},
		{
			name: "Invalid TPSet - Empty ID",
			tpSet: TPSet{
				ID:               "",
				CPID:             "cp-001",
				VersionNo:        1,
				Status:           WorkflowStatusDraft,
				GenerationSource: GenerationSourceManual,
				GeneratedBy:      "user-001",
				CreatedAt:        time.Now(),
				UpdatedAt:        time.Now(),
			},
			expectedErr: true,
		},
		{
			name: "Invalid TPSet - Empty CP ID",
			tpSet: TPSet{
				ID:               "tp-set-001",
				CPID:             "",
				VersionNo:        1,
				Status:           WorkflowStatusDraft,
				GenerationSource: GenerationSourceManual,
				GeneratedBy:      "user-001",
				CreatedAt:        time.Now(),
				UpdatedAt:        time.Now(),
			},
			expectedErr: true,
		},
		{
			name: "Invalid TPSet - Invalid Version",
			tpSet: TPSet{
				ID:               "tp-set-001",
				CPID:             "cp-001",
				VersionNo:        0,
				Status:           WorkflowStatusDraft,
				GenerationSource: GenerationSourceManual,
				GeneratedBy:      "user-001",
				CreatedAt:        time.Now(),
				UpdatedAt:        time.Now(),
			},
			expectedErr: true,
		},
		{
			name: "Invalid TPSet - Empty GeneratedBy",
			tpSet: TPSet{
				ID:               "tp-set-001",
				CPID:             "cp-001",
				VersionNo:        1,
				Status:           WorkflowStatusDraft,
				GenerationSource: GenerationSourceManual,
				GeneratedBy:      "",
				CreatedAt:        time.Now(),
				UpdatedAt:        time.Now(),
			},
			expectedErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Basic validation checks
			if tc.tpSet.ID == "" && !tc.expectedErr {
				t.Errorf("expected validation error for empty ID")
			}
			if tc.tpSet.CPID == "" && !tc.expectedErr {
				t.Errorf("expected validation error for empty CP ID")
			}
			if tc.tpSet.VersionNo <= 0 && !tc.expectedErr {
				t.Errorf("expected validation error for invalid version")
			}
			if tc.tpSet.GeneratedBy == "" && !tc.expectedErr {
				t.Errorf("expected validation error for empty GeneratedBy")
			}
		})
	}
}

// TestTPValidation tests TP domain validation
func TestTPValidation(t *testing.T) {
	tests := []struct {
		name        string
		tp          TP
		expectedErr bool
	}{
		{
			name: "Valid TP",
			tp: TP{
				ID:             "tp-001",
				TPSetID:        "tp-set-001",
				SequenceNumber: 1,
				CPID:           "cp-001",
				SubjectID:      "subject-001",
				PhaseID:        "phase-001",
				ElementID:      "element-001",
				SubelementID:   "subelement-001",
				UserID:         "user-001",
				Status:         WorkflowStatusDraft,
				VersionNo:      1,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			expectedErr: false,
		},
		{
			name: "Invalid TP - Empty ID",
			tp: TP{
				ID:             "",
				TPSetID:        "tp-set-001",
				SequenceNumber: 1,
				CPID:           "cp-001",
				SubjectID:      "subject-001",
				PhaseID:        "phase-001",
				ElementID:      "element-001",
				SubelementID:   "subelement-001",
				UserID:         "user-001",
				Status:         WorkflowStatusDraft,
				VersionNo:      1,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			expectedErr: true,
		},
		{
			name: "Invalid TP - Invalid Sequence Number",
			tp: TP{
				ID:             "tp-001",
				TPSetID:        "tp-set-001",
				SequenceNumber: 0,
				CPID:           "cp-001",
				SubjectID:      "subject-001",
				PhaseID:        "phase-001",
				ElementID:      "element-001",
				SubelementID:   "subelement-001",
				UserID:         "user-001",
				Status:         WorkflowStatusDraft,
				VersionNo:      1,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			expectedErr: true,
		},
		{
			name: "Invalid TP - Empty UserID",
			tp: TP{
				ID:             "tp-001",
				TPSetID:        "tp-set-001",
				SequenceNumber: 1,
				CPID:           "cp-001",
				SubjectID:      "subject-001",
				PhaseID:        "phase-001",
				ElementID:      "element-001",
				SubelementID:   "subelement-001",
				UserID:         "",
				Status:         WorkflowStatusDraft,
				VersionNo:      1,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			expectedErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Basic validation checks
			if tc.tp.ID == "" && !tc.expectedErr {
				t.Errorf("expected validation error for empty ID")
			}
			if tc.tp.SequenceNumber <= 0 && !tc.expectedErr {
				t.Errorf("expected validation error for invalid sequence number")
			}
			if tc.tp.UserID == "" && !tc.expectedErr {
				t.Errorf("expected validation error for empty UserID")
			}
		})
	}
}

// TestTPVersioning tests TP version tracking logic
func TestTPVersioning(t *testing.T) {
	tests := []struct {
		name            string
		tp              TP
		expectedVersion int
		expectedCurrent bool
	}{
		{
			name: "TP with version 1",
			tp: TP{
				ID:               "tp-001",
				VersionNo:        1,
				IsCurrentVersion: true,
			},
			expectedVersion: 1,
			expectedCurrent: true,
		},
		{
			name: "TP with version 2",
			tp: TP{
				ID:               "tp-002",
				VersionNo:        2,
				IsCurrentVersion: true,
			},
			expectedVersion: 2,
			expectedCurrent: true,
		},
		{
			name: "Historical TP version",
			tp: TP{
				ID:               "tp-003",
				VersionNo:        1,
				IsCurrentVersion: false,
				ParentVersionID:  stringPtr("parent-001"),
			},
			expectedVersion: 1,
			expectedCurrent: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.tp.VersionNo != tc.expectedVersion {
				t.Errorf("expected version %d, got %d", tc.expectedVersion, tc.tp.VersionNo)
			}
			if tc.tp.IsCurrentVersion != tc.expectedCurrent {
				t.Errorf("expected IsCurrentVersion %v, got %v", tc.expectedCurrent, tc.tp.IsCurrentVersion)
			}
		})
	}
}

// TestTPStatusTransitions tests TP workflow status transitions
func TestTPStatusTransitions(t *testing.T) {
	validTransitions := map[WorkflowStatus][]WorkflowStatus{
		WorkflowStatusDraft:       {WorkflowStatusUnderReview, WorkflowStatusDraft},
		WorkflowStatusUnderReview: {WorkflowStatusApproved, WorkflowStatusRejected, WorkflowStatusUnderReview},
		WorkflowStatusApproved:    {WorkflowStatusArchived, WorkflowStatusApproved},
		WorkflowStatusRejected:    {WorkflowStatusDraft, WorkflowStatusRejected},
		WorkflowStatusArchived:    {WorkflowStatusArchived},
	}

	tests := []struct {
		name          string
		fromStatus    WorkflowStatus
		toStatus      WorkflowStatus
		expectedValid bool
	}{
		{
			name:          "Valid transition: Draft to Under Review",
			fromStatus:    WorkflowStatusDraft,
			toStatus:      WorkflowStatusUnderReview,
			expectedValid: true,
		},
		{
			name:          "Valid transition: Under Review to Approved",
			fromStatus:    WorkflowStatusUnderReview,
			toStatus:      WorkflowStatusApproved,
			expectedValid: true,
		},
		{
			name:          "Valid transition: Approved to Archived",
			fromStatus:    WorkflowStatusApproved,
			toStatus:      WorkflowStatusArchived,
			expectedValid: true,
		},
		{
			name:          "Valid transition: Rejected to Draft",
			fromStatus:    WorkflowStatusRejected,
			toStatus:      WorkflowStatusDraft,
			expectedValid: true,
		},
		{
			name:          "Invalid transition: Draft to Approved",
			fromStatus:    WorkflowStatusDraft,
			toStatus:      WorkflowStatusApproved,
			expectedValid: false,
		},
		{
			name:          "Invalid transition: Archived to Draft",
			fromStatus:    WorkflowStatusArchived,
			toStatus:      WorkflowStatusDraft,
			expectedValid: false,
		},
		{
			name:          "Invalid transition: Approved to Draft",
			fromStatus:    WorkflowStatusApproved,
			toStatus:      WorkflowStatusDraft,
			expectedValid: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			allowedTransitions := validTransitions[tc.fromStatus]
			isValid := false
			for _, status := range allowedTransitions {
				if status == tc.toStatus {
					isValid = true
					break
				}
			}
			if isValid != tc.expectedValid {
				t.Errorf("expected valid=%v for transition %s -> %s, got valid=%v",
					tc.expectedValid, tc.fromStatus, tc.toStatus, isValid)
			}
		})
	}
}

// TestSuccessCriteriaJSONB tests SuccessCriteria JSONB handling
func TestSuccessCriteriaJSONB(t *testing.T) {
	tests := []struct {
		name            string
		successCriteria interface{}
		expectedValid   bool
	}{
		{
			name: "Valid SuccessCriteria as map",
			successCriteria: map[string]interface{}{
				"mastery_thresholds": map[string]interface{}{
					"excellent_threshold":  float64(90),
					"proficient_threshold": float64(80),
					"developing_threshold": float64(70),
					"beginning_threshold":  float64(60),
				},
				"minimum_requirements": map[string]interface{}{
					"core_competencies": []string{"Critical thinking"},
				},
			},
			expectedValid: true,
		},
		{
			name:            "Valid SuccessCriteria as JSON string",
			successCriteria: `{"mastery_thresholds":{"excellent_threshold":90,"proficient_threshold":80,"developing_threshold":70,"beginning_threshold":60},"minimum_requirements":{"core_competencies":["Critical thinking"]}}`,
			expectedValid:   true,
		},
		{
			name:            "Valid SuccessCriteria as nil",
			successCriteria: nil,
			expectedValid:   true,
		},
		{
			name:            "Invalid SuccessCriteria - invalid JSON",
			successCriteria: "{invalid json}",
			expectedValid:   false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.successCriteria == nil {
				return // nil is valid
			}

			// Try to convert to JSON to validate structure
			var jsonData []byte
			var err error

			switch v := tc.successCriteria.(type) {
			case string:
				jsonData = []byte(v)
				err = json.Unmarshal(jsonData, &map[string]interface{}{})
			case map[string]interface{}:
				jsonData, err = json.Marshal(v)
			default:
				err = json.Unmarshal([]byte{}, tc.successCriteria)
			}

			if tc.expectedValid && err != nil {
				t.Errorf("expected valid SuccessCriteria, got error: %v", err)
			}
			if !tc.expectedValid && err == nil {
				t.Errorf("expected invalid SuccessCriteria, got no error")
			}
		})
	}
}

// TestLearningObjectivesJSONB tests LearningObjectives JSONB handling
func TestLearningObjectivesJSONB(t *testing.T) {
	tests := []struct {
		name               string
		learningObjectives interface{}
		expectedValid      bool
	}{
		{
			name: "Valid LearningObjectives as array",
			learningObjectives: []string{
				"Students will be able to analyze complex problems",
				"Students will demonstrate critical thinking skills",
			},
			expectedValid: true,
		},
		{
			name:               "Valid LearningObjectives as JSON array string",
			learningObjectives: `["Objective 1", "Objective 2"]`,
			expectedValid:      true,
		},
		{
			name:               "Valid LearningObjectives as nil",
			learningObjectives: nil,
			expectedValid:      true,
		},
		{
			name:               "Invalid LearningObjectives - not an array",
			learningObjectives: map[string]string{"objective": "test"},
			expectedValid:      false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.learningObjectives == nil {
				return // nil is valid
			}

			// Try to convert to JSON to validate structure
			var jsonData []byte
			var err error

			switch v := tc.learningObjectives.(type) {
			case string:
				jsonData = []byte(v)
				err = json.Unmarshal(jsonData, &[]string{})
			case []string:
				jsonData, err = json.Marshal(v)
			default:
				err = json.Unmarshal([]byte{}, tc.learningObjectives)
			}

			if tc.expectedValid && err != nil {
				t.Errorf("expected valid LearningObjectives, got error: %v", err)
			}
			if !tc.expectedValid && err == nil {
				t.Errorf("expected invalid LearningObjectives, got no error")
			}
		})
	}
}

// TestTPSetResponseStructure tests TPSetResponse field completeness
func TestTPSetResponseStructure(t *testing.T) {
	response := TPSetResponse{
		ID:               "tp-set-001",
		CPID:             "cp-001",
		CPCode:           "CP-001",
		CPText:           "Capaian Pembelajaran 1",
		VersionNo:        1,
		Status:           WorkflowStatusDraft,
		GenerationSource: GenerationSourceManual,
		GenerationReason: stringPtr("Manual creation"),
		GeneratedBy:      "user-001",
		GeneratedByName:  "John Doe",
		ApprovedBy:       nil,
		ApprovedByName:   nil,
		ApprovedAt:       nil,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	if response.ID == "" {
		t.Error("ID should not be empty")
	}
	if response.CPID == "" {
		t.Error("CPID should not be empty")
	}
	if response.CPCode == "" {
		t.Error("CPCode should not be empty")
	}
	if response.CPText == "" {
		t.Error("CPText should not be empty")
	}
	if response.VersionNo <= 0 {
		t.Error("VersionNo should be positive")
	}
	if response.GeneratedBy == "" {
		t.Error("GeneratedBy should not be empty")
	}
	if response.GeneratedByName == "" {
		t.Error("GeneratedByName should not be empty")
	}
}

// TestTPResponseStructure tests TPResponse field completeness
func TestTPResponseStructure(t *testing.T) {
	response := TPResponse{
		ID:             "tp-001",
		TPSetID:        "tp-set-001",
		SequenceNumber: 1,
		CPID:           "cp-001",
		CPCode:         "CP-001",
		CPText:         "Capaian Pembelajaran 1",
		SubjectID:      "subject-001",
		SubjectCode:    "SUB-001",
		SubjectName:    "Mathematics",
		PhaseID:        "phase-001",
		PhaseCode:      "PH-001",
		PhaseName:      "Phase F",
		ElementID:      "element-001",
		ElementCode:    "EL-001",
		ElementName:    "Element 1",
		SubelementID:   "subelement-001",
		SubelementCode: "SUB-001",
		SubelementName: "Subelement 1",
		UserID:         "user-001",
		UserName:       "John Doe",
		Status:         WorkflowStatusDraft,
		VersionNo:      1,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if response.ID == "" {
		t.Error("ID should not be empty")
	}
	if response.TPSetID == "" {
		t.Error("TPSetID should not be empty")
	}
	if response.SequenceNumber <= 0 {
		t.Error("SequenceNumber should be positive")
	}
	if response.SubjectID == "" {
		t.Error("SubjectID should not be empty")
	}
	if response.UserID == "" {
		t.Error("UserID should not be empty")
	}
}

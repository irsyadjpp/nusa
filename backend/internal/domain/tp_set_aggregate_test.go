package domain

import (
	"errors"
	"testing"
)

// Test TPSet Aggregate Invariants

func TestTPSetAggregate_CreationRequiresSchoolID(t *testing.T) {
	// TP-INV-001: School Ownership Required
	tpSet := &TPSet{
		ID:        "test-id",
		CPID:      "cp-123",
		VersionNo: 1,
	}

	_, err := NewTPSetAggregate(tpSet, "")
	if err == nil {
		t.Error("Expected SchoolOwnershipRequiredException, got nil")
	}

	var schoolErr *SchoolOwnershipRequiredException
	if err == nil || !errors.As(err, &schoolErr) {
		t.Errorf("Expected SchoolOwnershipRequiredException, got %v", err)
	}
}

func TestTPSetAggregate_CreationRequiresCPID(t *testing.T) {
	// TP-INV-002: CP Reference Required
	tpSet := &TPSet{
		ID:        "test-id",
		CPID:      "",
		VersionNo: 1,
	}

	_, err := NewTPSetAggregate(tpSet, "school-123")
	if err == nil {
		t.Error("Expected CPReferenceRequiredException, got nil")
	}

	var cpErr *CPReferenceRequiredException
	if err == nil || !errors.As(err, &cpErr) {
		t.Errorf("Expected CPReferenceRequiredException, got %v", err)
	}
}

func TestTPSetAggregate_CreationRequiresInitialVersion(t *testing.T) {
	// TP-INV-003: At least one version required
	tpSet := &TPSet{
		ID:        "test-id",
		CPID:      "cp-123",
		VersionNo: 1,
	}

	aggregate, err := NewTPSetAggregate(tpSet, "school-123")
	if err != nil {
		t.Fatalf("Failed to create aggregate: %v", err)
	}

	// Try to add version with wrong version number
	version := &TPVersion{
		ID:               "version-1",
		TPSetID:          "test-id",
		VersionNo:        2, // Should be 1 for first version
		IsCurrentVersion: true,
		Status:           WorkflowStatusDraft,
	}

	err = aggregate.AddVersion(version)
	if err == nil {
		t.Error("Expected AtLeastOneVersionRequiredException, got nil")
	}

	var versionErr *AtLeastOneVersionRequiredException
	if err == nil || !errors.As(err, &versionErr) {
		t.Errorf("Expected AtLeastOneVersionRequiredException, got %v", err)
	}
}

func TestTPSetAggregate_OnlyOneCurrentVersion(t *testing.T) {
	// TP-INV-004: Only one current version
	tpSet := &TPSet{
		ID:        "test-id",
		CPID:      "cp-123",
		VersionNo: 1,
	}

	aggregate, err := NewTPSetAggregate(tpSet, "school-123")
	if err != nil {
		t.Fatalf("Failed to create aggregate: %v", err)
	}

	// Add first current version
	version1 := &TPVersion{
		ID:               "version-1",
		TPSetID:          "test-id",
		VersionNo:        1,
		IsCurrentVersion: true,
		Status:           WorkflowStatusDraft,
	}

	err = aggregate.AddVersion(version1)
	if err != nil {
		t.Fatalf("Failed to add version: %v", err)
	}

	// Try to add second current version
	version2 := &TPVersion{
		ID:               "version-2",
		TPSetID:          "test-id",
		VersionNo:        2,
		IsCurrentVersion: true, // This should fail
		Status:           WorkflowStatusDraft,
	}

	err = aggregate.AddVersion(version2)
	if err == nil {
		t.Error("Expected MultipleCurrentVersionException, got nil")
	}

	var currentErr *MultipleCurrentVersionException
	if err == nil || !errors.As(err, &currentErr) {
		t.Errorf("Expected MultipleCurrentVersionException, got %v", err)
	}
}

func TestTPSetAggregate_VersionNumbersMustBeSequential(t *testing.T) {
	// TP-INV-005: Sequential version numbers
	tpSet := &TPSet{
		ID:        "test-id",
		CPID:      "cp-123",
		VersionNo: 1,
	}

	aggregate, err := NewTPSetAggregate(tpSet, "school-123")
	if err != nil {
		t.Fatalf("Failed to create aggregate: %v", err)
	}

	// Add first version
	version1 := &TPVersion{
		ID:               "version-1",
		TPSetID:          "test-id",
		VersionNo:        1,
		IsCurrentVersion: true,
		Status:           WorkflowStatusDraft,
	}

	err = aggregate.AddVersion(version1)
	if err != nil {
		t.Fatalf("Failed to add version: %v", err)
	}

	// Try to add version with non-sequential number
	version2 := &TPVersion{
		ID:               "version-2",
		TPSetID:          "test-id",
		VersionNo:        3, // Should be 2
		IsCurrentVersion: false,
		Status:           WorkflowStatusDraft,
	}

	err = aggregate.AddVersion(version2)
	if err == nil {
		t.Error("Expected NonSequentialVersionNumberException, got nil")
	}

	var seqErr *NonSequentialVersionNumberException
	if err == nil || !errors.As(err, &seqErr) {
		t.Errorf("Expected NonSequentialVersionNumberException, got %v", err)
	}
}

func TestTPSetAggregate_PublishedVersionImmutability(t *testing.T) {
	// TP-INV-006: Published version immutability
	tpSet := &TPSet{
		ID:        "test-id",
		CPID:      "cp-123",
		VersionNo: 1,
		Status:    WorkflowStatusApproved,
	}

	aggregate, err := NewTPSetAggregate(tpSet, "school-123")
	if err != nil {
		t.Fatalf("Failed to create aggregate: %v", err)
	}

	// Add published version
	version := &TPVersion{
		ID:               "version-1",
		TPSetID:          "test-id",
		VersionNo:        1,
		IsCurrentVersion: true,
		Status:           WorkflowStatusApproved,
	}

	err = aggregate.AddVersion(version)
	if err != nil {
		t.Fatalf("Failed to add version: %v", err)
	}

	// Try to modify published version
	err = aggregate.ModifyCurrentVersion(func(v *TPVersion) error {
		v.Content = "modified content"
		return nil
	})

	if err == nil {
		t.Error("Expected ImmutableVersionException, got nil")
	}

	var immErr *ImmutableVersionException
	if err == nil || !errors.As(err, &immErr) {
		t.Errorf("Expected ImmutableVersionException, got %v", err)
	}
}

func TestTPSetAggregate_StatusTransitions(t *testing.T) {
	// TP-INV-007: Valid workflow status
	tests := []struct {
		name        string
		fromStatus  WorkflowStatus
		toStatus    WorkflowStatus
		shouldError bool
	}{
		{
			name:        "Draft to UnderReview allowed",
			fromStatus:  WorkflowStatusDraft,
			toStatus:    WorkflowStatusUnderReview,
			shouldError: false,
		},
		{
			name:        "UnderReview to Approved allowed",
			fromStatus:  WorkflowStatusUnderReview,
			toStatus:    WorkflowStatusApproved,
			shouldError: false,
		},
		{
			name:        "UnderReview to Draft allowed",
			fromStatus:  WorkflowStatusUnderReview,
			toStatus:    WorkflowStatusDraft,
			shouldError: false,
		},
		{
			name:        "Draft to Approved forbidden",
			fromStatus:  WorkflowStatusDraft,
			toStatus:    WorkflowStatusApproved,
			shouldError: true,
		},
		{
			name:        "Approved to Draft forbidden",
			fromStatus:  WorkflowStatusApproved,
			toStatus:    WorkflowStatusDraft,
			shouldError: true,
		},
		{
			name:        "Approved to UnderReview forbidden",
			fromStatus:  WorkflowStatusApproved,
			toStatus:    WorkflowStatusUnderReview,
			shouldError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tpSet := &TPSet{
				ID:        "test-id",
				CPID:      "cp-123",
				VersionNo: 1,
				Status:    tt.fromStatus,
			}

			aggregate, err := NewTPSetAggregate(tpSet, "school-123")
			if err != nil {
				t.Fatalf("Failed to create aggregate: %v", err)
			}

			err = aggregate.TransitionStatus(tt.toStatus)
			if tt.shouldError && err == nil {
				t.Error("Expected error, got nil")
			}
			if !tt.shouldError && err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
		})
	}
}

func TestTPSetAggregate_ActivateVersion(t *testing.T) {
	tpSet := &TPSet{
		ID:        "test-id",
		CPID:      "cp-123",
		VersionNo: 1,
	}

	aggregate, err := NewTPSetAggregate(tpSet, "school-123")
	if err != nil {
		t.Fatalf("Failed to create aggregate: %v", err)
	}

	// Add two versions
	version1 := &TPVersion{
		ID:               "version-1",
		TPSetID:          "test-id",
		VersionNo:        1,
		IsCurrentVersion: true,
		Status:           WorkflowStatusDraft,
	}

	version2 := &TPVersion{
		ID:               "version-2",
		TPSetID:          "test-id",
		VersionNo:        2,
		IsCurrentVersion: false,
		Status:           WorkflowStatusDraft,
	}

	err = aggregate.AddVersion(version1)
	if err != nil {
		t.Fatalf("Failed to add version1: %v", err)
	}

	err = aggregate.AddVersion(version2)
	if err != nil {
		t.Fatalf("Failed to add version2: %v", err)
	}

	// Activate version 2
	err = aggregate.ActivateVersion(2)
	if err != nil {
		t.Fatalf("Failed to activate version: %v", err)
	}

	// Verify version 1 is no longer current
	if version1.IsCurrentVersion {
		t.Error("Version 1 should not be current after activating version 2")
	}

	// Verify version 2 is current
	if !version2.IsCurrentVersion {
		t.Error("Version 2 should be current after activation")
	}
}

func TestTPSetAggregate_GetCurrentVersion(t *testing.T) {
	tpSet := &TPSet{
		ID:        "test-id",
		CPID:      "cp-123",
		VersionNo: 1,
	}

	aggregate, err := NewTPSetAggregate(tpSet, "school-123")
	if err != nil {
		t.Fatalf("Failed to create aggregate: %v", err)
	}

	// Try to get current version without any versions
	_, err = aggregate.GetCurrentVersion()
	if err == nil {
		t.Error("Expected AtLeastOneVersionRequiredException, got nil")
	}

	// Add a version
	version := &TPVersion{
		ID:               "version-1",
		TPSetID:          "test-id",
		VersionNo:        1,
		IsCurrentVersion: true,
		Status:           WorkflowStatusDraft,
	}

	err = aggregate.AddVersion(version)
	if err != nil {
		t.Fatalf("Failed to add version: %v", err)
	}

	// Get current version
	current, err := aggregate.GetCurrentVersion()
	if err != nil {
		t.Fatalf("Failed to get current version: %v", err)
	}

	if current.ID != "version-1" {
		t.Errorf("Expected version-1, got %s", current.ID)
	}
}

// Test Value Objects

func TestCPCode_Creation(t *testing.T) {
	_, err := NewCPCode("")
	if err == nil {
		t.Error("Expected error for empty CP code")
	}

	code, err := NewCPCode("CP-123")
	if err != nil {
		t.Errorf("Failed to create CP code: %v", err)
	}
	if code.String() != "CP-123" {
		t.Errorf("Expected CP-123, got %s", code.String())
	}
}

func TestCPText_Creation(t *testing.T) {
	_, err := NewCPText("")
	if err == nil {
		t.Error("Expected error for empty CP text")
	}

	text, err := NewCPText("Sample CP text")
	if err != nil {
		t.Errorf("Failed to create CP text: %v", err)
	}
	if text.String() != "Sample CP text" {
		t.Errorf("Expected 'Sample CP text', got %s", text.String())
	}
}

func TestLearningObjective_Creation(t *testing.T) {
	_, err := NewLearningObjective("")
	if err == nil {
		t.Error("Expected error for empty learning objective")
	}

	obj, err := NewLearningObjective("Students will understand...")
	if err != nil {
		t.Errorf("Failed to create learning objective: %v", err)
	}
	if obj.String() != "Students will understand..." {
		t.Errorf("Expected 'Students will understand...', got %s", obj.String())
	}
}

func TestTimeAllocation_Creation(t *testing.T) {
	_, err := NewTimeAllocation(-1, 0, 0)
	if err == nil {
		t.Error("Expected error for negative weeks")
	}

	_, err = NewTimeAllocation(0, -1, 0)
	if err == nil {
		t.Error("Expected error for negative hours")
	}

	_, err = NewTimeAllocation(0, 0, -1)
	if err == nil {
		t.Error("Expected error for negative minutes")
	}

	_, err = NewTimeAllocation(0, 0, 60)
	if err == nil {
		t.Error("Expected error for minutes >= 60")
	}

	alloc, err := NewTimeAllocation(2, 3, 30)
	if err != nil {
		t.Errorf("Failed to create time allocation: %v", err)
	}
	if alloc.GetWeeks() != 2 {
		t.Errorf("Expected 2 weeks, got %d", alloc.GetWeeks())
	}
	if alloc.GetHours() != 3 {
		t.Errorf("Expected 3 hours, got %d", alloc.GetHours())
	}
	if alloc.GetMinutes() != 30 {
		t.Errorf("Expected 30 minutes, got %d", alloc.GetMinutes())
	}
}

func TestSuccessCriteria_Creation(t *testing.T) {
	_, err := NewSuccessCriteria(nil)
	if err == nil {
		t.Error("Expected error for nil criteria")
	}

	criteria, err := NewSuccessCriteria(map[string]interface{}{"key": "value"})
	if err != nil {
		t.Errorf("Failed to create success criteria: %v", err)
	}
	if criteria.GetCriteria() == nil {
		t.Error("Expected non-nil criteria")
	}
}

func TestGenerationReason_Creation(t *testing.T) {
	_, err := NewGenerationReason("")
	if err == nil {
		t.Error("Expected error for empty reason")
	}

	longReason := ""
	for i := 0; i < 501; i++ {
		longReason += "a"
	}
	_, err = NewGenerationReason(longReason)
	if err == nil {
		t.Error("Expected error for reason exceeding 500 characters")
	}

	reason, err := NewGenerationReason("Valid generation reason")
	if err != nil {
		t.Errorf("Failed to create generation reason: %v", err)
	}
	if reason.String() != "Valid generation reason" {
		t.Errorf("Expected 'Valid generation reason', got %s", reason.String())
	}
}

func TestTPSetID_Creation(t *testing.T) {
	_, err := NewTPSetID("")
	if err == nil {
		t.Error("Expected error for empty TPSet ID")
	}

	id, err := NewTPSetID("tp-set-123")
	if err != nil {
		t.Errorf("Failed to create TPSet ID: %v", err)
	}
	if id.String() != "tp-set-123" {
		t.Errorf("Expected 'tp-set-123', got %s", id.String())
	}
}

func TestSchoolID_Creation(t *testing.T) {
	_, err := NewSchoolID("")
	if err == nil {
		t.Error("Expected error for empty school ID")
	}

	id, err := NewSchoolID("school-123")
	if err != nil {
		t.Errorf("Failed to create school ID: %v", err)
	}
	if id.String() != "school-123" {
		t.Errorf("Expected 'school-123', got %s", id.String())
	}
}

func TestUserID_Creation(t *testing.T) {
	_, err := NewUserID("")
	if err == nil {
		t.Error("Expected error for empty user ID")
	}

	id, err := NewUserID("user-123")
	if err != nil {
		t.Errorf("Failed to create user ID: %v", err)
	}
	if id.String() != "user-123" {
		t.Errorf("Expected 'user-123', got %s", id.String())
	}
}

func TestTPSetAggregate_ModifyCurrentVersion_NoCurrentVersion(t *testing.T) {
	tpSet := &TPSet{
		ID:        "test-id",
		CPID:      "cp-123",
		VersionNo: 1,
	}

	aggregate, err := NewTPSetAggregate(tpSet, "school-123")
	if err != nil {
		t.Fatalf("Failed to create aggregate: %v", err)
	}

	// Try to modify when there's no current version
	err = aggregate.ModifyCurrentVersion(func(v *TPVersion) error {
		v.Content = "modified content"
		return nil
	})

	if err == nil {
		t.Error("Expected error when no current version exists")
	}
}

func TestTPSetAggregate_ModifyCurrentVersion_ModificationError(t *testing.T) {
	tpSet := &TPSet{
		ID:        "test-id",
		CPID:      "cp-123",
		VersionNo: 1,
	}

	aggregate, err := NewTPSetAggregate(tpSet, "school-123")
	if err != nil {
		t.Fatalf("Failed to create aggregate: %v", err)
	}

	// Add current version
	version := &TPVersion{
		ID:               "version-1",
		TPSetID:          "test-id",
		VersionNo:        1,
		IsCurrentVersion: true,
		Status:           WorkflowStatusDraft,
	}

	err = aggregate.AddVersion(version)
	if err != nil {
		t.Fatalf("Failed to add version: %v", err)
	}

	// Try to modify with an error in the modification function
	err = aggregate.ModifyCurrentVersion(func(v *TPVersion) error {
		return errors.New("modification error")
	})

	if err == nil {
		t.Error("Expected error from modification function")
	}
}

func TestTPSetAggregate_AddVersion_WrongVersionNumber(t *testing.T) {
	tpSet := &TPSet{
		ID:        "test-id",
		CPID:      "cp-123",
		VersionNo: 1,
	}

	aggregate, err := NewTPSetAggregate(tpSet, "school-123")
	if err != nil {
		t.Fatalf("Failed to create aggregate: %v", err)
	}

	// Add first version
	version1 := &TPVersion{
		ID:               "version-1",
		TPSetID:          "test-id",
		VersionNo:        1,
		IsCurrentVersion: true,
		Status:           WorkflowStatusDraft,
	}

	err = aggregate.AddVersion(version1)
	if err != nil {
		t.Fatalf("Failed to add version: %v", err)
	}

	// Try to add version with wrong version number (should be 2)
	version2 := &TPVersion{
		ID:               "version-2",
		TPSetID:          "test-id",
		VersionNo:        1, // Wrong - should be 2
		IsCurrentVersion: false,
		Status:           WorkflowStatusDraft,
	}

	err = aggregate.AddVersion(version2)
	if err == nil {
		t.Error("Expected error for wrong version number")
	}
}

func TestTPSetAggregate_Getters(t *testing.T) {
	tpSet := &TPSet{
		ID:        "test-id",
		CPID:      "cp-123",
		VersionNo: 1,
	}

	aggregate, err := NewTPSetAggregate(tpSet, "school-123")
	if err != nil {
		t.Fatalf("Failed to create aggregate: %v", err)
	}

	// Test GetTPSet
	retrievedTPSet := aggregate.GetTPSet()
	if retrievedTPSet.ID != tpSet.ID {
		t.Errorf("Expected TPSet ID %s, got %s", tpSet.ID, retrievedTPSet.ID)
	}

	// Test GetVersions
	versions := aggregate.GetVersions()
	if versions == nil {
		t.Error("Expected versions to be initialized")
	}

	// Test GetSchoolID
	schoolID := aggregate.GetSchoolID()
	if schoolID != "school-123" {
		t.Errorf("Expected school ID school-123, got %s", schoolID)
	}
}

func TestTPSetAggregate_ErrorMethods(t *testing.T) {
	// Test SchoolOwnershipRequiredException
	err1 := &SchoolOwnershipRequiredException{}
	if err1.Error() == "" {
		t.Error("Expected error message, got empty string")
	}

	// Test CPReferenceRequiredException
	err2 := &CPReferenceRequiredException{}
	if err2.Error() == "" {
		t.Error("Expected error message, got empty string")
	}

	// Test AtLeastOneVersionRequiredException
	err3 := &AtLeastOneVersionRequiredException{}
	if err3.Error() == "" {
		t.Error("Expected error message, got empty string")
	}

	// Test MultipleCurrentVersionException
	err4 := &MultipleCurrentVersionException{}
	if err4.Error() == "" {
		t.Error("Expected error message, got empty string")
	}

	// Test NonSequentialVersionNumberException
	err5 := &NonSequentialVersionNumberException{}
	if err5.Error() == "" {
		t.Error("Expected error message, got empty string")
	}

	// Test ImmutableVersionException
	err6 := &ImmutableVersionException{}
	if err6.Error() == "" {
		t.Error("Expected error message, got empty string")
	}

	// Test InvalidWorkflowStatusException
	err7 := &InvalidWorkflowStatusException{
		Status: "invalid",
	}
	if err7.Error() == "" {
		t.Error("Expected error message, got empty string")
	}

	// Test NonSequentialVersionNumberException with fields
	err8 := &NonSequentialVersionNumberException{
		Expected: 2,
		Actual:   3,
	}
	if err8.Error() == "" {
		t.Error("Expected error message, got empty string")
	}

	// Test OwnerModificationException
	err9 := &OwnerModificationException{}
	if err9.Error() == "" {
		t.Error("Expected error message, got empty string")
	}
}

func TestValueObject_Equals(t *testing.T) {
	// Test CPCode Equals
	cpCode1, _ := NewCPCode("CP-123")
	cpCode2, _ := NewCPCode("CP-123")
	cpCode3, _ := NewCPCode("CP-456")

	if !cpCode1.Equals(cpCode2) {
		t.Error("Expected CPCode1 to equal CPCode2")
	}
	if cpCode1.Equals(cpCode3) {
		t.Error("Expected CPCode1 to not equal CPCode3")
	}
	if cpCode1.Equals(nil) {
		t.Error("Expected CPCode1 to not equal nil")
	}
	var nilCPCode *CPCode
	if nilCPCode.Equals(cpCode1) {
		t.Error("Expected nil CPCode to not equal CPCode1")
	}

	// Test CPText Equals
	cpText1, _ := NewCPText("Sample CP text")
	cpText2, _ := NewCPText("Sample CP text")
	cpText3, _ := NewCPText("Different text")

	if !cpText1.Equals(cpText2) {
		t.Error("Expected CPText1 to equal CPText2")
	}
	if cpText1.Equals(cpText3) {
		t.Error("Expected CPText1 to not equal CPText3")
	}

	// Test LearningObjective Equals
	lo1, _ := NewLearningObjective("Students will understand...")
	lo2, _ := NewLearningObjective("Students will understand...")
	lo3, _ := NewLearningObjective("Different objective")

	if !lo1.Equals(lo2) {
		t.Error("Expected LO1 to equal LO2")
	}
	if lo1.Equals(lo3) {
		t.Error("Expected LO1 to not equal LO3")
	}

	// Test TimeAllocation Equals
	ta1, _ := NewTimeAllocation(2, 3, 30)
	ta2, _ := NewTimeAllocation(2, 3, 30)
	ta3, _ := NewTimeAllocation(2, 3, 45)

	if !ta1.Equals(ta2) {
		t.Error("Expected TA1 to equal TA2")
	}
	if ta1.Equals(ta3) {
		t.Error("Expected TA1 to not equal TA3")
	}

	// Test SuccessCriteria Equals
	sc1, _ := NewSuccessCriteria("criteria1")
	sc2, _ := NewSuccessCriteria("criteria1")
	sc3, _ := NewSuccessCriteria("criteria2")

	if !sc1.Equals(sc2) {
		t.Error("Expected SC1 to equal SC2")
	}
	if sc1.Equals(sc3) {
		t.Error("Expected SC1 to not equal SC3")
	}

	// Test GenerationReason Equals
	gr1, _ := NewGenerationReason("Initial generation")
	gr2, _ := NewGenerationReason("Initial generation")
	gr3, _ := NewGenerationReason("Updated generation")

	if !gr1.Equals(gr2) {
		t.Error("Expected GR1 to equal GR2")
	}
	if gr1.Equals(gr3) {
		t.Error("Expected GR1 to not equal GR3")
	}

	// Test TPSetID Equals
	tpSetID1, _ := NewTPSetID("tpset-123")
	tpSetID2, _ := NewTPSetID("tpset-123")
	tpSetID3, _ := NewTPSetID("tpset-456")

	if !tpSetID1.Equals(tpSetID2) {
		t.Error("Expected TPSetID1 to equal TPSetID2")
	}
	if tpSetID1.Equals(tpSetID3) {
		t.Error("Expected TPSetID1 to not equal TPSetID3")
	}

	// Test SchoolID Equals
	schoolID1, _ := NewSchoolID("school-123")
	schoolID2, _ := NewSchoolID("school-123")
	schoolID3, _ := NewSchoolID("school-456")

	if !schoolID1.Equals(schoolID2) {
		t.Error("Expected SchoolID1 to equal SchoolID2")
	}
	if schoolID1.Equals(schoolID3) {
		t.Error("Expected SchoolID1 to not equal SchoolID3")
	}

	// Test UserID Equals
	userID1, _ := NewUserID("user-123")
	userID2, _ := NewUserID("user-123")
	userID3, _ := NewUserID("user-456")

	if !userID1.Equals(userID2) {
		t.Error("Expected UserID1 to equal UserID2")
	}
	if userID1.Equals(userID3) {
		t.Error("Expected UserID1 to not equal UserID3")
	}

	// Test Equals with nil receiver for all value objects
	var nilCPCodeReceiver *CPCode
	if nilCPCodeReceiver.Equals(cpCode1) {
		t.Error("Expected nil CPCode to not equal CPCode1")
	}

	var nilCPTextReceiver *CPText
	if nilCPTextReceiver.Equals(cpText1) {
		t.Error("Expected nil CPText to not equal CPText1")
	}

	var nilLOReceiver *LearningObjective
	if nilLOReceiver.Equals(lo1) {
		t.Error("Expected nil LO to not equal LO1")
	}

	var nilTAReceiver *TimeAllocation
	if nilTAReceiver.Equals(ta1) {
		t.Error("Expected nil TA to not equal TA1")
	}

	var nilSCReceiver *SuccessCriteria
	if nilSCReceiver.Equals(sc1) {
		t.Error("Expected nil SC to not equal SC1")
	}

	var nilGRReceiver *GenerationReason
	if nilGRReceiver.Equals(gr1) {
		t.Error("Expected nil GR to not equal GR1")
	}

	var nilTPSetIDReceiver *TPSetID
	if nilTPSetIDReceiver.Equals(tpSetID1) {
		t.Error("Expected nil TPSetID to not equal TPSetID1")
	}

	var nilSchoolIDReceiver *SchoolID
	if nilSchoolIDReceiver.Equals(schoolID1) {
		t.Error("Expected nil SchoolID to not equal SchoolID1")
	}

	var nilUserIDReceiver *UserID
	if nilUserIDReceiver.Equals(userID1) {
		t.Error("Expected nil UserID to not equal UserID1")
	}
}

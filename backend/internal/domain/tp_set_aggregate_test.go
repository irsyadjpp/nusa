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

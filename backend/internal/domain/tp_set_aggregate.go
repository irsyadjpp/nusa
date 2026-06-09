package domain

import (
	"errors"
	"fmt"
	"time"
)

// TPSetAggregate represents the TPSet aggregate root
// This is the aggregate root that enforces all invariants for TPSet
type TPSetAggregate struct {
	tpSet      *TPSet
	versions   []*TPVersion
	schoolID   string
}

// NewTPSetAggregate creates a new TPSet aggregate
func NewTPSetAggregate(tpSet *TPSet, schoolID string) (*TPSetAggregate, error) {
	if tpSet == nil {
		return nil, errors.New("tp_set cannot be nil")
	}
	if schoolID == "" {
		return nil, &SchoolOwnershipRequiredException{}
	}
	if tpSet.CPID == "" {
		return nil, &CPReferenceRequiredException{}
	}

	return &TPSetAggregate{
		tpSet:      tpSet,
		versions:   []*TPVersion{},
		schoolID:   schoolID,
	}, nil
}

// AddVersion adds a new version to the TPSet aggregate
func (a *TPSetAggregate) AddVersion(version *TPVersion) error {
	if version == nil {
		return errors.New("version cannot be nil")
	}

	// TP-INV-003: At least one version required
	if len(a.versions) == 0 && version.VersionNo != 1 {
		return &AtLeastOneVersionRequiredException{}
	}

	// TP-INV-005: Sequential version numbers
	expectedVersionNo := len(a.versions) + 1
	if version.VersionNo != expectedVersionNo {
		return &NonSequentialVersionNumberException{
			Expected: expectedVersionNo,
			Actual:   version.VersionNo,
		}
	}

	// TP-INV-004: Only one current version
	if version.IsCurrentVersion {
		for _, v := range a.versions {
			if v.IsCurrentVersion {
				return &MultipleCurrentVersionException{}
			}
		}
	}

	a.versions = append(a.versions, version)
	return nil
}

// ActivateVersion activates a specific version
func (a *TPSetAggregate) ActivateVersion(versionNo int) error {
	// TP-INV-004: Only one current version
	for _, v := range a.versions {
		if v.VersionNo == versionNo {
			v.IsCurrentVersion = true
		} else {
			v.IsCurrentVersion = false
		}
	}
	return nil
}

// GetCurrentVersion returns the current version
func (a *TPSetAggregate) GetCurrentVersion() (*TPVersion, error) {
	for _, v := range a.versions {
		if v.IsCurrentVersion {
			return v, nil
		}
	}
	return nil, &AtLeastOneVersionRequiredException{}
}

// ModifyCurrentVersion modifies the current version
func (a *TPSetAggregate) ModifyCurrentVersion(modifier func(*TPVersion) error) error {
	current, err := a.GetCurrentVersion()
	if err != nil {
		return err
	}

	// TP-INV-006: Published version immutability
	if current.Status == WorkflowStatusApproved {
		return &ImmutableVersionException{}
	}

	return modifier(current)
}

// TransitionStatus transitions the TPSet status
func (a *TPSetAggregate) TransitionStatus(newStatus WorkflowStatus) error {
	currentStatus := a.tpSet.Status

	// TP-INV-007: Valid workflow status
	if !isValidWorkflowStatus(newStatus) {
		return &InvalidWorkflowStatusException{Status: string(newStatus)}
	}

	// Validate transitions
	switch currentStatus {
	case WorkflowStatusDraft:
		if newStatus != WorkflowStatusUnderReview {
			return &InvalidWorkflowStatusException{
				Status: fmt.Sprintf("Cannot transition from %s to %s", currentStatus, newStatus),
			}
		}
	case WorkflowStatusUnderReview:
		if newStatus != WorkflowStatusApproved && newStatus != WorkflowStatusDraft {
			return &InvalidWorkflowStatusException{
				Status: fmt.Sprintf("Cannot transition from %s to %s", currentStatus, newStatus),
			}
		}
	case WorkflowStatusApproved:
		// TP-INV-006: Published version immutability - cannot transition from approved
		return &ImmutableVersionException{}
	}

	a.tpSet.Status = newStatus
	return nil
}

// GetTPSet returns the TPSet
func (a *TPSetAggregate) GetTPSet() *TPSet {
	return a.tpSet
}

// GetVersions returns all versions
func (a *TPSetAggregate) GetVersions() []*TPVersion {
	return a.versions
}

// GetSchoolID returns the school ID
func (a *TPSetAggregate) GetSchoolID() string {
	return a.schoolID
}

// TPVersion represents a TPSet version entity
type TPVersion struct {
	ID               string        `json:"id" db:"id"`
	TPSetID          string        `json:"tp_set_id" db:"tp_set_id"`
	VersionNo        int           `json:"version_no" db:"version_no"`
	IsCurrentVersion bool          `json:"is_current_version" db:"is_current_version"`
	Status           WorkflowStatus `json:"status" db:"status"`
	Content          interface{}   `json:"content" db:"content"`
	CreatedAt        time.Time     `json:"created_at" db:"created_at"`
	CreatedBy        string        `json:"created_by" db:"created_by"`
}

// Domain Exceptions for TPSet Aggregate

// SchoolOwnershipRequiredException is thrown when TPSet does not have school ownership
type SchoolOwnershipRequiredException struct{}

func (e *SchoolOwnershipRequiredException) Error() string {
	return "TPSet must belong to exactly one School"
}

// CPReferenceRequiredException is thrown when TPSet does not reference a valid CP
type CPReferenceRequiredException struct{}

func (e *CPReferenceRequiredException) Error() string {
	return "TPSet must reference a valid CP"
}

// AtLeastOneVersionRequiredException is thrown when TPSet does not have at least one version
type AtLeastOneVersionRequiredException struct{}

func (e *AtLeastOneVersionRequiredException) Error() string {
	return "TPSet must have at least one version"
}

// MultipleCurrentVersionException is thrown when attempting to have multiple current versions
type MultipleCurrentVersionException struct{}

func (e *MultipleCurrentVersionException) Error() string {
	return "Only one version can be current for a TPSet"
}

// NonSequentialVersionNumberException is thrown when version numbers are not sequential
type NonSequentialVersionNumberException struct {
	Expected int
	Actual   int
}

func (e *NonSequentialVersionNumberException) Error() string {
	return fmt.Sprintf("Version number must be sequential. Expected: %d, Actual: %d", e.Expected, e.Actual)
}

// ImmutableVersionException is thrown when attempting to modify a published version
type ImmutableVersionException struct{}

func (e *ImmutableVersionException) Error() string {
	return "Published versions are immutable"
}

// InvalidWorkflowStatusException is thrown when status is invalid
type InvalidWorkflowStatusException struct {
	Status string
}

func (e *InvalidWorkflowStatusException) Error() string {
	return fmt.Sprintf("Invalid workflow status: %s", e.Status)
}

// OwnerModificationException is thrown when attempting to modify the owner
type OwnerModificationException struct{}

func (e *OwnerModificationException) Error() string {
	return "TPSet owner cannot be modified after creation"
}

// Helper functions

func isValidWorkflowStatus(status WorkflowStatus) bool {
	switch status {
	case WorkflowStatusDraft, WorkflowStatusUnderReview, WorkflowStatusApproved:
		return true
	default:
		return false
	}
}

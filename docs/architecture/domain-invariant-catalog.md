# NUSA PLATFORM — Domain Invariant Catalog

**Document Version**: 1.0
**Effective Date**: 2026-06-09
**Status**: FINAL SOURCE OF TRUTH FOR DOMAIN INVARIANTS
**Governance**: This document is the single source of truth for all domain invariants before implementation begins

---

# Document Purpose

This catalog defines all domain invariants, business rules, validation rules, lifecycle rules, state transitions, versioning rules, authorization assumptions, invariant violation scenarios, domain exceptions, and required unit tests for all aggregates in the NUSA Platform.

## Why Invariants Exist

Domain invariants are the fundamental business rules that must always hold true for an aggregate. They represent the core business logic that protects data integrity, ensures consistency, and prevents invalid states. Invariants are the contract between the domain and the rest of the system.

## Why Invariants Must Be Enforced in Aggregate Roots

Aggregate roots are the sole guardians of invariants. Only the aggregate root can modify its internal state. This ensures that:
- Invariants cannot be bypassed by external code
- All state transitions go through invariant validation
- The aggregate remains consistent at all times
- No external code can directly modify internal entities

## Why Handlers/Services/Controllers Must Never Bypass Invariants

Handlers, services, and controllers are application layer concerns. They must never:
- Directly modify aggregate internal state
- Skip invariant validation
- Modify entities bypassing the aggregate root
- Create or modify entities outside aggregate boundaries

Bypassing invariants leads to:
- Data corruption
- Invalid states
- Inconsistent business logic
- Security vulnerabilities
- Unpredictable system behavior

---

# Aggregate: TPSet

## Aggregate Purpose

TPSet (Tujuan Pembelajaran Set) represents a set of learning objectives derived from a Capaian Pembelajaran (CP). It is the foundational aggregate in the Learning Planning Context and serves as the basis for ATP Sets, Modul Ajar Sets, and Assessments.

## Aggregate Root

TPSet

## Entities

- TPSet (root)
- TPVersion

## Value Objects

- CPCode
- CPText
- LearningObjective
- TimeAllocation
- SuccessCriteria
- GenerationReason
- WorkflowStatus
- TPSetID
- SchoolID
- UserID

## Invariants

### TP-INV-001: School Ownership Required

**Rule**: A TPSet must belong to exactly one School.

**Violation**: Creation or modification without valid school_id.

**Domain Exception**: `SchoolOwnershipRequiredException`

**Test Requirement**: Must reject TPSet creation without school_id. Must reject TPSet modification that attempts to change school_id.

---

### TP-INV-002: CP Reference Required

**Rule**: A TPSet must reference a valid CP (Capaian Pembelajaran).

**Violation**: Creation or modification without valid cp_id.

**Domain Exception**: `CPReferenceRequiredException`

**Test Requirement**: Must reject TPSet creation without cp_id. Must reject TPSet modification with invalid cp_id.

---

### TP-INV-003: At Least One Version Required

**Rule**: A TPSet must have at least one version.

**Violation**: Creation without initial version. Deletion of all versions.

**Domain Exception**: `AtLeastOneVersionRequiredException`

**Test Requirement**: Must reject TPSet creation without initial version. Must reject deletion of the last version.

---

### TP-INV-004: Only One Current Version

**Rule**: Only one version may have is_current_version = true for a TPSet.

**Violation**: Attempting to activate a version when another current version exists. Attempting to create a version with is_current_version = true without deactivating the current version.

**Domain Exception**: `MultipleCurrentVersionException`

**Test Requirement**: Must reject activation when another current version exists. Must reject creation of current version without deactivating existing current version.

---

### TP-INV-005: Sequential Version Numbers

**Rule**: Version numbers must be sequential integers with no gaps.

**Violation**: Creating a version with non-sequential version number. Creating a version with a version number that skips a number.

**Domain Exception**: `NonSequentialVersionNumberException`

**Test Requirement**: Must reject version creation with non-sequential version number. Must enforce version_no = max(version_no) + 1.

---

### TP-INV-006: Published Version Immutability

**Rule**: Published versions (status = APPROVED) are immutable.

**Violation**: Attempting to modify content fields of a published version. Attempting to delete a published version.

**Domain Exception**: `ImmutableVersionException`

**Test Requirement**: Must reject update of published version content fields. Must reject deletion of published version. Must allow update of mutable fields (status, approved_by, approved_at) even for published versions.

---

### TP-INV-007: Valid Workflow Status

**Rule**: TPSet status must be one of: DRAFT, UNDER_REVIEW, APPROVED.

**Violation**: Setting status to invalid value. Transitioning to invalid status.

**Domain Exception**: `InvalidWorkflowStatusException`

**Test Requirement**: Must reject invalid status values. Must validate all status transitions.

---

### TP-INV-008: Owner Modification Restriction

**Rule**: TPSet owner (generated_by) cannot be modified after creation.

**Violation**: Attempting to change generated_by field.

**Domain Exception**: `OwnerModificationException`

**Test Requirement**: Must reject modification of generated_by field after creation.

---

## Lifecycle Rules

**Allowed Transitions**:
- DRAFT → UNDER_REVIEW
- UNDER_REVIEW → APPROVED
- UNDER_REVIEW → DRAFT (return to draft for revisions)

**Forbidden Transitions**:
- APPROVED → DRAFT (cannot return to draft after approval)
- APPROVED → UNDER_REVIEW (cannot return to review after approval)
- DRAFT → APPROVED (must go through review first)

**State Machine**:
```
DRAFT → UNDER_REVIEW → APPROVED
  ↑         ↓
  └─────────┘
```

## Versioning Rules

**Version Creation Triggers**:
- Content changes (learning_objectives, time_allocation, success_criteria)
- Status transitions to APPROVED
- Regeneration with new parameters
- Explicit version creation requested

**Version Number Generation**:
- Initial version: version_no = 1
- Subsequent versions: version_no = max(version_no) + 1
- No gaps allowed

**Current Version Rules**:
- Only one version has is_current_version = true
- New version becomes current when created
- Previous current version is set to is_current_version = false

**Historical Version Rules**:
- Historical versions are preserved indefinitely
- Historical versions cannot be modified (except mutable fields)
- Historical versions can be queried for audit

## Authorization Assumptions

**Teacher**:
- Can create TPSet for own school
- Can modify own TPSet only
- Can delete own TPSet only
- Cannot modify TPSet owned by other teachers
- Cannot access TPSet from other schools

**School Admin**:
- Can read all TPSet in school
- Can modify any TPSet in school
- Can approve TPSet in school
- Can delete any TPSet in school
- Cannot access TPSet from other schools

**System Admin**:
- Can read all TPSet across all schools
- Can modify any TPSet (emergency override)
- Can approve any TPSet (emergency override)
- Can delete any TPSet (emergency override)
- Read-only access unless emergency override required

## Invariant Violation Scenarios

**Scenario 1**: Teacher attempts to create TPSet without school_id
- Violation: TP-INV-001
- Exception: SchoolOwnershipRequiredException
- Resolution: Require school_id from authenticated user context

**Scenario 2**: Teacher attempts to create TPSet without cp_id
- Violation: TP-INV-002
- Exception: CPReferenceRequiredException
- Resolution: Require cp_id from request

**Scenario 3**: Teacher attempts to activate second current version
- Violation: TP-INV-004
- Exception: MultipleCurrentVersionException
- Resolution: Deactivate current version before activating new version

**Scenario 4**: Teacher attempts to modify published version content
- Violation: TP-INV-006
- Exception: ImmutableVersionException
- Resolution: Create new version for content changes

**Scenario 5**: Teacher attempts to change owner after creation
- Violation: TP-INV-008
- Exception: OwnerModificationException
- Resolution: Reject owner modification

## Domain Exceptions

```typescript
class SchoolOwnershipRequiredException extends DomainException {
  constructor() {
    super("TPSet must belong to exactly one School");
  }
}

class CPReferenceRequiredException extends DomainException {
  constructor() {
    super("TPSet must reference a valid CP");
  }
}

class AtLeastOneVersionRequiredException extends DomainException {
  constructor() {
    super("TPSet must have at least one version");
  }
}

class MultipleCurrentVersionException extends DomainException {
  constructor() {
    super("Only one version can be current for a TPSet");
  }
}

class NonSequentialVersionNumberException extends DomainException {
  constructor(expected: number, actual: number) {
    super(`Version number must be sequential. Expected: ${expected}, Actual: ${actual}`);
  }
}

class ImmutableVersionException extends DomainException {
  constructor() {
    super("Published versions are immutable");
  }
}

class InvalidWorkflowStatusException extends DomainException {
  constructor(status: string) {
    super(`Invalid workflow status: ${status}`);
  }
}

class OwnerModificationException extends DomainException {
  constructor() {
    super("TPSet owner cannot be modified after creation");
  }
}
```

## Required Unit Tests

- `test_TPSet_creation_requires_school_id`
- `test_TPSet_creation_requires_cp_id`
- `test_TPSet_creation_requires_initial_version`
- `test_TPSet_activation_rejects_multiple_current_versions`
- `test_TPSet_version_numbers_must_be_sequential`
- `test_TPSet_published_version_content_is_immutable`
- `test_TPSet_status_must_be_valid`
- `test_TPSet_owner_cannot_be_modified`
- `test_TPSet_draft_to_under_review_transition_allowed`
- `test_TPSet_under_review_to_approved_transition_allowed`
- `test_TPSet_under_review_to_draft_transition_allowed`
- `test_TPSet_approved_to_draft_transition_forbidden`
- `test_TPSet_approved_to_under_review_transition_forbidden`
- `test_TPSet_draft_to_approved_transition_forbidden`

## Aggregate Boundary

**Inside Aggregate**:
- TPSet (root)
- TPVersion (entity)

**Outside Aggregate**:
- CP (external reference, not owned by TPSet)
- School (external reference, not owned by TPSet)
- User (external reference, not owned by TPSet)

## Aggregate Consistency Boundary

TPSet and all its versions must be transactionally consistent. All operations on TPSet and its versions must occur within a single transaction.

## External Dependencies

**May Reference**:
- CP (Capaian Pembelajaran) - read-only reference
- School - read-only reference for ownership
- User - read-only reference for ownership

## Forbidden Dependencies

**Must Not Reference**:
- ATPSet (ATPSet references TPSet, not vice versa)
- ModulAjarSet (ModulAjarSet references TPSet via ATPSet, not vice versa)
- Assessment (Assessment references TPSet, not vice versa)

## Ownership Rules

**Teacher Ownership**:
- Teacher owns TPSet they created
- Teacher can modify own TPSet
- Teacher can delete own TPSet
- Teacher cannot modify TPSet owned by others

**School Ownership**:
- TPSet belongs to teacher's school
- School Admin can access all school's TPSet
- School Admin can modify any school's TPSet
- Cross-school access forbidden

**Admin Ownership**:
- System Admin has emergency override access
- System Admin can modify any TPSet
- System Admin access requires audit logging

## School Isolation Rules

- TPSet queries must include school scope filter
- Cross-school TPSet access must return 404 (not 403)
- School scope derived from user.school_id
- System Admin bypasses school scope filter

---

# Aggregate: ATPSet

## Aggregate Purpose

ATPSet (Alokasi Waktu Tujuan Pembelajaran Set) represents the allocation of time for learning objectives across weeks. It is derived from a TPSet and serves as the basis for Modul Ajar Sets.

## Aggregate Root

ATPSet

## Entities

- ATPSet (root)
- ATPVersion
- ATPItem

## Value Objects

- TPSetID
- TPSetVersionNo
- WeekNumber
- CPCode
- CPText
- GenerationReason
- WorkflowStatus
- ATPSetID
- SchoolID
- UserID

## Invariants

### ATP-INV-001: TPSet Reference Required

**Rule**: ATPSet must reference a valid TPSet.

**Violation**: Creation or modification without valid tp_set_id.

**Domain Exception**: `TPSetReferenceRequiredException`

**Test Requirement**: Must reject ATPSet creation without tp_set_id. Must reject ATPSet modification with invalid tp_set_id.

---

### ATP-INV-002: TPSet Version Reference Required

**Rule**: ATPSet must reference a specific TPSet version.

**Violation**: Creation or modification without valid tp_set_version_no.

**Domain Exception**: `TPSetVersionReferenceRequiredException`

**Test Requirement**: Must reject ATPSet creation without tp_set_version_no. Must validate tp_set_version_no exists for referenced TPSet.

---

### ATP-INV-003: School Ownership Required

**Rule**: ATPSet must belong to exactly one School.

**Violation**: Creation or modification without valid school_id.

**Domain Exception**: `SchoolOwnershipRequiredException`

**Test Requirement**: Must reject ATPSet creation without school_id. Must reject ATPSet modification that attempts to change school_id.

---

### ATP-INV-004: At Least One Version Required

**Rule**: ATPSet must have at least one version.

**Violation**: Creation without initial version. Deletion of all versions.

**Domain Exception**: `AtLeastOneVersionRequiredException`

**Test Requirement**: Must reject ATPSet creation without initial version. Must reject deletion of the last version.

---

### ATP-INV-005: Only One Current Version

**Rule**: Only one version may have is_current_version = true for an ATPSet.

**Violation**: Attempting to activate a version when another current version exists.

**Domain Exception**: `MultipleCurrentVersionException`

**Test Requirement**: Must reject activation when another current version exists.

---

### ATP-INV-006: Sequential Version Numbers

**Rule**: Version numbers must be sequential integers with no gaps.

**Violation**: Creating a version with non-sequential version number.

**Domain Exception**: `NonSequentialVersionNumberException`

**Test Requirement**: Must reject version creation with non-sequential version number.

---

### ATP-INV-007: No Week Overlap

**Rule**: Week allocations within an ATPSet must not overlap.

**Violation**: Creating ATPItems with overlapping week numbers.

**Domain Exception**: `WeekOverlapException`

**Test Requirement**: Must reject ATPItem creation with overlapping week numbers.

---

### ATP-INV-008: Published Version Immutability

**Rule**: Published versions (status = APPROVED) are immutable.

**Violation**: Attempting to modify content fields of a published version.

**Domain Exception**: `ImmutableVersionException`

**Test Requirement**: Must reject update of published version content fields.

---

## Lifecycle Rules

**Allowed Transitions**:
- DRAFT → UNDER_REVIEW
- UNDER_REVIEW → APPROVED
- UNDER_REVIEW → DRAFT

**Forbidden Transitions**:
- APPROVED → DRAFT
- APPROVED → UNDER_REVIEW
- DRAFT → APPROVED

## Versioning Rules

**Version Creation Triggers**:
- TPSet reference changes
- ATP content changes (week allocations)
- Status transitions to APPROVED

**Version Number Generation**:
- Initial version: version_no = 1
- Subsequent versions: version_no = max(version_no) + 1

## Authorization Assumptions

**Teacher**: Can modify own ATPSet only
**School Admin**: Can modify any ATPSet in school
**System Admin**: Emergency override access

## Invariant Violation Scenarios

**Scenario 1**: Teacher attempts to create ATPSet without tp_set_id
- Violation: ATP-INV-001
- Exception: TPSetReferenceRequiredException

**Scenario 2**: Teacher attempts to create ATPItems with overlapping weeks
- Violation: ATP-INV-007
- Exception: WeekOverlapException

## Domain Exceptions

```typescript
class TPSetReferenceRequiredException extends DomainException {
  constructor() {
    super("ATPSet must reference a valid TPSet");
  }
}

class TPSetVersionReferenceRequiredException extends DomainException {
  constructor() {
    super("ATPSet must reference a specific TPSet version");
  }
}

class WeekOverlapException extends DomainException {
  constructor(week1: number, week2: number) {
    super(`Week allocations must not overlap. Conflict: ${week1} and ${week2}`);
  }
}
```

## Required Unit Tests

- `test_ATPSet_creation_requires_tp_set_id`
- `test_ATPSet_creation_requires_tp_set_version_no`
- `test_ATPSet_creation_requires_school_id`
- `test_ATPSet_week_allocations_must_not_overlap`
- `test_ATPSet_version_numbers_must_be_sequential`
- `test_ATPSet_published_version_content_is_immutable`

## Aggregate Boundary

**Inside Aggregate**:
- ATPSet (root)
- ATPVersion (entity)
- ATPItem (entity)

**Outside Aggregate**:
- TPSet (external reference)
- School (external reference)
- User (external reference)

## Aggregate Consistency Boundary

ATPSet, all its versions, and ATPItems must be transactionally consistent.

## External Dependencies

**May Reference**:
- TPSet (read-only reference)
- School (read-only reference)
- User (read-only reference)

## Forbidden Dependencies

**Must Not Reference**:
- ModulAjarSet (ModulAjarSet references ATPSet, not vice versa)

## Ownership Rules

**Teacher Ownership**: Teacher owns ATPSet they generated
**School Ownership**: ATPSet belongs to teacher's school
**Admin Ownership**: System Admin has emergency override

## School Isolation Rules

- ATPSet queries must include school scope filter
- Cross-school ATPSet access must return 404

---

# Aggregate: ModulAjarSet

## Aggregate Purpose

ModulAjarSet represents teaching modules derived from an ATPSet. It contains detailed content for teaching activities.

## Aggregate Root

ModulAjarSet

## Entities

- ModulAjarSet (root)
- ModulAjarVersion
- ModulAjarItem

## Value Objects

- ATPSetID
- ATPSetVersionNo
- SequenceNumber
- Content
- GenerationReason
- WorkflowStatus
- ModulAjarSetID
- SchoolID
- UserID

## Invariants

### MA-INV-001: ATPSet Reference Required

**Rule**: ModulAjarSet must reference a valid ATPSet.

**Violation**: Creation or modification without valid atp_set_id.

**Domain Exception**: `ATPSetReferenceRequiredException`

**Test Requirement**: Must reject ModulAjarSet creation without atp_set_id.

---

### MA-INV-002: ATPSet Version Reference Required

**Rule**: ModulAjarSet must reference a specific ATPSet version.

**Violation**: Creation or modification without valid atp_set_version_no.

**Domain Exception**: `ATPSetVersionReferenceRequiredException`

**Test Requirement**: Must reject ModulAjarSet creation without atp_set_version_no.

---

### MA-INV-003: School Ownership Required

**Rule**: ModulAjarSet must belong to exactly one School.

**Violation**: Creation or modification without valid school_id.

**Domain Exception**: `SchoolOwnershipRequiredException`

**Test Requirement**: Must reject ModulAjarSet creation without school_id.

---

### MA-INV-004: At Least One Version Required

**Rule**: ModulAjarSet must have at least one version.

**Violation**: Creation without initial version.

**Domain Exception**: `AtLeastOneVersionRequiredException`

**Test Requirement**: Must reject ModulAjarSet creation without initial version.

---

### MA-INV-005: Only One Current Version

**Rule**: Only one version may have is_current_version = true.

**Violation**: Attempting to activate a version when another current version exists.

**Domain Exception**: `MultipleCurrentVersionException`

**Test Requirement**: Must reject activation when another current version exists.

---

### MA-INV-006: Sequential Version Numbers

**Rule**: Version numbers must be sequential integers.

**Violation**: Creating a version with non-sequential version number.

**Domain Exception**: `NonSequentialVersionNumberException`

**Test Requirement**: Must reject version creation with non-sequential version number.

---

### MA-INV-007: Sequential Item Numbers

**Rule**: ModulAjarItem sequence numbers must be sequential.

**Violation**: Creating items with non-sequential sequence numbers.

**Domain Exception**: `NonSequentialItemNumberException`

**Test Requirement**: Must reject item creation with non-sequential sequence number.

---

### MA-INV-008: Published Version Immutability

**Rule**: Published versions are immutable.

**Violation**: Attempting to modify content of published version.

**Domain Exception**: `ImmutableVersionException`

**Test Requirement**: Must reject update of published version content.

---

## Lifecycle Rules

**Allowed Transitions**:
- DRAFT → UNDER_REVIEW
- UNDER_REVIEW → APPROVED
- UNDER_REVIEW → DRAFT

**Forbidden Transitions**:
- APPROVED → DRAFT
- APPROVED → UNDER_REVIEW
- DRAFT → APPROVED

## Versioning Rules

**Version Creation Triggers**:
- ATPSet reference changes
- Content changes
- Status transitions to APPROVED

## Authorization Assumptions

**Teacher**: Can modify own ModulAjarSet only
**School Admin**: Can modify any ModulAjarSet in school
**System Admin**: Emergency override access

## Domain Exceptions

```typescript
class ATPSetReferenceRequiredException extends DomainException {
  constructor() {
    super("ModulAjarSet must reference a valid ATPSet");
  }
}

class NonSequentialItemNumberException extends DomainException {
  constructor(expected: number, actual: number) {
    super(`Item numbers must be sequential. Expected: ${expected}, Actual: ${actual}`);
  }
}
```

## Required Unit Tests

- `test_ModulAjarSet_creation_requires_atp_set_id`
- `test_ModulAjarSet_creation_requires_atp_set_version_no`
- `test_ModulAjarSet_creation_requires_school_id`
- `test_ModulAjarSet_item_numbers_must_be_sequential`
- `test_ModulAjarSet_version_numbers_must_be_sequential`
- `test_ModulAjarSet_published_version_content_is_immutable`

## Aggregate Boundary

**Inside Aggregate**:
- ModulAjarSet (root)
- ModulAjarVersion (entity)
- ModulAjarItem (entity)

**Outside Aggregate**:
- ATPSet (external reference)
- School (external reference)
- User (external reference)

## Aggregate Consistency Boundary

ModulAjarSet, all its versions, and ModulAjarItems must be transactionally consistent.

## External Dependencies

**May Reference**:
- ATPSet (read-only reference)
- School (read-only reference)
- User (read-only reference)

## Forbidden Dependencies

**Must Not Reference**:
- None (ModulAjarSet is leaf aggregate in Learning Planning)

## Ownership Rules

**Teacher Ownership**: Teacher owns ModulAjarSet they generated
**School Ownership**: ModulAjarSet belongs to teacher's school
**Admin Ownership**: System Admin has emergency override

## School Isolation Rules

- ModulAjarSet queries must include school scope filter
- Cross-school ModulAjarSet access must return 404

---

# Aggregate: Assessment

## Aggregate Purpose

Assessment represents a test or evaluation instrument for measuring student achievement against learning objectives.

## Aggregate Root

Assessment

## Entities

- Assessment (root)
- AssessmentVersion
- AssessmentItem
- AnswerKey
- ScoringGuideline

## Value Objects

- TPSetID
- TPSetVersionNo
- AssessmentID
- TotalScore
- PerformanceLevel
- WorkflowStatus
- SchoolID
- UserID

## Invariants

### AS-INV-001: TPSet Reference Required

**Rule**: Assessment must reference a valid TPSet.

**Violation**: Creation or modification without valid tp_id.

**Domain Exception**: `TPSetReferenceRequiredException`

**Test Requirement**: Must reject Assessment creation without tp_id.

---

### AS-INV-002: TPSet Version Reference Required

**Rule**: Assessment must reference a specific TPSet version.

**Violation**: Creation or modification without valid tp_version_no.

**Domain Exception**: `TPSetVersionReferenceRequiredException`

**Test Requirement**: Must reject Assessment creation without tp_version_no.

---

### AS-INV-003: School Ownership Required

**Rule**: Assessment must belong to exactly one School.

**Violation**: Creation or modification without valid school_id.

**Domain Exception**: `SchoolOwnershipRequiredException`

**Test Requirement**: Must reject Assessment creation without school_id.

---

### AS-INV-004: At Least One Version Required

**Rule**: Assessment must have at least one version.

**Violation**: Creation without initial version.

**Domain Exception**: `AtLeastOneVersionRequiredException`

**Test Requirement**: Must reject Assessment creation without initial version.

---

### AS-INV-005: Only One Current Version

**Rule**: Only one version may have is_current_version = true.

**Violation**: Attempting to activate a version when another current version exists.

**Domain Exception**: `MultipleCurrentVersionException`

**Test Requirement**: Must reject activation when another current version exists.

---

### AS-INV-006: Sequential Version Numbers

**Rule**: Version numbers must be sequential integers.

**Violation**: Creating a version with non-sequential version number.

**Domain Exception**: `NonSequentialVersionNumberException`

**Test Requirement**: Must reject version creation with non-sequential version number.

---

### AS-INV-007: At Least One Item Required

**Rule**: Assessment must have at least one item.

**Violation**: Creation without items. Deletion of all items.

**Domain Exception**: `AtLeastOneItemRequiredException`

**Test Requirement**: Must reject Assessment creation without items. Must reject deletion of the last item.

---

### AS-INV-008: Sequential Item Numbers

**Rule**: AssessmentItem sequence numbers must be sequential.

**Violation**: Creating items with non-sequential sequence numbers.

**Domain Exception**: `NonSequentialItemNumberException`

**Test Requirement**: Must reject item creation with non-sequential sequence number.

---

### AS-INV-009: Valid Score Range

**Rule**: AssessmentItem max_score must be within valid range (0-100 or higher).

**Violation**: Creating item with invalid max_score.

**Domain Exception**: `InvalidScoreRangeException`

**Test Requirement**: Must reject item creation with max_score < 0.

---

### AS-INV-010: Published Version Immutability

**Rule**: Published versions are immutable.

**Violation**: Attempting to modify content of published version.

**Domain Exception**: `ImmutableVersionException`

**Test Requirement**: Must reject update of published version content.

---

## Lifecycle Rules

**Allowed Transitions**:
- DRAFT → UNDER_REVIEW
- UNDER_REVIEW → APPROVED
- UNDER_REVIEW → DRAFT

**Forbidden Transitions**:
- APPROVED → DRAFT
- APPROVED → UNDER_REVIEW
- DRAFT → APPROVED

## Versioning Rules

**Version Creation Triggers**:
- Items change
- Answer key changes
- Scoring guidelines change
- Status transitions to APPROVED

## Authorization Assumptions

**Teacher**: Can modify own Assessment only
**School Admin**: Can modify any Assessment in school
**System Admin**: Emergency override access

## Domain Exceptions

```typescript
class AtLeastOneItemRequiredException extends DomainException {
  constructor() {
    super("Assessment must have at least one item");
  }
}

class InvalidScoreRangeException extends DomainException {
  constructor(score: number) {
    super(`Score must be >= 0. Invalid score: ${score}`);
  }
}
```

## Required Unit Tests

- `test_Assessment_creation_requires_tp_id`
- `test_Assessment_creation_requires_tp_version_no`
- `test_Assessment_creation_requires_school_id`
- `test_Assessment_creation_requires_at_least_one_item`
- `test_Assessment_item_numbers_must_be_sequential`
- `test_Assessment_item_score_must_be_valid`
- `test_Assessment_version_numbers_must_be_sequential`
- `test_Assessment_published_version_content_is_immutable`

## Aggregate Boundary

**Inside Aggregate**:
- Assessment (root)
- AssessmentVersion (entity)
- AssessmentItem (entity)
- AnswerKey (entity)
- ScoringGuideline (entity)

**Outside Aggregate**:
- TPSet (external reference)
- School (external reference)
- User (external reference)

## Aggregate Consistency Boundary

Assessment, all its versions, items, answer key, and scoring guidelines must be transactionally consistent.

## External Dependencies

**May Reference**:
- TPSet (read-only reference)
- School (read-only reference)
- User (read-only reference)

## Forbidden Dependencies

**Must Not Reference**:
- Evaluation (Evaluation references Assessment, not vice versa)
- Evidence (Evidence references Assessment, not vice versa)

## Ownership Rules

**Teacher Ownership**: Teacher owns Assessment they created
**School Ownership**: Assessment belongs to teacher's school
**Admin Ownership**: System Admin has emergency override

## School Isolation Rules

- Assessment queries must include school scope filter
- Cross-school Assessment access must return 404

---

# Aggregate: Evaluation

## Aggregate Purpose

Evaluation represents the assessment of a student's performance on an Assessment.

## Aggregate Root

Evaluation

## Entities

- Evaluation (root)
- EvaluationFeedbackHistory

## Value Objects

- AssessmentID
- StudentID
- TeacherID
- EvaluationID
- TotalScore
- PerformanceLevel
- TeacherFeedback
- EvaluationCriteriaScores
- WorkflowStatus
- SchoolID

## Invariants

### EV-INV-001: Assessment Reference Required

**Rule**: Evaluation must reference a valid Assessment.

**Violation**: Creation or modification without valid assessment_id.

**Domain Exception**: `AssessmentReferenceRequiredException`

**Test Requirement**: Must reject Evaluation creation without assessment_id.

---

### EV-INV-002: Student Reference Required

**Rule**: Evaluation must reference a valid Student.

**Violation**: Creation or modification without valid student_id.

**Domain Exception**: `StudentReferenceRequiredException`

**Test Requirement**: Must reject Evaluation creation without student_id.

---

### EV-INV-003: Teacher Reference Required

**Rule**: Evaluation must reference a valid Teacher (User).

**Violation**: Creation or modification without valid teacher_id.

**Domain Exception**: `TeacherReferenceRequiredException`

**Test Requirement**: Must reject Evaluation creation without teacher_id.

---

### EV-INV-004: School Ownership Required

**Rule**: Evaluation must belong to exactly one School (derived from student's school).

**Violation**: Creation or modification with invalid school scope.

**Domain Exception**: `SchoolOwnershipRequiredException`

**Test Requirement**: Must reject Evaluation creation for student from different school than teacher.

---

### EV-INV-005: Valid Score Range

**Rule**: Total score must be within valid range (0-100).

**Violation**: Setting total_score outside 0-100 range.

**Domain Exception**: `InvalidScoreRangeException`

**Test Requirement**: Must reject total_score < 0 or > 100.

---

### EV-INV-006: Sequential Revision Numbers

**Rule**: Revision numbers must be sequential integers.

**Violation**: Creating a revision with non-sequential revision_no.

**Domain Exception**: `NonSequentialRevisionNumberException`

**Test Requirement**: Must reject revision creation with non-sequential revision_no.

---

### EV-INV-007: Only One Current Revision

**Rule**: Only one revision may have is_current_version = true.

**Violation**: Attempting to activate a revision when another current revision exists.

**Domain Exception**: `MultipleCurrentVersionException`

**Test Requirement**: Must reject activation when another current revision exists.

---

### EV-INV-008: Feedback History Preservation

**Rule**: Feedback history must be preserved across all revisions.

**Violation**: Creating a revision without preserving feedback history.

**Domain Exception**: `FeedbackHistoryNotPreservedException`

**Test Requirement**: Must reject revision creation without feedback history preservation.

---

### EV-INV-009: No In-Place Updates

**Rule**: Evaluation content updates must create new revision, not update in-place.

**Violation**: Attempting to update evaluation content without creating new revision.

**Domain Exception**: `InPlaceUpdateException`

**Test Requirement**: Must reject in-place updates of evaluation content.

---

## Lifecycle Rules

**Allowed Transitions**:
- DRAFT → UNDER_REVIEW
- UNDER_REVIEW → APPROVED
- UNDER_REVIEW → DRAFT

**Forbidden Transitions**:
- APPROVED → DRAFT
- APPROVED → UNDER_REVIEW
- DRAFT → APPROVED

## Versioning Rules

**Revision Creation Triggers**:
- Feedback changes
- Score changes
- Performance level changes

**Revision Number Generation**:
- Initial evaluation: revision_no = 1
- Subsequent revisions: revision_no = max(revision_no) + 1

**Current Revision Rules**:
- Only one revision has is_current_version = true
- New revision becomes current when created
- Previous current revision is set to is_current_version = false

## Authorization Assumptions

**Teacher**: Can modify own Evaluation only
**School Admin**: Can modify any Evaluation in school
**System Admin**: Emergency override access

## Domain Exceptions

```typescript
class StudentReferenceRequiredException extends DomainException {
  constructor() {
    super("Evaluation must reference a valid Student");
  }
}

class TeacherReferenceRequiredException extends DomainException {
  constructor() {
    super("Evaluation must reference a valid Teacher");
  }
}

class NonSequentialRevisionNumberException extends DomainException {
  constructor(expected: number, actual: number) {
    super(`Revision numbers must be sequential. Expected: ${expected}, Actual: ${actual}`);
  }
}

class FeedbackHistoryNotPreservedException extends DomainException {
  constructor() {
    super("Feedback history must be preserved across revisions");
  }
}

class InPlaceUpdateException extends DomainException {
  constructor() {
    super("Evaluation updates must create new revision, not update in-place");
  }
}
```

## Required Unit Tests

- `test_Evaluation_creation_requires_assessment_id`
- `test_Evaluation_creation_requires_student_id`
- `test_Evaluation_creation_requires_teacher_id`
- `test_Evaluation_creation_requires_school_ownership`
- `test_Evaluation_total_score_must_be_valid_range`
- `test_Evaluation_revision_numbers_must_be_sequential`
- `test_Evaluation_only_one_current_revision_allowed`
- `test_Evaluation_feedback_history_must_be_preserved`
- `test_Evaluation_updates_must_create_new_revision`
- `test_Evaluation_in_place_updates_forbidden`

## Aggregate Boundary

**Inside Aggregate**:
- Evaluation (root)
- EvaluationFeedbackHistory (entity)

**Outside Aggregate**:
- Assessment (external reference)
- Student (external reference)
- User (external reference, as Teacher)

## Aggregate Consistency Boundary

Evaluation and all its feedback history must be transactionally consistent.

## External Dependencies

**May Reference**:
- Assessment (read-only reference)
- Student (read-only reference)
- User (read-only reference)

## Forbidden Dependencies

**Must Not Reference**:
- Evidence (Evidence references Evaluation, not vice versa)
- Achievement (Achievement references Evaluation, not vice versa)

## Ownership Rules

**Teacher Ownership**: Teacher owns Evaluation they created
**School Ownership**: Evaluation belongs to student's school
**Admin Ownership**: System Admin has emergency override

## School Isolation Rules

- Evaluation queries must include school scope filter (derived from student's school)
- Cross-school Evaluation access must return 404

---

# Aggregate: Evidence

## Aggregate Purpose

Evidence represents supporting materials (documents, images, videos, audio) uploaded for an Assessment or Evaluation.

## Aggregate Root

Evidence

## Entities

- Evidence (root)

## Value Objects

- EvidenceID
- StudentID
- AssessmentID
- FileID
- StorageKey
- FileName
- MimeType
- FileSizeBytes
- FileHash
- EvidenceType
- SchoolID
- UserID

## Invariants

### EVI-INV-001: Student Reference Required

**Rule**: Evidence must reference a valid Student.

**Violation**: Creation or modification without valid student_id.

**Domain Exception**: `StudentReferenceRequiredException`

**Test Requirement**: Must reject Evidence creation without student_id.

---

### EVI-INV-002: Assessment Reference Required

**Rule**: Evidence must reference a valid Assessment.

**Violation**: Creation or modification without valid assessment_id.

**Domain Exception**: `AssessmentReferenceRequiredException`

**Test Requirement**: Must reject Evidence creation without assessment_id.

---

### EVI-INV-003: School Ownership Required

**Rule**: Evidence must belong to exactly one School (derived from student's school).

**Violation**: Creation or modification with invalid school scope.

**Domain Exception**: `SchoolOwnershipRequiredException`

**Test Requirement**: Must reject Evidence creation for student from different school than uploader.

---

### EVI-INV-004: Valid File Size

**Rule**: File size must not exceed 50 MB (52,428,800 bytes).

**Violation**: Uploading file larger than 50 MB.

**Domain Exception**: `FileSizeExceededException`

**Test Requirement**: Must reject file upload larger than 50 MB.

---

### EVI-INV-005: Valid File Type

**Rule**: File type must be in allowed list (DOCUMENT, IMAGE, VIDEO, AUDIO).

**Violation**: Uploading file with disallowed type.

**Domain Exception**: `InvalidFileTypeException`

**Test Requirement**: Must reject file upload with disallowed type.

---

### EVI-INV-006: File Storage Required

**Rule**: Evidence must have a valid file stored in object storage.

**Violation**: Creation without valid file_id or storage_key.

**Domain Exception**: `FileStorageRequiredException`

**Test Requirement**: Must reject Evidence creation without file storage reference.

---

### EVI-INV-007: File Immutability

**Rule**: Evidence file content is immutable after upload.

**Violation**: Attempting to modify file content after upload.

**Domain Exception**: `FileImmutableException`

**Test Requirement**: Must reject file content modification after upload.

---

### EVI-INV-008: Valid MIME Type

**Rule**: MIME type must match file type.

**Violation**: MIME type does not match declared file type.

**Domain Exception**: `InvalidMimeTypeException`

**Test Requirement**: Must reject file upload with mismatched MIME type.

---

## Lifecycle Rules

**Allowed Transitions**:
- UPLOADED → PROCESSING
- PROCESSING → PROCESSED
- PROCESSING → ERROR

**Forbidden Transitions**:
- PROCESSED → UPLOADED
- ERROR → PROCESSED (must re-upload)

## Versioning Rules

**No Versioning**: Evidence does not have versioning. File content is immutable after upload.

## Authorization Assumptions

**Teacher**: Can upload Evidence for own students only
**School Admin**: Can access all Evidence in school
**System Admin**: Emergency override access

## Domain Exceptions

```typescript
class FileSizeExceededException extends DomainException {
  constructor(size: number, maxSize: number) {
    super(`File size exceeds maximum. Size: ${size}, Max: ${maxSize}`);
  }
}

class InvalidFileTypeException extends DomainException {
  constructor(type: string) {
    super(`Invalid file type: ${type}. Allowed: DOCUMENT, IMAGE, VIDEO, AUDIO`);
  }
}

class FileStorageRequiredException extends DomainException {
  constructor() {
    super("Evidence must have a valid file stored in object storage");
  }
}

class FileImmutableException extends DomainException {
  constructor() {
    super("Evidence file content is immutable after upload");
  }
}

class InvalidMimeTypeException extends DomainException {
  constructor(mimeType: string) {
    super(`Invalid MIME type: ${mimeType}`);
  }
}
```

## Required Unit Tests

- `test_Evidence_creation_requires_student_id`
- `test_Evidence_creation_requires_assessment_id`
- `test_Evidence_creation_requires_school_ownership`
- `test_Evidence_file_size_must_not_exceed_50mb`
- `test_Evidence_file_type_must_be_valid`
- `test_Evidence_file_storage_required`
- `test_Evidence_file_content_immutable_after_upload`
- `test_Evidence_mime_type_must_match_file_type`

## Aggregate Boundary

**Inside Aggregate**:
- Evidence (root)

**Outside Aggregate**:
- Student (external reference)
- Assessment (external reference)
- User (external reference, as uploader)

## Aggregate Consistency Boundary

Evidence entity is its own consistency boundary.

## External Dependencies

**May Reference**:
- Student (read-only reference)
- Assessment (read-only reference)
- User (read-only reference)

## Forbidden Dependencies

**Must Not Reference**:
- Evaluation (Evaluation may reference Evidence, not vice versa)

## Ownership Rules

**Teacher Ownership**: Teacher owns Evidence they uploaded
**School Ownership**: Evidence belongs to student's school
**Admin Ownership**: System Admin has emergency override

## School Isolation Rules

- Evidence queries must include school scope filter (derived from student's school)
- Cross-school Evidence access must return 404

---

# Aggregate: Achievement

## Aggregate Purpose

Achievement represents system-calculated student achievement based on Evaluations.

## Aggregate Root

Achievement

## Entities

- Achievement (root)
- AchievementCriteria
- AchievementSnapshot

## Value Objects

- AchievementID
- StudentID
- CompetencyLevel
- ProgressPercentage
- AchievementDate
- SchoolID

## Invariants

### ACH-INV-001: Student Reference Required

**Rule**: Achievement must reference a valid Student.

**Violation**: Creation or modification without valid student_id.

**Domain Exception**: `StudentReferenceRequiredException`

**Test Requirement**: Must reject Achievement creation without student_id.

---

### ACH-INV-002: School Ownership Required

**Rule**: Achievement must belong to exactly one School (derived from student's school).

**Violation**: Creation or modification with invalid school scope.

**Domain Exception**: `SchoolOwnershipRequiredException`

**Test Requirement**: Must reject Achievement creation for student from different school.

---

### ACH-INV-003: Valid Progress Range

**Rule**: Progress percentage must be between 0-100.

**Violation**: Setting progress_percentage outside 0-100 range.

**Domain Exception**: `InvalidProgressRangeException`

**Test Requirement**: Must reject progress_percentage < 0 or > 100.

---

### ACH-INV-004: System-Generated Only

**Rule**: Achievement must be system-calculated, not manually created.

**Violation**: Manual creation or modification of Achievement.

**Domain Exception**: `ManualCreationException`

**Test Requirement**: Must reject manual Achievement creation. Must reject manual Achievement modification.

---

### ACH-INV-005: Read-Only

**Rule**: Achievement is read-only after calculation.

**Violation**: Attempting to modify Achievement after calculation.

**Domain Exception**: `ReadOnlyException`

**Test Requirement**: Must reject Achievement modification after calculation.

---

## Lifecycle Rules

**No Lifecycle Transitions**: Achievement is system-calculated and read-only.

## Versioning Rules

**No Versioning**: Achievement does not have versioning. Snapshots are created periodically for audit.

## Authorization Assumptions

**Teacher**: Can read Achievement for own students only
**School Admin**: Can read all Achievement in school
**System Admin**: Can read all Achievement across all schools

## Domain Exceptions

```typescript
class InvalidProgressRangeException extends DomainException {
  constructor(progress: number) {
    super(`Progress percentage must be between 0-100. Invalid: ${progress}`);
  }
}

class ManualCreationException extends DomainException {
  constructor() {
    super("Achievement must be system-calculated, not manually created");
  }
}

class ReadOnlyException extends DomainException {
  constructor() {
    super("Achievement is read-only after calculation");
  }
}
```

## Required Unit Tests

- `test_Achievement_creation_requires_student_id`
- `test_Achievement_creation_requires_school_ownership`
- `test_Achievement_progress_percentage_must_be_valid_range`
- `test_Achievement_must_be_system_calculated`
- `test_Achievement_is_read_only_after_calculation`

## Aggregate Boundary

**Inside Aggregate**:
- Achievement (root)
- AchievementCriteria (entity)
- AchievementSnapshot (entity)

**Outside Aggregate**:
- Student (external reference)

## Aggregate Consistency Boundary

Achievement, criteria, and snapshots must be transactionally consistent.

## External Dependencies

**May Reference**:
- Student (read-only reference)

## Forbidden Dependencies

**Must Not Reference**:
- Evaluation (Achievement references Evaluation data, but not as direct dependency)
- NarrativeReport (NarrativeReport references Achievement, not vice versa)

## Ownership Rules

**Teacher Ownership**: No single owner (system-generated)
**School Ownership**: Achievement belongs to student's school
**Admin Ownership**: System Admin has read-only access

## School Isolation Rules

- Achievement queries must include school scope filter (derived from student's school)
- Cross-school Achievement access must return 404

---

# Aggregate: NarrativeReport

## Aggregate Purpose

NarrativeReport represents a narrative progress report for a student, integrating Achievement data.

## Aggregate Root

NarrativeReport

## Entities

- NarrativeReport (root)
- NarrativeReportVersion

## Value Objects

- ReportID
- StudentID
- ReportPeriod
- NarrativeContent
- AchievementData
- WorkflowStatus
- SchoolID
- UserID

## Invariants

### NR-INV-001: Student Reference Required

**Rule**: NarrativeReport must reference a valid Student.

**Violation**: Creation or modification without valid student_id.

**Domain Exception**: `StudentReferenceRequiredException`

**Test Requirement**: Must reject NarrativeReport creation without student_id.

---

### NR-INV-002: School Ownership Required

**Rule**: NarrativeReport must belong to exactly one School (derived from student's school).

**Violation**: Creation or modification with invalid school scope.

**Domain Exception**: `SchoolOwnershipRequiredException`

**Test Requirement**: Must reject NarrativeReport creation for student from different school than generator.

---

### NR-INV-003: At Least One Version Required

**Rule**: NarrativeReport must have at least one version.

**Violation**: Creation without initial version.

**Domain Exception**: `AtLeastOneVersionRequiredException`

**Test Requirement**: Must reject NarrativeReport creation without initial version.

---

### NR-INV-004: Only One Current Version

**Rule**: Only one version may have is_current_version = true.

**Violation**: Attempting to activate a version when another current version exists.

**Domain Exception**: `MultipleCurrentVersionException`

**Test Requirement**: Must reject activation when another current version exists.

---

### NR-INV-005: Sequential Version Numbers

**Rule**: Version numbers must be sequential integers.

**Violation**: Creating a version with non-sequential version number.

**Domain Exception**: `NonSequentialVersionNumberException`

**Test Requirement**: Must reject version creation with non-sequential version number.

---

### NR-INV-006: Published Version Immutability

**Rule**: Published versions are immutable.

**Violation**: Attempting to modify content of published version.

**Domain Exception**: `ImmutableVersionException`

**Test Requirement**: Must reject update of published version content.

---

## Lifecycle Rules

**Allowed Transitions**:
- DRAFT → UNDER_REVIEW
- UNDER_REVIEW → APPROVED
- UNDER_REVIEW → DRAFT

**Forbidden Transitions**:
- APPROVED → DRAFT
- APPROVED → UNDER_REVIEW
- DRAFT → APPROVED

## Versioning Rules

**Version Creation Triggers**:
- Content changes
- Achievement data refresh
- Status transitions to APPROVED

## Authorization Assumptions

**Teacher**: Can generate NarrativeReport for own students only
**School Admin**: Can access all NarrativeReport in school
**System Admin**: Emergency override access

## Domain Exceptions

```typescript
class StudentReferenceRequiredException extends DomainException {
  constructor() {
    super("NarrativeReport must reference a valid Student");
  }
}
```

## Required Unit Tests

- `test_NarrativeReport_creation_requires_student_id`
- `test_NarrativeReport_creation_requires_school_ownership`
- `test_NarrativeReport_creation_requires_initial_version`
- `test_NarrativeReport_only_one_current_version_allowed`
- `test_NarrativeReport_version_numbers_must_be_sequential`
- `test_NarrativeReport_published_version_content_is_immutable`

## Aggregate Boundary

**Inside Aggregate**:
- NarrativeReport (root)
- NarrativeReportVersion (entity)

**Outside Aggregate**:
- Student (external reference)
- Achievement (external reference, via achievement_data)
- User (external reference, as generator)

## Aggregate Consistency Boundary

NarrativeReport and all its versions must be transactionally consistent.

## External Dependencies

**May Reference**:
- Student (read-only reference)
- Achievement (read-only reference, via achievement_data)
- User (read-only reference)

## Forbidden Dependencies

**Must Not Reference**:
- None (NarrativeReport is leaf aggregate in Reporting)

## Ownership Rules

**Teacher Ownership**: Teacher owns NarrativeReport they generated
**School Ownership**: NarrativeReport belongs to student's school
**Admin Ownership**: System Admin has emergency override

## School Isolation Rules

- NarrativeReport queries must include school scope filter (derived from student's school)
- Cross-school NarrativeReport access must return 404

---

# Aggregate: School

## Aggregate Purpose

School represents a school in the multi-tenant NUSA Platform.

## Aggregate Root

School

## Entities

- School (root)

## Value Objects

- SchoolID
- SchoolName
- SchoolCode
- SchoolAddress
- SchoolPhone
- SchoolEmail
- SchoolStatus

## Invariants

### SCH-INV-001: Unique School Code

**Rule**: School code must be unique across all schools.

**Violation**: Creating school with duplicate school_code.

**Domain Exception**: `DuplicateSchoolCodeException`

**Test Requirement**: Must reject school creation with duplicate school_code.

---

### SCH-INV-002: Valid School Status

**Rule**: School status must be one of: ACTIVE, INACTIVE.

**Violation**: Setting status to invalid value.

**Domain Exception**: `InvalidSchoolStatusException`

**Test Requirement**: Must reject invalid status values.

---

### SCH-INV-003: School Code Immutability

**Rule**: School code cannot be modified after creation.

**Violation**: Attempting to change school_code after creation.

**Domain Exception**: `SchoolCodeImmutableException`

**Test Requirement**: Must reject modification of school_code after creation.

---

## Lifecycle Rules

**Allowed Transitions**:
- ACTIVE → INACTIVE
- INACTIVE → ACTIVE

**Forbidden Transitions**:
- None

## Versioning Rules

**No Versioning**: School does not have versioning.

## Authorization Assumptions

**System Admin**: Can create, modify, delete schools
**School Admin**: Can read own school only
**Teacher**: Can read own school only

## Domain Exceptions

```typescript
class DuplicateSchoolCodeException extends DomainException {
  constructor(code: string) {
    super(`School code must be unique. Duplicate: ${code}`);
  }
}

class InvalidSchoolStatusException extends DomainException {
  constructor(status: string) {
    super(`Invalid school status: ${status}. Allowed: ACTIVE, INACTIVE`);
  }
}

class SchoolCodeImmutableException extends DomainException {
  constructor() {
    super("School code cannot be modified after creation");
  }
}
```

## Required Unit Tests

- `test_School_code_must_be_unique`
- `test_School_status_must_be_valid`
- `test_School_code_cannot_be_modified_after_creation`

## Aggregate Boundary

**Inside Aggregate**:
- School (root)

**Outside Aggregate**:
- None

## Aggregate Consistency Boundary

School entity is its own consistency boundary.

## External Dependencies

**May Reference**:
- None

## Forbidden Dependencies

**Must Not Reference**:
- None

## Ownership Rules

**System Admin Ownership**: System Admin owns all schools
**School Ownership**: School belongs to System Admin
**Teacher Ownership**: Teachers belong to School

## School Isolation Rules

- School queries must include school scope filter for school-scoped entities
- System Admin bypasses school scope filter

---

# Aggregate: User

## Aggregate Purpose

User represents a user in the NUSA Platform (teachers, school admins, system admins).

## Aggregate Root

User

## Entities

- User (root)

## Value Objects

- UserID
- SchoolID
- Username
- Email
- PasswordHash
- FullName
- UserRole
- UserStatus

## Invariants

### USR-INV-001: School Ownership Required

**Rule**: User must belong to exactly one School.

**Violation**: Creation or modification without valid school_id.

**Domain Exception**: `SchoolOwnershipRequiredException`

**Test Requirement**: Must reject User creation without school_id.

---

### USR-INV-002: Unique Username

**Rule**: Username must be unique across all users.

**Violation**: Creating user with duplicate username.

**Domain Exception**: `DuplicateUsernameException`

**Test Requirement**: Must reject user creation with duplicate username.

---

### USR-INV-003: Unique Email

**Rule**: Email must be unique across all users.

**Violation**: Creating user with duplicate email.

**Domain Exception**: `DuplicateEmailException`

**Test Requirement**: Must reject user creation with duplicate email.

---

### USR-INV-004: Valid User Role

**Rule**: User role must be one of: TEACHER, SCHOOL_ADMIN, SYSTEM_ADMIN.

**Violation**: Setting role to invalid value.

**Domain Exception**: `InvalidUserRoleException`

**Test Requirement**: Must reject invalid role values.

---

### USR-INV-005: Valid User Status

**Rule**: User status must be one of: ACTIVE, INACTIVE.

**Violation**: Setting status to invalid value.

**Domain Exception**: `InvalidUserStatusException`

**Test Requirement**: Must reject invalid status values.

---

### USR-INV-006: Username Immutability

**Rule**: Username cannot be modified after creation.

**Violation**: Attempting to change username after creation.

**Domain Exception**: `UsernameImmutableException`

**Test Requirement**: Must reject modification of username after creation.

---

### USR-INV-007: Email Immutability

**Rule**: Email cannot be modified after creation.

**Violation**: Attempting to change email after creation.

**Domain Exception**: `EmailImmutableException`

**Test Requirement**: Must reject modification of email after creation.

---

### USR-INV-008: School Immutability

**Rule**: School cannot be modified after creation.

**Violation**: Attempting to change school_id after creation.

**Domain Exception**: `SchoolImmutableException`

**Test Requirement**: Must reject modification of school_id after creation.

---

## Lifecycle Rules

**Allowed Transitions**:
- ACTIVE → INACTIVE
- INACTIVE → ACTIVE

**Forbidden Transitions**:
- None

## Versioning Rules

**No Versioning**: User does not have versioning.

## Authorization Assumptions

**System Admin**: Can create, modify, delete users
**School Admin**: Can create, modify users in own school
**Teacher**: Can read own profile only

## Domain Exceptions

```typescript
class DuplicateUsernameException extends DomainException {
  constructor(username: string) {
    super(`Username must be unique. Duplicate: ${username}`);
  }
}

class DuplicateEmailException extends DomainException {
  constructor(email: string) {
    super(`Email must be unique. Duplicate: ${email}`);
  }
}

class InvalidUserRoleException extends DomainException {
  constructor(role: string) {
    super(`Invalid user role: ${role}. Allowed: TEACHER, SCHOOL_ADMIN, SYSTEM_ADMIN`);
  }
}

class InvalidUserStatusException extends DomainException {
  constructor(status: string) {
    super(`Invalid user status: ${status}. Allowed: ACTIVE, INACTIVE`);
  }
}

class UsernameImmutableException extends DomainException {
  constructor() {
    super("Username cannot be modified after creation");
  }
}

class EmailImmutableException extends DomainException {
  constructor() {
    super("Email cannot be modified after creation");
  }
}

class SchoolImmutableException extends DomainException {
  constructor() {
    super("School cannot be modified after creation");
  }
}
```

## Required Unit Tests

- `test_User_creation_requires_school_id`
- `test_User_username_must_be_unique`
- `test_User_email_must_be_unique`
- `test_User_role_must_be_valid`
- `test_User_status_must_be_valid`
- `test_User_username_cannot_be_modified_after_creation`
- `test_User_email_cannot_be_modified_after_creation`
- `test_User_school_cannot_be_modified_after_creation`

## Aggregate Boundary

**Inside Aggregate**:
- User (root)

**Outside Aggregate**:
- School (external reference)
- Role (external reference)
- Permission (external reference)

## Aggregate Consistency Boundary

User entity is its own consistency boundary.

## External Dependencies

**May Reference**:
- School (read-only reference)
- Role (read-only reference)
- Permission (read-only reference)

## Forbidden Dependencies

**Must Not Reference**:
- None

## Ownership Rules

**System Admin Ownership**: System Admin owns all users
**School Ownership**: User belongs to School
**Teacher Ownership**: Teacher is a User with TEACHER role

## School Isolation Rules

- User queries must include school scope filter for school-scoped entities
- System Admin bypasses school scope filter

---

# Aggregate: Role

## Aggregate Purpose

Role represents a role in the RBAC system.

## Aggregate Root

Role

## Entities

- Role (root)

## Value Objects

- RoleID
- RoleName
- RoleDescription

## Invariants

### ROL-INV-001: Unique Role Name

**Rule**: Role name must be unique across all roles.

**Violation**: Creating role with duplicate role_name.

**Domain Exception**: `DuplicateRoleNameException`

**Test Requirement**: Must reject role creation with duplicate role_name.

---

### ROL-INV-002: Role Name Immutability

**Rule**: Role name cannot be modified after creation.

**Violation**: Attempting to change role_name after creation.

**Domain Exception**: `RoleNameImmutableException`

**Test Requirement**: Must reject modification of role_name after creation.

---

## Lifecycle Rules

**No Lifecycle Transitions**: Role is static configuration.

## Versioning Rules

**No Versioning**: Role does not have versioning.

## Authorization Assumptions

**System Admin**: Can create, modify, delete roles
**School Admin**: Can read roles only
**Teacher**: Can read roles only

## Domain Exceptions

```typescript
class DuplicateRoleNameException extends DomainException {
  constructor(name: string) {
    super(`Role name must be unique. Duplicate: ${name}`);
  }
}

class RoleNameImmutableException extends DomainException {
  constructor() {
    super("Role name cannot be modified after creation");
  }
}
```

## Required Unit Tests

- `test_Role_name_must_be_unique`
- `test_Role_name_cannot_be_modified_after_creation`

## Aggregate Boundary

**Inside Aggregate**:
- Role (root)

**Outside Aggregate**:
- Permission (external reference)

## Aggregate Consistency Boundary

Role entity is its own consistency boundary.

## External Dependencies

**May Reference**:
- Permission (read-only reference)

## Forbidden Dependencies

**Must Not Reference**:
- None

## Ownership Rules

**System Admin Ownership**: System Admin owns all roles
**School Ownership**: Roles are global (not school-scoped)
**Teacher Ownership**: Teachers are assigned to roles (not owned by role)

## School Isolation Rules

- Roles are global (not school-scoped)
- No school scope filtering for roles

---

# Aggregate: Permission

## Aggregate Purpose

Permission represents a permission in the RBAC system.

## Aggregate Root

Permission

## Entities

- Permission (root)

## Value Objects

- PermissionID
- Resource
- Action
- PermissionDescription

## Invariants

### PER-INV-001: Unique Resource-Action Combination

**Rule**: Resource and action combination must be unique across all permissions.

**Violation**: Creating permission with duplicate resource-action combination.

**Domain Exception**: `DuplicatePermissionException`

**Test Requirement**: Must reject permission creation with duplicate resource-action combination.

---

### PER-INV-002: Resource Immutability

**Rule**: Resource cannot be modified after creation.

**Violation**: Attempting to change resource after creation.

**Domain Exception**: `ResourceImmutableException`

**Test Requirement**: Must reject modification of resource after creation.

---

### PER-INV-003: Action Immutability

**Rule**: Action cannot be modified after creation.

**Violation**: Attempting to change action after creation.

**Domain Exception**: `ActionImmutableException`

**Test Requirement**: Must reject modification of action after creation.

---

## Lifecycle Rules

**No Lifecycle Transitions**: Permission is static configuration.

## Versioning Rules

**No Versioning**: Permission does not have versioning.

## Authorization Assumptions

**System Admin**: Can create, modify, delete permissions
**School Admin**: Can read permissions only
**Teacher**: Can read permissions only

## Domain Exceptions

```typescript
class DuplicatePermissionException extends DomainException {
  constructor(resource: string, action: string) {
    super(`Permission must be unique. Duplicate: ${resource}:${action}`);
  }
}

class ResourceImmutableException extends DomainException {
  constructor() {
    super("Resource cannot be modified after creation");
  }
}

class ActionImmutableException extends DomainException {
  constructor() {
    super("Action cannot be modified after creation");
  }
}
```

## Required Unit Tests

- `test_Permission_resource_action_must_be_unique`
- `test_Permission_resource_cannot_be_modified_after_creation`
- `test_Permission_action_cannot_be_modified_after_creation`

## Aggregate Boundary

**Inside Aggregate**:
- Permission (root)

**Outside Aggregate**:
- Role (external reference)

## Aggregate Consistency Boundary

Permission entity is its own consistency boundary.

## External Dependencies

**May Reference**:
- None

## Forbidden Dependencies

**Must Not Reference**:
- None

## Ownership Rules

**System Admin Ownership**: System Admin owns all permissions
**School Ownership**: Permissions are global (not school-scoped)
**Teacher Ownership**: Teachers are granted permissions via roles (not owned by permission)

## School Isolation Rules

- Permissions are global (not school-scoped)
- No school scope filtering for permissions

---

# Aggregate: PermissionChange

## Aggregate Purpose

PermissionChange represents an audit trail of permission modifications.

## Aggregate Root

PermissionChange

## Entities

- PermissionChange (root)

## Value Objects

- PermissionChangeID
- RoleID
- Resource
- Action
- ChangeType
- ChangedBy
- ChangedAt
- Reason

## Invariants

### PC-INV-001: Valid Change Type

**Rule**: Change type must be one of: GRANT, REVOKE.

**Violation**: Setting change_type to invalid value.

**Domain Exception**: `InvalidChangeTypeException`

**Test Requirement**: Must reject invalid change_type values.

---

### PC-INV-002: Changed By Required

**Rule**: PermissionChange must reference a valid User who made the change.

**Violation**: Creation without valid changed_by.

**Domain Exception**: `ChangedByRequiredException`

**Test Requirement**: Must reject PermissionChange creation without changed_by.

---

### PC-INV-003: Role Reference Required

**Rule**: PermissionChange must reference a valid Role.

**Violation**: Creation without valid role_id.

**Domain Exception**: `RoleReferenceRequiredException`

**Test Requirement**: Must reject PermissionChange creation without role_id.

---

### PC-INV-004: Immutability

**Rule**: PermissionChange is immutable after creation.

**Violation**: Attempting to modify PermissionChange after creation.

**Domain Exception**: `ImmutableException`

**Test Requirement**: Must reject modification of PermissionChange after creation.

---

## Lifecycle Rules

**No Lifecycle Transitions**: PermissionChange is audit trail (immutable).

## Versioning Rules

**No Versioning**: PermissionChange does not have versioning.

## Authorization Assumptions

**System Admin**: Can read all PermissionChange
**School Admin**: Can read PermissionChange for school roles only
**Teacher**: Cannot read PermissionChange

## Domain Exceptions

```typescript
class InvalidChangeTypeException extends DomainException {
  constructor(type: string) {
    super(`Invalid change type: ${type}. Allowed: GRANT, REVOKE`);
  }
}

class ChangedByRequiredException extends DomainException {
  constructor() {
    super("PermissionChange must reference a valid User who made the change");
  }
}

class RoleReferenceRequiredException extends DomainException {
  constructor() {
    super("PermissionChange must reference a valid Role");
  }
}

class ImmutableException extends DomainException {
  constructor() {
    super("PermissionChange is immutable after creation");
  }
}
```

## Required Unit Tests

- `test_PermissionChange_change_type_must_be_valid`
- `test_PermissionChange_changed_by_required`
- `test_PermissionChange_role_reference_required`
- `test_PermissionChange_immutable_after_creation`

## Aggregate Boundary

**Inside Aggregate**:
- PermissionChange (root)

**Outside Aggregate**:
- Role (external reference)
- User (external reference, as changed_by)

## Aggregate Consistency Boundary

PermissionChange entity is its own consistency boundary.

## External Dependencies

**May Reference**:
- Role (read-only reference)
- User (read-only reference)

## Forbidden Dependencies

**Must Not Reference**:
- None

## Ownership Rules

**System Admin Ownership**: System Admin owns all PermissionChange
**School Ownership**: PermissionChange is global (not school-scoped)
**Teacher Ownership**: Teachers cannot modify permissions

## School Isolation Rules

- PermissionChange is global (not school-scoped)
- No school scope filtering for PermissionChange

---

# Cross Aggregate Invariants

## CROSS-INV-001: Assessment Must Reference Approved TPSet Version

**Rule**: Assessment must reference an approved TPSet version.

**Violation**: Creating Assessment that references non-approved TPSet version.

**Domain Exception**: `TPSetVersionNotApprovedException`

**Test Requirement**: Must reject Assessment creation if referenced TPSet version is not APPROVED.

---

## CROSS-INV-002: ATPSet Must Reference Valid TPSet

**Rule**: ATPSet must reference a valid TPSet that exists.

**Violation**: Creating ATPSet that references non-existent TPSet.

**Domain Exception**: `TPSetNotFoundException`

**Test Requirement**: Must reject ATPSet creation if referenced TPSet does not exist.

---

## CROSS-INV-003: ModulAjarSet Must Reference Valid ATPSet

**Rule**: ModulAjarSet must reference a valid ATPSet that exists.

**Violation**: Creating ModulAjarSet that references non-existent ATPSet.

**Domain Exception**: `ATPSetNotFoundException`

**Test Requirement**: Must reject ModulAjarSet creation if referenced ATPSet does not exist.

---

## CROSS-INV-004: Evaluation Must Reference Valid Assessment

**Rule**: Evaluation must reference a valid Assessment that exists.

**Violation**: Creating Evaluation that references non-existent Assessment.

**Domain Exception**: `AssessmentNotFoundException`

**Test Requirement**: Must reject Evaluation creation if referenced Assessment does not exist.

---

## CROSS-INV-005: Evaluation Must Reference Same School as Student

**Rule**: Evaluation must be created by Teacher from same school as Student.

**Violation**: Teacher from School A creating Evaluation for Student from School B.

**Domain Exception**: `CrossSchoolEvaluationException`

**Test Requirement**: Must reject Evaluation creation if teacher.school_id != student.school_id.

---

## CROSS-INV-006: Evidence Must Reference Same School as Student

**Rule**: Evidence must be uploaded by Teacher from same school as Student.

**Violation**: Teacher from School A uploading Evidence for Student from School B.

**Domain Exception**: `CrossSchoolEvidenceException`

**Test Requirement**: Must reject Evidence upload if uploader.school_id != student.school_id.

---

## CROSS-INV-007: Achievement Cannot Be Calculated Without Evaluation

**Rule**: Achievement calculation requires at least one Evaluation for the student.

**Violation**: Attempting to calculate Achievement without any Evaluations.

**Domain Exception**: `EvaluationRequiredException`

**Test Requirement**: Must reject Achievement calculation if no Evaluations exist for student.

---

## CROSS-INV-008: NarrativeReport Cannot Be Finalized Without Achievement

**Rule**: NarrativeReport generation requires Achievement data to exist.

**Violation**: Attempting to generate NarrativeReport without Achievement data.

**Domain Exception**: `AchievementRequiredException`

**Test Requirement**: Must reject NarrativeReport generation if Achievement data does not exist for student.

---

## CROSS-INV-009: TPSet Version Consistency

**Rule**: ATPSet must reference the same TPSet version that was current when ATPSet was created.

**Violation**: ATPSet referencing a TPSet version that was not current when ATPSet was created.

**Domain Exception**: `TPSetVersionConsistencyException`

**Test Requirement**: Must validate ATPSet references current TPSet version at creation time.

---

## CROSS-INV-010: Assessment Version Consistency

**Rule**: Evaluation must reference the same Assessment version that was current when Evaluation was created.

**Violation**: Evaluation referencing an Assessment version that was not current when Evaluation was created.

**Domain Exception**: `AssessmentVersionConsistencyException`

**Test Requirement**: Must validate Evaluation references current Assessment version at creation time.

---

## Cross Aggregate Domain Exceptions

```typescript
class TPSetVersionNotApprovedException extends DomainException {
  constructor() {
    super("Assessment must reference an approved TPSet version");
  }
}

class TPSetNotFoundException extends DomainException {
  constructor(tpSetId: string) {
    super(`TPSet not found: ${tpSetId}`);
  }
}

class ATPSetNotFoundException extends DomainException {
  constructor(atpSetId: string) {
    super(`ATPSet not found: ${atpSetId}`);
  }
}

class AssessmentNotFoundException extends DomainException {
  constructor(assessmentId: string) {
    super(`Assessment not found: ${assessmentId}`);
  }
}

class CrossSchoolEvaluationException extends DomainException {
  constructor() {
    super("Evaluation must be created by Teacher from same school as Student");
  }
}

class CrossSchoolEvidenceException extends DomainException {
  constructor() {
    super("Evidence must be uploaded by Teacher from same school as Student");
  }
}

class EvaluationRequiredException extends DomainException {
  constructor() {
    super("Achievement calculation requires at least one Evaluation");
  }
}

class AchievementRequiredException extends DomainException {
  constructor() {
    super("NarrativeReport generation requires Achievement data");
  }
}

class TPSetVersionConsistencyException extends DomainException {
  constructor() {
    super("ATPSet must reference the TPSet version that was current when ATPSet was created");
  }
}

class AssessmentVersionConsistencyException extends DomainException {
  constructor() {
    super("Evaluation must reference the Assessment version that was current when Evaluation was created");
  }
}
```

## Cross Aggregate Required Unit Tests

- `test_Assessment_must_reference_approved_tp_set_version`
- `test_ATPSet_must_reference_valid_tp_set`
- `test_ModulAjarSet_must_reference_valid_atp_set`
- `test_Evaluation_must_reference_valid_assessment`
- `test_Evaluation_must_reference_same_school_as_student`
- `test_Evidence_must_reference_same_school_as_student`
- `test_Achievement_cannot_be_calculated_without_evaluation`
- `test_NarrativeReport_cannot_be_finalized_without_achievement`
- `test_TPSet_version_consistency_across_aggregates`
- `test_Assessment_version_consistency_across_aggregates`

---

# Document Control

**Version**: 1.0
**Effective Date**: 2026-06-09
**Status**: FINAL SOURCE OF TRUTH FOR DOMAIN INVARIANTS
**Parent Documents**:
- Architecture Freeze v2
- Database Schema Freeze v1
- Repository Governance

**Change Process**:
1. Submit Domain Invariant Change Request (DICR)
2. Review by Principal DDD Architect
3. Update document if approved
4. Increment version number
5. Communicate changes to all stakeholders

**Approval Required**:
- Principal DDD Architect
- Principal Backend Architect
- Principal QA Architect

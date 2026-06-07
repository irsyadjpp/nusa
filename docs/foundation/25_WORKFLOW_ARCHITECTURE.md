# 25_WORKFLOW_ARCHITECTURE.md

## Foundation Document for NUSA Education Platform

**Version**: 1.0
**Date**: June 2026
**Status**: FOUNDATION DOCUMENT
**Alignment**: Aligned with 06_APPLICATION_ARCHITECTURE.md, 14_DATABASE_SCHEMA.md, 24_TP_GENERATION_ARCHITECTURE.md

**Purpose**: Define a reusable workflow architecture for all educational artifacts in NUSA, establishing consistent approval workflows, audit trails, and notification mechanisms across TP, ATP, Modul Ajar, Assessment, Rubric, and Narrative Report.

---

# SECTION 1 — Executive Summary

## Why Workflow Architecture Matters

All educational artifacts in NUSA follow a consistent approval workflow to ensure:
- Human oversight and accountability
- Quality control before artifacts become official
- Complete audit trail for compliance
- Consistent user experience across artifact types
- Reusable workflow logic to avoid duplication

## Core Principles

- **Human Approval Mandatory**: AI-generated artifacts require human approval before becoming official
- **AI Never Publishes Official Artifacts**: AI can generate recommendations but never auto-approves
- **Consistent States**: All artifacts use the same workflow states
- **Complete Audit Trail**: Every workflow action is logged
- **AI Generation Traceability**: Every workflow action must be traceable back to originating AI generation
- **Reusable Engine**: Single workflow service handles all artifact types
- **Modular Monolith**: Workflow engine is a module within the monolith, not a microservice

---

# SECTION 2 — Workflow Principles

## Principle 1: Human Approval is Mandatory

**Rule**: All educational artifacts must be approved by a human user before becoming official.

**Implications**:
- AI-generated artifacts start in DRAFT status
- DRAFT artifacts cannot be used for downstream generation
- Only APPROVED artifacts can be used in production
- Approval requires explicit user action

**Rationale**: Teachers and administrators must maintain authority over educational content to ensure quality, alignment with local context, and compliance with regulations.

## Principle 2: AI Never Publishes Official Artifacts

**Rule**: AI agents can generate artifact recommendations but cannot automatically approve or publish them.

**Implications**:
- AI output always goes to DRAFT status
- No auto-approval based on confidence scores
- No auto-publication based on time or other triggers
- Human review is non-bypassable

**Rationale**: AI may generate high-quality content, but human judgment is required for educational decisions involving student learning.

## Principle 3: Consistent Workflow Across Artifacts

**Rule**: All artifact types (TP, ATP, Modul Ajar, Assessment, Rubric, Narrative Report) use the same workflow states and transitions.

**Implications**:
- Single workflow engine handles all artifact types
- Consistent UI patterns across artifact review screens
- Unified audit trail across all artifacts
- Simplified training for users

**Rationale**: Consistency reduces complexity, improves user experience, and simplifies maintenance.

## Principle 4: AI Generation Traceability Chain

**Rule**: Every workflow action must be traceable back to the originating AI generation.

**Implications**:
- All artifact sets (tp_sets, atp_sets, modul_ajar_sets) include ai_generation_id field
- ai_generation_log becomes the root traceability record for all AI-generated artifacts
- Complete audit trail from AI generation → Artifact → Workflow History
- Enables debugging, cost analysis, and quality improvement

**Traceability Chain Example**:

```
AI Generation (ai_generation_logs)
  ↓
TP Set (tp_sets.ai_generation_id)
  ↓
Workflow History (workflow_history.artifact_id)
```

**Audit Trail Example**:

```
1. AI generates TP Set (ai_generation_logs.id = gen_123)
2. TP Set created (tp_sets.ai_generation_id = gen_123)
3. Teacher reviews TP Set (workflow_history.artifact_id = tp_set_456)
4. Teacher edits TP Set (workflow_history.artifact_id = tp_set_456)
5. Teacher approves TP Set (workflow_history.artifact_id = tp_set_456)
```

**Rationale**: Complete traceability ensures accountability, enables cost governance, supports debugging, and provides data for AI quality improvement.

---

# SECTION 3 — Workflow States

## Workflow State Standard

All workflow-driven artifacts in NUSA use the same standardized workflow states to ensure consistency across the system.

### Standardized States

| State | Description |
|-------|-------------|
| DRAFT | Initial state for newly created artifacts (AI-generated or manually created) |
| UNDER_REVIEW | Artifact submitted for human review |
| APPROVED | Artifact approved and becomes official |
| REJECTED | Artifact rejected during review |
| ARCHIVED | Historical artifact no longer in active use |

### Applicable Artifacts

The standardized workflow states apply to:
- TP Sets (tp_sets)
- TP Items (tp)
- ATP Sets (atp_sets)
- ATP Items (atp)
- Modul Ajar Sets (modul_ajar_sets)
- Modul Ajar Items (modul_ajar)
- Assessments (assessments)
- Rubrics (rubrics)
- Narrative Reports (narrative_reports)
- Future artifacts

### State Transition Matrix

**Allowed Transitions**:

```
DRAFT → UNDER_REVIEW
UNDER_REVIEW → APPROVED
UNDER_REVIEW → REJECTED
APPROVED → ARCHIVED
REJECTED → DRAFT
```

**Transition Rules**:

| From | To | Trigger | Conditions |
|------|----|---------|------------|
| DRAFT | UNDER_REVIEW | User submits for review | Artifact has required fields populated |
| UNDER_REVIEW | APPROVED | Reviewer approves | Reviewer has approval permissions |
| UNDER_REVIEW | REJECTED | Reviewer rejects | Reviewer provides rejection reason |
| APPROVED | ARCHIVED | System or user action | Artifact superseded or period ended |
| REJECTED | DRAFT | User edits and resubmits | User makes changes to address rejection |

**Forbidden Transitions**:

- DRAFT → APPROVED (must go through review)
- DRAFT → REJECTED (must go through review first)
- DRAFT → ARCHIVED (must be approved first)
- UNDER_REVIEW → ARCHIVED (must be approved or rejected first)
- APPROVED → DRAFT (regeneration creates new version, does not change existing)
- APPROVED → REJECTED (approved artifacts cannot be rejected)
- ARCHIVED → Any state (archived is terminal)

### State Transition Diagram

```
┌─────────┐
│ DRAFT   │
└────┬────┘
     │ (submit for review)
     ↓
┌─────────────┐
│UNDER_REVIEW │
└─────┬───────┘
      │
      ├─→ APPROVED (approve)
      │      ↓
      │   ┌─────────┐
      │   │ARCHIVED │
      │   └─────────┘
      │
      └─→ REJECTED (reject)
             ↓
         ┌─────────┐
         │ DRAFT   │ (edit and resubmit)
         └─────────┘
```

## State Definitions

### DRAFT

**Description**: Initial state for newly created artifacts (AI-generated or manually created).

**Characteristics**:
- Artifact is editable
- Artifact is not visible to students
- Artifact cannot be used for downstream generation
- No version number assigned yet
- Can be deleted

**Use Cases**:
- AI agent generates artifact
- Teacher creates artifact manually
- Artifact is regenerated (new version)

### UNDER_REVIEW

**Description**: Artifact submitted for human review.

**Characteristics**:
- Artifact is locked from editing (except by reviewer)
- Artifact is not visible to students
- Artifact cannot be used for downstream generation
- Reviewer can approve or reject
- Reviewer can request changes (returns to DRAFT)

**Use Cases**:
- Teacher submits artifact for approval
- School administrator reviews artifact
- Peer review process

### APPROVED

**Description**: Artifact approved and becomes official.

**Characteristics**:
- Artifact is immutable (cannot be edited directly)
- Artifact is visible to appropriate users
- Artifact can be used for downstream generation
- Version number assigned
- Approval metadata recorded
- Can be regenerated (creates new version)

**Use Cases**:
- Teacher approves own artifact
- School administrator approves artifact
- Artifact approved after review cycle

### REJECTED

**Description**: Artifact rejected during review.

**Characteristics**:
- Artifact is editable
- Artifact is not visible to students
- Artifact cannot be used for downstream generation
- Rejection reason recorded
- Can be resubmitted after changes

**Use Cases**:
- Reviewer rejects artifact due to quality issues
- Reviewer rejects artifact due to misalignment
- Reviewer rejects artifact due to incomplete content

### ARCHIVED

**Description**: Historical artifact no longer in active use.

**Characteristics**:
- Artifact is read-only
- Artifact preserved for historical reporting
- Artifact cannot be regenerated
- Artifact cannot be used for new downstream generation
- Can be viewed for reference

**Use Cases**:
- Artifact superseded by new version
- Academic year completed
- School curriculum changed

## State Summary Table

| State | Editable | Visible to Students | Used for Downstream | Immutable |
|-------|----------|---------------------|---------------------|-----------|
| DRAFT | Yes | No | No | No |
| UNDER_REVIEW | Limited | No | No | No |
| APPROVED | No | Yes | Yes | Yes |
| REJECTED | Yes | No | No | No |
| ARCHIVED | No | No | No | Yes |

---

# SECTION 4 — State Transition Rules

## Allowed Transitions

### DRAFT → UNDER_REVIEW

**Trigger**: User submits artifact for review

**Preconditions**:
- Artifact is in DRAFT status
- Artifact has required fields populated
- User has permission to submit for review

**Postconditions**:
- Artifact status changes to UNDER_REVIEW
- Workflow history entry created
- Notification sent to reviewer

**Example**:
```json
{
  "from_state": "DRAFT",
  "to_state": "UNDER_REVIEW",
  "trigger": "SUBMIT_FOR_REVIEW",
  "performed_by": "user_uuid",
  "comments": "Submitted for school admin review"
}
```

### UNDER_REVIEW → APPROVED

**Trigger**: Reviewer approves artifact

**Preconditions**:
- Artifact is in UNDER_REVIEW status
- Reviewer has permission to approve
- Reviewer has reviewed artifact content

**Postconditions**:
- Artifact status changes to APPROVED
- approved_by and approved_at populated
- Version number assigned
- Workflow history entry created
- Notification sent to submitter
- Artifact available for downstream generation

**Example**:
```json
{
  "from_state": "UNDER_REVIEW",
  "to_state": "APPROVED",
  "trigger": "APPROVE",
  "performed_by": "user_uuid",
  "comments": "Approved - meets all requirements"
}
```

### UNDER_REVIEW → REJECTED

**Trigger**: Reviewer rejects artifact

**Preconditions**:
- Artifact is in UNDER_REVIEW status
- Reviewer has permission to reject
- Rejection reason provided

**Postconditions**:
- Artifact status changes to REJECTED
- rejected_by, rejected_at, rejection_reason populated
- Workflow history entry created
- Notification sent to submitter with rejection reason

**Example**:
```json
{
  "from_state": "UNDER_REVIEW",
  "to_state": "REJECTED",
  "trigger": "REJECT",
  "performed_by": "user_uuid",
  "comments": "Rejected - learning objectives not aligned with CP"
}
```

### REJECTED → DRAFT

**Trigger**: Submitter makes changes and resubmits

**Preconditions**:
- Artifact is in REJECTED status
- User has permission to edit
- Changes made to address rejection reason

**Postconditions**:
- Artifact status changes to DRAFT
- Rejection metadata cleared
- Workflow history entry created

**Example**:
```json
{
  "from_state": "REJECTED",
  "to_state": "DRAFT",
  "trigger": "RESUBMIT",
  "performed_by": "user_uuid",
  "comments": "Addressed rejection comments - resubmitting"
}
```

### APPROVED → ARCHIVED

**Trigger**: System or user archives artifact

**Preconditions**:
- Artifact is in APPROVED status
- No active dependencies (or dependencies handled)
- Archival policy conditions met

**Postconditions**:
- Artifact status changes to ARCHIVED
- Workflow history entry created
- Artifact becomes read-only

**Example**:
```json
{
  "from_state": "APPROVED",
  "to_state": "ARCHIVED",
  "trigger": "ARCHIVE",
  "performed_by": "system",
  "comments": "Auto-archived - academic year completed"
}
```

## Prohibited Transitions

The following transitions are NOT allowed:

| From | To | Reason |
|------|-----|--------|
| DRAFT | APPROVED | Must go through review first |
| DRAFT | REJECTED | Cannot reject without review |
| DRAFT | ARCHIVED | Cannot archive draft |
| UNDER_REVIEW | DRAFT | Must use explicit "request changes" action |
| UNDER_REVIEW | ARCHIVED | Cannot archive while under review |
| APPROVED | DRAFT | Must regenerate to create new version |
| APPROVED | REJECTED | Cannot reject approved artifact |
| REJECTED | APPROVED | Must resubmit as DRAFT first |
| REJECTED | ARCHIVED | Cannot archive rejected artifact |
| ARCHIVED | Any | Terminal state, no outgoing transitions |

## State Transition Diagram

```
┌─────────┐
│  DRAFT  │
└────┬────┘
     │ Submit for Review
     ↓
┌─────────────┐
│UNDER_REVIEW │
└──────┬──────┘
       │
       ├─→ Approve
       │     ↓
       │  ┌─────────┐
       │  │ APPROVED│
       │  └────┬────┘
       │       │ Archive
       │       ↓
       │  ┌─────────┐
       │  │ARCHIVED │
       │  └─────────┘
       │
       └─→ Reject
             ↓
       ┌─────────┐
       │ REJECTED│
       └────┬────┘
            │ Resubmit
            ↓
       ┌─────────┐
       │  DRAFT  │
       └─────────┘
```

---

# SECTION 5 — Workflow History

## Workflow History Model

### workflow_history Table

**Purpose**: Track all workflow state transitions for audit and compliance.

**Schema**:

```sql
CREATE TABLE workflow_history (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    entity_type VARCHAR(50) NOT NULL,
    entity_id UUID NOT NULL,
    from_state VARCHAR(20) NOT NULL,
    to_state VARCHAR(20) NOT NULL,
    action VARCHAR(50) NOT NULL,
    performed_by UUID NOT NULL,
    performed_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    comments TEXT,
    metadata JSONB,
    
    -- Foreign Keys
    FOREIGN KEY (performed_by) REFERENCES users(id)
);

CREATE INDEX idx_workflow_history_entity ON workflow_history(entity_type, entity_id);
CREATE INDEX idx_workflow_history_performed_by ON workflow_history(performed_by);
CREATE INDEX idx_workflow_history_performed_at ON workflow_history(performed_at);
CREATE INDEX idx_workflow_history_action ON workflow_history(action);
```

### Field Descriptions

| Field | Type | Description |
|-------|------|-------------|
| `id` | UUID | Unique history entry identifier |
| `entity_type` | VARCHAR(50) | Type of entity (TP, ATP, MODUL_AJAR, ASSESSMENT, RUBRIC, NARRATIVE_REPORT) |
| `entity_id` | UUID | ID of the entity that changed state |
| `from_state` | VARCHAR(20) | Previous state |
| `to_state` | VARCHAR(20) | New state |
| `action` | VARCHAR(50) | Action that triggered transition (SUBMIT_FOR_REVIEW, APPROVE, REJECT, etc.) |
| `performed_by` | UUID | User who performed the action |
| `performed_at` | TIMESTAMP | When the action was performed |
| `comments` | TEXT | Optional comments about the transition |
| `metadata` | JSONB | Additional context-specific metadata |

### Entity Type Values

| Entity Type | Description |
|-------------|-------------|
| TP | Teaching Plan |
| ATP | Alur Tujuan Pembelajaran |
| MODUL_AJAR | Modul Ajar |
| ASSESSMENT | Assessment |
| RUBRIC | Rubric |
| NARRATIVE_REPORT | Narrative Report |

### Action Values

| Action | Description |
|--------|-------------|
| CREATE | Entity created (initial DRAFT) |
| SUBMIT_FOR_REVIEW | Submitted for review |
| REQUEST_CHANGES | Reviewer requested changes |
| APPROVE | Approved by reviewer |
| REJECT | Rejected by reviewer |
| RESUBMIT | Resubmitted after rejection |
| REGENERATE | Regenerated (new version) |
| ARCHIVE | Archived |
| RESTORE | Restored from archive (if supported) |

## History Query Patterns

### Get Entity History

```sql
SELECT * FROM workflow_history
WHERE entity_type = 'TP'
  AND entity_id = :entity_id
ORDER BY performed_at DESC;
```

### Get User Actions

```sql
SELECT * FROM workflow_history
WHERE performed_by = :user_id
ORDER BY performed_at DESC
LIMIT 100;
```

### Get State Transition Statistics

```sql
SELECT 
    entity_type,
    action,
    COUNT(*) as count
FROM workflow_history
WHERE performed_at >= :start_date
GROUP BY entity_type, action
ORDER BY count DESC;
```

---

# SECTION 6 — Approval Metadata

## Approval Fields

All artifact tables must include the following approval metadata fields:

### Standard Approval Metadata

| Field | Type | Nullable | Description |
|-------|------|----------|-------------|
| `approved_by` | UUID | Yes | User who approved the artifact |
| `approved_at` | TIMESTAMP WITH TIME ZONE | Yes | When approval occurred |
| `rejected_by` | UUID | Yes | User who rejected the artifact |
| `rejected_at` | TIMESTAMP WITH TIME ZONE | Yes | When rejection occurred |
| `rejection_reason` | TEXT | Yes | Reason for rejection |

### Field Population Rules

**approved_by**:
- Populated when artifact transitions to APPROVED
- Cleared if artifact is rejected and resubmitted
- References users(id)

**approved_at**:
- Populated when artifact transitions to APPROVED
- Set to current timestamp
- Cleared if artifact is rejected and resubmitted

**rejected_by**:
- Populated when artifact transitions to REJECTED
- Cleared if artifact is resubmitted
- References users(id)

**rejected_at**:
- Populated when artifact transitions to REJECTED
- Set to current timestamp
- Cleared if artifact is resubmitted

**rejection_reason**:
- Populated when artifact transitions to REJECTED
- Required field for rejection
- Cleared if artifact is resubmitted

## Example Schema Addition

```sql
-- Add to any artifact table
ALTER TABLE tp ADD COLUMN approved_by UUID;
ALTER TABLE tp ADD COLUMN approved_at TIMESTAMP WITH TIME ZONE;
ALTER TABLE tp ADD COLUMN rejected_by UUID;
ALTER TABLE tp ADD COLUMN rejected_at TIMESTAMP WITH TIME ZONE;
ALTER TABLE tp ADD COLUMN rejection_reason TEXT;

ALTER TABLE tp ADD CONSTRAINT fk_tp_approved_by FOREIGN KEY (approved_by) REFERENCES users(id);
ALTER TABLE tp ADD CONSTRAINT fk_tp_rejected_by FOREIGN KEY (rejected_by) REFERENCES users(id);

CREATE INDEX idx_tp_approved_by ON tp(approved_by);
CREATE INDEX idx_tp_rejected_by ON tp(rejected_by);
```

---

# SECTION 7 — Audit Requirements

## Audit Principle

**Rule**: Every workflow action must be auditable.

**Implications**:
- All state transitions logged in workflow_history
- All user actions tracked with timestamp
- No actions can be performed without audit trail
- Audit logs are immutable
- Audit logs are retained per retention policy

## Audit Data Requirements

### Mandatory Audit Fields

Every workflow action must record:

1. **Who**: User ID of performer
2. **What**: Action performed
3. **When**: Timestamp of action
4. **Where**: Entity type and ID
5. **Why**: Comments or reason (when applicable)
6. **From/To**: State transition

### Optional Audit Fields

1. **Metadata**: Additional context (JSONB)
2. **IP Address**: Network origin (if applicable)
3. **User Agent**: Client application (if applicable)

## Audit Retention Policy

**Retention Period**: 7 years minimum

**Rationale**: Educational records may be needed for long-term compliance, accreditation, and legal requirements.

**Implementation**:
- Audit logs never deleted within retention period
- After retention period, logs may be archived to cold storage
- Archived logs remain accessible for compliance audits

## Audit Query Capabilities

### Compliance Audit

**Query**: All approvals by a specific user within a time range

```sql
SELECT 
    wh.entity_type,
    wh.entity_id,
    wh.action,
    wh.performed_at,
    wh.comments,
    u.name as performer_name
FROM workflow_history wh
JOIN users u ON wh.performed_by = u.id
WHERE wh.action = 'APPROVE'
  AND wh.performed_by = :user_id
  AND wh.performed_at BETWEEN :start_date AND :end_date
ORDER BY wh.performed_at DESC;
```

### Artifact Lifecycle Audit

**Query**: Complete lifecycle of a specific artifact

```sql
SELECT 
    wh.from_state,
    wh.to_state,
    wh.action,
    wh.performed_at,
    wh.comments,
    u.name as performer_name
FROM workflow_history wh
JOIN users u ON wh.performed_by = u.id
WHERE wh.entity_type = :entity_type
  AND wh.entity_id = :entity_id
ORDER BY wh.performed_at ASC;
```

### Rejection Analysis

**Query**: All rejections with reasons

```sql
SELECT 
    wh.entity_type,
    wh.rejection_reason,
    COUNT(*) as count
FROM workflow_history wh
JOIN (
    SELECT entity_id, MAX(performed_at) as last_rejected_at
    FROM workflow_history
    WHERE action = 'REJECT'
    GROUP BY entity_id
) last_rejections ON wh.entity_id = last_rejections.entity_id 
    AND wh.performed_at = last_rejections.last_rejected_at
WHERE wh.action = 'REJECT'
GROUP BY wh.entity_type, wh.rejection_reason
ORDER BY count DESC;
```

---

# SECTION 8 — Notification Triggers

## Notification Events

### Draft Submitted

**Trigger**: Artifact transitions from DRAFT to UNDER_REVIEW

**Recipients**:
- Reviewer (if specified)
- School administrator (if no reviewer specified)

**Notification Content**:
- Artifact type and ID
- Submitter name
- Submission timestamp
- Link to review screen

**Channels**: In-app notification, email (optional)

### Approved

**Trigger**: Artifact transitions from UNDER_REVIEW to APPROVED

**Recipients**:
- Submitter
- School administrator (if different from reviewer)

**Notification Content**:
- Artifact type and ID
- Approver name
- Approval timestamp
- Link to view approved artifact

**Channels**: In-app notification, email (optional)

### Rejected

**Trigger**: Artifact transitions from UNDER_REVIEW to REJECTED

**Recipients**:
- Submitter

**Notification Content**:
- Artifact type and ID
- Rejecter name
- Rejection timestamp
- Rejection reason
- Link to edit and resubmit

**Channels**: In-app notification, email (optional)

### Regenerated

**Trigger**: Artifact is regenerated (new version created)

**Recipients**:
- User who requested regeneration

**Notification Content**:
- Artifact type and ID
- New version number
- Generation timestamp
- Link to review new version

**Channels**: In-app notification

### Archived

**Trigger**: Artifact transitions from APPROVED to ARCHIVED

**Recipients**:
- Owner of artifact
- School administrator

**Notification Content**:
- Artifact type and ID
- Archival timestamp
- Reason for archival
- Link to view archived artifact

**Channels**: In-app notification

## Notification Model

### notifications Table

```sql
CREATE TABLE notifications (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    notification_type VARCHAR(50) NOT NULL,
    title VARCHAR(255) NOT NULL,
    message TEXT NOT NULL,
    entity_type VARCHAR(50),
    entity_id UUID,
    action_link VARCHAR(500),
    is_read BOOLEAN NOT NULL DEFAULT false,
    read_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE INDEX idx_notifications_user_id ON notifications(user_id);
CREATE INDEX idx_notifications_is_read ON notifications(is_read);
CREATE INDEX idx_notifications_created_at ON notifications(created_at);
CREATE INDEX idx_notifications_type ON notifications(notification_type);
```

### Notification Types

| Type | Description |
|------|-------------|
| DRAFT_SUBMITTED | Artifact submitted for review |
| APPROVED | Artifact approved |
| REJECTED | Artifact rejected |
| REGENERATED | Artifact regenerated |
| ARCHIVED | Artifact archived |
| REVIEW_REQUESTED | Review requested from user |

---

# SECTION 9 — Reusable Workflow Engine

## Design Principles

### Artifact-Agnostic

**Rule**: Workflow engine must not contain artifact-specific logic.

**Implications**:
- Engine works with generic entity_type and entity_id
- No hardcoded business rules for specific artifact types
- Artifact-specific validation handled by calling services
- Engine focuses on state management and transitions

### Service-Oriented

**Rule**: Workflow engine is a service module within the modular monolith.

**Implications**:
- Clear interface (API) for workflow operations
- Dependency injection for artifact services
- Testable in isolation
- Can be extended without modifying core logic

## Workflow Engine Interface

### Core Operations

```typescript
interface IWorkflowEngine {
  // Initialize workflow for new entity
  initializeWorkflow(
    entity_type: string,
    entity_id: string,
    initial_state: string = 'DRAFT',
    performed_by: string
  ): Promise<WorkflowHistoryEntry>;

  // Transition entity to new state
  transitionState(
    entity_type: string,
    entity_id: string,
    to_state: string,
    action: string,
    performed_by: string,
    comments?: string,
    metadata?: Record<string, any>
  ): Promise<WorkflowHistoryEntry>;

  // Get current state of entity
  getCurrentState(
    entity_type: string,
    entity_id: string
  ): Promise<string>;

  // Get workflow history for entity
  getHistory(
    entity_type: string,
    entity_id: string
  ): Promise<WorkflowHistoryEntry[]>;

  // Check if transition is valid
  isValidTransition(
    from_state: string,
    to_state: string
  ): boolean;

  // Get allowed transitions from current state
  getAllowedTransitions(
    current_state: string
  ): string[];
}
```

## State Transition Configuration

### Transition Rules Table

```sql
CREATE TABLE workflow_transition_rules (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    from_state VARCHAR(20) NOT NULL,
    to_state VARCHAR(20) NOT NULL,
    action VARCHAR(50) NOT NULL,
    is_allowed BOOLEAN NOT NULL DEFAULT true,
    requires_permission VARCHAR(50),
    validation_hook VARCHAR(100),
    notification_trigger VARCHAR(50),
    
    UNIQUE(from_state, to_state, action)
);

CREATE INDEX idx_transition_rules_from_state ON workflow_transition_rules(from_state);
CREATE INDEX idx_transition_rules_to_state ON workflow_transition_rules(to_state);
```

### Seed Data

```sql
INSERT INTO workflow_transition_rules (from_state, to_state, action, is_allowed, notification_trigger) VALUES
('DRAFT', 'UNDER_REVIEW', 'SUBMIT_FOR_REVIEW', true, 'DRAFT_SUBMITTED'),
('UNDER_REVIEW', 'APPROVED', 'APPROVE', true, 'APPROVED'),
('UNDER_REVIEW', 'REJECTED', 'REJECT', true, 'REJECTED'),
('REJECTED', 'DRAFT', 'RESUBMIT', true, null),
('APPROVED', 'ARCHIVED', 'ARCHIVE', true, 'ARCHIVED');
```

## Workflow Engine Implementation

### Transition Validation

```typescript
async transitionState(
  entity_type: string,
  entity_id: string,
  to_state: string,
  action: string,
  performed_by: string,
  comments?: string,
  metadata?: Record<string, any>
): Promise<WorkflowHistoryEntry> {
  // 1. Get current state
  const current_state = await this.getCurrentState(entity_type, entity_id);
  
  // 2. Validate transition
  if (!this.isValidTransition(current_state, to_state)) {
    throw new WorkflowTransitionError(
      `Invalid transition from ${current_state} to ${to_state}`
    );
  }
  
  // 3. Check permissions
  await this.checkPermission(performed_by, action);
  
  // 4. Run validation hooks if configured
  await this.runValidationHooks(entity_type, entity_id, to_state);
  
  // 5. Update entity state
  await this.updateEntityState(entity_type, entity_id, to_state);
  
  // 6. Update approval metadata if needed
  await this.updateApprovalMetadata(entity_type, entity_id, to_state, performed_by, comments);
  
  // 7. Log workflow history
  const history_entry = await this.logHistory(
    entity_type,
    entity_id,
    current_state,
    to_state,
    action,
    performed_by,
    comments,
    metadata
  );
  
  // 8. Trigger notifications
  await this.triggerNotifications(entity_type, entity_id, to_state, performed_by);
  
  return history_entry;
}
```

### Validation Hooks

**Purpose**: Allow artifact-specific validation without hardcoding in workflow engine.

**Implementation**:
- Validation hooks registered per entity_type and transition
- Hooks called before state transition
- Hooks can throw errors to block transition

**Example Hook Registration**:

```typescript
workflowEngine.registerValidationHook(
  'TP',
  'UNDER_REVIEW',
  'APPROVE',
  async (entity_id: string) => {
    const tp = await tpService.getById(entity_id);
    if (!tp.learning_objectives || tp.learning_objectives.length === 0) {
      throw new ValidationError('TP must have at least one learning objective');
    }
  }
);
```

## Integration with Artifact Services

### Artifact Service Responsibilities

Artifact services (TPService, ATPService, etc.) are responsible for:

1. **Business Logic Validation**: Artifact-specific validation rules
2. **State Persistence**: Updating the artifact's status field
3. **Side Effects**: Triggering downstream operations when approved
4. **Version Management**: Handling version-specific logic

### Workflow Engine Responsibilities

Workflow engine is responsible for:

1. **State Transition Validation**: Ensuring valid transitions
2. **Permission Checking**: Verifying user permissions
3. **Audit Logging**: Recording all transitions
4. **Notification Triggering**: Sending appropriate notifications
5. **Approval Metadata**: Managing approval/rejection fields

### Collaboration Pattern

```typescript
// In TPService
async approveTP(tp_id: string, approved_by: string, comments?: string) {
  // 1. Business logic validation
  const tp = await this.getById(tp_id);
  this.validateTPForApproval(tp);
  
  // 2. Call workflow engine for state transition
  await workflowEngine.transitionState(
    'TP',
    tp_id,
    'APPROVED',
    'APPROVE',
    approved_by,
    comments
  );
  
  // 3. Handle side effects (e.g., version assignment)
  await this.assignVersionNumber(tp_id);
  
  // 4. Trigger downstream availability
  await this.notifyDownstreamServices(tp_id);
}
```

---

# SECTION 10 — Module Architecture

## Workflow Module Structure

### Module Components

```
workflow/
├── core/
│   ├── WorkflowEngine.ts          # Core workflow engine
│   ├── StateTransitionValidator.ts
│   ├── PermissionChecker.ts
│   └── AuditLogger.ts
├── config/
│   ├── TransitionRules.ts         # Transition configuration
│   └── ValidationHooks.ts        # Hook registry
├── notifications/
│   ├── NotificationService.ts
│   └── NotificationTemplates.ts
├── repository/
│   ├── WorkflowHistoryRepository.ts
│   └── TransitionRulesRepository.ts
└── api/
    └── WorkflowController.ts
```

### Dependencies

**Workflow Module Depends On**:
- User Module (for user validation)
- Notification Module (for sending notifications)

**Artifact Modules Depend On**:
- Workflow Module (for state transitions)

**No Circular Dependencies**: Workflow module does not depend on artifact modules.

---

# SECTION 11 — API Contract

## Workflow Endpoints

### Get Current State

```
GET /api/v1/workflow/state/{entity_type}/{entity_id}
```

**Response**:
```json
{
  "success": true,
  "data": {
    "entity_type": "TP",
    "entity_id": "uuid",
    "current_state": "UNDER_REVIEW",
    "allowed_transitions": ["APPROVE", "REJECT"]
  }
}
```

### Transition State

```
POST /api/v1/workflow/transition
```

**Request**:
```json
{
  "entity_type": "TP",
  "entity_id": "uuid",
  "to_state": "APPROVED",
  "action": "APPROVE",
  "comments": "Approved - meets requirements"
}
```

**Response**:
```json
{
  "success": true,
  "data": {
    "history_entry": {
      "id": "uuid",
      "entity_type": "TP",
      "entity_id": "uuid",
      "from_state": "UNDER_REVIEW",
      "to_state": "APPROVED",
      "action": "APPROVE",
      "performed_by": "uuid",
      "performed_at": "2026-06-05T12:00:00Z",
      "comments": "Approved - meets requirements"
    }
  }
}
```

### Get Workflow History

```
GET /api/v1/workflow/history/{entity_type}/{entity_id}
```

**Response**:
```json
{
  "success": true,
  "data": {
    "history": [
      {
        "id": "uuid",
        "from_state": "DRAFT",
        "to_state": "UNDER_REVIEW",
        "action": "SUBMIT_FOR_REVIEW",
        "performed_by": "uuid",
        "performed_at": "2026-06-05T10:00:00Z",
        "comments": "Submitted for review"
      },
      {
        "id": "uuid",
        "from_state": "UNDER_REVIEW",
        "to_state": "APPROVED",
        "action": "APPROVE",
        "performed_by": "uuid",
        "performed_at": "2026-06-05T12:00:00Z",
        "comments": "Approved - meets requirements"
      }
    ]
  }
}
```

---

# SECTION 12 — Security Considerations

## Permission Requirements

### State Transition Permissions

| Action | Required Permission |
|--------|---------------------|
| SUBMIT_FOR_REVIEW | artifact:submit |
| APPROVE | artifact:approve |
| REJECT | artifact:reject |
| RESUBMIT | artifact:edit |
| ARCHIVE | artifact:archive |
| VIEW_HISTORY | artifact:view |

### Role-Based Access

| Role | Permissions |
|------|-------------|
| Teacher | artifact:submit, artifact:edit, artifact:view for own artifacts |
| School Admin | artifact:approve, artifact:reject, artifact:archive for school artifacts |
| System Admin | All permissions |

## Audit Security

**Immutable Logs**: Workflow history entries cannot be modified or deleted

**Access Control**: Audit logs accessible only to authorized users

**Tamper Detection**: Cryptographic hashing of audit entries (optional, future enhancement)

---

# SECTION 13 — Summary

## Key Deliverables

✓ **Workflow State Diagram**: Visual representation of states and transitions
✓ **Workflow Rules**: Complete state transition rules with preconditions and postconditions
✓ **Audit Rules**: Comprehensive audit requirements with 7-year retention
✓ **Notification Rules**: Notification triggers for all workflow events
✓ **Workflow Engine Design**: Reusable, artifact-agnostic workflow service

## Design Principles Adhered To

- ✓ Human approval mandatory
- ✓ AI never publishes official artifacts
- ✓ Consistent workflow states across all artifact types
- ✓ Complete audit trail for all actions
- ✓ Reusable workflow engine (no artifact-specific logic)
- ✓ Modular monolith architecture
- ✓ Clear separation between workflow engine and artifact services

## Applicable Artifacts

The workflow architecture applies to:
- TP (Teaching Plan)
- ATP (Alur Tujuan Pembelajaran)
- Modul Ajar
- Assessment
- Rubric
- Narrative Report

## Next Steps

1. Implement workflow_history table
2. Implement workflow_transition_rules table
3. Implement notifications table
4. Add approval metadata fields to all artifact tables
5. Implement WorkflowEngine service
6. Implement validation hook registry
7. Implement notification service integration
8. Create workflow API endpoints
9. Implement permission checks
10. Add workflow UI components for artifact review screens

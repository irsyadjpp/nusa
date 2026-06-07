# 31_WORKFLOW_ENGINE_DESIGN.md

## Foundation Document for NUSA Education Platform

**Version**: 1.0
**Date**: June 2026
**Status**: FOUNDATION DOCUMENT
**Alignment**: Aligned with 14_DATABASE_SCHEMA.md, 25_WORKFLOW_ARCHITECTURE.md, 27_BACKEND_FOUNDATION_DESIGN.md, 28_AUTHENTICATION_DESIGN.md

**Purpose**: Design a reusable workflow engine for NUSA. This document defines the workflow service design, state machine, history tracking, approval process, notification triggers, audit design, and database mapping. The workflow engine is reusable across all artifacts (TP, ATP, Modul Ajar, Assessment, Rubric, Narrative Report).

---

# SECTION 1 — Executive Summary

## Why Workflow Engine Design Matters

A reusable workflow engine ensures:
- Consistent approval processes across all artifact types
- Standardized state transitions and validation
- Complete audit trail for compliance
- Reduced code duplication
- Easy addition of new artifact types
- Centralized workflow logic

## Supported Artifacts

**Current (MVP)**:
- TP (Teaching Plan)

**Future (Wave 2)**:
- ATP (Annual Teaching Plan)
- Modul Ajar (Lesson Plan)
- Assessment
- Rubric
- Narrative Report

## Core Principles

- **Reusable**: Single workflow engine for all artifact types
- **State Machine**: Standardized states and transitions
- **Audit Trail**: Complete history of all workflow actions
- **Permission-Based**: Role-based access control
- **Event-Driven**: Notification triggers on state changes
- **MVP-Focused**: Simple synchronous workflow for MVP

---

# SECTION 2 — Workflow Service Design

## Service Architecture

### Workflow Engine Components

```
┌─────────────────────────────────────────────────────────────┐
│                    Workflow Engine                           │
└─────────────────────────────────────────────────────────────┘
                            │
        ┌───────────────────┼───────────────────┐
        │                   │                   │
        ↓                   ↓                   ↓
┌───────────────┐  ┌───────────────┐  ┌───────────────┐
│ State Machine │  │   History     │  │ Notification  │
│               │  │   Service     │  │   Service     │
└───────────────┘  └───────────────┘  └───────────────┘
```

### Workflow Service Interface

```go
package workflow

import "context"

type ArtifactType string

const (
    ArtifactTypeTP           ArtifactType = "tp_set"
    ArtifactTypeATP          ArtifactType = "atp_set"
    ArtifactTypeModulAjar    ArtifactType = "modul_ajar_set"
    ArtifactTypeAssessment   ArtifactType = "assessment"
    ArtifactTypeRubric       ArtifactType = "rubric"
    ArtifactTypeNarrative    ArtifactType = "narrative_report"
)

type State string

const (
    StateDraft       State = "DRAFT"
    StateUnderReview State = "UNDER_REVIEW"
    StateApproved    State = "APPROVED"
    StateRejected    State = "REJECTED"
    StateArchived    State = "ARCHIVED"
)

type Action string

const (
    ActionCreate      Action = "CREATE"
    ActionUpdate      Action = "UPDATE"
    ActionSubmit     Action = "SUBMIT"
    ActionApprove    Action = "APPROVE"
    ActionReject     Action = "REJECT"
    ActionArchive    Action = "ARCHIVE"
    ActionRegenerate Action = "REGENERATE"
)

type TransitionRequest struct {
    ArtifactID   string
    ArtifactType ArtifactType
    CurrentState State
    TargetState State
    Action      Action
    UserID      string
    Reason      string
    Metadata    map[string]interface{}
}

type TransitionResult struct {
    Success     bool
    NewState    State
    HistoryID   string
    Error       string
}

type Service interface {
    // State Transitions
    Transition(ctx context.Context, req *TransitionRequest) (*TransitionResult, error)
    CanTransition(ctx context.Context, artifactType ArtifactType, from, to State) bool
    
    // History
    GetHistory(ctx context.Context, artifactID string, artifactType ArtifactType) ([]*History, error)
    GetHistoryByUser(ctx context.Context, userID string, artifactType ArtifactType) ([]*History, error)
    
    // State Queries
    GetCurrentState(ctx context.Context, artifactID string, artifactType ArtifactType) (State, error)
    GetArtifactsByState(ctx context.Context, artifactType ArtifactType, state State) ([]string, error)
    
    // Approval
    SubmitForApproval(ctx context.Context, artifactID string, artifactType ArtifactType, userID string, reason string) error
    Approve(ctx context.Context, artifactID string, artifactType ArtifactType, userID string, reason string) error
    Reject(ctx context.Context, artifactID string, artifactType ArtifactType, userID string, reason string) error
    Archive(ctx context.Context, artifactID string, artifactType ArtifactType, userID string) error
    
    // Validation
    ValidateTransition(ctx context.Context, req *TransitionRequest) error
}
```

## Service Implementation

### Workflow Service

```go
package workflow

import (
    "context"
    "errors"
    "fmt"
    
    "github.com/google/uuid"
)

type service struct {
    stateMachine *StateMachine
    historyRepo  HistoryRepository
    notifier    NotificationService
}

func NewService(historyRepo HistoryRepository, notifier NotificationService) Service {
    return &service{
        stateMachine: NewStateMachine(),
        historyRepo:  historyRepo,
        notifier:    notifier,
    }
}

func (s *service) Transition(ctx context.Context, req *TransitionRequest) (*TransitionResult, error) {
    // Validate transition
    if err := s.ValidateTransition(ctx, req); err != nil {
        return &TransitionResult{
            Success: false,
            Error:   err.Error(),
        }, err
    }
    
    // Check if transition is allowed
    if !s.stateMachine.CanTransition(req.ArtifactType, req.CurrentState, req.TargetState) {
        return &TransitionResult{
            Success: false,
            Error:   fmt.Sprintf("invalid transition from %s to %s", req.CurrentState, req.TargetState),
        }, ErrInvalidTransition
    }
    
    // Execute transition (update artifact status)
    if err := s.updateArtifactStatus(ctx, req); err != nil {
        return &TransitionResult{
            Success: false,
            Error:   err.Error(),
        }, err
    }
    
    // Record history
    history := &History{
        ID:          uuid.New().String(),
        ArtifactID:  req.ArtifactID,
        ArtifactType: req.ArtifactType,
        Action:      req.Action,
        FromState:   req.CurrentState,
        ToState:     req.TargetState,
        UserID:      req.UserID,
        Reason:      req.Reason,
        Metadata:    req.Metadata,
        CreatedAt:   time.Now(),
    }
    
    if err := s.historyRepo.CreateHistory(ctx, history); err != nil {
        return &TransitionResult{
            Success: false,
            Error:   err.Error(),
        }, err
    }
    
    // Trigger notifications
    if err := s.notifier.Notify(ctx, req); err != nil {
        // Log error but don't fail transition
        log.Printf("notification failed: %v", err)
    }
    
    return &TransitionResult{
        Success:   true,
        NewState:  req.TargetState,
        HistoryID: history.ID,
    }, nil
}

func (s *service) ValidateTransition(ctx context.Context, req *TransitionRequest) error {
    // Validate artifact ID
    if req.ArtifactID == "" {
        return errors.New("artifact ID is required")
    }
    
    // Validate artifact type
    if !s.isValidArtifactType(req.ArtifactType) {
        return errors.New("invalid artifact type")
    }
    
    // Validate states
    if req.CurrentState == "" {
        return errors.New("current state is required")
    }
    
    if req.TargetState == "" {
        return errors.New("target state is required")
    }
    
    // Validate user ID
    if req.UserID == "" {
        return errors.New("user ID is required")
    }
    
    return nil
}

func (s *service) updateArtifactStatus(ctx context.Context, req *TransitionRequest) error {
    // This is implemented by calling the specific artifact service
    // The workflow engine doesn't directly update artifact tables
    // Instead, it delegates to the artifact-specific service
    return nil
}

func (s *service) SubmitForApproval(ctx context.Context, artifactID string, artifactType ArtifactType, userID string, reason string) error {
    currentState, err := s.GetCurrentState(ctx, artifactID, artifactType)
    if err != nil {
        return err
    }
    
    req := &TransitionRequest{
        ArtifactID:   artifactID,
        ArtifactType: artifactType,
        CurrentState: currentState,
        TargetState:  StateUnderReview,
        Action:       ActionSubmit,
        UserID:       userID,
        Reason:       reason,
    }
    
    result, err := s.Transition(ctx, req)
    if err != nil {
        return err
    }
    
    if !result.Success {
        return errors.New(result.Error)
    }
    
    return nil
}

func (s *service) Approve(ctx context.Context, artifactID string, artifactType ArtifactType, userID string, reason string) error {
    currentState, err := s.GetCurrentState(ctx, artifactID, artifactType)
    if err != nil {
        return err
    }
    
    req := &TransitionRequest{
        ArtifactID:   artifactID,
        ArtifactType: artifactType,
        CurrentState: currentState,
        TargetState:  StateApproved,
        Action:       ActionApprove,
        UserID:       userID,
        Reason:       reason,
    }
    
    result, err := s.Transition(ctx, req)
    if err != nil {
        return err
    }
    
    if !result.Success {
        return errors.New(result.Error)
    }
    
    return nil
}

func (s *service) Reject(ctx context.Context, artifactID string, artifactType ArtifactType, userID string, reason string) error {
    currentState, err := s.GetCurrentState(ctx, artifactID, artifactType)
    if err != nil {
        return err
    }
    
    req := &TransitionRequest{
        ArtifactID:   artifactID,
        ArtifactType: artifactType,
        CurrentState: currentState,
        TargetState:  StateRejected,
        Action:       ActionReject,
        UserID:       userID,
        Reason:       reason,
    }
    
    result, err := s.Transition(ctx, req)
    if err != nil {
        return err
    }
    
    if !result.Success {
        return errors.New(result.Error)
    }
    
    return nil
}

func (s *service) Archive(ctx context.Context, artifactID string, artifactType ArtifactType, userID string) error {
    currentState, err := s.GetCurrentState(ctx, artifactID, artifactType)
    if err != nil {
        return err
    }
    
    req := &TransitionRequest{
        ArtifactID:   artifactID,
        ArtifactType: artifactType,
        CurrentState: currentState,
        TargetState:  StateArchived,
        Action:       ActionArchive,
        UserID:       userID,
        Reason:       "Archived",
    }
    
    result, err := s.Transition(ctx, req)
    if err != nil {
        return err
    }
    
    if !result.Success {
        return errors.New(result.Error)
    }
    
    return nil
}

func (s *service) isValidArtifactType(at ArtifactType) bool {
    validTypes := []ArtifactType{
        ArtifactTypeTP,
        ArtifactTypeATP,
        ArtifactTypeModulAjar,
        ArtifactTypeAssessment,
        ArtifactTypeRubric,
        ArtifactTypeNarrative,
    }
    
    for _, validType := range validTypes {
        if at == validType {
            return true
        }
    }
    return false
}
```

---

# SECTION 3 — Workflow State Machine

## State Machine Definition

### Standardized States

All artifacts use the same standardized states:

| State | Description | Terminal |
|-------|-------------|----------|
| DRAFT | Initial state for newly created artifacts | No |
| UNDER_REVIEW | Artifact submitted for review | No |
| APPROVED | Artifact approved and official | No |
| REJECTED | Artifact rejected during review | No |
| ARCHIVED | Historical artifact no longer in use | Yes |

### State Transitions

**Allowed Transitions**:

```
DRAFT → UNDER_REVIEW
DRAFT → ARCHIVED
UNDER_REVIEW → APPROVED
UNDER_REVIEW → REJECTED
APPROVED → ARCHIVED
REJECTED → DRAFT
```

**Forbidden Transitions**:

- DRAFT → APPROVED (must go through review)
- DRAFT → REJECTED (must go through review first)
- UNDER_REVIEW → ARCHIVED (must be approved or rejected first)
- UNDER_REVIEW → DRAFT (must be rejected first)
- APPROVED → DRAFT (regeneration creates new version)
- APPROVED → REJECTED (approved artifacts cannot be rejected)
- APPROVED → UNDER_REVIEW (approved artifacts cannot go back to review)
- REJECTED → APPROVED (must go through review again)
- REJECTED → ARCHIVED (must go back to draft first)
- ARCHIVED → Any state (archived is terminal)

### State Machine Implementation

```go
package workflow

type StateMachine struct {
    transitions map[ArtifactType]map[State][]State
}

func NewStateMachine() *StateMachine {
    sm := &StateMachine{
        transitions: make(map[ArtifactType]map[State][]State),
    }
    
    // Initialize standard transitions for all artifact types
    artifactTypes := []ArtifactType{
        ArtifactTypeTP,
        ArtifactTypeATP,
        ArtifactTypeModulAjar,
        ArtifactTypeAssessment,
        ArtifactTypeRubric,
        ArtifactTypeNarrative,
    }
    
    for _, at := range artifactTypes {
        sm.transitions[at] = map[State][]State{
            StateDraft:       {StateUnderReview, StateArchived},
            StateUnderReview: {StateApproved, StateRejected},
            StateApproved:    {StateArchived},
            StateRejected:    {StateDraft},
            StateArchived:    {},
        }
    }
    
    return sm
}

func (sm *StateMachine) CanTransition(at ArtifactType, from, to State) bool {
    transitions, ok := sm.transitions[at]
    if !ok {
        return false
    }
    
    allowedStates, ok := transitions[from]
    if !ok {
        return false
    }
    
    for _, allowedState := range allowedStates {
        if allowedState == to {
            return true
        }
    }
    
    return false
}

func (sm *StateMachine) GetAllowedTransitions(at ArtifactType, from State) []State {
    transitions, ok := sm.transitions[at]
    if !ok {
        return []State{}
    }
    
    allowedStates, ok := transitions[from]
    if !ok {
        return []State{}
    }
    
    return allowedStates
}
```

---

# SECTION 4 — Workflow History

## History Entity

### Workflow History Table

| Field | Type | Constraints | Description |
|-------|------|-------------|-------------|
| `id` | UUID | PRIMARY KEY | Unique history identifier |
| `artifact_id` | UUID | NOT NULL | Artifact identifier |
| `artifact_type` | VARCHAR(50) | NOT NULL | Artifact type (tp_set, atp_set, etc.) |
| `action` | VARCHAR(50) | NOT NULL | Action performed (CREATE, SUBMIT, APPROVE, etc.) |
| `from_state` | VARCHAR(20) | NOT NULL | Previous state |
| `to_state` | VARCHAR(20) | NOT NULL | New state |
| `user_id` | UUID | NOT NULL, FOREIGN KEY → users(id) | User who performed action |
| `reason` | TEXT | NULLABLE | Reason for action |
| `metadata` | JSONB | NULLABLE | Additional metadata |
| `ai_generation_id` | UUID | NULLABLE, FOREIGN KEY → ai_generation_logs(id) | AI generation reference |
| `created_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Action timestamp |

### Constraints

```sql
-- Artifact and type index
CREATE INDEX idx_workflow_history_artifact ON workflow_history(artifact_id, artifact_type);

-- User index
CREATE INDEX idx_workflow_history_user_id ON workflow_history(user_id);

-- Action index
CREATE INDEX idx_workflow_history_action ON workflow_history(action);

-- Created at index
CREATE INDEX idx_workflow_history_created_at ON workflow_history(created_at);

-- AI generation index
CREATE INDEX idx_workflow_history_ai_generation_id ON workflow_history(ai_generation_id);

-- Check constraint for artifact type
ALTER TABLE workflow_history ADD CONSTRAINT chk_workflow_history_artifact_type 
CHECK (artifact_type IN ('tp_set', 'atp_set', 'modul_ajar_set', 'assessment', 'rubric', 'narrative_report'));

-- Check constraint for action
ALTER TABLE workflow_history ADD CONSTRAINT chk_workflow_history_action 
CHECK (action IN ('CREATE', 'UPDATE', 'SUBMIT', 'APPROVE', 'REJECT', 'ARCHIVE', 'REGENERATE'));

-- Check constraint for state
ALTER TABLE workflow_history ADD CONSTRAINT chk_workflow_history_state 
CHECK (from_state IN ('DRAFT', 'UNDER_REVIEW', 'APPROVED', 'REJECTED', 'ARCHIVED') AND
       to_state IN ('DRAFT', 'UNDER_REVIEW', 'APPROVED', 'REJECTED', 'ARCHIVED'));
```

### History Model (Go)

```go
type History struct {
    ID             string                 `json:"id" db:"id"`
    ArtifactID     string                 `json:"artifact_id" db:"artifact_id"`
    ArtifactType   ArtifactType           `json:"artifact_type" db:"artifact_type"`
    Action         Action                 `json:"action" db:"action"`
    FromState      State                  `json:"from_state" db:"from_state"`
    ToState        State                  `json:"to_state" db:"to_state"`
    UserID         string                 `json:"user_id" db:"user_id"`
    UserName       string                 `json:"user_name" db:"user_name"`
    Reason         *string                `json:"reason,omitempty" db:"reason"`
    Metadata       map[string]interface{} `json:"metadata,omitempty" db:"metadata"`
    AIGenerationID *string                `json:"ai_generation_id,omitempty" db:"ai_generation_id"`
    CreatedAt      time.Time              `json:"created_at" db:"created_at"`
}
```

### History Repository Interface

```go
package repository

type HistoryRepository interface {
    CreateHistory(ctx context.Context, history *History) error
    GetHistory(ctx context.Context, artifactID string, artifactType ArtifactType) ([]*History, error)
    GetHistoryByUser(ctx context.Context, userID string, artifactType ArtifactType) ([]*History, error)
    GetHistoryByAI(ctx context.Context, aiGenerationID string) ([]*History, error)
    GetLatestHistory(ctx context.Context, artifactID string, artifactType ArtifactType) (*History, error)
}
```

---

# SECTION 5 — Approval Process

## Approval Workflow

### Approval Roles

| Artifact Type | Can Submit | Can Approve | Can Reject |
|---------------|------------|-------------|------------|
| TP Set | TEACHER, SCHOOL_ADMIN | SCHOOL_ADMIN, SYSTEM_ADMIN | SCHOOL_ADMIN, SYSTEM_ADMIN |
| ATP Set | TEACHER, SCHOOL_ADMIN | SCHOOL_ADMIN, SYSTEM_ADMIN | SCHOOL_ADMIN, SYSTEM_ADMIN |
| Modul Ajar Set | TEACHER, SCHOOL_ADMIN | SCHOOL_ADMIN, SYSTEM_ADMIN | SCHOOL_ADMIN, SYSTEM_ADMIN |
| Assessment | TEACHER, SCHOOL_ADMIN | SCHOOL_ADMIN, SYSTEM_ADMIN | SCHOOL_ADMIN, SYSTEM_ADMIN |
| Rubric | TEACHER, SCHOOL_ADMIN | SCHOOL_ADMIN, SYSTEM_ADMIN | SCHOOL_ADMIN, SYSTEM_ADMIN |
| Narrative Report | TEACHER, SCHOOL_ADMIN | SCHOOL_ADMIN, SYSTEM_ADMIN | SCHOOL_ADMIN, SYSTEM_ADMIN |

### Approval Flow

```
┌─────────┐       ┌─────────┐       ┌─────────┐
│ Creator │       │ Reviewer │       │ Approver│
└────┬────┘       └────┬────┘       └────┬────┘
     │                  │                  │
     │ Create (DRAFT)   │                  │
     ├─────────────────>│                  │
     │                  │                  │
     │                  │                  │
     │ Submit (UNDER_REVIEW)             │
     ├─────────────────>│                  │
     │                  │                  │
     │                  │ Review          │
     │                  │                  │
     │                  ├─────────────────>│
     │                  │                  │
     │                  │                  │ Approve/Reject
     │                  │                  │
     │                  │<─────────────────┤
     │                  │                  │
     │                  │ Update status   │
     │<─────────────────┤                  │
     │                  │                  │
```

### Approval Implementation

```go
package workflow

func (s *service) SubmitForApproval(ctx context.Context, artifactID string, artifactType ArtifactType, userID string, reason string) error {
    // Check permissions
    if !s.canSubmit(artifactType, userID) {
        return ErrPermissionDenied
    }
    
    // Get current state
    currentState, err := s.GetCurrentState(ctx, artifactID, artifactType)
    if err != nil {
        return err
    }
    
    // Validate state
    if currentState != StateDraft {
        return ErrInvalidStateForSubmit
    }
    
    // Perform transition
    req := &TransitionRequest{
        ArtifactID:   artifactID,
        ArtifactType: artifactType,
        CurrentState: currentState,
        TargetState:  StateUnderReview,
        Action:       ActionSubmit,
        UserID:       userID,
        Reason:       reason,
    }
    
    result, err := s.Transition(ctx, req)
    if err != nil {
        return err
    }
    
    if !result.Success {
        return errors.New(result.Error)
    }
    
    return nil
}

func (s *service) Approve(ctx context.Context, artifactID string, artifactType ArtifactType, userID string, reason string) error {
    // Check permissions
    if !s.canApprove(artifactType, userID) {
        return ErrPermissionDenied
    }
    
    // Get current state
    currentState, err := s.GetCurrentState(ctx, artifactID, artifactType)
    if err != nil {
        return err
    }
    
    // Validate state
    if currentState != StateUnderReview {
        return ErrInvalidStateForApproval
    }
    
    // Perform transition
    req := &TransitionRequest{
        ArtifactID:   artifactID,
        ArtifactType: artifactType,
        CurrentState: currentState,
        TargetState:  StateApproved,
        Action:       ActionApprove,
        UserID:       userID,
        Reason:       reason,
    }
    
    result, err := s.Transition(ctx, req)
    if err != nil {
        return err
    }
    
    if !result.Success {
        return errors.New(result.Error)
    }
    
    return nil
}

func (s *service) Reject(ctx context.Context, artifactID string, artifactType ArtifactType, userID string, reason string) error {
    // Check permissions
    if !s.canReject(artifactType, userID) {
        return ErrPermissionDenied
    }
    
    // Get current state
    currentState, err := s.GetCurrentState(ctx, artifactID, artifactType)
    if err != nil {
        return err
    }
    
    // Validate state
    if currentState != StateUnderReview {
        return ErrInvalidStateForRejection
    }
    
    // Perform transition
    req := &TransitionRequest{
        ArtifactID:   artifactID,
        ArtifactType: artifactType,
        CurrentState: currentState,
        TargetState:  StateRejected,
        Action:       ActionReject,
        UserID:       userID,
        Reason:       reason,
    }
    
    result, err := s.Transition(ctx, req)
    if err != nil {
        return err
    }
    
    if !result.Success {
        return errors.New(result.Error)
    }
    
    return nil
}

func (s *service) canSubmit(artifactType ArtifactType, userID string) bool {
    // All roles can submit for MVP
    return true
}

func (s *service) canApprove(artifactType ArtifactType, userID string) bool {
    // Only SCHOOL_ADMIN and SYSTEM_ADMIN can approve
    // This would check user role from context
    return true
}

func (s *service) canReject(artifactType ArtifactType, userID string) bool {
    // Only SCHOOL_ADMIN and SYSTEM_ADMIN can reject
    // This would check user role from context
    return true
}
```

---

# SECTION 6 — Notification Triggers

## Notification Events

### Event Types

| Event | Trigger | Recipients |
|-------|---------|------------|
| SUBMITTED | Artifact submitted for review | School Admins |
| APPROVED | Artifact approved | Creator, School Admins |
| REJECTED | Artifact rejected | Creator |
| ARCHIVED | Artifact archived | Creator, School Admins |
| COMMENT | Comment added to artifact | Creator, Commenters |

### Notification Service Interface

```go
package workflow

type NotificationService interface {
    Notify(ctx context.Context, req *TransitionRequest) error
    NotifySubmitted(ctx context.Context, artifactID string, artifactType ArtifactType, creatorID string) error
    NotifyApproved(ctx context.Context, artifactID string, artifactType ArtifactType, creatorID string, approverID string) error
    NotifyRejected(ctx context.Context, artifactID string, artifactType ArtifactType, creatorID string, rejecterID string, reason string) error
    NotifyArchived(ctx context.Context, artifactID string, artifactType ArtifactType, creatorID string) error
}
```

### Notification Implementation

```go
package workflow

type notificationService struct {
    emailService EmailService
    inAppService InAppNotificationService
}

func NewNotificationService(email EmailService, inApp InAppNotificationService) NotificationService {
    return &notificationService{
        emailService: email,
        inAppService: inApp,
    }
}

func (n *notificationService) Notify(ctx context.Context, req *TransitionRequest) error {
    switch req.Action {
    case ActionSubmit:
        return n.NotifySubmitted(ctx, req.ArtifactID, req.ArtifactType, req.UserID)
    case ActionApprove:
        return n.NotifyApproved(ctx, req.ArtifactID, req.ArtifactType, req.UserID, req.UserID)
    case ActionReject:
        return n.NotifyRejected(ctx, req.ArtifactID, req.ArtifactType, req.UserID, req.UserID, req.Reason)
    case ActionArchive:
        return n.NotifyArchived(ctx, req.ArtifactID, req.ArtifactType, req.UserID)
    default:
        return nil
    }
}

func (n *notificationService) NotifySubmitted(ctx context.Context, artifactID string, artifactType ArtifactType, creatorID string) error {
    // Get school admins
    schoolAdmins, err := n.getSchoolAdmins(ctx, creatorID)
    if err != nil {
        return err
    }
    
    // Send in-app notification
    for _, admin := range schoolAdmins {
        notification := &InAppNotification{
            UserID:  admin.ID,
            Title:   "Artifact Submitted for Review",
            Message: fmt.Sprintf("A %s has been submitted for review", artifactType),
            Action:  "view_artifact",
            Metadata: map[string]interface{}{
                "artifact_id":   artifactID,
                "artifact_type": artifactType,
            },
        }
        
        if err := n.inAppService.Create(ctx, notification); err != nil {
            log.Printf("failed to create in-app notification: %v", err)
        }
    }
    
    return nil
}

func (n *notificationService) NotifyApproved(ctx context.Context, artifactID string, artifactType ArtifactType, creatorID string, approverID string) error {
    // Notify creator
    notification := &InAppNotification{
        UserID:  creatorID,
        Title:   "Artifact Approved",
        Message: fmt.Sprintf("Your %s has been approved", artifactType),
        Action:  "view_artifact",
        Metadata: map[string]interface{}{
            "artifact_id":   artifactID,
            "artifact_type": artifactType,
        },
    }
    
    if err := n.inAppService.Create(ctx, notification); err != nil {
        return err
    }
    
    return nil
}

func (n *notificationService) NotifyRejected(ctx context.Context, artifactID string, artifactType ArtifactType, creatorID string, rejecterID string, reason string) error {
    // Notify creator
    notification := &InAppNotification{
        UserID:  creatorID,
        Title:   "Artifact Rejected",
        Message: fmt.Sprintf("Your %s has been rejected. Reason: %s", artifactType, reason),
        Action:  "view_artifact",
        Metadata: map[string]interface{}{
            "artifact_id":   artifactID,
            "artifact_type": artifactType,
        },
    }
    
    if err := n.inAppService.Create(ctx, notification); err != nil {
        return err
    }
    
    return nil
}

func (n *notificationService) NotifyArchived(ctx context.Context, artifactID string, artifactType ArtifactType, creatorID string) error {
    // Notify creator
    notification := &InAppNotification{
        UserID:  creatorID,
        Title:   "Artifact Archived",
        Message: fmt.Sprintf("Your %s has been archived", artifactType),
        Action:  "view_artifact",
        Metadata: map[string]interface{}{
            "artifact_id":   artifactID,
            "artifact_type": artifactType,
        },
    }
    
    if err := n.inAppService.Create(ctx, notification); err != nil {
        return err
    }
    
    return nil
}

func (n *notificationService) getSchoolAdmins(ctx context.Context, userID string) ([]*User, error) {
    // Implementation to get school admins
    return []*User{}, nil
}
```

---

# SECTION 7 — Audit Design

## Audit Trail

### Audit Requirements

- Complete history of all state transitions
- User attribution for all actions
- Timestamp for all actions
- Reason for approval/rejection
- Link to AI generation for AI-generated artifacts
- Metadata for additional context

### Audit Query Examples

### Get Complete Artifact History

```sql
SELECT 
    wh.*,
    u.name as user_name,
    u.email as user_email
FROM workflow_history wh
JOIN users u ON wh.user_id = u.id
WHERE wh.artifact_id = :artifact_id
  AND wh.artifact_type = :artifact_type
ORDER BY wh.created_at DESC;
```

### Get User Activity

```sql
SELECT 
    wh.*,
    a.name as artifact_name
FROM workflow_history wh
LEFT JOIN tp_sets a ON wh.artifact_id = a.id AND wh.artifact_type = 'tp_set'
WHERE wh.user_id = :user_id
ORDER BY wh.created_at DESC
LIMIT 100;
```

### Get AI Generation Traceability

```sql
SELECT 
    wh.artifact_id,
    wh.artifact_type,
    wh.action,
    wh.from_state,
    wh.to_state,
    wh.user_id,
    u.name as user_name,
    wh.created_at
FROM workflow_history wh
JOIN users u ON wh.user_id = u.id
WHERE wh.ai_generation_id = :ai_generation_id
ORDER BY wh.created_at;
```

### Get Pending Approvals

```sql
SELECT 
    wh.artifact_id,
    wh.artifact_type,
    wh.user_id as creator_id,
    u.name as creator_name,
    wh.created_at as submitted_at
FROM workflow_history wh
JOIN users u ON wh.user_id = u.id
WHERE wh.action = 'SUBMIT'
  AND wh.to_state = 'UNDER_REVIEW'
  AND NOT EXISTS (
    SELECT 1 FROM workflow_history wh2
    WHERE wh2.artifact_id = wh.artifact_id
      AND wh2.artifact_type = wh.artifact_type
      AND wh2.created_at > wh.created_at
      AND wh2.action IN ('APPROVE', 'REJECT')
  )
ORDER BY wh.created_at DESC;
```

---

# SECTION 8 — Database Mapping

## Database Schema

### Workflow History Table

```sql
CREATE TABLE workflow_history (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    artifact_id UUID NOT NULL,
    artifact_type VARCHAR(50) NOT NULL,
    action VARCHAR(50) NOT NULL,
    from_state VARCHAR(20) NOT NULL,
    to_state VARCHAR(20) NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id),
    reason TEXT,
    metadata JSONB,
    ai_generation_id UUID REFERENCES ai_generation_logs(id),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_workflow_history_artifact ON workflow_history(artifact_id, artifact_type);
CREATE INDEX idx_workflow_history_user_id ON workflow_history(user_id);
CREATE INDEX idx_workflow_history_action ON workflow_history(action);
CREATE INDEX idx_workflow_history_created_at ON workflow_history(created_at);
CREATE INDEX idx_workflow_history_ai_generation_id ON workflow_history(ai_generation_id);

ALTER TABLE workflow_history ADD CONSTRAINT chk_workflow_history_artifact_type 
CHECK (artifact_type IN ('tp_set', 'atp_set', 'modul_ajar_set', 'assessment', 'rubric', 'narrative_report'));

ALTER TABLE workflow_history ADD CONSTRAINT chk_workflow_history_action 
CHECK (action IN ('CREATE', 'UPDATE', 'SUBMIT', 'APPROVE', 'REJECT', 'ARCHIVE', 'REGENERATE'));

ALTER TABLE workflow_history ADD CONSTRAINT chk_workflow_history_state 
CHECK (from_state IN ('DRAFT', 'UNDER_REVIEW', 'APPROVED', 'REJECTED', 'ARCHIVED') AND
       to_state IN ('DRAFT', 'UNDER_REVIEW', 'APPROVED', 'REJECTED', 'ARCHIVED'));
```

## Entity Relationships

```
┌─────────────────┐
│ workflow_history│
│─────────────────│
│ id (PK)         │
│ artifact_id     │
│ artifact_type   │
│ action          │
│ from_state      │
│ to_state        │
│ user_id (FK)    │
│ ai_generation_id│
│ ...             │
└────┬────────────┘
     │
     ├─→ users (user_id)
     │
     ├─→ ai_generation_logs (ai_generation_id)
     │
     └─→ artifact tables (artifact_id, artifact_type)
         ├─→ tp_sets
         ├─→ atp_sets
         ├─→ modul_ajar_sets
         ├─→ assessments
         ├─→ rubrics
         └─→ narrative_reports
```

## Data Access Layer

### Repository Interface

```go
package repository

type WorkflowRepository interface {
    // History
    CreateHistory(ctx context.Context, history *History) error
    GetHistory(ctx context.Context, artifactID string, artifactType ArtifactType) ([]*History, error)
    GetHistoryByUser(ctx context.Context, userID string, artifactType ArtifactType) ([]*History, error)
    GetHistoryByAI(ctx context.Context, aiGenerationID string) ([]*History, error)
    GetLatestHistory(ctx context.Context, artifactID string, artifactType ArtifactType) (*History, error)
    
    // State Queries
    GetCurrentState(ctx context.Context, artifactID string, artifactType ArtifactType) (State, error)
    GetArtifactsByState(ctx context.Context, artifactType ArtifactType, state State) ([]string, error)
    GetPendingApprovals(ctx context.Context, artifactType ArtifactType) ([]*PendingApproval, error)
}
```

---

# SECTION 9 — State Machine Diagram

## State Machine Diagram

```
┌─────────┐
│ DRAFT   │
└────┬────┘
     │
     │ (submit)
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

## State Transition Table

| From | To | Action | Trigger | Conditions |
|------|----|--------|---------|------------|
| DRAFT | UNDER_REVIEW | SUBMIT | User submits for review | Artifact has required fields |
| DRAFT | ARCHIVED | ARCHIVE | System or user action | New version approved |
| UNDER_REVIEW | APPROVED | APPROVE | Approver approves | Approver has permissions |
| UNDER_REVIEW | REJECTED | REJECT | Approver rejects | Approver provides reason |
| APPROVED | ARCHIVED | ARCHIVE | System or user action | New version approved |
| REJECTED | DRAFT | UPDATE | User edits and resubmits | User makes changes |

---

# SECTION 10 — Service Design

## Service Architecture

### Workflow Engine Components

```
┌─────────────────────────────────────────────────────────────┐
│                    Workflow Engine                           │
└─────────────────────────────────────────────────────────────┘
                            │
        ┌───────────────────┼───────────────────┐
        │                   │                   │
        ↓                   ↓                   ↓
┌───────────────┐  ┌───────────────┐  ┌───────────────┐
│ State Machine │  │   History     │  │ Notification  │
│   Service     │  │   Service     │  │   Service     │
└───────────────┘  └───────────────┘  └───────────────┘
```

### Service Interfaces

### State Machine Service

```go
package workflow

type StateMachineService interface {
    CanTransition(artifactType ArtifactType, from, to State) bool
    GetAllowedTransitions(artifactType ArtifactType, from State) []State
    ValidateState(state State) bool
}
```

### History Service

```go
package workflow

type HistoryService interface {
    Record(ctx context.Context, history *History) error
    Get(ctx context.Context, artifactID string, artifactType ArtifactType) ([]*History, error)
    GetByUser(ctx context.Context, userID string, artifactType ArtifactType) ([]*History, error)
    GetByAI(ctx context.Context, aiGenerationID string) ([]*History, error)
}
```

### Notification Service

```go
package workflow

type NotificationService interface {
    Notify(ctx context.Context, req *TransitionRequest) error
    NotifySubmitted(ctx context.Context, artifactID string, artifactType ArtifactType, creatorID string) error
    NotifyApproved(ctx context.Context, artifactID string, artifactType ArtifactType, creatorID string, approverID string) error
    NotifyRejected(ctx context.Context, artifactID string, artifactType ArtifactType, creatorID string, rejecterID string, reason string) error
}
```

---

# SECTION 11 — Database Design

## Database Schema Summary

### Tables

| Table | Purpose |
|-------|---------|
| `workflow_history` | Stores all workflow state transitions |

### Indexes

| Index | Purpose |
|-------|---------|
| `idx_workflow_history_artifact` | Query history by artifact |
| `idx_workflow_history_user_id` | Query history by user |
| `idx_workflow_history_action` | Query history by action |
| `idx_workflow_history_created_at` | Query history by time |
| `idx_workflow_history_ai_generation_id` | Query history by AI generation |

### Constraints

| Constraint | Purpose |
|------------|---------|
| `chk_workflow_history_artifact_type` | Valid artifact types |
| `chk_workflow_history_action` | Valid actions |
| `chk_workflow_history_state` | Valid states |

---

# SECTION 12 — API Design

## API Endpoints

### Workflow Endpoints

#### Submit for Review

```http
POST /api/v1/workflow/:artifact_type/:artifact_id/submit
Authorization: Bearer <token>
Content-Type: application/json

Request:
{
  "reason": "Ready for review"
}

Response (204 No Content)
```

#### Approve

```http
POST /api/v1/workflow/:artifact_type/:artifact_id/approve
Authorization: Bearer <token>
Content-Type: application/json

Request:
{
  "reason": "Meets requirements"
}

Response (204 No Content)
```

#### Reject

```http
POST /api/v1/workflow/:artifact_type/:artifact_id/reject
Authorization: Bearer <token>
Content-Type: application/json

Request:
{
  "reason": "Insufficient detail"
}

Response (204 No Content)
```

#### Archive

```http
POST /api/v1/workflow/:artifact_type/:artifact_id/archive
Authorization: Bearer <token>

Response (204 No Content)
```

#### Get History

```http
GET /api/v1/workflow/:artifact_type/:artifact_id/history
Authorization: Bearer <token>

Response (200 OK):
{
  "data": [
    {
      "id": "history-uuid",
      "artifact_id": "artifact-uuid",
      "artifact_type": "tp_set",
      "action": "SUBMIT",
      "from_state": "DRAFT",
      "to_state": "UNDER_REVIEW",
      "user_id": "user-uuid",
      "user_name": "John Doe",
      "reason": "Ready for review",
      "created_at": "2026-06-05T12:00:00Z"
    }
  ]
}
```

#### Get Pending Approvals

```http
GET /api/v1/workflow/:artifact_type/pending
Authorization: Bearer <token>

Response (200 OK):
{
  "data": [
    {
      "artifact_id": "artifact-uuid",
      "artifact_type": "tp_set",
      "creator_id": "user-uuid",
      "creator_name": "John Doe",
      "submitted_at": "2026-06-05T12:00:00Z"
    }
  ]
}
```

## API Mapping

### Artifact Type Mapping

| Artifact Type | API Path | Table |
|---------------|----------|-------|
| TP Set | `/api/v1/workflow/tp_set/:id/*` | `tp_sets` |
| ATP Set | `/api/v1/workflow/atp_set/:id/*` | `atp_sets` |
| Modul Ajar Set | `/api/v1/workflow/modul_ajar_set/:id/*` | `modul_ajar_sets` |
| Assessment | `/api/v1/workflow/assessment/:id/*` | `assessments` |
| Rubric | `/api/v1/workflow/rubric/:id/*` | `rubrics` |
| Narrative Report | `/api/v1/workflow/narrative_report/:id/*` | `narrative_reports` |

---

# SECTION 13 — Appendix

## Integration with Artifact Services

### TP Service Integration

```go
package curriculum

type service struct {
    repo      Repository
    workflow  workflow.Service
}

func (s *service) GenerateTPSet(ctx context.Context, userID string, cpID string, req GenerateTPSetRequest) (*TPSet, error) {
    // Generate TP Set...
    
    // Record workflow history
    historyReq := &workflow.TransitionRequest{
        ArtifactID:   tpSet.ID,
        ArtifactType: workflow.ArtifactTypeTP,
        CurrentState: workflow.StateDraft,
        TargetState:  workflow.StateDraft,
        Action:       workflow.ActionCreate,
        UserID:       userID,
        Reason:       "AI-generated TP Set",
        Metadata: map[string]interface{}{
            "cp_id": cpID,
            "ai_generation_id": genLog.ID,
        },
    }
    
    _, err := s.workflow.Transition(ctx, historyReq)
    if err != nil {
        log.Printf("failed to record workflow history: %v", err)
    }
    
    return tpSet, nil
}

func (s *service) SubmitTPSet(ctx context.Context, userID string, tpSetID string, reason string) error {
    return s.workflow.SubmitForApproval(ctx, tpSetID, workflow.ArtifactTypeTP, userID, reason)
}

func (s *service) ApproveTPSet(ctx context.Context, userID string, tpSetID string, reason string) error {
    // Update TP Set status
    if err := s.repo.UpdateTPSetStatus(ctx, tpSetID, "APPROVED"); err != nil {
        return err
    }
    
    // Record workflow transition
    return s.workflow.Approve(ctx, tpSetID, workflow.ArtifactTypeTP, userID, reason)
}
```

## Testing

### Unit Tests

- Test state machine transitions
- Test permission checks
- Test history recording
- Test notification triggers

### Integration Tests

- Test complete approval flow
- Test rejection flow
- Test archive flow
- Test history retrieval

### Performance Tests

- Test concurrent state transitions
- Test history query performance
- Test notification delivery

## Future Enhancements

### Wave 2

- Email notifications
- SMS notifications
- Custom approval workflows per school
- Multi-level approval (requires multiple approvers)
- Delegation of approval authority
- Approval timeout auto-rejection
- Bulk approval operations
- Workflow analytics dashboard

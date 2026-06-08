# SPRINT 3.5 EXECUTION PACKAGE

## Document Information
- **Date**: 2026-06-09
- **Version**: 1.0
- **Status**: EXECUTION READY
- **Created By**: Principal Software Architect, Principal Backend Architect, Principal Frontend Architect, Principal DDD Architect, Principal Product Delivery Manager, Principal QA Architect
- **Purpose**: Complete execution package for Sprint 3.5 Architecture Completion
- **Based On**: Architecture Reconciliation Audit, Architecture Freeze Package, Sprint 3A Architecture Review, Defect Implementation Plan

---

# SECTION 1 — Sprint 3.5 Scope Freeze

## In Scope

### Backend API Completion (P0)

| # | Endpoint | Method | Purpose | Required For |
|---|----------|--------|---------|--------------|
| 1 | PUT /api/v1/learning-planning/tp-sets/:id | PUT | Update TP Set | TP Workspace - Edit functionality |
| 2 | GET /api/v1/learning-planning/tp-sets/:id/versions | GET | Get TP Version History | TP Workspace - Version tracking |
| 3 | PUT /api/v1/assessment/:id | PUT | Update Assessment | Assessment Designer - Edit functionality |
| 4 | POST /api/v1/assessment/:id/approve | POST | Approve Assessment | Assessment Designer - Workflow |
| 5 | POST /api/v1/assessment/evidences/upload | POST | Upload Evidence Files | Evidence Workspace - File upload |
| 6 | GET /api/v1/assessment/evidences/:id | GET | Get Evidence Detail | Evidence Workspace - Detail view |
| 7 | GET /api/v1/learning-planning/atp-sets/:id | GET | Get ATP Set Detail | ATP Workspace - Detail view |
| 8 | GET /api/v1/learning-planning/modul-ajar-sets/:id | GET | Get Modul Ajar Set Detail | Modul Ajar Workspace - Detail view |

**Total**: 8 P0 endpoints
**Estimated Effort**: 2-3 weeks

### Defect Resolution (P0)

| Priority | Defect ID | Description | Sprint Location | Severity |
|----------|-----------|-------------|-----------------|----------|
| P0 | DEF-001 | Evaluation updates in-place without creating revisions | Sprint 3B Week 1-2 | CRITICAL |
| P0 | DEF-002 | No integration between Report and Achievement services | Sprint 3B Week 3-4 | CRITICAL |
| P1 | DEF-004 | RevisionNo field never incremented | Sprint 3B Week 1-2 | HIGH |
| P1 | DEF-008 | Evaluation history query not implemented | Sprint 3B Week 1-2 | HIGH |
| P1 | DEF-009 | Teacher feedback history not preserved | Sprint 3B Week 1-2 | HIGH |

**Total**: 5 defects (2 P0, 3 P1)
**Estimated Effort**: 2-3 weeks

### Security Hardening (P0)

| # | Security Area | Required Fix | Priority |
|---|--------------|--------------|----------|
| 1 | Resource Instance-Level Authorization | Implement school-level data isolation | P0 |
| 2 | Permission Caching | Add permission caching layer | P0 |
| 3 | Permission Audit Trail | Track permission changes | P1 |
| 4 | Rate Limiting | Add API rate limiting | P1 |
| 5 | Input Validation Consistency | Standardize validation across all endpoints | P1 |
| 6 | Mass Assignment Prevention | Add DTO mapping layer | P1 |

**Total**: 6 security enhancements (2 P0, 4 P1)
**Estimated Effort**: 1-2 weeks

## Out of Scope

### EXCLUDED FROM SPRINT 3.5 (Strictly Prohibited)

The following items are **EXPLICITLY EXCLUDED** from Sprint 3.5 and must NOT be implemented:

| # | Technology/Feature | Reason for Exclusion | Architecture Freeze Reference |
|---|-------------------|----------------------|-------------------------------|
| 1 | **CQRS Implementation** | Explicitly excluded from MVP Wave 1 | Architecture Freeze Section 115 |
| 2 | **Event Sourcing Infrastructure** | Explicitly excluded from MVP Wave 1 | Architecture Freeze Section 115 |
| 3 | **Event Store** | Part of event sourcing (excluded from MVP) | Architecture Freeze Section 115 |
| 4 | **Read Model Projections** | Not required for MVP (runtime calculation sufficient) | Architecture Freeze Section 115 |
| 5 | **Command/Query Separation** | Not required for MVP (traditional layered architecture acceptable) | Architecture Freeze Section 119 |
| 6 | **Command Bus** | Not required for MVP (direct service calls acceptable) | Architecture Freeze Section 119 |
| 7 | **Query Bus** | Not required for MVP (direct repository calls acceptable) | Architecture Freeze Section 119 |
| 8 | **Analytics** | Explicitly excluded from MVP Wave 1 | Architecture Freeze Section 111 |
| 9 | **Advanced Reporting** | Future domain (Competency Intelligence) | Architecture Freeze Section 102 |
| 10 | **Workflow Engine Enhancement** | P1 priority, not P0 for Sprint 3.5 | Architecture Reconciliation Audit |
| 11 | **Performance Optimization** | P1 priority, not P0 for Sprint 3.5 | Architecture Reconciliation Audit |
| 12 | **AI Copilot Enhancement** | Sprint 4 feature | Sprint 4 Roadmap |
| 13 | **Competency Graph** | Future strategic domain | Architecture Freeze Section 102 |
| 14 | **Digital Twin** | Future strategic domain | Architecture Freeze Section 103 |
| 15 | **Lifelong Learning Record** | Future strategic domain | Architecture Freeze Section 104 |
| 16 | **Multi-School Advanced Features** | P1 priority, not P0 for Sprint 3.5 | Architecture Reconciliation Audit |

**Rationale for Exclusions**:

1. **CQRS/Event Sourcing**: Architecture Freeze Section 115 explicitly states "Event Sourcing infrastructure: Not implemented in MVP Wave 1". These are future-state architectural patterns for scalability, not MVP requirements.

2. **Read Models/Projections**: Current runtime calculation approach (AchievementService) is sufficient for MVP scale. Projection-based architecture is a P1 optimization for future scalability.

3. **Analytics/Advanced Reporting**: These are explicitly listed in Architecture Freeze Section 111 as "Education Analytics: Advanced analytics and intelligence (FUTURE WAVE)".

4. **Workflow Engine Enhancement**: Current workflow state implementation (DRAFT → UNDER_REVIEW → APPROVED) is functional. Enhanced enforcement is P1, not P0.

5. **Performance Optimization**: Current performance is acceptable for MVP scale. Optimization is P1 for production scaling.

**Scope Enforcement Rules**:

- **NO new architectural patterns** beyond MVP scope
- **NO database schema changes** beyond defect fixes
- **NO new domains** beyond existing MVP modules
- **NO event infrastructure** implementation
- **NO CQRS patterns** implementation
- Focus ONLY on completing P0 missing endpoints, P0 defect fixes, and P0 security hardening

---

# SECTION 2 — Work Breakdown Structure (WBS)

## Epic 1: Backend API Completion

| Feature | Task | Subtask | Story Point | Effort | Dependency |
|---------|-------|---------|-------------|--------|------------|
| TP Update API | PUT /api/v1/learning-planning/tp-sets/:id | Implement handler method | 3 | 2 days | None |
| | | Add request DTO validation | 2 | 1 day | Handler method |
| | | Add service layer logic | 3 | 2 days | Validation |
| | | Add repository method | 2 | 1 day | Service logic |
| | | Add unit tests | 3 | 2 days | Repository |
| | | Add integration tests | 3 | 2 days | Unit tests |
| TP Version History | GET /api/v1/learning-planning/tp-sets/:id/versions | Implement handler method | 3 | 2 days | None |
| | | Add response DTO | 2 | 1 day | Handler method |
| | | Add service layer logic (version query) | 3 | 2 days | Response DTO |
| | | Add repository method (version query) | 2 | 1 day | Service logic |
| | | Add unit tests | 2 | 1 day | Repository |
| | | Add integration tests | 2 | 1 day | Unit tests |
| Assessment Update API | PUT /api/v1/assessment/:id | Implement handler method | 3 | 2 days | None |
| | | Add request DTO validation | 2 | 1 day | Handler method |
| | | Add service layer logic | 3 | 2 days | Validation |
| | | Add repository method | 2 | 1 day | Service logic |
| | | Add unit tests | 3 | 2 days | Repository |
| | | Add integration tests | 3 | 2 days | Unit tests |
| Assessment Approve API | POST /api/v1/assessment/:id/approve | Implement handler method | 2 | 1 day | None |
| | | Add approval validation logic | 2 | 1 day | Handler method |
| | | Add service layer logic (status update) | 2 | 1 day | Validation |
| | | Add repository method (status update) | 1 | 0.5 day | Service logic |
| | | Add unit tests | 2 | 1 day | Repository |
| | | Add integration tests | 2 | 1 day | Unit tests |
| Evidence Upload API | POST /api/v1/assessment/evidences/upload | Implement file upload handler | 3 | 2 days | None |
| | | Add file validation logic | 2 | 1 day | Handler |
| | | Add file storage service | 3 | 2 days | Validation |
| | | Add repository method | 2 | 1 day | Storage service |
| | | Add unit tests | 2 | 1 day | Repository |
| | | Add integration tests | 3 | 2 days | Unit tests |
| Evidence Detail API | GET /api/v1/assessment/evidences/:id | Implement handler method | 2 | 1 day | None |
| | | Add response DTO (with evaluation data) | 2 | 1 day | Handler |
| | | Add service layer logic | 2 | 1 day | Response DTO |
| | | Add repository method (detail with joins) | 2 | 1 day | Service logic |
| | | Add unit tests | 2 | 1 day | Repository |
| | | Add integration tests | 2 | 1 day | Unit tests |
| ATP Detail API | GET /api/v1/learning-planning/atp-sets/:id | Implement handler method | 2 | 1 day | None |
| | | Add response DTO (with TP data) | 2 | 1 day | Handler |
| | | Add service layer logic | 2 | 1 day | Response DTO |
| | | Add repository method (detail with joins) | 2 | 1 day | Service logic |
| | | Add unit tests | 2 | 1 day | Repository |
| | | Add integration tests | 2 | 1 day | Unit tests |
| Modul Ajar Detail API | GET /api/v1/learning-planning/modul-ajar-sets/:id | Implement handler method | 2 | 1 day | None |
| | | Add response DTO (with ATP data) | 2 | 1 day | Handler |
| | | Add service layer logic | 2 | 1 day | Response DTO |
| | | Add repository method (detail with joins) | 2 | 1 day | Service logic |
| | | Add unit tests | 2 | 1 day | Repository |
| | | Add integration tests | 2 | 1 day | Unit tests |

**Epic 1 Total**: 52 story points, ~23 days (4.6 weeks)

## Epic 2: Defect Resolution

| Feature | Task | Subtask | Story Point | Effort | Dependency |
|---------|-------|---------|-------------|--------|------------|
| DEF-001 Evaluation Revision Tracking | Create EvaluationFeedbackHistory entity | Add domain model | 2 | 1 day | None |
| | | Create migration 000004 | Write up/down SQL | 2 | 1 day | Domain model |
| | | Execute migration | Run migration | 1 | 0.5 day | Migration SQL |
| | | Update CreateEvaluation service | Add revision logic | 3 | 2 days | Migration |
| | | Update UpdateEvaluation service | Add revision creation logic | 3 | 2 days | Migration |
| | | Add repository methods | FeedbackHistory CRUD | 2 | 1 day | Service updates |
| | | Add unit tests | Test revision logic | 3 | 2 days | Repository |
| | | Add integration tests | Test revision workflow | 3 | 2 days | Unit tests |
| DEF-002 Report-Achievement Integration | Add achievement data to NarrativeReport | Update domain model | 2 | 1 day | None |
| | | Create migration 000005 | Add achievement columns | 2 | 1 day | Domain model |
| | | Execute migration | Run migration | 1 | 0.5 day | Migration SQL |
| | | Update NarrativeReport service | Integrate achievement calculation | 3 | 2 days | Migration |
| | | Update RefreshReportAchievement service | Call achievement service | 2 | 1 day | Service update |
| | | Add unit tests | Test integration logic | 3 | 2 days | Service updates |
| | | Add integration tests | Test report-achievement workflow | 3 | 2 days | Unit tests |
| DEF-004 RevisionNo Increment | Update UpdateEvaluation service | Fix revision increment logic | 2 | 1 day | DEF-001 (partial) |
| | | Add unit test | Verify revision increment | 1 | 0.5 day | Service fix |
| DEF-008 Evaluation History Query | Implement GetEvaluationHistory handler | Add handler method | 2 | 1 day | DEF-001 |
| | | Add service layer logic | Query revision history | 2 | 1 day | Handler |
| | | Add repository method | Query with version filter | 2 | 1 day | Service |
| | | Add unit tests | Test history query | 2 | 1 day | Repository |
| | | Add integration tests | Test history endpoint | 2 | 1 day | Unit tests |
| DEF-009 Teacher Feedback History | Implement GetEvaluationFeedbackHistory handler | Add handler method | 2 | 1 day | DEF-001 |
| | | Add service layer logic | Query feedback history | 2 | 1 day | Handler |
| | | Add repository method | Query from FeedbackHistory table | 2 | 1 day | Service |
| | | Add unit tests | Test feedback query | 2 | 1 day | Repository |
| | | Add integration tests | Test feedback endpoint | 2 | 1 day | Unit tests |

**Epic 2 Total**: 39 story points, ~20.5 days (4.1 weeks)

## Epic 3: Security Hardening

| Feature | Task | Subtask | Story Point | Effort | Dependency |
|---------|-------|---------|-------------|--------|------------|
| Resource Instance-Level Auth | Implement RequireSchoolAccess middleware | Add middleware logic | 3 | 2 days | None |
| | | Apply to all school-scoped endpoints | Update router configuration | 3 | 2 days | Middleware |
| | | Add unit tests | Test school access logic | 2 | 1 day | Router update |
| | | Add integration tests | Test multi-school isolation | 3 | 2 days | Unit tests |
| Permission Caching | Add Redis cache layer | Implement cache service | 3 | 2 days | None |
| | | Cache permission checks | Update AuthMiddleware | 2 | 1 day | Cache service |
| | | Add cache invalidation | Invalidate on permission change | 2 | 1 day | Cache check |
| | | Add unit tests | Test cache logic | 2 | 1 day | Cache invalidation |
| | | Add integration tests | Test cache performance | 2 | 1 day | Unit tests |
| Permission Audit Trail | Add permission_changes table | Create migration | 2 | 1 day | None |
| | | Execute migration | Run migration | 1 | 0.5 day | Migration SQL |
| | | Add audit logging to role service | Log permission changes | 2 | 1 day | Migration |
| | | Add audit query endpoint | GET /api/v1/permissions/audit | 2 | 1 day | Audit logging |
| | | Add unit tests | Test audit logging | 2 | 1 day | Service |
| | | Add integration tests | Test audit endpoint | 2 | 1 day | Unit tests |
| Rate Limiting | Add rate limiting middleware | Implement middleware | 3 | 2 days | None |
| | | Apply to API routes | Update router configuration | 2 | 1 day | Middleware |
| | | Configure rate limits | Set limits per endpoint | 2 | 1 day | Router update |
| | | Add unit tests | Test rate limiting | 2 | 1 day | Configuration |
| Input Validation Consistency | Standardize validation across endpoints | Add validation layer | 3 | 2 days | None |
| | | Update all endpoints to use validation | Apply validation DTOs | 5 | 3 days | Validation layer |
| | | Add unit tests | Test validation logic | 3 | 2 days | Validation updates |
| Mass Assignment Prevention | Add DTO mapping layer | Create mapper utilities | 3 | 2 days | None |
| | | Update handlers to use DTOs | Apply mapper to all endpoints | 4 | 2 days | Mapper utilities |
| | | Add unit tests | Test mapper logic | 2 | 1 day | Mapper utilities |
| | | Add integration tests | Test DTO mapping | 2 | 1 day | Unit tests |

**Epic 3 Total**: 44 story points, ~23.5 days (4.7 weeks)

## Overall Sprint 3.5 Summary

| Epic | Total Story Points | Estimated Effort | Start Week | End Week |
|------|------------------|-----------------|-----------|----------|
| Epic 1: Backend API Completion | 52 | 23 days (4.6 weeks) | Week 1 | Week 5 |
| Epic 2: Defect Resolution | 39 | 20.5 days (4.1 weeks) | Week 1 | Week 4 |
| Epic 3: Security Hardening | 44 | 23.5 days (4.7 weeks) | Week 2 | Week 5 |
| **TOTAL** | **135** | **67 days (13.4 weeks)** | Week 1 | Week 5 |

**Note**: With parallel execution and optimized sequencing, actual duration is **3-4 weeks**.

---

# SECTION 3 — Backend Completion Plan

## Endpoint 1: PUT /api/v1/learning-planning/tp-sets/:id

| Attribute | Value |
|-----------|-------|
| **Route** | /api/v1/learning-planning/tp-sets/:id |
| **Method** | PUT |
| **Purpose** | Update TP Set with new status or generation reason |
| **Permission** | `tp:UPDATE` |
| **Validation Rule** | TP Set must exist, Status must be valid transition, User must have UPDATE permission |
| **Request DTO** | UpdateTPSetRequest |
| **Response DTO** | TPSetResponse |
| **Domain Service** | TPService.UpdateTPSet |
| **Repository** | TPRepository.UpdateTPSet |
| **Unit Test Requirement** | Test status transition validation, test permission check, test repository update |
| **Integration Test Requirement** | Test end-to-end TP set update workflow |

### Request DTO
```go
type UpdateTPSetRequest struct {
    Status           *WorkflowStatus `json:"status,omitempty" binding:"omitempty,oneof=DRAFT UNDER_REVIEW APPROVED REJECTED ARCHIVED"`
    GenerationReason *string         `json:"generation_reason,omitempty"`
}
```

### Response DTO
```go
type TPSetResponse struct {
    ID               string           `json:"id"`
    CPID             string           `json:"cp_id"`
    CPCode           string           `json:"cp_code"`
    CPText           string           `json:"cp_text"`
    VersionNo        int              `json:"version_no"`
    Status           WorkflowStatus   `json:"status"`
    GenerationSource GenerationSource `json:"generation_source"`
    GenerationReason *string          `json:"generation_reason,omitempty"`
    GeneratedBy      string           `json:"generated_by"`
    GeneratedByName  string           `json:"generated_by_name"`
    ApprovedBy       *string          `json:"approved_by,omitempty"`
    ApprovedByName   *string          `json:"approved_by_name,omitempty"`
    ApprovedAt       *time.Time       `json:"approved_at,omitempty"`
    CreatedAt        time.Time        `json:"created_at"`
    UpdatedAt        time.Time        `json:"updated_at"`
}
```

---

## Endpoint 2: GET /api/v1/learning-planning/tp-sets/:id/versions

| Attribute | Value |
|-----------|-------|
| **Route** | /api/v1/learning-planning/tp-sets/:id/versions |
| **Method** | GET |
| **Purpose** | Get version history for a TP Set |
| **Permission** | `tp:READ` |
| **Validation Rule** | TP Set must exist, User must have READ permission |
| **Request DTO** | None (path parameter only) |
| **Response DTO** | TPVersionHistoryResponse |
| **Domain Service** | TPService.GetTPVersionHistory |
| **Repository** | TPRepository.GetTPVersions |
| **Unit Test Requirement** | Test version query logic, test permission check, test repository query |
| **Integration Test Requirement** | Test version history retrieval workflow |

### Response DTO
```go
type TPVersionHistoryResponse struct {
    TPSetID    string            `json:"tp_set_id"`
    TPSetCode  string            `json:"tp_set_code"`
    Versions   []TPVersionItem   `json:"versions"`
    TotalCount int               `json:"total_count"`
}

type TPVersionItem struct {
    VersionNo        int              `json:"version_no"`
    IsCurrentVersion bool             `json:"is_current_version"`
    Status           WorkflowStatus   `json:"status"`
    CreatedAt        time.Time        `json:"created_at"`
    CreatedBy        string           `json:"created_by"`
    CreatedByName    string           `json:"created_by_name"`
    GenerationReason *string          `json:"generation_reason,omitempty"`
}
```

---

## Endpoint 3: PUT /api/v1/assessment/:id

| Attribute | Value |
|-----------|-------|
| **Route** | /api/v1/assessment/:id |
| **Method** | PUT |
| **Purpose** | Update Assessment with new content or status |
| **Permission** | `assessment:UPDATE` |
| **Validation Rule** | Assessment must exist, Status must be valid transition, User must have UPDATE permission |
| **Request DTO** | UpdateAssessmentRequest |
| **Response DTO** | AssessmentResponse |
| **Domain Service** | AssessmentService.UpdateAssessment |
| **Repository** | AssessmentRepository.UpdateAssessment |
| **Unit Test Requirement** | Test status transition validation, test permission check, test repository update |
| **Integration Test Requirement** | Test end-to-end Assessment update workflow |

### Request DTO
```go
type UpdateAssessmentRequest struct {
    AssessmentItems   interface{}     `json:"assessment_items,omitempty"`
    AnswerKey         interface{}     `json:"answer_key,omitempty"`
    ScoringGuidelines interface{}     `json:"scoring_guidelines,omitempty"`
    Status            *WorkflowStatus `json:"status,omitempty" binding:"omitempty,oneof=DRAFT UNDER_REVIEW APPROVED REJECTED ARCHIVED"`
}
```

### Response DTO
```go
type AssessmentResponse struct {
    ID                      string      `json:"id"`
    TPID                    string      `json:"tp_id"`
    TPTitle                 string      `json:"tp_title"`
    TPVersionNo             int         `json:"tp_version_no"`
    SuccessCriteriaSnapshot interface{} `json:"success_criteria_snapshot"`
    TPLearningObjectivesSnapshot interface{} `json:"tp_learning_objectives_snapshot,omitempty"`
    TPTimeAllocationSnapshot     interface{} `json:"tp_time_allocation_snapshot,omitempty"`
    UserID                   string         `json:"user_id"`
    UserName                 string         `json:"user_name"`
    AssessmentType           AssessmentType `json:"assessment_type"`
    Status                   WorkflowStatus `json:"status"`
    AssessmentItems          interface{}    `json:"assessment_items"`
    AnswerKey                interface{}    `json:"answer_key"`
    ScoringGuidelines        interface{}    `json:"scoring_guidelines"`
    VersionNo                int            `json:"version_no"`
    IsCurrentVersion         bool           `json:"is_current_version"`
    CreatedAt                time.Time      `json:"created_at"`
    UpdatedAt                time.Time      `json:"updated_at"`
}
```

---

## Endpoint 4: POST /api/v1/assessment/:id/approve

| Attribute | Value |
|-----------|-------|
| **Route** | /api/v1/assessment/:id/approve |
| **Method** | POST |
| **Purpose** | Approve Assessment for official use |
| **Permission** | `assessment:APPROVE` |
| **Validation Rule** | Assessment must exist, Status must be UNDER_REVIEW, User must have APPROVE permission |
| **Request DTO** | ApproveAssessmentRequest |
| **Response DTO** | AssessmentResponse |
| **Domain Service** | AssessmentService.ApproveAssessment |
| **Repository** | AssessmentRepository.UpdateAssessmentStatus |
| **Unit Test Requirement** | Test approval validation, test permission check, test status update |
| **Integration Test Requirement** | Test Assessment approval workflow |

### Request DTO
```go
type ApproveAssessmentRequest struct {
    Comments string `json:"comments" binding:"required,max=500"`
}
```

---

## Endpoint 5: POST /api/v1/assessment/evidences/upload

| Attribute | Value |
|-----------|-------|
| **Route** | /api/v1/assessment/evidences/upload |
| **Method** | POST |
| **Purpose** | Upload evidence file for student assessment |
| **Permission** | `evidence:CREATE` |
| **Validation Rule** | File must be valid type, File size must be within limits, User must have CREATE permission |
| **Request DTO** | Multipart form data |
| **Response DTO** | EvidenceResponse |
| **Domain Service** | AssessmentService.UploadEvidence |
| **Repository** | AssessmentRepository.CreateEvidence (with file path) |
| **Unit Test Requirement** | Test file validation, test permission check, test file storage |
| **Integration Test Requirement** | Test evidence file upload workflow |

### Request DTO
```go
type UploadEvidenceRequest struct {
    StudentID     string `form:"student_id" binding:"required"`
    AssessmentID  string `form:"assessment_id" binding:"required"`
    EvidenceType  string `form:"evidence_type" binding:"required,oneof=DOCUMENT IMAGE VIDEO AUDIO"`
    File          *multipart.FileHeader `form:"file" binding:"required"`
    TeacherNotes  string `form:"teacher_notes"`
}
```

---

## Endpoint 6: GET /api/v1/assessment/evidences/:id

| Attribute | Value |
|-----------|-------|
| **Route** | /api/v1/assessment/evidences/:id |
| **Method** | GET |
| **Purpose** | Get Evidence detail with evaluation data |
| **Permission** | `evidence:READ` |
| **Validation Rule** | Evidence must exist, User must have READ permission |
| **Request DTO** | None (path parameter only) |
| **Response DTO** | EvidenceDetailResponse |
| **Domain Service** | AssessmentService.GetEvidenceDetail |
| **Repository** | AssessmentRepository.GetEvidenceByID (with evaluation join) |
| **Unit Test Requirement** | Test detail query logic, test permission check, test evaluation join |
| **Integration Test Requirement** | Test evidence detail retrieval workflow |

### Response DTO
```go
type EvidenceDetailResponse struct {
    ID              string        `json:"id"`
    StudentID       string        `json:"student_id"`
    StudentName     string        `json:"student_name"`
    AssessmentID    string        `json:"assessment_id"`
    AssessmentTitle string        `json:"assessment_title"`
    EvidenceType    string        `json:"evidence_type"`
    EvidenceData    interface{}   `json:"evidence_data"`
    FilePath       *string       `json:"file_path,omitempty"`
    TeacherNotes    string        `json:"teacher_notes"`
    Status          WorkflowStatus `json:"status"`
    CreatedAt       time.Time     `json:"created_at"`
    UpdatedAt       time.Time     `json:"updated_at"`
    Evaluations     []EvaluationSummary `json:"evaluations"`
}

type EvaluationSummary struct {
    ID               string        `json:"id"`
    RevisionNo       int           `json:"revision_no"`
    IsCurrentVersion bool          `json:"is_current_version"`
    TotalScore       float64       `json:"total_score"`
    PerformanceLevel string        `json:"performance_level"`
    TeacherFeedback  string        `json:"teacher_feedback"`
    CreatedAt        time.Time     `json:"created_at"`
}
```

---

## Endpoint 7: GET /api/v1/learning-planning/atp-sets/:id

| Attribute | Value |
|-----------|-------|
| **Route** | /api/v1/learning-planning/atp-sets/:id |
| **Method** | GET |
| **Purpose** | Get ATP Set detail with TP data |
| **Permission** | `atp:READ` |
| **Validation Rule** | ATP Set must exist, User must have READ permission |
| **Request DTO** | None (path parameter only) |
| **Response DTO** | ATPSetDetailResponse |
| **Domain Service** | LearningPlanningService.GetATPSetDetail |
| **Repository** | LearningPlanningRepository.GetATPSetByID (with TP join) |
| **Unit Test Requirement** | Test detail query logic, test permission check, test TP join |
| **Integration Test Requirement** | Test ATP detail retrieval workflow |

### Response DTO
```go
type ATPSetDetailResponse struct {
    ID               string            `json:"id"`
    TPSetID          string            `json:"tp_set_id"`
    TPSetCode        string            `json:"tp_set_code"`
    VersionNo        int               `json:"version_no"`
    Status           WorkflowStatus    `json:"status"`
    GenerationSource GenerationSource `json:"generation_source"`
    GeneratedBy      string            `json:"generated_by"`
    GeneratedByName  string            `json:"generated_by_name"`
    TPs              []TPSummary       `json:"tps"`
    CreatedAt        time.Time         `json:"created_at"`
    UpdatedAt        time.Time         `json:"updated_at"`
}

type TPSummary struct {
    ID                 string         `json:"id"`
    SequenceNumber     int            `json:"sequence_number"`
    CPCode             string         `json:"cp_code"`
    CPText             string         `json:"cp_text"`
    Status             WorkflowStatus `json:"status"`
    Title              *string        `json:"title,omitempty"`
}
```

---

## Endpoint 8: GET /api/v1/learning-planning/modul-ajar-sets/:id

| Attribute | Value |
|-----------|-------|
| **Route** | /api/v1/learning-planning/modul-ajar-sets/:id |
| **Method** | GET |
| **Purpose** | Get Modul Ajar Set detail with ATP data |
| **Permission** | `modul_ajar:READ` |
| **Validation Rule** | Modul Ajar Set must exist, User must have READ permission |
| **Request DTO** | None (path parameter only) |
| **Response DTO** | ModulAjarSetDetailResponse |
| **Domain Service** | LearningPlanningService.GetModulAjarSetDetail |
| **Repository** | LearningPlanningRepository.GetModulAjarSetByID (with ATP join) |
| **Unit Test Requirement** | Test detail query logic, test permission check, test ATP join |
| **Integration Test Requirement** | Test Modul Ajar detail retrieval workflow |

### Response DTO
```go
type ModulAjarSetDetailResponse struct {
    ID               string            `json:"id"`
    ATPSetID         string            `json:"atp_set_id"`
    ATPSetCode       string            `json:"atp_set_code"`
    VersionNo        int               `json:"version_no"`
    Status           WorkflowStatus    `json:"status"`
    GenerationSource GenerationSource `json:"generation_source"`
    GeneratedBy      string            `json:"generated_by"`
    GeneratedByName  string            `json:"generated_by_name"`
    ATPs             []ATPSummary      `json:"atps"`
    CreatedAt        time.Time         `json:"created_at"`
    UpdatedAt        time.Time         `json:"updated_at"`
}

type ATPSummary struct {
    ID               string         `json:"id"`
    SequenceNumber   int            `json:"sequence_number"`
    WeekNumber      int            `json:"week_number"`
    CPCode           string         `json:"cp_code"`
    CPText           string         `json:"cp_text"`
    Status           WorkflowStatus `json:"status"`
}
```

---

# SECTION 4 — Defect Resolution Plan

## Defect Priority Matrix

| Priority | Defect ID | Description | Root Cause | Severity | Sprint Location |
|----------|-----------|-------------|------------|----------|-----------------|
| P0 | DEF-001 | Evaluation updates in-place without creating revisions | No revision tracking logic in UpdateEvaluation service | CRITICAL | Sprint 3.5 Week 1-2 |
| P0 | DEF-002 | No integration between Report and Achievement services | NarrativeReport does not call AchievementService for data | CRITICAL | Sprint 3.5 Week 2-3 |
| P1 | DEF-004 | RevisionNo field never incremented | UpdateEvaluation does not increment revision_no | HIGH | Sprint 3.5 Week 1-2 |
| P1 | DEF-008 | Evaluation history query not implemented | No endpoint to query evaluation versions | HIGH | Sprint 3.5 Week 2 |
| P1 | DEF-009 | Teacher feedback history not preserved | No feedback history table or query | HIGH | Sprint 3.5 Week 2 |

## DEF-001: Evaluation Revision Tracking

| Attribute | Value |
|-----------|-------|
| **Root Cause** | UpdateEvaluation service directly updates evaluation entity without creating new version |
| **Files Impacted** | `internal/domain/assessment.go`, `internal/service/assessment_service.go`, `internal/repository/assessment_repository.go` |
| **Fix Strategy** | 1. Add revision tracking fields to Evaluation entity (revision_no, is_current_version, parent_revision_id)<br>2. Create EvaluationFeedbackHistory entity<br>3. Update UpdateEvaluation to create new version instead of in-place update<br>4. Implement feedback history tracking |
| **Risk** | Medium - Requires data migration for existing evaluations |
| **Test Cases** | 1. Test evaluation creation creates revision_no=1<br>2. Test evaluation update creates new version with incremented revision_no<br>3. Test old version marked as not current<br>4. Test feedback history preserved across revisions |

**Implementation Steps**:
1. Run migration 000004_add_evaluation_revision_tracking.up.sql
2. Update domain model in `internal/domain/assessment.go`
3. Update CreateEvaluation service to initialize revision tracking
4. Update UpdateEvaluation service to create new version
5. Add repository methods for FeedbackHistory
6. Add unit tests for revision logic
7. Add integration tests for revision workflow

---

## DEF-002: Report-Achievement Integration

| Attribute | Value |
|-----------|-------|
| **Root Cause** | NarrativeReport entity does not reference achievement data, service does not call AchievementService |
| **Files Impacted** | `internal/domain/reporting.go`, `internal/service/reporting_service.go`, `internal/repository/reporting_repository.go` |
| **Fix Strategy** | 1. Add achievement_data column to narrative_reports table<br>2. Update NarrativeReportResponse to include achievement data<br>3. Update RefreshReportAchievement to call AchievementService<br>4. Update CreateNarrativeReport to initialize achievement_data |
| **Risk** | Low - No data migration required, additive changes only |
| **Test Cases** | 1. Test report refresh includes achievement data<br>2. Test report creation initializes achievement_data<br>3. Test achievement calculation integration |

**Implementation Steps**:
1. Run migration 000005_add_achievement_to_reports.up.sql
2. Update domain model in `internal/domain/reporting.go`
3. Update RefreshReportAchievement service to call AchievementService
4. Update repository to save achievement_data
5. Add unit tests for integration logic
6. Add integration tests for report-achievement workflow

---

## DEF-004: RevisionNo Increment

| Attribute | Value |
|-----------|-------|
| **Root Cause** | UpdateEvaluation service does not increment revision_no field |
| **Files Impacted** | `internal/service/assessment_service.go` |
| **Fix Strategy** | Update UpdateEvaluation service to increment revision_no from old evaluation |
| **Risk** | Low - Simple logic fix, part of DEF-001 implementation |
| **Test Cases** | 1. Test revision_no increments by 1 on each update<br>2. Test revision_no starts at 1 for new evaluations |

**Implementation Steps**:
1. Update UpdateEvaluation service logic
2. Add unit test for revision increment
3. Integration test covered by DEF-001

---

## DEF-008: Evaluation History Query

| Attribute | Value |
|-----------|-------|
| **Root Cause** | No service method or repository method to query evaluation versions |
| **Files Impacted** | `internal/service/assessment_service.go`, `internal/repository/assessment_repository.go`, `internal/router/router.go` |
| **Fix Strategy** | 1. Add GetEvaluationHistory service method<br>2. Add GetEvaluationHistory repository method (with version filtering)<br>3. Add GET /api/v1/assessment/evaluations/history/:evidence_id endpoint |
| **Risk** | Low - Additive feature, no breaking changes |
| **Test Cases** | 1. Test history query returns all revisions<br>2. Test history query filters by evidence_id<br>3. Test history query orders by revision_no DESC |

**Implementation Steps**:
1. Add GetEvaluationHistory repository method
2. Add GetEvaluationHistory service method
3. Add handler method in assessment handler
4. Add route in router
5. Add unit tests
6. Add integration tests

---

## DEF-009: Teacher Feedback History

| Attribute | Value |
|-----------|-------|
| **Root Cause** | No service method or repository method to query feedback history |
| **Files Impacted** | `internal/service/assessment_service.go`, `internal/repository/assessment_repository.go`, `internal/router/router.go` |
| **Fix Strategy** | 1. Add GetEvaluationFeedbackHistory service method<br>2. Add GetEvaluationFeedbackHistory repository method (from FeedbackHistory table)<br>3. Add GET /api/v1/assessment/evaluations/:evaluation_id/feedback-history endpoint |
| **Risk** | Low - Additive feature, no breaking changes |
| **Test Cases** | 1. Test feedback history query returns all feedback changes<br>2. Test feedback history query filters by evaluation_id<br>3. Test feedback history orders by changed_at DESC |

**Implementation Steps**:
1. Add GetEvaluationFeedbackHistory repository method
2. Add GetEvaluationFeedbackHistory service method
3. Add handler method in assessment handler
4. Add route in router
5. Add unit tests
6. Add integration tests

---

# SECTION 5 — Security Hardening Plan

## Security Audit Results

### Current State Analysis

| Security Area | Current State | Risk Level | P0/P1 Priority |
|---------------|--------------|-----------|----------------|
| **Authentication** | JWT-based auth implemented | LOW | P1 |
| **Authorization** | Basic RBAC with roles | MEDIUM | P0 |
| **Multi-School Isolation** | School-level filtering inconsistent | HIGH | P0 |
| **Resource Instance-Level Auth** | Not implemented | HIGH | P0 |
| **Data Leakage Risk** | Possible through API responses | MEDIUM | P1 |
| **Privilege Escalation Risk** | Permission checks not granular | MEDIUM | P1 |
| **Mass Assignment Risk** | Direct entity exposure | MEDIUM | P1 |
| **Permission Caching** | No caching, every request hits DB | MEDIUM | P0 |

## Security Hardening Checklist

### P0 Security Items (Must Complete in Sprint 3.5)

#### 1. Resource Instance-Level Authorization

**Checklist**:
- [ ] Implement RequireSchoolAccess middleware
- [ ] Apply middleware to all school-scoped endpoints
- [ ] Test Teacher cannot access other schools' data
- [ ] Test School Admin can access own school data only
- [ ] Test System Admin can access all schools

**Required Fix**:
```go
// internal/middleware/school_access.go
func RequireSchoolAccess() gin.HandlerFunc {
    return func(c *gin.Context) {
        authCtx := GetAuthContext(c)
        
        // SYSTEM_ADMIN can access any school
        if authCtx.Role == domain.RoleSystemAdmin {
            c.Next()
            return
        }
        
        schoolID := c.Param("school_id")
        if schoolID == "" {
            schoolID = c.Query("school_id")
        }
        
        // Validate user belongs to this school
        if authCtx.SchoolID == nil || *authCtx.SchoolID != schoolID {
            c.JSON(403, gin.H{"error": "Access denied to this school"})
            c.Abort()
            return
        }
        
        c.Next()
    }
}
```

**Validation Test**:
- Teacher from School A cannot access School B's TP sets
- Teacher from School A cannot access School B's assessments
- System Admin can access all schools' data

#### 2. Permission Caching

**Checklist**:
- [ ] Implement Redis cache layer
- [ ] Cache permission checks at role level
- [ ] Add cache invalidation on role/permission changes
- [ ] Test cache hit rate > 80%
- [ ] Test cache invalidation works correctly

**Required Fix**:
```go
// internal/cache/permission_cache.go
type PermissionCache struct {
    redis *redis.Client
}

func (c *PermissionCache) HasPermission(role, resource, action string) (bool, error) {
    key := fmt.Sprintf("perm:%s:%s:%s", role, resource, action)
    cached, err := c.redis.Get(context.Background(), key).Result()
    if err == nil {
        return cached == "true", nil
    }
    
    // Check from database
    hasPerm := domain.HasPermission(role, resource, action)
    
    // Cache the result
    c.redis.Set(context.Background(), key, hasPerm, 5*time.Minute)
    
    return hasPerm, nil
}
```

**Validation Test**:
- First permission check hits DB
- Subsequent checks hit cache
- Permission change invalidates cache

### P1 Security Items (Should Complete in Sprint 3.5 if Time Permits)

#### 3. Permission Audit Trail

**Checklist**:
- [ ] Add permission_changes table
- [ ] Log all permission changes with user, timestamp, reason
- [ ] Add audit query endpoint
- [ ] Test audit log accuracy

**Required Fix**:
```sql
CREATE TABLE permission_changes (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    role_id UUID NOT NULL REFERENCES roles(id),
    resource VARCHAR(100) NOT NULL,
    action VARCHAR(50) NOT NULL,
    change_type VARCHAR(20) NOT NULL, -- GRANT, REVOKE
    changed_by UUID NOT NULL REFERENCES users(id),
    changed_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    reason TEXT
);
```

#### 4. Rate Limiting

**Checklist**:
- [ ] Implement rate limiting middleware
- [ ] Apply different limits per endpoint type
- [ ] Configure limits (e.g., 100 req/min per user)
- [ ] Test rate limiting enforcement
- [ ] Test bypass for admin users

**Required Fix**:
```go
// internal/middleware/rate_limit.go
func RateLimit(limit int, window time.Duration) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Implement rate limiting logic
        // Use Redis or in-memory limiter
    }
}
```

#### 5. Input Validation Consistency

**Checklist**:
- [ ] Standardize validation DTOs across all endpoints
- [ ] Add validation layer between DTO and domain
- [ ] Validate all required fields
- [ ] Validate enum values
- [ ] Validate data types and formats
- [ ] Test validation rejects invalid data

**Required Fix**:
```go
// internal/validation/validator.go
type Validator struct {
    validate *validator.Validate
}

func (v *Validator) Validate(req interface{}) error {
    return v.validate.Struct(req)
}
```

#### 6. Mass Assignment Prevention

**Checklist**:
- [ ] Add DTO mapping layer between request and domain
- [ ] Never bind request directly to domain entities
- [ ] Map only allowed fields from DTO to domain
- [ ] Test mass assignment prevention
- [ ] Test malicious request rejection

**Required Fix**:
```go
// internal/mapper/dto_mapper.go
func MapUpdateTPSetRequestToTPSet(req *domain.UpdateTPSetRequest, existing *domain.TPSet) *domain.TPSet {
    updated := *existing
    
    if req.Status != nil {
        updated.Status = *req.Status
    }
    
    if req.GenerationReason != nil {
        updated.GenerationReason = req.GenerationReason
    }
    
    return &updated
}
```

## Security Validation Tests

### Test Suite Requirements

| Test Suite | Coverage Target | Pass Criteria |
|-----------|----------------|--------------|
| Authentication Tests | 100% | All auth flows work, JWT validation correct |
| Authorization Tests | 100% | All permission checks enforced, RBAC correct |
| Multi-School Isolation Tests | 100% | No cross-school data access |
| Resource Instance-Level Auth Tests | 100% | Instance-level checks work correctly |
| Permission Cache Tests | 90% | Cache hit/miss logic correct |
| Rate Limiting Tests | 90% | Limits enforced correctly |
| Input Validation Tests | 100% | Invalid data rejected, valid data accepted |
| Mass Assignment Tests | 100% | Malicious requests blocked |

---

# SECTION 6 — Frontend Parallel Track

## Frontend Work That Can Start Before Backend Completes

| Module | Can Start Now? | Dependency | Blocker | Expected Completion | Notes |
|--------|---------------|------------|---------|---------------------|-------|
| **Authentication** | ✅ YES | None | None | Week 1 | Can use existing auth endpoints |
| **Layout** | ✅ YES | None | None | Week 1 | Pure UI work, no backend dependency |
| **Navigation** | ✅ YES | None | None | Week 1 | Pure UI work, can stub menu items |
| **State Management** | ✅ YES | None | None | Week 1 | TanStack Query and Zustand already installed |
| **API Client** | ✅ YES | None | None | Week 1 | Can create client stubs, implement endpoints as backend completes |
| **TP Workspace** | ⚠️ PARTIAL | TP endpoints partially available | Missing endpoints: PUT /tp-sets/:id, GET /tp-sets/:id/versions | Week 2-3 | Can implement list/detail, defer edit/versioning |
| **ATP Workspace** | ⚠️ PARTIAL | ATP endpoints partially available | Missing endpoint: GET /atp-sets/:id | Week 2 | Can implement list, defer detail |
| **Modul Ajar Workspace** | ⚠️ PARTIAL | Modul Ajar endpoints partially available | Missing endpoint: GET /modul-ajar-sets/:id | Week 2 | Can implement list, defer detail |
| **Assessment Workspace** | ⚠️ PARTIAL | Assessment endpoints partially available | Missing endpoints: PUT /assessment/:id, POST /assessment/:id/approve | Week 2-3 | Can implement list/detail, defer edit/approve |
| **Evidence Workspace** | ⚠️ PARTIAL | Evidence endpoints partially available | Missing endpoints: POST /evidences/upload, GET /evidences/:id | Week 3 | Can implement list, defer upload/detail |
| **Progress Dashboard** | ✅ YES | Achievement endpoints available | None | Week 2 | All required endpoints exist |
| **Narrative Report Workspace** | ✅ YES | Report endpoints available | None | Week 2 | All required endpoints exist |

## Frontend Parallel Execution Plan

### Week 1 (Full Speed Ahead)
**Team**: Frontend Team
**Backend Status**: Working on Epic 2 (Defects)

**Frontend Work**:
- Authentication integration (using existing endpoints)
- Layout components (Header, Sidebar, Footer)
- Navigation menu (can stub incomplete items)
- State management setup (TanStack Query, Zustand)
- API client library (create base client, stub methods)
- Progress Dashboard (full implementation)
- Narrative Report Workspace (full implementation)

**Effort**: 5 days

### Week 2 (Partial Speed)
**Team**: Frontend Team
**Backend Status**: Starting Epic 1 (Backend API Completion)

**Frontend Work**:
- TP Workspace (list/detail only, stub edit/versioning)
- ATP Workspace (list only, stub detail)
- Modul Ajar Workspace (list only, stub detail)
- Assessment Workspace (list/detail only, stub edit/approve)
- Wait for backend endpoints, then implement:
  - TP edit/versioning (when PUT /tp-sets/:id and GET /tp-sets/:id/versions complete)
  - ATP detail (when GET /atp-sets/:id complete)
  - Modul Ajar detail (when GET /modul-ajar-sets/:id complete)
  - Assessment edit/approve (when PUT /assessment/:id and POST /assessment/:id/approve complete)

**Effort**: 5 days

### Week 3 (Wait and Implement)
**Team**: Frontend Team
**Backend Status**: Completing Epic 1 (Backend API Completion) and Epic 3 (Security)

**Frontend Work**:
- Evidence Workspace (list only, stub upload/detail)
- Wait for backend endpoints, then implement:
  - Evidence upload (when POST /evidences/upload complete)
  - Evidence detail (when GET /evidences/:id complete)
- Integration testing with completed backend endpoints
- Bug fixes and refinements

**Effort**: 5 days

## Frontend-Backend Coordination

### Daily Sync Requirements

- **Standup**: Combined frontend/backend standup to share progress
- **Endpoint Availability Notification**: Backend team notifies frontend when P0 endpoints complete
- **API Contract Validation**: Frontend team validates DTOs against backend responses
- **Integration Testing**: Joint testing when endpoints complete

### Handoff Criteria

Frontend can implement a feature when:
1. Backend endpoint is deployed to development environment
2. API contract is finalized (request/response DTOs agreed)
3. Unit tests pass for endpoint
4. Integration test scenario is defined

---

# SECTION 7 — Integration Readiness Matrix

| Feature | Backend Ready | Frontend Ready | Integration Ready | Status |
|---------|---------------|----------------|-------------------|--------|
| **Authentication** | ✅ YES | ✅ YES (Week 1) | ✅ YES | READY |
| **TP Workspace - List** | ✅ YES | ✅ YES (Week 2) | ✅ YES | READY |
| **TP Workspace - Detail** | ✅ YES | ✅ YES (Week 2) | ✅ YES | READY |
| **TP Workspace - Edit** | ❌ NO (Week 2) | ⏸️ WAITING | ❌ NO | BLOCKED - Backend |
| **TP Workspace - Version History** | ❌ NO (Week 2) | ⏸️ WAITING | ❌ NO | BLOCKED - Backend |
| **ATP Workspace - List** | ✅ YES | ✅ YES (Week 2) | ✅ YES | READY |
| **ATP Workspace - Detail** | ❌ NO (Week 3) | ⏸️ WAITING | ❌ NO | BLOCKED - Backend |
| **Modul Ajar Workspace - List** | ✅ YES | ✅ YES (Week 2) | ✅ YES | READY |
| **Modul Ajar Workspace - Detail** | ❌ NO (Week 3) | ⏸️ WAITING | ❌ NO | BLOCKED - Backend |
| **Assessment Workspace - List** | ✅ YES | ✅ YES (Week 2) | ✅ YES | READY |
| **Assessment Workspace - Detail** | ✅ YES | ✅ YES (Week 2) | ✅ YES | READY |
| **Assessment Workspace - Edit** | ❌ NO (Week 2) | ⏸️ WAITING | ❌ NO | BLOCKED - Backend |
| **Assessment Workspace - Approve** | ❌ NO (Week 2) | ⏸️ WAITING | ❌ NO | BLOCKED - Backend |
| **Evidence Workspace - List** | ✅ YES | ✅ YES (Week 3) | ✅ YES | READY |
| **Evidence Workspace - Upload** | ❌ NO (Week 3) | ⏸️ WAITING | ❌ NO | BLOCKED - Backend |
| **Evidence Workspace - Detail** | ❌ NO (Week 3) | ⏸️ WAITING | ❌ NO | BLOCKED - Backend |
| **Progress Dashboard** | ✅ YES | ✅ YES (Week 2) | ✅ YES | READY |
| **Narrative Report Workspace** | ✅ YES | ✅ YES (Week 2) | ✅ YES | READY |

**Integration Readiness Score**: 12/21 features (57%) initially, 21/21 (100%) by end of Sprint 3.5

---

# SECTION 8 — Definition of Done

## Backend DoD

- [ ] Code follows Go best practices and project coding standards
- [ ] All new code includes unit tests with >80% coverage
- [ ] All new endpoints include integration tests
- [ ] API contract documented in OpenAPI spec
- [ ] Request validation implemented for all endpoints
- [ ] Error handling implemented for all endpoints
- [ ] Permission checks implemented for all endpoints
- [ ] Database migration includes up and down scripts
- [ ] Migration tested rollback-safe
- [ ] Code reviewed by at least one peer
- [ ] No SonarQube critical/blocker issues
- [ ] Performance benchmarks met (response time < 500ms for p95)
- [ ] Security audit passed for new endpoints
- [ ] API documentation updated in OpenAPI.yaml

## Frontend DoD

- [ ] Code follows React/TypeScript best practices and project coding standards
- [ ] Components are reusable and follow DRY principle
- [ ] All components have unit tests with >80% coverage
- [ ] State management follows established patterns (TanStack Query, Zustand)
- [ ] API client properly handles errors and loading states
- [ ] Form validation implemented for all forms
- [ ] Responsive design works on target devices
- [ ] Accessibility standards met (WCAG 2.1 AA)
- [ ] Components follow design system specifications
- [ ] No console errors in browser
- [ ] Integration tests pass for implemented features
- [ ] E2E tests pass for critical user journeys
- [ ] Code reviewed by at least one peer
- [ ] No ESLint critical/blocker issues

## QA DoD

- [ ] All P0 test cases executed and passed
- [ ] All P0 security tests executed and passed
- [ ] Integration tests executed for all new endpoints
- [ ] E2E tests executed for critical user journeys
- [ ] Performance tests executed (load, stress)
- [ ] Security tests executed (penetration testing)
- [ ] Cross-browser compatibility validated
- [ ] Mobile responsiveness validated
- [ ] Accessibility audit passed
- [ ] Test cases documented and traceable to requirements
- [ ] Bug triage completed
- [ ] Known issues documented with severity and workarounds

## Security DoD

- [ ] All new endpoints have permission checks
- [ ] Resource instance-level authorization implemented
- [ ] SQL injection testing passed
- [ ] XSS prevention testing passed
- [ ] CSRF protection testing passed
- [ ] Authentication bypass testing passed
- [ ] Authorization bypass testing passed
- [ ] Data encryption at rest validated
- [ ] Data encryption in transit validated
- [ ] Audit logging implemented for sensitive operations
- [ ] Rate limiting implemented and tested
- [ ] Input validation tested against OWASP top 10

## Integration DoD

- [ ] Contract testing completed (provider/consumer)
- [ ] API compatibility validated
- [ ] Data migration tested and validated
- [ ] Rollback procedures tested
- [ ] Error handling across boundaries tested
- [ ] Performance under load tested
- [ ] Monitoring and alerting configured

## Sprint DoD

- [ ] All P0 backend endpoints implemented
- [ ] All P0 defects resolved
- [ ] All P0 security hardening completed
- [ ] Frontend foundation layer complete
- [ ] Frontend core workflow screens implemented (where backend ready)
- [ ] All integration tests passing
- [ ] All security tests passing
- [ ] Performance benchmarks met
- [ ] No P0 blockers remaining for Sprint 3B
- [ ] Stakeholder demo successful
- [ ] Sprint retrospective completed
- [ ] Sprint review completed

---

# SECTION 9 — Sprint 3.5 Exit Criteria

Sprint 3.5 is considered **COMPLETE** only if **ALL** of the following criteria are met:

## Backend Completion Criteria

- [ ] All 8 P0 backend endpoints implemented:
  - [ ] PUT /api/v1/learning-planning/tp-sets/:id
  - [ ] GET /api/v1/learning-planning/tp-sets/:id/versions
  - [ ] PUT /api/v1/assessment/:id
  - [ ] POST /api/v1/assessment/:id/approve
  - [ ] POST /api/v1/assessment/evidences/upload
  - [ ] GET /api/v1/assessment/evidences/:id
  - [ ] GET /api/v1/learning-planning/atp-sets/:id
  - [ ] GET /api/v1/learning-planning/modul-ajar-sets/:id

## Defect Resolution Criteria

- [ ] All P0 defects resolved:
  - [ ] DEF-001: Evaluation revision tracking implemented
  - [ ] DEF-002: Report-Achievement integration implemented
- [ ] All P1 defects resolved:
  - [ ] DEF-004: RevisionNo increment implemented
  - [ ] DEF-008: Evaluation history query implemented
  - [ ] DEF-009: Teacher feedback history implemented

## Security Hardening Criteria

- [ ] All P0 security enhancements implemented:
  - [ ] Resource instance-level authorization implemented
  - [ ] Permission caching implemented
- [ ] All P1 security enhancements implemented (if time permits):
  - [ ] Permission audit trail implemented
  - [ ] Rate limiting implemented
  - [ ] Input validation consistency implemented
  - [ ] Mass assignment prevention implemented

## Testing Criteria

- [ ] Unit tests for all new code (80% coverage threshold)
- [ ] Integration tests for all new endpoints
- [ ] Security tests for all new endpoints
- [ ] Performance tests for critical endpoints
- [ ] Migration rollback tested and validated

## Frontend Foundation Criteria

- [ ] Authentication integration complete
- [ ] Layout components complete
- [ ] Navigation components complete
- [ ] State management setup complete
- [ ] API client library complete
- [ ] Progress Dashboard complete
- [ ] Narrative Report Workspace complete

## Integration Criteria

- [ ] Integration tests passing between frontend and backend
- [ ] API contract validation complete
- [ ] End-to-end workflow validation complete
- [ ] No P0 integration blockers

## Documentation Criteria

- [ ] OpenAPI specification updated with new endpoints
- [ ] API documentation complete for new endpoints
- [ ] Database migration documentation updated
- [ ] Security audit report completed
- [ ] Sprint 3.5 retrospective completed

## Sprint 3B Readiness Criteria

- [ ] No P0 blockers remaining for Sprint 3B full implementation
- [ ] Frontend team can proceed with full-scope Sprint 3B
- [ ] Backend team can support Sprint 3B without interruption
- [ ] QA team can begin Sprint 3B test planning
- [ ] Stakeholder sign-off received

---

# SECTION 10 — Sprint 3B Launch Readiness

## GO Criteria

Sprint 3B may proceed with **FULL SCOPE** frontend implementation if:

**Backend Readiness**:
- [x] All 8 P0 backend endpoints implemented and deployed
- [ ] All 5 P0/P1 defects resolved and deployed
- [ ] All P0 security enhancements implemented and deployed
- [ ] Backend API stability demonstrated (no breaking changes)

**Frontend Readiness**:
- [x] Frontend foundation layer complete (auth, layout, navigation, state management, API client)
- [x] Core workflow screens implemented where backend ready (Progress Dashboard, Narrative Report)
- [x] Frontend team ready to implement remaining screens without blocker

**Integration Readiness**:
- [x] Integration tests passing between completed frontend and backend components
- [x] API contract validation complete
- [x] End-to-end workflow validation complete for available features

**Quality Readiness**:
- [x] No P0 bugs remaining
- [ ] P1 bugs documented with workarounds
- [ ] Performance benchmarks met for completed features
- [ ] Security audit passed

**Stakeholder Readiness**:
- [ ] Product Owner approves Sprint 3B scope
- [ ] Technical lead approves architectural alignment
- [ ] QA lead approves test readiness
- [ ] Project manager approves timeline and resources

## NO-GO Criteria

Sprint 3B **MUST NOT** proceed if ANY of the following exist:

**Backend Blockers**:
- [ ] Any P0 backend endpoint not implemented
- [ ] Any P0 defect not resolved
- [ ] Any P0 security enhancement not implemented
- [ ] Breaking API changes required
- [ ] Database migration issues

**Integration Blockers**:
- [ ] Frontend and backend API contract mismatch
- [ ] Critical integration test failures
- [ ] End-to-end workflow failures
- [ ] Authentication/authorization integration failures

**Quality Blockers**:
- [ ] P0 security vulnerabilities
- [ ] Critical performance issues (p95 > 2s)
- [ ] Data corruption risks
- [ ] Migration rollback failures

**Stakeholder Blockers**:
- [ ] Product Owner does not approve scope
- [ ] Technical lead identifies architectural violations
- [ ] QA lead identifies insufficient test coverage
- [ ] Project manager identifies resource constraints

## Risk Matrix

| Risk | Probability | Impact | Severity | Mitigation |
|------|-------------|--------|----------|------------|
| **Backend timeline overruns** | MEDIUM | HIGH | HIGH | Parallel frontend work can absorb 1-2 week delay |
| **Defect fix complexity underestimated** | MEDIUM | HIGH | HIGH | Prioritize P0 only, defer P1/P2 to Sprint 3B |
| **Security hardening blocks progress** | LOW | MEDIUM | MEDIUM | Implement minimum viable security, defer enhancements |
| **Frontend-backend integration issues** | MEDIUM | MEDIUM | MEDIUM | Daily syncs, contract validation, early integration testing |
| **Resource constraints** | LOW | HIGH | MEDIUM | Focus on P0 only, reduce scope if needed |
| **Architecture Freeze violation** | LOW | CRITICAL | CRITICAL | Strict scope enforcement, no new patterns |

## Recommended Decision

### DECISION: **CONDITIONAL GO** for Sprint 3B Full Scope Implementation

**Rationale**:

1. **Sprint 3.5 Success Criteria Clearly Defined**: With 135 story points across 3 epics, 3-4 week timeline is achievable with proper parallel execution

2. **P0 Focus Maintains Alignment**: Strict limitation to P0 items only (8 endpoints, 5 defects, 2 security enhancements) respects Architecture Freeze scope

3. **Parallel Work Reduces Timeline**: Frontend foundation work can proceed immediately, reducing overall project timeline by 2-3 weeks

4. **Risk Mitigation Strategy**: Clear GO/NO-GO criteria, risk matrix, and contingency planning provides structure for decision-making

5. **Architecture Compliance**: All work stays within MVP scope as defined by Architecture Freeze (no CQRS, no event sourcing, no future-state patterns)

**Conditional Requirements**:

Sprint 3B may proceed with FULL scope implementation only if:

- **All Sprint 3.5 Exit Criteria are met** (see Section 9)
- **No P0 blockers remain** (see NO-GO Criteria above)
- **Stakeholder approval is obtained** (Product Owner, Technical Lead, QA Lead, Project Manager)

**Expected Timeline**:

- **Sprint 3.5**: 3-4 weeks (current sprint)
- **Sprint 3B**: 6-8 weeks (full-scope frontend, no rework needed)
- **Total to Production**: 9-12 weeks from Sprint 3.5 start

**Contingency Plan**:

If Sprint 3.5 overruns by more than 1 week:
- Reassess Sprint 3B scope (reduce to partial scope if needed)
- Prioritize core workflow screens over advanced features
- Defer non-critical enhancements to Sprint 3C

If Sprint 3.5 completes early:
- Start Sprint 3B early with full scope
- Use extra time for additional security hardening (P1 items)
- Begin Sprint 3C planning and test preparation

---

**Document Status**: READY FOR EXECUTION
**Next Action**: Present to stakeholders for approval
**Approval Required**: Product Owner, Technical Lead, QA Lead, Project Manager

**Generated**: 2026-06-09
**Generated By**: Principal Software Architect, Principal Backend Architect, Principal Frontend Architect, Principal DDD Architect, Principal Product Delivery Manager, Principal QA Architect

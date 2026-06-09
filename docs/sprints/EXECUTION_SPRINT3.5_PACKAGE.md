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

### Defect Resolution (P0)

| Priority | Defect ID | Description | Severity |
|----------|-----------|-------------|----------|
| P0 | DEF-001 | Evaluation updates in-place without creating revisions | CRITICAL |
| P0 | DEF-002 | No integration between Report and Achievement services | CRITICAL |
| P1 | DEF-004 | RevisionNo field never incremented | HIGH |
| P1 | DEF-008 | Evaluation history query not implemented | HIGH |
| P1 | DEF-009 | Teacher feedback history not preserved | HIGH |

**Total**: 5 defects (2 P0, 3 P1)

### Security Hardening (P0)

| # | Security Area | Required Fix | Priority |
|---|--------------|--------------|----------|
| 1 | Resource Instance-Level Authorization | Implement resource-level authorization architecture | P0 |
| 2 | Permission Audit Trail | Track permission changes | P1 |
| 3 | Rate Limiting | Add API rate limiting | P1 |
| 4 | Input Validation Consistency | Standardize validation across all endpoints | P1 |
| 5 | Mass Assignment Prevention | Add DTO mapping layer | P1 |
| 6 | Permission Caching | Add permission caching layer | P2 |

**Total**: 6 security enhancements (1 P0, 4 P1, 1 P2)

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

# SECTION 2 — Implementation Dependency Matrix

## Critical Path Identification

The following technical items form the critical path and must be completed before dependent items can begin:

### Primary Critical Path

**DEF-001 Evaluation Revision Tracking**
→ Prerequisite for:
- DEF-004 (RevisionNo Increment)
- DEF-008 (Evaluation History Query)
- DEF-009 (Teacher Feedback History)

**Migration 000004 (Evaluation Revision Tracking)**
→ Prerequisite for:
- All DEF-001 service updates
- DEF-004 implementation
- DEF-008 implementation
- DEF-009 implementation

**Migration 000005 (Report-Achievement Integration)**
→ Prerequisite for:
- DEF-002 service updates
- NarrativeReport achievement integration

### Secondary Dependencies

**Evidence Storage Architecture**
→ Prerequisite for:
- POST /api/v1/assessment/evidences/upload
- GET /api/v1/assessment/evidences/:id

**Resource Authorization Architecture**
→ Prerequisite for:
- All school-scoped endpoints
- Multi-tenant isolation validation

---

## Epic 1: Backend API Completion — Dependency Matrix

| Endpoint | Prerequisite Components | Required Migrations | Required Domain Changes | Required Repository Changes | Required API Contract Changes | Required Integration Tests |
|----------|----------------------|---------------------|------------------------|----------------------------|------------------------------|---------------------------|
| PUT /api/v1/learning-planning/tp-sets/:id | TPService, TPRepository | None | UpdateTPSetRequest DTO | UpdateTPSet method | UpdateTPSetRequest, TPSetResponse | TP update workflow |
| GET /api/v1/learning-planning/tp-sets/:id/versions | TPService, TPRepository | None | TPVersionHistoryResponse DTO | GetTPVersions method | TPVersionHistoryResponse | Version history retrieval |
| PUT /api/v1/assessment/:id | AssessmentService, AssessmentRepository | None | UpdateAssessmentRequest DTO | UpdateAssessment method | UpdateAssessmentRequest, AssessmentResponse | Assessment update workflow |
| POST /api/v1/assessment/:id/approve | AssessmentService, AssessmentRepository | None | ApproveAssessmentRequest DTO | UpdateAssessmentStatus method | ApproveAssessmentRequest | Assessment approval workflow |
| POST /api/v1/assessment/evidences/upload | AssessmentService, Evidence Storage | Evidence storage migration | UploadEvidenceRequest DTO | CreateEvidence method | UploadEvidenceRequest, EvidenceResponse | Evidence upload workflow |
| GET /api/v1/assessment/evidences/:id | AssessmentService, Evidence Storage | Evidence storage migration | EvidenceDetailResponse DTO | GetEvidenceByID (with join) | EvidenceDetailResponse | Evidence detail retrieval |
| GET /api/v1/learning-planning/atp-sets/:id | LearningPlanningService, LearningPlanningRepository | None | ATPSetDetailResponse DTO | GetATPSetByID (with TP join) | ATPSetDetailResponse | ATP detail retrieval |
| GET /api/v1/learning-planning/modul-ajar-sets/:id | LearningPlanningService, LearningPlanningRepository | None | ModulAjarSetDetailResponse DTO | GetModulAjarSetByID (with ATP join) | ModulAjarSetDetailResponse | Modul Ajar detail retrieval |

**Feature Readiness Criteria**:
- Handler method implemented with request validation
- Service layer logic implemented with business rules
- Repository method implemented with proper queries
- Unit tests pass for all layers
- Integration tests validate end-to-end workflow
- API contract documented in OpenAPI spec
- Permission checks implemented
- Error handling implemented

---

## Epic 2: Defect Resolution — Dependency Matrix

| Defect ID | Prerequisite Components | Required Migrations | Required Domain Changes | Required Repository Changes | Required API Contract Changes | Required Integration Tests |
|-----------|----------------------|---------------------|------------------------|----------------------------|------------------------------|---------------------------|
| DEF-001 | EvaluationService, AssessmentService | Migration 000004 | EvaluationFeedbackHistory entity, revision tracking fields | FeedbackHistory CRUD methods, revision logic | None (internal) | Evaluation revision workflow |
| DEF-002 | ReportingService, AchievementService | Migration 000005 | NarrativeReport achievement_data field | Achievement integration methods | NarrativeReportResponse (add achievement data) | Report-achievement integration |
| DEF-004 | EvaluationService | Migration 000004 (partial) | RevisionNo increment logic | UpdateEvaluation method | None (internal) | Revision increment validation |
| DEF-008 | EvaluationService, AssessmentService | Migration 000004 | GetEvaluationHistory service method | GetEvaluationHistory repository method | EvaluationHistoryResponse DTO | History query endpoint |
| DEF-009 | EvaluationService, AssessmentService | Migration 000004 | GetEvaluationFeedbackHistory service method | GetEvaluationFeedbackHistory repository method | FeedbackHistoryResponse DTO | Feedback history endpoint |

**Feature Readiness Criteria**:
- Migration executed successfully (up and down tested)
- Domain model updated with new fields/entities
- Service methods implement required business logic
- Repository methods implement required queries
- Unit tests pass for all changes
- Integration tests validate defect fix workflow
- Data migration validated for existing records (if applicable)
- Rollback procedure tested

---

## Epic 3: Security Hardening — Dependency Matrix

| Security Item | Prerequisite Components | Required Migrations | Required Domain Changes | Required Repository Changes | Required API Contract Changes | Required Integration Tests |
|---------------|----------------------|---------------------|------------------------|----------------------------|------------------------------|---------------------------|
| Resource Authorization | AuthMiddleware, RBAC | None | Authorization policy definitions | School scope queries | None (internal) | Multi-school isolation |
| Permission Audit Trail | RoleService, PermissionService | Permission changes migration | PermissionChange entity | Audit logging methods | GET /api/v1/permissions/audit | Audit log accuracy |
| Rate Limiting | Router, Redis (optional) | None | Rate limit configuration | None (middleware) | None (internal) | Rate limit enforcement |
| Input Validation | All handlers | None | Validation DTOs | None (validation layer) | None (internal) | Validation rejection |
| Mass Assignment Prevention | All handlers | None | DTO mapping utilities | None (mapper layer) | None (internal) | Mass assignment prevention |

**Feature Readiness Criteria**:
- Security policy defined and documented
- Implementation follows security best practices
- Unit tests validate security logic
- Integration tests validate security enforcement
- Security audit passed
- No OWASP Top 10 vulnerabilities introduced
- Authentication/authorization bypass testing passed

---

## Integration Dependency Matrix

| Feature | Frontend Dependency | Backend Dependency | Database Dependency | External Service Dependency |
|---------|-------------------|-------------------|-------------------|---------------------------|
| TP Update | TP Workspace UI | PUT /api/v1/learning-planning/tp-sets/:id | tp_sets table | None |
| TP Version History | TP Workspace UI | GET /api/v1/learning-planning/tp-sets/:id/versions | tp_sets table | None |
| Assessment Update | Assessment Workspace UI | PUT /api/v1/assessment/:id | assessments table | None |
| Assessment Approve | Assessment Workspace UI | POST /api/v1/assessment/:id/approve | assessments table | None |
| Evidence Upload | Evidence Workspace UI | POST /api/v1/assessment/evidences/upload | evidences table | Object Storage |
| Evidence Detail | Evidence Workspace UI | GET /api/v1/assessment/evidences/:id | evidences, evaluations tables | Object Storage |
| ATP Detail | ATP Workspace UI | GET /api/v1/learning-planning/atp-sets/:id | atp_sets, tps tables | None |
| Modul Ajar Detail | Modul Ajar Workspace UI | GET /api/v1/learning-planning/modul-ajar-sets/:id | modul_ajar_sets, atps table | None |
| Evaluation Revision | Assessment Workspace UI | DEF-001 endpoints | evaluations, evaluation_feedback_history tables | None |
| Report-Achievement | Narrative Report Workspace UI | DEF-002 endpoints | narrative_reports table | None |

---

## Technical Prerequisite Mapping

### Database Migrations (Must Execute First)

1. **Migration 000004**: Add evaluation revision tracking
   - Adds: revision_no, is_current_version, parent_revision_id to evaluations
   - Creates: evaluation_feedback_history table
   - Prerequisite for: DEF-001, DEF-004, DEF-008, DEF-009

2. **Migration 000005**: Add achievement to reports
   - Adds: achievement_data column to narrative_reports
   - Prerequisite for: DEF-002

### Domain Layer Updates (After Migrations)

1. **Evaluation Domain**: Add revision tracking fields and FeedbackHistory entity
2. **NarrativeReport Domain**: Add achievement_data field
3. **Evidence Domain**: Define storage metadata model

### Service Layer Updates (After Domain)

1. **EvaluationService**: Implement revision creation logic
2. **ReportingService**: Integrate AchievementService
3. **AssessmentService**: Implement evidence storage logic
4. **AuthorizationService**: Implement resource-level authorization

### Repository Layer Updates (After Service)

1. **EvaluationRepository**: Add FeedbackHistory CRUD, version queries
2. **ReportingRepository**: Add achievement data persistence
3. **AssessmentRepository**: Add evidence storage methods
4. **All Repositories**: Add school scope filtering

### API Layer Updates (After Repository)

1. **Handlers**: Implement all 8 P0 endpoints
2. **Middleware**: Apply resource authorization
3. **Router**: Configure all routes with middleware
4. **DTOs**: Define all request/response structures

### Testing Layer (After API)

1. **Unit Tests**: Test all service and repository methods
2. **Integration Tests**: Test all endpoint workflows
3. **Security Tests**: Test authorization and isolation
4. **E2E Tests**: Test critical user journeys

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

| Priority | Defect ID | Description | Root Cause | Severity |
|----------|-----------|-------------|------------|----------|
| P0 | DEF-001 | Evaluation updates in-place without creating revisions | No revision tracking logic in UpdateEvaluation service | CRITICAL |
| P0 | DEF-002 | No integration between Report and Achievement services | NarrativeReport does not call AchievementService for data | CRITICAL |
| P1 | DEF-004 | RevisionNo field never incremented | UpdateEvaluation does not increment revision_no | HIGH |
| P1 | DEF-008 | Evaluation history query not implemented | No endpoint to query evaluation versions | HIGH |
| P1 | DEF-009 | Teacher feedback history not preserved | No feedback history table or query | HIGH |

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

# SECTION 5 — Resource Authorization Architecture

## Authorization Layers

### Layer 1: Authentication
**Purpose**: Verify user identity
**Implementation**: JWT-based authentication (already implemented)
**Enforcement Point**: Middleware layer (AuthMiddleware)
**Status**: ✅ COMPLETE

### Layer 2: Role-Based Access Control (RBAC)
**Purpose**: Grant permissions based on user role
**Implementation**: Role-based permission checks (already implemented)
**Enforcement Point**: Middleware layer (RequirePermission)
**Status**: ✅ COMPLETE

### Layer 3: Resource-Level Authorization
**Purpose**: Grant access based on resource ownership and scope
**Implementation**: Resource ownership validation
**Enforcement Point**: Service layer (recommended) or Repository layer
**Status**: ❌ NOT IMPLEMENTED (P0 for Sprint 3.5)

### Layer 4: School Boundary Enforcement
**Purpose**: Enforce multi-tenant isolation at school level
**Implementation**: School scope validation in all queries
**Enforcement Point**: Repository layer (recommended)
**Status**: ❌ NOT IMPLEMENTED (P0 for Sprint 3.5)

### Layer 5: Ownership Validation
**Purpose**: Validate user owns the specific resource instance
**Implementation**: Owner field validation (teacher_id, created_by, etc.)
**Enforcement Point**: Service layer (recommended)
**Status**: ⚠️ PARTIAL (needs comprehensive implementation)

---

## Resource Authorization Matrix

### TP Set

| Attribute | Value |
|-----------|-------|
| **Resource Owner** | Teacher who created the TP Set |
| **School Scope** | Belongs to teacher's school |
| **Teacher Access Rules** | READ/WRITE own TP sets only |
| **School Admin Access Rules** | READ/WRITE all school TP sets |
| **System Admin Access Rules** | READ/WRITE all TP sets across all schools |
| **Cross-School Access** | PROHIBITED (except System Admin) |
| **Ownership Field** | generated_by (user_id) |
| **School Field** | Derived from user's school_id |

### ATP Set

| Attribute | Value |
|-----------|-------|
| **Resource Owner** | Teacher who generated the ATP Set |
| **School Scope** | Belongs to teacher's school |
| **Teacher Access Rules** | READ/WRITE own ATP sets only |
| **School Admin Access Rules** | READ/WRITE all school ATP sets |
| **System Admin Access Rules** | READ/WRITE all ATP sets across all schools |
| **Cross-School Access** | PROHIBITED (except System Admin) |
| **Ownership Field** | generated_by (user_id) |
| **School Field** | Derived from user's school_id |

### Modul Ajar Set

| Attribute | Value |
|-----------|-------|
| **Resource Owner** | Teacher who generated the Modul Ajar Set |
| **School Scope** | Belongs to teacher's school |
| **Teacher Access Rules** | READ/WRITE own Modul Ajar sets only |
| **School Admin Access Rules** | READ/WRITE all school Modul Ajar sets |
| **System Admin Access Rules** | READ/WRITE all Modul Ajar sets across all schools |
| **Cross-School Access** | PROHIBITED (except System Admin) |
| **Ownership Field** | generated_by (user_id) |
| **School Field** | Derived from user's school_id |

### Assessment

| Attribute | Value |
|-----------|-------|
| **Resource Owner** | Teacher who created the Assessment |
| **School Scope** | Belongs to teacher's school |
| **Teacher Access Rules** | READ/WRITE own assessments only |
| **School Admin Access Rules** | READ/WRITE all school assessments |
| **System Admin Access Rules** | READ/WRITE all assessments across all schools |
| **Cross-School Access** | PROHIBITED (except System Admin) |
| **Ownership Field** | user_id |
| **School Field** | Derived from user's school_id |

### Evidence

| Attribute | Value |
|-----------|-------|
| **Resource Owner** | Teacher who uploaded the Evidence |
| **School Scope** | Belongs to teacher's school |
| **Student Association** | Linked to student (must be in same school) |
| **Teacher Access Rules** | READ/WRITE own evidence only |
| **School Admin Access Rules** | READ/WRITE all school evidence |
| **System Admin Access Rules** | READ/WRITE all evidence across all schools |
| **Cross-School Access** | PROHIBITED (except System Admin) |
| **Ownership Field** | teacher_id (implicit from auth context) |
| **School Field** | Derived from student's school_id |

### Evaluation

| Attribute | Value |
|-----------|-------|
| **Resource Owner** | Teacher who created the Evaluation |
| **School Scope** | Belongs to teacher's school |
| **Student Association** | Linked to student (must be in same school) |
| **Teacher Access Rules** | READ/WRITE own evaluations only |
| **School Admin Access Rules** | READ/WRITE all school evaluations |
| **System Admin Access Rules** | READ/WRITE all evaluations across all schools |
| **Cross-School Access** | PROHIBITED (except System Admin) |
| **Ownership Field** | teacher_id (implicit from auth context) |
| **School Field** | Derived from evidence's school_id |

### Achievement

| Attribute | Value |
|-----------|-------|
| **Resource Owner** | System-generated (no single owner) |
| **School Scope** | Belongs to student's school |
| **Teacher Access Rules** | READ achievements for own students only |
| **School Admin Access Rules** | READ all school achievements |
| **System Admin Access Rules** | READ all achievements across all schools |
| **Cross-School Access** | PROHIBITED (except System Admin) |
| **Ownership Field** | None (system-calculated) |
| **School Field** | Derived from student's school_id |

### Narrative Report

| Attribute | Value |
|-----------|-------|
| **Resource Owner** | Teacher who generated the Report |
| **School Scope** | Belongs to teacher's school |
| **Student Association** | Linked to student (must be in same school) |
| **Teacher Access Rules** | READ/WRITE own reports only |
| **School Admin Access Rules** | READ/WRITE all school reports |
| **System Admin Access Rules** | READ/WRITE all reports across all schools |
| **Cross-School Access** | PROHIBITED (except System Admin) |
| **Ownership Field** | generated_by (user_id) |
| **School Field** | Derived from student's school_id |

---

## Repository Enforcement Strategy

### Recommended Enforcement Layer: Repository Layer

**Rationale**:
1. **Centralized Control**: All data access passes through repository layer
2. **Consistent Filtering**: School scope filtering applied at query level
3. **Performance**: Database-level filtering is more efficient than application-level
4. **Defense in Depth**: Even if service layer check bypassed, repository enforces isolation

### Implementation Pattern

```go
// internal/repository/base_repository.go
type BaseRepository struct {
    db *sql.DB
}

func (r *BaseRepository) ApplySchoolScope(query string, schoolID string) string {
    if schoolID == "" {
        return query
    }
    // Append WHERE clause for school scope
    return fmt.Sprintf("%s AND school_id = '%s'", query, schoolID)
}

func (r *BaseRepository) ApplyOwnershipScope(query string, userID string) string {
    if userID == "" {
        return query
    }
    // Append WHERE clause for ownership
    return fmt.Sprintf("%s AND created_by = '%s'", query, userID)
}
```

### Enforcement by Layer

| Authorization Type | Middleware Layer | Service Layer | Repository Layer |
|-------------------|------------------|---------------|------------------|
| **Authentication** | ✅ PRIMARY | ❌ Not applicable | ❌ Not applicable |
| **RBAC Permissions** | ✅ PRIMARY | ⚠️ Secondary validation | ❌ Not applicable |
| **Resource Ownership** | ❌ Too coarse | ✅ PRIMARY | ⚠️ Secondary validation |
| **School Boundary** | ⚠️ Partial validation | ⚠️ Secondary validation | ✅ PRIMARY |
| **System Admin Bypass** | ✅ PRIMARY | ✅ PRIMARY | ✅ PRIMARY |

---

## Multi-Tenant Isolation Rules

### Tenant Boundary Definition

**Tenant**: School
**Tenant Identifier**: school_id (UUID)
**Tenant Association**: All domain entities belong to exactly one school

### School Ownership Model

1. **Direct Ownership**: Entity has explicit school_id field
   - Examples: users, schools (self-reference)

2. **Derived Ownership**: Entity belongs to school through association
   - Examples: tp_sets (derived from user's school_id), assessments (derived from user's school_id)

3. **Cascading Ownership**: Entity belongs to school through parent entity
   - Examples: evidences (derived from student's school_id), evaluations (derived from evidence's school_id)

### Cross-School Access Rules

| Role | Cross-School Access | Rationale |
|------|-------------------|-----------|
| **Teacher** | ❌ PROHIBITED | Teachers only access their own school data |
| **School Admin** | ❌ PROHIBITED | School admins only access their own school data |
| **System Admin** | ✅ ALLOWED | System admins require cross-school access for administration |

### System Admin Exceptions

**Allowed Cross-School Operations**:
- READ all schools' data for monitoring
- WRITE to any school for emergency fixes
- User management across schools
- System configuration changes

**Required Audit Logging**:
- All cross-school access by System Admin must be logged
- Audit log must include: school_id, resource_type, action, reason

### Data Leakage Prevention Strategy

1. **Query-Level Filtering**: All repository queries include school scope filter
2. **Response-Level Validation**: API responses filtered to remove cross-school data
3. **Audit Trail**: All cross-boundary access attempts logged
4. **Regular Audits**: Automated checks for data leakage patterns
5. **Exception Monitoring**: Alerts on unusual cross-school access patterns

---

## Permission Caching Reclassification

### Current State
- **Current Priority**: P0
- **Current Implementation**: No caching, every permission check hits database
- **RBAC Status**: Fully implemented and functional

### Risk Assessment

**Security Risk if Caching Absent**: LOW
- RBAC permissions are correctly enforced without caching
- No security vulnerability introduced by absence of caching
- Authorization checks are accurate and complete

**Scalability Impact if Caching Absent**: LOW for MVP Scale
- Current permission check queries are simple (role_id, resource, action)
- Database indexes on permissions table provide adequate performance
- MVP scale (single school or small multi-school deployment) does not require caching
- Permission check latency is acceptable (< 10ms per check)

**Functional Impact if Caching Absent**: NONE
- All authorization functionality works correctly without caching
- No missing features or broken workflows

### Recommendation

**Reclassify Permission Caching from P0 to P2**

**Rationale**:
1. **Correctness Priority**: Authorization correctness is more important than performance optimization
2. **MVP Scale**: Current implementation is performant enough for MVP deployment
3. **No Security Risk**: Absence of caching does not introduce security vulnerabilities
4. **Future Optimization**: Caching is a performance optimization, not a functional requirement
5. **Architecture Freeze Compliance**: Removing P0 status aligns with MVP scope constraints

**New Priority**: P2 (Future Optimization)
**Implementation Timing**: Sprint 3C or post-MVP scaling phase
**Trigger for Implementation**: When permission check latency exceeds acceptable thresholds or when scaling to 100+ schools

---

## Security Hardening Checklist (Revised)

### P0 Security Items (Must Complete in Sprint 3.5)

#### 1. Resource-Level Authorization Implementation

**Checklist**:
- [ ] Implement resource ownership validation in service layer
- [ ] Implement school boundary enforcement in repository layer
- [ ] Apply authorization to all domain entities (TP, ATP, Modul Ajar, Assessment, Evidence, Evaluation, Achievement, Narrative Report)
- [ ] Test Teacher cannot access other schools' data
- [ ] Test School Admin can access own school data only
- [ ] Test System Admin can access all schools
- [ ] Test ownership validation for all resource types
- [ ] Audit all cross-school access by System Admin

**Implementation Approach**:
```go
// internal/authorization/resource_authorizer.go
type ResourceAuthorizer struct {
    repo *repository.BaseRepository
}

func (a *ResourceAuthorizer) CanAccessResource(userID, schoolID, resourceID, resourceType string, action string) error {
    // Check RBAC permission first
    if !a.hasPermission(userID, resourceType, action) {
        return ErrUnauthorized
    }
    
    // Check school boundary
    if !a.isInSchoolScope(resourceID, schoolID, resourceType) {
        return ErrCrossSchoolAccess
    }
    
    // Check ownership (if not admin)
    if !a.isSystemAdmin(userID) && !a.isOwner(userID, resourceID, resourceType) {
        return ErrNotOwner
    }
    
    return nil
}
```

### P1 Security Items (Should Complete in Sprint 3.5 if Time Permits)

#### 2. Permission Audit Trail

**Checklist**:
- [ ] Add permission_changes table
- [ ] Log all permission changes with user, timestamp, reason
- [ ] Add audit query endpoint
- [ ] Test audit log accuracy

#### 3. Rate Limiting

**Checklist**:
- [ ] Implement rate limiting middleware
- [ ] Apply different limits per endpoint type
- [ ] Configure limits (e.g., 100 req/min per user)
- [ ] Test rate limiting enforcement
- [ ] Test bypass for admin users

#### 4. Input Validation Consistency

**Checklist**:
- [ ] Standardize validation DTOs across all endpoints
- [ ] Add validation layer between DTO and domain
- [ ] Validate all required fields
- [ ] Validate enum values
- [ ] Validate data types and formats
- [ ] Test validation rejects invalid data

#### 5. Mass Assignment Prevention

**Checklist**:
- [ ] Add DTO mapping layer between request and domain
- [ ] Never bind request directly to domain entities
- [ ] Map only allowed fields from DTO to domain
- [ ] Test mass assignment prevention
- [ ] Test malicious request rejection

### P2 Security Items (Future Optimization)

#### 6. Permission Caching

**Checklist**:
- [ ] Implement Redis cache layer
- [ ] Cache permission checks at role level
- [ ] Add cache invalidation on role/permission changes
- [ ] Test cache hit rate > 80%
- [ ] Test cache invalidation works correctly

**Implementation Timing**: Sprint 3C or post-MVP scaling phase

---

# SECTION 5.5 — Unified Versioning Architecture

## Versioning Scope

This section defines the unified versioning strategy for all versioned domain entities in the MVP scope:

- TP Sets
- ATP Sets
- Modul Ajar Sets
- Assessments
- Evaluations
- Narrative Reports

---

## TP Set Versioning

### Version Trigger
A new TP Set version is created when:
- Content is updated (CP text, learning objectives, etc.)
- Status transitions to APPROVED
- Regeneration is requested with new parameters
- Explicit version creation requested by user

### Version Number Rules
- **Initial Version**: version_no = 1
- **Increment Rule**: version_no increments by 1 for each new version
- **Current Version**: The version with is_current_version = true
- **Historical Preservation**: All previous versions retained with is_current_version = false

### Immutable vs Mutable Fields

**Immutable Fields (Create New Version)**:
- cp_id
- cp_code
- cp_text
- learning_objectives
- time_allocation
- success_criteria

**Mutable Fields (In-Place Update)**:
- status
- generation_reason
- approved_by
- approved_at
- approved_by_name

### Audit Requirements
- **CreatedBy**: user_id of creator
- **CreatedAt**: timestamp of creation
- **PreviousVersionId**: parent_revision_id (null for initial version)
- **VersionReason**: reason for version creation (regeneration, content update, etc.)
- **ApprovalMetadata**: approved_by, approved_at, approval_comments

### Retrieval Rules
- **Get Current Version**: Query WHERE is_current_version = true
- **Get Specific Version**: Query WHERE version_no = X AND tp_set_id = Y
- **Get Version History**: Query WHERE tp_set_id = Y ORDER BY version_no DESC

---

## ATP Set Versioning

### Version Trigger
A new ATP Set version is created when:
- TP Set reference changes
- ATP content is updated (week allocations, CP mappings)
- Status transitions to APPROVED
- Regeneration is requested

### Version Number Rules
- **Initial Version**: version_no = 1
- **Increment Rule**: version_no increments by 1 for each new version
- **Current Version**: The version with is_current_version = true
- **Historical Preservation**: All previous versions retained with is_current_version = false

### Immutable vs Mutable Fields

**Immutable Fields (Create New Version)**:
- tp_set_id
- tp_set_version_no
- atp_items (week allocations, CP mappings)

**Mutable Fields (In-Place Update)**:
- status
- generation_reason
- approved_by
- approved_at

### Audit Requirements
- **CreatedBy**: user_id of creator
- **CreatedAt**: timestamp of creation
- **PreviousVersionId**: parent_revision_id (null for initial version)
- **VersionReason**: reason for version creation
- **ApprovalMetadata**: approved_by, approved_at, approval_comments

### Retrieval Rules
- **Get Current Version**: Query WHERE is_current_version = true
- **Get Specific Version**: Query WHERE version_no = X AND atp_set_id = Y
- **Get Version History**: Query WHERE atp_set_id = Y ORDER BY version_no DESC

---

## Modul Ajar Set Versioning

### Version Trigger
A new Modul Ajar Set version is created when:
- ATP Set reference changes
- Modul Ajar content is updated
- Status transitions to APPROVED
- Regeneration is requested

### Version Number Rules
- **Initial Version**: version_no = 1
- **Increment Rule**: version_no increments by 1 for each new version
- **Current Version**: The version with is_current_version = true
- **Historical Preservation**: All previous versions retained with is_current_version = false

### Immutable vs Mutable Fields

**Immutable Fields (Create New Version)**:
- atp_set_id
- atp_set_version_no
- modul_ajar_items

**Mutable Fields (In-Place Update)**:
- status
- generation_reason
- approved_by
- approved_at

### Audit Requirements
- **CreatedBy**: user_id of creator
- **CreatedAt**: timestamp of creation
- **PreviousVersionId**: parent_revision_id (null for initial version)
- **VersionReason**: reason for version creation
- **ApprovalMetadata**: approved_by, approved_at, approval_comments

### Retrieval Rules
- **Get Current Version**: Query WHERE is_current_version = true
- **Get Specific Version**: Query WHERE version_no = X AND modul_ajar_set_id = Y
- **Get Version History**: Query WHERE modul_ajar_set_id = Y ORDER BY version_no DESC

---

## Assessment Versioning

### Version Trigger
A new Assessment version is created when:
- Assessment items are updated
- Answer key is updated
- Scoring guidelines are updated
- Status transitions to APPROVED

### Version Number Rules
- **Initial Version**: version_no = 1
- **Increment Rule**: version_no increments by 1 for each new version
- **Current Version**: The version with is_current_version = true
- **Historical Preservation**: All previous versions retained with is_current_version = false

### Immutable vs Mutable Fields

**Immutable Fields (Create New Version)**:
- tp_id
- tp_version_no
- success_criteria_snapshot
- tp_learning_objectives_snapshot
- tp_time_allocation_snapshot
- assessment_items
- answer_key
- scoring_guidelines

**Mutable Fields (In-Place Update)**:
- status

### Audit Requirements
- **CreatedBy**: user_id of creator
- **CreatedAt**: timestamp of creation
- **PreviousVersionId**: parent_revision_id (null for initial version)
- **VersionReason**: reason for version creation
- **ApprovalMetadata**: approved_by, approved_at, approval_comments

### Retrieval Rules
- **Get Current Version**: Query WHERE is_current_version = true
- **Get Specific Version**: Query WHERE version_no = X AND assessment_id = Y
- **Get Version History**: Query WHERE assessment_id = Y ORDER BY version_no DESC

---

## Evaluation Versioning

### Version Trigger
A new Evaluation version is created when:
- Teacher feedback is updated
- Total score is updated
- Performance level is updated
- Evaluation is re-submitted for review

### Version Number Rules
- **Initial Version**: revision_no = 1
- **Increment Rule**: revision_no increments by 1 for each new version
- **Current Version**: The version with is_current_version = true
- **Historical Preservation**: All previous versions retained with is_current_version = false
- **Feedback History**: Preserved in evaluation_feedback_history table

### Immutable vs Mutable Fields

**Immutable Fields (Create New Version)**:
- total_score
- performance_level
- teacher_feedback
- evaluation_criteria_scores

**Mutable Fields (In-Place Update)**:
- status (DRAFT → UNDER_REVIEW → APPROVED)

### Audit Requirements
- **CreatedBy**: user_id of creator (teacher)
- **CreatedAt**: timestamp of creation
- **PreviousVersionId**: parent_revision_id (null for initial version)
- **VersionReason**: reason for revision (feedback update, score correction, etc.)
- **ApprovalMetadata**: approved_by, approved_at, approval_comments
- **FeedbackHistory**: All feedback changes tracked in evaluation_feedback_history

### Retrieval Rules
- **Get Current Version**: Query WHERE is_current_version = true
- **Get Specific Version**: Query WHERE revision_no = X AND evaluation_id = Y
- **Get Version History**: Query WHERE evaluation_id = Y ORDER BY revision_no DESC
- **Get Feedback History**: Query evaluation_feedback_history WHERE evaluation_id = Y ORDER BY changed_at DESC

---

## Narrative Report Versioning

### Version Trigger
A new Narrative Report version is created when:
- Report content is updated
- Achievement data is refreshed
- Report is regenerated
- Status transitions to APPROVED

### Version Number Rules
- **Initial Version**: version_no = 1
- **Increment Rule**: version_no increments by 1 for each new version
- **Current Version**: The version with is_current_version = true
- **Historical Preservation**: All previous versions retained with is_current_version = false

### Immutable vs Mutable Fields

**Immutable Fields (Create New Version)**:
- student_id
- report_period
- narrative_content
- achievement_data (snapshot at time of generation)

**Mutable Fields (In-Place Update)**:
- status
- generated_by (if re-generated by different user)

### Audit Requirements
- **CreatedBy**: user_id of creator
- **CreatedAt**: timestamp of creation
- **PreviousVersionId**: parent_revision_id (null for initial version)
- **VersionReason**: reason for version creation (content update, achievement refresh, etc.)
- **ApprovalMetadata**: approved_by, approved_at, approval_comments

### Retrieval Rules
- **Get Current Version**: Query WHERE is_current_version = true
- **Get Specific Version**: Query WHERE version_no = X AND narrative_report_id = Y
- **Get Version History**: Query WHERE narrative_report_id = Y ORDER BY version_no DESC

---

## Cross-Entity Version Consistency

### Version Cascade Rules

When a parent entity is versioned, dependent entities may require versioning:

1. **TP Set Version Change** → May trigger ATP Set versioning
2. **ATP Set Version Change** → May trigger Modul Ajar Set versioning
3. **Assessment Version Change** → Does NOT trigger Evaluation versioning (evaluations reference assessment snapshot)

### Version Reference Integrity

- All versioned entities reference parent entity version (e.g., assessment references tp_version_no)
- Version references are immutable snapshots, not live references
- Historical versions preserve the state of referenced entities at time of creation

---

## Version Cleanup and Retention

### Retention Policy
- **Minimum Retention**: All versions retained for audit purposes
- **Active Retention**: Current version + previous 5 versions readily accessible
- **Archive Retention**: Versions older than 5 may be archived to cold storage
- **Deletion Policy**: Versions never deleted, only archived

### Soft Delete Behavior
- Versions are never hard deleted
- Soft delete marks version as archived (is_archived = true)
- Archived versions remain queryable but with performance impact

---

## Security and Authorization

### Version Access Rules
- **Current Version**: All authorized users can access current version
- **Historical Versions**: Only users with READ permission can access historical versions
- **Version Creation**: Only users with WRITE permission can create new versions
- **Version Deletion**: PROHIBITED (versions never deleted, only archived)

### Version Audit Logging
- All version creations logged with user, timestamp, reason
- All version accesses logged for audit trail
- Version history queries logged for compliance

---

# SECTION 5.6 — Evidence Storage Architecture

## Supported Evidence Types

### Document Types
- **PDF**: application/pdf
- **Word Document**: application/msword, application/vnd.openxmlformats-officedocument.wordprocessingml.document
- **Excel Spreadsheet**: application/vnd.ms-excel, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet
- **PowerPoint Presentation**: application/vnd.ms-powerpoint, application/vnd.openxmlformats-officedocument.presentationml.presentation
- **Plain Text**: text/plain

### Image Types
- **JPEG**: image/jpeg
- **PNG**: image/png
- **GIF**: image/gif
- **WebP**: image/webp
- **SVG**: image/svg+xml

### Audio Types
- **MP3**: audio/mpeg
- **WAV**: audio/wav
- **AAC**: audio/aac
- **OGG**: audio/ogg

### Video Types
- **MP4**: video/mp4
- **WebM**: video/webm
- **MOV**: video/quicktime
- **AVI**: video/x-msvideo

### Validation Rules
- **Maximum File Size**: 50 MB per file
- **Allowed MIME Types**: Only types listed above
- **File Extension Validation**: Extension must match MIME type
- **Content-Type Validation**: Server-side validation of actual file content
- **Malicious File Detection**: Basic file signature validation

---

## Storage Strategy

### Development Environment
- **Storage Type**: Local filesystem
- **Storage Path**: ./storage/evidences/
- **Recommendation**: Acceptable for development only

### Production Environment
- **Storage Type**: S3-compatible object storage
- **Recommended Solutions**:
  - **AWS S3**: Production-grade, highly available
  - **MinIO**: Self-hosted S3-compatible, suitable for on-premise deployment
  - **DigitalOcean Spaces**: Cost-effective S3-compatible
  - **Google Cloud Storage**: Alternative cloud storage
- **Storage Path**: Organized by school_id and evidence_id
  - Example: s3://nusa-evidences/{school_id}/{evidence_id}/{filename}

### Storage Configuration

```go
// internal/storage/evidence_storage.go
type EvidenceStorage struct {
    client     *s3.Client
    bucketName string
}

func NewEvidenceStorage(config StorageConfig) *EvidenceStorage {
    cfg, _ := config.LoadDefaultConfig(context.TODO())
    client := s3.NewFromConfig(cfg)
    return &EvidenceStorage{
        client:     client,
        bucketName: config.BucketName,
    }
}

func (s *EvidenceStorage) StoreFile(evidenceID, schoolID, filename string, file io.Reader) (string, error) {
    key := fmt.Sprintf("%s/%s/%s", schoolID, evidenceID, filename)
    _, err := s.client.PutObject(context.TODO(), &s3.PutObjectInput{
        Bucket: aws.String(s.bucketName),
        Key:    aws.String(key),
        Body:   file,
    })
    return key, err
}

func (s *EvidenceStorage) GetFile(key string) (io.ReadCloser, error) {
    result, err := s.client.GetObject(context.TODO(), &s3.GetObjectInput{
        Bucket: aws.String(s.bucketName),
        Key:    aws.String(key),
    })
    return result.Body, err
}
```

---

## Metadata Model

### Evidence Metadata (Database)

```sql
CREATE TABLE evidences (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    student_id UUID NOT NULL REFERENCES students(id),
    assessment_id UUID NOT NULL REFERENCES assessments(id),
    evidence_type VARCHAR(50) NOT NULL, -- DOCUMENT, IMAGE, VIDEO, AUDIO
    file_id UUID NOT NULL UNIQUE,
    storage_key VARCHAR(500) NOT NULL,
    file_name VARCHAR(255) NOT NULL,
    mime_type VARCHAR(100) NOT NULL,
    file_size_bytes BIGINT NOT NULL,
    file_hash VARCHAR(64) NOT NULL, -- SHA-256 hash
    teacher_notes TEXT,
    status VARCHAR(50) NOT NULL DEFAULT 'UPLOADED',
    uploaded_by UUID NOT NULL REFERENCES users(id),
    uploaded_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);
```

### Evidence Binary Content (Object Storage)

**Storage Key Format**: `{school_id}/{evidence_id}/{filename}`

**Metadata Headers** (stored with object):
- Content-Type
- Content-Length
- x-amz-meta-uploaded-by
- x-amz-meta-uploaded-at
- x-amz-meta-evidence-id
- x-amz-meta-student-id

### Separation Rationale

1. **Metadata in Database**: Enables querying, filtering, and relational operations
2. **Binary Content in Object Storage**: Scalable, cost-effective for large files
3. **Separation of Concerns**: Database manages metadata, storage manages binary data
4. **Performance**: Metadata queries fast, binary content retrieved on-demand

---

## Security Requirements

### File Type Validation

**Multi-Layer Validation**:
1. **Client-Side**: File extension check (user bypassable)
2. **Server-Side**: MIME type validation from file header
3. **Server-Side**: Magic number validation (file signature)
4. **Server-Side**: Content-Type validation against allowed list

**Implementation**:
```go
func ValidateFileType(file io.Reader, filename string) error {
    // Read first 512 bytes for magic number detection
    buffer := make([]byte, 512)
    _, err := file.Read(buffer)
    if err != nil {
        return err
    }

    // Detect MIME type from content
    mimeType := http.DetectContentType(buffer)

    // Validate against allowed types
    if !isAllowedMimeType(mimeType) {
        return ErrInvalidFileType
    }

    // Validate extension matches MIME type
    if !extensionMatchesMime(filename, mimeType) {
        return ErrExtensionMismatch
    }

    return nil
}
```

### Virus Scanning

**Requirement**: Optional for MVP, Recommended for Production

**Implementation Options**:
1. **ClamAV Integration**: Open-source antivirus scanning
2. **Cloud-Based Scanning**: AWS VirusScan, Google Cloud Virus Scanning
3. **Sandbox Execution**: Execute file in isolated environment

**Scanning Workflow**:
1. Upload file to temporary quarantine location
2. Scan file for viruses/malware
3. If clean, move to permanent storage
4. If infected, delete file and notify user

**MVP Approach**: Skip virus scanning for MVP, implement in Sprint 3C or post-MVP

### Download Authorization

**Access Control**:
- **Teacher**: Can download evidence for own students only
- **School Admin**: Can download all school evidence
- **System Admin**: Can download all evidence
- **Student**: Can download own evidence (if feature enabled)

**Signed URL Strategy** (Recommended for Production):
```go
func (s *EvidenceStorage) GeneratePresignedURL(key string, expiration time.Duration) (string, error) {
    presignClient := s3.NewPresignClient(s.client)
    result, _ := presignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
        Bucket: aws.String(s.bucketName),
        Key:    aws.String(key),
    }, s3.WithPresignExpires(expiration))
    return result.URL, nil
}
```

**Benefits**:
- Temporary access with expiration
- No need to proxy downloads through application
- Reduced server load
- Direct download from storage

### Content-Type Validation

**Validation Rules**:
- Validate Content-Type header on upload
- Validate Content-Type header on download
- Prevent MIME type sniffing attacks
- Force download for executable files

---

## Lifecycle Management

### Retention Rules

**Default Retention Policy**:
- **Active Evidence**: Retain indefinitely while student is enrolled
- **Graduated Students**: Retain for 7 years post-graduation
- **Deleted Students**: Retain for 1 year post-deletion
- **Archived Evidence**: Move to cold storage after 3 years

**Configuration**:
```go
type RetentionPolicy struct {
    ActiveRetention       time.Duration // 0 = indefinite
    GraduatedRetention    time.Duration // 7 years
    DeletedRetention      time.Duration // 1 year
    ArchiveThreshold      time.Duration // 3 years
}
```

### Soft Delete Behavior

**Implementation**:
1. Mark evidence as deleted in database (is_deleted = true)
2. Move file to archive storage location
3. Remove from active queries
4. Retain for retention period
5. Permanent delete after retention period expires

**Archive Location**: `{school_id}/archived/{evidence_id}/{filename}`

### Cleanup Strategy

**Automated Cleanup Job**:
1. **Daily**: Scan for evidence past retention period
2. **Weekly**: Delete files from archive storage past retention
3. **Monthly**: Generate retention audit report
4. **Quarterly**: Review and update retention policies

**Cleanup Workflow**:
```go
func CleanupExpiredEvidence(repo *EvidenceRepository, storage *EvidenceStorage) error {
    expired, _ := repo.FindExpiredEvidence()
    for _, evidence := range expired {
        // Delete from storage
        storage.DeleteFile(evidence.StorageKey)
        // Mark as permanently deleted
        repo.PermanentlyDelete(evidence.ID)
    }
    return nil
}
```

---

## Storage Migration Strategy

### Migration from Local to Object Storage

**Phase 1: Dual Write**
- Write to both local storage and object storage
- Read from local storage (primary), object storage (fallback)

**Phase 2: Switch Read**
- Write to both local storage and object storage
- Read from object storage (primary), local storage (fallback)

**Phase 3: Object Storage Only**
- Write to object storage only
- Read from object storage only
- Remove local storage dependency

**Phase 4: Cleanup**
- Delete files from local storage
- Remove local storage code

---

## Error Handling and Recovery

### Upload Failure Handling

**Retry Strategy**:
- **Transient Errors**: Retry up to 3 times with exponential backoff
- **Permanent Errors**: Return error to user with clear message
- **Network Errors**: Retry with backoff, timeout after 30 seconds

**Partial Upload Handling**:
- Use multipart upload for large files
- Track upload progress
- Allow resume of interrupted uploads
- Clean up partial uploads on failure

### Storage Unavailability

**Fallback Strategy**:
- **Primary Storage**: Object storage
- **Fallback Storage**: Local filesystem (temporary)
- **Graceful Degradation**: Reject uploads if both unavailable
- **Alerting**: Notify operations team of storage issues

---

## Monitoring and Metrics

### Storage Metrics

**Business Metrics**:
- Evidence uploads per day
- Evidence downloads per day
- Storage usage per school
- Average file size
- File type distribution

**System Metrics**:
- Upload success rate
- Download success rate
- Storage latency (p50, p95, p99)
- Storage error rate
- Storage capacity utilization

### Alerting

**Alert Conditions**:
- Storage error rate > 5%
- Upload latency p95 > 10 seconds
- Download latency p95 > 5 seconds
- Storage capacity > 80%
- Storage service unavailable

---

# SECTION 5.7 — Observability and Operational Readiness

## Structured Logging

### Logging Requirements

**Mandatory Logging Events**:
- Authentication events (login, logout, token refresh)
- Authorization failures (permission denied, access denied)
- Workflow transitions (status changes, approvals)
- Assessment approvals
- Evidence uploads
- Evaluation revisions
- Report generation
- Data modifications (create, update, delete operations)
- Status transitions
- Approval actions
- Permission changes

### Log Format

**Structured JSON Format**:
```json
{
  "timestamp": "2026-06-09T10:30:00Z",
  "level": "INFO",
  "service": "assessment-service",
  "event": "assessment.approved",
  "user_id": "uuid",
  "school_id": "uuid",
  "assessment_id": "uuid",
  "correlation_id": "uuid",
  "duration_ms": 150,
  "metadata": {
    "previous_status": "UNDER_REVIEW",
    "new_status": "APPROVED",
    "approver_id": "uuid"
  }
}
```

### Required Log Fields

| Field | Type | Description |
|-------|------|-------------|
| timestamp | string | ISO 8601 timestamp |
| level | string | DEBUG, INFO, WARN, ERROR, FATAL |
| service | string | Service name (e.g., assessment-service) |
| event | string | Event type (e.g., assessment.approved) |
| user_id | string | User UUID (if applicable) |
| school_id | string | School UUID (if applicable) |
| correlation_id | string | Request correlation ID for tracing |
| duration_ms | number | Operation duration in milliseconds |
| metadata | object | Event-specific metadata |

### Log Levels

- **DEBUG**: Detailed diagnostic information
- **INFO**: Normal operational events
- **WARN**: Warning conditions that may require attention
- **ERROR**: Error events that do not prevent operation
- **FATAL**: Critical errors that prevent operation

---

## Audit Logging

### Audit Requirements

**Mandatory Audit Events**:
- Data modifications (create, update, delete)
- Status transitions
- Approval actions
- Permission changes
- Cross-school access by System Admin
- Authentication failures (after 3 consecutive failures)
- Authorization failures
- Configuration changes

### Audit Log Format

```json
{
  "audit_id": "uuid",
  "timestamp": "2026-06-09T10:30:00Z",
  "actor_id": "uuid",
  "actor_role": "TEACHER",
  "school_id": "uuid",
  "action": "UPDATE",
  "resource_type": "assessment",
  "resource_id": "uuid",
  "changes": {
    "status": {
      "from": "UNDER_REVIEW",
      "to": "APPROVED"
    }
  },
  "reason": "Assessment approved for use",
  "ip_address": "192.168.1.100",
  "user_agent": "Mozilla/5.0..."
}
```

### Audit Log Retention

- **Retention Period**: 7 years minimum
- **Storage**: Separate audit database or dedicated audit log storage
- **Immutability**: Audit logs must be immutable (append-only)
- **Access**: Audit logs accessible only to System Admin and auditors

---

## Metrics

### Business Metrics

**Assessment Metrics**:
- Assessments created (count per day)
- Assessments approved (count per day)
- Assessments rejected (count per day)
- Average approval time (duration)

**Evaluation Metrics**:
- Evaluations completed (count per day)
- Evaluations revised (count per day)
- Average evaluation time (duration)
- Evaluation revision rate (percentage)

**Report Metrics**:
- Reports generated (count per day)
- Report generation time (duration p50, p95, p99)
- Report generation failures (count per day)

**Evidence Metrics**:
- Evidence uploads (count per day)
- Evidence downloads (count per day)
- Average file size (bytes)
- Storage usage per school (bytes)

### System Metrics

**Request Metrics**:
- Request count (per endpoint, per status code)
- Request latency (p50, p95, p99)
- Error rate (percentage)
- Throughput (requests per second)

**Database Metrics**:
- Query duration (p50, p95, p99)
- Connection pool utilization (percentage)
- Query count (per query type)
- Slow queries (count, duration > 1s)

**Cache Metrics** (if implemented):
- Cache hit rate (percentage)
- Cache miss rate (percentage)
- Cache eviction count
- Cache size (bytes)

**External Service Metrics**:
- AI service call duration (p50, p95, p99)
- AI service error rate (percentage)
- Storage operation duration (p50, p95, p99)
- Storage error rate (percentage)

### Metrics Collection

**Recommended Tools**:
- **Prometheus**: Metrics collection and storage
- **Grafana**: Metrics visualization and dashboards
- **OpenTelemetry**: Metrics instrumentation

**Implementation**:
```go
import "github.com/prometheus/client_golang/prometheus"

var (
    assessmentsCreated = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "assessments_created_total",
            Help: "Total number of assessments created",
        },
        []string{"school_id"},
    )

    assessmentApprovalDuration = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name: "assessment_approval_duration_seconds",
            Help: "Duration of assessment approval process",
        },
        []string{"school_id"},
    )
)
```

---

## Health Checks

### Liveness Check

**Purpose**: Determine if the application is running
**Endpoint**: GET /health/live
**Response**: 200 OK if application is running

```json
{
  "status": "healthy",
  "timestamp": "2026-06-09T10:30:00Z"
}
```

### Readiness Check

**Purpose**: Determine if the application is ready to handle requests
**Endpoint**: GET /health/ready
**Response**: 200 OK if all dependencies are healthy

```json
{
  "status": "ready",
  "timestamp": "2026-06-09T10:30:00Z",
  "checks": {
    "database": "healthy",
    "redis": "healthy",
    "object_storage": "healthy",
    "ai_service": "healthy"
  }
}
```

### Dependency Checks

**PostgreSQL Check**:
- Connection established
- Query execution successful
- Connection pool available

**Redis Check** (if used):
- Connection established
- SET/GET operation successful

**Object Storage Check**:
- Connection established
- Bucket access successful
- PUT/GET operation successful

**AI Service Check**:
- Connection established
- Health endpoint accessible
- API key valid

### Health Check Configuration

```go
// internal/health/health.go
type HealthChecker struct {
    db           *sql.DB
    redis        *redis.Client
    storage      *EvidenceStorage
    aiService    *AIService
}

func (h *HealthChecker) CheckLiveness() error {
    // Check if application is running
    return nil
}

func (h *HealthChecker) CheckReadiness() map[string]string {
    checks := make(map[string]string)

    // Check database
    if err := h.db.Ping(); err != nil {
        checks["database"] = "unhealthy"
    } else {
        checks["database"] = "healthy"
    }

    // Check Redis
    if h.redis != nil {
        if err := h.redis.Ping(context.Background()).Err(); err != nil {
            checks["redis"] = "unhealthy"
        } else {
            checks["redis"] = "healthy"
        }
    }

    // Check object storage
    if err := h.storage.CheckHealth(); err != nil {
        checks["object_storage"] = "unhealthy"
    } else {
        checks["object_storage"] = "healthy"
    }

    // Check AI service
    if err := h.aiService.CheckHealth(); err != nil {
        checks["ai_service"] = "unhealthy"
    } else {
        checks["ai_service"] = "healthy"
    }

    return checks
}
```

---

## Tracing

### Distributed Tracing Requirements

**Trace Scopes**:
- API request flow (from ingress to egress)
- Service execution (service method calls)
- Database operations (query execution)
- External integrations (AI service, object storage)

### Trace Context Propagation

**Trace Headers**:
- `traceparent`: W3C trace context format
- `X-Correlation-ID`: Custom correlation ID

**Implementation**:
```go
import "go.opentelemetry.io/otel"

func (h *AssessmentHandler) ApproveAssessment(c *gin.Context) {
    ctx := c.Request.Context()
    tracer := otel.Tracer("assessment-service")

    ctx, span := tracer.Start(ctx, "ApproveAssessment")
    defer span.End()

    span.SetAttributes(
        attribute.String("assessment.id", assessmentID),
        attribute.String("user.id", userID),
    )

    // Service call
    err := h.service.ApproveAssessment(ctx, assessmentID, userID)
    if err != nil {
        span.RecordError(err)
        span.SetStatus(codes.Error, err.Error())
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    span.SetStatus(codes.Ok, "Assessment approved")
    c.JSON(200, gin.H{"status": "approved"})
}
```

### Span Attributes

**Required Attributes**:
- `service.name`: Service name
- `operation.name`: Operation name
- `user.id`: User UUID (if applicable)
- `school.id`: School UUID (if applicable)
- `resource.id`: Resource UUID (if applicable)

**Optional Attributes**:
- `db.statement`: SQL query (sanitized)
- `db.system`: Database type (PostgreSQL)
- `http.method`: HTTP method
- `http.url`: HTTP URL
- `http.status_code`: HTTP status code

### Tracing Tools

**Recommended Tools**:
- **OpenTelemetry**: Tracing instrumentation
- **Jaeger**: Distributed tracing platform
- **Zipkin**: Distributed tracing system

---

## Operational Alerts

### Alert Conditions

**Authentication Alerts**:
- Authentication failure rate > 10% (5-minute window)
- Consecutive authentication failures > 5 from same IP
- Token refresh failure rate > 5%

**Authorization Alerts**:
- Authorization failure rate > 5% (5-minute window)
- Cross-school access attempt by non-admin
- Permission escalation attempt

**Application Alerts**:
- Error rate > 5% (5-minute window)
- Request latency p95 > 2 seconds (5-minute window)
- Request rate > 1000 req/sec (5-minute window)

**Database Alerts**:
- Database connection pool utilization > 80%
- Slow query count > 10 per minute (duration > 1s)
- Database connection failures > 5 per minute

**Storage Alerts**:
- Storage error rate > 5% (5-minute window)
- Upload latency p95 > 10 seconds (5-minute window)
- Storage capacity > 80%

**AI Service Alerts**:
- AI service error rate > 10% (5-minute window)
- AI service latency p95 > 30 seconds (5-minute window)
- AI service unavailable

### Alert Severity Levels

**Critical**:
- Application down (health check failing)
- Database unavailable
- Storage unavailable
- Data corruption risk

**Warning**:
- High error rate
- High latency
- Resource utilization high
- Degraded performance

**Info**:
- Scheduled maintenance
- Configuration changes
- Deployment events

### Alert Notification Channels

**Channels**:
- **Email**: For critical alerts
- **Slack/Teams**: For warning and critical alerts
- **PagerDuty/Opsgenie**: For critical alerts (on-call rotation)
- **SMS**: For critical alerts only

---

## Monitoring Dashboard Requirements

### Required Dashboards

**System Overview Dashboard**:
- Request rate (per minute)
- Error rate (percentage)
- Request latency (p50, p95, p99)
- Database connection pool utilization
- Storage error rate
- AI service availability

**Business Metrics Dashboard**:
- Assessments created (per day)
- Evaluations completed (per day)
- Reports generated (per day)
- Evidence uploads (per day)
- Average approval time
- Average evaluation time

**Security Dashboard**:
- Authentication failures (per hour)
- Authorization failures (per hour)
- Cross-school access attempts
- Permission changes (per day)
- Suspicious activity alerts

**Database Dashboard**:
- Query duration (p50, p95, p99)
- Slow queries (count, duration)
- Connection pool utilization
- Query count (per query type)
- Database size

---

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

# SECTION 6 — Frontend Dependency-Based Execution Plan

## Frontend Module Dependency Analysis

| Module | Backend Dependency Status | Prerequisite Endpoints | Blocking Endpoints | Implementation Strategy |
|--------|--------------------------|----------------------|-------------------|------------------------|
| **Authentication** | ✅ AVAILABLE | Existing auth endpoints | None | Full implementation |
| **Layout** | ✅ NO DEPENDENCY | None | None | Full implementation |
| **Navigation** | ✅ NO DEPENDENCY | None | None | Full implementation (can stub incomplete items) |
| **State Management** | ✅ NO DEPENDENCY | None | None | Full implementation |
| **API Client** | ✅ NO DEPENDENCY | None | None | Create base client, stub methods |
| **TP Workspace** | ⚠️ PARTIAL | GET /tp-sets, GET /tp-sets/:id | PUT /tp-sets/:id, GET /tp-sets/:id/versions | Implement list/detail, stub edit/versioning |
| **ATP Workspace** | ⚠️ PARTIAL | GET /atp-sets | GET /atp-sets/:id | Implement list, stub detail |
| **Modul Ajar Workspace** | ⚠️ PARTIAL | GET /modul-ajar-sets | GET /modul-ajar-sets/:id | Implement list, stub detail |
| **Assessment Workspace** | ⚠️ PARTIAL | GET /assessments, GET /assessments/:id | PUT /assessment/:id, POST /assessment/:id/approve | Implement list/detail, stub edit/approve |
| **Evidence Workspace** | ⚠️ PARTIAL | GET /evidences | POST /evidences/upload, GET /evidences/:id | Implement list, stub upload/detail |
| **Progress Dashboard** | ✅ AVAILABLE | Achievement endpoints | None | Full implementation |
| **Narrative Report Workspace** | ✅ AVAILABLE | Report endpoints | None | Full implementation |

## Frontend Implementation Sequence

### Phase 1: Foundation Layer (No Backend Dependencies)

**Prerequisites**: None
**Readiness Criteria**: None

**Implementation Items**:
- Authentication integration (using existing endpoints)
- Layout components (Header, Sidebar, Footer)
- Navigation menu (can stub incomplete items)
- State management setup (TanStack Query, Zustand)
- API client library (create base client, stub methods)

**Completion Criteria**:
- Authentication flow functional
- Layout components render correctly
- Navigation structure defined
- State management configured
- API client base structure established

### Phase 2: Independent Modules (Backend Dependencies Available)

**Prerequisites**: Phase 1 complete
**Backend Dependencies**: Achievement endpoints, Report endpoints

**Implementation Items**:
- Progress Dashboard (full implementation)
- Narrative Report Workspace (full implementation)

**Completion Criteria**:
- Progress Dashboard displays correct data
- Narrative Report Workspace functional
- API integration validated

### Phase 3: Partial Implementation (Backend Dependencies Partial)

**Prerequisites**: Phase 1 complete, Phase 2 complete
**Backend Dependencies**: List endpoints available, detail/update endpoints pending

**Implementation Items**:
- TP Workspace (list/detail only, stub edit/versioning)
- ATP Workspace (list only, stub detail)
- Modul Ajar Workspace (list only, stub detail)
- Assessment Workspace (list/detail only, stub edit/approve)
- Evidence Workspace (list only, stub upload/detail)

**Completion Criteria**:
- List views functional for all modules
- Detail views functional where endpoints available
- Stub components in place for blocked features
- Error handling for missing endpoints

### Phase 4: Complete Implementation (All Backend Dependencies Available)

**Prerequisites**: Phase 3 complete, all P0 backend endpoints deployed
**Backend Dependencies**: All Sprint 3.5 endpoints

**Implementation Items**:
- TP Workspace edit/versioning (when PUT /tp-sets/:id and GET /tp-sets/:id/versions complete)
- ATP Workspace detail (when GET /atp-sets/:id complete)
- Modul Ajar Workspace detail (when GET /modul-ajar-sets/:id complete)
- Assessment Workspace edit/approve (when PUT /assessment/:id and POST /assessment/:id/approve complete)
- Evidence Workspace upload/detail (when POST /evidences/upload and GET /evidences/:id complete)

**Completion Criteria**:
- All stubbed features implemented
- All endpoints integrated
- End-to-end workflows functional
- Integration tests passing

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
| **Authentication** | ✅ YES | ✅ YES | ✅ YES | READY |
| **TP Workspace - List** | ✅ YES | ✅ YES | ✅ YES | READY |
| **TP Workspace - Detail** | ✅ YES | ✅ YES | ✅ YES | READY |
| **TP Workspace - Edit** | ❌ NO | ⏸️ WAITING | ❌ NO | BLOCKED - Backend |
| **TP Workspace - Version History** | ❌ NO | ⏸️ WAITING | ❌ NO | BLOCKED - Backend |
| **ATP Workspace - List** | ✅ YES | ✅ YES | ✅ YES | READY |
| **ATP Workspace - Detail** | ❌ NO | ⏸️ WAITING | ❌ NO | BLOCKED - Backend |
| **Modul Ajar Workspace - List** | ✅ YES | ✅ YES | ✅ YES | READY |
| **Modul Ajar Workspace - Detail** | ❌ NO | ⏸️ WAITING | ❌ NO | BLOCKED - Backend |
| **Assessment Workspace - List** | ✅ YES | ✅ YES | ✅ YES | READY |
| **Assessment Workspace - Detail** | ✅ YES | ✅ YES | ✅ YES | READY |
| **Assessment Workspace - Edit** | ❌ NO | ⏸️ WAITING | ❌ NO | BLOCKED - Backend |
| **Assessment Workspace - Approve** | ❌ NO | ⏸️ WAITING | ❌ NO | BLOCKED - Backend |
| **Evidence Workspace - List** | ✅ YES | ✅ YES | ✅ YES | READY |
| **Evidence Workspace - Upload** | ❌ NO | ⏸️ WAITING | ❌ NO | BLOCKED - Backend |
| **Evidence Workspace - Detail** | ❌ NO | ⏸️ WAITING | ❌ NO | BLOCKED - Backend |
| **Progress Dashboard** | ✅ YES | ✅ YES | ✅ YES | READY |
| **Narrative Report Workspace** | ✅ YES | ✅ YES | ✅ YES | READY |

**Integration Readiness Score**: 12/21 features (57%) initially, 21/21 (100%) after Sprint 3.5 completion

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
  - [ ] Resource-level authorization architecture implemented
- [ ] All P1 security enhancements implemented (if time permits):
  - [ ] Permission audit trail implemented
  - [ ] Rate limiting implemented
  - [ ] Input validation consistency implemented
  - [ ] Mass assignment prevention implemented
- [ ] P2 security enhancements (future optimization):
  - [ ] Permission caching (deferred to Sprint 3C or post-MVP)

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

# SECTION 10 — Sprint 3B Technical Readiness

## GO Criteria

Sprint 3B may proceed with **FULL SCOPE** frontend implementation if:

**Backend Readiness**:
- [ ] All 8 P0 backend endpoints implemented and deployed
- [ ] All 5 P0/P1 defects resolved and deployed
- [ ] All P0 security enhancements implemented and deployed
- [ ] Backend API stability demonstrated (no breaking changes)
- [ ] Resource authorization architecture implemented
- [ ] Evidence storage architecture implemented
- [ ] Observability infrastructure configured

**Frontend Readiness**:
- [ ] Frontend foundation layer complete (auth, layout, navigation, state management, API client)
- [ ] Core workflow screens implemented where backend ready (Progress Dashboard, Narrative Report)
- [ ] Frontend team ready to implement remaining screens without technical blockers

**Integration Readiness**:
- [ ] Integration tests passing between completed frontend and backend components
- [ ] API contract validation complete
- [ ] End-to-end workflow validation complete for available features
- [ ] Multi-school isolation validated

**Quality Readiness**:
- [ ] No P0 bugs remaining
- [ ] P1 bugs documented with workarounds
- [ ] Performance benchmarks met for completed features
- [ ] Security audit passed
- [ ] All P0 security tests passing

**Operational Readiness**:
- [ ] Structured logging implemented
- [ ] Audit logging implemented
- [ ] Metrics collection configured
- [ ] Health checks implemented
- [ ] Distributed tracing configured
- [ ] Operational alerts configured

**Stakeholder Readiness**:
- [ ] Product Owner approves Sprint 3B scope
- [ ] Technical lead approves architectural alignment
- [ ] QA lead approves test readiness
- [ ] Operations team approves operational readiness

## NO-GO Criteria

Sprint 3B **MUST NOT** proceed if ANY of the following exist:

**Backend Blockers**:
- [ ] Any P0 backend endpoint not implemented
- [ ] Any P0 defect not resolved
- [ ] Any P0 security enhancement not implemented
- [ ] Breaking API changes required
- [ ] Database migration issues
- [ ] Resource authorization not implemented
- [ ] Evidence storage not implemented

**Integration Blockers**:
- [ ] Frontend and backend API contract mismatch
- [ ] Critical integration test failures
- [ ] End-to-end workflow failures
- [ ] Authentication/authorization integration failures
- [ ] Multi-school isolation not validated

**Quality Blockers**:
- [ ] P0 security vulnerabilities
- [ ] Critical performance issues (p95 > 2s)
- [ ] Data corruption risks
- [ ] Migration rollback failures

**Operational Blockers**:
- [ ] Observability infrastructure not configured
- [ ] Health checks not implemented
- [ ] Audit logging not implemented
- [ ] Operational alerts not configured

**Stakeholder Blockers**:
- [ ] Product Owner does not approve scope
- [ ] Technical lead identifies architectural violations
- [ ] QA lead identifies insufficient test coverage
- [ ] Operations team identifies operational gaps

## Technical Risk Matrix

| Risk | Probability | Impact | Severity | Mitigation |
|------|-------------|--------|----------|------------|
| **Defect fix complexity underestimated** | MEDIUM | HIGH | HIGH | Prioritize P0 only, defer P1/P2 to Sprint 3B |
| **Security hardening blocks progress** | LOW | MEDIUM | MEDIUM | Implement minimum viable security, defer enhancements |
| **Frontend-backend integration issues** | MEDIUM | MEDIUM | MEDIUM | Daily syncs, contract validation, early integration testing |
| **Architecture Freeze violation** | LOW | CRITICAL | CRITICAL | Strict scope enforcement, no new patterns |
| **Resource authorization complexity** | MEDIUM | HIGH | HIGH | Implement repository-level enforcement first, service-level second |
| **Evidence storage migration complexity** | LOW | MEDIUM | MEDIUM | Use local storage for MVP, plan object storage migration |
| **Observability gaps** | LOW | MEDIUM | MEDIUM | Implement minimum viable observability (logging, metrics, health checks) |

## Recommended Decision

### DECISION: **CONDITIONAL GO** for Sprint 3B Full Scope Implementation

**Rationale**:

1. **Sprint 3.5 Success Criteria Clearly Defined**: Dependency-based execution plan ensures technical prerequisites are met before Sprint 3B begins

2. **P0 Focus Maintains Alignment**: Strict limitation to P0 items only (8 endpoints, 5 defects, resource authorization, evidence storage) respects Architecture Freeze scope

3. **Parallel Work Reduces Dependencies**: Frontend foundation work can proceed immediately based on dependency analysis, reducing technical dependencies

4. **Risk Mitigation Strategy**: Clear GO/NO-GO criteria, technical risk matrix, and dependency mapping provides structure for decision-making

5. **Architecture Compliance**: All work stays within MVP scope as defined by Architecture Freeze (no CQRS, no event sourcing, no future-state patterns)

6. **Operational Readiness**: Observability and operational readiness requirements ensure Sprint 3B can proceed with proper monitoring and alerting

**Conditional Requirements**:

Sprint 3B may proceed with FULL scope implementation only if:

- **All Sprint 3.5 Exit Criteria are met** (see Section 9)
- **No P0 blockers remain** (see NO-GO Criteria above)
- **Stakeholder approval is obtained** (Product Owner, Technical Lead, QA Lead, Operations Team)

**Technical Readiness Prerequisites**:

- All P0 backend endpoints deployed and stable
- Resource authorization architecture implemented and tested
- Evidence storage architecture implemented (local storage acceptable for MVP)
- Observability infrastructure configured (logging, metrics, health checks)
- All P0 security enhancements implemented
- Multi-school isolation validated

**Contingency Plan**:

If Sprint 3.5 has unresolved P0 blockers:
- Sprint 3B scope reduced to partial implementation (only features with completed backend)
- Prioritize core workflow screens over advanced features
- Defer blocked features to Sprint 3C

If Sprint 3.5 completes with all P0 items:
- Sprint 3B proceeds with full scope
- Use any additional time for P1 security enhancements
- Begin Sprint 3C planning and test preparation

---

**Document Status**: READY FOR EXECUTION
**Next Action**: Present to stakeholders for approval
**Approval Required**: Product Owner, Technical Lead, QA Lead, Project Manager

**Generated**: 2026-06-09
**Generated By**: Principal Software Architect, Principal Backend Architect, Principal Frontend Architect, Principal DDD Architect, Principal Product Delivery Manager, Principal QA Architect

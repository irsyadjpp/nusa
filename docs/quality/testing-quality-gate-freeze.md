# NUSA Platform — Testing & Quality Gate Freeze v1

**Version**: 1.0  
**Date**: June 2026  
**Status**: LOCKED — DEFINITION OF "DONE" FOR AI IMPLEMENTATION  
**Parent Documents**:
- [`NUSA_ARCHITECTURE_FREEZE_V2.md`](NUSA_ARCHITECTURE_FREEZE_V2.md)
- [`NUSA_DATABASE_SCHEMA_FREEZE_V1.md`](NUSA_DATABASE_SCHEMA_FREEZE_V1.md)
- [`SPRINT_3.5_EXECUTION_PACKAGE.md`](SPRINT_3.5_EXECUTION_PACKAGE.md)

---

# Testing Philosophy

## Purpose

Testing is the **enforcement mechanism** for Architecture Freeze v2 and Database Schema Freeze v1. Code is not "done" when it compiles — it is done when measurable quality gates pass.

## Core Principles

| Principle | Rule |
|-----------|------|
| **Architecture-first** | Tests validate bounded contexts, aggregate invariants, and school isolation — not implementation convenience |
| **Contract-driven** | API behavior is validated against frozen contracts; undocumented behavior is a defect |
| **Fail fast** | CI must block merge on any P0 test failure |
| **No placeholder tests** | `assert.True(t, true)` and log-only tests are **forbidden** in merged code |
| **Real persistence** | Integration tests that claim persistence validation must use a real PostgreSQL test database |
| **Security is functional** | RBAC and tenant isolation are tested as behavior, not assumed from middleware presence |
| **Revision integrity** | Evaluation and artifact versioning tests are mandatory — in-place mutation of approved/revisioned data is P0 |
| **MVP scope** | No CQRS, projection, event-store, or load-test infrastructure in this freeze |

## Test Pyramid (MVP)

```text
        ┌─────────────┐
        │  E2E (8)    │  Critical user journeys — Playwright (target)
        ├─────────────┤
        │ Integration │  Every Sprint 3.5 endpoint × 5 scenarios
        │   (~120+)   │  + defect workflows + migration tests
        ├─────────────┤
        │    Unit     │  Domain, service, repository, middleware
        │  (80% new)  │
        └─────────────┘
```

## Definition of DONE (Universal)

A work item is **DONE** only when ALL of the following are true:

1. Unit tests exist for changed packages at or above coverage threshold
2. Integration tests exist for every new or modified HTTP endpoint (5 mandatory scenarios each)
3. Security tests pass for RBAC, resource ownership, and school isolation
4. OpenAPI spec updated and contract validation passes
5. No open P0 or P1 bugs in scope
6. CI pipeline green (`go test`, lint, migration up/down on test DB)
7. AI code validation checklist completed (see §AI Generated Code Validation)

---

# Quality Gates

## Gate G0 — Pre-Implementation (Design)

| Criterion | Pass Condition | Blocker If |
|-----------|----------------|------------|
| Architecture reference | PR cites Architecture Freeze v2 section | None cited |
| API contract | Endpoint defined in Part 3.3 or Sprint 3.5 package | Ambiguous request/response |
| DB schema | Table/column exists in Schema Freeze v1 | New column without Board approval |
| Test plan | Integration test IDs listed in PR description | No test IDs |

**Status**: Must pass before coding begins.

## Gate G1 — Unit Test Gate

| Criterion | Pass Condition | Measurement |
|-----------|----------------|-------------|
| New/changed packages | ≥ **80%** line coverage | `go test -coverprofile=coverage.out ./...` |
| Domain validators | ≥ **90%** line coverage | Per-package report |
| Zero skipped security unit tests | No `t.Skip` in `*_security_*` tests without Board waiver | Grep CI check |
| All tests pass | `go test ./...` exit 0 | CI job `backend-test` |

## Gate G2 — Integration Test Gate

| Criterion | Pass Condition | Measurement |
|-----------|----------------|-------------|
| Sprint 3.5 endpoints | 5 scenarios × 8 endpoints = **40 tests** minimum | Test file inventory |
| Defect fixes | 5 defect test suites (DEF-001–009) | Named test functions |
| Education route regression | School isolation on all school-scoped routes | `TestSchoolIsolation_*` |
| Database used | Tests use `backend/tests/integration` with test DB | No httptest-only persistence claims |
| All pass | `go test ./tests/integration/...` exit 0 | CI with `TEST_DATABASE_URL` |

## Gate G3 — Security Gate

| Criterion | Pass Condition | Measurement |
|-----------|----------------|-------------|
| RBAC enforcement | Wrong role → HTTP **403** on protected admin routes | `authorization_security_test.go` |
| Unauthenticated | No token → HTTP **401** | Auth integration tests |
| School isolation | Cross-school resource access → HTTP **404** (not 403) | Multitenancy tests |
| JWT validation | Expired/tampered token → **401** | `jwt_security_test.go` |
| SQL injection | Parameterized queries — injection payloads return safe errors | `sql_injection_test.go` |
| Mass assignment | Extra JSON fields do not modify protected columns | Per-endpoint tests |
| P0 security items | Resource instance-level authorization implemented | Manual + automated checklist |

## Gate G4 — Migration Gate

| Criterion | Pass Condition | Measurement |
|-----------|----------------|-------------|
| Forward migration | `000001`–`000008` apply cleanly on empty DB | Migration CI job |
| Rollback migration | Each `.down.sql` reverses `.up.sql` without error | Paired test |
| Data preservation | Post-migration validation queries pass (000003, 000004) | SQL assertions in test |
| Schema drift | `information_schema` matches Schema Freeze v1 | Schema diff script |
| UUID consistency | All PKs use `gen_uuid_v7()` | No `gen_random_uuid()` in new migrations |

## Gate G5 — API Contract Gate

| Criterion | Pass Condition | Measurement |
|-----------|----------------|-------------|
| OpenAPI sync | `backend/docs/api/openapi.yaml` includes all Sprint 3.5 endpoints | Diff review |
| Response envelope | `{ success, data/error, timestamp }` on all endpoints | Contract tests |
| Error codes | Frozen error codes from Architecture Freeze v2 §3.2 | Assert `error.code` |
| Status codes | Match frozen mapping (401/403/404/409/422) | Per-scenario assertions |
| Backward compatibility | Existing clients: no removed fields without version bump | Breaking change review |

## Gate G6 — Release Gate

See §Release Gate. **All G1–G5 must pass.** Any P0 bug blocks release.

---

# Unit Testing Standards

## General Rules (All Layers)

| Rule | Requirement |
|------|-------------|
| Framework | Go `testing` + `testify/assert` + `testify/require` |
| File naming | `{source}_test.go` in same package or `backend/tests/unit/` |
| Table-driven | Required for validation matrices and RBAC permutations |
| Mocking | Interfaces mocked at layer boundary only (service mocks repo, not DB in unit tests) |
| Parallel | `t.Parallel()` allowed for pure domain tests; forbidden for DB tests |
| Short mode | `testing.Short()` may skip integration only — **not** unit security tests |

### Forbidden Shortcuts

| Forbidden | Why |
|-----------|-----|
| `assert.True(t, true)` placeholder | False confidence |
| Testing only happy path | Misses validation and auth failures |
| Unit testing through HTTP for business rules | Belongs in service/domain tests |
| Copy-paste tests without scenario name | Untraceable failures |
| `@todo test later` in merged code | Gate violation |
| Skipping revision/version assertions | Caused DEF-001, DEF-004 |
| Testing repository with production DB | Must use test database or sqlmock with query assertions |

---

## Domain Layer

**Location**: `backend/internal/domain/*_test.go`

| Category | Mandatory Tests | Coverage |
|----------|-----------------|----------|
| Value object validation | `KKTPCriteria.Validate()` threshold ordering, empty competencies | ≥ 90% |
| Workflow transitions | Valid/invalid `WorkflowStatus` transitions | ≥ 90% |
| Permission matrix | `HasPermission()` for all roles × resources × actions | ≥ 90% |
| Enum constants | AssessmentType, EvidenceType, PerformanceLevel boundaries | ≥ 90% |
| Invariant documentation | Test name references aggregate invariant ID | Required |

**Example test IDs**: `TestKKTPCriteria_Validate_ThresholdOrder`, `TestHasPermission_TeacherCannotApprove`

---

## Service Layer

**Location**: `backend/internal/service/*_test.go`, `backend/modules/*/service_test.go`

| Category | Mandatory Tests | Coverage |
|----------|-----------------|----------|
| Business rule enforcement | Cannot approve DRAFT without prerequisites | ≥ 80% |
| School boundary | `SchoolBoundaryValidator` rejects cross-tenant access | ≥ 80% |
| Transaction behavior | Rollback on mid-operation failure (mock repo error) | ≥ 80% |
| Versioning logic | New row created; `is_current_version` flipped | ≥ 80% |
| Snapshot immutability | Assessment snapshot not updated after TP change | ≥ 80% |
| Evaluation revision | Update creates new revision; increments `revision_no` | ≥ 80% |
| Achievement integration | `RefreshReportAchievement` populates `achievement_data` | ≥ 80% |
| Error mapping | Domain errors return correct sentinel types | ≥ 80% |

**Forbidden**: Service tests that call real HTTP or real PostgreSQL (use integration tests).

---

## Repository Layer

**Location**: `backend/tests/integration/repository_test.go` (integration-style) or sqlmock unit tests

| Category | Mandatory Tests | Coverage |
|----------|-----------------|----------|
| CRUD operations | Create, read, update for each aggregate repo | ≥ 60% unit / integration |
| FK constraints | Cascade behavior matches Schema Freeze delete matrix | Integration |
| Unique constraints | Duplicate `(cp_id, version_no)` → error | Integration |
| Version queries | `is_current_version = true` filter correct | Integration |
| Revision queries | `ORDER BY revision_no DESC` returns correct chain | Integration |
| JSONB persistence | `success_criteria`, `evidence_data` round-trip | Integration |
| Parameterized SQL | No string concatenation in queries | Static review + sql injection test |

---

## Middleware Layer

**Location**: `backend/internal/middleware/*_test.go`, `backend/tests/unit/middleware_test.go`

| Category | Mandatory Tests | Coverage |
|----------|-----------------|----------|
| AuthMiddleware | Valid JWT sets context; missing/invalid → 401 | ≥ 85% |
| RequirePermission | Has permission → pass; missing → 403 | ≥ 85% |
| RequireRole | Role match; mismatch → 403 | ≥ 85% |
| RequireSchoolAccess | SYSTEM_ADMIN bypass; cross-school → 403/404 | ≥ 85% |
| RequestID | Header propagated to context | ≥ 85% |
| Recovery | Panic → 500 without crash | ≥ 85% |

---

# Integration Testing Standards

## Test Environment

| Requirement | Specification |
|-------------|---------------|
| Database | PostgreSQL 18+ (or 16+ compatible); dedicated `nusa_test` database |
| Migrations | Full `000001`–`000008` applied before suite |
| Isolation | Each test case uses transactions rolled back OR fresh seed per test |
| Auth helper | `loginAs(role, schoolID)` returns valid Bearer token |
| Seed data | Minimum 2 schools, 3 users (SYSTEM_ADMIN, SCHOOL_ADMIN×2, TEACHER×2 cross-school) |
| Command | `TEST_DATABASE_URL=... go test -v ./tests/integration/...` |

## Mandatory Scenarios (Every Endpoint)

For **each** endpoint below, implement **exactly these 5 test cases**:

| # | Scenario | Expected HTTP | Assertion Focus |
|---|----------|---------------|-----------------|
| S1 | **Happy path** | 200/201 | Response schema, DB row exists, correct fields |
| S2 | **Validation failure** | 422 | Invalid/missing required fields; `error.code = VALIDATION_ERROR` |
| S3 | **Authorization failure** | 401 or 403 | Wrong role or missing permission |
| S4 | **School isolation** | 404 | User from School B cannot access School A resource |
| S5 | **Persistence validation** | — | Direct SQL confirms DB state matches API response |

---

## Sprint 3.5 Endpoint Test Matrix

### EP-01: `PUT /api/v1/learning-planning/tp-sets/:id`

| Scenario | Test ID | Key Setup | Key Assertion |
|----------|---------|-----------|---------------|
| S1 Happy | `IT_EP01_S1_UpdateTPSet_Draft` | TEACHER owns DRAFT set | Status unchanged or valid transition; `updated_at` changed |
| S2 Validation | `IT_EP01_S2_UpdateTPSet_InvalidStatus` | Invalid status value | 422, no DB change |
| S3 Auth | `IT_EP01_S3_UpdateTPSet_NoPermission` | Token without `tp:UPDATE` | 403 |
| S4 Isolation | `IT_EP01_S4_UpdateTPSet_CrossSchool` | Teacher B accesses School A set | 404 |
| S5 Persistence | `IT_EP01_S5_UpdateTPSet_DBState` | SQL `SELECT status, updated_at FROM tp_sets WHERE id = ?` | Matches response |

### EP-02: `GET /api/v1/learning-planning/tp-sets/:id/versions`

| Scenario | Test ID | Key Setup | Key Assertion |
|----------|---------|-----------|---------------|
| S1 Happy | `IT_EP02_S1_GetTPVersions` | TP set with 2+ TP item versions | Array ordered by `version_no` |
| S2 Validation | `IT_EP02_S2_GetTPVersions_InvalidID` | Malformed UUID | 422 |
| S3 Auth | `IT_EP02_S3_GetTPVersions_Unauthenticated` | No token | 401 |
| S4 Isolation | `IT_EP02_S4_GetTPVersions_CrossSchool` | Other school's set | 404 |
| S5 Persistence | `IT_EP02_S5_GetTPVersions_Count` | SQL count of `tp` versions = response length | Match |

### EP-03: `PUT /api/v1/assessment/:id`

| Scenario | Test ID | Key Setup | Key Assertion |
|----------|---------|-----------|---------------|
| S1 Happy | `IT_EP03_S1_UpdateAssessment_Draft` | DRAFT assessment | Fields updated; snapshot columns **unchanged** |
| S2 Validation | `IT_EP03_S2_UpdateAssessment_Approved` | APPROVED assessment edit attempt | 409 `INVALID_STATE_TRANSITION` |
| S3 Auth | `IT_EP03_S3_UpdateAssessment_NoPermission` | TEACHER without `assessment:UPDATE` | 403 |
| S4 Isolation | `IT_EP03_S4_UpdateAssessment_CrossSchool` | Other school's assessment | 404 |
| S5 Persistence | `IT_EP03_S5_UpdateAssessment_SnapshotImmutable` | SQL: `success_criteria_snapshot` unchanged after update | Match pre-update value |

### EP-04: `POST /api/v1/assessment/:id/approve`

| Scenario | Test ID | Key Setup | Key Assertion |
|----------|---------|-----------|---------------|
| S1 Happy | `IT_EP04_S1_ApproveAssessment` | SCHOOL_ADMIN; DRAFT assessment | `status = APPROVED`, `approved_by` set |
| S2 Validation | `IT_EP04_S2_ApproveAssessment_AlreadyApproved` | APPROVED assessment | 409 `ALREADY_APPROVED` |
| S3 Auth | `IT_EP04_S3_ApproveAssessment_Teacher` | TEACHER token | 403 |
| S4 Isolation | `IT_EP04_S4_ApproveAssessment_CrossSchool` | Admin from School B | 404 |
| S5 Persistence | `IT_EP04_S5_ApproveAssessment_Audit` | SQL + `audit_logs` entry | Approval recorded |

### EP-05: `POST /api/v1/assessment/evidences/upload`

| Scenario | Test ID | Key Setup | Key Assertion |
|----------|---------|-----------|---------------|
| S1 Happy | `IT_EP05_S1_UploadEvidence` | Multipart file ≤ 10MB, valid MIME | `evidence_data.files[]` populated with storage reference |
| S2 Validation | `IT_EP05_S2_UploadEvidence_Oversized` | File > 10MB | 422 |
| S3 Auth | `IT_EP05_S3_UploadEvidence_Unauthenticated` | No token | 401 |
| S4 Isolation | `IT_EP05_S4_UploadEvidence_CrossSchool` | Assessment from other school | 404 |
| S5 Persistence | `IT_EP05_S5_UploadEvidence_Checksum` | SQL JSONB contains `integrity.checksum_sha256` | Present |

### EP-06: `GET /api/v1/assessment/evidences/:id`

| Scenario | Test ID | Key Setup | Key Assertion |
|----------|---------|-----------|---------------|
| S1 Happy | `IT_EP06_S1_GetEvidenceDetail` | Evidence with linked evaluation | Response includes evidence + evaluation summary |
| S2 Validation | `IT_EP06_S2_GetEvidence_NotFound` | Random UUID | 404 `EVIDENCE_NOT_FOUND` |
| S3 Auth | `IT_EP06_S3_GetEvidence_NoPermission` | Missing `assessment:READ` | 403 |
| S4 Isolation | `IT_EP06_S4_GetEvidence_CrossSchool` | Other school's evidence | 404 |
| S5 Persistence | `IT_EP06_S5_GetEvidence_Join` | SQL join assessment→user→school | School matches token |

### EP-07: `GET /api/v1/learning-planning/atp-sets/:id`

| Scenario | Test ID | Key Setup | Key Assertion |
|----------|---------|-----------|---------------|
| S1 Happy | `IT_EP07_S1_GetATPSetDetail` | ATP set with ATP items + TP join | Nested ATP items with TP titles |
| S2 Validation | `IT_EP07_S2_GetATPSet_InvalidID` | Bad UUID | 422 |
| S3 Auth | `IT_EP07_S3_GetATPSet_Unauthenticated` | No token | 401 |
| S4 Isolation | `IT_EP07_S4_GetATPSet_CrossSchool` | Other school's set | 404 |
| S5 Persistence | `IT_EP07_S5_GetATPSet_ItemCount` | SQL count `atp` rows = response item count | Match |

### EP-08: `GET /api/v1/learning-planning/modul-ajar-sets/:id`

| Scenario | Test ID | Key Setup | Key Assertion |
|----------|---------|-----------|---------------|
| S1 Happy | `IT_EP08_S1_GetModulAjarSetDetail` | Set with modul ajar items | Nested items with ATP week |
| S2 Validation | `IT_EP08_S2_GetModulAjarSet_NotFound` | Nonexistent ID | 404 |
| S3 Auth | `IT_EP08_S3_GetModulAjarSet_NoPermission` | Missing `tp:READ` | 403 |
| S4 Isolation | `IT_EP08_S4_GetModulAjarSet_CrossSchool` | Cross-tenant | 404 |
| S5 Persistence | `IT_EP08_S5_GetModulAjarSet_DB` | SQL verify `modul_ajar` rows | Match response |

---

## Defect Fix Integration Tests (Mandatory)

| Defect | Test ID | Pass Criteria |
|--------|---------|---------------|
| DEF-001 | `IT_DEF001_EvaluationCreatesRevision` | PUT evaluation creates new row; prior `is_current_version = false` |
| DEF-002 | `IT_DEF002_ReportAchievementIntegration` | `POST .../refresh-achievement` sets `achievement_data` NOT NULL |
| DEF-004 | `IT_DEF004_RevisionNoIncrements` | Second revision has `revision_no = prior + 1` |
| DEF-008 | `IT_DEF008_EvaluationHistoryQuery` | `GET .../evaluations/history/:evidence_id` returns all revisions ordered |
| DEF-009 | `IT_DEF009_FeedbackHistoryPreserved` | Feedback change creates `evaluation_feedback_history` row |

---

## Regression Integration Tests (Existing Routes — Security Hardening)

Education routes currently lack `RequirePermission` middleware. After P0 security hardening, **each route group** must have:

| Route Group | Minimum Tests |
|-------------|---------------|
| `/api/v1/curriculum/*` | S1 list CP + S4 isolation N/A (global read) + S3 unauthenticated |
| `/api/v1/learning-planning/*` | S4 on tp-sets, atp-sets, modul-ajar-sets (existing GET/POST) |
| `/api/v1/assessment/*` | S4 on assessments, evidences, evaluations |
| `/api/v1/reporting/*` | S4 on narrative-reports |
| `/api/v1/students/*`, `/api/v1/classes/*` | S4 achievement endpoints |

**Regression test ID pattern**: `IT_REG_{Module}_{Endpoint}_{Scenario}`

---

# Security Testing Standards

## RBAC Matrix Tests

Automated tests must validate the frozen permission matrix (Architecture Freeze v2 Appendix B):

| Test ID | Role | Permission | Endpoint Example | Expected |
|---------|------|------------|------------------|----------|
| `SEC_RBAC_001` | TEACHER | `school:CREATE` | `POST /api/v1/schools` | 403 |
| `SEC_RBAC_002` | SCHOOL_ADMIN | `tp:APPROVE` | `POST .../tp-sets/:id/approve` | 200/201 |
| `SEC_RBAC_003` | TEACHER | `tp:APPROVE` | `POST .../tp-sets/:id/approve` | 403 |
| `SEC_RBAC_004` | SYSTEM_ADMIN | `user:DELETE` | `DELETE /api/v1/users/:id` | 200/204 |
| `SEC_RBAC_005` | TEACHER | `user:DELETE` | `DELETE /api/v1/users/:id` | 403 |

**Pass criteria**: 100% of matrix cells tested for Sprint 3.5 touched resources (`tp`, `assessment`, `reporting`).

## Resource-Level Authorization

| Test ID | Scenario | Expected |
|---------|----------|----------|
| `SEC_RES_001` | Teacher updates own artifact | 200 |
| `SEC_RES_002` | Teacher updates another teacher's artifact (same school) | 403 |
| `SEC_RES_003` | SCHOOL_ADMIN approves artifact in own school | 200 |
| `SEC_RES_004` | Owner check uses `user_id` / `generated_by` from DB | SQL-verified |

## School Isolation

| Test ID | Scenario | Expected |
|---------|----------|----------|
| `SEC_ISO_001` | School A teacher reads School A resource | 200 |
| `SEC_ISO_002` | School A teacher reads School B resource by UUID | **404** (not 403) |
| `SEC_ISO_003` | SYSTEM_ADMIN reads any school resource | 200 (read) |
| `SEC_ISO_004` | Enumeration resistance: sequential UUID probe | 404, no data leak in body |

## JWT Validation

| Test ID | Scenario | Expected |
|---------|----------|----------|
| `SEC_JWT_001` | Valid access token | 200 on `/api/v1/auth/me` |
| `SEC_JWT_002` | Expired access token | 401 `TOKEN_EXPIRED` |
| `SEC_JWT_003` | Tampered signature | 401 |
| `SEC_JWT_004` | Missing `Authorization` header | 401 |
| `SEC_JWT_005` | Refresh rotation: old refresh token rejected after use | 401 |
| `SEC_JWT_006` | Logout invalidates refresh tokens | 401 on refresh |

## Input Validation

| Test ID | Attack Vector | Expected |
|---------|---------------|----------|
| `SEC_INP_001` | SQL injection in query params | 422/400, no DB error exposed |
| `SEC_INP_002` | Oversized JSON body (> 1MB) | 413 or 422 |
| `SEC_INP_003` | Invalid UUID in path | 422 |
| `SEC_INP_004` | Invalid enum value | 422 `VALIDATION_ERROR` |
| `SEC_INP_005` | XSS payload in text fields | Stored as-is; no execution (API returns escaped or raw per contract) |
| `SEC_INP_006` | Path traversal in upload filename | 422, sanitized `object_key` |

## Mass Assignment Prevention

| Test ID | Scenario | Expected |
|---------|----------|----------|
| `SEC_MA_001` | Request includes `approved_by` on create | Field ignored; `approved_by` NULL until approve |
| `SEC_MA_002` | Request includes `is_current_version: true` on create | Server sets per business rules |
| `SEC_MA_003` | Request includes `school_id` on artifact create | Ignored; school derived from user |
| `SEC_MA_004` | Request includes `role_id` on user self-update | 403 or field stripped |

---

# Migration Testing Standards

## Frozen Migration Sequence

Per Database Schema Freeze v1 — migrations `000001` through `000008`.

## Mandatory Tests Per Migration

| Test ID | Type | Procedure | Pass Criteria |
|---------|------|-----------|---------------|
| `MIG_FWD_ALL` | Forward | Apply `000001`→`000008` on empty DB | Exit 0; all tables exist |
| `MIG_ROLL_008` | Rollback | Down `000008`, verify `class_id` removed | Schema matches post-007 |
| `MIG_ROLL_ALL` | Rollback | Down to `000000` in reverse order | Clean empty schema |
| `MIG_IDEM` | Idempotency | Up twice (where applicable) | Second run no error or no-op |
| `MIG_003_DATA` | Data | Seed assessment with modul_ajar; run 000003 | `tp_id` NOT NULL; `modul_ajar_id` absent |
| `MIG_004_REV` | Data | Evaluations exist pre-004 | All have `revision_no = 1`, `is_current_version = true` |
| `MIG_006_TP` | Schema | After 000006 | `tp.version_no`, `tp.is_current_version`, `tp.parent_version_id` exist |
| `MIG_008_CLASS` | Schema | After 000008 | `narrative_reports.class_id` NOT NULL |

## Data Preservation Validation

After migrations on seeded production-like data:

```sql
-- Post-000003: no orphan assessments
SELECT COUNT(*) FROM assessments WHERE tp_id IS NULL;  -- must be 0

-- Post-000004: single current evaluation per evidence
SELECT evidence_id, COUNT(*) FROM evaluations
WHERE is_current_version = true GROUP BY evidence_id HAVING COUNT(*) > 1;  -- must return 0 rows

-- Post-000005: achievement columns nullable (no data loss)
SELECT COUNT(*) FROM narrative_reports;  -- count unchanged from pre-migration
```

## Migration Correction Tests (Pre-Apply)

| Issue | Validation Test |
|-------|-----------------|
| 000006 targets `tps` | Static check: migration file contains `ALTER TABLE tp` |
| 000003 typo | Data migration sets `success_criteria_snapshot` not `success_criteria` |
| 000004 duplicate column | `revision_no` add is idempotent |

---

# API Contract Testing

## Tools & Location

| Tool | Purpose |
|------|---------|
| `backend/docs/api/openapi.yaml` | Contract source |
| `backend/tests/integration/handler_test.go` | Handler contract tests |
| `docs/qa/API_CONTRACT_VALIDATION_REPORT.md` | Manual validation log |

## Required Validations

| Category | Rule | Test Method |
|----------|------|-------------|
| Request validation | Required fields enforced per OpenAPI `required` | S2 integration scenarios |
| Response validation | Response body matches OpenAPI schema | JSON schema assert or struct field check |
| Error envelope | `{ success: false, error: { code, message, details }, timestamp }` | All error scenarios |
| Pagination | `page`, `limit`, `total`, `total_pages` when list endpoint | `IT_REG_Pagination_*` |
| Content-Type | `application/json` except upload (`multipart/form-data`) | Header assert |
| Auth header | `Authorization: Bearer <token>` on protected routes | SEC_JWT_* |

## Backward Compatibility Rules

| Change Type | Allowed in Sprint 3.5? | Test Required |
|-------------|------------------------|---------------|
| Add optional response field | Yes | Existing tests still pass |
| Add required request field | **No** without version bump | N/A — forbidden |
| Rename response field | **No** | Contract diff fails |
| Change error code | **No** without Architecture amendment | Contract test update |
| Change HTTP status for same condition | **No** | Regression test fails |

**Contract test command** (target CI job):

```bash
# Validate OpenAPI parses
npx @redocly/cli lint backend/docs/api/openapi.yaml

# Integration tests assert contract
go test -v ./tests/integration/... -run Contract
```

---

# End-to-End Testing

## Tooling

| Tool | Scope |
|------|-------|
| Playwright | Frontend E2E (target Sprint 3B) |
| API-only E2E | `backend/tests/integration/full_flow_test.go` — must be implemented, not placeholder |

## Critical User Journeys (Mandatory)

| Journey ID | Name | Steps | Pass Criteria |
|------------|------|-------|---------------|
| E2E-01 | Teacher Login | Login → dashboard → `/auth/me` | Token stored; role displayed |
| E2E-02 | TP Workflow | Browse CP → create TP set → view detail → edit (EP-01) → approve | Set reaches APPROVED |
| E2E-03 | TP Version History | Approve TP → revise → GET versions (EP-02) | ≥ 2 versions visible |
| E2E-04 | Assessment Workflow | Create assessment from TP → edit (EP-03) → approve (EP-04) | Snapshots preserved |
| E2E-05 | Evidence Collection | Upload file (EP-05) → view detail (EP-06) → link rubric | Status COLLECTED→LINKED |
| E2E-06 | Evaluation Revision | Evaluate → revise feedback → history (DEF-008, DEF-009) | revision_no incremented |
| E2E-07 | ATP / Modul Ajar Detail | View ATP set (EP-07) → view Modul Ajar set (EP-08) | Nested items render |
| E2E-08 | Narrative Report | Create report → refresh achievement (DEF-002) → verify content | `achievement_data` populated |

**E2E pass criteria**: All 8 journeys pass in staging environment with seeded data.

---

# Performance Validation

MVP-only requirements. No load testing infrastructure required.

| Metric | Threshold | Measurement | Scope |
|--------|-----------|-------------|-------|
| API p95 latency | **< 500ms** | `backend/tests/integration/performance_test.go` | Sprint 3.5 endpoints |
| API p99 latency | **< 1000ms** | Same | Sprint 3.5 endpoints |
| Health check | **< 50ms** | `GET /health` | Always |
| DB query per request | **< 10** queries | Log count in integration test | Detail endpoints (EP-06–08) |
| Concurrent users | **10** simultaneous | Basic goroutine test | Login + read only |
| File upload | **10MB** max completes < 5s | EP-05 | Local MinIO or mock |

**Forbidden in MVP**: k6/Gatling suites, read replicas, caching benchmarks, projection latency tests.

**Pass criteria**: `performance_test.go` reports all thresholds met; no P1 perf regression > 20% from baseline.

---

# Bug Severity Matrix

## Severity Definitions

| Severity | Code | Definition | Example |
|----------|------|------------|---------|
| **P0 — Critical** | P0 | Data loss, security breach, broken core workflow, production down | Cross-school data leak; evaluation overwrites without revision |
| **P1 — High** | P1 | Major feature broken; no workaround | Approve endpoint returns 500; migration rollback fails |
| **P2 — Medium** | P2 | Feature degraded; workaround exists | Pagination off-by-one; non-critical field missing in response |
| **P3 — Low** | P3 | Cosmetic, docs, minor UX | Typo in error message; logging format inconsistency |

## Release Blocking Criteria

| Severity | Sprint 3.5 Release | Sprint 3B Release | Production |
|----------|-------------------|-------------------|------------|
| P0 open | **BLOCK** | **BLOCK** | **BLOCK** |
| P1 open | **BLOCK** (all must be fixed per Sprint 3.5 package) | **BLOCK** | **BLOCK** |
| P2 open | Allowed with documented waiver | ≤ 5 open | ≤ 3 open |
| P3 open | Unlimited | ≤ 20 open | ≤ 10 open |

## Defect Severity Mapping (Sprint 3.5)

| Defect ID | Severity | Release Blocker |
|-----------|----------|-----------------|
| DEF-001 | P0 | Yes |
| DEF-002 | P0 | Yes |
| DEF-004 | P1 | Yes |
| DEF-008 | P1 | Yes |
| DEF-009 | P1 | Yes |

## Security Finding Severity

| Finding | Severity |
|---------|----------|
| Authentication bypass | P0 |
| Cross-tenant data access | P0 |
| Missing RBAC on admin endpoint | P0 |
| Missing RBAC on education endpoint | P0 (Sprint 3.5 P0 security item) |
| SQL injection exploitable | P0 |
| Mass assignment on privileged field | P1 |
| Missing rate limiting | P1 |
| Verbose error stack trace in production | P2 |

---

# Release Gate

Implementation is **REJECTED** and merge blocked if **any** condition below is true:

## Automatic Rejection (CI)

| # | Condition |
|---|-----------|
| R1 | `go test ./...` fails |
| R2 | Unit coverage on changed packages < 80% |
| R3 | Any `IT_EP##_S*` integration test missing or failing |
| R4 | Any `IT_DEF00*` test failing |
| R5 | Any `SEC_*` P0 security test failing |
| R6 | Migration forward or rollback fails on test DB |
| R7 | OpenAPI lint fails |
| R8 | Placeholder test detected (`assert.True(t, true)` in new code) |

## Manual Rejection (QA Sign-off)

| # | Condition |
|---|-----------|
| R9 | Any open P0 or P1 bug in sprint scope |
| R10 | Resource instance-level authorization not implemented |
| R11 | Education routes lack school isolation tests |
| R12 | Evaluation revision creates in-place update (DEF-001 regression) |
| R13 | Assessment snapshot mutated on update (Architecture violation) |
| R14 | Schema drift from Database Schema Freeze v1 |
| R15 | New migration outside `000001`–`000008` without Board approval |
| R16 | CQRS, event sourcing, or projection code introduced |
| R17 | E2E-01 through E2E-06 not passing in staging |

## Release Approval Requires

- [ ] All G0–G6 quality gates green
- [ ] All 8 Sprint 3.5 endpoints deployed and tested
- [ ] All 5 defects verified fixed
- [ ] Security audit checklist 100% for P0 items
- [ ] QA sign-off recorded in sprint review
- [ ] No R1–R17 rejection conditions active

---

# AI Generated Code Validation

Every PR containing AI-generated code must complete this checklist **before merge request**.

## Step 1 — Architecture Compliance (Mandatory)

| Check | Pass | Fail Action |
|-------|------|-------------|
| Read Architecture Freeze v2 relevant section | ☐ | Stop — do not implement |
| Read Database Schema Freeze v1 table definitions | ☐ | Stop — schema mismatch |
| No forbidden patterns (CQRS, generic repo, handler business logic) | ☐ | Reject PR |
| Layer boundaries: Handler → Service → Repository | ☐ | Reject PR |

## Step 2 — Code Review Checks (Mandatory)

| Check | Pass |
|-------|------|
| No direct repository calls from handler | ☐ |
| No SQL in service layer | ☐ |
| DTO used for request/response (no raw entity bind) | ☐ |
| School boundary validated in service for school-scoped resources | ☐ |
| Transaction wraps multi-table writes | ☐ |
| Errors use frozen error codes | ☐ |
| No hardcoded secrets or credentials | ☐ |

## Step 3 — Test Generation (Mandatory)

| Check | Pass |
|-------|------|
| Unit tests for every new service method | ☐ |
| Unit tests for every new domain validation | ☐ |
| Integration test S1–S5 for every new endpoint | ☐ |
| Security test for RBAC + isolation on new endpoint | ☐ |
| No placeholder tests | ☐ |
| Test IDs follow naming convention (`IT_EP##_S#_*`) | ☐ |

## Step 4 — Automated Validation (Mandatory)

```bash
# Run before every merge
cd backend
go vet ./...
go test ./... -count=1
go test ./tests/integration/... -count=1 -tags=integration  # with TEST_DATABASE_URL
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out | tail -1  # verify ≥ 80% on changed packages
```

| Check | Pass |
|-------|------|
| `go vet` clean | ☐ |
| All unit tests pass | ☐ |
| All integration tests pass | ☐ |
| Coverage ≥ 80% on changed packages | ☐ |
| Linter clean (project standards) | ☐ |

## Step 5 — Contract & Migration (If Applicable)

| Check | Pass |
|-------|------|
| OpenAPI updated for API changes | ☐ |
| Migration up + down tested | ☐ |
| Data preservation queries pass | ☐ |
| Schema matches Freeze v1 | ☐ |

## Step 6 — Human Review Triggers (Mandatory Review)

AI code **requires human review** when any of:

- Security middleware or auth logic changed
- Migration includes data transformation
- Evaluation versioning logic changed
- School isolation logic changed
- New external dependency added
- Performance-sensitive query (JOIN > 3 tables)

## AI Merge Rejection Rules

Reject immediately if AI output:

1. Creates tables not in Schema Freeze v1
2. Adds `school_id` to artifact tables (violates isolation model)
3. Reintroduces `modul_ajar_id` on assessments
4. Updates evaluation in-place without revision row
5. Skips tests with `t.Skip()` for security scenarios
6. Uses `gen_random_uuid()` instead of `gen_uuid_v7()`
7. Implements CQRS, event bus, or projection tables

---

# Appendix A — CI Pipeline Requirements (Target State)

```yaml
# Required CI jobs for merge
jobs:
  backend-unit:
    - go vet ./...
    - go test ./... -coverprofile=coverage.out
    - coverage threshold check ≥ 80%

  backend-integration:
    services: [postgres:18]
    - migrate up 000001-000008
    - go test ./tests/integration/... -tags=integration

  backend-security:
    - go test ./... -run 'SEC_|Security'

  openapi-lint:
    - npx @redocly/cli lint backend/docs/api/openapi.yaml

  migration-rollback:
    - migrate up && migrate down 1 && migrate up
```

## Appendix B — Test Traceability

| Requirement Source | Test Prefix |
|--------------------|-------------|
| Sprint 3.5 EP-01–08 | `IT_EP##_S#_*` |
| Defect fixes | `IT_DEF00#_*` |
| Security | `SEC_*` |
| Migrations | `MIG_*` |
| E2E journeys | `E2E-0#` |
| Regression | `IT_REG_*` |

## Appendix C — Document Authority

This document defines **what DONE means** for Sprint 3.5 and all subsequent AI implementation. Conflicts with ad-hoc testing approaches are resolved in favor of this freeze.

**Status**: LOCKED  
**Enforcement**: CI + QA sign-off + Release Gate R1–R17  
**Amendment**: Architecture Board approval required

# NUSA Platform — Repository Modification Map

**Version**: 1.0  
**Date**: June 2026  
**Status**: LOCKED — GOVERNANCE FOR ALL AI CODING AGENTS  
**Audience**: Devin, Codex, Claude Code, Cursor Agent, OpenHands, Roo Code, and similar autonomous systems  
**Parent Documents**:
- [`NUSA_ARCHITECTURE_FREEZE_V2.md`](NUSA_ARCHITECTURE_FREEZE_V2.md)
- [`NUSA_DATABASE_SCHEMA_FREEZE_V1.md`](NUSA_DATABASE_SCHEMA_FREEZE_V1.md)
- [`NUSA_TESTING_QUALITY_GATE_FREEZE_V1.md`](NUSA_TESTING_QUALITY_GATE_FREEZE_V1.md)
- [`SPRINT_3.5_EXECUTION_SEQUENCE.md`](SPRINT_3.5_EXECUTION_SEQUENCE.md)
- [`docs/foundation/09_REPOSITORY_ARCHITECTURE.md`](foundation/09_REPOSITORY_ARCHITECTURE.md)

---

# Repository Governance Principles

| # | Principle | Enforcement |
|---|-----------|-------------|
| G1 | **Frozen architecture is law** | Implementation must conform to Architecture Freeze v2. Ambiguity → stop and escalate; do not invent patterns. |
| G2 | **Modify only in allowed zones** | Changes outside ALLOWED/CONDITIONAL zones are rejected at PR review. |
| G3 | **One agent, one ownership file** | High-conflict files have single-writer rules (see File Ownership Matrix). |
| G4 | **Layer boundaries are immutable** | Handler → Service → Repository → PostgreSQL. No shortcuts. |
| G5 | **Contracts are frozen** | API paths, error codes, and response envelopes change only via Architecture Board amendment. |
| G6 | **Schema is frozen** | Migrations `000001`–`000008` only; no new tables without Board approval. |
| G7 | **Tests accompany behavior** | No production code merge without tests per Testing Quality Gate Freeze. |
| G8 | **No scope expansion** | CQRS, event sourcing, new domains, and framework swaps are prohibited. |
| G9 | **Traceability** | Every PR cites freeze doc section + Sprint 3.5 work item (EP-*, DEF-*, SEC-*). |
| G10 | **Smallest diff wins** | Modify only files required for the assigned work package. |

---

# Allowed Modification Zones

**Legend**

| Classification | Meaning |
|----------------|---------|
| **ALLOWED** | AI agents may create/modify freely within sprint scope and layer rules |
| **CONDITIONAL** | Modification allowed only with stated prerequisites and review |
| **PROHIBITED** | No AI modification; human Architecture Board only |

---

## Root & DevOps

| Path | Classification | Conditions |
|------|----------------|------------|
| `README.md` | CONDITIONAL | Minor factual updates only; no architecture changes |
| `.env.example` | CONDITIONAL | New env vars for approved features only; no secrets |
| `.gitignore` | CONDITIONAL | DevOps review |
| `.github/workflows/` | CONDITIONAL | DevOps + Platform Architect review; no weakening CI gates |
| `docker-compose.yml` | CONDITIONAL | Service additions require Platform Architect approval |
| `podman-compose.yml` | CONDITIONAL | Same as docker-compose |
| `.devin/` | PROHIBITED | Local agent config — not part of product |
| `LICENSE` | PROHIBITED | Legal |

---

## Documentation

| Path | Classification | Conditions |
|------|----------------|------------|
| `docs/foundation/` | **PROHIBITED** | Locked foundation; amend via Board process only |
| `docs/NUSA_*_FREEZE*.md` | **PROHIBITED** | Governance source of truth |
| `docs/SPRINT_3.5_EXECUTION_*.md` | CONDITIONAL | Status checkbox updates only in Phase 11 |
| `docs/qa/`, `docs/sprints/` | CONDITIONAL | QA-authored reports; agents may append validation results |
| `docs/research/` | PROHIBITED | Reference only |
| `backend/docs/api/openapi.yaml` | **ALLOWED** | Phase 10; must match frozen contracts |
| `backend/docs/CODE_STANDARDS.md` | PROHIBITED | Standards locked |
| `backend/docs/cqrs/` | **PROHIBITED** | Future reference — must not drive implementation |
| `backend/docs/sprints/`, `quality/`, `security/` | CONDITIONAL | Report updates only |
| `frontend/docs/` | CONDITIONAL | Frontend team scope |

---

## Backend — Application Code

| Path | Classification | Conditions |
|------|----------------|------------|
| `backend/internal/domain/` | **ALLOWED** | Sprint 3.5 domain changes only; no new aggregate roots outside freeze |
| `backend/internal/service/` | **ALLOWED** | Business logic; transactions here; follow execution sequence |
| `backend/internal/repository/` | **ALLOWED** | SQL only; parameterized queries; school-scope joins |
| `backend/internal/middleware/` | CONDITIONAL | Security hardening (Phase 6); Platform + Security review |
| `backend/internal/router/` | CONDITIONAL | **Single-writer**; route registration only; no business logic |
| `backend/internal/handler/` | CONDITIONAL | Legacy — prefer `modules/*/handler.go`; do not expand pattern |
| `backend/internal/error/` | CONDITIONAL | Frozen error codes only |
| `backend/internal/config/` | CONDITIONAL | New config keys for approved features only |
| `backend/internal/database/` | CONDITIONAL | Transaction helpers; no schema logic |
| `backend/internal/db/` | CONDITIONAL | Connection/retry; avoid duplicating `internal/database/` |
| `backend/internal/bootstrap/` | CONDITIONAL | Wiring only; Platform Architect review |
| `backend/internal/server/` | CONDITIONAL | Server lifecycle; rare changes |
| `backend/internal/auth/` | CONDITIONAL | Security review mandatory |
| `backend/internal/infrastructure/` | CONDITIONAL | Avoid expansion; prefer `internal/repository/` |
| `backend/internal/logger/` | PROHIBITED | Unless observability Board approval |
| `backend/modules/*/handler.go` | **ALLOWED** | HTTP layer only; thin handlers |
| `backend/modules/*/` (other) | CONDITIONAL | Prefer `internal/service` + `internal/repository` |
| `backend/pkg/jwt/` | **PROHIBITED** | Immutable security core (bugfix only, Security review) |
| `backend/pkg/errors/` | CONDITIONAL | Add frozen error codes only |
| `backend/pkg/response/` | CONDITIONAL | Response envelope shape frozen |
| `backend/pkg/rabbitmq/` | CONDITIONAL | AI orchestration scope only |
| `backend/cmd/` | CONDITIONAL | Entrypoints; minimal changes |
| `backend/cmd/seed/` | CONDITIONAL | Test/seed data for sprint scope |

**Note**: There is no `internal/validation/` or `internal/mapper/` directory today. Validation lives on DTOs (`binding` tags) and in services. **Do not create** `internal/validation/` or generic mapper frameworks — use module DTOs and explicit mapper functions per Architecture Freeze v2.

---

## Backend — Data & Migrations

| Path | Classification | Conditions |
|------|----------------|------------|
| `backend/migrations/` | CONDITIONAL | **Only** `000001`–`000008`; see Database Modification Rules |
| `backend/services/postgres/` | CONDITIONAL | Extensions SQL; DBA review |

---

## Backend — Tests

| Path | Classification | Conditions |
|------|----------------|------------|
| `backend/tests/unit/` | **ALLOWED** | Required for all changed packages |
| `backend/tests/integration/` | **ALLOWED** | Required per endpoint/defect; real DB |
| `backend/internal/**/*_test.go` | **ALLOWED** | Co-located tests encouraged |
| `backend/modules/**/*_test.go` | **ALLOWED** | Handler tests with mocked service |

---

## Frontend

| Path | Classification | Conditions |
|------|----------------|------------|
| `frontend/src/api/` | **ALLOWED** | Must match frozen API contract |
| `frontend/src/features/` | **ALLOWED** | Feature modules per screen |
| `frontend/src/pages/` | **ALLOWED** | Route pages |
| `frontend/src/services/queries/` | **ALLOWED** | TanStack Query hooks |
| `frontend/src/services/commands/` | **ALLOWED** | Mutation hooks |
| `frontend/src/components/` | **ALLOWED** | Shared UI |
| `frontend/src/hooks/` | **ALLOWED** | Shared hooks |
| `frontend/src/shared/types/` | **ALLOWED** | Types mirroring API DTOs |
| `frontend/src/theme/` | CONDITIONAL | Design system; no drive-by refactors |
| `frontend/src/icons/` | PROHIBITED | Generated/vendor icon set — do not bulk edit |
| `frontend/package.json` | **PROHIBITED** | No new dependencies without Board approval |
| `frontend/package-lock.json` | **PROHIBITED** | Follows package.json |
| `frontend/vite.config.ts` | CONDITIONAL | Build config; Platform review |
| `frontend/node_modules/` | **PROHIBITED** | Never commit or modify |

---

## AI Runtime

| Path | Classification | Conditions |
|------|----------------|------------|
| `ai-runtime/app/` | CONDITIONAL | Sprint scope AI agents only; AI Engineer review |
| `ai-runtime/tests/` | **ALLOWED** | Agent tests |
| `ai-runtime/requirements.txt` | **PROHIBITED** | Approved stack frozen |

---

## Scripts & Deploy

| Path | Classification | Conditions |
|------|----------------|------------|
| `scripts/` | CONDITIONAL | DevOps review; no production secrets |
| `deploy/` | CONDITIONAL | Platform Architect review |
| `backend/Makefile` | CONDITIONAL | Add targets for approved workflows only |

---

# File Ownership Matrix

## Platform & Cross-Cutting

| Module / File | Owner | Modification Rules | Review |
|---------------|-------|-------------------|--------|
| `backend/internal/router/router.go` | Backend Lead | One agent at a time; route registration only | Backend Architect |
| `backend/internal/bootstrap/bootstrap.go` | Platform | DI wiring only | Platform Architect |
| `backend/pkg/jwt/*` | Security | Bugfix only; no algorithm change | Security Architect |
| `backend/internal/middleware/auth_middleware.go` | Security | Auth flow frozen | Security Architect |
| `backend/internal/middleware/role.go` | Security | RBAC middleware | Security Architect |
| `backend/migrations/*` | DBA / Backend Lead | Sequential `000001`–`000008` only | DBA + Architect |
| `backend/docs/api/openapi.yaml` | API Contract | Sync with handlers | Backend + QA |
| `.github/workflows/ci.yml` | DevOps | Never remove test jobs | DevOps |

## Domain Modules (Backend)

| Module | Owner | Modification Rules | Review |
|--------|-------|-------------------|--------|
| **Identity** (`modules/auth`, `modules/users`, `modules/schools`, `modules/roles`) | Backend Identity | Auth/RBAC scope; no OAuth/Keycloak | Security Architect |
| **Curriculum** (`modules/curriculum`, `internal/domain/curriculum.go`, `curriculum_*`) | Backend Curriculum | Read/import CP; no new curriculum domains | Backend Architect |
| **Learning Planning** (`modules/learning_planning`, `tp.go`, `learning_planning.go`) | Backend LP | TP/ATP/Modul Ajar sets only | Backend Architect |
| **Assessment** (`modules/assessment`, `assessment.go`, `assessment_service.go`) | Backend Assessment | **Single-writer** for service file; sequential DEF then EP | Backend Architect |
| **Reporting** (`modules/reporting`, `reporting.go`) | Backend Reporting | Narrative reports only; no analytics | Backend Architect |
| **Achievement** (`modules/achievement`, `achievement.go`) | Backend Reporting | Runtime computation only; no persistence tables | Backend Architect |

## Service / Repository Pairing

| Service File | Repository File | Single-Writer |
|--------------|-----------------|---------------|
| `internal/service/assessment_service.go` | `internal/repository/assessment_repository.go` | **Yes** — one agent |
| `internal/service/learning_planning_service.go` | `internal/repository/learning_planning_repository.go` | Prefer one agent per PR |
| `internal/service/reporting_service.go` | `internal/repository/reporting_repository.go` | DEF-002 owner |
| `internal/service/tp_service.go` | TP methods in `learning_planning_repository.go` | EP-01, EP-02 owner |
| `internal/service/user_service.go` | `internal/repository/user_repository.go` | Identity changes only |

## Frontend Modules

| Area | Owner | Modification Rules | Review |
|------|-------|-------------------|--------|
| `frontend/src/api/` | Frontend Lead | Axios clients; match OpenAPI | Frontend Architect |
| `frontend/src/features/*` | Feature owner | One feature per agent | Frontend peer review |
| Auth / RBAC components | Frontend Security | Match backend permissions | Security + Frontend |

---

# AI Agent Modification Rules

## Allowed

| Action | Scope | Constraints |
|--------|-------|-------------|
| Create endpoint handlers | `backend/modules/*/handler.go` | Thin; call service only |
| Create / extend services | `backend/internal/service/` | Business logic, transactions, authorization calls |
| Create / extend repositories | `backend/internal/repository/` | SQL only; school-scope joins |
| Create domain types & DTOs | `backend/internal/domain/` | Must match Schema Freeze v1 |
| Create unit tests | `backend/tests/unit/`, `*_test.go` | ≥ 80% coverage on changed code |
| Create integration tests | `backend/tests/integration/` | 5 scenarios per endpoint |
| Register routes | `backend/internal/router/router.go` | After handler exists; follow merge order |
| Update OpenAPI | `backend/docs/api/openapi.yaml` | Phase 10; examples required |
| Implement Sprint 3.5 endpoints | EP-01 through EP-08 | Follow execution sequence |
| Fix defects | DEF-001, 002, 004, 008, 009 | Follow defect merge order |
| Add permission middleware | Education route groups | Phase 6 |
| Frontend API clients & screens | `frontend/src/api/`, `features/`, `pages/` | Contract-aligned |

## Forbidden

| Action | Reason |
|--------|--------|
| Redesign architecture | Violates Architecture Freeze v2 |
| Introduce CQRS (command/query buses) | Explicitly excluded |
| Introduce event sourcing / event store | Explicitly excluded |
| Introduce read models / projections | Explicitly excluded |
| Add domains: Analytics, Competency Graph, Digital Twin, LLR | Future wave |
| Add `students` / `classes` / `achievement_*` tables | Schema freeze |
| Add migration `000009+` without Board approval | Schema freeze |
| Replace Gin, PostgreSQL, React, RabbitMQ, JWT library | ADR-locked stack |
| Add npm/Go dependencies | Requires Board approval |
| Create `internal/validation/` framework | Over-engineering |
| Generic repository `Repository<T>` | Forbidden pattern |
| Handler → repository direct calls | Layer violation |
| Business logic in handlers | Layer violation |
| Repository calling repository | Layer violation |
| Modify `docs/foundation/*` | Locked |
| Modify freeze governance docs | Board only |
| Implement from `backend/docs/cqrs/*` | Future-only docs |
| Bulk refactor / rename unrelated code | Scope creep |
| Disable or skip CI tests | Quality gate violation |
| Commit secrets, `.env`, credentials | Security violation |
| Change JWT algorithm (HS256 → RS256) | Security freeze |
| Add Keycloak, OAuth2, Kafka, Kubernetes manifests | Excluded technologies |
| Change assessment `tp_id` back to `modul_ajar_id` | Schema freeze |

## Approved Libraries (Immutable Set)

**Backend (Go)** — from `backend/go.mod`:

`gin`, `golang-jwt/jwt`, `sqlx`, `pgx`, `pq`, `rabbitmq/amqp091-go`, `viper`, `testify`, `zap`, `golang.org/x/crypto`

**Frontend** — from `frontend/package.json` (no additions):

`react`, `react-router-dom`, `@tanstack/react-query`, `axios`, `@mui/material`, `vite`, `typescript`, `formik`, `dayjs`, `i18next`

**AI Runtime** — from `ai-runtime/requirements.txt`:

`fastapi`, `uvicorn`, `pydantic`, `langgraph`, `openai`, `anthropic`, `httpx`

---

# Database Modification Rules

## When Migrations ARE Allowed

| Condition | Requirement |
|-----------|-------------|
| Migration number is `000001`–`000008` | Listed in Schema Freeze v1 |
| Migration is in assigned sprint work | e.g., 000004 for DEF-001, 000008 for `class_id` |
| Paired `.up.sql` and `.down.sql` exist | Rollback mandatory |
| Schema matches freeze document exactly | No extra tables/columns |
| `gen_uuid_v7()` used for all new PKs | No `gen_random_uuid()` |
| Forward + rollback tested | `MIG_*` tests pass |
| Single migration PR merges before dependent code | Per execution sequence Phase 1 |
| Correction to existing migration (000003, 000006) | Documented in Schema Freeze; DBA review |

## When Migrations ARE Forbidden

| Condition | Action |
|-----------|--------|
| New migration `000009` or higher | **Stop** — Architecture Board approval required |
| New table not in Schema Freeze v1 | **Forbidden** |
| `school_id` column on artifact tables | **Forbidden** — use user join isolation |
| `modul_ajar_id` on `assessments` | **Forbidden** — removed in 000003 |
| CQRS/event store/projection tables | **Forbidden** |
| `curriculum_versions`, `workflow_history` | Wave 2 — forbidden in Sprint 3.5 |
| Editing applied migrations in production | **Forbidden** — create corrective migration only if Board approves |
| DROP TABLE on production artifacts | **Forbidden** — use `status = ARCHIVED` |
| Changing `gen_uuid_v7()` function semantics | Platform Board only |

## Migration File Authority

| File | Agent Access |
|------|--------------|
| `000001`–`000003` | **PROHIBITED** — already applied; fix only via Board-approved correction PR |
| `000004`–`000008` | CONDITIONAL — sprint-assigned agent only |
| New files | **PROHIBITED** without Board |

---

# API Contract Protection Rules

## Frozen Contract Sources (Priority Order)

1. [`NUSA_ARCHITECTURE_FREEZE_V2.md`](NUSA_ARCHITECTURE_FREEZE_V2.md) Part 3 — API Contract Freeze  
2. [`SPRINT_3.5_EXECUTION_PACKAGE.md`](SPRINT_3.5_EXECUTION_PACKAGE.md) — Sprint 3.5 endpoints  
3. [`backend/docs/api/openapi.yaml`](../../backend/docs/api/openapi.yaml) — generated mirror  

**Rule**: OpenAPI follows the freeze; OpenAPI does not lead the freeze.

## Immutable Contract Elements

| Element | Rule |
|---------|------|
| Base URL | `/api/v1` |
| Public auth paths | `/api/v1/public/auth/login`, `/api/v1/public/auth/refresh` |
| Response envelope | `{ success, data/error, timestamp }` |
| Error shape | `{ code, message, details }` |
| Frozen error codes | See Architecture Freeze v2 §3.2 — no renames |
| HTTP status semantics | 404 for cross-tenant; 403 for RBAC; 422 validation |
| Assessment link field | `tp_id` + snapshots — not `modul_ajar_id` |
| Sprint 3.5 endpoint paths | EP-01 through EP-08 paths are fixed |

## Contract Change Process

| Change Type | Allowed? | Process |
|-------------|----------|---------|
| Add optional response field | Yes | Update OpenAPI + backward compat test |
| Add required request field | **No** | Board amendment + version bump |
| Remove field | **No** | Board amendment |
| Rename endpoint | **No** | Board amendment |
| Change error code for same condition | **No** | Board amendment |

## Agent Rule

If implementation requires a contract change → **stop**, document gap, escalate. Do not implement and "fix docs later."

---

# Security Modification Rules

## Immutable Security Components

| Component | Path | Rule |
|-----------|------|------|
| JWT algorithm | `backend/pkg/jwt/` | HS256 only; no OIDC |
| Password hashing | User service | bcrypt cost 12 |
| Token claims schema | JWT service | `user_id`, `school_id`, `role`, `permissions` |
| Refresh token storage | `refresh_tokens` table | DB-backed rotation |
| Auth middleware order | `router.go` | Recovery → CORS → RequestID → Auth → Permission |
| RBAC role names | `domain/role.go` | `SYSTEM_ADMIN`, `SCHOOL_ADMIN`, `TEACHER` |
| Permission format | `resource:ACTION` | No alternate format |
| Cross-tenant response | Service layer | **404**, not 403 |
| Public routes | `/api/v1/public/auth/*` only | No new public education routes |

## Conditionally Mutable (Security Review Required)

| Component | Allowed Change |
|-----------|----------------|
| `RequirePermission` on education routes | **Required** Sprint 3.5 Phase 6 |
| `SchoolBoundaryValidator` | Add/implement — P0 |
| Resource ownership checks | Add to services — P0 |
| Rate limiting middleware | P1 — optional Sprint 3.5 |
| CORS origins | Config only; no `*` in production |
| Account lockout thresholds | Config constants; document change |

## Prohibited Security Changes

- Disable auth middleware for convenience  
- Add `SYSTEM_ADMIN` bypass without audit logging  
- Store passwords in plaintext or reversible encryption  
- Log access tokens or refresh tokens  
- Expose `password_hash` in API responses  
- Skip school boundary check "for MVP"  
- Add backdoor/debug auth endpoints  

---

# Architectural Violation Matrix

## Acceptable Changes

| Change | Example |
|--------|---------|
| Add handler method for frozen endpoint | `PUT tp-sets/:id` |
| Add service method with transaction | `UpdateEvaluation` creates revision row |
| Add repository query with school join | `WHERE u.school_id = $1` |
| Add DTO with validation tags | `UpdateTPSetRequest` |
| Add integration test | `IT_EP01_S1_UpdateTPSet_Draft` |
| Add `RequirePermission("tp:UPDATE")` to route | Phase 6 security |
| Fix migration typo (000006 `tp` not `tps`) | Board-documented correction |
| Extend OpenAPI with new path | EP-05 upload |
| Add frontend feature component | Evidence upload panel |
| Append QA validation report | `docs/qa/` |

## Prohibited Changes

| Change | Violation |
|--------|-----------|
| Add `events` table + event publisher | Event sourcing |
| Split into `commands/` and `queries/` packages | CQRS |
| Add `projections` package | Read models |
| Create `competency_graph` module | Future domain |
| Replace Gin with Echo/Fiber | Framework swap |
| Add Prisma/GORM alongside sqlx | Unapproved ORM |
| Handler calls `assessment_repository` directly | Layer bypass |
| `UPDATE evaluations SET ...` for revision | DEF-001 violation |
| Add `school_id` to `evidences` table | Schema/isolation model |
| Change `POST /api/v1/auth/login` path | Contract break |
| Introduce GraphQL gateway | Out of scope |
| Add Redis caching layer | Premature optimization |
| Refactor entire `internal/` package layout | Architecture drift |
| Auto-generate mappers via reflection | Forbidden pattern |
| Merge EP-06 before DEF-001 | Execution sequence violation |

---

# Pull Request Validation Checklist

Every PR from an AI agent **must** include this checklist in the description. Reviewer marks pass/fail.

## Identity & Scope

- [ ] PR title references work item: `EP-0X`, `DEF-00X`, `SEC-`, or `MIG-00X`
- [ ] Cites Architecture Freeze v2 section and/or Schema Freeze v1 table
- [ ] Diff scope limited to assigned work package — no unrelated files
- [ ] Branch name follows `s35/*` convention (per Execution Sequence)

## Architecture & Layers

- [ ] Handler → Service → Repository only; no layer skips
- [ ] No business logic in handlers
- [ ] No SQL in services
- [ ] DTOs used for request/response; no mass assignment
- [ ] No CQRS, event sourcing, projections, or new domains
- [ ] No new Go/npm/Python dependencies

## Database

- [ ] No migration outside `000001`–`000008` unless Board-approved
- [ ] Migrations have `.down.sql` if new/changed
- [ ] `gen_uuid_v7()` for new UUID columns
- [ ] No forbidden columns (`assessments.modul_ajar_id`, artifact `school_id`)

## API Contract

- [ ] Endpoints match frozen paths and methods
- [ ] Response envelope and error codes match freeze
- [ ] OpenAPI updated if API surface changed
- [ ] No breaking changes to existing fields

## Security

- [ ] Permission middleware on new/modified routes (or tracked in Phase 6 PR)
- [ ] School boundary validated in service layer
- [ ] Cross-tenant access returns 404
- [ ] No secrets in code
- [ ] JWT/auth code unchanged unless Security-reviewed

## Testing

- [ ] Unit tests added; ≥ 80% coverage on changed packages
- [ ] Integration tests: 5 scenarios per new endpoint
- [ ] Defect tests for DEF-* work
- [ ] No placeholder tests (`assert.True(t, true)`)
- [ ] `go test ./...` passes locally

## Merge & Sequence

- [ ] Prerequisites merged to `main` (per Execution Sequence)
- [ ] No conflict with single-writer files (`router.go`, `assessment_service.go`)
- [ ] Rebased on latest `main`

## Reviewer Sign-off

| Role | Required When | Approved |
|------|---------------|----------|
| Peer developer | All PRs | ☐ |
| Backend Architect | Service/repo/schema logic | ☐ |
| Security Architect | Auth, middleware, isolation | ☐ |
| QA | New endpoints / defects | ☐ |
| Platform Architect | CI, deploy, migrations | ☐ |

---

# Final Governance Rules

**Immutable constraints for all AI agents.** Violation → PR rejected; agent must rollback.

## Architecture

1. Modular monolith — Handler → Service → Repository → PostgreSQL.  
2. Six MVP domains only: Identity, Curriculum, Learning Planning, Assessment, Evidence (within Assessment), Reporting, Achievement (computed).  
3. No CQRS, event sourcing, event store, command bus, query bus, saga, workflow engine expansion.  
4. No microservice extraction or new deployable services.  
5. No generic repository pattern or reflection mapping.

## Data

6. PostgreSQL single database, schema `public`.  
7. Migrations limited to `000001`–`000008` unless Architecture Board approves otherwise.  
8. Assessment references `tp_id` with snapshots — never `modul_ajar_id`.  
9. School isolation via `users.school_id` join — no `school_id` on artifact tables.  
10. Evaluation revisions insert new rows — never in-place update of revisioned fields.

## API & Security

11. Frozen API contract — paths, envelopes, error codes.  
12. JWT custom HS256 — no Keycloak/OAuth2.  
13. RBAC roles and permission format frozen.  
14. Education routes must enforce permissions (Sprint 3.5 Phase 6).  
15. Cross-tenant access returns 404.

## Process

16. Follow [`SPRINT_3.5_EXECUTION_SEQUENCE.md`](SPRINT_3.5_EXECUTION_SEQUENCE.md) merge order.  
17. Pass [`NUSA_TESTING_QUALITY_GATE_FREEZE_V1.md`](NUSA_TESTING_QUALITY_GATE_FREEZE_V1.md) gates before merge.  
18. One agent per high-conflict file at a time.  
19. Smallest possible diff; no drive-by refactors.  
20. When uncertain → **stop and escalate**; never guess architecture.

## Agent Preamble (Copy into Agent Instructions)

```text
You are implementing NUSA Sprint 3.5 under frozen governance.
Before editing: read NUSA_ARCHITECTURE_FREEZE_V2.md (relevant section),
NUSA_DATABASE_SCHEMA_FREEZE_V1.md (tables), SPRINT_3.5_EXECUTION_SEQUENCE.md (order).
Modify only ALLOWED zones. Do not add migrations beyond 000008.
Do not introduce CQRS, event sourcing, or new domains.
Follow Handler → Service → Repository. Include tests.
If contract or schema change seems needed: STOP and report — do not implement.
```

---

**Document Status**: LOCKED  
**Enforcement**: PR checklist + CI + Architecture Board escalation  
**Amendment**: Architecture Board written approval only

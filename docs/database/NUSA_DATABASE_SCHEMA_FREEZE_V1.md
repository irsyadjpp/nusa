# NUSA Platform — Database Schema Freeze v1

**Version**: 1.0  
**Date**: June 2026  
**Status**: LOCKED — FINAL DATABASE SOURCE OF TRUTH  
**Parent Document**: [`NUSA_ARCHITECTURE_FREEZE_V2.md`](NUSA_ARCHITECTURE_FREEZE_V2.md) (approved)  
**Scope**: MVP Wave 1 — PostgreSQL single schema (`public`)

---

# Executive Summary

This document freezes the **physical PostgreSQL schema** for NUSA MVP Wave 1. It formalizes the database model required by Architecture Freeze v2 without redesigning application architecture.

**What this freeze establishes:**

- **25 tables** across 6 domain areas (Identity, Curriculum, Learning Planning, Assessment, Reporting, Platform)
- **1 database function** (`gen_uuid_v7()`)
- **8 locked migrations** (`000001`–`000008`) — no further migrations without Architecture Board approval
- **3 versioning models**: set-level generation sessions, row-level artifact versions, evaluation revision chains
- **Row-level multi-school isolation** via `users.school_id` (no `school_id` on artifact tables)

**What this freeze explicitly excludes:**

- CQRS, read models, projections, event store tables
- `students`, `classes`, `achievement_*` tables (deferred — external IDs or runtime computation)
- `curriculum_versions`, `workflow_history` (Wave 2 — not in MVP migration sequence)
- Object storage infrastructure (MinIO) — only `evidences.evidence_data` metadata is defined here

**Known gaps closed by this freeze:**

| Gap | Resolution |
|-----|------------|
| Migration `000006` references wrong table `tps` | Frozen correction: target table `tp` |
| Migration `000003` data migration typo `success_criteria` | Frozen correction: `success_criteria_snapshot` |
| Migration `000004` duplicates `revision_no` from `000003` | Frozen: `000004` adds only `is_current_version`, `parent_revision_id`, `evaluation_feedback_history` |
| `narrative_reports.class_id` in domain, missing in DB | Frozen migration `000008` |
| `evaluation_feedback_history.id` uses `gen_random_uuid()` | Frozen: must use `gen_uuid_v7()` |

---

# Database Principles

## UUID Strategy

| Rule | Specification |
|------|---------------|
| Primary keys | `UUID` on every table |
| Generation | `DEFAULT gen_uuid_v7()` on all new rows |
| Function | `gen_uuid_v7()` defined in `000001` — native `uuid_generate_v7()` on PG 18+, fallback `uuid_generate_v4()` |
| Extension | `uuid-ossp` enabled in `000001` |
| External references | `student_id`, `class_id` are `VARCHAR(50)` — external SIS identifiers, not UUID FKs |

## Naming Conventions

| Element | Convention | Example |
|---------|------------|---------|
| Tables | `snake_case`, plural or domain term | `tp_sets`, `narrative_reports` |
| Columns | `snake_case` | `created_at`, `tp_version_no` |
| Primary key | `id` | `id UUID PRIMARY KEY` |
| Foreign keys | `{entity}_id` | `tp_set_id`, `user_id` |
| Indexes | `idx_{table}_{columns}` | `idx_tp_tp_set_id` |
| Unique constraints | `uq_{table}_{columns}` or `unique_{name}` | `uq_tp_sets_cp_version` |
| Check constraints | `chk_{table}_{column}` | `chk_tp_status` |
| Booleans | `is_{adjective}` | `is_active`, `is_current_version` |
| Timestamps | `TIMESTAMP WITH TIME ZONE` | `created_at`, `updated_at` |

## Soft Delete Strategy

MVP does **not** use `deleted_at` columns.

| Pattern | Implementation |
|---------|----------------|
| User deactivation | `users.is_active = false` |
| School deactivation | `schools.is_active = false` |
| Artifact retirement | `status = 'ARCHIVED'` on set entities (`tp_sets`, `atp_sets`, `modul_ajar_sets`) |
| Physical delete | **Forbidden** on approved artifacts; CASCADE only on draft chains during set rollback |
| Audit preservation | `audit_logs` rows are never deleted |

## Audit Fields Strategy

| Tier | Fields | Tables |
|------|--------|--------|
| **Mandatory** | `created_at`, `updated_at` | All transactional tables |
| **Approval** | `approved_by`, `approved_at` | Set tables, assessments, rubrics, narrative_reports |
| **Ownership** | `user_id` or `generated_by` | All teacher-created artifacts |
| **Provenance** | `created_by`, `updated_by` | `schools`, `users` |
| **Centralized** | `audit_logs` | Cross-entity actions (login, approve, admin) |
| **Revision audit** | `evaluation_feedback_history` | Evaluation feedback changes |

`updated_at` must be set by application on every UPDATE (no automatic trigger required in MVP).

## Foreign Key Strategy

| Rule | Specification |
|------|---------------|
| Curriculum hierarchy | `ON DELETE CASCADE` — children die with parent |
| Set → items | `ON DELETE CASCADE` |
| TP → assessments | `ON DELETE CASCADE` |
| Assessment → evidences | `ON DELETE CASCADE` |
| Evidence → evaluations | `ON DELETE CASCADE` |
| User → refresh_tokens | `ON DELETE CASCADE` |
| User → artifacts (creator) | `ON DELETE RESTRICT` — cannot delete user with artifacts |
| School → users | `ON DELETE SET NULL` on `users.school_id` (existing); prefer deactivation over delete |
| Self-referential version | `parent_version_id`, `parent_revision_id` — no CASCADE (nullable FK) |
| Optional links | `rubric_id` on evidences — `ON DELETE` not set (default NO ACTION) |

## Versioning Strategy

Three distinct mechanisms — **do not conflate**:

| Mechanism | Scope | Key Columns |
|-----------|-------|-------------|
| **Set versioning** | Generation sessions | `tp_sets.version_no`, `atp_sets.version_no`, `modul_ajar_sets.version_no` — unique per parent |
| **Row versioning** | Mutable artifacts | `version_no`, `is_current_version`, `parent_version_id` on `tp`, `assessments`, `rubrics`, `narrative_reports` |
| **Point-in-time snapshot** | Assessment integrity | `tp_version_no`, `success_criteria_snapshot`, `tp_*_snapshot` on `assessments` |
| **Revision chain** | Evaluations | `revision_no`, `is_current_version`, `parent_revision_id` on `evaluations` |

## Multi-School Isolation Strategy

| Rule | Specification |
|------|---------------|
| Tenant key | `users.school_id` → `schools.id` |
| SYSTEM_ADMIN | `school_id IS NULL` |
| Artifact scoping | **No `school_id` column on artifact tables** — isolation via JOIN to `users` on `user_id` / `generated_by` |
| Curriculum data | Global reference data — not school-scoped |
| AI logs | `ai_generation_logs.school_id` denormalized for query performance |
| Application enforcement | Every artifact query: `JOIN users u ON artifact.user_id = u.id WHERE u.school_id = :school_id` |
| Student/class IDs | School-scoped by convention — validated in service layer |

---

# Complete ERD

## Aggregate: School (Identity & Access)

| Attribute | Value |
|-----------|-------|
| **Tables** | `schools` |
| **Ownership** | SYSTEM_ADMIN |
| **Relationships** | `schools` 1→N `users`; `schools` 1→N `ai_generation_logs` (optional) |

## Aggregate: User (Identity & Access)

| Attribute | Value |
|-----------|-------|
| **Tables** | `users`, `refresh_tokens` |
| **Ownership** | SCHOOL_ADMIN (school users), SYSTEM_ADMIN (platform) |
| **Relationships** | `users` N→1 `roles`, N→1 `schools`; `users` 1→N `refresh_tokens` |

## Aggregate: Role (Identity & Access)

| Attribute | Value |
|-----------|-------|
| **Tables** | `roles`, `permissions` |
| **Ownership** | SYSTEM_ADMIN |
| **Relationships** | `roles` 1→N `permissions`; `roles` 1→N `users` |

## Aggregate: Curriculum Reference (Curriculum)

| Attribute | Value |
|-----------|-------|
| **Tables** | `curriculum_subjects`, `curriculum_phases`, `curriculum_elements`, `curriculum_subelements`, `cp` |
| **Ownership** | SYSTEM_ADMIN (write), all roles (read) |
| **Relationships** | Hierarchy: subject+phase → element → subelement → cp |

```text
curriculum_subjects ──┐
curriculum_phases  ──┼── curriculum_elements ── curriculum_subelements ── cp
```

## Aggregate: TPSet (Learning Planning)

| Attribute | Value |
|-----------|-------|
| **Tables** | `tp_sets`, `tp` |
| **Ownership** | Teacher (`generated_by` / `user_id`), school via user |
| **Relationships** | `cp` 1→N `tp_sets` 1→N `tp`; `tp` self-ref `parent_version_id` |

## Aggregate: ATPSet (Learning Planning)

| Attribute | Value |
|-----------|-------|
| **Tables** | `atp_sets`, `atp` |
| **Ownership** | Teacher, school via user |
| **Relationships** | `tp_sets` 1→N `atp_sets` 1→N `atp`; `tp` 1→N `atp` |

## Aggregate: ModulAjarSet (Learning Planning)

| Attribute | Value |
|-----------|-------|
| **Tables** | `modul_ajar_sets`, `modul_ajar` |
| **Ownership** | Teacher, school via user |
| **Relationships** | `atp_sets` 1→N `modul_ajar_sets` 1→N `modul_ajar`; `atp` 1→N `modul_ajar` |

## Aggregate: Assessment (Assessment)

| Attribute | Value |
|-----------|-------|
| **Tables** | `assessments` |
| **Ownership** | Teacher (`user_id`), school via user |
| **Relationships** | `tp` 1→N `assessments`; self-ref `parent_version_id` |
| **Note** | References **TP**, not Modul Ajar (Architecture Freeze v2 §1.2) |

## Aggregate: Rubric (Assessment)

| Attribute | Value |
|-----------|-------|
| **Tables** | `rubrics` |
| **Ownership** | Teacher, school via user |
| **Relationships** | `assessments` 1→N `rubrics`; self-ref `parent_version_id` |

## Aggregate: Evidence (Assessment)

| Attribute | Value |
|-----------|-------|
| **Tables** | `evidences`, `evaluations`, `evaluation_feedback_history` |
| **Ownership** | Teacher, school via user |
| **Relationships** | `assessments` 1→N `evidences` 1→N `evaluations`; `rubrics` N→1 `evaluations`; `evaluations` 1→N `evaluation_feedback_history` |

## Aggregate: NarrativeReport (Reporting)

| Attribute | Value |
|-----------|-------|
| **Tables** | `narrative_reports` |
| **Ownership** | Teacher, school via user |
| **Relationships** | `users` 1→N `narrative_reports`; self-ref `parent_version_id` |
| **Note** | Achievement is cached JSONB, not a separate table |

## Platform Tables (No Aggregate Root)

| Table | Purpose |
|-------|---------|
| `audit_logs` | Cross-cutting audit trail |
| `ai_generation_logs` | AI orchestration provenance |

## Entity Exclusion List (Prevent Duplicates)

| Excluded Table | Reason | Alternative |
|----------------|--------|-------------|
| `students` | MVP — external SIS | `student_id VARCHAR(50)` on evidences, evaluations, narrative_reports |
| `classes` | MVP — external SIS | `class_id VARCHAR(50)` on narrative_reports |
| `achievement_*` | Runtime computed domain | `narrative_reports.achievement_data` JSONB cache |
| `teaching_plans` / `tps` | Wrong name | Canonical table: `tp` |
| `tujuan_pembelajaran` | Duplicate concept | `tp` items inside `tp_sets` |
| `kk_tp` / `success_criteria` table | Value object, not entity | `tp.success_criteria` JSONB |
| `workflow_history` | Wave 2 | `audit_logs` + status columns for MVP |
| `curriculum_versions` | Wave 2 | `cp.version` string for MVP |

---

# Final Table Definitions

## Function: `gen_uuid_v7()`

| Attribute | Value |
|-----------|-------|
| **Purpose** | Time-ordered UUID primary key generation |
| **Returns** | `UUID` |
| **Behavior** | `uuid_generate_v7()` if available, else `uuid_generate_v4()` |

---

## Table: `schools`

| Column | Type | Nullable | Default | Constraints |
|--------|------|----------|---------|-------------|
| `id` | UUID | NO | `gen_uuid_v7()` | PRIMARY KEY |
| `name` | VARCHAR(255) | NO | — | — |
| `code` | VARCHAR(50) | NO | — | UNIQUE |
| `address` | TEXT | YES | — | — |
| `phone` | VARCHAR(50) | YES | — | — |
| `email` | VARCHAR(255) | YES | — | — |
| `is_active` | BOOLEAN | NO | `true` | — |
| `created_at` | TIMESTAMPTZ | NO | `NOW()` | — |
| `updated_at` | TIMESTAMPTZ | NO | `NOW()` | — |
| `created_by` | UUID | YES | — | — |
| `updated_by` | UUID | YES | — | — |

**Indexes:** `idx_schools_code`, `idx_schools_is_active`

---

## Table: `roles`

| Column | Type | Nullable | Default | Constraints |
|--------|------|----------|---------|-------------|
| `id` | UUID | NO | `gen_uuid_v7()` | PRIMARY KEY |
| `name` | VARCHAR(50) | NO | — | UNIQUE |
| `description` | TEXT | YES | — | — |
| `is_active` | BOOLEAN | NO | `true` | — |
| `created_at` | TIMESTAMPTZ | NO | `NOW()` | — |
| `updated_at` | TIMESTAMPTZ | NO | `NOW()` | — |

**Indexes:** `idx_roles_name`, `idx_roles_is_active`

**Seed values:** `SYSTEM_ADMIN`, `SCHOOL_ADMIN`, `TEACHER`

---

## Table: `permissions`

| Column | Type | Nullable | Default | Constraints |
|--------|------|----------|---------|-------------|
| `id` | UUID | NO | `gen_uuid_v7()` | PRIMARY KEY |
| `role_id` | UUID | NO | — | FK → `roles(id)` ON DELETE CASCADE |
| `resource` | VARCHAR(100) | NO | — | — |
| `action` | VARCHAR(50) | NO | — | — |
| `created_at` | TIMESTAMPTZ | NO | `NOW()` | — |

**Constraints:** `UNIQUE (role_id, resource, action)`

**Indexes:** `idx_permissions_role_id`, `idx_permissions_resource`

---

## Table: `users`

| Column | Type | Nullable | Default | Constraints |
|--------|------|----------|---------|-------------|
| `id` | UUID | NO | `gen_uuid_v7()` | PRIMARY KEY |
| `email` | VARCHAR(255) | NO | — | UNIQUE, `chk_users_email_format` |
| `password_hash` | VARCHAR(255) | NO | — | — |
| `name` | VARCHAR(255) | NO | — | — |
| `role_id` | UUID | NO | — | FK → `roles(id)` |
| `school_id` | UUID | YES | — | FK → `schools(id)` ON DELETE SET NULL |
| `is_active` | BOOLEAN | NO | `true` | — |
| `failed_login_attempts` | INTEGER | NO | `0` | `chk_users_failed_login_attempts` (≥ 0) |
| `locked_until` | TIMESTAMPTZ | YES | — | — |
| `created_at` | TIMESTAMPTZ | NO | `NOW()` | — |
| `updated_at` | TIMESTAMPTZ | NO | `NOW()` | — |
| `created_by` | UUID | YES | — | FK → `users(id)` |
| `updated_by` | UUID | YES | — | FK → `users(id)` |

**Indexes:** `idx_users_email`, `idx_users_role_id`, `idx_users_school_id`, `idx_users_is_active`, `idx_users_created_at`

**Business rule (application):** SYSTEM_ADMIN → `school_id IS NULL`; SCHOOL_ADMIN/TEACHER → `school_id IS NOT NULL`

---

## Table: `refresh_tokens`

| Column | Type | Nullable | Default | Constraints |
|--------|------|----------|---------|-------------|
| `id` | UUID | NO | `gen_uuid_v7()` | PRIMARY KEY |
| `user_id` | UUID | NO | — | FK → `users(id)` ON DELETE CASCADE |
| `token` | TEXT | NO | — | UNIQUE |
| `expires_at` | TIMESTAMPTZ | NO | — | `chk_refresh_tokens_expires_at` |
| `revoked_at` | TIMESTAMPTZ | YES | — | — |
| `ip_address` | INET | YES | — | — |
| `created_at` | TIMESTAMPTZ | NO | `NOW()` | — |
| `created_by` | UUID | YES | — | — |

**Indexes:** `idx_refresh_tokens_user_id`, `idx_refresh_tokens_token`, `idx_refresh_tokens_expires_at`

---

## Table: `ai_generation_logs`

| Column | Type | Nullable | Default | Constraints |
|--------|------|----------|---------|-------------|
| `id` | UUID | NO | `gen_uuid_v7()` | PRIMARY KEY |
| `user_id` | UUID | NO | — | FK → `users(id)` ON DELETE CASCADE |
| `school_id` | UUID | YES | — | FK → `schools(id)` ON DELETE SET NULL |
| `artifact_type` | VARCHAR(50) | NO | — | — |
| `artifact_id` | UUID | YES | — | — |
| `provider` | VARCHAR(50) | NO | — | — |
| `model` | VARCHAR(100) | YES | — | — |
| `tokens_used` | INTEGER | YES | — | — |
| `estimated_cost` | DECIMAL(10,4) | YES | — | — |
| `response_time_ms` | INTEGER | YES | — | — |
| `status` | VARCHAR(20) | NO | — | — |
| `error_message` | TEXT | YES | — | — |
| `prompt_snapshot` | TEXT | YES | — | — |
| `response_snapshot` | TEXT | YES | — | — |
| `created_at` | TIMESTAMPTZ | NO | `NOW()` | — |

**Indexes:** `idx_ai_generation_logs_user_id`, `idx_ai_generation_logs_school_id`, `idx_ai_generation_logs_artifact`, `idx_ai_generation_logs_created_at`

---

## Table: `curriculum_subjects`

| Column | Type | Nullable | Default | Constraints |
|--------|------|----------|---------|-------------|
| `id` | UUID | NO | `gen_uuid_v7()` | PRIMARY KEY |
| `code` | VARCHAR(50) | NO | — | UNIQUE |
| `name` | VARCHAR(255) | NO | — | — |
| `name_en` | VARCHAR(255) | YES | — | — |
| `description` | TEXT | YES | — | — |
| `is_active` | BOOLEAN | NO | `true` | — |
| `created_at` | TIMESTAMPTZ | NO | `NOW()` | — |
| `updated_at` | TIMESTAMPTZ | NO | `NOW()` | — |

**Indexes:** `idx_curriculum_subjects_code`, `idx_curriculum_subjects_is_active`

---

## Table: `curriculum_phases`

| Column | Type | Nullable | Default | Constraints |
|--------|------|----------|---------|-------------|
| `id` | UUID | NO | `gen_uuid_v7()` | PRIMARY KEY |
| `code` | VARCHAR(50) | NO | — | UNIQUE |
| `name` | VARCHAR(255) | NO | — | — |
| `name_en` | VARCHAR(255) | YES | — | — |
| `description` | TEXT | YES | — | — |
| `grade_level_start` | INTEGER | YES | — | — |
| `grade_level_end` | INTEGER | YES | — | — |
| `is_active` | BOOLEAN | NO | `true` | — |
| `created_at` | TIMESTAMPTZ | NO | `NOW()` | — |
| `updated_at` | TIMESTAMPTZ | NO | `NOW()` | — |

**Indexes:** `idx_curriculum_phases_code`, `idx_curriculum_phases_is_active`

---

## Table: `curriculum_elements`

| Column | Type | Nullable | Default | Constraints |
|--------|------|----------|---------|-------------|
| `id` | UUID | NO | `gen_uuid_v7()` | PRIMARY KEY |
| `subject_id` | UUID | NO | — | FK → `curriculum_subjects(id)` CASCADE |
| `phase_id` | UUID | NO | — | FK → `curriculum_phases(id)` CASCADE |
| `code` | VARCHAR(50) | NO | — | UNIQUE(subject_id, phase_id, code) |
| `name` | VARCHAR(255) | NO | — | — |
| `name_en` | VARCHAR(255) | YES | — | — |
| `description` | TEXT | YES | — | — |
| `is_active` | BOOLEAN | NO | `true` | — |
| `created_at` | TIMESTAMPTZ | NO | `NOW()` | — |
| `updated_at` | TIMESTAMPTZ | NO | `NOW()` | — |

**Indexes:** `idx_curriculum_elements_subject_id`, `idx_curriculum_elements_phase_id`, `idx_curriculum_elements_code`, `idx_curriculum_elements_is_active`

---

## Table: `curriculum_subelements`

| Column | Type | Nullable | Default | Constraints |
|--------|------|----------|---------|-------------|
| `id` | UUID | NO | `gen_uuid_v7()` | PRIMARY KEY |
| `element_id` | UUID | NO | — | FK → `curriculum_elements(id)` CASCADE |
| `code` | VARCHAR(50) | NO | — | UNIQUE(element_id, code) |
| `name` | VARCHAR(255) | NO | — | — |
| `name_en` | VARCHAR(255) | YES | — | — |
| `description` | TEXT | YES | — | — |
| `is_active` | BOOLEAN | NO | `true` | — |
| `created_at` | TIMESTAMPTZ | NO | `NOW()` | — |
| `updated_at` | TIMESTAMPTZ | NO | `NOW()` | — |

**Indexes:** `idx_curriculum_subelements_element_id`, `idx_curriculum_subelements_code`, `idx_curriculum_subelements_is_active`

---

## Table: `cp`

| Column | Type | Nullable | Default | Constraints |
|--------|------|----------|---------|-------------|
| `id` | UUID | NO | `gen_uuid_v7()` | PRIMARY KEY |
| `subject_id` | UUID | NO | — | FK CASCADE |
| `phase_id` | UUID | NO | — | FK CASCADE |
| `element_id` | UUID | NO | — | FK CASCADE |
| `subelement_id` | UUID | NO | — | FK CASCADE |
| `code` | VARCHAR(50) | NO | — | UNIQUE(full hierarchy, code) |
| `description` | TEXT | NO | — | — |
| `competency_code` | VARCHAR(50) | YES | — | — |
| `learning_objectives` | JSONB | NO | — | — |
| `competency_standards` | JSONB | NO | — | — |
| `time_allocation_hours` | INTEGER | NO | — | > 0 |
| `hours_per_week` | INTEGER | NO | — | 1–40 |
| `version` | VARCHAR(20) | NO | — | Curriculum version label (MVP) |
| `is_active` | BOOLEAN | NO | `true` | — |
| `imported_at` | TIMESTAMPTZ | NO | `NOW()` | — |
| `imported_by` | UUID | YES | — | FK → `users(id)` |

**Indexes:** `idx_cp_subject_id`, `idx_cp_phase_id`, `idx_cp_element_id`, `idx_cp_subelement_id`, `idx_cp_code`, `idx_cp_version`, `idx_cp_is_active`

---

## Table: `tp_sets`

| Column | Type | Nullable | Default | Constraints |
|--------|------|----------|---------|-------------|
| `id` | UUID | NO | `gen_uuid_v7()` | PRIMARY KEY |
| `cp_id` | UUID | NO | — | FK → `cp(id)` CASCADE |
| `version_no` | INTEGER | NO | `1` | UNIQUE(cp_id, version_no) |
| `status` | VARCHAR(20) | NO | — | DRAFT, UNDER_REVIEW, APPROVED, REJECTED, ARCHIVED |
| `generation_source` | VARCHAR(50) | NO | — | AI_GENERATED, MANUAL |
| `generation_reason` | TEXT | YES | — | — |
| `generated_by` | UUID | NO | — | FK → `users(id)` |
| `ai_generation_id` | UUID | YES | — | FK → `ai_generation_logs(id)` |
| `approved_by` | UUID | YES | — | FK → `users(id)` |
| `approved_at` | TIMESTAMPTZ | YES | — | — |
| `created_at` | TIMESTAMPTZ | NO | `NOW()` | — |
| `updated_at` | TIMESTAMPTZ | NO | `NOW()` | — |

**Indexes:** `idx_tp_sets_cp_id`, `idx_tp_sets_version_no`, `idx_tp_sets_status`, `idx_tp_sets_generated_by`, `idx_tp_sets_approved_by`, `idx_tp_sets_ai_generation_id`

---

## Table: `tp`

| Column | Type | Nullable | Default | Constraints |
|--------|------|----------|---------|-------------|
| `id` | UUID | NO | `gen_uuid_v7()` | PRIMARY KEY |
| `tp_set_id` | UUID | NO | — | FK → `tp_sets(id)` CASCADE |
| `sequence_number` | INTEGER | NO | — | > 0 |
| `cp_id` | UUID | NO | — | FK CASCADE |
| `subject_id` | UUID | NO | — | FK CASCADE |
| `phase_id` | UUID | NO | — | FK CASCADE |
| `element_id` | UUID | NO | — | FK CASCADE |
| `subelement_id` | UUID | NO | — | FK CASCADE |
| `user_id` | UUID | NO | — | FK → `users(id)` |
| `status` | VARCHAR(20) | NO | — | Workflow statuses |
| `title` | VARCHAR(255) | YES | — | — |
| `learning_objectives` | JSONB | NO | — | — |
| `time_allocation` | JSONB | NO | — | — |
| `prerequisites` | JSONB | YES | — | — |
| `estimated_weeks` | INTEGER | YES | — | > 0 if set |
| `success_criteria` | JSONB | YES* | — | *Required before set approval (app) |
| `version_no` | INTEGER | NO | `1` | Item version |
| `is_current_version` | BOOLEAN | NO | `true` | — |
| `parent_version_id` | UUID | YES | — | FK → `tp(id)` |
| `created_at` | TIMESTAMPTZ | NO | `NOW()` | — |
| `updated_at` | TIMESTAMPTZ | NO | `NOW()` | — |

**Indexes:** `idx_tp_tp_set_id`, `idx_tp_sequence_number`, `idx_tp_tp_set_id_sequence`, hierarchy FK indexes, `idx_tp_status`, `idx_tp_user_id`, `idx_tp_success_criteria` (GIN), `idx_tp_set_version` (tp_set_id, version_no), `idx_tp_parent_version`

---

## Table: `atp_sets`

| Column | Type | Nullable | Default | Constraints |
|--------|------|----------|---------|-------------|
| `id` | UUID | NO | `gen_uuid_v7()` | PRIMARY KEY |
| `tp_set_id` | UUID | NO | — | FK → `tp_sets(id)` CASCADE |
| `version_no` | INTEGER | NO | `1` | UNIQUE(tp_set_id, version_no) |
| `status` | VARCHAR(20) | NO | — | Workflow statuses |
| `generation_source` | VARCHAR(50) | NO | — | AI_GENERATED, MANUAL |
| `generation_reason` | TEXT | YES | — | — |
| `generated_by` | UUID | NO | — | FK → `users(id)` |
| `ai_generation_id` | UUID | YES | — | FK → `ai_generation_logs(id)` |
| `approved_by` | UUID | YES | — | FK → `users(id)` |
| `approved_at` | TIMESTAMPTZ | YES | — | — |
| `created_at` | TIMESTAMPTZ | NO | `NOW()` | — |
| `updated_at` | TIMESTAMPTZ | NO | `NOW()` | — |

**Indexes:** `idx_atp_sets_tp_set_id`, `idx_atp_sets_version_no`, `idx_atp_sets_status`, `idx_atp_sets_generated_by`, `idx_atp_sets_approved_by`, `idx_atp_sets_ai_generation_id`

---

## Table: `atp`

| Column | Type | Nullable | Default | Constraints |
|--------|------|----------|---------|-------------|
| `id` | UUID | NO | `gen_uuid_v7()` | PRIMARY KEY |
| `atp_set_id` | UUID | NO | — | FK → `atp_sets(id)` CASCADE |
| `tp_id` | UUID | NO | — | FK → `tp(id)` CASCADE |
| `user_id` | UUID | NO | — | FK → `users(id)` |
| `status` | VARCHAR(20) | NO | — | Workflow statuses |
| `academic_calendar` | JSONB | NO | — | — |
| `class_schedule` | JSONB | NO | — | — |
| `weekly_sequence` | JSONB | NO | — | — |
| `assessment_schedule` | JSONB | YES | — | — |
| `created_at` | TIMESTAMPTZ | NO | `NOW()` | — |
| `updated_at` | TIMESTAMPTZ | NO | `NOW()` | — |

**Indexes:** `idx_atp_atp_set_id`, `idx_atp_tp_id`, `idx_atp_user_id`, `idx_atp_status`, `idx_atp_created_at`

**Versioning note:** ATP items do **not** have row-level version columns. Version context is `atp_sets.version_no`.

---

## Table: `modul_ajar_sets`

| Column | Type | Nullable | Default | Constraints |
|--------|------|----------|---------|-------------|
| `id` | UUID | NO | `gen_uuid_v7()` | PRIMARY KEY |
| `atp_set_id` | UUID | NO | — | FK → `atp_sets(id)` CASCADE |
| `version_no` | INTEGER | NO | `1` | UNIQUE(atp_set_id, version_no) |
| `status` | VARCHAR(20) | NO | — | Workflow statuses |
| `generation_source` | VARCHAR(50) | NO | — | AI_GENERATED, MANUAL |
| `generation_reason` | TEXT | YES | — | — |
| `generated_by` | UUID | NO | — | FK → `users(id)` |
| `ai_generation_id` | UUID | YES | — | FK → `ai_generation_logs(id)` |
| `approved_by` | UUID | YES | — | FK → `users(id)` |
| `approved_at` | TIMESTAMPTZ | YES | — | — |
| `created_at` | TIMESTAMPTZ | NO | `NOW()` | — |
| `updated_at` | TIMESTAMPTZ | NO | `NOW()` | — |

**Indexes:** `idx_modul_ajar_sets_atp_set_id`, `idx_modul_ajar_sets_version_no`, `idx_modul_ajar_sets_status`, `idx_modul_ajar_sets_generated_by`, `idx_modul_ajar_sets_approved_by`, `idx_modul_ajar_sets_ai_generation_id`

---

## Table: `modul_ajar`

| Column | Type | Nullable | Default | Constraints |
|--------|------|----------|---------|-------------|
| `id` | UUID | NO | `gen_uuid_v7()` | PRIMARY KEY |
| `modul_ajar_set_id` | UUID | NO | — | FK → `modul_ajar_sets(id)` CASCADE |
| `atp_id` | UUID | NO | — | FK → `atp(id)` CASCADE |
| `week` | INTEGER | NO | — | ≥ 1 |
| `topic` | JSONB | NO | — | — |
| `resources` | JSONB | YES | — | — |
| `class_characteristics` | JSONB | YES | — | — |
| `learning_activities` | JSONB | NO | — | — |
| `resource_requirements` | JSONB | YES | — | — |
| `assessment_methods` | JSONB | YES | — | — |
| `status` | VARCHAR(20) | NO | — | Workflow statuses |
| `created_at` | TIMESTAMPTZ | NO | `NOW()` | — |
| `updated_at` | TIMESTAMPTZ | NO | `NOW()` | — |

**Indexes:** `idx_modul_ajar_modul_ajar_set_id`, `idx_modul_ajar_atp_id`, `idx_modul_ajar_week`, `idx_modul_ajar_status`, `idx_modul_ajar_created_at`

**Versioning note:** Modul Ajar items do **not** have row-level version columns. Version context is `modul_ajar_sets.version_no`.

---

## Table: `assessments`

| Column | Type | Nullable | Default | Constraints |
|--------|------|----------|---------|-------------|
| `id` | UUID | NO | `gen_uuid_v7()` | PRIMARY KEY |
| `tp_id` | UUID | NO | — | FK → `tp(id)` CASCADE |
| `tp_version_no` | INTEGER | NO | `1` | Snapshot reference |
| `success_criteria_snapshot` | JSONB | NO | — | Immutable after create |
| `tp_title_snapshot` | TEXT | YES | — | — |
| `tp_learning_objectives_snapshot` | JSONB | YES | — | — |
| `tp_time_allocation_snapshot` | JSONB | YES | — | — |
| `user_id` | UUID | NO | — | FK → `users(id)` |
| `assessment_type` | VARCHAR(20) | NO | — | FORMATIVE, SUMMATIVE |
| `status` | VARCHAR(20) | NO | — | DRAFT, APPROVED, REJECTED |
| `assessment_items` | JSONB | NO | — | — |
| `answer_key` | JSONB | NO | — | — |
| `scoring_guidelines` | JSONB | NO | — | — |
| `ai_confidence_score` | DECIMAL(3,2) | YES | — | 0.00–1.00 |
| `ai_generated_at` | TIMESTAMPTZ | YES | — | — |
| `ai_agent_version` | VARCHAR(20) | YES | — | — |
| `version_no` | INTEGER | NO | `1` | Row version |
| `is_current_version` | BOOLEAN | NO | `true` | — |
| `parent_version_id` | UUID | YES | — | FK → `assessments(id)` |
| `created_at` | TIMESTAMPTZ | NO | `NOW()` | — |
| `updated_at` | TIMESTAMPTZ | NO | `NOW()` | — |
| `approved_at` | TIMESTAMPTZ | YES | — | — |
| `approved_by` | UUID | YES | — | FK → `users(id)` |

**Removed column (frozen):** `modul_ajar_id` — must not be reintroduced.

**Indexes:** `idx_assessments_tp_id`, `idx_assessments_tp_version`, `idx_assessments_success_criteria` (GIN), `idx_assessments_tp_snapshot`, `idx_assessments_user_id`, `idx_assessments_status`, `idx_assessments_created_at`, `idx_assessments_approved_at`

---

## Table: `rubrics`

| Column | Type | Nullable | Default | Constraints |
|--------|------|----------|---------|-------------|
| `id` | UUID | NO | `gen_uuid_v7()` | PRIMARY KEY |
| `assessment_id` | UUID | NO | — | FK → `assessments(id)` CASCADE |
| `user_id` | UUID | NO | — | FK → `users(id)` |
| `rubric_type` | VARCHAR(20) | NO | — | ANALYTIC, HOLISTIC |
| `status` | VARCHAR(20) | NO | — | DRAFT, APPROVED, REJECTED |
| `performance_criteria` | JSONB | NO | — | — |
| `performance_levels` | JSONB | NO | — | — |
| `scoring_guidelines` | JSONB | NO | — | — |
| `ai_confidence_score` | DECIMAL(3,2) | YES | — | 0.00–1.00 |
| `ai_generated_at` | TIMESTAMPTZ | YES | — | — |
| `ai_agent_version` | VARCHAR(20) | YES | — | — |
| `version_no` | INTEGER | NO | `1` | Row version |
| `is_current_version` | BOOLEAN | NO | `true` | — |
| `parent_version_id` | UUID | YES | — | FK → `rubrics(id)` |
| `created_at` | TIMESTAMPTZ | NO | `NOW()` | — |
| `updated_at` | TIMESTAMPTZ | NO | `NOW()` | — |
| `approved_at` | TIMESTAMPTZ | YES | — | — |
| `approved_by` | UUID | YES | — | FK → `users(id)` |

**Indexes:** `idx_rubrics_assessment_id`, `idx_rubrics_user_id`, `idx_rubrics_status`, `idx_rubrics_created_at`, `idx_rubrics_approved_at`

---

## Table: `evidences`

| Column | Type | Nullable | Default | Constraints |
|--------|------|----------|---------|-------------|
| `id` | UUID | NO | `gen_uuid_v7()` | PRIMARY KEY |
| `student_id` | VARCHAR(50) | NO | — | External SIS ID |
| `assessment_id` | UUID | NO | — | FK → `assessments(id)` CASCADE |
| `user_id` | UUID | NO | — | FK → `users(id)` |
| `evidence_type` | VARCHAR(50) | NO | — | STUDENT_WORK, ASSESSMENT_RESULT, OBSERVATION |
| `status` | VARCHAR(20) | NO | — | COLLECTED, LINKED, EVALUATED |
| `evidence_data` | JSONB | NO | — | See Evidence Metadata Model |
| `teacher_notes` | TEXT | YES | — | — |
| `rubric_id` | UUID | YES | — | FK → `rubrics(id)` |
| `linked_criteria` | JSONB | YES | — | — |
| `evaluation_notes` | TEXT | YES | — | — |
| `created_at` | TIMESTAMPTZ | NO | `NOW()` | — |
| `updated_at` | TIMESTAMPTZ | NO | `NOW()` | — |

**Indexes:** `idx_evidences_student_id`, `idx_evidences_assessment_id`, `idx_evidences_user_id`, `idx_evidences_evidence_type`, `idx_evidences_status`, `idx_evidences_rubric_id`, `idx_evidences_created_at`

---

## Table: `evaluations`

| Column | Type | Nullable | Default | Constraints |
|--------|------|----------|---------|-------------|
| `id` | UUID | NO | `gen_uuid_v7()` | PRIMARY KEY |
| `student_id` | VARCHAR(50) | NO | — | — |
| `rubric_id` | UUID | NO | — | FK → `rubrics(id)` CASCADE |
| `evidence_id` | UUID | NO | — | FK → `evidences(id)` CASCADE |
| `user_id` | UUID | NO | — | FK → `users(id)` |
| `performance_scores` | JSONB | NO | — | — |
| `total_score` | INTEGER | NO | — | ≥ 0, ≤ max_score |
| `max_score` | INTEGER | NO | — | > 0 |
| `performance_level` | VARCHAR(20) | NO | — | EXCELLENT, PROFICIENT, DEVELOPING, BEGINNING |
| `teacher_feedback` | TEXT | YES | — | — |
| `revision_no` | INTEGER | NO | `1` | Monotonic per evidence |
| `is_current_version` | BOOLEAN | NO | `true` | — |
| `parent_revision_id` | UUID | YES | — | FK → `evaluations(id)` |
| `evaluated_at` | TIMESTAMPTZ | NO | `NOW()` | — |
| `created_at` | TIMESTAMPTZ | NO | `NOW()` | — |
| `updated_at` | TIMESTAMPTZ | NO | `NOW()` | — |

**Indexes:** `idx_evaluations_student_id`, `idx_evaluations_rubric_id`, `idx_evaluations_evidence_id`, `idx_evaluations_user_id`, `idx_evaluations_performance_level`, `idx_evaluations_evaluated_at`, `idx_evaluations_revision`, `idx_evaluations_evidence_revision`, `idx_evaluations_parent_revision`, `idx_evaluations_student_evidence_revision`, `idx_evaluations_created_at`, `idx_evaluations_updated_at`, `idx_evaluations_teacher_feedback` (partial)

---

## Table: `evaluation_feedback_history`

| Column | Type | Nullable | Default | Constraints |
|--------|------|----------|---------|-------------|
| `id` | UUID | NO | `gen_uuid_v7()` | PRIMARY KEY |
| `evaluation_id` | UUID | NO | — | FK → `evaluations(id)` CASCADE |
| `teacher_feedback` | TEXT | NO | — | — |
| `changed_by` | UUID | NO | — | FK → `users(id)` |
| `changed_at` | TIMESTAMPTZ | NO | `NOW()` | — |

**Indexes:** `idx_feedback_history_evaluation`, `idx_feedback_history_changed_at`

---

## Table: `narrative_reports`

| Column | Type | Nullable | Default | Constraints |
|--------|------|----------|---------|-------------|
| `id` | UUID | NO | `gen_uuid_v7()` | PRIMARY KEY |
| `student_id` | VARCHAR(50) | NO | — | External SIS ID |
| `class_id` | VARCHAR(50) | NO | — | External SIS ID (migration 000008) |
| `user_id` | UUID | NO | — | FK → `users(id)` |
| `status` | VARCHAR(20) | NO | — | DRAFT, APPROVED, REJECTED |
| `report_period` | JSONB | NO | — | — |
| `language` | VARCHAR(20) | NO | — | INDONESIAN, ENGLISH |
| `content` | JSONB | NO | — | — |
| `achievement_summary_id` | UUID | YES | — | Runtime reference (not FK) |
| `achievement_data` | JSONB | YES | — | Cached achievement snapshot |
| `last_achievement_calculated_at` | TIMESTAMPTZ | YES | — | — |
| `ai_confidence_score` | DECIMAL(3,2) | YES | — | 0.00–1.00 |
| `ai_generated_at` | TIMESTAMPTZ | YES | — | — |
| `ai_agent_version` | VARCHAR(20) | YES | — | — |
| `version_no` | INTEGER | NO | `1` | Row version |
| `is_current_version` | BOOLEAN | NO | `true` | — |
| `parent_version_id` | UUID | YES | — | FK → `narrative_reports(id)` |
| `created_at` | TIMESTAMPTZ | NO | `NOW()` | — |
| `updated_at` | TIMESTAMPTZ | NO | `NOW()` | — |
| `approved_at` | TIMESTAMPTZ | YES | — | — |
| `approved_by` | UUID | YES | — | FK → `users(id)` |

**Indexes:** `idx_narrative_reports_student_id`, `idx_narrative_reports_user_id`, `idx_narrative_reports_status`, `idx_narrative_reports_language`, `idx_narrative_reports_created_at`, `idx_narrative_reports_approved_at`, `idx_narrative_reports_achievement`, `idx_narrative_reports_student_period`, `idx_narrative_reports_class_id` (000008)

---

## Table: `audit_logs`

| Column | Type | Nullable | Default | Constraints |
|--------|------|----------|---------|-------------|
| `id` | UUID | NO | `gen_uuid_v7()` | PRIMARY KEY |
| `user_id` | UUID | YES | — | FK → `users(id)` |
| `action` | VARCHAR(50) | NO | — | — |
| `entity_type` | VARCHAR(50) | NO | — | — |
| `entity_id` | UUID | NO | — | — |
| `changes` | JSONB | YES | — | — |
| `ip_address` | INET | YES | — | — |
| `user_agent` | TEXT | YES | — | — |
| `created_at` | TIMESTAMPTZ | NO | `NOW()` | — |

**Indexes:** `idx_audit_logs_user_id`, `idx_audit_logs_action`, `idx_audit_logs_entity_type`, `idx_audit_logs_entity_id`, `idx_audit_logs_created_at`, `idx_audit_logs_user_id_created_at`

---

# Versioning Model

## TP Versioning

### Set-level (`tp_sets`)

| Rule | Specification |
|------|---------------|
| Scope | One AI generation / manual planning session per CP |
| Key | `UNIQUE (cp_id, version_no)` |
| Generation | `version_no = COALESCE(MAX(version_no), 0) + 1` per `cp_id` |
| Current set | Highest `version_no` with `status IN ('DRAFT', 'UNDER_REVIEW', 'APPROVED')` for active work; APPROVED sets are immutable |

### Item-level (`tp`)

| Rule | Specification |
|------|---------------|
| Trigger | Edit to approved TP content |
| Generation | Insert new row: `version_no = prior.version_no + 1`, `parent_version_id = prior.id`, set prior `is_current_version = false` |
| Current identification | `is_current_version = true` AND `tp_set_id = :set` |
| Historical query | `SELECT * FROM tp WHERE tp_set_id = :id AND sequence_number = :seq ORDER BY version_no` |
| Logical identity | `(tp_set_id, sequence_number)` chain via `parent_version_id` |

## ATP Versioning

### Set-level (`atp_sets`)

| Rule | Specification |
|------|---------------|
| Scope | One generation session per TP Set |
| Key | `UNIQUE (tp_set_id, version_no)` |
| Generation | `version_no = MAX(version_no) + 1` per `tp_set_id` |
| Prerequisite | Parent `tp_sets.status = 'APPROVED'` |

### Item-level (`atp`)

| Rule | Specification |
|------|---------------|
| Row versioning | **Not used in MVP** |
| Version context | Items inherit `atp_sets.version_no` |
| Historical query | Join `atp` to `atp_sets` on `atp_set_id` filtered by `atp_sets.version_no` |

## Modul Ajar Versioning

### Set-level (`modul_ajar_sets`)

| Rule | Specification |
|------|---------------|
| Scope | One generation session per ATP Set |
| Key | `UNIQUE (atp_set_id, version_no)` |
| Generation | `version_no = MAX(version_no) + 1` per `atp_set_id` |
| Prerequisite | Parent `atp_sets.status = 'APPROVED'` |

### Item-level (`modul_ajar`)

| Rule | Specification |
|------|---------------|
| Row versioning | **Not used in MVP** |
| Version context | Items inherit `modul_ajar_sets.version_no` |
| Historical query | Join `modul_ajar` to `modul_ajar_sets` filtered by `version_no` |

## Assessment Versioning

| Rule | Specification |
|------|---------------|
| Row versioning | `version_no`, `is_current_version`, `parent_version_id` on `assessments` |
| Snapshot (immutable) | `tp_version_no`, `success_criteria_snapshot`, `tp_*_snapshot` set at INSERT only |
| Trigger | Edit to assessment content while DRAFT; new version when editing approved (via new row) |
| Generation | `version_no = prior.version_no + 1`; prior `is_current_version = false` |
| Current identification | `is_current_version = true` per logical assessment chain |
| Historical query | `SELECT * FROM assessments WHERE tp_id = :tp AND is_current_version = false ORDER BY version_no` OR walk `parent_version_id` chain |
| Rubric versioning | Independent row versioning on `rubrics` table, same pattern |

## Evaluation Revision Tracking

| Rule | Specification |
|------|---------------|
| Scope | Revisions within single `evidence_id` |
| Generation | `revision_no = MAX(revision_no) + 1` WHERE `evidence_id = :id` |
| Current identification | `is_current_version = true` per `evidence_id` (exactly one) |
| Parent link | `parent_revision_id` → prior evaluation row |
| Feedback audit | On feedback change: INSERT `evaluation_feedback_history`; UPDATE current row |
| Historical query | `SELECT * FROM evaluations WHERE evidence_id = :id ORDER BY revision_no DESC` |
| History endpoint | Current + full chain via `revision_no` ordering |

### Version Query Patterns (Canonical SQL)

```sql
-- Current TP item
SELECT * FROM tp WHERE tp_set_id = $1 AND sequence_number = $2 AND is_current_version = true;

-- TP version history
SELECT * FROM tp WHERE tp_set_id = $1 AND sequence_number = $2 ORDER BY version_no;

-- Current assessment for TP
SELECT * FROM assessments WHERE tp_id = $1 AND is_current_version = true;

-- Current evaluation for evidence
SELECT * FROM evaluations WHERE evidence_id = $1 AND is_current_version = true;

-- Evaluation revision history
SELECT * FROM evaluations WHERE evidence_id = $1 ORDER BY revision_no DESC;
```

---

# Evidence Storage Metadata Model

Database stores **metadata only**. Binary files live in external object storage (out of scope for this document).

## `evidence_data` JSONB Schema (Frozen)

```json
{
  "description": "string, required, max 2000 chars",
  "metadata": {
    "collected_at": "ISO8601, required",
    "collection_context": "string, optional",
    "subject_code": "string, optional",
    "tags": ["string"]
  },
  "files": [
    {
      "file_id": "uuid, required",
      "storage_reference": {
        "bucket": "string, required",
        "object_key": "string, required"
      },
      "original_filename": "string, required",
      "mime_type": "string, required",
      "size_bytes": "integer, required, max 10485760",
      "uploaded_at": "ISO8601, required",
      "uploaded_by": "uuid, required"
    }
  ],
  "integrity": {
    "checksum_sha256": "string, hex, required per file",
    "verified_at": "ISO8601, optional"
  },
  "retention": {
    "retain_until": "ISO8601, optional",
    "retention_policy": "STANDARD | EXTENDED, default STANDARD"
  }
}
```

## Field Definitions

### Metadata Fields

| Field | Location | Purpose |
|-------|----------|---------|
| `description` | root | Human-readable evidence summary |
| `metadata.collected_at` | metadata | When evidence was captured |
| `metadata.collection_context` | metadata | Classroom activity context |
| `metadata.tags` | metadata | Searchable labels |
| `teacher_notes` | table column | Teacher commentary (separate from JSONB) |
| `linked_criteria` | table column | Rubric criteria mapping |
| `evaluation_notes` | table column | Post-evaluation notes |

### Storage Reference Fields

| Field | Purpose |
|-------|---------|
| `files[].file_id` | Stable identifier for DB ↔ storage correlation |
| `files[].storage_reference.bucket` | Object storage bucket name |
| `files[].storage_reference.object_key` | Opaque storage path (school-scoped by convention) |
| `files[].mime_type` | Validation record |
| `files[].size_bytes` | Quota enforcement record |
| `files[].uploaded_by` | Provenance |

### File Integrity Fields

| Field | Purpose |
|-------|---------|
| `integrity.checksum_sha256` | Detect tampering / corruption |
| `integrity.verified_at` | Last successful verification timestamp |

### Retention Fields

| Field | Purpose |
|-------|---------|
| `retention.retain_until` | Earliest eligible purge date (application-enforced) |
| `retention.retention_policy` | STANDARD = 7 years (education compliance default); EXTENDED = manual hold |

**MVP rule:** No automated purge job — retention fields are populated for future governance only.

---

# Migration Freeze

## Approved Sequence (LOCKED)

No migration may be added, reordered, or modified after freeze approval without **Architecture Board** written approval.

| # | File | Purpose | Tables Affected |
|---|------|---------|-----------------|
| **000001** | `init_schema` | Platform foundation, UUID v7, auth | `schools`, `roles`, `permissions`, `users`, `refresh_tokens`, `ai_generation_logs` |
| **000002** | `education_domain_tables` | Curriculum hierarchy, planning, assessment, reporting, audit | All `curriculum_*`, `cp`, `tp_sets`, `tp`, `atp_sets`, `atp`, `modul_ajar_sets`, `modul_ajar`, `assessments`, `rubrics`, `evidences`, `evaluations`, `narrative_reports`, `audit_logs` |
| **000003** | `success_criteria_and_assessment_refactor` | KKTPCriteria on TP; assessment→TP; evaluation base revision columns | `tp`, `assessments`, `evaluations` |
| **000004** | `evaluation_revision_tracking` | Evaluation version flags + feedback history | `evaluations`, `evaluation_feedback_history` |
| **000005** | `achievement_reports` | Achievement cache on reports | `narrative_reports` |
| **000006** | `tp_item_versioning` | TP row versioning (**corrected: table `tp`**) | `tp` |
| **000007** | `assessment_snapshot_expand` | Expanded TP snapshots on assessment | `assessments` |
| **000008** | `narrative_reports_class_id` | Close domain/DB gap for class reference | `narrative_reports` |

## Mandatory Corrections Before Production Apply

| Migration | Issue | Required Fix |
|-----------|-------|--------------|
| 000003 | UPDATE sets `success_criteria` on assessments | Must set `success_criteria_snapshot` |
| 000004 | Re-adds `revision_no` already in 000003 | Make idempotent: skip if column exists |
| 000006 | `ALTER TABLE tps` | Change to `ALTER TABLE tp`; index names `idx_tp_*` |
| 000004 | `gen_random_uuid()` on feedback history | Change to `gen_uuid_v7()` |

## Explicitly Excluded from MVP Migration Sequence

| Proposed Migration | Status |
|--------------------|--------|
| `curriculum_versions` | Wave 2 — Architecture Board only |
| `workflow_history` | Wave 2 — Architecture Board only |
| `students` / `classes` tables | Wave 2 — Architecture Board only |
| `achievement_*` tables | **Permanently excluded** — runtime computation |
| Any CQRS/projection/event tables | **Forbidden** |

## Rollback Policy

- Every `.up.sql` must have paired `.down.sql`
- Production rollback: reverse-order down migrations only with Board approval
- Data migrations (000003) require backup before apply

---

# Index Strategy

## Complete Index Inventory

### Identity & Platform (17)

`idx_schools_code`, `idx_schools_is_active`, `idx_roles_name`, `idx_roles_is_active`, `idx_permissions_role_id`, `idx_permissions_resource`, `idx_users_email`, `idx_users_role_id`, `idx_users_school_id`, `idx_users_is_active`, `idx_users_created_at`, `idx_refresh_tokens_user_id`, `idx_refresh_tokens_token`, `idx_refresh_tokens_expires_at`, `idx_ai_generation_logs_user_id`, `idx_ai_generation_logs_school_id`, `idx_ai_generation_logs_artifact`, `idx_ai_generation_logs_created_at`

### Curriculum (14)

`idx_curriculum_subjects_code`, `idx_curriculum_subjects_is_active`, `idx_curriculum_phases_code`, `idx_curriculum_phases_is_active`, `idx_curriculum_elements_subject_id`, `idx_curriculum_elements_phase_id`, `idx_curriculum_elements_code`, `idx_curriculum_elements_is_active`, `idx_curriculum_subelements_element_id`, `idx_curriculum_subelements_code`, `idx_curriculum_subelements_is_active`, `idx_cp_subject_id`, `idx_cp_phase_id`, `idx_cp_element_id`, `idx_cp_subelement_id`, `idx_cp_code`, `idx_cp_version`, `idx_cp_is_active`

### Learning Planning (28)

Set tables: status, version, generated_by, approved_by, ai_generation_id indexes per set table.

`tp`: `idx_tp_tp_set_id`, `idx_tp_sequence_number`, `idx_tp_tp_set_id_sequence`, hierarchy indexes, `idx_tp_status`, `idx_tp_user_id`, `idx_tp_success_criteria` (GIN), `idx_tp_set_version`, `idx_tp_parent_version`

`atp`, `modul_ajar`: set_id, parent FK, status, created_at indexes.

### Assessment (25)

Assessment TP/snapshot indexes, rubric indexes, evidence indexes, evaluation revision indexes (see table definitions).

### Reporting (8)

`idx_narrative_reports_*` including `idx_narrative_reports_student_period`, `idx_narrative_reports_achievement`, `idx_narrative_reports_class_id`

### Audit (6)

`idx_audit_logs_*` composite on `(user_id, created_at)`

## Index Design Rules

| Rule | Specification |
|------|---------------|
| FK columns | Always indexed |
| Tenant queries | `users.school_id` indexed; artifact queries join through `user_id` |
| JSONB search | GIN on `success_criteria`, `success_criteria_snapshot` only (MVP) |
| Partial indexes | Allowed for nullable filter columns (e.g., `teacher_feedback IS NOT NULL`) |
| Unique constraints | Implicit unique indexes on all UNIQUE constraints |

---

# Query Optimization Requirements

MVP-required optimizations only. No materialized views, no read replicas, no caching layer.

| Query Pattern | Required Optimization |
|---------------|----------------------|
| List artifacts by school | `JOIN users u ON a.user_id = u.id WHERE u.school_id = $1` — requires `idx_users_school_id` + `idx_*_user_id` |
| Current version lookups | `WHERE is_current_version = true` — consider partial index `WHERE is_current_version` on high-volume tables (optional MVP) |
| TP set by CP | `idx_tp_sets_cp_id` + `uq_tp_sets_cp_version` |
| Assessment by TP | `idx_assessments_tp_id`, `idx_assessments_tp_version` |
| Evaluation history | `idx_evaluations_student_evidence_revision` |
| Narrative report by student+period | `idx_narrative_reports_student_period` |
| Audit trail by user | `idx_audit_logs_user_id_created_at` |
| AI cost tracking | `idx_ai_generation_logs_school_id`, `idx_ai_generation_logs_created_at` |
| Pagination | `ORDER BY created_at DESC` with `idx_*_created_at` |

**Forbidden optimizations:** CQRS read models, projection tables, event replay queries, cross-aggregate materialized views.

---

# Schema Risks

## Future Migration Risks

| Risk | Impact | Mitigation |
|------|--------|------------|
| `curriculum_versions` addition | FK backfill on all hierarchy + artifacts | Deferred to Wave 2; `cp.version` string suffices for MVP |
| `students`/`classes` normalization | `student_id`/`class_id` VARCHAR migration to UUID FK | Plan nullable transition column; Board approval required |
| JSONB schema evolution | Breaking changes to `evidence_data`, `content` | Version field inside JSONB documents |
| Index bloat on version tables | Storage growth from row versioning | Archival policy via `status = 'ARCHIVED'`; no physical delete |
| Migration 000003–000007 corrections | Fresh vs existing DB divergence | Apply mandatory corrections before any new environment |

## Data Integrity Risks

| Risk | Impact | Mitigation |
|------|--------|------------|
| Missing school boundary JOIN | Cross-tenant data leak | Mandatory service-layer validator; integration tests |
| Orphan evaluations | `is_current_version` out of sync | Transaction: unset prior + insert new in single TX |
| Multiple current versions | Query ambiguity | UNIQUE partial index (Wave 2) or app-level constraint check |
| Assessment without snapshot | Historical inconsistency | NOT NULL on `success_criteria_snapshot`; service captures at INSERT |
| CASCADE delete on approved chain | Accidental data loss | RESTRICT on user delete; soft-archive instead of delete |
| `student_id` without school validation | Cross-school evidence | Service validates student belongs to teacher's school |

## Scaling Risks

| Risk | Threshold | Mitigation |
|------|-----------|------------|
| JSONB column size | Large assessment_items / content blobs | Monitor row size; externalize large media to object storage metadata only |
| `audit_logs` growth | Unbounded append | Partition by `created_at` (Wave 2); retention policy |
| Version row proliferation | TP/assessment revisions | Archive old versions; index maintenance |
| Single DB multi-tenant | Connection pool pressure | Connection pooling (PgBouncer) — ops concern, not schema |
| GIN index maintenance | Write amplification on TP/assessments | Limit GIN to required columns only |

---

# Final Freeze Decisions

The following decisions are **immutable** after approval of this document:

1. **Single PostgreSQL database**, schema `public`, modular monolith — no per-tenant schemas.
2. **UUID v7** (via `gen_uuid_v7()`) for all primary keys — no serial integers, no `gen_random_uuid()`.
3. **Table name `tp`** — not `tps`, `teaching_plans`, or `tujuan_pembelajaran`.
4. **Assessment references `tp_id`** — `modul_ajar_id` on assessments is permanently removed.
5. **Success criteria embedded in `tp.success_criteria` JSONB** — no separate `kk_tp` or `success_criteria` table.
6. **Set-level versioning** on `tp_sets`, `atp_sets`, `modul_ajar_sets` via `version_no`.
7. **Row-level versioning** on `tp`, `assessments`, `rubrics`, `narrative_reports` via `version_no` + `is_current_version` + `parent_version_id`.
8. **ATP and Modul Ajar items** do not get row-level version columns in MVP — set `version_no` is the version boundary.
9. **Evaluation revisions** use `revision_no` + `is_current_version` + `parent_revision_id` — separate from artifact versioning.
10. **No `students` or `classes` tables** in MVP — `VARCHAR(50)` external IDs only.
11. **No `achievement_*` tables** — achievement computed at runtime, cached in `narrative_reports.achievement_data`.
12. **No `deleted_at` columns** — deactivation via `is_active` or `status = 'ARCHIVED'`.
13. **School isolation via `users.school_id`** — no denormalized `school_id` on artifact tables (except `ai_generation_logs`).
14. **Eight migrations maximum** for MVP (`000001`–`000008`) — further changes require Architecture Board approval.
15. **Evidence file metadata in `evidence_data` JSONB** — no `evidence_files` relational table in MVP.
16. **No event store, workflow_history, or curriculum_versions tables** in MVP migration sequence.
17. **`narrative_reports.class_id` is required** — delivered via migration `000008`.
18. **Workflow status enums frozen** — set tables use 5 states; assessment/rubric/report use 3 states (DRAFT, APPROVED, REJECTED).
19. **Immutability after approval** — approved rows are not UPDATEd; revisions insert new rows.
20. **Audit via `audit_logs` + `evaluation_feedback_history`** — no event sourcing substitute.

---

**Document Status:** LOCKED  
**Database Schema Freeze v1:** ACTIVE  
**Parent Architecture:** NUSA Architecture Freeze v2 — APPROVED  
**Next Action:** Apply mandatory migration corrections (`000003`, `000004`, `000006`) and implement `000008` before narrative reporting production deploy.

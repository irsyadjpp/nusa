# 14_DATABASE_SCHEMA.md

## Foundation Document for NUSA Education Platform

**Version**: 1.0
**Date**: June 2026
**Status**: FOUNDATION DOCUMENT
**Alignment**: Validated against Foundation Architecture (00A, 00B, 00C, 01, 02, 03, 04, 05, 06, 07, 08, 09, 10, 11, 12, 13)

**Purpose**: Define the physical MVP database schema for NUSA Wave 1, serving as the official database design document. This document is implementation-ready for PostgreSQL migration generation and provides the complete table definitions, relationships, constraints, and indexes.

---

# SECTION 1 — Database Overview

## Audit Field Standard

All transactional entities should support audit fields for traceability and governance.

## Mandatory Fields

All transactional entities must include:

- **created_at**: Timestamp when the record was created
- **updated_at**: Timestamp when the record was last updated

## Recommended Fields

Transactional entities should include where operationally justified:

- **created_by**: User ID of the user who created the record
- **updated_by**: User ID of the user who last updated the record

## Optional Fields

Transactional entities may include for soft delete functionality:

- **deleted_at**: Timestamp when the record was soft deleted
- **deleted_by**: User ID of the user who soft deleted the record

## Approval-Based Entities

Entities that require human approval (artifacts, configurations, policies) additionally support:

- **approved_by**: User ID of the user who approved the record
- **approved_at**: Timestamp when the record was approved

## Implementation Guidance

MVP implementation may introduce audit fields incrementally where operationally justified.

Do not force unnecessary schema complexity. Apply audit fields based on actual governance and traceability requirements.

The centralized audit_logs table provides system-wide audit trail regardless of individual entity audit field implementation.

---

## Database Information

- **Database Name**: `nusa`
- **Database Engine**: PostgreSQL 18+
- **Character Set**: UTF8
- **Collation**: en_US.UTF-8
- **Timezone**: UTC
- **UUID Version**: UUID v7 (time-ordered, sortable)

## Schema Organization

The database uses a single schema `public` for MVP Wave 1. Tables are organized by module:
- Authentication Module: `users`, `roles`, `permissions`, `refresh_tokens`
- Curriculum Module: `cp`, `tp`
- Learning Planning Module: `atp`, `modul_ajar`
- Assessment Module: `assessments`, `rubrics`, `evidences`, `evaluations`
- Reporting Module: `narrative_reports`
- Administration Module: `audit_logs`

## Naming Conventions

- **Table Names**: `snake_case` (e.g., `teaching_plans`)
- **Column Names**: `snake_case` (e.g., `created_at`)
- **Primary Keys**: `id` (UUID v7)
- **Foreign Keys**: `{table}_id` (e.g., `user_id`)
- **Timestamps**: `created_at`, `updated_at`
- **Boolean Columns**: `is_{status}` (e.g., `is_active`)
- **UUID Generation**: `gen_uuid_v7()` function for all UUID columns

---

# SECTION 2 — Authentication Module Tables

## Table: users

**Module Ownership**: Authentication Module

**Description**: Stores user account information for teachers and administrators.

### Columns

| Column | Data Type | Constraints | Description |
|--------|-----------|-------------|-------------|
| `id` | UUID | PRIMARY KEY | Unique user identifier |
| `email` | VARCHAR(255) | UNIQUE, NOT NULL | User email address |
| `password_hash` | VARCHAR(255) | NOT NULL | Bcrypt hashed password |
| `name` | VARCHAR(255) | NOT NULL | User full name |
| `role_id` | UUID | FOREIGN KEY → roles(id), NOT NULL | User role reference |
| `is_active` | BOOLEAN | NOT NULL, DEFAULT true | Account active status |
| `failed_login_attempts` | INTEGER | NOT NULL, DEFAULT 0 | Failed login attempt count |
| `locked_until` | TIMESTAMP WITH TIME ZONE | NULLABLE | Account lock expiration |
| `created_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Account creation timestamp |
| `updated_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Last update timestamp |
| `created_by` | UUID | FOREIGN KEY → users(id) | User who created account |
| `updated_by` | UUID | FOREIGN KEY → users(id) | User who last updated account |

### Indexes

```sql
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_role_id ON users(role_id);
CREATE INDEX idx_users_is_active ON users(is_active);
CREATE INDEX idx_users_created_at ON users(created_at);
```

### Constraints

```sql
ALTER TABLE users ADD CONSTRAINT chk_users_email_format CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$');
ALTER TABLE users ADD CONSTRAINT chk_users_failed_login_attempts CHECK (failed_login_attempts >= 0);
```

---

## Table: roles

**Module Ownership**: Authentication Module

**Description**: Defines user roles and their permissions.

### Columns

| Column | Data Type | Constraints | Description |
|--------|-----------|-------------|-------------|
| `id` | UUID | PRIMARY KEY | Unique role identifier |
| `name` | VARCHAR(50) | UNIQUE, NOT NULL | Role name (TEACHER, ADMINISTRATOR) |
| `description` | TEXT | NULLABLE | Role description |
| `is_active` | BOOLEAN | NOT NULL, DEFAULT true | Role active status |
| `created_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Role creation timestamp |
| `updated_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Last update timestamp |

### Indexes

```sql
CREATE INDEX idx_roles_name ON roles(name);
CREATE INDEX idx_roles_is_active ON roles(is_active);
```

---

## Table: permissions

**Module Ownership**: Authentication Module

**Description**: Defines permissions for each role.

### Columns

| Column | Data Type | Constraints | Description |
|--------|-----------|-------------|-------------|
| `id` | UUID | PRIMARY KEY | Unique permission identifier |
| `role_id` | UUID | FOREIGN KEY → roles(id), NOT NULL | Role reference |
| `resource` | VARCHAR(100) | NOT NULL | Resource identifier (e.g., "curriculum.tp") |
| `action` | VARCHAR(50) | NOT NULL | Action (CREATE, READ, UPDATE, DELETE, APPROVE) |
| `created_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Permission creation timestamp |

### Indexes

```sql
CREATE INDEX idx_permissions_role_id ON permissions(role_id);
CREATE INDEX idx_permissions_resource ON permissions(resource);
CREATE UNIQUE INDEX idx_permissions_role_resource_action ON permissions(role_id, resource, action);
```

---

## Table: refresh_tokens

**Module Ownership**: Authentication Module

**Description**: Stores JWT refresh tokens for token refresh functionality.

### Columns

| Column | Data Type | Constraints | Description |
|--------|-----------|-------------|-------------|
| `id` | UUID | PRIMARY KEY | Unique token identifier |
| `user_id` | UUID | FOREIGN KEY → users(id), NOT NULL | User reference |
| `token` | TEXT | UNIQUE, NOT NULL | Refresh token string |
| `expires_at` | TIMESTAMP WITH TIME ZONE | NOT NULL | Token expiration timestamp |
| `revoked_at` | TIMESTAMP WITH TIME ZONE | NULLABLE | Token revocation timestamp |
| `created_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Token creation timestamp |
| `ip_address` | INET | NULLABLE | IP address used for token generation |

### Indexes

```sql
CREATE INDEX idx_refresh_tokens_user_id ON refresh_tokens(user_id);
CREATE INDEX idx_refresh_tokens_token ON refresh_tokens(token);
CREATE INDEX idx_refresh_tokens_expires_at ON refresh_tokens(expires_at);
```

### Constraints

```sql
ALTER TABLE refresh_tokens ADD CONSTRAINT chk_refresh_tokens_expires_at CHECK (expires_at > created_at);
```

---

# SECTION 3 — Curriculum Module Tables

## Table: curriculum_subjects

**Module Ownership**: Curriculum Module

**Description**: Stores curriculum subjects (Mata Pelajaran) as the top level of the curriculum hierarchy.

### Columns

| Column | Data Type | Constraints | Description |
|--------|-----------|-------------|-------------|
| `id` | UUID | PRIMARY KEY | Unique subject identifier |
| `code` | VARCHAR(50) | NOT NULL, UNIQUE | Subject code (e.g., MTK, BIN, IPA) |
| `name` | VARCHAR(255) | NOT NULL | Subject name (e.g., Matematika, Bahasa Indonesia) |
| `name_en` | VARCHAR(255) | NULLABLE | Subject name in English |
| `description` | TEXT | NULLABLE | Subject description |
| `is_active` | BOOLEAN | NOT NULL, DEFAULT true | Subject active status |
| `created_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Creation timestamp |
| `updated_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Last update timestamp |

### Indexes

```sql
CREATE INDEX idx_curriculum_subjects_code ON curriculum_subjects(code);
CREATE INDEX idx_curriculum_subjects_is_active ON curriculum_subjects(is_active);
```

---

## Table: curriculum_phases

**Module Ownership**: Curriculum Module

**Description**: Stores curriculum phases (Fase) as the second level of the curriculum hierarchy.

### Columns

| Column | Data Type | Constraints | Description |
|--------|-----------|-------------|-------------|
| `id` | UUID | PRIMARY KEY | Unique phase identifier |
| `code` | VARCHAR(50) | NOT NULL, UNIQUE | Phase code (e.g., FASE_FONDASI, FASE_A, FASE_B) |
| `name` | VARCHAR(255) | NOT NULL | Phase name (e.g., Fase Fondasi, Fase A) |
| `name_en` | VARCHAR(255) | NULLABLE | Phase name in English |
| `description` | TEXT | NULLABLE | Phase description |
| `grade_level_start` | INTEGER | NULLABLE | Starting grade level for this phase |
| `grade_level_end` | INTEGER | NULLABLE | Ending grade level for this phase |
| `is_active` | BOOLEAN | NOT NULL, DEFAULT true | Phase active status |
| `created_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Creation timestamp |
| `updated_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Last update timestamp |

### Indexes

```sql
CREATE INDEX idx_curriculum_phases_code ON curriculum_phases(code);
CREATE INDEX idx_curriculum_phases_is_active ON curriculum_phases(is_active);
```

---

## Table: curriculum_elements

**Module Ownership**: Curriculum Module

**Description**: Stores curriculum elements (Elemen) as the third level of the curriculum hierarchy, grouped by subject and phase.

### Columns

| Column | Data Type | Constraints | Description |
|--------|-----------|-------------|-------------|
| `id` | UUID | PRIMARY KEY | Unique element identifier |
| `subject_id` | UUID | FOREIGN KEY → curriculum_subjects(id), NOT NULL | Subject reference |
| `phase_id` | UUID | FOREIGN KEY → curriculum_phases(id), NOT NULL | Phase reference |
| `code` | VARCHAR(50) | NOT NULL | Element code (e.g., NUM_OPS, ALGEBRA, GEOM) |
| `name` | VARCHAR(255) | NOT NULL | Element name (e.g., Number and Operations, Algebra) |
| `name_en` | VARCHAR(255) | NULLABLE | Element name in English |
| `description` | TEXT | NULLABLE | Element description |
| `is_active` | BOOLEAN | NOT NULL, DEFAULT true | Element active status |
| `created_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Creation timestamp |
| `updated_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Last update timestamp |

### Indexes

```sql
CREATE INDEX idx_curriculum_elements_subject_id ON curriculum_elements(subject_id);
CREATE INDEX idx_curriculum_elements_phase_id ON curriculum_elements(phase_id);
CREATE INDEX idx_curriculum_elements_code ON curriculum_elements(code);
CREATE INDEX idx_curriculum_elements_is_active ON curriculum_elements(is_active);
CREATE UNIQUE INDEX idx_curriculum_elements_subject_phase_code ON curriculum_elements(subject_id, phase_id, code);
```

---

## Table: curriculum_subelements

**Module Ownership**: Curriculum Module

**Description**: Stores curriculum subelements (Subelemen) as the fourth level of the curriculum hierarchy, grouped by element.

### Columns

| Column | Data Type | Constraints | Description |
|--------|-----------|-------------|-------------|
| `id` | UUID | PRIMARY KEY | Unique subelement identifier |
| `element_id` | UUID | FOREIGN KEY → curriculum_elements(id), NOT NULL | Element reference |
| `code` | VARCHAR(50) | NOT NULL | Subelement code (e.g., WHOLE_NUM, FRACTIONS, DECIMALS) |
| `name` | VARCHAR(255) | NOT NULL | Subelement name (e.g., Whole Numbers, Fractions, Decimals) |
| `name_en` | VARCHAR(255) | NULLABLE | Subelement name in English |
| `description` | TEXT | NULLABLE | Subelement description |
| `is_active` | BOOLEAN | NOT NULL, DEFAULT true | Subelement active status |
| `created_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Creation timestamp |
| `updated_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Last update timestamp |

### Indexes

```sql
CREATE INDEX idx_curriculum_subelements_element_id ON curriculum_subelements(element_id);
CREATE INDEX idx_curriculum_subelements_code ON curriculum_subelements(code);
CREATE INDEX idx_curriculum_subelements_is_active ON curriculum_subelements(is_active);
CREATE UNIQUE INDEX idx_curriculum_subelements_element_code ON curriculum_subelements(element_id, code);
```

---

## Table: cp (Curriculum Plans)

**Module Ownership**: Curriculum Module

**Description**: Stores national curriculum plan data (Capaian Pembelajaran) imported from government sources, with full hierarchy traceability.

**Data Model Mapping Note**: Database structures and AI prompt payload structures may differ. Application-layer mapping is responsible for transforming between Database Format and Prompt Payload Format. For example, the database stores `time_allocation_hours` and `hours_per_week` as separate columns, while the AI prompt expects a nested `time_allocation` object with `total_hours` and `hours_per_week`. This transformation is intentional and does not require schema redesign.

### Columns

| Column | Data Type | Constraints | Description |
|--------|-----------|-------------|-------------|
| `id` | UUID | PRIMARY KEY | Unique CP identifier |
| `subject_id` | UUID | FOREIGN KEY → curriculum_subjects(id), NOT NULL | Subject reference |
| `phase_id` | UUID | FOREIGN KEY → curriculum_phases(id), NOT NULL | Phase reference |
| `element_id` | UUID | FOREIGN KEY → curriculum_elements(id), NOT NULL | Element reference |
| `subelement_id` | UUID | FOREIGN KEY → curriculum_subelements(id), NOT NULL | Subelement reference |
| `code` | VARCHAR(50) | NOT NULL | CP code (e.g., CP.10.1.1) |
| `description` | TEXT | NOT NULL | CP description (learning achievement expectation) |
| `competency_code` | VARCHAR(50) | NULLABLE | Competency code reference |
| `learning_objectives` | JSONB | NOT NULL | Learning objectives array |
| `competency_standards` | JSONB | NOT NULL | Competency standards array |
| `time_allocation_hours` | INTEGER | NOT NULL | Total time allocation in hours |
| `hours_per_week` | INTEGER | NOT NULL | Hours per week |
| `version` | VARCHAR(20) | NOT NULL | Curriculum version |
| `is_active` | BOOLEAN | NOT NULL, DEFAULT true | CP active status |
| `imported_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Import timestamp |
| `imported_by` | UUID | FOREIGN KEY → users(id) | User who imported CP |

### Indexes

```sql
CREATE INDEX idx_cp_subject_id ON cp(subject_id);
CREATE INDEX idx_cp_phase_id ON cp(phase_id);
CREATE INDEX idx_cp_element_id ON cp(element_id);
CREATE INDEX idx_cp_subelement_id ON cp(subelement_id);
CREATE INDEX idx_cp_code ON cp(code);
CREATE INDEX idx_cp_version ON cp(version);
CREATE INDEX idx_cp_is_active ON cp(is_active);
CREATE UNIQUE INDEX idx_cp_hierarchy_code ON cp(subject_id, phase_id, element_id, subelement_id, code);
```

### Constraints

```sql
ALTER TABLE cp ADD CONSTRAINT chk_cp_time_allocation_hours CHECK (time_allocation_hours > 0);
ALTER TABLE cp ADD CONSTRAINT chk_cp_hours_per_week CHECK (hours_per_week > 0 AND hours_per_week <= 40);
```

---

## Table: tp_sets (Teaching Plan Sets)

**Module Ownership**: Curriculum Module

**Description**: Represents one AI generation session for a CP, containing multiple TP Items. A TP Set is the first-class domain entity for TP generation workflow.

### Columns

| Column | Data Type | Constraints | Description |
|--------|-----------|-------------|-------------|
| `id` | UUID | PRIMARY KEY | Unique TP Set identifier |
| `cp_id` | UUID | FOREIGN KEY → cp(id), NOT NULL | CP reference |
| `version_no` | INTEGER | NOT NULL, DEFAULT 1 | Version number for this CP |
| `status` | VARCHAR(20) | NOT NULL | Status (DRAFT, UNDER_REVIEW, APPROVED, REJECTED, ARCHIVED) |
| `generation_source` | VARCHAR(50) | NOT NULL | Generation source (AI_GENERATED, MANUAL) |
| `generation_reason` | TEXT | NULLABLE | Reason for generation (e.g., "Initial", "Regeneration after CP update") |
| `generated_by` | UUID | FOREIGN KEY → users(id), NOT NULL | User who requested generation |
| `ai_generation_id` | UUID | NULLABLE | Reference to ai_generation_logs (if AI-generated) |
| `approved_by` | UUID | FOREIGN KEY → users(id) | User who approved TP Set |
| `approved_at` | TIMESTAMP WITH TIME ZONE | NULLABLE | Approval timestamp |
| `created_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | TP Set creation timestamp |
| `updated_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Last update timestamp |

### Indexes

```sql
CREATE INDEX idx_tp_sets_cp_id ON tp_sets(cp_id);
CREATE INDEX idx_tp_sets_version_no ON tp_sets(version_no);
CREATE INDEX idx_tp_sets_status ON tp_sets(status);
CREATE INDEX idx_tp_sets_generated_by ON tp_sets(generated_by);
CREATE INDEX idx_tp_sets_approved_by ON tp_sets(approved_by);
CREATE INDEX idx_tp_sets_ai_generation_id ON tp_sets(ai_generation_id);
CREATE UNIQUE INDEX idx_tp_sets_cp_version ON tp_sets(cp_id, version_no);
```

---

## Table: tp (Teaching Plan Items)

**Module Ownership**: Curriculum Module

**Description**: Stores individual teaching plan items within a TP Set, with full curriculum traceability.

### Columns

| Column | Data Type | Constraints | Description |
|--------|-----------|-------------|-------------|
| `id` | UUID | PRIMARY KEY | Unique TP identifier |
| `tp_set_id` | UUID | FOREIGN KEY → tp_sets(id), NOT NULL | TP Set reference |
| `sequence_number` | INTEGER | NOT NULL | Sequence within TP Set |
| `cp_id` | UUID | FOREIGN KEY → cp(id), NOT NULL | CP reference |
| `subject_id` | UUID | FOREIGN KEY → curriculum_subjects(id), NOT NULL | Subject reference (for traceability) |
| `phase_id` | UUID | FOREIGN KEY → curriculum_phases(id), NOT NULL | Phase reference (for traceability) |
| `element_id` | UUID | FOREIGN KEY → curriculum_elements(id), NOT NULL | Element reference (for traceability) |
| `subelement_id` | UUID | FOREIGN KEY → curriculum_subelements(id), NOT NULL | Subelement reference (for traceability) |
| `user_id` | UUID | FOREIGN KEY → users(id), NOT NULL | Teacher who created TP |
| `status` | VARCHAR(20) | NOT NULL | Status (DRAFT, UNDER_REVIEW, APPROVED, REJECTED, ARCHIVED) |
| `title` | VARCHAR(255) | NULLABLE | TP title |
| `learning_objectives` | JSONB | NOT NULL | Learning objectives array |
| `time_allocation` | JSONB | NOT NULL | Time allocation object |
| `prerequisites` | JSONB | NULLABLE | Prerequisite relationships |
| `estimated_weeks` | INTEGER | NULLABLE | Estimated weeks for this TP |
| `created_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | TP creation timestamp |
| `updated_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Last update timestamp |

### Indexes

```sql
CREATE INDEX idx_tp_tp_set_id ON tp(tp_set_id);
CREATE INDEX idx_tp_sequence_number ON tp(tp_set_id, sequence_number);
CREATE INDEX idx_tp_cp_id ON tp(cp_id);
CREATE INDEX idx_tp_subject_id ON tp(subject_id);
CREATE INDEX idx_tp_phase_id ON tp(phase_id);
CREATE INDEX idx_tp_element_id ON tp(element_id);
CREATE INDEX idx_tp_subelement_id ON tp(subelement_id);
CREATE INDEX idx_tp_user_id ON tp(user_id);
CREATE INDEX idx_tp_status ON tp(status);
CREATE INDEX idx_tp_created_at ON tp(created_at);
```

### Constraints

```sql
ALTER TABLE tp ADD CONSTRAINT chk_tp_sequence_number CHECK (sequence_number > 0);
ALTER TABLE tp ADD CONSTRAINT chk_tp_estimated_weeks CHECK (estimated_weeks IS NULL OR estimated_weeks > 0);
ALTER TABLE tp ADD CONSTRAINT chk_tp_status CHECK (status IN ('DRAFT', 'UNDER_REVIEW', 'APPROVED', 'REJECTED', 'ARCHIVED'));
```

---

# SECTION 4 — Learning Planning Module Tables

## Table: atp_sets (Annual Teaching Plan Sets)

**Module Ownership**: Learning Planning Module

**Description**: Represents one AI generation session for ATP, containing multiple ATP Items. An ATP Set is the first-class domain entity for ATP generation workflow.

### Columns

| Column | Data Type | Constraints | Description |
|--------|-----------|-------------|-------------|
| `id` | UUID | PRIMARY KEY | Unique ATP Set identifier |
| `tp_set_id` | UUID | FOREIGN KEY → tp_sets(id), NOT NULL | TP Set reference |
| `version_no` | INTEGER | NOT NULL, DEFAULT 1 | Version number for this TP Set |
| `status` | VARCHAR(20) | NOT NULL | Status (DRAFT, UNDER_REVIEW, APPROVED, REJECTED, ARCHIVED) |
| `generation_source` | VARCHAR(50) | NOT NULL | Generation source (AI_GENERATED, MANUAL) |
| `generation_reason` | TEXT | NULLABLE | Reason for generation |
| `generated_by` | UUID | FOREIGN KEY → users(id), NOT NULL | User who requested generation |
| `ai_generation_id` | UUID | NULLABLE | Reference to ai_generation_logs (if AI-generated) |
| `approved_by` | UUID | FOREIGN KEY → users(id) | User who approved ATP Set |
| `approved_at` | TIMESTAMP WITH TIME ZONE | NULLABLE | Approval timestamp |
| `created_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | ATP Set creation timestamp |
| `updated_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Last update timestamp |

### Indexes

```sql
CREATE INDEX idx_atp_sets_tp_set_id ON atp_sets(tp_set_id);
CREATE INDEX idx_atp_sets_version_no ON atp_sets(version_no);
CREATE INDEX idx_atp_sets_status ON atp_sets(status);
CREATE INDEX idx_atp_sets_generated_by ON atp_sets(generated_by);
CREATE INDEX idx_atp_sets_approved_by ON atp_sets(approved_by);
CREATE INDEX idx_atp_sets_ai_generation_id ON atp_sets(ai_generation_id);
CREATE UNIQUE INDEX idx_atp_sets_tp_version ON atp_sets(tp_set_id, version_no);
```

---

## Table: atp (Annual Teaching Plan Items)

**Module Ownership**: Learning Planning Module

**Description**: Stores individual annual teaching plan items within an ATP Set.

### Columns

| Column | Data Type | Constraints | Description |
|--------|-----------|-------------|-------------|
| `id` | UUID | PRIMARY KEY | Unique ATP identifier |
| `atp_set_id` | UUID | FOREIGN KEY → atp_sets(id), NOT NULL | ATP Set reference |
| `tp_id` | UUID | FOREIGN KEY → tp(id), NOT NULL | TP reference |
| `user_id` | UUID | FOREIGN KEY → users(id), NOT NULL | Teacher who created ATP |
| `status` | VARCHAR(20) | NOT NULL | Status (DRAFT, UNDER_REVIEW, APPROVED, REJECTED, ARCHIVED) |
| `academic_calendar` | JSONB | NOT NULL | Academic calendar object |
| `class_schedule` | JSONB | NOT NULL | Class schedule object |
| `weekly_sequence` | JSONB | NOT NULL | Weekly sequence array |
| `assessment_schedule` | JSONB | NULLABLE | Assessment schedule array |
| `created_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | ATP creation timestamp |
| `updated_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Last update timestamp |

### Indexes

```sql
CREATE INDEX idx_atp_atp_set_id ON atp(atp_set_id);
CREATE INDEX idx_atp_tp_id ON atp(tp_id);
CREATE INDEX idx_atp_user_id ON atp(user_id);
CREATE INDEX idx_atp_status ON atp(status);
CREATE INDEX idx_atp_created_at ON atp(created_at);
```

### Constraints

```sql
ALTER TABLE atp ADD CONSTRAINT chk_atp_status CHECK (status IN ('DRAFT', 'UNDER_REVIEW', 'APPROVED', 'REJECTED', 'ARCHIVED'));
```

---

## Table: modul_ajar_sets (Modul Ajar Sets)

**Module Ownership**: Learning Planning Module

**Description**: Represents one AI generation session for Modul Ajar, containing multiple Modul Ajar Items. A Modul Ajar Set is the first-class domain entity for Modul Ajar generation workflow.

### Columns

| Column | Data Type | Constraints | Description |
|--------|-----------|-------------|-------------|
| `id` | UUID | PRIMARY KEY | Unique Modul Ajar Set identifier |
| `atp_set_id` | UUID | FOREIGN KEY → atp_sets(id), NOT NULL | ATP Set reference |
| `version_no` | INTEGER | NOT NULL, DEFAULT 1 | Version number for this ATP Set |
| `status` | VARCHAR(20) | NOT NULL | Status (DRAFT, UNDER_REVIEW, APPROVED, REJECTED, ARCHIVED) |
| `generation_source` | VARCHAR(50) | NOT NULL | Generation source (AI_GENERATED, MANUAL) |
| `generation_reason` | TEXT | NULLABLE | Reason for generation |
| `generated_by` | UUID | FOREIGN KEY → users(id), NOT NULL | User who requested generation |
| `ai_generation_id` | UUID | NULLABLE | Reference to ai_generation_logs (if AI-generated) |
| `approved_by` | UUID | FOREIGN KEY → users(id) | User who approved Modul Ajar Set |
| `approved_at` | TIMESTAMP WITH TIME ZONE | NULLABLE | Approval timestamp |
| `created_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Modul Ajar Set creation timestamp |
| `updated_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Last update timestamp |

### Indexes

```sql
CREATE INDEX idx_modul_ajar_sets_atp_set_id ON modul_ajar_sets(atp_set_id);
CREATE INDEX idx_modul_ajar_sets_version_no ON modul_ajar_sets(version_no);
CREATE INDEX idx_modul_ajar_sets_status ON modul_ajar_sets(status);
CREATE INDEX idx_modul_ajar_sets_generated_by ON modul_ajar_sets(generated_by);
CREATE INDEX idx_modul_ajar_sets_approved_by ON modul_ajar_sets(approved_by);
CREATE INDEX idx_modul_ajar_sets_ai_generation_id ON modul_ajar_sets(ai_generation_id);
CREATE UNIQUE INDEX idx_modul_ajar_sets_atp_version ON modul_ajar_sets(atp_set_id, version_no);
```

---

## Table: modul_ajar (Modul Ajar Items)

**Module Ownership**: Learning Planning Module

**Description**: Stores individual lesson plan items within a Modul Ajar Set.

### Columns

| Column | Data Type | Constraints | Description |
|--------|-----------|-------------|-------------|
| `id` | UUID | PRIMARY KEY | Unique Modul Ajar identifier |
| `modul_ajar_set_id` | UUID | FOREIGN KEY → modul_ajar_sets(id), NOT NULL | Modul Ajar Set reference |
| `atp_id` | UUID | FOREIGN KEY → atp(id), NOT NULL | ATP reference |
| `week` | INTEGER | NOT NULL | Week number in ATP |
| `topic` | JSONB | NOT NULL | Topic object |
| `resources` | JSONB | NULLABLE | Resources object |
| `class_characteristics` | JSONB | NULLABLE | Class characteristics object |
| `learning_activities` | JSONB | NOT NULL | Learning activities array |
| `resource_requirements` | JSONB | NULLABLE | Resource requirements array |
| `assessment_methods` | JSONB | NULLABLE | Assessment methods array |
| `status` | VARCHAR(20) | NOT NULL | Status (DRAFT, UNDER_REVIEW, APPROVED, REJECTED, ARCHIVED) |
| `created_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Modul Ajar creation timestamp |
| `updated_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Last update timestamp |

### Indexes

```sql
CREATE INDEX idx_modul_ajar_modul_ajar_set_id ON modul_ajar(modul_ajar_set_id);
CREATE INDEX idx_modul_ajar_atp_id ON modul_ajar(atp_id);
CREATE INDEX idx_modul_ajar_week ON modul_ajar(week);
CREATE INDEX idx_modul_ajar_status ON modul_ajar(status);
CREATE INDEX idx_modul_ajar_created_at ON modul_ajar(created_at);
```

### Constraints

```sql
ALTER TABLE modul_ajar ADD CONSTRAINT chk_modul_ajar_week CHECK (week >= 1);
ALTER TABLE modul_ajar ADD CONSTRAINT chk_modul_ajar_status CHECK (status IN ('DRAFT', 'UNDER_REVIEW', 'APPROVED', 'REJECTED', 'ARCHIVED'));
```

---

# SECTION 5 — Assessment Module Tables

## Table: assessments

**Module Ownership**: Assessment Module

**Description**: Stores assessments generated from Modul Ajar.

### Columns

| Column | Data Type | Constraints | Description |
|--------|-----------|-------------|-------------|
| `id` | UUID | PRIMARY KEY | Unique assessment identifier |
| `modul_ajar_id` | UUID | FOREIGN KEY → modul_ajar(id), NOT NULL | Modul Ajar reference |
| `user_id` | UUID | FOREIGN KEY → users(id), NOT NULL | Teacher who created assessment |
| `assessment_type` | VARCHAR(20) | NOT NULL | Assessment type (FORMATIVE, SUMMATIVE) |
| `status` | VARCHAR(20) | NOT NULL | Status (DRAFT, APPROVED, REJECTED) |
| `assessment_items` | JSONB | NOT NULL | Assessment items array |
| `answer_key` | JSONB | NOT NULL | Answer key object |
| `scoring_guidelines` | JSONB | NOT NULL | Scoring guidelines object |
| `ai_confidence_score` | DECIMAL(3,2) | NULLABLE | AI confidence score (0.00-1.00) |
| `ai_generated_at` | TIMESTAMP WITH TIME ZONE | NULLABLE | AI generation timestamp |
| `ai_agent_version` | VARCHAR(20) | NULLABLE | AI agent version |
| `version_no` | INTEGER | NOT NULL, DEFAULT 1 | Version number |
| `is_current_version` | BOOLEAN | NOT NULL, DEFAULT true | Current version flag |
| `parent_version_id` | UUID | FOREIGN KEY → assessments(id), NULLABLE | Parent version reference |
| `created_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Assessment creation timestamp |
| `updated_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Last update timestamp |
| `approved_at` | TIMESTAMP WITH TIME ZONE | NULLABLE | Approval timestamp |
| `approved_by` | UUID | FOREIGN KEY → users(id) | User who approved assessment |

### Indexes

```sql
CREATE INDEX idx_assessments_modul_ajar_id ON assessments(modul_ajar_id);
CREATE INDEX idx_assessments_user_id ON assessments(user_id);
CREATE INDEX idx_assessments_assessment_type ON assessments(assessment_type);
CREATE INDEX idx_assessments_status ON assessments(status);
CREATE INDEX idx_assessments_created_at ON assessments(created_at);
CREATE INDEX idx_assessments_approved_at ON assessments(approved_at);
```

### Constraints

```sql
ALTER TABLE assessments ADD CONSTRAINT chk_assessments_assessment_type CHECK (assessment_type IN ('FORMATIVE', 'SUMMATIVE'));
ALTER TABLE assessments ADD CONSTRAINT chk_assessments_status CHECK (status IN ('DRAFT', 'APPROVED', 'REJECTED'));
ALTER TABLE assessments ADD CONSTRAINT chk_assessments_ai_confidence_score CHECK (ai_confidence_score >= 0.00 AND ai_confidence_score <= 1.00);
```

---

## Table: rubrics

**Module Ownership**: Assessment Module

**Description**: Stores rubrics generated from assessments.

### Columns

| Column | Data Type | Constraints | Description |
|--------|-----------|-------------|-------------|
| `id` | UUID | PRIMARY KEY | Unique rubric identifier |
| `assessment_id` | UUID | FOREIGN KEY → assessments(id), NOT NULL | Assessment reference |
| `user_id` | UUID | FOREIGN KEY → users(id), NOT NULL | Teacher who created rubric |
| `rubric_type` | VARCHAR(20) | NOT NULL | Rubric type (ANALYTIC, HOLISTIC) |
| `status` | VARCHAR(20) | NOT NULL | Status (DRAFT, APPROVED, REJECTED) |
| `performance_criteria` | JSONB | NOT NULL | Performance criteria array |
| `performance_levels` | JSONB | NOT NULL | Performance levels array |
| `scoring_guidelines` | JSONB | NOT NULL | Scoring guidelines object |
| `ai_confidence_score` | DECIMAL(3,2) | NULLABLE | AI confidence score (0.00-1.00) |
| `ai_generated_at` | TIMESTAMP WITH TIME ZONE | NULLABLE | AI generation timestamp |
| `ai_agent_version` | VARCHAR(20) | NULLABLE | AI agent version |
| `version_no` | INTEGER | NOT NULL, DEFAULT 1 | Version number |
| `is_current_version` | BOOLEAN | NOT NULL, DEFAULT true | Current version flag |
| `parent_version_id` | UUID | FOREIGN KEY → rubrics(id), NULLABLE | Parent version reference |
| `created_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Rubric creation timestamp |
| `updated_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Last update timestamp |
| `approved_at` | TIMESTAMP WITH TIME ZONE | NULLABLE | Approval timestamp |
| `approved_by` | UUID | FOREIGN KEY → users(id) | User who approved rubric |

### Indexes

```sql
CREATE INDEX idx_rubrics_assessment_id ON rubrics(assessment_id);
CREATE INDEX idx_rubrics_user_id ON rubrics(user_id);
CREATE INDEX idx_rubrics_rubric_type ON rubrics(rubric_type);
CREATE INDEX idx_rubrics_status ON rubrics(status);
CREATE INDEX idx_rubrics_created_at ON rubrics(created_at);
CREATE INDEX idx_rubrics_approved_at ON rubrics(approved_at);
```

### Constraints

```sql
ALTER TABLE rubrics ADD CONSTRAINT chk_rubrics_rubric_type CHECK (rubric_type IN ('ANALYTIC', 'HOLISTIC'));
ALTER TABLE rubrics ADD CONSTRAINT chk_rubrics_status CHECK (status IN ('DRAFT', 'APPROVED', 'REJECTED'));
ALTER TABLE rubrics ADD CONSTRAINT chk_rubrics_ai_confidence_score CHECK (ai_confidence_score >= 0.00 AND ai_confidence_score <= 1.00);
```

---

## Table: evidences

**Module Ownership**: Assessment Module

**Description**: Stores evidence of student learning.

### Columns

| Column | Data Type | Constraints | Description |
|--------|-----------|-------------|-------------|
| `id` | UUID | PRIMARY KEY | Unique evidence identifier |
| `student_id` | VARCHAR(50) | NOT NULL | Student identifier |
| `assessment_id` | UUID | FOREIGN KEY → assessments(id), NOT NULL | Assessment reference |
| `user_id` | UUID | FOREIGN KEY → users(id), NOT NULL | Teacher who recorded evidence |
| `evidence_type` | VARCHAR(50) | NOT NULL | Evidence type (STUDENT_WORK, ASSESSMENT_RESULT, OBSERVATION) |
| `status` | VARCHAR(20) | NOT NULL | Status (COLLECTED, LINKED, EVALUATED) |
| `evidence_data` | JSONB | NOT NULL | Evidence data object |
| `teacher_notes` | TEXT | NULLABLE | Teacher notes |
| `rubric_id` | UUID | FOREIGN KEY → rubrics(id) | NULLABLE | Linked rubric reference |
| `linked_criteria` | JSONB | NULLABLE | Linked rubric criteria array |
| `evaluation_notes` | TEXT | NULLABLE | Evaluation notes |
| `created_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Evidence creation timestamp |
| `updated_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Last update timestamp |

### Indexes

```sql
CREATE INDEX idx_evidences_student_id ON evidences(student_id);
CREATE INDEX idx_evidences_assessment_id ON evidences(assessment_id);
CREATE INDEX idx_evidences_user_id ON evidences(user_id);
CREATE INDEX idx_evidences_evidence_type ON evidences(evidence_type);
CREATE INDEX idx_evidences_status ON evidences(status);
CREATE INDEX idx_evidences_rubric_id ON evidences(rubric_id);
CREATE INDEX idx_evidences_created_at ON evidences(created_at);
```

### Constraints

```sql
ALTER TABLE evidences ADD CONSTRAINT chk_evidences_evidence_type CHECK (evidence_type IN ('STUDENT_WORK', 'ASSESSMENT_RESULT', 'OBSERVATION'));
ALTER TABLE evidences ADD CONSTRAINT chk_evidences_status CHECK (status IN ('COLLECTED', 'LINKED', 'EVALUATED'));
```

---

## Table: evaluations

**Module Ownership**: Assessment Module

**Description**: Stores student performance evaluations based on rubrics.

### Columns

| Column | Data Type | Constraints | Description |
|--------|-----------|-------------|-------------|
| `id` | UUID | PRIMARY KEY | Unique evaluation identifier |
| `student_id` | VARCHAR(50) | NOT NULL | Student identifier |
| `rubric_id` | UUID | FOREIGN KEY → rubrics(id), NOT NULL | Rubric reference |
| `evidence_id` | UUID | FOREIGN KEY → evidences(id), NOT NULL | Evidence reference |
| `user_id` | UUID | FOREIGN KEY → users(id), NOT NULL | Teacher who performed evaluation |
| `performance_scores` | JSONB | NOT NULL | Performance scores array |
| `total_score` | INTEGER | NOT NULL | Total score |
| `max_score` | INTEGER | NOT NULL | Maximum possible score |
| `performance_level` | VARCHAR(20) | NOT NULL | Performance level (EXCELLENT, PROFICIENT, DEVELOPING, BEGINNING) |
| `evaluated_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Evaluation timestamp |

### Indexes

```sql
CREATE INDEX idx_evaluations_student_id ON evaluations(student_id);
CREATE INDEX idx_evaluations_rubric_id ON evaluations(rubric_id);
CREATE INDEX idx_evaluations_evidence_id ON evaluations(evidence_id);
CREATE INDEX idx_evaluations_user_id ON evaluations(user_id);
CREATE INDEX idx_evaluations_performance_level ON evaluations(performance_level);
CREATE INDEX idx_evaluations_evaluated_at ON evaluations(evaluated_at);
```

### Constraints

```sql
ALTER TABLE evaluations ADD CONSTRAINT chk_evaluations_total_score CHECK (total_score >= 0);
ALTER TABLE evaluations ADD CONSTRAINT chk_evaluations_max_score CHECK (max_score > 0);
ALTER TABLE evaluations ADD CONSTRAINT chk_evaluations_total_score_max CHECK (total_score <= max_score);
ALTER TABLE evaluations ADD CONSTRAINT chk_evaluations_performance_level CHECK (performance_level IN ('EXCELLENT', 'PROFICIENT', 'DEVELOPING', 'BEGINNING'));
```

---

# SECTION 6 — Reporting Module Tables

## Table: narrative_reports

**Module Ownership**: Reporting Module

**Description**: Stores narrative reports generated from evidence and evaluations.

### Columns

| Column | Data Type | Constraints | Description |
|--------|-----------|-------------|-------------|
| `id` | UUID | PRIMARY KEY | Unique narrative report identifier |
| `student_id` | VARCHAR(50) | NOT NULL | Student identifier |
| `user_id` | UUID | FOREIGN KEY → users(id), NOT NULL | Teacher who created report |
| `status` | VARCHAR(20) | NOT NULL | Status (DRAFT, APPROVED, REJECTED) |
| `report_period` | JSONB | NOT NULL | Report period object |
| `language` | VARCHAR(20) | NOT NULL | Report language (INDONESIAN, ENGLISH) |
| `content` | JSONB | NOT NULL | Report content object |
| `ai_confidence_score` | DECIMAL(3,2) | NULLABLE | AI confidence score (0.00-1.00) |
| `ai_generated_at` | TIMESTAMP WITH TIME ZONE | NULLABLE | AI generation timestamp |
| `ai_agent_version` | VARCHAR(20) | NULLABLE | AI agent version |
| `version_no` | INTEGER | NOT NULL, DEFAULT 1 | Version number |
| `is_current_version` | BOOLEAN | NOT NULL, DEFAULT true | Current version flag |
| `parent_version_id` | UUID | FOREIGN KEY → narrative_reports(id), NULLABLE | Parent version reference |
| `created_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Report creation timestamp |
| `updated_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Last update timestamp |
| `approved_at` | TIMESTAMP WITH TIME ZONE | NULLABLE | Approval timestamp |
| `approved_by` | UUID | FOREIGN KEY → users(id) | User who approved report |

### Indexes

```sql
CREATE INDEX idx_narrative_reports_student_id ON narrative_reports(student_id);
CREATE INDEX idx_narrative_reports_user_id ON narrative_reports(user_id);
CREATE INDEX idx_narrative_reports_status ON narrative_reports(status);
CREATE INDEX idx_narrative_reports_language ON narrative_reports(language);
CREATE INDEX idx_narrative_reports_created_at ON narrative_reports(created_at);
CREATE INDEX idx_narrative_reports_approved_at ON narrative_reports(approved_at);
```

### Constraints

```sql
ALTER TABLE narrative_reports ADD CONSTRAINT chk_narrative_reports_status CHECK (status IN ('DRAFT', 'APPROVED', 'REJECTED'));
ALTER TABLE narrative_reports ADD CONSTRAINT chk_narrative_reports_language CHECK (language IN ('INDONESIAN', 'ENGLISH'));
ALTER TABLE narrative_reports ADD CONSTRAINT chk_narrative_reports_ai_confidence_score CHECK (ai_confidence_score >= 0.00 AND ai_confidence_score <= 1.00);
```

---

# SECTION 7 — Administration Module Tables

## Table: audit_logs

**Module Ownership**: Administration Module

**Description**: Stores audit trail for all system changes.

### Columns

| Column | Data Type | Constraints | Description |
|--------|-----------|-------------|-------------|
| `id` | UUID | PRIMARY KEY | Unique audit log identifier |
| `user_id` | UUID | FOREIGN KEY → users(id) | User who performed action |
| `action` | VARCHAR(50) | NOT NULL | Action performed (CREATE, UPDATE, DELETE, APPROVE) |
| `entity_type` | VARCHAR(50) | NOT NULL | Entity type (TP, ATP, MODUL_AJAR, etc.) |
| `entity_id` | UUID | NOT NULL | Entity identifier |
| `changes` | JSONB | NULLABLE | Changes made (old and new values) |
| `ip_address` | INET | NULLABLE | IP address of user |
| `user_agent` | TEXT | NULLABLE | User agent string |
| `created_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Audit log creation timestamp |

### Indexes

```sql
CREATE INDEX idx_audit_logs_user_id ON audit_logs(user_id);
CREATE INDEX idx_audit_logs_action ON audit_logs(action);
CREATE INDEX idx_audit_logs_entity_type ON audit_logs(entity_type);
CREATE INDEX idx_audit_logs_entity_id ON audit_logs(entity_id);
CREATE INDEX idx_audit_logs_created_at ON audit_logs(created_at);
CREATE INDEX idx_audit_logs_user_id_created_at ON audit_logs(user_id, created_at);
```

---

# SECTION 8 — Entity Relationship Diagram (ERD)

## ERD Description

### Core Relationships

```
users (1) ──────< (N) tp
users (1) ──────< (N) atp
users (1) ──────< (N) modul_ajar
users (1) ──────< (N) assessments
users (1) ──────< (N) rubrics
users (1) ──────< (N) evidences
users (1) ──────< (N) evaluations
users (1) ──────< (N) narrative_reports
users (1) ──────< (N) audit_logs

roles (1) ──────< (N) users
roles (1) ──────< (N) permissions

users (1) ──────< (N) refresh_tokens

cp (1) ──────< (N) tp
tp (1) ──────< (N) atp
atp (1) ──────< (N) modul_ajar
modul_ajar (1) ──────< (N) assessments
assessments (1) ──────< (N) rubrics
assessments (1) ──────< (N) evidences
rubrics (1) ──────< (N) evidences
evidences (1) ──────< (N) evaluations
```

### Workflow Relationships

```
CP → TP → ATP → Modul Ajar → Assessment → Rubric → Evidence → Evaluation → Narrative Report
```

### Cardinality

- **One-to-Many**: users to all artifact tables
- **One-to-Many**: cp to tp
- **One-to-Many**: tp to atp
- **One-to-Many**: atp to modul_ajar
- **One-to-Many**: modul_ajar to assessments
- **One-to-Many**: assessments to rubrics
- **One-to-Many**: assessments to evidences
- **One-to-Many**: rubrics to evidences
- **One-to-Many**: evidences to evaluations

---

# SECTION 9 — Schema Ownership by Module

## Authentication Module

**Tables**:
- `users`
- `roles`
- `permissions`
- `refresh_tokens`

**Ownership**: Authentication Module Owner

**Responsibilities**:
- User account management
- Role and permission management
- Token management
- Authentication security

---

## Curriculum Module

**Tables**:
- `cp`
- `tp`

**Ownership**: Curriculum Domain Owner

**Responsibilities**:
- National curriculum data management
- Teaching plan generation and management
- Curriculum-to-teaching plan workflow

---

## Learning Planning Module

**Tables**:
- `atp`
- `modul_ajar`

**Ownership**: Learning Planning Domain Owner

**Responsibilities**:
- Annual teaching plan generation and management
- Modul Ajar generation and management
- Teaching schedule management

---

## Assessment Module

**Tables**:
- `assessments`
- `rubrics`
- `evidences`
- `evaluations`

**Ownership**: Assessment Domain Owner

**Responsibilities**:
- Assessment generation and management
- Rubric generation and management
- Evidence collection and management
- Student performance evaluation

---

## Reporting Module

**Tables**:
- `narrative_reports`

**Ownership**: Reporting Domain Owner

**Responsibilities**:
- Narrative report generation and management
- Parent communication
- Student progress reporting

---

## Administration Module

**Tables**:
- `audit_logs`

**Ownership**: Platform Owner

**Responsibilities**:
- Audit trail management
- System change tracking
- Compliance and governance

---

# SECTION 10 — Migration Strategy

## Migration Tool

**Tool**: golang-migrate

**Migration Path**: `backend/migrations/`

## Migration Naming Convention

```
NNNNNN_description.up.sql
NNNNNN_description.down.sql
```

Example:
```
000001_init_schema.up.sql
000001_init_schema.down.sql
000002_users.up.sql
000002_users.down.sql
```

## Migration Order

1. **000001_init_schema**: Create extension functions, enums, types
2. **000002_users**: Create users, roles, permissions, refresh_tokens tables
3. **000003_curriculum**: Create cp, tp tables
4. **000004_learning_planning**: Create atp, modul_ajar tables
5. **000005_assessment**: Create assessments, rubrics, evidences, evaluations tables
6. **000006_reporting**: Create narrative_reports table
7. **000007_administration**: Create audit_logs table
8. **000008_seed_data**: Insert seed data (roles, permissions)

---

# SECTION 11 — Sample Migration SQL

## Migration 000001_init_schema.up.sql

```sql
-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create enums
CREATE TYPE user_role AS ENUM ('TEACHER', 'ADMINISTRATOR');
CREATE TYPE artifact_status AS ENUM ('DRAFT', 'APPROVED', 'REJECTED');
CREATE TYPE assessment_type AS ENUM ('FORMATIVE', 'SUMMATIVE');
CREATE TYPE rubric_type AS ENUM ('ANALYTIC', 'HOLISTIC');
CREATE TYPE evidence_type AS ENUM ('STUDENT_WORK', 'ASSESSMENT_RESULT', 'OBSERVATION');
CREATE TYPE evidence_status AS ENUM ('COLLECTED', 'LINKED', 'EVALUATED');
CREATE TYPE performance_level AS ENUM ('EXCELLENT', 'PROFICIENT', 'DEVELOPING', 'BEGINNING');
CREATE TYPE report_language AS ENUM ('INDONESIAN', 'ENGLISH');
CREATE TYPE audit_action AS ENUM ('CREATE', 'UPDATE', 'DELETE', 'APPROVE');

-- Create updated_at trigger function
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';
```

## Migration 000002_users.up.sql

```sql
-- Create roles table
CREATE TABLE roles (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    name VARCHAR(50) UNIQUE NOT NULL,
    description TEXT,
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- Create users table
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    role_id UUID NOT NULL REFERENCES roles(id),
    is_active BOOLEAN NOT NULL DEFAULT true,
    failed_login_attempts INTEGER NOT NULL DEFAULT 0,
    locked_until TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    created_by UUID REFERENCES users(id),
    updated_by UUID REFERENCES users(id),
    CONSTRAINT chk_users_email_format CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'),
    CONSTRAINT chk_users_failed_login_attempts CHECK (failed_login_attempts >= 0)
);

-- Create permissions table
CREATE TABLE permissions (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    role_id UUID NOT NULL REFERENCES roles(id),
    resource VARCHAR(100) NOT NULL,
    action VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    UNIQUE (role_id, resource, action)
);

-- Create refresh_tokens table
CREATE TABLE refresh_tokens (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    user_id UUID NOT NULL REFERENCES users(id),
    token TEXT UNIQUE NOT NULL,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    revoked_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    ip_address INET,
    CONSTRAINT chk_refresh_tokens_expires_at CHECK (expires_at > created_at)
);

-- Create indexes
CREATE INDEX idx_roles_name ON roles(name);
CREATE INDEX idx_roles_is_active ON roles(is_active);
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_role_id ON users(role_id);
CREATE INDEX idx_users_is_active ON users(is_active);
CREATE INDEX idx_users_created_at ON users(created_at);
CREATE INDEX idx_permissions_role_id ON permissions(role_id);
CREATE INDEX idx_permissions_resource ON permissions(resource);
CREATE INDEX idx_refresh_tokens_user_id ON refresh_tokens(user_id);
CREATE INDEX idx_refresh_tokens_token ON refresh_tokens(token);
CREATE INDEX idx_refresh_tokens_expires_at ON refresh_tokens(expires_at);

-- Create triggers
CREATE TRIGGER update_roles_updated_at BEFORE UPDATE ON roles
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON users
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
```

---

# SECTION 12 — Workflow History

## Overview

The Workflow History table provides complete workflow traceability for all lifecycle transitions across MVP entities. This table ensures every status change is auditable and traceable.

## Module Ownership

**Module**: Cross-Module (used by Curriculum, Learning Planning, Assessment, Reporting)

**Purpose**: Workflow traceability and audit trail

---

## Table: workflow_history

**Module Ownership**: Cross-Module

**Description**: Logs all workflow status transitions for entities requiring lifecycle management.

### Columns

| Column | Data Type | Constraints | Description |
|--------|-----------|-------------|-------------|
| `id` | UUID | PRIMARY KEY | Unique workflow history identifier |
| `entity_type` | VARCHAR(100) | NOT NULL | Entity type (TP, ATP, ModulAjar, Assessment, Rubric, NarrativeReport) |
| `entity_id` | UUID | NOT NULL | Reference to entity ID |
| `from_status` | VARCHAR(50) | NOT NULL | Previous status |
| `to_status` | VARCHAR(50) | NOT NULL | New status |
| `changed_by` | UUID | FOREIGN KEY → users(id), NOT NULL | User who made the change |
| `changed_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Change timestamp |
| `notes` | TEXT | Optional notes about the change | |

### Indexes

```sql
CREATE INDEX idx_workflow_history_entity_type ON workflow_history(entity_type);
CREATE INDEX idx_workflow_history_entity_id ON workflow_history(entity_id);
CREATE INDEX idx_workflow_history_changed_at ON workflow_history(changed_at);
CREATE INDEX idx_workflow_history_changed_by ON workflow_history(changed_by);
```

### Constraints

```sql
ALTER TABLE workflow_history ADD CONSTRAINT chk_workflow_history_entity_type 
    CHECK (entity_type IN ('TP', 'ATP', 'ModulAjar', 'Assessment', 'Rubric', 'NarrativeReport'));
```

---

## Migration SQL

```sql
-- Enable UUID extension if not already enabled
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create workflow_history table
CREATE TABLE workflow_history (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    entity_type VARCHAR(100) NOT NULL,
    entity_id UUID NOT NULL,
    from_status VARCHAR(50) NOT NULL,
    to_status VARCHAR(50) NOT NULL,
    changed_by UUID NOT NULL REFERENCES users(id),
    changed_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    notes TEXT,
    CONSTRAINT chk_workflow_history_entity_type 
        CHECK (entity_type IN ('TP', 'ATP', 'ModulAjar', 'Assessment', 'Rubric', 'NarrativeReport'))
);

-- Create indexes
CREATE INDEX idx_workflow_history_entity_type ON workflow_history(entity_type);
CREATE INDEX idx_workflow_history_entity_id ON workflow_history(entity_id);
CREATE INDEX idx_workflow_history_changed_at ON workflow_history(changed_at);
CREATE INDEX idx_workflow_history_changed_by ON workflow_history(changed_by);
```

---

## Usage Requirements

### Entities Using Workflow History

The following entities must generate workflow_history records on every lifecycle transition:

- **TP** (Teaching Plan): Curriculum Module
- **ATP** (Annual Teaching Plan): Learning Planning Module
- **Modul Ajar** (Lesson Plan): Learning Planning Module
- **Assessment**: Assessment Module
- **Rubric**: Assessment Module
- **Narrative Report**: Reporting Module

### Lifecycle Transitions

Every lifecycle transition must generate a workflow_history record:

- DRAFT → GENERATED
- GENERATED → REVIEWED
- REVIEWED → APPROVED
- APPROVED → PUBLISHED
- PUBLISHED → ARCHIVED

### Audit Requirements

- Every status change must be logged
- User who made the change must be recorded
- Timestamp must be accurate
- Optional notes may be provided for context

---

## Conclusion

The Workflow History table provides complete workflow traceability for all lifecycle transitions, ensuring auditability and governance across all MVP entities requiring lifecycle management.

---

# SECTION 13 — Artifact Versioning Strategy

## Overview

The Artifact Versioning Strategy provides support for versioning of key educational artifacts, enabling AI regeneration, teacher edits, auditability, and rollback capabilities.

## Module Ownership

**Module**: Cross-Module (Curriculum, Learning Planning, Assessment, Reporting)

**Purpose**: Artifact versioning and rollback support

---

## Entities Requiring Versioning

The following entities must support versioning:

- **TP** (Teaching Plan): Curriculum Module
- **ATP** (Annual Teaching Plan): Learning Planning Module
- **Modul Ajar** (Lesson Plan): Learning Planning Module
- **Assessment**: Assessment Module
- **Rubric**: Assessment Module
- **Narrative Report**: Reporting Module

---

## Standard Versioning Fields

Each versioned entity must include the following standard fields:

| Field | Data Type | Constraints | Description |
|-------|-----------|-------------|-------------|
| `version_no` | INTEGER | NOT NULL, DEFAULT 1 | Version number (starts from 1) |
| `is_current_version` | BOOLEAN | NOT NULL, DEFAULT true | Whether this is the current version |
| `parent_version_id` | UUID | FOREIGN KEY → self(id) | Reference to parent version |

---

## Versioning Rules

### Version Number

- `version_no` starts from 1 for the first version
- Every approved modification creates a new version
- Version numbers increment sequentially (1, 2, 3, ...)
- Version numbers are immutable once assigned

### Current Version

- Only one version may have `is_current_version = TRUE` per entity
- When a new version is created, the previous version's `is_current_version` is set to FALSE
- The current version is the version used in all operations
- Historical versions remain immutable

### Parent Version

- `parent_version_id` references the previous version
- The first version has `parent_version_id = NULL`
- Parent version reference enables version tree traversal
- Parent version reference supports rollback capability

### Version Creation Triggers

A new version is created when:
- AI regeneration occurs (new AI-generated content)
- Teacher edits are made and approved
- Manual version creation is initiated
- Rollback to a previous version is performed

---

## Version Query Patterns

### Get Current Version

```sql
SELECT * FROM {table} 
WHERE entity_id = ? AND is_current_version = TRUE;
```

### Get Version History

```sql
SELECT * FROM {table} 
WHERE entity_id = ? 
ORDER BY version_no DESC;
```

### Get Specific Version

```sql
SELECT * FROM {table} 
WHERE entity_id = ? AND version_no = ?;
```

### Rollback to Previous Version

```sql
-- Set current version to FALSE
UPDATE {table} 
SET is_current_version = FALSE 
WHERE entity_id = ? AND is_current_version = TRUE;

-- Create new version from parent
INSERT INTO {table} (...)
SELECT ..., parent_version_id = current_version_id
FROM {table}
WHERE id = current_version_id;

-- Set new version as current
UPDATE {table} 
SET is_current_version = TRUE 
WHERE id = new_version_id;
```

---

## Version Constraints

### CHECK Constraints

```sql
ALTER TABLE {table} ADD CONSTRAINT chk_{table}_version_no_positive 
    CHECK (version_no >= 1);

ALTER TABLE {table} ADD CONSTRAINT chk_{table}_single_current_version 
    CHECK (
        entity_id IN (
            SELECT entity_id FROM {table} 
            WHERE is_current_version = TRUE 
            GROUP BY entity_id 
            HAVING COUNT(*) = 1
        )
    );
```

### TRIGGER for Current Version Enforcement

```sql
CREATE OR REPLACE FUNCTION enforce_single_current_version()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.is_current_version = TRUE THEN
        UPDATE {table}
        SET is_current_version = FALSE
        WHERE entity_id = NEW.entity_id 
        AND id != NEW.id 
        AND is_current_version = TRUE;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_enforce_single_current_version
    BEFORE INSERT OR UPDATE ON {table}
    FOR EACH ROW EXECUTE FUNCTION enforce_single_current_version();
```

---

## Version Audit Trail

All version changes must be recorded in `workflow_history`:

- Version creation: DRAFT → GENERATED (new version)
- Version approval: REVIEWED → APPROVED (new version)
- Version rollback: APPROVED → APPROVED (rollback to previous version)

---

## Conclusion

The Artifact Versioning Strategy provides comprehensive versioning support for key educational artifacts, enabling AI regeneration, teacher edits, auditability, and rollback capabilities while maintaining data integrity and traceability.

---

# SECTION 14 — Lifecycle Status Standard

## Overview

The Lifecycle Status Standard defines the unified workflow lifecycle for all key educational artifacts, ensuring consistent status management across all modules.

## Module Ownership

**Module**: Cross-Module (Curriculum, Learning Planning, Assessment, Reporting)

**Purpose**: Unified lifecycle status management

---

## Entities Requiring Lifecycle Management

The following entities must implement lifecycle management:

- **TP** (Teaching Plan): Curriculum Module
- **ATP** (Annual Teaching Plan): Learning Planning Module
- **Modul Ajar** (Lesson Plan): Learning Planning Module
- **Assessment**: Assessment Module
- **Rubric**: Assessment Module
- **Narrative Report**: Reporting Module

---

## Standard Status Field

Each lifecycle-managed entity must include the following standard field:

| Field | Data Type | Constraints | Description |
|-------|-----------|-------------|-------------|
| `status` | VARCHAR(50) | NOT NULL, DEFAULT 'DRAFT' | Current lifecycle status |

---

## Standard Status Values

The following status values are defined for all lifecycle-managed entities:

| Status | Description |
|--------|-------------|
| `DRAFT` | Initial state, artifact is being created or edited |
| `GENERATED` | Artifact has been generated by AI or system |
| `REVIEWED` | Artifact has been reviewed by a human |
| `APPROVED` | Artifact has been approved for use |
| `PUBLISHED` | Artifact is published and available for use |
| `ARCHIVED` | Artifact is archived and no longer active |

---

## Lifecycle Transitions

### Valid Status Transitions

The following transitions are valid for all lifecycle-managed entities:

```
DRAFT → GENERATED → REVIEWED → APPROVED → PUBLISHED → ARCHIVED
```

### Transition Rules

- **DRAFT → GENERATED**: Artifact generation (AI or manual)
- **GENERATED → REVIEWED**: Human review initiated
- **REVIEWED → APPROVED**: Approval granted
- **APPROVED → PUBLISHED**: Publication for use
- **PUBLISHED → ARCHIVED**: Artifact archival
- **REVIEWED → DRAFT**: Rejection requiring revision
- **APPROVED → REVIEWED**: Re-approval required after changes

### Transition Requirements

Every status transition must:
1. Be recorded in `workflow_history` table
2. Include the user who made the change
3. Include optional notes for context
4. Follow the valid transition rules

---

## Status Constraints

### CHECK Constraint

```sql
ALTER TABLE {table} ADD CONSTRAINT chk_{table}_status 
    CHECK (status IN ('DRAFT', 'GENERATED', 'REVIEWED', 'APPROVED', 'PUBLISHED', 'ARCHIVED'));
```

### TRIGGER for Workflow History Recording

```sql
CREATE OR REPLACE FUNCTION record_workflow_history()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.status IS DISTINCT FROM NEW.status THEN
        INSERT INTO workflow_history (
            entity_type,
            entity_id,
            from_status,
            to_status,
            changed_by,
            changed_at,
            notes
        ) VALUES (
            TG_TABLE_NAME,
            NEW.id,
            OLD.status,
            NEW.status,
            NEW.updated_by,
            NOW(),
            NULL
        );
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_record_workflow_history
    AFTER UPDATE ON {table}
    FOR EACH ROW EXECUTE FUNCTION record_workflow_history();
```

---

## Status Query Patterns

### Get Current Status

```sql
SELECT id, status FROM {table} WHERE id = ?;
```

### Get Entities by Status

```sql
SELECT * FROM {table} WHERE status = ?;
```

### Get Workflow History

```sql
SELECT * FROM workflow_history 
WHERE entity_type = ? AND entity_id = ? 
ORDER BY changed_at DESC;
```

---

## Status-Based Access Control

### Status-Based Permissions

- **DRAFT**: Creator and editors can modify
- **GENERATED**: Reviewers can review
- **REVIEWED**: Approvers can approve
- **APPROVED**: Publishers can publish
- **PUBLISHED**: Read-only for most users
- **ARCHIVED**: Read-only for audit purposes

### Status-Based Notifications

- **GENERATED**: Notify reviewers
- **REVIEWED**: Notify approvers
- **APPROVED**: Notify publishers
- **PUBLISHED**: Notify stakeholders
- **ARCHIVED**: Notify for record-keeping

---

## Conclusion

The Lifecycle Status Standard provides a unified workflow lifecycle model for all key educational artifacts, ensuring consistent status management, auditability, and governance across all modules. This lifecycle model becomes the official workflow standard for NUSA MVP Wave 1.

---

# SECTION 15 — Audit Field Standard

## Overview

The Audit Field Standard defines the required audit fields for all tables based on their classification, ensuring comprehensive auditability and traceability across the database.

## Module Ownership

**Module**: Cross-Module (all modules)

**Purpose**: Audit field standardization and compliance

---

## Table Classification

### Tier A — Transaction Tables

**Definition**: Tables that capture business transactions and operational activities with high audit requirements.

**Examples**:
- TP (Teaching Plan)
- ATP (Annual Teaching Plan)
- Modul Ajar (Lesson Plan)
- Assessment
- Rubric
- Evidence
- Narrative Report

**Required Fields**:

| Field | Data Type | Constraints | Description |
|-------|-----------|-------------|-------------|
| `created_at` | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() | Record creation timestamp |
| `created_by` | UUID | FOREIGN KEY → users(id) | User who created the record |
| `updated_at` | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() | Last update timestamp |
| `updated_by` | UUID | FOREIGN KEY → users(id) | User who last updated the record |
| `deleted_at` | TIMESTAMPTZ | NULLABLE | Soft delete timestamp |
| `deleted_by` | UUID | FOREIGN KEY → users(id), NULLABLE | User who deleted the record |

**Rules**:
- All 6 fields are required
- `created_by`, `updated_by`, `deleted_by` must represent actual actors (no random UUID defaults)
- Soft delete pattern: set `deleted_at` and `deleted_by` instead of hard delete
- `updated_at` must be updated on every modification

### Tier B — Master Tables

**Definition**: Core business entities that change infrequently and are referenced by transactional data.

**Examples**:
- School
- Teacher
- Student
- Subject
- Grade

**Required Fields**:

| Field | Data Type | Constraints | Description |
|-------|-----------|-------------|-------------|
| `created_at` | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() | Record creation timestamp |
| `updated_at` | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() | Last update timestamp |

**Optional Fields**:

| Field | Data Type | Constraints | Description |
|-------|-----------|-------------|-------------|
| `created_by` | UUID | FOREIGN KEY → users(id), NULLABLE | User who created the record |
| `updated_by` | UUID | FOREIGN KEY → users(id), NULLABLE | User who last updated the record |

**Rules**:
- `created_at` and `updated_at` are required
- `created_by` and `updated_by` are optional based on business requirements
- If used, must represent actual actors (no random UUID defaults)

### Tier C — Reference Tables

**Definition**: Static or semi-static data that provides context and classification.

**Examples**:
- Phase
- Religion
- Gender
- Academic Calendar Reference

**Required Fields**:

| Field | Data Type | Constraints | Description |
|-------|-----------|-------------|-------------|
| `created_at` | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() | Record creation timestamp |

**Rules**:
- Only `created_at` is required
- No actor tracking needed for reference data
- Reference data is typically system-managed

---

## Audit Field Implementation

### TRIGGER for Updated Timestamp

```sql
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_update_updated_at
    BEFORE UPDATE ON {table}
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
```

### Soft Delete Pattern

```sql
-- Soft delete
UPDATE {table}
SET deleted_at = NOW(),
    deleted_by = ?
WHERE id = ?;

-- Query excluding deleted records
SELECT * FROM {table} WHERE deleted_at IS NULL;

-- Query including deleted records
SELECT * FROM {table} WHERE deleted_at IS NOT NULL;
```

---

## Audit Field Constraints

### CHECK Constraints

```sql
-- Ensure created_at <= updated_at
ALTER TABLE {table} ADD CONSTRAINT chk_{table}_created_at_before_updated_at 
    CHECK (created_at <= updated_at);

-- Ensure deleted_at >= updated_at (if applicable)
ALTER TABLE {table} ADD CONSTRAINT chk_{table}_deleted_at_after_updated_at 
    CHECK (deleted_at IS NULL OR deleted_at >= updated_at);
```

### Actor Field Validation

**IMPORTANT**: Do NOT assign random UUID defaults to `created_by`, `updated_by`, or `deleted_by`.

These fields must represent actual actors:
- Must reference valid user records in `users` table
- Must be set by application logic based on authenticated user
- Must not use DEFAULT uuid_generate_v4()
- Must be NULLABLE if actor is not applicable

---

## Audit Query Patterns

### Get Creation History

```sql
SELECT id, created_at, created_by FROM {table} WHERE id = ?;
```

### Get Update History

```sql
SELECT id, updated_at, updated_by FROM {table} WHERE id = ?;
```

### Get Deleted Records

```sql
SELECT * FROM {table} WHERE deleted_at IS NOT NULL;
```

### Get Active Records Only

```sql
SELECT * FROM {table} WHERE deleted_at IS NULL;
```

---

## Conclusion

The Audit Field Standard provides comprehensive auditability and traceability across all database tables based on their classification. This standard ensures consistent audit field implementation while respecting the different requirements of transaction, master, and reference tables.

---

# SECTION 17 — Approval Ownership

## Overview

The Approval Ownership standard defines the approval fields required for entities that require human approval before publication, ensuring educational accountability and governance.

## Module Ownership

**Module**: Cross-Module (Curriculum, Learning Planning, Assessment, Reporting)

**Purpose**: Approval ownership and educational accountability

---

## Entities Requiring Approval Ownership

The following entities require approval ownership fields:

- **TP** (Teaching Plan): Curriculum Module
- **ATP** (Annual Teaching Plan): Learning Planning Module
- **Modul Ajar** (Lesson Plan): Learning Planning Module
- **Assessment**: Assessment Module
- **Rubric**: Assessment Module
- **Narrative Report**: Reporting Module

---

## Standard Approval Fields

Each approval-managed entity must include the following standard fields:

| Field | Data Type | Constraints | Description |
|-------|-----------|-------------|-------------|
| `approved_by` | UUID | FOREIGN KEY → users(id), NULLABLE | User who approved the artifact |
| `approved_at` | TIMESTAMPTZ | NULLABLE | Approval timestamp |

---

## Approval Rules

### Approval Transition

Approval can only occur from:
- **REVIEWED → APPROVED**

Approval cannot occur from any other status.

### Authorization Requirements

Only authorized users may approve:
- Users with APPROVER role
- Users with appropriate permissions
- Users who are not the creator (separation of duties)

### Approval Process

1. Artifact reaches REVIEWED status
2. Authorized reviewer reviews artifact
3. If approved, set `approved_by` and `approved_at`
4. Status transitions to APPROVED
5. Workflow history record is generated

### Approval Revocation

Approval can be revoked:
- Set `approved_by` to NULL
- Set `approved_at` to NULL
- Status transitions back to REVIEWED
- Workflow history record is generated

---

## Approval Constraints

### CHECK Constraint

```sql
ALTER TABLE {table} ADD CONSTRAINT chk_{table}_approval_consistency 
    CHECK (
        (approved_by IS NULL AND approved_at IS NULL) OR
        (approved_by IS NOT NULL AND approved_at IS NOT NULL)
    );
```

### TRIGGER for Workflow History Recording

```sql
CREATE OR REPLACE FUNCTION record_approval_workflow_history()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.approved_by IS DISTINCT FROM NEW.approved_by THEN
        IF NEW.approved_by IS NOT NULL THEN
            -- Approval granted
            INSERT INTO workflow_history (
                entity_type,
                entity_id,
                from_status,
                to_status,
                changed_by,
                changed_at,
                notes
            ) VALUES (
                TG_TABLE_NAME,
                NEW.id,
                'REVIEWED',
                'APPROVED',
                NEW.approved_by,
                NEW.approved_at,
                'Approval granted'
            );
        ELSE
            -- Approval revoked
            INSERT INTO workflow_history (
                entity_type,
                entity_id,
                from_status,
                to_status,
                changed_by,
                changed_at,
                notes
            ) VALUES (
                TG_TABLE_NAME,
                NEW.id,
                'APPROVED',
                'REVIEWED',
                NEW.updated_by,
                NOW(),
                'Approval revoked'
            );
        END IF;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_record_approval_workflow_history
    AFTER UPDATE ON {table}
    FOR EACH ROW EXECUTE FUNCTION record_approval_workflow_history();
```

---

## Approval Query Patterns

### Get Approved Artifacts

```sql
SELECT * FROM {table} 
WHERE approved_by IS NOT NULL 
AND status = 'APPROVED';
```

### Get Pending Approvals

```sql
SELECT * FROM {table} 
WHERE approved_by IS NULL 
AND status = 'REVIEWED';
```

### Get Approval History

```sql
SELECT * FROM workflow_history 
WHERE entity_type = ? 
AND entity_id = ? 
AND to_status = 'APPROVED'
ORDER BY changed_at DESC;
```

---

## Approval Audit Requirements

### Audit Trail

All approval actions must be auditable:
- Approval timestamp recorded
- Approver identity recorded
- Approval reason recorded (optional)
- Workflow history generated

### Accountability

Approval ownership ensures:
- Educational accountability for content
- Traceability of approval decisions
- Support for governance reviews
- Compliance with educational standards

---

## Conclusion

The Approval Ownership standard provides comprehensive approval tracking for key educational artifacts, ensuring educational accountability and governance while maintaining auditability and traceability.

---

# SECTION 18 — Database Migration Governance

## Purpose

Ensure all database changes remain controlled, reproducible, and auditable.

---

## Migration Principles

Database schema changes must never be applied manually in production environments.

Every schema modification must be represented by a migration file.

---

## Required Workflow

Schema Change Request

↓

Migration File Creation

↓

Code Review

↓

Migration Validation

↓

Deployment

↓

Migration Execution

---

## Migration Requirements

Every migration must:

- Have a unique migration identifier
- Be version controlled
- Be reviewed before deployment
- Be executable multiple times safely when applicable
- Include rollback guidance

---

## Naming Convention

YYYYMMDDHHMM_<description>

Examples:

202606050900_create_tp_table

202606051200_add_status_column

---

## Rollback Rule

Every migration must document:

- Impact
- Rollback approach
- Recovery procedure

---

## Production Rule

Direct database modification is prohibited.

All production schema changes must originate from approved migration files.

---

## Source of Truth

The migration repository is the official source of truth for database structure.

The live database must always be reproducible from migration history.

---

This migration governance model becomes the official database change management standard for MVP Wave 1.

---

# SECTION 19 — Conclusion

## Overview

The AI Audit and Observability Schema provides comprehensive auditability, monitoring, governance, and debugging capabilities for all AI operations. This schema ensures all AI-generated content, prompt versions, human reviews, and AI operations are fully traceable and auditable.

## Module Ownership

**Module**: AI Orchestration Module

**Purpose**: AI governance, auditability, and observability

---

## Table: ai_generation_log

**Module Ownership**: AI Orchestration Module

**Description**: Logs all AI generation requests, including input/output hashes, provider information, and performance metrics. This table provides complete traceability for all AI-generated content.

### Columns

| Column | Data Type | Constraints | Description |
|--------|-----------|-------------|-------------|
| `id` | UUID | PRIMARY KEY | Unique generation log identifier |
| `agent_name` | VARCHAR(50) | NOT NULL | AI agent name (e.g., TP_GENERATOR, ATP_GENERATOR) |
| `prompt_version` | VARCHAR(20) | NOT NULL | Prompt version used for generation |
| `request_id` | UUID | NOT NULL | Unique request identifier |
| `user_id` | UUID | FOREIGN KEY → users(id) | User who initiated the generation |
| `feature_id` | VARCHAR(20) | NOT NULL | Feature ID from Feature Traceability Matrix |
| `entity_type` | VARCHAR(100) | NOT NULL | Entity type (TP, ATP, ModulAjar, Assessment, Rubric, NarrativeReport) |
| `entity_id` | UUID | NOT NULL | Reference to target entity ID |
| `input_hash` | VARCHAR(64) | NOT NULL | SHA-256 hash of input data |
| `output_hash` | VARCHAR(64) | NOT NULL | SHA-256 hash of output data |
| `generation_status` | VARCHAR(20) | NOT NULL | Generation status (SUCCESS, FAILED, TIMEOUT, VALIDATION_FAILED) |
| `provider_name` | VARCHAR(50) | NOT NULL | LLM provider name (e.g., OpenAI, Anthropic) |
| `model_name` | VARCHAR(100) | NOT NULL | LLM model name (e.g., gpt-4, claude-3-opus) |
| `duration_ms` | INTEGER | NOT NULL | Generation duration in milliseconds |
| `token_input` | INTEGER | NOT NULL | Number of input tokens used |
| `token_output` | INTEGER | NOT NULL | Number of output tokens generated |
| `confidence_score` | DECIMAL(3,2) | CHECK (confidence_score >= 0 AND confidence_score <= 1) | AI confidence score (0.00-1.00) |
| `error_code` | VARCHAR(50) | Error code if generation failed | |
| `error_message` | TEXT | Error message if generation failed | |
| `created_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Generation timestamp |

### Indexes

```sql
CREATE INDEX idx_ai_generation_log_created_at ON ai_generation_log(created_at);
CREATE INDEX idx_ai_generation_log_agent_name ON ai_generation_log(agent_name);
CREATE INDEX idx_ai_generation_log_generation_status ON ai_generation_log(generation_status);
CREATE INDEX idx_ai_generation_log_provider_name ON ai_generation_log(provider_name);
CREATE INDEX idx_ai_generation_log_user_id ON ai_generation_log(user_id);
CREATE INDEX idx_ai_generation_log_request_id ON ai_generation_log(request_id);
CREATE INDEX idx_ai_generation_log_feature_id ON ai_generation_log(feature_id);
CREATE INDEX idx_ai_generation_log_entity_type ON ai_generation_log(entity_type);
CREATE INDEX idx_ai_generation_log_entity_id ON ai_generation_log(entity_id);
CREATE INDEX idx_ai_generation_log_input_hash ON ai_generation_log(input_hash);
CREATE INDEX idx_ai_generation_log_output_hash ON ai_generation_log(output_hash);
```

### Constraints

```sql
ALTER TABLE ai_generation_log ADD CONSTRAINT chk_ai_generation_log_generation_status 
    CHECK (generation_status IN ('SUCCESS', 'FAILED', 'TIMEOUT', 'VALIDATION_FAILED'));
ALTER TABLE ai_generation_log ADD CONSTRAINT chk_ai_generation_log_duration_ms 
    CHECK (duration_ms >= 0);
ALTER TABLE ai_generation_log ADD CONSTRAINT chk_ai_generation_log_token_input 
    CHECK (token_input >= 0);
ALTER TABLE ai_generation_log ADD CONSTRAINT chk_ai_generation_log_token_output 
    CHECK (token_output >= 0);
```

---

## Table: ai_review_log

**Module Ownership**: AI Orchestration Module

**Description**: Logs all human review actions on AI-generated content, including approvals, rejections, and revision notes. This table provides complete traceability for human-in-the-loop governance.

### Columns

| Column | Data Type | Constraints | Description |
|--------|-----------|-------------|-------------|
| `id` | UUID | PRIMARY KEY | Unique review log identifier |
| `generation_id` | UUID | NOT NULL, FOREIGN KEY → ai_generation_log(id) | Reference to AI generation log |
| `reviewer_id` | UUID | NOT NULL, FOREIGN KEY → users(id) | User who performed the review |
| `review_action` | VARCHAR(20) | NOT NULL | Review action (APPROVED, REJECTED, REVISED, FLAGGED) |
| `review_notes` | TEXT | Reviewer's notes and feedback | |
| `revision_count` | INTEGER | NOT NULL, DEFAULT 0 | Number of revisions made |
| `revision_details` | JSONB | Details of revisions made | |
| `created_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Review timestamp |

### Indexes

```sql
CREATE INDEX idx_ai_review_log_generation_id ON ai_review_log(generation_id);
CREATE INDEX idx_ai_review_log_reviewer_id ON ai_review_log(reviewer_id);
CREATE INDEX idx_ai_review_log_review_action ON ai_review_log(review_action);
CREATE INDEX idx_ai_review_log_created_at ON ai_review_log(created_at);
```

### Constraints

```sql
ALTER TABLE ai_review_log ADD CONSTRAINT chk_ai_review_log_review_action 
    CHECK (review_action IN ('APPROVED', 'REJECTED', 'REVISED', 'FLAGGED'));
ALTER TABLE ai_review_log ADD CONSTRAINT chk_ai_review_log_revision_count 
    CHECK (revision_count >= 0);
```

---

## Table: ai_feedback_log

**Module Ownership**: AI Orchestration Module

**Description**: Logs user feedback on AI-generated content, including satisfaction scores and qualitative feedback. This table enables continuous improvement of AI agents.

### Columns

| Column | Data Type | Constraints | Description |
|--------|-----------|-------------|-------------|
| `id` | UUID | PRIMARY KEY | Unique feedback log identifier |
| `generation_id` | UUID | NOT NULL, FOREIGN KEY → ai_generation_log(id) | Reference to AI generation log |
| `user_id` | UUID | NOT NULL, FOREIGN KEY → users(id) | User who provided feedback |
| `feedback_type` | VARCHAR(20) | NOT NULL | Feedback type (SATISFACTION, QUALITY, ACCURACY, USABILITY) |
| `feedback_score` | INTEGER | CHECK (feedback_score >= 1 AND feedback_score <= 5) | Feedback score (1-5 scale) |
| `feedback_notes` | TEXT | User's qualitative feedback | |
| `would_regenerate` | BOOLEAN | Would user regenerate this output | |
| `created_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Feedback timestamp |

### Indexes

```sql
CREATE INDEX idx_ai_feedback_log_generation_id ON ai_feedback_log(generation_id);
CREATE INDEX idx_ai_feedback_log_user_id ON ai_feedback_log(user_id);
CREATE INDEX idx_ai_feedback_log_feedback_type ON ai_feedback_log(feedback_type);
CREATE INDEX idx_ai_feedback_log_feedback_score ON ai_feedback_log(feedback_score);
CREATE INDEX idx_ai_feedback_log_created_at ON ai_feedback_log(created_at);
```

### Constraints

```sql
ALTER TABLE ai_feedback_log ADD CONSTRAINT chk_ai_feedback_log_feedback_type 
    CHECK (feedback_type IN ('SATISFACTION', 'QUALITY', 'ACCURACY', 'USABILITY'));
ALTER TABLE ai_feedback_log ADD CONSTRAINT chk_ai_feedback_log_feedback_score 
    CHECK (feedback_score >= 1 AND feedback_score <= 5);
```

---

## Migration SQL

```sql
-- Enable UUID extension if not already enabled
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create ai_generation_log table
CREATE TABLE ai_generation_log (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    agent_name VARCHAR(50) NOT NULL,
    prompt_version VARCHAR(20) NOT NULL,
    request_id UUID NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id),
    feature_id VARCHAR(20) NOT NULL,
    entity_type VARCHAR(100) NOT NULL,
    entity_id UUID NOT NULL,
    input_hash VARCHAR(64) NOT NULL,
    output_hash VARCHAR(64) NOT NULL,
    generation_status VARCHAR(20) NOT NULL,
    provider_name VARCHAR(50) NOT NULL,
    model_name VARCHAR(100) NOT NULL,
    duration_ms INTEGER NOT NULL,
    token_input INTEGER NOT NULL,
    token_output INTEGER NOT NULL,
    confidence_score DECIMAL(3,2) CHECK (confidence_score >= 0 AND confidence_score <= 1),
    error_code VARCHAR(50),
    error_message TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_ai_generation_log_generation_status 
        CHECK (generation_status IN ('SUCCESS', 'FAILED', 'TIMEOUT', 'VALIDATION_FAILED')),
    CONSTRAINT chk_ai_generation_log_duration_ms 
        CHECK (duration_ms >= 0),
    CONSTRAINT chk_ai_generation_log_token_input 
        CHECK (token_input >= 0),
    CONSTRAINT chk_ai_generation_log_token_output 
        CHECK (token_output >= 0)
);

-- Create ai_review_log table
CREATE TABLE ai_review_log (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    generation_id UUID NOT NULL REFERENCES ai_generation_log(id),
    reviewer_id UUID NOT NULL REFERENCES users(id),
    review_action VARCHAR(20) NOT NULL,
    review_notes TEXT,
    revision_count INTEGER NOT NULL DEFAULT 0,
    revision_details JSONB,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_ai_review_log_review_action 
        CHECK (review_action IN ('APPROVED', 'REJECTED', 'REVISED', 'FLAGGED')),
    CONSTRAINT chk_ai_review_log_revision_count 
        CHECK (revision_count >= 0)
);

-- Create ai_feedback_log table
CREATE TABLE ai_feedback_log (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    generation_id UUID NOT NULL REFERENCES ai_generation_log(id),
    user_id UUID NOT NULL REFERENCES users(id),
    feedback_type VARCHAR(20) NOT NULL,
    feedback_score INTEGER NOT NULL CHECK (feedback_score >= 1 AND feedback_score <= 5),
    feedback_notes TEXT,
    would_regenerate BOOLEAN,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_ai_feedback_log_feedback_type 
        CHECK (feedback_type IN ('SATISFACTION', 'QUALITY', 'ACCURACY', 'USABILITY'))
);

-- Create indexes for ai_generation_log
CREATE INDEX idx_ai_generation_log_created_at ON ai_generation_log(created_at);
CREATE INDEX idx_ai_generation_log_agent_name ON ai_generation_log(agent_name);
CREATE INDEX idx_ai_generation_log_generation_status ON ai_generation_log(generation_status);
CREATE INDEX idx_ai_generation_log_provider_name ON ai_generation_log(provider_name);
CREATE INDEX idx_ai_generation_log_user_id ON ai_generation_log(user_id);
CREATE INDEX idx_ai_generation_log_request_id ON ai_generation_log(request_id);
CREATE INDEX idx_ai_generation_log_feature_id ON ai_generation_log(feature_id);
CREATE INDEX idx_ai_generation_log_entity_type ON ai_generation_log(entity_type);
CREATE INDEX idx_ai_generation_log_entity_id ON ai_generation_log(entity_id);
CREATE INDEX idx_ai_generation_log_input_hash ON ai_generation_log(input_hash);
CREATE INDEX idx_ai_generation_log_output_hash ON ai_generation_log(output_hash);

-- Create indexes for ai_review_log
CREATE INDEX idx_ai_review_log_generation_id ON ai_review_log(generation_id);
CREATE INDEX idx_ai_review_log_reviewer_id ON ai_review_log(reviewer_id);
CREATE INDEX idx_ai_review_log_review_action ON ai_review_log(review_action);
CREATE INDEX idx_ai_review_log_created_at ON ai_review_log(created_at);

-- Create indexes for ai_feedback_log
CREATE INDEX idx_ai_feedback_log_generation_id ON ai_feedback_log(generation_id);
CREATE INDEX idx_ai_feedback_log_user_id ON ai_feedback_log(user_id);
CREATE INDEX idx_ai_feedback_log_feedback_type ON ai_feedback_log(feedback_type);
CREATE INDEX idx_ai_feedback_log_feedback_score ON ai_feedback_log(feedback_score);
CREATE INDEX idx_ai_feedback_log_created_at ON ai_feedback_log(created_at);
```

---

## Audit Requirements

### Traceability Requirements

All AI operations must be fully traceable:

- **AI-Generated Content**: Every AI-generated artifact must have a corresponding `ai_generation_log` entry with input/output hashes for verification
- **Prompt Versions**: Every generation must log the prompt version used for traceability to prompt lifecycle management
- **Human Reviews**: Every human review action must be logged in `ai_review_log` with reviewer, action, and notes
- **AI Operations**: All AI operations (success, failure, timeout) must be logged with complete context

### Monitoring Requirements

The system must monitor:

- **Generation Success Rate**: Percentage of successful generations per agent
- **Average Generation Time**: Average duration_ms per agent and provider
- **Token Usage**: Total token consumption per agent and provider
- **Confidence Scores**: Average confidence scores per agent
- **Error Rates**: Error rate by error code and agent
- **Review Actions**: Approval/rejection rates per agent
- **User Feedback**: Average feedback scores per agent

### Governance Requirements

The system must support:

- **Audit Queries**: Query all generations by user, agent, date range, or status
- **Replay Capability**: Replay any generation using input_hash and prompt_version
- **Version Tracking**: Track which prompt versions were used for which generations
- **Performance Analysis**: Analyze performance trends over time
- **Cost Tracking**: Track token usage for cost analysis
- **Quality Metrics**: Calculate quality metrics from review and feedback logs

---

## Data Retention

### Retention Policy

- **ai_generation_log**: 7 years (full audit trail)
- **ai_review_log**: 7 years (full audit trail)
- **ai_feedback_log**: 7 years (full audit trail)

### Archival Strategy

After 1 year, move data to archival storage while maintaining query capability:
- Keep recent 1 year in hot storage (fast queries)
- Move older data to cold storage (cost optimization)
- Maintain indexes for audit queries

---

## Conclusion

This AI Audit and Observability Schema provides comprehensive auditability, monitoring, governance, and debugging capabilities for all AI operations. The three audit tables ensure complete traceability of AI-generated content, prompt versions, human reviews, and AI operations, making this the official AI governance schema for NUSA MVP Wave 1.

---

# SECTION 13 — Conclusion

## Database Schema Summary

This Database Schema (14) provides the complete physical database design for NUSA MVP Wave 1:

### Table Count
- **Total Tables**: 18
- **Authentication Module**: 4 tables
- **Curriculum Module**: 2 tables
- **Learning Planning Module**: 2 tables
- **Assessment Module**: 4 tables
- **Reporting Module**: 1 table
- **Administration Module**: 1 table
- **AI Orchestration Module**: 3 tables (ai_generation_log, ai_review_log, ai_feedback_log)

### Schema Characteristics
- **Database Engine**: PostgreSQL 18+
- **Primary Keys**: UUID for all tables
- **Foreign Keys**: Properly defined with CASCADE options
- **Constraints**: Comprehensive CHECK constraints
- **Indexes**: Strategic indexes for performance
- **Timestamps**: Created_at and updated_at on all tables
- **Audit Trail**: Comprehensive audit_logs table
- **JSONB**: Flexible data storage for complex structures

### Implementation Readiness
This database schema is:
- ✅ Complete with all MVP tables
- ✅ Aligned with frozen architecture decisions
- ✅ Ready for PostgreSQL migration generation
- ✅ Ready for backend implementation
- ✅ Optimized for query performance
- ✅ Designed for data integrity

The database schema is officially approved for NUSA MVP Wave 1 implementation.

---

**Document Status**: FOUNDATION DOCUMENT - LOCKED

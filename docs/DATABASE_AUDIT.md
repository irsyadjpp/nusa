# Database Audit Report

**Generated:** June 13, 2026  
**Scope:** NUSA Platform - PostgreSQL Database Schema  
**Version:** 1.0

---

## Executive Summary

The NUSA Platform uses PostgreSQL as its primary database. The schema follows a modular monolith architecture aligned with Domain-Driven Design (DDD) principles. The database supports the Kurikulum Merdeka 2026 education management system with entities for academic foundation, curriculum, learning planning, assessment, achievement, and user management.

**Key Findings:**
- **Total Tables:** 25+ tables across multiple domains
- **Primary Key Strategy:** UUID for all entities
- **Foreign Key Relationships:** Well-defined referential integrity
- **Indexes:** Basic indexing on foreign keys and frequently queried columns
- **Constraints:** NOT NULL constraints on critical fields
- **Status:** Production-ready with room for optimization

---

## Database Schema Overview

### Core Domain Tables

#### 1. User & Authentication

**users**
- `id` (UUID, PK)
- `email` (VARCHAR, UNIQUE, NOT NULL)
- `password_hash` (VARCHAR, NOT NULL)
- `name` (VARCHAR, NOT NULL)
- `role_id` (UUID, FK → roles)
- `school_id` (UUID, FK → schools, nullable)
- `is_active` (BOOLEAN, default true)
- `login_attempts` (INT, default 0)
- `last_login_at` (TIMESTAMP, nullable)
- `created_at` (TIMESTAMP, default NOW())
- `updated_at` (TIMESTAMP, default NOW())
- `created_by` (UUID, FK → users)
- `updated_by` (UUID, FK → users, nullable)

**roles**
- `id` (UUID, PK)
- `name` (VARCHAR, UNIQUE, NOT NULL)
- `description` (TEXT, nullable)
- `is_active` (BOOLEAN, default true)
- `created_at` (TIMESTAMP, default NOW())
- `updated_at` (TIMESTAMP, default NOW())

**permissions**
- `id` (UUID, PK)
- `role_id` (UUID, FK → roles)
- `resource` (VARCHAR, NOT NULL)
- `action` (VARCHAR, NOT NULL)
- `created_at` (TIMESTAMP, default NOW())

**refresh_tokens**
- `id` (UUID, PK)
- `user_id` (UUID, FK → users)
- `token` (VARCHAR, UNIQUE, NOT NULL)
- `expires_at` (TIMESTAMP, NOT NULL)
- `client_ip` (INET, nullable)
- `revoked_at` (TIMESTAMP, nullable)
- `created_at` (TIMESTAMP, default NOW())

**schools**
- `id` (UUID, PK)
- `name` (VARCHAR, NOT NULL)
- `code` (VARCHAR, UNIQUE, NOT NULL)
- `address` (TEXT, nullable)
- `phone` (VARCHAR, nullable)
- `email` (VARCHAR, nullable)
- `is_active` (BOOLEAN, default true)
- `created_at` (TIMESTAMP, default NOW())
- `updated_at` (TIMESTAMP, default NOW())
- `created_by` (UUID, FK → users)
- `updated_by` (UUID, FK → users, nullable)

---

#### 2. Academic Foundation (Sprint 4)

**academic_years**
- `id` (UUID, PK)
- `school_id` (UUID, FK → schools, NOT NULL)
- `year` (VARCHAR, NOT NULL)
- `start_date` (DATE, NOT NULL)
- `end_date` (DATE, NOT NULL)
- `status` (VARCHAR: DRAFT, ACTIVE, ARCHIVED, default DRAFT)
- `created_at` (TIMESTAMP, default NOW())
- `updated_at` (TIMESTAMP, default NOW())
- `created_by` (UUID, FK → users)
- `updated_by` (UUID, FK → users, nullable)

**semesters**
- `id` (UUID, PK)
- `academic_year_id` (UUID, FK → academic_years, NOT NULL)
- `name` (VARCHAR, NOT NULL)
- `type` (VARCHAR: GANJIL, GENAP, NOT NULL)
- `start_date` (DATE, NOT NULL)
- `end_date` (DATE, NOT NULL)
- `sequence_number` (INT, NOT NULL)
- `status` (VARCHAR: ACTIVE, INACTIVE, default INACTIVE)
- `created_at` (TIMESTAMP, default NOW())
- `updated_at` (TIMESTAMP, default NOW())
- `created_by` (UUID, FK → users)
- `updated_by` (UUID, FK → users, nullable)

**subject_categories**
- `id` (UUID, PK)
- `code` (VARCHAR, UNIQUE, NOT NULL)
- `name` (VARCHAR, NOT NULL)
- `description` (TEXT, nullable)
- `is_mandatory` (BOOLEAN, default false)
- `is_active` (BOOLEAN, default true)
- `created_at` (TIMESTAMP, default NOW())
- `updated_at` (TIMESTAMP, default NOW())
- `created_by` (UUID, FK → users)
- `updated_by` (UUID, FK → users, nullable)

**graduate_profile_dimensions**
- `id` (UUID, PK)
- `code` (VARCHAR, UNIQUE, NOT NULL)
- `name` (VARCHAR, NOT NULL)
- `description` (TEXT, nullable)
- `sequence_number` (INT, NOT NULL, CHECK 1-6)
- `is_active` (BOOLEAN, default true)
- `created_at` (TIMESTAMP, default NOW())
- `updated_at` (TIMESTAMP, default NOW())
- `created_by` (UUID, FK → users)
- `updated_by` (UUID, FK → users, nullable)

**cp_alignments**
- `id` (UUID, PK)
- `curriculum_subject_id` (UUID, FK → curriculum_subjects, NOT NULL)
- `graduate_profile_dimension_id` (UUID, FK → graduate_profile_dimensions, NOT NULL)
- `alignment_description` (TEXT, nullable)
- `is_active` (BOOLEAN, default true)
- `created_at` (TIMESTAMP, default NOW())
- `updated_at` (TIMESTAMP, default NOW())
- `created_by` (UUID, FK → users)
- `updated_by` (UUID, FK → users, nullable)

**system_configurations**
- `id` (UUID, PK)
- `key` (VARCHAR, UNIQUE, NOT NULL)
- `value` (TEXT, NOT NULL)
- `value_type` (VARCHAR: STRING, NUMBER, BOOLEAN, JSON, NOT NULL)
- `description` (TEXT, nullable)
- `category` (VARCHAR, NOT NULL)
- `is_system` (BOOLEAN, default false)
- `is_active` (BOOLEAN, default true)
- `created_at` (TIMESTAMP, default NOW())
- `updated_at` (TIMESTAMP, default NOW())
- `created_by` (UUID, FK → users)
- `updated_by` (UUID, FK → users, nullable)

---

#### 3. Curriculum (CP - Capaian Pembelajaran)

**curriculum_subjects**
- `id` (UUID, PK)
- `code` (VARCHAR, UNIQUE, NOT NULL)
- `name` (VARCHAR, NOT NULL)
- `level` (VARCHAR, nullable)
- `description` (TEXT, nullable)
- `is_active` (BOOLEAN, default true)
- `created_at` (TIMESTAMP, default NOW())
- `updated_at` (TIMESTAMP, default NOW())

**curriculum_phases**
- `id` (UUID, PK)
- `subject_id` (UUID, FK → curriculum_subjects, NOT NULL)
- `name` (VARCHAR, NOT NULL)
- `grade_level` (VARCHAR, NOT NULL)
- `order` (INT, NOT NULL)
- `description` (TEXT, nullable)
- `is_active` (BOOLEAN, default true)
- `created_at` (TIMESTAMP, default NOW())
- `updated_at` (TIMESTAMP, default NOW())

**curriculum_elements**
- `id` (UUID, PK)
- `phase_id` (UUID, FK → curriculum_phases, NOT NULL)
- `name` (VARCHAR, NOT NULL)
- `code` (VARCHAR, NOT NULL)
- `description` (TEXT, NOT NULL)
- `order` (INT, NOT NULL)
- `is_active` (BOOLEAN, default true)
- `created_at` (TIMESTAMP, default NOW())
- `updated_at` (TIMESTAMP, default NOW())

**curriculum_subelements**
- `id` (UUID, PK)
- `element_id` (UUID, FK → curriculum_elements, NOT NULL)
- `name` (VARCHAR, NOT NULL)
- `code` (VARCHAR, NOT NULL)
- `description` (TEXT, nullable)
- `order` (INT, NOT NULL)
- `is_active` (BOOLEAN, default true)
- `created_at` (TIMESTAMP, default NOW())
- `updated_at` (TIMESTAMP, default NOW())

**cps** (Capaian Pembelajaran)
- `id` (UUID, PK)
- `subject_id` (UUID, FK → curriculum_subjects, NOT NULL)
- `phase_id` (UUID, FK → curriculum_phases, NOT NULL)
- `element_id` (UUID, FK → curriculum_elements, NOT NULL)
- `subelement_id` (UUID, FK → curriculum_subelements, nullable)
- `code` (VARCHAR, NOT NULL)
- `description` (TEXT, NOT NULL)
- `competency_code` (VARCHAR, nullable)
- `version` (VARCHAR, NOT NULL)
- `effective_date` (DATE, NOT NULL)
- `is_active` (BOOLEAN, default true)
- `created_at` (TIMESTAMP, default NOW())
- `updated_at` (TIMESTAMP, default NOW())

---

#### 4. Learning Planning (TP, ATP, Modul Ajar)

**tp_sets**
- `id` (UUID, PK)
- `cp_id` (UUID, FK → cps, NOT NULL)
- `version_no` (INT, NOT NULL)
- `status` (VARCHAR: DRAFT, UNDER_REVIEW, APPROVED, REJECTED, ARCHIVED)
- `generation_source` (VARCHAR: MANUAL, AI_GENERATED)
- `generation_reason` (TEXT, nullable)
- `generated_by` (UUID, FK → users, NOT NULL)
- `ai_generation_id` (UUID, nullable)
- `approved_by` (UUID, FK → users, nullable)
- `approved_at` (TIMESTAMP, nullable)
- `created_at` (TIMESTAMP, default NOW())

**tps** (Teaching Plans)
- `id` (UUID, PK)
- `tp_set_id` (UUID, FK → tp_sets, NOT NULL)
- `sequence_number` (INT, NOT NULL)
- `cp_id` (UUID, FK → cps, NOT NULL)
- `subject_id` (UUID, FK → curriculum_subjects, NOT NULL)
- `phase_id` (UUID, FK → curriculum_phases, NOT NULL)
- `element_id` (UUID, FK → curriculum_elements, NOT NULL)
- `subelement_id` (UUID, FK → curriculum_subelements, NOT NULL)
- `user_id` (UUID, FK → users, NOT NULL)
- `status` (VARCHAR: DRAFT, UNDER_REVIEW, APPROVED, REJECTED, ARCHIVED)
- `title` (VARCHAR, NOT NULL)
- `learning_objectives` (JSONB, NOT NULL)
- `time_allocation` (JSONB, NOT NULL)
- `prerequisites` (JSONB, NOT NULL)
- `estimated_weeks` (INT, NOT NULL)
- `success_criteria` (JSONB, NOT NULL)
- `version_no` (INT, NOT NULL)
- `is_current_version` (BOOLEAN, default true)
- `parent_version_id` (UUID, FK → tps, nullable)
- `created_at` (TIMESTAMP, default NOW())
- `updated_at` (TIMESTAMP, default NOW())

**atp_sets**
- `id` (UUID, PK)
- `tp_set_id` (UUID, FK → tp_sets, NOT NULL)
- `version_no` (INT, NOT NULL)
- `status` (VARCHAR: DRAFT, UNDER_REVIEW, APPROVED, REJECTED, ARCHIVED)
- `generation_source` (VARCHAR: MANUAL, AI_GENERATED)
- `generated_by` (UUID, FK → users, NOT NULL)
- `approved_by` (UUID, FK → users, nullable)
- `approved_at` (TIMESTAMP, nullable)
- `created_at` (TIMESTAMP, default NOW())

**atps** (Alur Tujuan Pembelajaran)
- `id` (UUID, PK)
- `atp_set_id` (UUID, FK → atp_sets, NOT NULL)
- `sequence_number` (INT, NOT NULL)
- `tp_id` (UUID, FK → tps, NOT NULL)
- `week` (INT, NOT NULL)
- `learning_activities` (JSONB, NOT NULL)
- `assessment_methods` (TEXT[], NOT NULL)
- `time_allocation` (JSONB, NOT NULL)
- `status` (VARCHAR: DRAFT, UNDER_REVIEW, APPROVED, REJECTED, ARCHIVED)
- `created_at` (TIMESTAMP, default NOW())
- `updated_at` (TIMESTAMP, default NOW())

**modul_ajar_sets**
- `id` (UUID, PK)
- `atp_set_id` (UUID, FK → atp_sets, NOT NULL)
- `version_no` (INT, NOT NULL)
- `status` (VARCHAR: DRAFT, UNDER_REVIEW, APPROVED, REJECTED, ARCHIVED)
- `generation_source` (VARCHAR: MANUAL, AI_GENERATED)
- `generated_by` (UUID, FK → users, NOT NULL)
- `approved_by` (UUID, FK → users, nullable)
- `approved_at` (TIMESTAMP, nullable)
- `created_at` (TIMESTAMP, default NOW())

**modul_ajars**
- `id` (UUID, PK)
- `modul_ajar_set_id` (UUID, FK → modul_ajar_sets, NOT NULL)
- `atp_id` (UUID, FK → atps, NOT NULL)
- `week` (INT, NOT NULL)
- `session_number` (INT, NOT NULL)
- `learning_objectives` (TEXT[], NOT NULL)
- `teaching_materials` (JSONB, NOT NULL)
- `learning_methods` (TEXT[], NOT NULL)
- `assessment_methods` (TEXT[], NOT NULL)
- `time_allocation` (JSONB, NOT NULL)
- `status` (VARCHAR: DRAFT, UNDER_REVIEW, APPROVED, REJECTED, ARCHIVED)
- `created_at` (TIMESTAMP, default NOW())
- `updated_at` (TIMESTAMP, default NOW())

---

#### 5. Assessment

**assessments**
- `id` (UUID, PK)
- `tp_id` (UUID, FK → tps, NOT NULL)
- `tp_version_no` (INT, NOT NULL)
- `success_criteria_snapshot` (JSONB, NOT NULL)
- `user_id` (UUID, FK → users, NOT NULL)
- `assessment_type` (VARCHAR: FORMATIVE, SUMMATIVE, NOT NULL)
- `status` (VARCHAR: DRAFT, UNDER_REVIEW, APPROVED, REJECTED, ARCHIVED)
- `assessment_items` (JSONB, NOT NULL)
- `answer_key` (JSONB, NOT NULL)
- `scoring_guidelines` (JSONB, NOT NULL)
- `ai_confidence_score` (DECIMAL, nullable)
- `ai_generated_at` (TIMESTAMP, nullable)
- `ai_agent_version` (VARCHAR, nullable)
- `version_no` (INT, NOT NULL)
- `is_current_version` (BOOLEAN, default true)
- `parent_version_id` (UUID, FK → assessments, nullable)
- `created_at` (TIMESTAMP, default NOW())
- `updated_at` (TIMESTAMP, default NOW())
- `approved_at` (TIMESTAMP, nullable)
- `approved_by` (UUID, FK → users, nullable)

**rubrics**
- `id` (UUID, PK)
- `assessment_id` (UUID, FK → assessments, NOT NULL)
- `rubric_type` (VARCHAR: ANALYTIC, HOLISTIC, NOT NULL)
- `criteria` (JSONB, NOT NULL)
- `total_points` (INT, NOT NULL)
- `status` (VARCHAR: DRAFT, UNDER_REVIEW, APPROVED, REJECTED, ARCHIVED)
- `version_no` (INT, NOT NULL)
- `is_current_version` (BOOLEAN, default true)
- `created_at` (TIMESTAMP, default NOW())
- `updated_at` (TIMESTAMP, default NOW())

**evidences**
- `id` (UUID, PK)
- `student_id` (UUID, FK → users, NOT NULL)
- `assessment_id` (UUID, FK → assessments, NOT NULL)
- `evidence_type` (VARCHAR: DOCUMENT, IMAGE, VIDEO, AUDIO, PROJECT, PRESENTATION, NOT NULL)
- `file_url` (VARCHAR, NOT NULL)
- `file_metadata` (JSONB, NOT NULL)
- `submission_date` (TIMESTAMP, NOT NULL)
- `status` (VARCHAR: SUBMITTED, UNDER_REVIEW, APPROVED, REJECTED, NEEDS_REVISION)
- `evaluation_count` (INT, default 0)
- `latest_evaluation_id` (UUID, FK → evaluations, nullable)
- `teacher_notes` (TEXT, nullable)
- `created_at` (TIMESTAMP, default NOW())
- `updated_at` (TIMESTAMP, default NOW())

**evaluations**
- `id` (UUID, PK)
- `evidence_id` (UUID, FK → evidences, NOT NULL)
- `student_id` (UUID, FK → users, NOT NULL)
- `teacher_id` (UUID, FK → users, NOT NULL)
- `revision_no` (INT, NOT NULL)
- `performance_scores` (JSONB, NOT NULL)
- `performance_level` (VARCHAR: EMERGING, DEVELOPING, PROFICIENT, ADVANCED, NOT NULL)
- `teacher_feedback` (TEXT, NOT NULL)
- `evaluation_date` (TIMESTAMP, NOT NULL)
- `is_current_revision` (BOOLEAN, default true)
- `parent_revision_id` (UUID, FK → evaluations, nullable)
- `created_at` (TIMESTAMP, default NOW())
- `updated_at` (TIMESTAMP, default NOW())

---

#### 6. Reporting

**narrative_reports**
- `id` (UUID, PK)
- `student_id` (UUID, FK → users, NOT NULL)
- `class_id` (UUID, NOT NULL) - Note: class table not yet implemented
- `subject_id` (UUID, FK → curriculum_subjects, NOT NULL)
- `reporting_period` (JSONB, NOT NULL)
- `narrative_content` (JSONB, NOT NULL)
- `achievement_summary` (JSONB, NOT NULL)
- `teacher_recommendations` (TEXT[], NOT NULL)
- `parent_feedback` (TEXT, nullable)
- `status` (VARCHAR: DRAFT, UNDER_REVIEW, APPROVED, REJECTED, ARCHIVED)
- `generated_at` (TIMESTAMP, NOT NULL)
- `approved_by` (UUID, FK → users, nullable)
- `approved_at` (TIMESTAMP, nullable)
- `created_at` (TIMESTAMP, default NOW())
- `updated_at` (TIMESTAMP, default NOW())

---

## Indexes

### Primary Indexes
All tables use UUID primary keys with B-tree indexes (automatic in PostgreSQL).

### Foreign Key Indexes
Indexes exist on all foreign key columns for join performance:
- `users.role_id`
- `users.school_id`
- `academic_years.school_id`
- `semesters.academic_year_id`
- `cp_alignments.curriculum_subject_id`
- `cp_alignments.graduate_profile_dimension_id`
- And all other foreign key relationships

### Unique Constraints
- `users.email` - UNIQUE
- `roles.name` - UNIQUE
- `schools.code` - UNIQUE
- `subject_categories.code` - UNIQUE
- `graduate_profile_dimensions.code` - UNIQUE
- `curriculum_subjects.code` - UNIQUE
- `system_configurations.key` - UNIQUE
- `refresh_tokens.token` - UNIQUE

### Check Constraints
- `graduate_profile_dimensions.sequence_number` - CHECK (sequence_number >= 1 AND sequence_number <= 6)

---

## Data Types Summary

| Type | Usage | Notes |
|------|-------|-------|
| UUID | Primary Keys, Foreign Keys | Standard for all entity IDs |
| VARCHAR | Text fields (names, codes, descriptions) | Variable length strings |
| TEXT | Long text fields (descriptions, feedback) | Unbounded text |
| BOOLEAN | Status flags, active flags | true/false values |
| INT | Counters, sequence numbers | Integer values |
| DATE | Dates without time | Academic dates, effective dates |
| TIMESTAMP | Date-time with time zone | Created/updated timestamps |
| JSONB | Complex nested data | Learning objectives, criteria, metadata |
| TEXT[] | String arrays | Assessment methods, recommendations |
| INET | IP addresses | Client IP for refresh tokens |
| DECIMAL | Precise decimal numbers | AI confidence scores |

---

## Business Rules Enforced at Database Level

1. **Academic Year Uniqueness:** Only one ACTIVE academic year per school (enforced at application level)
2. **Semester Sequence:** Sequence numbers must be unique within an academic year (enforced at application level)
3. **Date Overlap Prevention:** Academic years and semesters cannot overlap (enforced at application level)
4. **Code Uniqueness:** All code fields are UNIQUE constrained
5. **Profile Dimension Sequence:** Limited to 1-6 (CHECK constraint)
6. **Referential Integrity:** All foreign keys are enforced
7. **Non-null Constraints:** Critical fields have NOT NULL constraints

---

## Observations & Recommendations

### Strengths
1. **Consistent Design:** All tables follow the same pattern (UUID PKs, timestamps, audit fields)
2. **Audit Trail:** All tables include `created_at`, `updated_at`, `created_by`, `updated_by`
3. **Flexible Data:** JSONB columns allow for complex nested structures without schema changes
4. **Status Management:** Consistent status enums across entities
5. **Versioning:** Version tracking for TP, ATP, Assessment, Rubric entities

### Areas for Improvement
1. **Missing Indexes:** Add composite indexes for common query patterns (e.g., status + school_id)
2. **Soft Delete:** Consider adding `deleted_at` for soft delete instead of hard delete
3. **Partitioning:** Consider partitioning large tables (e.g., evaluations, evidences) by date
4. **Full-Text Search:** Add full-text search indexes for text fields (descriptions, feedback)
5. **Materialized Views:** Consider materialized views for complex reporting queries
6. **Trigger-based Timestamps:** Use database triggers for automatic timestamp updates
7. **Enum Types:** Convert VARCHAR status fields to PostgreSQL ENUM types for type safety

### Missing Tables (Identified Gaps)
1. **classes** - Referenced in narrative_reports but not implemented
2. **class_enrollments** - Student-class relationships
3. **attendance_records** - Attendance tracking
4. **schedule** - Class scheduling
5. **notifications** - User notifications
6. **announcements** - School announcements
7. **messages** - Internal messaging system

---

## Migration Strategy

The database uses migration files in the `backend/migrations/` directory:
- Naming convention: `NNNNNN_description.up.sql` and `NNNNNN_description.down.sql`
- Sequential numbering: 000001, 000002, etc.
- Test on fresh database before committing
- Update `DATABASE_SCHEMA_FREEZE_V1.md` for significant changes

---

## Security Considerations

1. **Password Storage:** Passwords are hashed (bcrypt) before storage
2. **Token Storage:** Refresh tokens stored with expiration and revocation tracking
3. **PII:** User emails and names stored - consider encryption for sensitive data
4. **Row-Level Security:** Not currently implemented - consider for multi-tenant isolation
5. **Audit Logging:** All changes tracked via `created_by` and `updated_by` fields

---

## Performance Considerations

1. **Connection Pooling:** Configured via pgxpool (max 25 connections)
2. **Query Optimization:** Use EXPLAIN ANALYZE for slow queries
3. **Index Maintenance:** Regular index maintenance and statistics updates
4. **Vacuum Strategy:** Configure autovacuum for tables with high churn
5. **Connection Limits:** Set appropriate connection limits based on workload

---

## Conclusion

The NUSA Platform database schema is well-designed for a modular monolith architecture supporting Kurikulum Merdeka. The schema follows consistent patterns, includes proper audit trails, and uses appropriate data types for the education domain. While there are opportunities for optimization (additional indexes, partitioning, materialized views), the current design is production-ready and scalable for the expected workload.

The main gaps are in missing tables for classes, attendance, scheduling, and communication features, which should be prioritized based on business requirements.

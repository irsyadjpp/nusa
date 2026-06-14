# Database Design

**Version:** 1.0  
**Date:** June 13, 2026  
**Based On:** DATABASE_AUDIT.md, GAP_ANALYSIS.md, TARGET_ARCHITECTURE.md

---

## Executive Summary

This document describes the target database design for the NUSA Platform, including new tables, modifications to existing tables, index strategy, and migration strategy. The design maintains the existing PostgreSQL schema while adding critical missing functionality identified in the gap analysis.

**Key Changes:**
- **New Tables:** 15 tables (classes, attendance, scheduling, communication, audit, etc.)
- **Modified Tables:** 25 tables (add deleted_at, audit fields, indexes)
- **Removed Tables:** 0
- **New Indexes:** 40+ composite and partial indexes
- **Migration Strategy:** Incremental with rollback support

---

## Module: Class Management

### New Tables

#### classes

```sql
CREATE TABLE classes (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  school_id UUID NOT NULL REFERENCES schools(id) ON DELETE CASCADE,
  academic_year_id UUID NOT NULL REFERENCES academic_years(id) ON DELETE CASCADE,
  semester_id UUID NOT NULL REFERENCES semesters(id) ON DELETE CASCADE,
  subject_id UUID NOT NULL REFERENCES curriculum_subjects(id) ON DELETE CASCADE,
  teacher_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  name VARCHAR(255) NOT NULL,
  grade_level VARCHAR(50) NOT NULL,
  room VARCHAR(100),
  max_students INTEGER DEFAULT 40,
  is_active BOOLEAN DEFAULT true,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  created_by UUID REFERENCES users(id),
  updated_by UUID REFERENCES users(id),
  deleted_at TIMESTAMP
);

-- Indexes
CREATE INDEX idx_classes_school_id ON classes(school_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_classes_academic_year_id ON classes(academic_year_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_classes_semester_id ON classes(semester_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_classes_subject_id ON classes(subject_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_classes_teacher_id ON classes(teacher_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_classes_school_semester ON classes(school_id, semester_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_classes_active ON classes(is_active) WHERE deleted_at IS NULL AND is_active = true;
```

#### class_enrollments

```sql
CREATE TABLE class_enrollments (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  class_id UUID NOT NULL REFERENCES classes(id) ON DELETE CASCADE,
  student_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  enrollment_date DATE NOT NULL DEFAULT CURRENT_DATE,
  status VARCHAR(50) DEFAULT 'ACTIVE' CHECK (status IN ('ACTIVE', 'INACTIVE', 'WITHDRAWN', 'COMPLETED')),
  notes TEXT,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  deleted_at TIMESTAMP,
  UNIQUE(class_id, student_id)
);

-- Indexes
CREATE INDEX idx_class_enrollments_class_id ON class_enrollments(class_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_class_enrollments_student_id ON class_enrollments(student_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_class_enrollments_status ON class_enrollments(status) WHERE deleted_at IS NULL;
CREATE INDEX idx_class_enrollments_active ON class_enrollments(status) WHERE deleted_at IS NULL AND status = 'ACTIVE';
```

---

## Module: Attendance

### New Tables

#### attendance_records

```sql
CREATE TABLE attendance_records (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  class_id UUID NOT NULL REFERENCES classes(id) ON DELETE CASCADE,
  student_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  date DATE NOT NULL,
  status VARCHAR(50) NOT NULL CHECK (status IN ('PRESENT', 'ABSENT', 'LATE', 'EXCUSED')),
  notes TEXT,
  recorded_by UUID NOT NULL REFERENCES users(id),
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  deleted_at TIMESTAMP,
  UNIQUE(class_id, student_id, date)
);

-- Indexes
CREATE INDEX idx_attendance_records_class_id ON attendance_records(class_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_attendance_records_student_id ON attendance_records(student_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_attendance_records_date ON attendance_records(date) WHERE deleted_at IS NULL;
CREATE INDEX idx_attendance_records_status ON attendance_records(status) WHERE deleted_at IS NULL;
CREATE INDEX idx_attendance_records_class_date ON attendance_records(class_id, date) WHERE deleted_at IS NULL;
CREATE INDEX idx_attendance_records_student_date ON attendance_records(student_id, date) WHERE deleted_at IS NULL;
CREATE INDEX idx_attendance_records_class_date_student ON attendance_records(class_id, date, student_id) WHERE deleted_at IS NULL;
```

---

## Module: Scheduling

### New Tables

#### schedules

```sql
CREATE TABLE schedules (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  class_id UUID NOT NULL REFERENCES classes(id) ON DELETE CASCADE,
  day_of_week INTEGER NOT NULL CHECK (day_of_week BETWEEN 1 AND 7),
  start_time TIME NOT NULL,
  end_time TIME NOT NULL,
  room VARCHAR(100),
  is_active BOOLEAN DEFAULT true,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  created_by UUID REFERENCES users(id),
  updated_by UUID REFERENCES users(id),
  deleted_at TIMESTAMP
);

-- Indexes
CREATE INDEX idx_schedules_class_id ON schedules(class_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_schedules_day_of_week ON schedules(day_of_week) WHERE deleted_at IS NULL;
CREATE INDEX idx_schedules_time ON schedules(start_time, end_time) WHERE deleted_at IS NULL;
CREATE INDEX idx_schedules_room ON schedules(room) WHERE deleted_at IS NULL;
CREATE INDEX idx_schedules_active ON schedules(is_active) WHERE deleted_at IS NULL AND is_active = true;
CREATE INDEX idx_schedules_day_time ON schedules(day_of_week, start_time, end_time) WHERE deleted_at IS NULL;
```

---

## Module: Exam & Assignment

### New Tables

#### exams

```sql
CREATE TABLE exams (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  class_id UUID NOT NULL REFERENCES classes(id) ON DELETE CASCADE,
  assessment_id UUID NOT NULL REFERENCES assessments(id) ON DELETE CASCADE,
  exam_date DATE NOT NULL,
  start_time TIME NOT NULL,
  duration_minutes INTEGER NOT NULL,
  room VARCHAR(100),
  status VARCHAR(50) DEFAULT 'SCHEDULED' CHECK (status IN ('SCHEDULED', 'IN_PROGRESS', 'COMPLETED', 'CANCELLED')),
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  created_by UUID REFERENCES users(id),
  updated_by UUID REFERENCES users(id),
  deleted_at TIMESTAMP
);

-- Indexes
CREATE INDEX idx_exams_class_id ON exams(class_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_exams_assessment_id ON exams(assessment_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_exams_exam_date ON exams(exam_date) WHERE deleted_at IS NULL;
CREATE INDEX idx_exams_status ON exams(status) WHERE deleted_at IS NULL;
CREATE INDEX idx_exams_class_date ON exams(class_id, exam_date) WHERE deleted_at IS NULL;
CREATE INDEX idx_exams_upcoming ON exams(exam_date) WHERE deleted_at IS NULL AND exam_date >= CURRENT_DATE AND status = 'SCHEDULED';
```

#### assignments

```sql
CREATE TABLE assignments (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  class_id UUID NOT NULL REFERENCES classes(id) ON DELETE CASCADE,
  assessment_id UUID NOT NULL REFERENCES assessments(id) ON DELETE CASCADE,
  title VARCHAR(255) NOT NULL,
  description TEXT,
  due_date TIMESTAMP NOT NULL,
  max_score INTEGER,
  status VARCHAR(50) DEFAULT 'ASSIGNED' CHECK (status IN ('ASSIGNED', 'IN_PROGRESS', 'SUBMITTED', 'GRADED', 'CANCELLED')),
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  created_by UUID REFERENCES users(id),
  updated_by UUID REFERENCES users(id),
  deleted_at TIMESTAMP
);

-- Indexes
CREATE INDEX idx_assignments_class_id ON assignments(class_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_assignments_assessment_id ON assignments(assessment_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_assignments_due_date ON assignments(due_date) WHERE deleted_at IS NULL;
CREATE INDEX idx_assignments_status ON assignments(status) WHERE deleted_at IS NULL;
CREATE INDEX idx_assignments_class_due ON assignments(class_id, due_date) WHERE deleted_at IS NULL;
CREATE INDEX idx_assignments_upcoming ON assignments(due_date) WHERE deleted_at IS NULL AND due_date >= NOW() AND status = 'ASSIGNED';
```

#### exam_results

```sql
CREATE TABLE exam_results (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  exam_id UUID NOT NULL REFERENCES exams(id) ON DELETE CASCADE,
  student_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  score DECIMAL(5,2),
  grade VARCHAR(10),
  remarks TEXT,
  submitted_at TIMESTAMP,
  graded_at TIMESTAMP,
  graded_by UUID REFERENCES users(id),
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  deleted_at TIMESTAMP,
  UNIQUE(exam_id, student_id)
);

-- Indexes
CREATE INDEX idx_exam_results_exam_id ON exam_results(exam_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_exam_results_student_id ON exam_results(student_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_exam_results_score ON exam_results(score) WHERE deleted_at IS NULL;
CREATE INDEX idx_exam_results_grade ON exam_results(grade) WHERE deleted_at IS NULL;
CREATE INDEX idx_exam_results_exam_student ON exam_results(exam_id, student_id) WHERE deleted_at IS NULL;
```

---

## Module: Communication

### New Tables

#### notifications

```sql
CREATE TABLE notifications (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  title VARCHAR(255) NOT NULL,
  message TEXT NOT NULL,
  type VARCHAR(50) NOT NULL CHECK (type IN ('INFO', 'WARNING', 'ERROR', 'SUCCESS')),
  is_read BOOLEAN DEFAULT false,
  read_at TIMESTAMP,
  action_url VARCHAR(500),
  metadata JSONB,
  created_at TIMESTAMP DEFAULT NOW(),
  deleted_at TIMESTAMP
);

-- Indexes
CREATE INDEX idx_notifications_user_id ON notifications(user_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_notifications_type ON notifications(type) WHERE deleted_at IS NULL;
CREATE INDEX idx_notifications_is_read ON notifications(is_read) WHERE deleted_at IS NULL;
CREATE INDEX idx_notifications_created_at ON notifications(created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_notifications_user_unread ON notifications(user_id, is_read) WHERE deleted_at IS NULL AND is_read = false;
CREATE INDEX idx_notifications_user_created ON notifications(user_id, created_at DESC) WHERE deleted_at IS NULL;
```

#### announcements

```sql
CREATE TABLE announcements (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  school_id UUID NOT NULL REFERENCES schools(id) ON DELETE CASCADE,
  title VARCHAR(255) NOT NULL,
  content TEXT NOT NULL,
  priority VARCHAR(50) DEFAULT 'NORMAL' CHECK (priority IN ('LOW', 'NORMAL', 'HIGH', 'URGENT')),
  target_audience VARCHAR(50) CHECK (target_audience IN ('ALL', 'TEACHERS', 'STUDENTS', 'PARENTS', 'ADMIN')),
  published_by UUID NOT NULL REFERENCES users(id),
  published_at TIMESTAMP DEFAULT NOW(),
  expires_at TIMESTAMP,
  is_active BOOLEAN DEFAULT true,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  deleted_at TIMESTAMP
);

-- Indexes
CREATE INDEX idx_announcements_school_id ON announcements(school_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_announcements_priority ON announcements(priority) WHERE deleted_at IS NULL;
CREATE INDEX idx_announcements_target_audience ON announcements(target_audience) WHERE deleted_at IS NULL;
CREATE INDEX idx_announcements_published_at ON announcements(published_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_announcements_active ON announcements(is_active, expires_at) WHERE deleted_at IS NULL AND is_active = true AND (expires_at IS NULL OR expires_at > NOW());
```

#### messages

```sql
CREATE TABLE messages (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  sender_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  receiver_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  subject VARCHAR(255),
  content TEXT NOT NULL,
  is_read BOOLEAN DEFAULT false,
  read_at TIMESTAMP,
  parent_message_id UUID REFERENCES messages(id) ON DELETE CASCADE,
  created_at TIMESTAMP DEFAULT NOW(),
  deleted_at TIMESTAMP
);

-- Indexes
CREATE INDEX idx_messages_sender_id ON messages(sender_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_messages_receiver_id ON messages(receiver_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_messages_is_read ON messages(is_read) WHERE deleted_at IS NULL;
CREATE INDEX idx_messages_created_at ON messages(created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_messages_conversation ON messages(LEAST(sender_id, receiver_id), GREATEST(sender_id, receiver_id), created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_messages_receiver_unread ON messages(receiver_id, is_read) WHERE deleted_at IS NULL AND is_read = false;
```

---

## Module: Audit & Logging

### New Tables

#### audit_logs

```sql
CREATE TABLE audit_logs (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID REFERENCES users(id) ON DELETE SET NULL,
  action VARCHAR(100) NOT NULL,
  entity_type VARCHAR(100) NOT NULL,
  entity_id UUID,
  old_values JSONB,
  new_values JSONB,
  ip_address INET,
  user_agent TEXT,
  request_id VARCHAR(100),
  created_at TIMESTAMP DEFAULT NOW()
);

-- Indexes
CREATE INDEX idx_audit_logs_user_id ON audit_logs(user_id);
CREATE INDEX idx_audit_logs_entity ON audit_logs(entity_type, entity_id);
CREATE INDEX idx_audit_logs_action ON audit_logs(action);
CREATE INDEX idx_audit_logs_created_at ON audit_logs(created_at DESC);
CREATE INDEX idx_audit_logs_user_created ON audit_logs(user_id, created_at DESC);
CREATE INDEX idx_audit_logs_entity_created ON audit_logs(entity_type, entity_id, created_at DESC);
```

#### job_queue

```sql
CREATE TABLE job_queue (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  job_type VARCHAR(100) NOT NULL,
  payload JSONB NOT NULL,
  status VARCHAR(50) DEFAULT 'PENDING' CHECK (status IN ('PENDING', 'PROCESSING', 'COMPLETED', 'FAILED', 'CANCELLED')),
  priority INTEGER DEFAULT 5,
  max_retries INTEGER DEFAULT 3,
  retry_count INTEGER DEFAULT 0,
  error_message TEXT,
  started_at TIMESTAMP,
  completed_at TIMESTAMP,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);

-- Indexes
CREATE INDEX idx_job_queue_status ON job_queue(status);
CREATE INDEX idx_job_queue_job_type ON job_queue(job_type);
CREATE INDEX idx_job_queue_priority ON job_queue(priority);
CREATE INDEX idx_job_queue_created_at ON job_queue(created_at);
CREATE INDEX idx_job_queue_pending ON job_queue(status, priority, created_at) WHERE status = 'PENDING';
CREATE INDEX idx_job_queue_retry ON job_queue(status, retry_count, max_retries) WHERE status = 'FAILED' AND retry_count < max_retries;
```

---

## Modified Tables

### Add deleted_at to All Tables

```sql
-- Add soft delete to all existing tables
ALTER TABLE users ADD COLUMN deleted_at TIMESTAMP;
ALTER TABLE roles ADD COLUMN deleted_at TIMESTAMP;
ALTER TABLE permissions ADD COLUMN deleted_at TIMESTAMP;
ALTER TABLE refresh_tokens ADD COLUMN deleted_at TIMESTAMP;
ALTER TABLE schools ADD COLUMN deleted_at TIMESTAMP;
ALTER TABLE academic_years ADD COLUMN deleted_at TIMESTAMP;
ALTER TABLE semesters ADD COLUMN deleted_at TIMESTAMP;
ALTER TABLE subject_categories ADD COLUMN deleted_at TIMESTAMP;
ALTER TABLE graduate_profile_dimensions ADD COLUMN deleted_at TIMESTAMP;
ALTER TABLE cp_alignments ADD COLUMN deleted_at TIMESTAMP;
ALTER TABLE system_configurations ADD COLUMN deleted_at TIMESTAMP;
ALTER TABLE curriculum_subjects ADD COLUMN deleted_at TIMESTAMP;
ALTER TABLE curriculum_phases ADD COLUMN deleted_at TIMESTAMP;
ALTER TABLE curriculum_elements ADD COLUMN deleted_at TIMESTAMP;
ALTER TABLE curriculum_subelements ADD COLUMN deleted_at TIMESTAMP;
ALTER TABLE cps ADD COLUMN deleted_at TIMESTAMP;
ALTER TABLE tp_sets ADD COLUMN deleted_at TIMESTAMP;
ALTER TABLE tps ADD COLUMN deleted_at TIMESTAMP;
ALTER TABLE atp_sets ADD COLUMN deleted_at TIMESTAMP;
ALTER TABLE atps ADD COLUMN deleted_at TIMESTAMP;
ALTER TABLE modul_ajar_sets ADD COLUMN deleted_at TIMESTAMP;
ALTER TABLE modul_ajars ADD COLUMN deleted_at TIMESTAMP;
ALTER TABLE assessments ADD COLUMN deleted_at TIMESTAMP;
ALTER TABLE rubrics ADD COLUMN deleted_at TIMESTAMP;
ALTER TABLE evidences ADD COLUMN deleted_at TIMESTAMP;
ALTER TABLE evaluations ADD COLUMN deleted_at TIMESTAMP;
ALTER TABLE narrative_reports ADD COLUMN deleted_at TIMESTAMP;
```

### Add Partial Indexes for Soft Delete

```sql
-- Add partial indexes for active records
CREATE INDEX idx_users_active ON users(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_roles_active ON roles(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_schools_active ON schools(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_academic_years_active ON academic_years(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_semesters_active ON semesters(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_subject_categories_active ON subject_categories(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_graduate_profile_dimensions_active ON graduate_profile_dimensions(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_cp_alignments_active ON cp_alignments(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_system_configurations_active ON system_configurations(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_curriculum_subjects_active ON curriculum_subjects(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_curriculum_phases_active ON curriculum_phases(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_curriculum_elements_active ON curriculum_elements(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_curriculum_subelements_active ON curriculum_subelements(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_cps_active ON cps(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_tp_sets_active ON tp_sets(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_tps_active ON tps(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_atp_sets_active ON atp_sets(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_atps_active ON atps(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_modul_ajar_sets_active ON modul_ajar_sets(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_modul_ajars_active ON modul_ajars(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_assessments_active ON assessments(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_rubrics_active ON rubrics(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_evidences_active ON evidences(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_evaluations_active ON evaluations(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_narrative_reports_active ON narrative_reports(id) WHERE deleted_at IS NULL;
```

### Add Composite Indexes for Common Queries

```sql
-- Academic Foundation
CREATE INDEX idx_academic_years_school_status ON academic_years(school_id, status) WHERE deleted_at IS NULL;
CREATE INDEX idx_semesters_academic_year_status ON semesters(academic_year_id, status) WHERE deleted_at IS NULL;
CREATE INDEX idx_cp_alignments_subject_dimension ON cp_alignments(curriculum_subject_id, graduate_profile_dimension_id) WHERE deleted_at IS NULL;

-- Curriculum
CREATE INDEX idx_cps_subject_phase ON cps(subject_id, phase_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_cps_phase_element ON cps(phase_id, element_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_curriculum_phases_subject ON curriculum_phases(subject_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_curriculum_elements_phase ON curriculum_elements(phase_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_curriculum_subelements_element ON curriculum_subelements(element_id) WHERE deleted_at IS NULL;

-- Learning Planning
CREATE INDEX idx_tps_tp_set ON tps(tp_set_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_tps_subject_phase ON tps(subject_id, phase_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_atps_atp_set ON atps(atp_set_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_modul_ajars_modul_ajar_set ON modul_ajars(modul_ajar_set_id) WHERE deleted_at IS NULL;

-- Assessment
CREATE INDEX idx_assessments_tp ON assessments(tp_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_assessments_user_status ON assessments(user_id, status) WHERE deleted_at IS NULL;
CREATE INDEX idx_rubrics_assessment ON rubrics(assessment_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_evidences_assessment ON evidences(assessment_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_evidences_student ON evidences(student_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_evaluations_evidence ON evaluations(evidence_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_evaluations_student ON evaluations(student_id) WHERE deleted_at IS NULL;

-- User Management
CREATE INDEX idx_users_school_role ON users(school_id, role_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_users_email_active ON users(email) WHERE deleted_at IS NULL AND is_active = true;
```

### Add Full-Text Search Indexes

```sql
-- Add full-text search for text fields
CREATE INDEX idx_cps_text_search ON cps USING gin(to_tsvector('english', description)) WHERE deleted_at IS NULL;
CREATE INDEX idx_tps_title_search ON tps USING gin(to_tsvector('english', title)) WHERE deleted_at IS NULL;
CREATE INDEX idx_assessments_title_search ON assessments USING gin(to_tsvector('english', title)) WHERE deleted_at IS NULL;
CREATE INDEX idx_announcements_content_search ON announcements USING gin(to_tsvector('english', content)) WHERE deleted_at IS NULL;
CREATE INDEX idx_messages_content_search ON messages USING gin(to_tsvector('english', content)) WHERE deleted_at IS NULL;
```

---

## ERD Description

### Entity Relationships

```
schools (1) ────< (N) academic_years
schools (1) ────< (N) semesters (via academic_years)
schools (1) ────< (N) users
schools (1) ────< (N) classes
schools (1) ────< (N) announcements

academic_years (1) ────< (N) semesters
academic_years (1) ────< (N) classes

semesters (1) ────< (N) classes

curriculum_subjects (1) ────< (N) curriculum_phases
curriculum_subjects (1) ────< (N) cps
curriculum_subjects (1) ────< (N) cp_alignments
curriculum_subjects (1) ────< (N) classes

curriculum_phases (1) ────< (N) curriculum_elements
curriculum_phases (1) ────< (N) cps

curriculum_elements (1) ────< (N) curriculum_subelements
curriculum_elements (1) ────< (N) cps

curriculum_subelements (1) ────< (N) cps

cps (1) ────< (N) tp_sets
cps (1) ────< (N) cp_alignments

tp_sets (1) ────< (N) tps
tp_sets (1) ────< (N) atp_sets

tps (1) ────< (N) atps
tps (1) ────< (N) assessments

atp_sets (1) ────< (N) atps
atp_sets (1) ────< (N) modul_ajar_sets

atps (1) ────< (N) modul_ajars

modul_ajar_sets (1) ────< (N) modul_ajars

assessments (1) ────< (N) rubrics
assessments (1) ────< (N) evidences
assessments (1) ────< (N) exams
assessments (1) ────< (N) assignments

evidences (1) ────< (N) evaluations

classes (1) ────< (N) class_enrollments
classes (1) ────< (N) attendance_records
classes (1) ────< (N) schedules
classes (1) ────< (N) exams
classes (1) ────< (N) assignments

users (1) ────< (N) class_enrollments (as student)
users (1) ────< (N) classes (as teacher)
users (1) ────< (N) attendance_records (as student)
users (1) ────< (N) attendance_records (as recorded_by)
users (1) ────< (N) exam_results (as student)
users (1) ────< (N) exam_results (as graded_by)
users (1) ────< (N) notifications
users (1) ────< (N) messages (as sender)
users (1) ────< (N) messages (as receiver)
users (1) ────< (N) audit_logs
users (1) ────< (N) narrative_reports (as student)
users (1) ────< (N) narrative_reports (as approved_by)

graduate_profile_dimensions (1) ────< (N) cp_alignments
```

---

## Index Strategy

### Primary Indexes
- All tables use UUID primary keys with B-tree indexes (automatic in PostgreSQL)

### Foreign Key Indexes
- All foreign key columns have indexes for join performance

### Composite Indexes
- Added for common query patterns (school_id + status, academic_year_id + status, etc.)
- Optimized for WHERE clause filtering

### Partial Indexes
- Added for soft delete (WHERE deleted_at IS NULL)
- Added for active records (WHERE is_active = true)
- Added for upcoming items (WHERE date >= CURRENT_DATE)

### Full-Text Search Indexes
- Added GIN indexes for text search on description, title, content fields
- Uses English language configuration

### Unique Constraints
- All code fields have UNIQUE constraints
- Email has UNIQUE constraint
- Composite unique constraints for business rules (class_id + student_id, etc.)

---

## Migration Strategy

### Migration File Naming Convention

```
NNNNNN_description.up.sql
NNNNNN_description.down.sql
```

### Migration Order

#### Phase 1: Infrastructure Foundation

**Migration 000001: Add deleted_at to all tables**
- Add deleted_at column to all existing tables
- Add partial indexes for active records
- Rollback: Remove deleted_at columns and partial indexes

**Migration 000002: Add composite indexes**
- Add composite indexes for common queries
- Rollback: Remove composite indexes

**Migration 000003: Add full-text search indexes**
- Add GIN indexes for text search
- Rollback: Remove GIN indexes

#### Phase 2: Class Management Module

**Migration 000004: Create classes table**
- Create classes table with indexes
- Rollback: Drop classes table

**Migration 000005: Create class_enrollments table**
- Create class_enrollments table with indexes
- Rollback: Drop class_enrollments table

#### Phase 3: Attendance Module

**Migration 000006: Create attendance_records table**
- Create attendance_records table with indexes
- Rollback: Drop attendance_records table

#### Phase 4: Scheduling Module

**Migration 000007: Create schedules table**
- Create schedules table with indexes
- Rollback: Drop schedules table

#### Phase 5: Exam & Assignment Module

**Migration 000008: Create exams table**
- Create exams table with indexes
- Rollback: Drop exams table

**Migration 000009: Create assignments table**
- Create assignments table with indexes
- Rollback: Drop assignments table

**Migration 000010: Create exam_results table**
- Create exam_results table with indexes
- Rollback: Drop exam_results table

#### Phase 6: Communication Module

**Migration 000011: Create notifications table**
- Create notifications table with indexes
- Rollback: Drop notifications table

**Migration 000012: Create announcements table**
- Create announcements table with indexes
- Rollback: Drop announcements table

**Migration 000013: Create messages table**
- Create messages table with indexes
- Rollback: Drop messages table

#### Phase 7: Audit & Logging Module

**Migration 000014: Create audit_logs table**
- Create audit_logs table with indexes
- Rollback: Drop audit_logs table

**Migration 000015: Create job_queue table**
- Create job_queue table with indexes
- Rollback: Drop job_queue table

### Migration Testing

**Pre-Migration Checklist:**
- [ ] Backup database
- [ ] Test migration on staging database
- [ ] Verify rollback on staging database
- [ ] Review migration SQL for errors

**Post-Migration Checklist:**
- [ ] Verify table creation
- [ ] Verify index creation
- [ ] Verify foreign key constraints
- [ ] Verify data integrity
- [ ] Run application smoke tests

### Rollback Strategy

**Automatic Rollback Conditions:**
- Migration fails mid-execution
- Post-migration verification fails
- Application errors detected

**Manual Rollback Procedure:**
1. Stop application
2. Run down migration
3. Verify rollback success
4. Restart application
5. Verify application functionality

---

## Data Migration

### No Data Migration Required

All new tables start empty. No existing data needs to be migrated.

### Data Seeding

**System Configurations:**
- Seed default system configurations
- Seed default roles and permissions
- Seed default subject categories
- Seed default graduate profile dimensions

---

## Performance Considerations

### Index Maintenance

- Run `ANALYZE` after index creation
- Run `VACUUM ANALYZE` weekly for high-churn tables
- Monitor index usage with `pg_stat_user_indexes`

### Query Optimization

- Use EXPLAIN ANALYZE for slow queries
- Add indexes based on query patterns
- Monitor slow query logs

### Partitioning Strategy

**Consider partitioning for:**
- audit_logs (by created_at, monthly)
- attendance_records (by date, monthly)
- notifications (by created_at, monthly)

**Partitioning benefits:**
- Faster query performance on recent data
- Easier data archival
- Better maintenance operations

---

## Backup Strategy

### Automated Backups

**Daily Full Backups:**
- pg_dump to S3
- Retention: 30 days

**Weekly Full Backups:**
- pg_dump to local storage
- Retention: 12 weeks

**Continuous WAL Archiving:**
- Point-in-time recovery capability
- Retention: 7 days

### Backup Verification

- Weekly restore test
- Backup integrity check
- Recovery time objective verification

---

## Security Considerations

### Row-Level Security

Consider adding RLS policies for multi-tenant isolation:
- School-based data isolation
- User-based data access

### Data Encryption

- At-rest encryption for sensitive fields (password_hash, personal data)
- TLS 1.3 for data in transit

### Audit Trail

- All changes tracked via audit_logs table
- Include user_id, action, entity_type, entity_id, old_values, new_values

---

## Conclusion

The target database design adds 15 new tables to support class management, attendance, scheduling, exams, assignments, and communication features. All existing tables are modified to support soft delete with partial indexes. Comprehensive indexing strategy includes composite indexes, partial indexes, and full-text search indexes for optimal query performance.

The migration strategy is incremental with rollback support, ensuring safe deployment. The design maintains consistency with existing patterns while addressing critical gaps identified in the audit.

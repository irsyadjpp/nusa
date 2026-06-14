# Sprint 5.5 — Class Management Implementation Plan

**Document Version**: 1.0  
**Date**: June 11, 2026  
**Status**: IMPLEMENTATION READY  
**Sprint Sequence**: Sprint 5.5 (Blocking Sprint - Pre-requisite for Sprint 5)  
**Architecture Alignment**: Architecture Freeze v2, DDD Lite, Layered Architecture  
**Priority**: CRITICAL - Blocks Sprint 5 Student Lifecycle Implementation

---

# Executive Summary

Sprint 5.5 implements the Class Management (Class/Rombel) foundation that was identified as a critical blocker for Sprint 5 (Student Lifecycle). This sprint establishes the Class/Rombel aggregate, MinIO storage integration, students table schema migration, and phase-to-class mapping capabilities required to enable Sprint 5 implementation.

**Scope**: Class Management and Storage Infrastructure only. Student lifecycle, attendance recording, document verification workflow, and class scheduling are explicitly out of scope.

**Dependencies**: Builds on Sprint 4 (Academic Foundation) and extends initial schema foundation.

---

# Table of Contents

1. [Gap Analysis](#gap-analysis)
2. [Business Requirements Document (BRD)](#business-requirements-document-brd)
3. [Functional Requirements](#functional-requirements)
4. [Domain Model](#domain-model)
5. [Database Design](#database-design)
6. [API Specifications](#api-specifications)
7. [Storage Integration](#storage-integration)
8. [Security Requirements](#security-requirements)
9. [Test Strategy](#test-strategy)
10. [Implementation Plan](#implementation-plan)

---

# Gap Analysis

## Gap Analysis Matrix

| Requirement | Existing Module | Reuse | Extend | New | Rationale |
| ----------- | --------------- | ----- | ------ | --- | --------- |
| **1. Class/Rombel Management** | ❌ None | | | ✅ New | No class structure exists - students table has class_name string only |
| **2. Wali Kelas Assignment** | ❌ None | | | ✅ New | No teacher-to-class assignment mechanism exists |
| **3. MinIO Storage** | ❌ None | | | ✅ New | No object storage integration exists for document/file handling |
| **4. Students Table Migration** | ⚠️ `students` table | | ✅ Extend | | students table needs class_id foreign key, not class_name string |
| **5. Phase-to-Class Mapping** | ⚠️ `curriculum_phases` | ✅ Reuse | | | Phase table exists, just needs class relationship |

## Summary Statistics

- **Total Requirements**: 5
- **Fully Covered**: 1 (20%) - Phase table exists
- **Extend Existing**: 1 (20%) - Students table
- **New Required**: 3 (60%) - Class aggregate, MinIO, Wali Kelas
- **New Tables Required**: 3 (classes, wali_kelas_assignments, wali_kelas_history)
- **Tables to Migrate**: 1 (students)
- **New Bounded Contexts Required**: 1 (Class Management)
- **Storage Infrastructure Required**: 1 (MinIO)

---

# Business Requirements Document (BRD)

## Business Objectives

### Primary Objective
Establish the Class Management infrastructure required to support Student Lifecycle, Attendance Recording, Assessment, and Reporting capabilities in the NUSA Platform.

### Secondary Objectives

1. **Class/Rombel Management**: Implement proper class (rombel) structure with unique identifiers, capacity management, and metadata.
2. **Teacher Assignment**: Enable assignment of Wali Kelas (class teachers) to classes with role-based responsibilities.
3. **Storage Infrastructure**: Implement MinIO object storage integration for document and file management.
4. **Schema Migration**: Migrate students table from string-based class names to proper foreign key relationships.
5. **Phase Integration**: Establish proper relationship between classes and Kurikulum Merdeka phases.

## Success Metrics

### Quantitative Metrics

| Metric | Target | Measurement Method |
| ------ | ------ | ------------------ |
| Class Creation Coverage | 100% of schools have classes configured | Database query count of classes per school |
| Wali Kelas Assignment | 100% of active classes have assigned Wali Kelas | Database query count of classes with wali kelas |
| Students Table Migration | 100% of student records have valid class_id references | Database query validation |
| MinIO Integration | MinIO service running and accessible | Health check ping test |
| Storage Performance | <2s file upload, <1s file retrieval | Performance test |
| Zero Downtime Migration | <30 seconds for schema migration | Migration execution time |
| Data Integrity | 100% referential integrity after migration | Foreign key validation query |

### Qualitative Metrics

- School admins can independently manage class structure without IT support
- Wali Kelas can view their assigned classes and students
- Teachers can understand class-to-phase relationships
- Document upload/retrieval works seamlessly through MinIO
- Class-based authorization works correctly
- Audit trail exists for all class management changes

## Scope

### In Scope

1. **Class/Rombel Management**
   - Create, read, update, deactivate classes
   - Set class metadata (grade level, phase, capacity, room)
   - Manage class status (ACTIVE, INACTIVE, ARCHIVED)
   - Class uniqueness validation (school + grade + name)

2. **Wali Kelas Assignment**
   - Assign Wali Kelas (class teachers) to classes
   - Support multiple Wali Kelas per class (primary, secondary)
   - Track Wali Kelas assignment history
   - Wali Kelas removal and reassignment

3. **MinIO Storage Integration**
   - Configure MinIO connection and bucket management
   - Implement file upload/download pre-signed URLs
   - Document versioning and metadata
   - Access control integration

4. **Students Table Migration**
   - Add class_id foreign key column to students table
   - Create classes table and populate with existing class_name data
   - Migrate data from class_name strings to class_id references
   - Remove class_name column after successful migration
   - Add indexes and constraints

5. **Phase-to-Class Mapping**
   - Add phase_id foreign key to classes table
   - Enable class querying by phase
   - Support phase-based class organization
   - Validate grade-level-to-phase rules

### Out of Scope

1. **Student Lifecycle** - Student enrollment, status management (Sprint 5)
2. **Class Scheduling** - Timetable, teacher assignment to time slots (future sprint)
3. **Attendance Recording** - Daily attendance tracking (Sprint 5)
4. **Class Capacity Planning** - Capacity planning algorithm (future sprint)
5. **Document Verification Workflow** - Manual verification of student documents (Sprint 5)
6. **Dapodik Integration** - External system synchronization (future sprint)
7. **Class Grouping** - Multiple class sections per grade (future enhancement)

## Stakeholders

### Primary Stakeholders

| Stakeholder | Role | Responsibilities |
| ----------- | ---- | --------------- |
| School Admin | Configure classes, assign Wali Kelas | Class structure management |
| Wali Kelas | View assigned classes and students | Class management context |
| System Admin | Configure MinIO, storage management | Platform infrastructure |
| Curriculum Admin | Ensure class-phase alignment | Curriculum compliance |

### Secondary Stakeholders

| Stakeholder | Role | Responsibilities |
| ----------- | ---- | --------------- |
| Principal | Review class configuration, approve changes | Oversight and approval |
| Teachers | View class information (read-only) | Information access |

## Business Processes

### Process 1: Class Configuration

**Actors**: School Admin

**Flow**:
1. School Admin creates new class for academic year
2. School Admin sets class metadata (grade level, phase, capacity, room)
3. School Admin assigns Wali Kelas (primary and secondary)
4. System validates class uniqueness (school + grade + name)
5. Class is created and marked as ACTIVE
6. Class is available for student assignment

**Business Rules**:
- Class name must be unique within school + grade level
- Class capacity cannot be negative
- Each class must have at least one Wali Kelas
- Only one primary Wali Kelas per class
- Class grade level must map to valid phase

### Process 2: Students Table Migration

**Actors**: System Admin (executed via migration script)

**Flow**:
1. Migration script creates classes table
2. Migration script populates classes from existing class_name data
3. Migration script adds class_id column to students table (nullable initially)
4. Migration script maps class_name strings to class_id references
5. Migration script validates all mappings
6. Migration script sets class_id NOT NULL constraint
7. Migration script removes class_name column
8. Indexes are created for performance

**Business Rules**:
- Migration must be reversible (down migration)
- No student data loss during migration
- Class name collisions must be handled (create unique classes)
- Backwards compatibility maintained until migration complete

### Process 3: MinIO Configuration

**Actors**: System Admin

**Flow**:
1. System Admin configures MinIO connection parameters
2. System Admin creates required buckets (documents, profile photos)
3. System Admin tests connectivity and access permissions
4. MinIO service is marked as healthy and operational
5. Application can store and retrieve files

**Business Rules**:
- MinIO must be accessible from application server
- Bucket names must be unique and follow naming convention
- File size limits must be enforced at application level
- Access policies must be configured for security

---

# Functional Requirements

## FR-CL-001: Class Creation

The system shall allow School Admin to create classes with the following information:
- School reference
- Class name (e.g., "1A", "2B", "X IPA 1")
- Grade level (e.g., "1", "2", "X", "XI", "XII")
- Phase reference (e.g., "A", "B", "C")
- Room number (optional)
- Capacity (optional)
- Academic year reference
- Status (ACTIVE, INACTIVE, ARCHIVED)

## FR-CL-002: Class Update

The system shall allow School Admin to update class information for:
- Room number
- Capacity
- Status

Updates to class name, grade level, or phase shall require special validation and audit logging.

## FR-CL-003: Class Deactivation

The system shall support soft archival of class records when:
- Academic year ends
- Class is dissolved
- Class is reorganized

Archived records shall be retained for audit purposes but not displayed in active class lists.

## FR-CL-004: Class Uniqueness

The system shall enforce class uniqueness at school level:
- Class name must be unique within school + grade level combination
- Example: "1A" can exist for Grade 1 but not multiple times for Grade 1

## FR-CL-005: Phase Validation

The system shall validate that class grade level maps to correct phase:
- Grades 1-2 → Fase A
- Grades 3-4 → Fase B  
- Grades 5-6 → Fase C

Class creation with invalid grade-phase mapping shall be rejected.

## FR-WK-001: Wali Kelas Assignment

The system shall allow School Admin to assign Wali Kelas to classes with:
- Teacher reference
- Assignment type (PRIMARY, SECONDARY)
- Assignment date

## FR-WK-002: Multiple Wali Kelas

The system shall support multiple Wali Kelas per class with:
- One primary Wali Kelas designation
- Multiple secondary Wali Kelas designations
- Priority ordering

## FR-WK-003: Wali Kelas History

The system shall maintain history of Wali Kelas assignments for audit purposes.

## FR-WK-004: Wali Kelas Removal

The system shall allow removal of Wali Kelas assignments with:
- Automatic reassignment of students if needed
- Assignment termination date
- Audit trail

## FR-ST-MIG-001: Students Table Migration

The system shall migrate the students table schema with:
- Addition of class_id foreign key column (nullable initially)
- Data migration from class_name strings to class_id references
- Removal of class_name column after successful migration
- Index creation for performance optimization

## FR-ST-MIG-002: Class Creation from Existing Data

The system shall automatically create classes from existing class_name data during migration:
- Parse class_name strings (e.g., "1A", "2B", "X IPA 1")
- Extract grade level from class name
- Create classes with default metadata
- Map students to created classes

## FR-ST-MIG-003: Migration Rollback

The system shall support migration rollback with:
- Down migration script to restore class_name column
- Removal of class_id column
- Deletion of created classes if needed

## FR-IO-001: MinIO Configuration

The system shall allow System Admin to configure MinIO with:
- Connection parameters (endpoint, access key, secret key)
- Bucket creation and configuration
- Health check and connectivity test

## FR-IO-002: File Upload

The system shall allow file upload to MinIO with:
- Bucket selection
- File type validation
- File size validation
- Metadata attachment
- Version management

## FR-IO-003: File Retrieval

The system shall allow file retrieval from MinIO with:
- Pre-signed URL generation
- Access permission validation
- File metadata retrieval

## FR-IO-004: Storage Health Monitoring

The system shall monitor MinIO storage health with:
- Periodic connectivity checks
- Bucket availability validation
- Storage capacity monitoring
- Alert generation for issues

---

# Business Rules

## BR-CL-001: Class Name Format

Class name must follow Indonesian naming convention:
- Primary: Grade + Class letter (e.g., "1A", "2B")
- Secondary: Grade + Program + Class letter (e.g., "X IPA 1", "XI MIPA 1")
- Max length: 50 characters
- Alphanumeric + spaces only

## BR-CL-002: Capacity Validation

Class capacity must be:
- Greater than 0 if specified
- Not negative
- Maximum 60 students per class (Indonesian regulation)

## BR-CL-003: Phase-Grade Mapping

Grade level must map to correct Kurikulum Merdeka phase:
- Grades 1-2 → Fase A
- Grades 3-4 → Fase B
- Grades 5-6 → Fase C

Class creation with invalid mapping must be rejected.

## BR-CL-004: Class Uniqueness

Class name must be unique within school + grade level:
- Example: "1A" can exist once for Grade 1
- "1A" can exist for Grade 2 (different grade level)

## BR-CL-005: Status Transition Rules

- ACTIVE can transition to INACTIVE, ARCHIVED
- INACTIVE can transition to ACTIVE, ARCHIVED
- ARCHIVED is terminal (no reactivation)

## BR-WK-001: Primary Wali Kelas Uniqueness

Only one primary Wali Kelas per class at any given time.

## BR-WK-002: Wali Kelas Teacher Validation

Wali Kelas must be an active teacher with proper authorization.

## BR-MIG-001: Migration Safety

Migration must:
- Be reversible with down migration script
- Preserve all student data
- Allow zero-downtime if possible
- Provide rollback capability

## BR-MIG-002: Class Name Parsing

Class name parsing must handle:
- Standard format: Grade + Letter (e.g., "1A", "2B")
- Extended format: Grade + Program + Letter (e.g., "X IPA 1")
- Format variations (spaces, no spaces)
- Handle invalid formats gracefully (create default class)

## BR-IO-001: File Size Limits

File size must not exceed 10 MB for document uploads.

## BR-IO-002: File Type Validation

Only allowed file types:
- Documents: PDF, DOC, DOCX
- Images: JPEG, PNG, WebP
- Others: Based on configuration

## BR-IO-003: Bucket Naming

Bucket names must:
- Be lowercase
- Use hyphens only
- Start with letter or number
- Be globally unique within MinIO instance

---

# Domain Model

## Bounded Context: Class Management

### Aggregate: Class

**Aggregate Root**: `Class`

**Purpose**: Class/Rombel (study group) management with teacher assignment

**Entities**:
- `Class` (root)
- `WaliKelasAssignment` (child)
- `WaliKelasAssignmentHistory` (child)

**Value Objects**:
- `ClassID` (UUID)
- `ClassName` (string, e.g., "1A", "X IPA 1")
- `GradeLevel` (string, e.g., "1", "2", "X", "XI", "XII")
- `PhaseID` (UUID reference to curriculum_phases)
- `ClassStatus` (ACTIVE, INACTIVE, ARCHIVED)
- `AssignmentType` (PRIMARY, SECONDARY)
- `RoomNumber` (optional string)
- `Capacity` (optional integer)

**Domain Services**:
- `ClassCreationService` (validates class creation rules)
- `ClassMigrationService` (migrates students table from class_name to class_id)
- `WaliKelasAssignmentService` (manages teacher assignments)

**Invariants**:
- CL-INV-001: Class name must be unique within school + grade level
- CL-INV-002: Class grade level must map to correct phase
- CL-INV-003: Class must have at least one Wali Kelas when ACTIVE
- CL-INV-004: Class capacity must be positive if specified
- CL-INV-005: Only one primary Wali Kelas per class

**Lifecycle Rules**:
```
ACTIVE → INACTIVE → ACTIVE
ACTIVE → ARCHIVED (terminal)
INACTIVE → ARCHIVED (terminal)
```

**Ownership Rules**:
- School Admin owns classes within their school
- System Admin has read-only access across all schools
- Wali Kelas have read-only access to their assigned classes

---

# Database Design

## ERD Overview

```
classes (1) ── (N) wali_kelas_assignments
classes (N) ── (1) schools
classes (N) ── (1) curriculum_phases
classes (N) ── (1) academic_years
students (N) ── (1) classes (after migration)
users (N) ── (N) wali_kelas_assignments
wali_kelas_assignments (1) ── (N) wali_kelas_assignment_history
```

## Table Definitions

### Table: classes (NEW)

| Column | Type | Nullable | Default | Constraints |
|--------|------|----------|---------|-------------|
| `id` | UUID | NO | `gen_uuid_v7()` | PRIMARY KEY |
| `school_id` | UUID | NO | — | FK → `schools(id)` ON DELETE RESTRICT |
| `academic_year_id` | UUID | NO | — | FK → `academic_years(id)` ON DELETE RESTRICT |
| `phase_id` | UUID | YES | — | FK → `curriculum_phases(id)` ON DELETE SET NULL |
| `name` | VARCHAR(50) | NO | — | — |
| `grade_level` | VARCHAR(10) | NO | — | CHECK (grade_level IN ('1','2','3','4','5','6','7','8','9','10','11','12','X','XI','XII')) |
| `room_number` | VARCHAR(20) | YES | — | — |
| `capacity` | INTEGER | YES | — | CHECK (capacity > 0 AND capacity <= 60) |
| `status` | VARCHAR(20) | NO | 'ACTIVE' | CHECK (status IN ('ACTIVE','INACTIVE','ARCHIVED')) |
| `created_by` | UUID | NO | — | FK → `users(id)` ON DELETE RESTRICT |
| `updated_by` | UUID | YES | — | FK → `users(id)` ON DELETE SET NULL |
| `created_at` | TIMESTAMPTZ | NO | `NOW()` | — |
| `updated_at` | TIMESTAMPTZ | NO | `NOW()` | — |

**Indexes:**
- `idx_classes_school_id`
- `idx_classes_academic_year_id`
- `idx_classes_phase_id`
- `idx_classes_grade_level`
- `idx_classes_status`
- `idx_classes_school_grade_name` (unique: school_id, grade_level, name)

**Constraints:**
- UNIQUE: (school_id, grade_level, name)
- CHECK: Class grade level maps to phase (application-level)

### Table: wali_kelas_assignments (NEW)

| Column | Type | Nullable | Default | Constraints |
|--------|------|----------|---------|-------------|
| `id` | UUID | NO | `gen_uuid_v7()` | PRIMARY KEY |
| `class_id` | UUID | NO | — | FK → `classes(id)` ON DELETE CASCADE |
| `user_id` | UUID | NO | — | FK → `users(id)` ON DELETE CASCADE |
| `assignment_type` | VARCHAR(20) | NO | 'SECONDARY' | CHECK (assignment_type IN ('PRIMARY','SECONDARY')) |
| `assigned_date` | TIMESTAMPTZ | NO | `NOW()` | — |
| `terminated_date` | TIMESTAMPTZ | YES | — | — |
| `is_active` | BOOLEAN | NO | `true` | — |
| `created_by` | UUID | NO | — | FK → `users(id)` ON DELETE RESTRICT |
| `created_at` | TIMESTAMPTZ | NO | `NOW()` | — |

**Indexes:**
- `idx_wali_kelas_assignments_class_id`
- `idx_wali_kelas_assignments_user_id`
- `idx_wali_kelas_assignments_is_active`
- `idx_wali_kelas_assignments_class_type_active` (class_id, assignment_type, is_active)

**Constraints:**
- CHECK: Only one PRIMARY assignment per class (application-level)
- CHECK: User must be TEACHER or SCHOOL_ADMIN role (application-level)

### Table: wali_kelas_assignment_history (NEW)

| Column | Type | Nullable | Default | Constraints |
|--------|------|----------|---------|-------------|
| `id` | UUID | NO | `gen_uuid_v7()` | PRIMARY KEY |
| `class_id` | UUID | NO | — | FK → `classes(id)` ON DELETE CASCADE |
| `user_id` | UUID | NO | — | FK → `users(id)` ON DELETE CASCADE |
| `assignment_type` | VARCHAR(20) | NO | — | CHECK (assignment_type IN ('PRIMARY','SECONDARY')) |
| `assigned_date` | TIMESTAMPTZ | NO | — | — |
| `terminated_date` | TIMESTAMPTZ | YES | — | — |
| `changed_by` | UUID | NO | — | FK → `users(id)` ON DELETE RESTRICT |
| `changed_at` | TIMESTAMPTZ | NO | `NOW()` | — |

**Indexes:**
- `idx_wali_kelas_history_class_id`
- `idx_wali_kelas_history_user_id`

### Table: students (MODIFY)

**Add Column:**
```sql
ALTER TABLE students ADD COLUMN class_id UUID REFERENCES classes(id) ON DELETE SET NULL;
CREATE INDEX idx_students_class_id ON students(class_id);
```

**Remove Column (after migration):**
```sql
ALTER TABLE students DROP COLUMN class_name;
```

## Storage Integration Design

### MinIO Configuration

| Setting | Environment Variable | Required | Example |
|----------|---------------------|----------|---------|
| MinIO Endpoint | `MINIO_ENDPOINT` | Yes | `http://localhost:9000` |
| MinIO Access Key | `MINIO_ACCESS_KEY` | Yes | `minioadmin` |
| MinIO Secret Key | `MINIO_SECRET_KEY` | Yes | `minioadmin` |
| MinIO SSL | `MINIO_SSL` | No | `false` |
| MinIO Bucket Prefix | `MINIO_BUCKET_PREFIX` | No | `nusa` |

### Buckets

| Bucket Name | Purpose | Public | Versioning |
|------------|---------|--------|-------------|
| `nusa-documents` | Student documents (KK, birth cert, etc.) | No | Yes |
| `nusa-photos` | Student profile photos | No | Yes |
| `nusa-resources` | Learning resources | Yes | No |

---

# API Specifications

## Class Management Endpoints

### POST /api/v1/classes

**Description**: Create a new class

**Request Body**:
```json
{
  "school_id": "uuid",
  "academic_year_id": "uuid",
  "phase_id": "uuid",
  "name": "1A",
  "grade_level": "1",
  "room_number": "101",
  "capacity": 30
}
```

**Response**: 201 Created with Class entity

### GET /api/v1/classes

**Description**: Get list of classes with optional filters

**Query Parameters**:
- `school_id` (optional) - Filter by school
- `academic_year_id` (optional) - Filter by academic year
- `phase_id` (optional) - Filter by phase
- `grade_level` (optional) - Filter by grade level
- `status` (optional) - Filter by status

**Response**: 200 OK with array of Class entities

### GET /api/v1/classes/{id}

**Description**: Get class by ID

**Response**: 200 OK with Class entity

### PUT /api/v1/classes/{id}

**Description**: Update class

**Request Body**:
```json
{
  "room_number": "102",
  "capacity": 32,
  "status": "INACTIVE"
}
```

**Response**: 200 OK with updated Class entity

### DELETE /api/v1/classes/{id}

**Description**: Deactivate class (soft delete)

**Response**: 204 No Content

## Wali Kelas Assignment Endpoints

### POST /api/v1/classes/{class_id}/wali-kelas

**Description**: Assign Wali Kelas to class

**Request Body**:
```json
{
  "user_id": "uuid",
  "assignment_type": "PRIMARY"
}
```

**Response**: 201 Created

### GET /api/v1/classes/{class_id}/wali-kelas

**Description**: Get active Wali Kelas for class

**Response**: 200 OK with array of assignments

### DELETE /api/v1/classes/{class_id}/wali-kelas/{user_id}

**Description**: Remove Wali Kelas assignment

**Response**: 204 No Content

## Migration Endpoints

### POST /api/v1/admin/migrate-students-to-classes

**Description**: Trigger students table migration

**Request Body**:
```json
{
  "dry_run": false
}
```

**Response**: 200 OK with migration results

### GET /api/v1/admin/migration-status

**Description**: Check migration status

**Response**: 200 OK with status information

## Storage Endpoints

### POST /api/v1/storage/upload

**Description**: Upload file to MinIO

**Request**: Multipart form data with file

**Response**: 200 OK with file object key and pre-signed URL

### GET /api/v1/storage/download/{object_key}

**Description**: Get pre-signed download URL for file

**Response**: 200 OK with download URL

---

# Security Requirements

## Authorization Matrix

| Resource | Create | Read | Update | Delete | School Admin | Wali Kelas | System Admin |
|----------|--------| ----- | ------ | ------ | -------------| ------------| -------------|
| Classes | ✅ | ✅ | ✅ | ✅ | Own school | Own classes | All schools |
| Wali Kelas Assignment | ✅ | ✅ | ✅ | ✅ | Own school | Own assignments | All |
| Storage Upload | ✅ | ✅ | ❌ | ❌ | Own school | No | All |
| Migration Trigger | ❌ | ✅ | ❌ | ❌ | No | No | Yes |

---

# Test Strategy

## Unit Tests

### Domain Tests
- Class creation with valid data
- Class creation with invalid phase mapping
- Class creation with duplicate name
- Class status transitions
- Wali Kelas assignment rules
- Primary Wali Kelas uniqueness

### Repository Tests
- Class CRUD operations
- Wali Kelas CRUD operations
- Foreign key relationships
- Index performance

### Migration Tests
- Students table migration (class_name → class_id)
- Rollback migration
- Data integrity validation
- Performance with large datasets

## Integration Tests

### MinIO Integration
- File upload and retrieval
- Pre-signed URL generation
- Access control validation
- Storage health monitoring

### API Tests
- Class management endpoints
- Wali Kelas assignment endpoints
- Migration endpoints
- Storage endpoints

## Performance Tests

- Class querying by school with 1000+ classes
- Student class_id index performance
- File upload/download performance
- Migration performance with 10,000+ student records

---

# Implementation Plan

## Phase 1: Storage Infrastructure (Days 1-2)

**Objective**: Establish MinIO integration

**Tasks**:
1. Install and configure MinIO
2. Create MinIO buckets
3. Implement storage service layer
4. Add MinIO configuration to environment
5. Create storage repository
6. Implement file upload/download handlers
7. Add storage health check endpoint
8. Write storage integration tests

**Deliverables**:
- MinIO configuration
- Storage service and repository
- File upload/download handlers
- Storage tests

## Phase 2: Class Aggregate Implementation (Days 3-5)

**Objective**: Implement Class/Rombel domain and database

**Tasks**:
1. Create classes table migration
2. Create wali_kelas_assignments table migration
3. Create wali_kelas_assignment_history table migration
4. Implement Class domain model
5. Implement WaliKelasAssignment domain model
6. Implement Class repository
7. Implement WaliKelasAssignment repository
8. Implement Class application service
9. Implement WaliKelasAssignment application service
10. Add CURRICULUM_ADMIN permissions for class management
11. Write domain unit tests
12. Write repository tests

**Deliverables**:
- Database migration files
- Domain models (Class, WaliKelasAssignment)
- Repositories
- Application services
- Domain and repository tests

## Phase 3: API Layer (Days 6-7)

**Objective**: Implement REST API endpoints

**Tasks**:
1. Create Class DTOs
2. Create Class handlers
3. Create Wali Kelas handlers
4. Add class management routes to router
5. Add Wali Kelas assignment routes
6. Implement request validation
7. Add error handling
8. Write API integration tests

**Deliverables**:
- DTOs
- HTTP handlers
- Route registration
- API tests

## Phase 4: Students Table Migration (Days 8-9)

**Objective**: Migrate students table from class_name to class_id

**Tasks**:
1. Add class_id column to students table (nullable)
2. Create class_id index
3. Implement class_name parsing logic
4. Create migration data script
5. Map existing class_name values to new classes
6. Validate all students have class_id
7. Set class_id NOT NULL constraint
8. Remove class_name column
9. Create down migration script
10. Test migration on development database
11. Test migration rollback
12. Run migration on production

**Deliverables**:
- Migration scripts (up/down)
- Migration data processing logic
- Migration tests

## Phase 5: Frontend Class Management (Days 10-12)

**Objective**: Create React components for class management

**Tasks**:
1. Create Class API client
2. Create Class query and command services
3. Create ClassManagement component
4. Create WaliKelasAssignment component
5. Add class management to menu
6. Create class management pages
7. Implement form validation
8. Add loading/error states
9. Write frontend tests

**Deliverables**:
- API client
- Query/command services
- React components
- Menu integration
- Frontend tests

## Phase 6: Testing and Documentation (Days 13-14)

**Objective**: Comprehensive testing and documentation

**Tasks**:
1. Run all test suites
2. Fix any failing tests
3. Update API documentation
4. Create migration documentation
5. Create MinIO setup guide
6. Update CHANGELOG
7. Create sprint completion report
8. Prepare Sprint 5 start checklist

**Deliverables**:
- Passing test suite
- Updated documentation
- Migration guide
- Sprint completion report

## Risk Mitigation

### Risk 1: Migration Failure
- **Mitigation**: Comprehensive testing on development database, rollback script available
- **Contingency**: Manual data recovery from backup if needed

### Risk 2: Class Name Parsing Errors
- **Mitigation**: Handle invalid formats gracefully, create default classes, manual review of unmapped records
- **Contingency**: Manual data cleanup script

### Risk 3: MinIO Configuration Issues
- **Mitigation**: Detailed setup guide, configuration validation, health checks
- **Contingency**: Use temporary local storage or defer document management

### Risk 4: Schedule Overrun
- **Mitigation**: Defer nice-to-have features, prioritize critical path
- **Contingency**: Phase out storage validation, simplify migration logic

## Acceptance Criteria

- [ ] All class management CRUD operations working
- [ ] Wali Kelas assignment and removal working
- [ ] Students table successfully migrated (class_name → class_id)
- [ ] MinIO storage integration working
- [ ] All tests passing (unit, integration, performance)
- [ ] Documentation updated
- [ ] Sprint 5 dependencies satisfied
- [ ] Zero data loss during migration
- [ ] Migration successfully rolled back and re-applied

---

# Timeline Summary

**Total Duration**: 14 days

| Phase | Days | Status |
|-------|------|--------|
| Phase 1: Storage Infrastructure | 2 | Pending |
| Phase 2: Class Aggregate Implementation | 3 | Pending |
| Phase 3: API Layer | 2 | Pending |
| Phase 4: Students Table Migration | 2 | Pending |
| Phase 5: Frontend Class Management | 3 | Pending |
| Phase 6: Testing and Documentation | 2 | Pending |

---

# Dependencies on Sprint 4

This sprint requires Sprint 4 to be fully completed:

- ✅ academic_years table exists
- ✅ semesters table exists
- ✅ curriculum_phases table exists (from migration 000002)
- ✅ schools table exists (from initial schema)
- ✅ users table exists (from initial schema)
- ✅ CURRICULUM_ADMIN role exists

All Sprint 4 dependencies are satisfied.

---

# Success Criteria for Sprint 5.5

Upon completion of Sprint 5.5, Sprint 5 (Student Lifecycle) will be unblocked because:

1. **Class/Rombel Aggregate**: Will exist for StudentClassAssignment, StudentPromotion, StudentAttendance
2. **Foreign Key Relationships**: Students table will have proper class_id references
3. **Phase-to-Class Mapping**: Classes will have phase_id for promotion validation
4. **MinIO Storage**: Will be available for StudentDocument file storage
5. **Referential Integrity**: All foreign key relationships will be properly established
6. **Audit Trail**: Wali Kelas history will provide teacher assignment tracking
7. **Authorization**: Class-based access control will be possible

**Sprint 5 can proceed safely after Sprint 5.5 completion.**
# Sprint 4 Requirement Package: Academic Foundation

**Document Version**: 1.0  
**Date**: 2026-06-11  
**Status**: Implementation Ready  
**Authors**: Principal Solution Architect, Product Owner, Education Domain Expert, DDD Reviewer, Technical Analyst

---

# Table of Contents

1. [PART 1 – GAP ANALYSIS](#part-1--gap-analysis)
2. [PART 2 – BUSINESS REQUIREMENTS DOCUMENT (BRD)](#part-2--business-requirements-document-brd)
3. [PART 3 – FUNCTIONAL REQUIREMENTS](#part-3--functional-requirements)
4. [PART 4 – DOMAIN MODEL](#part-4--domain-model)
5. [PART 5 – DATABASE DESIGN](#part-5--database-design)
6. [PART 6 – API DESIGN](#part-6--api-design)
7. [PART 7 – FRONTEND REQUIREMENTS](#part-7--frontend-requirements)
8. [PART 8 – SECURITY REQUIREMENTS](#part-8--security-requirements)
9. [PART 9 – TEST STRATEGY](#part-9--test-strategy)
10. [PART 10 – IMPLEMENTATION PLAN](#part-10--implementation-plan)

---

# PART 1 – GAP ANALYSIS

## Gap Analysis Matrix

| Requirement | Existing Module | Reuse | Extend | New | Rationale |
| ----------- | --------------- | ----- | ------ | --- | --------- |
| **1. Academic Year** | ❌ None | | | ✅ New | No temporal scoping exists in current curriculum structure |
| **2. Semester** | ❌ None | | | ✅ New | No semester structure exists; required for academic year division |
| **3. Subject Category** | ⚠️ `curriculum_subjects` | | ✅ Extend | | Subject table exists but lacks categorization (Intrakurikuler/Kokurikuler/Ekstrakurikuler) |
| **4. Profil Lulusan 8 Dimensi** | ❌ None | | | ✅ New | No graduate profile dimension tracking exists |
| **5. Kurikulum 2026 Alignment** | ✅ `cp`, `tp`, `atp`, `modul_ajar` | ✅ Reuse | | | Core curriculum modules are already Kurikulum Merdeka compliant |
| **6. Koding** | ✅ `curriculum_subjects` | ✅ Reuse | | | Can add "Koding" as new subject record |
| **7. AI** | ✅ `curriculum_subjects` | ✅ Reuse | | | Can add "AI" as new subject record |
| **8. Numerasi** | ✅ `curriculum_subjects` | ✅ Reuse | | | Can add "Numerasi" as new subject record or integrate with existing subjects |
| **9. Academic Structure Governance** | ✅ `users`, `roles`, `permissions` | ✅ Reuse | ✅ Extend | | Existing auth/authorization can be extended with new permissions |

## Summary Statistics

- **Total Requirements**: 9
- **Fully Covered**: 3 (33%)
- **Extend Existing**: 2 (22%)
- **New Required**: 4 (44%)
- **New Tables Required**: 4 (academic_years, semesters, subject_categories, graduate_profile_dimensions)
- **Tables to Extend**: 1 (curriculum_subjects)
- **New Bounded Contexts Required**: 0 (All extensions fit within existing Curriculum Context)

---

# PART 2 – BUSINESS REQUIREMENTS DOCUMENT (BRD)

## Business Objectives

### Primary Objective
Establish the academic foundation infrastructure required to support Kurikulum 2026 implementation and Deep Learning pedagogy, enabling the platform to transition from curriculum content management to comprehensive academic governance.

### Secondary Objectives

1. **Temporal Scoping**: Implement academic year and semester management to enable curriculum planning, delivery, and reporting within defined time boundaries.
2. **Subject Categorization**: Categorize subjects according to Kurikulum Merdeka components (Intrakurikuler, Kokurikuler, Ekstrakurikuler) to support comprehensive curriculum delivery.
3. **Graduate Profile Tracking**: Implement the 8-dimensional Profil Lulusan framework to enable curriculum alignment with national graduate outcomes.
4. **Modern Subject Support**: Enable support for Koding, AI, and Numerasi subjects as required by Kurikulum 2026.
5. **Governance Framework**: Establish academic structure governance to enable curriculum administrators to manage academic configuration without developer intervention.

## Success Metrics

### Quantitative Metrics

| Metric | Target | Measurement Method |
| ------ | ------ | ------------------ |
| Academic Year Coverage | 100% of schools have active academic year configured | Database query count of active academic_years |
| Semester Configuration | 100% of academic years have 2 semesters configured | Database query count of semesters per academic year |
| Subject Categorization | 100% of subjects assigned to category | Database query of subjects with category_id |
| Profil Lulusan Dimensions | 8 dimensions configured and active | Database query count of active graduate_profile_dimensions |
| Kurikulum Alignment | 100% of CP aligned to graduate profile dimensions | Database query of CP with dimension assignments |
| Configuration Time | <5 minutes to configure academic year | Time measurement test |
| Zero Downtime Migration | <30 seconds for schema migration | Migration execution time |

### Qualitative Metrics

- School admins can independently configure academic structure without IT support
- Curriculum administrators can verify Kurikulum 2026 alignment via UI
- Teachers can understand subject categorization in curriculum planning
- Audit trail exists for all academic structure changes
- Configuration changes are reversible with rollback capability

## Scope

### In Scope

1. **Academic Year Management**
   - Create, read, update, deactivate academic years
   - Set academic year date ranges
   - Mark current/active academic year
   - Academic year approval workflow

2. **Semester Management**
   - Create semesters within academic years
   - Set semester date ranges
   - Configure semester sequence (Ganjil/Genap)
   - Semester activation/deactivation

3. **Subject Category Management**
   - Define Intrakurikuler, Kokurikuler, Ekstrakurikuler categories
   - Categorize existing and new subjects
   - Manage category metadata (description, guidelines)
   - Category governance (create, update, deactivate)

4. **Profil Lulusan Management**
   - Configure 8 graduate profile dimensions:
     1. Keimanan & Ketakwaan
     2. Kewargaan
     3. Berakhlak Mulia
     4. Berani Bertanggung Jawab
     5. Peduli
     6. Gotong Royong
     7. Mandiri
     8. Kreatif
   - Define dimension descriptions and indicators
   - Link dimensions to CP for alignment tracking

5. **Kurikulum Alignment Management**
   - Link CP to graduate profile dimensions
   - Track alignment percentage per CP
   - Generate alignment reports
   - Identify gaps in curriculum coverage

6. **Koding & AI Subject Support**
   - Add Koding subject to curriculum_subjects
   - Add AI subject to curriculum_subjects
   - Add Numerasi subject to curriculum_subjects
   - Categorize as Intrakurikuler

### Out of Scope

1. **Student Management** - Student lifecycle, enrollment, records
2. **PPDB** - New student admission workflow
3. **Class Management** - Rombel, wali kelas, scheduling
4. **Attendance** - Student attendance tracking
5. **Scheduling** - Class timetable, teacher assignments
6. **Grading** - Grade calculation, report cards
7. **Dapodik Integration** - External system synchronization
8. **Deep Learning Pedagogy** - Teaching methodology implementation
9. **Assessment Alignment** - Linking assessments to graduate profiles (future sprint)

## Stakeholders

### Primary Stakeholders

| Stakeholder | Role | Responsibilities |
| ----------- | ---- | --------------- |
| School Admin | Configure academic year/semester for school | School-level academic configuration |
| Curriculum Admin | Manage subject categories, Profil Lulusan, alignment | System-wide curriculum governance |
| Teacher | View academic structure, understand subject categorization | Curriculum planning context |
| System Admin | Platform-level configuration, permissions | System administration |

### Secondary Stakeholders

| Stakeholder | Role | Responsibilities |
| ----------- | ---- | --------------- |
| Principal | Review academic configuration, approve changes | Oversight and approval |
- | Ministry of Education | Compliance with Kurikulum 2026 | Regulatory compliance |
| Parents | View academic calendar (read-only) | Information access |

## Business Processes

### Process 1: Academic Year Configuration

**Actors**: School Admin, System Admin

**Flow**:
1. School Admin requests new academic year creation
2. System Admin approves academic year creation request
3. School Admin configures academic year date range
4. School Admin defines semesters within academic year
5. School Admin submits for approval
6. System Admin approves academic year configuration
7. Academic year becomes active on specified date

**Business Rules**:
- Only one academic year can be active at a time per school
- Academic year date ranges cannot overlap
- Academic years must be configured at least 30 days before start
- Past academic years cannot be modified (read-only)
- Active academic year cannot be deactivated without approval

### Process 2: Subject Categorization

**Actors**: Curriculum Admin

**Flow**:
1. Curriculum Admin creates subject category
2. Curriculum Admin defines category metadata
3. Curriculum Admin categorizes subjects
4. Curriculum Admin publishes categorization
5. Teachers view categorized subjects in curriculum planning

**Business Rules**:
- Each subject must belong to exactly one category
- Categories cannot be deleted if subjects are assigned
- Category changes require approval
- Historical categorization is preserved

### Process 3: Profil Lulusan Configuration

**Actors**: Curriculum Admin

**Flow**:
1. Curriculum Admin configures 8 graduate profile dimensions
2. Curriculum Admin defines dimension indicators
3. Curriculum Admin links CP to dimensions
4. System generates alignment reports
5. Curriculum Admin reviews gaps
6. Curriculum Admin adjusts CP-dimension links

**Business Rules**:
- CP must align to at least 1 graduate profile dimension
- CP can align to multiple dimensions
- Dimension configuration is system-wide (not per school)
- Alignment percentage is calculated automatically

## Business Rules

### BR-001: Academic Year Uniqueness
Each school can have only one academic year active at any given time.

### BR-002: Academic Year Non-Overlap
Academic year date ranges for a school cannot overlap with existing academic years.

### BR-003: Academic Year Lead Time
New academic years must be created at least 30 days before the start date.

### BR-004: Semester Sequence
Each academic year must have exactly 2 semesters in sequence: Ganjil (odd) followed by Genap (even).

### BR-005: Semester Date Coverage
Semester date ranges must fully cover the academic year date range without gaps or overlaps.

### BR-006: Subject Category Exclusivity
Each subject must belong to exactly one subject category (Intrakurikuler, Kokurikuler, or Ekstrakurikuler).

### BR-007: Profil Lulusan Completeness
CP must align to at least 1 of the 8 graduate profile dimensions.

### BR-008: Kurikulum Alignment Minimum
Each CP must have a minimum alignment score of 60% to graduate profile dimensions.

### BR-009: Modern Subject Support
Koding, AI, and Numerasi subjects must be categorized as Intrakurikuler.

### BR-010: Governance Approval
All academic structure changes (academic year, semester, category, dimension) require approval before becoming active.

## Assumptions

### Technical Assumptions

1. PostgreSQL 18+ with UUID v7 support is available
2. Existing authentication and authorization infrastructure will be reused
3. Existing audit logging infrastructure will be reused
4. Database migrations will follow existing pattern (up/down SQL files)
5. API will follow existing REST patterns and naming conventions

### Business Assumptions

1. Schools operate on a standard 2-semester academic year structure
2. Academic years follow calendar year structure (e.g., 2026/2027)
3. Subject categorization follows Kurikulum Merdeka standard definitions
4. Graduate profile dimensions are consistent across all schools (national standard)
5. School admins have authority to configure academic years for their schools
6. Curriculum admins have system-wide authority for categories and dimensions
7. Academic structure changes are infrequent (annually or semi-annually)

### Integration Assumptions

1. Dapodik integration will be deferred to future sprint
2. External calendar systems (Google Calendar, etc.) will not be integrated in Sprint 4
3. No real-time synchronization with external systems required

## Constraints

### Technical Constraints

1. **Architecture Compliance**: Must follow Architecture Freeze v2 (no CQRS, no Event Sourcing, no new bounded contexts without approval)
2. **Database Schema**: Must follow Database Schema Freeze v1 patterns (UUID primary keys, audit fields, soft delete via status)
3. **API Standards**: Must follow existing REST API patterns (versioned endpoints, consistent error codes)
4. **Frontend Stack**: Must use existing React/TypeScript stack with MUI components
5. **Performance**: API response time must be <500ms for read operations, <2s for write operations

### Business Constraints

1. **Kurikulum Compliance**: Must align with Kurikulum Merdeka 2026 requirements
2. **National Standards**: Graduate profile dimensions must match Permendikdasmen No. 10 Tahun 2025
3. **School Autonomy**: Schools must have autonomy to configure academic years within national guidelines
4. **Data Privacy**: No student personal data will be stored in Sprint 4 (deferred to future sprint)

### Resource Constraints

1. **Timeline**: Sprint 4 duration is 8 weeks
2. **Team**: Solo developer with limited time
3. **Budget**: No external dependencies or paid services
4. **Infrastructure**: Single PostgreSQL instance, no distributed systems

### Regulatory Constraints

1. **Indonesian Language**: All UI text must be in Indonesian
2. **Data Sovereignty**: All data must reside in Indonesia
3. **Education Standards**: Must comply with SNP (Standar Nasional Pendidikan)

---

# PART 3 – FUNCTIONAL REQUIREMENTS

## Feature 1: Academic Year Management

### Description
Enable school administrators to create, configure, and manage academic years with defined date ranges, approval workflows, and semester divisions.

### Actors
- School Admin (primary)
- System Admin (approval)
- Curriculum Admin (read access)
- Teacher (read access)

### Preconditions
- User must be authenticated
- User must have SCHOOL_ADMIN role (for create/update) or appropriate read permission
- School must exist and be active

### Main Flow

**1. Create Academic Year**
1. School Admin navigates to Academic Year Management page
2. School Admin clicks "Create New Academic Year"
3. System displays academic year creation form
4. School Admin enters:
   - Academic year name (e.g., "2026/2027")
   - Start date
   - End date
   - Description (optional)
5. System validates:
   - Start date < end date
   - No overlap with existing academic years for school
   - Start date is at least 30 days in future
6. System creates academic year in DRAFT status
7. System creates audit log entry
8. System returns success response with academic year ID

**2. Configure Semesters**
1. School Admin selects created academic year
2. School Admin clicks "Configure Semesters"
3. System displays semester configuration form
4. School Admin enters for Semester Ganjil:
   - Name (default: "Semester Ganjil")
   - Start date
   - End date
5. School Admin enters for Semester Genap:
   - Name (default: "Semester Genap")
   - Start date
   - End date
6. System validates:
   - Semester dates within academic year range
   - No gaps between semesters
   - No overlaps between semesters
   - Ganjil before Genap
7. System creates semesters in DRAFT status
8. System creates audit log entry
9. System returns success response

**3. Submit for Approval**
1. School Admin clicks "Submit for Approval"
2. System validates:
   - Academic year has 2 semesters configured
   - All required fields populated
3. System changes academic year status to UNDER_REVIEW
4. System sends notification to System Admin
5. System creates audit log entry
6. System returns success response

**4. Approve Academic Year**
1. System Admin navigates to approval queue
2. System Admin selects academic year
3. System Admin reviews configuration
4. System Admin clicks "Approve"
5. System changes academic year status to APPROVED
6. System activates academic year on start date (scheduled)
7. System creates audit log entry
8. System sends notification to School Admin
9. System returns success response

### Alternative Flow

**A1: Reject Academic Year**
1. System Admin selects academic year in UNDER_REVIEW status
2. System Admin clicks "Reject"
3. System Admin enters rejection reason
4. System changes academic year status to REJECTED
5. System creates audit log entry
6. System sends notification to School Admin with reason
7. System returns success response

**A2: Deactivate Academic Year**
1. School Admin selects active academic year
2. School Admin clicks "Deactivate"
3. System requires confirmation and reason
4. System validates:
   - No active curriculum planning in progress
   - No linked data that would be orphaned
5. System changes academic year status to INACTIVE
6. System creates audit log entry
7. System returns success response

### Error Flow

**E1: Overlapping Date Range**
1. School Admin enters academic year dates that overlap with existing academic year
2. System returns validation error: "Academic year dates cannot overlap with existing academic years"
3. System highlights conflicting academic year
4. School Admin adjusts dates

**E2: Insufficient Lead Time**
1. School Admin enters start date less than 30 days in future
2. System returns validation error: "Academic year must be created at least 30 days before start date"
3. School Admin adjusts start date

**E3: Semester Gap or Overlap**
1. School Admin enters semester dates with gap or overlap
2. System returns validation error: "Semesters must fully cover academic year without gaps or overlaps"
3. System visualizes gap/overlap
4. School Admin adjusts semester dates

### Validation Rules

| Field | Validation | Error Message |
| ----- | ---------- | ------------- |
| name | Required, max 100 characters | "Academic year name is required and must be less than 100 characters" |
| start_date | Required, date format, >= today + 30 days | "Start date is required and must be at least 30 days in future" |
| end_date | Required, date format, > start_date | "End date is required and must be after start date" |
| description | Optional, max 500 characters | "Description must be less than 500 characters" |
| semester.name | Required, max 50 characters | "Semester name is required" |
| semester.start_date | Required, date format, >= academic_year.start_date | "Semester start date must be within academic year" |
| semester.end_date | Required, date format, <= academic_year.end_date | "Semester end date must be within academic year" |

### Acceptance Criteria

**AC-001**: School Admin can create academic year with valid date range
**AC-002**: System prevents overlapping academic years for same school
**AC-003**: System requires 30-day lead time for new academic years
**AC-004**: School Admin can configure 2 semesters within academic year
**AC-005**: System validates semester coverage (no gaps, no overlaps)
**AC-006**: System enforces Ganjil before Genap sequence
**AC-007**: School Admin can submit academic year for approval
**AC-008**: System Admin can approve or reject academic year
**AC-009**: System automatically activates academic year on start date
**AC-010**: School Admin can deactivate academic year with approval
**AC-011**: All changes are logged in audit trail
**AC-012**: Past academic years are read-only

---

## Feature 2: Semester Management

### Description
Enable school administrators to create, configure, and manage semesters within academic years with date ranges, sequence validation, and activation workflows.

### Actors
- School Admin (primary)
- System Admin (approval)
- Curriculum Admin (read access)
- Teacher (read access)

### Preconditions
- User must be authenticated
- User must have SCHOOL_ADMIN role
- Academic year must exist and be in DRAFT or UNDER_REVIEW status

### Main Flow

**1. Create Semester**
1. School Admin navigates to Academic Year detail page
2. School Admin clicks "Add Semester"
3. System displays semester creation form
4. School Admin enters:
   - Semester name (e.g., "Semester Ganjil")
   - Sequence (1 for Ganjil, 2 for Genap)
   - Start date
   - End date
   - Description (optional)
5. System validates:
   - Sequence is unique within academic year
   - Dates within academic year range
   - No overlap with existing semesters
6. System creates semester in DRAFT status
7. System links semester to academic year
8. System creates audit log entry
9. System returns success response

**2. Update Semester**
1. School Admin selects existing semester
2. School Admin clicks "Edit"
3. System displays semester edit form
4. School Admin modifies fields
5. System validates updated values
6. System updates semester
7. System creates audit log entry
8. System returns success response

**3. Activate Semester**
1. Academic year is approved
2. System automatically activates first semester on academic year start date
3. System automatically activates second semester on first semester end date
4. System creates audit log entry
5. System sends notification to School Admin

### Alternative Flow

**A1: Delete Semester**
1. School Admin selects semester in DRAFT status
2. School Admin clicks "Delete"
3. System requires confirmation
4. System validates:
   - No linked curriculum planning
   - Not yet activated
5. System deletes semester
6. System creates audit log entry
7. System returns success response

### Error Flow

**E1: Invalid Sequence**
1. School Admin enters sequence that conflicts with existing semester
2. System returns validation error: "Sequence must be unique within academic year"
3. School Admin adjusts sequence

**E2: Date Range Violation**
1. School Admin enters dates outside academic year range
2. System returns validation error: "Semester dates must be within academic year range"
3. School Admin adjusts dates

### Validation Rules

| Field | Validation | Error Message |
| ----- | ---------- | ------------- |
| name | Required, max 50 characters | "Semester name is required" |
| sequence | Required, integer 1 or 2 | "Sequence must be 1 or 2" |
| start_date | Required, date format, >= academic_year.start_date | "Start date must be within academic year" |
| end_date | Required, date format, <= academic_year.end_date | "End date must be within academic year" |
| description | Optional, max 500 characters | "Description must be less than 500 characters" |

### Acceptance Criteria

**AC-013**: School Admin can create semester within academic year
**AC-014**: System enforces unique sequence within academic year
**AC-015**: System validates semester dates within academic year
**AC-016**: System prevents overlapping semesters
**AC-017**: School Admin can update semester in DRAFT status
**AC-018**: School Admin can delete semester in DRAFT status
**AC-019**: System automatically activates semesters on schedule
**AC-020**: All changes are logged in audit trail

---

## Feature 3: Subject Category Management

### Description
Enable curriculum administrators to define and manage subject categories (Intrakurikuler, Kokurikuler, Ekstrakurikuler) and categorize subjects according to Kurikulum Merdeka standards.

### Actors
- Curriculum Admin (primary)
- School Admin (read access)
- Teacher (read access)

### Preconditions
- User must be authenticated
- User must have CURRICULUM_ADMIN role for create/update
- User must have appropriate read permission

### Main Flow

**1. Create Subject Category**
1. Curriculum Admin navigates to Subject Category Management page
2. Curriculum Admin clicks "Create Category"
3. System displays category creation form
4. Curriculum Admin enters:
   - Category code (e.g., "INTRAKURIKULER")
   - Category name (e.g., "Intrakurikuler")
   - Category name (English, optional)
   - Description
   - Guidelines
5. System validates:
   - Code is unique
   - Name is unique
6. System creates category in ACTIVE status
7. System creates audit log entry
8. System returns success response

**2. Categorize Subject**
1. Curriculum Admin navigates to Subject Management page
2. Curriculum Admin selects subject
3. Curriculum Admin clicks "Edit"
4. System displays subject edit form
5. Curriculum Admin selects category from dropdown
6. System saves category assignment
7. System creates audit log entry
8. System returns success response

**3. Bulk Categorize Subjects**
1. Curriculum Admin navigates to Subject Category Management page
2. Curriculum Admin selects category
3. Curriculum Admin clicks "Assign Subjects"
4. System displays subject selection list
5. Curriculum Admin selects multiple subjects
6. System assigns selected subjects to category
7. System creates audit log entry
8. System returns success response

### Alternative Flow

**A1: Deactivate Category**
1. Curriculum Admin selects category
2. Curriculum Admin clicks "Deactivate"
3. System validates:
   - No subjects currently assigned
4. System changes category status to INACTIVE
5. System creates audit log entry
6. System returns success response

### Error Flow

**E1: Category with Subjects**
1. Curriculum Admin attempts to deactivate category with assigned subjects
2. System returns error: "Cannot deactivate category with assigned subjects"
3. System displays count of assigned subjects
4. Curriculum Admin reassigns subjects first

### Validation Rules

| Field | Validation | Error Message |
| ----- | ---------- | ------------- |
| code | Required, unique, uppercase, max 50 characters | "Category code is required and must be unique" |
| name | Required, unique, max 100 characters | "Category name is required and must be unique" |
| name_en | Optional, max 100 characters | "English name must be less than 100 characters" |
| description | Required, max 1000 characters | "Description is required" |
| guidelines | Optional, max 2000 characters | "Guidelines must be less than 2000 characters" |

### Acceptance Criteria

**AC-021**: Curriculum Admin can create subject category
**AC-022**: System enforces unique category codes and names
**AC-023**: Curriculum Admin can categorize subjects
**AC-024**: System supports bulk subject categorization
**AC-025**: System prevents category deletion with assigned subjects
**AC-026**: Curriculum Admin can deactivate unused categories
**AC-027**: All changes are logged in audit trail

---

## Feature 4: Profil Lulusan Management

### Description
Enable curriculum administrators to configure the 8-dimensional Profil Lulusan framework with descriptions, indicators, and alignment tracking to Kurikulum 2026 graduate outcomes.

### Actors
- Curriculum Admin (primary)
- School Admin (read access)
- Teacher (read access)

### Preconditions
- User must be authenticated
- User must have CURRICULUM_ADMIN role for create/update
- User must have appropriate read permission

### Main Flow

**1. Create Graduate Profile Dimension**
1. Curriculum Admin navigates to Profil Lulusan Management page
2. Curriculum Admin clicks "Create Dimension"
3. System displays dimension creation form
4. Curriculum Admin enters:
   - Dimension code (e.g., "KEIMANAN_KETAKWAAN")
   - Dimension name (e.g., "Keimanan & Ketakwaan")
   - Dimension name (English, optional)
   - Description
   - Indicators (JSON array)
   - Weight (for alignment calculation)
5. System validates:
   - Code is unique
   - Name is unique
   - Weight is positive
6. System creates dimension in ACTIVE status
7. System creates audit log entry
8. System returns success response

**2. Configure 8 Dimensions**
1. Curriculum Admin creates all 8 dimensions:
   - Keimanan & Ketakwaan
   - Kewargaan
   - Berakhlak Mulia
   - Berani Bertanggung Jawab
   - Peduli
   - Gotong Royong
   - Mandiri
   - Kreatif
2. System validates that exactly 8 dimensions exist
3. System validates that all dimensions are active
4. System publishes dimension configuration
5. System creates audit log entry
6. System returns success response

**3. Link CP to Dimensions**
1. Curriculum Admin navigates to CP Management page
2. Curriculum Admin selects CP
3. Curriculum Admin clicks "Align to Profil Lulusan"
4. System displays dimension selection form
5. Curriculum Admin selects one or more dimensions
6. Curriculum Admin assigns alignment strength (Strong, Medium, Weak)
7. System saves alignment
8. System calculates alignment percentage
9. System creates audit log entry
10. System returns success response

**4. View Alignment Report**
1. Curriculum Admin navigates to Alignment Report page
2. System selects academic year and phase
3. System generates alignment report showing:
   - CP count per dimension
   - Alignment percentage per dimension
   - Gaps in dimension coverage
4. System visualizes alignment distribution
5. Curriculum Admin exports report as PDF

### Alternative Flow

**A1: Update Dimension**
1. Curriculum Admin selects existing dimension
2. Curriculum Admin clicks "Edit"
3. System displays dimension edit form
4. Curriculum Admin updates fields
5. System validates updates
6. System updates dimension
7. System creates audit log entry
8. System returns success response

**A2: Deactivate Dimension**
1. Curriculum Admin selects dimension
2. Curriculum Admin clicks "Deactivate"
3. System validates:
   - No CP currently aligned to dimension
4. System changes dimension status to INACTIVE
5. System creates audit log entry
6. System returns success response

### Error Flow

**E1: Dimension with Aligned CP**
1. Curriculum Admin attempts to deactivate dimension with aligned CP
2. System returns error: "Cannot deactivate dimension with aligned CP"
3. System displays count of aligned CP
4. Curriculum Admin removes alignments first

### Validation Rules

| Field | Validation | Error Message |
| ----- | ---------- | ------------- |
| code | Required, unique, uppercase, max 50 characters | "Dimension code is required and must be unique" |
| name | Required, unique, max 100 characters | "Dimension name is required and must be unique" |
| name_en | Optional, max 100 characters | "English name must be less than 100 characters" |
| description | Required, max 1000 characters | "Description is required" |
| indicators | Required, valid JSON array | "Indicators must be valid JSON array" |
| weight | Required, positive decimal, max 1.0 | "Weight must be positive and <= 1.0" |

### Acceptance Criteria

**AC-028**: Curriculum Admin can create graduate profile dimension
**AC-029**: System enforces unique dimension codes and names
**AC-030**: Curriculum Admin can configure exactly 8 dimensions
**AC-031**: Curriculum Admin can link CP to dimensions
**AC-032**: System calculates alignment percentage automatically
**AC-033**: Curriculum Admin can view alignment reports
**AC-034**: System identifies gaps in dimension coverage
**AC-035**: System prevents dimension deletion with aligned CP
**AC-036**: All changes are logged in audit trail

---

## Feature 5: Kurikulum Alignment Management

### Description
Enable curriculum administrators to track and manage alignment between CP (Capaian Pembelajaran) and graduate profile dimensions, ensuring curriculum coverage of all 8 dimensions with quantitative alignment scoring.

### Actors
- Curriculum Admin (primary)
- School Admin (read access)
- Teacher (read access)

### Preconditions
- User must be authenticated
- User must have CURRICULUM_ADMIN role for manage alignment
- User must have appropriate read permission
- Graduate profile dimensions must be configured
- CP must exist

### Main Flow

**1. Create Alignment**
1. Curriculum Admin navigates to CP detail page
2. Curriculum Admin clicks "Add Alignment"
3. System displays alignment creation form
4. Curriculum Admin selects:
   - Graduate profile dimension
   - Alignment strength (Strong = 100%, Medium = 75%, Weak = 50%)
   - Rationale (optional)
5. System validates:
   - CP not already aligned to same dimension
6. System creates alignment
7. System recalculates CP overall alignment percentage
8. System creates audit log entry
9. System returns success response

**2. Update Alignment**
1. Curriculum Admin selects existing alignment
2. Curriculum Admin clicks "Edit"
3. System displays alignment edit form
4. Curriculum Admin modifies alignment strength
5. System recalculates CP overall alignment percentage
6. System updates alignment
7. System creates audit log entry
8. System returns success response

**3. Delete Alignment**
1. Curriculum Admin selects alignment
2. Curriculum Admin clicks "Delete"
3. System requires confirmation
4. System deletes alignment
5. System recalculates CP overall alignment percentage
6. System creates audit log entry
7. System returns success response

**4. Generate Alignment Report**
1. Curriculum Admin navigates to Alignment Report page
2. Curriculum Admin selects filters:
   - Academic year
   - Phase
   - Subject
3. System generates alignment report showing:
   - Total CP count
   - CP aligned count
   - Overall alignment percentage
   - Alignment per dimension
   - CP below 60% alignment threshold
4. System visualizes data with charts
5. Curriculum Admin exports report as PDF/CSV

### Alternative Flow

**A1: Bulk Align CP**
1. Curriculum Admin navigates to CP Management page
2. Curriculum Admin selects multiple CP
3. Curriculum Admin clicks "Bulk Align"
4. System displays bulk alignment form
5. Curriculum Admin selects dimension(s) for all selected CP
6. Curriculum Admin sets default alignment strength
7. System creates alignments for all selected CP
8. System recalculates alignment percentages
9. System creates audit log entry
10. System returns success response

### Error Flow

**E1: Minimum Alignment Violation**
1. Curriculum Admin deletes alignment causing CP to fall below 60% overall alignment
2. System returns warning: "CP alignment will fall below 60% threshold"
3. System requires confirmation to proceed
4. Curriculum Admin confirms or adjusts

### Validation Rules

| Field | Validation | Error Message |
| ----- | ---------- | ------------- |
| dimension_id | Required, exists | "Dimension is required" |
| alignment_strength | Required, enum (STRONG, MEDIUM, WEAK) | "Alignment strength is required" |
| rationale | Optional, max 500 characters | "Rationale must be less than 500 characters" |

### Acceptance Criteria

**AC-037**: Curriculum Admin can create CP-dimension alignment
**AC-038**: System prevents duplicate alignments
**AC-039**: System calculates alignment percentage automatically
**AC-040**: Curriculum Admin can update alignment strength
**AC-041**: Curriculum Admin can delete alignment
**AC-042**: System warns when CP falls below 60% alignment
**AC-043**: Curriculum Admin can generate alignment reports
**AC-044**: System visualizes alignment distribution
**AC-045**: Curriculum Admin can bulk align CP
**AC-046**: All changes are logged in audit trail

---

## Feature 6: Koding & AI Subject Support

### Description
Enable the platform to support Koding, AI, and Numerasi as formal subjects within the curriculum structure, properly categorized and integrated with existing curriculum planning workflows.

### Actors
- Curriculum Admin (primary)
- School Admin (read access)
- Teacher (read access)

### Preconditions
- User must be authenticated
- User must have CURRICULUM_ADMIN role for create subjects
- User must have appropriate read permission
- Subject categories must be configured

### Main Flow

**1. Add Koding Subject**
1. Curriculum Admin navigates to Subject Management page
2. Curriculum Admin clicks "Create Subject"
3. System displays subject creation form
4. Curriculum Admin enters:
   - Code: "KODING"
   - Name: "Koding"
   - Name (English): "Coding"
   - Description: "Mata pelajaran koding dan pemrograman"
   - Category: Intrakurikuler
   - Is Active: true
5. System validates:
   - Code is unique
   - Name is unique
   - Category exists
6. System creates Koding subject
7. System creates audit log entry
8. System returns success response

**2. Add AI Subject**
1. Curriculum Admin repeats process for AI subject
2. Curriculum Admin enters:
   - Code: "AI"
   - Name: "Kecerdasan Buatan"
   - Name (English): "Artificial Intelligence"
   - Description: "Mata pelajaran kecerdasan buatan"
   - Category: Intrakurikuler
   - Is Active: true
3. System creates AI subject
4. System creates audit log entry
5. System returns success response

**3. Add Numerasi Subject**
1. Curriculum Admin repeats process for Numerasi subject
2. Curriculum Admin enters:
   - Code: "NUMERASI"
   - Name: "Numerasi"
   - Name (English): "Numeracy"
   - Description: "Mata pelajaran numerasi"
   - Category: Intrakurikuler
   - Is Active: true
3. System creates Numerasi subject
4. System creates audit log entry
5. System returns success response

**4. Verify Integration**
1. Curriculum Admin navigates to CP Management
2. Curriculum Admin selects Koding subject
3. Curriculum Admin creates CP for Koding
4. System verifies CP creation succeeds
5. Curriculum Admin creates TP for Koding CP
6. System verifies TP creation succeeds
7. System confirms integration is working

### Alternative Flow

**A1: Categorize Existing Subjects**
1. Curriculum Admin navigates to Subject Management page
2. Curriculum Admin selects existing subject (e.g., Matematika)
3. Curriculum Admin clicks "Edit"
4. Curriculum Admin selects category (e.g., Intrakurikuler)
5. System saves category assignment
6. System creates audit log entry
7. System returns success response

### Error Flow

**E1: Duplicate Subject**
1. Curriculum Admin attempts to create subject with existing code
2. System returns error: "Subject code already exists"
3. Curriculum Admin adjusts code

### Validation Rules

| Field | Validation | Error Message |
| ----- | ---------- | ------------- |
| code | Required, unique, uppercase, max 50 characters | "Subject code is required and must be unique" |
| name | Required, unique, max 255 characters | "Subject name is required and must be unique" |
| name_en | Optional, max 255 characters | "English name must be less than 255 characters" |
| description | Required, max 1000 characters | "Description is required" |
| category_id | Required, exists | "Category is required" |

### Acceptance Criteria

**AC-047**: Curriculum Admin can add Koding subject
**AC-048**: Curriculum Admin can add AI subject
**AC-049**: Curriculum Admin can add Numerasi subject
**AC-050**: System categorizes new subjects as Intrakurikuler
**AC-051**: New subjects integrate with existing CP/TP workflow
**AC-052**: Curriculum Admin can categorize existing subjects
**AC-053**: All changes are logged in audit trail

---

# PART 4 – DOMAIN MODEL

## Bounded Context Decision

### Decision: NO NEW BOUNDED CONTEXTS

**Rationale:**
1. All Sprint 4 requirements are extensions to the existing Curriculum Context
2. Academic years, semesters, subject categories, and graduate profile dimensions are reference data for curriculum planning
3. These entities do not represent new business domains - they are infrastructure for the existing curriculum domain
4. Following DDD principle: "Bounded contexts should be based on business capabilities, not technical entities"
5. Adding new bounded contexts would violate the "minimize new bounded contexts" constraint

**Conclusion:**
All Sprint 4 features will be implemented within the existing **Curriculum Context** as extensions to reference data and governance structures.

---

## Aggregates

### Aggregate 1: AcademicYear

**Purpose**: Manage academic year configuration with semesters

**Aggregate Root**: `AcademicYear`

**Entities**:
- `AcademicYear` (Aggregate Root)
- `Semester` (Entity)

**Value Objects**:
- `DateRange` (Value Object)
- `AcademicYearName` (Value Object)

**Domain Services**:
- `AcademicYearValidationService` - Validates academic year rules (no overlap, 30-day lead time)
- `AcademicYearActivationService` - Manages automatic activation scheduling

**Repository**: `AcademicYearRepository`

**Domain Events**: None (no cross-aggregate communication required)

### Aggregate 2: SubjectCategory

**Purpose**: Manage subject categorization according to Kurikulum Merdeka

**Aggregate Root**: `SubjectCategory`

**Entities**:
- `SubjectCategory` (Aggregate Root)

**Value Objects**:
- `CategoryCode` (Value Object)
- `CategoryName` (Value Object)
- `CategoryGuidelines` (Value Object)

**Domain Services**: None

**Repository**: `SubjectCategoryRepository`

**Domain Events**: None

### Aggregate 3: GraduateProfileDimension

**Purpose**: Manage graduate profile dimensions and CP alignment

**Aggregate Root**: `GraduateProfileDimension`

**Entities**:
- `GraduateProfileDimension` (Aggregate Root)
- `CPAlignment` (Entity)

**Value Objects**:
- `DimensionCode` (Value Object)
- `DimensionIndicators` (Value Object)
- `AlignmentStrength` (Value Object)

**Domain Services**:
- `AlignmentCalculationService` - Calculates CP alignment percentage

**Repository**: `GraduateProfileDimensionRepository`

**Domain Events**: None

---

## Entities

### AcademicYear

```go
type AcademicYear struct {
    ID              UUID
    SchoolID        UUID
    Name            string          // e.g., "2026/2027"
    StartDate       time.Time
    EndDate         time.Time
    Status          AcademicYearStatus
    ApprovalStatus  ApprovalStatus
    ApprovedBy      *UUID
    ApprovedAt      *time.Time
    CreatedBy       UUID
    CreatedAt       time.Time
    UpdatedAt       time.Time
    
    // Child entities
    Semesters       []Semester
}

type AcademicYearStatus string
const (
    AcademicYearStatusDraft      AcademicYearStatus = "DRAFT"
    AcademicYearStatusUnderReview AcademicYearStatus = "UNDER_REVIEW"
    AcademicYearStatusApproved   AcademicYearStatus = "APPROVED"
    AcademicYearStatusActive     AcademicYearStatus = "ACTIVE"
    AcademicYearStatusInactive   AcademicYearStatus = "INACTIVE"
)

func (ay *AcademicYear) Validate() error {
    // Business rules:
    // 1. StartDate < EndDate
    // 2. StartDate >= today + 30 days (for new creation)
    // 3. Must have exactly 2 semesters
    // 4. Semesters must fully cover date range
}
```

### Semester

```go
type Semester struct {
    ID              UUID
    AcademicYearID  UUID
    Name            string          // e.g., "Semester Ganjil"
    Sequence        int             // 1 or 2
    StartDate       time.Time
    EndDate         time.Time
    Status          SemesterStatus
    CreatedAt       time.Time
    UpdatedAt       time.Time
}

type SemesterStatus string
const (
    SemesterStatusDraft      SemesterStatus = "DRAFT"
    SemesterStatusActive     SemesterStatus = "ACTIVE"
    SemesterStatusInactive   SemesterStatus = "INACTIVE"
)

func (s *Semester) Validate(academicYear AcademicYear) error {
    // Business rules:
    // 1. Sequence is unique within academic year
    // 2. Dates within academic year range
    // 3. No overlap with other semesters
}
```

### SubjectCategory

```go
type SubjectCategory struct {
    ID          UUID
    Code        string          // e.g., "INTRAKURIKULER"
    Name        string          // e.g., "Intrakurikuler"
    NameEN      *string
    Description string
    Guidelines  *string
    Status      SubjectCategoryStatus
    CreatedBy   UUID
    CreatedAt   time.Time
    UpdatedAt   time.Time
}

type SubjectCategoryStatus string
const (
    SubjectCategoryStatusActive   SubjectCategoryStatus = "ACTIVE"
    SubjectCategoryStatusInactive SubjectCategoryStatus = "INACTIVE"
)

func (sc *SubjectCategory) Validate() error {
    // Business rules:
    // 1. Code is unique
    // 2. Name is unique
}
```

### GraduateProfileDimension

```go
type GraduateProfileDimension struct {
    ID          UUID
    Code        string          // e.g., "KEIMANAN_KETAKWAAN"
    Name        string          // e.g., "Keimanan & Ketakwaan"
    NameEN      *string
    Description string
    Indicators  JSON            // Array of indicators
    Weight      float64         // For alignment calculation
    Status      DimensionStatus
    CreatedBy   UUID
    CreatedAt   time.Time
    UpdatedAt   time.Time
    
    // Child entities
    Alignments  []CPAlignment
}

type DimensionStatus string
const (
    DimensionStatusActive   DimensionStatus = "ACTIVE"
    DimensionStatusInactive DimensionStatus = "INACTIVE"
)

func (gpd *GraduateProfileDimension) Validate() error {
    // Business rules:
    // 1. Code is unique
    // 2. Name is unique
    // 3. Weight is positive and <= 1.0
    // 4. Exactly 8 dimensions can be active
}
```

### CPAlignment

```go
type CPAlignment struct {
    ID                  UUID
    CPID                UUID
    DimensionID         UUID
    AlignmentStrength   AlignmentStrength
    Rationale           *string
    CreatedBy           UUID
    CreatedAt           time.Time
    UpdatedAt           time.Time
}

type AlignmentStrength string
const (
    AlignmentStrengthStrong  AlignmentStrength = "STRONG"  // 100%
    AlignmentStrengthMedium  AlignmentStrength = "MEDIUM"  // 75%
    AlignmentStrengthWeak    AlignmentStrength = "WEAK"    // 50%
)

func (cpa *CPAlignment) Validate() error {
    // Business rules:
    // 1. CP cannot have duplicate alignment to same dimension
    // 2. Alignment strength is valid enum
}
```

---

## Value Objects

### DateRange

```go
type DateRange struct {
    StartDate time.Time
    EndDate   time.Time
}

func (dr DateRange) IsValid() bool {
    return dr.StartDate.Before(dr.EndDate)
}

func (dr DateRange) Overlaps(other DateRange) bool {
    return !(dr.EndDate.Before(other.StartDate) || dr.StartDate.After(other.EndDate))
}
```

### AcademicYearName

```go
type AcademicYearName string

func NewAcademicYearName(name string) (AcademicYearName, error) {
    if len(name) == 0 || len(name) > 100 {
        return "", errors.New("academic year name must be 1-100 characters")
    }
    return AcademicYearName(name), nil
}
```

### CategoryCode

```go
type CategoryCode string

func NewCategoryCode(code string) (CategoryCode, error) {
    if len(code) == 0 || len(code) > 50 {
        return "", errors.New("category code must be 1-50 characters")
    }
    if code != strings.ToUpper(code) {
        return "", errors.New("category code must be uppercase")
    }
    return CategoryCode(code), nil
}
```

### DimensionIndicators

```go
type DimensionIndicators []string

func NewDimensionIndicators(indicators []string) (DimensionIndicators, error) {
    if len(indicators) == 0 {
        return nil, errors.New("at least one indicator is required")
    }
    return DimensionIndicators(indicators), nil
}
```

---

## Domain Services

### AcademicYearValidationService

```go
type AcademicYearValidationService struct{}

func (s *AcademicYearValidationService) ValidateNewAcademicYear(
    schoolID UUID,
    dateRange DateRange,
    existingYears []AcademicYear,
) error {
    // Validate 30-day lead time
    if dateRange.StartDate.Before(time.Now().AddDate(0, 0, 30)) {
        return errors.New("academic year must be created at least 30 days in advance")
    }
    
    // Validate no overlap with existing academic years
    for _, year := range existingYears {
        if year.SchoolID == schoolID {
            existingRange := DateRange{StartDate: year.StartDate, EndDate: year.EndDate}
            if dateRange.Overlaps(existingRange) {
                return errors.New("academic year dates cannot overlap with existing academic years")
            }
        }
    }
    
    return nil
}

func (s *AcademicYearValidationService) ValidateSemesters(
    academicYear AcademicYear,
    semesters []Semester,
) error {
    // Validate exactly 2 semesters
    if len(semesters) != 2 {
        return errors.New("academic year must have exactly 2 semesters")
    }
    
    // Validate sequence (1 and 2)
    sequences := make(map[int]bool)
    for _, sem := range semesters {
        if sequences[sem.Sequence] {
            return errors.New("semester sequences must be unique")
        }
        sequences[sem.Sequence] = true
    }
    if !sequences[1] || !sequences[2] {
        return errors.New("semesters must have sequences 1 and 2")
    }
    
    // Validate full coverage without gaps
    academicYearRange := DateRange{StartDate: academicYear.StartDate, EndDate: academicYear.EndDate}
    totalCoverage := 0
    for _, sem := range semesters {
        semesterRange := DateRange{StartDate: sem.StartDate, EndDate: sem.EndDate}
        if !academicYearRange.Contains(semesterRange) {
            return errors.New("semester dates must be within academic year range")
        }
        totalCoverage += int(semesterRange.EndDate.Sub(semesterRange.StartDate).Hours())
    }
    
    expectedCoverage := int(academicYearRange.EndDate.Sub(academicYearRange.StartDate).Hours())
    if totalCoverage != expectedCoverage {
        return errors.New("semesters must fully cover academic year without gaps")
    }
    
    return nil
}
```

### AlignmentCalculationService

```go
type AlignmentCalculationService struct{}

func (s *AlignmentCalculationService) CalculateCPAlignment(
    alignments []CPAlignment,
    dimensions []GraduateProfileDimension,
) (float64, error) {
    if len(alignments) == 0 {
        return 0, nil
    }
    
    totalWeight := 0.0
    totalScore := 0.0
    
    for _, alignment := range alignments {
        dimension := findDimension(dimensions, alignment.DimensionID)
        if dimension == nil {
            continue
        }
        
        strength := s.strengthToPercentage(alignment.AlignmentStrength)
        weightedScore := strength * dimension.Weight
        totalScore += weightedScore
        totalWeight += dimension.Weight
    }
    
    if totalWeight == 0 {
        return 0, nil
    }
    
    return (totalScore / totalWeight) * 100, nil
}

func (s *AlignmentCalculationService) strengthToPercentage(strength AlignmentStrength) float64 {
    switch strength {
    case AlignmentStrengthStrong:
        return 1.0  // 100%
    case AlignmentStrengthMedium:
        return 0.75 // 75%
    case AlignmentStrengthWeak:
        return 0.5  // 50%
    default:
        return 0.0
    }
}
```

---

## Repositories

### AcademicYearRepository Interface

```go
type AcademicYearRepository interface {
    Create(academicYear *AcademicYear) error
    GetByID(id UUID) (*AcademicYear, error)
    GetBySchoolID(schoolID UUID) ([]AcademicYear, error)
    GetActiveBySchoolID(schoolID UUID) (*AcademicYear, error)
    Update(academicYear *AcademicYear) error
    Delete(id UUID) error
    CheckOverlap(schoolID UUID, startDate, endDate time.Time) (bool, error)
}
```

### SubjectCategoryRepository Interface

```go
type SubjectCategoryRepository interface {
    Create(category *SubjectCategory) error
    GetByID(id UUID) (*SubjectCategory, error)
    GetAll() ([]SubjectCategory, error)
    GetActive() ([]SubjectCategory, error)
    Update(category *SubjectCategory) error
    Delete(id UUID) error
    CheckSubjectCount(categoryID UUID) (int, error)
}
```

### GraduateProfileDimensionRepository Interface

```go
type GraduateProfileDimensionRepository interface {
    Create(dimension *GraduateProfileDimension) error
    GetByID(id UUID) (*GraduateProfileDimension, error)
    GetAll() ([]GraduateProfileDimension, error)
    GetActive() ([]GraduateProfileDimension, error)
    Update(dimension *GraduateProfileDimension) error
    Delete(id UUID) error
    GetActiveCount() (int, error)
    CheckCPCount(dimensionID UUID) (int, error)
}
```

### CPAlignmentRepository Interface

```go
type CPAlignmentRepository interface {
    Create(alignment *CPAlignment) error
    GetByID(id UUID) (*CPAlignment, error)
    GetByCPID(cpID UUID) ([]CPAlignment, error)
    GetByDimensionID(dimensionID UUID) ([]CPAlignment, error)
    Update(alignment *CPAlignment) error
    Delete(id UUID) error
    CheckDuplicate(cpID, dimensionID UUID) (bool, error)
    GetAlignmentReport(filters AlignmentReportFilters) (*AlignmentReport, error)
}
```

---

## Domain Events

### Decision: NO DOMAIN EVENTS

**Rationale:**
1. All Sprint 4 aggregates operate within a single bounded context (Curriculum Context)
2. No cross-context communication is required
3. Academic year, semester, category, and dimension changes are administrative in nature
4. No event-driven workflows are triggered by these changes
5. Following YAGNI (You Aren't Gonna Need It) principle
6. Adding events would violate the "no Event Bus" constraint in Architecture Freeze v2

**Conclusion:**
No domain events are needed for Sprint 4. All communication happens via direct repository and service calls within the same bounded context.

---

# PART 5 – DATABASE DESIGN

## New Tables

### Table: academic_years

**Purpose**: Store academic year configuration with approval workflow and activation status

**Columns**:
| Column | Type | Constraints | Description |
| ------ | ---- | ----------- | ----------- |
| id | UUID | PRIMARY KEY, DEFAULT gen_uuid_v7() | Unique identifier |
| school_id | UUID | NOT NULL, FK schools(id) ON DELETE RESTRICT | School ownership |
| name | VARCHAR(100) | NOT NULL | Academic year name (e.g., "2026/2027") |
| start_date | TIMESTAMP WITH TIME ZONE | NOT NULL | Academic year start date |
| end_date | TIMESTAMP WITH TIME ZONE | NOT NULL | Academic year end date |
| status | VARCHAR(20) | NOT NULL, CHECK (status IN ('DRAFT', 'UNDER_REVIEW', 'APPROVED', 'ACTIVE', 'INACTIVE')) | Workflow status |
| approved_by | UUID | FK users(id) ON DELETE SET NULL | Approver user reference |
| approved_at | TIMESTAMP WITH TIME ZONE | | Approval timestamp |
| created_by | UUID | NOT NULL, FK users(id) ON DELETE RESTRICT | Creator user reference |
| created_at | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Creation timestamp |
| updated_at | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Last update timestamp |

**Indexes**:
- `idx_academic_years_school_id` on (school_id)
- `idx_academic_years_status` on (status)
- `idx_academic_years_dates` on (school_id, start_date, end_date)
- `idx_academic_years_school_status` on (school_id, status) WHERE status = 'ACTIVE'

**Unique Constraints**:
- `uq_academic_years_school_name` on (school_id, name)

**Foreign Keys**:
- `fk_academic_years_school_id` → schools(id)
- `fk_academic_years_approved_by` → users(id)
- `fk_academic_years_created_by` → users(id)

**Audit Fields**: created_by, created_at, updated_at

---

### Table: semesters

**Purpose**: Store semester configuration within academic years

**Columns**:
| Column | Type | Constraints | Description |
| ------ | ---- | ----------- | ----------- |
| id | UUID | PRIMARY KEY, DEFAULT gen_uuid_v7() | Unique identifier |
| academic_year_id | UUID | NOT NULL, FK academic_years(id) ON DELETE CASCADE | Parent academic year |
| name | VARCHAR(50) | NOT NULL | Semester name (e.g., "Semester Ganjil") |
| sequence | INTEGER | NOT NULL, CHECK (sequence IN (1, 2)) | Semester sequence (1=Ganjil, 2=Genap) |
| start_date | TIMESTAMP WITH TIME ZONE | NOT NULL | Semester start date |
| end_date | TIMESTAMP WITH TIME ZONE | NOT NULL | Semester end date |
| status | VARCHAR(20) | NOT NULL, CHECK (status IN ('DRAFT', 'ACTIVE', 'INACTIVE')) | Status |
| created_at | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Creation timestamp |
| updated_at | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Last update timestamp |

**Indexes**:
- `idx_semesters_academic_year_id` on (academic_year_id)
- `idx_semesters_sequence` on (academic_year_id, sequence)
- `idx_semesters_dates` on (academic_year_id, start_date, end_date)
- `idx_semesters_status` on (status)

**Unique Constraints**:
- `uq_semesters_academic_year_sequence` on (academic_year_id, sequence)

**Foreign Keys**:
- `fk_semesters_academic_year_id` → academic_years(id)

**Audit Fields**: created_at, updated_at

---

### Table: subject_categories

**Purpose**: Store subject category definitions (Intrakurikuler, Kokurikuler, Ekstrakurikuler)

**Columns**:
| Column | Type | Constraints | Description |
| ------ | ---- | ----------- | ----------- |
| id | UUID | PRIMARY KEY, DEFAULT gen_uuid_v7() | Unique identifier |
| code | VARCHAR(50) | NOT NULL, UNIQUE | Category code (e.g., "INTRAKURIKULER") |
| name | VARCHAR(100) | NOT NULL, UNIQUE | Category name (e.g., "Intrakurikuler") |
| name_en | VARCHAR(100) | | Category name in English |
| description | TEXT | NOT NULL | Category description |
| guidelines | TEXT | | Category implementation guidelines |
| status | VARCHAR(20) | NOT NULL, CHECK (status IN ('ACTIVE', 'INACTIVE')) | Status |
| created_by | UUID | NOT NULL, FK users(id) ON DELETE RESTRICT | Creator user reference |
| created_at | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Creation timestamp |
| updated_at | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Last update timestamp |

**Indexes**:
- `idx_subject_categories_code` on (code)
- `idx_subject_categories_name` on (name)
- `idx_subject_categories_status` on (status)

**Unique Constraints**:
- `uq_subject_categories_code` on (code)
- `uq_subject_categories_name` on (name)

**Foreign Keys**:
- `fk_subject_categories_created_by` → users(id)

**Audit Fields**: created_by, created_at, updated_at

---

### Table: graduate_profile_dimensions

**Purpose**: Store graduate profile dimension definitions (8 dimensions of Profil Lulusan)

**Columns**:
| Column | Type | Constraints | Description |
| ------ | ---- | ----------- | ----------- |
| id | UUID | PRIMARY KEY, DEFAULT gen_uuid_v7() | Unique identifier |
| code | VARCHAR(50) | NOT NULL, UNIQUE | Dimension code (e.g., "KEIMANAN_KETAKWAAN") |
| name | VARCHAR(100) | NOT NULL, UNIQUE | Dimension name (e.g., "Keimanan & Ketakwaan") |
| name_en | VARCHAR(100) | | Dimension name in English |
| description | TEXT | NOT NULL | Dimension description |
| indicators | JSONB | NOT NULL | Array of dimension indicators |
| weight | DECIMAL(5,4) | NOT NULL, CHECK (weight > 0 AND weight <= 1.0) | Weight for alignment calculation |
| status | VARCHAR(20) | NOT NULL, CHECK (status IN ('ACTIVE', 'INACTIVE')) | Status |
| created_by | UUID | NOT NULL, FK users(id) ON DELETE RESTRICT | Creator user reference |
| created_at | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Creation timestamp |
| updated_at | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Last update timestamp |

**Indexes**:
- `idx_graduate_profile_dimensions_code` on (code)
- `idx_graduate_profile_dimensions_name` on (name)
- `idx_graduate_profile_dimensions_status` on (status)
- `idx_graduate_profile_dimensions_indicators` on (indicators) USING GIN

**Unique Constraints**:
- `uq_graduate_profile_dimensions_code` on (code)
- `uq_graduate_profile_dimensions_name` on (name)

**Foreign Keys**:
- `fk_graduate_profile_dimensions_created_by` → users(id)

**Audit Fields**: created_by, created_at, updated_at

---

### Table: cp_alignments

**Purpose**: Store CP to graduate profile dimension alignments with strength scoring

**Columns**:
| Column | Type | Constraints | Description |
| ------ | ---- | ----------- | ----------- |
| id | UUID | PRIMARY KEY, DEFAULT gen_uuid_v7() | Unique identifier |
| cp_id | UUID | NOT NULL, FK cp(id) ON DELETE CASCADE | CP reference |
| dimension_id | UUID | NOT NULL, FK graduate_profile_dimensions(id) ON DELETE CASCADE | Dimension reference |
| alignment_strength | VARCHAR(20) | NOT NULL, CHECK (alignment_strength IN ('STRONG', 'MEDIUM', 'WEAK')) | Alignment strength |
| rationale | TEXT | | Alignment rationale |
| created_by | UUID | NOT NULL, FK users(id) ON DELETE RESTRICT | Creator user reference |
| created_at | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Creation timestamp |
| updated_at | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Last update timestamp |

**Indexes**:
- `idx_cp_alignments_cp_id` on (cp_id)
- `idx_cp_alignments_dimension_id` on (dimension_id)
- `idx_cp_alignments_strength` on (alignment_strength)

**Unique Constraints**:
- `uq_cp_alignments_cp_dimension` on (cp_id, dimension_id)

**Foreign Keys**:
- `fk_cp_alignments_cp_id` → cp(id)
- `fk_cp_alignments_dimension_id` → graduate_profile_dimensions(id)
- `fk_cp_alignments_created_by` → users(id)

**Audit Fields**: created_by, created_at, updated_at

---

## Table Extensions

### Extension: curriculum_subjects

**Purpose**: Add subject categorization to existing subjects table

**New Column**:
| Column | Type | Constraints | Description |
| ------ | ---- | ----------- | ----------- |
| subject_category_id | UUID | FK subject_categories(id) ON DELETE SET NULL | Subject category reference |

**Migration Action**: Add column, create index, update existing records

**Index**: `idx_curriculum_subjects_category_id` on (subject_category_id)

**Foreign Key**: `fk_curriculum_subjects_category_id` → subject_categories(id)

**Note**: Column will be nullable initially, then populated via data migration

---

### Extension: cp

**Purpose**: Add academic year scoping to existing CP table

**New Columns**:
| Column | Type | Constraints | Description |
| ------ | ---- | ----------- | ----------- |
| academic_year_id | UUID | FK academic_years(id) ON DELETE SET NULL | Academic year reference |
| semester_id | UUID | FK semesters(id) ON DELETE SET NULL | Semester reference |

**Migration Action**: Add columns, create indexes

**Indexes**:
- `idx_cp_academic_year_id` on (academic_year_id)
- `idx_cp_semester_id` on (semester_id)
- `idx_cp_academic_semester` on (academic_year_id, semester_id)

**Foreign Keys**:
- `fk_cp_academic_year_id` → academic_years(id)
- `fk_cp_semester_id` → semesters(id)

**Note**: Columns will be nullable to support existing CP without academic year scoping

---

## Migration Strategy

### Migration File: 000010_sprint4_academic_foundation.up.sql

```sql
-- Sprint 4 Academic Foundation Migration
-- Purpose: Add academic year, semester, subject category, graduate profile dimension infrastructure
-- Risk Level: MEDIUM (new tables, existing table extensions)
-- Tables: 5 new, 2 extended

-- ============================================================================
-- NEW TABLES
-- ============================================================================

-- Table: academic_years
CREATE TABLE academic_years (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    school_id UUID NOT NULL,
    name VARCHAR(100) NOT NULL,
    start_date TIMESTAMP WITH TIME ZONE NOT NULL,
    end_date TIMESTAMP WITH TIME ZONE NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'DRAFT' CHECK (status IN ('DRAFT', 'UNDER_REVIEW', 'APPROVED', 'ACTIVE', 'INACTIVE')),
    approved_by UUID,
    approved_at TIMESTAMP WITH TIME ZONE,
    created_by UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_academic_years_school_id FOREIGN KEY (school_id) REFERENCES schools(id) ON DELETE RESTRICT,
    CONSTRAINT fk_academic_years_approved_by FOREIGN KEY (approved_by) REFERENCES users(id) ON DELETE SET NULL,
    CONSTRAINT fk_academic_years_created_by FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE RESTRICT,
    CONSTRAINT uq_academic_years_school_name UNIQUE (school_id, name),
    CONSTRAINT chk_academic_years_dates CHECK (start_date < end_date)
);

CREATE INDEX idx_academic_years_school_id ON academic_years(school_id);
CREATE INDEX idx_academic_years_status ON academic_years(status);
CREATE INDEX idx_academic_years_dates ON academic_years(school_id, start_date, end_date);
CREATE INDEX idx_academic_years_school_status ON academic_years(school_id, status) WHERE status = 'ACTIVE';

-- Table: semesters
CREATE TABLE semesters (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    academic_year_id UUID NOT NULL,
    name VARCHAR(50) NOT NULL,
    sequence INTEGER NOT NULL CHECK (sequence IN (1, 2)),
    start_date TIMESTAMP WITH TIME ZONE NOT NULL,
    end_date TIMESTAMP WITH TIME ZONE NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'DRAFT' CHECK (status IN ('DRAFT', 'ACTIVE', 'INACTIVE')),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_semesters_academic_year_id FOREIGN KEY (academic_year_id) REFERENCES academic_years(id) ON DELETE CASCADE,
    CONSTRAINT uq_semesters_academic_year_sequence UNIQUE (academic_year_id, sequence),
    CONSTRAINT chk_semesters_dates CHECK (start_date < end_date)
);

CREATE INDEX idx_semesters_academic_year_id ON semesters(academic_year_id);
CREATE INDEX idx_semesters_sequence ON semesters(academic_year_id, sequence);
CREATE INDEX idx_semesters_dates ON semesters(academic_year_id, start_date, end_date);
CREATE INDEX idx_semesters_status ON semesters(status);

-- Table: subject_categories
CREATE TABLE subject_categories (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    code VARCHAR(50) NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL UNIQUE,
    name_en VARCHAR(100),
    description TEXT NOT NULL,
    guidelines TEXT,
    status VARCHAR(20) NOT NULL DEFAULT 'ACTIVE' CHECK (status IN ('ACTIVE', 'INACTIVE')),
    created_by UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_subject_categories_created_by FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE RESTRICT
);

CREATE INDEX idx_subject_categories_code ON subject_categories(code);
CREATE INDEX idx_subject_categories_name ON subject_categories(name);
CREATE INDEX idx_subject_categories_status ON subject_categories(status);

-- Table: graduate_profile_dimensions
CREATE TABLE graduate_profile_dimensions (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    code VARCHAR(50) NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL UNIQUE,
    name_en VARCHAR(100),
    description TEXT NOT NULL,
    indicators JSONB NOT NULL,
    weight DECIMAL(5,4) NOT NULL CHECK (weight > 0 AND weight <= 1.0),
    status VARCHAR(20) NOT NULL DEFAULT 'ACTIVE' CHECK (status IN ('ACTIVE', 'INACTIVE')),
    created_by UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_graduate_profile_dimensions_created_by FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE RESTRICT
);

CREATE INDEX idx_graduate_profile_dimensions_code ON graduate_profile_dimensions(code);
CREATE INDEX idx_graduate_profile_dimensions_name ON graduate_profile_dimensions(name);
CREATE INDEX idx_graduate_profile_dimensions_status ON graduate_profile_dimensions(status);
CREATE INDEX idx_graduate_profile_dimensions_indicators ON graduate_profile_dimensions USING GIN (indicators);

-- Table: cp_alignments
CREATE TABLE cp_alignments (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    cp_id UUID NOT NULL,
    dimension_id UUID NOT NULL,
    alignment_strength VARCHAR(20) NOT NULL CHECK (alignment_strength IN ('STRONG', 'MEDIUM', 'WEAK')),
    rationale TEXT,
    created_by UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_cp_alignments_cp_id FOREIGN KEY (cp_id) REFERENCES cp(id) ON DELETE CASCADE,
    CONSTRAINT fk_cp_alignments_dimension_id FOREIGN KEY (dimension_id) REFERENCES graduate_profile_dimensions(id) ON DELETE CASCADE,
    CONSTRAINT fk_cp_alignments_created_by FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE RESTRICT,
    CONSTRAINT uq_cp_alignments_cp_dimension UNIQUE (cp_id, dimension_id)
);

CREATE INDEX idx_cp_alignments_cp_id ON cp_alignments(cp_id);
CREATE INDEX idx_cp_alignments_dimension_id ON cp_alignments(dimension_id);
CREATE INDEX idx_cp_alignments_strength ON cp_alignments(alignment_strength);

-- ============================================================================
-- TABLE EXTENSIONS
-- ============================================================================

-- Extend: curriculum_subjects
ALTER TABLE curriculum_subjects
ADD COLUMN subject_category_id UUID;

ALTER TABLE curriculum_subjects
ADD CONSTRAINT fk_curriculum_subjects_category_id 
FOREIGN KEY (subject_category_id) REFERENCES subject_categories(id) ON DELETE SET NULL;

CREATE INDEX idx_curriculum_subjects_category_id ON curriculum_subjects(subject_category_id);

-- Extend: cp
ALTER TABLE cp
ADD COLUMN academic_year_id UUID;

ALTER TABLE cp
ADD CONSTRAINT fk_cp_academic_year_id 
FOREIGN KEY (academic_year_id) REFERENCES academic_years(id) ON DELETE SET NULL;

CREATE INDEX idx_cp_academic_year_id ON cp(academic_year_id);

ALTER TABLE cp
ADD COLUMN semester_id UUID;

ALTER TABLE cp
ADD CONSTRAINT fk_cp_semester_id 
FOREIGN KEY (semester_id) REFERENCES semesters(id) ON DELETE SET NULL;

CREATE INDEX idx_cp_semester_id ON cp(semester_id);
CREATE INDEX idx_cp_academic_semester ON cp(academic_year_id, semester_id);

-- ============================================================================
-- DATA MIGRATION
-- ============================================================================

-- Seed subject categories
INSERT INTO subject_categories (id, code, name, name_en, description, guidelines, status, created_by, created_at, updated_at) VALUES
('550e8400-e29b-41d4-a716-446655440010', 'INTRAKURIKULER', 'Intrakurikuler', 'Intracurricular', 'Mata pelajaran utama kurikulum', 'Mata pelajaran inti yang wajib diikuti semua siswa sesuai fase', 'ACTIVE', (SELECT id FROM users WHERE role = 'SYSTEM_ADMIN' LIMIT 1), NOW(), NOW()),
('550e8400-e29b-41d4-a716-446655440011', 'KOKURIKULER', 'Kokurikuler', 'Cocurricular', 'Kegiatan pengembangan diri', 'Kegiatan untuk mengembangkan potensi dan bakat siswa', 'ACTIVE', (SELECT id FROM users WHERE role = 'SYSTEM_ADMIN' LIMIT 1), NOW(), NOW()),
('550e8400-e29b-41d4-a716-446655440012', 'EKSTRAKURIKULER', 'Ekstrakurikuler', 'Extracurricular', 'Kegiatan tambahan di luar kurikulum', 'Kegiatan pilihan untuk pengembangan minat dan bakat', 'ACTIVE', (SELECT id FROM users WHERE role = 'SYSTEM_ADMIN' LIMIT 1), NOW(), NOW());

-- Categorize existing subjects as Intrakurikuler by default
UPDATE curriculum_subjects
SET subject_category_id = '550e8400-e29b-41d4-a716-446655440010'
WHERE subject_category_id IS NULL;

-- Seed graduate profile dimensions (8 dimensions)
INSERT INTO graduate_profile_dimensions (id, code, name, name_en, description, indicators, weight, status, created_by, created_at, updated_at) VALUES
('550e8400-e29b-41d4-a716-446655440020', 'KEIMANAN_KETAKWAAN', 'Keimanan & Ketakwaan', 'Faith and Piety', 'Dimensi keimanan dan ketakwaan kepada Tuhan Yang Maha Esa', '["Berakhlak mulia", "Menjaga kebersihan hati", "Melaksanakan ibadah"]'::jsonb, 0.125, 'ACTIVE', (SELECT id FROM users WHERE role = 'SYSTEM_ADMIN' LIMIT 1), NOW(), NOW()),
('550e8400-e29b-41d4-a716-446655440021', 'KEWARGAAN', 'Kewargaan', 'Citizenship', 'Dimensi kesadaran berbangsa dan bernegara', '["Cinta tanah air", "Menghargai keberagaman", "Taats aturan"]'::jsonb, 0.125, 'ACTIVE', (SELECT id FROM users WHERE role = 'SYSTEM_ADMIN' LIMIT 1), NOW(), NOW()),
('550e8400-e29b-41d4-a716-446655440022', 'BERAKHLAK_MULIA', 'Berakhlak Mulia', 'Noble Character', 'Dimensi pembentukan karakter mulia', '["Jujur", "Disiplin", "Tanggung jawab"]'::jsonb, 0.125, 'ACTIVE', (SELECT id FROM users WHERE role = 'SYSTEM_ADMIN' LIMIT 1), NOW(), NOW()),
('550e8400-e29b-41d4-a716-446655440023', 'BERANI_BERTANGGUNG_JAWAB', 'Berani Bertanggung Jawab', 'Courageous and Responsible', 'Dimensi keberanian dan tanggung jawab', '["Berani mengambil keputusan", "Pertanggung jawaban atas tindakan"]'::jsonb, 0.125, 'ACTIVE', (SELECT id FROM users WHERE role = 'SYSTEM_ADMIN' LIMIT 1), NOW(), NOW()),
('550e8400-e29b-41d4-a716-446655440024', 'PEDULI', 'Peduli', 'Caring', 'Dimensi kepedulian terhadap sesama', '["Empati", "Saling membantu", "Toleransi"]'::jsonb, 0.125, 'ACTIVE', (SELECT id FROM users WHERE role = 'SYSTEM_ADMIN' LIMIT 1), NOW(), NOW()),
('550e8400-e29b-41d4-a716-446655440025', 'GOTONG_ROYONG', 'Gotong Royong', 'Collaboration', 'Dimensi kerja sama dan gotong royong', '["Kerja tim", "Solidaritas", "Kolaborasi"]'::jsonb, 0.125, 'ACTIVE', (SELECT id FROM users WHERE role = 'SYSTEM_ADMIN' LIMIT 1), NOW(), NOW()),
('550e8400-e29b-41d4-a716-446655440026', 'MANDIRI', 'Mandiri', 'Independent', 'Dimensi kemandirian dan otonomi', '["Berpikir kritis", "Mengambil inisiatif", "Mandiri belajar"]'::jsonb, 0.125, 'ACTIVE', (SELECT id FROM users WHERE role = 'SYSTEM_ADMIN' LIMIT 1), NOW(), NOW()),
('550e8400-e29b-41d4-a716-446655440027', 'KREATIF', 'Kreatif', 'Creative', 'Dimensi kreativitas dan inovasi', '["Berpikir kreatif", "Inovasi", "Pemecahan masalah"]'::jsonb, 0.125, 'ACTIVE', (SELECT id FROM users WHERE role = 'SYSTEM_ADMIN' LIMIT 1), NOW(), NOW());
```

### Migration File: 000010_sprint4_academic_foundation.down.sql

```sql
-- Sprint 4 Academic Foundation Rollback
-- Risk Level: MEDIUM (data loss if not backed up)

-- Drop indexes for table extensions
DROP INDEX IF EXISTS idx_cp_academic_semester;
DROP INDEX IF EXISTS idx_cp_semester_id;
DROP INDEX IF EXISTS idx_cp_academic_year_id;

-- Drop foreign keys for table extensions
ALTER TABLE cp DROP CONSTRAINT IF EXISTS fk_cp_semester_id;
ALTER TABLE cp DROP CONSTRAINT IF EXISTS fk_cp_academic_year_id;

-- Drop columns for table extensions
ALTER TABLE cp DROP COLUMN IF EXISTS semester_id;
ALTER TABLE cp DROP COLUMN IF EXISTS academic_year_id;

DROP INDEX IF EXISTS idx_curriculum_subjects_category_id;
ALTER TABLE curriculum_subjects DROP CONSTRAINT IF EXISTS fk_curriculum_subjects_category_id;
ALTER TABLE curriculum_subjects DROP COLUMN IF EXISTS subject_category_id;

-- Drop indexes for new tables
DROP INDEX IF EXISTS idx_cp_alignments_strength;
DROP INDEX IF EXISTS idx_cp_alignments_dimension_id;
DROP INDEX IF EXISTS idx_cp_alignments_cp_id;

DROP INDEX IF EXISTS idx_graduate_profile_dimensions_indicators;
DROP INDEX IF EXISTS idx_graduate_profile_dimensions_status;
DROP INDEX IF EXISTS idx_graduate_profile_dimensions_name;
DROP INDEX IF EXISTS idx_graduate_profile_dimensions_code;

DROP INDEX IF EXISTS idx_subject_categories_status;
DROP INDEX IF EXISTS idx_subject_categories_name;
DROP INDEX IF EXISTS idx_subject_categories_code;

DROP INDEX IF EXISTS idx_semesters_status;
DROP INDEX IF EXISTS idx_semesters_dates;
DROP INDEX IF EXISTS idx_semesters_sequence;
DROP INDEX IF EXISTS idx_semesters_academic_year_id;

DROP INDEX IF EXISTS idx_academic_years_school_status;
DROP INDEX IF EXISTS idx_academic_years_dates;
DROP INDEX IF EXISTS idx_academic_years_status;
DROP INDEX IF EXISTS idx_academic_years_school_id;

-- Drop new tables
DROP TABLE IF EXISTS cp_alignments;
DROP TABLE IF EXISTS graduate_profile_dimensions;
DROP TABLE IF EXISTS subject_categories;
DROP TABLE IF EXISTS semesters;
DROP TABLE IF EXISTS academic_years;
```

---

## Existing Table Reuse Validation

### curriculum_phases

**Decision**: REUSE WITHOUT EXTENSION

**Rationale**:
- Existing `curriculum_phases` table already supports Fase A-E structure
- Grade level ranges (grade_level_start, grade_level_end) are sufficient
- No additional fields required for Sprint 4
- Kurikulum 2026 phase structure is already supported

**Validation**: ✅ Existing structure is sufficient

---

### curriculum_subjects

**Decision**: EXTEND WITH subject_category_id

**Rationale**:
- Existing table has all required fields for subjects
- Missing categorization (Intrakurikuler/Kokurikuler/Ekstrakurikuler)
- Extension is minimal (single nullable column)
- Backward compatible (nullable allows existing records)

**Extension**: Add `subject_category_id` FK to `subject_categories` table

**Validation**: ✅ Extension is minimal and safe

---

### cp

**Decision**: EXTEND WITH academic_year_id AND semester_id

**Rationale**:
- Existing CP table has comprehensive curriculum structure
- Missing temporal scoping (academic year, semester)
- Extensions enable future features (reporting by academic year, curriculum planning per semester)
- Backward compatible (nullable columns allow existing CP)

**Extensions**: 
- Add `academic_year_id` FK to `academic_years` table
- Add `semester_id` FK to `semesters` table

**Validation**: ✅ Extensions enable future capabilities while maintaining backward compatibility

---

# PART 6 – API DESIGN

## Academic Year API

### Endpoints

#### 1. Create Academic Year

**Method**: `POST`  
**URL**: `/api/v1/academic/academic-years`  
**Authorization**: `SCHOOL_ADMIN` or `SYSTEM_ADMIN`  
**Permission**: `academic_year:CREATE`

**Request**:
```json
{
  "school_id": "uuid",
  "name": "2026/2027",
  "start_date": "2026-07-15T00:00:00Z",
  "end_date": "2027-06-30T23:59:59Z",
  "description": "Tahun ajaran 2026/2027"
}
```

**Response**: `201 Created`
```json
{
  "id": "uuid",
  "school_id": "uuid",
  "name": "2026/2027",
  "start_date": "2026-07-15T00:00:00Z",
  "end_date": "2027-06-30T23:59:59Z",
  "status": "DRAFT",
  "approved_by": null,
  "approved_at": null,
  "created_by": "uuid",
  "created_at": "2026-06-11T10:00:00Z",
  "updated_at": "2026-06-11T10:00:00Z"
}
```

**Validation**:
- `school_id`: Required, must exist and belong to user's school
- `name`: Required, max 100 characters, unique per school
- `start_date`: Required, valid date, >= today + 30 days
- `end_date`: Required, valid date, > start_date
- `description`: Optional, max 500 characters

**Error Codes**:
- `400 BAD_REQUEST`: Validation error
- `403 FORBIDDEN`: Insufficient permissions
- `409 CONFLICT`: Academic year name already exists for school
- `409 CONFLICT`: Date range overlaps with existing academic year
- `422 UNPROCESSABLE_ENTITY`: Start date less than 30 days in future

---

#### 2. List Academic Years

**Method**: `GET`  
**URL**: `/api/v1/academic/academic-years`  
**Authorization**: `SCHOOL_ADMIN`, `SYSTEM_ADMIN`, `CURRICULUM_ADMIN`, `TEACHER`  
**Permission**: `academic_year:READ`

**Query Parameters**:
- `school_id` (optional): Filter by school
- `status` (optional): Filter by status
- `year` (optional): Filter by year (e.g., "2026")

**Response**: `200 OK`
```json
{
  "data": [
    {
      "id": "uuid",
      "school_id": "uuid",
      "school_name": "SD NUSA 01",
      "name": "2026/2027",
      "start_date": "2026-07-15T00:00:00Z",
      "end_date": "2027-06-30T23:59:59Z",
      "status": "ACTIVE",
      "approved_by": "uuid",
      "approved_by_name": "Admin NUSA",
      "approved_at": "2026-06-01T10:00:00Z",
      "created_at": "2026-05-15T10:00:00Z",
      "semesters_count": 2
    }
  ],
  "pagination": {
    "page": 1,
    "per_page": 20,
    "total": 10,
    "total_pages": 1
  }
}
```

**Error Codes**:
- `403 FORBIDDEN`: Insufficient permissions

---

#### 3. Get Academic Year by ID

**Method**: `GET`  
**URL**: `/api/v1/academic/academic-years/:id`  
**Authorization**: `SCHOOL_ADMIN`, `SYSTEM_ADMIN`, `CURRICULUM_ADMIN`, `TEACHER`  
**Permission**: `academic_year:READ`

**Response**: `200 OK`
```json
{
  "id": "uuid",
  "school_id": "uuid",
  "school_name": "SD NUSA 01",
  "name": "2026/2027",
  "start_date": "2026-07-15T00:00:00Z",
  "end_date": "2027-06-30T23:59:59Z",
  "status": "ACTIVE",
  "approved_by": "uuid",
  "approved_by_name": "Admin NUSA",
  "approved_at": "2026-06-01T10:00:00Z",
  "created_by": "uuid",
  "created_by_name": "Admin NUSA",
  "created_at": "2026-05-15T10:00:00Z",
  "updated_at": "2026-06-01T10:00:00Z",
  "semesters": [
    {
      "id": "uuid",
      "name": "Semester Ganjil",
      "sequence": 1,
      "start_date": "2026-07-15T00:00:00Z",
      "end_date": "2026-12-31T23:59:59Z",
      "status": "ACTIVE"
    },
    {
      "id": "uuid",
      "name": "Semester Genap",
      "sequence": 2,
      "start_date": "2027-01-01T00:00:00Z",
      "end_date": "2027-06-30T23:59:59Z",
      "status": "INACTIVE"
    }
  ]
}
```

**Error Codes**:
- `403 FORBIDDEN`: Insufficient permissions or school access
- `404 NOT FOUND`: Academic year not found

---

#### 4. Update Academic Year

**Method**: `PUT`  
**URL**: `/api/v1/academic/academic-years/:id`  
**Authorization**: `SCHOOL_ADMIN` or `SYSTEM_ADMIN`  
**Permission**: `academic_year:UPDATE`

**Request**:
```json
{
  "name": "2026/2027 (Updated)",
  "start_date": "2026-07-15T00:00:00Z",
  "end_date": "2027-06-30T23:59:59Z",
  "description": "Updated description"
}
```

**Response**: `200 OK`
```json
{
  "id": "uuid",
  "name": "2026/2027 (Updated)",
  "start_date": "2026-07-15T00:00:00Z",
  "end_date": "2027-06-30T23:59:59Z",
  "status": "DRAFT",
  "updated_at": "2026-06-11T11:00:00Z"
}
```

**Validation**:
- Can only update academic years in DRAFT or UNDER_REVIEW status
- Date range validation same as create
- Name uniqueness validation

**Error Codes**:
- `400 BAD_REQUEST`: Validation error
- `403 FORBIDDEN`: Insufficient permissions
- `404 NOT_FOUND`: Academic year not found
- `409 CONFLICT`: Cannot update active or past academic years

---

#### 5. Submit for Approval

**Method**: `POST`  
**URL**: `/api/v1/academic/academic-years/:id/submit`  
**Authorization**: `SCHOOL_ADMIN`  
**Permission**: `academic_year:SUBMIT`

**Request**: (empty body)

**Response**: `200 OK`
```json
{
  "id": "uuid",
  "status": "UNDER_REVIEW",
  "updated_at": "2026-06-11T11:00:00Z"
}
```

**Validation**:
- Academic year must have exactly 2 semesters configured
- Academic year must be in DRAFT status
- All required fields must be populated

**Error Codes**:
- `400 BAD_REQUEST`: Academic year not ready for approval
- `403 FORBIDDEN`: Insufficient permissions
- `404 NOT_FOUND`: Academic year not found

---

#### 6. Approve Academic Year

**Method**: `POST`  
**URL**: `/api/v1/academic/academic-years/:id/approve`  
**Authorization**: `SYSTEM_ADMIN`  
**Permission**: `academic_year:APPROVE`

**Request**:
```json
{
  "reason": "Configuration approved for 2026/2027"
}
```

**Response**: `200 OK`
```json
{
  "id": "uuid",
  "status": "APPROVED",
  "approved_by": "uuid",
  "approved_at": "2026-06-11T11:00:00Z",
  "updated_at": "2026-06-11T11:00:00Z"
}
```

**Validation**:
- Academic year must be in UNDER_REVIEW status
- Only SYSTEM_ADMIN can approve

**Error Codes**:
- `400 BAD_REQUEST`: Academic year not in UNDER_REVIEW status
- `403 FORBIDDEN`: Insufficient permissions (not SYSTEM_ADMIN)
- `404 NOT_FOUND`: Academic year not found

---

#### 7. Deactivate Academic Year

**Method**: `POST`  
**URL`: `/api/v1/academic/academic-years/:id/deactivate`  
**Authorization**: `SCHOOL_ADMIN` or `SYSTEM_ADMIN`  
**Permission**: `academic_year:DEACTIVATE`

**Request**:
```json
{
  "reason": "Year end completed"
}
```

**Response**: `200 OK`
```json
{
  "id": "uuid",
  "status": "INACTIVE",
  "updated_at": "2026-06-11T11:00:00Z"
}
```

**Validation**:
- Academic year must be ACTIVE status
- Must provide reason
- Cannot deactivate if active curriculum planning exists

**Error Codes**:
- `400 BAD_REQUEST`: Cannot deactivate (has active dependencies)
- `403 FORBIDDEN`: Insufficient permissions
- `404 NOT_FOUND`: Academic year not found

---

## Semester API

### Endpoints

#### 1. Create Semester

**Method**: `POST`  
**URL**: `/api/v1/academic/semesters`  
**Authorization**: `SCHOOL_ADMIN` or `SYSTEM_ADMIN`  
**Permission**: `semester:CREATE`

**Request**:
```json
{
  "academic_year_id": "uuid",
  "name": "Semester Ganjil",
  "sequence": 1,
  "start_date": "2026-07-15T00:00:00Z",
  "end_date": "2026-12-31T23:59:59Z",
  "description": "Semester ganjil 2026/2027"
}
```

**Response**: `201 Created`
```json
{
  "id": "uuid",
  "academic_year_id": "uuid",
  "name": "Semester Ganjil",
  "sequence": 1,
  "start_date": "2026-07-15T00:00:00Z",
  "end_date": "2026-12-31T23:59:59Z",
  "status": "DRAFT",
  "created_at": "2026-06-11T10:00:00Z",
  "updated_at": "2026-06-11T10:00:00Z"
}
```

**Error Codes**:
- `400 BAD_REQUEST`: Validation error
- `403 FORBIDDEN`: Insufficient permissions
- `404 NOT_FOUND`: Academic year not found
- `409 CONFLICT`: Sequence conflict or date overlap

---

#### 2. List Semesters

**Method**: `GET`  
**URL**: `/api/v1/academic/semesters`  
**Authorization**: `SCHOOL_ADMIN`, `SYSTEM_ADMIN`, `CURRICULUM_ADMIN`, `TEACHER`  
**Permission**: `semester:READ`

**Query Parameters**:
- `academic_year_id` (optional): Filter by academic year
- `status` (optional): Filter by status

**Response**: `200 OK`
```json
{
  "data": [
    {
      "id": "uuid",
      "academic_year_id": "uuid",
      "academic_year_name": "2026/2027",
      "name": "Semester Ganjil",
      "sequence": 1,
      "start_date": "2026-07-15T00:00:00Z",
      "end_date": "2026-12-31T23:59:59Z",
      "status": "ACTIVE"
    }
  ]
}
```

**Error Codes**:
- `403 FORBIDDEN`: Insufficient permissions

---

#### 3. Update Semester

**Method**: `PUT`  
**URL**: `/api/v1/academic/semesters/:id`  
**Authorization**: `SCHOOL_ADMIN` or `SYSTEM_ADMIN`  
**Permission**: `semester:UPDATE`

**Request**:
```json
{
  "name": "Semester Ganjil (Updated)",
  "start_date": "2026-07-15T00:00:00Z",
  "end_date": "2026-12-31T23:59:59Z"
}
```

**Response**: `200 OK`
```json
{
  "id": "uuid",
  "name": "Semester Ganjil (Updated)",
  "updated_at": "2026-06-11T11:00:00Z"
}
```

**Error Codes**:
- `400 BAD_REQUEST`: Validation error
- `403 FORBIDDEN`: Insufficient permissions
- `404 NOT_FOUND`: Semester not found
- `409 CONFLICT`: Cannot update active semester

---

#### 4. Delete Semester

**Method**: `DELETE`  
**URL**: `/api/v1/academic/semesters/:id`  
**Authorization**: `SCHOOL_ADMIN` or `SYSTEM_ADMIN`  
**Permission**: `semester:DELETE`

**Response**: `204 No Content`

**Validation**:
- Can only delete semesters in DRAFT status
- Cannot delete if linked to CP or other entities

**Error Codes**:
- `400 BAD_REQUEST`: Cannot delete (has dependencies)
- `403 FORBIDDEN`: Insufficient permissions
- `404 NOT_FOUND`: Semester not found

---

## Subject Category API

### Endpoints

#### 1. Create Subject Category

**Method**: `POST`  
**URL**: `/api/v1/curriculum/subject-categories`  
**Authorization**: `CURRICULUM_ADMIN`  
**Permission**: `subject_category:CREATE`

**Request**:
```json
{
  "code": "INTRAKURIKULER",
  "name": "Intrakurikuler",
  "name_en": "Intracurricular",
  "description": "Mata pelajaran utama kurikulum",
  "guidelines": "Mata pelajaran inti yang wajib diikuti semua siswa sesuai fase"
}
```

**Response**: `201 Created`
```json
{
  "id": "uuid",
  "code": "INTRAKURIKULER",
  "name": "Intrakurikuler",
  "name_en": "Intracurricular",
  "description": "Mata pelajaran utama kurikulum",
  "guidelines": "Mata pelajaran inti yang wajib diikuti semua siswa sesuai fase",
  "status": "ACTIVE",
  "created_by": "uuid",
  "created_by_name": "Admin NUSA",
  "created_at": "2026-06-11T10:00:00Z",
  "updated_at": "2026-06-11T10:00:00Z"
}
```

**Error Codes**:
- `400 BAD_REQUEST`: Validation error
- `403 FORBIDDEN`: Insufficient permissions
- `409 CONFLICT`: Code or name already exists

---

#### 2. List Subject Categories

**Method**: `GET`  
**URL**: `/api/v1/curriculum/subject-categories`  
**Authorization**: `SCHOOL_ADMIN`, `SYSTEM_ADMIN`, `CURRICULUM_ADMIN`, `TEACHER`  
**Permission**: `subject_category:READ`

**Query Parameters**:
- `status` (optional): Filter by status

**Response**: `200 OK`
```json
{
  "data": [
    {
      "id": "uuid",
      "code": "INTRAKURIKULER",
      "name": "Intrakurikuler",
      "name_en": "Intracurricular",
      "description": "Mata pelajaran utama kurikulum",
      "guidelines": "Mata pelajaran inti yang wajib diikuti semua siswa sesuai fase",
      "status": "ACTIVE",
      "subjects_count": 15,
      "created_at": "2026-06-11T10:00:00Z"
    }
  ]
}
```

**Error Codes**:
- `403 FORBIDDEN`: Insufficient permissions

---

#### 3. Update Subject Category

**Method**: `PUT`  
**URL**: `/api/v1/curriculum/subject-categories/:id`  
**Authorization**: `CURRICULUM_ADMIN`  
**Permission**: `subject_category:UPDATE`

**Request**:
```json
{
  "name": "Intrakurikuler (Updated)",
  "description": "Updated description"
}
```

**Response**: `200 OK`
```json
{
  "id": "uuid",
  "name": "Intrakurikuler (Updated)",
  "description": "Updated description",
  "updated_at": "2026-06-11T11:00:00Z"
}
```

**Error Codes**:
- `400 BAD_REQUEST`: Validation error
- `403 FORBIDDEN`: Insufficient permissions
- `404 NOT_FOUND`: Category not found

---

#### 4. Deactivate Subject Category

**Method**: `POST`  
**URL**: `/api/v1/curriculum/subject-categories/:id/deactivate`  
**Authorization**: `CURRICULUM_ADMIN`  
**Permission**: `subject_category:DEACTIVATE`

**Request**:
```json
{
  "reason": "No longer used"
}
```

**Response**: `200 OK`
```json
{
  "id": "uuid",
  "status": "INACTIVE",
  "updated_at": "2026-06-11T11:00:00Z"
}
```

**Validation**:
- Cannot deactivate if subjects are assigned

**Error Codes**:
- `400 BAD_REQUEST`: Cannot deactivate (has assigned subjects)
- `403 FORBIDDEN`: Insufficient permissions
- `404 NOT_FOUND`: Category not found

---

## Graduate Profile Dimension API

### Endpoints

#### 1. Create Graduate Profile Dimension

**Method**: `POST`  
**URL**: `/api/v1/curriculum/graduate-profile-dimensions`  
**Authorization**: `CURRICULUM_ADMIN`  
**Permission**: `graduate_profile_dimension:CREATE`

**Request**:
```json
{
  "code": "KEIMANAN_KETAKWAAN",
  "name": "Keimanan & Ketakwaan",
  "name_en": "Faith and Piety",
  "description": "Dimensi keimanan dan ketakwaan kepada Tuhan Yang Maha Esa",
  "indicators": ["Berakhlak mulia", "Menjaga kebersihan hati", "Melaksanakan ibadah"],
  "weight": 0.125
}
```

**Response**: `201 Created`
```json
{
  "id": "uuid",
  "code": "KEIMANAN_KETAKWAAN",
  "name": "Keimanan & Ketakwaan",
  "name_en": "Faith and Piety",
  "description": "Dimensi keimanan dan ketakwaan kepada Tuhan Yang Maha Esa",
  "indicators": ["Berakhlak mulia", "Menjaga kebersihan hati", "Melaksanakan ibadah"],
  "weight": 0.125,
  "status": "ACTIVE",
  "created_by": "uuid",
  "created_by_name": "Admin NUSA",
  "created_at": "2026-06-11T10:00:00Z",
  "updated_at": "2026-06-11T10:00:00Z"
}
```

**Error Codes**:
- `400 BAD_REQUEST`: Validation error
- `403 FORBIDDEN`: Insufficient permissions
- `409 CONFLICT`: Code or name already exists
- `422 UNPROCESSABLE ENTITY`: More than 8 active dimensions

---

#### 2. List Graduate Profile Dimensions

**Method**: `GET`  
**URL**: `/api/v1/curriculum/graduate-profile-dimensions`  
**Authorization**: `SCHOOL_ADMIN`, `SYSTEM_ADMIN`, `CURRICULUM_ADMIN`, `TEACHER`  
**Permission**: `graduate_profile_dimension:READ`

**Query Parameters**:
- `status` (optional): Filter by status

**Response**: `200 OK`
```json
{
  "data": [
    {
      "id": "uuid",
      "code": "KEIMANAN_KETAKWAAN",
      "name": "Keimanan & Ketakwaan",
      "name_en": "Faith and Piety",
      "description": "Dimensi keimanan dan ketakwaan kepada Tuhan Yang Maha Esa",
      "indicators": ["Berakhlak mulia", "Menjaga kebersihan hati", "Melaksanakan ibadah"],
      "weight": 0.125,
      "status": "ACTIVE",
      "cp_alignments_count": 45,
      "created_at": "2026-06-11T10:00:00Z"
    }
  ]
}
```

**Error Codes**:
- `403 FORBIDDEN`: Insufficient permissions

---

#### 3. Update Graduate Profile Dimension

**Method**: `PUT`  
**URL**: `/api/v1/curriculum/graduate-profile-dimensions/:id`  
**Authorization**: `CURRICULUM_ADMIN`  
**Permission**: `graduate_profile_dimension:UPDATE`

**Request**:
```json
{
  "name": "Keimanan & Ketakwaan (Updated)",
  "description": "Updated description",
  "indicators": ["Updated indicator 1", "Updated indicator 2"],
  "weight": 0.125
}
```

**Response**: `200 OK`
```json
{
  "id": "uuid",
  "name": "Keimanan & Ketakwaan (Updated)",
  "description": "Updated description",
  "indicators": ["Updated indicator 1", "Updated indicator 2"],
  "weight": 0.125,
  "updated_at": "2026-06-11T11:00:00Z"
}
```

**Error Codes**:
- `400 BAD_REQUEST`: Validation error
- `403 FORBIDDEN`: Insufficient permissions
- `404 NOT_FOUND`: Dimension not found

---

#### 4. Deactivate Graduate Profile Dimension

**Method**: `POST`  
**URL**: `/api/v1/curriculum/graduate-profile-dimensions/:id/deactivate`  
**Authorization**: `CURRICULUM_ADMIN`  
**Permission**: `graduate_profile_dimension:DEACTIVATE`

**Request**:
```json
{
  "reason": "Dimension retired"
}
```

**Response**: `200 OK`
```json
{
  "id": "uuid",
  "status": "INACTIVE",
  "updated_at": "2026-06-11T11:00:00Z"
}
```

**Validation**:
- Cannot deactivate if CP are aligned
- Cannot deactivate if it would result in less than 8 active dimensions

**Error Codes**:
- `400 BAD_REQUEST`: Cannot deactivate (has aligned CP or would violate 8-dimension rule)
- `403 FORBIDDEN`: Insufficient permissions
- `404 NOT_FOUND`: Dimension not found

---

## Curriculum Alignment API

### Endpoints

#### 1. Create CP Alignment

**Method**: `POST`  
**URL**: `/api/v1/curriculum/cp-alignments`  
**Authorization**: `CURRICULUM_ADMIN`  
**Permission**: `cp_alignment:CREATE`

**Request**:
```json
{
  "cp_id": "uuid",
  "dimension_id": "uuid",
  "alignment_strength": "STRONG",
  "rationale": "Strong alignment to faith dimension"
}
```

**Response**: `201 Created`
```json
{
  "id": "uuid",
  "cp_id": "uuid",
  "cp_code": "CP.001",
  "cp_description": "Capaian pembelajaran keimanan",
  "dimension_id": "uuid",
  "dimension_code": "KEIMANAN_KETAKWAAN",
  "dimension_name": "Keimanan & Ketakwaan",
  "alignment_strength": "STRONG",
  "rationale": "Strong alignment to faith dimension",
  "created_by": "uuid",
  "created_by_name": "Admin NUSA",
  "created_at": "2026-06-11T10:00:00Z",
  "updated_at": "2026-06-11T10:00:00Z"
}
```

**Error Codes**:
- `400 BAD_REQUEST`: Validation error
- `403 FORBIDDEN`: Insufficient permissions
- `404 NOT_FOUND`: CP or dimension not found
- `409 CONFLICT`: Alignment already exists for CP-dimension pair

---

#### 2. List CP Alignments

**Method**: `GET`  
**URL**: `/api/v1/curriculum/cp-alignments`  
**Authorization**: `SCHOOL_ADMIN`, `SYSTEM_ADMIN`, `CURRICULUM_ADMIN`, `TEACHER`  
**Permission**: `cp_alignment:READ`

**Query Parameters**:
- `cp_id` (optional): Filter by CP
- `dimension_id` (optional): Filter by dimension
- `alignment_strength` (optional): Filter by strength

**Response**: `200 OK`
```json
{
  "data": [
    {
      "id": "uuid",
      "cp_id": "uuid",
      "cp_code": "CP.001",
      "cp_description": "Capaian pembelajaran keimanan",
      "dimension_id": "uuid",
      "dimension_code": "KEIMANAN_KETAKWAAN",
      "dimension_name": "Keimanan & Ketakwaan",
      "alignment_strength": "STRONG",
      "rationale": "Strong alignment to faith dimension",
      "created_at": "2026-06-11T10:00:00Z"
    }
  ]
}
```

**Error Codes**:
- `403 FORBIDDEN`: Insufficient permissions

---

#### 3. Update CP Alignment

**Method**: `PUT`  
**URL**: `/api/v1/curriculum/cp-alignments/:id`  
**Authorization**: `CURRICULUM_ADMIN`  
**Permission**: `cp_alignment:UPDATE`

**Request**:
```json
{
  "alignment_strength": "MEDIUM",
  "rationale": "Updated rationale"
}
```

**Response**: `200 OK`
```json
{
  "id": "uuid",
  "alignment_strength": "MEDIUM",
  "rationale": "Updated rationale",
  "updated_at": "2026-06-11T11:00:00Z"
}
```

**Error Codes**:
- `400 BAD_REQUEST`: Validation error
- `403 FORBIDDEN`: Insufficient permissions
- `404 NOT_FOUND`: Alignment not found

---

#### 4. Delete CP Alignment

**Method**: `DELETE`  
**URL**: `/api/v1/curriculum/cp-alignments/:id`  
**Authorization**: `CURRICULUM_ADMIN`  
**Permission**: `cp_alignment:DELETE`

**Response**: `204 No Content`

**Validation**:
- Warn if CP will fall below 60% alignment threshold

**Error Codes**:
- `400 BAD_REQUEST`: Cannot delete (would violate 60% threshold)
- `403 FORBIDDEN`: Insufficient permissions
- `404 NOT_FOUND`: Alignment not found

---

#### 5. Get Alignment Report

**Method**: `GET`  
**URL**: `/api/v1/curriculum/alignment-report`  
**Authorization**: `CURRICULUM_ADMIN`, `SCHOOL_ADMIN`  
**Permission**: `alignment_report:READ`

**Query Parameters**:
- `academic_year_id` (optional): Filter by academic year
- `phase_id` (optional): Filter by phase
- `subject_id` (optional): Filter by subject

**Response**: `200 OK`
```json
{
  "summary": {
    "total_cp": 150,
    "aligned_cp": 142,
    "overall_alignment_percentage": 94.7,
    "below_threshold_count": 3
  },
  "by_dimension": [
    {
      "dimension_id": "uuid",
      "dimension_code": "KEIMANAN_KETAKWAAN",
      "dimension_name": "Keimanan & Ketakwaan",
      "cp_count": 18,
      "average_alignment": 92.5
    }
  ],
  "below_threshold_cp": [
    {
      "cp_id": "uuid",
      "cp_code": "CP.045",
      "cp_description": "Description",
      "alignment_percentage": 55.0,
      "alignments_count": 1
    }
  ]
}
```

**Error Codes**:
- `403 FORBIDDEN`: Insufficient permissions

---

# PART 7 – FRONTEND REQUIREMENTS

## Pages Overview

### School Admin Pages

#### 1. Academic Year Management
**Path**: `/admin/academic/academic-years`  
**Purpose**: Manage academic years for school  
**Actors**: School Admin

**Components**:
- **Table**: List academic years with columns (Name, Start Date, End Date, Status, Actions)
- **Filters**: Status (All, Draft, Under Review, Approved, Active, Inactive)
- **Search**: Search by academic year name
- **Actions**: Create, View, Edit (if Draft), Submit (if Draft), Approve (if System Admin), Deactivate (if Active)
- **Pagination**: 20 items per page

**Form Fields (Create/Edit)**:
- Academic Year Name (text, required, max 100 chars)
- Start Date (date picker, required, >= today + 30 days)
- End Date (date picker, required, > start date)
- Description (textarea, optional, max 500 chars)

**Validation**:
- Real-time validation for date constraints
- Warning if dates overlap with existing academic years
- Disable submit if validation fails

**Permissions**: `academic_year:READ`, `academic_year:CREATE`, `academic_year:UPDATE`, `academic_year:SUBMIT`, `academic_year:DEACTIVATE`

---

#### 2. Semester Configuration
**Path**: `/admin/academic/academic-years/:id/semesters`  
**Purpose**: Configure semesters within academic year  
**Actors**: School Admin

**Components**:
- **Table**: List semesters with columns (Name, Sequence, Start Date, End Date, Status, Actions)
- **Visualization**: Timeline showing academic year with semesters
- **Actions**: Create, Edit (if Draft), Delete (if Draft)
- **Status Badges**: DRAFT (gray), ACTIVE (green), INACTIVE (red)

**Form Fields (Create/Edit)**:
- Semester Name (text, required, max 50 chars, default "Semester Ganjil"/"Semester Genap")
- Sequence (dropdown, required, options: 1, 2)
- Start Date (date picker, required, within academic year range)
- End Date (date picker, required, within academic year range)
- Description (textarea, optional, max 500 chars)

**Validation**:
- Sequence must be unique within academic year
- Dates must be within academic year range
- No gaps or overlaps between semesters
- Visual gap/overlap indicators in timeline

**Permissions**: `semester:READ`, `semester:CREATE`, `semester:UPDATE`, `semester:DELETE`

---

### Curriculum Admin Pages

#### 3. Subject Category Management
**Path**: `/admin/curriculum/subject-categories`  
**Purpose**: Manage subject categories (Intrakurikuler, Kokurikuler, Ekstrakurikuler)  
**Actors**: Curriculum Admin

**Components**:
- **Table**: List categories with columns (Code, Name, Name EN, Status, Subjects Count, Actions)
- **Filters**: Status (All, Active, Inactive)
- **Search**: Search by code or name
- **Actions**: Create, View, Edit, Deactivate
- **Stats Cards**: Total categories, Active categories, Total subjects categorized

**Form Fields (Create/Edit)**:
- Category Code (text, required, uppercase, max 50 chars)
- Category Name (text, required, max 100 chars)
- Category Name (English) (text, optional, max 100 chars)
- Description (textarea, required, max 1000 chars)
- Guidelines (textarea, optional, max 2000 chars)

**Validation**:
- Code must be uppercase and unique
- Name must be unique
- Cannot deactivate if subjects are assigned

**Permissions**: `subject_category:READ`, `subject_category:CREATE`, `subject_category:UPDATE`, `subject_category:DEACTIVATE`

---

#### 4. Subject Categorization
**Path**: `/admin/curriculum/subjects/categorization`  
**Purpose**: Categorize subjects into categories  
**Actors**: Curriculum Admin

**Components**:
- **Table**: List subjects with columns (Code, Name, Current Category, Actions)
- **Filters**: Category (All, Intrakurikuler, Kokurikuler, Ekstrakurikuler, Uncategorized)
- **Bulk Actions**: Assign category to selected subjects
- **Progress Bar**: Show categorization completion percentage

**Form Fields (Bulk Assign)**:
- Category (dropdown, required)
- Selected subjects (checkbox list)

**Validation**:
- At least one subject must be selected
- Category must be active

**Permissions**: `curriculum_subjects:UPDATE`

---

#### 5. Graduate Profile Dimensions
**Path**: `/admin/curriculum/graduate-profile-dimensions`  
**Purpose**: Manage 8-dimensional Profil Lulusan  
**Actors**: Curriculum Admin

**Components**:
- **Table**: List dimensions with columns (Code, Name, Name EN, Weight, Status, CP Alignments Count, Actions)
- **Filters**: Status (All, Active, Inactive)
- **Stats Cards**: Total dimensions (8), Active dimensions, Total CP aligned
- **Visualization**: Radar chart showing dimension coverage

**Form Fields (Create/Edit)**:
- Dimension Code (text, required, uppercase, max 50 chars)
- Dimension Name (text, required, max 100 chars)
- Dimension Name (English) (text, optional, max 100 chars)
- Description (textarea, required, max 1000 chars)
- Indicators (dynamic tags input, required, at least 1)
- Weight (number, required, 0 < weight <= 1.0, default 0.125)

**Validation**:
- Code must be uppercase and unique
- Name must be unique
- Weight must be positive and <= 1.0
- Exactly 8 dimensions can be active
- Cannot deactivate if CP are aligned

**Permissions**: `graduate_profile_dimension:READ`, `graduate_profile_dimension:CREATE`, `graduate_profile_dimension:UPDATE`, `graduate_profile_dimension:DEACTIVATE`

---

#### 6. CP Alignment Management
**Path**: `/admin/curriculum/cp-alignment`  
**Purpose**: Align CP to graduate profile dimensions  
**Actors**: Curriculum Admin

**Components**:
- **Table**: List CP with columns (Code, Description, Alignment Percentage, Dimensions Aligned, Actions)
- **Filters**: Subject, Phase, Alignment Level (All, <60%, 60-80%, >80%)
- **Search**: Search by CP code or description
- **Alignment Visual**: Progress bar showing alignment percentage
- **Bulk Actions**: Align selected CP to dimensions

**Form Fields (Align CP)**:
- Select CP (checkbox list or dropdown)
- Select Dimension(s) (multi-select dropdown)
- Alignment Strength (radio buttons: Strong, Medium, Weak)
- Rationale (textarea, optional, max 500 chars)

**Validation**:
- At least one CP must be selected
- At least one dimension must be selected
- Alignment strength must be selected
- Warn if CP will fall below 60% alignment

**Permissions**: `cp_alignment:READ`, `cp_alignment:CREATE`, `cp_alignment:UPDATE`, `cp_alignment:DELETE`

---

#### 7. Alignment Report
**Path**: `/admin/curriculum/alignment-report`  
**Purpose**: View curriculum alignment reports  
**Actors**: Curriculum Admin, School Admin

**Components**:
- **Summary Cards**: Total CP, Aligned CP, Overall Alignment %, Below Threshold Count
- **Filters**: Academic Year, Phase, Subject
- **Charts**: 
  - Bar chart: Alignment per dimension
  - Pie chart: CP alignment distribution
  - Line chart: Alignment trend over time
- **Table**: CP below 60% threshold with details
- **Actions**: Export PDF, Export CSV

**Validation**: None (read-only page)

**Permissions**: `alignment_report:READ`

---

#### 8. Koding & AI Subject Setup
**Path**: `/admin/curriculum/modern-subjects`  
**Purpose**: Add Koding, AI, Numerasi subjects  
**Actors**: Curriculum Admin

**Components**:
- **Table**: List modern subjects (Koding, AI, Numerasi) with status
- **Form**: Subject creation form for new modern subjects
- **Instructions**: Guide for adding modern subjects
- **Integration Check**: Verify subjects integrate with CP/TP workflow

**Form Fields**:
- Subject Code (text, required, uppercase, max 50 chars)
- Subject Name (text, required, max 255 chars)
- Subject Name (English) (text, optional, max 255 chars)
- Description (textarea, required, max 1000 chars)
- Category (dropdown, required, default Intrakurikuler)
- Phase (multi-select, required)
- Is Active (toggle, required)

**Validation**:
- Code must be uppercase and unique
- Name must be unique
- Category must be active
- At least one phase must be selected

**Permissions**: `curriculum_subjects:CREATE`

---

### Teacher Pages

#### 9. Academic Calendar View
**Path**: `/teacher/academic-calendar`  
**Purpose**: View academic calendar for planning  
**Actors**: Teacher

**Components**:
- **Calendar View**: Monthly calendar showing academic year, semester dates
- **Academic Year Info**: Current academic year, semester dates
- **Upcoming Events**: Important dates (semester start/end, holidays)
- **Read-only**: No editing capabilities

**Validation**: None (read-only page)

**Permissions**: `academic_year:READ`

---

#### 10. Subject Categories View
**Path**: `/teacher/curriculum/subject-categories`  
**Purpose**: View subject categories for understanding  
**Actors**: Teacher

**Components**:
- **Table**: List categories with descriptions and guidelines
- **Subject List**: Show subjects under each category
- **Read-only**: No editing capabilities

**Validation**: None (read-only page)

**Permissions**: `subject_category:READ`

---

#### 11. Graduate Profile Reference
**Path**: `/teacher/curriculum/graduate-profile`  
**Purpose**: Reference graduate profile dimensions for lesson planning  
**Actors**: Teacher

**Components**:
- **Cards**: 8 dimension cards with descriptions and indicators
- **Search**: Search dimensions by name or indicator
- **CP Alignment**: Show which CP align to which dimensions
- **Read-only**: No editing capabilities

**Validation**: None (read-only page)

**Permissions**: `graduate_profile_dimension:READ`

---

## Navigation Structure

### School Admin Navigation
```
Dashboard
├── Academic Management
│   ├── Academic Years
│   └── Semester Configuration
└── Reports
    └── Alignment Report
```

### Curriculum Admin Navigation
```
Dashboard
├── Curriculum Management
│   ├── Subjects
│   │   ├── Subject Categories
│   │   ├── Subject Categorization
│   │   └── Modern Subjects (Koding, AI, Numerasi)
│   ├── Phases
│   ├── Elements & Subelements
│   ├── CP
│   ├── Graduate Profile Dimensions
│   ├── CP Alignment
│   └── Alignment Report
```

### Teacher Navigation
```
Dashboard
├── Planning
│   ├── Academic Calendar
│   ├── Curriculum
│   │   ├── Subject Categories
│   │   └── Graduate Profile Reference
│   ├── CP
│   ├── TP
│   ├── ATP
│   └── Modul Ajar
```

---

## UI/UX Guidelines

### Consistent Patterns

1. **Data Tables**: All tables use consistent pagination (20 items), filters, search, and action menus
2. **Form Validation**: Real-time validation with inline error messages
3. **Loading States**: Skeleton loaders for all async operations
4. **Empty States**: Friendly empty state messages with call-to-action
5. **Confirmation Dialogs**: All destructive actions require confirmation
6. **Success Notifications**: Toast notifications for successful operations
7. **Error Handling**: User-friendly error messages with retry options

### Indonesian Language

All UI text must be in Indonesian:
- "Academic Year" → "Tahun Ajaran"
- "Semester" → "Semester"
- "Subject Category" → "Kategori Mata Pelajaran"
- "Graduate Profile Dimension" → "Dimensi Profil Lulusan"
- "Alignment" → "Aligment" (or "Kesesuaian")

### Accessibility

- WCAG 2.1 AA compliance
- Keyboard navigation support
- Screen reader support
- Color contrast ratio >= 4.5:1
- Focus indicators on all interactive elements

---

# PART 8 – SECURITY REQUIREMENTS

## Roles

### Existing Roles (Reuse)

- **SYSTEM_ADMIN**: Platform-level administration
- **SCHOOL_ADMIN**: School-level administration
- **TEACHER**: Curriculum planning and delivery
- **CURRICULUM_ADMIN**: (New Role) System-wide curriculum governance

### New Role: CURRICULUM_ADMIN

**Purpose**: Manage curriculum governance across all schools  
**Scope**: System-wide (not school-scoped)  
**Responsibilities**:
- Configure subject categories (system-wide)
- Configure graduate profile dimensions (system-wide)
- Manage CP alignments
- Generate alignment reports
- Approve curriculum structure changes

**Permissions**: See below

---

## Permissions

### Academic Year Permissions

| Permission | Description | Roles |
| ---------- | ----------- | ------- |
| `academic_year:READ` | View academic years | SCHOOL_ADMIN, SYSTEM_ADMIN, CURRICULUM_ADMIN, TEACHER |
| `academic_year:CREATE` | Create academic years | SCHOOL_ADMIN, SYSTEM_ADMIN |
| `academic_year:UPDATE` | Update academic years | SCHOOL_ADMIN, SYSTEM_ADMIN |
| `academic_year:DELETE` | Delete academic years | SYSTEM_ADMIN |
| `academic_year:SUBMIT` | Submit for approval | SCHOOL_ADMIN |
| `academic_year:APPROVE` | Approve academic years | SYSTEM_ADMIN |
| `academic_year:DEACTIVATE` | Deactivate academic years | SCHOOL_ADMIN, SYSTEM_ADMIN |

### Semester Permissions

| Permission | Description | Roles |
| ---------- | ----------- | ------- |
| `semester:READ` | View semesters | SCHOOL_ADMIN, SYSTEM_ADMIN, CURRICULUM_ADMIN, TEACHER |
| `semester:CREATE` | Create semesters | SCHOOL_ADMIN, SYSTEM_ADMIN |
| `semester:UPDATE` | Update semesters | SCHOOL_ADMIN, SYSTEM_ADMIN |
| `semester:DELETE` | Delete semesters | SCHOOL_ADMIN, SYSTEM_ADMIN |

### Subject Category Permissions

| Permission | Description | Roles |
| ---------- | ----------- | ------- |
| `subject_category:READ` | View subject categories | SCHOOL_ADMIN, SYSTEM_ADMIN, CURRICULUM_ADMIN, TEACHER |
| `subject_category:CREATE` | Create subject categories | CURRICULUM_ADMIN, SYSTEM_ADMIN |
| `subject_category:UPDATE` | Update subject categories | CURRICULUM_ADMIN, SYSTEM_ADMIN |
| `subject_category:DELETE` | Delete subject categories | CURRICULUM_ADMIN, SYSTEM_ADMIN |
| `subject_category:DEACTIVATE` | Deactivate subject categories | CURRICULUM_ADMIN, SYSTEM_ADMIN |

### Graduate Profile Dimension Permissions

| Permission | Description | Roles |
| ---------- | ----------- | ------- |
| `graduate_profile_dimension:READ` | View dimensions | SCHOOL_ADMIN, SYSTEM_ADMIN, CURRICULUM_ADMIN, TEACHER |
| `graduate_profile_dimension:CREATE` | Create dimensions | CURRICULUM_ADMIN, SYSTEM_ADMIN |
| `graduate_profile_dimension:UPDATE` | Update dimensions | CURRICULUM_ADMIN, SYSTEM_ADMIN |
| `graduate_profile_dimension:DELETE` | Delete dimensions | CURRICULUM_ADMIN, SYSTEM_ADMIN |
| `graduate_profile_dimension:DEACTIVATE` | Deactivate dimensions | CURRICULUM_ADMIN, SYSTEM_ADMIN |

### CP Alignment Permissions

| Permission | Description | Roles |
| ---------- | ----------- | ------- |
| `cp_alignment:READ` | View CP alignments | SCHOOL_ADMIN, SYSTEM_ADMIN, CURRICULUM_ADMIN, TEACHER |
| `cp_alignment:CREATE` | Create CP alignments | CURRICULUM_ADMIN |
| `cp_alignment:UPDATE` | Update CP alignments | CURRICULUM_ADMIN |
| `cp_alignment:DELETE` | Delete CP alignments | CURRICULUM_ADMIN |

### Alignment Report Permissions

| Permission | Description | Roles |
| ---------- | ----------- | ------- |
| `alignment_report:READ` | View alignment reports | CURRICULUM_ADMIN, SCHOOL_ADMIN |

---

## Resource Ownership

### Academic Year Ownership

**Model**: School-scoped  
**Rule**: School Admin can only manage academic years for their own school  
**Implementation**: 
- All queries filter by `school_id` from JWT token
- School Admin can only access `WHERE school_id = current_user.school_id`
- System Admin can access all schools

### Subject Category Ownership

**Model**: System-wide (no ownership)  
**Rule**: Curriculum Admin manages categories for all schools  
**Implementation**: No ownership filtering - categories are reference data

### Graduate Profile Dimension Ownership

**Model**: System-wide (no ownership)  
**Rule**: Curriculum Admin manages dimensions for all schools  
**Implementation**: No ownership filtering - dimensions are reference data

### CP Alignment Ownership

**Model**: School-scoped (via CP ownership)  
**Rule**: Curriculum Admin can manage alignments for all schools; School Admin for their school only  
**Implementation**: 
- Filter by CP ownership (via school_id through curriculum hierarchy)
- School Admin: `WHERE cp.school_id = current_user.school_id`
- Curriculum Admin: No filtering

---

## Audit Logging

### Audit Events

All Sprint 4 operations must be logged to `audit_logs` table:

| Event | Entity | Action | Fields Logged |
| ----- | ------ | ------ | ------------- |
| Academic Year Created | academic_year | CREATE | id, school_id, name, start_date, end_date |
| Academic Year Updated | academic_year | UPDATE | id, changes (old/new values) |
| Academic Year Submitted | academic_year | SUBMIT | id, status (DRAFT → UNDER_REVIEW) |
| Academic Year Approved | academic_year | APPROVE | id, status (UNDER_REVIEW → APPROVED), approved_by |
| Academic Year Deactivated | academic_year | DEACTIVATE | id, status, reason |
| Semester Created | semester | CREATE | id, academic_year_id, name, sequence, dates |
| Semester Updated | semester | UPDATE | id, changes |
| Subject Category Created | subject_category | CREATE | id, code, name |
| Subject Category Updated | subject_category | UPDATE | id, changes |
| Subject Category Deactivated | subject_category | DEACTIVATE | id, reason |
| Dimension Created | graduate_profile_dimension | CREATE | id, code, name |
| Dimension Updated | graduate_profile_dimension | UPDATE | id, changes |
| Dimension Deactivated | graduate_profile_dimension | DEACTIVATE | id, reason |
| CP Alignment Created | cp_alignment | CREATE | id, cp_id, dimension_id, alignment_strength |
| CP Alignment Updated | cp_alignment | UPDATE | id, changes (old/new strength) |
| CP Alignment Deleted | cp_alignment | DELETE | id, cp_id, dimension_id |

### Audit Log Fields

```go
type AuditLog struct {
    ID        UUID
    UserID    UUID
    Action    string      // CREATE, UPDATE, DELETE, SUBMIT, APPROVE, DEACTIVATE
    EntityType string      // academic_year, semester, subject_category, etc.
    EntityID  UUID
    Changes   JSONB       // Old/new values
    IPAddress INET
    UserAgent string
    CreatedAt time.Time
}
```

### Audit Query API

**Endpoint**: `GET /api/v1/audit/logs`  
**Authorization**: SYSTEM_ADMIN, CURRICULUM_ADMIN  
**Filters**: entity_type, entity_id, user_id, action, date_range  
**Response**: Paginated audit log entries

---

## Data Retention

### Retention Policy

| Entity | Retention Period | Rationale |
| ------ | --------------- | --------- |
| Audit logs | 7 years | Compliance with Indonesian data retention laws |
| Academic year records | 10 years | Historical academic records |
- | Deactivated entities | 5 years | Reference for rollback and analysis |
- | Soft-deleted data | 1 year | Recovery window |

### Data Deletion

**Automatic Deletion**:
- Audit logs older than 7 years are automatically purged via scheduled job
- Soft-deleted entities older than 5 years are permanently deleted

**Manual Deletion**:
- SYSTEM_ADMIN can manually purge data via admin API
- Requires confirmation and audit log entry

---

## Soft Delete Strategy

### Implementation

Sprint 4 uses **status-based soft delete** (not `deleted_at` columns):

**Academic Year**:
- Status: `INACTIVE` (soft delete)
- Cannot be deleted if has dependencies (CP, planning)
- Past academic years automatically marked `INACTIVE`

**Subject Category**:
- Status: `INACTIVE` (soft delete)
- Cannot be deactivated if subjects are assigned

**Graduate Profile Dimension**:
- Status: `INACTIVE` (soft delete)
- Cannot be deactivated if CP are aligned

**Semester**:
- Status: `INACTIVE` (soft delete)
- Cascades with academic year

**CP Alignment**:
- Physical delete (CASCADE with CP)
- Alignment is not critical for historical records

### Query Implications

All read queries must filter by status:
```sql
WHERE status = 'ACTIVE'
```

Admin queries can include inactive entities:
```sql
WHERE status IN ('ACTIVE', 'INACTIVE')
```

---

# PART 9 – TEST STRATEGY

## Unit Test Matrix

### Backend Unit Tests

| Module | Test Cases | Coverage Target | Priority |
| ------ | ---------- | --------------- | -------- |
| Academic Year Domain | 15 | 90% | P0 |
| Semester Domain | 10 | 90% | P0 |
| Subject Category Domain | 8 | 85% | P1 |
| Graduate Profile Dimension Domain | 12 | 90% | P0 |
| CP Alignment Domain | 10 | 85% | P1 |
| Academic Year Validation Service | 8 | 95% | P0 |
| Alignment Calculation Service | 6 | 95% | P0 |
| Academic Year Repository | 12 | 80% | P1 |
| Semester Repository | 8 | 80% | P1 |
| Subject Category Repository | 6 | 80% | P2 |
| Graduate Profile Dimension Repository | 8 | 80% | P1 |
| CP Alignment Repository | 10 | 80% | P1 |

**Total Backend Unit Tests**: 93 test cases

---

### Frontend Unit Tests

| Component | Test Cases | Coverage Target | Priority |
| --------- | ---------- | --------------- | -------- |
| Academic Year Form | 10 | 85% | P0 |
| Semester Form | 8 | 85% | P0 |
| Subject Category Form | 6 | 80% | P1 |
| Graduate Profile Dimension Form | 10 | 85% | P0 |
| CP Alignment Form | 8 | 80% | P1 |
| Alignment Report | 5 | 75% | P1 |
| Academic Year Table | 6 | 80% | P1 |
| Semester Timeline | 4 | 75% | P2 |
| Dimension Radar Chart | 3 | 70% | P2 |

**Total Frontend Unit Tests**: 60 test cases

---

## Integration Test Matrix

### Backend Integration Tests

| Scenario | Test Cases | Priority |
| -------- | ---------- | -------- |
| Create Academic Year + Semesters | 5 | P0 |
| Submit + Approve Academic Year Workflow | 4 | P0 |
| Academic Year Date Overlap Validation | 3 | P0 |
| Semester Coverage Validation | 3 | P0 |
| Subject Category + Subject Assignment | 4 | P1 |
| Graduate Profile Dimension + CP Alignment | 5 | P0 |
| Alignment Calculation | 3 | P0 |
| CP Alignment Below Threshold Warning | 2 | P1 |
| Audit Logging for All Operations | 8 | P1 |
| Permission-based Access Control | 6 | P0 |
| Soft Delete + Query Filtering | 4 | P1 |

**Total Backend Integration Tests**: 47 test cases

---

### Frontend Integration Tests

| Scenario | Test Cases | Priority |
| -------- | ---------- | -------- |
| Academic Year CRUD Flow | 4 | P0 |
| Semester Configuration Flow | 4 | P0 |
| Subject Category Management Flow | 3 | P1 |
| CP Alignment Management Flow | 4 | P0 |
| Alignment Report Generation | 2 | P1 |
| Permission-based UI Access | 5 | P0 |
| Form Validation Error Handling | 6 | P0 |
| API Error Handling | 5 | P1 |

**Total Frontend Integration Tests**: 33 test cases

---

## API Test Matrix

### Contract Tests

| Endpoint | Test Cases | Priority |
| -------- | ---------- | -------- |
| POST /academic/academic-years | 4 | P0 |
| GET /academic/academic-years | 3 | P0 |
| PUT /academic/academic-years/:id | 3 | P0 |
| POST /academic/academic-years/:id/submit | 2 | P0 |
| POST /academic/academic-years/:id/approve | 2 | P0 |
| POST /academic/academic-years/:id/deactivate | 2 | P0 |
| POST /academic/semesters | 3 | P0 |
| GET /academic/semesters | 2 | P0 |
| PUT /academic/semesters/:id | 2 | P0 |
| DELETE /academic/semesters/:id | 2 | P1 |
| POST /curriculum/subject-categories | 3 | P1 |
| GET /curriculum/subject-categories | 2 | P1 |
| PUT /curriculum/subject-categories/:id | 2 | P1 |
| POST /curriculum/graduate-profile-dimensions | 3 | P0 |
| GET /curriculum/graduate-profile-dimensions | 2 | P0 |
| PUT /curriculum/graduate-profile-dimensions/:id | 2 | P0 |
| POST /curriculum/cp-alignments | 3 | P0 |
| GET /curriculum/cp-alignments | 2 | P0 |
| PUT /curriculum/cp-alignments/:id | 2 | P0 |
| DELETE /curriculum/cp-alignments/:id | 2 | P1 |
| GET /curriculum/alignment-report | 2 | P1 |

**Total API Contract Tests**: 44 test cases

---

## E2E Test Matrix

### User Journeys

| Journey | Test Cases | Priority |
| ------- | ---------- | -------- |
| School Admin: Create Academic Year | 3 | P0 |
| School Admin: Configure Semesters | 3 | P0 |
| School Admin: Submit for Approval | 2 | P0 |
| System Admin: Approve Academic Year | 2 | P0 |
| Curriculum Admin: Create Subject Category | 2 | P1 |
| Curriculum Admin: Configure Graduate Profile Dimensions | 4 | P0 |
| Curriculum Admin: Align CP to Dimensions | 3 | P0 |
| Curriculum Admin: Generate Alignment Report | 2 | P1 |
| Curriculum Admin: Add Koding Subject | 2 | P1 |
| Teacher: View Academic Calendar | 2 | P2 |

**Total E2E Test Cases**: 25 test cases

---

## UAT Checklist

### School Admin UAT

- [ ] Can create academic year with valid date range
- [ ] System prevents overlapping academic years
- [ ] System requires 30-day lead time for new academic years
- [ ] Can configure 2 semesters within academic year
- [ ] System validates semester coverage (no gaps, no overlaps)
- [ ] Can submit academic year for approval
- [ ] Cannot approve academic year (not SYSTEM_ADMIN)
- [ ] Can deactivate academic year with reason
- [ ] Can view academic year list with filters
- [ ] Can view academic year details with semesters
- [ ] Audit logs are generated for all actions
- [ ] Permission-based access control works correctly

### Curriculum Admin UAT

- [ ] Can create subject category
- [ ] Cannot create duplicate category codes or names
- [ ] Can categorize subjects (single and bulk)
- [ ] Cannot deactivate category with assigned subjects
- [ ] Can create graduate profile dimension
- [ ] Can configure exactly 8 dimensions
- [ ] Can align CP to dimensions
- [ ] System calculates alignment percentage automatically
- [ ] System warns when CP falls below 60% alignment
- [ ] Can view alignment reports
- [ ] Can export alignment reports (PDF/CSV)
- [ ] Can add Koding, AI, Numerasi subjects
- [ ] New subjects integrate with CP/TP workflow
- [ ] Audit logs are generated for all actions
- [ ] Permission-based access control works correctly

### Teacher UAT

- [ ] Can view academic calendar (read-only)
- [ ] Can view subject categories (read-only)
- [ ] Can view graduate profile dimensions (read-only)
- [ ] Cannot edit academic structure
- [ ] UI displays in Indonesian language
- [ ] Page loading times are acceptable (<2s)

### System Admin UAT

- [ ] Can approve academic year submissions
- [ ] Can reject academic year submissions
- [ ] Can view audit logs
- [ ] Can create CURRICULUM_ADMIN role
- [ ] Can assign permissions to CURRICULUM_ADMIN
- [ ] Can manage academic years for all schools

---

## Performance Testing

### Performance Targets

| Operation | Target | Measurement |
| --------- | ------ | ----------- |
| API Response Time (Read) | <500ms | 95th percentile |
| API Response Time (Write) | <2s | 95th percentile |
| Page Load Time | <2s | 95th percentile |
| Database Query Time | <100ms | Average |
| Alignment Report Generation | <5s | For 1000 CP |

### Load Testing

- Concurrent users: 100
- Requests per second: 50
- Test duration: 30 minutes
- Success rate target: >99.9%

---

# PART 10 – IMPLEMENTATION PLAN

## Task Breakdown

### Backend Tasks

| ID | Task | Dependencies | Estimate | Priority |
| ---- | ---- | ------------ | -------- | -------- |
| B-001 | Create database migration (000010) | None | 4 hours | P0 |
| B-002 | Create AcademicYear domain models | B-001 | 6 hours | P0 |
| B-003 | Create Semester domain models | B-001 | 4 hours | P0 |
| B-004 | Create SubjectCategory domain models | B-001 | 3 hours | P1 |
| B-005 | Create GraduateProfileDimension domain models | B-001 | 5 hours | P0 |
| B-006 | Create CPAlignment domain models | B-001, B-005 | 4 hours | P1 |
| B-007 | Create AcademicYearValidationService | B-002 | 4 hours | P0 |
| B-008 | Create AlignmentCalculationService | B-006 | 3 hours | P0 |
| B-009 | Create AcademicYearRepository | B-002, B-003 | 4 hours | P0 |
| B-010 | Create SubjectCategoryRepository | B-004 | 3 hours | P1 |
| B-011 | Create GraduateProfileDimensionRepository | B-005 | 4 hours | P0 |
| B-012 | Create CPAlignmentRepository | B-006 | 4 hours | P1 |
| B-013 | Create AcademicYearService | B-002, B-007, B-009 | 6 hours | P0 |
| B-014 | Create SubjectCategoryService | B-004, B-010 | 4 hours | P1 |
| B-015 | Create GraduateProfileDimensionService | B-005, B-008, B-011 | 5 hours | P0 |
| B-016 | Create CPAlignmentService | B-006, B-008, B-012 | 5 hours | P1 |
| B-017 | Create AcademicYearHandler | B-013 | 4 hours | P0 |
| B-018 | Create SemesterHandler | B-013 | 3 hours | P0 |
| B-019 | Create SubjectCategoryHandler | B-014 | 3 hours | P1 |
| B-020 | Create GraduateProfileDimensionHandler | B-015 | 4 hours | P0 |
| B-021 | Create CPAlignmentHandler | B-016 | 4 hours | P1 |
| B-022 | Register routes in router.go | B-017, B-018, B-019, B-020, B-021 | 2 hours | P0 |
| B-023 | Add CURRICULUM_ADMIN role to seed data | None | 1 hour | P0 |
| B-024 | Add Sprint 4 permissions to seed data | B-023 | 2 hours | P0 |
| B-025 | Write backend unit tests | B-001 to B-024 | 16 hours | P0 |
| B-026 | Write backend integration tests | B-001 to B-024 | 12 hours | P1 |

**Total Backend Effort**: 107 hours (approx. 13.5 days)

---

### Frontend Tasks

| ID | Task | Dependencies | Estimate | Priority |
| ---- | ---- | ------------ | -------- | -------- |
| F-001 | Create Academic Year Management page | None | 8 hours | P0 |
| F-002 | Create Semester Configuration page | F-001 | 6 hours | P0 |
| F-003 | Create Subject Category Management page | None | 6 hours | P1 |
| F-004 | Create Subject Categorization page | F-003 | 4 hours | P1 |
| F-005 | Create Graduate Profile Dimensions page | None | 8 hours | P0 |
| F-006 | Create CP Alignment Management page | F-005 | 8 hours | P1 |
| F-007 | Create Alignment Report page | F-006 | 6 hours | P1 |
| F-008 | Create Koding & AI Subject Setup page | None | 6 hours | P1 |
| F-009 | Create Academic Calendar view (Teacher) | F-001 | 3 hours | P2 |
| F-010 | Create Subject Categories view (Teacher) | F-003 | 2 hours | P2 |
| F-011 | Create Graduate Profile Reference (Teacher) | F-005 | 3 hours | P2 |
| F-012 | Add navigation menu items | F-001, F-003, F-005 | 2 hours | P0 |
| F-013 | Implement permission-based UI access | F-012 | 4 hours | P0 |
| F-014 | Implement Indonesian translations | F-001 to F-011 | 6 hours | P1 |
| F-015 | Write frontend unit tests | F-001 to F-014 | 12 hours | P1 |
| F-016 | Write frontend integration tests | F-001 to F-014 | 10 hours | P1 |

**Total Frontend Effort**: 94 hours (approx. 12 days)

---

### Database Tasks

| ID | Task | Dependencies | Estimate | Priority |
| ---- | ---- | ------------ | -------- | -------- |
| D-001 | Review migration script for syntax errors | B-001 | 1 hour | P0 |
| D-002 | Test migration on fresh database | D-001 | 1 hour | P0 |
| D-003 | Test migration rollback | D-002 | 1 hour | P0 |
| D-004 | Verify data migration (categories, dimensions) | D-003 | 1 hour | P0 |
| D-005 | Create database performance indexes if needed | D-004 | 1 hour | P2 |
| D-006 | Document migration in DATABASE_SCHEMA_FREEZE | D-005 | 2 hours | P1 |

**Total Database Effort**: 7 hours (approx. 1 day)

---

### QA Tasks

| ID | Task | Dependencies | Estimate | Priority |
| ---- | ---- | ------------ | -------- | -------- |
| Q-001 | Review test plans | B-026, F-016 | 2 hours | P0 |
| Q-002 | Execute backend unit tests | B-025 | 2 hours | P0 |
| Q-003 | Execute backend integration tests | B-026 | 3 hours | P0 |
| Q-004 | Execute frontend unit tests | F-015 | 2 hours | P0 |
| Q-005 | Execute frontend integration tests | F-016 | 3 hours | P0 |
| Q-006 | Execute API contract tests | B-026 | 2 hours | P0 |
| Q-007 | Execute E2E tests | All | 4 hours | P0 |
| Q-008 | Perform UAT with stakeholders | All | 4 hours | P0 |
| Q-009 | Performance testing | All | 3 hours | P1 |
| Q-010 | Security testing | All | 2 hours | P1 |
| Q-011 | Document test results | Q-002 to Q-010 | 2 hours | P1 |

**Total QA Effort**: 27 hours (approx. 3.5 days)

---

## Implementation Sequence

### Phase 1: Foundation (Week 1)

**Goal**: Database migration and core domain models

**Tasks**:
- D-001, D-002, D-003 (Migration)
- B-001 (Database migration script)
- B-002, B-003 (AcademicYear, Semester models)
- B-004 (SubjectCategory model)
- B-005, B-006 (GraduateProfileDimension, CPAlignment models)
- B-023, B-024 (Role and permission seed data)

**Deliverables**:
- Working migration with rollback
- All domain models defined
- CURRICULUM_ADMIN role and permissions created

---

### Phase 2: Backend Services (Week 2)

**Goal**: Repositories and services

**Tasks**:
- B-007, B-008 (Domain services)
- B-009 to B-012 (Repositories)
- B-013 to B-016 (Application services)
- B-025 (Backend unit tests)

**Deliverables**:
- All repositories implemented
- All services implemented
- Unit tests passing

---

### Phase 3: Backend API (Week 2-3)

**Goal**: Handlers and routes

**Tasks**:
- B-017 to B-021 (Handlers)
- B-022 (Route registration)
- B-026 (Backend integration tests)

**Deliverables**:
- All API endpoints functional
- Integration tests passing

---

### Phase 4: Frontend Pages (Week 3-4)

**Goal**: Frontend UI implementation

**Tasks**:
- F-001 to F-008 (All pages)
- F-012, F-013 (Navigation and permissions)
- F-015 (Frontend unit tests)

**Deliverables**:
- All pages implemented
- Permission-based access working
- Unit tests passing

---

### Phase 5: Frontend Polish (Week 4)

**Goal**: Teacher pages and localization

**Tasks**:
- F-009 to F-011 (Teacher pages)
- F-014 (Indonesian translations)
- F-016 (Frontend integration tests)

**Deliverables**:
- Teacher pages implemented
- Indonesian localization complete
- Integration tests passing

---

### Phase 6: QA & UAT (Week 5)

**Goal**: Testing and validation

**Tasks**:
- D-004 to D-006 (Database validation)
- Q-001 to Q-008 (All testing)
- Q-009, Q-010 (Performance and security testing)

**Deliverables**:
- All tests passing
- UAT sign-off from stakeholders
- Performance and security validation

---

### Phase 7: Documentation & Handoff (Week 5)

**Goal**: Documentation and deployment readiness

**Tasks**:
- Q-011 (Test results documentation)
- Update DATABASE_SCHEMA_FREEZE_V1.md
- Update API documentation
- Create deployment guide

**Deliverables**:
- Complete documentation
- Deployment-ready package

---

## Total Effort Summary

| Category | Hours | Days | Weeks |
| -------- | ----- | ---- | ----- |
| Backend | 107 | 13.5 | 1.7 |
| Frontend | 94 | 12.0 | 1.5 |
| Database | 7 | 1.0 | 0.1 |
| QA | 27 | 3.5 | 0.4 |
| **Total** | **235** | **29.5** | **3.7** |

**Recommended Timeline**: 5 weeks (including buffer)

**Optimistic Timeline**: 4 weeks (if no blocking issues)

**Conservative Timeline**: 6 weeks (including risk mitigation)

---

## Dependencies

### External Dependencies

None - Sprint 4 uses only existing infrastructure and libraries

### Internal Dependencies

1. **Sprint 1-3 Modules**: Must be stable and not in active development
2. **Database Schema Freeze v1**: Must be approved for extension
3. **Architecture Freeze v2**: Must be compliant with new role and permissions
4. **Authentication/Authorization**: Existing JWT and RBAC must be working

### Blocking Dependencies

- Architecture Board approval for database schema extension
- Architecture Board approval for CURRICULUM_ADMIN role
- Database migration approval from DBA
- UAT sign-off from Curriculum Admin stakeholders

---

## Risks and Mitigations

### Risk 1: Architecture Board Approval Delay

**Probability**: Medium  
**Impact**: High  
**Mitigation**: Submit architecture change request early (Week -1), provide detailed rationale

### Risk 2: Migration Complexity

**Probability**: Low  
**Impact**: Medium  
**Mitigation**: Test migration on staging database first, have rollback plan ready

### Risk 3: Frontend-Backend Integration Issues

**Probability**: Medium  
**Impact**: High  
**Mitigation**: API contract tests first, parallel development with mock data

### Risk 4: Stakeholder UAT Rejection

**Probability**: Low  
**Impact**: Medium  
**Mitigation**: Early demo of prototypes, collect feedback before implementation

### Risk 5: Performance Issues with Alignment Calculation

**Probability**: Low  
**Impact**: Medium  
**Mitigation**: Implement caching for alignment reports, optimize database queries

### Risk 6: Solo Developer Time Constraints

**Probability**: High  
**Impact**: High  
**Mitigation**: Prioritize P0 features first, defer P2 features if needed

---

## Success Criteria

Sprint 4 will be considered successful when:

1. **Functional Requirements**: All P0 features are implemented and working
2. **Quality Gates**: All unit tests passing (>90% coverage), all integration tests passing
3. **Performance**: API response times meet targets (<500ms read, <2s write)
4. **UAT Sign-off**: All stakeholders approve implementation
5. **Documentation**: Complete documentation including API, database, and deployment guides
6. **Architecture Compliance**: No violations of Architecture Freeze v2
7. **Kurikulum Compliance**: All features align with Kurikulum Merdeka 2026
8. **Deployment Ready**: Migration tested and rollback plan validated

---

## Appendix A: Kurikulum 2026 Reference

### 8 Profil Lulusan Dimensions

1. **Keimanan & Ketakwaan**
   - Faith and piety toward God Almighty
   - Indicators: Noble character, maintaining purity of heart, performing worship

2. **Kewargaan**
   - Awareness of nationality and statehood
   - Indicators: Love for the homeland, respect for diversity, obedience to rules

3. **Berakhlak Mulia**
   - Formation of noble character
   - Indicators: Honesty, discipline, responsibility

4. **Berani Bertanggung Jawab**
   - Courage and responsibility
   - Indicators: Courage to make decisions, accountability for actions

5. **Peduli**
   - Concern for others
   - Indicators: Empathy, helping each other, tolerance

6. **Gotong Royong**
   - Cooperation and mutual assistance
   - Indicators: Teamwork, solidarity, collaboration

7. **Mandiri**
   - Independence and autonomy
   - Indicators: Critical thinking, initiative, independent learning

8. **Kreatif**
   - Creativity and innovation
   - Indicators: Creative thinking, innovation, problem-solving

### Subject Categories (Kurikulum Merdeka)

1. **Intrakurikuler**
   - Core curriculum subjects
   - Mandatory for all students according to phase
   - Examples: Bahasa Indonesia, Matematika, IPAS

2. **Kokurikuler**
   - Self-development activities
   - Developing potential and talents
   - Examples: Scout activities, religious activities

3. **Ekstrakurikuler**
   - Additional activities outside curriculum
   - Developing interests and talents
   - Examples: Sports, arts, clubs

---

## Appendix B: API Error Codes Reference

| Code | HTTP Status | Description |
| ---- | ---------- | ----------- |
| `ACADEMIC_YEAR_OVERLAP` | 409 | Academic year dates overlap with existing academic year |
| `ACADEMIC_YEAR_LEAD_TIME` | 422 | Academic year must be created at least 30 days in advance |
| `SEMESTER_COVERAGE_GAP` | 422 | Semesters must fully cover academic year without gaps |
| `SEMESTER_COVERAGE_OVERLAP` | 422 | Semesters cannot overlap |
| `CATEGORY_HAS_SUBJECTS` | 409 | Cannot deactivate category with assigned subjects |
| `DIMENSION_HAS_CP` | 409 | Cannot deactivate dimension with aligned CP |
| `DIMENSION_COUNT_VIOLATION` | 422 | Exactly 8 dimensions must be active |
| `ALIGNMENT_BELOW_THRESHOLD` | 400 | CP alignment will fall below 60% threshold |
| `ALIGNMENT_DUPLICATE` | 409 | CP already aligned to this dimension |

---

## Appendix C: Database Migration Checklist

- [ ] Migration script reviewed for syntax errors
- [ ] Migration tested on fresh database
- [ ] Rollback tested and working
- [ ] Data migration verified (categories, dimensions)
- [ ] Indexes created and verified
- [ ] Foreign keys verified
- [ ] Audit fields included
- [ ] Performance tested with sample data
- [ ] Documentation updated (DATABASE_SCHEMA_FREEZE_V1.md)
- [ ] Backup plan documented
- [ ] Rollback plan documented
- [ ] Migration approved by DBA
- [ ] Migration scheduled for non-peak hours

---

**Document Status**: COMPLETE - READY FOR IMPLEMENTATION

**Next Steps**:
1. Architecture Board review and approval
2. Database migration approval
3. Begin Phase 1 implementation
4. Daily standups to track progress
5. Weekly reviews with stakeholders

I'll add additional sections to make the Sprint 4 Requirement Package even more comprehensive and implementation-ready.

---

# PART 11 – EXECUTIVE SUMMARY

## Overview

Sprint 4: Academic Foundation is a 5-week implementation sprint that establishes the academic infrastructure required to support Kurikulum 2026 and Deep Learning pedagogy in the NUSA Platform. This sprint focuses on governance and reference data rather than student lifecycle or class management features.

## Key Deliverables

1. **Academic Year Management**: Enable schools to configure academic years with semester divisions
2. **Subject Categorization**: Implement Kurikulum Merdeka subject categories (Intrakurikuler, Kokurikuler, Ekstrakurikuler)
3. **Graduate Profile Dimensions**: Configure the 8-dimensional Profil Lulusan framework
4. **Curriculum Alignment**: Track and manage CP alignment to graduate profile dimensions
5. **Modern Subject Support**: Enable Koding, AI, and Numerasi as formal curriculum subjects

## Strategic Value

- **Foundation for Future Sprints**: Academic infrastructure is prerequisite for Student Lifecycle, Class Management, and Enrollment
- **Kurikulum 2026 Compliance**: Ensures platform alignment with national curriculum standards
- **Governance Capability**: Enables curriculum administrators to manage academic structure without developer intervention
- **Data-Driven Insights**: Alignment reporting provides visibility into curriculum coverage

## Technical Approach

- **No New Bounded Contexts**: All extensions fit within existing Curriculum Context
- **Minimal Database Changes**: 5 new tables, 2 table extensions
- **Architecture Compliance**: Follows existing DDD and Clean Architecture patterns
- **Reuse Strategy**: Maximizes reuse of Sprint 1-3 modules

## Resource Requirements

- **Development**: 235 hours (29.5 days) across backend, frontend, database, and QA
- **Timeline**: 5 weeks recommended (6 weeks conservative)
- **Team**: Solo developer (requires prioritization of P0 features)
- **Infrastructure**: No new infrastructure requirements

## Success Metrics

- 100% of schools able to configure academic years
- 100% of subjects categorized according to Kurikulum Merdeka
- 8 graduate profile dimensions configured and active
- CP alignment tracking operational with <60% threshold warnings
- API response times <500ms (read), <2s (write)
- Test coverage >90% for critical paths

## Risk Summary

| Risk | Level | Mitigation |
| ---- | ---- | ---------- |
| Architecture Board approval delay | Medium | Submit early, provide detailed rationale |
| Migration complexity | Low | Test on staging, rollback plan ready |
| Frontend-backend integration | Medium | API contract tests first, parallel dev |
| Stakeholder UAT rejection | Low | Early prototype demos, feedback loop |
| Performance issues | Low | Caching for reports, query optimization |
| Solo developer time constraints | High | Prioritize P0, defer P2 if needed |

## Go/No-Go Recommendation

**CONDITIONAL GO**

**Prerequisites**:
1. Architecture Board approval for database schema extension
2. Architecture Board approval for CURRICULUM_ADMIN role
3. Database migration approval from DBA
4. Stakeholder sign-off on functional requirements

**Recommended Start Date**: Week after approvals obtained

**Critical Success Factor**: Prioritize P0 features (Academic Year, Semester, Graduate Profile Dimensions) over P1/P2 features if timeline constraints emerge.

---

# PART 12 – DATA MIGRATION SCRIPTS

## Script 1: Extend curriculum_subjects with category_id

```sql
-- Data Migration: Assign default category to existing subjects
-- Run this after migration 000010 is applied

-- Step 1: Verify migration
SELECT COUNT(*) as uncategorized_subjects 
FROM curriculum_subjects 
WHERE subject_category_id IS NULL;

-- Step 2: Assign Intrakurikuler as default category
-- This assumes subject category with code 'INTRAKURIKULER' was created in migration
UPDATE curriculum_subjects
SET subject_category_id = (
    SELECT id FROM subject_categories 
    WHERE code = 'INTRAKURIKULER' 
    LIMIT 1
)
WHERE subject_category_id IS NULL;

-- Step 3: Verify result
SELECT 
    sc.code as category_code,
    sc.name as category_name,
    COUNT(cs.id) as subject_count
FROM subject_categories sc
LEFT JOIN curriculum_subjects cs ON cs.subject_category_id = sc.id
GROUP BY sc.id, sc.code, sc.name;
```

## Script 2: Validate 8 Dimensions Configuration

```sql
-- Validation Script: Verify exactly 8 active graduate profile dimensions
-- Run after migration to ensure configuration is correct

-- Check active dimension count
SELECT 
    COUNT(*) as active_dimensions,
    COUNT(CASE WHEN status = 'ACTIVE' THEN 1 END) as expected
FROM graduate_profile_dimensions;

-- Verify all 8 required dimensions exist
SELECT 
    code, 
    name, 
    status,
    weight
FROM graduate_profile_dimensions
WHERE code IN (
    'KEIMANAN_KETAKWAAN',
    'KEWARGAAN',
    'BERAKHLAK_MULIA',
    'BERANI_BERTANGGUNG_JAWAB',
    'PEDULI',
    'GOTONG_ROYONG',
    'MANDIRI',
    'KREATIF'
)
ORDER BY code;

-- Verify weight distribution (should be 0.125 each for equal distribution)
SELECT 
    code,
    name,
    weight,
    weight * 100 as percentage
FROM graduate_profile_dimensions
WHERE status = 'ACTIVE'
ORDER BY weight;
```

## Script 3: Sample Data for Testing

```sql
-- Test Data: Create sample academic year with semesters
-- Use for development and testing purposes

-- Create sample academic year for a test school
INSERT INTO academic_years (
    id, school_id, name, start_date, end_date, 
    status, created_by, created_at, updated_at
)
VALUES (
    gen_uuid_v7(),
    (SELECT id FROM schools LIMIT 1),
    '2026/2027',
    '2026-07-15T00:00:00Z',
    '2027-06-30T23:59:59Z',
    'DRAFT',
    (SELECT id FROM users WHERE role = 'SYSTEM_ADMIN' LIMIT 1),
    NOW(),
    NOW()
);

-- Create semesters for the academic year
INSERT INTO semesters (
    id, academic_year_id, name, sequence, 
    start_date, end_date, status, created_at, updated_at
)
SELECT 
    gen_uuid_v7(),
    ay.id,
    'Semester Ganjil',
    1,
    ay.start_date,
    '2026-12-31T23:59:59Z',
    'DRAFT',
    NOW(),
    NOW()
FROM academic_years ay
WHERE ay.name = '2026/2027';

INSERT INTO semesters (
    id, academic_year_id, name, sequence, 
    start_date, end_date, status, created_at, updated_at
)
SELECT 
    gen_uuid_v7(),
    ay.id,
    'Semester Genap',
    2,
    '2027-01-01T00:00:00Z',
    ay.end_date,
    'DRAFT',
    NOW(),
    NOW()
FROM academic_years ay
WHERE ay.name = '2026/2027';
```

## Script 4: Audit Log Verification

```sql
-- Audit Verification: Verify audit logs are being created
-- Run after implementing features to test audit logging

-- Check audit logs for academic year operations
SELECT 
    entity_type,
    action,
    COUNT(*) as operation_count,
    MIN(created_at) as first_operation,
    MAX(created_at) as last_operation
FROM audit_logs
WHERE entity_type IN ('academic_year', 'semester', 'subject_category', 'graduate_profile_dimension')
GROUP BY entity_type, action
ORDER BY entity_type, action;

-- Verify audit log fields
SELECT 
    al.id,
    al.user_id,
    u.full_name as user_name,
    al.action,
    al.entity_type,
    al.entity_id,
    al.created_at
FROM audit_logs al
JOIN users u ON al.user_id = u.id
WHERE al.entity_type = 'academic_year'
ORDER BY al.created_at DESC
LIMIT 10;
```

---

# PART 13 – DEPLOYMENT CHECKLIST

## Pre-Deployment Checklist

### Code Review
- [ ] All code reviewed against Architecture Freeze v2
- [ ] No CQRS patterns introduced
- [ ] No Event Sourcing patterns introduced
- [ ] No new bounded contexts without approval
- [ ] Code follows existing patterns (Handler → Service → Repository)
- [ ] No business logic in handlers
- [ ] Repository layer only contains database operations
- [ ] Domain logic in domain layer only
- [ ] Proper error handling throughout
- [ ] No hardcoded values (use environment variables)

### Database Review
- [ ] Migration script reviewed by DBA
- [ ] Migration tested on staging database
- [ ] Rollback script tested and working
- [ ] Data migration verified
- [ ] Indexes created and verified
- [ ] Foreign keys verified
- [ ] No orphaned records after migration
- [ ] Backup taken before migration
- [ ] Rollback plan documented

### Testing Review
- [ ] Unit tests passing (>90% coverage for critical paths)
- [ ] Integration tests passing
- [ ] API contract tests passing
- [ ] E2E tests passing
- [ ] Performance tests meeting targets
- [ ] Security tests passing
- [ ] UAT sign-off obtained
- [ ] Test results documented

### Security Review
- [ ] CURRICULUM_ADMIN role approved
- [ ] All permissions defined and granted
- [ ] Permission-based access control tested
- [ ] Audit logging verified
- [ ] No SQL injection vulnerabilities
- [ ] No XSS vulnerabilities
- [ ] CSRF protection enabled
- [ ] Rate limiting configured
- [ ] Sensitive data not logged

### Documentation Review
- [ ] Database schema documentation updated
- [ ] API documentation updated
- [ ] Deployment guide created
- [ ] Rollback procedure documented
- [ ] Monitoring guide created
- [ ] Runbook created
- [ ] CHANGELOG.md updated

---

## Deployment Procedure

### Phase 1: Preparation (Day -1)

1. **Backup Database**
   ```bash
   podman exec nusa-postgres pg_dump -U nusa_user nusa_db > backup_pre_sprint4.sql
   ```

2. **Stop Backend Services**
   ```bash
   podman stop nusa-backend
   ```

3. **Verify Backup**
   ```bash
   # Verify backup file exists and is not empty
   ls -lh backup_pre_sprint4.sql
   ```

### Phase 2: Migration Execution (Day 0 - Maintenance Window)

1. **Apply Migration**
   ```bash
   podman exec nusa-postgres psql -U nusa_user -d nusa_db -f /migrations/000010_sprint4_academic_foundation.up.sql
   ```

2. **Verify Migration Success**
   ```bash
   # Check new tables exist
   podman exec nusa-postgres psql -U nusa_user -d nusa_db -c "\dt academic_years"
   
   # Verify row counts
   podman exec nusa-postgres psql -U nusa_user -d nusa_db -c "SELECT COUNT(*) FROM subject_categories"
   
   # Verify data migration
   podman exec nusa-postgres psql -U nusa_user -d nusa_db -c "SELECT COUNT(*) FROM curriculum_subjects WHERE subject_category_id IS NULL"
   ```

3. **Seed Data**
   ```bash
   # Run seed data scripts if not included in migration
   podman exec nusa-postgres psql -U nusa_user -d nusa_db -f /scripts/seed_sprint4_data.sql
   ```

### Phase 3: Backend Deployment (Day 0)

1. **Build New Docker Image**
   ```bash
   cd backend
   podman build -t nusa-backend:v1.4.0 .
   ```

2. **Stop Old Container**
   ```bash
   podman stop nusa-backend
   podman rm nusa-backend
   ```

3. **Start New Container**
   ```bash
   podman run -d \
     --name nusa-backend \
     -p 8081:8080 \
     -e DB_HOST=host.docker.internal \
     -e DB_PORT=5432 \
     -e DB_NAME=nusa_db \
     -e DB_USER=nusa_user \
     -e DB_PASSWORD=nusa_password \
     nusa-backend:v1.4.0
   ```

4. **Health Check**
   ```bash
   curl http://localhost:8081/health
   curl http://localhost:8081/ready
   curl http://localhost:8081/version
   ```

### Phase 4: Frontend Deployment (Day 0)

1. **Build Frontend**
   ```bash
   cd frontend
   npm run build
   ```

2. **Deploy to Production**
   ```bash
   # Copy build artifacts to web server
   cp -r build/* /var/www/nusa-frontend/
   ```

3. **Verify Frontend**
   ```bash
   curl https://nusa.example.com/
   ```

### Phase 5: Smoke Tests (Day 0)

1. **Test Authentication**
   ```bash
   curl -X POST https://api.nusa.example.com/api/v1/public/auth/login \
     -H "Content-Type: application/json" \
     -d '{"username":"admin@nusa.id","password":"admin123"}'
   ```

2. **Test Academic Year API**
   ```bash
   curl -X GET https://api.nusa.example.com/api/v1/academic/academic-years \
     -H "Authorization: Bearer <token>"
   ```

3. **Test Subject Category API**
   ```bash
   curl -X GET https://api.nusa.example.com/api/v1/curriculum/subject-categories \
     -H "Authorization: Bearer <token>"
   ```

4. **Test Graduate Profile API**
   ```bash
   curl -X GET https://api.nusa.example.com/api/v1/curriculum/graduate-profile-dimensions \
     -H "Authorization: Bearer <token>"
   ```

### Phase 6: Monitoring Setup (Day 0)

1. **Configure Alerts**
   - API error rate > 1% → Alert
   - API response time > 2s → Alert
   - Database connection failures → Alert
   - Audit log failures → Alert

2. **Verify Monitoring**
   - Check dashboard connectivity
   - Verify alert delivery
   - Test alert escalation

---

## Rollback Procedure

### Trigger Conditions

Rollback should be triggered if:
- Critical bugs discovered in production
- Performance degradation > 50%
- Data corruption detected
- Security vulnerability identified
- Stakeholder requests rollback

### Rollback Steps

1. **Stop Services**
   ```bash
   podman stop nusa-backend
   ```

2. **Rollback Database Migration**
   ```bash
   podman exec nusa-postgres psql -U nusa_user -d nusa_db -f /migrations/000010_sprint4_academic_foundation.down.sql
   ```

3. **Verify Rollback**
   ```bash
   # Verify new tables are dropped
   podman exec nusa-postgres psql -U nusa_user -d nusa_db -c "\dt" | grep -E "academic_years|semesters|subject_categories"
   
   # Should return empty
   ```

4. **Restore Previous Backend Version**
   ```bash
   podman stop nusa-backend
   podman rm nusa-backend
   podman run -d \
     --name nusa-backend \
     -p 8081:8081 \
     nusa-backend:v1.3.0
   ```

5. **Restore Previous Frontend Version**
   ```bash
   cp -r /var/www/nusa-frontend-backup/* /var/www/nusa-frontend/
   ```

6. **Start Services**
   ```bash
   podman start nusa-backend
   ```

7. **Verify System**
   ```bash
   curl http://localhost:8081/health
   ```

### Rollback Verification

- [ ] Database schema matches pre-deployment state
- [ ] Backend services healthy
- [ ] Frontend accessible
- [ ] Authentication working
- [ ] No data loss
- [ ] Audit log of rollback created

---

# PART 14 – MONITORING AND ALERTING

## Metrics to Monitor

### Application Metrics

| Metric | Type | Threshold | Alert Level |
| ------ | ---- | ---------- | ---------- |
| API Error Rate | Percentage | > 1% | WARNING |
| API Error Rate | Percentage | > 5% | CRITICAL |
| API Response Time (P50) | Milliseconds | > 500ms | WARNING |
| API Response Time (P95) | Milliseconds | > 2000ms | WARNING |
| API Response Time (P95) | Milliseconds | > 5000ms | CRITICAL |
| Database Query Time | Milliseconds | > 100ms | WARNING |
| Database Query Time | Milliseconds | > 500ms | CRITICAL |
| Audit Log Failures | Count | > 0 | CRITICAL |
| Authentication Failures | Percentage | > 10% | WARNING |
| Authentication Failures | Percentage | > 25% | CRITICAL |

### Business Metrics

| Metric | Type | Threshold | Alert Level |
| ------ | ---- | ---------- | ---------- |
| Schools Without Active Academic Year | Count | > 0 | INFO |
| Schools Without Active Academic Year | Count | > 5 | WARNING |
- | Subjects Without Category | Count | > 0 | INFO |
| Subjects Without Category | Count | > 10% | WARNING |
| CP Below 60% Alignment | Count | > 0 | INFO |
| CP Below 60% Alignment | Count | > 20% | WARNING |

### Database Metrics

| Metric | Type | Threshold | Alert Level |
| ------ | ---- | ---------- | ---------- |
| Database Connection Pool Usage | Percentage | > 80% | WARNING |
| Database Connection Pool Usage | Percentage | > 95% | CRITICAL |
| Database Slow Queries | Count | > 10/min | WARNING |
| Database Deadlocks | Count | > 0 | CRITICAL |
| Database Disk Usage | Percentage | > 80% | WARNING |
| Database Disk Usage | Percentage | > 90% | CRITICAL |

## Dashboard Configuration

### Grafana Dashboard Panels

1. **API Health Panel**
   - Request rate per endpoint
   - Error rate per endpoint
   - Response time percentiles (P50, P95, P99)
   - Active connections

2. **Database Health Panel**
   - Connection pool usage
   - Query performance
   - Slow query log
   - Table sizes

3. **Business Metrics Panel**
   - Schools with active academic years
   - Subject categorization coverage
   - CP alignment distribution
   - Graduate profile dimension coverage

4. **Audit Log Panel**
   - Audit log volume
   - Audit log failures
   - Audit log by entity type
   - Audit log by action type

### Log Aggregation

**Critical Logs to Aggregate**:
- All audit logs
- All error logs
- All authentication failures
- All authorization failures
- All database query failures
- All API errors (4xx, 5xx)

**Log Retention**:
- Application logs: 30 days
- Audit logs: 7 years (compliance)
- Error logs: 90 days

---

# PART 15 – CONFIGURATION MANAGEMENT

## Environment Variables

### New Environment Variables Required

```bash
# No new environment variables required for Sprint 4
# All configuration uses existing patterns
```

### Configuration Validation

Add to application startup validation:

```go
// Validate required configuration on startup
func ValidateConfig() error {
    // Existing validation...
    
    // Sprint 4 specific validation
    if os.Getenv("DB_HOST") == "" {
        return errors.New("DB_HOST is required")
    }
    
    // No new Sprint 4 environment variables needed
    
    return nil
}
```

## Feature Flags

### Sprint 4 Feature Flags

```go
const (
    FeatureAcademicYearManagement = "academic_year_management"
    FeatureSubjectCategorization = "subject_categorization"
    FeatureGraduateProfileDimensions = "graduate_profile_dimensions"
    FeatureCPAlignment = "cp_alignment"
)

// Feature flag configuration
var FeatureFlags = map[string]bool{
    FeatureAcademicYearManagement:     true,
    FeatureSubjectCategorization:     true,
    FeatureGraduateProfileDimensions:  true,
    FeatureCPAlignment:               true,
}
```

### Feature Flag Usage

```go
// In handlers
if !FeatureFlags[FeatureAcademicYearManagement] {
    c.JSON(503, gin.H{"error": "Feature not enabled"})
    return
}
```

---

# PART 16 – COMMUNICATION PLAN

## Stakeholder Communication

### Pre-Implementation Communication

**Audience**: School Admins, Curriculum Admins, Teachers

**Message**: "Sprint 4 Academic Foundation akan dilaksanakan mulai [tanggal]. Fitur baru mencakup manajemen tahun ajaran, kategori mata pelajaran, dan dimensi profil lulusan."

**Channels**:
- Email announcement
- In-app notification
- Portal banner

**Timing**: 2 weeks before deployment

---

### Deployment Communication

**Audience**: All users

**Message**: "Sprint 4 Academic Foundation telah diluncurkan pada [tanggal]. Sistem akan mengalami maintenance selama [durasi]."

**Channels**:
- Email announcement
- In-app notification
- Portal banner
- SMS (for critical notifications)

**Timing**: 48 hours before deployment

---

### Post-Deployment Communication

**Audience**: School Admins, Curriculum Admins

**Message**: "Sprint 4 telah berhasil diluncurkan. Silakan konfigurasi tahun ajaran dan kategorisasi mata pelajaran. Panduan tersedia di [link]."

**Channels**:
- Email announcement
- In-app notification
- Portal banner
- Training session (optional)

**Timing**: Immediately after deployment

---

### Training Materials

**User Guides** (Indonesian):
1. Panduan Konfigurasi Tahun Ajaran
2. Panduan Kategorisasi Mata Pelajaran
3. Panduan Dimensi Profil Lulusan
4. Panduan Aligment Kurikulum

**Video Tutorials**:
1. Cara membuat tahun ajaran baru
2. Cara mengategorikan mata pelajaran
3. Cara melakukan aligment CP

**FAQ**:
- Pertanyaan umum tentang Sprint 4
- Troubleshooting common issues
- Kontak support

---

# PART 17 – ROLLBACK PLAN

## Rollback Triggers

### Automatic Rollback Triggers

- Deployment script failure
- Database migration failure
- Critical error during smoke tests
- Health check failures > 5 consecutive attempts

### Manual Rollback Triggers

- Critical bugs discovered in production
- Performance degradation > 50%
- Data corruption detected
- Security vulnerability identified
- Stakeholder request rollback

## Rollback Timeline

| Phase | Duration | Actions |
| ---- | -------- | ------- |
| Decision | 5 minutes | Stakeholder approves rollback |
- | Stop Services | 2 minutes | Stop backend and frontend |
- | Database Rollback | 5 minutes | Apply rollback migration |
- | Code Rollback | 3 minutes | Deploy previous version |
- | Verification | 10 minutes | Smoke tests |
- | Communication | 5 minutes | Notify stakeholders |
- | **Total** | **30 minutes** | |

## Rollback Verification Checklist

- [ ] Database schema verified (matches pre-deployment)
- [ ] Data integrity verified (no data loss)
- [ ] Backend services healthy
- [ ] Frontend accessible
- [ ] Authentication working
- [ ] Critical workflows working
- [ ] No error spikes in logs
- [ ] Performance metrics normal
- [ ] Stakeholders notified
- [ ] Incident report created

## Rollback Communication

**Message Template**:

```
SUBJECT: ROLLBACK NOTICE - Sprint 4 Academic Foundation

Dikirim kepada: School Admins, Curriculum Admins, System Admins

Waktu: [timestamp]
Aksi: Rollback dari Sprint 4 ke versi sebelumnya
Alasan: [reason]
Estimasi downtime: 30 minutes

Status saat ini: [IN_PROGRESS/COMPLETED/FAILED]

Langkah berikutnya:
- [Untuk School Admin]: Silakan konfigurasi ulang tahun ajaran setelah sistem kembali online
- [Untuk Curriculum Admin]: Kategorisasi mata pelajaran akan dikembalikan ke status sebelumnya

Support: support@nusa.id
```

---

# PART 18 – POST-IMPLEMENTATION ACTIVITIES

## Week 1 Post-Deployment

### Monitoring and Stabilization

**Daily Activities**:
- Monitor error rates and response times
- Review audit logs for anomalies
- Address user-reported issues
- Fine-tune database indexes if needed
- Adjust alert thresholds based on actual traffic

**Deliverables**:
- Stability report
- Issue log and resolution status
- Performance optimization recommendations

### User Training

**Activities**:
- Conduct live training sessions for School Admins
- Conduct live training sessions for Curriculum Admins
- Record training sessions for on-demand viewing
- Distribute user guides
- Monitor user adoption

**Deliverables**:
- Training completion report
- User adoption metrics
- Training feedback summary

---

## Week 2 Post-Deployment

### Feature Enhancement

**Activities**:
- Prioritize backlog based on user feedback
- Implement quick wins (high impact, low effort)
- Address critical bugs
- Improve user experience based on feedback

**Deliverables**:
- Bug fix report
- Enhancement completion report
- User satisfaction survey

### Documentation Updates

**Activities**:
- Update API documentation with actual examples
- Update user guides based on feedback
- Update troubleshooting guides
- Update runbooks based on actual incidents

**Deliverables**:
- Updated documentation package
- Knowledge base articles

---

## Week 3 Post-Deployment

### Optimization

**Activities**:
- Analyze query performance
- Optimize slow queries
- Review and optimize indexes
- Implement caching for frequently accessed data (alignment reports)

**Deliverables**:
- Performance optimization report
- Cache strategy documentation

### Retrospective

**Activities**:
- Conduct team retrospective
- Document lessons learned
- Identify improvement opportunities
- Update best practices document

**Deliverables**:
- Retrospective report
- Lessons learned document
- Updated development process

---

# PART 19 – SUPPORT PROCEDURES

## Common Issues and Resolutions

### Issue 1: Academic Year Creation Fails

**Symptom**: Error when creating academic year

**Possible Causes**:
- Date range overlaps with existing academic year
- Start date is less than 30 days in future
- User lacks permissions

**Resolution Steps**:
1. Check error message for specific validation failure
2. Verify date range does not overlap
3. Verify start date is at least 30 days in future
4. Verify user has SCHOOL_ADMIN role and academic_year:CREATE permission

**Escalation**: If issue persists, check audit logs and contact support

---

### Issue 2: Semester Configuration Fails

**Symptom**: Cannot configure semesters

**Possible Causes**:
- Dates outside academic year range
- Semester sequence conflict
- Gap or overlap in semester dates

**Resolution Steps**:
1. Verify dates are within academic year range
2. Verify sequence is unique (1 and 2)
3. Use timeline visualization to identify gaps/overlaps
4. Adjust dates accordingly

**Escalation**: If issue persists, check audit logs and contact support

---

### Issue 3: Subject Category Deactivation Fails

**Symptom**: Cannot deactivate subject category

**Possible Causes**:
- Subjects are assigned to category
- User lacks permissions

**Resolution Steps**:
1. Count subjects assigned to category
2. Reassign subjects to different category
3. Retry deactivation
4. Verify user has CURRICULUM_ADMIN role

**Escalation**: If issue persists, check audit logs and contact support

---

### Issue 4: CP Alignment Calculation Incorrect

**Symptom**: Alignment percentage seems wrong

**Possible Causes**:
- Weight values incorrect
- Alignment strength mapping incorrect
- CP not aligned to enough dimensions

**Resolution Steps**:
1. Verify dimension weights (should sum to 1.0)
2. Verify alignment strength mapping (Strong=100%, Medium=75%, Weak=50%)
3. Check if CP has sufficient alignments
4. Recalculate manually for verification

**Escalation**: If calculation is definitely wrong, log bug and contact support

---

## Support Tiers

### Tier 1: Self-Service

**Scope**: Common issues with documented resolutions

**Response Time**: Immediate (user resolves using documentation)

**Channels**:
- Knowledge base
- FAQ
- User guides

---

### Tier 2: Admin Support

**Scope**: Issues requiring School Admin or Curriculum Admin intervention

**Response Time**: < 4 hours

**Channels**:
- In-app support request
- Email to support@nusa.id

**Examples**:
- Permission issues
- Configuration issues
- Data validation failures

---

### Tier 3: Technical Support

**Scope**: Technical issues requiring development team intervention

**Response Time**: < 24 hours (P0), < 48 hours (P1)

**Channels**:
- Email to support@nusa.id
- Phone (for P0 issues only)

**Examples**:
- System bugs
- Performance issues
- Data corruption
- Security incidents

---

## Support Contact Information

**Email**: support@nusa.id  
**Phone**: +62-21-XXXX-XXXX (P0 issues only, 24/7)  
**Emergency Contact**: on-call@nusa.id (for P0 incidents)

---

# PART 20 – SUCCESS CRITERIA DETAILED

## Functional Success Criteria

### Academic Year Management

- [ ] School Admin can create academic year with valid date range
- [ ] System prevents overlapping academic years for same school
- [ ] System requires 30-day lead time for new academic years
- [ ] School Admin can configure 2 semesters within academic year
- [ ] System validates semester coverage (no gaps, no overlaps)
- [ ] System enforces Ganjil before Genap sequence
- [ ] School Admin can submit academic year for approval
- [ ] System Admin can approve or reject academic year
- [ ] System automatically activates academic year on start date
- [ ] School Admin can deactivate academic year with approval
- [ ] All changes are logged in audit trail
- [ ] Past academic years are read-only
- [ ] Academic year status transitions are enforced (DRAFT → UNDER_REVIEW → APPROVED → ACTIVE → INACTIVE)

### Semester Management

- [ ] School Admin can create semester within academic year
- [ ] System enforces unique sequence within academic year
- [ ] System validates semester dates within academic year
- [ ] System prevents overlapping semesters
- [ ] School Admin can update semester in DRAFT status
- [ ] School Admin can delete semester in DRAFT status
- [ ] System automatically activates semesters on schedule
- [ ] All changes are logged in audit trail

### Subject Category Management

- [ ] Curriculum Admin can create subject category
- [ ] System enforces unique category codes and names
- [ ] Curriculum Admin can categorize subjects
- [ ] System supports bulk subject categorization
- [ ] System prevents category deletion with assigned subjects
- [ ] Curriculum Admin can deactivate unused categories
- [ ] All changes are logged in audit trail

### Graduate Profile Management

- [ ] Curriculum Admin can create graduate profile dimension
- [ ] System enforces unique dimension codes and names
- [ ] Curriculum Admin can configure exactly 8 dimensions
- [ ] Curriculum Admin can link CP to dimensions
- [ ] System calculates alignment percentage automatically
- [ ] Curriculum Admin can view alignment reports
- [ ] System identifies gaps in dimension coverage
- [ ] System prevents dimension deletion with aligned CP
- [ ] All changes are logged in audit trail

### Kurikulum Alignment Management

- [ ] Curriculum Admin can create CP-dimension alignment
- [ ] System prevents duplicate alignments
- [ ] System calculates alignment percentage automatically
- [ ] Curriculum Admin can update alignment strength
- [ ] Curriculum Admin can delete alignment
- [ ] System warns when CP falls below 60% alignment
- [ ] Curriculum Admin can generate alignment reports
- [ ] System visualizes alignment distribution
- [ ] Curriculum Admin can bulk align CP
- [ ] All changes are logged in audit trail

### Koding & AI Subject Support

- [ ] Curriculum Admin can add Koding subject
- [ ] Curriculum Admin can add AI subject
- [ ] Curriculum Admin can add Numerasi subject
- [ ] System categorizes new subjects as Intrakurikuler
- [ ] New subjects integrate with existing CP/TP workflow
- [ ] Curriculum Admin can categorize existing subjects
- [ ] All changes are logged in audit trail

## Technical Success Criteria

### Performance

- [ ] API response time (read) <500ms (95th percentile)
- [ ] API response time (write) <2s (95th percentile)
- [ ] Page load time <2s (95th percentile)
- [ ] Database query time <100ms (average)
- [ ] Alignment report generation <5s (for 1000 CP)

### Quality

- [ ] Unit test coverage >90% for critical paths
- [ ] All integration tests passing
- [ ] All API contract tests passing
- [ ] All E2E tests passing
- [ ] Zero critical bugs in production
- [ ] Zero security vulnerabilities

### Architecture

- [ ] No violations of Architecture Freeze v2
- [ ] No CQRS patterns introduced
- [ ] No Event Sourcing patterns introduced
- [ ] No new bounded contexts without approval
- [ ] Follows Handler → Service → Repository pattern
- [ ] No business logic in handlers
- [ ] Repository layer contains only database operations

### Security

- [ ] Permission-based access control working correctly
- [ ] Audit logging functional for all operations
- [ ] No SQL injection vulnerabilities
- [ ] No XSS vulnerabilities
- [ ] CSRF protection enabled
- [ ] Rate limiting configured
- [ ] Sensitive data not logged

## Business Success Criteria

### Adoption

- [ ] 100% of schools have active academic year configured within 30 days
- [ ] 100% of subjects categorized within 30 days
- [ ] 8 graduate profile dimensions configured and active
- [ ] CP alignment tracking operational within 30 days

### User Satisfaction

- [ ] UAT sign-off from all stakeholders
- [ ] User satisfaction score > 4.0/5.0
- [ ] Support ticket volume < 5% of user base
- [ ] Average issue resolution time < 24 hours

### Compliance

- [ ] Kurikulum 2026 compliance verified
- [ ] 8 Profil Lulusan dimensions match national standards
- [ ] Subject categorization follows Kurikulum Merdeka
- [ ] Data retention policies compliant with Indonesian laws

---

# DOCUMENT STATUS: FINAL

**Version**: 1.0  
**Date**: 2026-06-11  
**Status**: COMPLETE AND READY FOR IMPLEMENTATION  
**Next Action**: Submit for Architecture Board Review

---

This completes the comprehensive Sprint 4 Requirement Package. The document contains 20 parts covering all aspects needed for implementation, including detailed functional requirements, domain models, database design, API specifications, frontend requirements, security considerations, testing strategy, implementation plan, deployment procedures, monitoring, configuration management, communication plans, rollback procedures, post-implementation activities, support procedures, and success criteria.

The package is implementation-ready and provides sufficient detail for backend engineers, frontend engineers, QA engineers, and database engineers to begin coding immediately.
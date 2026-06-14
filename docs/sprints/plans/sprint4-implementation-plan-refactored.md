# Sprint 4 Requirement Package: Academic Foundation (Refactored)

**Document Version**: 2.0  
**Date**: 2026-06-11  
**Status**: Implementation Ready  
**Authors**: Principal Solution Architect, Product Owner, Education Domain Expert, DDD Reviewer, Technical Analyst

---

# Executive Summary

Sprint 4: Academic Foundation adalah sprint yang berfokus pada implementasi infrastruktur akademik dasar yang diperlukan untuk mendukung Kurikulum 2026 dan Deep Learning pedagogy. Sprint ini **TIDAK** mencakup Student Management, PPDB, atau Class Management - fitur tersebut ditangguh ke sprint berikutnya.

**Scope yang Difokuskan:**
- Academic Year Management (Tahun Ajaran)
- Semester Management
- Subject Category Management (Intrakurikuler, Kokurikuler, Ekstrakurikuler)
- Graduate Profile Dimension Management (8 Dimensi Profil Lulusan)
- CP Alignment Management
- Koding & AI Subject Support

**Out of Scope:**
- Student Management
- PPDB/Enrollment
- Class Management (Rombel, Wali Kelas)
- Attendance
- Scheduling
- Dapodik Integration

**Perubahan Utama dari Versi Sebelumnya:**
1. **Simplified Academic Year Workflow**: Dari Draft → Under Review → Approved → Active → Inactive menjadi Draft → Active → Archived (tanpa approval System Admin)
2. **Configurable Alignment Threshold**: CP alignment threshold 60% dijadikan configurable via system configuration
3. **Removed Dependencies**: Tidak ada dependensi ke PPDB, Student Management, atau Class Management

---

# PART 1 – GAP ANALYSIS

## Gap Analysis Matrix

| Requirement | Existing Module | Reuse | Extend | New | Rationale |
| ----------- | --------------- | ----- | ------ | --- | --------- |
| **1. Academic Year** | ❌ None | | | ✅ New | Tidak ada temporal scoping di struktur kurikulum saat ini |
| **2. Semester** | ❌ None | | | ✅ New | Tidak ada struktur semester; dibutuhkan untuk pembagian tahun ajaran |
| **3. Subject Category** | ⚠️ `curriculum_subjects` | | ✅ Extend | | Tabel subject ada tapi tidak ada kategorisasi (Intrakurikuler/Kokurikuler/Ekstrakurikuler) |
| **4. Profil Lulusan 8 Dimensi** | ❌ None | | | ✅ New | Tidak ada tracking dimensi profil lulusan |
| **5. Kurikulum 2026 Alignment** | ✅ `cp`, `tp`, `atp`, `modul_ajar` | ✅ Reuse | | | Modul kurikulum inti sudah sesuai Kurikulum Merdeka |
| **6. Koding** | ✅ `curriculum_subjects` | ✅ Reuse | | | Bisa tambahkan "Koding" sebagai subject baru |
| **7. AI** | ✅ `curriculum_subjects` | ✅ Reuse | | | Bisa tambahkan "AI" sebagai subject baru |
| **8. Numerasi** | ✅ `curriculum_subjects` | ✅ Reuse | | | Bisa tambahkan "Numerasi" sebagai subject baru |
| **9. Academic Structure Governance** | ✅ `users`, `roles`, `permissions` | ✅ Reuse | ✅ Extend | | Auth/authorization yang ada bisa diperluas dengan permission baru |

## Summary Statistics

- **Total Requirements**: 9
- **Fully Covered**: 3 (33%)
- **Extend Existing**: 2 (22%)
- **New Required**: 4 (44%)
- **New Tables Required**: 4 (academic_years, semesters, subject_categories, graduate_profile_dimensions)
- **Tables to Extend**: 1 (curriculum_subjects)
- **New Bounded Contexts Required**: 0 (Semua ekstensi masuk dalam Curriculum Context yang sudah ada)

---

# PART 2 – BUSINESS REQUIREMENTS DOCUMENT (BRD)

## Business Objectives

### Primary Objective
Membangun infrastruktur akademik dasar yang diperlukan untuk mendukung implementasi Kurikulum 2026 dan Deep Learning pedagogy, memungkinkan platform transisi dari manajemen konten kurikulum ke tata kelola akademik yang komprehensif.

### Secondary Objectives

1. **Temporal Scoping**: Implement manajemen tahun ajaran dan semester untuk memungkinkan perencanaan kurikulum, pengiriman, dan reporting dalam batas waktu yang terdefinisi.
2. **Subject Categorization**: Mengategorikan mata pelajaran sesuai komponen Kurikulum Merdeka (Intrakurikuler, Kokurikuler, Ekstrakurikuler) untuk mendukung pengiriman kurikulum yang komprehensif.
3. **Graduate Profile Tracking**: Implement framework Profil Lulusan 8 dimensi untuk memungkinkan alignment kurikulum dengan outcome lulusan nasional.
4. **Modern Subject Support**: Mendukung Koding, AI, dan Numerasi sebagai mata pelajaran formal sesuai Kurikulum 2026.
5. **Governance Framework**: Membangun tata kelola struktur akademik untuk memungkinkan curriculum admin mengelola konfigurasi akademik tanpa bantuan developer.

## Success Metrics

### Quantitative Metrics

| Metric | Target | Measurement Method |
| ------ | ------ | ------------------ |
| Academic Year Coverage | 100% sekolah memiliki tahun ajaran aktif dikonfigurasi | Query database academic_years aktif |
| Semester Configuration | 100% tahun ajaran memiliki 2 semester dikonfigurasi | Query database semester per academic_year |
| Subject Categorization | 100% subject dikategorikan | Query database subject dengan category_id |
| Profil Lulusan Dimensions | 8 dimensi dikonfigurasi dan aktif | Query database graduate_profile_dimensions aktif |
| Kurikulum Alignment | 100% CP align ke dimensi profil lulusan | Query database CP dengan alignment |
| Configuration Time | <5 menit untuk konfigurasi tahun ajaran | Time measurement test |
| Zero Downtime Migration | <30 detik untuk schema migration | Migration execution time |

### Qualitative Metrics

- School admin dapat mengonfigurasi struktur akademik secara mandiri tanpa support IT
- Curriculum admin dapat memverifikasi alignment Kurikulum 2026 via UI
- Guru dapat memahami kategorisasi subject dalam perencanaan kurikulum
- Audit trail ada untuk semua perubahan struktur akademik
- Perubahan konfigurasi dapat di-reverse dengan rollback capability

## Scope

### In Scope

1. **Academic Year Management**
   - Create, read, update, archive academic years
   - Set tanggal tahun ajaran
   - Tandai tahun ajaran aktif
   - **Simplified workflow**: Draft → Active → Archived (tanpa approval System Admin)

2. **Semester Management**
   - Create semester dalam tahun ajaran
   - Set tanggal semester
   - Configure urutan semester (Ganjil/Genap)
   - Semester activation/deactivation

3. **Subject Category Management**
   - Define Intrakurikuler, Kokurikuler, Ekstrakurikuler categories
   - Categorize subject yang ada dan baru
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
   - Define dimension descriptions dan indicators
   - Link dimensions ke CP untuk alignment tracking

5. **Kurikulum Alignment Management**
   - Link CP ke graduate profile dimensions
   - Track alignment percentage per CP
   - Generate alignment reports
   - Identify gaps dalam coverage kurikulum
   - **Configurable threshold**: Alignment threshold default 60% tapi dapat diubah via system configuration

6. **Koding & AI Subject Support**
   - Add Koding subject ke curriculum_subjects
   - Add AI subject ke curriculum_subjects
   - Add Numerasi subject ke curriculum_subjects
   - Categorize sebagai Intrakurikuler

### Out of Scope

1. **Student Management** - Student lifecycle, enrollment, records
2. **PPDB** - Workflow penerimaan siswa baru
3. **Class Management** - Rombel, wali kelas, scheduling
4. **Attendance** - Student attendance tracking
5. **Scheduling** - Class timetable, teacher assignments
6. **Grading** - Grade calculation, report cards
7. **Dapodik Integration** - External system synchronization
8. **Deep Learning Pedagogy** - Teaching methodology implementation
9. **Assessment Alignment** - Linking assessments ke graduate profiles (future sprint)
10. **Approval Workflow** - System Admin approval untuk academic year (dihapus di v2)

## Stakeholders

### Primary Stakeholders

| Stakeholder | Role | Responsibilities |
| ----------- | ---- | --------------- |
| School Admin | Configure academic year/semester untuk school | School-level academic configuration |
| Curriculum Admin | Manage subject categories, Profil Lulusan, alignment | System-wide curriculum governance |
| Teacher | View academic structure, understand subject categorization | Curriculum planning context |

### Secondary Stakeholders

| Stakeholder | Role | Responsibilities |
| ----------- | ---- | --------------- |
| Ministry of Education | Compliance dengan Kurikulum 2026 | Regulatory compliance |

## Business Processes

### Process 1: Academic Year Configuration (Simplified)

**Actors**: School Admin

**Flow**:
1. School Admin membuka halaman Academic Year Management
2. School Admin klik "Create New Academic Year"
3. School Admin isi form:
   - Academic year name (misal: "2026/2027")
   - Start date
   - End date
   - Description (optional)
4. System validasi:
   - Start date < end date
   - Tidak ada overlap dengan academic year yang ada
   - Start date >= today + 30 hari
5. System buat academic year dalam status DRAFT
6. School Admin configure semesters (ganjil/genap)
7. System validasi semester coverage
8. School Admin klik "Activate"
9. System ubah status academic year ke ACTIVE
10. Academic year aktif pada start date yang ditentukan

**Business Rules**:
- Hanya satu academic year bisa aktif pada satu waktu per school
- Academic year date ranges tidak boleh overlap
- Academic years harus dikonfigurasi minimal 30 hari sebelum start
- Academic year masa lalu read-only (status ARCHIVED)
- **TIDAK ada approval System Admin** (simplified di v2)

### Process 2: Subject Categorization

**Actors**: Curriculum Admin

**Flow**:
1. Curriculum Admin buat subject category
2. Curriculum Admin define category metadata
3. Curriculum Admin categorize subjects
4. Curriculum Admin publish categorization
5. Teachers view categorized subjects dalam curriculum planning

**Business Rules**:
- Setiap subject harus belong ke exactly satu category
- Categories tidak bisa dihapus jika subjects assigned
- Historical categorization di-preserve

### Process 3: Profil Lulusan Configuration

**Actors**: Curriculum Admin

**Flow**:
1. Curriculum Admin configure 8 graduate profile dimensions
2. Curriculum Admin define dimension indicators
3. Curriculum Admin link CP ke dimensions
4. System generate alignment reports
5. Curriculum Admin review gaps
6. Curriculum Admin adjust CP-dimension links

**Business Rules**:
- CP harus align ke minimal 1 graduate profile dimension
- CP bisa align ke multiple dimensions
- Dimension configuration adalah system-wide (tidak per school)
- Alignment percentage dihitung otomatis
- **Alignment threshold configurable** (default 60% tapi bisa diubah via config)

## Business Rules

### BR-001: Academic Year Uniqueness
Setiap school hanya bisa memiliki satu academic year aktif pada satu waktu.

### BR-002: Academic Year Non-Overlap
Academic year date ranges untuk satu school tidak boleh overlap dengan academic year yang sudah ada.

### BR-003: Academic Year Lead Time
Academic year baru harus dibuat minimal 30 hari sebelum start date.

### BR-004: Semester Sequence
Setiap academic year harus memiliki tepat 2 semester berurutan: Ganjil (ganjil) diikuti Genap (genap).

### BR-005: Semester Date Coverage
Semester date ranges harus fully cover academic year date range tanpa gaps atau overlaps.

### BR-006: Subject Category Exclusivity
Setiap subject harus belong ke exactly satu subject category (Intrakurikuler, Kokurikuler, atau Ekstrakurikuler).

### BR-007: Profil Lulusan Completeness
CP harus align ke minimal 1 dari 8 graduate profile dimensions.

### BR-008: Kurikulum Alignment Minimum (Configurable)
Setiap CP harus memiliki minimum alignment score ke graduate profile dimensions. Threshold ini **configurable** via system configuration, default value adalah 60%.

### BR-009: Modern Subject Support
Koding, AI, dan Numerasi subjects harus dikategorikan sebagai Intrakurikuler.

### BR-010: Governance Self-Service
School admin dapat mengonfigurasi academic year secara mandiri tanpa approval dari System Admin (simplified di v2).

## Assumptions

### Technical Assumptions

1. PostgreSQL 18+ dengan UUID v7 support tersedia
2. Authentication dan authorization infrastructure yang ada akan di-reuse
3. Audit logging infrastructure yang ada akan di-reuse
4. Database migrations akan mengikuti pattern yang ada (up/down SQL files)
5. API akan mengikuti REST pattern dan naming conventions yang sudah ada

### Business Assumptions

1. Schools mengoperasikan pada struktur tahun ajaran 2-semester standar
2. Academic years mengikuti struktur tahun kalender (misal: 2026/2027)
3. Subject categorization mengikuti definisi standar Kurikulum Merdeka
4. Graduate profile dimensions konsisten di semua schools (national standard)
5. School admin memiliki authority untuk mengonfigurasi academic years untuk schools mereka

### Integration Assumptions

1. Dapodik integration ditunda ke future sprint
2. External calendar systems (Google Calendar, dll) tidak akan diintegrasikan di Sprint 4
3. Tidak ada real-time synchronization dengan external systems

## Constraints

### Technical Constraints

1. **Architecture Compliance**: Harus mengikuti Architecture Freeze v2 (tidak ada CQRS, tidak ada Event Sourcing, tidak ada new bounded contexts tanpa approval)
2. **Database Schema**: Harus mengikuti Database Schema Freeze v1 patterns (UUID primary keys, audit fields, soft delete via status)
3. **API Standards**: Harus mengikuti REST API pattern yang sudah ada (versioned endpoints, consistent error codes)
4. **Frontend Stack**: Harus menggunakan React/TypeScript stack dengan MUI components yang sudah ada
5. **Performance**: API response time harus <500ms untuk read operations, <2s untuk write operations

### Business Constraints

1. **Kurikulum Compliance**: Harus align dengan Kurikulum Merdeka 2026 requirements
2. **National Standards**: Graduate profile dimensions harus match Permendikdasmen No. 10 Tahun 2025
3. **School Autonomy**: Schools harus memiliki autonomy untuk mengonfigurasi academic years dalam national guidelines
4. **Data Privacy**: Tidak ada student personal data yang akan disimpan di Sprint 4 (ditunda ke future sprint)

### Resource Constraints

1. **Timeline**: Sprint 4 durasi 4 minggu (dikurangi dari 5 minggu karena simplifikasi workflow)
2. **Team**: Solo developer dengan waktu terbatas
3. **Budget**: Tidak ada external dependencies atau paid services
4. **Infrastructure**: Single PostgreSQL instance, tidak ada distributed systems

### Regulatory Constraints

1. **Indonesian Language**: Semua teks UI harus dalam Bahasa Indonesia
2. **Data Sovereignty**: Semua data harus reside di Indonesia
3. **Education Standards**: Harus comply dengan SNP (Standar Nasional Pendidikan)

---

# PART 3 – FUNCTIONAL REQUIREMENTS

## Feature 1: Academic Year Management (Simplified)

### Description
Enable school administrators untuk create, configure, dan manage academic years dengan defined date ranges dan semester divisions. **Workflow di-sederhanakan** untuk MVP: Draft → Active → Archived (tanpa approval System Admin).

### Actors
- School Admin (primary)
- Curriculum Admin (read access)
- Teacher (read access)

### Preconditions
- User must be authenticated
- User must have SCHOOL_ADMIN role (untuk create/update)
- School harus exist dan aktif

### Main Flow

**1. Create Academic Year**
1. School Admin navigasi ke halaman Academic Year Management
2. School Admin klik "Create New Academic Year"
3. System display academic year creation form
4. School Admin isi:
   - Academic year name (misal: "2026/2027")
   - Start date
   - End date
   - Description (optional)
5. System validasi:
   - Start date < end date
   - Tidak ada overlap dengan academic year yang ada untuk school
   - Start date >= today + 30 hari
6. System buat academic year dalam status DRAFT
7. System create audit log entry
8. System return success response dengan academic year ID

**2. Configure Semesters**
1. School Admin pilih academic year yang dibuat
2. School Admin klik "Configure Semesters"
3. System display semester configuration form
4. School Admin isi untuk Semester Ganjil:
   - Name (default: "Semester Ganjil")
   - Start date
   - End date
5. School Admin isi untuk Semester Genap:
   - Name (default: "Semester Genap")
   - Start date
   - End date
6. System validasi:
   - Semester dates dalam academic year range
   - Tidak ada gaps antara semester
   - Tidak ada overlaps antara semester
   - Ganjil sebelum Genap
7. System buat semesters dalam status DRAFT
8. System create audit log entry
9. System return success response

**3. Activate Academic Year**
1. School Admin klik "Activate" pada academic year
2. System validasi:
   - Academic year memiliki 2 semester dikonfigurasi
   - Semua required fields populated
   - Tidak ada academic year lain yang sedang aktif untuk school
3. System ubah academic year status ke ACTIVE
4. System create audit log entry
5. System return success response

**4. Archive Academic Year**
1. School Admin pilih academic year yang bukan lagi aktif
2. School Admin klik "Archive"
3. System validasi:
   - Academic year status bukan ACTIVE
   - Tidak ada active curriculum planning yang sedang berlangsung
4. System ubah academic year status ke ARCHIVED
5. System create audit log entry
6. System return success response

### Alternative Flow

**A1: Update Academic Year**
1. School Admin pilih academic year dalam status DRAFT
2. School Admin klik "Edit"
3. System display academic year edit form
4. School Admin modify fields
5. System validasi updated values
6. System update academic year
7. System create audit log entry
8. System return success response

### Error Flow

**E1: Overlapping Date Range**
1. School Admin isi academic year dates yang overlap dengan academic year yang ada
2. System return validation error: "Academic year dates tidak boleh overlap dengan academic year yang sudah ada"
3. System highlight conflicting academic year
4. School Admin adjust dates

**E2: Insufficient Lead Time**
1. School Admin isi start date kurang dari 30 hari di future
2. System return validation error: "Academic year harus dibuat minimal 30 hari sebelum start date"
3. School Admin adjust start date

**E3: Semester Gap atau Overlap**
1. School Admin isi semester dates dengan gap atau overlap
2. System return validation error: "Semesters harus fully cover academic year tanpa gaps atau overlaps"
3. System visualisasi gap/overlap
4. School Admin adjust semester dates

### Validation Rules

| Field | Validation | Error Message |
| ----- | ---------- | ------------- |
| name | Required, max 100 characters | "Academic year name required dan harus kurang dari 100 characters" |
| start_date | Required, date format, >= today + 30 days | "Start date required dan harus minimal 30 hari di future" |
| end_date | Required, date format, > start_date | "End date required dan harus setelah start date" |
| description | Optional, max 500 characters | "Description harus kurang dari 500 characters" |
| semester.name | Required, max 50 characters | "Semester name required" |
| semester.start_date | Required, date format, >= academic_year.start_date | "Semester start date harus dalam academic year" |
| semester.end_date | Required, date format, <= academic_year.end_date | "Semester end date harus dalam academic year" |

### Acceptance Criteria

**AC-001**: School Admin bisa buat academic year dengan date range yang valid
**AC-002**: System prevent overlapping academic years untuk school yang sama
**AC-003**: System require 30-day lead time untuk academic year baru
**AC-004**: School Admin bisa configure 2 semester dalam academic year
**AC-005**: System validasi semester coverage (tidak ada gaps, tidak ada overlaps)
**AC-006**: System enforce Ganjil sebelum Genap sequence
**AC-007**: School Admin bisa activate academic year tanpa approval System Admin (simplified)
**AC-008**: School Admin bisa archive academic year yang sudah selesai
**AC-009**: Semua perubahan di-log dalam audit trail
**AC-010**: Academic year masa lalu read-only (status ARCHIVED)
**AC-011**: **TIDAK ada approval workflow System Admin** (simplified di v2)

---

## Feature 2: Semester Management

### Description
Enable school administrators untuk create, configure, dan manage semesters dalam academic years dengan date ranges, sequence validation, dan activation workflows.

### Actors
- School Admin (primary)
- Curriculum Admin (read access)
- Teacher (read access)

### Preconditions
- User must be authenticated
- User must have SCHOOL_ADMIN role
- Academic year harus exist dan dalam status DRAFT

### Main Flow

**1. Create Semester**
1. School Admin navigasi ke Academic Year detail page
2. School Admin klik "Add Semester"
3. System display semester creation form
4. School Admin isi:
   - Semester name (misal: "Semester Ganjil")
   - Sequence (1 untuk Ganjil, 2 untuk Genap)
   - Start date
   - End date
   - Description (optional)
5. System validasi:
   - Sequence unique dalam academic year
   - Dates dalam academic year range
   - Tidak ada overlap dengan semester yang ada
6. System buat semester dalam status DRAFT
7. System link semester ke academic year
8. System create audit log entry
9. System return success response

**2. Update Semester**
1. School Admin pilih semester yang ada
2. School Admin klik "Edit"
3. System display semester edit form
4. School Admin modify fields
5. System validasi updated values
6. System update semester
7. System create audit log entry
8. System return success response

**3. Activate Semester**
1. Academic year diactivate
2. System otomatis activate first semester pada academic year start date
3. System otomatis activate second semester pada first semester end date
4. System create audit log entry
5. System kirim notification ke School Admin

### Alternative Flow

**A1: Delete Semester**
1. School Admin pilih semester dalam status DRAFT
2. School Admin klik "Delete"
3. System require confirmation
4. System validasi:
   - Tidak ada linked curriculum planning
   - Belum di-activate
5. System delete semester
6. System create audit log entry
7. System return success response

### Error Flow

**E1: Invalid Sequence**
1. School Admin isi sequence yang konflik dengan semester yang ada
2. System return validation error: "Sequence harus unique dalam academic year"
3. School Admin adjust sequence

**E2: Date Range Violation**
1. School Admin isi dates di luar academic year range
2. System return validation error: "Semester dates harus dalam academic year range"
3. School Admin adjust dates

### Validation Rules

| Field | Validation | Error Message |
| ----- | ---------- | ------------- |
| name | Required, max 50 characters | "Semester name required" |
| sequence | Required, integer 1 atau 2 | "Sequence harus 1 atau 2" |
| start_date | Required, date format, >= academic_year.start_date | "Start date harus dalam academic year" |
| end_date | Required, date format, <= academic_year.end_date | "End date harus dalam academic year" |
| description | Optional, max 500 characters | "Description harus kurang dari 500 characters" |

### Acceptance Criteria

**AC-012**: School Admin bisa buat semester dalam academic year
**AC-013**: System enforce unique sequence dalam academic year
**AC-014**: System validasi semester dates dalam academic year
**AC-015**: System prevent overlapping semesters
**AC-016**: School Admin bisa update semester dalam status DRAFT
**AC-017**: School Admin bisa delete semester dalam status DRAFT
**AC-018**: System otomatis activate semesters pada schedule
**AC-019**: Semua perubahan di-log dalam audit trail

---

## Feature 3: Subject Category Management

### Description
Enable curriculum administrators untuk define dan manage subject categories (Intrakurikuler, Kokurikuler, Ekstrakurikuler) dan categorize subjects sesuai Kurikulum Merdeka standards.

### Actors
- Curriculum Admin (primary)
- School Admin (read access)
- Teacher (read access)

### Preconditions
- User must be authenticated
- User must have CURRICULUM_ADMIN role untuk create/update
- User must memiliki appropriate read permission

### Main Flow

**1. Create Subject Category**
1. Curriculum Admin navigasi ke halaman Subject Category Management
2. Curriculum Admin klik "Create Category"
3. System display category creation form
4. Curriculum Admin isi:
   - Category code (misal: "INTRAKURIKULER")
   - Category name (misal: "Intrakurikuler")
   - Category name (English, optional)
   - Description
   - Guidelines
5. System validasi:
   - Code adalah unique
   - Name adalah unique
6. System buat category dalam status ACTIVE
7. System create audit log entry
8. System return success response

**2. Categorize Subject**
1. Curriculum Admin navigasi ke halaman Subject Management
2. Curriculum Admin pilih subject
3. Curriculum Admin klik "Edit"
4. System display subject edit form
5. Curriculum Admin pilih category dari dropdown
6. System simpan category assignment
7. System create audit log entry
8. System return success response

**3. Bulk Categorize Subjects**
1. Curriculum Admin navigasi ke halaman Subject Category Management
2. Curriculum Admin pilih category
3. Curriculum Admin klik "Assign Subjects"
4. System display subject selection list
5. Curriculum Admin pilih multiple subjects
6. System assign selected subjects ke category
7. System create audit log entry
8. System return success response

### Alternative Flow

**A1: Deactivate Category**
1. Curriculum Admin pilih category
2. Curriculum Admin klik "Deactivate"
3. System validasi:
   - Tidak ada subjects yang sedang assigned
4. System ubah category status ke INACTIVE
5. System create audit log entry
6. System return success response

### Error Flow

**E1: Category dengan Subjects**
1. Curriculum Admin mencoba deactivate category dengan assigned subjects
2. System return error: "Tidak bisa deactivate category dengan assigned subjects"
3. System display count dari assigned subjects
4. Curriculum Admin reassign subjects terlebih dahulu

### Validation Rules

| Field | Validation | Error Message |
| ----- | ---------- | ------------- |
| code | Required, unique, uppercase, max 50 characters | "Category code required dan harus unique" |
| name | Required, unique, max 100 characters | "Category name required dan harus unique" |
| name_en | Optional, max 100 characters | "English name harus kurang dari 100 characters" |
| description | Required, max 1000 characters | "Description required" |
| guidelines | Optional, max 2000 characters | "Guidelines harus kurang dari 2000 characters" |

### Acceptance Criteria

**AC-020**: Curriculum Admin bisa buat subject category
**AC-021**: System enforce unique category codes dan names
**AC-022**: Curriculum Admin bisa categorize subjects
**AC-023**: System support bulk subject categorization
**AC-024**: System prevent category deletion dengan assigned subjects
**AC-025**: Curriculum Admin bisa deactivate unused categories
**AC-026**: Semua perubahan di-log dalam audit trail

---

## Feature 4: Profil Lulusan Management

### Description
Enable curriculum administrators untuk configure 8-dimensional Profil Lulusan framework dengan descriptions, indicators, dan alignment tracking ke Kurikulum 2026 graduate outcomes.

### Actors
- Curriculum Admin (primary)
- School Admin (read access)
- Teacher (read access)

### Preconditions
- User must be authenticated
- User must have CURRICULUM_ADMIN role untuk create/update
- User must memiliki appropriate read permission

### Main Flow

**1. Create Graduate Profile Dimension**
1. Curriculum Admin navigasi ke halaman Profil Lulusan Management
2. Curriculum Admin klik "Create Dimension"
3. System display dimension creation form
4. Curriculum Admin isi:
   - Dimension code (misal: "KEIMANAN_KETAKWAAN")
   - Dimension name (misal: "Keimanan & Ketakwaan")
   - Dimension name (English, optional)
   - Description
   - Indicators (JSON array)
   - Weight (untuk alignment calculation)
5. System validasi:
   - Code adalah unique
   - Name adalah unique
   - Weight adalah positive
6. System buat dimension dalam status ACTIVE
7. System create audit log entry
8. System return success response

**2. Configure 8 Dimensions**
1. Curriculum Admin buat semua 8 dimensions:
   - Keimanan & Ketakwaan
   - Kewargaan
   - Berakhlak Mulia
   - Berani Bertanggung Jawab
   - Peduli
   - Gotong Royong
   - Mandiri
   - Kreatif
2. System validasi bahwa exactly 8 dimensions exist
3. System validasi bahwa semua dimensions aktif
4. System publish dimension configuration
5. System create audit log entry
6. System return success response

**3. Link CP ke Dimensions**
1. Curriculum Admin navigasi ke halaman CP Management
2. Curriculum Admin pilih CP
3. Curriculum Admin klik "Align ke Profil Lulusan"
4. System display dimension selection form
5. Curriculum Admin pilih satu atau lebih dimensions
6. Curriculum Admin assign alignment strength (Strong, Medium, Weak)
7. System simpan alignment
8. System calculate alignment percentage
9. System create audit log entry
10. System return success response

**4. View Alignment Report**
1. Curriculum Admin navigasi ke halaman Alignment Report
2. System select academic year dan phase
3. System generate alignment report showing:
   - CP count per dimension
   - Alignment percentage per dimension
   - Gaps dalam dimension coverage
4. System visualisasi alignment distribution
5. Curriculum Admin export report sebagai PDF

### Alternative Flow

**A1: Update Dimension**
1. Curriculum Admin pilih dimension yang ada
2. Curriculum Admin klik "Edit"
3. System display dimension edit form
4. Curriculum Admin update fields
5. System validasi updates
6. System update dimension
7. System create audit log entry
8. System return success response

**A2: Deactivate Dimension**
1. Curriculum Admin pilih dimension
2. Curriculum Admin klik "Deactivate"
3. System validasi:
   - Tidak ada CP yang sedang align ke dimension
4. System ubah dimension status ke INACTIVE
5. System create audit log entry
6. System return success response

### Error Flow

**E1: Dimension dengan Aligned CP**
1. Curriculum Admin mencoba deactivate dimension dengan aligned CP
2. System return error: "Tidak bisa deactivate dimension dengan aligned CP"
3. System display count dari aligned CP
4. Curriculum Admin remove alignments terlebih dahulu

### Validation Rules

| Field | Validation | Error Message |
| ----- | ---------- | ------------- |
| code | Required, unique, uppercase, max 50 characters | "Dimension code required dan harus unique" |
| name | Required, unique, max 100 characters | "Dimension name required dan harus unique" |
| name_en | Optional, max 100 characters | "English name harus kurang dari 100 characters" |
| description | Required, max 1000 characters | "Description required" |
| indicators | Required, valid JSON array | "Indicators harus valid JSON array" |
| weight | Required, positive decimal, max 1.0 | "Weight harus positive dan <= 1.0" |

### Acceptance Criteria

**AC-027**: Curriculum Admin bisa buat graduate profile dimension
**AC-028**: System enforce unique dimension codes dan names
**AC-029**: Curriculum Admin bisa configure exactly 8 dimensions
**AC-030**: Curriculum Admin bisa link CP ke dimensions
**AC-031**: System calculate alignment percentage otomatis
**AC-032**: Curriculum Admin bisa view alignment reports
**AC-033**: System identify gaps dalam dimension coverage
**AC-034**: System prevent dimension deletion dengan aligned CP
**AC-035**: Semua perubahan di-log dalam audit trail

---

## Feature 5: Kurikulum Alignment Management

### Description
Enable curriculum administrators untuk track dan manage alignment antara CP (Capaian Pembelajaran) dan graduate profile dimensions, memastikan curriculum coverage dari semua 8 dimensions dengan quantitative alignment scoring.

### Actors
- Curriculum Admin (primary)
- School Admin (read access)
- Teacher (read access)

### Preconditions
- User must be authenticated
- User must have CURRICULUM_ADMIN role untuk manage alignment
- User must memiliki appropriate read permission
- Graduate profile dimensions harus dikonfigurasi
- CP harus exist

### Main Flow

**1. Create Alignment**
1. Curriculum Admin navigasi ke CP detail page
2. Curriculum Admin klik "Add Alignment"
3. System display alignment creation form
4. Curriculum Admin pilih:
   - Graduate profile dimension
   - Alignment strength (Strong = 100%, Medium = 75%, Weak = 50%)
   - Rationale (optional)
5. System validasi:
   - CP belum align ke dimension yang sama
6. System buat alignment
7. System recalculate CP overall alignment percentage
8. System create audit log entry
9. System return success response

**2. Update Alignment**
1. Curriculum Admin pilih alignment yang ada
2. Curriculum Admin klik "Edit"
3. System display alignment edit form
4. Curriculum Admin modify alignment strength
5. System recalculate CP overall alignment percentage
6. System update alignment
7. System create audit log entry
8. System return success response

**3. Delete Alignment**
1. Curriculum Admin pilih alignment
2. Curriculum Admin klik "Delete"
3. System require confirmation
4. System delete alignment
5. System recalculate CP overall alignment percentage
6. System create audit log entry
7. System return success response

**4. Generate Alignment Report**
1. Curriculum Admin navigasi ke halaman Alignment Report
2. Curriculum Admin select filters:
   - Academic year
   - Phase
   - Subject
3. System generate alignment report showing:
   - Total CP count
   - CP aligned count
   - Overall alignment percentage
   - Alignment per dimension
   - CP **dibawah konfigurasi threshold** (default 60% tapi configurable)
4. System visualisasi data dengan charts
5. Curriculum Admin export report sebagai PDF/CSV

### Alternative Flow

**A1: Bulk Align CP**
1. Curriculum Admin navigasi ke halaman CP Management
2. Curriculum Admin pilih multiple CP
3. Curriculum Admin klik "Bulk Align"
4. System display bulk alignment form
5. Curriculum Admin pilih dimension(s) untuk semua selected CP
6. Curriculum Admin set default alignment strength
7. System buat alignments untuk semua selected CP
8. System recalculate alignment percentages
9. System create audit log entry
10. System return success response

### Error Flow

**E1: Minimum Alignment Violation (Configurable Threshold)**
1. Curriculum Admin delete alignment menyebabkan CP fall **dibawah konfigurasi threshold** (default 60% tapi configurable)
2. System return warning: "CP alignment akan fall di bawah [threshold]% threshold. Lanjutkan?"
3. System require confirmation
4. Curriculum Admin confirm atau adjust

### Validation Rules

| Field | Validation | Error Message |
| ----- | ---------- | ------------- |
| dimension_id | Required, exists | "Dimension required" |
| alignment_strength | Required, enum (STRONG, MEDIUM, WEAK) | "Alignment strength required" |
| rationale | Optional, max 500 characters | "Rationale harus kurang dari 500 characters" |

### Acceptance Criteria

**AC-036**: Curriculum Admin bisa buat CP-dimension alignment
**AC-037**: System prevent duplicate alignments
**AC-038**: System calculate alignment percentage otomatis
**AC-039**: Curriculum Admin bisa update alignment strength
**AC-040**: Curriculum Admin bisa delete alignment
**AC-041**: System warns ketika CP fall **dibawah konfigurasi threshold** (configurable, default 60%)
**AC-042**: Curriculum Admin bisa generate alignment reports
**AC-043**: System visualisasi alignment distribution
**AC-044**: Curriculum Admin bisa bulk align CP
**AC-045**: Semua perubahan di-log dalam audit trail
**AC-046**: **Alignment threshold configurable** via system configuration, default 60%

---

## Feature 6: Koding & AI Subject Support

### Description
Enable platform untuk support Koding, AI, dan Numerasi sebagai formal subjects dalam kurikulum structure, properly categorized dan integrated dengan existing curriculum planning workflows.

### Actors
- Curriculum Admin (primary)
- School Admin (read access)
- Teacher (read access)

### Preconditions
- User must be authenticated
- User must have CURRICULUM_ADMIN role untuk create subjects
- User must memiliki appropriate read permission
- Subject categories harus dikonfigurasi

### Main Flow

**1. Add Koding Subject**
1. Curriculum Admin navigasi ke halaman Subject Management
2. Curriculum Admin klik "Create Subject"
3. System display subject creation form
4. Curriculum Admin isi:
   - Code: "KODING"
   - Name: "Koding"
   - Name (English): "Coding"
   - Description: "Mata pelajaran koding dan pemrograman"
   - Category: Intrakurikuler
   - Is Active: true
5. System validasi:
   - Code adalah unique
   - Name adalah unique
   - Category ada
6. System buat Koding subject
7. System create audit log entry
8. System return success response

**2. Add AI Subject**
1. Curriculum Admin repeat process untuk AI subject
2. Curriculum Admin isi:
   - Code: "AI"
   - Name: "Kecerdasan Buatan"
   - Name (English): "Artificial Intelligence"
   - Description: "Mata pelajaran kecerdasan buatan"
   - Category: Intrakurikuler
   - Is Active: true
3. System buat AI subject
4. System create audit log entry
5. System return success response

**3. Add Numerasi Subject**
1. Curriculum Admin repeat process untuk Numerasi subject
2. Curriculum Admin isi:
   - Code: "NUMERASI"
   - Name: "Numerasi"
   - Name (English): "Numeracy"
   - Description: "Mata pelajaran numerasi"
   - Category: Intrakurikuler
   - Is Active: true
3. System buat Numerasi subject
4. System create audit log entry
5. System return success response

**4. Verify Integration**
1. Curriculum Admin navigasi ke halaman CP Management
2. Curriculum Admin pilih Koding subject
3. Curriculum Admin buat CP untuk Koding
4. System verifikasi CP creation berhasil
5. Curriculum Admin buat TP untuk Koding CP
6. System verifikasi TP creation berhasil
7. System konfirmasi integration sedang bekerja

### Alternative Flow

**A1: Categorize Existing Subjects**
1. Curriculum Admin navigasi ke halaman Subject Management
2. Curriculum Admin pilih subject yang ada (misal: Matematika)
3. Curriculum Admin klik "Edit"
4. Curriculum Admin pilih category (misal: Intrakurikuler)
5. System simpan category assignment
6. System create audit log entry
7. System return success response

### Error Flow

**E1: Duplicate Subject**
1. Curriculum Admin mencoba buat subject dengan code yang sudah ada
2. System return error: "Subject code sudah ada"
3. Curriculum Admin adjust code

### Validation Rules

| Field | Validation | Error Message |
| ----- | ---------- | ------------- |
| code | Required, unique, uppercase, max 50 characters | "Subject code required dan harus unique" |
| name | Required, unique, max 255 characters | "Subject name required dan harus unique" |
| name_en | Optional, max 255 characters | "English name harus kurang dari 255 characters" |
| description | Required, max 1000 characters | "Description required" |
| category_id | Required, exists | "Category required" |

### Acceptance Criteria

**AC-047**: Curriculum Admin bisa tambah Koding subject
**AC-048**: Curriculum Admin bisa tambah AI subject
**AC-049**: Curriculum Admin bisa tambah Numerasi subject
**AC-050**: System categorize new subjects sebagai Intrakurikuler
**AC-051**: New subjects integrate dengan CP/TP workflow yang sudah ada
**AC-052**: Curriculum Admin bisa categorize subjects yang sudah ada
**AC-053**: Semua perubahan di-log dalam audit trail

---

# PART 4 – DOMAIN MODEL

## Bounded Context Decision

### Decision: NO NEW BOUNDED CONTEXTS

**Rationale**:
1. Semua Sprint 4 requirements adalah extensions ke Curriculum Context yang sudah ada
2. Academic years, semesters, subject categories, dan graduate profile dimensions adalah reference data untuk curriculum planning
3. Entities ini tidak merepresent business domain baru - mereka adalah infrastructure untuk domain kurikulum yang sudah ada
4. Mengikuti prinsip DDD: "Bounded contexts harus berdasarkan business capabilities, bukan technical entities"
5. Menambahkan new bounded contexts akan melanggar "minimize new bounded contexts" constraint

**Conclusion**:
Semua fitur Sprint 4 akan diimplement dalam Curriculum Context yang sudah ada sebagai extensions ke reference data dan governance structures.

---

## Aggregates

### Aggregate 1: AcademicYear

**Purpose**: Manage academic year configuration dengan semesters

**Aggregate Root**: `AcademicYear`

**Entities**:
- `AcademicYear` (Aggregate Root)
- `Semester` (Entity)

**Value Objects**:
- `DateRange` (Value Object)
- `AcademicYearName` (Value Object)

**Domain Services**:
- `AcademicYearValidationService` - Validasi academic year rules (tidak ada overlap, 30-day lead time)
- `AcademicYearActivationService` - Manage otomatisasi activation scheduling

**Repository**: `AcademicYearRepository`

**Domain Events**: None (tidak ada cross-aggregate communication required)

### Aggregate 2: SubjectCategory

**Purpose**: Manage subject categorization sesuai Kurikulum Merdeka

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

**Purpose**: Manage graduate profile dimensions dan CP alignment

**Aggregate Root**: `GraduateProfileDimension`

**Entities**:
- `GraduateProfileDimension` (Aggregate Root)
- `CPAlignment` (Entity)

**Value Objects**:
- `DimensionCode` (Value Object)
- `DimensionIndicators` (Value Object)
- `AlignmentStrength` (Value Object)
- **AlignmentThreshold** (Value Object) - **NEW: Configurable threshold**

**Domain Services**:
- `AlignmentCalculationService` - Calculate CP alignment percentage dengan configurable threshold
- `ConfigurationService` - **NEW: Read system configuration untuk alignment threshold**

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
    CreatedBy       UUID
    CreatedAt       time.Time
    UpdatedAt       time.Time
    
    // Child entities
    Semesters       []Semester
}

type AcademicYearStatus string
const (
    AcademicYearStatusDraft   AcademicYearStatus = "DRAFT"
    AcademicYearStatusActive  AcademicYearStatus = "ACTIVE"
    AcademicYearStatusArchived AcademicYearStatus = "ARCHIVED"
)

// Removed: UnderReview, Approved, ApprovedBy, ApprovedAt
// Simplified workflow: DRAFT → ACTIVE → ARCHIVED

func (ay *AcademicYear) Validate() error {
    // Business rules:
    // 1. StartDate < EndDate
    // 2. StartDate >= today + 30 days (untuk creation baru)
    // 3. Must have exactly 2 semesters
    // 4. Semesters harus fully cover date range
    // 5. TIDAK ada approval System Admin (removed)
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
    SemesterStatusDraft   SemesterStatus = "DRAFT"
    SemesterStatusActive  SemesterStatus = "ACTIVE"
    SemesterStatusInactive SemesterStatus = "INACTIVE"
)

func (s *Semester) Validate(academicYear AcademicYear) error {
    // Business rules:
    // 1. Sequence unique dalam academic year
    // 2. Dates dalam academic year range
    // 3. Tidak ada overlap dengan semester lain
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
    // 1. Code adalah unique
    // 2. Name adalah unique
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
    Weight      float64         // Untuk alignment calculation
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
    // 1. Code adalah unique
    // 2. Name adalah unique
    // 3. Weight adalah positive dan <= 1.0
    // 4. Exactly 8 dimensions bisa aktif
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
    // 1. CP tidak bisa punya duplicate alignment ke dimension yang sama
    // 2. Alignment strength adalah valid enum
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
        return "", errors.New("academic year name harus 1-100 characters")
    }
    return AcademicYearName(name), nil
}
```

### CategoryCode

```go
type CategoryCode string

func NewCategoryCode(code string) (CategoryCode, error) {
    if len(code) == 0 || len(code) > 50 {
        return "", errors.New("category code harus 1-50 characters")
    }
    if code != strings.ToUpper(code) {
        return "", errors.New("category code harus uppercase")
    }
    return CategoryCode(code), nil
}
```

### DimensionIndicators

```go
type DimensionIndicators []string

func NewDimensionIndicators(indicators []string) (DimensionIndicators, error) {
    if len(indicators) == 0 {
        return nil, errors.New("minimal satu indicator diperlukan")
    }
    return DimensionIndicators(indicators), nil
}
```

### AlignmentThreshold (NEW)

```go
// AlignmentThreshold is a Value Object that represents the configurable threshold
type AlignmentThreshold struct {
    Value float64
}

func NewAlignmentThreshold(value float64) (AlignmentThreshold, error) {
    if value < 0 || value > 100 {
        return "", errors.New("threshold harus antara 0 dan 100")
    }
    return AlignmentThreshold{Value: value}, nil
}

// GetThreshold returns the threshold value, defaulting to 60% if not configured
func (at AlignmentThreshold) GetThreshold() float64 {
    if at.Value == 0 {
        return 60.0 // Default threshold
    }
    return at.Value
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
    // Validasi 30-day lead time
    if dateRange.StartDate.Before(time.Now().AddDate(0, 0, 30)) {
        return errors.New("academic year harus dibuat minimal 30 hari di advance")
    }
    
    // Validasi tidak ada overlap dengan existing academic years
    for _, year := range existingYears {
        if year.SchoolID == schoolID {
            existingRange := DateRange{StartDate: year.StartDate, EndDate: year.EndDate}
            if dateRange.Overlaps(existingRange) {
                return errors.New("academic year dates tidak boleh overlap dengan academic year yang sudah ada")
            }
        }
    }
    
    return nil
}

func (s *AcademicYearValidationService) ValidateSemesters(
    academicYear AcademicYear,
    semesters []Semester,
) error {
    // Validasi exactly 2 semesters
    if len(semesters) != 2 {
        return errors.New("academic year harus memiliki tepat 2 semester")
    }
    
    // Validasi sequence (1 dan 2)
    sequences := make(map[int]bool)
    for _, sem := range semesters {
        if sequences[sem.Sequence] {
            return errors.New("semester sequences harus unique")
        }
        sequences[sem.Sequence] = true
    }
    if !sequences[1] || !sequences[2] {
        return errors.New("semesters harus memiliki sequences 1 dan 2")
    }
    
    // Validasi full coverage tanpa gaps
    academicYearRange := DateRange{StartDate: academicYear.StartDate, EndDate: academicYear.EndDate}
    totalCoverage := 0
    for _, sem := range semesters {
        semesterRange := DateRange{StartDate: sem.StartDate, EndDate: sem.EndDate}
        if !academicYearRange.Contains(semesterRange) {
            return errors.New("semester dates harus dalam academic year range")
        }
        totalCoverage += int(semesterRange.EndDate.Sub(semesterRange.StartDate).Hours())
    }
    
    expectedCoverage := int(academicYearRange.EndDate.Sub(academicYearRange.StartDate).Hours())
    if totalCoverage != expectedCoverage {
        return errors.New("semesters harus fully cover academic year tanpa gaps")
    }
    
    return nil
}
```

### AlignmentCalculationService (Updated)

```go
type AlignmentCalculationService struct {
    configService ConfigurationService // NEW: For configurable threshold
}

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

// CheckAgainstThreshold checks if CP alignment meets the configurable threshold
func (s *AlignmentCalculationService) CheckAgainstThreshold(alignmentPercentage float64) (bool, error) {
    // Get threshold from configuration
    threshold, err := s.configService.GetAlignmentThreshold()
    if err != nil {
        return false, err
    }
    
    thresholdVO, err := NewAlignmentThreshold(threshold)
    if err != nil {
        return false, err
    }
    
    return alignmentPercentage >= thresholdVO.GetThreshold(), nil
}
```

### ConfigurationService (NEW)

```go
// ConfigurationService reads system configuration values
type ConfigurationService struct {
    // Could read from database table, environment variables, or config file
    // For MVP, use environment variable or database table
}

func (cs *ConfigurationService) GetAlignmentThreshold() (float64, error) {
    // Try to read from configuration source
    // Priority: 1. Database table, 2. Environment variable, 3. Default value
    
    // For Sprint 4 MVP, use default if not configured
    // Future: Implement database table for configuration
    
    // Read from environment variable if exists
    if env := os.Getenv("CP_ALIGNMENT_THRESHOLD"); env != "" {
        threshold, err := strconv.ParseFloat(env, 64)
        if err == nil {
            return threshold, nil
        }
    }
    
    // Default value
    return 60.0, nil
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
    Archive(id UUID) error  // Changed from Delete to Archive
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
    Deactivate(id UUID) error  // Changed from Delete to Deactivate
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
    Deactivate(id UUID) error // Changed from Delete to Deactivate
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

# PART 5 – DATABASE DESIGN

## New Tables

### Table: academic_years

**Purpose**: Store academic year configuration dengan simplified workflow (tanpa approval)

**Columns**:
| Column | Type | Constraints | Description |
| ------ | ---- | ----------- | ----------- |
| id | UUID | PRIMARY KEY, DEFAULT gen_uuid_v7() | Unique identifier |
| school_id | UUID | NOT NULL, FK schools(id) ON DELETE RESTRICT | School ownership |
| name | VARCHAR(100) | NOT NULL | Academic year name (e.g., "2026/2027") |
| start_date | TIMESTAMP WITH TIME ZONE | NOT NULL | Academic year start date |
| end_date | TIMESTAMP WITH TIME ZONE | NOT NULL | Academic year end date |
| status | VARCHAR(20) | NOT NULL, CHECK (status IN ('DRAFT', 'ACTIVE', 'ARCHIVED')) | Workflow status (simplified: 3 states only) |
| created_by | UUID | NOT NULL, FK users(id) ON DELETE RESTRICT | Creator user reference |
| created_at | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Creation timestamp |
| updated_at | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Last update timestamp |

**Removed Fields** (compared to v1):
- ~~approved_by~~
- ~~approved_at~~

**Indexes**:
- `idx_academic_years_school_id` on (school_id)
- `idx_academic_years_status` on (status)
- `idx_academic_years_dates` on (school_id, start_date, end_date)
- `idx_academic_years_school_status` on (school_id, status) WHERE status = 'ACTIVE'

**Unique Constraints**:
- `uq_academic_years_school_name` on (school_id, name)

**Foreign Keys**:
- `fk_academic_years_school_id` → schools(id)
- `fk_academic_years_created_by` → users(id)

**Audit Fields**: created_by, created_at, updated_at

---

### Table: semesters

**Purpose**: Store semester configuration dalam academic years

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

**Purpose**: Store graduate profile dimension definitions (8 dimensions dari Profil Lulusan)

**Columns**:
| Column | Type | Constraints | Description |
| ------ | ---- | ----------- | ----------- |
| id | UUID | PRIMARY KEY, DEFAULT gen_uuid_v7() | Unique identifier |
| code | VARCHAR(50) | NOT NULL, UNIQUE | Dimension code (e.g., "KEIMANAN_KETAKWAAN") |
| name | VARCHAR(100) | NOT NULL, UNIQUE | Dimension name (e.g., "Keimanan & Ketakwaan") |
| name_en | VARCHAR(100) | | Dimension name in English |
| description | TEXT | NOT NULL | Dimension description |
| indicators | JSONB | NOT NULL | Array of dimension indicators |
| weight | DECIMAL(5,4) | NOT NULL, CHECK (weight > 0 AND weight <= 1.0) | Weight untuk alignment calculation |
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

**Purpose**: Store CP ke graduate profile dimension alignments dengan strength scoring

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

### Table: system_configuration (NEW)

**Purpose**: Store system configuration values termasuk CP alignment threshold yang configurable

**Columns**:
| Column | Type | Constraints | Description |
| ------ | ---- | ----------- | ----------- |
| id | UUID | PRIMARY KEY, DEFAULT gen_uuid_v7() | Unique identifier |
| key | VARCHAR(100) | NOT NULL, UNIQUE | Configuration key |
| value | TEXT | NOT NULL | Configuration value (JSON string for complex values) |
| description | TEXT | | Configuration description |
| created_at | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Creation timestamp |
| updated_at | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Last update timestamp |

**Indexes**:
- `idx_system_configuration_key` on (key)

**Unique Constraints**:
- `uq_system_configuration_key` on (key)

**Data Seeding**:
```sql
INSERT INTO system_configuration (id, key, value, description, created_at, updated_at) VALUES
(gen_uuid_v7(), 'cp_alignment_threshold', '60.0', 'CP alignment threshold percentage (default: 60%)', NOW(), NOW());
```

---

## Table Extensions

### Extension: curriculum_subjects

**Purpose**: Add subject categorization ke existing subjects table

**New Column**:
| Column | Type | Constraints | Description |
| ------ | ---- | ----------- | ----------- |
| subject_category_id | UUID | FK subject_categories(id) ON DELETE SET NULL | Subject category reference |

**Migration Action**: Add column, create index, update existing records

**Index**: `idx_curriculum_subjects_category_id` on (subject_category_id)`

**Foreign Key**: `fk_curriculum_subjects_category_id` → subject_categories(id)

**Note**: Column akan nullable initially, kemudian di-populate via data migration

---

### Extension: cp

**Purpose**: Add academic year scoping ke existing CP table

**New Columns**:
| Column | Type | Constraints | Description |
| ------ | ---- | ----------- | ----------- |
| academic_year_id | UUID | FK academic_years(id) ON DELETE SET NULL | Academic year reference |
| semester_id | UUID | FK semesters(id) ON DELETE SET NULL | Semester reference |

**Migration Action**: Add columns, create indexes

**Indexes**:
- `idx_cp_academic_year_id` on (academic_year_id)
- `idx_cp_semester_id` pada (semester_id)
- `idx_cp_academic_semester` pada (academic_year_id, semester_id)

**Foreign Keys**:
- `fk_cp_academic_year_id` → academic_years(id)
- `fk_cp_semester_id` → semesters(id)

**Note**: Columns akan nullable untuk support CP yang sudah ada tanpa academic year scoping

---

## Migration Strategy

### Migration File: 000010_sprint4_academic_foundation.up.sql

```sql
-- Sprint 4 Academic Foundation Migration (Refactored v2)
-- Purpose: Add academic year, semester, subject category, graduate profile dimension infrastructure
-- Risk Level: MEDIUM (new tables, existing table extensions)
-- Tables: 5 new, 2 extended, 1 configuration
-- Changes from v1: Simplified academic year workflow (removed approval fields), added system configuration table

-- ============================================================================
-- NEW TABLES
-- ============================================================================

-- Table: system_configuration (NEW)
-- Table: academic_years
-- Table: semesters
-- Table: subject_categories
-- Table: graduate_profile_dimensions
-- Table: cp_alignments
-- Table extensions: curriculum_subjects, cp

-- ============================================================================
-- Table: system_configuration (NEW)
-- ============================================================================

CREATE TABLE system_configuration (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    key VARCHAR(100) NOT NULL UNIQUE,
    value TEXT NOT NULL,
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_system_configuration_key ON system_configuration(key);

-- Seed default configuration
INSERT INTO system_configuration (id, key, value, description, created_at, updated_at) VALUES
(gen_uuid_v7(), 'cp_alignment_threshold', '60.0', 'CP alignment threshold percentage (default: 60%)', NOW(), NOW());

-- ============================================================================
-- Table: academic_years (NEW - Simplified)
-- ============================================================================

CREATE TABLE academic_years (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    school_id UUID NOT NULL,
    name VARCHAR(100) NOT NULL,
    start_date TIMESTAMP WITH TIME ZONE NOT NULL,
    end_date TIMESTAMP WITH TIME ZONE NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'DRAFT' CHECK (status IN ('DRAFT', 'ACTIVE', 'ARCHIVED')),
    created_by UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_academic_years_school_id FOREIGN KEY (school_id) REFERENCES schools(id) ON DELETE RESTRICT,
    CONSTRAINT fk_academic_years_created_by FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE RESTRICT,
    CONSTRAINT uq_academic_years_school_name UNIQUE (school_id, name),
    CONSTRAINT chk_academic_years_dates CHECK (start_date < end_date)
);

CREATE INDEX idx_academic_years_school_id ON academic_years(school_id);
CREATE INDEX idx_academic_years_status ON academic_years(status);
CREATE INDEX idx_academic_years_dates ON academic_years(school_id, start_date, end_date);
CREATE INDEX idx_academic_years_school_status ON academic_years(school_id, status) WHERE status = 'ACTIVE';

-- ============================================================================
-- Table: semesters (NEW)
-- ============================================================================

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

-- ============================================================================
-- Table: subject_categories (NEW)
-- ============================================================================

CREATE TABLE subject_categories (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    code VARCHAR(50) NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL UNIQUE,
    name_en VARCHAR(202) NOT NULL UNIQUE,
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

-- ============================================================================
-- Table: graduate_profile_dimensions (NEW)
-- ============================================================================

CREATE TABLE graduate_profile_dimensions (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    code VARCHAR(50) NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL UNIQUE,
    name_en VARCHAR(202) NOT NULL UNIQUE,
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

-- ============================================================================
-- Table: cp_alignments (NEW)
-- ============================================================================

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
('550e8400-e29b-41d4-a716-446655440021', 'KEWARGAAN', 'Kewargaan', 'Citizenship', 'Dimensi kesadaran berbangsa dan bernegara', '["Cinta tanah air", "Menghargai keberagaman", "Taats aturan"]'::jsonb, 0.125, 'ACTIVE', (SELECT id FROM WHERE role = 'SYSTEM_ADMIN' LIMIT 1), NOW(), NOW()),
('550e8400-e29b-41d4-a716-446655440022', 'BERAKHLAK_MULIA', 'Berakhlak Mulia', 'Noble Character', 'Dimensi pembentukan karakter mulia', '["Jujur", "Disiplin", "Tanggung jawab"]'::jsonb, 0.125, 'ACTIVE', (SELECT id FROM users WHERE role = 'SYSTEM_ADMIN' LIMIT 1), NOW(), NOW()),
('550e8400-e29b-41d4-a716-446655440023', 'BERANI_BERTANGGUNG_JAWAB', 'Berani Bertanggung Jawab', 'Courageous and Responsible', 'Dimensi keberanian dan tanggung jawab', '["Berani mengambil keputusan", "Pertanggung jawaban atas tindakan"]'::jsonb, 0.125, 'ACTIVE', (SELECT id FROM users WHERE role = 'SYSTEM_ADMIN' LIMIT 1), NOW(), NOW()),
('550e8400-e29b-41d4-a716-446655440024', 'PEDULI', 'Peduli', 'Caring', 'Dimensi kepedulian terhadap sesama', '["Empati", "Saling membantu", "Toleransi"]'::jsonb, 0.125, 'ACTIVE', (SELECT id FROM users WHERE role = 'SYSTEM_ADMIN' LIMIT 1), NOW(), NOW()),
('550e8400-e29b-41d4-a716-446655440025', 'GOTONG_ROYONG', 'Gotong Royong', 'Collaboration', 'Dimensi kerja sama dan gotong royong', '["Kerja tim", "Solidaritas", "Kolaborasi"]'::jsonb, 0.125, 'ACTIVE', (SELECT id FROM users WHERE role = 'SYSTEM_ADMIN' LIMIT 1), NOW(), NOW()),
('550e8400-e29b-41d4-a716-446655440026', 'MANDIRI', 'Mandiri', 'Independent', 'Dimensi kemandirian dan otonomi', '["Berpikir kritis", "Mengambil inisiatif", "Mandiri belajar"]'::jsonb, 0.125, 'ACTIVE', (SELECT id FROM users WHERE role = 'SYSTEM_ADMIN' LIMIT 1), NOW(), NOW()),
('550e8400-e29b-41d4-a716-446655440027', 'KREATIF', 'Kreatif', 'Creative', 'Dimensi kreativitas dan inovasi', '["Berpikir kreatif", "Inovasi", "Pemecahan masalah"]'::jsonb, 0.125, 'ACTIVE', (SELECT id FROM users WHERE role = 'SYSTEM_ADMIN' LIMIT 1), NOW(), NOW());
```

### Migration File: 000010_sprint4_academic_foundation.down.sql

```sql
-- Sprint 4 Academic Foundation Rollback (Refactored v2)
-- Risk Level: MEDIUM (data loss jika tidak di-backup)

-- Drop indexes untuk table extensions
DROP INDEX IF EXISTS idx_cp_academic_semester;
DROP INDEX IF EXISTS idx_cp_semester_id;
DROP INDEX IF EXISTS idx_cp_academic_year_id;

-- Drop foreign keys untuk table extensions
ALTER TABLE cp DROP CONSTRAINT IF EXISTS fk_cp_semester_id;
ALTER TABLE cp DROP CONSTRAINT IF EXISTS fk_cp_academic_year_id;

-- Drop columns untuk table extensions
ALTER TABLE cp DROP COLUMN IF EXISTS semester_id;
ALTER TABLE cp DROP COLUMN IF EXISTS academic_year_id;

DROP INDEX IF EXISTS idx_curriculum_subjects_category_id;
ALTER TABLE curriculum_subjects DROP CONSTRAINT IF EXISTS fk_curriculum_subjects_category_id;
ALTER TABLE curriculum_subjects DROP COLUMN IF EXISTS subject_category_id;

-- Drop indexes untuk new tables
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
DROP INDEX IF EXISTS idxacademic_years_status;
DROP INDEX IF EXISTS idx_academic_years_school_id;

DROP INDEX IF EXISTS idx_system_configuration_key;

-- Drop new tables
DROP TABLE IF EXISTS cp_alignments;
DROP TABLE IF EXISTS graduate_profile_dimensions;
DROP TABLE IF EXISTS subject_categories;
DROP TABLE IF EXISTS semesters;
DROP TABLE IF EXISTS academic_years;
DROP TABLE IF EXISTS system_configuration;
```

---

# PART 6 – API DESIGN

## Academic Year API (Simplified)

### Endpoints

#### 1. Create Academic Year

**Method**: `POST`  
**URL**: `/api/v1/academic/academic-years`  
**Authorization**: `SCHOOL_ADMIN`  
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
  "created_by": "uuid",
  "created_at": "2026-06-11T10:00:00Z",
  "updated_at": "2026-06-11T10:00:00Z"
}
```

**Removed from v1**:
- ~~approved_by~~
- ~~approved_at~~

**Validation**:
- `school_id`: Required, must exist dan belong ke user's school
- `name`: Required, max 100 characters, unique per school
- `start_date`: Required, valid date, >= today + 30 days
- `end_date`: Required, valid date, > start_date
- `description`: Optional, max 500 characters

**Error Codes**:
- `400 BAD_REQUEST`: Validation error
- `403 FORBIDDEN`: Insufficient permissions
- `409 CONFLICT`: Academic year name sudah ada untuk school
- `409 CONFLICT`: Date range overlap dengan academic year yang ada
- `422 UNPROCESSABLE_ENTITY`: Start date kurang dari 30 hari di future

---

#### 2. List Academic Years

**Method**: `GET`  
**URL**: `/api/v1/academic/academic-years`  
**Authorization**: `SCHOOL_ADMIN`, `SYSTEM_ADMIN`, `CURRICULUM_ADMIN`, `TEACHER`  
**Permission**: `academic_year:READ`

**Query Parameters**:
- `school_id` (optional): Filter by school
- `status` (optional): Filter by status (DRAFT, ACTIVE, ARCHIVED)

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
      "created_at": "2026-06-01T10:00:00Z",
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

**Removed from v1**:
- ~~approved_by~~
- ~~approved_by_name~~
- ~~approved_at~~

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
  "start_date": "2026-07-15T00:00:00:00Z",
  "end_date": "2027-06-30T23:59:59Z",
  "status": "ACTIVE",
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

**Removed from v1**:
- ~~approved_by~~
- ~~approved_by_name~~
- ~~approved_at~~

**Error Codes**:
- `403 FORBIDDEN`: Insufficient permissions atau school access
- `404 NOT_FOUND`: Academic year not found

---

#### 4. Update Academic Year

**Method**: `PUT`  
**URL**: `/api/v1/academic/academic-years/:id`  
**Authorization**: `SCHOOL_ADMIN`  
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
- Hanya bisa update academic years dalam status DRAFT
- Date range validation sama seperti create
- Name uniqueness validation sama seperti create

**Error Codes**:
- `400 BAD_REQUEST`: Validation error
- `403 FORBIDDEN`: Insufficient permissions
- `404 NOT_FOUND`: Academic year not found
- `409 CONFLICT`: Cannot update active or archived academic year

---

#### 5. Activate Academic Year (NEW - Simplified)

**Method**: `POST`  
**URL**: `/api/v1/academic/academic-years/:id/activate`  
**Authorization**: `SCHOOL_ADMIN`  
**Permission**: `academic_year:ACTIVATE`

**Request**:
```json
{
  "reason": "Siap untuk tahun ajaran 2026/2027"
}
```

**Response**: `200 OK`
```json
{
  "id": "uuid",
  "status": "ACTIVE",
  "updated_at": "2026-06-11T11:00:00Z"
}
```

**Validation**:
- Academic year harus memiliki status DRAFT
- Academic year harus memiliki 2 semester dikonfigurasi
- Semua required fields populated
- Tidak ada academic year lain yang sedang ACTIVE untuk school

**Error Codes**:
- `400 BAD_REQUEST`: Academic year tidak siap untuk activation
- `403 FORBIDDEN`: Insufficient permissions
- `404 NOT_FOUND`: Academic year not found
- `409 CONFLICT`: Tidak bisa activate (ada academic year lain yang aktif)

**Removed from v1**:
- ~~Submit for approval endpoint~~
- ~~Approve endpoint~~

---

#### 6. Archive Academic Year (NEW - Replaces Deactivate)

**Method**: `POST`  
**URL**: `/api/v1/academic/academic-years/:id/archive`  
**Authorization**: `SCHOOL_ADMIN`  
**Permission**: `academic_year:ARCHIVE`

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
  "status": "ARCHIVED",
  "updated_at": "2026-06-11T11:00:00Z"
}
```

**Validation**:
- Academic year harus dalam status ACTIVE
- Must provide reason
- Tidak ada active curriculum planning (CP/TP/ATP yang sedang aktif)

**Error Codes**:
- `400 BAD_REQUEST`: Cannot archive (has active dependencies)
- `403 FORBIDDEN`: Insufficient permissions
- `404 NOT_FOUND`: Academic year not found

**Removed from v1**:
- ~~Deactivate endpoint~~

---

## Semester API

### Endpoints

#### 1. Create Semester

**Method**: `POST`  
**URL**: `/api/v1/academic/semesters`  
**Authorization**: `SCHOOL_ADMIN`  
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
  "start_date": "20206-07-15T00:00:00Z",
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
- `409 CONFLICT`: Sequence conflict atau date overlap

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
**Authorization**: `SCHOOL_ADMIN`  
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
  "updated_at": "2026-06-11T11:00:00:00Z"
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
**Authorization**: `SCHOOL_ADMIN`  
**Permission**: `semester:DELETE`

**Response**: `204 No Content`

**Validation**:
- Hanya bisa delete semesters dalam status DRAFT
- Tidak bisa delete jika linked ke CP atau entities lain

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
- `409 CONFLICT`: Code atau name sudah ada

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
- Tidak bisa deactivate jika subjects assigned

**Error Codes**:
- `400 BAD_REQUEST`: Tidak bisa deactivate (has assigned subjects)
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
- `409 CONFLICT`: Code atau name sudah ada
- `422 UNPROCESSABLE_ENTITY`: Lebih dari 8 active dimensions

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
      "name_en": "FAITH AND PIETY",
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
**Permission`: `graduate_profile_dimension:DEACTIVATE`

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
- Tidak bisa deactivate jika CP align ke dimension
- Tidak bisa deactivate jika akan menghasilkan kurang dari 8 active dimensions

**Error Codes**:
- `400 BAD_REQUEST`: Tidak bisa deactivate (has aligned CP atau < 8 dimensions)
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
  "rationale": "Strong alignment ke faith dimension"
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
  "rationale": "Strong alignment ke faith dimension",
  "created_by": "uuid",
  "created_by_name": "Admin NUSA",
  "created_at": "2026-06-11T10:00:00Z",
  "updated_at": "2026-06-11T10:00:00:00Z"
}
```

**Error Codes**:
- `400 BAD_REQUEST`: Validation error
- `403 FORBIDDEN`: Insufficient permissions
- `404 NOT_FOUND`: CP atau dimension not found
- `409 CONFLICT`: Alignment already exists untuk CP-dimension pair

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
      "rationale": "Strong alignment ke faith dimension",
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
**Authorization**: `CP_ALIGNMENT:CURRICULUM_ADMIN`  
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
**Authorization**: `CP_ALIGNMENT:CURRICULUM_ADMIN`  
**Permission**: `cp_alignment:DELETE`

**Response**: `204 No Content`

**Validation**:
- Warn jika CP akan fall di **konfigurasi threshold** (dibaca dari system configuration)
- Error message menggunakan nilai threshold aktual dari konfigurasi

**Error Codes**:
- `400 BAD_REQUEST`: Tidak bisa delete (akan melanggarui **konfigurasi threshold**)
- `403 FORBIDDEN`: Insufficient permissions
- `404 NOT_FOUND`: Alignment not found

---

#### 5. Get Alignment Report (Updated)

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
    "below_threshold_count": 3,
    "alignment_threshold": 60.0
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
      "alignments_count": 1,
      "alignment_threshold": 60.0
    }
  ]
}
```

**Added in v2**: `alignment_threshold` field showing the configured threshold

**Error Codes**:
- `403 FORBIDDEN`: Insufficient permissions

---

## System Configuration API (NEW)

### Endpoints

#### 1. Get Configuration

**Method**: `GET`  
**URL**: `/api/v1/system/configuration`  
**Authorization**: `SYSTEM_ADMIN`  
**Permission**: `system_config:READ`

**Query Parameters**:
- `key` (optional): Filter by configuration key

**Response**: `200 OK`
```json
{
  "data": [
    {
      "id": "uuid",
      "key": "cp_alignment_threshold",
      "value": "60.0",
      "description": "CP alignment threshold percentage (default: 60%)"
    }
  ]
}
```

**Error Codes**:
- `403 FORBIDDEN`: Insufficient permissions (SYSTEM_ADMIN only)

---

#### 2. Update Configuration

**Method**: `PUT`  
**URL**: `/api/v1/system/configuration/:key`  
**Authorization**: `SYSTEM_ADMIN`  
**Permission**: `system_config:UPDATE`

**Request**:
```json
{
  "value": "70.0",
  "description": "Increased threshold to 70% for pilot phase"
}
```

**Response**: `200 OK`
```json
{
  "id": "uuid",
  "key": "cp_alignment_threshold",
  "value": "70.0",
  "description": "Increased threshold to 70% for pilot phase",
  "updated_at": "2026-06-11T11:00:00:00Z"
}
```

**Validation**:
- Key must exist
- Value must be valid number antara 0 dan 100
- Only SYSTEM_ADMIN can update configuration

**Error Codes**:
- `400 BAD_REQUEST`: Invalid key or value
- `403 FORBIDDEN`: Insufficient permissions (SYSTEM_ADMIN only)
- `404 NOT_FOUND`: Configuration key not found

---

# PART 7 – FRONTEND REQUIREMENTS

## Pages Overview

### School Admin Pages

#### 1. Academic Year Management (Simplified)

**Path**: `/admin/academic/academic-years`  
**Purpose**: Manage academic years untuk school  
**Actors**: School Admin

**Components**:
- **Table**: List academic years dengan columns (Name, Start Date, End Date, Status, Actions)
- **Filters**: Status (All, Draft, Active, Archived)
- **Search**: Search by academic year name
- **Actions**: Create, View, Edit (jika Draft), Activate, Archive (jika tidak aktif)
- **Pagination**: 20 items per page

**Form Fields (Create/Edit)**:
- Academic Year Name (text, required, max 100 chars)
- Start Date (date picker, required, >= today + 30 days)
- End Date (date picker, required, > start date)
- Description (textarea, optional, max 500 chars)

**Validation**:
- Real-time validation untuk date constraints
- Warning jika dates overlap dengan existing academic years
- Disable submit jika validation fails

**Removed in v2**:
- ~~Submit for approval button~~
- ~~Approve button~~
- ~~Approval status~~
- ~~Approval workflow UI~~

**Permissions**: `academic_year:READ`, `academic_year:CREATE`, `academic_year:UPDATE`, `academic_year:ACTIVATE`, `academic_year:ARCHIVE`

---

#### 2. Semester Configuration

**Path**: `/admin/academic/academic-years/:id/semesters`  
**Purpose**: Configure semesters dalam academic year  
**Actors**: School Admin

**Components**:
- **Table**: List semesters dengan columns (Name, Sequence, Start Date, End Date, Status, Actions)
- **Visualization**: Timeline showing academic year dengan semesters
- **Actions**: Create, Edit (jika Draft), Archive (jika tidak aktif)
- **Status Badges**: DRAFT (gray), ACTIVE (green), INACTIVE (red)

**Form Fields (Create/Edit)**:
- Semester Name (text, required, max 50 chars, default "Semester Ganjil"/"Semester Genap")
- Sequence (dropdown, required, options: 1, 2)
- Start Date (date picker, required, dalam academic year range)
- End Date (date picker, required, dalam academic year range)
- Description (textarea, optional, max 500 chars)

**Validation**:
- Sequence harus unique dalam academic year
- Dates harus dalam academic year range
- Tidak ada gaps atau overlaps antara semesters
- Visual gap/overlap indicators di timeline

**Permissions**: `semester:READ`, `semester:CREATE`, `semester:UPDATE`, `semester:DELETE`, `semester:ARCHIVE`

---

### Curriculum Admin Pages

#### 3. Subject Category Management

**Path**: `/admin/curriculum/subject-categories`  
**Purpose**: Manage subject categories (Intrakurikuler, Kokurikuler, Ekstrakurikuler)  
**Actors**: Curriculum Admin

**Components**:
- **Table**: List categories dengan columns (Code, Name, Name EN, Status, Subjects Count, Actions)
- **Filters**: Status (All, Active, Inactive)
- **Search**: Search by code atau name
- **Actions**: Create, View, Edit, Deactivate
- **Stats Cards**: Total categories, Active categories, Total subjects categorized

**Form Fields (Create/Edit)**:
- Category Code (text, required, uppercase, max 50 chars)
- Category Name (text, required, max 100 chars)
- Category Name (English) (text, optional, max 100 chars)
- Description (textarea, required, max 1000 chars)
- Guidelines (textarea, optional, max 2000 chars)

**Validation**:
- Code harus uppercase dan unique
- Name harus unique
- Tidak bisa deactivate jika subjects assigned

**Permissions**: `subject_category:READ`, `subject_category:CREATE`, `subject_category:UPDATE`, `subject_category:DEACTIVATE`

---

#### 4. Subject Categorization

**Path**: `/admin/curriculum/subjects/categorization`  
**Purpose**: Categorize subjects ke categories  
**Actors**: Curriculum Admin

**Components**:
- **Table**: List subjects dengan columns (Code, Name, Current Category, Actions)
- **Filters**: Category (All, Intrakurikuler, Kokurikuler, Ekstrakurikuler, Uncategorized)
- **Bulk Actions**: Assign category ke selected subjects
- **Progress Bar**: Show categorization completion percentage

**Form Fields (Bulk Assign)**:
- Category (dropdown, required)
- Selected subjects (checkbox list)

**Validation**:
- Minimal satu subject harus dipilih
- Category harus aktif

**Permissions**: `curriculum_subjects:UPDATE`

---

#### 5. Graduate Profile Dimensions

**Path**: `/admin/curriculum/graduate-profile-dimensions`  
**Purpose**: Manage 8-dimensional Profil Lulusan  
**Actors**: Curriculum Admin

**Components**:
- **Table**: List dimensions dengan columns (Code, Name, Name EN, Weight, Status, CP Alignments Count, Actions)
- **Filters**: Status (All, Active, Inactive)
- **Stats Cards**: Total dimensions (8), Active dimensions, Total CP aligned
- **Visualization**: Radar chart showing dimension coverage

**Form Fields (Create/Edit)**:
- Dimension Code (text, required, uppercase, max 50 chars)
- Dimension Name (text, required, max 100 chars)
- Dimension Name (English) (text, optional, max 100 chars)
- Description (textarea, required, max 1000 chars)
- Indicators (dynamic tags input, required, minimal 1)
- Weight (number, required, 0 < weight <= 1.0, default 0.125)

**Validation**:
- Code harus uppercase dan unique
- Name harus unique
- Weight harus positive dan <= 1.0
- Exactly 8 dimensions bisa aktif
- Tidak bisa deactivate jika CP align ke dimension

**Permissions**: `graduate_profile_dimension:READ`, `graduate_profile_dimension:CREATE`, `graduate_profile_dimension:UPDATE`, `graduate_profile_dimension:DEACTIVATE`

---

#### 6. CP Alignment Management

**Path**: `/admin/curriculum/cp-alignment`  
**Purpose: Align CP ke graduate profile dimensions  
**Actors**: Curriculum Admin

**Components**:
- **Table**: List CP dengan columns (Code, Description, Alignment Percentage, Dimensions Aligned, Actions)
- **Filters**: Subject, Phase, Alignment Level (All, <threshold, 60-80%, >80%) - threshold configurable
- **Search**: Search by CP code atau description
- **Alignment Visual**: Progress bar showing alignment percentage
- **Bulk Actions**: Align selected CP ke dimensions

**Form Fields (Align CP)**:
- Select CP (checkbox list atau dropdown)
- Select Dimension(s) (multi-select dropdown)
- Alignment Strength (radio buttons: Strong, Medium, Weak)
- Rationale (textarea, optional, max 500 chars)

**Validation**:
- Minimal satu CP harus dipilih
- Minimal satu dimension harus dipilih
- Alignment strength harus dipilih
- Warn jika CP akan fall di **konfigurasi threshold** (default 60 tapi configurable via system config)

**Permissions**: `cp_alignment:READ`, `cp_alignment:CREATE`, `cp_alignment:UPDATE`, `cp_alignment:DELETE`

---

#### 7. Alignment Report (Updated)

**Path**: `/admin/curriculum/alignment-report`  
**Purpose**: View curriculum alignment reports  
**Actors**: Curriculum Admin, School Admin

**Components**:
- **Summary Cards**: Total CP, Aligned CP, Overall Alignment %, Below Threshold Count, **Alignment Threshold**
- **Filters**: Academic Year, Phase, Subject
- **Charts**: 
  - Bar chart: Alignment per dimension
  - Pie chart: CP alignment distribution
  - Line chart: Alignment trend over time
- **Table**: CP **dibawah threshold** dengan details - threshold configurable
- **Actions**: Export PDF, Export CSV

**Validation**: None (read-only page)

**Permissions**: `alignment_report:READ`

---

#### 8. Koding & AI Subject Setup

**Path**: `/admin/curriculum/modern-subjects`  
**Purpose**: Add Koding, AI, Numerasi subjects  
**Actors**: Curriculum Admin

**Components**:
- **Table**: List modern subjects (Koding, AI, Numerasi) dengan status
- **Form**: Subject creation form untuk new modern subjects
- **Instructions**: Guide untuk menambah modern subjects
- **Integration Check**: Verify subjects integrate dengan CP/TP workflow

**Form Fields**:
- Subject Code (text, required, uppercase, max 50 chars)
- Subject Name (text, required, max 255 chars)
- Subject Name (English) (text, optional, max 255 chars)
- Description (textarea, required, max 1000 chars)
- Category (dropdown, required, default Intrakurikuler)
- Phase (multi-select, required)
- Is Active (toggle, required)

**Validation**:
- Code harus uppercase dan unique
- Name harus unique
- Category harus aktif
- Minimal satu phase harus dipilih

**Permissions**: `curriculum_subjects:CREATE`

---

### Teacher Pages

#### 9. Academic Calendar View

**Path**: `/teacher/academic-calendar`  
**Purpose**: View academic calendar untuk planning  
**Actors**: Teacher

**Components**:
- **Calendar View**: Monthly calendar showing academic year, semester dates
- **Academic Year Info**: Current academic year, semester dates
- **Upcoming Events**: Important dates (semester start/end)
- **Read-only**: Tidak ada editing capabilities

**Validation**: None (read-only page)

**Permissions**: `academic_year:READ`

---

#### 10. Subject Categories View

**Path**: `/teacher/curriculum/subject-categories`  
**Purpose**: View subject categories untuk understanding  
**Actors**: Teacher

**Components**:
- **Table**: List categories dengan descriptions dan guidelines
- **Subject List**: Show subjects di bawah masing-masing category
- **Read-only**: Tidak ada editing capabilities

**Validation**: None (read-only page)

**Permissions**: `subject_category:READ`

---

#### 11. Graduate Profile Reference

**Path**: `/teacher/curriculum/graduate-profile`  
**Purpose**: Reference graduate profile dimensions untuk lesson planning  
**Acters**: Teacher

**Components**:
- **Cards**: 8 dimension cards dengan descriptions dan indicators
- **Search**: Search dimensions by name atau indicator
- **CP Alignment**: Show CP mana yang align ke dimension mana
- **Read-only**: Tidak ada editing capabilities

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

1. **Data Tables**: Semua table menggunakan pagination standar (20 items), filters, search, action menus
2. **Form Validation**: Real-time validation dengan inline error messages
3. **Loading States**: Skeleton loaders untuk semua async operations
4. Empty States: Friendly empty state messages dengan call-to-action
5. Confirmation Dialogs: Semua destructive action memerlukan confirmation
6. Success Notifications: Toast notifications untuk operasi yang berhasil
7. Error Handling: User-friendly error messages dengan retry options

### Indonesian Language

Semua teks UI harus dalam Bahasa Indonesia:
- "Academic Year" → "Tahun Ajaran"
- "Semester" → "Semester"
- "Subject Category" → "Kategori Mata Pelajaran"
- "Graduate Profile Dimension" → "Dimensi Profil Lulusan"
- "Alignment" → "Kesesuaian" atau "Aligment"
- "Activate" → "Aktifkan"
- "Archive" → "Arsipkan"

### Accessibility

- WCAG 2.1 AA compliance
- Keyboard navigation support
- Screen reader support
- Color contrast ratio >= 4.5:1
- Focus indicators pada semua interactive elements

---

# PART 8 – SECURITY REQUIREMENTS

## Roles

### Existing Roles (Reuse)

- **SYSTEM_ADMIN**: Platform-level administration
- **SCHOOL_ADMIN**: School-level administration
- **TEACHER**: Curriculum planning dan delivery
- **CURRICULUM_ADMIN**: (Existing but should be verified) System-wide curriculum governance

### Existing Role Verification

Let me check if CURRICULUM_ADMIN role exists in the codebase.<tool_call>grep<arg_key>case_insensitive</arg_key><arg_value>true
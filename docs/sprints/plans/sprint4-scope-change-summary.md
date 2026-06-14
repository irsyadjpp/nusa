# Sprint 4 Scope Change Summary

**Document Version**: 1.0  
**Date**: 2026-06-11  
**Status**: Final  
**Purpose**: Document scope changes made to Sprint 4 to ensure focus on Academic Foundation only

---

## Executive Summary

Sprint 4 telah di-refactor untuk fokus secara eksklusif pada Academic Foundation. Perubahan ini bertujuan untuk:
1. Menghilangkan kompleksitas yang tidak perlu
2. Mempercepat timeline development (dari 5 minggu ke 4 minggu)
3. Menghilangkan dependensi ke modul yang belum ada (Student Management, PPDB, Class Management)
4. Menyederhanakan workflow untuk MVP

**Total Perubahan**: 6 area utama di-refactor

---

## Scope Changes Summary

### Change 1: Removed Approval Workflow for Academic Year

**Original Scope (v1)**:
- Academic Year workflow: Draft → Under Review → Approved → Active → Inactive
- System Admin approval required
- Fields: `approved_by`, `approved_at`
- Approval permissions: `academic_year:SUBMIT`, `academic_year:APPROVE`

**New Scope (v2)**:
- Academic Year workflow: Draft → Active → Archived
- School Admin dapat activate langsung tanpa approval
- Fields removed: `approved_by`, `approved_at`
- Removed permissions: `academic_year:SUBMIT`, `academic_year:APPROVE`
- New permission: `academic_year:ACTIVATE`

**Rationale**:
- Approval workflow menambah kompleksitas tanpa value signifikan untuk MVP
- School Admin memiliki authority penuh untuk konfigurasi academic year
- Simplifikasi mengurangi development time dan maintenance burden
- Alignment dengan prinsip "self-service governance"

**Impact**:
- Database: Removed 2 columns dari `academic_years` table
- API: Removed 3 endpoints (submit, approve, reject)
- Frontend: Removed approval UI components
- Backend: Removed approval service logic
- Timeline: Mengurangi development time ~2 hari

---

### Change 2: CP Alignment Threshold Made Configurable

**Original Scope (v1)**:
- CP alignment threshold hardcoded ke 60%
- Threshold tidak bisa diubah tanpa redeploy
- Error message: "CP alignment will fall below 60% threshold"
- Validation logic: `if alignmentPercentage < 60.0 { return error }`

**New Scope (v2)**:
- CP alignment threshold configurable via system configuration
- Default value: 60% (jika konfigurasi belum tersedia)
- New table: `system_configuration` untuk menyimpan konfigurasi
- New API endpoint: `GET/PUT /api/v1/system/configuration`
- Error message: "CP alignment will fall below [threshold]% threshold" (dinamis)
- Validation logic: `if alignmentPercentage < config.GetThreshold() { return error }`

**Rationale**:
- Different schools mungkin memiliki standar alignment yang berbeda
- Threshold bisa di-adjust berdasarkan pilot phase feedback
- Tidak perlu redeploy untuk mengubah threshold
- Menambah fleksibilitas tanpa menambah kompleksitas signifikan

**Impact**:
- Database: Added 1 new table (`system_configuration`)
- API: Added 2 new endpoints (get config, update config)
- Backend: Added `ConfigurationService` domain service
- Frontend: Added system configuration management page (System Admin only)
- Timeline: Menambah development time ~0.5 hari

---

### Change 3: Removed Dependencies to PPDB/Student Management/Class Management

**Original Scope (v1)**:
- Beberapa business rules mengasumsikan keberadaan Student Management
- Referensi ke "school-level academic configuration" yang mengasumsikan student data
- Alignment reports mengasumsikan class-level reporting

**New Scope (v2)**:
- Semua use cases Sprint 4 dapat berjalan tanpa student data
- Academic Year dan Semester adalah school-level configuration independent
- Alignment reports adalah curriculum-level reporting (bukan class-level)
- Tidak ada foreign key ke student, class, atau PPDB tables

**Rationale**:
- Sprint 4 adalah Academic Foundation, bukan Student Lifecycle
- Student Management akan ditangani di sprint terpisah
- Menghindari circular dependencies antar sprint
- Setiap sprint harus dapat berjalan secara independen

**Impact**:
- Database: Tidak ada foreign key ke student/class/PPDB tables
- API: Tidak ada endpoints yang memerlukan student context
- Business Logic: Semua rules independent dari student data
- Timeline: Tidak ada impact (was already independent)

---

### Change 4: Simplified Academic Year States

**Original Scope (v1)**:
- Status enum: `DRAFT`, `UNDER_REVIEW`, `APPROVED`, `ACTIVE`, `INACTIVE`
- 5 states dengan complex transitions

**New Scope (v2)**:
- Status enum: `DRAFT`, `ACTIVE`, `ARCHIVED`
- 3 states dengan simple transitions

**Rationale**:
- Mengurangi kompleksitas state machine
- Memudahkan testing dan debugging
- Align dengan MVP requirements
- Archive lebih jelas daripada Inactive (implies historical retention)

**Impact**:
- Database: Simplified CHECK constraint
- Backend: Simplified validation logic
- Frontend: Simplified status badges dan transitions
- Testing: Mengurangi test cases dari 10 ke 5
- Timeline: Mengurangi development time ~0.5 hari

---

### Change 5: Removed Principal as Stakeholder

**Original Scope (v1)**:
- Principal listed sebagai Primary Stakeholder dengan responsibility "Oversight and approval"

**New Scope (v2)**:
- Principal removed dari stakeholder list
- School Admin memiliki authority penuh untuk academic configuration

**Rationale**:
- Principal approval handled di organisational level, bukan di software level
- School Admin (admin sekolah) memiliki authority teknis untuk konfigurasi
- Simplifikasi stakeholder management

**Impact**:
- BRD: Updated stakeholder section
- UI: Tidak ada approval workflow untuk Principal
- Timeline: Tidak ada impact (documentation only)

---

### Change 6: Updated Success Metrics

**Original Scope (v1)**:
- "Academic year coverage" dengan target 100%
- "Configuration time" dengan target <5 menit

**New Scope (v2)**:
- Metrics tetap sama tetapi dengan clarifications
- Configuration time target diupdate untuk reflect simplified workflow
- Added metric: "Self-service activation rate" (school admin yang dapat activate tanpa support)

**Rationale**:
- Metrics harus align dengan workflow yang baru
- Self-service capability adalah key success factor untuk MVP

**Impact**:
- BRD: Updated success metrics section
- Monitoring: Updated metrics untuk dashboard
- Timeline: Tidak ada impact (documentation only)

---

## Items Moved to Backlog/Future Sprint

### Moved to Sprint 5 (Student Lifecycle)

**Original Sprint 4 Scope** (removed):
- Student Management (lifecycle, enrollment, records)
- Student biodata tracking
- Student attendance tracking
- Student health tracking

**Rationale**:
- Di luar scope Academic Foundation
- Memerlukan student data yang belum ada
- Sprint 5 akan fokus pada Student Lifecycle

---

### Moved to Sprint 6 (Class Management)

**Original Sprint 4 Scope** (removed):
- Class Management (Rombel, Wali Kelas)
- Scheduling (class timetable, teacher assignments)
- Teacher workload tracking (beban mengajar)

**Rationale**:
- Di luar scope Academic Foundation
- Memerlukan student data dan teacher assignment
- Sprint 6 akan fokus pada Class Management

---

### Moved to Sprint 7 (PPDB/Enrollment)

**Original Sprint 4 Scope** (removed):
- PPDB (Penerimaan Peserta Didik Baru)
- New student admission workflow
- Document verification (PPDB dokumen)
- Zonasi mapping
- Dukcapil integration

**Rationale**:
- Di luar scope Academic Foundation
- Memerlukan integration eksternal yang kompleks
- Sprint 7 akan fokus pada PPDB

---

### Moved to Sprint 8 (Dapodik Integration)

**Original Sprint 4 Scope** (removed):
- Dapodik Integration (external system synchronization)
- Data synchronization dengan Ministry of Education

**Rationale**:
- Di luar scope Academic Foundation
- Memerlukan integration infrastructure yang belum ada
- Sprint 8 akan fokus pada Dapodik

---

### Moved to Sprint 9 (Deep Learning Pedagogy)

**Original Sprint 4 Scope** (removed):
- Deep Learning Pedagogy implementation
- Teaching methodology tracking
- Pedagogical assessment

**Rationale**:
- Di luar scope Academic Foundation
- Memerlukan pedagogical framework yang belum didefinisikan
- Sprint 9 akan fokus pada pedagogy

---

### Moved to Sprint 10 (Assessment Alignment)

**Original Sprint 4 Scope** (removed):
- Assessment Alignment (linking assessments ke graduate profiles)
- Rubric alignment ke dimensions
- Assessment reporting dengan graduate profile context

**Rationale**:
- Di luar scope Academic Foundation
- Assessment alignment lebih cocok setelah Assessment module mature
- Sprint 10 akan fokus pada Assessment Integration

---

## Timeline Impact

### Original Timeline (v1)
- Backend: 13.5 days
- Frontend: 12 days
- Database: 1 day
- QA: 3.5 days
- **Total**: 30 days (~5 weeks)

### New Timeline (v2)
- Backend: 11 days (removed approval logic)
- Frontend: 10 days (removed approval UI)
- Database: 1 day (added system_configuration table)
- QA: 3 days (simplified state transitions)
- **Total**: 25 days (~4 weeks)

**Net Change**: -5 hari (dari 5 minggu ke 4 minggu)

---

## Risk Assessment

### Risks Removed (v1 → v2)

1. **Approval Workflow Complexity** ✅ REMOVED
   - Risk: System Admin bottleneck
   - Risk: Delay dalam activation karena approval queue
   - Risk: Complex state transitions causing bugs

2. **Hardcoded Threshold Inflexibility** ✅ REMOVED
   - Risk: Tidak bisa adapt ke different school standards
   - Risk: Perlu redeploy untuk threshold adjustment
   - Risk: One-size-fits-all approach tidak cocok untuk pilot phase

### Risks Added (v1 → v2)

1. **School Admin Error Rate** ⚠️ MODERATE
   - Risk: School admin mungkin activate academic year dengan error
   - Mitigation: Validation rules tetap strict (30-day lead time, no overlap, dll.)
   - Mitigation: Audit trail untuk semua perubahan
   - Mitigation: Ability to archive (soft delete) untuk recovery

2. **Configuration Management** ⚠️ LOW
   - Risk: System configuration table bisa di-incorrectly modified
   - Mitigation: Only System Admin dapat update configuration
   - Mitigation: Audit trail untuk configuration changes
   - Mitigation: Default values untuk fallback

### Overall Risk Assessment

**v1 Risk Level**: MEDIUM-HIGH (complex approval workflow + hardcoded threshold)

**v2 Risk Level**: MEDIUM (simplified workflow + configurable threshold)

**Conclusion**: v2 memiliki risk profile yang lebih baik karena lebih simple dan lebih fleksibel

---

## Validation Checklist

### After Refactoring

- [x] Tidak ada references ke PPDB dalam requirement
- [x] Tidak ada references ke Student Management dalam requirement
- [x] Tidak ada references ke Class Management dalam requirement
- [x] Academic Year workflow simplified ke 3 states
- [x] Approval workflow removed
- [x] CP alignment threshold configurable
- [x] System configuration table added
- [x] All use cases dapat berjalan tanpa student data
- [x] All use cases dapat berjalan tanpa class data
- [x] Timeline reduced dari 5 minggu ke 4 minggu
- [x] Risk profile improved

---

## Approval

**Approved By**: Product Owner  
**Approval Date**: 2026-06-11  
**Approval Status**: APPROVED

**Next Steps**:
1. Update sprint4-implementation-plan.md dengan scope yang direfactor
2. Begin implementation sesuai timeline v2
3. Monitor self-service activation rate sebagai key success metric
4. Collect feedback pada configurable threshold usage

---

**Document Status**: FINAL - READY FOR IMPLEMENTATION
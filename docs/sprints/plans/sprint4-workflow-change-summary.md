# Sprint 4 Academic Year Workflow Change Summary

**Document Version**: 1.0  
**Date**: 2026-06-11  
**Status**: Final  
**Purpose**: Document workflow simplification for Academic Year management from approval-based to self-service

---

## Executive Summary

Academic Year workflow telah disederhanakan dari model approval-based (dengan System Admin approval) ke model self-service (School Admin dapat activate langsung). Perubahan ini mengurangi kompleksitas state machine dari 5 states ke 3 states, dan menghilangkan dependency ke System Admin untuk approval.

**Key Changes**:
- States: 5 → 3
- Approval steps: 2 → 0
- Database columns: 2 removed (approved_by, approved_at)
- API endpoints: 3 removed (submit, approve, reject)
- UI components: Approval workflow removed

---

## Original Workflow (v1)

### State Machine

```
┌─────────────┐
│   DRAFT     │
└──────┬──────┘
       │ Submit
       ↓
┌─────────────────┐
│  UNDER_REVIEW   │
└──────┬──────────┘
       │ Approve
       ↓
┌─────────────────┐
│    APPROVED     │
└──────┬──────────┘
       │ Activate (scheduled)
       ↓
┌─────────────────┐
│     ACTIVE      │
└──────┬──────────┘
       │ End date passed
       ↓
┌─────────────────┐
│    INACTIVE     │
└─────────────────┘
```

### States

| State | Description | Who Can Transition To | Who Can Transition From |
| ----- | ----------- | --------------------- | ---------------------- |
| DRAFT | Initial state, configuration in progress | School Admin (create) | School Admin (submit), System Admin (reject) |
| UNDER_REVIEW | Submitted for approval | School Admin (submit) | System Admin (approve, reject) |
| APPROVED | Approved, waiting for activation date | System Admin (approve) | System (auto-activate) |
| ACTIVE | Currently active academic year | System (auto-activate) | School Admin (deactivate) |
| INACTIVE | Past academic year, read-only | School Admin (deactivate) | - |

### Business Rules

1. **BR-001**: Academic Year Uniqueness
   - Each school can have only one academic year active at any given time

2. **BR-002**: Academic Year Non-Overlap
   - Academic year date ranges for a school cannot overlap with existing academic years

3. **BR-003**: Academic Year Lead Time
   - New academic years must be created at least 30 days before the start date

4. **BR-004**: Approval Required
   - Academic years must be submitted and approved before becoming active

5. **BR-005**: System Admin Approval
   - Only System Admin can approve academic year submissions

6. **BR-006**: Read-Only Past Years
   - Past academic years (INACTIVE) cannot be modified

### Actors

| Actor | Responsibilities | Permissions |
| ----- | --------------- | ------------ |
| School Admin | Create, configure, submit academic years | academic_year:CREATE, academic_year:UPDATE, academic_year:SUBMIT, academic_year:DEACTIVATE |
| System Admin | Approve/reject academic year submissions | academic_year:APPROVE, academic_year:REJECT |
- | Curriculum Admin | Read access to academic years | academic_year:READ |
| Teacher | Read access to academic years | academic_year:READ |

### Database Schema

```sql
CREATE TABLE academic_years (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    school_id UUID NOT NULL,
    name VARCHAR(100) NOT NULL,
    start_date TIMESTAMP WITH TIME ZONE NOT NULL,
    end_date TIMESTAMP WITH TIME ZONE NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'DRAFT' CHECK (status IN ('DRAFT', 'UNDER_REVIEW', 'APPROVED', 'ACTIVE', 'INACTIVE')),
    approved_by UUID,                    -- ← REMOVED IN v2
    approved_at TIMESTAMP WITH TIME ZONE, -- ← REMOVED IN v2
    created_by UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);
```

### API Endpoints

| Method | URL | Description | Permission |
| ------ | --- | ----------- | ---------- |
| POST | /api/v1/academic/academic-years | Create academic year | academic_year:CREATE |
| PUT | /api/v1/academic/academic-years/:id | Update academic year | academic_year:UPDATE |
| POST | /api/v1/academic/academic-years/:id/submit | Submit for approval | academic_year:SUBMIT |
| POST | /api/v1/academic/academic-years/:id/approve | Approve submission | academic_year:APPROVE |
| POST | /api/v1/academic/academic-years/:id/reject | Reject submission | academic_year:REJECT |
| POST | /api/v1/academic/academic-years/:id/deactivate | Deactivate active year | academic_year:DEACTIVATE |
| GET | /api/v1/academic/academic-years | List academic years | academic_year:READ |
| GET | /api/v1/academic/academic-years/:id | Get academic year by ID | academic_year:READ |

---

## New Workflow (v2)

### State Machine

```
┌─────────────┐
│   DRAFT     │
└──────┬──────┘
       │ Activate
       ↓
┌─────────────────┐
│     ACTIVE      │
└──────┬──────────┘
       │ Archive
       ↓
┌─────────────────┐
│   ARCHIVED      │
└─────────────────┘
```

### States

| State | Description | Who Can Transition To | Who Can Transition From |
| ----- | ----------- | --------------------- | ---------------------- |
| DRAFT | Initial state, configuration in progress | School Admin (create) | School Admin (activate) |
| ACTIVE | Currently active academic year | School Admin (activate) | School Admin (archive), System (auto-archive) |
| ARCHIVED | Past academic year, read-only | School Admin (archive), System (auto-archive) | - |

### Business Rules

1. **BR-001**: Academic Year Uniqueness (UNCHANGED)
   - Each school can have only one academic year active at any given time

2. **BR-002**: Academic Year Non-Overlap (UNCHANGED)
   - Academic year date ranges for a school cannot overlap with existing academic years

3. **BR-003**: Academic Year Lead Time (UNCHANGED)
   - New academic years must be created at least 30 days before the start date

4. **BR-004**: Self-Service Activation (NEW)
   - School Admin can activate academic year directly without approval

5. **BR-005**: Activation Validation (NEW)
   - Academic year must have exactly 2 semesters configured before activation
   - Academic year must have no overlaps with existing years

6. **BR-006**: Read-Only Archived Years (RENAMED from INACTIVE)
   - Archived academic years cannot be modified

### Actors

| Actor | Responsibilities | Permissions |
| ----- | --------------- | ------------ |
| School Admin | Create, configure, activate, archive academic years | academic_year:CREATE, academic_year:UPDATE, academic_year:ACTIVATE, academic_year:ARCHIVE |
- | Curriculum Admin | Read access to academic years | academic_year:READ |
| Teacher | Read access to academic years | academic_year:READ |

**Removed**: System Admin no longer involved in academic year workflow

### Database Schema

```sql
CREATE TABLE academic_years (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    school_id UUID NOT NULL,
    name VARCHAR(100) NOT NULL,
    start_date TIMESTAMP WITH TIME ZONE NOT NULL,
    end_date TIMESTAMP WITH TIME ZONE NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'DRAFT' CHECK (status IN ('DRAFT', 'ACTIVE', 'ARCHIVED')),
    -- REMOVED: approved_by UUID
    -- REMOVED: approved_at TIMESTAMP WITH TIME ZONE
    created_by UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);
```

### API Endpoints

| Method | URL | Description | Permission |
| ------ | --- | ----------- | ---------- |
| POST | /api/v1/academic/academic-years | Create academic year | academic_year:CREATE |
| PUT | /api/v1/academic/academic-years/:id | Update academic year | academic_year:UPDATE |
| POST | /api/v1/academic/academic-years/:id/activate | Activate academic year | academic_year:ACTIVATE |
| POST | /api/v1/academic/academic-years/:id/archive | Archive academic year | academic_year:ARCHIVE |
| GET | /api/v1/academic/academic-years | List academic years | academic_year:READ |
| GET | /api/v1/academic/academic-years/:id | Get academic year by ID | academic_year:READ |

**Removed Endpoints**:
- ~~POST /api/v1/academic/academic-years/:id/submit~~
- ~~POST /api/v1/academic/academic-years/:id/approve~~
- ~~POST /api/v1/academic/academic-years/:id/reject~~
- ~~POST /api/v1/academic/academic-years/:id/deactivate~~ (replaced with archive)

---

## Detailed Comparison

### State Changes

| Aspect | v1 (Approval-Based) | v2 (Self-Service) | Change |
| ------ | ------------------- | ------------------ | ------ |
| Total States | 5 | 3 | -40% |
| State Names | DRAFT, UNDER_REVIEW, APPROVED, ACTIVE, INACTIVE | DRAFT, ACTIVE, ARCHIVED | Simplified |
| State Transitions | 7 | 3 | -57% |
| Approval States | 2 (UNDER_REVIEW, APPROVED) | 0 | -100% |
- | Intermediate States | 3 (UNDER_REVIEW, APPROVED, INACTIVE) | 0 (ARCHIVED is terminal) | Simplified |

### Business Rule Changes

| Rule | v1 | v2 | Change |
| ---- | -- | -- | ------ |
| BR-001 (Uniqueness) | Unchanged | Unchanged | No change |
| BR-002 (Non-Overlap) | Unchanged | Unchanged | No change |
| BR-003 (Lead Time) | Unchanged | Unchanged | No change |
| BR-004 (Approval) | Approval required | Self-service activation | REMOVED |
| BR-005 (System Admin) | System Admin approves | No approval | REMOVED |
| BR-006 (Read-Only) | INACTIVE is read-only | ARCHIVED is read-only | RENAMED |
| BR-007 (Activation Validation) | - | Must have 2 semesters | ADDED |
| BR-008 (Archive Validation) | - | No active dependencies | ADDED |

### Permission Changes

| Permission | v1 | v2 | Change |
| ---------- | -- | -- | ------ |
| academic_year:CREATE | School Admin | School Admin | No change |
| academic_year:UPDATE | School Admin | School Admin | No change |
- | academic_year:SUBMIT | School Admin | - | REMOVED |
- | academic_year:APPROVE | System Admin | - | REMOVED |
- | academic_year:REJECT | System Admin | - | REMOVED |
- | academic_year:DEACTIVATE | School Admin | - | REMOVED |
| academic_year:ACTIVATE | - | School Admin | ADDED |
| academic_year:ARCHIVE | - | School Admin | ADDED |
| academic_year:READ | All roles | All roles | No change |

### Database Changes

| Aspect | v1 | v2 | Change |
| ------ | -- | -- | ------ |
| Table: academic_years | 11 columns | 9 columns | -2 columns |
| Column: approved_by | UUID, nullable | - | REMOVED |
| Column: approved_at | TIMESTAMP, nullable | - | REMOVED |
| Status CHECK | 5 states | 3 states | Simplified |
| Indexes | 4 indexes | 4 indexes | No change |
| Foreign Keys | 2 FKs | 2 FKs | No change |

### API Changes

| Aspect | v1 | v2 | Change |
| ------ | -- | -- | ------ |
| Total Endpoints | 9 | 6 | -3 endpoints |
| Create Endpoint | POST /create | POST /create | No change |
| Update Endpoint | PUT /:id | PUT /:id | No change |
| Submit Endpoint | POST /:id/submit | - | REMOVED |
| Approve Endpoint | POST /:id/approve | - | REMOVED |
| Reject Endpoint | POST /:id/reject | - | REMOVED |
| Activate Endpoint | - (auto) | POST /:id/activate | ADDED |
| Deactivate Endpoint | POST /:id/deactivate | - (replaced) | REMOVED |
| Archive Endpoint | - | POST /:id/archive | ADDED |
| List Endpoint | GET / | GET / | No change |
| Get Endpoint | GET /:id | GET /:id | No change |

### Frontend Changes

| Component | v1 | v2 | Change |
| --------- | -- | -- | ------ |
| Academic Year Form | Standard | Standard | No change |
| Submit Button | "Submit for Approval" | "Activate" | Renamed |
| Approval Queue | Separate page | - | REMOVED |
- | Approve Button | - | REMOVED |
- | Reject Button | - | REMOVED |
| Status Badges | 5 badges | 3 badges | Simplified |
| Transition UI | 5 steps | 2 steps | Simplified |
| Archive Button | - | "Archive" | ADDED |

---

## Implementation Impact

### Backend Changes

#### Domain Model

```go
// v1
type AcademicYearStatus string
const (
    AcademicYearStatusDraft      AcademicYearStatus = "DRAFT"
    AcademicYearStatusUnderReview AcademicYearStatus = "UNDER_REVIEW"
    AcademicYearStatusApproved   AcademicYearStatus = "APPROVED"
    AcademicYearStatusActive     AcademicYearStatus = "ACTIVE"
    AcademicYearStatusInactive   AcademicYearStatus = "INACTIVE"
)

// v2
type AcademicYearStatus string
const (
    AcademicYearStatusDraft   AcademicYearStatus = "DRAFT"
    AcademicYearStatusActive  AcademicYearStatus = "ACTIVE"
    AcademicYearStatusArchived AcademicYearStatus = "ARCHIVED"
)
```

#### Domain Service

```go
// v1
func (s *AcademicYearService) SubmitForApproval(ayID UUID) error {
    // Validate academic year
    // Change status to UNDER_REVIEW
    // Send notification to System Admin
}

func (s *AcademicYearService) Approve(ayID UUID, approverID UUID) error {
    // Validate approval authority
    // Change status to APPROVED
    // Set approved_by and approved_at
}

// v2
func (s *AcademicYearService) Activate(ayID UUID) error {
    // Validate academic year has 2 semesters
    // Validate no overlaps
    // Validate no other active year
    // Change status to ACTIVE
}

func (s *AcademicYearService) Archive(ayID UUID) error {
    // Validate no active dependencies
    // Change status to ARCHIVED
}
```

#### Repository

```go
// v1
func (r *AcademicYearRepository) UpdateApproval(id UUID, approverID UUID) error {
    // Update approved_by and approved_at
}

// v2
// No approval-specific repository methods
```

### Frontend Changes

```typescript
// v1
const useAcademicYearActions = () => {
  const submitForApproval = (id: string) => { /* ... */ }
  const approve = (id: string) => { /* ... */ }
  const reject = (id: string) => { /* ... */ }
  const deactivate = (id: string) => { /* ... */ }
  
  return { submitForApproval, approve, reject, deactivate }
}

// v2
const useAcademicYearActions = () => {
  const activate = (id: string) => { /* ... */ }
  const archive = (id: string) => { /* ... */ }
  
  return { activate, archive }
}
```

---

## Migration Impact

### Database Migration

```sql
-- v1 → v2 migration
-- Step 1: Remove columns
ALTER TABLE academic_years DROP COLUMN IF EXISTS approved_by;
ALTER TABLE academic_years DROP COLUMN IF EXISTS approved_at;

-- Step 2: Update status enum constraint
ALTER TABLE academic_years DROP CONSTRAINT IF EXISTS chk_academic_years_status;
ALTER TABLE academic_years 
  ADD CONSTRAINT chk_academic_years_status 
  CHECK (status IN ('DRAFT', 'ACTIVE', 'ARCHIVED'));

-- Step 3: Migrate existing data
UPDATE academic_years 
SET status = 'ACTIVE' 
WHERE status IN ('APPROVED', 'ACTIVE');

UPDATE academic_years 
SET status = 'ARCHIVED' 
WHERE status = 'INACTIVE';

-- Step 4: Remove APPROVED and UNDER_REVIEW records
DELETE FROM academic_years 
WHERE status IN ('UNDER_REVIEW');
```

### Rollback Strategy

```sql
-- v2 → v1 rollback
-- Step 1: Add back columns
ALTER TABLE academic_years ADD COLUMN approved_by UUID;
ALTER TABLE academic_years ADD COLUMN approved_at TIMESTAMP WITH TIME ZONE;

-- Step 2: Update status enum constraint
ALTER TABLE academic_years DROP CONSTRAINT chk_academic_years_status;
ALTER TABLE academic_years 
  ADD CONSTRAINT chk_academic_years_status 
  CHECK (status IN ('DRAFT', 'UNDER_REVIEW', 'APPROVED', 'ACTIVE', 'INACTIVE'));

-- Step 3: Cannot restore approval data (data loss accepted)
```

---

## Risk Analysis

### Risks Removed

1. **Approval Bottleneck** ✅ REMOVED
   - Risk: System Admin unavailability delays activation
   - Impact: Critical
   - Mitigation: Self-service eliminates bottleneck

2. **Complex State Machine** ✅ REMOVED
   - Risk: Bugs in state transitions
   - Impact: Medium
   - Mitigation: Simplified state machine

3. **Approval Queue Management** ✅ REMOVED
   - Risk: Approval queue backlog
   - Impact: Medium
   - Mitigation: No queue needed

### Risks Added

1. **School Admin Errors** ⚠️ MODERATE
   - Risk: School admin may activate incorrect configuration
   - Impact: Medium
   - Mitigation: Strict validation rules (30-day lead time, no overlap, 2 semesters)
   - Mitigation: Audit trail for all changes
   - Mitigation: Ability to archive (soft delete) for recovery

2. **Lack of Governance Oversight** ⚠️ LOW
   - Risk: No System Admin oversight on academic year configuration
   - Impact: Low
   - Mitigation: Audit logs track all changes
   - Mitigation: Principal oversight at organizational level (not software)
   - Mitigation: Configuration can be archived and reviewed

### Net Risk Assessment

**v1 Risk Level**: MEDIUM-HIGH (complex approval workflow + potential bottleneck)

**v2 Risk Level**: MEDIUM (self-service with validation + audit trail)

**Conclusion**: v2 memiliki risk profile yang lebih baik karena lebih simple dan fewer points of failure

---

## Testing Impact

### Unit Test Cases

| Test Suite | v1 Cases | v2 Cases | Change |
| ---------- | -------- | -------- | ------ |
| Academic Year Domain | 15 | 9 | -6 (-40%) |
| Academic Year Service | 12 | 7 | -5 (-42%) |
- | Academic Year Approval | 5 | 0 | -5 (-100%) |
- | Academic Year Validation | 7 | 7 | No change |
| **Total** | **32** | **16** | **-16 (-50%)** |

### Integration Test Cases

| Test Suite | v1 Cases | v2 Cases | Change |
| ---------- | -------- | -------- | ------ |
| Academic Year CRUD | 5 | 5 | No change |
- | Approval Workflow | 4 | 0 | -4 (-100%) |
- | Activation Schedule | 3 | 2 | -1 (-33%) |
| Archive Flow | 0 | 2 | +2 (+100%) |
| **Total** | **12** | **9** | **-3 (-25%)** |

### E2E Test Cases

| Test Suite | v1 Cases | v2 Cases | Change |
| ---------- | -------- | -------- | ------ |
| Academic Year Creation | 2 | 2 | No change |
- | Approval Process | 2 | 0 | -2 (-100%) |
| Activation | 2 | 2 | No change |
| Archival | 0 | 1 | +1 (+100%) |
| **Total** | **6** | **5** | **-1 (-17%)** |

---

## User Experience Impact

### School Admin Experience

**v1 (Approval Workflow)**:
1. Create academic year → DRAFT
2. Configure semesters
3. Click "Submit for Approval" → UNDER_REVIEW
4. Wait for System Admin approval (could take hours/days)
5. System Admin approves → APPROVED
6. Wait for activation date → ACTIVE

**v2 (Self-Service)**:
1. Create academic year → DRAFT
2. Configure semesters
3. Click "Activate" → ACTIVE

**Time Savings**: ~1-3 days (eliminates approval wait time)

**User Satisfaction**: Improved (no dependency on System Admin availability)

### System Admin Experience

**v1 (Approval Workflow)**:
- Must review and approve every academic year submission
- Manage approval queue
- Handle rejected submissions

**v2 (Self-Service)**:
- No involvement in academic year workflow
- Focus on other platform administration tasks

**Time Savings**: ~2-4 hours per school per year (eliminates approval overhead)

---

## Documentation Updates Required

### Documentation Files to Update

1. **sprint4-implementation-plan.md**
   - Part 2: Business Requirements (remove approval workflow)
   - Part 3: Functional Requirements (update Academic Year feature)
   - Part 4: Domain Model (update AcademicYearStatus enum)
   - Part 5: Database Design (remove approved_by, approved_at)
   - Part 6: API Design (remove approval endpoints)
   - Part 7: Frontend Requirements (remove approval UI)
   - Part 8: Security Requirements (update permissions)
   - Part 9: Test Strategy (update test cases)
   - Part 10: Implementation Plan (update tasks and estimates)

2. **DATABASE_SCHEMA_FREEZE_V1.md**
   - Update academic_years table definition
   - Remove approved_by and approved_at columns
   - Update status enum to 3 states

3. **API_DOCUMENTATION.md**
   - Remove approval-related endpoints
   - Add activate and archive endpoints
   - Update response schemas

4. **CHANGELOG.md**
   - Add entry for Sprint 4 workflow simplification
   - Document removed features
   - Document new features

---

## Rollback Plan

### If Self-Service Workflow Causes Issues

**Trigger**:
- School admin errors causing incorrect academic year activations
- Governance concerns arise about lack of oversight

**Rollback Steps**:
1. Re-add approved_by and approved_at columns to academic_years table
2. Update status enum back to 5 states
3. Re-implement approval endpoints
4. Re-add approval UI components
5. Update documentation

**Estimated Time**: 2-3 days

**Data Loss**:
- Approval history will not be restored (acceptable tradeoff)

---

## Success Metrics

### Metrics to Track

| Metric | Target | Measurement Method |
| ------ | ------ | ------------------ |
| Self-Service Activation Rate | >95% | Count of activations via self-service vs total activations |
- | Activation Error Rate | <1% | Count of activation failures / total activations |
| Configuration Time Reduction | >80% | Compare v1 vs v2 configuration time |
| School Admin Satisfaction | >4.0/5.0 | User satisfaction survey |
- | System Admin Time Savings | >90% | Measure System Admin time spent on approvals |

### Monitoring

- Track activation success rate
- Track activation error rate
- Monitor for rapid state transitions (potential errors)
- Audit log monitoring for suspicious activity

---

## Approval

**Approved By**: Product Owner  
**Approval Date**: 2026-06-11  
**Approval Status**: APPROVED

**Next Steps**:
1. Update all documentation files
2. Implement workflow changes in code
3. Update database migration script
4. Update API tests
5. Update UI components
6. Monitor self-service activation rate
7. Collect user feedback on simplified workflow

---

**Document Status**: FINAL - READY FOR IMPLEMENTATION
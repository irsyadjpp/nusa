# 23_CURRICULUM_VERSIONING_ARCHITECTURE.md

## Foundation Document for NUSA Education Platform

**Version**: 1.0
**Date**: June 2026
**Status**: FOUNDATION DOCUMENT
**Alignment**: Aligned with 01_EDUCATION_DOMAIN_MODEL.md, 14_DATABASE_SCHEMA.md, 06_APPLICATION_ARCHITECTURE.md

**Purpose**: Define the curriculum versioning architecture for NUSA to preserve historical data while supporting future curriculum revisions from the Indonesian National Curriculum.

---

# SECTION 1 — Executive Summary

## Why Curriculum Versioning Matters

The Indonesian National Curriculum (Kurikulum Nasional) undergoes periodic updates by the government. These updates may include:
- Changes to learning objectives (CP)
- Restructuring of curriculum hierarchy (Element, Subelement)
- Updates to competency standards
- New phase definitions

NUSA must support these changes while:
- Preserving historical data integrity
- Maintaining traceability of approved artifacts
- Enabling schools to migrate at their own pace
- Supporting coexistence of multiple curriculum versions

## Design Principles

- **Immutability**: Historical curriculum versions remain immutable
- **Traceability**: All artifacts remain traceable to their originating curriculum version
- **Coexistence**: Multiple curriculum versions can coexist in the system
- **Gradual Migration**: Schools can migrate gradually without data loss
- **MVP-Friendly**: Simple, pragmatic approach without event sourcing or microservices

---

# SECTION 2.5 — MVP Simplification Rules

## MVP Scope

For MVP, the curriculum versioning architecture is simplified to reduce complexity while preserving the foundation for future multi-version support.

### MVP Simplification Rule

**Rule**: Exactly one ACTIVE curriculum version exists system-wide.

**Implications**:
- Only one curriculum version can have status = ACTIVE at any time
- All schools use the current ACTIVE curriculum version
- Schools cannot choose curriculum version
- No school_curriculum_mapping table required for MVP
- Curriculum selection is system-wide, not per-school

**Rationale**: MVP focuses on single curriculum deployment to reduce complexity. Multi-version support is deferred to Wave 2.

### MVP Implementation

**Active Version Management**:
- System administrator sets one curriculum version to ACTIVE
- All other versions are DRAFT, DEPRECATED, or ARCHIVED
- New curriculum versions can be imported but must not be set to ACTIVE until ready
- When a new version becomes ACTIVE, the previous version automatically becomes DEPRECATED

**School Behavior**:
- All schools automatically use the ACTIVE curriculum version
- No UI for school curriculum selection
- No per-school curriculum configuration
- All CP queries default to ACTIVE curriculum version

**Artifact Traceability**:
- All artifacts (TP Sets, ATP Sets, Modul Ajar Sets) reference curriculum_version_id
- Traceability preserved even with single ACTIVE version
- Historical artifacts remain linked to their originating curriculum version

### Future Enhancement Rule

**Wave 2: Multi-Version Support**

Future enhancement will enable:
- Multiple ACTIVE curriculum versions to coexist
- School-level curriculum selection via school_curriculum_mapping table
- Schools to migrate at their own pace
- Per-school curriculum configuration UI
- Gradual migration without data loss

**Foundation Preserved**:
- Curriculum version table structure supports multi-version
- All curriculum hierarchy entities include curriculum_version_id
- Artifact traceability via curriculum_version_id preserved
- No schema changes required for Wave 2

---

---

# SECTION 2 — Curriculum Version Model

## Version Entity

### Curriculum Version

**Purpose**: Represents a specific version of the national curriculum with effective dates and status.

**Fields**:

| Field | Type | Constraints | Description |
|-------|------|-------------|-------------|
| `id` | UUID | PRIMARY KEY | Unique version identifier |
| `curriculum_version_code` | VARCHAR(50) | NOT NULL, UNIQUE | Version code (e.g., "KUR2025", "KUR2027") |
| `effective_year` | INTEGER | NOT NULL | Academic year when version becomes effective |
| `effective_from` | DATE | NOT NULL | Effective start date |
| `effective_until` | DATE | NULLABLE | Effective end date (null for current) |
| `status` | VARCHAR(20) | NOT NULL | Version status |
| `description` | TEXT | NULLABLE | Version description |
| `created_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Creation timestamp |
| `updated_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Last update timestamp |

### Status Values

| Status | Description | Transitions |
|--------|-------------|-------------|
| DRAFT | Version under preparation | DRAFT → ACTIVE |
| ACTIVE | Currently in use | ACTIVE → DEPRECATED |
| DEPRECATED | No longer recommended for new use | DEPRECATED → ARCHIVED |
| ARCHIVED | Historical reference only | No transitions |

## Curriculum Hierarchy Entities with Versioning

All curriculum hierarchy entities include `curriculum_version_id` to associate them with a specific curriculum version.

### Subject

**Additional Fields**:
- `curriculum_version_id` (UUID, FOREIGN KEY → curriculum_version(id), NOT NULL)

### Phase

**Additional Fields**:
- `curriculum_version_id` (UUID, FOREIGN KEY → curriculum_version(id), NOT NULL)

### Element

**Additional Fields**:
- `curriculum_version_id` (UUID, FOREIGN KEY → curriculum_version(id), NOT NULL)

### Subelement

**Additional Fields**:
- `curriculum_version_id` (UUID, FOREIGN KEY → curriculum_version(id), NOT NULL)

### CP (Capaian Pembelajaran)

**Additional Fields**:
- `curriculum_version_id` (UUID, FOREIGN KEY → curriculum_version(id), NOT NULL)

---

# SECTION 3 — Versioning Rules

## Core Principles

### 1. Immutability Rule

**Rule**: Once a curriculum version is marked as ACTIVE, it becomes immutable.

**Implications**:
- No modifications to subjects, phases, elements, subelements, or CP in an ACTIVE version
- New versions must be created as DRAFT, then promoted to ACTIVE
- Historical data integrity is preserved

### 2. Historical Reference Rule

**Rule**: Historical TP references remain valid regardless of curriculum version changes.

**Implications**:
- Approved TPs retain their original CP references
- TP → CP foreign key relationships are never broken
- Historical artifacts can always be traced back to their originating curriculum

### 3. Artifact Traceability Rule

**Rule**: Approved artifacts must remain traceable to the originating curriculum version.

**Implications**:
- All artifacts (TP, ATP, Modul Ajar, Assessment, Rubric) store curriculum_version_id
- Reports always use the curriculum version that existed when artifact was approved
- Audit trails preserve the original curriculum context

## Version Lifecycle

```
DRAFT
    ↓ (Admin approval)
ACTIVE
    ↓ (New version becomes active)
DEPRECATED
    ↓ (Archival period expires)
ARCHIVED
```

### Lifecycle Transitions

**DRAFT → ACTIVE**
- Trigger: Administrator approval
- Validation: All required entities (subjects, phases, elements, subelements, CP) are populated
- Effect: Version becomes available for school selection

**ACTIVE → DEPRECATED**
- Trigger: New curriculum version becomes ACTIVE
- Validation: New version is fully populated and approved
- Effect: Existing schools can continue using deprecated version, new schools must use active version

**DEPRECATED → ARCHIVED**
- Trigger: Configured archival period expires (e.g., 5 years)
- Validation: No schools are actively using the version
- Effect: Version becomes read-only reference

---

# SECTION 4 — Artifact Impact Matrix

## Impact Classification

When CP changes (new curriculum version), downstream artifacts are classified by impact level:

| Impact Level | Description | Required Action |
|--------------|-------------|-----------------|
| NO IMPACT | Artifact remains valid | No action required |
| REVIEW REQUIRED | Artifact should be reviewed | Teacher reviews for alignment |
| REGENERATE RECOMMENDED | Artifact should be regenerated | AI regeneration recommended |
| INVALIDATED | Artifact is no longer valid | Must be regenerated |

## Impact Matrix by Artifact Type

### TP (Teaching Plan)

| CP Change Type | Impact Level | Rationale |
|----------------|-------------|-----------|
| CP description update | REVIEW REQUIRED | May affect learning objectives |
| CP learning objectives added | REVIEW REQUIRED | New objectives may need to be addressed |
| CP learning objectives removed | REGENERATE RECOMMENDED | Objectives no longer exist |
| CP competency code change | REGENERATE RECOMMENDED | Alignment may be broken |
| CP hierarchy restructure | INVALIDATED | CP no longer exists in same context |

### ATP (Alur Tujuan Pembelajaran)

| CP Change Type | Impact Level | Rationale |
|----------------|-------------|-----------|
| CP description update | NO IMPACT | ATP sequencing independent of CP description |
| CP learning objectives added | REVIEW REQUIRED | May require new sequence entries |
| CP learning objectives removed | REGENERATE RECOMMENDED | Sequence may have gaps |
| CP competency code change | REVIEW REQUIRED | May affect alignment |
| CP hierarchy restructure | INVALIDATED | Underlying CP structure changed |

### Modul Ajar

| CP Change Type | Impact Level | Rationale |
|----------------|-------------|-----------|
| CP description update | NO IMPACT | Lesson plans independent of CP description |
| CP learning objectives added | REVIEW REQUIRED | May need additional lesson plans |
| CP learning objectives removed | REVIEW REQUIRED | Some lesson plans may be obsolete |
| CP competency code change | NO IMPACT | Lesson plans focus on activities, not codes |
| CP hierarchy restructure | INVALIDATED | Curriculum context for lessons changed |

### Assessment

| CP Change Type | Impact Level | Rationale |
|----------------|-------------|-----------|
| CP description update | NO IMPACT | Assessment items independent |
| CP learning objectives added | REVIEW REQUIRED | May need new assessment items |
| CP learning objectives removed | REVIEW REQUIRED | Some items may be obsolete |
| CP competency code change | REVIEW REQUIRED | Alignment may need verification |
| CP hierarchy restructure | INVALIDATED | Assessment context changed |

### Rubric

| CP Change Type | Impact Level | Rationale |
|----------------|-------------|-----------|
| CP description update | NO IMPACT | Rubric criteria independent |
| CP learning objectives added | REVIEW REQUIRED | May need new criteria |
| CP learning objectives removed | REVIEW REQUIRED | Some criteria may be obsolete |
| CP competency code change | NO IMPACT | Rubric focuses on performance, not codes |
| CP hierarchy restructure | INVALIDATED | Rubric context changed |

### Narrative Report

| CP Change Type | Impact Level | Rationale |
|----------------|-------------|-----------|
| CP description update | NO IMPACT | Reports based on evidence, not CP |
| CP learning objectives added | NO IMPACT | Historical reports immutable |
| CP learning objectives removed | NO IMPACT | Historical reports immutable |
| CP competency code change | NO IMPACT | Historical reports immutable |
| CP hierarchy restructure | NO IMPACT | Historical reports immutable |

**Note**: Narrative Reports are historical records and are never affected by curriculum changes. They always reflect the curriculum version that existed when the report was generated.

---

# SECTION 5 — Migration Strategy

## Coexistence Model

### Multiple Version Support

The system supports coexistence of multiple curriculum versions simultaneously:

```
School A → Curriculum 2025 (ACTIVE)
School B → Curriculum 2025 (ACTIVE)
School C → Curriculum 2027 (ACTIVE)
School D → Curriculum 2025 (DEPRECATED, transitioning)
```

### School-Level Version Selection

Each school can select which curriculum version to use:

**Database Schema**:

```sql
ALTER TABLE schools ADD COLUMN curriculum_version_id UUID FOREIGN KEY → curriculum_version(id);
```

**Default Behavior**:
- New schools default to the ACTIVE curriculum version
- Existing schools retain their current version until they explicitly migrate

### Teacher-Level Version Selection

Within a school, teachers can view curriculum from the school's selected version:

**Query Pattern**:
```sql
SELECT * FROM cp 
WHERE curriculum_version_id = (SELECT curriculum_version_id FROM schools WHERE id = :school_id)
  AND subject_id = :subject_id
  AND phase_id = :phase_id
  AND element_id = :element_id
  AND subelement_id = :subelement_id;
```

## Migration Path

### Phased Adoption

**Phase 1: Preparation**
- New curriculum version created as DRAFT
- Admin reviews and populates hierarchy entities
- Schools notified of upcoming change

**Phase 2: Parallel Operation**
- New version promoted to ACTIVE
- Old version marked as DEPRECATED
- Both versions available in system
- Schools can migrate at their own pace

**Phase 3: Migration**
- Schools update their curriculum_version_id
- Teachers review affected artifacts
- Regeneration of impacted artifacts as needed

**Phase 4: Cleanup**
- All schools migrated to new version
- Old version marked as ARCHIVED
- Read-only access maintained for historical reference

### Migration Triggers

**Manual Migration**:
- School administrator initiates migration
- System validates readiness (no in-progress artifacts)
- Migration proceeds with artifact impact analysis

**Automatic Migration** (Optional, Future Wave):
- System detects school with DEPRECATED version
- Sends migration recommendations
- Provides guided migration workflow

## Rollback Support

### Version Rollback

**Scenario**: New curriculum version has critical issues

**Process**:
1. New version marked as DEPRECATED
2. Previous version (if still in system) can be reactivated
3. Schools can rollback their curriculum_version_id
4. Historical artifacts remain valid

**Limitation**: If previous version has been archived, rollback requires re-import from backup

---

# SECTION 6 — Database Changes

## New Tables

### curriculum_versions

```sql
CREATE TABLE curriculum_versions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    curriculum_version_code VARCHAR(50) NOT NULL UNIQUE,
    effective_year INTEGER NOT NULL,
    effective_from DATE NOT NULL,
    effective_until DATE,
    status VARCHAR(20) NOT NULL CHECK (status IN ('DRAFT', 'ACTIVE', 'DEPRECATED', 'ARCHIVED')),
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_curriculum_versions_code ON curriculum_versions(curriculum_version_code);
CREATE INDEX idx_curriculum_versions_status ON curriculum_versions(status);
CREATE INDEX idx_curriculum_versions_effective_year ON curriculum_versions(effective_year);
```

## Modified Tables

### curriculum_subjects

```sql
ALTER TABLE curriculum_subjects ADD COLUMN curriculum_version_id UUID NOT NULL;
ALTER TABLE curriculum_subjects ADD CONSTRAINT fk_curriculum_subjects_version 
    FOREIGN KEY (curriculum_version_id) REFERENCES curriculum_versions(id);
CREATE INDEX idx_curriculum_subjects_version ON curriculum_subjects(curriculum_version_id);
```

### curriculum_phases

```sql
ALTER TABLE curriculum_phases ADD COLUMN curriculum_version_id UUID NOT NULL;
ALTER TABLE curriculum_phases ADD CONSTRAINT fk_curriculum_phases_version 
    FOREIGN KEY (curriculum_version_id) REFERENCES curriculum_versions(id);
CREATE INDEX idx_curriculum_phases_version ON curriculum_phases(curriculum_version_id);
```

### curriculum_elements

```sql
ALTER TABLE curriculum_elements ADD COLUMN curriculum_version_id UUID NOT NULL;
ALTER TABLE curriculum_elements ADD CONSTRAINT fk_curriculum_elements_version 
    FOREIGN KEY (curriculum_version_id) REFERENCES curriculum_versions(id);
CREATE INDEX idx_curriculum_elements_version ON curriculum_elements(curriculum_version_id);
```

### curriculum_subelements

```sql
ALTER TABLE curriculum_subelements ADD COLUMN curriculum_version_id UUID NOT NULL;
ALTER TABLE curriculum_subelements ADD CONSTRAINT fk_curriculum_subelements_version 
    FOREIGN KEY (curriculum_version_id) REFERENCES curriculum_versions(id);
CREATE INDEX idx_curriculum_subelements_version ON curriculum_subelements(curriculum_version_id);
```

### cp

```sql
ALTER TABLE cp ADD COLUMN curriculum_version_id UUID NOT NULL;
ALTER TABLE cp ADD CONSTRAINT fk_cp_version 
    FOREIGN KEY (curriculum_version_id) REFERENCES curriculum_versions(id);
CREATE INDEX idx_cp_version ON cp(curriculum_version_id);
```

## Artifact Tables (Curriculum Traceability)

### tp

```sql
ALTER TABLE tp ADD COLUMN curriculum_version_id UUID NOT NULL;
ALTER TABLE tp ADD CONSTRAINT fk_tp_version 
    FOREIGN KEY (curriculum_version_id) REFERENCES curriculum_versions(id);
CREATE INDEX idx_tp_version ON tp(curriculum_version_id);
```

### atp

```sql
ALTER TABLE atp ADD COLUMN curriculum_version_id UUID NOT NULL;
ALTER TABLE atp ADD CONSTRAINT fk_atp_version 
    FOREIGN KEY (curriculum_version_id) REFERENCES curriculum_versions(id);
CREATE INDEX idx_atp_version ON atp(curriculum_version_id);
```

### modul_ajar

```sql
ALTER TABLE modul_ajar ADD COLUMN curriculum_version_id UUID NOT NULL;
ALTER TABLE modul_ajar ADD CONSTRAINT fk_modul_ajar_version 
    FOREIGN KEY (curriculum_version_id) REFERENCES curriculum_versions(id);
CREATE INDEX idx_modul_ajar_version ON modul_ajar(curriculum_version_id);
```

### assessment

```sql
ALTER TABLE assessment ADD COLUMN curriculum_version_id UUID NOT NULL;
ALTER TABLE assessment ADD CONSTRAINT fk_assessment_version 
    FOREIGN KEY (curriculum_version_id) REFERENCES curriculum_versions(id);
CREATE INDEX idx_assessment_version ON assessment(curriculum_version_id);
```

### rubric

```sql
ALTER TABLE rubric ADD COLUMN curriculum_version_id UUID NOT NULL;
ALTER TABLE rubric ADD CONSTRAINT fk_rubric_version 
    FOREIGN KEY (curriculum_version_id) REFERENCES curriculum_versions(id);
CREATE INDEX idx_rubric_version ON rubric(curriculum_version_id);
```

### narrative_report

```sql
ALTER TABLE narrative_report ADD COLUMN curriculum_version_id UUID NOT NULL;
ALTER TABLE narrative_report ADD CONSTRAINT fk_narrative_report_version 
    FOREIGN KEY (curriculum_version_id) REFERENCES curriculum_versions(id);
CREATE INDEX idx_narrative_report_version ON narrative_report(curriculum_version_id);
```

## School Table Modification

### schools

```sql
ALTER TABLE schools ADD COLUMN curriculum_version_id UUID;
ALTER TABLE schools ADD CONSTRAINT fk_schools_version 
    FOREIGN KEY (curriculum_version_id) REFERENCES curriculum_versions(id);
CREATE INDEX idx_schools_version ON schools(curriculum_version_id);
```

---

# SECTION 7 — Reporting Rules

## Historical Reporting Principle

**Rule**: Historical reports must always use the curriculum version that existed when the artifact was approved.

### Implementation

**Report Query Pattern**:

```sql
-- When generating a report for an approved TP
SELECT 
    tp.*,
    cp.*,
    curriculum_versions.curriculum_version_code,
    curriculum_versions.effective_year
FROM tp
JOIN cp ON tp.cp_id = cp.id
JOIN curriculum_versions ON tp.curriculum_version_id = curriculum_versions.id
WHERE tp.id = :tp_id
  AND tp.status = 'APPROVED';
```

**Key Points**:
- Report queries join through the artifact's curriculum_version_id
- Never use the school's current curriculum_version_id for historical reports
- Always preserve the original curriculum context

## Version Display in Reports

### Report Header Information

All reports include curriculum version information:

```json
{
  "report_metadata": {
    "generated_at": "2026-06-05T12:00:00Z",
    "curriculum_version": {
      "code": "KUR2025",
      "effective_year": 2025,
      "effective_from": "2025-01-01",
      "status": "ACTIVE"
    }
  },
  "artifact_data": { ... }
}
```

### Cross-Version Reporting

**Scenario**: School wants to compare student performance across curriculum versions

**Approach**:
- Generate separate reports for each version
- Use curriculum_version_code as grouping dimension
- Do not attempt to normalize across versions (different structures)

## Audit Trail Preservation

### Change Tracking

When curriculum_version_id changes for a school:

```sql
INSERT INTO school_curriculum_history (
    school_id,
    old_curriculum_version_id,
    new_curriculum_version_id,
    changed_at,
    changed_by
) VALUES (
    :school_id,
    :old_version_id,
    :new_version_id,
    NOW(),
    :user_id
);
```

**Purpose**: Maintain audit trail of curriculum version changes for compliance and analysis

---

# SECTION 8 — Lifecycle Diagram

## Version Lifecycle Flowchart

```
┌─────────────┐
│   DRAFT     │
│ (Preparation)│
└──────┬──────┘
       │ Admin Approval
       │ Validation Complete
       ↓
┌─────────────┐
│   ACTIVE    │
│ (In Use)    │
└──────┬──────┘
       │ New Version Activated
       ↓
┌─────────────┐
│ DEPRECATED  │
│ (Phasing Out)│
└──────┬──────┘
       │ Archival Period Expires
       │ No Active Schools
       ↓
┌─────────────┐
│  ARCHIVED   │
│ (Read-Only) │
└─────────────┘
```

## School Migration Flowchart

```
┌─────────────────┐
│ School on       │
│ Curriculum 2025 │
└────────┬────────┘
         │
         │ New Version Available
         ↓
┌─────────────────┐
│ Review Impact   │
│ Analysis        │
└────────┬────────┘
         │
         │ Decision: Migrate?
         ↓ Yes
┌─────────────────┐
│ Update School   │
│ curriculum_     │
│ version_id      │
└────────┬────────┘
         │
         │ Review Artifacts
         ↓
┌─────────────────┐
│ Regenerate      │
│ Impacted        │
│ Artifacts       │
└────────┬────────┘
         │
         │ Complete Migration
         ↓
┌─────────────────┐
│ School on       │
│ Curriculum 2027 │
└─────────────────┘
```

---

# SECTION 9 — Migration Rules

## Data Migration

### Initial Migration (Existing Data)

**Scenario**: System already has curriculum data without versioning

**Process**:
1. Create curriculum_version record for existing data (e.g., "KUR2025")
2. Backfill curriculum_version_id for all existing curriculum entities
3. Backfill curriculum_version_id for all existing artifacts
4. Set school curriculum_version_id to the new version

**SQL Example**:
```sql
-- Step 1: Create version
INSERT INTO curriculum_versions (
    id, curriculum_version_code, effective_year, effective_from, status
) VALUES (
    gen_random_uuid(), 'KUR2025', 2025, '2025-01-01', 'ACTIVE'
);

-- Step 2: Backfill curriculum entities
UPDATE curriculum_subjects SET curriculum_version_id = (SELECT id FROM curriculum_versions WHERE curriculum_version_code = 'KUR2025');
UPDATE curriculum_phases SET curriculum_version_id = (SELECT id FROM curriculum_versions WHERE curriculum_version_code = 'KUR2025');
UPDATE curriculum_elements SET curriculum_version_id = (SELECT id FROM curriculum_versions WHERE curriculum_version_code = 'KUR2025');
UPDATE curriculum_subelements SET curriculum_version_id = (SELECT id FROM curriculum_versions WHERE curriculum_version_code = 'KUR2025');
UPDATE cp SET curriculum_version_id = (SELECT id FROM curriculum_versions WHERE curriculum_version_code = 'KUR2025');

-- Step 3: Backfill artifacts
UPDATE tp SET curriculum_version_id = (SELECT id FROM curriculum_versions WHERE curriculum_version_code = 'KUR2025');
UPDATE atp SET curriculum_version_id = (SELECT id FROM curriculum_versions WHERE curriculum_version_code = 'KUR2025');
UPDATE modul_ajar SET curriculum_version_id = (SELECT id FROM curriculum_versions WHERE curriculum_version_code = 'KUR2025');
UPDATE assessment SET curriculum_version_id = (SELECT id FROM curriculum_versions WHERE curriculum_version_code = 'KUR2025');
UPDATE rubric SET curriculum_version_id = (SELECT id FROM curriculum_versions WHERE curriculum_version_code = 'KUR2025');
UPDATE narrative_report SET curriculum_version_id = (SELECT id FROM curriculum_versions WHERE curriculum_version_code = 'KUR2025');

-- Step 4: Set school version
UPDATE schools SET curriculum_version_id = (SELECT id FROM curriculum_versions WHERE curriculum_version_code = 'KUR2025');
```

### New Version Import

**Process**:
1. Create new curriculum_version record as DRAFT
2. Import new curriculum hierarchy entities with new version_id
3. Validate completeness
4. Promote to ACTIVE
5. Notify schools

## Validation Rules

### Version Completeness Check

Before promoting a version to ACTIVE:

```sql
-- Check that all required entities exist
SELECT 
    (SELECT COUNT(*) FROM curriculum_subjects WHERE curriculum_version_id = :version_id) as subject_count,
    (SELECT COUNT(*) FROM curriculum_phases WHERE curriculum_version_id = :version_id) as phase_count,
    (SELECT COUNT(*) FROM curriculum_elements WHERE curriculum_version_id = :version_id) as element_count,
    (SELECT COUNT(*) FROM curriculum_subelements WHERE curriculum_version_id = :version_id) as subelement_count,
    (SELECT COUNT(*) FROM cp WHERE curriculum_version_id = :version_id) as cp_count;
```

**Validation**: All counts must be > 0

### School Migration Readiness Check

Before a school migrates:

```sql
-- Check for in-progress artifacts
SELECT COUNT(*) as in_progress_count
FROM tp
WHERE user_id IN (SELECT id FROM users WHERE school_id = :school_id)
  AND status = 'DRAFT';
```

**Validation**: in_progress_count must be 0 (no in-progress artifacts)

---

# SECTION 10 — API Considerations

## Version-Aware Endpoints

### Curriculum Query Endpoints

All curriculum query endpoints must accept optional `curriculum_version_id` parameter:

```
GET /api/v1/curriculum/subjects?curriculum_version_id={version_id}
GET /api/v1/curriculum/phases?curriculum_version_id={version_id}
GET /api/v1/curriculum/elements?curriculum_version_id={version_id}&subject_id={subject_id}&phase_id={phase_id}
GET /api/v1/curriculum/subelements?curriculum_version_id={version_id}&element_id={element_id}
GET /api/v1/curriculum/cp?curriculum_version_id={version_id}&subject_id={subject_id}&phase_id={phase_id}&element_id={element_id}&subelement_id={subelement_id}
```

**Default Behavior**: If `curriculum_version_id` not provided, use school's current version

### School Version Management

```
GET /api/v1/schools/{school_id}/curriculum-version
PUT /api/v1/schools/{school_id}/curriculum-version
```

**Request Body**:
```json
{
  "curriculum_version_id": "version_uuid"
}
```

### Version List Endpoint

```
GET /api/v1/curriculum/versions
```

**Response**:
```json
{
  "success": true,
  "data": {
    "versions": [
      {
        "id": "version_uuid",
        "curriculum_version_code": "KUR2025",
        "effective_year": 2025,
        "effective_from": "2025-01-01",
        "effective_until": null,
        "status": "ACTIVE",
        "description": "Kurikulum 2025"
      }
    ]
  }
}
```

---

# SECTION 11 — Implementation Notes

## MVP Considerations

### MVP Wave 1 Scope

**Included**:
- Basic curriculum_version table
- curriculum_version_id on curriculum entities
- Single ACTIVE version support
- School-level version selection

**Deferred to Future Waves**:
- Automated migration workflows
- Cross-version analytics
- Version comparison tools
- Advanced impact analysis

### Simplified Approach

For MVP, the system can start with:
1. Single curriculum version (no versioning initially)
2. Add versioning schema when first curriculum update is needed
3. Manual migration process for initial rollout

This keeps MVP simple while preserving upgrade path.

## Performance Considerations

### Indexing Strategy

All curriculum_version_id fields are indexed for efficient querying:
- Single-version queries: `WHERE curriculum_version_id = ?`
- Cross-version queries: `GROUP BY curriculum_version_id`

### Query Optimization

**Pattern**: Always include curriculum_version_id in WHERE clauses for curriculum entities

```sql
-- Efficient
SELECT * FROM cp 
WHERE curriculum_version_id = :version_id 
  AND subject_id = :subject_id;

-- Less efficient (avoids if possible)
SELECT * FROM cp 
WHERE subject_id = :subject_id;
```

---

# SECTION 12 — Summary

## Key Deliverables

✓ **Versioning Architecture**: Complete model for curriculum versioning with lifecycle management
✓ **Lifecycle Diagram**: Visual representation of version and migration flows
✓ **Migration Rules**: Clear rules for data migration and school migration
✓ **Database Impact**: Comprehensive schema changes for versioning support
✓ **Reporting Rules**: Guidelines for historical reporting with version preservation

## Design Principles Adhered To

- ✓ No event sourcing
- ✓ No microservices
- ✓ MVP-friendly approach
- ✓ Immutable historical versions
- ✓ Complete artifact traceability
- ✓ Gradual migration support
- ✓ Coexistence of multiple versions

## Next Steps

1. Implement database schema changes
2. Update API contracts to include version parameters
3. Develop version management UI for administrators
4. Create school migration workflow
5. Update reporting to use version-aware queries
6. Document migration procedures for operators

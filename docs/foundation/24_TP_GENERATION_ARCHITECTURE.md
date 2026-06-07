# 24_TP_GENERATION_ARCHITECTURE.md

## Foundation Document for NUSA Education Platform

**Version**: 1.0
**Date**: June 2026
**Status**: FOUNDATION DOCUMENT
**Alignment**: Aligned with 01_EDUCATION_DOMAIN_MODEL.md, 14_DATABASE_SCHEMA.md, 15_AI_PROMPT_SPECIFICATION.md, 06_APPLICATION_ARCHITECTURE.md

**Purpose**: Define the complete TP (Teaching Plan) Generation Architecture for NUSA, establishing the workflow, lifecycle, versioning, traceability, and human review requirements for AI-assisted TP generation.

---

# SECTION 1 — Executive Summary

## Why TP Generation Architecture Matters

Teaching Plans (TP) are the foundational artifacts that translate national curriculum (CP) into actionable teaching objectives. TP generation is the first step in the curriculum-to-report pipeline and sets the foundation for all subsequent artifacts (ATP, Modul Ajar, Assessment, Rubric).

## Core Principles

- **AI-Assisted, Not AI-Autonomous**: AI generates recommendations, teachers make decisions
- **Human Authority**: Teachers remain the final authority over TP content
- **TP Set as First-Class Entity**: One AI generation session creates a TP Set containing multiple TP Items
- **TP Set-Level Approval**: Approval applies at TP Set level, not individual TP level
- **Traceability**: Every TP Set and TP Item must be traceable to its originating curriculum hierarchy and AI generation
- **Version Control**: TP Set versions support iteration while preserving history
- **Review Required**: AI output cannot become official TP Set without human review

---

# SECTION 2 — TP Generation Principles

## Principle 1: AI Generates Recommendations

**Rule**: AI generates TP recommendations based on CP, curriculum context, and teacher preferences.

**Implications**:
- AI provides structured, actionable teaching objectives
- AI suggests time allocation and prerequisites
- AI considers curriculum hierarchy context (Subject → Phase → Element → Subelement → CP)
- AI output is always a recommendation, never final

## Principle 2: Teacher Remains Final Authority

**Rule**: Teachers have complete authority to accept, modify, or reject AI-generated TP recommendations.

**Implications**:
- Teachers can edit any AI-generated content
- Teachers can regenerate TP sets as needed
- Teachers can create TP manually without AI assistance
- Teacher approval is required for TP to become official

## Principle 3: AI Never Creates Official TP Automatically

**Rule**: AI-generated TP Set never becomes official without explicit teacher approval.

**Implications**:
- All AI-generated TP Sets start in DRAFT status
- DRAFT TP Sets cannot be used for downstream artifact generation
- Only APPROVED TP Sets can generate ATP, Modul Ajar, etc.
- System enforces approval workflow at TP Set level
- Individual TP Items within a TP Set inherit the TP Set's status

---

# SECTION 3 — Input Model

## Generation Context

The TP generation input must include complete curriculum hierarchy and teaching context:

### Curriculum Hierarchy

```json
{
  "curriculum_hierarchy": {
    "subject": {
      "id": "uuid",
      "code": "MTK",
      "name": "Matematika",
      "name_en": "Mathematics"
    },
    "phase": {
      "id": "uuid",
      "code": "FASE_E",
      "name": "Fase E",
      "name_en": "Phase E",
      "grade_level_start": 10,
      "grade_level_end": 12
    },
    "element": {
      "id": "uuid",
      "code": "ALGEBRA",
      "name": "Algebra",
      "name_en": "Algebra",
      "description": "Algebraic thinking and operations"
    },
    "subelement": {
      "id": "uuid",
      "code": "LINEAR_EQ",
      "name": "Linear Equations",
      "name_en": "Linear Equations",
      "description": "Linear equations and inequalities"
    }
  }
}
```

### CP Data

```json
{
  "cp": {
    "id": "uuid",
    "code": "CP.10.1.1",
    "description": "Students can solve linear equations and inequalities",
    "competency_code": "CP.10.1",
    "learning_objectives": [
      {
        "id": "uuid",
        "code": "LO.10.1",
        "description": "Students understand algebraic concepts",
        "competency_code": "CP.10.1"
      }
    ],
    "competency_standards": [
      {
        "id": "uuid",
        "code": "CS.10.1",
        "description": "Algebraic thinking"
      }
    ],
    "time_allocation": {
      "total_hours": 120,
      "hours_per_week": 4
    },
    "version": "2026"
  }
}
```

### Teaching Context

```json
{
  "class_info": {
    "grade": "10",
    "subject": "Matematika",
    "academic_year": "2026"
  },
  "teaching_schedule": {
    "hours_per_week": 4,
    "weeks_per_semester": 18
  },
  "school_context": {
    "school_type": "SMA",
    "school_level": "SMA",
    "region": "DKI Jakarta"
  },
  "preferences": {
    "focus_areas": ["algebra", "geometry"],
    "teaching_style": "interactive",
    "differentiation_needs": ["visual_learners", "advanced_students"]
  }
}
```

## Complete Input Schema

```json
{
  "curriculum_hierarchy": { /* hierarchy object */ },
  "cp": { /* CP object */ },
  "class_info": { /* class info object */ },
  "teaching_schedule": { /* schedule object */ },
  "school_context": { /* school context object */ },
  "preferences": { /* teacher preferences object */ }
}
```

---

# SECTION 4 — Output Model

## TP Set Concept

**Definition**: A TP Set is a first-class domain entity representing one AI generation session for a CP. It contains multiple TP Items.

**Rationale**:
- CP often contains multiple learning objectives that span different topics
- Grouping objectives into logical TP Items improves teachability
- Sequencing TP Items creates coherent learning progression
- Teachers can implement TP Items incrementally
- TP Set-level approval ensures atomicity of the generation session
- TP Set versioning supports regeneration without losing history

## TP Set Structure

```json
{
  "tp_set": {
    "id": "uuid",
    "cp_id": "uuid",
    "version_no": 1,
    "status": "DRAFT",
    "generation_source": "AI_GENERATED",
    "generation_reason": "Initial generation",
    "generated_by": "uuid",
    "ai_generation_id": "uuid",
    "tps": [
      {
        "id": "uuid",
        "tp_set_id": "uuid",
        "sequence_number": 1,
        "title": "Introduction to Linear Equations",
        "learning_objectives": [
          {
            "id": "uuid",
            "cp_objective_id": "uuid",
            "description": "Students understand the concept of linear equations",
            "time_allocation_hours": 4
          }
        ],
        "time_allocation": {
          "total_hours": 12,
          "hours_per_week": 4
        },
        "prerequisites": [],
        "estimated_weeks": 3
      },
      {
        "id": "uuid",
        "sequence_number": 2,
        "title": "Solving Linear Equations",
        "learning_objectives": [
          {
            "id": "uuid",
            "cp_objective_id": "uuid",
            "description": "Students can solve linear equations using various methods",
            "time_allocation_hours": 6
          }
        ],
        "time_allocation": {
          "total_hours": 16,
          "hours_per_week": 4
        },
        "prerequisites": [
          {
            "objective_id": "uuid",
            "required_for": ["uuid"]
          }
        ],
        "estimated_weeks": 4
      }
    ],
    "total_tps": 2,
    "total_hours": 28,
    "estimated_weeks": 7,
    "ai_metadata": {
      "confidence_score": 0.92,
      "generated_at": "2026-06-05T12:00:00Z",
      "agent_version": "1.0",
      "model_version": "gpt-4",
      "prompt_version": "1.0"
    }
  }
}
```

## TP Set Characteristics

### Size Range

**Typical Range**: 3–10 TP Items per TP Set

**Factors Influencing Size**:
- Number of CP learning objectives
- Complexity of subject matter
- Time allocation (total hours / hours per week)
- Teacher preferences for granularity

**Guidelines**:
- Each TP Item should cover 2–6 learning objectives
- Each TP Item should span 2–4 weeks of instruction
- Total TP Set should cover the entire CP within the semester

### Sequencing

TP Items within a TP Set are sequenced to create logical learning progression:

```
TP Item 1 (Foundational Concepts)
  ↓
TP Item 2 (Application)
  ↓
TP Item 3 (Advanced Topics)
  ↓
...
```

**Seencing Rules**:
- Prerequisites defined between TPs
- Foundational concepts covered in earlier TPs
- Complex topics in later TPs
- Clear progression from simple to complex

---

# SECTION 5 — TP Generation Lifecycle

## Lifecycle States

```
DRAFT
  ↓ (Teacher Review)
REVIEW
  ↓ (Teacher Edit)
DRAFT
  ↓ (Teacher Approval)
APPROVED
  ↓ (Regeneration)
DRAFT (new version)
  ↓ (Archival)
ARCHIVED
```

## State Definitions

### DRAFT

**Description**: Initial state for AI-generated or manually created TP

**Characteristics**:
- Not yet reviewed by teacher
- Can be edited freely
- Cannot be used for downstream artifact generation
- No version number assigned

**Transitions**:
- DRAFT → REVIEW (teacher initiates review)
- DRAFT → APPROVED (teacher approves without review step)

### REVIEW

**Description**: TP under teacher review

**Characteristics**:
- Teacher is actively reviewing content
- Edits allowed during review
- Can be sent back to DRAFT for further work
- Can be approved directly from REVIEW

**Transitions**:
- REVIEW → DRAFT (teacher requests changes)
- REVIEW → APPROVED (teacher approves)

### APPROVED

**Description**: Official TP approved by teacher

**Characteristics**:
- Can be used for downstream artifact generation (ATP, Modul Ajar, etc.)
- Becomes immutable (cannot be edited directly)
- Version number assigned
- Can be regenerated (creates new version)

**Transitions**:
- APPROVED → DRAFT (regeneration creates new version)
- APPROVED → ARCHIVED (archived after period of inactivity)

### ARCHIVED

**Description**: Historical TP no longer in active use

**Characteristics**:
- Read-only reference
- Preserved for historical reporting
- Cannot be regenerated
- Can be viewed but not used for new artifacts

**Transitions**:
- No outgoing transitions (terminal state)

## Lifecycle Diagram

```
┌─────────┐
│  DRAFT  │
│ (AI Gen)│
└────┬────┘
     │ Teacher Review
     ↓
┌─────────┐
│ REVIEW  │
│(Teacher)│
└────┬────┘
     │
     ├─→ Edit → DRAFT
     │
     └─→ Approve
         ↓
     ┌─────────┐
     │ APPROVED│
     │(Official)│
     └────┬────┘
          │
          ├─→ Regenerate → DRAFT (new version)
          │
          └─→ Archive
              ↓
          ┌─────────┐
          │ARCHIVED │
          │(History)│
          └─────────┘
```

---

# SECTION 6 — Regeneration Rules

## Regeneration Triggers

Teachers may regenerate TP sets in the following scenarios:

### Full Regeneration

**Trigger**: Teacher requests complete regeneration of TP Set

**Process**:
1. Create new DRAFT version of TP Set
2. Send same input to AI agent
3. Generate new TP recommendations
4. Preserve previous APPROVED version
5. Teacher reviews new DRAFT

**Rule**: Regeneration never overwrites approved TP. Creates new version instead.

### Partial Regeneration

**Trigger**: Teacher regenerates specific TP within TP Set

**Process**:
1. Create new DRAFT version of specific TP
2. Send CP and context to AI agent
3. Generate new TP recommendations
4. Preserve other TPs in TP Set
5. Teacher reviews regenerated TP

**Rule**: Only selected TP is regenerated, others remain unchanged.

### Context Change Regeneration

**Trigger**: Underlying CP or curriculum context changes

**Process**:
1. System detects CP change
2. Notifies teacher of potential impact
3. Teacher initiates regeneration if needed
4. New version created with updated context
5. Teacher reviews and approves

**Rule**: Automatic regeneration not performed. Teacher must initiate.

## Regeneration Constraints

### Cannot Overwrite Approved TP

**Rule**: Approved TP cannot be modified directly. Regeneration creates new version.

**Implementation**:
```sql
-- When regenerating, create new version
INSERT INTO tp (parent_version_id, ...) 
VALUES (:approved_tp_id, ...);

-- Original approved TP remains unchanged
UPDATE tp SET is_current_version = false WHERE id = :approved_tp_id;
```

### Version Continuity

**Rule**: Regenerated TP maintains traceability to original CP and curriculum hierarchy.

**Implementation**:
- curriculum_version_id preserved from original
- cp_id, subject_id, phase_id, element_id, subelement_id preserved
- New version_no incremented
- parent_version_id references original

### Approval Required

**Rule**: Regenerated TP requires approval before becoming official.

**Process**:
- Regenerated TP starts in DRAFT status
- Teacher must review and approve
- Cannot skip approval step

---

# SECTION 7 — Versioning

## Version Fields

### TP Version Attributes

| Field | Type | Description |
|-------|------|-------------|
| `version_no` | INTEGER | Version number (starts at 1) |
| `is_current_version` | BOOLEAN | Whether this is the current version |
| `parent_version_id` | UUID | Reference to previous version (null for initial) |

### Version Behavior

**Initial Version**:
- version_no = 1
- is_current_version = true
- parent_version_id = null

**Subsequent Versions**:
- version_no = parent.version_no + 1
- is_current_version = true
- parent_version_id = parent.id
- Previous version: is_current_version = false

## Version Query Patterns

### Get Current Version

```sql
SELECT * FROM tp 
WHERE cp_id = :cp_id 
  AND user_id = :user_id 
  AND is_current_version = true;
```

### Get Version History

```sql
SELECT * FROM tp 
WHERE cp_id = :cp_id 
  AND user_id = :user_id 
ORDER BY version_no DESC;
```

### Get Specific Version

```sql
SELECT * FROM tp 
WHERE id = :tp_id;
```

## Version Coexistence

Multiple versions of the same TP can coexist:

```
TP v1 (APPROVED, is_current_version = false) ← Historical reference
TP v2 (DRAFT, is_current_version = true) ← Work in progress
```

**Rules**:
- Only one version can be current (is_current_version = true)
- Historical versions remain accessible
- Downstream artifacts reference specific version IDs

---

# SECTION 8 — Traceability

## Required Traceability Fields

Every TP must store the following traceability information:

### Curriculum Traceability

| Field | Type | Description |
|-------|------|-------------|
| `cp_id` | UUID | Reference to originating CP |
| `subject_id` | UUID | Reference to Subject in hierarchy |
| `phase_id` | UUID | Reference to Phase in hierarchy |
| `element_id` | UUID | Reference to Element in hierarchy |
| `subelement_id` | UUID | Reference to Subelement in hierarchy |
| `curriculum_version_id` | UUID | Reference to Curriculum Version |

### Generation Traceability

| Field | Type | Description |
|-------|------|-------------|
| `prompt_version` | VARCHAR(20) | Version of AI prompt template used |
| `model_version` | VARCHAR(50) | AI model version used (e.g., "gpt-4") |
| `ai_generation_id` | UUID | Unique identifier for this generation request |

### Audit Traceability

| Field | Type | Description |
|-------|------|-------------|
| `user_id` | UUID | Teacher who requested generation |
| `created_at` | TIMESTAMP | Generation timestamp |
| `updated_at` | TIMESTAMP | Last modification timestamp |
| `approved_at` | TIMESTAMP | Approval timestamp (null if not approved) |
| `approved_by` | UUID | User who approved (null if not approved) |

## Complete TP Schema with Traceability

```sql
CREATE TABLE tp (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    cp_id UUID NOT NULL,
    subject_id UUID NOT NULL,
    phase_id UUID NOT NULL,
    element_id UUID NOT NULL,
    subelement_id UUID NOT NULL,
    curriculum_version_id UUID NOT NULL,
    
    user_id UUID NOT NULL,
    status VARCHAR(20) NOT NULL,
    
    learning_objectives JSONB NOT NULL,
    time_allocation JSONB NOT NULL,
    prerequisites JSONB,
    
    -- Versioning
    version_no INTEGER NOT NULL DEFAULT 1,
    is_current_version BOOLEAN NOT NULL DEFAULT true,
    parent_version_id UUID,
    
    -- AI Metadata
    ai_confidence_score DECIMAL(3,2),
    ai_generated_at TIMESTAMP WITH TIME ZONE,
    ai_agent_version VARCHAR(20),
    prompt_version VARCHAR(20),
    model_version VARCHAR(50),
    ai_generation_id UUID,
    
    -- Audit
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    approved_at TIMESTAMP WITH TIME ZONE,
    approved_by UUID,
    
    -- Foreign Keys
    FOREIGN KEY (cp_id) REFERENCES cp(id),
    FOREIGN KEY (subject_id) REFERENCES curriculum_subjects(id),
    FOREIGN KEY (phase_id) REFERENCES curriculum_phases(id),
    FOREIGN KEY (element_id) REFERENCES curriculum_elements(id),
    FOREIGN KEY (subelement_id) REFERENCES curriculum_subelements(id),
    FOREIGN KEY (curriculum_version_id) REFERENCES curriculum_versions(id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (parent_version_id) REFERENCES tp(id),
    FOREIGN KEY (approved_by) REFERENCES users(id)
);

CREATE INDEX idx_tp_cp_id ON tp(cp_id);
CREATE INDEX idx_tp_subject_id ON tp(subject_id);
CREATE INDEX idx_tp_phase_id ON tp(phase_id);
CREATE INDEX idx_tp_element_id ON tp(element_id);
CREATE INDEX idx_tp_subelement_id ON tp(subelement_id);
CREATE INDEX idx_tp_curriculum_version_id ON tp(curriculum_version_id);
CREATE INDEX idx_tp_user_id ON tp(user_id);
CREATE INDEX idx_tp_status ON tp(status);
CREATE INDEX idx_tp_version_no ON tp(version_no);
CREATE INDEX idx_tp_is_current_version ON tp(is_current_version);
CREATE INDEX idx_tp_parent_version_id ON tp(parent_version_id);
CREATE INDEX idx_tp_ai_generation_id ON tp(ai_generation_id);
```

## Traceability Use Cases

### Historical Analysis

**Query**: Analyze TP quality by AI model version

```sql
SELECT 
    model_version,
    AVG(ai_confidence_score) as avg_confidence,
    COUNT(*) as tp_count
FROM tp
WHERE status = 'APPROVED'
GROUP BY model_version;
```

### Curriculum Change Impact

**Query**: Find all TPs affected by a CP change

```sql
SELECT tp.id, tp.version_no, tp.status
FROM tp
JOIN cp ON tp.cp_id = cp.id
WHERE cp.id = :cp_id
  AND tp.status = 'APPROVED';
```

### Audit Trail

**Query**: Trace TP generation history

```sql
SELECT 
    tp.id,
    tp.version_no,
    tp.status,
    tp.ai_generation_id,
    tp.model_version,
    tp.prompt_version,
    tp.created_at,
    tp.approved_at,
    u.name as teacher_name
FROM tp
JOIN users u ON tp.user_id = u.id
WHERE tp.cp_id = :cp_id
ORDER BY tp.version_no;
```

---

# SECTION 9 — Human Review Architecture

## Review Requirement

**Rule**: Teacher must review AI-generated TP before approval.

**Implications**:
- AI output cannot skip review step
- System enforces review workflow
- Teacher can approve or request changes
- Review is mandatory for all AI-generated TP

## Review Process

### Step 1: AI Generation

**Action**: AI generates TP Set in DRAFT status

**System Behavior**:
- TP Set created with status = DRAFT
- TP Set not visible to students
- TP Set cannot be used for downstream generation
- Notification sent to teacher

### Step 2: Teacher Review

**Action**: Teacher reviews TP Set content

**Review Checklist**:
- Learning objectives align with CP
- Time allocation is realistic
- Prerequisites are logical
- TP sequencing is appropriate
- Content is appropriate for grade level
- Content matches teaching style preferences

**Teacher Actions**:
- Approve entire TP Set
- Approve individual TPs
- Request changes (send back to DRAFT)
- Edit TP content directly
- Regenerate specific TP

### Step 3: Approval

**Action**: Teacher approves TP Set

**System Behavior**:
- TP Set status changes to APPROVED
- Version number assigned
- TP Set becomes official
- TP Set available for downstream generation
- Approval timestamp recorded

**Approval Rules**:
- Only teacher who generated (or has permission) can approve
- All TPs in TP Set must be reviewed
- Teacher must explicitly confirm approval

## Review UI Requirements

### TP Set Review Screen

**Components**:
1. Curriculum context display (Subject → Phase → Element → Subelement → CP)
2. TP Set overview (total TPs, total hours, estimated weeks)
3. Individual TP cards with:
   - Title and sequence number
   - Learning objectives
   - Time allocation
   - Prerequisites
4. AI confidence score display
5. Edit buttons for each TP
6. Regenerate buttons for each TP
7. Approve/Reject buttons for TP Set

### TP Detail Review Screen

**Components**:
1. Full TP content display
2. Side-by-side comparison with CP objectives
3. Edit mode for manual modifications
4. Regenerate option
5. Save/Cancel buttons
6. Approve button

## Review Quality Metrics

### Review Completion Rate

**Metric**: Percentage of AI-generated TP Sets that complete review

**Calculation**:
```
Review Completion Rate = (Approved TP Sets / Generated TP Sets) × 100
```

**Target**: > 90%

### Approval Rate

**Metric**: Percentage of reviewed TP Sets that are approved

**Calculation**:
```
Approval Rate = (Approved TP Sets / Reviewed TP Sets) × 100
```

**Target**: > 80%

### Edit Rate

**Metric**: Percentage of approved TP Sets that required teacher edits

**Calculation**:
```
Edit Rate = (Edited TP Sets / Approved TP Sets) × 100
```

**Target**: < 50% (indicates AI quality)

---

# SECTION 10 — TP Generation Workflow

## Complete Workflow Diagram

```
┌─────────────────────────────────────────────────────────┐
│ 1. Teacher Navigates Curriculum Hierarchy               │
│    Subject → Phase → Element → Subelement → CP          │
└────────────────────┬────────────────────────────────────┘
                     │
                     ↓
┌─────────────────────────────────────────────────────────┐
│ 2. Teacher Selects CP and Provides Context              │
│    - Class info (grade, subject, academic year)          │
│    - Teaching schedule (hours/week, weeks/semester)      │
│    - School context (type, level, region)                │
│    - Preferences (focus areas, teaching style)            │
└────────────────────┬────────────────────────────────────┘
                     │
                     ↓
┌─────────────────────────────────────────────────────────┐
│ 3. System Sends Generation Request to AI Agent           │
│    Input: Full curriculum hierarchy + context             │
└────────────────────┬────────────────────────────────────┘
                     │
                     ↓
┌─────────────────────────────────────────────────────────┐
│ 4. AI Agent Generates TP Set                            │
│    Output: Multiple TPs (3–10 typical)                  │
│    Status: DRAFT                                         │
└────────────────────┬────────────────────────────────────┘
                     │
                     ↓
┌─────────────────────────────────────────────────────────┐
│ 5. Teacher Reviews TP Set                              │
│    - Reviews each TP                                    │
│    - Edits if needed                                    │
│    - Regenerates if needed                              │
└────────────────────┬────────────────────────────────────┘
                     │
                     ↓
          ┌──────────┴──────────┐
          │                     │
          ↓                     ↓
    Request Changes        Approve
          │                     │
          ↓                     ↓
┌─────────────────┐   ┌─────────────────┐
│ Edit/Regenerate │   │ TP Set Status  │
│ (Back to DRAFT) │   │ → APPROVED      │
└─────────────────┘   │ Version #1      │
          │             └─────────────────┘
          │                     │
          └─────────────────────┘
                                │
                                ↓
                  ┌─────────────────────────┐
                  │ TP Set Available for     │
                  │ Downstream Generation    │
                  │ (ATP, Modul Ajar, etc.)  │
                  └─────────────────────────┘
```

## Workflow States and Transitions

### State Transition Table

| Current State | Event | Next State | Action |
|---------------|-------|------------|--------|
| None | Teacher selects CP | Ready for Generation | Display generation form |
| Ready for Generation | Teacher submits request | Generating | Call AI agent |
| Generating | AI completes | DRAFT | Display TP Set for review |
| DRAFT | Teacher edits | DRAFT | Save changes |
| DRAFT | Teacher regenerates | DRAFT | Create new TP |
| DRAFT | Teacher submits for review | REVIEW | Lock for review |
| REVIEW | Teacher requests changes | DRAFT | Unlock for editing |
| REVIEW | Teacher approves | APPROVED | Assign version, lock |
| APPROVED | Teacher regenerates | DRAFT (v2) | Create new version |
| APPROVED | System archives | ARCHIVED | Mark as read-only |

---

# SECTION 11 — Error Handling

## Generation Failures

### AI Service Unavailable

**Detection**: AI agent returns error or times out

**System Behavior**:
- Display error message to teacher
- Preserve teacher input context
- Allow retry
- Do not create partial TP Set

### Invalid Input

**Detection**: Validation fails on input parameters

**System Behavior**:
- Display specific validation error
- Highlight invalid fields
- Prevent generation until corrected

### Low Confidence Score

**Detection**: AI returns confidence score < threshold (e.g., 0.70)

**System Behavior**:
- Display warning to teacher
- Show confidence score
- Allow teacher to proceed with caution
- Recommend manual review

## Review Failures

### Concurrent Edit Conflict

**Detection**: Two users attempt to edit same TP simultaneously

**System Behavior**:
- Detect conflict on save
- Display conflict resolution UI
- Allow user to choose which version to keep
- Preserve both versions if needed

### Approval Constraint Violation

**Detection**: Attempt to approve TP with unresolved issues

**System Behavior**:
- Block approval
- Display specific issues
- Require resolution before approval

---

# SECTION 12 — Performance Considerations

## Generation Performance

### Target Response Times

| Operation | Target P95 |
|-----------|------------|
| TP Set Generation (3-5 TPs) | < 10 seconds |
| TP Set Generation (6-10 TPs) | < 20 seconds |
| Single TP Regeneration | < 5 seconds |

### Optimization Strategies

- Cache curriculum hierarchy data
- Pre-fetch CP and related entities
- Use streaming responses for large TP Sets
- Implement request queuing for concurrent generations

## Storage Performance

### TP Set Storage

**Estimated Size**:
- Single TP: ~5-10 KB (JSON)
- TP Set (5 TPs): ~25-50 KB
- With metadata: ~30-60 KB

**Storage Strategy**:
- Store learning objectives as JSONB
- Index frequently queried fields (cp_id, user_id, status)
- Partition by academic_year if needed

---

# SECTION 13 — Security Considerations

## Access Control

### Generation Permissions

**Required Permissions**:
- `curriculum:cp:view` - View CP
- `curriculum:tp:generate` - Generate TP
- `curriculum:tp:edit` - Edit TP
- `curriculum:tp:approve` - Approve TP

**Role-Based Access**:
- Teacher: Generate, edit, approve own TP
- School Admin: Generate, edit, approve any TP in school
- System Admin: Full access

### Data Privacy

**Sensitive Data**:
- Teacher preferences
- School context
- Student-related context (if added later)

**Protection**:
- Encrypt at rest
- Encrypt in transit
- Access logging
- Audit trail

---

# SECTION 14 — Summary

## Key Deliverables

✓ **TP Generation Workflow**: Complete workflow from curriculum navigation to TP approval
✓ **Lifecycle Diagram**: Visual representation of TP states and transitions
✓ **Versioning Rules**: Clear rules for TP versioning and regeneration
✓ **Traceability Rules**: Comprehensive traceability requirements for curriculum and generation
✓ **Human Review Architecture**: Mandatory review process with quality metrics

## Design Principles Adhered To

- ✓ AI generates recommendations, teacher has final authority
- ✓ AI never creates official TP automatically
- ✓ TP Set concept (1 CP → multiple TP)
- ✓ Complete curriculum hierarchy in generation context
- ✓ Version control with history preservation
- ✓ Comprehensive traceability
- ✓ Mandatory human review
- ✓ Focus only on TP (no ATP discussion)

## Next Steps

1. Implement TP generation API endpoints
2. Develop TP review UI components
3. Implement version control logic
4. Add traceability field population
5. Create generation monitoring and metrics
6. Develop error handling and retry logic
7. Implement access control and permissions
8. Create teacher training materials

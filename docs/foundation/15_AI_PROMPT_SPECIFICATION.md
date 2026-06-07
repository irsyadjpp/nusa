# 15_AI_PROMPT_SPECIFICATION.md

## Foundation Document for NUSA Education Platform

**Version**: 1.0
**Date**: June 2026
**Status**: FOUNDATION DOCUMENT
**Alignment**: Validated against Foundation Architecture (00A, 00B, 00C, 01, 02, 03, 04, 05, 06, 07, 08, 09, 10, 11, 12, 13, 14)

**Purpose**: Define the official AI prompt specifications for NUSA MVP Wave 1, serving as the AI contract between application modules and AI agents. This document provides complete prompt templates, input/output schemas, validation rules, and quality criteria for all 6 AI agents.

---

# SECTION 1 — AI Agent Overview

## Data Model Mapping Note

Database structures and AI prompt payload structures may differ. Application-layer mapping is responsible for transforming between Database Format and Prompt Payload Format.

**Example Transformation:**

Database Format:
- time_allocation_hours (INTEGER)
- hours_per_week (INTEGER)

Prompt Payload Format:
- time_allocation:
  - total_hours (INTEGER)
  - hours_per_week (INTEGER)

This transformation is intentional and does not require schema redesign. The application layer handles all necessary mapping between database records and AI agent input/output schemas.

## Agent List

| Agent | Module | Purpose |
|-------|--------|---------|
| TP Agent | Curriculum Module | Generate Teaching Plans from Curriculum Plans |
| ATP Agent | Learning Planning Module | Generate Annual Teaching Plans from Teaching Plans |
| Modul Ajar Agent | Learning Planning Module | Generate Modul Ajar from Annual Teaching Plans |
| Assessment Agent | Assessment Module | Generate Assessments from Modul Ajar |
| Rubric Agent | Assessment Module | Generate Rubrics from Assessments |
| Narrative Report Agent | Reporting Module | Generate Narrative Reports from Evidence and Evaluations |

## Agent Architecture

```
Application Module → AI Orchestration Module → AI Agent → LLM Provider → AI Agent → AI Orchestration Module → Application Module
```

## Output Format

All AI agent outputs must be:
- Structured and machine-processable
- JSON format preferred
- Validatable against defined schemas
- Include confidence scores
- Include generation metadata

---

# SECTION 2 — TP Agent Specification

## Agent Information

**Agent Name**: TP Generator Agent
**Agent ID**: `agent_tp_generator`
**Version**: 1.0.0
**Module**: Curriculum Module

---

## Purpose

Generate Teaching Plans (TP) from national Curriculum Plans (CP) by analyzing learning objectives, competency standards, and time allocation guidelines to create structured teaching plans aligned with national curriculum standards.

---

## Inputs

### Input Schema

```json
{
  "curriculum_hierarchy": {
    "subject": {
      "id": "string",
      "code": "string",
      "name": "string",
      "name_en": "string"
    },
    "phase": {
      "id": "string",
      "code": "string",
      "name": "string",
      "name_en": "string",
      "grade_level_start": "integer",
      "grade_level_end": "integer"
    },
    "element": {
      "id": "string",
      "code": "string",
      "name": "string",
      "name_en": "string",
      "description": "string"
    },
    "subelement": {
      "id": "string",
      "code": "string",
      "name": "string",
      "name_en": "string",
      "description": "string"
    }
  },
  "cp": {
    "id": "string",
    "code": "string",
    "description": "string",
    "competency_code": "string",
    "learning_objectives": [
      {
        "id": "string",
        "code": "string",
        "description": "string",
        "competency_code": "string"
      }
    ],
    "competency_standards": [
      {
        "id": "string",
        "code": "string",
        "description": "string"
      }
    ],
    "time_allocation": {
      "total_hours": "integer",
      "hours_per_week": "integer"
    },
    "version": "string"
  },
  "class_info": {
    "grade": "integer",
    "subject": "string",
    "academic_year": "integer"
  },
  "teaching_schedule": {
    "hours_per_week": "integer",
    "weeks_per_semester": "integer"
  },
  "preferences": {
    "focus_areas": ["string"],
    "teaching_style": "string"
  }
}
```

---

## Output Schema

```json
{
  "tp_set": {
    "set_id": "string",
    "cp_id": "string",
    "tps": [
      {
        "id": "string",
        "learning_objectives": [
          {
            "id": "string",
            "cp_objective_id": "string",
            "description": "string",
            "time_allocation_hours": "integer"
          }
        ],
        "time_allocation": {
          "total_hours": "integer",
          "hours_per_week": "integer"
        },
        "prerequisites": [
          {
            "objective_id": "string",
            "required_for": ["string"]
          }
        ]
      }
    ]
  },
  "metadata": {
    "confidence_score": "number",
    "generated_at": "string",
    "agent_version": "string"
  }
}
```

---

## Prompt Template

```
You are an expert curriculum designer for the Indonesian education system. Your task is to generate a Teaching Plan (TP) from the provided Curriculum Plan (CP) with full curriculum hierarchy context.

## Curriculum Hierarchy Context
- Subject: {curriculum_hierarchy.subject.name} ({curriculum_hierarchy.subject.code})
- Phase: {curriculum_hierarchy.phase.name} ({curriculum_hierarchy.phase.code}) - Grade Levels {curriculum_hierarchy.phase.grade_level_start} to {curriculum_hierarchy.phase.grade_level_end}
- Element: {curriculum_hierarchy.element.name} ({curriculum_hierarchy.element.code})
- Subelement: {curriculum_hierarchy.subelement.name} ({curriculum_hierarchy.subelement.code})

## Context
- Subject: {curriculum_hierarchy.subject.name}
- Phase: {curriculum_hierarchy.phase.name}
- Grade Level: {class_info.grade}
- Academic Year: {class_info.academic_year}
- Teaching Schedule: {teaching_schedule.hours_per_week} hours per week for {teaching_schedule.weeks_per_semester} weeks
- Total Available Hours: {teaching_schedule.hours_per_week * teaching_schedule.weeks_per_semester}

## Curriculum Plan (CP)
The CP contains the following learning objectives and competency standards:

### CP Description
{cp.description}

### Learning Objectives
{foreach learning_objective in cp.learning_objectives}
- {learning_objective.code}: {learning_objective.description}
- Competency: {learning_objective.competency_code}
{/foreach}

### Competency Standards
{foreach standard in cp.competency_standards}
- {standard.code}: {standard.description}
{/foreach}

### Time Allocation Guidelines
- Total Hours: {cp.time_allocation.total_hours}
- Hours per Week: {cp.time_allocation.hours_per_week}

## Teacher Preferences
{if preferences.focus_areas}
Focus Areas: {join(preferences.focus_areas, ", ")}
{/if}
{if preferences.teaching_style}
Teaching Style: {preferences.teaching_style}
{/if}

## Instructions
Generate a Teaching Plan Set (TP Set) containing multiple Teaching Plans (TP) from the provided Curriculum Plan (CP) with the following requirements:

1. **TP Set Structure**: Generate multiple TPs from a single CP
   - Each TP should represent a logical grouping of learning objectives
   - TPs should be sequenced to form a coherent learning progression
   - The TP Set as a whole should cover all CP learning objectives
   - Each TP should be independently teachable while maintaining coherence with the set

2. **Learning Objectives**: Transform CP learning objectives into actionable teaching objectives across multiple TPs
   - Each objective must be clear, measurable, and achievable
   - Allocate appropriate time for each objective based on complexity
   - Ensure alignment with competency standards
   - Total time allocation across all TPs must not exceed available hours
   - Consider the curriculum hierarchy context (element and subelement) for appropriate depth and scope

3. **Time Allocation**: Distribute time across TPs and objectives
   - Consider complexity of each objective
   - Account for assessment and review time within each TP
   - Ensure balanced distribution across the semester
   - Each TP should have appropriate time allocation for its objectives

4. **Prerequisites**: Identify prerequisite relationships
   - Determine which objectives must be learned before others
   - Create logical learning sequences within and across TPs
   - Ensure foundational concepts are covered first in earlier TPs
   - Maintain clear progression from one TP to the next

## Output Format
Provide your response as a JSON object with the following structure:

```json
{
  "tp_set": {
    "set_id": "unique_set_id",
    "cp_id": "cp_id_from_input",
    "tps": [
      {
        "id": "unique_tp_id_1",
        "learning_objectives": [
          {
            "id": "unique_objective_id",
            "cp_objective_id": "cp_objective_id",
            "description": "actionable teaching objective",
            "time_allocation_hours": integer
          }
        ],
        "time_allocation": {
          "total_hours": integer,
          "hours_per_week": integer
        },
        "prerequisites": [
          {
            "objective_id": "objective_id",
            "required_for": ["objective_id_1", "objective_id_2"]
          }
        ]
      },
      {
        "id": "unique_tp_id_2",
        "learning_objectives": [
          {
            "id": "unique_objective_id",
            "cp_objective_id": "cp_objective_id",
            "description": "actionable teaching objective",
            "time_allocation_hours": integer
          }
        ],
        "time_allocation": {
          "total_hours": integer,
          "hours_per_week": integer
        },
        "prerequisites": [
          {
            "objective_id": "objective_id",
            "required_for": ["objective_id_1", "objective_id_2"]
          }
        ]
      }
    ]
  },
  "metadata": {
    "confidence_score": 0.95,
    "generated_at": "2026-06-05T10:00:00Z",
    "agent_version": "1.0"
  }
}
```

## Quality Criteria
- Learning objectives must be specific and measurable
- Time allocation must be realistic and balanced
- Prerequisites must be logical and educationally sound
- Total hours must not exceed available teaching hours
- All objectives must align with competency standards

Generate the Teaching Plan now.
```

---

## Validation Rules

### Schema Validation
- Output must be valid JSON
- All required fields must be present
- All arrays must have at least one item
- All IDs must be unique within the response

### Data Validation
- `time_allocation.total_hours` must be ≤ available teaching hours
- `time_allocation.hours_per_week` must match teaching schedule
- `learning_objectives[].time_allocation_hours` must be > 0
- `prerequisites[].objective_id` must reference valid learning objective IDs
- `prerequisites[].required_for` must reference valid learning objective IDs

### Educational Validation
- Learning objectives must be measurable (use action verbs)
- Learning objectives must align with CP learning objectives
- Time allocation must be proportional to objective complexity
- Prerequisites must be logically ordered

---

## Failure Handling

### Failure Scenarios
- **Invalid Input**: CP data missing or malformed
- **Time Allocation Failure**: Cannot fit objectives within available hours
- **Ambiguous Objectives**: CP objectives unclear or incomplete
- **LLM Service Unavailable**: External LLM provider down
- **Generation Timeout**: LLM response exceeds time limit

### Failure Responses

```json
{
  "success": false,
  "error": {
    "code": "GENERATION_FAILED",
    "message": "Failed to generate Teaching Plan",
    "details": {
      "reason": "Time allocation exceeds available hours",
      "available_hours": 72,
      "required_hours": 80
    }
  },
  "metadata": {
    "attempted_at": "2026-06-03T12:00:00Z",
    "agent_version": "1.0.0"
  }
}
```

### Retry Strategy
- Retry on transient failures (LLM timeout, network issues)
- Do not retry on validation failures (invalid input, data issues)
- Maximum 3 retry attempts
- Exponential backoff: 1s, 2s, 4s

---

## Human Review Rules

### Review Checklist
- [ ] Learning objectives are clear and measurable
- [ ] Time allocation is realistic and balanced
- [ ] Prerequisites are logical and educationally sound
- [ ] All objectives align with competency standards
- [ ] Total hours do not exceed available time
- [ ] Language is appropriate for Indonesian education context

### Review Actions
- **Approve**: Teacher approves TP for ATP generation
- **Edit**: Teacher modifies TP before approval
- **Reject**: Teacher rejects TP and requests regeneration
- **Regenerate**: Teacher requests regeneration with different parameters

---

## Approval Rules

### Approval Requirements
- TP must be in DRAFT status
- All learning objectives must be complete
- Time allocation must be validated
- Human review must be completed
- Teacher must explicitly approve

### Approval Workflow
1. AI generates TP (status: DRAFT)
2. Teacher reviews TP
3. Teacher edits TP if needed
4. Teacher approves TP (status: APPROVED)
5. TP becomes available for ATP generation

### Rejection Workflow
1. Teacher rejects TP (status: REJECTED)
2. System records rejection reason
3. Teacher can request regeneration
4. New TP generated (status: DRAFT)

---

## Versioning Strategy

### Agent Versioning
- Format: `MAJOR.MINOR.PATCH` (e.g., 1.0.0)
- MAJOR: Breaking changes to prompt or output schema
- MINOR: Non-breaking changes to prompt or output schema
- PATCH: Bug fixes or minor improvements

### Prompt Versioning
- Each prompt version is stored in database
- Prompt versions are immutable
- Historical prompts are retained for audit trail

### Output Schema Versioning
- Output schema versions are tied to agent versions
- Schema changes require agent version increment
- Backward compatibility maintained within major version

---

## Logging Requirements

### Log Events
- **Generation Started**: Timestamp, input hash, user ID
- **Generation Completed**: Timestamp, output hash, confidence score
- **Generation Failed**: Timestamp, error code, error details
- **Human Review**: Timestamp, reviewer ID, review outcome
- **Approval**: Timestamp, approver ID, approval status

### Log Format

```json
{
  "event": "tp_generation_completed",
  "timestamp": "2026-06-03T12:00:00Z",
  "agent_id": "agent_tp_generator",
  "agent_version": "1.0.0",
  "user_id": "usr_1234567890",
  "input_hash": "sha256_hash",
  "output_hash": "sha256_hash",
  "confidence_score": 0.92,
  "duration_ms": 1500
}
```

### Log Retention
- All logs retained for 1 year
- Logs stored in audit_logs table
- Logs indexed by user_id, event_type, timestamp

---

## Quality Criteria

### Confidence Score Thresholds
- **High Confidence**: ≥ 0.80 (proceed with review)
- **Medium Confidence**: 0.50 - 0.79 (careful review required)
- **Low Confidence**: < 0.50 (manual creation recommended)

### Quality Metrics
- **Completeness**: All required fields present and populated
- **Accuracy**: Learning objectives align with CP
- **Coherence**: Prerequisites are logically ordered
- **Feasibility**: Time allocation is realistic
- **Clarity**: Language is clear and unambiguous

### Quality Monitoring
- Track confidence scores over time
- Monitor rejection rates
- Analyze common failure patterns
- Collect teacher feedback on quality

---

# SECTION 3 — ATP Agent Specification

## Agent Information

**Agent Name**: ATP Generator Agent
**Agent ID**: `agent_atp_generator`
**Version**: 1.0.0
**Module**: Learning Planning Module

---

## Purpose

Generate Annual Teaching Plans (ATP) from approved Teaching Plans (TP) by sequencing learning objectives across the academic year, optimizing time allocation, and identifying assessment points while respecting academic calendar constraints.

---

## Inputs

### Input Schema

```json
{
  "tp": {
    "id": "string",
    "learning_objectives": [
      {
        "id": "string",
        "description": "string",
        "time_allocation_hours": "integer"
      }
    ],
    "time_allocation": {
      "total_hours": "integer",
      "hours_per_week": "integer"
    },
    "prerequisites": [
      {
        "objective_id": "string",
        "required_for": ["string"]
      }
    ]
  },
  "academic_calendar": {
    "start_date": "string",
    "end_date": "string",
    "holidays": ["string"]
  },
  "class_schedule": {
    "days_per_week": "integer",
    "periods_per_day": "integer",
    "available_hours_per_week": "integer"
  }
}
```

---

## Output Schema

```json
{
  "atp": {
    "weekly_sequence": [
      {
        "week": "integer",
        "topics": [
          {
            "learning_objective_id": "string",
            "hours": "integer",
            "start_date": "string",
            "end_date": "string"
          }
        ]
      }
    ],
    "assessment_schedule": [
      {
        "week": "integer",
        "type": "FORMATIVE",
        "topics_covered": ["string"]
      }
    ]
  },
  "metadata": {
    "confidence_score": "number",
    "generated_at": "string",
    "agent_version": "string"
  }
}
```

---

## Prompt Template

```
You are an expert curriculum scheduler for the Indonesian education system. Your task is to generate an Annual Teaching Plan (ATP) from the provided Teaching Plan (TP).

## Context
- Academic Year: {tp.time_allocation.hours_per_week} hours per week
- Academic Calendar: {academic_calendar.start_date} to {academic_calendar.end_date}
- Class Schedule: {class_schedule.days_per_week} days per week, {class_schedule.periods_per_day} periods per day
- Available Hours per Week: {class_schedule.available_hours_per_week}
- Holidays: {join(academic_calendar.holidays, ", ")}

## Teaching Plan (TP)
The TP contains the following learning objectives and prerequisites:

### Learning Objectives
{foreach objective in tp.learning_objectives}
- {objective.id}: {objective.description} ({objective.time_allocation_hours} hours)
{/foreach}

### Prerequisites
{foreach prereq in tp.prerequisites}
- {prereq.objective_id} required for: {join(prereq.required_for, ", ")}
{/foreach}

### Total Time Allocation
- Total Hours: {tp.time_allocation.total_hours}
- Hours per Week: {tp.time_allocation.hours_per_week}

## Instructions
Generate an Annual Teaching Plan with the following requirements:

1. **Weekly Sequence**: Sequence learning objectives across the academic year
   - Respect prerequisite relationships
   - Balance workload across weeks
   - Account for holidays in scheduling
   - Ensure logical progression of topics

2. **Time Allocation**: Distribute hours across weeks
   - Match available hours per week
   - Account for assessment time
   - Ensure no week exceeds available hours
   - Provide buffer time for review and catch-up

3. **Assessment Schedule**: Identify assessment points
   - Schedule formative assessments every 3-4 weeks
   - Align assessments with topic completion
   - Ensure adequate preparation time
   - Balance assessment load across semester

## Output Format
Provide your response as a JSON object with the following structure:

```json
{
  "atp": {
    "weekly_sequence": [
      {
        "week": integer,
        "topics": [
          {
            "learning_objective_id": "objective_id",
            "hours": integer,
            "start_date": "YYYY-MM-DD",
            "end_date": "YYYY-MM-DD"
          }
        ]
      }
    ],
    "assessment_schedule": [
      {
        "week": integer,
        "type": "FORMATIVE",
        "topics_covered": ["objective_id_1", "objective_id_2"]
      }
    ]
  }
}
```

## Quality Criteria
- Weekly sequence must respect prerequisite relationships
- Time allocation must not exceed available hours per week
- Holidays must be respected in scheduling
- Assessments must be logically placed
- Workload must be balanced across weeks

Generate the Annual Teaching Plan now.
```

---

## Validation Rules

### Schema Validation
- Output must be valid JSON
- All required fields must be present
- All arrays must have at least one item
- All IDs must be unique within the response

### Data Validation
- `weekly_sequence[].week` must be sequential starting from 1
- `weekly_sequence[].topics[].hours` must sum to ≤ available hours per week
- `weekly_sequence[].topics[].start_date` must be before `end_date`
- `weekly_sequence[].topics[].learning_objective_id` must reference valid TP objective IDs
- `assessment_schedule[].topics_covered` must reference completed objectives

### Scheduling Validation
- Prerequisites must be satisfied before dependent objectives
- Holidays must not have scheduled topics
- Total weeks must fit within academic calendar
- Assessment weeks must have adequate preparation time

---

## Failure Handling

### Failure Scenarios
- **Invalid Input**: TP data missing or malformed
- **Scheduling Conflict**: Cannot fit objectives within calendar
- **Prerequisite Violation**: Cannot satisfy prerequisite relationships
- **LLM Service Unavailable**: External LLM provider down
- **Generation Timeout**: LLM response exceeds time limit

### Failure Responses

```json
{
  "success": false,
  "error": {
    "code": "GENERATION_FAILED",
    "message": "Failed to generate Annual Teaching Plan",
    "details": {
      "reason": "Cannot satisfy prerequisite relationships within calendar",
      "conflicting_objectives": ["obj_1", "obj_2"]
    }
  },
  "metadata": {
    "attempted_at": "2026-06-03T12:00:00Z",
    "agent_version": "1.0.0"
  }
}
```

---

## Human Review Rules

### Review Checklist
- [ ] Weekly sequence respects prerequisites
- [ ] Time allocation is balanced across weeks
- [ ] Holidays are respected
- [ ] Assessments are logically placed
- [ ] Workload is realistic
- [ ] Schedule fits within academic calendar

---

## Approval Rules

### Approval Requirements
- ATP must be in DRAFT status
- All weeks must be scheduled
- Time allocation must be validated
- Human review must be completed
- Teacher must explicitly approve

---

## Versioning Strategy

Same as TP Agent (Section 2).

---

## Logging Requirements

Same as TP Agent (Section 2), with event types:
- `atp_generation_started`
- `atp_generation_completed`
- `atp_generation_failed`

---

## Quality Criteria

### Confidence Score Thresholds
- **High Confidence**: ≥ 0.80 (proceed with review)
- **Medium Confidence**: 0.50 - 0.79 (careful review required)
- **Low Confidence**: < 0.50 (manual creation recommended)

### Quality Metrics
- **Completeness**: All weeks scheduled
- **Accuracy**: Prerequisites satisfied
- **Coherence**: Logical progression
- **Feasibility**: Realistic workload
- **Compliance**: Holidays respected

---

# SECTION 4 — Modul Ajar Agent Specification

## Agent Information

**Agent Name**: Modul Ajar Agent
**Agent ID**: `agent_modul_ajar_generator`
**Version**: 1.0.0
**Module**: Learning Planning Module

---

## Purpose

Generate Modul Ajar (lesson plans) from approved Annual Teaching Plans (ATP) by creating detailed lesson plans with learning activities, resource requirements, and assessment methods that are logically sequenced and matched to available resources.

---

## Inputs

### Input Schema

```json
{
  "atp": {
    "id": "string",
    "weekly_sequence": [
      {
        "week": "integer",
        "topics": [
          {
            "learning_objective_id": "string",
            "hours": "integer"
          }
        ]
      }
    ]
  },
  "week": "integer",
  "topic": {
    "learning_objective_id": "string",
    "title": "string"
  },
  "resources": {
    "textbooks": ["string"],
    "materials": ["string"]
  },
  "class_characteristics": {
    "student_count": "integer",
    "skill_level": "string"
  }
}
```

---

## Output Schema

```json
{
  "modul_ajar": {
    "learning_activities": [
      {
        "id": "string",
        "sequence": "integer",
        "activity_type": "string",
        "description": "string",
        "duration_minutes": "integer",
        "resources": ["string"]
      }
    ],
    "resource_requirements": [
      {
        "resource": "string",
        "quantity": "integer"
      }
    ],
    "assessment_methods": [
      {
        "type": "string",
        "description": "string"
      }
    ]
  },
  "metadata": {
    "confidence_score": "number",
    "generated_at": "string",
    "agent_version": "string"
  }
}
```

---

## Prompt Template

```
You are an expert lesson plan designer for the Indonesian education system. Your task is to generate a Modul Ajar (lesson plan) from the provided ATP topic.

## Context
- Week: {week}
- Topic: {topic.title}
- Learning Objective ID: {topic.learning_objective_id}
- Class Characteristics: {class_characteristics.student_count} students, {class_characteristics.skill_level} skill level
- Available Resources: {join(resources.textbooks, ", ")}, {join(resources.materials, ", ")}

## ATP Topic
This topic is scheduled for {atp.weekly_sequence[week-1].topics[0].hours} hours.

## Instructions
Generate a Modul Ajar with the following requirements:

1. **Learning Activities**: Create a sequence of learning activities
   - Start with introduction/hook activity
   - Include main learning activities
   - Include practice/application activities
   - End with summary/closure activity
   - Sequence activities logically
   - Total duration must match topic hours

2. **Resource Requirements**: Specify required resources
   - Use available resources where possible
   - Specify quantities for consumable resources
   - Ensure resources are appropriate for class size

3. **Assessment Methods**: Define assessment methods
   - Include formative assessment during lesson
   - Align assessment with learning objective
   - Provide clear assessment criteria

## Output Format
Provide your response as a JSON object with the following structure:

```json
{
  "modul_ajar": {
    "learning_activities": [
      {
        "id": "unique_id",
        "sequence": integer,
        "activity_type": "INTRODUCTION|MAIN|PRACTICE|CLOSURE",
        "description": "activity description",
        "duration_minutes": integer,
        "resources": ["resource_1", "resource_2"]
      }
    ],
    "resource_requirements": [
      {
        "resource": "resource_name",
        "quantity": integer
      }
    ],
    "assessment_methods": [
      {
        "type": "FORMATIVE|SUMMATIVE",
        "description": "assessment description"
      }
    ]
  }
}
```

## Quality Criteria
- Activities must be logically sequenced
- Total duration must match topic hours
- Resources must be available and appropriate
- Assessment methods must align with objective
- Activities must be appropriate for skill level

Generate the Modul Ajar now.
```

---

## Validation Rules

### Schema Validation
- Output must be valid JSON
- All required fields must be present
- All arrays must have at least one item
- All IDs must be unique within the response

### Data Validation
- `learning_activities[].sequence` must be sequential starting from 1
- `learning_activities[].duration_minutes` must sum to topic hours × 60
- `learning_activities[].activity_type` must be valid type
- `resource_requirements[].quantity` must be > 0
- `assessment_methods[].type` must be valid type

### Educational Validation
- Activities must be logically sequenced
- Activities must be appropriate for skill level
- Resources must be appropriate for class size
- Assessment methods must align with learning objective

---

## Failure Handling

### Failure Scenarios
- **Invalid Input**: ATP data missing or malformed
- **Duration Mismatch**: Cannot fit activities within topic hours
- **Resource Unavailable**: Required resources not available
- **LLM Service Unavailable**: External LLM provider down
- **Generation Timeout**: LLM response exceeds time limit

---

## Human Review Rules

### Review Checklist
- [ ] Activities are logically sequenced
- [ ] Duration matches topic hours
- [ ] Resources are available
- [ ] Assessment methods are appropriate
- [ ] Activities are appropriate for skill level
- [ ] Lesson is engaging and effective

---

## Approval Rules

### Approval Requirements
- Modul Ajar must be in DRAFT status
- All activities must be complete
- Duration must be validated
- Human review must be completed
- Teacher must explicitly approve

---

## Versioning Strategy

Same as TP Agent (Section 2).

---

## Logging Requirements

Same as TP Agent (Section 2), with event types:
- `modul_ajar_generation_started`
- `modul_ajar_generation_completed`
- `modul_ajar_generation_failed`

---

## Quality Criteria

### Confidence Score Thresholds
- **High Confidence**: ≥ 0.80 (proceed with review)
- **Medium Confidence**: 0.50 - 0.79 (careful review required)
- **Low Confidence**: < 0.50 (manual creation recommended)

### Quality Metrics
- **Completeness**: All activities defined
- **Accuracy**: Duration matches topic hours
- **Coherence**: Logical activity sequence
- **Feasibility**: Resources available
- **Engagement**: Activities are engaging

---

# SECTION 5 — Assessment Agent Specification

## Agent Information

**Agent Name**: Assessment Generator Agent
**Agent ID**: `agent_assessment_generator`
**Version**: 1.0.0
**Module**: Assessment Module

---

## Purpose

Generate assessments from approved Modul Ajar by creating assessment items aligned with learning objectives, providing varied question types, and ensuring appropriate difficulty levels and scoring guidelines.

---

## Inputs

### Input Schema

```json
{
  "modul_ajar": {
    "id": "string",
    "learning_activities": [
      {
        "id": "string",
        "description": "string",
        "learning_objective_id": "string"
      }
    ]
  },
  "assessment_type": "string",
  "question_count": "integer",
  "difficulty_level": "string",
  "time_allocation_minutes": "integer"
}
```

---

## Output Schema

```json
{
  "assessment": {
    "assessment_items": [
      {
        "id": "string",
        "type": "string",
        "question": "string",
        "options": ["string"],
        "correct_answer": "string",
        "points": "integer",
        "learning_objective_id": "string"
      }
    ],
    "answer_key": {
      "item_id": "answer"
    },
    "scoring_guidelines": {
      "total_points": "integer",
      "passing_score": "integer"
    }
  },
  "metadata": {
    "confidence_score": "number",
    "generated_at": "string",
    "agent_version": "string"
  }
}
```

---

## Prompt Template

```
You are an expert assessment designer for the Indonesian education system. Your task is to generate an assessment from the provided Modul Ajar.

## Context
- Assessment Type: {assessment_type}
- Question Count: {question_count}
- Difficulty Level: {difficulty_level}
- Time Allocation: {time_allocation_minutes} minutes

## Modul Ajar
The Modul Ajar contains the following learning activities:

{foreach activity in modul_ajar.learning_activities}
- {activity.description}
- Learning Objective: {activity.learning_objective_id}
{/foreach}

## Instructions
Generate an assessment with the following requirements:

1. **Assessment Items**: Create assessment items
   - Generate {question_count} questions
   - Use varied question types (multiple choice, essay, short answer)
   - Align questions with learning objectives
   - Match difficulty level to specified level
   - Ensure questions are clear and unambiguous

2. **Answer Key**: Provide correct answers
   - Provide correct answer for each item
   - For essay questions, provide rubric-based evaluation guidance

3. **Scoring Guidelines**: Define scoring
   - Assign points to each item
   - Calculate total points
   - Define passing score (typically 60-70% of total)

## Output Format
Provide your response as a JSON object with the following structure:

```json
{
  "assessment": {
    "assessment_items": [
      {
        "id": "unique_id",
        "type": "MULTIPLE_CHOICE|ESSAY|SHORT_ANSWER",
        "question": "question text",
        "options": ["option_1", "option_2", "option_3", "option_4"],
        "correct_answer": "correct_option",
        "points": integer,
        "learning_objective_id": "objective_id"
      }
    ],
    "answer_key": {
      "item_id": "answer"
    },
    "scoring_guidelines": {
      "total_points": integer,
      "passing_score": integer
    }
  }
}
```

## Quality Criteria
- Questions must align with learning objectives
- Questions must be clear and unambiguous
- Difficulty must match specified level
- Question types must be varied
- Scoring must be fair and consistent

Generate the assessment now.
```

---

## Validation Rules

### Schema Validation
- Output must be valid JSON
- All required fields must be present
- All arrays must have at least one item
- All IDs must be unique within the response

### Data Validation
- `assessment_items[].type` must be valid type
- `assessment_items[].points` must be > 0
- `assessment_items[].options` required for MULTIPLE_CHOICE
- `scoring_guidelines.total_points` must equal sum of item points
- `scoring_guidelines.passing_score` must be 60-70% of total_points

### Assessment Validation
- Questions must align with learning objectives
- Questions must be clear and unambiguous
- Difficulty must match specified level
- Answer key must be accurate
- Scoring must be fair

---

## Failure Handling

### Failure Scenarios
- **Invalid Input**: Modul Ajar data missing or malformed
- **Question Generation Failed**: Cannot generate valid questions
- **Difficulty Mismatch**: Cannot match difficulty level
- **LLM Service Unavailable**: External LLM provider down
- **Generation Timeout**: LLM response exceeds time limit

---

## Human Review Rules

### Review Checklist
- [ ] Questions align with learning objectives
- [ ] Questions are clear and unambiguous
- [ ] Difficulty matches specified level
- [ ] Question types are varied
- [ ] Answer key is accurate
- [ ] Scoring is fair and consistent

---

## Approval Rules

### Approval Requirements
- Assessment must be in DRAFT status
- All questions must be complete
- Answer key must be validated
- Human review must be completed
- Teacher must explicitly approve

---

## Versioning Strategy

Same as TP Agent (Section 2).

---

## Logging Requirements

Same as TP Agent (Section 2), with event types:
- `assessment_generation_started`
- `assessment_generation_completed`
- `assessment_generation_failed`

---

## Quality Criteria

### Confidence Score Thresholds
- **High Confidence**: ≥ 0.80 (proceed with review)
- **Medium Confidence**: 0.50 - 0.79 (careful review required)
- **Low Confidence**: < 0.50 (manual creation recommended)

### Quality Metrics
- **Completeness**: All questions defined
- **Accuracy**: Questions align with objectives
- **Clarity**: Questions are unambiguous
- **Fairness**: Scoring is fair
- **Variety**: Question types varied

---

# SECTION 6 — Rubric Agent Specification

## Agent Information

**Agent Name**: Rubric Generator Agent
**Agent ID**: `agent_rubric_generator`
**Version**: 1.0.0
**Module**: Assessment Module

---

## Purpose

Generate rubrics from approved assessments by defining clear performance criteria, creating performance level descriptors, and providing scoring guidelines that ensure fair and consistent evaluation.

---

## Inputs

### Input Schema

```json
{
  "assessment": {
    "id": "string",
    "assessment_items": [
      {
        "id": "string",
        "type": "string",
        "question": "string",
        "learning_objective_id": "string"
      }
    ]
  },
  "rubric_type": "string",
  "performance_levels": "integer",
  "criteria_categories": ["string"]
}
```

---

## Output Schema

```json
{
  "rubric": {
    "performance_criteria": [
      {
        "id": "string",
        "category": "string",
        "description": "string",
        "weight": "number"
      }
    ],
    "performance_levels": [
      {
        "level": "integer",
        "label": "string",
        "description": "string"
      }
    ],
    "scoring_guidelines": {
      "total_points": "integer",
      "passing_score": "integer"
    }
  },
  "metadata": {
    "confidence_score": "number",
    "generated_at": "string",
    "agent_version": "string"
  }
}
```

---

## Prompt Template

```
You are an expert rubric designer for the Indonesian education system. Your task is to generate a rubric from the provided assessment.

## Context
- Rubric Type: {rubric_type}
- Performance Levels: {performance_levels}
- Criteria Categories: {join(criteria_categories, ", ")}

## Assessment
The assessment contains the following items:

{foreach item in assessment.assessment_items}
- {item.type}: {item.question}
- Learning Objective: {item.learning_objective_id}
{/foreach}

## Instructions
Generate a rubric with the following requirements:

1. **Performance Criteria**: Define performance criteria
   - Use specified criteria categories
   - Provide clear descriptions for each criterion
   - Assign weights to each criterion (must sum to 1.0)
   - Ensure criteria are measurable and observable

2. **Performance Levels**: Define performance level descriptors
   - Create {performance_levels} performance levels
   - Provide clear labels for each level (e.g., Excellent, Proficient, Developing, Beginning)
   - Provide descriptive text for each level at each criterion
   - Ensure levels are clearly differentiated

3. **Scoring Guidelines**: Define scoring
   - Calculate total points based on weights
   - Define passing score (typically 60-70% of total)
   - Ensure scoring is fair and consistent

## Output Format
Provide your response as a JSON object with the following structure:

```json
{
  "rubric": {
    "performance_criteria": [
      {
        "id": "unique_id",
        "category": "category_name",
        "description": "criterion description",
        "weight": 0.4
      }
    ],
    "performance_levels": [
      {
        "level": 4,
        "label": "Excellent",
        "description": "level description"
      }
    ],
    "scoring_guidelines": {
      "total_points": integer,
      "passing_score": integer
    }
  }
}
```

## Quality Criteria
- Criteria must be measurable and observable
- Weights must sum to 1.0
- Performance levels must be clearly differentiated
- Descriptions must be clear and specific
- Scoring must be fair and consistent

Generate the rubric now.
```

---

## Validation Rules

### Schema Validation
- Output must be valid JSON
- All required fields must be present
- All arrays must have at least one item
- All IDs must be unique within the response

### Data Validation
- `performance_criteria[].weight` must be > 0
- Sum of `performance_criteria[].weight` must equal 1.0
- `performance_levels[].level` must be sequential starting from 1
- `performance_levels` count must match requested performance_levels
- `scoring_guidelines.passing_score` must be 60-70% of total_points

### Rubric Validation
- Criteria must be measurable and observable
- Performance levels must be clearly differentiated
- Descriptions must be clear and specific
- Weights must be appropriate for criteria importance

---

## Failure Handling

### Failure Scenarios
- **Invalid Input**: Assessment data missing or malformed
- **Criteria Generation Failed**: Cannot generate valid criteria
- **Weight Calculation Failed**: Cannot calculate valid weights
- **LLM Service Unavailable**: External LLM provider down
- **Generation Timeout**: LLM response exceeds time limit

---

## Human Review Rules

### Review Checklist
- [ ] Criteria are measurable and observable
- [ ] Weights sum to 1.0
- [ ] Performance levels are clearly differentiated
- [ ] Descriptions are clear and specific
- [ ] Scoring is fair and consistent
- [ ] Rubric is appropriate for assessment type

---

## Approval Rules

### Approval Requirements
- Rubric must be in DRAFT status
- All criteria must be complete
- Performance levels must be defined
- Human review must be completed
- Teacher must explicitly approve

---

## Versioning Strategy

Same as TP Agent (Section 2).

---

## Logging Requirements

Same as TP Agent (Section 2), with event types:
- `rubric_generation_started`
- `rubric_generation_completed`
- `rubric_generation_failed`

---

## Quality Criteria

### Confidence Score Thresholds
- **High Confidence**: ≥ 0.80 (proceed with review)
- **Medium Confidence**: 0.50 - 0.79 (careful review required)
- **Low Confidence**: < 0.50 (manual creation recommended)

### Quality Metrics
- **Completeness**: All criteria defined
- **Accuracy**: Criteria align with assessment
- **Clarity**: Descriptions are clear
- **Fairness**: Weights are appropriate
- **Differentiation**: Levels are distinct

---

# SECTION 7 — Narrative Report Agent Specification

## Agent Information

**Agent Name**: Narrative Report Agent
**Agent ID**: `agent_narrative_report_generator`
**Version**: 1.0.0
**Module**: Reporting Module

---

## Purpose

Generate narrative reports for parents from collected evidence and student evaluations by synthesizing evidence into clear, parent-friendly progress summaries, identifying strengths and areas for improvement, and providing actionable recommendations.

---

## Inputs

### Input Schema

```json
{
  "student_id": "string",
  "evidences": [
    {
      "id": "string",
      "evidence_type": "string",
      "evidence_data": {
        "description": "string"
      },
      "teacher_notes": "string"
    }
  ],
  "evaluations": [
    {
      "id": "string",
      "rubric_id": "string",
      "total_score": "integer",
      "max_score": "integer",
      "performance_level": "string",
      "performance_scores": [
        {
          "criteria_id": "string",
          "score": "integer",
          "notes": "string"
        }
      ]
    }
  ],
  "report_period": {
    "type": "string",
    "semester": "integer",
    "academic_year": "integer"
  },
  "language": "string"
}
```

---

## Output Schema

```json
{
  "narrative_report": {
    "content": {
      "progress_summary": "string",
      "strengths": ["string"],
      "areas_for_improvement": ["string"],
      "recommendations": ["string"]
    }
  },
  "metadata": {
    "confidence_score": "number",
    "generated_at": "string",
    "agent_version": "string"
  }
}
```

---

## Prompt Template

```
You are an expert educational reporter for the Indonesian education system. Your task is to generate a narrative report for parents from the provided evidence and evaluations.

## Context
- Student ID: {student_id}
- Report Period: {report_period.type} {report_period.semester}, {report_period.academic_year}
- Language: {language}

## Evidence
The following evidence has been collected:

{foreach evidence in evidences}
- Type: {evidence.evidence_type}
- Description: {evidence.evidence_data.description}
- Teacher Notes: {evidence.teacher_notes}
{/foreach}

## Evaluations
The following evaluations have been completed:

{foreach evaluation in evaluations}
- Total Score: {evaluation.total_score}/{evaluation.max_score}
- Performance Level: {evaluation.performance_level}
- Performance Scores:
  {foreach score in evaluation.performance_scores}
  - Criteria {score.criteria_id}: {score.score} - {score.notes}
  {/foreach}
{/foreach}

## Instructions
Generate a narrative report with the following requirements:

1. **Progress Summary**: Synthesize evidence into a clear summary
   - Describe overall progress during the report period
   - Highlight key achievements
   - Note any challenges or concerns
   - Use parent-friendly language
   - Maintain professional and supportive tone

2. **Strengths**: Identify student strengths
   - Based on evidence and evaluations
   - Provide specific examples where possible
   - Focus on positive achievements
   - Limit to 3-5 key strengths

3. **Areas for Improvement**: Identify areas needing improvement
   - Based on evidence and evaluations
   - Be constructive and supportive
   - Provide specific examples where possible
   - Limit to 3-5 key areas

4. **Recommendations**: Provide actionable recommendations
   - Suggest specific actions for improvement
   - Provide practical advice for parents
   - Align with identified areas for improvement
   - Limit to 3-5 key recommendations

## Output Format
Provide your response as a JSON object with the following structure:

```json
{
  "narrative_report": {
    "content": {
      "progress_summary": "comprehensive progress summary",
      "strengths": ["strength_1", "strength_2", "strength_3"],
      "areas_for_improvement": ["area_1", "area_2", "area_3"],
      "recommendations": ["recommendation_1", "recommendation_2", "recommendation_3"]
    }
  }
}
```

## Quality Criteria
- Progress summary must be comprehensive and clear
- Strengths must be evidence-based
- Areas for improvement must be constructive
- Recommendations must be actionable
- Language must be parent-friendly
- Tone must be professional and supportive

Generate the narrative report now.
```

---

## Validation Rules

### Schema Validation
- Output must be valid JSON
- All required fields must be present
- All arrays must have at least one item
- All arrays must have ≤ 10 items

### Data Validation
- `content.progress_summary` must be 100-2000 characters
- `content.strengths` must have 1-5 items
- `content.areas_for_improvement` must have 1-5 items
- `content.recommendations` must have 1-5 items
- All string items must be 10-500 characters

### Report Validation
- Progress summary must be comprehensive
- Strengths must be evidence-based
- Areas for improvement must be constructive
- Recommendations must be actionable
- Language must be appropriate for parents

---

## Failure Handling

### Failure Scenarios
- **Invalid Input**: Evidence or evaluation data missing or malformed
- **Synthesis Failed**: Cannot synthesize evidence into summary
- **Language Mismatch**: Cannot generate in specified language
- **LLM Service Unavailable**: External LLM provider down
- **Generation Timeout**: LLM response exceeds time limit

---

## Human Review Rules

### Review Checklist
- [ ] Progress summary is comprehensive and clear
- [ ] Strengths are evidence-based
- [ ] Areas for improvement are constructive
- [ ] Recommendations are actionable
- [ ] Language is parent-friendly
- [ ] Tone is professional and supportive
- [ ] Report is appropriate for student

---

## Approval Rules

### Approval Requirements
- Narrative report must be in DRAFT status
- All sections must be complete
- Human review must be completed
- Teacher must explicitly approve

---

## Versioning Strategy

Same as TP Agent (Section 2).

---

## Logging Requirements

Same as TP Agent (Section 2), with event types:
- `narrative_report_generation_started`
- `narrative_report_generation_completed`
- `narrative_report_generation_failed`

---

## Quality Criteria

### Confidence Score Thresholds
- **High Confidence**: ≥ 0.80 (proceed with review)
- **Medium Confidence**: 0.50 - 0.79 (careful review required)
- **Low Confidence**: < 0.50 (manual creation recommended)

### Quality Metrics
- **Completeness**: All sections complete
- **Accuracy**: Content aligns with evidence
- **Clarity**: Language is clear
- **Appropriateness**: Tone is appropriate
- **Actionability**: Recommendations are actionable

---

# SECTION 8 — Cross-Agent Considerations

## Agent Coordination

### Workflow Sequence
```
TP Agent → ATP Agent → Modul Ajar Agent → Assessment Agent → Rubric Agent → Narrative Report Agent
```

### Data Flow
- Each agent receives output from previous agent as input
- Output schemas are designed to be compatible with next agent's input
- AI Orchestration Module manages data flow between agents

### Error Propagation
- Agent failure stops workflow at that point
- Previous agent outputs are preserved
- Teacher can retry failed agent with modified parameters

---

## Common Patterns

### Input Validation
- All agents validate input schemas before processing
- Invalid inputs return structured error responses
- Validation failures do not trigger retries

### Confidence Scoring
- All agents calculate confidence scores
- Confidence scores are logged for monitoring
- Low confidence scores trigger careful review requirement

### Metadata
- All agents include generation metadata
- Metadata includes: confidence_score, generated_at, agent_version
- Metadata is stored with generated artifacts

### Human Review
- All AI-generated artifacts require human review
- Review checklists are agent-specific
- Approval workflow is consistent across agents

---

## Quality Assurance

### Automated Validation
- Schema validation for all outputs
- Data validation for all outputs
- Educational validation where applicable

### Human Review
- Required for all AI-generated artifacts
- Review checklists provided for each agent
- Approval workflow enforced by system

### Monitoring
- Track confidence scores over time
- Monitor rejection rates
- Analyze common failure patterns
- Collect teacher feedback

---

# SECTION 9 — Implementation Guidelines

## Prompt Engineering Best Practices

### Prompt Structure
- Clear context and instructions
- Explicit output format requirements
- Quality criteria defined upfront
- Examples provided where helpful

### Prompt Optimization
- Use clear, unambiguous language
- Avoid ambiguous instructions
- Provide specific constraints
- Test prompts with sample inputs

### Prompt Versioning
- Version prompts alongside agents
- Maintain prompt history
- Document prompt changes
- A/B test prompt variations

---

## Error Handling Guidelines

### Retry Strategy
- Retry on transient failures only
- Maximum 3 retry attempts
- Exponential backoff
- Do not retry validation failures

### Error Messages
- Provide clear error codes
- Include actionable error details
- Log all errors for analysis
- Display user-friendly messages

### Graceful Degradation
- Allow manual creation when AI fails
- Preserve previous work
- Provide clear error messages
- Enable workflow continuation

---

## Logging Guidelines

### Log Levels
- **INFO**: Normal operations (generation started, completed)
- **WARN**: Non-critical issues (low confidence, manual review required)
- **ERROR**: Failures (generation failed, validation failed)

### Log Content
- Timestamp for all events
- User ID for all user-initiated events
- Input/output hashes for traceability
- Confidence scores for monitoring
- Error details for debugging

### Log Retention
- All logs retained for 1 year
- Logs stored in audit_logs table
- Logs indexed for querying
- Logs exported for analysis

---

# SECTION 10 — AI Quality Evaluation Framework

## Purpose

Define measurable success criteria for AI-generated outputs to ensure quality, consistency, and educational validity. This framework provides objective metrics for evaluating AI agent performance and establishes clear thresholds for human intervention.

## Quality Metrics Overview

All AI-generated outputs are evaluated against the following quality dimensions:

### Core Metrics

- **Completeness Score**: Percentage of required elements present in the output (0-100)
- **Curriculum Alignment Score**: Alignment with national curriculum standards (0-100)
- **Structure Compliance Score**: Adherence to required output structure (0-100)
- **AI Confidence Score**: LLM's self-assessed confidence (0-100)
- **Teacher Acceptance Rate**: Percentage of outputs approved without major revision (0-100)
- **Revision Rate**: Percentage of outputs requiring significant revision (0-100)
- **Regeneration Rate**: Percentage of outputs requiring complete regeneration (0-100)

### Calculation Method

All scores are calculated as follows:
- **Automated Metrics**: Computed by validation rules and schema checks
- **Human Metrics**: Tracked through approval workflow and revision tracking
- **Aggregate Metrics**: Computed over rolling 30-day window

---

## TP Agent Quality Criteria

### Quality Metrics

| Metric | Calculation Method | Target Threshold |
|--------|-------------------|------------------|
| Completeness Score | (Required fields present / Total required fields) × 100 | ≥ 90% |
| Curriculum Alignment Score | (Learning objectives mapped / Total CP objectives) × 100 | ≥ 95% |
| Structure Compliance Score | (Valid JSON structure / Required structure) × 100 | 100% |
| AI Confidence Score | LLM-provided confidence score | ≥ 80% |
| Time Allocation Accuracy | (Allocated hours within ±10% of CP allocation) | 100% |

### Acceptance Threshold

**Minimum Acceptance Score**: 85%

An output is considered acceptable if:
- Completeness Score ≥ 90%
- Curriculum Alignment Score ≥ 95%
- Structure Compliance Score = 100%
- AI Confidence Score ≥ 80%
- Time Allocation Accuracy = 100%

### Human Review Requirement

**Mandatory Review**: Always required for TP Agent outputs

**Review Focus Areas**:
- Learning objective mapping accuracy
- Time allocation reasonableness
- Prerequisite relationship validity
- Educational pedagogical soundness

### Escalation Rule

**Escalate to Administrator** if:
- Completeness Score < 70%
- Curriculum Alignment Score < 80%
- AI Confidence Score < 60%
- Teacher rejection rate > 30% in 30-day window

**Action**: Administrator reviews prompt template and LLM configuration

---

## ATP Agent Quality Criteria

### Quality Metrics

| Metric | Calculation Method | Target Threshold |
|--------|-------------------|------------------|
| Completeness Score | (Required fields present / Total required fields) × 100 | ≥ 90% |
| TP Alignment Score | (ATP topics aligned with TP objectives / Total TP objectives) × 100 | ≥ 95% |
| Calendar Compliance Score | (Dates within academic calendar / Total dates) × 100 | 100% |
| Structure Compliance Score | (Valid JSON structure / Required structure) × 100 | 100% |
| AI Confidence Score | LLM-provided confidence score | ≥ 80% |

### Acceptance Threshold

**Minimum Acceptance Score**: 85%

An output is considered acceptable if:
- Completeness Score ≥ 90%
- TP Alignment Score ≥ 95%
- Calendar Compliance Score = 100%
- Structure Compliance Score = 100%
- AI Confidence Score ≥ 80%

### Human Review Requirement

**Mandatory Review**: Always required for ATP Agent outputs

**Review Focus Areas**:
- Topic sequencing logic
- Time allocation per topic
- Assessment point placement
- Holiday and break accommodation

### Escalation Rule

**Escalate to Administrator** if:
- Completeness Score < 70%
- TP Alignment Score < 80%
- Calendar Compliance Score < 90%
- Teacher rejection rate > 30% in 30-day window

**Action**: Administrator reviews prompt template and calendar integration

---

## Modul Ajar Agent Quality Criteria

### Quality Metrics

| Metric | Calculation Method | Target Threshold |
|--------|-------------------|------------------|
| Completeness Score | (Required fields present / Total required fields) × 100 | ≥ 90% |
| ATP Alignment Score | (Modul Ajar aligned with ATP topic / ATP topic) × 100 | 100% |
| Activity Sequence Score | (Logical activity flow / Total activities) × 100 | ≥ 85% |
| Resource Match Score | (Resources available for activities / Total activities) × 100 | ≥ 90% |
| Structure Compliance Score | (Valid JSON structure / Required structure) × 100 | 100% |
| AI Confidence Score | LLM-provided confidence score | ≥ 80% |

### Acceptance Threshold

**Minimum Acceptance Score**: 85%

An output is considered acceptable if:
- Completeness Score ≥ 90%
- ATP Alignment Score = 100%
- Activity Sequence Score ≥ 85%
- Resource Match Score ≥ 90%
- Structure Compliance Score = 100%
- AI Confidence Score ≥ 80%

### Human Review Requirement

**Mandatory Review**: Always required for Modul Ajar Agent outputs

**Review Focus Areas**:
- Learning activity variety and engagement
- Time allocation per activity
- Resource availability and appropriateness
- Assessment method alignment

### Escalation Rule

**Escalate to Administrator** if:
- Completeness Score < 70%
- Activity Sequence Score < 70%
- Resource Match Score < 70%
- Teacher rejection rate > 30% in 30-day window

**Action**: Administrator reviews prompt template and resource database

---

## Assessment Agent Quality Criteria

### Quality Metrics

| Metric | Calculation Method | Target Threshold |
|--------|-------------------|------------------|
| Completeness Score | (Required fields present / Total required fields) × 100 | ≥ 90% |
| Objective Alignment Score | (Questions aligned with learning objectives / Total questions) × 100 | ≥ 95% |
| Question Variety Score | (Different question types / Total questions) × 100 | ≥ 70% |
| Difficulty Balance Score | (Appropriate difficulty distribution / Total questions) × 100 | ≥ 80% |
| Structure Compliance Score | (Valid JSON structure / Required structure) × 100 | 100% |
| AI Confidence Score | LLM-provided confidence score | ≥ 80% |

### Acceptance Threshold

**Minimum Acceptance Score**: 85%

An output is considered acceptable if:
- Completeness Score ≥ 90%
- Objective Alignment Score ≥ 95%
- Question Variety Score ≥ 70%
- Difficulty Balance Score ≥ 80%
- Structure Compliance Score = 100%
- AI Confidence Score ≥ 80%

### Human Review Requirement

**Mandatory Review**: Always required for Assessment Agent outputs

**Review Focus Areas**:
- Question clarity and unambiguous wording
- Answer key accuracy
- Scoring guidelines clarity
- Time allocation reasonableness

### Escalation Rule

**Escalate to Administrator** if:
- Completeness Score < 70%
- Objective Alignment Score < 80%
- Answer key accuracy < 90%
- Teacher rejection rate > 30% in 30-day window

**Action**: Administrator reviews prompt template and assessment design principles

---

## Rubric Agent Quality Criteria

### Quality Metrics

| Metric | Calculation Method | Target Threshold |
|--------|-------------------|------------------|
| Completeness Score | (Required fields present / Total required fields) × 100 | ≥ 90% |
| Assessment Alignment Score | (Rubric criteria aligned with assessment items / Total assessment items) × 100 | ≥ 95% |
| Performance Clarity Score | (Clear performance descriptors / Total descriptors) × 100 | ≥ 90% |
| Scoring Consistency Score | (Consistent scoring ranges / Total criteria) × 100 | 100% |
| Structure Compliance Score | (Valid JSON structure / Required structure) × 100 | 100% |
| AI Confidence Score | LLM-provided confidence score | ≥ 80% |

### Acceptance Threshold

**Minimum Acceptance Score**: 85%

An output is considered acceptable if:
- Completeness Score ≥ 90%
- Assessment Alignment Score ≥ 95%
- Performance Clarity Score ≥ 90%
- Scoring Consistency Score = 100%
- Structure Compliance Score = 100%
- AI Confidence Score ≥ 80%

### Human Review Requirement

**Mandatory Review**: Always required for Rubric Agent outputs

**Review Focus Areas**:
- Performance criterion clarity
- Performance level descriptor specificity
- Scoring guideline objectivity
- Fairness and bias assessment

### Escalation Rule

**Escalate to Administrator** if:
- Completeness Score < 70%
- Performance Clarity Score < 70%
- Scoring Consistency Score < 90%
- Teacher rejection rate > 30% in 30-day window

**Action**: Administrator reviews prompt template and rubric design standards

---

## Narrative Report Agent Quality Criteria

### Quality Metrics

| Metric | Calculation Method | Target Threshold |
|--------|-------------------|------------------|
| Completeness Score | (Required sections present / Total required sections) × 100 | ≥ 90% |
| Evidence Integration Score | (Evidence referenced in report / Total evidence) × 100 | ≥ 90% |
| Tone Appropriateness Score | (Parent-friendly language / Total words) × 100 | ≥ 95% |
| Recommendation Actionability Score | (Actionable recommendations / Total recommendations) × 100 | ≥ 85% |
| Structure Compliance Score | (Valid JSON structure / Required structure) × 100 | 100% |
| AI Confidence Score | LLM-provided confidence score | ≥ 80% |

### Acceptance Threshold

**Minimum Acceptance Score**: 85%

An output is considered acceptable if:
- Completeness Score ≥ 90%
- Evidence Integration Score ≥ 90%
- Tone Appropriateness Score ≥ 95%
- Recommendation Actionability Score ≥ 85%
- Structure Compliance Score = 100%
- AI Confidence Score ≥ 80%

### Human Review Requirement

**Mandatory Review**: Always required for Narrative Report Agent outputs

**Review Focus Areas**:
- Report accuracy and factual correctness
- Parent communication appropriateness
- Student progress representation fairness
- Recommendation specificity and helpfulness

### Escalation Rule

**Escalate to Administrator** if:
- Completeness Score < 70%
- Evidence Integration Score < 70%
- Tone Appropriateness Score < 80%
- Teacher rejection rate > 30% in 30-day window

**Action**: Administrator reviews prompt template and communication guidelines

---

## Aggregate Quality Metrics

### Teacher Acceptance Rate

**Definition**: Percentage of AI-generated outputs approved by teachers without requiring major revision.

**Calculation**: (Approved without major revision / Total generated) × 100

**Target Threshold**: ≥ 70%

**Monitoring**: Tracked per agent over rolling 30-day window

**Alert Threshold**: < 50% triggers administrator review

### Revision Rate

**Definition**: Percentage of AI-generated outputs requiring significant teacher revision.

**Calculation**: (Requiring significant revision / Total generated) × 100

**Target Threshold**: ≤ 30%

**Monitoring**: Tracked per agent over rolling 30-day window

**Alert Threshold**: > 50% triggers administrator review

### Regeneration Rate

**Definition**: Percentage of AI-generated outputs requiring complete regeneration.

**Calculation**: (Completely regenerated / Total generated) × 100

**Target Threshold**: ≤ 10%

**Monitoring**: Tracked per agent over rolling 30-day window

**Alert Threshold**: > 20% triggers administrator review

### Average AI Confidence Score

**Definition**: Average confidence score across all AI-generated outputs.

**Calculation**: Sum of all confidence scores / Total generated

**Target Threshold**: ≥ 80%

**Monitoring**: Tracked per agent over rolling 30-day window

**Alert Threshold**: < 70% triggers administrator review

---

## Quality Monitoring Dashboard

### Dashboard Metrics

The system must provide a quality monitoring dashboard displaying:

**Per-Agent Metrics**:
- Total generations (30-day window)
- Teacher acceptance rate
- Revision rate
- Regeneration rate
- Average AI confidence score
- Average completeness score
- Average alignment score

**System-Wide Metrics**:
- Overall teacher acceptance rate
- Overall revision rate
- Overall regeneration rate
- Average AI confidence score across all agents
- Total generations by module

**Trend Analysis**:
- 30-day trend lines for all metrics
- Week-over-week change indicators
- Anomaly detection alerts

### Alert Configuration

**Critical Alerts** (Immediate notification):
- Teacher acceptance rate < 50% for any agent
- AI confidence score < 60% for any agent
- System-wide regeneration rate > 20%

**Warning Alerts** (Daily summary):
- Teacher acceptance rate < 70% for any agent
- Revision rate > 40% for any agent
- AI confidence score < 75% for any agent

**Informational Alerts** (Weekly report):
- Trends in quality metrics
- Comparison with previous periods
- Recommendations for improvement

---

## Quality Improvement Process

### Continuous Improvement Cycle

1. **Monitor**: Track quality metrics via dashboard
2. **Analyze**: Identify patterns and root causes of quality issues
3. **Improve**: Update prompt templates, validation rules, or LLM configuration
4. **Validate**: Measure impact of improvements on quality metrics
5. **Iterate**: Continue cycle based on results

### Prompt Optimization

**Trigger**: When quality metrics fall below thresholds

**Process**:
1. Review failing outputs for common patterns
2. Identify prompt template weaknesses
3. Update prompt template with clearer instructions
4. A/B test new prompt template against current
5. Roll out improved template if metrics improve

### Validation Rule Enhancement

**Trigger**: When validation rules fail to catch quality issues

**Process**:
1. Analyze outputs that passed validation but required major revision
2. Identify missing validation checks
3. Add new validation rules
4. Test validation rules on historical data
5. Deploy updated validation rules

---

## Quality Governance

### Quality Standards Committee

**Composition**:
- AI Architect (Chair)
- Lead Backend Engineer
- Lead Frontend Engineer
- Education Specialist
- Quality Assurance Lead

**Responsibilities**:
- Review quality metrics monthly
- Approve prompt template changes
- Set quality thresholds
- Escalate critical quality issues
- Define quality improvement priorities

### Quality Review Schedule

**Weekly**:
- Review quality dashboard
- Address critical alerts
- Track improvement initiatives

**Monthly**:
- Comprehensive quality review
- Trend analysis
- Prompt template optimization review
- Threshold adjustment review

**Quarterly**:
- Quality framework review
- LLM provider evaluation
- Quality standards update
- Improvement roadmap planning

---

## Conclusion

This AI Quality Evaluation Framework defines measurable success criteria for all 6 AI agents, ensuring quality, consistency, and educational validity through objective metrics, clear thresholds, and systematic quality improvement processes.

---

# SECTION 11 — Prompt Lifecycle Management

## Purpose

Control the evolution of AI prompts through structured versioning, state management, and change tracking. This framework ensures prompt changes are auditable, traceable, and governed by clear approval processes.

## Prompt States

### State Definitions

| State | Description | Transitions Allowed |
|-------|-------------|---------------------|
| **DRAFT** | Prompt under development, not yet in production | DRAFT → ACTIVE, DRAFT → ARCHIVED |
| **ACTIVE** | Prompt currently in use for AI generation | ACTIVE → DEPRECATED, ACTIVE → ARCHIVED |
| **DEPRECATED** | Prompt no longer recommended but still supported | DEPRECATED → ARCHIVED |
| **ARCHIVED** | Prompt retired and no longer available | None (terminal state) |

### State Transition Rules

- **DRAFT → ACTIVE**: Requires Quality Standards Committee approval
- **ACTIVE → DEPRECATED**: Requires Quality Standards Committee approval and 30-day notice
- **DEPRECATED → ARCHIVED**: Requires Quality Standards Committee approval
- **DRAFT → ARCHIVED**: Direct archival without activation (abandoned drafts)
- **ACTIVE → ARCHIVED**: Direct archival (emergency deprecation)

---

## Version Format

### Semantic Versioning

**Format**: `Major.Minor.Patch`

**Examples**:
- `1.0.0` - Initial release
- `1.1.0` - Prompt improvement (backward compatible)
- `1.0.1` - Typo correction (backward compatible)
- `2.0.0` - Breaking changes (not backward compatible)

### Version Rules

#### Major Version (X.0.0)

**Trigger**: Breaking changes that require:
- Input schema changes
- Output schema changes
- Fundamental prompt structure changes
- LLM provider changes
- Significant pedagogical approach changes

**Impact**: Requires full regression testing, may require data migration

#### Minor Version (0.X.0)

**Trigger**: Prompt improvements that are backward compatible:
- Enhanced instructions
- Additional context
- Improved examples
- Clarified requirements
- Performance optimizations

**Impact**: Requires validation testing, no data migration needed

#### Patch Version (0.0.X)

**Trigger**: Minor corrections that are backward compatible:
- Typo fixes
- Formatting corrections
- Grammar corrections
- Minor clarifications
- Documentation updates

**Impact**: Requires basic validation, no testing needed

---

## Prompt Registry

### TP Agent Prompt

**Prompt ID**: `prompt_tp_generator`

**Current Version**: `1.0.0`

**Current Status**: `ACTIVE`

**Effective Date**: `2026-06-01`

**Change History**:

| Version | Status | Effective Date | Change Type | Description | Changed By |
|---------|--------|---------------|-------------|-------------|------------|
| 1.0.0 | ACTIVE | 2026-06-01 | Major | Initial release with complete TP generation prompt | AI Architect |

---

### ATP Agent Prompt

**Prompt ID**: `prompt_atp_generator`

**Current Version**: `1.0.0`

**Current Status**: `ACTIVE`

**Effective Date**: `2026-06-01`

**Change History**:

| Version | Status | Effective Date | Change Type | Description | Changed By |
|---------|--------|---------------|-------------|-------------|------------|
| 1.0.0 | ACTIVE | 2026-06-01 | Major | Initial release with complete ATP generation prompt | AI Architect |

---

### Modul Ajar Agent Prompt

**Prompt ID**: `prompt_modul_ajar_generator`

**Current Version**: `1.0.0`

**Current Status**: `ACTIVE`

**Effective Date**: `2026-06-01`

**Change History**:

| Version | Status | Effective Date | Change Type | Description | Changed By |
|---------|--------|---------------|-------------|-------------|------------|
| 1.0.0 | ACTIVE | 2026-06-01 | Major | Initial release with complete Modul Ajar generation prompt | AI Architect |

---

### Assessment Agent Prompt

**Prompt ID**: `prompt_assessment_generator`

**Current Version**: `1.0.0`

**Current Status**: `ACTIVE`

**Effective Date**: `2026-06-01`

**Change History**:

| Version | Status | Effective Date | Change Type | Description | Changed By |
|---------|--------|---------------|-------------|-------------|------------|
| 1.0.0 | ACTIVE | 2026-06-01 | Major | Initial release with complete Assessment generation prompt | AI Architect |

---

### Rubric Agent Prompt

**Prompt ID**: `prompt_rubric_generator`

**Current Version**: `1.0.0`

**Current Status**: `ACTIVE`

**Effective Date**: `2026-06-01`

**Change History**:

| Version | Status | Effective Date | Change Type | Description | Changed By |
|---------|--------|---------------|-------------|-------------|------------|
| 1.0.0 | ACTIVE | 2026-06-01 | Major | Initial release with complete Rubric generation prompt | AI Architect |

---

### Narrative Report Agent Prompt

**Prompt ID**: `prompt_narrative_report_generator`

**Current Version**: `1.0.0`

**Current Status**: `ACTIVE`

**Effective Date**: `2026-06-01`

**Change History**:

| Version | Status | Effective Date | Change Type | Description | Changed By |
|---------|--------|---------------|-------------|-------------|------------|
| 1.0.0 | ACTIVE | 2026-06-01 | Major | Initial release with complete Narrative Report generation prompt | AI Architect |

---

## Change Management Process

### Change Request

**Initiation**: Any team member can submit a prompt change request

**Change Request Form**:
- Prompt ID
- Current Version
- Proposed Version
- Change Type (Major/Minor/Patch)
- Change Description
- Justification
- Expected Impact
- Testing Requirements

### Change Review

**Review Body**: Quality Standards Committee

**Review Criteria**:
- Alignment with educational standards
- Technical feasibility
- Impact assessment
- Testing adequacy
- Rollback plan

**Review Timeline**:
- Minor/Patch changes: 1 week review
- Major changes: 2 weeks review

### Change Approval

**Approval Authority**:
- **Minor/Patch**: AI Architect (Chair)
- **Major**: Quality Standards Committee (unanimous)

**Approval Checklist**:
- Change request reviewed
- Impact assessment completed
- Testing requirements defined
- Rollback plan documented
- Stakeholders notified

### Change Deployment

**Deployment Process**:
1. Update prompt version in registry
2. Deploy new prompt to staging environment
3. Execute testing plan
4. Deploy to production (with approval)
5. Monitor quality metrics
6. Rollback if issues detected

**Deployment Timeline**:
- Minor/Patch changes: Immediate deployment after approval
- Major changes: Deploy during maintenance window

---

## Audit Trail

### Audit Requirements

All prompt changes must be auditable and traceable:

**Audit Record Includes**:
- Prompt ID and version
- Change timestamp
- Changed by (user ID)
- Change type (Major/Minor/Patch)
- Change description
- Approval timestamp
- Approver (user ID)
- Deployment timestamp
- Rollback timestamp (if applicable)

### Audit Storage

**Storage Location**: `audit_logs` table

**Retention Period**: 7 years

**Access Control**: Administrator and Quality Standards Committee only

### Audit Query

**Audit Report Queries**:
- All changes for a specific prompt
- All changes by a specific user
- All changes within a date range
- All changes of a specific type (Major/Minor/Patch)
- All deployments within a date range

---

## Rollback Procedures

### Rollback Triggers

**Automatic Rollback**:
- Quality metrics fall below critical thresholds
- Error rate exceeds 10%
- Teacher rejection rate exceeds 50%

**Manual Rollback**:
- Quality Standards Committee decision
- Critical bug discovered
- LLM provider issues

### Rollback Process

1. **Initiation**: Quality Standards Committee initiates rollback
2. **Version Selection**: Select previous stable version
3. **Deployment**: Deploy previous version to production
4. **Verification**: Verify quality metrics return to acceptable levels
5. **Documentation**: Document rollback reason and outcome

### Rollback Timeline

- **Critical Rollback**: Within 1 hour of trigger
- **Standard Rollback**: Within 4 hours of trigger
- **Scheduled Rollback**: During next maintenance window

---

## Prompt Testing

### Testing Requirements

**Minor/Patch Changes**:
- Schema validation
- Basic functionality test
- Sample output review

**Major Changes**:
- Full regression testing
- Integration testing
- Performance testing
- Quality metrics validation
- A/B testing against previous version

### Testing Environment

**Staging Environment**: Mirrors production configuration
- Test data set
- LLM provider integration
- Quality monitoring
- Performance monitoring

### Test Criteria

**Pass Criteria**:
- All schema validations pass
- Quality metrics meet thresholds
- Error rate < 1%
- Performance within SLA
- Teacher acceptance rate ≥ 70%

**Fail Criteria**:
- Any schema validation fails
- Quality metrics below thresholds
- Error rate ≥ 5%
- Performance outside SLA
- Teacher acceptance rate < 50%

---

## Prompt Governance

### Governance Responsibilities

**Quality Standards Committee**:
- Approve all prompt changes
- Set prompt quality standards
- Review prompt performance metrics
- Escalate prompt quality issues
- Define prompt improvement priorities

**AI Architect**:
- Maintain prompt registry
- Implement approved changes
- Monitor prompt performance
- Execute rollback procedures
- Document prompt changes

**Education Specialist**:
- Review prompt educational validity
- Validate pedagogical approach
- Assess teacher feedback
- Recommend prompt improvements
- Ensure alignment with curriculum standards

### Governance Meetings

**Weekly**: Prompt performance review
**Monthly**: Prompt change review
**Quarterly**: Prompt strategy review

---

## Conclusion

This Prompt Lifecycle Management framework provides structured governance for AI prompt evolution, ensuring all changes are auditable, traceable, and approved through clear processes. This framework becomes the official prompt governance model for NUSA MVP Wave 1.

---

# SECTION 12 — Prompt Snapshot Strategy

## Overview

The Prompt Snapshot Strategy ensures prompt reproducibility by preserving complete prompt context for every AI generation. Prompt versions alone are insufficient for historical audit trails and governance reviews.

## Module Ownership

**Module**: AI Orchestration Module

**Purpose**: Prompt reproducibility and audit trail preservation

---

## Prompt Snapshot Requirements

### Snapshot Storage

For every AI generation, the following must be stored:

| Field | Description |
|-------|-------------|
| `prompt_version` | Semantic version of the prompt used (e.g., v1.0.0) |
| `prompt_template_id` | Unique identifier of the prompt template |
| `prompt_snapshot_hash` | SHA-256 hash of the complete prompt snapshot |

### Snapshot Content

The prompt snapshot must include:
- Complete prompt template text
- All variable substitutions applied
- System prompt (if applicable)
- User prompt
- Any context or reference data included
- Temperature and other generation parameters

---

## Snapshot Hash Calculation

### Hash Algorithm

Use SHA-256 for prompt snapshot hashing:

```python
import hashlib

def calculate_prompt_snapshot_hash(prompt_snapshot):
    """
    Calculate SHA-256 hash of prompt snapshot.
    
    Args:
        prompt_snapshot: Complete prompt snapshot as string
    
    Returns:
        SHA-256 hash as hexadecimal string
    """
    return hashlib.sha256(prompt_snapshot.encode('utf-8')).hexdigest()
```

### Hash Storage

Store the hash in `ai_generation_log.prompt_snapshot_hash` field.

---

## Reproducibility Requirements

### Historical Generations

Historical generations must remain reproducible:
- Prompt snapshots must be preserved for audit trail retention period (7 years)
- Hash verification must be possible for any historical generation
- Prompt changes must not invalidate historical audit trails

### Reproduction Process

To reproduce a historical generation:

1. Retrieve prompt snapshot from storage
2. Verify snapshot hash matches recorded hash
3. Reconstruct prompt with original parameters
4. Re-execute generation with same LLM provider and model
5. Compare output hash with original output hash

---

## Storage Strategy

### Snapshot Storage Location

Prompt snapshots are stored in:
- `ai_generation_log` table: `prompt_snapshot_hash` field
- External storage (e.g., S3, object storage): Full prompt snapshot content
- Indexed by generation ID for efficient retrieval

### Storage Format

Prompt snapshots are stored as:
- JSON format for structured prompts
- Plain text for simple prompts
- Include metadata (timestamp, agent name, user ID)

---

## Governance Requirements

### Audit Trail

Prompt snapshots support:
- Future investigation of AI-generated content
- Governance reviews of prompt effectiveness
- Debugging of generation issues
- Compliance audits

### Change Impact Analysis

Prompt snapshots enable:
- Analysis of prompt changes on output quality
- Comparison of outputs across prompt versions
- Identification of prompt regression issues
- Validation of prompt improvement effectiveness

---

## Implementation Guidelines

### Snapshot Generation

```python
def create_prompt_snapshot(prompt_template, variables, parameters):
    """
    Create complete prompt snapshot.
    
    Args:
        prompt_template: Prompt template string
        variables: Dictionary of variable substitutions
        parameters: Generation parameters (temperature, etc.)
    
    Returns:
        Complete prompt snapshot as string
    """
    snapshot = {
        'template': prompt_template,
        'variables': variables,
        'parameters': parameters,
        'timestamp': datetime.utcnow().isoformat()
    }
    return json.dumps(snapshot, sort_keys=True)
```

### Hash Verification

```python
def verify_prompt_snapshot_hash(generation_id, stored_hash, snapshot):
    """
    Verify prompt snapshot hash matches stored hash.
    
    Args:
        generation_id: Generation record ID
        stored_hash: Hash stored in ai_generation_log
        snapshot: Retrieved prompt snapshot
    
    Returns:
        Boolean indicating hash match
    """
    calculated_hash = calculate_prompt_snapshot_hash(snapshot)
    return calculated_hash == stored_hash
```

---

## Conclusion

The Prompt Snapshot Strategy ensures complete prompt reproducibility for all AI generations, enabling historical audit trails, governance reviews, and debugging capabilities. Prompt snapshots preserve the complete context of every generation, ensuring that prompt changes do not invalidate historical audit trails.

---

# SECTION 13 — Conclusion

## AI Prompt Specification Summary

This AI Prompt Specification (15) provides the complete AI contract for NUSA MVP Wave 1:

### Agent Count
- **Total Agents**: 6
- **Curriculum Module**: 1 agent (TP Agent)
- **Learning Planning Module**: 2 agents (ATP Agent, Modul Ajar Agent)
- **Assessment Module**: 2 agents (Assessment Agent, Rubric Agent)
- **Reporting Module**: 1 agent (Narrative Report Agent)

### Specification Characteristics
- Complete prompt templates for all agents
- Structured input/output schemas (JSON)
- Comprehensive validation rules
- Failure handling strategies
- Human review requirements
- Approval workflows
- Versioning strategies
- Logging requirements
- Quality criteria

### Implementation Readiness
This AI prompt specification is:
- ✅ Complete with all MVP agents
- ✅ Aligned with frozen architecture decisions
- ✅ Ready for AI Orchestration Module implementation
- ✅ Ready for LLM provider integration
- ✅ Ready for quality monitoring
- ✅ Ready for human-in-the-loop governance

The AI prompt specification is officially approved for NUSA MVP Wave 1 implementation.

---

**Document Status**: FOUNDATION DOCUMENT - LOCKED

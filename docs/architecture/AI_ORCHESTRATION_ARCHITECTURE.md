# AI Orchestration Architecture - NUSA Education System

## Executive Summary

This document defines the complete AI orchestration layer for the NUSA Education Operating System, implementing Domain-Driven Design (DDD) principles with robust prompt engineering, validation, and approval workflows for curriculum planning, learning design, assessment, and reporting.

---

## Architecture Overview

### Core Principles

1. **Domain-Driven Design**: Each workflow is bounded by its domain context
2. **Prompt Versioning**: All prompts are versioned and tracked
3. **Human-in-the-Loop**: Critical checkpoints require human approval
4. **Hallucination Mitigation**: Multi-layer validation and fact-checking
5. **Audit Trail**: Complete traceability of all AI-generated content
6. **Rollback Strategy**: Ability to revert to previous versions

### System Components

```
┌─────────────────────────────────────────────────────────────┐
│                    AI Orchestration Layer                    │
├─────────────────────────────────────────────────────────────┤
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐       │
│  │  Prompt      │  │  Validation  │  │  Approval    │       │
│  │  Engine      │  │  Engine      │  │  Engine      │       │
│  └──────────────┘  └──────────────┘  └──────────────┘       │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐       │
│  │  Workflow    │  │  Audit       │  │  Rollback    │       │
│  │  Orchestrator│  │  Trail       │  │  Manager    │       │
│  └──────────────┘  └──────────────┘  └──────────────┘       │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                    Domain Workflows                          │
├─────────────────────────────────────────────────────────────┤
│  TP │ ATP │ Modul Ajar │ Assessment │ Rubric │ Narrative    │
└─────────────────────────────────────────────────────────────┘
```

---

## Workflow #1: TP (Teaching Plan) Generation Workflow

### Domain Context
**Teaching Plan Domain** - Core curriculum planning component that defines learning objectives, teaching strategies, and assessment methods for specific learning objectives within a curriculum plan.

### User Inputs

```typescript
interface TPGenerationInput {
  // Curriculum Context
  curriculumPlanId: string;
  phase: 'A' | 'B' | 'C' | 'D' | 'E' | 'F';
  grade: string;
  subject: string;
  
  // Learning Objectives
  learningObjectives: string[];
  essentialQuestions: string[];
  
  // Time Allocation
  durationWeeks: number;
  hoursPerWeek: number;
  
  // Student Context
  studentProfile: {
    averageAge: number;
    learningLevel: 'beginner' | 'intermediate' | 'advanced';
    specialNeeds: string[];
  };
  
  // Resource Constraints
  availableResources: string[];
  timeConstraints: string[];
  
  // Teacher Preferences
  teachingStyle: 'direct' | 'inquiry' | 'collaborative' | 'problem-based';
  assessmentPreference: 'formative' | 'summative' | 'both';
}
```

### AI Prompts

#### System Prompt (Version 1.0.0)

```
You are an expert curriculum designer specializing in the Indonesian Kurikulum Merdeka. 
Your task is to create comprehensive Teaching Plans (TP) that align with national standards 
while being adaptable to local contexts.

CORE PRINCIPLES:
1. Student-centered learning
2. Differentiated instruction
3. Formative assessment integration
4. Cultural relevance
5. 21st-century skills development

OUTPUT STRUCTURE:
- Learning objectives (aligned with CP)
- Teaching strategies with time allocation
- Learning activities (opening, core, closing)
- Assessment methods and criteria
- Resource requirements
- Differentiation strategies

CONSTRAINTS:
- Must align with Kurikulum Merdeka standards
- Include P5 (Projek Penguatan Profil Pelajar Pancasila) elements
- Consider local wisdom and cultural context
- Ensure inclusivity for diverse learners
```

#### User Prompt Template

```
Generate a Teaching Plan for the following context:

CURRICULUM CONTEXT:
- Phase: {phase}
- Grade: {grade}
- Subject: {subject}
- Duration: {durationWeeks} weeks ({hoursPerWeek} hours/week)

LEARNING OBJECTIVES:
{learningObjectives}

ESSENTIAL QUESTIONS:
{essentialQuestions}

STUDENT PROFILE:
- Age: {averageAge}
- Level: {learningLevel}
- Special Needs: {specialNeeds}

TEACHING PREFERENCES:
- Style: {teachingStyle}
- Assessment: {assessmentPreference}

AVAILABLE RESOURCES:
{availableResources}

Please generate a comprehensive TP that includes:
1. Detailed learning objectives with Bloom's taxonomy alignment
2. Week-by-week breakdown with time allocation
3. Teaching strategies and learning activities
4. Assessment methods with success criteria
5. Required resources and materials
6. Differentiation strategies for diverse learners
7. P5 integration opportunities
```

### Validation Rules

#### Structural Validation

```typescript
const structuralRules = [
  {
    rule: 'has_learning_objectives',
    check: (tp) => tp.learningObjectives?.length > 0,
    message: 'TP must have at least one learning objective'
  },
  {
    rule: 'objectives_aligned_with_cp',
    check: (tp, input) => {
      return tp.learningObjectives.every(obj => 
        input.learningObjectives.some(cpObj => 
          obj.includes(cpObj.substring(0, 20))
        )
      );
    },
    message: 'Learning objectives must align with curriculum plan'
  },
  {
    rule: 'has_weekly_breakdown',
    check: (tp) => tp.weeklyPlan?.length === input.durationWeeks,
    message: 'Weekly plan must match duration'
  },
  {
    rule: 'time_allocation_valid',
    check: (tp) => {
      const totalHours = tp.weeklyPlan.reduce((sum, week) => 
        sum + week.hours, 0
      );
      return totalHours <= input.durationWeeks * input.hoursPerWeek;
    },
    message: 'Time allocation must not exceed available hours'
  },
  {
    rule: 'includes_assessment',
    check: (tp) => tp.assessments?.length > 0,
    message: 'TP must include assessment methods'
  },
  {
    rule: 'has_p5_integration',
    check: (tp) => tp.p5Elements?.length > 0,
    message: 'TP must include P5 elements'
  }
];
```

#### Content Validation

```typescript
const contentValidationRules = [
  {
    rule: 'bloom_taxonomy_alignment',
    check: async (tp) => {
      const bloomLevels = ['remember', 'understand', 'apply', 'analyze', 'evaluate', 'create'];
      return tp.learningObjectives.some(obj => 
        bloomLevels.some(level => obj.toLowerCase().includes(level))
      );
    },
    message: 'Learning objectives should span Bloom\'s taxonomy levels'
  },
  {
    rule: 'cultural_relevance',
    check: async (tp) => {
      const indonesianContext = [
        'budaya', 'lokal', 'wisata', 'kearifan', 'tradisi'
      ];
      return tp.activities.some(act => 
        indonesianContext.some(ctx => act.description.toLowerCase().includes(ctx))
      );
    },
    message: 'Activities should include local cultural context'
  },
  {
    rule: 'differentiation_present',
    check: (tp) => tp.differentiationStrategies?.length > 0,
    message: 'Must include differentiation strategies'
  }
];
```

### Approval Workflow

```
┌─────────────────────────────────────────────────────────────┐
│                    TP Approval Workflow                      │
└─────────────────────────────────────────────────────────────┘

1. DRAFT STATUS
   ├─ AI generates TP
   ├─ Structural validation passes
   ├─ Content validation passes
   └─ Status: PENDING_REVIEW

2. TEACHER REVIEW (Level 1)
   ├─ Teacher reviews generated TP
   ├─ Can request regeneration with feedback
   ├─ Can make minor edits
   └─ Status: PENDING_APPROVAL or BACK_TO_GENERATION

3. PEER REVIEW (Level 2) - Optional
   ├─ Peer teacher reviews TP
   ├─ Checks for pedagogical soundness
   └─ Status: PENDING_FINAL_APPROVAL or BACK_TO_TEACHER

4. CURRICULUM COORDINATOR APPROVAL (Level 3)
   ├─ Coordinator reviews alignment with CP
   ├─ Checks Kurikulum Merdeka compliance
   └─ Status: APPROVED or REJECTED

5. FINAL APPROVAL
   ├─ TP marked as approved
   ├─ Version created (v1.0)
   └─ Ready for ATP generation
```

### Regeneration Workflow

```typescript
interface RegenerationRequest {
  tpId: string;
  feedback: {
    category: 'content' | 'structure' | 'alignment' | 'resources';
    specificIssues: string[];
    suggestions: string[];
  };
  regenerateFull: boolean;
  preserveSections: string[];
}

const regenerationStrategy = {
  full: async (request) => {
    // Complete regeneration with feedback incorporated
    const enhancedPrompt = basePrompt + `
PREVIOUS VERSION ISSUES:
${request.feedback.specificIssues.join('\n')}

SUGGESTED IMPROVEMENTS:
${request.feedback.suggestions.join('\n')}
`;
    return await generateTP(enhancedPrompt);
  },
  
  partial: async (request) => {
    // Regenerate only specified sections
    const preservedTP = await getTP(request.tpId);
    const regeneratedSections = await generateSections(
      request.feedback.category,
      preservedTP
    );
    return mergeTP(preservedTP, regeneratedSections, request.preserveSections);
  }
};
```

### Audit Trail

```typescript
interface TPAuditEntry {
  id: string;
  tpId: string;
  timestamp: ISO8601;
  eventType: 'generated' | 'validated' | 'approved' | 'rejected' | 'edited' | 'regenerated';
  actor: {
    type: 'ai' | 'user' | 'system';
    id: string;
    name: string;
  };
  changes: {
    before?: any;
    after?: any;
    diff?: string;
  };
  metadata: {
    promptVersion: string;
    modelVersion: string;
    validationResults: ValidationResult[];
    approvalLevel: number;
  };
}
```

---

## Workflow #2: ATP (Annual Teaching Plan) Generation Workflow

### Domain Context
**Annual Teaching Plan Domain** - Time-bound planning component that distributes teaching plans across an academic year, considering holidays, assessments, and school events.

### User Inputs

```typescript
interface ATPGenerationInput {
  // Teaching Plans to Schedule
  teachingPlanIds: string[];
  
  // Academic Year Context
  academicYear: string;
  semester: 1 | 2;
  startDate: Date;
  endDate: Date;
  
  // School Calendar
  holidays: Holiday[];
  schoolEvents: SchoolEvent[];
  assessmentPeriods: AssessmentPeriod[];
  
  // Time Constraints
  teachingDaysPerWeek: number;
  hoursPerDay: number;
  
  // Scheduling Preferences
  distributionStrategy: 'balanced' | 'front-loaded' | 'back-loaded' | 'custom';
  bufferDays: number;
  reviewDays: number;
}
```

### AI Prompts

#### System Prompt (Version 1.0.0)

```
You are an expert academic scheduler specializing in Indonesian school calendars.
Your task is to create Annual Teaching Plans (ATP) that optimally distribute teaching 
content across an academic year while accounting for holidays, assessments, and school events.

CORE PRINCIPLES:
1. Balanced workload distribution
2. Assessment preparation time
3. Holiday and event accommodation
4. Buffer time for contingencies
5. Review and remediation periods

OUTPUT STRUCTURE:
- Week-by-week schedule
- Teaching plan assignments
- Assessment windows
- Holiday and event markers
- Buffer and review periods
- Total hours allocation

CONSTRAINTS:
- Must respect school calendar
- Ensure adequate preparation time for assessments
- Include buffer days for unexpected events
- Balance workload across weeks
- Align with semester structure
```

#### User Prompt Template

```
Generate an Annual Teaching Plan for the following context:

ACADEMIC YEAR:
- Year: {academicYear}
- Semester: {semester}
- Start Date: {startDate}
- End Date: {endDate}

TEACHING PLANS TO SCHEDULE:
{teachingPlansSummary}

SCHOOL CALENDAR:
- Holidays: {holidays}
- School Events: {schoolEvents}
- Assessment Periods: {assessmentPeriods}

TIME CONSTRAINTS:
- Teaching Days/Week: {teachingDaysPerWeek}
- Hours/Day: {hoursPerDay}

SCHEDULING PREFERENCES:
- Strategy: {distributionStrategy}
- Buffer Days: {bufferDays}
- Review Days: {reviewDays}

Please generate an ATP that includes:
1. Week-by-week breakdown with dates
2. Teaching plan assignments with time allocation
3. Assessment windows with preparation time
4. Holiday and event accommodations
5. Buffer periods for contingencies
6. Review and remediation periods
7. Total hours tracking per teaching plan
```

### Validation Rules

```typescript
const atpValidationRules = [
  {
    rule: 'all_tps_scheduled',
    check: (atp, input) => {
      const scheduledTPs = new Set(atp.weeks.flatMap(w => w.teachingPlans.map(tp => tp.id)));
      return input.teachingPlanIds.every(id => scheduledTPs.has(id));
    },
    message: 'All teaching plans must be scheduled'
  },
  {
    rule: 'no_holiday_conflicts',
    check: (atp, input) => {
      return atp.weeks.every(week => {
        const weekHolidays = input.holidays.filter(h => 
          isDateInRange(h.date, week.startDate, week.endDate)
        );
        return weekHolidays.length === 0 || week.isHolidayWeek;
      });
    },
    message: 'Teaching cannot be scheduled on holidays'
  },
  {
    rule: 'assessment_preparation_time',
    check: (atp) => {
      return atp.assessmentWindows.every(window => {
        const preparationWeeks = atp.weeks.filter(w =>
          isDateInRange(w.startDate, window.preparationStart, window.assessmentStart)
        );
        return preparationWeeks.length >= 2; // At least 2 weeks preparation
      });
    },
    message: 'Assessments must have adequate preparation time'
  },
  {
    rule: 'balanced_workload',
    check: (atp) => {
      const weeklyHours = atp.weeks.map(w => w.totalHours);
      const avgHours = weeklyHours.reduce((a, b) => a + b, 0) / weeklyHours.length;
      const variance = weeklyHours.reduce((sum, h) => sum + Math.pow(h - avgHours, 2), 0) / weeklyHours.length;
      return variance < Math.pow(avgHours * 0.3, 2); // Within 30% variance
    },
    message: 'Workload should be balanced across weeks'
  }
];
```

### Approval Workflow

```
┌─────────────────────────────────────────────────────────────┐
│                   ATP Approval Workflow                      │
└─────────────────────────────────────────────────────────────┘

1. DRAFT STATUS
   ├─ AI generates ATP
   ├─ Validation passes
   └─ Status: PENDING_REVIEW

2. TEACHER REVIEW (Level 1)
   ├─ Teacher reviews schedule
   ├─ Can adjust time allocations
   ├─ Can request regeneration
   └─ Status: PENDING_COORDINATOR

3. CURRICULUM COORDINATOR REVIEW (Level 2)
   ├─ Checks alignment with school calendar
   ├─ Validates assessment windows
   ├─ Ensures workload balance
   └─ Status: PENDING_PRINCIPAL or BACK_TO_TEACHER

4. PRINCIPAL APPROVAL (Level 3)
   ├─ Final approval for academic year
   ├─ Locks ATP for the semester
   └─ Status: APPROVED or REJECTED

5. FINAL APPROVAL
   ├─ ATP marked as approved
   ├─ Version created (v1.0)
   └─ Ready for Modul Ajar generation
```

---

## Workflow #3: Modul Ajar Generation Workflow

### Domain Context
**Modul Ajar Domain** - Learning material creation domain that generates comprehensive teaching modules including lesson plans, activities, worksheets, and assessment materials.

### User Inputs

```typescript
interface ModulAjarGenerationInput {
  // Teaching Plan Reference
  teachingPlanId: string;
  atpWeekId: string;
  
  // Module Context
  moduleType: 'opening' | 'core' | 'closing' | 'assessment' | 'enrichment';
  duration: number; // in minutes
  
  // Learning Focus
  specificLearningObjectives: string[];
  keyConcepts: string[];
  
  // Student Context
  classSize: number;
  groupSize: number;
  availableEquipment: string[];
  
  // Teaching Approach
  pedagogicalApproach: 'direct' | 'inquiry' | 'collaborative' | 'problem-based' | 'project-based';
  differentiationLevel: 'none' | 'basic' | 'advanced';
  
  // Resource Preferences
  includeWorksheet: boolean;
  includeAssessment: boolean;
  includeHomework: boolean;
  language: 'indonesian' | 'bilingual' | 'local';
}
```

### AI Prompts

#### System Prompt (Version 1.2.0)

```
You are an expert instructional designer specializing in the Indonesian Kurikulum Merdeka.
Your task is to create comprehensive Modul Ajar (Teaching Modules) that are engaging, 
inclusive, and aligned with 21st-century learning principles.

CORE PRINCIPLES:
1. Active learning engagement
2. Student-centered activities
3. Differentiated instruction
4. Formative assessment integration
5. Cultural and local relevance

OUTPUT STRUCTURE:
- Learning objectives (specific and measurable)
- Prerequisite knowledge
- Required materials and resources
- Lesson phases (opening, core, closing)
- Detailed activity instructions
- Student worksheets (if requested)
- Assessment rubrics (if requested)
- Differentiation strategies
- Homework assignments (if requested)

CONSTRAINTS:
- Must align with Kurikulum Merdeka standards
- Include character education (P5) elements
- Use age-appropriate language
- Ensure inclusivity for diverse learners
- Incorporate local wisdom and culture
- Provide clear time allocations
```

#### User Prompt Template

```
Generate a Modul Ajar for the following context:

TEACHING PLAN REFERENCE:
- TP ID: {teachingPlanId}
- ATP Week: {atpWeekId}
- Module Type: {moduleType}
- Duration: {duration} minutes

LEARNING FOCUS:
- Specific Objectives: {specificLearningObjectives}
- Key Concepts: {keyConcepts}

STUDENT CONTEXT:
- Class Size: {classSize}
- Group Size: {groupSize}
- Available Equipment: {availableEquipment}

TEACHING APPROACH:
- Pedagogy: {pedagogicalApproach}
- Differentiation: {differentiationLevel}

RESOURCE PREFERENCES:
- Worksheet: {includeWorksheet}
- Assessment: {includeAssessment}
- Homework: {includeHomework}
- Language: {language}

Please generate a comprehensive Modul Ajar that includes:
1. Clear learning objectives with success criteria
2. Prerequisite knowledge check
3. Detailed materials and resources list
4. Lesson phases with time allocation:
   - Opening activity (10-15% of time)
   - Core learning activities (70-75% of time)
   - Closing activity (10-15% of time)
5. Step-by-step activity instructions
6. Student worksheets (if requested)
7. Assessment rubrics (if requested)
8. Differentiation strategies for diverse learners
9. Homework assignments (if requested)
10. P5 character education integration
11. Local cultural references where appropriate
```

### Validation Rules

```typescript
const modulAjarValidationRules = [
  {
    rule: 'time_allocation_valid',
    check: (modul, input) => {
      const totalTime = modul.phases.reduce((sum, phase) => sum + phase.duration, 0);
      return totalTime === input.duration;
    },
    message: 'Phase durations must sum to total duration'
  },
  {
    rule: 'has_opening_core_closing',
    check: (modul) => {
      const phases = modul.phases.map(p => p.type);
      return phases.includes('opening') && 
             phases.includes('core') && 
             phases.includes('closing');
    },
    message: 'Module must have opening, core, and closing phases'
  },
  {
    rule: 'objectives_measurable',
    check: (modul) => {
      const measurableVerbs = ['menganalisis', 'mengidentifikasi', 'menerapkan', 'mengevaluasi', 'mencipta'];
      return modul.learningObjectives.every(obj =>
        measurableVerbs.some(verb => obj.toLowerCase().includes(verb))
      );
    },
    message: 'Learning objectives must use measurable verbs'
  },
  {
    rule: 'includes_differentiation',
    check: (modul, input) => {
      if (input.differentiationLevel !== 'none') {
        return modul.differentiationStrategies?.length > 0;
      }
      return true;
    },
    message: 'Must include differentiation strategies when requested'
  },
  {
    rule: 'p5_integration',
    check: (modul) => {
      const p5Dimensions = [
        'beriman', 'berkebinekaan', 'gotong_royong', 
        'kreatif', 'berkarakter', 'mandiri'
      ];
      return modul.activities.some(act =>
        p5Dimensions.some(dim => act.p5Dimension === dim)
      );
    },
    message: 'Activities should integrate P5 dimensions'
  }
];
```

### Approval Workflow

```
┌─────────────────────────────────────────────────────────────┐
│                Modul Ajar Approval Workflow                   │
└─────────────────────────────────────────────────────────────┘

1. DRAFT STATUS
   ├─ AI generates Modul Ajar
   ├─ Validation passes
   └─ Status: PENDING_REVIEW

2. TEACHER REVIEW (Level 1)
   ├─ Teacher reviews module content
   ├─ Can edit activities
   ├─ Can request regeneration
   └─ Status: PENDING_PEER_REVIEW or BACK_TO_GENERATION

3. PEER REVIEW (Level 2) - Optional
   ├─ Peer reviews pedagogical soundness
   ├─ Checks activity feasibility
   └─ Status: PENDING_FINAL_APPROVAL or BACK_TO_TEACHER

4. FINAL APPROVAL
   ├─ Module marked as approved
   ├─ Version created (v1.0)
   └─ Ready for classroom use
```

---

## Workflow #4: Assessment Generation Workflow

### Domain Context
**Assessment Domain** - Evaluation component that creates various assessment types including formative, summative, diagnostic, and performance-based assessments aligned with learning objectives.

### User Inputs

```typescript
interface AssessmentGenerationInput {
  // Learning Context
  teachingPlanId: string;
  modulAjarIds: string[];
  learningObjectives: string[];
  
  // Assessment Parameters
  assessmentType: 'diagnostic' | 'formative' | 'summative' | 'performance' | 'portfolio';
  duration: number; // in minutes
  totalPoints: number;
  
  // Question Distribution
  questionTypes: {
    multipleChoice: number;
    shortAnswer: number;
    essay: number;
    practical: number;
  };
  
  // Difficulty Distribution
  difficultyLevels: {
    easy: number;    // percentage
    medium: number; // percentage
    hard: number;   // percentage
  };
  
  // Content Focus
  topicsToCover: string[];
  skillsToAssess: string[];
  
  // Special Requirements
  includeRubric: boolean;
  includeAnswerKey: boolean;
  allowPartialCredit: boolean;
  timeAllocation: 'strict' | 'flexible';
}
```

### AI Prompts

#### System Prompt (Version 1.1.0)

```
You are an expert assessment designer specializing in Indonesian education standards.
Your task is to create high-quality assessments that validly measure student learning 
while being fair, inclusive, and aligned with learning objectives.

CORE PRINCIPLES:
1. Alignment with learning objectives
2. Valid and reliable measurement
3. Fairness and inclusivity
4. Age-appropriate content
5. Clear instructions and criteria

OUTPUT STRUCTURE:
- Assessment overview
- Instructions for students
- Questions by type with point values
- Scoring rubric (if requested)
- Answer key (if requested)
- Time allocation per section
- Differentiation considerations

CONSTRAINTS:
- Questions must align with learning objectives
- Difficulty levels must match specified distribution
- Language must be clear and age-appropriate
- Include culturally relevant contexts
- Ensure accessibility for diverse learners
- Provide clear success criteria
```

#### User Prompt Template

```
Generate an assessment for the following context:

LEARNING CONTEXT:
- Teaching Plan: {teachingPlanId}
- Learning Objectives: {learningObjectives}
- Topics: {topicsToCover}
- Skills: {skillsToAssess}

ASSESSMENT PARAMETERS:
- Type: {assessmentType}
- Duration: {duration} minutes
- Total Points: {totalPoints}

QUESTION DISTRIBUTION:
- Multiple Choice: {multipleChoice} questions
- Short Answer: {shortAnswer} questions
- Essay: {essay} questions
- Practical: {practical} tasks

DIFFICULTY DISTRIBUTION:
- Easy: {easy}%
- Medium: {medium}%
- Hard: {hard}%

REQUIREMENTS:
- Include Rubric: {includeRubric}
- Include Answer Key: {includeAnswerKey}
- Partial Credit: {allowPartialCredit}
- Time Allocation: {timeAllocation}

Please generate an assessment that includes:
1. Clear assessment overview and objectives
2. Student instructions
3. Questions organized by type with point values
4. Time allocation per section
5. Scoring rubric (if requested)
6. Answer key with point allocations (if requested)
7. Differentiation considerations for diverse learners
8. Cultural and age-appropriate contexts
9. Clear success criteria for each question type
```

### Validation Rules

```typescript
const assessmentValidationRules = [
  {
    rule: 'objective_alignment',
    check: (assessment, input) => {
      return assessment.questions.every(q =>
        input.learningObjectives.some(obj =>
          q.objectives?.includes(obj) || 
          q.content.toLowerCase().includes(obj.substring(0, 15))
        )
      );
    },
    message: 'Questions must align with learning objectives'
  },
  {
    rule: 'difficulty_distribution',
    check: (assessment, input) => {
      const distribution = assessment.questions.reduce((acc, q) => {
        acc[q.difficulty]++;
        return acc;
      }, { easy: 0, medium: 0, hard: 0 });
      
      const total = assessment.questions.length;
      const easyPct = (distribution.easy / total) * 100;
      const mediumPct = (distribution.medium / total) * 100;
      const hardPct = (distribution.hard / total) * 100;
      
      return Math.abs(easyPct - input.difficultyLevels.easy) <= 10 &&
             Math.abs(mediumPct - input.difficultyLevels.medium) <= 10 &&
             Math.abs(hardPct - input.difficultyLevels.hard) <= 10;
    },
    message: 'Difficulty distribution must match specifications'
  },
  {
    rule: 'point_allocation_valid',
    check: (assessment, input) => {
      const totalPoints = assessment.questions.reduce((sum, q) => sum + q.points, 0);
      return totalPoints === input.totalPoints;
    },
    message: 'Total points must match specification'
  },
  {
    rule: 'time_allocation_feasible',
    check: (assessment, input) => {
      const estimatedTime = assessment.sections.reduce((sum, s) => 
        sum + s.estimatedTime, 0
      );
      return estimatedTime <= input.duration;
    },
    message: 'Estimated time must not exceed duration'
  },
  {
    rule: 'question_type_distribution',
    check: (assessment, input) => {
      const types = assessment.questions.reduce((acc, q) => {
        acc[q.type]++;
        return acc;
      }, { multipleChoice: 0, shortAnswer: 0, essay: 0, practical: 0 });
      
      return types.multipleChoice === input.questionTypes.multipleChoice &&
             types.shortAnswer === input.questionTypes.shortAnswer &&
             types.essay === input.questionTypes.essay &&
             types.practical === input.questionTypes.practical;
    },
    message: 'Question type distribution must match specification'
  }
];
```

### Approval Workflow

```
┌─────────────────────────────────────────────────────────────┐
│                Assessment Approval Workflow                   │
└─────────────────────────────────────────────────────────────┘

1. DRAFT STATUS
   ├─ AI generates assessment
   ├─ Validation passes
   └─ Status: PENDING_REVIEW

2. TEACHER REVIEW (Level 1)
   ├─ Teacher reviews questions
   ├─ Can edit questions
   ├─ Can adjust point values
   ├─ Can request regeneration
   └─ Status: PENDING_PEER_REVIEW or BACK_TO_GENERATION

3. PEER REVIEW (Level 2) - Required for Summative
   ├─ Peer checks question quality
   ├─ Validates difficulty levels
   ├─ Reviews answer key accuracy
   └─ Status: PENDING_COORDINATOR or BACK_TO_TEACHER

4. CURRICULUM COORDINATOR REVIEW (Level 3) - Summative Only
   ├─ Validates alignment with standards
   ├─ Checks fairness and bias
   └─ Status: APPROVED or REJECTED

5. FINAL APPROVAL
   ├─ Assessment marked as approved
   ├─ Version created (v1.0)
   └─ Ready for administration
```

---

## Workflow #5: Rubric Generation Workflow

### Domain Context
**Rubric Domain** - Assessment criteria definition component that creates detailed scoring rubrics for performance-based assessments, projects, and complex tasks.

### User Inputs

```typescript
interface RubricGenerationInput {
  // Assessment Context
  assessmentId: string;
  assessmentType: string;
  learningObjectives: string[];
  
  // Rubric Parameters
  rubricType: 'analytic' | 'holistic' | 'single-point';
  performanceLevels: number; // 3-5 levels
  totalPoints: number;
  
  // Criteria Definition
  criteriaCategories: string[];
  skillsToAssess: string[];
  
  // Performance Level Labels
  levelLabels?: {
    [key: number]: string; // e.g., { 4: 'Excellent', 3: 'Proficient', ... }
  };
  
  // Descriptive Detail
  includeDescriptors: boolean;
  includeExamples: boolean;
  includeWeighting: boolean;
  
  // Special Requirements
  ageAppropriate: boolean;
  studentFriendly: boolean;
  parentFriendly: boolean;
}
```

### AI Prompts

#### System Prompt (Version 1.0.0)

```
You are an expert rubric designer specializing in educational assessment.
Your task is to create clear, valid, and reliable rubrics that provide meaningful 
feedback to students while ensuring fair and consistent assessment.

CORE PRINCIPLES:
1. Clear and specific criteria
2. Distinct performance levels
3. Age-appropriate language
4. Actionable feedback
5. Alignment with learning objectives

OUTPUT STRUCTURE:
- Rubric overview and purpose
- Performance level descriptions
- Criteria with descriptors
- Scoring guidelines
- Feedback suggestions
- Student-friendly version (if requested)

CONSTRAINTS:
- Criteria must align with learning objectives
- Performance levels must be distinct
- Language must be clear and specific
- Descriptors must be observable
- Provide constructive feedback
- Ensure cultural relevance
```

#### User Prompt Template

```
Generate a rubric for the following context:

ASSESSMENT CONTEXT:
- Assessment ID: {assessmentId}
- Type: {assessmentType}
- Learning Objectives: {learningObjectives}

RUBRIC PARAMETERS:
- Type: {rubricType}
- Performance Levels: {performanceLevels}
- Total Points: {totalPoints}

CRITERIA:
- Categories: {criteriaCategories}
- Skills: {skillsToAssess}

PERFORMANCE LEVELS:
{levelLabels}

REQUIREMENTS:
- Include Descriptors: {includeDescriptors}
- Include Examples: {includeExamples}
- Include Weighting: {includeWeighting}
- Age Appropriate: {ageAppropriate}
- Student Friendly: {studentFriendly}
- Parent Friendly: {parentFriendly}

Please generate a comprehensive rubric that includes:
1. Clear rubric overview and purpose
2. Performance level descriptions with labels
3. Detailed criteria with specific descriptors
4. Scoring guidelines and point allocations
5. Feedback suggestions for each level
6. Student-friendly version (if requested)
7. Weighting system (if requested)
8. Examples of work at each level (if requested)
9. Clear distinction between performance levels
10. Alignment with learning objectives
```

### Validation Rules

```typescript
const rubricValidationRules = [
  {
    rule: 'criteria_alignment',
    check: (rubric, input) => {
      return rubric.criteria.every(crit =>
        input.learningObjectives.some(obj =>
          crit.objectives?.includes(obj) ||
          crit.description.toLowerCase().includes(obj.substring(0, 15))
        )
      );
    },
    message: 'Criteria must align with learning objectives'
  },
  {
    rule: 'performance_levels_distinct',
    check: (rubric, input) => {
      return rubric.performanceLevels.length === input.performanceLevels;
    },
    message: 'Must have specified number of performance levels'
  },
  {
    rule: 'descriptors_observable',
    check: (rubric) => {
      const observableVerbs = ['demonstrates', 'shows', 'uses', 'applies', 'creates', 'analyzes'];
      return rubric.criteria.every(crit =>
        crit.descriptors.every(desc =>
          observableVerbs.some(verb => desc.toLowerCase().includes(verb))
        )
      );
    },
    message: 'Descriptors must use observable verbs'
  },
  {
    rule: 'point_allocation_valid',
    check: (rubric, input) => {
      const totalPoints = rubric.criteria.reduce((sum, crit) => sum + crit.maxPoints, 0);
      return totalPoints === input.totalPoints;
    },
    message: 'Total points must match specification'
  },
  {
    rule: 'level_progression',
    check: (rubric) => {
      // Check that performance levels show progression
      for (let i = 1; i < rubric.performanceLevels.length; i++) {
        const current = rubric.performanceLevels[i];
        const previous = rubric.performanceLevels[i - 1];
        if (current.points <= previous.points) {
          return false;
        }
      }
      return true;
    },
    message: 'Performance levels must show point progression'
  }
];
```

### Approval Workflow

```
┌─────────────────────────────────────────────────────────────┐
│                  Rubric Approval Workflow                    │
└─────────────────────────────────────────────────────────────┘

1. DRAFT STATUS
   ├─ AI generates rubric
   ├─ Validation passes
   └─ Status: PENDING_REVIEW

2. TEACHER REVIEW (Level 1)
   ├─ Teacher reviews criteria
   ├─ Can edit descriptors
   ├─ Can adjust point values
   ├─ Can request regeneration
   └─ Status: PENDING_PEER_REVIEW or BACK_TO_GENERATION

3. PEER REVIEW (Level 2)
   ├─ Peer checks clarity and fairness
   ├─ Validates distinctness of levels
   └─ Status: PENDING_FINAL_APPROVAL or BACK_TO_TEACHER

4. FINAL APPROVAL
   ├─ Rubric marked as approved
   ├─ Version created (v1.0)
   └─ Ready for use with assessment
```

---

## Workflow #6: Narrative Report Generation Workflow

### Domain Context
**Narrative Report Domain** - Student progress reporting component that generates qualitative assessments of student performance, achievements, and areas for improvement.

### User Inputs

```typescript
interface NarrativeReportGenerationInput {
  // Student Context
  studentId: string;
  studentName: string;
  grade: string;
  class: string;
  reportingPeriod: {
    startDate: Date;
    endDate: Date;
    type: 'midterm' | 'final' | 'progress';
  };
  
  // Academic Performance
  assessmentResults: AssessmentResult[];
  learningObjectiveProgress: LearningObjectiveProgress[];
  attendance: AttendanceRecord;
  behavior: BehaviorRecord;
  
  // Report Parameters
  reportType: 'academic' | 'behavioral' | 'comprehensive';
  tone: 'formal' | 'encouraging' | 'constructive';
  language: 'indonesian' | 'bilingual';
  
  // Focus Areas
  highlightStrengths: boolean;
  includeRecommendations: boolean;
  includeGoals: boolean;
  includeParentSuggestions: boolean;
  
  // Special Considerations
  specialAccomplishments: string[];
  areasOfConcern: string[];
  personalCircumstances: string[];
}
```

### AI Prompts

#### System Prompt (Version 1.0.0)

```
You are an expert educational reporter specializing in student progress communication.
Your task is to generate narrative reports that provide meaningful, constructive, and 
actionable feedback to students and parents while celebrating achievements and 
identifying areas for growth.

CORE PRINCIPLES:
1. Strength-based approach
2. Specific and evidence-based
3. Constructive and actionable
4. Age-appropriate language
5. Culturally sensitive
6. Confidential and professional

OUTPUT STRUCTURE:
- Student overview
- Academic performance summary
- Strengths and achievements
- Areas for improvement
- Specific recommendations
- Goals for next period
- Parent communication suggestions

CONSTRAINTS:
- Must be based on provided evidence
- Language must be age-appropriate
- Tone must be constructive
- Include specific examples
- Avoid generalizations
- Respect confidentiality
- Consider cultural context
```

#### User Prompt Template

```
Generate a narrative report for the following student:

STUDENT INFORMATION:
- Name: {studentName}
- Grade: {grade}
- Class: {class}
- Reporting Period: {reportingPeriod.type} ({reportingPeriod.startDate} to {reportingPeriod.endDate})

ACADEMIC PERFORMANCE:
- Assessment Results: {assessmentResultsSummary}
- Learning Objective Progress: {learningObjectiveProgressSummary}
- Attendance: {attendanceSummary}
- Behavior: {behaviorSummary}

REPORT PARAMETERS:
- Type: {reportType}
- Tone: {tone}
- Language: {language}

FOCUS AREAS:
- Highlight Strengths: {highlightStrengths}
- Include Recommendations: {includeRecommendations}
- Include Goals: {includeGoals}
- Parent Suggestions: {includeParentSuggestions}

SPECIAL CONSIDERATIONS:
- Accomplishments: {specialAccomplishments}
- Concerns: {areasOfConcern}
- Circumstances: {personalCircumstances}

Please generate a narrative report that includes:
1. Student overview with positive opening
2. Academic performance summary with specific evidence
3. Strengths and achievements with examples
4. Areas for improvement with constructive feedback
5. Specific, actionable recommendations
6. Goals for the next reporting period
7. Suggestions for parent support
8. Age-appropriate language
9. Culturally sensitive content
10. Evidence-based statements
```

### Validation Rules

```typescript
const narrativeReportValidationRules = [
  {
    rule: 'evidence_based',
    check: (report, input) => {
      // Check that claims are supported by provided data
      const claims = extractClaims(report.content);
      return claims.every(claim =>
        isSupportedByEvidence(claim, input.assessmentResults, input.learningObjectiveProgress)
      );
    },
    message: 'All claims must be supported by provided evidence'
  },
  {
    rule: 'constructive_tone',
    check: (report) => {
      const negativePhrases = ['cannot', 'unable', 'fails', 'poor', 'bad'];
      const content = report.content.toLowerCase();
      const negativeCount = negativePhrases.filter(phrase => 
        content.includes(phrase)
      ).length;
      const positivePhrases = ['excellent', 'good', 'improves', 'achieves', 'successful'];
      const positiveCount = positivePhrases.filter(phrase => 
        content.includes(phrase)
      ).length;
      return positiveCount >= negativeCount * 2; // At least 2:1 positive to negative
    },
    message: 'Report should maintain constructive tone'
  },
  {
    rule: 'specific_examples',
    check: (report) => {
      const vaguePhrases = ['always', 'never', 'usually', 'sometimes'];
      const content = report.content.toLowerCase();
      const vagueCount = vaguePhrases.filter(phrase => 
        content.includes(phrase)
      ).length;
      return vagueCount < 3; // Limit vague generalizations
    },
    message: 'Report should include specific examples rather than generalizations'
  },
  {
    rule: 'includes_recommendations',
    check: (report, input) => {
      if (input.includeRecommendations) {
        return report.recommendations?.length > 0;
      }
      return true;
    },
    message: 'Must include recommendations when requested'
  },
  {
    rule: 'age_appropriate',
    check: (report, input) => {
      const gradeLevel = parseInt(input.grade);
      const avgSentenceLength = calculateAvgSentenceLength(report.content);
      // Younger grades should have shorter sentences
      const maxLength = gradeLevel <= 3 ? 12 : gradeLevel <= 6 ? 15 : 20;
      return avgSentenceLength <= maxLength;
    },
    message: 'Language should be age-appropriate'
  }
];
```

### Approval Workflow

```
┌─────────────────────────────────────────────────────────────┐
│             Narrative Report Approval Workflow               │
└─────────────────────────────────────────────────────────────┘

1. DRAFT STATUS
   ├─ AI generates narrative report
   ├─ Validation passes
   └─ Status: PENDING_REVIEW

2. TEACHER REVIEW (Level 1)
   ├─ Teacher reviews accuracy
   ├─ Can edit content
   ├─ Can add personal observations
   ├─ Can request regeneration
   └─ Status: PENDING_COORDINATOR or BACK_TO_GENERATION

3. CURRICULUM COORDINATOR REVIEW (Level 2)
   ├─ Checks tone and professionalism
   ├─ Validates evidence-based claims
   ├─ Ensures cultural sensitivity
   └─ Status: PENDING_PRINCIPAL or BACK_TO_TEACHER

4. PRINCIPAL APPROVAL (Level 3) - Final Reports Only
   ├─ Final approval for official reports
   └─ Status: APPROVED or REJECTED

5. FINAL APPROVAL
   ├─ Report marked as approved
   ├─ Version created (v1.0)
   └─ Ready for distribution to parents
```

---

## Final AI Orchestration Architecture

### System Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                    AI Orchestration System                       │
├─────────────────────────────────────────────────────────────────┤
│                                                                  │
│  ┌──────────────────────────────────────────────────────────┐   │
│  │              API Gateway Layer                            │   │
│  │  (REST + GraphQL + WebSocket)                            │   │
│  └──────────────────────────────────────────────────────────┘   │
│                              │                                  │
│                              ▼                                  │
│  ┌──────────────────────────────────────────────────────────┐   │
│  │            Workflow Orchestrator                          │   │
│  │  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐    │   │
│  │  │  TP Workflow │  │ ATP Workflow │  │ Modul Workflow│    │   │
│  │  └──────────────┘  └──────────────┘  └──────────────┘    │   │
│  │  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐    │   │
│  │  │ Assessment   │  │   Rubric     │  │  Narrative   │    │   │
│  │  │  Workflow    │  │  Workflow    │  │  Workflow    │    │   │
│  │  └──────────────┘  └──────────────┘  └──────────────┘    │   │
│  └──────────────────────────────────────────────────────────┘   │
│                              │                                  │
│                              ▼                                  │
│  ┌──────────────────────────────────────────────────────────┐   │
│  │              Core Services Layer                         │   │
│  │  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐    │   │
│  │  │ Prompt       │  │ Validation   │  │ Approval     │    │   │
│  │  │ Engine       │  │ Engine       │  │ Engine       │    │   │
│  │  └──────────────┘  └──────────────┘  └──────────────┘    │   │
│  │  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐    │   │
│  │  │ Audit        │  │ Rollback     │  │ Version      │    │   │
│  │  │ Trail        │  │ Manager      │  │ Control      │    │   │
│  │  └──────────────┘  └──────────────┘  └──────────────┘    │   │
│  └──────────────────────────────────────────────────────────┘   │
│                              │                                  │
│                              ▼                                  │
│  ┌──────────────────────────────────────────────────────────┐   │
│  │              AI Provider Layer                             │   │
│  │  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐    │   │
│  │  │   OpenAI     │  │   Gemini     │  │   Claude     │    │   │
│  │  │   Provider   │  │   Provider   │  │   Provider   │    │   │
│  │  └──────────────┘  └──────────────┘  └──────────────┘    │   │
│  └──────────────────────────────────────────────────────────┘   │
│                              │                                  │
│                              ▼                                  │
│  ┌──────────────────────────────────────────────────────────┐   │
│  │              Data Layer                                   │   │
│  │  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐    │   │
│  │  │ PostgreSQL   │  │   Redis      │  │   Qdrant     │    │   │
│  │  │   (Primary)  │  │   (Cache)    │  │   (Vectors)  │    │   │
│  │  └──────────────┘  └──────────────┘  └──────────────┘    │   │
│  │  ┌──────────────┐  ┌──────────────┐                       │   │
│  │  │   MinIO      │  │  RabbitMQ    │                       │   │
│  │  │  (Storage)   │  │   (Queue)    │                       │   │
│  │  └──────────────┘  └──────────────┘                       │   │
│  └──────────────────────────────────────────────────────────┘   │
│                                                                  │
└─────────────────────────────────────────────────────────────────┘
```

### Prompt Versioning System

```typescript
interface PromptVersion {
  id: string;
  workflow: 'TP' | 'ATP' | 'MODUL_AJAR' | 'ASSESSMENT' | 'RUBRIC' | 'NARRATIVE';
  version: string; // Semantic versioning (e.g., 1.0.0)
  type: 'system' | 'user';
  content: string;
  
  // Metadata
  createdAt: ISO8601;
  createdBy: string;
  changeLog: string;
  
  // Performance Metrics
  metrics: {
    averageQualityScore: number;
    approvalRate: number;
    regenerationRate: number;
    averageTokens: number;
  };
  
  // Status
  status: 'draft' | 'active' | 'deprecated' | 'archived';
  
  // A/B Testing
  abTestGroup?: 'A' | 'B';
  abTestResults?: {
    groupSize: number;
    averageQualityScore: number;
    approvalRate: number;
  };
}

class PromptVersionManager {
  async createVersion(workflow: string, content: string, changeLog: string): Promise<PromptVersion> {
    const latestVersion = await this.getLatestVersion(workflow);
    const newVersion = incrementVersion(latestVersion.version);
    
    return await this.repository.create({
      id: generateId(),
      workflow,
      version: newVersion,
      type: 'system',
      content,
      createdAt: new Date().toISOString(),
      createdBy: 'system',
      changeLog,
      metrics: {
        averageQualityScore: 0,
        approvalRate: 0,
        regenerationRate: 0,
        averageTokens: 0
      },
      status: 'draft'
    });
  }
  
  async activateVersion(versionId: string): Promise<void> {
    // Deactivate current active version
    await this.repository.deactivateActiveVersions(workflow);
    
    // Activate new version
    await this.repository.update(versionId, { status: 'active' });
    
    // Archive old versions
    await this.repository.archiveOldVersions(workflow);
  }
  
  async rollbackVersion(workflow: string, targetVersion: string): Promise<void> {
    // Validate target version exists
    const target = await this.repository.findByVersion(workflow, targetVersion);
    if (!target) {
      throw new Error(`Version ${targetVersion} not found`);
    }
    
    // Create rollback record
    await this.auditTrail.record({
      action: 'prompt_rollback',
      workflow,
      fromVersion: await this.getActiveVersion(workflow),
      toVersion: targetVersion,
      timestamp: new Date().toISOString()
    });
    
    // Activate target version
    await this.activateVersion(target.id);
  }
  
  async getActiveVersion(workflow: string): Promise<PromptVersion> {
    return await this.repository.findActive(workflow);
  }
  
  async compareVersions(versionA: string, versionB: string): Promise<DiffResult> {
    const vA = await this.repository.findById(versionA);
    const vB = await this.repository.findById(versionB);
    
    return {
      contentDiff: computeDiff(vA.content, vB.content),
      metricsDiff: {
        qualityScore: vB.metrics.averageQualityScore - vA.metrics.averageQualityScore,
        approvalRate: vB.metrics.approvalRate - vA.metrics.approvalRate,
        regenerationRate: vB.metrics.regenerationRate - vA.metrics.regenerationRate
      }
    };
  }
}
```

### Human Approval Checkpoints

```typescript
interface ApprovalCheckpoint {
  id: string;
  workflow: string;
  level: number;
  name: string;
  description: string;
  
  // Approval Criteria
  requiredRoles: string[];
  requiredApprovals: number;
  
  // Conditions
  autoApproveConditions?: AutoApproveCondition[];
  skipConditions?: SkipCondition[];
  
  // Actions
  onApprove: string[];
  onReject: string[];
  onRequestChanges: string[];
}

class ApprovalEngine {
  checkpoints: Map<string, ApprovalCheckpoint> = new Map();
  
  constructor() {
    this.initializeCheckpoints();
  }
  
  private initializeCheckpoints() {
    // TP Workflow Checkpoints
    this.checkpoints.set('TP_LEVEL_1', {
      id: 'TP_LEVEL_1',
      workflow: 'TP',
      level: 1,
      name: 'Teacher Review',
      description: 'Teacher reviews generated TP for accuracy and completeness',
      requiredRoles: ['teacher'],
      requiredApprovals: 1,
      autoApproveConditions: [
        { condition: 'quality_score >= 0.9', action: 'auto_approve' },
        { condition: 'no_validation_errors', action: 'auto_approve' }
      ],
      onApprove: ['status:PENDING_PEER_REVIEW', 'notify:peer_reviewers'],
      onReject: ['status:REJECTED', 'notify:teacher'],
      onRequestChanges: ['status:BACK_TO_GENERATION', 'regenerate_with_feedback']
    });
    
    this.checkpoints.set('TP_LEVEL_2', {
      id: 'TP_LEVEL_2',
      workflow: 'TP',
      level: 2,
      name: 'Peer Review',
      description: 'Peer teacher reviews TP for pedagogical soundness',
      requiredRoles: ['teacher'],
      requiredApprovals: 1,
      skipConditions: [
        { condition: 'peer_review_disabled', action: 'skip' }
      ],
      onApprove: ['status:PENDING_COORDINATOR', 'notify:coordinator'],
      onReject: ['status:BACK_TO_TEACHER', 'notify:teacher'],
      onRequestChanges: ['status:BACK_TO_TEACHER', 'notify:teacher']
    });
    
    this.checkpoints.set('TP_LEVEL_3', {
      id: 'TP_LEVEL_3',
      workflow: 'TP',
      level: 3,
      name: 'Coordinator Approval',
      description: 'Curriculum coordinator approves TP for Kurikulum Merdeka compliance',
      requiredRoles: ['curriculum_coordinator'],
      requiredApprovals: 1,
      onApprove: ['status:APPROVED', 'create_version:v1.0', 'notify:teacher'],
      onReject: ['status:REJECTED', 'notify:teacher'],
      onRequestChanges: ['status:BACK_TO_TEACHER', 'notify:teacher']
    });
    
    // Similar checkpoints for other workflows...
  }
  
  async processApproval(workflow: string, level: number, action: 'approve' | 'reject' | 'request_changes', context: any): Promise<ApprovalResult> {
    const checkpoint = this.checkpoints.get(`${workflow}_LEVEL_${level}`);
    if (!checkpoint) {
      throw new Error(`Checkpoint not found for ${workflow} level ${level}`);
    }
    
    // Check if user has required role
    if (!checkpoint.requiredRoles.includes(context.user.role)) {
      throw new Error('User does not have required role');
    }
    
    // Check auto-approve conditions
    if (action === 'approve' && checkpoint.autoApproveConditions) {
      for (const condition of checkpoint.autoApproveConditions) {
        if (await this.evaluateCondition(condition.condition, context)) {
          return await this.executeActions(condition.action, context);
        }
      }
    }
    
    // Execute action
    const actions = checkpoint[`on${action.charAt(0).toUpperCase() + action.slice(1)}`];
    return await this.executeActions(actions, context);
  }
  
  private async evaluateCondition(condition: string, context: any): Promise<boolean> {
    // Implement condition evaluation logic
    return true;
  }
  
  private async executeActions(actions: string[], context: any): Promise<ApprovalResult> {
    const results = [];
    for (const action of actions) {
      const [command, ...args] = action.split(':');
      const result = await this.executeCommand(command, args, context);
      results.push(result);
    }
    return { success: true, actions: results };
  }
  
  private async executeCommand(command: string, args: string[], context: any): Promise<any> {
    // Implement command execution logic
    switch (command) {
      case 'status':
        return await this.updateStatus(context.entityId, args[0]);
      case 'notify':
        return await this.sendNotification(args[0], context);
      case 'create_version':
        return await this.createVersion(context.entityId, args[0]);
      case 'regenerate_with_feedback':
        return await this.regenerate(context.entityId, context.feedback);
      default:
        throw new Error(`Unknown command: ${command}`);
    }
  }
}
```

### Rollback Strategy

```typescript
interface RollbackStrategy {
  type: 'full' | 'partial' | 'content_only' | 'metadata_only';
  targetVersion: string;
  preserveFields: string[];
  rollbackReason: string;
}

class RollbackManager {
  async executeRollback(entityId: string, strategy: RollbackStrategy): Promise<RollbackResult> {
    // Get current version
    const currentVersion = await this.versionControl.getCurrentVersion(entityId);
    
    // Get target version
    const targetVersion = await this.versionControl.getVersion(entityId, strategy.targetVersion);
    if (!targetVersion) {
      throw new Error(`Target version ${strategy.targetVersion} not found`);
    }
    
    // Create backup of current version
    const backupId = await this.createBackup(currentVersion);
    
    // Execute rollback based on strategy
    switch (strategy.type) {
      case 'full':
        return await this.fullRollback(entityId, targetVersion, strategy);
      case 'partial':
        return await this.partialRollback(entityId, targetVersion, strategy);
      case 'content_only':
        return await this.contentRollback(entityId, targetVersion, strategy);
      case 'metadata_only':
        return await this.metadataRollback(entityId, targetVersion, strategy);
    }
  }
  
  private async fullRollback(entityId: string, targetVersion: any, strategy: RollbackStrategy): Promise<RollbackResult> {
    // Rollback all fields
    await this.repository.update(entityId, {
      content: targetVersion.content,
      metadata: targetVersion.metadata,
      status: 'rolled_back',
      rollbackVersion: targetVersion.version,
      rollbackReason: strategy.rollbackReason,
      rollbackTimestamp: new Date().toISOString()
    });
    
    // Record audit trail
    await this.auditTrail.record({
      action: 'full_rollback',
      entityId,
      fromVersion: await this.getCurrentVersion(entityId),
      toVersion: strategy.targetVersion,
      reason: strategy.rollbackReason,
      timestamp: new Date().toISOString()
    });
    
    return { success: true, backupId, newVersion: strategy.targetVersion };
  }
  
  private async partialRollback(entityId: string, targetVersion: any, strategy: RollbackStrategy): Promise<RollbackResult> {
    // Rollback only specified fields
    const updates = {};
    for (const field of strategy.preserveFields) {
      updates[field] = targetVersion[field];
    }
    
    await this.repository.update(entityId, {
      ...updates,
      status: 'partially_rolled_back',
      rollbackVersion: targetVersion.version,
      rollbackReason: strategy.rollbackReason
    });
    
    return { success: true, newVersion: await this.getCurrentVersion(entityId) };
  }
  
  private async createBackup(version: any): Promise<string> {
    const backupId = generateId();
    await this.backupRepository.create({
      id: backupId,
      entityId: version.entityId,
      version: version.version,
      content: version.content,
      metadata: version.metadata,
      createdAt: new Date().toISOString()
    });
    return backupId;
  }
  
  async restoreFromBackup(backupId: string): Promise<RestoreResult> {
    const backup = await this.backupRepository.findById(backupId);
    if (!backup) {
      throw new Error(`Backup ${backupId} not found`);
    }
    
    await this.repository.update(backup.entityId, {
      content: backup.content,
      metadata: backup.metadata,
      status: 'restored',
      restoreTimestamp: new Date().toISOString()
    });
    
    return { success: true, restoredVersion: backup.version };
  }
}
```

### Hallucination Mitigation System

```typescript
interface HallucinationCheck {
  type: 'factual' | 'logical' | 'consistency' | 'alignment';
  severity: 'critical' | 'high' | 'medium' | 'low';
  check: (content: any, context: any) => Promise<boolean>;
  message: string;
  mitigation: string;
}

class HallucinationMitigationEngine {
  checks: HallucinationCheck[] = [];
  
  constructor() {
    this.initializeChecks();
  }
  
  private initializeChecks() {
    // Factual Accuracy Checks
    this.checks.push({
      type: 'factual',
      severity: 'critical',
      check: async (content, context) => {
        // Check against knowledge base
        const facts = extractFacts(content);
        for (const fact of facts) {
          const verified = await this.knowledgeBase.verify(fact);
          if (!verified) return false;
        }
        return true;
      },
      message: 'Content contains unverified factual claims',
      mitigation: 'Flag for human review and provide source verification'
    });
    
    // Logical Consistency Checks
    this.checks.push({
      type: 'logical',
      severity: 'high',
      check: async (content, context) => {
        // Check for logical contradictions
        return !hasContradictions(content);
      },
      message: 'Content contains logical contradictions',
      mitigation: 'Request regeneration with consistency constraints'
    });
    
    // Alignment Checks
    this.checks.push({
      type: 'alignment',
      severity: 'high',
      check: async (content, context) => {
        // Check alignment with input parameters
        return isAligned(content, context.input);
      },
      message: 'Content does not align with input parameters',
      mitigation: 'Regenerate with stronger alignment constraints'
    });
    
    // Consistency Checks
    this.checks.push({
      type: 'consistency',
      severity: 'medium',
      check: async (content, context) => {
        // Check internal consistency
        return isInternallyConsistent(content);
      },
      message: 'Content has internal inconsistencies',
      mitigation: 'Request regeneration with consistency checks'
    });
  }
  
  async validate(content: any, context: any): Promise<ValidationResult> {
    const results = [];
    let overallSeverity = 'low';
    
    for (const check of this.checks) {
      const passed = await check.check(content, context);
      results.push({
        type: check.type,
        severity: check.severity,
        passed,
        message: passed ? null : check.message,
        mitigation: passed ? null : check.mitigation
      });
      
      if (!passed && this.getSeverityLevel(check.severity) > this.getSeverityLevel(overallSeverity)) {
        overallSeverity = check.severity;
      }
    }
    
    return {
      passed: results.every(r => r.passed),
      severity: overallSeverity,
      checks: results
    };
  }
  
  private getSeverityLevel(severity: string): number {
    const levels = { critical: 4, high: 3, medium: 2, low: 1 };
    return levels[severity] || 0;
  }
  
  async mitigate(content: any, validation: ValidationResult, context: any): Promise<MitigationResult> {
    const failedChecks = validation.checks.filter(c => !c.passed);
    
    for (const check of failedChecks) {
      switch (check.type) {
        case 'factual':
          return await this.mitigateFactual(content, context);
        case 'logical':
          return await this.mitigateLogical(content, context);
        case 'alignment':
          return await this.mitigateAlignment(content, context);
        case 'consistency':
          return await this.mitigateConsistency(content, context);
      }
    }
    
    return { success: true, mitigatedContent: content };
  }
  
  private async mitigateFactual(content: any, context: any): Promise<MitigationResult> {
    // Add fact-checking prompt
    const enhancedPrompt = context.prompt + `
IMPORTANT: Verify all factual claims against the provided knowledge base.
If a fact cannot be verified, either:
1. Remove the claim
2. Add a qualifier (e.g., "according to some sources")
3. Request human verification
`;
    
    const regenerated = await this.aiProvider.generate(enhancedPrompt);
    return { success: true, mitigatedContent: regenerated };
  }
  
  private async mitigateLogical(content: any, context: any): Promise<MitigationResult> {
    // Add consistency constraints
    const enhancedPrompt = context.prompt + `
IMPORTANT: Ensure logical consistency throughout the response.
Check for:
- Contradictory statements
- Inconsistent terminology
- Logical flow issues
`;
    
    const regenerated = await this.aiProvider.generate(enhancedPrompt);
    return { success: true, mitigatedContent: regenerated };
  }
  
  private async mitigateAlignment(content: any, context: any): Promise<MitigationResult> {
    // Strengthen alignment constraints
    const enhancedPrompt = context.prompt + `
IMPORTANT: Strictly align with the following parameters:
${JSON.stringify(context.input, null, 2)}

Do not deviate from these requirements.
`;
    
    const regenerated = await this.aiProvider.generate(enhancedPrompt);
    return { success: true, mitigatedContent: regenerated };
  }
  
  private async mitigateConsistency(content: any, context: any): Promise<MitigationResult> {
    // Add internal consistency checks
    const enhancedPrompt = context.prompt + `
IMPORTANT: Ensure internal consistency:
- Use consistent terminology throughout
- Maintain consistent tone and style
- Ensure all sections reference each other correctly
`;
    
    const regenerated = await this.aiProvider.generate(enhancedPrompt);
    return { success: true, mitigatedContent: regenerated };
  }
}
```

### Complete Audit Trail System

```typescript
interface AuditEntry {
  id: string;
  entityType: 'TP' | 'ATP' | 'MODUL_AJAR' | 'ASSESSMENT' | 'RUBRIC' | 'NARRATIVE';
  entityId: string;
  timestamp: ISO8601;
  
  // Event Details
  eventType: 'generated' | 'validated' | 'approved' | 'rejected' | 'edited' | 'regenerated' | 'rolled_back';
  
  // Actor Information
  actor: {
    type: 'ai' | 'user' | 'system';
    id: string;
    name: string;
    role?: string;
  };
  
  // Change Details
  changes: {
    before?: any;
    after?: any;
    diff?: string;
    fieldsChanged?: string[];
  };
  
  // Technical Details
  metadata: {
    promptVersion: string;
    modelVersion: string;
    modelProvider: string;
    tokensUsed?: number;
    latency?: number;
    validationResults?: ValidationResult[];
    approvalLevel?: number;
    checkpointId?: string;
  };
  
  // Context
  context: {
    requestId?: string;
    correlationId?: string;
    sessionId?: string;
    ipAddress?: string;
    userAgent?: string;
  };
}

class AuditTrailManager {
  async record(entry: Omit<AuditEntry, 'id' | 'timestamp'>): Promise<string> {
    const auditEntry: AuditEntry = {
      id: generateId(),
      timestamp: new Date().toISOString(),
      ...entry
    };
    
    await this.repository.create(auditEntry);
    
    // Also send to event stream for real-time monitoring
    await this.eventStream.publish('audit_event', auditEntry);
    
    return auditEntry.id;
  }
  
  async getHistory(entityId: string): Promise<AuditEntry[]> {
    return await this.repository.findByEntityId(entityId);
  }
  
  async getHistoryByType(entityType: string, startDate: Date, endDate: Date): Promise<AuditEntry[]> {
    return await this.repository.findByEntityTypeAndDateRange(entityType, startDate, endDate);
  }
  
  async getHistoryByActor(actorId: string): Promise<AuditEntry[]> {
    return await this.repository.findByActorId(actorId);
  }
  
  async generateReport(filters: AuditFilters): Promise<AuditReport> {
    const entries = await this.repository.findWithFilters(filters);
    
    return {
      totalEvents: entries.length,
      eventsByType: this.groupByType(entries),
      eventsByActor: this.groupByActor(entries),
      eventsByDate: this.groupByDate(entries),
      averageLatency: this.calculateAverageLatency(entries),
      totalTokens: this.calculateTotalTokens(entries),
      approvalRate: this.calculateApprovalRate(entries),
      regenerationRate: this.calculateRegenerationRate(entries)
    };
  }
  
  private groupByType(entries: AuditEntry[]): Record<string, number> {
    return entries.reduce((acc, entry) => {
      acc[entry.eventType] = (acc[entry.eventType] || 0) + 1;
      return acc;
    }, {});
  }
  
  private groupByActor(entries: AuditEntry[]): Record<string, number> {
    return entries.reduce((acc, entry) => {
      const key = `${entry.actor.type}:${entry.actor.name}`;
      acc[key] = (acc[key] || 0) + 1;
      return acc;
    }, {});
  }
  
  private groupByDate(entries: AuditEntry[]): Record<string, number> {
    return entries.reduce((acc, entry) => {
      const date = entry.timestamp.split('T')[0];
      acc[date] = (acc[date] || 0) + 1;
      return acc;
    }, {});
  }
  
  private calculateAverageLatency(entries: AuditEntry[]): number {
    const latencies = entries
      .map(e => e.metadata.latency)
      .filter(l => l !== undefined);
    return latencies.length > 0 
      ? latencies.reduce((a, b) => a + b, 0) / latencies.length 
      : 0;
  }
  
  private calculateTotalTokens(entries: AuditEntry[]): number {
    return entries
      .map(e => e.metadata.tokensUsed || 0)
      .reduce((a, b) => a + b, 0);
  }
  
  private calculateApprovalRate(entries: AuditEntry[]): number {
    const approvals = entries.filter(e => e.eventType === 'approved').length;
    const total = entries.filter(e => e.eventType === 'approved' || e.eventType === 'rejected').length;
    return total > 0 ? (approvals / total) * 100 : 0;
  }
  
  private calculateRegenerationRate(entries: AuditEntry[]): number {
    const regenerations = entries.filter(e => e.eventType === 'regenerated').length;
    const total = entries.filter(e => e.eventType === 'generated' || e.eventType === 'regenerated').length;
    return total > 0 ? (regenerations / total) * 100 : 0;
  }
}
```

### Implementation Summary

This AI Orchestration Architecture provides:

1. **Complete Workflow Coverage**: All 6 educational content generation workflows with detailed specifications
2. **Robust Validation**: Multi-layer validation (structural, content, factual, logical)
3. **Human-in-the-Loop**: Multi-level approval workflows with role-based access
4. **Prompt Versioning**: Semantic versioning with A/B testing and rollback capabilities
5. **Hallucination Mitigation**: Multi-layer checks with automatic mitigation strategies
6. **Complete Audit Trail**: Full traceability of all AI-generated content
7. **Rollback Strategy**: Multiple rollback strategies with backup and restore
8. **Scalable Architecture**: Modular design supporting multiple AI providers
9. **Kurikulum Merdeka Alignment**: All workflows aligned with Indonesian curriculum standards
10. **Production-Ready**: Error handling, monitoring, and performance metrics

The architecture is designed to be implemented incrementally, with each workflow able to be deployed independently while sharing core services (validation, approval, audit trail, etc.).

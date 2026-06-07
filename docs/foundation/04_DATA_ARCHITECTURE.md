# 04_DATA_ARCHITECTURE.md

## Foundation Document for NUSA Education Platform

**Version**: 1.0
**Date**: June 2026
**Status**: FOUNDATION DOCUMENT
**Alignment**: Validated against Foundation Architecture (00A, 00B, 00C, 01, 02, 03)

**Purpose**: Define the complete data architecture for NUSA Education Platform, serving as the official translation from Business Process Architecture to Data Architecture. This document is the single source of truth for Data Model, Data Classification, Data Lifecycle, Domain Event Strategy, Data Governance, and Data Security.

---

# SECTION 1 — Executive Summary

## Why Data Architecture is Required

Data Architecture is the critical layer that defines WHAT data entities and structures are needed to support the business processes defined in the Business Process Architecture (03). While the Business Process Architecture defines HOW capabilities are executed through workflows, the Data Architecture defines WHAT data entities capture educational activities and outcomes.

### Architecture Translation Chain

```
Business Process (03)
    ↓ generates
Events
    → updates
Data (04)
    → triggers
AI Actions (05)
    → implemented by
Applications (06)
```

**Business Process**: The workflows and procedures that execute capabilities (CP → TP transformation, Assessment evidence collection).

**Events**: The significant state changes that trigger workflows and AI actions.

**Data**: The information entities that capture educational activities and outcomes.

**AI Actions**: The automated responses and recommendations that AI provides based on events and data.

**Applications**: The software systems that implement the processes and manage the data.

### Primary Objectives

The Data Architecture is designed to support:

#### Competency Graph as Educational Brain
- All data entities must contribute to Competency Graph updates
- Competency Graph serves as the central intelligence backbone
- Every learning activity must be traceable to competency development

#### Digital Twin as Living Representation
- Data architecture supports real-time Digital Twin updates
- Student state model captures current competency status
- Digital Twin enables personalized learning and prediction

#### Lifelong Learning Record as Continuous Identity
- Data architecture supports learning identity across educational phases
- Learning records persist from PAUD through adulthood
- Credentials and achievements are recognized and portable

#### AI-Ready Data Foundation
- Data structures designed for AI consumption
- Domain events enable AI to learn from operational patterns
- Data quality standards ensure reliable AI decision-making

#### Human Governed Data
- Critical data requires human validation
- Data governance ensures privacy and security
- Human oversight maintains educational judgment

### Relationship with Foundation Documents

This Data Architecture is derived from and validated against:

- **00A_NATIONAL_EDUCATION_DIRECTION_2045.md**: The strategic direction and national vision for education transformation
- **00B_PRODUCT_VISION.md**: The product vision for AI-Native NUSA Education Platform with 90% AI assistance
- **00C_EDUCATION_OPERATING_SYSTEM_PRINCIPLES.md**: The architectural principles (Curriculum-Centered, Learning > Administration)
- **01_EDUCATION_DOMAIN_MODEL.md**: The domain entities and relationships that data must capture
- **02_CAPABILITY_MODEL.md**: The 16 Level 1 capabilities that data must support
- **03_BUSINESS_PROCESS_ARCHITECTURE.md**: The business processes that produce and consume data

All data entities defined in this document are traceable to processes defined in 03, which are traceable to capabilities defined in 02, which are traceable to domains defined in 01, which are aligned with principles in 00C, which support the vision in 00B, which enables the direction in 00A.

### Core Principle

**Data entities are the information structures that capture educational activities and outcomes.**

Every data entity must be traceable to a business process defined in 03, which must be traceable to a capability defined in 02. If a data entity is not defined here, it should not exist in the system.

---

# Future-State Strategic Data Domains

Competency Graph, Digital Twin, and Lifelong Learning Record remain part of the future-state NUSA architecture.

These domains are included to preserve long-term architectural direction and strategic alignment.

They are not implemented, deployed, or developed in MVP Wave 1.

MVP Wave 1 focuses only on:

- Curriculum
- Learning Planning
- Assessment
- Reporting
- Administration
- Authentication
- AI Orchestration

Future-state domains must not introduce MVP implementation requirements.

All references throughout this document to Competency Graph, Digital Twin, and Lifelong Learning Record are for long-term architectural planning purposes only and do not represent MVP Wave 1 implementation scope.

---

# SECTION 2 — Data Architecture Principles

## Principle 1: Single Source of Truth

**Statement**: Every data entity has one authoritative source and one canonical definition.

**Rationale**: Duplicate or inconsistent data definitions lead to data quality issues, integration problems, and unreliable analytics. Single Source of Truth ensures data consistency across the system.

**Implementation**:
- Each data entity has one canonical definition in this document
- All systems reference the same data definition
- Data ownership is clearly assigned
- Data changes are managed through defined processes

**Examples**:
- Student data has one canonical definition across all systems
- Competency Graph has one authoritative source
- Assessment evidence has one consistent structure

## Principle 2: Event Driven Data

**Statement**: Data changes are captured as operational domain events for workflow coordination and notifications.

**Rationale**: Operational domain events enable asynchronous workflows, module coordination, and real-time notifications through RabbitMQ. Events provide the foundation for workflow progression and process coordination.

**Implementation**:
- Operational events are published for state changes
- Events are consumed via RabbitMQ for workflow coordination
- PostgreSQL transactional tables remain the source of truth
- Events enable asynchronous processing and notifications

**Examples**:
- TPGenerated event triggers ATP generation workflow
- TPApproved event notifies downstream modules
- AssessmentCompleted event triggers evaluation workflow
- AIGenerationCompleted event notifies AI orchestration completion

## Principle 3: Competency First

**Statement**: All data entities must contribute to competency development and Competency Graph updates.

**Rationale**: The ultimate goal of education is competency development. All data must be designed to support this goal and enable Competency Graph as the educational brain.

**Implementation**:
- Every data entity must define its contribution to competency
- Data structures must support Competency Graph updates
- Competency traceability is maintained for all data
- No data entity exists without competency impact

**Examples**:
- Learning activity data includes competency dimension mapping
- Assessment evidence includes competency evaluation
- Student data includes competency profile

## Principle 4: AI Ready Data

**Statement**: Data structures are designed for AI consumption, enabling automated analysis and decision-making.

**Rationale**: AI requires structured, high-quality data to function effectively. Data architecture must support AI requirements for training, inference, and continuous learning.

**Implementation**:
- Data structures follow AI-friendly schemas
- Data quality standards ensure reliable AI operation
- Domain events provide operational data for AI training
- Data is enriched with metadata for AI context

**Examples**:
- Learning activity data includes engagement metrics for AI analysis
- Assessment data includes rubric scores for AI evaluation
- Student data includes learning preferences for AI personalization

## Principle 5: Human Governed Data

**Statement**: Critical data requires human validation and oversight to maintain educational judgment.

**Rationale**: AI cannot replace human educational judgment. Data governance ensures that critical decisions remain under human control and that data privacy and security are maintained.

**Implementation**:
- Critical data requires human validation before use
- Data governance defines validation requirements
- Human oversight maintains educational values
- Privacy and security protect student data

**Examples**:
- Assessment grades require teacher validation
- Student progress reports require teacher review
- Competency evaluations require human confirmation

---

# SECTION 3 — Education Canonical Data Model

## Overview

The Education Canonical Data Model defines the core data entities that capture educational activities and outcomes. These entities are derived from the domains defined in 01_EDUCATION_DOMAIN_MODEL.md and the processes defined in 03_BUSINESS_PROCESS_ARCHITECTURE.md.

## Core Data Entities

### Student

**Purpose**: Capture student identity, enrollment, and educational journey

**Key Attributes**:
- Student ID (unique identifier)
- NISN (national student number)
- Name
- Date of Birth
- Gender
- Enrollment Status
- Current Phase (Fase Fondasi, A, B, C, D, E, F)
- Current School
- Parent/Guardian Information
- Special Needs
- Learning Preferences
- Competency Profile (reference to Competency Graph)

**Relationships**:
- Belongs to School
- Has Parents/Guardians
- Participates in Learning Activities
- Completes Assessments
- Has Competency Graph
- Has Digital Twin
- Has Lifelong Learning Record

**Domain**: Student Development
**Primary Process**: P7 — Student Development

### Teacher

**Purpose**: Capture teacher identity, qualifications, and professional development

**Key Attributes**:
- Teacher ID (unique identifier)
- NIP (national teacher number)
- Name
- Qualifications
- Subject Expertise
- Teaching Experience
- Current School
- Professional Development Records
- Performance Evaluations

**Relationships**:
- Belongs to School
- Teaches Subjects
- Facilitates Learning Activities
- Conducts Assessments
- Has Professional Development Records

**Domain**: Teacher Development
**Primary Process**: P8 — Teacher Development

### School

**Purpose**: Capture school identity, characteristics, and quality metrics

**Key Attributes**:
- School ID (unique identifier)
- NPSN (national school number)
- Name
- Type (PAUD, SD, SMP, SMA, SMK)
- Level
- Location
- Principal
- Contact Information
- Quality Indicators
- SPMI Status
- Resources

**Relationships**:
- Has Students
- Has Teachers
- Has Parents
- Implements Curriculum
- Conducts Assessments
- Generates Reports

**Domain**: School Improvement
**Primary Process**: P10 — School Improvement

### Parent

**Purpose**: Capture parent identity and engagement in education

**Key Attributes**:
- Parent ID (unique identifier)
- Name
- Contact Information
- Relationship to Student
- Engagement Level
- Communication Preferences

**Relationships**:
- Has Children (Students)
- Communicates with School
- Receives Reports
- Participates in Learning Support

**Domain**: Parent Partnership
**Primary Process**: P9 — Parent Partnership

### Graduate Profile

**Purpose**: Define the 8 dimensions of the graduate profile and their developmental phases

**Key Attributes**:
- Profile ID (unique identifier)
- Dimension (Keimanan & Ketakwaan, Kewargaan, Penalaran Kritis, Kreativitas, Kolaborasi, Kemandirian, Kesehatan, Komunikasi)
- Phase (Fase Fondasi, A, B, C, D, E, F)
- Indicators
- Developmental Milestones
- Assessment Framework

**Relationships**:
- Informs Curriculum
- Guides Learning Activities
- Measured by Assessments
- Reflected in Competency Graph
- Communicated in Reports

**Domain**: Graduate Profile
**Primary Process**: P1 — Graduate Outcome Management

### Competency

**Purpose**: Define specific competencies and their relationships

**Key Attributes**:
- Competency ID (unique identifier)
- Name
- Description
- Dimension (reference to Graduate Profile)
- Level (Emerging, Developing, Proficient, Advanced)
- Prerequisites
- Related Competencies
- Evidence Requirements

**Relationships**:
- Part of Graduate Profile Dimension
- Has Prerequisites
- Related to Other Competencies
- Measured by Assessments
- Tracked in Competency Graph

**Domain**: Graduate Profile / Curriculum
**Primary Process**: P1 — Graduate Outcome Management, P2 — Curriculum Management

### Learning Activity

**Purpose**: Capture individual learning activities and their outcomes

**Key Attributes**:
- Activity ID (unique identifier)
- Type (Understand, Apply, Reflect)
- Subject
- Topic
- Date/Time
- Duration
- Participants (Students)
- Facilitator (Teacher)
- Resources Used
- Learning Objectives
- Outcomes
- Engagement Metrics
- Competency Impact (mapping to Competency)

**Relationships**:
- Conducted by Teacher
- Participated by Students
- Based on Modul Ajar
- Produces Assessment Evidence
- Updates Competency Graph
- Updates Digital Twin

**Domain**: Learning Delivery
**Primary Process**: P4 — Learning Delivery

### Assessment

**Purpose**: Define assessment instruments and their structure

**Key Attributes**:
- Assessment ID (unique identifier)
- Type (Formative, Summative, Diagnostic)
- Subject
- Topic
- Learning Objectives
- Questions/Items
- Rubric
- Scoring Criteria
- Competency Mapping
- Date/Time
- Administrator (Teacher)

**Relationships**:
- Created by Teacher
- Based on Learning Objectives
- Measures Competencies
- Produces Assessment Results
- Generates Evidence

**Domain**: Assessment
**Primary Process**: P5 — Assessment Management

### Evidence

**Purpose**: Capture evidence of learning and achievement

**Key Attributes**:
- Evidence ID (unique identifier)
- Type (Artifact, Observation, Performance, Portfolio)
- Source (Student, Teacher, System)
- Date/Time
- Description
- Artifact (file, image, video, text)
- Associated Assessment
- Associated Learning Activity
- Competency Evaluation
- Quality Rating

**Relationships**:
- Produced by Learning Activity
- Collected by Assessment
- Evaluates Competency
- Updates Competency Graph
- Part of Digital Twin

**Domain**: Assessment
**Primary Process**: P5 — Assessment Management

### Evaluation

**Purpose**: Capture evaluation of evidence against competency standards

**Key Attributes**:
- Evaluation ID (unique identifier)
- Evidence ID (reference)
- Competency ID (reference)
- Evaluator (Teacher/AI)
- Score
- Level (Emerging, Developing, Proficient, Advanced)
- Feedback
- Date/Time
- Validation Status

**Relationships**:
- Evaluates Evidence
- Assesses Competency
- Performed by Teacher or AI
- Updates Competency Graph
- Requires Human Validation

**Domain**: Assessment
**Primary Process**: P5 — Assessment Management

### Competency Graph

**Purpose**: Capture the complete competency development history and current state for each student

**Key Attributes**:
- Graph ID (unique identifier)
- Student ID (reference)
- Competency Nodes (with levels and relationships)
- Achievement History (timeline of competency development)
- Current State (snapshot of all competencies)
- Growth Trajectory (predicted development path)
- Last Updated

**Relationships**:
- Belongs to Student
- Contains Competency Nodes
- Updated by Evaluations
- Updated by Learning Activities
- Used by Digital Twin
- Used by AI Agents

**Domain**: Competency Intelligence
**Primary Process**: P12 — Competency Intelligence

### Digital Twin

**Purpose**: Capture the living, real-time representation of each student's educational journey

**Key Attributes**:
- Twin ID (unique identifier)
- Student ID (reference)
- Graduate Profile (current state across 8 dimensions)
- Competency Profile (current competency status)
- Learning History (all learning activities)
- Learning Preferences (personal learning style)
- Growth Trajectory (predicted learning path)
- Engagement Patterns
- Intervention History
- Privacy Settings
- Last Updated

**Relationships**:
- Belongs to Student
- Based on Competency Graph
- Updated by Learning Activities
- Updated by Assessments
- Used by AI Agents
- Used by Teachers
- Used by Parents

**Domain**: Competency Intelligence
**Primary Process**: P12 — Competency Intelligence

### Credential

**Purpose**: Capture formal credentials and achievements

**Key Attributes**:
- Credential ID (unique identifier)
- Type (Certificate, Diploma, Badge, Certification)
- Name
- Description
- Issuer
- Issue Date
- Expiration Date
- Verification Status
- Blockchain Hash (for verification)
- Associated Competencies
- Student ID (reference)

**Relationships**:
- Issued to Student
- Based on Competency Achievement
- Part of Lifelong Learning Record
- Verifiable through Blockchain

**Domain**: Credential & Achievement
**Primary Process**: P13 — Credential Management

### Achievement

**Purpose**: Capture informal achievements and milestones

**Key Attributes**:
- Achievement ID (unique identifier)
- Type (Competition, Project, Volunteer, Recognition)
- Name
- Description
- Date
- Issuer
- Evidence
- Associated Competencies
- Student ID (reference)

**Relationships**:
- Earned by Student
- Based on Competency Development
- Part of Lifelong Learning Record
- Recognized by System

**Domain**: Credential & Achievement
**Primary Process**: P13 — Credential Management

### Lifelong Learning Record

**Purpose**: Capture the complete learning identity across educational phases and lifelong learning

**Key Attributes**:
- Record ID (unique identifier)
- Student ID (reference)
- Learning Identity (unique lifelong identifier)
- Educational Phases (PAUD, SD, SMP, SMA, Higher Education, Workplace)
- Formal Learning Records
- Informal Learning Records
- Credentials
- Achievements
- Competency History
- Career Readiness Assessment
- Portability Status
- Privacy Settings

**Relationships**:
- Belongs to Student
- Aggregates Learning Activities
- Aggregates Credentials
- Aggregates Achievements
- Based on Competency Graph
- Used for Career Readiness

**Domain**: Lifelong Learning Record
**Primary Process**: P14 — Lifelong Learning Record

---

# SECTION 4 — Data Classification

## Overview

Data classification categorizes data entities based on their characteristics, usage patterns, and governance requirements. This classification enables appropriate data management strategies, security controls, and lifecycle policies.

## Classification Categories

### Master Data

**Definition**: Core business entities that change infrequently and are referenced by transactional data.

**Characteristics**:
- Long-lived
- Low volume of changes
- Referenced by multiple processes
- High business value

**Examples**:
- Student
- Teacher
- School
- Parent
- Graduate Profile
- Competency

**Governance**: Strict change control, high availability, data quality standards

### Reference Data

**Definition**: Static or semi-static data that provides context and classification for master and transactional data.

**Characteristics**:
- Relatively stable
- Used for lookup and validation
- Hierarchical or categorical
- Standardized across system

**Examples**:
- National Standards
- Phase Specifications
- Subject Classifications
- Assessment Types
- Competency Levels
- Graduate Profile Dimensions

**Governance**: Version-controlled, standardized definitions, periodic review

### Transaction Data

**Definition**: Data that captures business transactions and operational activities.

**Characteristics**:
- High volume
- Time-sensitive
- Event-driven
- Immutable (domain events)

**Examples**:
- Learning Activity Records
- Assessment Results
- Evidence Records
- Evaluation Records
- Attendance Records
- Communication Logs

**Governance**: Domain events, audit trail, retention policies, privacy controls

### Analytical Data

**Definition**: Aggregated and processed data used for reporting, analytics, and decision-making.

**Characteristics**:
- Derived from transaction data
- Aggregated and summarized
- Time-series analysis
- Multi-dimensional

**Examples**:
- Progress Reports
- Aggregate Assessment Analytics
- Competency Development Trends
- School Quality Indicators
- National Education Intelligence
- Predictive Models

**Governance**: Data quality validation, refresh schedules, access controls, privacy aggregation

### AI Data

**Definition**: Data specifically structured and prepared for AI consumption and machine learning.

**Characteristics**:
- Feature-engineered
- Labeled (for supervised learning)
- Enriched with metadata
- High-dimensional
- Time-series for temporal models

**Examples**:
- Learning Activity Features (engagement metrics, duration, participation)
- Assessment Features (rubric scores, difficulty, discrimination)
- Student Features (learning preferences, historical performance)
- Competency Graph Features (node embeddings, relationship weights)
- Digital Twin Features (state vectors, growth trajectories)

**Governance**: Feature versioning, model training data, drift monitoring, bias detection

---

# SECTION 5 — Domain Data Ownership Matrix

## Overview

The Domain Data Ownership Matrix defines which domains own which data entities and which domains consume which data. This ensures clear data ownership, accountability, and traceability.

## Ownership Matrix

| Domain | Owns Data | Consumes Data |
| ------- | --------- | ------------- |
| Graduate Profile | Graduate Profile, Competency | Curriculum, Learning Planning, Assessment, Reporting |
| Curriculum | CP, TP, ATP, Modul Ajar, Curriculum Mapping | Learning Planning, Learning Delivery, Assessment |
| Learning Planning | Lesson Plan, Resource Allocations, Differentiation Strategies | Learning Delivery, Assessment |
| Learning Delivery | Learning Activity Records, Session Logs, Engagement Metrics | Assessment, Competency Intelligence, Reporting |
| Deep Learning Operating Layer | Deep Learning Cycle Records, Metacognition Data | Learning Delivery, Assessment, Competency Intelligence |
| Assessment | Assessment Instruments, Assessment Results, Evidence, Evaluation | Competency Intelligence, Reporting, Digital Twin |
| Reporting | Progress Reports, Narrative Reports, Stakeholder Dashboards | Student Development, Teacher Development, School Improvement, Parent Partnership |
| Literacy | Literacy Assessment Data, Reading Progress, Writing Samples | Learning Delivery, Assessment, Competency Intelligence |
| Numeracy | Numeracy Assessment Data, Math Progress, Problem-Solving Samples | Learning Delivery, Assessment, Competency Intelligence |
| Coding | Coding Projects, Programming Assessments, Code Quality Metrics | Learning Delivery, Assessment, Competency Intelligence |
| AI Literacy | AI Projects, Prompt Engineering Assessments, AI Ethics Records | Learning Delivery, Assessment, Competency Intelligence |
| Character Development | Character Observations, Behavioral Records, Values Assessment | Learning Delivery, Assessment, Reporting |
| Student Wellbeing | Wellbeing Assessments, Mental Health Records, Physical Fitness Data | Learning Delivery, Assessment, Reporting, Parent Partnership |
| Parent Partnership | Communication Logs, Engagement Records, Feedback | Learning Delivery, Assessment, Reporting |
| Career & Future Readiness | Career Assessments, Interest Inventories, Work Readiness Records | Learning Planning, Assessment, Reporting, Lifelong Learning Record |
| Lifelong Learning Record | Learning Identity, Educational Phase Records, Credential Records | Credential Management, Career & Future Readiness, AI Orchestration |
| Student Development | Student Data, Progress Records, Development Plans | Learning Planning, Learning Delivery, Assessment, Reporting |
| Teacher Development | Teacher Data, Performance Records, Professional Development Plans | Learning Delivery, Assessment, Reporting |
| School Improvement | School Data, Quality Indicators, SPMI Records | Learning Delivery, Assessment, Reporting, Education Analytics |
| Education Analytics | Aggregate Analytics, Intelligence Reports, Predictive Models | All domains (consumes data from all) |
| Competency Intelligence | Competency Graph, Competency Evaluations, Competency Analytics | All domains (provides competency insights to all) |
| Credential & Achievement | Credentials, Achievements, Verification Records | Lifelong Learning Record, Career & Future Readiness |
| Governance & Compliance | Compliance Records, Policy Documents, Audit Logs | All domains (enforces governance across all) |

## Ownership Principles

### Single Owner
Each data entity has one primary owning domain responsible for its definition, quality, and lifecycle management.

### Shared Consumption
Multiple domains may consume the same data entity for different purposes, but only the owning domain can modify the entity definition.

### Traceability
All data consumption must be traceable to a business process defined in 03_BUSINESS_PROCESS_ARCHITECTURE.md.

### Change Control
Changes to data entity definitions require approval from the owning domain and impact analysis on consuming domains.

---

# SECTION 6 — Aggregate Root Design

## Overview

Aggregate Root Design follows Domain-Driven Design (DDD) principles to define consistency boundaries around related entities. Each aggregate has one root entity that ensures consistency within the aggregate boundary.

## Aggregate Roots

### Student Aggregate

**Root Entity**: Student

**Contained Entities**:
- Student (root)
- Enrollment Records
- Learning Preferences
- Special Needs
- Parent/Guardian Relationships

**Consistency Boundary**: All data related to a single student's identity and enrollment

**Invariants**:
- Student must have unique Student ID
- Student must be enrolled in exactly one school at a time
- Student must have at least one parent/guardian

**Business Operations**:
- Enroll Student
- Transfer Student
- Update Student Information
- Manage Parent/Guardian Relationships

**Domain**: Student Development
**Primary Process**: P7 — Student Development

### Curriculum Aggregate

**Root Entity**: CP (Capaian Pembelajaran)

**Contained Entities**:
- CP (root)
- TP (Tujuan Pembelajaran)
- ATP (Alur Tujuan Pembelajaran)
- Modul Ajar
- Curriculum Mapping
- Alignment Records

**Consistency Boundary**: All data related to curriculum definition and structure

**Invariants**:
- CP must be defined for each phase and subject
- TP must derive from CP
- ATP must sequence TP
- Modul Ajar must implement ATP

**Business Operations**:
- Define CP
- Generate TP from CP
- Sequence ATP from TP
- Generate Modul Ajar from ATP
- Align Curriculum with Graduate Profile

**Domain**: Curriculum
**Primary Process**: P2 — Curriculum Management

### Assessment Aggregate

**Root Entity**: Assessment

**Contained Entities**:
- Assessment (root)
- Assessment Items/Questions
- Rubric
- Scoring Criteria
- Assessment Results
- Evidence Records
- Evaluation Records

**Consistency Boundary**: All data related to a single assessment instrument and its execution

**Invariants**:
- Assessment must have at least one item
- Assessment must have scoring criteria
- Assessment results must be linked to assessment instrument
- Evidence must be linked to assessment

**Business Operations**:
- Design Assessment
- Execute Assessment
- Collect Evidence
- Evaluate Evidence
- Generate Assessment Results

**Domain**: Assessment
**Primary Process**: P5 — Assessment Management

### Competency Graph Aggregate

**Root Entity**: Competency Graph

**Contained Entities**:
- Competency Graph (root)
- Competency Nodes
- Competency Relationships
- Achievement History
- Growth Trajectory

**Consistency Boundary**: All data related to a single student's competency development

**Invariants**:
- Competency Graph must belong to exactly one student
- Competency nodes must be unique
- Relationships must be valid (prerequisites must exist)
- Achievement history must be chronological

**Business Operations**:
- Initialize Competency Graph
- Update Competency Node
- Add Competency Relationship
- Record Achievement
- Predict Growth Trajectory

**Domain**: Competency Intelligence
**Primary Process**: P12 — Competency Intelligence

### Digital Twin Aggregate

**Root Entity**: Digital Twin

**Contained Entities**:
- Digital Twin (root)
- Graduate Profile State
- Competency Profile State
- Learning History
- Learning Preferences
- Growth Trajectory
- Engagement Patterns
- Intervention History

**Consistency Boundary**: All data related to a single student's digital twin representation

**Invariants**:
- Digital Twin must belong to exactly one student
- Digital Twin must be synchronized with Competency Graph
- Digital Twin must update in real-time
- Privacy settings must be respected

**Business Operations**:
- Initialize Digital Twin
- Update Graduate Profile State
- Update Competency Profile State
- Record Learning Activity
- Update Growth Trajectory
- Generate Interventions

**Domain**: Competency Intelligence
**Primary Process**: P12 — Competency Intelligence

### Learning Record Aggregate

**Root Entity**: Lifelong Learning Record

**Contained Entities**:
- Lifelong Learning Record (root)
- Educational Phase Records
- Formal Learning Records
- Informal Learning Records
- Credentials
- Achievements
- Competency History
- Career Readiness Assessment

**Consistency Boundary**: All data related to a single student's lifelong learning identity

**Invariants**:
- Lifelong Learning Record must have unique Learning Identity
- Learning Identity must persist across educational phases
- Credentials must be verifiable
- Achievements must be recognized

**Business Operations**:
- Initialize Lifelong Learning Record
- Record Educational Phase
- Aggregate Formal Learning
- Aggregate Informal Learning
- Issue Credential
- Record Achievement
- Assess Career Readiness

**Domain**: Lifelong Learning Record
**Primary Process**: P14 — Lifelong Learning Record

---

# SECTION 7 — Competency Graph Data Architecture

## Overview

The Competency Graph is the educational brain of the system. It captures the complete competency development history and current state for each student, enabling AI to understand learning progress and provide personalized recommendations.

## Competency Graph as Educational Brain

### Graph Structure

The Competency Graph is a directed graph where:
- **Nodes**: Individual competencies with current achievement levels
- **Edges**: Relationships between competencies (prerequisites, enhancements, transfers)
- **Node Attributes**: Competency ID, name, description, dimension, level, achievement date, evidence
- **Edge Attributes**: Relationship type, strength, confidence

### Graph Properties

**Dynamic**: The graph evolves continuously as students learn and develop competencies

**Personalized**: Each student has their own unique competency graph based on their learning journey

**Traceable**: Every competency achievement is traceable to specific learning activities and assessments

**Predictive**: The graph enables prediction of future learning pathways and competency development

**Explainable**: AI can explain recommendations based on graph structure and relationships

---

# Audit Field Standard

All transactional entities should support audit fields for traceability and governance.

## Mandatory Fields

All transactional entities must include:

- **created_at**: Timestamp when the record was created
- **updated_at**: Timestamp when the record was last updated

## Recommended Fields

Transactional entities should include where operationally justified:

- **created_by**: User ID of the user who created the record
- **updated_by**: User ID of the user who last updated the record

## Optional Fields

Transactional entities may include for soft delete functionality:

- **deleted_at**: Timestamp when the record was soft deleted
- **deleted_by**: User ID of the user who soft deleted the record

## Approval-Based Entities

Entities that require human approval (artifacts, configurations, policies) additionally support:

- **approved_by**: User ID of the user who approved the record
- **approved_at**: Timestamp when the record was approved

## Implementation Guidance

MVP implementation may introduce audit fields incrementally where operationally justified.

Do not force unnecessary schema complexity. Apply audit fields based on actual governance and traceability requirements.

The centralized audit_logs table provides system-wide audit trail regardless of individual entity audit field implementation.

## Data Model

### Competency Graph Entity

```
CompetencyGraph {
  graphId: UUID
  studentId: UUID
  nodes: CompetencyNode[]
  edges: CompetencyEdge[]
  achievementHistory: AchievementEvent[]
  currentSnapshot: GraphSnapshot
  growthTrajectory: GrowthPrediction
  lastUpdated: Timestamp
}
```

### Competency Node

```
CompetencyNode {
  competencyId: UUID
  name: String
  description: String
  dimension: GraduateProfileDimension
  currentLevel: CompetencyLevel
  targetLevel: CompetencyLevel
  achievementDate: Timestamp
  evidence: Evidence[]
  confidence: Float
  trend: TrendDirection
}
```

### Competency Edge

```
CompetencyEdge {
  sourceNodeId: UUID
  targetNodeId: UUID
  relationshipType: RelationshipType
  strength: Float
  confidence: Float
}
```

### Achievement Event

```
AchievementEvent {
  eventId: UUID
  competencyId: UUID
  previousLevel: CompetencyLevel
  newLevel: CompetencyLevel
  timestamp: Timestamp
  source: Source (LearningActivity, Assessment, Evaluation)
  evidence: EvidenceId
  validated: Boolean
}
```

## Graph Operations

### Initialization

Create initial competency graph for a new student based on:
- Graduate Profile dimensions
- Phase-specific competency expectations
- Prior learning records (if any)
- Developmental milestones

### Update

Update competency graph when:
- Learning activity completed
- Assessment evidence collected
- Evaluation performed
- Competency achieved
- Competency level improved

### Query

Query competency graph for:
- Current competency status
- Competency relationships
- Learning pathways
- Prerequisites for target competencies
- Recommended next competencies

### Prediction

Predict competency development based on:
- Historical achievement patterns
- Learning velocity
- Competency relationships
- Growth trajectory models

## AI Integration

### Graph Embeddings

Generate graph embeddings for:
- Competency similarity analysis
- Learning pathway recommendation
- Student clustering
- Predictive modeling

### Graph Neural Networks

Apply graph neural networks for:
- Competency achievement prediction
- Learning pathway optimization
- Personalized recommendation generation
- Anomaly detection

### Natural Language Processing

Use NLP for:
- Competency description understanding
- Evidence text analysis
- Feedback generation
- Report generation

## Data Quality

### Validation

Validate competency graph data for:
- Node consistency (no duplicate nodes)
- Edge validity (prerequisites must exist)
- Level progression (must follow valid sequence)
- Evidence linkage (every achievement must have evidence)

### Consistency

Ensure competency graph consistency with:
- Learning activity records
- Assessment results
- Evaluation records
- Digital twin state

### Completeness

Ensure competency graph completeness for:
- All graduate profile dimensions
- All phase-specific competencies
- All prerequisite relationships
- Complete achievement history

---

# SECTION 8 — Digital Twin Data Architecture

## Overview

The Digital Twin is a living, real-time representation of each student's educational journey. It captures the current state of the student's development, learning history, preferences, and predicted growth trajectory.

## Digital Twin as Living Representation

### State Model

The Digital Twin maintains a comprehensive state model including:

**Graduate Profile State**: Current achievement status across 8 dimensions
- Dimension level (Emerging, Developing, Proficient, Advanced)
- Progress indicators
- Trend analysis
- Comparison with phase expectations

**Competency Profile State**: Current competency status across all competencies
- Competency levels
- Achievement history
- Growth velocity
- Gap analysis

**Learning History**: Complete record of all learning activities
- Learning activities completed
- Assessments taken
- Evidence collected
- Feedback received
- Interventions applied

**Learning Preferences**: Personal learning characteristics
- Learning style (visual, auditory, kinesthetic)
- Engagement patterns
- Motivation factors
- Preferred content types
- Optimal learning times

**Growth Trajectory**: Predicted learning path and outcomes
- Predicted competency achievement
- Recommended learning activities
- Intervention needs
- Timeline predictions

**Engagement Patterns**: Behavioral patterns and engagement metrics
- Participation rates
- Attention spans
- Collaboration preferences
- Self-regulation indicators

**Intervention History**: Record of all interventions applied
- Intervention type
- Intervention date
- Intervention outcome
- Effectiveness assessment

## Data Model

### Digital Twin Entity

```
DigitalTwin {
  twinId: UUID
  studentId: UUID
  graduateProfileState: GraduateProfileState
  competencyProfileState: CompetencyProfileState
  learningHistory: LearningHistory
  learningPreferences: LearningPreferences
  growthTrajectory: GrowthTrajectory
  engagementPatterns: EngagementPatterns
  interventionHistory: InterventionHistory
  privacySettings: PrivacySettings
  lastUpdated: Timestamp
}
```

### Graduate Profile State

```
GraduateProfileState {
  dimensions: DimensionState[]
  overallProgress: Float
  trend: TrendDirection
  phaseComparison: PhaseComparison
}
```

### Dimension State

```
DimensionState {
  dimension: GraduateProfileDimension
  currentLevel: CompetencyLevel
  targetLevel: CompetencyLevel
  progress: Float
  trend: TrendDirection
  evidence: Evidence[]
  lastUpdated: Timestamp
}
```

### Competency Profile State

```
CompetencyProfileState {
  competencies: CompetencyState[]
  overallCompetencyLevel: CompetencyLevel
  growthVelocity: Float
  gaps: CompetencyGap[]
}
```

### Competency State

```
CompetencyState {
  competencyId: UUID
  currentLevel: CompetencyLevel
  targetLevel: CompetencyLevel
  progress: Float
  lastAchieved: Timestamp
  evidence: Evidence[]
  prerequisites: PrerequisiteStatus[]
}
```

### Learning History

```
LearningHistory {
  activities: LearningActivityRecord[]
  assessments: AssessmentRecord[]
  evidence: EvidenceRecord[]
  feedback: FeedbackRecord[]
  interventions: InterventionRecord[]
}
```

### Learning Preferences

```
LearningPreferences {
  learningStyle: LearningStyle
  engagementPatterns: EngagementPattern[]
  motivationFactors: MotivationFactor[]
  preferredContentTypes: ContentType[]
  optimalLearningTimes: TimeOfDay[]
}
```

### Growth Trajectory

```
GrowthTrajectory {
  predictedCompetencies: PredictedCompetency[]
  recommendedActivities: RecommendedActivity[]
  interventionNeeds: InterventionNeed[]
  timeline: TimelinePrediction
  confidence: Float
}
```

## Update Mechanism

### Real-Time Updates

Digital Twin updates in real-time when:
- Learning activity completed
- Assessment evidence collected
- Evaluation performed
- Competency achieved
- Intervention applied

### Event-Driven Updates

Digital Twin updates are triggered by events:
- LearningActivityCompleted event
- AssessmentResulted event
- CompetencyAchieved event
- InterventionApplied event
- PreferenceUpdated event

### Synchronization

Digital Twin synchronizes with:
- Competency Graph (competency state)
- Learning Activity Records (learning history)
- Assessment Results (assessment history)
- Evaluation Records (competency evaluations)

## AI Foundation

### Personalization

Digital Twin enables AI personalization for:
- Learning activity recommendations
- Assessment adaptation
- Intervention targeting
- Content matching
- Learning pathway optimization

### Prediction

Digital Twin enables AI prediction for:
- Competency achievement timing
- Learning outcomes
- Intervention effectiveness
- Career readiness
- At-risk identification

### Explanation

Digital Twin enables AI explanation for:
- Recommendation rationale
- Prediction basis
- Intervention justification
- Progress explanation
- Trend analysis

## Privacy and Governance

### Privacy Settings

Digital Twin includes privacy settings for:
- Data sharing preferences
- Access control
- Anonymization options
- Data retention preferences

### Human Oversight

Digital Twin requires human oversight for:
- Intervention decisions
- At-risk identification
- Sensitive recommendations
- Privacy violations
- Anomaly detection

### Data Protection

Digital Twin data protection includes:
- Encryption at rest and in transit
- Access logging
- Data minimization
- Purpose limitation
- Storage limitation

---

# SECTION 9 — Lifelong Learning Record Architecture

## Overview

The Lifelong Learning Record captures the complete learning identity of each student across educational phases and lifelong learning. It provides a portable, verifiable record of learning achievements, credentials, and competencies.

## Lifelong Learning Record as Continuous Identity

### Learning Identity

The Lifelong Learning Record establishes a unique learning identity that:
- Persists across educational phases (PAUD, SD, SMP, SMA, Higher Education, Workplace)
- Aggregates formal and informal learning
- Recognizes credentials and achievements
- Supports career readiness
- Enables portability

### Educational Phases

The record captures learning across phases:
- **PAUD**: Early childhood development, foundational competencies
- **SD**: Primary education, basic competencies
- **SMP**: Secondary education, intermediate competencies
- **SMA**: High school education, advanced competencies
- **Higher Education**: University and vocational education
- **Workplace**: On-the-job learning, professional development

### Learning Types

The record aggregates:
- **Formal Learning**: Structured education in accredited institutions
- **Informal Learning**: Self-directed learning, online courses, workshops
- **Workplace Learning**: On-the-job training, professional development
- **Experiential Learning**: Projects, volunteering, competitions

## Data Model

### Lifelong Learning Record Entity

```
LifelongLearningRecord {
  recordId: UUID
  studentId: UUID
  learningIdentity: UUID (unique lifelong identifier)
  educationalPhases: EducationalPhaseRecord[]
  formalLearningRecords: FormalLearningRecord[]
  informalLearningRecords: InformalLearningRecord[]
  credentials: Credential[]
  achievements: Achievement[]
  competencyHistory: CompetencyHistory
  careerReadinessAssessment: CareerReadinessAssessment
  portabilityStatus: PortabilityStatus
  privacySettings: PrivacySettings
  lastUpdated: Timestamp
}
```

### Educational Phase Record

```
EducationalPhaseRecord {
  phaseId: UUID
  phaseType: PhaseType (PAUD, SD, SMP, SMA, HigherEducation, Workplace)
  institution: Institution
  startDate: Date
  endDate: Date
  learningRecords: LearningRecord[]
  competenciesAchieved: Competency[]
  credentialsEarned: Credential[]
  achievements: Achievement[]
}
```

### Formal Learning Record

```
FormalLearningRecord {
  recordId: UUID
  institution: Institution
  program: String
  startDate: Date
  endDate: Date
  courses: Course[]
  grades: Grade[]
  credits: Credit
  competenciesAchieved: Competency[]
  credentialsEarned: Credential[]
}
```

### Informal Learning Record

```
InformalLearningRecord {
  recordId: UUID
  provider: String
  program: String
  completionDate: Date
  duration: Duration
  competenciesDeveloped: Competency[]
  evidence: Evidence[]
  verification: VerificationStatus
}
```

### Competency History

```
CompetencyHistory {
  competencyId: UUID
  name: String
  achievementTimeline: AchievementEvent[]
  currentLevel: CompetencyLevel
  growthTrajectory: GrowthTrajectory
  sources: Source[]
}
```

### Career Readiness Assessment

```
CareerReadinessAssessment {
  assessmentId: UUID
  assessmentDate: Date
  careerInterests: CareerInterest[]
  skillAlignment: SkillAlignment[]
  readinessScore: Float
  recommendations: Recommendation[]
  gaps: SkillGap[]
}
```

## Credential Architecture

### Credential Types

- **Certificate**: Completion of a course or program
- **Diploma**: Completion of a degree program
- **Badge**: Achievement of a specific skill or competency
- **Certification**: Professional certification from industry body

### Credential Verification

Credentials are verified through:
- Blockchain-based verification
- Digital signatures
- Issuer verification
- Competency validation

### Credential Portability

Credentials are portable through:
- Standardized format (Verifiable Credentials)
- Blockchain-based verification
- Cross-institution recognition
- Industry acceptance

## Achievement Architecture

### Achievement Types

- **Competition**: Participation and achievement in competitions
- **Project**: Completion of significant projects
- **Volunteer**: Volunteer service and community engagement
- **Recognition**: Awards and honors

### Achievement Recognition

Achievements are recognized through:
- System validation
- Teacher verification
- Institution endorsement
- Industry recognition

## Data Integration

### Phase Transitions

Data continuity across phase transitions:
- Competency history preserved
- Learning records aggregated
- Credentials recognized
- Achievements acknowledged

### Cross-Phase Learning

Recognition of learning across phases:
- Prior learning recognition
- Credit transfer
- Competency validation
- Achievement acknowledgment

### Workplace Integration

Integration with workplace learning:
- Competency mapping to job requirements
- Credential verification
- Achievement recognition
- Learning record portability

## Privacy and Security

### Data Privacy

Lifelong Learning Record privacy includes:
- Student consent for data sharing
- Purpose limitation
- Data minimization
- Right to be forgotten

### Data Security

Lifelong Learning Record security includes:
- Encryption at rest and in transit
- Access control
- Audit logging
- Blockchain verification

### Data Governance

Lifelong Learning Record governance includes:
- Data ownership
- Access rights
- Retention policies
- Deletion procedures

---

# SECTION 10 — Data Lifecycle Architecture

## Overview

Data Lifecycle Architecture defines how data entities are created, updated, archived, retained, and deleted throughout their lifecycle. This ensures proper data management, compliance with regulations, and efficient resource utilization.

## Lifecycle Stages

### Create

**Purpose**: Create new data entities when business processes generate new information

**Triggers**:
- Student enrollment
- Curriculum definition
- Learning activity execution
- Assessment administration
- Credential issuance

**Validation**:
- Data quality checks
- Business rule validation
- Referential integrity
- Human validation (for critical data)

**Examples**:
- Create Student record when student enrolls
- Create CP when curriculum is defined
- Create Learning Activity when learning occurs
- Create Assessment when assessment is administered
- Create Credential when credential is issued

### Update

**Purpose**: Modify existing data entities when information changes

**Triggers**:
- Student information changes
- Curriculum revisions
- Learning activity updates
- Assessment result updates
- Competency achievement updates

**Validation**:
- Data quality checks
- Business rule validation
- Change authorization
- Audit logging

**Examples**:
- Update Student when information changes
- Update CP when curriculum is revised
- Update Learning Activity when additional data is collected
- Update Assessment when results are modified
- Update Competency Graph when competency is achieved

### Archive

**Purpose**: Move inactive or historical data to archive storage

**Triggers**:
- Student graduates or leaves school
- Academic year ends
- Curriculum version becomes obsolete
- Assessment period ends

**Policies**:
- Archive after defined retention period
- Compress archived data
- Maintain access for audit purposes
- Delete after legal retention period expires

**Examples**:
- Archive Student records after graduation
- Archive CP when curriculum version becomes obsolete
- Archive Assessment results after academic year ends
- Archive Learning Activity records after phase transition

### Retain

**Purpose**: Maintain data for legal, regulatory, or business requirements

**Retention Periods**:
- Student records: Until student reaches age 25 or 10 years after graduation (whichever is longer)
- Assessment records: 5 years
- Curriculum records: Indefinitely (for historical reference)
- Competency Graph: Indefinitely (lifelong learning record)
- Credentials: Indefinitely (lifelong learning record)

**Compliance**:
- PDPA (Personal Data Protection Act)
- Education regulations
- Institutional policies
- International standards

### Delete

**Purpose**: Permanently remove data when retention period expires or when requested by data subject

**Triggers**:
- Retention period expires
- Data subject requests deletion (right to be forgotten)
- Data is no longer needed for business purposes

**Policies**:
- Secure deletion (data wiping)
- Verification of deletion
- Audit logging of deletion
- Compliance with legal requirements

**Examples**:
- Delete Student records after retention period expires
- Delete Assessment records after retention period expires
- Delete personal data upon request (right to be forgotten)

## Lifecycle by Data Type

### Master Data Lifecycle

**Create**: When entity is first registered (student enrollment, teacher hiring, school establishment)

**Update**: When entity information changes (student transfer, teacher promotion, school accreditation)

**Archive**: When entity becomes inactive (student graduates, teacher retires, school closes)

**Retain**: Indefinitely for historical reference and audit

**Delete**: Only in exceptional circumstances with legal authorization

### Reference Data Lifecycle

**Create**: When reference data is defined (national standards, phase specifications)

**Update**: When reference data is revised (curriculum updates, standard changes)

**Archive**: When reference data version becomes obsolete

**Retain**: Indefinitely for historical reference

**Delete**: Never (reference data is never deleted, only archived)

### Transaction Data Lifecycle

**Create**: When business transaction occurs (learning activity, assessment, communication)

**Update**: Never (transaction data is immutable, captured as events)

**Archive**: After defined retention period (end of academic year, phase transition)

**Retain**: For defined retention period (5 years for assessments, 10 years for student records)

**Delete**: After retention period expires or upon data subject request

### Analytical Data Lifecycle

**Create**: When analytics are generated (reports, dashboards, predictions)

**Update**: When underlying data changes and analytics are refreshed

**Archive**: When analytics become obsolete (end of reporting period)

**Retain**: For defined retention period (1 year for most analytics, longer for trend analysis)

**Delete**: After retention period expires

### AI Data Lifecycle

**Create**: When AI features are engineered from transaction data

**Update**: When AI models are retrained with new data

**Archive**: When AI model version becomes obsolete

**Retain**: For model training history and audit

**Delete**: After model retention period expires or when data subject requests deletion

## Domain Event Strategy

### Domain Events

Domain Events are used for:
- Workflow Progression
- Internal Notifications
- Process Coordination

Examples:
- TPGenerated
- ATPGenerated
- AssessmentCompleted
- ReportGenerated

NUSA MVP Wave 1 does not implement Event Sourcing.
NUSA MVP Wave 1 does not implement Event Store infrastructure.
Domain Events are operational events and not an Event Sourcing architecture.

---

# SECTION 11 — Operational Domain Event Strategy

## Overview

Domain Events are used for workflow progression, internal notifications, and process coordination through RabbitMQ. Domain Events are operational integration events only and are NOT used as the source of truth. PostgreSQL transactional tables remain the source of truth.

## Event Design Principles

### Operational Events

Events are operational integration events:
- Events are published for state changes
- Events are consumed via RabbitMQ for workflow coordination
- Events enable asynchronous processing
- Events provide notifications to downstream modules

### Event Types

Define operational event types for:
- Curriculum events (TPGenerated, TPUpdated, TPApproved)
- Learning Planning events (ATPGenerated, ATPApproved, ModulAjarGenerated, ModulAjarApproved)
- Assessment events (AssessmentGenerated, RubricGenerated, EvidenceRecorded)
- Reporting events (NarrativeReportGenerated, NarrativeReportApproved)
- Administration events (UserCreated, UserRoleChanged)
- AI Orchestration events (AIGenerationRequested, AIGenerationCompleted, AIGenerationFailed)

### Event Structure

Each event includes:
- Event ID (unique identifier)
- Event Type
- Timestamp
- Producer Module
- Consumer Module
- Event Data (specific to event type)
- Correlation ID (for request tracing)

## Event Purpose

### Workflow Coordination

Events enable workflow progression:
- TPGenerated triggers ATP generation
- ATPApproved triggers Modul Ajar generation
- AssessmentGenerated triggers Rubric generation
- AIGenerationCompleted notifies completion

### Module Notifications

Events provide notifications:
- TPApproved notifies Learning Planning Module
- AssessmentCompleted notifies Reporting Module
- UserCreated notifies Administration Module

### Asynchronous Processing

Events enable asynchronous processing:
- Long-running AI operations
- Cross-module workflows
- Background processing tasks

## Event Infrastructure

### RabbitMQ Configuration

Events are published to:
- RabbitMQ topic exchange (nusa.events)
- Routing keys by module and event type
- Durable queues for reliability
- At-least-once delivery guarantee

### Event Consumption

Events are consumed by:
- Module-specific consumers
- Idempotent processing
- Error handling and retry
- Dead letter queue for failed events

## Event Governance

### Event Authorization

Event publishing requires:
- Module authorization
- User authentication (if applicable)
- Business rule validation
- Human validation (for critical events)

### Event Validation

Event validation includes:
- Schema validation
- Business rule validation
- Referential integrity
- Data quality checks

### Event Auditing

Event auditing provides:
- Complete audit trail in audit_logs table
- Change history
- Accountability
- Compliance verification

---

# SECTION 12 — Data Governance

## Overview

Data Governance establishes the policies, procedures, and standards for data management across the NUSA Education Platform. It ensures data quality, security, privacy, and compliance with regulations.

## Governance Framework

### Data Ownership

**Domain Owners**: Each domain owns the data entities it produces
- Responsible for data definition
- Responsible for data quality
- Responsible for data lifecycle
- Responsible for data access authorization

**Data Stewards**: Responsible for day-to-day data management
- Maintain data quality
- Resolve data issues
- Coordinate data changes
- Communicate with stakeholders

### Data Quality

**Quality Dimensions**:
- Accuracy: Data is correct and reliable
- Completeness: Data is complete and comprehensive
- Consistency: Data is consistent across systems
- Timeliness: Data is current and up-to-date
- Validity: Data conforms to business rules
- Uniqueness: Data is not duplicated

**Quality Controls**:
- Data validation rules
- Data quality checks
- Data cleansing procedures
- Data quality monitoring
- Data quality reporting

### Data Access Control

**Access Principles**:
- Least privilege
- Need-to-know
- Role-based access
- Context-aware access

**Access Types**:
- Read access
- Write access
- Delete access
- Admin access

**Access Authorization**:
- Role-based access control (RBAC)
- Attribute-based access control (ABAC)
- Context-aware access control
- Just-in-time access provisioning

### Data Privacy

**Privacy Principles**:
- Data minimization
- Purpose limitation
- Consent management
- Data subject rights

**Privacy Controls**:
- Anonymization
- Pseudonymization
- Encryption
- Access logging
- Privacy impact assessment

### Data Security

**Security Measures**:
- Encryption at rest and in transit
- Secure authentication
- Secure authorization
- Audit logging
- Intrusion detection
- Data breach response

### Data Compliance

**Regulatory Compliance**:
- PDPA (Personal Data Protection Act)
- Education regulations
- Child protection laws
- International standards (GDPR, if applicable)

**Compliance Measures**:
- Compliance monitoring
- Compliance reporting
- Compliance audits
- Risk assessment
- Remediation procedures

## Governance Processes

### Data Definition Process

1. Data entity proposed by domain owner
2. Data entity reviewed by architecture board
3. Data entity approved and documented
4. Data entity implemented in systems
5. Data entity monitored for quality

### Data Change Process

1. Data change requested
2. Impact analysis performed
3. Change approved by domain owner
4. Change implemented
5. Change validated
6. Stakeholders notified

### Data Issue Resolution Process

1. Data issue identified
2. Issue logged and prioritized
3. Root cause analysis performed
4. Resolution implemented
5. Resolution validated
6. Stakeholders notified

### Data Access Request Process

1. Access request submitted
2. Request reviewed by data owner
3. Access approved or denied
4. Access provisioned
5. Access monitored
6. Access reviewed periodically

## Governance Roles

### Data Governance Board

**Responsibilities**:
- Approve data standards
- Resolve data governance issues
- Monitor data governance compliance
- Establish data governance policies

### Data Owners

**Responsibilities**:
- Own data entities
- Define data quality standards
- Authorize data access
- Manage data lifecycle

### Data Stewards

**Responsibilities**:
- Maintain data quality
- Resolve data issues
- Coordinate data changes
- Communicate with stakeholders

### Data Custodians

**Responsibilities**:
- Implement data controls
- Monitor data systems
- Perform data backups
- Respond to data incidents

---

# SECTION 13 — Data Security & Privacy

## Overview

Data Security & Privacy establishes the technical and procedural controls to protect data from unauthorized access, use, disclosure, modification, or destruction.

## Security Architecture

### Defense in Depth

Multiple layers of security:
- Network security (firewalls, intrusion detection)
- Application security (authentication, authorization)
- Data security (encryption, access control)
- Physical security (data center security)

### Encryption

**Encryption at Rest**:
- Database encryption
- File system encryption
- Backup encryption
- Key management

**Encryption in Transit**:
- TLS/SSL for all communications
- API encryption
- Message encryption
- Key exchange protocols

### Authentication

**Authentication Methods**:
- Multi-factor authentication (MFA)
- Single sign-on (SSO)
- Biometric authentication
- Certificate-based authentication

**Authentication Policies**:
- Strong password requirements
- Password rotation
- Session timeout
- Failed login lockout

### Authorization

**Authorization Models**:
- Role-based access control (RBAC)
- Attribute-based access control (ABAC)
- Policy-based access control
- Just-in-time authorization

**Authorization Enforcement**:
- Least privilege
- Need-to-know
- Separation of duties
- Context-aware authorization

## Privacy Architecture

### Privacy by Design

Privacy principles embedded in:
- System design
- Data collection
- Data processing
- Data storage
- Data sharing

### Data Minimization

Collect only data that is:
- Necessary for business purpose
- Relevant to educational outcomes
- Proportionate to need
- Time-bound

### Purpose Limitation

Use data only for:
- Stated purpose
- Educational purposes
- Student benefit
- Compliance requirements

### Consent Management

Consent for:
- Data collection
- Data processing
- Data sharing
- Data retention
- Data deletion

### Data Subject Rights

Rights include:
- Right to access
- Right to rectification
- Right to erasure
- Right to restrict processing
- Right to data portability
- Right to object

## Privacy Controls

### Anonymization

Anonymization techniques:
- Data masking
- Pseudonymization
- Generalization
- Data swapping

### Pseudonymization

Pseudonymization methods:
- Token-based pseudonyms
- Hash-based pseudonyms
- Random pseudonyms
- Reversible pseudonyms (with key management)

### Data Masking

Masking techniques:
- Dynamic masking
- Static masking
- Partial masking
- Redaction

### Access Logging

Logging includes:
- Access attempts
- Access granted/denied
- Data viewed
- Data modified
- Data exported

## Incident Response

### Data Breach Response

Response procedures:
1. Identify breach
2. Contain breach
3. Notify stakeholders
4. Investigate breach
5. Remediate vulnerabilities
6. Document lessons learned

### Data Incident Response

Response procedures:
1. Identify incident
2. Assess impact
3. Notify stakeholders
4. Resolve incident
5. Prevent recurrence
6. Document incident

## Compliance

### Regulatory Compliance

Compliance with:
- PDPA (Personal Data Protection Act)
- Education regulations
- Child protection laws
- International standards

### Compliance Monitoring

Monitoring includes:
- Compliance audits
- Privacy impact assessments
- Security assessments
- Risk assessments
- Penetration testing

---

# SECTION 14 — Data Traceability Matrix

## Overview

The Data Traceability Matrix ensures complete traceability from national vision to data entities, enabling auditability, accountability, and alignment with strategic objectives.

## Traceability Chain

```
Indonesia Emas 2045
    ↓ defines
Graduate Profile (8 Dimensions)
    ↓ informs
Domain (01)
    ↓ translates to
Capability (02)
    ↓ enables
Business Process (03)
    ↓ generates
Events
    → updates
Data (04)
```

## Traceability Matrix

### Graduate Profile → Domain → Capability → Process → Data

| Graduate Profile Dimension | Domain | Capability | Process | Data Entity |
| ------------------------- | ------ | ---------- | ------- | ----------- |
| Keimanan & Ketakwaan | Graduate Profile | Graduate Profile Management | P1.1 Graduate Profile Definition | Graduate Profile |
| Kewargaan | Graduate Profile | Graduate Profile Management | P1.1 Graduate Profile Definition | Graduate Profile |
| Penalaran Kritis | Graduate Profile | Graduate Profile Management | P1.1 Graduate Profile Definition | Graduate Profile |
| Kreativitas | Graduate Profile | Graduate Profile Management | P1.1 Graduate Profile Definition | Graduate Profile |
| Kolaborasi | Graduate Profile | Graduate Profile Management | P1.1 Graduate Profile Definition | Graduate Profile |
| Kemandirian | Graduate Profile | Graduate Profile Management | P1.1 Graduate Profile Definition | Graduate Profile |
| Kesehatan | Graduate Profile | Graduate Profile Management | P1.1 Graduate Profile Definition | Graduate Profile |
| Komunikasi | Graduate Profile | Graduate Profile Management | P1.1 Graduate Profile Definition | Graduate Profile |

### Domain → Capability → Process → Data

| Domain | Capability | Process | Data Entity |
| ------ | ---------- | ------- | ----------- |
| Graduate Profile | Graduate Profile Management | P1.1 Graduate Profile Definition | Graduate Profile, Competency |
| Curriculum | Curriculum Management | P2.1 CP Management | CP, TP, ATP, Modul Ajar |
| Learning Planning | Learning Planning | P3.1 Learning Objective Design | Lesson Plan, Resource Allocations |
| Learning Delivery | Learning Delivery | P4.1 Learning Session Planning | Learning Activity Record, Session Log |
| Assessment | Assessment Management | P5.1 Assessment Design | Assessment, Assessment Result, Evidence |
| Reporting | Reporting Management | P6.1 Progress Report Generation | Progress Report, Narrative Report |
| Student Development | Student Development | P7.1 Student Progress Monitoring | Student, Progress Record |
| Teacher Development | Teacher Development | P8.1 Teacher Performance Monitoring | Teacher, Performance Record |
| Parent Partnership | Parent Partnership | P9.1 Communication Management | Parent, Communication Log |
| School Improvement | School Improvement | P10.1 Quality Indicator Monitoring | School, Quality Indicator |
| Education Analytics | Education Analytics | P11.1 Data Collection | Aggregate Analytics, Intelligence Report |
| Competency Intelligence | Competency Intelligence | P12.1 Competency Graph Update | Competency Graph, Digital Twin |
| Credential Management | Credential Management | P13.1 Credential Definition | Credential, Achievement |
| Lifelong Learning Record | Lifelong Learning Record | P14.1 Record Aggregation | Lifelong Learning Record |
| AI Orchestration | AI Orchestration | P15.1 Agent Coordination | AI Data, Model Metadata |
| Governance & Compliance | Governance & Compliance | P16.1 Compliance Monitoring | Compliance Record, Audit Log |

### Process → Data Production

| Process | Produces Data | Consumes Data |
| ------- | ------------- | ------------- |
| P1 — Graduate Outcome Management | Graduate Profile, Competency | National Standards, Cultural Values |
| P2 — Curriculum Management | CP, TP, ATP, Modul Ajar | Graduate Profile, National Standards |
| P3 — Learning Planning | Lesson Plan, Modul Ajar | CP, TP, ATP, Graduate Profile |
| P4 — Learning Delivery | Learning Activity Record, Session Log | Lesson Plan, Modul Ajar, Student Data |
| P5 — Assessment Management | Assessment, Assessment Result, Evidence | Learning Activity Data, Competency Framework |
| P6 — Reporting Management | Progress Report, Narrative Report | Assessment Results, Competency Data |
| P7 — Student Development | Progress Record, Development Plan | Student Data, Assessment Data |
| P8 — Teacher Development | Performance Record, PD Record | Teacher Data, Assessment Data |
| P9 — Parent Partnership | Communication Log, Engagement Record | Student Data, School Data |
| P10 — School Improvement | Quality Indicator, SPMI Record | School Data, Assessment Data |
| P11 — Education Analytics | Aggregate Analytics, Intelligence Report | All transaction data |
| P12 — Competency Intelligence | Competency Graph, Digital Twin | Assessment Data, Learning Activity Data |
| P13 — Credential Management | Credential, Achievement | Competency Data, Learning Record |
| P14 — Lifelong Learning Record | Lifelong Learning Record | All learning data, Credential data |
| P15 — AI Orchestration | AI Data, Model Metadata | All data for AI training |
| P16 — Governance & Compliance | Compliance Record, Audit Log | All system data |

## Traceability Validation

### Completeness Check

Every data entity must be traceable to:
- A business process that produces it
- A capability that enables the process
- A domain that owns the capability
- A graduate profile dimension that the domain supports

### Consistency Check

Traceability must be consistent:
- No orphaned data entities
- No data entities without process traceability
- No processes without capability traceability
- No capabilities without domain traceability

### Alignment Check

Traceability must align with:
- National education vision
- Graduate profile dimensions
- Educational outcomes
- Business objectives

---

# SECTION 15 — Architecture Validation

## Overview

Architecture Validation ensures that the Data Architecture meets all requirements and aligns with all foundation documents.

## Validation Checklist

### Alignment with Foundation Documents

✓ **Aligned with 00A_NATIONAL_EDUCATION_DIRECTION_2045.md**
- Data entities support Indonesia Emas 2045 vision
- Data structures enable national education intelligence
- Data lifecycle supports human capital development

✓ **Aligned with 00B_PRODUCT_VISION.md**
- Data architecture supports AI-Native NUSA Education Platform
- Data structures enable 90% AI assistance
- Data quality ensures reliable AI operation

✓ **Aligned with 00C_EDUCATION_OPERATING_SYSTEM_PRINCIPLES.md**
- Data architecture follows Curriculum-Centered principle
- Data structures prioritize Learning over Administration
- Data design supports Deep Learning pedagogy

✓ **Aligned with 01_EDUCATION_DOMAIN_MODEL.md**
- Data entities capture approved domains
- Data relationships reflect domain relationships
- No new domains introduced through data

✓ **Aligned with 02_CAPABILITY_MODEL.md**
- Data entities support approved capabilities
- Data structures enable capability execution
- No new capabilities introduced through data

✓ **Aligned with 03_BUSINESS_PROCESS_ARCHITECTURE.md**
- All data entities traceable to processes
- Data flows match process data flow matrix
- No data entities without process traceability

### Data Model Completeness

✓ **Canonical Data Model Complete**
- All required data entities defined
- All entity attributes specified
- All relationships defined
- All invariants specified

✓ **Aggregate Roots Complete**
- Student Aggregate defined
- Curriculum Aggregate defined
- Assessment Aggregate defined
- Competency Graph Aggregate defined
- Digital Twin Aggregate defined
- Learning Record Aggregate defined

### Strategic Support

✓ **Competency Graph Support**
- Competency Graph data architecture defined
- Competency Graph as educational brain established
- Competency Graph AI integration defined

✓ **Digital Twin Support**
- Digital Twin data architecture defined
- Digital Twin state model defined
- Digital Twin real-time update mechanism defined

✓ **Lifelong Learning Record Support**
- Lifelong Learning Record architecture defined
- Learning identity continuity established
- Credential portability defined

### Architecture Level Compliance

✓ **Data Architecture Level Maintained**
- Document focuses on data entities and structures
- No technical implementation details
- No database design specifications
- No service deployment details
- Document serves as foundation for Application Architecture

### Data Governance

✓ **Governance Framework Established**
- Data ownership defined
- Data quality standards defined
- Data access control defined
- Data privacy controls defined

### Data Security

✓ **Security Architecture Defined**
- Encryption strategy defined
- Authentication methods defined
- Authorization models defined
- Incident response defined

### Data Lifecycle

✓ **Lifecycle Architecture Defined**
- Create, Update, Archive, Retain, Delete defined
- Domain event strategy defined
- Retention policies defined
- Compliance requirements defined

## Validation Status

**Overall Status**: ✓ PASSED

The Data Architecture document:
- Is validated against all foundation documents
- Does not introduce new domains or capabilities
- Maintains appropriate Data Architecture level boundaries
- Provides complete data model definitions
- Addresses all strategic requirements
- Establishes comprehensive governance framework
- Serves as a solid foundation for downstream architectures

## Next Steps

This document is now ready to serve as the foundation for:
- 05_AI_ARCHITECTURE.md
- 06_APPLICATION_ARCHITECTURE.md

All downstream architecture documents must:
- Trace their services and agents to data entities defined in this document
- Maintain alignment with the principles and constraints established here
- Respect the Data Architecture level boundaries
- Support the strategic objectives of Indonesia Emas 2045

---

# SECTION 16 — Conclusion

## Strategic Positioning

The Data Architecture serves as the critical foundation that defines WHAT data entities and structures are needed to support the business processes defined in the Business Process Architecture. It translates business process requirements into data models that capture educational activities and outcomes.

## Architecture Translation Chain

This Data Architecture (04) is the official translation layer in the architecture hierarchy:

```
Business Process Architecture (03)
    ↓ generates
Events
    → updates
Data Architecture (04)
    → enables
AI Architecture (05)
    → implemented by
Application Architecture (06)
```

**Business Process Architecture (03)**: Defines HOW capabilities are executed through workflows that deliver educational value.

**Data Architecture (04)**: Defines WHAT data entities and structures are needed to support the processes.

**AI Architecture (05)**: Defines HOW AI agents use data to assist and automate the processes.

**Application Architecture (06)**: Defines WHAT applications and services implement the processes and manage the data.

## Single Source of Truth

This Data Architecture is the single source of truth for:

- **Data Entity Definition**: All data entities must be defined in this document
- **Data Structure**: All data structures must follow this specification
- **Data Classification**: All data must be classified according to this framework
- **Data Lifecycle**: All data lifecycle policies must follow this specification
- **Data Governance**: All data governance must follow this framework
- **Data Security**: All data security must follow this specification

No data entity, data structure, or data policy should exist in the system without being defined in this document.

## Foundation for Downstream Architectures

### AI Architecture (05)
AI agents are designed to consume data entities defined here. Every AI agent must be mapped to one or more data entities it uses.

### Application Architecture (06)
Applications and services are built to manage data entities defined here. Every application must be traceable to one or more data entities it manages.

### Integration Architecture
Integration points are defined by data entity boundaries. Every integration must serve data entity operations.

### Security Architecture
Security requirements are derived from data classification and governance. Every security control must protect data entities defined here.

## Strategic Impact

### AI-Native Foundation
Data architecture provides the foundation for 90% AI assistance by:
- Designing AI-ready data structures
- Enabling domain events for AI training
- Providing high-quality data for AI operation
- Supporting AI explainability through data traceability

### Competency Graph Intelligence (FUTURE STRATEGIC DOMAIN)
**Status**: FUTURE STRATEGIC DOMAIN - Excluded from MVP Wave 1
Data architecture enables Competency Graph as educational brain by:
- Designing graph data structures
- Supporting real-time graph updates
- Enabling graph-based AI algorithms
- Providing graph explainability

### Digital Twin Personalization (FUTURE STRATEGIC DOMAIN)
**Status**: FUTURE STRATEGIC DOMAIN - Excluded from MVP Wave 1
Data architecture enables Digital Twin personalization by:
- Designing state model data structures
- Supporting real-time twin updates
- Enabling predictive analytics
- Providing personalized recommendations

### Lifelong Learning Continuity (FUTURE STRATEGIC DOMAIN)
**Status**: FUTURE STRATEGIC DOMAIN - Excluded from MVP Wave 1
Data architecture enables lifelong learning continuity by:
- Designing portable credential structures
- Supporting cross-phase data aggregation
- Enabling learning identity persistence
- Providing career readiness insights

## Continuous Evolution

The Data Architecture is designed for continuous evolution:

- **MVP Phase**: Implement core data entities (Student, Teacher, School, Curriculum, Assessment, Competency Graph)
- **Expansion Phase**: Add supporting data entities (Digital Twin, Lifelong Learning Record, Credential)
- **Maturity Phase**: Add advanced data structures (AI Data, Analytical Data, Event Store)
- **Optimization Phase**: Continuous data optimization based on usage and performance

## MVP Scope Protection

**MVP Wave 1 Data Scope is Strictly Limited To**:
- Core data entities for Graduate Profile, Curriculum, Learning Planning, Assessment, and Reporting
- Domain events for internal communication
- Basic data structures for AI agent operation

**Explicitly Excluded from MVP Wave 1**:
- Competency Graph data structures (FUTURE STRATEGIC DOMAIN)
- Digital Twin data structures (FUTURE STRATEGIC DOMAIN)
- Lifelong Learning Record data structures (FUTURE STRATEGIC DOMAIN)
- Event Store infrastructure
- Advanced analytical data structures

No future strategic data structure shall be included in MVP Wave 1 implementation without explicit architecture freeze amendment approved by Chief Enterprise Architect.

## Governance and Maintenance

### Architecture Governance
- Data Architecture Owner maintains this document
- Changes require approval from Architecture Review Board
- Alignment with foundation documents must be maintained

### Version Control
- This document is version-controlled
- Changes are tracked with rationale and approval
- Impact analysis is required for changes

### Stakeholder Communication
- Data changes are communicated to all stakeholders
- Training is provided for data changes
- Feedback is collected and incorporated

## Conclusion

The Data Architecture is the foundation upon which the entire NUSA Education Platform's data layer is built. It ensures that:

- Every data entity serves educational value
- Every data structure supports competency development
- Every data flow enables AI assistance
- Every data control maintains human authority
- Every data event enables real-time responsiveness
- Every data lifecycle ensures compliance

By adhering to this Data Architecture, the NUSA Education Platform will achieve its vision of building an AI-Native NUSA Education Platform that reduces teacher administrative burden, improves learning quality, strengthens student competency, and delivers the national education outcomes required for Indonesia Emas 2045.

---

**Document Status**: FOUNDATION DOCUMENT - LOCKED

**Version**: 1.0
**Freeze Date**: June 2026
**Next Review**: June 2027

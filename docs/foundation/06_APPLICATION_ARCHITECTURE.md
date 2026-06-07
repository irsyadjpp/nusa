# 06_APPLICATION_ARCHITECTURE.md

## Foundation Document for Education Operating System Indonesia 2045

**Version**: 2.0
**Date**: June 2026
**Status**: FOUNDATION DOCUMENT
**Alignment**: Validated against Foundation Architecture (00A, 00B, 00C, 01, 02, 03, 05)

**Purpose**: Define the complete application architecture for NUSA using Domain-Driven Design (DDD) principles, serving as the official translation from AI Architecture to Application Architecture. This document is the single source of truth for Bounded Contexts, Context Maps, Modules, Event-Driven Architecture, APIs, Integrations, AI Integration, Identity & Access, Notifications, Analytics, and Cross-Cutting Modules.

---

# SECTION 1 — Executive Summary

## Why Application Architecture is Required

Application Architecture is the critical layer that translates business domains into software boundaries using Domain-Driven Design (DDD) principles. It defines Bounded Contexts, Context Maps, Modules, Event-Driven Architecture, APIs, and Cross-Cutting Modules that implement the business processes defined in the Business Process Architecture (03), host the AI agents defined in the AI Architecture (05), and manage the data entities defined in the Data Architecture (04).

### Architecture Translation Chain

```
Domain Model (01)
    → bounded by
Application Architecture (06)
    → implemented by
SDLC Architecture (08)
```

**Domain Model (01)**: Defines the business domains and their relationships.

**Application Architecture (06)**: Defines the software boundaries (Bounded Contexts) and their interactions.

**SDLC Architecture (08)**: Defines the development methodology to implement the architecture.

### Primary Objectives

The Application Architecture is designed to build an AI-Native Education Operating System capable of:

#### Bounded Context Definition
- Define clear software boundaries aligned with business domains
- Establish Context Maps for inter-context communication
- Implement Anti-Corruption Layers (ACL) for external integration
- Define Shared Kernel for common domain concepts

#### Module Architecture
- Design modules aligned with Bounded Contexts
- Enable modular development within a single deployable application
- Support internal communication patterns
- Implement module-to-module integration patterns

#### Event-Driven Architecture
- Define domain events for state changes
- Enable real-time responsiveness
- Support eventual consistency

#### AI Integration
- Integrate AI agents within Bounded Contexts
- Enable AI agent orchestration across contexts
- Implement AI governance and monitoring

### Relationship with Foundation Documents

This Application Architecture is derived from and validated against:

- **00A_NATIONAL_EDUCATION_DIRECTION_2045.md**: The strategic direction and national vision for education transformation
- **00B_PRODUCT_VISION.md**: The product vision for AI-Native NUSA with 90% AI assistance
- **00C_EDUCATION_OPERATING_SYSTEM_PRINCIPLES.md**: The architectural principles (Curriculum-Centered, Learning > Administration)
- **01_EDUCATION_DOMAIN_MODEL.md**: The domain entities and relationships that define Bounded Contexts
- **02_CAPABILITY_MODEL.md**: The capabilities that modules implement
- **03_BUSINESS_PROCESS_ARCHITECTURE.md**: The business processes that modules execute
- **05_AI_ARCHITECTURE.md**: The AI agents that modules host

All Bounded Contexts defined in this document are traceable to domains defined in 01, which are traceable to capabilities defined in 02, which are traceable to processes defined in 03, which host AI agents defined in 05.

### Core Principle

**Bounded Contexts are the software boundaries that encapsulate business domains and their interactions.**

Every Bounded Context must be traceable to one or more domains defined in 01. If a Bounded Context is not defined here, it should not exist in the system.

---

# SECTION 2 — Architecture Principles

## Principle 1: Bounded Context Alignment

**Statement**: Bounded Contexts must align with business domains defined in the Domain Model (01).

**Rationale**: DDD requires clear software boundaries that encapsulate business domains. Misaligned boundaries lead to confusion, coupling, and maintenance issues.

**Implementation**:
- Each Bounded Context maps to one or more related domains
- Context boundaries are explicit and enforced
- Ubiquitous Language is consistent within each context
- Context mapping defines inter-context relationships

**Examples**:
- Curriculum Bounded Context aligns with Curriculum Domain
- Assessment Bounded Context aligns with Assessment Domain
- Competency Intelligence Bounded Context aligns with Graduate Profile Domain

---

## Principle 2: Strategic Domain First

**Statement**: Strategic domains receive dedicated Bounded Contexts with appropriate investment.

**Rationale**: Not all domains are equally important. Strategic domains (Curriculum, Assessment, Competency Intelligence) require dedicated contexts and investment.

**Implementation**:
- Strategic domains have dedicated Bounded Contexts
- Supporting domains may share contexts
- Generic domains use off-the-shelf solutions
- Investment prioritized by domain strategic value

**Examples**:
- Curriculum Domain is strategic → dedicated Curriculum Bounded Context
- Authentication is generic → use off-the-shelf identity provider

---

## Principle 3: Event-Driven Communication

**Statement**: Bounded Contexts communicate through domain events for loose coupling and eventual consistency.

**Rationale**: Event-driven communication enables loose coupling, scalability, and real-time responsiveness.

**Implementation**:
- State changes published as domain events
- Contexts subscribe to relevant events
- Event bus enables event distribution

**Examples**:
- AssessmentCompleted event triggers Competency Graph update
- CompetencyAchieved event triggers Digital Twin update
- LessonPlanCreated event triggers notification

---

## Principle 4: Anti-Corruption Layer

**Statement**: External system integration uses Anti-Corruption Layers (ACL) to protect domain integrity.

**Rationale**: External systems (SIS, LMS, government systems) have different models. ACLs translate external models to internal domain models, protecting domain integrity.

**Implementation**:
- ACLs for all external system integrations
- Translation layer between external and internal models
- Isolation from external system changes
- Versioning for external API changes

**Examples**:
- SIS ACL translates SIS student data to internal Student Profile
- LMS ACL translates LMS activity data to internal Learning Activity
- Government ACL translates government reporting format to internal format

---

## Principle 5: Shared Kernel Minimization

**Statement**: Shared Kernel is minimized to essential shared concepts only.

**Rationale**: Shared Kernel creates coupling between contexts. Minimizing shared kernel reduces coupling and maintains context autonomy.

**Implementation**:
- Shared kernel limited to essential shared concepts
- Shared concepts versioned carefully
- Changes to shared kernel require coordination
- Prefer domain events over shared kernel

**Examples**:
- Shared kernel includes: Student ID, Teacher ID, School ID
- Shared kernel excludes: Curriculum structures, Assessment models

---

## Principle 6: AI-Native Module Design

**Statement**: Modules are designed to host AI agents and support AI workflows.

**Rationale**: NUSA is AI-Native with 90% AI assistance. Modules must support AI agent execution and AI governance.

**Implementation**:
- Modules designed to host AI agents
- AI agent orchestration infrastructure
- AI governance and monitoring

**Examples**:
- Curriculum Module hosts Curriculum Agent and ATP Agent
- Assessment Module hosts Assessment Agent
- AI Orchestration Module coordinates all AI agents

---

# SECTION 3 — Bounded Context Architecture

## Overview

Bounded Context Architecture defines the software boundaries that encapsulate business domains from the Domain Model (01). Each Bounded Context represents a specific area of the business with its own Ubiquitous Language, domain model, and business rules.

## Bounded Context Definition

### BC1: Graduate Profile Context

**Purpose**: Manage the 8 dimensions of the graduate profile as the North Star for all educational activities.

**Aligned Domains**: Graduate Profile (Domain 1)

**Ubiquitous Language**: Dimension, Indicator, Developmental Phase, Assessment Framework

**Core Entities**:
- GraduateProfile
- Dimension
- DimensionIndicator
- DimensionAssessment

**Responsibilities**:
- Define graduate profile dimensions
- Define developmental phases
- Define assessment frameworks
- Track dimension development

**AI Agents**: None (strategic domain, human-governed)

---

### BC2: Curriculum Context

**Purpose**: Manage curriculum structures (CP, TP, ATP, Modul Ajar) as the execution layer for graduate profile.

**Aligned Domains**: Curriculum (Domain 3), Learning Planning (Domain 4)

**Ubiquitous Language**: CP, TP, ATP, Modul Ajar, Kurikulum Satuan Pendidikan, Cross-Reference

**Core Entities**:
- CapaianPembelajaran (CP)
- TujuanPembelajaran (TP)
- AlurTujuanPembelajaran (ATP)
- ModulAjar
- KurikulumSatuanPendidikan
- CurriculumMapping

**Responsibilities**:
- Manage CP definitions
- Generate TP from CP
- Sequence ATP
- Generate Modul Ajar
- Manage cross-references

**AI Agents**: Curriculum Agent, ATP Agent, Modul Ajar Agent

---

### BC3: Assessment Context

**Purpose**: Manage assessment design, delivery, evidence collection, and evaluation.

**Aligned Domains**: Assessment (Domain 7)

**Ubiquitous Language**: Assessment, Evidence, Evaluation, Rubric, Formative, Summative

**Core Entities**:
- Assessment
- AssessmentResult
- Evidence
- Evaluation
- Rubric

**Responsibilities**:
- Design assessments
- Deliver assessments
- Collect evidence
- Evaluate evidence
- Determine competency mastery

**AI Agents**: Assessment Agent, Rubric Agent

---

### BC4: Reporting Context

**Purpose**: Communicate student progress and achievement to stakeholders.

**Aligned Domains**: Reporting (Domain 8)

**Ubiquitous Language**: Progress Report, Narrative Report, Stakeholder, Communication

**Core Entities**:
- ProgressReport
- NarrativeReport
- ParentCommunication
- StakeholderDashboard

**Responsibilities**:
- Generate progress reports
- Generate narrative reports
- Manage parent communication
- Provide stakeholder dashboards

**AI Agents**: Narrative Report Agent

---

### BC5: Competency Intelligence Context

**Purpose**: Maintain and leverage the Competency Graph for intelligent recommendations and personalization.

**Aligned Domains**: Graduate Profile (Domain 1), Foundational Competencies (Domain 2)

**Ubiquitous Language**: Competency Graph, Competency Node, Competency Edge, Achievement, Growth Prediction

**Core Entities**:
- CompetencyGraph
- CompetencyNode
- CompetencyEdge
- AchievementEvent
- GrowthPrediction

**Responsibilities**:
- Maintain Competency Graph
- Query Competency Graph
- Analyze competency development
- Predict competency growth
- Generate competency recommendations

**AI Agents**: Competency Intelligence Agent

---

### BC6: Digital Twin Context

**Purpose**: Maintain real-time digital twins of students for personalized learning and intervention.

**Aligned Domains**: Student Wellbeing (Domain 14), Character Development (Domain 13)

**Ubiquitous Language**: Digital Twin, Learning State, Wellbeing State, Character State, Personalization Profile

**Core Entities**:
- DigitalTwin
- LearningState
- WellbeingState
- CharacterState
- PersonalizationProfile

**Responsibilities**:
- Maintain digital twin
- Update learning state
- Update wellbeing state
- Update character state
- Generate personalization recommendations

**AI Agents**: Student Coach Agent, Intervention Agent

---

### BC7: Lifelong Learning Record Context

**Purpose**: Become the national lifelong learning record connecting all educational phases.

**Aligned Domains**: Lifelong Learning Record (Domain 17), Career & Future Readiness (Domain 16)

**Ubiquitous Language**: Learning Passport, Digital Competency Transcript, Graduate Readiness Score, Career Readiness Profile

**Core Entities**:
- LifelongLearningRecord
- LearningPassport
- DigitalCompetencyTranscript
- GraduateReadinessScore
- CareerReadinessProfile

**Responsibilities**:
- Aggregate learning records
- Track educational phases
- Aggregate credentials
- Assess career readiness
- Enable portability

**AI Agents**: Lifelong Learning Agent

---

### BC8: Teacher Growth Context

**Purpose**: Manage teacher professional development and performance.

**Aligned Domains**: Teacher Professional Growth (Domain 18), Teacher Performance & Workload (Domain 19)

**Ubiquitous Language**: Professional Development, Performance, Workload, Growth, Reflection

**Core Entities**:
- Teacher
- ProfessionalDevelopment
- PerformanceRecord
- WorkloadRecord
- Reflection

**Responsibilities**:
- Track teacher performance
- Manage professional development
- Monitor workload
- Support reflection
- Generate growth recommendations

**AI Agents**: Teacher Coach Agent

---

### BC9: School Improvement Context

**Purpose**: Support school quality improvement through SPMI and quality cycles.

**Aligned Domains**: School Improvement (Domain 20), Quality Assurance & Accreditation (Domain 21)

**Ubiquitous Language**: SPMI, Quality Cycle, Quality Indicator, Benchmark, Improvement Plan

**Core Entities**:
- School
- QualityIndicator
- SPMIRecord
- BenchmarkData
- ImprovementPlan

**Responsibilities**:
- Monitor quality indicators
- Manage SPMI cycles
- Conduct benchmarking
- Generate improvement plans
- Support accreditation

**AI Agents**: None (human-governed strategic domain)

---

### BC10: Parent Partnership Context

**Purpose**: Engage parents and families in children's education.

**Aligned Domains**: Parent Partnership (Domain 15)

**Ubiquitous Language**: Partnership, Communication, Engagement, Parent Education

**Core Entities**:
- Parent
- CommunicationLog
- EngagementRecord
- ParentEducation

**Responsibilities**:
- Manage communication
- Track engagement
- Provide parent education
- Support partnership

**AI Agents**: Parent Communication Agent

---

### BC11: Education Analytics Context

**Purpose**: Provide education intelligence and analytics for decision-making.

**Aligned Domains**: All domains (cross-cutting)

**Ubiquitous Language**: Analytics, Intelligence, Prediction, Recommendation, Insight

**Core Entities**:
- AggregateAnalytics
- IntelligenceReport
- PredictiveModel
- DataVisualization

**Responsibilities**:
- Collect aggregate data
- Generate analytics
- Provide intelligence
- Support decision-making

**AI Agents**: Analytics Agent, Policy Intelligence Agent

---

### BC12: AI Orchestration Context

**Purpose**: Coordinate and orchestrate AI agents across all contexts.

**Aligned Domains**: All domains (cross-cutting)

**Ubiquitous Language**: Agent, Orchestration, Workflow, Coordination, Governance

**Core Entities**:
- AgentConfiguration
- AgentExecutionLog
- AgentPerformanceMetric
- AgentGovernanceRecord

**Responsibilities**:
- Coordinate AI agents
- Schedule agent execution
- Manage agent resources
- Monitor agent performance
- Enforce agent governance

**AI Agents**: All AI Agents (orchestration only)

---

# SECTION 4 — Context Map

## Overview

Context Map defines the relationships between Bounded Contexts, specifying communication patterns, dependencies, and integration strategies.

## Context Relationships

### Upstream/Downstream Relationships

**Graduate Profile Context (Upstream)**
- Downstream: All contexts
- Relationship: All contexts depend on graduate profile definitions
- Communication: Domain events (DimensionDefined, DimensionUpdated)

**Curriculum Context (Upstream)**
- Downstream: Learning Planning, Assessment, Reporting
- Relationship: Learning activities depend on curriculum structures
- Communication: Domain events (CPDefined, TPGenerated, ATPSequenced, ModulAjarGenerated)

**Assessment Context (Upstream)**
- Downstream: Reporting, Competency Intelligence, Digital Twin
- Relationship: Reports and intelligence depend on assessment results
- Communication: Domain events (AssessmentCompleted, EvidenceEvaluated, CompetencyValidated)

**Competency Intelligence Context (Bidirectional)**
- Upstream: Assessment, Learning Delivery
- Downstream: Digital Twin, Reporting
- Relationship: Competency data flows both directions
- Communication: Domain events (CompetencyAchieved, CompetencyGapIdentified, GrowthPredicted)

**Digital Twin Context (Downstream)**
- Upstream: All learning contexts
- Relationship: Digital twin aggregates data from all learning activities
- Communication: Domain events (LearningActivityCompleted, AssessmentCompleted, CompetencyAchieved)

**Lifelong Learning Record Context (Downstream)**
- Upstream: All contexts
- Relationship: Lifelong learning record aggregates data from all phases
- Communication: Domain events (AchievementRecorded, CredentialIssued, PhaseCompleted)

---

### Shared Kernel

**Shared Kernel Components**:
- Student ID (universal identifier)
- Teacher ID (universal identifier)
- School ID (universal identifier)
- Phase (educational phase: Fondasi, A-F)
- Timestamp (universal time reference)

**Shared Kernel Governance**:
- Changes require coordination across all contexts
- Versioned to enable backward compatibility
- Minimal to reduce coupling
- Documented in shared kernel specification

---

### Anti-Corruption Layers (ACL)

**SIS ACL**
- Purpose: Integrate with School Information Systems
- Translation: SIS student data → internal Student Profile
- Isolation: Protects internal domain from SIS model changes
- Versioning: Supports multiple SIS versions

**LMS ACL**
- Purpose: Integrate with Learning Management Systems
- Translation: LMS activity data → internal Learning Activity
- Isolation: Protects internal domain from LMS model changes
- Versioning: Supports multiple LMS versions

**Government ACL**
- Purpose: Integrate with government systems (Rapor Pendidikan, National Assessment)
- Translation: Government format → internal format
- Isolation: Protects internal domain from government format changes
- Versioning: Supports multiple government system versions

---

### Context Integration Patterns

**Curriculum Context → Assessment Context**
- Pattern: Upstream/Downstream
- Communication: Domain events (ModulAjarGenerated → AssessmentDesigned)
- Data Flow: Curriculum structures inform assessment design

**Assessment Context → Competency Intelligence Context**
- Pattern: Upstream/Downstream
- Communication: Domain events (CompetencyValidated → CompetencyGraphUpdated)
- Data Flow: Assessment results update competency graph

**Competency Intelligence Context → Digital Twin Context**
- Pattern: Upstream/Downstream
- Communication: Domain events (CompetencyAchieved → DigitalTwinUpdated)
- Data Flow: Competency data updates digital twin

**All Contexts → Lifelong Learning Record Context**
- Pattern: Upstream/Downstream
- Communication: Domain events (AchievementRecorded → LifelongRecordUpdated)
- Data Flow: All achievements update lifelong learning record

---

# SECTION 5 — Module Architecture

## Overview

Module Architecture defines the internal structure of each Bounded Context, organizing code into cohesive modules that align with domain concepts and business capabilities.

## Module Design Principles

### Domain-Driven Modules

**Principle**: Modules align with domain concepts, not technical concerns.

**Implementation**:
- Each module represents a domain concept or aggregate
- Modules contain related entities, value objects, and domain services
- Module boundaries enforce encapsulation
- Modules communicate through well-defined interfaces

**Examples**:
- Curriculum Context modules: CP Module, TP Module, ATP Module, Modul Ajar Module
- Assessment Context modules: Assessment Module, Evidence Module, Evaluation Module

---

### Aggregate Root Modules

**Principle**: Each aggregate root has its own module to enforce consistency boundaries.

**Implementation**:
- Aggregate root is the entry point for the module
- Module enforces aggregate invariants
- External access only through aggregate root
- Module maintains aggregate consistency

**Examples**:
- CP Aggregate Root Module: Manages CP and its relationships
- Assessment Aggregate Root Module: Manages assessment and its results
- Competency Graph Aggregate Root Module: Manages competency graph and its nodes

---

### Module Dependencies

**Principle**: Module dependencies follow domain dependencies, not technical dependencies.

**Implementation**:
- Upstream modules have no dependencies on downstream modules
- Downstream modules depend on upstream modules through interfaces
- Circular dependencies are prohibited
- Module dependencies are explicit and documented

**Examples**:
- CP Module has no dependencies (upstream)
- Assessment Module depends on CP Module (downstream)
- Competency Graph Module depends on Assessment Module (downstream)

---

## Module Structure by Context

### Curriculum Context Modules

**CP Module**
- Aggregate Root: CapaianPembelajaran
- Entities: CP, Subject, Phase, Element, Subelement
- Value Objects: LearningAchievement, CurriculumHierarchy
- Domain Services: CPAlignmentService, CurriculumHierarchyService
- Interfaces: CPRepository, CPService, CurriculumHierarchyRepository

**TP Module**
- Aggregate Root: TujuanPembelajaran
- Entities: TP
- Value Objects: LearningObjective, CurriculumTraceability (subject_id, phase_id, element_id, subelement_id)
- Domain Services: TPGenerationService
- Interfaces: TPRepository, TPService

**ATP Module**
- Aggregate Root: AlurTujuanPembelajaran
- Entities: ATP
- Value Objects: Sequence, TimeAllocation, Prerequisite
- Domain Services: ATPSequencingService
- Interfaces: ATPRepository, ATPService

**Modul Ajar Module**
- Aggregate Root: ModulAjar
- Entities: ModulAjar, ResourceAllocation, ActivitySequence
- Value Objects: LessonPlan, DifferentiationStrategy
- Domain Services: ModulAjarGenerationService
- Interfaces: ModulAjarRepository, ModulAjarService

---

### Assessment Context Modules

**Assessment Module**
- Aggregate Root: Assessment
- Entities: Assessment, AssessmentItem
- Value Objects: AssessmentType, DifficultyLevel
- Domain Services: AssessmentDesignService
- Interfaces: AssessmentRepository, AssessmentService

**Evidence Module**
- Aggregate Root: Evidence
- Entities: Evidence, EvidenceItem
- Value Objects: EvidenceType, ContentType
- Domain Services: EvidenceCollectionService
- Interfaces: EvidenceRepository, EvidenceService

**Evaluation Module**
- Aggregate Root: Evaluation
- Entities: Evaluation, EvaluationCriteria
- Value Objects: Score, Level, Feedback
- Domain Services: EvidenceEvaluationService
- Interfaces: EvaluationRepository, EvaluationService

---

### Competency Intelligence Context Modules

**Competency Graph Module**
- Aggregate Root: CompetencyGraph
- Entities: CompetencyGraph, CompetencyNode, CompetencyEdge
- Value Objects: CompetencyLevel, ConfidenceScore
- Domain Services: CompetencyGraphService
- Interfaces: CompetencyGraphRepository, CompetencyGraphQueryService

**Achievement Module**
- Aggregate Root: AchievementEvent
- Entities: AchievementEvent, AchievementRecord
- Value Objects: AchievementType, Timestamp
- Domain Services: AchievementTrackingService
- Interfaces: AchievementRepository, AchievementService

**Prediction Module**
- Aggregate Root: GrowthPrediction
- Entities: GrowthPrediction, PredictionModel
- Value Objects: PredictionConfidence, TimeHorizon
- Domain Services: GrowthPredictionService
- Interfaces: PredictionRepository, PredictionService

---

# SECTION 6 — Module Architecture

## Overview

Module Architecture defines the modules that implement the Bounded Contexts, ensuring clear module boundaries within a single deployable application.

Modules are logical boundaries inside a single deployable Modular Monolith application.

All modules are deployed together as a single application.

Modules are logical boundaries and are not independently deployable.

---

## Module Design Principles

### Context-Module Alignment

**Principle**: Each Bounded Context is implemented as one or more modules.

**Implementation**:
- Strategic contexts have dedicated modules
- Supporting contexts may share modules
- Module boundaries align with context boundaries
- Modules maintain context autonomy

**Examples**:
- Curriculum Context → Curriculum Module
- Assessment Context → Assessment Module
- Competency Intelligence Context → Competency Module

---

### Module Communication

**Principle**: Modules communicate through well-defined interfaces.

**Implementation**:
- Direct method calls for synchronous communication
- Application services for coordination
- Domain events for internal event-driven communication

**Examples**:
- Curriculum Module exposes application services for curriculum operations
- Assessment Module publishes domain events for assessment completion
- Competency Module subscribes to assessment domain events

---

# Internal Module Communication

Modules communicate through:

- Direct Method Calls
- Application Services
- Domain Events (internal only)

Domain Events are application-level patterns and do not require distributed infrastructure.

---

# Application Module Structure

## Core Education Modules

- Curriculum Module
- Learning Planning Module
- Assessment Module
- Evidence Module
- Reporting Module

---

## User Modules

- Teacher Module
- Administration Module

---

## AI Modules

- AI Orchestration Module
- TP Agent Module
- ATP Agent Module
- Modul Ajar Agent Module
- Assessment Agent Module
- Rubric Agent Module
- Narrative Report Agent Module

---

## Platform Modules

- Authentication & Authorization Module
- Notification Module
- Audit Module
- Configuration Module

---

All modules are deployed together as a single application.

---

# Data Access Strategy

All modules use a shared PostgreSQL database.

Logical separation may be implemented using:

- Schemas
- Module ownership conventions

but not separate databases.

Single database is mandatory for MVP Wave 1.

---

# Integration Strategy

## Internal Integration

- Module-to-Module Interaction
- Application Services
- Domain Services

---

## External Integration

- LLM Providers
- Email Providers
- Future SIS Integration

---

External integrations must remain isolated through adapter interfaces.

---

# Deployment Architecture

## Single Deployable Application

Components:

- React Frontend
- Backend API (Go)
- AI Runtime (Python + LangGraph)
- PostgreSQL
- RabbitMQ

Deployment Method:

Docker Compose

---

# Application Component Architecture

## Backend API (Go)

**Responsibilities:**
- Authentication
- Authorization
- User Management
- Workflow Orchestration
- API Layer
- Persistence
- Audit Logging
- Notification

**Communication:**
- Frontend → Backend API (HTTP REST)
- Backend API → AI Runtime (HTTP REST)

## AI Runtime (Python + LangGraph)

**Responsibilities:**
- LangGraph Execution
- Prompt Execution
- LLM Integration
- Structured Output Generation
- AI Validation

**Communication:**
- Backend API → AI Runtime (HTTP REST)

**Note:** The AI Runtime is NOT a business microservice and NOT an independently deployable product. It exists solely to execute AI workflows and LLM interactions on behalf of the Backend API.

## Communication Flow

```
React Frontend
    ↓ HTTP REST
Backend API (Go)
    ↓ HTTP REST (for AI operations)
AI Runtime (Python + LangGraph)
    ↓ (LLM API calls)
LLM Providers
```

**Protocol Selection:**
- Backend API → AI Runtime uses HTTP REST to reduce implementation complexity during MVP
- Future protocols may include gRPC, RabbitMQ, NATS, or Kafka
- No business workflow changes shall be required when communication protocol changes

---

# Application Layer Architecture

## Presentation Layer

- React UI

---

## API Layer

- REST Controllers

---

## Application Layer

- Use Cases
- Application Services

---

## Domain Layer

- Domain Models
- Domain Rules
- Domain Events

---

## Infrastructure Layer

- PostgreSQL
- RabbitMQ
- AI Providers

---

Dependencies flow inward only.

---

# MVP Application Flow

CP
↓
TP
↓
ATP
↓
Modul Ajar
↓
Assessment
↓
Rubric
↓
Evidence
↓
Narrative Report

Application modules must support this workflow directly.

Application architecture must not introduce unnecessary abstraction layers that obscure this workflow.

---

# Architecture Optimization Principle

The architecture is optimized for:

- maintainability
- clarity
- implementation speed
- educational workflow automation

rather than hypothetical future scale.

---

# Application Architecture Freeze Declaration

The following decisions are frozen:

- Modular Monolith
- React
- Go
- PostgreSQL
- RabbitMQ
- REST API
- Docker Compose

No application architecture changes may be introduced during MVP implementation without Architecture Owner approval.

---

# MVP Implementation Scope

Purpose:

Prevent misunderstanding between architectural coverage and implementation scope.

The Application Architecture represents the long-term target architecture of the NUSA platform.

Not all modules are implemented during MVP Wave 1.

---

## MVP Wave 1 Modules

Implemented:

- Curriculum Module
- Learning Planning Module
- Assessment Module
- Evidence Module
- Reporting Module
- Teacher Module
- Administration Module
- Authentication & Authorization Module
- AI Orchestration Module
- TP Agent Module
- ATP Agent Module
- Modul Ajar Agent Module
- Assessment Agent Module
- Rubric Agent Module
- Narrative Report Agent Module

---

## Future Architecture Modules

Architecturally defined but not implemented in MVP Wave 1:

- Competency Module
- Digital Twin Module
- Lifelong Learning Module
- Teacher Growth Module
- Parent Partnership Module
- School Improvement Module
- Education Analytics Module

Future modules remain part of platform evolution but are excluded from MVP delivery planning.

---

# Domain Prioritization Matrix

Purpose:

Align implementation priorities with Domain Model and Capability Model.

| Domain Context                 | Classification    | MVP Status |
| ------------------------------ | ----------------- | ---------- |
| Curriculum                     | Core Domain       | MVP        |
| Learning Planning              | Core Domain       | MVP        |
| Assessment                     | Core Domain       | MVP        |
| Evidence                       | Core Domain       | MVP        |
| Reporting                      | Core Domain       | MVP        |
| Teacher                        | Supporting Domain | MVP        |
| Administration                 | Supporting Domain | MVP        |
| Authentication & Authorization | Supporting Domain | MVP        |
| AI Orchestration               | Supporting Domain | MVP        |
| Competency                     | Core Domain       | Future     |
| Digital Twin                   | Strategic Domain  | Future     |
| Lifelong Learning              | Strategic Domain  | Future     |
| Teacher Growth                 | Strategic Domain  | Future     |
| Parent Partnership             | Supporting Domain | Future     |
| School Improvement             | Supporting Domain | Future     |
| Education Analytics            | Strategic Domain  | Future     |

**Domain Classifications**:

- **Core Domain**: Contains business-critical capabilities that differentiate the platform and require specialized domain expertise. These domains are essential for the core value proposition.

- **Supporting Domain**: Provides necessary capabilities that support the core domains but do not provide competitive differentiation. These domains can often leverage off-the-shelf solutions or standard implementations.

- **Strategic Domain**: Contains capabilities that are critical for long-term platform evolution but are not required for MVP Wave 1. These domains represent future strategic initiatives.

---

# Platform Identity

NUSA is the official platform identity.

NUSA is the official platform identity across business, architecture, engineering, and implementation documentation.

NUSA functions as a national education platform that orchestrates curriculum planning, learning design, assessment, evidence collection, reporting, and future lifelong learning modules.

---

## Module Definitions

### Curriculum Module

**Purpose**: Implement Curriculum Context capabilities.

**Bounded Context**: Curriculum Context

**Responsibilities**:
- CP management
- TP generation
- ATP sequencing
- Modul Ajar generation
- Curriculum alignment

**AI Agents**: Curriculum Agent, ATP Agent, Modul Ajar Agent

**Data Store**: PostgreSQL (shared database)

**Events Published**:
- TPGenerated
- ATPSequenced
- ModulAjarGenerated
- CurriculumAligned

**Events Subscribed**:
- CPDefined (from Graduate Profile Context)

---

### Assessment Module

**Purpose**: Implement Assessment Context capabilities.

**Bounded Context**: Assessment Context

**Responsibilities**:
- Assessment design
- Assessment delivery
- Evidence collection
- Evidence evaluation
- Competency validation

**AI Agents**: Assessment Agent, Rubric Agent

**Data Store**: PostgreSQL (shared database)

**Events Published**:
- AssessmentCreated
- AssessmentDelivered
- EvidenceCollected
- EvidenceEvaluated
- CompetencyValidated

**Events Subscribed**:
- ModulAjarGenerated (from Curriculum Context)

---

### Reporting Module

**Purpose**: Implement Reporting Context capabilities.

**Bounded Context**: Reporting Context

**Responsibilities**:
- Progress report generation
- Narrative report generation
- Parent communication
- Stakeholder dashboards

**AI Agents**: Narrative Report Agent

**Data Store**: PostgreSQL (shared database)

**Events Published**:
- ReportGenerated
- NarrativeGenerated
- CommunicationSent

**Events Subscribed**:
- CompetencyValidated (from Assessment Context)
- CompetencyAchieved (from Competency Intelligence Context)

---

### Competency Module

**Purpose**: Implement Competency Intelligence Context capabilities.

**Bounded Context**: Competency Intelligence Context

**Responsibilities**:
- Competency Graph management
- Competency Graph query
- Competency Graph analysis
- Competency Graph prediction
- Competency Graph recommendation

**AI Agents**: Competency Intelligence Agent

**Data Store**: PostgreSQL (shared database)

**Events Published**:
- CompetencyGraphUpdated
- CompetencyAchieved
- CompetencyGapIdentified
- GrowthPredicted

**Events Subscribed**:
- CompetencyValidated (from Assessment Context)
- LearningActivityCompleted (from Learning Delivery Context)

---

### Digital Twin Module

**Purpose**: Implement Digital Twin Context capabilities.

**Bounded Context**: Digital Twin Context

**Responsibilities**:
- Digital twin management
- Learning state update
- Wellbeing state update
- Character state update
- Personalization recommendation

**AI Agents**: Student Coach Agent, Intervention Agent

**Data Store**: PostgreSQL (shared database)

**Events Published**:
- DigitalTwinUpdated
- PersonalizationRecommended
- InterventionSuggested

**Events Subscribed**:
- All learning events from all contexts

---

### Lifelong Learning Module

**Purpose**: Implement Lifelong Learning Record Context capabilities.

**Bounded Context**: Lifelong Learning Record Context

**Responsibilities**:
- Record aggregation
- Educational phase tracking
- Credential aggregation
- Career readiness assessment
- Portability management

**AI Agents**: Lifelong Learning Agent

**Data Store**: PostgreSQL (shared database)

**Events Published**:
- RecordAggregated
- PhaseCompleted
- CredentialAdded
- CareerAssessmentGenerated

**Events Subscribed**:
- AchievementRecorded (from all contexts)

---

### Teacher Growth Module

**Purpose**: Implement Teacher Growth Context capabilities.

**Bounded Context**: Teacher Growth Context

**Responsibilities**:
- Teacher performance tracking
- Professional development management
- Workload monitoring
- Reflection support
- Growth recommendation

**AI Agents**: Teacher Coach Agent

**Data Store**: PostgreSQL

**Events Published**:
- PerformanceUpdated
- PDCompleted
- ReflectionRecorded
- GrowthUpdated

**Events Subscribed**:
- None (teacher-initiated activities)

---

### School Improvement Module

**Purpose**: Implement School Improvement Context capabilities.

**Bounded Context**: School Improvement Context

**Responsibilities**:
- Quality indicator monitoring
- SPMI cycle management
- Benchmarking
- Improvement planning
- Quality reporting

**AI Agents**: None (human-governed)

**Data Store**: PostgreSQL

**Events Published**:
- QualityIndicatorUpdated
- SPMICycleCompleted
- BenchmarkCompleted
- ImprovementPlanGenerated

**Events Subscribed**:
- Assessment data from Assessment Context
- Teacher performance from Teacher Growth Context

---

### Parent Partnership Module

**Purpose**: Implement Parent Partnership Context capabilities.

**Bounded Context**: Parent Partnership Context

**Responsibilities**:
- Communication management
- Engagement tracking
- Parent education
- Partnership support

**AI Agents**: Parent Communication Agent

**Data Store**: PostgreSQL

**Events Published**:
- CommunicationSent
- EngagementRecorded
- ParentEducationCompleted

**Events Subscribed**:
- ReportGenerated (from Reporting Context)

---

### Education Analytics Module

**Purpose**: Implement Education Analytics Context capabilities.

**Bounded Context**: Education Analytics Context

**Responsibilities**:
- Data collection
- Descriptive analytics
- Predictive analytics
- Prescriptive analytics
- Intelligence reporting

**AI Agents**: Analytics Agent, Policy Intelligence Agent

**Data Store**: PostgreSQL (shared database)

**Events Published**:
- AnalyticsGenerated
- PredictionGenerated
- RecommendationGenerated
- ReportPublished

**Events Subscribed**:
- All events from all contexts (for data collection)

---

### Authentication & Authorization Module

**Purpose**: Implement authentication and authorization capabilities.

**Bounded Context**: Identity & Access Context

**Responsibilities**:
- Login
- Access Token Management
- Refresh Token Management
- Role Validation
- Session Validation
- Route Authorization

**Authentication Strategy**: Custom JWT

**Authorization Strategy**: RBAC

**Roles**:
- Administrator
- Teacher

**Data Store**: PostgreSQL (shared database)

**Events Published**:
- UserLoggedIn
- UserLoggedOut
- TokenRefreshed
- RoleAssigned

**Events Subscribed**:
- None (module-initiated activities)

---

### AI Orchestration Module

**Purpose**: Implement AI Orchestration Context capabilities.

**Bounded Context**: AI Orchestration Context

**Responsibilities**:
- Agent coordination
- Agent scheduling
- Agent resource management
- Agent performance monitoring
- Agent governance enforcement

**AI Agents**: All AI Agents (orchestration only)

**Data Store**: PostgreSQL (shared database)

**Events Published**:
- AgentExecuted
- AgentCompleted
- AgentFailed
- AgentPerformanceUpdated

**Events Subscribed**:
- Workflow trigger events from all contexts

---

# SECTION 7 — Event Driven Architecture

## Overview

Event-Driven Architecture enables loose coupling between modules through asynchronous event-based communication, supporting real-time responsiveness and scalability.

## Event Design Principles

### Domain Events

**Principle**: State changes are published as domain events.

**Implementation**:
- Events represent domain-relevant state changes
- Events are immutable
- Events have unique identifiers
- Events include timestamp and source

**Examples**:
- TPGenerated: TP generated from CP
- AssessmentCompleted: Assessment delivery completed
- CompetencyAchieved: Competency milestone achieved

---

### Event Schema

**Principle**: Events follow a consistent schema for interoperability.

**Implementation**:
- Event header: eventId, eventType, timestamp, source, correlationId
- Event body: domain-specific payload
- Event metadata: version, schemaUri
- Event validation: schema validation at publish time

**Schema Example**:
```json
{
  "eventId": "uuid",
  "eventType": "TPGenerated",
  "timestamp": "2026-06-02T12:00:00Z",
  "source": "CurriculumModule",
  "correlationId": "uuid",
  "version": "1.0",
  "schemaUri": "https://schema.nusa.id/events/TPGenerated/v1",
  "body": {
    "tpId": "uuid",
    "cpId": "uuid",
    "phase": "A",
    "subject": "Matematika"
  }
}
```

---

## Event Bus

### Event Bus Implementation

**Technology**: RabbitMQ for internal event communication

**Features**:
- Event queues for each event type
- Consumer groups for event processing

**Topics**:
- curriculum.events: CP, TP, ATP, Modul Ajar events
- assessment.events: Assessment, evidence, evaluation events
- competency.events: Competency graph, achievement events
- digital-twin.events: Digital twin, personalization events
- lifelong-learning.events: Record, credential events

---

### Event Publishing

**Publishing Process**:
1. Module detects state change
2. Module creates domain event
3. Module validates event schema
4. Module publishes event to event bus
5. Module acknowledges publish success

**Publishing Guarantees**:
- At-least-once delivery
- Dead letter queue for failed events

---

### Event Subscription

**Subscription Process**:
1. Module subscribes to event topic
2. Module creates consumer group
3. Module processes events
4. Module acknowledges processing
5. Module handles processing failures

**Subscription Patterns**:
- Event filtering: Subscribe to specific event types
- Event routing: Route events based on content
- Event aggregation: Aggregate multiple events
- Event transformation: Transform events for consumption

---

## Event Flows

### Curriculum Flow

```
CPDefined (Graduate Profile Context)
    ↓
TPGenerated (Curriculum Context)
    ↓
ATPSequenced (Curriculum Context)
    ↓
ModulAjarGenerated (Curriculum Context)
    ↓
AssessmentDesigned (Assessment Context)
```

---

### Assessment Flow

```
AssessmentCreated (Assessment Context)
    ↓
AssessmentDelivered (Assessment Context)
    ↓
EvidenceCollected (Assessment Context)
    ↓
EvidenceEvaluated (Assessment Context)
    ↓
CompetencyValidated (Assessment Context)
    ↓
CompetencyAchieved (Competency Intelligence Context)
    ↓
DigitalTwinUpdated (Digital Twin Context)
```

---

### Reporting Flow

```
CompetencyValidated (Assessment Context)
    ↓
CompetencyAchieved (Competency Intelligence Context)
    ↓
ReportGenerated (Reporting Context)
    ↓
CommunicationSent (Reporting Context)
```

---

# Domain Event Strategy

Purpose:

Support internal module communication.

Domain Events are used only for:

- Workflow Progression
- Audit Notifications
- Internal Process Coordination

Examples:

- TPGenerated
- ATPGenerated
- AssessmentCompleted
- ReportGenerated

Domain Events are internal application patterns.

They do not require:

- Event Store
- Event Replay
- Event Sourcing Infrastructure

Event Sourcing is not implemented in MVP Wave 1.

---

# Internal Event Strategy

The application uses Internal Domain Events for workflow coordination.

## RabbitMQ Usage

RabbitMQ is used for:

- ATP Generation Jobs
- Modul Ajar Generation Jobs
- Narrative Report Generation Jobs
- Notification Dispatching
- Long-Running AI Tasks

RabbitMQ is NOT used for:

- Authentication
- CRUD Operations
- Source of Truth
- Event Sourcing
- Transaction Processing

PostgreSQL remains the system of record.

RabbitMQ is used only for asynchronous processing.

Domain Events are application-level constructs and do not require Kafka infrastructure.

Kafka is not part of MVP Wave 1 architecture.

---

# Domain Event Usage

Domain Events are used for:

- Workflow Progression
- Audit Notifications
- Internal Process Coordination

Examples:

- TPGenerated
- ATPGenerated
- AssessmentCompleted
- ReportGenerated

Event replay capability is not implemented in MVP Wave 1.

Domain Events are not persisted as an Event Store.

---

# SECTION 8 — API Architecture

## Overview

API Architecture defines the principles and patterns for API design, ensuring consistency, interoperability, and maintainability across all modules.

## API Design Principles

### RESTful Design

**Principle**: APIs follow RESTful design principles for simplicity and interoperability.

**Implementation**:
- Resource-based URLs
- HTTP methods (GET, POST, PUT, DELETE, PATCH)
- Status codes for responses
- JSON for request/response bodies
- HATEOAS for navigation

**Examples**:
- GET /api/curriculum/cp/{id}
- POST /api/curriculum/tp
- PUT /api/curriculum/tp/{id}
- DELETE /api/curriculum/tp/{id}

---

### API Versioning

**Principle**: APIs are versioned to enable evolution without breaking existing clients.

**Implementation**:
- URL-based versioning (/api/v1/resource)
- Semantic versioning (major.minor.patch)
- Deprecation policy for old versions
- Backward compatibility maintenance

**Versioning Strategy**:
- Major version: Breaking changes
- Minor version: Non-breaking additions
- Patch version: Bug fixes

---

### API Documentation

**Principle**: All APIs are documented using OpenAPI/Swagger for discoverability and interoperability.

**Implementation**:
- OpenAPI specification for all APIs
- Interactive API documentation (Swagger UI)
- API examples and use cases
- API change documentation

**Documentation Components**:
- API endpoints
- Request/response schemas
- Authentication requirements
- Error responses
- Rate limits

---

### API Security

**Principle**: APIs are secured through authentication, authorization, and rate limiting.

**Implementation**:
- Custom JWT for user authentication
- Role-based access control (RBAC)
- Rate limiting per user

**Security Layers**:
- Authentication: Verify identity
- Authorization: Verify permissions
- Rate limiting: Prevent abuse
- Input validation: Prevent injection

---

# API Strategy

The application exposes REST APIs.

REST APIs serve:

- Teacher Portal
- Administration Portal
- Future Integrations

REST API is the only supported API style in MVP Wave 1.

---

## API Communication Patterns

### Synchronous Communication

**Pattern**: Request-response for immediate responses.

**Use Cases**:
- Query operations (GET)
- Command operations requiring immediate confirmation (POST, PUT, DELETE)
- Real-time data retrieval

**Implementation**:
- REST APIs over HTTP/HTTPS
- Timeout handling
- Retry logic

---

### Asynchronous Communication

**Pattern**: Event-driven for eventual consistency.

**Use Cases**:
- Long-running operations
- High-throughput operations
- Decoupled modules
- Event-driven workflows

**Implementation**:
- Internal event bus (RabbitMQ)
- Event subscription
- Event acknowledgment

---

# SECTION 9 — Integration Architecture

## Overview

Integration Architecture defines how NUSA integrates with external systems (SIS, LMS, government systems) through Anti-Corruption Layers (ACL) to protect domain integrity.

## Integration Patterns

### Anti-Corruption Layer (ACL)

**Principle**: External system integration uses ACL to translate external models to internal domain models.

**Implementation**:
- Translation layer between external and internal models
- Isolation from external system changes
- Versioning for external API changes
- Error handling and retry logic

**ACL Components**:
- External API client
- Model translator
- Internal domain model mapper
- Error handler

---

### SIS Integration

**Purpose**: Integrate with School Information Systems for student and teacher data.

**External System**: School Information System (SIS)

**ACL Responsibilities**:
- Translate SIS student data to internal Student Profile
- Translate SIS teacher data to internal Teacher Profile
- Translate SIS school data to internal School Profile
- Sync data changes from SIS to NUSA

**Data Flow**:
```
SIS → SIS ACL → Student Profile Context → Internal Domain
```

**Integration Points**:
- Student enrollment
- Teacher assignment
- School information
- Class schedules

---

### LMS Integration

**Purpose**: Integrate with Learning Management Systems for learning activity data.

**External System**: Learning Management System (LMS)

**ACL Responsibilities**:
- Translate LMS activity data to internal Learning Activity
- Translate LMS submission data to internal Evidence
- Translate LMS grade data to internal Assessment Result
- Sync learning activities from LMS to NUSA

**Data Flow**:
```
LMS → LMS ACL → Assessment Context → Internal Domain
```

**Integration Points**:
- Learning activities
- Student submissions
- Activity grades
- Attendance data

---

### Government Integration

**Purpose**: Integrate with government systems for reporting and compliance.

**External Systems**: Rapor Pendidikan, National Assessment, Government Databases

**ACL Responsibilities**:
- Translate government reporting format to internal format
- Translate internal format to government reporting format
- Submit reports to government systems
- Receive government data and translate to internal format

**Data Flow**:
```
NUSA → Government ACL → Government System
Government System → Government ACL → NUSA
```

**Integration Points**:
- Rapor Pendidikan submission
- National Assessment data
- Student registry
- Teacher certification

---

## Integration Governance

### Integration Security

**Principle**: All integrations are secured with authentication, authorization, and encryption.

**Implementation**:
- TLS encryption for all communications
- IP whitelisting for external systems

---

### Integration Monitoring

**Principle**: All integrations are monitored for availability, performance, and errors.

**Implementation**:
- Integration health checks
- Integration performance metrics
- Integration error logging
- Integration alerting

**Monitoring Metrics**:
- Integration availability
- Integration latency
- Integration error rate
- Integration throughput

---

### Integration Error Handling

**Principle**: All integrations have robust error handling and retry logic.

**Implementation**:
- Retry logic for transient failures
- Dead letter queue for failed messages
- Circuit breaker for failing integrations
- Error notification and escalation

**Error Handling Strategies**:
- Retry with exponential backoff
- Dead letter queue for manual review
- Circuit breaker to prevent cascading failures
- Fallback to manual processes

---

# SECTION 10 — AI Integration Architecture

## Overview

AI Integration Architecture defines how AI agents are integrated within modules, enabling AI-Native operations with 90% AI assistance while maintaining 10% human governance.

## AI Agent Integration

### Agent Hosting

**Principle**: AI agents are hosted within modules that align with their domain responsibilities.

**Implementation**:
- Each module hosts relevant AI agents
- Agents have dedicated execution environments
- Agents are isolated from module logic

**Agent Hosting Examples**:
- Curriculum Module hosts: Curriculum Agent, ATP Agent, Modul Ajar Agent
- Assessment Module hosts: Assessment Agent, Rubric Agent
- Competency Module hosts: Competency Intelligence Agent
- Digital Twin Module hosts: Student Coach Agent, Intervention Agent

---

### Agent Orchestration

**Principle**: AI Orchestration Module coordinates agent execution across modules.

**Implementation**:
- Central orchestration module
- Agent scheduling and resource management
- Agent workflow coordination
- Agent performance monitoring

**Orchestration Responsibilities**:
- Schedule agent execution
- Manage agent resources (compute, memory)
- Coordinate multi-agent workflows
- Monitor agent performance
- Enforce agent governance

---

### Agent Communication

**Principle**: AI agents communicate through well-defined interfaces and events.

**Implementation**:
- Agent-to-agent communication through events
- Agent-to-module communication through APIs
- Agent-to-human communication through notifications

**Communication Patterns**:
- Request-response for immediate agent responses
- Event-driven for agent workflows
- Notification for human-in-the-loop
- State persistence for agent continuity

---

# AI Development Principles

The MVP AI architecture is prompt-centric.

Supported practices:

- Prompt Versioning
- Prompt Evaluation
- Human Review
- Output Quality Assessment
- Continuous Prompt Improvement

The MVP does not implement custom model training infrastructure.

The MVP uses external LLM providers through the AI Orchestration Module.

---

## AI Governance

### Human-in-the-Loop

**Principle**: AI agent decisions require human validation for high-stakes decisions.

**Implementation**:
- Validation points defined in AI Architecture
- Human approval workflow
- Override mechanism for human decisions
- Audit trail for AI decisions

**Validation Points**:
- Curriculum Agent: Teacher validates generated curriculum
- Assessment Agent: Teacher validates assessment design
- Competency Intelligence Agent: Teacher validates competency assessment
- Student Coach Agent: Teacher validates intervention recommendations

---

### AI Safety

**Principle**: AI agents are designed with safety guardrails to prevent harmful outputs.

**Implementation**:
- Input validation and sanitization
- Output filtering and moderation
- Fail-safe mechanisms
- Human override capability

**Safety Measures**:
- Content filtering for inappropriate content
- Bias detection and mitigation
- Error handling for edge cases
- Human escalation for uncertain decisions

---

### AI Auditability

**Principle**: AI agent decisions are logged and auditable for transparency and accountability.

**Implementation**:
- Decision logging
- Explanation generation
- Audit trail
- Compliance reporting

**Audit Components**:
- Agent execution logs
- Decision explanations
- Human approval logs
- Performance metrics

---

# SECTION 11 — Identity & Access Architecture

## Overview

Identity & Access Architecture defines authentication, authorization, and access control for NUSA, ensuring security and compliance while enabling seamless user experience.

## Authentication

### Authentication Strategy

**Principle**: Use Custom JWT for standardized authentication and identity management.

**Implementation**:
- Custom JWT for authentication and identity management
- RBAC for authorization
- Access token management
- Refresh token management
- Session validation

The MVP does not depend on external identity providers.

Authentication and authorization are implemented internally through the Authentication & Authorization Module.

---

# Authentication Strategy

Authentication:

Custom JWT

Authorization:

RBAC

Components:

- Access Token
- Refresh Token
- Session Validation
- Role Validation

Roles:

- Administrator
- Teacher

OAuth2 and OpenID Connect are not part of MVP Wave 1 architecture.

---

# MVP User Types

Supported:

- Administrator
- Teacher

**Teachers**
- Authentication: Custom JWT
- Roles: Teacher, Subject Teacher, Homeroom Teacher
- Permissions: Curriculum management, assessment, reporting

**Administrators**
- Authentication: Custom JWT
- Roles: School Admin, District Admin, National Admin
- Permissions: School management, analytics, governance

---

# Future User Types

Future platform expansion may include:

- Student
- Parent
- School Leader
- Supervisor

Future user types are excluded from MVP Wave 1 implementation planning.

---

## Authorization

### Role-Based Access Control (RBAC)

**Principle**: Use RBAC for authorization based on user roles.

**Implementation**:
- Role definitions for each user type
- Permission assignments to roles
- Role inheritance for hierarchical permissions
- Permission checks at module boundaries

**Role Hierarchy**:
- National Admin (highest permissions)
- District Admin
- School Admin
- Teacher
- Student
- Parent (lowest permissions)

---

# Authorization Strategy

Authorization is implemented using Role-Based Access Control (RBAC).

Supported Roles:

- Administrator
- Teacher

ABAC is not implemented in MVP Wave 1.

---

## Access Control

### Module-Level Access Control

**Principle**: Role-Based Access Control (RBAC) for authorization.

**Implementation**:
- Authentication & Authorization Module enforces RBAC
- Modules check user roles
- Permission checks on all API calls
- Audit logging for access attempts

**Roles**:
- Administrator
- Teacher

**Access Control Flow**:
1. API validates authentication
2. API extracts user context
3. Module validates authorization
4. Module checks permissions
5. Module logs access

---

### Data-Level Access Control

**Principle**: Data access is controlled based on user permissions and data ownership.

**Implementation**:
- Row-level security in databases
- Data filtering based on user context
- Data masking for sensitive fields
- Data encryption at rest and in transit

**Data Access Rules**:
- Teachers can access their students' data
- Parents can access their children's data
- Students can access their own data
- Administrators can access data within their scope

---

# SECTION 12 — Notification Architecture

## Overview

Notification Architecture defines how NUSA communicates with users through various channels, ensuring timely and relevant information delivery.

## Notification Channels

### Email Notifications

**Purpose**: Send formal communications and reports via email.

**Use Cases**:
- Report generation notifications
- Parent communication
- System alerts
- Administrative notifications

**Implementation**:
- Email service for sending emails
- Email templates for different notification types
- Email scheduling and batching
- Email delivery tracking

---

### Push Notifications

**Purpose**: Send real-time notifications to mobile and web applications.

**Use Cases**:
- Assessment completion alerts
- Report generation alerts
- AI agent recommendations
- Urgent notifications

**Implementation**:
- Push notification service (Firebase Cloud Messaging)
- Mobile app integration
- Web push notifications
- Notification preferences

---

### SMS Notifications

**Purpose**: Send critical notifications via SMS for users without app access.

**Use Cases**:
- Urgent alerts
- Attendance notifications
- Critical system notifications
- Emergency communications

**Implementation**:
- SMS service integration
- SMS templates
- SMS delivery tracking
- SMS cost optimization

---

### In-App Notifications

**Purpose**: Display notifications within the application for immediate visibility.

**Use Cases**:
- Task reminders
- Approval requests
- System updates
- User notifications

**Implementation**:
- In-app notification service
- Notification center
- Notification badges
- Notification history

---

## Notification Types

### System Notifications

**Purpose**: Notify users about system events and updates.

**Examples**:
- System maintenance
- Feature updates
- Security alerts
- Performance issues

---

### Business Notifications

**Purpose**: Notify users about business events and workflows.

**Examples**:
- Assessment completed
- Report generated
- Curriculum approved
- Intervention recommended

---

### AI Agent Notifications

**Purpose**: Notify users about AI agent activities and recommendations.

**Examples**:
- AI agent recommendation
- AI agent approval request
- AI agent completion
- AI agent error

---

## Notification Governance

### Notification Preferences

**Principle**: Users can control their notification preferences.

**Implementation**:
- User notification preferences
- Channel preferences (email, push, SMS, in-app)
- Frequency preferences
- Notification type preferences

---

### Notification Throttling

**Principle**: Notifications are throttled to prevent notification fatigue.

**Implementation**:
- Rate limiting per user
- Batching for similar notifications
- Digest mode for non-urgent notifications
- Quiet hours for non-urgent notifications

---

### Notification Compliance

**Principle**: Notifications comply with privacy and communication regulations.

**Implementation**:
- Opt-in/opt-out mechanisms
- Privacy compliance (GDPR, local regulations)
- Communication consent management
- Audit logging for notifications

---

# SECTION 13 — Analytics Architecture

## Overview

Analytics Architecture defines how NUSA collects, processes, and analyzes data to provide insights for decision-making, policy intelligence, and continuous improvement.

**Future Domain - Not Implemented in MVP Wave 1**

These capabilities are architectural placeholders for future platform evolution and are excluded from MVP delivery planning.

## Analytics Types

### Descriptive Analytics

**Purpose**: Describe what happened in the past.

**Implementation**:
- Data aggregation from event store
- Statistical analysis
- Data visualization
- Report generation

**Use Cases**:
- Student progress reports
- Teacher performance reports
- School quality reports
- National education statistics

---

### Diagnostic Analytics

**Purpose**: Explain why something happened.

**Implementation**:
- Root cause analysis
- Correlation analysis
- Drill-down capabilities
- Comparative analysis

**Use Cases**:
- Why student performance declined
- Why teacher workload increased
- Why school quality improved
- Why intervention succeeded

---

### Predictive Analytics

**Purpose**: Predict what will happen in the future.

**Implementation**:
- Machine learning models
- Predictive modeling
- Forecasting
- Scenario analysis

**Use Cases**:
- Predict student growth
- Predict teacher burnout
- Predict school quality trends
- Predict policy impact

---

### Prescriptive Analytics

**Purpose**: Recommend actions to achieve desired outcomes.

**Implementation**:
- Recommendation engines
- Optimization algorithms
- Decision support systems
- AI-powered suggestions

**Use Cases**:
- Recommend learning interventions
- Recommend professional development
- Recommend school improvement strategies
- Recommend policy adjustments

---

## Data Collection

### Event-Based Collection

**Principle**: Data is collected through domain events for comprehensive coverage.

**Implementation**:
- Event subscription for data collection
- Event aggregation for analytics
- Event filtering for relevant data
- Event transformation for analysis

**Data Sources**:
- Curriculum events (CP, TP, ATP, Modul Ajar)
- Assessment events (assessment, evidence, evaluation)
- Competency events (achievement, gap, prediction)
- Digital twin events (learning state, wellbeing state)

---

### Survey-Based Collection

**Principle**: Data is collected through surveys for subjective feedback.

**Implementation**:
- Survey design and distribution
- Survey response collection
- Survey data analysis
- Survey reporting

**Survey Types**:
- Student satisfaction surveys
- Teacher satisfaction surveys
- Parent satisfaction surveys
- School quality surveys

---

## Analytics Infrastructure

### Data Warehouse

**Purpose**: Centralized data storage for analytics.

**Implementation**:
- Data warehouse (PostgreSQL)
- Data lake (S3 for unstructured data)
- ETL pipelines for data transformation
- Data modeling for analytics

**Data Warehouse Schema**:
- Fact tables (assessment, learning, intervention)
- Dimension tables (student, teacher, school, time)
- Aggregate tables for performance
- Materialized views for common queries

---

### Analytics Engine

**Purpose**: Process and analyze data for insights.

**Implementation**:
- Analytics engine (Apache Spark)
- Machine learning platform (MLflow)
- Data visualization (Grafana, Tableau)
- Reporting engine (JasperReports)

**Analytics Capabilities**:
- Batch processing for historical analysis
- Stream processing for real-time analytics
- Machine learning for predictive analytics
- Visualization for data exploration

---

## Analytics Use Cases

### Student Analytics

**Purpose**: Provide insights into student learning and development.

**Future User Ecosystem - Not Implemented in MVP Wave 1**

Student and Parent capabilities are part of future platform expansion.

MVP Wave 1 supports only: Administrator, Teacher.

**Analytics**:
- Student progress tracking
- Competency development analysis
- Learning pattern analysis
- Intervention effectiveness analysis

---

### Teacher Analytics

**Purpose**: Provide insights into teacher performance and growth.

**Analytics**:
- Teacher performance tracking
- Workload analysis
- Professional development analysis
- Growth pattern analysis

---

### School Analytics

**Purpose**: Provide insights into school quality and improvement.

**Analytics**:
- School quality tracking
- SPMI cycle analysis
- Benchmarking analysis
- Improvement effectiveness analysis

---

### National Analytics

**Purpose**: Provide insights into national education outcomes and policy impact.

**Analytics**:
- National progress tracking
- Policy impact analysis
- Regional comparison analysis
- Trend analysis

---

# SECTION 14 — Competency Graph Module

## Overview

Competency Graph Module implements the Competency Intelligence Context, maintaining and leveraging the Competency Graph for intelligent recommendations and personalization.

**Future Strategic Domain - Not Implemented in MVP Wave 1**

Competency Graph remains part of long-term architecture and serves as a foundation for future Digital Twin and Lifelong Learning capabilities.

It is not included in MVP implementation planning.

## Data Architecture

### Graph Data

**Implementation**: Competency graph data is stored in PostgreSQL using adjacency list pattern.

**Graph Schema**:
- Nodes: Competencies, Skills, Knowledge, Dimensions
- Edges: Prerequisite, RelatedTo, PartOf, Achieves
- Properties: Level, Confidence, Timestamp

**Graph Structure**:
```
Graduate Profile Dimension
    ↓ (PartOf)
Subject Competency
    ↓ (PartOf)
Topic Competency
    ↓ (PartOf)
Skill
    ↓ (RelatedTo)
Knowledge
```

---

### Graph Operations

### Graph Construction

**Purpose**: Build and maintain the Competency Graph from curriculum and assessment data.

**Implementation**:
- Extract competencies from CP and TP
- Extract skills from curriculum
- Extract knowledge from curriculum
- Build relationships between competencies
- Build relationships to graduate profile dimensions

---

### Graph Query

**Purpose**: Query the Competency Graph for competency information.

**Implementation**:
- Query competency relationships
- Query competency prerequisites
- Query competency achievements
- Query competency gaps
- Query competency paths

**Query Examples**:
- Find all prerequisites for a competency
- Find all competencies achieved by a student
- Find competency gaps for a student
- Find optimal learning path for a competency

---

### Graph Analysis

**Purpose**: Analyze the Competency Graph for insights and predictions.

**Implementation**:
- Competency gap analysis
- Growth prediction
- Learning path optimization
- Competency clustering

**Analysis Types**:
- Gap analysis: Identify missing competencies
- Growth prediction: Predict future competency achievement
- Path optimization: Optimize learning paths
- Clustering: Group similar competencies

---

## AI Integration

### Competency Intelligence Agent

**Purpose**: AI agent for competency intelligence and recommendations.

**Responsibilities**:
- Competency gap analysis
- Competency growth prediction
- Learning path recommendation
- Competency clustering

**Inputs**:
- Assessment results
- Learning activities
- Competency graph state
- Student profile

**Outputs**:
- Competency gap report
- Growth prediction
- Learning path recommendation
- Competency cluster analysis

---

## Module APIs

### Graph Query API

**Purpose**: API for querying the Competency Graph.

**Endpoints**:
- GET /api/competency/graph/nodes: Query graph nodes
- GET /api/competency/graph/edges: Query graph edges
- GET /api/competency/graph/path: Query learning path
- GET /api/competency/graph/gaps: Query competency gaps

---

### Graph Analysis API

**Purpose**: API for analyzing the Competency Graph.

**Endpoints**:
- POST /api/competency/analysis/gap: Analyze competency gaps
- POST /api/competency/analysis/prediction: Predict competency growth
- POST /api/competency/analysis/path: Optimize learning path
- POST /api/competency/analysis/cluster: Cluster competencies

---

# SECTION 15 — Digital Twin Module

## Overview

Digital Twin Module implements the Digital Twin Context, maintaining real-time digital twins of students for personalized learning and intervention.

**Future Strategic Domain - Not Implemented in MVP Wave 1**

Digital Twin capabilities are part of long-term architecture and serve as a foundation for future personalized learning and intervention features.

It is not included in MVP implementation planning.

## Module Architecture

### Digital Twin Model

**Purpose**: Real-time representation of student learning, wellbeing, and character state.

**Digital Twin Components**:
- Learning State: Current learning progress and activities
- Wellbeing State: Physical and mental wellbeing indicators
- Character State: Character development indicators
- Personalization Profile: Learning preferences and needs

**State Updates**:
- Real-time updates from learning activities
- Periodic updates from assessments
- Event-driven updates from all contexts
- Predictive updates from AI models

---

### State Management

### Learning State

**Purpose**: Track student learning progress and activities.

**Components**:
- Current competencies
- Learning activities
- Assessment results
- Learning patterns

**Update Triggers**:
- Learning activity completed
- Assessment completed
- Competency achieved
- Intervention delivered

---

### Wellbeing State

**Purpose**: Track student physical and mental wellbeing.

**Components**:
- Physical health indicators
- Mental health indicators
- Stress levels
- Engagement levels

**Update Triggers**:
- Wellbeing assessments
- Behavioral observations
- Self-reports
- Teacher observations

---

### Character State

**Purpose**: Track student character development.

**Components**:
- Graduate profile dimension progress
- Character evidence
- Behavioral observations
- Service activities

**Update Triggers**:
- Character assessments
- Behavioral observations
- Service activities
- Teacher feedback

---

### Personalization Profile

**Purpose**: Maintain learning preferences and needs for personalization.

**Components**:
- Learning style
- Interests
- Strengths
- Areas for improvement

**Update Triggers**:
- Learning pattern analysis
- AI model updates
- Teacher feedback
- Student self-assessment

---

## AI Integration

### Student Coach Agent

**Purpose**: AI agent for personalized learning coaching.

**Responsibilities**:
- Learning state analysis
- Personalization recommendation
- Learning path adjustment
- Motivation support

**Inputs**:
- Digital twin state
- Learning activities
- Assessment results
- Student preferences

**Outputs**:
- Personalization recommendations
- Learning path adjustments
- Motivation messages
- Progress updates

---

### Intervention Agent

**Purpose**: AI agent for identifying and recommending interventions.

**Responsibilities**:
- Wellbeing monitoring
- Early warning detection
- Intervention recommendation
- Intervention tracking

**Inputs**:
- Digital twin state
- Wellbeing indicators
- Behavioral patterns
- Assessment results

**Outputs**:
- Intervention recommendations
- Early warning alerts
- Intervention tracking
- Progress monitoring

---

## Module APIs

### Digital Twin Query API

**Purpose**: API for querying digital twin state.

**Endpoints**:
- GET /api/digital-twin/{studentId}: Get digital twin state
- GET /api/digital-twin/{studentId}/learning: Get learning state
- GET /api/digital-twin/{studentId}/wellbeing: Get wellbeing state
- GET /api/digital-twin/{studentId}/character: Get character state

---

### Personalization API

**Purpose**: API for personalization recommendations.

**Endpoints**:
- POST /api/digital-twin/{studentId}/personalization: Get personalization recommendations
- POST /api/digital-twin/{studentId}/learning-path: Get learning path recommendations
- POST /api/digital-twin/{studentId}/resources: Get resource recommendations

---

# SECTION 16 — Lifelong Learning Record Module

## Overview

Lifelong Learning Record Module implements the Lifelong Learning Record Context, becoming the national lifelong learning record that connects all educational phases from early childhood through professional development.

## Module Architecture

### Record Aggregation

**Purpose**: Aggregate learning records from all educational phases.

**Educational Phases**:
- PAUD (Early Childhood Education)
- SD (Primary School)
- SMP (Junior Secondary School)
- SMA/SMK (Senior Secondary School)
- Higher Education (University/Vocational)
- Professional Development
- Lifelong Learning

**Aggregation Strategy**:
- Event-based aggregation from all contexts
- Phase transition tracking
- Record verification and validation
- Record deduplication

---

### Core Outputs

### Learning Passport

**Purpose**: Comprehensive record of all learning experiences across phases.

**Components**:
- Educational history
- Competencies
- Achievements
- Certifications
- Projects

**Use Cases**:
- College admissions
- Job applications
- Credential verification
- Lifelong learning planning

---

### Digital Competency Transcript

**Purpose**: Machine-readable transcript of competencies and achievements.

**Components**:
- Competency levels
- Achievement dates
- Verification status
- Evidence links

**Use Cases**:
- Automated credential verification
- Skill matching
- Workforce planning

---

### Graduate Readiness Score

**Purpose**: Aggregate score measuring readiness for next educational phase or career.

**Components**:
- Academic readiness
- Character readiness
- Career readiness
- Wellbeing readiness

**Use Cases**:
- Transition planning
- Intervention targeting
- Resource allocation

---

### Career Readiness Profile

**Purpose**: Detailed profile of career-related skills and readiness.

**Components**:
- Industry-specific skills
- Soft skills
- Work experience
- Career interests

**Use Cases**:
- Career guidance
- Job matching
- Workforce development

---

## AI Integration

### Lifelong Learning Agent

**Purpose**: AI agent for lifelong learning intelligence and recommendations.

**Responsibilities**:
- Competency gap analysis
- Career recommendation
- Personalized learning path
- Workforce readiness prediction
- Portfolio curation

**Inputs**:
- Learning records
- Competency data
- Career interests
- Industry trends

**Outputs**:
- Competency gap analysis
- Career recommendations
- Learning path recommendations
- Workforce readiness prediction
- Portfolio curation

---

## Module APIs

### Record Query API

**Purpose**: API for querying lifelong learning records.

**Endpoints**:
- GET /api/lifelong-learning/{studentId}: Get lifelong learning record
- GET /api/lifelong-learning/{studentId}/passport: Get learning passport
- GET /api/lifelong-learning/{studentId}/transcript: Get digital competency transcript
- GET /api/lifelong-learning/{studentId}/readiness: Get graduate readiness score

---

### Career API

**Purpose**: API for career readiness and recommendations.

**Endpoints**:
- GET /api/lifelong-learning/{studentId}/career: Get career readiness profile
- POST /api/lifelong-learning/{studentId}/career/recommendation: Get career recommendations
- POST /api/lifelong-learning/{studentId}/career/path: Get career path recommendations

---

# SECTION 17 — Cross-Cutting Modules

## Overview

Cross-Cutting Modules provide common capabilities across all Bounded Contexts, ensuring consistency and reducing duplication.

## Authentication Module

### Purpose

Provide centralized authentication and identity management for all modules.

### Responsibilities

- User authentication
- Token issuance and validation
- Session management
- Multi-factor authentication

### Implementation

- Custom JWT token issuance
- Token revocation
- Session management

---

## Audit Module

### Purpose

Provide centralized audit logging for compliance and security.

### Responsibilities

- Audit log collection
- Audit log storage
- Audit log query
- Audit log retention

### Implementation

- Centralized audit log service
- Audit log API
- Audit log retention policy
- Audit log query interface

---

## Logging Module

### Purpose

Provide centralized logging for observability and debugging.

### Responsibilities

- Log collection
- Log aggregation
- Log storage
- Log query

### Implementation

- Centralized logging service (ELK stack)
- Structured logging
- Log correlation
- Log retention policy

---

## Monitoring Module

### Purpose

Provide centralized monitoring for system health and performance.

### Responsibilities

- Metrics collection
- Metrics aggregation
- Alerting
- Dashboard

### Implementation

- Application logs
- Docker logs
- PostgreSQL monitoring
- RabbitMQ monitoring

---

## Search Module

### Purpose

Provide centralized search capabilities across all contexts.

### Responsibilities

- Index management
- Search query
- Search ranking
- Search analytics

### Implementation

- PostgreSQL full-text search
- Search API

---

# SECTION 18 — Architecture Validation

## Validation Checklist

### Foundation Document Alignment

**00A_NATIONAL_EDUCATION_DIRECTION_2045.md**
- [ ] Architecture supports Indonesia Emas 2045 vision
- [ ] Architecture enables human capital development
- [ ] Architecture aligns with national education direction

**00B_PRODUCT_VISION.md**
- [ ] Architecture enables 90% AI assistance
- [ ] Architecture maintains 10% human governance
- [ ] Architecture supports AI-Native Education Operating System

**00C_EDUCATION_OPERATING_SYSTEM_PRINCIPLES.md**
- [ ] Architecture is curriculum-centered
- [ ] Architecture prioritizes learning over administration
- [ ] Architecture follows architectural guardrails

**01_EDUCATION_DOMAIN_MODEL.md**
- [ ] All Bounded Contexts are traceable to domains
- [ ] Bounded Context boundaries align with domain boundaries
- [ ] Ubiquitous Language is consistent within contexts

**02_CAPABILITY_MODEL.md**
- [ ] All modules implement defined capabilities
- [ ] Module responsibilities align with capability definitions
- [ ] AI agents support capability automation targets

**03_BUSINESS_PROCESS_ARCHITECTURE.md**
- [ ] All modules implement defined processes
- [ ] Module workflows align with process definitions
- [ ] Event flows support process execution

**05_AI_ARCHITECTURE.md**
- [ ] All AI agents are hosted in appropriate modules
- [ ] AI agent orchestration is defined
- [ ] AI governance is implemented

---

### DDD Principles Validation

**Bounded Context Alignment**
- [ ] Each Bounded Context maps to one or more domains
- [ ] Context boundaries are explicit and enforced
- [ ] Ubiquitous Language is consistent within each context

**Context Map Validation**
- [ ] Upstream/Downstream relationships are defined
- [ ] Shared Kernel is minimized
- [ ] Anti-Corruption Layers are defined for external integration

**Module Architecture Validation**
- [ ] Modules align with domain concepts
- [ ] Aggregate roots have dedicated modules
- [ ] Module dependencies follow domain dependencies

**Module Architecture Validation**
- [ ] Each Bounded Context has dedicated module
- [ ] All modules are deployed together as a single application
- [ ] Modules communicate through well-defined interfaces

---

### Event-Driven Architecture Validation

**Event Design**
- [ ] Events represent domain-relevant state changes
- [ ] Events follow consistent schema
- [ ] Events are immutable

**Event Bus**
- [ ] Event bus is defined (RabbitMQ)
- [ ] Event topics are defined

**Domain Events**
- [ ] Domain events are defined
- [ ] Event topics are defined

---

### AI Integration Validation

**Agent Hosting**
- [ ] All AI agents are hosted in appropriate modules
- [ ] Agents have dedicated execution environments

**Agent Orchestration**
- [ ] AI Orchestration Module is defined
- [ ] Agent scheduling is defined
- [ ] Agent resource management is defined

**AI Governance**
- [ ] Human-in-the-loop validation points are defined
- [ ] AI safety measures are defined
- [ ] AI auditability is implemented

---

### Cross-Cutting Modules Validation

**Authentication**
- [ ] Authentication module is defined
- [ ] Custom JWT is implemented
- [ ] Multi-factor authentication is supported

**Audit**
- [ ] Audit module is defined
- [ ] Audit logging is centralized
- [ ] Audit retention policy is defined

**Logging**
- [ ] Logging module is defined
- [ ] Structured logging is implemented
- [ ] Log correlation is supported

**Monitoring**
- [ ] Monitoring module is defined
- [ ] Metrics collection is implemented
- [ ] Alerting is configured

**Search**
- [ ] Search module is defined
- [ ] PostgreSQL full-text search is implemented
- [ ] Search API is defined

---

# Future Architecture Evolution

The following technologies are part of long-term platform evolution and are not implemented in MVP Wave 1:

- ELK Stack
- Prometheus
- Grafana
- AlertManager
- Spark
- MLflow
- Tableau
- JasperReports

MVP monitoring relies on:

- Application Logs
- Docker Logs
- PostgreSQL Monitoring
- RabbitMQ Monitoring

---

# SECTION 19 — Conclusion

## Strategic Positioning

The Application Architecture (06) serves as the critical translation layer between business domains and software implementation. Using Domain-Driven Design (DDD) principles, it defines Bounded Contexts, Context Maps, Modules, Services, Event-Driven Architecture, APIs, and Cross-Cutting Services that implement the business processes defined in the Business Process Architecture (03), host the AI agents defined in the AI Architecture (05), and manage the data entities defined in the Data Architecture (04).

## Architecture Alignment

This Application Architecture is fully aligned with all foundation documents:

- **00A_NATIONAL_EDUCATION_DIRECTION_2045.md**: Architecture supports Indonesia Emas 2045 vision through human capital development
- **00B_PRODUCT_VISION.md**: Architecture enables 90% AI assistance while maintaining 10% human governance
- **00C_EDUCATION_OPERATING_SYSTEM_PRINCIPLES.md**: Architecture is curriculum-centered and prioritizes learning over administration
- **01_EDUCATION_DOMAIN_MODEL.md**: All Bounded Contexts are traceable to domains defined in 01
- **02_CAPABILITY_MODEL.md**: All modules implement capabilities defined in 02
- **03_BUSINESS_PROCESS_ARCHITECTURE.md**: All modules implement processes defined in 03
- **05_AI_ARCHITECTURE.md**: All modules host AI agents defined in 05

## Key Architectural Decisions

### Bounded Context Architecture

- 12 Bounded Contexts defined, each aligned with business domains
- Context Map defines upstream/downstream relationships
- Shared Kernel minimized to essential shared concepts
- Anti-Corruption Layers protect domain integrity from external integration

### Module Architecture

- Modules implement Bounded Contexts
- All modules are deployed together as a single application
- Modules communicate through REST APIs and domain events
- All modules use a shared PostgreSQL database

### Event-Driven Architecture

- Domain events for state changes
- Internal event bus for module communication
- Event flows support business process execution

### AI Integration

- AI agents hosted within modules
- AI Orchestration Module coordinates agent execution
- Human-in-the-loop validation points defined
- AI governance, safety, and auditability implemented

### Cross-Cutting Modules

- Authentication & Authorization Module for centralized identity management
- Audit Module for compliance and security
- Logging Module for observability and debugging
- Monitoring Module for system health and performance
- Search Module for centralized search capabilities

## Next Steps

This Application Architecture (06) provides the foundation for:

1. **MVP Architecture (07)**: Define the scope of the initial implementation
2. **SDLC Architecture (08)**: Define the development methodology to implement the architecture
3. **Implementation**: Build the modules and applications defined in this architecture
4. **Integration**: Integrate with external systems through Anti-Corruption Layers
5. **Deployment**: Deploy application using Docker Compose

## Architecture Governance

This Application Architecture is a foundation document and must be validated against all foundation documents before implementation. Any changes to this architecture must be traceable to changes in foundation documents and must maintain alignment with architectural principles.

---

**End of 06_APPLICATION_ARCHITECTURE.md**

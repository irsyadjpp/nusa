# 02_CAPABILITY_MODEL.md

## Foundation Document for Education Operating System Indonesia 2045

**Version**: 3.0
**Date**: June 2026
**Status**: FOUNDATION DOCUMENT - LOCKED
**Alignment**: 100% aligned with 00A_NATIONAL_EDUCATION_DIRECTION_2045.md, 00B_PRODUCT_VISION.md, 00C_EDUCATION_OPERATING_SYSTEM_PRINCIPLES.md, and 01_EDUCATION_DOMAIN_MODEL.md

**Purpose**: Define the complete capability model for Education Operating System (NUSA), serving as the official translation from Domain Architecture to Capability Architecture. This document is the single source of truth for Business Capability, AI Capability, Human Capability, MVP Scoping, Module Identification, Application Architecture, AI Architecture, and Process Architecture.

---

# SECTION 1 — Executive Summary

## Why Capability Model is Required

The Capability Model is the critical bridge between strategic vision and operational execution. While the Domain Architecture (01) defines WHAT exists in the education system, the Capability Model defines WHAT the platform must be able to DO to deliver value.

### Relationship with Domain Model

The Capability Model is the direct translation of the Domain Architecture:

```
Domain Architecture (01)
    ↓ translates to
Capability Model (02)
    ↓ enables
Business Process Architecture (03)
    ↓ enables
Data Architecture (04)
    ↓ enables
AI Architecture (05)
    ↓ enables
Application Architecture (06)
```

**Domain = What exists**: The stable business areas and entities in the education system (e.g., Curriculum Domain, Assessment Domain, Student Domain).

**Capability = What can be done**: The abilities and functions the platform must provide (e.g., Manage CP, Design Assessment, Track Student Growth).

**Process = How it works**: The workflows and procedures that execute capabilities (e.g., CP → TP transformation workflow, Assessment evidence collection workflow).

### Relationship with AI First Architecture

The Capability Model is AI-native by design. Every capability is evaluated for AI automation potential:

- **Human Driven**: Capabilities requiring human judgment, empathy, or ethical decision-making
- **AI Assisted**: Capabilities where AI provides recommendations and humans decide
- **AI Automated**: Capabilities where AI can execute autonomously within defined parameters

This evaluation ensures that the 90% AI Assistance, 10% Human Governance vision from 00B is operationalized at the capability level.

### Relationship with 90% Automation Vision

The Capability Model provides the foundation for achieving 90% AI-assisted operations:

- Each capability has an automation target
- AI agents are mapped to capabilities they can assist or automate
- Human validation points are defined for each capability
- Automation progress can be measured at capability level

### Strategic Importance

The Capability Model serves as:

1. **Single Source of Truth**: No capability should appear in subsequent architecture documents without being defined here
2. **MVP Scoping Foundation**: MVP scope is defined by selecting a subset of capabilities
3. **Module Identification**: Module boundaries are derived from capability boundaries
4. **AI Architecture Foundation**: AI agents are designed based on capability requirements
5. **Process Architecture Foundation**: Business processes are designed to execute capabilities
6. **Application Architecture Foundation**: Applications are built to deliver capabilities

## Core Principle

**Capabilities are the atomic units of value delivery in NUSA.**

Every feature, every module, every AI agent, and every business process must be traceable to one or more capabilities defined in this document. If a capability is not defined here, it should not exist in the system.

---

# SECTION 2 — Capability Architecture Principles

## Capability First

**Principle**: Platform is built based on capabilities, not features, not menus.

### Implementation
- All platform development starts with capability definition
- Features are implementations of capabilities
- Menus are interfaces to access capabilities
- No feature exists without a corresponding capability

### Rationale
- Capabilities are stable; features change
- Capabilities are strategic; features are tactical
- Capabilities enable strategic planning; features enable tactical execution
- Capability-first approach ensures long-term architectural coherence

---

## Outcome Driven

**Principle**: Every capability must contribute to the development of Profil Lulusan 8 Dimensi.

### Implementation
- Each capability is mapped to one or more Graduate Profile dimensions
- Capability success metrics include outcome contribution
- Capabilities without clear outcome contribution are rejected
- Outcome traceability is documented for each capability

### Rationale
- Ensures platform alignment with national education goals
- Guarantees that all capabilities serve educational purpose
- Enables measurement of platform impact on graduate outcomes
- Prevents capability creep without educational value

---

## AI Native

**Principle**: Every capability is evaluated for AI automation potential.

### Implementation
- Each capability is classified as: Human Driven, AI Assisted, or AI Automated
- AI automation targets are defined for each capability
- Human validation points are identified for AI-assisted capabilities
- AI agents are mapped to capabilities they can support

### Rationale
- Enables systematic achievement of 90% AI assistance target
- Ensures AI is considered from the beginning, not added later
- Provides clear framework for AI investment prioritization
- Maintains human governance while leveraging AI capabilities

---

## Human Governance

**Principle**: Strategic decisions remain under human authority.

### Implementation
- Critical decisions require human approval regardless of AI capability
- Human authority domains are explicitly defined
- AI agents are designed as co-pilots, not decision-makers
- Human accountability is maintained for all outcomes

### Rationale
- Preserves human authority in education
- Ensures ethical and pedagogical judgment remains human
- Maintains accountability for educational outcomes
- Aligns with Human Authority Principle from 00C

---

## Traceability

**Principle**: All capabilities must be traceable to Graduate Profile.

### Implementation
- Each capability documents its contribution to Graduate Profile dimensions
- Traceability chain is maintained: Capability → Competency → Graduate Profile → Indonesia Emas 2045
- Capabilities without traceability are rejected
- Traceability is auditable and verifiable

### Rationale
- Ensures all platform work contributes to national goals
- Enables measurement of platform impact on human capital development
- Prevents misalignment between platform capabilities and educational vision
- Supports outcome-based education paradigm

---

# SECTION 3 — Education Outcome Alignment

## Graduate Profile Dimension to Capability Group Mapping

This matrix ensures that all 8 dimensions of the Profil Lulusan are supported by platform capabilities.

| Graduate Profile Dimension | Capability Group | Key Capabilities |
| ------------------------- | ---------------- | ---------------- |
| **Dimensi 1: Beriman, Bertakwa kepada Tuhan YME, dan Berakhlak Mulia** | Character Development | Character Development Tracking, Values Integration, Ethical Decision Support |
| **Dimensi 2: Berkebinekaan Global** | Character Development | Cultural Competency, Global Perspective, Multilingual Support |
| **Dimensi 3: Gotong Royong** | Character Development | Collaboration Support, Community Engagement, Social Responsibility |
| **Dimensi 4: Mandiri** | Student Development | Self-Regulation, Personal Growth, Decision Making Support |
| **Dimensi 5: Bernalar Kritis** | Learning Delivery | Critical Thinking Support, Problem Solving, Inquiry-Based Learning |
| **Dimensi 6: Kreatif** | Learning Delivery | Creativity Support, Innovation, Project-Based Learning |
| **Dimensi 7: Bercocok Tanam** | Career Readiness | Career Exploration, Skill Development, Industry Connection |
| **Dimensi 8: Sehat Jasmani dan Rohani** | Student Wellbeing | Physical Health Tracking, Mental Health Support, Wellbeing Monitoring |

## Outcome Coverage Statement

All 8 dimensions of the Profil Lulusan are covered by platform capabilities:

### Complete Coverage
- Each dimension has dedicated capabilities
- Capabilities are designed to develop specific dimension competencies
- Assessment capabilities measure progress toward dimension outcomes
- Reporting capabilities communicate dimension development

### Cross-Dimension Integration
- Many capabilities contribute to multiple dimensions
- Integrated capabilities support holistic development
- Competency Graph enables cross-dimension traceability
- Learning Digital Twin tracks multi-dimensional growth

### Strategic Alignment
- Capability design is explicitly aligned with Graduate Profile
- Capability success metrics include dimension development
- AI recommendations consider dimension balance
- Platform evolution respects dimension priorities

---

# SECTION 4 — Capability Map Level 1

Level 1 capabilities represent the highest-level functional areas of the Education Operating System.

## 1. Graduate Profile Management

**Purpose**: Define, manage, and track the national graduate profile (8 Dimensions) as the North Star for all educational activities.

**Business Outcome**: All educational activities are aligned with and contribute to the development of the Profil Lulusan 8 Dimensi.

**Strategic Importance**: This is the foundation capability. Without clear graduate profile definition, all other capabilities lack strategic direction. This capability enables outcome traceability from daily activities to national goals.

---

## 2. Curriculum Management

**Purpose**: Define, manage, and distribute the national curriculum (CP, TP, ATP) as the single source of truth for learning content.

**Business Outcome**: Teachers have access to accurate, up-to-date curriculum materials that are aligned with the Graduate Profile.

**Strategic Importance**: Curriculum is the bridge between desired outcomes (Graduate Profile) and daily learning activities. This capability ensures that what is taught aligns with what is intended.

---

## 3. Learning Planning

**Purpose**: Enable teachers to create detailed learning plans (Modul Ajar) that translate curriculum into actionable classroom activities.

**Business Outcome**: Teachers can efficiently create high-quality lesson plans that are aligned with curriculum and personalized to student needs.

**Strategic Importance**: Learning planning is where curriculum becomes actionable. This capability directly impacts teacher workload and teaching quality.

---

## 4. Learning Delivery

**Purpose**: Support the delivery of learning activities in the classroom, including resource management, activity execution, and student engagement.

**Business Outcome**: Teachers can effectively deliver learning activities that engage students and achieve learning objectives.

**Strategic Importance**: Learning delivery is where teaching happens. This capability directly impacts student learning outcomes.

---

## 5. Assessment Management

**Purpose**: Design, deliver, and evaluate assessments to measure student learning and competency development.

**Business Outcome**: Teachers can efficiently create and administer assessments that accurately measure competency achievement.

**Strategic Importance**: Assessment provides the evidence of learning. This capability enables outcome measurement and data-driven decision-making.

---

## 6. Reporting Management

**Purpose**: Generate and communicate student progress reports to parents and stakeholders.

**Business Outcome**: Parents receive timely, accurate, and meaningful reports on their child's learning progress.

**Strategic Importance**: Reporting is the primary communication channel with parents. This capability enables parent engagement and partnership.

---

## 7. Student Development

**Purpose**: Monitor and support student growth across academic, social-emotional, and wellbeing dimensions.

**Business Outcome**: Students receive appropriate support to develop holistically across all dimensions.

**Strategic Importance**: Student development ensures that education goes beyond academics to support the whole child.

---

## 8. Teacher Development

**Purpose**: Support teacher professional growth through reflection, feedback, and professional development resources.

**Business Outcome**: Teachers continuously improve their practice and grow professionally.

**Strategic Importance**: Teacher quality is the most significant factor in student learning outcomes. This capability enables teacher capacity building.

---

## 9. School Improvement

**Purpose**: Enable schools to plan, implement, and monitor continuous quality improvement initiatives.

**Business Outcome**: Schools systematically improve their quality through data-driven improvement cycles.

**Strategic Importance**: School improvement ensures that educational institutions continuously evolve to better serve students.

---

## 10. Parent Partnership

**Purpose**: Facilitate effective communication and collaboration between schools and parents.

**Business Outcome**: Parents are actively engaged in their child's education and partner with teachers.

**Strategic Importance**: Parent engagement is a critical factor in student success. This capability enables home-school partnership.

---

## 11. Education Analytics

**Purpose**: Aggregate, analyze, and visualize education data to provide insights for decision-making.

**Business Outcome**: Education stakeholders have access to actionable insights derived from education data.

**Strategic Importance**: Analytics enables data-driven decision-making at all levels of the education system.

---

## 12. AI Orchestration

**Purpose**: Coordinate AI agents across domains to provide seamless AI assistance throughout the platform.

**Business Outcome**: Users receive consistent, high-quality AI assistance that enhances their work without adding burden.

**Strategic Importance**: AI orchestration is the foundation for achieving 90% AI assistance while maintaining human governance.

---

## 13. Competency Intelligence

**Purpose**: Maintain and leverage the National Competency Graph to enable intelligent recommendations and personalization.

**Business Outcome**: The Competency Graph serves as the educational brain of NUSA, enabling intelligent learning pathways and personalization.

**Strategic Importance**: Competency Intelligence is the foundation for AI-native education at scale.

---

## 14. Lifelong Learning Record

**Purpose**: Maintain a comprehensive longitudinal record of an individual's learning journey across all educational phases.

**Business Outcome**: Individuals have a complete, portable record of their learning that supports career readiness and lifelong learning.

**Strategic Importance**: Lifelong Learning Record enables national human capital intelligence and workforce planning.

---

## 15. Credential & Achievement

**Purpose**: Manage and verify credentials, certifications, and achievements throughout an individual's learning journey.

**Business Outcome**: Individuals have verified credentials that are recognized across educational phases and by employers.

**Strategic Importance**: Credentials provide the evidence of competency achievement that enables career progression.

---

## 16. Governance & Compliance

**Purpose**: Ensure platform operations comply with education policies, standards, and regulations.

**Business Outcome**: Platform operations are compliant with national education policies and standards.

**Strategic Importance**: Governance and compliance ensure platform legitimacy and sustainability.

---

# SECTION 5 — Capability Decomposition Level 2

Level 2 capabilities decompose Level 1 capabilities into more specific functional areas.

## 1. Graduate Profile Management

- **Graduate Profile Definition**: Define and maintain the 8 dimensions of the Profil Lulusan
- **Outcome Mapping**: Map learning outcomes to Graduate Profile dimensions
- **Progress Tracking**: Track progress toward Graduate Profile development at individual and aggregate levels
- **Alignment Verification**: Verify that curriculum, learning, and assessment align with Graduate Profile

## 2. Curriculum Management

- **CP Management**: Define and maintain Capaian Pembelajaran (CP)
- **TP Management**: Define and maintain Tujuan Pembelajaran (TP)
- **ATP Management**: Define and maintain Alur Tujuan Pembelajaran (ATP)
- **Curriculum Mapping**: Map curriculum across grades and subjects
- **Curriculum Alignment**: Verify alignment with Graduate Profile and standards
- **Curriculum Review**: Support curriculum review and update processes

## 3. Learning Planning

- **Lesson Planning**: Create detailed lesson plans (Modul Ajar)
- **Resource Matching**: Match learning resources to lesson plans
- **Differentiation Planning**: Plan for differentiated instruction
- **Activity Sequencing**: Sequence learning activities optimally
- **Plan Sharing**: Share and collaborate on lesson plans

## 4. Learning Delivery

- **Session Management**: Manage learning sessions and schedules
- **Activity Execution**: Execute learning activities in the classroom
- **Resource Access**: Provide access to learning resources during delivery
- **Student Engagement**: Monitor and support student engagement
- **Real-time Adaptation**: Adapt delivery based on real-time feedback

## 5. Assessment Management

- **Assessment Design**: Create assessments aligned with learning objectives
- **Assessment Delivery**: Administer assessments to students
- **Evidence Collection**: Collect evidence of student learning
- **Evidence Evaluation**: Evaluate evidence against competency criteria
- **Competency Validation**: Validate competency achievement
- **Assessment Moderation**: Ensure assessment quality and consistency

## 6. Reporting Management

- **Report Generation**: Generate student progress reports
- **Narrative Writing**: Create narrative descriptions of student progress
- **Multi-language Support**: Support reports in multiple languages
- **Parent Communication**: Communicate reports to parents
- **Trend Analysis**: Analyze progress trends over time

## 7. Student Development

- **Growth Tracking**: Track student growth across multiple dimensions
- **Wellbeing Monitoring**: Monitor student physical and mental wellbeing
- **Intervention Planning**: Plan and track interventions for student support
- **Strength Identification**: Identify and develop student strengths
- **Gap Analysis**: Identify learning gaps and remediation needs

## 8. Teacher Development

- **Reflection Support**: Support teacher reflection on practice
- **Feedback Collection**: Collect feedback on teacher performance
- **PD Resource Access**: Provide access to professional development resources
- **Growth Planning**: Plan and track teacher professional growth
- **Collaboration Support**: Enable teacher collaboration and knowledge sharing

## 9. School Improvement

- **Quality Planning**: Plan school quality improvement initiatives
- **Data Analysis**: Analyze school performance data
- **Benchmarking**: Benchmark against other schools
- **Improvement Tracking**: Track improvement initiative progress
- **SPMI Support**: Support School Quality Management System (SPMI) processes

## 10. Parent Partnership

- **Communication Management**: Manage communication with parents
- **Engagement Tracking**: Track parent engagement levels
- **Resource Sharing**: Share resources with parents
- **Feedback Collection**: Collect feedback from parents
- **Event Management**: Manage parent-school events

## 11. Education Analytics

- **Data Aggregation**: Aggregate data from across the platform
- **Descriptive Analytics**: Provide descriptive analysis of education data
- **Predictive Analytics**: Provide predictive insights
- **Prescriptive Analytics**: Provide actionable recommendations
- **Visualization**: Visualize data for stakeholders

## 12. AI Orchestration

- **Agent Coordination**: Coordinate AI agents across domains
- **Workflow Orchestration**: Orchestrate AI-assisted workflows
- **Human Validation**: Manage human validation points
- **Performance Monitoring**: Monitor AI agent performance
- **Governance Enforcement**: Enforce AI governance rules

## 13. Competency Intelligence

- **Competency Graph Maintenance**: Maintain the National Competency Graph
- **Relationship Mapping**: Map relationships between competencies
- **Pathway Analysis**: Analyze optimal learning pathways
- **Gap Analysis**: Identify competency gaps
- **Recommendation Engine**: Generate personalized recommendations

## 14. Lifelong Learning Record

- **Record Aggregation**: Aggregate learning records across phases
- **Competency Tracking**: Track competency development over time
- **Portfolio Management**: Manage learning portfolios
- **Credential Verification**: Verify credentials and achievements
- **Portability**: Ensure record portability across systems

## 15. Credential & Achievement

- **Credential Issuance**: Issue credentials for achievements
- **Credential Verification**: Verify credential authenticity
- **Badge Management**: Manage digital badges and micro-credentials
- **Achievement Tracking**: Track achievements and milestones
- **Recognition**: Recognize and celebrate achievements

## 16. Governance & Compliance

- **Policy Management**: Manage education policies and standards
- **Compliance Monitoring**: Monitor compliance with policies
- **Audit Support**: Support audit processes
- **Risk Management**: Identify and manage risks
- **Access Control**: Manage access control and permissions

---

# SECTION 6 — Capability Decomposition Level 3

Level 3 capabilities define operational capabilities that can be directly implemented as modules or features.

## Critical Level 3 Capabilities

### Assessment Evidence Collection

- **Collect Evidence**: Collect various types of evidence (work samples, projects, observations)
- **Validate Evidence**: Validate evidence authenticity and quality
- **Tag Evidence**: Tag evidence with competency metadata
- **Link Evidence to Competency**: Link evidence to specific competencies
- **Store Evidence**: Store evidence securely with proper metadata
- **Retrieve Evidence**: Retrieve evidence for assessment and reporting

### Lesson Plan Generation

- **Generate Draft**: Generate initial lesson plan draft based on curriculum
- **Customize Plan**: Customize plan for specific classroom context
- **Add Resources**: Add and link learning resources
- **Sequence Activities**: Sequence learning activities optimally
- **Differentiate**: Plan for differentiated instruction
- **Review and Approve**: Review and approve final lesson plan

### Competency Evaluation

- **Analyze Evidence**: Analyze collected evidence against competency criteria
- **Determine Mastery**: Determine level of competency mastery
- **Provide Feedback**: Generate feedback for improvement
- **Track Progress**: Track progress over time
- **Identify Gaps**: Identify learning gaps
- **Recommend Interventions**: Recommend appropriate interventions

### Report Narrative Generation

- **Analyze Data**: Analyze student performance data
- **Identify Strengths**: Identify student strengths and achievements
- **Identify Areas for Improvement**: Identify areas needing support
- **Generate Narrative**: Generate narrative description of progress
- **Translate**: Translate to multiple languages
- **Customize**: Customize for specific parent context

### AI Agent Coordination

- **Route Request**: Route user requests to appropriate AI agent
- **Coordinate Agents**: Coordinate multiple agents for complex tasks
- **Manage Context**: Maintain context across agent interactions
- **Validate Output**: Validate AI agent outputs
- **Request Human Approval**: Request human approval when required
- **Log Interactions**: Log all AI agent interactions

### Competency Graph Query

- **Query Competency**: Query competency information
- **Trace Relationships**: Trace competency relationships
- **Identify Prerequisites**: Identify prerequisite competencies
- **Suggest Pathways**: Suggest optimal learning pathways
- **Validate Alignment**: Validate alignment with standards
- **Update Graph**: Update competency graph as needed

---

# SECTION 7 — Core vs Supporting Capabilities

## Core Capabilities

Capabilities that create the primary value of the Education Operating System. These are essential for the platform to fulfill its mission.

### List of Core Capabilities

1. **Graduate Profile Management** - Defines the North Star for all education
2. **Curriculum Management** - Provides the single source of truth for learning content
3. **Learning Planning** - Enables teachers to translate curriculum into action
4. **Learning Delivery** - Supports the core teaching and learning process
5. **Assessment Management** - Provides evidence of learning
6. **Reporting Management** - Communicates progress to stakeholders

**Strategic Importance**: Without these capabilities, the platform cannot deliver its core educational value. These capabilities must be implemented first and maintained at the highest quality.

---

## Strategic Capabilities

Capabilities that create competitive advantage and enable the platform to achieve its vision of AI-native education.

### List of Strategic Capabilities

1. **AI Orchestration** - Enables 90% AI assistance while maintaining human governance
2. **Competency Intelligence** - Provides the educational brain for personalization
3. **Lifelong Learning Record** - Enables national human capital intelligence
4. **Student Development** - Enables holistic student development
5. **Teacher Development** - Enables continuous teacher capacity building

**Strategic Importance**: These capabilities differentiate NUSA from traditional education systems and enable the achievement of the Indonesia Emas 2045 vision.

---

## Supporting Capabilities

Capabilities that support the operation of core and strategic capabilities but do not directly create educational value.

### List of Supporting Capabilities

1. **Parent Partnership** - Supports core education through parent engagement
2. **School Improvement** - Supports core education through quality improvement
3. **Education Analytics** - Supports decision-making across all capabilities
4. **Governance & Compliance** - Ensures platform legitimacy and sustainability

**Strategic Importance**: These capabilities are necessary for effective operation but can be implemented with simpler solutions initially and enhanced over time.

---

## Platform Capabilities

Capabilities that orchestrate the system and enable other capabilities to function.

### List of Platform Capabilities

1. **Credential & Achievement** - Provides credentialing infrastructure
2. **AI Orchestration** (also Strategic) - Provides AI infrastructure
3. **Competency Intelligence** (also Strategic) - Provides intelligence infrastructure

**Strategic Importance**: These capabilities provide the foundational infrastructure that enables other capabilities to function effectively.

---

# SECTION 8 — AI Capability Model

For each capability, we define the level of AI automation and human involvement.

## AI Capability Matrix

| Capability | Human | AI Assisted | AI Automated | Automation Target |
| ---------- | ----- | ----------- | ------------ | ---------------- |
| **Graduate Profile Definition** | 100% | 0% | 0% | 0% |
| **Outcome Mapping** | 80% | 20% | 0% | 20% |
| **Progress Tracking** | 20% | 80% | 0% | 80% |
| **CP Management** | 70% | 30% | 0% | 30% |
| **TP Management** | 30% | 70% | 0% | 70% |
| **ATP Management** | 20% | 80% | 0% | 80% |
| **Lesson Planning** | 20% | 80% | 0% | 80% |
| **Resource Matching** | 10% | 90% | 0% | 90% |
| **Activity Sequencing** | 20% | 80% | 0% | 80% |
| **Assessment Design** | 30% | 70% | 0% | 70% |
| **Assessment Delivery** | 50% | 50% | 0% | 50% |
| **Evidence Collection** | 40% | 60% | 0% | 60% |
| **Evidence Evaluation** | 20% | 80% | 0% | 80% |
| **Competency Validation** | 30% | 70% | 0% | 70% |
| **Report Generation** | 10% | 90% | 0% | 90% |
| **Narrative Writing** | 10% | 90% | 0% | 90% |
| **Growth Tracking** | 20% | 80% | 0% | 80% |
| **Wellbeing Monitoring** | 50% | 50% | 0% | 50% |
| **Reflection Support** | 40% | 60% | 0% | 60% |
| **Data Aggregation** | 10% | 90% | 0% | 90% |
| **Descriptive Analytics** | 20% | 80% | 0% | 80% |
| **Predictive Analytics** | 40% | 60% | 0% | 60% |
| **Agent Coordination** | 20% | 80% | 0% | 80% |
| **Competency Graph Query** | 10% | 90% | 0% | 90% |
| **Pathway Analysis** | 20% | 80% | 0% | 80% |
| **Record Aggregation** | 10% | 90% | 0% | 90% |
| **Credential Verification** | 30% | 70% | 0% | 70% |

## AI Automation Principles

### Human Driven (0-20% AI)
Capabilities requiring human judgment, ethical decision-making, or definition of strategic direction. AI provides minimal support.

### AI Assisted (60-90% AI)
Capabilities where AI provides significant assistance but humans make final decisions. AI handles routine work, humans handle judgment and validation.

### AI Automated (Future State)
Capabilities where AI can execute autonomously within defined parameters. Currently, no capabilities are fully automated to maintain human governance.

## Overall Target

**90% AI Assistance, 10% Human Governance**

The weighted average across all capabilities targets 90% AI assistance, with human governance maintained at critical decision points.

---

# SECTION 9 — AI Agent Capability Matrix

## Official AI Agents of NUSA Platform

### 1. Curriculum Agent

**Purpose**: Assist with curriculum-related tasks including CP → TP transformation, TP → ATP generation, and curriculum alignment verification.

**Inputs**:
- Graduate Profile (8 Dimensions)
- CP (Capaian Pembelajaran)
- Subject and grade context
- Teacher preferences

**Outputs**:
- TP (Tujuan Pembelajaran)
- ATP (Alur Tujuan Pembelajaran)
- Alignment reports
- Curriculum recommendations

**Decisions**:
- Sequencing of learning objectives
- Prerequisite relationships
- Alignment with standards

**Human Approval Requirement**: Teacher must review and approve all curriculum outputs before use.

---

### 2. Lesson Planning Agent

**Purpose**: Assist teachers in creating detailed lesson plans (Modul Ajar) that are aligned with curriculum and personalized to student needs.

**Inputs**:
- ATP (Alur Tujuan Pembelajaran)
- Student data and learning profiles
- Available resources
- Teacher preferences

**Outputs**:
- Modul Ajar draft
- Resource recommendations
- Activity sequences
- Differentiation suggestions

**Decisions**:
- Activity selection and sequencing
- Resource matching
- Differentiation strategies

**Human Approval Requirement**: Teacher must review, customize, and approve lesson plans before use.

---

### 3. Assessment Agent

**Purpose**: Assist with assessment design, evidence evaluation, and competency validation.

**Inputs**:
- TP (Tujuan Pembelajaran)
- Competency criteria
- Student evidence
- Assessment rubrics

**Outputs**:
- Assessment item recommendations
- Evidence evaluation results
- Competency mastery levels
- Feedback for improvement

**Decisions**:
- Assessment item selection
- Evidence quality evaluation
- Mastery level determination

**Human Approval Requirement**: Teacher must review and approve all assessment evaluations and competency validations.

---

### 4. Evidence Agent

**Purpose**: Assist with evidence collection, validation, tagging, and linkage to competencies.

**Inputs**:
- Student work samples
- Project artifacts
- Observation notes
- Competency definitions

**Outputs**:
- Validated evidence
- Evidence tags
- Competency linkages
- Portfolio entries

**Decisions**:
- Evidence quality validation
- Competency matching
- Portfolio organization

**Human Approval Requirement**: Teacher must review evidence classifications and competency linkages.

---

### 5. Student Growth Agent

**Purpose**: Monitor student growth, identify learning gaps, and recommend interventions.

**Inputs**:
- Assessment data
- Learning progress data
- Wellbeing indicators
- Historical performance

**Outputs**:
- Growth trajectories
- Gap analysis
- Intervention recommendations
- Progress reports

**Decisions**:
- Growth trend analysis
- Gap identification
- Intervention prioritization

**Human Approval Requirement**: Teacher must review intervention recommendations and decide on implementation.

---

### 6. Teacher Coach Agent

**Purpose**: Support teacher reflection, provide feedback, and recommend professional development resources.

**Inputs**:
- Teacher performance data
- Student outcomes
- Classroom observations
- Professional goals

**Outputs**:
- Reflection summaries
- Performance insights
- PD recommendations
- Growth suggestions

**Decisions**:
- Strength identification
- Area for improvement identification
- PD resource matching

**Human Approval Requirement**: Teacher must review feedback and decide on professional development actions.

---

### 7. Parent Communication Agent

**Purpose**: Assist with parent communication, report generation, and engagement tracking.

**Inputs**:
- Student progress data
- Communication history
- Parent preferences
- School policies

**Outputs**:
- Progress reports
- Communication drafts
- Engagement insights
- Translation to multiple languages

**Decisions**:
- Report content generation
- Communication timing
- Engagement strategy

**Human Approval Requirement**: Teacher must review and approve all communications before sending to parents.

---

### 8. Analytics Agent

**Purpose**: Aggregate, analyze, and visualize education data to provide actionable insights.

**Inputs**:
- Platform data across all domains
- External benchmarks
- Historical data
- Stakeholder queries

**Outputs**:
- Descriptive analytics
- Predictive insights
- Prescriptive recommendations
- Data visualizations

**Decisions**:
- Data aggregation strategies
- Analysis model selection
- Insight prioritization

**Human Approval Requirement**: School leaders must review analytics insights before making strategic decisions.

---

### 9. Competency Intelligence Agent

**Purpose**: Maintain and leverage the National Competency Graph to enable intelligent recommendations and personalization.

**Inputs**:
- Competency definitions
- Student competency data
- Learning progress data
- Career requirements

**Outputs**:
- Learning pathway recommendations
- Competency gap analysis
- Personalization suggestions
- Alignment reports

**Decisions**:
- Pathway optimization
- Gap identification
- Personalization strategy

**Human Approval Requirement**: Teachers must review personalization recommendations before implementation.

---

### 10. Lifelong Learning Agent

**Purpose**: Maintain the Lifelong Learning Record and provide career readiness insights.

**Inputs**:
- Learning records across phases
- Competency achievements
- Credential data
- Industry requirements

**Outputs**:
- Lifelong learning record
- Career readiness assessment
- Skill gap analysis
- Credential verification

**Decisions**:
- Record aggregation
- Career pathway mapping
- Credential validation

**Human Approval Requirement**: Individuals must review their learning records and career recommendations.

---

### 11. Governance Agent

**Purpose**: Enforce AI governance rules, monitor AI agent performance, and ensure human oversight.

**Inputs**:
- AI agent logs
- Governance rules
- Human approval records
- Performance metrics

**Outputs**:
- Compliance reports
- Performance insights
- Risk alerts
- Governance recommendations

**Decisions**:
- Compliance validation
- Risk assessment
- Governance enforcement

**Human Approval Requirement**: Governance board must review compliance reports and risk assessments.

---

# SECTION 10 — Human Capability Model

## Student Capabilities

### Responsibilities
- Engage in learning activities
- Complete assessments
- Reflect on learning
- Set learning goals
- Seek help when needed

### Decisions
- Learning path choices (within guidance)
- Resource selection
- Collaboration choices
- Goal setting
- Help-seeking

### Governance Rights
- Access to personal learning data
- Privacy protection
- Voice in learning process
- Appeal of assessments
- Participation in improvement

---

## Teacher Capabilities

### Responsibilities
- Plan and deliver learning
- Assess student learning
- Provide feedback
- Communicate with parents
- Engage in professional development

### Decisions
- Lesson planning
- Assessment design
- Intervention selection
- Communication strategy
- Professional growth

### Governance Rights
- Curriculum input
- Assessment authority
- Pedagogical autonomy
- Professional development choice
- School improvement participation

---

## Parent Capabilities

### Responsibilities
- Support child's learning
- Communicate with teachers
- Participate in school activities
- Provide learning environment
- Monitor child's progress

### Decisions
- Engagement level
- Communication preferences
- Support strategies
- Participation choices
- Advocacy

### Governance Rights
- Access to child's data
- Communication with school
- Participation in governance
- Feedback on school
- Appeal of decisions

---

## School Leader Capabilities

### Responsibilities
- Lead school improvement
- Manage school operations
- Support teacher development
- Ensure quality standards
- Engage with community

### Decisions
- Improvement strategy
- Resource allocation
- Staffing decisions
- Program selection
- Community engagement

### Governance Rights
- School autonomy
- Budget authority
- Staffing authority
- Policy implementation
- Accountability for outcomes

---

## Government Capabilities

### Responsibilities
- Set education policy
- Define standards
- Provide resources
- Monitor quality
- Ensure equity

### Decisions
- Policy formulation
- Standard setting
- Resource allocation
- Quality monitoring
- Intervention decisions

### Governance Rights
- Policy authority
- Standard authority
- Funding authority
- Regulatory authority
- Accountability for national outcomes

---

## Industry Capabilities

### Responsibilities
- Define skill requirements
- Provide learning opportunities
- Participate in curriculum development
- Hire and develop talent
- Provide feedback on education

### Decisions
- Skill requirements
- Partnership strategies
- Hiring criteria
- Training investments
- Feedback provision

### Governance Rights
- Input on curriculum
- Participation in assessment
- Access to talent data
- Partnership terms
- Quality expectations

---

# SECTION 11 — Capability Ownership Matrix

| Capability | Domain Owner | Human Owner | AI Agent |
| ---------- | ------------ | ----------- | -------- |
| **Graduate Profile Definition** | Graduate Profile Domain | Ministry of Education | None |
| **Outcome Mapping** | Graduate Profile Domain | Curriculum Expert | Curriculum Agent |
| **Progress Tracking** | Graduate Profile Domain | Teacher | Analytics Agent |
| **CP Management** | Curriculum Domain | Ministry of Education | None |
| **TP Management** | Curriculum Domain | Curriculum Expert | Curriculum Agent |
| **ATP Management** | Curriculum Domain | Teacher | Curriculum Agent |
| **Lesson Planning** | Learning Planning Domain | Teacher | Lesson Planning Agent |
| **Resource Matching** | Learning Planning Domain | Teacher | Lesson Planning Agent |
| **Activity Sequencing** | Learning Planning Domain | Teacher | Lesson Planning Agent |
| **Assessment Design** | Assessment Domain | Teacher | Assessment Agent |
| **Assessment Delivery** | Assessment Domain | Teacher | Assessment Agent |
| **Evidence Collection** | Assessment Domain | Teacher | Evidence Agent |
| **Evidence Evaluation** | Assessment Domain | Teacher | Evidence Agent |
| **Competency Validation** | Assessment Domain | Teacher | Assessment Agent |
| **Report Generation** | Reporting Domain | Teacher | Parent Communication Agent |
| **Narrative Writing** | Reporting Domain | Teacher | Parent Communication Agent |
| **Growth Tracking** | Student Wellbeing Domain | Teacher | Student Growth Agent |
| **Wellbeing Monitoring** | Student Wellbeing Domain | Counselor | Student Growth Agent |
| **Reflection Support** | Teacher Growth Domain | Teacher | Teacher Coach Agent |
| **Data Aggregation** | Education Data Domain | Data Analyst | Analytics Agent |
| **Descriptive Analytics** | Education Data Domain | Data Analyst | Analytics Agent |
| **Predictive Analytics** | Education Data Domain | Data Analyst | Analytics Agent |
| **Agent Coordination** | AI Orchestration Domain | AI Engineer | Governance Agent |
| **Competency Graph Query** | Lifelong Learning Domain | Competency Expert | Competency Intelligence Agent |
| **Pathway Analysis** | Lifelong Learning Domain | Competency Expert | Competency Intelligence Agent |
| **Record Aggregation** | Lifelong Learning Domain | System Admin | Lifelong Learning Agent |
| **Credential Verification** | Lifelong Learning Domain | Credential Authority | Lifelong Learning Agent |

---

# SECTION 12 — Capability Dependency Map

## Core Education Flow Dependencies

```
Graduate Profile Management
    ↓ defines
Curriculum Management
    ↓ enables
Learning Planning
    ↓ enables
Learning Delivery
    ↓ generates
Assessment Management
    ↓ produces
Reporting Management
    ↓ informs
Student Development
    ↓ feeds
Lifelong Learning Record
    ↓ enables
Human Capital Outcome
```

## Growth and Improvement Dependencies

```
Assessment Management
    ↓ provides data for
Student Development
    ↓ informs
Teacher Development
    ↓ enables
School Improvement
    ↓ monitored by
Governance & Compliance
```

## Intelligence and Analytics Dependencies

```
All Operational Capabilities
    ↓ provide data to
Education Analytics
    ↓ enables
Competency Intelligence
    ↓ powers
AI Orchestration
    ↓ governed by
Governance & Compliance
```

## Partnership and Communication Dependencies

```
Reporting Management
    ↓ enables
Parent Partnership
    ↓ supports
Student Development
    ↓ informs
School Improvement
```

## Credential and Achievement Dependencies

```
Assessment Management
    ↓ validates
Competency Achievement
    ↓ recorded in
Credential & Achievement
    ↓ feeds
Lifelong Learning Record
    ↓ enables
Career Readiness
```

## Dependency Principles

### Upstream Dependency
- Downstream capabilities depend on upstream capabilities
- Upstream capabilities can function independently
- Downstream capabilities cannot function without upstream

### Data Flow Direction
- Data flows from upstream to downstream capabilities
- Downstream capabilities consume upstream data
- Upstream capabilities are not dependent on downstream

### Implementation Sequencing
- Upstream capabilities should be implemented before downstream
- This ensures data availability for dependent capabilities
- Phased implementation follows dependency hierarchy

---

# SECTION 13 — Capability Maturity Model

## Maturity Levels

### Level 1 — Manual
- Processes are entirely manual
- No digital support
- High human effort
- Low consistency

### Level 2 — Digital
- Processes are digitized
- Basic digital tools available
- Reduced human effort
- Improved consistency

### Level 3 — Assisted
- AI provides assistance
- Humans make decisions
- Significant automation of routine tasks
- High consistency

### Level 4 — Intelligent
- AI provides intelligent recommendations
- Humans validate and approve
- Most routine work automated
- Very high consistency
- Personalized experiences

### Level 5 — Autonomous
- AI executes autonomously within parameters
- Humans monitor and intervene only when needed
- Near-complete automation
- Perfect consistency
- Highly personalized

## NUSA Maturity Targets

### Target State
- **Majority of capabilities at Level 4 (Intelligent)**
- **Critical decision capabilities at Level 3 (Assisted)**
- **Strategic definition capabilities at Level 1-2 (Manual/Digital)**

### Capability Maturity Targets

| Capability | Current State | Target State | Timeline |
| ---------- | ------------- | ------------ | -------- |
| **Graduate Profile Definition** | Level 1 | Level 2 | Phase 1 |
| **Curriculum Management** | Level 2 | Level 3 | Phase 1 |
| **Learning Planning** | Level 1 | Level 4 | Phase 1 |
| **Learning Delivery** | Level 2 | Level 3 | Phase 2 |
| **Assessment Management** | Level 2 | Level 4 | Phase 1 |
| **Reporting Management** | Level 1 | Level 4 | Phase 1 |
| **Student Development** | Level 1 | Level 3 | Phase 2 |
| **Teacher Development** | Level 1 | Level 3 | Phase 2 |
| **Education Analytics** | Level 1 | Level 4 | Phase 2 |
| **AI Orchestration** | Level 1 | Level 4 | Phase 1 |
| **Competency Intelligence** | Level 1 | Level 4 | Phase 2 |
| **Lifelong Learning Record** | Level 1 | Level 3 | Phase 3 |

---

# SECTION 14 — MVP Capability Prioritization

## Wave 1 (30 Day MVP)

Capabilities that must be built to deliver immediate value and reduce teacher burden.

### Capabilities

1. **Graduate Profile Management** - Graduate Profile Definition, Outcome Mapping
2. **Curriculum Management** - CP Management, TP Management, ATP Management
3. **Learning Planning** - Lesson Planning, Resource Matching
4. **Assessment Management** - Assessment Design, Evidence Collection, Evidence Evaluation
5. **Reporting Management** - Report Generation, Narrative Writing
6. **AI Orchestration** - Agent Coordination, Human Validation

### Rationale
- These capabilities directly reduce teacher workload
- They form the core education flow
- They provide immediate value to teachers
- They are foundational for all other capabilities
- They align with Phase 1 MVP priority: Reducing Teacher Burden

---

## Wave 2

Capabilities that enhance the core capabilities and enable deeper functionality.

### Capabilities

1. **Learning Delivery** - Session Management, Activity Execution
2. **Student Development** - Growth Tracking, Wellbeing Monitoring
3. **Teacher Development** - Reflection Support, PD Resource Access
4. **Education Analytics** - Data Aggregation, Descriptive Analytics
5. **Competency Intelligence** - Competency Graph Query, Pathway Analysis

### Rationale
- These capabilities build on Wave 1 foundations
- They enable deeper personalization and analytics
- They support holistic student development
- They enhance teacher capacity building
- They align with Phase 2 focus: Deep Learning and Analytics

---

## Wave 3

Capabilities that enable longitudinal tracking and lifelong learning.

### Capabilities

1. **Lifelong Learning Record** - Record Aggregation, Competency Tracking
2. **Credential & Achievement** - Credential Issuance, Credential Verification
3. **Parent Partnership** - Communication Management, Engagement Tracking
4. **School Improvement** - Quality Planning, Data Analysis

### Rationale
- These capabilities enable longitudinal value
- They support lifelong learning and career readiness
- They enhance parent engagement
- They support school quality improvement
- They align with Phase 3 focus: Lifelong Learning and Quality

---

## Wave 4

Capabilities that provide advanced intelligence and optimization.

### Capabilities

1. **Education Analytics** - Predictive Analytics, Prescriptive Analytics
2. **AI Orchestration** - Performance Monitoring, Governance Enforcement
3. **Governance & Compliance** - Policy Management, Compliance Monitoring
4. **Competency Intelligence** - Gap Analysis, Recommendation Engine

### Rationale
- These capabilities provide advanced intelligence
- They enable predictive and prescriptive analytics
- They ensure governance and compliance
- They optimize system performance
- They align with Phase 4 focus: Advanced Intelligence and Governance

---

# SECTION 15 — Capability Success Metrics

## Core Capability Success Metrics

### Graduate Profile Management

**Adoption Metric**
- Percentage of schools using Graduate Profile as alignment framework

**Outcome Metric**
- Percentage of students demonstrating progress across all 8 dimensions

**Automation Metric**
- Percentage of alignment checks automated (Target: 80%)

**Quality Metric**
- Accuracy of outcome mapping to Graduate Profile dimensions

---

### Curriculum Management

**Adoption Metric**
- Percentage of teachers accessing curriculum materials through platform

**Outcome Metric**
- Alignment percentage of lesson plans with curriculum (Target: 95%)

**Automation Metric**
- Percentage of CP → TP transformation automated (Target: 70%)

**Quality Metric**
- Curriculum update cycle time

---

### Learning Planning

**Adoption Metric**
- Percentage of teachers using platform for lesson planning

**Outcome Metric**
- Lesson plan quality score (based on rubric)

**Automation Metric**
- Percentage of lesson plan generation automated (Target: 80%)

**Quality Metric**
- Teacher time saved on lesson planning (Target: 90% reduction)

---

### Assessment Management

**Adoption Metric**
- Assessment completion rate

**Outcome Metric**
- Evidence quality score (based on rubric)

**Automation Metric**
- Percentage of evidence evaluation automated (Target: 80%)

**Quality Metric**
- Competency coverage percentage (Target: 100%)

---

### Reporting Management

**Adoption Metric**
- Report delivery rate to parents

**Outcome Metric**
- Parent understanding of child's progress (survey-based)

**Automation Metric**
- Percentage of narrative generation automated (Target: 90%)

**Quality Metric**
- Report generation time (Target: 90% reduction)

---

### AI Orchestration

**Adoption Metric**
- AI agent usage rate across platform

**Outcome Metric**
- User satisfaction with AI assistance (survey-based)

**Automation Metric**
- Percentage of workflows AI-assisted (Target: 90%)

**Quality Metric**
- AI recommendation acceptance rate (Target: 85%)

---

## Strategic Capability Success Metrics

### Competency Intelligence

**Adoption Metric**
- Competency Graph query rate

**Outcome Metric**
- Learning pathway optimization impact on learning outcomes

**Automation Metric**
- Percentage of pathway recommendations automated (Target: 90%)

**Quality Metric**
- Competency Graph accuracy and completeness

---

### Lifelong Learning Record

**Adoption Metric**
- Record portability across educational phases

**Outcome Metric**
- Career readiness assessment accuracy

**Automation Metric**
- Percentage of record aggregation automated (Target: 90%)

**Quality Metric**
- Record completeness across learning phases

---

# SECTION 16 — Capability to Domain Traceability

## Domain to Capability Mapping

This matrix ensures that every capability is owned by a domain and every domain has capabilities.

| Domain | Capabilities |
| ------ | ------------ |
| **Graduate Profile Domain** | Graduate Profile Definition, Outcome Mapping, Progress Tracking, Alignment Verification |
| **Curriculum Domain** | CP Management, TP Management, ATP Management, Curriculum Mapping, Curriculum Alignment, Curriculum Review |
| **Learning Planning Domain** | Lesson Planning, Resource Matching, Differentiation Planning, Activity Sequencing, Plan Sharing |
| **Learning Domain** | Session Management, Activity Execution, Resource Access, Student Engagement, Real-time Adaptation |
| **Assessment Domain** | Assessment Design, Assessment Delivery, Evidence Collection, Evidence Evaluation, Competency Validation, Assessment Moderation |
| **Reporting Domain** | Report Generation, Narrative Writing, Multi-language Support, Parent Communication, Trend Analysis |
| **Student Wellbeing Domain** | Growth Tracking, Wellbeing Monitoring, Intervention Planning, Strength Identification, Gap Analysis |
| **Teacher Professional Growth Domain** | Reflection Support, Feedback Collection, PD Resource Access, Growth Planning, Collaboration Support |
| **School Improvement Domain** | Quality Planning, Data Analysis, Benchmarking, Improvement Tracking, SPMI Support |
| **Parent Partnership Domain** | Communication Management, Engagement Tracking, Resource Sharing, Feedback Collection, Event Management |
| **Education Data Domain** | Data Aggregation, Descriptive Analytics, Predictive Analytics, Prescriptive Analytics, Visualization |
| **AI Orchestration Domain** | Agent Coordination, Workflow Orchestration, Human Validation, Performance Monitoring, Governance Enforcement |
| **Lifelong Learning Record Domain** | Record Aggregation, Competency Tracking, Portfolio Management, Credential Verification, Portability |
| **Career & Future Readiness Domain** | Career Exploration, Skill Development, Industry Connection, Career Guidance, Pathway Planning |
| **Quality Assurance Domain** | Quality Monitoring, Compliance Checking, Accreditation Support, Audit Support, Risk Management |

## Traceability Principles

### No Capability Without Domain
- Every capability must have a domain owner
- No capability exists outside domain boundaries
- Capability ownership follows domain ownership from 01

### No Domain Without Capability
- Every domain must have at least one capability
- Domains without capabilities indicate architectural gaps
- Domain capability completeness is verified

### One-to-Many Mapping
- One domain can have multiple capabilities
- One capability has exactly one domain owner
- Clear ownership enables effective governance

---

# SECTION 17 — Capability to AI Traceability

## Capability to AI Agent Mapping

This matrix ensures that every AI capability has an AI agent and every AI agent has capabilities.

| Capability | AI Agent |
| ---------- | -------- |
| **TP Management** | Curriculum Agent |
| **ATP Management** | Curriculum Agent |
| **Lesson Planning** | Lesson Planning Agent |
| **Resource Matching** | Lesson Planning Agent |
| **Activity Sequencing** | Lesson Planning Agent |
| **Assessment Design** | Assessment Agent |
| **Assessment Delivery** | Assessment Agent |
| **Evidence Collection** | Evidence Agent |
| **Evidence Evaluation** | Evidence Agent |
| **Competency Validation** | Assessment Agent |
| **Report Generation** | Parent Communication Agent |
| **Narrative Writing** | Parent Communication Agent |
| **Growth Tracking** | Student Growth Agent |
| **Wellbeing Monitoring** | Student Growth Agent |
| **Reflection Support** | Teacher Coach Agent |
| **Data Aggregation** | Analytics Agent |
| **Descriptive Analytics** | Analytics Agent |
| **Predictive Analytics** | Analytics Agent |
| **Agent Coordination** | Governance Agent |
| **Competency Graph Query** | Competency Intelligence Agent |
| **Pathway Analysis** | Competency Intelligence Agent |
| **Record Aggregation** | Lifelong Learning Agent |
| **Credential Verification** | Lifelong Learning Agent |

## Traceability Principles

### No AI Capability Without Agent
- Every AI-assisted capability must have an assigned AI agent
- No AI capability exists without agent ownership
- AI agent assignment follows capability requirements

### No AI Agent Without Capability
- Every AI agent must have at least one capability
- AI agents without capabilities indicate architectural gaps
- AI agent capability completeness is verified

### Human Validation Points
- Every AI-assisted capability has defined human validation points
- Human validation requirements are documented
- AI agent outputs require human approval before use

---

# SECTION 18 — Foundation Alignment Validation

## Alignment with 00A National Education Direction

### Indonesia Emas 2045 Vision
- All capabilities contribute to human capital development goals
- Lifelong Learning Record capability enables workforce planning
- Competency Intelligence capability enables skills gap analysis
- Career Readiness capability supports industry alignment

### Graduate Profile (8 Dimensions)
- All capabilities are mapped to Graduate Profile dimensions
- Outcome traceability ensures contribution to Profil Lulusan
- Assessment capabilities measure dimension development
- Reporting capabilities communicate dimension progress

### Outcome-Driven Education
- Every capability has clear outcome contribution
- Capability success metrics include outcome measurement
- Outcome traceability is maintained from capability to national goals

---

## Alignment with 00B Product Vision

### AI-First Architecture
- Every capability is evaluated for AI automation potential
- AI Capability Model defines automation targets
- AI Agent Capability Matrix defines agent responsibilities
- 90% AI Assistance target is operationalized at capability level

### 90% AI Assistance, 10% Human Governance
- AI automation targets average 90% across capabilities
- Human validation points are defined for all AI-assisted capabilities
- Human authority is maintained at critical decision points
- AI agents serve as co-pilots, not decision-makers

### Curriculum First, Pedagogy Before Technology
- Core capabilities prioritize curriculum and pedagogy
- Technology capabilities support, not replace, educational work
- Teacher capabilities maintain pedagogical authority
- AI capabilities enhance, not replace, human judgment

---

## Alignment with 00C Operating Principles

### Human Authority Principle
- Strategic capabilities require human decision-making
- Human approval requirements are defined for all AI-assisted capabilities
- Human governance is maintained at all levels
- AI agents are designed as co-pilots, not decision-makers

### Single Source of Truth
- Curriculum Management capability provides SSOT for curriculum
- Data ownership follows domain ownership from 01
- Capability ownership follows domain ownership
- No duplicate authority across capabilities

### Platform Strategy
- Platform capabilities orchestrate system operations
- AI Orchestration capability enables multi-agent coordination
- Competency Intelligence capability provides platform intelligence
- Lifelong Learning Record capability provides longitudinal data

---

## Alignment with 01 Education Domain Model

### Domain Ownership
- Every capability has a domain owner from 01
- Capability ownership follows domain ownership matrix from 01
- No capability exists outside domain boundaries
- Domain capability completeness is verified

### Domain Dependencies
- Capability dependency map follows domain dependency hierarchy from 01
- Implementation sequencing respects domain dependencies
- Data flow follows domain dependency patterns
- Critical dependencies are identified and monitored

### AI Agent Mapping
- AI Agent Capability Matrix follows AI Agent ↔ Domain Mapping from 01
- AI agent responsibilities align with domain responsibilities
- Human validation points follow Human Governance Layer from 01
- AI automation targets follow AI-Native Domain Matrix from 01

### Competency Architecture
- Competency Intelligence capability implements Competency Graph from 01
- Lifelong Learning Record capability implements Lifelong Learning Record from 01
- Learning Digital Twin capability implements Learning Digital Twin from 01
- National Human Capital Intelligence capability implements intelligence layer from 01

---

## Complete Alignment Verification

### Graduate Profile (8 Dimensions)
- ✅ All 8 dimensions are covered by capabilities
- ✅ Outcome traceability is maintained
- ✅ Assessment capabilities measure dimension development
- ✅ Reporting capabilities communicate dimension progress

### Deep Learning
- ✅ Learning Delivery capability supports deep learning pedagogy
- ✅ Assessment capability supports evidence-based assessment
- ✅ Teacher Development capability supports pedagogical growth

### Competency Graph
- ✅ Competency Intelligence capability implements Competency Graph
- ✅ AI agents leverage Competency Graph for recommendations
- ✅ Learning pathways are derived from Competency Graph

### Digital Twin
- ✅ Student Development capability implements Learning Digital Twin
- ✅ Growth tracking maintains twin data
- ✅ Personalization uses twin data

### Lifelong Learning Record
- ✅ Lifelong Learning Record capability implements longitudinal record
- ✅ Credential capability supports record verification
- ✅ Portability is ensured across phases

### Human Governance
- ✅ Human approval requirements are defined for all AI capabilities
- ✅ Human authority is maintained at critical decisions
- ✅ AI agents serve as co-pilots, not decision-makers

### AI First Architecture
- ✅ Every capability is evaluated for AI automation
- ✅ AI Capability Model defines automation targets
- ✅ AI Agent Capability Matrix defines agent responsibilities
- ✅ 90% AI Assistance target is operationalized

---

# SECTION 19 — Final Architecture Decision

## Official Reference Status

This Capability Model (02_CAPABILITY_MODEL.md) is established as the **official reference** for all subsequent architecture work:

### For Business Process Architecture (03)
- Business processes are designed to execute capabilities
- Process boundaries respect capability boundaries
- Process flows follow capability dependencies
- No process exists without a corresponding capability

### For Data Architecture (04)
- Data entities are derived from capability requirements
- Data ownership follows capability ownership
- Data flows follow capability dependencies
- No data entity exists without a corresponding capability

### For AI Architecture (05)
- AI agents are designed based on capability requirements
- AI agent responsibilities align with capability automation targets
- AI governance follows capability human validation points
- No AI agent exists without a corresponding capability

### For Application Architecture (06)
- Applications are built to deliver capabilities
- Module boundaries follow capability boundaries
- Application features implement capabilities
- No application feature exists without a corresponding capability

### For MVP Architecture
- MVP scope is defined by selecting Wave 1 capabilities
- MVP implementation follows capability prioritization
- MVP success is measured by capability success metrics
- No MVP feature exists without a corresponding capability

### For SDLC Architecture
- Development work is organized by capability
- Testing validates capability delivery
- Deployment follows capability dependencies
- No development work exists without a corresponding capability

## Architectural Compliance Requirement

**No process, module, AI agent, database entity, or feature may exist that cannot be traced to a capability defined in this document.**

### Compliance Checklist
- [ ] Process can be traced to capability
- [ ] Module can be traced to capability
- [ ] AI agent can be traced to capability
- [ ] Database entity can be traced to capability
- [ ] Feature can be traced to capability
- [ ] Capability is defined in this document
- [ ] Capability has domain owner
- [ ] Capability has human owner
- [ ] Capability has AI agent (if AI-assisted)
- [ ] Capability has success metrics

### Non-Compliance Consequence
Any design that cannot be traced to a capability defined in this document must be rejected or modified to achieve compliance.

## Strategic Mandate

The Capability Model v1.0 is established as:

**The authoritative capability reference for all Education Operating System (NUSA) development.**

This mandate ensures that:
- All development is aligned with educational outcomes
- All architecture is coherent and consistent
- All systems are built on stable foundations
- All evolution is governed and controlled
- All investments deliver measurable value

## Final Statement

**The Capability Model Phase is complete.**

The Capability Model (02_CAPABILITY_MODEL.md) provides a complete, coherent, and authoritative translation from Domain Architecture (01) to Capability Architecture.

With this capability model defined, the foundation is established for:
- 03_BUSINESS_PROCESS_ARCHITECTURE.md
- 04_DATA_ARCHITECTURE.md
- 05_AI_ARCHITECTURE.md
- 06_APPLICATION_ARCHITECTURE.md
- MVP Architecture
- SDLC Architecture

The architectural risk for subsequent work is significantly reduced with this capability model in place.

**Capability Model v1.0 is declared.**

---

# SECTION 20 — Capability → Module Boundary Mapping

## Purpose

This section translates capabilities into module boundaries that will be used in Application Architecture, Module Architecture, and Team Architecture. Module boundaries are derived from capability boundaries to ensure architectural coherence.

## Module Boundary Principles

### One Module = One Bounded Context
- Each module corresponds to one bounded context from Domain-Driven Design
- Module boundaries align with domain boundaries
- Modules maintain autonomy within their bounded context

### One Module = One Primary Capability Group
- Each module implements one primary capability group
- Modules may support multiple related capabilities within the group
- Module scope is defined by capability group boundaries

### Module Ownership Follows Domain Ownership
- Module ownership follows domain ownership from 01
- Primary domain owner is the module owner
- Module governance follows domain governance

## Capability to Module Boundary Mapping

| Capability Group | Module Boundary | Primary Domain |
| ---------------- | ---------------- | -------------- |
| **Graduate Profile Management** | Graduate Profile Module | Graduate Profile Domain |
| **Curriculum Management** | Curriculum Module | Curriculum Domain |
| **Learning Planning** | Planning Module | Learning Planning Domain |
| **Learning Delivery** | Learning Module | Learning Delivery Domain |
| **Assessment Management** | Assessment Module | Assessment Domain |
| **Reporting Management** | Reporting Module | Reporting Domain |
| **Student Development** | Student Growth Module | Student Wellbeing Domain |
| **Teacher Development** | Teacher Growth Module | Teacher Professional Growth Domain |
| **Parent Partnership** | Parent Module | Parent Partnership Domain |
| **School Improvement** | School Improvement Module | School Improvement Domain |
| **Education Analytics** | Analytics Module | Education Data Domain |
| **Competency Intelligence** | Competency Graph Module | Lifelong Learning Record Domain |
| **Lifelong Learning Record** | Learning Identity Module | Lifelong Learning Record Domain |
| **Credential & Achievement** | Credential Module | Lifelong Learning Record Domain |
| **AI Orchestration** | AI Orchestration Module | AI Orchestration Domain |
| **Governance & Compliance** | Governance Module | Quality Assurance Domain |

## Core Principles

### Capability ≠ Module
- Capabilities are what the platform can do
- Modules are how capabilities are implemented
- One capability may be implemented by multiple modules
- One module may implement multiple related capabilities

### Module Traceability
- Every module must be traceable to one or more capabilities
- No module exists without a corresponding capability
- Module boundaries respect capability boundaries
- Module dependencies follow capability dependencies

### Module Autonomy
- Modules are autonomous within their bounded context
- Modules own their data within their domain
- Modules communicate through well-defined interfaces
- Module evolution follows capability evolution

## Strategic Implications

### For Application Architecture
- Application boundaries follow module boundaries
- Applications are composed of modules
- Application features implement capabilities
- No application feature exists without capability traceability

### For Module Architecture
- Module boundaries follow capability boundaries
- Each module implements one capability group
- Module autonomy is ensured by bounded context
- Module communication follows module interfaces

### For Team Architecture
- Team boundaries follow module boundaries
- Teams are organized by module
- Team dependencies follow module dependencies
- Team ownership follows module ownership
- Each team owns one or more modules
- Team autonomy is ensured by module autonomy
- Team collaboration follows module interfaces

## Core Principle

**Module boundaries are derived from capability boundaries to ensure architectural coherence.**

This principle ensures that:
- Modules implement well-defined capabilities
- Module boundaries are stable and meaningful
- Module ownership is clear and effective
- Module evolution is controlled and coherent

---

# SECTION 21 — AI Automation Classification

## Purpose

This section explicitly defines the automation levels to achieve the target of 90% AI Automation and 10% Human Governance. Each capability is classified by its automation level to guide implementation and measure progress.

## Automation Levels

### Level 0 — Human Only
- No AI assistance
- All work is done by humans
- Full human control and responsibility
- Used for strategic decision-making and ethical judgment

### Level 1 — AI Assisted
- AI provides assistance and recommendations
- Humans perform the core work
- AI augments human capability
- Humans maintain full control

### Level 2 — AI Accelerated
- AI generates drafts and initial outputs
- Humans review and refine
- AI handles routine work
- Humans ensure quality and accuracy

### Level 3 — AI Automated
- AI executes the process autonomously
- Humans handle exceptions and edge cases
- AI operates within defined parameters
- Humans monitor and intervene when needed

### Level 4 — AI Autonomous
- AI executes the full process autonomously
- Humans audit and oversee
- AI makes decisions within governance framework
- Humans provide strategic direction

## AI Automation Classification Matrix

| Capability | Automation Level | Human Involvement |
| ---------- | ---------------- | ----------------- |

### Curriculum Capabilities

| Capability | Automation Level | Human Involvement |
| ---------- | ---------------- | ----------------- |
| CP Mapping | Level 2 (AI Accelerated) | Review and approve curriculum mappings |
| TP Generation | Level 2 (AI Accelerated) | Review and approve learning objectives |
| ATP Generation | Level 3 (AI Automated) | Monitor and adjust learning paths |
| Curriculum Alignment | Level 2 (AI Accelerated) | Review alignment reports |

### Planning Capabilities

| Capability | Automation Level | Human Involvement |
| ---------- | ---------------- | ----------------- |
| Lesson Planning | Level 2 (AI Accelerated) | Customize and approve lesson plans |
| Modul Ajar Generation | Level 2 (AI Accelerated) | Review and refine teaching materials |
| Resource Matching | Level 3 (AI Automated) | Override resource recommendations |
| Activity Sequencing | Level 2 (AI Accelerated) | Adjust activity sequences |

### Assessment Capabilities

| Capability | Automation Level | Human Involvement |
| ---------- | ---------------- | ----------------- |
| Assessment Design | Level 2 (AI Accelerated) | Review and approve assessments |
| Evidence Analysis | Level 3 (AI Automated) | Override evidence evaluations |
| Competency Evaluation | Level 2 (AI Accelerated) | Validate competency mastery |
| Assessment Moderation | Level 1 (AI Assisted) | Moderate assessment quality |

### Reporting Capabilities

| Capability | Automation Level | Human Involvement |
| ---------- | ---------------- | ----------------- |
| Report Generation | Level 3 (AI Automated) | Customize report content |
| Parent Summary | Level 2 (AI Accelerated) | Review and approve communications |
| Narrative Writing | Level 2 (AI Accelerated) | Refine narrative descriptions |
| Trend Analysis | Level 3 (AI Automated) | Interpret trend insights |

### Analytics Capabilities

| Capability | Automation Level | Human Involvement |
| ---------- | ---------------- | ----------------- |
| Risk Detection | Level 3 (AI Automated) | Review risk alerts |
| Intervention Recommendation | Level 2 (AI Accelerated) | Approve intervention plans |
| Predictive Analytics | Level 2 (AI Accelerated) | Interpret predictions |
| Prescriptive Analytics | Level 2 (AI Accelerated) | Review recommendations |

### Competency Intelligence Capabilities

| Capability | Automation Level | Human Involvement |
| ---------- | ---------------- | ----------------- |
| Competency Graph Update | Level 3 (AI Automated) | Validate graph updates |
| Digital Twin Update | Level 3 (AI Automated) | Review twin accuracy |
| Pathway Analysis | Level 3 (AI Automated) | Override pathway recommendations |
| Gap Analysis | Level 2 (AI Accelerated) | Review gap analysis |

### Lifelong Learning Capabilities

| Capability | Automation Level | Human Involvement |
| ---------- | ---------------- | ----------------- |
| Learning Passport Update | Level 3 (AI Automated) | Verify learning records |
| Credential Verification | Level 3 (AI Automated) | Handle verification disputes |
| Career Readiness Assessment | Level 2 (AI Accelerated) | Review career recommendations |
| Record Aggregation | Level 3 (AI Automated) | Monitor aggregation quality |

## Automation Targets

### MVP Target (30 Days)
- **70–80% Automation** across Wave 1 capabilities
- Focus on reducing teacher burden in core workflows
- Human validation maintained at critical decision points

### Phase 2 Target
- **80–85% Automation** across Wave 1 and Wave 2 capabilities
- Expand automation to analytics and intelligence capabilities
- Maintain human governance for strategic decisions

### Long Term Target (2045)
- **90% Automation** across all capabilities
- AI operates autonomously within governance framework
- Human focus shifts to strategic direction and exception handling

## Core Principle

**AI automation levels guide implementation and measure progress toward 90% AI assistance target.**

This principle ensures that:
- Automation targets are clear and measurable
- Human involvement is appropriate to capability nature
- AI capabilities are leveraged effectively
- Human governance is maintained

---

# SECTION 22 — Capability Critical Path

## Purpose

This section identifies the critical capabilities that must be implemented first for NUSA to deliver real value. The critical path defines the implementation sequence that maximizes early value while respecting dependencies.

## Critical Path 1 — Graduate Outcome Flow

The core education flow that delivers graduate outcomes:

```
Graduate Profile Management
    ↓ defines
Curriculum Management
    ↓ enables
Learning Planning
    ↓ enables
Learning Delivery
    ↓ generates
Assessment Management
    ↓ produces
Reporting Management
    ↓ informs
Student Development
    ↓ validates
Graduate Profile Validation
```

**Critical Capabilities**: Graduate Profile, Curriculum, Planning, Learning, Assessment, Reporting, Student Development

**Value Delivered**: Complete education workflow from curriculum definition to outcome validation

**Implementation Priority**: Wave 1 (MVP)

---

## Critical Path 2 — Competency Intelligence Flow

The intelligence flow that enables personalization and lifelong learning:

```
Learning Delivery
    ↓ generates
Assessment Evidence
    ↓ feeds
Competency Evaluation
    ↓ updates
Competency Graph
    → feeds
Digital Twin
    → aggregates into
Lifelong Learning Record
```

**Critical Capabilities**: Learning Delivery, Assessment, Competency Intelligence, Lifelong Learning

**Value Delivered**: Personalized learning and longitudinal tracking

**Implementation Priority**: Wave 2

---

## Critical Path 3 — AI Automation Flow

The automation flow that delivers AI assistance:

```
Curriculum Management
    ↓ enables
Learning Planning
    ↓ enables
Assessment Management
    ↓ enables
Reporting Management
    ↓ provides data for
Education Analytics
    ↓ enables
AI Recommendations
```

**Critical Capabilities**: Curriculum, Planning, Assessment, Reporting, Analytics, AI Orchestration

**Value Delivered**: AI assistance that reduces teacher burden

**Implementation Priority**: Wave 1 (MVP) for core, Wave 2 for advanced analytics

---

## Critical Path Principles

### Dependency Respect
- Critical paths respect capability dependencies
- Upstream capabilities must be implemented before downstream
- Implementation sequence follows dependency hierarchy

### Value Maximization
- Critical paths prioritize capabilities that deliver highest value
- Early implementation focuses on teacher burden reduction
- Subsequent implementation expands to intelligence and analytics

### Risk Mitigation
- Critical paths identify high-risk dependencies
- Redundancy planned for critical capabilities
- Failure recovery prioritizes critical path capabilities

## Strategic Implications

### For MVP Prioritization
- MVP scope is defined by Critical Path 1 capabilities
- MVP delivers complete core education workflow
- MVP provides immediate value to teachers

### For Module Implementation Order
- Modules are implemented in critical path sequence
- Upstream modules are implemented before downstream
- Module dependencies follow critical path dependencies

### For Technical Roadmap
- Technical roadmap follows critical path sequence
- Infrastructure supports critical path capabilities first
- Technical debt is managed to support critical path evolution

## Core Principle

**Critical path defines the implementation sequence that maximizes early value while respecting dependencies.**

This principle ensures that:
- Early implementation delivers maximum value
- Implementation sequence respects dependencies
- Technical roadmap supports business priorities
- Risk is managed through dependency awareness

---

# SECTION 23 — Capability Outcome Chain

## Purpose

This section ensures that every capability produces measurable educational impact. The outcome chain traces how capabilities contribute from immediate outcomes to strategic outcomes.

## Outcome Chain Model

The complete outcome chain from learning activity to national outcome:

```
Learning Activity
    ↓ produces
Assessment Evidence
    ↓ validates
Competency Validation
    ↓ updates
Competency Graph
    → feeds
Digital Twin
    → aggregates into
Graduate Profile Progress
    → achieves
Graduate Profile Achievement
    → contributes to
Human Capital Outcome
    → supports
National Education Outcome
    → realizes
Indonesia Emas 2045
```

## Capability Outcome Matrix

| Capability | Immediate Outcome | Intermediate Outcome | Strategic Outcome |
| ---------- | ----------------- | -------------------- | ----------------- |

### Assessment Management

| Capability | Immediate Outcome | Intermediate Outcome | Strategic Outcome |
| ---------- | ----------------- | -------------------- | ----------------- |
| **Assessment Management** | Evidence of learning | Competency validation | Graduate Profile achievement |

**Explanation**: Assessment collects evidence (immediate), validates competency (intermediate), which contributes to graduate profile achievement (strategic).

---

### Learning Planning

| Capability | Immediate Outcome | Intermediate Outcome | Strategic Outcome |
| ---------- | ----------------- | -------------------- | ----------------- |
| **Learning Planning** | Learning plan | Learning quality | Student growth |

**Explanation**: Planning creates learning plans (immediate), which improves learning quality (intermediate), leading to student growth (strategic).

---

### Competency Intelligence

| Capability | Immediate Outcome | Intermediate Outcome | Strategic Outcome |
| ---------- | ----------------- | -------------------- | ----------------- |
| **Competency Intelligence** | Competency Graph | Digital Twin | Lifelong Learning Record |

**Explanation**: Intelligence maintains Competency Graph (immediate), which feeds Digital Twin (intermediate), enabling Lifelong Learning Record (strategic).

---

### Lifelong Learning Record

| Capability | Immediate Outcome | Intermediate Outcome | Strategic Outcome |
| ---------- | ----------------- | -------------------- | ----------------- |
| **Lifelong Learning Record** | Verified learning history | Career readiness | Human capital development |

**Explanation**: Record maintains verified history (immediate), which enables career readiness (intermediate), contributing to human capital development (strategic).

---

### Curriculum Management

| Capability | Immediate Outcome | Intermediate Outcome | Strategic Outcome |
| ---------- | ----------------- | -------------------- | ----------------- |
| **Curriculum Management** | Aligned curriculum | Learning alignment | Graduate Profile alignment |

**Explanation**: Curriculum provides aligned materials (immediate), ensuring learning alignment (intermediate), which aligns with Graduate Profile (strategic).

---

### AI Orchestration

| Capability | Immediate Outcome | Intermediate Outcome | Strategic Outcome |
| ---------- | ----------------- | -------------------- | ----------------- |
| **AI Orchestration** | AI assistance | Teacher efficiency | Education quality improvement |

**Explanation**: Orchestration provides AI assistance (immediate), improving teacher efficiency (intermediate), leading to education quality improvement (strategic).

---

## Outcome Chain Principles

### No Capability Stands Alone
- Every capability produces outcomes
- Every outcome contributes to higher-level outcomes
- Capabilities are connected through outcome chains

### Measurable Outcomes
- Every capability has measurable immediate outcomes
- Intermediate outcomes are defined and tracked
- Strategic outcomes are aligned with national goals

### Traceability to National Goals
- Every outcome chain traces to Profil Lulusan 8 Dimensi
- Every outcome chain traces to Indonesia Emas 2045
- Every outcome chain traces to Human Capital Development

### Outcome Validation
- Outcomes are validated through evidence
- Outcome achievement is measured
- Outcome impact is assessed

## Strategic Implications

### For Capability Design
- Capabilities are designed with clear outcome chains
- Capability success is measured by outcome achievement
- Capability evolution is guided by outcome impact

### For Investment Prioritization
- Capabilities with strong outcome chains are prioritized
- Investment follows outcome impact
- ROI is measured by outcome achievement

### For Stakeholder Communication
- Outcome chains communicate capability value
- Stakeholders understand capability impact
- Alignment with national goals is demonstrated

## Core Principle

**Every capability produces measurable educational impact through outcome chains.**

This principle ensures that:
- Capabilities deliver real educational value
- Outcomes are measurable and traceable
- Investment is aligned with impact
- Stakeholders understand capability value

---

# SECTION 24 — Capability Data Ownership & Data Flow Mapping

## Objective

This section explains the relationship between capabilities and the data they produce and consume. This mapping serves as the bridge to Data Architecture (04_DATA_ARCHITECTURE.md).

## Purpose

- Serve as the bridge to 04_DATA_ARCHITECTURE.md
- Ensure every data entity has a clear origin
- Ensure no data exists without a corresponding capability
- Maintain Single Source of Truth (SSOT) from Domain Architecture

## Data Flow Principles

### Capability Creates Data
- Capabilities produce data as part of their operation
- Data production is a capability output
- Data creation is traceable to the producing capability

### Capability Consumes Data
- Capabilities consume data as input for their operation
- Data consumption is a capability input
- Data usage is traceable to the consuming capability

### Capability Never Owns Data Directly
- Capabilities interact with data but do not own data
- Data ownership remains with domains as defined in 01
- Capability data interactions follow domain data ownership

### Data Ownership Remains in Domain Architecture
- Data ownership is defined in Domain Architecture (01)
- Capability data interactions respect domain ownership
- No capability can claim data ownership outside its domain

### Capability Only Interacts with Data
- Capabilities create, read, update, and delete data
- Capability data interactions are governed by domain ownership
- Data authority remains with the owning domain

## Capability Data Flow Matrix

| Capability | Produces Data | Consumes Data |
| ---------- | ------------- | ------------- |

### Graduate Profile Management

| Capability | Produces Data | Consumes Data |
| ---------- | ------------- | ------------- |
| **Graduate Profile Management** | Graduate Profile, Graduate Outcome Definition | National Education Standards |

**Data Flow**: National Education Standards → Graduate Profile Definition → Graduate Profile

---

### Curriculum Management

| Capability | Produces Data | Consumes Data |
| ---------- | ------------- | ------------- |
| **Curriculum Management** | CP, TP, ATP, Curriculum Mapping | Graduate Profile |

**Data Flow**: Graduate Profile → CP Definition → TP Definition → ATP Generation → Curriculum Mapping

---

### Learning Planning

| Capability | Produces Data | Consumes Data |
| ---------- | ------------- | ------------- |
| **Learning Planning** | Learning Plan, Modul Ajar, Learning Objectives | CP, TP, ATP |

**Data Flow**: CP, TP, ATP → Learning Plan Generation → Modul Ajar Creation → Learning Objectives Definition

---

### Learning Delivery

| Capability | Produces Data | Consumes Data |
| ---------- | ------------- | ------------- |
| **Learning Delivery** | Learning Activities, Learning Sessions, Learning Participation | Learning Plan |

**Data Flow**: Learning Plan → Learning Session Execution → Learning Activity Delivery → Learning Participation Recording

---

### Assessment Management

| Capability | Produces Data | Consumes Data |
| ---------- | ------------- | ------------- |
| **Assessment Management** | Assessment, Assessment Evidence, Competency Evaluation | Student, Learning Activity, TP, ATP |

**Data Flow**: Student, Learning Activity, TP, ATP → Assessment Design → Assessment Delivery → Evidence Collection → Competency Evaluation

---

### Reporting Management

| Capability | Produces Data | Consumes Data |
| ---------- | ------------- | ------------- |
| **Reporting Management** | Progress Report, Narrative Report, Parent Summary | Assessment Evidence, Competency Evaluation |

**Data Flow**: Assessment Evidence, Competency Evaluation → Progress Analysis → Narrative Generation → Parent Summary Creation

---

### Student Development

| Capability | Produces Data | Consumes Data |
| ---------- | ------------- | ------------- |
| **Student Development** | Growth Plan, Intervention Plan | Student Progress, Assessment Results |

**Data Flow**: Student Progress, Assessment Results → Growth Analysis → Intervention Planning → Growth Plan Creation

---

### Teacher Development

| Capability | Produces Data | Consumes Data |
| ---------- | ------------- | ------------- |
| **Teacher Development** | Reflection Summary, Growth Plan | Teacher Performance, Student Outcomes |

**Data Flow**: Teacher Performance, Student Outcomes → Reflection Support → Growth Analysis → Growth Plan Creation

---

### School Improvement

| Capability | Produces Data | Consumes Data |
| ---------- | ------------- | ------------- |
| **School Improvement** | Improvement Plan, Quality Metrics | School Performance Data, Benchmark Data |

**Data Flow**: School Performance Data, Benchmark Data → Quality Analysis → Improvement Planning → Improvement Plan Creation

---

### Education Analytics

| Capability | Produces Data | Consumes Data |
| ---------- | ------------- | ------------- |
| **Education Analytics** | Insights, Risk Indicators, Recommendations | Learning Data, Assessment Data, Development Data |

**Data Flow**: Learning Data, Assessment Data, Development Data → Data Aggregation → Analytics Processing → Insight Generation

---

### Competency Intelligence

| Capability | Produces Data | Consumes Data |
| ---------- | ------------- | ------------- |
| **Competency Intelligence** | Competency Graph, Digital Twin Update | Competency Evaluation, Assessment Evidence |

**Data Flow**: Competency Evaluation, Assessment Evidence → Competency Graph Update → Digital Twin Update → Pathway Analysis

---

### Lifelong Learning Record

| Capability | Produces Data | Consumes Data |
| ---------- | ------------- | ------------- |
| **Lifelong Learning Record** | Learning Passport, Lifelong Learning Record | Competency Graph, Credentials, Achievements |

**Data Flow**: Competency Graph, Credentials, Achievements → Record Aggregation → Learning Passport Update → Lifelong Learning Record Maintenance

---

### Credential & Achievement

| Capability | Produces Data | Consumes Data |
| ---------- | ------------- | ------------- |
| **Credential & Achievement** | Credential, Achievement Record | Competency Evaluation, Achievement Evidence |

**Data Flow**: Competency Evaluation, Achievement Evidence → Credential Issuance → Achievement Recording → Credential Verification

---

### AI Orchestration

| Capability | Produces Data | Consumes Data |
| ---------- | ------------- | ------------- |
| **AI Orchestration** | AI Recommendation, Agent Log | User Request, Context Data, Capability Data |

**Data Flow**: User Request, Context Data, Capability Data → Agent Coordination → Recommendation Generation → Log Recording

---

## Core Principle

**Data flow must always be traceable to the capability that produces the data.**

This principle ensures that:
- Every data entity has a clear origin
- Data creation is traceable to capabilities
- Data consumption is traceable to capabilities
- Data ownership remains with domains
- Single Source of Truth is maintained

## Strategic Implications

### For Data Architecture (04)
- Data entities are derived from capability data flow
- Data models respect capability data interactions
- Data ownership follows domain ownership
- No data entity exists without capability traceability

### For Application Architecture (06)
- Applications implement capabilities
- Application data access follows capability data flow
- Application data interactions respect domain ownership
- No application data access exists without capability traceability

### For AI Architecture (05)
- AI agents consume and produce data through capabilities
- AI data interactions follow capability data flow
- AI data recommendations respect domain ownership
- No AI data interaction exists without capability traceability

---

# SECTION 25 — Capability Event Architecture

## Objective

This section defines the primary events generated by each capability. Events serve as the foundation for Process Architecture, AI Architecture, Event-Driven Architecture, and Module Architecture.

## Purpose

- Define the event model for the Education Operating System
- Enable event-driven workflows
- Enable AI agent triggering
- Enable Competency Graph updates
- Enable Digital Twin updates
- Enable Lifelong Learning Record updates

## Event Architecture Principles

### Capabilities Generate Events
- Every capability generates events as part of its operation
- Events represent significant state changes
- Events are the primary communication mechanism between capabilities

### Events Trigger Workflows
- Events trigger business process workflows
- Workflows are event-driven
- Process state is maintained through events

### Events Trigger AI Agents
- Events trigger AI agent actions
- AI agents respond to specific events
- AI agent coordination is event-driven

### Events Update Competency Graph
- Competency achievement events update the Competency Graph
- Learning events contribute to competency tracking
- Competency Graph is event-sourced

### Events Update Digital Twin
- Learning events update the Digital Twin
- Assessment events update competency mastery
- Digital Twin is maintained through events

### Events Update Lifelong Learning Record
- Credential events update the Lifelong Learning Record
- Achievement events contribute to the record
- Lifelong Learning Record is event-sourced

## Capability Event Matrix

| Capability | Primary Events |
| ---------- | ------------- |
| **Graduate Profile Management** | GraduateProfileDefined |
| **Curriculum Management** | CurriculumPublished |
| **Learning Planning** | LessonPlanCreated |
| **Learning Delivery** | LearningSessionCompleted |
| **Assessment Management** | AssessmentSubmitted, EvidenceValidated |
| **Reporting Management** | ReportGenerated |
| **Student Development** | InterventionTriggered |
| **Teacher Development** | TeacherGrowthUpdated |
| **School Improvement** | ImprovementPlanCreated |
| **Education Analytics** | RiskDetected |
| **Competency Intelligence** | CompetencyAchieved, CompetencyGraphUpdated |
| **Lifelong Learning Record** | LearningRecordUpdated |
| **Credential & Achievement** | CredentialIssued |
| **AI Orchestration** | AIRecommendationGenerated |

## Critical Education Events

The following are the most critical events in the platform:

### LearningCompleted
**Trigger**: Learning session completion
**Purpose**: Marks completion of a learning activity
**Consumers**: Assessment, Competency Intelligence, Digital Twin

### AssessmentSubmitted
**Trigger**: Student submits assessment
**Purpose**: Initiates assessment evaluation
**Consumers**: Assessment Management, Competency Intelligence

### EvidenceValidated
**Trigger**: Evidence is validated against competency criteria
**Purpose**: Confirms evidence quality and relevance
**Consumers**: Competency Intelligence, Digital Twin

### CompetencyAchieved
**Trigger**: Student demonstrates competency mastery
**Purpose**: Records competency achievement
**Consumers**: Competency Graph, Digital Twin, Lifelong Learning Record

### CompetencyGraphUpdated
**Trigger**: Competency Graph is updated with new data
**Purpose**: Signals graph state change
**Consumers**: AI Agents, Analytics, Personalization

### DigitalTwinUpdated
**Trigger**: Digital Twin is updated with new learning data
**Purpose**: Signals twin state change
**Consumers**: AI Agents, Analytics, Personalization

### LearningRecordUpdated
**Trigger**: Lifelong Learning Record is updated
**Purpose**: Signals record state change
**Consumers**: Career Intelligence, Workforce Planning

### CredentialIssued
**Trigger**: Credential is issued to individual
**Purpose**: Records credential achievement
**Consumers**: Lifelong Learning Record, Career Intelligence

### InterventionTriggered
**Trigger**: Intervention is triggered for student
**Purpose**: Signals need for student support
**Consumers**: Student Development, Teacher Development, Parents

### GraduateProfileProgressUpdated
**Trigger**: Progress toward Graduate Profile is updated
**Purpose**: Tracks progress toward national outcomes
**Consumers**: Analytics, Reporting, National Intelligence

## Event Flow Examples

### Learning Flow
```
LearningSessionCompleted
    ↓ triggers
AssessmentSubmitted
    ↓ triggers
EvidenceValidated
    ↓ triggers
CompetencyAchieved
    ↓ triggers
CompetencyGraphUpdated
    ↓ triggers
DigitalTwinUpdated
    ↓ triggers
GraduateProfileProgressUpdated
```

### Credential Flow
```
CompetencyAchieved
    ↓ triggers
CredentialIssued
    ↓ triggers
LearningRecordUpdated
    ↓ triggers
DigitalTwinUpdated
```

### Intervention Flow
```
RiskDetected
    ↓ triggers
InterventionTriggered
    ↓ triggers
GrowthPlanCreated
    ↓ triggers
TeacherGrowthUpdated
```

## Core Principle

**All education processes in NUSA must be viewable as a series of historically traceable events.**

This principle ensures that:
- Every significant state change is captured as an event
- Event history provides complete audit trail
- Events enable event-driven architecture
- Events enable real-time intelligence
- Events enable replay and analysis

## Strategic Implications

### For Process Architecture (03)
- Business processes are event-driven
- Process workflows are triggered by events
- Process state is maintained through events

### For AI Architecture (05)
- AI agents are event-driven
- AI agents respond to specific events
- AI agent coordination is event-based

### For Event-Driven Architecture
- System architecture is event-driven
- Modules communicate through events
- Event sourcing is used for state management

### For Module Architecture (06)
- Modules are event-driven
- Module boundaries align with event boundaries
- Module communication is event-based

---

# SECTION 26 — MVP Capability Lock (30-Day MVP)

## Objective

This section officially establishes the capability scope for the 30-Day MVP. The purpose is to prevent scope creep and provide a clear engineering roadmap foundation.

## Purpose

- Officially establish MVP capability scope
- Prevent scope creep during MVP development
- Provide clear foundation for engineering roadmap
- Define MVP success outcomes

## MVP Wave 1 (Mandatory)

The following capabilities are mandatory and must be completed in the first MVP:

### Graduate Profile Management
- Graduate Profile Definition
- Outcome Mapping
- Progress Tracking

### Curriculum Management
- CP Management
- TP Management
- ATP Management
- Curriculum Alignment

### Learning Planning
- Lesson Planning
- Modul Ajar Generation
- Resource Matching

### Assessment Management
- Assessment Design
- Evidence Collection
- Evidence Evaluation
- Competency Validation

### Reporting Management
- Report Generation
- Narrative Writing
- Parent Communication

### AI Orchestration
- Agent Coordination
- Human Validation
- Workflow Orchestration

## MVP Success Outcome

The MVP must be capable of:

- **Generating ATP automatically** from CP and TP
- **Generating Modul Ajar automatically** from ATP
- **Generating Assessment automatically** aligned with learning objectives
- **Generating Narrative Reports automatically** from assessment evidence
- **Connecting all activities to Graduate Profile** (8 Dimensions)
- **Running AI-assisted workflow end-to-end** from curriculum to report

## Explicitly Out of Scope

The following capabilities are explicitly OUT OF SCOPE for the 30-Day MVP:

### Teacher Development
- Teacher Reflection
- Professional Growth
- PD Resource Access

### School Improvement
- Quality Planning
- Improvement Monitoring
- SPMI Support

### Parent Partnership
- Communication Management
- Engagement Tracking
- Resource Sharing

### Credential & Achievement
- Credential Issuance
- Credential Verification
- Badge Management

### Lifelong Learning Record
- Record Aggregation
- Competency Tracking
- Portfolio Management

### National Education Analytics
- Predictive Analytics
- Prescriptive Analytics
- National Intelligence

### Workforce Intelligence
- Career Readiness
- Industry Connection
- Workforce Planning

### Human Capital Intelligence
- National Human Capital Analysis
- Workforce Readiness
- Policy Intelligence

## Wave 2 (Post-MVP)

Capabilities that support growth and collaboration:

- Student Development (Growth Tracking, Wellbeing Monitoring)
- Teacher Development (Reflection Support, PD Resource Access)
- Parent Partnership (Communication Management)
- Education Analytics (Descriptive Analytics)

## Wave 3 (Post-Wave 2)

Competency intelligence and digital twin:

- Competency Intelligence (Competency Graph Query, Pathway Analysis)
- Student Development (Intervention Planning)
- School Improvement (Quality Planning)

## Wave 4 (Post-Wave 3)

Lifelong learning and national intelligence:

- Lifelong Learning Record (Record Aggregation, Competency Tracking)
- Credential & Achievement (Credential Issuance, Credential Verification)
- Education Analytics (Predictive Analytics, Prescriptive Analytics)
- Career Readiness (Career Exploration, Pathway Planning)

## Core Principle

**MVP is not built to cover all domains. MVP is built to prove the AI-Native Curriculum-to-Report Workflow.**

This principle ensures that:
- MVP scope is focused and achievable
- MVP delivers clear, measurable value
- MVP proves the core value proposition
- MVP provides foundation for subsequent waves
- Scope creep is prevented

## Strategic Implications

### For MVP Architecture
- MVP architecture implements Wave 1 capabilities only
- MVP modules are limited to Wave 1 module boundaries
- MVP data models are limited to Wave 1 data entities
- MVP AI agents are limited to Wave 1 agents

### For Engineering Roadmap
- Wave 1 capabilities are implemented in MVP (30 days)
- Wave 2 capabilities are implemented in Phase 2 (60 days)
- Wave 3 capabilities are implemented in Phase 3 (90 days)
- Wave 4 capabilities are implemented in Phase 4 (120 days)

### For Resource Allocation
- MVP resources focus on Wave 1 capabilities
- Subsequent waves allocate resources to additional capabilities
- Resource planning follows wave prioritization

---

# SECTION 27 — Capability Architecture Final Statement

## Declaration

This Capability Model (02_CAPABILITY_MODEL.md) is the official translation from:

- **00A_NATIONAL_EDUCATION_DIRECTION_2045.md** (National Education Direction)
- **00B_PRODUCT_VISION.md** (Product Vision)
- **00C_EDUCATION_OPERATING_SYSTEM_PRINCIPLES.md** (Operating Principles)
- **01_EDUCATION_DOMAIN_MODEL.md** (Domain Architecture)

## Complete Capability Architecture Coverage

This document now provides comprehensive coverage of:

### Capability Structure
- 16 Level 1 capabilities
- Comprehensive Level 2 decomposition
- Critical Level 3 operational capabilities
- Capability classification (Core, Strategic, Supporting, Platform)

### Capability Ownership
- Capability ownership matrix with domain owners
- Human owner assignments
- AI agent assignments
- Clear governance structure

### Capability Module Mapping
- 16 module boundaries mapped to capability groups
- Module boundary principles (DDD-aligned)
- Module ownership follows domain ownership
- Strategic implications for architecture

### Capability Data Mapping
- Capability data flow matrix
- Data production and consumption by capability
- Data ownership remains with domains
- Bridge to Data Architecture (04)

### Capability Event Mapping
- Capability event matrix
- Critical education events defined
- Event flow examples
- Foundation for event-driven architecture

### AI Capability Mapping
- AI automation classification (5 levels)
- AI automation targets (90% AI, 10% Human)
- AI Agent Capability Matrix (11 official agents)
- Human validation points defined

### Capability Dependencies
- Capability dependency map
- Critical paths defined (3 paths)
- Implementation sequencing
- Risk mitigation strategies

### Capability Outcomes
- Capability outcome chains
- Outcome traceability to national goals
- Success metrics defined
- Strategic impact measurement

### MVP Capability Scope
- MVP Wave 1 capabilities locked
- Explicitly out-of-scope capabilities
- Wave 2, 3, 4 planning
- MVP success outcomes defined

## Official Reference Status

This Capability Model serves as the **sole official reference** for all subsequent architecture work:

### For Business Process Architecture (03)
- Business processes are designed to execute capabilities
- Process boundaries respect capability boundaries
- Process flows follow capability dependencies
- No process exists without a corresponding capability

### For Data Architecture (04)
- Data entities are derived from capability data flow
- Data ownership follows domain ownership
- Data flows follow capability dependencies
- No data entity exists without a corresponding capability

### For AI Architecture (05)
- AI agents are designed based on capability requirements
- AI agent responsibilities align with capability automation targets
- AI governance follows capability human validation points
- No AI agent exists without a corresponding capability

### For Application Architecture (06)
- Applications are built to deliver capabilities
- Module boundaries follow capability module mapping
- Application features implement capabilities
- No application feature exists without a corresponding capability

### For MVP Architecture
- MVP scope is defined by MVP Wave 1 capabilities
- MVP implementation follows capability prioritization
- MVP success is measured by capability success metrics
- No MVP feature exists without a corresponding capability

### For SDLC Architecture
- Development work is organized by capability
- Testing validates capability delivery
- Deployment follows capability dependencies
- No development work exists without a corresponding capability

## Architectural Compliance Requirement

**No process, workflow, module, event, AI agent, data entity, or MVP feature may exist that cannot be traced to a capability defined in this document.**

### Compliance Checklist
- [ ] Process can be traced to capability
- [ ] Workflow can be traced to capability
- [ ] Module can be traced to capability
- [ ] Event can be traced to capability
- [ ] AI agent can be traced to capability
- [ ] Data entity can be traced to capability
- [ ] MVP feature can be traced to capability
- [ ] Capability is defined in this document
- [ ] Capability has domain owner
- [ ] Capability has human owner
- [ ] Capability has AI agent (if AI-assisted)
- [ ] Capability has success metrics
- [ ] Capability has outcome chain
- [ ] Capability has data flow mapping
- [ ] Capability has event mapping

### Non-Compliance Consequence
Any design that cannot be traced to a capability defined in this document must be rejected or modified to achieve compliance.

## Strategic Mandate

The Capability Architecture v1.0 is established as:

**The authoritative capability reference for all Education Operating System (NUSA) development.**

This mandate ensures that:
- All development is aligned with educational outcomes
- All architecture is coherent and consistent
- All systems are built on stable foundations
- All evolution is governed and controlled
- All investments deliver measurable value
- All traceability is maintained from feature to national outcome

## Final Statement

**The Capability Architecture Phase is complete.**

The Capability Model (02_CAPABILITY_MODEL.md) provides a complete, coherent, and authoritative translation from Domain Architecture (01) to Capability Architecture, including:

- Capability Structure and Ownership
- Capability Module Mapping
- Capability Data Mapping
- Capability Event Mapping
- AI Capability Mapping
- Capability Dependencies and Critical Paths
- Capability Outcome Chains
- MVP Capability Scope

With this capability architecture established, the foundation is established for:
- 03_BUSINESS_PROCESS_ARCHITECTURE.md
- 04_DATA_ARCHITECTURE.md
- 05_AI_ARCHITECTURE.md
- 06_APPLICATION_ARCHITECTURE.md
- 07_MVP_ARCHITECTURE.md
- 08_SDLC_ARCHITECTURE.md

The architectural risk for subsequent work is significantly reduced with this comprehensive capability architecture in place.

---

# SECTION 28 — Document Status

**Version**: 3.0
**Status**: FOUNDATION DOCUMENT
**Alignment**: 100% aligned with 00A_NATIONAL_EDUCATION_DIRECTION_2045.md, 00B_PRODUCT_VISION.md, 00C_EDUCATION_OPERATING_SYSTEM_PRINCIPLES.md, and 01_EDUCATION_DOMAIN_MODEL.md
**Purpose**: Capability model for Education Operating System Indonesia 2045
**Governance**: Changes to this document require review and approval from Chief Education Architect

This document serves as the official reference for:
- Business Capability Architecture
- AI Capability Architecture
- Human Capability Architecture
- MVP Scoping
- Module Identification
- Application Architecture
- AI Architecture
- Process Architecture

## Version History

- **Version 3.0** (June 2026): Added Capability Data Ownership & Data Flow Mapping (comprehensive data flow matrix for all 16 capabilities with data production and consumption mapping, data ownership principles, and bridge to Data Architecture), Capability Event Architecture (capability event matrix defining primary events for all capabilities, 10 critical education events with event flow examples, and foundation for event-driven architecture), and MVP Capability Lock (official MVP Wave 1 scope with 6 mandatory capabilities, explicitly out-of-scope capabilities, wave 2-4 planning, and MVP success outcomes). Updated Capability Architecture Final Statement to include comprehensive coverage of capability structure, ownership, service mapping, data mapping, event mapping, AI capability mapping, dependencies, outcomes, and MVP scope. Compliance checklist expanded to include process, workflow, service, event, AI agent, data entity, and MVP feature traceability. Document now serves as Capability Architecture v1.0 with complete coverage, ready as official reference for 03_BUSINESS_PROCESS_ARCHITECTURE.md, 04_DATA_ARCHITECTURE.md, 05_AI_ARCHITECTURE.md, 06_APPLICATION_ARCHITECTURE.md, 07_MVP_ARCHITECTURE.md, and 08_SDLC_ARCHITECTURE.md with significantly reduced architectural risk.
- **Version 2.0** (June 2026): Added Capability → Service Boundary Mapping (16 service boundaries mapped to capability groups with DDD principles), AI Automation Classification (5 automation levels from Human Only to AI Autonomous with detailed classification matrix for all capabilities), Capability Critical Path (3 critical paths defining implementation sequence for graduate outcome flow, competency intelligence flow, and AI automation flow), Capability Outcome Chain (complete outcome chain model from learning activity to Indonesia Emas 2045 with outcome matrix for key capabilities), and Capability Architecture Final Statement (official declaration with compliance requirements for all subsequent architecture work). Document now serves as Capability Architecture v1.0, ready as official reference for 03_BUSINESS_PROCESS_ARCHITECTURE.md, 04_DATA_ARCHITECTURE.md, 05_AI_ARCHITECTURE.md, 06_APPLICATION_ARCHITECTURE.md, MVP Architecture, and SDLC Architecture with significantly reduced architectural risk.
- **Version 1.0** (June 2026): Initial Capability Model defining 16 Level 1 capabilities, comprehensive Level 2 decomposition, critical Level 3 capabilities, AI Capability Model with 90% automation target, AI Agent Capability Matrix with 11 official AI agents, Human Capability Model for 6 stakeholder groups, Capability Ownership Matrix, Capability Dependency Map, Capability Maturity Model with 5 levels, MVP Capability Prioritization across 4 waves, Capability Success Metrics, complete traceability to domains and AI agents, and full alignment validation with foundation documents (00A, 00B, 00C, 01). Document serves as official reference for all subsequent architecture work with significantly reduced architectural risk.

---

**Curriculum Before Code. Pedagogy Before Technology. AI-First Architecture. Human Flourishing Through AI.**

# 17_SPRINT_1_EXECUTION_PLAN.md

## Foundation Document for NUSA Education Platform

**Version**: 1.0
**Date**: June 2026
**Status**: IMPLEMENTATION DOCUMENT
**Alignment**: Validated against Foundation Architecture (00A, 00B, 00C, 01, 02, 03, 04, 05, 06, 07, 08, 09, 10, 11, 12, 13, 14, 15, 16)

**Purpose**: Define Sprint 1 implementation plan for NUSA MVP Wave 1, serving as the execution blueprint for the first end-to-end vertical slice delivery. This document is the primary planning artifact for Sprint 1.

---

# Sprint 1 Objective

## Sprint 1: MVP Delivery Sprint

Sprint 1 represents the complete MVP Wave 1 implementation effort. The entire MVP Wave 1 scope is planned for completion within this single 20-day implementation cycle.

## User Journey

Teacher can:

1. **Login** to the system
2. **Select CP** (Graduate Profile)
3. **Generate TP** using AI
4. **Generate ATP** using AI
5. **Generate Modul Ajar** using AI
6. **Generate Assessment** using AI
7. **Generate Rubric** using AI
8. **Generate Narrative Report** using AI
9. **Review and Approve** all artifacts

## Success Statement

Sprint 1 is successful when a teacher can complete the entire end-to-end workflow from login through narrative report generation in the staging environment.

---

# Sprint Scope

## Included Modules

- **Authentication Module**: Custom JWT authentication
- **Administration Module**: School and Teacher management
- **Curriculum Module**: Graduate Profile (CP) and Tujuan Pembelajaran (TP)
- **Learning Planning Module**: Alur Tujuan Pembelajaran (ATP) and Modul Ajar
- **Assessment Module**: Assessment, Rubric, and Evidence management
- **Reporting Module**: Narrative Report generation
- **AI Orchestration Module**: All 6 AI Agents (TP, ATP, Modul Ajar, Assessment, Rubric, Narrative Report)
- **Workflow Module**: Workflow history and lifecycle
- **Audit Module**: Audit fields and traceability

## Technical Scope

### Backend
- Go + Gin framework
- PostgreSQL database
- RabbitMQ messaging
- Custom JWT authentication
- All 6 AI Agents integration (TP, ATP, Modul Ajar, Assessment, Rubric, Narrative Report)

### Frontend
- React + TypeScript
- Login page
- CP selection page
- TP generation page
- TP review page
- TP approval page
- ATP generation page
- ATP review page
- ATP approval page
- Modul Ajar generation page
- Modul Ajar review page
- Modul Ajar approval page
- Assessment generation page
- Assessment review page
- Assessment approval page
- Rubric generation page
- Rubric review page
- Rubric approval page
- Narrative Report generation page
- Narrative Report review page
- Narrative Report approval page

### Database
- Users table
- Schools table
- Teachers table
- CP dimensions table
- CP table
- TP table
- ATP table
- Modul Ajar table
- Assessments table
- Rubrics table
- Evidences table
- Evaluations table
- Narrative Reports table
- Workflow history table
- AI audit table

### AI
- All 6 AI Agent prompt specifications (TP, ATP, Modul Ajar, Assessment, Rubric, Narrative Report)
- All 6 AI Agent integrations
- LLM provider connection

---

# Sprint Timeline

## Duration

**20 Days**

## Method

**Mini Scrum**

- Daily standups (15 minutes)
- Weekly sprint planning (1 hour)
- Weekly sprint review (1 hour)
- Weekly retrospective (1 hour)

## Calendar

| Day | Activity |
|-----|----------|
| Day 1 | Sprint planning, task breakdown, environment setup |
| Day 2-5 | Development (Backend: Authentication, Administration, Curriculum) |
| Day 6-10 | Development (Backend: Learning Planning, Assessment, Reporting, AI Orchestration) |
| Day 11-14 | Development (Frontend: All artifact generation and review pages) |
| Day 15-17 | Development (Frontend: Approval workflows, AI integration) |
| Day 18 | Integration testing, bug fixes |
| Day 19 | End-to-end testing, deployment to staging |
| Day 20 | Sprint review, retrospective, MVP delivery |

---

# Deliverables

## Backend Deliverables

### Authentication Module
- [ ] JWT token generation and validation
- [ ] Login endpoint
- [ ] User authentication middleware
- [ ] Password hashing
- [ ] Session management

### Administration Module
- [ ] School CRUD endpoints
- [ ] Teacher CRUD endpoints
- [ ] User-School relationship
- [ ] Teacher-School assignment

### Curriculum Module
- [ ] CP dimension CRUD endpoints
- [ ] CP CRUD endpoints
- [ ] TP CRUD endpoints
- [ ] CP-TP relationship
- [ ] TP versioning (version_no, parent_version_id, is_current_version)

### AI Orchestration Module
- [ ] TP Agent integration
- [ ] LLM provider client
- [ ] Prompt builder
- [ ] AI generation endpoint
- [ ] AI audit logging

### Workflow Module
- [ ] Workflow history table
- [ ] Workflow state transitions
- [ ] Workflow event logging

### Audit Module
- [ ] Audit fields (created_by, updated_by, created_at, updated_at)
- [ ] Audit logging for all entities
- [ ] Audit trail queries

### API Contract
- [ ] Authentication API endpoints
- [ ] School API endpoints
- [ ] Teacher API endpoints
- [ ] CP API endpoints
- [ ] TP API endpoints
- [ ] AI generation API endpoints

## Frontend Deliverables

### Authentication
- [ ] Login page
- [ ] Login form validation
- [ ] JWT token storage
- [ ] Auth context
- [ ] Protected route wrapper

### Administration
- [ ] School list page
- [ ] Teacher list page
- [ ] School detail page
- [ ] Teacher detail page

### Curriculum
- [ ] CP list page
- [ ] CP detail page
- [ ] CP selection component
- [ ] TP list page
- [ ] TP detail page
- [ ] TP version display

### AI Integration
- [ ] TP generation page
- [ ] AI generation trigger
- [ ] Generation progress indicator
- [ ] AI response display
- [ ] TP review page
- [ ] TP approval page
- [ ] TP save functionality

### Common Components
- [ ] Button component
- [ ] Input component
- [ ] Modal component
- [ ] Table component
- [ ] Card component
- [ ] Loading spinner
- [ ] Error message display

## Database Deliverables

### Migrations
- [ ] 000001_init_schema.up.sql
- [ ] 000001_init_schema.down.sql
- [ ] 000002_users.up.sql
- [ ] 000002_users.down.sql
- [ ] 000003_schools.up.sql
- [ ] 000003_schools.down.sql
- [ ] 000004_teachers.up.sql
- [ ] 000004_teachers.down.sql
- [ ] 000005_cp_dimensions.up.sql
- [ ] 000005_cp_dimensions.down.sql
- [ ] 000006_cp.up.sql
- [ ] 000006_cp.down.sql
- [ ] 000007_tp.up.sql
- [ ] 000007_tp.down.sql
- [ ] 000008_workflow_history.up.sql
- [ ] 000008_workflow_history.down.sql
- [ ] 000009_ai_audit.up.sql
- [ ] 000009_ai_audit.down.sql

### Seed Data
- [ ] Seed schools
- [ ] Seed teachers
- [ ] Seed CP dimensions
- [ ] Seed sample CP

## AI Deliverables

### TP Agent
- [ ] TP prompt specification
- [ ] TP prompt template
- [ ] TP prompt examples
- [ ] TP prompt versioning
- [ ] LLM provider configuration
- [ ] AI generation service
- [ ] AI validation rules
- [ ] AI error handling

---

# Task Breakdown

## Developer 1: Backend Authentication & Administration

### Day 1-2: Environment Setup
- [ ] Set up Go development environment
- [ ] Set up PostgreSQL database
- [ ] Set up RabbitMQ
- [ ] Initialize Go modules
- [ ] Set up project structure
- [ ] Configure environment variables

### Day 3-5: Authentication Module
- [ ] Implement JWT token generation
- [ ] Implement JWT token validation
- [ ] Implement password hashing (bcrypt)
- [ ] Create login endpoint
- [ ] Create authentication middleware
- [ ] Write unit tests for authentication
- [ ] Write integration tests for login

### Day 6-8: Administration Module
- [ ] Create School entity and repository
- [ ] Create Teacher entity and repository
- [ ] Implement School CRUD endpoints
- [ ] Implement Teacher CRUD endpoints
- [ ] Implement User-School relationship
- [ ] Implement Teacher-School assignment
- [ ] Write unit tests for administration
- [ ] Write integration tests for administration

### Day 9-10: Integration & Bug Fixes
- [ ] Integrate authentication with administration
- [ ] Fix authentication bugs
- [ ] Fix administration bugs
- [ ] Code review
- [ ] Documentation

## Developer 2: Backend Curriculum Module

### Day 1-2: Environment Setup
- [ ] Set up Go development environment
- [ ] Set up PostgreSQL database
- [ ] Initialize Go modules
- [ ] Set up project structure
- [ ] Configure environment variables

### Day 3-6: CP Module
- [ ] Create CP Dimension entity and repository
- [ ] Create CP entity and repository
- [ ] Implement CP Dimension CRUD endpoints
- [ ] Implement CP CRUD endpoints
- [ ] Implement CP-TP relationship
- [ ] Implement CP versioning
- [ ] Write unit tests for CP
- [ ] Write integration tests for CP

### Day 7-10: TP Module
- [ ] Create TP entity and repository
- [ ] Implement TP CRUD endpoints
- [ ] Implement TP versioning (version_no, parent_version_id, is_current_version)
- [ ] Implement TP lifecycle states (DRAFT, UNDER_REVIEW, APPROVED, REJECTED)
- [ ] Implement TP audit fields
- [ ] Write unit tests for TP
- [ ] Write integration tests for TP

### Day 11-12: Integration & Bug Fixes
- [ ] Integrate CP with TP
- [ ] Fix CP bugs
- [ ] Fix TP bugs
- [ ] Code review
- [ ] Documentation

## Developer 3: Backend AI Orchestration & Workflow

### Day 1-2: Environment Setup
- [ ] Set up Go development environment
- [ ] Set up LLM provider account
- [ ] Configure LLM API keys
- [ ] Initialize Go modules
- [ ] Set up project structure
- [ ] Configure environment variables

### Day 3-6: AI Orchestration Module
- [ ] Create AI Agent entity and repository
- [ ] Implement LLM provider client
- [ ] Implement prompt builder
- [ ] Implement TP Agent integration
- [ ] Implement AI generation endpoint
- [ ] Implement AI audit logging
- [ ] Write unit tests for AI
- [ ] Write integration tests for AI

### Day 7-10: Workflow Module
- [ ] Create Workflow History entity and repository
- [ ] Implement workflow state transitions
- [ ] Implement workflow event logging
- [ ] Integrate workflow with TP lifecycle
- [ ] Implement workflow queries
- [ ] Write unit tests for workflow
- [ ] Write integration tests for workflow

### Day 11-12: Integration & Bug Fixes
- [ ] Integrate AI with workflow
- [ ] Integrate AI with TP
- [ ] Fix AI bugs
- [ ] Fix workflow bugs
- [ ] Code review
- [ ] Documentation

## Developer 4: Frontend Authentication & Administration

### Day 1-2: Environment Setup
- [ ] Set up React development environment
- [ ] Set up TypeScript
- [ ] Set up Vite
- [ ] Set up TailwindCSS
- [ ] Initialize npm modules
- [ ] Set up project structure
- [ ] Configure environment variables

### Day 3-5: Authentication
- [ ] Create login page
- [ ] Create login form
- [ ] Implement form validation
- [ ] Implement JWT token storage
- [ ] Create auth context
- [ ] Create protected route wrapper
- [ ] Write unit tests for authentication
- [ ] Write integration tests for authentication

### Day 6-8: Administration
- [ ] Create school list page
- [ ] Create teacher list page
- [ ] Create school detail page
- [ ] Create teacher detail page
- [ ] Implement API client for administration
- [ ] Write unit tests for administration
- [ ] Write integration tests for administration

### Day 9-10: Common Components
- [ ] Create Button component
- [ ] Create Input component
- [ ] Create Modal component
- [ ] Create Table component
- [ ] Create Card component
- [ ] Create Loading spinner
- [ ] Create Error message display
- [ ] Write unit tests for components

### Day 11-12: Integration & Bug Fixes
- [ ] Integrate authentication with administration
- [ ] Fix authentication bugs
- [ ] Fix administration bugs
- [ ] Code review
- [ ] Documentation

## Developer 5: Frontend Curriculum & AI Integration

### Day 1-2: Environment Setup
- [ ] Set up React development environment
- [ ] Set up TypeScript
- [ ] Set up Vite
- [ ] Set up TailwindCSS
- [ ] Initialize npm modules
- [ ] Set up project structure
- [ ] Configure environment variables

### Day 3-6: Curriculum
- [ ] Create CP list page
- [ ] Create CP detail page
- [ ] Create CP selection component
- [ ] Create TP list page
- [ ] Create TP detail page
- [ ] Implement TP version display
- [ ] Implement API client for curriculum
- [ ] Write unit tests for curriculum
- [ ] Write integration tests for curriculum

### Day 7-10: AI Integration
- [ ] Create TP generation page
- [ ] Implement AI generation trigger
- [ ] Implement generation progress indicator
- [ ] Implement AI response display
- [ ] Create TP review page
- [ ] Create TP approval page
- [ ] Implement TP save functionality
- [ ] Write unit tests for AI integration
- [ ] Write integration tests for AI integration

### Day 11-12: Integration & Bug Fixes
- [ ] Integrate curriculum with AI
- [ ] Fix curriculum bugs
- [ ] Fix AI integration bugs
- [ ] Code review
- [ ] Documentation

---

# Risks

## Technical Risks

### Risk 1: LLM Provider Integration
- **Description**: LLM provider may have rate limits, downtime, or API changes
- **Probability**: Medium
- **Impact**: High (blocks TP generation)
- **Mitigation**: 
  - Implement retry logic with exponential backoff
  - Implement fallback to alternative LLM provider
  - Implement graceful degradation with error messages
  - Monitor LLM provider status
- **Owner**: Developer 3

### Risk 2: Database Schema Changes
- **Description**: Database schema may require changes during development
- **Probability**: Medium
- **Impact**: Medium (requires migration updates)
- **Mitigation**:
  - Follow database migration governance
  - Use version-controlled migrations
  - Test migrations in development before production
  - Document rollback procedures
- **Owner**: All Developers

### Risk 3: AI Prompt Quality
- **Description**: AI prompts may not generate high-quality TP
- **Probability**: Medium
- **Impact**: High (affects user experience)
- **Mitigation**:
  - Use prompt specifications from 15_AI_PROMPT_SPECIFICATION.md
  - Implement prompt versioning
  - Implement human review workflow
  - Collect feedback and iterate on prompts
- **Owner**: Developer 3 + AI Engineer

### Risk 4: Integration Complexity
- **Description**: Frontend and backend integration may have issues
- **Probability**: Medium
- **Impact**: Medium (delays end-to-end testing)
- **Mitigation**:
  - Define API contract early
  - Use API contract testing
  - Implement mock services for frontend development
  - Daily integration testing
- **Owner**: All Developers

## Project Risks

### Risk 5: Timeline Pressure
- **Description**: 20-day timeline may be insufficient for complete implementation
- **Probability**: Medium
- **Impact**: High (sprint failure)
- **Mitigation**:
  - Prioritize critical path features
  - Defer non-critical features to Sprint 2
  - Daily progress tracking
  - Early escalation of blockers
- **Owner**: Principal Delivery Manager

### Risk 6: Resource Availability
- **Description**: Developers may be unavailable due to illness or other commitments
- **Probability**: Low
- **Impact**: Medium (delays implementation)
- **Mitigation**:
  - Cross-train developers on multiple modules
  - Maintain task documentation
  - Have backup plan for critical tasks
- **Owner**: Principal Delivery Manager

---

# Dependencies

## External Dependencies

### LLM Provider
- **Dependency**: LLM provider API access and API key
- **Required By**: Day 3
- **Owner**: Principal Technical Lead
- **Status**: Pending

### PostgreSQL Database
- **Dependency**: PostgreSQL database instance
- **Required By**: Day 1
- **Owner**: DevOps
- **Status**: Pending

### RabbitMQ
- **Dependency**: RabbitMQ message broker instance
- **Required By**: Day 1
- **Owner**: DevOps
- **Status**: Pending

## Internal Dependencies

### Authentication Module
- **Dependency**: Database schema (users table)
- **Required By**: Day 3
- **Owner**: Developer 1
- **Status**: Pending

### Administration Module
- **Dependency**: Authentication module
- **Required By**: Day 6
- **Owner**: Developer 1
- **Status**: Pending

### Curriculum Module
- **Dependency**: Administration module (for school/teacher context)
- **Required By**: Day 3
- **Owner**: Developer 2
- **Status**: Pending

### AI Orchestration Module
- **Dependency**: Curriculum module (for TP generation)
- **Required By**: Day 3
- **Owner**: Developer 3
- **Status**: Pending

### Workflow Module
- **Dependency**: AI Orchestration module (for workflow events)
- **Required By**: Day 7
- **Owner**: Developer 3
- **Status**: Pending

### Frontend Authentication
- **Dependency**: Backend authentication API
- **Required By**: Day 3
- **Owner**: Developer 4
- **Status**: Pending

### Frontend Administration
- **Dependency**: Backend administration API
- **Required By**: Day 6
- **Owner**: Developer 4
- **Status**: Pending

### Frontend Curriculum
- **Dependency**: Backend curriculum API
- **Required By**: Day 3
- **Owner**: Developer 5
- **Status**: Pending

### Frontend AI Integration
- **Dependency**: Backend AI API
- **Required By**: Day 7
- **Owner**: Developer 5
- **Status**: Pending

---

# Definition of Done

## Functional Criteria

- [ ] Teacher can login with valid credentials
- [ ] Teacher can select a CP
- [ ] Teacher can generate TP using AI
- [ ] Teacher can review generated TP
- [ ] Teacher can approve TP
- [ ] Teacher can save TP
- [ ] TP is saved with audit fields
- [ ] Workflow history is recorded
- [ ] AI generation is audited

## Code Quality Criteria

- [ ] All code committed to repository
- [ ] All code reviewed by another developer
- [ ] No critical review findings remain
- [ ] Code follows coding standards from 08_SDLC_ARCHITECTURE.md

## Testing Criteria

- [ ] Unit tests pass for all modules
- [ ] Integration tests pass for all modules
- [ ] End-to-end test passes for TP generation flow
- [ ] Manual testing completed

## API Criteria

- [ ] API implementation matches API Contract (13_API_CONTRACT.md)
- [ ] Request and response structures validated
- [ ] Error handling implemented
- [ ] Authentication middleware implemented

## Database Criteria

- [ ] Migration scripts created
- [ ] Migration successfully executed
- [ ] Rollback strategy verified
- [ ] Seed data loaded

## Frontend Criteria

- [ ] UI implemented according to feature requirements
- [ ] Validation and error states handled
- [ ] Responsive behavior verified
- [ ] Loading states implemented

## AI Criteria

- [ ] Prompt specification implemented (15_AI_PROMPT_SPECIFICATION.md)
- [ ] AI output validated
- [ ] Human review flow verified
- [ ] AI audit logging implemented

## Documentation Criteria

- [ ] API documentation updated
- [ ] Database documentation updated
- [ ] Implementation notes documented

## Deployment Criteria

- [ ] Successfully deployed to staging environment
- [ ] No critical runtime errors
- [ ] End-to-end flow works in staging

---

# Sprint Success Criteria

## Primary Success Criterion

**Sprint is successful when the complete end-to-end TP generation flow works in the staging environment.**

A teacher can:
1. Login to the system
2. Select a CP
3. Generate TP using AI
4. Review the generated TP
5. Approve the TP
6. Save the TP

## Secondary Success Criteria

- [ ] All Definition of Done criteria met
- [ ] No critical bugs remaining
- [ ] Code review completed for all modules
- [ ] Unit test coverage > 80%
- [ ] Integration test coverage > 70%
- [ ] End-to-end test passes
- [ ] Deployment to staging successful
- [ ] Performance targets met (AI generation P95 < 5 seconds)
- [ ] Security review passed
- [ ] Documentation complete

## Failure Criteria

Sprint is considered a failure if:
- End-to-end TP generation flow does not work in staging
- Critical bugs remain at end of sprint
- Security vulnerabilities identified
- Performance targets not met
- Definition of Done criteria not met

---

# Sprint Review

## Review Date

Day 20

## Review Participants

- Principal Delivery Manager
- Principal Technical Lead
- All 5 Developers
- AI Engineer
- QA Engineer (if available)

## Review Agenda

1. Sprint goal review
2. Demo of end-to-end TP generation flow
3. Review of completed deliverables
4. Review of Definition of Done compliance
5. Discussion of risks and issues
6. Lessons learned
7. Sprint 2 planning

## Review Output

- Sprint success/failure determination
- Lessons learned document
- Improvement actions for Sprint 2
- Sprint 2 planning input

---

# Sprint Retrospective

## Retrospective Date

Day 20 (after Sprint Review)

## Retrospective Participants

- All 5 Developers
- Principal Delivery Manager
- Principal Technical Lead

## Retrospective Format

**What went well?**
- List 3-5 things that went well during the sprint

**What didn't go well?**
- List 3-5 things that didn't go well during the sprint

**What can we improve?**
- List 3-5 action items for improvement in Sprint 2

## Retrospective Output

- Lessons learned document
- Improvement action items
- Process updates for Sprint 2

---

# Communication Plan

## Daily Standups

- **Time**: 9:00 AM daily
- **Duration**: 15 minutes
- **Participants**: All 5 Developers, Principal Technical Lead
- **Format**: 
  - What did you complete yesterday?
  - What will you complete today?
  - Any blockers?

## Weekly Planning

- **Time**: Monday 9:30 AM
- **Duration**: 1 hour
- **Participants**: All 5 Developers, Principal Delivery Manager, Principal Technical Lead
- **Format**: Review progress, plan week, identify risks

## Weekly Review

- **Time**: Friday 3:00 PM
- **Duration**: 1 hour
- **Participants**: All 5 Developers, Principal Delivery Manager, Principal Technical Lead
- **Format**: Demo completed work, review progress, plan next week

## Ad-hoc Communication

- **Slack**: For daily communication and quick questions
- **Email**: For formal communication and documentation
- **Video Call**: For urgent issues requiring discussion

---

# Tools and Infrastructure

## Development Tools

- **IDE**: VS Code
- **Version Control**: Git
- **CI/CD**: GitHub Actions
- **Project Management**: GitHub Issues/Projects
- **Communication**: Slack

## Infrastructure

- **Development**: Local development environment
- **Staging**: Docker Compose deployment
- **Database**: PostgreSQL 18+
- **Message Broker**: RabbitMQ
- **LLM Provider**: OpenAI (or alternative)

## Monitoring

- **Logs**: Application logs
- **Metrics**: Basic metrics (response time, error rate)
- **Alerts**: Error alerts for critical failures

---

**Document Status**: IMPLEMENTATION DOCUMENT

**Sprint Status**: PLANNING

**Sprint Start**: TBD

**Sprint End**: TBD

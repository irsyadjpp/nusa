# Target Architecture

**Version:** 1.0  
**Date:** June 13, 2026  
**Based On:** Audit Deliverables (DATABASE_AUDIT.md, API_INVENTORY.md, UI_INVENTORY.md, GAP_ANALYSIS.md, TECHNICAL_DEBT.md, AUDIT_REPORT.md)

---

## Executive Summary

The target architecture maintains the existing modular monolith pattern with DDD Lite principles while addressing critical gaps identified in the audit. The architecture prioritizes maintainability, scalability, simplicity, testability, security, and observability without introducing overengineering.

**Key Decisions:**
- **Pattern:** Continue with Layered Architecture + DDD Lite (no CQRS, no Event Sourcing)
- **Deployment:** Single monolithic deployment with horizontal scaling capability
- **Data:** PostgreSQL with Redis caching layer
- **Communication:** REST API with WebSocket for real-time features
- **Background Processing:** RabbitMQ for async jobs
- **Observability:** Structured logging + Sentry + APM

---

## System Architecture

### High-Level Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                         Client Layer                             │
├─────────────────────────────────────────────────────────────────┤
│  Web Browser (React + TypeScript)                                │
│  - Admin Portal                                                  │
│  - Teacher Portal                                                │
│  - Student Portal                                                │
└─────────────────────────────────────────────────────────────────┘
                              │
                              │ HTTPS
                              │
┌─────────────────────────────────────────────────────────────────┐
│                      API Gateway Layer                           │
├─────────────────────────────────────────────────────────────────┤
│  - Rate Limiting (Redis-based)                                   │
│  - Security Headers                                             │
│  - Request ID Generation                                        │
│  - CORS Configuration                                           │
└─────────────────────────────────────────────────────────────────┘
                              │
                              │
┌─────────────────────────────────────────────────────────────────┐
│                    Application Layer (Go)                         │
├─────────────────────────────────────────────────────────────────┤
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐          │
│  │   Handlers   │  │  App Services│  │   Services   │          │
│  │  (HTTP only) │  │ (Orchestrate)│  │ (Business    │          │
│  └──────────────┘  └──────────────┘  │  Logic)      │          │
│         │                 │           └──────────────┘          │
│         │                 │                  │                  │
│         └─────────────────┴──────────────────┘                  │
│                              │                                  │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐          │
│  │  Middleware  │  │   Domain     │  │ Repositories │          │
│  │  (Auth,      │  │   Layer      │  │ (Data Access)│          │
│  │   RBAC,      │  │ (Entities,   │  │              │          │
│  │   Validation)│  │  Value Obj)  │  └──────────────┘          │
│  └──────────────┘  └──────────────┘          │                  │
│                                             │                  │
└─────────────────────────────────────────────┼──────────────────┘
                                              │
┌─────────────────────────────────────────────┼──────────────────┐
│              Data Layer                      │                  │
├─────────────────────────────────────────────┼──────────────────┤
│  ┌──────────────┐  ┌──────────────┐         │                  │
│  │  PostgreSQL  │  │    Redis     │         │                  │
│  │  (Primary)   │  │   (Cache)    │         │                  │
│  └──────────────┘  └──────────────┘         │                  │
│                                              │                  │
│  ┌──────────────┐  ┌──────────────┐         │                  │
│  │   RabbitMQ   │  │   S3/Cloud   │         │                  │
│  │  (Jobs)      │  │   Storage    │         │                  │
│  └──────────────┘  └──────────────┘         │                  │
└─────────────────────────────────────────────┼──────────────────┘
                                              │
┌─────────────────────────────────────────────┼──────────────────┐
│           Infrastructure & Observability      │                  │
├─────────────────────────────────────────────┼──────────────────┤
│  - Sentry (Error Tracking)                  │                  │
│  - APM (Performance Monitoring)             │                  │
│  - Structured Logging (JSON)                │                  │
│  - Metrics (Prometheus/Grafana)             │                  │
│  - Health Checks                            │                  │
│  - Automated Backups                        │                  │
└─────────────────────────────────────────────┴──────────────────┘
```

---

## Module Boundaries

### Core Modules (Existing)

#### 1. Authentication & Authorization
- **Responsibility:** User authentication, JWT token management, role-based access control
- **Boundaries:** Self-contained, no dependencies on business modules
- **Integration:** Used by all modules via middleware

#### 2. User Management
- **Responsibility:** User CRUD, role assignment, school assignment
- **Boundaries:** Depends on Authentication, Roles
- **Integration:** Used by all modules for user context

#### 3. School Management
- **Responsibility:** School CRUD, school configuration
- **Boundaries:** Self-contained
- **Integration:** Used by Academic Foundation for scoping

#### 4. Academic Foundation
- **Responsibility:** Academic years, semesters, subject categories, graduate profile dimensions, CP alignments, system configurations
- **Boundaries:** Depends on School Management
- **Integration:** Used by Curriculum, Learning Planning

#### 5. Curriculum (CP)
- **Responsibility:** Curriculum subjects, phases, elements, subelements, CPs
- **Boundaries:** Depends on Academic Foundation
- **Integration:** Used by Learning Planning (TP)

#### 6. Learning Planning
- **Responsibility:** TP, TP Sets, ATP, ATP Sets, Modul Ajar, Modul Ajar Sets
- **Boundaries:** Depends on Curriculum, Academic Foundation
- **Integration:** Used by Assessment

#### 7. Assessment
- **Responsibility:** Assessments, rubrics, evidence, evaluations
- **Boundaries:** Depends on Learning Planning
- **Integration:** Used by Achievement, Reporting

#### 8. Achievement
- **Responsibility:** Student achievement, competency progress, class achievement (runtime calculation)
- **Boundaries:** Depends on Assessment
- **Integration:** Used by Reporting

#### 9. Reporting
- **Responsibility:** Narrative reports, achievement summaries
- **Boundaries:** Depends on Achievement
- **Integration:** Final output module

### New Modules (To Be Added)

#### 10. Class Management
- **Responsibility:** Classes, class enrollments, teacher assignments
- **Boundaries:** Depends on Academic Foundation, User Management
- **Integration:** Used by Attendance, Scheduling, Assessment, Reporting

#### 11. Attendance
- **Responsibility:** Attendance records, attendance reports, attendance tracking
- **Boundaries:** Depends on Class Management
- **Integration:** Used by Reporting

#### 12. Scheduling
- **Responsibility:** Class schedules, timetables, conflict detection, room allocation
- **Boundaries:** Depends on Class Management
- **Integration:** Used by Teacher Portal, Student Portal

#### 13. Exam & Assignment
- **Responsibility:** Exams, assignments, exam results, grade book
- **Boundaries:** Depends on Class Management, Assessment
- **Integration:** Used by Reporting, Teacher Portal, Student Portal

#### 14. Communication
- **Responsibility:** Notifications, announcements, messages, email integration
- **Boundaries:** Self-contained, depends on User Management
- **Integration:** Used by all modules for user communication

#### 15. Audit & Logging
- **Responsibility:** Audit logs, change history, activity tracking
- **Boundaries:** Cross-cutting concern, no dependencies
- **Integration:** Used by all modules via middleware

---

## Domain Boundaries

### Bounded Contexts

#### 1. Identity Context
- **Entities:** User, Role, Permission, RefreshToken
- **Value Objects:** Email, PasswordHash, Token
- **Aggregates:** User (with roles and permissions)
- **Responsibility:** Authentication, authorization, user management

#### 2. Organization Context
- **Entities:** School
- **Value Objects:** SchoolCode, Address
- **Aggregates:** School
- **Responsibility:** School management, multi-tenancy scoping

#### 3. Academic Foundation Context
- **Entities:** AcademicYear, Semester, SubjectCategory, GraduateProfileDimension, CPAlignment, SystemConfiguration
- **Value Objects:** DateRange, SequenceNumber, ConfigurationValue
- **Aggregates:** AcademicYear (with semesters)
- **Responsibility:** Academic calendar, curriculum governance

#### 4. Curriculum Context
- **Entities:** CurriculumSubject, CurriculumPhase, CurriculumElement, CurriculumSubelement, CP
- **Value Objects:** CPCode, PhaseLevel, ElementOrder
- **Aggregates:** CP (with subject, phase, element, subelement)
- **Responsibility:** Curriculum planning, CP management

#### 5. Learning Planning Context
- **Entities:** TP, TPSet, ATP, ATPSet, ModulAjar, ModulAjarSet
- **Value Objects:** LearningObjectives, TimeAllocation, KKTPCriteria, Prerequisites
- **Aggregates:** TPSet (with TPs), ATPSet (with ATPs), ModulAjarSet (with ModulAjars)
- **Responsibility:** Teaching plan, ATP, modul ajar generation

#### 6. Assessment Context
- **Entities:** Assessment, Rubric, Evidence, Evaluation
- **Value Objects:** AssessmentItems, AnswerKey, ScoringGuidelines, PerformanceScores
- **Aggregates:** Assessment (with rubric), Evidence (with evaluations)
- **Responsibility:** Assessment design, evidence collection, evaluation

#### 7. Class Management Context
- **Entities:** Class, ClassEnrollment
- **Value Objects:** GradeLevel, Room, Schedule
- **Aggregates:** Class (with enrollments)
- **Responsibility:** Class organization, student-teacher relationships

#### 8. Attendance Context
- **Entities:** AttendanceRecord
- **Value Objects:** AttendanceStatus, AttendanceDate
- **Aggregates:** AttendanceRecord
- **Responsibility:** Attendance tracking, reporting

#### 9. Scheduling Context
- **Entities:** Schedule
- **Value Objects:** DayOfWeek, TimeSlot, Room
- **Aggregates:** Schedule
- **Responsibility:** Class scheduling, conflict detection

#### 10. Exam Context
- **Entities:** Exam, Assignment, ExamResult
- **Value Objects:** ExamDate, Score, Grade
- **Aggregates:** Exam (with results)
- **Responsibility:** Exam management, grade tracking

#### 11. Communication Context
- **Entities:** Notification, Announcement, Message
- **Value Objects:** NotificationType, MessageContent
- **Aggregates:** Notification, Announcement, Message
- **Responsibility:** User communication, notifications

#### 12. Achievement Context
- **Entities:** StudentAchievement, CompetencyProgress, ClassAchievement
- **Value Objects:** MasteryLevel, ProgressPercentage
- **Aggregates:** StudentAchievement (with competency progress)
- **Responsibility:** Achievement calculation, reporting

#### 13. Reporting Context
- **Entities:** NarrativeReport
- **Value Objects:** NarrativeContent, ReportingPeriod
- **Aggregates:** NarrativeReport
- **Responsibility:** Report generation, distribution

---

## Responsibilities

### Layer Responsibilities

#### Handler Layer
- **Responsibility:** HTTP request/response handling only
- **Rules:**
  - No business logic
  - Request validation (DTOs)
  - Response formatting
  - Error translation to HTTP status codes
- **Integration:** Calls Application Services

#### Application Service Layer
- **Responsibility:** Use case orchestration
- **Rules:**
  - Coordinate between Services
  - Transaction management
  - Permission checks
  - Response assembly
- **Integration:** Calls Services, Repositories

#### Service Layer
- **Responsibility:** Business logic
- **Rules:**
  - Domain logic implementation
  - Business rule enforcement
  - Domain object manipulation
- **Integration:** Uses Domain objects, calls Repositories

#### Repository Layer
- **Responsibility:** Data access only
- **Rules:**
  - Database operations only
  - No business logic
  - SQL queries only
- **Integration:** Database only

#### Domain Layer
- **Responsibility:** Business invariants and behavior
- **Rules:**
  - Entity behavior
  - Value object validation
  - Aggregate consistency
  - Business rule enforcement
- **Integration:** No external dependencies

### Module Responsibilities

#### Authentication Module
- JWT token generation and validation
- Refresh token rotation
- Password hashing
- Login/logout
- Session management

#### User Management Module
- User CRUD operations
- Role assignment
- School assignment
- User status management
- User profile management

#### Academic Foundation Module
- Academic year lifecycle (DRAFT → ACTIVE → ARCHIVED)
- Semester management
- Subject category management
- Graduate profile dimension management
- CP alignment management
- System configuration management
- Business rule enforcement (BR-001 to BR-004)

#### Curriculum Module
- Curriculum subject management
- Phase management
- Element management
- Subelement management
- CP management
- CP import/export

#### Learning Planning Module
- TP set generation (manual/AI)
- TP management with versioning
- ATP set generation
- ATP management
- Modul Ajar set generation
- Modul Ajar management
- Approval workflows

#### Assessment Module
- Assessment creation and management
- Rubric management
- Evidence upload and management
- Evaluation with revision tracking
- AI-assisted assessment generation

#### Class Management Module
- Class creation and management
- Student enrollment
- Teacher assignment
- Class configuration

#### Attendance Module
- Attendance recording
- Attendance tracking
- Attendance reporting
- Attendance analytics

#### Scheduling Module
- Class schedule creation
- Timetable management
- Conflict detection
- Room allocation
- Teacher workload tracking

#### Exam Module
- Exam creation and scheduling
- Assignment creation and management
- Exam result recording
- Grade book management
- Grade calculation

#### Communication Module
- Notification creation and delivery
- Announcement management
- Messaging system
- Email integration
- Real-time notifications (WebSocket)

#### Achievement Module
- Student achievement calculation (runtime)
- Competency progress tracking
- Class achievement calculation
- Mastery level determination

#### Reporting Module
- Narrative report generation
- Achievement summary generation
- Report distribution
- Report approval workflow

#### Audit Module
- Audit log creation
- Change history tracking
- Activity logging
- Audit report generation

---

## Integration Flows

### Authentication Flow

```
Client → API Gateway → Handler (Login) → Application Service (Login)
  → Service (User) → Repository (User) → Database
  → Service (JWT) → Generate Tokens → Application Service
  → Handler → Response (Tokens)
```

### Learning Planning Flow

```
Client → API Gateway → Handler (Create TP) → Application Service (Create TP)
  → Service (TP) → Domain (TP) → Validate Business Rules
  → Repository (TP) → Database
  → Application Service → Handler → Response
```

### Assessment Flow

```
Client → API Gateway → Handler (Create Assessment) → Application Service
  → Service (Assessment) → Domain (Assessment) → Validate
  → Repository (Assessment) → Database
  → Application Service → Handler → Response
```

### Achievement Calculation Flow

```
Client → API Gateway → Handler (Get Achievement) → Application Service
  → Service (Achievement) → Repository (Evidence) → Database
  → Service (Achievement) → Repository (Evaluation) → Database
  → Service (Achievement) → Calculate Achievement (Runtime)
  → Application Service → Handler → Response
```

### Notification Flow

```
Application Service (Any Module) → Service (Notification)
  → Repository (Notification) → Database
  → Service (Notification) → RabbitMQ (Notification Queue)
  → Background Worker → WebSocket → Client
  → Service (Notification) → Email Service → Email Provider
```

### Audit Logging Flow

```
Any Operation → Middleware (Audit) → Service (Audit)
  → Repository (Audit) → Database
  → Continue with original operation
```

### Background Job Flow

```
Application Service → RabbitMQ (Job Queue)
  → Background Worker → Process Job
  → Service (Domain) → Repository → Database
  → Update Job Status → RabbitMQ (Result Queue)
```

---

## Cross-Cutting Concerns

### Security
- Authentication: JWT with refresh token rotation
- Authorization: RBAC with permission checks
- Rate Limiting: Redis-based per-user and per-IP limits
- Security Headers: CSP, HSTS, X-Frame-Options, X-XSS-Protection
- Input Validation: Comprehensive validation layer
- SQL Injection Prevention: Parameterized queries only

### Observability
- Structured Logging: JSON format with correlation IDs
- Error Tracking: Sentry integration
- Performance Monitoring: APM integration
- Metrics: Prometheus metrics for key operations
- Health Checks: /health, /ready endpoints
- Audit Logging: Comprehensive audit trail

### Caching
- Strategy: Redis for server-side caching
- Cache Keys: Namespace-based (e.g., `academic_year:{id}`)
- Cache Invalidation: Event-based on mutations
- Cache Warming: Pre-populate frequently accessed data
- Cache TTL: Configurable per data type

### Background Processing
- Queue: RabbitMQ for async jobs
- Job Types: Email sending, report generation, data export, AI processing
- Retry Strategy: Exponential backoff with max retries
- Dead Letter Queue: Failed job handling
- Job Status Tracking: Persistent job status in database

### Data Consistency
- Transactions: Database transactions for multi-step operations
- Optimistic Locking: Version-based for concurrent updates
- Eventual Consistency: For background operations
- Data Validation: Domain-level validation before persistence

---

## Scalability Strategy

### Horizontal Scaling
- Stateless application layer (multiple instances)
- Load balancer for API gateway
- Database read replicas for read-heavy operations
- Redis clustering for cache layer
- RabbitMQ clustering for message queue

### Vertical Scaling
- Connection pooling configuration
- Database indexing optimization
- Query optimization
- Bundle size optimization (frontend)

### Data Partitioning
- School-based partitioning (multi-tenant)
- Time-based partitioning for audit logs
- Consider sharding for large tables (evidence, evaluations)

---

## Technology Stack

### Backend
- **Language:** Go 1.21+
- **Framework:** Gin Gonic
- **Database:** PostgreSQL 15+
- **Cache:** Redis 7+
- **Message Queue:** RabbitMQ 3.12+
- **File Storage:** S3-compatible storage
- **Logging:** Structured JSON logging
- **Error Tracking:** Sentry
- **APM:** Datadog/New Relic

### Frontend
- **Framework:** React 19+
- **Language:** TypeScript 5.8+
- **Build Tool:** Vite 7+
- **UI Framework:** Material-UI v7
- **Styling:** Tailwind CSS v4
- **State Management:** TanStack Query + Zustand
- **Routing:** React Router v7
- **Forms:** Formik + Yup
- **Real-time:** WebSocket client

### Infrastructure
- **Containerization:** Docker
- **Deployment:** Kubernetes or Docker Compose
- **CI/CD:** GitHub Actions
- **Monitoring:** Prometheus + Grafana
- **Logging:** ELK Stack or CloudWatch
- **Backup:** Automated PostgreSQL backups with point-in-time recovery

---

## Non-Functional Requirements

### Performance
- API response time: < 200ms (p95)
- Page load time: < 2s (p95)
- Database query time: < 100ms (p95)
- Cache hit rate: > 80%

### Availability
- Uptime: 99.5% (monthly)
- Recovery Time Objective (RTO): 4 hours
- Recovery Point Objective (RPO): 1 hour

### Security
- Authentication: JWT with 24h access token expiration
- Authorization: RBAC with permission-based access
- Data encryption: TLS 1.3 in transit, at-rest encryption for sensitive data
- Compliance: GDPR-like data protection principles

### Maintainability
- Code coverage: > 70%
- Documentation: Comprehensive API and developer documentation
- Code review: Required for all changes
- Automated testing: Unit, integration, E2E tests

### Scalability
- Concurrent users: 10,000+
- Data volume: 1TB+ (with partitioning)
- API requests: 1,000+ requests/second

---

## Deployment Architecture

### Development Environment
- Single instance deployment
- Local PostgreSQL and Redis
- No monitoring (basic logging only)
- Manual deployments

### Staging Environment
- Multi-instance deployment (2 instances)
- Managed PostgreSQL and Redis
- Basic monitoring (Sentry, health checks)
- CI/CD deployments

### Production Environment
- Multi-instance deployment (3+ instances with auto-scaling)
- Managed PostgreSQL with read replicas
- Redis cluster
- RabbitMQ cluster
- Comprehensive monitoring (Sentry, APM, metrics, logging)
- Automated deployments with blue-green strategy
- Automated backups with disaster recovery

---

## Migration Strategy

### From Current to Target Architecture

**Phase 1: Infrastructure Foundation**
- Add Redis caching layer
- Add RabbitMQ for background jobs
- Add Sentry for error tracking
- Add security headers middleware
- Add rate limiting middleware
- Implement database backup strategy

**Phase 2: Critical Features**
- Implement class management module
- Implement attendance module
- Implement scheduling module
- Implement communication module
- Implement exam/assignment module

**Phase 3: Quality & Observability**
- Add comprehensive audit logging
- Add performance monitoring
- Add database indexing
- Implement soft delete
- Add test coverage

**Phase 4: Enhancement**
- Optimize caching strategy
- Add advanced search
- Add export functionality
- Enhance dashboard
- Add analytics

---

## Conclusion

The target architecture maintains the existing modular monolith pattern with DDD Lite principles while addressing critical gaps identified in the audit. The architecture prioritizes maintainability, scalability, simplicity, testability, security, and observability without introducing overengineering.

The key additions are:
- Class management, attendance, scheduling, communication, and exam modules
- Redis caching layer for performance
- RabbitMQ for background processing
- Comprehensive observability (Sentry, APM, structured logging)
- Security enhancements (rate limiting, security headers, audit logging)
- Quality improvements (test coverage, database indexing, soft delete)

This architecture provides a solid foundation for scaling the NUSA Platform to support Indonesian schools implementing Kurikulum Merdeka 2026.

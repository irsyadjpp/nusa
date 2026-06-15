# Backend Test Strategy

## Overview

This document defines the comprehensive testing strategy for the NUSA Platform backend, based on official Go testing patterns, Context7 documentation research, and the test inventory analysis.

**Generated:** 2026-06-15
**Backend Technology:** Go (Gin, sqlx, PostgreSQL)
**Testing Framework:** Go testing package + gomock for mocking
**Quality Gates:** Service ≥ 90%, Repository ≥ 90%, Controller ≥ 90%, Critical Business Rule = 100%, Authorization = 100%, Validation = 100%, Transaction = 100%

---

## Testing Philosophy

### Goals
- **NOT** to maximize code coverage percentage
- **NOT** to test implementation details
- **YES** to verify backend behaves correctly under real-world conditions
- **YES** to test business rules and domain invariants
- **YES** to ensure security and data integrity

### Principles
1. **Test behavior, not implementation** - Focus on what the code does, not how it does it
2. **Use official testing patterns** - Follow Go testing best practices from Context7 documentation
3. **Mock appropriately** - Mock external dependencies, not domain logic
4. **Test boundaries** - Verify layer boundaries and contracts
5. **Real database for repositories** - Don't mock database operations

---

## Component Classification

### Pure Unit Tests

**Definition:** Tests that don't require external dependencies, databases, or network calls.

**Examples:**
- Validators (DTO validation, business rule validation)
- DTO Mappers (request/response serialization)
- Utility Functions (string manipulation, date formatting)
- Value Objects (domain value object validation)
- Domain Model methods (invariants, validation logic)

**Testing Approach:**
- Use standard Go testing package
- Table-driven tests for multiple scenarios
- Test happy path and error cases
- No mocking required
- Fast execution (milliseconds)

**Context7 Documentation Used:** `/golang/go` - Go testing patterns, table-driven tests

### Service Unit Tests

**Definition:** Tests for application service layer business logic with mocked dependencies.

**Examples:**
- Business Rules (assessment validation, achievement calculation)
- Domain Logic (state transitions, invariants)
- Permission Checks (authorization logic)
- Orchestration (coordinating multiple repositories)

**Testing Approach:**
- Use gomock for repository interfaces
- Mock all external dependencies
- Test business rules independently
- Test error handling and edge cases
- Verify correct repository method calls
- No database access

**Context7 Documentation Used:** `/uber-go/mock` - Mock generation and expectations

### Integration Tests

**Definition:** Tests that require real external dependencies (database, cache, message queue).

**Examples:**
- Repository CRUD operations (database operations)
- Database Constraints (foreign keys, unique constraints)
- Complex Queries (joins, aggregations)
- Cache Operations (Redis get/set/delete)
- API Endpoints (HTTP request/response flow)

**Testing Approach:**
- Use dedicated test database
- Use test Redis instance
- Real HTTP requests using httptest
- Complete request flow testing
- Transaction rollback testing
- Constraint validation testing

**Context7 Documentation Used:** `/jmoiron/sqlx` - Database testing patterns, `/golang/go` - HTTP testing with httptest

### System Tests

**Definition:** End-to-end tests that verify complete workflows across multiple components.

**Examples:**
- Full Request Flow (HTTP → Handler → Service → Repository → Database)
- Authentication Flow (Login → Token → Protected Resource)
- Authorization Flow (Role-based access control)
- Business Workflows (Assessment → Evidence → Evaluation → Achievement)

**Testing Approach:**
- Complete system setup (database, cache, queue)
- Real HTTP requests
- Full integration of all layers
- Workflow validation
- Performance testing (optional)

---

## Module-by-Module Strategy

### Authentication & Authorization Module

**Components:**
- auth_middleware.go
- role.go (middleware)
- resource_authorization.go
- pkg/jwt/service.go

**Test Classification:**
- Pure Unit: JWT token generation/validation logic
- Service Unit: Permission checks, role verification
- Integration: Authentication flow, authorization flow
- System: Complete auth workflow

**Test Scope:**
```
Pure Unit Tests (JWT Service):
- Token generation with valid claims
- Token validation with valid token
- Token validation with expired token
- Token validation with invalid signature
- Token validation with malformed token
- Claim extraction and verification

Service Unit Tests (Authorization):
- Permission check for valid user
- Permission check for invalid user
- Role verification for valid role
- Role verification for invalid role
- Resource ownership validation
- School boundary enforcement

Integration Tests (Auth Middleware):
- Valid token passes middleware
- Invalid token rejected with 401
- Expired token rejected with 401
- Missing token rejected with 401
- Token refresh flow
- Logout flow

System Tests (Complete Auth Flow):
- User registration
- Email verification (if implemented)
- Login with valid credentials
- Login with invalid credentials
- Password reset (if implemented)
- Protected resource access with valid token
- Protected resource access with expired token
```

**Mocking Requirements:**
- Mock user repository for token generation
- Mock cache for token blacklist (if implemented)
- Real JWT library for crypto operations

**Database Requirements:**
- Test database for user CRUD operations
- Test database for token storage (if persistent)

**Context7 Documentation:**
- `/golang/go` - HTTP testing with httptest
- `/uber-go/mock` - Mock repository interfaces

---

### Achievement Module

**Components:**
- achievement_service.go
- achievement_handler.go
- achievement_repository.go
- achievement.go (domain)

**Test Classification:**
- Pure Unit: Achievement calculation logic, domain invariants
- Service Unit: Achievement service business rules
- Integration: Repository CRUD operations, API endpoints
- System: Complete achievement calculation workflow

**Test Scope:**
```
Pure Unit Tests (Achievement Domain):
- Student achievement calculation logic
- Competency progress calculation
- Achievement summary generation
- Class achievement aggregation
- Domain invariant validation (achievement thresholds)

Service Unit Tests (Achievement Service):
- Calculate student achievement
- Calculate competency progress
- Generate achievement summary
- Generate class achievement
- Error handling for missing data
- Permission checks for achievement access

Integration Tests (Achievement Repository):
- Create achievement record
- Read achievement by student ID
- Update achievement scores
- Delete achievement record
- Query achievements by competency
- Query achievements by class
- Pagination and sorting

API Integration Tests (Achievement Handler):
- GET /achievement/students/:id/achievement
- GET /achievement/students/:id/progress
- GET /achievement/classes/:id/achievement
- GET /achievement/reports/:id/achievement-summary
- Authorization checks
- Error handling (404, 403, 500)

System Tests (Complete Workflow):
- Evidence submission → Achievement calculation
- Multiple assessments → Aggregate achievement
- Competency progress over time
- Class achievement summary generation
```

**Mocking Requirements:**
- Mock evidence repository for achievement calculation
- Mock evaluation repository for scoring
- Mock assessment repository for reference data
- Mock reporting repository for narrative report integration

**Database Requirements:**
- Test database for achievement CRUD operations
- Test database for evidence and evaluation data
- Test database for assessment reference data

**Context7 Documentation:**
- `/jmoiron/sqlx` - Database operations and transactions
- `/uber-go/mock` - Mock repository interfaces
- `/golang/go` - Table-driven tests for calculation scenarios

---

### Assessment Module

**Components:**
- assessment_service.go
- assessment_handler.go
- assessment_repository.go
- assessment.go (domain)
- assessment_dto.go

**Test Classification:**
- Pure Unit: Assessment validation, DTO mapping, domain invariants
- Service Unit: Assessment lifecycle, versioning, snapshot logic
- Integration: Repository CRUD operations, API endpoints
- System: Complete assessment workflow

**Test Scope:**
```
Pure Unit Tests (Assessment Domain):
- Assessment validation rules
- Success criteria validation
- TP reference validation
- Assessment state transitions
- Version increment logic
- Snapshot creation logic

Service Unit Tests (Assessment Service):
- Create assessment with valid data
- Create assessment with invalid data
- Update assessment version
- Archive assessment
- Assessment permission checks
- Assessment business rules

Integration Tests (Assessment Repository):
- Create assessment
- Read assessment by ID
- Update assessment
- Delete assessment (soft delete)
- Query assessments by TP
- Query assessments by subject
- Version history queries

API Integration Tests (Assessment Handler):
- POST /assessments
- GET /assessments/:id
- PUT /assessments/:id
- DELETE /assessments/:id
- GET /assessments (with filters)
- Validation errors (400, 422)
- Authorization checks

DTO Tests (Assessment DTO):
- Request DTO validation
- Response DTO serialization
- JSON binding/unbinding
- Type conversion and mapping

System Tests (Complete Workflow):
- TP creation → Assessment creation
- Assessment update → Version increment
- Assessment archive → Status change
- Assessment snapshot → Historical consistency
```

**Mocking Requirements:**
- Mock TP repository for reference validation
- Mock user repository for permission checks
- Mock evaluation repository for cascade operations

**Database Requirements:**
- Test database for assessment CRUD operations
- Test database for TP reference data
- Test database for evaluation cascade

**Context7 Documentation:**
- `/jmoiron/sqlx` - Database operations and versioning
- `/uber-go/mock` - Mock repository interfaces
- `/golang/go` - HTTP testing with httptest

---

### TP (Tujuan Pembelajaran) Module

**Components:**
- tp_service.go
- tp_handler.go
- tp_repository.go
- tp.go (domain)
- learning_planning_service.go
- learning_planning_repository.go
- learning_planning.go (domain)

**Test Classification:**
- Pure Unit: TP validation, KKTP criteria validation, ATP/Modul Ajar logic
- Service Unit: TP lifecycle, versioning, approval workflow
- Integration: Repository CRUD operations, API endpoints
- System: Complete TP workflow (TP → ATP → Modul Ajar)

**Test Scope:**
```
Pure Unit Tests (TP Domain):
- TP validation rules
- KKTP criteria validation (mastery thresholds, performance indicators)
- TP state transitions (draft → approved → published)
- Version increment logic
- Parent version tracking
- Approval workflow logic

Service Unit Tests (TP Service):
- Create TP with valid data
- Create TP with invalid KKTP
- Update TP with versioning
- Approve TP
- Reject TP
- Archive TP
- TP permission checks

Integration Tests (TP Repository):
- Create TP
- Read TP by ID
- Update TP
- Delete TP (soft delete)
- Query TPs by CP
- Query TPs by phase
- Version history queries
- Current version queries

API Integration Tests (TP Handler):
- POST /tp
- GET /tp/:id
- PUT /tp/:id
- DELETE /tp/:id
- POST /tp/:id/approve
- POST /tp/:id/reject
- GET /tp (with filters)
- KKTP validation errors

System Tests (Complete Workflow):
- CP creation → TP creation
- TP approval → ATP generation
- ATP → Modul Ajar creation
- TP versioning → Historical tracking
- TP rejection → Workflow termination
```

**Mocking Requirements:**
- Mock CP repository for reference validation
- Mock ATP repository for cascade operations
- Mock user repository for permission checks

**Database Requirements:**
- Test database for TP CRUD operations
- Test database for CP reference data
- Test database for ATP cascade operations

**Context7 Documentation:**
- `/jmoiron/sqlx` - Database operations and versioning
- `/uber-go/mock` - Mock repository interfaces
- `/golang/go` - HTTP testing with httptest

---

### Class Management Module

**Components:**
- class_service.go
- class_handler.go
- class_repository.go
- class.go (domain)
- class_dto.go

**Test Classification:**
- Pure Unit: Class validation, DTO mapping, domain invariants
- Service Unit: Class lifecycle, student assignment, Wali Kelas assignment
- Integration: Repository CRUD operations, API endpoints
- System: Complete class workflow

**Test Scope:**
```
Pure Unit Tests (Class Domain):
- Class validation rules
- Class capacity validation
- Class schedule validation
- Student assignment validation
- Wali Kelas assignment validation
- Class state transitions

Service Unit Tests (Class Service):
- Create class with valid data
- Assign student to class
- Remove student from class
- Assign Wali Kelas
- Update class details
- Archive class
- Class permission checks

Integration Tests (Class Repository):
- Create class
- Read class by ID
- Update class
- Delete class (soft delete)
- Query classes by school
- Query classes by academic year
- Student assignment queries
- Wali Kelas assignment queries

API Integration Tests (Class Handler):
- POST /classes
- GET /classes/:id
- PUT /classes/:id
- DELETE /classes/:id
- POST /classes/:id/students
- DELETE /classes/:id/students/:studentId
- POST /classes/:id/wali-kelas
- GET /classes (with filters)
- Capacity validation errors

DTO Tests (Class DTO):
- Request DTO validation
- Response DTO serialization
- Student assignment DTO validation
- JSON binding/unbinding

System Tests (Complete Workflow):
- Class creation → Student enrollment
- Student assignment → Class capacity validation
- Wali Kelas assignment → Permission verification
- Class update → Historical tracking
- Class archive → Status change
```

**Mocking Requirements:**
- Mock user repository for student validation
- Mock school repository for school validation
- Mock academic year repository for period validation

**Database Requirements:**
- Test database for class CRUD operations
- Test database for student reference data
- Test database for school and academic year data

**Context7 Documentation:**
- `/jmoiron/sqlx` - Database operations and relationship queries
- `/uber-go/mock` - Mock repository interfaces
- `/golang/go` - HTTP testing with httptest

---

### Evidence & Evaluation Module

**Components:**
- evidence_service.go (in assessment_service.go)
- evaluation_service.go
- evidence_handler.go
- evaluation_handler.go
- evidence_repository.go
- evaluation_repository.go
- evidence.go (domain in assessment.go)
- evaluation.go (domain)

**Test Classification:**
- Pure Unit: Evidence validation, evaluation scoring, domain invariants
- Service Unit: Evidence upload, evaluation submission, revision logic
- Integration: Repository CRUD operations, file storage operations
- System: Complete evidence → evaluation → achievement workflow

**Test Scope:**
```
Pure Unit Tests (Evidence Domain):
- Evidence validation rules
- File metadata validation
- Evidence state transitions
- Evaluation scoring logic
- Revision increment logic
- Teacher feedback validation

Service Unit Tests (Evidence Service):
- Upload evidence with valid data
- Upload evidence with invalid file
- Update evidence metadata
- Archive evidence
- Evidence permission checks

Service Unit Tests (Evaluation Service):
- Submit evaluation with valid score
- Submit evaluation with invalid score
- Update evaluation with revision
- Evaluation revision logic
- Teacher feedback validation
- Evaluation permission checks

Integration Tests (Evidence Repository):
- Create evidence
- Read evidence by ID
- Update evidence
- Delete evidence (soft delete)
- Query evidence by assessment
- Query evidence by student
- File metadata queries

Integration Tests (Evaluation Repository):
- Create evaluation
- Read evaluation by ID
- Update evaluation
- Delete evaluation (soft delete)
- Query evaluations by evidence
- Query evaluations by teacher
- Revision history queries

API Integration Tests (Evidence Handler):
- POST /evidence
- GET /evidence/:id
- PUT /evidence/:id
- DELETE /evidence/:id
- GET /evidence (with filters)
- File upload validation
- Authorization checks

API Integration Tests (Evaluation Handler):
- POST /evaluations
- GET /evaluations/:id
- PUT /evaluations/:id
- DELETE /evaluations/:id
- GET /evaluations (with filters)
- Revision history endpoint
- Score validation errors

System Tests (Complete Workflow):
- Evidence upload → File storage
- Evidence submission → Evaluation creation
- Evaluation submission → Score validation
- Evaluation revision → Historical tracking
- Evaluation completion → Achievement calculation
```

**Mocking Requirements:**
- Mock MinIO storage for file operations
- Mock assessment repository for reference validation
- Mock user repository for permission checks
- Mock achievement service for trigger

**Database Requirements:**
- Test database for evidence CRUD operations
- Test database for evaluation CRUD operations
- Test database for assessment reference data
- Test MinIO instance for file storage (or mock)

**Context7 Documentation:**
- `/jmoiron/sqlx` - Database operations and revision tracking
- `/uber-go/mock` - Mock repository interfaces and storage
- `/golang/go` - HTTP testing with httptest

---

### Reporting Module

**Components:**
- reporting_service.go
- reporting_handler.go
- reporting_repository.go
- reporting.go (domain)
- reporting_dto.go

**Test Classification:**
- Pure Unit: Report validation, narrative content validation, domain invariants
- Service Unit: Report generation, achievement data refresh, publish logic
- Integration: Repository CRUD operations, API endpoints
- System: Complete reporting workflow (achievement → report → publish)

**Test Scope:**
```
Pure Unit Tests (Reporting Domain):
- Narrative report validation
- Report content validation
- Report state transitions
- Achievement summary validation
- Publish workflow validation

Service Unit Tests (Reporting Service):
- Generate narrative report
- Generate academic report
- Refresh achievement data
- Publish report
- Unpublish report
- Archive report
- Report permission checks

Integration Tests (Reporting Repository):
- Create narrative report
- Read report by ID
- Update report
- Delete report (soft delete)
- Query reports by student
- Query reports by class
- Query reports by academic period
- Published report queries

API Integration Tests (Reporting Handler):
- POST /narrative-reports
- GET /narrative-reports/:id
- PUT /narrative-reports/:id
- DELETE /narrative-reports/:id
- POST /narrative-reports/:id/publish
- POST /narrative-reports/:id/unpublish
- GET /narrative-reports (with filters)
- Achievement refresh endpoint

DTO Tests (Reporting DTO):
- Request DTO validation
- Response DTO serialization
- Narrative content validation
- JSON binding/unbinding

System Tests (Complete Workflow):
- Achievement calculation → Report generation
- Report generation → Achievement integration
- Report publish → Status change
- Report update → Revision tracking
- Achievement refresh → Data consistency
```

**Mocking Requirements:**
- Mock achievement service for data refresh
- Mock evaluation repository for evidence data
- Mock assessment repository for reference data
- Mock user repository for permission checks

**Database Requirements:**
- Test database for report CRUD operations
- Test database for achievement reference data
- Test database for evaluation and assessment data

**Context7 Documentation:**
- `/jmoiron/sqlx` - Database operations and complex queries
- `/uber-go/mock` - Mock repository interfaces
- `/golang/go` - HTTP testing with httptest

---

### User Management Module

**Components:**
- user_service.go
- user_handler.go
- user_repository.go
- user.go (domain)
- user_dto.go

**Test Classification:**
- Pure Unit: User validation, DTO mapping, domain invariants
- Service Unit: User CRUD, password hashing, authentication logic
- Integration: Repository CRUD operations, API endpoints
- System: Complete user workflow (registration → verification → login)

**Test Scope:**
```
Pure Unit Tests (User Domain):
- User validation rules
- Email validation
- Password strength validation
- User state transitions
- Role assignment validation

Service Unit Tests (User Service):
- Create user with valid data
- Create user with invalid email
- Update user details
- Change password
- Assign role
- Archive user
- User permission checks

Integration Tests (User Repository):
- Create user
- Read user by ID
- Read user by email
- Update user
- Delete user (soft delete)
- Query users by school
- Query users by role
- Permission queries

API Integration Tests (User Handler):
- POST /users
- GET /users/:id
- PUT /users/:id
- DELETE /users/:id
- POST /users/:id/change-password
- POST /users/:id/assign-role
- GET /users (with filters)
- Email validation errors
- Password strength validation errors

DTO Tests (User DTO):
- Request DTO validation
- Response DTO serialization
- Password change DTO validation
- JSON binding/unbinding

System Tests (Complete Workflow):
- User registration → Email verification (if implemented)
- User login → Token generation
- Password change → Hash validation
- Role assignment → Permission update
- User archive → Status change
```

**Mocking Requirements:**
- Mock JWT service for token generation
- Mock school repository for school validation
- Mock role repository for role validation

**Database Requirements:**
- Test database for user CRUD operations
- Test database for school and role reference data

**Context7 Documentation:**
- `/jmoiron/sqlx` - Database operations and authentication queries
- `/uber-go/mock` - Mock repository interfaces
- `/golang/go` - HTTP testing with httptest

---

### Curriculum Management Module

**Components:**
- curriculum_service.go
- curriculum_handler.go
- curriculum_repository.go
- curriculum.go (domain)
- curriculum_dto.go
- academic_year_service.go
- semester_service.go
- subject_category_repository.go

**Test Classification:**
- Pure Unit: Curriculum validation, DTO mapping, domain invariants
- Service Unit: Curriculum CRUD, academic year management, semester management
- Integration: Repository CRUD operations, API endpoints
- System: Complete curriculum workflow

**Test Scope:**
```
Pure Unit Tests (Curriculum Domain):
- Curriculum validation rules
- Subject validation
- Phase validation
- Element validation
- Subelement validation
- CP (Capaian Pembelajaran) validation

Service Unit Tests (Curriculum Service):
- Create curriculum with valid data
- Create subject with valid data
- Create phase with valid data
- Update curriculum structure
- Archive curriculum
- Curriculum permission checks

Service Unit Tests (Academic Year Service):
- Create academic year
- Update academic year dates
- Activate academic year
- Deactivate academic year
- Academic year permission checks

Service Unit Tests (Semester Service):
- Create semester
- Update semester dates
- Activate semester
- Deactivate semester
- Semester permission checks

Integration Tests (Curriculum Repository):
- Create curriculum subject
- Read subject by ID
- Update subject
- Delete subject (soft delete)
- Query subjects by school
- Query subjects by academic year
- Phase and element queries

API Integration Tests (Curriculum Handler):
- POST /curriculum/subjects
- GET /curriculum/subjects/:id
- PUT /curriculum/subjects/:id
- DELETE /curriculum/subjects/:id
- POST /academic-years
- GET /academic-years/:id
- PUT /academic-years/:id
- POST /semesters
- GET /semesters/:id
- Curriculum validation errors

DTO Tests (Curriculum DTO):
- Request DTO validation
- Response DTO serialization
- Curriculum structure validation
- JSON binding/unbinding

System Tests (Complete Workflow):
- Subject creation → Phase assignment
- Phase creation → Element definition
- Element creation → Subelement definition
- Academic year activation → Semester activation
- Curriculum archive → Status change
```

**Mocking Requirements:**
- Mock school repository for school validation
- Mock academic year repository for period validation
- Mock user repository for permission checks

**Database Requirements:**
- Test database for curriculum CRUD operations
- Test database for academic year and semester data
- Test database for school reference data

**Context7 Documentation:**
- `/jmoiron/sqlx` - Database operations and hierarchical queries
- `/uber-go/mock` - Mock repository interfaces
- `/golang/go` - HTTP testing with httptest

---

### Exam & Assignment Module

**Components:**
- exam_service.go
- exam_handler.go
- exam_repository.go
- exam.go (domain)
- exam_result_service.go
- exam_result_handler.go
- exam_result_repository.go
- exam_result.go (domain)
- assignment_service.go
- assignment_handler.go
- assignment_repository.go
- assignment.go (domain)

**Test Classification:**
- Pure Unit: Exam/assignment validation, result calculation, domain invariants
- Service Unit: Exam/assignment lifecycle, result processing, grading logic
- Integration: Repository CRUD operations, API endpoints
- System: Complete exam/assignment workflow

**Test Scope:**
```
Pure Unit Tests (Exam Domain):
- Exam validation rules
- Exam schedule validation
- Exam conflict detection
- Exam state transitions

Pure Unit Tests (Assignment Domain):
- Assignment validation rules
- Assignment due date validation
- Assignment state transitions

Pure Unit Tests (Exam Result Domain):
- Result validation rules
- Score validation
- Grade calculation logic
- Result aggregation logic

Service Unit Tests (Exam Service):
- Create exam with valid data
- Create exam with schedule conflict
- Update exam details
- Schedule exam
- Cancel exam
- Exam permission checks

Service Unit Tests (Assignment Service):
- Create assignment with valid data
- Update assignment details
- Archive assignment
- Assignment permission checks

Service Unit Tests (Exam Result Service):
- Submit exam result with valid score
- Calculate grade from score
- Aggregate class results
- Result permission checks

Integration Tests (Exam Repository):
- Create exam
- Read exam by ID
- Update exam
- Delete exam (soft delete)
- Query exams by class
- Query exams by subject
- Schedule conflict queries

Integration Tests (Exam Result Repository):
- Create exam result
- Read result by ID
- Update result
- Delete result (soft delete)
- Query results by exam
- Query results by student
- Aggregate result queries

API Integration Tests (Exam Handler):
- POST /exams
- GET /exams/:id
- PUT /exams/:id
- DELETE /exams/:id
- POST /exams/:id/schedule
- GET /exams (with filters)
- Schedule conflict errors

API Integration Tests (Assignment Handler):
- POST /assignments
- GET /assignments/:id
- PUT /assignments/:id
- DELETE /assignments/:id
- GET /assignments (with filters)

API Integration Tests (Exam Result Handler):
- POST /exam-results
- GET /exam-results/:id
- PUT /exam-results/:id
- GET /exam-results (with filters)
- Score validation errors

System Tests (Complete Workflow):
- Exam creation → Schedule validation
- Exam scheduling → Conflict detection
- Exam completion → Result submission
- Result submission → Grade calculation
- Grade calculation → Achievement update
```

**Mocking Requirements:**
- Mock class repository for schedule validation
- Mock student repository for enrollment validation
- Mock achievement service for result processing

**Database Requirements:**
- Test database for exam and assignment CRUD operations
- Test database for exam result CRUD operations
- Test database for class and student reference data

**Context7 Documentation:**
- `/jmoiron/sqlx` - Database operations and scheduling queries
- `/uber-go/mock` - Mock repository interfaces
- `/golang/go` - HTTP testing with httptest

---

### Schedule Module

**Components:**
- schedule_service.go
- schedule_handler.go
- schedule_repository.go
- schedule.go (domain)

**Test Classification:**
- Pure Unit: Schedule validation, conflict detection, domain invariants
- Service Unit: Schedule CRUD, conflict resolution, time slot management
- Integration: Repository CRUD operations, API endpoints
- System: Complete scheduling workflow

**Test Scope:**
```
Pure Unit Tests (Schedule Domain):
- Schedule validation rules
- Time slot validation
- Schedule conflict detection
- Resource availability validation
- Schedule state transitions

Service Unit Tests (Schedule Service):
- Create schedule with valid data
- Create schedule with conflict
- Update schedule
- Resolve schedule conflict
- Archive schedule
- Schedule permission checks

Integration Tests (Schedule Repository):
- Create schedule
- Read schedule by ID
- Update schedule
- Delete schedule (soft delete)
- Query schedules by class
- Query schedules by teacher
- Query schedules by room
- Conflict queries

API Integration Tests (Schedule Handler):
- POST /schedules
- GET /schedules/:id
- PUT /schedules/:id
- DELETE /schedules/:id
- GET /schedules (with filters)
- Conflict detection errors
- Time slot validation errors

System Tests (Complete Workflow):
- Schedule creation → Conflict detection
- Conflict resolution → Schedule update
- Schedule validation → Resource booking
- Schedule archive → Status change
```

**Mocking Requirements:**
- Mock class repository for class validation
- Mock teacher repository for teacher validation
- Mock room repository for room validation (if exists)

**Database Requirements:**
- Test database for schedule CRUD operations
- Test database for class, teacher, and room reference data

**Context7 Documentation:**
- `/jmoiron/sqlx` - Database operations and scheduling queries
- `/uber-go/mock` - Mock repository interfaces
- `/golang/go` - HTTP testing with httptest

---

### Notification & Message Module

**Components:**
- notification_service.go
- notification_handler.go
- notification_repository.go
- notification.go (domain)
- message_service.go
- message_handler.go
- message_repository.go
- message.go (domain)

**Test Classification:**
- Pure Unit: Notification validation, message validation, domain invariants
- Service Unit: Notification delivery, message sending, priority logic
- Integration: Repository CRUD operations, API endpoints
- System: Complete notification/message workflow

**Test Scope:**
```
Pure Unit Tests (Notification Domain):
- Notification validation rules
- Notification type validation
- Priority validation
- Recipient validation

Pure Unit Tests (Message Domain):
- Message validation rules
- Message content validation
- Message state transitions

Service Unit Tests (Notification Service):
- Create notification with valid data
- Send notification to recipients
- Mark notification as read
- Archive notification
- Notification permission checks

Service Unit Tests (Message Service):
- Create message with valid data
- Send message to recipients
- Mark message as read
- Archive message
- Message permission checks

Integration Tests (Notification Repository):
- Create notification
- Read notification by ID
- Update notification
- Delete notification (soft delete)
- Query notifications by user
- Query notifications by type
- Unread notification queries

Integration Tests (Message Repository):
- Create message
- Read message by ID
- Update message
- Delete message (soft delete)
- Query messages by user
- Query messages by thread
- Unread message queries

API Integration Tests (Notification Handler):
- POST /notifications
- GET /notifications/:id
- PUT /notifications/:id/read
- DELETE /notifications/:id
- GET /notifications (with filters)
- Notification validation errors

API Integration Tests (Message Handler):
- POST /messages
- GET /messages/:id
- PUT /messages/:id/read
- DELETE /messages/:id
- GET /messages (with filters)
- Message validation errors

System Tests (Complete Workflow):
- Notification creation → Recipient validation
- Notification delivery → Read status update
- Message creation → Thread management
- Message sending → Delivery confirmation
- Notification archive → Status change
```

**Mocking Requirements:**
- Mock user repository for recipient validation
- Mock cache for notification queue (if implemented)
- Mock email/SMS service for delivery (if implemented)

**Database Requirements:**
- Test database for notification and message CRUD operations
- Test database for user reference data

**Context7 Documentation:**
- `/jmoiron/sqlx` - Database operations and recipient queries
- `/uber-go/mock` - Mock repository interfaces
- `/golang/go` - HTTP testing with httptest

---

### System Configuration Module

**Components:**
- system_configuration_service.go
- system_configuration_handler.go
- system_configuration_repository.go
- system_configuration.go (domain)

**Test Classification:**
- Pure Unit: Configuration validation, domain invariants
- Service Unit: Configuration CRUD, validation logic, default values
- Integration: Repository CRUD operations, API endpoints
- System: Complete configuration workflow

**Test Scope:**
```
Pure Unit Tests (System Configuration Domain):
- Configuration validation rules
- Configuration type validation
- Configuration value validation
- Configuration scope validation

Service Unit Tests (System Configuration Service):
- Create configuration with valid data
- Update configuration value
- Reset configuration to default
- Archive configuration
- Configuration permission checks

Integration Tests (System Configuration Repository):
- Create configuration
- Read configuration by key
- Update configuration
- Delete configuration (soft delete)
- Query configurations by scope
- Query configurations by type
- Default value queries

API Integration Tests (System Configuration Handler):
- POST /system-configurations
- GET /system-configurations/:id
- PUT /system-configurations/:id
- DELETE /system-configurations/:id
- GET /system-configurations (with filters)
- Configuration validation errors
- Type validation errors

System Tests (Complete Workflow):
- Configuration creation → Validation
- Configuration update → Value validation
- Configuration reset → Default value restoration
- Configuration archive → Status change
```

**Mocking Requirements:**
- Mock cache for configuration caching (if implemented)
- Mock school repository for scope validation

**Database Requirements:**
- Test database for configuration CRUD operations
- Test database for school reference data

**Context7 Documentation:**
- `/jmoiron/sqlx` - Database operations and configuration queries
- `/uber-go/mock` - Mock repository interfaces
- `/golang/go` - HTTP testing with httptest

---

## Testing Infrastructure

### Test Database Setup

**Requirements:**
- Dedicated test PostgreSQL database
- Separate from development and production databases
- Auto-migration before test runs
- Cleanup after test runs
- Seed data for consistent test scenarios

**Configuration:**
```go
// Use environment variable for test database
TEST_DB_URL=postgresql://user:password@localhost:5432/nusa_test
```

**Context7 Documentation:** `/jmoiron/sqlx` - Database connection and schema setup

### Mock Generation

**Tool:** `mockgen` from uber-go/mock

**Installation:**
```bash
go install go.uber.org/mock/mockgen@latest
```

**Usage:**
```bash
// Generate mocks for repository interfaces
mockgen -source=internal/repository/tp_repository.go -destination=internal/repository/mock/mock_tp_repository.go

// Generate mocks for service interfaces
mockgen -source=internal/service/tp_service.go -destination=internal/service/mock/mock_tp_service.go
```

**Context7 Documentation:** `/uber-go/mock` - Mock generation and usage

### Test Organization

**Directory Structure:**
```
backend/
├── internal/
│   ├── domain/
│   │   ├── tp_test.go              # Domain unit tests
│   │   ├── assessment_test.go      # Domain unit tests
│   │   └── ...
│   ├── service/
│   │   ├── tp_service_test.go      # Service unit tests
│   │   ├── assessment_service_test.go
│   │   └── mock/                   # Generated mocks
│   │       ├── mock_tp_repository.go
│   │       └── ...
│   ├── repository/
│   │   ├── tp_repository_test.go   # Repository integration tests
│   │   ├── assessment_repository_test.go
│   │   └── ...
│   ├── handler/
│   │   ├── tp_handler_test.go      # API integration tests
│   │   ├── assessment_handler_test.go
│   │   └── ...
│   └── middleware/
│       ├── auth_middleware_test.go
│       └── ...
├── pkg/
│   ├── jwt/
│   │   ├── service_test.go         # Package-level tests
│   │   └── security_test.go
│   └── ...
└── tests/
    ├── integration/
    │   ├── auth_test.go            # Integration tests
    │   ├── full_flow_test.go
    │   └── ...
    └── main_test.go                # Test setup
```

### Test Data Management

**Fixtures:**
- Reusable test data in `testdata/` directories
- JSON fixtures for complex data structures
- Factory functions for domain objects
- Builder pattern for test object creation

**Example:**
```go
// testdata/tp_fixture.json
{
  "id": "test-tp-001",
  "cp_id": "test-cp-001",
  "title": "Test TP",
  "learning_objectives": ["Objective 1", "Objective 2"],
  "success_criteria": {
    "mastery_threshold": 75,
    "performance_indicators": ["Indicator 1"]
  }
}
```

### Test Execution

**Run all tests:**
```bash
cd backend
go test ./...
```

**Run specific package tests:**
```bash
go test ./internal/service/
go test ./internal/repository/
go test ./internal/handler/
```

**Run with coverage:**
```bash
go test ./... -cover
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

**Run with race detector:**
```bash
go test ./... -race
```

**Context7 Documentation:** `/golang/go` - Test execution and coverage

---

## Quality Gates

### Coverage Requirements

| Layer | Target Coverage | Critical Rules |
|-------|----------------|----------------|
| Services (Application) | ≥ 90% | Business rules, domain invariants |
| Repositories (Data Access) | ≥ 90% | CRUD operations, transactions |
| Controllers (Handlers) | ≥ 90% | API endpoints, authorization |
| Critical Business Rules | 100% | Achievement calculation, assessment logic |
| Authorization | 100% | Authentication, permission checks |
| Validation | 100% | Input validation, sanitization |
| Transaction | 100% | Rollback scenarios, commit verification |

### Success Criteria

**Testing is complete when:**
1. All quality gates pass
2. All P0 CRITICAL components have comprehensive tests
3. All integration tests pass consistently
4. No race conditions detected
5. Coverage targets met for all layers
6. Security tests pass (authentication, authorization, input validation)
7. Transaction tests pass (rollback scenarios)
8. Performance tests meet requirements (if applicable)

---

## Context7 Documentation References

The following Context7 documentation was used to inform this test strategy:

1. **Go Testing Framework** (`/golang/go`)
   - Basic test function structure
   - Table-driven tests for multiple scenarios
   - Test organization and naming conventions
   - Subtests for test organization
   - Coverage reporting
   - Race detection

2. **Go Mock** (`/uber-go/mock`)
   - Mock generation using mockgen
   - Setting expectations on mocked methods
   - Custom return values and behavior
   - Controller lifecycle management
   - DoAndReturn for custom logic
   - AnyTimes for call frequency

3. **SQLx Database Testing** (`/jmoiron/sqlx`)
   - Database connection setup for testing
   - Transaction testing patterns
   - Query testing with sqlx
   - Named parameter testing
   - Schema setup and cleanup
   - Test data management

4. **HTTP Testing** (`/golang/go` - net/http/httptest)
   - ResponseRecorder for handler testing
   - Test server creation
   - HTTP request creation for testing
   - TLS server testing
   - Request context handling

---

## Implementation Phases

### Phase 3: Unit Testing (In Progress)
- Focus on P0 CRITICAL components
- Pure unit tests for domain models
- Service unit tests with mocks
- DTO validation tests

### Phase 4: Repository Integration Tests
- Real database testing
- CRUD operation testing
- Transaction testing
- Constraint validation

### Phase 5: API Integration Tests
- HTTP endpoint testing
- Authentication and authorization
- Input validation
- Error handling

### Phase 6: Transaction Tests
- Rollback scenarios
- Commit verification
- Concurrent transaction testing

### Phase 7: Event Tests
- Domain event emission
- Event consumption
- Event failure handling

### Phase 8: Background Job Tests
- Job processing
- Retry logic
- Timeout handling
- Idempotency verification

### Phase 9: Concurrency Tests
- Race condition detection
- Deadlock prevention
- Concurrent request handling

### Phase 10: Security Tests
- Authentication bypass
- Authorization escalation
- Input validation
- SQL injection protection

### Phase 11: Coverage Analysis
- Coverage report generation
- Gap analysis
- Risk assessment
- Coverage optimization

### Phase 12: Quality Gates Verification
- Verify all quality gates pass
- Generate final test report
- Document remaining risks
- Production readiness assessment

---

**Document Status:** ✅ Complete
**Next Phase:** Phase 3 - Unit Testing Implementation

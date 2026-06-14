# Backend Design

**Version:** 1.0  
**Date:** June 13, 2026  
**Based On:** AUDIT_REPORT.md, TARGET_ARCHITECTURE.md, DATABASE_DESIGN.md, API_SPEC.md

---

## Executive Summary

This document describes the backend design for the NUSA Platform, including services, repositories, domain logic, event flows, background jobs, queue design, and caching strategy. The design maintains the existing layered architecture with DDD Lite principles while adding new modules and infrastructure components.

**Key Decisions:**
- **Pattern:** Continue with Layered Architecture (Handler → Application Service → Service → Repository → Domain)
- **DDD Lite:** Aggregates, bounded contexts, domain invariants
- **Background Processing:** RabbitMQ for async jobs
- **Caching:** Redis for server-side caching
- **No CQRS/Event Sourcing:** Per architecture freeze

---

## Service Layer Design

### Service Responsibilities

**Service Layer (Business Logic):**
- Domain logic implementation
- Business rule enforcement
- Domain object manipulation
- No external dependencies (except repositories)
- No HTTP concerns

**Application Service Layer (Orchestration):**
- Use case orchestration
- Transaction management
- Permission checks
- Response assembly
- Calls Services and Repositories

### Service Structure

#### Existing Services

**AuthenticationService**
- JWT token generation and validation
- Refresh token rotation
- Password hashing
- Login/logout logic

**UserService**
- User CRUD business logic
- Role assignment validation
- School assignment validation
- User status transitions

**AcademicYearService**
- Academic year lifecycle (DRAFT → ACTIVE → ARCHIVED)
- Business rule enforcement (BR-001 to BR-003)
- Date overlap prevention
- Lead time validation

**SemesterService**
- Semester creation and management
- Business rule enforcement (sequence uniqueness, date overlap)
- Semester activation logic

**SubjectCategoryService**
- Subject category CRUD
- Code uniqueness validation
- Mandatory flag validation

**GraduateProfileDimensionService**
- Graduate profile dimension CRUD
- Code and sequence uniqueness validation
- Sequence range validation (1-6)

**CPAlignmentService**
- CP alignment CRUD
- Bulk creation
- Duplicate prevention
- Coverage report generation

**SystemConfigurationService**
- System configuration CRUD
- Key uniqueness validation
- System configuration protection

**CurriculumService**
- Curriculum subject/phase/element/subelement CRUD
- CP management
- CP import/export
- Code uniqueness validation

**TPService**
- TP set generation (manual/AI)
- TP CRUD with versioning
- Business rule enforcement
- Approval workflow

**ATPService**
- ATP set generation
- ATP CRUD
- Week sequence validation

**ModulAjarService**
- Modul Ajar set generation
- Modul Ajar CRUD
- Session number validation

**AssessmentService**
- Assessment CRUD with versioning
- AI-assisted generation
- Approval workflow

**RubricService**
- Rubric CRUD
- Assessment association

**EvidenceService**
- Evidence CRUD
- File upload handling
- Status transitions

**EvaluationService**
- Evaluation CRUD with revision tracking
- Performance score calculation
- Teacher feedback preservation

**AchievementService**
- Student achievement calculation (runtime)
- Competency progress calculation
- Class achievement calculation
- Mastery level determination

**NarrativeReportService**
- Narrative report generation
- Achievement summary integration
- Approval workflow

#### New Services

**ClassService**
```go
type ClassService struct {
    classRepo repository.ClassRepository
    academicYearRepo repository.AcademicYearRepository
    semesterRepo repository.SemesterRepository
    subjectRepo repository.SubjectRepository
    userRepo repository.UserRepository
}

// Business Logic
- Class creation validation (academic year active, semester active)
- Teacher assignment validation (teacher role, availability)
- Student capacity validation
- Class activation/deactivation
- Class archiving
```

**ClassEnrollmentService**
```go
type ClassEnrollmentService struct {
    enrollmentRepo repository.ClassEnrollmentRepository
    classRepo repository.ClassRepository
    userRepo repository.UserRepository
}

// Business Logic
- Enrollment validation (class active, student not enrolled)
- Capacity check
- Status transitions (ACTIVE → INACTIVE → WITHDRAWN → COMPLETED)
- Duplicate prevention
```

**AttendanceService**
```go
type AttendanceService struct {
    attendanceRepo repository.AttendanceRepository
    classRepo repository.ClassRepository
    userRepo repository.UserRepository
}

// Business Logic
- Attendance recording validation (class active, date valid)
- Status validation (PRESENT, ABSENT, LATE, EXCUSED)
- Duplicate prevention (class_id + student_id + date)
- Attendance rate calculation
- Report generation
```

**ScheduleService**
```go
type ScheduleService struct {
    scheduleRepo repository.ScheduleRepository
    classRepo repository.ClassRepository
    userRepo repository.UserRepository
}

// Business Logic
- Schedule creation validation (class active, time valid)
- Conflict detection (teacher, room, time overlap)
- Day of week validation (1-7)
- Time validation (end_time > start_time)
- Teacher workload tracking
```

**ExamService**
```go
type ExamService struct {
    examRepo repository.ExamRepository
    assessmentRepo repository.AssessmentRepository
    classRepo repository.ClassRepository
}

// Business Logic
- Exam creation validation (assessment approved, class active)
- Date/time validation
- Duration validation
- Status transitions (SCHEDULED → IN_PROGRESS → COMPLETED → CANCELLED)
- Conflict detection
```

**AssignmentService**
```go
type AssignmentService struct {
    assignmentRepo repository.AssignmentRepository
    assessmentRepo repository.AssessmentRepository
    classRepo repository.ClassRepository
}

// Business Logic
- Assignment creation validation (assessment approved, class active)
- Due date validation
- Max score validation
- Status transitions (ASSIGNED → IN_PROGRESS → SUBMITTED → GRADED → CANCELLED)
```

**ExamResultService**
```go
type ExamService struct {
    examResultRepo repository.ExamResultRepository
    examRepo repository.ExamRepository
    userRepo repository.UserRepository
}

// Business Logic
- Result recording validation (exam completed, student enrolled)
- Score validation (0-100)
- Grade calculation
- Duplicate prevention (exam_id + student_id)
```

**NotificationService**
```go
type NotificationService struct {
    notificationRepo repository.NotificationRepository
    userRepo repository.UserRepository
    queueService *QueueService
}

// Business Logic
- Notification creation
- User validation
- Type validation (INFO, WARNING, ERROR, SUCCESS)
- Queue for delivery (WebSocket, email)
- Read status management
```

**AnnouncementService**
```go
type AnnouncementService struct {
    announcementRepo repository.AnnouncementRepository
    schoolRepo repository.SchoolRepository
    userRepo repository.UserRepository
}

// Business Logic
- Announcement creation validation (school active)
- Priority validation (LOW, NORMAL, HIGH, URGENT)
- Target audience validation
- Expiration validation
- Active status management
```

**MessageService**
```go
type MessageService struct {
    messageRepo repository.MessageRepository
    userRepo repository.UserRepository
}

// Business Logic
- Message creation validation (sender/receiver valid)
- Subject validation
- Content validation
- Reply validation (parent_message_id exists)
- Read status management
- Conversation threading
```

**AuditService**
```go
type AuditService struct {
    auditRepo repository.AuditRepository
}

// Business Logic
- Audit log creation
- Change tracking (old_values, new_values)
- User context capture
- IP address capture
- Request ID capture
- Audit report generation
```

---

## Repository Layer Design

### Repository Responsibilities

**Repository Layer (Data Access):**
- Database operations only
- SQL queries only
- No business logic
- Return domain objects or DTOs

### Repository Structure

#### New Repositories

**ClassRepository**
```go
type ClassRepository interface {
    Create(ctx context.Context, class *domain.Class) error
    GetByID(ctx context.Context, id string) (*domain.Class, error)
    List(ctx context.Context, filter ClassFilter) ([]*domain.Class, error)
    Update(ctx context.Context, class *domain.Class) error
    Delete(ctx context.Context, id string) error
    GetStudentCount(ctx context.Context, classID string) (int, error)
    ExistsByName(ctx context.Context, schoolID, name string) (bool, error)
}
```

**ClassEnrollmentRepository**
```go
type ClassEnrollmentRepository interface {
    Create(ctx context.Context, enrollment *domain.ClassEnrollment) error
    GetByID(ctx context.Context, id string) (*domain.ClassEnrollment, error)
    List(ctx context.Context, filter EnrollmentFilter) ([]*domain.ClassEnrollment, error)
    Update(ctx context.Context, enrollment *domain.ClassEnrollment) error
    Delete(ctx context.Context, id string) error
    IsEnrolled(ctx context.Context, classID, studentID string) (bool, error)
    GetActiveEnrollments(ctx context.Context, classID string) ([]*domain.ClassEnrollment, error)
}
```

**AttendanceRepository**
```go
type AttendanceRepository interface {
    Create(ctx context.Context, attendance *domain.Attendance) error
    CreateBulk(ctx context.Context, attendances []*domain.Attendance) error
    GetByID(ctx context.Context, id string) (*domain.Attendance, error)
    List(ctx context.Context, filter AttendanceFilter) ([]*domain.Attendance, error)
    Update(ctx context.Context, attendance *domain.Attendance) error
    Delete(ctx context.Context, id string) error
    Exists(ctx context.Context, classID, studentID, date string) (bool, error)
    GetAttendanceRate(ctx context.Context, classID, dateFrom, dateTo string) (float64, error)
}
```

**ScheduleRepository**
```go
type ScheduleRepository interface {
    Create(ctx context.Context, schedule *domain.Schedule) error
    GetByID(ctx context.Context, id string) (*domain.Schedule, error)
    List(ctx context.Context, filter ScheduleFilter) ([]*domain.Schedule, error)
    Update(ctx context.Context, schedule *domain.Schedule) error
    Delete(ctx context.Context, id string) error
    CheckTeacherConflict(ctx context.Context, teacherID string, dayOfWeek int, startTime, endTime string) (bool, error)
    CheckRoomConflict(ctx context.Context, room string, dayOfWeek int, startTime, endTime string) (bool, error)
}
```

**ExamRepository**
```go
type ExamRepository interface {
    Create(ctx context.Context, exam *domain.Exam) error
    GetByID(ctx context.Context, id string) (*domain.Exam, error)
    List(ctx context.Context, filter ExamFilter) ([]*domain.Exam, error)
    Update(ctx context.Context, exam *domain.Exam) error
    Delete(ctx context.Context, id string) error
    GetUpcoming(ctx context.Context, classID string) ([]*domain.Exam, error)
}
```

**AssignmentRepository**
```go
type AssignmentRepository interface {
    Create(ctx context.Context, assignment *domain.Assignment) error
    GetByID(ctx context.Context, id string) (*domain.Assignment, error)
    List(ctx context.Context, filter AssignmentFilter) ([]*domain.Assignment, error)
    Update(ctx context.Context, assignment *domain.Assignment) error
    Delete(ctx context.Context, id string) error
    GetUpcoming(ctx context.Context, classID string) ([]*domain.Assignment, error)
}
```

**ExamResultRepository**
```go
type ExamResultRepository interface {
    Create(ctx context.Context, result *domain.ExamResult) error
    GetByID(ctx context.Context, id string) (*domain.ExamResult, error)
    List(ctx context.Context, filter ExamResultFilter) ([]*domain.ExamResult, error)
    Update(ctx context.Context, result *domain.ExamResult) error
    Delete(ctx context.Context, id string) error
    Exists(ctx context.Context, examID, studentID string) (bool, error)
    GetByExam(ctx context.Context, examID string) ([]*domain.ExamResult, error)
    GetByStudent(ctx context.Context, studentID string) ([]*domain.ExamResult, error)
}
```

**NotificationRepository**
```go
type NotificationRepository interface {
    Create(ctx context.Context, notification *domain.Notification) error
    GetByID(ctx context.Context, id string) (*domain.Notification, error)
    List(ctx context.Context, filter NotificationFilter) ([]*domain.Notification, error)
    Update(ctx context.Context, notification *domain.Notification) error
    Delete(ctx context.Context, id string) error
    MarkAsRead(ctx context.Context, id string) error
    MarkAllAsRead(ctx context.Context, userID string) error
    GetUnreadCount(ctx context.Context, userID string) (int, error)
}
```

**AnnouncementRepository**
```go
type AnnouncementRepository interface {
    Create(ctx context.Context, announcement *domain.Announcement) error
    GetByID(ctx context.Context, id string) (*domain.Announcement, error)
    List(ctx context.Context, filter AnnouncementFilter) ([]*domain.Announcement, error)
    Update(ctx context.Context, announcement *domain.Announcement) error
    Delete(ctx context.Context, id string) error
    GetActive(ctx context.Context, schoolID string) ([]*domain.Announcement, error)
}
```

**MessageRepository**
```go
type MessageRepository interface {
    Create(ctx context.Context, message *domain.Message) error
    GetByID(ctx context.Context, id string) (*domain.Message, error)
    List(ctx context.Context, filter MessageFilter) ([]*domain.Message, error)
    Update(ctx context.Context, message *domain.Message) error
    Delete(ctx context.Context, id string) error
    MarkAsRead(ctx context.Context, id string) error
    GetConversation(ctx context.Context, userID1, userID2 string) ([]*domain.Message, error)
}
```

**AuditRepository**
```go
type AuditRepository interface {
    Create(ctx context.Context, audit *domain.AuditLog) error
    GetByID(ctx context.Context, id string) (*domain.AuditLog, error)
    List(ctx context.Context, filter AuditFilter) ([]*domain.AuditLog, error)
    GetByEntity(ctx context.Context, entityType, entityID string) ([]*domain.AuditLog, error)
    GetByUser(ctx context.Context, userID string) ([]*domain.AuditLog, error)
}
```

---

## Domain Logic Design

### Domain Entities

#### New Domain Entities

**Class**
```go
type Class struct {
    ID              string
    SchoolID        string
    AcademicYearID  string
    SemesterID      string
    SubjectID       string
    TeacherID       string
    Name            string
    GradeLevel      string
    Room            string
    MaxStudents     int
    IsActive        bool
    CreatedAt       time.Time
    UpdatedAt       time.Time
    CreatedBy       string
    UpdatedBy       string
    DeletedAt       *time.Time
}

// Domain Methods
func (c *Class) Validate() error {
    if c.Name == "" {
        return errors.New("class name is required")
    }
    if c.GradeLevel == "" {
        return errors.New("grade level is required")
    }
    if c.MaxStudents < 1 || c.MaxStudents > 100 {
        return errors.New("max students must be between 1 and 100")
    }
    return nil
}

func (c *Class) CanActivate() error {
    if c.DeletedAt != nil {
        return errors.New("cannot activate deleted class")
    }
    return nil
}
```

**ClassEnrollment**
```go
type ClassEnrollment struct {
    ID             string
    ClassID        string
    StudentID      string
    EnrollmentDate time.Time
    Status         EnrollmentStatus
    Notes          string
    CreatedAt      time.Time
    UpdatedAt      time.Time
    DeletedAt      *time.Time
}

type EnrollmentStatus string

const (
    EnrollmentStatusActive    EnrollmentStatus = "ACTIVE"
    EnrollmentStatusInactive  EnrollmentStatus = "INACTIVE"
    EnrollmentStatusWithdrawn EnrollmentStatus = "WITHDRAWN"
    EnrollmentStatusCompleted EnrollmentStatus = "COMPLETED"
)

// Domain Methods
func (e *ClassEnrollment) Validate() error {
    if e.EnrollmentDate.IsZero() {
        return errors.New("enrollment date is required")
    }
    return e.ValidateStatus()
}

func (e *ClassEnrollment) ValidateStatus() error {
    switch e.Status {
    case EnrollmentStatusActive, EnrollmentStatusInactive, 
         EnrollmentStatusWithdrawn, EnrollmentStatusCompleted:
        return nil
    default:
        return errors.New("invalid enrollment status")
    }
}

func (e *ClassEnrollment) CanWithdraw() error {
    if e.Status != EnrollmentStatusActive {
        return errors.New("only active enrollments can be withdrawn")
    }
    return nil
}
```

**AttendanceRecord**
```go
type AttendanceRecord struct {
    ID          string
    ClassID     string
    StudentID   string
    Date        time.Time
    Status      AttendanceStatus
    Notes       string
    RecordedBy  string
    CreatedAt   time.Time
    UpdatedAt   time.Time
    DeletedAt   *time.Time
}

type AttendanceStatus string

const (
    AttendanceStatusPresent AttendanceStatus = "PRESENT"
    AttendanceStatusAbsent  AttendanceStatus = "ABSENT"
    AttendanceStatusLate    AttendanceStatus = "LATE"
    AttendanceStatusExcused AttendanceStatus = "EXCUSED"
)

// Domain Methods
func (a *AttendanceRecord) Validate() error {
    if a.Date.IsZero() {
        return errors.New("date is required")
    }
    return a.ValidateStatus()
}

func (a *AttendanceRecord) ValidateStatus() error {
    switch a.Status {
    case AttendanceStatusPresent, AttendanceStatusAbsent,
         AttendanceStatusLate, AttendanceStatusExcused:
        return nil
    default:
        return errors.New("invalid attendance status")
    }
}
```

**Schedule**
```go
type Schedule struct {
    ID          string
    ClassID     string
    DayOfWeek   int
    StartTime   string
    EndTime     string
    Room        string
    IsActive    bool
    CreatedAt   time.Time
    UpdatedAt   time.Time
    CreatedBy   string
    UpdatedBy   string
    DeletedAt   *time.Time
}

// Domain Methods
func (s *Schedule) Validate() error {
    if s.DayOfWeek < 1 || s.DayOfWeek > 7 {
        return errors.New("day of week must be between 1 and 7")
    }
    if err := s.ValidateTime(); err != nil {
        return err
    }
    return nil
}

func (s *Schedule) ValidateTime() error {
    start, err := time.Parse("15:04:05", s.StartTime)
    if err != nil {
        return errors.New("invalid start time format")
    }
    end, err := time.Parse("15:04:05", s.EndTime)
    if err != nil {
        return errors.New("invalid end time format")
    }
    if end.Before(start) || end.Equal(start) {
        return errors.New("end time must be after start time")
    }
    return nil
}

func (s *Schedule) Overlaps(other *Schedule) bool {
    if s.DayOfWeek != other.DayOfWeek {
        return false
    }
    sStart, _ := time.Parse("15:04:05", s.StartTime)
    sEnd, _ := time.Parse("15:04:05", s.EndTime)
    oStart, _ := time.Parse("15:04:05", other.StartTime)
    oEnd, _ := time.Parse("15:04:05", other.EndTime)
    
    return (sStart.Before(oEnd) && sEnd.After(oStart))
}
```

**Exam**
```go
type Exam struct {
    ID             string
    ClassID        string
    AssessmentID   string
    ExamDate       time.Time
    StartTime      string
    DurationMinutes int
    Room           string
    Status         ExamStatus
    CreatedAt      time.Time
    UpdatedAt      time.Time
    CreatedBy      string
    UpdatedBy      string
    DeletedAt      *time.Time
}

type ExamStatus string

const (
    ExamStatusScheduled  ExamStatus = "SCHEDULED"
    ExamStatusInProgress ExamStatus = "IN_PROGRESS"
    ExamStatusCompleted  ExamStatus = "COMPLETED"
    ExamStatusCancelled  ExamStatus = "CANCELLED"
)

// Domain Methods
func (e *Exam) Validate() error {
    if e.ExamDate.IsZero() {
        return errors.New("exam date is required")
    }
    if e.DurationMinutes < 1 {
        return errors.New("duration must be at least 1 minute")
    }
    return e.ValidateStatus()
}

func (e *Exam) ValidateStatus() error {
    switch e.Status {
    case ExamStatusScheduled, ExamStatusInProgress,
         ExamStatusCompleted, ExamStatusCancelled:
        return nil
    default:
        return errors.New("invalid exam status")
    }
}

func (e *Exam) CanStart() error {
    if e.Status != ExamStatusScheduled {
        return errors.New("only scheduled exams can be started")
    }
    return nil
}

func (e *Exam) CanComplete() error {
    if e.Status != ExamStatusInProgress {
        return errors.New("only in-progress exams can be completed")
    }
    return nil
}
```

**Assignment**
```go
type Assignment struct {
    ID          string
    ClassID     string
    AssessmentID string
    Title       string
    Description string
    DueDate     time.Time
    MaxScore    int
    Status      AssignmentStatus
    CreatedAt   time.Time
    UpdatedAt   time.Time
    CreatedBy   string
    UpdatedBy   string
    DeletedAt   *time.Time
}

type AssignmentStatus string

const (
    AssignmentStatusAssigned   AssignmentStatus = "ASSIGNED"
    AssignmentStatusInProgress AssignmentStatus = "IN_PROGRESS"
    AssignmentStatusSubmitted  AssignmentStatus = "SUBMITTED"
    AssignmentStatusGraded     AssignmentStatus = "GRADED"
    AssignmentStatusCancelled  AssignmentStatus = "CANCELLED"
)

// Domain Methods
func (a *Assignment) Validate() error {
    if a.Title == "" {
        return errors.New("title is required")
    }
    if a.DueDate.IsZero() {
        return errors.New("due date is required")
    }
    if a.MaxScore < 0 {
        return errors.New("max score cannot be negative")
    }
    return a.ValidateStatus()
}

func (a *Assignment) ValidateStatus() error {
    switch a.Status {
    case AssignmentStatusAssigned, AssignmentStatusInProgress,
         AssignmentStatusSubmitted, AssignmentStatusGraded, AssignmentStatusCancelled:
        return nil
    default:
        return errors.New("invalid assignment status")
    }
}

func (a *Assignment) IsOverdue() bool {
    return time.Now().After(a.DueDate) && a.Status == AssignmentStatusAssigned
}
```

**Notification**
```go
type Notification struct {
    ID         string
    UserID     string
    Title      string
    Message    string
    Type       NotificationType
    IsRead     bool
    ReadAt     *time.Time
    ActionURL  string
    Metadata   map[string]interface{}
    CreatedAt  time.Time
    DeletedAt  *time.Time
}

type NotificationType string

const (
    NotificationTypeInfo    NotificationType = "INFO"
    NotificationTypeWarning NotificationType = "WARNING"
    NotificationTypeError   NotificationType = "ERROR"
    NotificationTypeSuccess NotificationType = "SUCCESS"
)

// Domain Methods
func (n *Notification) Validate() error {
    if n.Title == "" {
        return errors.New("title is required")
    }
    if n.Message == "" {
        return errors.New("message is required")
    }
    return n.ValidateType()
}

func (n *Notification) ValidateType() error {
    switch n.Type {
    case NotificationTypeInfo, NotificationTypeWarning,
         NotificationTypeError, NotificationTypeSuccess:
        return nil
    default:
        return errors.New("invalid notification type")
    }
}

func (n *Notification) MarkAsRead() error {
    if n.IsRead {
        return errors.New("notification already read")
    }
    n.IsRead = true
    now := time.Now()
    n.ReadAt = &now
    return nil
}
```

**Announcement**
```go
type Announcement struct {
    ID            string
    SchoolID      string
    Title         string
    Content       string
    Priority      AnnouncementPriority
    TargetAudience AnnouncementTargetAudience
    PublishedBy   string
    PublishedAt   time.Time
    ExpiresAt     *time.Time
    IsActive      bool
    CreatedAt     time.Time
    UpdatedAt     time.Time
    DeletedAt     *time.Time
}

type AnnouncementPriority string

const (
    AnnouncementPriorityLow    AnnouncementPriority = "LOW"
    AnnouncementPriorityNormal AnnouncementPriority = "NORMAL"
    AnnouncementPriorityHigh   AnnouncementPriority = "HIGH"
    AnnouncementPriorityUrgent AnnouncementPriority = "URGENT"
)

type AnnouncementTargetAudience string

const (
    AnnouncementTargetAudienceAll      AnnouncementTargetAudience = "ALL"
    AnnouncementTargetAudienceTeachers AnnouncementTargetAudience = "TEACHERS"
    AnnouncementTargetAudienceStudents AnnouncementTargetAudience = "STUDENTS"
    AnnouncementTargetAudienceParents  AnnouncementTargetAudience = "PARENTS"
    AnnouncementTargetAudienceAdmin    AnnouncementTargetAudience = "ADMIN"
)

// Domain Methods
func (a *Announcement) Validate() error {
    if a.Title == "" {
        return errors.New("title is required")
    }
    if a.Content == "" {
        return errors.New("content is required")
    }
    if err := a.ValidatePriority(); err != nil {
        return err
    }
    return a.ValidateTargetAudience()
}

func (a *Announcement) ValidatePriority() error {
    switch a.Priority {
    case AnnouncementPriorityLow, AnnouncementPriorityNormal,
         AnnouncementPriorityHigh, AnnouncementPriorityUrgent:
        return nil
    default:
        return errors.New("invalid announcement priority")
    }
}

func (a *Announcement) ValidateTargetAudience() error {
    switch a.TargetAudience {
    case AnnouncementTargetAudienceAll, AnnouncementTargetAudienceTeachers,
         AnnouncementTargetAudienceStudents, AnnouncementTargetAudienceParents,
         AnnouncementTargetAudienceAdmin:
        return nil
    default:
        return errors.New("invalid target audience")
    }
}

func (a *Announcement) IsExpired() bool {
    if a.ExpiresAt == nil {
        return false
    }
    return time.Now().After(*a.ExpiresAt)
}
```

**Message**
```go
type Message struct {
    ID              string
    SenderID        string
    ReceiverID      string
    Subject         string
    Content         string
    IsRead          bool
    ReadAt          *time.Time
    ParentMessageID *string
    CreatedAt       time.Time
    DeletedAt       *time.Time
}

// Domain Methods
func (m *Message) Validate() error {
    if m.Content == "" {
        return errors.New("content is required")
    }
    if m.SenderID == m.ReceiverID {
        return errors.New("sender and receiver cannot be the same")
    }
    return nil
}

func (m *Message) MarkAsRead() error {
    if m.IsRead {
        return errors.New("message already read")
    }
    m.IsRead = true
    now := time.Now()
    m.ReadAt = &now
    return nil
}

func (m *Message) IsReply() bool {
    return m.ParentMessageID != nil
}
```

**AuditLog**
```go
type AuditLog struct {
    ID          string
    UserID      *string
    Action      string
    EntityType  string
    EntityID    *string
    OldValues   map[string]interface{}
    NewValues   map[string]interface{}
    IPAddress   string
    UserAgent   string
    RequestID   string
    CreatedAt   time.Time
}

// Domain Methods
func (a *AuditLog) Validate() error {
    if a.Action == "" {
        return errors.New("action is required")
    }
    if a.EntityType == "" {
        return errors.New("entity type is required")
    }
    return nil
}
```

---

## Event Flows

### Audit Logging Flow

**Trigger:** Any data mutation operation

**Flow:**
```
1. Middleware captures request (user_id, action, entity_type, entity_id)
2. Middleware captures old values (before mutation)
3. Operation executes
4. Middleware captures new values (after mutation)
5. AuditService.Create() called
6. AuditLog persisted to database
```

**Implementation:**
```go
func AuditMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Capture request details
        userID := c.GetString("user_id")
        action := c.Request.Method + " " + c.Request.URL.Path
        requestID := c.GetHeader("X-Request-ID")
        
        // Get old values if applicable
        var oldValues map[string]interface{}
        if c.Request.Method != "POST" {
            oldValues = getOldValues(c)
        }
        
        // Execute handler
        c.Next()
        
        // Get new values
        newValues := getNewValues(c)
        
        // Create audit log
        audit := &domain.AuditLog{
            UserID:     &userID,
            Action:     action,
            EntityType: getEntityType(c),
            EntityID:   getEntityID(c),
            OldValues:  oldValues,
            NewValues:  newValues,
            IPAddress:  c.ClientIP(),
            UserAgent:  c.GetHeader("User-Agent"),
            RequestID:  requestID,
            CreatedAt:  time.Now(),
        }
        
        auditService.Create(c, audit)
    }
}
```

### Notification Flow

**Trigger:** Business events (assignment created, exam scheduled, etc.)

**Flow:**
```
1. Business event occurs
2. NotificationService.Create() called
3. Notification persisted to database
4. Notification queued to RabbitMQ
5. Background worker processes queue
6. WebSocket notification sent to user
7. Email notification sent (if applicable)
```

**Implementation:**
```go
func (s *NotificationService) CreateAndNotify(ctx context.Context, notification *domain.Notification) error {
    // Create notification
    if err := s.notificationRepo.Create(ctx, notification); err != nil {
        return err
    }
    
    // Queue for delivery
    job := &Job{
        Type:    "notification_delivery",
        Payload: notification,
    }
    
    return s.queueService.Enqueue(ctx, job)
}
```

### Background Job Flow

**Trigger:** Async operations (email sending, report generation, etc.)

**Flow:**
```
1. Application Service creates job
2. Job persisted to job_queue table
3. Job published to RabbitMQ
4. Background worker consumes job
5. Worker executes job logic
6. Job status updated to COMPLETED or FAILED
7. Result published to result queue
8. Application processes result
```

**Implementation:**
```go
func (s *QueueService) Enqueue(ctx context.Context, job *Job) error {
    // Persist job
    job.Status = "PENDING"
    if err := s.jobRepo.Create(ctx, job); err != nil {
        return err
    }
    
    // Publish to RabbitMQ
    message, _ := json.Marshal(job)
    return s.rabbitMQ.Publish("jobs", message)
}

func (w *Worker) ProcessJob(job *Job) error {
    // Update status to PROCESSING
    job.Status = "PROCESSING"
    job.StartedAt = time.Now()
    w.jobRepo.Update(context.Background(), job)
    
    // Execute job
    var err error
    switch job.Type {
    case "notification_delivery":
        err = w.processNotificationDelivery(job)
    case "report_generation":
        err = w.processReportGeneration(job)
    case "email_sending":
        err = w.processEmailSending(job)
    }
    
    // Update status
    if err != nil {
        job.Status = "FAILED"
        job.ErrorMessage = err.Error()
        job.RetryCount++
    } else {
        job.Status = "COMPLETED"
    }
    job.CompletedAt = time.Now()
    
    return w.jobRepo.Update(context.Background(), job)
}
```

---

## Background Jobs Design

### Job Queue

**Table:** job_queue

**Job Types:**
- `notification_delivery` - Send notifications via WebSocket/email
- `report_generation` - Generate narrative reports
- `email_sending` - Send emails
- `data_export` - Export data to CSV/PDF
- `ai_processing` - AI-assisted content generation

**Job Priority:**
- 1: Highest (urgent notifications)
- 3: High (email sending)
- 5: Normal (report generation)
- 7: Low (data export)
- 9: Lowest (AI processing)

**Retry Strategy:**
- Max retries: 3
- Retry delay: Exponential backoff (1s, 2s, 4s)
- Dead letter queue for failed jobs after max retries

### Worker Design

**Worker Pool:**
- Number of workers: 5 (configurable)
- Worker type: Dedicated workers per job type
- Concurrency: 1 job per worker

**Worker Lifecycle:**
1. Worker starts
2. Connects to RabbitMQ
3. Subscribes to job queue
4. Polls for jobs
5. Processes job
6. Updates job status
7. Acknowledges/rejects message
8. Repeat

**Implementation:**
```go
type Worker struct {
    jobRepo     repository.JobRepository
    rabbitMQ    *RabbitMQClient
    jobType     string
    concurrency int
}

func (w *Worker) Start() error {
    // Connect to RabbitMQ
    if err := w.rabbitMQ.Connect(); err != nil {
        return err
    }
    
    // Start worker goroutines
    for i := 0; i < w.concurrency; i++ {
        go w.processJobs()
    }
    
    return nil
}

func (w *Worker) processJobs() {
    for {
        // Get job from queue
        job, err := w.rabbitMQ.Consume(w.jobType)
        if err != nil {
            log.Error("Failed to consume job", err)
            continue
        }
        
        // Process job
        if err := w.ProcessJob(job); err != nil {
            log.Error("Failed to process job", err)
            w.rabbitMQ.Reject(job, true) // Requeue for retry
        } else {
            w.rabbitMQ.Ack(job)
        }
    }
}
```

---

## Queue Design

### RabbitMQ Configuration

**Exchanges:**
- `nusa.direct` - Direct exchange for job routing
- `nusa.fanout` - Fanout exchange for notifications
- `nusa.topic` - Topic exchange for event routing

**Queues:**
- `jobs.high` - High priority jobs
- `jobs.normal` - Normal priority jobs
- `jobs.low` - Low priority jobs
- `notifications` - Notification delivery
- `emails` - Email sending

**Bindings:**
- `jobs.high` → `nusa.direct` (routing key: `job.high`)
- `jobs.normal` → `nusa.direct` (routing key: `job.normal`)
- `jobs.low` → `nusa.direct` (routing key: `job.low`)
- `notifications` → `nusa.fanout`
- `emails` → `nusa.direct` (routing key: `email`)

**Message Format:**
```json
{
  "job_id": "uuid",
  "type": "notification_delivery",
  "priority": 5,
  "payload": { ... },
  "created_at": "2026-06-13T10:00:00Z"
}
```

---

## Caching Strategy

### Redis Configuration

**Cache Keys:**
- Namespace-based: `{module}:{entity}:{id}`
- Examples: `academic_year:uuid`, `class:uuid`, `user:uuid`

**Cache TTL:**
- User data: 1 hour
- Academic foundation data: 30 minutes
- Curriculum data: 1 hour
- Class data: 30 minutes
- Attendance data: 15 minutes
- Notification data: 5 minutes

**Cache Invalidation:**
- Event-based: Invalidate on mutation
- Time-based: TTL expiration
- Manual: Explicit invalidation

**Cache Implementation:**
```go
type CacheService struct {
    redis *redis.Client
}

func (s *CacheService) Get(ctx context.Context, key string, dest interface{}) error {
    val, err := s.redis.Get(ctx, key).Result()
    if err != nil {
        return err
    }
    return json.Unmarshal([]byte(val), dest)
}

func (s *CacheService) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
    val, err := json.Marshal(value)
    if err != nil {
        return err
    }
    return s.redis.Set(ctx, key, val, ttl).Err()
}

func (s *CacheService) Invalidate(ctx context.Context, pattern string) error {
    iter := s.redis.Scan(ctx, 0, pattern, 0).Iterator()
    for iter.Next(ctx) {
        s.redis.Del(ctx, iter.Val())
    }
    return iter.Err()
}
```

**Cache Middleware:**
```go
func CacheMiddleware(cacheService *CacheService, ttl time.Duration) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Try cache
        key := fmt.Sprintf("%s:%s", c.Request.URL.Path, c.Request.URL.RawQuery)
        var result interface{}
        
        if err := cacheService.Get(c, key, &result); err == nil {
            c.JSON(200, result)
            c.Abort()
            return
        }
        
        // Execute handler
        c.Next()
        
        // Cache response
        if c.Writer.Status() == 200 {
            cacheService.Set(c, key, c.Keys["response"], ttl)
        }
    }
}
```

---

## Transaction Management

### Transaction Strategy

**Service Layer Transactions:**
- Use database transactions for multi-step operations
- Rollback on error
- Commit on success

**Implementation:**
```go
func (s *ApplicationService) CreateClass(ctx context.Context, req CreateClassRequest) (*Class, error) {
    tx, err := s.db.Begin(ctx)
    if err != nil {
        return nil, err
    }
    defer tx.Rollback(ctx)
    
    // Create class
    class := &domain.Class{...}
    if err := s.classRepo.Create(ctx, class); err != nil {
        return nil, err
    }
    
    // Create default schedule
    schedule := &domain.Schedule{...}
    if err := s.scheduleRepo.Create(ctx, schedule); err != nil {
        return nil, err
    }
    
    // Commit transaction
    if err := tx.Commit(ctx); err != nil {
        return nil, err
    }
    
    return class, nil
}
```

---

## Error Handling

### Error Types

**Domain Errors:**
- ValidationError - Input validation failed
- BusinessRuleError - Business rule violated
- NotFoundError - Resource not found
- ConflictError - Resource conflict

**Infrastructure Errors:**
- DatabaseError - Database operation failed
- CacheError - Cache operation failed
- QueueError - Queue operation failed

**Implementation:**
```go
type AppError struct {
    Code    string
    Message string
    Cause   error
}

func (e *AppError) Error() string {
    return e.Message
}

func NewValidationError(message string) *AppError {
    return &AppError{
        Code:    "VAL_001",
        Message: message,
    }
}

func NewBusinessRuleError(message string) *AppError {
    return &AppError{
        Code:    "BIZ_001",
        Message: message,
    }
}
```

---

## Logging Strategy

### Structured Logging

**Log Format:** JSON

**Log Levels:**
- DEBUG - Detailed debugging information
- INFO - General informational messages
- WARN - Warning messages
- ERROR - Error messages
- FATAL - Critical errors

**Log Fields:**
- timestamp
- level
- request_id
- user_id
- action
- entity_type
- entity_id
- error
- duration_ms

**Implementation:**
```go
type Logger struct {
    *logrus.Logger
}

func (l *Logger) WithRequestID(requestID string) *logrus.Entry {
    return l.WithField("request_id", requestID)
}

func (l *Logger) WithUserID(userID string) *logrus.Entry {
    return l.WithField("user_id", userID)
}

func (l *Logger) WithAction(action string) *logrus.Entry {
    return l.WithField("action", action)
}
```

---

## Conclusion

The backend design maintains the existing layered architecture with DDD Lite principles while adding new services, repositories, and domain entities for class management, attendance, scheduling, exams, assignments, and communication features. The design includes comprehensive background job processing with RabbitMQ, caching strategy with Redis, and structured logging for observability.

The design prioritizes maintainability, scalability, simplicity, testability, security, and observability without introducing overengineering or forbidden patterns (CQRS, Event Sourcing).

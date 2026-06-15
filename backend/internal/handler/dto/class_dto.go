package dto

import (
	"encoding/json"
	"time"
)

// CreateClassRequest represents the request to create a class
type CreateClassRequest struct {
	SchoolID       string  `json:"school_id" binding:"required"`
	AcademicYearID string  `json:"academic_year_id" binding:"required"`
	SemesterID     string  `json:"semester_id" binding:"required"`
	SubjectID      string  `json:"subject_id" binding:"required"`
	TeacherID      string  `json:"teacher_id" binding:"required"`
	Name           string  `json:"name" binding:"required,min=2,max=255"`
	GradeLevel     string  `json:"grade_level" binding:"required,max=50"`
	Room           *string `json:"room,omitempty" binding:"omitempty,max=100"`
	MaxStudents    int     `json:"max_students" binding:"required,min=1,max=100"`
}

// UpdateClassRequest represents the request to update a class
type UpdateClassRequest struct {
	Name        *string `json:"name,omitempty" binding:"omitempty,min=2,max=255"`
	GradeLevel  *string `json:"grade_level,omitempty" binding:"omitempty,max=50"`
	Room        *string `json:"room,omitempty" binding:"omitempty,max=100"`
	MaxStudents *int    `json:"max_students,omitempty" binding:"omitempty,min=1,max=100"`
	IsActive    *bool   `json:"is_active,omitempty"`
}

// ClassResponse represents the class response
type ClassResponse struct {
	ID              string  `json:"id"`
	SchoolID        string  `json:"school_id"`
	SchoolName      *string `json:"school_name,omitempty"`
	AcademicYearID  string  `json:"academic_year_id"`
	AcademicYear    *string `json:"academic_year,omitempty"`
	SemesterID      string  `json:"semester_id"`
	SemesterName    *string `json:"semester_name,omitempty"`
	SubjectID       string  `json:"subject_id"`
	SubjectName     *string `json:"subject_name,omitempty"`
	TeacherID       string  `json:"teacher_id"`
	TeacherName     *string `json:"teacher_name,omitempty"`
	Name            string  `json:"name"`
	GradeLevel      string  `json:"grade_level"`
	Room            *string `json:"room,omitempty"`
	MaxStudents     int     `json:"max_students"`
	CurrentStudents int     `json:"current_students"`
	IsActive        bool    `json:"is_active"`
}

// ClassListResponse represents the paginated class list response
type ClassListResponse struct {
	Classes  []*ClassResponse `json:"classes"`
	Total    int              `json:"total"`
	Page     int              `json:"page"`
	PageSize int              `json:"page_size"`
}

// CreateClassEnrollmentRequest represents the request to enroll a student
type CreateClassEnrollmentRequest struct {
	ClassID   string  `json:"class_id" binding:"required"`
	StudentID string  `json:"student_id" binding:"required"`
	Notes     *string `json:"notes,omitempty"`
}

// UpdateClassEnrollmentRequest represents the request to update enrollment
type UpdateClassEnrollmentRequest struct {
	Status EnrollmentStatus `json:"status" binding:"required,oneof=ACTIVE INACTIVE WITHDRAWN COMPLETED"`
	Notes  *string          `json:"notes,omitempty"`
}

// EnrollmentStatus represents enrollment status
type EnrollmentStatus string

const (
	EnrollmentStatusActive    EnrollmentStatus = "ACTIVE"
	EnrollmentStatusInactive  EnrollmentStatus = "INACTIVE"
	EnrollmentStatusWithdrawn EnrollmentStatus = "WITHDRAWN"
	EnrollmentStatusCompleted EnrollmentStatus = "COMPLETED"
)

// AttendanceStatus represents attendance status
type AttendanceStatus string

const (
	AttendanceStatusPresent AttendanceStatus = "PRESENT"
	AttendanceStatusAbsent  AttendanceStatus = "ABSENT"
	AttendanceStatusLate    AttendanceStatus = "LATE"
	AttendanceStatusExcused AttendanceStatus = "EXCUSED"
)

// ClassEnrollmentResponse represents the enrollment response
type ClassEnrollmentResponse struct {
	ID             string           `json:"id"`
	ClassID        string           `json:"class_id"`
	ClassName      *string          `json:"class_name,omitempty"`
	StudentID      string           `json:"student_id"`
	StudentName    *string          `json:"student_name,omitempty"`
	EnrollmentDate string           `json:"enrollment_date"`
	Status         EnrollmentStatus `json:"status"`
	Notes          *string          `json:"notes,omitempty"`
	CreatedAt      string           `json:"created_at"`
	UpdatedAt      string           `json:"updated_at"`
}

// ClassEnrollmentListResponse represents the paginated enrollment list response
type ClassEnrollmentListResponse struct {
	Enrollments []*ClassEnrollmentResponse `json:"enrollments"`
	Total       int                        `json:"total"`
	Page        int                        `json:"page"`
	PageSize    int                        `json:"page_size"`
}

// CreateAttendanceRequest represents the request to record attendance
type CreateAttendanceRequest struct {
	ClassID   string           `json:"class_id" binding:"required"`
	StudentID string           `json:"student_id" binding:"required"`
	Date      string           `json:"date" binding:"required"`
	Status    AttendanceStatus `json:"status" binding:"required,oneof=PRESENT ABSENT LATE EXCUSED"`
	Notes     *string          `json:"notes,omitempty"`
}

// UpdateAttendanceRequest represents the request to update attendance
type UpdateAttendanceRequest struct {
	Status AttendanceStatus `json:"status" binding:"required,oneof=PRESENT ABSENT LATE EXCUSED"`
	Notes  *string          `json:"notes,omitempty"`
}

// AttendanceResponse represents the attendance response
type AttendanceResponse struct {
	ID             string           `json:"id"`
	ClassID        string           `json:"class_id"`
	ClassName      *string          `json:"class_name,omitempty"`
	StudentID      string           `json:"student_id"`
	StudentName    *string          `json:"student_name,omitempty"`
	Date           string           `json:"date"`
	Status         AttendanceStatus `json:"status"`
	Notes          *string          `json:"notes,omitempty"`
	RecordedBy     string           `json:"recorded_by"`
	RecordedByName *string          `json:"recorded_by_name,omitempty"`
}

// AttendanceListResponse represents the paginated attendance list response
type AttendanceListResponse struct {
	Attendances []*AttendanceResponse `json:"attendances"`
	Total       int                   `json:"total"`
	Page        int                   `json:"page"`
	PageSize    int                   `json:"page_size"`
}

// AttendanceStatsResponse represents attendance statistics
type AttendanceStatsResponse struct {
	Present int `json:"present"`
	Absent  int `json:"absent"`
	Late    int `json:"late"`
	Excused int `json:"excused"`
	Total   int `json:"total"`
}

// CreateScheduleRequest represents the request to create a schedule
type CreateScheduleRequest struct {
	ClassID   string  `json:"class_id" binding:"required"`
	DayOfWeek int     `json:"day_of_week" binding:"required,min=1,max=7"`
	StartTime string  `json:"start_time" binding:"required"`
	EndTime   string  `json:"end_time" binding:"required"`
	Room      *string `json:"room,omitempty" binding:"omitempty,max=100"`
}

// UpdateScheduleRequest represents the request to update a schedule
type UpdateScheduleRequest struct {
	DayOfWeek *int    `json:"day_of_week,omitempty" binding:"omitempty,min=1,max=7"`
	StartTime *string `json:"start_time,omitempty" binding:"omitempty"`
	EndTime   *string `json:"end_time,omitempty" binding:"omitempty"`
	Room      *string `json:"room,omitempty" binding:"omitempty,max=100"`
	IsActive  *bool   `json:"is_active,omitempty"`
}

// ScheduleResponse represents the schedule response
type ScheduleResponse struct {
	ID        string  `json:"id"`
	ClassID   string  `json:"class_id"`
	ClassName *string `json:"class_name,omitempty"`
	DayOfWeek int     `json:"day_of_week"`
	DayName   string  `json:"day_name"`
	StartTime string  `json:"start_time"`
	EndTime   string  `json:"end_time"`
	Room      *string `json:"room,omitempty"`
	IsActive  bool    `json:"is_active"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

// ScheduleListResponse represents the paginated schedule list response
type ScheduleListResponse struct {
	Schedules []*ScheduleResponse `json:"schedules"`
	Total     int                 `json:"total"`
	Page      int                 `json:"page"`
	PageSize  int                 `json:"page_size"`
}

// NotificationType represents the type of notification
type NotificationType string

const (
	NotificationTypeInfo    NotificationType = "INFO"
	NotificationTypeWarning NotificationType = "WARNING"
	NotificationTypeError   NotificationType = "ERROR"
	NotificationTypeSuccess NotificationType = "SUCCESS"
)

// CreateNotificationRequest represents the request to create a notification
type CreateNotificationRequest struct {
	UserID    string           `json:"user_id" binding:"required"`
	Title     string           `json:"title" binding:"required,max=255"`
	Message   string           `json:"message" binding:"required"`
	Type      NotificationType `json:"type" binding:"required,oneof=INFO WARNING ERROR SUCCESS"`
	ActionURL *string          `json:"action_url,omitempty" binding:"omitempty,max=500"`
	Metadata  json.RawMessage  `json:"metadata,omitempty"`
}

// UpdateNotificationRequest represents the request to update a notification
type UpdateNotificationRequest struct {
	IsRead *bool `json:"is_read,omitempty"`
}

// NotificationResponse represents the notification response
type NotificationResponse struct {
	ID        string           `json:"id"`
	UserID    string           `json:"user_id"`
	UserName  *string          `json:"user_name,omitempty"`
	Title     string           `json:"title"`
	Message   string           `json:"message"`
	Type      NotificationType `json:"type"`
	IsRead    bool             `json:"is_read"`
	ReadAt    *string          `json:"read_at,omitempty"`
	ActionURL *string          `json:"action_url,omitempty"`
	Metadata  json.RawMessage  `json:"metadata,omitempty"`
	CreatedAt string           `json:"created_at"`
}

// NotificationListResponse represents the paginated notification list response
type NotificationListResponse struct {
	Notifications []*NotificationResponse `json:"notifications"`
	Total         int                     `json:"total"`
	Page          int                     `json:"page"`
	PageSize      int                     `json:"page_size"`
}

// AnnouncementPriority represents the priority level of an announcement
type AnnouncementPriority string

const (
	AnnouncementPriorityLow    AnnouncementPriority = "LOW"
	AnnouncementPriorityNormal AnnouncementPriority = "NORMAL"
	AnnouncementPriorityHigh   AnnouncementPriority = "HIGH"
	AnnouncementPriorityUrgent AnnouncementPriority = "URGENT"
)

// TargetAudience represents the target audience for an announcement
type TargetAudience string

const (
	TargetAudienceAll      TargetAudience = "ALL"
	TargetAudienceTeachers TargetAudience = "TEACHERS"
	TargetAudienceStudents TargetAudience = "STUDENTS"
	TargetAudienceParents  TargetAudience = "PARENTS"
	TargetAudienceAdmin    TargetAudience = "ADMIN"
)

// CreateAnnouncementRequest represents the request to create an announcement
type CreateAnnouncementRequest struct {
	SchoolID       string               `json:"school_id" binding:"required"`
	Title          string               `json:"title" binding:"required,max=255"`
	Content        string               `json:"content" binding:"required"`
	Priority       AnnouncementPriority `json:"priority" binding:"required,oneof=LOW NORMAL HIGH URGENT"`
	TargetAudience TargetAudience       `json:"target_audience" binding:"required,oneof=ALL TEACHERS STUDENTS PARENTS ADMIN"`
	ExpiresAt      *time.Time           `json:"expires_at,omitempty"`
}

// UpdateAnnouncementRequest represents the request to update an announcement
type UpdateAnnouncementRequest struct {
	Title          *string               `json:"title,omitempty" binding:"omitempty,max=255"`
	Content        *string               `json:"content,omitempty"`
	Priority       *AnnouncementPriority `json:"priority,omitempty" binding:"omitempty,oneof=LOW NORMAL HIGH URGENT"`
	TargetAudience *TargetAudience       `json:"target_audience,omitempty" binding:"omitempty,oneof=ALL TEACHERS STUDENTS PARENTS ADMIN"`
	ExpiresAt      *time.Time            `json:"expires_at,omitempty"`
	IsActive       *bool                 `json:"is_active,omitempty"`
}

// AnnouncementResponse represents the announcement response
type AnnouncementResponse struct {
	ID              string               `json:"id"`
	SchoolID        string               `json:"school_id"`
	SchoolName      *string              `json:"school_name,omitempty"`
	Title           string               `json:"title"`
	Content         string               `json:"content"`
	Priority        AnnouncementPriority `json:"priority"`
	TargetAudience  TargetAudience       `json:"target_audience"`
	PublishedBy     string               `json:"published_by"`
	PublishedByName *string              `json:"published_by_name,omitempty"`
	PublishedAt     string               `json:"published_at"`
	ExpiresAt       *string              `json:"expires_at,omitempty"`
	IsActive        bool                 `json:"is_active"`
	CreatedAt       string               `json:"created_at"`
	UpdatedAt       string               `json:"updated_at"`
}

// AnnouncementListResponse represents the paginated announcement list response
type AnnouncementListResponse struct {
	Announcements []*AnnouncementResponse `json:"announcements"`
	Total         int                     `json:"total"`
	Page          int                     `json:"page"`
	PageSize      int                     `json:"page_size"`
}

// CreateMessageRequest represents the request to create a message
type CreateMessageRequest struct {
	SenderID        string  `json:"sender_id" binding:"required"`
	ReceiverID      string  `json:"receiver_id" binding:"required"`
	Subject         *string `json:"subject,omitempty" binding:"omitempty,max=255"`
	Content         string  `json:"content" binding:"required"`
	ParentMessageID *string `json:"parent_message_id,omitempty"`
}

// UpdateMessageRequest represents the request to update a message
type UpdateMessageRequest struct {
	IsRead *bool `json:"is_read,omitempty"`
}

// MessageResponse represents the message response
type MessageResponse struct {
	ID              string  `json:"id"`
	SenderID        string  `json:"sender_id"`
	SenderName      *string `json:"sender_name,omitempty"`
	ReceiverID      string  `json:"receiver_id"`
	ReceiverName    *string `json:"receiver_name,omitempty"`
	Subject         *string `json:"subject,omitempty"`
	Content         string  `json:"content"`
	IsRead          bool    `json:"is_read"`
	ReadAt          *string `json:"read_at,omitempty"`
	ParentMessageID *string `json:"parent_message_id,omitempty"`
	CreatedAt       string  `json:"created_at"`
}

// MessageListResponse represents the paginated message list response
type MessageListResponse struct {
	Messages []*MessageResponse `json:"messages"`
	Total    int                `json:"total"`
	Page     int                `json:"page"`
	PageSize int                `json:"page_size"`
}

// ExamStatus represents the status of an exam
type ExamStatus string

const (
	ExamStatusScheduled  ExamStatus = "SCHEDULED"
	ExamStatusInProgress ExamStatus = "IN_PROGRESS"
	ExamStatusCompleted  ExamStatus = "COMPLETED"
	ExamStatusCancelled  ExamStatus = "CANCELLED"
)

// CreateExamRequest represents the request to create an exam
type CreateExamRequest struct {
	ClassID         string    `json:"class_id" binding:"required"`
	AssessmentID    string    `json:"assessment_id" binding:"required"`
	ExamDate        time.Time `json:"exam_date" binding:"required"`
	StartTime       string    `json:"start_time" binding:"required"`
	DurationMinutes int       `json:"duration_minutes" binding:"required,min=1"`
	Room            *string   `json:"room,omitempty" binding:"omitempty,max=100"`
}

// UpdateExamRequest represents the request to update an exam
type UpdateExamRequest struct {
	ExamDate        *time.Time  `json:"exam_date,omitempty"`
	StartTime       *string     `json:"start_time,omitempty" binding:"omitempty"`
	DurationMinutes *int        `json:"duration_minutes,omitempty" binding:"omitempty,min=1"`
	Room            *string     `json:"room,omitempty" binding:"omitempty,max=100"`
	Status          *ExamStatus `json:"status,omitempty" binding:"omitempty,oneof=SCHEDULED IN_PROGRESS COMPLETED CANCELLED"`
}

// ExamResponse represents the exam response
type ExamResponse struct {
	ID              string     `json:"id"`
	ClassID         string     `json:"class_id"`
	ClassName       *string    `json:"class_name,omitempty"`
	AssessmentID    string     `json:"assessment_id"`
	AssessmentType  *string    `json:"assessment_type,omitempty"`
	ExamDate        string     `json:"exam_date"`
	StartTime       string     `json:"start_time"`
	DurationMinutes int        `json:"duration_minutes"`
	Room            *string    `json:"room,omitempty"`
	Status          ExamStatus `json:"status"`
	CreatedAt       string     `json:"created_at"`
	UpdatedAt       string     `json:"updated_at"`
	CreatedBy       *string    `json:"created_by,omitempty"`
	CreatedByName   *string    `json:"created_by_name,omitempty"`
	UpdatedBy       *string    `json:"updated_by,omitempty"`
	UpdatedByName   *string    `json:"updated_by_name,omitempty"`
}

// ExamListResponse represents the paginated exam list response
type ExamListResponse struct {
	Exams    []*ExamResponse `json:"exams"`
	Total    int             `json:"total"`
	Page     int             `json:"page"`
	PageSize int             `json:"page_size"`
}

// AssignmentStatus represents the status of an assignment
type AssignmentStatus string

const (
	AssignmentStatusAssigned   AssignmentStatus = "ASSIGNED"
	AssignmentStatusInProgress AssignmentStatus = "IN_PROGRESS"
	AssignmentStatusSubmitted  AssignmentStatus = "SUBMITTED"
	AssignmentStatusGraded     AssignmentStatus = "GRADED"
	AssignmentStatusCancelled  AssignmentStatus = "CANCELLED"
)

// CreateAssignmentRequest represents the request to create an assignment
type CreateAssignmentRequest struct {
	ClassID      string    `json:"class_id" binding:"required"`
	AssessmentID string    `json:"assessment_id" binding:"required"`
	Title        string    `json:"title" binding:"required,max=255"`
	Description  *string   `json:"description,omitempty"`
	DueDate      time.Time `json:"due_date" binding:"required"`
	MaxScore     int       `json:"max_score" binding:"required,min=1"`
}

// UpdateAssignmentRequest represents the request to update an assignment
type UpdateAssignmentRequest struct {
	Title       *string           `json:"title,omitempty" binding:"omitempty,max=255"`
	Description *string           `json:"description,omitempty"`
	DueDate     *time.Time        `json:"due_date,omitempty"`
	MaxScore    *int              `json:"max_score,omitempty" binding:"omitempty,min=1"`
	Status      *AssignmentStatus `json:"status,omitempty" binding:"omitempty,oneof=ASSIGNED IN_PROGRESS SUBMITTED GRADED CANCELLED"`
}

// AssignmentResponse represents the assignment response
type AssignmentResponse struct {
	ID             string           `json:"id"`
	ClassID        string           `json:"class_id"`
	ClassName      *string          `json:"class_name,omitempty"`
	AssessmentID   string           `json:"assessment_id"`
	AssessmentType *string          `json:"assessment_type,omitempty"`
	Title          string           `json:"title"`
	Description    *string          `json:"description,omitempty"`
	DueDate        string           `json:"due_date"`
	MaxScore       int              `json:"max_score"`
	Status         AssignmentStatus `json:"status"`
	CreatedAt      string           `json:"created_at"`
	UpdatedAt      string           `json:"updated_at"`
	CreatedBy      *string          `json:"created_by,omitempty"`
	CreatedByName  *string          `json:"created_by_name,omitempty"`
	UpdatedBy      *string          `json:"updated_by,omitempty"`
	UpdatedByName  *string          `json:"updated_by_name,omitempty"`
}

// AssignmentListResponse represents the paginated assignment list response
type AssignmentListResponse struct {
	Assignments []*AssignmentResponse `json:"assignments"`
	Total       int                   `json:"total"`
	Page        int                   `json:"page"`
	PageSize    int                   `json:"page_size"`
}

// CreateExamResultRequest represents the request to create an exam result
type CreateExamResultRequest struct {
	ExamID    string   `json:"exam_id" binding:"required"`
	StudentID string   `json:"student_id" binding:"required"`
	Score     *float64 `json:"score,omitempty" binding:"omitempty,gte=0,lte=100"`
	Grade     *string  `json:"grade,omitempty" binding:"omitempty,max=10"`
	Remarks   *string  `json:"remarks,omitempty"`
}

// UpdateExamResultRequest represents the request to update an exam result
type UpdateExamResultRequest struct {
	Score    *float64 `json:"score,omitempty" binding:"omitempty,gte=0,lte=100"`
	Grade    *string  `json:"grade,omitempty" binding:"omitempty,max=10"`
	Remarks  *string  `json:"remarks,omitempty"`
	GradedBy *string  `json:"graded_by,omitempty"`
}

// ExamResultResponse represents the exam result response
type ExamResultResponse struct {
	ID           string   `json:"id"`
	ExamID       string   `json:"exam_id"`
	ExamDate     *string  `json:"exam_date,omitempty"`
	ExamTitle    *string  `json:"exam_title,omitempty"`
	StudentID    string   `json:"student_id"`
	StudentName  *string  `json:"student_name,omitempty"`
	Score        *float64 `json:"score,omitempty"`
	Grade        *string  `json:"grade,omitempty"`
	Remarks      *string  `json:"remarks,omitempty"`
	GradedAt     *string  `json:"graded_at,omitempty"`
	GradedBy     *string  `json:"graded_by,omitempty"`
	GradedByName *string  `json:"graded_by_name,omitempty"`
	CreatedAt    string   `json:"created_at"`
	UpdatedAt    string   `json:"updated_at"`
}

// ExamResultListResponse represents the paginated exam result list response
type ExamResultListResponse struct {
	ExamResults []*ExamResultResponse `json:"exam_results"`
	Total       int                   `json:"total"`
	Page        int                   `json:"page"`
	PageSize    int                   `json:"page_size"`
}

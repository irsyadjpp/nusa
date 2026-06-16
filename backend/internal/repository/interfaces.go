package repository

import (
	"context"
	"time"

	"github.com/nusa/backend/internal/domain"
)

// UserRepositoryInterface defines the interface for user data access
type UserRepositoryInterface interface {
	Create(ctx context.Context, user *domain.User) error
	GetByID(ctx context.Context, id string) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) error
	UpdateStatus(ctx context.Context, id string, isActive bool, lockedUntil *string, failedAttempts int) error
	List(ctx context.Context, schoolID, roleID *string, isActive *bool, limit, offset int) ([]*domain.User, error)
	Count(ctx context.Context, schoolID, roleID *string, isActive *bool) (int, error)
	Delete(ctx context.Context, id string) error
	GetUserSchoolID(ctx context.Context, userID string) (*string, error)
	ListUsersBySchool(ctx context.Context, schoolID string, limit, offset int) ([]*domain.User, error)
}

// RoleRepositoryInterface defines the interface for role data access
type RoleRepositoryInterface interface {
	Create(ctx context.Context, role *domain.Role) error
	GetByID(ctx context.Context, id string) (*domain.Role, error)
	GetByName(ctx context.Context, name string) (*domain.Role, error)
	Update(ctx context.Context, role *domain.Role) error
	List(ctx context.Context, isActive *bool) ([]*domain.Role, error)
	AddPermission(ctx context.Context, roleID, resource, action string) error
	GetPermissions(ctx context.Context, roleID string) ([]*domain.Permission, error)
	RemovePermission(ctx context.Context, roleID, resource, action string) error
}

// AssessmentRepository defines the interface for assessment data access
type AssessmentRepositoryInterface interface {
	// Assessment operations
	CreateAssessment(ctx context.Context, assessment *domain.Assessment) error
	GetAssessmentByID(ctx context.Context, id string) (*domain.Assessment, error)
	ListAssessments(ctx context.Context, tpID, userID *string, assessmentType *domain.AssessmentType, status *domain.WorkflowStatus, limit, offset int) ([]*domain.Assessment, error)
	UpdateAssessment(ctx context.Context, assessment *domain.Assessment) error
	DeleteAssessment(ctx context.Context, id string) error

	// Rubric operations
	CreateRubric(ctx context.Context, rubric *domain.Rubric) error
	GetRubricByID(ctx context.Context, id string) (*domain.Rubric, error)
	ListRubrics(ctx context.Context, assessmentID, userID *string, rubricType *domain.RubricType, status *domain.WorkflowStatus, limit, offset int) ([]*domain.Rubric, error)
	UpdateRubric(ctx context.Context, rubric *domain.Rubric) error
	DeleteRubric(ctx context.Context, id string) error

	// Evidence operations
	CreateEvidence(ctx context.Context, evidence *domain.Evidence) error
	GetEvidenceByID(ctx context.Context, id string) (*domain.Evidence, error)
	ListEvidences(ctx context.Context, studentID, assessmentID *string, evidenceType *domain.EvidenceType, status *domain.EvidenceStatus, limit, offset int) ([]*domain.Evidence, error)
	UpdateEvidence(ctx context.Context, evidence *domain.Evidence) error
	DeleteEvidence(ctx context.Context, id string) error

	// Evaluation operations
	CreateEvaluation(ctx context.Context, evaluation *domain.Evaluation) error
	GetEvaluationByID(ctx context.Context, id string) (*domain.Evaluation, error)
	ListEvaluations(ctx context.Context, evidenceID, studentID *string, performanceLevel *domain.PerformanceLevel, limit, offset int) ([]*domain.Evaluation, error)
	UpdateEvaluation(ctx context.Context, evaluation *domain.Evaluation) error
	DeleteEvaluation(ctx context.Context, id string) error
	GetEvaluationHistory(ctx context.Context, evaluationID string) ([]*domain.Evaluation, error)

	// Feedback history operations
	CreateFeedbackHistory(ctx context.Context, feedback *domain.EvaluationFeedbackHistory) error
	GetFeedbackHistory(ctx context.Context, evaluationID string) ([]*domain.EvaluationFeedbackHistory, error)

	// Archive operations
	ArchiveCurrentRevision(ctx context.Context, id string) error

	// Status operations
	UpdateAssessmentStatus(ctx context.Context, id string, status domain.WorkflowStatus, approvedBy *string) error
}

// TPRepositoryInterface defines the interface for TP (Tujuan Pembelajaran) data access
type TPRepositoryInterface interface {
	// TP Set operations
	CreateTPSet(ctx context.Context, tpSet *domain.TPSet) error
	GetTPSetByID(ctx context.Context, id string) (*domain.TPSet, error)
	GetTPSetByCPAndVersion(ctx context.Context, cpID string, versionNo int) (*domain.TPSet, error)
	ListTPSets(ctx context.Context, cpID *string, status *domain.WorkflowStatus, schoolID *string, limit, offset int) ([]*domain.TPSet, error)
	UpdateTPSet(ctx context.Context, tpSet *domain.TPSet) error
	UpdateTPSetStatus(ctx context.Context, id string, status domain.WorkflowStatus, approvedBy *string, rejectedReason *string) error
	DeleteTPSet(ctx context.Context, id string) error

	// TP operations
	CreateTP(ctx context.Context, tp *domain.TP) error
	GetTPByID(ctx context.Context, id string) (*domain.TP, error)
	ListTPs(ctx context.Context, tpSetID, subjectID, phaseID, status *string, limit, offset int) ([]*domain.TP, error)
	UpdateTP(ctx context.Context, tp *domain.TP) error
	DeleteTP(ctx context.Context, id string) error
	GetTPVersionHistory(ctx context.Context, tpSetID string) ([]*domain.TP, error)
	HasDownstreamAssessments(ctx context.Context, tpID string) (bool, error)
}

// LearningPlanningRepository defines the interface for ATP and Modul Ajar data access
type LearningPlanningRepositoryInterface interface {
	// ATP Set operations
	CreateATPSet(ctx context.Context, atpSet *domain.ATPSet) error
	GetATPSetByID(ctx context.Context, id string) (*domain.ATPSet, error)
	GetATPSetByTPAndVersion(ctx context.Context, tpSetID string, versionNo int) (*domain.ATPSet, error)
	ListATPSets(ctx context.Context, tpSetID *string, status *domain.WorkflowStatus, limit, offset int) ([]*domain.ATPSet, error)
	UpdateATPSet(ctx context.Context, atpSet *domain.ATPSet) error
	DeleteATPSet(ctx context.Context, id string) error

	// ATP operations
	CreateATP(ctx context.Context, atp *domain.ATP) error
	GetATPByID(ctx context.Context, id string) (*domain.ATP, error)
	ListATPs(ctx context.Context, atpSetID *string, limit, offset int) ([]*domain.ATP, error)
	UpdateATP(ctx context.Context, atp *domain.ATP) error
	DeleteATP(ctx context.Context, id string) error

	// Modul Ajar Set operations
	CreateModulAjarSet(ctx context.Context, set *domain.ModulAjarSet) error
	GetModulAjarSetByID(ctx context.Context, id string) (*domain.ModulAjarSet, error)
	ListModulAjarSets(ctx context.Context, atpSetID *string, status *domain.WorkflowStatus, limit, offset int) ([]*domain.ModulAjarSet, error)
	UpdateModulAjarSet(ctx context.Context, set *domain.ModulAjarSet) error
	DeleteModulAjarSet(ctx context.Context, id string) error

	// Modul Ajar operations
	CreateModulAjar(ctx context.Context, modulAjar *domain.ModulAjar) error
	GetModulAjarByID(ctx context.Context, id string) (*domain.ModulAjar, error)
	ListModulAjars(ctx context.Context, modulAjarSetID, atpID *string, limit, offset int) ([]*domain.ModulAjar, error)
	UpdateModulAjar(ctx context.Context, modulAjar *domain.ModulAjar) error
	DeleteModulAjar(ctx context.Context, id string) error
}

// ReportingRepository defines the interface for reporting data access
type ReportingRepositoryInterface interface {
	// Narrative Report operations
	CreateNarrativeReport(ctx context.Context, report *domain.NarrativeReport) error
	GetNarrativeReportByID(ctx context.Context, id string) (*domain.NarrativeReport, error)
	ListNarrativeReports(ctx context.Context, studentID, userID *string, language *domain.ReportLanguage, status *domain.WorkflowStatus, limit, offset int) ([]*domain.NarrativeReport, error)
	UpdateNarrativeReport(ctx context.Context, report *domain.NarrativeReport) error
	DeleteNarrativeReport(ctx context.Context, id string) error
}

// SchoolRepository defines the interface for school data access
type SchoolRepositoryInterface interface {
	Create(ctx context.Context, school *domain.School) error
	GetByID(ctx context.Context, id string) (*domain.School, error)
	GetByCode(ctx context.Context, code string) (*domain.School, error)
	Update(ctx context.Context, school *domain.School) error
	UpdateStatus(ctx context.Context, id string, isActive bool) error
	List(ctx context.Context, isActive *bool, limit, offset int) ([]*domain.School, error)
	Count(ctx context.Context, isActive *bool) (int, error)
	Delete(ctx context.Context, id string) error
}

// CurriculumRepository defines the interface for curriculum data access
type CurriculumRepositoryInterface interface {
	// CP (Capaian Pembelajaran) operations
	CreateCP(ctx context.Context, cp *domain.CP) error
	GetCPByID(ctx context.Context, id string) (*domain.CP, error)
	ListCPs(ctx context.Context, subjectID, phaseID *string, limit, offset int) ([]*domain.CP, error)
	UpdateCP(ctx context.Context, cp *domain.CP) error
	DeleteCP(ctx context.Context, id string) error

	// Curriculum Subject operations
	CreateCurriculumSubject(ctx context.Context, subject *domain.CurriculumSubject) error
	GetCurriculumSubjectByID(ctx context.Context, id string) (*domain.CurriculumSubject, error)
	GetCurriculumSubjectByCode(ctx context.Context, code string) (*domain.CurriculumSubject, error)
	ListCurriculumSubjects(ctx context.Context, isActive *bool, limit, offset int) ([]*domain.CurriculumSubject, error)
	UpdateCurriculumSubject(ctx context.Context, subject *domain.CurriculumSubject) error
	DeleteCurriculumSubject(ctx context.Context, id string) error

	// Curriculum Phase operations
	CreateCurriculumPhase(ctx context.Context, phase *domain.CurriculumPhase) error
	GetCurriculumPhaseByID(ctx context.Context, id string) (*domain.CurriculumPhase, error)
	ListCurriculumPhases(ctx context.Context, isActive *bool, limit, offset int) ([]*domain.CurriculumPhase, error)
	UpdateCurriculumPhase(ctx context.Context, phase *domain.CurriculumPhase) error
	DeleteCurriculumPhase(ctx context.Context, id string) error

	// Curriculum Element operations
	CreateCurriculumElement(ctx context.Context, element *domain.CurriculumElement) error
	GetCurriculumElementByID(ctx context.Context, id string) (*domain.CurriculumElement, error)
	ListCurriculumElements(ctx context.Context, subjectID, phaseID *string, isActive *bool, limit, offset int) ([]*domain.CurriculumElement, error)
	UpdateCurriculumElement(ctx context.Context, element *domain.CurriculumElement) error
	DeleteCurriculumElement(ctx context.Context, id string) error

	// Curriculum Subelement operations
	CreateCurriculumSubelement(ctx context.Context, subelement *domain.CurriculumSubelement) error
	GetCurriculumSubelementByID(ctx context.Context, id string) (*domain.CurriculumSubelement, error)
	ListCurriculumSubelements(ctx context.Context, elementID *string, isActive *bool, limit, offset int) ([]*domain.CurriculumSubelement, error)
	UpdateCurriculumSubelement(ctx context.Context, subelement *domain.CurriculumSubelement) error
	DeleteCurriculumSubelement(ctx context.Context, id string) error
}

// ClassRepository defines the interface for class data access
type ClassRepositoryInterface interface {
	Create(ctx context.Context, class *domain.Class) error
	GetByID(ctx context.Context, id string) (*domain.Class, error)
	List(ctx context.Context, schoolID, academicYearID, semesterID, subjectID, teacherID *string, isActive *bool, limit, offset int) ([]*domain.Class, error)
	Count(ctx context.Context, schoolID, academicYearID, semesterID, subjectID, teacherID *string, isActive *bool) (int, error)
	Update(ctx context.Context, class *domain.Class) error
	Delete(ctx context.Context, id string) error
	GetStudentCount(ctx context.Context, classID string) (int, error)
}

// ClassEnrollmentRepository defines the interface for class enrollment data access
type ClassEnrollmentRepositoryInterface interface {
	Create(ctx context.Context, enrollment *domain.ClassEnrollment) error
	GetByID(ctx context.Context, id string) (*domain.ClassEnrollment, error)
	List(ctx context.Context, classID, studentID *string, status *string, limit, offset int) ([]*domain.ClassEnrollment, error)
	Update(ctx context.Context, enrollment *domain.ClassEnrollment) error
	Delete(ctx context.Context, id string) error
	CheckEnrollment(ctx context.Context, classID, studentID string) (bool, error)
}

// AnnouncementRepository defines the interface for announcement data access
type AnnouncementRepositoryInterface interface {
	Create(ctx context.Context, announcement *domain.Announcement) error
	GetByID(ctx context.Context, id string) (*domain.Announcement, error)
	List(ctx context.Context, schoolID, priority, targetAudience *string, isActive *bool, limit, offset int) ([]*domain.Announcement, error)
	Count(ctx context.Context, schoolID, priority, targetAudience *string, isActive *bool) (int, error)
	Update(ctx context.Context, announcement *domain.Announcement) error
	Delete(ctx context.Context, id string) error
	GetBySchoolID(ctx context.Context, schoolID string) ([]*domain.Announcement, error)
}

// AssignmentRepository defines the interface for assignment data access
type AssignmentRepositoryInterface interface {
	Create(ctx context.Context, assignment *domain.Assignment) error
	GetByID(ctx context.Context, id string) (*domain.Assignment, error)
	List(ctx context.Context, classID, assessmentID, status *string, limit, offset int) ([]*domain.Assignment, error)
	Count(ctx context.Context, classID, assessmentID, status *string) (int, error)
	Update(ctx context.Context, assignment *domain.Assignment) error
	Delete(ctx context.Context, id string) error
	GetByClassID(ctx context.Context, classID string) ([]*domain.Assignment, error)
}

// AttendanceRepository defines the interface for attendance data access
type AttendanceRepositoryInterface interface {
	Create(ctx context.Context, attendance *domain.AttendanceRecord) error
	GetByID(ctx context.Context, id string) (*domain.AttendanceRecord, error)
	List(ctx context.Context, classID, studentID *string, status *string, startDate, endDate *time.Time, limit, offset int) ([]*domain.AttendanceRecord, error)
	Count(ctx context.Context, classID, studentID *string, status *string, startDate, endDate *time.Time) (int, error)
	Update(ctx context.Context, attendance *domain.AttendanceRecord) error
	Delete(ctx context.Context, id string) error
	GetAttendanceStats(ctx context.Context, classID string, startDate, endDate time.Time) (map[string]int, error)
	GetStudentAttendanceStats(ctx context.Context, studentID string, startDate, endDate time.Time) (map[string]int, error)
}

// ScheduleRepository defines the interface for schedule data access
type ScheduleRepositoryInterface interface {
	Create(ctx context.Context, schedule *domain.Schedule) error
	GetByID(ctx context.Context, id string) (*domain.Schedule, error)
	List(ctx context.Context, classID *string, dayOfWeek *int, isActive *bool, limit, offset int) ([]*domain.Schedule, error)
	Count(ctx context.Context, classID *string, dayOfWeek *int, isActive *bool) (int, error)
	Update(ctx context.Context, schedule *domain.Schedule) error
	Delete(ctx context.Context, id string) error
	GetByClassID(ctx context.Context, classID string) ([]*domain.Schedule, error)
}

// NotificationRepository defines the interface for notification data access
type NotificationRepositoryInterface interface {
	Create(ctx context.Context, notification *domain.Notification) error
	GetByID(ctx context.Context, id string) (*domain.Notification, error)
	List(ctx context.Context, userID, notificationType *string, isRead *bool, limit, offset int) ([]*domain.Notification, error)
	Count(ctx context.Context, userID, notificationType *string, isRead *bool) (int, error)
	Update(ctx context.Context, notification *domain.Notification) error
	MarkAsRead(ctx context.Context, id string) error
	Delete(ctx context.Context, id string) error
	GetUnreadCount(ctx context.Context, userID string) (int, error)
}

// ExamRepository defines the interface for exam data access
type ExamRepositoryInterface interface {
	Create(ctx context.Context, exam *domain.Exam) error
	GetByID(ctx context.Context, id string) (*domain.Exam, error)
	List(ctx context.Context, classID, assessmentID, status *string, limit, offset int) ([]*domain.Exam, error)
	Count(ctx context.Context, classID, assessmentID, status *string) (int, error)
	Update(ctx context.Context, exam *domain.Exam) error
	Delete(ctx context.Context, id string) error
	GetByClassID(ctx context.Context, classID string) ([]*domain.Exam, error)
}

// ExamResultRepository defines the interface for exam result data access
type ExamResultRepositoryInterface interface {
	Create(ctx context.Context, result *domain.ExamResult) error
	GetByID(ctx context.Context, id string) (*domain.ExamResult, error)
	GetByExamAndStudent(ctx context.Context, examID, studentID string) (*domain.ExamResult, error)
	List(ctx context.Context, examID, studentID, grade *string, limit, offset int) ([]*domain.ExamResult, error)
	Count(ctx context.Context, examID, studentID, grade *string) (int, error)
	Update(ctx context.Context, result *domain.ExamResult) error
	Delete(ctx context.Context, id string) error
	GetByExamID(ctx context.Context, examID string) ([]*domain.ExamResult, error)
	GetByStudentID(ctx context.Context, studentID string) ([]*domain.ExamResult, error)
}

// MessageRepository defines the interface for message data access
type MessageRepositoryInterface interface {
	Create(ctx context.Context, message *domain.Message) error
	GetByID(ctx context.Context, id string) (*domain.Message, error)
	List(ctx context.Context, senderID, receiverID *string, isRead *bool, limit, offset int) ([]*domain.Message, error)
	Count(ctx context.Context, senderID, receiverID *string, isRead *bool) (int, error)
	Update(ctx context.Context, message *domain.Message) error
	MarkAsRead(ctx context.Context, id string) error
	Delete(ctx context.Context, id string) error
	GetConversation(ctx context.Context, userID1, userID2 string, limit, offset int) ([]*domain.Message, error)
	GetUnreadCount(ctx context.Context, userID string) (int, error)
}

// AcademicYearRepository defines the interface for academic year data access
type AcademicYearRepositoryInterface interface {
	Create(ctx context.Context, academicYear *domain.AcademicYear) error
	GetByID(ctx context.Context, id string) (*domain.AcademicYear, error)
	GetAcademicYearByID(ctx context.Context, id string) (*domain.AcademicYear, error)
	GetByYear(ctx context.Context, year int) (*domain.AcademicYear, error)
	List(ctx context.Context, isActive *bool, limit, offset int) ([]*domain.AcademicYear, error)
	Update(ctx context.Context, academicYear *domain.AcademicYear) error
	Delete(ctx context.Context, id string) error
	GetCurrent(ctx context.Context) (*domain.AcademicYear, error)
}

// SemesterRepository defines the interface for semester data access
type SemesterRepositoryInterface interface {
	Create(ctx context.Context, semester *domain.Semester) error
	GetByID(ctx context.Context, id string) (*domain.Semester, error)
	GetSemesterByID(ctx context.Context, id string) (*domain.Semester, error)
	GetByCode(ctx context.Context, code string) (*domain.Semester, error)
	List(ctx context.Context, academicYearID *string, isActive *bool, limit, offset int) ([]*domain.Semester, error)
	Update(ctx context.Context, semester *domain.Semester) error
	Delete(ctx context.Context, id string) error
	GetCurrent(ctx context.Context, academicYearID string) (*domain.Semester, error)
}

package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/handler/dto"
	"github.com/nusa/backend/internal/middleware"
	"github.com/nusa/backend/internal/service"
)

// ClassHandler handles HTTP requests for class endpoints
type ClassHandler struct {
	classService *service.ClassService
}

// NewClassHandler creates a new class handler
func NewClassHandler(classService *service.ClassService) *ClassHandler {
	return &ClassHandler{
		classService: classService,
	}
}

// CreateClass creates a new class
// POST /api/v1/classes
func (h *ClassHandler) CreateClass(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	var req dto.CreateClassRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	domainReq := &domain.CreateClassRequest{
		SchoolID:       req.SchoolID,
		AcademicYearID: req.AcademicYearID,
		SemesterID:     req.SemesterID,
		SubjectID:      req.SubjectID,
		TeacherID:      req.TeacherID,
		Name:           req.Name,
		GradeLevel:     req.GradeLevel,
		Room:           req.Room,
		MaxStudents:    req.MaxStudents,
	}

	class, err := h.classService.CreateClass(c.Request.Context(), domainReq, authCtx.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	currentStudents, _ := h.classService.GetStudentCount(c.Request.Context(), class.ID)
	domainResponse := class.ToClassResponse("", "", "", "", "", currentStudents)

	response := &dto.ClassResponse{
		ID:              domainResponse.ID,
		SchoolID:        domainResponse.SchoolID,
		SchoolName:      domainResponse.SchoolName,
		AcademicYearID:  domainResponse.AcademicYearID,
		AcademicYear:    domainResponse.AcademicYear,
		SemesterID:      domainResponse.SemesterID,
		SemesterName:    domainResponse.SemesterName,
		SubjectID:       domainResponse.SubjectID,
		SubjectName:     domainResponse.SubjectName,
		TeacherID:       domainResponse.TeacherID,
		TeacherName:     domainResponse.TeacherName,
		Name:            domainResponse.Name,
		GradeLevel:      domainResponse.GradeLevel,
		Room:            domainResponse.Room,
		MaxStudents:     domainResponse.MaxStudents,
		CurrentStudents: domainResponse.CurrentStudents,
		IsActive:        domainResponse.IsActive,
	}

	c.JSON(http.StatusCreated, response)
}

// GetClass retrieves a class by ID
// GET /api/v1/classes/:id
func (h *ClassHandler) GetClass(c *gin.Context) {
	id := c.Param("id")

	class, err := h.classService.GetClass(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Class not found"})
		return
	}

	currentStudents, _ := h.classService.GetStudentCount(c.Request.Context(), class.ID)
	domainResponse := class.ToClassResponse("", "", "", "", "", currentStudents)

	response := &dto.ClassResponse{
		ID:              domainResponse.ID,
		SchoolID:        domainResponse.SchoolID,
		SchoolName:      domainResponse.SchoolName,
		AcademicYearID:  domainResponse.AcademicYearID,
		AcademicYear:    domainResponse.AcademicYear,
		SemesterID:      domainResponse.SemesterID,
		SemesterName:    domainResponse.SemesterName,
		SubjectID:       domainResponse.SubjectID,
		SubjectName:     domainResponse.SubjectName,
		TeacherID:       domainResponse.TeacherID,
		TeacherName:     domainResponse.TeacherName,
		Name:            domainResponse.Name,
		GradeLevel:      domainResponse.GradeLevel,
		Room:            domainResponse.Room,
		MaxStudents:     domainResponse.MaxStudents,
		CurrentStudents: domainResponse.CurrentStudents,
		IsActive:        domainResponse.IsActive,
	}

	c.JSON(http.StatusOK, response)
}

// ListClasses retrieves classes with pagination and filters
// GET /api/v1/classes
func (h *ClassHandler) ListClasses(c *gin.Context) {
	var schoolID, academicYearID, semesterID, subjectID, teacherID *string
	var isActive *bool
	var err error

	if schoolIDStr := c.Query("school_id"); schoolIDStr != "" {
		schoolID = &schoolIDStr
	}
	if academicYearIDStr := c.Query("academic_year_id"); academicYearIDStr != "" {
		academicYearID = &academicYearIDStr
	}
	if semesterIDStr := c.Query("semester_id"); semesterIDStr != "" {
		semesterID = &semesterIDStr
	}
	if subjectIDStr := c.Query("subject_id"); subjectIDStr != "" {
		subjectID = &subjectIDStr
	}
	if teacherIDStr := c.Query("teacher_id"); teacherIDStr != "" {
		teacherID = &teacherIDStr
	}
	if isActiveStr := c.Query("is_active"); isActiveStr != "" {
		isActiveBool, err := strconv.ParseBool(isActiveStr)
		if err == nil {
			isActive = &isActiveBool
		}
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	classes, total, err := h.classService.ListClasses(c.Request.Context(), schoolID, academicYearID, semesterID, subjectID, teacherID, isActive, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	classResponses := make([]*dto.ClassResponse, len(classes))
	for i, class := range classes {
		currentStudents, _ := h.classService.GetStudentCount(c.Request.Context(), class.ID)
		domainResponse := class.ToClassResponse("", "", "", "", "", currentStudents)
		classResponses[i] = &dto.ClassResponse{
			ID:              domainResponse.ID,
			SchoolID:        domainResponse.SchoolID,
			SchoolName:      domainResponse.SchoolName,
			AcademicYearID:  domainResponse.AcademicYearID,
			AcademicYear:    domainResponse.AcademicYear,
			SemesterID:      domainResponse.SemesterID,
			SemesterName:    domainResponse.SemesterName,
			SubjectID:       domainResponse.SubjectID,
			SubjectName:     domainResponse.SubjectName,
			TeacherID:       domainResponse.TeacherID,
			TeacherName:     domainResponse.TeacherName,
			Name:            domainResponse.Name,
			GradeLevel:      domainResponse.GradeLevel,
			Room:            domainResponse.Room,
			MaxStudents:     domainResponse.MaxStudents,
			CurrentStudents: domainResponse.CurrentStudents,
			IsActive:        domainResponse.IsActive,
		}
	}

	c.JSON(http.StatusOK, dto.ClassListResponse{
		Classes:  classResponses,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	})
}

// UpdateClass updates class information
// PUT /api/v1/classes/:id
func (h *ClassHandler) UpdateClass(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	id := c.Param("id")

	var req dto.UpdateClassRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	domainReq := &domain.UpdateClassRequest{
		Name:        req.Name,
		GradeLevel:  req.GradeLevel,
		Room:        req.Room,
		MaxStudents: req.MaxStudents,
		IsActive:    req.IsActive,
	}

	class, err := h.classService.UpdateClass(c.Request.Context(), id, domainReq, authCtx.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	currentStudents, _ := h.classService.GetStudentCount(c.Request.Context(), class.ID)
	domainResponse := class.ToClassResponse("", "", "", "", "", currentStudents)

	response := &dto.ClassResponse{
		ID:              domainResponse.ID,
		SchoolID:        domainResponse.SchoolID,
		SchoolName:      domainResponse.SchoolName,
		AcademicYearID:  domainResponse.AcademicYearID,
		AcademicYear:    domainResponse.AcademicYear,
		SemesterID:      domainResponse.SemesterID,
		SemesterName:    domainResponse.SemesterName,
		SubjectID:       domainResponse.SubjectID,
		SubjectName:     domainResponse.SubjectName,
		TeacherID:       domainResponse.TeacherID,
		TeacherName:     domainResponse.TeacherName,
		Name:            domainResponse.Name,
		GradeLevel:      domainResponse.GradeLevel,
		Room:            domainResponse.Room,
		MaxStudents:     domainResponse.MaxStudents,
		CurrentStudents: domainResponse.CurrentStudents,
		IsActive:        domainResponse.IsActive,
	}

	c.JSON(http.StatusOK, response)
}

// DeleteClass soft deletes a class
// DELETE /api/v1/classes/:id
func (h *ClassHandler) DeleteClass(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	id := c.Param("id")

	if err := h.classService.DeleteClass(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// EnrollStudent enrolls a student in a class
// POST /api/v1/classes/:id/enrollments
func (h *ClassHandler) EnrollStudent(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	classID := c.Param("id")

	var req dto.CreateClassEnrollmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.ClassID = classID

	domainReq := &domain.CreateClassEnrollmentRequest{
		ClassID:   req.ClassID,
		StudentID: req.StudentID,
		Notes:     req.Notes,
	}

	enrollment, err := h.classService.EnrollStudent(c.Request.Context(), domainReq, authCtx.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	domainResponse := enrollment.ToClassEnrollmentResponse("", "")
	response := &dto.ClassEnrollmentResponse{
		ID:             domainResponse.ID,
		ClassID:        domainResponse.ClassID,
		ClassName:      domainResponse.ClassName,
		StudentID:      domainResponse.StudentID,
		StudentName:    domainResponse.StudentName,
		EnrollmentDate: enrollment.EnrollmentDate.Format(time.RFC3339),
		Status:         dto.EnrollmentStatus(domainResponse.Status),
		Notes:          domainResponse.Notes,
	}

	c.JSON(http.StatusCreated, response)
}

// WithdrawStudent withdraws a student from a class
// DELETE /api/v1/classes/:id/enrollments/:student_id
func (h *ClassHandler) WithdrawStudent(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	classID := c.Param("id")
	studentID := c.Param("student_id")

	if err := h.classService.WithdrawStudent(c.Request.Context(), classID, studentID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// ListEnrollments retrieves enrollments with pagination and filters
// GET /api/v1/classes/:id/enrollments
func (h *ClassHandler) ListEnrollments(c *gin.Context) {
	classID := c.Param("id")
	var studentID *string
	var status *string

	if studentIDStr := c.Query("student_id"); studentIDStr != "" {
		studentID = &studentIDStr
	}
	if statusStr := c.Query("status"); statusStr != "" {
		status = &statusStr
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	enrollments, err := h.classService.ListEnrollments(c.Request.Context(), &classID, studentID, status, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	enrollmentResponses := make([]*dto.ClassEnrollmentResponse, len(enrollments))
	for i, enrollment := range enrollments {
		domainResponse := enrollment.ToClassEnrollmentResponse("", "")
		enrollmentResponses[i] = &dto.ClassEnrollmentResponse{
			ID:             domainResponse.ID,
			ClassID:        domainResponse.ClassID,
			ClassName:      domainResponse.ClassName,
			StudentID:      domainResponse.StudentID,
			StudentName:    domainResponse.StudentName,
			EnrollmentDate: enrollment.EnrollmentDate.Format(time.RFC3339),
			Status:         dto.EnrollmentStatus(domainResponse.Status),
			Notes:          domainResponse.Notes,
		}
	}

	c.JSON(http.StatusOK, dto.ClassEnrollmentListResponse{
		Enrollments: enrollmentResponses,
		Total:       len(enrollmentResponses),
		Page:        page,
		PageSize:    pageSize,
	})
}

// UpdateEnrollment updates enrollment status
// PUT /api/v1/classes/:id/enrollments/:enrollment_id
func (h *ClassHandler) UpdateEnrollment(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	enrollmentID := c.Param("enrollment_id")

	var req dto.UpdateClassEnrollmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	domainReq := &domain.UpdateClassEnrollmentRequest{
		Status: domain.EnrollmentStatus(req.Status),
		Notes:  req.Notes,
	}

	if err := h.classService.UpdateEnrollment(c.Request.Context(), enrollmentID, domainReq); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Enrollment updated successfully"})
}

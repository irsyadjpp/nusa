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

// AttendanceHandler handles HTTP requests for attendance endpoints
type AttendanceHandler struct {
	attendanceService *service.AttendanceService
}

// NewAttendanceHandler creates a new attendance handler
func NewAttendanceHandler(attendanceService *service.AttendanceService) *AttendanceHandler {
	return &AttendanceHandler{
		attendanceService: attendanceService,
	}
}

// RecordAttendance records attendance for a student
// POST /api/v1/attendance
func (h *AttendanceHandler) RecordAttendance(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	var req dto.CreateAttendanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
		return
	}

	domainReq := &domain.CreateAttendanceRequest{
		ClassID:   req.ClassID,
		StudentID: req.StudentID,
		Date:      date,
		Status:    domain.AttendanceStatus(req.Status),
		Notes:     req.Notes,
	}

	attendance, err := h.attendanceService.RecordAttendance(c.Request.Context(), domainReq, authCtx.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	domainResponse := attendance.ToAttendanceResponse("", "", "")
	response := &dto.AttendanceResponse{
		ID:             domainResponse.ID,
		ClassID:        domainResponse.ClassID,
		ClassName:      domainResponse.ClassName,
		StudentID:      domainResponse.StudentID,
		StudentName:    domainResponse.StudentName,
		Date:           domainResponse.Date,
		Status:         dto.AttendanceStatus(domainResponse.Status),
		Notes:          domainResponse.Notes,
		RecordedBy:     domainResponse.RecordedBy,
		RecordedByName: domainResponse.RecordedByName,
	}

	c.JSON(http.StatusCreated, response)
}

// GetAttendance retrieves an attendance record by ID
// GET /api/v1/attendance/:id
func (h *AttendanceHandler) GetAttendance(c *gin.Context) {
	id := c.Param("id")

	attendance, err := h.attendanceService.GetAttendance(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Attendance record not found"})
		return
	}

	domainResponse := attendance.ToAttendanceResponse("", "", "")
	response := &dto.AttendanceResponse{
		ID:             domainResponse.ID,
		ClassID:        domainResponse.ClassID,
		ClassName:      domainResponse.ClassName,
		StudentID:      domainResponse.StudentID,
		StudentName:    domainResponse.StudentName,
		Date:           domainResponse.Date,
		Status:         dto.AttendanceStatus(domainResponse.Status),
		Notes:          domainResponse.Notes,
		RecordedBy:     domainResponse.RecordedBy,
		RecordedByName: domainResponse.RecordedByName,
	}

	c.JSON(http.StatusOK, response)
}

// ListAttendances retrieves attendance records with filters and pagination
// GET /api/v1/attendance
func (h *AttendanceHandler) ListAttendances(c *gin.Context) {
	var classID, studentID, status *string
	var startDate, endDate *time.Time
	var err error

	if classIDStr := c.Query("class_id"); classIDStr != "" {
		classID = &classIDStr
	}
	if studentIDStr := c.Query("student_id"); studentIDStr != "" {
		studentID = &studentIDStr
	}
	if statusStr := c.Query("status"); statusStr != "" {
		status = &statusStr
	}
	if startDateStr := c.Query("start_date"); startDateStr != "" {
		parsedDate, err := time.Parse("2006-01-02", startDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_date format. Use YYYY-MM-DD"})
			return
		}
		startDate = &parsedDate
	}
	if endDateStr := c.Query("end_date"); endDateStr != "" {
		parsedDate, err := time.Parse("2006-01-02", endDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_date format. Use YYYY-MM-DD"})
			return
		}
		endDate = &parsedDate
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	attendances, total, err := h.attendanceService.ListAttendances(c.Request.Context(), classID, studentID, status, startDate, endDate, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	attendanceResponses := make([]*dto.AttendanceResponse, len(attendances))
	for i, attendance := range attendances {
		domainResponse := attendance.ToAttendanceResponse("", "", "")
		attendanceResponses[i] = &dto.AttendanceResponse{
			ID:             domainResponse.ID,
			ClassID:        domainResponse.ClassID,
			ClassName:      domainResponse.ClassName,
			StudentID:      domainResponse.StudentID,
			StudentName:    domainResponse.StudentName,
			Date:           domainResponse.Date,
			Status:         dto.AttendanceStatus(domainResponse.Status),
			Notes:          domainResponse.Notes,
			RecordedBy:     domainResponse.RecordedBy,
			RecordedByName: domainResponse.RecordedByName,
		}
	}

	c.JSON(http.StatusOK, dto.AttendanceListResponse{
		Attendances: attendanceResponses,
		Total:       total,
		Page:        page,
		PageSize:    pageSize,
	})
}

// UpdateAttendance updates an attendance record
// PUT /api/v1/attendance/:id
func (h *AttendanceHandler) UpdateAttendance(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	id := c.Param("id")

	var req dto.UpdateAttendanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	domainReq := &domain.UpdateAttendanceRequest{
		Status: domain.AttendanceStatus(req.Status),
		Notes:  req.Notes,
	}

	attendance, err := h.attendanceService.UpdateAttendance(c.Request.Context(), id, domainReq, authCtx.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	domainResponse := attendance.ToAttendanceResponse("", "", "")
	response := &dto.AttendanceResponse{
		ID:             domainResponse.ID,
		ClassID:        domainResponse.ClassID,
		ClassName:      domainResponse.ClassName,
		StudentID:      domainResponse.StudentID,
		StudentName:    domainResponse.StudentName,
		Date:           domainResponse.Date,
		Status:         dto.AttendanceStatus(domainResponse.Status),
		Notes:          domainResponse.Notes,
		RecordedBy:     domainResponse.RecordedBy,
		RecordedByName: domainResponse.RecordedByName,
	}

	c.JSON(http.StatusOK, response)
}

// DeleteAttendance soft deletes an attendance record
// DELETE /api/v1/attendance/:id
func (h *AttendanceHandler) DeleteAttendance(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	id := c.Param("id")

	if err := h.attendanceService.DeleteAttendance(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// GetClassAttendanceStats returns attendance statistics for a class
// GET /api/v1/attendance/class/:classId/stats
func (h *AttendanceHandler) GetClassAttendanceStats(c *gin.Context) {
	classID := c.Param("classId")

	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	var startDate, endDate time.Time
	var err error

	if startDateStr != "" {
		startDate, err = time.Parse("2006-01-02", startDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_date format. Use YYYY-MM-DD"})
			return
		}
	} else {
		startDate = time.Now().AddDate(0, -1, 0)
	}

	if endDateStr != "" {
		endDate, err = time.Parse("2006-01-02", endDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_date format. Use YYYY-MM-DD"})
			return
		}
	} else {
		endDate = time.Now()
	}

	stats, err := h.attendanceService.GetClassAttendanceStats(c.Request.Context(), classID, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := &dto.AttendanceStatsResponse{
		Present: stats["PRESENT"],
		Absent:  stats["ABSENT"],
		Late:    stats["LATE"],
		Excused: stats["EXCUSED"],
		Total:   stats["PRESENT"] + stats["ABSENT"] + stats["LATE"] + stats["EXCUSED"],
	}

	c.JSON(http.StatusOK, response)
}

// GetStudentAttendanceStats returns attendance statistics for a student
// GET /api/v1/attendance/student/:studentId/stats
func (h *AttendanceHandler) GetStudentAttendanceStats(c *gin.Context) {
	studentID := c.Param("studentId")

	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	var startDate, endDate time.Time
	var err error

	if startDateStr != "" {
		startDate, err = time.Parse("2006-01-02", startDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_date format. Use YYYY-MM-DD"})
			return
		}
	} else {
		startDate = time.Now().AddDate(0, -1, 0)
	}

	if endDateStr != "" {
		endDate, err = time.Parse("2006-01-02", endDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_date format. Use YYYY-MM-DD"})
			return
		}
	} else {
		endDate = time.Now()
	}

	stats, err := h.attendanceService.GetStudentAttendanceStats(c.Request.Context(), studentID, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := &dto.AttendanceStatsResponse{
		Present: stats["PRESENT"],
		Absent:  stats["ABSENT"],
		Late:    stats["LATE"],
		Excused: stats["EXCUSED"],
		Total:   stats["PRESENT"] + stats["ABSENT"] + stats["LATE"] + stats["EXCUSED"],
	}

	c.JSON(http.StatusOK, response)
}

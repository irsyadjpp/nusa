package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/handler/dto"
	"github.com/nusa/backend/internal/middleware"
	"github.com/nusa/backend/internal/service"
)

// ExamHandler handles HTTP requests for exam endpoints
type ExamHandler struct {
	examService *service.ExamService
}

// NewExamHandler creates a new exam handler
func NewExamHandler(examService *service.ExamService) *ExamHandler {
	return &ExamHandler{
		examService: examService,
	}
}

// CreateExam creates a new exam
// POST /api/v1/exams
func (h *ExamHandler) CreateExam(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	var req dto.CreateExamRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	domainReq := &domain.CreateExamRequest{
		ClassID:        req.ClassID,
		AssessmentID:   req.AssessmentID,
		ExamDate:       req.ExamDate,
		StartTime:      req.StartTime,
		DurationMinutes: req.DurationMinutes,
		Room:           req.Room,
	}

	exam, err := h.examService.CreateExam(c.Request.Context(), domainReq, authCtx.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	domainResponse := exam.ToExamResponse("", "", "", "")
	response := &dto.ExamResponse{
		ID:             domainResponse.ID,
		ClassID:        domainResponse.ClassID,
		ClassName:      domainResponse.ClassName,
		AssessmentID:   domainResponse.AssessmentID,
		AssessmentType: domainResponse.AssessmentType,
		ExamDate:       domainResponse.ExamDate,
		StartTime:      domainResponse.StartTime,
		DurationMinutes: domainResponse.DurationMinutes,
		Room:           domainResponse.Room,
		Status:         dto.ExamStatus(domainResponse.Status),
		CreatedAt:      domainResponse.CreatedAt,
		UpdatedAt:      domainResponse.UpdatedAt,
		CreatedBy:      domainResponse.CreatedBy,
		CreatedByName:  domainResponse.CreatedByName,
		UpdatedBy:      domainResponse.UpdatedBy,
		UpdatedByName:  domainResponse.UpdatedByName,
	}

	c.JSON(http.StatusCreated, response)
}

// GetExam retrieves an exam by ID
// GET /api/v1/exams/:id
func (h *ExamHandler) GetExam(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	id := c.Param("id")

	exam, err := h.examService.GetExam(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Exam not found"})
		return
	}

	domainResponse := exam.ToExamResponse("", "", "", "")
	response := &dto.ExamResponse{
		ID:             domainResponse.ID,
		ClassID:        domainResponse.ClassID,
		ClassName:      domainResponse.ClassName,
		AssessmentID:   domainResponse.AssessmentID,
		AssessmentType: domainResponse.AssessmentType,
		ExamDate:       domainResponse.ExamDate,
		StartTime:      domainResponse.StartTime,
		DurationMinutes: domainResponse.DurationMinutes,
		Room:           domainResponse.Room,
		Status:         dto.ExamStatus(domainResponse.Status),
		CreatedAt:      domainResponse.CreatedAt,
		UpdatedAt:      domainResponse.UpdatedAt,
		CreatedBy:      domainResponse.CreatedBy,
		CreatedByName:  domainResponse.CreatedByName,
		UpdatedBy:      domainResponse.UpdatedBy,
		UpdatedByName:  domainResponse.UpdatedByName,
	}

	c.JSON(http.StatusOK, response)
}

// ListExams retrieves exams with filters and pagination
// GET /api/v1/exams
func (h *ExamHandler) ListExams(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	var classID, assessmentID, status *string

	if classIDStr := c.Query("class_id"); classIDStr != "" {
		classID = &classIDStr
	}
	if assessmentIDStr := c.Query("assessment_id"); assessmentIDStr != "" {
		assessmentID = &assessmentIDStr
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

	exams, total, err := h.examService.ListExams(c.Request.Context(), classID, assessmentID, status, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	examResponses := make([]*dto.ExamResponse, len(exams))
	for i, exam := range exams {
		domainResponse := exam.ToExamResponse("", "", "", "")
		examResponses[i] = &dto.ExamResponse{
			ID:             domainResponse.ID,
			ClassID:        domainResponse.ClassID,
			ClassName:      domainResponse.ClassName,
			AssessmentID:   domainResponse.AssessmentID,
			AssessmentType: domainResponse.AssessmentType,
			ExamDate:       domainResponse.ExamDate,
			StartTime:      domainResponse.StartTime,
			DurationMinutes: domainResponse.DurationMinutes,
			Room:           domainResponse.Room,
			Status:         dto.ExamStatus(domainResponse.Status),
			CreatedAt:      domainResponse.CreatedAt,
			UpdatedAt:      domainResponse.UpdatedAt,
			CreatedBy:      domainResponse.CreatedBy,
			CreatedByName:  domainResponse.CreatedByName,
			UpdatedBy:      domainResponse.UpdatedBy,
			UpdatedByName:  domainResponse.UpdatedByName,
		}
	}

	c.JSON(http.StatusOK, dto.ExamListResponse{
		Exams:    examResponses,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	})
}

// UpdateExam updates exam information
// PUT /api/v1/exams/:id
func (h *ExamHandler) UpdateExam(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	id := c.Param("id")

	var req dto.UpdateExamRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	domainReq := &domain.UpdateExamRequest{
		ExamDate:       req.ExamDate,
		StartTime:      req.StartTime,
		DurationMinutes: req.DurationMinutes,
		Room:           req.Room,
		Status:         (*domain.ExamStatus)(req.Status),
	}

	exam, err := h.examService.UpdateExam(c.Request.Context(), id, domainReq, authCtx.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	domainResponse := exam.ToExamResponse("", "", "", "")
	response := &dto.ExamResponse{
		ID:             domainResponse.ID,
		ClassID:        domainResponse.ClassID,
		ClassName:      domainResponse.ClassName,
		AssessmentID:   domainResponse.AssessmentID,
		AssessmentType: domainResponse.AssessmentType,
		ExamDate:       domainResponse.ExamDate,
		StartTime:      domainResponse.StartTime,
		DurationMinutes: domainResponse.DurationMinutes,
		Room:           domainResponse.Room,
		Status:         dto.ExamStatus(domainResponse.Status),
		CreatedAt:      domainResponse.CreatedAt,
		UpdatedAt:      domainResponse.UpdatedAt,
		CreatedBy:      domainResponse.CreatedBy,
		CreatedByName:  domainResponse.CreatedByName,
		UpdatedBy:      domainResponse.UpdatedBy,
		UpdatedByName:  domainResponse.UpdatedByName,
	}

	c.JSON(http.StatusOK, response)
}

// DeleteExam soft deletes an exam
// DELETE /api/v1/exams/:id
func (h *ExamHandler) DeleteExam(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	id := c.Param("id")

	if err := h.examService.DeleteExam(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// GetClassExams retrieves all exams for a class
// GET /api/v1/classes/:classId/exams
func (h *ExamHandler) GetClassExams(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	classID := c.Param("classId")

	exams, err := h.examService.GetClassExams(c.Request.Context(), classID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	examResponses := make([]*dto.ExamResponse, len(exams))
	for i, exam := range exams {
		domainResponse := exam.ToExamResponse("", "", "", "")
		examResponses[i] = &dto.ExamResponse{
			ID:             domainResponse.ID,
			ClassID:        domainResponse.ClassID,
			ClassName:      domainResponse.ClassName,
			AssessmentID:   domainResponse.AssessmentID,
			AssessmentType: domainResponse.AssessmentType,
			ExamDate:       domainResponse.ExamDate,
			StartTime:      domainResponse.StartTime,
			DurationMinutes: domainResponse.DurationMinutes,
			Room:           domainResponse.Room,
			Status:         dto.ExamStatus(domainResponse.Status),
			CreatedAt:      domainResponse.CreatedAt,
			UpdatedAt:      domainResponse.UpdatedAt,
			CreatedBy:      domainResponse.CreatedBy,
			CreatedByName:  domainResponse.CreatedByName,
			UpdatedBy:      domainResponse.UpdatedBy,
			UpdatedByName:  domainResponse.UpdatedByName,
		}
	}

	c.JSON(http.StatusOK, dto.ExamListResponse{
		Exams:    examResponses,
		Total:    len(examResponses),
		Page:     1,
		PageSize: len(examResponses),
	})
}

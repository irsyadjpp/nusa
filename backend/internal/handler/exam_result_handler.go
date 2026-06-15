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

// ExamResultHandler handles HTTP requests for exam result endpoints
type ExamResultHandler struct {
	examResultService *service.ExamResultService
}

// NewExamResultHandler creates a new exam result handler
func NewExamResultHandler(examResultService *service.ExamResultService) *ExamResultHandler {
	return &ExamResultHandler{
		examResultService: examResultService,
	}
}

// CreateExamResult creates a new exam result
// POST /api/v1/exam-results
func (h *ExamResultHandler) CreateExamResult(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	var req dto.CreateExamResultRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	domainReq := &domain.CreateExamResultRequest{
		ExamID:    req.ExamID,
		StudentID: req.StudentID,
		Score:     req.Score,
		Grade:     req.Grade,
		Remarks:   req.Remarks,
	}

	examResult, err := h.examResultService.CreateExamResult(c.Request.Context(), domainReq, authCtx.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	domainResponse := examResult.ToExamResultResponse("", "", "", "")
	response := &dto.ExamResultResponse{
		ID:           domainResponse.ID,
		ExamID:       domainResponse.ExamID,
		ExamDate:     domainResponse.ExamDate,
		ExamTitle:    domainResponse.ExamTitle,
		StudentID:    domainResponse.StudentID,
		StudentName:  domainResponse.StudentName,
		Score:        domainResponse.Score,
		Grade:        domainResponse.Grade,
		Remarks:      domainResponse.Remarks,
		GradedAt:     domainResponse.GradedAt,
		GradedBy:     domainResponse.GradedBy,
		GradedByName: domainResponse.GradedByName,
		CreatedAt:    domainResponse.CreatedAt,
		UpdatedAt:    domainResponse.UpdatedAt,
	}

	c.JSON(http.StatusCreated, response)
}

// GetExamResult retrieves an exam result by ID
// GET /api/v1/exam-results/:id
func (h *ExamResultHandler) GetExamResult(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	id := c.Param("id")

	examResult, err := h.examResultService.GetExamResult(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Exam result not found"})
		return
	}

	domainResponse := examResult.ToExamResultResponse("", "", "", "")
	response := &dto.ExamResultResponse{
		ID:           domainResponse.ID,
		ExamID:       domainResponse.ExamID,
		ExamDate:     domainResponse.ExamDate,
		ExamTitle:    domainResponse.ExamTitle,
		StudentID:    domainResponse.StudentID,
		StudentName:  domainResponse.StudentName,
		Score:        domainResponse.Score,
		Grade:        domainResponse.Grade,
		Remarks:      domainResponse.Remarks,
		GradedAt:     domainResponse.GradedAt,
		GradedBy:     domainResponse.GradedBy,
		GradedByName: domainResponse.GradedByName,
		CreatedAt:    domainResponse.CreatedAt,
		UpdatedAt:    domainResponse.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

// ListExamResults retrieves exam results with filters and pagination
// GET /api/v1/exam-results
func (h *ExamResultHandler) ListExamResults(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	var examID, studentID, grade *string

	if examIDStr := c.Query("exam_id"); examIDStr != "" {
		examID = &examIDStr
	}
	if studentIDStr := c.Query("student_id"); studentIDStr != "" {
		studentID = &studentIDStr
	}
	if gradeStr := c.Query("grade"); gradeStr != "" {
		grade = &gradeStr
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	examResults, total, err := h.examResultService.ListExamResults(c.Request.Context(), examID, studentID, grade, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	examResultResponses := make([]*dto.ExamResultResponse, len(examResults))
	for i, examResult := range examResults {
		domainResponse := examResult.ToExamResultResponse("", "", "", "")
		examResultResponses[i] = &dto.ExamResultResponse{
			ID:           domainResponse.ID,
			ExamID:       domainResponse.ExamID,
			ExamDate:     domainResponse.ExamDate,
			ExamTitle:    domainResponse.ExamTitle,
			StudentID:    domainResponse.StudentID,
			StudentName:  domainResponse.StudentName,
			Score:        domainResponse.Score,
			Grade:        domainResponse.Grade,
			Remarks:      domainResponse.Remarks,
			GradedAt:     domainResponse.GradedAt,
			GradedBy:     domainResponse.GradedBy,
			GradedByName: domainResponse.GradedByName,
			CreatedAt:    domainResponse.CreatedAt,
			UpdatedAt:    domainResponse.UpdatedAt,
		}
	}

	c.JSON(http.StatusOK, dto.ExamResultListResponse{
		ExamResults: examResultResponses,
		Total:       total,
		Page:        page,
		PageSize:    pageSize,
	})
}

// UpdateExamResult updates exam result information
// PUT /api/v1/exam-results/:id
func (h *ExamResultHandler) UpdateExamResult(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	id := c.Param("id")

	var req dto.UpdateExamResultRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	domainReq := &domain.UpdateExamResultRequest{
		Score:    req.Score,
		Grade:    req.Grade,
		Remarks:  req.Remarks,
		GradedBy: req.GradedBy,
	}

	examResult, err := h.examResultService.UpdateExamResult(c.Request.Context(), id, domainReq, authCtx.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	domainResponse := examResult.ToExamResultResponse("", "", "", "")
	response := &dto.ExamResultResponse{
		ID:           domainResponse.ID,
		ExamID:       domainResponse.ExamID,
		ExamDate:     domainResponse.ExamDate,
		ExamTitle:    domainResponse.ExamTitle,
		StudentID:    domainResponse.StudentID,
		StudentName:  domainResponse.StudentName,
		Score:        domainResponse.Score,
		Grade:        domainResponse.Grade,
		Remarks:      domainResponse.Remarks,
		GradedAt:     domainResponse.GradedAt,
		GradedBy:     domainResponse.GradedBy,
		GradedByName: domainResponse.GradedByName,
		CreatedAt:    domainResponse.CreatedAt,
		UpdatedAt:    domainResponse.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

// DeleteExamResult soft deletes an exam result
// DELETE /api/v1/exam-results/:id
func (h *ExamResultHandler) DeleteExamResult(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	id := c.Param("id")

	if err := h.examResultService.DeleteExamResult(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// GetExamResultsByExam retrieves all exam results for an exam
// GET /api/v1/exams/:examId/results
func (h *ExamResultHandler) GetExamResultsByExam(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	examID := c.Param("examId")

	examResults, err := h.examResultService.GetExamResultsByExam(c.Request.Context(), examID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	examResultResponses := make([]*dto.ExamResultResponse, len(examResults))
	for i, examResult := range examResults {
		domainResponse := examResult.ToExamResultResponse("", "", "", "")
		examResultResponses[i] = &dto.ExamResultResponse{
			ID:           domainResponse.ID,
			ExamID:       domainResponse.ExamID,
			ExamDate:     domainResponse.ExamDate,
			ExamTitle:    domainResponse.ExamTitle,
			StudentID:    domainResponse.StudentID,
			StudentName:  domainResponse.StudentName,
			Score:        domainResponse.Score,
			Grade:        domainResponse.Grade,
			Remarks:      domainResponse.Remarks,
			GradedAt:     domainResponse.GradedAt,
			GradedBy:     domainResponse.GradedBy,
			GradedByName: domainResponse.GradedByName,
			CreatedAt:    domainResponse.CreatedAt,
			UpdatedAt:    domainResponse.UpdatedAt,
		}
	}

	c.JSON(http.StatusOK, dto.ExamResultListResponse{
		ExamResults: examResultResponses,
		Total:       len(examResultResponses),
		Page:        1,
		PageSize:    len(examResultResponses),
	})
}

// GetExamResultsByStudent retrieves all exam results for a student
// GET /api/v1/students/:studentId/exam-results
func (h *ExamResultHandler) GetExamResultsByStudent(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	studentID := c.Param("studentId")

	examResults, err := h.examResultService.GetExamResultsByStudent(c.Request.Context(), studentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	examResultResponses := make([]*dto.ExamResultResponse, len(examResults))
	for i, examResult := range examResults {
		domainResponse := examResult.ToExamResultResponse("", "", "", "")
		examResultResponses[i] = &dto.ExamResultResponse{
			ID:           domainResponse.ID,
			ExamID:       domainResponse.ExamID,
			ExamDate:     domainResponse.ExamDate,
			ExamTitle:    domainResponse.ExamTitle,
			StudentID:    domainResponse.StudentID,
			StudentName:  domainResponse.StudentName,
			Score:        domainResponse.Score,
			Grade:        domainResponse.Grade,
			Remarks:      domainResponse.Remarks,
			GradedAt:     domainResponse.GradedAt,
			GradedBy:     domainResponse.GradedBy,
			GradedByName: domainResponse.GradedByName,
			CreatedAt:    domainResponse.CreatedAt,
			UpdatedAt:    domainResponse.UpdatedAt,
		}
	}

	c.JSON(http.StatusOK, dto.ExamResultListResponse{
		ExamResults: examResultResponses,
		Total:       len(examResultResponses),
		Page:        1,
		PageSize:    len(examResultResponses),
	})
}

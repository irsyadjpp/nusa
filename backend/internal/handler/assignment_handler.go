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

// AssignmentHandler handles HTTP requests for assignment endpoints
type AssignmentHandler struct {
	assignmentService *service.AssignmentService
}

// NewAssignmentHandler creates a new assignment handler
func NewAssignmentHandler(assignmentService *service.AssignmentService) *AssignmentHandler {
	return &AssignmentHandler{
		assignmentService: assignmentService,
	}
}

// CreateAssignment creates a new assignment
// POST /api/v1/assignments
func (h *AssignmentHandler) CreateAssignment(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	var req dto.CreateAssignmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	domainReq := &domain.CreateAssignmentRequest{
		ClassID:      req.ClassID,
		AssessmentID: req.AssessmentID,
		Title:        req.Title,
		Description:  req.Description,
		DueDate:      req.DueDate,
		MaxScore:     req.MaxScore,
	}

	assignment, err := h.assignmentService.CreateAssignment(c.Request.Context(), domainReq, authCtx.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	domainResponse := assignment.ToAssignmentResponse("", "", "", "")
	response := &dto.AssignmentResponse{
		ID:             domainResponse.ID,
		ClassID:        domainResponse.ClassID,
		ClassName:      domainResponse.ClassName,
		AssessmentID:   domainResponse.AssessmentID,
		AssessmentType: domainResponse.AssessmentType,
		Title:          domainResponse.Title,
		Description:    domainResponse.Description,
		DueDate:        domainResponse.DueDate,
		MaxScore:       domainResponse.MaxScore,
		Status:         dto.AssignmentStatus(domainResponse.Status),
		CreatedAt:      domainResponse.CreatedAt,
		UpdatedAt:      domainResponse.UpdatedAt,
		CreatedBy:      domainResponse.CreatedBy,
		CreatedByName:  domainResponse.CreatedByName,
		UpdatedBy:      domainResponse.UpdatedBy,
		UpdatedByName:  domainResponse.UpdatedByName,
	}

	c.JSON(http.StatusCreated, response)
}

// GetAssignment retrieves an assignment by ID
// GET /api/v1/assignments/:id
func (h *AssignmentHandler) GetAssignment(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	id := c.Param("id")

	assignment, err := h.assignmentService.GetAssignment(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Assignment not found"})
		return
	}

	domainResponse := assignment.ToAssignmentResponse("", "", "", "")
	response := &dto.AssignmentResponse{
		ID:             domainResponse.ID,
		ClassID:        domainResponse.ClassID,
		ClassName:      domainResponse.ClassName,
		AssessmentID:   domainResponse.AssessmentID,
		AssessmentType: domainResponse.AssessmentType,
		Title:          domainResponse.Title,
		Description:    domainResponse.Description,
		DueDate:        domainResponse.DueDate,
		MaxScore:       domainResponse.MaxScore,
		Status:         dto.AssignmentStatus(domainResponse.Status),
		CreatedAt:      domainResponse.CreatedAt,
		UpdatedAt:      domainResponse.UpdatedAt,
		CreatedBy:      domainResponse.CreatedBy,
		CreatedByName:  domainResponse.CreatedByName,
		UpdatedBy:      domainResponse.UpdatedBy,
		UpdatedByName:  domainResponse.UpdatedByName,
	}

	c.JSON(http.StatusOK, response)
}

// ListAssignments retrieves assignments with filters and pagination
// GET /api/v1/assignments
func (h *AssignmentHandler) ListAssignments(c *gin.Context) {
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

	assignments, total, err := h.assignmentService.ListAssignments(c.Request.Context(), classID, assessmentID, status, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	assignmentResponses := make([]*dto.AssignmentResponse, len(assignments))
	for i, assignment := range assignments {
		domainResponse := assignment.ToAssignmentResponse("", "", "", "")
		assignmentResponses[i] = &dto.AssignmentResponse{
			ID:             domainResponse.ID,
			ClassID:        domainResponse.ClassID,
			ClassName:      domainResponse.ClassName,
			AssessmentID:   domainResponse.AssessmentID,
			AssessmentType: domainResponse.AssessmentType,
			Title:          domainResponse.Title,
			Description:    domainResponse.Description,
			DueDate:        domainResponse.DueDate,
			MaxScore:       domainResponse.MaxScore,
			Status:         dto.AssignmentStatus(domainResponse.Status),
			CreatedAt:      domainResponse.CreatedAt,
			UpdatedAt:      domainResponse.UpdatedAt,
			CreatedBy:      domainResponse.CreatedBy,
			CreatedByName:  domainResponse.CreatedByName,
			UpdatedBy:      domainResponse.UpdatedBy,
			UpdatedByName:  domainResponse.UpdatedByName,
		}
	}

	c.JSON(http.StatusOK, dto.AssignmentListResponse{
		Assignments: assignmentResponses,
		Total:        total,
		Page:         page,
		PageSize:     pageSize,
	})
}

// UpdateAssignment updates assignment information
// PUT /api/v1/assignments/:id
func (h *AssignmentHandler) UpdateAssignment(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	id := c.Param("id")

	var req dto.UpdateAssignmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	domainReq := &domain.UpdateAssignmentRequest{
		Title:       req.Title,
		Description: req.Description,
		DueDate:     req.DueDate,
		MaxScore:    req.MaxScore,
		Status:      (*domain.AssignmentStatus)(req.Status),
	}

	assignment, err := h.assignmentService.UpdateAssignment(c.Request.Context(), id, domainReq, authCtx.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	domainResponse := assignment.ToAssignmentResponse("", "", "", "")
	response := &dto.AssignmentResponse{
		ID:             domainResponse.ID,
		ClassID:        domainResponse.ClassID,
		ClassName:      domainResponse.ClassName,
		AssessmentID:   domainResponse.AssessmentID,
		AssessmentType: domainResponse.AssessmentType,
		Title:          domainResponse.Title,
		Description:    domainResponse.Description,
		DueDate:        domainResponse.DueDate,
		MaxScore:       domainResponse.MaxScore,
		Status:         dto.AssignmentStatus(domainResponse.Status),
		CreatedAt:      domainResponse.CreatedAt,
		UpdatedAt:      domainResponse.UpdatedAt,
		CreatedBy:      domainResponse.CreatedBy,
		CreatedByName:  domainResponse.CreatedByName,
		UpdatedBy:      domainResponse.UpdatedBy,
		UpdatedByName:  domainResponse.UpdatedByName,
	}

	c.JSON(http.StatusOK, response)
}

// DeleteAssignment soft deletes an assignment
// DELETE /api/v1/assignments/:id
func (h *AssignmentHandler) DeleteAssignment(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	id := c.Param("id")

	if err := h.assignmentService.DeleteAssignment(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// GetClassAssignments retrieves all assignments for a class
// GET /api/v1/classes/:classId/assignments
func (h *AssignmentHandler) GetClassAssignments(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	classID := c.Param("classId")

	assignments, err := h.assignmentService.GetClassAssignments(c.Request.Context(), classID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	assignmentResponses := make([]*dto.AssignmentResponse, len(assignments))
	for i, assignment := range assignments {
		domainResponse := assignment.ToAssignmentResponse("", "", "", "")
		assignmentResponses[i] = &dto.AssignmentResponse{
			ID:             domainResponse.ID,
			ClassID:        domainResponse.ClassID,
			ClassName:      domainResponse.ClassName,
			AssessmentID:   domainResponse.AssessmentID,
			AssessmentType: domainResponse.AssessmentType,
			Title:          domainResponse.Title,
			Description:    domainResponse.Description,
			DueDate:        domainResponse.DueDate,
			MaxScore:       domainResponse.MaxScore,
			Status:         dto.AssignmentStatus(domainResponse.Status),
			CreatedAt:      domainResponse.CreatedAt,
			UpdatedAt:      domainResponse.UpdatedAt,
			CreatedBy:      domainResponse.CreatedBy,
			CreatedByName:  domainResponse.CreatedByName,
			UpdatedBy:      domainResponse.UpdatedBy,
			UpdatedByName:  domainResponse.UpdatedByName,
		}
	}

	c.JSON(http.StatusOK, dto.AssignmentListResponse{
		Assignments: assignmentResponses,
		Total:       len(assignmentResponses),
		Page:        1,
		PageSize:    len(assignmentResponses),
	})
}

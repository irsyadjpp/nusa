package handler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nusa/backend/internal/application"
	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/handler/dto"
	"github.com/nusa/backend/internal/middleware"
)

// ITPSetApplicationService defines the interface for TP Set application service
type ITPSetApplicationService interface {
	CreateTPSet(ctx context.Context, cmd *application.CreateTPSetCommand) (*application.CreateTPSetResponse, error)
	ListTPSets(ctx context.Context, query *application.ListTPSetsQuery) (*application.ListTPSetsResponse, error)
	GetTPSet(ctx context.Context, query *application.GetTPSetQuery) (*application.GetTPSetResponse, error)
	ApproveTPSet(ctx context.Context, cmd *application.ApproveTPSetCommand) (*application.ApproveTPSetResponse, error)
	CreateTP(ctx context.Context, cmd *application.CreateTPCommand) (*application.CreateTPResponse, error)
	ListTPs(ctx context.Context, query *application.ListTPsQuery) (*application.ListTPsResponse, error)
	GetTP(ctx context.Context, query *application.GetTPQuery) (*application.GetTPResponse, error)
}

// TPSetHandler handles HTTP requests for TP Set endpoints
type TPSetHandler struct {
	tpSetApplicationService ITPSetApplicationService
}

// NewTPSetHandler creates a new TP Set handler
func NewTPSetHandler(tpSetApplicationService *application.TPSetApplicationService) *TPSetHandler {
	return &TPSetHandler{
		tpSetApplicationService: tpSetApplicationService,
	}
}

// CreateTPSet creates a new TP Set
// POST /tp-sets
// @Summary Create a new TP Set
// @Description Create a new Teaching Plan Set from a CP
// @Tags TP (Teaching Plan)
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body dto.CreateTPSetRequest true "Create TP Set request"
// @Success 201 {object} dto.TPSetResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /tp-sets [post]
func (h *TPSetHandler) CreateTPSet(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "UNAUTHORIZED",
				Message: "Authentication required",
			},
		})
		return
	}

	var req dto.CreateTPSetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INVALID_REQUEST",
				Message: err.Error(),
			},
		})
		return
	}

	cmd := &application.CreateTPSetCommand{
		UserID:           authCtx.UserID,
		CPID:             req.CPID,
		VersionNo:        req.VersionNo,
		GenerationSource: domain.GenerationSource(req.GenerationSource),
		GenerationReason: &req.GenerationReason,
	}

	resp, err := h.tpSetApplicationService.CreateTPSet(c.Request.Context(), cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	// Fetch the created TP Set to return full response
	tpSet, err := h.tpSetApplicationService.GetTPSet(c.Request.Context(), &application.GetTPSetQuery{
		UserID:  authCtx.UserID,
		TPSetID: resp.TPSetID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	tpSetResp := h.mapToTPSetResponse(tpSet.TPSet)
	c.JSON(http.StatusCreated, tpSetResp)
}

// ListTPSets lists TP Sets
// GET /tp-sets
// @Summary List TP Sets
// @Description Get a paginated list of TP Sets
// @Tags TP (Teaching Plan)
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param cp_id query string false "Filter by CP ID"
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(20)
// @Success 200 {object} dto.ListTPSetsResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /tp-sets [get]
func (h *TPSetHandler) ListTPSets(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "UNAUTHORIZED",
				Message: "Authentication required",
			},
		})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	var cpID *string
	if cpIDStr := c.Query("cp_id"); cpIDStr != "" {
		cpID = &cpIDStr
	}

	query := &application.ListTPSetsQuery{
		UserID:   authCtx.UserID,
		CPID:     cpID,
		Page:     page,
		PageSize: pageSize,
	}

	resp, err := h.tpSetApplicationService.ListTPSets(c.Request.Context(), query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	tpSetResponses := make([]dto.TPSetResponse, 0, len(resp.TPSets))
	for _, tpSet := range resp.TPSets {
		tpSetResponses = append(tpSetResponses, h.mapToTPSetResponse(tpSet))
	}

	c.JSON(http.StatusOK, dto.ListTPSetsResponse{
		TPSets:   tpSetResponses,
		Total:    resp.Total,
		Page:     resp.Page,
		PageSize: resp.PageSize,
	})
}

// GetTPSet gets a TP Set by ID
// GET /tp-sets/{id}
// @Summary Get TP Set by ID
// @Description Retrieve a specific TP Set
// @Tags TP (Teaching Plan)
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "TP Set ID"
// @Success 200 {object} dto.TPSetResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /tp-sets/{id} [get]
func (h *TPSetHandler) GetTPSet(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "UNAUTHORIZED",
				Message: "Authentication required",
			},
		})
		return
	}

	id := c.Param("id")
	query := &application.GetTPSetQuery{
		UserID:  authCtx.UserID,
		TPSetID: id,
	}

	resp, err := h.tpSetApplicationService.GetTPSet(c.Request.Context(), query)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "NOT_FOUND",
				Message: "TP Set not found",
			},
		})
		return
	}

	c.JSON(http.StatusOK, h.mapToTPSetResponse(resp.TPSet))
}

// ApproveTPSet approves a TP Set
// POST /tp-sets/{id}/approve
// @Summary Approve TP Set
// @Description Approve a TP Set
// @Tags TP (Teaching Plan)
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "TP Set ID"
// @Param request body dto.ApproveTPSetRequest true "Approve TP Set request"
// @Success 200 {object} dto.ApproveTPSetResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /tp-sets/{id}/approve [post]
func (h *TPSetHandler) ApproveTPSet(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "UNAUTHORIZED",
				Message: "Authentication required",
			},
		})
		return
	}

	// RBAC: Only School Admin or System Admin can approve
	if authCtx.Role != "SCHOOL_ADMIN" && authCtx.Role != "SYSTEM_ADMIN" {
		c.JSON(http.StatusForbidden, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "FORBIDDEN",
				Message: "Only School Admin or System Admin can approve TP Sets",
			},
		})
		return
	}

	id := c.Param("id")

	var req dto.ApproveTPSetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INVALID_REQUEST",
				Message: err.Error(),
			},
		})
		return
	}

	cmd := &application.ApproveTPSetCommand{
		ApproverID: authCtx.UserID,
		TPSetID:    id,
		Reason:     req.Reason,
	}

	_, err := h.tpSetApplicationService.ApproveTPSet(c.Request.Context(), cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, dto.ApproveTPSetResponse{
		Message: "TP Set approved successfully",
	})
}

// CreateTP creates a new TP
// POST /tps
// @Summary Create a new TP
// @Description Create a new Teaching Plan item
// @Tags TP (Teaching Plan)
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body dto.CreateTPRequest true "Create TP request"
// @Success 201 {object} dto.TPResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /tps [post]
func (h *TPSetHandler) CreateTP(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "UNAUTHORIZED",
				Message: "Authentication required",
			},
		})
		return
	}

	var req dto.CreateTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INVALID_REQUEST",
				Message: err.Error(),
			},
		})
		return
	}

	// Convert DTO types to application service types
	var elementID *string
	if req.ElementID != "" {
		elementID = &req.ElementID
	}
	var subelementID *string
	if req.SubelementID != "" {
		subelementID = &req.SubelementID
	}

	learningObjectives, _ := req.LearningObjectives.(map[string]interface{})
	timeAllocation, _ := req.TimeAllocation.(map[string]interface{})

	var prerequisites string
	if req.Prerequisites != nil {
		if s, ok := req.Prerequisites.(string); ok {
			prerequisites = s
		}
	}

	var successCriteria map[string]interface{}
	if req.SuccessCriteria != nil {
		if m, ok := req.SuccessCriteria.(map[string]interface{}); ok {
			successCriteria = m
		}
	} else {
		successCriteria = make(map[string]interface{})
	}

	cmd := &application.CreateTPCommand{
		TPSetID:            req.TPSetID,
		SequenceNumber:     req.SequenceNumber,
		CPID:               req.CPID,
		SubjectID:          req.SubjectID,
		PhaseID:            req.PhaseID,
		ElementID:          elementID,
		SubelementID:       subelementID,
		Title:              req.Title,
		LearningObjectives: learningObjectives,
		TimeAllocation:     timeAllocation,
		Prerequisites:      prerequisites,
		EstimatedWeeks:     req.EstimatedWeeks,
		SuccessCriteria:    successCriteria,
		UserID:             authCtx.UserID,
	}

	resp, err := h.tpSetApplicationService.CreateTP(c.Request.Context(), cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	// Fetch the created TP to return full response
	tp, err := h.tpSetApplicationService.GetTP(c.Request.Context(), &application.GetTPQuery{
		UserID: authCtx.UserID,
		TPID:   resp.TPID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	tpResp := h.mapToTPResponse(tp.TP)
	c.JSON(http.StatusCreated, tpResp)
}

// ListTPs lists TPs
// GET /tps
// @Summary List TPs
// @Description Get a paginated list of Teaching Plans
// @Tags TP (Teaching Plan)
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param tp_set_id query string false "Filter by TP Set ID"
// @Param status query string false "Filter by status" Enums(DRAFT, PENDING, APPROVED, REJECTED)
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(20)
// @Success 200 {object} dto.ListTPsResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /tps [get]
func (h *TPSetHandler) ListTPs(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "UNAUTHORIZED",
				Message: "Authentication required",
			},
		})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	var tpSetID *string
	if tpSetIDStr := c.Query("tp_set_id"); tpSetIDStr != "" {
		tpSetID = &tpSetIDStr
	}

	var status *domain.WorkflowStatus
	if statusStr := c.Query("status"); statusStr != "" {
		s := domain.WorkflowStatus(statusStr)
		status = &s
	}

	query := &application.ListTPsQuery{
		TPSetID:  tpSetID,
		Status:   status,
		UserID:   authCtx.UserID,
		Page:     page,
		PageSize: pageSize,
	}

	resp, err := h.tpSetApplicationService.ListTPs(c.Request.Context(), query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	tpResponses := make([]dto.TPResponse, 0, len(resp.TPs))
	for _, tp := range resp.TPs {
		tpResponses = append(tpResponses, h.mapToTPResponse(tp))
	}

	c.JSON(http.StatusOK, dto.ListTPsResponse{
		TPs:      tpResponses,
		Total:    resp.Total,
		Page:     resp.Page,
		PageSize: resp.PageSize,
	})
}

// GetTP gets a TP by ID
// GET /tps/{id}
// @Summary Get TP by ID
// @Description Retrieve a specific Teaching Plan
// @Tags TP (Teaching Plan)
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "TP ID"
// @Success 200 {object} dto.TPResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /tps/{id} [get]
func (h *TPSetHandler) GetTP(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "UNAUTHORIZED",
				Message: "Authentication required",
			},
		})
		return
	}

	id := c.Param("id")
	query := &application.GetTPQuery{
		UserID: authCtx.UserID,
		TPID:   id,
	}

	resp, err := h.tpSetApplicationService.GetTP(c.Request.Context(), query)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "NOT_FOUND",
				Message: "TP not found",
			},
		})
		return
	}

	c.JSON(http.StatusOK, h.mapToTPResponse(resp.TP))
}

// Helper functions

func (h *TPSetHandler) mapToTPSetResponse(tpSet *domain.TPSet) dto.TPSetResponse {
	resp := dto.TPSetResponse{
		ID:               tpSet.ID,
		CPID:             tpSet.CPID,
		CPCode:           "", // Would need to fetch from CP repository
		CPText:           "", // Would need to fetch from CP repository
		VersionNo:        tpSet.VersionNo,
		Status:           string(tpSet.Status),
		GenerationSource: string(tpSet.GenerationSource),
		GenerationReason: tpSet.GenerationReason,
		GeneratedBy:      tpSet.GeneratedBy,
		GeneratedByName:  "", // Would need to fetch from user repository
		AIGenerationID:   tpSet.AIGenerationID,
		ApprovedBy:       tpSet.ApprovedBy,
		ApprovedByName:   nil, // Would need to fetch from user repository
		CreatedAt:        tpSet.CreatedAt.Format(time.RFC3339),
		UpdatedAt:        tpSet.UpdatedAt.Format(time.RFC3339),
	}

	if tpSet.ApprovedAt != nil {
		approvedAt := tpSet.ApprovedAt.Format(time.RFC3339)
		resp.ApprovedAt = &approvedAt
	}

	return resp
}

func (h *TPSetHandler) mapToTPResponse(tp *domain.TP) dto.TPResponse {
	resp := dto.TPResponse{
		ID:                 tp.ID,
		TPSetID:            tp.TPSetID,
		SequenceNumber:     tp.SequenceNumber,
		CPID:               tp.CPID,
		CPCode:             "", // Would need to fetch from CP repository
		CPText:             "", // Would need to fetch from CP repository
		SubjectID:          tp.SubjectID,
		SubjectCode:        "", // Would need to fetch from subject repository
		SubjectName:        "", // Would need to fetch from subject repository
		PhaseID:            tp.PhaseID,
		PhaseCode:          "", // Would need to fetch from phase repository
		PhaseName:          "", // Would need to fetch from phase repository
		ElementID:          tp.ElementID,
		ElementCode:        "", // Would need to fetch from element repository
		ElementName:        "", // Would need to fetch from element repository
		SubelementID:       tp.SubelementID,
		SubelementCode:     "", // Would need to fetch from subelement repository
		SubelementName:     "", // Would need to fetch from subelement repository
		UserID:             tp.UserID,
		UserName:           "", // Would need to fetch from user repository
		Status:             string(tp.Status),
		Title:              tp.Title,
		LearningObjectives: tp.LearningObjectives,
		TimeAllocation:     tp.TimeAllocation,
		Prerequisites:      tp.Prerequisites,
		EstimatedWeeks:     tp.EstimatedWeeks,
		SuccessCriteria:    tp.SuccessCriteria,
		CreatedAt:          tp.CreatedAt.Format(time.RFC3339),
		UpdatedAt:          tp.UpdatedAt.Format(time.RFC3339),
	}

	return resp
}

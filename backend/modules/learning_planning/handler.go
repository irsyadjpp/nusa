package learning_planning

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/middleware"
	"github.com/nusa/backend/internal/service"
	"github.com/nusa/backend/pkg/response"
)

type Handler struct {
	tpService    *service.TPService
	atpService   *service.LearningPlanningService
	modulAjarService *service.LearningPlanningService
}

func NewHandler(
	tpService *service.TPService,
	atpService *service.LearningPlanningService,
	modulAjarService *service.LearningPlanningService,
) *Handler {
	return &Handler{
		tpService:    tpService,
		atpService:   atpService,
		modulAjarService: modulAjarService,
	}
}

// TP handlers
func (h *Handler) CreateTPSet(c *gin.Context) {
	ctx := context.Background()
	authCtx := middleware.GetAuthContext(c)
	
	var req domain.CreateTPSetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "Invalid request")
		return
	}

	tpSet, err := h.tpService.CreateTPSet(ctx, &req, authCtx.UserID)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, tpSet)
}

func (h *Handler) ListTPSets(c *gin.Context) {
	ctx := context.Background()
	
	var cpID *string
	if id := c.Query("cp_id"); id != "" {
		cpID = &id
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	sets, total, err := h.tpService.ListTPSets(ctx, cpID, nil, page, pageSize)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, gin.H{"tp_sets": sets, "total": total, "page": page, "page_size": pageSize})
}

func (h *Handler) GetTPSet(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")
	
	set, err := h.tpService.GetTPSet(ctx, id)
	if err != nil {
		response.Error(c, 404, "Not found")
		return
	}

	response.Success(c, set)
}

func (h *Handler) ApproveTPSet(c *gin.Context) {
	ctx := context.Background()
	authCtx := middleware.GetAuthContext(c)
	id := c.Param("id")
	
	err := h.tpService.ApproveTPSet(ctx, id, authCtx.UserID)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "TP set approved"})
}

// ATP handlers
func (h *Handler) CreateATPSet(c *gin.Context) {
	ctx := context.Background()
	authCtx := middleware.GetAuthContext(c)
	
	var req domain.CreateATPSetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "Invalid request")
		return
	}

	atpSet, err := h.atpService.CreateATPSet(ctx, &req, authCtx.UserID)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, atpSet)
}

func (h *Handler) ListATPSets(c *gin.Context) {
	ctx := context.Background()
	
	var tpSetID *string
	if id := c.Query("tp_set_id"); id != "" {
		tpSetID = &id
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	sets, total, err := h.atpService.ListATPSets(ctx, tpSetID, nil, page, pageSize)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, gin.H{"atp_sets": sets, "total": total, "page": page, "page_size": pageSize})
}

// Modul Ajar handlers
func (h *Handler) CreateModulAjarSet(c *gin.Context) {
	ctx := context.Background()
	authCtx := middleware.GetAuthContext(c)
	
	var req domain.CreateModulAjarSetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "Invalid request")
		return
	}

	set, err := h.modulAjarService.CreateModulAjarSet(ctx, &req, authCtx.UserID)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, set)
}

func (h *Handler) ListModulAjarSets(c *gin.Context) {
	ctx := context.Background()
	
	var atpSetID *string
	if id := c.Query("atp_set_id"); id != "" {
		atpSetID = &id
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	sets, total, err := h.modulAjarService.ListModulAjarSets(ctx, atpSetID, nil, page, pageSize)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, gin.H{"modul_ajar_sets": sets, "total": total, "page": page, "page_size": pageSize})
}

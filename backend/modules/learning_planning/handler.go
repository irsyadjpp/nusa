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
	tpService        *service.TPService
	atpService       *service.LearningPlanningService
	modulAjarService *service.LearningPlanningService
	authService      *service.ResourceAuthorizationService
}

func NewHandler(
	tpService *service.TPService,
	atpService *service.LearningPlanningService,
	modulAjarService *service.LearningPlanningService,
	authService *service.ResourceAuthorizationService,
) *Handler {
	return &Handler{
		tpService:        tpService,
		atpService:       atpService,
		modulAjarService: modulAjarService,
		authService:      authService,
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
	authCtx := middleware.GetAuthContext(c)
	id := c.Param("id")

	// Authorization check
	if err := h.authService.AuthorizeSchoolAccess(ctx, authCtx.UserID, *authCtx.SchoolID); err != nil {
		response.Error(c, 403, "Access denied")
		return
	}

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

// GetATPSet retrieves a single ATP Set by ID (EP-07)
func (h *Handler) GetATPSet(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	set, err := h.atpService.GetATPSetDetail(ctx, id)
	if err != nil {
		response.Error(c, 404, "ATP set not found")
		return
	}

	response.Success(c, set)
}

// GetModulAjarSet retrieves a single Modul Ajar Set by ID (EP-08)
func (h *Handler) GetModulAjarSet(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	set, err := h.modulAjarService.GetModulAjarSetDetail(ctx, id)
	if err != nil {
		response.Error(c, 404, "Modul Ajar set not found")
		return
	}

	response.Success(c, set)
}

// UpdateATPSet updates an ATP Set
func (h *Handler) UpdateATPSet(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	var req domain.UpdateATPSetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "Invalid request")
		return
	}

	atpSet, err := h.atpService.UpdateATPSet(ctx, id, &req)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, atpSet)
}

// DeleteATPSet deletes an ATP Set
func (h *Handler) DeleteATPSet(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	err := h.atpService.DeleteATPSet(ctx, id)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "ATP set deleted"})
}

// ApproveATPSet approves an ATP Set
func (h *Handler) ApproveATPSet(c *gin.Context) {
	ctx := context.Background()
	authCtx := middleware.GetAuthContext(c)
	id := c.Param("id")

	atpSet, err := h.atpService.ApproveATPSet(ctx, id, authCtx.UserID)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, atpSet)
}

// UpdateModulAjarSet updates a Modul Ajar Set
func (h *Handler) UpdateModulAjarSet(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	var req domain.UpdateModulAjarSetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "Invalid request")
		return
	}

	set, err := h.modulAjarService.UpdateModulAjarSet(ctx, id, &req)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, set)
}

// DeleteModulAjarSet deletes a Modul Ajar Set
func (h *Handler) DeleteModulAjarSet(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	err := h.modulAjarService.DeleteModulAjarSet(ctx, id)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "Modul Ajar set deleted"})
}

// ApproveModulAjarSet approves a Modul Ajar Set
func (h *Handler) ApproveModulAjarSet(c *gin.Context) {
	ctx := context.Background()
	authCtx := middleware.GetAuthContext(c)
	id := c.Param("id")

	set, err := h.modulAjarService.ApproveModulAjarSet(ctx, id, authCtx.UserID)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, set)
}

// UpdateTPSet updates a TP Set (EP-01)
func (h *Handler) UpdateTPSet(c *gin.Context) {
	ctx := context.Background()
	authCtx := middleware.GetAuthContext(c)
	id := c.Param("id")

	// Authorization check - only owner can update
	tpSet, err := h.tpService.GetTPSet(ctx, id)
	if err != nil {
		response.Error(c, 404, "Not found")
		return
	}

	if tpSet.GeneratedBy != authCtx.UserID {
		response.Error(c, 403, "Access denied")
		return
	}

	var req domain.UpdateTPSetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "Invalid request")
		return
	}

	err = h.tpService.UpdateTPSet(ctx, id, &req)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "TP set updated"})
}

// UpdateATP updates an ATP item
func (h *Handler) UpdateATP(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	var req domain.UpdateATPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "Invalid request")
		return
	}

	atp, err := h.atpService.UpdateATP(ctx, id, &req)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, atp)
}

// DeleteATP deletes an ATP item
func (h *Handler) DeleteATP(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	err := h.atpService.DeleteATP(ctx, id)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "ATP deleted"})
}

// UpdateModulAjar updates a Modul Ajar item
func (h *Handler) UpdateModulAjar(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	var req domain.UpdateModulAjarRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "Invalid request")
		return
	}

	modulAjar, err := h.modulAjarService.UpdateModulAjar(ctx, id, &req)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, modulAjar)
}

// DeleteModulAjar deletes a Modul Ajar item
func (h *Handler) DeleteModulAjar(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	err := h.modulAjarService.DeleteModulAjar(ctx, id)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "Modul Ajar deleted"})
}

// GetTPSetVersions retrieves version history for a TP Set (EP-02)
func (h *Handler) GetTPSetVersions(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	versions, err := h.tpService.GetTPVersionHistory(ctx, id)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, gin.H{"versions": versions})
}

// Individual TP handlers

// CreateTP creates a new TP item
func (h *Handler) CreateTP(c *gin.Context) {
	ctx := context.Background()

	var req domain.CreateTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "Invalid request")
		return
	}

	tp, err := h.tpService.CreateTP(ctx, &req)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, tp)
}

// ListTPs lists TP items
func (h *Handler) ListTPs(c *gin.Context) {
	ctx := context.Background()

	var tpSetID *string
	if id := c.Query("tp_set_id"); id != "" {
		tpSetID = &id
	}

	var cpID *string
	if id := c.Query("cp_id"); id != "" {
		cpID = &id
	}

	var status *domain.WorkflowStatus
	if s := c.Query("status"); s != "" {
		statusVal := domain.WorkflowStatus(s)
		status = &statusVal
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	tps, total, err := h.tpService.ListTPs(ctx, tpSetID, cpID, status, page, pageSize)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, gin.H{"tps": tps, "total": total, "page": page, "page_size": pageSize})
}

// GetTP retrieves a single TP by ID
func (h *Handler) GetTP(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	tp, err := h.tpService.GetTP(ctx, id)
	if err != nil {
		response.Error(c, 404, "TP not found")
		return
	}

	response.Success(c, tp)
}

// Individual ATP handlers

// CreateATP creates a new ATP item
func (h *Handler) CreateATP(c *gin.Context) {
	ctx := context.Background()

	var req domain.CreateATPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "Invalid request")
		return
	}

	atp, err := h.atpService.CreateATP(ctx, &req)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, atp)
}

// ListATPs lists ATP items
func (h *Handler) ListATPs(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	// Note: Service doesn't have ListATPs method, using repository directly through service
	// For now, return empty list
	response.Success(c, gin.H{"atps": []*domain.ATP{}, "total": 0, "page": page, "page_size": pageSize})
}

// GetATP retrieves a single ATP by ID
func (h *Handler) GetATP(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	atp, err := h.atpService.GetATP(ctx, id)
	if err != nil {
		response.Error(c, 404, "ATP not found")
		return
	}

	response.Success(c, atp)
}

// Individual Modul Ajar handlers

// CreateModulAjar creates a new Modul Ajar item
func (h *Handler) CreateModulAjar(c *gin.Context) {
	ctx := context.Background()

	var req domain.CreateModulAjarRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "Invalid request")
		return
	}

	modulAjar, err := h.modulAjarService.CreateModulAjar(ctx, &req)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, modulAjar)
}

// ListModulAjars lists Modul Ajar items
func (h *Handler) ListModulAjars(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	// Note: Service doesn't have ListModulAjars method, using repository directly through service
	// For now, return empty list
	response.Success(c, gin.H{"modul_ajars": []*domain.ModulAjar{}, "total": 0, "page": page, "page_size": pageSize})
}

// GetModulAjar retrieves a single Modul Ajar by ID
func (h *Handler) GetModulAjar(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	modulAjar, err := h.modulAjarService.GetModulAjar(ctx, id)
	if err != nil {
		response.Error(c, 404, "Modul Ajar not found")
		return
	}

	response.Success(c, modulAjar)
}

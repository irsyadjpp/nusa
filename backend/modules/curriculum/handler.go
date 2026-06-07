package curriculum

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
	curriculumService *service.CurriculumService
}

func NewHandler(curriculumService *service.CurriculumService) *Handler {
	return &Handler{curriculumService: curriculumService}
}

// CurriculumSubject handlers
func (h *Handler) CreateCurriculumSubject(c *gin.Context) {
	var req domain.CreateCurriculumSubjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, map[string]string{"error": "Invalid request"})
		return
	}

	ctx := context.Background()
	subject, err := h.curriculumService.CreateCurriculumSubject(ctx, &req)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, subject)
}

func (h *Handler) ListCurriculumSubjects(c *gin.Context) {
	ctx := context.Background()
	
	var isActive *bool
	if active := c.Query("is_active"); active != "" {
		b, _ := strconv.ParseBool(active)
		isActive = &b
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	subjects, total, err := h.curriculumService.ListCurriculumSubjects(ctx, isActive, page, pageSize)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, gin.H{"subjects": subjects, "total": total, "page": page, "page_size": pageSize})
}

func (h *Handler) GetCurriculumSubject(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")
	
	subject, err := h.curriculumService.GetCurriculumSubject(ctx, id)
	if err != nil {
		response.Error(c, 404, "Not found")
		return
	}

	response.Success(c, subject)
}

// CP handlers
func (h *Handler) ImportCP(c *gin.Context) {
	ctx := context.Background()
	authCtx := middleware.GetAuthContext(c)
	
	var req domain.ImportCPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "Invalid request")
		return
	}

	cps, err := h.curriculumService.ImportCP(ctx, &req, authCtx.UserID)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, gin.H{"imported": len(cps), "cps": cps})
}

func (h *Handler) ListCPs(c *gin.Context) {
	ctx := context.Background()
	
	var subjectID, phaseID, elementID *string
	if s := c.Query("subject_id"); s != "" {
		subjectID = &s
	}
	if p := c.Query("phase_id"); p != "" {
		phaseID = &p
	}
	if e := c.Query("element_id"); e != "" {
		elementID = &e
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	cps, total, err := h.curriculumService.ListCPs(ctx, subjectID, phaseID, elementID, nil, nil, page, pageSize)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, gin.H{"cps": cps, "total": total, "page": page, "page_size": pageSize})
}

func (h *Handler) GetCP(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")
	
	cp, err := h.curriculumService.GetCP(ctx, id)
	if err != nil {
		response.Error(c, 404, "Not found")
		return
	}

	response.Success(c, cp)
}

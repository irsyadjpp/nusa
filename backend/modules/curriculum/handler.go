package curriculum

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/handler/dto"
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
	var req dto.CreateCurriculumSubjectRequest
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

func (h *Handler) UpdateCurriculumSubject(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	var req dto.UpdateCurriculumSubjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, map[string]string{"error": "Invalid request"})
		return
	}

	subject, err := h.curriculumService.UpdateCurriculumSubject(ctx, id, &req)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, subject)
}

func (h *Handler) DeleteCurriculumSubject(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	if err := h.curriculumService.DeleteCurriculumSubject(ctx, id); err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "Curriculum subject deleted successfully"})
}

// CurriculumPhase handlers
func (h *Handler) CreateCurriculumPhase(c *gin.Context) {
	var req dto.CreateCurriculumPhaseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, map[string]string{"error": "Invalid request"})
		return
	}

	ctx := context.Background()
	phase, err := h.curriculumService.CreateCurriculumPhase(ctx, &req)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, phase)
}

func (h *Handler) ListCurriculumPhases(c *gin.Context) {
	ctx := context.Background()

	var isActive *bool
	if active := c.Query("is_active"); active != "" {
		b, _ := strconv.ParseBool(active)
		isActive = &b
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	phases, total, err := h.curriculumService.ListCurriculumPhases(ctx, isActive, page, pageSize)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, gin.H{"phases": phases, "total": total, "page": page, "page_size": pageSize})
}

func (h *Handler) GetCurriculumPhase(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	phase, err := h.curriculumService.GetCurriculumPhase(ctx, id)
	if err != nil {
		response.Error(c, 404, "Not found")
		return
	}

	response.Success(c, phase)
}

func (h *Handler) UpdateCurriculumPhase(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	var req dto.UpdateCurriculumPhaseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, map[string]string{"error": "Invalid request"})
		return
	}

	phase, err := h.curriculumService.UpdateCurriculumPhase(ctx, id, &req)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, phase)
}

func (h *Handler) DeleteCurriculumPhase(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	if err := h.curriculumService.DeleteCurriculumPhase(ctx, id); err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "Curriculum phase deleted successfully"})
}

// CurriculumElement handlers
func (h *Handler) CreateCurriculumElement(c *gin.Context) {
	var req dto.CreateCurriculumElementRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, map[string]string{"error": "Invalid request"})
		return
	}

	ctx := context.Background()
	element, err := h.curriculumService.CreateCurriculumElement(ctx, &req)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, element)
}

func (h *Handler) ListCurriculumElements(c *gin.Context) {
	ctx := context.Background()

	var subjectID, phaseID *string
	if s := c.Query("subject_id"); s != "" {
		subjectID = &s
	}
	if p := c.Query("phase_id"); p != "" {
		phaseID = &p
	}

	var isActive *bool
	if active := c.Query("is_active"); active != "" {
		b, _ := strconv.ParseBool(active)
		isActive = &b
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	elements, total, err := h.curriculumService.ListCurriculumElements(ctx, subjectID, phaseID, isActive, page, pageSize)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, gin.H{"elements": elements, "total": total, "page": page, "page_size": pageSize})
}

func (h *Handler) GetCurriculumElement(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	element, err := h.curriculumService.GetCurriculumElement(ctx, id)
	if err != nil {
		response.Error(c, 404, "Not found")
		return
	}

	response.Success(c, element)
}

func (h *Handler) UpdateCurriculumElement(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	var req dto.UpdateCurriculumElementRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, map[string]string{"error": "Invalid request"})
		return
	}

	element, err := h.curriculumService.UpdateCurriculumElement(ctx, id, &req)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, element)
}

func (h *Handler) DeleteCurriculumElement(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	if err := h.curriculumService.DeleteCurriculumElement(ctx, id); err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "Curriculum element deleted successfully"})
}

// CurriculumSubelement handlers
func (h *Handler) CreateCurriculumSubelement(c *gin.Context) {
	var req dto.CreateCurriculumSubelementRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, map[string]string{"error": "Invalid request"})
		return
	}

	ctx := context.Background()
	subelement, err := h.curriculumService.CreateCurriculumSubelement(ctx, &req)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, subelement)
}

func (h *Handler) ListCurriculumSubelements(c *gin.Context) {
	ctx := context.Background()

	var elementID *string
	if e := c.Query("element_id"); e != "" {
		elementID = &e
	}

	var isActive *bool
	if active := c.Query("is_active"); active != "" {
		b, _ := strconv.ParseBool(active)
		isActive = &b
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	subelements, total, err := h.curriculumService.ListCurriculumSubelements(ctx, elementID, isActive, page, pageSize)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, gin.H{"subelements": subelements, "total": total, "page": page, "page_size": pageSize})
}

func (h *Handler) GetCurriculumSubelement(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	subelement, err := h.curriculumService.GetCurriculumSubelement(ctx, id)
	if err != nil {
		response.Error(c, 404, "Not found")
		return
	}

	response.Success(c, subelement)
}

func (h *Handler) UpdateCurriculumSubelement(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	var req dto.UpdateCurriculumSubelementRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, map[string]string{"error": "Invalid request"})
		return
	}

	subelement, err := h.curriculumService.UpdateCurriculumSubelement(ctx, id, &req)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, subelement)
}

func (h *Handler) DeleteCurriculumSubelement(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	if err := h.curriculumService.DeleteCurriculumSubelement(ctx, id); err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "Curriculum subelement deleted successfully"})
}

// CP handlers
func (h *Handler) ImportCP(c *gin.Context) {
	ctx := context.Background()
	authCtx := middleware.GetAuthContext(c)

	var req dto.ImportCPRequest
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

	var subjectID, phaseID *string
	if s := c.Query("subject_id"); s != "" {
		subjectID = &s
	}
	if p := c.Query("phase_id"); p != "" {
		phaseID = &p
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	cps, total, err := h.curriculumService.ListCPs(ctx, subjectID, phaseID, page, pageSize)
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

func (h *Handler) CreateCP(c *gin.Context) {
	ctx := context.Background()

	var req dto.CreateCPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, map[string]string{"error": "Invalid request"})
		return
	}

	cp, err := h.curriculumService.CreateCP(ctx, &req)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, cp)
}

func (h *Handler) UpdateCP(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	var req dto.UpdateCPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, map[string]string{"error": "Invalid request"})
		return
	}

	cp, err := h.curriculumService.UpdateCP(ctx, id, &req)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, cp)
}

func (h *Handler) DeleteCP(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	if err := h.curriculumService.DeleteCP(ctx, id); err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "CP deleted successfully"})
}

func (h *Handler) ExportCPs(c *gin.Context) {
	ctx := context.Background()

	var subjectID, phaseID *string
	if s := c.Query("subject_id"); s != "" {
		subjectID = &s
	}
	if p := c.Query("phase_id"); p != "" {
		phaseID = &p
	}

	format := c.DefaultQuery("format", "csv")

	cps, _, err := h.curriculumService.ListCPs(ctx, subjectID, phaseID, 1, 10000)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	if format == "csv" {
		h.exportCPsAsCSV(c, cps)
	} else {
		response.Error(c, 400, "Unsupported format")
	}
}

func (h *Handler) exportCPsAsCSV(c *gin.Context, cps []*domain.CP) {
	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", "attachment; filename=cp_export.csv")

	csv := "ID,SubjectID,PhaseID,ElementID,SubelementID,Code,Description,CompetencyCode,Version,IsActive\n"
	for _, cp := range cps {
		csv += fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s,%t\n",
			cp.ID, cp.SubjectID, cp.PhaseID, cp.ElementID, cp.SubelementID,
			cp.Code, escapeCSV(cp.Description), stringOrEmpty(cp.CompetencyCode),
			cp.Version, cp.IsActive)
	}

	c.String(200, csv)
}

func escapeCSV(s string) string {
	if strings.Contains(s, ",") || strings.Contains(s, "\"") || strings.Contains(s, "\n") {
		return fmt.Sprintf("\"%s\"", strings.ReplaceAll(s, "\"", "\"\""))
	}
	return s
}

func stringOrEmpty(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

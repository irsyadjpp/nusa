package reporting

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
	reportingService *service.ReportingService
}

func NewHandler(reportingService *service.ReportingService) *Handler {
	return &Handler{reportingService: reportingService}
}

// Narrative Report handlers
func (h *Handler) CreateNarrativeReport(c *gin.Context) {
	ctx := context.Background()
	authCtx := middleware.GetAuthContext(c)
	
	var req domain.CreateNarrativeReportRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "Invalid request")
		return
	}

	report, err := h.reportingService.CreateNarrativeReport(ctx, &req, authCtx.UserID)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, report)
}

func (h *Handler) ListNarrativeReports(c *gin.Context) {
	ctx := context.Background()
	
	var studentID *string
	if id := c.Query("student_id"); id != "" {
		studentID = &id
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	reports, total, err := h.reportingService.ListNarrativeReports(ctx, studentID, nil, nil, nil, page, pageSize)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, gin.H{"reports": reports, "total": total, "page": page, "page_size": pageSize})
}

func (h *Handler) GetNarrativeReport(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")
	
	report, err := h.reportingService.GetNarrativeReport(ctx, id)
	if err != nil {
		response.Error(c, 404, "Not found")
		return
	}

	response.Success(c, report)
}

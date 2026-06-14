package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nusa/backend/internal/application"
	"github.com/nusa/backend/internal/handler/dto"
	"github.com/nusa/backend/internal/middleware"
)

// ICurriculumGovernanceApplicationService defines the interface for Curriculum Governance application service
type ICurriculumGovernanceApplicationService interface {
	// Subject Category
	CreateSubjectCategory(ctx context.Context, cmd *application.CreateSubjectCategoryCommand) (*application.CreateSubjectCategoryResponse, error)
	UpdateSubjectCategory(ctx context.Context, cmd *application.UpdateSubjectCategoryCommand) (*application.UpdateSubjectCategoryResponse, error)
	DeleteSubjectCategory(ctx context.Context, cmd *application.DeleteSubjectCategoryCommand) (*application.DeleteSubjectCategoryResponse, error)
	ListSubjectCategories(ctx context.Context, cmd *application.ListSubjectCategoriesCommand) (*application.ListSubjectCategoriesResponse, error)

	// Graduate Profile Dimension
	CreateGraduateProfileDimension(ctx context.Context, cmd *application.CreateGraduateProfileDimensionCommand) (*application.CreateGraduateProfileDimensionResponse, error)
	UpdateGraduateProfileDimension(ctx context.Context, cmd *application.UpdateGraduateProfileDimensionCommand) (*application.UpdateGraduateProfileDimensionResponse, error)
	DeleteGraduateProfileDimension(ctx context.Context, cmd *application.DeleteGraduateProfileDimensionCommand) (*application.DeleteGraduateProfileDimensionResponse, error)
	ListGraduateProfileDimensions(ctx context.Context, cmd *application.ListGraduateProfileDimensionsCommand) (*application.ListGraduateProfileDimensionsResponse, error)

	// CP Alignment
	CreateCPAlignment(ctx context.Context, cmd *application.CreateCPAlignmentCommand) (*application.CreateCPAlignmentResponse, error)
	CreateCPAlignmentBulk(ctx context.Context, cmd *application.CreateCPAlignmentBulkCommand) (*application.CreateCPAlignmentBulkResponse, error)
	UpdateCPAlignment(ctx context.Context, cmd *application.UpdateCPAlignmentCommand) (*application.UpdateCPAlignmentResponse, error)
	DeleteCPAlignment(ctx context.Context, cmd *application.DeleteCPAlignmentCommand) (*application.DeleteCPAlignmentResponse, error)
	GetCPAlignmentsByCurriculumSubject(ctx context.Context, cmd *application.GetCPAlignmentsByCurriculumSubjectCommand) (*application.GetCPAlignmentsByCurriculumSubjectResponse, error)
	GenerateCPAlignmentReport(ctx context.Context, cmd *application.GenerateCPAlignmentReportCommand) (*application.GenerateCPAlignmentReportResponse, error)
	ListCPAlignments(ctx context.Context, cmd *application.ListCPAlignmentsCommand) (*application.ListCPAlignmentsResponse, error)
}

// CurriculumGovernanceHandler handles HTTP requests for Curriculum Governance endpoints
type CurriculumGovernanceHandler struct {
	curriculumGovernanceService ICurriculumGovernanceApplicationService
}

// NewCurriculumGovernanceHandler creates a new Curriculum Governance handler
func NewCurriculumGovernanceHandler(curriculumGovernanceService *application.CurriculumGovernanceApplicationService) *CurriculumGovernanceHandler {
	return &CurriculumGovernanceHandler{
		curriculumGovernanceService: curriculumGovernanceService,
	}
}

// ==================== Subject Category Endpoints ====================

// CreateSubjectCategory creates a new subject category
// POST /subject-categories
func (h *CurriculumGovernanceHandler) CreateSubjectCategory(c *gin.Context) {
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

	var req dto.CreateSubjectCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INVALID_REQUEST",
				Message: err.Error(),
			},
		})
		return
	}

	cmd := &application.CreateSubjectCategoryCommand{
		Code:        req.Code,
		Name:        req.Name,
		Description: req.Description,
		IsMandatory: req.IsMandatory,
		UserID:      authCtx.UserID,
	}

	resp, err := h.curriculumGovernanceService.CreateSubjectCategory(c.Request.Context(), cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusCreated, dto.SubjectCategoryResponse{
		ID: resp.SubjectCategoryID,
	})
}

// UpdateSubjectCategory updates a subject category
// PUT /subject-categories/:id
func (h *CurriculumGovernanceHandler) UpdateSubjectCategory(c *gin.Context) {
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
	if id == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INVALID_REQUEST",
				Message: "Subject Category ID is required",
			},
		})
		return
	}

	var req dto.UpdateSubjectCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INVALID_REQUEST",
				Message: err.Error(),
			},
		})
		return
	}

	cmd := &application.UpdateSubjectCategoryCommand{
		SubjectCategoryID: id,
		Name:              req.Name,
		Description:       req.Description,
		IsMandatory:       req.IsMandatory,
		IsActive:          req.IsActive,
		UserID:            authCtx.UserID,
	}

	resp, err := h.curriculumGovernanceService.UpdateSubjectCategory(c.Request.Context(), cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": resp.Success})
}

// DeleteSubjectCategory deletes a subject category
// DELETE /subject-categories/:id
func (h *CurriculumGovernanceHandler) DeleteSubjectCategory(c *gin.Context) {
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
	if id == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INVALID_REQUEST",
				Message: "Subject Category ID is required",
			},
		})
		return
	}

	cmd := &application.DeleteSubjectCategoryCommand{
		SubjectCategoryID: id,
		UserID:            authCtx.UserID,
	}

	resp, err := h.curriculumGovernanceService.DeleteSubjectCategory(c.Request.Context(), cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": resp.Success})
}

// ListSubjectCategories retrieves all subject categories
// GET /subject-categories?active_only=:active_only
func (h *CurriculumGovernanceHandler) ListSubjectCategories(c *gin.Context) {
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

	activeOnly := c.Query("active_only") == "true"

	cmd := &application.ListSubjectCategoriesCommand{
		ActiveOnly: activeOnly,
		UserID:     authCtx.UserID,
	}

	resp, err := h.curriculumGovernanceService.ListSubjectCategories(c.Request.Context(), cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	// Convert subject categories to DTO
	categoryResponses := make([]dto.SubjectCategoryResponse, len(resp.SubjectCategories))
	for i, sc := range resp.SubjectCategories {
		categoryResponses[i] = dto.SubjectCategoryResponse{
			ID:          sc.ID,
			Code:        sc.Code,
			Name:        sc.Name,
			Description: sc.Description,
			IsMandatory: sc.IsMandatory,
			IsActive:    sc.IsActive,
			CreatedBy:   sc.CreatedBy,
			CreatedAt:   sc.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   sc.UpdatedAt.Format(time.RFC3339),
		}
	}

	c.JSON(http.StatusOK, dto.ListSubjectCategoriesResponse{
		SubjectCategories: categoryResponses,
		Total:             len(resp.SubjectCategories),
	})
}

// ==================== Graduate Profile Dimension Endpoints ====================

// CreateGraduateProfileDimension creates a new graduate profile dimension
// POST /graduate-profile-dimensions
func (h *CurriculumGovernanceHandler) CreateGraduateProfileDimension(c *gin.Context) {
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

	var req dto.CreateGraduateProfileDimensionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INVALID_REQUEST",
				Message: err.Error(),
			},
		})
		return
	}

	cmd := &application.CreateGraduateProfileDimensionCommand{
		Code:           req.Code,
		Name:           req.Name,
		Description:    req.Description,
		SequenceNumber: req.SequenceNumber,
		UserID:         authCtx.UserID,
	}

	resp, err := h.curriculumGovernanceService.CreateGraduateProfileDimension(c.Request.Context(), cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusCreated, dto.GraduateProfileDimensionResponse{
		ID: resp.GraduateProfileDimensionID,
	})
}

// UpdateGraduateProfileDimension updates a graduate profile dimension
// PUT /graduate-profile-dimensions/:id
func (h *CurriculumGovernanceHandler) UpdateGraduateProfileDimension(c *gin.Context) {
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
	if id == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INVALID_REQUEST",
				Message: "Graduate Profile Dimension ID is required",
			},
		})
		return
	}

	var req dto.UpdateGraduateProfileDimensionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INVALID_REQUEST",
				Message: err.Error(),
			},
		})
		return
	}

	cmd := &application.UpdateGraduateProfileDimensionCommand{
		GraduateProfileDimensionID: id,
		Name:                       req.Name,
		Description:                req.Description,
		SequenceNumber:             req.SequenceNumber,
		IsActive:                   req.IsActive,
		UserID:                     authCtx.UserID,
	}

	resp, err := h.curriculumGovernanceService.UpdateGraduateProfileDimension(c.Request.Context(), cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": resp.Success})
}

// DeleteGraduateProfileDimension deletes a graduate profile dimension
// DELETE /graduate-profile-dimensions/:id
func (h *CurriculumGovernanceHandler) DeleteGraduateProfileDimension(c *gin.Context) {
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
	if id == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INVALID_REQUEST",
				Message: "Graduate Profile Dimension ID is required",
			},
		})
		return
	}

	cmd := &application.DeleteGraduateProfileDimensionCommand{
		GraduateProfileDimensionID: id,
		UserID:                     authCtx.UserID,
	}

	resp, err := h.curriculumGovernanceService.DeleteGraduateProfileDimension(c.Request.Context(), cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": resp.Success})
}

// ListGraduateProfileDimensions retrieves all graduate profile dimensions
// GET /graduate-profile-dimensions?active_only=:active_only
func (h *CurriculumGovernanceHandler) ListGraduateProfileDimensions(c *gin.Context) {
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

	activeOnly := c.Query("active_only") == "true"

	cmd := &application.ListGraduateProfileDimensionsCommand{
		ActiveOnly: activeOnly,
		UserID:     authCtx.UserID,
	}

	resp, err := h.curriculumGovernanceService.ListGraduateProfileDimensions(c.Request.Context(), cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	// Convert graduate profile dimensions to DTO
	dimensionResponses := make([]dto.GraduateProfileDimensionResponse, len(resp.GraduateProfileDimensions))
	for i, gpd := range resp.GraduateProfileDimensions {
		dimensionResponses[i] = dto.GraduateProfileDimensionResponse{
			ID:             gpd.ID,
			Code:           gpd.Code,
			Name:           gpd.Name,
			Description:    gpd.Description,
			SequenceNumber: gpd.SequenceNumber,
			IsActive:       gpd.IsActive,
			CreatedBy:      gpd.CreatedBy,
			CreatedAt:      gpd.CreatedAt.Format(time.RFC3339),
			UpdatedAt:      gpd.UpdatedAt.Format(time.RFC3339),
		}
	}

	c.JSON(http.StatusOK, dto.ListGraduateProfileDimensionsResponse{
		GraduateProfileDimensions: dimensionResponses,
		Total:                     len(resp.GraduateProfileDimensions),
	})
}

// ==================== CP Alignment Endpoints ====================

// CreateCPAlignment creates a new CP alignment
// POST /cp-alignments
func (h *CurriculumGovernanceHandler) CreateCPAlignment(c *gin.Context) {
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

	var req dto.CreateCPAlignmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INVALID_REQUEST",
				Message: err.Error(),
			},
		})
		return
	}

	cmd := &application.CreateCPAlignmentCommand{
		CurriculumSubjectID:        req.CurriculumSubjectID,
		GraduateProfileDimensionID: req.GraduateProfileDimensionID,
		AlignmentDescription:       req.AlignmentDescription,
		UserID:                     authCtx.UserID,
	}

	resp, err := h.curriculumGovernanceService.CreateCPAlignment(c.Request.Context(), cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusCreated, dto.CPAlignmentResponse{
		ID: resp.CPAlignmentID,
	})
}

// CreateCPAlignmentBulk creates multiple CP alignments for a curriculum subject
// POST /cp-alignments/bulk
func (h *CurriculumGovernanceHandler) CreateCPAlignmentBulk(c *gin.Context) {
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

	var req dto.CreateCPAlignmentBulkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INVALID_REQUEST",
				Message: err.Error(),
			},
		})
		return
	}

	cmd := &application.CreateCPAlignmentBulkCommand{
		CurriculumSubjectID:  req.CurriculumSubjectID,
		AlignmentIDs:         req.AlignmentIDs,
		AlignmentDescription: req.AlignmentDescription,
		UserID:               authCtx.UserID,
	}

	resp, err := h.curriculumGovernanceService.CreateCPAlignmentBulk(c.Request.Context(), cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"cp_alignment_ids": resp.CPAlignmentIDs,
	})
}

// UpdateCPAlignment updates a CP alignment
// PUT /cp-alignments/:id
func (h *CurriculumGovernanceHandler) UpdateCPAlignment(c *gin.Context) {
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
	if id == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INVALID_REQUEST",
				Message: "CP Alignment ID is required",
			},
		})
		return
	}

	var req dto.UpdateCPAlignmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INVALID_REQUEST",
				Message: err.Error(),
			},
		})
		return
	}

	cmd := &application.UpdateCPAlignmentCommand{
		CPAlignmentID:        id,
		AlignmentDescription: req.AlignmentDescription,
		IsActive:             req.IsActive,
		UserID:               authCtx.UserID,
	}

	resp, err := h.curriculumGovernanceService.UpdateCPAlignment(c.Request.Context(), cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": resp.Success})
}

// DeleteCPAlignment deletes a CP alignment
// DELETE /cp-alignments/:id
func (h *CurriculumGovernanceHandler) DeleteCPAlignment(c *gin.Context) {
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
	if id == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INVALID_REQUEST",
				Message: "CP Alignment ID is required",
			},
		})
		return
	}

	cmd := &application.DeleteCPAlignmentCommand{
		CPAlignmentID: id,
		UserID:        authCtx.UserID,
	}

	resp, err := h.curriculumGovernanceService.DeleteCPAlignment(c.Request.Context(), cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": resp.Success})
}

// GetCPAlignmentsByCurriculumSubject retrieves CP alignments for a curriculum subject
// GET /cp-alignments?curriculum_subject_id=:curriculum_subject_id
func (h *CurriculumGovernanceHandler) GetCPAlignmentsByCurriculumSubject(c *gin.Context) {
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

	curriculumSubjectID := c.Query("curriculum_subject_id")
	if curriculumSubjectID == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INVALID_REQUEST",
				Message: "Curriculum Subject ID is required",
			},
		})
		return
	}

	cmd := &application.GetCPAlignmentsByCurriculumSubjectCommand{
		CurriculumSubjectID: curriculumSubjectID,
		UserID:              authCtx.UserID,
	}

	resp, err := h.curriculumGovernanceService.GetCPAlignmentsByCurriculumSubject(c.Request.Context(), cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	// Convert CP alignments to DTO
	alignmentResponses := make([]dto.CPAlignmentResponse, len(resp.CPAlignments))
	for i, cpa := range resp.CPAlignments {
		alignmentResponses[i] = dto.CPAlignmentResponse{
			ID:                         cpa.ID,
			CurriculumSubjectID:        cpa.CurriculumSubjectID,
			GraduateProfileDimensionID: cpa.GraduateProfileDimensionID,
			AlignmentDescription:       cpa.AlignmentDescription,
			IsActive:                   cpa.IsActive,
			CreatedBy:                  cpa.CreatedBy,
			CreatedAt:                  cpa.CreatedAt.Format(time.RFC3339),
			UpdatedAt:                  cpa.UpdatedAt.Format(time.RFC3339),
		}
	}

	c.JSON(http.StatusOK, dto.ListCPAlignmentsResponse{
		CPAlignments: alignmentResponses,
		Total:        len(resp.CPAlignments),
	})
}

// GenerateCPAlignmentReport generates a CP alignment report
// GET /cp-alignments/report
func (h *CurriculumGovernanceHandler) GenerateCPAlignmentReport(c *gin.Context) {
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

	cmd := &application.GenerateCPAlignmentReportCommand{
		UserID: authCtx.UserID,
	}

	resp, err := h.curriculumGovernanceService.GenerateCPAlignmentReport(c.Request.Context(), cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	// Convert reports to DTO
	reportResponses := make([]dto.CPAlignmentReportResponse, len(resp.Reports))
	for i, report := range resp.Reports {
		reportResponses[i] = dto.CPAlignmentReportResponse{
			GraduateProfileDimensionID:   report.GraduateProfileDimensionID,
			GraduateProfileDimensionName: report.GraduateProfileDimensionName,
			TotalCPCount:                 report.TotalCPCount,
			AlignedCPCount:               report.AlignedCPCount,
			CoveragePercentage:           report.CoveragePercentage,
			MeetsThreshold:               report.MeetsThreshold,
		}
	}

	c.JSON(http.StatusOK, dto.GenerateCPAlignmentReportResponse{
		Reports: reportResponses,
	})
}

// ListCPAlignments retrieves all CP alignments
// GET /cp-alignments
func (h *CurriculumGovernanceHandler) ListCPAlignments(c *gin.Context) {
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

	cmd := &application.ListCPAlignmentsCommand{
		UserID: authCtx.UserID,
	}

	resp, err := h.curriculumGovernanceService.ListCPAlignments(c.Request.Context(), cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	// Convert CP alignments to DTO
	alignmentResponses := make([]dto.CPAlignmentResponse, len(resp.CPAlignments))
	for i, cpa := range resp.CPAlignments {
		alignmentResponses[i] = dto.CPAlignmentResponse{
			ID:                         cpa.ID,
			CurriculumSubjectID:        cpa.CurriculumSubjectID,
			GraduateProfileDimensionID: cpa.GraduateProfileDimensionID,
			AlignmentDescription:       cpa.AlignmentDescription,
			IsActive:                   cpa.IsActive,
			CreatedBy:                  cpa.CreatedBy,
			CreatedAt:                  cpa.CreatedAt.Format(time.RFC3339),
			UpdatedAt:                  cpa.UpdatedAt.Format(time.RFC3339),
		}
	}

	c.JSON(http.StatusOK, dto.ListCPAlignmentsResponse{
		CPAlignments: alignmentResponses,
		Total:        len(resp.CPAlignments),
	})
}

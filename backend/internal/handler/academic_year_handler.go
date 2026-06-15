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

// IAcademicYearApplicationService defines the interface for Academic Year application service
type IAcademicYearApplicationService interface {
	CreateAcademicYear(ctx context.Context, cmd *application.CreateAcademicYearCommand) (*application.CreateAcademicYearResponse, error)
	UpdateAcademicYear(ctx context.Context, cmd *application.UpdateAcademicYearCommand) (*application.UpdateAcademicYearResponse, error)
	ActivateAcademicYear(ctx context.Context, cmd *application.ActivateAcademicYearCommand) (*application.ActivateAcademicYearResponse, error)
	ArchiveAcademicYear(ctx context.Context, cmd *application.ArchiveAcademicYearCommand) (*application.ArchiveAcademicYearResponse, error)
	GetAcademicYear(ctx context.Context, cmd *application.GetAcademicYearCommand) (*application.GetAcademicYearResponse, error)
	ListAcademicYears(ctx context.Context, cmd *application.ListAcademicYearsCommand) (*application.ListAcademicYearsResponse, error)
}

// AcademicYearHandler handles HTTP requests for Academic Year endpoints
type AcademicYearHandler struct {
	academicYearService IAcademicYearApplicationService
}

// NewAcademicYearHandler creates a new Academic Year handler
func NewAcademicYearHandler(academicYearService *application.AcademicYearApplicationService) *AcademicYearHandler {
	return &AcademicYearHandler{
		academicYearService: academicYearService,
	}
}

// CreateAcademicYear creates a new academic year
// POST /academic-years
func (h *AcademicYearHandler) CreateAcademicYear(c *gin.Context) {
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

	var req dto.CreateAcademicYearRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INVALID_REQUEST",
				Message: err.Error(),
			},
		})
		return
	}

	cmd := &application.CreateAcademicYearCommand{
		SchoolID:    req.SchoolID,
		Name:        req.Name,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		Description: req.Description,
		UserID:      authCtx.UserID,
	}

	resp, err := h.academicYearService.CreateAcademicYear(c.Request.Context(), cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusCreated, dto.AcademicYearResponse{
		ID:     resp.AcademicYearID,
		Status: string(resp.Status),
	})
}

// UpdateAcademicYear updates an academic year
// PUT /academic-years/:id
func (h *AcademicYearHandler) UpdateAcademicYear(c *gin.Context) {
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
				Message: "Academic year ID is required",
			},
		})
		return
	}

	var req dto.UpdateAcademicYearRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INVALID_REQUEST",
				Message: err.Error(),
			},
		})
		return
	}

	cmd := &application.UpdateAcademicYearCommand{
		AcademicYearID: id,
		Name:           req.Name,
		StartDate:      req.StartDate,
		EndDate:        req.EndDate,
		UserID:         authCtx.UserID,
	}

	resp, err := h.academicYearService.UpdateAcademicYear(c.Request.Context(), cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, dto.AcademicYearResponse{
		ID:     resp.AcademicYearID,
		Status: string(resp.Status),
	})
}

// ActivateAcademicYear activates an academic year
// POST /academic-years/:id/activate
func (h *AcademicYearHandler) ActivateAcademicYear(c *gin.Context) {
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
				Message: "Academic year ID is required",
			},
		})
		return
	}

	cmd := &application.ActivateAcademicYearCommand{
		AcademicYearID: id,
		UserID:         authCtx.UserID,
	}

	resp, err := h.academicYearService.ActivateAcademicYear(c.Request.Context(), cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, dto.AcademicYearResponse{
		ID:     resp.AcademicYearID,
		Status: string(resp.Status),
	})
}

// ArchiveAcademicYear archives an academic year
// POST /academic-years/:id/archive
func (h *AcademicYearHandler) ArchiveAcademicYear(c *gin.Context) {
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
				Message: "Academic year ID is required",
			},
		})
		return
	}

	cmd := &application.ArchiveAcademicYearCommand{
		AcademicYearID: id,
		UserID:         authCtx.UserID,
	}

	resp, err := h.academicYearService.ArchiveAcademicYear(c.Request.Context(), cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, dto.AcademicYearResponse{
		ID:     resp.AcademicYearID,
		Status: string(resp.Status),
	})
}

// GetAcademicYear retrieves an academic year with its semesters
// GET /academic-years/:id
func (h *AcademicYearHandler) GetAcademicYear(c *gin.Context) {
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
				Message: "Academic year ID is required",
			},
		})
		return
	}

	cmd := &application.GetAcademicYearCommand{
		AcademicYearID: id,
		UserID:         authCtx.UserID,
	}

	resp, err := h.academicYearService.GetAcademicYear(c.Request.Context(), cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	// Convert semesters to DTO
	semesterResponses := make([]dto.SemesterResponse, len(resp.Semesters))
	for i, sem := range resp.Semesters {
		semesterResponses[i] = dto.SemesterResponse{
			ID:             sem.ID,
			AcademicYearID: sem.AcademicYearID,
			Type:           string(sem.Type),
			Name:           sem.Name,
			StartDate:      sem.StartDate.Format(time.RFC3339),
			EndDate:        sem.EndDate.Format(time.RFC3339),
			Status:         string(sem.Status),
			SequenceNumber: sem.SequenceNumber,
			CreatedBy:      sem.CreatedBy,
			CreatedAt:      sem.CreatedAt.Format(time.RFC3339),
			UpdatedAt:      sem.UpdatedAt.Format(time.RFC3339),
		}
	}

	c.JSON(http.StatusOK, dto.AcademicYearWithSemestersResponse{
		ID:        resp.ID,
		SchoolID:  resp.SchoolID,
		Name:      resp.Name,
		StartDate: resp.StartDate.Format(time.RFC3339),
		EndDate:   resp.EndDate.Format(time.RFC3339),
		Status:    string(resp.Status),
		CreatedBy: resp.CreatedBy,
		CreatedAt: resp.CreatedAt.Format(time.RFC3339),
		UpdatedAt: resp.UpdatedAt.Format(time.RFC3339),
		Semesters: semesterResponses,
	})
}

// ListAcademicYears retrieves all academic years for a school
// GET /academic-years?school_id=:school_id
func (h *AcademicYearHandler) ListAcademicYears(c *gin.Context) {
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

	schoolID := c.Query("school_id")
	if schoolID == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INVALID_REQUEST",
				Message: "School ID is required",
			},
		})
		return
	}

	cmd := &application.ListAcademicYearsCommand{
		SchoolID: schoolID,
		UserID:   authCtx.UserID,
	}

	resp, err := h.academicYearService.ListAcademicYears(c.Request.Context(), cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	// Convert academic years to DTO
	yearResponses := make([]dto.AcademicYearResponse, len(resp.AcademicYears))
	for i, ay := range resp.AcademicYears {
		yearResponses[i] = dto.AcademicYearResponse{
			ID:        ay.ID,
			SchoolID:  ay.SchoolID,
			Name:      ay.Name,
			StartDate: ay.StartDate.Format(time.RFC3339),
			EndDate:   ay.EndDate.Format(time.RFC3339),
			Status:    string(ay.Status),
			CreatedBy: ay.CreatedBy,
			CreatedAt: ay.CreatedAt.Format(time.RFC3339),
			UpdatedAt: ay.UpdatedAt.Format(time.RFC3339),
		}
	}

	c.JSON(http.StatusOK, dto.ListAcademicYearsResponse{
		AcademicYears: yearResponses,
		Total:         len(resp.AcademicYears),
	})
}

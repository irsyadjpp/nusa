package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nusa/backend/internal/application"
	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/handler/dto"
	"github.com/nusa/backend/internal/middleware"
)

// ISemesterApplicationService defines the interface for Semester application service
type ISemesterApplicationService interface {
	CreateSemester(ctx context.Context, cmd *application.CreateSemesterCommand) (*application.CreateSemesterResponse, error)
	UpdateSemester(ctx context.Context, cmd *application.UpdateSemesterCommand) (*application.UpdateSemesterResponse, error)
	DeleteSemester(ctx context.Context, cmd *application.DeleteSemesterCommand) (*application.DeleteSemesterResponse, error)
	GetSemester(ctx context.Context, cmd *application.GetSemesterCommand) (*application.GetSemesterResponse, error)
	ListSemesters(ctx context.Context, cmd *application.ListSemestersCommand) (*application.ListSemestersResponse, error)
}

// SemesterHandler handles HTTP requests for Semester endpoints
type SemesterHandler struct {
	semesterService ISemesterApplicationService
}

// NewSemesterHandler creates a new Semester handler
func NewSemesterHandler(semesterService *application.SemesterApplicationService) *SemesterHandler {
	return &SemesterHandler{
		semesterService: semesterService,
	}
}

// CreateSemester creates a new semester
// POST /semesters
func (h *SemesterHandler) CreateSemester(c *gin.Context) {
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

	var req dto.CreateSemesterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INVALID_REQUEST",
				Message: err.Error(),
			},
		})
		return
	}

	cmd := &application.CreateSemesterCommand{
		AcademicYearID: req.AcademicYearID,
		Type:           domain.SemesterType(req.Type),
		Name:           req.Name,
		StartDate:      req.StartDate,
		EndDate:        req.EndDate,
		SequenceNumber: req.SequenceNumber,
		UserID:         authCtx.UserID,
	}

	resp, err := h.semesterService.CreateSemester(c.Request.Context(), cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusCreated, dto.SemesterResponse{
		ID:     resp.SemesterID,
		Status: string(resp.Status),
	})
}

// UpdateSemester updates a semester
// PUT /semesters/:id
func (h *SemesterHandler) UpdateSemester(c *gin.Context) {
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
				Message: "Semester ID is required",
			},
		})
		return
	}

	var req dto.UpdateSemesterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INVALID_REQUEST",
				Message: err.Error(),
			},
		})
		return
	}

	var status *domain.SemesterStatus
	if req.Status != nil {
		s := domain.SemesterStatus(*req.Status)
		status = &s
	}

	cmd := &application.UpdateSemesterCommand{
		SemesterID: id,
		Name:       req.Name,
		StartDate:  req.StartDate,
		EndDate:    req.EndDate,
		Status:     status,
		UserID:     authCtx.UserID,
	}

	resp, err := h.semesterService.UpdateSemester(c.Request.Context(), cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, dto.SemesterResponse{
		ID:     resp.SemesterID,
		Status: string(resp.Status),
	})
}

// DeleteSemester deletes a semester
// DELETE /semesters/:id
func (h *SemesterHandler) DeleteSemester(c *gin.Context) {
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
				Message: "Semester ID is required",
			},
		})
		return
	}

	cmd := &application.DeleteSemesterCommand{
		SemesterID: id,
		UserID:     authCtx.UserID,
	}

	resp, err := h.semesterService.DeleteSemester(c.Request.Context(), cmd)
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

// GetSemester retrieves a semester
// GET /semesters/:id
func (h *SemesterHandler) GetSemester(c *gin.Context) {
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
				Message: "Semester ID is required",
			},
		})
		return
	}

	cmd := &application.GetSemesterCommand{
		SemesterID: id,
		UserID:     authCtx.UserID,
	}

	resp, err := h.semesterService.GetSemester(c.Request.Context(), cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, dto.SemesterResponse{
		ID:             resp.ID,
		AcademicYearID: resp.AcademicYearID,
		Type:           string(resp.Type),
		Name:           resp.Name,
		StartDate:      resp.StartDate.Format(time.RFC3339),
		EndDate:        resp.EndDate.Format(time.RFC3339),
		Status:         string(resp.Status),
		SequenceNumber: resp.SequenceNumber,
		CreatedBy:      resp.CreatedBy,
		CreatedAt:      resp.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      resp.UpdatedAt.Format(time.RFC3339),
	})
}

// ListSemesters retrieves all semesters for an academic year
// GET /semesters?academic_year_id=:academic_year_id
func (h *SemesterHandler) ListSemesters(c *gin.Context) {
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

	academicYearID := c.Query("academic_year_id")
	if academicYearID == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INVALID_REQUEST",
				Message: "Academic Year ID is required",
			},
		})
		return
	}

	cmd := &application.ListSemestersCommand{
		AcademicYearID: academicYearID,
		UserID:         authCtx.UserID,
	}

	resp, err := h.semesterService.ListSemesters(c.Request.Context(), cmd)
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

	c.JSON(http.StatusOK, dto.ListSemestersResponse{
		Semesters: semesterResponses,
		Total:     len(resp.Semesters),
	})
}

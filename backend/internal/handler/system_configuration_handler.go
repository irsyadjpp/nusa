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

// ISystemConfigurationApplicationService defines the interface for System Configuration application service
type ISystemConfigurationApplicationService interface {
	CreateSystemConfiguration(ctx context.Context, cmd *application.CreateSystemConfigurationCommand) (*application.CreateSystemConfigurationResponse, error)
	UpdateSystemConfiguration(ctx context.Context, cmd *application.UpdateSystemConfigurationCommand) (*application.UpdateSystemConfigurationResponse, error)
	DeleteSystemConfiguration(ctx context.Context, cmd *application.DeleteSystemConfigurationCommand) (*application.DeleteSystemConfigurationResponse, error)
	GetSystemConfiguration(ctx context.Context, cmd *application.GetSystemConfigurationCommand) (*application.GetSystemConfigurationResponse, error)
	GetSystemConfigurationByKey(ctx context.Context, cmd *application.GetSystemConfigurationByKeyCommand) (*application.GetSystemConfigurationByKeyResponse, error)
	ListSystemConfigurations(ctx context.Context, cmd *application.ListSystemConfigurationsCommand) (*application.ListSystemConfigurationsResponse, error)
}

// SystemConfigurationHandler handles HTTP requests for System Configuration endpoints
type SystemConfigurationHandler struct {
	systemConfigurationService ISystemConfigurationApplicationService
}

// NewSystemConfigurationHandler creates a new System Configuration handler
func NewSystemConfigurationHandler(systemConfigurationService *application.SystemConfigurationApplicationService) *SystemConfigurationHandler {
	return &SystemConfigurationHandler{
		systemConfigurationService: systemConfigurationService,
	}
}

// CreateSystemConfiguration creates a new system configuration
// POST /system-configurations
func (h *SystemConfigurationHandler) CreateSystemConfiguration(c *gin.Context) {
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

	var req dto.CreateSystemConfigurationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INVALID_REQUEST",
				Message: err.Error(),
			},
		})
		return
	}

	cmd := &application.CreateSystemConfigurationCommand{
		Key:         req.Key,
		Value:       req.Value,
		ValueType:   req.ValueType,
		Description: req.Description,
		Category:    req.Category,
		IsSystem:    req.IsSystem,
		UserID:      authCtx.UserID,
	}

	resp, err := h.systemConfigurationService.CreateSystemConfiguration(c.Request.Context(), cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusCreated, dto.SystemConfigurationResponse{
		ID: resp.SystemConfigurationID,
	})
}

// UpdateSystemConfiguration updates a system configuration
// PUT /system-configurations/:id
func (h *SystemConfigurationHandler) UpdateSystemConfiguration(c *gin.Context) {
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
				Message: "System Configuration ID is required",
			},
		})
		return
	}

	var req dto.UpdateSystemConfigurationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INVALID_REQUEST",
				Message: err.Error(),
			},
		})
		return
	}

	cmd := &application.UpdateSystemConfigurationCommand{
		SystemConfigurationID: id,
		Value:                 req.Value,
		ValueType:             req.ValueType,
		Description:           req.Description,
		Category:              req.Category,
		IsActive:              req.IsActive,
		UserID:                authCtx.UserID,
	}

	resp, err := h.systemConfigurationService.UpdateSystemConfiguration(c.Request.Context(), cmd)
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

// DeleteSystemConfiguration deletes a system configuration
// DELETE /system-configurations/:id
func (h *SystemConfigurationHandler) DeleteSystemConfiguration(c *gin.Context) {
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
				Message: "System Configuration ID is required",
			},
		})
		return
	}

	cmd := &application.DeleteSystemConfigurationCommand{
		SystemConfigurationID: id,
		UserID:                authCtx.UserID,
	}

	resp, err := h.systemConfigurationService.DeleteSystemConfiguration(c.Request.Context(), cmd)
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

// GetSystemConfiguration retrieves a system configuration
// GET /system-configurations/:id
func (h *SystemConfigurationHandler) GetSystemConfiguration(c *gin.Context) {
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
				Message: "System Configuration ID is required",
			},
		})
		return
	}

	cmd := &application.GetSystemConfigurationCommand{
		SystemConfigurationID: id,
		UserID:                authCtx.UserID,
	}

	resp, err := h.systemConfigurationService.GetSystemConfiguration(c.Request.Context(), cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, dto.SystemConfigurationResponse{
		ID:          resp.ID,
		Key:         resp.Key,
		Value:       resp.Value,
		ValueType:   resp.ValueType,
		Description: resp.Description,
		Category:    resp.Category,
		IsSystem:    resp.IsSystem,
		IsActive:    resp.IsActive,
		CreatedBy:   resp.CreatedBy,
		CreatedAt:   resp.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   resp.UpdatedAt.Format(time.RFC3339),
	})
}

// GetSystemConfigurationByKey retrieves a system configuration by key
// GET /system-configurations/by-key/:key
func (h *SystemConfigurationHandler) GetSystemConfigurationByKey(c *gin.Context) {
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

	key := c.Param("key")
	if key == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INVALID_REQUEST",
				Message: "Configuration key is required",
			},
		})
		return
	}

	cmd := &application.GetSystemConfigurationByKeyCommand{
		Key:    key,
		UserID: authCtx.UserID,
	}

	resp, err := h.systemConfigurationService.GetSystemConfigurationByKey(c.Request.Context(), cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, dto.SystemConfigurationResponse{
		ID:          resp.ID,
		Key:         resp.Key,
		Value:       resp.Value,
		ValueType:   resp.ValueType,
		Description: resp.Description,
		Category:    resp.Category,
		IsSystem:    resp.IsSystem,
		IsActive:    resp.IsActive,
		CreatedBy:   resp.CreatedBy,
		CreatedAt:   resp.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   resp.UpdatedAt.Format(time.RFC3339),
	})
}

// ListSystemConfigurations retrieves system configurations
// GET /system-configurations?category=:category&active_only=:active_only
func (h *SystemConfigurationHandler) ListSystemConfigurations(c *gin.Context) {
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

	var category *string
	if cat := c.Query("category"); cat != "" {
		category = &cat
	}

	activeOnly := c.Query("active_only") == "true"

	cmd := &application.ListSystemConfigurationsCommand{
		Category:   category,
		ActiveOnly: activeOnly,
		UserID:     authCtx.UserID,
	}

	resp, err := h.systemConfigurationService.ListSystemConfigurations(c.Request.Context(), cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: dto.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	// Convert system configurations to DTO
	configResponses := make([]dto.SystemConfigurationResponse, len(resp.SystemConfigurations))
	for i, sc := range resp.SystemConfigurations {
		configResponses[i] = dto.SystemConfigurationResponse{
			ID:          sc.ID,
			Key:         sc.Key,
			Value:       sc.Value,
			ValueType:   sc.ValueType,
			Description: sc.Description,
			Category:    sc.Category,
			IsSystem:    sc.IsSystem,
			IsActive:    sc.IsActive,
			CreatedBy:   sc.CreatedBy,
			CreatedAt:   sc.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   sc.UpdatedAt.Format(time.RFC3339),
		}
	}

	c.JSON(http.StatusOK, dto.ListSystemConfigurationsResponse{
		SystemConfigurations: configResponses,
		Total:                 len(resp.SystemConfigurations),
	})
}
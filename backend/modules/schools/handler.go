package schools

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
	schoolService *service.SchoolService
}

func NewHandler(schoolService *service.SchoolService) *Handler {
	return &Handler{
		schoolService: schoolService,
	}
}

func (h *Handler) CreateSchool(c *gin.Context) {
	var req domain.CreateSchoolRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, map[string]string{
			"name": "School name is required",
			"code": "School code is required",
		})
		return
	}

	ctx := context.Background()
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		response.Error(c, 401, "Authentication required")
		return
	}

	// Only SYSTEM_ADMIN can create schools
	if !middleware.HasPermission(authCtx.Role, "school", domain.ActionCreate) {
		response.Error(c, 403, "Insufficient permissions to create school")
		return
	}

	school, err := h.schoolService.CreateSchool(ctx, &req, authCtx.UserID)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	response.Success(c, school)
}

func (h *Handler) GetSchools(c *gin.Context) {
	ctx := context.Background()
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		response.Error(c, 401, "Authentication required")
		return
	}

	// Check permission
	if !middleware.HasPermission(authCtx.Role, "school", domain.ActionRead) {
		response.Error(c, 403, "Insufficient permissions to view schools")
		return
	}

	// Parse query parameters
	var isActive *bool
	if active := c.Query("is_active"); active != "" {
		b, _ := strconv.ParseBool(active)
		isActive = &b
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	schools, total, err := h.schoolService.ListSchools(ctx, isActive, page, pageSize)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, gin.H{
		"schools":   schools,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}

func (h *Handler) GetSchool(c *gin.Context) {
	ctx := context.Background()
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		response.Error(c, 401, "Authentication required")
		return
	}

	schoolID := c.Param("id")

	// Check permission
	if !middleware.HasPermission(authCtx.Role, "school", domain.ActionRead) {
		// Non-admins can only view their own school
		if authCtx.SchoolID == nil || *authCtx.SchoolID != schoolID {
			response.Error(c, 403, "Insufficient permissions")
			return
		}
	}

	school, err := h.schoolService.GetSchool(ctx, schoolID)
	if err != nil {
		response.Error(c, 404, "School not found")
		return
	}

	response.Success(c, school)
}

func (h *Handler) UpdateSchool(c *gin.Context) {
	ctx := context.Background()
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		response.Error(c, 401, "Authentication required")
		return
	}

	schoolID := c.Param("id")

	// Check permission
	if !middleware.HasPermission(authCtx.Role, "school", domain.ActionUpdate) {
		// Non-admins cannot update schools
		response.Error(c, 403, "Insufficient permissions to update school")
		return
	}

	var req domain.UpdateSchoolRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, map[string]string{
			"name": "Name is required if provided",
		})
		return
	}

	school, err := h.schoolService.UpdateSchool(ctx, schoolID, &req, authCtx.UserID)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	response.Success(c, school)
}

func (h *Handler) UpdateSchoolStatus(c *gin.Context) {
	ctx := context.Background()
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		response.Error(c, 401, "Authentication required")
		return
	}

	schoolID := c.Param("id")

	// Check permission
	if !middleware.HasPermission(authCtx.Role, "school", domain.ActionUpdate) {
		response.Error(c, 403, "Insufficient permissions to update school status")
		return
	}

	var req domain.UpdateSchoolStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "Invalid status")
		return
	}

	err := h.schoolService.UpdateSchoolStatus(ctx, schoolID, req.Status)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "School status updated successfully"})
}

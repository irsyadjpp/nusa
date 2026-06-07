package roles

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
	roleService *service.RoleService
}

func NewHandler(roleService *service.RoleService) *Handler {
	return &Handler{
		roleService: roleService,
	}
}

func (h *Handler) CreateRole(c *gin.Context) {
	var req domain.CreateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, map[string]string{
			"name": "Role name is required (min 2 characters)",
		})
		return
	}

	ctx := context.Background()
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		response.Error(c, 401, "Authentication required")
		return
	}

	// Only SYSTEM_ADMIN can create roles
	if authCtx.Role != domain.RoleSystemAdmin {
		response.Error(c, 403, "Insufficient permissions to create roles")
		return
	}

	role, err := h.roleService.CreateRole(ctx, &req)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	response.Created(c, role)
}

func (h *Handler) GetRoles(c *gin.Context) {
	ctx := context.Background()
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		response.Error(c, 401, "Authentication required")
		return
	}

	// Only SYSTEM_ADMIN can view all roles
	if authCtx.Role != domain.RoleSystemAdmin {
		response.Error(c, 403, "Insufficient permissions to view roles")
		return
	}

	var isActive *bool
	if active := c.Query("is_active"); active != "" {
		b, _ := strconv.ParseBool(active)
		isActive = &b
	}

	roles, err := h.roleService.ListRoles(ctx, isActive)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, gin.H{
		"roles": roles,
	})
}

func (h *Handler) GetRole(c *gin.Context) {
	ctx := context.Background()
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		response.Error(c, 401, "Authentication required")
		return
	}

	// Only SYSTEM_ADMIN can view roles
	if authCtx.Role != domain.RoleSystemAdmin {
		response.Error(c, 403, "Insufficient permissions to view roles")
		return
	}

	roleID := c.Param("id")

	role, err := h.roleService.GetRole(ctx, roleID)
	if err != nil {
		response.Error(c, 404, "Role not found")
		return
	}

	response.Success(c, role)
}

func (h *Handler) UpdateRole(c *gin.Context) {
	ctx := context.Background()
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		response.Error(c, 401, "Authentication required")
		return
	}

	// Only SYSTEM_ADMIN can update roles
	if authCtx.Role != domain.RoleSystemAdmin {
		response.Error(c, 403, "Insufficient permissions to update roles")
		return
	}

	roleID := c.Param("id")

	var req domain.UpdateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, map[string]string{
			"name": "Name is required if provided",
		})
		return
	}

	role, err := h.roleService.UpdateRole(ctx, roleID, &req)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	response.Success(c, role)
}

func (h *Handler) DeleteRole(c *gin.Context) {
	ctx := context.Background()
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		response.Error(c, 401, "Authentication required")
		return
	}

	// Only SYSTEM_ADMIN can delete roles
	if authCtx.Role != domain.RoleSystemAdmin {
		response.Error(c, 403, "Insufficient permissions to delete roles")
		return
	}

	roleID := c.Param("id")

	err := h.roleService.DeleteRole(ctx, roleID)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "Role deleted successfully"})
}

func (h *Handler) AddPermission(c *gin.Context) {
	ctx := context.Background()
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		response.Error(c, 401, "Authentication required")
		return
	}

	// Only SYSTEM_ADMIN can manage permissions
	if authCtx.Role != domain.RoleSystemAdmin {
		response.Error(c, 403, "Insufficient permissions to manage permissions")
		return
	}

	roleID := c.Param("id")

	var req domain.CreatePermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, map[string]string{
			"resource": "Resource is required",
			"action":   "Action is required",
		})
		return
	}

	err := h.roleService.AddPermission(ctx, roleID, req.Resource, req.Action)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "Permission added successfully"})
}

func (h *Handler) GetPermissions(c *gin.Context) {
	ctx := context.Background()
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		response.Error(c, 401, "Authentication required")
		return
	}

	// Only SYSTEM_ADMIN can view permissions
	if authCtx.Role != domain.RoleSystemAdmin {
		response.Error(c, 403, "Insufficient permissions to view permissions")
		return
	}

	roleID := c.Param("id")

	permissions, err := h.roleService.GetPermissions(ctx, roleID)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, gin.H{
		"permissions": permissions,
	})
}

func (h *Handler) RemovePermission(c *gin.Context) {
	ctx := context.Background()
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		response.Error(c, 401, "Authentication required")
		return
	}

	// Only SYSTEM_ADMIN can manage permissions
	if authCtx.Role != domain.RoleSystemAdmin {
		response.Error(c, 403, "Insufficient permissions to manage permissions")
		return
	}

	roleID := c.Param("id")
	resource := c.Query("resource")
	action := c.Query("action")

	if resource == "" || action == "" {
		response.Error(c, 400, "Resource and action are required")
		return
	}

	err := h.roleService.RemovePermission(ctx, roleID, resource, action)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "Permission removed successfully"})
}

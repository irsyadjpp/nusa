package auth

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/nusa/backend/internal/middleware"
	"github.com/nusa/backend/pkg/response"
)

func (h *Handler) Me(c *gin.Context) {
	ctx := context.Background()
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		response.Error(c, 401, "Authentication required")
		return
	}

	// Get user
	user, err := h.userService.GetUser(ctx, authCtx.UserID)
	if err != nil {
		response.Error(c, 404, "User not found")
		return
	}

	// Get role name
	role, err := h.roleRepo.GetByID(ctx, user.RoleID)
	if err != nil {
		response.Error(c, 500, "Failed to get user role")
		return
	}

	// Get school name if applicable
	var schoolName *string
	if user.SchoolID != nil {
		school, err := h.schoolRepo.GetByID(ctx, *user.SchoolID)
		if err == nil {
			schoolName = &school.Name
		}
	}

	// Get permissions
	permissions, err := h.roleRepo.GetPermissions(ctx, user.RoleID)
	if err != nil {
		response.Error(c, 500, "Failed to get user permissions")
		return
	}

	permissionStrings := make([]string, len(permissions))
	for i, perm := range permissions {
		permissionStrings[i] = perm.Resource + ":" + perm.Action
	}

	response.Success(c, gin.H{
		"user":        user.ToUserResponse(role.Name, stringOrEmpty(schoolName)),
		"role_name":   role.Name,
		"permissions": permissionStrings,
	})
}

package users

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/middleware"
	"github.com/nusa/backend/internal/repository"
	"github.com/nusa/backend/internal/service"
	"github.com/nusa/backend/pkg/response"
)

type Handler struct {
	userService *service.UserService
	roleRepo    *repository.RoleRepository
	schoolRepo  *repository.SchoolRepository
}

func NewHandler(
	userService *service.UserService,
	roleRepo *repository.RoleRepository,
	schoolRepo *repository.SchoolRepository,
) *Handler {
	return &Handler{
		userService: userService,
		roleRepo:    roleRepo,
		schoolRepo:  schoolRepo,
	}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var req domain.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, map[string]string{
			"email":    "Valid email is required",
			"password": "Password must be at least 8 characters",
			"name":     "Name is required",
			"role_id":  "Role ID is required",
		})
		return
	}

	ctx := context.Background()
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		response.Error(c, 401, "Authentication required")
		return
	}

	// Check permission
	if !middleware.HasPermission(authCtx.Role, "user", domain.ActionCreate) {
		response.Error(c, 403, "Insufficient permissions to create user")
		return
	}

	// If creating a school user, ensure the creator belongs to the same school
	if req.SchoolID != nil && authCtx.Role != domain.RoleSystemAdmin {
		if authCtx.SchoolID == nil || *authCtx.SchoolID != *req.SchoolID {
			response.Error(c, 403, "Cannot create user for different school")
			return
		}
	}

	user, err := h.userService.Register(ctx, &req, authCtx.UserID)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	response.Success(c, user)
}

func (h *Handler) GetUsers(c *gin.Context) {
	ctx := context.Background()
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		response.Error(c, 401, "Authentication required")
		return
	}

	// Check permission
	if !middleware.HasPermission(authCtx.Role, "user", domain.ActionRead) {
		response.Error(c, 403, "Insufficient permissions to view users")
		return
	}

	// Parse query parameters
	var schoolID *string
	if authCtx.Role == domain.RoleSchoolAdmin || authCtx.Role == domain.RoleTeacher {
		schoolID = authCtx.SchoolID
	} else {
		if s := c.Query("school_id"); s != "" {
			schoolID = &s
		}
	}

	var roleID *string
	if r := c.Query("role_id"); r != "" {
		roleID = &r
	}

	var isActive *bool
	if active := c.Query("is_active"); active != "" {
		b, _ := strconv.ParseBool(active)
		isActive = &b
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	users, total, err := h.userService.ListUsers(ctx, schoolID, roleID, isActive, page, pageSize)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, gin.H{
		"users": users,
		"total": total,
		"page":  page,
		"page_size": pageSize,
	})
}

func (h *Handler) GetUser(c *gin.Context) {
	ctx := context.Background()
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		response.Error(c, 401, "Authentication required")
		return
	}

	userID := c.Param("id")

	// Check permission
	if !middleware.HasPermission(authCtx.Role, "user", domain.ActionRead) {
		// Users can only view their own profile
		if authCtx.UserID != userID {
			response.Error(c, 403, "Insufficient permissions")
			return
		}
	}

	user, err := h.userService.GetUser(ctx, userID)
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

	response.Success(c, user.ToUserResponse(role.Name, stringOrEmpty(schoolName)))
}

func (h *Handler) UpdateUser(c *gin.Context) {
	ctx := context.Background()
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		response.Error(c, 401, "Authentication required")
		return
	}

	userID := c.Param("id")

	// Check permission
	if !middleware.HasPermission(authCtx.Role, "user", domain.ActionUpdate) {
		// Users can only update their own profile (excluding role and school)
		if authCtx.UserID != userID {
			response.Error(c, 403, "Insufficient permissions")
			return
		}
	}

	var req domain.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, map[string]string{
			"name": "Name is required if provided",
		})
		return
	}

	// Non-admins cannot change role or school
	if authCtx.Role != domain.RoleSystemAdmin && (req.RoleID != nil || req.SchoolID != nil) {
		response.Error(c, 403, "Cannot change role or school")
		return
	}

	user, err := h.userService.UpdateUser(ctx, userID, &req, authCtx.UserID)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	response.Success(c, user)
}

func (h *Handler) UpdateUserStatus(c *gin.Context) {
	ctx := context.Background()
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		response.Error(c, 401, "Authentication required")
		return
	}

	userID := c.Param("id")

	// Check permission
	if !middleware.HasPermission(authCtx.Role, "user", domain.ActionUpdate) {
		response.Error(c, 403, "Insufficient permissions to update user status")
		return
	}

	var req domain.UpdateUserStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "Invalid status")
		return
	}

	err := h.userService.UpdateUserStatus(ctx, userID, req.Status)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "User status updated successfully"})
}

func stringOrEmpty(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

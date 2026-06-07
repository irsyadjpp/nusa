package auth

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/repository"
	"github.com/nusa/backend/internal/service"
	jwtService "github.com/nusa/backend/pkg/jwt"
	"github.com/nusa/backend/pkg/response"
)

type Handler struct {
	userService      *service.UserService
	refreshTokenRepo *repository.RefreshTokenRepository
	jwtService       *jwtService.Service
	roleRepo         *repository.RoleRepository
	schoolRepo       *repository.SchoolRepository
}

func NewHandler(
	userService *service.UserService,
	refreshTokenRepo *repository.RefreshTokenRepository,
	jwtService *jwtService.Service,
	roleRepo *repository.RoleRepository,
	schoolRepo *repository.SchoolRepository,
) *Handler {
	return &Handler{
		userService:      userService,
		refreshTokenRepo: refreshTokenRepo,
		jwtService:       jwtService,
		roleRepo:         roleRepo,
		schoolRepo:       schoolRepo,
	}
}

func (h *Handler) Login(c *gin.Context) {
	var req domain.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, map[string]string{
			"email":    "Valid email is required",
			"password": "Password is required",
		})
		return
	}

	ctx := context.Background()

	// Validate credentials
	user, err := h.userService.ValidateCredentials(ctx, req.Email, req.Password)
	if err != nil {
		response.Error(c, 401, "Invalid credentials")
		return
	}

	// Check if user is active
	if !user.IsActive {
		response.Error(c, 403, "User account is not active")
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

	// Generate tokens
	accessToken, err := h.jwtService.GenerateAccessToken(user.ID, role.Name, user.SchoolID, permissionStrings)
	if err != nil {
		response.Error(c, 500, "Failed to generate access token")
		return
	}

	refreshToken, err := h.jwtService.GenerateRefreshToken(user.ID)
	if err != nil {
		response.Error(c, 500, "Failed to generate refresh token")
		return
	}

	// Store refresh token
	clientIP := c.ClientIP()
	err = h.refreshTokenRepo.Create(ctx, user.ID, refreshToken, time.Now().Add(7*24*time.Hour), &clientIP, nil)
	if err != nil {
		response.Error(c, 500, "Failed to store refresh token")
		return
	}

	// Update last login time
	user.UpdatedAt = time.Now()
	_, _ = h.userService.UpdateUser(ctx, user.ID, &domain.UpdateUserRequest{}, user.ID)

	response.Success(c, domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    3600,
		User:         user.ToUserResponse(role.Name, stringOrEmpty(schoolName)),
	})
}

func (h *Handler) Refresh(c *gin.Context) {
	var req domain.RefreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "Refresh token is required")
		return
	}

	ctx := context.Background()

	// Validate refresh token
	userID, err := h.refreshTokenRepo.GetByToken(ctx, req.RefreshToken)
	if err != nil {
		response.Error(c, 401, "Invalid or expired refresh token")
		return
	}

	// Get user
	user, err := h.userService.GetUser(ctx, *userID)
	if err != nil {
		response.Error(c, 401, "User not found")
		return
	}

	// Check if user is still active
	if !user.IsActive {
		response.Error(c, 401, "User account is not active")
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

	// Generate new access token
	accessToken, err := h.jwtService.GenerateAccessToken(user.ID, role.Name, user.SchoolID, permissionStrings)
	if err != nil {
		response.Error(c, 500, "Failed to generate access token")
		return
	}

	// Generate new refresh token (rotation)
	newRefreshToken, err := h.jwtService.GenerateRefreshToken(user.ID)
	if err != nil {
		response.Error(c, 500, "Failed to generate refresh token")
		return
	}

	// Revoke old refresh token
	_ = h.refreshTokenRepo.Revoke(ctx, req.RefreshToken)

	// Store new refresh token
	clientIP := c.ClientIP()
	err = h.refreshTokenRepo.Create(ctx, user.ID, newRefreshToken, time.Now().Add(7*24*time.Hour), &clientIP, nil)
	if err != nil {
		response.Error(c, 500, "Failed to store refresh token")
		return
	}

	response.Success(c, domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    3600,
		User:         user.ToUserResponse(role.Name, stringOrEmpty(schoolName)),
	})
}

func (h *Handler) Logout(c *gin.Context) {
	var req domain.RefreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "Refresh token is required")
		return
	}

	ctx := context.Background()

	// Revoke refresh token
	err := h.refreshTokenRepo.Revoke(ctx, req.RefreshToken)
	if err != nil {
		// Don't fail if token doesn't exist, just return success
		response.Success(c, gin.H{
			"message": "Logged out successfully",
		})
		return
	}

	response.Success(c, gin.H{
		"message": "Logged out successfully",
	})
}

func stringOrEmpty(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

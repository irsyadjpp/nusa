package domain

import (
	"time"

	"github.com/google/uuid"
)

// UserStatus represents the status of a user
type UserStatus string

const (
	UserStatusActive    UserStatus = "ACTIVE"
	UserStatusInactive  UserStatus = "INACTIVE"
	UserStatusSuspended UserStatus = "SUSPENDED"
)

// User represents a user in the system
type User struct {
	ID                  string     `json:"id" db:"id"`
	Email               string     `json:"email" db:"email"`
	PasswordHash        string     `json:"-" db:"password_hash"` // Never exposed in JSON
	Name                string     `json:"name" db:"name"`
	RoleID              string     `json:"role_id" db:"role_id"`
	SchoolID            *string    `json:"school_id,omitempty" db:"school_id"`
	IsActive            bool       `json:"is_active" db:"is_active"`
	Status              UserStatus `json:"status" db:"-"`
	FailedLoginAttempts int        `json:"failed_login_attempts" db:"failed_login_attempts"`
	LockedUntil         *time.Time `json:"locked_until,omitempty" db:"locked_until"`
	CreatedAt           time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at" db:"updated_at"`
	CreatedBy           *string    `json:"created_by,omitempty" db:"created_by"`
	UpdatedBy           *string    `json:"updated_by,omitempty" db:"updated_by"`
}

// CreateUserRequest represents the request to create a new user
type CreateUserRequest struct {
	Email    string  `json:"email" binding:"required,email"`
	Password string  `json:"password" binding:"required,min=8"`
	Name     string  `json:"name" binding:"required,min=2"`
	RoleID   string  `json:"role_id" binding:"required"`
	SchoolID *string `json:"school_id,omitempty"`
}

// UpdateUserRequest represents the request to update a user
type UpdateUserRequest struct {
	Name     *string `json:"name,omitempty"`
	RoleID   *string `json:"role_id,omitempty"`
	SchoolID *string `json:"school_id,omitempty"`
}

// UpdateUserStatusRequest represents the request to update user status
type UpdateUserStatusRequest struct {
	Status UserStatus `json:"status" binding:"required,oneof=ACTIVE INACTIVE SUSPENDED"`
}

// LoginRequest represents the login request
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse represents the login response
type LoginResponse struct {
	AccessToken  string        `json:"access_token"`
	RefreshToken string        `json:"refresh_token"`
	TokenType    string        `json:"token_type"`
	ExpiresIn    int64         `json:"expires_in"`
	User         *UserResponse `json:"user"`
}

// RefreshRequest represents the refresh token request
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// UserResponse represents the user data returned to clients
type UserResponse struct {
	ID         string     `json:"id"`
	Email      string     `json:"email"`
	Name       string     `json:"name"`
	RoleID     string     `json:"role_id"`
	RoleName   string     `json:"role_name"`
	SchoolID   *string    `json:"school_id,omitempty"`
	SchoolName *string    `json:"school_name,omitempty"`
	IsActive   bool       `json:"is_active"`
	Status     UserStatus `json:"status"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

// ToUserResponse converts User to UserResponse
func (u *User) ToUserResponse(roleName, schoolName string) *UserResponse {
	status := UserStatusActive
	if !u.IsActive {
		status = UserStatusInactive
	}
	if u.LockedUntil != nil && u.LockedUntil.After(time.Now()) {
		status = UserStatusSuspended
	}

	var schoolNamePtr *string
	if schoolName != "" {
		schoolNamePtr = &schoolName
	}

	return &UserResponse{
		ID:         u.ID,
		Email:      u.Email,
		Name:       u.Name,
		RoleID:     u.RoleID,
		RoleName:   roleName,
		SchoolID:   u.SchoolID,
		SchoolName: schoolNamePtr,
		IsActive:   u.IsActive,
		Status:     status,
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
	}
}

// NewID generates a new UUID v7 ID
// Note: Using UUID v4 as fallback since UUID v7 support is limited in Go libraries
// The database uses gen_uuid_v7() function for proper UUID v7 generation
func NewID() string {
	return uuid.New().String()
}

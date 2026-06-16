package service

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/repository"
)

// UserService handles business logic for user operations
type UserService struct {
	userRepo repository.UserRepositoryInterface
	roleRepo repository.RoleRepositoryInterface
}

// NewUserService creates a new user service
func NewUserService(userRepo repository.UserRepositoryInterface, roleRepo repository.RoleRepositoryInterface) *UserService {
	return &UserService{
		userRepo: userRepo,
		roleRepo: roleRepo,
	}
}

// Register creates a new user with hashed password
func (s *UserService) Register(ctx context.Context, req *domain.CreateUserRequest, creatorID string) (*domain.User, error) {
	// Hash the password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Verify role exists
	role, err := s.roleRepo.GetByID(ctx, req.RoleID)
	if err != nil {
		return nil, fmt.Errorf("invalid role: %w", err)
	}
	if !role.IsActive {
		return nil, fmt.Errorf("role is not active")
	}

	user := &domain.User{
		ID:                  domain.NewID(),
		Email:               req.Email,
		PasswordHash:        string(passwordHash),
		Name:                req.Name,
		RoleID:              req.RoleID,
		SchoolID:            req.SchoolID,
		IsActive:            true,
		FailedLoginAttempts: 0,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
		CreatedBy:           &creatorID,
		UpdatedBy:           &creatorID,
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}

// ValidateCredentials validates user credentials
func (s *UserService) ValidateCredentials(ctx context.Context, email, password string) (*domain.User, error) {
	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	// Check if user is locked
	if user.LockedUntil != nil && user.LockedUntil.After(time.Now()) {
		return nil, fmt.Errorf("account is locked until %s", user.LockedUntil.Format(time.RFC3339))
	}

	// Check if user is active
	if !user.IsActive {
		return nil, fmt.Errorf("account is not active")
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		// Increment failed login attempts
		s.incrementFailedAttempts(ctx, user.ID)
		return nil, fmt.Errorf("invalid credentials")
	}

	// Reset failed login attempts on successful login
	if user.FailedLoginAttempts > 0 {
		_ = s.userRepo.UpdateStatus(ctx, user.ID, true, nil, 0)
	}

	return user, nil
}

// incrementFailedAttempts increments failed login attempts and locks account after threshold
func (s *UserService) incrementFailedAttempts(ctx context.Context, userID string) {
	const maxAttempts = 5
	const lockDuration = 30 * time.Minute

	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return
	}

	newAttempts := user.FailedLoginAttempts + 1
	var lockUntil *string

	if newAttempts >= maxAttempts {
		lockTime := time.Now().Add(lockDuration)
		lockUntilStr := lockTime.Format(time.RFC3339)
		lockUntil = &lockUntilStr
	}

	_ = s.userRepo.UpdateStatus(ctx, userID, true, lockUntil, newAttempts)
}

// GetUser retrieves a user by ID
func (s *UserService) GetUser(ctx context.Context, id string) (*domain.User, error) {
	return s.userRepo.GetByID(ctx, id)
}

// ListUsers retrieves users with pagination and filters
func (s *UserService) ListUsers(ctx context.Context, schoolID, roleID *string, isActive *bool, page, pageSize int) ([]*domain.User, int, error) {
	limit := pageSize
	offset := (page - 1) * pageSize

	users, err := s.userRepo.List(ctx, schoolID, roleID, isActive, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list users: %w", err)
	}

	total, err := s.userRepo.Count(ctx, schoolID, roleID, isActive)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list users by school: %w", err)
	}

	return users, total, nil
}

// UpdateUser updates user information
func (s *UserService) UpdateUser(ctx context.Context, id string, req *domain.UpdateUserRequest, updaterID string) (*domain.User, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.RoleID != nil {
		// Verify role exists
		role, err := s.roleRepo.GetByID(ctx, *req.RoleID)
		if err != nil {
			return nil, fmt.Errorf("invalid role: %w", err)
		}
		if !role.IsActive {
			return nil, fmt.Errorf("role is not active")
		}
		user.RoleID = *req.RoleID
	}
	if req.SchoolID != nil {
		user.SchoolID = req.SchoolID
	}

	user.UpdatedBy = &updaterID
	user.UpdatedAt = time.Now()

	if err := s.userRepo.Update(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return user, nil
}

// UpdateUserStatus updates user status
func (s *UserService) UpdateUserStatus(ctx context.Context, id string, status domain.UserStatus) error {
	_, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	isActive := status == domain.UserStatusActive
	var lockUntil *string

	if status == domain.UserStatusSuspended {
		lockDuration := 24 * time.Hour // Default lock duration for suspension
		lockTime := time.Now().Add(lockDuration)
		lockUntilStr := lockTime.Format(time.RFC3339)
		lockUntil = &lockUntilStr
	}

	if status == domain.UserStatusActive {
		// Reset failed login attempts and lock when reactivating
		lockUntil = nil
	}

	return s.userRepo.UpdateStatus(ctx, id, isActive, lockUntil, 0)
}

// DeleteUser soft deletes a user
func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	return s.userRepo.Delete(ctx, id)
}

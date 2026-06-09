package application

import (
	"context"
	"fmt"
	"time"

	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/repository"
)

// TPSetApplicationService orchestrates TPSet use cases
// It coordinates repositories, authorization, validation, and transactions
// Business rules are enforced through the domain aggregate
type TPSetApplicationService struct {
	tpRepo     *repository.TPRepository
	userRepo   *repository.UserRepository
	schoolRepo *repository.SchoolRepository
}

// NewTPSetApplicationService creates a new TPSet application service
func NewTPSetApplicationService(
	tpRepo *repository.TPRepository,
	userRepo *repository.UserRepository,
	schoolRepo *repository.SchoolRepository,
) *TPSetApplicationService {
	return &TPSetApplicationService{
		tpRepo:     tpRepo,
		userRepo:   userRepo,
		schoolRepo: schoolRepo,
	}
}

// CreateTPSetCommand represents the command to create a TP Set
type CreateTPSetCommand struct {
	CPID             string
	VersionNo        int
	GenerationSource domain.GenerationSource
	GenerationReason *string
	UserID           string // Authenticated user ID
}

// CreateTPSetResponse represents the response for creating a TP Set
type CreateTPSetResponse struct {
	TPSetID string
	Status  domain.WorkflowStatus
}

// CreateTPSet creates a new TP Set
// Orchestrates: authorization, school scope validation, domain invariant enforcement, transaction
func (s *TPSetApplicationService) CreateTPSet(ctx context.Context, cmd *CreateTPSetCommand) (*CreateTPSetResponse, error) {
	// 1. Authorization: Get user and validate school scope
	user, err := s.userRepo.GetByID(ctx, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// 2. School scope: Validate user belongs to a school
	if user.SchoolID == nil || *user.SchoolID == "" {
		return nil, fmt.Errorf("user must belong to a school")
	}

	// 3. Create domain aggregate (enforces invariants)
	tpSet := &domain.TPSet{
		ID:               generateID(),
		CPID:             cmd.CPID,
		VersionNo:        cmd.VersionNo,
		Status:           domain.WorkflowStatusDraft,
		GenerationSource: cmd.GenerationSource,
		GenerationReason: cmd.GenerationReason,
		GeneratedBy:      cmd.UserID,
	}

	aggregate, err := domain.NewTPSetAggregate(tpSet, *user.SchoolID)
	if err != nil {
		return nil, fmt.Errorf("failed to create aggregate: %w", err)
	}

	// 4. Add initial version (enforces TP-INV-003, TP-INV-004, TP-INV-005)
	version := &domain.TPVersion{
		ID:               generateID(),
		TPSetID:          tpSet.ID,
		VersionNo:        1,
		IsCurrentVersion: true,
		Status:           domain.WorkflowStatusDraft,
		CreatedAt:        now(),
		CreatedBy:        cmd.UserID,
	}

	if err := aggregate.AddVersion(version); err != nil {
		return nil, fmt.Errorf("failed to add version: %w", err)
	}

	// 5. Transaction boundary: Persist aggregate
	if err := s.tpRepo.CreateTPSet(ctx, aggregate.GetTPSet()); err != nil {
		return nil, fmt.Errorf("failed to create TP set: %w", err)
	}

	// 6. Note: Version persistence would require additional repository method
	// For now, version is tracked in aggregate but not persisted separately
	// This is a simplification for the application layer

	return &CreateTPSetResponse{
		TPSetID: tpSet.ID,
		Status:  tpSet.Status,
	}, nil
}

// UpdateTPSetCommand represents the command to update a TP Set
type UpdateTPSetCommand struct {
	TPSetID          string
	Status           *domain.WorkflowStatus
	GenerationReason *string
	UserID           string // Authenticated user ID
}

// UpdateTPSetResponse represents the response for updating a TP Set
type UpdateTPSetResponse struct {
	TPSetID string
	Status  domain.WorkflowStatus
}

// UpdateTPSet updates a TP Set
// Orchestrates: authorization, ownership validation, domain invariant enforcement, transaction
func (s *TPSetApplicationService) UpdateTPSet(ctx context.Context, cmd *UpdateTPSetCommand) (*UpdateTPSetResponse, error) {
	// 1. Authorization: Get user
	user, err := s.userRepo.GetByID(ctx, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// 2. Load existing TP Set
	tpSet, err := s.tpRepo.GetTPSetByID(ctx, cmd.TPSetID)
	if err != nil {
		return nil, fmt.Errorf("TP set not found: %w", err)
	}

	// 3. Ownership validation: User must be owner or School Admin
	if tpSet.GeneratedBy != cmd.UserID && user.RoleID != "SCHOOL_ADMIN" && user.RoleID != "SYSTEM_ADMIN" {
		return nil, fmt.Errorf("user does not have permission to update this TP set")
	}

	// 4. School scope: Validate user belongs to same school as TP Set
	if user.SchoolID != nil && *user.SchoolID != "" {
		// Get TP Set owner's school
		owner, err := s.userRepo.GetByID(ctx, tpSet.GeneratedBy)
		if err != nil {
			return nil, fmt.Errorf("failed to get TP set owner: %w", err)
		}
		if owner.SchoolID != nil && user.SchoolID != nil && *user.SchoolID != *owner.SchoolID && user.RoleID != "SYSTEM_ADMIN" {
			return nil, fmt.Errorf("cross-school access not allowed")
		}
	}

	// 5. Create domain aggregate (enforces invariants)
	schoolID := ""
	if user.SchoolID != nil {
		schoolID = *user.SchoolID
	}
	aggregate, err := domain.NewTPSetAggregate(tpSet, schoolID)
	if err != nil {
		return nil, fmt.Errorf("failed to create aggregate: %w", err)
	}

	// 6. Apply updates through aggregate (enforces TP-INV-006, TP-INV-007)
	if cmd.Status != nil {
		if err := aggregate.TransitionStatus(*cmd.Status); err != nil {
			return nil, fmt.Errorf("failed to transition status: %w", err)
		}
	}

	if cmd.GenerationReason != nil {
		tpSet.GenerationReason = cmd.GenerationReason
	}

	// 7. Transaction boundary: Persist aggregate
	if err := s.tpRepo.UpdateTPSet(ctx, aggregate.GetTPSet()); err != nil {
		return nil, fmt.Errorf("failed to update TP set: %w", err)
	}

	return &UpdateTPSetResponse{
		TPSetID: tpSet.ID,
		Status:  tpSet.Status,
	}, nil
}

// ApproveTPSetCommand represents the command to approve a TP Set
type ApproveTPSetCommand struct {
	TPSetID    string
	ApproverID string
	Reason     string
}

// ApproveTPSetResponse represents the response for approving a TP Set
type ApproveTPSetResponse struct {
	TPSetID    string
	Status     domain.WorkflowStatus
	ApprovedBy string
	ApprovedAt string
}

// ApproveTPSet approves a TP Set
// Orchestrates: authorization, ownership validation, domain invariant enforcement, transaction
func (s *TPSetApplicationService) ApproveTPSet(ctx context.Context, cmd *ApproveTPSetCommand) (*ApproveTPSetResponse, error) {
	// 1. Authorization: Get approver
	approver, err := s.userRepo.GetByID(ctx, cmd.ApproverID)
	if err != nil {
		return nil, fmt.Errorf("approver not found: %w", err)
	}

	// 2. Authorization: Only School Admin or System Admin can approve
	if approver.RoleID != "SCHOOL_ADMIN" && approver.RoleID != "SYSTEM_ADMIN" {
		return nil, fmt.Errorf("only School Admin or System Admin can approve TP sets")
	}

	// 3. Load existing TP Set
	tpSet, err := s.tpRepo.GetTPSetByID(ctx, cmd.TPSetID)
	if err != nil {
		return nil, fmt.Errorf("TP set not found: %w", err)
	}

	// 4. School scope: Validate approver belongs to same school as TP Set
	if approver.RoleID == "SCHOOL_ADMIN" {
		owner, err := s.userRepo.GetByID(ctx, tpSet.GeneratedBy)
		if err != nil {
			return nil, fmt.Errorf("failed to get TP set owner: %w", err)
		}
		if owner.SchoolID != nil && approver.SchoolID != nil && *approver.SchoolID != *owner.SchoolID {
			return nil, fmt.Errorf("cross-school access not allowed")
		}
	}

	// 5. Create domain aggregate (enforces invariants)
	schoolID := ""
	if approver.SchoolID != nil {
		schoolID = *approver.SchoolID
	}
	aggregate, err := domain.NewTPSetAggregate(tpSet, schoolID)
	if err != nil {
		return nil, fmt.Errorf("failed to create aggregate: %w", err)
	}

	// 6. Transition status through aggregate (enforces TP-INV-006, TP-INV-007)
	if err := aggregate.TransitionStatus(domain.WorkflowStatusApproved); err != nil {
		return nil, fmt.Errorf("failed to approve: %w", err)
	}

	// 7. Set approval metadata
	now := now()
	tpSet.ApprovedBy = &cmd.ApproverID
	tpSet.ApprovedAt = &now

	// 8. Transaction boundary: Persist aggregate
	if err := s.tpRepo.UpdateTPSet(ctx, aggregate.GetTPSet()); err != nil {
		return nil, fmt.Errorf("failed to approve TP set: %w", err)
	}

	return &ApproveTPSetResponse{
		TPSetID:    tpSet.ID,
		Status:     tpSet.Status,
		ApprovedBy: cmd.ApproverID,
		ApprovedAt: now.Format("2006-01-02T15:04:05Z"),
	}, nil
}

// GetTPSetQuery represents the query to get a TP Set
type GetTPSetQuery struct {
	TPSetID string
	UserID  string // Authenticated user ID
}

// GetTPSetResponse represents the response for getting a TP Set
type GetTPSetResponse struct {
	TPSet *domain.TPSet
}

// GetTPSet retrieves a TP Set
// Orchestrates: authorization, school scope validation
func (s *TPSetApplicationService) GetTPSet(ctx context.Context, query *GetTPSetQuery) (*GetTPSetResponse, error) {
	// 1. Authorization: Get user
	user, err := s.userRepo.GetByID(ctx, query.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// 2. Load TP Set
	tpSet, err := s.tpRepo.GetTPSetByID(ctx, query.TPSetID)
	if err != nil {
		return nil, fmt.Errorf("TP set not found: %w", err)
	}

	// 3. School scope: Validate user belongs to same school as TP Set
	if user.RoleID != "SYSTEM_ADMIN" {
		owner, err := s.userRepo.GetByID(ctx, tpSet.GeneratedBy)
		if err != nil {
			return nil, fmt.Errorf("failed to get TP set owner: %w", err)
		}
		if owner.SchoolID != nil && user.SchoolID != nil && *user.SchoolID != *owner.SchoolID {
			return nil, fmt.Errorf("cross-school access not allowed")
		}
	}

	return &GetTPSetResponse{
		TPSet: tpSet,
	}, nil
}

// ListTPSetsQuery represents the query to list TP Sets
type ListTPSetsQuery struct {
	CPID     *string
	Status   *domain.WorkflowStatus
	UserID   string // Authenticated user ID
	Page     int
	PageSize int
}

// ListTPSetsResponse represents the response for listing TP Sets
type ListTPSetsResponse struct {
	TPSets   []*domain.TPSet
	Total    int
	Page     int
	PageSize int
}

// ListTPSets lists TP Sets
// Orchestrates: authorization, school scope filtering
func (s *TPSetApplicationService) ListTPSets(ctx context.Context, query *ListTPSetsQuery) (*ListTPSetsResponse, error) {
	// 1. Authorization: Get user
	user, err := s.userRepo.GetByID(ctx, query.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// 2. School scope: Filter by user's school unless System Admin
	var schoolID *string
	if user.RoleID != "SYSTEM_ADMIN" && user.SchoolID != nil {
		schoolID = user.SchoolID
	}

	limit := query.PageSize
	offset := (query.Page - 1) * query.PageSize

	tpSets, err := s.tpRepo.ListTPSets(ctx, query.CPID, query.Status, schoolID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list TP sets: %w", err)
	}

	return &ListTPSetsResponse{
		TPSets:   tpSets,
		Total:    len(tpSets),
		Page:     query.Page,
		PageSize: query.PageSize,
	}, nil
}

// Helper functions

func generateID() string {
	return fmt.Sprintf("tp-set-%d", now().UnixNano())
}

func now() time.Time {
	return time.Now().UTC()
}

package repository

import (
	"context"

	"github.com/nusa/backend/internal/domain"
)

// ITPSetRepository defines the interface for TPSet repository operations
// This follows the Repository Pattern and enables dependency injection
type ITPSetRepository interface {
	// CreateTPSet creates a new TP Set
	CreateTPSet(ctx context.Context, tpSet *domain.TPSet) error

	// GetTPSetByID retrieves a TP Set by ID
	GetTPSetByID(ctx context.Context, id string) (*domain.TPSet, error)

	// GetTPSetByCPAndVersion retrieves a TP Set by CP ID and version number
	GetTPSetByCPAndVersion(ctx context.Context, cpID string, versionNo int) (*domain.TPSet, error)

	// ListTPSets retrieves TP Sets with optional filters
	// Supports filtering by cpID, status, and school scope
	ListTPSets(ctx context.Context, cpID *string, status *domain.WorkflowStatus, schoolID *string, limit, offset int) ([]*domain.TPSet, error)

	// UpdateTPSet updates a TP Set
	UpdateTPSet(ctx context.Context, tpSet *domain.TPSet) error

	// UpdateTPSetStatus updates only the status of a TP Set
	UpdateTPSetStatus(ctx context.Context, id string, status domain.WorkflowStatus, approvedBy *string, approvedAt *interface{}) error

	// DeleteTPSet soft deletes a TP Set (if soft delete is supported)
	DeleteTPSet(ctx context.Context, id string) error
}

// ITPRepository defines the interface for TP (Teaching Plan Item) repository operations
type ITPRepository interface {
	// CreateTP creates a new TP
	CreateTP(ctx context.Context, tp *domain.TP) error

	// GetTPByID retrieves a TP by ID
	GetTPByID(ctx context.Context, id string) (*domain.TP, error)

	// ListTPsBySet retrieves TPs by TP Set ID
	ListTPsBySet(ctx context.Context, tpSetID string) ([]*domain.TP, error)

	// ListTPs retrieves TPs with optional filters
	// Supports filtering by tpSetID, cpID, status, and school scope
	ListTPs(ctx context.Context, tpSetID, cpID *string, status *domain.WorkflowStatus, schoolID *string, limit, offset int) ([]*domain.TP, error)

	// UpdateTP updates a TP
	UpdateTP(ctx context.Context, tp *domain.TP) error

	// HasDownstreamAssessments checks if a TP has downstream assessments
	HasDownstreamAssessments(ctx context.Context, tpID string) (bool, error)

	// GetTPVersionHistory retrieves all versions of a TP for a given TP set and sequence
	GetTPVersionHistory(ctx context.Context, tpSetID string, sequenceNumber int) ([]*domain.TP, error)

	// DeleteTP soft deletes a TP (if soft delete is supported)
	DeleteTP(ctx context.Context, id string) error
}

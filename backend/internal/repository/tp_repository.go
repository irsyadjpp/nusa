package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/nusa/backend/internal/domain"
)

// TPRepository handles database operations for TP entities
type TPRepository struct {
	db *sqlx.DB
}

// NewTPRepository creates a new TP repository
func NewTPRepository(db *sqlx.DB) *TPRepository {
	return &TPRepository{db: db}
}

// ==================== TP Set Operations ====================

// CreateTPSet creates a new TP Set
func (r *TPRepository) CreateTPSet(ctx context.Context, tpSet *domain.TPSet) error {
	query := `
		INSERT INTO tp_sets (id, cp_id, version_no, status, generation_source, generation_reason, 
		                    generated_by, ai_generation_id, approved_by, approved_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	_, err := r.db.ExecContext(ctx, query,
		tpSet.ID, tpSet.CPID, tpSet.VersionNo, tpSet.Status, tpSet.GenerationSource,
		tpSet.GenerationReason, tpSet.GeneratedBy, tpSet.AIGenerationID, tpSet.ApprovedBy, tpSet.ApprovedAt)
	return err
}

// GetTPSetByID retrieves a TP Set by ID
func (r *TPRepository) GetTPSetByID(ctx context.Context, id string) (*domain.TPSet, error) {
	query := `
		SELECT id, cp_id, version_no, status, generation_source, generation_reason, 
		       generated_by, ai_generation_id, approved_by, approved_at, created_at, updated_at
		FROM tp_sets WHERE id = $1
	`

	var tpSet domain.TPSet
	var generationReason, aiGenerationID, approvedBy sql.NullString
	var approvedAt sql.NullTime

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&tpSet.ID, &tpSet.CPID, &tpSet.VersionNo, &tpSet.Status, &tpSet.GenerationSource,
		&generationReason, &tpSet.GeneratedBy, &aiGenerationID, &approvedBy, &approvedAt,
		&tpSet.CreatedAt, &tpSet.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("tp set not found")
	}
	if err != nil {
		return nil, err
	}

	if generationReason.Valid {
		tpSet.GenerationReason = &generationReason.String
	}
	if aiGenerationID.Valid {
		tpSet.AIGenerationID = &aiGenerationID.String
	}
	if approvedBy.Valid {
		tpSet.ApprovedBy = &approvedBy.String
	}
	if approvedAt.Valid {
		tpSet.ApprovedAt = &approvedAt.Time
	}

	return &tpSet, nil
}

// GetTPSetByCPAndVersion retrieves a TP Set by CP ID and version number
func (r *TPRepository) GetTPSetByCPAndVersion(ctx context.Context, cpID string, versionNo int) (*domain.TPSet, error) {
	query := `
		SELECT id, cp_id, version_no, status, generation_source, generation_reason, 
		       generated_by, ai_generation_id, approved_by, approved_at, created_at, updated_at
		FROM tp_sets WHERE cp_id = $1 AND version_no = $2
	`

	var tpSet domain.TPSet
	var generationReason, aiGenerationID, approvedBy sql.NullString
	var approvedAt sql.NullTime

	err := r.db.QueryRowContext(ctx, query, cpID, versionNo).Scan(
		&tpSet.ID, &tpSet.CPID, &tpSet.VersionNo, &tpSet.Status, &tpSet.GenerationSource,
		&generationReason, &tpSet.GeneratedBy, &aiGenerationID, &approvedBy, &approvedAt,
		&tpSet.CreatedAt, &tpSet.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("tp set not found")
	}
	if err != nil {
		return nil, err
	}

	if generationReason.Valid {
		tpSet.GenerationReason = &generationReason.String
	}
	if aiGenerationID.Valid {
		tpSet.AIGenerationID = &aiGenerationID.String
	}
	if approvedBy.Valid {
		tpSet.ApprovedBy = &approvedBy.String
	}
	if approvedAt.Valid {
		tpSet.ApprovedAt = &approvedAt.Time
	}

	return &tpSet, nil
}

// ListTPSets retrieves TP Sets with optional filters
func (r *TPRepository) ListTPSets(ctx context.Context, cpID *string, status *domain.WorkflowStatus, limit, offset int) ([]*domain.TPSet, error) {
	query := `
		SELECT id, cp_id, version_no, status, generation_source, generation_reason, 
		       generated_by, ai_generation_id, approved_by, approved_at, created_at, updated_at
		FROM tp_sets
		WHERE 1=1
	`

	args := []interface{}{}
	argIndex := 1

	if cpID != nil {
		query += fmt.Sprintf(" AND cp_id = $%d", argIndex)
		args = append(args, *cpID)
		argIndex++
	}

	if status != nil {
		query += fmt.Sprintf(" AND status = $%d", argIndex)
		args = append(args, *status)
		argIndex++
	}

	query += " ORDER BY created_at DESC"

	if limit > 0 {
		query += fmt.Sprintf(" LIMIT $%d", argIndex)
		args = append(args, limit)
		argIndex++
	}

	if offset > 0 {
		query += fmt.Sprintf(" OFFSET $%d", argIndex)
		args = append(args, offset)
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tpSets []*domain.TPSet
	for rows.Next() {
		var tpSet domain.TPSet
		var generationReason, aiGenerationID, approvedBy sql.NullString
		var approvedAt sql.NullTime

		err := rows.Scan(
			&tpSet.ID, &tpSet.CPID, &tpSet.VersionNo, &tpSet.Status, &tpSet.GenerationSource,
			&generationReason, &tpSet.GeneratedBy, &aiGenerationID, &approvedBy, &approvedAt,
			&tpSet.CreatedAt, &tpSet.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		if generationReason.Valid {
			tpSet.GenerationReason = &generationReason.String
		}
		if aiGenerationID.Valid {
			tpSet.AIGenerationID = &aiGenerationID.String
		}
		if approvedBy.Valid {
			tpSet.ApprovedBy = &approvedBy.String
		}
		if approvedAt.Valid {
			tpSet.ApprovedAt = &approvedAt.Time
		}

		tpSets = append(tpSets, &tpSet)
	}

	return tpSets, nil
}

// UpdateTPSet updates a TP Set
func (r *TPRepository) UpdateTPSet(ctx context.Context, tpSet *domain.TPSet) error {
	query := `
		UPDATE tp_sets 
		SET status = $2, generation_reason = $3, approved_by = $4, approved_at = $5, updated_at = NOW()
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query,
		tpSet.ID, tpSet.Status, tpSet.GenerationReason, tpSet.ApprovedBy, tpSet.ApprovedAt)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("tp set not found")
	}

	return nil
}

// UpdateTPSetStatus updates only the status of a TP Set
func (r *TPRepository) UpdateTPSetStatus(ctx context.Context, id string, status domain.WorkflowStatus, approvedBy *string, approvedAt *interface{}) error {
	query := `
		UPDATE tp_sets 
		SET status = $2, approved_by = $3, approved_at = $4, updated_at = NOW()
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query, id, status, approvedBy, approvedAt)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("tp set not found")
	}

	return nil
}

// ==================== TP Item Operations ====================

// CreateTP creates a new TP
func (r *TPRepository) CreateTP(ctx context.Context, tp *domain.TP) error {
	query := `
		INSERT INTO tp (id, tp_set_id, sequence_number, cp_id, subject_id, phase_id, element_id, 
		              subelement_id, user_id, status, title, learning_objectives, time_allocation, 
		              prerequisites, estimated_weeks)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
	`

	_, err := r.db.ExecContext(ctx, query,
		tp.ID, tp.TPSetID, tp.SequenceNumber, tp.CPID, tp.SubjectID, tp.PhaseID, tp.ElementID,
		tp.SubelementID, tp.UserID, tp.Status, tp.Title, tp.LearningObjectives, tp.TimeAllocation,
		tp.Prerequisites, tp.EstimatedWeeks)
	return err
}

// GetTPByID retrieves a TP by ID
func (r *TPRepository) GetTPByID(ctx context.Context, id string) (*domain.TP, error) {
	query := `
		SELECT id, tp_set_id, sequence_number, cp_id, subject_id, phase_id, element_id, 
		       subelement_id, user_id, status, title, learning_objectives, time_allocation, 
		       prerequisites, estimated_weeks, created_at, updated_at
		FROM tp WHERE id = $1
	`

	var tp domain.TP
	var title, prerequisites sql.NullString
	var estimatedWeeks sql.NullInt32

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&tp.ID, &tp.TPSetID, &tp.SequenceNumber, &tp.CPID, &tp.SubjectID, &tp.PhaseID, &tp.ElementID,
		&tp.SubelementID, &tp.UserID, &tp.Status, &title, &tp.LearningObjectives, &tp.TimeAllocation,
		&prerequisites, &estimatedWeeks, &tp.CreatedAt, &tp.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("tp not found")
	}
	if err != nil {
		return nil, err
	}

	if title.Valid {
		tp.Title = &title.String
	}
	if prerequisites.Valid {
		tp.Prerequisites = prerequisites.String
	}
	if estimatedWeeks.Valid {
		weeks := int(estimatedWeeks.Int32)
		tp.EstimatedWeeks = &weeks
	}

	return &tp, nil
}

// ListTPsBySet retrieves TPs by TP Set ID
func (r *TPRepository) ListTPsBySet(ctx context.Context, tpSetID string) ([]*domain.TP, error) {
	query := `
		SELECT id, tp_set_id, sequence_number, cp_id, subject_id, phase_id, element_id, 
		       subelement_id, user_id, status, title, learning_objectives, time_allocation, 
		       prerequisites, estimated_weeks, created_at, updated_at
		FROM tp WHERE tp_set_id = $1 ORDER BY sequence_number ASC
	`

	rows, err := r.db.QueryContext(ctx, query, tpSetID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tps []*domain.TP
	for rows.Next() {
		var tp domain.TP
		var title, prerequisites sql.NullString
		var estimatedWeeks sql.NullInt32

		err := rows.Scan(
			&tp.ID, &tp.TPSetID, &tp.SequenceNumber, &tp.CPID, &tp.SubjectID, &tp.PhaseID, &tp.ElementID,
			&tp.SubelementID, &tp.UserID, &tp.Status, &title, &tp.LearningObjectives, &tp.TimeAllocation,
			&prerequisites, &estimatedWeeks, &tp.CreatedAt, &tp.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		if title.Valid {
			tp.Title = &title.String
		}
		if prerequisites.Valid {
			tp.Prerequisites = prerequisites.String
		}
		if estimatedWeeks.Valid {
			weeks := int(estimatedWeeks.Int32)
			tp.EstimatedWeeks = &weeks
		}

		tps = append(tps, &tp)
	}

	return tps, nil
}

// ListTPs retrieves TPs with optional filters
func (r *TPRepository) ListTPs(ctx context.Context, tpSetID, cpID *string, status *domain.WorkflowStatus, limit, offset int) ([]*domain.TP, error) {
	query := `
		SELECT id, tp_set_id, sequence_number, cp_id, subject_id, phase_id, element_id, 
		       subelement_id, user_id, status, title, learning_objectives, time_allocation, 
		       prerequisites, estimated_weeks, created_at, updated_at
		FROM tp
		WHERE 1=1
	`

	args := []interface{}{}
	argIndex := 1

	if tpSetID != nil {
		query += fmt.Sprintf(" AND tp_set_id = $%d", argIndex)
		args = append(args, *tpSetID)
		argIndex++
	}

	if cpID != nil {
		query += fmt.Sprintf(" AND cp_id = $%d", argIndex)
		args = append(args, *cpID)
		argIndex++
	}

	if status != nil {
		query += fmt.Sprintf(" AND status = $%d", argIndex)
		args = append(args, *status)
		argIndex++
	}

	query += " ORDER BY created_at DESC"

	if limit > 0 {
		query += fmt.Sprintf(" LIMIT $%d", argIndex)
		args = append(args, limit)
		argIndex++
	}

	if offset > 0 {
		query += fmt.Sprintf(" OFFSET $%d", argIndex)
		args = append(args, offset)
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tps []*domain.TP
	for rows.Next() {
		var tp domain.TP
		var title, prerequisites sql.NullString
		var estimatedWeeks sql.NullInt32

		err := rows.Scan(
			&tp.ID, &tp.TPSetID, &tp.SequenceNumber, &tp.CPID, &tp.SubjectID, &tp.PhaseID, &tp.ElementID,
			&tp.SubelementID, &tp.UserID, &tp.Status, &title, &tp.LearningObjectives, &tp.TimeAllocation,
			&prerequisites, &estimatedWeeks, &tp.CreatedAt, &tp.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		if title.Valid {
			tp.Title = &title.String
		}
		if prerequisites.Valid {
			tp.Prerequisites = prerequisites.String
		}
		if estimatedWeeks.Valid {
			weeks := int(estimatedWeeks.Int32)
			tp.EstimatedWeeks = &weeks
		}

		tps = append(tps, &tp)
	}

	return tps, nil
}

// UpdateTP updates a TP
func (r *TPRepository) UpdateTP(ctx context.Context, tp *domain.TP) error {
	query := `
		UPDATE tp 
		SET title = $2, learning_objectives = $3, time_allocation = $4, prerequisites = $5, 
		    estimated_weeks = $6, status = $7, updated_at = NOW()
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query,
		tp.ID, tp.Title, tp.LearningObjectives, tp.TimeAllocation, tp.Prerequisites, tp.EstimatedWeeks, tp.Status)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("tp not found")
	}

	return nil
}

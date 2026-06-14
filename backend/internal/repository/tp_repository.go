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
// Supports school scope filtering via JOIN with users table
func (r *TPRepository) ListTPSets(ctx context.Context, cpID *string, status *domain.WorkflowStatus, schoolID *string, limit, offset int) ([]*domain.TPSet, error) {
	query := `
		SELECT ts.id, ts.cp_id, ts.version_no, ts.status, ts.generation_source, ts.generation_reason, 
		       ts.generated_by, ts.ai_generation_id, ts.approved_by, ts.approved_at, ts.created_at, ts.updated_at
		FROM tp_sets ts
		WHERE 1=1
	`

	args := []interface{}{}
	argIndex := 1

	// School scope filter: JOIN with users table to filter by school_id
	if schoolID != nil {
		query += fmt.Sprintf(" AND EXISTS (SELECT 1 FROM users u WHERE u.id = ts.generated_by AND u.school_id = $%d)", argIndex)
		args = append(args, *schoolID)
		argIndex++
	}

	if cpID != nil {
		query += fmt.Sprintf(" AND ts.cp_id = $%d", argIndex)
		args = append(args, *cpID)
		argIndex++
	}

	if status != nil {
		query += fmt.Sprintf(" AND ts.status = $%d", argIndex)
		args = append(args, *status)
		argIndex++
	}

	query += " ORDER BY ts.created_at DESC"

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
		              prerequisites, estimated_weeks, success_criteria, version_no, is_current_version, parent_version_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)
	`

	_, err := r.db.ExecContext(ctx, query,
		tp.ID, tp.TPSetID, tp.SequenceNumber, tp.CPID, tp.SubjectID, tp.PhaseID, tp.ElementID,
		tp.SubelementID, tp.UserID, tp.Status, tp.Title, tp.LearningObjectives, tp.TimeAllocation,
		tp.Prerequisites, tp.EstimatedWeeks, tp.SuccessCriteria, tp.VersionNo, tp.IsCurrentVersion, tp.ParentVersionID)
	return err
}

// GetTPByID retrieves a TP by ID
func (r *TPRepository) GetTPByID(ctx context.Context, id string) (*domain.TP, error) {
	query := `
		SELECT id, tp_set_id, sequence_number, cp_id, subject_id, phase_id, element_id,
		       subelement_id, user_id, status, title, learning_objectives, time_allocation,
		       prerequisites, estimated_weeks, success_criteria, version_no, is_current_version, parent_version_id, created_at, updated_at
		FROM tp WHERE id = $1
	`

	var tp domain.TP
	var title, prerequisites, parentVersionID sql.NullString
	var estimatedWeeks sql.NullInt32

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&tp.ID, &tp.TPSetID, &tp.SequenceNumber, &tp.CPID, &tp.SubjectID, &tp.PhaseID, &tp.ElementID,
		&tp.SubelementID, &tp.UserID, &tp.Status, &title, &tp.LearningObjectives, &tp.TimeAllocation,
		&prerequisites, &estimatedWeeks, &tp.SuccessCriteria, &tp.VersionNo, &tp.IsCurrentVersion, &parentVersionID, &tp.CreatedAt, &tp.UpdatedAt,
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
	if parentVersionID.Valid {
		tp.ParentVersionID = &parentVersionID.String
	}

	return &tp, nil
}

// ListTPsBySet retrieves TPs by TP Set ID
func (r *TPRepository) ListTPsBySet(ctx context.Context, tpSetID string) ([]*domain.TP, error) {
	query := `
		SELECT id, tp_set_id, sequence_number, cp_id, subject_id, phase_id, element_id, 
		       subelement_id, user_id, status, title, learning_objectives, time_allocation, 
		       prerequisites, estimated_weeks, success_criteria, created_at, updated_at
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
			&prerequisites, &estimatedWeeks, &tp.SuccessCriteria, &tp.CreatedAt, &tp.UpdatedAt,
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
// Supports school scope filtering via JOIN with users table
func (r *TPRepository) ListTPs(ctx context.Context, tpSetID, cpID *string, status *domain.WorkflowStatus, schoolID *string, limit, offset int) ([]*domain.TP, error) {
	query := `
		SELECT t.id, t.tp_set_id, t.sequence_number, t.cp_id, t.subject_id, t.phase_id, t.element_id, 
		       t.subelement_id, t.user_id, t.status, t.title, t.learning_objectives, t.time_allocation, 
		       t.prerequisites, t.estimated_weeks, t.success_criteria, t.created_at, t.updated_at
		FROM tp t
		WHERE 1=1
	`

	args := []interface{}{}
	argIndex := 1

	// School scope filter: JOIN with users table to filter by school_id
	if schoolID != nil {
		query += fmt.Sprintf(" AND EXISTS (SELECT 1 FROM users u WHERE u.id = t.user_id AND u.school_id = $%d)", argIndex)
		args = append(args, *schoolID)
		argIndex++
	}

	if tpSetID != nil {
		query += fmt.Sprintf(" AND t.tp_set_id = $%d", argIndex)
		args = append(args, *tpSetID)
		argIndex++
	}

	if cpID != nil {
		query += fmt.Sprintf(" AND t.cp_id = $%d", argIndex)
		args = append(args, *cpID)
		argIndex++
	}

	if status != nil {
		query += fmt.Sprintf(" AND t.status = $%d", argIndex)
		args = append(args, *status)
		argIndex++
	}

	query += " ORDER BY t.created_at DESC"

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
			&prerequisites, &estimatedWeeks, &tp.SuccessCriteria, &tp.CreatedAt, &tp.UpdatedAt,
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
		    estimated_weeks = $6, status = $7, success_criteria = $8, version_no = $9, is_current_version = $10, parent_version_id = $11, updated_at = NOW()
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query,
		tp.ID, tp.Title, tp.LearningObjectives, tp.TimeAllocation, tp.Prerequisites, tp.EstimatedWeeks, tp.Status, tp.SuccessCriteria, tp.VersionNo, tp.IsCurrentVersion, tp.ParentVersionID)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("tp not found")
	}

	return nil
}

// HasDownstreamAssessments checks if a TP has downstream assessments
func (r *TPRepository) HasDownstreamAssessments(ctx context.Context, tpID string) (bool, error) {
	query := `
		SELECT COUNT(*) FROM assessments WHERE tp_version_no IN (
			SELECT version_no FROM tp WHERE id = $1
		)
	`

	var count int
	err := r.db.QueryRowContext(ctx, query, tpID).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to check downstream assessments: %w", err)
	}

	return count > 0, nil
}

// DeleteTPSet soft deletes a TP Set
// Note: Current schema does not support soft delete (no deleted_at column)
// This method is provided for future soft delete support
func (r *TPRepository) DeleteTPSet(ctx context.Context, id string) error {
	// Current implementation: Hard delete
	// Future implementation: UPDATE tp_sets SET deleted_at = NOW() WHERE id = $1
	query := `DELETE FROM tp_sets WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("tp set not found")
	}

	return nil
}

// DeleteTP soft deletes a TP
// Note: Current schema does not support soft delete (no deleted_at column)
// This method is provided for future soft delete support
func (r *TPRepository) DeleteTP(ctx context.Context, id string) error {
	// Current implementation: Hard delete
	// Future implementation: UPDATE tp SET deleted_at = NOW() WHERE id = $1
	query := `DELETE FROM tp WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("tp not found")
	}

	return nil
}

// ==================== TP Versioning Operations ====================

// GetTPVersions retrieves all versions of a specific TP
func (r *TPRepository) GetTPVersions(ctx context.Context, tpID string) ([]*domain.TP, error) {
	query := `
		SELECT id, tp_set_id, sequence_number, cp_id, subject_id, phase_id, element_id,
		       subelement_id, user_id, status, title, learning_objectives, time_allocation,
		       prerequisites, estimated_weeks, success_criteria, version_no, is_current_version,
		       parent_version_id, created_at, updated_at
		FROM tp
		WHERE id = $1 OR parent_version_id = $1
		ORDER BY version_no ASC
	`

	rows, err := r.db.QueryContext(ctx, query, tpID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tps []*domain.TP
	for rows.Next() {
		var tp domain.TP
		var title, parentVersionID sql.NullString
		var estimatedWeeks sql.NullInt64

		err := rows.Scan(
			&tp.ID, &tp.TPSetID, &tp.SequenceNumber, &tp.CPID, &tp.SubjectID, &tp.PhaseID, &tp.ElementID,
			&tp.SubelementID, &tp.UserID, &tp.Status, &title, &tp.LearningObjectives, &tp.TimeAllocation,
			&tp.Prerequisites, &estimatedWeeks, &tp.SuccessCriteria, &tp.VersionNo, &tp.IsCurrentVersion,
			&parentVersionID, &tp.CreatedAt, &tp.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		if title.Valid {
			tp.Title = &title.String
		}
		if estimatedWeeks.Valid {
			weeks := int(estimatedWeeks.Int64)
			tp.EstimatedWeeks = &weeks
		}
		if parentVersionID.Valid {
			tp.ParentVersionID = &parentVersionID.String
		}

		tps = append(tps, &tp)
	}

	return tps, nil
}

// GetTPVersionHistory retrieves the full version history for a TP Set
func (r *TPRepository) GetTPVersionHistory(ctx context.Context, tpSetID string) ([]*domain.TP, error) {
	query := `
		SELECT id, tp_set_id, sequence_number, cp_id, subject_id, phase_id, element_id,
		       subelement_id, user_id, status, title, learning_objectives, time_allocation,
		       prerequisites, estimated_weeks, success_criteria, version_no, is_current_version,
		       parent_version_id, created_at, updated_at
		FROM tp
		WHERE tp_set_id = $1
		ORDER BY version_no ASC, sequence_number ASC
	`

	rows, err := r.db.QueryContext(ctx, query, tpSetID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tps []*domain.TP
	for rows.Next() {
		var tp domain.TP
		var title, parentVersionID sql.NullString
		var estimatedWeeks sql.NullInt64

		err := rows.Scan(
			&tp.ID, &tp.TPSetID, &tp.SequenceNumber, &tp.CPID, &tp.SubjectID, &tp.PhaseID, &tp.ElementID,
			&tp.SubelementID, &tp.UserID, &tp.Status, &title, &tp.LearningObjectives, &tp.TimeAllocation,
			&tp.Prerequisites, &estimatedWeeks, &tp.SuccessCriteria, &tp.VersionNo, &tp.IsCurrentVersion,
			&parentVersionID, &tp.CreatedAt, &tp.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		if title.Valid {
			tp.Title = &title.String
		}
		if estimatedWeeks.Valid {
			weeks := int(estimatedWeeks.Int64)
			tp.EstimatedWeeks = &weeks
		}
		if parentVersionID.Valid {
			tp.ParentVersionID = &parentVersionID.String
		}

		tps = append(tps, &tp)
	}

	return tps, nil
}

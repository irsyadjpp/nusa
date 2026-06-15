package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/nusa/backend/internal/domain"
)

// CPAlignmentRepository handles database operations for CP alignments
type CPAlignmentRepository struct {
	db *sqlx.DB
}

// NewCPAlignmentRepository creates a new CP alignment repository
func NewCPAlignmentRepository(db *sqlx.DB) *CPAlignmentRepository {
	return &CPAlignmentRepository{db: db}
}

// CreateCPAlignment creates a new CP alignment
func (r *CPAlignmentRepository) CreateCPAlignment(ctx context.Context, cpa *domain.CPAlignment) error {
	query := `
		INSERT INTO cp_alignments (id, curriculum_subject_id, graduate_profile_dimension_id, alignment_description, is_active, created_by, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err := r.db.ExecContext(ctx, query,
		cpa.ID, cpa.CurriculumSubjectID, cpa.GraduateProfileDimensionID,
		cpa.AlignmentDescription, cpa.IsActive, cpa.CreatedBy, cpa.CreatedAt, cpa.UpdatedAt)
	return err
}

// GetCPAlignmentByID retrieves a CP alignment by ID
func (r *CPAlignmentRepository) GetCPAlignmentByID(ctx context.Context, id string) (*domain.CPAlignment, error) {
	query := `
		SELECT id, curriculum_subject_id, graduate_profile_dimension_id, alignment_description, is_active, created_by, created_at, updated_at, updated_by
		FROM cp_alignments WHERE id = $1
	`

	var cpa domain.CPAlignment
	var alignmentDescription, updatedBy sql.NullString

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&cpa.ID, &cpa.CurriculumSubjectID, &cpa.GraduateProfileDimensionID,
		&alignmentDescription, &cpa.IsActive, &cpa.CreatedBy, &cpa.CreatedAt, &cpa.UpdatedAt, &updatedBy,
	)
	if err != nil {
		return nil, err
	}

	if alignmentDescription.Valid {
		cpa.AlignmentDescription = &alignmentDescription.String
	}
	if updatedBy.Valid {
		cpa.UpdatedBy = &updatedBy.String
	}

	return &cpa, nil
}

// GetCPAlignmentsByCurriculumSubjectID retrieves all CP alignments for a curriculum subject
func (r *CPAlignmentRepository) GetCPAlignmentsByCurriculumSubjectID(ctx context.Context, curriculumSubjectID string) ([]*domain.CPAlignment, error) {
	query := `
		SELECT id, curriculum_subject_id, graduate_profile_dimension_id, alignment_description, is_active, created_by, created_at, updated_at, updated_by
		FROM cp_alignments WHERE curriculum_subject_id = $1 AND is_active = true
	`

	rows, err := r.db.QueryContext(ctx, query, curriculumSubjectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var alignments []*domain.CPAlignment
	for rows.Next() {
		var cpa domain.CPAlignment
		var alignmentDescription, updatedBy sql.NullString

		err := rows.Scan(
			&cpa.ID, &cpa.CurriculumSubjectID, &cpa.GraduateProfileDimensionID,
			&alignmentDescription, &cpa.IsActive, &cpa.CreatedBy, &cpa.CreatedAt, &cpa.UpdatedAt, &updatedBy,
		)
		if err != nil {
			return nil, err
		}

		if alignmentDescription.Valid {
			cpa.AlignmentDescription = &alignmentDescription.String
		}
		if updatedBy.Valid {
			cpa.UpdatedBy = &updatedBy.String
		}

		alignments = append(alignments, &cpa)
	}
	return alignments, nil
}

// GetCPAlignmentsByGraduateProfileDimensionID retrieves all CP alignments for a graduate profile dimension
func (r *CPAlignmentRepository) GetCPAlignmentsByGraduateProfileDimensionID(ctx context.Context, graduateProfileDimensionID string) ([]*domain.CPAlignment, error) {
	query := `
		SELECT id, curriculum_subject_id, graduate_profile_dimension_id, alignment_description, is_active, created_by, created_at, updated_at, updated_by
		FROM cp_alignments WHERE graduate_profile_dimension_id = $1 AND is_active = true
	`

	rows, err := r.db.QueryContext(ctx, query, graduateProfileDimensionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var alignments []*domain.CPAlignment
	for rows.Next() {
		var cpa domain.CPAlignment
		var alignmentDescription, updatedBy sql.NullString

		err := rows.Scan(
			&cpa.ID, &cpa.CurriculumSubjectID, &cpa.GraduateProfileDimensionID,
			&alignmentDescription, &cpa.IsActive, &cpa.CreatedBy, &cpa.CreatedAt, &cpa.UpdatedAt, &updatedBy,
		)
		if err != nil {
			return nil, err
		}

		if alignmentDescription.Valid {
			cpa.AlignmentDescription = &alignmentDescription.String
		}
		if updatedBy.Valid {
			cpa.UpdatedBy = &updatedBy.String
		}

		alignments = append(alignments, &cpa)
	}
	return alignments, nil
}

// GetAllCPAlignments retrieves all CP alignments
func (r *CPAlignmentRepository) GetAllCPAlignments(ctx context.Context) ([]*domain.CPAlignment, error) {
	query := `
		SELECT id, curriculum_subject_id, graduate_profile_dimension_id, alignment_description, is_active, created_by, created_at, updated_at, updated_by
		FROM cp_alignments WHERE is_active = true
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var alignments []*domain.CPAlignment
	for rows.Next() {
		var cpa domain.CPAlignment
		var alignmentDescription, updatedBy sql.NullString

		err := rows.Scan(
			&cpa.ID, &cpa.CurriculumSubjectID, &cpa.GraduateProfileDimensionID,
			&alignmentDescription, &cpa.IsActive, &cpa.CreatedBy, &cpa.CreatedAt, &cpa.UpdatedAt, &updatedBy,
		)
		if err != nil {
			return nil, err
		}

		if alignmentDescription.Valid {
			cpa.AlignmentDescription = &alignmentDescription.String
		}
		if updatedBy.Valid {
			cpa.UpdatedBy = &updatedBy.String
		}

		alignments = append(alignments, &cpa)
	}
	return alignments, nil
}

// UpdateCPAlignment updates a CP alignment
func (r *CPAlignmentRepository) UpdateCPAlignment(ctx context.Context, cpa *domain.CPAlignment) error {
	query := `
		UPDATE cp_alignments
		SET alignment_description = $1, is_active = $2, updated_at = $3, updated_by = $4
		WHERE id = $5
	`

	result, err := r.db.ExecContext(ctx, query,
		cpa.AlignmentDescription, cpa.IsActive, cpa.UpdatedAt, cpa.UpdatedBy, cpa.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("CP alignment not found")
	}
	return nil
}

// DeleteCPAlignment deletes a CP alignment
func (r *CPAlignmentRepository) DeleteCPAlignment(ctx context.Context, id string) error {
	query := `DELETE FROM cp_alignments WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("CP alignment not found")
	}
	return nil
}

// CheckAlignmentExists checks if a CP alignment already exists for the combination
func (r *CPAlignmentRepository) CheckAlignmentExists(ctx context.Context, curriculumSubjectID, graduateProfileDimensionID string, excludeID string) (bool, error) {
	query := `
		SELECT COUNT(*) FROM cp_alignments
		WHERE curriculum_subject_id = $1 AND graduate_profile_dimension_id = $2 AND is_active = true
	`
	args := []interface{}{curriculumSubjectID, graduateProfileDimensionID}

	if excludeID != "" {
		query += ` AND id != $3`
		args = append(args, excludeID)
	}

	var count int
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// GenerateCPAlignmentReport generates a report showing CP coverage across all graduate profile dimensions
// This is used for BR-004: Minimum CP coverage percentage
func (r *CPAlignmentRepository) GenerateCPAlignmentReport(ctx context.Context, threshold float64) ([]*domain.CPAlignmentReport, error) {
	query := `
		WITH total_cp AS (
			SELECT COUNT(*) as count FROM curriculum_subjects WHERE is_active = true
		),
		dimension_coverage AS (
			SELECT 
				gpd.id as graduate_profile_dimension_id,
				gpd.name as graduate_profile_dimension_name,
				total_cp.count as total_cp_count,
				COUNT(cpa.id) as aligned_cp_count
			FROM graduate_profile_dimensions gpd
			CROSS JOIN total_cp
			LEFT JOIN cp_alignments cpa ON cpa.graduate_profile_dimension_id = gpd.id 
				AND cpa.curriculum_subject_id IN (SELECT id FROM curriculum_subjects WHERE is_active = true)
				AND cpa.is_active = true
			WHERE gpd.is_active = true
			GROUP BY gpd.id, gpd.name, total_cp.count
		)
		SELECT 
			graduate_profile_dimension_id,
			graduate_profile_dimension_name,
			total_cp_count,
			aligned_cp_count,
			CASE 
				WHEN total_cp_count = 0 THEN 0
				ELSE (aligned_cp_count::float / total_cp_count::float) * 100
			END as coverage_percentage,
			CASE 
				WHEN total_cp_count = 0 THEN false
				ELSE (aligned_cp_count::float / total_cp_count::float) * 100 >= $1
			END as meets_threshold
		FROM dimension_coverage
		ORDER BY graduate_profile_dimension_name
	`

	rows, err := r.db.QueryContext(ctx, query, threshold)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []*domain.CPAlignmentReport
	for rows.Next() {
		var report domain.CPAlignmentReport
		err := rows.Scan(
			&report.GraduateProfileDimensionID,
			&report.GraduateProfileDimensionName,
			&report.TotalCPCount,
			&report.AlignedCPCount,
			&report.CoveragePercentage,
			&report.MeetsThreshold,
		)
		if err != nil {
			return nil, err
		}
		reports = append(reports, &report)
	}
	return reports, nil
}

// DeleteAlignmentsByCurriculumSubjectID deletes all alignments for a curriculum subject
func (r *CPAlignmentRepository) DeleteAlignmentsByCurriculumSubjectID(ctx context.Context, curriculumSubjectID string) error {
	query := `DELETE FROM cp_alignments WHERE curriculum_subject_id = $1`
	_, err := r.db.ExecContext(ctx, query, curriculumSubjectID)
	return err
}

// BulkCreateCPAlignments creates multiple CP alignments in a transaction
func (r *CPAlignmentRepository) BulkCreateCPAlignments(ctx context.Context, alignments []*domain.CPAlignment) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback() // Ignore error, transaction might already be committed
	}()

	query := `
		INSERT INTO cp_alignments (id, curriculum_subject_id, graduate_profile_dimension_id, alignment_description, is_active, created_by, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	for _, cpa := range alignments {
		_, err := tx.ExecContext(ctx, query,
			cpa.ID, cpa.CurriculumSubjectID, cpa.GraduateProfileDimensionID,
			cpa.AlignmentDescription, cpa.IsActive, cpa.CreatedBy, cpa.CreatedAt, cpa.UpdatedAt)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

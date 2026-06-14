package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/nusa/backend/internal/domain"
)

// ReportingRepository handles database operations for reporting entities
type ReportingRepository struct {
	db *sqlx.DB
}

// NewReportingRepository creates a new reporting repository
func NewReportingRepository(db *sqlx.DB) *ReportingRepository {
	return &ReportingRepository{db: db}
}

// ==================== Narrative Report Operations ====================

// CreateNarrativeReport creates a new narrative report
func (r *ReportingRepository) CreateNarrativeReport(ctx context.Context, report *domain.NarrativeReport) error {
	query := `
		INSERT INTO narrative_reports (id, student_id, user_id, status, report_period, language, content, 
		                        ai_confidence_score, ai_generated_at, ai_agent_version, version_no, is_current_version,
		                        parent_version_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	`

	_, err := r.db.ExecContext(ctx, query,
		report.ID, report.StudentID, report.UserID, report.Status, report.ReportPeriod, report.Language, report.Content,
		report.AiConfidenceScore, report.AiGeneratedAt, report.AiAgentVersion, report.VersionNo, report.IsCurrentVersion,
		report.ParentVersionID, report.CreatedAt, report.UpdatedAt)
	return err
}

// GetNarrativeReportByID retrieves a narrative report by ID
func (r *ReportingRepository) GetNarrativeReportByID(ctx context.Context, id string) (*domain.NarrativeReport, error) {
	query := `
		SELECT id, student_id, user_id, status, report_period, language, content, 
		       ai_confidence_score, ai_generated_at, ai_agent_version, version_no, is_current_version,
		       parent_version_id, created_at, updated_at, approved_at, approved_by
		FROM narrative_reports WHERE id = $1
	`

	var report domain.NarrativeReport
	var aiConfidenceScore sql.NullFloat64
	var aiGeneratedAt, aiAgentVersion, parentVersionID, approvedAt, approvedBy sql.NullString

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&report.ID, &report.StudentID, &report.UserID, &report.Status, &report.ReportPeriod, &report.Language, &report.Content,
		&aiConfidenceScore, &aiGeneratedAt, &aiAgentVersion, &report.VersionNo, &report.IsCurrentVersion,
		&parentVersionID, &report.CreatedAt, &report.UpdatedAt, &approvedAt, &approvedBy,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("narrative report not found")
	}
	if err != nil {
		return nil, err
	}

	if aiConfidenceScore.Valid {
		report.AiConfidenceScore = &aiConfidenceScore.Float64
	}
	if aiGeneratedAt.Valid {
		t := time.Time{}
		if _, err := time.Parse(time.RFC3339, aiGeneratedAt.String); err == nil {
			report.AiGeneratedAt = &t
		}
	}
	if aiAgentVersion.Valid {
		report.AiAgentVersion = &aiAgentVersion.String
	}
	if parentVersionID.Valid {
		report.ParentVersionID = &parentVersionID.String
	}
	if approvedAt.Valid {
		t := time.Time{}
		if _, err := time.Parse(time.RFC3339, approvedAt.String); err == nil {
			report.ApprovedAt = &t
		}
	}
	if approvedBy.Valid {
		report.ApprovedBy = &approvedBy.String
	}

	return &report, nil
}

// ListNarrativeReports retrieves narrative reports with optional filters
func (r *ReportingRepository) ListNarrativeReports(ctx context.Context, studentID, userID *string, language *domain.ReportLanguage, status *domain.WorkflowStatus, limit, offset int) ([]*domain.NarrativeReport, error) {
	query := `
		SELECT id, student_id, user_id, status, report_period, language, content, 
		       ai_confidence_score, ai_generated_at, ai_agent_version, version_no, is_current_version,
		       parent_version_id, created_at, updated_at, approved_at, approved_by
		FROM narrative_reports
		WHERE 1=1
	`

	args := []interface{}{}
	argIndex := 1

	if studentID != nil {
		query += fmt.Sprintf(" AND student_id = $%d", argIndex)
		args = append(args, *studentID)
		argIndex++
	}

	if userID != nil {
		query += fmt.Sprintf(" AND user_id = $%d", argIndex)
		args = append(args, *userID)
		argIndex++
	}

	if language != nil {
		query += fmt.Sprintf(" AND language = $%d", argIndex)
		args = append(args, *language)
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

	var reports []*domain.NarrativeReport
	for rows.Next() {
		var report domain.NarrativeReport
		var aiConfidenceScore sql.NullFloat64
		var aiGeneratedAt, aiAgentVersion, parentVersionID, approvedAt, approvedBy sql.NullString

		err := rows.Scan(
			&report.ID, &report.StudentID, &report.UserID, &report.Status, &report.ReportPeriod, &report.Language, &report.Content,
			&aiConfidenceScore, &aiGeneratedAt, &aiAgentVersion, &report.VersionNo, &report.IsCurrentVersion,
			&parentVersionID, &report.CreatedAt, &report.UpdatedAt, &approvedAt, &approvedBy,
		)
		if err != nil {
			return nil, err
		}

		if aiConfidenceScore.Valid {
			report.AiConfidenceScore = &aiConfidenceScore.Float64
		}
		if aiGeneratedAt.Valid {
			t := time.Time{}
			if _, err := time.Parse(time.RFC3339, aiGeneratedAt.String); err == nil {
				report.AiGeneratedAt = &t
			}
		}
		if aiAgentVersion.Valid {
			report.AiAgentVersion = &aiAgentVersion.String
		}
		if parentVersionID.Valid {
			report.ParentVersionID = &parentVersionID.String
		}
		if approvedAt.Valid {
			t := time.Time{}
			if _, err := time.Parse(time.RFC3339, approvedAt.String); err == nil {
				report.ApprovedAt = &t
			}
		}
		if approvedBy.Valid {
			report.ApprovedBy = &approvedBy.String
		}

		reports = append(reports, &report)
	}

	return reports, nil
}

// UpdateNarrativeReport updates a narrative report
func (r *ReportingRepository) UpdateNarrativeReport(ctx context.Context, report *domain.NarrativeReport) error {
	query := `
		UPDATE narrative_reports 
		SET report_period = $2, content = $3, status = $4, updated_at = NOW()
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query,
		report.ID, report.ReportPeriod, report.Content, report.Status)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("narrative report not found")
	}

	return nil
}

// DeleteNarrativeReport deletes a narrative report
func (r *ReportingRepository) DeleteNarrativeReport(ctx context.Context, id string) error {
	query := `DELETE FROM narrative_reports WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("narrative report not found")
	}

	return nil
}

// UpdateAchievementData updates the achievement data for a narrative report
func (r *ReportingRepository) UpdateAchievementData(ctx context.Context, reportID string, achievementData interface{}) error {
	query := `
		UPDATE narrative_reports 
		SET achievement_data = $2, last_achievement_calculated_at = NOW()
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query, reportID, achievementData)
	if err != nil {
		return fmt.Errorf("failed to update achievement data: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("narrative report not found")
	}

	return nil
}

// UpdateClassID updates the class_id for a narrative report
func (r *ReportingRepository) UpdateClassID(ctx context.Context, reportID, classID string) error {
	query := `
		UPDATE narrative_reports 
		SET class_id = $1
		WHERE id = $2
	`

	result, err := r.db.ExecContext(ctx, query, classID, reportID)
	if err != nil {
		return fmt.Errorf("failed to update class_id: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("narrative report not found")
	}

	return nil
}

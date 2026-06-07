package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/nusa/backend/internal/domain"
)

// AssessmentRepository handles database operations for assessment entities
type AssessmentRepository struct {
	db *sqlx.DB
}

// NewAssessmentRepository creates a new assessment repository
func NewAssessmentRepository(db *sqlx.DB) *AssessmentRepository {
	return &AssessmentRepository{db: db}
}

// ==================== Assessment Operations ====================

// CreateAssessment creates a new assessment
func (r *AssessmentRepository) CreateAssessment(ctx context.Context, assessment *domain.Assessment) error {
	query := `
		INSERT INTO assessments (id, modul_ajar_id, user_id, assessment_type, status, assessment_items, 
		                    answer_key, scoring_guidelines, ai_confidence_score, ai_generated_at, ai_agent_version,
		                    version_no, is_current_version, parent_version_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
	`

	_, err := r.db.ExecContext(ctx, query,
		assessment.ID, assessment.ModulAjarID, assessment.UserID, assessment.AssessmentType, assessment.Status,
		assessment.AssessmentItems, assessment.AnswerKey, assessment.ScoringGuidelines, assessment.AiConfidenceScore,
		assessment.AiGeneratedAt, assessment.AiAgentVersion, assessment.VersionNo, assessment.IsCurrentVersion,
		assessment.ParentVersionID, assessment.CreatedAt, assessment.UpdatedAt)
	return err
}

// GetAssessmentByID retrieves an assessment by ID
func (r *AssessmentRepository) GetAssessmentByID(ctx context.Context, id string) (*domain.Assessment, error) {
	query := `
		SELECT id, modul_ajar_id, user_id, assessment_type, status, assessment_items, 
		       answer_key, scoring_guidelines, ai_confidence_score, ai_generated_at, ai_agent_version,
		       version_no, is_current_version, parent_version_id, created_at, updated_at, 
		       approved_at, approved_by
		FROM assessments WHERE id = $1
	`

	var assessment domain.Assessment
	var aiConfidenceScore sql.NullFloat64
	var aiGeneratedAt, aiAgentVersion, parentVersionID, approvedAt, approvedBy sql.NullString

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&assessment.ID, &assessment.ModulAjarID, &assessment.UserID, &assessment.AssessmentType, &assessment.Status,
		&assessment.AssessmentItems, &assessment.AnswerKey, &assessment.ScoringGuidelines, &aiConfidenceScore,
		&aiGeneratedAt, &aiAgentVersion, &assessment.VersionNo, &assessment.IsCurrentVersion, &parentVersionID,
		&assessment.CreatedAt, &assessment.UpdatedAt, &approvedAt, &approvedBy,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("assessment not found")
	}
	if err != nil {
		return nil, err
	}

	if aiConfidenceScore.Valid {
		assessment.AiConfidenceScore = &aiConfidenceScore.Float64
	}
	if aiGeneratedAt.Valid {
		t := time.Time{}
		if _, err := time.Parse(time.RFC3339, aiGeneratedAt.String); err == nil {
			assessment.AiGeneratedAt = &t
		}
	}
	if aiAgentVersion.Valid {
		assessment.AiAgentVersion = &aiAgentVersion.String
	}
	if parentVersionID.Valid {
		assessment.ParentVersionID = &parentVersionID.String
	}
	if approvedAt.Valid {
		t := time.Time{}
		if _, err := time.Parse(time.RFC3339, approvedAt.String); err == nil {
			assessment.ApprovedAt = &t
		}
	}
	if approvedBy.Valid {
		assessment.ApprovedBy = &approvedBy.String
	}

	return &assessment, nil
}

// ListAssessments retrieves assessments with optional filters
func (r *AssessmentRepository) ListAssessments(ctx context.Context, modulAjarID *string, userID *string, assessmentType *domain.AssessmentType, status *domain.WorkflowStatus, limit, offset int) ([]*domain.Assessment, error) {
	query := `
		SELECT id, modul_ajar_id, user_id, assessment_type, status, assessment_items, 
		       answer_key, scoring_guidelines, ai_confidence_score, ai_generated_at, ai_agent_version,
		       version_no, is_current_version, parent_version_id, created_at, updated_at,
		       approved_at, approved_by
		FROM assessments
		WHERE 1=1
	`

	args := []interface{}{}
	argIndex := 1

	if modulAjarID != nil {
		query += fmt.Sprintf(" AND modul_ajar_id = $%d", argIndex)
		args = append(args, *modulAjarID)
		argIndex++
	}

	if userID != nil {
		query += fmt.Sprintf(" AND user_id = $%d", argIndex)
		args = append(args, *userID)
		argIndex++
	}

	if assessmentType != nil {
		query += fmt.Sprintf(" AND assessment_type = $%d", argIndex)
		args = append(args, *assessmentType)
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

	var assessments []*domain.Assessment
	for rows.Next() {
		var assessment domain.Assessment
		var aiConfidenceScore sql.NullFloat64
		var aiGeneratedAt, aiAgentVersion, parentVersionID, approvedAt, approvedBy sql.NullString

		err := rows.Scan(
			&assessment.ID, &assessment.ModulAjarID, &assessment.UserID, &assessment.AssessmentType, &assessment.Status,
			&assessment.AssessmentItems, &assessment.AnswerKey, &assessment.ScoringGuidelines, &aiConfidenceScore,
			&aiGeneratedAt, &aiAgentVersion, &assessment.VersionNo, &assessment.IsCurrentVersion, &parentVersionID,
			&assessment.CreatedAt, &assessment.UpdatedAt, &approvedAt, &approvedBy,
		)
		if err != nil {
			return nil, err
		}

		if aiConfidenceScore.Valid {
			assessment.AiConfidenceScore = &aiConfidenceScore.Float64
		}
		if aiGeneratedAt.Valid {
			t := time.Time{}
			if _, err := time.Parse(time.RFC3339, aiGeneratedAt.String); err == nil {
				assessment.AiGeneratedAt = &t
			}
		}
		if aiAgentVersion.Valid {
			assessment.AiAgentVersion = &aiAgentVersion.String
		}
		if parentVersionID.Valid {
			assessment.ParentVersionID = &parentVersionID.String
		}
		if approvedAt.Valid {
			t := time.Time{}
			if _, err := time.Parse(time.RFC3339, approvedAt.String); err == nil {
				assessment.ApprovedAt = &t
			}
		}
		if approvedBy.Valid {
			assessment.ApprovedBy = &approvedBy.String
		}

		assessments = append(assessments, &assessment)
	}

	return assessments, nil
}

// UpdateAssessment updates an assessment
func (r *AssessmentRepository) UpdateAssessment(ctx context.Context, assessment *domain.Assessment) error {
	query := `
		UPDATE assessments 
		SET assessment_items = $2, answer_key = $3, scoring_guidelines = $4, status = $5, updated_at = NOW()
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query,
		assessment.ID, assessment.AssessmentItems, assessment.AnswerKey, assessment.ScoringGuidelines, assessment.Status)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("assessment not found")
	}

	return nil
}

// ==================== Rubric Operations ====================

// CreateRubric creates a new rubric
func (r *AssessmentRepository) CreateRubric(ctx context.Context, rubric *domain.Rubric) error {
	query := `
		INSERT INTO rubrics (id, assessment_id, user_id, rubric_type, status, performance_criteria, 
		                 performance_levels, scoring_guidelines, ai_confidence_score, ai_generated_at, ai_agent_version,
		                 version_no, is_current_version, parent_version_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
	`

	_, err := r.db.ExecContext(ctx, query,
		rubric.ID, rubric.AssessmentID, rubric.UserID, rubric.RubricType, rubric.Status,
		rubric.PerformanceCriteria, rubric.PerformanceLevels, rubric.ScoringGuidelines, rubric.AiConfidenceScore,
		rubric.AiGeneratedAt, rubric.AiAgentVersion, rubric.VersionNo, rubric.IsCurrentVersion,
		rubric.ParentVersionID, rubric.CreatedAt, rubric.UpdatedAt)
	return err
}

// GetRubricByID retrieves a rubric by ID
func (r *AssessmentRepository) GetRubricByID(ctx context.Context, id string) (*domain.Rubric, error) {
	query := `
		SELECT id, assessment_id, user_id, rubric_type, status, performance_criteria, 
		       performance_levels, scoring_guidelines, ai_confidence_score, ai_generated_at, ai_agent_version,
		       version_no, is_current_version, parent_version_id, created_at, updated_at,
		       approved_at, approved_by
		FROM rubrics WHERE id = $1
	`

	var rubric domain.Rubric
	var aiConfidenceScore sql.NullFloat64
	var aiGeneratedAt, aiAgentVersion, parentVersionID, approvedAt, approvedBy sql.NullString

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&rubric.ID, &rubric.AssessmentID, &rubric.UserID, &rubric.RubricType, &rubric.Status,
		&rubric.PerformanceCriteria, &rubric.PerformanceLevels, &rubric.ScoringGuidelines, &aiConfidenceScore,
		&aiGeneratedAt, &aiAgentVersion, &rubric.VersionNo, &rubric.IsCurrentVersion, &parentVersionID,
		&rubric.CreatedAt, &rubric.UpdatedAt, &approvedAt, &approvedBy,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("rubric not found")
	}
	if err != nil {
		return nil, err
	}

	if aiConfidenceScore.Valid {
		rubric.AiConfidenceScore = &aiConfidenceScore.Float64
	}
	if aiGeneratedAt.Valid {
		t := time.Time{}
		if _, err := time.Parse(time.RFC3339, aiGeneratedAt.String); err == nil {
			rubric.AiGeneratedAt = &t
		}
	}
	if aiAgentVersion.Valid {
		rubric.AiAgentVersion = &aiAgentVersion.String
	}
	if parentVersionID.Valid {
		rubric.ParentVersionID = &parentVersionID.String
	}
	if approvedAt.Valid {
		t := time.Time{}
		if _, err := time.Parse(time.RFC3339, approvedAt.String); err == nil {
			rubric.ApprovedAt = &t
		}
	}
	if approvedBy.Valid {
		rubric.ApprovedBy = &approvedBy.String
	}

	return &rubric, nil
}

// ListRubrics retrieves rubrics with optional filters
func (r *AssessmentRepository) ListRubrics(ctx context.Context, assessmentID *string, userID *string, rubricType *domain.RubricType, status *domain.WorkflowStatus, limit, offset int) ([]*domain.Rubric, error) {
	query := `
		SELECT id, assessment_id, user_id, rubric_type, status, performance_criteria, 
		       performance_levels, scoring_guidelines, ai_confidence_score, ai_generated_at, ai_agent_version,
		       version_no, is_current_version, parent_version_id, created_at, updated_at,
		       approved_at, approved_by
		FROM rubrics
		WHERE 1=1
	`

	args := []interface{}{}
	argIndex := 1

	if assessmentID != nil {
		query += fmt.Sprintf(" AND assessment_id = $%d", argIndex)
		args = append(args, *assessmentID)
		argIndex++
	}

	if userID != nil {
		query += fmt.Sprintf(" AND user_id = $%d", argIndex)
		args = append(args, *userID)
		argIndex++
	}

	if rubricType != nil {
		query += fmt.Sprintf(" AND rubric_type = $%d", argIndex)
		args = append(args, *rubricType)
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

	var rubrics []*domain.Rubric
	for rows.Next() {
		var rubric domain.Rubric
		var aiConfidenceScore sql.NullFloat64
		var aiGeneratedAt, aiAgentVersion, parentVersionID, approvedAt, approvedBy sql.NullString

		err := rows.Scan(
			&rubric.ID, &rubric.AssessmentID, &rubric.UserID, &rubric.RubricType, &rubric.Status,
			&rubric.PerformanceCriteria, &rubric.PerformanceLevels, &rubric.ScoringGuidelines, &aiConfidenceScore,
			&aiGeneratedAt, &aiAgentVersion, &rubric.VersionNo, &rubric.IsCurrentVersion, &parentVersionID,
			&rubric.CreatedAt, &rubric.UpdatedAt, &approvedAt, &approvedBy,
		)
		if err != nil {
			return nil, err
		}

		if aiConfidenceScore.Valid {
			rubric.AiConfidenceScore = &aiConfidenceScore.Float64
		}
		if aiGeneratedAt.Valid {
			t := time.Time{}
			if _, err := time.Parse(time.RFC3339, aiGeneratedAt.String); err == nil {
				rubric.AiGeneratedAt = &t
			}
		}
		if aiAgentVersion.Valid {
			rubric.AiAgentVersion = &aiAgentVersion.String
		}
		if parentVersionID.Valid {
			rubric.ParentVersionID = &parentVersionID.String
		}
		if approvedAt.Valid {
			t := time.Time{}
			if _, err := time.Parse(time.RFC3339, approvedAt.String); err == nil {
				rubric.ApprovedAt = &t
			}
		}
		if approvedBy.Valid {
			rubric.ApprovedBy = &approvedBy.String
		}

		rubrics = append(rubrics, &rubric)
	}

	return rubrics, nil
}

// UpdateRubric updates a rubric
func (r *AssessmentRepository) UpdateRubric(ctx context.Context, rubric *domain.Rubric) error {
	query := `
		UPDATE rubrics 
		SET performance_criteria = $2, performance_levels = $3, scoring_guidelines = $4, status = $5, updated_at = NOW()
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query,
		rubric.ID, rubric.PerformanceCriteria, rubric.PerformanceLevels, rubric.ScoringGuidelines, rubric.Status)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("rubric not found")
	}

	return nil
}

// ==================== Evidence Operations ====================

// CreateEvidence creates a new evidence
func (r *AssessmentRepository) CreateEvidence(ctx context.Context, evidence *domain.Evidence) error {
	query := `
		INSERT INTO evidences (id, student_id, assessment_id, user_id, evidence_type, status, 
		                    evidence_data, teacher_notes, rubric_id, linked_criteria, evaluation_notes)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`

	_, err := r.db.ExecContext(ctx, query,
		evidence.ID, evidence.StudentID, evidence.AssessmentID, evidence.UserID, evidence.EvidenceType,
		evidence.Status, evidence.EvidenceData, evidence.TeacherNotes, evidence.RubricID,
		evidence.LinkedCriteria, evidence.EvaluationNotes)
	return err
}

// GetEvidenceByID retrieves an evidence by ID
func (r *AssessmentRepository) GetEvidenceByID(ctx context.Context, id string) (*domain.Evidence, error) {
	query := `
		SELECT id, student_id, assessment_id, user_id, evidence_type, status, 
		       evidence_data, teacher_notes, rubric_id, linked_criteria, evaluation_notes, created_at, updated_at
		FROM evidences WHERE id = $1
	`

	var evidence domain.Evidence
	var teacherNotes, rubricID, linkedCriteria, evaluationNotes sql.NullString

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&evidence.ID, &evidence.StudentID, &evidence.AssessmentID, &evidence.UserID, &evidence.EvidenceType,
		&evidence.Status, &evidence.EvidenceData, &teacherNotes, &rubricID, &linkedCriteria, &evaluationNotes,
		&evidence.CreatedAt, &evidence.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("evidence not found")
	}
	if err != nil {
		return nil, err
	}

	if teacherNotes.Valid {
		evidence.TeacherNotes = &teacherNotes.String
	}
	if rubricID.Valid {
		evidence.RubricID = &rubricID.String
	}
	if linkedCriteria.Valid {
		evidence.LinkedCriteria = linkedCriteria.String
	}
	if evaluationNotes.Valid {
		evidence.EvaluationNotes = &evaluationNotes.String
	}

	return &evidence, nil
}

// ListEvidences retrieves evidences with optional filters
func (r *AssessmentRepository) ListEvidences(ctx context.Context, studentID, assessmentID *string, evidenceType *domain.EvidenceType, status *domain.EvidenceStatus, limit, offset int) ([]*domain.Evidence, error) {
	query := `
		SELECT id, student_id, assessment_id, user_id, evidence_type, status, 
		       evidence_data, teacher_notes, rubric_id, linked_criteria, evaluation_notes, created_at, updated_at
		FROM evidences
		WHERE 1=1
	`

	args := []interface{}{}
	argIndex := 1

	if studentID != nil {
		query += fmt.Sprintf(" AND student_id = $%d", argIndex)
		args = append(args, *studentID)
		argIndex++
	}

	if assessmentID != nil {
		query += fmt.Sprintf(" AND assessment_id = $%d", argIndex)
		args = append(args, *assessmentID)
		argIndex++
	}

	if evidenceType != nil {
		query += fmt.Sprintf(" AND evidence_type = $%d", argIndex)
		args = append(args, *evidenceType)
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

	var evidences []*domain.Evidence
	for rows.Next() {
		var evidence domain.Evidence
		var teacherNotes, rubricID, linkedCriteria, evaluationNotes sql.NullString

		err := rows.Scan(
			&evidence.ID, &evidence.StudentID, &evidence.AssessmentID, &evidence.UserID, &evidence.EvidenceType,
			&evidence.Status, &evidence.EvidenceData, &teacherNotes, &rubricID, &linkedCriteria, &evaluationNotes,
			&evidence.CreatedAt, &evidence.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		if teacherNotes.Valid {
			evidence.TeacherNotes = &teacherNotes.String
		}
		if rubricID.Valid {
			evidence.RubricID = &rubricID.String
		}
		if linkedCriteria.Valid {
			evidence.LinkedCriteria = linkedCriteria.String
		}
		if evaluationNotes.Valid {
			evidence.EvaluationNotes = &evaluationNotes.String
		}

		evidences = append(evidences, &evidence)
	}

	return evidences, nil
}

// UpdateEvidence updates an evidence
func (r *AssessmentRepository) UpdateEvidence(ctx context.Context, evidence *domain.Evidence) error {
	query := `
		UPDATE evidences 
		SET evidence_data = $2, teacher_notes = $3, rubric_id = $4, linked_criteria = $5, 
		    evaluation_notes = $6, status = $7, updated_at = NOW()
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query,
		evidence.ID, evidence.EvidenceData, evidence.TeacherNotes, evidence.RubricID,
		evidence.LinkedCriteria, evidence.EvaluationNotes, evidence.Status)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("evidence not found")
	}

	return nil
}

// ==================== Evaluation Operations ====================

// CreateEvaluation creates a new evaluation
func (r *AssessmentRepository) CreateEvaluation(ctx context.Context, evaluation *domain.Evaluation) error {
	query := `
		INSERT INTO evaluations (id, student_id, rubric_id, evidence_id, user_id, performance_scores, 
		                    total_score, max_score, performance_level, evaluated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	_, err := r.db.ExecContext(ctx, query,
		evaluation.ID, evaluation.StudentID, evaluation.RubricID, evaluation.EvidenceID, evaluation.UserID,
		evaluation.PerformanceScores, evaluation.TotalScore, evaluation.MaxScore, evaluation.PerformanceLevel, evaluation.EvaluatedAt)
	return err
}

// GetEvaluationByID retrieves an evaluation by ID
func (r *AssessmentRepository) GetEvaluationByID(ctx context.Context, id string) (*domain.Evaluation, error) {
	query := `
		SELECT id, student_id, rubric_id, evidence_id, user_id, performance_scores, 
		       total_score, max_score, performance_level, evaluated_at
		FROM evaluations WHERE id = $1
	`

	var evaluation domain.Evaluation

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&evaluation.ID, &evaluation.StudentID, &evaluation.RubricID, &evaluation.EvidenceID, &evaluation.UserID,
		&evaluation.PerformanceScores, &evaluation.TotalScore, &evaluation.MaxScore, &evaluation.PerformanceLevel, &evaluation.EvaluatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("evaluation not found")
	}
	if err != nil {
		return nil, err
	}

	return &evaluation, nil
}

// ListEvaluations retrieves evaluations with optional filters
func (r *AssessmentRepository) ListEvaluations(ctx context.Context, studentID, rubricID, evidenceID *string, performanceLevel *domain.PerformanceLevel, limit, offset int) ([]*domain.Evaluation, error) {
	query := `
		SELECT id, student_id, rubric_id, evidence_id, user_id, performance_scores, 
		       total_score, max_score, performance_level, evaluated_at
		FROM evaluations
		WHERE 1=1
	`

	args := []interface{}{}
	argIndex := 1

	if studentID != nil {
		query += fmt.Sprintf(" AND student_id = $%d", argIndex)
		args = append(args, *studentID)
		argIndex++
	}

	if rubricID != nil {
		query += fmt.Sprintf(" AND rubric_id = $%d", argIndex)
		args = append(args, *rubricID)
		argIndex++
	}

	if evidenceID != nil {
		query += fmt.Sprintf(" AND evidence_id = $%d", argIndex)
		args = append(args, *evidenceID)
		argIndex++
	}

	if performanceLevel != nil {
		query += fmt.Sprintf(" AND performance_level = $%d", argIndex)
		args = append(args, *performanceLevel)
		argIndex++
	}

	query += " ORDER BY evaluated_at DESC"

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

	var evaluations []*domain.Evaluation
	for rows.Next() {
		var evaluation domain.Evaluation

		err := rows.Scan(
			&evaluation.ID, &evaluation.StudentID, &evaluation.RubricID, &evaluation.EvidenceID, &evaluation.UserID,
			&evaluation.PerformanceScores, &evaluation.TotalScore, &evaluation.MaxScore, &evaluation.PerformanceLevel, &evaluation.EvaluatedAt,
		)
		if err != nil {
			return nil, err
		}

		evaluations = append(evaluations, &evaluation)
	}

	return evaluations, nil
}

// UpdateEvaluation updates an evaluation
func (r *AssessmentRepository) UpdateEvaluation(ctx context.Context, evaluation *domain.Evaluation) error {
	query := `
		UPDATE evaluations 
		SET performance_scores = $2, total_score = $3, max_score = $4, performance_level = $5
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query,
		evaluation.ID, evaluation.PerformanceScores, evaluation.TotalScore, evaluation.MaxScore, evaluation.PerformanceLevel)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("evaluation not found")
	}

	return nil
}

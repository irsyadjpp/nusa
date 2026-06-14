package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/nusa/backend/internal/domain"
)

// LearningPlanningRepository handles database operations for ATP and Modul Ajar entities
type LearningPlanningRepository struct {
	db *sqlx.DB
}

// NewLearningPlanningRepository creates a new learning planning repository
func NewLearningPlanningRepository(db *sqlx.DB) *LearningPlanningRepository {
	return &LearningPlanningRepository{db: db}
}

// ==================== ATP Set Operations ====================

// CreateATPSet creates a new ATP Set
func (r *LearningPlanningRepository) CreateATPSet(ctx context.Context, atpSet *domain.ATPSet) error {
	query := `
		INSERT INTO atp_sets (id, tp_set_id, version_no, status, generation_source, generation_reason, 
		                    generated_by, ai_generation_id, approved_by, approved_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	_, err := r.db.ExecContext(ctx, query,
		atpSet.ID, atpSet.TPSetID, atpSet.VersionNo, atpSet.Status, atpSet.GenerationSource,
		atpSet.GenerationReason, atpSet.GeneratedBy, atpSet.AIGenerationID, atpSet.ApprovedBy, atpSet.ApprovedAt)
	return err
}

// GetATPSetByID retrieves an ATP Set by ID
func (r *LearningPlanningRepository) GetATPSetByID(ctx context.Context, id string) (*domain.ATPSet, error) {
	query := `
		SELECT id, tp_set_id, version_no, status, generation_source, generation_reason, 
		       generated_by, ai_generation_id, approved_by, approved_at, created_at, updated_at
		FROM atp_sets WHERE id = $1
	`

	var atpSet domain.ATPSet
	var generationReason, aiGenerationID, approvedBy sql.NullString
	var approvedAt sql.NullTime

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&atpSet.ID, &atpSet.TPSetID, &atpSet.VersionNo, &atpSet.Status, &atpSet.GenerationSource,
		&generationReason, &atpSet.GeneratedBy, &aiGenerationID, &approvedBy, &approvedAt,
		&atpSet.CreatedAt, &atpSet.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("atp set not found")
	}
	if err != nil {
		return nil, err
	}

	if generationReason.Valid {
		atpSet.GenerationReason = &generationReason.String
	}
	if aiGenerationID.Valid {
		atpSet.AIGenerationID = &aiGenerationID.String
	}
	if approvedBy.Valid {
		atpSet.ApprovedBy = &approvedBy.String
	}
	if approvedAt.Valid {
		atpSet.ApprovedAt = &approvedAt.Time
	}

	return &atpSet, nil
}

// GetATPSetByTPAndVersion retrieves an ATP Set by TP Set ID and version number
func (r *LearningPlanningRepository) GetATPSetByTPAndVersion(ctx context.Context, tpSetID string, versionNo int) (*domain.ATPSet, error) {
	query := `
		SELECT id, tp_set_id, version_no, status, generation_source, generation_reason, 
		       generated_by, ai_generation_id, approved_by, approved_at, created_at, updated_at
		FROM atp_sets WHERE tp_set_id = $1 AND version_no = $2
	`

	var atpSet domain.ATPSet
	var generationReason, aiGenerationID, approvedBy sql.NullString
	var approvedAt sql.NullTime

	err := r.db.QueryRowContext(ctx, query, tpSetID, versionNo).Scan(
		&atpSet.ID, &atpSet.TPSetID, &atpSet.VersionNo, &atpSet.Status, &atpSet.GenerationSource,
		&generationReason, &atpSet.GeneratedBy, &aiGenerationID, &approvedBy, &approvedAt,
		&atpSet.CreatedAt, &atpSet.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("atp set not found")
	}
	if err != nil {
		return nil, err
	}

	if generationReason.Valid {
		atpSet.GenerationReason = &generationReason.String
	}
	if aiGenerationID.Valid {
		atpSet.AIGenerationID = &aiGenerationID.String
	}
	if approvedBy.Valid {
		atpSet.ApprovedBy = &approvedBy.String
	}
	if approvedAt.Valid {
		atpSet.ApprovedAt = &approvedAt.Time
	}

	return &atpSet, nil
}

// ListATPSets retrieves ATP Sets with optional filters
func (r *LearningPlanningRepository) ListATPSets(ctx context.Context, tpSetID *string, status *domain.WorkflowStatus, limit, offset int) ([]*domain.ATPSet, error) {
	query := `
		SELECT id, tp_set_id, version_no, status, generation_source, generation_reason, 
		       generated_by, ai_generation_id, approved_by, approved_at, created_at, updated_at
		FROM atp_sets
		WHERE 1=1
	`

	args := []interface{}{}
	argIndex := 1

	if tpSetID != nil {
		query += fmt.Sprintf(" AND tp_set_id = $%d", argIndex)
		args = append(args, *tpSetID)
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

	var atpSets []*domain.ATPSet
	for rows.Next() {
		var atpSet domain.ATPSet
		var generationReason, aiGenerationID, approvedBy sql.NullString
		var approvedAt sql.NullTime

		err := rows.Scan(
			&atpSet.ID, &atpSet.TPSetID, &atpSet.VersionNo, &atpSet.Status, &atpSet.GenerationSource,
			&generationReason, &atpSet.GeneratedBy, &aiGenerationID, &approvedBy, &approvedAt,
			&atpSet.CreatedAt, &atpSet.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		if generationReason.Valid {
			atpSet.GenerationReason = &generationReason.String
		}
		if aiGenerationID.Valid {
			atpSet.AIGenerationID = &aiGenerationID.String
		}
		if approvedBy.Valid {
			atpSet.ApprovedBy = &approvedBy.String
		}
		if approvedAt.Valid {
			atpSet.ApprovedAt = &approvedAt.Time
		}

		atpSets = append(atpSets, &atpSet)
	}

	return atpSets, nil
}

// UpdateATPSet updates an ATP Set
func (r *LearningPlanningRepository) UpdateATPSet(ctx context.Context, atpSet *domain.ATPSet) error {
	query := `
		UPDATE atp_sets 
		SET status = $2, generation_reason = $3, approved_by = $4, approved_at = $5, updated_at = NOW()
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query,
		atpSet.ID, atpSet.Status, atpSet.GenerationReason, atpSet.ApprovedBy, atpSet.ApprovedAt)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("atp set not found")
	}

	return nil
}

// ==================== ATP Item Operations ====================

// CreateATP creates a new ATP
func (r *LearningPlanningRepository) CreateATP(ctx context.Context, atp *domain.ATP) error {
	query := `
		INSERT INTO atp (id, atp_set_id, tp_id, user_id, status, academic_calendar, 
		              class_schedule, weekly_sequence, assessment_schedule)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	_, err := r.db.ExecContext(ctx, query,
		atp.ID, atp.ATPSetID, atp.TPID, atp.UserID, atp.Status,
		atp.AcademicCalendar, atp.ClassSchedule, atp.WeeklySequence, atp.AssessmentSchedule)
	return err
}

// GetATPByID retrieves an ATP by ID
func (r *LearningPlanningRepository) GetATPByID(ctx context.Context, id string) (*domain.ATP, error) {
	query := `
		SELECT id, atp_set_id, tp_id, user_id, status, academic_calendar, 
		       class_schedule, weekly_sequence, assessment_schedule, created_at, updated_at
		FROM atp WHERE id = $1
	`

	var atp domain.ATP
	var assessmentSchedule sql.NullString

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&atp.ID, &atp.ATPSetID, &atp.TPID, &atp.UserID, &atp.Status,
		&atp.AcademicCalendar, &atp.ClassSchedule, &atp.WeeklySequence, &assessmentSchedule,
		&atp.CreatedAt, &atp.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("atp not found")
	}
	if err != nil {
		return nil, err
	}

	if assessmentSchedule.Valid {
		atp.AssessmentSchedule = assessmentSchedule.String
	}

	return &atp, nil
}

// ListATPsBySet retrieves ATPs by ATP Set ID
func (r *LearningPlanningRepository) ListATPsBySet(ctx context.Context, atpSetID string) ([]*domain.ATP, error) {
	query := `
		SELECT id, atp_set_id, tp_id, user_id, status, academic_calendar, 
		       class_schedule, weekly_sequence, assessment_schedule, created_at, updated_at
		FROM atp WHERE atp_set_id = $1 ORDER BY id ASC
	`

	rows, err := r.db.QueryContext(ctx, query, atpSetID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var atps []*domain.ATP
	for rows.Next() {
		var atp domain.ATP
		var assessmentSchedule sql.NullString

		err := rows.Scan(
			&atp.ID, &atp.ATPSetID, &atp.TPID, &atp.UserID, &atp.Status,
			&atp.AcademicCalendar, &atp.ClassSchedule, &atp.WeeklySequence, &assessmentSchedule,
			&atp.CreatedAt, &atp.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		if assessmentSchedule.Valid {
			atp.AssessmentSchedule = assessmentSchedule.String
		}

		atps = append(atps, &atp)
	}

	return atps, nil
}

// UpdateATP updates an ATP
func (r *LearningPlanningRepository) UpdateATP(ctx context.Context, atp *domain.ATP) error {
	query := `
		UPDATE atp 
		SET academic_calendar = $2, class_schedule = $3, weekly_sequence = $4, 
		    assessment_schedule = $5, status = $6, updated_at = NOW()
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query,
		atp.ID, atp.AcademicCalendar, atp.ClassSchedule, atp.WeeklySequence, atp.AssessmentSchedule, atp.Status)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("atp not found")
	}

	return nil
}

// DeleteATP deletes an ATP
func (r *LearningPlanningRepository) DeleteATP(ctx context.Context, id string) error {
	query := `DELETE FROM atp WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("atp not found")
	}

	return nil
}

// DeleteATPSet deletes an ATP Set
func (r *LearningPlanningRepository) DeleteATPSet(ctx context.Context, id string) error {
	query := `DELETE FROM atp_sets WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("atp set not found")
	}

	return nil
}

// ==================== Modul Ajar Set Operations ====================

// CreateModulAjarSet creates a new Modul Ajar Set
func (r *LearningPlanningRepository) CreateModulAjarSet(ctx context.Context, modulAjarSet *domain.ModulAjarSet) error {
	query := `
		INSERT INTO modul_ajar_sets (id, atp_set_id, version_no, status, generation_source, generation_reason, 
		                            generated_by, ai_generation_id, approved_by, approved_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	_, err := r.db.ExecContext(ctx, query,
		modulAjarSet.ID, modulAjarSet.ATPSetID, modulAjarSet.VersionNo, modulAjarSet.Status,
		modulAjarSet.GenerationSource, modulAjarSet.GenerationReason, modulAjarSet.GeneratedBy,
		modulAjarSet.AIGenerationID, modulAjarSet.ApprovedBy, modulAjarSet.ApprovedAt)
	return err
}

// GetModulAjarSetByID retrieves a Modul Ajar Set by ID
func (r *LearningPlanningRepository) GetModulAjarSetByID(ctx context.Context, id string) (*domain.ModulAjarSet, error) {
	query := `
		SELECT id, atp_set_id, version_no, status, generation_source, generation_reason, 
		       generated_by, ai_generation_id, approved_by, approved_at, created_at, updated_at
		FROM modul_ajar_sets WHERE id = $1
	`

	var modulAjarSet domain.ModulAjarSet
	var generationReason, aiGenerationID, approvedBy sql.NullString
	var approvedAt sql.NullTime

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&modulAjarSet.ID, &modulAjarSet.ATPSetID, &modulAjarSet.VersionNo, &modulAjarSet.Status,
		&modulAjarSet.GenerationSource, &generationReason, &modulAjarSet.GeneratedBy, &modulAjarSet.AIGenerationID,
		&approvedBy, &approvedAt, &modulAjarSet.CreatedAt, &modulAjarSet.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("modul ajar set not found")
	}
	if err != nil {
		return nil, err
	}

	if generationReason.Valid {
		modulAjarSet.GenerationReason = &generationReason.String
	}
	if aiGenerationID.Valid {
		modulAjarSet.AIGenerationID = &aiGenerationID.String
	}
	if approvedBy.Valid {
		modulAjarSet.ApprovedBy = &approvedBy.String
	}
	if approvedAt.Valid {
		modulAjarSet.ApprovedAt = &approvedAt.Time
	}

	return &modulAjarSet, nil
}

// GetModulAjarSetByATPAndVersion retrieves a Modul Ajar Set by ATP Set ID and version number
func (r *LearningPlanningRepository) GetModulAjarSetByATPAndVersion(ctx context.Context, atpSetID string, versionNo int) (*domain.ModulAjarSet, error) {
	query := `
		SELECT id, atp_set_id, version_no, status, generation_source, generation_reason, 
		       generated_by, ai_generation_id, approved_by, approved_at, created_at, updated_at
		FROM modul_ajar_sets WHERE atp_set_id = $1 AND version_no = $2
	`

	var modulAjarSet domain.ModulAjarSet
	var generationReason, aiGenerationID, approvedBy sql.NullString
	var approvedAt sql.NullTime

	err := r.db.QueryRowContext(ctx, query, atpSetID, versionNo).Scan(
		&modulAjarSet.ID, &modulAjarSet.ATPSetID, &modulAjarSet.VersionNo, &modulAjarSet.Status,
		&modulAjarSet.GenerationSource, &generationReason, &modulAjarSet.GeneratedBy, &modulAjarSet.AIGenerationID,
		&approvedBy, &approvedAt, &modulAjarSet.CreatedAt, &modulAjarSet.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("modul ajar set not found")
	}
	if err != nil {
		return nil, err
	}

	if generationReason.Valid {
		modulAjarSet.GenerationReason = &generationReason.String
	}
	if aiGenerationID.Valid {
		modulAjarSet.AIGenerationID = &aiGenerationID.String
	}
	if approvedBy.Valid {
		modulAjarSet.ApprovedBy = &approvedBy.String
	}
	if approvedAt.Valid {
		modulAjarSet.ApprovedAt = &approvedAt.Time
	}

	return &modulAjarSet, nil
}

// ListModulAjarSets retrieves Modul Ajar Sets with optional filters
func (r *LearningPlanningRepository) ListModulAjarSets(ctx context.Context, atpSetID *string, status *domain.WorkflowStatus, limit, offset int) ([]*domain.ModulAjarSet, error) {
	query := `
		SELECT id, atp_set_id, version_no, status, generation_source, generation_reason, 
		       generated_by, ai_generation_id, approved_by, approved_at, created_at, updated_at
		FROM modul_ajar_sets
		WHERE 1=1
	`

	args := []interface{}{}
	argIndex := 1

	if atpSetID != nil {
		query += fmt.Sprintf(" AND atp_set_id = $%d", argIndex)
		args = append(args, *atpSetID)
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

	var modulAjarSets []*domain.ModulAjarSet
	for rows.Next() {
		var modulAjarSet domain.ModulAjarSet
		var generationReason, aiGenerationID, approvedBy sql.NullString
		var approvedAt sql.NullTime

		err := rows.Scan(
			&modulAjarSet.ID, &modulAjarSet.ATPSetID, &modulAjarSet.VersionNo, &modulAjarSet.Status,
			&modulAjarSet.GenerationSource, &generationReason, &modulAjarSet.GeneratedBy, &modulAjarSet.AIGenerationID,
			&approvedBy, &approvedAt, &modulAjarSet.CreatedAt, &modulAjarSet.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		if generationReason.Valid {
			modulAjarSet.GenerationReason = &generationReason.String
		}
		if aiGenerationID.Valid {
			modulAjarSet.AIGenerationID = &aiGenerationID.String
		}
		if approvedBy.Valid {
			modulAjarSet.ApprovedBy = &approvedBy.String
		}
		if approvedAt.Valid {
			modulAjarSet.ApprovedAt = &approvedAt.Time
		}

		modulAjarSets = append(modulAjarSets, &modulAjarSet)
	}

	return modulAjarSets, nil
}

// UpdateModulAjarSet updates a Modul Ajar Set
func (r *LearningPlanningRepository) UpdateModulAjarSet(ctx context.Context, modulAjarSet *domain.ModulAjarSet) error {
	query := `
		UPDATE modul_ajar_sets 
		SET status = $2, generation_reason = $3, approved_by = $4, approved_at = $5, updated_at = NOW()
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query,
		modulAjarSet.ID, modulAjarSet.Status, modulAjarSet.GenerationReason, modulAjarSet.ApprovedBy, modulAjarSet.ApprovedAt)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("modul ajar set not found")
	}

	return nil
}

// DeleteModulAjarSet deletes a Modul Ajar Set
func (r *LearningPlanningRepository) DeleteModulAjarSet(ctx context.Context, id string) error {
	query := `DELETE FROM modul_ajar_sets WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("modul ajar set not found")
	}

	return nil
}

// ==================== Modul Ajar Item Operations ====================

// CreateModulAjar creates a new Modul Ajar
func (r *LearningPlanningRepository) CreateModulAjar(ctx context.Context, modulAjar *domain.ModulAjar) error {
	query := `
		INSERT INTO modul_ajar (id, modul_ajar_set_id, atp_id, week, topic, resources, 
		                 class_characteristics, learning_activities, resource_requirements, assessment_methods, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`

	_, err := r.db.ExecContext(ctx, query,
		modulAjar.ID, modulAjar.ModulAjarSetID, modulAjar.ATPID, modulAjar.Week, modulAjar.Topic,
		modulAjar.Resources, modulAjar.ClassCharacteristics, modulAjar.LearningActivities,
		modulAjar.ResourceRequirements, modulAjar.AssessmentMethods, modulAjar.Status)
	return err
}

// GetModulAjarByID retrieves a Modul Ajar by ID
func (r *LearningPlanningRepository) GetModulAjarByID(ctx context.Context, id string) (*domain.ModulAjar, error) {
	query := `
		SELECT id, modul_ajar_set_id, atp_id, week, topic, resources, 
		       class_characteristics, learning_activities, resource_requirements, assessment_methods, 
		       status, created_at, updated_at
		FROM modul_ajar WHERE id = $1
	`

	var modulAjar domain.ModulAjar
	var resources, classCharacteristics, resourceRequirements, assessmentMethods sql.NullString

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&modulAjar.ID, &modulAjar.ModulAjarSetID, &modulAjar.ATPID, &modulAjar.Week, &modulAjar.Topic,
		&resources, &classCharacteristics, &modulAjar.LearningActivities, &resourceRequirements, &assessmentMethods,
		&modulAjar.Status, &modulAjar.CreatedAt, &modulAjar.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("modul ajar not found")
	}
	if err != nil {
		return nil, err
	}

	if resources.Valid {
		modulAjar.Resources = resources.String
	}
	if classCharacteristics.Valid {
		modulAjar.ClassCharacteristics = classCharacteristics.String
	}
	if resourceRequirements.Valid {
		modulAjar.ResourceRequirements = resourceRequirements.String
	}
	if assessmentMethods.Valid {
		modulAjar.AssessmentMethods = assessmentMethods.String
	}

	return &modulAjar, nil
}

// ListModulAjarsBySet retrieves Modul Ajars by Modul Ajar Set ID
func (r *LearningPlanningRepository) ListModulAjarsBySet(ctx context.Context, modulAjarSetID string) ([]*domain.ModulAjar, error) {
	query := `
		SELECT id, modul_ajar_set_id, atp_id, week, topic, resources, 
		       class_characteristics, learning_activities, resource_requirements, assessment_methods, 
		       status, created_at, updated_at
		FROM modul_ajar WHERE modul_ajar_set_id = $1 ORDER BY week ASC
	`

	rows, err := r.db.QueryContext(ctx, query, modulAjarSetID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var modulAjars []*domain.ModulAjar
	for rows.Next() {
		var modulAjar domain.ModulAjar
		var resources, classCharacteristics, resourceRequirements, assessmentMethods sql.NullString

		err := rows.Scan(
			&modulAjar.ID, &modulAjar.ModulAjarSetID, &modulAjar.ATPID, &modulAjar.Week, &modulAjar.Topic,
			&resources, &classCharacteristics, &modulAjar.LearningActivities, &resourceRequirements, &assessmentMethods,
			&modulAjar.Status, &modulAjar.CreatedAt, &modulAjar.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		if resources.Valid {
			modulAjar.Resources = resources.String
		}
		if classCharacteristics.Valid {
			modulAjar.ClassCharacteristics = classCharacteristics.String
		}
		if resourceRequirements.Valid {
			modulAjar.ResourceRequirements = resourceRequirements.String
		}
		if assessmentMethods.Valid {
			modulAjar.AssessmentMethods = assessmentMethods.String
		}

		modulAjars = append(modulAjars, &modulAjar)
	}

	return modulAjars, nil
}

// UpdateModulAjar updates a Modul Ajar
func (r *LearningPlanningRepository) UpdateModulAjar(ctx context.Context, modulAjar *domain.ModulAjar) error {
	query := `
		UPDATE modul_ajar 
		SET topic = $2, resources = $3, class_characteristics = $4, learning_activities = $5, 
		    resource_requirements = $6, assessment_methods = $7, status = $8, updated_at = NOW()
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query,
		modulAjar.ID, modulAjar.Topic, modulAjar.Resources, modulAjar.ClassCharacteristics,
		modulAjar.LearningActivities, modulAjar.ResourceRequirements, modulAjar.AssessmentMethods, modulAjar.Status)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("modul ajar not found")
	}

	return nil
}

// DeleteModulAjar deletes a Modul Ajar
func (r *LearningPlanningRepository) DeleteModulAjar(ctx context.Context, id string) error {
	query := `DELETE FROM modul_ajar WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("modul ajar not found")
	}

	return nil
}

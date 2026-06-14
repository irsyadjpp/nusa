package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/nusa/backend/internal/domain"
)

// CurriculumRepository handles database operations for curriculum entities
type CurriculumRepository struct {
	db *sqlx.DB
}

// NewCurriculumRepository creates a new curriculum repository
func NewCurriculumRepository(db *sqlx.DB) *CurriculumRepository {
	return &CurriculumRepository{db: db}
}

// ==================== CurriculumSubject Operations ====================

// CreateCurriculumSubject creates a new curriculum subject
func (r *CurriculumRepository) CreateCurriculumSubject(ctx context.Context, subject *domain.CurriculumSubject) error {
	query := `
		INSERT INTO curriculum_subjects (id, code, name, description, is_active)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.db.ExecContext(ctx, query,
		subject.ID, subject.Code, subject.Name, subject.Description, subject.IsActive)
	return err
}

// GetCurriculumSubjectByID retrieves a curriculum subject by ID
func (r *CurriculumRepository) GetCurriculumSubjectByID(ctx context.Context, id string) (*domain.CurriculumSubject, error) {
	query := `
		SELECT id, code, name, description, is_active, created_at, updated_at
		FROM curriculum_subjects WHERE id = $1
	`

	var subject domain.CurriculumSubject
	var description sql.NullString

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&subject.ID, &subject.Code, &subject.Name, &description,
		&subject.IsActive, &subject.CreatedAt, &subject.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("curriculum subject not found")
	}
	if err != nil {
		return nil, err
	}

	if description.Valid {
		subject.Description = &description.String
	}

	return &subject, nil
}

// GetCurriculumSubjectByCode retrieves a curriculum subject by code
func (r *CurriculumRepository) GetCurriculumSubjectByCode(ctx context.Context, code string) (*domain.CurriculumSubject, error) {
	query := `
		SELECT id, code, name, description, is_active, created_at, updated_at
		FROM curriculum_subjects WHERE code = $1
	`

	var subject domain.CurriculumSubject
	var description sql.NullString

	err := r.db.QueryRowContext(ctx, query, code).Scan(
		&subject.ID, &subject.Code, &subject.Name, &description,
		&subject.IsActive, &subject.CreatedAt, &subject.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("curriculum subject not found")
	}
	if err != nil {
		return nil, err
	}

	if description.Valid {
		subject.Description = &description.String
	}

	return &subject, nil
}

// ListCurriculumSubjects retrieves curriculum subjects with optional filters
func (r *CurriculumRepository) ListCurriculumSubjects(ctx context.Context, isActive *bool, limit, offset int) ([]*domain.CurriculumSubject, error) {
	query := `
		SELECT id, code, name, description, is_active, created_at, updated_at
		FROM curriculum_subjects
		WHERE 1=1
	`

	args := []interface{}{}
	argIndex := 1

	if isActive != nil {
		query += fmt.Sprintf(" AND is_active = $%d", argIndex)
		args = append(args, *isActive)
		argIndex++
	}

	query += " ORDER BY name ASC"

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

	var subjects []*domain.CurriculumSubject
	for rows.Next() {
		var subject domain.CurriculumSubject
		var description sql.NullString

		err := rows.Scan(
			&subject.ID, &subject.Code, &subject.Name, &description,
			&subject.IsActive, &subject.CreatedAt, &subject.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		if description.Valid {
			subject.Description = &description.String
		}

		subjects = append(subjects, &subject)
	}

	return subjects, nil
}

// UpdateCurriculumSubject updates a curriculum subject
func (r *CurriculumRepository) UpdateCurriculumSubject(ctx context.Context, subject *domain.CurriculumSubject) error {
	query := `
		UPDATE curriculum_subjects
		SET name = $2, description = $3, is_active = $4, updated_at = NOW()
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query,
		subject.ID, subject.Name, subject.Description, subject.IsActive)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("curriculum subject not found")
	}

	return nil
}

// DeleteCurriculumSubject deletes a curriculum subject
func (r *CurriculumRepository) DeleteCurriculumSubject(ctx context.Context, id string) error {
	query := `DELETE FROM curriculum_subjects WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("curriculum subject not found")
	}

	return nil
}

// ==================== CurriculumPhase Operations ====================

// CreateCurriculumPhase creates a new curriculum phase
func (r *CurriculumRepository) CreateCurriculumPhase(ctx context.Context, phase *domain.CurriculumPhase) error {
	query := `
		INSERT INTO curriculum_phases (id, code, name, description, grade_level_start, grade_level_end, is_active)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.db.ExecContext(ctx, query,
		phase.ID, phase.Code, phase.Name, phase.Description,
		phase.GradeLevelStart, phase.GradeLevelEnd, phase.IsActive)
	return err
}

// GetCurriculumPhaseByID retrieves a curriculum phase by ID
func (r *CurriculumRepository) GetCurriculumPhaseByID(ctx context.Context, id string) (*domain.CurriculumPhase, error) {
	query := `
		SELECT id, code, name, description, grade_level_start, grade_level_end, is_active, created_at, updated_at
		FROM curriculum_phases WHERE id = $1
	`

	var phase domain.CurriculumPhase
	var description sql.NullString
	var gradeLevelStart, gradeLevelEnd sql.NullInt32

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&phase.ID, &phase.Code, &phase.Name, &description,
		&gradeLevelStart, &gradeLevelEnd, &phase.IsActive, &phase.CreatedAt, &phase.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("curriculum phase not found")
	}
	if err != nil {
		return nil, err
	}

	if description.Valid {
		phase.Description = &description.String
	}
	if gradeLevelStart.Valid {
		start := int(gradeLevelStart.Int32)
		phase.GradeLevelStart = &start
	}
	if gradeLevelEnd.Valid {
		end := int(gradeLevelEnd.Int32)
		phase.GradeLevelEnd = &end
	}

	return &phase, nil
}

// ListCurriculumPhases retrieves curriculum phases with optional filters
func (r *CurriculumRepository) ListCurriculumPhases(ctx context.Context, isActive *bool, limit, offset int) ([]*domain.CurriculumPhase, error) {
	query := `
		SELECT id, code, name, description, grade_level_start, grade_level_end, is_active, created_at, updated_at
		FROM curriculum_phases
		WHERE 1=1
	`

	args := []interface{}{}
	argIndex := 1

	if isActive != nil {
		query += fmt.Sprintf(" AND is_active = $%d", argIndex)
		args = append(args, *isActive)
		argIndex++
	}

	query += " ORDER BY grade_level_start ASC"

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

	var phases []*domain.CurriculumPhase
	for rows.Next() {
		var phase domain.CurriculumPhase
		var description sql.NullString
		var gradeLevelStart, gradeLevelEnd sql.NullInt32

		err := rows.Scan(
			&phase.ID, &phase.Code, &phase.Name, &description,
			&gradeLevelStart, &gradeLevelEnd, &phase.IsActive, &phase.CreatedAt, &phase.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		if description.Valid {
			phase.Description = &description.String
		}
		if gradeLevelStart.Valid {
			start := int(gradeLevelStart.Int32)
			phase.GradeLevelStart = &start
		}
		if gradeLevelEnd.Valid {
			end := int(gradeLevelEnd.Int32)
			phase.GradeLevelEnd = &end
		}

		phases = append(phases, &phase)
	}

	return phases, nil
}

// ==================== CurriculumElement Operations ====================

// CreateCurriculumElement creates a new curriculum element
func (r *CurriculumRepository) CreateCurriculumElement(ctx context.Context, element *domain.CurriculumElement) error {
	query := `
		INSERT INTO curriculum_elements (id, subject_id, phase_id, code, name, description, is_active)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.db.ExecContext(ctx, query,
		element.ID, element.SubjectID, element.PhaseID, element.Code,
		element.Name, element.Description, element.IsActive)
	return err
}

// GetCurriculumElementByID retrieves a curriculum element by ID
func (r *CurriculumRepository) GetCurriculumElementByID(ctx context.Context, id string) (*domain.CurriculumElement, error) {
	query := `
		SELECT id, subject_id, phase_id, code, name, description, is_active, created_at, updated_at
		FROM curriculum_elements WHERE id = $1
	`

	var element domain.CurriculumElement
	var description sql.NullString

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&element.ID, &element.SubjectID, &element.PhaseID, &element.Code,
		&element.Name, &description, &element.IsActive, &element.CreatedAt, &element.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("curriculum element not found")
	}
	if err != nil {
		return nil, err
	}

	if description.Valid {
		element.Description = &description.String
	}

	return &element, nil
}

// ListCurriculumElements retrieves curriculum elements with optional filters
func (r *CurriculumRepository) ListCurriculumElements(ctx context.Context, subjectID, phaseID *string, isActive *bool, limit, offset int) ([]*domain.CurriculumElement, error) {
	query := `
		SELECT id, subject_id, phase_id, code, name, description, is_active, created_at, updated_at
		FROM curriculum_elements
		WHERE 1=1
	`

	args := []interface{}{}
	argIndex := 1

	if subjectID != nil {
		query += fmt.Sprintf(" AND subject_id = $%d", argIndex)
		args = append(args, *subjectID)
		argIndex++
	}

	if phaseID != nil {
		query += fmt.Sprintf(" AND phase_id = $%d", argIndex)
		args = append(args, *phaseID)
		argIndex++
	}

	if isActive != nil {
		query += fmt.Sprintf(" AND is_active = $%d", argIndex)
		args = append(args, *isActive)
		argIndex++
	}

	query += " ORDER BY code ASC"

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

	var elements []*domain.CurriculumElement
	for rows.Next() {
		var element domain.CurriculumElement
		var description sql.NullString

		err := rows.Scan(
			&element.ID, &element.SubjectID, &element.PhaseID, &element.Code,
			&element.Name, &description, &element.IsActive, &element.CreatedAt, &element.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		if description.Valid {
			element.Description = &description.String
		}

		elements = append(elements, &element)
	}

	return elements, nil
}

// ==================== CurriculumSubelement Operations ====================

// CreateCurriculumSubelement creates a new curriculum subelement
func (r *CurriculumRepository) CreateCurriculumSubelement(ctx context.Context, subelement *domain.CurriculumSubelement) error {
	query := `
		INSERT INTO curriculum_subelements (id, element_id, code, name, description, is_active)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.ExecContext(ctx, query,
		subelement.ID, subelement.ElementID, subelement.Code,
		subelement.Name, subelement.Description, subelement.IsActive)
	return err
}

// GetCurriculumSubelementByID retrieves a curriculum subelement by ID
func (r *CurriculumRepository) GetCurriculumSubelementByID(ctx context.Context, id string) (*domain.CurriculumSubelement, error) {
	query := `
		SELECT id, element_id, code, name, description, is_active, created_at, updated_at
		FROM curriculum_subelements WHERE id = $1
	`

	var subelement domain.CurriculumSubelement
	var description sql.NullString

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&subelement.ID, &subelement.ElementID, &subelement.Code,
		&subelement.Name, &description, &subelement.IsActive, &subelement.CreatedAt, &subelement.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("curriculum subelement not found")
	}
	if err != nil {
		return nil, err
	}

	if description.Valid {
		subelement.Description = &description.String
	}

	return &subelement, nil
}

// ListCurriculumSubelements retrieves curriculum subelements with optional filters
func (r *CurriculumRepository) ListCurriculumSubelements(ctx context.Context, elementID *string, isActive *bool, limit, offset int) ([]*domain.CurriculumSubelement, error) {
	query := `
		SELECT id, element_id, code, name, description, is_active, created_at, updated_at
		FROM curriculum_subelements
		WHERE 1=1
	`

	args := []interface{}{}
	argIndex := 1

	if elementID != nil {
		query += fmt.Sprintf(" AND element_id = $%d", argIndex)
		args = append(args, *elementID)
		argIndex++
	}

	if isActive != nil {
		query += fmt.Sprintf(" AND is_active = $%d", argIndex)
		args = append(args, *isActive)
		argIndex++
	}

	query += " ORDER BY code ASC"

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

	var subelements []*domain.CurriculumSubelement
	for rows.Next() {
		var subelement domain.CurriculumSubelement
		var description sql.NullString

		err := rows.Scan(
			&subelement.ID, &subelement.ElementID, &subelement.Code,
			&subelement.Name, &description, &subelement.IsActive, &subelement.CreatedAt, &subelement.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		if description.Valid {
			subelement.Description = &description.String
		}

		subelements = append(subelements, &subelement)
	}

	return subelements, nil
}

// UpdateCurriculumSubelement updates a curriculum subelement
func (r *CurriculumRepository) UpdateCurriculumSubelement(ctx context.Context, subelement *domain.CurriculumSubelement) error {
	query := `
		UPDATE curriculum_subelements
		SET name = $2, description = $3, is_active = $4, updated_at = NOW()
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query,
		subelement.ID, subelement.Name, subelement.Description, subelement.IsActive)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("curriculum subelement not found")
	}

	return nil
}

// DeleteCurriculumSubelement deletes a curriculum subelement
func (r *CurriculumRepository) DeleteCurriculumSubelement(ctx context.Context, id string) error {
	query := `DELETE FROM curriculum_subelements WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("curriculum subelement not found")
	}

	return nil
}

// ==================== CP Operations ====================

// CreateCP creates a new CP
func (r *CurriculumRepository) CreateCP(ctx context.Context, cp *domain.CP) error {
	query := `
		INSERT INTO cp (id, subject_id, phase_id, element_id, subelement_id, code, description, 
		             competency_code, learning_objectives, competency_standards, time_allocation_hours, 
		             hours_per_week, version, is_active, imported_at, imported_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
	`

	_, err := r.db.ExecContext(ctx, query,
		cp.ID, cp.SubjectID, cp.PhaseID, cp.ElementID, cp.SubelementID, cp.Code, cp.Description,
		cp.CompetencyCode, cp.LearningObjectives, cp.CompetencyStandards, cp.TimeAllocationHours,
		cp.HoursPerWeek, cp.Version, cp.IsActive, cp.ImportedAt, cp.ImportedBy)
	return err
}

// GetCPByID retrieves a CP by ID
func (r *CurriculumRepository) GetCPByID(ctx context.Context, id string) (*domain.CP, error) {
	query := `
		SELECT id, subject_id, phase_id, element_id, subelement_id, code, description, 
		       competency_code, learning_objectives, competency_standards, time_allocation_hours, 
		       hours_per_week, version, is_active, imported_at, imported_by
		FROM cp WHERE id = $1
	`

	var cp domain.CP
	var competencyCode, importedBy sql.NullString

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&cp.ID, &cp.SubjectID, &cp.PhaseID, &cp.ElementID, &cp.SubelementID, &cp.Code, &cp.Description,
		&competencyCode, &cp.LearningObjectives, &cp.CompetencyStandards, &cp.TimeAllocationHours,
		&cp.HoursPerWeek, &cp.Version, &cp.IsActive, &cp.ImportedAt, &importedBy,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("cp not found")
	}
	if err != nil {
		return nil, err
	}

	if competencyCode.Valid {
		cp.CompetencyCode = &competencyCode.String
	}
	if importedBy.Valid {
		cp.ImportedBy = &importedBy.String
	}

	return &cp, nil
}

// GetCPByHierarchy retrieves a CP by its hierarchy identifiers
func (r *CurriculumRepository) GetCPByHierarchy(ctx context.Context, subjectID, phaseID, elementID, subelementID, code string) (*domain.CP, error) {
	query := `
		SELECT id, subject_id, phase_id, element_id, subelement_id, code, description, 
		       competency_code, learning_objectives, competency_standards, time_allocation_hours, 
		       hours_per_week, version, is_active, imported_at, imported_by
		FROM cp 
		WHERE subject_id = $1 AND phase_id = $2 AND element_id = $3 AND subelement_id = $4 AND code = $5
	`

	var cp domain.CP
	var competencyCode, importedBy sql.NullString

	err := r.db.QueryRowContext(ctx, query, subjectID, phaseID, elementID, subelementID, code).Scan(
		&cp.ID, &cp.SubjectID, &cp.PhaseID, &cp.ElementID, &cp.SubelementID, &cp.Code, &cp.Description,
		&competencyCode, &cp.LearningObjectives, &cp.CompetencyStandards, &cp.TimeAllocationHours,
		&cp.HoursPerWeek, &cp.Version, &cp.IsActive, &cp.ImportedAt, &importedBy,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("cp not found")
	}
	if err != nil {
		return nil, err
	}

	if competencyCode.Valid {
		cp.CompetencyCode = &competencyCode.String
	}
	if importedBy.Valid {
		cp.ImportedBy = &importedBy.String
	}

	return &cp, nil
}

// ListCPs retrieves CPs with optional filters
func (r *CurriculumRepository) ListCPs(ctx context.Context, subjectID, phaseID, elementID *string, version *string, isActive *bool, limit, offset int) ([]*domain.CP, error) {
	query := `
		SELECT id, subject_id, phase_id, element_id, subelement_id, code, description, 
		       competency_code, learning_objectives, competency_standards, time_allocation_hours, 
		       hours_per_week, version, is_active, imported_at, imported_by
		FROM cp
		WHERE 1=1
	`

	args := []interface{}{}
	argIndex := 1

	if subjectID != nil {
		query += fmt.Sprintf(" AND subject_id = $%d", argIndex)
		args = append(args, *subjectID)
		argIndex++
	}

	if phaseID != nil {
		query += fmt.Sprintf(" AND phase_id = $%d", argIndex)
		args = append(args, *phaseID)
		argIndex++
	}

	if elementID != nil {
		query += fmt.Sprintf(" AND element_id = $%d", argIndex)
		args = append(args, *elementID)
		argIndex++
	}

	if version != nil {
		query += fmt.Sprintf(" AND version = $%d", argIndex)
		args = append(args, *version)
		argIndex++
	}

	if isActive != nil {
		query += fmt.Sprintf(" AND is_active = $%d", argIndex)
		args = append(args, *isActive)
		argIndex++
	}

	query += " ORDER BY code ASC"

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

	var cps []*domain.CP
	for rows.Next() {
		var cp domain.CP
		var competencyCode, importedBy sql.NullString

		err := rows.Scan(
			&cp.ID, &cp.SubjectID, &cp.PhaseID, &cp.ElementID, &cp.SubelementID, &cp.Code, &cp.Description,
			&competencyCode, &cp.LearningObjectives, &cp.CompetencyStandards, &cp.TimeAllocationHours,
			&cp.HoursPerWeek, &cp.Version, &cp.IsActive, &cp.ImportedAt, &importedBy,
		)
		if err != nil {
			return nil, err
		}

		if competencyCode.Valid {
			cp.CompetencyCode = &competencyCode.String
		}
		if importedBy.Valid {
			cp.ImportedBy = &importedBy.String
		}

		cps = append(cps, &cp)
	}

	return cps, nil
}

// UpdateCP updates a CP
func (r *CurriculumRepository) UpdateCP(ctx context.Context, cp *domain.CP) error {
	query := `
		UPDATE cp 
		SET description = $2, competency_code = $3, learning_objectives = $4, competency_standards = $5,
		    time_allocation_hours = $6, hours_per_week = $7, version = $8, is_active = $9, updated_at = NOW()
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query,
		cp.ID, cp.Description, cp.CompetencyCode, cp.LearningObjectives, cp.CompetencyStandards,
		cp.TimeAllocationHours, cp.HoursPerWeek, cp.Version, cp.IsActive)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("cp not found")
	}

	return nil
}

// DeleteCP deletes a CP
func (r *CurriculumRepository) DeleteCP(ctx context.Context, id string) error {
	query := `DELETE FROM cp WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("cp not found")
	}

	return nil
}

// UpdateCurriculumPhase updates a curriculum phase
func (r *CurriculumRepository) UpdateCurriculumPhase(ctx context.Context, phase *domain.CurriculumPhase) error {
	query := `
		UPDATE curriculum_phases
		SET name = $2, description = $3, grade_level_start = $4, grade_level_end = $5, is_active = $6, updated_at = NOW()
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query,
		phase.ID, phase.Name, phase.Description, phase.GradeLevelStart, phase.GradeLevelEnd, phase.IsActive)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("curriculum phase not found")
	}

	return nil
}

// DeleteCurriculumPhase deletes a curriculum phase
func (r *CurriculumRepository) DeleteCurriculumPhase(ctx context.Context, id string) error {
	query := `DELETE FROM curriculum_phases WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("curriculum phase not found")
	}

	return nil
}

// UpdateCurriculumElement updates a curriculum element
func (r *CurriculumRepository) UpdateCurriculumElement(ctx context.Context, element *domain.CurriculumElement) error {
	query := `
		UPDATE curriculum_elements
		SET name = $2, description = $3, is_active = $4, updated_at = NOW()
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query,
		element.ID, element.Name, element.Description, element.IsActive)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("curriculum element not found")
	}

	return nil
}

// DeleteCurriculumElement deletes a curriculum element
func (r *CurriculumRepository) DeleteCurriculumElement(ctx context.Context, id string) error {
	query := `DELETE FROM curriculum_elements WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("curriculum element not found")
	}

	return nil
}

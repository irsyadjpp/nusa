-- Add composite indexes for common query patterns
-- This migration adds composite indexes to optimize query performance

-- Academic Foundation indexes
CREATE INDEX idx_academic_years_school_status ON academic_years(school_id, status) WHERE deleted_at IS NULL;
CREATE INDEX idx_semesters_academic_year_status ON semesters(academic_year_id, status) WHERE deleted_at IS NULL;
CREATE INDEX idx_cp_alignments_subject_dimension ON cp_alignments(curriculum_subject_id, graduate_profile_dimension_id) WHERE deleted_at IS NULL;

-- Curriculum indexes
CREATE INDEX idx_cps_subject_phase ON cps(subject_id, phase_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_cps_phase_element ON cps(phase_id, element_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_curriculum_phases_subject ON curriculum_phases(subject_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_curriculum_elements_phase ON curriculum_elements(phase_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_curriculum_subelements_element ON curriculum_subelements(element_id) WHERE deleted_at IS NULL;

-- Learning Planning indexes
CREATE INDEX idx_tps_tp_set ON tps(tp_set_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_tps_subject_phase ON tps(subject_id, phase_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_atps_atp_set ON atps(atp_set_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_modul_ajars_modul_ajar_set ON modul_ajars(modul_ajar_set_id) WHERE deleted_at IS NULL;

-- Assessment indexes
CREATE INDEX idx_assessments_tp ON assessments(tp_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_assessments_user_status ON assessments(user_id, status) WHERE deleted_at IS NULL;
CREATE INDEX idx_rubrics_assessment ON rubrics(assessment_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_evidences_assessment ON evidences(assessment_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_evidences_student ON evidences(student_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_evaluations_evidence ON evaluations(evidence_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_evaluations_student ON evaluations(student_id) WHERE deleted_at IS NULL;

-- User Management indexes
CREATE INDEX idx_users_school_role ON users(school_id, role_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_users_email_active ON users(email) WHERE deleted_at IS NULL AND is_active = true;

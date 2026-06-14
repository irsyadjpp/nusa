-- Remove composite indexes
-- This migration removes the composite indexes added in the up migration

-- User Management indexes
DROP INDEX IF EXISTS idx_users_email_active;
DROP INDEX IF EXISTS idx_users_school_role;

-- Assessment indexes
DROP INDEX IF EXISTS idx_evaluations_student;
DROP INDEX IF EXISTS idx_evaluations_evidence;
DROP INDEX IF EXISTS idx_evidences_student;
DROP INDEX IF EXISTS idx_evidences_assessment;
DROP INDEX IF EXISTS idx_rubrics_assessment;
DROP INDEX IF EXISTS idx_assessments_user_status;
DROP INDEX IF EXISTS idx_assessments_tp;

-- Learning Planning indexes
DROP INDEX IF EXISTS idx_modul_ajars_modul_ajar_set;
DROP INDEX IF EXISTS idx_atps_atp_set;
DROP INDEX IF EXISTS idx_tps_subject_phase;
DROP INDEX IF EXISTS idx_tps_tp_set;

-- Curriculum indexes
DROP INDEX IF EXISTS idx_curriculum_subelements_element;
DROP INDEX IF EXISTS idx_curriculum_elements_phase;
DROP INDEX IF EXISTS idx_curriculum_phases_subject;
DROP INDEX IF EXISTS idx_cps_phase_element;
DROP INDEX IF EXISTS idx_cps_subject_phase;

-- Academic Foundation indexes
DROP INDEX IF EXISTS idx_cp_alignments_subject_dimension;
DROP INDEX IF EXISTS idx_semesters_academic_year_status;
DROP INDEX IF EXISTS idx_academic_years_school_status;

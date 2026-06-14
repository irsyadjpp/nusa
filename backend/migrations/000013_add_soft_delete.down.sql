-- Remove soft delete functionality from all tables
-- This migration removes the deleted_at column and partial indexes

-- Drop partial indexes for active records
DROP INDEX IF EXISTS idx_users_active;
DROP INDEX IF EXISTS idx_roles_active;
DROP INDEX IF EXISTS idx_schools_active;
DROP INDEX IF EXISTS idx_academic_years_active;
DROP INDEX IF EXISTS idx_semesters_active;
DROP INDEX IF EXISTS idx_subject_categories_active;
DROP INDEX IF EXISTS idx_graduate_profile_dimensions_active;
DROP INDEX IF EXISTS idx_cp_alignments_active;
DROP INDEX IF EXISTS idx_system_configurations_active;
DROP INDEX IF EXISTS idx_curriculum_subjects_active;
DROP INDEX IF EXISTS idx_curriculum_phases_active;
DROP INDEX IF EXISTS idx_curriculum_elements_active;
DROP INDEX IF EXISTS idx_curriculum_subelements_active;
DROP INDEX IF EXISTS idx_cps_active;
DROP INDEX IF EXISTS idx_tp_sets_active;
DROP INDEX IF EXISTS idx_tps_active;
DROP INDEX IF EXISTS idx_atp_sets_active;
DROP INDEX IF EXISTS idx_atps_active;
DROP INDEX IF EXISTS idx_modul_ajar_sets_active;
DROP INDEX IF EXISTS idx_modul_ajars_active;
DROP INDEX IF EXISTS idx_assessments_active;
DROP INDEX IF EXISTS idx_rubrics_active;
DROP INDEX IF EXISTS idx_evidences_active;
DROP INDEX IF EXISTS idx_evaluations_active;
DROP INDEX IF EXISTS idx_narrative_reports_active;

-- Remove deleted_at column from all tables
ALTER TABLE narrative_reports DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE evaluations DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE evidences DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE rubrics DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE assessments DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE modul_ajars DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE modul_ajar_sets DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE atps DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE atp_sets DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE tps DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE tp_sets DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE cps DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE curriculum_subelements DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE curriculum_elements DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE curriculum_phases DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE curriculum_subjects DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE system_configurations DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE cp_alignments DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE graduate_profile_dimensions DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE subject_categories DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE semesters DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE academic_years DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE schools DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE refresh_tokens DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE permissions DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE roles DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE users DROP COLUMN IF EXISTS deleted_at;

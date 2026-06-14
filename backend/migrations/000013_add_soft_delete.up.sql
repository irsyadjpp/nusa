-- Add deleted_at column to all existing tables for soft delete functionality
-- This migration adds soft delete support to all existing tables

-- Add deleted_at to users table
ALTER TABLE users ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMP;

-- Add deleted_at to roles table
ALTER TABLE roles ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMP;

-- Add deleted_at to permissions table
ALTER TABLE permissions ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMP;

-- Add deleted_at to refresh_tokens table
ALTER TABLE refresh_tokens ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMP;

-- Add deleted_at to schools table
ALTER TABLE schools ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMP;

-- Add deleted_at to academic_years table
ALTER TABLE academic_years ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMP;

-- Add deleted_at to semesters table
ALTER TABLE semesters ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMP;

-- Add deleted_at to subject_categories table
ALTER TABLE subject_categories ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMP;

-- Add deleted_at to graduate_profile_dimensions table
ALTER TABLE graduate_profile_dimensions ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMP;

-- Add deleted_at to cp_alignments table
ALTER TABLE cp_alignments ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMP;

-- Add deleted_at to system_configurations table
ALTER TABLE system_configurations ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMP;

-- Add deleted_at to curriculum_subjects table
ALTER TABLE curriculum_subjects ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMP;

-- Add deleted_at to curriculum_phases table
ALTER TABLE curriculum_phases ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMP;

-- Add deleted_at to curriculum_elements table
ALTER TABLE curriculum_elements ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMP;

-- Add deleted_at to curriculum_subelements table
ALTER TABLE curriculum_subelements ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMP;

-- Add deleted_at to cps table
ALTER TABLE cps ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMP;

-- Add deleted_at to tp_sets table
ALTER TABLE tp_sets ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMP;

-- Add deleted_at to tps table
ALTER TABLE tps ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMP;

-- Add deleted_at to atp_sets table
ALTER TABLE atp_sets ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMP;

-- Add deleted_at to atps table
ALTER TABLE atps ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMP;

-- Add deleted_at to modul_ajar_sets table
ALTER TABLE modul_ajar_sets ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMP;

-- Add deleted_at to modul_ajars table
ALTER TABLE modul_ajars ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMP;

-- Add deleted_at to assessments table
ALTER TABLE assessments ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMP;

-- Add deleted_at to rubrics table
ALTER TABLE rubrics ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMP;

-- Add deleted_at to evidences table
ALTER TABLE evidences ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMP;

-- Add deleted_at to evaluations table
ALTER TABLE evaluations ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMP;

-- Add deleted_at to narrative_reports table
ALTER TABLE narrative_reports ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMP;

-- Add partial indexes for active records (deleted_at IS NULL)
CREATE INDEX idx_users_active ON users(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_roles_active ON roles(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_schools_active ON schools(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_academic_years_active ON academic_years(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_semesters_active ON semesters(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_subject_categories_active ON subject_categories(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_graduate_profile_dimensions_active ON graduate_profile_dimensions(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_cp_alignments_active ON cp_alignments(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_system_configurations_active ON system_configurations(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_curriculum_subjects_active ON curriculum_subjects(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_curriculum_phases_active ON curriculum_phases(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_curriculum_elements_active ON curriculum_elements(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_curriculum_subelements_active ON curriculum_subelements(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_cps_active ON cps(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_tp_sets_active ON tp_sets(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_tps_active ON tps(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_atp_sets_active ON atp_sets(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_atps_active ON atps(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_modul_ajar_sets_active ON modul_ajar_sets(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_modul_ajars_active ON modul_ajars(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_assessments_active ON assessments(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_rubrics_active ON rubrics(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_evidences_active ON evidences(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_evaluations_active ON evaluations(id) WHERE deleted_at IS NULL;
CREATE INDEX idx_narrative_reports_active ON narrative_reports(id) WHERE deleted_at IS NULL;

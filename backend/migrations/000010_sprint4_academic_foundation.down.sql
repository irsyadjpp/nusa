-- Sprint 4 Academic Foundation Rollback Migration
-- Based on: Sprint 4 Requirement Package (Refactored v2)
-- Migration Number: 000010
-- Purpose: Rollback Sprint 4 changes
-- Risk Level: MEDIUM (data loss if not backed up)
-- Notes: This rollback will remove Sprint 4 features

-- ============================================================================
-- ROLLBACK INSTRUCTIONS
-- 1. Ensure you have a recent database backup
-- 2. Run this migration in maintenance window
-- 3. Verify rollback results
-- 4. Document any data loss
-- ============================================================================

-- Drop indexes for table extensions
DROP INDEX IF EXISTS idx_cp_academic_semester;
DROP INDEX IF EXISTS idx_cp_semester_id;
DROP INDEX IF EXISTS idx_cp_academic_year_id;

-- Drop foreign keys for table extensions
ALTER TABLE cp DROP CONSTRAINT IF EXISTS fk_cp_semester_id;
ALTER TABLE cp DROP CONSTRAINT IF EXISTS fk_cp_academic_year_id;

-- Drop columns for table extensions
ALTER TABLE cp DROP COLUMN IF EXISTS semester_id;
ALTER TABLE cp DROP COLUMN IF EXISTS academic_year_id;

DROP INDEX IF EXISTS idx_curriculum_subjects_category_id;
ALTER TABLE curriculum_subjects DROP CONSTRAINT IF EXISTS fk_curriculum_subjects_category_id;
ALTER TABLE curriculum_subjects DROP COLUMN IF EXISTS subject_category_id;

-- Drop indexes for new tables
DROP INDEX IF EXISTS idx_cp_alignments_strength;
DROP INDEX IF EXISTS idx_cp_alignments_dimension_id;
DROP INDEX IF EXISTS idx_cp_alignments_cp_id;

DROP INDEX IF EXISTS idx_graduate_profile_dimensions_indicators;
DROP INDEX IF EXISTS idx_graduate_profile_dimensions_status;
DROP INDEX IF EXISTS idx_graduate_profile_dimensions_name;
DROP INDEX IF EXISTS idx_graduate_profile_dimensions_code;

DROP INDEX IF EXISTS idx_subject_categories_status;
DROP INDEX IF EXISTS idx_subject_categories_name;
DROP INDEX IF EXISTS idx_subject_categories_code;

DROP INDEX IF EXISTS idx_semesters_status;
DROP INDEX IF EXISTS idx_semesters_dates;
DROP INDEX IF EXISTS idx_semesters_sequence;
DROP INDEX IF EXISTS idx_semesters_academic_year_id;

DROP INDEX IF EXISTS idx_academic_years_school_status;
DROP INDEX IF EXISTS idx_academic_years_dates;
DROP INDEX IF EXISTS idx_academic_years_status;
DROP INDEX IF EXISTS idx_academic_years_school_id;

DROP INDEX IF EXISTS idx_system_configuration_key;

-- Drop foreign keys for new tables
ALTER TABLE cp_alignments DROP CONSTRAINT IF EXISTS fk_cp_alignments_created_by;
ALTER TABLE cp_alignments DROP CONSTRAINT IF EXISTS fk_cp_alignments_dimension_id;
ALTER TABLE cp_alignments DROP CONSTRAINT IF EXISTS fk_cp_alignments_cp_id;

ALTER TABLE graduate_profile_dimensions DROP CONSTRAINT IF EXISTS fk_graduate_profile_dimensions_created_by;

ALTER TABLE subject_categories DROP CONSTRAINT IF EXISTS fk_subject_categories_created_by;

ALTER TABLE semesters DROP CONSTRAINT IF EXISTS fk_semesters_academic_year_id;

ALTER TABLE academic_years DROP CONSTRAINT IF EXISTS fk_academic_years_created_by;
ALTER TABLE academic_years DROP CONSTRAINT IF EXISTS fk_academic_years_school_id;

-- Drop new tables
DROP TABLE IF EXISTS cp_alignments;
DROP TABLE IF EXISTS graduate_profile_dimensions;
DROP TABLE IF EXISTS subject_categories;
DROP TABLE IF EXISTS semesters;
DROP TABLE IF EXISTS academic_years;
DROP TABLE IF EXISTS system_configuration;

-- ============================================================================
-- END OF ROLLBACK MIGRATION
-- ============================================================================

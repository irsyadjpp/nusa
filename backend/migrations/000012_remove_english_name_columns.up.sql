-- Migration: Remove English name columns (name_en, description_en) from curriculum foundation tables
-- This removes bilingual support that was deemed unnecessary

-- Remove name_en from curriculum_subjects
ALTER TABLE curriculum_subjects DROP COLUMN IF EXISTS name_en;

-- Remove name_en from curriculum_phases
ALTER TABLE curriculum_phases DROP COLUMN IF EXISTS name_en;

-- Remove name_en from curriculum_elements
ALTER TABLE curriculum_elements DROP COLUMN IF EXISTS name_en;

-- Remove name_en from curriculum_subelements
ALTER TABLE curriculum_subelements DROP COLUMN IF EXISTS name_en;

-- Remove name_en from subject_categories
ALTER TABLE subject_categories DROP COLUMN IF EXISTS name_en;

-- Remove name_en from graduate_profile_dimensions
ALTER TABLE graduate_profile_dimensions DROP COLUMN IF EXISTS name_en;

-- Remove description_en from graduate_profile_dimensions
ALTER TABLE graduate_profile_dimensions DROP COLUMN IF EXISTS description_en;

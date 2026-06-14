-- Rollback: Add back English name columns to curriculum foundation tables

-- Add name_en back to curriculum_subjects
ALTER TABLE curriculum_subjects ADD COLUMN IF NOT EXISTS name_en VARCHAR(100);

-- Add name_en back to curriculum_phases
ALTER TABLE curriculum_phases ADD COLUMN IF NOT EXISTS name_en VARCHAR(100);

-- Add name_en back to curriculum_elements
ALTER TABLE curriculum_elements ADD COLUMN IF NOT EXISTS name_en VARCHAR(100);

-- Add name_en back to curriculum_subelements
ALTER TABLE curriculum_subelements ADD COLUMN IF NOT EXISTS name_en VARCHAR(100);

-- Add name_en back to subject_categories
ALTER TABLE subject_categories ADD COLUMN IF NOT EXISTS name_en VARCHAR(100);

-- Add name_en back to graduate_profile_dimensions
ALTER TABLE graduate_profile_dimensions ADD COLUMN IF NOT EXISTS name_en VARCHAR(100);

-- Add description_en back to graduate_profile_dimensions
ALTER TABLE graduate_profile_dimensions ADD COLUMN IF NOT EXISTS description_en TEXT;

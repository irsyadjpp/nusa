-- Sprint 3A Domain Changes Migration
-- This migration implements the domain model changes for Sprint 3A
-- Changes:
-- 1. Add success_criteria JSONB to TP table
-- 2. Modify assessments table to reference TP instead of ModulAjar
-- 3. Add revision history to evaluations table
-- 4. Create optimal indexes for new query patterns
-- ============================================
-- MIGRATION A: Add success_criteria to TP table
-- ============================================

-- Add success_criteria column to tp table
ALTER TABLE tp 
ADD COLUMN success_criteria JSONB;

-- Add comment for documentation
COMMENT ON COLUMN tp.success_criteria IS 'Embedded KKTPCriteria (Success Criteria) as JSONB - contains mastery thresholds, performance indicators, and minimum requirements';

-- ============================================
-- MIGRATION B: Assessment table changes
-- ============================================

-- Step 1: Add new columns for TP reference
ALTER TABLE assessments
ADD COLUMN tp_id UUID REFERENCES tp(id) ON DELETE CASCADE,
ADD COLUMN tp_version_no INTEGER NOT NULL DEFAULT 1,
ADD COLUMN success_criteria_snapshot JSONB;

-- Add comments for documentation
COMMENT ON COLUMN assessments.tp_id IS 'Reference to TP instead of ModulAjar - supports historical consistency';
COMMENT ON COLUMN assessments.tp_version_no IS 'Snapshot of TP version at assessment creation time';
COMMENT ON COLUMN assessments.success_criteria_snapshot IS 'Snapshot of TP SuccessCriteria at assessment creation time - ensures assessment remains valid even if TP changes';

-- Step 2: Migrate existing data from modul_ajar_id to tp_id
-- This is a data migration step - we need to trace through ModulAjar -> ATP -> TP
UPDATE assessments a
SET tp_id = ma.atp_id,
    tp_version_no = 1,
    success_criteria = COALESCE(t.success_criteria, '{}'::jsonb)
FROM modul_ajar ma
JOIN atp ON ma.atp_id = atp.id
JOIN tp ON atp.tp_id = tp.id
WHERE a.modul_ajar_id = ma.id;

-- Step 3: Make tp_id NOT NULL after data migration
ALTER TABLE assessments
ALTER COLUMN tp_id SET NOT NULL;

-- Step 4: Drop the old modul_ajar_id column and its index
DROP INDEX IF EXISTS idx_assessments_modul_ajar_id;
ALTER TABLE assessments
DROP COLUMN IF EXISTS modul_ajar_id;

-- ============================================
-- MIGRATION C: Evaluation table updates
-- ============================================

-- Add revision history columns to evaluations table
ALTER TABLE evaluations
ADD COLUMN teacher_feedback TEXT,
ADD COLUMN revision_no INTEGER NOT NULL DEFAULT 1,
ADD COLUMN created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
ADD COLUMN updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW();

-- Add comments for documentation
COMMENT ON COLUMN evaluations.teacher_feedback IS 'Teacher feedback on the evaluation';
COMMENT ON COLUMN evaluations.revision_no IS 'Revision number for tracking evaluation history';
COMMENT ON COLUMN evaluations.created_at IS 'Timestamp when evaluation was first created';
COMMENT ON COLUMN evaluations.updated_at IS 'Timestamp when evaluation was last updated';

-- Update existing evaluated_at to be consistent with new schema
-- For existing records, set created_at and updated_at to evaluated_at
UPDATE evaluations
SET created_at = evaluated_at,
    updated_at = evaluated_at
WHERE created_at IS NULL OR updated_at IS NULL;

-- ============================================
-- MIGRATION D: Create optimal indexes
-- ============================================

-- Indexes for TP table
CREATE INDEX idx_tp_success_criteria ON tp USING GIN (success_criteria);
CREATE INDEX idx_tp_tp_set_id_sequence ON tp(tp_set_id, sequence_number);

-- Indexes for assessments table (new query patterns)
CREATE INDEX idx_assessments_tp_id ON assessments(tp_id);
CREATE INDEX idx_assessments_tp_version ON assessments(tp_id, tp_version_no);
CREATE INDEX idx_assessments_success_criteria ON assessments USING GIN (success_criteria_snapshot);

-- Indexes for evaluations table (new query patterns)
CREATE INDEX idx_evaluations_revision ON evaluations(evidence_id, revision_no);
CREATE INDEX idx_evaluations_teacher_feedback ON evaluations(teacher_feedback) WHERE teacher_feedback IS NOT NULL;
CREATE INDEX idx_evaluations_created_at ON evaluations(created_at);
CREATE INDEX idx_evaluations_updated_at ON evaluations(updated_at);

-- Composite index for common lookup pattern (student + evidence + revision)
CREATE INDEX idx_evaluations_student_evidence_revision ON evaluations(student_id, evidence_id, revision_no DESC);

-- ============================================
-- DATA VALIDATION CHECKS
-- ============================================

-- Ensure all assessments have valid tp_id
DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM assessments WHERE tp_id IS NULL) THEN
        RAISE EXCEPTION 'Data validation failed: Some assessments have NULL tp_id after migration';
    END IF;
END $$;

-- Ensure all evaluations have revision_no
DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM evaluations WHERE revision_no IS NULL) THEN
        RAISE EXCEPTION 'Data validation failed: Some evaluations have NULL revision_no after migration';
    END IF;
END $$;

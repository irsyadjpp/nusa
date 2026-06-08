-- Sprint 3A Domain Changes Rollback Migration
-- This migration rolls back the domain model changes for Sprint 3A
-- ============================================
-- ROLLBACK MIGRATION D: Drop indexes
-- ============================================

DROP INDEX IF EXISTS idx_evaluations_student_evidence_revision;
DROP INDEX IF EXISTS idx_evaluations_updated_at;
DROP INDEX IF EXISTS idx_evaluations_created_at;
DROP INDEX IF EXISTS idx_evaluations_teacher_feedback;
DROP INDEX IF EXISTS idx_evaluations_revision;
DROP INDEX IF EXISTS idx_assessments_success_criteria;
DROP INDEX IF EXISTS idx_assessments_tp_version;
DROP INDEX IF EXISTS idx_assessments_tp_id;
DROP INDEX IF EXISTS idx_tp_tp_set_id_sequence;
DROP INDEX IF EXISTS idx_tp_success_criteria;

-- ============================================
-- ROLLBACK MIGRATION C: Evaluation table updates
-- ============================================

ALTER TABLE evaluations
DROP COLUMN IF EXISTS updated_at,
DROP COLUMN IF EXISTS created_at,
DROP COLUMN IF EXISTS revision_no,
DROP COLUMN IF EXISTS teacher_feedback;

-- ============================================
-- ROLLBACK MIGRATION B: Assessment table changes
-- ============================================

-- Step 1: Re-add modul_ajar_id column
ALTER TABLE assessments
ADD COLUMN modul_ajar_id UUID;

-- Step 2: Restore data from tp_id to modul_ajar_id
-- This is a best-effort restoration - some data may be lost if the chain is broken
UPDATE assessments a
SET modul_ajar_id = ma.id
FROM tp
JOIN atp ON atp.tp_id = tp.id
JOIN modul_ajar ma ON ma.atp_id = atp.id
WHERE a.tp_id = tp.id;

-- Step 3: Make modul_ajar_id NOT NULL (may fail if data restoration incomplete)
-- Commented out to allow rollback even if data is incomplete
-- ALTER TABLE assessments ALTER COLUMN modul_ajar_id SET NOT NULL;

-- Step 4: Recreate index for modul_ajar_id
CREATE INDEX idx_assessments_modul_ajar_id ON assessments(modul_ajar_id);

-- Step 5: Drop new TP-related columns
ALTER TABLE assessments
DROP COLUMN IF EXISTS success_criteria_snapshot,
DROP COLUMN IF EXISTS tp_version_no,
DROP COLUMN IF EXISTS tp_id;

-- ============================================
-- ROLLBACK MIGRATION A: Remove success_criteria from TP table
-- ============================================

ALTER TABLE tp
DROP COLUMN IF EXISTS success_criteria;

-- Remove class_id from narrative_reports
DROP INDEX IF EXISTS idx_narrative_reports_student_class;
DROP INDEX IF EXISTS idx_narrative_reports_class;
ALTER TABLE narrative_reports DROP COLUMN IF EXISTS class_id;

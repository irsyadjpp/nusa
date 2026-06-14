-- Drop assignments table
DROP INDEX IF EXISTS idx_assignments_class_due;
DROP INDEX IF EXISTS idx_assignments_status;
DROP INDEX IF EXISTS idx_assignments_due_date;
DROP INDEX IF EXISTS idx_assignments_assessment_id;
DROP INDEX IF EXISTS idx_assignments_class_id;
DROP TABLE IF EXISTS assignments;

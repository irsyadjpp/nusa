-- Drop class_enrollments table
DROP INDEX IF EXISTS idx_class_enrollments_active;
DROP INDEX IF EXISTS idx_class_enrollments_status;
DROP INDEX IF EXISTS idx_class_enrollments_student_id;
DROP INDEX IF EXISTS idx_class_enrollments_class_id;
DROP TABLE IF EXISTS class_enrollments;

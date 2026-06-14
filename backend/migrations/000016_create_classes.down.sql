-- Drop classes table
DROP INDEX IF EXISTS idx_classes_active;
DROP INDEX IF EXISTS idx_classes_school_semester;
DROP INDEX IF EXISTS idx_classes_teacher_id;
DROP INDEX IF EXISTS idx_classes_subject_id;
DROP INDEX IF EXISTS idx_classes_semester_id;
DROP INDEX IF EXISTS idx_classes_academic_year_id;
DROP INDEX IF EXISTS idx_classes_school_id;
DROP TABLE IF EXISTS classes;

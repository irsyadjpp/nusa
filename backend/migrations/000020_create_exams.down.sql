-- Drop exams table
DROP INDEX IF EXISTS idx_exams_class_date;
DROP INDEX IF EXISTS idx_exams_status;
DROP INDEX IF EXISTS idx_exams_exam_date;
DROP INDEX IF EXISTS idx_exams_assessment_id;
DROP INDEX IF EXISTS idx_exams_class_id;
DROP TABLE IF EXISTS exams;

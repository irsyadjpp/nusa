-- Drop attendance_records table
DROP INDEX IF EXISTS idx_attendance_records_student_date;
DROP INDEX IF EXISTS idx_attendance_records_class_date;
DROP INDEX IF EXISTS idx_attendance_records_status;
DROP INDEX IF EXISTS idx_attendance_records_date;
DROP INDEX IF EXISTS idx_attendance_records_student_id;
DROP INDEX IF EXISTS idx_attendance_records_class_id;
DROP TABLE IF EXISTS attendance_records;

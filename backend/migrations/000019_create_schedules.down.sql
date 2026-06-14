-- Drop schedules table
DROP INDEX IF EXISTS idx_schedules_room;
DROP INDEX IF EXISTS idx_schedules_active;
DROP INDEX IF EXISTS idx_schedules_day_of_week;
DROP INDEX IF EXISTS idx_schedules_class_id;
DROP TABLE IF EXISTS schedules;

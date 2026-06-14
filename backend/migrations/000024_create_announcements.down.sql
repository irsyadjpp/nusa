-- Drop announcements table
DROP INDEX IF EXISTS idx_announcements_school_active;
DROP INDEX IF EXISTS idx_announcements_active;
DROP INDEX IF EXISTS idx_announcements_target_audience;
DROP INDEX IF EXISTS idx_announcements_priority;
DROP INDEX IF EXISTS idx_announcements_school_id;
DROP TABLE IF EXISTS announcements;

DROP INDEX IF EXISTS idx_narrative_reports_student_period;
DROP INDEX IF EXISTS idx_narrative_reports_achievement;
ALTER TABLE narrative_reports DROP COLUMN IF EXISTS last_achievement_calculated_at;
ALTER TABLE narrative_reports DROP COLUMN IF EXISTS achievement_data;
ALTER TABLE narrative_reports DROP COLUMN IF EXISTS achievement_summary_id;

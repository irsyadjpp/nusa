-- Add achievement fields to narrative_reports
ALTER TABLE narrative_reports
ADD COLUMN achievement_summary_id UUID,
ADD COLUMN achievement_data JSONB,
ADD COLUMN last_achievement_calculated_at TIMESTAMP WITH TIME ZONE;

-- Create index for achievement data queries
CREATE INDEX idx_narrative_reports_achievement ON narrative_reports(achievement_summary_id);
CREATE INDEX idx_narrative_reports_student_period ON narrative_reports(student_id, report_period);

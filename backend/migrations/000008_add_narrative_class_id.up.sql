-- Add class_id to narrative_reports for school filtering
ALTER TABLE narrative_reports
ADD COLUMN class_id UUID REFERENCES classes(id) ON DELETE SET NULL;

-- Create index for class-based filtering
CREATE INDEX idx_narrative_reports_class ON narrative_reports(class_id);
CREATE INDEX idx_narrative_reports_student_class ON narrative_reports(student_id, class_id);

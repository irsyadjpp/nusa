-- Add expanded TP snapshot fields to assessments
ALTER TABLE assessments
ADD COLUMN tp_title_snapshot TEXT,
ADD COLUMN tp_learning_objectives_snapshot JSONB,
ADD COLUMN tp_time_allocation_snapshot JSONB;

-- Create index for snapshot queries
CREATE INDEX idx_assessments_tp_snapshot ON assessments(tp_version_no);

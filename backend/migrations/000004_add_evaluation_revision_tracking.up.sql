-- Add revision tracking fields to evaluations
ALTER TABLE evaluations 
ADD COLUMN revision_no INTEGER NOT NULL DEFAULT 1,
ADD COLUMN is_current_version BOOLEAN NOT NULL DEFAULT true,
ADD COLUMN parent_revision_id UUID REFERENCES evaluations(id);

-- Create index for revision queries
CREATE INDEX idx_evaluations_evidence_revision ON evaluations(evidence_id, revision_no);
CREATE INDEX idx_evaluations_parent_revision ON evaluations(parent_revision_id);

-- Create feedback history table
CREATE TABLE evaluation_feedback_history (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    evaluation_id UUID NOT NULL REFERENCES evaluations(id) ON DELETE CASCADE,
    teacher_feedback TEXT NOT NULL,
    changed_by UUID NOT NULL REFERENCES users(id),
    changed_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_feedback_history_evaluation ON evaluation_feedback_history(evaluation_id);
CREATE INDEX idx_feedback_history_changed_at ON evaluation_feedback_history(changed_at);

-- Update existing evaluations
UPDATE evaluations SET revision_no = 1, is_current_version = true WHERE revision_no IS NULL;

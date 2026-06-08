-- Add version tracking fields to tps
ALTER TABLE tps
ADD COLUMN version_no INTEGER NOT NULL DEFAULT 1,
ADD COLUMN is_current_version BOOLEAN NOT NULL DEFAULT true,
ADD COLUMN parent_version_id UUID REFERENCES tps(id);

-- Create index for version queries
CREATE INDEX idx_tps_set_version ON tps(tp_set_id, version_no);
CREATE INDEX idx_tps_parent_version ON tps(parent_version_id);

-- Update existing TPs
UPDATE tps SET version_no = 1, is_current_version = true WHERE version_no IS NULL;

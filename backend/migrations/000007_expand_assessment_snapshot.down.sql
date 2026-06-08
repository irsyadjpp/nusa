DROP INDEX IF EXISTS idx_assessments_tp_snapshot;
ALTER TABLE assessments DROP COLUMN IF EXISTS tp_time_allocation_snapshot;
ALTER TABLE assessments DROP COLUMN IF EXISTS tp_learning_objectives_snapshot;
ALTER TABLE assessments DROP COLUMN IF EXISTS tp_title_snapshot;

DROP INDEX IF EXISTS idx_tps_parent_version;
DROP INDEX IF EXISTS idx_tps_set_version;
ALTER TABLE tps DROP COLUMN IF EXISTS parent_version_id;
ALTER TABLE tps DROP COLUMN IF EXISTS is_current_version;
ALTER TABLE tps DROP COLUMN IF EXISTS version_no;

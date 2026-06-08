DROP TABLE IF EXISTS evaluation_feedback_history;
ALTER TABLE evaluations DROP COLUMN IF EXISTS parent_revision_id;
ALTER TABLE evaluations DROP COLUMN IF EXISTS is_current_version;
ALTER TABLE evaluations DROP COLUMN IF EXISTS revision_no;

-- Drop job_queue table
DROP INDEX IF EXISTS idx_job_queue_status_scheduled;
DROP INDEX IF EXISTS idx_job_queue_status_priority;
DROP INDEX IF EXISTS idx_job_queue_scheduled_at;
DROP INDEX IF EXISTS idx_job_queue_priority;
DROP INDEX IF EXISTS idx_job_queue_job_type;
DROP INDEX IF EXISTS idx_job_queue_status;
DROP TABLE IF EXISTS job_queue;

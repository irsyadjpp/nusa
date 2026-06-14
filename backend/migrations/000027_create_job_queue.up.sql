-- Create job_queue table for background job processing
CREATE TABLE job_queue (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  job_type VARCHAR(255) NOT NULL,
  payload JSONB NOT NULL,
  status VARCHAR(50) DEFAULT 'PENDING' CHECK (status IN ('PENDING', 'PROCESSING', 'COMPLETED', 'FAILED', 'CANCELLED')),
  priority INTEGER DEFAULT 5 CHECK (priority >= 1 AND priority <= 10),
  max_retries INTEGER DEFAULT 3,
  retry_count INTEGER DEFAULT 0,
  error_message TEXT,
  started_at TIMESTAMP,
  completed_at TIMESTAMP,
  scheduled_at TIMESTAMP DEFAULT NOW(),
  created_at TIMESTAMP DEFAULT NOW()
);

-- Indexes for job_queue table
CREATE INDEX idx_job_queue_status ON job_queue(status);
CREATE INDEX idx_job_queue_job_type ON job_queue(job_type);
CREATE INDEX idx_job_queue_priority ON job_queue(priority);
CREATE INDEX idx_job_queue_scheduled_at ON job_queue(scheduled_at);
CREATE INDEX idx_job_queue_status_priority ON job_queue(status, priority);
CREATE INDEX idx_job_queue_status_scheduled ON job_queue(status, scheduled_at);

package queue

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// JobStatus represents the status of a job
type JobStatus string

const (
	JobStatusPending    JobStatus = "PENDING"
	JobStatusProcessing JobStatus = "PROCESSING"
	JobStatusCompleted  JobStatus = "COMPLETED"
	JobStatusFailed     JobStatus = "FAILED"
	JobStatusCancelled  JobStatus = "CANCELLED"
)

// Job represents a job in the queue
type Job struct {
	ID           uuid.UUID      `json:"id"`
	JobType      string         `json:"job_type"`
	Payload      map[string]any `json:"payload"`
	Status       JobStatus      `json:"status"`
	Priority     int            `json:"priority"`
	MaxRetries   int            `json:"max_retries"`
	RetryCount   int            `json:"retry_count"`
	ErrorMessage string         `json:"error_message,omitempty"`
	StartedAt    *time.Time     `json:"started_at,omitempty"`
	CompletedAt  *time.Time     `json:"completed_at,omitempty"`
	ScheduledAt  time.Time      `json:"scheduled_at"`
	CreatedAt    time.Time      `json:"created_at"`
}

// JobHandler is a function that processes a job
type JobHandler func(ctx context.Context, job *Job) error

// JobQueue manages background jobs using PostgreSQL as the backend
type JobQueue struct {
	db       *sql.DB
	handlers map[string]JobHandler
}

// NewJobQueue creates a new job queue instance
func NewJobQueue(db *sql.DB) *JobQueue {
	return &JobQueue{
		db:       db,
		handlers: make(map[string]JobHandler),
	}
}

// RegisterHandler registers a handler for a specific job type
func (jq *JobQueue) RegisterHandler(jobType string, handler JobHandler) {
	jq.handlers[jobType] = handler
}

// Enqueue adds a new job to the queue
func (jq *JobQueue) Enqueue(ctx context.Context, jobType string, payload map[string]any, priority int, scheduledAt time.Time) (*Job, error) {
	job := &Job{
		ID:          uuid.New(),
		JobType:     jobType,
		Payload:     payload,
		Status:      JobStatusPending,
		Priority:    priority,
		MaxRetries:  3,
		RetryCount:  0,
		ScheduledAt: scheduledAt,
		CreatedAt:   time.Now(),
	}

	query := `
		INSERT INTO job_queue (id, job_type, payload, status, priority, max_retries, retry_count, scheduled_at, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	_, err = jq.db.ExecContext(ctx, query,
		job.ID,
		job.JobType,
		payloadJSON,
		job.Status,
		job.Priority,
		job.MaxRetries,
		job.RetryCount,
		job.ScheduledAt,
		job.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to enqueue job: %w", err)
	}

	return job, nil
}

// Dequeue retrieves the next available job from the queue
func (jq *JobQueue) Dequeue(ctx context.Context) (*Job, error) {
	query := `
		SELECT id, job_type, payload, status, priority, max_retries, retry_count, error_message, started_at, completed_at, scheduled_at, created_at
		FROM job_queue
		WHERE status = $1 AND scheduled_at <= $2
		ORDER BY priority DESC, created_at ASC
		LIMIT 1
		FOR UPDATE SKIP LOCKED
	`

	row := jq.db.QueryRowContext(ctx, query, JobStatusPending, time.Now())

	var job Job
	var payloadJSON []byte
	var errorMessage sql.NullString
	var startedAt, completedAt sql.NullTime

	err := row.Scan(
		&job.ID,
		&job.JobType,
		&payloadJSON,
		&job.Status,
		&job.Priority,
		&job.MaxRetries,
		&job.RetryCount,
		&errorMessage,
		&startedAt,
		&completedAt,
		&job.ScheduledAt,
		&job.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil // No jobs available
	}

	if err != nil {
		return nil, fmt.Errorf("failed to dequeue job: %w", err)
	}

	if err := json.Unmarshal(payloadJSON, &job.Payload); err != nil {
		return nil, fmt.Errorf("failed to unmarshal payload: %w", err)
	}

	if errorMessage.Valid {
		job.ErrorMessage = errorMessage.String
	}

	if startedAt.Valid {
		job.StartedAt = &startedAt.Time
	}

	if completedAt.Valid {
		job.CompletedAt = &completedAt.Time
	}

	// Mark job as processing
	now := time.Now()
	updateQuery := `
		UPDATE job_queue
		SET status = $1, started_at = $2
		WHERE id = $3
	`

	_, err = jq.db.ExecContext(ctx, updateQuery, JobStatusProcessing, now, job.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to mark job as processing: %w", err)
	}

	job.Status = JobStatusProcessing
	job.StartedAt = &now

	return &job, nil
}

// Complete marks a job as completed
func (jq *JobQueue) Complete(ctx context.Context, jobID uuid.UUID) error {
	query := `
		UPDATE job_queue
		SET status = $1, completed_at = $2
		WHERE id = $3
	`

	now := time.Now()
	_, err := jq.db.ExecContext(ctx, query, JobStatusCompleted, now, jobID)
	if err != nil {
		return fmt.Errorf("failed to complete job: %w", err)
	}

	return nil
}

// Fail marks a job as failed
func (jq *JobQueue) Fail(ctx context.Context, jobID uuid.UUID, errorMessage string) error {
	query := `
		UPDATE job_queue
		SET status = $1, error_message = $2, completed_at = $3
		WHERE id = $4
	`

	now := time.Now()
	_, err := jq.db.ExecContext(ctx, query, JobStatusFailed, errorMessage, now, jobID)
	if err != nil {
		return fmt.Errorf("failed to mark job as failed: %w", err)
	}

	return nil
}

// Retry increments the retry count and requeues the job
func (jq *JobQueue) Retry(ctx context.Context, jobID uuid.UUID, errorMessage string) error {
	query := `
		UPDATE job_queue
		SET status = $1, retry_count = retry_count + 1, error_message = $2, started_at = NULL
		WHERE id = $3
	`

	_, err := jq.db.ExecContext(ctx, query, JobStatusPending, errorMessage, jobID)
	if err != nil {
		return fmt.Errorf("failed to retry job: %w", err)
	}

	return nil
}

// Process processes a single job
func (jq *JobQueue) Process(ctx context.Context, job *Job) error {
	handler, exists := jq.handlers[job.JobType]
	if !exists {
		return fmt.Errorf("no handler registered for job type: %s", job.JobType)
	}

	if err := handler(ctx, job); err != nil {
		// Check if we should retry
		if job.RetryCount < job.MaxRetries {
			return jq.Retry(ctx, job.ID, err.Error())
		}
		return jq.Fail(ctx, job.ID, err.Error())
	}

	return jq.Complete(ctx, job.ID)
}

// StartWorker starts a worker that processes jobs from the queue
func (jq *JobQueue) StartWorker(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			job, err := jq.Dequeue(ctx)
			if err != nil {
				continue
			}

			if job != nil {
				if err := jq.Process(ctx, job); err != nil {
					// Log error but continue processing
					fmt.Printf("Failed to process job %s: %v\n", job.ID, err)
				}
			}
		}
	}
}

// GetJobStatus retrieves the status of a specific job
func (jq *JobQueue) GetJobStatus(ctx context.Context, jobID uuid.UUID) (*Job, error) {
	query := `
		SELECT id, job_type, payload, status, priority, max_retries, retry_count, error_message, started_at, completed_at, scheduled_at, created_at
		FROM job_queue
		WHERE id = $1
	`

	row := jq.db.QueryRowContext(ctx, query, jobID)

	var job Job
	var payloadJSON []byte
	var errorMessage sql.NullString
	var startedAt, completedAt sql.NullTime

	err := row.Scan(
		&job.ID,
		&job.JobType,
		&payloadJSON,
		&job.Status,
		&job.Priority,
		&job.MaxRetries,
		&job.RetryCount,
		&errorMessage,
		&startedAt,
		&completedAt,
		&job.ScheduledAt,
		&job.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get job status: %w", err)
	}

	if err := json.Unmarshal(payloadJSON, &job.Payload); err != nil {
		return nil, fmt.Errorf("failed to unmarshal payload: %w", err)
	}

	if errorMessage.Valid {
		job.ErrorMessage = errorMessage.String
	}

	if startedAt.Valid {
		job.StartedAt = &startedAt.Time
	}

	if completedAt.Valid {
		job.CompletedAt = &completedAt.Time
	}

	return &job, nil
}

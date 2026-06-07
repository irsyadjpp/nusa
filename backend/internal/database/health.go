package database

import (
	"context"
	"time"
)

type HealthStatus struct {
	Status    string
	Latency   time.Duration
	Timestamp time.Time
}

func (d *Database) HealthCheck(ctx context.Context) HealthStatus {
	start := time.Now()
	err := d.Ping(ctx)
	latency := time.Since(start)

	if err != nil {
		return HealthStatus{
			Status:    "unhealthy",
			Latency:   latency,
			Timestamp: time.Now(),
		}
	}

	return HealthStatus{
		Status:    "healthy",
		Latency:   latency,
		Timestamp: time.Now(),
	}
}

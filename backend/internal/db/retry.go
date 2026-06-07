package db

import (
	"context"
	"fmt"
	"time"

	"github.com/nusa/backend/internal/config"
)

type RetryConfig struct {
	MaxRetries      int
	InitialInterval time.Duration
	MaxInterval     time.Duration
	Multiplier      float64
}

func DefaultRetryConfig() RetryConfig {
	return RetryConfig{
		MaxRetries:      3,
		InitialInterval: 100 * time.Millisecond,
		MaxInterval:     5 * time.Second,
		Multiplier:      2.0,
	}
}

func WithRetry(ctx context.Context, cfg RetryConfig, fn func() error) error {
	var lastErr error
	interval := cfg.InitialInterval

	for attempt := 0; attempt <= cfg.MaxRetries; attempt++ {
		if attempt > 0 {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(interval):
			}
		}

		err := fn()
		if err == nil {
			return nil
		}

		lastErr = err

		interval = time.Duration(float64(interval) * cfg.Multiplier)
		if interval > cfg.MaxInterval {
			interval = cfg.MaxInterval
		}
	}

	return fmt.Errorf("after %d retries, last error: %w", cfg.MaxRetries, lastErr)
}

func (p *Postgres) ConnectWithRetry(ctx context.Context, cfg *config.DatabaseConfig, retryCfg RetryConfig) (*Postgres, error) {
	var pg *Postgres
	var err error

	err = WithRetry(ctx, retryCfg, func() error {
		pg, err = NewPostgres(cfg)
		return err
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database after retries: %w", err)
	}

	return pg, nil
}

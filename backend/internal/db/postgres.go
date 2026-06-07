package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jmoiron/sqlx"
	"github.com/nusa/backend/internal/config"
)

type Postgres struct {
	pool       *pgxpool.Pool
	connString string
}

func NewPostgres(cfg *config.DatabaseConfig) (*Postgres, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database config: %w", err)
	}

	config.MaxConns = int32(cfg.MaxOpenConns)
	config.MinConns = int32(cfg.MaxIdleConns)
	config.MaxConnLifetime = cfg.ConnMaxLifetime
	config.MaxConnIdleTime = cfg.ConnMaxIdleTime
	config.HealthCheckPeriod = 1 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Postgres{pool: pool, connString: connString}, nil
}

func (p *Postgres) GetPool() *pgxpool.Pool {
	return p.pool
}

// GetSQLXDB returns a sqlx.DB instance for compatibility with repositories
func (p *Postgres) GetSQLXDB() (*sqlx.DB, error) {
	return sqlx.Connect("postgres", p.connString)
}

func (p *Postgres) Close() error {
	p.pool.Close()
	return nil
}

func (p *Postgres) HealthCheck(ctx context.Context) error {
	return p.pool.Ping(ctx)
}

func (p *Postgres) BeginTx(ctx context.Context) (pgx.Tx, error) {
	return p.pool.Begin(ctx)
}

type Tx interface {
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

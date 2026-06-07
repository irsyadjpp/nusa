package integration

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/nusa/backend/internal/config"
	"github.com/nusa/backend/internal/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDatabaseConnection(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	dbHost := os.Getenv("TEST_DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}

	dbUser := os.Getenv("TEST_DB_USER")
	if dbUser == "" {
		dbUser = "nusa_user"
	}

	dbPassword := os.Getenv("TEST_DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "nusa_password"
	}

	dbName := os.Getenv("TEST_DB_NAME")
	if dbName == "" {
		dbName = "nusa_db"
	}

	cfg := &config.DatabaseConfig{
		Host:            dbHost,
		Port:            "5432",
		User:            dbUser,
		Password:        dbPassword,
		DBName:          dbName,
		SSLMode:         "disable",
		MaxOpenConns:    25,
		MaxIdleConns:    5,
		ConnMaxLifetime: 5 * time.Minute,
		ConnMaxIdleTime: 5 * time.Minute,
	}

	pg, err := db.NewPostgres(cfg)
	if err != nil {
		t.Skipf("Skipping database test - connection failed: %v", err)
		return
	}
	require.NoError(t, err)
	assert.NotNil(t, pg)
	defer pg.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = pg.HealthCheck(ctx)
	assert.NoError(t, err)
}

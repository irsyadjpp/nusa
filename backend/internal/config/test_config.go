package config

import (
	"os"
	"time"
)

// LoadTestConfig loads configuration for testing
func LoadTestConfig() *Config {
	return &Config{
		AppEnv: "test",
		Server: ServerConfig{
			Port:         ":8081",
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
			Environment:  "test",
		},
		Database: DatabaseConfig{
			Host:            getEnv("TEST_DB_HOST", "localhost"),
			Port:            getEnv("TEST_DB_PORT", "5432"),
			User:            getEnv("TEST_DB_USER", "nusa_test"),
			Password:        getEnv("TEST_DB_PASSWORD", "test_password"),
			DBName:          getEnv("TEST_DB_NAME", "nusa_test"),
			SSLMode:         getEnv("TEST_DB_SSLMODE", "disable"),
			MaxOpenConns:    getIntEnv("TEST_DB_MAX_OPEN_CONNS", 5),
			MaxIdleConns:    getIntEnv("TEST_DB_MAX_IDLE_CONNS", 2),
			ConnMaxLifetime: getDurationEnv("TEST_DB_CONN_MAX_LIFETIME", 5*time.Minute),
			ConnMaxIdleTime: getDurationEnv("TEST_DB_CONN_MAX_IDLE_TIME", 5*time.Minute),
		},
		JWT: JWTConfig{
			Secret:     "test-secret-key",
			Expiration: 1 * time.Hour,
		},
		RabbitMQ: RabbitMQConfig{
			Host:     "localhost",
			Port:     "5672",
			User:     "guest",
			Password: "guest",
			Queue:    "test_ai_generation",
		},
		AI: AIConfig{
			PrimaryProvider:  "openai",
			FallbackProvider: "gemini",
			OpenAIKey:        "",
			GeminiKey:        "",
			Timeout:          30 * time.Second,
			MaxRetries:       3,
		},
	}
}

// SetTestEnv sets environment variables for testing
func SetTestEnv() {
	os.Setenv("APP_ENV", "test")
	os.Setenv("TEST_DB_HOST", "localhost")
	os.Setenv("TEST_DB_PORT", "5432")
	os.Setenv("TEST_DB_USER", "nusa_test")
	os.Setenv("TEST_DB_PASSWORD", "test_password")
	os.Setenv("TEST_DB_NAME", "nusa_test")
	os.Setenv("JWT_SECRET", "test-secret-key")
}

// UnsetTestEnv unsets environment variables after testing
func UnsetTestEnv() {
	os.Unsetenv("APP_ENV")
	os.Unsetenv("TEST_DB_HOST")
	os.Unsetenv("TEST_DB_PORT")
	os.Unsetenv("TEST_DB_USER")
	os.Unsetenv("TEST_DB_PASSWORD")
	os.Unsetenv("TEST_DB_NAME")
	os.Unsetenv("JWT_SECRET")
}

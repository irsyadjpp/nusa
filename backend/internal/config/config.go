package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	AppEnv   string
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	RabbitMQ RabbitMQConfig
	AI       AIConfig
}

type ServerConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Environment  string
}

type DatabaseConfig struct {
	Host            string
	Port            string
	User            string
	Password        string
	DBName          string
	SSLMode         string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
}

type JWTConfig struct {
	Secret     string
	Expiration time.Duration
}

type RabbitMQConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Queue    string
}

type AIConfig struct {
	PrimaryProvider  string
	FallbackProvider string
	OpenAIKey        string
	GeminiKey        string
	Timeout          time.Duration
	MaxRetries       int
}

func Load() (*Config, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("../../")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		// If .env file doesn't exist, continue with environment variables
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	if err := validate(&config); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	return &config, nil
}

func validate(c *Config) error {
	// AppEnv
	if c.AppEnv == "" {
		c.AppEnv = getEnv("APP_ENV", "development")
	}

	// Server Config
	if c.Server.Port == "" {
		c.Server.Port = getEnv("SERVER_PORT", ":8080")
	}
	if c.Server.ReadTimeout == 0 {
		c.Server.ReadTimeout = getDurationEnv("SERVER_READ_TIMEOUT", 15*time.Second)
	}
	if c.Server.WriteTimeout == 0 {
		c.Server.WriteTimeout = getDurationEnv("SERVER_WRITE_TIMEOUT", 15*time.Second)
	}
	if c.Server.Environment == "" {
		c.Server.Environment = c.AppEnv
	}

	// Database Config
	if c.Database.Host == "" {
		c.Database.Host = getEnv("DB_HOST", "localhost")
	}
	if c.Database.Port == "" {
		c.Database.Port = getEnv("DB_PORT", "5432")
	}
	if c.Database.User == "" {
		c.Database.User = getEnv("DB_USER", "nusa")
	}
	if c.Database.Password == "" {
		c.Database.Password = getEnv("DB_PASSWORD", "")
		if c.Database.Password == "" {
			return fmt.Errorf("DB_PASSWORD is required")
		}
	}
	if c.Database.DBName == "" {
		c.Database.DBName = getEnv("DB_NAME", "nusa")
	}
	if c.Database.SSLMode == "" {
		c.Database.SSLMode = getEnv("DB_SSLMODE", "disable")
	}
	if c.Database.MaxOpenConns == 0 {
		c.Database.MaxOpenConns = getIntEnv("DB_MAX_OPEN_CONNS", 25)
	}
	if c.Database.MaxIdleConns == 0 {
		c.Database.MaxIdleConns = getIntEnv("DB_MAX_IDLE_CONNS", 5)
	}
	if c.Database.ConnMaxLifetime == 0 {
		c.Database.ConnMaxLifetime = getDurationEnv("DB_CONN_MAX_LIFETIME", 5*time.Minute)
	}
	if c.Database.ConnMaxIdleTime == 0 {
		c.Database.ConnMaxIdleTime = getDurationEnv("DB_CONN_MAX_IDLE_TIME", 5*time.Minute)
	}

	// JWT Config
	if c.JWT.Secret == "" {
		c.JWT.Secret = getEnv("JWT_SECRET", "change-me-in-production")
		if c.JWT.Secret == "change-me-in-production" && c.Server.Environment == "production" {
			return fmt.Errorf("JWT_SECRET must be set in production")
		}
	}
	if c.JWT.Expiration == 0 {
		c.JWT.Expiration = getDurationEnv("JWT_EXPIRATION", 24*time.Hour)
	}

	// RabbitMQ Config
	if c.RabbitMQ.Host == "" {
		c.RabbitMQ.Host = getEnv("RABBITMQ_HOST", "localhost")
	}
	if c.RabbitMQ.Port == "" {
		c.RabbitMQ.Port = getEnv("RABBITMQ_PORT", "5672")
	}
	if c.RabbitMQ.User == "" {
		c.RabbitMQ.User = getEnv("RABBITMQ_USER", "guest")
	}
	if c.RabbitMQ.Password == "" {
		c.RabbitMQ.Password = getEnv("RABBITMQ_PASSWORD", "guest")
	}
	if c.RabbitMQ.Queue == "" {
		c.RabbitMQ.Queue = getEnv("RABBITMQ_QUEUE", "ai_generation")
	}

	// AI Config
	if c.AI.PrimaryProvider == "" {
		c.AI.PrimaryProvider = getEnv("AI_PRIMARY_PROVIDER", "openai")
	}
	if c.AI.FallbackProvider == "" {
		c.AI.FallbackProvider = getEnv("AI_FALLBACK_PROVIDER", "gemini")
	}
	if c.AI.OpenAIKey == "" {
		c.AI.OpenAIKey = getEnv("AI_OPENAI_KEY", "")
	}
	if c.AI.GeminiKey == "" {
		c.AI.GeminiKey = getEnv("AI_GEMINI_KEY", "")
	}
	if c.AI.Timeout == 0 {
		c.AI.Timeout = getDurationEnv("AI_TIMEOUT", 30*time.Second)
	}
	if c.AI.MaxRetries == 0 {
		c.AI.MaxRetries = getIntEnv("AI_MAX_RETRIES", 3)
	}

	return nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getIntEnv(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getDurationEnv(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}

package logger

import (
	"fmt"
	"os"

	"github.com/nusa/backend/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger
}

func New(cfg *config.Config) (*Logger, error) {
	var zapConfig zap.Config

	if cfg.Server.Environment == "production" {
		zapConfig = zap.NewProductionConfig()
	} else {
		zapConfig = zap.NewDevelopmentConfig()
	}

	zapConfig.EncoderConfig.TimeKey = "timestamp"
	zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	zapConfig.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	zapConfig.OutputPaths = []string{"stdout"}
	zapConfig.ErrorOutputPaths = []string{"stderr"}

	logger, err := zapConfig.Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build logger: %w", err)
	}

	return &Logger{Logger: logger}, nil
}

func (l *Logger) WithRequestID(requestID string) *Logger {
	return &Logger{Logger: l.Logger.With(zap.String("request_id", requestID))}
}

func (l *Logger) WithCorrelationID(correlationID string) *Logger {
	return &Logger{Logger: l.Logger.With(zap.String("correlation_id", correlationID))}
}

func (l *Logger) WithUserID(userID string) *Logger {
	return &Logger{Logger: l.Logger.With(zap.String("user_id", userID))}
}

func (l *Logger) WithSchoolID(schoolID string) *Logger {
	return &Logger{Logger: l.Logger.With(zap.String("school_id", schoolID))}
}

func (l *Logger) WithField(key string, value interface{}) *Logger {
	return &Logger{Logger: l.Logger.With(zap.Any(key, value))}
}

func (l *Logger) WithFields(fields map[string]interface{}) *Logger {
	zapFields := make([]zap.Field, 0, len(fields))
	for k, v := range fields {
		zapFields = append(zapFields, zap.Any(k, v))
	}
	return &Logger{Logger: l.Logger.With(zapFields...)}
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.Logger.Info(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.Logger.Error(msg, fields...)
}

func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.Logger.Debug(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...zap.Field) {
	l.Logger.Warn(msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...zap.Field) {
	l.Logger.Fatal(msg, fields...)
	os.Exit(1)
}

func (l *Logger) Sync() error {
	return l.Logger.Sync()
}

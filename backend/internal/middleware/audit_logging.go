package middleware

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// AuditLogger handles audit logging to the database
type AuditLogger struct {
	db *sql.DB
}

// NewAuditLogger creates a new audit logger
func NewAuditLogger(db *sql.DB) *AuditLogger {
	return &AuditLogger{db: db}
}

// AuditLog represents an audit log entry
type AuditLog struct {
	ID         uuid.UUID      `json:"id"`
	UserID     *uuid.UUID     `json:"user_id,omitempty"`
	Action     string         `json:"action"`
	EntityType string         `json:"entity_type"`
	EntityID   *uuid.UUID     `json:"entity_id,omitempty"`
	OldValues  map[string]any `json:"old_values,omitempty"`
	NewValues  map[string]any `json:"new_values,omitempty"`
	IPAddress  string         `json:"ip_address"`
	UserAgent  string         `json:"user_agent"`
	RequestID  string         `json:"request_id"`
	CreatedAt  time.Time      `json:"created_at"`
}

// LogAudit logs an audit entry to the database
func (al *AuditLogger) LogAudit(ctx context.Context, log *AuditLog) error {
	query := `
		INSERT INTO audit_logs (id, user_id, action, entity_type, entity_id, old_values, new_values, ip_address, user_agent, request_id, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`

	oldValuesJSON, _ := json.Marshal(log.OldValues)
	newValuesJSON, _ := json.Marshal(log.NewValues)

	_, err := al.db.ExecContext(ctx, query,
		log.ID,
		log.UserID,
		log.Action,
		log.EntityType,
		log.EntityID,
		oldValuesJSON,
		newValuesJSON,
		log.IPAddress,
		log.UserAgent,
		log.RequestID,
		log.CreatedAt,
	)

	return err
}

// Global audit logger instance
var globalAuditLogger *AuditLogger

// InitAuditLogger initializes the global audit logger
func InitAuditLogger(db *sql.DB) {
	globalAuditLogger = NewAuditLogger(db)
}

// AuditLogging middleware logs all requests to the audit log
func AuditLogging() gin.HandlerFunc {
	if globalAuditLogger == nil {
		// If audit logger is not initialized, just pass through
		return func(c *gin.Context) {
			c.Next()
		}
	}

	return func(c *gin.Context) {
		// Skip audit logging for health checks and public endpoints
		if c.Request.URL.Path == "/health" || c.Request.URL.Path == "/api/v1/public/health" {
			c.Next()
			return
		}

		// Read request body for logging
		var requestBody []byte
		if c.Request.Body != nil {
			requestBody, _ = io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		// Get auth context
		authCtx := GetAuthContext(c)
		var userID *uuid.UUID
		if authCtx != nil {
			uid, _ := uuid.Parse(authCtx.UserID)
			userID = &uid
		}

		// Get request ID
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = c.GetString("request_id")
		}

		// Create audit log entry
		auditLog := &AuditLog{
			ID:         uuid.New(),
			UserID:     userID,
			Action:     c.Request.Method + " " + c.Request.URL.Path,
			EntityType: "http_request",
			IPAddress:  c.ClientIP(),
			UserAgent:  c.Request.UserAgent(),
			RequestID:  requestID,
			CreatedAt:  time.Now(),
		}

		// Add request body to new values if present
		if len(requestBody) > 0 {
			var bodyMap map[string]any
			if err := json.Unmarshal(requestBody, &bodyMap); err == nil {
				auditLog.NewValues = bodyMap
			}
		}

		// Log the audit entry asynchronously
		go func() {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			globalAuditLogger.LogAudit(ctx, auditLog)
		}()

		c.Next()
	}
}

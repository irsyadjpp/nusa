# 27_BACKEND_FOUNDATION_DESIGN.md

## Foundation Document for NUSA Education Platform

**Version**: 1.0
**Date**: June 2026
**Status**: FOUNDATION DOCUMENT
**Alignment**: Aligned with 06_APPLICATION_ARCHITECTURE.md, 14_DATABASE_SCHEMA.md, 13_API_CONTRACT.md, 26_AI_ORCHESTRATION_ARCHITECTURE.md

**Purpose**: Translate the NUSA architecture into implementation-ready backend design using Go, Gin, PostgreSQL, RabbitMQ, JWT, and Docker. This document provides the foundation for backend developers to implement the system consistently.

---

# SECTION 1 — Executive Summary

## Why Backend Foundation Design Matters

A well-defined backend foundation ensures:
- Consistent code structure across the codebase
- Clear separation of concerns and dependency management
- Standardized patterns for common operations
- Maintainable and testable code
- Scalable architecture that supports MVP and future enhancements

## Technology Stack

| Component | Technology | Purpose |
|-----------|------------|---------|
| Language | Go 1.21+ | Backend implementation language |
| Web Framework | Gin | HTTP routing and middleware |
| Database | PostgreSQL 18+ | Primary data store |
| Message Queue | RabbitMQ 3.12+ | Asynchronous AI generation tasks |
| Authentication | JWT | Token-based authentication |
| Containerization | Docker | Deployment and development environment |
| ORM | sqlc | Type-safe SQL code generation |
| Migration | golang-migrate | Database schema versioning |

## Core Principles

- **Clean Architecture**: Clear separation between handlers, services, and repositories
- **Dependency Injection**: Explicit dependencies for testability
- **Interface-Based Design**: Interfaces for services and repositories
- **Transaction Safety**: Database transactions at service layer
- **Error Consistency**: Standardized error types and handling
- **Structured Logging**: Consistent logging format and levels
- **Configuration Management**: Environment-based configuration
- **Migration-First**: Database changes tracked via migrations

---

# SECTION 2 — Package Structure

## Root Package Layout

```
nusa-backend/
├── cmd/
│   └── server/
│       └── main.go                 # Application entry point
├── internal/
│   ├── config/                     # Configuration loading
│   │   ├── config.go
│   │   └── env.go
│   ├── domain/                     # Domain models and interfaces
│   │   ├── curriculum/
│   │   │   ├── model.go
│   │   │   └── repository.go
│   │   ├── learning/
│   │   │   ├── model.go
│   │   │   └── repository.go
│   │   ├── assessment/
│   │   │   ├── model.go
│   │   │   └── repository.go
│   │   ├── ai/
│   │   │   ├── model.go
│   │   │   └── repository.go
│   │   ├── user/
│   │   │   ├── model.go
│   │   │   └── repository.go
│   │   └── workflow/
│   │       ├── model.go
│   │       └── repository.go
│   ├── handler/                    # HTTP handlers
│   │   ├── curriculum/
│   │   │   └── handler.go
│   │   ├── learning/
│   │   │   └── handler.go
│   │   ├── assessment/
│   │   │   └── handler.go
│   │   ├── ai/
│   │   │   └── handler.go
│   │   ├── user/
│   │   │   └── handler.go
│   │   └── workflow/
│   │       └── handler.go
│   ├── service/                    # Business logic
│   │   ├── curriculum/
│   │   │   └── service.go
│   │   ├── learning/
│   │   │   └── service.go
│   │   ├── assessment/
│   │   │   └── service.go
│   │   ├── ai/
│   │   │   └── service.go
│   │   ├── user/
│   │   │   └── service.go
│   │   └── workflow/
│   │       └── service.go
│   ├── repository/                 # Data access layer
│   │   ├── curriculum/
│   │   │   └── repository.go
│   │   ├── learning/
│   │   │   └── repository.go
│   │   ├── assessment/
│   │   │   └── repository.go
│   │   ├── ai/
│   │   │   └── repository.go
│   │   ├── user/
│   │   │   └── repository.go
│   │   └── workflow/
│   │       └── repository.go
│   ├── middleware/                 # HTTP middleware
│   │   ├── auth.go
│   │   ├── logging.go
│   │   ├── cors.go
│   │   ├── recovery.go
│   │   └── request_id.go
│   ├── ai/                        # AI orchestration
│   │   ├── gateway.go
│   │   ├── provider.go
│   │   ├── prompt_builder.go
│   │   └── response_processor.go
│   ├── queue/                     # RabbitMQ integration
│   │   ├── publisher.go
│   │   ├── consumer.go
│   │   └── job.go
│   ├── db/                        # Database setup
│   │   ├── db.go
│   │   ├── pool.go
│   │   └── transaction.go
│   ├── error/                     # Error handling
│   │   ├── error.go
│   │   └── handler.go
│   ├── logger/                    # Logging
│   │   └── logger.go
│   └── util/                      # Utilities
│       ├── jwt.go
│       ├── validator.go
│       └── time.go
├── migrations/                     # Database migrations
│   ├── 000001_init_schema.up.sql
│   └── 000001_init_schema.down.sql
├── queries/                       # sqlc generated queries
│   ├── curriculum.sql
│   ├── learning.sql
│   ├── assessment.sql
│   ├── ai.sql
│   ├── user.sql
│   └── workflow.sql
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── go.sum
```

## Package Responsibilities

### cmd/server
- Application entry point
- Dependency injection setup
- Server initialization
- Graceful shutdown

### internal/config
- Configuration loading from environment variables
- Configuration validation
- Configuration struct definitions

### internal/domain
- Domain model definitions
- Repository interfaces
- Domain-specific types and constants
- No business logic

### internal/handler
- HTTP request handling
- Request validation
- Response formatting
- Error translation to HTTP responses

### internal/service
- Business logic implementation
- Transaction management
- Orchestration of multiple repositories
- AI orchestration calls

### internal/repository
- Database access implementation
- SQL query execution
- Data mapping to domain models
- Transaction context handling

### internal/middleware
- HTTP middleware chain
- Authentication and authorization
- Request logging
- CORS handling
- Panic recovery

### internal/ai
- AI provider integration
- Prompt building
- Response processing
- AI generation orchestration

### internal/queue
- RabbitMQ integration
- Job publishing
- Job consumption
- Worker pool management

### internal/db
- Database connection pool setup
- Transaction management utilities
- Database health checks

### internal/error
- Error type definitions
- Error handling utilities
- HTTP error response formatting

### internal/logger
- Structured logging setup
- Log level management
- Context-aware logging

### internal/util
- JWT token generation and validation
- Input validation utilities
- Time formatting utilities

### migrations
- Database schema migrations
- Migration version tracking
- Rollback scripts

### queries
- SQL queries for sqlc code generation
- Type-safe query definitions
- Query result mapping

---

# SECTION 3 — Dependency Flow

## Dependency Diagram

```
┌─────────────────────────────────────────────────────────────┐
│                      cmd/server/main.go                      │
│                    (Application Entry Point)                  │
└────────────────────┬────────────────────────────────────────┘
                     │
                     ↓ (creates and injects)
┌─────────────────────────────────────────────────────────────┐
│                    internal/handler/*                        │
│                  (HTTP Request Handlers)                     │
└────────────────────┬────────────────────────────────────────┘
                     │
                     ↓ (calls)
┌─────────────────────────────────────────────────────────────┐
│                    internal/service/*                         │
│                   (Business Logic Layer)                      │
└────┬──────────────────────────────────┬─────────────────────┘
     │                                  │
     ↓ (uses)                          ↓ (uses)
┌──────────────────┐           ┌──────────────────┐
│ internal/ai/*   │           │ internal/repository/* │
│ (AI Orchestration)│           │   (Data Access)    │
└──────────────────┘           └──────┬─────────────┘
                                        │
                                        ↓ (uses)
                              ┌──────────────────┐
                              │  internal/db/*   │
                              │ (Database Layer)  │
                              └──────────────────┘
```

## Dependency Rules

1. **Handler → Service**: Handlers depend on service interfaces
2. **Service → Repository**: Services depend on repository interfaces
3. **Service → AI**: Services depend on AI gateway interface
4. **Repository → DB**: Repositories depend on database utilities
5. **No Circular Dependencies**: Dependencies flow in one direction
6. **Interface-Based**: Services and repositories use interfaces
7. **Domain Independence**: Domain layer has no dependencies on other layers

## Dependency Injection

```go
// cmd/server/main.go
func main() {
    // Load configuration
    cfg := config.Load()
    
    // Initialize database
    db := db.NewPool(cfg.Database)
    
    // Initialize repositories
    curriculumRepo := repository.NewCurriculumRepository(db)
    learningRepo := repository.NewLearningRepository(db)
    assessmentRepo := repository.NewAssessmentRepository(db)
    aiRepo := repository.NewAIRepository(db)
    userRepo := repository.NewUserRepository(db)
    workflowRepo := repository.NewWorkflowRepository(db)
    
    // Initialize AI gateway
    aiGateway := ai.NewGateway(cfg.AI)
    
    // Initialize queue
    queue := queue.New(cfg.RabbitMQ)
    
    // Initialize services
    curriculumService := service.NewCurriculumService(curriculumRepo)
    learningService := service.NewLearningService(learningRepo, aiGateway, queue)
    assessmentService := service.NewAssessmentService(assessmentRepo)
    aiService := service.NewAIService(aiRepo, aiGateway)
    userService := service.NewUserService(userRepo)
    workflowService := service.NewWorkflowService(workflowRepo)
    
    // Initialize handlers
    curriculumHandler := handler.NewCurriculumHandler(curriculumService)
    learningHandler := handler.NewLearningHandler(learningService)
    assessmentHandler := handler.NewAssessmentHandler(assessmentService)
    aiHandler := handler.NewAIHandler(aiService)
    userHandler := handler.NewUserHandler(userService)
    workflowHandler := handler.NewWorkflowHandler(workflowService)
    
    // Setup router
    router := setupRouter(cfg, handlers...)
    
    // Start server
    server := &http.Server{
        Addr:    cfg.Server.Port,
        Handler: router,
    }
    
    // Graceful shutdown
    gracefulShutdown(server, db, queue)
}
```

---

# SECTION 4 — Handler-Service-Repository Pattern

## Pattern Overview

The backend follows the classic three-layer architecture:

1. **Handler Layer**: HTTP request/response handling
2. **Service Layer**: Business logic and orchestration
3. **Repository Layer**: Data access and persistence

## Handler Layer

### Responsibilities
- Parse HTTP requests
- Validate request payloads
- Call service layer
- Format HTTP responses
- Translate service errors to HTTP status codes

### Example Handler

```go
package curriculum

import (
    "net/http"
    
    "github.com/gin-gonic/gin"
)

type Handler struct {
    service Service
}

func NewHandler(service Service) *Handler {
    return &Handler{service: service}
}

// GenerateTPSet generates a TP Set for a CP
func (h *Handler) GenerateTPSet(c *gin.Context) {
    var req GenerateTPSetRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    // Get user from context (set by auth middleware)
    userID := c.GetString("user_id")
    
    // Call service
    result, err := h.service.GenerateTPSet(c.Request.Context(), userID, req)
    if err != nil {
        handleError(c, err)
        return
    }
    
    c.JSON(http.StatusCreated, result)
}

// GetTPSet retrieves a TP Set by ID
func (h *Handler) GetTPSet(c *gin.Context) {
    id := c.Param("id")
    
    result, err := h.service.GetTPSet(c.Request.Context(), id)
    if err != nil {
        handleError(c, err)
        return
    }
    
    c.JSON(http.StatusOK, result)
}
```

### Request/Response DTOs

```go
type GenerateTPSetRequest struct {
    CPID           string                 `json:"cp_id" binding:"required,uuid"`
    Preferences    map[string]interface{} `json:"preferences"`
}

type TPSetResponse struct {
    ID              string       `json:"id"`
    CPID            string       `json:"cp_id"`
    VersionNo       int          `json:"version_no"`
    Status          string       `json:"status"`
    GenerationSource string      `json:"generation_source"`
    GeneratedBy     string       `json:"generated_by"`
    AIGenerationID  *string      `json:"ai_generation_id,omitempty"`
    TPs             []TPResponse `json:"tps"`
    CreatedAt       time.Time    `json:"created_at"`
    UpdatedAt       time.Time    `json:"updated_at"`
}
```

## Service Layer

### Responsibilities
- Implement business logic
- Orchestrate multiple repository calls
- Manage database transactions
- Call AI gateway for AI operations
- Enforce business rules and validations

### Example Service

```go
package curriculum

import (
    "context"
    "errors"
    
    "github.com/google/uuid"
)

type Service interface {
    GenerateTPSet(ctx context.Context, userID string, req GenerateTPSetRequest) (*TPSet, error)
    GetTPSet(ctx context.Context, id string) (*TPSet, error)
    ApproveTPSet(ctx context.Context, id string, userID string) error
    ArchiveTPSet(ctx context.Context, id string, userID string) error
}

type service struct {
    repo    Repository
    aiGateway ai.Gateway
    queue   queue.Publisher
}

func NewService(repo Repository, aiGateway ai.Gateway, queue queue.Publisher) Service {
    return &service{
        repo:      repo,
        aiGateway: aiGateway,
        queue:     queue,
    }
}

func (s *service) GenerateTPSet(ctx context.Context, userID string, req GenerateTPSetRequest) (*TPSet, error) {
    // Validate CP exists
    cp, err := s.repo.GetCP(ctx, req.CPID)
    if err != nil {
        return nil, err
    }
    
    // Build AI prompt
    prompt := s.buildPrompt(cp, req.Preferences)
    
    // Call AI gateway (synchronous for MVP)
    aiResponse, err := s.aiGateway.Generate(ctx, prompt)
    if err != nil {
        return nil, err
    }
    
    // Start transaction
    tx, err := s.repo.BeginTx(ctx)
    if err != nil {
        return nil, err
    }
    defer tx.Rollback()
    
    // Create AI generation log
    genLog := &AIGenerationLog{
        ID:           uuid.New().String(),
        UserID:       userID,
        SchoolID:     cp.SchoolID,
        ArtifactType: "tp_set",
        Provider:     aiResponse.Provider,
        Model:        aiResponse.Model,
        TokensUsed:   aiResponse.TokensUsed,
        EstimatedCost: aiResponse.EstimatedCost,
        ResponseTimeMs: aiResponse.ResponseTimeMs,
        Status:       "COMPLETED",
        PromptSnapshot: prompt,
        ResponseSnapshot: aiResponse.RawResponse,
    }
    
    if err := s.repo.CreateAIGenerationLog(ctx, genLog); err != nil {
        return nil, err
    }
    
    // Create TP Set
    tpSet := &TPSet{
        ID:              uuid.New().String(),
        CPID:            req.CPID,
        VersionNo:       1,
        Status:          "DRAFT",
        GenerationSource: "AI_GENERATED",
        GeneratedBy:     userID,
        AIGenerationID:  &genLog.ID,
    }
    
    if err := s.repo.CreateTPSet(ctx, tpSet); err != nil {
        return nil, err
    }
    
    // Create TP Items
    for i, tpData := range aiResponse.TPs {
        tp := &TP{
            ID:           uuid.New().String(),
            TPSetID:      tpSet.ID,
            SequenceNumber: i + 1,
            Title:        tpData.Title,
            LearningObjectives: tpData.LearningObjectives,
            EstimatedWeeks: tpData.EstimatedWeeks,
            Status:       "DRAFT",
        }
        
        if err := s.repo.CreateTP(ctx, tp); err != nil {
            return nil, err
        }
    }
    
    // Commit transaction
    if err := tx.Commit(); err != nil {
        return nil, err
    }
    
    return tpSet, nil
}

func (s *service) ApproveTPSet(ctx context.Context, id string, userID string) error {
    // Get TP Set
    tpSet, err := s.repo.GetTPSet(ctx, id)
    if err != nil {
        return err
    }
    
    // Validate status
    if tpSet.Status != "DRAFT" && tpSet.Status != "UNDER_REVIEW" {
        return errors.New("invalid status for approval")
    }
    
    // Update TP Set
    tpSet.Status = "APPROVED"
    tpSet.ApprovedBy = &userID
    now := time.Now()
    tpSet.ApprovedAt = &now
    
    return s.repo.UpdateTPSet(ctx, tpSet)
}
```

## Repository Layer

### Responsibilities
- Execute database queries
- Map database rows to domain models
- Handle transaction contexts
- Provide data access methods

### Example Repository

```go
package curriculum

import (
    "context"
    "database/sql"
    
    "github.com/jmoiron/sqlx"
)

type Repository interface {
    BeginTx(ctx context.Context) (Transaction, error)
    GetCP(ctx context.Context, id string) (*CP, error)
    CreateTPSet(ctx context.Context, tpSet *TPSet) error
    GetTPSet(ctx context.Context, id string) (*TPSet, error)
    UpdateTPSet(ctx context.Context, tpSet *TPSet) error
    CreateTP(ctx context.Context, tp *TP) error
    CreateAIGenerationLog(ctx context.Context, log *AIGenerationLog) error
}

type repository struct {
    db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
    return &repository{db: db}
}

func (r *repository) BeginTx(ctx context.Context) (Transaction, error) {
    tx, err := r.db.BeginTxx(ctx, nil)
    if err != nil {
        return nil, err
    }
    return &transaction{tx: tx}, nil
}

func (r *repository) GetCP(ctx context.Context, id string) (*CP, error) {
    var cp CP
    err := r.db.GetContext(ctx, &cp, "SELECT * FROM cp WHERE id = $1", id)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, ErrNotFound
        }
        return nil, err
    }
    return &cp, nil
}

func (r *repository) CreateTPSet(ctx context.Context, tpSet *TPSet) error {
    query := `
        INSERT INTO tp_sets (id, cp_id, version_no, status, generation_source, generated_by, ai_generation_id)
        VALUES (:id, :cp_id, :version_no, :status, :generation_source, :generated_by, :ai_generation_id)
    `
    _, err := r.db.NamedExecContext(ctx, query, tpSet)
    return err
}

func (r *repository) UpdateTPSet(ctx context.Context, tpSet *TPSet) error {
    query := `
        UPDATE tp_sets
        SET status = :status, approved_by = :approved_by, approved_at = :approved_at, updated_at = NOW()
        WHERE id = :id
    `
    _, err := r.db.NamedExecContext(ctx, query, tpSet)
    return err
}

// Transaction implementation
type transaction struct {
    tx *sqlx.Tx
}

func (t *transaction) Commit() error {
    return t.tx.Commit()
}

func (t *transaction) Rollback() error {
    return t.tx.Rollback()
}
```

---

# SECTION 5 — Configuration Loading

## Configuration Structure

Configuration is loaded from environment variables with sensible defaults.

### Configuration File

```go
package config

import (
    "os"
    "strconv"
    "time"
)

type Config struct {
    Server   ServerConfig
    Database DatabaseConfig
    AI       AIConfig
    RabbitMQ RabbitMQConfig
    JWT      JWTConfig
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
}

type AIConfig struct {
    PrimaryProvider   string
    FallbackProvider  string
    OpenAIKey         string
    GeminiKey         string
    Timeout           time.Duration
    MaxRetries        int
}

type RabbitMQConfig struct {
    Host     string
    Port     string
    User     string
    Password string
    Queue    string
}

type JWTConfig struct {
    Secret     string
    Expiration time.Duration
}

func Load() *Config {
    return &Config{
        Server: ServerConfig{
            Port:         getEnv("SERVER_PORT", ":8080"),
            ReadTimeout:  getDurationEnv("SERVER_READ_TIMEOUT", 15*time.Second),
            WriteTimeout: getDurationEnv("SERVER_WRITE_TIMEOUT", 15*time.Second),
            Environment:  getEnv("ENVIRONMENT", "development"),
        },
        Database: DatabaseConfig{
            Host:            getEnv("DB_HOST", "localhost"),
            Port:            getEnv("DB_PORT", "5432"),
            User:            getEnv("DB_USER", "nusa"),
            Password:        getEnv("DB_PASSWORD", ""),
            DBName:          getEnv("DB_NAME", "nusa"),
            SSLMode:         getEnv("DB_SSLMODE", "disable"),
            MaxOpenConns:    getIntEnv("DB_MAX_OPEN_CONNS", 25),
            MaxIdleConns:    getIntEnv("DB_MAX_IDLE_CONNS", 5),
            ConnMaxLifetime: getDurationEnv("DB_CONN_MAX_LIFETIME", 5*time.Minute),
        },
        AI: AIConfig{
            PrimaryProvider:  getEnv("AI_PRIMARY_PROVIDER", "openai"),
            FallbackProvider: getEnv("AI_FALLBACK_PROVIDER", "gemini"),
            OpenAIKey:        getEnv("AI_OPENAI_KEY", ""),
            GeminiKey:        getEnv("AI_GEMINI_KEY", ""),
            Timeout:          getDurationEnv("AI_TIMEOUT", 30*time.Second),
            MaxRetries:       getIntEnv("AI_MAX_RETRIES", 3),
        },
        RabbitMQ: RabbitMQConfig{
            Host:     getEnv("RABBITMQ_HOST", "localhost"),
            Port:     getEnv("RABBITMQ_PORT", "5672"),
            User:     getEnv("RABBITMQ_USER", "guest"),
            Password: getEnv("RABBITMQ_PASSWORD", "guest"),
            Queue:    getEnv("RABBITMQ_QUEUE", "ai_generation"),
        },
        JWT: JWTConfig{
            Secret:     getEnv("JWT_SECRET", "change-me-in-production"),
            Expiration: getDurationEnv("JWT_EXPIRATION", 24*time.Hour),
        },
    }
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
```

## Environment Variables

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| SERVER_PORT | HTTP server port | :8080 | No |
| ENVIRONMENT | Environment (development/production) | development | No |
| DB_HOST | Database host | localhost | Yes |
| DB_PORT | Database port | 5432 | No |
| DB_USER | Database user | nusa | Yes |
| DB_PASSWORD | Database password | - | Yes |
| DB_NAME | Database name | nusa | Yes |
| DB_SSLMODE | Database SSL mode | disable | No |
| AI_PRIMARY_PROVIDER | Primary AI provider | openai | No |
| AI_FALLBACK_PROVIDER | Fallback AI provider | gemini | No |
| AI_OPENAI_KEY | OpenAI API key | - | Yes |
| AI_GEMINI_KEY | Gemini API key | - | Yes |
| RABBITMQ_HOST | RabbitMQ host | localhost | Yes |
| RABBITMQ_PORT | RabbitMQ port | 5672 | No |
| RABBITMQ_USER | RabbitMQ user | guest | Yes |
| RABBITMQ_PASSWORD | RabbitMQ password | - | Yes |
| JWT_SECRET | JWT signing secret | - | Yes |

---

# SECTION 6 — Middleware Strategy

## Middleware Chain

Middleware is applied in the following order:

1. **Request ID** - Generate unique request ID
2. **Logger** - Log incoming request
3. **CORS** - Handle CORS headers
4. **Recovery** - Recover from panics
5. **Authentication** - Validate JWT token
6. **Authorization** - Check permissions (if needed)

## Middleware Implementation

### Request ID Middleware

```go
package middleware

import (
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

func RequestID() gin.HandlerFunc {
    return func(c *gin.Context) {
        requestID := c.GetHeader("X-Request-ID")
        if requestID == "" {
            requestID = uuid.New().String()
        }
        c.Set("request_id", requestID)
        c.Header("X-Request-ID", requestID)
        c.Next()
    }
}
```

### Logging Middleware

```go
package middleware

import (
    "time"
    
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

func Logging(logger *zap.Logger) gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        path := c.Request.URL.Path
        query := c.Request.URL.RawQuery
        
        c.Next()
        
        latency := time.Since(start)
        requestID := c.GetString("request_id")
        
        logger.Info("HTTP Request",
            zap.String("request_id", requestID),
            zap.String("method", c.Request.Method),
            zap.String("path", path),
            zap.String("query", query),
            zap.Int("status", c.Writer.Status()),
            zap.Duration("latency", latency),
            zap.String("ip", c.ClientIP()),
            zap.String("user_agent", c.Request.UserAgent()),
        )
    }
}
```

### CORS Middleware

```go
package middleware

import (
    "github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
        
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        
        c.Next()
    }
}
```

### Recovery Middleware

```go
package middleware

import (
    "net/http"
    
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

func Recovery(logger *zap.Logger) gin.HandlerFunc {
    return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
        requestID := c.GetString("request_id")
        
        logger.Error("Panic recovered",
            zap.String("request_id", requestID),
            zap.Any("panic", recovered),
            zap.String("path", c.Request.URL.Path),
        )
        
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Internal server error",
        })
    })
}
```

### Authentication Middleware

```go
package middleware

import (
    "net/http"
    "strings"
    
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
)

func Auth(jwtSecret string) gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
            c.Abort()
            return
        }
        
        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte(jwtSecret), nil
        })
        
        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }
        
        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
            c.Abort()
            return
        }
        
        c.Set("user_id", claims["sub"])
        c.Set("school_id", claims["school_id"])
        c.Set("role", claims["role"])
        
        c.Next()
    }
}
```

### Middleware Setup

```go
func setupRouter(cfg *config.Config, handlers ...Handler) *gin.Engine {
    router := gin.New()
    
    // Initialize logger
    logger, err := zap.NewProduction()
    if err != nil {
        panic(err)
    }
    
    // Apply global middleware
    router.Use(middleware.RequestID())
    router.Use(middleware.Logging(logger))
    router.Use(middleware.CORS())
    router.Use(middleware.Recovery(logger))
    
    // Public routes
    public := router.Group("/api/v1/public")
    {
        public.POST("/login", userHandler.Login)
    }
    
    // Protected routes
    protected := router.Group("/api/v1")
    protected.Use(middleware.Auth(cfg.JWT.Secret))
    {
        // Curriculum routes
        curriculum := protected.Group("/curriculum")
        {
            curriculum.GET("/cp", curriculumHandler.ListCPs)
            curriculum.GET("/cp/:id", curriculumHandler.GetCP)
            curriculum.POST("/cp/:id/tp-sets/generate", curriculumHandler.GenerateTPSet)
            curriculum.GET("/cp/:id/tp-sets", curriculumHandler.ListTPSets)
            curriculum.GET("/tp-sets/:id", curriculumHandler.GetTPSet)
            curriculum.POST("/tp-sets/:id/approve", curriculumHandler.ApproveTPSet)
            curriculum.POST("/tp-sets/:id/archive", curriculumHandler.ArchiveTPSet)
        }
        
        // Learning routes
        learning := protected.Group("/learning")
        {
            learning.POST("/tp-sets/:id/atp-sets/generate", learningHandler.GenerateATPSet)
            learning.GET("/tp-sets/:id/atp-sets", learningHandler.ListATPSets)
            learning.GET("/atp-sets/:id", learningHandler.GetATPSet)
        }
        
        // AI routes
        ai := protected.Group("/ai")
        {
            ai.POST("/generate", aiHandler.Generate)
            ai.GET("/logs/:id", aiHandler.GetGenerationLog)
        }
        
        // Workflow routes
        workflow := protected.Group("/workflow")
        {
            workflow.GET("/history/:artifact_id", workflowHandler.GetHistory)
        }
    }
    
    return router
}
```

---

# SECTION 7 — Transaction Strategy

## Transaction Management

Transactions are managed at the service layer to ensure atomicity of business operations.

### Transaction Interface

```go
package db

import "context"

type Transaction interface {
    Commit() error
    Rollback() error
}

type TransactionManager interface {
    BeginTx(ctx context.Context) (Transaction, error)
}
```

### Transaction Usage in Service

```go
func (s *service) GenerateTPSet(ctx context.Context, userID string, req GenerateTPSetRequest) (*TPSet, error) {
    // Start transaction
    tx, err := s.repo.BeginTx(ctx)
    if err != nil {
        return nil, err
    }
    defer tx.Rollback()
    
    // Perform multiple operations within transaction
    if err := s.repo.CreateAIGenerationLog(ctx, genLog); err != nil {
        return nil, err
    }
    
    if err := s.repo.CreateTPSet(ctx, tpSet); err != nil {
        return nil, err
    }
    
    for _, tp := range tps {
        if err := s.repo.CreateTP(ctx, tp); err != nil {
            return nil, err
        }
    }
    
    // Commit transaction
    if err := tx.Commit(); err != nil {
        return nil, err
    }
    
    return tpSet, nil
}
```

### Transaction Context

Repositories accept a context parameter that may contain a transaction:

```go
func (r *repository) CreateTPSet(ctx context.Context, tpSet *TPSet) error {
    // Check if transaction exists in context
    tx, ok := ctx.Value("tx").(*sqlx.Tx)
    if ok {
        // Use transaction
        _, err := tx.NamedExecContext(ctx, query, tpSet)
        return err
    }
    
    // Use regular connection
    _, err := r.db.NamedExecContext(ctx, query, tpSet)
    return err
}
```

### Transaction Helper

```go
package db

import (
    "context"
    "github.com/jmoiron/sqlx"
)

func WithTransaction(ctx context.Context, db *sqlx.DB, fn func(ctx context.Context) error) error {
    tx, err := db.BeginTxx(ctx, nil)
    if err != nil {
        return err
    }
    
    defer func() {
        if p := recover(); p != nil {
            _ = tx.Rollback()
            panic(p)
        }
    }()
    
    ctx = context.WithValue(ctx, "tx", tx)
    
    if err := fn(ctx); err != nil {
        if rbErr := tx.Rollback(); rbErr != nil {
            return fmt.Errorf("rollback failed: %v, original error: %w", rbErr, err)
        }
        return err
    }
    
    return tx.Commit()
}
```

---

# SECTION 8 — Error Handling Strategy

## Error Types

Define custom error types for different error scenarios:

```go
package error

import "net/http"

type AppError struct {
    Code       int
    Message    string
    Internal   error
    Context    map[string]interface{}
}

func (e *AppError) Error() string {
    return e.Message
}

func (e *AppError) Unwrap() error {
    return e.Internal
}

// Common error constructors
func NewBadRequestError(message string, internal error) *AppError {
    return &AppError{
        Code:     http.StatusBadRequest,
        Message:  message,
        Internal: internal,
    }
}

func NewNotFoundError(message string, internal error) *AppError {
    return &AppError{
        Code:     http.StatusNotFound,
        Message:  message,
        Internal: internal,
    }
}

func NewUnauthorizedError(message string, internal error) *AppError {
    return &AppError{
        Code:     http.StatusUnauthorized,
        Message:  message,
        Internal: internal,
    }
}

func NewForbiddenError(message string, internal error) *AppError {
    return &AppError{
        Code:     http.StatusForbidden,
        Message:  message,
        Internal: internal,
    }
}

func NewInternalServerError(message string, internal error) *AppError {
    return &AppError{
        Code:     http.StatusInternalServerError,
        Message:  message,
        Internal: internal,
    }
}

func NewValidationError(message string, context map[string]interface{}) *AppError {
    return &AppError{
        Code:    http.StatusBadRequest,
        Message: message,
        Context: context,
    }
}
```

## Error Handler

```go
package error

import (
    "net/http"
    
    "github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
    if appErr, ok := err.(*AppError); ok {
        c.JSON(appErr.Code, gin.H{
            "error":   appErr.Message,
            "context": appErr.Context,
        })
        return
    }
    
    // Default to internal server error
    c.JSON(http.StatusInternalServerError, gin.H{
        "error": "Internal server error",
    })
}
```

## Repository Errors

```go
package repository

import "github.com/nusa-backend/internal/error"

var (
    ErrNotFound = error.NewNotFoundError("Resource not found", nil)
    ErrDuplicate = error.NewBadRequestError("Resource already exists", nil)
    ErrInvalidInput = error.NewBadRequestError("Invalid input", nil)
)
```

## Service Errors

```go
package service

import "github.com/nusa-backend/internal/error"

var (
    ErrInvalidStatus = error.NewBadRequestError("Invalid status transition", nil)
    ErrPermissionDenied = error.NewForbiddenError("Permission denied", nil)
    ErrAIGenerationFailed = error.NewInternalServerError("AI generation failed", nil)
)
```

---

# SECTION 9 — Logging Strategy

## Logger Setup

Use structured logging with zap:

```go
package logger

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

func New(environment string) (*zap.Logger, error) {
    var config zap.Config
    
    if environment == "production" {
        config = zap.NewProductionConfig()
        config.EncoderConfig.TimeKey = "timestamp"
        config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
    } else {
        config = zap.NewDevelopmentConfig()
    }
    
    return config.Build()
}
```

## Logging Levels

- **Debug**: Detailed information for debugging
- **Info**: General informational messages
- **Warn**: Warning messages for potentially harmful situations
- **Error**: Error messages for errors that occur
- **Fatal**: Fatal errors that require immediate attention

## Logging Context

Include relevant context in log entries:

```go
logger.Info("TP Set generated",
    zap.String("request_id", requestID),
    zap.String("user_id", userID),
    zap.String("tp_set_id", tpSet.ID),
    zap.String("cp_id", tpSet.CPID),
    zap.String("ai_generation_id", *tpSet.AIGenerationID),
)
```

## Structured Logging

Use structured fields instead of string formatting:

```go
// Good
logger.Info("User logged in",
    zap.String("user_id", userID),
    zap.String("ip", ip),
)

// Bad
logger.Info(fmt.Sprintf("User %s logged in from %s", userID, ip))
```

---

# SECTION 10 — Migration Strategy

## Migration Tool

Use golang-migrate for database migrations:

```bash
# Install golang-migrate
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Create migration
migrate create -ext sql -dir migrations -seq create_tp_sets_table

# Run migrations
migrate -path migrations -database "postgres://user:password@localhost:5432/nusa?sslmode=disable" up

# Rollback migration
migrate -path migrations -database "postgres://user:password@localhost:5432/nusa?sslmode=disable" down 1
```

## Migration Naming Convention

Use sequential numbering with descriptive names:

```
000001_init_schema.up.sql
000001_init_schema.down.sql
000002_add_tp_sets_table.up.sql
000002_add_tp_sets_table.down.sql
000003_add_atp_sets_table.up.sql
000003_add_atp_sets_table.down.sql
```

## Migration Example

### Up Migration

```sql
-- 000002_add_tp_sets_table.up.sql

CREATE TABLE tp_sets (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    cp_id UUID NOT NULL REFERENCES cp(id),
    version_no INTEGER NOT NULL DEFAULT 1,
    status VARCHAR(20) NOT NULL CHECK (status IN ('DRAFT', 'UNDER_REVIEW', 'APPROVED', 'REJECTED', 'ARCHIVED')),
    generation_source VARCHAR(50) NOT NULL,
    generation_reason TEXT,
    generated_by UUID NOT NULL REFERENCES users(id),
    ai_generation_id UUID,
    approved_by UUID REFERENCES users(id),
    approved_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_tp_sets_cp_id ON tp_sets(cp_id);
CREATE INDEX idx_tp_sets_version_no ON tp_sets(version_no);
CREATE INDEX idx_tp_sets_status ON tp_sets(status);
CREATE INDEX idx_tp_sets_ai_generation_id ON tp_sets(ai_generation_id);
CREATE UNIQUE INDEX idx_tp_sets_cp_version ON tp_sets(cp_id, version_no);
```

### Down Migration

```sql
-- 000002_add_tp_sets_table.down.sql

DROP INDEX IF EXISTS idx_tp_sets_cp_version;
DROP INDEX IF EXISTS idx_tp_sets_ai_generation_id;
DROP INDEX IF EXISTS idx_tp_sets_status;
DROP INDEX IF EXISTS idx_tp_sets_version_no;
DROP INDEX IF EXISTS idx_tp_sets_cp_id;
DROP TABLE IF EXISTS tp_sets;
```

## Migration in Docker Compose

```yaml
version: '3.8'
services:
  migrate:
    image: migrate/migrate
    command: -path /migrations -database postgres://nusa:password@db:5432/nusa?sslmode=disable up
    volumes:
      - ./migrations:/migrations
    depends_on:
      - db
```

---

# SECTION 11 — Startup Sequence

## Startup Flow Diagram

```
┌─────────────────┐
│  main() starts  │
└────────┬────────┘
         │
         ↓
┌─────────────────┐
│ Load Config     │
│ (env vars)      │
└────────┬────────┘
         │
         ↓
┌─────────────────┐
│ Initialize      │
│ Logger          │
└────────┬────────┘
         │
         ↓
┌─────────────────┐
│ Connect to      │
│ Database        │
└────────┬────────┘
         │
         ↓
┌─────────────────┐
│ Run Migrations  │
│ (if configured) │
└────────┬────────┘
         │
         ↓
┌─────────────────┐
│ Connect to      │
│ RabbitMQ        │
└────────┬────────┘
         │
         ↓
┌─────────────────┐
│ Initialize      │
│ Repositories    │
└────────┬────────┘
         │
         ↓
┌─────────────────┐
│ Initialize      │
│ Services        │
└────────┬────────┘
         │
         ↓
┌─────────────────┐
│ Initialize      │
│ Handlers        │
└────────┬────────┘
         │
         ↓
┌─────────────────┐
│ Setup Router    │
│ & Middleware    │
└────────┬────────┘
         │
         ↓
┌─────────────────┐
│ Start HTTP      │
│ Server          │
└────────┬────────┘
         │
         ↓
┌─────────────────┐
│ Start Queue     │
│ Consumers       │
└────────┬────────┘
         │
         ↓
┌─────────────────┐
│ Wait for        │
│ Shutdown Signal │
└────────┬────────┘
         │
         ↓
┌─────────────────┐
│ Graceful        │
│ Shutdown        │
└─────────────────┘
```

## Startup Implementation

```go
package main

import (
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
    
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

func main() {
    // Load configuration
    cfg := config.Load()
    
    // Initialize logger
    logger, err := logger.New(cfg.Server.Environment)
    if err != nil {
        log.Fatalf("Failed to initialize logger: %v", err)
    }
    defer logger.Sync()
    
    logger.Info("Starting NUSA backend server",
        zap.String("environment", cfg.Server.Environment),
        zap.String("port", cfg.Server.Port),
    )
    
    // Connect to database
    db, err := db.NewPool(cfg.Database)
    if err != nil {
        logger.Fatal("Failed to connect to database", zap.Error(err))
    }
    defer db.Close()
    
    // Run migrations (optional, can be disabled)
    if cfg.Server.Environment == "development" {
        if err := runMigrations(cfg.Database); err != nil {
            logger.Warn("Migration failed", zap.Error(err))
        }
    }
    
    // Connect to RabbitMQ
    queue, err := queue.New(cfg.RabbitMQ)
    if err != nil {
        logger.Fatal("Failed to connect to RabbitMQ", zap.Error(err))
    }
    defer queue.Close()
    
    // Initialize repositories
    curriculumRepo := repository.NewCurriculumRepository(db)
    learningRepo := repository.NewLearningRepository(db)
    assessmentRepo := repository.NewAssessmentRepository(db)
    aiRepo := repository.NewAIRepository(db)
    userRepo := repository.NewUserRepository(db)
    workflowRepo := repository.NewWorkflowRepository(db)
    
    // Initialize AI gateway
    aiGateway := ai.NewGateway(cfg.AI)
    
    // Initialize services
    curriculumService := service.NewCurriculumService(curriculumRepo)
    learningService := service.NewLearningService(learningRepo, aiGateway, queue)
    assessmentService := service.NewAssessmentService(assessmentRepo)
    aiService := service.NewAIService(aiRepo, aiGateway)
    userService := service.NewUserService(userRepo)
    workflowService := service.NewWorkflowService(workflowRepo)
    
    // Initialize handlers
    curriculumHandler := handler.NewCurriculumHandler(curriculumService)
    learningHandler := handler.NewLearningHandler(learningService)
    assessmentHandler := handler.NewAssessmentHandler(assessmentService)
    aiHandler := handler.NewAIHandler(aiService)
    userHandler := handler.NewUserHandler(userService)
    workflowHandler := handler.NewWorkflowHandler(workflowService)
    
    // Setup router
    router := setupRouter(cfg, logger, curriculumHandler, learningHandler, assessmentHandler, aiHandler, userHandler, workflowHandler)
    
    // Start HTTP server
    server := &http.Server{
        Addr:         cfg.Server.Port,
        Handler:      router,
        ReadTimeout:  cfg.Server.ReadTimeout,
        WriteTimeout: cfg.Server.WriteTimeout,
    }
    
    // Start server in goroutine
    go func() {
        logger.Info("HTTP server listening", zap.String("addr", cfg.Server.Port))
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            logger.Fatal("HTTP server failed", zap.Error(err))
        }
    }()
    
    // Start queue consumers
    go func() {
        logger.Info("Starting queue consumers")
        if err := queue.StartConsumers(); err != nil {
            logger.Error("Queue consumer failed", zap.Error(err))
        }
    }()
    
    // Wait for shutdown signal
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    
    logger.Info("Shutdown signal received")
    
    // Graceful shutdown
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    // Shutdown HTTP server
    if err := server.Shutdown(ctx); err != nil {
        logger.Error("HTTP server shutdown failed", zap.Error(err))
    }
    
    // Stop queue consumers
    queue.StopConsumers()
    
    logger.Info("Server shutdown complete")
}

func runMigrations(cfg config.DatabaseConfig) error {
    // Implementation using golang-migrate
    return nil
}
```

## Graceful Shutdown

Graceful shutdown ensures:
- In-flight requests complete
- Database connections close properly
- Queue consumers finish current jobs
- Resources are released

---

# SECTION 12 — Package Diagram

## Package Diagram

```
┌─────────────────────────────────────────────────────────────┐
│                        cmd/server                           │
│                      (Application Entry)                     │
└─────────────────────────────────────────────────────────────┘
                            │
                            ↓
┌─────────────────────────────────────────────────────────────┐
│                      internal/handler                        │
│                    (HTTP Request Layer)                       │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐      │
│  │ curriculum   │  │  learning     │  │  assessment   │      │
│  └──────────────┘  └──────────────┘  └──────────────┘      │
└─────────────────────────────────────────────────────────────┘
                            │
                            ↓
┌─────────────────────────────────────────────────────────────┐
│                       internal/service                        │
│                    (Business Logic Layer)                     │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐      │
│  │ curriculum   │  │  learning     │  │  assessment   │      │
│  └──────────────┘  └──────────────┘  └──────────────┘      │
└─────────────────────────────────────────────────────────────┘
                            │
                ┌───────────┴───────────┐
                ↓                       ↓
┌──────────────────────────┐  ┌──────────────────────────┐
│      internal/ai          │  │    internal/repository     │
│   (AI Orchestration)       │  │     (Data Access)         │
│  ┌──────────────┐          │  │  ┌──────────────┐         │
│  │  gateway     │          │  │  │ curriculum   │         │
│  │  provider    │          │  │  │  learning     │         │
│  │  prompt      │          │  │  │  assessment   │         │
│  │  response    │          │  │  └──────────────┘         │
│  └──────────────┘          │  └──────────────────────────┘
└──────────────────────────┘                │
                                            ↓
                              ┌──────────────────────────┐
                              │       internal/db         │
                              │    (Database Layer)        │
                              └──────────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│                    internal/middleware                        │
│                    (Cross-Cutting Concerns)                   │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐      │
│  │    auth      │  │   logging     │  │    cors       │      │
│  └──────────────┘  └──────────────┘  └──────────────┘      │
└─────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│                   internal/infrastructure                      │
│                    (Infrastructure Support)                   │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐      │
│  │   config     │  │   logger     │  │   error       │      │
│  │   queue      │  │    util      │  │              │      │
│  └──────────────┘  └──────────────┘  └──────────────┘      │
└─────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│                      internal/domain                          │
│                    (Domain Models & Interfaces)               │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐      │
│  │ curriculum   │  │  learning     │  │  assessment   │      │
│  │    model     │  │    model      │  │    model      │      │
│  │ repository   │  │  repository   │  │  repository   │      │
│  └──────────────┘  └──────────────┘  └──────────────┘      │
└─────────────────────────────────────────────────────────────┘
```

---

# SECTION 13 — Dependency Diagram

## Dependency Graph

```
┌─────────────────────────────────────────────────────────────┐
│                        cmd/server                            │
│                      (No Dependencies)                        │
└─────────────────────────────────────────────────────────────┘
                            │ depends on
                            ↓
┌─────────────────────────────────────────────────────────────┐
│                    internal/handler                           │
│                    (Depends on: service, error)               │
└─────────────────────────────────────────────────────────────┘
                            │ depends on
                            ↓
┌─────────────────────────────────────────────────────────────┐
│                     internal/service                           │
│              (Depends on: repository, ai, queue)             │
└─────────────────────────────────────────────────────────────┘
                            │ depends on
            ┌───────────────┴───────────────┐
            ↓                               ↓
┌──────────────────────────┐    ┌──────────────────────────┐
│      internal/ai          │    │    internal/repository     │
│  (Depends on: config)     │    │ (Depends on: db, domain)  │
└──────────────────────────┘    └──────────────────────────┘
                                         │ depends on
                                         ↓
                              ┌──────────────────────────┐
                              │       internal/db         │
                              │  (Depends on: config)    │
                              └──────────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│                   internal/middleware                         │
│              (Depends on: logger, error, util)               │
└─────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│                  internal/infrastructure                       │
│  config: (No dependencies)                                    │
│  logger: (No dependencies)                                    │
│  error: (No dependencies)                                     │
│  util: (No dependencies)                                      │
│  queue: (Depends on: config)                                 │
└─────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│                      internal/domain                          │
│                  (No dependencies)                            │
└─────────────────────────────────────────────────────────────┘
```

---

# SECTION 14 — Startup Flow

## Startup Flowchart

```
START
  │
  ├─→ Load Configuration
  │   └─→ Read environment variables
  │   └─→ Set defaults
  │   └─→ Validate configuration
  │
  ├─→ Initialize Logger
  │   └─→ Create zap logger
  │   └─→ Set log level
  │
  ├─→ Connect to Database
  │   └─→ Create connection pool
  │   └─→ Test connection
  │   └─→ Set pool parameters
  │
  ├─→ Run Migrations (optional)
  │   └─→ Check migration flag
  │   └─→ Run golang-migrate
  │   └─→ Log migration status
  │
  ├─→ Connect to RabbitMQ
  │   └─→ Create connection
  │   └─→ Declare queues
  │   └─→ Test connection
  │
  ├─→ Initialize Repositories
  │   └─→ Create repository instances
  │   └─→ Inject database connection
  │
  ├─→ Initialize AI Gateway
  │   └─→ Load provider configurations
  │   └─→ Create provider instances
  │   └─→ Set primary and fallback
  │
  ├─→ Initialize Services
  │   └─→ Create service instances
  │   └─→ Inject repositories
  │   └─→ Inject AI gateway
  │   └─→ Inject queue
  │
  ├─→ Initialize Handlers
  │   └─→ Create handler instances
  │   └─→ Inject services
  │
  ├─→ Setup Router
  │   └─→ Create Gin router
  │   └─→ Apply middleware chain
  │   └─→ Register routes
  │
  ├─→ Start HTTP Server
  │   └─→ Listen on configured port
  │   └─→ Log server start
  │
  ├─→ Start Queue Consumers
  │   └─→ Start worker pool
  │   └─→ Begin consuming messages
  │
  ├─→ Wait for Shutdown Signal
  │   └─→ Listen for SIGINT/SIGTERM
  │
  └─→ Graceful Shutdown
      ├─→ Stop accepting new requests
      ├─→ Wait for in-flight requests (30s timeout)
      ├─→ Stop queue consumers
      ├─→ Close database connections
      ├─→ Close RabbitMQ connection
      └─→ Sync logger
      │
END
```

---

# SECTION 15 — Coding Standards

## Go Code Style

### Naming Conventions

- **Package names**: Lowercase, single word, no underscores or camelCase
- **Constants**: UPPER_SNAKE_CASE
- **Variables**: camelCase
- **Functions**: camelCase, exported functions start with uppercase
- **Interfaces**: Interface names should end with "er" or describe behavior
- **Structs**: PascalCase for exported, camelCase for unexported

### File Organization

- One struct/interface per file
- Related functions in same file
- Test files named `*_test.go` in same package
- Mock files named `*_mock.go` in separate package

### Error Handling

- Always handle errors explicitly
- Wrap errors with context using `fmt.Errorf`
- Use custom error types for application errors
- Log errors at appropriate level
- Return errors from functions

### Comments

- Exported functions must have comments
- Complex logic must have comments
- Package comments at top of file
- Use godoc format

### Example

```go
// Package curriculum handles curriculum-related business logic.
package curriculum

import (
    "context"
    "fmt"
)

// Service provides curriculum-related business operations.
type Service interface {
    // GenerateTPSet generates a TP Set for a given CP.
    GenerateTPSet(ctx context.Context, userID string, req GenerateTPSetRequest) (*TPSet, error)
}

type service struct {
    repo Repository
}

// NewService creates a new curriculum service.
func NewService(repo Repository) Service {
    return &service{repo: repo}
}

// GenerateTPSet generates a TP Set for a given CP using AI.
func (s *service) GenerateTPSet(ctx context.Context, userID string, req GenerateTPSetRequest) (*TPSet, error) {
    // Validate CP exists
    cp, err := s.repo.GetCP(ctx, req.CPID)
    if err != nil {
        return nil, fmt.Errorf("failed to get CP: %w", err)
    }
    
    // Generate TP Set
    // ... implementation
    
    return tpSet, nil
}
```

## Testing Standards

### Unit Tests

- Test individual functions in isolation
- Use table-driven tests for multiple cases
- Mock external dependencies
- Aim for >80% coverage

### Example

```go
func TestGenerateTPSet(t *testing.T) {
    tests := []struct {
        name       string
        userID     string
        req        GenerateTPSetRequest
        mockSetup  func(*mockRepository)
        wantErr    bool
        errType    error
    }{
        {
            name:   "success",
            userID: "user-1",
            req:    GenerateTPSetRequest{CPID: "cp-1"},
            mockSetup: func(m *mockRepository) {
                m.On("GetCP", mock.Anything, "cp-1").Return(&CP{ID: "cp-1"}, nil)
                m.On("BeginTx", mock.Anything).Return(&mockTransaction{}, nil)
                m.On("CreateAIGenerationLog", mock.Anything, mock.Anything).Return(nil)
                m.On("CreateTPSet", mock.Anything, mock.Anything).Return(nil)
                m.On("CreateTP", mock.Anything, mock.Anything).Return(nil)
            },
            wantErr: false,
        },
        {
            name:   "cp not found",
            userID: "user-1",
            req:    GenerateTPSetRequest{CPID: "cp-1"},
            mockSetup: func(m *mockRepository) {
                m.On("GetCP", mock.Anything, "cp-1").Return(nil, repository.ErrNotFound)
            },
            wantErr: true,
            errType: repository.ErrNotFound,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            mockRepo := new(mockRepository)
            tt.mockSetup(mockRepo)
            
            service := NewService(mockRepo)
            _, err := service.GenerateTPSet(context.Background(), tt.userID, tt.req)
            
            if tt.wantErr {
                if err == nil {
                    t.Errorf("expected error, got nil")
                }
                if !errors.Is(err, tt.errType) {
                    t.Errorf("expected error type %v, got %v", tt.errType, err)
                }
            } else {
                if err != nil {
                    t.Errorf("unexpected error: %v", err)
                }
            }
            
            mockRepo.AssertExpectations(t)
        })
    }
}
```

## Git Commit Standards

### Commit Message Format

```
<type>(<scope>): <subject>

<body>

<footer>
```

### Types

- **feat**: New feature
- **fix**: Bug fix
- **docs**: Documentation changes
- **style**: Code style changes (formatting)
- **refactor**: Code refactoring
- **test**: Test changes
- **chore**: Build process or auxiliary tool changes

### Example

```
feat(curriculum): add TP Set generation endpoint

- Add GenerateTPSet handler
- Add TP Set service logic
- Add TP Set repository methods
- Add AI generation integration

Closes #123
```

---

# SECTION 16 — Implementation Checklist

## Phase 1: Foundation Setup

- [ ] Initialize Go module
- [ ] Set up project structure
- [ ] Configure go.mod with dependencies
- [ ] Create Dockerfile
- [ ] Create docker-compose.yml
- [ ] Set up configuration loading
- [ ] Initialize logger
- [ ] Set up database connection
- [ ] Configure golang-migrate
- [ ] Create initial migration (init schema)

## Phase 2: Infrastructure Layer

- [ ] Implement error types
- [ ] Implement error handler
- [ ] Implement middleware (request ID, logging, CORS, recovery)
- [ ] Implement JWT utilities
- [ ] Implement validation utilities
- [ ] Implement time utilities
- [ ] Set up RabbitMQ connection
- [ ] Implement queue publisher
- [ ] Implement queue consumer
- [ ] Implement transaction manager

## Phase 3: Domain Layer

- [ ] Define curriculum domain models
- [ ] Define curriculum repository interfaces
- [ ] Define learning domain models
- [ ] Define learning repository interfaces
- [ ] Define assessment domain models
- [ ] Define assessment repository interfaces
- [ ] Define AI domain models
- [ ] Define AI repository interfaces
- [ ] Define user domain models
- [ ] Define user repository interfaces
- [ ] Define workflow domain models
- [ ] Define workflow repository interfaces

## Phase 4: Repository Layer

- [ ] Implement curriculum repository
- [ ] Implement learning repository
- [ ] Implement assessment repository
- [ ] Implement AI repository
- [ ] Implement user repository
- [ ] Implement workflow repository
- [ ] Write repository tests

## Phase 5: AI Orchestration Layer

- [ ] Implement AI gateway
- [ ] Implement OpenAI provider
- [ ] Implement Gemini provider
- [ ] Implement prompt builder
- [ ] Implement response processor
- [ ] Implement fallback logic
- [ ] Write AI tests

## Phase 6: Service Layer

- [ ] Implement curriculum service
- [ ] Implement learning service
- [ ] Implement assessment service
- [ ] Implement AI service
- [ ] Implement user service
- [ ] Implement workflow service
- [ ] Write service tests

## Phase 7: Handler Layer

- [ ] Implement curriculum handler
- [ ] Implement learning handler
- [ ] Implement assessment handler
- [ ] Implement AI handler
- [ ] Implement user handler
- [ ] Implement workflow handler
- [ ] Write handler tests

## Phase 8: Router Setup

- [ ] Set up Gin router
- [ ] Apply middleware chain
- [ ] Register curriculum routes
- [ ] Register learning routes
- [ ] Register assessment routes
- [ ] Register AI routes
- [ ] Register user routes
- [ ] Register workflow routes

## Phase 9: Startup Sequence

- [ ] Implement main.go
- [ ] Implement dependency injection
- [ ] Implement graceful shutdown
- [ ] Add health check endpoint
- [ ] Test startup sequence

## Phase 10: Testing

- [ ] Write integration tests
- [ ] Write end-to-end tests
- [ ] Set up test database
- [ ] Configure test environment
- [ ] Run test suite
- [ ] Measure code coverage

## Phase 11: Documentation

- [ ] Document API endpoints
- [ ] Document configuration options
- [ ] Document deployment process
- [ ] Document development setup
- [ ] Document troubleshooting

## Phase 12: Deployment

- [ ] Build Docker image
- [ ] Test Docker image locally
- [ ] Configure production environment
- [ ] Deploy to staging
- [ ] Run smoke tests
- [ ] Deploy to production
- [ ] Monitor deployment

---

# SECTION 17 — Appendix

## Dependencies

### Go Dependencies

```go
require (
    github.com/gin-gonic/gin v1.9.1
    github.com/jmoiron/sqlx v1.3.5
    github.com/lib/pq v1.10.9
    github.com/golang-jwt/jwt/v5 v5.0.0
    github.com/google/uuid v1.3.0
    github.com/streadway/amqp v1.1.0
    go.uber.org/zap v1.24.0
    golang.org/x/crypto v0.12.0
    golang.org/x/time v0.3.0
)
```

### Development Dependencies

```go
require (
    github.com/golang/mock v1.6.0
    github.com/stretchr/testify v1.8.4
    github.com/golang-migrate/migrate/v4 v4.15.2
)
```

## Useful Commands

```bash
# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Run linter
golangci-lint run

# Format code
go fmt ./...

# Vet code
go vet ./...

# Build
go build -o nusa-backend cmd/server/main.go

# Run
./nusa-backend

# Run with Docker
docker-compose up --build

# Run migrations
migrate -path migrations -database $DATABASE_URL up

# Generate sqlc queries
sqlc generate
```

## References

- [Gin Framework Documentation](https://gin-gonic.com/docs/)
- [sqlx Documentation](https://jmoiron.github.io/sqlx/)
- [golang-migrate Documentation](https://github.com/golang-migrate/migrate)
- [Zap Logging Documentation](https://github.com/uber-go/zap)
- [RabbitMQ Go Client](https://github.com/rabbitmq/amqp091-go)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)

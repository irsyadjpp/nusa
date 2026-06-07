# 18_TECHNICAL_FOUNDATION_BLUEPRINT.md

## Foundation Document for NUSA Education Platform

**Version**: 1.0
**Date**: June 2026
**Status**: FOUNDATION DOCUMENT - LOCKED
**Alignment**: Validated against Foundation Architecture (00A, 00B, 00C, 01, 02, 03, 04, 05, 06, 07, 08, 09, 10, 11, 12, 13, 14, 15, 16, 17)

**Purpose**: Define the technical foundation blueprint for NUSA MVP Wave 1 implementation, serving as the prerequisite for all feature development. This document establishes the foundational layers, frameworks, and patterns that must be implemented before any feature work begins.

---

# Technology Stack

## Backend Technology Stack

- **Language**: Go 1.21+
- **Web Framework**: Gin
- **Database**: PostgreSQL 18+
- **Message Broker**: RabbitMQ
- **Authentication**: Custom JWT
- **Migration Tool**: golang-migrate
- **ORM**: sqlx (no ORM, direct SQL)
- **Validation**: go-playground/validator
- **Logging**: logrus
- **Configuration**: viper

## Frontend Technology Stack

- **Framework**: React 18+
- **Language**: TypeScript 5+
- **Build Tool**: Vite
- **State Management**: React Context + Hooks
- **HTTP Client**: axios
- **Styling**: TailwindCSS
- **Validation**: zod
- **Routing**: React Router v6

## AI Runtime Technology Stack

- **Language**: Python 3.11+
- **AI Orchestration**: LangGraph
- **LLM Integration**: OpenAI API / Anthropic API
- **Web Framework**: FastAPI
- **Validation**: Pydantic
- **Communication**: HTTP REST

---

# Backend Foundation

## Database Connection Layer

### Purpose

Establish a single, reliable database connection pool for all database operations.

### Implementation

```go
// backend/internal/infrastructure/persistence/postgres/connection.go
package postgres

import (
    "context"
    "database/sql"
    "fmt"
    "time"

    _ "github.com/lib/pq"
    "github.com/jmoiron/sqlx"
)

type Connection struct {
    DB *sqlx.DB
}

func NewConnection(host, port, user, password, dbname string) (*Connection, error) {
    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    db, err := sqlx.Connect("postgres", dsn)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to database: %w", err)
    }

    // Configure connection pool
    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(5)
    db.SetConnMaxLifetime(5 * time.Minute)

    // Verify connection
    if err := db.Ping(); err != nil {
        return nil, fmt.Errorf("failed to ping database: %w", err)
    }

    return &Connection{DB: db}, nil
}

func (c *Connection) Close() error {
    return c.DB.Close()
}

func (c *Connection) WithTx(ctx context.Context, fn func(*sqlx.Tx) error) error {
    tx, err := c.DB.BeginTxx(ctx, nil)
    if err != nil {
        return err
    }

    defer func() {
        if p := recover(); p != nil {
            _ = tx.Rollback()
            panic(p)
        }
    }()

    if err := fn(tx); err != nil {
        if rbErr := tx.Rollback(); rbErr != nil {
            return fmt.Errorf("tx error: %v, rb error: %v", err, rbErr)
        }
        return err
    }

    return tx.Commit()
}
```

### Configuration

Environment variables:
- `DB_HOST`
- `DB_PORT`
- `DB_NAME`
- `DB_USER`
- `DB_PASSWORD`
- `DB_SSL_MODE`

### Dependencies

None (first layer)

### Implementation Order

**Priority 1** - Must be implemented first

---

## Migration Framework

### Purpose

Provide controlled, reproducible, and auditable database schema changes.

### Implementation

```go
// backend/scripts/migrate.go
package main

import (
    "log"
    "github.com/golang-migrate/migrate/v4"
    _ "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
    m, err := migrate.New(
        "file://migrations",
        "postgres://user:password@localhost:5432/nusa?sslmode=disable",
    )
    if err != nil {
        log.Fatal(err)
    }

    if err := m.Up(); err != nil {
        log.Fatal(err)
    }

    log.Println("Migrations completed successfully")
}
```

### Migration Naming Convention

YYYYMMDDHHMM_<description>

Examples:
- 202606050900_init_schema
- 202606051200_add_status_column

### Migration Structure

```
backend/migrations/
├── 000001_init_schema.up.sql
├── 000001_init_schema.down.sql
├── 000002_users.up.sql
├── 000002_users.down.sql
└── ...
```

### Dependencies

- Database Connection Layer

### Implementation Order

**Priority 2** - After Database Connection Layer

---

## Repository Pattern

### Purpose

Provide data access abstraction with clear separation between domain logic and data persistence.

### Interface Definition

```go
// backend/internal/modules/shared/domain/repositories/base_repository.go
package repositories

import "context"

type BaseRepository interface {
    Create(ctx context.Context, entity interface{}) error
    Update(ctx context.Context, entity interface{}) error
    Delete(ctx context.Context, id string) error
    FindByID(ctx context.Context, id string, dest interface{}) error
    FindAll(ctx context.Context, dest interface{}) error
}

type AuditableRepository interface {
    BaseRepository
    SetActor(ctx context.Context, actorID string)
}
```

### Implementation Template

```go
// backend/internal/modules/curriculum/infrastructure/repositories/cp_repository_impl.go
package repositories

import (
    "context"
    "database/sql"
    "fmt"
)

type CPRepositoryImpl struct {
    db *sqlx.DB
}

func NewCPRepository(db *sqlx.DB) *CPRepositoryImpl {
    return &CPRepositoryImpl{db: db}
}

func (r *CPRepositoryImpl) Create(ctx context.Context, cp *CP) error {
    query := `
        INSERT INTO cp (id, name, code, school_id, created_by, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
    `
    _, err := r.db.ExecContext(ctx, query, cp.ID, cp.Name, cp.Code, cp.SchoolID, cp.CreatedBy, cp.CreatedAt, cp.UpdatedAt)
    return err
}

func (r *CPRepositoryImpl) FindByID(ctx context.Context, id string) (*CP, error) {
    var cp CP
    query := `SELECT * FROM cp WHERE id = $1 AND deleted_at IS NULL`
    err := r.db.GetContext(ctx, &cp, query, id)
    if err == sql.ErrNoRows {
        return nil, fmt.Errorf("CP not found")
    }
    return &cp, err
}
```

### Dependencies

- Database Connection Layer

### Implementation Order

**Priority 3** - After Migration Framework

---

## Service Layer

### Purpose

Implement business logic and orchestrate repository operations.

### Interface Definition

```go
// backend/internal/modules/curriculum/application/services/cp_service.go
package services

import "context"

type CPService interface {
    CreateCP(ctx context.Context, req *CreateCPRequest) (*CP, error)
    UpdateCP(ctx context.Context, id string, req *UpdateCPRequest) (*CP, error)
    DeleteCP(ctx context.Context, id string) error
    GetCP(ctx context.Context, id string) (*CP, error)
    ListCPs(ctx context.Context, schoolID string) ([]*CP, error)
}
```

### Implementation Template

```go
// backend/internal/modules/curriculum/application/services/cp_service_impl.go
package services

import (
    "context"
    "fmt"
    "time"
)

type CPServiceImpl struct {
    repo repositories.CPRepository
    audit AuditService
}

func NewCPService(repo repositories.CPRepository, audit AuditService) *CPServiceImpl {
    return &CPServiceImpl{
        repo:  repo,
        audit: audit,
    }
}

func (s *CPServiceImpl) CreateCP(ctx context.Context, req *CreateCPRequest) (*CP, error) {
    actorID := s.audit.GetActor(ctx)

    cp := &CP{
        ID:        generateUUID(),
        Name:      req.Name,
        Code:      req.Code,
        SchoolID:  req.SchoolID,
        CreatedBy: actorID,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }

    if err := s.repo.Create(ctx, cp); err != nil {
        return nil, fmt.Errorf("failed to create CP: %w", err)
    }

    return cp, nil
}
```

### Dependencies

- Repository Pattern
- Audit Middleware

### Implementation Order

**Priority 4** - After Repository Pattern

---

## API Layer

### Purpose

Provide HTTP endpoints for client communication with proper request/response handling.

### Handler Template

```go
// backend/internal/api/handlers/curriculum/cp_handler.go
package curriculum

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type CPHandler struct {
    service services.CPService
}

func NewCPHandler(service services.CPService) *CPHandler {
    return &CPHandler{service: service}
}

func (h *CPHandler) CreateCP(c *gin.Context) {
    var req CreateCPRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    cp, err := h.service.CreateCP(c.Request.Context(), &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, cp)
}

func (h *CPHandler) GetCP(c *gin.Context) {
    id := c.Param("id")
    cp, err := h.service.GetCP(c.Request.Context(), id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, cp)
}
```

### Route Registration

```go
// backend/internal/api/routes/routes.go
package routes

import (
    "github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, handlers *Handlers) {
    api := r.Group("/api/v1")
    {
        curriculum := api.Group("/curriculum")
        {
            curriculum.POST("/cp", handlers.CP.CreateCP)
            curriculum.GET("/cp/:id", handlers.CP.GetCP)
            curriculum.PUT("/cp/:id", handlers.CP.UpdateCP)
            curriculum.DELETE("/cp/:id", handlers.CP.DeleteCP)
            curriculum.GET("/schools/:schoolId/cp", handlers.CP.ListCPs)
        }
    }
}
```

### Dependencies

- Service Layer
- Validation Framework
- Error Handling

### Implementation Order

**Priority 5** - After Service Layer

---

## JWT Authentication

### Purpose

Provide secure token-based authentication for API access.

### Implementation

```go
// backend/pkg/jwt/jwt.go
package jwt

import (
    "errors"
    "time"
    "github.com/golang-jwt/jwt/v5"
)

type Claims struct {
    UserID   string `json:"user_id"`
    Username string `json:"username"`
    Role     string `json:"role"`
    jwt.RegisteredClaims
}

type JWTManager struct {
    secretKey string
    duration  time.Duration
}

func NewJWTManager(secretKey string, duration time.Duration) *JWTManager {
    return &JWTManager{
        secretKey: secretKey,
        duration:  duration,
    }
}

func (m *JWTManager) GenerateToken(userID, username, role string) (string, error) {
    claims := &Claims{
        UserID:   userID,
        Username: username,
        Role:     role,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.duration)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(m.secretKey))
}

func (m *JWTManager) ValidateToken(tokenString string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(m.secretKey), nil
    })

    if err != nil {
        return nil, err
    }

    claims, ok := token.Claims.(*Claims)
    if !ok || !token.Valid {
        return nil, errors.New("invalid token")
    }

    return claims, nil
}
```

### Configuration

Environment variables:
- `JWT_SECRET`
- `JWT_EXPIRATION_HOURS`

### Dependencies

None (can be implemented in parallel with Database Connection Layer)

### Implementation Order

**Priority 1** - Can be implemented in parallel with Database Connection Layer

---

## RBAC Authorization

### Purpose

Enforce role-based access control for API endpoints.

### Implementation

```go
// backend/pkg/middleware/rbac.go
package middleware

import (
    "net/http"
    "strings"
    "github.com/gin-gonic/gin"
)

type RBACMiddleware struct {
    jwtManager *jwt.JWTManager
}

func NewRBACMiddleware(jwtManager *jwt.JWTManager) *RBACMiddleware {
    return &RBACMiddleware{jwtManager: jwtManager}
}

func (m *RBACMiddleware) RequireRole(roles ...string) gin.HandlerFunc {
    return func(c *gin.Context) {
        token := m.extractToken(c)
        if token == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
            return
        }

        claims, err := m.jwtManager.ValidateToken(token)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
            return
        }

        if !m.hasRole(claims.Role, roles) {
            c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "insufficient permissions"})
            return
        }

        c.Set("user_id", claims.UserID)
        c.Set("role", claims.Role)
        c.Next()
    }
}

func (m *RBACMiddleware) extractToken(c *gin.Context) string {
    authHeader := c.GetHeader("Authorization")
    if authHeader == "" {
        return ""
    }

    parts := strings.Split(authHeader, " ")
    if len(parts) != 2 || parts[0] != "Bearer" {
        return ""
    }

    return parts[1]
}

func (m *RBACMiddleware) hasRole(userRole string, requiredRoles []string) bool {
    for _, role := range requiredRoles {
        if userRole == role {
            return true
        }
    }
    return false
}
```

### Usage

```go
// backend/internal/api/routes/routes.go
api.POST("/cp", rbac.RequireRole("TEACHER", "ADMIN"), handlers.CP.CreateCP)
```

### Dependencies

- JWT Authentication

### Implementation Order

**Priority 2** - After JWT Authentication

---

## Audit Middleware

### Purpose

Automatically populate audit fields (created_by, updated_by, created_at, updated_at) for all database operations.

### Implementation

```go
// backend/internal/api/middleware/audit.go
package middleware

import (
    "context"
    "github.com/gin-gonic/gin"
)

const (
    ActorIDKey = "actor_id"
)

func AuditMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        userID := c.GetString("user_id")
        if userID != "" {
            ctx := context.WithValue(c.Request.Context(), ActorIDKey, userID)
            c.Request = c.Request.WithContext(ctx)
        }
        c.Next()
    }
}

func GetActor(ctx context.Context) string {
    actorID, ok := ctx.Value(ActorIDKey).(string)
    if !ok {
        return ""
    }
    return actorID
}
```

### Audit Service

```go
// backend/internal/modules/shared/infrastructure/audit_service.go
package infrastructure

import (
    "context"
    "time"
)

type AuditService struct{}

func NewAuditService() *AuditService {
    return &AuditService{}
}

func (s *AuditService) GetActor(ctx context.Context) string {
    return GetActor(ctx)
}

func (s *AuditService) SetAuditFields(entity interface{}, actorID string) {
    // Use reflection to set audit fields
    // created_by, updated_by, created_at, updated_at
}

func (s *AuditService) UpdateAuditFields(entity interface{}, actorID string) {
    // Use reflection to update audit fields
    // updated_by, updated_at
}
```

### Dependencies

- JWT Authentication

### Implementation Order

**Priority 2** - After JWT Authentication

---

## Error Handling

### Purpose

Provide consistent error handling and response formatting across all API endpoints.

### Implementation

```go
// backend/internal/api/middleware/error.go
package middleware

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type ErrorResponse struct {
    Error   string `json:"error"`
    Code    string `json:"code"`
    Details string `json:"details,omitempty"`
}

func ErrorHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next()

        if len(c.Errors) > 0 {
            err := c.Errors.Last().Err
            c.JSON(http.StatusInternalServerError, ErrorResponse{
                Error:   "Internal server error",
                Code:    "INTERNAL_ERROR",
                Details: err.Error(),
            })
        }
    }
}

func ValidationErrorHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next()

        if len(c.Errors) > 0 {
            err := c.Errors.Last()
            c.JSON(http.StatusBadRequest, ErrorResponse{
                Error:   "Validation error",
                Code:    "VALIDATION_ERROR",
                Details: err.Error(),
            })
        }
    }
}
```

### Custom Error Types

```go
// backend/internal/modules/shared/domain/errors/errors.go
package errors

import "errors"

var (
    ErrNotFound       = errors.New("resource not found")
    ErrUnauthorized   = errors.New("unauthorized")
    ErrForbidden      = errors.New("forbidden")
    ErrValidation     = errors.New("validation error")
    ErrConflict       = errors.New("conflict")
    ErrInternal       = errors.New("internal server error")
)

type AppError struct {
    Code    string
    Message string
    Err     error
}

func (e *AppError) Error() string {
    return e.Message
}

func (e *AppError) Unwrap() error {
    return e.Err
}
```

### Dependencies

None (can be implemented independently)

### Implementation Order

**Priority 1** - Can be implemented independently

---

## Validation Framework

### Purpose

Provide request validation using struct tags and automatic error handling.

### Implementation

```go
// backend/internal/api/middleware/validation.go
package middleware

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateRequest(req interface{}) gin.HandlerFunc {
    return func(c *gin.Context) {
        if err := c.ShouldBindJSON(req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "Validation failed",
                "details": err.Error(),
            })
            c.Abort()
            return
        }

        if err := validate.Struct(req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "Validation failed",
                "details": err.Error(),
            })
            c.Abort()
            return
        }

        c.Next()
    }
}
```

### Request DTO Example

```go
// backend/internal/modules/curriculum/application/dto/cp_dto.go
package dto

type CreateCPRequest struct {
    Name     string `json:"name" binding:"required,min=3,max=255"`
    Code     string `json:"code" binding:"required,min=2,max=50"`
    SchoolID string `json:"school_id" binding:"required,uuid"`
}
```

### Dependencies

None (can be implemented independently)

### Implementation Order

**Priority 1** - Can be implemented independently

---

## AI Client Abstraction

### Purpose

Provide a unified interface for LLM provider integration, enabling provider switching without code changes.

### Interface Definition

```go
// backend/internal/modules/ai/infrastructure/llm_client.go
package infrastructure

import "context"

type LLMClient interface {
    Generate(ctx context.Context, prompt string) (string, error)
    GenerateWithStream(ctx context.Context, prompt string) (<-chan string, error)
}
```

### OpenAI Implementation

```go
// backend/internal/modules/ai/infrastructure/openai_client.go
package infrastructure

import (
    "context"
    "github.com/sashabaranov/go-openai"
)

type OpenAIClient struct {
    client *openai.Client
}

func NewOpenAIClient(apiKey string) *OpenAIClient {
    return &OpenAIClient{
        client: openai.NewClient(apiKey),
    }
}

func (c *OpenAIClient) Generate(ctx context.Context, prompt string) (string, error) {
    resp, err := c.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
        Model: openai.GPT4,
        Messages: []openai.ChatCompletionMessage{
            {
                Role:    openai.ChatMessageRoleUser,
                Content: prompt,
            },
        },
    })

    if err != nil {
        return "", err
    }

    return resp.Choices[0].Message.Content, nil
}
```

### Factory Pattern

```go
// backend/internal/modules/ai/infrastructure/llm_factory.go
package infrastructure

type LLMProvider string

const (
    ProviderOpenAI LLMProvider = "openai"
    ProviderAnthropic LLMProvider = "anthropic"
)

func NewLLMClient(provider LLMProvider, apiKey string) LLMClient {
    switch provider {
    case ProviderOpenAI:
        return NewOpenAIClient(apiKey)
    default:
        return NewOpenAIClient(apiKey) // Default to OpenAI
    }
}
```

### Configuration

Environment variables:
- `LLM_PROVIDER`
- `LLM_API_KEY`
- `LLM_MODEL`

### Dependencies

None (can be implemented independently)

### Implementation Order

**Priority 1** - Can be implemented independently

---

## RabbitMQ Integration

### Purpose

Provide message publishing and consumption for domain events and workflow coordination.

### Implementation

```go
// backend/internal/infrastructure/messaging/rabbitmq/connection.go
package rabbitmq

import (
    "context"
    "amqp"
)

type Connection struct {
    conn *amqp.Connection
    ch   *amqp.Channel
}

func NewConnection(url string) (*Connection, error) {
    conn, err := amqp.Dial(url)
    if err != nil {
        return nil, err
    }

    ch, err := conn.Channel()
    if err != nil {
        conn.Close()
        return nil, err
    }

    return &Connection{conn: conn, ch: ch}, nil
}

func (c *Connection) Close() error {
    if c.ch != nil {
        c.ch.Close()
    }
    if c.conn != nil {
        c.conn.Close()
    }
    return nil
}

func (c *Connection) DeclareExchange(name, kind string) error {
    return c.ch.ExchangeDeclare(name, kind, true, false, false, false, nil)
}

func (c *Connection) DeclareQueue(name string) error {
    _, err := c.ch.QueueDeclare(name, true, false, false, false, nil)
    return err
}

func (c *Connection) BindQueue(queue, exchange, routingKey string) error {
    return c.ch.QueueBind(queue, routingKey, exchange, false, nil)
}
```

### Publisher

```go
// backend/internal/infrastructure/messaging/rabbitmq/publisher.go
package rabbitmq

import (
    "context"
    "encoding/json"
)

type Publisher struct {
    conn *Connection
}

func NewPublisher(conn *Connection) *Publisher {
    return &Publisher{conn: conn}
}

func (p *Publisher) Publish(ctx context.Context, exchange, routingKey string, message interface{}) error {
    body, err := json.Marshal(message)
    if err != nil {
        return err
    }

    return p.conn.ch.PublishWithContext(ctx, exchange, routingKey, false, false, amqp.Publishing{
        ContentType: "application/json",
        Body:        body,
    })
}
```

### Consumer

```go
// backend/internal/infrastructure/messaging/rabbitmq/consumer.go
package rabbitmq

import (
    "context"
    "encoding/json"
)

type Consumer struct {
    conn *Connection
}

func NewConsumer(conn *Connection) *Consumer {
    return &Consumer{conn: conn}
}

func (c *Consumer) Consume(ctx context.Context, queue string, handler func([]byte) error) error {
    msgs, err := c.conn.ch.Consume(queue, "", false, false, false, false, nil)
    if err != nil {
        return err
    }

    for {
        select {
        case <-ctx.Done():
            return ctx.Err()
        case msg := <-msgs:
            if err := handler(msg.Body); err != nil {
                msg.Nack(false, true)
            } else {
                msg.Ack(false)
            }
        }
    }
}
```

### Configuration

Environment variables:
- `RABBITMQ_HOST`
- `RABBITMQ_PORT`
- `RABBITMQ_USER`
- `RABBITMQ_PASSWORD`
- `RABBITMQ_VHOST`

### Dependencies

None (can be implemented independently)

### Implementation Order

**Priority 1** - Can be implemented independently

---

## Configuration Management

### Purpose

Centralize configuration loading from environment variables.

### Implementation

```go
// backend/internal/config/config.go
package config

import (
    "github.com/spf13/viper"
)

type Config struct {
    Database DatabaseConfig
    RabbitMQ RabbitMQConfig
    JWT      JWTConfig
    Server   ServerConfig
    AI       AIConfig
}

type DatabaseConfig struct {
    Host     string
    Port     string
    Name     string
    User     string
    Password string
    SSLMode  string
}

type RabbitMQConfig struct {
    Host     string
    Port     string
    User     string
    Password string
    VHost    string
}

type JWTConfig struct {
    Secret          string
    ExpirationHours int
}

type ServerConfig struct {
    Port string
    Mode string
}

type AIConfig struct {
    Provider string
    APIKey   string
    Model    string
}

func Load() (*Config, error) {
    viper.SetConfigName(".env")
    viper.SetConfigType("env")
    viper.AddConfigPath("../../")
    viper.AddConfigPath("./")
    viper.AutomaticEnv()

    if err := viper.ReadInConfig(); err != nil {
        return nil, err
    }

    var config Config
    if err := viper.Unmarshal(&config); err != nil {
        return nil, err
    }

    return &config, nil
}
```

### Dependencies

None (can be implemented independently)

### Implementation Order

**Priority 1** - Can be implemented independently

---

## Logging

### Purpose

Provide structured logging for debugging, monitoring, and audit trails.

### Implementation

```go
// backend/internal/infrastructure/logging/logger.go
package logging

import (
    "os"
    "github.com/sirupsen/logrus"
)

type Logger struct {
    *logrus.Logger
}

func NewLogger() *Logger {
    log := logrus.New()
    log.SetOutput(os.Stdout)
    log.SetFormatter(&logrus.JSONFormatter{})
    log.SetLevel(logrus.InfoLevel)

    return &Logger{Logger: log}
}

func (l *Logger) WithFields(fields logrus.Fields) *logrus.Entry {
    return l.Logger.WithFields(fields)
}

func (l *Logger) WithContext(ctx context.Context) *logrus.Entry {
    return l.Logger.WithContext(ctx)
}
```

### Usage

```go
logger := logging.NewLogger()
logger.WithFields(logrus.Fields{
    "user_id": userID,
    "action":  "create_cp",
}).Info("CP created successfully")
```

### Dependencies

None (can be implemented independently)

### Implementation Order

**Priority 1** - Can be implemented independently

---

## Health Check

### Purpose

Provide health check endpoint for monitoring and load balancer integration.

### Implementation

```go
// backend/internal/api/handlers/health_handler.go
package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type HealthHandler struct {
    db *sqlx.DB
}

func NewHealthHandler(db *sqlx.DB) *HealthHandler {
    return &HealthHandler{db: db}
}

func (h *HealthHandler) Health(c *gin.Context) {
    // Check database connection
    if err := h.db.Ping(); err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "status": "unhealthy",
            "checks": map[string]string{
                "database": "down",
            },
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "status": "healthy",
        "checks": map[string]string{
            "database": "up",
        },
    })
}
```

### Route

```go
api.GET("/health", handlers.Health.Health)
```

### Dependencies

- Database Connection Layer

### Implementation Order

**Priority 3** - After Database Connection Layer

---

# AI Runtime Foundation

## FastAPI Application Setup

### Purpose

Provide the FastAPI application foundation for the AI Runtime service.

### Implementation

```python
# ai-runtime/app/main.py
from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware

app = FastAPI(
    title="NUSA AI Runtime",
    description="AI orchestration service using LangGraph",
    version="1.0.0"
)

app.add_middleware(
    CORSMiddleware,
    allow_origins=["http://localhost:3000"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

@app.get("/health")
async def health_check():
    return {"status": "healthy"}
```

### Dependencies

None (can be implemented independently)

### Implementation Order

**Priority 1** - Can be implemented independently

---

## LangGraph Orchestration

### Purpose

Provide LangGraph workflow orchestration for AI agent execution.

### Implementation

```python
# ai-runtime/app/graph/workflow.py
from langgraph.graph import StateGraph, END
from typing import TypedDict, Annotated
import operator

class AgentState(TypedDict):
    messages: Annotated[list, operator.add]
    current_step: str
    result: dict

def create_tp_graph():
    workflow = StateGraph(AgentState)
    
    # Define nodes
    workflow.add_node("prompt_builder", build_prompt)
    workflow.add_node("llm_caller", call_llm)
    workflow.add_node("output_validator", validate_output)
    
    # Define edges
    workflow.set_entry_point("prompt_builder")
    workflow.add_edge("prompt_builder", "llm_caller")
    workflow.add_edge("llm_caller", "output_validator")
    workflow.add_edge("output_validator", END)
    
    return workflow.compile()
```

### Dependencies

- FastAPI Application Setup

### Implementation Order

**Priority 2** - After FastAPI Application Setup

---

## LLM Integration

### Purpose

Provide LLM provider integration for AI agent execution.

### Implementation

```python
# ai-runtime/app/llm/client.py
from openai import AsyncOpenAI
from anthropic import Anthropic

class LLMClient:
    def __init__(self, provider: str, api_key: str):
        if provider == "openai":
            self.client = AsyncOpenAI(api_key=api_key)
        elif provider == "anthropic":
            self.client = Anthropic(api_key=api_key)
    
    async def generate(self, prompt: str, model: str) -> str:
        if isinstance(self.client, AsyncOpenAI):
            response = await self.client.chat.completions.create(
                model=model,
                messages=[{"role": "user", "content": prompt}]
            )
            return response.choices[0].message.content
        # Add Anthropic implementation
```

### Dependencies

None (can be implemented independently)

### Implementation Order

**Priority 1** - Can be implemented independently

---

## HTTP Communication Layer

### Purpose

Provide HTTP endpoints for Backend API to communicate with AI Runtime.

### Implementation

```python
# ai-runtime/app/api/endpoints.py
from fastapi import HTTPException
from pydantic import BaseModel
from app.graph.workflow import create_tp_graph

class TPRequest(BaseModel):
    cp_id: str
    user_id: str
    context: dict

class TPResponse(BaseModel):
    result: dict
    confidence: float
    metadata: dict

@app.post("/api/v1/agents/tp/generate", response_model=TPResponse)
async def generate_tp(request: TPRequest):
    graph = create_tp_graph()
    result = await graph.ainvoke({
        "messages": [],
        "current_step": "start",
        "result": {},
        "context": request.context
    })
    return TPResponse(
        result=result["result"],
        confidence=0.85,
        metadata={"agent_version": "1.0.0"}
    )
```

### Dependencies

- LangGraph Orchestration
- LLM Integration

### Implementation Order

**Priority 3** - After LangGraph Orchestration and LLM Integration

---

## Configuration Management

### Purpose

Centralize configuration loading from environment variables.

### Implementation

```python
# ai-runtime/app/config/settings.py
from pydantic_settings import BaseSettings

class Settings(BaseSettings):
    # Server
    host: str = "0.0.0.0"
    port: int = 8000
    
    # LLM
    llm_provider: str = "openai"
    llm_api_key: str
    llm_model: str = "gpt-4"
    
    # Backend API
    backend_api_url: str = "http://localhost:8080"
    
    class Config:
        env_file = ".env"

settings = Settings()
```

### Dependencies

None (can be implemented independently)

### Implementation Order

**Priority 1** - Can be implemented independently

---

# Frontend Foundation

## Authentication

### Purpose

Provide authentication state management and token handling.

### Implementation

```typescript
// frontend/src/contexts/AuthContext.tsx
import React, { createContext, useContext, useState, useEffect } from 'react';
import { authService } from '../services/auth/authService';

interface AuthContextType {
  user: User | null;
  token: string | null;
  login: (username: string, password: string) => Promise<void>;
  logout: () => void;
  isAuthenticated: boolean;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const AuthProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const [user, setUser] = useState<User | null>(null);
  const [token, setToken] = useState<string | null>(null);

  useEffect(() => {
    const storedToken = localStorage.getItem('token');
    if (storedToken) {
      setToken(storedToken);
      // Validate token and fetch user
    }
  }, []);

  const login = async (username: string, password: string) => {
    const response = await authService.login(username, password);
    setToken(response.token);
    setUser(response.user);
    localStorage.setItem('token', response.token);
  };

  const logout = () => {
    setUser(null);
    setToken(null);
    localStorage.removeItem('token');
  };

  return (
    <AuthContext.Provider value={{ user, token, login, logout, isAuthenticated: !!token }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error('useAuth must be used within AuthProvider');
  }
  return context;
};
```

### Dependencies

None (can be implemented independently)

### Implementation Order

**Priority 1** - Can be implemented independently

---

## Protected Routes

### Purpose

Protect routes that require authentication.

### Implementation

```typescript
// frontend/src/components/layout/ProtectedRoute.tsx
import React from 'react';
import { Navigate } from 'react-router-dom';
import { useAuth } from '../../contexts/AuthContext';

interface ProtectedRouteProps {
  children: React.ReactNode;
}

export const ProtectedRoute: React.FC<ProtectedRouteProps> = ({ children }) => {
  const { isAuthenticated } = useAuth();

  if (!isAuthenticated) {
    return <Navigate to="/login" replace />;
  }

  return <>{children}</>;
};
```

### Usage

```typescript
<Route path="/dashboard" element={<ProtectedRoute><DashboardPage /></ProtectedRoute>} />
```

### Dependencies

- Authentication

### Implementation Order

**Priority 2** - After Authentication

---

## API Client

### Purpose

Provide centralized HTTP client with interceptors for authentication and error handling.

### Implementation

```typescript
// frontend/src/services/api/apiClient.ts
import axios, { AxiosInstance, AxiosError } from 'axios';
import { apiConfig } from '../../config/api.config';

class ApiClient {
  private client: AxiosInstance;

  constructor() {
    this.client = axios.create({
      baseURL: apiConfig.baseURL,
      timeout: apiConfig.timeout,
    });

    this.setupInterceptors();
  }

  private setupInterceptors() {
    // Request interceptor
    this.client.interceptors.request.use((config) => {
      const token = localStorage.getItem('token');
      if (token) {
        config.headers.Authorization = `Bearer ${token}`;
      }
      return config;
    });

    // Response interceptor
    this.client.interceptors.response.use(
      (response) => response,
      (error: AxiosError) => {
        if (error.response?.status === 401) {
          // Handle unauthorized
          localStorage.removeItem('token');
          window.location.href = '/login';
        }
        return Promise.reject(error);
      }
    );
  }

  public get<T>(url: string, config?: any) {
    return this.client.get<T>(url, config);
  }

  public post<T>(url: string, data?: any, config?: any) {
    return this.client.post<T>(url, data, config);
  }

  public put<T>(url: string, data?: any, config?: any) {
    return this.client.put<T>(url, data, config);
  }

  public delete<T>(url: string, config?: any) {
    return this.client.delete<T>(url, config);
  }
}

export const apiClient = new ApiClient();
```

### Dependencies

- Authentication (for token handling)

### Implementation Order

**Priority 2** - After Authentication

---

## State Management

### Purpose

Provide global state management for application data.

### Implementation

```typescript
// frontend/src/contexts/CPContext.tsx
import React, { createContext, useContext, useState, ReactNode } from 'react';
import { apiClient } from '../services/api/apiClient';

interface CP {
  id: string;
  name: string;
  code: string;
}

interface CPContextType {
  cps: CP[];
  loading: boolean;
  error: string | null;
  fetchCPs: (schoolId: string) => Promise<void>;
  createCP: (data: CreateCPRequest) => Promise<void>;
}

const CPContext = createContext<CPContextType | undefined>(undefined);

export const CPProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
  const [cps, setCps] = useState<CP[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const fetchCPs = async (schoolId: string) => {
    setLoading(true);
    setError(null);
    try {
      const response = await apiClient.get<CP[]>(`/api/v1/curriculum/schools/${schoolId}/cp`);
      setCps(response.data);
    } catch (err) {
      setError('Failed to fetch CPs');
    } finally {
      setLoading(false);
    }
  };

  const createCP = async (data: CreateCPRequest) => {
    setLoading(true);
    setError(null);
    try {
      await apiClient.post('/api/v1/curriculum/cp', data);
      await fetchCPs(data.school_id);
    } catch (err) {
      setError('Failed to create CP');
    } finally {
      setLoading(false);
    }
  };

  return (
    <CPContext.Provider value={{ cps, loading, error, fetchCPs, createCP }}>
      {children}
    </CPContext.Provider>
  );
};

export const useCP = () => {
  const context = useContext(CPContext);
  if (!context) {
    throw new Error('useCP must be used within CPProvider');
  }
  return context;
};
```

### Dependencies

- API Client

### Implementation Order

**Priority 3** - After API Client

---

## Layout System

### Purpose

Provide consistent layout structure across all pages.

### Implementation

```typescript
// frontend/src/layouts/MainLayout.tsx
import React from 'react';
import { Header } from '../components/layout/Header';
import { Sidebar } from '../components/layout/Sidebar';

interface MainLayoutProps {
  children: React.ReactNode;
}

export const MainLayout: React.FC<MainLayoutProps> = ({ children }) => {
  return (
    <div className="flex h-screen bg-gray-100">
      <Sidebar />
      <div className="flex-1 flex flex-col">
        <Header />
        <main className="flex-1 overflow-y-auto p-6">
          {children}
        </main>
      </div>
    </div>
  );
};
```

```typescript
// frontend/src/layouts/AuthLayout.tsx
import React from 'react';

interface AuthLayoutProps {
  children: React.ReactNode;
}

export const AuthLayout: React.FC<AuthLayoutProps> = ({ children }) => {
  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <div className="max-w-md w-full">
        {children}
      </div>
    </div>
  );
};
```

### Dependencies

None (can be implemented independently)

### Implementation Order

**Priority 1** - Can be implemented independently

---

## Notification System

### Purpose

Provide toast notifications for user feedback.

### Implementation

```typescript
// frontend/src/contexts/NotificationContext.tsx
import React, { createContext, useContext, useState, ReactNode } from 'react';

type NotificationType = 'success' | 'error' | 'info' | 'warning';

interface Notification {
  id: string;
  type: NotificationType;
  message: string;
}

interface NotificationContextType {
  notifications: Notification[];
  addNotification: (type: NotificationType, message: string) => void;
  removeNotification: (id: string) => void;
}

const NotificationContext = createContext<NotificationContextType | undefined>(undefined);

export const NotificationProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
  const [notifications, setNotifications] = useState<Notification[]>([]);

  const addNotification = (type: NotificationType, message: string) => {
    const id = Date.now().toString();
    setNotifications([...notifications, { id, type, message }]);
    
    // Auto-remove after 5 seconds
    setTimeout(() => {
      removeNotification(id);
    }, 5000);
  };

  const removeNotification = (id: string) => {
    setNotifications(notifications.filter(n => n.id !== id));
  };

  return (
    <NotificationContext.Provider value={{ notifications, addNotification, removeNotification }}>
      {children}
    </NotificationContext.Provider>
  );
};

export const useNotification = () => {
  const context = useContext(NotificationContext);
  if (!context) {
    throw new Error('useNotification must be used within NotificationProvider');
  }
  return context;
};
```

### Dependencies

None (can be implemented independently)

### Implementation Order

**Priority 1** - Can be implemented independently

---

## Error Handling

### Purpose

Provide consistent error handling across the application.

### Implementation

```typescript
// frontend/src/utils/errorHandler.ts
export const handleApiError = (error: any): string => {
  if (error.response) {
    const { data, status } = error.response;
    
    if (status === 401) {
      return 'Unauthorized. Please login again.';
    }
    
    if (status === 403) {
      return 'You do not have permission to perform this action.';
    }
    
    if (status === 404) {
      return 'Resource not found.';
    }
    
    if (data?.error) {
      return data.error;
    }
    
    if (data?.message) {
      return data.message;
    }
  }
  
  if (error.message) {
    return error.message;
  }
  
  return 'An unexpected error occurred. Please try again.';
};
```

### Dependencies

None (can be implemented independently)

### Implementation Order

**Priority 1** - Can be implemented independently

---

## Loading States

### Purpose

Provide consistent loading indicators across the application.

### Implementation

```typescript
// frontend/src/components/common/LoadingSpinner.tsx
import React from 'react';

export const LoadingSpinner: React.FC = () => {
  return (
    <div className="flex items-center justify-center">
      <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
    </div>
  );
};

export const FullPageLoading: React.FC = () => {
  return (
    <div className="flex items-center justify-center min-h-screen">
      <LoadingSpinner />
    </div>
  );
};
```

### Dependencies

None (can be implemented independently)

### Implementation Order

**Priority 1** - Can be implemented independently

---

## Permission Handling

### Purpose

Provide role-based permission checks for UI elements.

### Implementation

```typescript
// frontend/src/utils/permissions.ts
export const hasPermission = (userRole: string, requiredRoles: string[]): boolean => {
  return requiredRoles.includes(userRole);
};

export const canCreateCP = (userRole: string): boolean => {
  return hasPermission(userRole, ['TEACHER', 'ADMIN']);
};

export const canDeleteCP = (userRole: string): boolean => {
  return hasPermission(userRole, ['ADMIN']);
};
```

### Usage

```typescript
const { user } = useAuth();

{canCreateCP(user?.role) && (
  <Button onClick={handleCreate}>Create CP</Button>
)}
```

### Dependencies

- Authentication

### Implementation Order

**Priority 2** - After Authentication

---

## Environment Configuration

### Purpose

Centralize environment-specific configuration.

### Implementation

```typescript
// frontend/src/config/api.config.ts
export const apiConfig = {
  baseURL: import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api/v1',
  timeout: parseInt(import.meta.env.VITE_API_TIMEOUT) || 30000,
};

export const authConfig = {
  tokenStorage: import.meta.env.VITE_TOKEN_STORAGE || 'localStorage',
  tokenRefreshInterval: parseInt(import.meta.env.VITE_TOKEN_REFRESH_INTERVAL) || 3600000,
};

export const featureConfig = {
  enableAIAssistance: import.meta.env.VITE_ENABLE_AI_ASSISTANCE === 'true',
  enableAnalytics: import.meta.env.VITE_ENABLE_ANALYTICS === 'true',
};
```

### Environment Variables

```bash
# API
VITE_API_BASE_URL=http://localhost:8080/api/v1
VITE_API_TIMEOUT=30000

# Authentication
VITE_TOKEN_STORAGE=localStorage
VITE_TOKEN_REFRESH_INTERVAL=3600000

# Features
VITE_ENABLE_AI_ASSISTANCE=true
VITE_ENABLE_ANALYTICS=false
```

### Dependencies

None (can be implemented independently)

### Implementation Order

**Priority 1** - Can be implemented independently

---

# AI Foundation

## Prompt Loader

### Purpose

Load prompt templates from file system for AI agents.

### Implementation

```go
// backend/internal/modules/ai/infrastructure/prompt_loader.go
package infrastructure

import (
    "embed"
    "fmt"
    "os"
)

//go:embed prompts/*.txt
var promptFS embed.FS

type PromptLoader struct {
    prompts map[string]string
}

func NewPromptLoader() (*PromptLoader, error) {
    loader := &PromptLoader{
        prompts: make(map[string]string),
    }

    // Load prompts from embedded file system
    files, err := promptFS.ReadDir("prompts")
    if err != nil {
        return nil, fmt.Errorf("failed to read prompts directory: %w", err)
    }

    for _, file := range files {
        content, err := promptFS.ReadFile("prompts/" + file.Name())
        if err != nil {
            return nil, fmt.Errorf("failed to read prompt file %s: %w", file.Name(), err)
        }
        loader.prompts[file.Name()] = string(content)
    }

    return loader, nil
}

func (l *PromptLoader) GetPrompt(agentName string) (string, error) {
    prompt, ok := l.prompts[agentName+".txt"]
    if !ok {
        return "", fmt.Errorf("prompt not found for agent: %s", agentName)
    }
    return prompt, nil
}
```

### Dependencies

None (can be implemented independently)

### Implementation Order

**Priority 1** - Can be implemented independently

---

## Prompt Versioning

### Purpose

Track prompt versions for reproducibility and rollback.

### Implementation

```go
// backend/internal/modules/ai/domain/prompt_version.go
package domain

import "time"

type PromptVersion struct {
    ID          string    `json:"id"`
    AgentName   string    `json:"agent_name"`
    Version     string    `json:"version"`
    Content     string    `json:"content"`
    CreatedBy   string    `json:"created_by"`
    CreatedAt   time.Time `json:"created_at"`
    IsActive    bool      `json:"is_active"`
}

type PromptVersionRepository interface {
    Save(ctx context.Context, pv *PromptVersion) error
    GetActive(ctx context.Context, agentName string) (*PromptVersion, error)
    GetVersion(ctx context.Context, agentName, version string) (*PromptVersion, error)
    ListVersions(ctx context.Context, agentName string) ([]*PromptVersion, error)
}
```

### Dependencies

- Repository Pattern

### Implementation Order

**Priority 4** - After Repository Pattern

---

## AI Provider Adapter

### Purpose

Provide adapter pattern for different LLM providers.

### Implementation

```go
// backend/internal/modules/ai/infrastructure/adapter.go
package infrastructure

import "context"

type AIProviderAdapter interface {
    Generate(ctx context.Context, prompt string, options GenerationOptions) (*GenerationResult, error)
    GenerateWithStream(ctx context.Context, prompt string, options GenerationOptions) (<-chan GenerationChunk, error)
}

type GenerationOptions struct {
    Model       string
    Temperature float64
    MaxTokens   int
}

type GenerationResult struct {
    Content   string
    TokensUsed int
    Model     string
}

type GenerationChunk struct {
    Content string
    Done    bool
}
```

### Dependencies

- AI Client Abstraction

### Implementation Order

**Priority 2** - After AI Client Abstraction

---

## Generation Logging

### Purpose

Log all AI generations for auditability and debugging.

### Implementation

```go
// backend/internal/modules/ai/domain/generation_log.go
package domain

import "time"

type GenerationLog struct {
    ID          string    `json:"id"`
    AgentName   string    `json:"agent_name"`
    Prompt      string    `json:"prompt"`
    Response    string    `json:"response"`
    TokensUsed  int       `json:"tokens_used"`
    LatencyMs   int64     `json:"latency_ms"`
    Success     bool      `json:"success"`
    Error       string    `json:"error,omitempty"`
    CreatedBy   string    `json:"created_by"`
    CreatedAt   time.Time `json:"created_at"`
}

type GenerationLogRepository interface {
    Log(ctx context.Context, log *GenerationLog) error
    GetLogs(ctx context.Context, agentName string, limit int) ([]*GenerationLog, error)
}
```

### Dependencies

- Repository Pattern

### Implementation Order

**Priority 4** - After Repository Pattern

---

## Review Logging

### Purpose

Log human review decisions for AI-generated content.

### Implementation

```go
// backend/internal/modules/ai/domain/review_log.go
package domain

import "time"

type ReviewDecision string

const (
    ReviewApproved ReviewDecision = "APPROVED"
    ReviewRejected ReviewDecision = "REJECTED"
    ReviewModified ReviewDecision = "MODIFIED"
)

type ReviewLog struct {
    ID             string          `json:"id"`
    GenerationID   string          `json:"generation_id"`
    AgentName      string          `json:"agent_name"`
    Decision       ReviewDecision  `json:"decision"`
    Reason         string          `json:"reason"`
    ReviewedBy     string          `json:"reviewed_by"`
    ReviewedAt     time.Time       `json:"reviewed_at"`
}

type ReviewLogRepository interface {
    Log(ctx context.Context, log *ReviewLog) error
    GetLogs(ctx context.Context, generationID string) ([]*ReviewLog, error)
}
```

### Dependencies

- Repository Pattern

### Implementation Order

**Priority 4** - After Repository Pattern

---

## Feedback Logging

### Purpose

Log user feedback on AI-generated content for quality improvement.

### Implementation

```go
// backend/internal/modules/ai/domain/feedback_log.go
package domain

import "time"

type FeedbackRating int

const (
    RatingPoor FeedbackRating = 1
    RatingFair FeedbackRating = 2
    RatingGood FeedbackRating = 3
    RatingExcellent FeedbackRating = 4
)

type FeedbackLog struct {
    ID           string          `json:"id"`
    GenerationID string          `json:"generation_id"`
    AgentName    string          `json:"agent_name"`
    Rating       FeedbackRating  `json:"rating"`
    Comment      string          `json:"comment"`
    ProvidedBy   string          `json:"provided_by"`
    ProvidedAt   time.Time       `json:"provided_at"`
}

type FeedbackLogRepository interface {
    Log(ctx context.Context, log *FeedbackLog) error
    GetLogs(ctx context.Context, agentName string, limit int) ([]*FeedbackLog, error)
}
```

### Dependencies

- Repository Pattern

### Implementation Order

**Priority 4** - After Repository Pattern

---

# Operational Foundation

## Environment Strategy

### Development Environment

**Purpose**: Local development for individual developers.

**Configuration**:
- Local PostgreSQL database
- Local RabbitMQ instance
- Mock LLM provider or test API key
- Hot reload enabled
- Debug logging enabled

**Access**: Localhost

**Data**: Seed data for testing

### Staging Environment

**Purpose**: Pre-production testing and validation.

**Configuration**:
- Cloud PostgreSQL database
- Cloud RabbitMQ instance
- Production LLM provider API key
- Production-like configuration
- Info logging enabled

**Access**: Restricted access via VPN

**Data**: Sample production data (anonymized)

### Production Environment

**Purpose**: Live production deployment.

**Configuration**:
- Cloud PostgreSQL database (high availability)
- Cloud RabbitMQ instance (high availability)
- Production LLM provider API key
- Production configuration
- Error logging only

**Access**: Restricted access via VPN

**Data**: Real production data

---

# Non-Functional Requirements

## Security

### Authentication

- Custom JWT implementation with HS256 signing
- Token expiration: 24 hours
- Token refresh mechanism
- Secure token storage (httpOnly cookies recommended)

### Authorization

- Role-Based Access Control (RBAC)
- Permission checks on all API endpoints
- Permission checks on sensitive UI elements

### Data Protection

- Password hashing with bcrypt (cost factor 12)
- SQL injection prevention (parameterized queries)
- XSS prevention (input sanitization)
- CSRF protection (token-based)

### API Security

- Rate limiting (100 requests per minute per user)
- Request size limits (10MB max)
- CORS configuration (whitelist domains)
- Security headers (HSTS, X-Frame-Options, etc.)

### Audit Logging

- All CRUD operations logged
- AI generations logged
- Human reviews logged
- Audit trail retention: 90 days

## Performance

### Backend Performance

- API response time: P95 < 500ms (non-AI endpoints)
- AI generation: P95 < 5 seconds (simple), P95 < 15 seconds (complex)
- Database query time: P95 < 100ms
- Connection pool: 25 max connections, 5 idle connections

### Frontend Performance

- Initial page load: < 3 seconds
- Time to interactive: < 5 seconds
- Bundle size: < 500KB (gzipped)
- Lazy loading for routes and components

### Database Performance

- Indexes on frequently queried columns
- Query optimization
- Connection pooling
- Read replicas (future)

## Observability

### Logging

- Structured logging (JSON format)
- Log levels: DEBUG, INFO, WARN, ERROR
- Contextual logging (request ID, user ID)
- Log aggregation (future)

### Metrics

- Request rate
- Error rate
- Response time percentiles
- Database connection pool metrics
- AI generation metrics

### Tracing

- Request tracing (future)
- Distributed tracing (future)
- Performance profiling (future)

### Monitoring

- Health check endpoint
- Error alerting
- Performance alerting
- Uptime monitoring

## Scalability

### Horizontal Scaling

- Stateless application design
- Load balancer support
- Session state in JWT (no server-side sessions)
- Database connection pooling

### Vertical Scaling

- CPU-optimized instances for backend
- Memory-optimized instances for database
- GPU instances (future for AI)

### Database Scaling

- Read replicas (future)
- Database sharding (future)
- Caching layer (future)

## Maintainability

### Code Quality

- Code review required for all changes
- Unit test coverage > 80%
- Integration test coverage > 70%
- Linting and formatting (gofmt, ESLint)

### Documentation

- API documentation (OpenAPI/Swagger)
- Code comments for complex logic
- Architecture documentation
- Runbooks for operations

### Testing

- Unit tests for all business logic
- Integration tests for API endpoints
- End-to-end tests for critical flows
- Manual testing checklist

### Deployment

- Automated CI/CD pipeline
- Blue-green deployment (future)
- Canary deployment (future)
- Rollback capability

---

# Implementation Sequence

## Phase 1: Infrastructure Foundation (Days 1-2)

**Priority 1** - No dependencies

1. Configuration Management
2. Logging
3. Error Handling
4. Validation Framework
5. Database Connection Layer
6. Health Check
7. AI Client Abstraction
8. RabbitMQ Integration
9. JWT Authentication
10. Layout System (Frontend)
11. Notification System (Frontend)
12. Loading States (Frontend)
13. Error Handling (Frontend)
14. Environment Configuration (Frontend)

## Phase 2: Data Foundation (Days 3-4)

**Priority 2** - Depends on Phase 1

1. Migration Framework
2. Repository Pattern
3. Prompt Loader

## Phase 3: Business Foundation (Days 5-6)

**Priority 3** - Depends on Phase 2

1. Service Layer
2. RBAC Authorization
3. Audit Middleware
4. Prompt Versioning
5. AI Provider Adapter

## Phase 4: Integration Foundation (Days 7-8)

**Priority 4** - Depends on Phase 3

1. API Layer
2. Authentication (Frontend)
3. Protected Routes (Frontend)
4. API Client (Frontend)
5. State Management (Frontend)
6. Permission Handling (Frontend)

## Phase 5: AI Foundation (Days 9-10)

**Priority 5** - Depends on Phase 4

1. Generation Logging
2. Review Logging
3. Feedback Logging

---

# Dependency Order

## Backend Dependencies

```
Configuration Management
Logging
Error Handling
Validation Framework
↓
Database Connection Layer
Health Check
↓
Migration Framework
↓
Repository Pattern
↓
Service Layer
↓
API Layer
↓
JWT Authentication
↓
RBAC Authorization
Audit Middleware
↓
AI Client Abstraction
RabbitMQ Integration
↓
Prompt Loader
↓
Prompt Versioning
AI Provider Adapter
↓
Generation Logging
Review Logging
Feedback Logging
```

## Frontend Dependencies

```
Environment Configuration
Layout System
Notification System
Loading States
Error Handling
↓
Authentication
↓
Protected Routes
Permission Handling
↓
API Client
↓
State Management
```

## Cross-Platform Dependencies

```
Backend API Layer
↓
Frontend API Client
↓
Frontend State Management
```

---

**Document Status**: FOUNDATION DOCUMENT - LOCKED

**Implementation Status**: READY FOR IMPLEMENTATION

**Implementation Start**: TBD

**Implementation End**: TBD

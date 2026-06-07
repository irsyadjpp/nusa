# NUSA Backend Code Standards

## Layer Separation

The NUSA backend follows Clean Architecture principles with strict layer separation. Each layer has specific responsibilities and must not cross boundaries.

### Layer Hierarchy

```
Handler Layer (modules/*/handler.go)
    ↓
Service Layer (modules/*/service.go)
    ↓
Repository Layer (modules/*/repository.go)
    ↓
Database (internal/database)
```

### Layer Responsibilities

#### 1. Handler Layer (modules/*/handler.go)

**Responsibilities:**
- HTTP request/response handling
- Request validation
- Response formatting using pkg/response
- Calling service layer
- No business logic
- No direct database access

**Example:**
```go
func (h *Handler) Create(c *gin.Context) {
    var req CreateRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        response.ValidationError(c, map[string]string{"error": err.Error()})
        return
    }
    
    data, err := h.service.Create(c.Request.Context(), req)
    if err != nil {
        response.Error(c, http.StatusInternalServerError, err.Error())
        return
    }
    
    response.Created(c, data)
}
```

#### 2. Service Layer (modules/*/service.go)

**Responsibilities:**
- Business logic implementation
- Orchestrating repository calls
- Transaction management
- Domain validation
- No HTTP concerns
- No direct database queries (use repository)

**Example:**
```go
func (s *Service) Create(ctx context.Context, req CreateRequest) (*Entity, error) {
    if err := s.validateBusinessRules(ctx, req); err != nil {
        return nil, err
    }
    
    entity := s.toEntity(req)
    return s.repository.Create(ctx, entity)
}
```

#### 3. Repository Layer (modules/*/repository.go)

**Responsibilities:**
- Database operations
- SQL queries
- Data mapping (Entity ↔ DTO)
- No business logic
- No HTTP concerns

**Example:**
```go
func (r *Repository) Create(ctx context.Context, entity *Entity) (*Entity, error) {
    query := `INSERT INTO table (...) VALUES (...) RETURNING *`
    var result Entity
    err := r.db.GetContext(ctx, &result, query, ...)
    return &result, err
}
```

#### 4. DTO Layer (modules/*/dto.go)

**Responsibilities:**
- Data transfer objects for API requests/responses
- Serialization/deserialization
- Validation tags
- No business logic

**Example:**
```go
type CreateRequest struct {
    Name  string `json:"name" binding:"required,min=3,max=100"`
    Email string `json:"email" binding:"required,email"`
}
```

#### 5. Entity Layer (modules/*/entity.go)

**Responsibilities:**
- Domain entities
- Database table representations
- No validation tags (use DTOs)
- No business logic

**Example:**
```go
type User struct {
    ID        string    `db:"id"`
    Email     string    `db:"email"`
    CreatedAt time.Time `db:"created_at"`
}
```

### Forbidden Patterns

❌ **Handler calling repository directly**
```go
func (h *Handler) Create(c *gin.Context) {
    data, err := h.repository.Create(...) // FORBIDDEN
}
```

❌ **Service layer handling HTTP**
```go
func (s *Service) Create(c *gin.Context) { // FORBIDDEN
    // Service should not know about HTTP
}
```

❌ **Business logic in handlers**
```go
func (h *Handler) Create(c *gin.Context) {
    if user.Role != "ADMIN" { // FORBIDDEN - business logic
        return
    }
}
```

❌ **Direct SQL in service layer**
```go
func (s *Service) Create(ctx context.Context, req Request) {
    s.db.Query("INSERT INTO ...") // FORBIDDEN - use repository
}
```

### Required Patterns

✅ **Handler → Service → Repository**
```go
// Handler
func (h *Handler) Create(c *gin.Context) {
    data, err := h.service.Create(c.Request.Context(), req)
    response.Created(c, data)
}

// Service
func (s *Service) Create(ctx context.Context, req Request) (*Entity, error) {
    return s.repository.Create(ctx, entity)
}

// Repository
func (r *Repository) Create(ctx context.Context, entity *Entity) (*Entity, error) {
    // SQL operations
}
```

✅ **DTO for validation**
```go
type CreateRequest struct {
    Name string `json:"name" binding:"required,min=3"`
}
```

✅ **Entity for database**
```go
type Entity struct {
    ID   string `db:"id"`
    Name string `db:"name"`
}
```

### Package Structure per Module

Each module in `modules/` must follow this structure:

```
modules/example/
├── handler.go      # HTTP handlers
├── service.go      # Business logic
├── repository.go   # Database operations
├── dto.go          # Request/response DTOs
└── entity.go       # Domain entities
```

### Dependency Rules

- **Handler** depends on: Service, pkg/response, pkg/errors
- **Service** depends on: Repository, pkg/errors, pkg/jwt
- **Repository** depends on: internal/database, Entity
- **DTO** depends on: Nothing (pure data structures)
- **Entity** depends on: Nothing (pure data structures)

### Error Handling

All layers must use `pkg/errors` for consistent error handling:

```go
import "github.com/nusa/backend/pkg/errors"

// In service
if validationErr != nil {
    return errors.Wrap(validationErr, http.StatusBadRequest, "Validation failed")
}

// In handler
if err != nil {
    response.Error(c, http.StatusInternalServerError, err.Error())
}
```

### Transaction Management

Transactions must be managed in the Service layer using `internal/database/transaction`:

```go
func (s *Service) Create(ctx context.Context, req Request) error {
    tx, err := s.db.BeginTx(ctx)
    if err != nil {
        return err
    }
    defer func() {
        if err != nil {
            tx.Rollback()
        } else {
            tx.Commit()
        }
    }()
    
    // Use tx.GetTx() for repository operations
    return s.repository.CreateWithTx(ctx, tx.GetTx(), entity)
}
```

### Testing Standards

- **Handler tests**: Mock service layer
- **Service tests**: Mock repository layer
- **Repository tests**: Use test database or mock db

### Code Review Checklist

- [ ] Handler only calls service, not repository
- [ ] Service only calls repository, not database directly
- [ ] No business logic in handlers
- [ ] No HTTP concerns in services/repositories
- [ ] DTOs used for request/response
- [ ] Entities used for database operations
- [ ] Errors use pkg/errors
- [ ] Transactions managed in service layer
- [ ] Validation in DTOs using binding tags

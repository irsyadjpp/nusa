# 28_AUTHENTICATION_DESIGN.md

## Foundation Document for NUSA Education Platform

**Version**: 1.0
**Date**: June 2026
**Status**: FOUNDATION DOCUMENT
**Alignment**: Aligned with 14_DATABASE_SCHEMA.md, 13_API_CONTRACT.md, 27_BACKEND_FOUNDATION_DESIGN.md

**Purpose**: Design the authentication and authorization module for NUSA using custom JWT. This document defines user models, authentication flows, JWT structure, authorization middleware, and permission matrix for MVP.

---

# SECTION 1 — Executive Summary

## Why Authentication Design Matters

A well-designed authentication and authorization system ensures:
- Secure access to educational content and features
- Role-based access control for different user types
- Scalable session management for schools
- Audit trail for compliance
- Protection of sensitive student and teacher data

## Technology Stack

| Component | Technology | Purpose |
|-----------|------------|---------|
| Authentication | Custom JWT | Token-based authentication |
| Password Hashing | bcrypt | Secure password storage |
| Token Storage | Database (refresh_tokens table) | Refresh token persistence |
| Token Validation | JWT library | Access token validation |

## Core Principles

- **JWT-Based**: Stateless authentication using JWT access tokens
- **Role-Based Access Control (RBAC)**: Permissions based on user roles
- **School Isolation**: Users belong to schools with data isolation
- **Secure Password Storage**: bcrypt hashing for passwords
- **Token Rotation**: Refresh tokens for access token renewal
- **MVP-Focused**: Simple, pragmatic approach without OAuth/OIDC

---

# SECTION 2 — User Model

## User Entity

### User Table

| Field | Type | Constraints | Description |
|-------|------|-------------|-------------|
| `id` | UUID | PRIMARY KEY | Unique user identifier |
| `email` | VARCHAR(255) | NOT NULL, UNIQUE | User email address |
| `password_hash` | VARCHAR(255) | NOT NULL | Bcrypt hashed password |
| `name` | VARCHAR(255) | NOT NULL | Full name |
| `role` | VARCHAR(50) | NOT NULL | User role (SYSTEM_ADMIN, SCHOOL_ADMIN, TEACHER) |
| `school_id` | UUID | NULLABLE, FOREIGN KEY → schools(id) | Associated school (NULL for SYSTEM_ADMIN) |
| `is_active` | BOOLEAN | NOT NULL, DEFAULT TRUE | Account active status |
| `email_verified` | BOOLEAN | NOT NULL, DEFAULT FALSE | Email verification status |
| `last_login_at` | TIMESTAMP WITH TIME ZONE | NULLABLE | Last login timestamp |
| `created_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Account creation timestamp |
| `updated_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Last update timestamp |

### Constraints

```sql
-- Email uniqueness
CREATE UNIQUE INDEX idx_users_email ON users(email);

-- School and role index
CREATE INDEX idx_users_school_id ON users(school_id);
CREATE INDEX idx_users_role ON users(role);

-- Active users index
CREATE INDEX idx_users_is_active ON users(is_active);

-- Check constraint for role
ALTER TABLE users ADD CONSTRAINT chk_users_role 
CHECK (role IN ('SYSTEM_ADMIN', 'SCHOOL_ADMIN', 'TEACHER'));

-- Check constraint for school_id based on role
ALTER TABLE users ADD CONSTRAINT chk_users_school_role 
CHECK (
    (role = 'SYSTEM_ADMIN' AND school_id IS NULL) OR
    (role IN ('SCHOOL_ADMIN', 'TEACHER') AND school_id IS NOT NULL)
);
```

### User Model (Go)

```go
type User struct {
    ID            string     `json:"id" db:"id"`
    Email         string     `json:"email" db:"email"`
    PasswordHash  string     `json:"-" db:"password_hash"` // Never exposed
    Name          string     `json:"name" db:"name"`
    Role          string     `json:"role" db:"role"`
    SchoolID      *string    `json:"school_id,omitempty" db:"school_id"`
    IsActive      bool       `json:"is_active" db:"is_active"`
    EmailVerified bool       `json:"email_verified" db:"email_verified"`
    LastLoginAt   *time.Time `json:"last_login_at,omitempty" db:"last_login_at"`
    CreatedAt     time.Time  `json:"created_at" db:"created_at"`
    UpdatedAt     time.Time  `json:"updated_at" db:"updated_at"`
}
```

---

# SECTION 3 — Role Model

## Role Definitions

### SYSTEM_ADMIN

**Description**: Platform administrator with full system access.

**Scope**: System-wide access, not bound to any school.

**Responsibilities**:
- Manage schools (create, update, deactivate)
- Manage system-wide curriculum versions
- View system-wide analytics and reports
- Manage platform configuration
- Access all school data (read-only for audit)

### SCHOOL_ADMIN

**Description**: School administrator with school-level access.

**Scope**: Bound to a specific school.

**Responsibilities**:
- Manage school users (invite, deactivate)
- Manage school curriculum
- View school-level analytics and reports
- Approve/reject artifacts within school
- Manage school settings

### TEACHER

**Description**: Teacher with classroom-level access.

**Scope**: Bound to a specific school.

**Responsibilities**:
- Generate and manage teaching plans (TP Sets, ATP Sets, Modul Ajar Sets)
- Create and manage assessments
- View student data (assigned students only)
- Submit artifacts for approval
- View personal analytics

## Role Hierarchy

```
SYSTEM_ADMIN (highest)
    ↓
SCHOOL_ADMIN
    ↓
TEACHER (lowest)
```

**Inheritance**: Higher roles inherit all permissions of lower roles.

---

# SECTION 4 — School Model

## School Entity

### School Table

| Field | Type | Constraints | Description |
|-------|------|-------------|-------------|
| `id` | UUID | PRIMARY KEY | Unique school identifier |
| `name` | VARCHAR(255) | NOT NULL | School name |
| `code` | VARCHAR(50) | NOT NULL, UNIQUE | School code (e.g., "SDN-01-001") |
| `address` | TEXT | NULLABLE | School address |
| `phone` | VARCHAR(50) | NULLABLE | School phone number |
| `email` | VARCHAR(255) | NULLABLE | School email |
| `is_active` | BOOLEAN | NOT NULL, DEFAULT TRUE | School active status |
| `created_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | School creation timestamp |
| `updated_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Last update timestamp |

### Constraints

```sql
-- School code uniqueness
CREATE UNIQUE INDEX idx_schools_code ON schools(code);

-- Active schools index
CREATE INDEX idx_schools_is_active ON schools(is_active);
```

### School Model (Go)

```go
type School struct {
    ID        string    `json:"id" db:"id"`
    Name      string    `json:"name" db:"name"`
    Code      string    `json:"code" db:"code"`
    Address   *string   `json:"address,omitempty" db:"address"`
    Phone     *string   `json:"phone,omitempty" db:"phone"`
    Email     *string   `json:"email,omitempty" db:"email"`
    IsActive  bool      `json:"is_active" db:"is_active"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
```

---

# SECTION 5 — Login Flow

## Login Sequence Diagram

```
┌─────────┐       ┌─────────┐       ┌─────────┐       ┌─────────┐
│  Client │       │  Server │       │  Database│       │  JWT     │
└────┬────┘       └────┬────┘       └────┬────┘       └────┬────┘
     │                  │                  │                  │
     │ POST /api/v1/public/login         │                  │
     │ {email, password}                 │                  │
     ├─────────────────>│                  │                  │
     │                  │                  │                  │
     │                  │ SELECT user BY email              │
     │                  ├─────────────────>│                  │
     │                  │                  │                  │
     │                  │ User record     │                  │
     │                  │<─────────────────┤                  │
     │                  │                  │                  │
     │                  │ Verify password (bcrypt)          │
     │                  │                  │                  │
     │                  │ Generate access token             │
     │                  ├──────────────────────────────────>│
     │                  │                  │                  │
     │                  │ Access token    │                  │
     │                  │<──────────────────────────────────┤
     │                  │                  │                  │
     │                  │ Generate refresh token             │
     │                  ├──────────────────────────────────>│
     │                  │                  │                  │
     │                  │ Refresh token   │                  │
     │                  │<──────────────────────────────────┤
     │                  │                  │                  │
     │                  │ Save refresh token                │
     │                  ├─────────────────>│                  │
     │                  │                  │                  │
     │                  │ Update last_login_at              │
     │                  ├─────────────────>│                  │
     │                  │                  │                  │
     │                  │ Success         │                  │
     │                  │<─────────────────┤                  │
     │                  │                  │                  │
     │ {access_token, refresh_token, user} │                  │
     │<─────────────────┤                  │                  │
     │                  │                  │                  │
```

## Login API Flow

### Request

```http
POST /api/v1/public/login
Content-Type: application/json

{
  "email": "teacher@school.com",
  "password": "securepassword123"
}
```

### Response (Success)

```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "token_type": "Bearer",
  "expires_in": 3600,
  "user": {
    "id": "user-uuid",
    "email": "teacher@school.com",
    "name": "John Doe",
    "role": "TEACHER",
    "school_id": "school-uuid",
    "school_name": "SDN 01 Jakarta"
  }
}
```

### Response (Error)

```json
{
  "error": "Invalid credentials"
}
```

## Login Implementation

### Handler

```go
func (h *Handler) Login(c *gin.Context) {
    var req LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }
    
    result, err := h.service.Login(c.Request.Context(), req)
    if err != nil {
        if errors.Is(err, ErrInvalidCredentials) {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
            return
        }
        if errors.Is(err, ErrAccountInactive) {
            c.JSON(http.StatusForbidden, gin.H{"error": "Account is inactive"})
            return
        }
        handleError(c, err)
        return
    }
    
    c.JSON(http.StatusOK, result)
}
```

### Service

```go
func (s *service) Login(ctx context.Context, req LoginRequest) (*LoginResponse, error) {
    // Get user by email
    user, err := s.repo.GetUserByEmail(ctx, req.Email)
    if err != nil {
        if errors.Is(err, ErrNotFound) {
            return nil, ErrInvalidCredentials
        }
        return nil, err
    }
    
    // Verify password
    if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
        return nil, ErrInvalidCredentials
    }
    
    // Check if account is active
    if !user.IsActive {
        return nil, ErrAccountInactive
    }
    
    // Generate access token
    accessToken, err := s.jwt.GenerateAccessToken(user)
    if err != nil {
        return nil, err
    }
    
    // Generate refresh token
    refreshToken, err := s.jwt.GenerateRefreshToken(user.ID)
    if err != nil {
        return nil, err
    }
    
    // Save refresh token
    refreshRecord := &RefreshToken{
        ID:        uuid.New().String(),
        UserID:    user.ID,
        Token:     refreshToken,
        ExpiresAt: time.Now().Add(7 * 24 * time.Hour), // 7 days
        CreatedAt: time.Now(),
    }
    
    if err := s.repo.CreateRefreshToken(ctx, refreshRecord); err != nil {
        return nil, err
    }
    
    // Update last login
    now := time.Now()
    user.LastLoginAt = &now
    if err := s.repo.UpdateUser(ctx, user); err != nil {
        return nil, err
    }
    
    // Get school info if applicable
    var school *School
    if user.SchoolID != nil {
        school, err = s.repo.GetSchool(ctx, *user.SchoolID)
        if err != nil {
            return nil, err
        }
    }
    
    return &LoginResponse{
        AccessToken:  accessToken,
        RefreshToken: refreshToken,
        TokenType:    "Bearer",
        ExpiresIn:    3600, // 1 hour
        User: UserResponse{
            ID:         user.ID,
            Email:      user.Email,
            Name:       user.Name,
            Role:       user.Role,
            SchoolID:   user.SchoolID,
            SchoolName: school.Name,
        },
    }, nil
}
```

---

# SECTION 6 — Refresh Token Flow

## Refresh Token Sequence Diagram

```
┌─────────┐       ┌─────────┐       ┌─────────┐       ┌─────────┐
│  Client │       │  Server │       │  Database│       │  JWT     │
└────┬────┘       └────┬────┘       └────┬────┘       └────┬────┘
     │                  │                  │                  │
     │ POST /api/v1/public/refresh        │                  │
     │ {refresh_token}                    │                  │
     ├─────────────────>│                  │                  │
     │                  │                  │                  │
     │                  │ Validate refresh token             │
     │                  ├──────────────────────────────────>│
     │                  │                  │                  │
     │                  │ Valid           │                  │
     │                  │<──────────────────────────────────┤
     │                  │                  │                  │
     │                  │ Get refresh token record           │
     │                  ├─────────────────>│                  │
     │                  │                  │                  │
     │                  │ Refresh record  │                  │
     │                  │<─────────────────┤                  │
     │                  │                  │                  │
     │                  │ Check if expired/not revoked       │
     │                  │                  │                  │
     │                  │ Get user by user_id               │
     │                  ├─────────────────>│                  │
     │                  │                  │                  │
     │                  │ User record     │                  │
     │                  │<─────────────────┤                  │
     │                  │                  │                  │
     │                  │ Generate new access token          │
     │                  ├──────────────────────────────────>│
     │                  │                  │                  │
     │                  │ Access token    │                  │
     │                  │<──────────────────────────────────┤
     │                  │                  │                  │
     │                  │ Generate new refresh token         │
     │                  ├──────────────────────────────────>│
     │                  │                  │                  │
     │                  │ Refresh token   │                  │
     │                  │<──────────────────────────────────┤
     │                  │                  │                  │
     │                  │ Delete old refresh token          │
     │                  ├─────────────────>│                  │
     │                  │                  │                  │
     │                  │ Save new refresh token           │
     │                  ├─────────────────>│                  │
     │                  │                  │                  │
     │                  │ Success         │                  │
     │                  │<─────────────────┤                  │
     │                  │                  │                  │
     │ {access_token, refresh_token}       │                  │
     │<─────────────────┤                  │                  │
     │                  │                  │                  │
```

## Refresh Token API Flow

### Request

```http
POST /api/v1/public/refresh
Content-Type: application/json

{
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### Response (Success)

```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "token_type": "Bearer",
  "expires_in": 3600
}
```

### Response (Error)

```json
{
  "error": "Invalid or expired refresh token"
}
```

## Refresh Token Implementation

### Service

```go
func (s *service) RefreshToken(ctx context.Context, req RefreshTokenRequest) (*RefreshTokenResponse, error) {
    // Validate refresh token
    claims, err := s.jwt.ValidateRefreshToken(req.RefreshToken)
    if err != nil {
        return nil, ErrInvalidRefreshToken
    }
    
    // Get refresh token record
    refreshRecord, err := s.repo.GetRefreshToken(ctx, req.RefreshToken)
    if err != nil {
        if errors.Is(err, ErrNotFound) {
            return nil, ErrInvalidRefreshToken
        }
        return nil, err
    }
    
    // Check if expired
    if time.Now().After(refreshRecord.ExpiresAt) {
        return nil, ErrExpiredRefreshToken
    }
    
    // Check if revoked
    if refreshRecord.RevokedAt != nil {
        return nil, ErrRevokedRefreshToken
    }
    
    // Get user
    user, err := s.repo.GetUser(ctx, claims.UserID)
    if err != nil {
        return nil, err
    }
    
    // Check if account is active
    if !user.IsActive {
        return nil, ErrAccountInactive
    }
    
    // Generate new access token
    accessToken, err := s.jwt.GenerateAccessToken(user)
    if err != nil {
        return nil, err
    }
    
    // Generate new refresh token
    newRefreshToken, err := s.jwt.GenerateRefreshToken(user.ID)
    if err != nil {
        return nil, err
    }
    
    // Delete old refresh token
    if err := s.repo.DeleteRefreshToken(ctx, refreshRecord.ID); err != nil {
        return nil, err
    }
    
    // Save new refresh token
    newRefreshRecord := &RefreshToken{
        ID:        uuid.New().String(),
        UserID:    user.ID,
        Token:     newRefreshToken,
        ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
        CreatedAt: time.Now(),
    }
    
    if err := s.repo.CreateRefreshToken(ctx, newRefreshRecord); err != nil {
        return nil, err
    }
    
    return &RefreshTokenResponse{
        AccessToken:  accessToken,
        RefreshToken: newRefreshToken,
        TokenType:    "Bearer",
        ExpiresIn:    3600,
    }, nil
}
```

---

# SECTION 7 — Logout Flow

## Logout Sequence Diagram

```
┌─────────┐       ┌─────────┐       ┌─────────┐
│  Client │       │  Server │       │  Database│
└────┬────┘       └────┬────┘       └────┬────┘
     │                  │                  │
     │ POST /api/v1/auth/logout         │
     │ Authorization: Bearer <token>     │
     ├─────────────────>│                  │
     │                  │                  │
     │                  │ Extract user_id from JWT          │
     │                  │                  │
     │                  │ Delete refresh tokens for user    │
     │                  ├─────────────────>│                  │
     │                  │                  │                  │
     │                  │ Success         │                  │
     │                  │<─────────────────┤                  │
     │                  │                  │                  │
     │ 204 No Content   │                  │
     │<─────────────────┤                  │
     │                  │                  │
```

## Logout API Flow

### Request

```http
POST /api/v1/auth/logout
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

### Response

```http
HTTP/1.1 204 No Content
```

## Logout Implementation

### Service

```go
func (s *service) Logout(ctx context.Context, userID string) error {
    // Delete all refresh tokens for user
    if err := s.repo.DeleteRefreshTokensByUserID(ctx, userID); err != nil {
        return err
    }
    
    return nil
}
```

---

# SECTION 8 — JWT Claims Design

## JWT Structure

### Access Token Claims

```json
{
  "sub": "user-uuid",
  "email": "user@school.com",
  "name": "John Doe",
  "role": "TEACHER",
  "school_id": "school-uuid",
  "iat": 1717658400,
  "exp": 1717662000,
  "jti": "token-uuid"
}
```

### Claim Definitions

| Claim | Type | Description | Example |
|-------|------|-------------|---------|
| `sub` | string | User ID (subject) | "user-uuid" |
| `email` | string | User email | "user@school.com" |
| `name` | string | User name | "John Doe" |
| `role` | string | User role | "TEACHER" |
| `school_id` | string | School ID (NULL for SYSTEM_ADMIN) | "school-uuid" |
| `iat` | number | Issued at timestamp | 1717658400 |
| `exp` | number | Expiration timestamp | 1717662000 |
| `jti` | string | Token ID (unique identifier) | "token-uuid" |

### Refresh Token Claims

```json
{
  "sub": "user-uuid",
  "token_type": "refresh",
  "iat": 1717658400,
  "exp": 1718263200,
  "jti": "token-uuid"
}
```

### Refresh Token Claim Definitions

| Claim | Type | Description | Example |
|-------|------|-------------|---------|
| `sub` | string | User ID (subject) | "user-uuid" |
| `token_type` | string | Token type (always "refresh") | "refresh" |
| `iat` | number | Issued at timestamp | 1717658400 |
| `exp` | number | Expiration timestamp (7 days) | 1718263200 |
| `jti` | string | Token ID (unique identifier) | "token-uuid" |

## JWT Implementation

### JWT Service

```go
package jwt

import (
    "time"
    
    "github.com/golang-jwt/jwt/v5"
    "github.com/google/uuid"
)

type Claims struct {
    Sub       string `json:"sub"`
    Email     string `json:"email"`
    Name      string `json:"name"`
    Role      string `json:"role"`
    SchoolID  string `json:"school_id,omitempty"`
    TokenType string `json:"token_type,omitempty"`
    jwt.RegisteredClaims
}

type Service interface {
    GenerateAccessToken(user *User) (string, error)
    GenerateRefreshToken(userID string) (string, error)
    ValidateAccessToken(tokenString string) (*Claims, error)
    ValidateRefreshToken(tokenString string) (*Claims, error)
}

type service struct {
    secret          []byte
    accessTokenTTL  time.Duration
    refreshTokenTTL time.Duration
}

func NewService(secret string, accessTokenTTL, refreshTokenTTL time.Duration) Service {
    return &service{
        secret:          []byte(secret),
        accessTokenTTL:  accessTokenTTL,
        refreshTokenTTL: refreshTokenTTL,
    }
}

func (s *service) GenerateAccessToken(user *User) (string, error) {
    now := time.Now()
    claims := Claims{
        Sub:      user.ID,
        Email:    user.Email,
        Name:     user.Name,
        Role:     user.Role,
        SchoolID: *user.SchoolID,
        RegisteredClaims: jwt.RegisteredClaims{
            IssuedAt:  jwt.NewNumericDate(now),
            ExpiresAt: jwt.NewNumericDate(now.Add(s.accessTokenTTL)),
            ID:        uuid.New().String(),
        },
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(s.secret)
}

func (s *service) GenerateRefreshToken(userID string) (string, error) {
    now := time.Now()
    claims := Claims{
        Sub:       userID,
        TokenType: "refresh",
        RegisteredClaims: jwt.RegisteredClaims{
            IssuedAt:  jwt.NewNumericDate(now),
            ExpiresAt: jwt.NewNumericDate(now.Add(s.refreshTokenTTL)),
            ID:        uuid.New().String(),
        },
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(s.secret)
}

func (s *service) ValidateAccessToken(tokenString string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return s.secret, nil
    })
    
    if err != nil {
        return nil, err
    }
    
    claims, ok := token.Claims.(*Claims)
    if !ok || !token.Valid {
        return nil, jwt.ErrSignatureInvalid
    }
    
    return claims, nil
}

func (s *service) ValidateRefreshToken(tokenString string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return s.secret, nil
    })
    
    if err != nil {
        return nil, err
    }
    
    claims, ok := token.Claims.(*Claims)
    if !ok || !token.Valid {
        return nil, jwt.ErrSignatureInvalid
    }
    
    if claims.TokenType != "refresh" {
        return nil, jwt.ErrSignatureInvalid
    }
    
    return claims, nil
}
```

---

# SECTION 9 — Authorization Middleware

## Authentication Middleware

### Purpose

Validate JWT access token and extract user context.

### Implementation

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
        if tokenString == authHeader {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
            c.Abort()
            return
        }
        
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
        
        // Set user context
        c.Set("user_id", claims["sub"])
        c.Set("email", claims["email"])
        c.Set("name", claims["name"])
        c.Set("role", claims["role"])
        c.Set("school_id", claims["school_id"])
        
        c.Next()
    }
}
```

## Role-Based Authorization Middleware

### Purpose

Enforce role-based access control for protected endpoints.

### Implementation

```go
package middleware

import (
    "net/http"
    
    "github.com/gin-gonic/gin"
)

func RequireRole(allowedRoles ...string) gin.HandlerFunc {
    return func(c *gin.Context) {
        role := c.GetString("role")
        
        allowed := false
        for _, allowedRole := range allowedRoles {
            if role == allowedRole {
                allowed = true
                break
            }
        }
        
        if !allowed {
            c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
            c.Abort()
            return
        }
        
        c.Next()
    }
}

// Convenience functions
func RequireSystemAdmin() gin.HandlerFunc {
    return RequireRole("SYSTEM_ADMIN")
}

func RequireSchoolAdmin() gin.HandlerFunc {
    return RequireRole("SYSTEM_ADMIN", "SCHOOL_ADMIN")
}

func RequireTeacher() gin.HandlerFunc {
    return RequireRole("SYSTEM_ADMIN", "SCHOOL_ADMIN", "TEACHER")
}
```

## School Isolation Middleware

### Purpose

Ensure users can only access data from their own school.

### Implementation

```go
package middleware

import (
    "net/http"
    
    "github.com/gin-gonic/gin"
)

func RequireSchoolAccess() gin.HandlerFunc {
    return func(c *gin.Context) {
        role := c.GetString("role")
        userSchoolID := c.GetString("school_id")
        
        // SYSTEM_ADMIN can access all schools
        if role == "SYSTEM_ADMIN" {
            c.Next()
            return
        }
        
        // For SCHOOL_ADMIN and TEACHER, ensure school_id matches
        requestedSchoolID := c.Param("school_id")
        if requestedSchoolID != "" && requestedSchoolID != userSchoolID {
            c.JSON(http.StatusForbidden, gin.H{"error": "Access denied to this school"})
            c.Abort()
            return
        }
        
        c.Next()
    }
}
```

## Middleware Usage

```go
func setupRouter(cfg *config.Config, handlers ...Handler) *gin.Engine {
    router := gin.New()
    
    // Public routes
    public := router.Group("/api/v1/public")
    {
        public.POST("/login", userHandler.Login)
        public.POST("/refresh", userHandler.RefreshToken)
    }
    
    // Protected routes (require authentication)
    protected := router.Group("/api/v1")
    protected.Use(middleware.Auth(cfg.JWT.Secret))
    {
        // System admin only routes
        systemAdmin := protected.Group("/admin")
        systemAdmin.Use(middleware.RequireSystemAdmin())
        {
            systemAdmin.POST("/schools", adminHandler.CreateSchool)
            systemAdmin.GET("/schools", adminHandler.ListSchools)
        }
        
        // School admin routes
        schoolAdmin := protected.Group("/schools/:school_id")
        schoolAdmin.Use(middleware.RequireSchoolAdmin(), middleware.RequireSchoolAccess())
        {
            schoolAdmin.GET("/users", schoolHandler.ListUsers)
            schoolAdmin.POST("/users/invite", schoolHandler.InviteUser)
        }
        
        // Teacher routes
        curriculum := protected.Group("/curriculum")
        curriculum.Use(middleware.RequireTeacher())
        {
            curriculum.GET("/cp", curriculumHandler.ListCPs)
            curriculum.POST("/cp/:id/tp-sets/generate", curriculumHandler.GenerateTPSet)
        }
    }
    
    return router
}
```

---

# SECTION 10 — Permission Matrix

## Permission Matrix

| Resource | Action | SYSTEM_ADMIN | SCHOOL_ADMIN | TEACHER |
|----------|--------|-------------|--------------|---------|
| **Schools** | | | | |
| | Create | ✅ | ❌ | ❌ |
| | Read (all) | ✅ | ❌ | ❌ |
| | Read (own) | ✅ | ✅ | ✅ |
| | Update | ✅ | ✅ | ❌ |
| | Delete | ✅ | ❌ | ❌ |
| **Users** | | | | |
| | Create (invite) | ✅ | ✅ | ❌ |
| | Read (all schools) | ✅ | ❌ | ❌ |
| | Read (own school) | ✅ | ✅ | ✅ |
| | Update (all schools) | ✅ | ❌ | ❌ |
| | Update (own school) | ✅ | ✅ | ❌ |
| | Deactivate | ✅ | ✅ | ❌ |
| **Curriculum** | | | | |
| | Read CP (all schools) | ✅ | ❌ | ❌ |
| | Read CP (own school) | ✅ | ✅ | ✅ |
| | Generate TP Set | ✅ | ✅ | ✅ |
| | Approve TP Set | ✅ | ✅ | ❌ |
| | Archive TP Set | ✅ | ✅ | ✅ |
| **Learning** | | | | |
| | Generate ATP Set | ✅ | ✅ | ✅ |
| | Generate Modul Ajar Set | ✅ | ✅ | ✅ |
| | Approve ATP Set | ✅ | ✅ | ❌ |
| | Approve Modul Ajar Set | ✅ | ✅ | ❌ |
| **Assessment** | | | | |
| | Create Assessment | ✅ | ✅ | ✅ |
| | Read Assessment (own school) | ✅ | ✅ | ✅ |
| | Update Assessment (own) | ✅ | ✅ | ✅ |
| | Delete Assessment (own) | ✅ | ✅ | ✅ |
| **AI** | | | | |
| | Generate (any) | ✅ | ✅ | ✅ |
| | View Logs (all schools) | ✅ | ❌ | ❌ |
| | View Logs (own school) | ✅ | ✅ | ✅ |
| **Workflow** | | | | |
| | View History (all schools) | ✅ | ❌ | ❌ |
| | View History (own school) | ✅ | ✅ | ✅ |
| **Analytics** | | | | |
| | View System Analytics | ✅ | ❌ | ❌ |
| | View School Analytics | ✅ | ✅ | ✅ |
| | View Personal Analytics | ✅ | ✅ | ✅ |

## Permission Implementation

### Permission Checker

```go
package auth

type Permission string

const (
    PermissionSchoolCreate    Permission = "school:create"
    PermissionSchoolReadAll   Permission = "school:read_all"
    PermissionSchoolReadOwn   Permission = "school:read_own"
    PermissionSchoolUpdate    Permission = "school:update"
    PermissionSchoolDelete    Permission = "school:delete"
    
    PermissionUserCreate     Permission = "user:create"
    PermissionUserReadAll     Permission = "user:read_all"
    PermissionUserReadOwn     Permission = "user:read_own"
    PermissionUserUpdateAll   Permission = "user:update_all"
    PermissionUserUpdateOwn   Permission = "user:update_own"
    PermissionUserDeactivate  Permission = "user:deactivate"
    
    PermissionCPReadAll       Permission = "cp:read_all"
    PermissionCPReadOwn       Permission = "cp:read_own"
    PermissionTPGenerate      Permission = "tp:generate"
    PermissionTPApprove       Permission = "tp:approve"
    PermissionTPArchive       Permission = "tp:archive"
    
    PermissionATPGenerate     Permission = "atp:generate"
    PermissionATPApprove      Permission = "atp:approve"
    PermissionModulAjarGenerate Permission = "modul_ajar:generate"
    PermissionModulAjarApprove Permission = "modul_ajar:approve"
    
    PermissionAssessmentCreate Permission = "assessment:create"
    PermissionAssessmentReadOwn Permission = "assessment:read_own"
    PermissionAssessmentUpdate Permission = "assessment:update"
    PermissionAssessmentDelete Permission = "assessment:delete"
    
    PermissionAIGenerate      Permission = "ai:generate"
    PermissionAILogsReadAll   Permission = "ai:logs_read_all"
    PermissionAILogsReadOwn   Permission = "ai:logs_read_own"
    
    PermissionWorkflowReadAll Permission = "workflow:read_all"
    PermissionWorkflowReadOwn Permission = "workflow:read_own"
    
    PermissionAnalyticsSystem Permission = "analytics:system"
    PermissionAnalyticsSchool Permission = "analytics:school"
    PermissionAnalyticsPersonal Permission = "analytics:personal"
)

var rolePermissions = map[string][]Permission{
    "SYSTEM_ADMIN": {
        PermissionSchoolCreate, PermissionSchoolReadAll, PermissionSchoolReadOwn,
        PermissionSchoolUpdate, PermissionSchoolDelete,
        PermissionUserCreate, PermissionUserReadAll, PermissionUserReadOwn,
        PermissionUserUpdateAll, PermissionUserUpdateOwn, PermissionUserDeactivate,
        PermissionCPReadAll, PermissionCPReadOwn, PermissionTPGenerate,
        PermissionTPApprove, PermissionTPArchive,
        PermissionATPGenerate, PermissionATPApprove,
        PermissionModulAjarGenerate, PermissionModulAjarApprove,
        PermissionAssessmentCreate, PermissionAssessmentReadOwn,
        PermissionAssessmentUpdate, PermissionAssessmentDelete,
        PermissionAIGenerate, PermissionAILogsReadAll, PermissionAILogsReadOwn,
        PermissionWorkflowReadAll, PermissionWorkflowReadOwn,
        PermissionAnalyticsSystem, PermissionAnalyticsSchool, PermissionAnalyticsPersonal,
    },
    "SCHOOL_ADMIN": {
        PermissionSchoolReadOwn, PermissionSchoolUpdate,
        PermissionUserCreate, PermissionUserReadOwn,
        PermissionUserUpdateOwn, PermissionUserDeactivate,
        PermissionCPReadOwn, PermissionTPGenerate,
        PermissionTPApprove, PermissionTPArchive,
        PermissionATPGenerate, PermissionATPApprove,
        PermissionModulAjarGenerate, PermissionModulAjarApprove,
        PermissionAssessmentCreate, PermissionAssessmentReadOwn,
        PermissionAssessmentUpdate, PermissionAssessmentDelete,
        PermissionAIGenerate, PermissionAILogsReadOwn,
        PermissionWorkflowReadOwn,
        PermissionAnalyticsSchool, PermissionAnalyticsPersonal,
    },
    "TEACHER": {
        PermissionSchoolReadOwn,
        PermissionUserReadOwn,
        PermissionCPReadOwn, PermissionTPGenerate,
        PermissionTPArchive,
        PermissionATPGenerate,
        PermissionModulAjarGenerate,
        PermissionAssessmentCreate, PermissionAssessmentReadOwn,
        PermissionAssessmentUpdate, PermissionAssessmentDelete,
        PermissionAIGenerate, PermissionAILogsReadOwn,
        PermissionWorkflowReadOwn,
        PermissionAnalyticsPersonal,
    },
}

func HasPermission(role string, permission Permission) bool {
    permissions, ok := rolePermissions[role]
    if !ok {
        return false
    }
    
    for _, p := range permissions {
        if p == permission {
            return true
        }
    }
    
    return false
}
```

---

# SECTION 11 — Session Management Rules

## Token Lifecycle

### Access Token

- **Lifetime**: 1 hour (configurable)
- **Storage**: Client-side (localStorage, sessionStorage, or memory)
- **Validation**: Stateless, validated on each request
- **Revocation**: Not directly revocable (use short lifetime)

### Refresh Token

- **Lifetime**: 7 days (configurable)
- **Storage**: Database (refresh_tokens table)
- **Validation**: Database lookup on refresh
- **Revocation**: Database delete or mark as revoked

## Session Management Rules

### Rule 1: Single Device Login (MVP)

For MVP, users can have multiple active sessions across devices. Future enhancement may add single-device enforcement.

### Rule 2: Refresh Token Rotation

On each refresh token use:
- Delete old refresh token
- Generate new refresh token
- Client must update stored refresh token

### Rule 3: Logout Invalidates All Sessions

When user logs out:
- Delete all refresh tokens for that user
- Access tokens remain valid until expiration (short lifetime mitigates risk)

### Rule 4: Account Deactivation Invalidates Sessions

When user account is deactivated:
- Delete all refresh tokens for that user
- Access tokens remain valid until expiration

### Rule 5: Password Change Invalidates Sessions

When user changes password:
- Delete all refresh tokens for that user
- Force re-login on next request

## Refresh Token Table

### Schema

```sql
CREATE TABLE refresh_tokens (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token TEXT NOT NULL UNIQUE,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    revoked_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    last_used_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_refresh_tokens_user_id ON refresh_tokens(user_id);
CREATE INDEX idx_refresh_tokens_token ON refresh_tokens(token);
CREATE INDEX idx_refresh_tokens_expires_at ON refresh_tokens(expires_at);
```

### Model

```go
type RefreshToken struct {
    ID         string     `json:"id" db:"id"`
    UserID     string     `json:"user_id" db:"user_id"`
    Token      string     `json:"-" db:"token"` // Never exposed
    ExpiresAt  time.Time  `json:"expires_at" db:"expires_at"`
    RevokedAt  *time.Time `json:"revoked_at,omitempty" db:"revoked_at"`
    CreatedAt  time.Time  `json:"created_at" db:"created_at"`
    LastUsedAt *time.Time `json:"last_used_at,omitempty" db:"last_used_at"`
}
```

## Token Cleanup Job

### Purpose

Periodically clean up expired and revoked refresh tokens.

### Implementation

```go
func (s *service) CleanupExpiredTokens(ctx context.Context) error {
    // Delete expired tokens
    if err := s.repo.DeleteExpiredRefreshTokens(ctx); err != nil {
        return err
    }
    
    // Delete revoked tokens older than 24 hours
    if err := s.repo.DeleteRevokedRefreshTokens(ctx, 24*time.Hour); err != nil {
        return err
    }
    
    return nil
}
```

### Schedule

Run cleanup job daily using a cron job or background worker.

---

# SECTION 12 — Sequence Diagrams

## Complete Authentication Flow

```
┌─────────┐       ┌─────────┐       ┌─────────┐       ┌─────────┐
│  Client │       │  Server │       │  Database│       │  JWT     │
└────┬────┘       └────┬────┘       └────┬────┘       └────┬────┘
     │                  │                  │                  │
     │ 1. Login         │                  │                  │
     │ POST /login      │                  │                  │
     │ {email, password}│                 │                  │
     ├─────────────────>│                  │                  │
     │                  │                  │                  │
     │                  │ 2. Get user     │                  │
     │                  ├─────────────────>│                  │
     │                  │                  │                  │
     │                  │ User record     │                  │
     │                  │<─────────────────┤                  │
     │                  │                  │                  │
     │                  │ 3. Verify password                 │
     │                  │                  │                  │
     │                  │ 4. Generate access token          │
     │                  ├──────────────────────────────────>│
     │                  │                  │                  │
     │                  │ Access token    │                  │
     │                  │<──────────────────────────────────┤
     │                  │                  │                  │
     │                  │ 5. Generate refresh token         │
     │                  ├──────────────────────────────────>│
     │                  │                  │                  │
     │                  │ Refresh token   │                  │
     │                  │<──────────────────────────────────┤
     │                  │                  │                  │
     │                  │ 6. Save refresh token             │
     │                  ├─────────────────>│                  │
     │                  │                  │                  │
     │                  │ Success         │                  │
     │                  │<─────────────────┤                  │
     │                  │                  │                  │
     │ 7. Response      │                  │                  │
     │ {access_token,   │                  │                  │
     │  refresh_token,   │                  │                  │
     │  user}           │                  │                  │
     │<─────────────────┤                  │                  │
     │                  │                  │                  │
     │ 8. Use access token for API calls  │                  │
     │ GET /api/v1/curriculum/cp         │                  │
     │ Authorization: Bearer <token>      │                  │
     ├─────────────────>│                  │                  │
     │                  │                  │                  │
     │                  │ 9. Validate token                  │
     │                  ├──────────────────────────────────>│
     │                  │                  │                  │
     │                  │ Valid           │                  │
     │                  │<──────────────────────────────────┤
     │                  │                  │                  │
     │                  │ 10. Extract claims                 │
     │                  │                  │                  │
     │                  │ 11. Check permissions              │
     │                  │                  │                  │
     │                  │ 12. Query data   │                  │
     │                  ├─────────────────>│                  │
     │                  │                  │                  │
     │                  │ Data            │                  │
     │                  │<─────────────────┤                  │
     │                  │                  │                  │
     │ 13. Response     │                  │                  │
     │ {data}           │                  │                  │
     │<─────────────────┤                  │                  │
     │                  │                  │                  │
     │ 14. Access token expires (1 hour)  │                  │
     │                  │                  │                  │
     │ 15. Refresh token                 │                  │
     │ POST /refresh                     │                  │
     │ {refresh_token}                  │                  │
     ├─────────────────>│                  │                  │
     │                  │                  │                  │
     │                  │ 16. Validate refresh token        │
     │                  ├──────────────────────────────────>│
     │                  │                  │                  │
     │                  │ Valid           │                  │
     │                  │<──────────────────────────────────┤
     │                  │                  │                  │
     │                  │ 17. Get refresh token record       │
     │                  ├─────────────────>│                  │
     │                  │                  │                  │
     │                  │ Refresh record  │                  │
     │                  │<─────────────────┤                  │
     │                  │                  │                  │
     │                  │ 18. Generate new access token     │
     │                  ├──────────────────────────────────>│
     │                  │                  │                  │
     │                  │ Access token    │                  │
     │                  │<──────────────────────────────────┤
     │                  │                  │                  │
     │                  │ 19. Generate new refresh token    │
     │                  ├──────────────────────────────────>│
     │                  │                  │                  │
     │                  │ Refresh token   │                  │
     │                  │<──────────────────────────────────┤
     │                  │                  │                  │
     │                  │ 20. Delete old refresh token      │
     │                  ├─────────────────>│                  │
     │                  │                  │                  │
     │                  │ 21. Save new refresh token        │
     │                  ├─────────────────>│                  │
     │                  │                  │                  │
     │                  │ Success         │                  │
     │                  │<─────────────────┤                  │
     │                  │                  │                  │
     │ 22. Response     │                  │                  │
     │ {access_token,   │                  │                  │
     │  refresh_token}   │                  │                  │
     │<─────────────────┤                  │                  │
     │                  │                  │                  │
```

---

# SECTION 13 — API Flow

## Authentication API Endpoints

### Public Endpoints (No Authentication Required)

#### Login

```http
POST /api/v1/public/login
Content-Type: application/json

Request:
{
  "email": "user@school.com",
  "password": "password123"
}

Response (200 OK):
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "token_type": "Bearer",
  "expires_in": 3600,
  "user": {
    "id": "user-uuid",
    "email": "user@school.com",
    "name": "John Doe",
    "role": "TEACHER",
    "school_id": "school-uuid",
    "school_name": "SDN 01 Jakarta"
  }
}

Response (401 Unauthorized):
{
  "error": "Invalid credentials"
}

Response (403 Forbidden):
{
  "error": "Account is inactive"
}
```

#### Refresh Token

```http
POST /api/v1/public/refresh
Content-Type: application/json

Request:
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}

Response (200 OK):
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "token_type": "Bearer",
  "expires_in": 3600
}

Response (401 Unauthorized):
{
  "error": "Invalid or expired refresh token"
}
```

### Protected Endpoints (Authentication Required)

#### Logout

```http
POST /api/v1/auth/logout
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...

Response (204 No Content)
```

#### Get Current User

```http
GET /api/v1/auth/me
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...

Response (200 OK):
{
  "id": "user-uuid",
  "email": "user@school.com",
  "name": "John Doe",
  "role": "TEACHER",
  "school_id": "school-uuid",
  "school_name": "SDN 01 Jakarta",
  "is_active": true,
  "email_verified": true,
  "last_login_at": "2026-06-05T12:00:00Z",
  "created_at": "2026-01-01T00:00:00Z"
}
```

## Authorization Flow

### Request with Authorization

```http
GET /api/v1/curriculum/cp
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

### Middleware Processing

1. **Auth Middleware**
   - Extract token from Authorization header
   - Validate token signature and expiration
   - Extract claims (user_id, role, school_id)
   - Set context values

2. **Role Middleware** (if applicable)
   - Check user role against allowed roles
   - Allow or deny based on role

3. **School Isolation Middleware** (if applicable)
   - Check user school_id against requested school_id
   - Allow or deny based on school access

4. **Handler**
   - Access user context from Gin context
   - Perform business logic
   - Return response

---

# SECTION 14 — JWT Structure

## JWT Header

```json
{
  "alg": "HS256",
  "typ": "JWT"
}
```

## JWT Payload (Access Token)

```json
{
  "sub": "550e8400-e29b-41d4-a716-446655440000",
  "email": "teacher@school.com",
  "name": "John Doe",
  "role": "TEACHER",
  "school_id": "660e8400-e29b-41d4-a716-446655440001",
  "iat": 1717658400,
  "exp": 1717662000,
  "jti": "770e8400-e29b-41d4-a716-446655440002"
}
```

## JWT Payload (Refresh Token)

```json
{
  "sub": "550e8400-e29b-41d4-a716-446655440000",
  "token_type": "refresh",
  "iat": 1717658400,
  "exp": 1718263200,
  "jti": "880e8400-e29b-41d4-a716-446655440003"
}
```

## JWT Signature

HMACSHA256(
  base64UrlEncode(header) + "." + base64UrlEncode(payload),
  secret
)

## Token Validation

### Validation Steps

1. **Structure Check**: Verify JWT has three parts (header, payload, signature)
2. **Algorithm Check**: Verify algorithm is HS256
3. **Signature Check**: Verify signature using secret
4. **Expiration Check**: Verify `exp` claim is in the future
5. **Not Before Check**: Verify `nbf` claim (if present) is in the past
6. **Issued At Check**: Verify `iat` claim is reasonable
7. **Token Type Check**: Verify `token_type` claim (for refresh tokens)

### Error Handling

| Error | HTTP Status | Message |
|-------|------------|---------|
| Missing token | 401 | Authorization header required |
| Invalid token format | 401 | Invalid authorization header format |
| Invalid signature | 401 | Invalid token |
| Expired token | 401 | Token expired |
| Invalid claims | 401 | Invalid token claims |
| Wrong token type | 401 | Invalid token type |

---

# SECTION 15 — Database Mapping

## Database Schema

### Users Table

```sql
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL CHECK (role IN ('SYSTEM_ADMIN', 'SCHOOL_ADMIN', 'TEACHER')),
    school_id UUID REFERENCES schools(id) ON DELETE SET NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    email_verified BOOLEAN NOT NULL DEFAULT FALSE,
    last_login_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    
    CONSTRAINT chk_users_school_role CHECK (
        (role = 'SYSTEM_ADMIN' AND school_id IS NULL) OR
        (role IN ('SCHOOL_ADMIN', 'TEACHER') AND school_id IS NOT NULL)
    )
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_school_id ON users(school_id);
CREATE INDEX idx_users_role ON users(role);
CREATE INDEX idx_users_is_active ON users(is_active);
```

### Schools Table

```sql
CREATE TABLE schools (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    name VARCHAR(255) NOT NULL,
    code VARCHAR(50) NOT NULL UNIQUE,
    address TEXT,
    phone VARCHAR(50),
    email VARCHAR(255),
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_schools_code ON schools(code);
CREATE INDEX idx_schools_is_active ON schools(is_active);
```

### Refresh Tokens Table

```sql
CREATE TABLE refresh_tokens (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token TEXT NOT NULL UNIQUE,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    revoked_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    last_used_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_refresh_tokens_user_id ON refresh_tokens(user_id);
CREATE INDEX idx_refresh_tokens_token ON refresh_tokens(token);
CREATE INDEX idx_refresh_tokens_expires_at ON refresh_tokens(expires_at);
```

## Entity Relationships

```
┌─────────────┐
│   schools   │
│─────────────│
│ id (PK)     │
│ name        │
│ code        │
│ ...         │
└──────┬──────┘
       │ 1
       │
       │ N
┌──────┴──────┐
│    users    │
│─────────────│
│ id (PK)     │
│ email       │
│ password_hash│
│ name        │
│ role        │
│ school_id (FK)│
│ ...         │
└──────┬──────┘
       │ 1
       │
       │ N
┌──────┴──────┐
│refresh_tokens│
│─────────────│
│ id (PK)     │
│ user_id (FK)│
│ token       │
│ expires_at  │
│ ...         │
└─────────────┘
```

## Data Access Layer

### Repository Interface

```go
package repository

type UserRepository interface {
    CreateUser(ctx context.Context, user *User) error
    GetUser(ctx context.Context, id string) (*User, error)
    GetUserByEmail(ctx context.Context, email string) (*User, error)
    UpdateUser(ctx context.Context, user *User) error
    ListUsersBySchool(ctx context.Context, schoolID string) ([]*User, error)
    DeactivateUser(ctx context.Context, id string) error
}

type SchoolRepository interface {
    CreateSchool(ctx context.Context, school *School) error
    GetSchool(ctx context.Context, id string) (*School, error)
    GetSchoolByCode(ctx context.Context, code string) (*School, error)
    ListSchools(ctx context.Context) ([]*School, error)
    UpdateSchool(ctx context.Context, school *School) error
    DeactivateSchool(ctx context.Context, id string) error
}

type RefreshTokenRepository interface {
    CreateRefreshToken(ctx context.Context, token *RefreshToken) error
    GetRefreshToken(ctx context.Context, token string) (*RefreshToken, error)
    DeleteRefreshToken(ctx context.Context, id string) error
    DeleteRefreshTokensByUserID(ctx context.Context, userID string) error
    DeleteExpiredRefreshTokens(ctx context.Context) error
    DeleteRevokedRefreshTokens(ctx context.Context, olderThan time.Duration) error
}
```

---

# SECTION 16 — Appendix

## Security Considerations

### Password Security

- Use bcrypt with cost factor 12
- Never store plain text passwords
- Never expose password hash in API responses
- Enforce password complexity (minimum 8 characters, mixed case, numbers)

### Token Security

- Use strong secret key (minimum 32 characters)
- Store secret in environment variable
- Rotate secret key periodically (future enhancement)
- Use HTTPS in production

### Session Security

- Short access token lifetime (1 hour)
- Refresh token rotation on each use
- Logout invalidates all refresh tokens
- Account deactivation invalidates all sessions

### Rate Limiting

- Implement rate limiting on login endpoint (future enhancement)
- Implement rate limiting on refresh token endpoint (future enhancement)

## Configuration

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| JWT_SECRET | JWT signing secret | (required) |
| JWT_ACCESS_TOKEN_TTL | Access token lifetime | 1h |
| JWT_REFRESH_TOKEN_TTL | Refresh token lifetime | 168h (7 days) |
| BCRYPT_COST | Bcrypt cost factor | 12 |

## Testing

### Unit Tests

- Test password hashing and verification
- Test JWT generation and validation
- Test permission checking
- Test repository operations

### Integration Tests

- Test complete login flow
- Test refresh token flow
- Test logout flow
- Test authorization middleware

### Security Tests

- Test SQL injection protection
- Test token tampering
- Test expired token handling
- Test revoked token handling

## Future Enhancements

### Wave 2

- Email verification flow
- Password reset flow
- Multi-factor authentication (MFA)
- Single sign-on (SSO) integration
- OAuth/OIDC provider support
- Device management
- Session history
- Audit logging for authentication events

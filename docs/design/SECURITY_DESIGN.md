# Security Design

**Version:** 1.0  
**Date:** June 13, 2026  
**Based On:** AUDIT_REPORT.md, TECHNICAL_DEBT.md, TARGET_ARCHITECTURE.md

---

## Executive Summary

This document describes the security design for the NUSA Platform, including authentication, authorization, RBAC, audit logging, data protection, and API security. The design prioritizes security best practices while maintaining simplicity and avoiding overengineering.

**Security Principles:**
- Defense in depth
- Least privilege
- Zero trust
- Security by design
- Fail securely

---

## Authentication

### Authentication Flow

**Current Implementation:**
- JWT-based authentication
- Access tokens (24h expiration)
- Refresh tokens (30-day expiration)
- Refresh token rotation

**Enhanced Design:**

#### Token Management

**Access Token:**
- Type: JWT (HS256)
- Expiration: 24 hours
- Claims: user_id, role, school_id, permissions, iat, exp
- Storage: Client-side (localStorage for web, secure storage for mobile)

**Refresh Token:**
- Type: Random UUID
- Expiration: 30 days
- Storage: Database (refresh_tokens table)
- Attributes: user_id, token_hash, expires_at, client_ip, user_agent, revoked_at

**Token Rotation:**
- Rotate refresh token on every refresh
- Invalidate old refresh token
- Track refresh token usage
- Detect token reuse attacks

#### Authentication Endpoints

**POST /public/auth/login**
- Request: email, password
- Validation: Email format, password strength
- Rate limiting: 10 attempts per minute per IP
- Account lockout: 5 failed attempts, 15-minute lockout
- Response: access_token, refresh_token, user info

**POST /public/auth/refresh**
- Request: refresh_token
- Validation: Token exists, not expired, not revoked
- Rate limiting: 30 attempts per minute per IP
- Response: new access_token, new refresh_token

**POST /auth/logout**
- Request: refresh_token
- Action: Revoke refresh_token
- Response: Success message

#### Password Security

**Password Requirements:**
- Minimum 8 characters
- At least one uppercase letter
- At least one lowercase letter
- At least one number
- At least one special character
- No common passwords

**Password Hashing:**
- Algorithm: bcrypt
- Cost factor: 12
- Salt: Automatic per password

**Password Reset:**
- Token-based reset
- Token expiration: 1 hour
- Single-use token
- Email delivery

#### Session Management

**Session Timeout:**
- Inactivity timeout: 8 hours
- Absolute timeout: 24 hours
- Refresh on activity
- Warning before timeout (5 minutes)

**Session Revocation:**
- Manual logout
- Password change
- Role change
- School change
- Admin action

---

## Authorization

### Authorization Model

**Role-Based Access Control (RBAC):**
- Users assigned to roles
- Roles assigned to permissions
- Permissions grant access to resources

**Permission-Based Access Control (PBAC):**
- Fine-grained permissions
- Resource-action format
- Dynamic permission checks

### Roles

**System Roles:**
- `SYSTEM_ADMIN` - Full system access
- `SCHOOL_ADMIN` - School-level administration
- `CURRICULUM_ADMIN` - Curriculum management
- `TEACHER` - Teaching and assessment
- `STUDENT` - Student access (future)

**Role Hierarchy:**
```
SYSTEM_ADMIN
  ├── SCHOOL_ADMIN
  ├── CURRICULUM_ADMIN
  └── TEACHER
      └── STUDENT
```

### Permissions

**Permission Format:** `{resource}:{action}`

**Resource Types:**
- user
- school
- academic_year
- semester
- subject_category
- graduate_profile
- cp_alignment
- system_config
- curriculum_subject
- curriculum_phase
- curriculum_element
- curriculum_subelement
- cp
- tp
- atp
- modul_ajar
- assessment
- rubric
- evidence
- evaluation
- class
- attendance
- schedule
- exam
- assignment
- notification
- announcement
- message
- audit_log

**Actions:**
- CREATE
- READ
- UPDATE
- DELETE
- APPROVE
- ARCHIVE
- ACTIVATE
- EXPORT
- IMPORT

**Permission Examples:**
- `user:CREATE` - Create users
- `academic_year:ACTIVATE` - Activate academic year
- `assessment:APPROVE` - Approve assessments
- `class:UPDATE` - Update classes

### Permission Matrix

| Role | User | School | Academic Year | Curriculum | TP | Assessment | Class | Attendance |
|------|------|--------|---------------|------------|-----|------------|-------|------------|
| SYSTEM_ADMIN | CRUD | CRUD | CRUD | CRUD | CRUD | CRUD | CRUD | CRUD |
| SCHOOL_ADMIN | CRUD | CRUD | CRUD | CRUD | READ | READ | CRUD | CRUD |
| CURRICULUM_ADMIN | READ | READ | READ | CRUD | CRUD | READ | READ | READ |
| TEACHER | READ | READ | READ | READ | CRUD | CRUD | READ | CRUD |
| STUDENT | READ | READ | READ | READ | READ | READ | READ | READ |

---

## RBAC Implementation

### Middleware Design

**AuthMiddleware:**
```go
func AuthMiddleware(jwtService *JWTService) gin.HandlerFunc {
    return func(c *gin.Context) {
        token := extractToken(c)
        if token == "" {
            c.JSON(401, gin.H{"error": "Missing token"})
            c.Abort()
            return
        }
        
        claims, err := jwtService.ValidateToken(token)
        if err != nil {
            c.JSON(401, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }
        
        c.Set("user_id", claims.UserID)
        c.Set("role", claims.Role)
        c.Set("school_id", claims.SchoolID)
        c.Set("permissions", claims.Permissions)
        
        c.Next()
    }
}
```

**RequirePermissionMiddleware:**
```go
func RequirePermission(permission string) gin.HandlerFunc {
    return func(c *gin.Context) {
        permissions := c.GetStringSlice("permissions")
        if !contains(permissions, permission) {
            c.JSON(403, gin.H{"error": "Permission denied"})
            c.Abort()
            return
        }
        c.Next()
    }
}
```

**RequireRoleMiddleware:**
```go
func RequireRole(roles ...string) gin.HandlerFunc {
    return func(c *gin.Context) {
        userRole := c.GetString("role")
        if !contains(roles, userRole) {
            c.JSON(403, gin.H{"error": "Role not authorized"})
            c.Abort()
            return
        }
        c.Next()
    }
}
```

**RequireSchoolAccessMiddleware:**
```go
func RequireSchoolAccess() gin.HandlerFunc {
    return func(c *gin.Context) {
        userSchoolID := c.GetString("school_id")
        resourceSchoolID := c.Param("school_id")
        
        if userSchoolID != "" && resourceSchoolID != "" {
            if userSchoolID != resourceSchoolID {
                c.JSON(403, gin.H{"error": "School access denied"})
                c.Abort()
                return
            }
        }
        
        c.Next()
    }
}
```

### Permission Checks

**Service Layer:**
```go
func (s *UserService) CanCreateUser(ctx context.Context, userID string, targetSchoolID string) error {
    user, err := s.userRepo.GetByID(ctx, userID)
    if err != nil {
        return err
    }
    
    // SYSTEM_ADMIN can create users anywhere
    if user.Role == "SYSTEM_ADMIN" {
        return nil
    }
    
    // SCHOOL_ADMIN can only create users in their school
    if user.Role == "SCHOOL_ADMIN" {
        if user.SchoolID != targetSchoolID {
            return errors.New("cannot create user in different school")
        }
        return nil
    }
    
    return errors.New("insufficient permissions")
}
```

---

## Audit Logging

### Audit Log Design

**Audit Log Structure:**
```go
type AuditLog struct {
    ID          string
    UserID      *string
    Action      string
    EntityType  string
    EntityID    *string
    OldValues   map[string]interface{}
    NewValues   map[string]interface{}
    IPAddress   string
    UserAgent   string
    RequestID   string
    CreatedAt   time.Time
}
```

**Audit Middleware:**
```go
func AuditMiddleware(auditService *AuditService) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Capture request details
        userID := c.GetString("user_id")
        action := c.Request.Method + " " + c.Request.URL.Path
        requestID := c.GetHeader("X-Request-ID")
        
        // Get old values for UPDATE/DELETE
        var oldValues map[string]interface{}
        if c.Request.Method != "POST" {
            oldValues = getOldValues(c)
        }
        
        // Execute handler
        c.Next()
        
        // Get new values for POST/PUT
        var newValues map[string]interface{}
        if c.Request.Method != "DELETE" {
            newValues = getNewValues(c)
        }
        
        // Create audit log
        audit := &domain.AuditLog{
            UserID:     &userID,
            Action:     action,
            EntityType: getEntityType(c),
            EntityID:   getEntityID(c),
            OldValues:  oldValues,
            NewValues:  newValues,
            IPAddress:  c.ClientIP(),
            UserAgent:  c.GetHeader("User-Agent"),
            RequestID:  requestID,
            CreatedAt:  time.Now(),
        }
        
        auditService.Create(c, audit)
    }
}
```

**Audit Events:**
- User creation, update, deletion
- Role assignment changes
- School creation, update, deletion
- Academic year activation, archiving
- Class creation, enrollment changes
- Assessment approval
- Configuration changes
- Permission changes

**Audit Retention:**
- Retention period: 2 years
- Archive after 6 months
- Delete after 2 years

---

## Data Protection

### Data Encryption

**In Transit:**
- TLS 1.3 for all HTTP traffic
- Certificate validation
- HSTS header
- No insecure HTTP

**At Rest:**
- Passwords: bcrypt hashing
- PII: Consider encryption for sensitive fields (email, phone, address)
- Database: Transparent Data Encryption (TDE) if supported
- File storage: S3 server-side encryption

### Data Masking

**PII Fields:**
- Email: Mask in logs (user***@example.com)
- Phone: Mask in logs (***-***-1234)
- Address: Mask in logs (*** Street)
- Names: Mask in logs (J*** D***)

**Audit Logs:**
- Old values: Mask PII
- New values: Mask PII
- User ID: Keep for audit trail

### Data Minimization

**Collect Only Necessary Data:**
- User: Name, email, role, school (minimal PII)
- Student: Name, email, class enrollment (minimal PII)
- No unnecessary personal information

**Data Retention:**
- User data: Retain while active + 3 years after inactive
- Audit logs: 2 years
- Notifications: 30 days
- Messages: 1 year

---

## API Security

### Rate Limiting

**Rate Limiting Strategy:**
- Per user: 100 requests/minute
- Per IP: 200 requests/minute
- Per endpoint: Custom limits

**Implementation:**
```go
func RateLimitMiddleware(redis *redis.Client) gin.HandlerFunc {
    return func(c *gin.Context) {
        userID := c.GetString("user_id")
        key := fmt.Sprintf("ratelimit:%s", userID)
        
        count, err := redis.Incr(c, key).Result()
        if err != nil {
            c.Next()
            return
        }
        
        if count == 1 {
            redis.Expire(c, key, time.Minute)
        }
        
        if count > 100 {
            c.JSON(429, gin.H{
                "error": "Rate limit exceeded",
                "retry_after": 60,
            })
            c.Abort()
            return
        }
        
        c.Next()
    }
}
```

**Rate Limit Headers:**
```
X-RateLimit-Limit: 100
X-RateLimit-Remaining: 95
X-RateLimit-Reset: 1623542400
```

### Security Headers

**Required Headers:**
```
X-Content-Type-Options: nosniff
X-Frame-Options: DENY
X-XSS-Protection: 1; mode=block
Strict-Transport-Security: max-age=31536000; includeSubDomains
Content-Security-Policy: default-src 'self'; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline'
Referrer-Policy: strict-origin-when-cross-origin
Permissions-Policy: geolocation=(), microphone=(), camera=()
```

**Implementation:**
```go
func SecurityHeadersMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Header("X-Content-Type-Options", "nosniff")
        c.Header("X-Frame-Options", "DENY")
        c.Header("X-XSS-Protection", "1; mode=block")
        c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
        c.Header("Content-Security-Policy", "default-src 'self'")
        c.Header("Referrer-Policy", "strict-origin-when-cross-origin")
        c.Next()
    }
}
```

### CORS Configuration

**CORS Policy:**
- Allow origins: Configured whitelist
- Allow methods: GET, POST, PUT, DELETE, PATCH, OPTIONS
- Allow headers: Authorization, Content-Type, X-Request-ID
- Max age: 1 hour
- Credentials: Supported

**Implementation:**
```go
func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        origin := c.GetHeader("Origin")
        
        if isAllowedOrigin(origin) {
            c.Header("Access-Control-Allow-Origin", origin)
            c.Header("Access-Control-Allow-Credentials", "true")
            c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
            c.Header("Access-Control-Allow-Headers", "Authorization, Content-Type, X-Request-ID")
            c.Header("Access-Control-Max-Age", "3600")
        }
        
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        
        c.Next()
    }
}
```

### Input Validation

**Validation Strategy:**
- Schema validation for all requests
- Type checking
- Length limits
- Format validation
- SQL injection prevention (parameterized queries)
- XSS prevention (input sanitization)

**Implementation:**
```go
func ValidationMiddleware(schema interface{}) gin.HandlerFunc {
    return func(c *gin.Context) {
        if err := c.ShouldBindJSON(schema); err != nil {
            c.JSON(400, gin.H{
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

### Output Encoding

**JSON Encoding:**
- All responses as JSON
- Escape HTML entities
- No raw HTML in responses
- Content-Type: application/json

**Sanitization:**
- Sanitize user-generated content
- Strip HTML from text fields
- Escape special characters

---

## Infrastructure Security

### Database Security

**Connection Security:**
- TLS for database connections
- Connection string encryption
- No plaintext passwords

**Access Control:**
- Least privilege database user
- Separate read/write users
- No direct database access from internet

**Backup Security:**
- Encrypted backups
- Secure backup storage
- Backup access logging

### Redis Security

**Connection Security:**
- TLS for Redis connections
- AUTH password
- No plaintext passwords

**Access Control:**
- Redis AUTH enabled
- No public access
- Firewall rules

### RabbitMQ Security

**Connection Security:**
- TLS for AMQP connections
- SASL authentication
- Separate vhost per environment

**Access Control:**
- User permissions per queue
- No guest access
- Firewall rules

### File Storage Security

**S3 Security:**
- Bucket policies (least privilege)
- Presigned URLs for uploads
- Server-side encryption
- Versioning enabled

**File Validation:**
- File type validation
- File size limits
- Virus scanning (optional)

---

## Monitoring & Alerting

### Security Monitoring

**Metrics to Monitor:**
- Failed login attempts
- Rate limit violations
- Permission denials
- Unusual API usage patterns
- Audit log anomalies

**Alerting:**
- High failed login rate
- Brute force attack detected
- Unauthorized access attempts
- Data breach indicators

### Log Monitoring

**Security Events to Log:**
- Authentication failures
- Authorization failures
- Permission changes
- Role changes
- Configuration changes
- Data access patterns

**Log Analysis:**
- SIEM integration
- Anomaly detection
- Pattern recognition
- Threat intelligence

---

## Compliance

### Data Protection

**GDPR-like Principles:**
- Lawful basis for processing
- Data minimization
- Purpose limitation
- Data accuracy
- Storage limitation
- Integrity and confidentiality
- Accountability

**User Rights:**
- Right to access
- Right to rectification
- Right to erasure
- Right to restrict processing
- Right to data portability
- Right to object

### Audit Trail

**Comprehensive Logging:**
- All data mutations logged
- User actions tracked
- System changes recorded
- Access attempts monitored

**Audit Report:**
- User activity reports
- System change reports
- Access reports
- Compliance reports

---

## Security Best Practices

### Development

**Secure Coding:**
- Input validation
- Output encoding
- Parameterized queries
- No hardcoded secrets
- Dependency scanning

**Code Review:**
- Security review for all changes
- Penetration testing
- Static analysis
- Dependency vulnerability scanning

### Deployment

**Secure Deployment:**
- HTTPS only
- Security headers
- Firewall rules
- Network segmentation
- Regular updates

**Secrets Management:**
- Environment variables
- No secrets in code
- Secret rotation
- Access logging

### Operations

**Secure Operations:**
- Least privilege access
- Multi-factor authentication
- Access logging
- Regular audits
- Incident response plan

---

## Incident Response

### Incident Response Plan

**Detection:**
- Monitoring alerts
- User reports
- Automated detection

**Containment:**
- Isolate affected systems
- Revoke compromised credentials
- Block malicious IPs

**Eradication:**
- Remove malware
- Patch vulnerabilities
- Update configurations

**Recovery:**
- Restore from backups
- Verify system integrity
- Monitor for recurrence

**Lessons Learned:**
- Post-incident review
- Update procedures
- Improve monitoring

---

## Security Checklist

### Pre-Deployment

- [ ] TLS 1.3 enabled
- [ ] Security headers configured
- [ ] Rate limiting enabled
- [ ] Input validation implemented
- [ ] Output encoding implemented
- [ ] CORS configured
- [ ] Database encryption enabled
- [ ] Redis AUTH enabled
- [ ] RabbitMQ TLS enabled
- [ ] File storage encryption enabled
- [ ] Audit logging enabled
- [ ] Error tracking enabled
- [ ] Secrets management configured
- [ ] Firewall rules configured
- [ ] Security testing completed

### Post-Deployment

- [ ] Security headers verified
- [ ] Rate limiting tested
- [ ] Authentication tested
- [ ] Authorization tested
- [ ] Audit logs verified
- [ ] Monitoring configured
- [ ] Alerting configured
- [ ] Backup verification
- [ ] Penetration testing
- [ ] Security review

---

## Conclusion

The security design provides comprehensive coverage of authentication, authorization, RBAC, audit logging, data protection, and API security. The design prioritizes security best practices while maintaining simplicity and avoiding overengineering.

Key security features include JWT-based authentication with refresh token rotation, comprehensive RBAC with permission-based access control, detailed audit logging, rate limiting, security headers, and data encryption. The design ensures the NUSA Platform meets security requirements for educational data while maintaining usability and performance.

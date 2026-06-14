# Technical Debt Report

**Generated:** June 13, 2026  
**Scope:** NUSA Platform - Full Stack Technical Debt Assessment  
**Version:** 1.0

---

## Executive Summary

This technical debt report identifies areas of the codebase that require remediation to maintain long-term maintainability, performance, and scalability. The assessment covers code quality, architecture, testing, documentation, security, and operational concerns.

**Key Metrics:**
- **Total Debt Items:** 35
- **Critical Debt:** 8 (Immediate attention required)
- **High Debt:** 12 (Should address soon)
- **Medium Debt:** 10 (Plan to address)
- **Low Debt:** 5 (Nice to have)
- **Estimated Remediation Effort:** 8-12 weeks

---

## Critical Debt (Immediate Attention Required)

### 1. Missing Test Coverage

**Location:** Backend & Frontend  
**Severity:** Critical  
**Impact:** High risk of regressions, difficult refactoring, lower code quality

**Description:**
- No unit tests for backend services, repositories, or handlers
- No unit tests for frontend components or hooks
- No integration tests for API endpoints
- No E2E tests for critical user flows

**Remediation:**
```bash
# Backend testing setup
cd backend
go test ./... -v
go test ./internal/domain -cover
go test ./internal/service -cover
go test ./internal/repository -cover

# Frontend testing setup
cd frontend
npm test -- --coverage
npm run test:e2e
```

**Estimated Effort:** 3-4 weeks

---

### 2. No Error Tracking & Monitoring

**Location:** Backend & Frontend  
**Severity:** Critical  
**Impact:** Difficult debugging, poor incident response, limited visibility

**Description:**
- No error tracking (Sentry, Rollbar)
- No application performance monitoring (APM)
- No logging aggregation
- No alerting system

**Remediation:**
```go
// Add Sentry to backend
import "github.com/getsentry/sentry-go"

sentry.Init(sentry.ClientOptions{
    Dsn: os.Getenv("SENTRY_DSN"),
    Environment: os.Getenv("APP_ENV"),
})
```

```typescript
// Add Sentry to frontend
import * as Sentry from "@sentry/react";

Sentry.init({
  dsn: import.meta.env.VITE_SENTRY_DSN,
  environment: import.meta.env.MODE,
});
```

**Estimated Effort:** 1-2 weeks

---

### 3. No Database Backup Strategy

**Location:** Database  
**Severity:** Critical  
**Impact:** Risk of data loss, no disaster recovery, compliance concerns

**Description:**
- No automated backups
- No backup verification
- No point-in-time recovery
- No backup restoration testing

**Remediation:**
```bash
# Add automated backup script
pg_dump -h $DB_HOST -U $DB_USER -d $DB_NAME > backup_$(date +%Y%m%d).sql

# Add to cron job
0 2 * * * /path/to/backup-script.sh

# Add backup verification
pg_restore --list backup.sql | wc -l
```

**Estimated Effort:** 1 week

---

### 4. No Rate Limiting

**Location:** Backend API  
**Severity:** Critical  
**Impact:** Vulnerable to abuse, DDoS risk, no fair usage policy

**Description:**
- No API rate limiting
- No request throttling
- No abuse detection
- No IP-based blocking

**Remediation:**
```go
// Add rate limiting middleware
import "github.com/ulule/limiter/v3"

rate := limiter.Rate{
    Period: time.Hour,
    Limit:  1000,
}

store := memory.NewStore()
instance := limiter.New(store, rate)
```

**Estimated Effort:** 1 week

---

### 5. No Audit Logging

**Location:** Backend  
**Severity:** Critical  
**Impact:** Limited accountability, compliance concerns, difficult change tracking

**Description:**
- No comprehensive audit log
- No change history tracking
- No user action logging
- No audit report generation

**Remediation:**
```sql
CREATE TABLE audit_logs (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID REFERENCES users(id),
  action VARCHAR(100) NOT NULL,
  entity_type VARCHAR(100) NOT NULL,
  entity_id UUID,
  old_values JSONB,
  new_values JSONB,
  ip_address INET,
  user_agent TEXT,
  created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_audit_logs_user_id ON audit_logs(user_id);
CREATE INDEX idx_audit_logs_entity ON audit_logs(entity_type, entity_id);
CREATE INDEX idx_audit_logs_created_at ON audit_logs(created_at DESC);
```

**Estimated Effort:** 2 weeks

---

### 6. No Input Validation Layer

**Location:** Backend  
**Severity:** Critical  
**Impact:** Poor data quality, security risks, unclear error messages

**Description:**
- Basic validation in DTOs only
- No comprehensive validation rules
- No custom validators
- Limited error messages

**Remediation:**
```go
// Add validation library
import "github.com/go-playground/validator/v10"

validate := validator.New()
err := validate.Struct(request)
```

**Estimated Effort:** 1-2 weeks

---

### 7. No Security Headers

**Location:** Backend  
**Severity:** Critical  
**Impact:** Security vulnerabilities, compliance concerns

**Description:**
- No security headers (CSP, HSTS, X-Frame-Options)
- No CORS configuration
- No XSS protection headers
- No content-type validation

**Remediation:**
```go
// Add security headers middleware
c.Header("X-Content-Type-Options", "nosniff")
c.Header("X-Frame-Options", "DENY")
c.Header("X-XSS-Protection", "1; mode=block")
c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
c.Header("Content-Security-Policy", "default-src 'self'")
```

**Estimated Effort:** 3 days

---

### 8. No File Storage Strategy

**Location:** Backend & Database  
**Severity:** Critical  
**Impact:** Scalability concerns, no file backup, poor performance

**Description:**
- File URLs stored in database
- No file storage service integration
- No CDN integration
- No file versioning

**Remediation:**
```go
// Add S3 integration
import "github.com/aws/aws-sdk-go-v2/service/s3"

client := s3.New(s3.Options{
    Region: os.Getenv("AWS_REGION"),
    Credentials: cred.NewStaticCredentialsProvider(
        os.Getenv("AWS_ACCESS_KEY"),
        os.Getenv("AWS_SECRET_KEY"),
    ),
})
```

**Estimated Effort:** 2 weeks

---

## High Debt (Should Address Soon)

### 9. No Caching Layer

**Location:** Backend  
**Severity:** High  
**Impact:** Increased database load, slower response times, poor scalability

**Description:**
- TanStack Query provides client-side caching only
- No server-side caching (Redis)
- No cache invalidation strategy
- No cache warming

**Remediation:**
```go
// Add Redis caching
import "github.com/go-redis/redis/v8"

client := redis.NewClient(&redis.Options{
    Addr: os.Getenv("REDIS_ADDR"),
    Password: os.Getenv("REDIS_PASSWORD"),
    DB: 0,
})
```

**Estimated Effort:** 2-3 weeks

---

### 10. Limited Error Handling

**Location:** Backend & Frontend  
**Severity:** High  
**Impact:** Poor user experience, difficult debugging, limited visibility

**Description:**
- Basic error handling exists
- No error recovery mechanisms
- No error reporting UI
- Limited error context

**Remediation:**
```go
// Add comprehensive error handling
type AppError struct {
    Code    string
    Message string
    Context map[string]interface{}
    Cause   error
}

func (e *AppError) Error() string {
    return e.Message
}
```

**Estimated Effort:** 1-2 weeks

---

### 11. No Database Indexes for Common Queries

**Location:** Database  
**Severity:** High  
**Impact:** Slow queries, poor performance, database load

**Description:**
- Basic foreign key indexes only
- No composite indexes
- No full-text search indexes
- No partial indexes

**Remediation:**
```sql
-- Add composite indexes
CREATE INDEX idx_academic_years_school_status ON academic_years(school_id, status);
CREATE INDEX idx_semesters_academic_year_status ON semesters(academic_year_id, status);
CREATE INDEX idx_tps_subject_phase ON tps(subject_id, phase_id);

-- Add partial indexes
CREATE INDEX idx_active_users ON users(is_active) WHERE is_active = true;

-- Add full-text search indexes
CREATE INDEX idx_cp_text_search ON cps USING gin(to_tsvector('english', description));
```

**Estimated Effort:** 1 week

---

### 12. No Soft Delete Implementation

**Location:** Database  
**Severity:** High  
**Impact:** Data loss risk, no recovery, audit trail gaps

**Description:**
- Hard delete only
- No `deleted_at` timestamp
- No restore functionality
- No permanent delete workflow

**Remediation:**
```sql
-- Add deleted_at to all tables
ALTER TABLE users ADD COLUMN deleted_at TIMESTAMP;
ALTER TABLE academic_years ADD COLUMN deleted_at TIMESTAMP;
ALTER TABLE semesters ADD COLUMN deleted_at TIMESTAMP;

-- Add partial indexes
CREATE INDEX idx_users_not_deleted ON users(id) WHERE deleted_at IS NULL;
```

**Estimated Effort:** 2 weeks

---

### 13. No Database Connection Pooling Configuration

**Location:** Backend  
**Severity:** High  
**Impact:** Connection exhaustion, poor performance, scalability issues

**Description:**
- Default pgxpool configuration
- No connection limit tuning
- No connection timeout configuration
- No connection health checks

**Remediation:**
```go
// Configure connection pool
config, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
config.MaxConns = 25
config.MinConns = 5
config.MaxConnLifetime = time.Hour
config.MaxConnIdleTime = 30 * time.Minute
config.HealthCheckPeriod = 1 * time.Minute
```

**Estimated Effort:** 3 days

---

### 14. No API Versioning Strategy

**Location:** Backend API  
**Severity:** High  
**Impact:** Breaking changes, backward compatibility issues, migration complexity

**Description:**
- Single API version (`/api/v1`)
- No version negotiation
- No deprecation strategy
- No migration guide

**Remediation:**
```go
// Add version negotiation
func getVersion(c *gin.Context) string {
    version := c.GetHeader("API-Version")
    if version == "" {
        version = "v1"
    }
    return version
}
```

**Estimated Effort:** 2 weeks

---

### 15. No Request Validation Middleware

**Location:** Backend  
**Severity:** High  
**Impact:** Security risks, poor data quality, unclear error messages

**Description:**
- Validation scattered in handlers
- No centralized validation
- No schema validation
- No request size limits

**Remediation:**
```go
// Add validation middleware
func ValidationMiddleware(schema interface{}) gin.HandlerFunc {
    return func(c *gin.Context) {
        if err := c.ShouldBindJSON(schema); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            c.Abort()
            return
        }
        c.Next()
    }
}
```

**Estimated Effort:** 1 week

---

### 16. No Pagination Consistency

**Location:** Backend API  
**Severity:** High  
**Impact:** Inconsistent API responses, poor UX, performance issues

**Description:**
- Some endpoints use `page/page_size`
- Some use `limit/offset`
- No default pagination
- No maximum limits

**Remediation:**
```go
// Standardize pagination
type PaginationRequest struct {
    Page     int `form:"page" binding:"min=1"`
    PageSize int `form:"page_size" binding:"min=1,max=100"`
}

func (p *PaginationRequest) Defaults() {
    if p.Page == 0 {
        p.Page = 1
    }
    if p.PageSize == 0 {
        p.PageSize = 20
    }
}
```

**Estimated Effort:** 1 week

---

### 17. No Background Job Processing

**Location:** Backend  
**Severity:** High  
**Impact:** Poor UX for long operations, timeout issues, no job tracking

**Description:**
- All operations synchronous
- No job queue
- No job status tracking
- No retry mechanism

**Remediation:**
```go
// Add job queue (using existing RabbitMQ)
type Job struct {
    ID        string
    Type      string
    Payload   interface{}
    Status    string
    CreatedAt time.Time
    UpdatedAt time.Time
}

func (j *JobQueue) Enqueue(job *Job) error {
    // Enqueue job to RabbitMQ
}
```

**Estimated Effort:** 3 weeks

---

### 18. No Database Migration Rollback Testing

**Location:** Database Migrations  
**Severity:** High  
**Impact:** Risk of failed migrations, no rollback verification, deployment risk

**Description:**
- Migration files exist
- No rollback testing
- No migration verification
- No data integrity checks

**Remediation:**
```bash
# Add migration testing
migrate -path migrations -database $DATABASE_URL down
migrate -path migrations -database $DATABASE_URL up

# Add data integrity checks
pg_verify_checksums $DATA_DIR
```

**Estimated Effort:** 1 week

---

### 19. No Environment-Specific Configurations

**Location:** Backend & Frontend  
**Severity:** High  
**Impact:** Configuration errors, deployment issues, environment drift

**Description:**
- Single `.env` file
- No environment-specific configs
- No config validation
- No secret management

**Remediation:**
```bash
# Add environment-specific configs
.env.development
.env.staging
.env.production

# Add config validation
func LoadConfig() (*Config, error) {
    cfg := &Config{}
    if err := env.Parse(cfg); err != nil {
        return nil, err
    }
    if err := cfg.Validate(); err != nil {
        return nil, err
    }
    return cfg, nil
}
```

**Estimated Effort:** 1 week

---

### 20. No API Documentation for Consumers

**Location:** Backend API  
**Severity:** High  
**Impact:** Poor developer experience, integration difficulties, support burden

**Description:**
- Scalar documentation exists
- No usage examples
- No error code reference
- No integration guides

**Remediation:**
```markdown
# Add to API documentation
## Error Codes

| Code | Description | HTTP Status |
|------|-------------|-------------|
| AUTH_001 | Invalid credentials | 401 |
| AUTH_002 | Token expired | 401 |
| VAL_001 | Validation error | 400 |
```

**Estimated Effort:** 1 week

---

## Medium Debt (Plan to Address)

### 21. Inconsistent Naming Conventions

**Location:** Backend & Frontend  
**Severity:** Medium  
**Impact:** Code readability, maintenance burden

**Description:**
- Mixed naming styles (camelCase, snake_case)
- Inconsistent API field names
- Inconsistent database column names

**Remediation:**
- Establish naming convention guide
- Refactor inconsistent names
- Add linting rules

**Estimated Effort:** 1 week

---

### 22. No Code Comments

**Location:** Backend & Frontend  
**Severity:** Medium  
**Impact:** Code maintainability, onboarding difficulty

**Description:**
- Minimal code comments
- No complex logic documentation
- No API documentation in code

**Remediation:**
- Add godoc comments to Go code
- Add JSDoc comments to TypeScript
- Document complex algorithms

**Estimated Effort:** 2 weeks

---

### 23. No Component Library

**Location:** Frontend  
**Severity:** Medium  
**Impact:** Inconsistent UI, development overhead, maintenance burden

**Description:**
- Components scattered across features
- No design system
- No reusable component library
- Inconsistent styling

**Remediation:**
- Create component library (Storybook)
- Establish design tokens
- Document component usage
- Add component tests

**Estimated Effort:** 3-4 weeks

---

### 24. No State Management Strategy

**Location:** Frontend  
**Severity:** Medium  
**Impact:** State inconsistencies, prop drilling, complexity

**Description:**
- Zustand used but no strategy
- State scattered across components
- No state normalization
- No state persistence strategy

**Remediation:**
- Establish state management patterns
- Add state normalization
- Document state flow
- Add state debugging tools

**Estimated Effort:** 2 weeks

---

### 25. No Performance Budget

**Location:** Frontend  
**Severity:** Medium  
**Impact:** Poor performance, large bundles, slow load times

**Description:**
- No bundle size limits
- No performance budgets
- No performance monitoring
- No optimization strategy

**Remediation:**
```javascript
// Add performance budget to vite.config.ts
export default defineConfig({
  build: {
    rollupOptions: {
      output: {
        manualChunks: {
          vendor: ['react', 'react-dom'],
          mui: ['@mui/material', '@mui/icons-material'],
        },
      },
    },
    chunkSizeWarningLimit: 1000,
  },
});
```

**Estimated Effort:** 1 week

---

### 26. No Accessibility Audit

**Location:** Frontend  
**Severity:** Medium  
**Impact:** Limited accessibility, compliance concerns, user exclusion

**Description:**
- Basic ARIA attributes
- No WCAG compliance
- No keyboard navigation
- No screen reader testing

**Remediation:**
- Run accessibility audit (axe DevTools)
- Fix WCAG violations
- Add keyboard navigation
- Test with screen readers

**Estimated Effort:** 2 weeks

---

### 27. No Mobile Optimization

**Location:** Frontend  
**Severity:** Medium  
**Impact:** Poor mobile experience, limited reach

**Description:**
- Responsive design exists
- No mobile-specific features
- No touch optimization
- No PWA capabilities

**Remediation:**
- Add PWA manifest
- Add service worker
- Optimize touch targets
- Test on mobile devices

**Estimated Effort:** 2 weeks

---

### 28. No Internationalization Complete

**Location:** Frontend  
**Severity:** Medium  
**Impact:** Limited to Indonesian users, no multi-language support

**Description:**
- i18next configured
- Only Indonesian locale
- No translation files
- No language switcher

**Remediation:**
- Add English translations
- Add language switcher
- Add translation management
- Add locale detection

**Estimated Effort:** 1 week

---

### 29. No Analytics Integration

**Location:** Frontend  
**Severity:** Medium  
**Impact:** Limited insights, poor data-driven decisions

**Description:**
- No user analytics
- No feature usage tracking
- No conversion tracking
- No user journey analysis

**Remediation:**
```typescript
// Add analytics
import ReactGA from 'react-ga';

ReactGA.initialize('GA_TRACKING_ID');
ReactGA.pageview(window.location.pathname);
```

**Estimated Effort:** 1 week

---

### 30. No SEO Optimization

**Location:** Frontend  
**Severity:** Medium  
**Impact:** Poor discoverability, limited reach

**Description:**
- No meta tags
- No Open Graph tags
- No structured data
- No sitemap

**Remediation:**
```typescript
// Add meta tags
import { Helmet } from 'react-helmet';

<Helmet>
  <title>NUSA Platform - Education Management</title>
  <meta name="description" content="Kurikulum Merdeka 2026 education management system" />
</Helmet>
```

**Estimated Effort:** 3 days

---

## Low Debt (Nice to Have)

### 31. No Code Formatting Automation

**Location:** Backend & Frontend  
**Severity:** Low  
**Impact:** Code style inconsistencies, review overhead

**Description:**
- Prettier configured
- No auto-format on save
- No pre-commit hooks
- No CI formatting checks

**Remediation:**
- Enable auto-format on save
- Add pre-commit hooks
- Add CI formatting checks

**Estimated Effort:** 2 days

---

### 32. No Dependency Updates Automation

**Location:** Backend & Frontend  
**Severity:** Low  
**Impact:** Security vulnerabilities, outdated dependencies

**Description:**
- Manual dependency updates
- No Dependabot
- No security scanning
- No update notifications

**Remediation:**
- Enable Dependabot
- Add security scanning
- Schedule dependency updates

**Estimated Effort:** 2 days

---

### 33. No Documentation Site

**Location:** Project  
**Severity:** Low  
**Impact:** Poor developer experience, onboarding difficulty

**Description:**
- API documentation exists
- No developer documentation
- No architecture documentation
- No deployment guides

**Remediation:**
- Set up documentation site (Docusaurus)
- Add developer guides
- Add architecture diagrams
- Add deployment guides

**Estimated Effort:** 2 weeks

---

### 34. No Contribution Guidelines

**Location:** Project  
**Severity:** Low  
**Impact:** Poor contribution experience, code quality issues

**Description:**
- CONTRIBUTING.md exists
- No detailed guidelines
- No code review checklist
- No PR templates

**Remediation:**
- Expand CONTRIBUTING.md
- Add code review checklist
- Add PR templates
- Add issue templates

**Estimated Effort:** 2 days

---

### 35. No Changelog Automation

**Location:** Project  
**Severity:** Low  
**Impact**: Poor release tracking, communication gaps

**Description:**
- CHANGELOG.md exists
- Manual updates only
- No automated changelog
- No release notes

**Remediation:**
- Add conventional commits
- Add automated changelog
- Add release notes generation

**Estimated Effort:** 2 days

---

## Debt Prioritization Matrix

| Debt Item | Severity | Effort | Impact | Priority |
|-----------|----------|--------|--------|----------|
| Missing Test Coverage | Critical | 3-4 weeks | High | P0 |
| No Error Tracking | Critical | 1-2 weeks | High | P0 |
| No Database Backup | Critical | 1 week | High | P0 |
| No Rate Limiting | Critical | 1 week | High | P0 |
| No Audit Logging | Critical | 2 weeks | High | P0 |
| No Input Validation | Critical | 1-2 weeks | High | P0 |
| No Security Headers | Critical | 3 days | High | P0 |
| No File Storage | Critical | 2 weeks | High | P0 |
| No Caching Layer | High | 2-3 weeks | Medium | P1 |
| Limited Error Handling | High | 1-2 weeks | Medium | P1 |
| No Database Indexes | High | 1 week | High | P1 |
| No Soft Delete | High | 2 weeks | Medium | P1 |
| No Connection Pooling | High | 3 days | Medium | P1 |
| No API Versioning | High | 2 weeks | Medium | P2 |
| No Request Validation | High | 1 week | Medium | P1 |
| No Pagination Consistency | High | 1 week | Medium | P2 |
| No Background Jobs | High | 3 weeks | Medium | P2 |
| No Migration Testing | High | 1 week | High | P1 |
| No Env Configs | High | 1 week | Medium | P1 |
| No API Docs | High | 1 week | Medium | P2 |
| Inconsistent Naming | Medium | 1 week | Low | P3 |
| No Code Comments | Medium | 2 weeks | Low | P3 |
| No Component Library | Medium | 3-4 weeks | Medium | P2 |
| No State Strategy | Medium | 2 weeks | Medium | P3 |
| No Performance Budget | Medium | 1 week | Medium | P3 |
| No Accessibility | Medium | 2 weeks | Medium | P3 |
| No Mobile Optimization | Medium | 2 weeks | Medium | P3 |
| No i18n Complete | Medium | 1 week | Low | P3 |
| No Analytics | Medium | 1 week | Low | P3 |
| No SEO | Medium | 3 days | Low | P3 |
| No Code Formatting | Low | 2 days | Low | P4 |
| No Dependency Updates | Low | 2 days | Low | P4 |
| No Documentation Site | Low | 2 weeks | Low | P4 |
| No Contribution Guidelines | Low | 2 days | Low | P4 |
| No Changelog Automation | Low | 2 days | Low | P4 |

---

## Remediation Roadmap

### Phase 1: Critical Debt (Weeks 1-4)

**Week 1-2:**
- Add error tracking (Sentry)
- Add security headers
- Add input validation layer
- Add database backup strategy

**Week 3-4:**
- Add rate limiting
- Add audit logging
- Add file storage strategy
- Begin test coverage

### Phase 2: High Debt (Weeks 5-8)

**Week 5-6:**
- Add database indexes
- Add soft delete
- Add connection pooling configuration
- Add migration testing

**Week 7-8:**
- Add caching layer (Redis)
- Improve error handling
- Add request validation middleware
- Standardize pagination

### Phase 3: Medium Debt (Weeks 9-12)

**Week 9-10:**
- Add background job processing
- Add API versioning strategy
- Add environment-specific configurations
- Add API documentation

**Week 11-12:**
- Create component library
- Add state management strategy
- Add performance budget
- Complete test coverage

### Phase 4: Low Debt (Ongoing)

**Ongoing:**
- Code formatting automation
- Dependency updates automation
- Documentation site
- Contribution guidelines
- Changelog automation

---

## Debt Prevention Strategy

### Code Review Checklist

- [ ] Tests added for new functionality
- [ ] Error handling implemented
- [ ] Logging added
- [ ] Documentation updated
- [ ] Security review completed
- [ ] Performance review completed
- [ ] Accessibility review completed

### Pre-commit Hooks

```bash
# .husky/pre-commit
npm run lint
npm run format
npm run test
```

### CI/CD Pipeline

```yaml
# Add to CI pipeline
- Run tests
- Run linting
- Run security scanning
- Run dependency check
- Run performance audit
```

### Regular Debt Reviews

- Monthly debt review meetings
- Quarterly debt prioritization
- Annual debt assessment
- Debt metrics tracking

---

## Conclusion

The NUSA Platform has accumulated technical debt across multiple areas, with 8 critical items requiring immediate attention. The most pressing concerns are around testing, monitoring, backup strategy, and security. Addressing these items in the recommended phases will significantly improve the system's reliability, maintainability, and scalability.

The remediation roadmap provides a structured approach to addressing debt over 12 weeks, with ongoing maintenance for low-priority items. Implementing the debt prevention strategy will help minimize future debt accumulation and maintain code quality over time.

**Estimated Total Remediation Effort:** 8-12 weeks  
**Recommended Team Size:** 2-3 developers  
**Risk Level:** Medium (if addressed in phases)  
**Business Impact:** High (improved reliability, security, and maintainability)

# CP Alignment Threshold Configuration Design

**Document Version**: 1.0  
**Date**: 2026-06-11  
**Status**: Final  
**Purpose**: Design specification for configurable CP alignment threshold to replace hardcoded 60% threshold

---

## Executive Summary

CP alignment threshold has been changed from hardcoded value (60%) to configurable system value. This design allows schools and system administrators to adjust the threshold based on their specific needs without requiring application redeployment.

**Key Features**:
- System configuration table for storing threshold values
- Priority-based configuration resolution (database > environment variable > default)
- Configuration API for runtime updates
- Audit trail for configuration changes
- Default fallback to 60% if configuration not set

---

## Design Goals

### Primary Goals

1. **Flexibility**: Allow threshold adjustment without application redeploy
2. **Safety**: Provide default fallback if configuration is missing or invalid
3. **Auditability**: Track all configuration changes
4. **Simplicity**: Easy to understand and maintain
5. **Performance**: Minimal performance impact on alignment calculation

### Non-Goals

- Per-school threshold configuration (out of scope for Sprint 4)
- Per-dimension threshold configuration (out of scope for Sprint 4)
- Threshold based on subject or phase (out of scope for Sprint 4)

---

## Architecture Design

### Configuration Sources (Priority Order)

1. **Database Table** (Highest Priority)
   - Runtime configurable via API
   - Persisted across application restarts
   - Auditable changes

2. **Environment Variable** (Medium Priority)
   - Set at deployment time
   - Override default but not database
   - Useful for testing or staging environments

3. **Default Value** (Lowest Priority)
   - Hardcoded fallback
   - Value: 60.0
   - Ensures system always has a valid threshold

### Configuration Resolution Flow

```
┌─────────────────────┐
│  Read from Database  │
│  system_configuration│
└──────────┬──────────┘
           │
           ├─ Found and valid? ──→ Use database value
           │
           └─ Not found or invalid?
                │
                ↓
┌─────────────────────┐
│ Read from Environment │
│     Variable         │
│  CP_ALIGNMENT_THRESHOLD │
└──────────┬──────────┘
           │
           ├─ Found and valid? ──→ Use environment variable
           │
           └─ Not found or invalid?
                │
                ↓
┌─────────────────────┐
│   Use Default Value │
│        60.0         │
└─────────────────────┘
```

---

## Database Design

### Table: system_configuration

**Purpose**: Store system-level configuration values

**Columns**:
| Column | Type | Constraints | Description |
| ------ | ---- | ----------- | ----------- |
| id | UUID | PRIMARY KEY, DEFAULT gen_uuid_v7() | Unique identifier |
| key | VARCHAR(100) | NOT NULL, UNIQUE | Configuration key |
| value | TEXT | NOT NULL | Configuration value (JSON string for complex values) |
| description | TEXT | | Configuration description |
| created_at | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Creation timestamp |
| updated_at | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Last update timestamp |

**Indexes**:
- `idx_system_configuration_key` on (key)

**Unique Constraints**:
- `uq_system_configuration_key` on (key)

**Audit Fields**: created_at, updated_at

---

### Initial Data Seeding

```sql
-- Seed default CP alignment threshold
INSERT INTO system_configuration (id, key, value, description, created_at, updated_at) VALUES
(gen_uuid_v7(), 'cp_alignment_threshold', '60.0', 'CP alignment threshold percentage (default: 60%)', NOW(), NOW());
```

---

## Domain Model

### ConfigurationService

```go
type ConfigurationService struct {
    db *sql.DB
}

// GetAlignmentThreshold returns the CP alignment threshold value
// Resolution order: database > environment variable > default (60.0)
func (cs *ConfigurationService) GetAlignmentThreshold() (float64, error) {
    // Step 1: Try to read from database
    threshold, err := cs.getFromDatabase()
    if err == nil && threshold > 0 {
        return threshold, nil
    }

    // Step 2: Try to read from environment variable
    threshold, err = cs.getFromEnvironment()
    if err == nil && threshold > 0 {
        return threshold, nil
    }

    // Step 3: Use default value
    return 60.0, nil
}

func (cs *ConfigurationService) getFromDatabase() (float64, error) {
    var value string
    err := cs.db.QueryRow(
        "SELECT value FROM system_configuration WHERE key = $1",
        "cp_alignment_threshold",
    ).Scan(&value)
    
    if err != nil {
        if err == sql.ErrNoRows {
            return 0, nil // Not found, not an error
        }
        return 0, err
    }
    
    threshold, err := strconv.ParseFloat(value, 64)
    if err != nil {
        return 0, err
    }
    
    // Validate threshold range
    if threshold < 0 || threshold > 100 {
        return 0, errors.New("threshold must be between 0 and 100")
    }
    
    return threshold, nil
}

func (cs *ConfigurationService) getFromEnvironment() (float64, error) {
    env := os.Getenv("CP_ALIGNMENT_THRESHOLD")
    if env == "" {
        return 0, nil // Not set, not an error
    }
    
    threshold, err := strconv.ParseFloat(env, 64)
    if err != nil {
        return 0, err
    }
    
    // Validate threshold range
    if threshold < 0 || threshold > 100 {
        return 0, errors.New("threshold must be between 0 and 100")
    }
    
    return threshold, nil
}

// UpdateAlignmentThreshold updates the threshold in the database
func (cs *ConfigurationService) UpdateAlignmentThreshold(threshold float64, updaterID UUID) error {
    // Validate threshold
    if threshold < 0 || threshold > 100 {
        return errors.New("threshold must be between 0 and 100")
    }
    
    // Update or insert configuration
    query := `
        INSERT INTO system_configuration (id, key, value, description, created_at, updated_at)
        VALUES ($1, $2, $3, $4, NOW(), NOW())
        ON CONFLICT (key) DO UPDATE SET
            value = $3,
            updated_at = NOW()
    `
    
    _, err := cs.db.Exec(
        query,
        gen_uuid_v7(),
        "cp_alignment_threshold",
        fmt.Sprintf("%.2f", threshold),
        "CP alignment threshold percentage",
    )
    
    return err
}
```

---

### AlignmentCalculationService (Updated)

```go
type AlignmentCalculationService struct {
    configService ConfigurationService
}

func NewAlignmentCalculationService(configService ConfigurationService) *AlignmentCalculationService {
    return &AlignmentCalculationService{
        configService: configService,
    }
}

func (s *AlignmentCalculationService) CalculateCPAlignment(
    alignments []CPAlignment,
    dimensions []GraduateProfileDimension,
) (float64, error) {
    if len(alignments) == 0 {
        return 0, nil
    }
    
    totalWeight := 0.0
    totalScore := 0.0
    
    for _, alignment := range alignments {
        dimension := findDimension(dimensions, alignment.DimensionID)
        if dimension == nil {
            continue
        }
        
        strength := s.strengthToPercentage(alignment.AlignmentStrength)
        weightedScore := strength * dimension.Weight
        totalScore += weightedScore
        totalWeight += dimension.Weight
    }
    
    if totalWeight == 0 {
        return 0, nil
    }
    
    return (totalScore / totalWeight) * 100, nil
}

func (s *AlignmentCalculationService) strengthToPercentage(strength AlignmentStrength) float64 {
    switch strength {
    case AlignmentStrengthStrong:
        return 1.0  // 100%
    case AlignmentStrengthMedium:
        return 0.75 // 75%
    case AlignmentStrengthWeak:
        return 0.5  // 50%
    default:
        return 0.0
    }
}

// CheckAgainstThreshold checks if CP alignment meets the configurable threshold
func (s *AlignmentCalculationService) CheckAgainstThreshold(alignmentPercentage float64) (bool, error) {
    // Get threshold from configuration (resolves priority order automatically)
    threshold, err := s.configService.GetAlignmentThreshold()
    if err != nil {
        return false, fmt.Errorf("failed to get alignment threshold: %w", err)
    }
    
    return alignmentPercentage >= threshold, nil
}

// GetThresholdValue returns the current threshold value for display purposes
func (s *AlignmentCalculationService) GetThresholdValue() (float64, error) {
    return s.configService.GetAlignmentThreshold()
}

// GetThresholdGap returns the gap between alignment percentage and threshold
func (s *AlignmentCalculationService) GetThresholdGap(alignmentPercentage float64) (float64, error) {
    threshold, err := s.configService.GetAlignmentThreshold()
    if err != nil {
        return 0, err
    }
    
    gap := threshold - alignmentPercentage
    if gap < 0 {
        return 0, nil // Already above threshold
    }
    
    return gap, nil
}
```

---

## API Design

### Get Configuration

**Method**: `GET`  
**URL**: `/api/v1/system/configuration`  
**Authorization**: `SYSTEM_ADMIN` (read access) or `CURRICULUM_ADMIN` (read access)  
**Permission**: `system_config:READ`

**Query Parameters**:
- `key` (optional): Filter by configuration key

**Response**: `200 OK`
```json
{
  "data": [
    {
      "id": "uuid",
      "key": "cp_alignment_threshold",
      "value": "60.0",
      "description": "CP alignment threshold percentage (default: 60%)",
      "created_at": "2026-06-11T10:00:00Z",
      "updated_at": "2026-06-11T10:00:00Z"
    }
  ]
}
```

**Error Codes**:
- `403 FORBIDDEN`: Insufficient permissions
- `404 NOT_FOUND`: Configuration key not found (if specific key requested)

---

### Update Configuration

**Method**: `PUT`  
**URL**: `/api/v1/system/configuration/:key`  
**Authorization**: `SYSTEM_ADMIN` only  
**Permission**: `system_config:UPDATE`

**Request**:
```json
{
  "value": "70.0",
  "description": "Increased threshold to 70% for pilot phase"
}
```

**Response**: `200 OK`
```json
{
  "id": "uuid",
  "key": "cp_alignment_threshold",
  "value": "70.0",
  "description": "Increased threshold to 70% for pilot phase",
  "created_at": "2026-06-11T10:00:00Z",
  "updated_at": "2026-06-11T12:00:00Z"
}
```

**Validation**:
- Key must exist
- Value must be valid number antara 0 dan 100
- Only SYSTEM_ADMIN can update configuration
- Audit log entry created

**Error Codes**:
- `400 BAD_REQUEST`: Invalid key or value
- `403 FORBIDDEN`: Insufficient permissions (SYSTEM_ADMIN only)
- `404 NOT_FOUND`: Configuration key not found

---

### Get Alignment Threshold (Convenience Endpoint)

**Method**: `GET`  
**URL**: `/api/v1/system/configuration/alignment-threshold`  
**Authorization**: `SYSTEM_ADMIN`, `CURRICULUM_ADMIN`, `SCHOOL_ADMIN`, `TEACHER`  
**Permission**: `system_config:READ`

**Response**: `200 OK`
```json
{
  "threshold": 60.0,
  "source": "database",
  "description": "CP alignment threshold percentage"
}
```

**Notes**:
- `source` field indicates where value came from (database, environment, default)
- Convenience endpoint for quick access to threshold value

---

## Handler Implementation

### ConfigurationHandler

```go
type ConfigurationHandler struct {
    configService ConfigurationService
}

func NewConfigurationHandler(configService ConfigurationService) *ConfigurationHandler {
    return &ConfigurationHandler{
        configService: configService,
    }
}

// GetConfiguration retrieves configuration values
func (h *ConfigurationHandler) GetConfiguration(c *gin.Context) {
    key := c.Query("key")
    
    configs, err := h.configService.GetConfigurations(key)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(200, gin.H{"data": configs})
}

// UpdateConfiguration updates a configuration value
func (h *ConfigurationHandler) UpdateConfiguration(c *gin.Context) {
    var req UpdateConfigRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    
    key := c.Param("key")
    userID := c.GetString("user_id") // From JWT token
    
    err := h.configService.UpdateConfiguration(key, req.Value, userID)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(200, gin.H{"message": "Configuration updated successfully"})
}

// GetAlignmentThreshold retrieves the current alignment threshold
func (h *ConfigurationHandler) GetAlignmentThreshold(c *gin.Context) {
    threshold, source, err := h.configService.GetAlignmentThresholdWithSource()
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(200, gin.H{
        "threshold": threshold,
        "source": source,
        "description": "CP alignment threshold percentage",
    })
}
```

---

## Frontend Implementation

### System Configuration Page (System Admin only)

**Path**: `/admin/system/configuration`  
**Purpose**: Manage system-level configuration values  
**Actors**: System Admin only

**Components**:
- **Table**: List configuration items dengan columns (Key, Value, Description, Source, Last Updated)
- **Edit Modal**: Edit configuration value
- **Actions**: Update configuration
- **Refresh**: Reload configuration from database

**Form Fields**:
- Key (read-only text)
- Value (text input, validated as number 0-100)
- Description (textarea, optional)

**Validation**:
- Value must be valid number
- Value must be antara 0 dan 100
- Only SYSTEM_ADMIN can update

**Permissions**: `system_config:READ`, `system_config:UPDATE`

---

### Alignment Threshold Display (All Users)

**Components**:
- Display threshold value in alignment report header
- Display threshold value in CP alignment list
- Show threshold source indicator (database/environment/default)

---

## Error Handling

### Configuration Not Found

**Scenario**: Configuration key does not exist in database

**Behavior**:
- Return default value (60.0) instead of error
- Log warning that configuration not found
- Continue with default threshold

**Code**:
```go
threshold, err := cs.getFromDatabase()
if err == sql.ErrNoRows {
    // Not found, not an error - try next source
    return 0, nil
}
if err != nil {
    return 0, err // Actual error
}
```

---

### Invalid Configuration Value

**Scenario**: Configuration value is invalid (not a number, out of range)

**Behavior**:
- Log error
- Skip this source and try next
- Eventually fall back to default

**Code**:
```go
threshold, err := strconv.ParseFloat(value, 64)
if err != nil {
    log.Warn("Invalid threshold value in database: %v", value)
    return 0, nil // Try next source
}

if threshold < 0 || threshold > 100 {
    log.Warn("Threshold out of range in database: %v", threshold)
    return 0, nil // Try next source
}
```

---

## Security Considerations

### Access Control

- **Read Access**: SYSTEM_ADMIN, CURRICULUM_ADMIN, SCHOOL_ADMIN, TEACHER
- **Write Access**: SYSTEM_ADMIN only
- **Audit Trail**: All configuration changes logged

### Validation

- Value must be valid number
- Value must be in valid range (0-100)
- Key must match expected format

### Rate Limiting

- Configuration updates should be rate-limited
- Suggest: Max 10 updates per hour per user

---

## Performance Considerations

### Caching Strategy

**Option 1: No Caching (Current Design)**
- Pros: Always reflects latest configuration
- Cons: Database query on every alignment calculation

**Option 2: In-Memory Caching (Recommended for Future)**
- Pros: Reduced database load
- Cons: Requires cache invalidation strategy
- Cache TTL: 5 minutes
- Invalidation: On configuration update

**Recommendation**: Start with no caching (Option 1) for Sprint 4 MVP. Add caching in future if performance issues arise.

---

### Database Query Optimization

**Current Design**: Single query per configuration lookup

**Optimization**: Consider caching configuration values in application context during initialization

```go
type Config struct {
    CPAlignmentThreshold float64
}

func LoadConfig(db *sql.DB) (*Config, error) {
    // Load all configurations at startup
    config := &Config{
        CPAlignmentThreshold: 60.0, // Default
    }
    
    // Load from database
    var value string
    err := db.QueryRow("SELECT value FROM system_configuration WHERE key = $1", "cp_alignment_threshold").Scan(&value)
    if err == nil {
        config.CPAlignmentThreshold, _ = strconv.ParseFloat(value, 64)
    }
    
    return config, nil
}
```

---

## Testing Strategy

### Unit Tests

```go
func TestConfigurationService_GetAlignmentThreshold(t *testing.T) {
    tests := []struct {
        name           string
        dbValue        *string
        envValue       *string
        expected       float64
        expectedSource string
    }{
        {
            name:           "Database value used when present",
            dbValue:        stringPtr("70.0"),
            envValue:       nil,
            expected:       70.0,
            expectedSource: "database",
        },
        {
            name:           "Environment variable used when database empty",
            dbValue:        nil,
            envValue:       stringPtr("75.0"),
            expected:       75.0,
            expectedSource: "environment",
        },
        {
            name:           "Default value used when both empty",
            dbValue:        nil,
            envValue:       nil,
            expected:       60.0,
            expectedSource: "default",
        },
        {
            name:           "Invalid database value falls through to environment",
            dbValue:        stringPtr("invalid"),
            envValue:       stringPtr("65.0"),
            expected:       65.0,
            expectedSource: "environment",
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Setup test database and environment
            // Test configuration service
            // Verify result
        })
    }
}
```

---

### Integration Tests

```go
func TestAlignmentCalculationService_CheckAgainstThreshold(t *testing.T) {
    tests := []struct {
        name              string
        alignmentPercent  float64
        configuredThreshold float64
        expectedPass      bool
    }{
        {
            name:                "Above threshold passes",
            alignmentPercent:  70.0,
            configuredThreshold: 60.0,
            expectedPass:      true,
        },
        {
            name:                "Below threshold fails",
            alignmentPercent:  50.0,
            configuredThreshold: 60.0,
            expectedPass:      false,
        },
        {
            name:                "Exactly threshold passes",
            alignmentPercent:  60.0,
            configuredThreshold: 60.0,
            expectedPass:      true,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Setup configuration with threshold
            // Test alignment check
            // Verify result
        })
    }
}
```

---

## Migration Strategy

### Database Migration

```sql
-- Create system_configuration table
CREATE TABLE system_configuration (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    key VARCHAR(100) NOT NULL UNIQUE,
    value TEXT NOT NULL,
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_system_configuration_key ON system_configuration(key);

-- Seed default configuration
INSERT INTO system_configuration (id, key, value, description, created_at, updated_at) VALUES
(gen_uuid_v7(), 'cp_alignment_threshold', '60.0', 'CP alignment threshold percentage (default: 60%)', NOW(), NOW());
```

### Rollback

```sql
-- Cannot rollback configuration data (acceptable)
DROP TABLE IF EXISTS system_configuration;
```

---

## Monitoring

### Metrics to Track

| Metric | Description | Alert Threshold |
| ------ | ----------- | --------------- |
| Configuration Source | Percentage of threshold reads from each source | - |
| Configuration Update Rate | Number of threshold updates per day | >10/day |
- | Configuration Error Rate | Percentage of failed configuration reads | >1% |

### Logging

- Log configuration resolution source
- Log configuration updates with user context
- Log invalid configuration values

---

## Documentation Updates

### Files to Update

1. **DATABASE_SCHEMA_FREEZE_V1.md**
   - Add system_configuration table definition
   - Document configuration priority resolution

2. **API_DOCUMENTATION.md**
   - Add configuration API endpoints
   - Document error codes

3. **OPERATIONS_GUIDE.md**
   - Add configuration management procedures
   - Document threshold adjustment process

---

## Future Enhancements (Out of Scope for Sprint 4)

### Per-School Threshold Configuration

**Description**: Allow each school to have its own threshold value

**Design**:
- Add school_id column to system_configuration table
- Query threshold by school_id
- Add UI for school-specific configuration

**Rationale for Out of Scope**:
- Sprint 4 is Academic Foundation (system-wide governance)
- Per-school configuration adds complexity
- Can be added in future if needed

---

### Per-Dimension Threshold Configuration

**Description**: Allow different thresholds for different graduate profile dimensions

**Design**:
- Add dimension_id column to threshold configuration
- More granular alignment checking

**Rationale for Out of Scope**:
- Adds complexity without clear business need
- Current design (single threshold) is simpler and meets MVP requirements

---

### Threshold Based on Subject or Phase

**Description**: Different thresholds for different subjects or phases

**Design**:
- Add subject_id or phase_id to threshold configuration
- More granular alignment checking

**Rationale for Out of Scope**:
- Adds complexity
- Unclear business need
- Can be added in future if specific requirements emerge

---

## Approval

**Approved By**: Product Owner  
**Approval Date**: 2026-06-11  
**Approval Status**: APPROVED

**Next Steps**:
1. Implement system_configuration table in migration
2. Implement ConfigurationService
3. Update AlignmentCalculationService to use configurable threshold
4. Implement configuration API endpoints
5. Implement frontend configuration management page
6. Update alignment reports to display threshold
7. Update error messages to use threshold value
8. Add comprehensive tests
9. Document threshold adjustment procedure for operators

---

**Document Status**: FINAL - READY FOR IMPLEMENTATION
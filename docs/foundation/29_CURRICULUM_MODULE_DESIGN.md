# 29_CURRICULUM_MODULE_DESIGN.md

## Foundation Document for NUSA Education Platform

**Version**: 1.0
**Date**: June 2026
**Status**: FOUNDATION DOCUMENT
**Alignment**: Aligned with 14_DATABASE_SCHEMA.md, 23_CURRICULUM_VERSIONING_ARCHITECTURE.md, 27_BACKEND_FOUNDATION_DESIGN.md, 28_AUTHENTICATION_DESIGN.md

**Purpose**: Design the curriculum management module for NUSA. This document defines curriculum CRUD operations, versioning, import strategy, validation rules, search design, API design, permission rules, and database mapping for MVP.

---

# SECTION 1 — Executive Summary

## Why Curriculum Module Design Matters

A well-designed curriculum module ensures:
- Structured representation of the Indonesian National Curriculum
- Support for curriculum versioning and updates
- Efficient search and retrieval of curriculum elements
- Data integrity through validation rules
- Secure access control based on user roles
- Scalable import mechanism for bulk curriculum data

## Curriculum Hierarchy

```
Subject
  ↓
Phase
  ↓
Element
  ↓
Subelement
  ↓
CP (Capaian Pembelajaran)
```

**Description**:
- **Subject**: Academic subject (e.g., Mathematics, Indonesian Language)
- **Phase**: Educational phase (e.g., Phase A, Phase B, Phase C)
- **Element**: Curriculum element grouping (e.g., Reading, Writing)
- **Subelement**: Specific sub-element (e.g., Reading Comprehension)
- **CP**: Learning Outcome (Capaian Pembelajaran)

## Core Principles

- **Hierarchical Structure**: Clear parent-child relationships
- **Versioning Support**: Curriculum version tracking for updates
- **Validation First**: Strict validation rules for data integrity
- **Search-Optimized**: Efficient search and filtering
- **Role-Based Access**: Permissions based on user roles
- **MVP-Focused**: Single ACTIVE curriculum version for MVP

---

# SECTION 2 — Curriculum CRUD

## CRUD Operations

### Create

**Purpose**: Create new curriculum entities (Subject, Phase, Element, Subelement, CP).

**Scope**:
- MVP: Curriculum creation via import only (no manual CRUD)
- Future: Manual CRUD for curriculum administrators

**Implementation**: See Curriculum Import Strategy (Section 4).

### Read

**Purpose**: Retrieve curriculum entities with filtering and pagination.

**Operations**:
- Get single entity by ID
- List entities with filters
- Get hierarchy tree
- Search by text

**Implementation**: See Curriculum Search Design (Section 5).

### Update

**Purpose**: Update existing curriculum entities.

**Scope**:
- MVP: No updates allowed (immutable curriculum)
- Future: Version-based updates (create new version)

**Implementation**: See Curriculum Versioning (Section 3).

### Delete

**Purpose**: Delete curriculum entities.

**Scope**:
- MVP: No deletes allowed (immutable curriculum)
- Future: Soft delete with version tracking

**Implementation**: Not applicable for MVP.

---

# SECTION 3 — Curriculum Versioning

## Versioning Model

### Curriculum Version Entity

**Purpose**: Represents a specific version of the national curriculum.

**Fields**:

| Field | Type | Constraints | Description |
|-------|------|-------------|-------------|
| `id` | UUID | PRIMARY KEY | Unique version identifier |
| `curriculum_version_code` | VARCHAR(50) | NOT NULL, UNIQUE | Version code (e.g., "KUR2025") |
| `effective_year` | INTEGER | NOT NULL | Academic year when version becomes effective |
| `effective_from` | DATE | NOT NULL | Effective start date |
| `effective_until` | DATE | NULLABLE | Effective end date (null for current) |
| `status` | VARCHAR(20) | NOT NULL | Version status |
| `description` | TEXT | NULLABLE | Version description |
| `created_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Creation timestamp |
| `updated_at` | TIMESTAMP WITH TIME ZONE | NOT NULL, DEFAULT NOW() | Last update timestamp |

### Status Values

| Status | Description | Transitions |
|--------|-------------|-------------|
| DRAFT | Version under preparation | DRAFT → ACTIVE |
| ACTIVE | Currently in use | ACTIVE → DEPRECATED |
| DEPRECATED | No longer recommended for new use | DEPRECATED → ARCHIVED |
| ARCHIVED | Historical reference only | No transitions |

## MVP Simplification

### Single ACTIVE Version Rule

**Rule**: Exactly one ACTIVE curriculum version exists system-wide.

**Implications**:
- Only one curriculum version can have status = ACTIVE at any time
- All schools use the current ACTIVE curriculum version
- Schools cannot choose curriculum version
- No school_curriculum_mapping table required for MVP
- Curriculum selection is system-wide, not per-school

### Active Version Management

**System Administrator Actions**:
- Set one curriculum version to ACTIVE
- All other versions become DRAFT, DEPRECATED, or ARCHIVED
- New curriculum versions can be imported but must not be set to ACTIVE until ready
- When a new version becomes ACTIVE, the previous version automatically becomes DEPRECATED

### School Behavior

**Automatic Usage**:
- All schools automatically use the ACTIVE curriculum version
- No UI for school curriculum selection
- No per-school curriculum configuration
- All CP queries default to ACTIVE curriculum version

### Artifact Traceability

**Version Linking**:
- All artifacts (TP Sets, ATP Sets, Modul Ajar Sets) reference curriculum_version_id
- Traceability preserved even with single ACTIVE version
- Historical artifacts remain linked to their originating curriculum version

## Future Enhancement (Wave 2)

### Multi-Version Support

Future enhancement will enable:
- Multiple ACTIVE curriculum versions to coexist
- School-level curriculum selection via school_curriculum_mapping table
- Schools to migrate at their own pace
- Per-school curriculum configuration UI
- Gradual migration without data loss

**Foundation Preserved**:
- Curriculum version table structure supports multi-version
- All curriculum hierarchy entities include curriculum_version_id
- Artifact traceability via curriculum_version_id preserved
- No schema changes required for Wave 2

---

# SECTION 4 — Curriculum Import Strategy

## Import Format

### JSON Format

Curriculum data is imported via JSON file with hierarchical structure.

**Example Structure**:

```json
{
  "curriculum_version": {
    "code": "KUR2025",
    "effective_year": 2025,
    "effective_from": "2025-07-01",
    "description": "Kurikulum Merdeka 2025"
  },
  "subjects": [
    {
      "code": "MAT",
      "name": "Matematika",
      "name_en": "Mathematics",
      "phases": [
        {
          "code": "A",
          "name": "Fase A",
          "grade_levels": ["1", "2"],
          "elements": [
            {
              "code": "BIL",
              "name": "Bilangan",
              "name_en": "Numbers",
              "subelements": [
                {
                  "code": "BIL01",
                  "name": "Bilangan Cacah",
                  "name_en": "Whole Numbers",
                  "cps": [
                    {
                      "code": "CP.BIL.01",
                      "text": "Peserta didik dapat membaca bilangan cacah sampai 100",
                      "text_en": "Students can read whole numbers up to 100",
                      "fase": "A"
                    }
                  ]
                }
              ]
            }
          ]
        }
      ]
    }
  ]
}
```

## Import Process

### Sequence Diagram

```
┌─────────┐       ┌─────────┐       ┌─────────┐       ┌─────────┐
│  Client │       │  Server │       │  Database│       │  File    │
└────┬────┘       └────┬────┘       └────┬────┘       └────┬────┘
     │                  │                  │                  │
     │ 1. Upload JSON  │                  │                  │
     │ POST /api/v1/admin/curriculum/import│              │
     │ multipart/form-data               │                  │
     ├─────────────────>│                  │                  │
     │                  │                  │                  │
     │                  │ 2. Validate file                  │
     │                  ├──────────────────────────────────>│
     │                  │                  │                  │
     │                  │ Valid JSON     │                  │
     │                  │<──────────────────────────────────┤
     │                  │                  │                  │
     │                  │ 3. Validate structure             │
     │                  │                  │                  │
     │                  │ 4. Start transaction              │
     │                  ├─────────────────>│                  │
     │                  │                  │                  │
     │                  │ Transaction     │                  │
     │                  │<─────────────────┤                  │
     │                  │                  │                  │
     │                  │ 5. Create curriculum version       │
     │                  ├─────────────────>│                  │
     │                  │                  │                  │
     │                  │ Success         │                  │
     │                  │<─────────────────┤                  │
     │                  │                  │                  │
     │                  │ 6. Create subjects                 │
     │                  ├─────────────────>│                  │
     │                  │                  │                  │
     │                  │ Success         │                  │
     │                  │<─────────────────┤                  │
     │                  │                  │                  │
     │                  │ 7. Create phases                   │
     │                  ├─────────────────>│                  │
     │                  │                  │                  │
     │                  │ Success         │                  │
     │                  │<─────────────────┤                  │
     │                  │                  │                  │
     │                  │ 8. Create elements                 │
     │                  ├─────────────────>│                  │
     │                  │                  │                  │
     │                  │ Success         │                  │
     │                  │<─────────────────┤                  │
     │                  │                  │                  │
     │                  │ 9. Create subelements              │
     │                  ├─────────────────>│                  │
     │                  │                  │                  │
     │                  │ Success         │                  │
     │                  │<─────────────────┤                  │
     │                  │                  │                  │
     │                  │ 10. Create CPs                     │
     │                  ├─────────────────>│                  │
     │                  │                  │                  │
     │                  │ Success         │                  │
     │                  │<─────────────────┤                  │
     │                  │                  │                  │
     │                  │ 11. Commit transaction             │
     │                  ├─────────────────>│                  │
     │                  │                  │                  │
     │                  │ Success         │                  │
     │                  │<─────────────────┤                  │
     │                  │                  │                  │
     │ 12. Response     │                  │                  │
     │ {version_id,     │                  │                  │
     │  stats}          │                  │                  │
     │<─────────────────┤                  │                  │
     │                  │                  │                  │
```

## Import API

### Request

```http
POST /api/v1/admin/curriculum/import
Authorization: Bearer <token>
Content-Type: multipart/form-data

file: curriculum.json
```

### Response (Success)

```json
{
  "version_id": "version-uuid",
  "curriculum_version_code": "KUR2025",
  "stats": {
    "subjects": 10,
    "phases": 30,
    "elements": 150,
    "subelements": 450,
    "cps": 1800
  }
}
```

### Response (Error)

```json
{
  "error": "Invalid curriculum data",
  "details": "Subject code 'MAT' already exists in version KUR2024"
}
```

## Import Implementation

### Service

```go
func (s *service) ImportCurriculum(ctx context.Context, file multipart.File) (*ImportResult, error) {
    // Parse JSON file
    var data CurriculumImportData
    if err := json.NewDecoder(file).Decode(&data); err != nil {
        return nil, fmt.Errorf("invalid JSON: %w", err)
    }
    
    // Validate structure
    if err := s.validateImportData(ctx, &data); err != nil {
        return nil, err
    }
    
    // Start transaction
    tx, err := s.repo.BeginTx(ctx)
    if err != nil {
        return nil, err
    }
    defer tx.Rollback()
    
    // Create curriculum version
    version := &CurriculumVersion{
        ID:                     uuid.New().String(),
        CurriculumVersionCode:   data.CurriculumVersion.Code,
        EffectiveYear:           data.CurriculumVersion.EffectiveYear,
        EffectiveFrom:           data.CurriculumVersion.EffectiveFrom,
        Status:                 "DRAFT",
        Description:            data.CurriculumVersion.Description,
    }
    
    if err := s.repo.CreateCurriculumVersion(ctx, version); err != nil {
        return nil, err
    }
    
    // Import subjects
    stats := ImportStats{}
    for _, subjectData := range data.Subjects {
        subject := &Subject{
            ID:                   uuid.New().String(),
            CurriculumVersionID:  version.ID,
            Code:                 subjectData.Code,
            Name:                 subjectData.Name,
            NameEn:               subjectData.NameEn,
        }
        
        if err := s.repo.CreateSubject(ctx, subject); err != nil {
            return nil, err
        }
        
        stats.Subjects++
        
        // Import phases
        for _, phaseData := range subjectData.Phases {
            phase := &Phase{
                ID:                   uuid.New().String(),
                CurriculumVersionID:  version.ID,
                SubjectID:            subject.ID,
                Code:                 phaseData.Code,
                Name:                 phaseData.Name,
                GradeLevels:          phaseData.GradeLevels,
            }
            
            if err := s.repo.CreatePhase(ctx, phase); err != nil {
                return nil, err
            }
            
            stats.Phases++
            
            // Import elements
            for _, elementData := range phaseData.Elements {
                element := &Element{
                    ID:                   uuid.New().String(),
                    CurriculumVersionID:  version.ID,
                    PhaseID:              phase.ID,
                    Code:                 elementData.Code,
                    Name:                 elementData.Name,
                    NameEn:               elementData.NameEn,
                }
                
                if err := s.repo.CreateElement(ctx, element); err != nil {
                    return nil, err
                }
                
                stats.Elements++
                
                // Import subelements
                for _, subelementData := range elementData.Subelements {
                    subelement := &Subelement{
                        ID:                   uuid.New().String(),
                        CurriculumVersionID:  version.ID,
                        ElementID:            element.ID,
                        Code:                 subelementData.Code,
                        Name:                 subelementData.Name,
                        NameEn:               subelementData.NameEn,
                    }
                    
                    if err := s.repo.CreateSubelement(ctx, subelement); err != nil {
                        return nil, err
                    }
                    
                    stats.Subelements++
                    
                    // Import CPs
                    for _, cpData := range subelementData.CPs {
                        cp := &CP{
                            ID:                   uuid.New().String(),
                            CurriculumVersionID:  version.ID,
                            SubelementID:         subelement.ID,
                            Code:                 cpData.Code,
                            Text:                 cpData.Text,
                            TextEn:               cpData.TextEn,
                            Fase:                 cpData.Fase,
                        }
                        
                        if err := s.repo.CreateCP(ctx, cp); err != nil {
                            return nil, err
                        }
                        
                        stats.CPs++
                    }
                }
            }
        }
    }
    
    // Commit transaction
    if err := tx.Commit(); err != nil {
        return nil, err
    }
    
    return &ImportResult{
        VersionID:              version.ID,
        CurriculumVersionCode:  version.CurriculumVersionCode,
        Stats:                  stats,
    }, nil
}
```

---

# SECTION 5 — Curriculum Search Design

## Search Operations

### Get Single Entity

**Purpose**: Retrieve a single curriculum entity by ID.

**API**:

```http
GET /api/v1/curriculum/cp/:id
Authorization: Bearer <token>
```

**Response**:

```json
{
  "id": "cp-uuid",
  "code": "CP.BIL.01",
  "text": "Peserta didik dapat membaca bilangan cacah sampai 100",
  "text_en": "Students can read whole numbers up to 100",
  "fase": "A",
  "subelement": {
    "id": "subelement-uuid",
    "code": "BIL01",
    "name": "Bilangan Cacah",
    "element": {
      "id": "element-uuid",
      "code": "BIL",
      "name": "Bilangan",
      "phase": {
        "id": "phase-uuid",
        "code": "A",
        "name": "Fase A",
        "subject": {
          "id": "subject-uuid",
          "code": "MAT",
          "name": "Matematika"
        }
      }
    }
  }
}
```

### List Entities with Filters

**Purpose**: List curriculum entities with filtering and pagination.

**API**:

```http
GET /api/v1/curriculum/cp?subject_code=MAT&phase_code=A&page=1&page_size=50
Authorization: Bearer <token>
```

**Response**:

```json
{
  "data": [
    {
      "id": "cp-uuid",
      "code": "CP.BIL.01",
      "text": "Peserta didik dapat membaca bilangan cacah sampai 100",
      "fase": "A"
    }
  ],
  "pagination": {
    "page": 1,
    "page_size": 50,
    "total": 180,
    "total_pages": 4
  }
}
```

### Get Hierarchy Tree

**Purpose**: Retrieve curriculum hierarchy tree for a subject.

**API**:

```http
GET /api/v1/curriculum/subjects/:id/tree
Authorization: Bearer <token>
```

**Response**:

```json
{
  "id": "subject-uuid",
  "code": "MAT",
  "name": "Matematika",
  "phases": [
    {
      "id": "phase-uuid",
      "code": "A",
      "name": "Fase A",
      "grade_levels": ["1", "2"],
      "elements": [
        {
          "id": "element-uuid",
          "code": "BIL",
          "name": "Bilangan",
          "subelements": [
            {
              "id": "subelement-uuid",
              "code": "BIL01",
              "name": "Bilangan Cacah",
              "cps": [
                {
                  "id": "cp-uuid",
                  "code": "CP.BIL.01",
                  "text": "Peserta didik dapat membaca bilangan cacah sampai 100"
                }
              ]
            }
          ]
        }
      ]
    }
  ]
}
```

### Search by Text

**Purpose**: Search CPs by text content.

**API**:

```http
GET /api/v1/curriculum/cp/search?q=membaca bilangan&subject_code=MAT
Authorization: Bearer <token>
```

**Response**:

```json
{
  "data": [
    {
      "id": "cp-uuid",
      "code": "CP.BIL.01",
      "text": "Peserta didik dapat membaca bilangan cacah sampai 100",
      "fase": "A",
      "subject_code": "MAT",
      "subject_name": "Matematika"
    }
  ],
  "pagination": {
    "page": 1,
    "page_size": 20,
    "total": 15,
    "total_pages": 1
  }
}
```

## Search Implementation

### Repository

```go
func (r *repository) SearchCPs(ctx context.Context, query string, filters SearchFilters, pagination Pagination) ([]*CP, int, error) {
    var cps []*CP
    var total int
    
    // Build query
    sql := `
        SELECT cp.*, 
               sub.code as subelement_code,
               sub.name as subelement_name,
               el.code as element_code,
               el.name as element_name,
               ph.code as phase_code,
               ph.name as phase_name,
               su.code as subject_code,
               su.name as subject_name
        FROM cp
        JOIN subelements sub ON cp.subelement_id = sub.id
        JOIN elements el ON sub.element_id = el.id
        JOIN phases ph ON el.phase_id = ph.id
        JOIN subjects su ON ph.subject_id = su.id
        JOIN curriculum_versions cv ON su.curriculum_version_id = cv.id
        WHERE cv.status = 'ACTIVE'
    `
    
    args := []interface{}{}
    argCount := 1
    
    // Add text search
    if query != "" {
        sql += fmt.Sprintf(" AND (cp.text ILIKE $%d OR cp.text_en ILIKE $%d)", argCount, argCount+1)
        args = append(args, "%"+query+"%", "%"+query+"%")
        argCount += 2
    }
    
    // Add filters
    if filters.SubjectCode != "" {
        sql += fmt.Sprintf(" AND su.code = $%d", argCount)
        args = append(args, filters.SubjectCode)
        argCount++
    }
    
    if filters.PhaseCode != "" {
        sql += fmt.Sprintf(" AND ph.code = $%d", argCount)
        args = append(args, filters.PhaseCode)
        argCount++
    }
    
    if filters.Fase != "" {
        sql += fmt.Sprintf(" AND cp.fase = $%d", argCount)
        args = append(args, filters.Fase)
        argCount++
    }
    
    // Get total count
    countSQL := "SELECT COUNT(*) FROM (" + sql + ") as count_query"
    err := r.db.GetContext(ctx, &total, countSQL, args...)
    if err != nil {
        return nil, 0, err
    }
    
    // Add pagination
    sql += fmt.Sprintf(" ORDER BY su.code, ph.code, el.code, sub.code, cp.code")
    sql += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argCount, argCount+1)
    args = append(args, pagination.PageSize, (pagination.Page-1)*pagination.PageSize)
    
    // Execute query
    err = r.db.SelectContext(ctx, &cps, sql, args...)
    if err != nil {
        return nil, 0, err
    }
    
    return cps, total, nil
}
```

---

# SECTION 6 — Curriculum API Design

## API Endpoints

### Public Endpoints (Read-Only)

#### Get Active Curriculum Version

```http
GET /api/v1/public/curriculum/version
```

**Response**:

```json
{
  "id": "version-uuid",
  "curriculum_version_code": "KUR2025",
  "effective_year": 2025,
  "effective_from": "2025-07-01",
  "status": "ACTIVE",
  "description": "Kurikulum Merdeka 2025"
}
```

#### List Subjects

```http
GET /api/v1/public/curriculum/subjects
```

**Response**:

```json
{
  "data": [
    {
      "id": "subject-uuid",
      "code": "MAT",
      "name": "Matematika",
      "name_en": "Mathematics"
    }
  ]
}
```

#### Get Subject Tree

```http
GET /api/v1/public/curriculum/subjects/:id/tree
```

**Response**: See Section 5.

#### List CPs

```http
GET /api/v1/public/curriculum/cp?subject_code=MAT&phase_code=A&page=1&page_size=50
```

**Response**: See Section 5.

#### Search CPs

```http
GET /api/v1/public/curriculum/cp/search?q=membaca bilangan
```

**Response**: See Section 5.

### Admin Endpoints (System Admin Only)

#### Import Curriculum

```http
POST /api/v1/admin/curriculum/import
Authorization: Bearer <token>
Content-Type: multipart/form-data
```

**Request**: See Section 4.

**Response**: See Section 4.

#### List Curriculum Versions

```http
GET /api/v1/admin/curriculum/versions
Authorization: Bearer <token>
```

**Response**:

```json
{
  "data": [
    {
      "id": "version-uuid",
      "curriculum_version_code": "KUR2025",
      "effective_year": 2025,
      "effective_from": "2025-07-01",
      "status": "ACTIVE",
      "description": "Kurikulum Merdeka 2025",
      "created_at": "2025-01-01T00:00:00Z"
    }
  ]
}
```

#### Set Active Curriculum Version

```http
PUT /api/v1/admin/curriculum/versions/:id/activate
Authorization: Bearer <token>
```

**Response**:

```http
204 No Content
```

#### Delete Curriculum Version

```http
DELETE /api/v1/admin/curriculum/versions/:id
Authorization: Bearer <token>
```

**Response**:

```http
204 No Content
```

---

# SECTION 7 — Curriculum Permission Rules

## Permission Matrix

| Resource | Action | SYSTEM_ADMIN | SCHOOL_ADMIN | TEACHER |
|----------|--------|-------------|--------------|---------|
| **Curriculum Version** | | | | |
| | Import | ✅ | ❌ | ❌ |
| | List versions | ✅ | ❌ | ❌ |
| | Set active | ✅ | ❌ | ❌ |
| | Delete version | ✅ | ❌ | ❌ |
| **Subjects** | | | | |
| | Read (all versions) | ✅ | ❌ | ❌ |
| | Read (active version) | ✅ | ✅ | ✅ |
| | Create | ❌ | ❌ | ❌ |
| | Update | ❌ | ❌ | ❌ |
| | Delete | ❌ | ❌ | ❌ |
| **Phases** | | | | |
| | Read (all versions) | ✅ | ❌ | ❌ |
| | Read (active version) | ✅ | ✅ | ✅ |
| | Create | ❌ | ❌ | ❌ |
| | Update | ❌ | ❌ | ❌ |
| | Delete | ❌ | ❌ | ❌ |
| **Elements** | | | | |
| | Read (all versions) | ✅ | ❌ | ❌ |
| | Read (active version) | ✅ | ✅ | ✅ |
| | Create | ❌ | ❌ | ❌ |
| | Update | ❌ | ❌ | ❌ |
| | Delete | ❌ | ❌ | ❌ |
| **Subelements** | | | | |
| | Read (all versions) | ✅ | ❌ | ❌ |
| | Read (active version) | ✅ | ✅ | ✅ |
| | Create | ❌ | ❌ | ❌ |
| | Update | ❌ | ❌ | ❌ |
| | Delete | ❌ | ❌ | ❌ |
| **CPs** | | | | |
| | Read (all versions) | ✅ | ❌ | ❌ |
| | Read (active version) | ✅ | ✅ | ✅ |
| | Search (active version) | ✅ | ✅ | ✅ |
| | Create | ❌ | ❌ | ❌ |
| | Update | ❌ | ❌ | ❌ |
| | Delete | ❌ | ❌ | ❌ |

## Permission Implementation

### Middleware

```go
package middleware

func RequireSystemAdmin() gin.HandlerFunc {
    return RequireRole("SYSTEM_ADMIN")
}

func RequireCurriculumReadAccess() gin.HandlerFunc {
    return func(c *gin.Context) {
        role := c.GetString("role")
        
        // All roles can read active curriculum
        if c.Request.URL.Path == "/api/v1/public/curriculum" {
            c.Next()
            return
        }
        
        // Only SYSTEM_ADMIN can read all versions
        if role != "SYSTEM_ADMIN" {
            c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
            c.Abort()
            return
        }
        
        c.Next()
    }
}
```

---

# SECTION 8 — Curriculum Database Mapping

## Database Schema

### Curriculum Versions Table

```sql
CREATE TABLE curriculum_versions (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    curriculum_version_code VARCHAR(50) NOT NULL UNIQUE,
    effective_year INTEGER NOT NULL,
    effective_from DATE NOT NULL,
    effective_until DATE,
    status VARCHAR(20) NOT NULL CHECK (status IN ('DRAFT', 'ACTIVE', 'DEPRECATED', 'ARCHIVED')),
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_curriculum_versions_code ON curriculum_versions(curriculum_version_code);
CREATE INDEX idx_curriculum_versions_status ON curriculum_versions(status);
CREATE INDEX idx_curriculum_versions_effective_year ON curriculum_versions(effective_year);

-- Ensure only one ACTIVE version
CREATE UNIQUE INDEX idx_curriculum_versions_active ON curriculum_versions(status) 
WHERE status = 'ACTIVE';
```

### Subjects Table

```sql
CREATE TABLE subjects (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    curriculum_version_id UUID NOT NULL REFERENCES curriculum_versions(id) ON DELETE CASCADE,
    code VARCHAR(50) NOT NULL,
    name VARCHAR(255) NOT NULL,
    name_en VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    
    UNIQUE(curriculum_version_id, code)
);

CREATE INDEX idx_subjects_curriculum_version_id ON subjects(curriculum_version_id);
CREATE INDEX idx_subjects_code ON subjects(code);
```

### Phases Table

```sql
CREATE TABLE phases (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    curriculum_version_id UUID NOT NULL REFERENCES curriculum_versions(id) ON DELETE CASCADE,
    subject_id UUID NOT NULL REFERENCES subjects(id) ON DELETE CASCADE,
    code VARCHAR(50) NOT NULL,
    name VARCHAR(255) NOT NULL,
    grade_levels TEXT[],
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    
    UNIQUE(curriculum_version_id, subject_id, code)
);

CREATE INDEX idx_phases_curriculum_version_id ON phases(curriculum_version_id);
CREATE INDEX idx_phases_subject_id ON phases(subject_id);
CREATE INDEX idx_phases_code ON phases(code);
```

### Elements Table

```sql
CREATE TABLE elements (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    curriculum_version_id UUID NOT NULL REFERENCES curriculum_versions(id) ON DELETE CASCADE,
    phase_id UUID NOT NULL REFERENCES phases(id) ON DELETE CASCADE,
    code VARCHAR(50) NOT NULL,
    name VARCHAR(255) NOT NULL,
    name_en VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    
    UNIQUE(curriculum_version_id, phase_id, code)
);

CREATE INDEX idx_elements_curriculum_version_id ON elements(curriculum_version_id);
CREATE INDEX idx_elements_phase_id ON elements(phase_id);
CREATE INDEX idx_elements_code ON elements(code);
```

### Subelements Table

```sql
CREATE TABLE subelements (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    curriculum_version_id UUID NOT NULL REFERENCES curriculum_versions(id) ON DELETE CASCADE,
    element_id UUID NOT NULL REFERENCES elements(id) ON DELETE CASCADE,
    code VARCHAR(50) NOT NULL,
    name VARCHAR(255) NOT NULL,
    name_en VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    
    UNIQUE(curriculum_version_id, element_id, code)
);

CREATE INDEX idx_subelements_curriculum_version_id ON subelements(curriculum_version_id);
CREATE INDEX idx_subelements_element_id ON subelements(element_id);
CREATE INDEX idx_subelements_code ON subelements(code);
```

### CP Table

```sql
CREATE TABLE cp (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    curriculum_version_id UUID NOT NULL REFERENCES curriculum_versions(id) ON DELETE CASCADE,
    subelement_id UUID NOT NULL REFERENCES subelements(id) ON DELETE CASCADE,
    code VARCHAR(50) NOT NULL,
    text TEXT NOT NULL,
    text_en TEXT,
    fase VARCHAR(10),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    
    UNIQUE(curriculum_version_id, subelement_id, code)
);

CREATE INDEX idx_cp_curriculum_version_id ON cp(curriculum_version_id);
CREATE INDEX idx_cp_subelement_id ON cp(subelement_id);
CREATE INDEX idx_cp_code ON cp(code);
CREATE INDEX idx_cp_fase ON cp(fase);
CREATE INDEX idx_cp_text_search ON cp USING gin(to_tsvector('indonesian', text));
CREATE INDEX idx_cp_text_en_search ON cp USING gin(to_tsvector('english', text_en));
```

## Entity Relationships

```
┌─────────────────────┐
│ curriculum_versions │
│─────────────────────│
│ id (PK)             │
│ code                │
│ effective_year      │
│ status              │
│ ...                 │
└──────┬──────────────┘
       │ 1
       │
       │ N
┌──────┴──────────────┐
│     subjects        │
│─────────────────────│
│ id (PK)             │
│ curriculum_version_id│
│ code                │
│ name                │
│ ...                 │
└──────┬──────────────┘
       │ 1
       │
       │ N
┌──────┴──────────────┐
│      phases         │
│─────────────────────│
│ id (PK)             │
│ curriculum_version_id│
│ subject_id          │
│ code                │
│ ...                 │
└──────┬──────────────┘
       │ 1
       │
       │ N
┌──────┴──────────────┐
│     elements        │
│─────────────────────│
│ id (PK)             │
│ curriculum_version_id│
│ phase_id            │
│ code                │
│ ...                 │
└──────┬──────────────┘
       │ 1
       │
       │ N
┌──────┴──────────────┐
│   subelements      │
│─────────────────────│
│ id (PK)             │
│ curriculum_version_id│
│ element_id          │
│ code                │
│ ...                 │
└──────┬──────────────┘
       │ 1
       │
       │ N
┌──────┴──────────────┐
│        cp           │
│─────────────────────│
│ id (PK)             │
│ curriculum_version_id│
│ subelement_id       │
│ code                │
│ text                │
│ ...                 │
└─────────────────────┘
```

## Data Access Layer

### Repository Interface

```go
package repository

type CurriculumRepository interface {
    // Curriculum Version
    CreateCurriculumVersion(ctx context.Context, version *CurriculumVersion) error
    GetCurriculumVersion(ctx context.Context, id string) (*CurriculumVersion, error)
    GetActiveCurriculumVersion(ctx context.Context) (*CurriculumVersion, error)
    ListCurriculumVersions(ctx context.Context) ([]*CurriculumVersion, error)
    UpdateCurriculumVersion(ctx context.Context, version *CurriculumVersion) error
    DeleteCurriculumVersion(ctx context.Context, id string) error
    
    // Subject
    CreateSubject(ctx context.Context, subject *Subject) error
    GetSubject(ctx context.Context, id string) (*Subject, error)
    ListSubjects(ctx context.Context, versionID string) ([]*Subject, error)
    
    // Phase
    CreatePhase(ctx context.Context, phase *Phase) error
    GetPhase(ctx context.Context, id string) (*Phase, error)
    ListPhases(ctx context.Context, subjectID string) ([]*Phase, error)
    
    // Element
    CreateElement(ctx context.Context, element *Element) error
    GetElement(ctx context.Context, id string) (*Element, error)
    ListElements(ctx context.Context, phaseID string) ([]*Element, error)
    
    // Subelement
    CreateSubelement(ctx context.Context, subelement *Subelement) error
    GetSubelement(ctx context.Context, id string) (*Subelement, error)
    ListSubelements(ctx context.Context, elementID string) ([]*Subelement, error)
    
    // CP
    CreateCP(ctx context.Context, cp *CP) error
    GetCP(ctx context.Context, id string) (*CP, error)
    ListCPs(ctx context.Context, filters SearchFilters, pagination Pagination) ([]*CP, int, error)
    SearchCPs(ctx context.Context, query string, filters SearchFilters, pagination Pagination) ([]*CP, int, error)
    GetSubjectTree(ctx context.Context, subjectID string) (*SubjectTree, error)
}
```

---

# SECTION 9 — Module Diagram

## Curriculum Module Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                    Curriculum Module                          │
└─────────────────────────────────────────────────────────────┘
                            │
        ┌───────────────────┼───────────────────┐
        │                   │                   │
        ↓                   ↓                   ↓
┌───────────────┐  ┌───────────────┐  ┌───────────────┐
│   Handler     │  │   Service     │  │  Repository   │
│               │  │               │  │               │
│ - Import      │  │ - Import      │  │ - CRUD        │
│ - Read        │  │ - Validate    │  │ - Search      │
│ - Search      │  │ - Search      │  │ - Transaction │
│ - Tree        │  │ - Tree        │  │               │
└───────────────┘  └───────────────┘  └───────────────┘
        │                   │                   │
        └───────────────────┼───────────────────┘
                            ↓
                  ┌───────────────────┐
                  │   Database        │
                  │                   │
                  │ - curriculum_versions│
                  │ - subjects        │
                  │ - phases          │
                  │ - elements        │
                  │ - subelements     │
                  │ - cp              │
                  └───────────────────┘
```

## Component Interactions

```
┌─────────┐       ┌─────────┐       ┌─────────┐       ┌─────────┐
│  Client │       │ Handler │       │ Service │       │ Database│
└────┬────┘       └────┬────┘       └────┬────┘       └────┬────┘
     │                  │                  │                  │
     │ Import Request   │                  │                  │
     ├─────────────────>│                  │                  │
     │                  │                  │                  │
     │                  │ Validate JSON    │                  │
     │                  ├─────────────────>│                  │
     │                  │                  │                  │
     │                  │ Valid            │                  │
     │                  │<─────────────────┤                  │
     │                  │                  │                  │
     │                  │ Import Data      │                  │
     │                  ├─────────────────>│                  │
     │                  │                  │                  │
     │                  │                  │ Transaction      │
     │                  │                  ├─────────────────>│
     │                  │                  │                  │
     │                  │                  │ Insert Entities  │
     │                  │                  ├─────────────────>│
     │                  │                  │                  │
     │                  │                  │ Success          │
     │                  │                  │<─────────────────┤
     │                  │                  │                  │
     │                  │ Success          │                  │
     │                  │<─────────────────┤                  │
     │                  │                  │                  │
     │ Response         │                  │                  │
     │<─────────────────┤                  │                  │
     │                  │                  │                  │
```

---

# SECTION 10 — Validation Rules

## Import Validation

### File Validation

| Rule | Description | Error Message |
|------|-------------|--------------|
| File format | Must be JSON file | Invalid file format. Expected JSON. |
| File size | Maximum 10MB | File size exceeds 10MB limit. |
| JSON structure | Must have curriculum_version and subjects | Invalid JSON structure. Missing required fields. |

### Curriculum Version Validation

| Rule | Description | Error Message |
|------|-------------|--------------|
| Code required | curriculum_version_code must be present | Curriculum version code is required. |
| Code unique | curriculum_version_code must not exist | Curriculum version code already exists. |
| Effective year required | effective_year must be present | Effective year is required. |
| Effective from required | effective_from must be present | Effective from date is required. |
| Effective year format | Must be valid year (YYYY) | Invalid effective year format. |
| Effective from format | Must be valid date (YYYY-MM-DD) | Invalid effective from date format. |

### Subject Validation

| Rule | Description | Error Message |
|------|-------------|--------------|
| Code required | Subject code must be present | Subject code is required. |
| Code unique | Subject code must be unique within version | Subject code already exists in this version. |
| Name required | Subject name must be present | Subject name is required. |
| Code format | Must be uppercase alphanumeric, max 50 chars | Invalid subject code format. |
| Name format | Must be 1-255 characters | Invalid subject name format. |

### Phase Validation

| Rule | Description | Error Message |
|------|-------------|--------------|
| Code required | Phase code must be present | Phase code is required. |
| Code unique | Phase code must be unique within subject | Phase code already exists in this subject. |
| Name required | Phase name must be present | Phase name is required. |
| Grade levels required | At least one grade level must be present | Grade levels are required. |
| Code format | Must be uppercase alphanumeric, max 50 chars | Invalid phase code format. |
| Grade level format | Must be valid grade levels (1-12) | Invalid grade level format. |

### Element Validation

| Rule | Description | Error Message |
|------|-------------|--------------|
| Code required | Element code must be present | Element code is required. |
| Code unique | Element code must be unique within phase | Element code already exists in this phase. |
| Name required | Element name must be present | Element name is required. |
| Code format | Must be uppercase alphanumeric, max 50 chars | Invalid element code format. |

### Subelement Validation

| Rule | Description | Error Message |
|------|-------------|--------------|
| Code required | Subelement code must be present | Subelement code is required. |
| Code unique | Subelement code must be unique within element | Subelement code already exists in this element. |
| Name required | Subelement name must be present | Subelement name is required. |
| Code format | Must be uppercase alphanumeric, max 50 chars | Invalid subelement code format. |

### CP Validation

| Rule | Description | Error Message |
|------|-------------|--------------|
| Code required | CP code must be present | CP code is required. |
| Code unique | CP code must be unique within subelement | CP code already exists in this subelement. |
| Text required | CP text must be present | CP text is required. |
| Code format | Must be uppercase alphanumeric with dots, max 50 chars | Invalid CP code format. |
| Text format | Must be 1-5000 characters | Invalid CP text format. |
| Fase format | Must be valid phase (A, B, C, D, E, F) | Invalid fase format. |

## Validation Implementation

### Validator

```go
package curriculum

import (
    "regexp"
    "strings"
)

type Validator struct{}

func NewValidator() *Validator {
    return &Validator{}
}

func (v *Validator) ValidateImportData(data *CurriculumImportData) error {
    // Validate curriculum version
    if err := v.validateCurriculumVersion(&data.CurriculumVersion); err != nil {
        return err
    }
    
    // Validate subjects
    for _, subject := range data.Subjects {
        if err := v.validateSubject(&subject); err != nil {
            return err
        }
        
        // Validate phases
        for _, phase := range subject.Phases {
            if err := v.validatePhase(&phase); err != nil {
                return err
            }
            
            // Validate elements
            for _, element := range phase.Elements {
                if err := v.validateElement(&element); err != nil {
                    return err
                }
                
                // Validate subelements
                for _, subelement := range element.Subelements {
                    if err := v.validateSubelement(&subelement); err != nil {
                        return err
                    }
                    
                    // Validate CPs
                    for _, cp := range subelement.CPs {
                        if err := v.validateCP(&cp); err != nil {
                            return err
                        }
                    }
                }
            }
        }
    }
    
    return nil
}

func (v *Validator) validateCurriculumVersion(cv *CurriculumVersionData) error {
    if cv.Code == "" {
        return errors.New("curriculum version code is required")
    }
    
    if !v.isValidCode(cv.Code) {
        return errors.New("invalid curriculum version code format")
    }
    
    if cv.EffectiveYear == 0 {
        return errors.New("effective year is required")
    }
    
    if cv.EffectiveYear < 2000 || cv.EffectiveYear > 2100 {
        return errors.New("invalid effective year")
    }
    
    if cv.EffectiveFrom == "" {
        return errors.New("effective from date is required")
    }
    
    return nil
}

func (v *Validator) validateSubject(subject *SubjectData) error {
    if subject.Code == "" {
        return errors.New("subject code is required")
    }
    
    if !v.isValidCode(subject.Code) {
        return errors.New("invalid subject code format")
    }
    
    if subject.Name == "" {
        return errors.New("subject name is required")
    }
    
    if len(subject.Name) > 255 {
        return errors.New("subject name too long")
    }
    
    return nil
}

func (v *Validator) validatePhase(phase *PhaseData) error {
    if phase.Code == "" {
        return errors.New("phase code is required")
    }
    
    if !v.isValidCode(phase.Code) {
        return errors.New("invalid phase code format")
    }
    
    if phase.Name == "" {
        return errors.New("phase name is required")
    }
    
    if len(phase.GradeLevels) == 0 {
        return errors.New("grade levels are required")
    }
    
    for _, grade := range phase.GradeLevels {
        if !v.isValidGradeLevel(grade) {
            return errors.New("invalid grade level")
        }
    }
    
    return nil
}

func (v *Validator) validateElement(element *ElementData) error {
    if element.Code == "" {
        return errors.New("element code is required")
    }
    
    if !v.isValidCode(element.Code) {
        return errors.New("invalid element code format")
    }
    
    if element.Name == "" {
        return errors.New("element name is required")
    }
    
    return nil
}

func (v *Validator) validateSubelement(subelement *SubelementData) error {
    if subelement.Code == "" {
        return errors.New("subelement code is required")
    }
    
    if !v.isValidCode(subelement.Code) {
        return errors.New("invalid subelement code format")
    }
    
    if subelement.Name == "" {
        return errors.New("subelement name is required")
    }
    
    return nil
}

func (v *Validator) validateCP(cp *CPData) error {
    if cp.Code == "" {
        return errors.New("CP code is required")
    }
    
    if !v.isValidCPCode(cp.Code) {
        return errors.New("invalid CP code format")
    }
    
    if cp.Text == "" {
        return errors.New("CP text is required")
    }
    
    if len(cp.Text) > 5000 {
        return errors.New("CP text too long")
    }
    
    if cp.Fase != "" && !v.isValidFase(cp.Fase) {
        return errors.New("invalid fase format")
    }
    
    return nil
}

func (v *Validator) isValidCode(code string) bool {
    matched, _ := regexp.MatchString(`^[A-Z0-9]{1,50}$`, code)
    return matched
}

func (v *Validator) isValidCPCode(code string) bool {
    matched, _ := regexp.MatchString(`^[A-Z0-9\.]{1,50}$`, code)
    return matched
}

func (v *Validator) isValidGradeLevel(grade string) bool {
    validGrades := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}
    for _, validGrade := range validGrades {
        if grade == validGrade {
            return true
        }
    }
    return false
}

func (v *Validator) isValidFase(fase string) bool {
    validFases := []string{"A", "B", "C", "D", "E", "F"}
    for _, validFase := range validFases {
        if fase == validFase {
            return true
        }
    }
    return false
}
```

---

# SECTION 11 — Appendix

## Performance Considerations

### Database Indexes

- All foreign keys indexed for JOIN performance
- Code fields indexed for lookups
- Status indexed for filtering
- Full-text search indexes on CP text fields

### Query Optimization

- Use JOINs with proper indexes
- Implement pagination for large result sets
- Cache active curriculum version
- Use materialized views for complex queries (future)

### Caching Strategy

- Cache active curriculum version in memory
- Cache subject trees for frequently accessed subjects
- Implement Redis caching for search results (future)

## Testing

### Unit Tests

- Test validation rules
- Test repository CRUD operations
- Test service business logic

### Integration Tests

- Test complete import flow
- Test search functionality
- Test permission enforcement

### Performance Tests

- Test import performance with large datasets
- Test search performance with large CP sets
- Test concurrent read operations

## Future Enhancements

### Wave 2

- Manual CRUD for curriculum administrators
- Curriculum diff viewer
- Curriculum comparison between versions
- Bulk CP editing
- Curriculum export to multiple formats
- Curriculum approval workflow
- Curriculum change notifications

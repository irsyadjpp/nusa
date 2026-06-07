# 13_API_CONTRACT.md

## Foundation Document for NUSA Education Platform

**Version**: 1.0
**Date**: June 2026
**Status**: FOUNDATION DOCUMENT
**Alignment**: Validated against Foundation Architecture (00A, 00B, 00C, 01, 02, 03, 04, 05, 06, 07, 08, 09, 10, 11, 12)

**Purpose**: Define the complete MVP API catalog for NUSA Wave 1, serving as the official API contract between backend and frontend. This document is implementation-ready and convertible to OpenAPI/Swagger specification.

---

# SECTION 1 — API Overview

## API Base URL

**Development**: `http://localhost:8080/api/v1`
**Production**: `https://api.nusa.id/v1`

## API Versioning

- Current Version: `v1`
- Versioning Strategy: URL Path Versioning
- Backward Compatibility: Maintained within major version

## Authentication

All API endpoints (except authentication endpoints) require:
- **Authentication**: JWT Bearer Token in `Authorization` header
- **Token Format**: `Bearer <access_token>`
- **Token Expiration**: 24 hours for access token, 7 days for refresh token

## Response Format

All responses follow standard format:

```json
{
  "success": true,
  "data": { },
  "message": "Success message",
  "timestamp": "2026-06-03T12:00:00Z"
}
```

Error response format:

```json
{
  "success": false,
  "error": {
    "code": "ERROR_CODE",
    "message": "Error message",
    "details": { }
  },
  "timestamp": "2026-06-03T12:00:00Z"
}
```

---

# API Exposure Matrix

This matrix defines which database fields are exposed in API responses versus internal system fields.

## Artifact Entities

### TP (Teaching Plan)

| Database Field | API Exposure | Rationale |
|----------------|--------------|-----------|
| id | A. Exposed | Primary identifier for frontend operations |
| cp_id | A. Exposed | Required for CP reference |
| user_id | A. Exposed | Required for ownership tracking |
| status | A. Exposed | Required for workflow state |
| learning_objectives | A. Exposed | Core artifact content |
| time_allocation | A. Exposed | Core artifact content |
| prerequisites | A. Exposed | Core artifact content |
| ai_confidence_score | A. Exposed | Required for AI governance monitoring |
| ai_generated_at | A. Exposed | Required for AI governance monitoring |
| ai_agent_version | A. Exposed | Required for AI governance monitoring |
| version_no | A. Exposed | Required for version tracking |
| is_current_version | A. Exposed | Required for version filtering |
| parent_version_id | A. Exposed | Required for version history |
| created_at | A. Exposed | Required for audit trail |
| updated_at | A. Exposed | Required for audit trail |
| approved_at | A. Exposed | Required for approval workflow |
| approved_by | A. Exposed | Required for approval workflow |

### ATP (Annual Teaching Plan)

| Database Field | API Exposure | Rationale |
|----------------|--------------|-----------|
| id | A. Exposed | Primary identifier for frontend operations |
| tp_id | A. Exposed | Required for TP reference |
| user_id | A. Exposed | Required for ownership tracking |
| status | A. Exposed | Required for workflow state |
| academic_calendar | A. Exposed | Core artifact content |
| class_schedule | A. Exposed | Core artifact content |
| weekly_sequence | A. Exposed | Core artifact content |
| assessment_schedule | A. Exposed | Core artifact content |
| ai_confidence_score | A. Exposed | Required for AI governance monitoring |
| ai_generated_at | A. Exposed | Required for AI governance monitoring |
| ai_agent_version | A. Exposed | Required for AI governance monitoring |
| version_no | A. Exposed | Required for version tracking |
| is_current_version | A. Exposed | Required for version filtering |
| parent_version_id | A. Exposed | Required for version history |
| created_at | A. Exposed | Required for audit trail |
| updated_at | A. Exposed | Required for audit trail |
| approved_at | A. Exposed | Required for approval workflow |
| approved_by | A. Exposed | Required for approval workflow |

### Modul Ajar

| Database Field | API Exposure | Rationale |
|----------------|--------------|-----------|
| id | A. Exposed | Primary identifier for frontend operations |
| atp_id | A. Exposed | Required for ATP reference |
| week | A. Exposed | Required for scheduling |
| topic | A. Exposed | Core artifact content |
| resources | A. Exposed | Core artifact content |
| class_characteristics | A. Exposed | Core artifact content |
| learning_activities | A. Exposed | Core artifact content |
| resource_requirements | A. Exposed | Core artifact content |
| assessment_methods | A. Exposed | Core artifact content |
| status | A. Exposed | Required for workflow state |
| ai_confidence_score | A. Exposed | Required for AI governance monitoring |
| ai_generated_at | A. Exposed | Required for AI governance monitoring |
| ai_agent_version | A. Exposed | Required for AI governance monitoring |
| version_no | A. Exposed | Required for version tracking |
| is_current_version | A. Exposed | Required for version filtering |
| parent_version_id | A. Exposed | Required for version history |
| created_at | A. Exposed | Required for audit trail |
| updated_at | A. Exposed | Required for audit trail |
| approved_at | A. Exposed | Required for approval workflow |
| approved_by | A. Exposed | Required for approval workflow |

### Assessment

| Database Field | API Exposure | Rationale |
|----------------|--------------|-----------|
| id | A. Exposed | Primary identifier for frontend operations |
| modul_ajar_id | A. Exposed | Required for Modul Ajar reference |
| user_id | A. Exposed | Required for ownership tracking |
| assessment_type | A. Exposed | Required for categorization |
| status | A. Exposed | Required for workflow state |
| assessment_items | A. Exposed | Core artifact content |
| answer_key | A. Exposed | Core artifact content |
| scoring_guidelines | A. Exposed | Core artifact content |
| ai_confidence_score | A. Exposed | Required for AI governance monitoring |
| ai_generated_at | A. Exposed | Required for AI governance monitoring |
| ai_agent_version | A. Exposed | Required for AI governance monitoring |
| version_no | A. Exposed | Required for version tracking |
| is_current_version | A. Exposed | Required for version filtering |
| parent_version_id | A. Exposed | Required for version history |
| created_at | A. Exposed | Required for audit trail |
| updated_at | A. Exposed | Required for audit trail |
| approved_at | A. Exposed | Required for approval workflow |
| approved_by | A. Exposed | Required for approval workflow |

### Rubric

| Database Field | API Exposure | Rationale |
|----------------|--------------|-----------|
| id | A. Exposed | Primary identifier for frontend operations |
| assessment_id | A. Exposed | Required for Assessment reference |
| user_id | A. Exposed | Required for ownership tracking |
| rubric_type | A. Exposed | Required for categorization |
| status | A. Exposed | Required for workflow state |
| performance_criteria | A. Exposed | Core artifact content |
| performance_levels | A. Exposed | Core artifact content |
| scoring_guidelines | A. Exposed | Core artifact content |
| ai_confidence_score | A. Exposed | Required for AI governance monitoring |
| ai_generated_at | A. Exposed | Required for AI governance monitoring |
| ai_agent_version | A. Exposed | Required for AI governance monitoring |
| version_no | A. Exposed | Required for version tracking |
| is_current_version | A. Exposed | Required for version filtering |
| parent_version_id | A. Exposed | Required for version history |
| created_at | A. Exposed | Required for audit trail |
| updated_at | A. Exposed | Required for audit trail |
| approved_at | A. Exposed | Required for approval workflow |
| approved_by | A. Exposed | Required for approval workflow |

### Narrative Report

| Database Field | API Exposure | Rationale |
|----------------|--------------|-----------|
| id | A. Exposed | Primary identifier for frontend operations |
| student_id | A. Exposed | Required for student reference |
| user_id | A. Exposed | Required for ownership tracking |
| status | A. Exposed | Required for workflow state |
| report_period | A. Exposed | Core artifact content |
| language | A. Exposed | Required for localization |
| content | A. Exposed | Core artifact content |
| ai_confidence_score | A. Exposed | Required for AI governance monitoring |
| ai_generated_at | A. Exposed | Required for AI governance monitoring |
| ai_agent_version | A. Exposed | Required for AI governance monitoring |
| version_no | A. Exposed | Required for version tracking |
| is_current_version | A. Exposed | Required for version filtering |
| parent_version_id | A. Exposed | Required for version history |
| created_at | A. Exposed | Required for audit trail |
| updated_at | A. Exposed | Required for audit trail |
| approved_at | A. Exposed | Required for approval workflow |
| approved_by | A. Exposed | Required for approval workflow |

## HTTP Status Codes

- `200 OK`: Successful request
- `201 Created`: Resource created successfully
- `204 No Content`: Successful request with no response body
- `400 Bad Request`: Invalid request parameters
- `401 Unauthorized`: Authentication required or failed
- `403 Forbidden`: Authorization failed
- `404 Not Found`: Resource not found
- `409 Conflict`: Resource conflict
- `422 Unprocessable Entity`: Validation failed
- `500 Internal Server Error`: Server error

---

# SECTION 2 — Authentication APIs

## Overview

Authentication APIs handle user login, token refresh, and logout operations.

---

### API 2.1: User Login

**Endpoint Name**: User Login

**HTTP Method**: `POST`

**URL**: `/api/v1/auth/login`

**Authentication Requirement**: None

**Authorization Requirement**: None

**Request Schema**:

```json
{
  "email": "teacher@school.id",
  "password": "securePassword123"
}
```

**Validation Rules**:
- `email`: Required, valid email format, max 255 characters
- `password`: Required, min 8 characters, max 128 characters

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "user": {
      "id": "usr_1234567890",
      "email": "teacher@school.id",
      "name": "John Doe",
      "role": "TEACHER"
    },
    "tokens": {
      "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
      "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
      "token_type": "Bearer",
      "expires_in": 86400
    }
  },
  "message": "Login successful",
  "timestamp": "2026-06-03T12:00:00Z"
}
```

**Response Schema** (Error):

```json
{
  "success": false,
  "error": {
    "code": "INVALID_CREDENTIALS",
    "message": "Invalid email or password",
    "details": { }
  },
  "timestamp": "2026-06-03T12:00:00Z"
}
```

---

### API 2.2: Refresh Token

**Endpoint Name**: Refresh Token

**HTTP Method**: `POST`

**URL**: `/api/v1/auth/refresh`

**Authentication Requirement**: None

**Authorization Requirement**: None

**Request Schema**:

```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**Validation Rules**:
- `refresh_token`: Required, valid JWT format

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "token_type": "Bearer",
    "expires_in": 86400
  },
  "message": "Token refreshed successfully",
  "timestamp": "2026-06-03T12:00:00Z"
}
```

**Response Schema** (Error):

```json
{
  "success": false,
  "error": {
    "code": "INVALID_REFRESH_TOKEN",
    "message": "Invalid or expired refresh token",
    "details": { }
  },
  "timestamp": "2026-06-03T12:00:00Z"
}
```

---

### API 2.3: User Logout

**Endpoint Name**: User Logout

**HTTP Method**: `POST`

**URL**: `/api/v1/auth/logout`

**Authentication Requirement**: Required

**Authorization Requirement**: Required

**Request Schema**: None (empty body)

**Validation Rules**: None

**Response Schema** (Success):

```json
{
  "success": true,
  "data": { },
  "message": "Logout successful",
  "timestamp": "2026-06-03T12:00:00Z"
}
```

**Response Schema** (Error):

```json
{
  "success": false,
  "error": {
    "code": "LOGOUT_FAILED",
    "message": "Failed to logout",
    "details": { }
  },
  "timestamp": "2026-06-03T12:00:00Z"
}
```

---

# SECTION 3 — Curriculum APIs

## Overview

Curriculum APIs handle national curriculum viewing and Teaching Plan (TP) generation and management. The curriculum follows a hierarchical structure: Subject → Phase → Element → Subelement → CP.

---

### API 3.1: List Curriculum Subjects

**Endpoint Name**: List Curriculum Subjects

**HTTP Method**: `GET`

**URL**: `/api/v1/curriculum/subjects`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER, ADMINISTRATOR

**Query Parameters**: None

**Request Schema**: None

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "subjects": [
      {
        "id": "sub_1234567890",
        "code": "MTK",
        "name": "Matematika",
        "name_en": "Mathematics",
        "description": "Mathematics subject",
        "is_active": true
      }
    ]
  },
  "message": "Subjects retrieved successfully",
  "timestamp": "2026-06-05T12:00:00Z"
}
```

---

### API 3.2: List Curriculum Phases

**Endpoint Name**: List Curriculum Phases

**HTTP Method**: `GET`

**URL**: `/api/v1/curriculum/phases`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER, ADMINISTRATOR

**Query Parameters**: None

**Request Schema**: None

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "phases": [
      {
        "id": "phase_1234567890",
        "code": "FASE_A",
        "name": "Fase A",
        "name_en": "Phase A",
        "description": "Phase A for grades 1-2",
        "grade_level_start": 1,
        "grade_level_end": 2,
        "is_active": true
      }
    ]
  },
  "message": "Phases retrieved successfully",
  "timestamp": "2026-06-05T12:00:00Z"
}
```

---

### API 3.3: List Curriculum Elements

**Endpoint Name**: List Curriculum Elements

**HTTP Method**: `GET`

**URL**: `/api/v1/curriculum/elements`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER, ADMINISTRATOR

**Query Parameters**:
- `subject_id` (required): Subject identifier
- `phase_id` (required): Phase identifier

**Request Schema**: None (query parameters only)

**Validation Rules**:
- `subject_id`: Required, valid UUID format
- `phase_id`: Required, valid UUID format

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "elements": [
      {
        "id": "elem_1234567890",
        "subject_id": "sub_1234567890",
        "phase_id": "phase_1234567890",
        "code": "NUM_OPS",
        "name": "Number and Operations",
        "name_en": "Number and Operations",
        "description": "Number and operations element",
        "is_active": true
      }
    ]
  },
  "message": "Elements retrieved successfully",
  "timestamp": "2026-06-05T12:00:00Z"
}
```

---

### API 3.4: List Curriculum Subelements

**Endpoint Name**: List Curriculum Subelements

**HTTP Method**: `GET`

**URL**: `/api/v1/curriculum/subelements`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER, ADMINISTRATOR

**Query Parameters**:
- `element_id` (required): Element identifier

**Request Schema**: None (query parameters only)

**Validation Rules**:
- `element_id`: Required, valid UUID format

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "subelements": [
      {
        "id": "subelem_1234567890",
        "element_id": "elem_1234567890",
        "code": "WHOLE_NUM",
        "name": "Whole Numbers",
        "name_en": "Whole Numbers",
        "description": "Whole numbers subelement",
        "is_active": true
      }
    ]
  },
  "message": "Subelements retrieved successfully",
  "timestamp": "2026-06-05T12:00:00Z"
}
```

---

### API 3.5: View National Curriculum Plan (CP)

**Endpoint Name**: View National Curriculum Plan

**HTTP Method**: `GET`

**URL**: `/api/v1/curriculum/cp`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER, ADMINISTRATOR

**Query Parameters**:
- `subject_id` (required): Subject identifier
- `phase_id` (required): Phase identifier
- `element_id` (required): Element identifier
- `subelement_id` (required): Subelement identifier

**Request Schema**: None (query parameters only)

**Validation Rules**:
- `subject_id`: Required, valid UUID format
- `phase_id`: Required, valid UUID format
- `element_id`: Required, valid UUID format
- `subelement_id`: Required, valid UUID format

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "curriculum_hierarchy": {
      "subject": {
        "id": "sub_1234567890",
        "code": "MTK",
        "name": "Matematika",
        "name_en": "Mathematics"
      },
      "phase": {
        "id": "phase_1234567890",
        "code": "FASE_E",
        "name": "Fase E",
        "name_en": "Phase E",
        "grade_level_start": 10,
        "grade_level_end": 12
      },
      "element": {
        "id": "elem_1234567890",
        "code": "ALGEBRA",
        "name": "Algebra",
        "name_en": "Algebra",
        "description": "Algebraic thinking and operations"
      },
      "subelement": {
        "id": "subelem_1234567890",
        "code": "LINEAR_EQ",
        "name": "Linear Equations",
        "name_en": "Linear Equations",
        "description": "Linear equations and inequalities"
      }
    },
    "cp": {
      "id": "cp_1234567890",
      "code": "CP.10.1.1",
      "description": "Students can solve linear equations and inequalities",
      "competency_code": "CP.10.1",
      "learning_objectives": [
        {
          "id": "lo_1234567890",
          "code": "LO.10.1",
          "description": "Students understand algebraic concepts",
          "competency_code": "CP.10.1"
        }
      ],
      "competency_standards": [
        {
          "id": "cs_1234567890",
          "code": "CS.10.1",
          "description": "Algebraic thinking"
        }
      ],
      "time_allocation": {
        "total_hours": 120,
        "hours_per_week": 4
      },
      "version": "2026"
    }
  },
  "message": "Curriculum plan retrieved successfully",
  "timestamp": "2026-06-05T12:00:00Z"
}
```

**Response Schema** (Error):

```json
{
  "success": false,
  "error": {
    "code": "CP_NOT_FOUND",
    "message": "Curriculum plan not found for specified parameters",
    "details": { }
  },
  "timestamp": "2026-06-03T12:00:00Z"
}
```

---

### API 3.6: Generate TP Set with AI

**Endpoint Name**: Generate TP Set with AI

**HTTP Method**: `POST`

**URL**: `/api/v1/curriculum/cp/{cp_id}/tp-sets/generate`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Path Parameters**:
- `cp_id`: Required, valid UUID format - CP identifier

**Request Schema**:

```json
{
  "generation_reason": "Initial generation",
  "class_info": {
    "grade": "10",
    "subject": "Mathematics",
    "academic_year": "2026"
  },
  "teaching_schedule": {
    "hours_per_week": 4,
    "weeks_per_semester": 18
  },
  "preferences": {
    "focus_areas": ["algebra", "geometry"],
    "teaching_style": "interactive"
  }
}
```

**Validation Rules**:
- `generation_reason`: Optional, max 500 characters
- `class_info.grade`: Required, valid grade level
- `class_info.subject`: Required, max 255 characters
- `class_info.academic_year`: Required, valid year format (YYYY)
- `teaching_schedule.hours_per_week`: Required, min 1, max 40
- `teaching_schedule.weeks_per_semester`: Required, min 1, max 52
- `preferences`: Optional, object with optional fields

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "tp_set": {
      "id": "tps_1234567890",
      "cp_id": "cp_1234567890",
      "version_no": 1,
      "status": "DRAFT",
      "generation_source": "AI_GENERATED",
      "generation_reason": "Initial generation",
      "generated_by": "user_1234567890",
      "ai_generation_id": "ai_gen_1234567890",
      "created_at": "2026-06-03T12:00:00Z",
      "updated_at": "2026-06-03T12:00:00Z",
      "tps": [
        {
          "id": "tp_1234567890",
          "tp_set_id": "tps_1234567890",
          "sequence_number": 1,
          "title": "Introduction to Linear Equations",
          "learning_objectives": [
            {
              "id": "tlo_1234567890",
              "cp_objective_id": "lo_1234567890",
              "description": "Students understand algebraic concepts through interactive learning",
              "time_allocation_hours": 8
            }
          ],
          "time_allocation": {
            "total_hours": 36,
            "hours_per_week": 4
          },
          "prerequisites": [],
          "estimated_weeks": 3
        },
        {
          "id": "tp_1234567891",
          "tp_set_id": "tps_1234567890",
          "sequence_number": 2,
          "title": "Solving Linear Equations",
          "learning_objectives": [
            {
              "id": "tlo_1234567891",
              "cp_objective_id": "lo_1234567891",
              "description": "Students apply algebraic concepts to solve problems",
              "time_allocation_hours": 8
            }
          ],
          "time_allocation": {
            "total_hours": 36,
            "hours_per_week": 4
          },
          "prerequisites": [
            {
              "objective_id": "tlo_1234567890",
              "required_for": ["tlo_1234567891"]
            }
          ],
          "estimated_weeks": 4
        }
      ],
      "total_tps": 2,
      "total_hours": 72,
      "estimated_weeks": 7
    }
  }
}
```

---

### API 3.7: List TP Sets for CP

**Endpoint Name**: List TP Sets for CP

**HTTP Method**: `GET`

**URL**: `/api/v1/curriculum/cp/{cp_id}/tp-sets`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Path Parameters**:
- `cp_id`: Required, valid UUID format - CP identifier

**Query Parameters**:
- `status`: Optional, filter by status (DRAFT, UNDER_REVIEW, APPROVED, REJECTED, ARCHIVED)
- `page`: Optional, page number (default: 1)
- `per_page`: Optional, items per page (default: 20, max: 100)

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "tp_sets": [
      {
        "id": "tps_1234567890",
        "cp_id": "cp_1234567890",
        "version_no": 1,
        "status": "APPROVED",
        "generation_source": "AI_GENERATED",
        "generated_by": "user_1234567890",
        "approved_by": "user_1234567890",
        "approved_at": "2026-06-03T13:00:00Z",
        "created_at": "2026-06-03T12:00:00Z",
        "total_tps": 3,
        "total_hours": 108
      }
    ],
    "pagination": {
      "page": 1,
      "per_page": 20,
      "total_count": 1,
      "total_pages": 1
    }
  }
}
```

---

### API 3.8: Get TP Set Details

**Endpoint Name**: Get TP Set Details

**HTTP Method**: `GET`

**URL**: `/api/v1/curriculum/tp-sets/{tp_set_id}`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Path Parameters**:
- `tp_set_id`: Required, valid UUID format - TP Set identifier

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "tp_set": {
      "id": "tps_1234567890",
      "cp_id": "cp_1234567890",
      "version_no": 1,
      "status": "APPROVED",
      "generation_source": "AI_GENERATED",
      "generation_reason": "Initial generation",
      "generated_by": "user_1234567890",
      "ai_generation_id": "ai_gen_1234567890",
      "approved_by": "user_1234567890",
      "approved_at": "2026-06-03T13:00:00Z",
      "created_at": "2026-06-03T12:00:00Z",
      "updated_at": "2026-06-03T13:00:00Z",
      "tps": [
        {
          "id": "tp_1234567890",
          "tp_set_id": "tps_1234567890",
          "sequence_number": 1,
          "title": "Introduction to Linear Equations",
          "learning_objectives": [...],
          "time_allocation": {...},
          "prerequisites": [],
          "estimated_weeks": 3
        }
      ],
      "total_tps": 3,
      "total_hours": 108,
      "estimated_weeks": 9
    }
  }
}
```

---

### API 3.9: Approve TP Set

**Endpoint Name**: Approve TP Set

**HTTP Method**: `POST`

**URL**: `/api/v1/curriculum/tp-sets/{tp_set_id}/approve`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Path Parameters**:
- `tp_set_id`: Required, valid UUID format - TP Set identifier

**Request Schema**:

```json
{
  "comments": "Approved - meets all requirements"
}
```

**Validation Rules**:
- `comments`: Optional, max 1000 characters

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "tp_set": {
      "id": "tps_1234567890",
      "status": "APPROVED",
      "approved_by": "user_1234567890",
      "approved_at": "2026-06-03T13:00:00Z"
    }
  }
}
```

---

### API 3.10: Archive TP Set

**Endpoint Name**: Archive TP Set

**HTTP Method**: `POST`

**URL**: `/api/v1/curriculum/tp-sets/{tp_set_id}/archive`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Path Parameters**:
- `tp_set_id`: Required, valid UUID format - TP Set identifier

**Request Schema**:

```json
{
  "reason": "Superseded by new version"
}
```

**Validation Rules**:
- `reason`: Optional, max 500 characters

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "tp_set": {
      "id": "tps_1234567890",
      "status": "ARCHIVED"
    }
  }
}
```

---

### API 3.11: Get TP Item by ID

**Endpoint Name**: Get TP Item by ID

**HTTP Method**: `GET`

**URL**: `/api/v1/curriculum/tp/{tp_id}`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER, ADMINISTRATOR

**Path Parameters**:
- `tp_id`: Required, valid UUID format - TP Item identifier

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "tp": {
      "id": "tp_1234567890",
      "tp_set_id": "tps_1234567890",
      "sequence_number": 1,
      "cp_id": "cp_1234567890",
      "status": "APPROVED",
      "title": "Introduction to Linear Equations",
      "learning_objectives": [
        {
          "id": "tlo_1234567890",
          "cp_objective_id": "lo_1234567890",
          "description": "Students understand algebraic concepts through interactive learning",
          "time_allocation_hours": 8
        }
      ],
      "time_allocation": {
        "total_hours": 36,
        "hours_per_week": 4
      },
      "prerequisites": [],
      "estimated_weeks": 3,
      "created_at": "2026-06-03T10:00:00Z",
      "updated_at": "2026-06-03T11:00:00Z"
    }
  }
}
```

---

### API 3.12: Update TP Item

**Endpoint Name**: Update TP Item

**HTTP Method**: `PUT`

**URL**: `/api/v1/curriculum/tp/{tp_id}`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Path Parameters**:
- `tp_id`: Required, valid UUID format - TP Item identifier

**Request Schema**:

```json
{
  "title": "Updated title",
  "learning_objectives": [
    {
      "id": "tlo_1234567890",
      "description": "Updated description",
      "time_allocation_hours": 10
    }
  ],
  "time_allocation": {
    "total_hours": 80,
    "hours_per_week": 4
  }
}
```

**Validation Rules**:
- `tp_id`: Required, valid UUID format
- `learning_objectives`: Required, array of learning objectives
- `learning_objectives[].id`: Required, valid UUID format
- `learning_objectives[].description`: Required, max 1000 characters
- `learning_objectives[].time_allocation_hours`: Required, min 1, max 100
- `time_allocation.total_hours`: Required, min 1, max 1000
- `time_allocation.hours_per_week`: Required, min 1, max 40

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "tp": {
      "id": "tp_1234567890",
      "status": "DRAFT",
      "updated_at": "2026-06-03T13:00:00Z"
    }
  },
  "message": "Teaching plan updated successfully",
  "timestamp": "2026-06-03T13:00:00Z"
}
```

**Response Schema** (Error):

```json
{
  "success": false,
  "error": {
    "code": "TP_UPDATE_FAILED",
    "message": "Failed to update teaching plan",
    "details": { }
  },
  "timestamp": "2026-06-03T13:00:00Z"
}
```

---

### API 3.5: Approve Teaching Plan

**Endpoint Name**: Approve Teaching Plan

**HTTP Method**: `POST`

**URL**: `/api/v1/curriculum/tp/{tp_id}/approve`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Path Parameters**:
- `tp_id` (required): Teaching Plan identifier

**Request Schema**: None

**Validation Rules**:
- `tp_id`: Required, valid UUID format

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "tp": {
      "id": "tp_1234567890",
      "status": "APPROVED",
      "approved_at": "2026-06-03T14:00:00Z",
      "approved_by": "usr_1234567890"
    }
  },
  "message": "Teaching plan approved successfully",
  "timestamp": "2026-06-03T14:00:00Z"
}
```

**Response Schema** (Error):

```json
{
  "success": false,
  "error": {
    "code": "TP_APPROVAL_FAILED",
    "message": "Failed to approve teaching plan",
    "details": {
      "reason": "Teaching plan must be in DRAFT status"
    }
  },
  "timestamp": "2026-06-03T14:00:00Z"
}
```

---

### API 3.6: List Teaching Plans

**Endpoint Name**: List Teaching Plans

**HTTP Method**: `GET`

**URL**: `/api/v1/curriculum/tp`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER, ADMINISTRATOR

**Query Parameters**:
- `status` (optional): Filter by status (DRAFT, APPROVED, REJECTED)
- `subject_id` (optional): Filter by subject
- `grade_level` (optional): Filter by grade level
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 20, max: 100)

**Request Schema**: None (query parameters only)

**Validation Rules**:
- `status`: Optional, valid status (DRAFT, APPROVED, REJECTED)
- `subject_id`: Optional, valid UUID format
- `grade_level`: Optional, valid grade level (10, 11, 12)
- `page`: Optional, min 1
- `limit`: Optional, min 1, max 100

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "tps": [
      {
        "id": "tp_1234567890",
        "subject": "Mathematics",
        "grade_level": "10",
        "status": "APPROVED",
        "created_at": "2026-06-03T10:00:00Z"
      }
    ],
    "pagination": {
      "total": 50,
      "page": 1,
      "limit": 20,
      "total_pages": 3
    }
  },
  "message": "Teaching plans retrieved successfully",
  "timestamp": "2026-06-03T12:00:00Z"
}
```

---

# SECTION 4 — Learning Planning APIs

## Overview

Learning Planning APIs handle Annual Teaching Plan (ATP) and Modul Ajar generation and management.

---

### API 4.1: Generate Annual Teaching Plan (ATP) with AI

**Endpoint Name**: Generate Annual Teaching Plan with AI

**HTTP Method**: `POST`

**URL**: `/api/v1/learning-planning/atp/generate`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Request Schema**:

```json
{
  "tp_id": "tp_1234567890",
  "academic_calendar": {
    "start_date": "2026-07-15",
    "end_date": "2026-12-15",
    "holidays": [
      "2026-08-17",
      "2026-12-25"
    ]
  },
  "class_schedule": {
    "days_per_week": 5,
    "periods_per_day": 8,
    "available_hours_per_week": 4
  }
}
```

**Validation Rules**:
- `tp_id`: Required, valid UUID format
- `academic_calendar.start_date`: Required, valid date format (YYYY-MM-DD)
- `academic_calendar.end_date`: Required, valid date format (YYYY-MM-DD)
- `academic_calendar.holidays`: Optional, array of valid dates
- `class_schedule.days_per_week`: Required, min 1, max 7
- `class_schedule.periods_per_day`: Required, min 1, max 12
- `class_schedule.available_hours_per_week`: Required, min 1, max 40

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "atp": {
      "id": "atp_1234567890",
      "tp_id": "tp_1234567890",
      "status": "DRAFT",
      "weekly_sequence": [
        {
          "week": 1,
          "topics": [
            {
              "learning_objective_id": "tlo_1234567890",
              "hours": 4,
              "start_date": "2026-07-15",
              "end_date": "2026-07-19"
            }
          ]
        }
      ],
      "assessment_schedule": [
        {
          "week": 4,
          "type": "FORMATIVE",
          "topics_covered": ["tlo_1234567890", "tlo_1234567891"]
        }
      ],
      "ai_metadata": {
        "confidence_score": 0.89,
        "generated_at": "2026-06-03T12:00:00Z",
        "agent_version": "1.0"
      }
    }
  },
  "message": "Annual teaching plan generated successfully",
  "timestamp": "2026-06-03T12:00:00Z"
}
```

---

### API 4.2: Get Annual Teaching Plan by ID

**Endpoint Name**: Get Annual Teaching Plan by ID

**HTTP Method**: `GET`

**URL**: `/api/v1/learning-planning/atp/{atp_id}`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER, ADMINISTRATOR

**Path Parameters**:
- `atp_id` (required): Annual Teaching Plan identifier

**Request Schema**: None

**Validation Rules**:
- `atp_id`: Required, valid UUID format

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "atp": {
      "id": "atp_1234567890",
      "tp_id": "tp_1234567890",
      "status": "APPROVED",
      "weekly_sequence": [
        {
          "week": 1,
          "topics": [
            {
              "learning_objective_id": "tlo_1234567890",
              "hours": 4,
              "start_date": "2026-07-15",
              "end_date": "2026-07-19"
            }
          ]
        }
      ],
      "created_at": "2026-06-03T10:00:00Z",
      "updated_at": "2026-06-03T11:00:00Z",
      "approved_at": "2026-06-03T12:00:00Z",
      "approved_by": "usr_1234567890"
    }
  },
  "message": "Annual teaching plan retrieved successfully",
  "timestamp": "2026-06-03T12:00:00Z"
}
```

---

### API 4.3: Update Annual Teaching Plan

**Endpoint Name**: Update Annual Teaching Plan

**HTTP Method**: `PUT`

**URL**: `/api/v1/learning-planning/atp/{atp_id}`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Path Parameters**:
- `atp_id` (required): Annual Teaching Plan identifier

**Request Schema**:

```json
{
  "weekly_sequence": [
    {
      "week": 1,
      "topics": [
        {
          "learning_objective_id": "tlo_1234567890",
          "hours": 5,
          "start_date": "2026-07-15",
          "end_date": "2026-07-19"
        }
      ]
    }
  ]
}
```

**Validation Rules**:
- `atp_id`: Required, valid UUID format
- `weekly_sequence`: Required, array of weekly sequences
- `weekly_sequence[].week`: Required, min 1
- `weekly_sequence[].topics`: Required, array of topics
- `weekly_sequence[].topics[].learning_objective_id`: Required, valid UUID format
- `weekly_sequence[].topics[].hours`: Required, min 1, max 40
- `weekly_sequence[].topics[].start_date`: Required, valid date format
- `weekly_sequence[].topics[].end_date`: Required, valid date format

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "atp": {
      "id": "atp_1234567890",
      "status": "DRAFT",
      "updated_at": "2026-06-03T13:00:00Z"
    }
  },
  "message": "Annual teaching plan updated successfully",
  "timestamp": "2026-06-03T13:00:00Z"
}
```

---

### API 4.4: Approve Annual Teaching Plan

**Endpoint Name**: Approve Annual Teaching Plan

**HTTP Method**: `POST`

**URL**: `/api/v1/learning-planning/atp/{atp_id}/approve`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Path Parameters**:
- `atp_id` (required): Annual Teaching Plan identifier

**Request Schema**: None

**Validation Rules**:
- `atp_id`: Required, valid UUID format

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "atp": {
      "id": "atp_1234567890",
      "status": "APPROVED",
      "approved_at": "2026-06-03T14:00:00Z",
      "approved_by": "usr_1234567890"
    }
  },
  "message": "Annual teaching plan approved successfully",
  "timestamp": "2026-06-03T14:00:00Z"
}
```

---

### API 4.5: Generate Modul Ajar with AI

**Endpoint Name**: Generate Modul Ajar with AI

**HTTP Method**: `POST`

**URL**: `/api/v1/learning-planning/modul-ajar/generate`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Request Schema**:

```json
{
  "atp_id": "atp_1234567890",
  "week": 1,
  "topic": {
    "learning_objective_id": "tlo_1234567890",
    "title": "Introduction to Algebra"
  },
  "resources": {
    "textbooks": ["Algebra Fundamentals"],
    "materials": ["graph paper", "calculators"]
  },
  "class_characteristics": {
    "student_count": 30,
    "skill_level": "intermediate"
  }
}
```

**Validation Rules**:
- `atp_id`: Required, valid UUID format
- `week`: Required, min 1
- `topic.learning_objective_id`: Required, valid UUID format
- `topic.title`: Required, max 255 characters
- `resources.textbooks`: Optional, array of strings
- `resources.materials`: Optional, array of strings
- `class_characteristics.student_count`: Optional, min 1
- `class_characteristics.skill_level`: Optional, valid skill level

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "modul_ajar": {
      "id": "ma_1234567890",
      "atp_id": "atp_1234567890",
      "week": 1,
      "status": "DRAFT",
      "learning_activities": [
        {
          "id": "la_1234567890",
          "sequence": 1,
          "activity_type": "INTRODUCTION",
          "description": "Teacher introduces algebraic concepts with real-world examples",
          "duration_minutes": 15,
          "resources": ["whiteboard", "markers"]
        },
        {
          "id": "la_1234567891",
          "sequence": 2,
          "activity_type": "PRACTICE",
          "description": "Students practice solving algebraic equations",
          "duration_minutes": 30,
          "resources": ["worksheets", "calculators"]
        }
      ],
      "resource_requirements": [
        {
          "resource": "graph paper",
          "quantity": 30
        }
      ],
      "assessment_methods": [
        {
          "type": "FORMATIVE",
          "description": "Observe student participation in practice activities"
        }
      ],
      "ai_metadata": {
        "confidence_score": 0.91,
        "generated_at": "2026-06-03T12:00:00Z",
        "agent_version": "1.0"
      }
    }
  },
  "message": "Modul Ajar generated successfully",
  "timestamp": "2026-06-03T12:00:00Z"
}
```

---

### API 4.6: Get Modul Ajar by ID

**Endpoint Name**: Get Modul Ajar by ID

**HTTP Method**: `GET`

**URL**: `/api/v1/learning-planning/modul-ajar/{modul_ajar_id}`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER, ADMINISTRATOR

**Path Parameters**:
- `modul_ajar_id` (required): Modul Ajar identifier

**Request Schema**: None

**Validation Rules**:
- `modul_ajar_id`: Required, valid UUID format

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "modul_ajar": {
      "id": "ma_1234567890",
      "atp_id": "atp_1234567890",
      "week": 1,
      "status": "APPROVED",
      "learning_activities": [
        {
          "id": "la_1234567890",
          "sequence": 1,
          "activity_type": "INTRODUCTION",
          "description": "Teacher introduces algebraic concepts with real-world examples",
          "duration_minutes": 15,
          "resources": ["whiteboard", "markers"]
        }
      ],
      "created_at": "2026-06-03T10:00:00Z",
      "updated_at": "2026-06-03T11:00:00Z",
      "approved_at": "2026-06-03T12:00:00Z",
      "approved_by": "usr_1234567890"
    }
  },
  "message": "Modul Ajar retrieved successfully",
  "timestamp": "2026-06-03T12:00:00Z"
}
```

---

### API 4.7: Update Modul Ajar

**Endpoint Name**: Update Modul Ajar

**HTTP Method**: `PUT`

**URL**: `/api/v1/learning-planning/modul-ajar/{modul_ajar_id}`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Path Parameters**:
- `modul_ajar_id` (required): Modul Ajar identifier

**Request Schema**:

```json
{
  "learning_activities": [
    {
      "id": "la_1234567890",
      "sequence": 1,
      "activity_type": "INTRODUCTION",
      "description": "Updated description",
      "duration_minutes": 20,
      "resources": ["whiteboard", "markers", "projector"]
    }
  ]
}
```

**Validation Rules**:
- `modul_ajar_id`: Required, valid UUID format
- `learning_activities`: Required, array of learning activities
- `learning_activities[].id`: Required, valid UUID format
- `learning_activities[].sequence`: Required, min 1
- `learning_activities[].activity_type`: Required, valid activity type
- `learning_activities[].description`: Required, max 1000 characters
- `learning_activities[].duration_minutes`: Required, min 1, max 180
- `learning_activities[].resources`: Optional, array of strings

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "modul_ajar": {
      "id": "ma_1234567890",
      "status": "DRAFT",
      "updated_at": "2026-06-03T13:00:00Z"
    }
  },
  "message": "Modul Ajar updated successfully",
  "timestamp": "2026-06-03T13:00:00Z"
}
```

---

### API 4.8: Approve Modul Ajar

**Endpoint Name**: Approve Modul Ajar

**HTTP Method**: `POST`

**URL**: `/api/v1/learning-planning/modul-ajar/{modul_ajar_id}/approve`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Path Parameters**:
- `modul_ajar_id` (required): Modul Ajar identifier

**Request Schema**: None

**Validation Rules**:
- `modul_ajar_id`: Required, valid UUID format

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "modul_ajar": {
      "id": "ma_1234567890",
      "status": "APPROVED",
      "approved_at": "2026-06-03T14:00:00Z",
      "approved_by": "usr_1234567890"
    }
  },
  "message": "Modul Ajar approved successfully",
  "timestamp": "2026-06-03T14:00:00Z"
}
```

---

# SECTION 5 — Assessment APIs

## Overview

Assessment APIs handle assessment generation, rubric generation, and evidence collection.

---

### API 5.1: Generate Assessment with AI

**Endpoint Name**: Generate Assessment with AI

**HTTP Method**: `POST`

**URL**: `/api/v1/assessment/generate`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Request Schema**:

```json
{
  "modul_ajar_id": "ma_1234567890",
  "assessment_type": "FORMATIVE",
  "question_count": 10,
  "difficulty_level": "MEDIUM",
  "time_allocation_minutes": 45
}
```

**Validation Rules**:
- `modul_ajar_id`: Required, valid UUID format
- `assessment_type`: Required, valid type (FORMATIVE, SUMMATIVE)
- `question_count`: Required, min 1, max 50
- `difficulty_level`: Required, valid level (EASY, MEDIUM, HARD)
- `time_allocation_minutes`: Required, min 5, max 180

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "assessment": {
      "id": "ass_1234567890",
      "modul_ajar_id": "ma_1234567890",
      "status": "DRAFT",
      "assessment_type": "FORMATIVE",
      "assessment_items": [
        {
          "id": "ai_1234567890",
          "type": "MULTIPLE_CHOICE",
          "question": "What is the value of x in the equation 2x + 5 = 15?",
          "options": [
            "x = 5",
            "x = 10",
            "x = 15",
            "x = 20"
          ],
          "correct_answer": "x = 5",
          "points": 2,
          "learning_objective_id": "tlo_1234567890"
        },
        {
          "id": "ai_1234567891",
          "type": "ESSAY",
          "question": "Explain the process of solving algebraic equations step by step.",
          "max_points": 10,
          "learning_objective_id": "tlo_1234567890"
        }
      ],
      "answer_key": {
        "ai_1234567890": "x = 5",
        "ai_1234567891": "Rubric-based evaluation"
      },
      "scoring_guidelines": {
        "total_points": 12,
        "passing_score": 7
      },
      "ai_metadata": {
        "confidence_score": 0.88,
        "generated_at": "2026-06-03T12:00:00Z",
        "agent_version": "1.0"
      }
    }
  },
  "message": "Assessment generated successfully",
  "timestamp": "2026-06-03T12:00:00Z"
}
```

---

### API 5.2: Get Assessment by ID

**Endpoint Name**: Get Assessment by ID

**HTTP Method**: `GET`

**URL**: `/api/v1/assessment/{assessment_id}`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER, ADMINISTRATOR

**Path Parameters**:
- `assessment_id` (required): Assessment identifier

**Request Schema**: None

**Validation Rules**:
- `assessment_id`: Required, valid UUID format

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "assessment": {
      "id": "ass_1234567890",
      "modul_ajar_id": "ma_1234567890",
      "status": "APPROVED",
      "assessment_type": "FORMATIVE",
      "assessment_items": [
        {
          "id": "ai_1234567890",
          "type": "MULTIPLE_CHOICE",
          "question": "What is the value of x in the equation 2x + 5 = 15?",
          "options": ["x = 5", "x = 10", "x = 15", "x = 20"],
          "correct_answer": "x = 5",
          "points": 2
        }
      ],
      "created_at": "2026-06-03T10:00:00Z",
      "updated_at": "2026-06-03T11:00:00Z",
      "approved_at": "2026-06-03T12:00:00Z",
      "approved_by": "usr_1234567890"
    }
  },
  "message": "Assessment retrieved successfully",
  "timestamp": "2026-06-03T12:00:00Z"
}
```

---

### API 5.3: Update Assessment

**Endpoint Name**: Update Assessment

**HTTP Method**: `PUT`

**URL**: `/api/v1/assessment/{assessment_id}`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Path Parameters**:
- `assessment_id` (required): Assessment identifier

**Request Schema**:

```json
{
  "assessment_items": [
    {
      "id": "ai_1234567890",
      "type": "MULTIPLE_CHOICE",
      "question": "Updated question",
      "options": ["x = 5", "x = 10", "x = 15", "x = 20"],
      "correct_answer": "x = 5",
      "points": 3
    }
  ]
}
```

**Validation Rules**:
- `assessment_id`: Required, valid UUID format
- `assessment_items`: Required, array of assessment items
- `assessment_items[].id`: Required, valid UUID format
- `assessment_items[].type`: Required, valid type (MULTIPLE_CHOICE, ESSAY, SHORT_ANSWER)
- `assessment_items[].question`: Required, max 1000 characters
- `assessment_items[].options`: Optional, array of strings (required for MULTIPLE_CHOICE)
- `assessment_items[].correct_answer`: Required, max 500 characters
- `assessment_items[].points`: Required, min 1, max 20

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "assessment": {
      "id": "ass_1234567890",
      "status": "DRAFT",
      "updated_at": "2026-06-03T13:00:00Z"
    }
  },
  "message": "Assessment updated successfully",
  "timestamp": "2026-06-03T13:00:00Z"
}
```

---

### API 5.4: Approve Assessment

**Endpoint Name**: Approve Assessment

**HTTP Method**: `POST`

**URL**: `/api/v1/assessment/{assessment_id}/approve`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Path Parameters**:
- `assessment_id` (required): Assessment identifier

**Request Schema**: None

**Validation Rules**:
- `assessment_id`: Required, valid UUID format

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "assessment": {
      "id": "ass_1234567890",
      "status": "APPROVED",
      "approved_at": "2026-06-03T14:00:00Z",
      "approved_by": "usr_1234567890"
    }
  },
  "message": "Assessment approved successfully",
  "timestamp": "2026-06-03T14:00:00Z"
}
```

---

### API 5.5: Generate Rubric with AI

**Endpoint Name**: Generate Rubric with AI

**HTTP Method**: `POST`

**URL**: `/api/v1/assessment/rubric/generate`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Request Schema**:

```json
{
  "assessment_id": "ass_1234567890",
  "rubric_type": "ANALYTIC",
  "performance_levels": 4,
  "criteria_categories": ["content", "organization", "mechanics"]
}
```

**Validation Rules**:
- `assessment_id`: Required, valid UUID format
- `rubric_type`: Required, valid type (ANALYTIC, HOLISTIC)
- `performance_levels`: Required, min 2, max 5
- `criteria_categories`: Required, array of strings, min 1

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "rubric": {
      "id": "rub_1234567890",
      "assessment_id": "ass_1234567890",
      "status": "DRAFT",
      "rubric_type": "ANALYTIC",
      "performance_criteria": [
        {
          "id": "pc_1234567890",
          "category": "content",
          "description": "Accuracy and completeness of content",
          "weight": 0.4
        },
        {
          "id": "pc_1234567891",
          "category": "organization",
          "description": "Structure and flow of response",
          "weight": 0.3
        },
        {
          "id": "pc_1234567892",
          "category": "mechanics",
          "description": "Grammar, spelling, and punctuation",
          "weight": 0.3
        }
      ],
      "performance_levels": [
        {
          "level": 4,
          "label": "Excellent",
          "description": "Demonstrates complete understanding with no errors"
        },
        {
          "level": 3,
          "label": "Proficient",
          "description": "Demonstrates good understanding with minor errors"
        },
        {
          "level": 2,
          "label": "Developing",
          "description": "Demonstrates partial understanding with some errors"
        },
        {
          "level": 1,
          "label": "Beginning",
          "description": "Demonstrates limited understanding with significant errors"
        }
      ],
      "scoring_guidelines": {
        "total_points": 10,
        "passing_score": 6
      },
      "ai_metadata": {
        "confidence_score": 0.90,
        "generated_at": "2026-06-03T12:00:00Z",
        "agent_version": "1.0"
      }
    }
  },
  "message": "Rubric generated successfully",
  "timestamp": "2026-06-03T12:00:00Z"
}
```

---

### API 5.6: Get Rubric by ID

**Endpoint Name**: Get Rubric by ID

**HTTP Method**: `GET`

**URL**: `/api/v1/assessment/rubric/{rubric_id}`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER, ADMINISTRATOR

**Path Parameters**:
- `rubric_id` (required): Rubric identifier

**Request Schema**: None

**Validation Rules**:
- `rubric_id`: Required, valid UUID format

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "rubric": {
      "id": "rub_1234567890",
      "assessment_id": "ass_1234567890",
      "status": "APPROVED",
      "rubric_type": "ANALYTIC",
      "performance_criteria": [
        {
          "id": "pc_1234567890",
          "category": "content",
          "description": "Accuracy and completeness of content",
          "weight": 0.4
        }
      ],
      "created_at": "2026-06-03T10:00:00Z",
      "updated_at": "2026-06-03T11:00:00Z",
      "approved_at": "2026-06-03T12:00:00Z",
      "approved_by": "usr_1234567890"
    }
  },
  "message": "Rubric retrieved successfully",
  "timestamp": "2026-06-03T12:00:00Z"
}
```

---

### API 5.7: Update Rubric

**Endpoint Name**: Update Rubric

**HTTP Method**: `PUT`

**URL**: `/api/v1/assessment/rubric/{rubric_id}`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Path Parameters**:
- `rubric_id` (required): Rubric identifier

**Request Schema**:

```json
{
  "performance_criteria": [
    {
      "id": "pc_1234567890",
      "category": "content",
      "description": "Updated description",
      "weight": 0.5
    }
  ]
}
```

**Validation Rules**:
- `rubric_id`: Required, valid UUID format
- `performance_criteria`: Required, array of performance criteria
- `performance_criteria[].id`: Required, valid UUID format
- `performance_criteria[].category`: Required, max 100 characters
- `performance_criteria[].description`: Required, max 500 characters
- `performance_criteria[].weight`: Required, min 0, max 1

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "rubric": {
      "id": "rub_1234567890",
      "status": "DRAFT",
      "updated_at": "2026-06-03T13:00:00Z"
    }
  },
  "message": "Rubric updated successfully",
  "timestamp": "2026-06-03T13:00:00Z"
}
```

---

### API 5.8: Approve Rubric

**Endpoint Name**: Approve Rubric

**HTTP Method**: `POST`

**URL**: `/api/v1/assessment/rubric/{rubric_id}/approve`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Path Parameters**:
- `rubric_id` (required): Rubric identifier

**Request Schema**: None

**Validation Rules**:
- `rubric_id`: Required, valid UUID format

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "rubric": {
      "id": "rub_1234567890",
      "status": "APPROVED",
      "approved_at": "2026-06-03T14:00:00Z",
      "approved_by": "usr_1234567890"
    }
  },
  "message": "Rubric approved successfully",
  "timestamp": "2026-06-03T14:00:00Z"
}
```

---

### API 5.9: Record Evidence

**Endpoint Name**: Record Evidence

**HTTP Method**: `POST`

**URL**: `/api/v1/assessment/evidence`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Request Schema**:

```json
{
  "student_id": "stu_1234567890",
  "assessment_id": "ass_1234567890",
  "evidence_type": "STUDENT_WORK",
  "evidence_data": {
    "file_url": "https://storage.nusa.id/evidence/assignment1.pdf",
    "file_name": "assignment1.pdf",
    "file_size": 1024000
  },
  "teacher_notes": "Student demonstrated good understanding of concepts"
}
```

**Validation Rules**:
- `student_id`: Required, valid UUID format
- `assessment_id`: Required, valid UUID format
- `evidence_type`: Required, valid type (STUDENT_WORK, ASSESSMENT_RESULT, OBSERVATION)
- `evidence_data`: Required, object with evidence details
- `teacher_notes`: Optional, max 1000 characters

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "evidence": {
      "id": "evi_1234567890",
      "student_id": "stu_1234567890",
      "assessment_id": "ass_1234567890",
      "evidence_type": "STUDENT_WORK",
      "status": "COLLECTED",
      "created_at": "2026-06-03T12:00:00Z",
      "created_by": "usr_1234567890"
    }
  },
  "message": "Evidence recorded successfully",
  "timestamp": "2026-06-03T12:00:00Z"
}
```

---

### API 5.10: Link Evidence to Rubric Criteria

**Endpoint Name**: Link Evidence to Rubric Criteria

**HTTP Method**: `POST`

**URL**: `/api/v1/assessment/evidence/{evidence_id}/link`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Path Parameters**:
- `evidence_id` (required): Evidence identifier

**Request Schema**:

```json
{
  "rubric_id": "rub_1234567890",
  "criteria_ids": ["pc_1234567890", "pc_1234567891"],
  "evaluation_notes": "Student performed well in content and organization"
}
```

**Validation Rules**:
- `evidence_id`: Required, valid UUID format
- `rubric_id`: Required, valid UUID format
- `criteria_ids`: Required, array of valid UUID formats
- `evaluation_notes`: Optional, max 1000 characters

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "evidence": {
      "id": "evi_1234567890",
      "status": "LINKED",
      "rubric_id": "rub_1234567890",
      "linked_criteria": ["pc_1234567890", "pc_1234567891"],
      "updated_at": "2026-06-03T13:00:00Z"
    }
  },
  "message": "Evidence linked to rubric criteria successfully",
  "timestamp": "2026-06-03T13:00:00Z"
}
```

---

### API 5.11: Evaluate Student Performance

**Endpoint Name**: Evaluate Student Performance

**HTTP Method**: `POST`

**URL**: `/api/v1/assessment/evaluation`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Request Schema**:

```json
{
  "student_id": "stu_1234567890",
  "rubric_id": "rub_1234567890",
  "evidence_id": "evi_1234567890",
  "performance_scores": [
    {
      "criteria_id": "pc_1234567890",
      "performance_level": 4,
      "score": 4,
      "notes": "Excellent understanding demonstrated"
    },
    {
      "criteria_id": "pc_1234567891",
      "performance_level": 3,
      "score": 3,
      "notes": "Good organization with minor improvements needed"
    }
  ]
}
```

**Validation Rules**:
- `student_id`: Required, valid UUID format
- `rubric_id`: Required, valid UUID format
- `evidence_id`: Required, valid UUID format
- `performance_scores`: Required, array of performance scores
- `performance_scores[].criteria_id`: Required, valid UUID format
- `performance_scores[].performance_level`: Required, min 1, max 5
- `performance_scores[].score`: Required, min 0, max 10
- `performance_scores[].notes`: Optional, max 500 characters

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "evaluation": {
      "id": "eval_1234567890",
      "student_id": "stu_1234567890",
      "rubric_id": "rub_1234567890",
      "evidence_id": "evi_1234567890",
      "total_score": 7,
      "max_score": 10,
      "performance_level": "PROFICIENT",
      "evaluated_at": "2026-06-03T14:00:00Z",
      "evaluated_by": "usr_1234567890"
    }
  },
  "message": "Student performance evaluated successfully",
  "timestamp": "2026-06-03T14:00:00Z"
}
```

---

# SECTION 6 — Reporting APIs

## Overview

Reporting APIs handle narrative report generation and management.

---

### API 6.1: Generate Narrative Report with AI

**Endpoint Name**: Generate Narrative Report with AI

**HTTP Method**: `POST`

**URL**: `/api/v1/reporting/narrative-report/generate`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Request Schema**:

```json
{
  "student_id": "stu_1234567890",
  "evidence_ids": ["evi_1234567890", "evi_1234567891"],
  "evaluation_ids": ["eval_1234567890", "eval_1234567891"],
  "report_period": {
    "type": "SEMESTER",
    "semester": 1,
    "academic_year": "2026"
  },
  "language": "INDONESIAN"
}
```

**Validation Rules**:
- `student_id`: Required, valid UUID format
- `evidence_ids`: Required, array of valid UUID formats
- `evaluation_ids`: Required, array of valid UUID formats
- `report_period.type`: Required, valid type (SEMESTER, TRIMESTER)
- `report_period.semester`: Required, min 1, max 2
- `report_period.academic_year`: Required, valid year format (YYYY)
- `language`: Required, valid language (INDONESIAN, ENGLISH)

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "narrative_report": {
      "id": "nr_1234567890",
      "student_id": "stu_1234567890",
      "status": "DRAFT",
      "report_period": {
        "type": "SEMESTER",
        "semester": 1,
        "academic_year": "2026"
      },
      "content": {
        "progress_summary": "John has demonstrated consistent progress in Mathematics this semester. He has shown strong understanding of algebraic concepts and applies them effectively in problem-solving.",
        "strengths": [
          "Strong algebraic thinking skills",
          "Excellent problem-solving abilities",
          "Consistent participation in class activities"
        ],
        "areas_for_improvement": [
          "Needs more practice with complex equations",
          "Should improve time management in assessments"
        ],
        "recommendations": [
          "Continue practicing algebraic problems regularly",
          "Focus on time management techniques during assessments",
          "Consider peer tutoring for complex topics"
        ]
      },
      "ai_metadata": {
        "confidence_score": 0.87,
        "generated_at": "2026-06-03T12:00:00Z",
        "agent_version": "1.0"
      }
    }
  },
  "message": "Narrative report generated successfully",
  "timestamp": "2026-06-03T12:00:00Z"
}
```

---

### API 6.2: Get Narrative Report by ID

**Endpoint Name**: Get Narrative Report by ID

**HTTP Method**: `GET`

**URL**: `/api/v1/reporting/narrative-report/{report_id}`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER, ADMINISTRATOR

**Path Parameters**:
- `report_id` (required): Narrative Report identifier

**Request Schema**: None

**Validation Rules**:
- `report_id`: Required, valid UUID format

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "narrative_report": {
      "id": "nr_1234567890",
      "student_id": "stu_1234567890",
      "status": "APPROVED",
      "report_period": {
        "type": "SEMESTER",
        "semester": 1,
        "academic_year": "2026"
      },
      "content": {
        "progress_summary": "John has demonstrated consistent progress...",
        "strengths": ["Strong algebraic thinking skills"],
        "areas_for_improvement": ["Needs more practice with complex equations"],
        "recommendations": ["Continue practicing algebraic problems regularly"]
      },
      "created_at": "2026-06-03T10:00:00Z",
      "updated_at": "2026-06-03T11:00:00Z",
      "approved_at": "2026-06-03T12:00:00Z",
      "approved_by": "usr_1234567890"
    }
  },
  "message": "Narrative report retrieved successfully",
  "timestamp": "2026-06-03T12:00:00Z"
}
```

---

### API 6.3: Update Narrative Report

**Endpoint Name**: Update Narrative Report

**HTTP Method**: `PUT`

**URL**: `/api/v1/reporting/narrative-report/{report_id}`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Path Parameters**:
- `report_id` (required): Narrative Report identifier

**Request Schema**:

```json
{
  "content": {
    "progress_summary": "Updated progress summary",
    "strengths": ["Updated strength"],
    "areas_for_improvement": ["Updated area for improvement"],
    "recommendations": ["Updated recommendation"]
  }
}
```

**Validation Rules**:
- `report_id`: Required, valid UUID format
- `content.progress_summary`: Required, max 2000 characters
- `content.strengths`: Required, array of strings, max 10 items
- `content.areas_for_improvement`: Required, array of strings, max 10 items
- `content.recommendations`: Required, array of strings, max 10 items

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "narrative_report": {
      "id": "nr_1234567890",
      "status": "DRAFT",
      "updated_at": "2026-06-03T13:00:00Z"
    }
  },
  "message": "Narrative report updated successfully",
  "timestamp": "2026-06-03T13:00:00Z"
}
```

---

### API 6.4: Approve Narrative Report

**Endpoint Name**: Approve Narrative Report

**HTTP Method**: `POST`

**URL**: `/api/v1/reporting/narrative-report/{report_id}/approve`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Path Parameters**:
- `report_id` (required): Narrative Report identifier

**Request Schema**: None

**Validation Rules**:
- `report_id`: Required, valid UUID format

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "narrative_report": {
      "id": "nr_1234567890",
      "status": "APPROVED",
      "approved_at": "2026-06-03T14:00:00Z",
      "approved_by": "usr_1234567890"
    }
  },
  "message": "Narrative report approved successfully",
  "timestamp": "2026-06-03T14:00:00Z"
}
```

---

# SECTION 7 — Administration APIs

## Overview

Administration APIs handle user management, role management, and system configuration.

---

### API 7.1: Create User Account

**Endpoint Name**: Create User Account

**HTTP Method**: `POST`

**URL**: `/api/v1/admin/users`

**Authentication Requirement**: Required

**Authorization Requirement**: ADMINISTRATOR

**Request Schema**:

```json
{
  "email": "teacher@school.id",
  "password": "securePassword123",
  "name": "John Doe",
  "role": "TEACHER"
}
```

**Validation Rules**:
- `email`: Required, valid email format, max 255 characters, unique
- `password`: Required, min 8 characters, max 128 characters
- `name`: Required, max 255 characters
- `role`: Required, valid role (TEACHER, ADMINISTRATOR)

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "user": {
      "id": "usr_1234567890",
      "email": "teacher@school.id",
      "name": "John Doe",
      "role": "TEACHER",
      "status": "ACTIVE",
      "created_at": "2026-06-03T12:00:00Z"
    }
  },
  "message": "User account created successfully",
  "timestamp": "2026-06-03T12:00:00Z"
}
```

---

### API 7.2: Update User Account

**Endpoint Name**: Update User Account

**HTTP Method**: `PUT`

**URL**: `/api/v1/admin/users/{user_id}`

**Authentication Requirement**: Required

**Authorization Requirement**: ADMINISTRATOR

**Path Parameters**:
- `user_id` (required): User identifier

**Request Schema**:

```json
{
  "email": "updated@school.id",
  "name": "Updated Name",
  "status": "ACTIVE"
}
```

**Validation Rules**:
- `user_id`: Required, valid UUID format
- `email`: Optional, valid email format, max 255 characters, unique
- `name`: Optional, max 255 characters
- `status`: Optional, valid status (ACTIVE, INACTIVE)

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "user": {
      "id": "usr_1234567890",
      "email": "updated@school.id",
      "name": "Updated Name",
      "status": "ACTIVE",
      "updated_at": "2026-06-03T13:00:00Z"
    }
  },
  "message": "User account updated successfully",
  "timestamp": "2026-06-03T13:00:00Z"
}
```

---

### API 7.3: Deactivate User Account

**Endpoint Name**: Deactivate User Account

**HTTP Method**: `POST`

**URL**: `/api/v1/admin/users/{user_id}/deactivate`

**Authentication Requirement**: Required

**Authorization Requirement**: ADMINISTRATOR

**Path Parameters**:
- `user_id` (required): User identifier

**Request Schema**: None

**Validation Rules**:
- `user_id`: Required, valid UUID format

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "user": {
      "id": "usr_1234567890",
      "status": "INACTIVE",
      "deactivated_at": "2026-06-03T14:00:00Z",
      "deactivated_by": "usr_1234567891"
    }
  },
  "message": "User account deactivated successfully",
  "timestamp": "2026-06-03T14:00:00Z"
}
```

---

### API 7.4: Assign User Role

**Endpoint Name**: Assign User Role

**HTTP Method**: `POST`

**URL**: `/api/v1/admin/users/{user_id}/roles`

**Authentication Requirement**: Required

**Authorization Requirement**: ADMINISTRATOR

**Path Parameters**:
- `user_id` (required): User identifier

**Request Schema**:

```json
{
  "role": "ADMINISTRATOR"
}
```

**Validation Rules**:
- `user_id`: Required, valid UUID format
- `role`: Required, valid role (TEACHER, ADMINISTRATOR)

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "user": {
      "id": "usr_1234567890",
      "role": "ADMINISTRATOR",
      "role_updated_at": "2026-06-03T14:00:00Z",
      "role_updated_by": "usr_1234567891"
    }
  },
  "message": "User role assigned successfully",
  "timestamp": "2026-06-03T14:00:00Z"
}
```

---

### API 7.5: List User Accounts

**Endpoint Name**: List User Accounts

**HTTP Method**: `GET`

**URL**: `/api/v1/admin/users`

**Authentication Requirement**: Required

**Authorization Requirement**: ADMINISTRATOR

**Query Parameters**:
- `role` (optional): Filter by role (TEACHER, ADMINISTRATOR)
- `status` (optional): Filter by status (ACTIVE, INACTIVE)
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 20, max: 100)

**Request Schema**: None (query parameters only)

**Validation Rules**:
- `role`: Optional, valid role (TEACHER, ADMINISTRATOR)
- `status`: Optional, valid status (ACTIVE, INACTIVE)
- `page`: Optional, min 1
- `limit`: Optional, min 1, max 100

**Response Schema** (Success):

```json
{
  "success": true,
  "data": {
    "users": [
      {
        "id": "usr_1234567890",
        "email": "teacher@school.id",
        "name": "John Doe",
        "role": "TEACHER",
        "status": "ACTIVE",
        "created_at": "2026-06-03T10:00:00Z"
      }
    ],
    "pagination": {
      "total": 50,
      "page": 1,
      "limit": 20,
      "total_pages": 3
    }
  },
  "message": "User accounts retrieved successfully",
  "timestamp": "2026-06-03T12:00:00Z"
}
```

---

# SECTION 8 — Error Codes

## Standard Error Codes

| Error Code | HTTP Status | Description |
|------------|-------------|-------------|
| INVALID_REQUEST | 400 | Invalid request parameters |
| UNAUTHORIZED | 401 | Authentication required or failed |
| FORBIDDEN | 403 | Authorization failed |
| NOT_FOUND | 404 | Resource not found |
| CONFLICT | 409 | Resource conflict |
| VALIDATION_FAILED | 422 | Validation failed |
| INTERNAL_ERROR | 500 | Internal server error |
| SERVICE_UNAVAILABLE | 503 | Service temporarily unavailable |

## Domain-Specific Error Codes

| Error Code | Description |
|------------|-------------|
| INVALID_CREDENTIALS | Invalid email or password |
| INVALID_REFRESH_TOKEN | Invalid or expired refresh token |
| CP_NOT_FOUND | Curriculum plan not found |
| GENERATION_FAILED | AI generation failed |
| TP_NOT_FOUND | Teaching plan not found |
| TP_UPDATE_FAILED | Failed to update teaching plan |
| TP_APPROVAL_FAILED | Failed to approve teaching plan |
| ATP_NOT_FOUND | Annual teaching plan not found |
| MODUL_AJAR_NOT_FOUND | Modul Ajar not found |
| ASSESSMENT_NOT_FOUND | Assessment not found |
| RUBRIC_NOT_FOUND | Rubric not found |
| EVIDENCE_NOT_FOUND | Evidence not found |
| REPORT_NOT_FOUND | Narrative report not found |
| USER_NOT_FOUND | User account not found |
| EMAIL_ALREADY_EXISTS | Email already registered |

---

# SECTION 9 — Rate Limiting

## Rate Limit Rules

- **Unauthenticated Requests**: 10 requests per minute
- **Authenticated Requests**: 100 requests per minute
- **AI Generation Endpoints**: 10 requests per hour per user

## Rate Limit Headers

Rate limit information is included in response headers:

```
X-RateLimit-Limit: 100
X-RateLimit-Remaining: 95
X-RateLimit-Reset: 1620000000
```

---

# SECTION 10 — OpenAPI Conversion

## OpenAPI Specification

This API contract is designed to be convertible to OpenAPI 3.0 specification. Key mappings:

- **Endpoint Name** → `operationId`
- **HTTP Method** → HTTP method
- **URL** → `path` with path parameters
- **Request Schema** → `requestBody.content.application/json.schema`
- **Response Schema** → `responses.200.content.application/json.schema`
- **Validation Rules** → `schema.validation`
- **Authentication Requirement** → `security`
- **Authorization Requirement** → `security` with role-based scopes

## OpenAPI Tags

APIs are organized by module tags:
- `authentication` - Authentication APIs
- `curriculum` - Curriculum APIs
- `learning-planning` - Learning Planning APIs
- `assessment` - Assessment APIs
- `reporting` - Reporting APIs
- `administration` - Administration APIs

---

# SECTION 11 — Domain Event Catalog

## Purpose

Define all operational domain events used by the application for workflow progression, module coordination, and RabbitMQ asynchronous processing. These events enable loose coupling between modules and support the curriculum-to-report workflow.

## Important Note

**NUSA MVP does NOT implement Event Sourcing.** Domain events are used for:
- Workflow progression notifications
- Module coordination via RabbitMQ
- Asynchronous processing
- Audit trail enhancement

Domain events are NOT used for:
- State reconstruction
- Event sourcing
- Event store persistence
- Temporal queries

---

## Event Overview

### Event Flow Architecture

```
Producer Module → Domain Event → RabbitMQ → Consumer Module
```

### Event Characteristics

- **Format**: JSON
- **Transport**: RabbitMQ
- **Routing**: Topic-based routing keys
- **Persistence**: Transient (no event store)
- **Delivery**: At-least-once
- **Idempotency**: Consumers must handle duplicate events

---

## Curriculum Module Events

### Event: TPGenerated

**Event Name**: `TPGenerated`

**Producer Module**: Curriculum Module

**Consumer Module**: 
- Learning Planning Module (for ATP generation)
- AI Orchestration Module (for logging)

**Trigger Condition**: Teaching Plan (TP) is successfully generated by AI agent

**Payload Summary**:
```json
{
  "event_id": "evt_1234567890",
  "event_type": "TPGenerated",
  "timestamp": "2026-06-03T12:00:00Z",
  "producer": "curriculum_module",
  "payload": {
    "tp_id": "tp_1234567890",
    "cp_id": "cp_1234567890",
    "user_id": "usr_1234567890",
    "status": "DRAFT",
    "ai_confidence_score": 0.92,
    "ai_agent_version": "1.0.0"
  }
}
```

---

### Event: TPUpdated

**Event Name**: `TPUpdated`

**Producer Module**: Curriculum Module

**Consumer Module**: 
- Learning Planning Module (for ATP invalidation check)
- AI Orchestration Module (for logging)

**Trigger Condition**: Teaching Plan (TP) is updated by teacher

**Payload Summary**:
```json
{
  "event_id": "evt_1234567891",
  "event_type": "TPUpdated",
  "timestamp": "2026-06-03T13:00:00Z",
  "producer": "curriculum_module",
  "payload": {
    "tp_id": "tp_1234567890",
    "user_id": "usr_1234567890",
    "status": "DRAFT",
    "updated_fields": ["learning_objectives", "time_allocation"]
  }
}
```

---

### Event: TPApproved

**Event Name**: `TPApproved`

**Producer Module**: Curriculum Module

**Consumer Module**: 
- Learning Planning Module (for ATP generation trigger)
- AI Orchestration Module (for logging)
- Administration Module (for analytics)

**Trigger Condition**: Teaching Plan (TP) is approved by teacher

**Payload Summary**:
```json
{
  "event_id": "evt_1234567892",
  "event_type": "TPApproved",
  "timestamp": "2026-06-03T14:00:00Z",
  "producer": "curriculum_module",
  "payload": {
    "tp_id": "tp_1234567890",
    "cp_id": "cp_1234567890",
    "user_id": "usr_1234567890",
    "approved_by": "usr_1234567890",
    "approved_at": "2026-06-03T14:00:00Z"
  }
}
```

---

## Learning Planning Module Events

### Event: ATPGenerated

**Event Name**: `ATPGenerated`

**Producer Module**: Learning Planning Module

**Consumer Module**: 
- Learning Planning Module (for Modul Ajar generation)
- AI Orchestration Module (for logging)

**Trigger Condition**: Annual Teaching Plan (ATP) is successfully generated by AI agent

**Payload Summary**:
```json
{
  "event_id": "evt_1234567893",
  "event_type": "ATPGenerated",
  "timestamp": "2026-06-03T15:00:00Z",
  "producer": "learning_planning_module",
  "payload": {
    "atp_id": "atp_1234567890",
    "tp_id": "tp_1234567890",
    "user_id": "usr_1234567890",
    "status": "DRAFT",
    "ai_confidence_score": 0.89,
    "ai_agent_version": "1.0.0"
  }
}
```

---

### Event: ATPApproved

**Event Name**: `ATPApproved`

**Producer Module**: Learning Planning Module

**Consumer Module**: 
- Learning Planning Module (for Modul Ajar generation trigger)
- AI Orchestration Module (for logging)
- Administration Module (for analytics)

**Trigger Condition**: Annual Teaching Plan (ATP) is approved by teacher

**Payload Summary**:
```json
{
  "event_id": "evt_1234567894",
  "event_type": "ATPApproved",
  "timestamp": "2026-06-03T16:00:00Z",
  "producer": "learning_planning_module",
  "payload": {
    "atp_id": "atp_1234567890",
    "tp_id": "tp_1234567890",
    "user_id": "usr_1234567890",
    "approved_by": "usr_1234567890",
    "approved_at": "2026-06-03T16:00:00Z"
  }
}
```

---

### Event: ModulAjarGenerated

**Event Name**: `ModulAjarGenerated`

**Producer Module**: Learning Planning Module

**Consumer Module**: 
- Assessment Module (for assessment generation)
- AI Orchestration Module (for logging)

**Trigger Condition**: Modul Ajar is successfully generated by AI agent

**Payload Summary**:
```json
{
  "event_id": "evt_1234567895",
  "event_type": "ModulAjarGenerated",
  "timestamp": "2026-06-03T17:00:00Z",
  "producer": "learning_planning_module",
  "payload": {
    "modul_ajar_id": "ma_1234567890",
    "atp_id": "atp_1234567890",
    "week": 1,
    "user_id": "usr_1234567890",
    "status": "DRAFT",
    "ai_confidence_score": 0.91,
    "ai_agent_version": "1.0.0"
  }
}
```

---

### Event: ModulAjarApproved

**Event Name**: `ModulAjarApproved`

**Producer Module**: Learning Planning Module

**Consumer Module**: 
- Assessment Module (for assessment generation trigger)
- AI Orchestration Module (for logging)
- Administration Module (for analytics)

**Trigger Condition**: Modul Ajar is approved by teacher

**Payload Summary**:
```json
{
  "event_id": "evt_1234567896",
  "event_type": "ModulAjarApproved",
  "timestamp": "2026-06-03T18:00:00Z",
  "producer": "learning_planning_module",
  "payload": {
    "modul_ajar_id": "ma_1234567890",
    "atp_id": "atp_1234567890",
    "user_id": "usr_1234567890",
    "approved_by": "usr_1234567890",
    "approved_at": "2026-06-03T18:00:00Z"
  }
}
```

---

## Assessment Module Events

### Event: AssessmentGenerated

**Event Name**: `AssessmentGenerated`

**Producer Module**: Assessment Module

**Consumer Module**: 
- Assessment Module (for rubric generation)
- AI Orchestration Module (for logging)

**Trigger Condition**: Assessment is successfully generated by AI agent

**Payload Summary**:
```json
{
  "event_id": "evt_1234567897",
  "event_type": "AssessmentGenerated",
  "timestamp": "2026-06-03T19:00:00Z",
  "producer": "assessment_module",
  "payload": {
    "assessment_id": "ass_1234567890",
    "modul_ajar_id": "ma_1234567890",
    "user_id": "usr_1234567890",
    "assessment_type": "FORMATIVE",
    "status": "DRAFT",
    "ai_confidence_score": 0.88,
    "ai_agent_version": "1.0.0"
  }
}
```

---

### Event: RubricGenerated

**Event Name**: `RubricGenerated`

**Producer Module**: Assessment Module

**Consumer Module**: 
- Assessment Module (for evidence collection)
- AI Orchestration Module (for logging)

**Trigger Condition**: Rubric is successfully generated by AI agent

**Payload Summary**:
```json
{
  "event_id": "evt_1234567898",
  "event_type": "RubricGenerated",
  "timestamp": "2026-06-03T20:00:00Z",
  "producer": "assessment_module",
  "payload": {
    "rubric_id": "rub_1234567890",
    "assessment_id": "ass_1234567890",
    "user_id": "usr_1234567890",
    "rubric_type": "ANALYTIC",
    "status": "DRAFT",
    "ai_confidence_score": 0.90,
    "ai_agent_version": "1.0.0"
  }
}
```

---

### Event: EvidenceRecorded

**Event Name**: `EvidenceRecorded`

**Producer Module**: Assessment Module

**Consumer Module**: 
- Reporting Module (for narrative report generation)
- AI Orchestration Module (for logging)

**Trigger Condition**: Student evidence is recorded by teacher

**Payload Summary**:
```json
{
  "event_id": "evt_1234567899",
  "event_type": "EvidenceRecorded",
  "timestamp": "2026-06-03T21:00:00Z",
  "producer": "assessment_module",
  "payload": {
    "evidence_id": "evi_1234567890",
    "student_id": "stu_1234567890",
    "assessment_id": "ass_1234567890",
    "user_id": "usr_1234567890",
    "evidence_type": "STUDENT_WORK",
    "status": "COLLECTED"
  }
}
```

---

## Reporting Module Events

### Event: NarrativeReportGenerated

**Event Name**: `NarrativeReportGenerated`

**Producer Module**: Reporting Module

**Consumer Module**: 
- AI Orchestration Module (for logging)
- Administration Module (for analytics)

**Trigger Condition**: Narrative report is successfully generated by AI agent

**Payload Summary**:
```json
{
  "event_id": "evt_1234567900",
  "event_type": "NarrativeReportGenerated",
  "timestamp": "2026-06-03T22:00:00Z",
  "producer": "reporting_module",
  "payload": {
    "report_id": "nr_1234567890",
    "student_id": "stu_1234567890",
    "user_id": "usr_1234567890",
    "status": "DRAFT",
    "language": "INDONESIAN",
    "ai_confidence_score": 0.87,
    "ai_agent_version": "1.0.0"
  }
}
```

---

### Event: NarrativeReportApproved

**Event Name**: `NarrativeReportApproved`

**Producer Module**: Reporting Module

**Consumer Module**: 
- Administration Module (for analytics)
- AI Orchestration Module (for logging)

**Trigger Condition**: Narrative report is approved by teacher

**Payload Summary**:
```json
{
  "event_id": "evt_1234567901",
  "event_type": "NarrativeReportApproved",
  "timestamp": "2026-06-03T23:00:00Z",
  "producer": "reporting_module",
  "payload": {
    "report_id": "nr_1234567890",
    "student_id": "stu_1234567890",
    "user_id": "usr_1234567890",
    "approved_by": "usr_1234567890",
    "approved_at": "2026-06-03T23:00:00Z"
  }
}
```

---

## Administration Module Events

### Event: UserCreated

**Event Name**: `UserCreated`

**Producer Module**: Administration Module

**Consumer Module**: 
- Authentication Module (for session initialization)
- AI Orchestration Module (for logging)

**Trigger Condition**: User account is created by administrator

**Payload Summary**:
```json
{
  "event_id": "evt_1234567902",
  "event_type": "UserCreated",
  "timestamp": "2026-06-03T00:00:00Z",
  "producer": "administration_module",
  "payload": {
    "user_id": "usr_1234567890",
    "email": "teacher@school.id",
    "name": "John Doe",
    "role": "TEACHER",
    "created_by": "usr_1234567891"
  }
}
```

---

### Event: UserRoleChanged

**Event Name**: `UserRoleChanged`

**Producer Module**: Administration Module

**Consumer Module**: 
- Authentication Module (for permission cache update)
- AI Orchestration Module (for logging)

**Trigger Condition**: User role is changed by administrator

**Payload Summary**:
```json
{
  "event_id": "evt_1234567903",
  "event_type": "UserRoleChanged",
  "timestamp": "2026-06-03T01:00:00Z",
  "producer": "administration_module",
  "payload": {
    "user_id": "usr_1234567890",
    "old_role": "TEACHER",
    "new_role": "ADMINISTRATOR",
    "changed_by": "usr_1234567891"
  }
}
```

---

## AI Orchestration Module Events

### Event: AIGenerationRequested

**Event Name**: `AIGenerationRequested`

**Producer Module**: AI Orchestration Module

**Consumer Module**: 
- AI Orchestration Module (for pipeline execution)
- Administration Module (for analytics)

**Trigger Condition**: AI generation request is initiated by user action

**Payload Summary**:
```json
{
  "event_id": "evt_1234567904",
  "event_type": "AIGenerationRequested",
  "timestamp": "2026-06-03T02:00:00Z",
  "producer": "ai_orchestration_module",
  "payload": {
    "generation_id": "gen_1234567890",
    "agent_type": "TP_GENERATOR",
    "user_id": "usr_1234567890",
    "input_hash": "sha256_hash",
    "priority": "NORMAL"
  }
}
```

---

### Event: AIGenerationCompleted

**Event Name**: `AIGenerationCompleted`

**Producer Module**: AI Orchestration Module

**Consumer Module**: 
- Curriculum/Learning Planning/Assessment/Reporting Module (for artifact creation)
- Administration Module (for analytics)
- AI Orchestration Module (for logging)

**Trigger Condition**: AI generation is successfully completed

**Payload Summary**:
```json
{
  "event_id": "evt_1234567905",
  "event_type": "AIGenerationCompleted",
  "timestamp": "2026-06-03T03:00:00Z",
  "producer": "ai_orchestration_module",
  "payload": {
    "generation_id": "gen_1234567890",
    "agent_type": "TP_GENERATOR",
    "user_id": "usr_1234567890",
    "artifact_id": "tp_1234567890",
    "confidence_score": 0.92,
    "duration_ms": 1500
  }
}
```

---

### Event: AIGenerationFailed

**Event Name**: `AIGenerationFailed`

**Producer Module**: AI Orchestration Module

**Consumer Module**: 
- AI Orchestration Module (for retry logic)
- Administration Module (for analytics)
- AI Orchestration Module (for logging)

**Trigger Condition**: AI generation fails (LLM error, timeout, validation failure)

**Payload Summary**:
```json
{
  "event_id": "evt_1234567906",
  "event_type": "AIGenerationFailed",
  "timestamp": "2026-06-03T04:00:00Z",
  "producer": "ai_orchestration_module",
  "payload": {
    "generation_id": "gen_1234567890",
    "agent_type": "TP_GENERATOR",
    "user_id": "usr_1234567890",
    "error_code": "LLM_TIMEOUT",
    "error_message": "LLM provider timeout",
    "retry_count": 2
  }
}
```

---

## Event Ownership

### Event Ownership by Module

| Module | Events Owned | Total Events |
|--------|--------------|--------------|
| Curriculum Module | TPGenerated, TPUpdated, TPApproved | 3 |
| Learning Planning Module | ATPGenerated, ATPApproved, ModulAjarGenerated, ModulAjarApproved | 4 |
| Assessment Module | AssessmentGenerated, RubricGenerated, EvidenceRecorded | 3 |
| Reporting Module | NarrativeReportGenerated, NarrativeReportApproved | 2 |
| Administration Module | UserCreated, UserRoleChanged | 2 |
| AI Orchestration Module | AIGenerationRequested, AIGenerationCompleted, AIGenerationFailed | 3 |

**Total Events**: 17

---

## Event Flow Diagrams

### Curriculum-to-Report Workflow Event Flow

```
User Action
  ↓
AIGenerationRequested (AI Orchestration Module)
  ↓
AIGenerationCompleted (AI Orchestration Module)
  ↓
TPGenerated (Curriculum Module)
  ↓
TPApproved (Curriculum Module)
  ↓
ATPGenerated (Learning Planning Module)
  ↓
ATPApproved (Learning Planning Module)
  ↓
ModulAjarGenerated (Learning Planning Module)
  ↓
ModulAjarApproved (Learning Planning Module)
  ↓
AssessmentGenerated (Assessment Module)
  ↓
RubricGenerated (Assessment Module)
  ↓
EvidenceRecorded (Assessment Module)
  ↓
NarrativeReportGenerated (Reporting Module)
  ↓
NarrativeReportApproved (Reporting Module)
```

### User Management Event Flow

```
Administrator Action
  ↓
UserCreated (Administration Module)
  ↓
UserCreated (Authentication Module - consumer)
  ↓
Administrator Action
  ↓
UserRoleChanged (Administration Module)
  ↓
UserRoleChanged (Authentication Module - consumer)
```

### AI Generation Event Flow

```
User Action
  ↓
AIGenerationRequested (AI Orchestration Module)
  ↓
[Success]
  ↓
AIGenerationCompleted (AI Orchestration Module)
  ↓
Artifact-Specific Event (e.g., TPGenerated)
  ↓
[Failure]
  ↓
AIGenerationFailed (AI Orchestration Module)
  ↓
Retry or Manual Intervention
```

---

## RabbitMQ Configuration

### Exchange Configuration

**Exchange Name**: `nusa.events`
**Exchange Type**: `topic`

### Queue Configuration

| Queue | Routing Key Pattern | Purpose |
|-------|-------------------|---------|
| `curriculum_queue` | `curriculum.*` | Curriculum module events |
| `learning_planning_queue` | `learning_planning.*` | Learning planning module events |
| `assessment_queue` | `assessment.*` | Assessment module events |
| `reporting_queue` | `reporting.*` | Reporting module events |
| `administration_queue` | `administration.*` | Administration module events |
| `ai_orchestration_queue` | `ai.*` | AI orchestration module events |
| `audit_queue` | `*.audit` | Audit trail for all events |

### Routing Keys

| Event | Routing Key |
|-------|-------------|
| TPGenerated | `curriculum.tp.generated` |
| TPUpdated | `curriculum.tp.updated` |
| TPApproved | `curriculum.tp.approved` |
| ATPGenerated | `learning_planning.atp.generated` |
| ATPApproved | `learning_planning.atp.approved` |
| ModulAjarGenerated | `learning_planning.modul_ajar.generated` |
| ModulAjarApproved | `learning_planning.modul_ajar.approved` |
| AssessmentGenerated | `assessment.assessment.generated` |
| RubricGenerated | `assessment.rubric.generated` |
| EvidenceRecorded | `assessment.evidence.recorded` |
| NarrativeReportGenerated | `reporting.narrative_report.generated` |
| NarrativeReportApproved | `reporting.narrative_report.approved` |
| UserCreated | `administration.user.created` |
| UserRoleChanged | `administration.user.role_changed` |
| AIGenerationRequested | `ai.generation.requested` |
| AIGenerationCompleted | `ai.generation.completed` |
| AIGenerationFailed | `ai.generation.failed` |

---

## Event Consumer Guidelines

### Idempotency

All event consumers must be idempotent:
- Use event_id for deduplication
- Handle duplicate events gracefully
- Implement retry logic with exponential backoff

### Error Handling

Event consumers must:
- Log all processing errors
- Implement dead letter queue for failed events
- Alert on critical failures
- Not block event processing on transient errors

### Event Ordering

Event ordering is guaranteed within a single queue but not across queues:
- Consumers should not assume cross-queue ordering
- Use timestamps for temporal ordering if needed
- Implement compensation for out-of-order events

---

## Event Versioning

### Event Schema Versioning

- Event schemas include version field
- Consumers must handle multiple versions
- Deprecated events are supported for 6 months
- Breaking changes require new event type

### Event Version Format

```json
{
  "event_id": "evt_1234567890",
  "event_type": "TPGenerated",
  "event_version": "1.0",
  "timestamp": "2026-06-03T12:00:00Z",
  "producer": "curriculum_module",
  "payload": { }
}
```

---

## Conclusion

This Domain Event Catalog defines 17 operational domain events for NUSA MVP Wave 1, enabling workflow progression, module coordination, and RabbitMQ asynchronous processing without implementing Event Sourcing.

---

# SECTION 12 — API Error Catalog

## Purpose

Standardize API errors across all modules to ensure consistent error handling, clear error messages, and actionable resolution guidance for API consumers.

## Error Response Format

All error responses follow the standard format:

```json
{
  "success": false,
  "error": {
    "code": "ERROR_CODE",
    "message": "Human-readable error message",
    "details": {
      "field": "specific_field",
      "reason": "detailed_reason"
    }
  },
  "timestamp": "2026-06-03T12:00:00Z"
}
```

---

## Authentication Errors

### AUTH-001: Invalid Credentials

**Error Code**: `AUTH-001`

**HTTP Status**: `401 Unauthorized`

**Error Message**: Invalid email or password

**Description**: The provided email or password does not match any user account in the system.

**Resolution Guidance**:
- Verify email address is correct
- Check for typos in password
- Ensure account is active and not locked
- Contact administrator if password reset is needed

---

### AUTH-002: Token Expired

**Error Code**: `AUTH-002`

**HTTP Status**: `401 Unauthorized`

**Error Message**: Access token has expired

**Description**: The JWT access token has exceeded its 24-hour validity period.

**Resolution Guidance**:
- Use the refresh token to obtain a new access token via POST /api/v1/auth/refresh
- If refresh token is also expired, re-authenticate with login credentials
- Ensure client application handles token refresh automatically

---

### AUTH-003: Invalid Token

**Error Code**: `AUTH-003`

**HTTP Status**: `401 Unauthorized`

**Error Message**: Invalid or malformed token

**Description**: The provided JWT token is invalid, malformed, or has been revoked.

**Resolution Guidance**:
- Verify token format is correct (Bearer <token>)
- Check token has not been tampered with
- Re-authenticate to obtain a new valid token
- Check if user session has been terminated by administrator

---

## Authorization Errors

### RBAC-001: Insufficient Permissions

**Error Code**: `RBAC-001`

**HTTP Status**: `403 Forbidden`

**Error Message**: Insufficient permissions to access this resource

**Description**: The authenticated user does not have the required role or permission to access the requested resource.

**Resolution Guidance**:
- Verify user has the correct role (Administrator or Teacher)
- Check role permissions for the requested resource
- Contact administrator to request additional permissions if needed
- Ensure user is accessing resources within their authorized scope

---

### RBAC-002: Resource Access Denied

**Error Code**: `RBAC-002`

**HTTP Status**: `403 Forbidden`

**Error Message**: Access to this resource is denied

**Description**: The user does not have ownership or access rights to the specific resource being requested.

**Resolution Guidance**:
- Verify user owns the resource or has been granted access
- Check if resource belongs to the user's school or organization
- Contact resource owner or administrator for access
- Ensure correct resource ID is being used

---

## Validation Errors

### VAL-001: Invalid Request Parameters

**Error Code**: `VAL-001`

**HTTP Status**: `400 Bad Request`

**Error Message**: Invalid request parameters

**Description**: One or more request parameters fail validation rules (format, type, length, or value constraints).

**Resolution Guidance**:
- Review request body against API schema
- Check for missing required fields
- Verify data types match expected format
- Ensure string fields meet length requirements
- Validate enum values against allowed values

---

### VAL-002: Data Constraint Violation

**Error Code**: `VAL-002`

**HTTP Status**: `422 Unprocessable Entity`

**Error Message**: Data constraint violation

**Description**: The submitted data violates business rules or database constraints (unique constraints, foreign key constraints, etc.).

**Resolution Guidance**:
- Check for duplicate records (email, username, etc.)
- Verify referenced entities exist
- Ensure data relationships are valid
- Review business rule constraints
- Check for circular dependencies

---

## Curriculum Errors

### CUR-001: Curriculum Plan Not Found

**Error Code**: `CUR-001`

**HTTP Status**: `404 Not Found`

**Error Message**: Curriculum Plan not found

**Description**: The requested Curriculum Plan (CP) does not exist or has been deactivated.

**Resolution Guidance**:
- Verify CP ID is correct
- Check if CP is active and not archived
- Ensure CP exists for the specified subject, grade, and academic year
- Contact administrator to import required curriculum data

---

### CUR-002: Teaching Plan Not Found

**Error Code**: `CUR-002`

**HTTP Status**: `404 Not Found`

**Error Message**: Teaching Plan not found

**Description**: The requested Teaching Plan (TP) does not exist or has been deleted.

**Resolution Guidance**:
- Verify TP ID is correct
- Check if TP belongs to the authenticated user
- Ensure TP has not been deleted or archived
- Generate a new TP if needed from available CP

---

## Learning Planning Errors

### PLN-001: Annual Teaching Plan Not Found

**Error Code**: `PLN-001`

**HTTP Status**: `404 Not Found`

**Error Message**: Annual Teaching Plan not found

**Description**: The requested Annual Teaching Plan (ATP) does not exist or has been deleted.

**Resolution Guidance**:
- Verify ATP ID is correct
- Check if ATP belongs to the authenticated user
- Ensure ATP has not been deleted or archived
- Generate a new ATP if needed from approved TP

---

### PLN-002: Modul Ajar Not Found

**Error Code**: `PLN-002`

**HTTP Status**: `404 Not Found`

**Error Message**: Modul Ajar not found

**Description**: The requested Modul Ajar (lesson plan) does not exist or has been deleted.

**Resolution Guidance**:
- Verify Modul Ajar ID is correct
- Check if Modul Ajar belongs to the authenticated user
- Ensure Modul Ajar has not been deleted or archived
- Generate a new Modul Ajar if needed from approved ATP

---

## Assessment Errors

### ASM-001: Assessment Not Found

**Error Code**: `ASM-001`

**HTTP Status**: `404 Not Found`

**Error Message**: Assessment not found

**Description**: The requested assessment does not exist or has been deleted.

**Resolution Guidance**:
- Verify assessment ID is correct
- Check if assessment belongs to the authenticated user
- Ensure assessment has not been deleted or archived
- Generate a new assessment if needed from approved Modul Ajar

---

### ASM-002: Rubric Not Found

**Error Code**: `ASM-002`

**HTTP Status**: `404 Not Found`

**Error Message**: Rubric not found

**Description**: The requested rubric does not exist or has been deleted.

**Resolution Guidance**:
- Verify rubric ID is correct
- Check if rubric belongs to the authenticated user
- Ensure rubric has not been deleted or archived
- Generate a new rubric if needed from approved assessment

---

## Reporting Errors

### REP-001: Narrative Report Not Found

**Error Code**: `REP-001`

**HTTP Status**: `404 Not Found`

**Error Message**: Narrative report not found

**Description**: The requested narrative report does not exist or has been deleted.

**Resolution Guidance**:
- Verify report ID is correct
- Check if report belongs to the authenticated user
- Ensure report has not been deleted or archived
- Generate a new narrative report if needed from evidence and evaluations

---

### REP-002: Insufficient Evidence for Report

**Error Code**: `REP-002`

**HTTP Status**: `400 Bad Request`

**Error Message**: Insufficient evidence to generate narrative report

**Description**: Not enough student evidence and evaluations have been collected to generate a meaningful narrative report.

**Resolution Guidance**:
- Collect additional student evidence
- Complete student evaluations using approved rubrics
- Ensure minimum evidence requirements are met
- Verify evidence is linked to appropriate rubric criteria

---

## AI Errors

### AI-001: Prompt Validation Failed

**Error Code**: `AI-001`

**HTTP Status**: `400 Bad Request`

**Error Message**: AI prompt validation failed

**Description**: The generated prompt fails validation rules (missing required fields, invalid format, or exceeds size limits).

**Resolution Guidance**:
- Check prompt template configuration
- Verify input data is complete and valid
- Ensure prompt does not exceed token limits
- Review prompt validation logs for specific failure reason
- Contact administrator if issue persists

---

### AI-002: Provider Timeout

**Error Code**: `AI-002`

**HTTP Status**: `504 Gateway Timeout`

**Error Message**: AI provider timeout

**Description**: The external LLM provider did not respond within the configured timeout period.

**Resolution Guidance**:
- Retry the generation request
- Check if LLM provider service is operational
- Verify network connectivity to LLM provider
- Consider increasing timeout configuration if needed
- Contact administrator if timeout occurs repeatedly

---

### AI-003: Invalid Response

**Error Code**: `AI-003`

**HTTP Status**: `502 Bad Gateway`

**Error Message**: Invalid AI provider response

**Description**: The LLM provider returned an invalid or malformed response that could not be processed.

**Resolution Guidance**:
- Retry the generation request
- Check LLM provider API status
- Verify provider API configuration is correct
- Review provider response logs for specific error details
- Contact administrator if issue persists

---

### AI-004: Generation Failed

**Error Code**: `AI-004`

**HTTP Status**: `500 Internal Server Error`

**Error Message**: AI generation failed

**Description**: The AI generation process failed due to an internal error or validation failure.

**Resolution Guidance**:
- Review generation logs for specific error details
- Check input data quality and completeness
- Verify AI agent configuration is correct
- Retry generation with modified inputs if possible
- Contact administrator if issue persists

---

## Administration Errors

### ADM-001: User Not Found

**Error Code**: `ADM-001`

**HTTP Status**: `404 Not Found`

**Error Message**: User account not found

**Description**: The requested user account does not exist or has been deactivated.

**Resolution Guidance**:
- Verify user ID is correct
- Check if user account is active
- Ensure user has not been deleted
- Create new user account if needed

---

### ADM-002: Email Already Exists

**Error Code**: `ADM-002`

**HTTP Status**: `409 Conflict`

**Error Message**: Email address already registered

**Description**: The provided email address is already associated with an existing user account.

**Resolution Guidance**:
- Verify email address is correct
- Check if user already exists with different role
- Use different email address for new account
- Update existing user account instead of creating duplicate

---

## Error Code Summary

### Error Code Distribution

| Category | Error Codes | Total |
|----------|-------------|-------|
| Authentication | AUTH-001, AUTH-002, AUTH-003 | 3 |
| Authorization | RBAC-001, RBAC-002 | 2 |
| Validation | VAL-001, VAL-002 | 2 |
| Curriculum | CUR-001, CUR-002 | 2 |
| Learning Planning | PLN-001, PLN-002 | 2 |
| Assessment | ASM-001, ASM-002 | 2 |
| Reporting | REP-001, REP-002 | 2 |
| AI | AI-001, AI-002, AI-003, AI-004 | 4 |
| Administration | ADM-001, ADM-002 | 2 |

**Total Error Codes**: 21

### HTTP Status Distribution

| HTTP Status | Count | Error Codes |
|-------------|-------|-------------|
| 400 Bad Request | 3 | VAL-001, AI-001, REP-002 |
| 401 Unauthorized | 3 | AUTH-001, AUTH-002, AUTH-003 |
| 403 Forbidden | 2 | RBAC-001, RBAC-002 |
| 404 Not Found | 6 | CUR-001, CUR-002, PLN-001, PLN-002, ASM-001, ASM-002, REP-001 |
| 409 Conflict | 1 | ADM-002 |
| 422 Unprocessable Entity | 1 | VAL-002 |
| 500 Internal Server Error | 1 | AI-004 |
| 502 Bad Gateway | 1 | AI-003 |
| 504 Gateway Timeout | 1 | AI-002 |

---

## Error Handling Guidelines

### Client-Side Error Handling

All API clients must:
- Parse error response format consistently
- Display user-friendly error messages to end users
- Implement retry logic for transient errors (5xx status codes)
- Log error codes and timestamps for debugging
- Handle authentication token expiration gracefully

### Server-Side Error Handling

All API endpoints must:
- Return error codes from this catalog
- Include descriptive error messages
- Log detailed error information server-side
- Sanitize error messages to avoid exposing sensitive information
- Use appropriate HTTP status codes

### Error Monitoring

The system must:
- Track error rates by error code
- Alert on critical errors (5xx status codes)
- Monitor error trends over time
- Generate error reports for administrators
- Maintain error logs for troubleshooting

---

## Conclusion

This API Error Catalog standardizes 21 error codes across all NUSA MVP Wave 1 modules, ensuring consistent error handling, clear error messages, and actionable resolution guidance for API consumers.

---

# SECTION 13 — Conclusion

## API Contract Summary

This API Contract (13) provides the complete API catalog for NUSA MVP Wave 1:

### API Count
- **Total Endpoints**: 30
- **Authentication APIs**: 3
- **Curriculum APIs**: 6
- **Learning Planning APIs**: 8
- **Assessment APIs**: 11
- **Reporting APIs**: 4
- **Administration APIs**: 5

### API Characteristics
- RESTful design principles
- JWT authentication
- Role-based authorization
- Standardized response format
- Comprehensive validation rules
- Detailed error codes
- Rate limiting support
- OpenAPI-ready

### Implementation Readiness
This API contract is:
- ✅ Complete with all MVP endpoints
- ✅ Aligned with frozen architecture decisions
- ✅ Ready for backend implementation
- ✅ Ready for frontend integration
- ✅ Convertible to OpenAPI specification

The API contract is officially approved for NUSA MVP Wave 1 implementation.

---

**Document Status**: FOUNDATION DOCUMENT - LOCKED

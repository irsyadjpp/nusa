# 19_OPENAPI_PREPARATION.md

## Foundation Document for NUSA Education Platform

**Version**: 1.0
**Date**: June 2026
**Status**: FOUNDATION DOCUMENT - LOCKED
**Alignment**: Validated against Foundation Architecture (00A, 00B, 00C, 01, 02, 03, 04, 05, 06, 07, 08, 09, 10, 11, 12, 13, 14, 15, 16, 17, 18)

**Purpose**: Prepare OpenAPI specification generation based on the API Contract (13_API_CONTRACT.md). This document provides implementation-ready endpoint definitions for Swagger/OpenAPI generation, serving as the single source of truth for API documentation.

**Source of Truth**: 13_API_CONTRACT.md

**Constraint**: No new endpoints may be introduced. All endpoints must match the API Contract exactly.

---

# OpenAPI Specification Overview

## OpenAPI Version

**Version**: 3.0.3

## API Information

```yaml
openapi: 3.0.3
info:
  title: NUSA Education Platform API
  version: 1.0.0
  description: API for NUSA Education Platform MVP Wave 1
  contact:
    name: NUSA Platform Team
    email: api@nusa.id
  license:
    name: Proprietary
servers:
  - url: http://localhost:8080/api/v1
    description: Development server
  - url: https://api.nusa.id/v1
    description: Production server
```

## Security Schemes

```yaml
components:
  securitySchemes:
    bearerAuth:
      type: http
      scheme: bearer
      bearerFormat: JWT
      description: JWT Bearer Token authentication
```

---

# Authentication APIs

## API 2.1: User Login

**Endpoint**: `/api/v1/auth/login`

**HTTP Method**: `POST`

**Operation ID**: `userLogin`

**Tags**: `authentication`

**Authentication Requirement**: None

**Authorization Requirement**: None

**Request Schema**:

```yaml
requestBody:
  required: true
  content:
    application/json:
      schema:
        type: object
        required:
          - email
          - password
        properties:
          email:
            type: string
            format: email
            maxLength: 255
            description: User email address
            example: "teacher@school.id"
          password:
            type: string
            minLength: 8
            maxLength: 128
            description: User password
            example: "securePassword123"
```

**Response Schema** (Success - 200):

```yaml
responses:
  '200':
    description: Login successful
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: true
            data:
              type: object
              properties:
                user:
                  type: object
                  properties:
                    id:
                      type: string
                      format: uuid
                      example: "usr_1234567890"
                    email:
                      type: string
                      format: email
                      example: "teacher@school.id"
                    name:
                      type: string
                      example: "John Doe"
                    role:
                      type: string
                      enum: [TEACHER, ADMINISTRATOR]
                      example: "TEACHER"
                tokens:
                  type: object
                  properties:
                    access_token:
                      type: string
                      example: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
                    refresh_token:
                      type: string
                      example: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
                    token_type:
                      type: string
                      example: "Bearer"
                    expires_in:
                      type: integer
                      example: 86400
            message:
              type: string
              example: "Login successful"
            timestamp:
              type: string
              format: date-time
              example: "2026-06-03T12:00:00Z"
```

**Response Schema** (Error - 401):

```yaml
  '401':
    description: Invalid credentials
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: false
            error:
              type: object
              properties:
                code:
                  type: string
                  example: "INVALID_CREDENTIALS"
                message:
                  type: string
                  example: "Invalid email or password"
                details:
                  type: object
            timestamp:
              type: string
              format: date-time
```

**Error Codes**:
- `INVALID_CREDENTIALS` (401): Invalid email or password

**Related Feature**: User Authentication

**Related Database Tables**:
- `users`

---

## API 2.2: Refresh Token

**Endpoint**: `/api/v1/auth/refresh`

**HTTP Method**: `POST`

**Operation ID**: `refreshToken`

**Tags**: `authentication`

**Authentication Requirement**: None

**Authorization Requirement**: None

**Request Schema**:

```yaml
requestBody:
  required: true
  content:
    application/json:
      schema:
        type: object
        required:
          - refresh_token
        properties:
          refresh_token:
            type: string
            description: JWT refresh token
            example: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
```

**Response Schema** (Success - 200):

```yaml
responses:
  '200':
    description: Token refreshed successfully
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: true
            data:
              type: object
              properties:
                access_token:
                  type: string
                  example: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
                refresh_token:
                  type: string
                  example: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
                token_type:
                  type: string
                  example: "Bearer"
                expires_in:
                  type: integer
                  example: 86400
            message:
              type: string
              example: "Token refreshed successfully"
            timestamp:
              type: string
              format: date-time
```

**Response Schema** (Error - 401):

```yaml
  '401':
    description: Invalid or expired refresh token
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: false
            error:
              type: object
              properties:
                code:
                  type: string
                  example: "INVALID_REFRESH_TOKEN"
                message:
                  type: string
                  example: "Invalid or expired refresh token"
                details:
                  type: object
            timestamp:
              type: string
              format: date-time
```

**Error Codes**:
- `INVALID_REFRESH_TOKEN` (401): Invalid or expired refresh token

**Related Feature**: Token Refresh

**Related Database Tables**:
- `users`

---

## API 2.3: User Logout

**Endpoint**: `/api/v1/auth/logout`

**HTTP Method**: `POST`

**Operation ID**: `userLogout`

**Tags**: `authentication`

**Authentication Requirement**: Required

**Authorization Requirement**: Required

**Security**: `bearerAuth`

**Request Schema**: None (empty body)

**Response Schema** (Success - 200):

```yaml
responses:
  '200':
    description: Logout successful
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: true
            data:
              type: object
              nullable: true
            message:
              type: string
              example: "Logout successful"
            timestamp:
              type: string
              format: date-time
```

**Response Schema** (Error - 500):

```yaml
  '500':
    description: Logout failed
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: false
            error:
              type: object
              properties:
                code:
                  type: string
                  example: "LOGOUT_FAILED"
                message:
                  type: string
                  example: "Failed to logout"
                details:
                  type: object
            timestamp:
              type: string
              format: date-time
```

**Error Codes**:
- `LOGOUT_FAILED` (500): Failed to logout

**Related Feature**: User Logout

**Related Database Tables**:
- `users`

---

# Curriculum APIs

## API 3.1: View National Curriculum Plan (CP)

**Endpoint**: `/api/v1/curriculum/cp`

**HTTP Method**: `GET`

**Operation ID**: `viewCurriculumPlan`

**Tags**: `curriculum`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER, ADMINISTRATOR

**Security**: `bearerAuth`

**Query Parameters**:

```yaml
parameters:
  - name: subject_id
    in: query
    required: true
    schema:
      type: string
      format: uuid
    description: Subject identifier
    example: "sub_1234567890"
  - name: grade_level
    in: query
    required: true
    schema:
      type: string
      enum: ["10", "11", "12"]
    description: Grade level
    example: "10"
  - name: academic_year
    in: query
    required: true
    schema:
      type: string
      pattern: '^\d{4}$'
    description: Academic year
    example: "2026"
```

**Request Schema**: None (query parameters only)

**Response Schema** (Success - 200):

```yaml
responses:
  '200':
    description: Curriculum plan retrieved successfully
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: true
            data:
              type: object
              properties:
                cp:
                  type: object
                  properties:
                    id:
                      type: string
                      format: uuid
                      example: "cp_1234567890"
                    subject_id:
                      type: string
                      format: uuid
                      example: "sub_1234567890"
                    subject_name:
                      type: string
                      example: "Mathematics"
                    grade_level:
                      type: string
                      example: "10"
                    academic_year:
                      type: string
                      example: "2026"
                    learning_objectives:
                      type: array
                      items:
                        type: object
                        properties:
                          id:
                            type: string
                            format: uuid
                            example: "lo_1234567890"
                          code:
                            type: string
                            example: "LO.10.1"
                          description:
                            type: string
                            example: "Students understand algebraic concepts"
                          competency_code:
                            type: string
                            example: "CP.10.1"
                    competency_standards:
                      type: array
                      items:
                        type: object
                        properties:
                          id:
                            type: string
                            format: uuid
                            example: "cs_1234567890"
                          code:
                            type: string
                            example: "CS.10.1"
                          description:
                            type: string
                            example: "Algebraic thinking"
                    time_allocation:
                      type: object
                      properties:
                        total_hours:
                          type: integer
                          example: 120
                        hours_per_week:
                          type: integer
                          example: 4
            message:
              type: string
              example: "Curriculum plan retrieved successfully"
            timestamp:
              type: string
              format: date-time
```

**Response Schema** (Error - 404):

```yaml
  '404':
    description: Curriculum plan not found
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: false
            error:
              type: object
              properties:
                code:
                  type: string
                  example: "CP_NOT_FOUND"
                message:
                  type: string
                  example: "Curriculum plan not found for specified parameters"
                details:
                  type: object
            timestamp:
              type: string
              format: date-time
```

**Error Codes**:
- `CP_NOT_FOUND` (404): Curriculum plan not found for specified parameters

**Related Feature**: View Curriculum Plan

**Related Database Tables**:
- `cp_dimensions`
- `cp`

---

## API 3.2: Generate Teaching Plan (TP) with AI

**Endpoint**: `/api/v1/curriculum/tp/generate`

**HTTP Method**: `POST`

**Operation ID**: `generateTeachingPlan`

**Tags**: `curriculum`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Security**: `bearerAuth`

**Request Schema**:

```yaml
requestBody:
  required: true
  content:
    application/json:
      schema:
        type: object
        required:
          - cp_id
          - class_info
          - teaching_schedule
        properties:
          cp_id:
            type: string
            format: uuid
            description: Curriculum Plan identifier
            example: "cp_1234567890"
          class_info:
            type: object
            required:
              - grade
              - subject
              - academic_year
            properties:
              grade:
                type: string
                example: "10"
              subject:
                type: string
                maxLength: 255
                example: "Mathematics"
              academic_year:
                type: string
                pattern: '^\d{4}$'
                example: "2026"
          teaching_schedule:
            type: object
            required:
              - hours_per_week
              - weeks_per_semester
            properties:
              hours_per_week:
                type: integer
                minimum: 1
                maximum: 40
                example: 4
              weeks_per_semester:
                type: integer
                minimum: 1
                maximum: 52
                example: 18
          preferences:
            type: object
            properties:
              focus_areas:
                type: array
                items:
                  type: string
                example: ["algebra", "geometry"]
              teaching_style:
                type: string
                example: "interactive"
```

**Response Schema** (Success - 200):

```yaml
responses:
  '200':
    description: Teaching plan generated successfully
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: true
            data:
              type: object
              properties:
                tp:
                  type: object
                  properties:
                    id:
                      type: string
                      format: uuid
                      example: "tp_1234567890"
                    cp_id:
                      type: string
                      format: uuid
                      example: "cp_1234567890"
                    status:
                      type: string
                      enum: [DRAFT, UNDER_REVIEW, APPROVED, REJECTED]
                      example: "DRAFT"
                    learning_objectives:
                      type: array
                      items:
                        type: object
                        properties:
                          id:
                            type: string
                            format: uuid
                            example: "tlo_1234567890"
                          cp_objective_id:
                            type: string
                            format: uuid
                            example: "lo_1234567890"
                          description:
                            type: string
                            example: "Students understand algebraic concepts through interactive learning"
                          time_allocation_hours:
                            type: integer
                            example: 8
                    time_allocation:
                      type: object
                      properties:
                        total_hours:
                          type: integer
                          example: 72
                        hours_per_week:
                          type: integer
                          example: 4
                    prerequisites:
                      type: array
                      items:
                        type: object
                        properties:
                          objective_id:
                            type: string
                            format: uuid
                            example: "tlo_1234567890"
                          required_for:
                            type: array
                            items:
                              type: string
                              format: uuid
                            example: ["tlo_1234567891"]
                    ai_metadata:
                      type: object
                      properties:
                        confidence_score:
                          type: number
                          format: float
                          example: 0.92
                        generated_at:
                          type: string
                          format: date-time
                          example: "2026-06-03T12:00:00Z"
                        agent_version:
                          type: string
                          example: "1.0"
            message:
              type: string
              example: "Teaching plan generated successfully"
            timestamp:
              type: string
              format: date-time
```

**Response Schema** (Error - 500):

```yaml
  '500':
    description: Generation failed
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: false
            error:
              type: object
              properties:
                code:
                  type: string
                  example: "GENERATION_FAILED"
                message:
                  type: string
                  example: "Failed to generate teaching plan"
                details:
                  type: object
                  properties:
                    reason:
                      type: string
                      example: "AI service unavailable"
            timestamp:
              type: string
              format: date-time
```

**Error Codes**:
- `GENERATION_FAILED` (500): Failed to generate teaching plan

**Related Feature**: AI TP Generation

**Related Database Tables**:
- `tp`
- `cp`
- `ai_audit`
- `workflow_history`

---

## API 3.3: Get Teaching Plan by ID

**Endpoint**: `/api/v1/curriculum/tp/{tp_id}`

**HTTP Method**: `GET`

**Operation ID**: `getTeachingPlan`

**Tags**: `curriculum`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER, ADMINISTRATOR

**Security**: `bearerAuth`

**Path Parameters**:

```yaml
parameters:
  - name: tp_id
    in: path
    required: true
    schema:
      type: string
      format: uuid
    description: Teaching Plan identifier
    example: "tp_1234567890"
```

**Request Schema**: None

**Response Schema** (Success - 200):

```yaml
responses:
  '200':
    description: Teaching plan retrieved successfully
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: true
            data:
              type: object
              properties:
                tp:
                  type: object
                  properties:
                    id:
                      type: string
                      format: uuid
                      example: "tp_1234567890"
                    cp_id:
                      type: string
                      format: uuid
                      example: "cp_1234567890"
                    status:
                      type: string
                      enum: [DRAFT, UNDER_REVIEW, APPROVED, REJECTED]
                      example: "APPROVED"
                    learning_objectives:
                      type: array
                      items:
                        type: object
                        properties:
                          id:
                            type: string
                            format: uuid
                            example: "tlo_1234567890"
                          cp_objective_id:
                            type: string
                            format: uuid
                            example: "lo_1234567890"
                          description:
                            type: string
                            example: "Students understand algebraic concepts through interactive learning"
                          time_allocation_hours:
                            type: integer
                            example: 8
                    time_allocation:
                      type: object
                      properties:
                        total_hours:
                          type: integer
                          example: 72
                        hours_per_week:
                          type: integer
                          example: 4
                    created_at:
                      type: string
                      format: date-time
                      example: "2026-06-03T10:00:00Z"
                    updated_at:
                      type: string
                      format: date-time
                      example: "2026-06-03T11:00:00Z"
                    approved_at:
                      type: string
                      format: date-time
                      example: "2026-06-03T12:00:00Z"
                    approved_by:
                      type: string
                      format: uuid
                      example: "usr_1234567890"
            message:
              type: string
              example: "Teaching plan retrieved successfully"
            timestamp:
              type: string
              format: date-time
```

**Response Schema** (Error - 404):

```yaml
  '404':
    description: Teaching plan not found
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: false
            error:
              type: object
              properties:
                code:
                  type: string
                  example: "TP_NOT_FOUND"
                message:
                  type: string
                  example: "Teaching plan not found"
                details:
                  type: object
            timestamp:
              type: string
              format: date-time
```

**Error Codes**:
- `TP_NOT_FOUND` (404): Teaching plan not found

**Related Feature**: View Teaching Plan

**Related Database Tables**:
- `tp`

---

## API 3.4: Update Teaching Plan

**Endpoint**: `/api/v1/curriculum/tp/{tp_id}`

**HTTP Method**: `PUT`

**Operation ID**: `updateTeachingPlan`

**Tags**: `curriculum`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Security**: `bearerAuth`

**Path Parameters**:

```yaml
parameters:
  - name: tp_id
    in: path
    required: true
    schema:
      type: string
      format: uuid
    description: Teaching Plan identifier
    example: "tp_1234567890"
```

**Request Schema**:

```yaml
requestBody:
  required: true
  content:
    application/json:
      schema:
        type: object
        required:
          - learning_objectives
          - time_allocation
        properties:
          learning_objectives:
            type: array
            items:
              type: object
              required:
                - id
                - description
                - time_allocation_hours
              properties:
                id:
                  type: string
                  format: uuid
                  example: "tlo_1234567890"
                description:
                  type: string
                  maxLength: 1000
                  example: "Updated description"
                time_allocation_hours:
                  type: integer
                  minimum: 1
                  maximum: 100
                  example: 10
          time_allocation:
            type: object
            required:
              - total_hours
              - hours_per_week
            properties:
              total_hours:
                type: integer
                minimum: 1
                maximum: 1000
                example: 80
              hours_per_week:
                type: integer
                minimum: 1
                maximum: 40
                example: 4
```

**Response Schema** (Success - 200):

```yaml
responses:
  '200':
    description: Teaching plan updated successfully
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: true
            data:
              type: object
              properties:
                tp:
                  type: object
                  properties:
                    id:
                      type: string
                      format: uuid
                      example: "tp_1234567890"
                    status:
                      type: string
                      enum: [DRAFT, UNDER_REVIEW, APPROVED, REJECTED]
                      example: "DRAFT"
                    updated_at:
                      type: string
                      format: date-time
                      example: "2026-06-03T13:00:00Z"
            message:
              type: string
              example: "Teaching plan updated successfully"
            timestamp:
              type: string
              format: date-time
```

**Response Schema** (Error - 500):

```yaml
  '500':
    description: Update failed
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: false
            error:
              type: object
              properties:
                code:
                  type: string
                  example: "TP_UPDATE_FAILED"
                message:
                  type: string
                  example: "Failed to update teaching plan"
                details:
                  type: object
            timestamp:
              type: string
              format: date-time
```

**Error Codes**:
- `TP_UPDATE_FAILED` (500): Failed to update teaching plan

**Related Feature**: Edit Teaching Plan

**Related Database Tables**:
- `tp`
- `workflow_history`

---

## API 3.5: Approve Teaching Plan

**Endpoint**: `/api/v1/curriculum/tp/{tp_id}/approve`

**HTTP Method**: `POST`

**Operation ID**: `approveTeachingPlan`

**Tags**: `curriculum`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Security**: `bearerAuth`

**Path Parameters**:

```yaml
parameters:
  - name: tp_id
    in: path
    required: true
    schema:
      type: string
      format: uuid
    description: Teaching Plan identifier
    example: "tp_1234567890"
```

**Request Schema**: None

**Response Schema** (Success - 200):

```yaml
responses:
  '200':
    description: Teaching plan approved successfully
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: true
            data:
              type: object
              properties:
                tp:
                  type: object
                  properties:
                    id:
                      type: string
                      format: uuid
                      example: "tp_1234567890"
                    status:
                      type: string
                      enum: [DRAFT, UNDER_REVIEW, APPROVED, REJECTED]
                      example: "APPROVED"
                    approved_at:
                      type: string
                      format: date-time
                      example: "2026-06-03T14:00:00Z"
                    approved_by:
                      type: string
                      format: uuid
                      example: "usr_1234567890"
            message:
              type: string
              example: "Teaching plan approved successfully"
            timestamp:
              type: string
              format: date-time
```

**Response Schema** (Error - 500):

```yaml
  '500':
    description: Approval failed
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: false
            error:
              type: object
              properties:
                code:
                  type: string
                  example: "TP_APPROVAL_FAILED"
                message:
                  type: string
                  example: "Failed to approve teaching plan"
                details:
                  type: object
                  properties:
                    reason:
                      type: string
                      example: "Teaching plan must be in DRAFT status"
            timestamp:
              type: string
              format: date-time
```

**Error Codes**:
- `TP_APPROVAL_FAILED` (500): Failed to approve teaching plan

**Related Feature**: Approve Teaching Plan

**Related Database Tables**:
- `tp`
- `workflow_history`

---

## API 3.6: List Teaching Plans

**Endpoint**: `/api/v1/curriculum/tp`

**HTTP Method**: `GET`

**Operation ID**: `listTeachingPlans`

**Tags**: `curriculum`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER, ADMINISTRATOR

**Security**: `bearerAuth`

**Query Parameters**:

```yaml
parameters:
  - name: status
    in: query
    required: false
    schema:
      type: string
      enum: [DRAFT, APPROVED, REJECTED]
    description: Filter by status
  - name: subject_id
    in: query
    required: false
    schema:
      type: string
      format: uuid
    description: Filter by subject
  - name: grade_level
    in: query
    required: false
    schema:
      type: string
      enum: ["10", "11", "12"]
    description: Filter by grade level
  - name: page
    in: query
    required: false
    schema:
      type: integer
      minimum: 1
      default: 1
    description: Page number
  - name: limit
    in: query
    required: false
    schema:
      type: integer
      minimum: 1
      maximum: 100
      default: 20
    description: Items per page
```

**Request Schema**: None (query parameters only)

**Response Schema** (Success - 200):

```yaml
responses:
  '200':
    description: Teaching plans retrieved successfully
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: true
            data:
              type: object
              properties:
                tps:
                  type: array
                  items:
                    type: object
                    properties:
                      id:
                        type: string
                        format: uuid
                        example: "tp_1234567890"
                      subject:
                        type: string
                        example: "Mathematics"
                      grade_level:
                        type: string
                        example: "10"
                      status:
                        type: string
                        enum: [DRAFT, APPROVED, REJECTED]
                        example: "APPROVED"
                      created_at:
                        type: string
                        format: date-time
                        example: "2026-06-03T10:00:00Z"
                pagination:
                  type: object
                  properties:
                    total:
                      type: integer
                      example: 50
                    page:
                      type: integer
                      example: 1
                    limit:
                      type: integer
                      example: 20
                    total_pages:
                      type: integer
                      example: 3
            message:
              type: string
              example: "Teaching plans retrieved successfully"
            timestamp:
              type: string
              format: date-time
```

**Error Codes**: None specific to this endpoint (standard error codes apply)

**Related Feature**: List Teaching Plans

**Related Database Tables**:
- `tp`
- `cp`

---

# Learning Planning APIs

## API 4.1: Generate Annual Teaching Plan (ATP) with AI

**Endpoint**: `/api/v1/learning-planning/atp/generate`

**HTTP Method**: `POST`

**Operation ID**: `generateAnnualTeachingPlan`

**Tags**: `learning-planning`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Security**: `bearerAuth`

**Request Schema**:

```yaml
requestBody:
  required: true
  content:
    application/json:
      schema:
        type: object
        required:
          - tp_id
          - academic_calendar
          - class_schedule
        properties:
          tp_id:
            type: string
            format: uuid
            description: Teaching Plan identifier
            example: "tp_1234567890"
          academic_calendar:
            type: object
            required:
              - start_date
              - end_date
            properties:
              start_date:
                type: string
                format: date
                example: "2026-07-15"
              end_date:
                type: string
                format: date
                example: "2026-12-15"
              holidays:
                type: array
                items:
                  type: string
                  format: date
                example: ["2026-08-17", "2026-12-25"]
          class_schedule:
            type: object
            required:
              - days_per_week
              - periods_per_day
              - available_hours_per_week
            properties:
              days_per_week:
                type: integer
                minimum: 1
                maximum: 7
                example: 5
              periods_per_day:
                type: integer
                minimum: 1
                maximum: 12
                example: 8
              available_hours_per_week:
                type: integer
                minimum: 1
                maximum: 40
                example: 4
```

**Response Schema** (Success - 200):

```yaml
responses:
  '200':
    description: Annual teaching plan generated successfully
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: true
            data:
              type: object
              properties:
                atp:
                  type: object
                  properties:
                    id:
                      type: string
                      format: uuid
                      example: "atp_1234567890"
                    tp_id:
                      type: string
                      format: uuid
                      example: "tp_1234567890"
                    status:
                      type: string
                      enum: [DRAFT, UNDER_REVIEW, APPROVED, REJECTED]
                      example: "DRAFT"
                    weekly_sequence:
                      type: array
                      items:
                        type: object
                        properties:
                          week:
                            type: integer
                            example: 1
                          topics:
                            type: array
                            items:
                              type: object
                              properties:
                                learning_objective_id:
                                  type: string
                                  format: uuid
                                  example: "tlo_1234567890"
                                hours:
                                  type: integer
                                  example: 4
                                start_date:
                                  type: string
                                  format: date
                                  example: "2026-07-15"
                                end_date:
                                  type: string
                                  format: date
                                  example: "2026-07-19"
                    assessment_schedule:
                      type: array
                      items:
                        type: object
                        properties:
                          week:
                            type: integer
                            example: 4
                          type:
                            type: string
                            enum: [FORMATIVE, SUMMATIVE]
                            example: "FORMATIVE"
                          topics_covered:
                            type: array
                            items:
                              type: string
                              format: uuid
                            example: ["tlo_1234567890", "tlo_1234567891"]
                    ai_metadata:
                      type: object
                      properties:
                        confidence_score:
                          type: number
                          format: float
                          example: 0.89
                        generated_at:
                          type: string
                          format: date-time
                          example: "2026-06-03T12:00:00Z"
                        agent_version:
                          type: string
                          example: "1.0"
            message:
              type: string
              example: "Annual teaching plan generated successfully"
            timestamp:
              type: string
              format: date-time
```

**Error Codes**:
- `GENERATION_FAILED` (500): Failed to generate annual teaching plan

**Related Feature**: AI ATP Generation

**Related Database Tables**:
- `atp`
- `tp`
- `ai_audit`
- `workflow_history`

---

## API 4.2: Get Annual Teaching Plan by ID

**Endpoint**: `/api/v1/learning-planning/atp/{atp_id}`

**HTTP Method**: `GET`

**Operation ID**: `getAnnualTeachingPlan`

**Tags**: `learning-planning`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER, ADMINISTRATOR

**Security**: `bearerAuth`

**Path Parameters**:

```yaml
parameters:
  - name: atp_id
    in: path
    required: true
    schema:
      type: string
      format: uuid
    description: Annual Teaching Plan identifier
    example: "atp_1234567890"
```

**Request Schema**: None

**Response Schema** (Success - 200):

```yaml
responses:
  '200':
    description: Annual teaching plan retrieved successfully
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: true
            data:
              type: object
              properties:
                atp:
                  type: object
                  properties:
                    id:
                      type: string
                      format: uuid
                      example: "atp_1234567890"
                    tp_id:
                      type: string
                      format: uuid
                      example: "tp_1234567890"
                    status:
                      type: string
                      enum: [DRAFT, UNDER_REVIEW, APPROVED, REJECTED]
                      example: "APPROVED"
                    weekly_sequence:
                      type: array
                      items:
                        type: object
                        properties:
                          week:
                            type: integer
                            example: 1
                          topics:
                            type: array
                            items:
                              type: object
                              properties:
                                learning_objective_id:
                                  type: string
                                  format: uuid
                                  example: "tlo_1234567890"
                                hours:
                                  type: integer
                                  example: 4
                                start_date:
                                  type: string
                                  format: date
                                  example: "2026-07-15"
                                end_date:
                                  type: string
                                  format: date
                                  example: "2026-07-19"
                    created_at:
                      type: string
                      format: date-time
                      example: "2026-06-03T10:00:00Z"
                    updated_at:
                      type: string
                      format: date-time
                      example: "2026-06-03T11:00:00Z"
                    approved_at:
                      type: string
                      format: date-time
                      example: "2026-06-03T12:00:00Z"
                    approved_by:
                      type: string
                      format: uuid
                      example: "usr_1234567890"
            message:
              type: string
              example: "Annual teaching plan retrieved successfully"
            timestamp:
              type: string
              format: date-time
```

**Error Codes**:
- `ATP_NOT_FOUND` (404): Annual teaching plan not found

**Related Feature**: View Annual Teaching Plan

**Related Database Tables**:
- `atp`

---

## API 4.3: Update Annual Teaching Plan

**Endpoint**: `/api/v1/learning-planning/atp/{atp_id}`

**HTTP Method**: `PUT`

**Operation ID**: `updateAnnualTeachingPlan`

**Tags**: `learning-planning`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Security**: `bearerAuth`

**Path Parameters**:

```yaml
parameters:
  - name: atp_id
    in: path
    required: true
    schema:
      type: string
      format: uuid
    description: Annual Teaching Plan identifier
    example: "atp_1234567890"
```

**Request Schema**:

```yaml
requestBody:
  required: true
  content:
    application/json:
      schema:
        type: object
        required:
          - weekly_sequence
        properties:
          weekly_sequence:
            type: array
            items:
              type: object
              required:
                - week
                - topics
              properties:
                week:
                  type: integer
                  minimum: 1
                  example: 1
                topics:
                  type: array
                  items:
                    type: object
                    required:
                      - learning_objective_id
                      - hours
                      - start_date
                      - end_date
                    properties:
                      learning_objective_id:
                        type: string
                        format: uuid
                        example: "tlo_1234567890"
                      hours:
                        type: integer
                        minimum: 1
                        maximum: 40
                        example: 5
                      start_date:
                        type: string
                        format: date
                        example: "2026-07-15"
                      end_date:
                        type: string
                        format: date
                        example: "2026-07-19"
```

**Response Schema** (Success - 200):

```yaml
responses:
  '200':
    description: Annual teaching plan updated successfully
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: true
            data:
              type: object
              properties:
                atp:
                  type: object
                  properties:
                    id:
                      type: string
                      format: uuid
                      example: "atp_1234567890"
                    status:
                      type: string
                      enum: [DRAFT, UNDER_REVIEW, APPROVED, REJECTED]
                      example: "DRAFT"
                    updated_at:
                      type: string
                      format: date-time
                      example: "2026-06-03T13:00:00Z"
            message:
              type: string
              example: "Annual teaching plan updated successfully"
            timestamp:
              type: string
              format: date-time
```

**Error Codes**: None specific to this endpoint (standard error codes apply)

**Related Feature**: Edit Annual Teaching Plan

**Related Database Tables**:
- `atp`
- `workflow_history`

---

## API 4.4: Approve Annual Teaching Plan

**Endpoint**: `/api/v1/learning-planning/atp/{atp_id}/approve`

**HTTP Method**: `POST`

**Operation ID**: `approveAnnualTeachingPlan`

**Tags**: `learning-planning`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Security**: `bearerAuth`

**Path Parameters**:

```yaml
parameters:
  - name: atp_id
    in: path
    required: true
    schema:
      type: string
      format: uuid
    description: Annual Teaching Plan identifier
    example: "atp_1234567890"
```

**Request Schema**: None

**Response Schema** (Success - 200):

```yaml
responses:
  '200':
    description: Annual teaching plan approved successfully
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: true
            data:
              type: object
              properties:
                atp:
                  type: object
                  properties:
                    id:
                      type: string
                      format: uuid
                      example: "atp_1234567890"
                    status:
                      type: string
                      enum: [DRAFT, UNDER_REVIEW, APPROVED, REJECTED]
                      example: "APPROVED"
                    approved_at:
                      type: string
                      format: date-time
                      example: "2026-06-03T14:00:00Z"
                    approved_by:
                      type: string
                      format: uuid
                      example: "usr_1234567890"
            message:
              type: string
              example: "Annual teaching plan approved successfully"
            timestamp:
              type: string
              format: date-time
```

**Error Codes**: None specific to this endpoint (standard error codes apply)

**Related Feature**: Approve Annual Teaching Plan

**Related Database Tables**:
- `atp`
- `workflow_history`

---

## API 4.5: Generate Modul Ajar with AI

**Endpoint**: `/api/v1/learning-planning/modul-ajar/generate`

**HTTP Method**: `POST`

**Operation ID**: `generateModulAjar`

**Tags**: `learning-planning`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Security**: `bearerAuth`

**Request Schema**:

```yaml
requestBody:
  required: true
  content:
    application/json:
      schema:
        type: object
        required:
          - atp_id
          - week
          - topic
        properties:
          atp_id:
            type: string
            format: uuid
            description: Annual Teaching Plan identifier
            example: "atp_1234567890"
          week:
            type: integer
            minimum: 1
            description: Week number
            example: 1
          topic:
            type: object
            required:
              - learning_objective_id
              - title
            properties:
              learning_objective_id:
                type: string
                format: uuid
                example: "tlo_1234567890"
              title:
                type: string
                maxLength: 255
                example: "Introduction to Algebra"
          resources:
            type: object
            properties:
              textbooks:
                type: array
                items:
                  type: string
                example: ["Algebra Fundamentals"]
              materials:
                type: array
                items:
                  type: string
                example: ["graph paper", "calculators"]
          class_characteristics:
            type: object
            properties:
              student_count:
                type: integer
                minimum: 1
                example: 30
              skill_level:
                type: string
                example: "intermediate"
```

**Response Schema** (Success - 200):

```yaml
responses:
  '200':
    description: Modul Ajar generated successfully
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: true
            data:
              type: object
              properties:
                modul_ajar:
                  type: object
                  properties:
                    id:
                      type: string
                      format: uuid
                      example: "ma_1234567890"
                    atp_id:
                      type: string
                      format: uuid
                      example: "atp_1234567890"
                    week:
                      type: integer
                      example: 1
                    status:
                      type: string
                      enum: [DRAFT, UNDER_REVIEW, APPROVED, REJECTED]
                      example: "DRAFT"
                    learning_activities:
                      type: array
                      items:
                        type: object
                        properties:
                          id:
                            type: string
                            format: uuid
                            example: "la_1234567890"
                          sequence:
                            type: integer
                            example: 1
                          activity_type:
                            type: string
                            enum: [INTRODUCTION, PRACTICE, ASSESSMENT, CLOSING]
                            example: "INTRODUCTION"
                          description:
                            type: string
                            example: "Teacher introduces algebraic concepts with real-world examples"
                          duration_minutes:
                            type: integer
                            example: 15
                          resources:
                            type: array
                            items:
                              type: string
                            example: ["whiteboard", "markers"]
                    resource_requirements:
                      type: array
                      items:
                        type: object
                        properties:
                          resource:
                            type: string
                            example: "graph paper"
                          quantity:
                            type: integer
                            example: 30
                    assessment_methods:
                      type: array
                      items:
                        type: object
                        properties:
                          type:
                            type: string
                            enum: [FORMATIVE, SUMMATIVE]
                            example: "FORMATIVE"
                          description:
                            type: string
                            example: "Observe student participation in practice activities"
                    ai_metadata:
                      type: object
                      properties:
                        confidence_score:
                          type: number
                          format: float
                          example: 0.91
                        generated_at:
                          type: string
                          format: date-time
                          example: "2026-06-03T12:00:00Z"
                        agent_version:
                          type: string
                          example: "1.0"
            message:
              type: string
              example: "Modul Ajar generated successfully"
            timestamp:
              type: string
              format: date-time
```

**Error Codes**:
- `GENERATION_FAILED` (500): Failed to generate Modul Ajar

**Related Feature**: AI Modul Ajar Generation

**Related Database Tables**:
- `modul_ajar`
- `atp`
- `ai_audit`
- `workflow_history`

---

## API 4.6: Get Modul Ajar by ID

**Endpoint**: `/api/v1/learning-planning/modul-ajar/{modul_ajar_id}`

**HTTP Method**: `GET`

**Operation ID**: `getModulAjar`

**Tags**: `learning-planning`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER, ADMINISTRATOR

**Security**: `bearerAuth`

**Path Parameters**:

```yaml
parameters:
  - name: modul_ajar_id
    in: path
    required: true
    schema:
      type: string
      format: uuid
    description: Modul Ajar identifier
    example: "ma_1234567890"
```

**Request Schema**: None

**Response Schema** (Success - 200):

```yaml
responses:
  '200':
    description: Modul Ajar retrieved successfully
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: true
            data:
              type: object
              properties:
                modul_ajar:
                  type: object
                  properties:
                    id:
                      type: string
                      format: uuid
                      example: "ma_1234567890"
                    atp_id:
                      type: string
                      format: uuid
                      example: "atp_1234567890"
                    week:
                      type: integer
                      example: 1
                    status:
                      type: string
                      enum: [DRAFT, UNDER_REVIEW, APPROVED, REJECTED]
                      example: "APPROVED"
                    learning_activities:
                      type: array
                      items:
                        type: object
                        properties:
                          id:
                            type: string
                            format: uuid
                            example: "la_1234567890"
                          sequence:
                            type: integer
                            example: 1
                          activity_type:
                            type: string
                            enum: [INTRODUCTION, PRACTICE, ASSESSMENT, CLOSING]
                            example: "INTRODUCTION"
                          description:
                            type: string
                            example: "Teacher introduces algebraic concepts with real-world examples"
                          duration_minutes:
                            type: integer
                            example: 15
                          resources:
                            type: array
                            items:
                              type: string
                            example: ["whiteboard", "markers"]
                    created_at:
                      type: string
                      format: date-time
                      example: "2026-06-03T10:00:00Z"
                    updated_at:
                      type: string
                      format: date-time
                      example: "2026-06-03T11:00:00Z"
                    approved_at:
                      type: string
                      format: date-time
                      example: "2026-06-03T12:00:00Z"
                    approved_by:
                      type: string
                      format: uuid
                      example: "usr_1234567890"
            message:
              type: string
              example: "Modul Ajar retrieved successfully"
            timestamp:
              type: string
              format: date-time
```

**Error Codes**:
- `MODUL_AJAR_NOT_FOUND` (404): Modul Ajar not found

**Related Feature**: View Modul Ajar

**Related Database Tables**:
- `modul_ajar`

---

## API 4.7: Update Modul Ajar

**Endpoint**: `/api/v1/learning-planning/modul-ajar/{modul_ajar_id}`

**HTTP Method**: `PUT`

**Operation ID**: `updateModulAjar`

**Tags**: `learning-planning`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Security**: `bearerAuth`

**Path Parameters**:

```yaml
parameters:
  - name: modul_ajar_id
    in: path
    required: true
    schema:
      type: string
      format: uuid
    description: Modul Ajar identifier
    example: "ma_1234567890"
```

**Request Schema**:

```yaml
requestBody:
  required: true
  content:
    application/json:
      schema:
        type: object
        required:
          - learning_activities
        properties:
          learning_activities:
            type: array
            items:
              type: object
              required:
                - id
                - sequence
                - activity_type
                - description
                - duration_minutes
              properties:
                id:
                  type: string
                  format: uuid
                  example: "la_1234567890"
                sequence:
                  type: integer
                  minimum: 1
                  example: 1
                activity_type:
                  type: string
                  enum: [INTRODUCTION, PRACTICE, ASSESSMENT, CLOSING]
                  example: "INTRODUCTION"
                description:
                  type: string
                  maxLength: 1000
                  example: "Updated description"
                duration_minutes:
                  type: integer
                  minimum: 1
                  maximum: 180
                  example: 20
                resources:
                  type: array
                  items:
                    type: string
                  example: ["whiteboard", "markers", "projector"]
```

**Response Schema** (Success - 200):

```yaml
responses:
  '200':
    description: Modul Ajar updated successfully
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: true
            data:
              type: object
              properties:
                modul_ajar:
                  type: object
                  properties:
                    id:
                      type: string
                      format: uuid
                      example: "ma_1234567890"
                    status:
                      type: string
                      enum: [DRAFT, UNDER_REVIEW, APPROVED, REJECTED]
                      example: "DRAFT"
                    updated_at:
                      type: string
                      format: date-time
                      example: "2026-06-03T13:00:00Z"
            message:
              type: string
              example: "Modul Ajar updated successfully"
            timestamp:
              type: string
              format: date-time
```

**Error Codes**: None specific to this endpoint (standard error codes apply)

**Related Feature**: Edit Modul Ajar

**Related Database Tables**:
- `modul_ajar`
- `workflow_history`

---

## API 4.8: Approve Modul Ajar

**Endpoint**: `/api/v1/learning-planning/modul-ajar/{modul_ajar_id}/approve`

**HTTP Method**: `POST`

**Operation ID**: `approveModulAjar`

**Tags**: `learning-planning`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Security**: `bearerAuth`

**Path Parameters**:

```yaml
parameters:
  - name: modul_ajar_id
    in: path
    required: true
    schema:
      type: string
      format: uuid
    description: Modul Ajar identifier
    example: "ma_1234567890"
```

**Request Schema**: None

**Response Schema** (Success - 200):

```yaml
responses:
  '200':
    description: Modul Ajar approved successfully
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: true
            data:
              type: object
              properties:
                modul_ajar:
                  type: object
                  properties:
                    id:
                      type: string
                      format: uuid
                      example: "ma_1234567890"
                    status:
                      type: string
                      enum: [DRAFT, UNDER_REVIEW, APPROVED, REJECTED]
                      example: "APPROVED"
                    approved_at:
                      type: string
                      format: date-time
                      example: "2026-06-03T14:00:00Z"
                    approved_by:
                      type: string
                      format: uuid
                      example: "usr_1234567890"
            message:
              type: string
              example: "Modul Ajar approved successfully"
            timestamp:
              type: string
              format: date-time
```

**Error Codes**: None specific to this endpoint (standard error codes apply)

**Related Feature**: Approve Modul Ajar

**Related Database Tables**:
- `modul_ajar`
- `workflow_history`

---

# Assessment APIs

## API 5.1: Generate Assessment with AI

**Endpoint**: `/api/v1/assessment/generate`

**HTTP Method**: `POST`

**Operation ID**: `generateAssessment`

**Tags**: `assessment`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Security**: `bearerAuth`

**Request Schema**:

```yaml
requestBody:
  required: true
  content:
    application/json:
      schema:
        type: object
        required:
          - modul_ajar_id
          - assessment_type
          - question_count
          - difficulty_level
          - time_allocation_minutes
        properties:
          modul_ajar_id:
            type: string
            format: uuid
            description: Modul Ajar identifier
            example: "ma_1234567890"
          assessment_type:
            type: string
            enum: [FORMATIVE, SUMMATIVE]
            example: "FORMATIVE"
          question_count:
            type: integer
            minimum: 1
            maximum: 50
            example: 10
          difficulty_level:
            type: string
            enum: [EASY, MEDIUM, HARD]
            example: "MEDIUM"
          time_allocation_minutes:
            type: integer
            minimum: 5
            maximum: 180
            example: 45
```

**Response Schema** (Success - 200):

```yaml
responses:
  '200':
    description: Assessment generated successfully
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: true
            data:
              type: object
              properties:
                assessment:
                  type: object
                  properties:
                    id:
                      type: string
                      format: uuid
                      example: "ass_1234567890"
                    modul_ajar_id:
                      type: string
                      format: uuid
                      example: "ma_1234567890"
                    status:
                      type: string
                      enum: [DRAFT, UNDER_REVIEW, APPROVED, REJECTED]
                      example: "DRAFT"
                    assessment_type:
                      type: string
                      enum: [FORMATIVE, SUMMATIVE]
                      example: "FORMATIVE"
                    assessment_items:
                      type: array
                      items:
                        type: object
                        properties:
                          id:
                            type: string
                            format: uuid
                            example: "ai_1234567890"
                          type:
                            type: string
                            enum: [MULTIPLE_CHOICE, ESSAY, SHORT_ANSWER]
                            example: "MULTIPLE_CHOICE"
                          question:
                            type: string
                            example: "What is the value of x in the equation 2x + 5 = 15?"
                          options:
                            type: array
                            items:
                              type: string
                            example: ["x = 5", "x = 10", "x = 15", "x = 20"]
                          correct_answer:
                            type: string
                            example: "x = 5"
                          points:
                            type: integer
                            example: 2
                          learning_objective_id:
                            type: string
                            format: uuid
                            example: "tlo_1234567890"
                    answer_key:
                      type: object
                      additionalProperties:
                        type: string
                      example:
                        ai_1234567890: "x = 5"
                        ai_1234567891: "Rubric-based evaluation"
                    scoring_guidelines:
                      type: object
                      properties:
                        total_points:
                          type: integer
                          example: 12
                        passing_score:
                          type: integer
                          example: 7
                    ai_metadata:
                      type: object
                      properties:
                        confidence_score:
                          type: number
                          format: float
                          example: 0.88
                        generated_at:
                          type: string
                          format: date-time
                          example: "2026-06-03T12:00:00Z"
                        agent_version:
                          type: string
                          example: "1.0"
            message:
              type: string
              example: "Assessment generated successfully"
            timestamp:
              type: string
              format: date-time
```

**Error Codes**:
- `GENERATION_FAILED` (500): Failed to generate assessment

**Related Feature**: AI Assessment Generation

**Related Database Tables**:
- `assessment`
- `modul_ajar`
- `ai_audit`
- `workflow_history`

---

## API 5.2: Get Assessment by ID

**Endpoint**: `/api/v1/assessment/{assessment_id}`

**HTTP Method**: `GET`

**Operation ID**: `getAssessment`

**Tags**: `assessment`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER, ADMINISTRATOR

**Security**: `bearerAuth`

**Path Parameters**:

```yaml
parameters:
  - name: assessment_id
    in: path
    required: true
    schema:
      type: string
      format: uuid
    description: Assessment identifier
    example: "ass_1234567890"
```

**Request Schema**: None

**Response Schema** (Success - 200):

```yaml
responses:
  '200':
    description: Assessment retrieved successfully
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: true
            data:
              type: object
              properties:
                assessment:
                  type: object
                  properties:
                    id:
                      type: string
                      format: uuid
                      example: "ass_1234567890"
                    modul_ajar_id:
                      type: string
                      format: uuid
                      example: "ma_1234567890"
                    status:
                      type: string
                      enum: [DRAFT, UNDER_REVIEW, APPROVED, REJECTED]
                      example: "APPROVED"
                    assessment_type:
                      type: string
                      enum: [FORMATIVE, SUMMATIVE]
                      example: "FORMATIVE"
                    assessment_items:
                      type: array
                      items:
                        type: object
                        properties:
                          id:
                            type: string
                            format: uuid
                            example: "ai_1234567890"
                          type:
                            type: string
                            enum: [MULTIPLE_CHOICE, ESSAY, SHORT_ANSWER]
                            example: "MULTIPLE_CHOICE"
                          question:
                            type: string
                            example: "What is the value of x in the equation 2x + 5 = 15?"
                          options:
                            type: array
                            items:
                              type: string
                            example: ["x = 5", "x = 10", "x = 15", "x = 20"]
                          correct_answer:
                            type: string
                            example: "x = 5"
                          points:
                            type: integer
                            example: 2
                    created_at:
                      type: string
                      format: date-time
                      example: "2026-06-03T10:00:00Z"
                    updated_at:
                      type: string
                      format: date-time
                      example: "2026-06-03T11:00:00Z"
                    approved_at:
                      type: string
                      format: date-time
                      example: "2026-06-03T12:00:00Z"
                    approved_by:
                      type: string
                      format: uuid
                      example: "usr_1234567890"
            message:
              type: string
              example: "Assessment retrieved successfully"
            timestamp:
              type: string
              format: date-time
```

**Error Codes**:
- `ASSESSMENT_NOT_FOUND` (404): Assessment not found

**Related Feature**: View Assessment

**Related Database Tables**:
- `assessment`

---

## API 5.3: Update Assessment

**Endpoint**: `/api/v1/assessment/{assessment_id}`

**HTTP Method**: `PUT`

**Operation ID**: `updateAssessment`

**Tags**: `assessment`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Security**: `bearerAuth`

**Path Parameters**:

```yaml
parameters:
  - name: assessment_id
    in: path
    required: true
    schema:
      type: string
      format: uuid
    description: Assessment identifier
    example: "ass_1234567890"
```

**Request Schema**:

```yaml
requestBody:
  required: true
  content:
    application/json:
      schema:
        type: object
        required:
          - assessment_items
        properties:
          assessment_items:
            type: array
            items:
              type: object
              required:
                - id
                - type
                - question
                - points
              properties:
                id:
                  type: string
                  format: uuid
                  example: "ai_1234567890"
                type:
                  type: string
                  enum: [MULTIPLE_CHOICE, ESSAY, SHORT_ANSWER]
                  example: "MULTIPLE_CHOICE"
                question:
                  type: string
                  maxLength: 1000
                  example: "Updated question"
                options:
                  type: array
                  items:
                    type: string
                  example: ["x = 5", "x = 10", "x = 15", "x = 20"]
                correct_answer:
                  type: string
                  maxLength: 500
                  example: "x = 5"
                points:
                  type: integer
                  minimum: 1
                  maximum: 20
                  example: 3
```

**Response Schema** (Success - 200):

```yaml
responses:
  '200':
    description: Assessment updated successfully
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: true
            data:
              type: object
              properties:
                assessment:
                  type: object
                  properties:
                    id:
                      type: string
                      format: uuid
                      example: "ass_1234567890"
                    status:
                      type: string
                      enum: [DRAFT, UNDER_REVIEW, APPROVED, REJECTED]
                      example: "DRAFT"
                    updated_at:
                      type: string
                      format: date-time
                      example: "2026-06-03T13:00:00Z"
            message:
              type: string
              example: "Assessment updated successfully"
            timestamp:
              type: string
              format: date-time
```

**Error Codes**: None specific to this endpoint (standard error codes apply)

**Related Feature**: Edit Assessment

**Related Database Tables**:
- `assessment`
- `workflow_history`

---

## API 5.4: Approve Assessment

**Endpoint**: `/api/v1/assessment/{assessment_id}/approve`

**HTTP Method**: `POST`

**Operation ID**: `approveAssessment`

**Tags**: `assessment`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Security**: `bearerAuth`

**Path Parameters**:

```yaml
parameters:
  - name: assessment_id
    in: path
    required: true
    schema:
      type: string
      format: uuid
    description: Assessment identifier
    example: "ass_1234567890"
```

**Request Schema**: None

**Response Schema** (Success - 200):

```yaml
responses:
  '200':
    description: Assessment approved successfully
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: true
            data:
              type: object
              properties:
                assessment:
                  type: object
                  properties:
                    id:
                      type: string
                      format: uuid
                      example: "ass_1234567890"
                    status:
                      type: string
                      enum: [DRAFT, UNDER_REVIEW, APPROVED, REJECTED]
                      example: "APPROVED"
                    approved_at:
                      type: string
                      format: date-time
                      example: "2026-06-03T14:00:00Z"
                    approved_by:
                      type: string
                      format: uuid
                      example: "usr_1234567890"
            message:
              type: string
              example: "Assessment approved successfully"
            timestamp:
              type: string
              format: date-time
```

**Error Codes**: None specific to this endpoint (standard error codes apply)

**Related Feature**: Approve Assessment

**Related Database Tables**:
- `assessment`
- `workflow_history`

---

## API 5.5: Generate Rubric with AI

**Endpoint**: `/api/v1/assessment/rubric/generate`

**HTTP Method**: `POST`

**Operation ID**: `generateRubric`

**Tags**: `assessment`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Security**: `bearerAuth`

**Request Schema**:

```yaml
requestBody:
  required: true
  content:
    application/json:
      schema:
        type: object
        required:
          - assessment_id
          - rubric_type
          - performance_levels
          - criteria_categories
        properties:
          assessment_id:
            type: string
            format: uuid
            description: Assessment identifier
            example: "ass_1234567890"
          rubric_type:
            type: string
            enum: [ANALYTIC, HOLISTIC]
            example: "ANALYTIC"
          performance_levels:
            type: integer
            minimum: 2
            maximum: 5
            example: 4
          criteria_categories:
            type: array
            items:
              type: string
            minItems: 1
            example: ["content", "organization", "mechanics"]
```

**Response Schema** (Success - 200):

```yaml
responses:
  '200':
    description: Rubric generated successfully
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: true
            data:
              type: object
              properties:
                rubric:
                  type: object
                  properties:
                    id:
                      type: string
                      format: uuid
                      example: "rub_1234567890"
                    assessment_id:
                      type: string
                      format: uuid
                      example: "ass_1234567890"
                    status:
                      type: string
                      enum: [DRAFT, UNDER_REVIEW, APPROVED, REJECTED]
                      example: "DRAFT"
                    rubric_type:
                      type: string
                      enum: [ANALYTIC, HOLISTIC]
                      example: "ANALYTIC"
                    performance_criteria:
                      type: array
                      items:
                        type: object
                        properties:
                          id:
                            type: string
                            format: uuid
                            example: "pc_1234567890"
                          category:
                            type: string
                            example: "content"
                          description:
                            type: string
                            example: "Accuracy and completeness of content"
                          weight:
                            type: number
                            format: float
                            example: 0.4
                    performance_levels:
                      type: array
                      items:
                        type: object
                        properties:
                          level:
                            type: integer
                            example: 4
                          label:
                            type: string
                            example: "Excellent"
                          description:
                            type: string
                            example: "Demonstrates complete understanding with no errors"
                    scoring_guidelines:
                      type: object
                      properties:
                        total_points:
                          type: integer
                          example: 10
                        passing_score:
                          type: integer
                          example: 6
                    ai_metadata:
                      type: object
                      properties:
                        confidence_score:
                          type: number
                          format: float
                          example: 0.90
                        generated_at:
                          type: string
                          format: date-time
                          example: "2026-06-03T12:00:00Z"
                        agent_version:
                          type: string
                          example: "1.0"
            message:
              type: string
              example: "Rubric generated successfully"
            timestamp:
              type: string
              format: date-time
```

**Error Codes**:
- `GENERATION_FAILED` (500): Failed to generate rubric

**Related Feature**: AI Rubric Generation

**Related Database Tables**:
- `rubric`
- `assessment`
- `ai_audit`
- `workflow_history`

---

## API 5.6: Get Rubric by ID

**Endpoint**: `/api/v1/assessment/rubric/{rubric_id}`

**HTTP Method**: `GET`

**Operation ID**: `getRubric`

**Tags**: `assessment`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER, ADMINISTRATOR

**Security**: `bearerAuth`

**Path Parameters**:

```yaml
parameters:
  - name: rubric_id
    in: path
    required: true
    schema:
      type: string
      format: uuid
    description: Rubric identifier
    example: "rub_1234567890"
```

**Request Schema**: None

**Response Schema** (Success - 200):

```yaml
responses:
  '200':
    description: Rubric retrieved successfully
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: true
            data:
              type: object
              properties:
                rubric:
                  type: object
                  properties:
                    id:
                      type: string
                      format: uuid
                      example: "rub_1234567890"
                    assessment_id:
                      type: string
                      format: uuid
                      example: "ass_1234567890"
                    status:
                      type: string
                      enum: [DRAFT, UNDER_REVIEW, APPROVED, REJECTED]
                      example: "APPROVED"
                    rubric_type:
                      type: string
                      enum: [ANALYTIC, HOLISTIC]
                      example: "ANALYTIC"
                    performance_criteria:
                      type: array
                      items:
                        type: object
                        properties:
                          id:
                            type: string
                            format: uuid
                            example: "pc_1234567890"
                          category:
                            type: string
                            example: "content"
                          description:
                            type: string
                            example: "Accuracy and completeness of content"
                          weight:
                            type: number
                            format: float
                            example: 0.4
                    created_at:
                      type: string
                      format: date-time
                      example: "2026-06-03T10:00:00Z"
                    updated_at:
                      type: string
                      format: date-time
                      example: "2026-06-03T11:00:00Z"
                    approved_at:
                      type: string
                      format: date-time
                      example: "2026-06-03T12:00:00Z"
                    approved_by:
                      type: string
                      format: uuid
                      example: "usr_1234567890"
            message:
              type: string
              example: "Rubric retrieved successfully"
            timestamp:
              type: string
              format: date-time
```

**Error Codes**:
- `RUBRIC_NOT_FOUND` (404): Rubric not found

**Related Feature**: View Rubric

**Related Database Tables**:
- `rubric`

---

## API 5.7: Update Rubric

**Endpoint**: `/api/v1/assessment/rubric/{rubric_id}`

**HTTP Method**: `PUT`

**Operation ID**: `updateRubric`

**Tags**: `assessment`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Security**: `bearerAuth`

**Path Parameters**:

```yaml
parameters:
  - name: rubric_id
    in: path
    required: true
    schema:
      type: string
      format: uuid
    description: Rubric identifier
    example: "rub_1234567890"
```

**Request Schema**:

```yaml
requestBody:
  required: true
  content:
    application/json:
      schema:
        type: object
        required:
          - performance_criteria
        properties:
          performance_criteria:
            type: array
            items:
              type: object
              required:
                - id
                - category
                - description
                - weight
              properties:
                id:
                  type: string
                  format: uuid
                  example: "pc_1234567890"
                category:
                  type: string
                  maxLength: 100
                  example: "content"
                description:
                  type: string
                  maxLength: 500
                  example: "Updated description"
                weight:
                  type: number
                  format: float
                  minimum: 0
                  maximum: 1
                  example: 0.5
```

**Response Schema** (Success - 200):

```yaml
responses:
  '200':
    description: Rubric updated successfully
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: true
            data:
              type: object
              properties:
                rubric:
                  type: object
                  properties:
                    id:
                      type: string
                      format: uuid
                      example: "rub_1234567890"
                    status:
                      type: string
                      enum: [DRAFT, UNDER_REVIEW, APPROVED, REJECTED]
                      example: "DRAFT"
                    updated_at:
                      type: string
                      format: date-time
                      example: "2026-06-03T13:00:00Z"
            message:
              type: string
              example: "Rubric updated successfully"
            timestamp:
              type: string
              format: date-time
```

**Error Codes**: None specific to this endpoint (standard error codes apply)

**Related Feature**: Edit Rubric

**Related Database Tables**:
- `rubric`
- `workflow_history`

---

## API 5.8: Approve Rubric

**Endpoint**: `/api/v1/assessment/rubric/{rubric_id}/approve`

**HTTP Method**: `POST`

**Operation ID**: `approveRubric`

**Tags**: `assessment`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Security**: `bearerAuth`

**Path Parameters**:

```yaml
parameters:
  - name: rubric_id
    in: path
    required: true
    schema:
      type: string
      format: uuid
    description: Rubric identifier
    example: "rub_1234567890"
```

**Request Schema**: None

**Response Schema** (Success - 200):

```yaml
responses:
  '200':
    description: Rubric approved successfully
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: true
            data:
              type: object
              properties:
                rubric:
                  type: object
                  properties:
                    id:
                      type: string
                      format: uuid
                      example: "rub_1234567890"
                    status:
                      type: string
                      enum: [DRAFT, UNDER_REVIEW, APPROVED, REJECTED]
                      example: "APPROVED"
                    approved_at:
                      type: string
                      format: date-time
                      example: "2026-06-03T14:00:00Z"
                    approved_by:
                      type: string
                      format: uuid
                      example: "usr_1234567890"
            message:
              type: string
              example: "Rubric approved successfully"
            timestamp:
              type: string
              format: date-time
```

**Error Codes**: None specific to this endpoint (standard error codes apply)

**Related Feature**: Approve Rubric

**Related Database Tables**:
- `rubric`
- `workflow_history`

---

## API 5.9: Record Evidence

**Endpoint**: `/api/v1/assessment/evidence`

**HTTP Method**: `POST`

**Operation ID**: `recordEvidence`

**Tags**: `assessment`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Security**: `bearerAuth`

**Request Schema**:

```yaml
requestBody:
  required: true
  content:
    application/json:
      schema:
        type: object
        required:
          - student_id
          - assessment_id
          - evidence_type
          - evidence_data
        properties:
          student_id:
            type: string
            format: uuid
            description: Student identifier
            example: "stu_1234567890"
          assessment_id:
            type: string
            format: uuid
            description: Assessment identifier
            example: "ass_1234567890"
          evidence_type:
            type: string
            enum: [STUDENT_WORK, ASSESSMENT_RESULT, OBSERVATION]
            example: "STUDENT_WORK"
          evidence_data:
            type: object
            required:
              - file_url
              - file_name
              - file_size
            properties:
              file_url:
                type: string
                format: uri
                example: "https://storage.nusa.id/evidence/assignment1.pdf"
              file_name:
                type: string
                example: "assignment1.pdf"
              file_size:
                type: integer
                example: 1024000
          teacher_notes:
            type: string
            maxLength: 1000
            example: "Student demonstrated good understanding of concepts"
```

**Response Schema** (Success - 200):

```yaml
responses:
  '200':
    description: Evidence recorded successfully
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: true
            data:
              type: object
              properties:
                evidence:
                  type: object
                  properties:
                    id:
                      type: string
                      format: uuid
                      example: "evi_1234567890"
                    student_id:
                      type: string
                      format: uuid
                      example: "stu_1234567890"
                    assessment_id:
                      type: string
                      format: uuid
                      example: "ass_1234567890"
                    evidence_type:
                      type: string
                      enum: [STUDENT_WORK, ASSESSMENT_RESULT, OBSERVATION]
                      example: "STUDENT_WORK"
                    status:
                      type: string
                      enum: [COLLECTED, LINKED, EVALUATED]
                      example: "COLLECTED"
                    created_at:
                      type: string
                      format: date-time
                      example: "2026-06-03T12:00:00Z"
                    created_by:
                      type: string
                      format: uuid
                      example: "usr_1234567890"
            message:
              type: string
              example: "Evidence recorded successfully"
            timestamp:
              type: string
              format: date-time
```

**Error Codes**: None specific to this endpoint (standard error codes apply)

**Related Feature**: Record Evidence

**Related Database Tables**:
- `evidence`

---

## API 5.10: Link Evidence to Rubric Criteria

**Endpoint**: `/api/v1/assessment/evidence/{evidence_id}/link`

**HTTP Method**: `POST`

**Operation ID**: `linkEvidenceToRubric`

**Tags**: `assessment`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Security**: `bearerAuth`

**Path Parameters**:

```yaml
parameters:
  - name: evidence_id
    in: path
    required: true
    schema:
      type: string
      format: uuid
    description: Evidence identifier
    example: "evi_1234567890"
```

**Request Schema**:

```yaml
requestBody:
  required: true
  content:
    application/json:
      schema:
        type: object
        required:
          - rubric_id
          - criteria_ids
        properties:
          rubric_id:
            type: string
            format: uuid
            description: Rubric identifier
            example: "rub_1234567890"
          criteria_ids:
            type: array
            items:
              type: string
              format: uuid
            example: ["pc_1234567890", "pc_1234567891"]
          evaluation_notes:
            type: string
            maxLength: 1000
            example: "Student performed well in content and organization"
```

**Response Schema** (Success - 200):

```yaml
responses:
  '200':
    description: Evidence linked to rubric criteria successfully
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: true
            data:
              type: object
              properties:
                evidence:
                  type: object
                  properties:
                    id:
                      type: string
                      format: uuid
                      example: "evi_1234567890"
                    status:
                      type: string
                      enum: [COLLECTED, LINKED, EVALUATED]
                      example: "LINKED"
                    rubric_id:
                      type: string
                      format: uuid
                      example: "rub_1234567890"
                    linked_criteria:
                      type: array
                      items:
                        type: string
                        format: uuid
                      example: ["pc_1234567890", "pc_1234567891"]
                    updated_at:
                      type: string
                      format: date-time
                      example: "2026-06-03T13:00:00Z"
            message:
              type: string
              example: "Evidence linked to rubric criteria successfully"
            timestamp:
              type: string
              format: date-time
```

**Error Codes**: None specific to this endpoint (standard error codes apply)

**Related Feature**: Link Evidence to Rubric

**Related Database Tables**:
- `evidence`
- `rubric`

---

## API 5.11: Evaluate Student Performance

**Endpoint**: `/api/v1/assessment/evaluation`

**HTTP Method**: `POST`

**Operation ID**: `evaluateStudentPerformance`

**Tags**: `assessment`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Security**: `bearerAuth`

**Request Schema**:

```yaml
requestBody:
  required: true
  content:
    application/json:
      schema:
        type: object
        required:
          - student_id
          - rubric_id
          - evidence_id
          - performance_scores
        properties:
          student_id:
            type: string
            format: uuid
            description: Student identifier
            example: "stu_1234567890"
          rubric_id:
            type: string
            format: uuid
            description: Rubric identifier
            example: "rub_1234567890"
          evidence_id:
            type: string
            format: uuid
            description: Evidence identifier
            example: "evi_1234567890"
          performance_scores:
            type: array
            items:
              type: object
              required:
                - criteria_id
                - performance_level
                - score
              properties:
                criteria_id:
                  type: string
                  format: uuid
                  example: "pc_1234567890"
                performance_level:
                  type: integer
                  minimum: 1
                  maximum: 5
                  example: 4
                score:
                  type: number
                  minimum: 0
                  maximum: 10
                  example: 4
                notes:
                  type: string
                  maxLength: 500
                  example: "Excellent understanding demonstrated"
```

**Response Schema** (Success - 200):

```yaml
responses:
  '200':
    description: Student performance evaluated successfully
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: true
            data:
              type: object
              properties:
                evaluation:
                  type: object
                  properties:
                    id:
                      type: string
                      format: uuid
                      example: "eval_1234567890"
                    student_id:
                      type: string
                      format: uuid
                      example: "stu_1234567890"
                    rubric_id:
                      type: string
                      format: uuid
                      example: "rub_1234567890"
                    evidence_id:
                      type: string
                      format: uuid
                      example: "evi_1234567890"
                    total_score:
                      type: number
                      example: 7
                    max_score:
                      type: number
                      example: 10
                    performance_level:
                      type: string
                      enum: [BEGINNING, DEVELOPING, PROFICIENT, ADVANCED, EXCELLENT]
                      example: "PROFICIENT"
                    evaluated_at:
                      type: string
                      format: date-time
                      example: "2026-06-03T14:00:00Z"
                    evaluated_by:
                      type: string
                      format: uuid
                      example: "usr_1234567890"
            message:
              type: string
              example: "Student performance evaluated successfully"
            timestamp:
              type: string
              format: date-time
```

**Error Codes**: None specific to this endpoint (standard error codes apply)

**Related Feature**: Evaluate Student Performance

**Related Database Tables**:
- `evaluation`
- `evidence`
- `rubric`

---

# Reporting APIs

## API 6.1: Generate Narrative Report with AI

**Endpoint**: `/api/v1/reporting/narrative-report/generate`

**HTTP Method**: `POST`

**Operation ID**: `generateNarrativeReport`

**Tags**: `reporting`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Security**: `bearerAuth`

**Request Schema**:

```yaml
requestBody:
  required: true
  content:
    application/json:
      schema:
        type: object
        required:
          - student_id
          - evidence_ids
          - evaluation_ids
          - report_period
          - language
        properties:
          student_id:
            type: string
            format: uuid
            description: Student identifier
            example: "stu_1234567890"
          evidence_ids:
            type: array
            items:
              type: string
              format: uuid
            example: ["evi_1234567890", "evi_1234567891"]
          evaluation_ids:
            type: array
            items:
              type: string
              format: uuid
            example: ["eval_1234567890", "eval_1234567891"]
          report_period:
            type: object
            required:
              - type
              - semester
              - academic_year
            properties:
              type:
                type: string
                enum: [SEMESTER, TRIMESTER]
                example: "SEMESTER"
              semester:
                type: integer
                minimum: 1
                maximum: 2
                example: 1
              academic_year:
                type: string
                pattern: '^\d{4}$'
                example: "2026"
          language:
            type: string
            enum: [INDONESIAN, ENGLISH]
            example: "INDONESIAN"
```

**Response Schema** (Success - 200):

```yaml
responses:
  '200':
    description: Narrative report generated successfully
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: true
            data:
              type: object
              properties:
                narrative_report:
                  type: object
                  properties:
                    id:
                      type: string
                      format: uuid
                      example: "nr_1234567890"
                    student_id:
                      type: string
                      format: uuid
                      example: "stu_1234567890"
                    status:
                      type: string
                      enum: [DRAFT, UNDER_REVIEW, APPROVED, REJECTED]
                      example: "DRAFT"
                    report_period:
                      type: object
                      properties:
                        type:
                          type: string
                          enum: [SEMESTER, TRIMESTER]
                          example: "SEMESTER"
                        semester:
                          type: integer
                          example: 1
                        academic_year:
                          type: string
                          example: "2026"
                    content:
                      type: object
                      properties:
                        progress_summary:
                          type: string
                          example: "John has demonstrated consistent progress in Mathematics this semester. He has shown strong understanding of algebraic concepts and applies them effectively in problem-solving."
                        strengths:
                          type: array
                          items:
                            type: string
                          example: ["Strong algebraic thinking skills", "Excellent problem-solving abilities", "Consistent participation in class activities"]
                        areas_for_improvement:
                          type: array
                          items:
                            type: string
                          example: ["Needs more practice with complex equations", "Should improve time management in assessments"]
                        recommendations:
                          type: array
                          items:
                            type: string
                          example: ["Continue practicing algebraic problems regularly", "Focus on time management techniques during assessments", "Consider peer tutoring for complex topics"]
                    ai_metadata:
                      type: object
                      properties:
                        confidence_score:
                          type: number
                          format: float
                          example: 0.87
                        generated_at:
                          type: string
                          format: date-time
                          example: "2026-06-03T12:00:00Z"
                        agent_version:
                          type: string
                          example: "1.0"
            message:
              type: string
              example: "Narrative report generated successfully"
            timestamp:
              type: string
              format: date-time
```

**Error Codes**:
- `GENERATION_FAILED` (500): Failed to generate narrative report

**Related Feature**: AI Narrative Report Generation

**Related Database Tables**:
- `narrative_report`
- `evidence`
- `evaluation`
- `ai_audit`
- `workflow_history`

---

## API 6.2: Get Narrative Report by ID

**Endpoint**: `/api/v1/reporting/narrative-report/{report_id}`

**HTTP Method**: `GET`

**Operation ID**: `getNarrativeReport`

**Tags**: `reporting`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER, ADMINISTRATOR

**Security**: `bearerAuth`

**Path Parameters**:

```yaml
parameters:
  - name: report_id
    in: path
    required: true
    schema:
      type: string
      format: uuid
    description: Narrative Report identifier
    example: "nr_1234567890"
```

**Request Schema**: None

**Response Schema** (Success - 200):

```yaml
responses:
  '200':
    description: Narrative report retrieved successfully
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: true
            data:
              type: object
              properties:
                narrative_report:
                  type: object
                  properties:
                    id:
                      type: string
                      format: uuid
                      example: "nr_1234567890"
                    student_id:
                      type: string
                      format: uuid
                      example: "stu_1234567890"
                    status:
                      type: string
                      enum: [DRAFT, UNDER_REVIEW, APPROVED, REJECTED]
                      example: "APPROVED"
                    report_period:
                      type: object
                      properties:
                        type:
                          type: string
                          enum: [SEMESTER, TRIMESTER]
                          example: "SEMESTER"
                        semester:
                          type: integer
                          example: 1
                        academic_year:
                          type: string
                          example: "2026"
                    content:
                      type: object
                      properties:
                        progress_summary:
                          type: string
                          example: "John has demonstrated consistent progress..."
                        strengths:
                          type: array
                          items:
                            type: string
                          example: ["Strong algebraic thinking skills"]
                        areas_for_improvement:
                          type: array
                          items:
                            type: string
                          example: ["Needs more practice with complex equations"]
                        recommendations:
                          type: array
                          items:
                            type: string
                          example: ["Continue practicing algebraic problems regularly"]
                    created_at:
                      type: string
                      format: date-time
                      example: "2026-06-03T10:00:00Z"
                    updated_at:
                      type: string
                      format: date-time
                      example: "2026-06-03T11:00:00Z"
                    approved_at:
                      type: string
                      format: date-time
                      example: "2026-06-03T12:00:00Z"
                    approved_by:
                      type: string
                      format: uuid
                      example: "usr_1234567890"
            message:
              type: string
              example: "Narrative report retrieved successfully"
            timestamp:
              type: string
              format: date-time
```

**Error Codes**:
- `REPORT_NOT_FOUND` (404): Narrative report not found

**Related Feature**: View Narrative Report

**Related Database Tables**:
- `narrative_report`

---

## API 6.3: Update Narrative Report

**Endpoint**: `/api/v1/reporting/narrative-report/{report_id}`

**HTTP Method**: `PUT`

**Operation ID**: `updateNarrativeReport`

**Tags**: `reporting`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Security**: `bearerAuth`

**Path Parameters**:

```yaml
parameters:
  - name: report_id
    in: path
    required: true
    schema:
      type: string
      format: uuid
    description: Narrative Report identifier
    example: "nr_1234567890"
```

**Request Schema**:

```yaml
requestBody:
  required: true
  content:
    application/json:
      schema:
        type: object
        required:
          - content
        properties:
          content:
            type: object
            required:
              - progress_summary
              - strengths
              - areas_for_improvement
              - recommendations
            properties:
              progress_summary:
                type: string
                maxLength: 2000
                example: "Updated progress summary"
              strengths:
                type: array
                items:
                  type: string
                maxItems: 10
                example: ["Updated strength"]
              areas_for_improvement:
                type: array
                items:
                  type: string
                maxItems: 10
                example: ["Updated area for improvement"]
              recommendations:
                type: array
                items:
                  type: string
                maxItems: 10
                example: ["Updated recommendation"]
```

**Response Schema** (Success - 200):

```yaml
responses:
  '200':
    description: Narrative report updated successfully
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: true
            data:
              type: object
              properties:
                narrative_report:
                  type: object
                  properties:
                    id:
                      type: string
                      format: uuid
                      example: "nr_1234567890"
                    status:
                      type: string
                      enum: [DRAFT, UNDER_REVIEW, APPROVED, REJECTED]
                      example: "DRAFT"
                    updated_at:
                      type: string
                      format: date-time
                      example: "2026-06-03T13:00:00Z"
            message:
              type: string
              example: "Narrative report updated successfully"
            timestamp:
              type: string
              format: date-time
```

**Error Codes**: None specific to this endpoint (standard error codes apply)

**Related Feature**: Edit Narrative Report

**Related Database Tables**:
- `narrative_report`
- `workflow_history`

---

## API 6.4: Approve Narrative Report

**Endpoint**: `/api/v1/reporting/narrative-report/{report_id}/approve`

**HTTP Method**: `POST`

**Operation ID**: `approveNarrativeReport`

**Tags**: `reporting`

**Authentication Requirement**: Required

**Authorization Requirement**: TEACHER

**Security**: `bearerAuth`

**Path Parameters**:

```yaml
parameters:
  - name: report_id
    in: path
    required: true
    schema:
      type: string
      format: uuid
    description: Narrative Report identifier
    example: "nr_1234567890"
```

**Request Schema**: None

**Response Schema** (Success - 200):

```yaml
responses:
  '200':
    description: Narrative report approved successfully
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: true
            data:
              type: object
              properties:
                narrative_report:
                  type: object
                  properties:
                    id:
                      type: string
                      format: uuid
                      example: "nr_1234567890"
                    status:
                      type: string
                      enum: [DRAFT, UNDER_REVIEW, APPROVED, REJECTED]
                      example: "APPROVED"
                    approved_at:
                      type: string
                      format: date-time
                      example: "2026-06-03T14:00:00Z"
                    approved_by:
                      type: string
                      format: uuid
                      example: "usr_1234567890"
            message:
              type: string
              example: "Narrative report approved successfully"
            timestamp:
              type: string
              format: date-time
```

**Error Codes**: None specific to this endpoint (standard error codes apply)

**Related Feature**: Approve Narrative Report

**Related Database Tables**:
- `narrative_report`
- `workflow_history`

---

# Administration APIs

## API 7.1: Create User Account

**Endpoint**: `/api/v1/admin/users`

**HTTP Method**: `POST`

**Operation ID**: `createUser`

**Tags**: `administration`

**Authentication Requirement**: Required

**Authorization Requirement**: ADMINISTRATOR

**Security**: `bearerAuth`

**Request Schema**:

```yaml
requestBody:
  required: true
  content:
    application/json:
      schema:
        type: object
        required:
          - email
          - password
          - name
          - role
        properties:
          email:
            type: string
            format: email
            maxLength: 255
            description: User email address (must be unique)
            example: "teacher@school.id"
          password:
            type: string
            minLength: 8
            maxLength: 128
            description: User password
            example: "securePassword123"
          name:
            type: string
            maxLength: 255
            description: User full name
            example: "John Doe"
          role:
            type: string
            enum: [TEACHER, ADMINISTRATOR]
            description: User role
            example: "TEACHER"
```

**Response Schema** (Success - 201):

```yaml
responses:
  '201':
    description: User account created successfully
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: true
            data:
              type: object
              properties:
                user:
                  type: object
                  properties:
                    id:
                      type: string
                      format: uuid
                      example: "usr_1234567890"
                    email:
                      type: string
                      format: email
                      example: "teacher@school.id"
                    name:
                      type: string
                      example: "John Doe"
                    role:
                      type: string
                      enum: [TEACHER, ADMINISTRATOR]
                      example: "TEACHER"
                    status:
                      type: string
                      enum: [ACTIVE, INACTIVE]
                      example: "ACTIVE"
                    created_at:
                      type: string
                      format: date-time
                      example: "2026-06-03T12:00:00Z"
            message:
              type: string
              example: "User account created successfully"
            timestamp:
              type: string
              format: date-time
```

**Error Codes**:
- `EMAIL_ALREADY_EXISTS` (409): Email address already registered

**Related Feature**: User Management

**Related Database Tables**:
- `users`

---

## API 7.2: Update User Account

**Endpoint**: `/api/v1/admin/users/{user_id}`

**HTTP Method**: `PUT`

**Operation ID**: `updateUser`

**Tags**: `administration`

**Authentication Requirement**: Required

**Authorization Requirement**: ADMINISTRATOR

**Security**: `bearerAuth`

**Path Parameters**:

```yaml
parameters:
  - name: user_id
    in: path
    required: true
    schema:
      type: string
      format: uuid
    description: User identifier
    example: "usr_1234567890"
```

**Request Schema**:

```yaml
requestBody:
  required: true
  content:
    application/json:
      schema:
        type: object
        properties:
          email:
            type: string
            format: email
            maxLength: 255
            description: User email address (must be unique)
            example: "updated@school.id"
          name:
            type: string
            maxLength: 255
            description: User full name
            example: "Updated Name"
          status:
            type: string
            enum: [ACTIVE, INACTIVE]
            description: User account status
            example: "ACTIVE"
```

**Response Schema** (Success - 200):

```yaml
responses:
  '200':
    description: User account updated successfully
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: true
            data:
              type: object
              properties:
                user:
                  type: object
                  properties:
                    id:
                      type: string
                      format: uuid
                      example: "usr_1234567890"
                    email:
                      type: string
                      format: email
                      example: "updated@school.id"
                    name:
                      type: string
                      example: "Updated Name"
                    status:
                      type: string
                      enum: [ACTIVE, INACTIVE]
                      example: "ACTIVE"
                    updated_at:
                      type: string
                      format: date-time
                      example: "2026-06-03T13:00:00Z"
            message:
              type: string
              example: "User account updated successfully"
            timestamp:
              type: string
              format: date-time
```

**Error Codes**:
- `USER_NOT_FOUND` (404): User account not found
- `EMAIL_ALREADY_EXISTS` (409): Email address already registered

**Related Feature**: User Management

**Related Database Tables**:
- `users`

---

## API 7.3: Deactivate User Account

**Endpoint**: `/api/v1/admin/users/{user_id}/deactivate`

**HTTP Method**: `POST`

**Operation ID**: `deactivateUser`

**Tags**: `administration`

**Authentication Requirement**: Required

**Authorization Requirement**: ADMINISTRATOR

**Security**: `bearerAuth`

**Path Parameters**:

```yaml
parameters:
  - name: user_id
    in: path
    required: true
    schema:
      type: string
      format: uuid
    description: User identifier
    example: "usr_1234567890"
```

**Request Schema**: None

**Response Schema** (Success - 200):

```yaml
responses:
  '200':
    description: User account deactivated successfully
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: true
            data:
              type: object
              properties:
                user:
                  type: object
                  properties:
                    id:
                      type: string
                      format: uuid
                      example: "usr_1234567890"
                    status:
                      type: string
                      enum: [ACTIVE, INACTIVE]
                      example: "INACTIVE"
                    deactivated_at:
                      type: string
                      format: date-time
                      example: "2026-06-03T14:00:00Z"
                    deactivated_by:
                      type: string
                      format: uuid
                      example: "usr_1234567891"
            message:
              type: string
              example: "User account deactivated successfully"
            timestamp:
              type: string
              format: date-time
```

**Error Codes**:
- `USER_NOT_FOUND` (404): User account not found

**Related Feature**: User Management

**Related Database Tables**:
- `users`

---

## API 7.4: Assign User Role

**Endpoint**: `/api/v1/admin/users/{user_id}/roles`

**HTTP Method**: `POST`

**Operation ID**: `assignUserRole`

**Tags**: `administration`

**Authentication Requirement**: Required

**Authorization Requirement**: ADMINISTRATOR

**Security**: `bearerAuth`

**Path Parameters**:

```yaml
parameters:
  - name: user_id
    in: path
    required: true
    schema:
      type: string
      format: uuid
    description: User identifier
    example: "usr_1234567890"
```

**Request Schema**:

```yaml
requestBody:
  required: true
  content:
    application/json:
      schema:
        type: object
        required:
          - role
        properties:
          role:
            type: string
            enum: [TEACHER, ADMINISTRATOR]
            description: User role to assign
            example: "ADMINISTRATOR"
```

**Response Schema** (Success - 200):

```yaml
responses:
  '200':
    description: User role assigned successfully
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: true
            data:
              type: object
              properties:
                user:
                  type: object
                  properties:
                    id:
                      type: string
                      format: uuid
                      example: "usr_1234567890"
                    role:
                      type: string
                      enum: [TEACHER, ADMINISTRATOR]
                      example: "ADMINISTRATOR"
                    role_updated_at:
                      type: string
                      format: date-time
                      example: "2026-06-03T14:00:00Z"
                    role_updated_by:
                      type: string
                      format: uuid
                      example: "usr_1234567891"
            message:
              type: string
              example: "User role assigned successfully"
            timestamp:
              type: string
              format: date-time
```

**Error Codes**:
- `USER_NOT_FOUND` (404): User account not found

**Related Feature**: User Management

**Related Database Tables**:
- `users`

---

## API 7.5: List User Accounts

**Endpoint**: `/api/v1/admin/users`

**HTTP Method**: `GET`

**Operation ID**: `listUsers`

**Tags**: `administration`

**Authentication Requirement**: Required

**Authorization Requirement**: ADMINISTRATOR

**Security**: `bearerAuth`

**Query Parameters**:

```yaml
parameters:
  - name: role
    in: query
    required: false
    schema:
      type: string
      enum: [TEACHER, ADMINISTRATOR]
    description: Filter by role
  - name: status
    in: query
    required: false
    schema:
      type: string
      enum: [ACTIVE, INACTIVE]
    description: Filter by status
  - name: page
    in: query
    required: false
    schema:
      type: integer
      minimum: 1
      default: 1
    description: Page number
  - name: limit
    in: query
    required: false
    schema:
      type: integer
      minimum: 1
      maximum: 100
      default: 20
    description: Items per page
```

**Request Schema**: None (query parameters only)

**Response Schema** (Success - 200):

```yaml
responses:
  '200':
    description: User accounts retrieved successfully
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
              example: true
            data:
              type: object
              properties:
                users:
                  type: array
                  items:
                    type: object
                    properties:
                      id:
                        type: string
                        format: uuid
                        example: "usr_1234567890"
                      email:
                        type: string
                        format: email
                        example: "teacher@school.id"
                      name:
                        type: string
                        example: "John Doe"
                      role:
                        type: string
                        enum: [TEACHER, ADMINISTRATOR]
                        example: "TEACHER"
                      status:
                        type: string
                        enum: [ACTIVE, INACTIVE]
                        example: "ACTIVE"
                      created_at:
                        type: string
                        format: date-time
                        example: "2026-06-03T10:00:00Z"
                pagination:
                  type: object
                  properties:
                    total:
                      type: integer
                      example: 50
                    page:
                      type: integer
                      example: 1
                    limit:
                      type: integer
                      example: 20
                    total_pages:
                      type: integer
                      example: 3
            message:
              type: string
              example: "User accounts retrieved successfully"
            timestamp:
              type: string
              format: date-time
```

**Error Codes**: None specific to this endpoint (standard error codes apply)

**Related Feature**: User Management

**Related Database Tables**:
- `users`

---

# Standard Error Codes

## HTTP Status Codes

| HTTP Status | Description |
|-------------|-------------|
| 200 OK | Successful request |
| 201 Created | Resource created successfully |
| 204 No Content | Successful request with no response body |
| 400 Bad Request | Invalid request parameters |
| 401 Unauthorized | Authentication required or failed |
| 403 Forbidden | Authorization failed |
| 404 Not Found | Resource not found |
| 409 Conflict | Resource conflict |
| 422 Unprocessable Entity | Validation failed |
| 500 Internal Server Error | Server error |
| 502 Bad Gateway | Invalid AI provider response |
| 504 Gateway Timeout | AI provider timeout |

## Error Code Catalog

| Error Code | HTTP Status | Description |
|------------|-------------|-------------|
| INVALID_CREDENTIALS | 401 | Invalid email or password |
| INVALID_REFRESH_TOKEN | 401 | Invalid or expired refresh token |
| CP_NOT_FOUND | 404 | Curriculum plan not found |
| GENERATION_FAILED | 500 | AI generation failed |
| TP_NOT_FOUND | 404 | Teaching plan not found |
| TP_UPDATE_FAILED | 500 | Failed to update teaching plan |
| TP_APPROVAL_FAILED | 500 | Failed to approve teaching plan |
| ATP_NOT_FOUND | 404 | Annual teaching plan not found |
| MODUL_AJAR_NOT_FOUND | 404 | Modul Ajar not found |
| ASSESSMENT_NOT_FOUND | 404 | Assessment not found |
| RUBRIC_NOT_FOUND | 404 | Rubric not found |
| REPORT_NOT_FOUND | 404 | Narrative report not found |
| USER_NOT_FOUND | 404 | User account not found |
| EMAIL_ALREADY_EXISTS | 409 | Email address already registered |

---

# OpenAPI Tags

```yaml
tags:
  - name: authentication
    description: Authentication and token management
  - name: curriculum
    description: Curriculum and Teaching Plan management
  - name: learning-planning
    description: Annual Teaching Plan and Modul Ajar management
  - name: assessment
    description: Assessment, Rubric, and Evidence management
  - name: reporting
    description: Narrative Report generation and management
  - name: administration
    description: User and role management
```

---

# OpenAPI Components

## Common Schemas

```yaml
components:
  schemas:
    SuccessResponse:
      type: object
      properties:
        success:
          type: boolean
          example: true
        data:
          type: object
        message:
          type: string
        timestamp:
          type: string
          format: date-time
    
    ErrorResponse:
      type: object
      properties:
        success:
          type: boolean
          example: false
        error:
          type: object
          properties:
            code:
              type: string
            message:
              type: string
            details:
              type: object
        timestamp:
          type: string
          format: date-time
    
    Pagination:
      type: object
      properties:
        total:
          type: integer
        page:
          type: integer
        limit:
          type: integer
        total_pages:
          type: integer
    
    AuditFields:
      type: object
      properties:
        created_by:
          type: string
          format: uuid
        updated_by:
          type: string
          format: uuid
        created_at:
          type: string
          format: date-time
        updated_at:
          type: string
          format: date-time
```

---

# Endpoint Summary

## Total Endpoints

**Total**: 30 endpoints

## Endpoint Distribution

| Module | Endpoints |
|--------|-----------|
| Authentication | 3 |
| Curriculum | 6 |
| Learning Planning | 8 |
| Assessment | 11 |
| Reporting | 4 |
| Administration | 5 |

---

# Implementation Readiness

This OpenAPI preparation document is:
- ✅ Based on the single source of truth: 13_API_CONTRACT.md
- ✅ Implementation-ready for Swagger/OpenAPI generation
- ✅ No new endpoints introduced
- ✅ All endpoints match API Contract exactly
- ✅ Complete request/response schemas
- ✅ Complete error code mappings
- ✅ Complete authentication/authorization requirements
- ✅ Complete feature and database table mappings

**Document Status**: FOUNDATION DOCUMENT - LOCKED

**OpenAPI Generation Status**: READY

**Implementation Start**: TBD

**Implementation End**: TBD

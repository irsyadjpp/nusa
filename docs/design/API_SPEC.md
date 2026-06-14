# API Specification

**Version:** 1.0  
**Date:** June 13, 2026  
**Based On:** API_INVENTORY.md, GAP_ANALYSIS.md, DATABASE_DESIGN.md

---

## Executive Summary

This document specifies the complete API for the NUSA Platform, including existing endpoints from the audit and new endpoints for missing functionality. All endpoints follow RESTful conventions with consistent request/response formats, validation rules, error handling, and authorization rules.

**API Base URL:** `/api/v1`  
**Authentication:** JWT Bearer Token  
**Content-Type:** `application/json`  

---

## Common Specifications

### Request Format

**Headers:**
```
Authorization: Bearer {access_token}
Content-Type: application/json
X-Request-ID: {uuid} (optional)
```

**Query Parameters:**
- `page` (integer, default: 1) - Page number
- `page_size` (integer, default: 20, max: 100) - Items per page
- `sort` (string, format: `{field}:{direction}`) - Sorting
- `filter` (object) - Filter criteria

### Response Format

**Success Response:**
```json
{
  "data": { ... },
  "message": "Success message (optional)",
  "meta": {
    "total": 100,
    "page": 1,
    "page_size": 20,
    "total_pages": 5
  }
}
```

**Error Response:**
```json
{
  "error": "Error message",
  "code": "ERROR_CODE",
  "status": 400,
  "details": { ... }
}
```

### HTTP Status Codes

| Code | Description |
|------|-------------|
| 200 | Success |
| 201 | Created |
| 204 | No Content |
| 400 | Bad Request |
| 401 | Unauthorized |
| 403 | Forbidden |
| 404 | Not Found |
| 409 | Conflict |
| 422 | Unprocessable Entity |
| 429 | Too Many Requests |
| 500 | Internal Server Error |

### Validation Rules

**Common Validation:**
- UUID fields: Must be valid UUID v4
- Email fields: Must be valid email format
- Date fields: Must be ISO 8601 format (YYYY-MM-DD)
- DateTime fields: Must be ISO 8601 format (YYYY-MM-DDTHH:mm:ssZ)
- String fields: Trim whitespace, max length validation
- Integer fields: Min/max value validation
- Boolean fields: Must be true or false

### Error Handling

**Error Codes:**

| Code | Description | HTTP Status |
|------|-------------|-------------|
| AUTH_001 | Invalid credentials | 401 |
| AUTH_002 | Token expired | 401 |
| AUTH_003 | Token invalid | 401 |
| AUTH_004 | Permission denied | 403 |
| VAL_001 | Validation error | 400 |
| VAL_002 | Required field missing | 400 |
| VAL_003 | Invalid format | 400 |
| VAL_004 | Duplicate value | 409 |
| NOT_FOUND_001 | Resource not found | 404 |
| CONFLICT_001 | Resource conflict | 409 |
| RATE_LIMIT_001 | Rate limit exceeded | 429 |
| SERVER_001 | Internal server error | 500 |

### Authorization Rules

**Roles:**
- `SYSTEM_ADMIN` - Full system access
- `SCHOOL_ADMIN` - School-level administration
- `CURRICULUM_ADMIN` - Curriculum management
- `TEACHER` - Teaching and assessment
- `STUDENT` - Student access (future)

**Permissions:**
- Format: `{resource}:{action}`
- Examples: `user:CREATE`, `academic_year:READ`, `class:UPDATE`

**Middleware:**
- `AuthMiddleware` - Validates JWT token
- `RequirePermission` - Checks specific permission
- `RequireRole` - Checks role membership
- `RequireSchoolAccess` - Validates school scope
- `RateLimitMiddleware` - Enforces rate limits

---

## Public Endpoints (No Authentication)

### Health Checks

#### GET /health

**Description:** Health check endpoint

**Request:** None

**Response:**
```json
{
  "status": "healthy",
  "timestamp": "2026-06-13T10:00:00Z"
}
```

**Authorization:** None

---

#### GET /ready

**Description:** Readiness check endpoint

**Request:** None

**Response:**
```json
{
  "status": "ready",
  "checks": {
    "database": "healthy",
    "redis": "healthy",
    "rabbitmq": "healthy"
  }
}
```

**Authorization:** None

---

#### GET /version

**Description:** API version information

**Request:** None

**Response:**
```json
{
  "version": "1.0.0",
  "build": "2024-06-13",
  "environment": "production"
}
```

**Authorization:** None

---

### Authentication

#### POST /public/auth/login

**Description:** User login

**Request:**
```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

**Validation:**
- email: required, valid email format
- password: required, min 8 characters

**Response:**
```json
{
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIs...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIs...",
    "token_type": "Bearer",
    "expires_in": 86400,
    "user": {
      "id": "uuid",
      "email": "user@example.com",
      "name": "John Doe",
      "role_name": "TEACHER",
      "school_id": "uuid"
    }
  }
}
```

**Authorization:** None

**Error Codes:** AUTH_001, VAL_001, VAL_002

---

#### POST /public/auth/refresh

**Description:** Refresh access token

**Request:**
```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIs..."
}
```

**Validation:**
- refresh_token: required, valid JWT

**Response:**
```json
{
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIs...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIs...",
    "token_type": "Bearer",
    "expires_in": 86400
  }
}
```

**Authorization:** None

**Error Codes:** AUTH_002, AUTH_003, VAL_001

---

## Protected Endpoints (Requires Authentication)

### Authentication

#### POST /auth/logout

**Description:** User logout

**Request:**
```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIs..."
}
```

**Validation:**
- refresh_token: required

**Response:**
```json
{
  "message": "Logged out successfully"
}
```

**Authorization:** JWT required

**Error Codes:** AUTH_003

---

#### GET /auth/me

**Description:** Get current user info

**Request:** None

**Response:**
```json
{
  "data": {
    "user": {
      "id": "uuid",
      "email": "user@example.com",
      "name": "John Doe",
      "role_name": "TEACHER",
      "school_id": "uuid"
    },
    "permissions": ["user:READ", "academic_year:READ"]
  }
}
```

**Authorization:** JWT required

**Error Codes:** AUTH_003

---

### Class Management

#### POST /classes

**Description:** Create new class

**Permission:** `class:CREATE`

**Request:**
```json
{
  "school_id": "uuid",
  "academic_year_id": "uuid",
  "semester_id": "uuid",
  "subject_id": "uuid",
  "teacher_id": "uuid",
  "name": "Mathematics Class 10A",
  "grade_level": "10",
  "room": "Room 101",
  "max_students": 40
}
```

**Validation:**
- school_id: required, valid UUID
- academic_year_id: required, valid UUID
- semester_id: required, valid UUID
- subject_id: required, valid UUID
- teacher_id: required, valid UUID
- name: required, max 255 characters
- grade_level: required, max 50 characters
- room: optional, max 100 characters
- max_students: optional, integer, min 1, max 100

**Response:**
```json
{
  "data": {
    "id": "uuid",
    "school_id": "uuid",
    "academic_year_id": "uuid",
    "semester_id": "uuid",
    "subject_id": "uuid",
    "teacher_id": "uuid",
    "name": "Mathematics Class 10A",
    "grade_level": "10",
    "room": "Room 101",
    "max_students": 40,
    "is_active": true,
    "created_at": "2026-06-13T10:00:00Z"
  }
}
```

**Authorization:** JWT + `class:CREATE`

**Error Codes:** VAL_001, VAL_002, VAL_004, AUTH_004

---

#### GET /classes

**Description:** List classes

**Permission:** `class:READ`

**Query Parameters:**
- `school_id` (optional, UUID) - Filter by school
- `academic_year_id` (optional, UUID) - Filter by academic year
- `semester_id` (optional, UUID) - Filter by semester
- `subject_id` (optional, UUID) - Filter by subject
- `teacher_id` (optional, UUID) - Filter by teacher
- `is_active` (optional, boolean) - Filter by active status
- `page` (optional, integer) - Page number
- `page_size` (optional, integer) - Items per page

**Response:**
```json
{
  "data": [
    {
      "id": "uuid",
      "name": "Mathematics Class 10A",
      "grade_level": "10",
      "room": "Room 101",
      "teacher_id": "uuid",
      "teacher_name": "John Doe",
      "subject_id": "uuid",
      "subject_name": "Mathematics",
      "student_count": 35,
      "is_active": true
    }
  ],
  "meta": {
    "total": 100,
    "page": 1,
    "page_size": 20,
    "total_pages": 5
  }
}
```

**Authorization:** JWT + `class:READ`

**Error Codes:** AUTH_004

---

#### GET /classes/:id

**Description:** Get class by ID

**Permission:** `class:READ`

**Response:**
```json
{
  "data": {
    "id": "uuid",
    "school_id": "uuid",
    "academic_year_id": "uuid",
    "semester_id": "uuid",
    "subject_id": "uuid",
    "teacher_id": "uuid",
    "name": "Mathematics Class 10A",
    "grade_level": "10",
    "room": "Room 101",
    "max_students": 40,
    "student_count": 35,
    "is_active": true,
    "created_at": "2026-06-13T10:00:00Z",
    "updated_at": "2026-06-13T10:00:00Z"
  }
}
```

**Authorization:** JWT + `class:READ`

**Error Codes:** NOT_FOUND_001, AUTH_004

---

#### PUT /classes/:id

**Description:** Update class

**Permission:** `class:UPDATE`

**Request:**
```json
{
  "name": "Mathematics Class 10B",
  "room": "Room 102",
  "max_students": 35,
  "is_active": false
}
```

**Validation:**
- name: optional, max 255 characters
- room: optional, max 100 characters
- max_students: optional, integer, min 1, max 100
- is_active: optional, boolean

**Response:**
```json
{
  "data": {
    "id": "uuid",
    "name": "Mathematics Class 10B",
    "room": "Room 102",
    "max_students": 35,
    "is_active": false,
    "updated_at": "2026-06-13T11:00:00Z"
  }
}
```

**Authorization:** JWT + `class:UPDATE`

**Error Codes:** NOT_FOUND_001, VAL_001, AUTH_004

---

#### DELETE /classes/:id

**Description:** Delete class (soft delete)

**Permission:** `class:DELETE`

**Response:**
```json
{
  "message": "Class deleted successfully"
}
```

**Authorization:** JWT + `class:DELETE`

**Error Codes:** NOT_FOUND_001, AUTH_004

---

### Class Enrollments

#### POST /class-enrollments

**Description:** Enroll student in class

**Permission:** `class:UPDATE`

**Request:**
```json
{
  "class_id": "uuid",
  "student_id": "uuid",
  "enrollment_date": "2026-06-13",
  "notes": "Enrolled mid-semester"
}
```

**Validation:**
- class_id: required, valid UUID
- student_id: required, valid UUID
- enrollment_date: optional, valid date format
- notes: optional, text

**Response:**
```json
{
  "data": {
    "id": "uuid",
    "class_id": "uuid",
    "student_id": "uuid",
    "student_name": "Jane Doe",
    "enrollment_date": "2026-06-13",
    "status": "ACTIVE",
    "created_at": "2026-06-13T10:00:00Z"
  }
}
```

**Authorization:** JWT + `class:UPDATE`

**Error Codes:** VAL_001, VAL_002, VAL_004, AUTH_004

---

#### GET /class-enrollments

**Description:** List class enrollments

**Permission:** `class:READ`

**Query Parameters:**
- `class_id` (optional, UUID) - Filter by class
- `student_id` (optional, UUID) - Filter by student
- `status` (optional, string) - Filter by status
- `page` (optional, integer) - Page number
- `page_size` (optional, integer) - Items per page

**Response:**
```json
{
  "data": [
    {
      "id": "uuid",
      "class_id": "uuid",
      "class_name": "Mathematics Class 10A",
      "student_id": "uuid",
      "student_name": "Jane Doe",
      "enrollment_date": "2026-06-13",
      "status": "ACTIVE"
    }
  ],
  "meta": {
    "total": 50,
    "page": 1,
    "page_size": 20,
    "total_pages": 3
  }
}
```

**Authorization:** JWT + `class:READ`

**Error Codes:** AUTH_004

---

#### DELETE /class-enrollments/:id

**Description:** Remove student from class

**Permission:** `class:UPDATE`

**Response:**
```json
{
  "message": "Student removed from class successfully"
}
```

**Authorization:** JWT + `class:UPDATE`

**Error Codes:** NOT_FOUND_001, AUTH_004

---

### Attendance

#### POST /attendance

**Description:** Record attendance

**Permission:** `attendance:CREATE`

**Request:**
```json
{
  "class_id": "uuid",
  "student_id": "uuid",
  "date": "2026-06-13",
  "status": "PRESENT",
  "notes": "On time"
}
```

**Validation:**
- class_id: required, valid UUID
- student_id: required, valid UUID
- date: required, valid date format
- status: required, one of: PRESENT, ABSENT, LATE, EXCUSED
- notes: optional, text

**Response:**
```json
{
  "data": {
    "id": "uuid",
    "class_id": "uuid",
    "student_id": "uuid",
    "student_name": "Jane Doe",
    "date": "2026-06-13",
    "status": "PRESENT",
    "notes": "On time",
    "recorded_by": "uuid",
    "created_at": "2026-06-13T10:00:00Z"
  }
}
```

**Authorization:** JWT + `attendance:CREATE`

**Error Codes:** VAL_001, VAL_002, VAL_004, AUTH_004

---

#### POST /attendance/bulk

**Description:** Bulk record attendance

**Permission:** `attendance:CREATE`

**Request:**
```json
{
  "class_id": "uuid",
  "date": "2026-06-13",
  "records": [
    {
      "student_id": "uuid",
      "status": "PRESENT",
      "notes": "On time"
    },
    {
      "student_id": "uuid",
      "status": "ABSENT",
      "notes": "Sick"
    }
  ]
}
```

**Validation:**
- class_id: required, valid UUID
- date: required, valid date format
- records: required, array, max 100 items
  - student_id: required, valid UUID
  - status: required, one of: PRESENT, ABSENT, LATE, EXCUSED
  - notes: optional, text

**Response:**
```json
{
  "data": {
    "success_count": 35,
    "failed_count": 0,
    "records": [...]
  }
}
```

**Authorization:** JWT + `attendance:CREATE`

**Error Codes:** VAL_001, VAL_002, AUTH_004

---

#### GET /attendance

**Description:** List attendance records

**Permission:** `attendance:READ`

**Query Parameters:**
- `class_id` (optional, UUID) - Filter by class
- `student_id` (optional, UUID) - Filter by student
- `date_from` (optional, date) - Filter by date range (from)
- `date_to` (optional, date) - Filter by date range (to)
- `status` (optional, string) - Filter by status
- `page` (optional, integer) - Page number
- `page_size` (optional, integer) - Items per page

**Response:**
```json
{
  "data": [
    {
      "id": "uuid",
      "class_id": "uuid",
      "class_name": "Mathematics Class 10A",
      "student_id": "uuid",
      "student_name": "Jane Doe",
      "date": "2026-06-13",
      "status": "PRESENT",
      "notes": "On time",
      "recorded_by": "uuid",
      "created_at": "2026-06-13T10:00:00Z"
    }
  ],
  "meta": {
    "total": 500,
    "page": 1,
    "page_size": 20,
    "total_pages": 25
  }
}
```

**Authorization:** JWT + `attendance:READ`

**Error Codes:** AUTH_004

---

#### GET /attendance/report

**Description:** Generate attendance report

**Permission:** `attendance:READ`

**Query Parameters:**
- `class_id` (required, UUID) - Class ID
- `date_from` (required, date) - Report start date
- `date_to` (required, date) - Report end date

**Response:**
```json
{
  "data": {
    "class_id": "uuid",
    "class_name": "Mathematics Class 10A",
    "date_from": "2026-06-01",
    "date_to": "2026-06-30",
    "summary": {
      "total_students": 35,
      "total_days": 20,
      "present_rate": 95.5,
      "absent_rate": 3.2,
      "late_rate": 1.3
    },
    "students": [
      {
        "student_id": "uuid",
        "student_name": "Jane Doe",
        "present_days": 19,
        "absent_days": 1,
        "late_days": 0,
        "attendance_rate": 95.0
      }
    ]
  }
}
```

**Authorization:** JWT + `attendance:READ`

**Error Codes:** VAL_001, VAL_002, AUTH_004

---

#### PUT /attendance/:id

**Description:** Update attendance record

**Permission:** `attendance:UPDATE`

**Request:**
```json
{
  "status": "EXCUSED",
  "notes": "Medical reason"
}
```

**Validation:**
- status: optional, one of: PRESENT, ABSENT, LATE, EXCUSED
- notes: optional, text

**Response:**
```json
{
  "data": {
    "id": "uuid",
    "status": "EXCUSED",
    "notes": "Medical reason",
    "updated_at": "2026-06-13T11:00:00Z"
  }
}
```

**Authorization:** JWT + `attendance:UPDATE`

**Error Codes:** NOT_FOUND_001, VAL_001, AUTH_004

---

### Scheduling

#### POST /schedules

**Description:** Create schedule

**Permission:** `schedule:CREATE`

**Request:**
```json
{
  "class_id": "uuid",
  "day_of_week": 1,
  "start_time": "08:00:00",
  "end_time": "09:30:00",
  "room": "Room 101"
}
```

**Validation:**
- class_id: required, valid UUID
- day_of_week: required, integer, min 1, max 7
- start_time: required, valid time format
- end_time: required, valid time format
- room: optional, max 100 characters

**Response:**
```json
{
  "data": {
    "id": "uuid",
    "class_id": "uuid",
    "day_of_week": 1,
    "start_time": "08:00:00",
    "end_time": "09:30:00",
    "room": "Room 101",
    "is_active": true,
    "created_at": "2026-06-13T10:00:00Z"
  }
}
```

**Authorization:** JWT + `schedule:CREATE`

**Error Codes:** VAL_001, VAL_002, CONFLICT_001, AUTH_004

---

#### GET /schedules

**Description:** List schedules

**Permission:** `schedule:READ`

**Query Parameters:**
- `class_id` (optional, UUID) - Filter by class
- `day_of_week` (optional, integer) - Filter by day
- `room` (optional, string) - Filter by room
- `is_active` (optional, boolean) - Filter by active status
- `page` (optional, integer) - Page number
- `page_size` (optional, integer) - Items per page

**Response:**
```json
{
  "data": [
    {
      "id": "uuid",
      "class_id": "uuid",
      "class_name": "Mathematics Class 10A",
      "day_of_week": 1,
      "start_time": "08:00:00",
      "end_time": "09:30:00",
      "room": "Room 101",
      "is_active": true
    }
  ],
  "meta": {
    "total": 50,
    "page": 1,
    "page_size": 20,
    "total_pages": 3
  }
}
```

**Authorization:** JWT + `schedule:READ`

**Error Codes:** AUTH_004

---

#### GET /schedules/conflicts

**Description:** Check for schedule conflicts

**Permission:** `schedule:READ`

**Query Parameters:**
- `teacher_id` (optional, UUID) - Check teacher conflicts
- `room` (optional, string) - Check room conflicts
- `day_of_week` (optional, integer) - Filter by day
- `start_time` (optional, time) - Start time
- `end_time` (optional, time) - End time

**Response:**
```json
{
  "data": {
    "has_conflicts": true,
    "conflicts": [
      {
        "type": "TEACHER",
        "teacher_id": "uuid",
        "teacher_name": "John Doe",
        "existing_schedule": {
          "class_id": "uuid",
          "day_of_week": 1,
          "start_time": "08:00:00",
          "end_time": "09:30:00"
        },
        "proposed_schedule": {
          "day_of_week": 1,
          "start_time": "09:00:00",
          "end_time": "10:30:00"
        }
      }
    ]
  }
}
```

**Authorization:** JWT + `schedule:READ`

**Error Codes:** VAL_001, AUTH_004

---

#### PUT /schedules/:id

**Description:** Update schedule

**Permission:** `schedule:UPDATE`

**Request:**
```json
{
  "day_of_week": 2,
  "start_time": "09:00:00",
  "end_time": "10:30:00",
  "room": "Room 102",
  "is_active": false
}
```

**Validation:**
- day_of_week: optional, integer, min 1, max 7
- start_time: optional, valid time format
- end_time: optional, valid time format
- room: optional, max 100 characters
- is_active: optional, boolean

**Response:**
```json
{
  "data": {
    "id": "uuid",
    "day_of_week": 2,
    "start_time": "09:00:00",
    "end_time": "10:30:00",
    "room": "Room 102",
    "is_active": false,
    "updated_at": "2026-06-13T11:00:00Z"
  }
}
```

**Authorization:** JWT + `schedule:UPDATE`

**Error Codes:** NOT_FOUND_001, VAL_001, CONFLICT_001, AUTH_004

---

#### DELETE /schedules/:id

**Description:** Delete schedule

**Permission:** `schedule:DELETE`

**Response:**
```json
{
  "message": "Schedule deleted successfully"
}
```

**Authorization:** JWT + `schedule:DELETE`

**Error Codes:** NOT_FOUND_001, AUTH_004

---

### Exams

#### POST /exams

**Description:** Create exam

**Permission:** `exam:CREATE`

**Request:**
```json
{
  "class_id": "uuid",
  "assessment_id": "uuid",
  "exam_date": "2026-06-20",
  "start_time": "09:00:00",
  "duration_minutes": 90,
  "room": "Room 101"
}
```

**Validation:**
- class_id: required, valid UUID
- assessment_id: required, valid UUID
- exam_date: required, valid date format
- start_time: required, valid time format
- duration_minutes: required, integer, min 1
- room: optional, max 100 characters

**Response:**
```json
{
  "data": {
    "id": "uuid",
    "class_id": "uuid",
    "assessment_id": "uuid",
    "exam_date": "2026-06-20",
    "start_time": "09:00:00",
    "duration_minutes": 90,
    "room": "Room 101",
    "status": "SCHEDULED",
    "created_at": "2026-06-13T10:00:00Z"
  }
}
```

**Authorization:** JWT + `exam:CREATE`

**Error Codes:** VAL_001, VAL_002, CONFLICT_001, AUTH_004

---

#### GET /exams

**Description:** List exams

**Permission:** `exam:READ`

**Query Parameters:**
- `class_id` (optional, UUID) - Filter by class
- `assessment_id` (optional, UUID) - Filter by assessment
- `exam_date_from` (optional, date) - Filter by date range (from)
- `exam_date_to` (optional, date) - Filter by date range (to)
- `status` (optional, string) - Filter by status
- `page` (optional, integer) - Page number
- `page_size` (optional, integer) - Items per page

**Response:**
```json
{
  "data": [
    {
      "id": "uuid",
      "class_id": "uuid",
      "class_name": "Mathematics Class 10A",
      "assessment_id": "uuid",
      "assessment_title": "Midterm Exam",
      "exam_date": "2026-06-20",
      "start_time": "09:00:00",
      "duration_minutes": 90,
      "room": "Room 101",
      "status": "SCHEDULED"
    }
  ],
  "meta": {
    "total": 20,
    "page": 1,
    "page_size": 20,
    "total_pages": 1
  }
}
```

**Authorization:** JWT + `exam:READ`

**Error Codes:** AUTH_004

---

#### PUT /exams/:id

**Description:** Update exam

**Permission:** `exam:UPDATE`

**Request:**
```json
{
  "exam_date": "2026-06-21",
  "start_time": "10:00:00",
  "duration_minutes": 120,
  "room": "Room 102",
  "status": "CANCELLED"
}
```

**Validation:**
- exam_date: optional, valid date format
- start_time: optional, valid time format
- duration_minutes: optional, integer, min 1
- room: optional, max 100 characters
- status: optional, one of: SCHEDULED, IN_PROGRESS, COMPLETED, CANCELLED

**Response:**
```json
{
  "data": {
    "id": "uuid",
    "exam_date": "2026-06-21",
    "start_time": "10:00:00",
    "duration_minutes": 120,
    "room": "Room 102",
    "status": "CANCELLED",
    "updated_at": "2026-06-13T11:00:00Z"
  }
}
```

**Authorization:** JWT + `exam:UPDATE`

**Error Codes:** NOT_FOUND_001, VAL_001, CONFLICT_001, AUTH_004

---

### Assignments

#### POST /assignments

**Description:** Create assignment

**Permission:** `assignment:CREATE`

**Request:**
```json
{
  "class_id": "uuid",
  "assessment_id": "uuid",
  "title": "Homework Chapter 5",
  "description": "Complete exercises 1-10",
  "due_date": "2026-06-20T23:59:59Z",
  "max_score": 100
}
```

**Validation:**
- class_id: required, valid UUID
- assessment_id: required, valid UUID
- title: required, max 255 characters
- description: optional, text
- due_date: required, valid datetime format
- max_score: optional, integer, min 0

**Response:**
```json
{
  "data": {
    "id": "uuid",
    "class_id": "uuid",
    "assessment_id": "uuid",
    "title": "Homework Chapter 5",
    "description": "Complete exercises 1-10",
    "due_date": "2026-06-20T23:59:59Z",
    "max_score": 100,
    "status": "ASSIGNED",
    "created_at": "2026-06-13T10:00:00Z"
  }
}
```

**Authorization:** JWT + `assignment:CREATE`

**Error Codes:** VAL_001, VAL_002, AUTH_004

---

#### GET /assignments

**Description:** List assignments

**Permission:** `assignment:READ`

**Query Parameters:**
- `class_id` (optional, UUID) - Filter by class
- `assessment_id` (optional, UUID) - Filter by assessment
- `due_date_from` (optional, datetime) - Filter by due date range (from)
- `due_date_to` (optional, datetime) - Filter by due date range (to)
- `status` (optional, string) - Filter by status
- `page` (optional, integer) - Page number
- `page_size` (optional, integer) - Items per page

**Response:**
```json
{
  "data": [
    {
      "id": "uuid",
      "class_id": "uuid",
      "class_name": "Mathematics Class 10A",
      "assessment_id": "uuid",
      "title": "Homework Chapter 5",
      "description": "Complete exercises 1-10",
      "due_date": "2026-06-20T23:59:59Z",
      "max_score": 100,
      "status": "ASSIGNED"
    }
  ],
  "meta": {
    "total": 30,
    "page": 1,
    "page_size": 20,
    "total_pages": 2
  }
}
```

**Authorization:** JWT + `assignment:READ`

**Error Codes:** AUTH_004

---

#### PUT /assignments/:id

**Description:** Update assignment

**Permission:** `assignment:UPDATE`

**Request:**
```json
{
  "title": "Homework Chapter 5 (Updated)",
  "due_date": "2026-06-21T23:59:59Z",
  "max_score": 100,
  "status": "CANCELLED"
}
```

**Validation:**
- title: optional, max 255 characters
- description: optional, text
- due_date: optional, valid datetime format
- max_score: optional, integer, min 0
- status: optional, one of: ASSIGNED, IN_PROGRESS, SUBMITTED, GRADED, CANCELLED

**Response:**
```json
{
  "data": {
    "id": "uuid",
    "title": "Homework Chapter 5 (Updated)",
    "due_date": "2026-06-21T23:59:59Z",
    "max_score": 100,
    "status": "CANCELLED",
    "updated_at": "2026-06-13T11:00:00Z"
  }
}
```

**Authorization:** JWT + `assignment:UPDATE`

**Error Codes:** NOT_FOUND_001, VAL_001, AUTH_004

---

### Exam Results

#### POST /exam-results

**Description:** Record exam result

**Permission:** `exam:UPDATE`

**Request:**
```json
{
  "exam_id": "uuid",
  "student_id": "uuid",
  "score": 85.5,
  "grade": "A",
  "remarks": "Excellent performance"
}
```

**Validation:**
- exam_id: required, valid UUID
- student_id: required, valid UUID
- score: optional, decimal, min 0, max 100
- grade: optional, max 10 characters
- remarks: optional, text

**Response:**
```json
{
  "data": {
    "id": "uuid",
    "exam_id": "uuid",
    "student_id": "uuid",
    "student_name": "Jane Doe",
    "score": 85.5,
    "grade": "A",
    "remarks": "Excellent performance",
    "graded_at": "2026-06-13T10:00:00Z",
    "graded_by": "uuid"
  }
}
```

**Authorization:** JWT + `exam:UPDATE`

**Error Codes:** VAL_001, VAL_002, VAL_004, AUTH_004

---

#### GET /exam-results

**Description:** List exam results

**Permission:** `exam:READ`

**Query Parameters:**
- `exam_id` (optional, UUID) - Filter by exam
- `student_id` (optional, UUID) - Filter by student
- `grade` (optional, string) - Filter by grade
- `page` (optional, integer) - Page number
- `page_size` (optional, integer) - Items per page

**Response:**
```json
{
  "data": [
    {
      "id": "uuid",
      "exam_id": "uuid",
      "exam_title": "Midterm Exam",
      "student_id": "uuid",
      "student_name": "Jane Doe",
      "score": 85.5,
      "grade": "A",
      "remarks": "Excellent performance",
      "graded_at": "2026-06-13T10:00:00Z"
    }
  ],
  "meta": {
    "total": 35,
    "page": 1,
    "page_size": 20,
    "total_pages": 2
  }
}
```

**Authorization:** JWT + `exam:READ`

**Error Codes:** AUTH_004

---

### Notifications

#### POST /notifications

**Description:** Create notification

**Permission:** `notification:CREATE`

**Request:**
```json
{
  "user_id": "uuid",
  "title": "New Assignment",
  "message": "You have a new assignment due tomorrow",
  "type": "INFO",
  "action_url": "/dashboard/assignments/123"
}
```

**Validation:**
- user_id: required, valid UUID
- title: required, max 255 characters
- message: required, text
- type: required, one of: INFO, WARNING, ERROR, SUCCESS
- action_url: optional, max 500 characters
- metadata: optional, JSON object

**Response:**
```json
{
  "data": {
    "id": "uuid",
    "user_id": "uuid",
    "title": "New Assignment",
    "message": "You have a new assignment due tomorrow",
    "type": "INFO",
    "is_read": false,
    "action_url": "/dashboard/assignments/123",
    "created_at": "2026-06-13T10:00:00Z"
  }
}
```

**Authorization:** JWT + `notification:CREATE`

**Error Codes:** VAL_001, VAL_002, AUTH_004

---

#### GET /notifications

**Description:** List notifications

**Permission:** `notification:READ`

**Query Parameters:**
- `user_id` (optional, UUID) - Filter by user (defaults to current user)
- `type` (optional, string) - Filter by type
- `is_read` (optional, boolean) - Filter by read status
- `page` (optional, integer) - Page number
- `page_size` (optional, integer) - Items per page

**Response:**
```json
{
  "data": [
    {
      "id": "uuid",
      "title": "New Assignment",
      "message": "You have a new assignment due tomorrow",
      "type": "INFO",
      "is_read": false,
      "action_url": "/dashboard/assignments/123",
      "created_at": "2026-06-13T10:00:00Z"
    }
  ],
  "meta": {
    "total": 15,
    "page": 1,
    "page_size": 20,
    "total_pages": 1
  }
}
```

**Authorization:** JWT + `notification:READ`

**Error Codes:** AUTH_004

---

#### PUT /notifications/:id/read

**Description:** Mark notification as read

**Permission:** `notification:UPDATE`

**Response:**
```json
{
  "data": {
    "id": "uuid",
    "is_read": true,
    "read_at": "2026-06-13T11:00:00Z"
  }
}
```

**Authorization:** JWT + `notification:UPDATE`

**Error Codes:** NOT_FOUND_001, AUTH_004

---

### Announcements

#### POST /announcements

**Description:** Create announcement

**Permission:** `announcement:CREATE`

**Request:**
```json
{
  "school_id": "uuid",
  "title": "School Holiday",
  "content": "School will be closed on June 15 for Independence Day",
  "priority": "HIGH",
  "target_audience": "ALL",
  "expires_at": "2026-06-20T23:59:59Z"
}
```

**Validation:**
- school_id: required, valid UUID
- title: required, max 255 characters
- content: required, text
- priority: optional, one of: LOW, NORMAL, HIGH, URGENT
- target_audience: optional, one of: ALL, TEACHERS, STUDENTS, PARENTS, ADMIN
- expires_at: optional, valid datetime format

**Response:**
```json
{
  "data": {
    "id": "uuid",
    "school_id": "uuid",
    "title": "School Holiday",
    "content": "School will be closed on June 15 for Independence Day",
    "priority": "HIGH",
    "target_audience": "ALL",
    "published_by": "uuid",
    "published_at": "2026-06-13T10:00:00Z",
    "expires_at": "2026-06-20T23:59:59Z",
    "is_active": true
  }
}
```

**Authorization:** JWT + `announcement:CREATE`

**Error Codes:** VAL_001, VAL_002, AUTH_004

---

#### GET /announcements

**Description:** List announcements

**Permission:** `announcement:READ`

**Query Parameters:**
- `school_id` (optional, UUID) - Filter by school
- `priority` (optional, string) - Filter by priority
- `target_audience` (optional, string) - Filter by target audience
- `is_active` (optional, boolean) - Filter by active status
- `page` (optional, integer) - Page number
- `page_size` (optional, integer) - Items per page

**Response:**
```json
{
  "data": [
    {
      "id": "uuid",
      "school_id": "uuid",
      "title": "School Holiday",
      "content": "School will be closed on June 15 for Independence Day",
      "priority": "HIGH",
      "target_audience": "ALL",
      "published_by": "uuid",
      "published_by_name": "John Doe",
      "published_at": "2026-06-13T10:00:00Z",
      "expires_at": "2026-06-20T23:59:59Z",
      "is_active": true
    }
  ],
  "meta": {
    "total": 10,
    "page": 1,
    "page_size": 20,
    "total_pages": 1
  }
}
```

**Authorization:** JWT + `announcement:READ`

**Error Codes:** AUTH_004

---

### Messages

#### POST /messages

**Description:** Send message

**Permission:** `message:CREATE`

**Request:**
```json
{
  "receiver_id": "uuid",
  "subject": "Question about assignment",
  "content": "Can you clarify question 5?"
}
```

**Validation:**
- receiver_id: required, valid UUID
- subject: optional, max 255 characters
- content: required, text
- parent_message_id: optional, valid UUID (for replies)

**Response:**
```json
{
  "data": {
    "id": "uuid",
    "sender_id": "uuid",
    "sender_name": "Jane Doe",
    "receiver_id": "uuid",
    "receiver_name": "John Doe",
    "subject": "Question about assignment",
    "content": "Can you clarify question 5?",
    "is_read": false,
    "created_at": "2026-06-13T10:00:00Z"
  }
}
```

**Authorization:** JWT + `message:CREATE`

**Error Codes:** VAL_001, VAL_002, AUTH_004

---

#### GET /messages

**Description:** List messages

**Permission:** `message:READ`

**Query Parameters:**
- `user_id` (optional, UUID) - Filter by user (defaults to current user)
- `is_read` (optional, boolean) - Filter by read status
- `page` (optional, integer) - Page number
- `page_size` (optional, integer) - Items per page

**Response:**
```json
{
  "data": [
    {
      "id": "uuid",
      "sender_id": "uuid",
      "sender_name": "Jane Doe",
      "receiver_id": "uuid",
      "receiver_name": "John Doe",
      "subject": "Question about assignment",
      "content": "Can you clarify question 5?",
      "is_read": false,
      "created_at": "2026-06-13T10:00:00Z"
    }
  ],
  "meta": {
    "total": 25,
    "page": 1,
    "page_size": 20,
    "total_pages": 2
  }
}
```

**Authorization:** JWT + `message:READ`

**Error Codes:** AUTH_004

---

#### PUT /messages/:id/read

**Description:** Mark message as read

**Permission:** `message:UPDATE`

**Response:**
```json
{
  "data": {
    "id": "uuid",
    "is_read": true,
    "read_at": "2026-06-13T11:00:00Z"
  }
}
```

**Authorization:** JWT + `message:UPDATE`

**Error Codes:** NOT_FOUND_001, AUTH_004

---

## Rate Limiting

### Rate Limit Rules

**Per User:**
- Authentication endpoints: 10 requests/minute
- Read endpoints: 100 requests/minute
- Write endpoints: 30 requests/minute

**Per IP:**
- All endpoints: 200 requests/minute

**Rate Limit Headers:**
```
X-RateLimit-Limit: 100
X-RateLimit-Remaining: 95
X-RateLimit-Reset: 1623542400
```

**Rate Limit Response (429):**
```json
{
  "error": "Rate limit exceeded",
  "code": "RATE_LIMIT_001",
  "status": 429,
  "retry_after": 60
}
```

---

## Pagination

### Default Pagination

- Default page: 1
- Default page_size: 20
- Maximum page_size: 100

### Pagination Response

```json
{
  "data": [...],
  "meta": {
    "total": 100,
    "page": 1,
    "page_size": 20,
    "total_pages": 5
  }
}
```

---

## Sorting

### Sorting Format

```
sort={field}:{direction}
```

**Examples:**
- `sort=name:asc` - Sort by name ascending
- `sort=created_at:desc` - Sort by created_at descending

**Supported Directions:**
- `asc` - Ascending
- `desc` - Descending

---

## Filtering

### Filter Format

Filter as query parameters or JSON object in request body.

**Examples:**
- `?status=ACTIVE`
- `?school_id=uuid&status=ACTIVE`
- `?date_from=2026-06-01&date_to=2026-06-30`

---

## Conclusion

This API specification provides comprehensive coverage of all endpoints for the NUSA Platform, including existing endpoints from the audit and new endpoints for class management, attendance, scheduling, exams, assignments, and communication features. All endpoints follow consistent patterns for request/response formats, validation rules, error handling, and authorization rules.

The specification prioritizes maintainability, scalability, simplicity, testability, security, and observability while avoiding overengineering.

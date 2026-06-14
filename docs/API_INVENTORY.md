# API Inventory

**Generated:** June 13, 2026  
**Scope:** NUSA Platform REST API  
**Base URL:** `/api/v1`  
**Version:** 1.0

---

## Executive Summary

The NUSA Platform provides a RESTful API built with Gin Gonic framework. The API follows a layered architecture with JWT authentication, role-based access control, and permission-based authorization. All endpoints return JSON responses with consistent error handling.

**Key Metrics:**
- **Total Endpoints:** 80+
- **Authentication:** JWT Bearer tokens with refresh token rotation
- **Authorization:** Role-based (SYSTEM_ADMIN, SCHOOL_ADMIN, CURRICULUM_ADMIN, TEACHER)
- **Documentation:** Scalar API documentation available at `/scalar`
- **Status:** Production-ready

---

## Authentication & Authorization

### Authentication Flow

1. **Login:** POST `/api/v1/public/auth/login`
   - Request: `{ email, password }`
   - Response: `{ access_token, refresh_token, token_type, expires_in, user }`

2. **Refresh Token:** POST `/api/v1/public/auth/refresh`
   - Request: `{ refresh_token }`
   - Response: `{ access_token, refresh_token, token_type, expires_in, user }`

3. **Logout:** POST `/api/v1/auth/logout`
   - Request: `{ refresh_token }`
   - Response: `{ message }`
   - Requires: JWT authentication

4. **Get Current User:** GET `/api/v1/auth/me`
   - Response: `{ user, role_name, permissions }`
   - Requires: JWT authentication

### Authorization Model

**Roles:**
- `SYSTEM_ADMIN` - Full system access
- `SCHOOL_ADMIN` - School-level administration
- `CURRICULUM_ADMIN` - Curriculum management
- `TEACHER` - Teaching and assessment

**Permissions:** Resource-action format (e.g., `user:CREATE`, `academic_year:READ`)

**Middleware:**
- `AuthMiddleware` - Validates JWT tokens
- `RequirePermission` - Checks specific permissions
- `RequireRole` - Checks role membership
- `RequireSchoolAccess` - Validates school scope
- `ReadOnlyMiddleware` - Restricts teachers to GET requests

---

## Public Endpoints (No Authentication)

### Health Checks

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check endpoint |
| GET | `/ready` | Readiness check endpoint |
| GET | `/version` | API version information |

### Authentication

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/v1/public/auth/login` | User login |
| POST | `/api/v1/public/auth/refresh` | Refresh access token |

---

## Protected Endpoints (Requires Authentication)

### Authentication

| Method | Endpoint | Permission | Description |
|--------|----------|------------|-------------|
| POST | `/api/v1/auth/logout` | - | User logout |
| GET | `/api/v1/auth/me` | - | Get current user info |

---

### User Management

| Method | Endpoint | Permission | Description |
|--------|----------|------------|-------------|
| POST | `/api/v1/users` | `user:CREATE` | Create new user |
| GET | `/api/v1/users` | `user:READ` | List users (with filters) |
| GET | `/api/v1/users/:id` | `user:READ` | Get user by ID |
| PUT | `/api/v1/users/:id` | `user:UPDATE` | Update user |
| PATCH | `/api/v1/users/:id/status` | `user:UPDATE` | Update user status |

**Query Parameters (GET /users):**
- `school_id` (optional) - Filter by school
- `role_id` (optional) - Filter by role
- `is_active` (optional) - Filter by active status
- `page` (default: 1) - Page number
- `page_size` (default: 20) - Items per page

---

### School Management

| Method | Endpoint | Permission | Description |
|--------|----------|------------|-------------|
| POST | `/api/v1/schools` | `school:CREATE` | Create new school |
| GET | `/api/v1/schools` | `school:READ` | List schools |
| GET | `/api/v1/schools/:id` | `school:READ` | Get school by ID |
| PUT | `/api/v1/schools/:id` | `school:UPDATE` | Update school |
| PATCH | `/api/v1/schools/:id/status` | `school:UPDATE` | Update school status |

---

### Role Management

| Method | Endpoint | Permission | Description |
|--------|----------|------------|-------------|
| GET | `/api/v1/roles` | `user:READ` | List roles |
| GET | `/api/v1/roles/:id` | `user:READ` | Get role by ID |
| GET | `/api/v1/roles/:id/permissions` | `user:READ` | Get role permissions |
| POST | `/api/v1/roles` | `user:CREATE` + SYSTEM_ADMIN | Create new role |
| DELETE | `/api/v1/roles/:id` | `user:DELETE` + SYSTEM_ADMIN | Delete role |
| PUT | `/api/v1/roles/:id` | `user:UPDATE` + SYSTEM_ADMIN | Update role |
| POST | `/api/v1/roles/:id/permissions` | `user:UPDATE` | Add permission to role |
| DELETE | `/api/v1/roles/:id/permissions` | `user:UPDATE` | Remove permission from role |

---

### Curriculum (CP - Capaian Pembelajaran)

**Middleware:** ReadOnlyMiddleware (SYSTEM_ADMIN, SCHOOL_ADMIN have full access; TEACHER has read-only)

#### Subjects

| Method | Endpoint | Permission | Description |
|--------|----------|------------|-------------|
| POST | `/api/v1/curriculum/subjects` | - | Create curriculum subject |
| GET | `/api/v1/curriculum/subjects` | - | List curriculum subjects |
| GET | `/api/v1/curriculum/subjects/:id` | - | Get subject by ID |
| PUT | `/api/v1/curriculum/subjects/:id` | - | Update subject |
| DELETE | `/api/v1/curriculum/subjects/:id` | - | Delete subject |

#### Phases

| Method | Endpoint | Permission | Description |
|--------|----------|------------|-------------|
| POST | `/api/v1/curriculum/phases` | - | Create curriculum phase |
| GET | `/api/v1/curriculum/phases` | - | List curriculum phases |
| GET | `/api/v1/curriculum/phases/:id` | - | Get phase by ID |
| PUT | `/api/v1/curriculum/phases/:id` | - | Update phase |
| DELETE | `/api/v1/curriculum/phases/:id` | - | Delete phase |

#### Elements

| Method | Endpoint | Permission | Description |
|--------|----------|------------|-------------|
| POST | `/api/v1/curriculum/elements` | - | Create curriculum element |
| GET | `/api/v1/curriculum/elements` | - | List curriculum elements |
| GET | `/api/v1/curriculum/elements/:id` | - | Get element by ID |
| PUT | `/api/v1/curriculum/elements/:id` | - | Update element |
| DELETE | `/api/v1/curriculum/elements/:id` | - | Delete element |

#### Subelements

| Method | Endpoint | Permission | Description |
|--------|----------|------------|-------------|
| POST | `/api/v1/curriculum/subelements` | - | Create curriculum subelement |
| GET | `/api/v1/curriculum/subelements` | - | List curriculum subelements |
| GET | `/api/v1/curriculum/subelements/:id` | - | Get subelement by ID |
| PUT | `/api/v1/curriculum/subelements/:id` | - | Update subelement |
| DELETE | `/api/v1/curriculum/subelements/:id` | - | Delete subelement |

#### CP (Capaian Pembelajaran)

| Method | Endpoint | Permission | Description |
|--------|----------|------------|-------------|
| POST | `/api/v1/curriculum/cp` | - | Create CP |
| PUT | `/api/v1/curriculum/cp/:id` | - | Update CP |
| DELETE | `/api/v1/curriculum/cp/:id` | - | Delete CP |
| GET | `/api/v1/curriculum/cp/export` | - | Export CPs as CSV |
| POST | `/api/v1/curriculum/cp/import` | - | Bulk import CPs |
| GET | `/api/v1/curriculum/cp` | - | List CPs |
| GET | `/api/v1/curriculum/cp/:id` | - | Get CP by ID |

**Query Parameters (GET /curriculum/cp):**
- `subject_id` (optional) - Filter by subject
- `phase_id` (optional) - Filter by phase
- `element_id` (optional) - Filter by element
- `page` (default: 1) - Page number
- `page_size` (default: 20) - Items per page

---

### Learning Planning (TP, ATP, Modul Ajar)

**Middleware:** Requires `tp:READ` permission

#### TP Sets (Teaching Plan Sets)

| Method | Endpoint | Permission | Description |
|--------|----------|------------|-------------|
| POST | `/api/v1/learning-planning/tp-sets` | `tp:CREATE` | Create TP set |
| GET | `/api/v1/learning-planning/tp-sets` | - | List TP sets |
| GET | `/api/v1/learning-planning/tp-sets/:id` | - | Get TP set by ID |
| POST | `/api/v1/learning-planning/tp-sets/:id/approve` | `tp:APPROVE` | Approve TP set |
| PUT | `/api/v1/learning-planning/tp-sets/:id` | `tp:UPDATE` | Update TP set |
| GET | `/api/v1/learning-planning/tp-sets/:id/versions` | - | Get TP set versions |

#### Individual TPs

| Method | Endpoint | Permission | Description |
|--------|----------|------------|-------------|
| POST | `/api/v1/learning-planning/tps` | `tp:CREATE` | Create TP |
| GET | `/api/v1/learning-planning/tps` | - | List TPs |
| GET | `/api/v1/learning-planning/tps/:id` | - | Get TP by ID |

#### ATP Sets (Alur Tujuan Pembelajaran Sets)

| Method | Endpoint | Permission | Description |
|--------|----------|------------|-------------|
| POST | `/api/v1/learning-planning/atp-sets` | `tp:CREATE` | Create ATP set |
| GET | `/api/v1/learning-planning/atp-sets` | - | List ATP sets |
| GET | `/api/v1/learning-planning/atp-sets/:id` | - | Get ATP set by ID |
| PUT | `/api/v1/learning-planning/atp-sets/:id` | `tp:UPDATE` | Update ATP set |
| DELETE | `/api/v1/learning-planning/atp-sets/:id` | `tp:DELETE` | Delete ATP set |
| POST | `/api/v1/learning-planning/atp-sets/:id/approve` | `tp:APPROVE` | Approve ATP set |

#### Individual ATPs

| Method | Endpoint | Permission | Description |
|--------|----------|------------|-------------|
| POST | `/api/v1/learning-planning/atps` | `tp:CREATE` | Create ATP |
| GET | `/api/v1/learning-planning/atps` | - | List ATPs |
| GET | `/api/v1/learning-planning/atps/:id` | - | Get ATP by ID |
| PUT | `/api/v1/learning-planning/atps/:id` | `tp:UPDATE` | Update ATP |
| DELETE | `/api/v1/learning-planning/atps/:id` | `tp:DELETE` | Delete ATP |

#### Modul Ajar Sets

| Method | Endpoint | Permission | Description |
|--------|----------|------------|-------------|
| POST | `/api/v1/learning-planning/modul-ajar-sets` | `tp:CREATE` | Create Modul Ajar set |
| GET | `/api/v1/learning-planning/modul-ajar-sets` | - | List Modul Ajar sets |
| GET | `/api/v1/learning-planning/modul-ajar-sets/:id` | - | Get Modul Ajar set by ID |
| PUT | `/api/v1/learning-planning/modul-ajar-sets/:id` | `tp:UPDATE` | Update Modul Ajar set |
| DELETE | `/api/v1/learning-planning/modul-ajar-sets/:id` | `tp:DELETE` | Delete Modul Ajar set |
| POST | `/api/v1/learning-planning/modul-ajar-sets/:id/approve` | `tp:APPROVE` | Approve Modul Ajar set |

#### Individual Modul Ajar

| Method | Endpoint | Permission | Description |
|--------|----------|------------|-------------|
| POST | `/api/v1/learning-planning/modul-ajar` | `tp:CREATE` | Create Modul Ajar |
| GET | `/api/v1/learning-planning/modul-ajar` | - | List Modul Ajar |
| GET | `/api/v1/learning-planning/modul-ajar/:id` | - | Get Modul Ajar by ID |
| PUT | `/api/v1/learning-planning/modul-ajar/:id` | `tp:UPDATE` | Update Modul Ajar |
| DELETE | `/api/v1/learning-planning/modul-ajar/:id` | `tp:DELETE` | Delete Modul Ajar |

---

### Assessment

**Middleware:** Requires `assessment:READ` permission

#### Assessments

| Method | Endpoint | Permission | Description |
|--------|----------|------------|-------------|
| POST | `/api/v1/assessment` | `assessment:CREATE` | Create assessment |
| GET | `/api/v1/assessment` | - | List assessments |
| GET | `/api/v1/assessment/:id` | - | Get assessment by ID |
| PUT | `/api/v1/assessment/:id` | `assessment:UPDATE` | Update assessment |
| POST | `/api/v1/assessment/:id/approve` | `assessment:APPROVE` | Approve assessment |

#### Rubrics

| Method | Endpoint | Permission | Description |
|--------|----------|------------|-------------|
| POST | `/api/v1/assessment/rubrics` | `assessment:CREATE` | Create rubric |
| GET | `/api/v1/assessment/rubrics` | - | List rubrics |
| GET | `/api/v1/assessment/rubrics/:id` | - | Get rubric by ID |
| PUT | `/api/v1/assessment/rubrics/:id` | `assessment:UPDATE` | Update rubric |
| DELETE | `/api/v1/assessment/rubrics/:id` | `assessment:DELETE` | Delete rubric |

#### Evidences

| Method | Endpoint | Permission | Description |
|--------|----------|------------|-------------|
| POST | `/api/v1/assessment/evidences/upload` | `assessment:CREATE` | Upload evidence file |
| GET | `/api/v1/assessment/evidences/:id` | - | Get evidence by ID |
| POST | `/api/v1/assessment/evidences` | `assessment:CREATE` | Create evidence |
| GET | `/api/v1/assessment/evidences` | - | List evidences |
| PUT | `/api/v1/assessment/evidences/:id` | `assessment:UPDATE` | Update evidence |
| DELETE | `/api/v1/assessment/evidences/:id` | `assessment:DELETE` | Delete evidence |

#### Evaluations

| Method | Endpoint | Permission | Description |
|--------|----------|------------|-------------|
| POST | `/api/v1/assessment/evaluations` | `assessment:CREATE` | Create evaluation |
| GET | `/api/v1/assessment/evaluations` | - | List evaluations |
| GET | `/api/v1/assessment/evaluations/:id` | - | Get evaluation by ID |
| PUT | `/api/v1/assessment/evaluations/:id` | `assessment:UPDATE` | Update evaluation |
| GET | `/api/v1/assessment/evaluations/history/:evidence_id` | - | Get evaluation history |
| GET | `/api/v1/assessment/evaluations/:id/feedback-history` | - | Get feedback history |

---

### Reporting

**Middleware:** Requires `reporting:READ` permission

#### Narrative Reports

| Method | Endpoint | Permission | Description |
|--------|----------|------------|-------------|
| POST | `/api/v1/reporting/narrative-reports` | `reporting:CREATE` | Create narrative report |
| GET | `/api/v1/reporting/narrative-reports` | - | List narrative reports |
| GET | `/api/v1/reporting/narrative-reports/:id` | - | Get narrative report by ID |
| PUT | `/api/v1/reporting/narrative-reports/:id` | `reporting:UPDATE` | Update narrative report |
| DELETE | `/api/v1/reporting/narrative-reports/:id` | `reporting:DELETE` | Delete narrative report |
| POST | `/api/v1/reporting/narrative-reports/:id/refresh-achievement` | `reporting:UPDATE` | Refresh achievement data |

---

### Achievement

**Middleware:** Requires `reporting:READ` permission

#### Student Achievement

| Method | Endpoint | Permission | Description |
|--------|----------|------------|-------------|
| GET | `/api/v1/students/:id/achievement` | - | Get student achievement |
| GET | `/api/v1/students/:id/progress` | - | Get student competency progress |

#### Class Achievement

| Method | Endpoint | Permission | Description |
|--------|----------|------------|-------------|
| GET | `/api/v1/classes/:id/achievement` | - | Get class achievement summary |

---

### Academic Foundation (Sprint 4)

#### Academic Years

**Middleware:** Requires `academic_year:READ` permission

| Method | Endpoint | Permission | Description |
|--------|----------|------------|-------------|
| POST | `/api/v1/academic-years` | `academic_year:CREATE` | Create academic year |
| GET | `/api/v1/academic-years` | - | List academic years |
| GET | `/api/v1/academic-years/:id` | - | Get academic year by ID |
| PUT | `/api/v1/academic-years/:id` | `academic_year:UPDATE` | Update academic year |
| POST | `/api/v1/academic-years/:id/activate` | `academic_year:ACTIVATE` | Activate academic year |
| POST | `/api/v1/academic-years/:id/archive` | `academic_year:ARCHIVE` | Archive academic year |

**Query Parameters (GET /academic-years):**
- `school_id` (optional) - Filter by school
- `status` (optional) - Filter by status
- `page` (default: 1) - Page number
- `page_size` (default: 20) - Items per page

#### Semesters

**Middleware:** Requires `semester:READ` permission

| Method | Endpoint | Permission | Description |
|--------|----------|------------|-------------|
| POST | `/api/v1/semesters` | `semester:CREATE` | Create semester |
| GET | `/api/v1/semesters` | - | List semesters |
| GET | `/api/v1/semesters/:id` | - | Get semester by ID |
| PUT | `/api/v1/semesters/:id` | `semester:UPDATE` | Update semester |
| DELETE | `/api/v1/semesters/:id` | `semester:DELETE` | Delete semester |

**Query Parameters (GET /semesters):**
- `academic_year_id` (optional) - Filter by academic year
- `status` (optional) - Filter by status
- `page` (default: 1) - Page number
- `page_size` (default: 20) - Items per page

#### Subject Categories

**Middleware:** Requires `subject_category:READ` permission

| Method | Endpoint | Permission | Description |
|--------|----------|------------|-------------|
| POST | `/api/v1/subject-categories` | `subject_category:CREATE` | Create subject category |
| GET | `/api/v1/subject-categories` | - | List subject categories |
| PUT | `/api/v1/subject-categories/:id` | `subject_category:UPDATE` | Update subject category |
| DELETE | `/api/v1/subject-categories/:id` | `subject_category:DELETE` | Delete subject category |

**Query Parameters (GET /subject-categories):**
- `active_only` (optional, boolean) - Filter active only
- `page` (default: 1) - Page number
- `page_size` (default: 20) - Items per page

#### Graduate Profile Dimensions

**Middleware:** Requires `graduate_profile:READ` permission

| Method | Endpoint | Permission | Description |
|--------|----------|------------|-------------|
| POST | `/api/v1/graduate-profile-dimensions` | `graduate_profile:CREATE` | Create graduate profile dimension |
| GET | `/api/v1/graduate-profile-dimensions` | - | List graduate profile dimensions |
| PUT | `/api/v1/graduate-profile-dimensions/:id` | `graduate_profile:UPDATE` | Update graduate profile dimension |
| DELETE | `/api/v1/graduate-profile-dimensions/:id` | `graduate_profile:DELETE` | Delete graduate profile dimension |

**Query Parameters (GET /graduate-profile-dimensions):**
- `active_only` (optional, boolean) - Filter active only
- `page` (default: 1) - Page number
- `page_size` (default: 20) - Items per page

#### CP Alignments

**Middleware:** Requires `cp_alignment:READ` permission

| Method | Endpoint | Permission | Description |
|--------|----------|------------|-------------|
| POST | `/api/v1/cp-alignments` | `cp_alignment:CREATE` | Create CP alignment |
| POST | `/api/v1/cp-alignments/bulk` | `cp_alignment:CREATE` | Bulk create CP alignments |
| GET | `/api/v1/cp-alignments` | - | List CP alignments |
| GET | `/api/v1/cp-alignments/report` | - | Generate CP alignment report |
| PUT | `/api/v1/cp-alignments/:id` | `cp_alignment:UPDATE` | Update CP alignment |
| DELETE | `/api/v1/cp-alignments/:id` | `cp_alignment:DELETE` | Delete CP alignment |

**Query Parameters (GET /cp-alignments):**
- `curriculum_subject_id` (optional) - Filter by subject
- `graduate_profile_dimension_id` (optional) - Filter by dimension
- `active_only` (optional, boolean) - Filter active only
- `page` (default: 1) - Page number
- `page_size` (default: 20) - Items per page

#### System Configurations

**Middleware:** Requires `system_config:READ` permission

| Method | Endpoint | Permission | Description |
|--------|----------|------------|-------------|
| POST | `/api/v1/system-configurations` | `system_config:CREATE` | Create system configuration |
| GET | `/api/v1/system-configurations` | - | List system configurations |
| GET | `/api/v1/system-configurations/:id` | - | Get system configuration by ID |
| GET | `/api/v1/system-configurations/by-key/:key` | - | Get system configuration by key |
| PUT | `/api/v1/system-configurations/:id` | `system_config:UPDATE` | Update system configuration |
| DELETE | `/api/v1/system-configurations/:id` | `system_config:DELETE` | Delete system configuration |

**Query Parameters (GET /system-configurations):**
- `category` (optional) - Filter by category
- `active_only` (optional, boolean) - Filter active only
- `page` (default: 1) - Page number
- `page_size` (default: 20) - Items per page

---

## Response Format

### Success Response

```json
{
  "data": { ... },
  "message": "Success message (optional)"
}
```

### List Response with Pagination

```json
{
  "data": { ... },
  "meta": {
    "total": 100,
    "page": 1,
    "page_size": 20,
    "total_pages": 5
  }
}
```

### Error Response

```json
{
  "error": "Error message",
  "code": "ERROR_CODE (optional)",
  "status": 400
}
```

### Common HTTP Status Codes

| Code | Description |
|------|-------------|
| 200 | Success |
| 201 | Created |
| 400 | Bad Request |
| 401 | Unauthorized |
| 403 | Forbidden |
| 404 | Not Found |
| 500 | Internal Server Error |

---

## Rate Limiting

Currently not implemented. Consider adding rate limiting for:
- Authentication endpoints (login, refresh)
- Public endpoints
- Write operations

---

## CORS Configuration

CORS is enabled for cross-origin requests. Configuration is in `internal/middleware/cors.go`.

---

## API Documentation

Interactive API documentation is available at:
- **Scalar UI:** `/scalar`
- **Swagger UI:** `/swagger`
- **OpenAPI Spec:** `/openapi.json`

---

## Business Rules Enforced

### Academic Year
- **BR-001:** Only one ACTIVE academic year per school
- **BR-002:** Academic year dates cannot overlap
- **BR-003:** Academic year must be created with lead time (configurable)
- Only DRAFT academic years can be modified or activated
- Exactly two semesters required before activation

### Semester
- Only DRAFT academic years can have semesters added/modified
- Semesters cannot overlap in dates within the same academic year
- Semester sequence numbers must be unique within academic year

### Curriculum Governance
- **BR-004:** Minimum CP coverage percentage (configurable via system config)
- Subject category codes must be unique
- Graduate profile dimension codes and sequence numbers must be unique
- CP alignments cannot have duplicate subject-dimension combinations

### System Configuration
- Only SYSTEM_ADMIN can create/delete system configurations
- System configurations (is_system=true) cannot be deleted
- Configuration keys must be unique

---

## Security Best Practices

1. **JWT Tokens:** Short-lived access tokens (24h) with refresh token rotation
2. **Password Hashing:** bcrypt with appropriate cost factor
3. **SQL Injection:** Parameterized queries via sqlx
4. **XSS Prevention:** Input validation and sanitization
5. **CSRF Protection:** Consider implementing CSRF tokens for state-changing operations
6. **HTTPS:** Enforce HTTPS in production

---

## Future Enhancements

1. **GraphQL:** Consider GraphQL for complex queries
2. **Webhooks:** Add webhook support for event notifications
3. **Batch Operations:** Expand bulk operations beyond CP alignments
4. **Field Selection:** Add field selection (GraphQL-like) for list endpoints
5. **Sorting:** Add configurable sorting for list endpoints
6. **Filtering:** Expand filtering capabilities with advanced operators
7. **API Versioning:** Implement proper API versioning strategy
8. **Request Validation:** Add comprehensive request validation schemas

---

## Conclusion

The NUSA Platform API is well-structured with consistent patterns, proper authentication/authorization, and comprehensive coverage of the Kurikulum Merdeka education domain. The API follows RESTful principles and provides a solid foundation for the frontend application. The main areas for improvement are around rate limiting, advanced filtering/sorting, and expanding bulk operations.

# 33_INTEGRATION_TEST_PLAN.md

## Foundation Document for NUSA Education Platform

**Version**: 1.0
**Date**: June 2026
**Status**: FOUNDATION DOCUMENT
**Alignment**: Aligned with 27_BACKEND_FOUNDATION_DESIGN.md, 28_AUTHENTICATION_DESIGN.md, 29_CURRICULUM_MODULE_DESIGN.md, 30_TP_GENERATION_MODULE_DESIGN.md, 31_WORKFLOW_ENGINE_DESIGN.md, 32_FRONTEND_MODULE_DESIGN.md

**Purpose**: Define MVP integration testing strategy for NUSA. This document defines happy path tests, failure tests, security tests, permission tests, AI failure scenarios, and database consistency tests for authentication, curriculum, TP generation, workflow, and AI integration.

---

# SECTION 1 — Executive Summary

## Why Integration Testing Matters

Integration testing ensures:
- Components work together correctly
- API contracts are honored
- Data flows correctly through the system
- Security controls are enforced
- Error handling is robust
- Database consistency is maintained

## Test Coverage Areas

**Authentication**: Login, logout, token refresh, permission checks

**Curriculum**: CP retrieval, subject tree, search functionality

**TP Generation**: AI generation, regeneration, workflow transitions

**Workflow**: State transitions, approval flow, history tracking

**AI Integration**: AI gateway calls, failure handling, cost tracking

## Testing Tools

| Tool | Purpose |
|------|---------|
| Go testing framework | Backend integration tests |
| Testcontainers | Database container for tests |
| Mock AI Gateway | Mock AI provider responses |
| Postman/Newman | API contract testing |
| Playwright | E2E testing (future) |

---

# SECTION 2 — Happy Path Tests

## Authentication Happy Path

### Test: Successful Login

**Description**: User logs in with valid credentials

**Steps**:
1. Create test user with known credentials
2. POST /api/v1/public/login with email and password
3. Verify response contains access_token, refresh_token, and user data
4. Verify tokens are stored in database
5. Verify user.last_login_at is updated

**Expected Result**:
- HTTP 200 OK
- Response contains access_token, refresh_token, token_type, expires_in, user
- Refresh token saved in refresh_tokens table
- User last_login_at updated

**Acceptance Criteria**:
- [ ] Login succeeds with valid credentials
- [ ] Access token is valid JWT
- [ ] Refresh token is saved with 7-day expiration
- [ ] User last_login_at is updated

### Test: Successful Token Refresh

**Description**: User refreshes access token using valid refresh token

**Steps**:
1. Login to get refresh token
2. POST /api/v1/public/refresh with refresh_token
3. Verify response contains new access_token and refresh_token
4. Verify old refresh token is deleted
5. Verify new refresh token is saved

**Expected Result**:
- HTTP 200 OK
- Response contains new access_token and refresh_token
- Old refresh token deleted from database
- New refresh token saved with 7-day expiration

**Acceptance Criteria**:
- [ ] Token refresh succeeds with valid refresh token
- [ ] Old refresh token is deleted
- [ ] New refresh token is saved
- [ ] Access token has 1-hour expiration

### Test: Successful Logout

**Description**: User logs out successfully

**Steps**:
1. Login to get tokens
2. POST /api/v1/auth/logout with access token
3. Verify all refresh tokens for user are deleted
4. Verify access token still valid until expiration

**Expected Result**:
- HTTP 204 No Content
- All refresh tokens for user deleted from database

**Acceptance Criteria**:
- [ ] Logout succeeds
- [ ] All refresh tokens deleted
- [ ] Access token remains valid until expiration

## Curriculum Happy Path

### Test: Get Active Curriculum Version

**Description**: Retrieve active curriculum version

**Steps**:
1. Create curriculum version with status ACTIVE
2. GET /api/v1/public/curriculum/version
3. Verify response contains active version

**Expected Result**:
- HTTP 200 OK
- Response contains curriculum_version_code, effective_year, status = ACTIVE

**Acceptance Criteria**:
- [ ] Active curriculum version retrieved
- [ ] Only one ACTIVE version returned
- [ ] Response contains required fields

### Test: List Subjects

**Description**: List all subjects from active curriculum

**Steps**:
1. Create subjects with active curriculum version
2. GET /api/v1/public/curriculum/subjects
3. Verify response contains all subjects

**Expected Result**:
- HTTP 200 OK
- Response contains array of subjects with code, name

**Acceptance Criteria**:
- [ ] All subjects retrieved
- [ ] Only subjects from active version
- [ ] Response contains required fields

### Test: Get Subject Tree

**Description**: Retrieve subject hierarchy tree

**Steps**:
1. Create subject with phases, elements, subelements, CPs
2. GET /api/v1/public/curriculum/subjects/:id/tree
3. Verify response contains complete hierarchy

**Expected Result**:
- HTTP 200 OK
- Response contains subject → phases → elements → subelements → CPs

**Acceptance Criteria**:
- [ ] Complete hierarchy retrieved
- [ ] All levels present
- [ ] CPs included in subelements

### Test: Search CPs

**Description**: Search CPs by text

**Steps**:
1. Create CPs with various text
2. GET /api/v1/public/curriculum/cp/search?q=search_term
3. Verify response contains matching CPs

**Expected Result**:
- HTTP 200 OK
- Response contains CPs matching search term

**Acceptance Criteria**:
- [ ] Search returns matching CPs
- [ ] Case-insensitive search
- [ ] Pagination works correctly

## TP Generation Happy Path

### Test: Generate TP Set

**Description**: Generate TP Set from CP using AI

**Steps**:
1. Create test CP
2. Mock AI gateway to return valid TP data
3. POST /api/v1/curriculum/cp/:cp_id/tp-sets/generate with preferences
4. Verify TP Set created with DRAFT status
5. Verify TP Items created
6. Verify AI generation log created
7. Verify workflow history created

**Expected Result**:
- HTTP 201 Created
- Response contains TP Set with id, version_no = 1, status = DRAFT
- TP Items created with correct sequence
- AI generation log saved
- Workflow history entry created

**Acceptance Criteria**:
- [ ] TP Set created successfully
- [ ] Version number = 1
- [ ] Status = DRAFT
- [ ] TP Items created
- [ ] AI generation log created
- [ ] Workflow history created

### Test: Regenerate TP Set

**Description**: Regenerate TP Set creating new version

**Steps**:
1. Create TP Set in DRAFT status
2. Mock AI gateway to return new TP data
3. POST /api/v1/curriculum/tp-sets/:id/regenerate
4. Verify new TP Set created with version_no = 2
5. Verify old TP Set remains unchanged

**Expected Result**:
- HTTP 201 Created
- New TP Set with version_no = 2
- Old TP Set unchanged

**Acceptance Criteria**:
- [ ] New TP Set created
- [ ] Version number incremented
- [ ] Old TP Set unchanged
- [ ] New AI generation log created

### Test: Submit TP Set for Review

**Description**: Submit TP Set for approval

**Steps**:
1. Create TP Set in DRAFT status
2. POST /api/v1/curriculum/tp-sets/:id/submit with reason
3. Verify status changed to UNDER_REVIEW
4. Verify workflow history created

**Expected Result**:
- HTTP 204 No Content
- TP Set status = UNDER_REVIEW
- Workflow history entry with action = SUBMIT

**Acceptance Criteria**:
- [ ] Status changed to UNDER_REVIEW
- [ ] Workflow history created
- [ ] Reason saved

### Test: Approve TP Set

**Description**: Approve TP Set

**Steps**:
1. Create TP Set in UNDER_REVIEW status
2. POST /api/v1/curriculum/tp-sets/:id/approve with reason
3. Verify status changed to APPROVED
4. Verify approved_by and approved_at set
5. Verify workflow history created

**Expected Result**:
- HTTP 204 No Content
- TP Set status = APPROVED
- approved_by set
- approved_at set
- Workflow history entry with action = APPROVE

**Acceptance Criteria**:
- [ ] Status changed to APPROVED
- [ ] approved_by set
- [ ] approved_at set
- [ ] Workflow history created

## Workflow Happy Path

### Test: Complete Workflow Flow

**Description**: Complete workflow from generation to approval

**Steps**:
1. Generate TP Set (DRAFT)
2. Submit for review (UNDER_REVIEW)
3. Approve (APPROVED)
4. Verify workflow history has all entries
5. Verify state transitions valid

**Expected Result**:
- All state transitions succeed
- Workflow history has 3 entries (CREATE, SUBMIT, APPROVE)

**Acceptance Criteria**:
- [ ] All transitions succeed
- [ ] Workflow history complete
- [ ] States valid per state machine

### Test: Get Workflow History

**Description**: Retrieve workflow history for artifact

**Steps**:
1. Create TP Set with workflow actions
2. GET /api/v1/workflow/tp_set/:id/history
3. Verify response contains all history entries

**Expected Result**:
- HTTP 200 OK
- Response contains array of history entries in reverse chronological order

**Acceptance Criteria**:
- [ ] All history entries retrieved
- [ ] Entries ordered by created_at DESC
- [ ] User information included

---

# SECTION 3 — Failure Tests

## Authentication Failure Tests

### Test: Invalid Credentials

**Description**: Login with invalid credentials

**Steps**:
1. POST /api/v1/public/login with invalid email/password
2. Verify error response

**Expected Result**:
- HTTP 401 Unauthorized
- Response contains error message

**Acceptance Criteria**:
- [ ] Login fails with invalid credentials
- [ ] Error message returned
- [ ] No tokens returned

### Test: Invalid Refresh Token

**Description**: Refresh with invalid refresh token

**Steps**:
1. POST /api/v1/public/refresh with invalid refresh_token
2. Verify error response

**Expected Result**:
- HTTP 401 Unauthorized
- Response contains error message

**Acceptance Criteria**:
- [ ] Refresh fails with invalid token
- [ ] Error message returned

### Test: Expired Refresh Token

**Description**: Refresh with expired refresh token

**Steps**:
1. Create refresh token with past expiration
2. POST /api/v1/public/refresh with expired token
3. Verify error response

**Expected Result**:
- HTTP 401 Unauthorized
- Response contains error message

**Acceptance Criteria**:
- [ ] Refresh fails with expired token
- [ ] Error message returned
- [ ] Token deleted from database

### Test: Missing Authorization Header

**Description**: Access protected endpoint without token

**Steps**:
1. GET /api/v1/curriculum/tp-sets without Authorization header
2. Verify error response

**Expected Result**:
- HTTP 401 Unauthorized
- Response contains error message

**Acceptance Criteria**:
- [ ] Request fails without auth header
- [ ] Error message returned

## Curriculum Failure Tests

### Test: CP Not Found

**Description**: Request non-existent CP

**Steps**:
1. GET /api/v1/public/curriculum/cp/:invalid_id
2. Verify error response

**Expected Result**:
- HTTP 404 Not Found
- Response contains error message

**Acceptance Criteria**:
- [ ] Request fails for non-existent CP
- [ ] Error message returned

### Test: Invalid Search Parameters

**Description**: Search with invalid parameters

**Steps**:
1. GET /api/v1/public/curriculum/cp/search with invalid params
2. Verify error response

**Expected Result**:
- HTTP 400 Bad Request
- Response contains error message

**Acceptance Criteria**:
- [ ] Request fails with invalid params
- [ ] Error message returned

## TP Generation Failure Tests

### Test: Generate with Invalid CP

**Description**: Generate TP Set for non-existent CP

**Steps**:
1. POST /api/v1/curriculum/cp/:invalid_id/tp-sets/generate
2. Verify error response

**Expected Result**:
- HTTP 404 Not Found
- Response contains error message

**Acceptance Criteria**:
- [ ] Generation fails for non-existent CP
- [ ] Error message returned
- [ ] No TP Set created

### Test: Regenerate Approved TP Set

**Description**: Attempt to regenerate approved TP Set

**Steps**:
1. Create TP Set in APPROVED status
2. POST /api/v1/curriculum/tp-sets/:id/regenerate
3. Verify error response

**Expected Result**:
- HTTP 400 Bad Request
- Response contains error message

**Acceptance Criteria**:
- [ ] Regeneration fails for approved TP Set
- [ ] Error message returned
- [ ] No new version created

### Test: Submit Invalid Status

**Description**: Submit TP Set in invalid status

**Steps**:
1. Create TP Set in APPROVED status
2. POST /api/v1/curriculum/tp-sets/:id/submit
3. Verify error response

**Expected Result**:
- HTTP 400 Bad Request
- Response contains error message

**Acceptance Criteria**:
- [ ] Submit fails for invalid status
- [ ] Error message returned
- [ ] Status unchanged

### Test: Approve Invalid Status

**Description**: Approve TP Set in invalid status

**Steps**:
1. Create TP Set in DRAFT status
2. POST /api/v1/curriculum/tp-sets/:id/approve
3. Verify error response

**Expected Result**:
- HTTP 400 Bad Request
- Response contains error message

**Acceptance Criteria**:
- [ ] Approve fails for invalid status
- [ ] Error message returned
- [ ] Status unchanged

## Workflow Failure Tests

### Test: Invalid State Transition

**Description**: Attempt invalid state transition

**Steps**:
1. Create TP Set in DRAFT status
2. Attempt transition to APPROVED (skipping UNDER_REVIEW)
3. Verify error response

**Expected Result**:
- HTTP 400 Bad Request
- Response contains error message

**Acceptance Criteria**:
- [ ] Invalid transition fails
- [ ] Error message returned
- [ ] Status unchanged

---

# SECTION 4 — Security Tests

## Authentication Security Tests

### Test: SQL Injection in Login

**Description**: Attempt SQL injection via email field

**Steps**:
1. POST /api/v1/public/login with email containing SQL injection
2. Verify request fails safely

**Expected Result**:
- HTTP 401 Unauthorized
- No SQL error exposed

**Acceptance Criteria**:
- [ ] SQL injection attempt fails
- [ ] No SQL error exposed
- [ ] Error message generic

### Test: Brute Force Protection

**Description**: Multiple failed login attempts

**Steps**:
1. Attempt 10 failed logins
2. Verify rate limiting (if implemented)

**Expected Result**:
- All requests fail
- Rate limiting may be enforced (future)

**Acceptance Criteria**:
- [ ] All failed attempts rejected
- [ ] No account lockout (MVP)
- [ ] Rate limiting optional (future)

### Test: Token Theft Protection

**Description**: Use stolen access token after logout

**Steps**:
1. Login to get access token
2. Logout
3. Attempt to use access token
4. Verify request fails

**Expected Result**:
- Request succeeds (access token valid until expiration)
- MVP does not revoke access tokens on logout

**Acceptance Criteria**:
- [ ] Access token remains valid after logout (MVP)
- [ ] Refresh tokens deleted on logout
- [ ] Future: access token revocation

## Authorization Security Tests

### Test: Unauthorized Access to Admin Endpoint

**Description**: Teacher attempts to access admin endpoint

**Steps**:
1. Login as TEACHER
2. POST /api/v1/admin/curriculum/import
3. Verify error response

**Expected Result**:
- HTTP 403 Forbidden
- Response contains error message

**Acceptance Criteria**:
- [ ] Teacher cannot access admin endpoint
- [ ] Error message returned
- [ ] No data modified

### Test: School Isolation

**Description**: Teacher attempts to access other school's data

**Steps**:
1. Login as TEACHER from School A
2. Attempt to access TP Set from School B
3. Verify error response

**Expected Result**:
- HTTP 403 Forbidden
- Response contains error message

**Acceptance Criteria**:
- [ ] Teacher cannot access other school's data
- [ ] Error message returned
- [ ] No data exposed

### Test: Cross-School Data Access

**Description**: School Admin attempts to access other school's data

**Steps**:
1. Login as SCHOOL_ADMIN from School A
2. Attempt to access TP Set from School B
3. Verify error response

**Expected Result**:
- HTTP 403 Forbidden
- Response contains error message

**Acceptance Criteria**:
- [ ] School Admin cannot access other school's data
- [ ] Error message returned
- [ ] No data exposed

## Data Security Tests

### Test: Sensitive Data Exposure

**Description**: Verify password hash never exposed

**Steps**:
1. GET /api/v1/auth/me
2. Verify response does not contain password_hash

**Expected Result**:
- HTTP 200 OK
- Response does not contain password_hash

**Acceptance Criteria**:
- [ ] Password hash not exposed
- [ ] Only safe fields returned
- [ ] Refresh token not exposed

### Test: Token Storage

**Description**: Verify refresh tokens stored securely

**Steps**:
1. Login
2. Check database for refresh token
3. Verify token is hashed or encrypted (future)

**Expected Result**:
- Refresh token stored in database
- MVP: stored as plain text (acceptable for MVP)
- Future: hashed or encrypted

**Acceptance Criteria**:
- [ ] Refresh token stored
- [ ] MVP: plain text acceptable
- [ ] Future: implement hashing

---

# SECTION 5 — Permission Tests

## Role-Based Permission Tests

### Test: SYSTEM_ADMIN Permissions

**Description**: Verify SYSTEM_ADMIN can access all resources

**Steps**:
1. Login as SYSTEM_ADMIN
2. Access curriculum, TP generation, workflow endpoints
3. Verify all requests succeed

**Expected Result**:
- All requests succeed
- Full access granted

**Acceptance Criteria**:
- [ ] Can access all curriculum endpoints
- [ ] Can access all TP generation endpoints
- [ ] Can access all workflow endpoints
- [ ] Can access admin endpoints

### Test: SCHOOL_ADMIN Permissions

**Description**: Verify SCHOOL_ADMIN has correct permissions

**Steps**:
1. Login as SCHOOL_ADMIN
2. Access school-level resources
3. Attempt to access system-level resources
4. Verify correct access/denial

**Expected Result**:
- School-level access granted
- System-level access denied

**Acceptance Criteria**:
- [ ] Can access own school's curriculum
- [ ] Can access own school's TP Sets
- [ ] Can approve/reject own school's TP Sets
- [ ] Cannot access other schools' data
- [ ] Cannot access admin endpoints

### Test: TEACHER Permissions

**Description**: Verify TEACHER has correct permissions

**Steps**:
1. Login as TEACHER
2. Access teacher-level resources
3. Attempt to access admin resources
4. Verify correct access/denial

**Expected Result**:
- Teacher-level access granted
- Admin access denied

**Acceptance Criteria**:
- [ ] Can access curriculum (read-only)
- [ ] Can generate TP Sets
- [ ] Can submit TP Sets for review
- [ ] Cannot approve/reject TP Sets
- [ ] Cannot access admin endpoints

## Permission Matrix Tests

### Test: Curriculum Permissions

**Description**: Verify curriculum permissions per role

**Steps**:
1. Test each role against curriculum endpoints
2. Verify access matches permission matrix

**Expected Result**:
- Access matches defined permission matrix

**Acceptance Criteria**:
- [ ] SYSTEM_ADMIN: full access
- [ ] SCHOOL_ADMIN: read-only active version
- [ ] TEACHER: read-only active version

### Test: TP Generation Permissions

**Description**: Verify TP generation permissions per role

**Steps**:
1. Test each role against TP generation endpoints
2. Verify access matches permission matrix

**Expected Result**:
- Access matches defined permission matrix

**Acceptance Criteria**:
- [ ] SYSTEM_ADMIN: full access
- [ ] SCHOOL_ADMIN: generate, submit, approve, reject
- [ ] TEACHER: generate, submit

### Test: Workflow Permissions

**Description**: Verify workflow permissions per role

**Steps**:
1. Test each role against workflow endpoints
2. Verify access matches permission matrix

**Expected Result**:
- Access matches defined permission matrix

**Acceptance Criteria**:
- [ ] SYSTEM_ADMIN: full access
- [ ] SCHOOL_ADMIN: approve, reject, view history
- [ ] TEACHER: submit, view own history

---

# SECTION 6 — AI Failure Scenarios

## AI Gateway Failure Tests

### Test: AI Gateway Timeout

**Description**: AI gateway times out

**Steps**:
1. Mock AI gateway to timeout
2. POST /api/v1/curriculum/cp/:cp_id/tp-sets/generate
3. Verify error handling

**Expected Result**:
- HTTP 504 Gateway Timeout or 500 Internal Server Error
- Error message returned
- No TP Set created
- No AI generation log created

**Acceptance Criteria**:
- [ ] Timeout handled gracefully
- [ ] Error message returned
- [ ] No partial data created
- [ ] User can retry

### Test: AI Gateway Returns Error

**Description**: AI gateway returns error response

**Steps**:
1. Mock AI gateway to return error
2. POST /api/v1/curriculum/cp/:cp_id/tp-sets/generate
3. Verify error handling

**Expected Result**:
- HTTP 500 Internal Server Error
- Error message returned
- No TP Set created
- AI generation log created with status = FAILED

**Acceptance Criteria**:
- [ ] Error handled gracefully
- [ ] Error message returned
- [ ] AI generation log created with FAILED status
- [ ] No partial data created

### Test: AI Gateway Returns Invalid JSON

**Description**: AI gateway returns invalid JSON

**Steps**:
1. Mock AI gateway to return invalid JSON
2. POST /api/v1/curriculum/cp/:cp_id/tp-sets/generate
3. Verify error handling

**Expected Result**:
- HTTP 500 Internal Server Error
- Error message returned
- No TP Set created
- AI generation log created with status = FAILED

**Acceptance Criteria**:
- [ ] Invalid JSON handled gracefully
- [ ] Error message returned
- [ ] AI generation log created with FAILED status
- [ ] No partial data created

### Test: AI Gateway Returns Empty Response

**Description**: AI gateway returns empty TP array

**Steps**:
1. Mock AI gateway to return empty TP array
2. POST /api/v1/curriculum/cp/:cp_id/tp-sets/generate
3. Verify error handling

**Expected Result**:
- HTTP 400 Bad Request
- Error message returned
- No TP Set created

**Acceptance Criteria**:
- [ ] Empty response handled
- [ ] Error message returned
- [ ] No TP Set created
- [ ] Validation error

### Test: AI Gateway Returns Malformed TP Data

**Description**: AI gateway returns TP data with missing required fields

**Steps**:
1. Mock AI gateway to return malformed TP data
2. POST /api/v1/curriculum/cp/:cp_id/tp-sets/generate
3. Verify error handling

**Expected Result**:
- HTTP 400 Bad Request
- Error message returned
- No TP Set created

**Acceptance Criteria**:
- [ ] Malformed data handled
- [ ] Error message returned
- [ ] No TP Set created
- [ ] Validation error

## AI Cost Tracking Tests

### Test: Cost Tracking on Success

**Description**: Verify cost tracked on successful generation

**Steps**:
1. Mock AI gateway to return cost data
2. POST /api/v1/curriculum/cp/:cp_id/tp-sets/generate
3. Verify AI generation log has cost data

**Expected Result**:
- AI generation log contains tokens_used, estimated_cost

**Acceptance Criteria**:
- [ ] Tokens used tracked
- [ ] Estimated cost calculated
- [ ] Provider logged
- [ ] Model logged

### Test: Cost Tracking on Failure

**Description**: Verify cost tracked even on failure

**Steps**:
1. Mock AI gateway to return error with cost data
2. POST /api/v1/curriculum/cp/:cp_id/tp-sets/generate
3. Verify AI generation log has cost data

**Expected Result**:
- AI generation log contains tokens_used, estimated_cost, status = FAILED

**Acceptance Criteria**:
- [ ] Cost tracked even on failure
- [ ] Status = FAILED
- [ ] Error logged

---

# SECTION 7 — Database Consistency Tests

## Foreign Key Consistency Tests

### Test: TP Set CP Foreign Key

**Description**: Verify TP Set CP foreign key constraint

**Steps**:
1. Attempt to create TP Set with non-existent cp_id
2. Verify foreign key constraint enforced

**Expected Result**:
- Database error
- No TP Set created

**Acceptance Criteria**:
- [ ] Foreign key constraint enforced
- [ ] Error returned
- [ ] No orphaned records

### Test: TP Item TP Set Foreign Key

**Description**: Verify TP Item TP Set foreign key constraint

**Steps**:
1. Attempt to create TP Item with non-existent tp_set_id
2. Verify foreign key constraint enforced

**Expected Result**:
- Database error
- No TP Item created

**Acceptance Criteria**:
- [ ] Foreign key constraint enforced
- [ ] Error returned
- [ ] No orphaned records

### Test: Cascade Delete

**Description**: Verify cascade delete on TP Set deletion

**Steps**:
1. Create TP Set with TP Items
2. Delete TP Set
3. Verify TP Items also deleted

**Expected Result**:
- TP Set deleted
- TP Items deleted (cascade)

**Acceptance Criteria**:
- [ ] Cascade delete works
- [ ] No orphaned TP Items
- [ ] Foreign key constraints maintained

## Unique Constraint Tests

### Test: TP Set Version Uniqueness

**Description**: Verify unique constraint on (cp_id, version_no)

**Steps**:
1. Create TP Set with cp_id and version_no = 1
2. Attempt to create another TP Set with same cp_id and version_no
3. Verify unique constraint enforced

**Expected Result**:
- Database error
- Second TP Set not created

**Acceptance Criteria**:
- [ ] Unique constraint enforced
- [ ] Error returned
- [ ] No duplicate records

### Test: Refresh Token Uniqueness

**Description**: Verify unique constraint on refresh token

**Steps**:
1. Create refresh token
2. Attempt to create another with same token
3. Verify unique constraint enforced

**Expected Result**:
- Database error
- Second refresh token not created

**Acceptance Criteria**:
- [ ] Unique constraint enforced
- [ ] Error returned
- [ ] No duplicate records

## Transaction Rollback Tests

### Test: TP Generation Rollback on AI Failure

**Description**: Verify transaction rollback on AI failure

**Steps**:
1. Mock AI gateway to return error
2. POST /api/v1/curriculum/cp/:cp_id/tp-sets/generate
3. Verify no partial data created

**Expected Result**:
- No TP Set created
- No TP Items created
- No AI generation log created (or created with FAILED status)

**Acceptance Criteria**:
- [ ] Transaction rolled back
- [ ] No partial data
- [ ] Database consistent

### Test: Workflow History Rollback

**Description**: Verify workflow history roll back on failure

**Steps**:
1. Mock workflow service to fail after history creation
2. Attempt workflow transition
3. Verify history rolled back

**Expected Result**:
- No history entry created
- No state change

**Acceptance Criteria**:
- [ ] Transaction rolled back
- [ ] No partial data
- [ ] Database consistent

## Data Integrity Tests

### Test: Status Check Constraint

**Description**: Verify status check constraint

**Steps**:
1. Attempt to set invalid status
2. Verify check constraint enforced

**Expected Result**:
- Database error
- Status not updated

**Acceptance Criteria**:
- [ ] Check constraint enforced
- [ ] Error returned
- [ ] Only valid statuses allowed

### Test: Estimated Weeks Check Constraint

**Description**: Verify estimated_weeks check constraint

**Steps**:
1. Attempt to set estimated_weeks = 0
2. Verify check constraint enforced

**Expected Result**:
- Database error
- TP Item not created

**Acceptance Criteria**:
- [ ] Check constraint enforced
- [ ] Error returned
- [ ] Only valid values allowed

---

# SECTION 8 — Test Matrix

## Test Matrix Overview

| Test Category | Test Count | Priority | Status |
|---------------|-----------|----------|--------|
| Authentication Happy Path | 3 | High | Pending |
| Authentication Failure | 4 | High | Pending |
| Authentication Security | 3 | High | Pending |
| Curriculum Happy Path | 4 | High | Pending |
| Curriculum Failure | 2 | Medium | Pending |
| TP Generation Happy Path | 4 | High | Pending |
| TP Generation Failure | 4 | High | Pending |
| Workflow Happy Path | 2 | High | Pending |
| Workflow Failure | 1 | Medium | Pending |
| Authorization Security | 3 | High | Pending |
| Data Security | 2 | High | Pending |
| Permission Tests | 3 | High | Pending |
| Permission Matrix Tests | 3 | High | Pending |
| AI Gateway Failure | 5 | High | Pending |
| AI Cost Tracking | 2 | Medium | Pending |
| Foreign Key Consistency | 3 | High | Pending |
| Unique Constraint | 2 | High | Pending |
| Transaction Rollback | 2 | High | Pending |
| Data Integrity | 2 | High | Pending |
| **Total** | **50** | - | **Pending** |

## Test Priority Matrix

| Priority | Test Categories | Test Count |
|----------|----------------|-----------|
| High | Authentication Happy Path, Authentication Failure, Authentication Security, Curriculum Happy Path, TP Generation Happy Path, TP Generation Failure, Workflow Happy Path, Authorization Security, Data Security, Permission Tests, Permission Matrix Tests, AI Gateway Failure, Foreign Key Consistency, Unique Constraint, Transaction Rollback, Data Integrity | 42 |
| Medium | Curriculum Failure, AI Cost Tracking | 6 |
| Low | None | 0 |

---

# SECTION 9 — Test Scenarios

## Scenario 1: Complete User Journey

**Description**: End-to-end user journey from login to TP approval

**Steps**:
1. Login as TEACHER
2. Navigate to curriculum
3. Select subject and CP
4. Generate TP Set
5. Submit TP Set for review
6. Logout
7. Login as SCHOOL_ADMIN
8. Navigate to pending approvals
9. Review and approve TP Set
10. Verify TP Set status = APPROVED

**Expected Result**:
- All steps succeed
- TP Set approved
- Workflow history complete

**Acceptance Criteria**:
- [ ] Login successful
- [ ] Curriculum navigation works
- [ ] TP generation successful
- [ ] Submit successful
- [ ] Logout successful
- [ ] Admin login successful
- [ ] Pending approvals accessible
- [ ] Approval successful
- [ ] TP Set approved
- [ ] Workflow history complete

## Scenario 2: AI Failure Recovery

**Description**: User handles AI generation failure

**Steps**:
1. Login as TEACHER
2. Navigate to CP
3. Generate TP Set (AI fails)
4. Verify error message displayed
5. Retry generation (AI succeeds)
6. Verify TP Set created

**Expected Result**:
- First attempt fails gracefully
- Second attempt succeeds
- TP Set created

**Acceptance Criteria**:
- [ ] First attempt fails with error
- [ ] Error message clear
- [ ] Retry succeeds
- [ ] TP Set created
- [ ] AI generation log shows both attempts

## Scenario 3: Permission Enforcement

**Description**: Verify permissions enforced across roles

**Steps**:
1. Login as TEACHER
2. Attempt to approve TP Set
3. Verify access denied
4. Logout
5. Login as SCHOOL_ADMIN
6. Approve TP Set
7. Verify success

**Expected Result**:
- Teacher cannot approve
- School Admin can approve

**Acceptance Criteria**:
- [ ] Teacher approval denied
- [ ] Error message clear
- [ ] School Admin approval succeeds
- [ ] Workflow history shows correct user

## Scenario 4: Data Consistency After Failure

**Description**: Verify database consistent after failures

**Steps**:
1. Generate TP Set (AI fails mid-transaction)
2. Check database for partial data
3. Verify no partial data
4. Generate TP Set (AI succeeds)
5. Verify complete data

**Expected Result**:
- No partial data after failure
- Complete data after success

**Acceptance Criteria**:
- [ ] No TP Set after failure
- [ ] No TP Items after failure
- [ ] No AI generation log (or FAILED status)
- [ ] Complete TP Set after success
- [ ] Complete TP Items after success
- [ ] AI generation log with SUCCESS status

---

# SECTION 10 — Acceptance Criteria

## MVP Acceptance Criteria

### Authentication

- [ ] Users can login with valid credentials
- [ ] Users cannot login with invalid credentials
- [ ] Access tokens are valid for 1 hour
- [ ] Refresh tokens are valid for 7 days
- [ ] Token refresh works correctly
- [ ] Logout deletes refresh tokens
- [ ] Passwords are hashed with bcrypt
- [ ] Password hash never exposed

### Authorization

- [ ] SYSTEM_ADMIN has full access
- [ ] SCHOOL_ADMIN has school-level access
- [ ] TEACHER has teacher-level access
- [ ] School isolation enforced
- [ ] Permission matrix enforced
- [ ] Unauthorized access denied with 403

### Curriculum

- [ ] Active curriculum version accessible
- [ ] Subjects list accessible
- [ ] Subject tree accessible
- [ ] CP detail accessible
- [ ] CP search works correctly
- [ ] Pagination works correctly
- [ ] Only active version returned to non-admins

### TP Generation

- [ ] TP Set generation works
- [ ] TP Set regeneration works
- [ ] Version numbering correct
- [ ] TP Items created correctly
- [ ] AI generation log created
- [ ] Workflow history created
- [ ] Cannot regenerate approved TP Set

### Workflow

- [ ] State transitions work correctly
- [ ] State machine enforced
- [ ] Submit for review works
- [ ] Approve works
- [ ] Reject works
- [ ] Archive works
- [ ] Workflow history tracked
- [ ] Invalid transitions denied

### AI Integration

- [ ] AI gateway calls work
- [ ] AI failures handled gracefully
- [ ] Cost tracking works
- [ ] No partial data on failure
- [ ] Retry possible after failure
- [ ] AI generation log accurate

### Database

- [ ] Foreign key constraints enforced
- [ ] Unique constraints enforced
- [ ] Check constraints enforced
- [ ] Cascade delete works
- [ ] Transactions roll back on failure
- [ ] No orphaned records
- [ ] Data integrity maintained

### Security

- [ ] SQL injection prevented
- [ ] XSS prevented (future)
- [ ] CSRF prevented (future)
- [ ] Sensitive data not exposed
- [ ] Authorization headers required
- [ ] Invalid tokens rejected

### Performance

- [ ] API response time < 2s for 95% of requests
- [ ] Database queries optimized
- [ ] Indexes used appropriately
- [ ] No N+1 queries

---

# SECTION 11 — Appendix

## Test Environment Setup

### Database Setup

```bash
# Start test database with Docker
docker run -d \
  --name nusa-test-db \
  -e POSTGRES_PASSWORD=testpass \
  -e POSTGRES_DB=nusa_test \
  -p 5433:5432 \
  postgres:15

# Run migrations
migrate -path ./migrations -database "postgres://postgres:testpass@localhost:5433/nusa_test?sslmode=disable" up
```

### Mock AI Gateway

```go
// test/mock_ai_gateway.go
package mock

type MockAIGateway struct {
    Response interface{}
    Error    error
}

func (m *MockAIGateway) Generate(ctx context.Context, prompt string) (*AIResponse, error) {
    if m.Error != nil {
        return nil, m.Error
    }
    return m.Response.(*AIResponse), nil
}
```

## Test Data Setup

### Seed Test Data

```go
// test/seed.go
func SeedTestData(db *sql.DB) error {
    // Create test users
    users := []User{
        {Email: "admin@test.com", PasswordHash: hash("password"), Role: "SYSTEM_ADMIN"},
        {Email: "school_admin@test.com", PasswordHash: hash("password"), Role: "SCHOOL_ADMIN", SchoolID: "school-1"},
        {Email: "teacher@test.com", PasswordHash: hash("password"), Role: "TEACHER", SchoolID: "school-1"},
    }
    
    for _, user := range users {
        _, err := db.Exec("INSERT INTO users (email, password_hash, role, school_id) VALUES ($1, $2, $3, $4)",
            user.Email, user.PasswordHash, user.Role, user.SchoolID)
        if err != nil {
            return err
        }
    }
    
    // Create test school
    _, err := db.Exec("INSERT INTO schools (id, name, code) VALUES ($1, $2, $3)",
        "school-1", "Test School", "SCH-001")
    if err != nil {
        return err
    }
    
    // Create test curriculum version
    _, err = db.Exec("INSERT INTO curriculum_versions (id, curriculum_version_code, effective_year, status) VALUES ($1, $2, $3, $4)",
        "version-1", "KUR2025", 2025, "ACTIVE")
    if err != nil {
        return err
    }
    
    return nil
}
```

## Test Execution

### Run Integration Tests

```bash
# Run all integration tests
go test ./tests/integration/... -v

# Run specific test suite
go test ./tests/integration/auth/... -v

# Run with coverage
go test ./tests/integration/... -v -cover
```

## Continuous Integration

### GitHub Actions Configuration

```yaml
# .github/workflows/integration-tests.yml
name: Integration Tests

on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    
    services:
      postgres:
        image: postgres:15
        env:
          POSTGRES_PASSWORD: testpass
          POSTGRES_DB: nusa_test
        options: >-
          --health-cmd pg_isready
          --health-interval 10s
          --health-timeout 5s
          --health-retries 5
        ports:
          - 5432:5432
    
    steps:
      - uses: actions/checkout@v3
      
      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.21'
      
      - name: Install dependencies
        run: go mod download
      
      - name: Run migrations
        run: |
          go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
          migrate -path ./migrations -database "postgres://postgres:testpass@localhost:5432/nusa_test?sslmode=disable" up
      
      - name: Run integration tests
        run: go test ./tests/integration/... -v -cover
```

## Future Enhancements

### Wave 2

- E2E tests with Playwright
- Load testing with k6
- Security scanning with OWASP ZAP
- Performance testing with JMeter
- API contract testing with Postman
- Visual regression testing

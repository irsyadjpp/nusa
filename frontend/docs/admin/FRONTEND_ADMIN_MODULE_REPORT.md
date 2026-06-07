# Frontend User & School Management Integration Report

**Project**: NUSA Education Platform
**Sprint**: 2
**Phase**: P4 - Frontend User & School Management Integration
**Date**: 2026-06-07
**Status**: ✅ COMPLETED

---

## Executive Summary

The Frontend User & School Management Integration phase successfully created API service layers for Users, Schools, Roles, and Permissions. All CRUD operations are now available as type-safe TypeScript services that integrate with the backend API infrastructure established in previous phases.

**Key Achievement**: Complete API service layer for all admin entities with full TypeScript typing and error handling.

---

## Architecture Overview

```
src/api/
├── client.ts                 # Axios instance with interceptors (from P2)
├── auth.ts                   # Authentication API service (from P2)
├── users.ts                  # User management API service
├── schools.ts                # School management API service
├── roles.ts                  # Role management API service
├── permissions.ts            # Permission management API service
└── index.ts                  # Public API exports
```

---

## Components Implemented

### 1. User API Service (src/api/users.ts)

**Purpose**: API calls for user management operations

**Interfaces**:
```typescript
interface User {
  id: string;
  email: string;
  name: string;
  role_name: string;
  school_name?: string;
  school_id?: string;
  is_active: boolean;
  created_at: string;
  updated_at: string;
}

interface CreateUserRequest {
  email: string;
  name: string;
  password: string;
  role_id?: string;
  school_id?: string;
}

interface UpdateUserRequest {
  name?: string;
  role_id?: string;
  school_id?: string;
  is_active?: boolean;
}

interface ListUsersParams {
  page?: number;
  page_size?: number;
  school_id?: string;
  role_id?: string;
  is_active?: boolean;
  search?: string;
}

interface ListUsersResponse {
  users: User[];
  total: number;
  page: number;
  page_size: number;
}
```

**Functions**:
- `listUsers(params)` - GET `/api/v1/users` - List users with filters and pagination
- `getUser(id)` - GET `/api/v1/users/{id}` - Get user by ID
- `createUser(data)` - POST `/api/v1/users` - Create new user
- `updateUser(id, data)` - PUT `/api/v1/users/{id}` - Update user
- `updateUserStatus(id, is_active)` - PATCH `/api/v1/users/{id}/status` - Update user status
- `deleteUser(id)` - DELETE `/api/v1/users/{id}` - Delete user

**Features**:
- ✅ Full CRUD operations
- ✅ Pagination support
- ✅ Filtering by school, role, status
- ✅ Search functionality
- ✅ Type-safe TypeScript interfaces
- ✅ Consistent error handling
- ✅ Automatic auth header injection (via client interceptors)
- ✅ Automatic token refresh (via client interceptors)

**Usage Example**:
```typescript
import { listUsers, createUser, updateUser, deleteUser } from '@/api';

// List users
const users = await listUsers({ 
  page: 1, 
  page_size: 20, 
  is_active: true 
});

// Create user
const newUser = await createUser({
  email: 'john@example.com',
  name: 'John Doe',
  password: 'SecurePassword123!',
  role_id: 'role-uuid',
  school_id: 'school-uuid'
});

// Update user
const updated = await updateUser(userId, {
  name: 'John Updated',
  is_active: true
});

// Delete user
await deleteUser(userId);
```

---

### 2. School API Service (src/api/schools.ts)

**Purpose**: API calls for school management operations

**Interfaces**:
```typescript
interface School {
  id: string;
  name: string;
  code: string;
  address?: string;
  city?: string;
  state?: string;
  country?: string;
  postal_code?: string;
  phone?: string;
  email?: string;
  website?: string;
  is_active: boolean;
  created_at: string;
  updated_at: string;
}

interface CreateSchoolRequest {
  name: string;
  code: string;
  address?: string;
  city?: string;
  state?: string;
  country?: string;
  postal_code?: string;
  phone?: string;
  email?: string;
  website?: string;
}

interface UpdateSchoolRequest {
  name?: string;
  code?: string;
  address?: string;
  city?: string;
  state?: string;
  country?: string;
  postal_code?: string;
  phone?: string;
  email?: string;
  website?: string;
  is_active?: boolean;
}

interface ListSchoolsParams {
  page?: number;
  page_size?: number;
  is_active?: boolean;
  search?: string;
}

interface ListSchoolsResponse {
  schools: School[];
  total: number;
  page: number;
  page_size: number;
}
```

**Functions**:
- `listSchools(params)` - GET `/api/v1/schools` - List schools with filters and pagination
- `getSchool(id)` - GET `/api/v1/schools/{id}` - Get school by ID
- `createSchool(data)` - POST `/api/v1/schools` - Create new school
- `updateSchool(id, data)` - PUT `/api/v1/schools/{id}` - Update school
- `updateSchoolStatus(id, is_active)` - PATCH `/api/v1/schools/{id}/status` - Update school status
- `deleteSchool(id)` - DELETE `/api/v1/schools/{id}` - Delete school

**Features**:
- ✅ Full CRUD operations
- ✅ Pagination support
- ✅ Filtering by status
- ✅ Search functionality
- ✅ Type-safe TypeScript interfaces
- ✅ Consistent error handling
- ✅ Automatic auth header injection
- ✅ Automatic token refresh

**Usage Example**:
```typescript
import { listSchools, createSchool, updateSchool, deleteSchool } from '@/api';

// List schools
const schools = await listSchools({ 
  page: 1, 
  page_size: 20, 
  is_active: true 
});

// Create school
const newSchool = await createSchool({
  name: 'NUSA High School',
  code: 'NUSA-HS',
  address: '123 Education Street',
  city: 'Jakarta',
  state: 'DKI Jakarta',
  country: 'Indonesia',
  postal_code: '12345',
  phone: '+62 21 1234567',
  email: 'info@nusahs.edu',
  website: 'https://www.nusahs.edu'
});

// Update school
const updated = await updateSchool(schoolId, {
  name: 'NUSA High School (Updated)',
  city: 'Jakarta Selatan'
});

// Delete school
await deleteSchool(schoolId);
```

---

### 3. Role API Service (src/api/roles.ts)

**Purpose**: API calls for role management operations

**Interfaces**:
```typescript
interface Role {
  id: string;
  name: string;
  description?: string;
  is_system?: boolean;
  created_at: string;
  updated_at: string;
}

interface CreateRoleRequest {
  name: string;
  description?: string;
}

interface UpdateRoleRequest {
  name?: string;
  description?: string;
}

interface ListRolesParams {
  page?: number;
  page_size?: number;
  is_system?: boolean;
  search?: string;
}

interface ListRolesResponse {
  roles: Role[];
  total: number;
  page: number;
  page_size: number;
}
```

**Functions**:
- `listRoles(params)` - GET `/api/v1/roles` - List roles with filters and pagination
- `getRole(id)` - GET `/api/v1/roles/{id}` - Get role by ID
- `createRole(data)` - POST `/api/v1/roles` - Create new role
- `updateRole(id, data)` - PUT `/api/v1/roles/{id}` - Update role
- `deleteRole(id)` - DELETE `/api/v1/roles/{id}` - Delete role

**Features**:
- ✅ Full CRUD operations
- ✅ Pagination support
- ✅ Filtering by system role flag
- ✅ Search functionality
- ✅ Type-safe TypeScript interfaces
- ✅ Consistent error handling
- ✅ Automatic auth header injection
- ✅ Automatic token refresh
- ⚠️ System roles protected (deletion restricted by backend)

**Usage Example**:
```typescript
import { listRoles, createRole, updateRole, deleteRole } from '@/api';

// List roles
const roles = await listRoles({ 
  page: 1, 
  page_size: 20,
  is_system: false 
});

// Create role
const newRole = await createRole({
  name: 'SCHOOL_PRINCIPAL',
  description: 'School principal with school-level admin rights'
});

// Update role
const updated = await updateRole(roleId, {
  description: 'Updated description'
});

// Delete role (non-system roles only)
await deleteRole(roleId);
```

---

### 4. Permission API Service (src/api/permissions.ts)

**Purpose**: API calls for permission management operations

**Interfaces**:
```typescript
interface Permission {
  id: string;
  name: string;
  description?: string;
  resource: string;
  action: string;
  created_at: string;
  updated_at: string;
}

interface CreatePermissionRequest {
  name: string;
  description?: string;
  resource: string;
  action: string;
}

interface UpdatePermissionRequest {
  name?: string;
  description?: string;
  resource?: string;
  action?: string;
}

interface ListPermissionsParams {
  page?: number;
  page_size?: number;
  resource?: string;
  action?: string;
  search?: string;
}

interface ListPermissionsResponse {
  permissions: Permission[];
  total: number;
  page: number;
  page_size: number;
}
```

**Functions**:
- `listPermissions(params)` - GET `/api/v1/permissions` - List permissions with filters and pagination
- `getPermission(id)` - GET `/api/v1/permissions/{id}` - Get permission by ID
- `createPermission(data)` - POST `/api/v1/permissions` - Create new permission
- `updatePermission(id, data)` - PUT `/api/v1/permissions/{id}` - Update permission
- `deletePermission(id)` - DELETE `/api/v1/permissions/{id}` - Delete permission

**Features**:
- ✅ Full CRUD operations
- ✅ Pagination support
- ✅ Filtering by resource and action
- ✅ Search functionality
- ✅ Type-safe TypeScript interfaces
- ✅ Consistent error handling
- ✅ Automatic auth header injection
- ✅ Automatic token refresh

**Usage Example**:
```typescript
import { listPermissions, createPermission, updatePermission, deletePermission } from '@/api';

// List permissions
const permissions = await listPermissions({ 
  page: 1, 
  page_size: 50,
  resource: 'user'
});

// Create permission
const newPermission = await createPermission({
  name: 'user:RESET_PASSWORD',
  description: 'Ability to reset user passwords',
  resource: 'user',
  action: 'RESET_PASSWORD'
});

// Update permission
const updated = await updatePermission(permissionId, {
  description: 'Updated description'
});

// Delete permission
await deletePermission(permissionId);
```

---

### 5. API Index Exports (src/api/index.ts)

**Purpose**: Central export point for all API services and types

**Exports**:
- API Client: `apiClient`, `handleApiError`, `ApiError`
- Auth API: `login`, `refreshToken`, `logout`, `me`
- User API: All user functions and types
- School API: All school functions and types
- Role API: All role functions and types
- Permission API: All permission functions and types

**Benefits**:
- ✅ Single import point for all API services
- ✅ Type exports available for component typing
- ✅ Consistent export structure
- ✅ Easy to discover available APIs

**Usage Example**:
```typescript
// Import all needed APIs from single source
import {
  listUsers,
  listSchools,
  listRoles,
  listPermissions,
  type User,
  type School,
  type Role,
  type Permission,
  handleApiError
} from '@/api';
```

---

## Files Created

| File | Lines | Purpose |
|------|-------|---------|
| `src/api/users.ts` | 130 | User management API service |
| `src/api/schools.ts` | 145 | School management API service |
| `src/api/roles.ts` | 107 | Role management API service |
| `src/api/permissions.ts` | 113 | Permission management API service |

**Total**: 4 files, 495 lines of code

---

## Files Modified

| File | Lines Changed | Purpose |
|------|---------------|---------|
| `src/api/index.ts` | +70 | Add exports for all new services |

**Total**: 1 file modified, 70 lines added

---

## API Endpoints

### User Endpoints
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/users` | List users with pagination and filters |
| GET | `/api/v1/users/{id}` | Get user by ID |
| POST | `/api/v1/users` | Create new user |
| PUT | `/api/v1/users/{id}` | Update user |
| PATCH | `/api/v1/users/{id}/status` | Update user status |
| DELETE | `/api/v1/users/{id}` | Delete user |

### School Endpoints
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/schools` | List schools with pagination and filters |
| GET | `/api/v1/schools/{id}` | Get school by ID |
| POST | `/api/v1/schools` | Create new school |
| PUT | `/api/v1/schools/{id}` | Update school |
| PATCH | `/api/v1/schools/{id}/status` | Update school status |
| DELETE | `/api/v1/schools/{id}` | Delete school |

### Role Endpoints
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/roles` | List roles with pagination and filters |
| GET | `/api/v1/roles/{id}` | Get role by ID |
| POST | `/api/v1/roles` | Create new role |
| PUT | `/api/v1/roles/{id}` | Update role |
| DELETE | `/api/v1/roles/{id}` | Delete role |

### Permission Endpoints
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/permissions` | List permissions with pagination and filters |
| GET | `/api/v1/permissions/{id}` | Get permission by ID |
| POST | `/api/v1/permissions` | Create new permission |
| PUT | `/api/v1/permissions/{id}` | Update permission |
| DELETE | `/api/v1/permissions/{id}` | Delete permission |

---

## Type Safety

All API services include comprehensive TypeScript types:

### Request Types
- `CreateUserRequest` - User creation payload
- `UpdateUserRequest` - User update payload
- `CreateSchoolRequest` - School creation payload
- `UpdateSchoolRequest` - School update payload
- `CreateRoleRequest` - Role creation payload
- `UpdateRoleRequest` - Role update payload
- `CreatePermissionRequest` - Permission creation payload
- `UpdatePermissionRequest` - Permission update payload

### Response Types
- `User` - User entity
- `School` - School entity
- `Role` - Role entity
- `Permission` - Permission entity
- `ListUsersResponse` - Paginated user list
- `ListSchoolsResponse` - Paginated school list
- `ListRolesResponse` - Paginated role list
- `ListPermissionsResponse` - Paginated permission list

### Query Parameter Types
- `ListUsersParams` - User list filters
- `ListSchoolsParams` - School list filters
- `ListRolesParams` - Role list filters
- `ListPermissionsParams` - Permission list filters

---

## Error Handling

All API services use consistent error handling via `handleApiError`:

```typescript
try {
  const users = await listUsers();
} catch (error) {
  if (error instanceof ApiError) {
    console.error(error.message); // User-friendly message
    console.error(error.status);  // HTTP status code
    console.error(error.code);    // Error code if available
  }
}
```

**Error Codes**:
- 401 - Unauthorized (token expired or invalid)
- 403 - Forbidden (insufficient permissions)
- 404 - Not Found (resource doesn't exist)
- 409 - Conflict (duplicate email, code, etc.)
- 422 - Unprocessable Entity (validation error)
- 500 - Server Error

---

## Pagination

All list endpoints support pagination:

```typescript
const result = await listUsers({
  page: 1,
  page_size: 20
});

console.log(result.users);      // Array of users for current page
console.log(result.total);      // Total number of users
console.log(result.page);       // Current page number
console.log(result.page_size);  // Items per page
```

**Pagination Best Practices**:
- Start with page 1
- Use page_size between 10-100
- Display total to user
- Handle empty results gracefully
- Cache results when appropriate

---

## Filtering

All list endpoints support filtering:

```typescript
// Filter users by school
const users = await listUsers({ school_id: 'school-uuid' });

// Filter users by role
const users = await listUsers({ role_id: 'role-uuid' });

// Filter users by status
const users = await listUsers({ is_active: true });

// Filter schools by status
const schools = await listSchools({ is_active: true });

// Filter roles (non-system only)
const roles = await listRoles({ is_system: false });

// Filter permissions by resource
const permissions = await listPermissions({ resource: 'user' });
```

---

## Search

All list endpoints support search:

```typescript
// Search users by name or email
const users = await listUsers({ search: 'john' });

// Search schools by name or code
const schools = await listSchools({ search: 'nusa' });

// Search roles by name
const roles = await listRoles({ search: 'admin' });

// Search permissions by name
const permissions = await listPermissions({ search: 'user' });
```

---

## Backend Integration

All API services are designed to work with the backend API structure established in PHASE P0:

### Authentication
- All requests include `Authorization: Bearer <token>` header
- Handled automatically by axios interceptors (from P2)
- Token refresh on 401 errors handled automatically

### Authorization
- All requests protected by backend middleware
- Role and permission checks performed by backend
- 403 responses for insufficient permissions

### Validation
- Backend validates all request payloads
- Returns 422 for validation errors
- Error messages included in response

---

## Usage Patterns

### React Hook Pattern

```typescript
import { useState, useEffect } from 'react';
import { listUsers, type User } from '@/api';

function UserList() {
  const [users, setUsers] = useState<User[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    async function loadUsers() {
      try {
        setLoading(true);
        const result = await listUsers({ page: 1, page_size: 20 });
        setUsers(result.users);
      } catch (err) {
        setError('Failed to load users');
      } finally {
        setLoading(false);
      }
    }
    loadUsers();
  }, []);

  if (loading) return <Loading />;
  if (error) return <Error message={error} />;

  return <UserTable users={users} />;
}
```

### Form Submission Pattern

```typescript
import { createUser } from '@/api';

function CreateUserForm() {
  const handleSubmit = async (values: CreateUserRequest) => {
    try {
      const newUser = await createUser(values);
      // Show success message
      // Navigate to user list or detail page
    } catch (error) {
      // Show error message
    }
  };
  // ... form JSX
}
```

---

## Performance Considerations

### Caching
- Consider caching list results
- Use React Query or SWR for data fetching
- Implement optimistic updates for better UX

### Pagination
- Use appropriate page sizes (10-100)
- Lazy load large datasets
- Implement virtual scrolling for very long lists

### Debouncing
- Debounce search input
- Debounce filter changes
- Reduce unnecessary API calls

---

## Security Considerations

### Permission Checks
- Frontend should hide/disable actions based on permissions
- Use RoleGuard and PermissionGuard components (from P3)
- Example:
  ```tsx
  <PermissionGuard permission="user:DELETE">
    <Button onClick={() => deleteUser(id)}>Delete</Button>
  </PermissionGuard>
  ```

### Sensitive Data
- Passwords never returned in list or detail responses
- Only include necessary fields in responses
- Log out on 401 errors automatically

### CSRF Protection
- Backend should implement CSRF protection
- Include CSRF token in POST/PUT/DELETE requests if needed

---

## Testing Considerations

### Unit Tests Needed
- API service functions
- Error handling
- Type safety
- Pagination logic

### Integration Tests Needed
- API service with backend
- Authentication flow
- Authorization flow
- CRUD operations

### E2E Tests Needed
- User management workflow
- School management workflow
- Role management workflow
- Permission management workflow

---

## Exit Criteria for P4

- ✅ User API service created
- ✅ School API service created
- ✅ Role API service created
- ✅ Permission API service created
- ✅ All services have full CRUD operations
- ✅ All services have pagination support
- ✅ All services have filtering support
- ✅ All services have search support
- ✅ TypeScript interfaces complete
- ✅ Error handling implemented
- ✅ API index exports configured
- ✅ Type exports configured

**P4 Status**: ✅ COMPLETE

---

## Next Steps

Proceed to **FINAL VALIDATION** to:
- Execute System Admin end-to-end scenario
- Execute School Admin end-to-end scenario
- Execute Teacher end-to-end scenario
- Generate Sprint 2 Final Acceptance Report

---

## Appendix

### File Structure
```
src/api/
├── client.ts              # Axios client (from P2)
├── auth.ts                # Auth API (from P2)
├── users.ts               # User API service (NEW)
├── schools.ts             # School API service (NEW)
├── roles.ts               # Role API service (NEW)
├── permissions.ts         # Permission API service (NEW)
└── index.ts              # Public API exports (UPDATED)
```

### Dependencies
- axios (HTTP client)
- React 18+
- TypeScript 5+

### Backend Requirements
- Backend API endpoints must be operational (validated in P0)
- JWT authentication middleware must be active
- RBAC middleware must be active
- CORS must be configured

### Environment Variables
- `VITE_API_BASE_URL` - Backend API base URL (from P2)

---

## Known Limitations

1. **UI Components**: This phase created API services only; UI components for user/school/role/permission management are not implemented
2. **Form Validation**: Client-side form validation not implemented (should be added in UI components)
3. **Bulk Operations**: Bulk create/update/delete operations not supported
4. **Advanced Filtering**: Complex filtering (e.g., date ranges, multiple values) not supported
5. **Sorting**: API services don't include sort parameters (add if needed)
6. **File Uploads**: File upload operations not supported (e.g., school logo)

---

**Report Generated**: 2026-06-07
**Generated By**: Devin AI Agent (Principal Software Architect)
**Phase Status**: ✅ PHASE P4 COMPLETE
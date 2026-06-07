/**
 * Auth Feature Exports
 */

// Context & Provider
export { AuthProvider, useAuthContext } from './auth-context';
export { default as AuthContext } from './auth-context';

// Hooks
export { useAuth } from './use-auth';
export { default as useAuthDefault } from './use-auth';

// Storage
export { AuthStorage } from './auth-storage';

// Types
export type { Role, User, LoginCredentials, RegisterCredentials, AuthResponse, AuthState, AuthContextValue } from './types';

// Components
export { ProtectedRoute } from './protected-route';
export { default as ProtectedRouteDefault } from './protected-route';

export { PermissionGuard } from './permission-guard';
export { default as PermissionGuardDefault } from './permission-guard';

export { RoleGuard } from './role-guard';
export { default as RoleGuardDefault } from './role-guard';
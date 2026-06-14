/**
 * useAuth Hook
 * Convenient hook for accessing authentication state and methods
 */

import { useAuthContext } from './auth-context';
import { LoginCredentials } from './types';

export const useAuth = () => {
  const {
    isAuthenticated,
    user,
    permissions,
    loading,
    error,
    login,
    logout,
    refresh,
    me,
    hasPermission,
    hasRole,
  } = useAuthContext();

  return {
    // State
    isAuthenticated,
    user,
    permissions,
    loading,
    error,
    
    // Methods
    login: (credentials: LoginCredentials) => login(credentials),
    logout: () => logout(),
    refresh: () => refresh(),
    me: () => me(),
    hasPermission: (permission: string) => hasPermission(permission),
    hasRole: (role: string) => hasRole(role),
    
    // Helpers
    isSystemAdmin: user?.role_name === 'SYSTEM_ADMIN',
    isSchoolAdmin: user?.role_name === 'SCHOOL_ADMIN',
    isTeacher: user?.role_name === 'TEACHER',
    canCreateUsers: hasPermission('user:CREATE'),
    canUpdateUsers: hasPermission('user:UPDATE'),
    canDeleteUsers: hasPermission('user:DELETE'),
    canCreateSchools: hasPermission('school:CREATE'),
    canUpdateSchools: hasPermission('school:UPDATE'),
    canDeleteSchools: hasPermission('school:DELETE'),
  };
};

export default useAuth;
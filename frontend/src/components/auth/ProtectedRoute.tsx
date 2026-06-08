/**
 * ProtectedRoute Component
 * Wraps routes with role-based access control
 */

import React from 'react';
import { Navigate } from 'react-router-dom';
import { useAuthContext } from '@/features/auth';

interface ProtectedRouteProps {
  children: React.ReactNode;
  allowedRoles?: string[];
}

export const ProtectedRoute: React.FC<ProtectedRouteProps> = ({
  children,
  allowedRoles,
}) => {
  const { user, isAuthenticated } = useAuthContext();

  // If not authenticated, redirect to login
  if (!isAuthenticated) {
    return <Navigate to="/" replace />;
  }

  // If no roles are specified, allow access to all authenticated users
  if (!allowedRoles || allowedRoles.length === 0) {
    return <>{children}</>;
  }

  // Check if user's role is in the allowed roles
  if (user?.role_name && allowedRoles.includes(user.role_name)) {
    return <>{children}</>;
  }

  // Redirect to dashboard if not authorized
  return <Navigate to="/dashboard" replace />;
};

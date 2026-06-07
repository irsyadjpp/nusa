/**
 * Role Guard Component
 * Conditionally renders children based on user role
 */

import React from 'react';
import { useAuth } from './use-auth';

interface RoleGuardProps {
  allowedRoles: string[];
  fallback?: React.ReactNode;
  children: React.ReactNode;
}

export const RoleGuard: React.FC<RoleGuardProps> = ({
  allowedRoles,
  fallback,
  children,
}) => {
  const { hasRole } = useAuth();

  if (allowedRoles.some(role => hasRole(role))) {
    return <>{children}</>;
  }

  return <>{fallback || null}</>;
};

export default RoleGuard;
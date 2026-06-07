/**
 * Permission Guard Component
 * Conditionally renders children based on user permissions
 */

import React from 'react';
import { useAuth } from './use-auth';

interface PermissionGuardProps {
  permission: string;
  fallback?: React.ReactNode;
  children: React.ReactNode;
}

export const PermissionGuard: React.FC<PermissionGuardProps> = ({
  permission,
  fallback,
  children,
}) => {
  const { hasPermission } = useAuth();

  if (hasPermission(permission)) {
    return <>{children}</>;
  }

  return <>{fallback || null}</>;
};

export default PermissionGuard;
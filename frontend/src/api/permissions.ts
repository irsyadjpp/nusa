/**
 * Permission API Service
 * API calls for permission management operations
 */

import apiClient, { handleApiError } from './client';

// TypeScript interfaces
export interface Permission {
  id: string;
  name: string;
  description?: string;
  resource: string;
  action: string;
  created_at: string;
  updated_at: string;
}

export interface CreatePermissionRequest {
  name: string;
  description?: string;
  resource: string;
  action: string;
}

export interface UpdatePermissionRequest {
  name?: string;
  description?: string;
  resource?: string;
  action?: string;
}

export interface ListPermissionsResponse {
  permissions: Permission[];
  total: number;
  page: number;
  page_size: number;
}

export interface ListPermissionsParams {
  page?: number;
  page_size?: number;
  resource?: string;
  action?: string;
  search?: string;
}

/**
 * Get all permissions with optional filters
 */
export const listPermissions = async (params?: ListPermissionsParams): Promise<ListPermissionsResponse> => {
  try {
    const response = await apiClient.get('/permissions', { params });
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get permission by ID
 */
export const getPermission = async (id: string): Promise<Permission> => {
  try {
    const response = await apiClient.get(`/permissions/${id}`);
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Create new permission
 */
export const createPermission = async (data: CreatePermissionRequest): Promise<Permission> => {
  try {
    const response = await apiClient.post('/permissions', data);
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Update permission
 */
export const updatePermission = async (id: string, data: UpdatePermissionRequest): Promise<Permission> => {
  try {
    const response = await apiClient.put(`/permissions/${id}`, data);
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Delete permission
 */
export const deletePermission = async (id: string): Promise<void> => {
  try {
    await apiClient.delete(`/permissions/${id}`);
  } catch (error) {
    throw handleApiError(error);
  }
};

export default {
  listPermissions,
  getPermission,
  createPermission,
  updatePermission,
  deletePermission,
};
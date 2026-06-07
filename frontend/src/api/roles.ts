/**
 * Role API Service
 * API calls for role management operations
 */

import apiClient, { handleApiError } from './client';

// TypeScript interfaces
export interface Role {
  id: string;
  name: string;
  description?: string;
  is_system?: boolean;
  created_at: string;
  updated_at: string;
}

export interface CreateRoleRequest {
  name: string;
  description?: string;
}

export interface UpdateRoleRequest {
  name?: string;
  description?: string;
}

export interface ListRolesResponse {
  roles: Role[];
  total: number;
  page: number;
  page_size: number;
}

export interface ListRolesParams {
  page?: number;
  page_size?: number;
  is_system?: boolean;
  search?: string;
}

/**
 * Get all roles with optional filters
 */
export const listRoles = async (params?: ListRolesParams): Promise<ListRolesResponse> => {
  try {
    const response = await apiClient.get('/roles', { params });
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get role by ID
 */
export const getRole = async (id: string): Promise<Role> => {
  try {
    const response = await apiClient.get(`/roles/${id}`);
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Create new role
 */
export const createRole = async (data: CreateRoleRequest): Promise<Role> => {
  try {
    const response = await apiClient.post('/roles', data);
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Update role
 */
export const updateRole = async (id: string, data: UpdateRoleRequest): Promise<Role> => {
  try {
    const response = await apiClient.put(`/roles/${id}`, data);
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Delete role
 */
export const deleteRole = async (id: string): Promise<void> => {
  try {
    await apiClient.delete(`/roles/${id}`);
  } catch (error) {
    throw handleApiError(error);
  }
};

export default {
  listRoles,
  getRole,
  createRole,
  updateRole,
  deleteRole,
};
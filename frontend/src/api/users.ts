/**
 * User API Service
 * API calls for user management operations
 */

import apiClient, { handleApiError } from './client';

// TypeScript interfaces
export interface User {
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

export interface CreateUserRequest {
  email: string;
  name: string;
  password: string;
  role_id?: string;
  school_id?: string;
}

export interface UpdateUserRequest {
  name?: string;
  role_id?: string;
  school_id?: string;
  is_active?: boolean;
}

export interface ListUsersResponse {
  users: User[];
  total: number;
  page: number;
  page_size: number;
}

export interface ListUsersParams {
  page?: number;
  page_size?: number;
  school_id?: string;
  role_id?: string;
  is_active?: boolean;
  search?: string;
}

/**
 * Get all users with optional filters
 */
export const listUsers = async (params?: ListUsersParams): Promise<ListUsersResponse> => {
  try {
    const response = await apiClient.get('/users', { params });
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get user by ID
 */
export const getUser = async (id: string): Promise<User> => {
  try {
    const response = await apiClient.get(`/users/${id}`);
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Create new user
 */
export const createUser = async (data: CreateUserRequest): Promise<User> => {
  try {
    const response = await apiClient.post('/users', data);
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Update user
 */
export const updateUser = async (id: string, data: UpdateUserRequest): Promise<User> => {
  try {
    const response = await apiClient.put(`/users/${id}`, data);
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Update user status (activate/deactivate)
 */
export const updateUserStatus = async (id: string, is_active: boolean): Promise<User> => {
  try {
    const response = await apiClient.patch(`/users/${id}/status`, { is_active });
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Delete user
 */
export const deleteUser = async (id: string): Promise<void> => {
  try {
    await apiClient.delete(`/users/${id}`);
  } catch (error) {
    throw handleApiError(error);
  }
};

export default {
  listUsers,
  getUser,
  createUser,
  updateUser,
  updateUserStatus,
  deleteUser,
};
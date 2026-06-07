/**
 * Auth API Service
 * API calls for authentication operations
 */

import apiClient, { handleApiError, ApiError } from './client';
import { LoginCredentials, AuthResponse, User } from '@/features/auth/types';

/**
 * Login user
 */
export const login = async (credentials: LoginCredentials): Promise<AuthResponse> => {
  try {
    const response = await apiClient.post('/public/auth/login', credentials);
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Refresh access token
 */
export const refreshToken = async (refreshToken: string): Promise<{
  access_token: string;
  refresh_token: string;
  token_type: string;
  expires_in: number;
}> => {
  try {
    const response = await apiClient.post('/public/auth/refresh', {
      refresh_token: refreshToken,
    });
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Logout user
 */
export const logout = async (refreshToken: string): Promise<void> => {
  try {
    await apiClient.post('/auth/logout', {
      refresh_token: refreshToken,
    });
  } catch (error) {
    // Don't throw error on logout - always clear local storage
    console.error('Logout error:', handleApiError(error).message);
  }
};

/**
 * Get current user
 */
export const me = async (): Promise<{
  user: User;
  role_name: string;
  permissions: string[];
}> => {
  try {
    const response = await apiClient.get('/auth/me');
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

export default {
  login,
  refreshToken,
  logout,
  me,
};
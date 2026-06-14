/**
 * API Client Configuration
 * Axios instance with interceptors, auth header injection, and token refresh mechanism
 */

import axios, { AxiosInstance, AxiosError, InternalAxiosRequestConfig } from 'axios';
import { AuthStorage } from '@/features/auth/auth-storage';

// API base URL - should be configured via environment variable
const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8081';

/**
 * Create configured axios instance
 */
const createApiClient = (): AxiosInstance => {
  const client = axios.create({
    baseURL: `${API_BASE_URL}/api/v1`,
    timeout: 10000, // 10 seconds
    headers: {
      'Content-Type': 'application/json',
    },
  });

  // Request interceptor - Add auth header
  client.interceptors.request.use(
    (config) => {
      const token = AuthStorage.getAccessToken();
      if (token) {
        config.headers.Authorization = `Bearer ${token}`;
      }
      return config;
    },
    (error) => {
      return Promise.reject(error);
    }
  );

  // Response interceptor - Handle errors and token refresh
  client.interceptors.response.use(
    (response) => response,
    async (error: AxiosError) => {
      const originalRequest = error.config as InternalAxiosRequestConfig & {
        _retry?: boolean;
      };

      // Handle 401 errors - try token refresh
      if (error.response?.status === 401 && !originalRequest._retry) {
        originalRequest._retry = true;

        try {
          const refreshToken = AuthStorage.getRefreshToken();
          if (!refreshToken) {
            // No refresh token - logout user
            AuthStorage.clear();
            window.location.href = '/';
            return Promise.reject(error);
          }

          // Refresh token
          const response = await axios.post(`${API_BASE_URL}/api/v1/public/auth/refresh`, {
            refresh_token: refreshToken,
          });

          const { access_token, refresh_token } = response.data.data;

          // Store new tokens
          AuthStorage.setAccessToken(access_token);
          AuthStorage.setRefreshToken(refresh_token);

          // Retry original request with new token
          originalRequest.headers.Authorization = `Bearer ${access_token}`;
          return client(originalRequest);
        } catch (refreshError) {
          // Refresh failed - logout user
          AuthStorage.clear();
          window.location.href = '/';
          return Promise.reject(refreshError);
        }
      }

      // Handle other errors
      return Promise.reject(error);
    }
  );

  return client;
};

// Export singleton instance
export const apiClient = createApiClient();

/**
 * API Error class
 */
export class ApiError extends Error {
  constructor(
    public message: string,
    public status?: number,
    public code?: string
  ) {
    super(message);
    this.name = 'ApiError';
  }
}

/**
 * Handle API error responses
 */
export const handleApiError = (error: unknown): ApiError => {
  if (axios.isAxiosError(error)) {
    const status = error.response?.status;
    const data = error.response?.data as any;

    if (data?.error) {
      return new ApiError(data.error, status, data.code);
    }

    if (status === 401) {
      return new ApiError('Unauthorized - Please log in again', status);
    }

    if (status === 403) {
      return new ApiError('Forbidden - You do not have permission', status);
    }

    if (status === 404) {
      return new ApiError('Resource not found', status);
    }

    if (status === 500) {
      return new ApiError('Server error - Please try again later', status);
    }

    return new ApiError('An unexpected error occurred', status);
  }

  if (error instanceof Error) {
    return new ApiError(error.message);
  }

  return new ApiError('An unknown error occurred');
};

export default apiClient;
/**
 * School API Service
 * API calls for school management operations
 */

import apiClient, { handleApiError } from './client';

// TypeScript interfaces
export interface School {
  id: string;
  name: string;
  code: string;
  address?: string;
  city?: string;
  state?: string;
  country?: string;
  postal_code?: string;
  phone?: string;
  email?: string;
  website?: string;
  is_active: boolean;
  created_at: string;
  updated_at: string;
}

export interface CreateSchoolRequest {
  name: string;
  code: string;
  address?: string;
  city?: string;
  state?: string;
  country?: string;
  postal_code?: string;
  phone?: string;
  email?: string;
  website?: string;
}

export interface UpdateSchoolRequest {
  name?: string;
  code?: string;
  address?: string;
  city?: string;
  state?: string;
  country?: string;
  postal_code?: string;
  phone?: string;
  email?: string;
  website?: string;
  is_active?: boolean;
}

export interface ListSchoolsResponse {
  schools: School[];
  total: number;
  page: number;
  page_size: number;
}

export interface ListSchoolsParams {
  page?: number;
  page_size?: number;
  is_active?: boolean;
  search?: string;
}

/**
 * Get all schools with optional filters
 */
export const listSchools = async (params?: ListSchoolsParams): Promise<ListSchoolsResponse> => {
  try {
    const response = await apiClient.get('/schools', { params });
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get school by ID
 */
export const getSchool = async (id: string): Promise<School> => {
  try {
    const response = await apiClient.get(`/schools/${id}`);
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Create new school
 */
export const createSchool = async (data: CreateSchoolRequest): Promise<School> => {
  try {
    const response = await apiClient.post('/schools', data);
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Update school
 */
export const updateSchool = async (id: string, data: UpdateSchoolRequest): Promise<School> => {
  try {
    const response = await apiClient.put(`/schools/${id}`, data);
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Update school status (activate/deactivate)
 */
export const updateSchoolStatus = async (id: string, is_active: boolean): Promise<School> => {
  try {
    const response = await apiClient.patch(`/schools/${id}/status`, { is_active });
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Delete school
 */
export const deleteSchool = async (id: string): Promise<void> => {
  try {
    await apiClient.delete(`/schools/${id}`);
  } catch (error) {
    throw handleApiError(error);
  }
};

export default {
  listSchools,
  getSchool,
  createSchool,
  updateSchool,
  updateSchoolStatus,
  deleteSchool,
};
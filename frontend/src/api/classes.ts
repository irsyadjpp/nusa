/**
 * Class API Service
 * API calls for class management operations
 */

import apiClient, { handleApiError } from './client';

// TypeScript interfaces
export interface Class {
  id: string;
  name: string;
  code: string;
  grade_level: string;
  academic_year_id: string;
  semester_id: string;
  homeroom_teacher_id?: string;
  school_id: string;
  is_active: boolean;
  created_at: string;
  updated_at: string;
}

export interface CreateClassRequest {
  name: string;
  code: string;
  grade_level: string;
  academic_year_id: string;
  semester_id: string;
  homeroom_teacher_id?: string;
  school_id: string;
}

export interface UpdateClassRequest {
  name?: string;
  code?: string;
  grade_level?: string;
  academic_year_id?: string;
  semester_id?: string;
  homeroom_teacher_id?: string;
  is_active?: boolean;
}

export interface ListClassesResponse {
  classes: Class[];
  total: number;
  page: number;
  page_size: number;
}

export interface ListClassesParams {
  page?: number;
  page_size?: number;
  school_id?: string;
  academic_year_id?: string;
  semester_id?: string;
  grade_level?: string;
  is_active?: boolean;
  search?: string;
}

/**
 * Get all classes with optional filters
 */
export const listClasses = async (params?: ListClassesParams): Promise<ListClassesResponse> => {
  try {
    const response = await apiClient.get('/classes', { params });
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get class by ID
 */
export const getClass = async (id: string): Promise<Class> => {
  try {
    const response = await apiClient.get(`/classes/${id}`);
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Create new class
 */
export const createClass = async (data: CreateClassRequest): Promise<Class> => {
  try {
    const response = await apiClient.post('/classes', data);
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Update class
 */
export const updateClass = async (id: string, data: UpdateClassRequest): Promise<Class> => {
  try {
    const response = await apiClient.put(`/classes/${id}`, data);
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Delete class (soft delete)
 */
export const deleteClass = async (id: string): Promise<void> => {
  try {
    await apiClient.delete(`/classes/${id}`);
  } catch (error) {
    throw handleApiError(error);
  }
};

export default {
  listClasses,
  getClass,
  createClass,
  updateClass,
  deleteClass,
};

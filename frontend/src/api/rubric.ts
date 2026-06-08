/**
 * Rubric API Client
 * Handles all Rubric-related API calls
 */

import apiClient, { handleApiError } from './client';

// Types
export interface Rubric {
  id: string;
  user_id: string;
  rubric_type: string;
  title: string;
  description: string;
  criteria: any;
  max_score: number;
  performance_levels: any;
  ai_generated: boolean;
  ai_confidence_score?: number;
  ai_agent_version?: string;
  created_at: string;
  updated_at: string;
}

export interface RubricResponse {
  id: string;
  user_id: string;
  user_name: string;
  rubric_type: string;
  title: string;
  description: string;
  criteria: any;
  max_score: number;
  performance_levels: any;
  ai_generated: boolean;
  ai_confidence_score?: number;
  ai_agent_version?: string;
  created_at: string;
  updated_at: string;
}

export interface CreateRubricRequest {
  rubric_type: string;
  title: string;
  description: string;
  criteria: any;
  max_score: number;
  performance_levels: any;
}

export interface UpdateRubricRequest {
  title?: string;
  description?: string;
  criteria?: any;
  max_score?: number;
  performance_levels?: any;
}

/**
 * Get all rubrics with optional filters
 */
export const getRubrics = async (params?: {
  user_id?: string;
  rubric_type?: string;
  limit?: number;
  offset?: number;
}): Promise<Rubric[]> => {
  try {
    const response = await apiClient.get('/rubrics', { params });
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get rubric by ID
 */
export const getRubricById = async (id: string): Promise<RubricResponse> => {
  try {
    const response = await apiClient.get(`/rubrics/${id}`);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Create new rubric
 */
export const createRubric = async (data: CreateRubricRequest, userId: string): Promise<Rubric> => {
  try {
    const response = await apiClient.post('/rubrics', {
      ...data,
      user_id: userId,
    });
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Update rubric
 */
export const updateRubric = async (id: string, data: UpdateRubricRequest): Promise<Rubric> => {
  try {
    const response = await apiClient.put(`/rubrics/${id}`, data);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Delete rubric
 */
export const deleteRubric = async (id: string): Promise<void> => {
  try {
    await apiClient.delete(`/rubrics/${id}`);
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get rubrics by type
 */
export const getRubricsByType = async (rubricType: string, params?: {
  limit?: number;
  offset?: number;
}): Promise<Rubric[]> => {
  try {
    const response = await apiClient.get(`/rubrics/type/${rubricType}`, { params });
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

export default {
  getRubrics,
  getRubricById,
  createRubric,
  updateRubric,
  deleteRubric,
  getRubricsByType,
};

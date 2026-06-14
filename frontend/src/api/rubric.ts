/**
 * Rubric API Client
 * Handles all Rubric-related API calls with proper types
 */

import apiClient, { handleApiError } from './client';
import {
  Rubric,
  RubricType,
  RubricCriteria,
  AssessmentStatus,
  PaginationParams,
  FilterParams
} from '@/shared/types/domain';

// API-specific request types
export interface RubricCreateRequest {
  assessment_id?: string;
  rubric_type: RubricType;
  rubric_criteria: RubricCriteria[];
  total_points: number;
  status?: AssessmentStatus;
}

export interface RubricUpdateRequest {
  rubric_criteria?: RubricCriteria[];
  total_points?: number;
  status?: AssessmentStatus;
}

/**
 * Get all rubrics with optional filters
 */
export const getRubrics = async (params?: PaginationParams & FilterParams & {
  user_id?: string;
  assessment_id?: string;
  rubric_type?: RubricType;
}): Promise<Rubric[]> => {
  try {
    const response = await apiClient.get('/assessment/rubrics', { params });
    return response.data.rubrics || response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get rubric by ID
 */
export const getRubricById = async (id: string): Promise<Rubric> => {
  try {
    const response = await apiClient.get(`/assessment/rubrics/${id}`);
    return response.data || response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Create new rubric
 */
export const createRubric = async (data: RubricCreateRequest, userId: string): Promise<Rubric> => {
  try {
    const response = await apiClient.post('assessment/rubrics', {
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
export const updateRubric = async (id: string, data: RubricUpdateRequest): Promise<Rubric> => {
  try {
    const response = await apiClient.put(`assessment/rubrics/${id}`, data);
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
    await apiClient.delete(`assessment/rubrics/${id}`);
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get rubrics by type
 */
export const getRubricsByType = async (rubricType: RubricType, params?: PaginationParams): Promise<Rubric[]> => {
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

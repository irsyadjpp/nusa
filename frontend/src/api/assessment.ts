/**
 * Assessment API Client
 * Handles all Assessment-related API calls with proper types
 */

import apiClient, { handleApiError } from './client';
import {
  Assessment,
  AssessmentType,
  CreateAssessmentRequest,
  UpdateAssessmentRequest,
  PaginationParams,
  FilterParams
} from '@/shared/types/domain';

/**
 * Get all assessments with optional filters
 */
export const getAssessments = async (params?: PaginationParams & FilterParams & {
  tp_id?: string;
  user_id?: string;
  assessment_type?: AssessmentType;
}): Promise<Assessment[]> => {
  try {
    const response = await apiClient.get('/assessment', { params });
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get assessment by ID
 */
export const getAssessmentById = async (id: string): Promise<Assessment> => {
  try {
    const response = await apiClient.get(`/assessment/${id}`);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Create new assessment
 */
export const createAssessment = async (data: CreateAssessmentRequest, userId: string): Promise<Assessment> => {
  try {
    const response = await apiClient.post('/assessment', {
      ...data,
      user_id: userId,
    });
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Update assessment
 */
export const updateAssessment = async (id: string, data: UpdateAssessmentRequest): Promise<Assessment> => {
  try {
    const response = await apiClient.put(`/assessment/${id}`, data);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Delete assessment
 */
export const deleteAssessment = async (id: string): Promise<void> => {
  try {
    await apiClient.delete(`/assessment/${id}`);
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Approve assessment
 */
export const approveAssessment = async (id: string, userId: string): Promise<Assessment> => {
  try {
    const response = await apiClient.post(`/assessment/${id}/approve`, {
      approved_by: userId,
    });
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Reject assessment
 */
export const rejectAssessment = async (id: string, userId: string): Promise<Assessment> => {
  try {
    const response = await apiClient.put(`/assessment/${id}/reject`, {
      rejected_by: userId,
    });
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

export default {
  getAssessments,
  getAssessmentById,
  createAssessment,
  updateAssessment,
  deleteAssessment,
  approveAssessment,
  rejectAssessment,
};

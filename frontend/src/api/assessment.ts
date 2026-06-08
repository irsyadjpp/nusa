/**
 * Assessment API Client
 * Handles all Assessment-related API calls
 */

import apiClient, { handleApiError } from './client';

// Types
export interface Assessment {
  id: string;
  tp_id: string;
  tp_version_no: number;
  success_criteria_snapshot: any;
  user_id: string;
  assessment_type: string;
  status: string;
  assessment_items: any;
  answer_key: any;
  scoring_guidelines: any;
  ai_confidence_score?: number;
  ai_generated_at?: string;
  ai_agent_version?: string;
  version_no: number;
  is_current_version: boolean;
  parent_version_id?: string;
  created_at: string;
  updated_at: string;
  approved_at?: string;
  approved_by?: string;
}

export interface AssessmentResponse {
  id: string;
  tp_id: string;
  tp_title: string;
  tp_version_no: number;
  success_criteria_snapshot: any;
  user_id: string;
  user_name: string;
  assessment_type: string;
  status: string;
  assessment_items: any;
  answer_key: any;
  scoring_guidelines: any;
  ai_confidence_score?: number;
  ai_generated_at?: string;
  ai_agent_version?: string;
  version_no: number;
  is_current_version: boolean;
  parent_version_id?: string;
  created_at: string;
  updated_at: string;
  approved_at?: string;
  approved_by?: string;
}

export interface CreateAssessmentRequest {
  tp_id: string;
  tp_version_no: number;
  success_criteria_snapshot: any;
  assessment_type: string;
  assessment_items: any;
  answer_key: any;
  scoring_guidelines: any;
}

export interface UpdateAssessmentRequest {
  assessment_type?: string;
  assessment_items?: any;
  answer_key?: any;
  scoring_guidelines?: any;
  status?: string;
}

/**
 * Get all assessments with optional filters
 */
export const getAssessments = async (params?: {
  tp_id?: string;
  user_id?: string;
  assessment_type?: string;
  status?: string;
  limit?: number;
  offset?: number;
}): Promise<Assessment[]> => {
  try {
    const response = await apiClient.get('/assessments', { params });
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get assessment by ID
 */
export const getAssessmentById = async (id: string): Promise<AssessmentResponse> => {
  try {
    const response = await apiClient.get(`/assessments/${id}`);
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
    const response = await apiClient.post('/assessments', {
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
    const response = await apiClient.put(`/assessments/${id}`, data);
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
    await apiClient.delete(`/assessments/${id}`);
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Approve assessment
 */
export const approveAssessment = async (id: string, userId: string): Promise<Assessment> => {
  try {
    const response = await apiClient.put(`/assessments/${id}/approve`, {
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
    const response = await apiClient.put(`/assessments/${id}/reject`, {
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

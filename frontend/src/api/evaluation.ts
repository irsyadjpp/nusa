/**
 * Evaluation API Client
 * Handles all Evaluation-related API calls
 */

import apiClient, { handleApiError } from './client';

// Types
export interface Evaluation {
  id: string;
  student_id: string;
  rubric_id: string;
  evidence_id: string;
  user_id: string;
  performance_scores: any;
  total_score: number;
  max_score: number;
  performance_level: string;
  teacher_feedback?: string;
  revision_no: number;
  evaluated_at: string;
  created_at: string;
  updated_at: string;
}

export interface EvaluationResponse {
  id: string;
  student_id: string;
  student_name: string;
  rubric_id: string;
  rubric_type: string;
  evidence_id: string;
  evidence_type: string;
  user_id: string;
  user_name: string;
  performance_scores: any;
  total_score: number;
  max_score: number;
  performance_level: string;
  teacher_feedback?: string;
  revision_no: number;
  evaluated_at: string;
  created_at: string;
  updated_at: string;
}

export interface CreateEvaluationRequest {
  student_id: string;
  rubric_id: string;
  evidence_id: string;
  performance_scores: any;
  total_score: number;
  max_score: number;
  performance_level: string;
  teacher_feedback?: string;
}

export interface UpdateEvaluationRequest {
  performance_scores?: any;
  total_score?: number;
  max_score?: number;
  performance_level?: string;
  teacher_feedback?: string;
}

/**
 * Get all evaluations with optional filters
 */
export const getEvaluations = async (params?: {
  student_id?: string;
  rubric_id?: string;
  evidence_id?: string;
  user_id?: string;
  limit?: number;
  offset?: number;
}): Promise<Evaluation[]> => {
  try {
    const response = await apiClient.get('/evaluations', { params });
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get evaluation by ID
 */
export const getEvaluationById = async (id: string): Promise<EvaluationResponse> => {
  try {
    const response = await apiClient.get(`/evaluations/${id}`);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Create new evaluation
 */
export const createEvaluation = async (data: CreateEvaluationRequest, userId: string): Promise<Evaluation> => {
  try {
    const response = await apiClient.post('/evaluations', {
      ...data,
      user_id: userId,
    });
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Update evaluation
 */
export const updateEvaluation = async (id: string, data: UpdateEvaluationRequest): Promise<Evaluation> => {
  try {
    const response = await apiClient.put(`/evaluations/${id}`, data);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Delete evaluation
 */
export const deleteEvaluation = async (id: string): Promise<void> => {
  try {
    await apiClient.delete(`/evaluations/${id}`);
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get evaluations by student ID
 */
export const getEvaluationsByStudent = async (studentId: string, params?: {
  rubric_id?: string;
  limit?: number;
  offset?: number;
}): Promise<Evaluation[]> => {
  try {
    const response = await apiClient.get(`/students/${studentId}/evaluations`, { params });
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get evaluations by evidence ID
 */
export const getEvaluationsByEvidence = async (evidenceId: string, params?: {
  limit?: number;
  offset?: number;
}): Promise<Evaluation[]> => {
  try {
    const response = await apiClient.get(`/evidences/${evidenceId}/evaluations`, { params });
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get evaluation history for evidence
 */
export const getEvaluationHistory = async (evidenceId: string): Promise<Evaluation[]> => {
  try {
    const response = await apiClient.get(`/evidences/${evidenceId}/evaluations/history`);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

export default {
  getEvaluations,
  getEvaluationById,
  createEvaluation,
  updateEvaluation,
  deleteEvaluation,
  getEvaluationsByStudent,
  getEvaluationsByEvidence,
  getEvaluationHistory,
};

/**
 * Evaluation API Client
 * Handles all Evaluation-related API calls with proper types
 */

import apiClient, { handleApiError } from './client';
import {
  Evaluation,
  PerformanceScores,
  MasteryLevel,
  CreateEvaluationRequest,
  PaginationParams,
  FilterParams
} from '@/shared/types/domain';

// API-specific request types
export interface EvaluationUpdateRequest {
  performance_scores?: PerformanceScores;
  performance_level?: MasteryLevel;
  teacher_feedback?: string;
}

/**
 * Get all evaluations with optional filters
 */
export const getEvaluations = async (params?: PaginationParams & FilterParams & {
  student_id?: string;
  rubric_id?: string;
  evidence_id?: string;
  user_id?: string;
}): Promise<Evaluation[]> => {
  try {
    const response = await apiClient.get('assessment/evaluations', { params });
    return response.data.evaluations || response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get evaluation by ID
 */
export const getEvaluationById = async (id: string): Promise<Evaluation> => {
  try {
    const response = await apiClient.get(`assessment/evaluations/${id}`);
    return response.data || response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Create new evaluation
 */
export const createEvaluation = async (data: CreateEvaluationRequest, userId: string): Promise<Evaluation> => {
  try {
    const response = await apiClient.post('assessment/evaluations', {
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
export const updateEvaluation = async (id: string, data: EvaluationUpdateRequest): Promise<Evaluation> => {
  try {
    const response = await apiClient.put(`assessment/evaluations/${id}`, data);
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
    await apiClient.delete(`assessment/evaluations/${id}`);
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get evaluations by student ID
 */
export const getEvaluationsByStudent = async (studentId: string, params?: PaginationParams & FilterParams & {
  rubric_id?: string;
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
export const getEvaluationsByEvidence = async (evidenceId: string, params?: PaginationParams): Promise<Evaluation[]> => {
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

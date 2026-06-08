/**
 * Evidence API Client
 * Handles all Evidence-related API calls
 */

import apiClient, { handleApiError } from './client';

// Types
export interface Evidence {
  id: string;
  student_id: string;
  assessment_id: string;
  user_id: string;
  evidence_type: string;
  status: string;
  evidence_data: any;
  teacher_notes?: string;
  rubric_id?: string;
  linked_criteria?: any;
  evaluations: any[];
  created_at: string;
  updated_at: string;
}

export interface EvidenceResponse {
  id: string;
  student_id: string;
  student_name: string;
  assessment_id: string;
  assessment_type: string;
  user_id: string;
  user_name: string;
  evidence_type: string;
  status: string;
  evidence_data: any;
  teacher_notes?: string;
  rubric_id?: string;
  linked_criteria?: any;
  evaluation_notes?: string;
  created_at: string;
  updated_at: string;
}

export interface CreateEvidenceRequest {
  student_id: string;
  assessment_id: string;
  evidence_type: string;
  evidence_data: any;
  teacher_notes?: string;
  rubric_id?: string;
  linked_criteria?: any;
}

export interface UpdateEvidenceRequest {
  evidence_type?: string;
  evidence_data?: any;
  teacher_notes?: string;
  rubric_id?: string;
  linked_criteria?: any;
  status?: string;
}

/**
 * Get all evidences with optional filters
 */
export const getEvidences = async (params?: {
  student_id?: string;
  assessment_id?: string;
  user_id?: string;
  evidence_type?: string;
  status?: string;
  limit?: number;
  offset?: number;
}): Promise<Evidence[]> => {
  try {
    const response = await apiClient.get('/evidences', { params });
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get evidence by ID
 */
export const getEvidenceById = async (id: string): Promise<EvidenceResponse> => {
  try {
    const response = await apiClient.get(`/evidences/${id}`);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Create new evidence
 */
export const createEvidence = async (data: CreateEvidenceRequest, userId: string): Promise<Evidence> => {
  try {
    const response = await apiClient.post('/evidences', {
      ...data,
      user_id: userId,
    });
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Update evidence
 */
export const updateEvidence = async (id: string, data: UpdateEvidenceRequest): Promise<Evidence> => {
  try {
    const response = await apiClient.put(`/evidences/${id}`, data);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Delete evidence
 */
export const deleteEvidence = async (id: string): Promise<void> => {
  try {
    await apiClient.delete(`/evidences/${id}`);
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get evidences by student ID
 */
export const getEvidencesByStudent = async (studentId: string, params?: {
  assessment_id?: string;
  status?: string;
  limit?: number;
  offset?: number;
}): Promise<Evidence[]> => {
  try {
    const response = await apiClient.get(`/students/${studentId}/evidences`, { params });
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get evidences by assessment ID
 */
export const getEvidencesByAssessment = async (assessmentId: string, params?: {
  status?: string;
  limit?: number;
  offset?: number;
}): Promise<Evidence[]> => {
  try {
    const response = await apiClient.get(`/assessments/${assessmentId}/evidences`, { params });
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

export default {
  getEvidences,
  getEvidenceById,
  createEvidence,
  updateEvidence,
  deleteEvidence,
  getEvidencesByStudent,
  getEvidencesByAssessment,
};

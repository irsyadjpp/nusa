/**
 * Evidence API Client
 * Handles all Evidence-related API calls with proper types
 */

import apiClient, { handleApiError } from './client';
import {
  Evidence,
  EvidenceType,
  EvidenceStatus,
  FileMetadata,
  CreateEvidenceRequest,
  PaginationParams,
  FilterParams
} from '@/shared/types/domain';

// API-specific request types
export interface EvidenceUpdateRequest {
  evidence_type?: EvidenceType;
  evidence_data?: any;
  teacher_notes?: string;
  rubric_id?: string;
  linked_criteria?: any;
  status?: EvidenceStatus;
}

export interface UploadEvidenceRequest {
  student_id: string;
  assessment_id: string;
  evidence_type: EvidenceType;
  file_url: string;
  file_metadata: FileMetadata;
}

/**
 * Get all evidences with optional filters
 */
export const getEvidences = async (params?: PaginationParams & FilterParams & {
  student_id?: string;
  assessment_id?: string;
  user_id?: string;
  evidence_type?: EvidenceType;
}): Promise<Evidence[]> => {
  try {
    const response = await apiClient.get('/assessment/evidences', { params });
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get evidence by ID
 */
export const getEvidenceById = async (id: string): Promise<Evidence> => {
  try {
    const response = await apiClient.get(`/assessment/evidences/${id}`);
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
    const response = await apiClient.post('/assessment/evidences', {
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
export const updateEvidence = async (id: string, data: EvidenceUpdateRequest): Promise<Evidence> => {
  try {
    const response = await apiClient.put(`/assessment/evidences/${id}`, data);
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
    await apiClient.delete(`/assessment/evidences/${id}`);
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get evidences by student ID
 */
export const getEvidencesByStudent = async (studentId: string, params?: PaginationParams & FilterParams & {
  assessment_id?: string;
  status?: EvidenceStatus;
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
export const getEvidencesByAssessment = async (assessmentId: string, params?: PaginationParams & FilterParams & {
  status?: EvidenceStatus;
}): Promise<Evidence[]> => {
  try {
    const response = await apiClient.get(`/assessment/${assessmentId}/evidences`, { params });
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Upload evidence with metadata
 */
export const uploadEvidence = async (data: UploadEvidenceRequest, userId: string): Promise<Evidence> => {
  try {
    const response = await apiClient.post(`/assessment/evidences/upload`, {
      ...data,
      user_id: userId,
    });
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
  uploadEvidence,
};

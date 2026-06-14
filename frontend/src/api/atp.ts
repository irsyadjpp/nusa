/**
 * ATP (Alur Tujuan Pembelajaran) API Client
 * Handles all ATP-related API calls with proper types
 */

import apiClient, { handleApiError } from './client';
import {
  ATP,
  ATPSet,
  LearningActivities,
  TimeAllocation,
  TPStatus,
  PaginationParams,
  FilterParams
} from '@/shared/types/domain';

// Extended request types for API specific needs
export interface ATPCreateRequest {
  atp_set_id: string;
  tp_id: string;
  sequence_number: number;
  week: number;
  learning_activities: LearningActivities;
  assessment_methods: string[];
  time_allocation: TimeAllocation;
}

export interface ATPUpdateRequest {
  sequence_number?: number;
  week?: number;
  learning_activities?: LearningActivities;
  assessment_methods?: string[];
  time_allocation?: TimeAllocation;
  status?: TPStatus;
}

export interface ATPSetCreateRequest {
  tp_set_id: string;
  subject_id: string;
  phase_id: string;
  grade: string;
  semester: string;
  generation_source?: 'MANUAL' | 'AI_GENERATED';
  generation_reason?: string;
}

export interface ATPSetUpdateRequest {
  status?: TPStatus;
}

/**
 * Get all ATPs with optional filters
 */
export const getATPs = async (params?: PaginationParams & FilterParams & {
  atp_set_id?: string;
  tp_id?: string;
}): Promise<ATP[]> => {
  try {
    const response = await apiClient.get('learning-planning/atps', { params });
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get ATP by ID
 */
export const getATPById = async (id: string): Promise<ATP> => {
  try {
    const response = await apiClient.get(`learning-planning/atps/${id}`);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get ATPs by ATP Set ID
 */
export const getATPsBySet = async (atpSetId: string): Promise<ATP[]> => {
  try {
    const response = await apiClient.get(`learning-planning/atps/atp-set/${atpSetId}`);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Create new ATP
 */
export const createATP = async (data: ATPCreateRequest): Promise<ATP> => {
  try {
    const response = await apiClient.post('learning-planning/atps', data);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Update ATP
 */
export const updateATP = async (id: string, data: ATPUpdateRequest): Promise<ATP> => {
  try {
    const response = await apiClient.put(`learning-planning/atps/${id}`, data);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Delete ATP
 */
export const deleteATP = async (id: string): Promise<void> => {
  try {
    await apiClient.delete(`learning-planning/atps/${id}`);
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get ATP Sets
 */
export const getATPSets = async (params?: PaginationParams & FilterParams & {
  tp_set_id?: string;
  subject_id?: string;
  phase_id?: string;
}): Promise<ATPSet[]> => {
  try {
    const response = await apiClient.get('learning-planning/atp-sets', { params });
    return response.data.atp_sets || response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get ATP Set by ID (Sprint 3.5 EP-07)
 */
export const getATPSetById = async (id: string): Promise<ATPSet> => {
  try {
    const response = await apiClient.get(`learning-planning/atp-sets/${id}`);
    return response.data || response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Create ATP Set
 */
export const createATPSet = async (data: ATPSetCreateRequest): Promise<ATPSet> => {
  try {
    const response = await apiClient.post('learning-planning/atp-sets', data);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Update ATP Set
 */
export const updateATPSet = async (id: string, data: ATPSetUpdateRequest): Promise<ATPSet> => {
  try {
    const response = await apiClient.put(`learning-planning/atp-sets/${id}`, data);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Approve ATP Set
 */
export const approveATPSet = async (id: string): Promise<ATPSet> => {
  try {
    const response = await apiClient.post(`learning-planning/atp-sets/${id}/approve`);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Delete ATP Set
 */
export const deleteATPSet = async (id: string): Promise<void> => {
  try {
    await apiClient.delete(`learning-planning/atp-sets/${id}`);
  } catch (error) {
    throw handleApiError(error);
  }
};

export default {
  getATPs,
  getATPById,
  getATPsBySet,
  createATP,
  updateATP,
  deleteATP,
  getATPSets,
  getATPSetById,
  createATPSet,
  updateATPSet,
  approveATPSet,
  deleteATPSet,
};

/**
 * TP (Teaching Plan) API Client
 * Handles all TP-related API calls with proper types
 */

import apiClient, { handleApiError } from './client';
import {
  TP,
  TPSet,
  CreateTPRequest,
  UpdateTPRequest,
  CreateTPSetRequest,
  PaginationParams,
  FilterParams
} from '@/shared/types/domain';

// Re-export types for convenience
export type { TP, TPSet, CreateTPRequest, UpdateTPRequest, CreateTPSetRequest } from '@/shared/types/domain';

/**
 * Get all TPs with optional filters
 */
export const getTPs = async (params?: PaginationParams & FilterParams & {
  tp_set_id?: string;
}): Promise<TP[]> => {
  try {
    const response = await apiClient.get('/tp', { params });
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get TP by ID
 */
export const getTPById = async (id: string): Promise<TP> => {
  try {
    const response = await apiClient.get(`/tp/${id}`);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get TPs by TP Set ID
 */
export const getTPsBySet = async (tpSetId: string): Promise<TP[]> => {
  try {
    const response = await apiClient.get(`/tp/tp-set/${tpSetId}`);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Create new TP
 */
export const createTP = async (data: CreateTPRequest): Promise<TP> => {
  try {
    const response = await apiClient.post('/tp', data);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Update TP
 */
export const updateTP = async (id: string, data: UpdateTPRequest): Promise<TP> => {
  try {
    const response = await apiClient.put(`/tp/${id}`, data);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Delete TP
 */
export const deleteTP = async (id: string): Promise<void> => {
  try {
    await apiClient.delete(`/tp/${id}`);
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get TP Sets
 */
export const getTPSets = async (params?: PaginationParams & FilterParams & {
  cp_id?: string;
}): Promise<TPSet[]> => {
  try {
    const response = await apiClient.get('/learning-planning/tp-sets', { params });
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get TP Set by ID
 */
export const getTPSetById = async (id: string): Promise<TPSet> => {
  try {
    const response = await apiClient.get(`/learning-planning/tp-sets/${id}`);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Create TP Set
 */
export const createTPSet = async (data: CreateTPSetRequest): Promise<TPSet> => {
  try {
    const response = await apiClient.post('/learning-planning/tp-sets', data);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Approve TP Set
 */
export const approveTPSet = async (id: string): Promise<TPSet> => {
  try {
    const response = await apiClient.post(`/learning-planning/tp-sets/${id}/approve`);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Update TP Set
 */
export const updateTPSet = async (id: string, data: {
  generation_reason?: string;
}): Promise<TPSet> => {
  try {
    const response = await apiClient.put(`/learning-planning/tp-sets/${id}`, data);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get TP Set Version History
 */
export const getTPSetVersions = async (id: string): Promise<TP[]> => {
  try {
    const response = await apiClient.get(`/learning-planning/tp-sets/${id}/versions`);
    return response.data.versions || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

export default {
  getTPs,
  getTPById,
  getTPsBySet,
  createTP,
  updateTP,
  deleteTP,
  getTPSets,
  getTPSetById,
  createTPSet,
  approveTPSet,
  updateTPSet,
  getTPSetVersions,
};

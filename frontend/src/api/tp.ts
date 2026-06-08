/**
 * TP (Teaching Plan) API Client
 * Handles all TP-related API calls
 */

import apiClient, { handleApiError } from './client';

// Types
export interface TP {
  id: string;
  tp_set_id: string;
  sequence_number: number;
  cp_id: string;
  subject_id: string;
  phase_id: string;
  element_id: string;
  subelement_id: string;
  user_id: string;
  status: string;
  title: string;
  learning_objectives: any;
  time_allocation: any;
  prerequisites: any;
  estimated_weeks: number;
  success_criteria: any;
  created_at: string;
  updated_at: string;
}

export interface TPSet {
  id: string;
  cp_id: string;
  version_no: number;
  status: string;
  generation_source: string;
  generation_reason?: string;
  generated_by: string;
  ai_generation_id?: string;
  approved_by?: string;
  approved_at?: string;
  created_at: string;
}

export interface CreateTPRequest {
  tp_set_id: string;
  sequence_number: number;
  cp_id: string;
  subject_id: string;
  phase_id: string;
  element_id: string;
  subelement_id: string;
  title: string;
  learning_objectives: any;
  time_allocation: any;
  prerequisites: any;
  estimated_weeks: number;
  success_criteria?: any;
}

export interface UpdateTPRequest {
  title?: string;
  learning_objectives?: any;
  time_allocation?: any;
  prerequisites?: any;
  estimated_weeks?: number;
  status?: string;
  success_criteria?: any;
}

export interface TPResponse {
  id: string;
  tp_set_id: string;
  sequence_number: number;
  cp_id: string;
  subject_id: string;
  phase_id: string;
  element_id: string;
  subelement_id: string;
  user_id: string;
  status: string;
  title: string;
  learning_objectives: any;
  time_allocation: any;
  prerequisites: any;
  estimated_weeks: number;
  success_criteria: any;
  created_at: string;
  updated_at: string;
}

/**
 * Get all TPs with optional filters
 */
export const getTPs = async (params?: {
  tp_set_id?: string;
  subject_id?: string;
  phase_id?: string;
  status?: string;
  limit?: number;
  offset?: number;
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
export const getTPSets = async (params?: {
  cp_id?: string;
  subject_id?: string;
  phase_id?: string;
  status?: string;
  limit?: number;
  offset?: number;
}): Promise<TPSet[]> => {
  try {
    const response = await apiClient.get('/tp-sets', { params });
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
    const response = await apiClient.get(`/tp-sets/${id}`);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Create TP Set
 */
export const createTPSet = async (data: {
  cp_id: string;
  subject_id: string;
  phase_id: string;
  generation_source: string;
  generation_reason?: string;
}): Promise<TPSet> => {
  try {
    const response = await apiClient.post('/tp-sets', data);
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
    const response = await apiClient.post(`/tp-sets/${id}/approve`);
    return response.data.data || response.data;
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
};

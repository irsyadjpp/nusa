/**
 * ATP (Alur Tujuan Pembelajaran) API Client
 * Handles all ATP-related API calls
 */

import apiClient, { handleApiError } from './client';

// Types
export interface ATP {
  id: string;
  atp_set_id: string;
  tp_id: string;
  sequence_number: number;
  week_number: number;
  estimated_hours: number;
  status: string;
  created_at: string;
  updated_at: string;
}

export interface ATPSet {
  id: string;
  tp_set_id: string;
  subject_id: string;
  phase_id: string;
  grade: string;
  semester: string;
  status: string;
  approved_by?: string;
  approved_at?: string;
  created_by: string;
  created_at: string;
  updated_at: string;
}

export interface CreateATPRequest {
  atp_set_id: string;
  tp_id: string;
  sequence_number: number;
  week_number: number;
  estimated_hours: number;
}

export interface UpdateATPRequest {
  sequence_number?: number;
  week_number?: number;
  estimated_hours?: number;
  status?: string;
}

export interface CreateATPSetRequest {
  tp_set_id: string;
  subject_id: string;
  phase_id: string;
  grade: string;
  semester: string;
}

export interface UpdateATPSetRequest {
  status?: string;
}

/**
 * Get all ATPs with optional filters
 */
export const getATPs = async (params?: {
  atp_set_id?: string;
  tp_id?: string;
  status?: string;
  limit?: number;
  offset?: number;
}): Promise<ATP[]> => {
  try {
    const response = await apiClient.get('/learning-planning/atps', { params });
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
    const response = await apiClient.get(`/learning-planning/atps/${id}`);
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
    const response = await apiClient.get(`/learning-planning/atps/atp-set/${atpSetId}`);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Create new ATP
 */
export const createATP = async (data: CreateATPRequest): Promise<ATP> => {
  try {
    const response = await apiClient.post('/learning-planning/atps', data);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Update ATP
 */
export const updateATP = async (id: string, data: UpdateATPRequest): Promise<ATP> => {
  try {
    const response = await apiClient.put(`/learning-planning/atps/${id}`, data);
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
    await apiClient.delete(`/learning-planning/atps/${id}`);
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get ATP Sets
 */
export const getATPSets = async (params?: {
  tp_set_id?: string;
  subject_id?: string;
  phase_id?: string;
  status?: string;
  limit?: number;
  offset?: number;
}): Promise<ATPSet[]> => {
  try {
    const response = await apiClient.get('/learning-planning/atp-sets', { params });
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get ATP Set by ID
 */
export const getATPSetById = async (id: string): Promise<ATPSet> => {
  try {
    const response = await apiClient.get(`/learning-planning/atp-sets/${id}`);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Create ATP Set
 */
export const createATPSet = async (data: CreateATPSetRequest): Promise<ATPSet> => {
  try {
    const response = await apiClient.post('/learning-planning/atp-sets', data);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Update ATP Set
 */
export const updateATPSet = async (id: string, data: UpdateATPSetRequest): Promise<ATPSet> => {
  try {
    const response = await apiClient.put(`/learning-planning/atp-sets/${id}`, data);
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
    const response = await apiClient.post(`/learning-planning/atp-sets/${id}/approve`);
    return response.data.data || response.data;
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
};

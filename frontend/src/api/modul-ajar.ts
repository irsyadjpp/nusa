/**
 * Modul Ajar API Client
 * Handles all Modul Ajar-related API calls
 */

import apiClient, { handleApiError } from './client';

// Types
export interface ModulAjar {
  id: string;
  modul_ajar_set_id: string;
  atp_id: string;
  tp_id: string;
  sequence_number: number;
  title: string;
  learning_activities: any;
  teaching_methods: string[];
  learning_media: string[];
  learning_resources: string[];
  attachments: any[];
  status: string;
  created_by: string;
  created_at: string;
  updated_at: string;
}

export interface ModulAjarSet {
  id: string;
  atp_set_id: string;
  subject_id: string;
  phase_id: string;
  grade: string;
  semester: string;
  status: string;
  created_by: string;
  created_at: string;
  updated_at: string;
}

export interface CreateModulAjarRequest {
  modul_ajar_set_id: string;
  atp_id: string;
  tp_id: string;
  sequence_number: number;
  title: string;
  learning_activities: any;
  teaching_methods: string[];
  learning_media: string[];
  learning_resources: string[];
  attachments?: any[];
}

export interface UpdateModulAjarRequest {
  title?: string;
  learning_activities?: any;
  teaching_methods?: string[];
  learning_media?: string[];
  learning_resources?: string[];
  attachments?: any[];
  status?: string;
}

export interface CreateModulAjarSetRequest {
  atp_set_id: string;
  subject_id: string;
  phase_id: string;
  grade: string;
  semester: string;
}

export interface UpdateModulAjarSetRequest {
  status?: string;
}

/**
 * Get all Modul Ajar with optional filters
 */
export const getModulAjars = async (params?: {
  modul_ajar_set_id?: string;
  atp_id?: string;
  tp_id?: string;
  status?: string;
  limit?: number;
  offset?: number;
}): Promise<ModulAjar[]> => {
  try {
    const response = await apiClient.get('/learning-planning/modul-ajar', { params });
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get Modul Ajar by ID
 */
export const getModulAjarById = async (id: string): Promise<ModulAjar> => {
  try {
    const response = await apiClient.get(`/learning-planning/modul-ajar/${id}`);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get Modul Ajars by Modul Ajar Set ID
 */
export const getModulAjarsBySet = async (modulAjarSetId: string): Promise<ModulAjar[]> => {
  try {
    const response = await apiClient.get(`/learning-planning/modul-ajar/modul-ajar-set/${modulAjarSetId}`);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Create new Modul Ajar
 */
export const createModulAjar = async (data: CreateModulAjarRequest): Promise<ModulAjar> => {
  try {
    const response = await apiClient.post('/learning-planning/modul-ajar', data);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Update Modul Ajar
 */
export const updateModulAjar = async (id: string, data: UpdateModulAjarRequest): Promise<ModulAjar> => {
  try {
    const response = await apiClient.put(`/learning-planning/modul-ajar/${id}`, data);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Delete Modul Ajar
 */
export const deleteModulAjar = async (id: string): Promise<void> => {
  try {
    await apiClient.delete(`/learning-planning/modul-ajar/${id}`);
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get Modul Ajar Sets
 */
export const getModulAjarSets = async (params?: {
  atp_set_id?: string;
  subject_id?: string;
  phase_id?: string;
  status?: string;
  limit?: number;
  offset?: number;
}): Promise<ModulAjarSet[]> => {
  try {
    const response = await apiClient.get('/learning-planning/modul-ajar-sets', { params });
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get Modul Ajar Set by ID
 */
export const getModulAjarSetById = async (id: string): Promise<ModulAjarSet> => {
  try {
    const response = await apiClient.get(`/learning-planning/modul-ajar-sets/${id}`);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Create Modul Ajar Set
 */
export const createModulAjarSet = async (data: CreateModulAjarSetRequest): Promise<ModulAjarSet> => {
  try {
    const response = await apiClient.post('/learning-planning/modul-ajar-sets', data);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Update Modul Ajar Set
 */
export const updateModulAjarSet = async (id: string, data: UpdateModulAjarSetRequest): Promise<ModulAjarSet> => {
  try {
    const response = await apiClient.put(`/learning-planning/modul-ajar-sets/${id}`, data);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

export default {
  getModulAjars,
  getModulAjarById,
  getModulAjarsBySet,
  createModulAjar,
  updateModulAjar,
  deleteModulAjar,
  getModulAjarSets,
  getModulAjarSetById,
  createModulAjarSet,
  updateModulAjarSet,
};

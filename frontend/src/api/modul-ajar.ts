/**
 * Modul Ajar API Client
 * Handles all Modul Ajar-related API calls with proper types
 */

import apiClient, { handleApiError } from './client';
import {
  ModulAjar,
  ModulAjarSet,
  TeachingMaterials,
  TimeAllocation,
  TPStatus,
  LearningActivities,
  PaginationParams,
  FilterParams
} from '@/shared/types/domain';

// API-specific request types
export interface ModulAjarCreateRequest {
  modul_ajar_set_id: string;
  atp_id: string;
  week: number;
  session_number: number;
  learning_objectives: string[];
  learning_activities?: LearningActivities; // Convenience property matching component expectations
  teaching_materials: TeachingMaterials;
  learning_methods: string[];
  assessment_methods: string[];
  time_allocation: TimeAllocation;
}

export interface ModulAjarUpdateRequest {
  learning_objectives?: string[];
  teaching_materials?: TeachingMaterials;
  learning_methods?: string[];
  assessment_methods?: string[];
  time_allocation?: TimeAllocation;
  status?: TPStatus;
}

export interface ModulAjarSetCreateRequest {
  atp_set_id: string;
  subject_id: string;
  phase_id: string;
  generation_source?: 'MANUAL' | 'AI_GENERATED';
  generation_reason?: string;
}

export interface ModulAjarSetUpdateRequest {
  status?: TPStatus;
}

/**
 * Get all Modul Ajar with optional filters
 */
export const getModulAjars = async (params?: PaginationParams & FilterParams & {
  modul_ajar_set_id?: string;
  atp_id?: string;
}): Promise<ModulAjar[]> => {
  try {
    const response = await apiClient.get('learning-planning/modul-ajar', { params });
    return response.data.modul_ajar || response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get Modul Ajar by ID
 */
export const getModulAjarById = async (id: string): Promise<ModulAjar> => {
  try {
    const response = await apiClient.get(`learning-planning/modul-ajar/${id}`);
    return response.data || response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get Modul Ajars by Modul Ajar Set ID
 */
export const getModulAjarsBySet = async (modulAjarSetId: string): Promise<ModulAjar[]> => {
  try {
    const response = await apiClient.get(`learning-planning/modul-ajar/modul-ajar-set/${modulAjarSetId}`);
    return response.data.modul_ajar || response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Create new Modul Ajar
 */
export const createModulAjar = async (data: ModulAjarCreateRequest): Promise<ModulAjar> => {
  try {
    const response = await apiClient.post('learning-planning/modul-ajar', data);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Update Modul Ajar
 */
export const updateModulAjar = async (id: string, data: ModulAjarUpdateRequest): Promise<ModulAjar> => {
  try {
    const response = await apiClient.put(`learning-planning/modul-ajar/${id}`, data);
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
    await apiClient.delete(`learning-planning/modul-ajar/${id}`);
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get Modul Ajar Sets
 */
export const getModulAjarSets = async (params?: PaginationParams & FilterParams & {
  atp_set_id?: string;
  subject_id?: string;
  phase_id?: string;
}): Promise<ModulAjarSet[]> => {
  try {
    const response = await apiClient.get('learning-planning/modul-ajar-sets', { params });
    return response.data.modul_ajar_sets || response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get Modul Ajar Set by ID
 */
export const getModulAjarSetById = async (id: string): Promise<ModulAjarSet> => {
  try {
    const response = await apiClient.get(`learning-planning/modul-ajar-sets/${id}`);
    return response.data || response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Create Modul Ajar Set
 */
export const createModulAjarSet = async (data: ModulAjarSetCreateRequest): Promise<ModulAjarSet> => {
  try {
    const response = await apiClient.post('learning-planning/modul-ajar-sets', data);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Update Modul Ajar Set
 */
export const updateModulAjarSet = async (id: string, data: ModulAjarSetUpdateRequest): Promise<ModulAjarSet> => {
  try {
    const response = await apiClient.put(`learning-planning/modul-ajar-sets/${id}`, data);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Approve Modul Ajar Set
 */
export const approveModulAjarSet = async (id: string): Promise<ModulAjarSet> => {
  try {
    const response = await apiClient.post(`learning-planning/modul-ajar-sets/${id}/approve`);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Delete Modul Ajar Set
 */
export const deleteModulAjarSet = async (id: string): Promise<void> => {
  try {
    await apiClient.delete(`learning-planning/modul-ajar-sets/${id}`);
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
  approveModulAjarSet,
  deleteModulAjarSet,
};

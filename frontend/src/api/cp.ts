/**
 * CP (Capaian Pembelajaran) API Client
 * Handles all CP-related API calls
 */

import apiClient, { handleApiError } from './client';

// Types
export interface CP {
  id: string;
  subject_id: string;
  phase_id: string;
  element_id: string;
  subelement_id?: string;
  description: string;
  code: string;
  status: string;
  created_at: string;
  updated_at: string;
}

export interface Subject {
  id: string;
  name: string;
  code: string;
  status: string;
  created_at: string;
  updated_at: string;
}

export interface Phase {
  id: string;
  name: string;
  level: string;
  grade_range: string;
  status: string;
  created_at: string;
  updated_at: string;
}

export interface Element {
  id: string;
  phase_id: string;
  name: string;
  code: string;
  status: string;
  created_at: string;
  updated_at: string;
}

export interface Subelement {
  id: string;
  element_id: string;
  name: string;
  code: string;
  status: string;
  created_at: string;
  updated_at: string;
}

/**
 * Get all subjects
 */
export const getSubjects = async (params?: {
  status?: string;
  limit?: number;
  offset?: number;
}): Promise<Subject[]> => {
  try {
    const response = await apiClient.get('/curriculum/subjects', { params });
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get subject by ID
 */
export const getSubjectById = async (id: string): Promise<Subject> => {
  try {
    const response = await apiClient.get(`/curriculum/subjects/${id}`);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get all phases
 */
export const getPhases = async (params?: {
  status?: string;
  limit?: number;
  offset?: number;
}): Promise<Phase[]> => {
  try {
    const response = await apiClient.get('/curriculum/phases', { params });
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get phase by ID
 */
export const getPhaseById = async (id: string): Promise<Phase> => {
  try {
    const response = await apiClient.get(`/curriculum/phases/${id}`);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get all CPs with optional filters
 */
export const getCPs = async (params?: {
  subject_id?: string;
  phase_id?: string;
  element_id?: string;
  status?: string;
  limit?: number;
  offset?: number;
}): Promise<CP[]> => {
  try {
    const response = await apiClient.get('/curriculum/cp', { params });
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get CP by ID
 */
export const getCPById = async (id: string): Promise<CP> => {
  try {
    const response = await apiClient.get(`/curriculum/cp/${id}`);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get elements by phase
 */
export const getElementsByPhase = async (phaseId: string): Promise<Element[]> => {
  try {
    const response = await apiClient.get(`/curriculum/elements`, {
      params: { phase_id: phaseId },
    });
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get element by ID
 */
export const getElementById = async (id: string): Promise<Element> => {
  try {
    const response = await apiClient.get(`/curriculum/elements/${id}`);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get subelements by element
 */
export const getSubelementsByElement = async (elementId: string): Promise<Subelement[]> => {
  try {
    const response = await apiClient.get(`/curriculum/subelements`, {
      params: { element_id: elementId },
    });
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get subelement by ID
 */
export const getSubelementById = async (id: string): Promise<Subelement> => {
  try {
    const response = await apiClient.get(`/curriculum/subelements/${id}`);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

export default {
  getSubjects,
  getSubjectById,
  getPhases,
  getPhaseById,
  getCPs,
  getCPById,
  getElementsByPhase,
  getElementById,
  getSubelementsByElement,
  getSubelementById,
};

/**
 * CP (Capaian Pembelajaran) API Client
 * Handles all CP-related API calls with proper types
 */

import apiClient, { handleApiError } from './client';
import {
  CP,
  CurriculumSubject,
  CurriculumPhase,
  CurriculumElement,
  CurriculumSubelement,
} from '@/shared/types/domain';

// Type alias for consistency
type Subelement = CurriculumSubelement;

// API-specific request types
export interface SubjectCreateRequest {
  code: string;
  name: string;
  description?: string;
}

export interface PhaseCreateRequest {
  code: string;
  name: string;
  description?: string;
  grade_level_start?: number;
  grade_level_end?: number;
}

export interface ElementCreateRequest {
  subject_id: string;
  phase_id: string;
  name: string;
  code: string;
  description?: string;
}

export interface SubelementCreateRequest {
  element_id: string;
  name: string;
  code: string;
  description?: string;
}

export interface CPCreateRequest {
  subject_id: string;
  phase_id: string;
  element_id: string;
  subelement_id?: string;
  code: string;
  cp_text: string;
  competency_code?: string;
  learning_objectives: any;
  competency_standards: any;
  time_allocation_hours: number;
  hours_per_week: number;
  version: string;
}

/**
 * Get all subjects
 */
export const getSubjects = async (): Promise<CurriculumSubject[]> => {
  try {
    const response = await apiClient.get('curriculum/subjects');
    const result = response.data.subjects || response.data.data || response.data;
    return Array.isArray(result) ? result : [];
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get subject by ID
 */
export const getSubjectById = async (id: string): Promise<CurriculumSubject> => {
  try {
    const response = await apiClient.get(`curriculum/subjects/${id}`);
    return response.data || response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Create a new subject
 */
export const createSubject = async (data: SubjectCreateRequest): Promise<CurriculumSubject> => {
  try {
    const response = await apiClient.post('curriculum/subjects', data);
    return response.data || response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Update a subject
 */
export const updateSubject = async (
  id: string,
  data: {
    name?: string;
    description?: string;
    is_active?: boolean;
  }
): Promise<CurriculumSubject> => {
  try {
    const response = await apiClient.put(`curriculum/subjects/${id}`, data);
    return response.data || response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Delete a subject
 */
export const deleteSubject = async (id: string): Promise<void> => {
  try {
    await apiClient.delete(`curriculum/subjects/${id}`);
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get all phases
 */
export const getPhases = async (): Promise<CurriculumPhase[]> => {
  try {
    const response = await apiClient.get('curriculum/phases');
    const result = response.data.phases || response.data.data || response.data;
    return Array.isArray(result) ? result : [];
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get phase by ID
 */
export const getPhaseById = async (id: string): Promise<CurriculumPhase> => {
  try {
    const response = await apiClient.get(`curriculum/phases/${id}`);
    return response.data || response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Create a new phase
 */
export const createPhase = async (data: PhaseCreateRequest): Promise<CurriculumPhase> => {
  try {
    const response = await apiClient.post('curriculum/phases', data);
    return response.data || response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Update a phase
 */
export const updatePhase = async (
  id: string,
  data: {
    name?: string;
    description?: string;
    grade_level_start?: number;
    grade_level_end?: number;
    is_active?: boolean;
  }
): Promise<CurriculumPhase> => {
  try {
    const response = await apiClient.put(`curriculum/phases/${id}`, data);
    return response.data || response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Delete a phase
 */
export const deletePhase = async (id: string): Promise<void> => {
  try {
    await apiClient.delete(`curriculum/phases/${id}`);
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
    const response = await apiClient.get('curriculum/cp', { params });
    const result = response.data.cps || response.data.data || response.data;
    return Array.isArray(result) ? result : [];
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get CP by ID
 */
export const getCPById = async (id: string): Promise<CP> => {
  try {
    const response = await apiClient.get(`curriculum/cp/${id}`);
    return response.data || response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Create a new CP
 */
export const createCP = async (data: {
  subject_id: string;
  phase_id: string;
  element_id: string;
  subelement_id: string;
  code: string;
  description: string;
  competency_code?: string;
  learning_objectives: any;
  competency_standards: any;
  time_allocation_hours: number;
  hours_per_week: number;
  version: string;
}): Promise<CP> => {
  try {
    const response = await apiClient.post('curriculum/cp', data);
    return response.data || response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Update a CP
 */
export const updateCP = async (
  id: string,
  data: {
    description?: string;
    competency_code?: string;
    learning_objectives?: any;
    competency_standards?: any;
    time_allocation_hours?: number;
    hours_per_week?: number;
    version?: string;
    is_active?: boolean;
  }
): Promise<CP> => {
  try {
    const response = await apiClient.put(`curriculum/cp/${id}`, data);
    return response.data || response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Delete a CP
 */
export const deleteCP = async (id: string): Promise<void> => {
  try {
    await apiClient.delete(`curriculum/cp/${id}`);
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get elements by phase
 */
export const getElementsByPhase = async (phaseId: string): Promise<CurriculumElement[]> => {
  try {
    const response = await apiClient.get(`curriculum/elements`, {
      params: { phase_id: phaseId },
    });
    const result = response.data.elements || response.data.data || response.data;
    return Array.isArray(result) ? result : [];
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get element by ID
 */
export const getElementById = async (id: string): Promise<CurriculumElement> => {
  try {
    const response = await apiClient.get(`curriculum/elements/${id}`);
    return response.data || response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Create a new element
 */
export const createElement = async (data: {
  subject_id: string;
  phase_id: string;
  code: string;
  name: string;
  description?: string;
}): Promise<CurriculumElement> => {
  try {
    const response = await apiClient.post('curriculum/elements', data);
    return response.data || response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Update an element
 */
export const updateElement = async (
  id: string,
  data: {
    name?: string;
    description?: string;
    is_active?: boolean;
  }
): Promise<CurriculumElement> => {
  try {
    const response = await apiClient.put(`curriculum/elements/${id}`, data);
    return response.data || response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Delete an element
 */
export const deleteElement = async (id: string): Promise<void> => {
  try {
    await apiClient.delete(`curriculum/elements/${id}`);
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get subelements by element
 */
export const getSubelementsByElement = async (elementId: string): Promise<Subelement[]> => {
  try {
    const response = await apiClient.get(`curriculum/subelements`, {
      params: { element_id: elementId },
    });
    const result = response.data.subelements || response.data.data || response.data;
    return Array.isArray(result) ? result : [];
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get subelement by ID
 */
export const getSubelementById = async (id: string): Promise<Subelement> => {
  try {
    const response = await apiClient.get(`curriculum/subelements/${id}`);
    return response.data || response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Create a new subelement
 */
export const createSubelement = async (data: {
  element_id: string;
  code: string;
  name: string;
  description?: string;
}): Promise<Subelement> => {
  try {
    const response = await apiClient.post('curriculum/subelements', data);
    return response.data || response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Update a subelement
 */
export const updateSubelement = async (
  id: string,
  data: {
    name?: string;
    description?: string;
    is_active?: boolean;
  }
): Promise<Subelement> => {
  try {
    const response = await apiClient.put(`curriculum/subelements/${id}`, data);
    return response.data || response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Delete a subelement
 */
export const deleteSubelement = async (id: string): Promise<void> => {
  try {
    await apiClient.delete(`curriculum/subelements/${id}`);
  } catch (error) {
    throw handleApiError(error);
  }
};

export default {
  getSubjects,
  getSubjectById,
  createSubject,
  updateSubject,
  deleteSubject,
  getPhases,
  getPhaseById,
  createPhase,
  updatePhase,
  deletePhase,
  getCPs,
  getCPById,
  createCP,
  updateCP,
  deleteCP,
  getElementsByPhase,
  getElementById,
  createElement,
  updateElement,
  deleteElement,
  getSubelementsByElement,
  getSubelementById,
  createSubelement,
  updateSubelement,
  deleteSubelement,
};

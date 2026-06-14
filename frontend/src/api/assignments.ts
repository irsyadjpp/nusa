/**
 * Assignment API Service
 * API calls for assignment management operations
 */

import apiClient, { handleApiError } from './client';

// TypeScript interfaces
export interface Assignment {
  id: string;
  class_id: string;
  assessment_id: string;
  title: string;
  description?: string;
  due_date: string;
  max_score: number;
  status: 'draft' | 'published' | 'closed';
  created_by: string;
  created_at: string;
  updated_at: string;
}

export interface CreateAssignmentRequest {
  class_id: string;
  assessment_id: string;
  title: string;
  description?: string;
  due_date: string;
  max_score: number;
}

export interface UpdateAssignmentRequest {
  title?: string;
  description?: string;
  due_date?: string;
  max_score?: number;
  status?: 'draft' | 'published' | 'closed';
}

export interface ListAssignmentsResponse {
  assignments: Assignment[];
  total: number;
  page: number;
  page_size: number;
}

export interface ListAssignmentsParams {
  page?: number;
  page_size?: number;
  class_id?: string;
  assessment_id?: string;
  status?: string;
  due_date_from?: string;
  due_date_to?: string;
}

/**
 * Create assignment
 */
export const createAssignment = async (data: CreateAssignmentRequest): Promise<Assignment> => {
  try {
    const response = await apiClient.post('/assignments', data);
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * List assignments with filters
 */
export const listAssignments = async (params?: ListAssignmentsParams): Promise<ListAssignmentsResponse> => {
  try {
    const response = await apiClient.get('/assignments', { params });
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get assignments by class
 */
export const getAssignmentsByClass = async (class_id: string): Promise<Assignment[]> => {
  try {
    const response = await apiClient.get(`/assignments/class/${class_id}`);
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Update assignment
 */
export const updateAssignment = async (id: string, data: UpdateAssignmentRequest): Promise<Assignment> => {
  try {
    const response = await apiClient.put(`/assignments/${id}`, data);
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Delete assignment (soft delete)
 */
export const deleteAssignment = async (id: string): Promise<void> => {
  try {
    await apiClient.delete(`/assignments/${id}`);
  } catch (error) {
    throw handleApiError(error);
  }
};

export default {
  createAssignment,
  listAssignments,
  getAssignmentsByClass,
  updateAssignment,
  deleteAssignment,
};

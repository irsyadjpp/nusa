/**
 * Exam API Service
 * API calls for exam management operations
 */

import apiClient, { handleApiError } from './client';

// TypeScript interfaces
export interface Exam {
  id: string;
  class_id: string;
  assessment_id: string;
  exam_date: string;
  start_time: string;
  duration_minutes: number;
  room?: string;
  status: 'scheduled' | 'in_progress' | 'completed' | 'cancelled';
  created_by: string;
  created_at: string;
  updated_at: string;
}

export interface CreateExamRequest {
  class_id: string;
  assessment_id: string;
  exam_date: string;
  start_time: string;
  duration_minutes: number;
  room?: string;
}

export interface UpdateExamRequest {
  exam_date?: string;
  start_time?: string;
  duration_minutes?: number;
  room?: string;
  status?: 'scheduled' | 'in_progress' | 'completed' | 'cancelled';
}

export interface ListExamsResponse {
  exams: Exam[];
  total: number;
  page: number;
  page_size: number;
}

export interface ListExamsParams {
  page?: number;
  page_size?: number;
  class_id?: string;
  assessment_id?: string;
  status?: string;
  date_from?: string;
  date_to?: string;
}

/**
 * Create exam
 */
export const createExam = async (data: CreateExamRequest): Promise<Exam> => {
  try {
    const response = await apiClient.post('/exams', data);
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * List exams with filters
 */
export const listExams = async (params?: ListExamsParams): Promise<ListExamsResponse> => {
  try {
    const response = await apiClient.get('/exams', { params });
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get exams by class
 */
export const getExamsByClass = async (class_id: string): Promise<Exam[]> => {
  try {
    const response = await apiClient.get(`/exams/class/${class_id}`);
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Update exam
 */
export const updateExam = async (id: string, data: UpdateExamRequest): Promise<Exam> => {
  try {
    const response = await apiClient.put(`/exams/${id}`, data);
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Delete exam (soft delete)
 */
export const deleteExam = async (id: string): Promise<void> => {
  try {
    await apiClient.delete(`/exams/${id}`);
  } catch (error) {
    throw handleApiError(error);
  }
};

export default {
  createExam,
  listExams,
  getExamsByClass,
  updateExam,
  deleteExam,
};

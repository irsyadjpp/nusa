/**
 * Exam Result API Service
 * API calls for exam result management operations
 */

import apiClient, { handleApiError } from './client';

// TypeScript interfaces
export interface ExamResult {
  id: string;
  exam_id: string;
  student_id: string;
  score: number;
  grade?: string;
  remarks?: string;
  graded_at?: string;
  graded_by?: string;
  created_at: string;
  updated_at: string;
}

export interface CreateExamResultRequest {
  exam_id: string;
  student_id: string;
  score: number;
  grade?: string;
  remarks?: string;
}

export interface UpdateExamResultRequest {
  score?: number;
  grade?: string;
  remarks?: string;
}

export interface ListExamResultsResponse {
  results: ExamResult[];
  total: number;
  page: number;
  page_size: number;
}

export interface ListExamResultsParams {
  page?: number;
  page_size?: number;
  exam_id?: string;
  student_id?: string;
  graded?: boolean;
}

/**
 * Create exam result
 */
export const createExamResult = async (data: CreateExamResultRequest): Promise<ExamResult> => {
  try {
    const response = await apiClient.post('/exam-results', data);
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * List exam results with filters
 */
export const listExamResults = async (params?: ListExamResultsParams): Promise<ListExamResultsResponse> => {
  try {
    const response = await apiClient.get('/exam-results', { params });
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get exam results by exam
 */
export const getExamResultsByExam = async (exam_id: string): Promise<ExamResult[]> => {
  try {
    const response = await apiClient.get(`/exam-results/exam/${exam_id}`);
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get exam results by student
 */
export const getExamResultsByStudent = async (student_id: string): Promise<ExamResult[]> => {
  try {
    const response = await apiClient.get(`/exam-results/student/${student_id}`);
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Update exam result
 */
export const updateExamResult = async (id: string, data: UpdateExamResultRequest): Promise<ExamResult> => {
  try {
    const response = await apiClient.put(`/exam-results/${id}`, data);
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Delete exam result (soft delete)
 */
export const deleteExamResult = async (id: string): Promise<void> => {
  try {
    await apiClient.delete(`/exam-results/${id}`);
  } catch (error) {
    throw handleApiError(error);
  }
};

export default {
  createExamResult,
  listExamResults,
  getExamResultsByExam,
  getExamResultsByStudent,
  updateExamResult,
  deleteExamResult,
};

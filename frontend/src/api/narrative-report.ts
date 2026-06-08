/**
 * Narrative Report API Client
 * Handles all narrative report-related API calls
 */

import apiClient, { handleApiError } from './client';

// Types
export interface NarrativeReport {
  id: string;
  student_id: string;
  period_id: string;
  status: string;
  narrative_content: string;
  achievement_data: any;
  created_by: string;
  published_by?: string;
  published_at?: string;
  created_at: string;
  updated_at: string;
}

export interface CreateNarrativeReportRequest {
  student_id: string;
  period_id: string;
  narrative_content: string;
  achievement_data?: any;
}

export interface UpdateNarrativeReportRequest {
  narrative_content?: string;
  achievement_data?: any;
  status?: string;
}

export interface NarrativeReportResponse {
  id: string;
  student_id: string;
  period_id: string;
  status: string;
  narrative_content: string;
  achievement_data: any;
  created_by: string;
  published_by?: string;
  published_at?: string;
  created_at: string;
  updated_at: string;
}

/**
 * Get all narrative reports with optional filters
 */
export const getNarrativeReports = async (params?: {
  student_id?: string;
  period_id?: string;
  status?: string;
  limit?: number;
  offset?: number;
}): Promise<NarrativeReport[]> => {
  try {
    const response = await apiClient.get('/reports', { params });
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get narrative report by ID
 */
export const getNarrativeReportById = async (id: string): Promise<NarrativeReport> => {
  try {
    const response = await apiClient.get(`/reports/${id}`);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Create new narrative report
 */
export const createNarrativeReport = async (
  data: CreateNarrativeReportRequest
): Promise<NarrativeReport> => {
  try {
    const response = await apiClient.post('/reports', data);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Update narrative report
 */
export const updateNarrativeReport = async (
  id: string,
  data: UpdateNarrativeReportRequest
): Promise<NarrativeReport> => {
  try {
    const response = await apiClient.put(`/reports/${id}`, data);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Delete narrative report
 */
export const deleteNarrativeReport = async (id: string): Promise<void> => {
  try {
    await apiClient.delete(`/reports/${id}`);
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Publish narrative report
 */
export const publishNarrativeReport = async (id: string): Promise<NarrativeReport> => {
  try {
    const response = await apiClient.post(`/reports/${id}/publish`);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get achievement summary for report
 */
export const getAchievementSummary = async (id: string): Promise<any> => {
  try {
    const response = await apiClient.get(`/reports/${id}/achievement-summary`);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

export default {
  getNarrativeReports,
  getNarrativeReportById,
  createNarrativeReport,
  updateNarrativeReport,
  deleteNarrativeReport,
  publishNarrativeReport,
  getAchievementSummary,
};

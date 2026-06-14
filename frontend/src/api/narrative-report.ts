/**
 * Narrative Report API Client
 * Handles all narrative report-related API calls with proper types
 */

import apiClient, { handleApiError } from './client';
import {
  NarrativeReport,
  NarrativeContent,
  StudentAchievement,
  AssessmentStatus,
  PaginationParams,
  FilterParams
} from '@/shared/types/domain';

// API-specific request types
export interface NarrativeReportCreateRequest {
  student_id: string;
  class_id: string;
  subject_id: string;
  reporting_period: {
    semester: string;
    academic_year: string;
    start_date: string;
    end_date: string;
  };
  narrative_content: NarrativeContent;
}

export interface NarrativeReportUpdateRequest {
  narrative_content?: NarrativeContent;
  teacher_recommendations?: string[];
  status?: AssessmentStatus;
}

/**
 * Get all narrative reports with optional filters
 */
export const getNarrativeReports = async (params?: PaginationParams & FilterParams & {
  student_id?: string;
  class_id?: string;
  subject_id?: string;
}): Promise<NarrativeReport[]> => {
  try {
    const response = await apiClient.get('reporting/narrative-reports', { params });
    return response.data.narrative_reports || response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get narrative report by ID
 */
export const getNarrativeReportById = async (id: string): Promise<NarrativeReport> => {
  try {
    const response = await apiClient.get(`reporting/narrative-reports/${id}`);
    return response.data.narrative_report || response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Create new narrative report
 */
export const createNarrativeReport = async (
  data: NarrativeReportCreateRequest
): Promise<NarrativeReport> => {
  try {
    const response = await apiClient.post('reporting/narrative-reports', data);
    return response.data.narrative_report || response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Update narrative report
 */
export const updateNarrativeReport = async (
  id: string,
  data: NarrativeReportUpdateRequest
): Promise<NarrativeReport> => {
  try {
    const response = await apiClient.put(`reporting/narrative-reports/${id}`, data);
    return response.data.narrative_report || response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Delete narrative report
 */
export const deleteNarrativeReport = async (id: string): Promise<void> => {
  try {
    await apiClient.delete(`reporting/narrative-reports/${id}`);
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Publish narrative report
 */
export const publishNarrativeReport = async (id: string): Promise<NarrativeReport> => {
  try {
    const response = await apiClient.post(`reporting/narrative-reports/${id}/publish`);
    return response.data.narrative_report || response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get achievement summary for report
 */
export const getAchievementSummary = async (id: string, params?: PaginationParams & {
  student_id?: string;
  class_id?: string;
}): Promise<StudentAchievement> => {
  try {
    const response = await apiClient.get(`reporting/narrative-reports/${id}/achievement-summary`, { params });
    return response.data.achievement_summary || response.data.data || response.data;
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

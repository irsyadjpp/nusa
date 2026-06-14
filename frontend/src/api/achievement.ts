/**
 * Achievement API Client
 * Handles all Achievement-related API calls with proper types
 */

import apiClient, { handleApiError } from './client';
import {
  StudentAchievement,
  ClassAchievement,
  CompetencyProgress,
  MasteryLevel,
  PaginationParams
} from '@/shared/types/domain';

// API-specific types for achievement data
export interface StudentTrajectory {
  student_id: string;
  student_name: string;
  competency_id: string;
  trajectory_points: {
    date: string;
    score: number;
    mastery_level: MasteryLevel;
  }[];
}

export interface AchievementSummary {
  report_id: string;
  student_id: string;
  student_name: string;
  period_start: string;
  period_end: string;
  overall_mastery: number;
  competency_summary: {
    competency_id: string;
    competency_name: string;
    mastery_level: MasteryLevel;
    score: number;
    max_score: number;
  }[];
  achievements: string[];
  areas_for_improvement: string[];
  recommendations: string[];
}

/**
 * Get student achievement
 */
export const getStudentAchievement = async (studentId: string, params?: PaginationParams & {
  tp_id?: string;
}): Promise<StudentAchievement[]> => {
  try {
    const response = await apiClient.get(`students/${studentId}/achievement`, { params });
    return response.data.student_achievement || response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get student progress
 */
export const getStudentProgress = async (studentId: string, params?: PaginationParams & {
  subject_id?: string;
  phase_id?: string;
}): Promise<CompetencyProgress[]> => {
  try {
    const response = await apiClient.get(`students/${studentId}/progress`, { params });
    return response.data.competency_progress || response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get class achievement
 */
export const getClassAchievement = async (classId: string, params?: PaginationParams & {
  subject_id?: string;
}): Promise<ClassAchievement> => {
  try {
    const response = await apiClient.get(`classes/${classId}/achievement`, { params });
    return response.data.class_achievement || response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get report achievement summary
 */
export const getReportAchievementSummary = async (reportId: string, params?: PaginationParams & {
  student_id?: string;
  class_id?: string;
}): Promise<AchievementSummary> => {
  try {
    const response = await apiClient.get(`reports/${reportId}/achievement-summary`, { params });
    return response.data.achievement_summary || response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get student trajectory (progress over time)
 */
export const getStudentTrajectory = async (studentId: string, params?: PaginationParams & {
  competency_id?: string;
  start_date?: string;
  end_date?: string;
}): Promise<StudentTrajectory> => {
  try {
    const response = await apiClient.get(`/students/${studentId}/trajectory`, { params });
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

export default {
  getStudentAchievement,
  getStudentProgress,
  getClassAchievement,
  getReportAchievementSummary,
  getStudentTrajectory,
};

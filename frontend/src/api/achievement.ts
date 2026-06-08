/**
 * Achievement API Client
 * Handles all Achievement-related API calls
 */

import apiClient, { handleApiError } from './client';

// Types
export interface StudentAchievement {
  student_id: string;
  student_name: string;
  tp_id: string;
  tp_title: string;
  competency_id: string;
  competency_name: string;
  mastery_level: string;
  score: number;
  max_score: number;
  percentage: number;
  achieved_criteria: string[];
  pending_criteria: string[];
  last_updated: string;
}

export interface CompetencyProgress {
  competency_id: string;
  competency_name: string;
  total_assessments: number;
  completed_assessments: number;
  average_score: number;
  mastery_level: string;
  progress_percentage: number;
  criteria_progress: {
    criteria_id: string;
    criteria_name: string;
    achieved: boolean;
    evidence_count: number;
  }[];
}

export interface ClassAchievement {
  class_id: string;
  class_name: string;
  subject_id: string;
  subject_name: string;
  total_students: number;
  average_mastery: number;
  competency_achievements: {
    competency_id: string;
    competency_name: string;
    average_score: number;
    mastery_distribution: {
      excellent: number;
      proficient: number;
      developing: number;
      beginning: number;
    };
  }[];
  top_performers: {
    student_id: string;
    student_name: string;
    average_score: number;
  }[];
  areas_for_improvement: {
    competency_id: string;
    competency_name: string;
    average_score: number;
    struggling_students: number;
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
    mastery_level: string;
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
export const getStudentAchievement = async (studentId: string, params?: {
  tp_id?: string;
  competency_id?: string;
}): Promise<StudentAchievement[]> => {
  try {
    const response = await apiClient.get(`/students/${studentId}/achievement`, { params });
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get student progress
 */
export const getStudentProgress = async (studentId: string, params?: {
  tp_id?: string;
  competency_id?: string;
}): Promise<CompetencyProgress[]> => {
  try {
    const response = await apiClient.get(`/students/${studentId}/progress`, { params });
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get class achievement
 */
export const getClassAchievement = async (classId: string, params?: {
  subject_id?: string;
  tp_id?: string;
}): Promise<ClassAchievement> => {
  try {
    const response = await apiClient.get(`/classes/${classId}/achievement`, { params });
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get report achievement summary
 */
export const getReportAchievementSummary = async (reportId: string): Promise<AchievementSummary> => {
  try {
    const response = await apiClient.get(`/reports/${reportId}/achievement-summary`);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get student trajectory (progress over time)
 */
export const getStudentTrajectory = async (studentId: string, params?: {
  competency_id?: string;
  start_date?: string;
  end_date?: string;
}): Promise<{
  student_id: string;
  student_name: string;
  trajectory_points: {
    date: string;
    score: number;
    mastery_level: string;
  }[];
}> => {
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

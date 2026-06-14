/**
 * Schedule API Service
 * API calls for schedule management operations
 */

import apiClient, { handleApiError } from './client';

// TypeScript interfaces
export interface Schedule {
  id: string;
  class_id: string;
  subject_id?: string;
  teacher_id: string;
  day_of_week: string;
  start_time: string;
  end_time: string;
  room?: string;
  academic_year_id: string;
  semester_id: string;
  is_active: boolean;
  created_at: string;
  updated_at: string;
}

export interface CreateScheduleRequest {
  class_id: string;
  subject_id?: string;
  teacher_id: string;
  day_of_week: string;
  start_time: string;
  end_time: string;
  room?: string;
  academic_year_id: string;
  semester_id: string;
}

export interface UpdateScheduleRequest {
  class_id?: string;
  subject_id?: string;
  teacher_id?: string;
  day_of_week?: string;
  start_time?: string;
  end_time?: string;
  room?: string;
  is_active?: boolean;
}

export interface ListSchedulesResponse {
  schedules: Schedule[];
  total: number;
  page: number;
  page_size: number;
}

export interface ListSchedulesParams {
  page?: number;
  page_size?: number;
  class_id?: string;
  teacher_id?: string;
  subject_id?: string;
  academic_year_id?: string;
  semester_id?: string;
  day_of_week?: string;
  is_active?: boolean;
}

export interface ScheduleConflict {
  has_conflict: boolean;
  conflicts: Array<{
    schedule_id: string;
    reason: string;
  }>;
}

/**
 * Create new schedule
 */
export const createSchedule = async (data: CreateScheduleRequest): Promise<Schedule> => {
  try {
    const response = await apiClient.post('/schedules', data);
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * List schedules with filters
 */
export const listSchedules = async (params?: ListSchedulesParams): Promise<ListSchedulesResponse> => {
  try {
    const response = await apiClient.get('/schedules', { params });
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Check for schedule conflicts
 */
export const checkScheduleConflicts = async (data: CreateScheduleRequest): Promise<ScheduleConflict> => {
  try {
    const response = await apiClient.get('/schedules/conflicts', { params: data });
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Update schedule
 */
export const updateSchedule = async (id: string, data: UpdateScheduleRequest): Promise<Schedule> => {
  try {
    const response = await apiClient.put(`/schedules/${id}`, data);
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Delete schedule
 */
export const deleteSchedule = async (id: string): Promise<void> => {
  try {
    await apiClient.delete(`/schedules/${id}`);
  } catch (error) {
    throw handleApiError(error);
  }
};

export default {
  createSchedule,
  listSchedules,
  checkScheduleConflicts,
  updateSchedule,
  deleteSchedule,
};

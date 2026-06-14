/**
 * Attendance API Service
 * API calls for attendance management operations
 */

import apiClient, { handleApiError } from './client';

// TypeScript interfaces
export interface AttendanceRecord {
  id: string;
  class_id: string;
  student_id: string;
  date: string;
  status: 'present' | 'absent' | 'late' | 'excused';
  notes?: string;
  recorded_by: string;
  created_at: string;
  updated_at: string;
}

export interface CreateAttendanceRequest {
  class_id: string;
  student_id: string;
  date: string;
  status: 'present' | 'absent' | 'late' | 'excused';
  notes?: string;
}

export interface BulkAttendanceRequest {
  class_id: string;
  date: string;
  records: Array<{
    student_id: string;
    status: 'present' | 'absent' | 'late' | 'excused';
    notes?: string;
  }>;
}

export interface UpdateAttendanceRequest {
  status?: 'present' | 'absent' | 'late' | 'excused';
  notes?: string;
}

export interface ListAttendanceResponse {
  records: AttendanceRecord[];
  total: number;
  page: number;
  page_size: number;
}

export interface ListAttendanceParams {
  page?: number;
  page_size?: number;
  class_id?: string;
  student_id?: string;
  date_from?: string;
  date_to?: string;
  status?: string;
}

export interface AttendanceReport {
  class_id: string;
  date: string;
  total_students: number;
  present_count: number;
  absent_count: number;
  late_count: number;
  excused_count: number;
  attendance_rate: number;
}

/**
 * Record attendance for a single student
 */
export const createAttendance = async (data: CreateAttendanceRequest): Promise<AttendanceRecord> => {
  try {
    const response = await apiClient.post('/attendance', data);
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Bulk record attendance for multiple students
 */
export const bulkCreateAttendance = async (data: BulkAttendanceRequest): Promise<AttendanceRecord[]> => {
  try {
    const response = await apiClient.post('/attendance/bulk', data);
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * List attendance records with filters
 */
export const listAttendance = async (params?: ListAttendanceParams): Promise<ListAttendanceResponse> => {
  try {
    const response = await apiClient.get('/attendance', { params });
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Generate attendance report
 */
export const getAttendanceReport = async (class_id: string, date: string): Promise<AttendanceReport> => {
  try {
    const response = await apiClient.get('/attendance/report', { params: { class_id, date } });
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Update attendance record
 */
export const updateAttendance = async (id: string, data: UpdateAttendanceRequest): Promise<AttendanceRecord> => {
  try {
    const response = await apiClient.put(`/attendance/${id}`, data);
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

export default {
  createAttendance,
  bulkCreateAttendance,
  listAttendance,
  getAttendanceReport,
  updateAttendance,
};

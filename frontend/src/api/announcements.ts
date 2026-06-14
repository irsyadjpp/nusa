/**
 * Announcement API Service
 * API calls for announcement management operations
 */

import apiClient, { handleApiError } from './client';

// TypeScript interfaces
export interface Announcement {
  id: string;
  school_id?: string;
  title: string;
  content: string;
  type: 'general' | 'urgent' | 'event';
  target_audience: 'all' | 'teachers' | 'students' | 'admins';
  publish_date: string;
  expiry_date?: string;
  created_by: string;
  is_active: boolean;
  created_at: string;
  updated_at: string;
}

export interface CreateAnnouncementRequest {
  school_id?: string;
  title: string;
  content: string;
  type: 'general' | 'urgent' | 'event';
  target_audience: 'all' | 'teachers' | 'students' | 'admins';
  publish_date: string;
  expiry_date?: string;
}

export interface UpdateAnnouncementRequest {
  title?: string;
  content?: string;
  type?: 'general' | 'urgent' | 'event';
  target_audience?: 'all' | 'teachers' | 'students' | 'admins';
  publish_date?: string;
  expiry_date?: string;
  is_active?: boolean;
}

export interface ListAnnouncementsResponse {
  announcements: Announcement[];
  total: number;
  page: number;
  page_size: number;
}

export interface ListAnnouncementsParams {
  page?: number;
  page_size?: number;
  school_id?: string;
  type?: string;
  target_audience?: string;
  is_active?: boolean;
  search?: string;
}

/**
 * Create announcement
 */
export const createAnnouncement = async (data: CreateAnnouncementRequest): Promise<Announcement> => {
  try {
    const response = await apiClient.post('/announcements', data);
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * List announcements with filters
 */
export const listAnnouncements = async (params?: ListAnnouncementsParams): Promise<ListAnnouncementsResponse> => {
  try {
    const response = await apiClient.get('/announcements', { params });
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get school announcements
 */
export const getSchoolAnnouncements = async (school_id: string): Promise<Announcement[]> => {
  try {
    const response = await apiClient.get(`/announcements/school/${school_id}`);
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Update announcement
 */
export const updateAnnouncement = async (id: string, data: UpdateAnnouncementRequest): Promise<Announcement> => {
  try {
    const response = await apiClient.put(`/announcements/${id}`, data);
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Delete announcement (soft delete)
 */
export const deleteAnnouncement = async (id: string): Promise<void> => {
  try {
    await apiClient.delete(`/announcements/${id}`);
  } catch (error) {
    throw handleApiError(error);
  }
};

export default {
  createAnnouncement,
  listAnnouncements,
  getSchoolAnnouncements,
  updateAnnouncement,
  deleteAnnouncement,
};

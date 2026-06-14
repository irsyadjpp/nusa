/**
 * Notification API Service
 * API calls for notification management operations
 */

import apiClient, { handleApiError } from './client';

// TypeScript interfaces
export interface Notification {
  id: string;
  recipient_id: string;
  title: string;
  message: string;
  type: 'info' | 'warning' | 'error' | 'success';
  is_read: boolean;
  created_at: string;
  updated_at: string;
}

export interface CreateNotificationRequest {
  recipient_id: string;
  title: string;
  message: string;
  type: 'info' | 'warning' | 'error' | 'success';
}

export interface ListNotificationsResponse {
  notifications: Notification[];
  total: number;
  unread_count: number;
  page: number;
  page_size: number;
}

export interface ListNotificationsParams {
  page?: number;
  page_size?: number;
  recipient_id?: string;
  is_read?: boolean;
  type?: string;
}

/**
 * Create notification
 */
export const createNotification = async (data: CreateNotificationRequest): Promise<Notification> => {
  try {
    const response = await apiClient.post('/notifications', data);
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * List notifications with filters
 */
export const listNotifications = async (params?: ListNotificationsParams): Promise<ListNotificationsResponse> => {
  try {
    const response = await apiClient.get('/notifications', { params });
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Mark notification as read
 */
export const markNotificationAsRead = async (id: string): Promise<Notification> => {
  try {
    const response = await apiClient.put(`/notifications/${id}/read`);
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

export default {
  createNotification,
  listNotifications,
  markNotificationAsRead,
};

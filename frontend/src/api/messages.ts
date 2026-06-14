/**
 * Message API Service
 * API calls for message management operations
 */

import apiClient, { handleApiError } from './client';

// TypeScript interfaces
export interface Message {
  id: string;
  sender_id: string;
  recipient_id: string;
  subject: string;
  content: string;
  is_read: boolean;
  created_at: string;
  updated_at: string;
}

export interface CreateMessageRequest {
  recipient_id: string;
  subject: string;
  content: string;
}

export interface UpdateMessageRequest {
  subject?: string;
  content?: string;
  is_read?: boolean;
}

export interface ListMessagesResponse {
  messages: Message[];
  total: number;
  unread_count: number;
  page: number;
  page_size: number;
}

export interface ListMessagesParams {
  page?: number;
  page_size?: number;
  sender_id?: string;
  recipient_id?: string;
  is_read?: boolean;
  search?: string;
}

export interface Conversation {
  user_id: string;
  user_name: string;
  last_message: string;
  last_message_date: string;
  unread_count: number;
}

/**
 * Create message
 */
export const createMessage = async (data: CreateMessageRequest): Promise<Message> => {
  try {
    const response = await apiClient.post('/messages', data);
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * List messages with filters
 */
export const listMessages = async (params?: ListMessagesParams): Promise<ListMessagesResponse> => {
  try {
    const response = await apiClient.get('/messages', { params });
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get conversation with a user
 */
export const getConversation = async (user_id: string): Promise<Message[]> => {
  try {
    const response = await apiClient.get(`/messages/conversation/${user_id}`);
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get all conversations
 */
export const getConversations = async (): Promise<Conversation[]> => {
  try {
    const response = await apiClient.get('/messages/conversations');
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Update message
 */
export const updateMessage = async (id: string, data: UpdateMessageRequest): Promise<Message> => {
  try {
    const response = await apiClient.put(`/messages/${id}`, data);
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Mark message as read
 */
export const markMessageAsRead = async (id: string): Promise<Message> => {
  try {
    const response = await apiClient.put(`/messages/${id}/read`);
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Delete message (soft delete)
 */
export const deleteMessage = async (id: string): Promise<void> => {
  try {
    await apiClient.delete(`/messages/${id}`);
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Get unread count
 */
export const getUnreadCount = async (): Promise<number> => {
  try {
    const response = await apiClient.get('/messages/unread-count');
    return response.data.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

export default {
  createMessage,
  listMessages,
  getConversation,
  getConversations,
  updateMessage,
  markMessageAsRead,
  deleteMessage,
  getUnreadCount,
};

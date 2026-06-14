/**
 * Messages Query Service
 * Provides query operations for message data using TanStack Query
 */

import { useQuery, UseQueryOptions } from '@tanstack/react-query';
import * as messagesApi from '@/api/messages';

// Query Keys
export const messagesKeys = {
  all: ['messages'] as const,
  list: (params?: any) => ['messages', 'list', params] as const,
  conversation: (user_id: string) => ['messages', 'conversation', user_id] as const,
  conversations: () => ['messages', 'conversations'] as const,
  unreadCount: () => ['messages', 'unread-count'] as const,
};

/**
 * Message Queries
 */
export const useMessages = (
  params?: any,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: messagesKeys.list(params),
    queryFn: () => messagesApi.listMessages(params),
    staleTime: 60000, // 1 minute - messages change frequently
    ...options,
  });
};

export const useConversation = (
  user_id: string,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: messagesKeys.conversation(user_id),
    queryFn: () => messagesApi.getConversation(user_id),
    staleTime: 60000, // 1 minute
    ...options,
  });
};

export const useConversations = (
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: messagesKeys.conversations(),
    queryFn: () => messagesApi.getConversations(),
    staleTime: 60000, // 1 minute
    ...options,
  });
};

export const useUnreadCount = (
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: messagesKeys.unreadCount(),
    queryFn: () => messagesApi.getUnreadCount(),
    staleTime: 30000, // 30 seconds - unread count changes frequently
    ...options,
  });
};

/**
 * Query invalidation functions
 */
export const invalidateMessageQueries = (queryClient: any) => {
  queryClient.invalidateQueries({ queryKey: messagesKeys.all });
};

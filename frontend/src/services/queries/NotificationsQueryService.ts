/**
 * Notifications Query Service
 * Provides query operations for notification data using TanStack Query
 */

import { useQuery, UseQueryOptions } from '@tanstack/react-query';
import * as notificationsApi from '@/api/notifications';

// Query Keys
export const notificationsKeys = {
  all: ['notifications'] as const,
  list: (params?: any) => ['notifications', 'list', params] as const,
};

/**
 * Notification Queries
 */
export const useNotifications = (
  params?: any,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: notificationsKeys.list(params),
    queryFn: () => notificationsApi.listNotifications(params),
    staleTime: 60000, // 1 minute - notifications change frequently
    ...options,
  });
};

/**
 * Query invalidation functions
 */
export const invalidateNotificationQueries = (queryClient: any) => {
  queryClient.invalidateQueries({ queryKey: notificationsKeys.all });
};

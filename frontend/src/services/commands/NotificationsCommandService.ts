/**
 * Notifications Command Service
 * Provides command operations for notification data using TanStack Query mutations
 */

import { useMutation, useQueryClient, UseMutationOptions } from '@tanstack/react-query';
import * as notificationsApi from '@/api/notifications';
import { notificationsKeys } from '@/services/queries/NotificationsQueryService';

/**
 * Create Notification Mutation
 */
export const useCreateNotification = (
  options?: Omit<UseMutationOptions<any, Error, any>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: any) => notificationsApi.createNotification(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: notificationsKeys.all });
    },
    ...options,
  });
};

/**
 * Mark Notification as Read Mutation
 */
export const useMarkNotificationAsRead = (
  options?: Omit<UseMutationOptions<any, Error, string>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => notificationsApi.markNotificationAsRead(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: notificationsKeys.all });
    },
    ...options,
  });
};

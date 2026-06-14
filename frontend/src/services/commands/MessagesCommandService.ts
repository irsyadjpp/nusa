/**
 * Messages Command Service
 * Provides command operations for message data using TanStack Query mutations
 */

import { useMutation, useQueryClient, UseMutationOptions } from '@tanstack/react-query';
import * as messagesApi from '@/api/messages';
import { messagesKeys } from '@/services/queries/MessagesQueryService';

/**
 * Create Message Mutation
 */
export const useCreateMessage = (
  options?: Omit<UseMutationOptions<any, Error, any>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: any) => messagesApi.createMessage(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: messagesKeys.all });
    },
    ...options,
  });
};

/**
 * Update Message Mutation
 */
export const useUpdateMessage = (
  options?: Omit<UseMutationOptions<any, Error, any>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, data }: { id: string; data: any }) =>
      messagesApi.updateMessage(id, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: messagesKeys.all });
    },
    ...options,
  });
};

/**
 * Mark Message as Read Mutation
 */
export const useMarkMessageAsRead = (
  options?: Omit<UseMutationOptions<any, Error, string>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => messagesApi.markMessageAsRead(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: messagesKeys.all });
    },
    ...options,
  });
};

/**
 * Delete Message Mutation
 */
export const useDeleteMessage = (
  options?: Omit<UseMutationOptions<void, Error, string>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => messagesApi.deleteMessage(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: messagesKeys.all });
    },
    ...options,
  });
};

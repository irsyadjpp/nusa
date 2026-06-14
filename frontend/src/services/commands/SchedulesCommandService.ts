/**
 * Schedules Command Service
 * Provides command operations for schedule data using TanStack Query mutations
 */

import { useMutation, useQueryClient, UseMutationOptions } from '@tanstack/react-query';
import * as schedulesApi from '@/api/schedules';
import { schedulesKeys } from '@/services/queries/SchedulesQueryService';

/**
 * Create Schedule Mutation
 */
export const useCreateSchedule = (
  options?: Omit<UseMutationOptions<any, Error, any>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: any) => schedulesApi.createSchedule(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: schedulesKeys.all });
    },
    ...options,
  });
};

/**
 * Update Schedule Mutation
 */
export const useUpdateSchedule = (
  options?: Omit<UseMutationOptions<any, Error, any>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, data }: { id: string; data: any }) =>
      schedulesApi.updateSchedule(id, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: schedulesKeys.all });
    },
    ...options,
  });
};

/**
 * Delete Schedule Mutation
 */
export const useDeleteSchedule = (
  options?: Omit<UseMutationOptions<void, Error, string>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => schedulesApi.deleteSchedule(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: schedulesKeys.all });
    },
    ...options,
  });
};

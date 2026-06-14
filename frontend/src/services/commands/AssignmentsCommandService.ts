/**
 * Assignments Command Service
 * Provides command operations for assignment data using TanStack Query mutations
 */

import { useMutation, useQueryClient, UseMutationOptions } from '@tanstack/react-query';
import * as assignmentsApi from '@/api/assignments';
import { assignmentsKeys } from '@/services/queries/AssignmentsQueryService';

/**
 * Create Assignment Mutation
 */
export const useCreateAssignment = (
  options?: Omit<UseMutationOptions<any, Error, any>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: any) => assignmentsApi.createAssignment(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: assignmentsKeys.all });
    },
    ...options,
  });
};

/**
 * Update Assignment Mutation
 */
export const useUpdateAssignment = (
  options?: Omit<UseMutationOptions<any, Error, any>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, data }: { id: string; data: any }) =>
      assignmentsApi.updateAssignment(id, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: assignmentsKeys.all });
    },
    ...options,
  });
};

/**
 * Delete Assignment Mutation
 */
export const useDeleteAssignment = (
  options?: Omit<UseMutationOptions<void, Error, string>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => assignmentsApi.deleteAssignment(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: assignmentsKeys.all });
    },
    ...options,
  });
};

/**
 * Rubric Command Service
 * Provides command operations for Rubric data using TanStack Query mutations
 */

import { useMutation, useQueryClient, UseMutationOptions } from '@tanstack/react-query';
import * as rubricApi from '@/api/rubric';
import { rubricKeys } from '../queries/RubricQueryService';

/**
 * Create Rubric mutation
 */
export const useCreateRubric = (
  options?: Omit<UseMutationOptions<any, Error, { data: any; userId: string }>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ data, userId }) => rubricApi.createRubric(data, userId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: rubricKeys.all });
    },
    ...options,
  });
};

/**
 * Update Rubric mutation
 */
export const useUpdateRubric = (
  options?: Omit<UseMutationOptions<any, Error, { id: string; data: any }>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, data }) => rubricApi.updateRubric(id, data),
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({ queryKey: rubricKeys.detail(variables.id) });
      queryClient.invalidateQueries({ queryKey: rubricKeys.all });
    },
    ...options,
  });
};

/**
 * Delete Rubric mutation
 */
export const useDeleteRubric = (
  options?: Omit<UseMutationOptions<void, Error, string>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => rubricApi.deleteRubric(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: rubricKeys.all });
    },
    ...options,
  });
};

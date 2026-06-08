/**
 * Evaluation Command Service
 * Provides command operations for Evaluation data using TanStack Query mutations
 */

import { useMutation, useQueryClient, UseMutationOptions } from '@tanstack/react-query';
import * as evaluationApi from '@/api/evaluation';
import { evaluationKeys } from '../queries/EvaluationQueryService';

/**
 * Create Evaluation mutation
 */
export const useCreateEvaluation = (
  options?: Omit<UseMutationOptions<any, Error, any>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: any) => evaluationApi.createEvaluation(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: evaluationKeys.all });
    },
    ...options,
  });
};

/**
 * Update Evaluation mutation
 */
export const useUpdateEvaluation = (
  options?: Omit<UseMutationOptions<any, Error, { id: string; data: any }>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, data }) => evaluationApi.updateEvaluation(id, data),
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({ queryKey: evaluationKeys.detail(variables.id) });
      queryClient.invalidateQueries({ queryKey: evaluationKeys.all });
    },
    ...options,
  });
};

/**
 * Delete Evaluation mutation
 */
export const useDeleteEvaluation = (
  options?: Omit<UseMutationOptions<void, Error, string>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => evaluationApi.deleteEvaluation(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: evaluationKeys.all });
    },
    ...options,
  });
};

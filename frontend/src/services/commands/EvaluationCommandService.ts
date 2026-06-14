/**
 * Evaluation Command Service
 * Provides command operations for Evaluation data using TanStack Query mutations with proper types
 */

import { useMutation, useQueryClient, UseMutationOptions } from '@tanstack/react-query';
import * as evaluationApi from '@/api/evaluation';
import { evaluationKeys } from '../queries/EvaluationQueryService';
import { Evaluation, CreateEvaluationRequest, EvaluationUpdateRequest } from '@/shared/types/domain';

/**
 * Create Evaluation mutation
 */
export const useCreateEvaluation = (
  options?: Omit<UseMutationOptions<Evaluation, Error, { data: CreateEvaluationRequest; userId: string }>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation<Evaluation, Error, { data: CreateEvaluationRequest; userId: string }>({
    mutationFn: ({ data, userId }) => evaluationApi.createEvaluation(data, userId),
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
  options?: Omit<UseMutationOptions<Evaluation, Error, { id: string; data: EvaluationUpdateRequest }>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation<Evaluation, Error, { id: string; data: EvaluationUpdateRequest }>({
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

  return useMutation<void, Error, string>({
    mutationFn: (id: string) => evaluationApi.deleteEvaluation(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: evaluationKeys.all });
    },
    ...options,
  });
};

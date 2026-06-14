/**
 * Exam Results Command Service
 * Provides command operations for exam result data using TanStack Query mutations
 */

import { useMutation, useQueryClient, UseMutationOptions } from '@tanstack/react-query';
import * as examResultsApi from '@/api/exam-results';
import { examResultsKeys } from '@/services/queries/ExamResultsQueryService';

/**
 * Create Exam Result Mutation
 */
export const useCreateExamResult = (
  options?: Omit<UseMutationOptions<any, Error, any>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: any) => examResultsApi.createExamResult(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: examResultsKeys.all });
    },
    ...options,
  });
};

/**
 * Update Exam Result Mutation
 */
export const useUpdateExamResult = (
  options?: Omit<UseMutationOptions<any, Error, any>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, data }: { id: string; data: any }) =>
      examResultsApi.updateExamResult(id, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: examResultsKeys.all });
    },
    ...options,
  });
};

/**
 * Delete Exam Result Mutation
 */
export const useDeleteExamResult = (
  options?: Omit<UseMutationOptions<void, Error, string>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => examResultsApi.deleteExamResult(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: examResultsKeys.all });
    },
    ...options,
  });
};

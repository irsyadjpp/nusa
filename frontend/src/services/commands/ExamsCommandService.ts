/**
 * Exams Command Service
 * Provides command operations for exam data using TanStack Query mutations
 */

import { useMutation, useQueryClient, UseMutationOptions } from '@tanstack/react-query';
import * as examsApi from '@/api/exams';
import { examsKeys } from '@/services/queries/ExamsQueryService';

/**
 * Create Exam Mutation
 */
export const useCreateExam = (
  options?: Omit<UseMutationOptions<any, Error, any>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: any) => examsApi.createExam(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: examsKeys.all });
    },
    ...options,
  });
};

/**
 * Update Exam Mutation
 */
export const useUpdateExam = (
  options?: Omit<UseMutationOptions<any, Error, any>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, data }: { id: string; data: any }) =>
      examsApi.updateExam(id, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: examsKeys.all });
    },
    ...options,
  });
};

/**
 * Delete Exam Mutation
 */
export const useDeleteExam = (
  options?: Omit<UseMutationOptions<void, Error, string>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => examsApi.deleteExam(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: examsKeys.all });
    },
    ...options,
  });
};

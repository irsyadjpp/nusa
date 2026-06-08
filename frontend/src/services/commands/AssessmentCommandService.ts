/**
 * Assessment Command Service
 * Provides command operations for Assessment data using TanStack Query mutations
 */

import { useMutation, useQueryClient, UseMutationOptions } from '@tanstack/react-query';
import * as assessmentApi from '@/api/assessment';
import { assessmentKeys } from '../queries/AssessmentQueryService';

/**
 * Create Assessment mutation
 */
export const useCreateAssessment = (
  options?: Omit<UseMutationOptions<any, Error, any>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: any) => assessmentApi.createAssessment(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: assessmentKeys.all });
    },
    ...options,
  });
};

/**
 * Update Assessment mutation
 */
export const useUpdateAssessment = (
  options?: Omit<UseMutationOptions<any, Error, { id: string; data: any }>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, data }) => assessmentApi.updateAssessment(id, data),
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({ queryKey: assessmentKeys.detail(variables.id) });
      queryClient.invalidateQueries({ queryKey: assessmentKeys.all });
    },
    ...options,
  });
};

/**
 * Delete Assessment mutation
 */
export const useDeleteAssessment = (
  options?: Omit<UseMutationOptions<void, Error, string>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => assessmentApi.deleteAssessment(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: assessmentKeys.all });
    },
    ...options,
  });
};

/**
 * Approve Assessment mutation
 */
export const useApproveAssessment = (
  options?: Omit<UseMutationOptions<any, Error, string>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => assessmentApi.approveAssessment(id),
    onSuccess: (_, id) => {
      queryClient.invalidateQueries({ queryKey: assessmentKeys.detail(id) });
      queryClient.invalidateQueries({ queryKey: assessmentKeys.all });
    },
    ...options,
  });
};

/**
 * Reject Assessment mutation
 */
export const useRejectAssessment = (
  options?: Omit<UseMutationOptions<any, Error, string>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => assessmentApi.rejectAssessment(id),
    onSuccess: (_, id) => {
      queryClient.invalidateQueries({ queryKey: assessmentKeys.detail(id) });
      queryClient.invalidateQueries({ queryKey: assessmentKeys.all });
    },
    ...options,
  });
};

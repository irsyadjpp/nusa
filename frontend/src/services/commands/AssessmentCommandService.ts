/**
 * Assessment Command Service
 * Provides command operations for Assessment data using TanStack Query mutations with proper types
 */

import { useMutation, useQueryClient, UseMutationOptions } from '@tanstack/react-query';
import * as assessmentApi from '@/api/assessment';
import { assessmentKeys } from '../queries/AssessmentQueryService';
import { Assessment, CreateAssessmentRequest, UpdateAssessmentRequest } from '@/shared/types/domain';

/**
 * Create Assessment mutation
 */
export const useCreateAssessment = (
  options?: Omit<UseMutationOptions<Assessment, Error, { data: CreateAssessmentRequest; userId: string }>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation<Assessment, Error, { data: CreateAssessmentRequest; userId: string }>({
    mutationFn: ({ data, userId }) => assessmentApi.createAssessment(data, userId),
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
  options?: Omit<UseMutationOptions<Assessment, Error, { id: string; data: UpdateAssessmentRequest }>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation<Assessment, Error, { id: string; data: UpdateAssessmentRequest }>({
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

  return useMutation<void, Error, string>({
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
  options?: Omit<UseMutationOptions<Assessment, Error, { id: string; userId: string }>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation<Assessment, Error, { id: string; userId: string }>({
    mutationFn: ({ id, userId }) => assessmentApi.approveAssessment(id, userId),
    onSuccess: (_, { id }) => {
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
  options?: Omit<UseMutationOptions<Assessment, Error, { id: string; userId: string }>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation<Assessment, Error, { id: string; userId: string }>({
    mutationFn: ({ id, userId }) => assessmentApi.rejectAssessment(id, userId),
    onSuccess: (_, { id }) => {
      queryClient.invalidateQueries({ queryKey: assessmentKeys.detail(id) });
      queryClient.invalidateQueries({ queryKey: assessmentKeys.all });
    },
    ...options,
  });
};

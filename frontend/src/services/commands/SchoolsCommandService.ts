/**
 * Schools Command Service
 * Provides command operations for school data using TanStack Query mutations
 */

import { useMutation, useQueryClient, UseMutationOptions } from '@tanstack/react-query';
import * as schoolsApi from '@/api/schools';
import { schoolsKeys } from '@/services/queries/SchoolsQueryService';

/**
 * Create School Mutation
 */
export const useCreateSchool = (
  options?: Omit<UseMutationOptions<any, Error, any>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: any) => schoolsApi.createSchool(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: schoolsKeys.all });
    },
    ...options,
  });
};

/**
 * Update School Mutation
 */
export const useUpdateSchool = (
  options?: Omit<UseMutationOptions<any, Error, any>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, data }: { id: string; data: any }) =>
      schoolsApi.updateSchool(id, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: schoolsKeys.all });
    },
    ...options,
  });
};

/**
 * Update School Status Mutation
 */
export const useUpdateSchoolStatus = (
  options?: Omit<UseMutationOptions<any, Error, any>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, is_active }: { id: string; is_active: boolean }) =>
      schoolsApi.updateSchoolStatus(id, is_active),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: schoolsKeys.all });
    },
    ...options,
  });
};

/**
 * Delete School Mutation
 */
export const useDeleteSchool = (
  options?: Omit<UseMutationOptions<void, Error, string>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => schoolsApi.deleteSchool(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: schoolsKeys.all });
    },
    ...options,
  });
};

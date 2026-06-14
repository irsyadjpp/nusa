/**
 * Classes Command Service
 * Provides command operations for class data using TanStack Query mutations
 */

import { useMutation, useQueryClient, UseMutationOptions } from '@tanstack/react-query';
import * as classesApi from '@/api/classes';
import { classesKeys } from '@/services/queries/ClassesQueryService';

/**
 * Create Class Mutation
 */
export const useCreateClass = (
  options?: Omit<UseMutationOptions<any, Error, any>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: any) => classesApi.createClass(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: classesKeys.all });
    },
    ...options,
  });
};

/**
 * Update Class Mutation
 */
export const useUpdateClass = (
  options?: Omit<UseMutationOptions<any, Error, any>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, data }: { id: string; data: any }) =>
      classesApi.updateClass(id, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: classesKeys.all });
    },
    ...options,
  });
};

/**
 * Delete Class Mutation
 */
export const useDeleteClass = (
  options?: Omit<UseMutationOptions<void, Error, string>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => classesApi.deleteClass(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: classesKeys.all });
    },
    ...options,
  });
};

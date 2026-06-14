/**
 * ATP Command Service
 * Provides command operations for ATP data using TanStack Query mutations
 */

import { useMutation, useQueryClient, UseMutationOptions } from '@tanstack/react-query';
import * as atpApi from '@/api/atp';
import { atpKeys } from '../queries/ATPQueryService';

/**
 * Create ATP mutation
 */
export const useCreateATP = (
  options?: Omit<UseMutationOptions<any, Error, any>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: any) => atpApi.createATP(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: atpKeys.all });
    },
    ...options,
  });
};

/**
 * Update ATP mutation
 */
export const useUpdateATP = (
  options?: Omit<UseMutationOptions<any, Error, { id: string; data: any }>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, data }) => atpApi.updateATP(id, data),
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({ queryKey: atpKeys.detail(variables.id) });
      queryClient.invalidateQueries({ queryKey: atpKeys.all });
    },
    ...options,
  });
};

/**
 * Delete ATP mutation
 */
export const useDeleteATP = (
  options?: Omit<UseMutationOptions<void, Error, string>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => atpApi.deleteATP(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: atpKeys.all });
    },
    ...options,
  });
};

/**
 * Create ATP Set mutation
 */
export const useCreateATPSet = (
  options?: Omit<UseMutationOptions<any, Error, any>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: any) => atpApi.createATPSet(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: atpKeys.sets() });
    },
    ...options,
  });
};

/**
 * Approve ATP Set mutation
 */
export const useApproveATPSet = (
  options?: Omit<UseMutationOptions<any, Error, string>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => atpApi.approveATPSet(id),
    onSuccess: (_, id) => {
      queryClient.invalidateQueries({ queryKey: atpKeys.setDetail(id) });
      queryClient.invalidateQueries({ queryKey: atpKeys.sets() });
    },
    ...options,
  });
};

/**
 * Update ATP Set mutation
 */
export const useUpdateATPSet = (
  options?: Omit<UseMutationOptions<any, Error, { id: string; data: any }>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, data }) => atpApi.updateATPSet(id, data),
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({ queryKey: atpKeys.setDetail(variables.id) });
      queryClient.invalidateQueries({ queryKey: atpKeys.sets() });
    },
    ...options,
  });
};

/**
 * Delete ATP Set mutation
 */
export const useDeleteATPSet = (
  options?: Omit<UseMutationOptions<void, Error, string>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => atpApi.deleteATPSet(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: atpKeys.sets() });
    },
    ...options,
  });
};

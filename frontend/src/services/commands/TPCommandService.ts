/**
 * TP Command Service
 * Provides command operations for TP data using TanStack Query mutations
 */

import { useMutation, useQueryClient, UseMutationOptions } from '@tanstack/react-query';
import * as tpApi from '@/api/tp';
import { tpKeys } from '../queries/TPQueryService';

/**
 * Create TP mutation
 */
export const useCreateTP = (
  options?: Omit<UseMutationOptions<any, Error, any>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: any) => tpApi.createTP(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: tpKeys.all });
    },
    ...options,
  });
};

/**
 * Update TP mutation
 */
export const useUpdateTP = (
  options?: Omit<UseMutationOptions<any, Error, { id: string; data: any }>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, data }) => tpApi.updateTP(id, data),
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({ queryKey: tpKeys.detail(variables.id) });
      queryClient.invalidateQueries({ queryKey: tpKeys.all });
    },
    ...options,
  });
};

/**
 * Delete TP mutation
 */
export const useDeleteTP = (
  options?: Omit<UseMutationOptions<void, Error, string>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => tpApi.deleteTP(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: tpKeys.all });
    },
    ...options,
  });
};

/**
 * Create TP Set mutation
 */
export const useCreateTPSet = (
  options?: Omit<UseMutationOptions<any, Error, any>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: any) => tpApi.createTPSet(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: tpKeys.sets() });
    },
    ...options,
  });
};

/**
 * Approve TP Set mutation
 */
export const useApproveTPSet = (
  options?: Omit<UseMutationOptions<any, Error, string>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => tpApi.approveTPSet(id),
    onSuccess: (_, id) => {
      queryClient.invalidateQueries({ queryKey: tpKeys.setDetail(id) });
      queryClient.invalidateQueries({ queryKey: tpKeys.sets() });
    },
    ...options,
  });
};

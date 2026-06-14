/**
 * TP Command Service
 * Provides command operations for TP data using TanStack Query mutations with proper types
 */

import { useMutation, useQueryClient, UseMutationOptions } from '@tanstack/react-query';
import * as tpApi from '@/api/tp';
import { tpKeys } from '../queries/TPQueryService';
import { TP, TPSet, CreateTPRequest, UpdateTPRequest, CreateTPSetRequest } from '@/shared/types/domain';

/**
 * Create TP mutation
 */
export const useCreateTP = (
  options?: Omit<UseMutationOptions<TP, Error, CreateTPRequest>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation<TP, Error, CreateTPRequest>({
    mutationFn: (data: CreateTPRequest) => tpApi.createTP(data),
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
  options?: Omit<UseMutationOptions<TP, Error, { id: string; data: UpdateTPRequest }>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation<TP, Error, { id: string; data: UpdateTPRequest }>({
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

  return useMutation<void, Error, string>({
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
  options?: Omit<UseMutationOptions<TPSet, Error, CreateTPSetRequest>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation<TPSet, Error, CreateTPSetRequest>({
    mutationFn: (data: CreateTPSetRequest) => tpApi.createTPSet(data),
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
  options?: Omit<UseMutationOptions<TPSet, Error, string>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation<TPSet, Error, string>({
    mutationFn: (id: string) => tpApi.approveTPSet(id),
    onSuccess: (_, id) => {
      queryClient.invalidateQueries({ queryKey: tpKeys.setDetail(id) });
      queryClient.invalidateQueries({ queryKey: tpKeys.sets() });
    },
    ...options,
  });
};

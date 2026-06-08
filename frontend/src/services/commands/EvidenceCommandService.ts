/**
 * Evidence Command Service
 * Provides command operations for Evidence data using TanStack Query mutations
 */

import { useMutation, useQueryClient, UseMutationOptions } from '@tanstack/react-query';
import * as evidenceApi from '@/api/evidence';
import { evidenceKeys } from '../queries/EvidenceQueryService';

/**
 * Create Evidence mutation
 */
export const useCreateEvidence = (
  options?: Omit<UseMutationOptions<any, Error, { data: any; userId: string }>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ data, userId }) => evidenceApi.createEvidence(data, userId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: evidenceKeys.all });
    },
    ...options,
  });
};

/**
 * Update Evidence mutation
 */
export const useUpdateEvidence = (
  options?: Omit<UseMutationOptions<any, Error, { id: string; data: any }>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, data }) => evidenceApi.updateEvidence(id, data),
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({ queryKey: evidenceKeys.detail(variables.id) });
      queryClient.invalidateQueries({ queryKey: evidenceKeys.all });
    },
    ...options,
  });
};

/**
 * Delete Evidence mutation
 */
export const useDeleteEvidence = (
  options?: Omit<UseMutationOptions<void, Error, string>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => evidenceApi.deleteEvidence(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: evidenceKeys.all });
    },
    ...options,
  });
};

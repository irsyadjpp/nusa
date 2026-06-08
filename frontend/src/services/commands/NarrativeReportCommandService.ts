/**
 * Narrative Report Command Service
 * Provides command operations for Narrative Report data using TanStack Query mutations
 */

import { useMutation, useQueryClient, UseMutationOptions } from '@tanstack/react-query';
import * as narrativeReportApi from '@/api/narrative-report';
import { narrativeReportKeys } from '../queries/NarrativeReportQueryService';

/**
 * Create Narrative Report mutation
 */
export const useCreateNarrativeReport = (
  options?: Omit<UseMutationOptions<any, Error, any>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: any) => narrativeReportApi.createNarrativeReport(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: narrativeReportKeys.all });
    },
    ...options,
  });
};

/**
 * Update Narrative Report mutation
 */
export const useUpdateNarrativeReport = (
  options?: Omit<UseMutationOptions<any, Error, { id: string; data: any }>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, data }) => narrativeReportApi.updateNarrativeReport(id, data),
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({ queryKey: narrativeReportKeys.detail(variables.id) });
      queryClient.invalidateQueries({ queryKey: narrativeReportKeys.all });
    },
    ...options,
  });
};

/**
 * Delete Narrative Report mutation
 */
export const useDeleteNarrativeReport = (
  options?: Omit<UseMutationOptions<void, Error, string>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => narrativeReportApi.deleteNarrativeReport(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: narrativeReportKeys.all });
    },
    ...options,
  });
};

/**
 * Publish Narrative Report mutation
 */
export const usePublishNarrativeReport = (
  options?: Omit<UseMutationOptions<any, Error, string>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => narrativeReportApi.publishNarrativeReport(id),
    onSuccess: (_, id) => {
      queryClient.invalidateQueries({ queryKey: narrativeReportKeys.detail(id) });
      queryClient.invalidateQueries({ queryKey: narrativeReportKeys.all });
    },
    ...options,
  });
};

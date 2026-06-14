/**
 * Modul Ajar Command Service
 * Provides command operations for Modul Ajar data using TanStack Query mutations
 */

import { useMutation, useQueryClient, UseMutationOptions } from '@tanstack/react-query';
import * as modulAjarApi from '@/api/modul-ajar';
import { modulAjarKeys } from '../queries/ModulAjarQueryService';

/**
 * Create Modul Ajar mutation
 */
export const useCreateModulAjar = (
  options?: Omit<UseMutationOptions<any, Error, any>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: any) => modulAjarApi.createModulAjar(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: modulAjarKeys.all });
    },
    ...options,
  });
};

/**
 * Update Modul Ajar mutation
 */
export const useUpdateModulAjar = (
  options?: Omit<UseMutationOptions<any, Error, { id: string; data: any }>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, data }) => modulAjarApi.updateModulAjar(id, data),
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({ queryKey: modulAjarKeys.detail(variables.id) });
      queryClient.invalidateQueries({ queryKey: modulAjarKeys.all });
    },
    ...options,
  });
};

/**
 * Delete Modul Ajar mutation
 */
export const useDeleteModulAjar = (
  options?: Omit<UseMutationOptions<void, Error, string>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => modulAjarApi.deleteModulAjar(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: modulAjarKeys.all });
    },
    ...options,
  });
};

/**
 * Create Modul Ajar Set mutation
 */
export const useCreateModulAjarSet = (
  options?: Omit<UseMutationOptions<any, Error, any>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: any) => modulAjarApi.createModulAjarSet(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: modulAjarKeys.sets() });
    },
    ...options,
  });
};

/**
 * Update Modul Ajar Set mutation
 */
export const useUpdateModulAjarSet = (
  options?: Omit<UseMutationOptions<any, Error, { id: string; data: any }>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, data }) => modulAjarApi.updateModulAjarSet(id, data),
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({ queryKey: modulAjarKeys.setDetail(variables.id) });
      queryClient.invalidateQueries({ queryKey: modulAjarKeys.sets() });
    },
    ...options,
  });
};

/**
 * Delete Modul Ajar Set mutation
 */
export const useDeleteModulAjarSet = (
  options?: Omit<UseMutationOptions<void, Error, string>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => modulAjarApi.deleteModulAjarSet(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: modulAjarKeys.sets() });
    },
    ...options,
  });
};

/**
 * Approve Modul Ajar Set mutation
 */
export const useApproveModulAjarSet = (
  options?: Omit<UseMutationOptions<any, Error, string>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => modulAjarApi.approveModulAjarSet(id),
    onSuccess: (_, id) => {
      queryClient.invalidateQueries({ queryKey: modulAjarKeys.setDetail(id) });
      queryClient.invalidateQueries({ queryKey: modulAjarKeys.sets() });
    },
    ...options,
  });
};

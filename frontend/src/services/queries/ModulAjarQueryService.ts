/**
 * Modul Ajar Query Service
 * Provides query operations for Modul Ajar data using TanStack Query with proper types
 */

import { useQuery, UseQueryOptions } from '@tanstack/react-query';
import * as modulAjarApi from '@/api/modul-ajar';
import { ModulAjar, ModulAjarSet, PaginationParams, FilterParams } from '@/shared/types/domain';

// Query Keys
export const modulAjarKeys = {
  all: ['modul-ajar'] as const,
  list: (params?: PaginationParams & FilterParams & {
    modul_ajar_set_id?: string;
    atp_id?: string;
  }) => ['modul-ajar', 'list', params] as const,
  detail: (id: string) => ['modul-ajar', 'detail', id] as const,
  bySet: (setId: string) => ['modul-ajar', 'set', setId] as const,
  sets: (params?: PaginationParams & FilterParams & {
    atp_set_id?: string;
    subject_id?: string;
    phase_id?: string;
  }) => ['modul-ajar', 'sets', params] as const,
  setDetail: (id: string) => ['modul-ajar', 'set', 'detail', id] as const,
} as const;

/**
 * Get Modul Ajars list
 */
export const useModulAjars = (
  params?: PaginationParams & FilterParams & {
    modul_ajar_set_id?: string;
    atp_id?: string;
  },
  options?: Omit<UseQueryOptions<ModulAjar[], Error, ModulAjar[]>, 'queryKey' | 'queryFn'>
) => {
  return useQuery<ModulAjar[], Error, ModulAjar[]>({
    queryKey: modulAjarKeys.list(params),
    queryFn: () => modulAjarApi.getModulAjars(params),
    staleTime: 300000, // 5 minutes - Modul Ajar data changes infrequently
    ...options,
  });
};

/**
 * Get Modul Ajar by ID
 */
export const useModulAjar = (
  id: string,
  options?: Omit<UseQueryOptions<ModulAjar, Error, ModulAjar>, 'queryKey' | 'queryFn'>
) => {
  return useQuery<ModulAjar, Error, ModulAjar>({
    queryKey: modulAjarKeys.detail(id),
    queryFn: () => modulAjarApi.getModulAjarById(id),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

/**
 * Get Modul Ajars by Modul Ajar Set
 */
export const useModulAjarsBySet = (
  setId: string,
  options?: Omit<UseQueryOptions<ModulAjar[], Error, ModulAjar[]>, 'queryKey' | 'queryFn'>
) => {
  return useQuery<ModulAjar[], Error, ModulAjar[]>({
    queryKey: modulAjarKeys.bySet(setId),
    queryFn: () => modulAjarApi.getModulAjarsBySet(setId),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

/**
 * Get Modul Ajar Sets
 */
export const useModulAjarSets = (
  params?: PaginationParams & FilterParams & {
    atp_set_id?: string;
    subject_id?: string;
    phase_id?: string;
  },
  options?: Omit<UseQueryOptions<ModulAjarSet[], Error, ModulAjarSet[]>, 'queryKey' | 'queryFn'>
) => {
  return useQuery<ModulAjarSet[], Error, ModulAjarSet[]>({
    queryKey: modulAjarKeys.sets(params),
    queryFn: () => modulAjarApi.getModulAjarSets(params),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

/**
 * Get Modul Ajar Set by ID
 */
export const useModulAjarSet = (
  id: string,
  options?: Omit<UseQueryOptions<ModulAjarSet, Error, ModulAjarSet>, 'queryKey' | 'queryFn'>
) => {
  return useQuery<ModulAjarSet, Error, ModulAjarSet>({
    queryKey: modulAjarKeys.setDetail(id),
    queryFn: () => modulAjarApi.getModulAjarSetById(id),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

/**
 * Invalidate Modul Ajar queries
 */
export const invalidateModulAjarQueries = (queryClient: import('@tanstack/react-query').QueryClient) => {
  queryClient.invalidateQueries({ queryKey: modulAjarKeys.all });
};

/**
 * Invalidate Modul Ajar detail
 */
export const invalidateModulAjar = (queryClient: import('@tanstack/react-query').QueryClient, id: string) => {
  queryClient.invalidateQueries({ queryKey: modulAjarKeys.detail(id) });
};

/**
 * Invalidate Modul Ajar Set
 */
export const invalidateModulAjarSet = (queryClient: import('@tanstack/react-query').QueryClient, id: string) => {
  queryClient.invalidateQueries({ queryKey: modulAjarKeys.setDetail(id) });
};

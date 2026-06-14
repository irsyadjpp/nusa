/**
 * TP Query Service
 * Provides query operations for TP data using TanStack Query with proper types
 */

import { useQuery, UseQueryOptions } from '@tanstack/react-query';
import * as tpApi from '@/api/tp';
import { TP, TPSet, PaginationParams, FilterParams } from '@/shared/types/domain';

// Query Keys
export const tpKeys = {
  all: ['tp'] as const,
  list: (params?: PaginationParams & FilterParams & { tp_set_id?: string }) => ['tp', 'list', params] as const,
  detail: (id: string) => ['tp', 'detail', id] as const,
  bySet: (setId: string) => ['tp', 'set', setId] as const,
  sets: (params?: PaginationParams & FilterParams & { cp_id?: string }) => ['tp', 'sets', params] as const,
  setDetail: (id: string) => ['tp', 'set', 'detail', id] as const,
} as const;

/**
 * Get TPs list
 */
export const useTPs = (
  params?: PaginationParams & FilterParams & { tp_set_id?: string },
  options?: Omit<UseQueryOptions<TP[], Error, TP[]>, 'queryKey' | 'queryFn'>
) => {
  return useQuery<TP[], Error, TP[]>({
    queryKey: tpKeys.list(params),
    queryFn: () => tpApi.getTPs(params),
    staleTime: 300000, // 5 minutes - TP data changes infrequently
    ...options,
  });
};

/**
 * Get TP by ID
 */
export const useTP = (
  id: string,
  options?: Omit<UseQueryOptions<TP, Error, TP>, 'queryKey' | 'queryFn'>
) => {
  return useQuery<TP, Error, TP>({
    queryKey: tpKeys.detail(id),
    queryFn: () => tpApi.getTPById(id),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

/**
 * Get TPs by TP Set
 */
export const useTPsBySet = (
  setId: string,
  options?: Omit<UseQueryOptions<TP[], Error, TP[]>, 'queryKey' | 'queryFn'>
) => {
  return useQuery<TP[], Error, TP[]>({
    queryKey: tpKeys.bySet(setId),
    queryFn: () => tpApi.getTPsBySet(setId),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

/**
 * Get TP Sets
 */
export const useTPSets = (
  params?: PaginationParams & FilterParams & { cp_id?: string },
  options?: Omit<UseQueryOptions<TPSet[], Error, TPSet[]>, 'queryKey' | 'queryFn'>
) => {
  return useQuery<TPSet[], Error, TPSet[]>({
    queryKey: tpKeys.sets(params),
    queryFn: () => tpApi.getTPSets(params),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

/**
 * Get TP Set by ID
 */
export const useTPSet = (
  id: string,
  options?: Omit<UseQueryOptions<TPSet, Error, TPSet>, 'queryKey' | 'queryFn'>
) => {
  return useQuery<TPSet, Error, TPSet>({
    queryKey: tpKeys.setDetail(id),
    queryFn: () => tpApi.getTPSetById(id),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

/**
 * Invalidate TP queries
 */
export const invalidateTPQueries = (queryClient: import('@tanstack/react-query').QueryClient) => {
  queryClient.invalidateQueries({ queryKey: tpKeys.all });
};

/**
 * Invalidate TP detail
 */
export const invalidateTP = (queryClient: import('@tanstack/react-query').QueryClient, id: string) => {
  queryClient.invalidateQueries({ queryKey: tpKeys.detail(id) });
};

/**
 * Invalidate TP Set
 */
export const invalidateTPSet = (queryClient: import('@tanstack/react-query').QueryClient, id: string) => {
  queryClient.invalidateQueries({ queryKey: tpKeys.setDetail(id) });
};

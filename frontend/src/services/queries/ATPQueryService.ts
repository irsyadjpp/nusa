/**
 * ATP Query Service
 * Provides query operations for ATP data using TanStack Query with proper types
 */

import { useQuery, UseQueryOptions } from '@tanstack/react-query';
import * as atpApi from '@/api/atp';
import { ATP, ATPSet, PaginationParams, FilterParams } from '@/shared/types/domain';

// Query Keys
export const atpKeys = {
  all: ['atp'] as const,
  list: (params?: PaginationParams & FilterParams & { 
    atp_set_id?: string; 
    tp_id?: string;
  }) => ['atp', 'list', params] as const,
  detail: (id: string) => ['atp', 'detail', id] as const,
  bySet: (setId: string) => ['atp', 'set', setId] as const,
  sets: (params?: PaginationParams & FilterParams & { 
    tp_set_id?: string; 
    subject_id?: string; 
    phase_id?: string;
  }) => ['atp', 'sets', params] as const,
  setDetail: (id: string) => ['atp', 'set', 'detail', id] as const,
} as const;

/**
 * Get ATPs list
 */
export const useATPs = (
  params?: PaginationParams & FilterParams & { 
    atp_set_id?: string; 
    tp_id?: string;
  },
  options?: Omit<UseQueryOptions<ATP[], Error, ATP[]>, 'queryKey' | 'queryFn'>
) => {
  return useQuery<ATP[], Error, ATP[]>({
    queryKey: atpKeys.list(params),
    queryFn: () => atpApi.getATPs(params),
    staleTime: 300000, // 5 minutes - ATP data changes infrequently
    ...options,
  });
};

/**
 * Get ATP by ID
 */
export const useATP = (
  id: string,
  options?: Omit<UseQueryOptions<ATP, Error, ATP>, 'queryKey' | 'queryFn'>
) => {
  return useQuery<ATP, Error, ATP>({
    queryKey: atpKeys.detail(id),
    queryFn: () => atpApi.getATPById(id),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

/**
 * Get ATPs by ATP Set
 */
export const useATPsBySet = (
  setId: string,
  options?: Omit<UseQueryOptions<ATP[], Error, ATP[]>, 'queryKey' | 'queryFn'>
) => {
  return useQuery<ATP[], Error, ATP[]>({
    queryKey: atpKeys.bySet(setId),
    queryFn: () => atpApi.getATPsBySet(setId),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

/**
 * Get ATP Sets
 */
export const useATPSets = (
  params?: PaginationParams & FilterParams & { 
    tp_set_id?: string; 
    subject_id?: string; 
    phase_id?: string;
  },
  options?: Omit<UseQueryOptions<ATPSet[], Error, ATPSet[]>, 'queryKey' | 'queryFn'>
) => {
  return useQuery<ATPSet[], Error, ATPSet[]>({
    queryKey: atpKeys.sets(params),
    queryFn: () => atpApi.getATPSets(params),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

/**
 * Get ATP Set by ID
 */
export const useATPSet = (
  id: string,
  options?: Omit<UseQueryOptions<ATPSet, Error, ATPSet>, 'queryKey' | 'queryFn'>
) => {
  return useQuery<ATPSet, Error, ATPSet>({
    queryKey: atpKeys.setDetail(id),
    queryFn: () => atpApi.getATPSetById(id),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

/**
 * Invalidate ATP queries
 */
export const invalidateATPQueries = (queryClient: import('@tanstack/react-query').QueryClient) => {
  queryClient.invalidateQueries({ queryKey: atpKeys.all });
};

/**
 * Invalidate ATP detail
 */
export const invalidateATP = (queryClient: import('@tanstack/react-query').QueryClient, id: string) => {
  queryClient.invalidateQueries({ queryKey: atpKeys.detail(id) });
};

/**
 * Invalidate ATP Set
 */
export const invalidateATPSet = (queryClient: import('@tanstack/react-query').QueryClient, id: string) => {
  queryClient.invalidateQueries({ queryKey: atpKeys.setDetail(id) });
};

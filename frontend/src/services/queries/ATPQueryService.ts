/**
 * ATP Query Service
 * Provides query operations for ATP data using TanStack Query
 */

import { useQuery, UseQueryOptions } from '@tanstack/react-query';
import * as atpApi from '@/api/atp';

// Query Keys
export const atpKeys = {
  all: ['atp'] as const,
  list: (params?: any) => ['atp', 'list', params] as const,
  detail: (id: string) => ['atp', 'detail', id] as const,
  bySet: (setId: string) => ['atp', 'set', setId] as const,
  sets: (params?: any) => ['atp', 'sets', params] as const,
  setDetail: (id: string) => ['atp', 'set', 'detail', id] as const,
};

/**
 * Get ATPs list
 */
export const useATPs = (
  params?: any,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
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
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
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
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
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
  params?: any,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
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
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: atpKeys.setDetail(id),
    queryFn: () => atpApi.getATPSetById(id),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

/**
 * Invalidate ATP queries
 */
export const invalidateATPQueries = (queryClient: any) => {
  queryClient.invalidateQueries({ queryKey: atpKeys.all });
};

/**
 * Invalidate ATP detail
 */
export const invalidateATP = (queryClient: any, id: string) => {
  queryClient.invalidateQueries({ queryKey: atpKeys.detail(id) });
};

/**
 * Invalidate ATP Set
 */
export const invalidateATPSet = (queryClient: any, id: string) => {
  queryClient.invalidateQueries({ queryKey: atpKeys.setDetail(id) });
};

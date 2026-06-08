/**
 * TP Query Service
 * Provides query operations for TP data using TanStack Query
 */

import { useQuery, UseQueryOptions } from '@tanstack/react-query';
import * as tpApi from '@/api/tp';

// Query Keys
export const tpKeys = {
  all: ['tp'] as const,
  list: (params?: any) => ['tp', 'list', params] as const,
  detail: (id: string) => ['tp', 'detail', id] as const,
  bySet: (setId: string) => ['tp', 'set', setId] as const,
  sets: (params?: any) => ['tp', 'sets', params] as const,
  setDetail: (id: string) => ['tp', 'set', 'detail', id] as const,
};

/**
 * Get TPs list
 */
export const useTPs = (
  params?: any,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
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
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
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
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
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
  params?: any,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
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
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: tpKeys.setDetail(id),
    queryFn: () => tpApi.getTPSetById(id),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

/**
 * Invalidate TP queries
 */
export const invalidateTPQueries = (queryClient: any) => {
  queryClient.invalidateQueries({ queryKey: tpKeys.all });
};

/**
 * Invalidate TP detail
 */
export const invalidateTP = (queryClient: any, id: string) => {
  queryClient.invalidateQueries({ queryKey: tpKeys.detail(id) });
};

/**
 * Invalidate TP Set
 */
export const invalidateTPSet = (queryClient: any, id: string) => {
  queryClient.invalidateQueries({ queryKey: tpKeys.setDetail(id) });
};

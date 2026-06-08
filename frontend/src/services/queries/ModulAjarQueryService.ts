/**
 * Modul Ajar Query Service
 * Provides query operations for Modul Ajar data using TanStack Query
 */

import { useQuery, UseQueryOptions } from '@tanstack/react-query';
import * as modulAjarApi from '@/api/modul-ajar';

// Query Keys
export const modulAjarKeys = {
  all: ['modul-ajar'] as const,
  list: (params?: any) => ['modul-ajar', 'list', params] as const,
  detail: (id: string) => ['modul-ajar', 'detail', id] as const,
  bySet: (setId: string) => ['modul-ajar', 'set', setId] as const,
  sets: (params?: any) => ['modul-ajar', 'sets', params] as const,
  setDetail: (id: string) => ['modul-ajar', 'set', 'detail', id] as const,
};

/**
 * Get Modul Ajars list
 */
export const useModulAjars = (
  params?: any,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
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
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
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
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
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
  params?: any,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
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
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: modulAjarKeys.setDetail(id),
    queryFn: () => modulAjarApi.getModulAjarSetById(id),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

/**
 * Invalidate Modul Ajar queries
 */
export const invalidateModulAjarQueries = (queryClient: any) => {
  queryClient.invalidateQueries({ queryKey: modulAjarKeys.all });
};

/**
 * Invalidate Modul Ajar detail
 */
export const invalidateModulAjar = (queryClient: any, id: string) => {
  queryClient.invalidateQueries({ queryKey: modulAjarKeys.detail(id) });
};

/**
 * Invalidate Modul Ajar Set
 */
export const invalidateModulAjarSet = (queryClient: any, id: string) => {
  queryClient.invalidateQueries({ queryKey: modulAjarKeys.setDetail(id) });
};

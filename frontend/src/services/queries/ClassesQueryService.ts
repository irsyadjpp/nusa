/**
 * Classes Query Service
 * Provides query operations for class data using TanStack Query
 */

import { useQuery, UseQueryOptions } from '@tanstack/react-query';
import * as classesApi from '@/api/classes';

// Query Keys
export const classesKeys = {
  all: ['classes'] as const,
  list: (params?: any) => ['classes', 'list', params] as const,
  detail: (id: string) => ['classes', 'detail', id] as const,
};

/**
 * Class Queries
 */
export const useClasses = (
  params?: any,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: classesKeys.list(params),
    queryFn: () => classesApi.listClasses(params),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

export const useClass = (
  id: string,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: classesKeys.detail(id),
    queryFn: () => classesApi.getClass(id),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

/**
 * Query invalidation functions
 */
export const invalidateClassQueries = (queryClient: any) => {
  queryClient.invalidateQueries({ queryKey: classesKeys.all });
};

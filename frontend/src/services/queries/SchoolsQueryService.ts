/**
 * Schools Query Service
 * Provides query operations for school data using TanStack Query
 */

import { useQuery, UseQueryOptions } from '@tanstack/react-query';
import * as schoolsApi from '@/api/schools';

// Query Keys
export const schoolsKeys = {
  all: ['schools'] as const,
  list: (params?: any) => ['schools', 'list', params] as const,
  detail: (id: string) => ['schools', 'detail', id] as const,
};

/**
 * School Queries
 */
export const useSchools = (
  params?: any,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: schoolsKeys.list(params),
    queryFn: () => schoolsApi.listSchools(params),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

export const useSchool = (
  id: string,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: schoolsKeys.detail(id),
    queryFn: () => schoolsApi.getSchool(id),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

/**
 * Query invalidation functions
 */
export const invalidateSchoolQueries = (queryClient: any) => {
  queryClient.invalidateQueries({ queryKey: schoolsKeys.all });
};

/**
 * Rubric Query Service
 * Provides query operations for Rubric data using TanStack Query
 */

import { useQuery, UseQueryOptions } from '@tanstack/react-query';
import * as rubricApi from '@/api/rubric';

// Query Keys
export const rubricKeys = {
  all: ['rubric'] as const,
  list: (params?: any) => ['rubric', 'list', params] as const,
  detail: (id: string) => ['rubric', 'detail', id] as const,
  byType: (rubricType: string) => ['rubric', 'type', rubricType] as const,
};

/**
 * Get rubrics list
 */
export const useRubrics = (
  params?: any,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: rubricKeys.list(params),
    queryFn: () => rubricApi.getRubrics(params),
    staleTime: 300000, // 5 minutes - rubric data changes infrequently
    ...options,
  });
};

/**
 * Get rubric by ID
 */
export const useRubric = (
  id: string,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: rubricKeys.detail(id),
    queryFn: () => rubricApi.getRubricById(id),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

/**
 * Get rubrics by type
 */
export const useRubricsByType = (
  rubricType: string,
  params?: any,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: rubricKeys.byType(rubricType),
    queryFn: () => rubricApi.getRubricsByType(rubricType, params),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

/**
 * Invalidate rubric queries
 */
export const invalidateRubricQueries = (queryClient: any) => {
  queryClient.invalidateQueries({ queryKey: rubricKeys.all });
};

/**
 * Invalidate rubric detail
 */
export const invalidateRubric = (queryClient: any, id: string) => {
  queryClient.invalidateQueries({ queryKey: rubricKeys.detail(id) });
};

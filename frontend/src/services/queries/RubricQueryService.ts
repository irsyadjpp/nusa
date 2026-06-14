/**
 * Rubric Query Service
 * Provides query operations for Rubric data using TanStack Query with proper types
 */

import { useQuery, UseQueryOptions } from '@tanstack/react-query';
import * as rubricApi from '@/api/rubric';
import { Rubric, RubricType, PaginationParams, FilterParams } from '@/shared/types/domain';

// Query Keys
export const rubricKeys = {
  all: ['rubric'] as const,
  list: (params?: PaginationParams & FilterParams & {
    user_id?: string;
    assessment_id?: string;
    rubric_type?: RubricType;
  }) => ['rubric', 'list', params] as const,
  detail: (id: string) => ['rubric', 'detail', id] as const,
  byType: (rubricType: RubricType, params?: PaginationParams) => ['rubric', 'type', rubricType, params] as const,
} as const;

/**
 * Get rubrics list
 */
export const useRubrics = (
  params?: PaginationParams & FilterParams & {
    user_id?: string;
    assessment_id?: string;
    rubric_type?: RubricType;
  },
  options?: Omit<UseQueryOptions<Rubric[], Error, Rubric[]>, 'queryKey' | 'queryFn'>
) => {
  return useQuery<Rubric[], Error, Rubric[]>({
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
  options?: Omit<UseQueryOptions<Rubric, Error, Rubric>, 'queryKey' | 'queryFn'>
) => {
  return useQuery<Rubric, Error, Rubric>({
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
  rubricType: RubricType,
  params?: PaginationParams,
  options?: Omit<UseQueryOptions<Rubric[], Error, Rubric[]>, 'queryKey' | 'queryFn'>
) => {
  return useQuery<Rubric[], Error, Rubric[]>({
    queryKey: rubricKeys.byType(rubricType, params),
    queryFn: () => rubricApi.getRubricsByType(rubricType, params),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

/**
 * Invalidate rubric queries
 */
export const invalidateRubricQueries = (queryClient: import('@tanstack/react-query').QueryClient) => {
  queryClient.invalidateQueries({ queryKey: rubricKeys.all });
};

/**
 * Invalidate rubric detail
 */
export const invalidateRubric = (queryClient: import('@tanstack/react-query').QueryClient, id: string) => {
  queryClient.invalidateQueries({ queryKey: rubricKeys.detail(id) });
};

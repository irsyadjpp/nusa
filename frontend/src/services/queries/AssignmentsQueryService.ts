/**
 * Assignments Query Service
 * Provides query operations for assignment data using TanStack Query
 */

import { useQuery, UseQueryOptions } from '@tanstack/react-query';
import * as assignmentsApi from '@/api/assignments';

// Query Keys
export const assignmentsKeys = {
  all: ['assignments'] as const,
  list: (params?: any) => ['assignments', 'list', params] as const,
  class: (class_id: string) => ['assignments', 'class', class_id] as const,
};

/**
 * Assignment Queries
 */
export const useAssignments = (
  params?: any,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: assignmentsKeys.list(params),
    queryFn: () => assignmentsApi.listAssignments(params),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

export const useAssignmentsByClass = (
  class_id: string,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: assignmentsKeys.class(class_id),
    queryFn: () => assignmentsApi.getAssignmentsByClass(class_id),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

/**
 * Query invalidation functions
 */
export const invalidateAssignmentQueries = (queryClient: any) => {
  queryClient.invalidateQueries({ queryKey: assignmentsKeys.all });
};

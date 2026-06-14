/**
 * Schedules Query Service
 * Provides query operations for schedule data using TanStack Query
 */

import { useQuery, UseQueryOptions } from '@tanstack/react-query';
import * as schedulesApi from '@/api/schedules';

// Query Keys
export const schedulesKeys = {
  all: ['schedules'] as const,
  list: (params?: any) => ['schedules', 'list', params] as const,
  conflicts: (data: any) => ['schedules', 'conflicts', data] as const,
};

/**
 * Schedule Queries
 */
export const useSchedules = (
  params?: any,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: schedulesKeys.list(params),
    queryFn: () => schedulesApi.listSchedules(params),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

export const useScheduleConflicts = (
  data: any,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: schedulesKeys.conflicts(data),
    queryFn: () => schedulesApi.checkScheduleConflicts(data),
    staleTime: 0, // Always check for conflicts
    ...options,
  });
};

/**
 * Query invalidation functions
 */
export const invalidateScheduleQueries = (queryClient: any) => {
  queryClient.invalidateQueries({ queryKey: schedulesKeys.all });
};

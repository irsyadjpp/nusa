/**
 * Exams Query Service
 * Provides query operations for exam data using TanStack Query
 */

import { useQuery, UseQueryOptions } from '@tanstack/react-query';
import * as examsApi from '@/api/exams';

// Query Keys
export const examsKeys = {
  all: ['exams'] as const,
  list: (params?: any) => ['exams', 'list', params] as const,
  class: (class_id: string) => ['exams', 'class', class_id] as const,
};

/**
 * Exam Queries
 */
export const useExams = (
  params?: any,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: examsKeys.list(params),
    queryFn: () => examsApi.listExams(params),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

export const useExamsByClass = (
  class_id: string,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: examsKeys.class(class_id),
    queryFn: () => examsApi.getExamsByClass(class_id),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

/**
 * Query invalidation functions
 */
export const invalidateExamQueries = (queryClient: any) => {
  queryClient.invalidateQueries({ queryKey: examsKeys.all });
};

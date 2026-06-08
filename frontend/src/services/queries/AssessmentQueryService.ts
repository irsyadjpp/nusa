/**
 * Assessment Query Service
 * Provides query operations for Assessment data using TanStack Query
 */

import { useQuery, UseQueryOptions } from '@tanstack/react-query';
import * as assessmentApi from '@/api/assessment';

// Query Keys
export const assessmentKeys = {
  all: ['assessment'] as const,
  list: (params?: any) => ['assessment', 'list', params] as const,
  detail: (id: string) => ['assessment', 'detail', id] as const,
};

/**
 * Get assessments list
 */
export const useAssessments = (
  params?: any,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: assessmentKeys.list(params),
    queryFn: () => assessmentApi.getAssessments(params),
    staleTime: 60000, // 1 minute - assessment data changes moderately
    ...options,
  });
};

/**
 * Get assessment by ID
 */
export const useAssessment = (
  id: string,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: assessmentKeys.detail(id),
    queryFn: () => assessmentApi.getAssessmentById(id),
    staleTime: 300000, // 5 minutes - individual assessment data changes less frequently
    ...options,
  });
};

/**
 * Invalidate assessment queries
 */
export const invalidateAssessmentQueries = (queryClient: any) => {
  queryClient.invalidateQueries({ queryKey: assessmentKeys.all });
};

/**
 * Invalidate assessment detail
 */
export const invalidateAssessment = (queryClient: any, id: string) => {
  queryClient.invalidateQueries({ queryKey: assessmentKeys.detail(id) });
};

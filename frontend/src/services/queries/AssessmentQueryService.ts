/**
 * Assessment Query Service
 * Provides query operations for Assessment data using TanStack Query with proper types
 */

import { useQuery, UseQueryOptions } from '@tanstack/react-query';
import * as assessmentApi from '@/api/assessment';
import { Assessment, PaginationParams, FilterParams, AssessmentType } from '@/shared/types/domain';

// Query Keys
export const assessmentKeys = {
  all: ['assessment'] as const,
  list: (params?: PaginationParams & FilterParams & { 
    tp_id?: string; 
    user_id?: string; 
    assessment_type?: AssessmentType 
  }) => ['assessment', 'list', params] as const,
  detail: (id: string) => ['assessment', 'detail', id] as const,
} as const;

/**
 * Get assessments list
 */
export const useAssessments = (
  params?: PaginationParams & FilterParams & { 
    tp_id?: string; 
    user_id?: string; 
    assessment_type?: AssessmentType 
  },
  options?: Omit<UseQueryOptions<Assessment[], Error, Assessment[]>, 'queryKey' | 'queryFn'>
) => {
  return useQuery<Assessment[], Error, Assessment[]>({
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
  options?: Omit<UseQueryOptions<Assessment, Error, Assessment>, 'queryKey' | 'queryFn'>
) => {
  return useQuery<Assessment, Error, Assessment>({
    queryKey: assessmentKeys.detail(id),
    queryFn: () => assessmentApi.getAssessmentById(id),
    staleTime: 300000, // 5 minutes - individual assessment data changes less frequently
    ...options,
  });
};

/**
 * Invalidate assessment queries
 */
export const invalidateAssessmentQueries = (queryClient: import('@tanstack/react-query').QueryClient) => {
  queryClient.invalidateQueries({ queryKey: assessmentKeys.all });
};

/**
 * Invalidate assessment detail
 */
export const invalidateAssessment = (queryClient: import('@tanstack/react-query').QueryClient, id: string) => {
  queryClient.invalidateQueries({ queryKey: assessmentKeys.detail(id) });
};

/**
 * Narrative Report Query Service
 * Provides query operations for Narrative Report data using TanStack Query with proper types
 */

import { useQuery, UseQueryOptions } from '@tanstack/react-query';
import * as narrativeReportApi from '@/api/narrative-report';
import { NarrativeReport, StudentAchievement, PaginationParams, FilterParams } from '@/shared/types/domain';

// Query Keys
export const narrativeReportKeys = {
  all: ['narrative-report'] as const,
  list: (params?: PaginationParams & FilterParams & {
    student_id?: string;
    class_id?: string;
    subject_id?: string;
  }) => ['narrative-report', 'list', params] as const,
  detail: (id: string) => ['narrative-report', 'detail', id] as const,
  achievementSummary: (id: string, params?: PaginationParams) => ['narrative-report', id, 'achievement-summary', params] as const,
} as const;

/**
 * Get narrative reports list
 */
export const useNarrativeReports = (
  params?: PaginationParams & FilterParams & {
    student_id?: string;
    class_id?: string;
    subject_id?: string;
  },
  options?: Omit<UseQueryOptions<NarrativeReport[], Error, NarrativeReport[]>, 'queryKey' | 'queryFn'>
) => {
  return useQuery<NarrativeReport[], Error, NarrativeReport[]>({
    queryKey: narrativeReportKeys.list(params),
    queryFn: () => narrativeReportApi.getNarrativeReports(params),
    staleTime: 60000, // 1 minute - narrative report data changes moderately
    ...options,
  });
};

/**
 * Get narrative report by ID
 */
export const useNarrativeReport = (
  id: string,
  options?: Omit<UseQueryOptions<NarrativeReport, Error, NarrativeReport>, 'queryKey' | 'queryFn'>
) => {
  return useQuery<NarrativeReport, Error, NarrativeReport>({
    queryKey: narrativeReportKeys.detail(id),
    queryFn: () => narrativeReportApi.getNarrativeReportById(id),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

/**
 * Get achievement summary for report
 */
export const useAchievementSummary = (
  id: string,
  params?: PaginationParams,
  options?: Omit<UseQueryOptions<StudentAchievement, Error, StudentAchievement>, 'queryKey' | 'queryFn'>
) => {
  return useQuery<StudentAchievement, Error, StudentAchievement>({
    queryKey: narrativeReportKeys.achievementSummary(id, params),
    queryFn: () => narrativeReportApi.getAchievementSummary(id, params),
    staleTime: 30000, // 30 seconds - achievement summary changes frequently
    ...options,
  });
};

/**
 * Invalidate narrative report queries
 */
export const invalidateNarrativeReportQueries = (queryClient: import('@tanstack/react-query').QueryClient) => {
  queryClient.invalidateQueries({ queryKey: narrativeReportKeys.all });
};

/**
 * Invalidate narrative report detail
 */
export const invalidateNarrativeReport = (queryClient: import('@tanstack/react-query').QueryClient, id: string) => {
  queryClient.invalidateQueries({ queryKey: narrativeReportKeys.detail(id) });
};

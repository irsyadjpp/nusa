/**
 * Narrative Report Query Service
 * Provides query operations for Narrative Report data using TanStack Query
 */

import { useQuery, UseQueryOptions } from '@tanstack/react-query';
import * as narrativeReportApi from '@/api/narrative-report';

// Query Keys
export const narrativeReportKeys = {
  all: ['narrative-report'] as const,
  list: (params?: any) => ['narrative-report', 'list', params] as const,
  detail: (id: string) => ['narrative-report', 'detail', id] as const,
  achievementSummary: (id: string) => ['narrative-report', id, 'achievement-summary'] as const,
};

/**
 * Get narrative reports list
 */
export const useNarrativeReports = (
  params?: any,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
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
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
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
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: narrativeReportKeys.achievementSummary(id),
    queryFn: () => narrativeReportApi.getAchievementSummary(id),
    staleTime: 30000, // 30 seconds - achievement summary changes frequently
    ...options,
  });
};

/**
 * Invalidate narrative report queries
 */
export const invalidateNarrativeReportQueries = (queryClient: any) => {
  queryClient.invalidateQueries({ queryKey: narrativeReportKeys.all });
};

/**
 * Invalidate narrative report detail
 */
export const invalidateNarrativeReport = (queryClient: any, id: string) => {
  queryClient.invalidateQueries({ queryKey: narrativeReportKeys.detail(id) });
};

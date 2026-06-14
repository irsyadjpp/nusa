/**
 * Announcements Query Service
 * Provides query operations for announcement data using TanStack Query
 */

import { useQuery, UseQueryOptions } from '@tanstack/react-query';
import * as announcementsApi from '@/api/announcements';

// Query Keys
export const announcementsKeys = {
  all: ['announcements'] as const,
  list: (params?: any) => ['announcements', 'list', params] as const,
  school: (school_id: string) => ['announcements', 'school', school_id] as const,
};

/**
 * Announcement Queries
 */
export const useAnnouncements = (
  params?: any,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: announcementsKeys.list(params),
    queryFn: () => announcementsApi.listAnnouncements(params),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

export const useSchoolAnnouncements = (
  school_id: string,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: announcementsKeys.school(school_id),
    queryFn: () => announcementsApi.getSchoolAnnouncements(school_id),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

/**
 * Query invalidation functions
 */
export const invalidateAnnouncementQueries = (queryClient: any) => {
  queryClient.invalidateQueries({ queryKey: announcementsKeys.all });
};

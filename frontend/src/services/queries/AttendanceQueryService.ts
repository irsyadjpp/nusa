/**
 * Attendance Query Service
 * Provides query operations for attendance data using TanStack Query
 */

import { useQuery, UseQueryOptions } from '@tanstack/react-query';
import * as attendanceApi from '@/api/attendance';

// Query Keys
export const attendanceKeys = {
  all: ['attendance'] as const,
  list: (params?: any) => ['attendance', 'list', params] as const,
  report: (class_id: string, date: string) => ['attendance', 'report', class_id, date] as const,
};

/**
 * Attendance Queries
 */
export const useAttendance = (
  params?: any,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: attendanceKeys.list(params),
    queryFn: () => attendanceApi.listAttendance(params),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

export const useAttendanceReport = (
  class_id: string,
  date: string,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: attendanceKeys.report(class_id, date),
    queryFn: () => attendanceApi.getAttendanceReport(class_id, date),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

/**
 * Query invalidation functions
 */
export const invalidateAttendanceQueries = (queryClient: any) => {
  queryClient.invalidateQueries({ queryKey: attendanceKeys.all });
};

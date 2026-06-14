/**
 * Attendance Command Service
 * Provides command operations for attendance data using TanStack Query mutations
 */

import { useMutation, useQueryClient, UseMutationOptions } from '@tanstack/react-query';
import * as attendanceApi from '@/api/attendance';
import { attendanceKeys } from '@/services/queries/AttendanceQueryService';

/**
 * Create Attendance Mutation
 */
export const useCreateAttendance = (
  options?: Omit<UseMutationOptions<any, Error, any>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: any) => attendanceApi.createAttendance(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: attendanceKeys.all });
    },
    ...options,
  });
};

/**
 * Bulk Create Attendance Mutation
 */
export const useBulkCreateAttendance = (
  options?: Omit<UseMutationOptions<any, Error, any>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: any) => attendanceApi.bulkCreateAttendance(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: attendanceKeys.all });
    },
    ...options,
  });
};

/**
 * Update Attendance Mutation
 */
export const useUpdateAttendance = (
  options?: Omit<UseMutationOptions<any, Error, any>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, data }: { id: string; data: any }) =>
      attendanceApi.updateAttendance(id, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: attendanceKeys.all });
    },
    ...options,
  });
};

/**
 * Achievement Query Service
 * Provides query operations for Achievement data using TanStack Query
 */

import { useQuery, UseQueryOptions } from '@tanstack/react-query';
import * as achievementApi from '@/api/achievement';

// Query Keys
export const achievementKeys = {
  all: ['achievement'] as const,
  student: (id: string) => ['achievement', 'student', id] as const,
  class: (id: string) => ['achievement', 'class', id] as const,
  studentProgress: (id: string) => ['achievement', 'student', id, 'progress'] as const,
  reportSummary: (id: string) => ['achievement', 'report', id, 'summary'] as const,
};

/**
 * Get student achievement
 */
export const useStudentAchievement = (
  studentId: string,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: achievementKeys.student(studentId),
    queryFn: () => achievementApi.getStudentAchievement(studentId),
    staleTime: 30000, // 30 seconds - achievement data changes frequently
    ...options,
  });
};

/**
 * Get class achievement
 */
export const useClassAchievement = (
  classId: string,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: achievementKeys.class(classId),
    queryFn: () => achievementApi.getClassAchievement(classId),
    staleTime: 30000, // 30 seconds
    ...options,
  });
};

/**
 * Get student progress
 */
export const useStudentProgress = (
  studentId: string,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: achievementKeys.studentProgress(studentId),
    queryFn: () => achievementApi.getStudentProgress(studentId),
    staleTime: 30000, // 30 seconds
    ...options,
  });
};

/**
 * Invalidate achievement queries
 */
export const invalidateAchievementQueries = (queryClient: any) => {
  queryClient.invalidateQueries({ queryKey: achievementKeys.all });
};

/**
 * Invalidate student achievement
 */
export const invalidateStudentAchievement = (queryClient: any, studentId: string) => {
  queryClient.invalidateQueries({ queryKey: achievementKeys.student(studentId) });
};

/**
 * Invalidate class achievement
 */
export const invalidateClassAchievement = (queryClient: any, classId: string) => {
  queryClient.invalidateQueries({ queryKey: achievementKeys.class(classId) });
};

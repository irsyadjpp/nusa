/**
 * Achievement Query Service
 * Provides query operations for Achievement data using TanStack Query with proper types
 */

import { useQuery, UseQueryOptions } from '@tanstack/react-query';
import * as achievementApi from '@/api/achievement';
import { StudentAchievement, ClassAchievement, CompetencyProgress, PaginationParams } from '@/shared/types/domain';

// Query Keys
export const achievementKeys = {
  all: ['achievement'] as const,
  student: (id: string, params?: PaginationParams & { tp_id?: string }) => ['achievement', 'student', id, params] as const,
  class: (id: string, params?: PaginationParams & { subject_id?: string }) => ['achievement', 'class', id, params] as const,
  studentProgress: (id: string, params?: PaginationParams & { subject_id?: string; phase_id?: string }) => ['achievement', 'student', id, 'progress', params] as const,
  reportSummary: (id: string, params?: PaginationParams) => ['achievement', 'report', id, 'summary', params] as const,
} as const;

/**
 * Get student achievement
 */
export const useStudentAchievement = (
  studentId: string,
  params?: PaginationParams & { tp_id?: string },
  options?: Omit<UseQueryOptions<StudentAchievement[], Error, StudentAchievement[]>, 'queryKey' | 'queryFn'>
) => {
  return useQuery<StudentAchievement[], Error, StudentAchievement[]>({
    queryKey: achievementKeys.student(studentId, params),
    queryFn: () => achievementApi.getStudentAchievement(studentId, params),
    staleTime: 30000, // 30 seconds - achievement data changes frequently
    ...options,
  });
};

/**
 * Get class achievement
 */
export const useClassAchievement = (
  classId: string,
  params?: PaginationParams & { subject_id?: string },
  options?: Omit<UseQueryOptions<ClassAchievement, Error, ClassAchievement>, 'queryKey' | 'queryFn'>
) => {
  return useQuery<ClassAchievement, Error, ClassAchievement>({
    queryKey: achievementKeys.class(classId, params),
    queryFn: () => achievementApi.getClassAchievement(classId, params),
    staleTime: 30000, // 30 seconds
    ...options,
  });
};

/**
 * Get student progress
 */
export const useStudentProgress = (
  studentId: string,
  params?: PaginationParams & { subject_id?: string; phase_id?: string },
  options?: Omit<UseQueryOptions<CompetencyProgress[], Error, CompetencyProgress[]>, 'queryKey' | 'queryFn'>
) => {
  return useQuery<CompetencyProgress[], Error, CompetencyProgress[]>({
    queryKey: achievementKeys.studentProgress(studentId, params),
    queryFn: () => achievementApi.getStudentProgress(studentId, params),
    staleTime: 30000, // 30 seconds
    ...options,
  });
};

/**
 * Invalidate achievement queries
 */
export const invalidateAchievementQueries = (queryClient: import('@tanstack/react-query').QueryClient) => {
  queryClient.invalidateQueries({ queryKey: achievementKeys.all });
};

/**
 * Invalidate student achievement
 */
export const invalidateStudentAchievement = (queryClient: import('@tanstack/react-query').QueryClient) => {
  queryClient.invalidateQueries({ queryKey: achievementKeys.all });
};

/**
 * Invalidate class achievement
 */
export const invalidateClassAchievement = (queryClient: import('@tanstack/react-query').QueryClient) => {
  queryClient.invalidateQueries({ queryKey: achievementKeys.all });
};

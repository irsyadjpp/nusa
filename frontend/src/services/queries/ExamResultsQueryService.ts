/**
 * Exam Results Query Service
 * Provides query operations for exam result data using TanStack Query
 */

import { useQuery, UseQueryOptions } from '@tanstack/react-query';
import * as examResultsApi from '@/api/exam-results';

// Query Keys
export const examResultsKeys = {
  all: ['exam-results'] as const,
  list: (params?: any) => ['exam-results', 'list', params] as const,
  exam: (exam_id: string) => ['exam-results', 'exam', exam_id] as const,
  student: (student_id: string) => ['exam-results', 'student', student_id] as const,
};

/**
 * Exam Result Queries
 */
export const useExamResults = (
  params?: any,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: examResultsKeys.list(params),
    queryFn: () => examResultsApi.listExamResults(params),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

export const useExamResultsByExam = (
  exam_id: string,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: examResultsKeys.exam(exam_id),
    queryFn: () => examResultsApi.getExamResultsByExam(exam_id),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

export const useExamResultsByStudent = (
  student_id: string,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: examResultsKeys.student(student_id),
    queryFn: () => examResultsApi.getExamResultsByStudent(student_id),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

/**
 * Query invalidation functions
 */
export const invalidateExamResultQueries = (queryClient: any) => {
  queryClient.invalidateQueries({ queryKey: examResultsKeys.all });
};

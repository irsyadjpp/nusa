/**
 * Evaluation Query Service
 * Provides query operations for Evaluation data using TanStack Query
 */

import { useQuery, UseQueryOptions } from '@tanstack/react-query';
import * as evaluationApi from '@/api/evaluation';

// Query Keys
export const evaluationKeys = {
  all: ['evaluation'] as const,
  list: (params?: any) => ['evaluation', 'list', params] as const,
  detail: (id: string) => ['evaluation', 'detail', id] as const,
  byEvidence: (evidenceId: string) => ['evaluation', 'evidence', evidenceId] as const,
  byStudent: (studentId: string) => ['evaluation', 'student', studentId] as const,
  history: (evidenceId: string) => ['evaluation', 'history', evidenceId] as const,
};

/**
 * Get evaluations list
 */
export const useEvaluations = (
  params?: any,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: evaluationKeys.list(params),
    queryFn: () => evaluationApi.getEvaluations(params),
    staleTime: 60000, // 1 minute - evaluation data changes moderately
    ...options,
  });
};

/**
 * Get evaluation by ID
 */
export const useEvaluation = (
  id: string,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: evaluationKeys.detail(id),
    queryFn: () => evaluationApi.getEvaluationById(id),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

/**
 * Get evaluations by evidence
 */
export const useEvaluationsByEvidence = (
  evidenceId: string,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: evaluationKeys.byEvidence(evidenceId),
    queryFn: () => evaluationApi.getEvaluationsByEvidence(evidenceId),
    staleTime: 60000, // 1 minute
    ...options,
  });
};

/**
 * Get evaluations by student
 */
export const useEvaluationsByStudent = (
  studentId: string,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: evaluationKeys.byStudent(studentId),
    queryFn: () => evaluationApi.getEvaluationsByStudent(studentId),
    staleTime: 60000, // 1 minute
    ...options,
  });
};

/**
 * Get evaluation history
 */
export const useEvaluationHistory = (
  evidenceId: string,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: evaluationKeys.history(evidenceId),
    queryFn: () => evaluationApi.getEvaluationHistory(evidenceId),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

/**
 * Invalidate evaluation queries
 */
export const invalidateEvaluationQueries = (queryClient: any) => {
  queryClient.invalidateQueries({ queryKey: evaluationKeys.all });
};

/**
 * Invalidate evaluation detail
 */
export const invalidateEvaluation = (queryClient: any, id: string) => {
  queryClient.invalidateQueries({ queryKey: evaluationKeys.detail(id) });
};

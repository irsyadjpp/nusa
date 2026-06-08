/**
 * Evidence Query Service
 * Provides query operations for Evidence data using TanStack Query
 */

import { useQuery, UseQueryOptions } from '@tanstack/react-query';
import * as evidenceApi from '@/api/evidence';

// Query Keys
export const evidenceKeys = {
  all: ['evidence'] as const,
  list: (params?: any) => ['evidence', 'list', params] as const,
  detail: (id: string) => ['evidence', 'detail', id] as const,
  byStudent: (studentId: string) => ['evidence', 'student', studentId] as const,
  byAssessment: (assessmentId: string) => ['evidence', 'assessment', assessmentId] as const,
};

/**
 * Get evidences list
 */
export const useEvidences = (
  params?: any,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: evidenceKeys.list(params),
    queryFn: () => evidenceApi.getEvidences(params),
    staleTime: 60000, // 1 minute - evidence data changes moderately
    ...options,
  });
};

/**
 * Get evidence by ID
 */
export const useEvidence = (
  id: string,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: evidenceKeys.detail(id),
    queryFn: () => evidenceApi.getEvidenceById(id),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

/**
 * Get evidences by student
 */
export const useEvidencesByStudent = (
  studentId: string,
  params?: any,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: evidenceKeys.byStudent(studentId),
    queryFn: () => evidenceApi.getEvidencesByStudent(studentId, params),
    staleTime: 60000, // 1 minute
    ...options,
  });
};

/**
 * Get evidences by assessment
 */
export const useEvidencesByAssessment = (
  assessmentId: string,
  params?: any,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: evidenceKeys.byAssessment(assessmentId),
    queryFn: () => evidenceApi.getEvidencesByAssessment(assessmentId, params),
    staleTime: 60000, // 1 minute
    ...options,
  });
};

/**
 * Invalidate evidence queries
 */
export const invalidateEvidenceQueries = (queryClient: any) => {
  queryClient.invalidateQueries({ queryKey: evidenceKeys.all });
};

/**
 * Invalidate evidence detail
 */
export const invalidateEvidence = (queryClient: any, id: string) => {
  queryClient.invalidateQueries({ queryKey: evidenceKeys.detail(id) });
};

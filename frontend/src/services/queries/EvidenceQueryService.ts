/**
 * Evidence Query Service
 * Provides query operations for Evidence data using TanStack Query with proper types
 */

import { useQuery, UseQueryOptions } from '@tanstack/react-query';
import * as evidenceApi from '@/api/evidence';
import { Evidence, EvidenceType, EvidenceStatus, PaginationParams, FilterParams } from '@/shared/types/domain';

// Query Keys
export const evidenceKeys = {
  all: ['evidence'] as const,
  list: (params?: PaginationParams & FilterParams & { 
    student_id?: string; 
    assessment_id?: string; 
    user_id?: string;
    evidence_type?: EvidenceType;
  }) => ['evidence', 'list', params] as const,
  detail: (id: string) => ['evidence', 'detail', id] as const,
  byStudent: (studentId: string, params?: PaginationParams & FilterParams & {
    assessment_id?: string;
    status?: EvidenceStatus;
  }) => ['evidence', 'student', studentId, params] as const,
  byAssessment: (assessmentId: string, params?: PaginationParams & FilterParams & {
    status?: EvidenceStatus;
  }) => ['evidence', 'assessment', assessmentId, params] as const,
} as const;

/**
 * Get evidences list
 */
export const useEvidences = (
  params?: PaginationParams & FilterParams & { 
    student_id?: string; 
    assessment_id?: string; 
    user_id?: string;
    evidence_type?: EvidenceType;
  },
  options?: Omit<UseQueryOptions<Evidence[], Error, Evidence[]>, 'queryKey' | 'queryFn'>
) => {
  return useQuery<Evidence[], Error, Evidence[]>({
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
  options?: Omit<UseQueryOptions<Evidence, Error, Evidence>, 'queryKey' | 'queryFn'>
) => {
  return useQuery<Evidence, Error, Evidence>({
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
  params?: PaginationParams & FilterParams & {
    assessment_id?: string;
    status?: EvidenceStatus;
  },
  options?: Omit<UseQueryOptions<Evidence[], Error, Evidence[]>, 'queryKey' | 'queryFn'>
) => {
  return useQuery<Evidence[], Error, Evidence[]>({
    queryKey: evidenceKeys.byStudent(studentId, params),
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
  params?: PaginationParams & FilterParams & {
    status?: EvidenceStatus;
  },
  options?: Omit<UseQueryOptions<Evidence[], Error, Evidence[]>, 'queryKey' | 'queryFn'>
) => {
  return useQuery<Evidence[], Error, Evidence[]>({
    queryKey: evidenceKeys.byAssessment(assessmentId, params),
    queryFn: () => evidenceApi.getEvidencesByAssessment(assessmentId, params),
    staleTime: 60000, // 1 minute
    ...options,
  });
};

/**
 * Invalidate evidence queries
 */
export const invalidateEvidenceQueries = (queryClient: import('@tanstack/react-query').QueryClient) => {
  queryClient.invalidateQueries({ queryKey: evidenceKeys.all });
};

/**
 * Invalidate evidence detail
 */
export const invalidateEvidence = (queryClient: import('@tanstack/react-query').QueryClient, id: string) => {
  queryClient.invalidateQueries({ queryKey: evidenceKeys.detail(id) });
};

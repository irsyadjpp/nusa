/**
 * CP Query Service using TanStack Query with proper types
 * Handles all CP-related data fetching
 */

import { useQuery, UseQueryOptions } from '@tanstack/react-query';
import { getCPs, getCPById, getSubjects, getPhases, getElementsByPhase, getSubelementsByElement } from '@/api/cp';
import { CP, CurriculumSubject, CurriculumPhase, CurriculumElement, CurriculumSubelement, PaginationParams, FilterParams } from '@/shared/types/domain';

// Query Keys
export const cpKeys = {
  all: ['cp'] as const,
  lists: () => [...cpKeys.all, 'list'] as const,
  list: (filters?: PaginationParams & FilterParams & {
    subject_id?: string;
    phase_id?: string;
    element_id?: string;
  }) => [...cpKeys.lists(), filters] as const,
  details: () => [...cpKeys.all, 'detail'] as const,
  detail: (id: string) => [...cpKeys.details(), id] as const,
  subjects: () => ['subjects'] as const,
  phases: () => ['phases'] as const,
  elements: (phaseId: string) => ['elements', phaseId] as const,
  subelements: (elementId: string) => ['subelements', elementId] as const,
} as const;

// Query Hooks
export const useCPs = (
  filters?: PaginationParams & FilterParams & {
    subject_id?: string;
    phase_id?: string;
    element_id?: string;
  },
  options?: Omit<UseQueryOptions<CP[], Error, CP[]>, 'queryKey' | 'queryFn'>
) => {
  return useQuery<CP[], Error, CP[]>({
    queryKey: cpKeys.list(filters),
    queryFn: () => getCPs(filters),
    staleTime: 5 * 60 * 1000, // 5 minutes
    ...options,
  });
};

export const useCPById = (
  id: string,
  options?: Omit<UseQueryOptions<CP, Error, CP>, 'queryKey' | 'queryFn'>
) => {
  return useQuery<CP, Error, CP>({
    queryKey: cpKeys.detail(id),
    queryFn: () => getCPById(id),
    enabled: !!id,
    ...options,
  });
};

export const useSubjects = (
  options?: Omit<UseQueryOptions<CurriculumSubject[], Error, CurriculumSubject[]>, 'queryKey' | 'queryFn'>
) => {
  return useQuery<CurriculumSubject[], Error, CurriculumSubject[]>({
    queryKey: cpKeys.subjects(),
    queryFn: () => getSubjects(),
    staleTime: 30 * 60 * 1000, // 30 minutes (curriculum data doesn't change often)
    ...options,
  });
};

export const usePhases = (
  options?: Omit<UseQueryOptions<CurriculumPhase[], Error, CurriculumPhase[]>, 'queryKey' | 'queryFn'>
) => {
  return useQuery<CurriculumPhase[], Error, CurriculumPhase[]>({
    queryKey: cpKeys.phases(),
    queryFn: () => getPhases(),
    staleTime: 30 * 60 * 1000, // 30 minutes
    ...options,
  });
};

export const useElementsByPhase = (
  phaseId: string,
  options?: Omit<UseQueryOptions<CurriculumElement[], Error, CurriculumElement[]>, 'queryKey' | 'queryFn'>
) => {
  return useQuery<CurriculumElement[], Error, CurriculumElement[]>({
    queryKey: cpKeys.elements(phaseId),
    queryFn: () => getElementsByPhase(phaseId),
    enabled: !!phaseId,
    ...options,
  });
};

export const useSubelementsByElement = (
  elementId: string,
  options?: Omit<UseQueryOptions<CurriculumSubelement[], Error, CurriculumSubelement[]>, 'queryKey' | 'queryFn'>
) => {
  return useQuery<CurriculumSubelement[], Error, CurriculumSubelement[]>({
    queryKey: cpKeys.subelements(elementId),
    queryFn: () => getSubelementsByElement(elementId),
    enabled: !!elementId,
    ...options,
  });
};

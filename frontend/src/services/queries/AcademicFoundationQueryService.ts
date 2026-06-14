/**
 * Academic Foundation Query Service
 * Provides query operations for academic foundation data using TanStack Query
 */

import { useQuery, UseQueryOptions } from '@tanstack/react-query';
import * as academicFoundationApi from '@/api/academic-foundation';

// Query Keys
export const academicFoundationKeys = {
  all: ['academic-foundation'] as const,
  academicYears: {
    all: ['academic-foundation', 'academic-years'] as const,
    list: (params?: any) => ['academic-foundation', 'academic-years', 'list', params] as const,
    detail: (id: string) => ['academic-foundation', 'academic-years', 'detail', id] as const,
  },
  semesters: {
    all: ['academic-foundation', 'semesters'] as const,
    list: (params?: any) => ['academic-foundation', 'semesters', 'list', params] as const,
    detail: (id: string) => ['academic-foundation', 'semesters', 'detail', id] as const,
  },
  subjectCategories: {
    all: ['academic-foundation', 'subject-categories'] as const,
    list: (params?: any) => ['academic-foundation', 'subject-categories', 'list', params] as const,
    detail: (id: string) => ['academic-foundation', 'subject-categories', 'detail', id] as const,
  },
  graduateProfileDimensions: {
    all: ['academic-foundation', 'graduate-profile-dimensions'] as const,
    list: (params?: any) => ['academic-foundation', 'graduate-profile-dimensions', 'list', params] as const,
    detail: (id: string) => ['academic-foundation', 'graduate-profile-dimensions', 'detail', id] as const,
  },
  cpAlignments: {
    all: ['academic-foundation', 'cp-alignments'] as const,
    list: (params?: any) => ['academic-foundation', 'cp-alignments', 'list', params] as const,
    detail: (id: string) => ['academic-foundation', 'cp-alignments', 'detail', id] as const,
  },
  systemConfigurations: {
    all: ['academic-foundation', 'system-configurations'] as const,
    list: (params?: any) => ['academic-foundation', 'system-configurations', 'list', params] as const,
    detail: (key: string) => ['academic-foundation', 'system-configurations', 'detail', key] as const,
  },
};

/**
 * Academic Year Queries
 */
export const useAcademicYears = (
  params?: any,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: academicFoundationKeys.academicYears.list(params),
    queryFn: () => academicFoundationApi.getAcademicYears(params),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

export const useAcademicYear = (
  id: string,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: academicFoundationKeys.academicYears.detail(id),
    queryFn: () => academicFoundationApi.getAcademicYearById(id),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

/**
 * Semester Queries
 */
export const useSemesters = (
  params?: any,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: academicFoundationKeys.semesters.list(params),
    queryFn: () => academicFoundationApi.getSemesters(params),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

export const useSemester = (
  id: string,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: academicFoundationKeys.semesters.detail(id),
    queryFn: () => academicFoundationApi.getSemesterById(id),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

/**
 * Subject Category Queries
 */
export const useSubjectCategories = (
  params?: any,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: academicFoundationKeys.subjectCategories.list(params),
    queryFn: () => academicFoundationApi.getSubjectCategories(params),
    staleTime: 600000, // 10 minutes - categories change infrequently
    ...options,
  });
};

export const useSubjectCategory = (
  id: string,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: academicFoundationKeys.subjectCategories.detail(id),
    queryFn: () => academicFoundationApi.getSubjectCategoryById(id),
    staleTime: 600000, // 10 minutes
    ...options,
  });
};

/**
 * Graduate Profile Dimension Queries
 */
export const useGraduateProfileDimensions = (
  params?: any,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: academicFoundationKeys.graduateProfileDimensions.list(params),
    queryFn: () => academicFoundationApi.getGraduateProfileDimensions(params),
    staleTime: 600000, // 10 minutes - profile dimensions change infrequently
    ...options,
  });
};

export const useGraduateProfileDimension = (
  id: string,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: academicFoundationKeys.graduateProfileDimensions.detail(id),
    queryFn: () => academicFoundationApi.getGraduateProfileDimensionById(id),
    staleTime: 600000, // 10 minutes
    ...options,
  });
};

/**
 * CP Alignment Queries
 */
export const useCPAlignments = (
  params?: any,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: academicFoundationKeys.cpAlignments.list(params),
    queryFn: () => academicFoundationApi.getCPAlignments(params),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

export const useCPAlignment = (
  id: string,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: academicFoundationKeys.cpAlignments.detail(id),
    queryFn: () => academicFoundationApi.getCPAlignmentById(id),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

/**
 * System Configuration Queries
 */
export const useSystemConfigurations = (
  params?: any,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: academicFoundationKeys.systemConfigurations.list(params),
    queryFn: () => academicFoundationApi.getSystemConfigurations(params),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

export const useSystemConfiguration = (
  key: string,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: academicFoundationKeys.systemConfigurations.detail(key),
    queryFn: () => academicFoundationApi.getSystemConfigurationByKey(key),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

/**
 * Query invalidation functions
 */
export const invalidateAcademicFoundationQueries = (queryClient: any) => {
  queryClient.invalidateQueries({ queryKey: academicFoundationKeys.all });
};

export const invalidateAcademicYearQueries = (queryClient: any) => {
  queryClient.invalidateQueries({ queryKey: academicFoundationKeys.academicYears.all });
};

export const invalidateSemesterQueries = (queryClient: any) => {
  queryClient.invalidateQueries({ queryKey: academicFoundationKeys.semesters.all });
};

export const invalidateSubjectCategoryQueries = (queryClient: any) => {
  queryClient.invalidateQueries({ queryKey: academicFoundationKeys.subjectCategories.all });
};

export const invalidateGraduateProfileDimensionQueries = (queryClient: any) => {
  queryClient.invalidateQueries({ queryKey: academicFoundationKeys.graduateProfileDimensions.all });
};

export const invalidateCPAlignmentQueries = (queryClient: any) => {
  queryClient.invalidateQueries({ queryKey: academicFoundationKeys.cpAlignments.all });
};

export const invalidateSystemConfigurationQueries = (queryClient: any) => {
  queryClient.invalidateQueries({ queryKey: academicFoundationKeys.systemConfigurations.all });
};
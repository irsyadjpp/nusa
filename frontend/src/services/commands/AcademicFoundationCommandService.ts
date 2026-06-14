/**
 * Academic Foundation Command Service
 * Provides command operations for academic foundation data using TanStack Query mutations
 */

import { useMutation, useQueryClient, UseMutationOptions } from '@tanstack/react-query';
import * as academicFoundationApi from '@/api/academic-foundation';
import { academicFoundationKeys } from '../queries/AcademicFoundationQueryService';

/**
 * Academic Year Mutations
 */
export const useCreateAcademicYear = (
  options?: Omit<UseMutationOptions<any, Error, any>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: any) => academicFoundationApi.createAcademicYear(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: academicFoundationKeys.academicYears.all });
    },
    ...options,
  });
};

export const useUpdateAcademicYear = (
  options?: Omit<UseMutationOptions<any, Error, { id: string; data: any }>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, data }) => academicFoundationApi.updateAcademicYear(id, data),
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({ queryKey: academicFoundationKeys.academicYears.detail(variables.id) });
      queryClient.invalidateQueries({ queryKey: academicFoundationKeys.academicYears.all });
    },
    ...options,
  });
};

export const useActivateAcademicYear = (
  options?: Omit<UseMutationOptions<any, Error, string>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => academicFoundationApi.activateAcademicYear(id),
    onSuccess: (_, id) => {
      queryClient.invalidateQueries({ queryKey: academicFoundationKeys.academicYears.detail(id) });
      queryClient.invalidateQueries({ queryKey: academicFoundationKeys.academicYears.all });
    },
    ...options,
  });
};

export const useArchiveAcademicYear = (
  options?: Omit<UseMutationOptions<any, Error, string>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => academicFoundationApi.archiveAcademicYear(id),
    onSuccess: (_, id) => {
      queryClient.invalidateQueries({ queryKey: academicFoundationKeys.academicYears.detail(id) });
      queryClient.invalidateQueries({ queryKey: academicFoundationKeys.academicYears.all });
    },
    ...options,
  });
};

/**
 * Semester Mutations
 */
export const useCreateSemester = (
  options?: Omit<UseMutationOptions<any, Error, any>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: any) => academicFoundationApi.createSemester(data),
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({ queryKey: academicFoundationKeys.semesters.all });
      if (variables.academic_year_id) {
        queryClient.invalidateQueries({ queryKey: academicFoundationKeys.academicYears.detail(variables.academic_year_id) });
      }
    },
    ...options,
  });
};

export const useUpdateSemester = (
  options?: Omit<UseMutationOptions<any, Error, { id: string; data: any }>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, data }) => academicFoundationApi.updateSemester(id, data),
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({ queryKey: academicFoundationKeys.semesters.detail(variables.id) });
      queryClient.invalidateQueries({ queryKey: academicFoundationKeys.semesters.all });
    },
    ...options,
  });
};

export const useDeleteSemester = (
  options?: Omit<UseMutationOptions<void, Error, string>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => academicFoundationApi.deleteSemester(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: academicFoundationKeys.semesters.all });
    },
    ...options,
  });
};

/**
 * Subject Category Mutations
 */
export const useCreateSubjectCategory = (
  options?: Omit<UseMutationOptions<any, Error, any>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: any) => academicFoundationApi.createSubjectCategory(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: academicFoundationKeys.subjectCategories.all });
    },
    ...options,
  });
};

export const useUpdateSubjectCategory = (
  options?: Omit<UseMutationOptions<any, Error, { id: string; data: any }>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, data }) => academicFoundationApi.updateSubjectCategory(id, data),
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({ queryKey: academicFoundationKeys.subjectCategories.detail(variables.id) });
      queryClient.invalidateQueries({ queryKey: academicFoundationKeys.subjectCategories.all });
    },
    ...options,
  });
};

export const useDeleteSubjectCategory = (
  options?: Omit<UseMutationOptions<void, Error, string>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => academicFoundationApi.deleteSubjectCategory(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: academicFoundationKeys.subjectCategories.all });
    },
    ...options,
  });
};

/**
 * Graduate Profile Dimension Mutations
 */
export const useCreateGraduateProfileDimension = (
  options?: Omit<UseMutationOptions<any, Error, any>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: any) => academicFoundationApi.createGraduateProfileDimension(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: academicFoundationKeys.graduateProfileDimensions.all });
    },
    ...options,
  });
};

export const useUpdateGraduateProfileDimension = (
  options?: Omit<UseMutationOptions<any, Error, { id: string; data: any }>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, data }) => academicFoundationApi.updateGraduateProfileDimension(id, data),
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({ queryKey: academicFoundationKeys.graduateProfileDimensions.detail(variables.id) });
      queryClient.invalidateQueries({ queryKey: academicFoundationKeys.graduateProfileDimensions.all });
    },
    ...options,
  });
};

export const useDeleteGraduateProfileDimension = (
  options?: Omit<UseMutationOptions<void, Error, string>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => academicFoundationApi.deleteGraduateProfileDimension(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: academicFoundationKeys.graduateProfileDimensions.all });
    },
    ...options,
  });
};

/**
 * CP Alignment Mutations
 */
export const useCreateCPAlignment = (
  options?: Omit<UseMutationOptions<any, Error, any>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: any) => academicFoundationApi.createCPAlignment(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: academicFoundationKeys.cpAlignments.all });
    },
    ...options,
  });
};

export const useUpdateCPAlignment = (
  options?: Omit<UseMutationOptions<any, Error, { id: string; data: any }>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, data }) => academicFoundationApi.updateCPAlignment(id, data),
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({ queryKey: academicFoundationKeys.cpAlignments.detail(variables.id) });
      queryClient.invalidateQueries({ queryKey: academicFoundationKeys.cpAlignments.all });
    },
    ...options,
  });
};

export const useDeleteCPAlignment = (
  options?: Omit<UseMutationOptions<void, Error, string>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => academicFoundationApi.deleteCPAlignment(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: academicFoundationKeys.cpAlignments.all });
    },
    ...options,
  });
};

/**
 * System Configuration Mutations
 */
export const useUpdateSystemConfiguration = (
  options?: Omit<UseMutationOptions<any, Error, { key: string; data: any }>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ key, data }) => academicFoundationApi.updateSystemConfiguration(key, data),
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({ queryKey: academicFoundationKeys.systemConfigurations.detail(variables.key) });
      queryClient.invalidateQueries({ queryKey: academicFoundationKeys.systemConfigurations.all });
    },
    ...options,
  });
};
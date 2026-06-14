/**
 * Academic Foundation API Client
 * Handles all academic foundation-related API calls
 */

import apiClient, { handleApiError } from './client';

// Types
export interface AcademicYear {
  id: string;
  school_id: string;
  name: string;
  start_date: string;
  end_date: string;
  status: string;
  current_semester_id?: string;
  created_by: string;
  created_at: string;
  updated_at: string;
}

export interface Semester {
  id: string;
  academic_year_id: string;
  type: string;
  name: string;
  start_date: string;
  end_date: string;
  status: string;
  sequence_number: number;
  created_by: string;
  created_at: string;
  updated_at: string;
}

export interface SubjectCategory {
  id: string;
  code: string;
  name: string;
  name_indonesian: string;
  description?: string;
  parent_id?: string;
  level: number;
  sort_order: number;
  is_active: boolean;
  kurikulum_version: string;
  created_by: string;
  created_at: string;
  updated_at: string;
}

export interface GraduateProfileDimension {
  id: string;
  code: string;
  name: string;
  name_indonesian: string;
  description?: string;
  element_name?: string;
  element_name_indonesian?: string;
  phase_level?: string;
  phase_level_indonesian?: string;
  sort_order: number;
  is_active: boolean;
  created_by: string;
  created_at: string;
  updated_at: string;
}

export interface CPAlignment {
  id: string;
  subject_category_id: string;
  graduate_profile_dimension_id: string;
  alignment_strength: string;
  justification?: string;
  created_by: string;
  created_at: string;
  updated_at: string;
}

export interface SystemConfiguration {
  id: string;
  key: string;
  value: string;
  type: string;
  category: string;
  description?: string;
  is_public: boolean;
  school_id?: string;
  created_by: string;
  created_at: string;
  updated_at: string;
}

// Request Types
export interface CreateAcademicYearRequest {
  school_id: string;
  name: string;
  start_date: string;
  end_date: string;
}

export interface UpdateAcademicYearRequest {
  name?: string;
  start_date?: string;
  end_date?: string;
}

export interface CreateSemesterRequest {
  academic_year_id: string;
  type: string;
  name: string;
  start_date: string;
  end_date: string;
  sequence_number: number;
  sequence?: number; // Convenience property matching component expectations
  is_active?: boolean; // Convenience property matching component expectations
}

export interface UpdateSemesterRequest {
  name?: string;
  start_date?: string;
  end_date?: string;
  status?: string;
}

export interface CreateSubjectCategoryRequest {
  code: string;
  name: string;
  name_indonesian: string;
  description?: string;
  parent_id?: string;
  level: number;
  sort_order: number;
  kurikulum_version: string;
}

export interface UpdateSubjectCategoryRequest {
  code?: string;
  name?: string;
  name_indonesian?: string;
  description?: string;
  parent_id?: string;
  level?: number;
  sort_order?: number;
  is_active?: boolean;
}

export interface CreateGraduateProfileDimensionRequest {
  code: string;
  name: string;
  name_indonesian: string;
  description?: string;
  element_name?: string;
  element_name_indonesian?: string;
  phase_level?: string;
  phase_level_indonesian?: string;
  sort_order: number;
}

export interface UpdateGraduateProfileDimensionRequest {
  code?: string;
  name?: string;
  name_indonesian?: string;
  description?: string;
  element_name?: string;
  element_name_indonesian?: string;
  phase_level?: string;
  phase_level_indonesian?: string;
  sort_order?: number;
  is_active?: boolean;
}

export interface CreateCPAlignmentRequest {
  subject_category_id: string;
  graduate_profile_dimension_id: string;
  alignment_strength: string;
  justification?: string;
}

export interface UpdateCPAlignmentRequest {
  alignment_strength?: string;
  justification?: string;
}

export interface UpdateSystemConfigurationRequest {
  value: string;
}

/**
 * Academic Year API endpoints
 */
export const getAcademicYears = async (params?: {
  school_id?: string;
  status?: string;
  limit?: number;
  offset?: number;
}): Promise<AcademicYear[]> => {
  try {
    const response = await apiClient.get('/academic-years', { params });
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

export const getAcademicYearById = async (id: string): Promise<AcademicYear> => {
  try {
    const response = await apiClient.get(`/academic-years/${id}`);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

export const createAcademicYear = async (data: CreateAcademicYearRequest): Promise<AcademicYear> => {
  try {
    const response = await apiClient.post('/academic-years', data);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

export const updateAcademicYear = async (id: string, data: UpdateAcademicYearRequest): Promise<AcademicYear> => {
  try {
    const response = await apiClient.put(`/academic-years/${id}`, data);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

export const activateAcademicYear = async (id: string): Promise<AcademicYear> => {
  try {
    const response = await apiClient.post(`/academic-years/${id}/activate`);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

export const archiveAcademicYear = async (id: string): Promise<AcademicYear> => {
  try {
    const response = await apiClient.post(`/academic-years/${id}/archive`);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Semester API endpoints
 */
export const getSemesters = async (params?: {
  academic_year_id?: string;
  status?: string;
  type?: string;
  limit?: number;
  offset?: number;
}): Promise<Semester[]> => {
  try {
    const response = await apiClient.get('/semesters', { params });
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

export const getSemesterById = async (id: string): Promise<Semester> => {
  try {
    const response = await apiClient.get(`/semesters/${id}`);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

export const createSemester = async (data: CreateSemesterRequest): Promise<Semester> => {
  try {
    const response = await apiClient.post('/semesters', data);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

export const updateSemester = async (id: string, data: UpdateSemesterRequest): Promise<Semester> => {
  try {
    const response = await apiClient.put(`/semesters/${id}`, data);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

export const deleteSemester = async (id: string): Promise<void> => {
  try {
    await apiClient.delete(`/semesters/${id}`);
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Subject Category API endpoints
 */
export const getSubjectCategories = async (params?: {
  parent_id?: string;
  level?: number;
  is_active?: boolean;
  kurikulum_version?: string;
  limit?: number;
  offset?: number;
}): Promise<SubjectCategory[]> => {
  try {
    const response = await apiClient.get('/subject-categories', { params });
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

export const getSubjectCategoryById = async (id: string): Promise<SubjectCategory> => {
  try {
    const response = await apiClient.get(`/subject-categories/${id}`);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

export const createSubjectCategory = async (data: CreateSubjectCategoryRequest): Promise<SubjectCategory> => {
  try {
    const response = await apiClient.post('/subject-categories', data);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

export const updateSubjectCategory = async (id: string, data: UpdateSubjectCategoryRequest): Promise<SubjectCategory> => {
  try {
    const response = await apiClient.put(`/subject-categories/${id}`, data);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

export const deleteSubjectCategory = async (id: string): Promise<void> => {
  try {
    await apiClient.delete(`/subject-categories/${id}`);
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * Graduate Profile Dimension API endpoints
 */
export const getGraduateProfileDimensions = async (params?: {
  phase_level?: string;
  is_active?: boolean;
  limit?: number;
  offset?: number;
}): Promise<GraduateProfileDimension[]> => {
  try {
    const response = await apiClient.get('/graduate-profile-dimensions', { params });
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

export const getGraduateProfileDimensionById = async (id: string): Promise<GraduateProfileDimension> => {
  try {
    const response = await apiClient.get(`/graduate-profile-dimensions/${id}`);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

export const createGraduateProfileDimension = async (data: CreateGraduateProfileDimensionRequest): Promise<GraduateProfileDimension> => {
  try {
    const response = await apiClient.post('/graduate-profile-dimensions', data);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

export const updateGraduateProfileDimension = async (id: string, data: UpdateGraduateProfileDimensionRequest): Promise<GraduateProfileDimension> => {
  try {
    const response = await apiClient.put(`/graduate-profile-dimensions/${id}`, data);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

export const deleteGraduateProfileDimension = async (id: string): Promise<void> => {
  try {
    await apiClient.delete(`/graduate-profile-dimensions/${id}`);
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * CP Alignment API endpoints
 */
export const getCPAlignments = async (params?: {
  subject_category_id?: string;
  graduate_profile_dimension_id?: string;
  limit?: number;
  offset?: number;
}): Promise<CPAlignment[]> => {
  try {
    const response = await apiClient.get('/cp-alignments', { params });
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

export const getCPAlignmentById = async (id: string): Promise<CPAlignment> => {
  try {
    const response = await apiClient.get(`/cp-alignments/${id}`);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

export const createCPAlignment = async (data: CreateCPAlignmentRequest): Promise<CPAlignment> => {
  try {
    const response = await apiClient.post('/cp-alignments', data);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

export const updateCPAlignment = async (id: string, data: UpdateCPAlignmentRequest): Promise<CPAlignment> => {
  try {
    const response = await apiClient.put(`/cp-alignments/${id}`, data);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

export const deleteCPAlignment = async (id: string): Promise<void> => {
  try {
    await apiClient.delete(`/cp-alignments/${id}`);
  } catch (error) {
    throw handleApiError(error);
  }
};

/**
 * System Configuration API endpoints
 */
export const getSystemConfigurations = async (params?: {
  category?: string;
  key?: string;
  is_public?: boolean;
  school_id?: string;
  limit?: number;
  offset?: number;
}): Promise<SystemConfiguration[]> => {
  try {
    const response = await apiClient.get('/system-configurations', { params });
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

export const getSystemConfigurationByKey = async (key: string): Promise<SystemConfiguration> => {
  try {
    const response = await apiClient.get(`/system-configurations/key/${key}`);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

export const updateSystemConfiguration = async (key: string, data: UpdateSystemConfigurationRequest): Promise<SystemConfiguration> => {
  try {
    const response = await apiClient.put(`/system-configurations/key/${key}`, data);
    return response.data.data || response.data;
  } catch (error) {
    throw handleApiError(error);
  }
};

export default {
  // Academic Year
  getAcademicYears,
  getAcademicYearById,
  createAcademicYear,
  updateAcademicYear,
  activateAcademicYear,
  archiveAcademicYear,
  // Semester
  getSemesters,
  getSemesterById,
  createSemester,
  updateSemester,
  deleteSemester,
  // Subject Category
  getSubjectCategories,
  getSubjectCategoryById,
  createSubjectCategory,
  updateSubjectCategory,
  deleteSubjectCategory,
  // Graduate Profile Dimension
  getGraduateProfileDimensions,
  getGraduateProfileDimensionById,
  createGraduateProfileDimension,
  updateGraduateProfileDimension,
  deleteGraduateProfileDimension,
  // CP Alignment
  getCPAlignments,
  getCPAlignmentById,
  createCPAlignment,
  updateCPAlignment,
  deleteCPAlignment,
  // System Configuration
  getSystemConfigurations,
  getSystemConfigurationByKey,
  updateSystemConfiguration,
};
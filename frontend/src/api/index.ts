/**
 * API Client Exports
 */

// API Client
export { apiClient, handleApiError, ApiError } from './client';
export { default as apiClientDefault } from './client';

// Auth API
export { login, refreshToken, logout, me } from './auth';
export { default as authApiDefault } from './auth';

// User API
export {
  listUsers,
  getUser,
  createUser,
  updateUser,
  updateUserStatus,
  deleteUser,
} from './users';
export type {
  User,
  CreateUserRequest,
  UpdateUserRequest,
  ListUsersResponse,
  ListUsersParams,
} from './users';
export { default as usersApiDefault } from './users';

// School API
export {
  listSchools,
  getSchool,
  createSchool,
  updateSchool,
  updateSchoolStatus,
  deleteSchool,
} from './schools';
export type {
  School,
  CreateSchoolRequest,
  UpdateSchoolRequest,
  ListSchoolsResponse,
  ListSchoolsParams,
} from './schools';
export { default as schoolsApiDefault } from './schools';

// Role API
export {
  listRoles,
  getRole,
  createRole,
  updateRole,
  deleteRole,
} from './roles';
export type {
  Role,
  CreateRoleRequest,
  UpdateRoleRequest,
  ListRolesResponse,
  ListRolesParams,
} from './roles';
export { default as rolesApiDefault } from './roles';

// Permission API
export {
  listPermissions,
  getPermission,
  createPermission,
  updatePermission,
  deletePermission,
} from './permissions';
export type {
  Permission,
  CreatePermissionRequest,
  UpdatePermissionRequest,
  ListPermissionsResponse,
  ListPermissionsParams,
} from './permissions';
export { default as permissionsApiDefault } from './permissions';

// Narrative Report API
export {
  getNarrativeReports,
  getNarrativeReportById,
  createNarrativeReport,
  updateNarrativeReport,
  deleteNarrativeReport,
  publishNarrativeReport,
  getAchievementSummary,
} from './narrative-report';
export type {
  NarrativeReport,
  CreateNarrativeReportRequest,
  UpdateNarrativeReportRequest,
  NarrativeReportResponse,
} from './narrative-report';
export { default as narrativeReportApiDefault } from './narrative-report';

// CP (Capaian Pembelajaran) API
export {
  getSubjects,
  getSubjectById,
  getPhases,
  getPhaseById,
  getCPs,
  getCPById,
  getElementsByPhase,
  getElementById,
  getSubelementsByElement,
  getSubelementById,
} from './cp';
export type {
  CP,
  Subject,
  Phase,
  Element,
  Subelement,
} from './cp';
export { default as cpApiDefault } from './cp';

// ATP (Alur Tujuan Pembelajaran) API
export {
  getATPs,
  getATPById,
  getATPsBySet,
  createATP,
  updateATP,
  deleteATP,
  getATPSets,
  getATPSetById,
  createATPSet,
  updateATPSet,
  approveATPSet,
} from './atp';
export type {
  ATP,
  ATPSet,
  CreateATPRequest,
  UpdateATPRequest,
  CreateATPSetRequest,
  UpdateATPSetRequest,
} from './atp';
export { default as atpApiDefault } from './atp';

// Modul Ajar API
export {
  getModulAjars,
  getModulAjarById,
  getModulAjarsBySet,
  createModulAjar,
  updateModulAjar,
  deleteModulAjar,
  getModulAjarSets,
  getModulAjarSetById,
  createModulAjarSet,
  updateModulAjarSet,
} from './modul-ajar';
export type {
  ModulAjar,
  ModulAjarSet,
  CreateModulAjarRequest,
  UpdateModulAjarRequest,
  CreateModulAjarSetRequest,
  UpdateModulAjarSetRequest,
} from './modul-ajar';
export { default as modulAjarApiDefault } from './modul-ajar';
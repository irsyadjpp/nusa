/**
 * Authentication Types
 */

export interface User {
  id: string;
  email: string;
  name: string;
  role_id: string;
  role_name: string;
  school_id?: string;
  school_name?: string;
  is_active: boolean;
  status: string;
  created_at: string;
  updated_at: string;
}

export interface LoginCredentials {
  email: string;
  password: string;
}

export interface RegisterCredentials {
  name: string;
  email: string;
  company: string;
  password: string;
}

export interface AuthResponse {
  access_token: string;
  refresh_token: string;
  token_type: string;
  expires_in: number;
  user: User;
}

export interface AuthState {
  isAuthenticated: boolean;
  user: User | null;
  permissions: string[];
  loading: boolean;
  error: string | null;
}

export interface AuthContextValue extends AuthState {
  login: (credentials: LoginCredentials) => Promise<void>;
  logout: () => Promise<void>;
  refresh: () => Promise<void>;
  me: () => Promise<User>;
  hasPermission: (permission: string) => boolean;
  hasRole: (role: string) => boolean;
}

export type Role = 'SYSTEM_ADMIN' | 'SCHOOL_ADMIN' | 'TEACHER';
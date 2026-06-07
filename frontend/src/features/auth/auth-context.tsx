/**
 * Authentication Context
 * Provides authentication state and methods to the application
 */

import React, { createContext, useContext, useEffect, useState } from 'react';
import { AuthContextValue, AuthState, User, LoginCredentials } from './types';
import { AuthStorage } from './auth-storage';
import { login as loginApi, refreshToken as refreshApi, logout as logoutApi, me as meApi } from '@/api/auth';

const AuthContext = createContext<AuthContextValue | undefined>(undefined);

const initialState: AuthState = {
  isAuthenticated: false,
  user: null,
  permissions: [],
  loading: false,
  error: null,
};

interface AuthProviderProps {
  children: React.ReactNode;
}

export const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
  const [authState, setAuthState] = useState<AuthState>(initialState);

  // Restore session on mount
  useEffect(() => {
    const token = AuthStorage.getAccessToken();
    const user = AuthStorage.getUser();

    if (token && user) {
      setAuthState({
        isAuthenticated: true,
        user,
        permissions: [], // Will be loaded from /me endpoint
        loading: false,
        error: null,
      });

      // Load full user data including permissions
      loadCurrentUser();
    }
  }, []);

  const loadCurrentUser = async () => {
    try {
      setAuthState(prev => ({ ...prev, loading: true, error: null }));

      const data = await meApi();

      setAuthState(prev => ({
        ...prev,
        isAuthenticated: true,
        user: data.user,
        permissions: data.permissions,
        loading: false,
        error: null,
      }));

      // Update stored user
      AuthStorage.setUser(data.user);
    } catch (error: any) {
      console.error('Failed to load current user:', error);
      // If we can't load user, clear auth state
      AuthStorage.clear();
      setAuthState(initialState);
    }
  };

  const login = async (credentials: LoginCredentials): Promise<void> => {
    setAuthState(prev => ({ ...prev, loading: true, error: null }));

    try {
      const authData = await loginApi(credentials);

      // Store tokens
      AuthStorage.setAccessToken(authData.access_token);
      AuthStorage.setRefreshToken(authData.refresh_token);
      AuthStorage.setUser(authData.user);

      setAuthState({
        isAuthenticated: true,
        user: authData.user,
        permissions: [], // Will be loaded from /me endpoint
        loading: false,
        error: null,
      });

      // Load full user data including permissions
      await loadCurrentUser();
    } catch (error: any) {
      const errorMessage = error.message || 'Login failed';
      setAuthState(prev => ({
        ...prev,
        loading: false,
        error: errorMessage,
      }));
      throw error;
    }
  };

  const logout = async (): Promise<void> => {
    setAuthState(prev => ({ ...prev, loading: true }));

    try {
      const refreshToken = AuthStorage.getRefreshToken();
      if (refreshToken) {
        await logoutApi(refreshToken);
      }
    } catch (error) {
      // Don't fail logout if API call fails
      console.error('Logout API error:', error);
    } finally {
      // Always clear local storage and state
      AuthStorage.clear();
      setAuthState(initialState);
    }
  };

  const refresh = async (): Promise<void> => {
    setAuthState(prev => ({ ...prev, loading: true }));

    try {
      const refreshToken = AuthStorage.getRefreshToken();
      if (!refreshToken) {
        throw new Error('No refresh token available');
      }

      const tokens = await refreshApi(refreshToken);

      AuthStorage.setAccessToken(tokens.access_token);
      AuthStorage.setRefreshToken(tokens.refresh_token);

      setAuthState(prev => ({
        ...prev,
        loading: false,
      }));
    } catch (error: any) {
      // Refresh failed - clear auth state
      AuthStorage.clear();
      setAuthState(initialState);
      throw error;
    }
  };

  const me = async (): Promise<User> => {
    setAuthState(prev => ({ ...prev, loading: true }));

    try {
      const data = await meApi();

      setAuthState(prev => ({
        ...prev,
        user: data.user,
        permissions: data.permissions,
        loading: false,
      }));

      return data.user;
    } catch (error: any) {
      setAuthState(prev => ({
        ...prev,
        loading: false,
        error: error.message || 'Failed to fetch user',
      }));
      throw error;
    }
  };

  const hasPermission = (permission: string): boolean => {
    return authState.permissions.includes(permission);
  };

  const hasRole = (role: string): boolean => {
    return authState.user?.role_name === role;
  };

  const contextValue: AuthContextValue = {
    ...authState,
    login,
    logout,
    refresh,
    me,
    hasPermission,
    hasRole,
  };

  return (
    <AuthContext.Provider value={contextValue}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuthContext = (): AuthContextValue => {
  const context = useContext(AuthContext);

  if (context === undefined) {
    throw new Error('useAuthContext must be used within an AuthProvider');
  }

  return context;
};

export default AuthContext;
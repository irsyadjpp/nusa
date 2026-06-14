/**
 * Authentication Storage
 * Handles JWT token storage in localStorage
 */

const STORAGE_KEYS = {
  ACCESS_TOKEN: 'nusa_access_token',
  REFRESH_TOKEN: 'nusa_refresh_token',
  USER: 'nusa_user',
};

export interface StoredUser {
  id: string;
  email: string;
  name: string;
  role_id: string;
  role_name: string;
  school_name?: string;
  school_id?: string;
  is_active: boolean;
  status: string;
  created_at: string;
  updated_at: string;
}

export class AuthStorage {
  /**
   * Store access token
   */
  static setAccessToken(token: string): void {
    if (typeof window !== 'undefined') {
      localStorage.setItem(STORAGE_KEYS.ACCESS_TOKEN, token);
    }
  }

  /**
   * Get access token
   */
  static getAccessToken(): string | null {
    if (typeof window !== 'undefined') {
      return localStorage.getItem(STORAGE_KEYS.ACCESS_TOKEN);
    }
    return null;
  }

  /**
   * Remove access token
   */
  static removeAccessToken(): void {
    if (typeof window !== 'undefined') {
      localStorage.removeItem(STORAGE_KEYS.ACCESS_TOKEN);
    }
  }

  /**
   * Store refresh token
   */
  static setRefreshToken(token: string): void {
    if (typeof window !== 'undefined') {
      localStorage.setItem(STORAGE_KEYS.REFRESH_TOKEN, token);
    }
  }

  /**
   * Get refresh token
   */
  static getRefreshToken(): string | null {
    if (typeof window !== 'undefined') {
      return localStorage.getItem(STORAGE_KEYS.REFRESH_TOKEN);
    }
    return null;
  }

  /**
   * Remove refresh token
   */
  static removeRefreshToken(): void {
    if (typeof window !== 'undefined') {
      localStorage.removeItem(STORAGE_KEYS.REFRESH_TOKEN);
    }
  }

  /**
   * Store user data
   */
  static setUser(user: StoredUser): void {
    if (typeof window !== 'undefined') {
      localStorage.setItem(STORAGE_KEYS.USER, JSON.stringify(user));
    }
  }

  /**
   * Get user data
   */
  static getUser(): StoredUser | null {
    if (typeof window !== 'undefined') {
      const userJson = localStorage.getItem(STORAGE_KEYS.USER);
      if (userJson) {
        try {
          return JSON.parse(userJson);
        } catch (error) {
          console.error('Failed to parse user data:', error);
          return null;
        }
      }
    }
    return null;
  }

  /**
   * Remove user data
   */
  static removeUser(): void {
    if (typeof window !== 'undefined') {
      localStorage.removeItem(STORAGE_KEYS.USER);
    }
  }

  /**
   * Clear all auth data
   */
  static clear(): void {
    this.removeAccessToken();
    this.removeRefreshToken();
    this.removeUser();
  }

  /**
   * Check if user is authenticated
   */
  static isAuthenticated(): boolean {
    return this.getAccessToken() !== null;
  }
}
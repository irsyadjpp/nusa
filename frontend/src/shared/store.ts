/**
 * Zustand Store Configuration
 * Global state management for UI state that doesn't come from the server
 */

import { create } from 'zustand';
import { persist } from 'zustand/middleware';

// Types
interface UIState {
  // Sidebar state
  sidebarOpen: boolean;
  setSidebarOpen: (open: boolean) => void;
  toggleSidebar: () => void;

  // Theme state
  theme: 'light' | 'dark';
  setTheme: (theme: 'light' | 'dark') => void;

  // Filter state (persisted across pages)
  filters: {
    subject?: string;
    phase?: string;
    status?: string;
  };
  setFilters: (filters: Partial<UIState['filters']>) => void;
  clearFilters: () => void;

  // User session state
  userSession: {
    userId?: string;
    userName?: string;
    userRole?: string;
  };
  setUserSession: (session: Partial<UIState['userSession']>) => void;
  clearUserSession: () => void;
}

// Create the store with persistence
export const useUIStore = create<UIState>()(
  persist(
    (set) => ({
      // Sidebar state
      sidebarOpen: true,
      setSidebarOpen: (open) => set({ sidebarOpen: open }),
      toggleSidebar: () => set((state) => ({ sidebarOpen: !state.sidebarOpen })),

      // Theme state
      theme: 'light',
      setTheme: (theme) => set({ theme }),

      // Filter state
      filters: {},
      setFilters: (filters) => set((state) => ({ filters: { ...state.filters, ...filters } })),
      clearFilters: () => set({ filters: {} }),

      // User session state
      userSession: {},
      setUserSession: (session) => set((state) => ({ userSession: { ...state.userSession, ...session } })),
      clearUserSession: () => set({ userSession: {} }),
    }),
    {
      name: 'nusa-ui-storage', // Local storage key
      partialize: (state) => ({
        // Only persist these fields
        sidebarOpen: state.sidebarOpen,
        theme: state.theme,
        filters: state.filters,
      }),
    }
  )
);

// Selectors for common use cases
export const selectSidebarOpen = (state: UIState) => state.sidebarOpen;
export const selectTheme = (state: UIState) => state.theme;
export const selectFilters = (state: UIState) => state.filters;
export const selectUserSession = (state: UIState) => state.userSession;

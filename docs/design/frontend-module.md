# 32_FRONTEND_MODULE_DESIGN.md

## Foundation Document for NUSA Education Platform

**Version**: 1.0
**Date**: June 2026
**Status**: FOUNDATION DOCUMENT
**Alignment**: Aligned with 27_BACKEND_FOUNDATION_DESIGN.md, 28_AUTHENTICATION_DESIGN.md, 29_CURRICULUM_MODULE_DESIGN.md, 30_TP_GENERATION_MODULE_DESIGN.md, 31_WORKFLOW_ENGINE_DESIGN.md

**Purpose**: Map the frontend implementation for NUSA. This document defines route structure, layout structure, feature module structure, API client strategy, state management strategy, authentication flow, TP generation screens, and workflow screens for MVP.

---

# SECTION 1 — Executive Summary

## Why Frontend Module Design Matters

A well-designed frontend ensures:
- Consistent user experience across the application
- Scalable and maintainable codebase
- Type-safe development with TypeScript
- Efficient state management
- Clear separation of concerns
- Reusable components

## Technology Stack

| Component | Technology | Purpose |
|-----------|------------|---------|
| Framework | React 18 | UI library |
| Language | TypeScript 5 | Type-safe JavaScript |
| Build Tool | Vite 5 | Fast build tool |
| UI Library | Material UI (MUI) v6 | Component library |
| Routing | React Router v6 | Client-side routing |
| State Management | React Context + Hooks | Global state |
| HTTP Client | Axios | API requests |
| Form Handling | React Hook Form | Form validation |
| Date Handling | date-fns | Date utilities |

## Core Principles

- **Component-Based**: Reusable, composable components
- **Type-Safe**: Full TypeScript coverage
- **Performance-Optimized**: Code splitting, lazy loading
- **Accessible**: WCAG 2.1 AA compliance
- **Responsive**: Mobile-first design
- **MVP-Focused**: Essential features only

---

# SECTION 2 — Route Structure

## Route Hierarchy

```
/ (public)
├── /login
├── /register (future)
└── /forgot-password (future)

/dashboard (protected)
├── /curriculum
│   ├── /subjects
│   ├── /subjects/:id
│   └── /cp/:id
├── /tp-generation
│   ├── /cp/:cp-id/generate
│   ├── /tp-sets
│   └── /tp-sets/:id
├── /workflow
│   ├── /pending
│   └── /history
├── /profile
└── /settings (future)
```

## Route Configuration

```typescript
// src/routes/index.tsx
import { createBrowserRouter } from 'react-router-dom';
import { ProtectedRoute } from '../components/ProtectedRoute';
import { PublicRoute } from '../components/PublicRoute';
import Layout from '../layouts/Layout';
import LoginPage from '../pages/auth/LoginPage';
import DashboardPage from '../pages/dashboard/DashboardPage';
import CurriculumPage from '../pages/curriculum/CurriculumPage';
import SubjectPage from '../pages/curriculum/SubjectPage';
import CPPage from '../pages/curriculum/CPPage';
import TPGenerationPage from '../pages/tp-generation/TPGenerationPage';
import TPSetPage from '../pages/tp-generation/TPSetPage';
import TPSetDetailPage from '../pages/tp-generation/TPSetDetailPage';
import WorkflowPendingPage from '../pages/workflow/WorkflowPendingPage';
import WorkflowHistoryPage from '../pages/workflow/WorkflowHistoryPage';
import ProfilePage from '../pages/profile/ProfilePage';

const router = createBrowserRouter([
  {
    path: '/login',
    element: (
      <PublicRoute>
        <LoginPage />
      </PublicRoute>
    ),
  },
  {
    path: '/',
    element: (
      <ProtectedRoute>
        <Layout />
      </ProtectedRoute>
    ),
    children: [
      {
        index: true,
        element: <DashboardPage />,
      },
      {
        path: 'curriculum',
        element: <CurriculumPage />,
      },
      {
        path: 'curriculum/subjects/:id',
        element: <SubjectPage />,
      },
      {
        path: 'curriculum/cp/:id',
        element: <CPPage />,
      },
      {
        path: 'tp-generation',
        element: <TPGenerationPage />,
      },
      {
        path: 'tp-generation/cp/:cpId/generate',
        element: <TPSetDetailPage mode="generate" />,
      },
      {
        path: 'tp-generation/tp-sets',
        element: <TPSetPage />,
      },
      {
        path: 'tp-generation/tp-sets/:id',
        element: <TPSetDetailPage mode="view" />,
      },
      {
        path: 'workflow/pending',
        element: <WorkflowPendingPage />,
      },
      {
        path: 'workflow/history',
        element: <WorkflowHistoryPage />,
      },
      {
        path: 'profile',
        element: <ProfilePage />,
      },
    ],
  },
]);

export default router;
```

## Route Guards

### Protected Route

```typescript
// src/components/ProtectedRoute.tsx
import { Navigate, useLocation } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';

interface ProtectedRouteProps {
  children: React.ReactNode;
}

export function ProtectedRoute({ children }: ProtectedRouteProps) {
  const { isAuthenticated, isLoading } = useAuth();
  const location = useLocation();

  if (isLoading) {
    return <CircularProgress />;
  }

  if (!isAuthenticated) {
    return <Navigate to="/login" state={{ from: location }} replace />;
  }

  return <>{children}</>;
}
```

### Public Route

```typescript
// src/components/PublicRoute.tsx
import { Navigate } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';

interface PublicRouteProps {
  children: React.ReactNode;
}

export function PublicRoute({ children }: PublicRouteProps) {
  const { isAuthenticated, isLoading } = useAuth();

  if (isLoading) {
    return <CircularProgress />;
  }

  if (isAuthenticated) {
    return <Navigate to="/dashboard" replace />;
  }

  return <>{children}</>;
}
```

---

# SECTION 3 — Layout Structure

## Layout Components

### Main Layout

```
Layout
├── AppBar
│   ├── Logo
│   ├── Navigation
│   └── UserMenu
├── Sidebar
│   ├── Navigation Items
│   └── Collapse Toggle
└── Main Content
    └── Page Content
```

### Layout Implementation

```typescript
// src/layouts/Layout.tsx
import { Box, AppBar, Toolbar, Typography, Drawer, List, ListItem, ListItemButton, ListItemText, IconButton } from '@mui/material';
import { Menu as MenuIcon, ChevronLeft } from '@mui/icons-material';
import { Outlet, useNavigate, useLocation } from 'react-router-dom';
import { useState } from 'react';
import { useAuth } from '../contexts/AuthContext';

const DRAWER_WIDTH = 240;

export default function Layout() {
  const [mobileOpen, setMobileOpen] = useState(false);
  const navigate = useNavigate();
  const location = useLocation();
  const { user, logout } = useAuth();

  const handleDrawerToggle = () => {
    setMobileOpen(!mobileOpen);
  };

  const menuItems = [
    { text: 'Dashboard', path: '/' },
    { text: 'Curriculum', path: '/curriculum' },
    { text: 'TP Generation', path: '/tp-generation' },
    { text: 'Workflow', path: '/workflow/pending' },
    { text: 'Profile', path: '/profile' },
  ];

  const drawer = (
    <div>
      <Toolbar />
      <List>
        {menuItems.map((item) => (
          <ListItem key={item.text} disablePadding>
            <ListItemButton
              selected={location.pathname === item.path}
              onClick={() => navigate(item.path)}
            >
              <ListItemText primary={item.text} />
            </ListItemButton>
          </ListItem>
        ))}
      </List>
    </div>
  );

  return (
    <Box sx={{ display: 'flex' }}>
      <AppBar
        position="fixed"
        sx={{
          width: { sm: `calc(100% - ${DRAWER_WIDTH}px)` },
          ml: { sm: `${DRAWER_WIDTH}px` },
        }}
      >
        <Toolbar>
          <IconButton
            color="inherit"
            edge="start"
            onClick={handleDrawerToggle}
            sx={{ mr: 2, display: { sm: 'none' } }}
          >
            <MenuIcon />
          </IconButton>
          <Typography variant="h6" noWrap component="div" sx={{ flexGrow: 1 }}>
            NUSA Platform
          </Typography>
          <Typography variant="body1" sx={{ mr: 2 }}>
            {user?.name}
          </Typography>
          <IconButton color="inherit" onClick={logout}>
            Logout
          </IconButton>
        </Toolbar>
      </AppBar>
      <Box
        component="nav"
        sx={{ width: { sm: DRAWER_WIDTH }, flexShrink: { sm: 0 } }}
      >
        <Drawer
          variant="temporary"
          open={mobileOpen}
          onClose={handleDrawerToggle}
          ModalProps={{ keepMounted: true }}
          sx={{
            display: { xs: 'block', sm: 'none' },
            '& .MuiDrawer-paper': { boxSizing: 'border-box', width: DRAWER_WIDTH },
          }}
        >
          {drawer}
        </Drawer>
        <Drawer
          variant="permanent"
          sx={{
            display: { xs: 'none', sm: 'block' },
            '& .MuiDrawer-paper': { boxSizing: 'border-box', width: DRAWER_WIDTH },
          }}
          open
        >
          {drawer}
        </Drawer>
      </Box>
      <Box
        component="main"
        sx={{
          flexGrow: 1,
          p: 3,
          width: { sm: `calc(100% - ${DRAWER_WIDTH}px)` },
        }}
      >
        <Toolbar />
        <Outlet />
      </Box>
    </Box>
  );
}
```

---

# SECTION 4 — Feature Module Structure

## Directory Structure

```
src/
├── assets/
│   ├── images/
│   └── fonts/
├── components/
│   ├── common/
│   │   ├── Button.tsx
│   │   ├── Card.tsx
│   │   ├── Dialog.tsx
│   │   └── Loading.tsx
│   ├── curriculum/
│   │   ├── SubjectCard.tsx
│   │   ├── CPList.tsx
│   │   └── CPTree.tsx
│   ├── tp-generation/
│   │   ├── TPSetCard.tsx
│   │   ├── TPItemCard.tsx
│   │   ├── GenerationForm.tsx
│   │   └── TPSetDetail.tsx
│   ├── workflow/
│   │   ├── StatusBadge.tsx
│   │   ├── HistoryTimeline.tsx
│   │   └── ActionButtons.tsx
│   └── ProtectedRoute.tsx
├── contexts/
│   ├── AuthContext.tsx
│   └── NotificationContext.tsx
├── hooks/
│   ├── useAuth.ts
│   ├── useAPI.ts
│   └── useWorkflow.ts
├── layouts/
│   └── Layout.tsx
├── lib/
│   ├── api/
│   │   ├── client.ts
│   │   ├── auth.ts
│   │   ├── curriculum.ts
│   │   ├── tp-generation.ts
│   │   └── workflow.ts
│   └── utils/
│       ├── validation.ts
│       └── formatting.ts
├── pages/
│   ├── auth/
│   │   └── LoginPage.tsx
│   ├── dashboard/
│   │   └── DashboardPage.tsx
│   ├── curriculum/
│   │   ├── CurriculumPage.tsx
│   │   ├── SubjectPage.tsx
│   │   └── CPPage.tsx
│   ├── tp-generation/
│   │   ├── TPGenerationPage.tsx
│   │   ├── TPSetPage.tsx
│   │   └── TPSetDetailPage.tsx
│   ├── workflow/
│   │   ├── WorkflowPendingPage.tsx
│   │   └── WorkflowHistoryPage.tsx
│   └── profile/
│       └── ProfilePage.tsx
├── routes/
│   └── index.tsx
├── types/
│   ├── auth.ts
│   ├── curriculum.ts
│   ├── tp-generation.ts
│   └── workflow.ts
├── App.tsx
└── main.tsx
```

## Component Organization

### Common Components

Reusable components used across the application:

- **Button**: Styled button with loading state
- **Card**: Card container with consistent styling
- **Dialog**: Modal dialog component
- **Loading**: Loading spinner component
- **ErrorBoundary**: Error boundary for error handling

### Feature Components

Components specific to feature modules:

- **Curriculum**: SubjectCard, CPList, CPTree
- **TP Generation**: TPSetCard, TPItemCard, GenerationForm, TPSetDetail
- **Workflow**: StatusBadge, HistoryTimeline, ActionButtons

---

# SECTION 5 — API Client Strategy

## API Client Configuration

### Axios Instance

```typescript
// src/lib/api/client.ts
import axios, { AxiosInstance, AxiosError } from 'axios';
import { useAuthStore } from '../contexts/AuthContext';

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080';

export const apiClient: AxiosInstance = axios.create({
  baseURL: API_BASE_URL,
  timeout: 30000,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Request interceptor to add auth token
apiClient.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('access_token');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// Response interceptor to handle errors
apiClient.interceptors.response.use(
  (response) => response,
  async (error: AxiosError) => {
    if (error.response?.status === 401) {
      // Token expired, try to refresh
      const refreshToken = localStorage.getItem('refresh_token');
      if (refreshToken) {
        try {
          const response = await axios.post(`${API_BASE_URL}/api/v1/public/refresh`, {
            refresh_token: refreshToken,
          });
          
          const { access_token, refresh_token: newRefreshToken } = response.data;
          localStorage.setItem('access_token', access_token);
          localStorage.setItem('refresh_token', newRefreshToken);
          
          // Retry original request
          if (error.config) {
            error.config.headers.Authorization = `Bearer ${access_token}`;
            return apiClient.request(error.config);
          }
        } catch (refreshError) {
          // Refresh failed, logout user
          localStorage.removeItem('access_token');
          localStorage.removeItem('refresh_token');
          window.location.href = '/login';
        }
      } else {
        window.location.href = '/login';
      }
    }
    return Promise.reject(error);
  }
);
```

## API Service Modules

### Auth API

```typescript
// src/lib/api/auth.ts
import { apiClient } from './client';
import { LoginRequest, LoginResponse, RefreshTokenRequest, RefreshTokenResponse } from '../../types/auth';

export const authAPI = {
  login: async (data: LoginRequest): Promise<LoginResponse> => {
    const response = await apiClient.post<LoginResponse>('/api/v1/public/login', data);
    return response.data;
  },

  refreshToken: async (data: RefreshTokenRequest): Promise<RefreshTokenResponse> => {
    const response = await apiClient.post<RefreshTokenResponse>('/api/v1/public/refresh', data);
    return response.data;
  },

  logout: async (): Promise<void> => {
    await apiClient.post('/api/v1/auth/logout');
  },

  me: async (): Promise<any> => {
    const response = await apiClient.get('/api/v1/auth/me');
    return response.data;
  },
};
```

### Curriculum API

```typescript
// src/lib/api/curriculum.ts
import { apiClient } from './client';
import { Subject, CP, CPListResponse, SubjectTree } from '../../types/curriculum';

export const curriculumAPI = {
  getSubjects: async (): Promise<Subject[]> => {
    const response = await apiClient.get<Subject[]>('/api/v1/public/curriculum/subjects');
    return response.data;
  },

  getSubjectTree: async (subjectId: string): Promise<SubjectTree> => {
    const response = await apiClient.get<SubjectTree>(`/api/v1/public/curriculum/subjects/${subjectId}/tree`);
    return response.data;
  },

  getCP: async (cpId: string): Promise<CP> => {
    const response = await apiClient.get<CP>(`/api/v1/public/curriculum/cp/${cpId}`);
    return response.data;
  },

  listCPs: async (filters: any, pagination: any): Promise<CPListResponse> => {
    const response = await apiClient.get<CPListResponse>('/api/v1/public/curriculum/cp', {
      params: { ...filters, ...pagination },
    });
    return response.data;
  },

  searchCPs: async (query: string, filters: any): Promise<CPListResponse> => {
    const response = await apiClient.get<CPListResponse>('/api/v1/public/curriculum/cp/search', {
      params: { q: query, ...filters },
    });
    return response.data;
  },
};
```

### TP Generation API

```typescript
// src/lib/api/tp-generation.ts
import { apiClient } from './client';
import { TPSet, TPSetListResponse, GenerateTPSetRequest, GenerateTPSetResponse } from '../../types/tp-generation';

export const tpGenerationAPI = {
  generateTPSet: async (cpId: string, data: GenerateTPSetRequest): Promise<GenerateTPSetResponse> => {
    const response = await apiClient.post<GenerateTPSetResponse>(`/api/v1/curriculum/cp/${cpId}/tp-sets/generate`, data);
    return response.data;
  },

  regenerateTPSet: async (tpSetId: string, data: GenerateTPSetRequest): Promise<GenerateTPSetResponse> => {
    const response = await apiClient.post<GenerateTPSetResponse>(`/api/v1/curriculum/tp-sets/${tpSetId}/regenerate`, data);
    return response.data;
  },

  getTPSet: async (tpSetId: string): Promise<TPSet> => {
    const response = await apiClient.get<TPSet>(`/api/v1/curriculum/tp-sets/${tpSetId}`);
    return response.data;
  },

  listTPSets: async (cpId: string, pagination: any): Promise<TPSetListResponse> => {
    const response = await apiClient.get<TPSetListResponse>(`/api/v1/curriculum/cp/${cpId}/tp-sets`, {
      params: pagination,
    });
    return response.data;
  },

  submitTPSet: async (tpSetId: string, reason: string): Promise<void> => {
    await apiClient.post(`/api/v1/curriculum/tp-sets/${tpSetId}/submit`, { reason });
  },

  approveTPSet: async (tpSetId: string, reason: string): Promise<void> => {
    await apiClient.post(`/api/v1/curriculum/tp-sets/${tpSetId}/approve`, { reason });
  },

  rejectTPSet: async (tpSetId: string, reason: string): Promise<void> => {
    await apiClient.post(`/api/v1/curriculum/tp-sets/${tpSetId}/reject`, { reason });
  },

  archiveTPSet: async (tpSetId: string): Promise<void> => {
    await apiClient.post(`/api/v1/curriculum/tp-sets/${tpSetId}/archive`);
  },
};
```

### Workflow API

```typescript
// src/lib/api/workflow.ts
import { apiClient } from './client';
import { WorkflowHistory, PendingApproval } from '../../types/workflow';

export const workflowAPI = {
  getHistory: async (artifactId: string, artifactType: string): Promise<WorkflowHistory[]> => {
    const response = await apiClient.get<WorkflowHistory[]>(`/api/v1/workflow/${artifactType}/${artifactId}/history`);
    return response.data;
  },

  getPendingApprovals: async (artifactType: string): Promise<PendingApproval[]> => {
    const response = await apiClient.get<PendingApproval[]>(`/api/v1/workflow/${artifactType}/pending`);
    return response.data;
  },

  submit: async (artifactId: string, artifactType: string, reason: string): Promise<void> => {
    await apiClient.post(`/api/v1/workflow/${artifactType}/${artifactId}/submit`, { reason });
  },

  approve: async (artifactId: string, artifactType: string, reason: string): Promise<void> => {
    await apiClient.post(`/api/v1/workflow/${artifactType}/${artifactId}/approve`, { reason });
  },

  reject: async (artifactId: string, artifactType: string, reason: string): Promise<void> => {
    await apiClient.post(`/api/v1/workflow/${artifactType}/${artifactId}/reject`, { reason });
  },

  archive: async (artifactId: string, artifactType: string): Promise<void> => {
    await apiClient.post(`/api/v1/workflow/${artifactType}/${artifactId}/archive`);
  },
};
```

---

# SECTION 6 — State Management Strategy

## Context-Based State Management

### Auth Context

```typescript
// src/contexts/AuthContext.tsx
import { createContext, useContext, useState, useEffect, ReactNode } from 'react';
import { authAPI } from '../lib/api/auth';
import { User } from '../types/auth';

interface AuthContextType {
  user: User | null;
  isAuthenticated: boolean;
  isLoading: boolean;
  login: (email: string, password: string) => Promise<void>;
  logout: () => Promise<void>;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function AuthProvider({ children }: { children: ReactNode }) {
  const [user, setUser] = useState<User | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    checkAuth();
  }, []);

  const checkAuth = async () => {
    const token = localStorage.getItem('access_token');
    if (token) {
      try {
        const userData = await authAPI.me();
        setUser(userData);
      } catch (error) {
        localStorage.removeItem('access_token');
        localStorage.removeItem('refresh_token');
      }
    }
    setIsLoading(false);
  };

  const login = async (email: string, password: string) => {
    const response = await authAPI.login({ email, password });
    localStorage.setItem('access_token', response.access_token);
    localStorage.setItem('refresh_token', response.refresh_token);
    setUser(response.user);
  };

  const logout = async () => {
    try {
      await authAPI.logout();
    } catch (error) {
      // Ignore logout errors
    }
    localStorage.removeItem('access_token');
    localStorage.removeItem('refresh_token');
    setUser(null);
  };

  return (
    <AuthContext.Provider
      value={{
        user,
        isAuthenticated: !!user,
        isLoading,
        login,
        logout,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
}

export function useAuth() {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
}
```

### Notification Context

```typescript
// src/contexts/NotificationContext.tsx
import { createContext, useContext, useState, ReactNode } from 'react';
import { Snackbar, Alert } from '@mui/material';

interface Notification {
  message: string;
  severity: 'success' | 'error' | 'warning' | 'info';
}

interface NotificationContextType {
  showNotification: (message: string, severity: Notification['severity']) => void;
}

const NotificationContext = createContext<NotificationContextType | undefined>(undefined);

export function NotificationProvider({ children }: { children: ReactNode }) {
  const [notification, setNotification] = useState<Notification | null>(null);
  const [open, setOpen] = useState(false);

  const showNotification = (message: string, severity: Notification['severity']) => {
    setNotification({ message, severity });
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
  };

  return (
    <NotificationContext.Provider value={{ showNotification }}>
      {children}
      <Snackbar open={open} autoHideDuration={6000} onClose={handleClose}>
        <Alert onClose={handleClose} severity={notification?.severity} sx={{ width: '100%' }}>
          {notification?.message}
        </Alert>
      </Snackbar>
    </NotificationContext.Provider>
  );
}

export function useNotification() {
  const context = useContext(NotificationContext);
  if (context === undefined) {
    throw new Error('useNotification must be used within a NotificationProvider');
  }
  return context;
}
```

## Custom Hooks

### useAPI Hook

```typescript
// src/hooks/useAPI.ts
import { useState, useEffect } from 'react';

interface UseAPIResult<T> {
  data: T | null;
  loading: boolean;
  error: Error | null;
  refetch: () => Promise<void>;
}

export function useAPI<T>(
  apiFunction: () => Promise<T>,
  dependencies: any[] = []
): UseAPIResult<T> {
  const [data, setData] = useState<T | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<Error | null>(null);

  const fetchData = async () => {
    setLoading(true);
    setError(null);
    try {
      const result = await apiFunction();
      setData(result);
    } catch (err) {
      setError(err as Error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchData();
  }, dependencies);

  return { data, loading, error, refetch: fetchData };
}
```

### useWorkflow Hook

```typescript
// src/hooks/useWorkflow.ts
import { useState } from 'react';
import { workflowAPI } from '../lib/api/workflow';
import { useNotification } from '../contexts/NotificationContext';

export function useWorkflow() {
  const [loading, setLoading] = useState(false);
  const { showNotification } = useNotification();

  const submit = async (artifactId: string, artifactType: string, reason: string) => {
    setLoading(true);
    try {
      await workflowAPI.submit(artifactId, artifactType, reason);
      showNotification('Submitted for review', 'success');
    } catch (error) {
      showNotification('Failed to submit', 'error');
      throw error;
    } finally {
      setLoading(false);
    }
  };

  const approve = async (artifactId: string, artifactType: string, reason: string) => {
    setLoading(true);
    try {
      await workflowAPI.approve(artifactId, artifactType, reason);
      showNotification('Approved successfully', 'success');
    } catch (error) {
      showNotification('Failed to approve', 'error');
      throw error;
    } finally {
      setLoading(false);
    }
  };

  const reject = async (artifactId: string, artifactType: string, reason: string) => {
    setLoading(true);
    try {
      await workflowAPI.reject(artifactId, artifactType, reason);
      showNotification('Rejected', 'warning');
    } catch (error) {
      showNotification('Failed to reject', 'error');
      throw error;
    } finally {
      setLoading(false);
    }
  };

  const archive = async (artifactId: string, artifactType: string) => {
    setLoading(true);
    try {
      await workflowAPI.archive(artifactId, artifactType);
      showNotification('Archived successfully', 'success');
    } catch (error) {
      showNotification('Failed to archive', 'error');
      throw error;
    } finally {
      setLoading(false);
    }
  };

  return { submit, approve, reject, archive, loading };
}
```

---

# SECTION 7 — Authentication Flow

## Login Screen

```typescript
// src/pages/auth/LoginPage.tsx
import { useState } from 'react';
import { useNavigate, useLocation } from 'react-router-dom';
import { Box, Container, TextField, Button, Typography, Alert } from '@mui/material';
import { useAuth } from '../../contexts/AuthContext';
import { useNotification } from '../../contexts/NotificationContext';

export default function LoginPage() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(false);
  
  const { login } = useAuth();
  const { showNotification } = useNotification();
  const navigate = useNavigate();
  const location = useLocation();

  const from = (location.state as any)?.from?.pathname || '/';

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');
    setLoading(true);

    try {
      await login(email, password);
      showNotification('Login successful', 'success');
      navigate(from, { replace: true });
    } catch (err: any) {
      setError(err.response?.data?.error || 'Login failed');
      showNotification('Login failed', 'error');
    } finally {
      setLoading(false);
    }
  };

  return (
    <Container maxWidth="sm">
      <Box sx={{ mt: 8, display: 'flex', flexDirection: 'column', alignItems: 'center' }}>
        <Typography component="h1" variant="h4" sx={{ mb: 4 }}>
          NUSA Platform
        </Typography>
        
        {error && <Alert severity="error" sx={{ mb: 2, width: '100%' }}>{error}</Alert>}
        
        <Box component="form" onSubmit={handleSubmit} sx={{ width: '100%' }}>
          <TextField
            margin="normal"
            required
            fullWidth
            id="email"
            label="Email"
            name="email"
            autoComplete="email"
            autoFocus
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
          <TextField
            margin="normal"
            required
            fullWidth
            name="password"
            label="Password"
            type="password"
            id="password"
            autoComplete="current-password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
          <Button
            type="submit"
            fullWidth
            variant="contained"
            sx={{ mt: 3, mb: 2 }}
            disabled={loading}
          >
            {loading ? 'Signing in...' : 'Sign In'}
          </Button>
        </Box>
      </Box>
    </Container>
  );
}
```

## Logout Flow

```typescript
// In Layout component or user menu
const handleLogout = async () => {
  await logout();
  navigate('/login');
};
```

---

# SECTION 8 — TP Generation Screens

## TP Generation Page

```typescript
// src/pages/tp-generation/TPGenerationPage.tsx
import { useState } from 'react';
import { Box, Container, Typography, Grid, Card, CardContent, Button, TextField } from '@mui/material';
import { useNavigate } from 'react-router-dom';
import { curriculumAPI } from '../../lib/api/curriculum';
import { useAPI } from '../../hooks/useAPI';
import { Subject } from '../../types/curriculum';

export default function TPGenerationPage() {
  const navigate = useNavigate();
  const [searchQuery, setSearchQuery] = useState('');
  
  const { data: subjects, loading } = useAPI(() => curriculumAPI.getSubjects());

  const filteredSubjects = subjects?.filter(subject =>
    subject.name.toLowerCase().includes(searchQuery.toLowerCase()) ||
    subject.code.toLowerCase().includes(searchQuery.toLowerCase())
  );

  return (
    <Container maxWidth="lg">
      <Box sx={{ mb: 4 }}>
        <Typography variant="h4" component="h1">
          TP Generation
        </Typography>
        <Typography variant="body1" color="text.secondary">
          Generate Teaching Plans from Curriculum CPs
        </Typography>
      </Box>

      <TextField
        fullWidth
        placeholder="Search subjects..."
        value={searchQuery}
        onChange={(e) => setSearchQuery(e.target.value)}
        sx={{ mb: 3 }}
      />

      {loading ? (
        <Typography>Loading...</Typography>
      ) : (
        <Grid container spacing={3}>
          {filteredSubjects?.map((subject) => (
            <Grid item xs={12} sm={6} md={4} key={subject.id}>
              <Card>
                <CardContent>
                  <Typography variant="h6" gutterBottom>
                    {subject.code} - {subject.name}
                  </Typography>
                  <Button
                    variant="contained"
                    fullWidth
                    onClick={() => navigate(`/curriculum/subjects/${subject.id}`)}
                  >
                    View CPs
                  </Button>
                </CardContent>
              </Card>
            </Grid>
          ))}
        </Grid>
      )}
    </Container>
  );
}
```

## CP Page with TP Generation

```typescript
// src/pages/curriculum/CPPage.tsx
import { useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { Box, Container, Typography, Button, Card, CardContent, Dialog, DialogTitle, DialogContent, DialogActions, TextField } from '@mui/material';
import { curriculumAPI } from '../../lib/api/curriculum';
import { tpGenerationAPI } from '../../lib/api/tp-generation';
import { useAPI } from '../../hooks/useAPI';
import { useNotification } from '../../contexts/NotificationContext';
import { CP } from '../../types/curriculum';

export default function CPPage() {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const { showNotification } = useNotification();
  const [openGenerateDialog, setOpenGenerateDialog] = useState(false);
  const [preferences, setPreferences] = useState({ duration_weeks: 12, focus_areas: [] });
  const [generating, setGenerating] = useState(false);

  const { data: cp } = useAPI(() => curriculumAPI.getCP(id!), [id]);

  const handleGenerate = async () => {
    setGenerating(true);
    try {
      const result = await tpGenerationAPI.generateTPSet(id!, preferences);
      showNotification('TP Set generated successfully', 'success');
      setOpenGenerateDialog(false);
      navigate(`/tp-generation/tp-sets/${result.id}`);
    } catch (error) {
      showNotification('Failed to generate TP Set', 'error');
    } finally {
      setGenerating(false);
    }
  };

  return (
    <Container maxWidth="lg">
      <Box sx={{ mb: 4 }}>
        <Button onClick={() => navigate(-1)} sx={{ mb: 2 }}>
          Back
        </Button>
        <Typography variant="h4" component="h1">
          {cp?.code}
        </Typography>
        <Typography variant="body1" paragraph>
          {cp?.text}
        </Typography>
        <Button
          variant="contained"
          onClick={() => setOpenGenerateDialog(true)}
        >
          Generate TP Set
        </Button>
      </Box>

      <Dialog open={openGenerateDialog} onClose={() => setOpenGenerateDialog(false)} maxWidth="sm" fullWidth>
        <DialogTitle>Generate TP Set</DialogTitle>
        <DialogContent>
          <TextField
            fullWidth
            label="Duration (weeks)"
            type="number"
            value={preferences.duration_weeks}
            onChange={(e) => setPreferences({ ...preferences, duration_weeks: parseInt(e.target.value) })}
            sx={{ mt: 2 }}
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setOpenGenerateDialog(false)}>Cancel</Button>
          <Button onClick={handleGenerate} disabled={generating} variant="contained">
            {generating ? 'Generating...' : 'Generate'}
          </Button>
        </DialogActions>
      </Dialog>
    </Container>
  );
}
```

## TP Set Detail Page

```typescript
// src/pages/tp-generation/TPSetDetailPage.tsx
import { useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { Box, Container, Typography, Button, Card, CardContent, Grid, Chip, Dialog, DialogTitle, DialogContent, DialogActions, TextField, Timeline, TimelineItem, TimelineSeparator, TimelineConnector, TimelineContent, TimelineDot } from '@mui/material';
import { tpGenerationAPI } from '../../lib/api/tp-generation';
import { workflowAPI } from '../../lib/api/workflow';
import { useAPI } from '../../hooks/useAPI';
import { useWorkflow } from '../../hooks/useWorkflow';
import { useNotification } from '../../contexts/NotificationContext';
import { TPSet } from '../../types/tp-generation';

interface TPSetDetailPageProps {
  mode?: 'generate' | 'view';
}

export default function TPSetDetailPage({ mode = 'view' }: TPSetDetailPageProps) {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const { showNotification } = useNotification();
  const [openSubmitDialog, setOpenSubmitDialog] = useState(false);
  const [openApproveDialog, setOpenApproveDialog] = useState(false);
  const [openRejectDialog, setOpenRejectDialog] = useState(false);
  const [reason, setReason] = useState('');
  const { submit, approve, reject, loading: workflowLoading } = useWorkflow();

  const { data: tpSet, loading, refetch } = useAPI(() => tpGenerationAPI.getTPSet(id!), [id]);
  const { data: history } = useAPI(() => workflowAPI.getHistory(id!, 'tp_set'), [id]);

  const handleSubmit = async () => {
    try {
      await submit(id!, 'tp_set', reason);
      refetch();
      setOpenSubmitDialog(false);
      setReason('');
    } catch (error) {
      // Error handled in hook
    }
  };

  const handleApprove = async () => {
    try {
      await approve(id!, 'tp_set', reason);
      refetch();
      setOpenApproveDialog(false);
      setReason('');
    } catch (error) {
      // Error handled in hook
    }
  };

  const handleReject = async () => {
    try {
      await reject(id!, 'tp_set', reason);
      refetch();
      setOpenRejectDialog(false);
      setReason('');
    } catch (error) {
      // Error handled in hook
    }
  };

  const getStatusColor = (status: string) => {
    switch (status) {
      case 'DRAFT': return 'default';
      case 'UNDER_REVIEW': return 'warning';
      case 'APPROVED': return 'success';
      case 'REJECTED': return 'error';
      case 'ARCHIVED': return 'info';
      default: return 'default';
    }
  };

  return (
    <Container maxWidth="lg">
      {loading ? (
        <Typography>Loading...</Typography>
      ) : tpSet && (
        <>
          <Box sx={{ mb: 4 }}>
            <Button onClick={() => navigate(-1)} sx={{ mb: 2 }}>
              Back
            </Button>
            <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 2 }}>
              <Typography variant="h4" component="h1">
                TP Set v{tpSet.version_no}
              </Typography>
              <Chip label={tpSet.status} color={getStatusColor(tpSet.status)} />
            </Box>
            <Typography variant="body1" paragraph>
              CP: {tpSet.cp_code} - {tpSet.cp_text}
            </Typography>
            
            <Box sx={{ mt: 2 }}>
              {tpSet.status === 'DRAFT' && (
                <Button variant="contained" onClick={() => setOpenSubmitDialog(true)}>
                  Submit for Review
                </Button>
              )}
              {tpSet.status === 'UNDER_REVIEW' && (
                <>
                  <Button variant="contained" color="success" onClick={() => setOpenApproveDialog(true)} sx={{ mr: 1 }}>
                    Approve
                  </Button>
                  <Button variant="contained" color="error" onClick={() => setOpenRejectDialog(true)}>
                    Reject
                  </Button>
                </>
              )}
            </Box>
          </Box>

          <Grid container spacing={3} sx={{ mb: 4 }}>
            {tpSet.tps.map((tp) => (
              <Grid item xs={12} key={tp.id}>
                <Card>
                  <CardContent>
                    <Typography variant="h6" gutterBottom>
                      {tp.sequence_number}. {tp.title}
                    </Typography>
                    <Typography variant="body2" color="text.secondary" paragraph>
                      Estimated: {tp.estimated_weeks} weeks
                    </Typography>
                    <Typography variant="body2" paragraph>
                      <strong>Learning Objectives:</strong>
                    </Typography>
                    <ul>
                      {tp.learning_objectives.map((obj, idx) => (
                        <li key={idx}>{obj}</li>
                      ))}
                    </ul>
                  </CardContent>
                </Card>
              </Grid>
            ))}
          </Grid>

          <Typography variant="h5" gutterBottom>
            Workflow History
          </Typography>
          <Timeline>
            {history?.map((item) => (
              <TimelineItem key={item.id}>
                <TimelineSeparator>
                  <TimelineDot />
                  <TimelineConnector />
                </TimelineSeparator>
                <TimelineContent>
                  <Typography variant="body1">{item.action}</Typography>
                  <Typography variant="body2" color="text.secondary">
                    {item.user_name} - {new Date(item.created_at).toLocaleString()}
                  </Typography>
                  {item.reason && <Typography variant="caption">{item.reason}</Typography>}
                </TimelineContent>
              </TimelineItem>
            ))}
          </Timeline>

          {/* Submit Dialog */}
          <Dialog open={openSubmitDialog} onClose={() => setOpenSubmitDialog(false)} maxWidth="sm" fullWidth>
            <DialogTitle>Submit for Review</DialogTitle>
            <DialogContent>
              <TextField
                fullWidth
                multiline
                rows={4}
                label="Reason"
                value={reason}
                onChange={(e) => setReason(e.target.value)}
                sx={{ mt: 2 }}
              />
            </DialogContent>
            <DialogActions>
              <Button onClick={() => setOpenSubmitDialog(false)}>Cancel</Button>
              <Button onClick={handleSubmit} disabled={workflowLoading || !reason} variant="contained">
                Submit
              </Button>
            </DialogActions>
          </Dialog>

          {/* Approve Dialog */}
          <Dialog open={openApproveDialog} onClose={() => setOpenApproveDialog(false)} maxWidth="sm" fullWidth>
            <DialogTitle>Approve TP Set</DialogTitle>
            <DialogContent>
              <TextField
                fullWidth
                multiline
                rows={4}
                label="Reason"
                value={reason}
                onChange={(e) => setReason(e.target.value)}
                sx={{ mt: 2 }}
              />
            </DialogContent>
            <DialogActions>
              <Button onClick={() => setOpenApproveDialog(false)}>Cancel</Button>
              <Button onClick={handleApprove} disabled={workflowLoading || !reason} variant="contained" color="success">
                Approve
              </Button>
            </DialogActions>
          </Dialog>

          {/* Reject Dialog */}
          <Dialog open={openRejectDialog} onClose={() => setOpenRejectDialog(false)} maxWidth="sm" fullWidth>
            <DialogTitle>Reject TP Set</DialogTitle>
            <DialogContent>
              <TextField
                fullWidth
                multiline
                rows={4}
                label="Reason"
                value={reason}
                onChange={(e) => setReason(e.target.value)}
                sx={{ mt: 2 }}
              />
            </DialogContent>
            <DialogActions>
              <Button onClick={() => setOpenRejectDialog(false)}>Cancel</Button>
              <Button onClick={handleReject} disabled={workflowLoading || !reason} variant="contained" color="error">
                Reject
              </Button>
            </DialogActions>
          </Dialog>
        </>
      )}
    </Container>
  );
}
```

---

# SECTION 9 — Workflow Screens

## Pending Approvals Page

```typescript
// src/pages/workflow/WorkflowPendingPage.tsx
import { Box, Container, Typography, Card, CardContent, Button, Chip, Grid } from '@mui/material';
import { useNavigate } from 'react-router-dom';
import { workflowAPI } from '../../lib/api/workflow';
import { useAPI } from '../../hooks/useAPI';
import { PendingApproval } from '../../types/workflow';

export default function WorkflowPendingPage() {
  const navigate = useNavigate();
  const { data: pendingApprovals, loading } = useAPI(() => workflowAPI.getPendingApprovals('tp_set'));

  const getStatusColor = (status: string) => {
    switch (status) {
      case 'UNDER_REVIEW': return 'warning';
      default: return 'default';
    }
  };

  return (
    <Container maxWidth="lg">
      <Box sx={{ mb: 4 }}>
        <Typography variant="h4" component="h1">
          Pending Approvals
        </Typography>
        <Typography variant="body1" color="text.secondary">
          TP Sets awaiting review
        </Typography>
      </Box>

      {loading ? (
        <Typography>Loading...</Typography>
      ) : (
        <Grid container spacing={3}>
          {pendingApprovals?.map((approval) => (
            <Grid item xs={12} key={approval.artifact_id}>
              <Card>
                <CardContent>
                  <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 2 }}>
                    <Typography variant="h6">
                      TP Set - {approval.artifact_id}
                    </Typography>
                    <Chip label="Under Review" color={getStatusColor('UNDER_REVIEW')} />
                  </Box>
                  <Typography variant="body2" color="text.secondary" paragraph>
                    Created by: {approval.creator_name}
                  </Typography>
                  <Typography variant="body2" color="text.secondary" paragraph>
                    Submitted: {new Date(approval.submitted_at).toLocaleString()}
                  </Typography>
                  <Button
                    variant="contained"
                    onClick={() => navigate(`/tp-generation/tp-sets/${approval.artifact_id}`)}
                  >
                    Review
                  </Button>
                </CardContent>
              </Card>
            </Grid>
          ))}
        </Grid>
      )}
    </Container>
  );
}
```

## Workflow History Page

```typescript
// src/pages/workflow/WorkflowHistoryPage.tsx
import { Box, Container, Typography, Card, CardContent, Chip, Grid, Timeline, TimelineItem, TimelineSeparator, TimelineConnector, TimelineContent, TimelineDot } from '@mui/material';
import { useNavigate } from 'react-router-dom';
import { workflowAPI } from '../../lib/api/workflow';
import { useAPI } from '../../hooks/useAPI';
import { WorkflowHistory } from '../../types/workflow';

export default function WorkflowHistoryPage() {
  const navigate = useNavigate();
  const { data: history, loading } = useAPI(() => workflowAPI.getHistory('all', 'tp_set'));

  const getStatusColor = (status: string) => {
    switch (status) {
      case 'DRAFT': return 'default';
      case 'UNDER_REVIEW': return 'warning';
      case 'APPROVED': return 'success';
      case 'REJECTED': return 'error';
      case 'ARCHIVED': return 'info';
      default: return 'default';
    }
  };

  const getActionColor = (action: string) => {
    switch (action) {
      case 'CREATE': return 'info';
      case 'SUBMIT': return 'warning';
      case 'APPROVE': return 'success';
      case 'REJECT': return 'error';
      case 'ARCHIVE': return 'default';
      default: return 'info';
    }
  };

  return (
    <Container maxWidth="lg">
      <Box sx={{ mb: 4 }}>
        <Typography variant="h4" component="h1">
          Workflow History
        </Typography>
        <Typography variant="body1" color="text.secondary">
          All TP Set workflow actions
        </Typography>
      </Box>

      {loading ? (
        <Typography>Loading...</Typography>
      ) : (
        <Timeline>
          {history?.map((item) => (
            <TimelineItem key={item.id}>
              <TimelineSeparator>
                <TimelineDot color={getActionColor(item.action)} />
                <TimelineConnector />
              </TimelineSeparator>
              <TimelineContent>
                <Card sx={{ mb: 2 }}>
                  <CardContent>
                    <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 2 }}>
                      <Typography variant="h6">
                        {item.action}
                      </Typography>
                      <Chip label={item.to_state} color={getStatusColor(item.to_state)} size="small" />
                    </Box>
                    <Typography variant="body2" color="text.secondary" paragraph>
                      Artifact: {item.artifact_type} - {item.artifact_id}
                    </Typography>
                    <Typography variant="body2" color="text.secondary" paragraph>
                      User: {item.user_name}
                    </Typography>
                    <Typography variant="body2" color="text.secondary" paragraph>
                      Time: {new Date(item.created_at).toLocaleString()}
                    </Typography>
                    {item.reason && (
                      <Typography variant="body2" paragraph>
                        <strong>Reason:</strong> {item.reason}
                      </Typography>
                    )}
                    <Button
                      size="small"
                      onClick={() => navigate(`/tp-generation/tp-sets/${item.artifact_id}`)}
                    >
                      View Artifact
                    </Button>
                  </CardContent>
                </Card>
              </TimelineContent>
            </TimelineItem>
          ))}
        </Timeline>
      )}
    </Container>
  );
}
```

---

# SECTION 10 — Route Map

## Complete Route Map

| Path | Component | Auth Required | Role Required | Description |
|------|-----------|---------------|--------------|-------------|
| `/login` | LoginPage | No | None | User login |
| `/` | DashboardPage | Yes | All | Dashboard |
| `/curriculum` | CurriculumPage | Yes | All | List subjects |
| `/curriculum/subjects/:id` | SubjectPage | Yes | All | Subject detail with CP tree |
| `/curriculum/cp/:id` | CPPage | Yes | All | CP detail with TP generation |
| `/tp-generation` | TPGenerationPage | Yes | All | TP generation landing |
| `/tp-generation/cp/:cpId/generate` | TPSetDetailPage (generate mode) | Yes | TEACHER+ | Generate TP Set |
| `/tp-generation/tp-sets` | TPSetPage | Yes | All | List TP Sets |
| `/tp-generation/tp-sets/:id` | TPSetDetailPage (view mode) | Yes | All | TP Set detail |
| `/workflow/pending` | WorkflowPendingPage | Yes | SCHOOL_ADMIN+ | Pending approvals |
| `/workflow/history` | WorkflowHistoryPage | Yes | All | Workflow history |
| `/profile` | ProfilePage | Yes | All | User profile |

---

# SECTION 11 — Page Inventory

## Page List

### Auth Pages

| Page | Path | Status | Priority |
|------|------|--------|----------|
| Login | `/login` | MVP | High |
| Register | `/register` | Future | Low |
| Forgot Password | `/forgot-password` | Future | Low |

### Curriculum Pages

| Page | Path | Status | Priority |
|------|------|--------|----------|
| Curriculum List | `/curriculum` | MVP | High |
| Subject Detail | `/curriculum/subjects/:id` | MVP | High |
| CP Detail | `/curriculum/cp/:id` | MVP | High |

### TP Generation Pages

| Page | Path | Status | Priority |
|------|------|--------|----------|
| TP Generation Landing | `/tp-generation` | MVP | High |
| TP Set List | `/tp-generation/tp-sets` | MVP | High |
| TP Set Detail | `/tp-generation/tp-sets/:id` | MVP | High |
| TP Set Generate | `/tp-generation/cp/:cpId/generate` | MVP | High |

### Workflow Pages

| Page | Path | Status | Priority |
|------|------|--------|----------|
| Pending Approvals | `/workflow/pending` | MVP | High |
| Workflow History | `/workflow/history` | MVP | High |

### Profile Pages

| Page | Path | Status | Priority |
|------|------|--------|----------|
| Profile | `/profile` | MVP | Medium |
| Settings | `/settings` | Future | Low |

---

# SECTION 12 — Component Structure

## Component Hierarchy

### Layout Components

```
Layout
├── AppBar
├── Sidebar
└── MainContent
```

### Common Components

```
Button
Card
Dialog
Loading
ErrorBoundary
```

### Curriculum Components

```
SubjectCard
├── SubjectCode
├── SubjectName
└── ViewButton

CPList
├── CPItem
│   ├── CPCode
│   ├── CPText
│   └── GenerateButton
└── Pagination

CPTree
├── SubjectNode
├── PhaseNode
├── ElementNode
├── SubelementNode
└── CPNode
```

### TP Generation Components

```
TPSetCard
├── TPSetInfo
├── StatusBadge
├── VersionBadge
└── ViewButton

TPItemCard
├── SequenceNumber
├── Title
├── LearningObjectives
├── EstimatedWeeks
└── Prerequisites

GenerationForm
├── DurationInput
├── FocusAreasInput
└── GenerateButton

TPSetDetail
├── TPSetHeader
├── TPSetInfo
├── TPList
├── ActionButtons
└── HistoryTimeline
```

### Workflow Components

```
StatusBadge
├── StatusLabel
└── StatusColor

HistoryTimeline
├── TimelineItem
│   ├── ActionLabel
│   ├── UserInfo
│   ├── Timestamp
│   └── Reason
└── TimelineConnector

ActionButtons
├── SubmitButton
├── ApproveButton
├── RejectButton
└── ArchiveButton
```

---

# SECTION 13 — API Integration Map

## API Integration Summary

| Component | API Module | Endpoints Used |
|-----------|------------|----------------|
| LoginPage | authAPI | login |
| Layout | authAPI | me, logout |
| CurriculumPage | curriculumAPI | getSubjects |
| SubjectPage | curriculumAPI | getSubjectTree |
| CPPage | curriculumAPI | getCP |
| TPGenerationPage | curriculumAPI | getSubjects |
| CPPage (generate) | tpGenerationAPI | generateTPSet |
| TPSetDetailPage | tpGenerationAPI | getTPSet |
| TPSetDetailPage (actions) | workflowAPI | submit, approve, reject, archive |
| TPSetDetailPage (history) | workflowAPI | getHistory |
| WorkflowPendingPage | workflowAPI | getPendingApprovals |
| WorkflowHistoryPage | workflowAPI | getHistory |

## API Call Flow

### Authentication Flow

```
LoginPage
  ↓ authAPI.login()
  → POST /api/v1/public/login
  ← access_token, refresh_token, user
  → Store in localStorage
  → Navigate to dashboard
```

### TP Generation Flow

```
CPPage
  ↓ curriculumAPI.getCP()
  → GET /api/v1/public/curriculum/cp/:id
  ← CP data
  ↓ tpGenerationAPI.generateTPSet()
  → POST /api/v1/curriculum/cp/:cpId/tp-sets/generate
  ← TP Set data
  → Navigate to TP Set detail
```

### Workflow Action Flow

```
TPSetDetailPage
  ↓ workflowAPI.submit()
  → POST /api/v1/workflow/tp_set/:id/submit
  ← Success
  → Refetch TP Set data
  → Update UI
```

---

# SECTION 14 — Appendix

## TypeScript Type Definitions

### Auth Types

```typescript
// src/types/auth.ts
export interface User {
  id: string;
  email: string;
  name: string;
  role: 'SYSTEM_ADMIN' | 'SCHOOL_ADMIN' | 'TEACHER';
  school_id?: string;
  school_name?: string;
}

export interface LoginRequest {
  email: string;
  password: string;
}

export interface LoginResponse {
  access_token: string;
  refresh_token: string;
  token_type: string;
  expires_in: number;
  user: User;
}

export interface RefreshTokenRequest {
  refresh_token: string;
}

export interface RefreshTokenResponse {
  access_token: string;
  refresh_token: string;
  token_type: string;
  expires_in: number;
}
```

### Curriculum Types

```typescript
// src/types/curriculum.ts
export interface Subject {
  id: string;
  code: string;
  name: string;
  name_en?: string;
}

export interface SubjectTree {
  id: string;
  code: string;
  name: string;
  phases: Phase[];
}

export interface Phase {
  id: string;
  code: string;
  name: string;
  grade_levels: string[];
  elements: Element[];
}

export interface Element {
  id: string;
  code: string;
  name: string;
  name_en?: string;
  subelements: Subelement[];
}

export interface Subelement {
  id: string;
  code: string;
  name: string;
  name_en?: string;
  cps: CP[];
}

export interface CP {
  id: string;
  code: string;
  text: string;
  text_en?: string;
  fase?: string;
  subelement: Subelement;
  element: Element;
  phase: Phase;
  subject: Subject;
}

export interface CPListResponse {
  data: CP[];
  pagination: {
    page: number;
    page_size: number;
    total: number;
    total_pages: number;
  };
}
```

### TP Generation Types

```typescript
// src/types/tp-generation.ts
export interface TPSet {
  id: string;
  cp_id: string;
  cp_code: string;
  cp_text: string;
  version_no: number;
  status: 'DRAFT' | 'UNDER_REVIEW' | 'APPROVED' | 'REJECTED' | 'ARCHIVED';
  generation_source: string;
  generated_by: string;
  generated_by_name: string;
  ai_generation_id?: string;
  approved_by?: string;
  approved_by_name?: string;
  approved_at?: string;
  tps: TP[];
  created_at: string;
  updated_at: string;
}

export interface TP {
  id: string;
  tp_set_id: string;
  sequence_number: number;
  title: string;
  learning_objectives: string[];
  estimated_weeks: number;
  prerequisites?: string[];
  status: string;
  created_at: string;
  updated_at: string;
}

export interface GenerateTPSetRequest {
  preferences: {
    duration_weeks?: number;
    focus_areas?: string[];
  };
}

export interface GenerateTPSetResponse {
  id: string;
  cp_id: string;
  version_no: number;
  status: string;
  generation_source: string;
  generated_by: string;
  ai_generation_id?: string;
  tps: TP[];
  created_at: string;
  updated_at: string;
}

export interface TPSetListResponse {
  data: TPSet[];
  pagination: {
    page: number;
    page_size: number;
    total: number;
    total_pages: number;
  };
}
```

### Workflow Types

```typescript
// src/types/workflow.ts
export interface WorkflowHistory {
  id: string;
  artifact_id: string;
  artifact_type: string;
  action: string;
  from_state: string;
  to_state: string;
  user_id: string;
  user_name: string;
  reason?: string;
  metadata?: Record<string, any>;
  ai_generation_id?: string;
  created_at: string;
}

export interface PendingApproval {
  artifact_id: string;
  artifact_type: string;
  creator_id: string;
  creator_name: string;
  submitted_at: string;
}
```

## Build Configuration

### Vite Config

```typescript
// vite.config.ts
import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import path from 'path';

export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
    },
  },
  server: {
    port: 3000,
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
    },
  },
});
```

## Environment Variables

```bash
# .env
VITE_API_BASE_URL=http://localhost:8080
```

## Testing Strategy

### Unit Tests

- Component rendering
- Hook behavior
- API client functions
- Utility functions

### Integration Tests

- Page navigation
- Form submissions
- API integration
- Authentication flow

### E2E Tests

- Complete user flows
- Cross-browser testing
- Performance testing

## Future Enhancements

### Wave 2

- ATP generation screens
- Modul Ajar generation screens
- Assessment screens
- Rubric screens
- Advanced search and filtering
- Bulk operations
- Export functionality
- Offline support
- PWA capabilities

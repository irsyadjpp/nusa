/**
 * Users Query Service
 * Provides query operations for user data using TanStack Query
 */

import { useQuery, UseQueryOptions } from '@tanstack/react-query';
import * as usersApi from '@/api/users';

// Query Keys
export const usersKeys = {
  all: ['users'] as const,
  list: (params?: any) => ['users', 'list', params] as const,
  detail: (id: string) => ['users', 'detail', id] as const,
};

/**
 * User Queries
 */
export const useUsers = (
  params?: any,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: usersKeys.list(params),
    queryFn: () => usersApi.listUsers(params),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

export const useUser = (
  id: string,
  options?: Omit<UseQueryOptions<any, Error, any>, 'queryKey' | 'queryFn'>
) => {
  return useQuery({
    queryKey: usersKeys.detail(id),
    queryFn: () => usersApi.getUser(id),
    staleTime: 300000, // 5 minutes
    ...options,
  });
};

/**
 * Query invalidation functions
 */
export const invalidateUserQueries = (queryClient: any) => {
  queryClient.invalidateQueries({ queryKey: usersKeys.all });
};

/**
 * Announcements Command Service
 * Provides command operations for announcement data using TanStack Query mutations
 */

import { useMutation, useQueryClient, UseMutationOptions } from '@tanstack/react-query';
import * as announcementsApi from '@/api/announcements';
import { announcementsKeys } from '@/services/queries/AnnouncementsQueryService';

/**
 * Create Announcement Mutation
 */
export const useCreateAnnouncement = (
  options?: Omit<UseMutationOptions<any, Error, any>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: any) => announcementsApi.createAnnouncement(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: announcementsKeys.all });
    },
    ...options,
  });
};

/**
 * Update Announcement Mutation
 */
export const useUpdateAnnouncement = (
  options?: Omit<UseMutationOptions<any, Error, any>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, data }: { id: string; data: any }) =>
      announcementsApi.updateAnnouncement(id, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: announcementsKeys.all });
    },
    ...options,
  });
};

/**
 * Delete Announcement Mutation
 */
export const useDeleteAnnouncement = (
  options?: Omit<UseMutationOptions<void, Error, string>, 'mutationFn'>
) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => announcementsApi.deleteAnnouncement(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: announcementsKeys.all });
    },
    ...options,
  });
};

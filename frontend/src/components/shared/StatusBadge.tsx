/**
 * StatusBadge Component
 * Displays a status badge with appropriate color coding
 */

import { Chip, ChipProps } from '@mui/material';

interface StatusBadgeProps extends Omit<ChipProps, 'color'> {
  status: string;
}

export const StatusBadge = ({ status, ...props }: StatusBadgeProps) => {
  const getStatusColor = (status: string): ChipProps['color'] => {
    const statusLower = status.toLowerCase();
    if (statusLower === 'approved' || statusLower === 'completed' || statusLower === 'published') {
      return 'success';
    }
    if (statusLower === 'pending' || statusLower === 'in_review' || statusLower === 'draft') {
      return 'warning';
    }
    if (statusLower === 'rejected' || statusLower === 'failed' || statusLower === 'error') {
      return 'error';
    }
    return 'default';
  };

  return (
    <Chip
      label={status}
      color={getStatusColor(status)}
      size="small"
      {...props}
    />
  );
};

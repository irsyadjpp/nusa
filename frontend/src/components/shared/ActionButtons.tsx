/**
 * ActionButtons Component
 * Displays a set of action buttons for CRUD operations
 */

import { Stack, IconButton, Tooltip } from '@mui/material';
import { Edit, Delete, Preview, Check, Close } from '@mui/icons-material';

interface ActionButtonsProps {
  onEdit?: () => void;
  onDelete?: () => void;
  onPreview?: () => void;
  onApprove?: () => void;
  onReject?: () => void;
  showEdit?: boolean;
  showDelete?: boolean;
  showPreview?: boolean;
  showApprove?: boolean;
  showReject?: boolean;
  size?: 'small' | 'medium' | 'large';
}

export const ActionButtons = ({
  onEdit,
  onDelete,
  onPreview,
  onApprove,
  onReject,
  showEdit = true,
  showDelete = true,
  showPreview = true,
  showApprove = false,
  showReject = false,
  size = 'small',
}: ActionButtonsProps) => {
  return (
    <Stack direction="row" spacing={1}>
      {showEdit && onEdit && (
        <Tooltip title="Edit">
          <IconButton size={size} onClick={onEdit} color="primary">
            <Edit fontSize={size} />
          </IconButton>
        </Tooltip>
      )}
      {showPreview && onPreview && (
        <Tooltip title="Preview">
          <IconButton size={size} onClick={onPreview} color="info">
            <Preview fontSize={size} />
          </IconButton>
        </Tooltip>
      )}
      {showApprove && onApprove && (
        <Tooltip title="Approve">
          <IconButton size={size} onClick={onApprove} color="success">
            <Check fontSize={size} />
          </IconButton>
        </Tooltip>
      )}
      {showReject && onReject && (
        <Tooltip title="Reject">
          <IconButton size={size} onClick={onReject} color="error">
            <Close fontSize={size} />
          </IconButton>
        </Tooltip>
      )}
      {showDelete && onDelete && (
        <Tooltip title="Delete">
          <IconButton size={size} onClick={onDelete} color="error">
            <Delete fontSize={size} />
          </IconButton>
        </Tooltip>
      )}
    </Stack>
  );
};

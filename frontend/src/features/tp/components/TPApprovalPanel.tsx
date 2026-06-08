/**
 * TP Approval Panel Component
 * Panel for approving or rejecting Teaching Plans
 */

import { Box, Typography, Button, Stack, Alert } from '@mui/material';
import { Check as CheckIcon, Close as CloseIcon } from '@mui/icons-material';

interface TPApprovalPanelProps {
  onApprove: () => void;
  onReject: () => void;
  canApprove?: boolean;
  canReject?: boolean;
  loading?: boolean;
}

export const TPApprovalPanel = ({
  onApprove,
  onReject,
  canApprove = true,
  canReject = true,
  loading = false,
}: TPApprovalPanelProps) => {
  return (
    <Box sx={{ p: 3, bgcolor: 'background.paper', borderRadius: 1 }}>
      <Typography variant="h6" gutterBottom>
        Approval Actions
      </Typography>
      <Alert severity="info" sx={{ mb: 2 }}>
        Review the TP details before approving or rejecting.
      </Alert>
      <Stack direction="row" spacing={2}>
        <Button
          variant="contained"
          color="success"
          startIcon={<CheckIcon />}
          onClick={onApprove}
          disabled={!canApprove || loading}
        >
          {loading ? 'Processing...' : 'Approve'}
        </Button>
        <Button
          variant="contained"
          color="error"
          startIcon={<CloseIcon />}
          onClick={onReject}
          disabled={!canReject || loading}
        >
          {loading ? 'Processing...' : 'Reject'}
        </Button>
      </Stack>
    </Box>
  );
};

/**
 * Evidence Review Panel Component
 * Panel for reviewing and approving/rejecting evidence
 */

import { Box, Typography, Paper, Divider, Button, Stack, Alert } from '@mui/material';
import { Check as CheckIcon, Close as CloseIcon } from '@mui/icons-material';
import { Evidence } from '@/api/evidence';

interface EvidenceReviewPanelProps {
  evidence: Evidence;
  onApprove: () => void;
  onReject: () => void;
  canApprove?: boolean;
  canReject?: boolean;
  loading?: boolean;
}

export const EvidenceReviewPanel = ({
  evidence,
  onApprove,
  onReject,
  canApprove = true,
  canReject = true,
  loading = false,
}: EvidenceReviewPanelProps) => {
  return (
    <Paper elevation={2} sx={{ p: 3 }}>
      <Typography variant="h5" gutterBottom>
        Evidence Review
      </Typography>
      <Divider sx={{ my: 2 }} />
      <Stack spacing={2}>
        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            Title
          </Typography>
          <Typography variant="body1">{evidence.title || `Evidence ${evidence.id}`}</Typography>
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            Student ID
          </Typography>
          <Typography variant="body1">{evidence.student_id}</Typography>
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            Assessment ID
          </Typography>
          <Typography variant="body1">{evidence.assessment_id}</Typography>
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            Status
          </Typography>
          <Typography variant="body1">{evidence.status}</Typography>
        </Box>

        <Divider />

        <Alert severity="info">
          Review the evidence details before approving or rejecting.
        </Alert>

        <Stack direction="row" spacing={2} justifyContent="flex-end">
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
      </Stack>
    </Paper>
  );
};

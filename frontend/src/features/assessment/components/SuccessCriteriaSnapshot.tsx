/**
 * Success Criteria Snapshot Component
 * Displays the success criteria snapshot from a TP
 */

import { Box, Typography, Paper, Divider } from '@mui/material';

interface SuccessCriteriaSnapshotProps {
  snapshot: any;
}

export const SuccessCriteriaSnapshot = ({ snapshot }: SuccessCriteriaSnapshotProps) => {
  return (
    <Paper elevation={1} sx={{ p: 2, mb: 2 }}>
      <Typography variant="subtitle2" gutterBottom>
        Success Criteria Snapshot
      </Typography>
      <Divider sx={{ my: 1 }} />
      <Typography variant="body2" whiteSpace="pre-wrap">
        {typeof snapshot === 'string' ? snapshot : JSON.stringify(snapshot, null, 2)}
      </Typography>
    </Paper>
  );
};

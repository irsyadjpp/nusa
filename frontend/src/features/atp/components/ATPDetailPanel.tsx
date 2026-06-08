/**
 * ATP Detail Panel Component
 * Displays detailed information about an ATP
 */

import { Box, Typography, Paper, Divider, Chip } from '@mui/material';
import { ATP } from '@/api/atp';

interface ATPDetailPanelProps {
  atp: ATP;
}

export const ATPDetailPanel = ({ atp }: ATPDetailPanelProps) => {
  return (
    <Paper elevation={2} sx={{ p: 3 }}>
      <Typography variant="h5" gutterBottom>
        ATP Details
      </Typography>
      <Divider sx={{ my: 2 }} />
      <Stack spacing={2}>
        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            Status
          </Typography>
          <Chip
            label={atp.status}
            size="small"
            color={atp.status === 'approved' ? 'success' : atp.status === 'draft' ? 'default' : 'warning'}
          />
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            Sequence Number
          </Typography>
          <Typography variant="body1">{atp.sequence_number}</Typography>
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            Week Number
          </Typography>
          <Typography variant="body1">{atp.week_number}</Typography>
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            Estimated Hours
          </Typography>
          <Typography variant="body1">{atp.estimated_hours} hours</Typography>
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            ATP Set ID
          </Typography>
          <Typography variant="body1">{atp.atp_set_id}</Typography>
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            TP ID
          </Typography>
          <Typography variant="body1">{atp.tp_id}</Typography>
        </Box>

        <Divider />

        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            Created At
          </Typography>
          <Typography variant="body1">
            {new Date(atp.created_at).toLocaleString()}
          </Typography>
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            Updated At
          </Typography>
          <Typography variant="body1">
            {new Date(atp.updated_at).toLocaleString()}
          </Typography>
        </Box>
      </Stack>
    </Paper>
  );
};

/**
 * TP Preview Panel Component
 * Displays a preview of the Teaching Plan
 */

import { Box, Typography, Paper, Divider, Chip } from '@mui/material';
import { TP } from '@/api/tp';

interface TPPreviewPanelProps {
  tp: TP;
}

export const TPPreviewPanel = ({ tp }: TPPreviewPanelProps) => {
  return (
    <Paper elevation={2} sx={{ p: 3 }}>
      <Typography variant="h5" gutterBottom>
        {tp.title}
      </Typography>
      <Divider sx={{ my: 2 }} />
      <Stack spacing={2}>
        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            Status
          </Typography>
          <Chip
            label={tp.status}
            size="small"
            color={tp.status === 'approved' ? 'success' : tp.status === 'draft' ? 'default' : 'warning'}
          />
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            Sequence Number
          </Typography>
          <Typography variant="body1">{tp.sequence_number}</Typography>
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            Subject
          </Typography>
          <Typography variant="body1">{tp.subject_id}</Typography>
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            Phase
          </Typography>
          <Typography variant="body1">{tp.phase_id}</Typography>
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            Estimated Weeks
          </Typography>
          <Typography variant="body1">{tp.estimated_weeks}</Typography>
        </Box>

        <Divider />

        <Box>
          <Typography variant="subtitle2" color="text.secondary" gutterBottom>
            Learning Objectives
          </Typography>
          <Typography variant="body1" whiteSpace="pre-wrap">
            {typeof tp.learning_objectives === 'string'
              ? tp.learning_objectives
              : JSON.stringify(tp.learning_objectives, null, 2)}
          </Typography>
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary" gutterBottom>
            Time Allocation
          </Typography>
          <Typography variant="body1" whiteSpace="pre-wrap">
            {typeof tp.time_allocation === 'string'
              ? tp.time_allocation
              : JSON.stringify(tp.time_allocation, null, 2)}
          </Typography>
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary" gutterBottom>
            Prerequisites
          </Typography>
          <Typography variant="body1" whiteSpace="pre-wrap">
            {typeof tp.prerequisites === 'string'
              ? tp.prerequisites
              : JSON.stringify(tp.prerequisites, null, 2)}
          </Typography>
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary" gutterBottom>
            Success Criteria
          </Typography>
          <Typography variant="body1" whiteSpace="pre-wrap">
            {typeof tp.success_criteria === 'string'
              ? tp.success_criteria
              : JSON.stringify(tp.success_criteria, null, 2)}
          </Typography>
        </Box>
      </Stack>
    </Paper>
  );
};

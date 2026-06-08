/**
 * Rubric Preview Panel Component
 * Displays a preview of the rubric
 */

import { Box, Typography, Paper, Divider, Chip } from '@mui/material';
import { Rubric } from '@/api/rubric';

interface RubricPreviewPanelProps {
  rubric: Rubric;
}

export const RubricPreviewPanel = ({ rubric }: RubricPreviewPanelProps) => {
  return (
    <Paper elevation={2} sx={{ p: 3 }}>
      <Typography variant="h5" gutterBottom>
        {rubric.title}
      </Typography>
      <Divider sx={{ my: 2 }} />
      <Stack spacing={2}>
        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            Status
          </Typography>
          <Chip
            label={rubric.status}
            size="small"
            color={rubric.status === 'approved' ? 'success' : rubric.status === 'draft' ? 'default' : 'warning'}
          />
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            Assessment ID
          </Typography>
          <Typography variant="body1">{rubric.assessment_id}</Typography>
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary" gutterBottom>
            Description
          </Typography>
          <Typography variant="body1" whiteSpace="pre-wrap">
            {rubric.description || 'No description provided'}
          </Typography>
        </Box>

        <Divider />

        <Box>
          <Typography variant="subtitle2" color="text.secondary" gutterBottom>
            Criteria
          </Typography>
          <Typography variant="body2" whiteSpace="pre-wrap">
            {typeof rubric.criteria === 'string'
              ? rubric.criteria
              : JSON.stringify(rubric.criteria, null, 2)}
          </Typography>
        </Box>

        <Divider />

        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            Created At
          </Typography>
          <Typography variant="body1">
            {new Date(rubric.created_at).toLocaleString()}
          </Typography>
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            Updated At
          </Typography>
          <Typography variant="body1">
            {new Date(rubric.updated_at).toLocaleString()}
          </Typography>
        </Box>
      </Stack>
    </Paper>
  );
};

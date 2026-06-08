/**
 * Evaluation Workspace Header Component
 * Displays the header for the Evaluation workspace with actions and filters
 */

import { Box, Typography, Button, Stack } from '@mui/material';
import { Refresh as RefreshIcon } from '@mui/icons-material';

interface EvaluationWorkspaceHeaderProps {
  title: string;
  onRefresh?: () => void;
}

export const EvaluationWorkspaceHeader = ({
  title,
  onRefresh,
}: EvaluationWorkspaceHeaderProps) => {
  return (
    <Box sx={{ mb: 3 }}>
      <Stack direction="row" justifyContent="space-between" alignItems="center">
        <Typography variant="h4" component="h1">
          {title}
        </Typography>
        <Stack direction="row" spacing={2}>
          {onRefresh && (
            <Button variant="outlined" startIcon={<RefreshIcon />} onClick={onRefresh}>
              Refresh
            </Button>
          )}
        </Stack>
      </Stack>
    </Box>
  );
};

/**
 * TP Workspace Header Component
 * Displays the header for the TP workspace with actions and filters
 */

import { Box, Typography, Button, Stack } from '@mui/material';
import { Add as AddIcon } from '@mui/icons-material';

interface TPWorkspaceHeaderProps {
  title: string;
  onCreateNew?: () => void;
  onRefresh?: () => void;
  showCreateButton?: boolean;
}

export const TPWorkspaceHeader = ({
  title,
  onCreateNew,
  onRefresh,
  showCreateButton = true,
}: TPWorkspaceHeaderProps) => {
  return (
    <Box sx={{ mb: 3 }}>
      <Stack direction="row" justifyContent="space-between" alignItems="center">
        <Typography variant="h4" component="h1">
          {title}
        </Typography>
        <Stack direction="row" spacing={2}>
          {onRefresh && (
            <Button variant="outlined" onClick={onRefresh}>
              Refresh
            </Button>
          )}
          {showCreateButton && onCreateNew && (
            <Button variant="contained" startIcon={<AddIcon />} onClick={onCreateNew}>
              Create New TP
            </Button>
          )}
        </Stack>
      </Stack>
    </Box>
  );
};

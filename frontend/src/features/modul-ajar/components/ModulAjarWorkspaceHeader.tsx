/**
 * Modul Ajar Workspace Header Component
 * Displays the header for the Modul Ajar workspace with actions and filters
 */

import { Box, Typography, Button, Stack } from '@mui/material';
import { Add as AddIcon } from '@mui/icons-material';

interface ModulAjarWorkspaceHeaderProps {
  title: string;
  onCreateNew?: () => void;
  onRefresh?: () => void;
  showCreateButton?: boolean;
}

export const ModulAjarWorkspaceHeader = ({
  title,
  onCreateNew,
  onRefresh,
  showCreateButton = true,
}: ModulAjarWorkspaceHeaderProps) => {
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
              Create New Modul Ajar
            </Button>
          )}
        </Stack>
      </Stack>
    </Box>
  );
};

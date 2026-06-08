/**
 * Narrative Report Builder Header Component
 * Displays the header for the Narrative Report Builder with actions
 */

import { Box, Typography, Button, Stack } from '@mui/material';
import { Add as AddIcon } from '@mui/icons-material';

interface NarrativeReportBuilderHeaderProps {
  title: string;
  onCreateNew?: () => void;
  onRefresh?: () => void;
  showCreateButton?: boolean;
}

export const NarrativeReportBuilderHeader = ({
  title,
  onCreateNew,
  onRefresh,
  showCreateButton = true,
}: NarrativeReportBuilderHeaderProps) => {
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
              Create New Report
            </Button>
          )}
        </Stack>
      </Stack>
    </Box>
  );
};

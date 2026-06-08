/**
 * Achievement Dashboard Header Component
 * Displays the header for the Achievement Dashboard
 */

import { Box, Typography, Button, Stack } from '@mui/material';
import { Download as DownloadIcon } from '@mui/icons-material';

interface AchievementDashboardHeaderProps {
  title: string;
  onExport?: () => void;
  onRefresh?: () => void;
}

export const AchievementDashboardHeader = ({
  title,
  onExport,
  onRefresh,
}: AchievementDashboardHeaderProps) => {
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
          {onExport && (
            <Button variant="contained" startIcon={<DownloadIcon />} onClick={onExport}>
              Export Report
            </Button>
          )}
        </Stack>
      </Stack>
    </Box>
  );
};

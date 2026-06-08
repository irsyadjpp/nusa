/**
 * Evidence Workspace Header Component
 * Displays the header for the Evidence workspace with actions and filters
 */

import { Box, Typography, Button, Stack } from '@mui/material';
import { Upload as UploadIcon } from '@mui/icons-material';

interface EvidenceWorkspaceHeaderProps {
  title: string;
  onUpload?: () => void;
  onRefresh?: () => void;
  showUploadButton?: boolean;
}

export const EvidenceWorkspaceHeader = ({
  title,
  onUpload,
  onRefresh,
  showUploadButton = true,
}: EvidenceWorkspaceHeaderProps) => {
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
          {showUploadButton && onUpload && (
            <Button variant="contained" startIcon={<UploadIcon />} onClick={onUpload}>
              Upload Evidence
            </Button>
          )}
        </Stack>
      </Stack>
    </Box>
  );
};

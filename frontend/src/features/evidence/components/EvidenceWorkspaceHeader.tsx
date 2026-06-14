/**
 * Evidence Workspace Header Component
 * Displays the header for the Evidence workspace with actions and filters
 */

import { Box, Typography, Button } from '@mui/material';
import { Upload as UploadIcon } from '@mui/icons-material';

interface EvidenceWorkspaceHeaderProps {
  title: string;
  onUpload?: () => void;
  onUploadNew?: () => void;
  onRefresh?: () => void;
  showUploadButton?: boolean;
}

export const EvidenceWorkspaceHeader = ({
  title,
  onUpload,
  onUploadNew,
  onRefresh,
  showUploadButton = true,
}: EvidenceWorkspaceHeaderProps) => {
  return (
    <Box sx={{ mb: 3 }}>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', gap: 2 }}>
        <Typography variant="h4" component="h1">
          {title}
        </Typography>
        <Box sx={{ display: 'flex', gap: 2 }}>
          {onRefresh && (
            <Button variant="outlined" onClick={onRefresh}>
              Refresh
            </Button>
          )}
          {showUploadButton && (onUpload || onUploadNew) && (
            <Button variant="contained" startIcon={<UploadIcon />} onClick={onUpload || onUploadNew}>
              Upload Evidence
            </Button>
          )}
        </Box>
      </Box>
    </Box>
  );
};

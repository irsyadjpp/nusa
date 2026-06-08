import React from 'react';
import { Box, Button, IconButton, Tooltip, Menu, MenuItem, Divider } from '@mui/material';
import { Refresh as RefreshIcon, Download as DownloadIcon, Share as ShareIcon, MoreVert as MoreVertIcon, Print as PrintIcon } from '@mui/icons-material';

interface ReportActionsProps {
  reportId: string;
  onRefreshAchievement?: () => void;
  onDownload?: () => void;
  onShare?: () => void;
  onPrint?: () => void;
}

export const ReportActions: React.FC<ReportActionsProps> = ({
  reportId,
  onRefreshAchievement,
  onDownload,
  onShare,
  onPrint,
}) => {
  const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);
  const [refreshing, setRefreshing] = React.useState(false);

  const handleMenuClick = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const handleMenuClose = () => {
    setAnchorEl(null);
  };

  const handleRefreshAchievement = async () => {
    setRefreshing(true);
    try {
      await fetch(`/api/v1/reporting/narrative-reports/${reportId}/refresh-achievement`, {
        method: 'POST',
      });
      if (onRefreshAchievement) onRefreshAchievement();
    } catch (error) {
      console.error('Failed to refresh achievement data:', error);
    } finally {
      setRefreshing(false);
    }
  };

  return (
    <Box display="flex" gap={1}>
      <Tooltip title="Refresh Achievement Data">
        <Button
          variant="outlined"
          startIcon={<RefreshIcon />}
          onClick={handleRefreshAchievement}
          disabled={refreshing}
          size="small"
        >
          {refreshing ? 'Refreshing...' : 'Refresh Achievement'}
        </Button>
      </Tooltip>

      <Tooltip title="Download Report">
        <IconButton onClick={onDownload} size="small">
          <DownloadIcon />
        </IconButton>
      </Tooltip>

      <Tooltip title="Print Report">
        <IconButton onClick={onPrint} size="small">
          <PrintIcon />
        </IconButton>
      </Tooltip>

      <Tooltip title="Share Report">
        <IconButton onClick={onShare} size="small">
          <ShareIcon />
        </IconButton>
      </Tooltip>

      <Tooltip title="More Options">
        <IconButton onClick={handleMenuClick} size="small">
          <MoreVertIcon />
        </IconButton>
      </Tooltip>

      <Menu
        anchorEl={anchorEl}
        open={Boolean(anchorEl)}
        onClose={handleMenuClose}
      >
        <MenuItem onClick={handleMenuClose}>View Version History</MenuItem>
        <MenuItem onClick={handleMenuClose}>Duplicate Report</MenuItem>
        <Divider />
        <MenuItem onClick={handleMenuClose} sx={{ color: 'error.main' }}>Delete Report</MenuItem>
      </Menu>
    </Box>
  );
};

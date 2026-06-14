import React from 'react';
import { Box, Typography, List, ListItem, ListItemText, ListItemIcon, Chip, Button, Divider } from '@mui/material';
import { Compare as CompareIcon } from '@mui/icons-material';

interface TPVersion {
  id: string;
  version_no: number;
  title: string;
  learning_objectives: any;
  time_allocation: any;
  is_current_version: boolean;
  created_at: string;
}

interface TPVersionHistoryProps {
  tpSetId: string;
  sequenceNumber: number;
}

export const TPVersionHistory: React.FC<TPVersionHistoryProps> = ({ tpSetId, sequenceNumber }) => {
  const [history, setHistory] = React.useState<TPVersion[]>([]);
  const [loading, setLoading] = React.useState(true);

  React.useEffect(() => {
    // Fetch TP version history
    fetch(`/api/v1/learning-planning/tp-sets/${tpSetId}/tps/${sequenceNumber}/versions`)
      .then(res => res.json())
      .then(data => setHistory(data.data))
      .finally(() => setLoading(false));
  }, [tpSetId, sequenceNumber]);

  if (loading) return <Box>Loading...</Box>;

  return (
    <Box>
      <Box display="flex" justifyContent="space-between" alignItems="center" mb={2}>
        <Typography variant="h6">TP Version History</Typography>
        <Button
          variant="outlined"
          startIcon={<CompareIcon />}
          size="small"
        >
          Compare Versions
        </Button>
      </Box>
      <List>
        {history.map((version, index) => (
          <React.Fragment key={version.id}>
            <ListItem>
              <ListItemIcon>
                <CompareIcon color={version.is_current_version ? "primary" : "disabled"} />
              </ListItemIcon>
              <ListItemText
                primary={
                  <Box display="flex" alignItems="center" gap={1}>
                    <Typography variant="body1">Version {version.version_no}</Typography>
                    {version.is_current_version && (
                      <Chip label="Current" color="primary" size="small" />
                    )}
                  </Box>
                }
                secondary={
                  <>
                    <Typography variant="body2" color="textSecondary">
                      {version.title}
                    </Typography>
                    <Typography variant="caption" color="textSecondary">
                      Created: {new Date(version.created_at).toLocaleString()}
                    </Typography>
                  </>
                }
              />
            </ListItem>
            {index < history.length - 1 && <Divider />}
          </React.Fragment>
        ))}
      </List>
    </Box>
  );
};

/**
 * Revision History Component
 * Displays the revision history for evidence
 */

import React from 'react';
import { Box, Typography, List, ListItem, ListItemText, ListItemIcon, Divider, Chip } from '@mui/material';
import { History as HistoryIcon } from '@mui/icons-material';

interface Revision {
  id: string;
  version_no: number;
  changes: string;
  created_at: string;
  created_by: string;
}

interface RevisionHistoryProps {
  revisions: Revision[];
}

export const RevisionHistory = ({ revisions }: RevisionHistoryProps) => {
  if (revisions.length === 0) {
    return (
      <Box sx={{ p: 3, textAlign: 'center' }}>
        <Typography color="text.secondary">No revision history</Typography>
      </Box>
    );
  }

  return (
    <Box>
      <Typography variant="h6" gutterBottom>
        Revision History
      </Typography>
      <List>
        {revisions.map((revision, index) => (
          <React.Fragment key={revision.id}>
            <ListItem>
              <ListItemIcon>
                <HistoryIcon color="primary" />
              </ListItemIcon>
              <ListItemText
                primary={
                  <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
                    <Typography variant="body1">Version {revision.version_no}</Typography>
                    <Chip label={new Date(revision.created_at).toLocaleString()} size="small" variant="outlined" />
                  </Box>
                }
                secondary={
                  <>
                    <Typography variant="body2" color="textSecondary">
                      {revision.changes}
                    </Typography>
                    <Typography variant="caption" color="textSecondary">
                      By {revision.created_by}
                    </Typography>
                  </>
                }
              />
            </ListItem>
            {index < revisions.length - 1 && <Divider />}
          </React.Fragment>
        ))}
      </List>
    </Box>
  );
};

/**
 * Revision History Component
 * Displays the revision history for evidence
 */

import { Box, Typography, Timeline, TimelineItem, TimelineSeparator, TimelineConnector, TimelineContent, TimelineDot } from '@mui/material';
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
      <Timeline>
        {revisions.map((revision) => (
          <TimelineItem key={revision.id}>
            <TimelineSeparator>
              <TimelineDot color="primary" />
              <TimelineConnector />
            </TimelineSeparator>
            <TimelineContent>
              <Typography variant="body1">Version {revision.version_no}</Typography>
              <Typography variant="body2" color="textSecondary">
                {revision.changes}
              </Typography>
              <Typography variant="caption" color="textSecondary">
                By {revision.created_by} • {new Date(revision.created_at).toLocaleString()}
              </Typography>
            </TimelineContent>
          </TimelineItem>
        ))}
      </Timeline>
    </Box>
  );
};

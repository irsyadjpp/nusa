/**
 * Report List Component
 * Displays a list of narrative reports
 */

import { Box, List, ListItem, ListItemText, ListItemButton, Chip, Typography } from '@mui/material';
import { NarrativeReport } from '@/api/narrative-report';

interface ReportListProps {
  reports: NarrativeReport[];
  selectedId?: string;
  onSelect?: (report: NarrativeReport) => void;
  loading?: boolean;
}

export const ReportList = ({ reports, selectedId, onSelect, loading }: ReportListProps) => {
  if (loading) {
    return (
      <Box sx={{ p: 3, textAlign: 'center' }}>
        <Typography>Loading Reports...</Typography>
      </Box>
    );
  }

  if (reports.length === 0) {
    return (
      <Box sx={{ p: 3, textAlign: 'center' }}>
        <Typography color="text.secondary">No Reports found</Typography>
      </Box>
    );
  }

  return (
    <List>
      {reports.map((report) => (
        <ListItem
          key={report.id}
          disablePadding
          selected={selectedId === report.id}
          secondaryAction={
            <Chip
              label={report.status}
              size="small"
              color={report.status === 'published' ? 'success' : report.status === 'draft' ? 'default' : 'warning'}
            />
          }
        >
          <ListItemButton onClick={() => onSelect?.(report)} selected={selectedId === report.id}>
            <ListItemText
              primary={report.title}
              secondary={`Student: ${report.student_id} | Period: ${report.period}`}
            />
          </ListItemButton>
        </ListItem>
      ))}
    </List>
  );
};

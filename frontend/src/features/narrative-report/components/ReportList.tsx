/**
 * Report List Component
 * Displays a list of narrative reports
 */

import React from 'react';
import { Box, List, ListItem, ListItemText, ListItemButton, Chip, Typography, Divider } from '@mui/material';
import { NarrativeReport } from '@/shared/types/domain';

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
      {reports.map((report, index) => (
        <React.Fragment key={report.id}>
          <ListItem
            disablePadding
            sx={{
              backgroundColor: selectedId === report.id ? 'action.selected' : 'transparent',
            }}
          >
            <ListItemButton onClick={() => onSelect?.(report)}>
              <ListItemText
                primary={report.title || `Report ${report.id}`}
                secondary={`Student: ${report.student_id} | Period: ${report.period || report.period_id}`}
              />
            </ListItemButton>
            <Box sx={{ display: 'flex', alignItems: 'center', pr: 2 }}>
              <Chip
                label="Draft"
                size="small"
                color="default"
              />
            </Box>
          </ListItem>
          {index < reports.length - 1 && <Divider />}
        </React.Fragment>
      ))}
    </List>
  );
};

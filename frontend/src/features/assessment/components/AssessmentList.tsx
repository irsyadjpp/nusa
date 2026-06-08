/**
 * Assessment List Component
 * Displays a list of assessments with filtering and pagination
 */

import { Box, List, ListItem, ListItemText, ListItemButton, Chip, Typography } from '@mui/material';
import { Assessment } from '@/api/assessment';

interface AssessmentListProps {
  assessments: Assessment[];
  selectedId?: string;
  onSelect?: (assessment: Assessment) => void;
  loading?: boolean;
}

export const AssessmentList = ({ assessments, selectedId, onSelect, loading }: AssessmentListProps) => {
  if (loading) {
    return (
      <Box sx={{ p: 3, textAlign: 'center' }}>
        <Typography>Loading Assessments...</Typography>
      </Box>
    );
  }

  if (assessments.length === 0) {
    return (
      <Box sx={{ p: 3, textAlign: 'center' }}>
        <Typography color="text.secondary">No Assessments found</Typography>
      </Box>
    );
  }

  return (
    <List>
      {assessments.map((assessment) => (
        <ListItem
          key={assessment.id}
          disablePadding
          selected={selectedId === assessment.id}
          secondaryAction={
            <Chip
              label={assessment.status}
              size="small"
              color={assessment.status === 'approved' ? 'success' : assessment.status === 'draft' ? 'default' : 'warning'}
            />
          }
        >
          <ListItemButton onClick={() => onSelect?.(assessment)} selected={selectedId === assessment.id}>
            <ListItemText
              primary={`Assessment v${assessment.version_no}`}
              secondary={`Type: ${assessment.assessment_type} | TP: ${assessment.tp_id}`}
            />
          </ListItemButton>
        </ListItem>
      ))}
    </List>
  );
};

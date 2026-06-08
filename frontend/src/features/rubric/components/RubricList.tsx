/**
 * Rubric List Component
 * Displays a list of rubrics with filtering and pagination
 */

import { Box, List, ListItem, ListItemText, ListItemButton, Chip, Typography } from '@mui/material';
import { Rubric } from '@/api/rubric';

interface RubricListProps {
  rubrics: Rubric[];
  selectedId?: string;
  onSelect?: (rubric: Rubric) => void;
  loading?: boolean;
}

export const RubricList = ({ rubrics, selectedId, onSelect, loading }: RubricListProps) => {
  if (loading) {
    return (
      <Box sx={{ p: 3, textAlign: 'center' }}>
        <Typography>Loading Rubrics...</Typography>
      </Box>
    );
  }

  if (rubrics.length === 0) {
    return (
      <Box sx={{ p: 3, textAlign: 'center' }}>
        <Typography color="text.secondary">No Rubrics found</Typography>
      </Box>
    );
  }

  return (
    <List>
      {rubrics.map((rubric) => (
        <ListItem
          key={rubric.id}
          disablePadding
          selected={selectedId === rubric.id}
          secondaryAction={
            <Chip
              label={rubric.status}
              size="small"
              color={rubric.status === 'approved' ? 'success' : rubric.status === 'draft' ? 'default' : 'warning'}
            />
          }
        >
          <ListItemButton onClick={() => onSelect?.(rubric)} selected={selectedId === rubric.id}>
            <ListItemText
              primary={rubric.title}
              secondary={`Assessment: ${rubric.assessment_id}`}
            />
          </ListItemButton>
        </ListItem>
      ))}
    </List>
  );
};

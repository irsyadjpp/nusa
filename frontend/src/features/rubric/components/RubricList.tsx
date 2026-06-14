/**
 * Rubric List Component
 * Displays a list of rubrics with filtering and pagination
 */

import React from 'react';
import { Box, List, ListItem, ListItemText, ListItemButton, Chip, Typography, Divider } from '@mui/material';
import { Rubric } from '@/shared/types/domain';

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
      {rubrics.map((rubric, index) => (
        <React.Fragment key={rubric.id}>
          <ListItem
            disablePadding
            sx={{
              backgroundColor: selectedId === rubric.id ? 'action.selected' : 'transparent',
            }}
          >
            <ListItemButton onClick={() => onSelect?.(rubric)}>
              <ListItemText
                primary={rubric.title}
                secondary={`Assessment: ${rubric.assessment_id}`}
              />
            </ListItemButton>
            <Box sx={{ display: 'flex', alignItems: 'center', pr: 2 }}>
              <Chip
                label={rubric.status}
                size="small"
                color={rubric.status === 'APPROVED' ? 'success' : rubric.status === 'DRAFT' ? 'default' : 'warning'}
              />
            </Box>
          </ListItem>
          {index < rubrics.length - 1 && <Divider />}
        </React.Fragment>
      ))}
    </List>
  );
};

/**
 * TP List Component
 * Displays a list of Teaching Plans with filtering and pagination
 */

import React from 'react';
import { Box, List, ListItem, ListItemText, ListItemButton, Chip, Typography, Divider } from '@mui/material';
import { TP } from '@/api/tp';

interface TPListProps {
  tps: TP[];
  selectedId?: string;
  onSelect?: (tp: TP) => void;
  loading?: boolean;
}

export const TPList = ({ tps, selectedId, onSelect, loading }: TPListProps) => {
  if (loading) {
    return (
      <Box sx={{ p: 3, textAlign: 'center' }}>
        <Typography>Loading TPs...</Typography>
      </Box>
    );
  }

  if (tps.length === 0) {
    return (
      <Box sx={{ p: 3, textAlign: 'center' }}>
        <Typography color="text.secondary">No TPs found</Typography>
      </Box>
    );
  }

  return (
    <List>
      {tps.map((tp, index) => (
        <React.Fragment key={tp.id}>
          <ListItem
            disablePadding
            sx={{
              backgroundColor: selectedId === tp.id ? 'action.selected' : 'transparent',
            }}
          >
            <ListItemButton onClick={() => onSelect?.(tp)}>
              <ListItemText
                primary={tp.title}
                secondary={`Sequence: ${tp.sequence_number} | Phase: ${tp.phase_id}`}
              />
            </ListItemButton>
            <Box sx={{ display: 'flex', alignItems: 'center', pr: 2 }}>
              <Chip
                label={tp.status}
                size="small"
                color={tp.status === 'APPROVED' ? 'success' : tp.status === 'DRAFT' ? 'default' : 'warning'}
              />
            </Box>
          </ListItem>
          {index < tps.length - 1 && <Divider />}
        </React.Fragment>
      ))}
    </List>
  );
};

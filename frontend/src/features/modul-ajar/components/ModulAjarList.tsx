/**
 * Modul Ajar List Component
 * Displays a list of Modul Ajars with filtering and pagination
 */

import React from 'react';
import { Box, List, ListItem, ListItemText, ListItemButton, Chip, Typography, Divider } from '@mui/material';
import { ModulAjar } from '@/shared/types/domain';

interface ModulAjarListProps {
  modulAjars: ModulAjar[];
  selectedId?: string;
  onSelect?: (modulAjar: ModulAjar) => void;
  loading?: boolean;
}

export const ModulAjarList = ({ modulAjars, selectedId, onSelect, loading }: ModulAjarListProps) => {
  if (loading) {
    return (
      <Box sx={{ p: 3, textAlign: 'center' }}>
        <Typography>Loading Modul Ajars...</Typography>
      </Box>
    );
  }

  if (modulAjars.length === 0) {
    return (
      <Box sx={{ p: 3, textAlign: 'center' }}>
        <Typography color="text.secondary">No Modul Ajars found</Typography>
      </Box>
    );
  }

  return (
    <List>
      {modulAjars.map((modulAjar, index) => (
        <React.Fragment key={modulAjar.id}>
          <ListItem
            disablePadding
            sx={{
              backgroundColor: selectedId === modulAjar.id ? 'action.selected' : 'transparent',
            }}
          >
            <ListItemButton onClick={() => onSelect?.(modulAjar)}>
              <ListItemText
                primary={modulAjar.title}
                secondary={`Sequence: ${modulAjar.sequence_number}`}
              />
            </ListItemButton>
            <Box sx={{ display: 'flex', alignItems: 'center', pr: 2 }}>
              <Chip
                label={modulAjar.status}
                size="small"
                color={modulAjar.status === 'APPROVED' ? 'success' : modulAjar.status === 'DRAFT' ? 'default' : 'warning'}
              />
            </Box>
          </ListItem>
          {index < modulAjars.length - 1 && <Divider />}
        </React.Fragment>
      ))}
    </List>
  );
};

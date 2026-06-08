/**
 * Modul Ajar List Component
 * Displays a list of Modul Ajars with filtering and pagination
 */

import { Box, List, ListItem, ListItemText, ListItemButton, Chip, Typography } from '@mui/material';
import { ModulAjar } from '@/api/modul-ajar';

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
      {modulAjars.map((modulAjar) => (
        <ListItem
          key={modulAjar.id}
          disablePadding
          selected={selectedId === modulAjar.id}
          secondaryAction={
            <Chip
              label={modulAjar.status}
              size="small"
              color={modulAjar.status === 'approved' ? 'success' : modulAjar.status === 'draft' ? 'default' : 'warning'}
            />
          }
        >
          <ListItemButton onClick={() => onSelect?.(modulAjar)} selected={selectedId === modulAjar.id}>
            <ListItemText
              primary={modulAjar.title}
              secondary={`Sequence: ${modulAjar.sequence_number}`}
            />
          </ListItemButton>
        </ListItem>
      ))}
    </List>
  );
};

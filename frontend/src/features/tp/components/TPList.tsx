/**
 * TP List Component
 * Displays a list of Teaching Plans with filtering and pagination
 */

import { Box, List, ListItem, ListItemText, ListItemButton, Chip, Typography } from '@mui/material';
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
      {tps.map((tp) => (
        <ListItem
          key={tp.id}
          disablePadding
          selected={selectedId === tp.id}
          secondaryAction={
            <Chip
              label={tp.status}
              size="small"
              color={tp.status === 'approved' ? 'success' : tp.status === 'draft' ? 'default' : 'warning'}
            />
          }
        >
          <ListItemButton onClick={() => onSelect?.(tp)} selected={selectedId === tp.id}>
            <ListItemText
              primary={tp.title}
              secondary={`Sequence: ${tp.sequence_number} | Phase: ${tp.phase_id}`}
            />
          </ListItemButton>
        </ListItem>
      ))}
    </List>
  );
};

/**
 * ATP List Component
 * Displays a list of ATPs with filtering and pagination
 */

import { Box, List, ListItem, ListItemText, ListItemButton, Chip, Typography } from '@mui/material';
import { ATP } from '@/api/atp';

interface ATPListProps {
  atps: ATP[];
  selectedId?: string;
  onSelect?: (atp: ATP) => void;
  loading?: boolean;
}

export const ATPList = ({ atps, selectedId, onSelect, loading }: ATPListProps) => {
  if (loading) {
    return (
      <Box sx={{ p: 3, textAlign: 'center' }}>
        <Typography>Loading ATPs...</Typography>
      </Box>
    );
  }

  if (atps.length === 0) {
    return (
      <Box sx={{ p: 3, textAlign: 'center' }}>
        <Typography color="text.secondary">No ATPs found</Typography>
      </Box>
    );
  }

  return (
    <List>
      {atps.map((atp) => (
        <ListItem
          key={atp.id}
          disablePadding
          selected={selectedId === atp.id}
          secondaryAction={
            <Chip
              label={atp.status}
              size="small"
              color={atp.status === 'approved' ? 'success' : atp.status === 'draft' ? 'default' : 'warning'}
            />
          }
        >
          <ListItemButton onClick={() => onSelect?.(atp)} selected={selectedId === atp.id}>
            <ListItemText
              primary={`Week ${atp.week_number}`}
              secondary={`Sequence: ${atp.sequence_number} | Hours: ${atp.estimated_hours}`}
            />
          </ListItemButton>
        </ListItem>
      ))}
    </List>
  );
};

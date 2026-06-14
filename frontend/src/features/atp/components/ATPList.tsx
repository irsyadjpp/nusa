/**
 * ATP List Component
 * Displays a list of ATPs with filtering and pagination
 */

import { Box, List, ListItem, ListItemText, ListItemButton, Chip, Typography } from '@mui/material';
import { ATP } from '@/shared/types/domain';

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
          sx={{
            backgroundColor: selectedId === atp.id ? 'action.selected' : 'transparent',
          }}
        >
          <ListItemButton onClick={() => onSelect?.(atp)}>
            <ListItemText
              primary={`Week ${atp.week_number}`}
              secondary={`Sequence: ${atp.sequence_number} | Hours: ${atp.estimated_hours}`}
            />
          </ListItemButton>
          <Box sx={{ display: 'flex', alignItems: 'center', pr: 2 }}>
            <Chip
              label={atp.status}
              size="small"
              color={atp.status === 'APPROVED' ? 'success' : atp.status === 'DRAFT' ? 'default' : 'warning'}
            />
          </Box>
        </ListItem>
      ))}
    </List>
  );
};

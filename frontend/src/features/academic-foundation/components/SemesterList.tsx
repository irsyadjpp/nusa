/**
 * Semester List Component
 * Displays a list of semesters with type and status indicators
 */

import { Box, List, ListItem, ListItemText, ListItemButton, Chip, Typography } from '@mui/material';
import { Semester } from '@/api/academic-foundation';

interface SemesterListProps {
  semesters: Semester[];
  selectedId?: string;
  onSelect?: (semester: Semester) => void;
  loading?: boolean;
}

export const SemesterList = ({ semesters, selectedId, onSelect, loading }: SemesterListProps) => {
  if (loading) {
    return (
      <Box sx={{ p: 3, textAlign: 'center' }}>
        <Typography>Loading semesters...</Typography>
      </Box>
    );
  }

  if (semesters.length === 0) {
    return (
      <Box sx={{ p: 3, textAlign: 'center' }}>
        <Typography color="text.secondary">No semesters found</Typography>
      </Box>
    );
  }

  return (
    <List>
      {semesters.map((semester) => (
        <ListItem
          key={semester.id}
          disablePadding
          sx={{
            backgroundColor: selectedId === semester.id ? 'action.selected' : 'transparent',
          }}
        >
          <ListItemButton onClick={() => onSelect?.(semester)}>
            <ListItemText
              primary={semester.name}
              secondary={`${new Date(semester.start_date).toLocaleDateString('id-ID')} - ${new Date(semester.end_date).toLocaleDateString('id-ID')} | Semester: ${semester.sequence_number}`}
            />
          </ListItemButton>
          <Box sx={{ display: 'flex', alignItems: 'center', gap: 1, pr: 2 }}>
            <Chip
              label={semester.type}
              size="small"
              color={semester.type === 'GANJIL' ? 'primary' : 'secondary'}
            />
            <Chip
              label={semester.status}
              size="small"
              color={semester.status === 'ACTIVE' ? 'success' : 'default'}
            />
          </Box>
        </ListItem>
      ))}
    </List>
  );
};
/**
 * Academic Year List Component
 * Displays a list of academic years with status indicators
 */

import { Box, List, ListItem, ListItemText, ListItemButton, Chip, Typography } from '@mui/material';
import { AcademicYear } from '@/api/academic-foundation';

interface AcademicYearListProps {
  academicYears: AcademicYear[];
  selectedId?: string;
  onSelect?: (academicYear: AcademicYear) => void;
  loading?: boolean;
}

export const AcademicYearList = ({ academicYears, selectedId, onSelect, loading }: AcademicYearListProps) => {
  if (loading) {
    return (
      <Box sx={{ p: 3, textAlign: 'center' }}>
        <Typography>Loading academic years...</Typography>
      </Box>
    );
  }

  if (academicYears.length === 0) {
    return (
      <Box sx={{ p: 3, textAlign: 'center' }}>
        <Typography color="text.secondary">No academic years found</Typography>
      </Box>
    );
  }

  return (
    <List>
      {academicYears.map((academicYear) => (
        <ListItem
          key={academicYear.id}
          disablePadding
          sx={{
            backgroundColor: selectedId === academicYear.id ? 'action.selected' : 'transparent',
          }}
        >
          <ListItemButton onClick={() => onSelect?.(academicYear)}>
            <ListItemText
              primary={academicYear.name}
              secondary={`${new Date(academicYear.start_date).toLocaleDateString('id-ID')} - ${new Date(academicYear.end_date).toLocaleDateString('id-ID')}`}
            />
          </ListItemButton>
          <Box sx={{ display: 'flex', alignItems: 'center', pr: 2 }}>
            <Chip
              label={academicYear.status}
              size="small"
              color={
                academicYear.status === 'ACTIVE' 
                  ? 'success' 
                  : academicYear.status === 'ARCHIVED' 
                    ? 'default' 
                    : 'warning'
              }
            />
          </Box>
        </ListItem>
      ))}
    </List>
  );
};
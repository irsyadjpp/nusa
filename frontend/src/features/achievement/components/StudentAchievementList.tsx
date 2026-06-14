/**
 * Student Achievement List Component
 * Displays a list of students with their achievement data
 */

import { Box, List, ListItem, ListItemText, ListItemButton, Typography, Chip } from '@mui/material';

interface StudentAchievement {
  id: string;
  name: string;
  averageScore: number;
  completedAssessments: number;
  totalAssessments: number;
}

interface StudentAchievementListProps {
  students: StudentAchievement[];
  selectedId?: string;
  onSelect?: (student: StudentAchievement) => void;
  loading?: boolean;
}

export const StudentAchievementList = ({ students, selectedId, onSelect, loading }: StudentAchievementListProps) => {
  if (loading) {
    return (
      <Box sx={{ p: 3, textAlign: 'center' }}>
        <Typography>Loading Students...</Typography>
      </Box>
    );
  }

  if (students.length === 0) {
    return (
      <Box sx={{ p: 3, textAlign: 'center' }}>
        <Typography color="text.secondary">No Students found</Typography>
      </Box>
    );
  }

  return (
    <List>
      {students.map((student) => (
        <ListItem
          key={student.id}
          disablePadding
          sx={{
            backgroundColor: selectedId === student.id ? 'action.selected' : 'transparent',
          }}
        >
          <ListItemButton onClick={() => onSelect?.(student)}>
            <ListItemText
              primary={student.name}
              secondary={`${student.completedAssessments}/${student.totalAssessments} assessments completed`}
            />
          </ListItemButton>
          <Box sx={{ display: 'flex', alignItems: 'center', pr: 2 }}>
            <Chip
              label={`${student.averageScore}%`}
              size="small"
              color={student.averageScore >= 70 ? 'success' : student.averageScore >= 50 ? 'warning' : 'error'}
            />
          </Box>
        </ListItem>
      ))}
    </List>
  );
};

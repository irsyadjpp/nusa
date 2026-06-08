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
          selected={selectedId === student.id}
          secondaryAction={
            <Chip
              label={`${student.averageScore}%`}
              size="small"
              color={student.averageScore >= 70 ? 'success' : student.averageScore >= 50 ? 'warning' : 'error'}
            />
          }
        >
          <ListItemButton onClick={() => onSelect?.(student)} selected={selectedId === student.id}>
            <ListItemText
              primary={student.name}
              secondary={`${student.completedAssessments}/${student.totalAssessments} assessments completed`}
            />
          </ListItemButton>
        </ListItem>
      ))}
    </List>
  );
};

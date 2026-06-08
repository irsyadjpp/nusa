/**
 * Student Detail Panel Component
 * Displays detailed information about a student's achievements
 */

import { Box, Typography, Paper, Divider, Stack } from '@mui/material';

interface StudentDetailPanelProps {
  student: {
    id: string;
    name: string;
    averageScore: number;
    completedAssessments: number;
    totalAssessments: number;
    recentAchievements: string[];
  };
}

export const StudentDetailPanel = ({ student }: StudentDetailPanelProps) => {
  return (
    <Paper elevation={2} sx={{ p: 3 }}>
      <Typography variant="h5" gutterBottom>
        {student.name}
      </Typography>
      <Divider sx={{ my: 2 }} />
      <Stack spacing={2}>
        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            Average Score
          </Typography>
          <Typography variant="h3">{student.averageScore}%</Typography>
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            Assessments Completed
          </Typography>
          <Typography variant="body1">
            {student.completedAssessments} / {student.totalAssessments}
          </Typography>
        </Box>

        <Divider />

        <Box>
          <Typography variant="subtitle2" color="text.secondary" gutterBottom>
            Recent Achievements
          </Typography>
          {student.recentAchievements.length > 0 ? (
            <Stack spacing={1}>
              {student.recentAchievements.map((achievement, index) => (
                <Typography key={index} variant="body2">
                  • {achievement}
                </Typography>
              ))}
            </Stack>
          ) : (
            <Typography variant="body2" color="text.secondary">
              No recent achievements
            </Typography>
          )}
        </Box>
      </Stack>
    </Paper>
  );
};

/**
 * Competency Progress Component
 * Displays progress across different competencies
 */

import { Box, Typography, Paper, Divider, LinearProgress } from '@mui/material';

interface Competency {
  name: string;
  progress: number;
  level: string;
}

interface CompetencyProgressProps {
  competencies: Competency[];
}

export const CompetencyProgress = ({ competencies }: CompetencyProgressProps) => {
  return (
    <Paper elevation={2} sx={{ p: 3 }}>
      <Typography variant="h5" gutterBottom>
        Competency Progress
      </Typography>
      <Divider sx={{ my: 2 }} />
      <Stack spacing={3}>
        {competencies.map((competency) => (
          <Box key={competency.name}>
            <Box sx={{ display: 'flex', justifyContent: 'space-between', mb: 1 }}>
              <Typography variant="body1">{competency.name}</Typography>
              <Typography variant="body2" color="text.secondary">
                {competency.level}
              </Typography>
            </Box>
            <LinearProgress
              variant="determinate"
              value={competency.progress}
              sx={{ height: 10, borderRadius: 5 }}
              color={competency.progress >= 70 ? 'success' : competency.progress >= 50 ? 'warning' : 'error'}
            />
            <Typography variant="caption" color="text.secondary" sx={{ mt: 0.5, display: 'block' }}>
              {competency.progress}% Complete
            </Typography>
          </Box>
        ))}
      </Stack>
    </Paper>
  );
};

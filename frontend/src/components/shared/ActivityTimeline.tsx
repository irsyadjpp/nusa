/**
 * ActivityTimeline Component
 * Displays a timeline of activities
 */

import { Typography, Paper, Divider, Stack, Chip, Box } from '@mui/material';

interface Activity {
  id: string;
  title: string;
  description?: string;
  timestamp: string;
  type?: 'info' | 'success' | 'warning' | 'error';
}

interface ActivityTimelineProps {
  activities: Activity[];
  showDescription?: boolean;
}

export const ActivityTimeline = ({ activities, showDescription = true }: ActivityTimelineProps) => {
  const getChipColor = (type?: string): 'default' | 'primary' | 'secondary' | 'success' | 'error' | 'info' | 'warning' => {
    switch (type) {
      case 'success':
        return 'success';
      case 'warning':
        return 'warning';
      case 'error':
        return 'error';
      default:
        return 'info';
    }
  };

  return (
    <Paper elevation={2} sx={{ p: 3 }}>
      <Typography variant="h5" gutterBottom>
        Activity Timeline
      </Typography>
      <Divider sx={{ my: 2 }} />
      <Stack spacing={2}>
        {activities.map((activity) => (
          <Box key={activity.id}>
            <Box sx={{ display: 'flex', alignItems: 'center', gap: 2, mb: 1 }}>
              <Chip
                label={activity.type || 'info'}
                size="small"
                color={getChipColor(activity.type)}
              />
              <Typography variant="caption" color="textSecondary">
                {new Date(activity.timestamp).toLocaleString()}
              </Typography>
            </Box>
            <Typography variant="body1" fontWeight="medium">
              {activity.title}
            </Typography>
            {showDescription && activity.description && (
              <Typography variant="body2" color="textSecondary" sx={{ mt: 0.5 }}>
                {activity.description}
              </Typography>
            )}
          </Box>
        ))}
      </Stack>
    </Paper>
  );
};

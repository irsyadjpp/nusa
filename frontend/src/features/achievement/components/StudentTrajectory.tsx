/**
 * Student Trajectory Component
 * Displays the learning trajectory of a student over time
 */

import { Box, Typography, Paper, Divider, Timeline, TimelineItem, TimelineSeparator, TimelineConnector, TimelineContent, TimelineDot } from '@mui/material';

interface TrajectoryPoint {
  date: string;
  milestone: string;
  score: number;
}

interface StudentTrajectoryProps {
  trajectory: TrajectoryPoint[];
}

export const StudentTrajectory = ({ trajectory }: StudentTrajectoryProps) => {
  if (trajectory.length === 0) {
    return (
      <Paper elevation={2} sx={{ p: 3 }}>
        <Typography variant="h5" gutterBottom>
          Learning Trajectory
        </Typography>
        <Divider sx={{ my: 2 }} />
        <Typography color="text.secondary">No trajectory data available</Typography>
      </Paper>
    );
  }

  return (
    <Paper elevation={2} sx={{ p: 3 }}>
      <Typography variant="h5" gutterBottom>
        Learning Trajectory
      </Typography>
      <Divider sx={{ my: 2 }} />
      <Timeline>
        {trajectory.map((point, index) => (
          <TimelineItem key={index}>
            <TimelineSeparator>
              <TimelineDot color={point.score >= 70 ? 'success' : point.score >= 50 ? 'warning' : 'error'} />
              {index < trajectory.length - 1 && <TimelineConnector />}
            </TimelineSeparator>
            <TimelineContent>
              <Typography variant="body1">{point.milestone}</Typography>
              <Typography variant="body2" color="textSecondary">
                Score: {point.score}%
              </Typography>
              <Typography variant="caption" color="textSecondary">
                {new Date(point.date).toLocaleString()}
              </Typography>
            </TimelineContent>
          </TimelineItem>
        ))}
      </Timeline>
    </Paper>
  );
};

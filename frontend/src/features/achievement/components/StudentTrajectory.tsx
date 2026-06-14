/**
 * Student Trajectory Component
 * Displays the learning trajectory of a student over time
 */

import React from 'react';
import { Typography, Paper, Divider, List, ListItem, ListItemText, ListItemIcon } from '@mui/material';
import { CheckCircle, Warning, Error } from '@mui/icons-material';

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
      <List>
        {trajectory.map((point, index) => (
          <React.Fragment key={index}>
            <ListItem>
              <ListItemIcon>
                {point.score >= 70 ? (
                  <CheckCircle color="success" />
                ) : point.score >= 50 ? (
                  <Warning color="warning" />
                ) : (
                  <Error color="error" />
                )}
              </ListItemIcon>
              <ListItemText
                primary={point.milestone}
                secondary={
                  <>
                    <Typography variant="body2" color="textSecondary">
                      Score: {point.score}%
                    </Typography>
                    <Typography variant="caption" color="textSecondary">
                      {new Date(point.date).toLocaleString()}
                    </Typography>
                  </>
                }
              />
            </ListItem>
            {index < trajectory.length - 1 && <Divider />}
          </React.Fragment>
        ))}
      </List>
    </Paper>
  );
};

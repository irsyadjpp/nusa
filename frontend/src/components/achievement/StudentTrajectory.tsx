/**
 * Student Trajectory Component
 * Chart showing progress over time
 */

import React from 'react';
import {
  Box,
  Typography,
  Card,
  CardContent,
} from '@mui/material';
import { LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, Legend, ResponsiveContainer } from '@mui/x-charts';

interface TrajectoryPoint {
  date: string;
  score: number;
  mastery_level: string;
}

interface StudentTrajectoryProps {
  student_id: string;
  student_name: string;
  trajectory_points: TrajectoryPoint[];
}

const StudentTrajectory: React.FC<StudentTrajectoryProps> = ({
  student_id,
  student_name,
  trajectory_points,
}) => {
  const getMasteryColor = (level: string): string => {
    switch (level) {
      case 'EXCELLENT':
        return '#4caf50';
      case 'PROFICIENT':
        return '#2196f3';
      case 'DEVELOPING':
        return '#ff9800';
      case 'BEGINNING':
        return '#f44336';
      default:
        return '#9e9e9e';
    }
  };

  const chartData = trajectory_points.map((point) => ({
    x: new Date(point.date).toLocaleDateString('id-ID', { month: 'short', day: 'numeric' }),
    y: point.score,
    mastery: point.mastery_level,
  }));

  return (
    <Card>
      <CardContent>
        <Typography variant="h6" gutterBottom>
          Trajektori Kemajuan - {student_name}
        </Typography>

        <ResponsiveContainer width="100%" height={300}>
          <LineChart data={chartData}>
            <CartesianGrid strokeDasharray="3 3" />
            <XAxis dataKey="x" />
            <YAxis domain={[0, 100]} />
            <Tooltip />
            <Legend />
            <Line
              type="monotone"
              dataKey="y"
              stroke="#2196f3"
              strokeWidth={2}
              dot={{ fill: '#2196f3', strokeWidth: 2, r: 4 }}
              activeDot={{ r: 6 }}
            />
          </LineChart>
        </ResponsiveContainer>

        <Box sx={{ mt: 2 }}>
          <Typography variant="subtitle2" gutterBottom>
            Riwayat Pencapaian
          </Typography>
          {trajectory_points.map((point, index) => (
            <Box
              key={index}
              sx={{
                display: 'flex',
                alignItems: 'center',
                gap: 2,
                p: 1,
                bgcolor: index === trajectory_points.length - 1 ? 'primary.50' : 'background.paper',
                borderRadius: 1,
                mb: 1,
              }}
            >
              <Box
                sx={{
                  width: 12,
                  height: 12,
                  borderRadius: '50%',
                  bgcolor: getMasteryColor(point.mastery_level),
                }}
              />
              <Typography variant="body2" sx={{ flexGrow: 1 }}>
                {new Date(point.date).toLocaleDateString('id-ID', {
                  weekday: 'long',
                  year: 'numeric',
                  month: 'long',
                  day: 'numeric',
                })}
              </Typography>
              <Typography variant="body2" fontWeight="bold">
                {point.score}%
              </Typography>
              <Typography variant="caption" color="text.secondary">
                ({point.mastery_level})
              </Typography>
            </Box>
          ))}
        </Box>
      </CardContent>
    </Card>
  );
};

export default StudentTrajectory;

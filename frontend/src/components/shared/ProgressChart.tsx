/**
 * ProgressChart Component
 * Displays a progress chart with multiple data points
 */

import { Box, Typography, LinearProgress, Stack } from '@mui/material';

interface ProgressData {
  label: string;
  value: number;
  color?: 'primary' | 'secondary' | 'success' | 'error' | 'info' | 'warning';
}

interface ProgressChartProps {
  data: ProgressData[];
  showLabels?: boolean;
  showPercentage?: boolean;
}

export const ProgressChart = ({ data, showLabels = true, showPercentage = true }: ProgressChartProps) => {
  return (
    <Box>
      <Stack spacing={2}>
        {data.map((item, index) => (
          <Box key={index}>
            {showLabels && (
              <Box sx={{ display: 'flex', justifyContent: 'space-between', mb: 0.5 }}>
                <Typography variant="body2">{item.label}</Typography>
                {showPercentage && (
                  <Typography variant="body2" color="text.secondary">
                    {item.value}%
                  </Typography>
                )}
              </Box>
            )}
            <LinearProgress
              variant="determinate"
              value={item.value}
              sx={{ height: 8, borderRadius: 4 }}
              color={item.color || 'primary'}
            />
          </Box>
        ))}
      </Stack>
    </Box>
  );
};

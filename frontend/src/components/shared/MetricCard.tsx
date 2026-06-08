/**
 * MetricCard Component
 * Displays a metric with icon, value, and label
 */

import { Card, CardContent, Box, Typography } from '@mui/material';

interface MetricCardProps {
  title: string;
  value: string | number;
  icon?: React.ReactNode;
  color?: 'primary' | 'secondary' | 'success' | 'error' | 'info' | 'warning';
  trend?: {
    value: number;
    isPositive: boolean;
  };
}

export const MetricCard = ({ title, value, icon, color = 'primary', trend }: MetricCardProps) => {
  return (
    <Card>
      <CardContent>
        <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
          {icon && (
            <Box sx={{ color: `${color}.main` }}>
              {icon}
            </Box>
          )}
          <Box sx={{ flex: 1 }}>
            <Typography variant="h4" component="div">
              {value}
            </Typography>
            <Typography variant="body2" color="text.secondary">
              {title}
            </Typography>
            {trend && (
              <Typography
                variant="caption"
                color={trend.isPositive ? 'success.main' : 'error.main'}
                sx={{ display: 'block', mt: 0.5 }}
              >
                {trend.isPositive ? '+' : ''}{trend.value}%
              </Typography>
            )}
          </Box>
        </Box>
      </CardContent>
    </Card>
  );
};

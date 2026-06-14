/**
 * Dashboard Recent Activity Component
 * Shows recent activity feed
 */

import { Box, Card, CardContent, Typography, Stack } from "@mui/material";

export default function DashboardRecentActivity() {
  const recentActivity = [
    { action: 'TP approved', item: 'Mathematics Grade 5', time: '2 hours ago' },
    { action: 'Assessment created', item: 'Science Quiz Chapter 3', time: '4 hours ago' },
    { action: 'Evidence submitted', item: 'Student John Doe - Project', time: '6 hours ago' },
    { action: 'Report published', item: 'Narrative Report - Class 5A', time: '1 day ago' },
  ];

  return (
    <Card sx={{ height: '100%' }}>
      <CardContent>
        <Typography variant="h6" gutterBottom>
          Recent Activity
        </Typography>
        <Stack spacing={2}>
          {recentActivity.map((activity, index) => (
            <Box key={index}>
              <Typography variant="body1" fontWeight="medium">
                {activity.action}
              </Typography>
              <Typography variant="body2" color="text.secondary">
                {activity.item}
              </Typography>
              <Typography variant="caption" color="text.secondary">
                {activity.time}
              </Typography>
              {index < recentActivity.length - 1 && (
                <Box sx={{ my: 1, borderBottom: '1px solid #e0e0e0' }} />
              )}
            </Box>
          ))}
        </Stack>
      </CardContent>
    </Card>
  );
}

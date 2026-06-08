import React from 'react';
import { Box, Grid, Card, CardContent, Typography, Button, Stack } from '@mui/material';
import { School, Assignment, TrendingUp, People, Add, ArrowForward } from '@mui/icons-material';

const DashboardPage: React.FC = () => {
  const metrics = [
    { title: 'Total Students', value: '1,234', icon: People, color: 'primary' },
    { title: 'Active Classes', value: '45', icon: School, color: 'info' },
    { title: 'Pending Approvals', value: '23', icon: Assignment, color: 'warning' },
    { title: 'Average Mastery', value: '78%', icon: TrendingUp, color: 'success' },
  ];

  const quickActions = [
    { title: 'Create Teaching Plan', description: 'Start a new teaching plan', path: '/app/tp' },
    { title: 'Create Assessment', description: 'Design a new assessment', path: '/app/assessment' },
    { title: 'Review Evidence', description: 'Review student evidence submissions', path: '/app/evidence' },
    { title: 'Generate Reports', description: 'Create narrative reports', path: '/app/narrative-report' },
  ];

  const recentActivity = [
    { action: 'TP approved', item: 'Mathematics Grade 5', time: '2 hours ago' },
    { action: 'Assessment created', item: 'Science Quiz Chapter 3', time: '4 hours ago' },
    { action: 'Evidence submitted', item: 'Student John Doe - Project', time: '6 hours ago' },
    { action: 'Report published', item: 'Narrative Report - Class 5A', time: '1 day ago' },
  ];

  return (
    <Box sx={{ p: 3 }}>
      <Typography variant="h4" component="h1" gutterBottom>
        Dashboard
      </Typography>
      <Typography variant="body1" color="text.secondary" sx={{ mb: 4 }}>
        Welcome to NUSA Education Operating System
      </Typography>

      {/* Metrics Cards */}
      <Grid container spacing={3} sx={{ mb: 4 }}>
        {metrics.map((metric, index) => (
          <Grid size={{ xs: 12, sm: 6, md: 3 }} key={index}>
            <Card>
              <CardContent>
                <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
                  <metric.icon color={metric.color as any} sx={{ fontSize: 40 }} />
                  <Box>
                    <Typography variant="h4">{metric.value}</Typography>
                    <Typography variant="body2" color="text.secondary">
                      {metric.title}
                    </Typography>
                  </Box>
                </Box>
              </CardContent>
            </Card>
          </Grid>
        ))}
      </Grid>

      <Grid container spacing={3}>
        {/* Quick Actions */}
        <Grid size={{ xs: 12, md: 6 }}>
          <Card sx={{ height: '100%' }}>
            <CardContent>
              <Typography variant="h6" gutterBottom>
                Quick Actions
              </Typography>
              <Stack spacing={2}>
                {quickActions.map((action, index) => (
                  <Button
                    key={index}
                    variant="outlined"
                    fullWidth
                    startIcon={<Add />}
                    endIcon={<ArrowForward />}
                    sx={{ justifyContent: 'space-between' }}
                  >
                    <Box sx={{ textAlign: 'left' }}>
                      <Typography variant="body1">{action.title}</Typography>
                      <Typography variant="caption" color="text.secondary">
                        {action.description}
                      </Typography>
                    </Box>
                  </Button>
                ))}
              </Stack>
            </CardContent>
          </Card>
        </Grid>

        {/* Recent Activity */}
        <Grid size={{ xs: 12, md: 6 }}>
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
        </Grid>
      </Grid>
    </Box>
  );
};

export default DashboardPage;

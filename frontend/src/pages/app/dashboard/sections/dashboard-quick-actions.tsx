/**
 * Dashboard Quick Actions Component
 * Quick action buttons for common tasks
 */

import { Box, Card, CardContent, Typography, Button, Stack } from "@mui/material";
import AddIcon from "@mui/icons-material/Add";
import ArrowForwardIcon from "@mui/icons-material/ArrowForward";
import { Link as RouterLink } from "react-router-dom";

export default function DashboardQuickActions() {
  const quickActions = [
    { title: 'Create Teaching Plan', description: 'Start a new teaching plan', path: '/tp' },
    { title: 'Create Assessment', description: 'Design a new assessment', path: '/assessment' },
    { title: 'Review Evidence', description: 'Review student evidence submissions', path: '/evidence' },
    { title: 'Generate Reports', description: 'Create narrative reports', path: '/reports' },
  ];

  return (
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
              startIcon={<AddIcon />}
              endIcon={<ArrowForwardIcon />}
              component={RouterLink}
              to={action.path}
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
  );
}

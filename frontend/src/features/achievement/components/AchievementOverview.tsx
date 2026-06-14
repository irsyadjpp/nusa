/**
 * Achievement Overview Component
 * Displays overview statistics for achievements
 */

import { Box, Card, CardContent, Typography } from '@mui/material';
import { TrendingUp, People, Assignment, CheckCircle } from '@mui/icons-material';

interface AchievementOverviewProps {
  stats: {
    totalStudents: number;
    totalAssessments: number;
    averageScore: number;
    completionRate: number;
  };
}

export const AchievementOverview = ({ stats }: AchievementOverviewProps) => {
  return (
    <Box sx={{ mb: 3 }}>
      <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 3 }}>
        <Box sx={{ width: { xs: '100%', sm: '50%', md: '25%' } }}>
          <Card>
            <CardContent>
              <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
                <People color="primary" sx={{ fontSize: 40 }} />
                <Box>
                  <Typography variant="h4">{stats.totalStudents}</Typography>
                  <Typography variant="body2" color="text.secondary">
                    Total Students
                  </Typography>
                </Box>
              </Box>
            </CardContent>
          </Card>
        </Box>

        <Box sx={{ width: { xs: '100%', sm: '50%', md: '25%' } }}>
          <Card>
            <CardContent>
              <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
                <Assignment color="info" sx={{ fontSize: 40 }} />
                <Box>
                  <Typography variant="h4">{stats.totalAssessments}</Typography>
                  <Typography variant="body2" color="text.secondary">
                    Total Assessments
                  </Typography>
                </Box>
              </Box>
            </CardContent>
          </Card>
        </Box>

        <Box sx={{ width: { xs: '100%', sm: '50%', md: '25%' } }}>
          <Card>
            <CardContent>
              <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
                <TrendingUp color="success" sx={{ fontSize: 40 }} />
                <Box>
                  <Typography variant="h4">{stats.averageScore}%</Typography>
                  <Typography variant="body2" color="text.secondary">
                    Average Score
                  </Typography>
                </Box>
              </Box>
            </CardContent>
          </Card>
        </Box>

        <Box sx={{ width: { xs: '100%', sm: '50%', md: '25%' } }}>
          <Card>
            <CardContent>
              <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
                <CheckCircle color="warning" sx={{ fontSize: 40 }} />
                <Box>
                  <Typography variant="h4">{stats.completionRate}%</Typography>
                  <Typography variant="body2" color="text.secondary">
                    Completion Rate
                  </Typography>
                </Box>
              </Box>
            </CardContent>
          </Card>
        </Box>
      </Box>
    </Box>
  );
};

import React from 'react';
import { Box, Typography, Card, CardContent, Button, Chip, CircularProgress } from '@mui/material';
import { Refresh as RefreshIcon } from '@mui/icons-material';

interface AchievementData {
  competency_achievements: Array<{
    competency_name: string;
    achievement_level: string;
    progress_percentage: number;
  }>;
  overall_achievement: {
    total_score: number;
    max_score: number;
    percentage: number;
    level: string;
  };
}

interface ReportAchievementDisplayProps {
  reportId: string;
  achievementData?: AchievementData;
  lastCalculatedAt?: string;
  onRefresh?: () => void;
}

export const ReportAchievementDisplay: React.FC<ReportAchievementDisplayProps> = ({
  reportId,
  achievementData,
  lastCalculatedAt,
  onRefresh,
}) => {
  const [refreshing, setRefreshing] = React.useState(false);

  const handleRefresh = async () => {
    setRefreshing(true);
    try {
      await fetch(`/api/v1/reporting/narrative-reports/${reportId}/refresh-achievement`, {
        method: 'POST',
      });
      if (onRefresh) onRefresh();
    } catch (error) {
      console.error('Failed to refresh achievement data:', error);
    } finally {
      setRefreshing(false);
    }
  };

  if (!achievementData) {
    return (
      <Card>
        <CardContent>
          <Box display="flex" justifyContent="space-between" alignItems="center">
            <Typography variant="h6">Achievement Data</Typography>
            <Button
              variant="outlined"
              startIcon={refreshing ? <CircularProgress size={20} /> : <RefreshIcon />}
              onClick={handleRefresh}
              disabled={refreshing}
            >
              {refreshing ? 'Refreshing...' : 'Calculate Achievement'}
            </Button>
          </Box>
          <Typography variant="body2" color="textSecondary" sx={{ mt: 2 }}>
            No achievement data available. Click "Calculate Achievement" to generate.
          </Typography>
        </CardContent>
      </Card>
    );
  }

  return (
    <Card>
      <CardContent>
        <Box display="flex" justifyContent="space-between" alignItems="center" mb={2}>
          <Typography variant="h6">Achievement Data</Typography>
          <Box display="flex" alignItems="center" gap={2}>
            {lastCalculatedAt && (
              <Typography variant="caption" color="textSecondary">
                Last updated: {new Date(lastCalculatedAt).toLocaleString()}
              </Typography>
            )}
            <Button
              variant="outlined"
              size="small"
              startIcon={refreshing ? <CircularProgress size={16} /> : <RefreshIcon />}
              onClick={handleRefresh}
              disabled={refreshing}
            >
              Refresh
            </Button>
          </Box>
        </Box>

        <Box mb={3}>
          <Typography variant="subtitle1" gutterBottom>Overall Achievement</Typography>
          <Box display="flex" alignItems="center" gap={2}>
            <Typography variant="h4">
              {achievementData.overall_achievement.percentage.toFixed(1)}%
            </Typography>
            <Chip label={achievementData.overall_achievement.level} color="primary" />
          </Box>
          <Typography variant="body2" color="textSecondary">
            Score: {achievementData.overall_achievement.total_score} / {achievementData.overall_achievement.max_score}
          </Typography>
        </Box>

        <Box>
          <Typography variant="subtitle1" gutterBottom>Competency Breakdown</Typography>
          {achievementData.competency_achievements.map((competency, index) => (
            <Box key={index} mb={1}>
              <Box display="flex" justifyContent="space-between" alignItems="center">
                <Typography variant="body2">{competency.competency_name}</Typography>
                <Box display="flex" alignItems="center" gap={1}>
                  <Typography variant="body2">{competency.progress_percentage.toFixed(1)}%</Typography>
                  <Chip label={competency.achievement_level} size="small" />
                </Box>
              </Box>
            </Box>
          ))}
        </Box>
      </CardContent>
    </Card>
  );
};

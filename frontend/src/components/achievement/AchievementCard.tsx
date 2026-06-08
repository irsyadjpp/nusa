/**
 * Achievement Card Component
 * Card displaying student achievement
 */

import React from 'react';
import {
  Box,
  Card,
  CardContent,
  Typography,
  Chip,
  LinearProgress,
  Avatar,
  Grid,
} from '@mui/material';
import {
  Star,
  TrendingUp,
  CheckCircle,
} from '@mui/icons-material';

interface StudentAchievement {
  student_id: string;
  student_name: string;
  tp_id: string;
  tp_title: string;
  competency_id: string;
  competency_name: string;
  mastery_level: string;
  score: number;
  max_score: number;
  percentage: number;
  achieved_criteria: string[];
  pending_criteria: string[];
  last_updated: string;
}

interface AchievementCardProps {
  achievement: StudentAchievement;
  onClick?: () => void;
}

const AchievementCard: React.FC<AchievementCardProps> = ({
  achievement,
  onClick,
}) => {
  const getMasteryColor = (level: string): 'success' | 'info' | 'warning' | 'error' => {
    switch (level) {
      case 'EXCELLENT':
        return 'success';
      case 'PROFICIENT':
        return 'info';
      case 'DEVELOPING':
        return 'warning';
      case 'BEGINNING':
        return 'error';
      default:
        return 'info';
    }
  };

  const getMasteryLabel = (level: string): string => {
    switch (level) {
      case 'EXCELLENT':
        return 'Sangat Baik';
      case 'PROFICIENT':
        return 'Baik';
      case 'DEVELOPING':
        return 'Sedang Berkembang';
      case 'BEGINNING':
        return 'Perlu Bimbingan';
      default:
        return level;
    }
  };

  return (
    <Card onClick={onClick} sx={{ cursor: onClick ? 'pointer' : 'default' }}>
      <CardContent>
        <Box sx={{ display: 'flex', alignItems: 'center', gap: 2, mb: 2 }}>
          <Avatar>{achievement.student_name.charAt(0)}</Avatar>
          <Box sx={{ flexGrow: 1 }}>
            <Typography variant="h6">{achievement.student_name}</Typography>
            <Typography variant="body2" color="text.secondary">
              {achievement.competency_name}
            </Typography>
          </Box>
          <Chip
            label={getMasteryLabel(achievement.mastery_level)}
            color={getMasteryColor(achievement.mastery_level)}
            icon={<Star />}
          />
        </Box>

        <Grid container spacing={2} sx={{ mb: 2 }}>
          <Grid item xs={6}>
            <Typography variant="caption" color="text.secondary">
              Skor
            </Typography>
            <Typography variant="body1" fontWeight="bold">
              {achievement.score}/{achievement.max_score}
            </Typography>
          </Grid>
          <Grid item xs={6}>
            <Typography variant="caption" color="text.secondary">
              Persentase
            </Typography>
            <Typography variant="body1" fontWeight="bold">
              {achievement.percentage}%
            </Typography>
          </Grid>
        </Grid>

        <Box sx={{ mb: 2 }}>
          <LinearProgress
            variant="determinate"
            value={achievement.percentage}
            color={getMasteryColor(achievement.mastery_level)}
            sx={{ height: 8, borderRadius: 4 }}
          />
        </Box>

        <Box sx={{ display: 'flex', gap: 1, flexWrap: 'wrap', mb: 1 }}>
          <Box sx={{ display: 'flex', alignItems: 'center', gap: 0.5 }}>
            <CheckCircle fontSize="small" color="success" />
            <Typography variant="caption">
              {achievement.achieved_criteria.length} Tercapai
            </Typography>
          </Box>
          <Box sx={{ display: 'flex', alignItems: 'center', gap: 0.5 }}>
            <TrendingUp fontSize="small" color="warning" />
            <Typography variant="caption">
              {achievement.pending_criteria.length} Pending
            </Typography>
          </Box>
        </Box>

        <Typography variant="caption" color="text.secondary">
          Terakhir diperbarui: {new Date(achievement.last_updated).toLocaleDateString('id-ID')}
        </Typography>
      </CardContent>
    </Card>
  );
};

export default AchievementCard;

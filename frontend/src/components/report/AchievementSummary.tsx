/**
 * Achievement Summary Component
 * Summary of achievements for reports
 */

import React from 'react';
import {
  Box,
  Typography,
  Card,
  CardContent,
  Chip,
  Divider,
  List,
  ListItem,
  ListItemText,
  ListItemIcon,
} from '@mui/material';
import {
  CheckCircle,
  Star,
  TrendingUp,
  Warning,
} from '@mui/icons-material';

interface CompetencySummary {
  competency_id: string;
  competency_name: string;
  mastery_level: string;
  score: number;
  max_score: number;
}

interface AchievementSummaryData {
  student_id: string;
  student_name: string;
  period_start: string;
  period_end: string;
  overall_mastery: number;
  competency_summary: CompetencySummary[];
  achievements: string[];
  areas_for_improvement: string[];
  recommendations: string[];
}

interface AchievementSummaryProps {
  data: AchievementSummaryData;
}

const AchievementSummary: React.FC<AchievementSummaryProps> = ({ data }) => {
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
    <Card>
      <CardContent>
        <Typography variant="h6" gutterBottom>
          Ringkasan Pencapaian
        </Typography>

        <Box sx={{ mb: 3 }}>
          <Typography variant="subtitle1" gutterBottom>
            {data.student_name}
          </Typography>
          <Typography variant="body2" color="text.secondary" gutterBottom>
            Periode: {new Date(data.period_start).toLocaleDateString('id-ID')} -{' '}
            {new Date(data.period_end).toLocaleDateString('id-ID')}
          </Typography>
          <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
            <Typography variant="body2">Penguasaan Keseluruhan:</Typography>
            <Chip
              label={`${data.overall_mastery.toFixed(1)}%`}
              color={data.overall_mastery >= 80 ? 'success' : data.overall_mastery >= 60 ? 'info' : 'warning'}
              size="small"
            />
          </Box>
        </Box>

        <Divider sx={{ my: 2 }} />

        <Typography variant="subtitle2" gutterBottom>
          Ringkasan Kompetensi
        </Typography>
        <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 2, mb: 3 }}>
          {data.competency_summary.map((competency) => (
            <Box sx={{ width: { xs: '100%', sm: '50%' } }} key={competency.competency_id}>
              <Box
                sx={{
                  p: 2,
                  border: 1,
                  borderColor: 'divider',
                  borderRadius: 1,
                }}
              >
                <Typography variant="body2" gutterBottom>
                  {competency.competency_name}
                </Typography>
                <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                  <Typography variant="caption" color="text.secondary">
                    {competency.score}/{competency.max_score}
                  </Typography>
                  <Chip
                    label={getMasteryLabel(competency.mastery_level)}
                    color={getMasteryColor(competency.mastery_level)}
                    size="small"
                  />
                </Box>
              </Box>
            </Box>
          ))}
        </Box>

        <Divider sx={{ my: 2 }} />

        <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 3 }}>
          <Box sx={{ width: { xs: '100%', sm: '33.33%' } }}>
            <Typography variant="subtitle2" gutterBottom sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
              <Star color="primary" fontSize="small" />
              Pencapaian
            </Typography>
            <List dense>
              {data.achievements.map((achievement, index) => (
                <ListItem key={index}>
                  <ListItemIcon>
                    <CheckCircle color="success" fontSize="small" />
                  </ListItemIcon>
                  <ListItemText primary={achievement} />
                </ListItem>
              ))}
            </List>
          </Box>

          <Box sx={{ width: { xs: '100%', sm: '33.33%' } }}>
            <Typography variant="subtitle2" gutterBottom sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
              <Warning color="warning" fontSize="small" />
              Area Perlu Perbaikan
            </Typography>
            <List dense>
              {data.areas_for_improvement.map((area, index) => (
                <ListItem key={index}>
                  <ListItemIcon>
                    <TrendingUp color="warning" fontSize="small" />
                  </ListItemIcon>
                  <ListItemText primary={area} />
                </ListItem>
              ))}
            </List>
          </Box>

          <Box sx={{ width: { xs: '100%', sm: '33.33%' } }}>
            <Typography variant="subtitle2" gutterBottom sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
              <TrendingUp color="info" fontSize="small" />
              Rekomendasi
            </Typography>
            <List dense>
              {data.recommendations.map((recommendation, index) => (
                <ListItem key={index}>
                  <ListItemIcon>
                    <Star color="info" fontSize="small" />
                  </ListItemIcon>
                  <ListItemText primary={recommendation} />
                </ListItem>
              ))}
            </List>
          </Box>
        </Box>
      </CardContent>
    </Card>
  );
};

export default AchievementSummary;

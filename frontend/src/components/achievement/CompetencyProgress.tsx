/**
 * Competency Progress Component
 * Progress bar/chart for competency
 */

import React from 'react';
import {
  Box,
  Typography,
  Card,
  CardContent,
  LinearProgress,
  Chip,
} from '@mui/material';
import {
  Assignment,
} from '@mui/icons-material';

interface CriteriaProgress {
  criteria_id: string;
  criteria_name: string;
  achieved: boolean;
  evidence_count: number;
}

interface CompetencyProgress {
  competency_id: string;
  competency_name: string;
  total_assessments: number;
  completed_assessments: number;
  average_score: number;
  mastery_level: string;
  progress_percentage: number;
  criteria_progress: CriteriaProgress[];
}

interface CompetencyProgressProps {
  progress: CompetencyProgress;
  showDetails?: boolean;
}

const CompetencyProgress: React.FC<CompetencyProgressProps> = ({
  progress,
  showDetails = false,
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

  const achievedCriteria = progress.criteria_progress.filter((c) => c.achieved).length;
  const totalCriteria = progress.criteria_progress.length;

  return (
    <Card>
      <CardContent>
        <Box sx={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between', mb: 2 }}>
          <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
            <Assignment color="primary" />
            <Typography variant="h6">{progress.competency_name}</Typography>
          </Box>
          <Chip
            label={getMasteryLabel(progress.mastery_level)}
            color={getMasteryColor(progress.mastery_level)}
            size="small"
          />
        </Box>

        <Box sx={{ display: 'grid', gridTemplateColumns: 'repeat(12, 1fr)', gap: 2, mb: 2 }}>
          <Box sx={{ gridColumn: { xs: 'span 6', sm: 'span 3' } }}>
            <Typography variant="caption" color="text.secondary">
              Total Asesmen
            </Typography>
            <Typography variant="body1" fontWeight="bold">
              {progress.total_assessments}
            </Typography>
          </Box>
          <Box sx={{ gridColumn: { xs: 'span 6', sm: 'span 3' } }}>
            <Typography variant="caption" color="text.secondary">
              Selesai
            </Typography>
            <Typography variant="body1" fontWeight="bold">
              {progress.completed_assessments}
            </Typography>
          </Box>
          <Box sx={{ gridColumn: { xs: 'span 6', sm: 'span 3' } }}>
            <Typography variant="caption" color="text.secondary">
              Rata-rata Skor
            </Typography>
            <Typography variant="body1" fontWeight="bold">
              {progress.average_score}
            </Typography>
          </Box>
          <Box sx={{ gridColumn: { xs: 'span 6', sm: 'span 3' } }}>
            <Typography variant="caption" color="text.secondary">
              Progress
            </Typography>
            <Typography variant="body1" fontWeight="bold">
              {progress.progress_percentage}%
            </Typography>
          </Box>
        </Box>

        <Box sx={{ mb: 2 }}>
          <Box sx={{ display: 'flex', justifyContent: 'space-between', mb: 1 }}>
            <Typography variant="body2" color="text.secondary">
              Progress Kompetensi
            </Typography>
            <Typography variant="body2" color="text.secondary">
              {achievedCriteria}/{totalCriteria} Kriteria
            </Typography>
          </Box>
          <LinearProgress
            variant="determinate"
            value={progress.progress_percentage}
            color={getMasteryColor(progress.mastery_level)}
            sx={{ height: 10, borderRadius: 5 }}
          />
        </Box>

        {showDetails && (
          <Box sx={{ mt: 2 }}>
            <Typography variant="subtitle2" gutterBottom>
              Progress Kriteria
            </Typography>
            <Box sx={{ display: 'grid', gridTemplateColumns: 'repeat(12, 1fr)', gap: 1 }}>
              {progress.criteria_progress.map((criteria) => (
                <Box sx={{ gridColumn: { xs: 'span 12', sm: 'span 6' } }} key={criteria.criteria_id}>
                  <Box
                    sx={{
                      p: 1,
                      border: 1,
                      borderColor: criteria.achieved ? 'success.main' : 'divider',
                      borderRadius: 1,
                      bgcolor: criteria.achieved ? 'success.50' : 'background.paper',
                    }}
                  >
                    <Typography variant="caption" display="block">
                      {criteria.criteria_name}
                    </Typography>
                    <Typography variant="caption" color="text.secondary">
                      {criteria.evidence_count} bukti
                    </Typography>
                  </Box>
                </Box>
              ))}
            </Box>
          </Box>
        )}
      </CardContent>
    </Card>
  );
};

export default CompetencyProgress;

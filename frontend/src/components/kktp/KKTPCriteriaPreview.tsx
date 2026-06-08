/**
 * KKTPCriteria Preview Component
 * Preview component for KKTP (Kriteria Ketuntasan Tujuan Pembelajaran)
 * Shows a simplified view suitable for quick reference
 */

import React from 'react';
import {
  Box,
  Typography,
  Card,
  CardContent,
  Chip,
  Grid,
  LinearProgress,
} from '@mui/material';
import {
  CheckCircle,
  Star,
  TrendingUp,
} from '@mui/icons-material';

interface MasteryThresholds {
  excellent_threshold: number;
  proficient_threshold: number;
  developing_threshold: number;
  beginning_threshold: number;
}

interface PerformanceIndicators {
  cognitive: string[];
  psychomotor: string[];
  affective: string[];
}

interface MinimumRequirements {
  core_competencies: string[];
  essential_skills: string[];
  required_evidence: string[];
}

interface KKTPCriteria {
  mastery_thresholds: MasteryThresholds;
  performance_indicators: PerformanceIndicators;
  minimum_requirements: MinimumRequirements;
}

interface KKTPCriteriaPreviewProps {
  data: KKTPCriteria;
  onClick?: () => void;
}

const KKTPCriteriaPreview: React.FC<KKTPCriteriaPreviewProps> = ({
  data,
  onClick,
}) => {
  const totalIndicators =
    data.performance_indicators.cognitive.length +
    data.performance_indicators.psychomotor.length +
    data.performance_indicators.affective.length;

  const totalRequirements =
    data.minimum_requirements.core_competencies.length +
    data.minimum_requirements.essential_skills.length +
    data.minimum_requirements.required_evidence.length;

  return (
    <Card onClick={onClick} sx={{ cursor: onClick ? 'pointer' : 'default' }}>
      <CardContent>
        <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
          <TrendingUp color="primary" sx={{ mr: 1 }} />
          <Typography variant="subtitle1" fontWeight="bold">
            Ambang Penguasaan
          </Typography>
        </Box>

        <Grid container spacing={1} sx={{ mb: 2 }}>
          <Grid item xs={6}>
            <Box>
              <Typography variant="caption" color="text.secondary">
                Sangat Baik
              </Typography>
              <LinearProgress
                variant="determinate"
                value={data.mastery_thresholds.excellent_threshold}
                color="success"
                sx={{ height: 6, borderRadius: 3 }}
              />
            </Box>
          </Grid>
          <Grid item xs={6}>
            <Box>
              <Typography variant="caption" color="text.secondary">
                Baik
              </Typography>
              <LinearProgress
                variant="determinate"
                value={data.mastery_thresholds.proficient_threshold}
                color="info"
                sx={{ height: 6, borderRadius: 3 }}
              />
            </Box>
          </Grid>
          <Grid item xs={6}>
            <Box>
              <Typography variant="caption" color="text.secondary">
                Sedang Berkembang
              </Typography>
              <LinearProgress
                variant="determinate"
                value={data.mastery_thresholds.developing_threshold}
                color="warning"
                sx={{ height: 6, borderRadius: 3 }}
              />
            </Box>
          </Grid>
          <Grid item xs={6}>
            <Box>
              <Typography variant="caption" color="text.secondary">
                Perlu Bimbingan
              </Typography>
              <LinearProgress
                variant="determinate"
                value={data.mastery_thresholds.beginning_threshold}
                color="error"
                sx={{ height: 6, borderRadius: 3 }}
              />
            </Box>
          </Grid>
        </Grid>

        <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
          <CheckCircle color="primary" sx={{ mr: 1, fontSize: 20 }} />
          <Typography variant="subtitle2" fontWeight="bold">
            {totalIndicators} Indikator Pencapaian
          </Typography>
        </Box>

        <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 0.5, mb: 2 }}>
          {data.performance_indicators.cognitive.slice(0, 3).map((indicator, index) => (
            <Chip key={index} label={indicator} size="small" variant="outlined" />
          ))}
          {data.performance_indicators.cognitive.length > 3 && (
            <Chip
              label={`+${data.performance_indicators.cognitive.length - 3} lagi`}
              size="small"
              variant="outlined"
            />
          )}
        </Box>

        <Box sx={{ display: 'flex', alignItems: 'center' }}>
          <Star color="primary" sx={{ mr: 1, fontSize: 20 }} />
          <Typography variant="subtitle2" fontWeight="bold">
            {totalRequirements} Persyaratan Minimum
          </Typography>
        </Box>

        <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 0.5, mt: 2 }}>
          {data.minimum_requirements.core_competencies.slice(0, 2).map((req, index) => (
            <Chip key={index} label={req} size="small" color="secondary" variant="outlined" />
          ))}
          {data.minimum_requirements.core_competencies.length > 2 && (
            <Chip
              label={`+${data.minimum_requirements.core_competencies.length - 2} lagi`}
              size="small"
              color="secondary"
              variant="outlined"
            />
          )}
        </Box>
      </CardContent>
    </Card>
  );
};

export default KKTPCriteriaPreview;

/**
 * KKTPCriteria Display Component
 * Read-only display of KKTP (Kriteria Ketuntasan Tujuan Pembelajaran)
 */

import React from 'react';
import {
  Box,
  Typography,
  Card,
  CardContent,
  Grid,
  Chip,
  LinearProgress,
  Divider,
} from '@mui/material';
import {
  CheckCircle,
  TrendingUp,
  Psychology,
  EmojiObjects,
  Star,
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

interface KKTPCriteriaDisplayProps {
  data: KKTPCriteria;
  compact?: boolean;
}

const KKTPCriteriaDisplay: React.FC<KKTPCriteriaDisplayProps> = ({
  data,
  compact = false,
}) => {
  const getMasteryColor = (level: string): string => {
    switch (level) {
      case 'excellent':
        return 'success';
      case 'proficient':
        return 'info';
      case 'developing':
        return 'warning';
      case 'beginning':
        return 'error';
      default:
        return 'default';
    }
  };

  const getMasteryLabel = (level: string): string => {
    switch (level) {
      case 'excellent':
        return 'Sangat Baik';
      case 'proficient':
        return 'Baik';
      case 'developing':
        return 'Sedang Berkembang';
      case 'beginning':
        return 'Perlu Bimbingan';
      default:
        return level;
    }
  };

  return (
    <Box>
      <Typography variant="h6" gutterBottom>
        Kriteria Ketuntasan Tujuan Pembelajaran (KKTP)
      </Typography>

      {/* Mastery Thresholds */}
      <Card sx={{ mb: 2 }}>
        <CardContent>
          <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
            <TrendingUp color="primary" sx={{ mr: 1 }} />
            <Typography variant="subtitle1" fontWeight="bold">
              Ambang Penguasaan
            </Typography>
          </Box>
          <Grid container spacing={2}>
            <Grid item xs={12} sm={6}>
              <Box sx={{ mb: 1 }}>
                <Typography variant="body2" color="text.secondary">
                  Sangat Baik
                </Typography>
                <LinearProgress
                  variant="determinate"
                  value={data.mastery_thresholds.excellent_threshold}
                  color="success"
                  sx={{ height: 8, borderRadius: 4 }}
                />
                <Typography variant="caption" color="text.secondary">
                  {data.mastery_thresholds.excellent_threshold}%
                </Typography>
              </Box>
            </Grid>
            <Grid item xs={12} sm={6}>
              <Box sx={{ mb: 1 }}>
                <Typography variant="body2" color="text.secondary">
                  Baik
                </Typography>
                <LinearProgress
                  variant="determinate"
                  value={data.mastery_thresholds.proficient_threshold}
                  color="info"
                  sx={{ height: 8, borderRadius: 4 }}
                />
                <Typography variant="caption" color="text.secondary">
                  {data.mastery_thresholds.proficient_threshold}%
                </Typography>
              </Box>
            </Grid>
            <Grid item xs={12} sm={6}>
              <Box sx={{ mb: 1 }}>
                <Typography variant="body2" color="text.secondary">
                  Sedang Berkembang
                </Typography>
                <LinearProgress
                  variant="determinate"
                  value={data.mastery_thresholds.developing_threshold}
                  color="warning"
                  sx={{ height: 8, borderRadius: 4 }}
                />
                <Typography variant="caption" color="text.secondary">
                  {data.mastery_thresholds.developing_threshold}%
                </Typography>
              </Box>
            </Grid>
            <Grid item xs={12} sm={6}>
              <Box sx={{ mb: 1 }}>
                <Typography variant="body2" color="text.secondary">
                  Perlu Bimbingan
                </Typography>
                <LinearProgress
                  variant="determinate"
                  value={data.mastery_thresholds.beginning_threshold}
                  color="error"
                  sx={{ height: 8, borderRadius: 4 }}
                />
                <Typography variant="caption" color="text.secondary">
                  {data.mastery_thresholds.beginning_threshold}%
                </Typography>
              </Box>
            </Grid>
          </Grid>
        </CardContent>
      </Card>

      {/* Performance Indicators */}
      <Card sx={{ mb: 2 }}>
        <CardContent>
          <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
            <Psychology color="primary" sx={{ mr: 1 }} />
            <Typography variant="subtitle1" fontWeight="bold">
              Indikator Pencapaian
            </Typography>
          </Box>
          <Grid container spacing={2}>
            <Grid item xs={12}>
              <Typography variant="body2" color="text.secondary" gutterBottom>
                Kognitif
              </Typography>
              <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 1 }}>
                {data.performance_indicators.cognitive.map((indicator, index) => (
                  <Chip
                    key={index}
                    label={indicator}
                    size="small"
                    color="primary"
                    variant="outlined"
                    icon={<CheckCircle fontSize="small" />}
                  />
                ))}
              </Box>
            </Grid>
            <Grid item xs={12}>
              <Typography variant="body2" color="text.secondary" gutterBottom>
                Psikomotorik
              </Typography>
              <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 1 }}>
                {data.performance_indicators.psychomotor.map((indicator, index) => (
                  <Chip
                    key={index}
                    label={indicator}
                    size="small"
                    color="primary"
                    variant="outlined"
                    icon={<CheckCircle fontSize="small" />}
                  />
                ))}
              </Box>
            </Grid>
            <Grid item xs={12}>
              <Typography variant="body2" color="text.secondary" gutterBottom>
                Afektif
              </Typography>
              <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 1 }}>
                {data.performance_indicators.affective.map((indicator, index) => (
                  <Chip
                    key={index}
                    label={indicator}
                    size="small"
                    color="primary"
                    variant="outlined"
                    icon={<CheckCircle fontSize="small" />}
                  />
                ))}
              </Box>
            </Grid>
          </Grid>
        </CardContent>
      </Card>

      {/* Minimum Requirements */}
      <Card>
        <CardContent>
          <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
            <Star color="primary" sx={{ mr: 1 }} />
            <Typography variant="subtitle1" fontWeight="bold">
              Persyaratan Minimum
            </Typography>
          </Box>
          <Grid container spacing={2}>
            <Grid item xs={12}>
              <Typography variant="body2" color="text.secondary" gutterBottom>
                Kompetensi Inti
              </Typography>
              <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 1 }}>
                {data.minimum_requirements.core_competencies.map((req, index) => (
                  <Chip
                    key={index}
                    label={req}
                    size="small"
                    color="secondary"
                    variant="outlined"
                    icon={<EmojiObjects fontSize="small" />}
                  />
                ))}
              </Box>
            </Grid>
            <Grid item xs={12}>
              <Typography variant="body2" color="text.secondary" gutterBottom>
                Keterampilan Esensial
              </Typography>
              <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 1 }}>
                {data.minimum_requirements.essential_skills.map((req, index) => (
                  <Chip
                    key={index}
                    label={req}
                    size="small"
                    color="secondary"
                    variant="outlined"
                    icon={<EmojiObjects fontSize="small" />}
                  />
                ))}
              </Box>
            </Grid>
            <Grid item xs={12}>
              <Typography variant="body2" color="text.secondary" gutterBottom>
                Bukti yang Diperlukan
              </Typography>
              <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 1 }}>
                {data.minimum_requirements.required_evidence.map((req, index) => (
                  <Chip
                    key={index}
                    label={req}
                    size="small"
                    color="secondary"
                    variant="outlined"
                    icon={<CheckCircle fontSize="small" />}
                  />
                ))}
              </Box>
            </Grid>
          </Grid>
        </CardContent>
      </Card>
    </Box>
  );
};

export default KKTPCriteriaDisplay;

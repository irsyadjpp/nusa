/**
 * Success Criteria Snapshot Component
 * Display of snapshot criteria from TP at assessment creation time
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
  Alert,
  AlertTitle,
} from '@mui/material';
import {
  Info,
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

interface SuccessCriteriaSnapshot {
  mastery_thresholds: MasteryThresholds;
  performance_indicators: PerformanceIndicators;
  minimum_requirements: MinimumRequirements;
}

interface SuccessCriteriaSnapshotProps {
  snapshot: SuccessCriteriaSnapshot;
  tpVersion: number;
  showWarning?: boolean;
}

const SuccessCriteriaSnapshot: React.FC<SuccessCriteriaSnapshotProps> = ({
  snapshot,
  tpVersion,
  showWarning = false,
}) => {
  return (
    <Box>
      {showWarning && (
        <Alert severity="info" sx={{ mb: 2 }}>
          <AlertTitle>Snapshot Kriteria Ketuntasan</AlertTitle>
          Ini adalah snapshot dari kriteria ketuntasan TP pada saat asesmen dibuat.
          Kriteria ini tidak akan berubah meskipun TP asli diperbarui.
        </Alert>
      )}

      <Card>
        <CardContent>
          <Box sx={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between', mb: 2 }}>
            <Box sx={{ display: 'flex', alignItems: 'center' }}>
              <Info color="primary" sx={{ mr: 1 }} />
              <Typography variant="subtitle1" fontWeight="bold">
                Snapshot Kriteria Ketuntasan
              </Typography>
            </Box>
            <Chip
              label={`Versi TP: ${tpVersion}`}
              size="small"
              color="primary"
              variant="outlined"
            />
          </Box>

          {/* Mastery Thresholds */}
          <Box sx={{ mb: 3 }}>
            <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
              <TrendingUp color="primary" sx={{ mr: 1, fontSize: 20 }} />
              <Typography variant="subtitle2" fontWeight="bold">
                Ambang Penguasaan
              </Typography>
            </Box>
            <Grid container spacing={2}>
              <Grid item xs={12} sm={6}>
                <Box sx={{ mb: 1 }}>
                  <Typography variant="caption" color="text.secondary">
                    Sangat Baik
                  </Typography>
                  <LinearProgress
                    variant="determinate"
                    value={snapshot.mastery_thresholds.excellent_threshold}
                    color="success"
                    sx={{ height: 6, borderRadius: 3 }}
                  />
                  <Typography variant="caption" color="text.secondary">
                    {snapshot.mastery_thresholds.excellent_threshold}%
                  </Typography>
                </Box>
              </Grid>
              <Grid item xs={12} sm={6}>
                <Box sx={{ mb: 1 }}>
                  <Typography variant="caption" color="text.secondary">
                    Baik
                  </Typography>
                  <LinearProgress
                    variant="determinate"
                    value={snapshot.mastery_thresholds.proficient_threshold}
                    color="info"
                    sx={{ height: 6, borderRadius: 3 }}
                  />
                  <Typography variant="caption" color="text.secondary">
                    {snapshot.mastery_thresholds.proficient_threshold}%
                  </Typography>
                </Box>
              </Grid>
              <Grid item xs={12} sm={6}>
                <Box sx={{ mb: 1 }}>
                  <Typography variant="caption" color="text.secondary">
                    Sedang Berkembang
                  </Typography>
                  <LinearProgress
                    variant="determinate"
                    value={snapshot.mastery_thresholds.developing_threshold}
                    color="warning"
                    sx={{ height: 6, borderRadius: 3 }}
                  />
                  <Typography variant="caption" color="text.secondary">
                    {snapshot.mastery_thresholds.developing_threshold}%
                  </Typography>
                </Box>
              </Grid>
              <Grid item xs={12} sm={6}>
                <Box sx={{ mb: 1 }}>
                  <Typography variant="caption" color="text.secondary">
                    Perlu Bimbingan
                  </Typography>
                  <LinearProgress
                    variant="determinate"
                    value={snapshot.mastery_thresholds.beginning_threshold}
                    color="error"
                    sx={{ height: 6, borderRadius: 3 }}
                  />
                  <Typography variant="caption" color="text.secondary">
                    {snapshot.mastery_thresholds.beginning_threshold}%
                  </Typography>
                </Box>
              </Grid>
            </Grid>
          </Box>

          {/* Performance Indicators */}
          <Box sx={{ mb: 3 }}>
            <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
              <CheckCircle color="primary" sx={{ mr: 1, fontSize: 20 }} />
              <Typography variant="subtitle2" fontWeight="bold">
                Indikator Pencapaian
              </Typography>
            </Box>
            <Grid container spacing={2}>
              <Grid item xs={12}>
                <Typography variant="caption" color="text.secondary" gutterBottom>
                  Kognitif
                </Typography>
                <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 0.5 }}>
                  {snapshot.performance_indicators.cognitive.map((indicator, index) => (
                    <Chip
                      key={index}
                      label={indicator}
                      size="small"
                      variant="outlined"
                    />
                  ))}
                </Box>
              </Grid>
              <Grid item xs={12}>
                <Typography variant="caption" color="text.secondary" gutterBottom>
                  Psikomotorik
                </Typography>
                <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 0.5 }}>
                  {snapshot.performance_indicators.psychomotor.map((indicator, index) => (
                    <Chip
                      key={index}
                      label={indicator}
                      size="small"
                      variant="outlined"
                    />
                  ))}
                </Box>
              </Grid>
              <Grid item xs={12}>
                <Typography variant="caption" color="text.secondary" gutterBottom>
                  Afektif
                </Typography>
                <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 0.5 }}>
                  {snapshot.performance_indicators.affective.map((indicator, index) => (
                    <Chip
                      key={index}
                      label={indicator}
                      size="small"
                      variant="outlined"
                    />
                  ))}
                </Box>
              </Grid>
            </Grid>
          </Box>

          {/* Minimum Requirements */}
          <Box>
            <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
              <Star color="primary" sx={{ mr: 1, fontSize: 20 }} />
              <Typography variant="subtitle2" fontWeight="bold">
                Persyaratan Minimum
              </Typography>
            </Box>
            <Grid container spacing={2}>
              <Grid item xs={12}>
                <Typography variant="caption" color="text.secondary" gutterBottom>
                  Kompetensi Inti
                </Typography>
                <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 0.5 }}>
                  {snapshot.minimum_requirements.core_competencies.map((req, index) => (
                    <Chip
                      key={index}
                      label={req}
                      size="small"
                      color="secondary"
                      variant="outlined"
                    />
                  ))}
                </Box>
              </Grid>
              <Grid item xs={12}>
                <Typography variant="caption" color="text.secondary" gutterBottom>
                  Keterampilan Esensial
                </Typography>
                <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 0.5 }}>
                  {snapshot.minimum_requirements.essential_skills.map((req, index) => (
                    <Chip
                      key={index}
                      label={req}
                      size="small"
                      color="secondary"
                      variant="outlined"
                    />
                  ))}
                </Box>
              </Grid>
              <Grid item xs={12}>
                <Typography variant="caption" color="text.secondary" gutterBottom>
                  Bukti yang Diperlukan
                </Typography>
                <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 0.5 }}>
                  {snapshot.minimum_requirements.required_evidence.map((req, index) => (
                    <Chip
                      key={index}
                      label={req}
                      size="small"
                      color="secondary"
                      variant="outlined"
                    />
                  ))}
                </Box>
              </Grid>
            </Grid>
          </Box>
        </CardContent>
      </Card>
    </Box>
  );
};

export default SuccessCriteriaSnapshot;

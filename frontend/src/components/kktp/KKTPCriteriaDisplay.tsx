/**
 * KKTP Criteria Display Component
 * Display-only component for Kriteria Ketuntasan Tujuan Pembelajaran (KKTP)
 */

import React from 'react';
import {
  Box,
  Typography,
  Card,
  CardContent,
  Divider,
  Chip,
} from '@mui/material';

interface MasteryThreshold {
  level: string;
  description: string;
  min_score: number;
}

interface PerformanceIndicator {
  indicator: string;
  description: string;
}

interface MinimumRequirement {
  requirement: string;
  description: string;
}

interface KKTPCriteriaData {
  mastery_thresholds?: MasteryThreshold[];
  performance_indicators?: PerformanceIndicator[];
  minimum_requirements?: MinimumRequirement[];
}

interface KKTPCriteriaDisplayProps {
  data?: KKTPCriteriaData;
}

const KKTPCriteriaDisplay: React.FC<KKTPCriteriaDisplayProps> = ({ data }) => {
  if (!data) {
    return (
      <Typography variant="body2" color="text.secondary">
        Tidak ada data KKTP
      </Typography>
    );
  }

  return (
    <Box>
      {/* Mastery Thresholds */}
      {data.mastery_thresholds && data.mastery_thresholds.length > 0 && (
        <Box sx={{ mb: 3 }}>
          <Typography variant="subtitle1" fontWeight="medium" gutterBottom>
            Ambang Batas Penguasaan
          </Typography>
          <Box sx={{ display: 'grid', gridTemplateColumns: 'repeat(12, 1fr)', gap: 2 }}>
            {data.mastery_thresholds.map((threshold, index) => (
              <Box sx={{ gridColumn: { xs: 'span 12', sm: 'span 6' } }} key={index}>
                <Card variant="outlined">
                  <CardContent>
                    <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 1 }}>
                      <Chip
                        label={threshold.level}
                        size="small"
                        color="primary"
                      />
                      <Typography variant="caption" color="text.secondary">
                        Min: {threshold.min_score}
                      </Typography>
                    </Box>
                    <Typography variant="body2" color="text.secondary">
                      {threshold.description}
                    </Typography>
                  </CardContent>
                </Card>
              </Box>
            ))}
          </Box>
        </Box>
      )}

      <Divider sx={{ my: 2 }} />

      {/* Performance Indicators */}
      {data.performance_indicators && data.performance_indicators.length > 0 && (
        <Box sx={{ mb: 3 }}>
          <Typography variant="subtitle1" fontWeight="medium" gutterBottom>
            Indikator Kinerja
          </Typography>
          {data.performance_indicators.map((indicator, index) => (
            <Box key={index} sx={{ mb: 2 }}>
              <Typography variant="body2" fontWeight="medium">
                {indicator.indicator}
              </Typography>
              <Typography variant="body2" color="text.secondary">
                {indicator.description}
              </Typography>
            </Box>
          ))}
        </Box>
      )}

      <Divider sx={{ my: 2 }} />

      {/* Minimum Requirements */}
      {data.minimum_requirements && data.minimum_requirements.length > 0 && (
        <Box>
          <Typography variant="subtitle1" fontWeight="medium" gutterBottom>
            Persyaratan Minimum
          </Typography>
          {data.minimum_requirements.map((requirement, index) => (
            <Box key={index} sx={{ mb: 2 }}>
              <Typography variant="body2" fontWeight="medium">
                {requirement.requirement}
              </Typography>
              <Typography variant="body2" color="text.secondary">
                {requirement.description}
              </Typography>
            </Box>
          ))}
        </Box>
      )}

      {/* Empty State */}
      {!data.mastery_thresholds?.length && 
       !data.performance_indicators?.length && 
       !data.minimum_requirements?.length && (
        <Typography variant="body2" color="text.secondary">
          Tidak ada data KKTP tersedia
        </Typography>
      )}
    </Box>
  );
};

export default KKTPCriteriaDisplay;

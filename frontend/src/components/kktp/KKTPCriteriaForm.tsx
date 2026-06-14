/**
 * KKTP Criteria Form Component
 * Form component for editing Kriteria Ketuntasan Tujuan Pembelajaran (KKTP)
 */

import React from 'react';
import {
  Box,
  Typography,
  TextField,
  Button,
  Card,
  CardContent,
  IconButton,
  Divider,
} from '@mui/material';
import {
  Add as AddIcon,
  Delete as DeleteIcon,
} from '@mui/icons-material';

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

interface KKTPCriteriaFormProps {
  data?: KKTPCriteriaData;
  onChange: (data: KKTPCriteriaData) => void;
}

const KKTPCriteriaForm: React.FC<KKTPCriteriaFormProps> = ({ data, onChange }) => {
  const handleChange = (field: keyof KKTPCriteriaData, value: any) => {
    onChange({
      ...data,
      [field]: value,
    });
  };

  const addMasteryThreshold = () => {
    const newThreshold: MasteryThreshold = {
      level: '',
      description: '',
      min_score: 70,
    };
    handleChange('mastery_thresholds', [
      ...(data?.mastery_thresholds || []),
      newThreshold,
    ]);
  };

  const updateMasteryThreshold = (index: number, field: keyof MasteryThreshold, value: any) => {
    const thresholds = data?.mastery_thresholds || [];
    const updated = [...thresholds];
    updated[index] = { ...updated[index], [field]: value };
    handleChange('mastery_thresholds', updated);
  };

  const removeMasteryThreshold = (index: number) => {
    const thresholds = data?.mastery_thresholds || [];
    handleChange('mastery_thresholds', thresholds.filter((_, i) => i !== index));
  };

  const addPerformanceIndicator = () => {
    const newIndicator: PerformanceIndicator = {
      indicator: '',
      description: '',
    };
    handleChange('performance_indicators', [
      ...(data?.performance_indicators || []),
      newIndicator,
    ]);
  };

  const updatePerformanceIndicator = (index: number, field: keyof PerformanceIndicator, value: any) => {
    const indicators = data?.performance_indicators || [];
    const updated = [...indicators];
    updated[index] = { ...updated[index], [field]: value };
    handleChange('performance_indicators', updated);
  };

  const removePerformanceIndicator = (index: number) => {
    const indicators = data?.performance_indicators || [];
    handleChange('performance_indicators', indicators.filter((_, i) => i !== index));
  };

  const addMinimumRequirement = () => {
    const newRequirement: MinimumRequirement = {
      requirement: '',
      description: '',
    };
    handleChange('minimum_requirements', [
      ...(data?.minimum_requirements || []),
      newRequirement,
    ]);
  };

  const updateMinimumRequirement = (index: number, field: keyof MinimumRequirement, value: any) => {
    const requirements = data?.minimum_requirements || [];
    const updated = [...requirements];
    updated[index] = { ...updated[index], [field]: value };
    handleChange('minimum_requirements', updated);
  };

  const removeMinimumRequirement = (index: number) => {
    const requirements = data?.minimum_requirements || [];
    handleChange('minimum_requirements', requirements.filter((_, i) => i !== index));
  };

  return (
    <Box>
      {/* Mastery Thresholds */}
      <Box sx={{ mb: 3 }}>
        <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 2 }}>
          <Typography variant="subtitle1" fontWeight="medium">
            Ambang Batas Penguasaan
          </Typography>
          <Button
            startIcon={<AddIcon />}
            onClick={addMasteryThreshold}
            size="small"
            variant="outlined"
          >
            Tambah
          </Button>
        </Box>
        {data?.mastery_thresholds?.map((threshold, index) => (
          <Card key={index} variant="outlined" sx={{ mb: 2 }}>
            <CardContent>
              <Box sx={{ display: 'grid', gridTemplateColumns: 'repeat(12, 1fr)', gap: 2 }}>
                <Box sx={{ gridColumn: { xs: 'span 12', sm: 'span 4' } }}>
                  <TextField
                    fullWidth
                    label="Level"
                    value={threshold.level}
                    onChange={(e) => updateMasteryThreshold(index, 'level', e.target.value)}
                    size="small"
                  />
                </Box>
                <Box sx={{ gridColumn: { xs: 'span 12', sm: 'span 3' } }}>
                  <TextField
                    fullWidth
                    label="Min Score"
                    type="number"
                    value={threshold.min_score}
                    onChange={(e) => updateMasteryThreshold(index, 'min_score', parseInt(e.target.value))}
                    size="small"
                  />
                </Box>
                <Box sx={{ gridColumn: { xs: 'span 10', sm: 'span 4' } }}>
                  <TextField
                    fullWidth
                    label="Deskripsi"
                    value={threshold.description}
                    onChange={(e) => updateMasteryThreshold(index, 'description', e.target.value)}
                    size="small"
                  />
                </Box>
                <Box sx={{ gridColumn: { xs: 'span 2', sm: 'span 1' } }}>
                  <IconButton
                    onClick={() => removeMasteryThreshold(index)}
                    color="error"
                    size="small"
                  >
                    <DeleteIcon />
                  </IconButton>
                </Box>
              </Box>
            </CardContent>
          </Card>
        ))}
      </Box>

      <Divider sx={{ my: 2 }} />

      {/* Performance Indicators */}
      <Box sx={{ mb: 3 }}>
        <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 2 }}>
          <Typography variant="subtitle1" fontWeight="medium">
            Indikator Kinerja
          </Typography>
          <Button
            startIcon={<AddIcon />}
            onClick={addPerformanceIndicator}
            size="small"
            variant="outlined"
          >
            Tambah
          </Button>
        </Box>
        {data?.performance_indicators?.map((indicator, index) => (
          <Card key={index} variant="outlined" sx={{ mb: 2 }}>
            <CardContent>
              <Box sx={{ display: 'grid', gridTemplateColumns: 'repeat(12, 1fr)', gap: 2 }}>
                <Box sx={{ gridColumn: { xs: 'span 10', sm: 'span 5' } }}>
                  <TextField
                    fullWidth
                    label="Indikator"
                    value={indicator.indicator}
                    onChange={(e) => updatePerformanceIndicator(index, 'indicator', e.target.value)}
                    size="small"
                  />
                </Box>
                <Box sx={{ gridColumn: { xs: 'span 10', sm: 'span 6' } }}>
                  <TextField
                    fullWidth
                    label="Deskripsi"
                    value={indicator.description}
                    onChange={(e) => updatePerformanceIndicator(index, 'description', e.target.value)}
                    size="small"
                  />
                </Box>
                <Box sx={{ gridColumn: { xs: 'span 2', sm: 'span 1' } }}>
                  <IconButton
                    onClick={() => removePerformanceIndicator(index)}
                    color="error"
                    size="small"
                  >
                    <DeleteIcon />
                  </IconButton>
                </Box>
              </Box>
            </CardContent>
          </Card>
        ))}
      </Box>

      <Divider sx={{ my: 2 }} />

      {/* Minimum Requirements */}
      <Box>
        <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 2 }}>
          <Typography variant="subtitle1" fontWeight="medium">
            Persyaratan Minimum
          </Typography>
          <Button
            startIcon={<AddIcon />}
            onClick={addMinimumRequirement}
            size="small"
            variant="outlined"
          >
            Tambah
          </Button>
        </Box>
        {data?.minimum_requirements?.map((requirement, index) => (
          <Card key={index} variant="outlined" sx={{ mb: 2 }}>
            <CardContent>
              <Box sx={{ display: 'grid', gridTemplateColumns: 'repeat(12, 1fr)', gap: 2 }}>
                <Box sx={{ gridColumn: { xs: 'span 10', sm: 'span 5' } }}>
                  <TextField
                    fullWidth
                    label="Persyaratan"
                    value={requirement.requirement}
                    onChange={(e) => updateMinimumRequirement(index, 'requirement', e.target.value)}
                    size="small"
                  />
                </Box>
                <Box sx={{ gridColumn: { xs: 'span 10', sm: 'span 6' } }}>
                  <TextField
                    fullWidth
                    label="Deskripsi"
                    value={requirement.description}
                    onChange={(e) => updateMinimumRequirement(index, 'description', e.target.value)}
                    size="small"
                  />
                </Box>
                <Box sx={{ gridColumn: { xs: 'span 2', sm: 'span 1' } }}>
                  <IconButton
                    onClick={() => removeMinimumRequirement(index)}
                    color="error"
                    size="small"
                  >
                    <DeleteIcon />
                  </IconButton>
                </Box>
              </Box>
            </CardContent>
          </Card>
        ))}
      </Box>
    </Box>
  );
};

export default KKTPCriteriaForm;

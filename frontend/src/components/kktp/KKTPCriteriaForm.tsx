/**
 * KKTPCriteria Form Component
 * Form for editing KKTP (Kriteria Ketuntasan Tujuan Pembelajaran)
 */

import React from 'react';
import {
  Box,
  Typography,
  Accordion,
  AccordionSummary,
  AccordionDetails,
  TextField,
  Grid,
  Slider,
  Chip,
  IconButton,
  Button,
} from '@mui/material';
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
import AddIcon from '@mui/icons-material/Add';
import DeleteIcon from '@mui/icons-material/Delete';

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

interface KKTPCriteriaFormProps {
  data: KKTPCriteria;
  onChange: (data: KKTPCriteria) => void;
  disabled?: boolean;
}

const KKTPCriteriaForm: React.FC<KKTPCriteriaFormProps> = ({
  data,
  onChange,
  disabled = false,
}) => {
  const handleMasteryThresholdChange = (field: keyof MasteryThresholds, value: number) => {
    onChange({
      ...data,
      mastery_thresholds: {
        ...data.mastery_thresholds,
        [field]: value,
      },
    });
  };

  const handlePerformanceIndicatorAdd = (category: keyof PerformanceIndicators, value: string) => {
    if (!value.trim()) return;
    onChange({
      ...data,
      performance_indicators: {
        ...data.performance_indicators,
        [category]: [...data.performance_indicators[category], value.trim()],
      },
    });
  };

  const handlePerformanceIndicatorRemove = (
    category: keyof PerformanceIndicators,
    index: number,
  ) => {
    onChange({
      ...data,
      performance_indicators: {
        ...data.performance_indicators,
        [category]: data.performance_indicators[category].filter((_, i) => i !== index),
      },
    });
  };

  const handleMinimumRequirementAdd = (category: keyof MinimumRequirements, value: string) => {
    if (!value.trim()) return;
    onChange({
      ...data,
      minimum_requirements: {
        ...data.minimum_requirements,
        [category]: [...data.minimum_requirements[category], value.trim()],
      },
    });
  };

  const handleMinimumRequirementRemove = (
    category: keyof MinimumRequirements,
    index: number,
  ) => {
    onChange({
      ...data,
      minimum_requirements: {
        ...data.minimum_requirements,
        [category]: data.minimum_requirements[category].filter((_, i) => i !== index),
      },
    });
  };

  const renderChips = (
    items: string[],
    category: keyof PerformanceIndicators | keyof MinimumRequirements,
    type: 'performance' | 'requirement',
  ) => (
    <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 1, mt: 1 }}>
      {items.map((item, index) => (
        <Chip
          key={index}
          label={item}
          onDelete={
            !disabled
              ? () => {
                  if (type === 'performance') {
                    handlePerformanceIndicatorRemove(category as keyof PerformanceIndicators, index);
                  } else {
                    handleMinimumRequirementRemove(category as keyof MinimumRequirements, index);
                  }
                }
              : undefined
          }
          color={type === 'performance' ? 'primary' : 'secondary'}
          size="small"
        />
      ))}
    </Box>
  );

  return (
    <Box>
      <Typography variant="h6" gutterBottom>
        Kriteria Ketuntasan Tujuan Pembelajaran (KKTP)
      </Typography>

      {/* Mastery Thresholds */}
      <Accordion defaultExpanded>
        <AccordionSummary expandIcon={<ExpandMoreIcon />}>
          <Typography variant="subtitle1">Ambang Penguasaan</Typography>
        </AccordionSummary>
        <AccordionDetails>
          <Grid container spacing={3}>
            <Grid item xs={12}>
              <Typography variant="body2" color="text.secondary" gutterBottom>
                Excellent (Sangat Baik)
              </Typography>
              <Slider
                value={data.mastery_thresholds.excellent_threshold}
                onChange={(_, value) => handleMasteryThresholdChange('excellent_threshold', value as number)}
                min={0}
                max={100}
                valueLabelDisplay="auto"
                disabled={disabled}
                sx={{ color: 'success.main' }}
              />
            </Grid>
            <Grid item xs={12}>
              <Typography variant="body2" color="text.secondary" gutterBottom>
                Proficient (Baik)
              </Typography>
              <Slider
                value={data.mastery_thresholds.proficient_threshold}
                onChange={(_, value) => handleMasteryThresholdChange('proficient_threshold', value as number)}
                min={0}
                max={100}
                valueLabelDisplay="auto"
                disabled={disabled}
                sx={{ color: 'info.main' }}
              />
            </Grid>
            <Grid item xs={12}>
              <Typography variant="body2" color="text.secondary" gutterBottom>
                Developing (Sedang Berkembang)
              </Typography>
              <Slider
                value={data.mastery_thresholds.developing_threshold}
                onChange={(_, value) => handleMasteryThresholdChange('developing_threshold', value as number)}
                min={0}
                max={100}
                valueLabelDisplay="auto"
                disabled={disabled}
                sx={{ color: 'warning.main' }}
              />
            </Grid>
            <Grid item xs={12}>
              <Typography variant="body2" color="text.secondary" gutterBottom>
                Beginning (Perlu Bimbingan)
              </Typography>
              <Slider
                value={data.mastery_thresholds.beginning_threshold}
                onChange={(_, value) => handleMasteryThresholdChange('beginning_threshold', value as number)}
                min={0}
                max={100}
                valueLabelDisplay="auto"
                disabled={disabled}
                sx={{ color: 'error.main' }}
              />
            </Grid>
          </Grid>
        </AccordionDetails>
      </Accordion>

      {/* Performance Indicators */}
      <Accordion>
        <AccordionSummary expandIcon={<ExpandMoreIcon />}>
          <Typography variant="subtitle1">Indikator Pencapaian</Typography>
        </AccordionSummary>
        <AccordionDetails>
          <Grid container spacing={3}>
            <Grid item xs={12}>
              <Typography variant="body2" gutterBottom>
                Kognitif
              </Typography>
              <Box sx={{ display: 'flex', gap: 1, mb: 1 }}>
                <TextField
                  size="small"
                  placeholder="Tambah indikator kognitif"
                  onKeyPress={(e) => {
                    if (e.key === 'Enter') {
                      handlePerformanceIndicatorAdd('cognitive', (e.target as HTMLInputElement).value);
                      (e.target as HTMLInputElement).value = '';
                    }
                  }}
                  disabled={disabled}
                  fullWidth
                />
                {!disabled && (
                  <IconButton
                    onClick={(e) => {
                      const input = e.currentTarget.previousElementSibling as HTMLInputElement;
                      handlePerformanceIndicatorAdd('cognitive', input.value);
                      input.value = '';
                    }}
                  >
                    <AddIcon />
                  </IconButton>
                )}
              </Box>
              {renderChips(data.performance_indicators.cognitive, 'cognitive', 'performance')}
            </Grid>
            <Grid item xs={12}>
              <Typography variant="body2" gutterBottom>
                Psikomotorik
              </Typography>
              <Box sx={{ display: 'flex', gap: 1, mb: 1 }}>
                <TextField
                  size="small"
                  placeholder="Tambah indikator psikomotorik"
                  onKeyPress={(e) => {
                    if (e.key === 'Enter') {
                      handlePerformanceIndicatorAdd('psychomotor', (e.target as HTMLInputElement).value);
                      (e.target as HTMLInputElement).value = '';
                    }
                  }}
                  disabled={disabled}
                  fullWidth
                />
                {!disabled && (
                  <IconButton
                    onClick={(e) => {
                      const input = e.currentTarget.previousElementSibling as HTMLInputElement;
                      handlePerformanceIndicatorAdd('psychomotor', input.value);
                      input.value = '';
                    }}
                  >
                    <AddIcon />
                  </IconButton>
                )}
              </Box>
              {renderChips(data.performance_indicators.psychomotor, 'psychomotor', 'performance')}
            </Grid>
            <Grid item xs={12}>
              <Typography variant="body2" gutterBottom>
                Afektif
              </Typography>
              <Box sx={{ display: 'flex', gap: 1, mb: 1 }}>
                <TextField
                  size="small"
                  placeholder="Tambah indikator afektif"
                  onKeyPress={(e) => {
                    if (e.key === 'Enter') {
                      handlePerformanceIndicatorAdd('affective', (e.target as HTMLInputElement).value);
                      (e.target as HTMLInputElement).value = '';
                    }
                  }}
                  disabled={disabled}
                  fullWidth
                />
                {!disabled && (
                  <IconButton
                    onClick={(e) => {
                      const input = e.currentTarget.previousElementSibling as HTMLInputElement;
                      handlePerformanceIndicatorAdd('affective', input.value);
                      input.value = '';
                    }}
                  >
                    <AddIcon />
                  </IconButton>
                )}
              </Box>
              {renderChips(data.performance_indicators.affective, 'affective', 'performance')}
            </Grid>
          </Grid>
        </AccordionDetails>
      </Accordion>

      {/* Minimum Requirements */}
      <Accordion>
        <AccordionSummary expandIcon={<ExpandMoreIcon />}>
          <Typography variant="subtitle1">Persyaratan Minimum</Typography>
        </AccordionSummary>
        <AccordionDetails>
          <Grid container spacing={3}>
            <Grid item xs={12}>
              <Typography variant="body2" gutterBottom>
                Kompetensi Inti
              </Typography>
              <Box sx={{ display: 'flex', gap: 1, mb: 1 }}>
                <TextField
                  size="small"
                  placeholder="Tambah kompetensi inti"
                  onKeyPress={(e) => {
                    if (e.key === 'Enter') {
                      handleMinimumRequirementAdd('core_competencies', (e.target as HTMLInputElement).value);
                      (e.target as HTMLInputElement).value = '';
                    }
                  }}
                  disabled={disabled}
                  fullWidth
                />
                {!disabled && (
                  <IconButton
                    onClick={(e) => {
                      const input = e.currentTarget.previousElementSibling as HTMLInputElement;
                      handleMinimumRequirementAdd('core_competencies', input.value);
                      input.value = '';
                    }}
                  >
                    <AddIcon />
                  </IconButton>
                )}
              </Box>
              {renderChips(data.minimum_requirements.core_competencies, 'core_competencies', 'requirement')}
            </Grid>
            <Grid item xs={12}>
              <Typography variant="body2" gutterBottom>
                Keterampilan Esensial
              </Typography>
              <Box sx={{ display: 'flex', gap: 1, mb: 1 }}>
                <TextField
                  size="small"
                  placeholder="Tambah keterampilan esensial"
                  onKeyPress={(e) => {
                    if (e.key === 'Enter') {
                      handleMinimumRequirementAdd('essential_skills', (e.target as HTMLInputElement).value);
                      (e.target as HTMLInputElement).value = '';
                    }
                  }}
                  disabled={disabled}
                  fullWidth
                />
                {!disabled && (
                  <IconButton
                    onClick={(e) => {
                      const input = e.currentTarget.previousElementSibling as HTMLInputElement;
                      handleMinimumRequirementAdd('essential_skills', input.value);
                      input.value = '';
                    }}
                  >
                    <AddIcon />
                  </IconButton>
                )}
              </Box>
              {renderChips(data.minimum_requirements.essential_skills, 'essential_skills', 'requirement')}
            </Grid>
            <Grid item xs={12}>
              <Typography variant="body2" gutterBottom>
                Bukti yang Diperlukan
              </Typography>
              <Box sx={{ display: 'flex', gap: 1, mb: 1 }}>
                <TextField
                  size="small"
                  placeholder="Tambah bukti yang diperlukan"
                  onKeyPress={(e) => {
                    if (e.key === 'Enter') {
                      handleMinimumRequirementAdd('required_evidence', (e.target as HTMLInputElement).value);
                      (e.target as HTMLInputElement).value = '';
                    }
                  }}
                  disabled={disabled}
                  fullWidth
                />
                {!disabled && (
                  <IconButton
                    onClick={(e) => {
                      const input = e.currentTarget.previousElementSibling as HTMLInputElement;
                      handleMinimumRequirementAdd('required_evidence', input.value);
                      input.value = '';
                    }}
                  >
                    <AddIcon />
                  </IconButton>
                )}
              </Box>
              {renderChips(data.minimum_requirements.required_evidence, 'required_evidence', 'requirement')}
            </Grid>
          </Grid>
        </AccordionDetails>
      </Accordion>
    </Box>
  );
};

export default KKTPCriteriaForm;

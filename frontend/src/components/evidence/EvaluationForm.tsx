/**
 * Evaluation Form Component
 * Form for evaluating evidence with rubric
 */

import React from 'react';
import {
  Box,
  Typography,
  Card,
  CardContent,
  TextField,
  Button,
  Grid,
  Slider,
  Chip,
  CircularProgress,
} from '@mui/material';
import { Formik, Form, Field, FormikHelpers } from 'formik';
import * as Yup from 'yup';
import TeacherFeedback from './TeacherFeedback';

interface EvaluationFormProps {
  initialValues?: {
    performance_scores?: any;
    total_score?: number;
    max_score?: number;
    performance_level?: string;
    teacher_feedback?: string;
  };
  rubric?: any;
  onSubmit: (values: any, helpers: FormikHelpers<any>) => void;
  onCancel?: () => void;
  loading?: boolean;
}

const validationSchema = Yup.object().shape({
  performance_scores: Yup.object().required('Skor performa harus diisi'),
  total_score: Yup.number().required('Total skor harus diisi').min(0),
  max_score: Yup.number().required('Maksimal skor harus diisi').min(0),
  performance_level: Yup.string().required('Tingkat performa harus dipilih'),
});

const EvaluationForm: React.FC<EvaluationFormProps> = ({
  initialValues,
  rubric,
  onSubmit,
  onCancel,
  loading = false,
}) => {
  const performanceLevels = [
    { value: 'EXCELLENT', label: 'Sangat Baik', color: 'success' },
    { value: 'PROFICIENT', label: 'Baik', color: 'info' },
    { value: 'DEVELOPING', label: 'Sedang Berkembang', color: 'warning' },
    { value: 'BEGINNING', label: 'Perlu Bimbingan', color: 'error' },
  ];

  return (
    <Formik
      initialValues={{
        performance_scores: initialValues?.performance_scores || {},
        total_score: initialValues?.total_score || 0,
        max_score: initialValues?.max_score || 100,
        performance_level: initialValues?.performance_level || 'DEVELOPING',
        teacher_feedback: initialValues?.teacher_feedback || '',
      }}
      validationSchema={validationSchema}
      onSubmit={onSubmit}
    >
      {({ values, errors, touched, setFieldValue, isSubmitting }) => (
        <Form>
          <Card>
            <CardContent>
              <Typography variant="h6" gutterBottom>
                Form Evaluasi
              </Typography>

              <Grid container spacing={3}>
                {/* Performance Scores */}
                <Grid item xs={12}>
                  <Typography variant="subtitle2" gutterBottom>
                    Skor Performa
                  </Typography>
                  <TextField
                    fullWidth
                    multiline
                    rows={4}
                    label="Skor Performa (JSON)"
                    name="performance_scores"
                    value={JSON.stringify(values.performance_scores, null, 2)}
                    onChange={(e) => {
                      try {
                        setFieldValue('performance_scores', JSON.parse(e.target.value));
                      } catch {
                        // Invalid JSON, don't update
                      }
                    }}
                    error={touched.performance_scores && !!errors.performance_scores}
                    helperText={touched.performance_scores && (errors.performance_scores as string)}
                    disabled={loading}
                  />
                </Grid>

                {/* Total Score */}
                <Grid item xs={12} sm={6}>
                  <Typography variant="subtitle2" gutterBottom>
                    Total Skor: {values.total_score}
                  </Typography>
                  <Slider
                    value={values.total_score}
                    onChange={(_, value) => setFieldValue('total_score', value)}
                    min={0}
                    max={values.max_score}
                    valueLabelDisplay="auto"
                    disabled={loading}
                  />
                </Grid>

                {/* Max Score */}
                <Grid item xs={12} sm={6}>
                  <TextField
                    fullWidth
                    label="Maksimal Skor"
                    type="number"
                    name="max_score"
                    value={values.max_score}
                    onChange={(e) => setFieldValue('max_score', parseInt(e.target.value))}
                    disabled={loading}
                  />
                </Grid>

                {/* Performance Level */}
                <Grid item xs={12}>
                  <Typography variant="subtitle2" gutterBottom>
                    Tingkat Performa
                  </Typography>
                  <Box sx={{ display: 'flex', gap: 1, flexWrap: 'wrap' }}>
                    {performanceLevels.map((level) => (
                      <Chip
                        key={level.value}
                        label={level.label}
                        color={values.performance_level === level.value ? level.color : 'default'}
                        variant={values.performance_level === level.value ? 'filled' : 'outlined'}
                        onClick={() => setFieldValue('performance_level', level.value)}
                        sx={{ cursor: 'pointer' }}
                      />
                    ))}
                  </Box>
                </Grid>

                {/* Teacher Feedback */}
                <Grid item xs={12}>
                  <TeacherFeedback
                    value={values.teacher_feedback}
                    onChange={(value) => setFieldValue('teacher_feedback', value)}
                    disabled={loading}
                  />
                </Grid>

                {/* Actions */}
                <Grid item xs={12}>
                  <Box sx={{ display: 'flex', gap: 2, justifyContent: 'flex-end' }}>
                    {onCancel && (
                      <Button
                        variant="outlined"
                        onClick={onCancel}
                        disabled={isSubmitting || loading}
                      >
                        Batal
                      </Button>
                    )}
                    <Button
                      type="submit"
                      variant="contained"
                      disabled={isSubmitting || loading}
                      startIcon={isSubmitting || loading ? <CircularProgress size={20} /> : undefined}
                    >
                      {isSubmitting || loading ? 'Menyimpan...' : 'Simpan Evaluasi'}
                    </Button>
                  </Box>
                </Grid>
              </Grid>
            </CardContent>
          </Card>
        </Form>
      )}
    </Formik>
  );
};

export default EvaluationForm;

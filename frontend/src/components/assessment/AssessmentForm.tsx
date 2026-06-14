/**
 * Assessment Form Component
 * Form for creating/editing assessments
 */

import React, { useState } from 'react';
import {
  Box,
  Typography,
  TextField,
  Select,
  MenuItem,
  FormControl,
  InputLabel,
  Button,
  Card,
  CardContent,
  CircularProgress,
} from '@mui/material';
import { Formik, Form, Field, FormikHelpers } from 'formik';
import * as Yup from 'yup';
import TPSelector from './TPSelector';
import SuccessCriteriaSnapshot from './SuccessCriteriaSnapshot';
import { getTPById } from '@/api/tp';

interface AssessmentFormProps {
  initialValues?: {
    tp_id?: string;
    tp_version_no?: number;
    assessment_type?: string;
    assessment_items?: any;
    answer_key?: any;
    scoring_guidelines?: any;
  };
  onSubmit: (values: any, helpers: FormikHelpers<any>) => void;
  onCancel?: () => void;
  loading?: boolean;
}

const validationSchema = Yup.object().shape({
  tp_id: Yup.string().required('TP harus dipilih'),
  assessment_type: Yup.string().required('Tipe asesmen harus dipilih'),
  assessment_items: Yup.object().required('Item asesmen harus diisi'),
});

const AssessmentForm: React.FC<AssessmentFormProps> = ({
  initialValues,
  onSubmit,
  onCancel,
  loading = false,
}) => {
  const [tpSuccessCriteria, setTPSuccessCriteria] = useState<any>(null);

  const handleTPChange = (tpId: string) => {
    loadTPSuccessCriteria(tpId);
  };

  const loadTPSuccessCriteria = async (tpId: string) => {
    try {
      const tp = await getTPById(tpId);
      setTPSuccessCriteria(tp.success_criteria);
    } catch (error) {
      console.error('Error loading TP success criteria:', error);
    }
  };

  return (
    <Formik
      initialValues={{
        tp_id: initialValues?.tp_id || '',
        tp_version_no: initialValues?.tp_version_no || 1,
        success_criteria_snapshot: tpSuccessCriteria || {},
        assessment_type: initialValues?.assessment_type || 'FORMATIVE',
        assessment_items: initialValues?.assessment_items || {},
        answer_key: initialValues?.answer_key || {},
        scoring_guidelines: initialValues?.scoring_guidelines || {},
      }}
      validationSchema={validationSchema}
      onSubmit={onSubmit}
      enableReinitialize
    >
      {({ values, errors, touched, setFieldValue, isSubmitting }) => (
        <Form>
          <Card>
            <CardContent>
              <Typography variant="h6" gutterBottom>
                Form Asesmen
              </Typography>

              <Box sx={{ display: 'flex', flexDirection: 'column', gap: 3 }}>
                {/* TP Selection */}
                <Box sx={{ width: '100%' }}>
                  <TPSelector
                    value={values.tp_id}
                    onChange={(tpId) => {
                      setFieldValue('tp_id', tpId);
                      handleTPChange(tpId);
                    }}
                    error={touched.tp_id && !!errors.tp_id}
                    helperText={touched.tp_id && (errors.tp_id as string) || undefined}
                  />
                </Box>

                {/* Assessment Type */}
                <Box sx={{ width: { xs: '100%', sm: '50%' } }}>
                  <FormControl fullWidth error={touched.assessment_type && !!errors.assessment_type}>
                    <InputLabel>Tipe Asesmen</InputLabel>
                    <Field
                      as={Select}
                      name="assessment_type"
                      label="Tipe Asesmen"
                      disabled={loading}
                    >
                      <MenuItem value="FORMATIVE">Formatif</MenuItem>
                      <MenuItem value="SUMMATIVE">Sumatif</MenuItem>
                    </Field>
                  </FormControl>
                </Box>

                {/* TP Version */}
                <Box sx={{ width: { xs: '100%', sm: '50%' } }}>
                  <TextField
                    fullWidth
                    label="Versi TP"
                    type="number"
                    value={values.tp_version_no}
                    onChange={(e) => setFieldValue('tp_version_no', parseInt(e.target.value))}
                    disabled
                    helperText="Otomatis diambil dari TP yang dipilih"
                  />
                </Box>

                {/* Success Criteria Snapshot Preview */}
                {tpSuccessCriteria && (
                  <Box sx={{ width: '100%' }}>
                    <SuccessCriteriaSnapshot
                      snapshot={tpSuccessCriteria}
                      tpVersion={values.tp_version_no}
                      showWarning
                    />
                  </Box>
                )}

                {/* Assessment Items */}
                <Box sx={{ width: '100%' }}>
                  <TextField
                    fullWidth
                    multiline
                    rows={4}
                    label="Item Asesmen (JSON)"
                    name="assessment_items"
                    value={JSON.stringify(values.assessment_items, null, 2)}
                    onChange={(e) => {
                      try {
                        setFieldValue('assessment_items', JSON.parse(e.target.value));
                      } catch {
                        // Invalid JSON, don't update
                      }
                    }}
                    error={touched.assessment_items && !!errors.assessment_items}
                    helperText={touched.assessment_items && (errors.assessment_items as string)}
                    disabled={loading}
                  />
                </Box>

                {/* Answer Key */}
                <Box sx={{ width: '100%' }}>
                  <TextField
                    fullWidth
                    multiline
                    rows={4}
                    label="Kunci Jawaban (JSON)"
                    name="answer_key"
                    value={JSON.stringify(values.answer_key, null, 2)}
                    onChange={(e) => {
                      try {
                        setFieldValue('answer_key', JSON.parse(e.target.value));
                      } catch {
                        // Invalid JSON, don't update
                      }
                    }}
                    disabled={loading}
                  />
                </Box>

                {/* Scoring Guidelines */}
                <Box sx={{ width: '100%' }}>
                  <TextField
                    fullWidth
                    multiline
                    rows={4}
                    label="Pedoman Penilaian (JSON)"
                    name="scoring_guidelines"
                    value={JSON.stringify(values.scoring_guidelines, null, 2)}
                    onChange={(e) => {
                      try {
                        setFieldValue('scoring_guidelines', JSON.parse(e.target.value));
                      } catch {
                        // Invalid JSON, don't update
                      }
                    }}
                    disabled={loading}
                  />
                </Box>

                {/* Actions */}
                <Box sx={{ width: '100%' }}>
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
                      {isSubmitting || loading ? 'Menyimpan...' : 'Simpan'}
                    </Button>
                  </Box>
                </Box>
              </Box>
            </CardContent>
          </Card>
        </Form>
      )}
    </Formik>
  );
};

export default AssessmentForm;

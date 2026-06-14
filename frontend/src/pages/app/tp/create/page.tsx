/**
 * TP Create Page - MIGRATED TO TANSTACK QUERY
 * Create new Teaching Plan (TP)
 */

import React, { useState, useEffect } from 'react';
import {
  Box,
  Typography,
  Button,
  Card,
  CardContent,
  TextField,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  CircularProgress,
  Alert,
  Grid,
} from '@mui/material';
import {
  ArrowBack,
  Save,
} from '@mui/icons-material';
import { useNavigate, useLocation } from 'react-router-dom';
import { Formik, Form, FormikHelpers } from 'formik';
import * as Yup from 'yup';
import { useCreateTP } from '@/services/commands/TPCommandService';
import { useTPSets } from '@/services/queries/TPQueryService';
import { TPStatus } from '@/shared/types/domain';
import { CP } from '@/shared/types/domain';
import KKTPCriteriaForm from '@/components/kktp/KKTPCriteriaForm';
import { CPSelector } from '@/features/cp';

const validationSchema = Yup.object().shape({
  tp_set_id: Yup.string().required('Set TP harus dipilih'),
  sequence_number: Yup.number().required('Urutan TP harus diisi').min(1),
  cp_id: Yup.string().required('CP harus dipilih'),
  subject_id: Yup.string().required('Mata pelajaran harus dipilih'),
  phase_id: Yup.string().required('Fase harus dipilih'),
  element_id: Yup.string().required('Elemen harus dipilih'),
  title: Yup.string().required('Judul TP harus diisi'),
  learning_objectives: Yup.string().required('Tujuan pembelajaran harus diisi'),
  time_allocation: Yup.string().required('Alokasi waktu harus diisi'),
  estimated_weeks: Yup.number().required('Estimasi minggu harus diisi').min(1),
});

const TPCreatePage: React.FC = () => {
  const navigate = useNavigate();
  const location = useLocation();
  const [error, setError] = useState<string | null>(null);
  const [successCriteria, setSuccessCriteria] = useState<any>(null);
  const [selectedCP, setSelectedCP] = useState<CP | undefined>(undefined);

  // ✅ Using TanStack Query hook for TP Sets instead of manual state management
  const { 
    data: tpSets = [], 
  } = useTPSets();

  // Using mutation hook for creating TP
  const createMutation = useCreateTP({
    onSuccess: (newTP) => {
      navigate(`/tp/${newTP.id}`);
    },
    onError: (err) => {
      setError(err.message || 'Gagal membuat TP');
    },
  });

  // Check if CP was passed from navigation state
  useEffect(() => {
    if (location.state?.selectedCP) {
      setSelectedCP(location.state.selectedCP);
    }
  }, [location.state]);

  const handleSubmit = async (values: any, helpers: FormikHelpers<any>) => {
    setError(null);
    try {
      const payload = {
        ...values,
        success_criteria: successCriteria,
      };
      createMutation.mutate(payload);
    } catch (err: any) {
      setError(err.message || 'Gagal membuat TP');
      helpers.setSubmitting(false);
    }
  };

  return (
    <Box>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
        <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
          <Button
            variant="outlined"
            startIcon={<ArrowBack />}
            onClick={() => navigate('/tp')}
          >
            Kembali
          </Button>
          <Typography variant="h4">Buat TP Baru</Typography>
        </Box>
      </Box>

      {error && (
        <Alert severity="error" sx={{ mb: 3 }}>
          {error}
        </Alert>
      )}

      <Formik
        initialValues={{
          tp_set_id: '',
          sequence_number: 1,
          cp_id: '',
          subject_id: '',
          phase_id: '',
          element_id: '',
          subelement_id: '',
          title: '',
          learning_objectives: '',
          time_allocation: '',
          prerequisites: '',
          estimated_weeks: 1,
          status: 'DRAFT' as TPStatus,
        }}
        validationSchema={validationSchema}
        onSubmit={handleSubmit}
      >
        {({ values, errors, touched, setFieldValue, isSubmitting }) => (
          <Form>
            <Grid container spacing={3}>
              <Grid size={{ xs: 12, md: 8 }}>
                <Card sx={{ mb: 3 }}>
                  <CardContent>
                    <Typography variant="h6" gutterBottom>
                      Informasi TP
                    </Typography>
                    <Grid container spacing={2}>
                      <Grid size={{ xs: 12, sm: 6 }}>
                        <FormControl fullWidth error={touched.tp_set_id && !!errors.tp_set_id}>
                          <InputLabel>Set TP</InputLabel>
                          <Select
                            name="tp_set_id"
                            label="Set TP"
                            value={values.tp_set_id}
                            onChange={(e) => setFieldValue('tp_set_id', e.target.value)}
                          >
                            {tpSets.map((tpSet) => (
                              <MenuItem key={tpSet.id} value={tpSet.id}>
                                {tpSet.name}
                              </MenuItem>
                            ))}
                          </Select>
                        </FormControl>
                      </Grid>
                      <Grid size={{ xs: 12, sm: 6 }}>
                        <TextField
                          fullWidth
                          label="Urutan TP"
                          type="number"
                          name="sequence_number"
                          value={values.sequence_number}
                          onChange={(e) => setFieldValue('sequence_number', parseInt(e.target.value))}
                          error={touched.sequence_number && !!errors.sequence_number}
                          helperText={touched.sequence_number && (errors.sequence_number as string)}
                        />
                      </Grid>
                      <Grid size={{ xs: 12 }}>
                        <CPSelector
                          selectedCP={selectedCP}
                          onSelect={(cp) => {
                            setSelectedCP(cp);
                            // Auto-populate curriculum hierarchy fields
                            setFieldValue('cp_id', cp.id);
                            setFieldValue('subject_id', cp.subject_id);
                            setFieldValue('phase_id', cp.phase_id);
                            setFieldValue('element_id', cp.element_id);
                            setFieldValue('subelement_id', cp.subelement_id || '');
                          }}
                        />
                      </Grid>
                      <Grid size={{ xs: 12, sm: 6 }}>
                        <TextField
                          fullWidth
                          label="ID Mata Pelajaran"
                          name="subject_id"
                          value={values.subject_id}
                          onChange={(e) => setFieldValue('subject_id', e.target.value)}
                          error={touched.subject_id && !!errors.subject_id}
                          helperText={touched.subject_id && (errors.subject_id as string)}
                          disabled={!!selectedCP} // Auto-populated from CP
                        />
                      </Grid>
                      <Grid size={{ xs: 12, sm: 6 }}>
                        <TextField
                          fullWidth
                          label="ID Fase"
                          name="phase_id"
                          value={values.phase_id}
                          onChange={(e) => setFieldValue('phase_id', e.target.value)}
                          error={touched.phase_id && !!errors.phase_id}
                          helperText={touched.phase_id && (errors.phase_id as string)}
                          disabled={!!selectedCP} // Auto-populated from CP
                        />
                      </Grid>
                      <Grid size={{ xs: 12, sm: 6 }}>
                        <TextField
                          fullWidth
                          label="ID Elemen"
                          name="element_id"
                          value={values.element_id}
                          onChange={(e) => setFieldValue('element_id', e.target.value)}
                          error={touched.element_id && !!errors.element_id}
                          helperText={touched.element_id && (errors.element_id as string)}
                          disabled={!!selectedCP} // Auto-populated from CP
                        />
                      </Grid>
                      <Grid size={{ xs: 12 }}>
                        <TextField
                          fullWidth
                          label="ID Sub-elemen (Opsional)"
                          name="subelement_id"
                          value={values.subelement_id}
                          onChange={(e) => setFieldValue('subelement_id', e.target.value)}
                          disabled={!!selectedCP} // Auto-populated from CP
                        />
                      </Grid>
                    </Grid>
                  </CardContent>
                </Card>

                <Card sx={{ mb: 3 }}>
                  <CardContent>
                    <Typography variant="h6" gutterBottom>
                      Detail TP
                    </Typography>
                    <Grid container spacing={2}>
                      <Grid size={{ xs: 12 }}>
                        <TextField
                          fullWidth
                          label="Judul TP"
                          name="title"
                          value={values.title}
                          onChange={(e) => setFieldValue('title', e.target.value)}
                          error={touched.title && !!errors.title}
                          helperText={touched.title && (errors.title as string)}
                        />
                      </Grid>
                      <Grid size={{ xs: 12 }}>
                        <TextField
                          fullWidth
                          multiline
                          rows={4}
                          label="Tujuan Pembelajaran"
                          name="learning_objectives"
                          value={values.learning_objectives}
                          onChange={(e) => setFieldValue('learning_objectives', e.target.value)}
                          error={touched.learning_objectives && !!errors.learning_objectives}
                          helperText={touched.learning_objectives && (errors.learning_objectives as string)}
                        />
                      </Grid>
                      <Grid size={{ xs: 12 }}>
                        <TextField
                          fullWidth
                          label="Alokasi Waktu"
                          name="time_allocation"
                          value={values.time_allocation}
                          onChange={(e) => setFieldValue('time_allocation', e.target.value)}
                          error={touched.time_allocation && !!errors.time_allocation}
                          helperText={touched.time_allocation && (errors.time_allocation as string)}
                        />
                      </Grid>
                      <Grid size={{ xs: 12 }}>
                        <TextField
                          fullWidth
                          multiline
                          rows={3}
                          label="Prasyarat (Opsional)"
                          name="prerequisites"
                          value={values.prerequisites}
                          onChange={(e) => setFieldValue('prerequisites', e.target.value)}
                        />
                      </Grid>
                      <Grid size={{ xs: 12, sm: 6 }}>
                        <TextField
                          fullWidth
                          label="Estimasi Minggu"
                          type="number"
                          name="estimated_weeks"
                          value={values.estimated_weeks}
                          onChange={(e) => setFieldValue('estimated_weeks', parseInt(e.target.value))}
                          error={touched.estimated_weeks && !!errors.estimated_weeks}
                          helperText={touched.estimated_weeks && (errors.estimated_weeks as string)}
                        />
                      </Grid>
                    </Grid>
                  </CardContent>
                </Card>
              </Grid>

              <Grid size={{ xs: 12, md: 4 }}>
                <Card sx={{ mb: 3 }}>
                  <CardContent>
                    <Typography variant="h6" gutterBottom>
                      Kriteria Ketuntasan (KKTP)
                    </Typography>
                    <KKTPCriteriaForm
                      data={successCriteria}
                      onChange={setSuccessCriteria}
                    />
                  </CardContent>
                </Card>

                <Card>
                  <CardContent>
                    <Typography variant="h6" gutterBottom>
                      Status
                    </Typography>
                    <FormControl fullWidth>
                      <Select
                        name="status"
                        value={values.status}
                        label="Status"
                        onChange={(e) => setFieldValue('status', e.target.value)}
                      >
                        <MenuItem value="DRAFT">Draft</MenuItem>
                        <MenuItem value="UNDER_REVIEW">Dalam Review</MenuItem>
                      </Select>
                    </FormControl>
                  </CardContent>
                </Card>
              </Grid>

              <Grid size={{ xs: 12 }}>
                <Box sx={{ display: 'flex', gap: 2, justifyContent: 'flex-end' }}>
                  <Button
                    variant="outlined"
                    onClick={() => navigate('/tp')}
                    disabled={isSubmitting || createMutation.isPending}
                  >
                    Batal
                  </Button>
                  <Button
                    type="submit"
                    variant="contained"
                    startIcon={isSubmitting || createMutation.isPending ? <CircularProgress size={20} /> : <Save />}
                    disabled={isSubmitting || createMutation.isPending}
                  >
                    {isSubmitting || createMutation.isPending ? 'Menyimpan...' : 'Simpan'}
                  </Button>
                </Box>
              </Grid>
            </Grid>
          </Form>
        )}
      </Formik>
    </Box>
  );
};

export default TPCreatePage;

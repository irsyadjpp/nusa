/**
 * TP Edit Page - MIGRATED TO TANSTACK QUERY
 * Edit existing Teaching Plan (TP)
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
import { useNavigate, useParams } from 'react-router-dom';
import { Formik, Form, FormikHelpers } from 'formik';
import * as Yup from 'yup';
import { useTP, useTPSets } from '@/services/queries/TPQueryService';
import { useUpdateTP } from '@/services/commands/TPCommandService';
import KKTPCriteriaForm from '@/components/kktp/KKTPCriteriaForm';

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

const TPEditPage: React.FC = () => {
  const navigate = useNavigate();
  const { id } = useParams<{ id: string }>();
  const [error, setError] = useState<string | null>(null);
  const [successCriteria, setSuccessCriteria] = useState<any>(null);

  // ✅ Using TanStack Query hooks instead of manual state management
  const { 
    data: tp, 
    isLoading 
  } = useTP(id!);

  const { 
    data: tpSets = [],
    isLoading: tpSetsLoading 
  } = useTPSets();

  // Using mutation hook for updating TP
  const updateMutation = useUpdateTP({
    onSuccess: () => {
      navigate(`/tp/${id}`);
    },
    onError: (err) => {
      setError(err.message || 'Gagal mengupdate TP');
    },
  });

  useEffect(() => {
    if (tp) {
      setSuccessCriteria(tp.success_criteria);
    }
  }, [tp]);

  const handleSubmit = async (values: any, helpers: FormikHelpers<any>) => {
    if (!id) return;
    setError(null);
    try {
      const payload = {
        ...values,
        success_criteria: successCriteria,
      };
      updateMutation.mutate({ id, data: payload });
    } catch (err: any) {
      setError(err.message || 'Gagal mengupdate TP');
      helpers.setSubmitting(false);
    }
  };

  if (isLoading || tpSetsLoading) {
    return (
      <Box sx={{ display: 'flex', justifyContent: 'center', py: 4 }}>
        <CircularProgress />
      </Box>
    );
  }

  if (error || !tp) {
    return (
      <Alert severity="error">
        {error || 'TP tidak ditemukan'}
      </Alert>
    );
  }

  return (
    <Box>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
        <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
          <Button
            variant="outlined"
            startIcon={<ArrowBack />}
            onClick={() => navigate(`/tp/${id}`)}
          >
            Kembali
          </Button>
          <Typography variant="h4">Edit TP</Typography>
        </Box>
      </Box>

      {error && (
        <Alert severity="error" sx={{ mb: 3 }}>
          {error}
        </Alert>
      )}

      <Formik
        initialValues={{
          tp_set_id: tp.tp_set_id,
          sequence_number: tp.sequence_number,
          cp_id: tp.cp_id,
          subject_id: tp.subject_id,
          phase_id: tp.phase_id,
          element_id: tp.element_id,
          subelement_id: tp.subelement_id || '',
          title: tp.title,
          learning_objectives: tp.learning_objectives,
          time_allocation: tp.time_allocation,
          prerequisites: tp.prerequisites || '',
          estimated_weeks: tp.estimated_weeks,
          status: tp.status,
        }}
        validationSchema={validationSchema}
        onSubmit={handleSubmit}
        enableReinitialize
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
                      <Grid size={{ xs: 12, sm: 6 }}>
                        <TextField
                          fullWidth
                          label="ID CP"
                          name="cp_id"
                          value={values.cp_id}
                          onChange={(e) => setFieldValue('cp_id', e.target.value)}
                          error={touched.cp_id && !!errors.cp_id}
                          helperText={touched.cp_id && (errors.cp_id as string)}
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
                        />
                      </Grid>
                      <Grid size={{ xs: 12 }}>
                        <TextField
                          fullWidth
                          label="ID Sub-elemen (Opsional)"
                          name="subelement_id"
                          value={values.subelement_id}
                          onChange={(e) => setFieldValue('subelement_id', e.target.value)}
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
                      <InputLabel>Status</InputLabel>
                      <Select
                        name="status"
                        label="Status"
                        value={values.status}
                        onChange={(e) => setFieldValue('status', e.target.value)}
                      >
                        <MenuItem value="DRAFT">Draft</MenuItem>
                        <MenuItem value="UNDER_REVIEW">Dalam Review</MenuItem>
                        <MenuItem value="APPROVED">Disetujui</MenuItem>
                        <MenuItem value="REJECTED">Ditolak</MenuItem>
                        <MenuItem value="ARCHIVED">Diarsipkan</MenuItem>
                      </Select>
                    </FormControl>
                  </CardContent>
                </Card>
              </Grid>

              <Grid size={{ xs: 12 }}>
                <Box sx={{ display: 'flex', gap: 2, justifyContent: 'flex-end' }}>
                  <Button
                    variant="outlined"
                    onClick={() => navigate(`/tp/${id}`)}
                    disabled={isSubmitting || updateMutation.isPending}
                  >
                    Batal
                  </Button>
                  <Button
                    type="submit"
                    variant="contained"
                    startIcon={isSubmitting || updateMutation.isPending ? <CircularProgress size={20} /> : <Save />}
                    disabled={isSubmitting || updateMutation.isPending}
                  >
                    {isSubmitting || updateMutation.isPending ? 'Menyimpan...' : 'Simpan'}
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

export default TPEditPage;

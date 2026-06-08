/**
 * Evidence Upload Page
 * Upload new Evidence
 */

import React, { useState } from 'react';
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
import { useNavigate } from 'react-router-dom';
import { Formik, Form, FormikHelpers } from 'formik';
import * as Yup from 'yup';
import { createEvidence } from '@/api/evidence';
import EvidenceUpload from '@/components/evidence/EvidenceUpload';

const validationSchema = Yup.object().shape({
  student_id: Yup.string().required('Siswa harus dipilih'),
  assessment_id: Yup.string().required('Asesmen harus dipilih'),
  evidence_type: Yup.string().required('Tipe bukti harus dipilih'),
});

const EvidenceUploadPage: React.FC = () => {
  const navigate = useNavigate();
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [uploadedFiles, setUploadedFiles] = useState<File[]>([]);

  const handleFileUpload = (files: File[]) => {
    setUploadedFiles(files);
  };

  const handleSubmit = async (values: any, helpers: FormikHelpers<any>) => {
    setLoading(true);
    setError(null);
    try {
      const formData = new FormData();
      formData.append('student_id', values.student_id);
      formData.append('assessment_id', values.assessment_id);
      formData.append('evidence_type', values.evidence_type);
      formData.append('user_id', values.user_id || '');
      uploadedFiles.forEach((file) => {
        formData.append('files', file);
      });

      const newEvidence = await createEvidence(formData);
      navigate(`/evidence/${newEvidence.id}`);
    } catch (err: any) {
      setError(err.message || 'Gagal mengupload bukti');
      helpers.setSubmitting(false);
    } finally {
      setLoading(false);
    }
  };

  return (
    <Box>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
        <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
          <Button
            variant="outlined"
            startIcon={<ArrowBack />}
            onClick={() => navigate('/evidence')}
          >
            Kembali
          </Button>
          <Typography variant="h4">Upload Bukti Baru</Typography>
        </Box>
      </Box>

      {error && (
        <Alert severity="error" sx={{ mb: 3 }}>
          {error}
        </Alert>
      )}

      <Formik
        initialValues={{
          student_id: '',
          assessment_id: '',
          evidence_type: 'STUDENT_WORK',
          user_id: '',
        }}
        validationSchema={validationSchema}
        onSubmit={handleSubmit}
      >
        {({ values, errors, touched, setFieldValue, isSubmitting }) => (
          <Form>
            <Grid container spacing={3}>
              <Grid item xs={12} md={8}>
                <Card sx={{ mb: 3 }}>
                  <CardContent>
                    <Typography variant="h6" gutterBottom>
                      Informasi Bukti
                    </Typography>
                    <Grid container spacing={2}>
                      <Grid item xs={12} sm={6}>
                        <TextField
                          fullWidth
                          label="ID Siswa"
                          name="student_id"
                          value={values.student_id}
                          onChange={(e) => setFieldValue('student_id', e.target.value)}
                          error={touched.student_id && !!errors.student_id}
                          helperText={touched.student_id && (errors.student_id as string)}
                        />
                      </Grid>
                      <Grid item xs={12} sm={6}>
                        <TextField
                          fullWidth
                          label="ID Asesmen"
                          name="assessment_id"
                          value={values.assessment_id}
                          onChange={(e) => setFieldValue('assessment_id', e.target.value)}
                          error={touched.assessment_id && !!errors.assessment_id}
                          helperText={touched.assessment_id && (errors.assessment_id as string)}
                        />
                      </Grid>
                      <Grid item xs={12} sm={6}>
                        <FormControl fullWidth error={touched.evidence_type && !!errors.evidence_type}>
                          <InputLabel>Tipe Bukti</InputLabel>
                          <Select
                            name="evidence_type"
                            label="Tipe Bukti"
                            value={values.evidence_type}
                            onChange={(e) => setFieldValue('evidence_type', e.target.value)}
                          >
                            <MenuItem value="STUDENT_WORK">Karya Siswa</MenuItem>
                            <MenuItem value="ASSESSMENT_RESULT">Hasil Asesmen</MenuItem>
                            <MenuItem value="OBSERVATION">Observasi</MenuItem>
                          </Select>
                        </FormControl>
                      </Grid>
                      <Grid item xs={12} sm={6}>
                        <TextField
                          fullWidth
                          label="ID User (Guru)"
                          name="user_id"
                          value={values.user_id}
                          onChange={(e) => setFieldValue('user_id', e.target.value)}
                        />
                      </Grid>
                    </Grid>
                  </CardContent>
                </Card>

                <Card>
                  <CardContent>
                    <Typography variant="h6" gutterBottom>
                      File Bukti
                    </Typography>
                    <EvidenceUpload
                      onUpload={handleFileUpload}
                      disabled={loading}
                    />
                  </CardContent>
                </Card>
              </Grid>

              <Grid item xs={12} md={4}>
                <Card>
                  <CardContent>
                    <Typography variant="h6" gutterBottom>
                      Status
                    </Typography>
                    <Typography variant="body2" color="text.secondary">
                      Bukti akan diset ke status "SUBMITTED" secara otomatis setelah upload.
                    </Typography>
                  </CardContent>
                </Card>
              </Grid>

              <Grid item xs={12}>
                <Box sx={{ display: 'flex', gap: 2, justifyContent: 'flex-end' }}>
                  <Button
                    variant="outlined"
                    onClick={() => navigate('/evidence')}
                    disabled={isSubmitting || loading}
                  >
                    Batal
                  </Button>
                  <Button
                    type="submit"
                    variant="contained"
                    startIcon={isSubmitting || loading ? <CircularProgress size={20} /> : <Save />}
                    disabled={isSubmitting || loading || uploadedFiles.length === 0}
                  >
                    {isSubmitting || loading ? 'Mengupload...' : 'Upload'}
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

export default EvidenceUploadPage;

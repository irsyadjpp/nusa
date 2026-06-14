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
} from '@mui/material';
import {
  ArrowBack,
  Save,
} from '@mui/icons-material';
import { useNavigate } from 'react-router-dom';
import { Formik, Form, FormikHelpers } from 'formik';
import * as Yup from 'yup';
import { uploadEvidence } from '@/api/evidence';
import { useAuth } from '@/features/auth';
// import EvidenceUpload from '@/components/evidence/EvidenceUpload'; // TODO: Implement EvidenceUpload component

const validationSchema = Yup.object().shape({
  student_id: Yup.string().required('Siswa harus dipilih'),
  assessment_id: Yup.string().required('Asesmen harus dipilih'),
  evidence_type: Yup.string().required('Tipe bukti harus dipilih'),
});

const EvidenceUploadPage: React.FC = () => {
  const navigate = useNavigate();
  const { user } = useAuth();
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [uploadedFiles, setUploadedFiles] = useState<File[]>([]);

  const handleFileUpload = (files: File[]) => {
    setUploadedFiles(files);
  };

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const files = event.target.files;
    if (files) {
      handleFileUpload(Array.from(files));
    }
  };

  const handleSubmit = async (values: any, helpers: FormikHelpers<any>) => {
    if (!user?.id) return;
    setLoading(true);
    setError(null);
    try {
      // For now, just use a placeholder URL since we don't have actual file upload
      const fileUrl = uploadedFiles.length > 0 ? `https://placeholder.url/${uploadedFiles[0].name}` : '';
      const fileMetadata = {
        filename: uploadedFiles.length > 0 ? uploadedFiles[0].name : 'placeholder.jpg',
        file_size: uploadedFiles.length > 0 ? uploadedFiles[0].size : 0,
        mime_type: uploadedFiles.length > 0 ? uploadedFiles[0].type : 'image/jpeg',
        file_format: uploadedFiles.length > 0 ? uploadedFiles[0].type.split('/')[1] || 'jpg' : 'jpg',
      };

      const evidenceData = {
        student_id: values.student_id,
        assessment_id: values.assessment_id,
        evidence_type: values.evidence_type,
        file_url: fileUrl,
        file_metadata: fileMetadata,
      };

      const newEvidence = await uploadEvidence(evidenceData, user.id);
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
            <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 3 }}>
              <Box sx={{ width: { xs: '100%', md: '66.67%' } }}>
                <Card sx={{ mb: 3 }}>
                  <CardContent>
                    <Typography variant="h6" gutterBottom>
                      Informasi Bukti
                    </Typography>
                    <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 2 }}>
                      <Box sx={{ width: { xs: '100%', sm: '50%' } }}>
                        <TextField
                          fullWidth
                          label="ID Siswa"
                          name="student_id"
                          value={values.student_id}
                          onChange={(e) => setFieldValue('student_id', e.target.value)}
                          error={touched.student_id && !!errors.student_id}
                          helperText={touched.student_id && (errors.student_id as string)}
                        />
                      </Box>
                      <Box sx={{ width: { xs: '100%', sm: '50%' } }}>
                        <TextField
                          fullWidth
                          label="ID Asesmen"
                          name="assessment_id"
                          value={values.assessment_id}
                          onChange={(e) => setFieldValue('assessment_id', e.target.value)}
                          error={touched.assessment_id && !!errors.assessment_id}
                          helperText={touched.assessment_id && (errors.assessment_id as string)}
                        />
                      </Box>
                      <Box sx={{ width: { xs: '100%', sm: '50%' } }}>
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
                      </Box>
                      <Box sx={{ width: { xs: '100%', sm: '50%' }}}>
                        <TextField
                          fullWidth
                          label="ID User (Guru)"
                          name="user_id"
                          value={values.user_id}
                          onChange={(e) => setFieldValue('user_id', e.target.value)}
                        />
                      </Box>
                    </Box>
                  </CardContent>
                </Card>

                <Card>
                  <CardContent>
                    <Typography variant="h6" gutterBottom>
                      File Bukti
                    </Typography>
                    <Box sx={{ border: '2px dashed #ccc', borderRadius: 1, p: 3, textAlign: 'center' }}>
                      <input
                        type="file"
                        onChange={handleFileChange}
                        disabled={loading}
                        accept=".pdf,.doc,.docx,.jpg,.jpeg,.png"
                      />
                      <Typography variant="caption" sx={{ mt: 1, display: 'block' }}>
                        PDF, DOC, DOCX, JPG, PNG (max 10MB)
                      </Typography>
                    </Box>
                  </CardContent>
                </Card>
              </Box>

              <Box sx={{ width: { xs: '100%', md: '33.33%' } }}>
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
              </Box>

              <Box sx={{ width: { xs: '100%' } }}>
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
              </Box>
            </Box>
          </Form>
        )}
      </Formik>
    </Box>
  );
};

export default EvidenceUploadPage;

/**
 * Report Generate Page
 * Generate new Narrative Report
 */

import React, { useState } from 'react';
import {
  Box,
  Typography,
  Button,
  Card,
  CardContent,
  TextField,
  CircularProgress,
  Alert,
} from '@mui/material';
import {
  ArrowBack,
  Save,
  Visibility,
} from '@mui/icons-material';
import { useNavigate } from 'react-router-dom';
import { Formik, Form, FormikHelpers } from 'formik';
import * as Yup from 'yup';
import NarrativeReportEditor from '@/components/report/NarrativeReportEditor';
import AchievementSummary from '@/components/report/AchievementSummary';
import ReportPreview from '@/components/report/ReportPreview';

const validationSchema = Yup.object().shape({
  student_id: Yup.string().required('Siswa harus dipilih'),
  period_start: Yup.string().required('Tanggal mulai periode harus diisi'),
  period_end: Yup.string().required('Tanggal akhir periode harus diisi'),
});

const ReportGeneratePage: React.FC = () => {
  const navigate = useNavigate();
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [narrativeContent, setNarrativeContent] = useState('');
  const [achievementData, setAchievementData] = useState<any>(null);
  const [showPreview, setShowPreview] = useState(false);

  const handleGenerateAchievement = async (studentId: string, periodStart: string, periodEnd: string) => {
    try {
      const response = await fetch(`/api/achievements/summary?student_id=${studentId}&period_start=${periodStart}&period_end=${periodEnd}`);
      const data = await response.json();
      setAchievementData(data);
    } catch (err: any) {
      setError(err.message || 'Gagal mengenerate ringkasan pencapaian');
    }
  };

  const handleSubmit = async (values: any, helpers: FormikHelpers<any>) => {
    setLoading(true);
    setError(null);
    try {
      const payload = {
        ...values,
        narrative_content: narrativeContent,
      };
      const response = await fetch('/api/reports', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(payload),
      });
      const newReport = await response.json();
      navigate(`/reports/${newReport.id}`);
    } catch (err: any) {
      setError(err.message || 'Gagal membuat rapor');
      helpers.setSubmitting(false);
    } finally {
      setLoading(false);
    }
  };

  const handlePrint = () => {
    window.print();
  };

  return (
    <Box>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
        <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
          <Button
            variant="outlined"
            startIcon={<ArrowBack />}
            onClick={() => navigate('/reports')}
          >
            Kembali
          </Button>
          <Typography variant="h4">Buat Rapor Naratif</Typography>
        </Box>
        <Button
          variant="outlined"
          startIcon={<Visibility />}
          onClick={() => setShowPreview(!showPreview)}
        >
          {showPreview ? 'Edit' : 'Preview'}
        </Button>
      </Box>

      {error && (
        <Alert severity="error" sx={{ mb: 3 }}>
          {error}
        </Alert>
      )}

      {!showPreview ? (
        <Formik
          initialValues={{
            student_id: '',
            period_start: new Date().toISOString().split('T')[0],
            period_end: new Date().toISOString().split('T')[0],
          }}
          validationSchema={validationSchema}
          onSubmit={handleSubmit}
        >
          {({ values, errors, touched, setFieldValue, isSubmitting }) => (
            <Form>
              <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 3 }}>
                <Box sx={{ width: { xs: '100%', md: '33.33%' }}}>
                  <Card sx={{ mb: 3 }}>
                    <CardContent>
                      <Typography variant="h6" gutterBottom>
                        Informasi Rapor
                      </Typography>
                      <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
                        <TextField
                          fullWidth
                          label="ID Siswa"
                          name="student_id"
                          value={values.student_id}
                          onChange={(e) => setFieldValue('student_id', e.target.value)}
                          error={touched.student_id && !!errors.student_id}
                          helperText={touched.student_id && (errors.student_id as string)}
                        />
                        <TextField
                          fullWidth
                          label="Tanggal Mulai Periode"
                          type="date"
                          name="period_start"
                          value={values.period_start}
                          onChange={(e) => setFieldValue('period_start', e.target.value)}
                          error={touched.period_start && !!errors.period_start}
                          helperText={touched.period_start && (errors.period_start as string)}
                          InputLabelProps={{ shrink: true }}
                        />
                        <TextField
                          fullWidth
                          label="Tanggal Akhir Periode"
                          type="date"
                          name="period_end"
                          value={values.period_end}
                          onChange={(e) => setFieldValue('period_end', e.target.value)}
                          error={touched.period_end && !!errors.period_end}
                          helperText={touched.period_end && (errors.period_end as string)}
                          InputLabelProps={{ shrink: true }}
                        />
                        <Button
                          fullWidth
                          variant="outlined"
                          onClick={() => handleGenerateAchievement(values.student_id, values.period_start, values.period_end)}
                          disabled={!values.student_id}
                        >
                            Generate Ringkasan
                          </Button>
                      </Box>
                    </CardContent>
                  </Card>
                </Box>

                <Box sx={{ width: { xs: '100%', md: '66.67%' }}}>
                  {achievementData && (
                    <Card sx={{ mb: 3 }}>
                      <CardContent>
                        <Typography variant="h6" gutterBottom>
                          Ringkasan Pencapaian
                        </Typography>
                        <AchievementSummary data={achievementData} />
                      </CardContent>
                    </Card>
                  )}

                  <Card>
                    <CardContent>
                      <Typography variant="h6" gutterBottom>
                        Narasi Guru
                      </Typography>
                      <NarrativeReportEditor
                        value={narrativeContent}
                        onChange={setNarrativeContent}
                        onSave={() => {}}
                        disabled={loading}
                      />
                    </CardContent>
                  </Card>
                </Box>

                <Box sx={{ width: '100%' }}>
                  <Box sx={{ display: 'flex', gap: 2, justifyContent: 'flex-end' }}>
                    <Button
                      variant="outlined"
                      onClick={() => navigate('/reports')}
                      disabled={isSubmitting || loading}
                    >
                      Batal
                    </Button>
                    <Button
                      type="submit"
                      variant="contained"
                      startIcon={isSubmitting || loading ? <CircularProgress size={20} /> : <Save />}
                      disabled={isSubmitting || loading}
                    >
                      {isSubmitting || loading ? 'Menyimpan...' : 'Simpan Draft'}
                    </Button>
                  </Box>
                </Box>
              </Box>
            </Form>
          )}
        </Formik>
      ) : (
        <ReportPreview
          studentName={achievementData?.student_name || 'Siswa'}
          periodStart={achievementData?.period_start || ''}
          periodEnd={achievementData?.period_end || ''}
          narrativeContent={narrativeContent}
          achievementSummary={achievementData}
          onPrint={handlePrint}
        />
      )}
    </Box>
  );
};

export default ReportGeneratePage;

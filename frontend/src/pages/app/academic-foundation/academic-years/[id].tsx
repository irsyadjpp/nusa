/**
 * Academic Year Form Page
 * Separate page for creating and editing academic years
 */

import React, { useState, useEffect } from 'react';
import {
  Box,
  Typography,
  Button,
  Card,
  CardContent,
  TextField,
  Alert,
  CircularProgress,
  Container,
} from '@mui/material';
import {
  ArrowBack as ArrowBackIcon,
  Save as SaveIcon,
} from '@mui/icons-material';
import { useNavigate, useParams } from 'react-router-dom';
import { CreateAcademicYearRequest } from '@/api/academic-foundation';
import { useAcademicYears } from '@/services/queries/AcademicFoundationQueryService';
import * as academicFoundationApi from '@/api/academic-foundation';

const AcademicYearFormPage: React.FC = () => {
  const navigate = useNavigate();
  const { id } = useParams<{ id?: string }>();

  const [formData, setFormData] = useState<CreateAcademicYearRequest>({
    school_id: 'default-school-id', // TODO: Get from auth context
    name: '',
    start_date: '',
    end_date: '',
  });
  const [formError, setFormError] = useState('');
  const [submitting, setSubmitting] = useState(false);

  const { data: academicYears, isLoading } = useAcademicYears({
    school_id: 'default-school-id'
  });

  // Load academic year data for editing
  useEffect(() => {
    if (id && academicYears) {
      const academicYear = academicYears.find((ay: any) => ay.id === id);
      if (academicYear) {
        setFormData({
          name: academicYear.name,
          start_date: academicYear.start_date,
          end_date: academicYear.end_date,
          school_id: academicYear.school_id,
        });
      }
    }
  }, [id, academicYears]);

  const validateForm = (): boolean => {
    if (!formData.name.trim()) {
      setFormError('Nama tahun ajaran harus diisi');
      return false;
    }
    if (!formData.start_date) {
      setFormError('Tanggal mulai harus diisi');
      return false;
    }
    if (!formData.end_date) {
      setFormError('Tanggal selesai harus diisi');
      return false;
    }
    if (new Date(formData.start_date) >= new Date(formData.end_date)) {
      setFormError('Tanggal selesai harus setelah tanggal mulai');
      return false;
    }
    return true;
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!validateForm()) {
      return;
    }

    setSubmitting(true);
    setFormError('');

    try {
      if (id) {
        // Update existing academic year
        await academicFoundationApi.updateAcademicYear(id, {
          name: formData.name,
          start_date: formData.start_date,
          end_date: formData.end_date,
        });
      } else {
        // Create new academic year
        await academicFoundationApi.createAcademicYear(formData);
      }
      navigate('/dashboard/academic-foundation/academic-years');
    } catch (error: any) {
      setFormError(error.response?.data?.message || 'Gagal menyimpan tahun ajaran');
    } finally {
      setSubmitting(false);
    }
  };

  if (isLoading) {
    return (
      <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
        <Box sx={{ display: 'flex', justifyContent: 'center', py: 8 }}>
          <CircularProgress />
        </Box>
      </Container>
    );
  }

  return (
    <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
      <Box sx={{ mb: 3 }}>
        <Button
          startIcon={<ArrowBackIcon />}
          onClick={() => navigate('/dashboard/academic-foundation/academic-years')}
          sx={{ mb: 2 }}
        >
          Kembali ke Daftar Tahun Ajaran
        </Button>
        <Typography variant="h4" component="h1">
          {id ? 'Edit Tahun Ajaran' : 'Tambah Tahun Ajaran Baru'}
        </Typography>
      </Box>

      <Card>
        <CardContent>
          <form onSubmit={handleSubmit}>
            <Box sx={{ display: 'flex', flexDirection: 'column', gap: 3 }}>
              {formError && (
                <Alert severity="error" onClose={() => setFormError('')}>
                  {formError}
                </Alert>
              )}

              <TextField
                label="Nama Tahun Ajaran"
                fullWidth
                value={formData.name}
                onChange={(e) => setFormData({ ...formData, name: e.target.value })}
                placeholder="Contoh: 2024/2025"
                error={!formData.name.trim()}
                helperText={!formData.name.trim() ? 'Nama tahun ajaran harus diisi' : ''}
                required
              />

              <TextField
                label="Tanggal Mulai"
                fullWidth
                type="date"
                value={formData.start_date}
                onChange={(e) => setFormData({ ...formData, start_date: e.target.value })}
                InputLabelProps={{
                  shrink: true,
                }}
                error={!formData.start_date}
                helperText={!formData.start_date ? 'Tanggal mulai harus diisi' : ''}
                required
              />
              
              <TextField
                label="Tanggal Selesai"
                fullWidth
                type="date"
                value={formData.end_date}
                onChange={(e) => setFormData({ ...formData, end_date: e.target.value })}
                InputLabelProps={{
                  shrink: true,
                }}
                error={!formData.end_date}
                helperText={!formData.end_date ? 'Tanggal selesai harus diisi' : ''}
                required
              />

              <Box sx={{ display: 'flex', gap: 2, justifyContent: 'flex-end', mt: 2 }}>
                <Button
                  variant="outlined"
                  onClick={() => navigate('/dashboard/academic-foundation/academic-years')}
                >
                  Batal
                </Button>
                <Button
                  type="submit"
                  variant="contained"
                  startIcon={<SaveIcon />}
                  disabled={submitting}
                >
                  {submitting ? 'Menyimpan...' : 'Simpan'}
                </Button>
              </Box>
            </Box>
          </form>
        </CardContent>
      </Card>
    </Container>
  );
};

export default AcademicYearFormPage;

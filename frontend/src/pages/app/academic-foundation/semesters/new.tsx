/**
 * Semester Form Page
 * Separate page for creating and editing semesters
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
import { useNavigate, useParams, useSearchParams } from 'react-router-dom';
import { CreateSemesterRequest } from '@/api/academic-foundation';
import { useSemesters } from '@/services/queries/AcademicFoundationQueryService';
import * as academicFoundationApi from '@/api/academic-foundation';

const SemesterFormPage: React.FC = () => {
  const navigate = useNavigate();
  const { id } = useParams<{ id?: string }>();
  const [searchParams] = useSearchParams();
  const academicYearId = searchParams.get('academicYearId') || '';

  const [formData, setFormData] = useState<CreateSemesterRequest>({
    academic_year_id: academicYearId,
    type: 'REGULAR',
    name: '',
    sequence_number: 1,
    sequence: 1,
    start_date: '',
    end_date: '',
    is_active: true,
  });
  const [formError, setFormError] = useState('');
  const [submitting, setSubmitting] = useState(false);

  const { data: semesters, isLoading } = useSemesters({
    academic_year_id: academicYearId
  });

  // Load semester data for editing
  useEffect(() => {
    if (id && semesters) {
      const semester = semesters.find((s: any) => s.id === id);
      if (semester) {
        setFormData({
          academic_year_id: semester.academic_year_id,
          type: semester.type || 'REGULAR',
          name: semester.name,
          sequence_number: semester.sequence_number || 1,
          sequence: semester.sequence_number || 1,
          start_date: semester.start_date,
          end_date: semester.end_date,
          is_active: semester.is_active,
        });
      }
    }
  }, [id, semesters]);

  const validateForm = (): boolean => {
    if (!formData.academic_year_id) {
      setFormError('Tahun ajaran harus dipilih');
      return false;
    }
    if (!formData.name.trim()) {
      setFormError('Nama semester harus diisi');
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
        // Update existing semester
        await academicFoundationApi.updateSemester(id, {
          name: formData.name,
          start_date: formData.start_date,
          end_date: formData.end_date,
        });
      } else {
        // Create new semester
        await academicFoundationApi.createSemester(formData);
      }
      navigate(`/dashboard/academic-foundation/academic-years/${academicYearId}`);
    } catch (error: any) {
      setFormError(error.response?.data?.message || 'Gagal menyimpan semester');
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
          onClick={() => navigate(`/dashboard/academic-foundation/academic-years/${academicYearId}`)}
          sx={{ mb: 2 }}
        >
          Kembali ke Daftar Semester
        </Button>
        <Typography variant="h4" component="h1">
          {id ? 'Edit Semester' : 'Tambah Semester Baru'}
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
                label="Nama Semester"
                fullWidth
                value={formData.name}
                onChange={(e) => setFormData({ ...formData, name: e.target.value })}
                placeholder="Contoh: Ganjil 2024/2025"
                error={!formData.name.trim()}
                helperText={!formData.name.trim() ? 'Nama semester harus diisi' : ''}
                required
              />
              
              <TextField
                label="Urutan Semester"
                fullWidth
                type="number"
                value={formData.sequence_number}
                onChange={(e) => setFormData({ ...formData, sequence_number: parseInt(e.target.value), sequence: parseInt(e.target.value) })}
                inputProps={{ min: 1 }}
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

              <Box sx={{ display: 'flex', gap: 2, alignItems: 'center' }}>
                <input
                  type="checkbox"
                  id="is-active"
                  checked={formData.is_active}
                  onChange={(e) => setFormData({ ...formData, is_active: e.target.checked })}
                />
                <label htmlFor="is-active">Status aktif</label>
              </Box>

              <Box sx={{ display: 'flex', gap: 2, justifyContent: 'flex-end', mt: 2 }}>
                <Button
                  variant="outlined"
                  onClick={() => navigate(`/dashboard/academic-foundation/academic-years/${academicYearId}`)}
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

export default SemesterFormPage;

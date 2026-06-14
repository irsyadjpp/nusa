/**
 * Curriculum Phase Form Page
 * Separate page for creating and editing curriculum phases
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
  Switch,
  FormControlLabel,
} from '@mui/material';
import {
  ArrowBack as ArrowBackIcon,
  Save as SaveIcon,
} from '@mui/icons-material';
import { useNavigate, useParams } from 'react-router-dom';
import { usePhases } from '@/services/queries/CPQueryService';
import { createPhase, updatePhase } from '@/api/cp';
import { useMutation, useQueryClient } from '@tanstack/react-query';

const PhaseFormPage: React.FC = () => {
  const navigate = useNavigate();
  const { id } = useParams<{ id?: string }>();

  const [formData, setFormData] = useState({
    code: '',
    name: '',
    description: '',
    grade_level_start: '',
    grade_level_end: '',
    is_active: true,
  });
  const [formError, setFormError] = useState('');
  const [submitting, setSubmitting] = useState(false);

  const { data: phases, isLoading, error } = usePhases();

  const queryClient = useQueryClient();

  const createMutation = useMutation({
    mutationFn: createPhase,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['phases'] });
      navigate('/dashboard/curriculum/phases');
    },
    onError: (error: any) => {
      setFormError(error.response?.data?.message || 'Gagal menyimpan fase');
      setSubmitting(false);
    },
  });

  const updateMutation = useMutation({
    mutationFn: ({ id, data }: { id: string; data: any }) => 
      updatePhase(id, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['phases'] });
      navigate('/dashboard/curriculum/phases');
    },
    onError: (error: any) => {
      setFormError(error.response?.data?.message || 'Gagal menyimpan fase');
      setSubmitting(false);
    },
  });

  // Load phase data for editing
  useEffect(() => {
    if (id && phases) {
      const phase = phases.find(p => p.id === id);
      if (phase) {
        setFormData({
          code: phase.code || '',
          name: phase.name || '',
          description: phase.description || '',
          grade_level_start: phase.grade_level_start?.toString() || '',
          grade_level_end: phase.grade_level_end?.toString() || '',
          is_active: phase.status === 'ACTIVE',
        });
      }
    }
  }, [id, phases]);

  const validateForm = (): boolean => {
    if (!formData.code.trim()) {
      setFormError('Kode fase harus diisi');
      return false;
    }
    if (!formData.name.trim()) {
      setFormError('Nama fase harus diisi');
      return false;
    }
    if (!formData.grade_level_start) {
      setFormError('Kelas mulai harus diisi');
      return false;
    }
    if (!formData.grade_level_end) {
      setFormError('Kelas selesai harus diisi');
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
        // Update existing phase
        await updateMutation.mutateAsync({
          id,
          data: {
            ...formData,
            grade_level_start: formData.grade_level_start ? parseInt(formData.grade_level_start) : undefined,
            grade_level_end: formData.grade_level_end ? parseInt(formData.grade_level_end) : undefined,
          },
        });
      } else {
        // Create new phase
        await createMutation.mutateAsync({
          ...formData,
          grade_level_start: formData.grade_level_start ? parseInt(formData.grade_level_start) : undefined,
          grade_level_end: formData.grade_level_end ? parseInt(formData.grade_level_end) : undefined,
        });
      }
    } catch (error) {
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

  if (error) {
    return (
      <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
        <Alert severity="error">
          Gagal memuat data fase
        </Alert>
      </Container>
    );
  }

  return (
    <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
      <Box sx={{ mb: 3 }}>
        <Button
          startIcon={<ArrowBackIcon />}
          onClick={() => navigate('/dashboard/curriculum/phases')}
          sx={{ mb: 2 }}
        >
          Kembali ke Daftar Fase
        </Button>
        <Typography variant="h4" component="h1">
          {id ? 'Edit Fase' : 'Tambah Fase Baru'}
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
                label="Kode Fase"
                fullWidth
                value={formData.code}
                onChange={(e) => setFormData({ ...formData, code: e.target.value })}
                placeholder="Contoh: A, B, C"
                error={!formData.code.trim()}
                helperText={!formData.code.trim() ? 'Kode fase harus diisi' : ''}
                required
              />
              
              <TextField
                label="Nama Fase"
                fullWidth
                value={formData.name}
                onChange={(e) => setFormData({ ...formData, name: e.target.value })}
                placeholder="Contoh: Fase A"
                error={!formData.name.trim()}
                helperText={!formData.name.trim() ? 'Nama fase harus diisi' : ''}
                required
              />
              
              <TextField
                label="Deskripsi"
                fullWidth
                multiline
                rows={3}
                value={formData.description}
                onChange={(e) => setFormData({ ...formData, description: e.target.value })}
                placeholder="Deskripsi singkat tentang fase ini"
              />

              <TextField
                label="Kelas Mulai"
                fullWidth
                value={formData.grade_level_start}
                onChange={(e) => setFormData({ ...formData, grade_level_start: e.target.value })}
                placeholder="Contoh: 1"
                error={!formData.grade_level_start}
                helperText={!formData.grade_level_start ? 'Kelas mulai harus diisi' : ''}
                required
              />

              <TextField
                label="Kelas Selesai"
                fullWidth
                value={formData.grade_level_end}
                onChange={(e) => setFormData({ ...formData, grade_level_end: e.target.value })}
                placeholder="Contoh: 2"
                error={!formData.grade_level_end}
                helperText={!formData.grade_level_end ? 'Kelas selesai harus diisi' : ''}
                required
              />

              <FormControlLabel
                control={
                  <Switch
                    checked={formData.is_active}
                    onChange={(e) => setFormData({ ...formData, is_active: e.target.checked })}
                  />
                }
                label="Status aktif"
              />

              <Box sx={{ display: 'flex', gap: 2, justifyContent: 'flex-end', mt: 2 }}>
                <Button
                  variant="outlined"
                  onClick={() => navigate('/dashboard/curriculum/phases')}
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

export default PhaseFormPage;

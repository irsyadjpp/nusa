/**
 * Curriculum Subject Form Page
 * Separate page for creating and editing curriculum subjects
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
import { useSubjects } from '@/services/queries/CPQueryService';
import { createSubject, updateSubject } from '@/api/cp';
import { useMutation, useQueryClient } from '@tanstack/react-query';

const SubjectFormPage: React.FC = () => {
  const navigate = useNavigate();
  const { id } = useParams<{ id?: string }>();

  const [formData, setFormData] = useState({
    code: '',
    name: '',
    description: '',
    is_active: true,
  });
  const [formError, setFormError] = useState('');
  const [submitting, setSubmitting] = useState(false);

  const { data: subjects, isLoading, error } = useSubjects();

  const queryClient = useQueryClient();

  const createMutation = useMutation({
    mutationFn: createSubject,
  });

  const updateMutation = useMutation({
    mutationFn: ({ id, data }: { id: string; data: any }) =>
      updateSubject(id, data),
  });

  // Load subject data for editing
  useEffect(() => {
    if (id && subjects) {
      const subject = subjects.find(s => s.id === id);
      if (subject) {
        setFormData({
          code: subject.code || '',
          name: subject.name || '',
          description: subject.description || '',
          is_active: subject.status === 'ACTIVE',
        });
      }
    }
  }, [id, subjects]);

  const validateForm = (): boolean => {
    if (!formData.code.trim()) {
      setFormError('Kode mata pelajaran harus diisi');
      return false;
    }
    if (!formData.name.trim()) {
      setFormError('Nama mata pelajaran harus diisi');
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
        // Update existing subject
        await updateMutation.mutateAsync({
          id,
          data: formData,
        });
      } else {
        // Create new subject
        await createMutation.mutateAsync(formData);
      }
      queryClient.invalidateQueries({ queryKey: ['subjects'] });
      navigate('/dashboard/curriculum/subjects');
    } catch (error: any) {
      setFormError(error.response?.data?.message || 'Gagal menyimpan mata pelajaran');
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
          Gagal memuat data mata pelajaran
        </Alert>
      </Container>
    );
  }

  return (
    <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
      <Box sx={{ mb: 3 }}>
        <Button
          startIcon={<ArrowBackIcon />}
          onClick={() => navigate('/dashboard/curriculum/subjects')}
          sx={{ mb: 2 }}
        >
          Kembali ke Daftar Mata Pelajaran
        </Button>
        <Typography variant="h4" component="h1">
          {id ? 'Edit Mata Pelajaran' : 'Tambah Mata Pelajaran Baru'}
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
                label="Kode Mata Pelajaran"
                fullWidth
                value={formData.code}
                onChange={(e) => setFormData({ ...formData, code: e.target.value })}
                placeholder="Contoh: MTK, IPA, IPS"
                error={!formData.code.trim()}
                helperText={!formData.code.trim() ? 'Kode mata pelajaran harus diisi' : ''}
                disabled={!!id}
                required
              />
              
              <TextField
                label="Nama Mata Pelajaran"
                fullWidth
                value={formData.name}
                onChange={(e) => setFormData({ ...formData, name: e.target.value })}
                placeholder="Contoh: Matematika"
                error={!formData.name.trim()}
                helperText={!formData.name.trim() ? 'Nama mata pelajaran harus diisi' : ''}
                required
              />
              
              <TextField
                label="Deskripsi"
                fullWidth
                multiline
                rows={3}
                value={formData.description}
                onChange={(e) => setFormData({ ...formData, description: e.target.value })}
                placeholder="Deskripsi singkat tentang mata pelajaran ini"
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
                  onClick={() => navigate('/dashboard/curriculum/subjects')}
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

export default SubjectFormPage;

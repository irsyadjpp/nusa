/**
 * Curriculum Element Form Page
 * Separate page for creating and editing curriculum elements
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
  Select,
  MenuItem,
  FormControl,
  InputLabel,
} from '@mui/material';
import {
  ArrowBack as ArrowBackIcon,
  Save as SaveIcon,
} from '@mui/icons-material';
import { useNavigate, useParams, useSearchParams } from 'react-router-dom';
import { useElementsByPhase, useSubjects, usePhases } from '@/services/queries/CPQueryService';
import { createElement, updateElement } from '@/api/cp';
import { useMutation, useQueryClient } from '@tanstack/react-query';

const ElementFormPage: React.FC = () => {
  const navigate = useNavigate();
  const { id } = useParams<{ id?: string }>();
  const [searchParams] = useSearchParams();
  const phaseId = searchParams.get('phaseId') || '';

  const [formData, setFormData] = useState({
    subject_id: '',
    phase_id: phaseId,
    code: '',
    name: '',
    description: '',
    is_active: true,
  });
  const [formError, setFormError] = useState('');
  const [submitting, setSubmitting] = useState(false);

  const { data: elements, isLoading, error } = useElementsByPhase(phaseId);
  const { data: subjects } = useSubjects();
  const { data: phases } = usePhases();

  const queryClient = useQueryClient();

  const createMutation = useMutation({
    mutationFn: createElement,
  });

  const updateMutation = useMutation({
    mutationFn: ({ id, data }: { id: string; data: any }) =>
      updateElement(id, data),
  });

  // Load element data for editing
  useEffect(() => {
    if (id && elements) {
      const element = elements.find(e => e.id === id);
      if (element) {
        setFormData({
          subject_id: element.subject_id || '',
          phase_id: element.phase_id || phaseId,
          code: element.code || '',
          name: element.name || '',
          description: element.description || '',
          is_active: element.is_active ?? true,
        });
      }
    }
  }, [id, elements, phaseId]);

  const validateForm = (): boolean => {
    if (!formData.subject_id) {
      setFormError('Mata pelajaran harus dipilih');
      return false;
    }
    if (!formData.phase_id) {
      setFormError('Fase harus dipilih');
      return false;
    }
    if (!formData.code.trim()) {
      setFormError('Kode elemen harus diisi');
      return false;
    }
    if (!formData.name.trim()) {
      setFormError('Nama elemen harus diisi');
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
        // Update existing element
        await updateMutation.mutateAsync({
          id,
          data: formData,
        });
      } else {
        // Create new element
        await createMutation.mutateAsync(formData);
      }
      queryClient.invalidateQueries({ queryKey: ['elements'] });
      navigate(`/dashboard/curriculum/elements?phaseId=${phaseId}`);
    } catch (error: any) {
      setFormError(error.response?.data?.message || 'Gagal menyimpan elemen');
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
          Gagal memuat data elemen
        </Alert>
      </Container>
    );
  }

  return (
    <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
      <Box sx={{ mb: 3 }}>
        <Button
          startIcon={<ArrowBackIcon />}
          onClick={() => navigate(`/dashboard/curriculum/elements?phaseId=${phaseId}`)}
          sx={{ mb: 2 }}
        >
          Kembali ke Daftar Elemen
        </Button>
        <Typography variant="h4" component="h1">
          {id ? 'Edit Elemen' : 'Tambah Elemen Baru'}
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
              
              <FormControl fullWidth>
                <InputLabel>Mata Pelajaran</InputLabel>
                <Select
                  value={formData.subject_id}
                  label="Mata Pelajaran"
                  onChange={(e) => setFormData({ ...formData, subject_id: e.target.value })}
                  required
                >
                  <MenuItem value="">Pilih Mata Pelajaran</MenuItem>
                  {subjects?.map((subject) => (
                    <MenuItem key={subject.id} value={subject.id}>
                      {subject.code} - {subject.name}
                    </MenuItem>
                  ))}
                </Select>
              </FormControl>
              
              <FormControl fullWidth>
                <InputLabel>Fase</InputLabel>
                <Select
                  value={formData.phase_id}
                  label="Fase"
                  onChange={(e) => setFormData({ ...formData, phase_id: e.target.value })}
                  required
                >
                  <MenuItem value="">Pilih Fase</MenuItem>
                  {phases?.map((phase) => (
                    <MenuItem key={phase.id} value={phase.id}>
                      {phase.code} - {phase.name}
                    </MenuItem>
                  ))}
                </Select>
              </FormControl>
              
              <TextField
                label="Kode Elemen"
                fullWidth
                value={formData.code}
                onChange={(e) => setFormData({ ...formData, code: e.target.value })}
                placeholder="Contoh: 1.1, 1.2"
                error={!formData.code.trim()}
                helperText={!formData.code.trim() ? 'Kode elemen harus diisi' : ''}
                required
              />
              
              <TextField
                label="Nama Elemen"
                fullWidth
                value={formData.name}
                onChange={(e) => setFormData({ ...formData, name: e.target.value })}
                placeholder="Contoh: Menghargai dan beriman"
                error={!formData.name.trim()}
                helperText={!formData.name.trim() ? 'Nama elemen harus diisi' : ''}
                required
              />
              
              <TextField
                label="Deskripsi"
                fullWidth
                multiline
                rows={3}
                value={formData.description}
                onChange={(e) => setFormData({ ...formData, description: e.target.value })}
                placeholder="Deskripsi singkat tentang elemen ini"
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
                  onClick={() => navigate(`/dashboard/curriculum/elements?phaseId=${phaseId}`)}
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

export default ElementFormPage;

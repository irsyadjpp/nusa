/**
 * Curriculum Subelement Form Page
 * Separate page for creating and editing curriculum subelements
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
import { useNavigate, useParams, useSearchParams } from 'react-router-dom';
import { useSubelementsByElement } from '@/services/queries/CPQueryService';
import { createSubelement, updateSubelement } from '@/api/cp';
import { useMutation } from '@tanstack/react-query';

const SubelementFormPage: React.FC = () => {
  const navigate = useNavigate();
  const { id } = useParams<{ id?: string }>();
  const [searchParams] = useSearchParams();
  const elementId = searchParams.get('elementId') || '';

  const [formData, setFormData] = useState({
    element_id: elementId,
    code: '',
    name: '',
    description: '',
    is_active: true,
  });
  const [formError, setFormError] = useState('');
  const [submitting, setSubmitting] = useState(false);

  const { data: subelements, isLoading, error } = useSubelementsByElement(elementId);

  const createMutation = useMutation({
    mutationFn: createSubelement,
  });

  const updateMutation = useMutation({
    mutationFn: ({ id, data }: { id: string; data: any }) =>
      updateSubelement(id, data),
  });

  // Load subelement data for editing
  useEffect(() => {
    if (id && subelements) {
      const subelement = subelements.find(s => s.id === id);
      if (subelement) {
        setFormData({
          element_id: subelement.element_id || elementId,
          code: subelement.code || '',
          name: subelement.name || '',
          description: subelement.description || '',
          is_active: subelement.is_active ?? true,
        });
      }
    }
  }, [id, subelements]);

  const validateForm = (): boolean => {
    if (!formData.element_id) {
      setFormError('Elemen harus dipilih');
      return false;
    }
    if (!formData.code.trim()) {
      setFormError('Kode sub-elemen harus diisi');
      return false;
    }
    if (!formData.name.trim()) {
      setFormError('Nama sub-elemen harus diisi');
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
        // Update existing subelement
        await updateMutation.mutateAsync({
          id,
          data: formData,
        });
      } else {
        // Create new subelement
        await createMutation.mutateAsync(formData);
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
          Gagal memuat data sub-elemen
        </Alert>
      </Container>
    );
  }

  return (
    <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
      <Box sx={{ mb: 3 }}>
        <Button
          startIcon={<ArrowBackIcon />}
          onClick={() => navigate(`/dashboard/curriculum/subelements?elementId=${elementId}`)}
          sx={{ mb: 2 }}
        >
          Kembali ke Daftar Sub-elemen
        </Button>
        <Typography variant="h4" component="h1">
          {id ? 'Edit Sub-elemen' : 'Tambah Sub-elemen Baru'}
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
                label="ID Elemen"
                fullWidth
                value={formData.element_id}
                onChange={(e) => setFormData({ ...formData, element_id: e.target.value })}
                disabled
                helperText="Elemen ditetapkan dari URL"
              />

              <TextField
                label="Kode Sub-elemen"
                fullWidth
                value={formData.code}
                onChange={(e) => setFormData({ ...formData, code: e.target.value })}
                placeholder="Contoh: 1.1.1, 1.1.2"
                error={!formData.code.trim()}
                helperText={!formData.code.trim() ? 'Kode sub-elemen harus diisi' : ''}
                required
              />
              
              <TextField
                label="Nama Sub-elemen"
                fullWidth
                value={formData.name}
                onChange={(e) => setFormData({ ...formData, name: e.target.value })}
                placeholder="Contoh: Menghargai dan beriman kepada Tuhan"
                error={!formData.name.trim()}
                helperText={!formData.name.trim() ? 'Nama sub-elemen harus diisi' : ''}
                required
              />
              
              <TextField
                label="Deskripsi"
                fullWidth
                multiline
                rows={3}
                value={formData.description}
                onChange={(e) => setFormData({ ...formData, description: e.target.value })}
                placeholder="Deskripsi singkat tentang sub-elemen ini"
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
                  onClick={() => navigate(`/dashboard/curriculum/subelements?elementId=${elementId}`)}
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

export default SubelementFormPage;

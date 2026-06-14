/**
 * Phase Edit Page
 * Edit an existing curriculum phase
 */

import React, { useState, useEffect } from 'react';
import {
  Box,
  Typography,
  Button,
  Card,
  CardContent,
  CircularProgress,
  Alert,
  TextField,
  Switch,
  FormControlLabel,
} from '@mui/material';
import {
  ArrowBack as ArrowBackIcon,
  Save as SaveIcon,
} from '@mui/icons-material';
import { useNavigate, useParams } from 'react-router-dom';
import { getPhaseById, updatePhase } from '@/api/cp';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';

const PhaseEditPage: React.FC = () => {
  const navigate = useNavigate();
  const { id } = useParams<{ id: string }>();
  const queryClient = useQueryClient();

  const [formData, setFormData] = useState({
    name: '',
    description: '',
    grade_level_start: '',
    grade_level_end: '',
    is_active: true,
  });

  const { data: phase, isLoading, error } = useQuery({
    queryKey: ['phase', id],
    queryFn: () => getPhaseById(id!),
    enabled: !!id,
  });

  // Handle data loading side effect (replaces onSuccess)
  useEffect(() => {
    if (phase) {
      setFormData({
        name: phase.name,
        description: phase.description || '',
        grade_level_start: phase.level || '',
        grade_level_end: phase.grade_range || '',
        is_active: phase.status === 'ACTIVE',
      });
    }
  }, [phase]);

  const updateMutation = useMutation({
    mutationFn: (data: any) => updatePhase(id!, data),
  });

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const data: any = {
      name: formData.name,
      description: formData.description || undefined,
      is_active: formData.is_active,
    };

    if (formData.grade_level_start) {
      data.grade_level_start = parseInt(formData.grade_level_start);
    }
    if (formData.grade_level_end) {
      data.grade_level_end = parseInt(formData.grade_level_end);
    }

    await updateMutation.mutateAsync(data);
    queryClient.invalidateQueries({ queryKey: ['phase', id] });
    queryClient.invalidateQueries({ queryKey: ['phases'] });
    navigate(`/dashboard/curriculum/phases/${id}`);
  };

  if (isLoading) {
    return (
      <Box sx={{ display: 'flex', justifyContent: 'center', p: 4 }}>
        <CircularProgress />
      </Box>
    );
  }

  if (error || !phase) {
    return <Alert severity="error">Error loading phase</Alert>;
  }

  return (
    <Box sx={{ p: 3, maxWidth: 800, margin: '0 auto' }}>
      <Box sx={{ display: 'flex', alignItems: 'center', mb: 3 }}>
        <Button
          startIcon={<ArrowBackIcon />}
          onClick={() => navigate(`/dashboard/curriculum/phases/${id}`)}
          sx={{ mr: 2 }}
        >
          Kembali
        </Button>
        <Typography variant="h4" component="h1">
          Edit Fase
        </Typography>
      </Box>

      <Card>
        <CardContent>
          <form onSubmit={handleSubmit}>
            <Box sx={{ display: 'flex', flexDirection: 'column', gap: 3 }}>
              <TextField
                label="Nama"
                fullWidth
                variant="outlined"
                value={formData.name}
                onChange={(e) => setFormData({ ...formData, name: e.target.value })}
                required
              />

              <TextField
                label="Deskripsi"
                fullWidth
                multiline
                rows={4}
                variant="outlined"
                value={formData.description}
                onChange={(e) => setFormData({ ...formData, description: e.target.value })}
              />

              <TextField
                label="Kelas Mulai"
                fullWidth
                type="number"
                variant="outlined"
                value={formData.grade_level_start}
                onChange={(e) => setFormData({ ...formData, grade_level_start: e.target.value })}
              />

              <TextField
                label="Kelas Akhir"
                fullWidth
                type="number"
                variant="outlined"
                value={formData.grade_level_end}
                onChange={(e) => setFormData({ ...formData, grade_level_end: e.target.value })}
              />

              <FormControlLabel
                control={
                  <Switch
                    checked={formData.is_active}
                    onChange={(e) => setFormData({ ...formData, is_active: e.target.checked })}
                  />
                }
                label="Aktif"
              />

              <Box sx={{ display: 'flex', gap: 2, justifyContent: 'flex-end' }}>
                <Button
                  variant="outlined"
                  onClick={() => navigate(`/dashboard/curriculum/phases/${id}`)}
                >
                  Batal
                </Button>
                <Button
                  type="submit"
                  variant="contained"
                  startIcon={<SaveIcon />}
                  disabled={updateMutation.isPending}
                >
                  {updateMutation.isPending ? 'Menyimpan...' : 'Simpan'}
                </Button>
              </Box>
            </Box>
          </form>
        </CardContent>
      </Card>
    </Box>
  );
};

export default PhaseEditPage;

/**
 * CP Edit Page
 * Edit an existing Curriculum Plan
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
import { getCPById, updateCP } from '@/api/cp';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';

const CPEditPage: React.FC = () => {
  const navigate = useNavigate();
  const { id } = useParams<{ id: string }>();
  const queryClient = useQueryClient();

  const [formData, setFormData] = useState({
    description: '',
    competency_code: '',
    learning_objectives: '{}',
    competency_standards: '{}',
    time_allocation_hours: '',
    hours_per_week: '',
    version: '',
    is_active: true,
  });

  const { data: cp, isLoading: isLoadingCP, error: cpError } = useQuery({
    queryKey: ['cp', id],
    queryFn: () => getCPById(id!),
    enabled: !!id,
  });

  // Handle data loading side effect (replaces onSuccess)
  useEffect(() => {
    if (cp) {
      setFormData({
        description: cp.description || '',
        competency_code: cp.competency_code || '',
        learning_objectives: typeof cp.learning_objectives === 'string'
          ? cp.learning_objectives
          : JSON.stringify(cp.learning_objectives || {}),
        competency_standards: typeof cp.competency_standards === 'string'
          ? cp.competency_standards
          : JSON.stringify(cp.competency_standards || {}),
        time_allocation_hours: cp.time_allocation_hours?.toString() || '',
        hours_per_week: cp.hours_per_week?.toString() || '',
        version: cp.version || '',
        is_active: cp.status === 'ACTIVE',
      });
    }
  }, [cp]);

  const updateMutation = useMutation({
    mutationFn: (data: any) => updateCP(id!, data),
  });

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    try {
      const data: any = {
        description: formData.description,
        competency_code: formData.competency_code || undefined,
        learning_objectives: JSON.parse(formData.learning_objectives),
        competency_standards: JSON.parse(formData.competency_standards),
        is_active: formData.is_active,
      };

      if (formData.time_allocation_hours) {
        data.time_allocation_hours = parseInt(formData.time_allocation_hours);
      }
      if (formData.hours_per_week) {
        data.hours_per_week = parseInt(formData.hours_per_week);
      }
      if (formData.version) {
        data.version = formData.version;
      }

      await updateMutation.mutateAsync(data);
      queryClient.invalidateQueries({ queryKey: ['cp', id] });
      queryClient.invalidateQueries({ queryKey: ['cps'] });
      navigate(`/cp/${id}`);
    } catch (error) {
      alert('Invalid JSON in learning objectives or competency standards');
    }
  };

  if (isLoadingCP) {
    return (
      <Box sx={{ display: 'flex', justifyContent: 'center', p: 4 }}>
        <CircularProgress />
      </Box>
    );
  }

  if (cpError || !cp) {
    return <Alert severity="error">Error loading CP</Alert>;
  }

  return (
    <Box sx={{ p: 3, maxWidth: 800, margin: '0 auto' }}>
      <Box sx={{ display: 'flex', alignItems: 'center', mb: 3 }}>
        <Button
          startIcon={<ArrowBackIcon />}
          onClick={() => navigate(`/cp/${id}`)}
          sx={{ mr: 2 }}
        >
          Kembali
        </Button>
        <Typography variant="h4" component="h1">
          Edit CP
        </Typography>
      </Box>

      <Card>
        <CardContent>
          <form onSubmit={handleSubmit}>
            <Box sx={{ display: 'flex', flexDirection: 'column', gap: 3 }}>
              <TextField
                label="Deskripsi"
                fullWidth
                multiline
                rows={4}
                variant="outlined"
                value={formData.description}
                onChange={(e) => setFormData({ ...formData, description: e.target.value })}
                required
              />

              <TextField
                label="Kompetensi Kode"
                fullWidth
                variant="outlined"
                value={formData.competency_code}
                onChange={(e) => setFormData({ ...formData, competency_code: e.target.value })}
              />

              <TextField
                label="Learning Objectives (JSON)"
                fullWidth
                multiline
                rows={4}
                variant="outlined"
                value={formData.learning_objectives}
                onChange={(e) => setFormData({ ...formData, learning_objectives: e.target.value })}
                required
              />

              <TextField
                label="Competency Standards (JSON)"
                fullWidth
                multiline
                rows={4}
                variant="outlined"
                value={formData.competency_standards}
                onChange={(e) => setFormData({ ...formData, competency_standards: e.target.value })}
                required
              />

              <TextField
                label="Alokasi Waktu (Jam)"
                fullWidth
                type="number"
                variant="outlined"
                value={formData.time_allocation_hours}
                onChange={(e) => setFormData({ ...formData, time_allocation_hours: e.target.value })}
              />

              <TextField
                label="Jam per Minggu"
                fullWidth
                type="number"
                variant="outlined"
                value={formData.hours_per_week}
                onChange={(e) => setFormData({ ...formData, hours_per_week: e.target.value })}
              />

              <TextField
                label="Versi"
                fullWidth
                variant="outlined"
                value={formData.version}
                onChange={(e) => setFormData({ ...formData, version: e.target.value })}
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
                  onClick={() => navigate(`/cp/${id}`)}
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

export default CPEditPage;

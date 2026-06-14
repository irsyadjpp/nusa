/**
 * Subject Edit Page
 * Edit an existing curriculum subject
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
import { getSubjectById, updateSubject } from '@/api/cp';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';

const SubjectEditPage: React.FC = () => {
  const navigate = useNavigate();
  const { id } = useParams<{ id: string }>();
  const queryClient = useQueryClient();

  const [formData, setFormData] = useState({
    name: '',
    description: '',
    is_active: true,
  });

  const { data: subject, isLoading, error } = useQuery({
    queryKey: ['subject', id],
    queryFn: () => getSubjectById(id!),
    enabled: !!id,
  });

  // Handle data loading side effect (replaces onSuccess)
  useEffect(() => {
    if (subject) {
      setFormData({
        name: subject.name || '',
        description: subject.description || '',
        is_active: subject.status === 'ACTIVE',
      });
    }
  }, [subject]);

  const updateMutation = useMutation({
    mutationFn: (data: any) => updateSubject(id!, data),
  });

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    await updateMutation.mutateAsync({
      name: formData.name,
      description: formData.description || undefined,
      is_active: formData.is_active,
    });
    queryClient.invalidateQueries({ queryKey: ['subject', id] });
    queryClient.invalidateQueries({ queryKey: ['subjects'] });
    navigate(`/dashboard/curriculum/subjects/${id}`);
  };

  if (isLoading) {
    return (
      <Box sx={{ display: 'flex', justifyContent: 'center', p: 4 }}>
        <CircularProgress />
      </Box>
    );
  }

  if (error || !subject) {
    return <Alert severity="error">Error loading subject</Alert>;
  }

  return (
    <Box sx={{ p: 3, maxWidth: 800, margin: '0 auto' }}>
      <Box sx={{ display: 'flex', alignItems: 'center', mb: 3 }}>
        <Button
          startIcon={<ArrowBackIcon />}
          onClick={() => navigate(`/dashboard/curriculum/subjects/${id}`)}
          sx={{ mr: 2 }}
        >
          Kembali
        </Button>
        <Typography variant="h4" component="h1">
          Edit Mata Pelajaran
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
                  onClick={() => navigate(`/dashboard/curriculum/subjects/${id}`)}
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

export default SubjectEditPage;

/**
 * Phases List Page
 * Display all curriculum phases with CRUD operations using page-based forms
 */

import React from 'react';
import {
  Box,
  Typography,
  Button,
  Card,
  CardContent,
  CardActions,
  Grid,
  CircularProgress,
  Alert,
} from '@mui/material';
import {
  Add as AddIcon,
  Edit as EditIcon,
  Delete as DeleteIcon,
} from '@mui/icons-material';
import { useNavigate } from 'react-router-dom';
import { usePhases } from '@/services/queries/CPQueryService';
import { deletePhase } from '@/api/cp';
import { CurriculumPhase } from '@/shared/types/domain';
import { useMutation, useQueryClient } from '@tanstack/react-query';

const PhasesListPage: React.FC = () => {
  const navigate = useNavigate();
  const queryClient = useQueryClient();

  const handleAdd = () => {
    navigate('/dashboard/curriculum/phases/new');
  };

  const handleEdit = (phaseId: string) => {
    navigate(`/dashboard/curriculum/phases/${phaseId}`);
  };

  const deleteMutation = useMutation({
    mutationFn: deletePhase,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['phases'] });
    },
  });

  const handleDelete = async (phase: CurriculumPhase) => {
    if (window.confirm(`Apakah Anda yakin ingin menghapus fase "${phase.name}"?`)) {
      try {
        await deleteMutation.mutateAsync(phase.id);
      } catch (error) {
        console.error('Error deleting phase:', error);
      }
    }
  };

  const { data: phases = [], isLoading, error } = usePhases();

  const handleViewDetail = (phase: CurriculumPhase) => {
    navigate(`/dashboard/curriculum/phases/${phase.id}`);
  };

  if (isLoading) {
    return (
      <Box sx={{ display: 'flex', justifyContent: 'center', p: 4 }}>
        <CircularProgress />
      </Box>
    );
  }

  if (error) {
    return <Alert severity="error">Error loading phases</Alert>;
  }

  return (
    <Box sx={{ p: 3 }}>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
        <Typography variant="h4">Fase</Typography>
        <Button variant="contained" startIcon={<AddIcon />} onClick={handleAdd}>
          Tambah Fase
        </Button>
      </Box>

      <Grid container spacing={3}>
        {phases.map((phase) => (
          <Grid size={{ xs: 12, md: 6, lg: 4 }} key={phase.id}>
            <Card>
              <CardContent>
                <Typography variant="h6" gutterBottom>
                  {phase.name}
                </Typography>
                <Typography variant="body2" color="text.secondary" gutterBottom>
                  Kode: {phase.code}
                </Typography>
                {phase.description && (
                  <Typography variant="body2" color="text.secondary">
                    {phase.description}
                  </Typography>
                )}
              </CardContent>
              <CardActions>
                <Button size="small" onClick={() => handleViewDetail(phase)}>
                  Detail
                </Button>
                <Button
                  size="small"
                  startIcon={<EditIcon />}
                  onClick={() => handleEdit(phase.id)}
                >
                  Edit
                </Button>
                <Button
                  size="small"
                  color="error"
                  startIcon={<DeleteIcon />}
                  onClick={() => handleDelete(phase)}
                >
                  Hapus
                </Button>
              </CardActions>
            </Card>
          </Grid>
        ))}
      </Grid>
    </Box>
  );
};

export default PhasesListPage;

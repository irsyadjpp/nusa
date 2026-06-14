/**
 * Phase Detail Page
 * Display detailed information about a curriculum phase
 */

import React from 'react';
import {
  Box,
  Typography,
  Button,
  Card,
  CardContent,
  CircularProgress,
  Alert,
  Chip,
  Divider,
} from '@mui/material';
import {
  ArrowBack as ArrowBackIcon,
  Edit as EditIcon,
} from '@mui/icons-material';
import { useNavigate, useParams } from 'react-router-dom';
import { getPhaseById } from '@/api/cp';
import { useQuery } from '@tanstack/react-query';

const PhaseDetailPage: React.FC = () => {
  const navigate = useNavigate();
  const { id } = useParams<{ id: string }>();

  const { data: phase, isLoading, error } = useQuery({
    queryKey: ['phase', id],
    queryFn: () => getPhaseById(id!),
    enabled: !!id,
  });

  if (isLoading) {
    return (
      <Box sx={{ display: 'flex', justifyContent: 'center', p: 4 }}>
        <CircularProgress />
      </Box>
    );
  }

  if (error) {
    return <Alert severity="error">Error loading phase</Alert>;
  }

  if (!phase) {
    return <Alert severity="info">Phase not found</Alert>;
  }

  return (
    <Box sx={{ p: 3 }}>
      <Box sx={{ display: 'flex', alignItems: 'center', mb: 3 }}>
        <Button
          startIcon={<ArrowBackIcon />}
          onClick={() => navigate('/dashboard/curriculum/phases')}
          sx={{ mr: 2 }}
        >
          Kembali
        </Button>
        <Typography variant="h4" component="h1">
          Detail Fase
        </Typography>
      </Box>

      <Card>
        <CardContent>
          <Box sx={{ mb: 3 }}>
            <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 2 }}>
              <Typography variant="h5">{phase.name}</Typography>
              <Chip
                label={phase.status === 'ACTIVE' ? 'Aktif' : 'Tidak Aktif'}
                color={phase.status === 'ACTIVE' ? 'success' : 'default'}
              />
            </Box>
            <Divider sx={{ my: 2 }} />
          </Box>

          <Box sx={{ mb: 2 }}>
            <Typography variant="subtitle2" color="text.secondary">
              Kode
            </Typography>
            <Typography variant="body1">{phase.code}</Typography>
          </Box>

          {phase.description && (
            <Box sx={{ mb: 2 }}>
              <Typography variant="subtitle2" color="text.secondary">
                Deskripsi
              </Typography>
              <Typography variant="body1">{phase.description}</Typography>
            </Box>
          )}

          <Box sx={{ mb: 2 }}>
            <Typography variant="subtitle2" color="text.secondary">
              Tingkat Kelas
            </Typography>
            <Typography variant="body1">{phase.grade_range}</Typography>
          </Box>

          <Box sx={{ mt: 3 }}>
            <Button
              variant="contained"
              startIcon={<EditIcon />}
              onClick={() => navigate(`/dashboard/curriculum/phases/${phase.id}/edit`)}
            >
              Edit
            </Button>
          </Box>
        </CardContent>
      </Card>
    </Box>
  );
};

export default PhaseDetailPage;

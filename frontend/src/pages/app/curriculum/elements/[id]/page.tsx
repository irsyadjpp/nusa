/**
 * Element Detail Page
 * Display detailed information about a curriculum element
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
import { getElementById } from '@/api/cp';
import { useQuery } from '@tanstack/react-query';
import { CurriculumElement } from '@/shared/types/domain';

const ElementDetailPage: React.FC = () => {
  const navigate = useNavigate();
  const { id } = useParams<{ id: string }>();

  const { data: element, isLoading, error } = useQuery({
    queryKey: ['element', id],
    queryFn: () => getElementById(id!),
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
    return <Alert severity="error">Error loading element</Alert>;
  }

  if (!element) {
    return <Alert severity="info">Element not found</Alert>;
  }

  const typedElement = element as unknown as CurriculumElement;

  return (
    <Box sx={{ p: 3 }}>
      <Box sx={{ display: 'flex', alignItems: 'center', mb: 3 }}>
        <Button
          startIcon={<ArrowBackIcon />}
          onClick={() => navigate('/dashboard/curriculum/elements')}
          sx={{ mr: 2 }}
        >
          Kembali
        </Button>
        <Typography variant="h4" component="h1">
          Detail Elemen
        </Typography>
      </Box>

      <Card>
        <CardContent>
          <Box sx={{ mb: 3 }}>
            <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 2 }}>
              <Typography variant="h5">{typedElement.name}</Typography>
              <Chip
                label={typedElement.status === 'ACTIVE' ? 'Aktif' : 'Tidak Aktif'}
                color={typedElement.status === 'ACTIVE' ? 'success' : 'default'}
              />
            </Box>
            <Divider sx={{ my: 2 }} />
          </Box>

          <Box sx={{ mb: 2 }}>
            <Typography variant="subtitle2" color="text.secondary">
              Kode
            </Typography>
            <Typography variant="body1">{typedElement.code}</Typography>
          </Box>

          <Box sx={{ mb: 2 }}>
            <Typography variant="subtitle2" color="text.secondary">
              Subject ID
            </Typography>
            <Typography variant="body1">{typedElement.subject_id}</Typography>
          </Box>

          <Box sx={{ mb: 2 }}>
            <Typography variant="subtitle2" color="text.secondary">
              Phase ID
            </Typography>
            <Typography variant="body1">{typedElement.phase_id}</Typography>
          </Box>

          {typedElement.description && (
            <Box sx={{ mb: 2 }}>
              <Typography variant="subtitle2" color="text.secondary">
                Deskripsi
              </Typography>
              <Typography variant="body1">{typedElement.description}</Typography>
            </Box>
          )}

          <Box sx={{ mt: 3 }}>
            <Button
              variant="contained"
              startIcon={<EditIcon />}
              onClick={() => navigate(`/dashboard/curriculum/elements/${typedElement.id}/edit`)}
            >
              Edit
            </Button>
          </Box>
        </CardContent>
      </Card>
    </Box>
  );
};

export default ElementDetailPage;

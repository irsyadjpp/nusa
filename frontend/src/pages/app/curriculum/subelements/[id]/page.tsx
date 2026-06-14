/**
 * SubElement Detail Page
 * Display detailed information about a curriculum subelement
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
import { getSubelementById } from '@/api/cp';
import { useQuery } from '@tanstack/react-query';

const SubElementDetailPage: React.FC = () => {
  const navigate = useNavigate();
  const { id } = useParams<{ id: string }>();

  const { data: subelement, isLoading, error } = useQuery({
    queryKey: ['subelement', id],
    queryFn: () => getSubelementById(id!),
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
    return <Alert severity="error">Error loading subelement</Alert>;
  }

  if (!subelement) {
    return <Alert severity="info">Subelement not found</Alert>;
  }

  return (
    <Box sx={{ p: 3 }}>
      <Box sx={{ display: 'flex', alignItems: 'center', mb: 3 }}>
        <Button
          startIcon={<ArrowBackIcon />}
          onClick={() => navigate('/dashboard/curriculum/subelements')}
          sx={{ mr: 2 }}
        >
          Kembali
        </Button>
        <Typography variant="h4" component="h1">
          Detail Subelemen
        </Typography>
      </Box>

      <Card>
        <CardContent>
          <Box sx={{ mb: 3 }}>
            <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 2 }}>
              <Typography variant="h5">{subelement.name}</Typography>
              <Chip
                label={subelement.status === 'ACTIVE' ? 'Aktif' : 'Tidak Aktif'}
                color={subelement.status === 'ACTIVE' ? 'success' : 'default'}
              />
            </Box>
            <Divider sx={{ my: 2 }} />
          </Box>

          <Box sx={{ mb: 2 }}>
            <Typography variant="subtitle2" color="text.secondary">
              Kode
            </Typography>
            <Typography variant="body1">{subelement.code}</Typography>
          </Box>

          <Box sx={{ mb: 2 }}>
            <Typography variant="subtitle2" color="text.secondary">
              Element ID
            </Typography>
            <Typography variant="body1">{subelement.element_id}</Typography>
          </Box>

          {subelement.description && (
            <Box sx={{ mb: 2 }}>
              <Typography variant="subtitle2" color="text.secondary">
                Deskripsi
              </Typography>
              <Typography variant="body1">{subelement.description}</Typography>
            </Box>
          )}

          <Box sx={{ mt: 3 }}>
            <Button
              variant="contained"
              startIcon={<EditIcon />}
              onClick={() => navigate(`/dashboard/curriculum/subelements/${subelement.id}/edit`)}
            >
              Edit
            </Button>
          </Box>
        </CardContent>
      </Card>
    </Box>
  );
};

export default SubElementDetailPage;

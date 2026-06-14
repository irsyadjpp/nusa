/**
 * Subject Detail Page
 * Display detailed information about a curriculum subject
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
import { getSubjectById } from '@/api/cp';
import { useQuery } from '@tanstack/react-query';

const SubjectDetailPage: React.FC = () => {
  const navigate = useNavigate();
  const { id } = useParams<{ id: string }>();

  const { data: subject, isLoading, error } = useQuery({
    queryKey: ['subject', id],
    queryFn: () => getSubjectById(id!),
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
    return <Alert severity="error">Error loading subject</Alert>;
  }

  if (!subject) {
    return <Alert severity="info">Subject not found</Alert>;
  }

  return (
    <Box sx={{ p: 3 }}>
      <Box sx={{ display: 'flex', alignItems: 'center', mb: 3 }}>
        <Button
          startIcon={<ArrowBackIcon />}
          onClick={() => navigate('/dashboard/curriculum/subjects')}
          sx={{ mr: 2 }}
        >
          Kembali
        </Button>
        <Typography variant="h4" component="h1">
          Detail Mata Pelajaran
        </Typography>
      </Box>

      <Card>
        <CardContent>
          <Box sx={{ mb: 3 }}>
            <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 2 }}>
              <Typography variant="h5">{subject.name}</Typography>
              <Chip
                label={subject.status === 'ACTIVE' ? 'Aktif' : 'Tidak Aktif'}
                color={subject.status === 'ACTIVE' ? 'success' : 'default'}
              />
            </Box>
            <Divider sx={{ my: 2 }} />
          </Box>

          <Box sx={{ mb: 2 }}>
            <Typography variant="subtitle2" color="text.secondary">
              Kode
            </Typography>
            <Typography variant="body1">{subject.code}</Typography>
          </Box>

          {subject.description && (
            <Box sx={{ mb: 2 }}>
              <Typography variant="subtitle2" color="text.secondary">
                Deskripsi
              </Typography>
              <Typography variant="body1">{subject.description}</Typography>
            </Box>
          )}

          <Box sx={{ mt: 3 }}>
            <Button
              variant="contained"
              startIcon={<EditIcon />}
              onClick={() => navigate(`/dashboard/curriculum/subjects/${subject.id}/edit`)}
            >
              Edit
            </Button>
          </Box>
        </CardContent>
      </Card>
    </Box>
  );
};

export default SubjectDetailPage;

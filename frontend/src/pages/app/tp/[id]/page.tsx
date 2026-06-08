/**
 * TP Detail Page
 * Detail view for Teaching Plan (TP)
 */

import React, { useState, useEffect } from 'react';
import {
  Box,
  Typography,
  Button,
  Card,
  CardContent,
  Chip,
  Grid,
  CircularProgress,
  Alert,
  Divider,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
} from '@mui/material';
import {
  ArrowBack,
  Edit,
  Delete,
  Visibility,
} from '@mui/icons-material';
import { useNavigate, useParams } from 'react-router-dom';
import { getTPById, deleteTP } from '@/api/tp';
import { TP } from '@/api/tp';
import KKTPCriteriaDisplay from '@/components/kktp/KKTPCriteriaDisplay';

const TPDetailPage: React.FC = () => {
  const navigate = useNavigate();
  const { id } = useParams<{ id: string }>();
  const [tp, setTP] = useState<TP | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [deleteDialogOpen, setDeleteDialogOpen] = useState(false);
  const [deleting, setDeleting] = useState(false);

  useEffect(() => {
    if (id) {
      loadTP(id);
    }
  }, [id]);

  const loadTP = async (tpId: string) => {
    setLoading(true);
    setError(null);
    try {
      const data = await getTPById(tpId);
      setTP(data);
    } catch (err: any) {
      setError(err.message || 'Gagal memuat data TP');
    } finally {
      setLoading(false);
    }
  };

  const handleDelete = async () => {
    if (!id) return;
    setDeleting(true);
    try {
      await deleteTP(id);
      navigate('/tp');
    } catch (err: any) {
      setError(err.message || 'Gagal menghapus TP');
      setDeleteDialogOpen(false);
    } finally {
      setDeleting(false);
    }
  };

  const getStatusColor = (status: string): 'success' | 'warning' | 'error' | 'info' | 'default' => {
    switch (status) {
      case 'APPROVED':
        return 'success';
      case 'PENDING':
        return 'warning';
      case 'REJECTED':
        return 'error';
      case 'DRAFT':
        return 'info';
      default:
        return 'default';
    }
  };

  if (loading) {
    return (
      <Box sx={{ display: 'flex', justifyContent: 'center', py: 4 }}>
        <CircularProgress />
      </Box>
    );
  }

  if (error || !tp) {
    return (
      <Alert severity="error">
        {error || 'TP tidak ditemukan'}
      </Alert>
    );
  }

  return (
    <Box>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
        <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
          <Button
            variant="outlined"
            startIcon={<ArrowBack />}
            onClick={() => navigate('/tp')}
          >
            Kembali
          </Button>
          <Typography variant="h4">Detail TP</Typography>
        </Box>
        <Box sx={{ display: 'flex', gap: 1 }}>
          <Button
            variant="outlined"
            startIcon={<Edit />}
            onClick={() => navigate(`/tp/${tp.id}/edit`)}
          >
            Edit
          </Button>
          <Button
            variant="outlined"
            color="error"
            startIcon={<Delete />}
            onClick={() => setDeleteDialogOpen(true)}
          >
            Hapus
          </Button>
        </Box>
      </Box>

      <Grid container spacing={3}>
        <Grid item xs={12} md={8}>
          <Card sx={{ mb: 3 }}>
            <CardContent>
              <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-start', mb: 2 }}>
                <Typography variant="h5">{tp.title}</Typography>
                <Chip
                  label={tp.status}
                  color={getStatusColor(tp.status)}
                />
              </Box>

              <Divider sx={{ my: 2 }} />

              <Grid container spacing={2}>
                <Grid item xs={6}>
                  <Typography variant="caption" color="text.secondary">
                    Urutan TP
                  </Typography>
                  <Typography variant="body1">
                    #{tp.sequence_number}
                  </Typography>
                </Grid>
                <Grid item xs={6}>
                  <Typography variant="caption" color="text.secondary">
                    Estimasi Minggu
                  </Typography>
                  <Typography variant="body1">
                    {tp.estimated_weeks} minggu
                  </Typography>
                </Grid>
                <Grid item xs={12}>
                  <Typography variant="caption" color="text.secondary">
                    Tujuan Pembelajaran
                  </Typography>
                  <Typography variant="body1">
                    {tp.learning_objectives}
                  </Typography>
                </Grid>
                <Grid item xs={12}>
                  <Typography variant="caption" color="text.secondary">
                    Alokasi Waktu
                  </Typography>
                  <Typography variant="body1">
                    {tp.time_allocation}
                  </Typography>
                </Grid>
                <Grid item xs={12}>
                  <Typography variant="caption" color="text.secondary">
                    Prasyarat
                  </Typography>
                  <Typography variant="body1">
                    {tp.prerequisites || '-'}
                  </Typography>
                </Grid>
              </Grid>
            </CardContent>
          </Card>

          {tp.success_criteria && (
            <Card>
              <CardContent>
                <Typography variant="h6" gutterBottom>
                  Kriteria Ketuntasan Tujuan Pembelajaran (KKTP)
                </Typography>
                <KKTPCriteriaDisplay data={tp.success_criteria} />
              </CardContent>
            </Card>
          )}
        </Grid>

        <Grid item xs={12} md={4}>
          <Card>
            <CardContent>
              <Typography variant="h6" gutterBottom>
                Informasi
              </Typography>
              <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
                <Box>
                  <Typography variant="caption" color="text.secondary">
                    ID TP
                  </Typography>
                  <Typography variant="body2">
                    {tp.id}
                  </Typography>
                </Box>
                <Box>
                  <Typography variant="caption" color="text.secondary">
                    ID Set TP
                  </Typography>
                  <Typography variant="body2">
                    {tp.tp_set_id}
                  </Typography>
                </Box>
                <Box>
                  <Typography variant="caption" color="text.secondary">
                    ID CP
                  </Typography>
                  <Typography variant="body2">
                    {tp.cp_id}
                  </Typography>
                </Box>
                <Box>
                  <Typography variant="caption" color="text.secondary">
                    ID Mata Pelajaran
                  </Typography>
                  <Typography variant="body2">
                    {tp.subject_id}
                  </Typography>
                </Box>
                <Box>
                  <Typography variant="caption" color="text.secondary">
                    ID Fase
                  </Typography>
                  <Typography variant="body2">
                    {tp.phase_id}
                  </Typography>
                </Box>
                <Box>
                  <Typography variant="caption" color="text.secondary">
                    ID Elemen
                  </Typography>
                  <Typography variant="body2">
                    {tp.element_id}
                  </Typography>
                </Box>
                <Box>
                  <Typography variant="caption" color="text.secondary">
                    ID Sub-elemen
                  </Typography>
                  <Typography variant="body2">
                    {tp.subelement_id || '-'}
                  </Typography>
                </Box>
                <Divider />
                <Box>
                  <Typography variant="caption" color="text.secondary">
                    Dibuat
                  </Typography>
                  <Typography variant="body2">
                    {new Date(tp.created_at).toLocaleString('id-ID')}
                  </Typography>
                </Box>
                <Box>
                  <Typography variant="caption" color="text.secondary">
                    Diperbarui
                  </Typography>
                  <Typography variant="body2">
                    {new Date(tp.updated_at).toLocaleString('id-ID')}
                  </Typography>
                </Box>
              </Box>
            </CardContent>
          </Card>
        </Grid>
      </Grid>

      <Dialog open={deleteDialogOpen} onClose={() => setDeleteDialogOpen(false)}>
        <DialogTitle>Hapus TP</DialogTitle>
        <DialogContent>
          Apakah Anda yakin ingin menghapus TP ini? Tindakan ini tidak dapat dibatalkan.
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setDeleteDialogOpen(false)}>Batal</Button>
          <Button
            variant="contained"
            color="error"
            onClick={handleDelete}
            disabled={deleting}
          >
            {deleting ? 'Menghapus...' : 'Hapus'}
          </Button>
        </DialogActions>
      </Dialog>
    </Box>
  );
};

export default TPDetailPage;

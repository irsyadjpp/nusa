/**
 * TP Detail Page - MIGRATED TO TANSTACK QUERY
 * Detail view for Teaching Plan (TP)
 */

import React, { useState } from 'react';
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
  History,
} from '@mui/icons-material';
import { useNavigate, useParams } from 'react-router-dom';
import { useTP } from '@/services/queries/TPQueryService';
import { useDeleteTP } from '@/services/commands/TPCommandService';
import { TPStatus } from '@/shared/types/domain';
import KKTPCriteriaDisplay from '@/components/kktp/KKTPCriteriaDisplay';
import TPVersionHistory from '@/components/tp/TPVersionHistory';

const TPDetailPage: React.FC = () => {
  const navigate = useNavigate();
  const { id } = useParams<{ id: string }>();
  const [deleteDialogOpen, setDeleteDialogOpen] = useState(false);
  const [versionHistoryOpen, setVersionHistoryOpen] = useState(false);

  // ✅ Using TanStack Query hook instead of manual state management
  const {
    data: tp,
    isLoading,
    error
  } = useTP(id!);

  const deleteMutation = useDeleteTP({
    onSuccess: () => {
      navigate('/tp');
    },
  });

  const handleDelete = async () => {
    if (!id) return;
    deleteMutation.mutate(id);
  };

  const getStatusColor = (status: TPStatus): 'success' | 'warning' | 'error' | 'info' | 'default' => {
    switch (status) {
      case 'APPROVED':
        return 'success';
      case 'UNDER_REVIEW':
        return 'warning';
      case 'REJECTED':
        return 'error';
      case 'DRAFT':
        return 'info';
      case 'ARCHIVED':
        return 'default';
      default:
        return 'default';
    }
  };

  const getStatusLabel = (status: TPStatus): string => {
    switch (status) {
      case 'APPROVED': return 'Disetujui';
      case 'UNDER_REVIEW': return 'Dalam Review';
      case 'REJECTED': return 'Ditolak';
      case 'DRAFT': return 'Draft';
      case 'ARCHIVED': return 'Diarsipkan';
      default: return status;
    }
  };

  if (isLoading) {
    return (
      <Box sx={{ display: 'flex', justifyContent: 'center', py: 4 }}>
        <CircularProgress />
      </Box>
    );
  }

  if (error || !tp) {
    return (
      <Alert severity="error">
        {error?.message || 'TP tidak ditemukan'}
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
            startIcon={<History />}
            onClick={() => setVersionHistoryOpen(true)}
          >
            Riwayat Versi
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
        <Grid size={{ xs: 12, md: 8 }}>
          <Card sx={{ mb: 3 }}>
            <CardContent>
              <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-start', mb: 2 }}>
                <Typography variant="h5">{tp.title}</Typography>
                <Chip
                  label={getStatusLabel(tp.status)}
                  color={getStatusColor(tp.status)}
                />
              </Box>

              <Divider sx={{ my: 2 }} />

              <Grid container spacing={2}>
                <Grid size={{ xs: 6 }}>
                  <Typography variant="caption" color="text.secondary">
                    Urutan TP
                  </Typography>
                  <Typography variant="body1">
                    #{tp.sequence_number}
                  </Typography>
                </Grid>
                <Grid size={{ xs: 6 }}>
                  <Typography variant="caption" color="text.secondary">
                    Estimasi Minggu
                  </Typography>
                  <Typography variant="body1">
                    {tp.estimated_weeks} minggu
                  </Typography>
                </Grid>
                <Grid size={{ xs: 12 }}>
                  <Typography variant="caption" color="text.secondary">
                    Tujuan Pembelajaran
                  </Typography>
                  <Typography variant="body1">
                    {typeof tp.learning_objectives === 'string' ? tp.learning_objectives : JSON.stringify(tp.learning_objectives)}
                  </Typography>
                </Grid>
                <Grid size={{ xs: 12 }}>
                  <Typography variant="caption" color="text.secondary">
                    Alokasi Waktu
                  </Typography>
                  <Typography variant="body1">
                    {typeof tp.time_allocation === 'string' ? tp.time_allocation : JSON.stringify(tp.time_allocation)}
                  </Typography>
                </Grid>
                <Grid size={{ xs: 12 }}>
                  <Typography variant="caption" color="text.secondary">
                    Prasyarat
                  </Typography>
                  <Typography variant="body1">
                    {typeof tp.prerequisites === 'string' ? tp.prerequisites : JSON.stringify(tp.prerequisites) || '-'}
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
                <KKTPCriteriaDisplay data={tp.success_criteria as any} />
              </CardContent>
            </Card>
          )}
        </Grid>

        <Grid size={{ xs: 12, md: 4 }}>
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
            disabled={deleteMutation.isPending}
          >
            {deleteMutation.isPending ? 'Menghapus...' : 'Hapus'}
          </Button>
        </DialogActions>
      </Dialog>

      <TPVersionHistory
        tpSetId={tp.tp_set_id}
        open={versionHistoryOpen}
        onClose={() => setVersionHistoryOpen(false)}
      />
    </Box>
  );
};

export default TPDetailPage;

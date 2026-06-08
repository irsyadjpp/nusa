/**
 * Assessment Detail Page
 * Detail view for Assessment
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
  CheckCircle,
  Cancel,
} from '@mui/icons-material';
import { useNavigate, useParams } from 'react-router-dom';
import { getAssessmentById, deleteAssessment, approveAssessment, rejectAssessment } from '@/api/assessment';
import { Assessment } from '@/api/assessment';
import AssessmentReview from '@/components/assessment/AssessmentReview';

const AssessmentDetailPage: React.FC = () => {
  const navigate = useNavigate();
  const { id } = useParams<{ id: string }>();
  const [assessment, setAssessment] = useState<Assessment | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [deleteDialogOpen, setDeleteDialogOpen] = useState(false);
  const [deleting, setDeleting] = useState(false);
  const [approving, setApproving] = useState(false);

  useEffect(() => {
    if (id) {
      loadAssessment(id);
    }
  }, [id]);

  const loadAssessment = async (assessmentId: string) => {
    setLoading(true);
    setError(null);
    try {
      const data = await getAssessmentById(assessmentId);
      setAssessment(data);
    } catch (err: any) {
      setError(err.message || 'Gagal memuat data Asesmen');
    } finally {
      setLoading(false);
    }
  };

  const handleDelete = async () => {
    if (!id) return;
    setDeleting(true);
    try {
      await deleteAssessment(id);
      navigate('/assessment');
    } catch (err: any) {
      setError(err.message || 'Gagal menghapus asesmen');
      setDeleteDialogOpen(false);
    } finally {
      setDeleting(false);
    }
  };

  const handleApprove = async () => {
    if (!id) return;
    setApproving(true);
    try {
      await approveAssessment(id);
      loadAssessment(id);
    } catch (err: any) {
      setError(err.message || 'Gagal menyetujui asesmen');
    } finally {
      setApproving(false);
    }
  };

  const handleReject = async () => {
    if (!id) return;
    setApproving(true);
    try {
      await rejectAssessment(id);
      loadAssessment(id);
    } catch (err: any) {
      setError(err.message || 'Gagal menolak asesmen');
    } finally {
      setApproving(false);
    }
  };

  if (loading) {
    return (
      <Box sx={{ display: 'flex', justifyContent: 'center', py: 4 }}>
        <CircularProgress />
      </Box>
    );
  }

  if (error || !assessment) {
    return (
      <Alert severity="error">
        {error || 'Asesmen tidak ditemukan'}
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
            onClick={() => navigate('/assessment')}
          >
            Kembali
          </Button>
          <Typography variant="h4">Detail Asesmen</Typography>
        </Box>
        <Box sx={{ display: 'flex', gap: 1 }}>
          {assessment.status === 'DRAFT' && (
            <Button
              variant="outlined"
              startIcon={<Edit />}
              onClick={() => navigate(`/assessment/${assessment.id}/edit`)}
            >
              Edit
            </Button>
          )}
          {assessment.status === 'PENDING' && (
            <>
              <Button
                variant="contained"
                color="success"
                startIcon={<CheckCircle />}
                onClick={handleApprove}
                disabled={approving}
              >
                Setujui
              </Button>
              <Button
                variant="contained"
                color="error"
                startIcon={<Cancel />}
                onClick={handleReject}
                disabled={approving}
              >
                Tolak
              </Button>
            </>
          )}
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
          <AssessmentReview
            assessment={assessment}
            onApprove={assessment.status === 'PENDING' ? handleApprove : undefined}
            onReject={assessment.status === 'PENDING' ? handleReject : undefined}
            onEdit={assessment.status === 'DRAFT' ? (id) => navigate(`/assessment/${id}/edit`) : undefined}
            onDelete={(id) => {
              setDeleteDialogOpen(true);
            }}
            loading={approving || deleting}
          />
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
                    ID Asesmen
                  </Typography>
                  <Typography variant="body2">
                    {assessment.id}
                  </Typography>
                </Box>
                <Box>
                  <Typography variant="caption" color="text.secondary">
                    ID TP
                  </Typography>
                  <Typography variant="body2">
                    {assessment.tp_id}
                  </Typography>
                </Box>
                <Box>
                  <Typography variant="caption" color="text.secondary">
                    Versi TP
                  </Typography>
                  <Typography variant="body2">
                    {assessment.tp_version_no}
                  </Typography>
                </Box>
                <Box>
                  <Typography variant="caption" color="text.secondary">
                    ID User
                  </Typography>
                  <Typography variant="body2">
                    {assessment.user_id}
                  </Typography>
                </Box>
                <Box>
                  <Typography variant="caption" color="text.secondary">
                    Tipe Asesmen
                  </Typography>
                  <Typography variant="body2">
                    {assessment.assessment_type}
                  </Typography>
                </Box>
                <Box>
                  <Typography variant="caption" color="text.secondary">
                    Versi
                  </Typography>
                  <Typography variant="body2">
                    {assessment.version_no}
                  </Typography>
                </Box>
                {assessment.ai_confidence_score && (
                  <Box>
                    <Typography variant="caption" color="text.secondary">
                      AI Confidence
                    </Typography>
                    <Typography variant="body2">
                      {assessment.ai_confidence_score.toFixed(2)}
                    </Typography>
                  </Box>
                )}
                <Divider />
                <Box>
                  <Typography variant="caption" color="text.secondary">
                    Dibuat
                  </Typography>
                  <Typography variant="body2">
                    {new Date(assessment.created_at).toLocaleString('id-ID')}
                  </Typography>
                </Box>
                <Box>
                  <Typography variant="caption" color="text.secondary">
                    Diperbarui
                  </Typography>
                  <Typography variant="body2">
                    {new Date(assessment.updated_at).toLocaleString('id-ID')}
                  </Typography>
                </Box>
                {assessment.approved_at && (
                  <Box>
                    <Typography variant="caption" color="text.secondary">
                      Disetujui
                    </Typography>
                    <Typography variant="body2">
                      {new Date(assessment.approved_at).toLocaleString('id-ID')}
                    </Typography>
                  </Box>
                )}
              </Box>
            </CardContent>
          </Card>
        </Grid>
      </Grid>

      <Dialog open={deleteDialogOpen} onClose={() => setDeleteDialogOpen(false)}>
        <DialogTitle>Hapus Asesmen</DialogTitle>
        <DialogContent>
          Apakah Anda yakin ingin menghapus asesmen ini? Tindakan ini tidak dapat dibatalkan.
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

export default AssessmentDetailPage;

/**
 * Evidence Detail Page
 * Detail view for Evidence
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
  ArrowForward,
} from '@mui/icons-material';
import { useNavigate, useParams } from 'react-router-dom';
import { getEvidenceById, deleteEvidence } from '@/api/evidence';
import { Evidence } from '@/shared/types/domain';
import { getEvaluationsByEvidence } from '@/api/evaluation';
// import EvidenceReview from '@/components/evidence/EvidenceReview'; // TODO: Implement
// import EvaluationForm from '@/components/evidence/EvaluationForm'; // TODO: Implement
// import RevisionHistory from '@/components/evidence/RevisionHistory'; // TODO: Implement

const EvidenceDetailPage: React.FC = () => {
  const navigate = useNavigate();
  const { id } = useParams<{ id: string }>();
  const [evidence, setEvidence] = useState<Evidence | null>(null);
  const [evaluations, setEvaluations] = useState<any[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [deleteDialogOpen, setDeleteDialogOpen] = useState(false);
  const [deleting, setDeleting] = useState(false);
  const [showEvaluationForm, setShowEvaluationForm] = useState(false);

  useEffect(() => {
    if (id) {
      loadData(id);
    }
  }, [id]);

  const loadData = async (evidenceId: string) => {
    setLoading(true);
    setError(null);
    try {
      const [evidenceData, evaluationsData] = await Promise.all([
        getEvidenceById(evidenceId),
        getEvaluationsByEvidence(evidenceId),
      ]);
      setEvidence(evidenceData);
      setEvaluations(evaluationsData);
    } catch (err: any) {
      setError(err.message || 'Gagal memuat data Bukti');
    } finally {
      setLoading(false);
    }
  };

  const handleDelete = async () => {
    if (!id) return;
    setDeleting(true);
    try {
      await deleteEvidence(id);
      navigate('/evidence');
    } catch (err: any) {
      setError(err.message || 'Gagal menghapus bukti');
      setDeleteDialogOpen(false);
    } finally {
      setDeleting(false);
    }
  };

  const handleGoToEvaluation = () => {
    navigate('/evaluation', { state: { evidenceId: id } });
  };

  if (loading) {
    return (
      <Box sx={{ display: 'flex', justifyContent: 'center', py: 4 }}>
        <CircularProgress />
      </Box>
    );
  }

  if (error || !evidence) {
    return (
      <Alert severity="error">
        {error || 'Bukti tidak ditemukan'}
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
            onClick={() => navigate('/evidence')}
          >
            Kembali
          </Button>
          <Typography variant="h4">Detail Bukti</Typography>
        </Box>
        <Box sx={{ display: 'flex', gap: 1 }}>
          {evidence.status === 'SUBMITTED' && (
            <>
              <Button
                variant="contained"
                startIcon={<ArrowForward />}
                onClick={handleGoToEvaluation}
              >
                Evaluate
              </Button>
            </>
          )}
          <Button
            variant="outlined"
            startIcon={<Edit />}
            onClick={() => navigate(`/evidence/${evidence.id}/edit`)}
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

      <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 3 }}>
        <Box sx={{ width: { xs: '100%', md: '66.67%' } }}>
          <Card>
            <CardContent>
              <Typography variant="h6" gutterBottom>
                Evidence Review
              </Typography>
              <Alert severity="info">
                Evidence Review component - TODO: Implement
              </Alert>
              <Typography variant="body2" sx={{ mt: 2 }}>
                Evidence ID: {evidence.id}
              </Typography>
            </CardContent>
          </Card>

          {showEvaluationForm && (
            <Box sx={{ mt: 3 }}>
              <Card>
                <CardContent>
                  <Typography variant="h6" gutterBottom>
                    Evaluation Form
                  </Typography>
                  <Alert severity="info">
                    Evaluation Form component - TODO: Implement
                  </Alert>
                  <Button onClick={() => setShowEvaluationForm(false)}>
                    Cancel
                  </Button>
                </CardContent>
              </Card>
            </Box>
          )}

          {evaluations.length > 0 && (
            <Box sx={{ mt: 3 }}>
              <Card>
                <CardContent>
                  <Typography variant="h6" gutterBottom>
                    Revision History
                  </Typography>
                  <Typography variant="body2" color="text.secondary">
                    {evaluations.length} evaluation(s)
                  </Typography>
                  <Typography variant="caption" sx={{ mt: 1 }}>
                    Revision History component - TODO: Implement
                  </Typography>
                </CardContent>
              </Card>
            </Box>
          )}
        </Box>

        <Box sx={{ width: { xs: '100%', md: '33.33%' } }}>
          <Box sx={{ display: 'flex', flexDirection: 'column', gap: 3 }}>
            <Box>
              <Typography variant="h6" gutterBottom>
                Informasi
              </Typography>
              <Box sx={{ p: 2, border: 1, borderColor: 'divider', borderRadius: 1 }}>
                <Box sx={{ mb: 2 }}>
                  <Typography variant="caption" color="text.secondary">
                    ID Bukti
                  </Typography>
                  <Typography variant="body2">
                    {evidence.id}
                  </Typography>
                </Box>
                <Box sx={{ mb: 2 }}>
                  <Typography variant="caption" color="text.secondary">
                    ID Siswa
                  </Typography>
                  <Typography variant="body2">
                    {evidence.student_id}
                  </Typography>
                </Box>
                <Box sx={{ mb: 2 }}>
                  <Typography variant="caption" color="text.secondary">
                    ID Asesmen
                  </Typography>
                  <Typography variant="body2">
                    {evidence.assessment_id}
                  </Typography>
                </Box>
                <Box sx={{ mb: 2 }}>
                  <Typography variant="caption" color="text.secondary">
                    Tipe Bukti
                  </Typography>
                  <Typography variant="body2">
                    {evidence.evidence_type}
                  </Typography>
                </Box>
                <Divider sx={{ my: 2 }} />
                <Box sx={{ mb: 2 }}>
                  <Typography variant="caption" color="text.secondary">
                    Dibuat
                  </Typography>
                  <Typography variant="body2">
                    {new Date(evidence.created_at).toLocaleString('id-ID')}
                  </Typography>
                </Box>
                <Box>
                  <Typography variant="caption" color="text.secondary">
                    Diperbarui
                  </Typography>
                  <Typography variant="body2">
                    {new Date(evidence.updated_at).toLocaleString('id-ID')}
                  </Typography>
                </Box>
              </Box>
            </Box>

            {evidence.teacher_notes && (
              <Box>
                <Typography variant="h6" gutterBottom>
                  Catatan Guru
                </Typography>
                <Box sx={{ p: 2, border: 1, borderColor: 'divider', borderRadius: 1 }}>
                  <Typography variant="body2">
                    {evidence.teacher_notes}
                  </Typography>
                </Box>
              </Box>
            )}
          </Box>
        </Box>
      </Box>

      <Dialog open={deleteDialogOpen} onClose={() => setDeleteDialogOpen(false)}>
        <DialogTitle>Hapus Bukti</DialogTitle>
        <DialogContent>
          Apakah Anda yakin ingin menghapus bukti ini? Tindakan ini tidak dapat dibatalkan.
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

export default EvidenceDetailPage;

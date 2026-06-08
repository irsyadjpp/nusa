/**
 * Evidence Detail Page
 * Detail view for Evidence
 */

import React, { useState, useEffect } from 'react';
import {
  Box,
  Typography,
  Button,
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
} from '@mui/icons-material';
import { useNavigate, useParams } from 'react-router-dom';
import { getEvidenceById, deleteEvidence } from '@/api/evidence';
import { getEvaluationsByEvidence } from '@/api/evaluation';
import { Evidence } from '@/api/evidence';
import EvidenceReview from '@/components/evidence/EvidenceReview';
import EvaluationForm from '@/components/evidence/EvaluationForm';
import RevisionHistory from '@/components/evidence/RevisionHistory';

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

  const handleEvaluationSubmit = async (values: any) => {
    if (!id) return;
    try {
      // Create evaluation
      await fetch(`/api/evaluations`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          evidence_id: id,
          ...values,
        }),
      });
      loadData(id);
      setShowEvaluationForm(false);
    } catch (err: any) {
      setError(err.message || 'Gagal menyimpan evaluasi');
    }
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
            <Button
              variant="contained"
              onClick={() => setShowEvaluationForm(true)}
            >
              Evaluasi
            </Button>
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

      <Grid container spacing={3}>
        <Grid item xs={12} md={8}>
          <EvidenceReview
            evidence={evidence}
            onView={() => {}}
            onEdit={(id) => navigate(`/evidence/${id}/edit`)}
            onDelete={(id) => setDeleteDialogOpen(true)}
          />

          {showEvaluationForm && (
            <Box sx={{ mt: 3 }}>
              <EvaluationForm
                onSubmit={handleEvaluationSubmit}
                onCancel={() => setShowEvaluationForm(false)}
              />
            </Box>
          )}

          {evaluations.length > 0 && (
            <Box sx={{ mt: 3 }}>
              <RevisionHistory evaluations={evaluations} />
            </Box>
          )}
        </Grid>

        <Grid item xs={12} md={4}>
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
        </Grid>
      </Grid>

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

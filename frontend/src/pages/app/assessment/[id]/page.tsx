/**
 * Assessment Detail Page - MIGRATED TO TANSTACK QUERY
 * Detail view for Assessment
 */

import React, { useState } from 'react';
import {
  Box,
  Typography,
  Button,
  Card,
  CardContent,
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
  NavigateNext,
  Description,
  AssignmentTurnedIn,
} from '@mui/icons-material';
import { useNavigate, useParams } from 'react-router-dom';
import { useAssessment } from '@/services/queries/AssessmentQueryService';
import { useDeleteAssessment } from '@/services/commands/AssessmentCommandService';
import AssessmentReview from '@/components/assessment/AssessmentReview';

const AssessmentDetailPage: React.FC = () => {
  const navigate = useNavigate();
  const { id } = useParams<{ id: string }>();
  const [deleteDialogOpen, setDeleteDialogOpen] = useState(false);

  // ✅ Using TanStack Query hooks instead of manual state management
  const { 
    data: assessment, 
    isLoading, 
    error,
    refetch
  } = useAssessment(id!);

  const deleteMutation = useDeleteAssessment({
    onSuccess: () => {
      navigate('/assessment');
    },
  });

  const handleDelete = async () => {
    if (!id) return;
    deleteMutation.mutate(id);
  };

  const handleApprove = async () => {
    // This would typically use a command mutation, but the API functions may not exist yet
    // For now, we'll keep the existing approach but could be enhanced later
    if (!id || !assessment) return;
    try {
      // TODO: Replace with proper TanStack Query mutation when command service is available
      // For now, we call the API directly and refetch
      await fetch(`/api/assessment/${id}/approve`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ user_id: assessment.user_id || 'current-user' }),
      });
      refetch();
    } catch (err: any) {
      console.error('Failed to approve assessment:', err);
    }
  };

  const handleReject = async () => {
    // This would typically use a command mutation, but the API functions may not exist yet
    // For now, we'll keep the existing approach but could be enhanced later
    if (!id || !assessment) return;
    try {
      // TODO: Replace with proper TanStack Query mutation when command service is available
      // For now, we call the API directly and refetch
      await fetch(`/api/assessment/${id}/reject`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ user_id: assessment.user_id || 'current-user' }),
      });
      refetch();
    } catch (err: any) {
      console.error('Failed to reject assessment:', err);
    }
  };

  if (isLoading) {
    return (
      <Box sx={{ display: 'flex', justifyContent: 'center', py: 4 }}>
        <CircularProgress />
      </Box>
    );
  }

  if (error || !assessment) {
    return (
      <Alert severity="error">
        {error?.message || 'Asesmen tidak ditemukan'}
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
            <>
              <Button
                variant="outlined"
                startIcon={<Edit />}
                onClick={() => navigate(`/assessment/${assessment.id}/edit`)}
              >
                Edit
              </Button>
              <Button
                variant="outlined"
                startIcon={<Description />}
                onClick={() => navigate('/rubric/create', { state: { assessmentId: assessment.id } })}
              >
                Create Rubric
              </Button>
              <Button
                variant="outlined"
                startIcon={<AssignmentTurnedIn />}
                onClick={() => navigate('/evidence/upload', { state: { assessmentId: assessment.id } })}
              >
                Upload Evidence
              </Button>
            </>
          )}
          {assessment.status === 'UNDER_REVIEW' && (
            <>
              <Button
                variant="contained"
                color="success"
                startIcon={<CheckCircle />}
                onClick={handleApprove}
              >
                Setujui
              </Button>
              <Button
                variant="contained"
                color="error"
                startIcon={<Cancel />}
                onClick={handleReject}
              >
                Tolak
              </Button>
            </>
          )}
          {assessment.status === 'APPROVED' && (
            <>
              <Button
                variant="outlined"
                startIcon={<Description />}
                onClick={() => navigate('/rubric/create', { state: { assessmentId: assessment.id } })}
              >
                Create Rubric
              </Button>
              <Button
                variant="outlined"
                startIcon={<AssignmentTurnedIn />}
                onClick={() => navigate('/evidence/upload', { state: { assessmentId: assessment.id } })}
              >
                Upload Evidence
              </Button>
              <Button
                variant="outlined"
                startIcon={<NavigateNext />}
                onClick={() => navigate('/evidence')}
              >
                View Evidence
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
        <Grid size={{ xs: 12, md: 8 }}>
          <AssessmentReview
            assessment={{ ...assessment, tp_title: assessment.tp_title || '' } as any}
            onApprove={assessment.status === 'UNDER_REVIEW' ? handleApprove : undefined}
            onReject={assessment.status === 'UNDER_REVIEW' ? handleReject : undefined}
            onEdit={assessment.status === 'DRAFT' ? (id) => navigate(`/assessment/${id}/edit`) : undefined}
            onDelete={() => {
              setDeleteDialogOpen(true);
            }}
            loading={deleteMutation.isPending}
          />
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
            disabled={deleteMutation.isPending}
          >
            {deleteMutation.isPending ? 'Menghapus...' : 'Hapus'}
          </Button>
        </DialogActions>
      </Dialog>
    </Box>
  );
};

export default AssessmentDetailPage;

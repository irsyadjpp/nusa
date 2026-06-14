/**
 * Assessment Review Component
 * Review component for assessments with approval workflow
 */

import React from 'react';
import {
  Box,
  Typography,
  Card,
  CardContent,
  Chip,
  Button,
  Divider,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  CircularProgress,
} from '@mui/material';
import {
  CheckCircle,
  Cancel,
  Visibility,
  Edit,
  Delete,
} from '@mui/icons-material';
import SuccessCriteriaSnapshot from './SuccessCriteriaSnapshot';

interface Assessment {
  id: string;
  tp_id: string;
  tp_title: string;
  tp_version_no: number;
  success_criteria_snapshot: any;
  user_id: string;
  user_name: string;
  assessment_type: string;
  status: string;
  assessment_items: any;
  answer_key: any;
  scoring_guidelines: any;
  created_at: string;
  updated_at: string;
}

interface AssessmentReviewProps {
  assessment: Assessment;
  onApprove?: (id: string) => void;
  onReject?: (id: string) => void;
  onEdit?: (id: string) => void;
  onDelete?: (id: string) => void;
  loading?: boolean;
}

const AssessmentReview: React.FC<AssessmentReviewProps> = ({
  assessment,
  onApprove,
  onReject,
  onEdit,
  onDelete,
  loading = false,
}) => {
  const [previewOpen, setPreviewOpen] = React.useState(false);

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

  const getAssessmentTypeLabel = (type: string): string => {
    switch (type) {
      case 'FORMATIVE':
        return 'Formatif';
      case 'SUMMATIVE':
        return 'Sumatif';
      default:
        return type;
    }
  };

  return (
    <Card>
      <CardContent>
        <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-start', mb: 2 }}>
          <Box>
            <Typography variant="h6" gutterBottom>
              {assessment.tp_title}
            </Typography>
            <Typography variant="body2" color="text.secondary">
              Dibuat oleh: {assessment.user_name}
            </Typography>
          </Box>
          <Chip
            label={assessment.status}
            color={getStatusColor(assessment.status)}
            size="small"
          />
        </Box>

        <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 2, mb: 2 }}>
          <Box sx={{ width: { xs: '50%', sm: '33.33%' } }}>
            <Typography variant="caption" color="text.secondary">
              Tipe Asesmen
            </Typography>
            <Typography variant="body2">
              {getAssessmentTypeLabel(assessment.assessment_type)}
            </Typography>
          </Box>
          <Box sx={{ width: { xs: '50%', sm: '33.33%' } }}>
            <Typography variant="caption" color="text.secondary">
              Versi TP
            </Typography>
            <Typography variant="body2">
              {assessment.tp_version_no}
            </Typography>
          </Box>
          <Box sx={{ width: { xs: '50%', sm: '33.33%' } }}>
            <Typography variant="caption" color="text.secondary">
              Dibuat
            </Typography>
            <Typography variant="body2">
              {new Date(assessment.created_at).toLocaleDateString('id-ID')}
            </Typography>
          </Box>
        </Box>

        <Divider sx={{ my: 2 }} />

        <Box sx={{ mb: 2 }}>
          <Box sx={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between', mb: 1 }}>
            <Typography variant="subtitle2" fontWeight="bold">
              Snapshot Kriteria Ketuntasan
            </Typography>
            <Button
              size="small"
              startIcon={<Visibility />}
              onClick={() => setPreviewOpen(true)}
            >
              Lihat Detail
            </Button>
          </Box>
        </Box>

        <Box sx={{ display: 'flex', gap: 1, justifyContent: 'flex-end' }}>
          {onEdit && assessment.status === 'DRAFT' && (
            <Button
              size="small"
              variant="outlined"
              startIcon={<Edit />}
              onClick={() => onEdit(assessment.id)}
              disabled={loading}
            >
              Edit
            </Button>
          )}
          {onApprove && assessment.status === 'PENDING' && (
            <Button
              size="small"
              variant="contained"
              color="success"
              startIcon={loading ? <CircularProgress size={16} /> : <CheckCircle />}
              onClick={() => onApprove(assessment.id)}
              disabled={loading}
            >
              Setujui
            </Button>
          )}
          {onReject && assessment.status === 'PENDING' && (
            <Button
              size="small"
              variant="contained"
              color="error"
              startIcon={<Cancel />}
              onClick={() => onReject(assessment.id)}
              disabled={loading}
            >
              Tolak
            </Button>
          )}
          {onDelete && (
            <Button
              size="small"
              variant="outlined"
              color="error"
              startIcon={<Delete />}
              onClick={() => onDelete(assessment.id)}
              disabled={loading}
            >
              Hapus
            </Button>
          )}
        </Box>

        {/* Preview Dialog */}
        <Dialog
          open={previewOpen}
          onClose={() => setPreviewOpen(false)}
          maxWidth="md"
          fullWidth
        >
          <DialogTitle>Detail Snapshot Kriteria Ketuntasan</DialogTitle>
          <DialogContent>
            <SuccessCriteriaSnapshot
              snapshot={assessment.success_criteria_snapshot}
              tpVersion={assessment.tp_version_no}
            />
          </DialogContent>
          <DialogActions>
            <Button onClick={() => setPreviewOpen(false)}>Tutup</Button>
          </DialogActions>
        </Dialog>
      </CardContent>
    </Card>
  );
};

export default AssessmentReview;
